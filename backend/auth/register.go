package auth

type Register struct {
    USER string `json:"username" binding : required`
  PASSWORD string `json:"password" binding : required`
  CONPASSWORD string `json:"password_confirm" binding : required`
  EMAIL string `json:"email" binding : required`
}
type Login struct {
  USER string `json:"username" binding : required`
  PASSWORD string `json:"password" binding : required`
}

type Movie struct {
  MNAME string `json:"mname" binding : required`
  MDESC string `json:"mdesc" binding : required`
  MDIRECTOR string `json:"mdirector" binding : required`
  MDURATION string `json:"mduration" binding : required`
  IMAGE []byte `json:"image" binding : required`
}
type Buy struct {
  MNAME string `json:"mname" binding : required`
  MTIME string `json:"mtime" binding : required`
  MDATE string `json:"mdate" binding : required`
  QUANTITY string `json:"quantity" binding : required`
}
