<template>
  <div class="container">
    <div v-if="app" class="content">
      <h4>{{ app.name }}</h4>
      <p style="white-space:pre-line;">
        {{ app.description }}
      </p>
      <p v-if="app.isPassword">Need password</p>
    </div>
  </div>
</template>
<script>
export default {
  async asyncData({ $axios, params }) {
    try {
      const res = await $axios.$get(`/apps/${params.id}`)
      return {
        app: res.result
      }
    } catch {}
  }
}
</script>
