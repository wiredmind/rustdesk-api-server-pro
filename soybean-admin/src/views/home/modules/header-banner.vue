<script setup lang="ts">
import { useAuthStore } from '@/store/modules/auth';

defineOptions({
  name: 'HeaderBanner'
});

const authStore = useAuthStore();
</script>

<template>
  <section class="dashboard-hero spotlight-card">
    <div class="hero-copy">
      <div class="eyebrow">
        <span class="status-dot"></span>
        Systems operational
      </div>
      <h1>
        Your remote world,
        <span>beautifully under control.</span>
      </h1>
      <p>
        Welcome back, {{ authStore.userInfo.userName }}. Every trusted computer, session and signal is right where
        it belongs.
      </p>
      <div class="hero-meta">
        <span><SvgIcon icon="solar:shield-keyhole-bold-duotone" /> Self-hosted</span>
        <span><SvgIcon icon="solar:server-square-cloud-bold-duotone" /> Private relay</span>
        <span><SvgIcon icon="solar:bolt-circle-bold-duotone" /> Live telemetry</span>
      </div>
    </div>
    <div class="hero-orbit" aria-hidden="true">
      <div class="orbit orbit-one"></div>
      <div class="orbit orbit-two"></div>
      <div class="hero-core">
        <SvgIcon icon="solar:monitor-smartphone-bold-duotone" />
      </div>
      <span class="orbit-node node-one"></span>
      <span class="orbit-node node-two"></span>
      <span class="orbit-node node-three"></span>
    </div>
  </section>
</template>

<style scoped>
.dashboard-hero {
  min-height: 280px;
  display: grid;
  grid-template-columns: minmax(0, 1.4fr) minmax(260px, 0.6fr);
  align-items: center;
  padding: 38px 42px;
}

.hero-copy {
  position: relative;
  z-index: 2;
  max-width: 760px;
}

.hero-copy h1 {
  margin: 20px 0 14px;
  color: var(--text-strong);
  font-size: clamp(32px, 4vw, 54px);
  font-weight: 780;
  letter-spacing: -0.055em;
  line-height: 1.02;
}

.hero-copy h1 span {
  display: block;
  background: linear-gradient(100deg, #c4b5fd 8%, #8b5cf6 48%, #67e8f9 94%);
  background-clip: text;
  color: transparent;
}

.hero-copy p {
  max-width: 650px;
  margin: 0;
  color: var(--text-muted);
  font-size: 15px;
  line-height: 1.7;
}

.hero-meta {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
  margin-top: 24px;
}

.hero-meta span {
  display: inline-flex;
  align-items: center;
  gap: 7px;
  border: 1px solid var(--surface-border);
  border-radius: 999px;
  background: var(--surface-muted);
  color: var(--text-muted);
  font-size: 11px;
  font-weight: 650;
  padding: 8px 11px;
}

.hero-meta span :deep(svg) {
  color: var(--accent);
  font-size: 15px;
}

.hero-orbit {
  position: relative;
  width: 240px;
  height: 240px;
  justify-self: center;
}

.hero-core {
  position: absolute;
  z-index: 2;
  top: 50%;
  left: 50%;
  display: grid;
  width: 82px;
  height: 82px;
  border: 1px solid rgba(196, 181, 253, 0.38);
  border-radius: 24px;
  background: linear-gradient(145deg, rgba(139, 92, 246, 0.95), rgba(79, 70, 229, 0.88));
  box-shadow:
    0 0 0 14px rgba(139, 92, 246, 0.08),
    0 26px 70px rgba(76, 29, 149, 0.5),
    inset 0 1px 0 rgba(255, 255, 255, 0.35);
  color: white;
  font-size: 42px;
  place-items: center;
  transform: translate(-50%, -50%);
}

.orbit {
  position: absolute;
  inset: 18px;
  border: 1px solid rgba(139, 92, 246, 0.22);
  border-radius: 50%;
}

.orbit-two {
  inset: 48px;
  border-color: rgba(34, 211, 238, 0.2);
}

.orbit-node {
  position: absolute;
  width: 11px;
  height: 11px;
  border-radius: 50%;
  background: var(--accent-cyan);
  box-shadow: 0 0 24px rgba(34, 211, 238, 0.8);
}

.node-one {
  top: 42px;
  left: 48px;
}

.node-two {
  right: 22px;
  bottom: 92px;
  background: var(--accent);
  box-shadow: 0 0 24px rgba(167, 139, 250, 0.8);
}

.node-three {
  bottom: 28px;
  left: 74px;
  width: 7px;
  height: 7px;
}

@media (max-width: 900px) {
  .dashboard-hero {
    grid-template-columns: 1fr;
    padding: 30px;
  }

  .hero-orbit {
    display: none;
  }
}

@media (max-width: 560px) {
  .dashboard-hero {
    min-height: auto;
    padding: 24px 20px;
  }

  .hero-copy h1 {
    font-size: 34px;
  }
}
</style>
