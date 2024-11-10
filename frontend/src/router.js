import { createRouter, createWebHistory } from 'vue-router'

import MainView from './views/MainView.vue'
import CategoryView from './views/CategoryView.vue'
import ProductView from './views/ProductView.vue'
import CartView from './views/CartView.vue'
import LoginView from './views/Login.vue'

export const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/api/category/:id', component: CategoryView },
    { path: '/', component: MainView },
    { path: '/api/product/:id', component: ProductView },
    { path: '/cart', component: CartView },
    { path: '/login', component: LoginView }
    ],
})
