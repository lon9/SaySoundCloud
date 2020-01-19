<template>
  <div class="container">
    <ErrorView :message="errorMsg" />
    <div v-if="app" class="content">
      <p class="title is-4">{{ app.name }}</p>
      <p class="subtitle is-6">@{{ app.user.name }}</p>
      <p class="is-6" style="white-space:pre-line;">
        {{ app.description }}
      </p>
      <p v-if="app.isPassword" class="is-6">password &#10003;</p>
      <p v-else class="is-6">password &#10005;</p>
      <nuxt-link
        :to="localePath({ name: 'apps-id-room', params: { id: app.ID } })"
        class="button"
      >
        Enter
      </nuxt-link>
      <nuxt-link
        v-if="user && user.ID == app.userId"
        :to="localePath({ name: 'apps-id-edit', params: { id: app.ID } })"
        class="button"
      >
        Edit
      </nuxt-link>
      <a @click="deleteApp" class="button">
        Delete
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
        this.$router.push({ path: '/' })
      } else {
        this.errorMsg = 'Failed to delete'
      }
    }
  }
}
</script>
