import Vue from 'vue';
import Vuex from 'vuex';
import { AxiosError } from 'axios';
import { SnackbarProgrammatic as Snackbar } from 'buefy';
import articleStore from './modules/article.store';
import authStore from './modules/auth.store';
import reviewStore from './modules/review.store';
import publisherStore from './modules/publisher.store';

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    as: articleStore,
    auth: authStore,
    rs: reviewStore,
    ps: publisherStore,
  },
  state: (): GlobalState => ({
    err: '',
  }),
  mutations: {
    error(state: GlobalState, err: AxiosError) {
      let errMsg = 'Something went wrong';
      if (err.response) {
        errMsg = `${ err.response?.data?.error } (${ err.response?.status })`;
      }
      state.err = errMsg;
      Snackbar.open(`${ errMsg }`);
      throw err;
    },
  },
  strict: true,
});

interface GlobalState {
  err: string;
}
