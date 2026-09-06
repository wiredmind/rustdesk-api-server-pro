<script setup lang="ts">

defineOptions({
  name: 'GlobalLogo'
});

interface Props {
  /** Whether to show the title */
  showTitle?: boolean;
}

withDefaults(defineProps<Props>(), {
  showTitle: true
});
</script>

<template>
  <RouterLink to="/" class="brand-lockup w-full nowrap-hidden">
    <span class="brand-mark">
      <SystemLogo class="text-24px" />
    </span>
    <span v-show="showTitle" class="brand-copy">
      <strong>West454</strong>
      <small>Remote operations</small>
    </span>
  </RouterLink>
</template>

<style scoped>
.brand-lockup {
  display: flex;
  align-items: center;
  gap: 11px;
  padding: 0 14px;
}

.brand-mark {
  display: grid;
  width: 38px;
  height: 38px;
  flex: 0 0 38px;
  place-items: center;
  border: 2px solid var(--surface-border);
  border-radius: var(--radius-sm);
  background: var(--accent);
  box-shadow: var(--shadow-hard-sm);
  color: white;
}

.brand-copy {
  min-width: 0;
  line-height: 1;
}

.brand-copy strong,
.brand-copy small {
  display: block;
  white-space: nowrap;
}

.brand-copy strong {
  color: var(--text-strong);
  font-size: 15px;
  font-weight: 800;
  letter-spacing: -0.02em;
}

.brand-copy small {
  margin-top: 6px;
  color: var(--text-muted);
  font-size: 8px;
  font-weight: 700;
  letter-spacing: 0.14em;
  text-transform: uppercase;
}
</style>

