import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  { path: '/login', component: () => import('../views/Login.vue'), meta: { public: true } },
  {
    path: '/',
    component: () => import('../components/Layout.vue'),
    meta: { requiresAuth: true },
    children: [
      { path: '', redirect: '/dashboard' },
      { path: 'dashboard', component: () => import('../views/Dashboard/Index.vue') },
      { path: 'budgets', component: () => import('../views/Budget/Index.vue') },
      { path: 'expenses', component: () => import('../views/Expenses/Index.vue') },
      { path: 'invoices', component: () => import('../views/Invoices/Index.vue') },
      { path: 'payroll', component: () => import('../views/Payroll/Index.vue') },
      { path: 'reports', component: () => import('../views/Reports/Index.vue') },
    ]
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('token')
  if (to.meta.requiresAuth && !token) return next('/login')
  if (to.path === '/login' && token) return next('/dashboard')
  next()
})

export default router
