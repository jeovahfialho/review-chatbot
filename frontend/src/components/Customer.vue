<template>
  <div class="customer-container">
    <h1 class="title">Cadastrar e Selecionar Cliente</h1>
    <div class="card">
      <div class="card-content">
        <h2 class="subtitle">Cadastrar Cliente</h2>
        <div class="field has-addons">
          <div class="control">
            <input class="input" type="text" v-model="newCustomerName" placeholder="Nome do Cliente">
          </div>
          <div class="control">
            <button class="button is-primary" @click="addCustomer">Adicionar</button>
          </div>
        </div>
        <h2 class="subtitle">Selecionar Cliente</h2>
        <div class="select">
          <select v-model="selectedCustomerID">
            <option v-for="customer in customers" :value="customer.ID" :key="customer.ID">{{ customer.Name }}</option>
          </select>
        </div>
        <button class="button is-link" @click="goToChatbot">Ir para Chatbot</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      newCustomerName: '',
      customers: [],
      selectedCustomerID: null
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
    },
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
        this.customers.push(newCustomer);
        this.selectedCustomerID = newCustomer.ID;
        this.newCustomerName = '';
        this.fetchCustomers();
      } catch (error) {
        console.error('Error adding customer:', error);
      }
    },
    goToChatbot() {
      if (this.selectedCustomerID) {
        this.$router.push({ path: '/chatbot', query: { customerID: this.selectedCustomerID } });
      } else {
        alert('Por favor, selecione um cliente.');
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
