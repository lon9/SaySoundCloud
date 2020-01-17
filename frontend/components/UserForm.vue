<template>
  <div class="container">
    <div class="field">
      <div class="label">Name</div>
      <div class="control">
        <input v-model="profile.name" class="input" type="text" />
      </div>
    </div>
    <div class="field">
      <label class="label">Description</label>
      <div class="control">
        <textarea v-model="profile.description" class="textarea"></textarea>
      </div>
    </div>
    <div class="field is-grouped">
      <div class="control">
        <button @click="onSubmit" class="button is-link">Submit</button>
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
