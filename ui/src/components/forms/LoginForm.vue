<script setup lang="ts">
import type { Authentication } from "@/common/models/auth.model";
import { useAuthStore } from "@/stores/auth.store";
import { reactive } from "vue";
import { useRoute, useRouter } from "vue-router";

const router = useRouter();
const route = useRoute();
const authStore = useAuthStore();
const formValue: Authentication = reactive({
  username: "",
  password: "",
});

const login = async () => {
  const { success } = await authStore.login(formValue);
  if (success) {
    const redirect = route.query?.redirect?.toString() ?? "archive";
    router.push(redirect);
  }
};
</script>

<template>
  <div class="w-80 m-auto rounded p-4 mt-8">
    <o-field label="Username">
      <o-input type="text" v-model="formValue.username"></o-input>
    </o-field>
    <o-field label="Password">
      <o-input type="password" v-model="formValue.password" password-reveal> </o-input>
    </o-field>
    <div class="flex mt-4 gap-2 justify-end">
      <o-button @click="login()" class="confirm-button">Login</o-button>
    </div>
  </div>
</template>
