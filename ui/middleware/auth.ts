export default defineNuxtRouteMiddleware((to, from) => {
  const user = usePB().authStore.isAuthRecord
  // In a real app you would probably not redirect every route to `/`
  // however it is important to check `to.path` before redirecting or you
  // might get an infinite redirect loop
  console.log('running auth middleware')

  if (to.path !== '/login' && !user)
    return navigateTo('/login')
})
