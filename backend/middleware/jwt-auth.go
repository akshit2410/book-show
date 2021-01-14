import (
  D "book_show/backend/databaseconn"
  "net/http"
)

func AuthorizeJWT() gin.HandlerFunc{
  return func (c *gin.Context)  {
    const BEARER_SCH = "Akshit"
    authHeader := c.GetHeader("Authorization")
    tokenString := authHeader[len(BEARER_SCH):]
    token ,err =: D.NewJWTService().ValidateToken(tokenString)
    if token.Valid {
      claims := token.Claims.(jwt.MapClaims)
    }else {
      c.AbortWithStatus(http.StatusUnauthorized)
    }
  }

}
