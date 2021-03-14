import Vue from 'vue';
import Vuex from 'vuex';
import { AxiosError } from 'axios';
import { get } from '@/api/publishers';
import { Publisher } from '@/common/models/publisher.model';

Vue.use(Vuex);

export default {
  state: (): PublisherState => ({
    publishers: [],
  }),
  getters: {
    getPublisherLabel: (state: PublisherState) => (id: string) => {
      const match = state.publishers.find(p => p.id === id)?.label;
      return match || '?';
    },
  },
  mutations: {
    setPublishers(state: PublisherState, publishers: Publisher[]) {
      state.publishers = publishers;
    },
  },
  actions: {
    async getPublishers({ commit }: { state: PublisherState; commit: Function }) {
      const res = await get()
        .catch((err: AxiosError) => commit('error', err));
      commit('setPublishers', res.data.publishers);
    },
  },
};

interface PublisherState {
  publishers: Publisher[];
}
