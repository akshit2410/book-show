
  function getdata() {
  var url='http://localhost:3000/'
  const data = fetch(url,{
        method:'GET',
        headers: {
          "Accept": "application/json",
          "Content-Type": "application/json"
      },
    })
    data.then(response => response.json())
    .then(data => console.log(data));
}
getdata();
