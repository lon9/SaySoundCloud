<template>
  <div class="container">
    <UserForm :profile="profile" :on-submit="onUserSubmit" />
    <div class="container">
      App place holder
    </div>
  </div>
</template>
<script>
import UserForm from '~/components/UserForm'
export default {
  components: { UserForm },
  data() {
    return {
      profile: {
        name: '',
        description: ''
      }
    }
  },
  mounted() {
    if (this.$store.state.user) {
      this.profile.name = this.$store.state.user.name
      this.profile.description = this.$store.state.user.description
    }
  },
  methods: {
    onUserSubmit() {
      this.$store.dispatch('updateUser', this.profile)
    }
  }
}
</script>
