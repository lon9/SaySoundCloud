<template>
  <div class="container">
    <ErrorView :message="errorMsg" />
    <div v-if="app" class="content">
      <p class="title is-4">{{ app.name }}</p>
      <p class="subtitle is-6">
        <nuxt-link
          :to="localePath({ name: 'users-id', params: { id: app.userId } })"
        >
          @{{ app.user.name }}
        </nuxt-link>
      </p>
      <div v-html="$md.render(app.description)" />
      <p v-if="app.isPassword" class="is-6">{{ $t('password') }} &#10003;</p>
      <p v-else class="is-6">{{ $t('password') }} &#10005;</p>
      <nuxt-link
        :to="localePath({ name: 'apps-id-room', params: { id: app.ID } })"
        class="button"
      >
        {{ $t('enter') }}
      </nuxt-link>
      <nuxt-link
        v-if="user && user.ID == app.userId"
        :to="localePath({ name: 'apps-id-edit', params: { id: app.ID } })"
        class="button"
      >
        {{ $t('edit') }}
      </nuxt-link>
      <a
        v-if="user && user.ID == app.userId"
        @click="deleteApp"
        class="button is-danger"
      >
        {{ $t('delete') }}
      </a>
    </div>
  </div>
</template>
<script>
import ErrorView from '~/components/ErrorView'
export default {
  components: { ErrorView },
  data() {
    return {
      errorMsg: ''
    }
  },
  computed: {
    user() {
      return this.$store.state.user
    }
  },
  async asyncData({ $axios, params }) {
    try {
      const res = await $axios.$get(`/apps/${params.id}`)
      return {
        app: res.result
      }
    } catch {}
  },
  methods: {
    async deleteApp() {
      this.errorMsg = ''
      if (await this.$store.dispatch('deleteApp', this.app.ID)) {
        this.$router.push(this.localePath({ path: '/' }))
      } else {
        this.errorMsg = this.$t('failedToDelete')
      }
    }
  }
}
</script>
