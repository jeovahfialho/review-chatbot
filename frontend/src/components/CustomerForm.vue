<template>
  <div class="form-container">
    <h1 class="title">Cadastrar Cliente</h1>
    <div class="field">
      <label class="label">Nome do Cliente</label>
      <div class="control">
        <input class="input" type="text" v-model="newCustomerName" placeholder="Nome do Cliente">
      </div>
    </div>
    <div class="field">
      <button class="button is-primary" @click="addCustomer">Adicionar</button>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      newCustomerName: ''
    };
  },
  methods: {
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
        this.newCustomerName = '';
        alert('Cliente adicionado com sucesso!');
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
