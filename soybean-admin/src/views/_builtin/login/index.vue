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
      <div class="login-mark">
        <span class="brand-mark"><SvgIcon icon="solar:monitor-smartphone-bold-duotone" /></span>
        <strong>West454</strong>
      </div>
      <h2>{{ $t(activeModule.label) }}</h2>
      <Transition :name="themeStore.page.animateMode" mode="out-in" appear>
        <component :is="activeModule.component" />
      </Transition>
    </NCard>
  </div>
</template>

<style scoped>
.login-shell {
  position: relative;
  display: grid;
  min-height: 100%;
  place-items: center;
  background-color: #0a0a0a !important;
}

.login-actions {
  position: absolute;
  top: 28px;
  right: 32px;
  display: flex;
  gap: 6px;
}

.login-card {
  width: min(100%, 420px);
  margin: 24px;
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

.login-mark {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 22px;
}

.brand-mark {
  display: grid;
  width: 34px;
  height: 34px;
  flex: 0 0 34px;
  place-items: center;
  border: 2px solid #f2f1ec;
  border-radius: var(--radius-sm);
  background: var(--accent-fill);
  color: white;
  font-size: 19px;
}

.login-mark strong {
  color: #f5f5f0;
  font-size: 14px;
  font-weight: 700;
}

.login-card h2 {
  margin: 0 0 22px;
  color: #f5f5f0;
  font-size: 22px;
  font-weight: 800;
  letter-spacing: -0.02em;
}

@media (max-width: 520px) {
  .login-card :deep(.n-card__content) {
    padding: 28px 22px !important;
  }
}
</style>
