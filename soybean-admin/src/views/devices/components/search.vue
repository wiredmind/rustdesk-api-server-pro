<script setup lang="ts">
import { ref } from 'vue';

defineOptions({
  name: 'DevicesSearch'
});

interface Emits {
  (e: 'reset'): void;
  (e: 'search'): void;
}

const emit = defineEmits<Emits>();

const model = defineModel<Record<string, any>>('model', { required: true });

const advancedOpen = ref(false);

function reset() {
  emit('reset');
}

function search() {
  emit('search');
}
</script>

<template>
  <section class="search-bar spotlight-card">
    <NInput
      v-model:value="model.keyword"
      class="search-keyword"
      clearable
      placeholder="Search hostname, operator, RustDesk ID, OS, last user"
      @keyup.enter="search"
    >
      <template #prefix>
        <SvgIcon icon="solar:magnifer-linear" class="text-icon" />
      </template>
    </NInput>
    <NSelect
      v-model:value="model.online"
      class="search-status"
      clearable
      placeholder="Status"
      :options="[
        { label: 'Online', value: 'online' },
        { label: 'Offline', value: 'offline' }
      ]"
    />
    <NInput v-model:value="model.os" class="search-os" clearable placeholder="OS" @keyup.enter="search" />
    <NInput v-model:value="model.version" class="search-version" clearable placeholder="Version" @keyup.enter="search" />
    <NButton size="medium" @click="search">
      <template #icon><icon-ic-round-search class="text-icon" /></template>
      Search
    </NButton>
    <NButton size="medium" quaternary @click="reset">
      <template #icon><icon-ic-round-refresh class="text-icon" /></template>
      Reset
    </NButton>
    <NPopover trigger="click" placement="bottom-end" :show-arrow="false" raw>
      <template #trigger>
        <button type="button" class="advanced-toggle" :class="{ 'is-active': advancedOpen }" @click="advancedOpen = !advancedOpen">
          <SvgIcon icon="solar:tuning-2-linear" />
          More filters
        </button>
      </template>
      <div class="advanced-panel">
        <NFormItem label="Computer name" label-placement="top">
          <NInput v-model:value="model.hostname" clearable placeholder="fdl-badrib" @keyup.enter="search" />
        </NFormItem>
        <NFormItem label="RustDesk ID" label-placement="top">
          <NInput v-model:value="model.rustdesk_id" clearable placeholder="000 000 000" @keyup.enter="search" />
        </NFormItem>
        <NFormItem label="Current operator" label-placement="top">
          <NInput v-model:value="model.username" clearable placeholder="Current operator" @keyup.enter="search" />
        </NFormItem>
        <NFormItem label="Last signed-in user" label-placement="top">
          <NInput v-model:value="model.last_user" clearable placeholder="Most recent user" @keyup.enter="search" />
        </NFormItem>
        <NButton type="primary" block @click="search">Apply filters</NButton>
      </div>
    </NPopover>
  </section>
</template>

<style scoped>
.search-bar {
  display: flex;
  align-items: center;
  flex-wrap: wrap;
  gap: 10px;
  padding: 12px 16px;
}

.search-keyword {
  min-width: 260px;
  flex: 2 1 260px;
}

.search-status {
  min-width: 130px;
  flex: 0 1 130px;
}

.search-os {
  min-width: 140px;
  flex: 0 1 140px;
}

.search-version {
  min-width: 110px;
  flex: 0 1 110px;
}

.advanced-toggle {
  display: inline-flex;
  align-items: center;
  gap: 7px;
  border: 2px solid var(--surface-border);
  border-radius: var(--radius-sm);
  background: var(--surface-muted);
  color: var(--text-strong);
  font-size: 12px;
  font-weight: 700;
  white-space: nowrap;
  padding: 8px 12px;
  box-shadow: var(--shadow-hard-sm);
  transition:
    transform 140ms ease,
    box-shadow 140ms ease;
}

.advanced-toggle:hover,
.advanced-toggle.is-active {
  background: var(--accent-fill);
  color: #ffffff;
  transform: translate(-1px, -1px);
  box-shadow: 3px 3px 0 0 var(--shadow-color);
}

.advanced-panel {
  display: flex;
  flex-direction: column;
  gap: 12px;
  width: 280px;
  background: var(--surface-strong);
  border: 3px solid var(--surface-border);
  border-radius: var(--radius);
  box-shadow: var(--shadow-hard-lg);
  padding: 18px;
}

.advanced-panel :deep(.n-form-item) {
  margin-bottom: 0;
}

.advanced-panel :deep(.n-form-item-label) {
  color: var(--text-muted);
  font-size: 10px;
  font-weight: 700;
  letter-spacing: 0.06em;
  text-transform: uppercase;
}

@media (max-width: 720px) {
  .search-bar {
    align-items: stretch;
    flex-direction: column;
  }

  .search-keyword,
  .search-status,
  .search-os,
  .search-version {
    flex: 1 1 auto;
    min-width: 0;
    width: 100%;
  }
}
</style>
