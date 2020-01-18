<template>
  <div class="container">
    <ErrorView :message="errorMsg" />
    <UserForm :profile="profile" :on-submit="onUserSubmit" />
    <div class="container">
      App place holder
    </div>
  </div>
</template>
<script>
import UserForm from '~/components/UserForm'
import ErrorView from '~/components/ErrorView'
export default {
  components: { UserForm, ErrorView },
  data() {
    return {
      profile: {
        name: '',
        description: ''
      },
      errorMsg: ''
    }
  },
  mounted() {
    if (this.$store.state.user) {
      this.profile.name = this.$store.state.user.name
      this.profile.description = this.$store.state.user.description
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
