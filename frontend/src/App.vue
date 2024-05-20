<template>
  <div id="app">
    <!-- Router view to render the matched component for the current route -->
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  name: 'App'
};
</script>

<style>
.container {
  max-width: 800px;
  margin: 0 auto;
}
</style>
