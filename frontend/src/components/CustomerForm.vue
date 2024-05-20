<template>
  <div class="form-container">
    <h1 class="title">Register Customer</h1>
    <div class="field">
      <label class="label">Customer Name</label>
      <!-- Input field for new customer name -->
      <div class="control">
        <input class="input" type="text" v-model="newCustomerName" placeholder="Customer Name">
      </div>
    </div>
    <!-- Button to add a new customer -->
    <div class="field">
      <button class="button is-primary" @click="addCustomer">Add</button>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      newCustomerName: '' // The new customer's name
    };
  },
  methods: {
    async addCustomer() {
      if (this.newCustomerName.trim() === '') return;

      try {
        // Send a request to the backend to add a new customer
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
        // Clear the input field and show a success message
        this.newCustomerName = '';
        alert('Customer added successfully!');
      } catch (error) {
        console.error('Error adding customer:', error);
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
