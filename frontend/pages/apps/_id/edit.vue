<template>
  <div v-if="app" class="container">
    <ErrorView :message="errorMsg" />
    <AppForm :app="app" :onSubmit="onAppSubmit" />
    <div class="access-token-container">
      <div class="field">
        <label class="label">{{ $t('accessTokenLabel') }}</label>
        <div class="control">
          <div class="field has-addons">
            <div class="control is-expanded">
              <input :value="accessToken" class="input" type="text" readonly />
            </div>
            <div class="control">
              <a @click="copy(accessToken)" class="button is-info is-outlined">
                {{ $t('copy') }}
              </a>
            </div>
            <div class="control">
              <a @click="renewToken" class="button is-info is-outlined">
                {{ $t('renewToken') }}
              </a>
            </div>
          </div>
        </div>
      </div>
      <div class="field">
        <label class="label">URL (POST)</label>
        <div class="control">
          <div class="field has-addons">
            <div class="control is-expanded">
              <input
                :value="`${baseUrl}/apps/${$route.params.id}/cmd`"
                class="input"
                type="text"
                readonly
              />
            </div>
            <div class="control">
              <a
                @click="copy(`${baseUrl}/apps/${$route.params.id}/cmd`)"
                class="button is-info is-outlined"
              >
                {{ $t('copy') }}
              </a>
            </div>
          </div>
        </div>
      </div>
      <div class="field">
        <label class="label">Body</label>
      </div>
      <div v-html="$md.render(body)" />
    </div>
  </div>
</template>
<script>
import AppForm from '~/components/AppForm'
import ErrorView from '~/components/ErrorView'
export default {
  components: { AppForm, ErrorView },
  data() {
    return {
      errorMsg: '',
      app: {
        name: '',
        description: ''
      },
      accessToken: '',
      baseUrl: process.env.BASE_URL
    }
  },
  computed: {
    body() {
      return `\`\`\`
{
    "name": "[cmdName]",
    "accessToken": "${this.accessToken}"
}
\`\`\``
    }
  },
  async mounted() {
    try {
      const res = await this.$axios.$get(`/apps/${this.$route.params.id}/owner`)
      this.app.name = res.result.name
      this.app.description = res.result.description
      this.accessToken = res.result.accessToken
    } catch (e) {}
  },
  methods: {
    async onAppSubmit() {
      this.errorMsg = ''
      try {
        const res = await this.$axios.$put(
          `/apps/${this.$route.params.id}`,
          this.app
        )
        this.$router.push(
          this.localePath({ path: `/apps/${res.result.ID}/edit` })
        )
      } catch {
        this.errorMsg = this.$t('failedToEdit')
      }
    },
    copy(text) {
      navigator.clipboard.writeText(text)
    },
    async renewToken() {
      this.errorMsg = ''
      try {
        const res = await this.$axios.$put(
          `/apps/${this.$route.params.id}/renewtoken`
        )
        this.accessToken = res.result.accessToken
      } catch {
        this.errorMsg = this.$t('failedToRenew')
      }
    }
  }
}
</script>
<style scoped>
.access-token-container {
  margin-top: 1em;
}
</style>
