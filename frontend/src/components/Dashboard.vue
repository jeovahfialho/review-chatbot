<template>
  <div class="dashboard-container">
    <h1 class="title">Dashboard</h1>
    <div class="menu">
      <router-link to="/add-customer" class="button is-primary">Cadastrar Cliente</router-link>
      <router-link to="/send-message" class="button is-link">Enviar Mensagem</router-link>
    </div>
    <div class="tables">
      <h2 class="subtitle">Lista de Clientes</h2>
      <CustomerTable />
      <ReviewTable />
    </div>
  </div>
</template>

<script>
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
