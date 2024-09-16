import Vue from 'vue'
import Toast from 'vue-toastification'
import 'vue-toastification/dist/index.css'

Vue.use(Toast)

export function useToast () {
  return Vue.prototype.$toast
}
