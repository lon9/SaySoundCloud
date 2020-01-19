<template>
  <nav class="navbar is-dark" role="navigation" aria-label="main navigation">
    <div class="navbar-brand">
      <nuxt-link to="/" class="navbar-item">
        {{ appTitle }}
      </nuxt-link>
    </div>
    <client-only>
      <div class="navbar-end">
        <nuxt-link
          :to="localePath({ name: 'users-id', params: { id: user.ID } })"
          v-if="user"
          class="navbar-item has-text-white"
        >
          {{ user.name }}
        </nuxt-link>
        <div class="navbar-item">
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
      </div>
    </client-only>
  </nav>
</template>
<script>
export default {
  data() {
    return {
      appTitle: process.env.APP_TITLE
    }
  },
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
