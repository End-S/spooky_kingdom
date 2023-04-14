<script setup lang="ts">
import { reactive } from "vue";
import { RouterView } from "vue-router";
import { HTTP } from "./common/http";
import ThemeDropdown from "./components/ThemeDropdown.vue";
import TopNav from "./components/TopNav.vue";
import { useGlobalStore } from "./stores/global.store";

const globalStore = reactive(useGlobalStore());
HTTP.interceptors.request.use((req) => {
  globalStore.incrementActiveRequests();
  return req;
});

HTTP.interceptors.response.use((res) => {
  globalStore.decrementActiveRequests();
  return res;
});
</script>

<template>
  <TopNav>
    <ThemeDropdown></ThemeDropdown>
  </TopNav>
  <RouterView />
</template>

<style scoped></style>
