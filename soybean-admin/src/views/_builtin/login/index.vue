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
        <span class="eyebrow login-card-eyebrow">Welcome home</span>
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
  background: #0a0a0a;
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
  border-right: 3px solid #f2f1ec;
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
  border: 2px solid #f2f1ec;
  border-radius: var(--radius-sm);
  background: var(--accent);
  box-shadow: 3px 3px 0 0 #f2f1ec;
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
  font-weight: 800;
}

.login-brand small {
  margin-top: 4px;
  color: #a1a1aa;
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
  color: #f5f5f0;
  font-size: clamp(46px, 6vw, 76px);
  font-weight: 800;
  letter-spacing: -0.055em;
  line-height: 0.98;
}

.story-copy h1 span {
  display: block;
  color: #c4b5fd;
}

.story-copy > p {
  max-width: 570px;
  margin: 0;
  color: #a1a1aa;
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
  border: 2px solid #f2f1ec;
  border-radius: 999px;
  background: #17171b;
  color: #e4e4e7;
  font-size: 11px;
  font-weight: 700;
  padding: 9px 12px;
}

.story-points :deep(svg) {
  color: #c4b5fd;
  font-size: 16px;
}

.story-foot {
  display: flex;
  align-items: center;
  gap: 10px;
  color: #a1a1aa;
  font-size: 10px;
  font-weight: 700;
  letter-spacing: 0.12em;
  text-transform: uppercase;
}

.signal-line {
  width: 44px;
  height: 3px;
  border-radius: 999px;
  background: #22c55e;
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
  border: 3px solid #f2f1ec !important;
  background: #101013 !important;
  box-shadow: 6px 6px 0 0 #f2f1ec !important;
  --surface-border: #f2f1ec;
  --surface: #101013;
  --text-strong: #f5f5f0;
  --text-muted: #a1a1aa;
}

.login-card :deep(.n-card__content) {
  padding: 38px !important;
}

.login-card-eyebrow {
  background: var(--accent) !important;
  border-color: #f2f1ec !important;
}

.login-card h2 {
  margin: 18px 0 8px;
  color: #f5f5f0;
  font-size: 31px;
  font-weight: 800;
  letter-spacing: -0.035em;
}

.login-subtitle {
  margin: 0 0 28px;
  color: #a1a1aa;
  line-height: 1.6;
}

.login-assurance {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 7px;
  margin-top: 24px;
  color: #a1a1aa;
  font-size: 10px;
  font-weight: 700;
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
