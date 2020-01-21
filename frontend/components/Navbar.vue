<template>
  <nav class="navbar is-dark" role="navigation" aria-label="main navigation">
    <div class="navbar-brand">
      <nuxt-link :to="localePath({ path: '/' })" class="navbar-item">
        {{ appTitle }}
      </nuxt-link>
    </div>
    <div class="navbar-menu is-active">
      <div class="navbar-start">
        <a
          :href="localePath({ path: '/sounds' })"
          target="_blank"
          class="navbar-item"
        >
          {{ $t('commandList') }}
        </a>
      </div>
      <client-only>
        <div class="navbar-end">
          <nuxt-link
            :to="localePath({ name: 'users-id', params: { id: user.ID } })"
            v-if="user"
            class="navbar-item"
          >
            {{ user.name }}
          </nuxt-link>
          <div class="navbar-item">
            <div v-if="user" class="buttons">
              <nuxt-link
                :to="localePath({ name: 'apps-create' })"
                class="button"
              >
                {{ $t('createApp') }}
              </nuxt-link>
              <nuxt-link
                :to="localePath({ name: 'users-edit' })"
                class="button"
              >
                {{ $t('editProfile') }}
              </nuxt-link>
              <div @click="signOut" class="button">
                {{ $t('signOut') }}
              </div>
            </div>
            <div v-else>
              <nuxt-link
                :to="localePath({ name: 'signin' })"
                class="button is-primary"
              >
                <strong>{{ $t('signIn') }}</strong>
              </nuxt-link>
            </div>
          </div>
        </div>
      </client-only>
    </div>
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
      this.$router.push(this.localePath({ path: '/' }))
    }
  }
}
</script>
<style scoped>
nav {
  margin-bottom: 1rem;
}
</style>
