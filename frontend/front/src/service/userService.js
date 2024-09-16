import http from '@/plugins/http'

// 用户注册
const register = ({ username, email, password }) => {
  return http.post('register', { username, email, password })
}

// 用户登陆
const login = ({ email, password }) => {
  return http.post('userlogin', { email, password })
}

// 获取用户信息
const info = () => {
  return http.get('user/info')
}

export default {
  register,
  login,
  info
}
