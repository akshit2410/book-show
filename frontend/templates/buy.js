var url='http://localhost:3000'
var body = document.body;
const thisform = document.getElementById('buy-form');
  body.addEventListener('submit',event => {
      event.preventDefault();
      const formData = new FormData(thisform).entries();
      const data = fetch(url.concat('/buy'),{
          method:'POST',
          headers: {
            "Accept": "application/json",
            "Content-Type": "application/json"
        },
        body: JSON.stringify(Object.fromEntries(formData))
      })
      data.then(response => response.json())
      .then(data => {
        if(data["status"]===200){
          location.assign("homepage.html")
          alert("u bought the ticket ")
        }else if (data["status"]===404) {
          location.assign("login.html")
        }else{
          alert("movie no available,fill the form again")

        }

      })
  })
