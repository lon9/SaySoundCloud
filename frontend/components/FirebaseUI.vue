<template>
  <div>
    <div id="firebaseui-auth-container"></div>
  </div>
</template>
<script>
export default {
  mounted() {
    if (process.browser) {
      const firebaseui = require('firebaseui')
      const ui =
        firebaseui.auth.AuthUI.getInstance() ||
        new firebaseui.auth.AuthUI(this.$auth)

      const config = {
        credentialHelper: firebaseui.auth.CredentialHelper.NONE,
        signInOptions: [
          this.$firebase.auth.GoogleAuthProvider.PROVIDER_ID,
          this.$firebase.auth.EmailAuthProvider.PROVIDER_ID
        ],
        signInFlow: 'popup',
        callbacks: {
          signInSuccessWithAuthResult: this.signInResult
        }
      }
      ui.disableAutoSignIn()
      if (this.$store.state.user) {
        this.$router.push(this.localePath({ path: '/' }))
      } else {
        ui.start('#firebaseui-auth-container', config)
      }
    }
  },
  methods: {
    auth() {
      return new Promise((resolve, reject) => {
        this.$auth.onAuthStateChanged((user) => {
          resolve(user || false)
        })
      })
    },
    async signInResult() {
      const user = await this.auth()
      const token = await user.getIdToken()
      this.$store.commit('setToken', token)
      if (await this.$store.dispatch('createUser')) {
        this.$router.push(this.localePath({ path: '/users/edit' }))
        return false
      }
      await this.$store.dispatch('getLoginUser')
      this.$router.push(this.localePath({ path: '/' }))
      return false
    }
  }
}
</script>
