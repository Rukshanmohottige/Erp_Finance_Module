<template>
  <div>
    <div class="page-header">
      <h2>Budgets</h2>
      <button class="btn btn-primary" @click="openModal()">+ Add Budget</button>
    </div>
    <div class="card">
      <table>
        <thead>
          <tr>
            <th>Department</th>
            <th>Fiscal Year</th>
            <th>Allocated</th>
            <th>Spent</th>
            <th>Remaining</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="b in budgets" :key="b.id">
            <td><strong>{{ b.department }}</strong></td>
            <td>{{ b.fiscal_year }}</td>
            <td>${{ fmt(b.allocated_amount) }}</td>
            <td>${{ fmt(b.spent_amount) }}</td>
            <td :style="{ color: remaining(b) < 0 ? 'var(--danger)' : 'var(--success)' }">
              ${{ fmt(remaining(b)) }}
            </td>
            <td>
              <button class="btn btn-outline" style="margin-right:6px" @click="openModal(b)">Edit</button>
              <button class="btn btn-danger" @click="del(b.id)">Delete</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <div v-if="showModal" class="modal-overlay" @click.self="showModal=false">
      <div class="modal">
        <h3>{{ editing ? 'Edit Budget' : 'New Budget' }}</h3>
        <div class="form-group">
          <label>Department</label>
          <input v-model="form.department" placeholder="e.g. Engineering" />
        </div>
        <div class="form-group">
          <label>Fiscal Year</label>
          <input v-model="form.fiscal_year" type="number" placeholder="2025" />
        </div>
        <div class="form-group">
          <label>Allocated Amount ($)</label>
          <input v-model="form.allocated_amount" type="number" placeholder="0.00" />
        </div>
        <div class="modal-actions">
          <button class="btn btn-outline" @click="showModal=false">Cancel</button>
          <button class="btn btn-primary" @click="save">Save</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import api from '../../api/index.js'

const budgets = ref([])
const showModal = ref(false)
const editing = ref(null)
const form = ref({ department: '', fiscal_year: 2025, allocated_amount: 0 })

const fmt = n => Number(n || 0).toLocaleString()
const remaining = b => b.allocated_amount - b.spent_amount

async function load() {
  const res = await api.get('/budgets')
  budgets.value = res.data || []
}

function openModal(b = null) {
  editing.value = b
  form.value = b ? { ...b } : { department: '', fiscal_year: 2025, allocated_amount: 0 }
  showModal.value = true
}

async function save() {
  if (editing.value) await api.put(`/budgets/${editing.value.id}`, form.value)
  else await api.post('/budgets', form.value)
  showModal.value = false
  load()
}

async function del(id) {
  if (confirm('Delete this budget?')) { await api.delete(`/budgets/${id}`); load() }
}

onMounted(load)
</script>
