import Vue from 'vue'
import axios from 'axios'

// axios.defaults.baseURL = 'http://localhost:3000/api/v1'
const URL = 'http://localhost:3000/api/v1'
const instance = axios.create({
  baseURL: URL, // 替换为你的 API 基础 URL
  timeout: 100000
})
// 添加请求拦截器
instance.interceptors.request.use(
  (config) => {
    const token = window.sessionStorage.getItem('token')
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)
Vue.prototype.$http = instance

// Vue.prototype.$http = axios
