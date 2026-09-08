<script setup lang="ts">
import { computed, reactive, ref } from 'vue';
import QRCode from 'qrcode';
import { useMessage } from 'naive-ui';
import { $t } from '@/locales';
import { useNaiveForm } from '@/hooks/common/form';
import { useAuthStore } from '@/store/modules/auth';

defineOptions({
  name: 'PwdLogin'
});

const authStore = useAuthStore();
const message = useMessage();
const { formRef, validate } = useNaiveForm();

type Stage = 'password' | 'enroll' | 'verify';

const stage = ref<Stage>('password');
const challenge = ref('');
const qrDataUrl = ref('');
const enrollSecret = ref('');
const mfaCode = ref('');

const model: Api.Form.LoginForm = reactive({
  username: '',
  password: ''
});

const rules = computed<Record<keyof Api.Form.LoginForm, App.Global.FormRule[]>>(() => {
  return {
    username: [
      {
        required: true,
        message: 'Username is required'
      }
    ],
    password: [
      {
        required: true,
        message: 'Password is required'
      }
    ]
  };
});

async function handleSubmit() {
  await validate();
  const { data, error } = await authStore.requestLoginChallenge(model);
  if (error || !data) {
    return;
  }

  challenge.value = data.challenge;

  if (data.stage === 'enroll' && data.secret && data.otpauth_url) {
    enrollSecret.value = data.secret;
    qrDataUrl.value = await QRCode.toDataURL(data.otpauth_url, { width: 220, margin: 1 });
    stage.value = 'enroll';
  } else {
    stage.value = 'verify';
  }
}

async function handleVerify() {
  if (!mfaCode.value.trim()) {
    return;
  }
  await authStore.completeLogin({ challenge: challenge.value, code: mfaCode.value.trim() });
}

function backToPassword() {
  stage.value = 'password';
  mfaCode.value = '';
  challenge.value = '';
  qrDataUrl.value = '';
  enrollSecret.value = '';
}

function copySecret() {
  if (navigator?.clipboard?.writeText) {
    void navigator.clipboard.writeText(enrollSecret.value).then(() => {
      message.success($t('page.login.pwdLogin.mfaCopied'));
    });
  }
}
</script>

<template>
  <NForm v-if="stage === 'password'" ref="formRef" :model="model" :rules="rules" size="large" :show-label="false" class="auth-form">
    <NFormItem path="username">
      <NInput v-model:value="model.username" :placeholder="$t('page.login.common.userNamePlaceholder')">
        <template #prefix><SvgIcon icon="solar:user-rounded-linear" /></template>
      </NInput>
    </NFormItem>
    <NFormItem path="password">
      <NInput
        v-model:value="model.password"
        type="password"
        show-password-on="click"
        :placeholder="$t('page.login.common.passwordPlaceholder')"
        @keyup.enter="handleSubmit"
      >
        <template #prefix><SvgIcon icon="solar:lock-keyhole-minimalistic-linear" /></template>
      </NInput>
    </NFormItem>
    <div class="form-meta">
      <NCheckbox>{{ $t('page.login.pwdLogin.rememberMe') }}</NCheckbox>
    </div>
    <NButton
      attr-type="submit"
      type="primary"
      size="large"
      block
      :loading="authStore.loginLoading"
      class="submit-button"
      @click="handleSubmit"
    >
      <span class="inline-flex items-center gap-8px">
        {{ $t('common.confirm') }}
        <SvgIcon icon="solar:arrow-right-linear" />
      </span>
    </NButton>
  </NForm>

  <div v-else-if="stage === 'enroll'" class="mfa-panel">
    <h3>{{ $t('page.login.pwdLogin.mfaEnrollTitle') }}</h3>
    <p>{{ $t('page.login.pwdLogin.mfaEnrollHint') }}</p>
    <img v-if="qrDataUrl" :src="qrDataUrl" width="220" height="220" class="mfa-qr" alt="TOTP QR code" />
    <div class="mfa-secret">
      <span>{{ $t('page.login.pwdLogin.mfaSecretLabel') }}</span>
      <button type="button" class="mfa-secret-value" @click="copySecret">
        <code>{{ enrollSecret }}</code>
        <SvgIcon icon="solar:copy-linear" />
      </button>
    </div>
    <NInput
      v-model:value="mfaCode"
      size="large"
      class="mfa-code-input"
      :placeholder="$t('page.login.pwdLogin.mfaCodePlaceholder')"
      @keyup.enter="handleVerify"
    />
    <NButton type="primary" size="large" block :loading="authStore.loginLoading" class="submit-button" @click="handleVerify">
      {{ $t('common.confirm') }}
    </NButton>
    <button type="button" class="mfa-back" @click="backToPassword">{{ $t('page.login.pwdLogin.mfaBack') }}</button>
  </div>

  <div v-else class="mfa-panel">
    <h3>{{ $t('page.login.pwdLogin.mfaVerifyTitle') }}</h3>
    <p>{{ $t('page.login.pwdLogin.mfaVerifyHint') }}</p>
    <NInput
      v-model:value="mfaCode"
      size="large"
      class="mfa-code-input"
      :placeholder="$t('page.login.pwdLogin.mfaCodePlaceholder')"
      autofocus
      @keyup.enter="handleVerify"
    >
      <template #prefix><SvgIcon icon="solar:shield-keyhole-linear" /></template>
    </NInput>
    <NButton type="primary" size="large" block :loading="authStore.loginLoading" class="submit-button" @click="handleVerify">
      <span class="inline-flex items-center gap-8px">
        {{ $t('common.confirm') }}
        <SvgIcon icon="solar:arrow-right-linear" />
      </span>
    </NButton>
    <button type="button" class="mfa-back" @click="backToPassword">{{ $t('page.login.pwdLogin.mfaBack') }}</button>
  </div>
</template>

<style scoped>
.auth-form :deep(.n-form-item) {
  margin-bottom: 5px;
}

.auth-form :deep(.n-input) {
  min-height: 48px;
  background: var(--surface) !important;
}

.auth-form :deep(.n-input__prefix) {
  margin-right: 8px;
  color: var(--accent);
  font-size: 18px;
}

.form-meta {
  display: flex;
  align-items: center;
  margin: 4px 0 20px;
}

.submit-button {
  min-height: 49px;
  font-size: 14px;
  letter-spacing: 0.015em;
}

.mfa-panel {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 12px;
  text-align: center;
}

.mfa-panel h3 {
  margin: 0;
  color: var(--text-strong);
  font-size: 18px;
  font-weight: 760;
}

.mfa-panel p {
  margin: 0;
  color: var(--text-muted);
  font-size: 12.5px;
  line-height: 1.6;
}

.mfa-qr {
  border: 2px solid var(--surface-border);
  border-radius: var(--radius-sm);
  background: #fff;
  padding: 10px;
}

.mfa-secret {
  display: flex;
  flex-direction: column;
  gap: 6px;
  width: 100%;
}

.mfa-secret > span {
  color: var(--text-muted);
  font-size: 10px;
  font-weight: 700;
  letter-spacing: 0.06em;
  text-transform: uppercase;
}

.mfa-secret-value {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 8px;
  border: 2px solid var(--surface-border);
  border-radius: var(--radius-sm);
  background: var(--surface);
  color: var(--text-strong);
  padding: 10px 12px;
}

.mfa-secret-value code {
  overflow: hidden;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 12px;
  letter-spacing: 0.04em;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.mfa-code-input {
  width: 100%;
}

.mfa-code-input :deep(.n-input) {
  min-height: 48px;
  background: var(--surface) !important;
  text-align: center;
  letter-spacing: 0.3em;
}

.mfa-back {
  color: var(--text-muted);
  font-size: 11px;
  font-weight: 650;
  letter-spacing: 0.04em;
  text-decoration: underline;
  text-underline-offset: 3px;
}
</style>
