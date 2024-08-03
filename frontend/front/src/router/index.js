import Vue from 'vue'
import VueRouter from 'vue-router'
import HomeView from '../views/Home.vue'
import ArtList from '../components/ArtList.vue'
import Detail from '../components/Details.vue'
import Category from '../components/Category.vue'
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
  }, {
    path: '/detail/:id',
    component: Detail,
    meta: { title: '文章详情' },
    props: true
  }, {
    path: '/category/:id',
    component: Category,
    meta: { title: '文章分类页' },
    props: true
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
