<template>
  <div class="table-container">
    <table class="table is-fullwidth is-striped">
      <thead>
        <tr>
          <th>ID</th>
          <th>Nome</th>
          <th>Email</th>
        </tr>
      </thead>
      <tbody>
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
      customers: []
    };
  },
  mounted() {
    this.fetchCustomers();
  },
  methods: {
    async fetchCustomers() {
      try {
        const response = await fetch('http://localhost:8080/api/customers');
        if (!response.ok) {
          throw new Error('Failed to fetch customers');
        }
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
