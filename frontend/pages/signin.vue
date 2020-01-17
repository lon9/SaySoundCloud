<template>
  <div>
    <client-only>
      <firebase-ui />
    </client-only>
  </div>
</template>
<script>
import FirebaseUi from '~/components/FirebaseUi'
export default {
  components: {
    FirebaseUi
  }
}
</script>
