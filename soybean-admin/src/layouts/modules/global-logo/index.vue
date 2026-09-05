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
  border: 1px solid rgba(196, 181, 253, 0.3);
  border-radius: 12px;
  background: linear-gradient(145deg, #8b5cf6, #4f46e5 58%, #0891b2);
  box-shadow: 0 10px 28px rgba(124, 58, 237, 0.34), inset 0 1px 0 rgba(255, 255, 255, 0.35);
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
  background: linear-gradient(90deg, var(--text-strong), var(--accent-bright));
  background-clip: text;
  color: transparent;
  font-size: 15px;
  font-weight: 760;
  letter-spacing: -0.025em;
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

