<template>
  <div class="field has-addons">
    <div class="control">
      <input
        v-model="query"
        @keydown.enter="$emit('search', query)"
        :placeholder="placeHolder"
        class="input"
        type="text"
      />
    </div>
    <div class="control">
      <a @click="$emit('search', query)" class="button is-info">
        {{ $t('search') }}
      </a>
    </div>
  </div>
</template>
<script>
export default {
  props: {
    placeHolder: {
      type: String,
      default: 'Search word'
    }
  },
  data() {
    return {
      query: ''
    }
  }
}
</script>
