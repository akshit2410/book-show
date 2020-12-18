var url='http://localhost:3000'
const data = fetch(url.concat('/home'),{
      method:'GET',
      headers: {
        "Accept": "application/json",
        "Content-Type": "application/json"
    },
  })
  data.then(response => response.json())
  .then(data => {
    if (data["status"]===404){
      // alert("login first")
      location.assign('login.html')
    }else{
    var container = document.getElementsByClassName('wrapper')
    for(item of data) {
      var obj = JSON.parse(item);
      console.log(obj.Mname)
      var container1 = document.createElement("div");
      container1.className = 'card'
      container[0].appendChild(container1)
      var image = document.createElement('img')
      image.src = "data:image/gif;base64,"+obj.Image;
      container1.appendChild(image)
      var desc = document.createElement('div')
      desc.className = 'descriptions'
      container1.appendChild(desc)
      var div1 = document.createElement('div')
      desc.appendChild(div1)
      var h1 = document.createElement('h1')
      h1.id = 'Mname'
      h1.innerHTML = obj.Mname
      div1.appendChild(h1)
      var p = document.createElement('p')
      p.innerHTML = obj.Mdesc
      div1.appendChild(p)
      var button = document.createElement("BUTTON")
      button.type = "submit"
      button.innerHTML = "BOOK"
      desc.appendChild(button)
    }


  }
  });
  // function logout() {
  //   const data1 = fetch(url.concat('/logout'),{
  //         method:'GET',
  //         headers: {
  //           "Accept": "application/json",
  //           "Content-Type": "application/json"
  //       },
  //     })
  //     data1.then(response => response.json())
  //     .then(data1 => {
  //       if(data1["status"]===200){
  //         location.assign('index.html')
  //       }
  //     })
  //
  // }
