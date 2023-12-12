<!-- <template>
  <div>

    {{ authData }}
    {{ pb.authStore.isValid }}
    {{ pb.authStore.token }}
    {{ pb.authStore?.model }}
    {{ url }}
    <button @click="pb.authStore.clear()">Logout</button>
    <button @click="pb.admins">Logout</button>
  </div>
</template>
<script lang="ts" setup>
const pb = usePB();
const authData = await pb.admins.authWithPassword( 'test@example.com', '1234567890' );
// pb.admins.
// after the above you can also access the auth data from the authStore
console.log( pb.authStore.isValid );
console.log( pb.authStore.token );
console.log( pb.authStore?.model?.id );

const url = pb.getFileUrl(pb.authStore.model!, pb.authStore.model!.avatar);
</script> -->
<script setup lang="ts">
import { z } from 'zod'
import type { FormSubmitEvent } from '#ui/types'

const schema = z.object({
  email: z.string().email('Invalid email'),
  password: z.string().min(8, 'Must be at least 8 characters')
})

type Schema = z.output<typeof schema>

const state = reactive({
  email: undefined,
  password: undefined
})

async function onSubmit (event: FormSubmitEvent<Schema>) {
  // Do something with data
  console.log(event.data)
}
</script>

<template>
  <UForm :schema="schema" :state="state" class="space-y-4" @submit="onSubmit">
    <UFormGroup label="Email" name="email">
      <UInput v-model="state.email" />
    </UFormGroup>

    <UFormGroup label="Password" name="password">
      <UInput v-model="state.password" type="password" />
    </UFormGroup>

    <UButton type="submit">
      Submit
    </UButton>
  </UForm>
</template>

