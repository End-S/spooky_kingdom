<script setup lang="ts">
import { ref, onMounted, watch, type Ref } from "vue";
const darkMode: Ref<boolean | null> = ref(null);

onMounted(() => {
  darkMode.value =
    localStorage.theme === "dark"
      ? true
      : localStorage.theme === "light"
      ? false
      : window.matchMedia("(prefers-color-scheme: dark)").matches;
});

watch(darkMode, (isDarkNow, wasDarkBefore) => {
  if (wasDarkBefore !== null) {
    localStorage.theme = isDarkNow ? "dark" : "light";
  }
  let html = document.querySelector("html")!;
  if (isDarkNow) {
    html.classList.add("dark");
  } else {
    html.classList.remove("dark");
  }
});
</script>

<template>
  <o-dropdown aria-role="list" v-model="darkMode" menu-class="min-w-fit bg-gray-900">
    <template #trigger>
      <o-button class="dark:bg-gray-900 !rounded-none">
        <div class="flex gap-1 items-center">
          <o-icon class="text-sm" icon="fa-circle-half-stroke"></o-icon>
          <span class="text-xs font-light hidden md:inline">Theme</span>
        </div>
      </o-button>
    </template>

    <o-dropdown-item aria-role="listitem" :value="true" @click="darkMode = true">
      <div class="flex gap-2 items-center">
        <o-icon class="text-lg md:text-sm" icon="moon"></o-icon>
        <span class="text-xl md:text-sm font-light">Dark</span>
      </div>
    </o-dropdown-item>
    <o-dropdown-item :value="false" aria-role="listitem" @click="darkMode = false">
      <div class="flex gap-2 items-center">
        <o-icon class="text-lg md:text-sm" icon="sun"></o-icon>
        <span class="text-xl md:text-sm font-light">Light</span>
      </div>
    </o-dropdown-item>
  </o-dropdown>
</template>
