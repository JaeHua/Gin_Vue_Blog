import Vue from 'vue'
import App from './App.vue'
import router from './router'
import './plugin/antui'
import axiosInstance from './plugin/axios'
import './assets/css/style.css'

Vue.prototype.$http = axiosInstance
Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
