package databaseconn
import (
  "fmt"
  "os"
  "time"
  "github.com/dgrijalva/jwt-go"
)
// type JWTService interface {
//   GenerateToken (name string,admin bool) string
//   ValidateToken (tokenString string)(*jwt.Token,error)
// }

type jwtCustomClaims struct {
  Name string `json:"name"`
  Admin bool `json:"name"`
  jwt.StandardClaims
}

type jwtService struct{
  secretKey string
  issuer string
}
var jwtnew jwtService
 func NewJWTService () *jwtService {
     return &jwtService{
       secretKey :getSecretKey(),
       issuer : "akshit",
       }
 }
func getSecretKey() string  {
  secret := os.Getenv("JWT_SECRET")
  if secret == "" {
    secret = "heywebsite"
  }
  return secret
}

func GenerateToken(username string, admin bool) string {
  fmt.Println(username,"hey")
	claims := &jwtCustomClaims{
		username,
		admin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
			Issuer:    jwtnew.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t,err := token.SignedString([]byte(jwtnew.secretKey))
  PanicError(err)
	return t
}
func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(jwtnew.secretKey), nil
	})
}
