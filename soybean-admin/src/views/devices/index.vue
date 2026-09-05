<script setup lang="tsx">
import { fetchDevicesList } from '@/service/api/devices';
import { $t } from '@/locales';
import { useAppStore } from '@/store/modules/app';
import { useTable } from '@/hooks/common/table';
import TableHeader from './components/table-header.vue';
import AuditBaseLogsSearch from './components/search.vue';

const appStore = useAppStore();

const {
  columns,
  columnChecks,
  data,
  getData,
  getDataByPage,
  loading,
  mobilePagination,
  searchParams,
  resetSearchParams
} = useTable({
  apiFn: fetchDevicesList,
  showTotal: true,
  apiParams: {
    current: 1,
    size: 10,
    // if you want to use the searchParams in Form, you need to define the following properties, and the value is null
    // the value can not be undefined, otherwise the property in Form will not be reactive
    hostname: null,
    username: null,
    rustdesk_id: null
  },
  columns: () => [
    {
      key: 'id',
      title: 'ID',
      align: 'center'
    },
    {
      key: 'rustdesk_id',
      title: $t('dataMap.device.rustdesk_id'),
      align: 'left',
      width: 150,
      render: row => (
        <span class="device-id-pill">
          <span class="device-id-dot"></span>
          {row.rustdesk_id}
        </span>
      )
    },
    {
      key: 'hostname',
      title: $t('dataMap.device.hostname'),
      align: 'left',
      width: 190,
      ellipsis: { tooltip: true },
      render: row => (
        <div class="device-name-cell">
          <span class="device-avatar">{(row.hostname || '?').slice(0, 1).toUpperCase()}</span>
          <strong>{row.hostname || 'Unnamed device'}</strong>
        </div>
      )
    },
    {
      key: 'username',
      title: $t('dataMap.device.username'),
      align: 'center'
    },
    {
      key: 'version',
      title: $t('dataMap.device.version'),
      align: 'center',
      width: 110,
      render: row => <span class="version-pill">v{row.version || '—'}</span>
    },
    {
      key: 'os',
      title: $t('dataMap.device.os'),
      align: 'center'
    },
    {
      key: 'memory',
      title: $t('dataMap.device.memory'),
      align: 'center'
    },
    {
      key: 'created_at',
      title: $t('dataMap.audit.created_at'),
      align: 'center'
    }
  ]
});
</script>

<template>
  <div class="aceternity-page devices-page">
    <header class="fleet-heading">
      <div>
        <span class="eyebrow">
          <span class="status-dot"></span>
          Private fleet
        </span>
        <h1>Every computer. One calm view.</h1>
        <p>Search, inspect and monitor the devices connected to your private RustDesk network.</p>
      </div>
      <div class="fleet-summary">
        <span class="summary-icon"><SvgIcon icon="solar:monitor-smartphone-bold-duotone" /></span>
        <div>
          <strong>{{ data.length }}</strong>
          <small>loaded on this page</small>
        </div>
      </div>
    </header>

    <AuditBaseLogsSearch v-model:model="searchParams" @reset="resetSearchParams" @search="getDataByPage" />

    <section class="device-table-card spotlight-card">
      <div class="table-toolbar">
        <div>
          <span class="table-kicker">Device inventory</span>
          <h2>{{ $t('route.devices') }}</h2>
        </div>
        <TableHeader v-model:columns="columnChecks" :loading="loading" @refresh="getData" />
      </div>
      <NDataTable
        :columns="columns"
        :data="data"
        size="small"
        :flex-height="!appStore.isMobile"
        :scroll-x="1060"
        :loading="loading"
        remote
        :row-key="row => row.id"
        :pagination="mobilePagination"
        class="device-table sm:h-full"
      />
    </section>
  </div>
</template>

<style scoped>
.devices-page {
  display: flex;
  min-height: 560px;
  flex-direction: column;
  gap: 18px;
}

.fleet-heading {
  display: flex;
  align-items: end;
  justify-content: space-between;
  gap: 24px;
  padding: 4px 4px 10px;
}

.fleet-heading h1 {
  margin: 16px 0 8px;
  color: var(--text-strong);
  font-size: clamp(28px, 3vw, 40px);
  font-weight: 770;
  letter-spacing: -0.045em;
  line-height: 1.05;
}

.fleet-heading p {
  margin: 0;
  color: var(--text-muted);
  font-size: 14px;
}

.fleet-summary {
  display: flex;
  min-width: 174px;
  align-items: center;
  gap: 12px;
  border: 1px solid var(--surface-border);
  border-radius: 15px;
  background: var(--surface);
  padding: 12px 14px;
}

.summary-icon {
  display: grid;
  width: 38px;
  height: 38px;
  place-items: center;
  border-radius: 11px;
  background: rgba(139, 92, 246, 0.14);
  color: var(--accent-bright);
  font-size: 21px;
}

.fleet-summary strong,
.fleet-summary small {
  display: block;
}

.fleet-summary strong {
  color: var(--text-strong);
  font-size: 19px;
  line-height: 1;
}

.fleet-summary small {
  margin-top: 4px;
  color: var(--text-muted);
  font-size: 10px;
  font-weight: 650;
  letter-spacing: 0.05em;
  text-transform: uppercase;
}

.device-table-card {
  display: flex;
  min-height: 430px;
  flex: 1;
  flex-direction: column;
  padding: 20px;
}

.table-toolbar {
  position: relative;
  z-index: 1;
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 18px;
  padding: 0 2px 16px;
}

.table-kicker {
  color: var(--accent-bright);
  font-size: 10px;
  font-weight: 750;
  letter-spacing: 0.12em;
  text-transform: uppercase;
}

.table-toolbar h2 {
  margin: 4px 0 0;
  color: var(--text-strong);
  font-size: 19px;
  font-weight: 720;
  letter-spacing: -0.025em;
}

.device-table {
  position: relative;
  z-index: 1;
  flex: 1;
}

:deep(.device-id-pill) {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  border: 1px solid var(--surface-border);
  border-radius: 999px;
  background: var(--surface-muted);
  color: var(--text-strong);
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 11px;
  padding: 5px 8px;
}

:deep(.device-id-dot) {
  width: 6px;
  height: 6px;
  border-radius: 50%;
  background: #34d399;
  box-shadow: 0 0 10px rgba(52, 211, 153, 0.75);
}

:deep(.device-name-cell) {
  display: flex;
  align-items: center;
  gap: 10px;
}

:deep(.device-name-cell strong) {
  overflow: hidden;
  color: var(--text-strong);
  font-size: 12px;
  font-weight: 650;
  text-overflow: ellipsis;
  white-space: nowrap;
}

:deep(.device-avatar) {
  display: grid;
  width: 29px;
  height: 29px;
  flex: 0 0 29px;
  place-items: center;
  border: 1px solid rgba(139, 92, 246, 0.24);
  border-radius: 9px;
  background: linear-gradient(145deg, rgba(139, 92, 246, 0.2), rgba(34, 211, 238, 0.08));
  color: var(--accent-bright);
  font-size: 11px;
  font-weight: 750;
}

:deep(.version-pill) {
  border: 1px solid rgba(34, 211, 238, 0.2);
  border-radius: 999px;
  background: rgba(34, 211, 238, 0.07);
  color: var(--accent-cyan);
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 10px;
  padding: 4px 7px;
}

@media (max-width: 700px) {
  .fleet-heading {
    align-items: flex-start;
    flex-direction: column;
  }

  .fleet-summary {
    width: 100%;
  }

  .device-table-card {
    min-height: 520px;
    padding: 14px;
  }

  .table-toolbar {
    align-items: flex-start;
    flex-direction: column;
  }
}
</style>
