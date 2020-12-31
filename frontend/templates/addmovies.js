var url='http://localhost:3000'
var body = document.body;
const thisform = document.getElementById('addmovies');
  body.addEventListener('submit',event => {
      event.preventDefault();
      const formData = new FormData(thisform).entries();
      const data = fetch(url.concat('/addmovies'),{
          method:'POST',
          headers: {
            "Accept": "application/json",
            "Content-Type": "application/json"
        },
        body: JSON.stringify(Object.fromEntries(formData))
      })
      data.then(response => response.json())
      .then(data => {
        if (data["status"]===404){
          location.assign('login.html')
        }
        else if (data["status"]===200){
            alert("movie updated")
            location.assign('adminhome.html');
        }else{
          alert("Error has occured.Please check.")
        }
      }
      );
  })
