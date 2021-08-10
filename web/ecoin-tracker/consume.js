async function Get(url){
    let request = new XMLHttpRequest()
    await request.open("GET", url, true)
    await request.send()
    return request.resposeText
    }

    function mainConsume(){
      var json_obj = Get("http://localhost:9000")
      console.log(json_obj)
    }

    mainConsume()