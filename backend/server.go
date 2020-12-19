package main
import (
  "fmt"
  R "book_show/backend/auth"
  // "os"
  "net/http"
  "github.com/gin-gonic/gin"
   "github.com/gin-contrib/cors"
   db "book_show/backend/databaseconn"
   //"github.com/gin-gonic/contrib/sessions"
  // "github.com/gin-contrib/sessions/cookie"
  // "github.com/go-session/gin-session"

)
type Use struct {
  Name string
}
var u Use
func main()  {
    fmt.Println("hey its working backend")
 //  port := os.Getenv("PORT")
 // if port == "" {
 //   port = "3000"
 // }
 port := "3000"
 db.Connect()
 type REG =R.Register
 type LOG = R.Login
 type BUY = R.Buy
 server :=gin.Default()
 server.Use(cors.Default())

  //get index


  server.GET("/",func (c *gin.Context){
    u.Name = ""
    c.Header("Access-Control-Allow-Origin", "*")
    c.JSON(http.StatusOK, gin.H{"message": "item created"})
    })

//get login

  // post login
  server.POST("/login",func(c *gin.Context){
        c.Header("Access-Control-Allow-Origin", "*")
        var log LOG
        c.BindJSON(&log)
        fmt.Println(log)
        res,username := db.CheckCredentials(log)
        u.Name = username
         // c.HTML(http.StatusOK, "http://localhost:1234/login.html", nil)
         c.JSON(200, gin.H{"status":res,"session":u.Name})
    })
  // post register
  server.POST("/register",func(c *gin.Context){
      c.Header("Access-Control-Allow-Origin", "*")
      var reg REG
      c.BindJSON(&reg)
      res := db.Create(reg)
       // c.HTML(http.StatusOK, "http://localhost:1234/login.html", nil)
       c.JSON(200, gin.H{"status":res})
    })
//get unloghome
      server.GET("/unloghome",func(c *gin.Context) {
        res := db.Connmovies()
        c.JSON(200, res)
      })
//post buyticket
  server.POST("/buy",func (c *gin.Context)  {
  c.Header("Access-Control-Allow-Origin", "*")
      var buy BUY
      c.BindJSON(&buy)
      fmt.Println(buy)
      if u.Name !=""{
        // res := db.Buytickets(buy)
        res := db.Buytickets(buy,u.Name)
        c.JSON(200, gin.H{"status":res})
      }else {
        c.JSON(200, gin.H{"status":404})
      }
  })
  //get home
  server.GET("/home",func(c *gin.Context) {
    if u.Name!=""{
    res := db.Connmovies()
    c.JSON(200, res)
  }else{
    c.JSON(200, gin.H{"status":404})
  }
  })
//mypurchase
server.GET("/mypurchase",func(c *gin.Context) {
  if u.Name!=""{
  res := db.Purchase(u.Name)
  fmt.Println(res)
  c.JSON(200, res)
}else{
  c.JSON(200, gin.H{"status":404})
}
})


  //Logout
server.GET("/logout",func (c* gin.Context)  {
  fmt.Println("hey")
u.Name = ""
c.JSON(200, gin.H{"status":200})

})
server.Run(":"+port)
}
