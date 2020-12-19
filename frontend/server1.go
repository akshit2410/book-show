package main
import (
  "fmt"
  "github.com/gin-gonic/contrib/static"
  "os"
  "github.com/gin-gonic/gin"
)
func main()  {
  port := os.Getenv("PORT")
 if port == "" {
   port = "2404"
 }
  server :=gin.Default()
  server.Use(static.Serve("/", static.LocalFile("./templates", true)))
  fmt.Println("hey its working",port)
server.Run(":"+port)
}
