<!-- components/Tabs.vue -->
<script setup lang="ts">
import { LazyTabsBuild, LazyTabsCompose, LazyTabsConfiguration, LazyTabsSecrets } from '#components'

const id = useRoute().params.id as string

const project = await usePB().collection('projects').getOne(id)

useActiveProject().value = project

const selectedTab = ref(0)

const tabs = computed(() => {
  if (true) {
    return [
      { label: 'configuration', component: LazyTabsConfiguration },
      { label: 'compose', component: LazyTabsCompose },
      { label: 'build', component: LazyTabsBuild },
      { label: 'secrets', component: LazyTabsSecrets },
    ]
  }
  return [
    { label: 'configuration', component: LazyTabsConfiguration },

    { label: 'build', component: LazyTabsBuild },
    { label: 'secrets', component: LazyTabsSecrets },
  ]
})

function selectTab(index: number) {
  selectedTab.value = index
}
</script>

<template>
  <div class="p-5 w-full">
    <div class="flex items-center justify-between mb-20">
      <div class="flex items-center space-x-2 ">
        <h1 class="text-2xl font-bold">
          Configurations
        </h1>
        <UBadge>Healthy</UBadge>
      </div>
      <!-- Left side with buttons and icons -->

      <TabsApplicationControlBar class="flex gap-2" />

      <!-- Right side with page title and badge -->
    </div>
    <div class="flex  px-5 gap-3">
      <div class="p-2  border h-fit rounded-lg">
        <RippleBtn
          v-for="(tab, index) in tabs " :key=" index " :class=" { 'bg-white text-black': selectedTab === index } "
          class="w-full p-2 mb-2 text-left max rounded" @click="selectTab(index)"
        >
          {{ tab.label }}
        </RippleBtn>
      </div>

      <div class=" w-full">
        <div v-for="(tab, index) in tabs " :key=" index ">
          <div v-if=" selectedTab === index ">
            <component :is="tab.component" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
