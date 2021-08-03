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
    function Get(url){
    let request = new XMLHttpRequest()
    request.open("GET", url, false)
    request.send()
    return request.resposeText
    }

    function mainConsume(){
      console.log(Get(fetch("http://localhost:10000")))
    }

    mainConsume()
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
