<template>
  <div class="chat-container">
    <h1 class="title">Chatbot Interface</h1>
    <div class="chat-box">
      <!-- Loop through the messages and display each one with appropriate styling -->
      <div v-for="message in messages" :class="messageClass(message.from)" :key="message.id">
        <p><strong>{{ message.from === 'user' ? `${customerName} (${customerID})` : 'Sales' }}:</strong> {{ message.content }}</p>
      </div>
    </div>
    <!-- Input field for the user to type their message -->
    <div class="field">
      <input type="text" class="input" v-model="userMessage" @keyup.enter="sendMessage">
    </div>
    <!-- Button to send the message -->
    <button class="button is-primary" @click="sendMessage">Send</button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      userMessage: '', // The user's message input
      messages: [
        // Initial message from the bot
        { id: 1, from: 'bot', content: 'Hello again! We noticed you\'ve recently received your iPhone 13. We\'d love to hear about your experience. Can you spare a few minutes to share your thoughts?' }
      ],
      customerID: null, // Customer ID from the URL query
      customerName: '', // Customer name to display
      reviewStep: 0, // Step in the review process
      rating: null, // User's rating
    };
  },
  mounted() {
    // Get the customer ID from the URL query and fetch the customer name
    this.customerID = parseInt(this.$route.query.customerID, 10);
    this.fetchCustomerName();
  },
  methods: {
    async fetchCustomerName() {
      try {
        const response = await fetch(`http://localhost:8080/api/customers/${this.customerID}`);
        if (!response.ok) {
          throw new Error('Failed to fetch customer name');
        }
        const customer = await response.json();
        this.customerName = customer.Name; // Set the customer name
      } catch (error) {
        console.error('Error fetching customer name:', error);
      }
    },
    async sendMessage() {
      if (this.userMessage.trim() === '' || !this.customerID) return;

      // Add the user's message to the messages array
      const newMessage = {
        id: this.messages.length + 1,
        from: 'user',
        content: this.userMessage
      };
      this.messages.push(newMessage);

      try {
        // Send the user's message to the backend API
        const response = await fetch('http://localhost:8080/api/chatbot', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({
            customer_id: this.customerID,
            message: this.userMessage
          })
        });

        if (!response.ok) {
          throw new Error('Network response was not ok');
        }

        const data = await response.json();
        // Add the bot's response to the messages array
        this.messages.push({
          id: this.messages.length + 1,
          from: 'bot',
          content: data.response
        });
      } catch (error) {
        console.error('Error:', error);
        // Show an error message if the request fails
        this.messages.push({
          id: this.messages.length + 1,
          from: 'bot',
          content: 'There was a problem processing your request. Please try again later.'
        });
      }

      // Clear the user's message input
      this.userMessage = '';
    },
    // Determine the CSS class for the message based on who sent it
    messageClass(from) {
      return {
        'message-sales': from !== 'user',
        'message-user': from === 'user',
        'has-text-right': from === 'user',
        'has-text-left': from !== 'user'
      };
    }
  }
};
</script>

<style>
.chat-container {
  max-width: 600px;
  margin: 0 auto;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0,0,0,0.1);
  background-color: #fff;
  margin-top: 50px;
}

.chat-box {
  height: 400px;
  overflow-y: scroll;
  border: 1px solid #ccc;
  padding: 10px;
  margin-bottom: 20px;
}

.message-sales {
  background-color: #f0f0f0;
  padding: 10px;
  border-radius: 8px;
  margin-bottom: 10px;
  text-align: left;
  max-width: 70%;
}

.message-user {
  background-color: #d1e7dd;
  padding: 10px;
  border-radius: 8px;
  margin-bottom: 10px;
  text-align: right;
  max-width: 70%;
  margin-left: auto;
}

.field {
  text-align: left;
}
</style>
