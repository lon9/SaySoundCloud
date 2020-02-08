<template>
  <div :class="`modal ${isActive ? 'is-active' : ''}`">
    <div class="modal-background"></div>
    <div class="modal-card">
      <header class="modal-card-head">
        <p class="modal-card-title">{{ $t('passwordRequired') }}</p>
      </header>
      <section class="modal-card-body">
        <div class="field">
          <label class="label">{{ $t('password') }}</label>
          <div class="control">
            <input
              v-model="password"
              @keydown.enter="$emit('onSubmit', password)"
              class="input"
              type="password"
            />
          </div>
        </div>
      </section>
      <footer class="modal-card-foot">
        <button @click="$emit('onSubmit', password)" class="button is-success">
          {{ $t('submitLabel') }}
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
