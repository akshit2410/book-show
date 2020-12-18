package main
import (
  //"fmt"
  "github.com/gin-gonic/contrib/static"
  "os"
  "github.com/gin-gonic/gin"
)
func main()  {
  port := os.Getenv("PORT")

 if port == "" {
   port = "1234"
 }
  server :=gin.Default()
  server.Use(static.Serve("/", static.LocalFile("./templates", true)))
server.Run(":"+port)
}
