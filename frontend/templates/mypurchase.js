var url='http://localhost:3000'
const data = fetch(url.concat('/mypurchase'),{
      method:'GET',
      headers: {
        "Accept": "application/json",
        "Content-Type": "application/json"
    },
  })
  data.then(response => response.json())
  .then(data => {
    if (data["status"]===404){
      location.assign('login.html')
    }else{
      console.log(data)
      var container = document.getElementsByTagName('tbody')
        for(item of data) {
          var obj = JSON.parse(item);
          var tr = document.createElement("tr")
          container[0].appendChild(tr)
          var mname = document.createElement("td")
          mname.innerHTML = obj.Mname
          var time = document.createElement("td")
          time.innerHTML = obj.Start_time
          var date = document.createElement("td")
          date.innerHTML = obj.Date
          var quantity = document.createElement("td")
          quantity.innerHTML = obj.Quantity

          tr.appendChild(mname)
          tr.appendChild(time)
          tr.appendChild(date)
          tr.appendChild(quantity)

        }


    }
  })
