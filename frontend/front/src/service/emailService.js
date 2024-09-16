import http from '@/plugins/http'
import { useToast } from '@/plugins/toast'
const toast = useToast()
// 发送验证码模块
const sendVerifyCode = async (mail) => {
  try {
    const response = await http.post('register/getcode', mail)
    // 请求成功
    toast.success('验证码发送成功', { timeout: 3000 })
    return response
  } catch (error) {
    toast.error('验证码发送失败', { timeout: 3000 })
    return await Promise.reject(error)
  }
}
// 验证验证码模块
const VerifyCode = async (data) => {
  try {
    const response = await http.post('register/verify', data)
    // 请求成功
    if (response.data.status !== 200) {
      toast.error(response.data.message, { timeout: 3000 })
      return false
    } else {
      // toast.success(response.data.message, { timeout: 3000 })
      return true
    }
    // return response
  } catch (error) {
    toast.error('验证服务错误', { timeout: 3000 })
    return await Promise.reject(error)
  }
}

export default {
  sendVerifyCode,
  VerifyCode
}
