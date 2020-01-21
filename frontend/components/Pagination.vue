<template>
  <div class="pagination is-centered" role="navigation" aria-label="pagination">
    <a v-if="page !== 1" @click="$emit('prev')" class="pagination-previous"
      >Prev</a
    >
    <a v-else class="pagination-previous" disabled>Prev</a>
    <a @click="$emit('next')" class="pagination-next">Next</a>
    <ul class="pagination-list">
      <li v-if="page !== 1">
        <a @click="$emit('prev')" class="pagination-link">
          {{ page - 1 }}
        </a>
      </li>
      <li>
        <a class="pagination-link is-current" aria-current="page">{{ page }}</a>
      </li>
      <li>
        <a @click="$emit('next')" class="pagination-link">
          {{ page + 1 }}
        </a>
      </li>
    </ul>
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
