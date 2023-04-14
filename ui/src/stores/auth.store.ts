import { login, type LoginResponse } from "@/api/auth";
import { isErrorResponse } from "@/common/models/api.model";
import type { Authentication } from "@/common/models/auth.model";
import { defineStore } from "pinia";
import { useGlobalStore } from "./global.store";

export interface AuthState {
  jwt: string;
}

export const useAuthStore = defineStore("authStore", {
  state: (): AuthState => ({
    jwt: "",
  }),
  actions: {
    setJwt(jwt: string) {
      this.jwt = jwt;
      localStorage.setItem("JWT", jwt);
    },
    async login(auth: Authentication): Promise<{ success: boolean }> {
      let success = false;
      const res: LoginResponse = await login(auth);

      if (isErrorResponse(res)) {
        const globalStore = useGlobalStore();
        globalStore.error(res);
      } else {
        success = true;
        this.setJwt(res.jwt);
      }

      return { success };
    },
  },
});
