<template>
  <div class="container">
    <div v-if="user" class="content">
      <h4>{{ user.name }}</h4>
      <p style="white-space:pre-line;">
        {{ user.description }}
      </p>
    </div>
  </div>
</template>
<script>
export default {
  async asyncData({ $axios, params }) {
    try {
      const res = await $axios.$get(`/users/${params.id}`)
      return {
        user: res.result
      }
    } catch {}
  }
}
</script>
