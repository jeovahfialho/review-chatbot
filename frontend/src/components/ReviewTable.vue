<template>
  <div>
    <h2 class="subtitle">Avaliações de Produtos</h2>
    <table class="table is-fullwidth is-striped">
      <thead>
        <tr>
          <th>ID</th>
          <th>Customer ID</th>
          <th>Product ID</th>
          <th>Rating</th>
          <th>Comments</th>
          <th>Review Time</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="review in reviews" :key="review.ID">
          <td>{{ review.ID }}</td>
          <td>{{ review.CustomerID }}</td>
          <td>{{ review.ProductID }}</td>
          <td>{{ review.Rating }}</td>
          <td>{{ review.Comments }}</td>
          <td>{{ new Date(review.ReviewTime).toLocaleString() }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  data() {
    return {
      reviews: []
    };
  },
  mounted() {
    this.fetchReviews();
  },
  methods: {
    async fetchReviews() {
      try {
        const response = await fetch('http://localhost:8080/api/reviews');
        if (!response.ok) {
          throw new Error('Failed to fetch reviews');
        }
        this.reviews = await response.json();
      } catch (error) {
        console.error('Error fetching reviews:', error);
      }
    }
  }
};
</script>
