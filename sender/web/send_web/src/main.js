/*
import { createApp } from 'vue'
import App from './App.vue'

createApp(App).mount('#app')
*/

import { createApp } from 'vue'
import App from './App.vue';
import ElementPlus from 'element-plus';
import 'element-ui/lib/theme-chalk/index.css';
const app = createApp(App)
app.use(ElementPlus)
app.mount('#app')