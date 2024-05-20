import Vue from 'vue';
import App from './App.vue';
import VueRouter from 'vue-router';
import Dashboard from './components/Dashboard.vue';
import CustomerForm from './components/CustomerForm.vue';
import MessageForm from './components/MessageForm.vue';
import Chatbot from './components/Chatbot.vue';

Vue.config.productionTip = false;

// Use the VueRouter plugin
Vue.use(VueRouter);

// Define the routes for the application
const routes = [
  { path: '/', component: Dashboard }, // Route for the dashboard
  { path: '/add-customer', component: CustomerForm }, // Route for the customer registration form
  { path: '/send-message', component: MessageForm }, // Route for the message form
  { path: '/chatbot', component: Chatbot } // Route for the chatbot interface
];

// Create a new router instance
const router = new VueRouter({
  routes // short for `routes: routes`
});

// Create and mount the root Vue instance
new Vue({
  router, // Inject the router into the Vue instance
  render: h => h(App), // Render the App component
}).$mount('#app'); // Mount the Vue instance to the element with id 'app'
