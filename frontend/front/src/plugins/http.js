import Vue from 'vue'
import axios from 'axios'
import storageService from '@/service/storageService.js'
// axios.defaults.baseURL = 'http://localhost:3000/api/v1'
const URL = 'http://150.109.93.8:80/api/v1'
const instance = axios.create({
  baseURL: URL, // 替换为你的 API 基础 URL
  timeout: 100000,
  headers: { Authorization: `Bearer ${storageService.get(storageService.USER_TOKEN)}` }
})
// 添加请求拦截器
instance.interceptors.request.use((config) => {
  // Do something before request is sent
  Object.assign(config.headers, { Authorization: `Bearer ${storageService.get(storageService.USER_TOKEN)}` })
  return config
}, (error) => {
  // Do something with request error
  return Promise.reject(error)
})
Vue.prototype.$http = instance

export default instance
// Vue.prototype.$http = axios
