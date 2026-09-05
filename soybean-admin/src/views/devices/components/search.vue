<script setup lang="ts">
import { $t } from '@/locales';
import { useNaiveForm } from '@/hooks/common/form';

defineOptions({
  name: 'AuditBaseLogsSearch'
});

interface Emits {
  (e: 'reset'): void;
  (e: 'search'): void;
}

const emit = defineEmits<Emits>();

const { formRef } = useNaiveForm();
const model = defineModel<Record<string, any>>('model', { required: true });

async function reset() {
  emit('reset');
}

async function search() {
  emit('search');
}
</script>

<template>
  <section class="search-panel spotlight-card">
    <div class="search-intro">
      <span class="search-icon"><SvgIcon icon="solar:magnifer-linear" /></span>
      <div>
        <strong>Find a computer</strong>
        <span>Sort, search and filter your private fleet</span>
      </div>
    </div>
    <NForm ref="formRef" :model="model" label-placement="top">
      <NGrid :x-gap="12" :y-gap="10" responsive="screen" item-responsive>
        <NFormItemGi span="24 s:12 l:6" label="Quick search" path="keyword">
          <NInput v-model:value="model.keyword" clearable placeholder="Search hostname, operator, ID, OS, last user" />
        </NFormItemGi>
        <NFormItemGi span="24 s:12 l:6" label="Status" path="online">
          <NSelect
            v-model:value="model.online"
            :options="[
              { label: 'Online', value: 'online' },
              { label: 'Offline', value: 'offline' }
            ]"
            clearable
            placeholder="Any"
          />
        </NFormItemGi>
        <NFormItemGi span="24 s:12 l:6" label="Operating system" path="os">
          <NInput v-model:value="model.os" clearable placeholder="Windows 11 Pro, Ubuntu 24.04, macOS 15" />
        </NFormItemGi>
        <NFormItemGi span="24 s:12 l:6" label="RustDesk version" path="version">
          <NInput v-model:value="model.version" clearable placeholder="1.4.6" />
        </NFormItemGi>
        <NFormItemGi span="24 s:12 l:6" label="Operator" path="username">
          <NInput v-model:value="model.username" clearable placeholder="Current operator" />
        </NFormItemGi>
        <NFormItemGi span="24 s:12 l:6" label="Last signed-in user" path="last_user">
          <NInput v-model:value="model.last_user" clearable placeholder="Most recent user" />
        </NFormItemGi>
        <NFormItemGi span="24 s:12 l:6" label="Computer name" path="hostname">
          <NInput v-model:value="model.hostname" clearable placeholder="fdl-badrib" />
        </NFormItemGi>
        <NFormItemGi span="24 s:12 l:6" label="RustDesk ID" path="rustdesk_id">
          <NInput v-model:value="model.rustdesk_id" clearable placeholder="000 000 000" />
        </NFormItemGi>
        <NFormItemGi span="24 s:12 l:6" class="action-field">
          <NSpace class="w-full" justify="end">
            <NButton @click="reset">
              <template #icon><icon-ic-round-refresh class="text-icon" /></template>
              Reset
            </NButton>
            <NButton type="primary" @click="search">
              <template #icon><icon-ic-round-search class="text-icon" /></template>
              Search
            </NButton>
          </NSpace>
        </NFormItemGi>
      </NGrid>
    </NForm>
  </section>
</template>

<style scoped>
.search-panel {
  display: grid;
  grid-template-columns: minmax(190px, 0.35fr) minmax(0, 1.65fr);
  align-items: end;
  gap: 22px;
  padding: 18px 20px 8px;
}

.search-intro {
  display: flex;
  align-items: center;
  gap: 11px;
  padding-bottom: 16px;
}

.search-icon {
  display: grid;
  width: 38px;
  height: 38px;
  flex: 0 0 38px;
  place-items: center;
  border: 1px solid rgba(139, 92, 246, 0.24);
  border-radius: 11px;
  background: rgba(139, 92, 246, 0.12);
  color: var(--accent-bright);
  font-size: 20px;
}

.search-intro strong,
.search-intro span {
  display: block;
}

.search-intro strong {
  color: var(--text-strong);
  font-size: 13px;
  font-weight: 700;
}

.search-intro span {
  margin-top: 4px;
  color: var(--text-muted);
  font-size: 10px;
}

.action-field {
  align-items: end;
}

:deep(.n-form-item-label) {
  color: var(--text-muted);
  font-size: 10px;
  font-weight: 700;
  letter-spacing: 0.08em;
  text-transform: uppercase;
}

@media (max-width: 1024px) {
  .search-panel {
    grid-template-columns: 1fr;
  }

  .search-intro {
    padding-bottom: 0;
  }
}
</style>
