<template>
  <div class="container">
    <ErrorView :message="errorMsg" />
    <AppForm :app="app" :onSubmit="onAppSubmit" />
  </div>
</template>
<script>
import AppForm from '~/components/AppForm'
import ErrorView from '~/components/ErrorView'
export default {
  components: { AppForm, ErrorView },
  data() {
    return {
      app: {
        name: '',
        description: '',
        password: ''
      },
      errorMsg: ''
    }
  },
  methods: {
    async onAppSubmit() {
      this.errorMsg = ''
      const app = await this.$store.dispatch('createApp', this.app)
      if (app) {
        this.$router.push({ path: `/apps/${app.ID}/edit` })
      } else {
        this.errorMsg = 'Failed to create'
      }
    }
  }
}
</script>
