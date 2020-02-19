<template>
  <div class="container">
    <div v-if="app" class="content">
      <p class="title is-4">{{ app.name }}</p>
      <p class="subtitle is-6">@{{ app.user.name }}</p>
      <div v-html="$md.render(app.description)" />
    </div>
    {{ $t('canCopyDesc') }}<br />
    <client-only>
      <vue-slider
        v-model="volume"
        :min="0"
        :max="2"
        :interval="0.01"
        :dragOnClick="true"
        @drag-end="onDragEnd"
      />
    </client-only>
    <div v-if="cmds.length !== 0" class="panel">
      <a
        v-for="(cmd, index) in cmds"
        :key="index"
        @click="copyToClipboard(cmd)"
        class="panel-block"
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
      audioCtx: null,
      gainNode: null,
      volume: 1
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
      this.$router.push(this.localePath({ path: '/' }))
    }
    this.volume = localStorage.getItem('volume')
    if (this.volume === null) {
      this.volume = 1
      localStorage.setItem('volume', this.volume)
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
    onDragEnd() {
      if (this.gainNode !== null) {
        this.gainNode.gain.value = this.volume
      }
      localStorage.setItem('volume', this.volume)
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
          const sound = JSON.parse(decodeURIComponent(atob(data.payload)))
          let url = ''
          if (process.env.USE_FIREBASE === 'true') {
            const ref = that.$storage.ref(
              `${process.env.SOUND_BASE_URL}/${sound.path}`
            )
            url = await ref.getDownloadURL()
          } else {
            url = `${process.env.SOUND_BASE_URL}/${sound.path}`
          }
          try {
            const src = await that.$axios.$get(url, {
              responseType: 'arraybuffer'
            })
            that.audioCtx.decodeAudioData(src, function(buffer) {
              const source = that.audioCtx.createBufferSource()
              source.buffer = buffer
              that.gainNode = that.audioCtx.createGain()
              source.connect(that.gainNode)
              that.gainNode.connect(that.audioCtx.destination)
              that.gainNode.gain.value = that.volume
              source.start(0)
              const time = new Date()
              that.cmds.unshift({
                name: sound.name,
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
