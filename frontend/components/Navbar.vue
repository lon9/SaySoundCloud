<template>
  <nav class="navbar" role="navigation" aria-label="main navigation">
    <div class="navbar-brand">
      <nuxt-link to="/" class="navbar-item">
        Soundboard
      </nuxt-link>
    </div>
    <client-only>
      <div class="navbar-end">
        <div v-if="user" class="navbar-item">
          {{ user.name }}
        </div>
        <div v-if="user" class="buttons">
          <nuxt-link to="/apps/create" class="button">
            Create app
          </nuxt-link>
          <nuxt-link to="/users/edit" class="button">
            {{ $t('editProfile') }}
          </nuxt-link>
          <div @click="signOut" class="button">
            {{ $t('signOut') }}
          </div>
        </div>
        <div v-else>
          <nuxt-link to="/signin" class="button is-primary">
            <strong>{{ $t('signIn') }}</strong>
          </nuxt-link>
        </div>
      </div>
    </client-only>
  </nav>
</template>
<script>
export default {
  computed: {
    user() {
      return this.$store.state.user
    }
  },
  methods: {
    async signOut() {
      await this.$store.dispatch('signOut')
      this.$router.push({ path: '/' })
    }
  }
}
</script>
