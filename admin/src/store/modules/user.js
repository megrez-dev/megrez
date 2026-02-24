import Vue from 'vue'
import { ACCESS_TOKEN, USER_ID } from '@/store/mutation-types'

const user = {
  state: {
    token: null,
    uid: null
  },
  mutations: {
    SET_TOKEN: (state, token) => {
      Vue.ls.set(ACCESS_TOKEN, token)
      state.token = token
    },
    CLEAR_TOKEN: state => {
      Vue.ls.remove(ACCESS_TOKEN)
      Vue.ls.remove(USER_ID)
      state.token = null
      state.uid = null
    },
    SET_USER: (state, uid) => {
      Vue.ls.set(USER_ID, uid)
      state.uid = uid
    }
  },
  actions: {
    login({ commit }, { token, uid }) {
      commit('SET_TOKEN', token)
      commit('SET_USER', uid)
    },
    logout({ commit }) {
      commit('CLEAR_TOKEN')
    },
  }
}

export default user
