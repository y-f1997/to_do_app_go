import Vue from 'vue'
import App from './App.vue'
import router from './router'
import 'bootstrap/dist/css/bootstrap.css' // add
import 'bootstrap-vue/dist/bootstrap-vue.css' // add
import jQuery from 'jquery'
import { library } from '@fortawesome/fontawesome-svg-core'
import { fas } from '@fortawesome/free-solid-svg-icons'
import { fab } from '@fortawesome/free-brands-svg-icons'
import { far } from '@fortawesome/free-regular-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import VCalendar from 'v-calendar';
import VueCtkDateTimePicker from 'vue-ctk-date-time-picker';
import 'vue-ctk-date-time-picker/dist/vue-ctk-date-time-picker.css';
import axios from 'axios';
Vue.component('VueCtkDateTimePicker', VueCtkDateTimePicker);

Vue.use(VCalendar);
library.add(fas, fab, far)

Vue.component('font-awesome-icon', FontAwesomeIcon)


axios.interceptors.request.use(req => {
  var cookie;
  cookie = document.cookie;
  var cookieArrayBeforeformat = cookie.split("; ");
  cookie = {};
  cookieArrayBeforeformat.forEach(v => {
    console.log(v)
    let arr = v.split('=')
    cookie[arr[0]] = arr[1];
  })
  console.log(cookie)
  req.headers.Authorization = cookie['isOnSession']
  return req;
})

axios.interceptors.response.use(
  response => {
    console.log('res', response)
    return response;
  },
  error => {
    console.log('error',error)
    if (error.response.data.code == '401'){
      router.push({name:'Login'})
      return Promise.reject(error);
    }else if(error.response.data.code == '500'){
      router.push({name:'Error500'})
      return Promise.reject(error);
    }
     self.error = "<div class='alert alert-warning'>" + error.response.data.code + " :" + error.response.data.error + "</div>"
    return Promise.reject(error);
  }
);

global.jquery = jQuery
global.$ = jQuery
window.$ = window.jQuery = require('jquery')
Vue.config.productionTip = false

new Vue({
  // router設定
  router: router,
  render: h => h(App),
}).$mount('#app')
