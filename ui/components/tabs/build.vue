<script lang="ts" setup>
import { eq } from 'typed-pocketbase';


const id = useRoute().params.id as string;

const logs = await usePB().collection( 'logs' ).getFullList( { 'filter': eq( 'project.id', id ) } );
const project = useActiveProject()


</script>

<template>
  <section class="flex  flex-col items-center  justify-center w-full space-y-3 ">
    <div class="lg:flex space-x-3 w-full">
      <label class="block dark:text-white text-gray-700 text-3xl  font-bold mb-2" for="build-logs">
        Build Logs
      </label>
          <!-- <UButton type="button" class=" font-bold py-1" @click="handleClick">
        makeshift build
      </UButton> -->
          <UTooltip>
            <UIcon name="uil:rocket" class="text-2xl" />
            <template #text>
              <span class="italic">Hello World!</span>
            </template>
          </UTooltip>
        </div>
        <UDivider class="w-full" />
        <div class="w-full flex p-2 gap-2">
          <div class="w-4/5">
            <div class="p-2 bg-zinc-700 rounded-md">
                <!-- <pre v-if="buildData.length" id="pre-build" class="w-full h-full overflow-auto whitespace-pre-wrap scrollable-pre"> {{ buildData }}</pre> -->
                <!-- <pre v-else id="pre-build" class="w-full h-full overflow-auto whitespace-pre-wrap scrollable-pre"> {{ 'no builds logged' }}</pre> -->
          </div>
        </div>
        <div class="w-1/5">
            <div v-for="  log   in   logs  " :key=" log.id " class="w-full flex justify-center items-center">
          <BuildLogCard :duration="log.buildTime" :date="log.created" />
        </div>
      </div>

      <!-- BUILD STATUS -->
      <!-- BUILD STATUS -->
    </div>
  </section>
</template>

<style>
.scrollable-pre {
  min-height: 55vh;
  max-height: 500px; /* Set the desired fixed height */
  overflow-y: auto; /* Enable vertical scrolling */
  border: 1px solid #ccc; /* Optional: add a border for styling */
  padding: 10px; /* Optional: add padding for better appearance */
}
</style>
