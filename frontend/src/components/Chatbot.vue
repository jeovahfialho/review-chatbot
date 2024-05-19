<template>
  <div class="chatbot-container">
    <div class="chat-window">
      <div v-for="message in messages" :key="message.id" class="message" :class="{'user-message': message.sender === 'user', 'bot-message': message.sender === 'bot'}">
        <span v-if="message.sender === 'bot'">Sales: </span>{{ message.text }}
      </div>
    </div>
    <div class="input-container">
      <input v-model="userInput" @keyup.enter="sendMessage" placeholder="Type a message..." />
      <button @click="sendMessage">Send</button>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      userInput: '',
      messages: [],
      customerId: 1 // Replace this with the actual customer ID as needed
    };
  },
  methods: {
    sendMessage() {
      if (this.userInput.trim() === '') return;

      // Add the user's message to the chat
      this.messages.push({ id: Date.now(), text: this.userInput, sender: 'user' });

      // Send the message to the backend
      axios.post('http://localhost:8080/api/chatbot', {
        customer_id: this.customerId,
        message: this.userInput
      })
      .then(response => {
        // Add the bot's response to the chat
        this.messages.push({ id: Date.now() + 1, text: response.data.response, sender: 'bot' });
      })
      .catch(error => {
        console.error('Error sending message:', error);
      });

      // Clear the input field
      this.userInput = '';
    }
  }
};
</script>

<style scoped>
.chatbot-container {
  width: 400px;
  margin: 0 auto;
  border: 1px solid #ccc;
  border-radius: 8px;
  overflow: hidden;
  display: flex;
  flex-direction: column;
  height: 500px;
}

.chat-window {
  flex-grow: 1;
  padding: 16px;
  overflow-y: auto;
}

.message {
  margin-bottom: 8px;
  word-wrap: break-word;
}

.user-message {
  text-align: right;
  color: blue;
}

.bot-message {
  text-align: left;
  color: green;
}

.input-container {
  display: flex;
  padding: 16px;
  border-top: 1px solid #ccc;
}

input {
  flex-grow: 1;
  padding: 8px;
  border: 1px solid #ccc;
  border-radius: 4px;
}

button {
  padding: 8px 16px;
  border: none;
  background-color: #007bff;
  color: white;
  border-radius: 4px;
  cursor: pointer;
}

button:hover {
  background-color: #0056b3;
}
</style>
