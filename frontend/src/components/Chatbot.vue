<template>
  <div class="chat-container">
    <h1 class="title">Chatbot Interface</h1>
    <div class="chat-box">
      <div v-for="message in messages" :class="messageClass(message.from)" :key="message.id">
        <p><strong>{{ message.from === 'user' ? `${customerName} (${customerID})` : 'Sales' }}:</strong> {{ message.content }}</p>
      </div>
    </div>
    <div class="field">
      <input type="text" class="input" v-model="userMessage" @keyup.enter="sendMessage">
    </div>
    <button class="button is-primary" @click="sendMessage">Enviar</button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      userMessage: '',
      messages: [
        { id: 1, from: 'bot', content: 'Hello again! We noticed you\'ve recently received your iPhone 13. We\'d love to hear about your experience. Can you spare a few minutes to share your thoughts?' }
      ],
      customerID: null,
      customerName: '',
      reviewStep: 0,
      rating: null,
    };
  },
  mounted() {
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
        this.customerName = customer.Name;
      } catch (error) {
        console.error('Error fetching customer name:', error);
      }
    },
    async sendMessage() {
      if (this.userMessage.trim() === '' || !this.customerID) return;

      const newMessage = {
        id: this.messages.length + 1,
        from: 'user',
        content: this.userMessage
      };
      this.messages.push(newMessage);

      try {
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
        this.messages.push({
          id: this.messages.length + 1,
          from: 'bot',
          content: data.response
        });
      } catch (error) {
        console.error('Error:', error);
        this.messages.push({
          id: this.messages.length + 1,
          from: 'bot',
          content: 'Houve um problema ao processar sua solicitação. Por favor, tente novamente mais tarde.'
        });
      }

      this.userMessage = '';
    },
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
