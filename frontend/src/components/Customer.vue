<template>
  <div class="customer-container">
    <h1 class="title">Register and Select Customer</h1>
    <div class="card">
      <div class="card-content">
        <h2 class="subtitle">Register Customer</h2>
        <div class="field has-addons">
          <!-- Input field for new customer name -->
          <div class="control">
            <input class="input" type="text" v-model="newCustomerName" placeholder="Customer Name">
          </div>
          <!-- Button to add a new customer -->
          <div class="control">
            <button class="button is-primary" @click="addCustomer">Add</button>
          </div>
        </div>
        <h2 class="subtitle">Select Customer</h2>
        <!-- Dropdown to select a customer from the list -->
        <div class="select">
          <select v-model="selectedCustomerID">
            <option v-for="customer in customers" :value="customer.ID" :key="customer.ID">{{ customer.Name }}</option>
          </select>
        </div>
        <!-- Button to navigate to the chatbot interface -->
        <button class="button is-link" @click="goToChatbot">Go to Chatbot</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      newCustomerName: '', // The new customer's name
      customers: [], // List of existing customers
      selectedCustomerID: null // ID of the selected customer
    };
  },
  mounted() {
    // Fetch the list of customers when the component is mounted
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
        this.customers = await response.json();
      } catch (error) {
        console.error('Error fetching customers:', error);
      }
    },
    // Add a new customer to the backend
    async addCustomer() {
      if (this.newCustomerName.trim() === '') return;

      try {
        const response = await fetch('http://localhost:8080/api/customers', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({ Name: this.newCustomerName, Email: `${this.newCustomerName}@example.com` })
        });
        if (!response.ok) {
          throw new Error('Failed to add customer');
        }
        const newCustomer = await response.json();
        this.customers.push(newCustomer); // Add the new customer to the list
        this.selectedCustomerID = newCustomer.ID; // Set the new customer as the selected customer
        this.newCustomerName = ''; // Clear the input field
        this.fetchCustomers(); // Refresh the customer list
      } catch (error) {
        console.error('Error adding customer:', error);
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
.customer-container {
  max-width: 600px;
  margin: 0 auto;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0,0,0,0.1);
  background-color: #fff;
  margin-top: 50px;
  text-align: center;
}

.card {
  margin: 0 auto;
  max-width: 500px;
  padding: 20px;
  border: 1px solid #ddd;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0,0,0,0.1);
}

.card-content {
  padding: 20px;
}

.field {
  margin-bottom: 20px;
}

.select {
  margin-bottom: 20px;
}
</style>
