<template>
  <div class="container">
    App room
    <PasswordModal
      :app="app"
      @onSubmit="onModalSubmit"
      :onClose="onModalClose"
      :isActive="isModalActive"
    />
  </div>
</template>
<script>
import PasswordModal from '~/components/PasswordModal'
export default {
  components: { PasswordModal },
  data() {
    return {
      accessToken: '',
      connection: null
    }
  },
  async asyncData({ $axios, params }) {
    try {
      let isModalActive = false
      const res = await $axios.$get(`/apps/${params.id}`)
      if (res.result.isPassword) {
        isModalActive = true
      }
      return {
        app: res.result,
        isModalActive
      }
    } catch {}
  },
  mounted() {
    if (!this.app.isPassword) {
      this.connectToWebsocket()
    }
  },
  methods: {
    async onModalSubmit(password) {
      try {
        const res = await this.$axios.$post(`/apps/${this.app.ID}/ws`, {
          password
        })
        this.accessToken = res.result
        this.isModalActive = false
        this.connectToWebsocket()
      } catch {}
    },
    onModalClose() {
      this.isModalActive = false
    },
    connectToWebsocket() {
      const baseUrl = process.env.BASE_URL.replace(/http/, 'ws')
      this.connection = new WebSocket(
        `${baseUrl}/apps/${this.app.ID}/ws?token=${this.accessToken}`
      )
      this.connection.onopen = function() {
        console.log('open')
      }
      this.connection.onmessage = function(e) {
        console.log(e.data)
      }
      this.connection.onclose = function() {
        console.log('close')
      }
      this.connection.onerror = function(err) {
        console.error(err)
      }
    }
  }
}
</script>
