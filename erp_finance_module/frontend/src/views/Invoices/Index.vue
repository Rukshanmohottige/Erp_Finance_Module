<template>
  <div>
    <div class="page-header">
      <h2>Invoices</h2>
      <button class="btn btn-primary" @click="openModal()">+ New Invoice</button>
    </div>
    <div class="card">
      <table>
        <thead><tr><th>Invoice #</th><th>Vendor</th><th>Amount</th><th>Due Date</th><th>Status</th><th>Actions</th></tr></thead>
        <tbody>
          <tr v-for="inv in invoices" :key="inv.id">
            <td><strong>{{ inv.invoice_number }}</strong></td>
            <td>{{ inv.vendor_name }}</td>
            <td>${{ fmt(inv.amount) }}</td>
            <td>{{ inv.due_date }}</td>
            <td><span :class="['badge', inv.status === 'paid' ? 'badge-success' : 'badge-warning']">{{ inv.status }}</span></td>
            <td>
              <button class="btn btn-outline" style="margin-right:6px" @click="openModal(inv)">Edit</button>
              <button class="btn btn-danger" @click="del(inv.id)">Delete</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <div v-if="showModal" class="modal-overlay" @click.self="showModal=false">
      <div class="modal">
        <h3>{{ editing ? 'Edit Invoice' : 'New Invoice' }}</h3>
        <div class="form-group"><label>Invoice Number</label><input v-model="form.invoice_number" placeholder="INV-2025-001" /></div>
        <div class="form-group"><label>Vendor Name</label><input v-model="form.vendor_name" /></div>
        <div class="form-group"><label>Amount ($)</label><input v-model="form.amount" type="number" /></div>
        <div class="form-group"><label>Due Date</label><input v-model="form.due_date" type="date" /></div>
        <div class="form-group">
          <label>Status</label>
          <select v-model="form.status">
            <option value="unpaid">Unpaid</option>
            <option value="paid">Paid</option>
            <option value="overdue">Overdue</option>
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

const invoices = ref([])
const showModal = ref(false)
const editing = ref(null)
const form = ref({ invoice_number: '', vendor_name: '', amount: 0, due_date: '', status: 'unpaid' })

const fmt = n => Number(n || 0).toLocaleString()

async function load() { invoices.value = (await api.get('/invoices')).data || [] }
function openModal(inv = null) {
  editing.value = inv
  form.value = inv ? { ...inv } : { invoice_number: '', vendor_name: '', amount: 0, due_date: '', status: 'unpaid' }
  showModal.value = true
}
async function save() {
  if (editing.value) await api.put(`/invoices/${editing.value.id}`, form.value)
  else await api.post('/invoices', form.value)
  showModal.value = false; load()
}
async function del(id) { if (confirm('Delete?')) { await api.delete(`/invoices/${id}`); load() } }

onMounted(load)
</script>
