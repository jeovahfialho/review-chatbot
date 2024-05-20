<template>
  <div class="table-container">
    <table class="table is-fullwidth is-striped">
      <thead>
        <tr>
          <th>ID</th>
          <th>Name</th>
          <th>Email</th>
        </tr>
      </thead>
      <tbody>
        <!-- Loop through the customers and display each one in a table row -->
        <tr v-for="customer in customers" :key="customer.ID">
          <td>{{ customer.ID }}</td>
          <td>{{ customer.Name }}</td>
          <td>{{ customer.Email }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  data() {
    return {
      customers: [] // Array to hold the list of customers
    };
  },
  mounted() {
    // Fetch the customers when the component is mounted
    this.fetchCustomers();
  },
  methods: {
    async fetchCustomers() {
      try {
        // Send a request to the backend to fetch the list of customers
        const response = await fetch('http://localhost:8080/api/customers');
        if (!response.ok) {
          throw new Error('Failed to fetch customers');
        }
        // Set the response data to the customers array
        this.customers = await response.json();
      } catch (error) {
        console.error('Error fetching customers:', error);
      }
    }
  }
};
</script>

<style>
.table-container {
  margin-top: 20px;
}
</style>
