<template>
  <div>
    <div class="page-header"><h2>Reports</h2></div>
    <div style="display:grid; grid-template-columns: 1fr 1fr; gap: 20px;">
      <div class="card">
        <h3 style="margin-bottom:16px; font-size:15px;">Budget vs Actual Spend</h3>
        <table>
          <thead><tr><th>Department</th><th>Allocated</th><th>Spent</th><th>%</th></tr></thead>
          <tbody>
            <tr v-for="r in budgetReport" :key="r.department">
              <td>{{ r.department }}</td>
              <td>${{ fmt(r.allocated) }}</td>
              <td>${{ fmt(r.spent) }}</td>
              <td :style="{ color: pct(r.spent, r.allocated) > 90 ? 'var(--danger)' : 'var(--success)' }">
                {{ pct(r.spent, r.allocated) }}%
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <div class="card">
        <h3 style="margin-bottom:16px; font-size:15px;">Expenses by Category</h3>
        <table>
          <thead><tr><th>Category</th><th>Total Spent</th></tr></thead>
          <tbody>
            <tr v-for="r in categoryReport" :key="r.category">
              <td>{{ r.category }}</td>
              <td><strong>${{ fmt(r.total) }}</strong></td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import api from '../../api/index.js'

const budgetReport = ref([])
const categoryReport = ref([])

const fmt = n => Number(n || 0).toLocaleString()
const pct = (s, t) => t > 0 ? ((s / t) * 100).toFixed(1) : '0.0'

onMounted(async () => {
  const [b, c] = await Promise.all([
    api.get('/reports/budget-vs-actual'),
    api.get('/reports/expense-by-category')
  ])
  budgetReport.value = b.data || []
  categoryReport.value = c.data || []
})
</script>
