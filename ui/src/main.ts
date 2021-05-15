import Vue from 'vue';
import Buefy from 'buefy';
import 'buefy/dist/buefy.css';
import dayjs from 'dayjs';
import { capitalize } from 'lodash-es';
import App from './App.vue';
import router from './router';
import store from './store';
import SvgIcon from './assets/SvgIcon.vue';

Vue.config.productionTip = false;
Vue.use(Buefy, {
  defaultIconPack: null,
  defaultIconComponent: SvgIcon,
});
Vue.filter('time', (d: string) => dayjs(d).format('Do MMM YYYY'));
Vue.filter('caps', (s: string) => capitalize(s));
Vue.filter('publisherLabel', (id: string) => store.getters.getPublisherLabel(id));

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount('#app');
