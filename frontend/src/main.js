import Vue from 'vue';
import App from './App.vue';
import VueRouter from 'vue-router';
import Dashboard from './components/Dashboard.vue';
import CustomerForm from './components/CustomerForm.vue';
import MessageForm from './components/MessageForm.vue';
import Chatbot from './components/Chatbot.vue';

Vue.config.productionTip = false;

Vue.use(VueRouter);

const routes = [
  { path: '/', component: Dashboard },
  { path: '/add-customer', component: CustomerForm },
  { path: '/send-message', component: MessageForm },
  { path: '/chatbot', component: Chatbot }
];

const router = new VueRouter({
  routes
});

new Vue({
  router,
  render: h => h(App),
}).$mount('#app');
