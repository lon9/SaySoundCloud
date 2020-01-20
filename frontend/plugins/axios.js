export default function({ $axios, store, redirect }) {
  $axios.onRequest((config) => {
    const token = store.state.token
    if (token) {
      config.headers.common.Authorization = 'Bearer ' + token
    }
  })
  $axios.onError((error) => {
    if (error.response.status === 401) {
      redirect('/signin')
    }
  })
}
