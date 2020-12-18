  // Form Submission
  var url='http://localhost:3000'
  var body = document.body;
  const thisform = document.getElementById('register-form');
  const thisform1 = document.getElementById('login-form');
  body.addEventListener('submit',event => {
    if (event.target !==thisform && event.target!==thisform1){
      return
    }
    else if (event.target===thisform) {
      event.preventDefault();
     const formData = new FormData(thisform).entries();
     const data = fetch(url.concat('/register'),{
         method:'POST',
         headers: {
           "Accept": "application/json",
           "Content-Type": "application/json"
       },
       body: JSON.stringify(Object.fromEntries(formData))
     })
     data.then(response => response.json())
     .then(data => {
       if (data["status"]===200){
           location.assign('login.html');
       }else if(data["status"]===201) {
         alert("user already exist")
       }else{
         alert("Error has occured.Please check.")
       }
     }
     );
    }
    else {
      event.preventDefault();
     const formData = new FormData(thisform1).entries();
     console.log(formData)
   const data = fetch(url.concat('/login'),{
         method:'POST',
         headers: {
           "Accept": "application/json",
           "Content-Type": "application/json"
       },
       body: JSON.stringify(Object.fromEntries(formData))
     })
      console.log(data)
     data.then(response => response.json())
     .then(data => {
       if (data["status"]===200){
         // const data = fetch('http://localhost:3000/homepage/',{
         //       method:'GET',
         //       headers: {
         //         "Accept": "application/json",
         //         "Content-Type": "application/json"
         //     },
         if(data["session"]){
         location.assign('homepage.html');
       }else{
         alert("you arent logged in yet")
       }

       }else if(data["status"]===201) {
         alert("Wrong Password!please, try again")
       }else{
         alert("User doesn't exist")
       }
     }
     );
    }
  })
