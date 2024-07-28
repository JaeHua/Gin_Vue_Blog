import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../views/Login.vue'
import Admin from '../views/Admin.vue'
import Index from '@/components/admin/Index.vue'
import AddArt from '@/components/article/AddArt.vue'
import ArtList from '@/components/article/ArtList.vue'
import CateList from '@/components/category/CateList.vue'
import UserList from '@/components/user/UserList.vue'
import Profile from '@/components/user/Profile.vue'

Vue.use(VueRouter)

// 页面路由组件

const routes = [
  {
    path: '/login',
    name: 'login',
    component: Login
  },
  {
    path: '/',
    name: 'admin',
    component: Admin,
    children: [
      { path: 'index', component: Index },
      { path: 'artlist', component: ArtList },
      { path: 'addart', component: AddArt },
      { path: 'addart/:id', component: AddArt, props: true },
      { path: 'catelist', component: CateList },
      { path: 'userlist', component: UserList },
      { path: 'profile', component: Profile }
    ]
  }
]

const router = new VueRouter({
  routes
})

router.beforeEach((to, from, next) => {
  const token = window.sessionStorage.getItem('token')
  if (to.path === '/login') return next()
  if (!token && to.path === '/') {
    next('/login')
  } else {
    next()
  }
})
export default router
