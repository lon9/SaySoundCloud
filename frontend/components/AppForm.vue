<template>
  <div class="container">
    <div class="field">
      <div class="label">Name</div>
      <div class="control">
        <input v-model="app.name" class="input" type="text" name="appname" />
      </div>
    </div>
    <div class="field">
      <label class="label">Description</label>
      <div class="control">
        <textarea
          v-model="app.description"
          class="textarea"
          name="appdescription"
        ></textarea>
      </div>
    </div>
    <div class="field">
      <label class="label">Password</label>
      <div class="control">
        <input
          v-model="app.password"
          class="input"
          type="password"
          name="apppassword"
        />
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
    app: {
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
