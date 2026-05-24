<template>
  <div class="layout">
    <aside class="sidebar">
      <div class="sidebar-logo">
        <span class="logo-icon">💰</span>
        <span class="logo-text">ERP Finance</span>
      </div>
      <nav class="sidebar-nav">
        <router-link v-for="item in navItems" :key="item.path" :to="item.path" class="nav-item">
          <span class="nav-icon">{{ item.icon }}</span>
          <span>{{ item.label }}</span>
        </router-link>
      </nav>
      <div class="sidebar-footer">
        <div class="user-info">
          <div class="avatar">{{ userInitial }}</div>
          <div>
            <div class="user-name">{{ user?.name }}</div>
            <div class="user-role">{{ user?.role }}</div>
          </div>
        </div>
        <button class="logout-btn" @click="logout">Logout</button>
      </div>
    </aside>
    <main class="content">
      <router-view />
    </main>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../store/auth.js'

const auth = useAuthStore()
const router = useRouter()
const user = computed(() => auth.user)
const userInitial = computed(() => user.value?.name?.[0]?.toUpperCase() || 'U')

const navItems = [
  { path: '/dashboard', icon: '📊', label: 'Dashboard' },
  { path: '/budgets',   icon: '📁', label: 'Budgets' },
  { path: '/expenses',  icon: '💳', label: 'Expenses' },
  { path: '/invoices',  icon: '🧾', label: 'Invoices' },
  { path: '/payroll',   icon: '👥', label: 'Payroll' },
  { path: '/reports',   icon: '📈', label: 'Reports' },
]

function logout() {
  auth.logout()
  router.push('/login')
}
</script>

<style scoped>
.layout { display: flex; height: 100vh; overflow: hidden; }
.sidebar {
  width: 240px; background: var(--sidebar-bg); display: flex;
  flex-direction: column; flex-shrink: 0;
}
.sidebar-logo {
  padding: 20px 16px; display: flex; align-items: center; gap: 10px;
  border-bottom: 1px solid rgba(255,255,255,0.08);
}
.logo-icon { font-size: 22px; }
.logo-text { color: white; font-weight: 700; font-size: 16px; }
.sidebar-nav { flex: 1; padding: 16px 8px; display: flex; flex-direction: column; gap: 2px; }
.nav-item {
  display: flex; align-items: center; gap: 10px; padding: 10px 12px;
  border-radius: 8px; color: var(--sidebar-text); text-decoration: none;
  font-size: 14px; transition: all 0.2s;
}
.nav-item:hover { background: rgba(255,255,255,0.06); color: white; }
.nav-item.router-link-active { background: var(--sidebar-active); color: white; }
.nav-icon { font-size: 16px; }
.sidebar-footer { padding: 16px; border-top: 1px solid rgba(255,255,255,0.08); }
.user-info { display: flex; align-items: center; gap: 10px; margin-bottom: 12px; }
.avatar {
  width: 34px; height: 34px; background: var(--primary); border-radius: 50%;
  display: flex; align-items: center; justify-content: center; color: white; font-weight: 700;
}
.user-name { color: white; font-size: 13px; font-weight: 500; }
.user-role { color: var(--sidebar-text); font-size: 11px; text-transform: capitalize; }
.logout-btn {
  width: 100%; padding: 8px; background: rgba(255,255,255,0.07); border: none;
  border-radius: 6px; color: var(--sidebar-text); cursor: pointer; font-size: 13px;
}
.logout-btn:hover { background: rgba(255,255,255,0.12); color: white; }
.content { flex: 1; overflow-y: auto; padding: 28px; }
</style>
