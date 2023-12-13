export const usePB = () => useNuxtApp().$pb
// const user = useState('user',() => usePB().authStore?.model ?? null)
// usePB().authStore?.onChange((token, model ) => {
//     user.value = model
// })
// export const useUser = () => user