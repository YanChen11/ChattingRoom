<template>
  <div id="app">
    <ChatRoom/>
  </div>
</template>

<script>
import ChatRoom from './components/ChatRoom.vue'

export default {
  name: 'App',
  components: {
    ChatRoom
  }
}
</script>

<style>
html, body {
  height: 100%;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  height: 100%;
}
</style>
