<script setup lang="ts">
import { computed } from 'vue';
import type { Component } from 'vue';
import { $t } from '@/locales';
import { useAppStore } from '@/store/modules/app';
import { useThemeStore } from '@/store/modules/theme';
import { loginModuleRecord } from '@/constants/app';
import PwdLogin from './modules/pwd-login.vue';

interface Props {
  /** The login module */
  module?: UnionKey.LoginModule;
}

const props = defineProps<Props>();

const appStore = useAppStore();
const themeStore = useThemeStore();

interface LoginModule {
  label: string;
  component: Component;
}

const moduleMap: Record<UnionKey.LoginModule, LoginModule> = {
  'pwd-login': { label: loginModuleRecord['pwd-login'], component: PwdLogin }
};

const activeModule = computed(() => moduleMap[props.module || 'pwd-login']);

</script>

<template>
  <div class="login-shell">
    <div class="login-grid" aria-hidden="true"></div>
    <section class="login-story">
      <div class="login-brand">
        <span class="brand-mark"><SvgIcon icon="solar:monitor-smartphone-bold-duotone" /></span>
        <span>
          <strong>West454</strong>
          <small>Remote operations</small>
        </span>
      </div>

      <div class="story-copy">
        <span class="eyebrow"><span class="status-dot"></span> Private infrastructure</span>
        <h1>Remote access that feels <span>effortless.</span></h1>
        <p>Your family fleet, protected by your server and organized in one focused workspace.</p>
        <div class="story-points">
          <span><SvgIcon icon="solar:shield-check-bold-duotone" /> Self-hosted identity</span>
          <span><SvgIcon icon="solar:lock-keyhole-minimalistic-bold-duotone" /> Encrypted connections</span>
          <span><SvgIcon icon="solar:devices-bold-duotone" /> Unified device view</span>
        </div>
      </div>

      <div class="story-foot">
        <span class="signal-line"></span>
        Secure relay online
      </div>
    </section>

    <section class="login-panel">
      <div class="login-actions">
        <ThemeSchemaSwitch
          :theme-schema="themeStore.themeScheme"
          :show-tooltip="false"
          @switch="themeStore.toggleThemeScheme"
        />
        <LangSwitch
          :lang="appStore.locale"
          :lang-options="appStore.localeOptions"
          :show-tooltip="false"
          @change-lang="appStore.changeLocale"
        />
      </div>

      <NCard :bordered="false" class="login-card">
        <div class="mobile-brand">
          <span class="brand-mark"><SvgIcon icon="solar:monitor-smartphone-bold-duotone" /></span>
          <strong>West454</strong>
        </div>
        <span class="eyebrow">Welcome home</span>
        <h2>{{ $t(activeModule.label) }}</h2>
        <p class="login-subtitle">Use your private console credentials to continue.</p>
        <Transition :name="themeStore.page.animateMode" mode="out-in" appear>
          <component :is="activeModule.component" />
        </Transition>
        <div class="login-assurance">
          <SvgIcon icon="solar:shield-check-linear" />
          Credentials stay on your server
        </div>
      </NCard>
    </section>
  </div>
</template>

<style scoped>
.login-shell {
  position: relative;
  display: grid;
  min-height: 100%;
  grid-template-columns: minmax(420px, 1.08fr) minmax(440px, 0.92fr);
  overflow: hidden;
  background:
    radial-gradient(circle at 18% 14%, rgba(124, 58, 237, 0.22), transparent 30rem),
    radial-gradient(circle at 90% 80%, rgba(6, 182, 212, 0.12), transparent 28rem),
    #040711;
}

.login-grid {
  position: absolute;
  inset: 0;
  background-image:
    linear-gradient(rgba(148, 163, 184, 0.055) 1px, transparent 1px),
    linear-gradient(90deg, rgba(148, 163, 184, 0.055) 1px, transparent 1px);
  background-size: 52px 52px;
  mask-image: radial-gradient(circle at 35% 45%, black, transparent 68%);
  pointer-events: none;
}

.login-story,
.login-panel {
  position: relative;
  z-index: 1;
}

.login-story {
  display: flex;
  min-height: 100vh;
  flex-direction: column;
  justify-content: space-between;
  border-right: 1px solid rgba(148, 163, 184, 0.1);
  padding: 42px clamp(38px, 6vw, 92px);
}

.login-brand {
  display: flex;
  align-items: center;
  gap: 12px;
}

.brand-mark {
  display: grid;
  width: 42px;
  height: 42px;
  place-items: center;
  border: 1px solid rgba(196, 181, 253, 0.34);
  border-radius: 13px;
  background: linear-gradient(145deg, #8b5cf6, #4f46e5 58%, #0891b2);
  box-shadow: 0 12px 34px rgba(124, 58, 237, 0.38);
  color: white;
  font-size: 24px;
}

.login-brand strong,
.login-brand small {
  display: block;
}

.login-brand strong {
  color: white;
  font-size: 16px;
  font-weight: 760;
}

.login-brand small {
  margin-top: 4px;
  color: #64748b;
  font-size: 9px;
  font-weight: 700;
  letter-spacing: 0.14em;
  text-transform: uppercase;
}

.story-copy {
  max-width: 650px;
}

.story-copy h1 {
  margin: 24px 0 18px;
  color: #f8fafc;
  font-size: clamp(46px, 6vw, 76px);
  font-weight: 780;
  letter-spacing: -0.065em;
  line-height: 0.98;
}

.story-copy h1 span {
  display: block;
  background: linear-gradient(95deg, #c4b5fd, #8b5cf6 48%, #67e8f9);
  background-clip: text;
  color: transparent;
}

.story-copy > p {
  max-width: 570px;
  margin: 0;
  color: #94a3b8;
  font-size: 16px;
  line-height: 1.75;
}

.story-points {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
  margin-top: 28px;
}

.story-points span {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  border: 1px solid rgba(148, 163, 184, 0.13);
  border-radius: 999px;
  background: rgba(15, 23, 42, 0.58);
  color: #cbd5e1;
  font-size: 11px;
  font-weight: 620;
  padding: 9px 12px;
}

.story-points :deep(svg) {
  color: #a78bfa;
  font-size: 16px;
}

.story-foot {
  display: flex;
  align-items: center;
  gap: 10px;
  color: #64748b;
  font-size: 10px;
  font-weight: 700;
  letter-spacing: 0.12em;
  text-transform: uppercase;
}

.signal-line {
  width: 44px;
  height: 2px;
  border-radius: 999px;
  background: linear-gradient(90deg, #8b5cf6, #22d3ee);
  box-shadow: 0 0 16px rgba(34, 211, 238, 0.5);
}

.login-panel {
  display: grid;
  min-height: 100vh;
  place-items: center;
  padding: 80px clamp(24px, 6vw, 90px) 48px;
}

.login-actions {
  position: absolute;
  top: 28px;
  right: 32px;
  display: flex;
  gap: 6px;
}

.login-card {
  width: min(100%, 470px);
  border-color: rgba(139, 92, 246, 0.24) !important;
  background:
    radial-gradient(circle at 100% 0, rgba(139, 92, 246, 0.13), transparent 22rem),
    rgba(11, 15, 28, 0.78) !important;
  box-shadow: 0 40px 110px rgba(0, 0, 0, 0.44) !important;
}

.login-card :deep(.n-card__content) {
  padding: 38px !important;
}

.login-card h2 {
  margin: 18px 0 8px;
  color: var(--text-strong);
  font-size: 31px;
  font-weight: 760;
  letter-spacing: -0.045em;
}

.login-subtitle {
  margin: 0 0 28px;
  color: var(--text-muted);
  line-height: 1.6;
}

.login-assurance {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 7px;
  margin-top: 24px;
  color: var(--text-muted);
  font-size: 10px;
  font-weight: 650;
  letter-spacing: 0.04em;
}

.login-assurance :deep(svg) {
  color: #34d399;
  font-size: 15px;
}

.mobile-brand {
  display: none;
}

@media (max-width: 900px) {
  .login-shell {
    grid-template-columns: 1fr;
  }

  .login-story {
    display: none;
  }

  .login-panel {
    padding: 76px 20px 28px;
  }

  .mobile-brand {
    display: flex;
    align-items: center;
    gap: 10px;
    margin-bottom: 24px;
    color: var(--text-strong);
  }
}

@media (max-width: 520px) {
  .login-card :deep(.n-card__content) {
    padding: 28px 22px !important;
  }
}
</style>
