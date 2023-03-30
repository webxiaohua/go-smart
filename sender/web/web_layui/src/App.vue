<template>
  <img alt="Vue logo" src="./assets/logo.png">
  <HelloWorld msg="Test" name="Robin"/>
  <LayuiOne message="hello" />
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'
import LayuiOne from './components/LayuiOne'
export default {
  name: 'App',
  components: {
    HelloWorld,
    LayuiOne
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
