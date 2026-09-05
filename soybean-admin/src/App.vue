<script setup lang="ts">
import { computed, onBeforeUnmount, onMounted } from 'vue';
import { NConfigProvider, darkTheme } from 'naive-ui';
import type { WatermarkProps } from 'naive-ui';
import { useAppStore } from './store/modules/app';
import { useThemeStore } from './store/modules/theme';
import { naiveDateLocales, naiveLocales } from './locales/naive';

defineOptions({
  name: 'App'
});

const appStore = useAppStore();
const themeStore = useThemeStore();

const naiveDarkTheme = computed(() => (themeStore.darkMode ? darkTheme : undefined));

const naiveLocale = computed(() => {
  return naiveLocales[appStore.locale];
});

const naiveDateLocale = computed(() => {
  return naiveDateLocales[appStore.locale];
});

const watermarkProps = computed<WatermarkProps>(() => {
  return {
    content: themeStore.watermark?.text || 'SoybeanAdmin',
    cross: true,
    fullscreen: true,
    fontSize: 16,
    lineHeight: 16,
    width: 384,
    height: 384,
    xOffset: 12,
    yOffset: 60,
    rotate: -15,
    zIndex: 9999
  };
});
let pointerFrame: number | undefined;

function updateAmbientPosition(event: PointerEvent) {
  if (pointerFrame !== undefined) return;

  pointerFrame = window.requestAnimationFrame(() => {
    document.documentElement.style.setProperty('--ambient-x', `${event.clientX}px`);
    document.documentElement.style.setProperty('--ambient-y', `${event.clientY}px`);
    pointerFrame = undefined;
  });
}

onMounted(() => {
  window.addEventListener('pointermove', updateAmbientPosition, { passive: true });
});

onBeforeUnmount(() => {
  window.removeEventListener('pointermove', updateAmbientPosition);
  if (pointerFrame !== undefined) window.cancelAnimationFrame(pointerFrame);
});
</script>

<template>
  <NConfigProvider
    :theme="naiveDarkTheme"
    :theme-overrides="themeStore.naiveTheme"
    :locale="naiveLocale"
    :date-locale="naiveDateLocale"
    class="h-full"
  >
    <AppProvider>
      <RouterView class="bg-layout" />
      <NWatermark v-if="themeStore.watermark?.visible" v-bind="watermarkProps" />
    </AppProvider>
  </NConfigProvider>
</template>

<style scoped></style>
