<template>
  <img alt="Vue logo" src="./assets/logo.png">
  <!--<HelloWorld msg="Welcome to Your Vue.js App"/>-->
  <LayoutOne />
</template>

<script>
/*import HelloWorld from './components/HelloWorld.vue'*/
import LayoutOne from './components/LayoutOne.vue'


export default {
  name: 'App',
  components: {
    LayoutOne
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
