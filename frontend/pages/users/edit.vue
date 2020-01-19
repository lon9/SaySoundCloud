<template>
  <div class="container">
    <ErrorView :message="errorMsg" />
    <UserForm :profile="profile" :on-submit="onUserSubmit" />
    <p class="is-size-4">Your apps</p>
    <div class="container">
      <AppView v-for="app in apps" :key="app.ID" :app="app" />
    </div>
  </div>
</template>
<script>
import AppView from '~/components/AppView'
import UserForm from '~/components/UserForm'
import ErrorView from '~/components/ErrorView'
export default {
  components: { UserForm, ErrorView, AppView },
  data() {
    return {
      profile: {
        name: '',
        description: ''
      },
      errorMsg: '',
      apps: []
    }
  },
  async mounted() {
    if (this.$store.state.user) {
      this.profile.name = this.$store.state.user.name
      this.profile.description = this.$store.state.user.description
      try {
        const res = await this.$axios.$get(`/apps`, {
          params: {
            userId: this.$store.state.user.ID
          }
        })
        this.apps = res.result
      } catch {}
    }
  },
  methods: {
    async onUserSubmit() {
      this.errMsg = ''
      const user = await this.$store.dispatch('updateUser', this.profile)
      if (user) {
        this.$router.push({ path: `/users/${this.$store.state.user.ID}` })
      } else {
        this.errorMsg = 'Failed to edit'
      }
    }
  }
}
</script>
