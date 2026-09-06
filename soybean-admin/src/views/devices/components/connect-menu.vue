<script setup lang="ts">
import { ref } from 'vue';
import { NPopconfirm, NPopover, useMessage } from 'naive-ui';
import { deleteDevices } from '@/service/api/devices';
import { $t } from '@/locales';

interface Props {
  deviceId?: number;
  rustdeskId: string;
  hostname?: string;
}

const props = defineProps<Props>();
const emit = defineEmits<{ deleted: [] }>();

const message = useMessage();
const popoverVisible = ref(false);

type ConnectAction = 'remote' | 'file' | 'mirror' | 'terminal';

// Maps a portal action to the RustDesk client's own rustdesk://<authority>/<id>
// uni-link scheme, as read from the RustDesk client source (core_main.rs):
// --connect, --play, --file-transfer map to authorities "connect", "play",
// "file-transfer"; the built-in terminal session uses the "terminal"
// authority (confirmed by the RustDesk maintainers). There is no supported
// way to inject or execute an arbitrary shell command through this link.
const actionAuthority: Record<ConnectAction, string> = {
  remote: 'connect',
  file: 'file-transfer',
  mirror: 'play',
  terminal: 'terminal'
};

function actionTitle(action: ConnectAction): string {
  switch (action) {
    case 'remote':
      return $t('page.devices.connect.remote' as App.I18n.I18nKey);
    case 'file':
      return $t('page.devices.connect.file' as App.I18n.I18nKey);
    case 'mirror':
      return $t('page.devices.connect.mirror' as App.I18n.I18nKey);
    case 'terminal':
      return $t('page.devices.connect.terminal' as App.I18n.I18nKey);
  }
}

function launch(action: ConnectAction) {
  if (!props.rustdeskId) {
    message.warning($t('page.devices.connect.missingId' as App.I18n.I18nKey));
    return;
  }
  const url = `rustdesk://${actionAuthority[action]}/${props.rustdeskId}`;
  try {
    window.open(url, '_blank', 'noopener,noreferrer');
    message.success(
      $t('page.devices.connect.launched' as App.I18n.I18nKey, {
        action: actionTitle(action),
        host: props.hostname || props.rustdeskId
      })
    );
  } catch {
    message.error($t('page.devices.connect.failed' as App.I18n.I18nKey));
  }
  popoverVisible.value = false;
}

function copyId() {
  if (!props.rustdeskId) {
    message.warning($t('page.devices.connect.missingId' as App.I18n.I18nKey));
    return;
  }
  if (navigator?.clipboard?.writeText) {
    void navigator.clipboard.writeText(props.rustdeskId).then(() => {
      message.success($t('page.devices.connect.copied' as App.I18n.I18nKey, { id: props.rustdeskId }));
    });
    return;
  }
  message.info(props.rustdeskId);
}

async function handleDelete() {
  if (!props.deviceId) {
    message.warning($t('page.devices.connect.missingId' as App.I18n.I18nKey));
    return;
  }
  const res = await deleteDevices({ ids: [props.deviceId] });
  if (res.error === null) {
    message.success($t('common.deleteSuccess'));
    popoverVisible.value = false;
    emit('deleted');
  }
}
</script>

<template>
  <NPopover v-model:show="popoverVisible" trigger="click" placement="bottom-end" :show-arrow="false" raw>
    <template #trigger>
      <button type="button" class="connect-button">
        <SvgIcon icon="solar:bolt-circle-bold-duotone" />
        {{ $t('page.devices.connect.cta' as App.I18n.I18nKey) }}
        <SvgIcon icon="solar:alt-arrow-down-linear" class="connect-caret" />
      </button>
    </template>
    <div class="connect-popover">
      <header class="connect-popover-head">
        <span class="connect-popover-eyebrow">Connect from portal</span>
        <strong>{{ props.hostname || props.rustdeskId }}</strong>
        <code>{{ props.rustdeskId }}</code>
      </header>

      <section class="connect-popover-actions">
        <button class="connect-card" type="button" @click="launch('remote')">
          <SvgIcon icon="solar:monitor-smartphone-bold-duotone" />
          <div>
            <strong>{{ $t('page.devices.connect.remote' as App.I18n.I18nKey) }}</strong>
            <small>Full control</small>
          </div>
        </button>
        <button class="connect-card" type="button" @click="launch('file')">
          <SvgIcon icon="solar:upload-square-bold-duotone" />
          <div>
            <strong>{{ $t('page.devices.connect.file' as App.I18n.I18nKey) }}</strong>
            <small>Send or receive</small>
          </div>
        </button>
        <button class="connect-card" type="button" @click="launch('mirror')">
          <SvgIcon icon="solar:tv-bold-duotone" />
          <div>
            <strong>{{ $t('page.devices.connect.mirror' as App.I18n.I18nKey) }}</strong>
            <small>View only session</small>
          </div>
        </button>
        <button class="connect-card" type="button" @click="launch('terminal')">
          <SvgIcon icon="solar:programming-bold-duotone" />
          <div>
            <strong>{{ $t('page.devices.connect.terminal' as App.I18n.I18nKey) }}</strong>
            <small>RustDesk built-in shell</small>
          </div>
        </button>
        <button class="connect-card connect-card-wide" type="button" @click="copyId">
          <SvgIcon icon="solar:copy-bold-duotone" />
          <div>
            <strong>Copy ID</strong>
            <small>Share with family</small>
          </div>
        </button>
      </section>

      <footer class="connect-popover-foot">
        <SvgIcon icon="solar:info-circle-linear" />
        <span>{{ $t('page.devices.connect.terminalHint' as App.I18n.I18nKey) }}</span>
      </footer>

      <NPopconfirm @positive-click="handleDelete">
        <template #trigger>
          <button type="button" class="connect-delete">
            <SvgIcon icon="solar:trash-bin-trash-bold-duotone" />
            {{ $t('common.delete') }}
          </button>
        </template>
        {{ $t('common.confirmDelete') }}
      </NPopconfirm>
    </div>
  </NPopover>
</template>

<style scoped>
.connect-button {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  border: 2px solid var(--surface-border);
  border-radius: var(--radius-sm);
  background: var(--accent-bright);
  color: #ffffff;
  font-size: 11px;
  font-weight: 800;
  letter-spacing: 0.05em;
  white-space: nowrap;
  padding: 7px 11px;
  box-shadow: var(--shadow-hard-sm);
  transition:
    transform 140ms ease,
    box-shadow 140ms ease;
}

.connect-button:hover {
  transform: translate(-1px, -1px);
  box-shadow: 3px 3px 0 0 var(--shadow-color);
}

.connect-button :deep(svg) {
  font-size: 14px;
}

.connect-caret {
  opacity: 0.7;
  font-size: 12px !important;
}

.connect-popover {
  position: relative;
  width: 340px;
  padding: 22px;
  background: var(--surface-strong);
  border: 3px solid var(--surface-border);
  border-radius: var(--radius);
  box-shadow: var(--shadow-hard-lg);
}

.connect-popover-head {
  display: flex;
  flex-direction: column;
  gap: 4px;
  border-bottom: 1px solid var(--surface-border);
  padding-bottom: 14px;
}

.connect-popover-eyebrow {
  color: var(--accent-bright);
  font-size: 9px;
  font-weight: 760;
  letter-spacing: 0.12em;
  text-transform: uppercase;
}

.connect-popover-head strong {
  color: var(--text-strong);
  font-size: 15px;
  font-weight: 720;
}

.connect-popover-head code {
  color: var(--text-muted);
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 10px;
}

.connect-popover-actions {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 10px;
  margin: 14px 0;
}

.connect-card {
  display: flex;
  align-items: center;
  gap: 10px;
  border: 2px solid var(--surface-border);
  border-radius: var(--radius-sm);
  background: var(--surface);
  color: var(--text-strong);
  font: inherit;
  text-align: left;
  padding: 10px 12px;
  transition:
    transform 140ms ease,
    box-shadow 140ms ease;
}

.connect-card-wide {
  grid-column: 1 / -1;
}

.connect-card:hover {
  transform: translate(-1px, -1px);
  box-shadow: var(--shadow-hard-sm);
}

.connect-card :deep(svg) {
  color: var(--accent-bright);
  font-size: 18px;
}

.connect-card strong,
.connect-card small {
  display: block;
}

.connect-card strong {
  font-size: 12px;
  font-weight: 720;
}

.connect-card small {
  margin-top: 2px;
  color: var(--text-muted);
  font-size: 9px;
  font-weight: 650;
  letter-spacing: 0.04em;
  text-transform: uppercase;
}

.connect-popover-foot {
  display: flex;
  align-items: flex-start;
  gap: 8px;
  margin-top: 4px;
  border-top: 1px solid var(--surface-border);
  padding-top: 14px;
  color: var(--text-muted);
  font-size: 10px;
  font-weight: 600;
  letter-spacing: 0.02em;
  line-height: 1.5;
}

.connect-popover-foot :deep(svg) {
  flex: 0 0 auto;
  margin-top: 1px;
  color: var(--accent-cyan);
  font-size: 14px;
}

.connect-delete {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  width: 100%;
  margin-top: 12px;
  border: 2px solid var(--surface-border);
  border-radius: var(--radius-sm);
  background: #fca5a5;
  color: #7f1d1d;
  font-size: 12px;
  font-weight: 800;
  letter-spacing: 0.02em;
  padding: 10px 12px;
  transition:
    transform 140ms ease,
    box-shadow 140ms ease;
}

.connect-delete:hover {
  transform: translate(-1px, -1px);
  box-shadow: var(--shadow-hard-sm);
}

.connect-delete :deep(svg) {
  font-size: 15px;
}
</style>
