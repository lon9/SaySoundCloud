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
      try {
        const res = await this.$axios.$post('/apps', this.app)
        if (res.status === 201) {
          this.$router.push({ path: `/apps/${res.result.ID}/edit` })
        } else {
          this.errorMsg = 'Failed to create'
        }
      } catch {}
    }
  }
}
</script>
