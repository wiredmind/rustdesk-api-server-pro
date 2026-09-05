<script setup lang="ts">
import HeaderBanner from './modules/header-banner.vue';
import CardData from './modules/card-data.vue';
import LineChart from './modules/line-chart.vue';
import PieChart from './modules/pie-chart.vue';
</script>

<template>
  <div class="aceternity-page dashboard-page">
    <HeaderBanner />
    <CardData />

    <section class="section-heading">
      <div>
        <span class="eyebrow">Signal intelligence</span>
        <h2>Infrastructure at a glance</h2>
        <p>Live connection activity and operating-system distribution across your private fleet.</p>
      </div>
      <div class="live-chip">
        <span class="status-dot"></span>
        Updating live
      </div>
    </section>

    <div class="analytics-grid">
      <LineChart />
      <PieChart />
    </div>
  </div>
</template>

<style scoped>
.dashboard-page {
  display: flex;
  flex-direction: column;
  gap: 18px;
}

.section-heading {
  display: flex;
  align-items: end;
  justify-content: space-between;
  gap: 24px;
  padding: 18px 4px 0;
}

.section-heading h2 {
  margin: 14px 0 7px;
  color: var(--text-strong);
  font-size: 25px;
  font-weight: 740;
  letter-spacing: -0.035em;
}

.section-heading p {
  margin: 0;
  color: var(--text-muted);
  line-height: 1.6;
}

.live-chip {
  display: inline-flex;
  flex: 0 0 auto;
  align-items: center;
  gap: 9px;
  border: 1px solid var(--surface-border);
  border-radius: 999px;
  background: var(--surface);
  color: var(--text-muted);
  font-size: 11px;
  font-weight: 650;
  padding: 9px 12px;
}

.analytics-grid {
  display: grid;
  grid-template-columns: minmax(0, 1.45fr) minmax(320px, 0.55fr);
  gap: 18px;
}

.analytics-grid :deep(.n-card) {
  height: 100%;
}

@media (max-width: 1024px) {
  .analytics-grid {
    grid-template-columns: 1fr;
  }
}

@media (max-width: 640px) {
  .section-heading {
    align-items: flex-start;
    flex-direction: column;
  }
}
</style>
