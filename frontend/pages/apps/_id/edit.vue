<template>
  <div v-if="app" class="container">
    <ErrorView :message="errorMsg" />
    <AppForm :app="app" :onSubmit="onAppSubmit" />
    <p>Access token: {{ accessToken }}</p>
  </div>
</template>
<script>
import AppForm from '~/components/AppForm'
import ErrorView from '~/components/ErrorView'
export default {
  components: { AppForm, ErrorView },
  data() {
    return {
      errorMsg: '',
      app: {
        name: '',
        description: ''
      },
      accessToken: ''
    }
  },
  async mounted() {
    try {
      const res = await this.$axios.$get(`/apps/${this.$route.params.id}/owner`)
      this.app.name = res.result.name
      this.app.description = res.result.description
      this.accessToken = res.result.accessToken
    } catch (e) {}
  },
  methods: {
    async onAppSubmit() {
      this.errorMsg = ''
      try {
        const res = await this.$axios.$put(
          `/apps/${this.$route.params.id}`,
          this.app
        )
        this.$router.push({ path: `/apps/${res.result.ID}/edit` })
      } catch {
        this.errorMsg = 'Failed to update'
      }
    }
  }
}
</script>
