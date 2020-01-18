export const state = () => ({
  token: '',
  user: null,
  apps: []
})

export const mutations = {
  setToken(state, token) {
    state.token = token
    localStorage.setItem('token', token)
  },
  setUser(state, user) {
    state.user = user
  },
  unshiftApp(state, app) {
    state.apps.unshift(app)
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
      return res.result
    } catch {}
  },
  async updateUser({ state, commit }, profile) {
    try {
      const res = await this.$axios.$put(`/users/${state.user.ID}`, profile)
      commit('setUser', res.result)
      return res.result
    } catch {}
  },
  async getLoginUser({ dispatch, commit }) {
    try {
      const res = await this.$axios.$get(`/users/me`)
      commit('setUser', res.result)
      return res.result
    } catch {
      dispatch('signOut')
    }
  },
  async createApp({ commit }, app) {
    try {
      const res = await this.$axios.$post('/apps', app)
      commit('unshiftApp', res.result)
      return res.result
    } catch {}
  }
}
