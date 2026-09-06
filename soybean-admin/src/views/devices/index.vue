<script setup lang="ts">
import { computed, h, ref, watch } from 'vue';
import { NTag, NTooltip, useMessage } from 'naive-ui';
import { deleteDevices, fetchDevicesList } from '@/service/api/devices';
import { $t } from '@/locales';
import { useAppStore } from '@/store/modules/app';
import { useTable, useTableOperate } from '@/hooks/common/table';
import TableHeader from './components/table-header.vue';
import DevicesSearch from './components/search.vue';
import ConnectMenu from './components/connect-menu.vue';

type DeviceRow = Api.Devices.Device & { index: number };

type SortOrder = 'ascend' | 'descend' | false;

interface SortState {
  columnKey: string | number | null;
  order: SortOrder;
}

const appStore = useAppStore();
const message = useMessage();

const sort = ref<SortState>({ columnKey: 'is_online', order: 'descend' });

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
    size: 15,
    hostname: null,
    username: null,
    rustdesk_id: null,
    last_user: null,
    keyword: null,
    version: null,
    os: null,
    online: null,
    sort_by: 'is_online',
    sort_order: 'desc'
  },
  columns: () => [
    {
      type: 'selection',
      width: 48,
      align: 'center',
      fixed: 'left',
      selectable: (row: DeviceRow) => Boolean(row.rustdesk_id)
    },
    {
      key: 'index',
      title: $t('common.index'),
      width: 62,
      align: 'center',
      render: (row: DeviceRow) =>
        h('span', { class: 'cell-index' }, String(row.index).padStart(2, '0'))
    },
    {
      key: 'hostname',
      title: $t('dataMap.device.hostname'),
      width: 260,
      minWidth: 180,
      maxWidth: 480,
      resizable: true,
      sorter: true,
      sortOrder: sort.value.columnKey === 'hostname' ? sort.value.order : false,
      render: (row: DeviceRow) =>
        h('div', { class: 'cell-host' }, [
          h('span', { class: 'cell-host-avatar' }, (row.hostname || row.rustdesk_id || '?').slice(0, 1).toUpperCase()),
          h('div', { class: 'cell-host-meta' }, [
            h('strong', { class: 'cell-host-name', title: row.hostname || '' }, row.hostname || 'Unnamed computer'),
            h('span', { class: 'cell-host-sub' }, [row.os, row.cpu].filter(Boolean).join(' • ') || 'System info unavailable')
          ])
        ])
    },
    {
      key: 'rustdesk_id',
      title: $t('dataMap.device.rustdesk_id'),
      width: 180,
      minWidth: 140,
      maxWidth: 260,
      resizable: true,
      sorter: true,
      sortOrder: sort.value.columnKey === 'rustdesk_id' ? sort.value.order : false,
      render: (row: DeviceRow) =>
        h('span', { class: 'cell-id' }, [
          h('span', { class: ['cell-id-dot', row.is_online ? 'is-online' : 'is-offline'] }),
          h('code', { class: 'cell-id-text' }, row.rustdesk_id || '—')
        ])
    },
    {
      key: 'username',
      title: $t('page.devices.table.columns.user' as App.I18n.I18nKey),
      width: 150,
      minWidth: 110,
      maxWidth: 240,
      resizable: true,
      sorter: true,
      sortOrder: sort.value.columnKey === 'username' ? sort.value.order : false,
      render: (row: DeviceRow) =>
        h('span', { class: 'cell-name' }, row.username || '—')
    },
    {
      key: 'last_user',
      title: $t('page.devices.table.columns.lastUser' as App.I18n.I18nKey),
      width: 160,
      minWidth: 110,
      maxWidth: 240,
      resizable: true,
      sorter: true,
      sortOrder: sort.value.columnKey === 'last_user' ? sort.value.order : false,
      render: (row: DeviceRow) =>
        h('span', { class: 'cell-name muted' }, row.last_user || row.username || '—')
    },
    {
      key: 'last_online',
      title: $t('page.devices.table.columns.lastOnline' as App.I18n.I18nKey),
      width: 190,
      minWidth: 150,
      maxWidth: 280,
      resizable: true,
      sorter: true,
      sortOrder: sort.value.columnKey === 'last_online' ? sort.value.order : false,
      render: (row: DeviceRow) =>
        h('div', { class: 'cell-time' }, [
          h(
            NTag,
            {
              bordered: false,
              round: true,
              size: 'small',
              type: row.is_online ? 'success' : 'default',
              class: 'cell-time-tag'
            },
            { default: () => (row.is_online ? $t('page.devices.table.status.online' as App.I18n.I18nKey) : $t('page.devices.table.status.offline' as App.I18n.I18nKey)) }
          ),
          h('span', { class: 'cell-time-text' }, row.last_online || 'Awaiting first heartbeat')
        ])
    },
    {
      key: 'version',
      title: $t('dataMap.device.version'),
      width: 150,
      minWidth: 110,
      maxWidth: 220,
      resizable: true,
      sorter: true,
      sortOrder: sort.value.columnKey === 'version' ? sort.value.order : false,
      render: (row: DeviceRow) =>
        h('div', { class: 'cell-version' }, [
          h('span', { class: 'cell-version-pill' }, row.version ? `v${row.version}` : '—'),
          h('span', { class: 'cell-version-meta' }, `${row.memory || '—'} • ${row.conns || 0} live`)
        ])
    },
    {
      key: 'os',
      title: $t('dataMap.device.os'),
      width: 150,
      minWidth: 100,
      maxWidth: 260,
      resizable: true,
      sorter: true,
      sortOrder: sort.value.columnKey === 'os' ? sort.value.order : false,
      render: (row: DeviceRow) =>
        h('span', { class: 'cell-os' }, row.os || '—')
    },
    {
      key: 'is_online',
      title: $t('page.devices.table.columns.status' as App.I18n.I18nKey),
      minWidth: 120,
      sorter: true,
      sortOrder: sort.value.columnKey === 'is_online' ? sort.value.order : false,
      render: (row: DeviceRow) => {
        const online = row.is_online;
        return h(
          'div',
          { class: ['cell-status', online ? 'is-online' : 'is-offline'] },
          [
            h('span', { class: 'cell-status-dot' }),
            h('span', { class: 'cell-status-label' }, online ? $t('page.devices.table.status.online' as App.I18n.I18nKey) : $t('page.devices.table.status.offline' as App.I18n.I18nKey))
          ]
        );
      }
    },
    {
      key: 'created_at',
      title: $t('dataMap.audit.created_at'),
      width: 170,
      minWidth: 140,
      maxWidth: 240,
      resizable: true,
      sorter: true,
      sortOrder: sort.value.columnKey === 'created_at' ? sort.value.order : false,
      render: (row: DeviceRow) =>
        h('span', { class: 'cell-time-text muted' }, row.created_at || '—')
    },
    {
      key: 'operate',
      title: $t('common.operate'),
      width: 170,
      fixed: 'right',
      align: 'center',
      render: (row: DeviceRow) =>
        h(
          NTooltip,
          { trigger: 'hover', placement: 'top' },
          {
            trigger: () =>
              h(ConnectMenu, {
                deviceId: row.id,
                rustdeskId: row.rustdesk_id,
                hostname: row.hostname,
                onDeleted: () => void getData()
              }),
            default: () => $t('page.devices.table.actions.connect' as App.I18n.I18nKey)
          }
        )
    }
  ]
});

const sortableKeys = new Set([
  'is_online', 'last_online', 'hostname', 'rustdesk_id',
  'username', 'last_user', 'version', 'os', 'memory', 'created_at'
]);

const { checkedRowKeys, onBatchDeleted } = useTableOperate(data, getData);

const selectedRustdeskIds = computed(() =>
  checkedRowKeys.value
    .map(key => data.value.find(item => String(item.id) === String(key))?.rustdesk_id)
    .filter((id): id is string => Boolean(id))
);

async function handleBatchDelete() {
  const res = await deleteDevices({ ids: checkedRowKeys.value.map(key => Number(key)) });
  if (res.error === null) {
    onBatchDeleted();
  }
}

watch(sort, () => {
  const field = String(sort.value.columnKey ?? '');
  searchParams.sort_by = sortableKeys.has(field) ? field : 'is_online';
  searchParams.sort_order = sort.value.order === 'ascend' ? 'asc' : 'desc';
  // Mutate the existing column objects in place instead of recreating the
  // columns array: naive-ui treats a resizable column's width state as
  // belonging to that exact object reference, so replacing the array would
  // silently reset any user-resized widths every time the sort changes.
  for (const col of columns.value) {
    if ('key' in col && sortableKeys.has(String(col.key))) {
      (col as Record<string, unknown>).sortOrder = col.key === sort.value.columnKey ? sort.value.order : false;
    }
  }
  void getData();
});

function onSorterChange(state: SortState | null | undefined) {
  if (!state || !state.columnKey) {
    sort.value = { columnKey: 'is_online', order: 'descend' };
    return;
  }
  const field = String(state.columnKey);
  if (!sortableKeys.has(field)) {
    return;
  }
  sort.value = { columnKey: field, order: state.order || 'descend' };
}

function copySelectedIds() {
  if (selectedRustdeskIds.value.length === 0) {
    message.warning($t('page.devices.table.bulk.empty' as App.I18n.I18nKey));
    return;
  }
  if (navigator?.clipboard?.writeText) {
    void navigator.clipboard.writeText(selectedRustdeskIds.value.join('\n')).then(() => {
      message.success($t('page.devices.table.bulk.copied' as App.I18n.I18nKey, { count: selectedRustdeskIds.value.length }));
    });
    return;
  }
  message.info(selectedRustdeskIds.value.join('\n'));
}

</script>

<template>
  <div class="aceternity-page devices-page">
    <DevicesSearch v-model:model="searchParams" @reset="resetSearchParams" @search="getDataByPage" />

    <section class="device-table-card spotlight-card">
      <div class="table-toolbar">
        <div class="table-toolbar-summary">
          <span class="table-kicker">Device inventory</span>
          <h2>{{ $t('route.devices') }}<span class="table-count">· {{ searchParams.size }} / page</span></h2>
        </div>
        <div class="table-toolbar-actions">
          <div class="selection-pill" :class="{ 'is-active': checkedRowKeys.length > 0 }">
            <span class="status-dot" />
            <strong>{{ checkedRowKeys.length || 0 }}</strong>
            <small>selected</small>
            <button
              type="button"
              class="selection-pill-action"
              :disabled="selectedRustdeskIds.length === 0"
              @click="copySelectedIds"
            >
              <SvgIcon icon="solar:copy-linear" />
              {{ $t('page.devices.table.bulk.copy' as App.I18n.I18nKey) }}
            </button>
            <NPopconfirm :disabled="checkedRowKeys.length === 0" @positive-click="handleBatchDelete">
              <template #trigger>
                <button type="button" class="selection-pill-action is-danger" :disabled="checkedRowKeys.length === 0">
                  <SvgIcon icon="solar:trash-bin-trash-linear" />
                  {{ $t('common.batchDelete') }}
                </button>
              </template>
              {{ $t('common.confirmDelete') }}
            </NPopconfirm>
          </div>
          <TableHeader v-model:columns="columnChecks" :loading="loading" @refresh="getData" />
        </div>
      </div>
      <NDataTable
        :columns="columns"
        :data="data"
        :flex-height="!appStore.isMobile"
        :scroll-x="1520"
        :loading="loading"
        remote
        :row-key="row => row.id"
        :pagination="mobilePagination"
        :row-class-name="() => 'device-row'"
        @update:sorter="onSorterChange"
        v-model:checked-row-keys="checkedRowKeys"
        class="device-table"
      />
    </section>
  </div>
</template>

<style scoped>
.devices-page {
  display: flex;
  min-height: 640px;
  flex-direction: column;
  gap: 20px;
  padding-bottom: 12px;
}

.device-table-card {
  display: flex;
  min-height: 600px;
  flex: 1;
  flex-direction: column;
  padding: 22px 22px 14px;
}

.table-toolbar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 18px;
  padding: 0 4px 18px;
}

.table-kicker {
  color: var(--accent-bright);
  font-size: 10px;
  font-weight: 760;
  letter-spacing: 0.14em;
  text-transform: uppercase;
}

.table-toolbar-summary h2 {
  margin: 6px 0 0;
  color: var(--text-strong);
  font-size: 22px;
  font-weight: 720;
  letter-spacing: -0.025em;
}

.table-count {
  margin-left: 12px;
  color: var(--text-muted);
  font-size: 11px;
  font-weight: 650;
  letter-spacing: 0.05em;
}

.table-toolbar-actions {
  display: flex;
  align-items: center;
  gap: 14px;
}

.selection-pill {
  display: inline-flex;
  align-items: center;
  gap: 10px;
  border: 2px solid var(--surface-border);
  border-radius: 999px;
  background: var(--surface-muted);
  color: var(--text-muted);
  font-size: 11px;
  font-weight: 700;
  letter-spacing: 0.05em;
  padding: 8px 8px 8px 14px;
  transition:
    background-color 160ms ease,
    color 160ms ease;
}

.selection-pill.is-active {
  background: var(--accent-bright);
  color: #ffffff;
}

.selection-pill strong {
  font-size: 13px;
  font-weight: 780;
}

.selection-pill small {
  margin-right: 4px;
  color: var(--text-muted);
  font-size: 10px;
  letter-spacing: 0.08em;
  text-transform: uppercase;
}

.selection-pill.is-active small {
  color: var(--accent-bright);
}

.selection-pill-action {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  border: 2px solid var(--surface-border);
  border-radius: 999px;
  background: var(--surface);
  color: var(--text-strong);
  font-size: 11px;
  font-weight: 700;
  letter-spacing: 0.04em;
  padding: 6px 12px;
  box-shadow: var(--shadow-hard-sm);
  transition:
    transform 140ms ease,
    box-shadow 140ms ease,
    opacity 140ms ease;
}

.selection-pill-action:hover:not([disabled]) {
  transform: translate(-1px, -1px);
  box-shadow: 3px 3px 0 0 var(--shadow-color);
}

.selection-pill-action:active:not([disabled]) {
  transform: translate(1px, 1px);
  box-shadow: 1px 1px 0 0 var(--shadow-color);
}

.selection-pill-action[disabled] {
  opacity: 0.45;
  cursor: not-allowed;
}

.selection-pill-action.is-danger {
  background: #fca5a5;
  color: #7f1d1d;
}

.selection-pill-action.is-danger:hover:not([disabled]) {
  background: #f87171;
}

.device-table {
  position: relative;
  z-index: 1;
  flex: 1;
}

.device-table :deep(.n-data-table-tr.device-row) {
  transition:
    background-color 180ms ease,
    box-shadow 180ms ease;
}

.device-table :deep(.n-data-table-tr.device-row:hover) {
  box-shadow: inset 3px 0 0 0 var(--accent-bright);
}

.device-table :deep(.n-data-table-tr.device-row.is-selected) {
  background: color-mix(in srgb, var(--accent-bright) 16%, var(--surface)) !important;
}

.device-table :deep(.cell-index) {
  color: var(--text-muted);
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 11px;
  letter-spacing: 0.06em;
}

.device-table :deep(.cell-host) {
  display: flex;
  align-items: center;
  gap: 12px;
}

.device-table :deep(.cell-host-avatar) {
  display: grid;
  width: 36px;
  height: 36px;
  flex: 0 0 36px;
  place-items: center;
  border: 2px solid var(--surface-border);
  border-radius: var(--radius-sm);
  background: var(--accent-bright);
  color: white;
  font-size: 13px;
  font-weight: 800;
  letter-spacing: -0.01em;
}

.device-table :deep(.cell-host-meta) {
  min-width: 0;
}

.device-table :deep(.cell-host-name) {
  display: block;
  overflow: hidden;
  color: var(--text-strong);
  font-size: 13px;
  font-weight: 700;
  letter-spacing: -0.01em;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.device-table :deep(.cell-host-sub) {
  display: block;
  margin-top: 3px;
  color: var(--text-muted);
  font-size: 10px;
  font-weight: 600;
  letter-spacing: 0.04em;
  text-overflow: ellipsis;
  white-space: nowrap;
  overflow: hidden;
}

.device-table :deep(.cell-id) {
  display: inline-flex;
  align-items: center;
  gap: 8px;
}

.device-table :deep(.cell-id-dot) {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: #94a3b8;
  border: 1.5px solid var(--surface-border);
}

.device-table :deep(.cell-id-dot.is-online) {
  background: #22c55e;
}

.device-table :deep(.cell-id-text) {
  color: var(--text-strong);
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 11px;
  letter-spacing: 0.02em;
}

.device-table :deep(.cell-name) {
  color: var(--text-strong);
  font-size: 12px;
  font-weight: 650;
}

.device-table :deep(.cell-name.muted) {
  color: var(--text-muted);
  font-weight: 600;
}

.device-table :deep(.cell-time) {
  display: inline-flex;
  flex-direction: column;
  align-items: flex-start;
  gap: 4px;
}

.device-table :deep(.cell-time-text) {
  color: var(--text-strong);
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 10.5px;
}

.device-table :deep(.cell-time-text.muted) {
  color: var(--text-muted);
}

.device-table :deep(.cell-version) {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  gap: 4px;
}

.device-table :deep(.cell-version-pill) {
  display: inline-flex;
  align-items: center;
  border: 1.5px solid var(--surface-border);
  border-radius: 999px;
  background: var(--surface-muted);
  color: var(--accent-cyan);
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 10px;
  font-weight: 700;
  padding: 4px 8px;
}

.device-table :deep(.cell-version-meta) {
  color: var(--text-muted);
  font-size: 10px;
  font-weight: 600;
  letter-spacing: 0.04em;
}

.device-table :deep(.cell-os) {
  color: var(--text-muted);
  font-size: 11px;
  font-weight: 650;
}

.device-table :deep(.cell-status) {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  font-size: 11px;
  font-weight: 760;
  letter-spacing: 0.05em;
}

.device-table :deep(.cell-status-dot) {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: #94a3b8;
  border: 1.5px solid var(--surface-border);
}

.device-table :deep(.cell-status.is-online .cell-status-dot) {
  background: #22c55e;
}

.device-table :deep(.cell-status.is-online .cell-status-label) {
  color: var(--accent-bright);
}

.device-table :deep(.cell-status.is-offline .cell-status-label) {
  color: var(--text-muted);
}

@media (max-width: 880px) {
  .device-table-card {
    padding: 14px;
  }
}
</style>
