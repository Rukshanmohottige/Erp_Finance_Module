<template>
  <div>
    <div class="page-header">
      <h2>Dashboard</h2>
      <span style="color: var(--text-muted); font-size: 14px;">Fiscal Year 2025</span>
    </div>

    <div class="stats-grid">
      <div class="stat-card">
        <div class="label">Total Budget</div>
        <div class="value" style="color: var(--primary)">${{ fmt(summary.total_budget) }}</div>
        <div class="sub">Allocated this year</div>
      </div>
      <div class="stat-card">
        <div class="label">Total Spent</div>
        <div class="value" style="color: var(--warning)">${{ fmt(summary.total_spent) }}</div>
        <div class="sub">{{ summary.budget_utilization?.toFixed(1) }}% utilized</div>
      </div>
      <div class="stat-card">
        <div class="label">Pending Expenses</div>
        <div class="value" style="color: var(--danger)">{{ summary.pending_expenses }}</div>
        <div class="sub">Awaiting approval</div>
      </div>
      <div class="stat-card">
        <div class="label">Unpaid Invoices</div>
        <div class="value" style="color: var(--warning)">{{ summary.unpaid_invoices }}</div>
        <div class="sub">Requires payment</div>
      </div>
      <div class="stat-card">
        <div class="label">Payroll Pending</div>
        <div class="value" style="color: var(--text)">{{ summary.pending_payroll }}</div>
        <div class="sub">To be processed</div>
      </div>
    </div>

    <div class="card">
      <h3 style="margin-bottom: 16px; font-size: 16px;">Budget Utilization by Department</h3>
      <div v-if="budgets.length" class="budget-bars">
        <div v-for="b in budgets" :key="b.id" class="budget-row">
          <div class="budget-dept">{{ b.department }}</div>
          <div class="budget-bar-wrap">
            <div class="budget-bar" :style="{ width: pct(b.spent_amount, b.allocated_amount) + '%' }"></div>
          </div>
          <div class="budget-nums">${{ fmt(b.spent_amount) }} / ${{ fmt(b.allocated_amount) }}</div>
        </div>
      </div>
      <p v-else style="color: var(--text-muted); font-size: 14px;">Loading...</p>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import api from '../../api/index.js'

const summary = ref({})
const budgets = ref([])

const fmt = n => Number(n || 0).toLocaleString()
const pct = (spent, total) => total > 0 ? Math.min((spent / total) * 100, 100).toFixed(1) : 0

onMounted(async () => {
  const [s, b] = await Promise.all([
    api.get('/dashboard/summary'),
    api.get('/budgets')
  ])
  summary.value = s.data
  budgets.value = b.data || []
})
</script>

<style scoped>
.budget-bars { display: flex; flex-direction: column; gap: 14px; }
.budget-row { display: flex; align-items: center; gap: 14px; }
.budget-dept { width: 120px; font-size: 13px; font-weight: 500; flex-shrink: 0; }
.budget-bar-wrap { flex: 1; background: var(--border); border-radius: 6px; height: 10px; overflow: hidden; }
.budget-bar { height: 100%; background: var(--primary); border-radius: 6px; transition: width 0.6s; }
.budget-nums { width: 160px; font-size: 12px; color: var(--text-muted); text-align: right; flex-shrink: 0; }
</style>
