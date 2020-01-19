<template>
  <div class="container">
    <div v-if="user" class="content">
      <p class="title is-4">{{ user.name }}</p>
      <p class="is-6" style="white-space:pre-line;">{{ user.description }}</p>
    </div>
    <p class="is-size-4">Apps</p>
    <div>
      <AppView v-for="app in apps" :key="app.ID" :app="app" />
    </div>
  </div>
</template>
<script>
import AppView from '~/components/AppView'
export default {
  components: { AppView },
  async asyncData({ $axios, params }) {
    try {
      const userRes = await $axios.$get(`/users/${params.id}`)
      const appsRes = await $axios.$get(`/apps`, {
        params: {
          userId: params.id
        }
      })
      return {
        user: userRes.result,
        apps: appsRes.result
      }
    } catch {}
  }
}
</script>
