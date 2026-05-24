<template>
  <div class="login-page">
    <div class="login-card">
      <div class="login-header">
        <div class="login-logo">💰</div>
        <h1>ERP Finance</h1>
        <p>Sign in to your account</p>
      </div>
      <div class="form-group">
        <label>Email</label>
        <input v-model="email" type="email" placeholder="admin@erp.com" />
      </div>
      <div class="form-group">
        <label>Password</label>
        <input v-model="password" type="password" placeholder="••••••••" @keyup.enter="login" />
      </div>
      <button class="btn btn-primary login-btn" @click="login" :disabled="loading">
        {{ loading ? 'Signing in...' : 'Sign In' }}
      </button>
      <p v-if="error" class="error-msg">{{ error }}</p>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../store/auth.js'

const auth = useAuthStore()
const router = useRouter()
const email = ref('')
const password = ref('')
const loading = ref(false)
const error = ref('')

async function login() {
  loading.value = true
  error.value = ''
  try {
    await auth.login(email.value, password.value)
    router.push('/dashboard')
  } catch {
    error.value = 'Invalid email or password'
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.login-page {
  min-height: 100vh; background: var(--sidebar-bg);
  display: flex; align-items: center; justify-content: center;
}
.login-card {
  background: white; border-radius: 16px; padding: 40px; width: 400px;
}
.login-header { text-align: center; margin-bottom: 28px; }
.login-logo { font-size: 40px; margin-bottom: 10px; }
.login-header h1 { font-size: 24px; font-weight: 700; }
.login-header p { color: var(--text-muted); font-size: 14px; margin-top: 4px; }
.login-btn { width: 100%; padding: 12px; font-size: 15px; margin-top: 4px; }
.error-msg { color: var(--danger); font-size: 13px; text-align: center; margin-top: 12px; }
</style>
