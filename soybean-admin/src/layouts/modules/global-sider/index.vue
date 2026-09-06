<script setup lang="ts">
import { computed } from 'vue';
import { useAppStore } from '@/store/modules/app';
import { useThemeStore } from '@/store/modules/theme';
import { GLOBAL_SIDER_MENU_ID } from '@/constants/app';
import GlobalLogo from '../global-logo/index.vue';

defineOptions({
  name: 'GlobalSider'
});

const appStore = useAppStore();
const themeStore = useThemeStore();

const isVerticalMix = computed(() => themeStore.layout.mode === 'vertical-mix');
const isHorizontalMix = computed(() => themeStore.layout.mode === 'horizontal-mix');
const darkMenu = computed(() => !themeStore.darkMode && !isHorizontalMix.value && themeStore.sider.inverted);
const showLogo = computed(() => !isVerticalMix.value && !isHorizontalMix.value);
const menuWrapperClass = computed(() => (showLogo.value ? 'flex-1-hidden' : 'h-full'));
</script>

<template>
  <DarkModeContainer class="app-sider size-full flex-col-stretch shadow-sider" :inverted="darkMenu">
    <GlobalLogo
      v-if="showLogo"
      :show-title="!appStore.siderCollapse"
      :style="{ height: themeStore.header.height + 'px' }"
    />
    <div :id="GLOBAL_SIDER_MENU_ID" :class="menuWrapperClass"></div>
    <div v-if="!appStore.siderCollapse" class="sider-signal">
      <div class="signal-icon">
        <SvgIcon icon="solar:shield-check-bold-duotone" />
      </div>
      <div>
        <strong>Private relay</strong>
        <span>Identity preserved</span>
      </div>
    </div>
  </DarkModeContainer>
</template>

<style scoped>
.sider-signal {
  display: flex;
  align-items: center;
  gap: 11px;
  margin: 10px;
  border: 2px solid var(--surface-border);
  border-radius: var(--radius-sm);
  background: var(--surface-muted);
  box-shadow: var(--shadow-hard-sm);
  padding: 12px;
}

.signal-icon {
  display: grid;
  width: 34px;
  height: 34px;
  flex: 0 0 34px;
  place-items: center;
  border: 2px solid var(--surface-border);
  border-radius: var(--radius-sm);
  background: var(--accent-bright);
  color: white;
  font-size: 20px;
}

.sider-signal strong,
.sider-signal span {
  display: block;
}

.sider-signal strong {
  color: var(--text-strong);
  font-size: 12px;
  line-height: 1.3;
}

.sider-signal span {
  margin-top: 2px;
  color: var(--text-muted);
  font-size: 10px;
}
</style>
