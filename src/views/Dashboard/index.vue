<template>
  <div>
    <div class="page-header">
      <h2>Expenses</h2>
      <button class="btn btn-primary" @click="openModal()">+ Add Expense</button>
    </div>
    <div class="card">
      <table>
        <thead><tr><th>Description</th><th>Category</th><th>Amount</th><th>Status</th><th>Actions</th></tr></thead>
        <tbody>
          <tr v-for="e in expenses" :key="e.id">
            <td>{{ e.description }}</td>
            <td>{{ e.category }}</td>
            <td>${{ fmt(e.amount) }}</td>
            <td><span :class="['badge', badgeClass(e.status)]">{{ e.status }}</span></td>
            <td>
              <button class="btn btn-outline" style="margin-right:6px" @click="openModal(e)">Edit</button>
              <button class="btn btn-danger" @click="del(e.id)">Delete</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <div v-if="showModal" class="modal-overlay" @click.self="showModal=false">
      <div class="modal">
        <h3>{{ editing ? 'Edit Expense' : 'New Expense' }}</h3>
        <div class="form-group"><label>Description</label><input v-model="form.description" /></div>
        <div class="form-group"><label>Amount ($)</label><input v-model="form.amount" type="number" /></div>
        <div class="form-group"><label>Category</label><input v-model="form.category" /></div>
        <div class="form-group">
          <label>Status</label>
          <select v-model="form.status">
            <option value="pending">Pending</option>
            <option value="approved">Approved</option>
            <option value="rejected">Rejected</option>
          </select>
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

const expenses = ref([])
const showModal = ref(false)
const editing = ref(null)
const form = ref({ description: '', amount: 0, category: '', status: 'pending', budget_id: 1, submitted_by: 1 })

const fmt = n => Number(n || 0).toLocaleString()
const badgeClass = s => ({ approved: 'badge-success', pending: 'badge-warning', rejected: 'badge-danger' }[s] || 'badge-warning')

async function load() { expenses.value = (await api.get('/expenses')).data || [] }

function openModal(e = null) {
  editing.value = e
  form.value = e ? { ...e } : { description: '', amount: 0, category: '', status: 'pending', budget_id: 1, submitted_by: 1 }
  showModal.value = true
}

async function save() {
  if (editing.value) await api.put(`/expenses/${editing.value.id}`, form.value)
  else await api.post('/expenses', form.value)
  showModal.value = false
  load()
}

async function del(id) {
  if (confirm('Delete?')) { await api.delete(`/expenses/${id}`); load() }
}

onMounted(load)
</script>
