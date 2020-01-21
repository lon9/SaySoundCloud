require('dotenv').config()

export default {
  mode: 'universal',
  /*
   ** Headers of the page
   */
  head: {
    title: process.env.APP_TITLE || '',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      {
        hid: 'description',
        name: 'description',
        content: process.env.APP_DESCRIPTION || ''
      },
      {
        hid: 'og:site_name',
        property: 'og:site_name',
        content: process.env.OG_SITE_NAME
      },
      {
        hid: 'og:type',
        property: 'og:type',
        content: 'website'
      },
      {
        hid: 'og:url',
        property: 'og:url',
        content: process.env.OG_URL
      },
      {
        hid: 'og:title',
        property: 'og:title',
        content: process.env.OG_TITLE
      },
      {
        hid: 'og:description',
        property: 'og:description',
        content: process.env.OG_DESCRIPTION
      },
      {
        hid: 'og:image',
        property: 'og:image',
        content: process.env.OG_IMAGE
      },
      {
        hid: 'twitter:card',
        name: 'twitter:card',
        content: process.env.TWITTER_CARD
      },
      {
        hid: 'twitter:site',
        name: 'twitter:site',
        content: process.env.TWITTER_SITE
      }
    ],
    link: [{ rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }]
  },
  /*
   ** Customize the progress-bar color
   */
  loading: { color: '#fff' },
  /*
   ** Global CSS
   */
  css: ['firebaseui/dist/firebaseui.css'],
  /*
   ** Plugins to load before mounting the App
   */
  plugins: [
    '~/plugins/firebase',
    '~/plugins/axios',
    { src: '~/plugins/nuxt-client-init.js', ssr: false },
    { src: '~/plugins/router-option.js', ssr: false }
  ],
  axios: {
    baseURL: process.env.BASE_URL
  },
  /*
   ** Nuxt.js dev-modules
   */
  buildModules: [
    // Doc: https://github.com/nuxt-community/eslint-module
    '@nuxtjs/eslint-module'
  ],
  /*
   ** Nuxt.js modules
   */
  modules: [
    // Doc: https://github.com/nuxt-community/modules/tree/master/packages/bulma
    '@nuxtjs/bulma',
    '@nuxtjs/pwa',
    // Doc: https://github.com/nuxt-community/dotenv-module
    '@nuxtjs/dotenv',
    '@nuxtjs/axios',
    '@nuxtjs/markdownit',
    [
      'nuxt-i18n',
      {
        locales: [
          {
            code: 'en',
            file: 'en-US.js'
          },
          {
            code: 'ja',
            file: 'ja-JP.js'
          }
        ],
        defaultLocale: 'en',
        langDir: 'lang/',
        lazy: true
      }
    ]
  ],
  markdownit: {
    injected: true
  },
  manifest: {
    name: process.env.PWA_NAME,
    shortname: process.env.PWA_SHORTNAME,
    description: process.env.PWA_DESCRIPTION
  },
  /*
   ** Build configuration
   */
  env: {
    // for firebase
    API_KEY: process.env.API_KEY,
    AUTH_DOMAIN: process.env.AUTH_DOMAIN,
    DATABASE_URL: process.env.DATABASE_URL,
    PROJECT_ID: process.env.PROJECT_ID,
    STORAGE_BUCKET: process.env.STORAGE_BUCKET,
    MESSAGING_SENDER_ID: process.env.MESSAGING_SENDER_ID,
    APP_ID: process.env.APP_ID,
    // for app
    BASE_URL: process.env.BASE_URL,
    APP_TITLE: process.env.APP_TITLE,
    APP_DESCRIPTION: process.env.APP_DESCRIPTION,
    USE_FIREBASE: process.env.USE_FIREBASE,
    SOUND_BASE_URL: process.env.SOUND_BASE_URL,
    // for open graph
    OG_SITE_NAME: process.env.OG_SITE_NAME,
    OG_URL: process.env.OG_URL,
    OG_TITLE: process.env.OG_URL,
    OG_DESCRIPTION: process.env.OG_DESCRIPTION,
    OG_IMAGE: process.env.OG_IMAGE,
    // for twitter card
    TWITTER_CARD: process.env.TWITTER_CARD,
    TWITTER_SITE: process.env.TWITTER_SITE,
    // for pwa
    PWA_NAME: process.env.PWA_NAME,
    PWA_SHORTNAME: process.env.PWA_SHORTNAME,
    PWA_DESCRIPTION: process.env.PWA_DESCRIPTION
  },
  build: {
    babel: {
      presets({ isServer }) {
        return [
          [
            require.resolve('@nuxt/babel-preset-app'),
            {
              buildTarget: isServer ? 'server' : 'client',
              corejs: { version: 3 }
            }
          ]
        ]
      }
    },
    postcss: {
      preset: {
        features: {
          customProperties: false
        }
      }
    },
    /*
     ** You can extend webpack config here
     */
    extend(config, ctx) {}
  }
}
