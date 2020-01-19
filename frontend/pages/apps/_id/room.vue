<template>
  <div class="container">
    <div v-if="app" class="content">
      <p class="title is-4">{{ app.name }}</p>
      <p class="subtitle is-6">@{{ app.user.name }}</p>
      <p class="is-6" style="white-space:pre-line;">
        {{ app.description }}
      </p>
    </div>
    You can copy the command by clicking the command name.
    <div v-if="cmds.length !== 0" class="list">
      <a
        v-for="(cmd, index) in cmds"
        :key="index"
        @click="copyToClipboard(cmd)"
        class="list-item"
      >
        {{ cmd.name }} <small>({{ cmd.time }})</small>
      </a>
    </div>
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
      cmds: [],
      audioCtx: null
    }
  },
  computed: {
    connection() {
      return this.$store.state.websocketConnection
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
      this.connect()
    }
    try {
      window.AudioContext = window.AudioContext || window.webkitAudioContext
      this.audioCtx = new AudioContext()
    } catch {
      this.$router.push({ path: '/' })
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
        this.connect()
      } catch {}
    },
    onModalClose() {
      this.isModalActive = false
    },
    copyToClipboard(cmd) {
      navigator.clipboard.writeText(cmd.name)
    },
    connect() {
      const that = this
      this.$store.dispatch('connectWebsocket', {
        id: this.app.ID,
        accessToken: this.accessToken
      })
      this.connection.onopen = function() {
        console.log('open')
      }
      this.connection.onclose = function() {
        console.log('close')
      }
      this.connection.onerror = function(e) {
        console.error(e)
      }
      this.connection.onmessage = async function(e) {
        const data = JSON.parse(e.data.toString())
        if (data.event === 'cmd') {
          const cmdName = decodeURIComponent(atob(data.payload))
          try {
            const source = await that.$axios.$get(
              `${process.env.SOUND_BASE_URL}/${cmdName}.wav`,
              {
                responseType: 'arraybuffer'
              }
            )
            that.audioCtx.decodeAudioData(source, function(buffer) {
              const source = that.audioCtx.createBufferSource()
              source.buffer = buffer
              source.connect(that.audioCtx.destination)
              source.start(0)
              const time = new Date()
              that.cmds.unshift({
                name: cmdName,
                time: `${time.getHours()}:${time.getMinutes()}:${time.getSeconds()}`
              })
            })
          } catch {}
        }
      }
    }
  }
}
</script>
