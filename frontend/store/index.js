export const state = () => ({
  token: '',
  user: null
})

export const mutations = {
  setToken(state, token) {
    state.token = token
    localStorage.setItem('token', token)
  },
  setUser(state, user) {
    state.user = user
  }
}

export const actions = {
  async nuxtClientInit({ dispatch, commit }) {
    const token = localStorage.getItem('token')
    if (token) {
      commit('setToken', token)
      await dispatch('getLoginUser')
    }
  },
  signOut({ commit }) {
    commit('setToken', '')
    commit('setUser', null)
  },
  async createUser({ commit }) {
    try {
      const res = await this.$axios.$post('/users')
      commit('setUser', res.result)
      return true
    } catch {
      return false
    }
  },
  async updateUser({ state, commit }, profile) {
    try {
      const res = await this.$axios.$put(`/users/${state.user.ID}`, profile)
      if (res.status === 200) commit('setUser', res.result)
    } catch {}
  },
  async getLoginUser({ dispatch, commit }) {
    try {
      const res = await this.$axios.$get(`/users/me`)
      if (res.status === 200) commit('setUser', res.result)
    } catch {
      dispatch('signOut')
    }
  }
}
