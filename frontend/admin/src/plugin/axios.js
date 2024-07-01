// src/axios.js
import axios from 'axios'

// 创建一个 axios 实例
const instance = axios.create({
  baseURL: 'http://localhost:3000/api/v1', // 替换为你的 API 基础 URL
  timeout: 10000
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

export default instance
