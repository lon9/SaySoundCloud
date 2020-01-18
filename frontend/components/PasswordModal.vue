<template>
  <div :class="`modal ${isActive ? 'is-active' : ''}`">
    <div class="modal-background"></div>
    <div class="modal-card">
      <header class="modal-card-head">
        <p class="modal-card-title">Password required</p>
      </header>
      <section class="modal-card-body">
        <div class="field">
          <label class="label">Password</label>
          <div class="control">
            <input v-model="password" class="input" type="password" />
          </div>
        </div>
      </section>
      <footer class="modal-card-foot">
        <button @click="$emit('onSubmit', password)" class="button is-success">
          Submit
        </button>
      </footer>
    </div>
  </div>
</template>
<script>
export default {
  props: {
    app: {
      type: Object,
      required: true
    },
    onClose: {
      type: Function,
      required: true
    },
    isActive: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      password: ''
    }
  }
}
</script>
