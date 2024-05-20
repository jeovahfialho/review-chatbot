<template>
  <div class="dashboard-container">
    <h1 class="title">Dashboard</h1>
    <div class="menu">
      <!-- Links to navigate to the customer registration and send message pages -->
      <router-link to="/add-customer" class="button is-primary">Register Customer</router-link>
      <router-link to="/send-message" class="button is-link">Send Message</router-link>
    </div>
    <div class="tables">
      <h2 class="subtitle">Customer List</h2>
      <!-- Component to display the customer table -->
      <CustomerTable />
      <!-- Component to display the review table -->
      <ReviewTable />
    </div>
  </div>
</template>

<script>
// Import the customer and review table components
import CustomerTable from './CustomerTable.vue';
import ReviewTable from './ReviewTable.vue';

export default {
  components: {
    CustomerTable,
    ReviewTable
  }
};
</script>

<style>
.dashboard-container {
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
}

.menu {
  display: flex;
  justify-content: space-between;
  margin-bottom: 20px;
}

.tables {
  margin-top: 20px;
}
</style>
