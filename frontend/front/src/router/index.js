import Vue from 'vue'
import VueRouter from 'vue-router'
import HomeView from '../views/Home.vue'
import ArtList from '../components/ArtList.vue'
Vue.use(VueRouter)
const routes =
[{
  path: '/',
  name: 'home',
  component: HomeView,
  children: [{
    path: '/',
    component: ArtList,
    meta: { title: '欢迎来到我的博客' }
  }]
}]

const router = new VueRouter({ mode: 'history', base: process.env.BASE_URL, routes })
router.beforeEach((to, from, next) => {
  if (to.meta.title) {
    document.title = to.meta.title
  }
  next()
})
export default router
