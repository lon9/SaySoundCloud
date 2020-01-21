<template>
  <div class="card">
    <div class="card-content">
      <p class="title is-4">{{ app.name }}</p>
      <p class="subtitle is-6">
        <nuxt-link
          :to="localePath({ name: 'users-id', params: { id: app.userId } })"
        >
          @{{ app.user.name }}
        </nuxt-link>
      </p>
    </div>
    <div class="card-footer">
      <nuxt-link
        :to="localePath({ name: 'apps-id', params: { id: app.ID } })"
        class="card-footer-item"
      >
        Detail
      </nuxt-link>
      <nuxt-link
        :to="localePath({ name: 'apps-id-room', params: { id: app.ID } })"
        class="card-footer-item"
      >
        Enter
      </nuxt-link>
      <client-only>
        <nuxt-link
          v-if="user && user.ID === app.userId"
          :to="localePath({ name: 'apps-id-edit', params: { id: app.ID } })"
          class="card-footer-item"
        >
          Edit
        </nuxt-link>
      </client-only>
    </div>
  </div>
</template>
<script>
export default {
  props: {
    app: {
      type: Object,
      required: true
    }
  },
  computed: {
    user() {
      return this.$store.state.user
    }
  }
}
</script>
<style scoped>
.card {
  margin-top: 0.5rem;
  margin-bottom: 0.5rem;
}
</style>
