<template>
  <div class="container">
    <div class="field has-addons">
      <div class="control">
        <input
          v-model="query"
          class="input"
          type="text"
          placeholder="Find an application"
        />
      </div>
      <div class="control">
        <a @click="searchApps" class="button is-info">
          Search
        </a>
      </div>
    </div>
    <div>
      <AppView
        v-for="app in apps"
        :key="app.ID"
        :app="app"
        @onEnter="onEnter"
      />
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
    },
    searchApps() {
      this.page = 1
      this.$store.dispatch('getApps', {
        offset: (this.page - 1) * LIMIT,
        limit: LIMIT,
        query: this.query
      })
    }
  }
}
</script>
