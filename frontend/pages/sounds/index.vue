<template>
  <div class="container">
    <p>{{ $t('canCopyDesc') }} {{ $t('numberOfSounds') }}: {{ soundCount }}</p>
    <SearchBox @search="searchSounds" :place-holder="$t('findSound')" />
    <Pagination @next="next" @prev="prev" :page="page" />
    <div class="panel">
      <a
        v-for="(sound, index) in sounds"
        :key="index"
        @click="copyToClipboard(sound)"
        class="panel-block"
      >
        {{ sound.name }}
      </a>
    </div>
    <Pagination @prev="prev" @next="next" :page="page" />
  </div>
</template>
<script>
import Pagination from '~/components/Pagination'
import SearchBox from '~/components/SearchBox'
const LIMIT = 100
export default {
  components: { Pagination, SearchBox },
  data() {
    return {
      page: 1,
      query: '',
      soundCount: 0
    }
  },
  computed: {
    sounds() {
      return this.$store.state.sounds
    }
  },
  async asyncData({ $axios }) {
    try {
      const res = await $axios.$get('/sounds/count')
      return {
        soundCount: res.result
      }
    } catch {}
  },
  async fetch({ store, params }) {
    await store.dispatch('getSounds', {
      offset: 0,
      limit: LIMIT,
      query: ''
    })
  },
  methods: {
    searchSounds(query) {
      this.query = query
      this.page = 1
      this.$store.dispatch('getSounds', {
        offset: (this.page - 1) * LIMIT,
        limit: LIMIT,
        query: this.query
      })
    },
    copyToClipboard(sound) {
      navigator.clipboard.writeText(sound.name)
    },
    next() {
      this.page++
      this.$store.dispatch('getSounds', {
        offset: (this.page - 1) * LIMIT,
        limit: LIMIT,
        query: this.query
      })
    },
    prev() {
      this.page--
      this.$store.dispatch('getSounds', {
        offset: (this.page - 1) * LIMIT,
        limit: LIMIT,
        query: this.query
      })
    }
  }
}
</script>
