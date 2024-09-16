// 本地缓存服务
const PREFIX = 'Vue_'

// user模块
const USER_PREFIX = `${PREFIX}user_`
const USER_TOKEN = `${USER_PREFIX}token`
const USER_INFO = `${USER_PREFIX}info`

// 存储
const set = (key, data) => {
  sessionStorage.setItem(key, data)
}

// 读取
const get = (key) => sessionStorage.getItem(key)

export default {
  set,
  get,
  USER_INFO,
  USER_TOKEN
}
