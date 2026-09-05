<script setup lang="ts">
import { computed, onMounted, reactive } from 'vue';
import { $t } from '@/locales';
import { useNaiveForm } from '@/hooks/common/form';
import { useAuthStore } from '@/store/modules/auth';
import { fetchCaptcha } from '@/service/api/auth';

defineOptions({
  name: 'PwdLogin'
});

const authStore = useAuthStore();
const { formRef, validate } = useNaiveForm();

const model: Api.Form.LoginForm = reactive({
  username: '',
  password: '',
  code: '',
  captchaId: ''
});

const captcha: Api.Auth.Captcha = reactive({
  id: '',
  img: ''
});

const rules = computed<Record<keyof Api.Form.LoginForm, App.Global.FormRule[]>>(() => {
  return {
    username: [
      {
        required: true,
        message: 'Username is required'
      }
    ],
    password: [
      {
        required: true,
        message: 'Password is required'
      }
    ],
    code: [
      {
        required: true,
        message: 'Verification code is required'
      }
    ],
    captchaId: [
      {
        required: true,
        message: 'Verification code is required'
      }
    ]
  };
});

async function handleSubmit() {
  await validate();
  const err = await authStore.login(model);
  if (err?.response?.data.message === 'CaptchaError') {
    handleCaptcha();
  }
}

async function handleCaptcha() {
  const c = await fetchCaptcha();
  captcha.id = c.data?.id || '';
  captcha.img = c.data?.img || '';
  model.captchaId = captcha.id || '';
}
onMounted(() => {
  handleCaptcha();
});
</script>

<template>
  <NForm ref="formRef" :model="model" :rules="rules" size="large" :show-label="false" class="auth-form">
    <NFormItem path="username">
      <NInput v-model:value="model.username" :placeholder="$t('page.login.common.userNamePlaceholder')">
        <template #prefix><SvgIcon icon="solar:user-rounded-linear" /></template>
      </NInput>
    </NFormItem>
    <NFormItem path="password">
      <NInput
        v-model:value="model.password"
        type="password"
        show-password-on="click"
        :placeholder="$t('page.login.common.passwordPlaceholder')"
      >
        <template #prefix><SvgIcon icon="solar:lock-keyhole-minimalistic-linear" /></template>
      </NInput>
    </NFormItem>
    <NFormItem path="code">
      <div class="captcha-row">
        <NInput v-model:value="model.code" clearable :placeholder="$t('page.login.common.codePlaceholder')">
          <template #prefix><SvgIcon icon="solar:shield-keyhole-linear" /></template>
        </NInput>
        <button class="captcha-frame" type="button" title="Refresh verification code" @click="handleCaptcha">
          <img width="152" height="40" :src="captcha.img" alt="Verification code" />
        </button>
      </div>
    </NFormItem>
    <div class="form-meta">
      <NCheckbox>{{ $t('page.login.pwdLogin.rememberMe') }}</NCheckbox>
      <span>Protected access</span>
    </div>
    <NButton
      attr-type="submit"
      type="primary"
      size="large"
      block
      :loading="authStore.loginLoading"
      class="submit-button"
      @click="handleSubmit"
    >
      <span class="inline-flex items-center gap-8px">
        {{ $t('common.confirm') }}
        <SvgIcon icon="solar:arrow-right-linear" />
      </span>
    </NButton>
  </NForm>
</template>

<style scoped>
.auth-form :deep(.n-form-item) {
  margin-bottom: 5px;
}

.auth-form :deep(.n-input) {
  min-height: 48px;
  border: 1px solid var(--surface-border);
  background: rgba(15, 23, 42, 0.58) !important;
}

.auth-form :deep(.n-input__prefix) {
  margin-right: 8px;
  color: var(--accent-bright);
  font-size: 18px;
}

.captcha-row {
  display: grid;
  width: 100%;
  grid-template-columns: minmax(0, 1fr) 134px;
  gap: 10px;
}

.captcha-frame {
  height: 48px;
  overflow: hidden;
  border: 1px solid var(--surface-border);
  border-radius: 11px;
  background: #f8fafc;
  cursor: pointer;
  padding: 3px;
  transition:
    border-color 160ms ease,
    transform 160ms ease;
}

.captcha-frame:hover {
  border-color: var(--surface-border-strong);
  transform: translateY(-1px);
}

.captcha-frame img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.form-meta {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin: 4px 0 20px;
}

.form-meta > span {
  color: var(--text-muted);
  font-size: 10px;
  font-weight: 650;
  letter-spacing: 0.08em;
  text-transform: uppercase;
}

.submit-button {
  min-height: 49px;
  font-size: 14px;
  letter-spacing: 0.015em;
}

@media (max-width: 420px) {
  .captcha-row {
    grid-template-columns: 1fr;
  }

  .captcha-frame {
    width: 100%;
  }
}
</style>
