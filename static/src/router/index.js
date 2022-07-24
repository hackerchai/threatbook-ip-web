import Vue from 'vue'
import VueRouter from 'vue-router'
import IocTable from '../views/IocTable'
import Index from '../views/Index'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'index',
    component: Index
  },
  {
    path: '/IocTable',
    name: 'iocTable',
    component: IocTable
  }
]

const router = new VueRouter({
  routes
})

export default router
