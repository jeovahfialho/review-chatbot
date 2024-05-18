<template>
  <div class="form-container">
    <h1 class="title">Enviar Mensagem</h1>
    <div class="field">
      <label class="label">Selecionar Cliente</label>
      <div class="control">
        <div class="select">
          <select v-model="selectedCustomerID">
            <option v-for="customer in customers" :value="customer.ID" :key="customer.ID">{{ customer.Name }}</option>
          </select>
        </div>
      </div>
    </div>
    <div class="field">
      <button class="button is-primary" @click="goToChatbot">Ir para Chatbot</button>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      customers: [],
      selectedCustomerID: null,
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
