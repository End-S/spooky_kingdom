import '@mdi/font/scss/materialdesignicons.scss';
import Vue from 'vue';
import Buefy from 'buefy';
import 'buefy/dist/buefy.css';
import dayjs from 'dayjs';
import { capitalize } from 'lodash-es';
import App from './App.vue';
import router from './router';
import store from './store';

Vue.config.productionTip = false;
Vue.use(Buefy, {});
Vue.filter('time', (d: string) => dayjs(d).format('Do MMM YYYY'));
Vue.filter('caps', (s: string) => capitalize(s));
Vue.filter('publisherLabel', (id: string) => store.getters.getPublisherLabel(id));

new Vue({
  router,
  store,
  data: {
    test: 'TRY PUTTING HTML DEF IN HERE',
  },
  render: (h) => h(App),
}).$mount('#app');
