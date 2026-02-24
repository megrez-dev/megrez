import Vue from 'vue';
import VueRouter from 'vue-router';
import App from './App.vue';
import router from './router';
import store from './store';
import StorageConfig from '@/config/storage';
import axiosInstance from '@/utils/request';

import VueClipboard from 'vue-clipboard2';
import TDesign from 'tdesign-vue';
import VueStorage from 'vue-ls';
import VueViewer from 'v-viewer';
import { timeAgo } from '@/utils/datetime';

import 'viewerjs/dist/viewer.css';
import 'tdesign-vue/dist/reset.css';

Vue.config.productionTip = false;

Vue.filter('timeAgo', timeAgo);
Vue.use(VueRouter);
Vue.use(TDesign);
Vue.use(VueStorage, StorageConfig);
Vue.use(VueClipboard);
Vue.use(VueViewer, {
  defaultOptions: {
    zIndex: 9999,
  },
});

Vue.prototype.$request = axiosInstance;

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount('#app');
