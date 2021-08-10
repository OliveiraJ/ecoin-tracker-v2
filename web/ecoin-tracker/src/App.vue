<template>
  <div>
    <table id="table"></table>
  </div>
</template>

<script>


export default {
  name: 'App',
  components: {
    
    },
  mounted(){
    const url = "http://localhost:9000/"
    fetch(url)
    .then(response => response.json())
    .then(data => console.log(data))
    .catch(err => console.log(err))
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
