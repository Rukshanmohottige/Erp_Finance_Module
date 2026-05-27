<template>
  <div>
    <div class="page-header">
      <h2>Payroll</h2>
      <button class="btn btn-primary" @click="openModal()">+ Add Payroll</button>
    </div>
    <div class="card">
      <table>
        <thead><tr><th>Employee ID</th><th>Name</th><th>Department</th><th>Basic</th><th>Net Salary</th><th>Period</th><th>Status</th><th>Actions</th></tr></thead>
        <tbody>
          <tr v-for="p in payroll" :key="p.id">
            <td>{{ p.employee_id }}</td>
            <td><strong>{{ p.employee_name }}</strong></td>
            <td>{{ p.department }}</td>
            <td>${{ fmt(p.basic_salary) }}</td>
            <td><strong>${{ fmt(p.net_salary) }}</strong></td>
            <td>{{ p.pay_period }}</td>
            <td><span :class="['badge', p.status === 'processed' ? 'badge-success' : 'badge-warning']">{{ p.status }}</span></td>
            <td>
              <button class="btn btn-outline" style="margin-right:6px" @click="openModal(p)">Edit</button>
              <button class="btn btn-danger" @click="del(p.id)">Delete</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <div v-if="showModal" class="modal-overlay" @click.self="showModal=false">
      <div class="modal">
        <h3>{{ editing ? 'Edit Payroll' : 'New Payroll Entry' }}</h3>
        <div class="form-group"><label>Employee ID</label><input v-model="form.employee_id" /></div>
        <div class="form-group"><label>Employee Name</label><input v-model="form.employee_name" /></div>
        <div class="form-group"><label>Department</label><input v-model="form.department" /></div>
        <div class="form-group"><label>Basic Salary ($)</label><input v-model="form.basic_salary" type="number" /></div>
        <div class="form-group"><label>Allowances ($)</label><input v-model="form.allowances" type="number" /></div>
        <div class="form-group"><label>Deductions ($)</label><input v-model="form.deductions" type="number" /></div>
        <div class="form-group"><label>Pay Period (e.g. 2025-05)</label><input v-model="form.pay_period" /></div>
        <div class="form-group">
          <label>Status</label>
          <select v-model="form.status"><option value="pending">Pending</option><option value="processed">Processed</option></select>
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

const payroll = ref([])
const showModal = ref(false)
const editing = ref(null)
const form = ref({ employee_id: '', employee_name: '', department: '', basic_salary: 0, allowances: 0, deductions: 0, pay_period: '', status: 'pending' })

const fmt = n => Number(n || 0).toLocaleString()

async function load() { payroll.value = (await api.get('/payroll')).data || [] }
function openModal(p = null) {
  editing.value = p
  form.value = p ? { ...p } : { employee_id: '', employee_name: '', department: '', basic_salary: 0, allowances: 0, deductions: 0, pay_period: '', status: 'pending' }
  showModal.value = true
}
async function save() {
  if (editing.value) await api.put(`/payroll/${editing.value.id}`, form.value)
  else await api.post('/payroll', form.value)
  showModal.value = false; load()
}
async function del(id) { if (confirm('Delete?')) { await api.delete(`/payroll/${id}`); load() } }

onMounted(load)
</script>
