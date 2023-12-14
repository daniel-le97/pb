<script setup lang="ts">
import { z } from 'zod'

const schema = z.object({
  repoUrl: z.string(),
  installCommand: z.string().nullable(),
  buildCommand: z.string().nullable(),
  startCommand: z.string().nullable(),
  buildPack: z.union([z.literal('nixpacks'), z.literal('dockerfile'), z.literal('docker-compose')]),
  // baseDirectory: z.string().default('/'),
  // publishDirectory: z.string().default('/'),
  // branch: z.string().default('main'),
})

type Schema = z.output<typeof schema>

const state = useActiveProject()
// console.log(state.value)

const options = [
  { label: 'nixpacks', value: 'nixpacks' },
  { label: 'dockerfile', value: 'dockerfile' },
  { label: 'docker-compose', value: 'docker-compose' },
]

async function onSubmit() {

}

// const needsRepo = computed(() => state?.value?.application?.buildCommand === 'nixpacks')
</script>

<template>
  <div v-if="state">
    <UForm :schema=" schema " :state=" state " class="space-y-4" @submit="onSubmit">
      <UFormGroup label="Repo URL" name="repoUrl">
        <UInput v-model=" state.repoURL " type="url" />
      </UFormGroup>

      <div class="flex space-x-4">
        <div class="w-1/2 flex flex-col space-y-3">
          <UFormGroup label="Install command" name="installCommand">
            <UInput v-model=" state.installCommand " placeholder="npm install" type="text" />
          </UFormGroup>
          <UFormGroup label="Build command" name="buildCommand">
            <UInput v-model=" state.buildCommand " placeholder="npm run build" type="text" />
          </UFormGroup>
          <UFormGroup label="Start command" name="startCommand">
            <UInput v-model=" state.startCommand " placeholder="npm run serve" type="text" />
          </UFormGroup>
        </div>
        <div class="w-1/2">
          <UFormGroup label="choose a build pack" name="buildPack">
            <USelect v-model=" state.buildpack" :options=" options " />
          </UFormGroup>
          <UFormGroup label="please specify ports" name="ports">
            <UInput v-model="state.ports" placeholder="3000,3001" />
          </UFormGroup>
          <UFormGroup label="please specify exposed" name="ports">
            <UInput v-model="state.ports" placeholder="3000" />
          </UFormGroup>
          <UFormGroup label="use our proxy?" name="proxy">
            <input v-model="state.managed" type="checkbox">
          </UFormGroup>
          <UFormGroup label="https" name="https">
            <input v-model="state.https" type="checkbox">
          </UFormGroup>
          <!-- <div class="flex gap-2">
            <div class="w-1/2">
              <UFormGroup label="base directory" name="buildPack">
                <USelect v-model=" state.buildPack" :options=" options " />
              </UFormGroup>
            </div>
            <div class="w-1/2">
              <UFormGroup label="publish directory" name="buildPack">
                <USelect v-model=" state.buildPack" :options=" options " />
              </UFormGroup>
            </div>
          </div> -->
        </div>
      </div>

      <RippleBtn type="submit" class="rounded bg-primary">
        Save
      </RippleBtn>
    </UForm>
  </div>
  <div v-else>
    loading...
  </div>
</template>
