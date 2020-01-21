<template>
  <div class="container">
    <div class="field">
      <div class="label">{{ $t('nameLabel') }}</div>
      <div class="control">
        <input
          v-model="profile.name"
          class="input"
          type="text"
          name="username"
        />
      </div>
    </div>
    <div class="field">
      <label class="label">{{ $t('descriptionLabel') }}</label>
      <div class="control">
        <textarea
          v-model="profile.description"
          class="textarea"
          name="userdescription"
        ></textarea>
      </div>
    </div>
    <div class="field is-grouped">
      <div class="control">
        <button @click="onSubmit" class="button is-link">
          {{ $t('submitLabel') }}
        </button>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  props: {
    profile: {
      type: Object,
      required: true
    },
    onSubmit: {
      type: Function,
      required: true
    }
  }
}
</script>
