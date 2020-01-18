<template>
  <div class="container">
    <div v-for="app in apps" :key="app.ID">
      <AppView :app="app" @onEnter="onEnter" />
    </div>
  </div>
</template>
<script>
import AppView from '~/components/AppView'
const LIMIT = 50
export default {
  components: { AppView },
  data() {
    return {
      page: 1,
      query: ''
    }
  },
  computed: {
    apps() {
      return this.$store.state.apps
    }
  },
  async fetch({ store, params }) {
    await store.dispatch('getApps', {
      offset: 0,
      limit: LIMIT,
      query: ''
    })
  },
  methods: {
    onEnter(app) {
      this.$router.push({ path: `/apps/${app.ID}/room` })
    }
  }
}
</script>
