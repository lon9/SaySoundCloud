export const state = () => ({
  token: '',
  user: null,
  apps: [],
  websocketConnection: null
})

export const mutations = {
  setToken(state, token) {
    state.token = token
    localStorage.setItem('token', token)
  },
  setUser(state, user) {
    state.user = user
  },
  setApps(state, apps) {
    state.apps = apps
  },
  setConnection(state, connection) {
    state.websocketConnection = connection
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
  async getApps({ commit }, { offset, limit, query }) {
    try {
      const res = await this.$axios.$get('/apps', {
        params: {
          offset,
          limit,
          q: query
        }
      })
      commit('setApps', res.result)
      return res.result
    } catch {}
  },
  async deleteApp(_, id) {
    try {
      await this.$axios.$delete(`/apps/${id}`)
      return true
    } catch {
      return false
    }
  },
  connectWebsocket({ commit }, { id, accessToken }) {
    const baseUrl = process.env.BASE_URL.replace(/http/, 'ws')
    const connection = new WebSocket(
      `${baseUrl}/apps/${id}/ws?token=${accessToken}`
    )
    commit('setConnection', connection)
  },
  closeConnection({ state, commit }) {
    if (state.websocketConnection !== null) {
      state.websocketConnection.close()
      commit('setConnection', null)
    }
  }
}
