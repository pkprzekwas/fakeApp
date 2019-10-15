<template>
  <div id="app">
    <Layout msg="Welcome"/>
  </div>
</template>

<script>
import Layout from './components/Layout.vue'

export default {
  name: 'app',
  components: {
    Layout
  }
}
</script>

<style>
  * {
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
  }

  body {
    background-color: #171717;
    color: white;
    padding: 100px 50px;
  }

  #app {
    font-family: 'Fira Code', sans-serif;
    text-align: center;
    margin-top: 60px;
  }
</style>
