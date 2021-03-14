import '@mdi/font/scss/materialdesignicons.scss';
import Vue from 'vue';
import Buefy from 'buefy';
import 'buefy/dist/buefy.css';
import App from './App.vue';
import router from './router';
import store from './store';

Vue.config.productionTip = false;
Vue.use(Buefy, {
  defaultFieldLabelPosition: 'on-border',
});

new Vue({
  router,
  store,
  data: {
    test: 'TRY PUTTING HTML DEF IN HERE',
  },
  render: (h) => h(App),
}).$mount('#app');
