import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    loggedIn: false,
    ID_email: ''
  },
  mutations: {
    setLoginState (state, status) {
      state.loggedIn = status
    },
    setProfile (state, profile) {
      state.ID_email = profile || ''
    },
    clearProfile (state) {
      state.ID_email = ''
    }
  },
  actions: {
    _login ({ commit }, profile) {
      commit('setLoginState', true)
      if (profile) {
        commit('setProfile', profile)
      }
    },
    _logout ({ commit }) {
      commit('setLoginState', false)
      commit('clearProfile')
    }
  },
  getters: {
    isLoggedIn: (state) => state.loggedIn,
    getProfile: (state) => state.ID_email
  }
})
export default store
