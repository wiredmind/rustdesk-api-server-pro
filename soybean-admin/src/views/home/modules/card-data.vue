<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import { $t } from '@/locales';
import { fetchStat } from '@/service/api/home';

defineOptions({
  name: 'CardData'
});

interface CardData {
  key: string;
  title: string;
  value: number;
  unit: string;
  color: {
    start: string;
    end: string;
  };
  icon: string;
}

const stat = ref<Api.Home.Stat>({
  userCount: 0,
  deviceCount: 0,
  onlineCount: 0,
  visitsCount: 0
});

const cardData = computed<CardData[]>(() => [
  {
    key: 'userCount',
    title: $t('page.home.userCount'),
    value: stat.value.userCount,
    unit: '',
    color: {
      start: '#ec4786',
      end: '#b955a4'
    },
    icon: 'gravity-ui:person'
  },
  {
    key: 'deviceCount',
    title: $t('page.home.deviceCount'),
    value: stat.value.deviceCount,
    unit: '',
    color: {
      start: '#865ec0',
      end: '#5144b4'
    },
    icon: 'fluent:desktop-32-regular'
  },
  {
    key: 'onlineCount',
    title: $t('page.home.onlineCount'),
    value: stat.value.onlineCount,
    unit: '',
    color: {
      start: '#56cdf3',
      end: '#719de3'
    },
    icon: 'fluent:desktop-checkmark-20-regular'
  },
  {
    key: 'visitCount',
    title: $t('page.home.visitsCount'),
    value: stat.value.visitsCount,
    unit: '',
    color: {
      start: '#fcbc25',
      end: '#f68057'
    },
    icon: 'ant-design:bar-chart-outlined'
  }
]);


onMounted(async () => {
  const s = (await fetchStat()).data;
  if (s !== null) {
    stat.value = s;
  }
});
</script>

<template>
  <section class="metric-grid">
    <article
      v-for="item in cardData"
      :key="item.key"
      class="metric-card spotlight-card"
      :style="{ '--metric-start': item.color.start, '--metric-end': item.color.end }"
    >
      <div class="metric-topline">
        <span class="metric-icon"><SvgIcon :icon="item.icon" /></span>
        <span class="metric-label">{{ item.title }}</span>
      </div>
      <div class="metric-value">
        <CountTo :prefix="item.unit" :start-value="0" :end-value="item.value" />
      </div>
      <div class="metric-footer">
        <span>Live overview</span>
        <span class="metric-pulse"></span>
      </div>
    </article>
  </section>
</template>

<style scoped>
.metric-grid {
  display: grid;
  grid-template-columns: repeat(4, minmax(0, 1fr));
  gap: 16px;
}

.metric-card {
  min-height: 168px;
  padding: 20px;
  transition:
    transform 220ms ease,
    border-color 220ms ease,
    box-shadow 220ms ease;
}

.metric-card::after {
  position: absolute;
  right: -38px;
  bottom: -60px;
  width: 150px;
  height: 150px;
  border-radius: 50%;
  background: radial-gradient(circle, var(--metric-start), transparent 68%);
  content: '';
  filter: blur(12px);
  opacity: 0.15;
  pointer-events: none;
}

.metric-card:hover {
  border-color: var(--surface-border-strong);
  box-shadow: 0 28px 80px rgba(2, 6, 23, 0.25);
  transform: translateY(-4px);
}

.metric-topline {
  position: relative;
  z-index: 1;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.metric-icon {
  display: grid;
  width: 40px;
  height: 40px;
  place-items: center;
  border: 1px solid color-mix(in srgb, var(--metric-start) 38%, transparent);
  border-radius: 12px;
  background: color-mix(in srgb, var(--metric-start) 14%, transparent);
  color: var(--metric-start);
  font-size: 22px;
}

.metric-label {
  color: var(--text-muted);
  font-size: 12px;
  font-weight: 650;
  letter-spacing: 0.04em;
}

.metric-value {
  position: relative;
  z-index: 1;
  margin-top: 18px;
  color: var(--text-strong);
  font-size: 38px;
  font-weight: 760;
  letter-spacing: -0.055em;
  line-height: 1;
}

.metric-footer {
  position: relative;
  z-index: 1;
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-top: 18px;
  color: var(--text-muted);
  font-size: 10px;
  font-weight: 650;
  letter-spacing: 0.08em;
  text-transform: uppercase;
}

.metric-pulse {
  width: 34px;
  height: 3px;
  border-radius: 999px;
  background: linear-gradient(90deg, var(--metric-start), var(--metric-end));
  box-shadow: 0 0 14px color-mix(in srgb, var(--metric-start) 55%, transparent);
}

@media (max-width: 1080px) {
  .metric-grid {
    grid-template-columns: repeat(2, minmax(0, 1fr));
  }
}

@media (max-width: 560px) {
  .metric-grid {
    grid-template-columns: 1fr;
  }
}
</style>
