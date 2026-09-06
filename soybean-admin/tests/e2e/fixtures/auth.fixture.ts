import crypto from 'node:crypto';
import process from 'node:process';
import { test as base, expect } from '@playwright/test';

type AuthFixtures = {
  loginAsAdmin: () => Promise<void>;
};

/** Decode an RFC 4648 base32 string (no padding) into raw bytes. */
function base32Decode(input: string): Buffer {
  const alphabet = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ234567';
  const clean = input.replace(/=+$/, '').toUpperCase();
  let bits = '';
  for (const char of clean) {
    const value = alphabet.indexOf(char);
    if (value === -1) continue;
    bits += value.toString(2).padStart(5, '0');
  }
  const bytes: number[] = [];
  for (let i = 0; i + 8 <= bits.length; i += 8) {
    bytes.push(Number.parseInt(bits.slice(i, i + 8), 2));
  }
  return Buffer.from(bytes);
}

/** Generate an RFC 6238 TOTP code (HMAC-SHA1, 30s step, 6 digits) - same algorithm the backend validates with (pquerna/otp defaults). */
function generateTotp(secret: string, timeStepSeconds = 30, digits = 6): string {
  const key = base32Decode(secret);
  const counter = Math.floor(Date.now() / 1000 / timeStepSeconds);
  const counterBuf = Buffer.alloc(8);
  counterBuf.writeBigUInt64BE(BigInt(counter));
  const hmac = crypto.createHmac('sha1', key).update(counterBuf).digest();
  const offset = hmac[hmac.length - 1] & 0xf;
  const code =
    ((hmac[offset] & 0x7f) << 24) |
    ((hmac[offset + 1] & 0xff) << 16) |
    ((hmac[offset + 2] & 0xff) << 8) |
    (hmac[offset + 3] & 0xff);
  return (code % 10 ** digits).toString().padStart(digits, '0');
}

// The admin account is provisioned once per CI run and its TOTP secret is
// only ever shown on first login (enrollment). Subsequent logins within the
// same run hit the "verify" stage instead, so the secret captured during
// enrollment is cached here to keep computing valid codes.
let cachedSecret: string | null = null;

export const test = base.extend<AuthFixtures>({
  loginAsAdmin: async ({ page, baseURL }, use) => {
    await use(async () => {
      const username = process.env.E2E_ADMIN_USER || 'admin';
      const password = process.env.E2E_ADMIN_PASS || 'admin123456';
      const loginPath = `${baseURL}/login/pwd-login`;

      await page.goto(loginPath, { waitUntil: 'domcontentloaded' });
      await page.getByPlaceholder(/Please enter user name|请输入用户名/).fill(username);
      await page.getByPlaceholder(/Please enter password|请输入密码/).fill(password);
      await page.getByRole('button', { name: /Confirm|确定|确认/ }).click();

      const secretCode = page.locator('.mfa-secret-value code');
      const mfaCodeInput = page.getByPlaceholder(/enter your 6-digit code/i);
      await mfaCodeInput.waitFor({ state: 'visible', timeout: 10_000 });

      if (await secretCode.isVisible().catch(() => false)) {
        cachedSecret = (await secretCode.textContent())?.trim() || null;
      }

      if (!cachedSecret) {
        throw new Error('No cached TOTP secret available to complete admin login verification.');
      }

      await mfaCodeInput.fill(generateTotp(cachedSecret));
      await page.getByRole('button', { name: /Confirm|确定|确认/ }).click();

      await page.waitForURL(/\/devices/, { timeout: 20_000 });
      await expect(page.getByText('West454').or(page.getByText('Rustdesk Api Server'))).toBeVisible();
    });
  }
});

export { expect };
