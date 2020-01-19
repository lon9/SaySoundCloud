<template>
  <div class="card">
    <div class="card-content">
      <p class="title is-4">{{ app.name }}</p>
      <p class="subtitle is-6">@{{ app.user.name }}</p>
    </div>
    <div class="card-footer">
      <nuxt-link
        :to="localePath({ name: 'apps-id', params: { id: app.ID } })"
        class="card-footer-item"
      >
        Detail
      </nuxt-link>
      <a @click="$emit('onEnter', app)" class="card-footer-item">
        Enter
      </a>
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
  }
}
</script>
<style scoped>
.card {
  margin-top: 0.5rem;
  margin-bottom: 0.5rem;
}
</style>
