import 'core-js/stable'
import Vue from 'vue'
//import CoreuiVuePro from '@coreui/vue-pro'
import CoreuiVuePro from '../node_modules/@coreui/vue-pro/src/index.js'
import App from './App'
import router from './router/index'
import { iconsSet as icons } from './assets/icons/icons.js'
import store from './store'
import i18n from './i18n.js'
import axios from 'axios'
import SmartTable from 'vuejs-smart-table'
import 'ant-design-vue/dist/antd.css';
import './assets/scss/index.css';
import jsPDF from 'jspdf'
import autoTable from 'jspdf-autotable'
// -----------------------------------------
import { Datetime } from 'vue-datetime'
// You need a specific loader for CSS files
import 'vue-datetime/dist/vue-datetime.css'
Vue.use(CoreuiVuePro)
Vue.use(SmartTable)

Vue.use(Datetime)
const api = axios.create({
    baseURL: process.env.VUE_APP_API_ENDPOINT,
    withCredentials: true
})
Vue.prototype.$log = console.log.bind(console)
Vue.prototype.$PDF = jsPDF
Vue.prototype.$Version = process.env.VUE_APP_VERSION
Vue.prototype.$autoTable = autoTable
new Vue({
    el: '#app',
    router,
    store,
    icons,
    i18n,
    provide: {
        api: api
    },
    template: '<App/>',
    components: {
        App
    }
})