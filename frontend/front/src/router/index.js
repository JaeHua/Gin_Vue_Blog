import Vue from 'vue'
import VueRouter from 'vue-router'
import HomeView from '../views/Home.vue'
import ArtList from '../components/ArtList.vue'
import Detail from '../components/Details.vue'
import Category from '../components/Category.vue'
import Result from '../components/Result.vue'
import Account from '@/components/Account.vue'
import NotFound from '@/components/404.vue'
import * as jwtdecode from 'jwt-decode'
import store from '@/store'
Vue.use(VueRouter)
const routes =
[{
  path: '/',
  component: HomeView,
  children: [{
    path: '/',
    name: 'home',
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
  }, {
    path: '/result',
    component: Result,
    meta: { title: '文章搜索页' }
  }]

}, {
  path: '/account',
  name: 'account',
  component: Account,
  meta: { title: '用户空间', auth: true }
},
{
  path: '/404',
  name: '404',
  component: NotFound

}]

const router = new VueRouter({ mode: 'history', base: process.env.BASE_URL, routes })
router.beforeEach((to, from, next) => {
  if (to.meta.title) {
    document.title = to.meta.title
  }

  if (to.matched.length === 0) {
    next({ name: '404' })
  } else {
    if (to.meta.auth) {
      const token = store.state.userModule.token

      if (token) {
        try {
          const decoded = jwtdecode.jwtDecode(token)
          const currentTime = Date.now() / 1000

          if (decoded.exp < currentTime) {
            store.dispatch('userModule/logout').then(() => {
              document.title = '欢迎来到我的博客'

              next({ name: 'home' })
            })
          } else {
            next()
          }
        } catch (e) {
          console.error('Token decoding error:', e)
          store.dispatch('userModule/logout').then(() => {
            document.title = '欢迎来到我的博客'

            next({ name: 'home' })
          })
        }
      } else {
        document.title = '欢迎来到我的博客'
        next({ name: 'home' })
      }
    } else {
      next()
    }
  }
})
export default router
