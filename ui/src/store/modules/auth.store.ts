import Vue from 'vue';
import Vuex from 'vuex';
import { login } from '@/api/auth';
import { Authentication } from '@/common/models/auth.model';
import router from '@/router';
import { AxiosError } from 'axios';

Vue.use(Vuex);

export default {
  state: (): AuthState => ({
    jwt: '',
  }),
  mutations: {
    setJWT(state: AuthState, jwt: string) {
      state.jwt = jwt;
      localStorage.setItem('JWT', jwt);
    },
  },
  actions: {
    async login({ commit }: { commit: Function }, auth: Authentication) {
      const res = await login(auth)
        .catch((err: AxiosError) => commit('error', err));
      commit('setJWT', res.data.jwtToken);
      await router.push('/review');
    },
  },
};

interface AuthState {
  jwt: string;
}
