<template>
  <div class="container">
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
    </div>
  </div>
</template>
<script>
export default {
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
  }
}
</script>
