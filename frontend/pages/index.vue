<template>
  <div class="container">
    <SearchBox @search="searchApps" place-holder="Find an application" />
    <div>
      <AppView v-for="app in apps" :key="app.ID" :app="app" />
    </div>
    <Pagination @next="next" @prev="prev" :page="page" />
  </div>
</template>
<script>
import AppView from '~/components/AppView'
import Pagination from '~/components/Pagination'
import SearchBox from '~/components/SearchBox'
const LIMIT = 50
export default {
  components: { AppView, Pagination, SearchBox },
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
    searchApps(query) {
      this.page = 1
      this.query = query
      this.$store.dispatch('getApps', {
        offset: (this.page - 1) * LIMIT,
        limit: LIMIT,
        query: this.query
      })
    },
    next() {
      this.page++
      this.$store.dispatch('getApps', {
        offset: (this.page - 1) * LIMIT,
        limit: LIMIT,
        query: this.query
      })
    },
    prev() {
      this.page--
      this.$store.dispatch('getApps', {
        offset: (this.page - 1) * LIMIT,
        limit: LIMIT,
        query: this.query
      })
    }
  }
}
</script>
