<template>
  <div class="dashboard-container">
    <h1 class="title">Dashboard</h1>
    <div class="menu">
      <router-link to="/add-customer" class="button is-primary">Cadastrar Cliente</router-link>
      <router-link to="/send-message" class="button is-link">Enviar Mensagem</router-link>
    </div>
    <div class="form-container">
      <h2 class="subtitle">Review Statistics</h2>
      <form @submit.prevent="fetchReviewStats">
        <div class="field">
          <label class="label">Start Date</label>
          <div class="control">
            <input type="date" v-model="startDate" class="input" required>
          </div>
        </div>
        <div class="field">
          <label class="label">End Date</label>
          <div class="control">
            <input type="date" v-model="endDate" class="input" required>
          </div>
        </div>
        <div class="field">
          <label class="label">Interval (minutes)</label>
          <div class="control">
            <input type="number" v-model="intervalMinutes" class="input" required>
          </div>
        </div>
        <div class="field">
          <div class="control">
            <button type="submit" class="button is-primary">Get Statistics</button>
          </div>
        </div>
      </form>
      <div v-if="averageRatingTotal !== null">
        <h2 class="subtitle">Average Rating: {{ averageRatingTotal }}</h2>
      </div>
      <canvas id="reviewChart" v-if="chartData"></canvas>
    </div>
    <div class="tables">
      <h2 class="subtitle">Lista de Clientes</h2>
      <CustomerTable />
      <ReviewTable />
    </div>
  </div>
</template>

<script>
import Vue from 'vue'
import { Line } from 'vue-chartjs'
import Chart from 'chart.js'
import CustomerTable from './CustomerTable.vue'
import ReviewTable from './ReviewTable.vue'

export default {
  components: {
    CustomerTable,
    ReviewTable,
    LineChart: Line
  },
  data() {
    return {
      startDate: '',
      endDate: '',
      intervalMinutes: 10,
      averageRatingTotal: null,
      chartData: null,
      chartInstance: null // Adicionando chartInstance ao data
    }
  },
  methods: {
    async fetchReviewStats() {
      try {
        console.log(`Fetching stats from ${this.startDate} to ${this.endDate} with interval ${this.intervalMinutes} minutes`);
        const response = await fetch(`http://localhost:8080/api/reviews/average?start_date=${this.startDate}&end_date=${this.endDate}&interval_minutes=${this.intervalMinutes}`);
        if (!response.ok) {
          throw new Error('Failed to fetch review statistics');
        }
        const data = await response.json();
        console.log('Received data:', data);
        this.averageRatingTotal = data.average_rating_total;
        this.chartData = {
          labels: data.interval_ratings.map(interval => `${interval.start_date} - ${interval.end_date}`),
          datasets: [
            {
              label: 'Average Rating',
              data: data.interval_ratings.map(interval => interval.average_rating),
              fill: false,
              borderColor: 'rgb(75, 192, 192)',
              tension: 0.1
            }
          ]
        }
        Vue.nextTick(() => {
          this.renderChart();
        });
      } catch (error) {
        console.error('Error fetching review statistics:', error);
      }
    },
    renderChart() {
      if (this.chartInstance) {
        this.chartInstance.destroy();
      }
      const ctx = document.getElementById('reviewChart').getContext('2d');
      this.chartInstance = new Chart(ctx, {
        type: 'line',
        data: this.chartData,
        options: {
          responsive: true,
          scales: {
            x: {
              display: true,
              title: {
                display: true,
                text: 'Interval'
              }
            },
            y: {
              display: true,
              title: {
                display: true,
                text: 'Average Rating'
              },
              suggestedMin: 0,
              suggestedMax: 5
            }
          }
        }
      });
    }
  }
}
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

.form-container {
  margin-bottom: 20px;
}

.tables {
  margin-top: 20px;
}

canvas {
  max-width: 100%;
}
</style>
