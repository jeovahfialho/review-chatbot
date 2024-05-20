<template>
  <div class="form-container">
    <h1 class="title">Send Message</h1>
    <div class="field">
      <label class="label">Select Customer</label>
      <div class="control">
        <!-- Dropdown to select a customer from the list -->
        <div class="select">
          <select v-model="selectedCustomerID">
            <option v-for="customer in customers" :value="customer.ID" :key="customer.ID">{{ customer.Name }}</option>
          </select>
        </div>
      </div>
    </div>
    <!-- Button to navigate to the chatbot interface -->
    <div class="field">
      <button class="button is-primary" @click="goToChatbot">Go to Chatbot</button>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      customers: [], // Array to hold the list of customers
      selectedCustomerID: null, // ID of the selected customer
    };
  },
  mounted() {
    // Fetch the customers when the component is mounted
    this.fetchCustomers();
  },
  methods: {
    // Fetch the list of customers from the backend
    async fetchCustomers() {
      try {
        const response = await fetch('http://localhost:8080/api/customers');
        if (!response.ok) {
          throw new Error('Failed to fetch customers');
        }
        // Set the response data to the customers array
        this.customers = await response.json();
      } catch (error) {
        console.error('Error fetching customers:', error);
      }
    },
    // Navigate to the chatbot interface with the selected customer ID
    goToChatbot() {
      if (this.selectedCustomerID) {
        this.$router.push({ path: '/chatbot', query: { customerID: this.selectedCustomerID } });
      } else {
        alert('Please select a customer.');
      }
    }
  }
};
</script>

<style>
.form-container {
  max-width: 600px;
  margin: 0 auto;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0,0,0,0.1);
  background-color: #fff;
  text-align: center;
}

.field {
  margin-bottom: 20px;
}
</style>
