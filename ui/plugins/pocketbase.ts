import type { TypedPocketBase } from 'typed-pocketbase'
import type { Schema } from '../Database'

export default defineNuxtPlugin(async (nuxt) => {
  const PocketBase = (await import('pocketbase')).default
  const pb = new PocketBase('http://localhost:8090') as TypedPocketBase<Schema>
  pb.authStore.onChange((token, model) => {
    console.log('authStore.onChange', token, model)
    useUser().value = model
  })

  const cookie = useCookie('pb_auth', {
    path: '/',
    secure: true,
    sameSite: 'strict',
    httpOnly: false, // change to "true" if you want only server-side access
    maxAge: 604800,
  })



  // console.log('cookie.value:', cookie.value)

  if (cookie.value) {
    // console.log('found cookie:', cookie.value);
    
    pb.authStore.loadFromCookie(cookie.value)
  }

  // load the store data from the cookie value
  // pb.authStore.save(cookie.value?.token, cookie.value?.model)

  // send back the default 'pb_auth' cookie to the client with the latest store state
  pb.authStore.onChange(() => {
    console.log('pb.authStore.onChange:', pb.authStore.token, pb.authStore.model)
    cookie.value = pb.authStore.exportToCookie()
  })

  try {
    // get an up-to-date auth store state by verifying and refreshing the loaded auth model (if any)
    pb.authStore.isValid && pb.admins.authRefresh()
  }
  catch (_) {
    // clear the auth store on failed refresh
    pb.authStore.clear()
  }
  return {
    provide: {
      pb,
    },
  }
})
