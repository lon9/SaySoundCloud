<template>
  <article v-if="message" class="message is-danger">
    <div class="message-body">
      {{ message }}
    </div>
  </article>
</template>
<script>
export default {
  props: {
    message: {
      type: String,
      required: true
    }
  }
}
</script>
