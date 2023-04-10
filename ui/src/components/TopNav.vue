<script setup lang="ts">
import { ref, type Ref } from "vue";
import { RouterLink } from "vue-router";

const menuExpanded: Ref<boolean> = ref(false);
const links: { text: string; route: string }[] = [
  { text: "About", route: "/" },
  { text: "Archive", route: "/archive" },
];
</script>

<template>
  <header class="relative">
    <div
      class="flex justify-between items-center h-12 bg-white dark:bg-gray-900 rounded-sm md:justify-start md:gap-8"
    >
      <h1
        class="text-2xl md:text-3xl ml-4 font-['Gabriela'] subpixel-antialiased text-gray-800 dark:text-gray-200"
      >
        <a href="/archive">Spooky Kingdom</a>
      </h1>
      <nav class="hidden md:flex gap-4 font-['Gabriela']" v-for="link in links" :key="link.route">
        <RouterLink :to="link.route">{{ link.text }}</RouterLink>
      </nav>
      <div class="flex flex-1 justify-end">
        <slot></slot>
      </div>
      <o-button @click="menuExpanded = !menuExpanded" class="block md:hidden" icon-right="bars" />
    </div>
    <div v-if="menuExpanded" class="absolute w-full z-10">
      <nav
        v-for="link in links"
        class="flex flex-col bg-neutral-50 dark:bg-gray-800 md:hidden"
        :key="link.route"
      >
        <RouterLink
          @click="menuExpanded = !menuExpanded"
          class="text-lg hover:bg-neutral-200 dark:hover:bg-gray-600 p-2"
          :to="link.route"
          >{{ link.text }}</RouterLink
        >
      </nav>
    </div>
  </header>
</template>
