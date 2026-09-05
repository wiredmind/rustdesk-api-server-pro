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

const model = defineModel<Api.Devices.DeviceSearchParams>('model', { required: true });

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
        <span>Filter by name, operator or RustDesk ID</span>
      </div>
    </div>
    <NForm ref="formRef" :model="model" label-placement="top">
      <NGrid :x-gap="12" :y-gap="10" responsive="screen" item-responsive>
        <NFormItemGi span="24 s:12 l:6" :label="$t('dataMap.device.hostname')" path="hostname">
          <NInput v-model:value="model.hostname" clearable placeholder="Computer name" />
        </NFormItemGi>
        <NFormItemGi span="24 s:12 l:6" :label="$t('dataMap.device.username')" path="username">
          <NInput v-model:value="model.username" clearable placeholder="Operator" />
        </NFormItemGi>
        <NFormItemGi span="24 s:12 l:6" :label="$t('dataMap.device.rustdesk_id')" path="rustdesk_id">
          <NInput v-model:value="model.rustdesk_id" clearable placeholder="000 000 000" />
        </NFormItemGi>
        <NFormItemGi span="24 s:12 l:6" class="action-field">
          <NSpace class="w-full" justify="end">
            <NButton @click="reset">
              <template #icon>
                <icon-ic-round-refresh class="text-icon" />
              </template>
              {{ $t('common.reset') }}
            </NButton>
            <NButton type="primary" @click="search">
              <template #icon>
                <icon-ic-round-search class="text-icon" />
              </template>
              {{ $t('common.search') }}
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
