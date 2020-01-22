<template>
  <div>
    <client-only>
      <firebase-ui />
    </client-only>
  </div>
</template>
<script>
import FirebaseUI from '~/components/FirebaseUI'
export default {
  components: {
    FirebaseUI
  }
}
</script>
