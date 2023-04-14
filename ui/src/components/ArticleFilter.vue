<script setup lang="ts">
import FilterForm from "./forms/FilterForm.vue";
import { ref, type Ref } from "vue";

const menuExpanded: Ref<boolean> = ref(false);
const modalActive: Ref<boolean> = ref(false);
</script>

<template>
  <!-- dropdown -->
  <div class="hidden md:inline">
    <!-- dropdown gauze -->
    <div
      v-if="menuExpanded"
      @click="menuExpanded = !menuExpanded"
      class="fixed bg-black/50 z-10 w-full h-full left-0 top-0"
    ></div>
    <!-- dropdown gauze ends -->
    <div class="relative">
      <!-- dropdown trigger button -->
      <o-button @click="menuExpanded = !menuExpanded">
        <div class="flex gap-1 items-center">
          <o-icon class="text-sm" icon="fa-filter"></o-icon>
          <span class="text-sm hidden md:inline">Filter</span>
        </div>
      </o-button>
      <!-- dropdown trigger button ends -->
      <!-- dropdown content -->
      <div
        v-if="menuExpanded"
        class="absolute right-0 bg-gray-100 dark:bg-gray-800 p-4 w-96 rounded-sm z-20"
      >
        <FilterForm @submit="menuExpanded = false" @cancel="menuExpanded = false"></FilterForm>
      </div>
      <!-- dropdown content ends -->
    </div>
  </div>
  <!-- dropdown ends -->

  <!-- modal (mobile) -->
  <div class="inline md:hidden">
    <!-- modal trigger button -->
    <o-button size="medium" variant="primary" @click="modalActive = !modalActive">
      <div class="flex gap-1 items-center">
        <o-icon class="text-sm" icon="fa-filter"></o-icon>
        <span class="text-sm hidden md:inline">Filter</span>
      </div>
    </o-button>
    <!-- modal trigger button ends -->

    <!-- modal content -->
    <o-modal v-model:active="modalActive" class="p-4">
      <FilterForm @submit="modalActive = false" @cancel="modalActive = false"></FilterForm>
    </o-modal>
    <!-- modal content ends -->
  </div>
  <!-- modal (mobile) ends -->
</template>
