<template>
  <div class="pagination is-centered" role="navigation" aria-label="pagination">
    <a v-if="page !== 1" @click="$emit('prev')" class="pagination-previous"
      >Prev</a
    >
    <a v-else class="pagination-previous" disabled>Prev</a>
    <a @click="$emit('next')" class="pagination-next">Next</a>
  </div>
</template>
<script>
export default {
  props: {
    page: {
      type: Number,
      required: true
    }
  }
}
</script>
