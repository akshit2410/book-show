//Query for returning rows and Exec and Prepare for repeated use without returning
package databaseconn
import(
    "fmt"
    "database/sql"
    R "book_show/backend/auth"
  _ "github.com/go-sql-driver/mysql"
  "golang.org/x/crypto/bcrypt"
)
//connecting to DB
func Connect() *sql.DB {
  db,err := sql.Open("mysql","root@tcp(127.0.0.1)/showtime")
  if err!=nil {
      fmt.Println(err.Error())
  }
  err= db.Ping()
  if err != nil {
    fmt.Println(err.Error())
  }
  return db
}


//returning error function
func PanicError(err error )  {
if err !=nil {
  panic(err.Error())
}
}


//checking if user already exist
func Checkuserexist(reg R.Register ) int  {
  db := Connect()
  querycheck := "SELECT username FROM user WHERE username = ?"
  rows,err :=db.Query(querycheck,reg.USER)
  PanicError(err)
  defer rows.Close()
  if rows.Next() {
    return 201
  }
  return 200

}

//hashing the password of the fucntion]
func hashandsaltpass(password string) string  {
  p := []byte (password)
  hash,err :=   bcrypt.GenerateFromPassword(p,bcrypt.DefaultCost)
  PanicError(err)
  return string(hash)

}

//comparing password at login time
func ComparePasswords(hashedPwd string, plainPwd string) bool {
    byteHash := []byte(hashedPwd)
    p :=[]byte(plainPwd)
    err := bcrypt.CompareHashAndPassword(byteHash,p)
    if err != nil {

        return false
    }
    return true
}

//creating a new user
func Create(reg R.Register) int {
  hashedpass := hashandsaltpass(reg.PASSWORD)
  db := Connect()
  queryinsert := "INSERT INTO user(username, password, email) VALUES(?, ?, ?)"
  if(Checkuserexist(reg)==201){
    return 201
  }else {
  insert,err :=  db.Prepare(queryinsert)
  PanicError(err)
  insert.Exec(reg.USER,hashedpass,reg.EMAIL)
  defer db.Close()
  return 200
}
}
//check if user present
func Checkuserpresence(log R.Login) (string,int)  {
  var (
    username string
    password string
    level int
  )
  db := Connect()
  querycheck := "SELECT username,password,level FROM user WHERE username = ?"
  rows,err :=db.Query(querycheck,log.USER)
  PanicError(err)
  defer rows.Close()
  if rows.Next() {
    err := rows.Scan(&username,&password,&level)
    PanicError(err)
    return password,level
  }
  return "",0
}
//logging in a user
func CheckCredentials(log R.Login) (int,string) {
  hashpass,level := Checkuserpresence(log)
  fmt.Println(level)
  if hashpass != "" {
    if ComparePasswords(hashpass,log.PASSWORD) && level==0 {
      return 200,log.USER
    }else if ComparePasswords(hashpass,log.PASSWORD) && level==1 {
      return 500,log.USER
      }else{
      return 201,""
    }
  }else {
    return 202,""
  }

}


func Addmovie(info R.Addmovie) int  {
  db :=  Connect()
  fmt.Println(info)
  queryinsert := "INSERT INTO movie(mname,mdesc,mdirector,mduration,image) VALUES(?,?,?,?,?)"
  insert,err :=  db.Prepare(queryinsert)
  PanicError(err)
  insert.Exec(info.MNAME,info.MDESC,info.MDIRECTOR,info.MDURATION,info.IMAGE)
  defer db.Close()
  return 200

}

func Addtime(info R.Addtime) int  {
  db :=  Connect()
  querycheck := "SELECT mname FROM movie WHERE mname = ?"
  rows,err :=db.Query(querycheck,info.MNAME)
  PanicError(err)
  defer rows.Close()
  if rows.Next() {
    queryinsert := "INSERT INTO timing(mname,start_time,end_time,date,quantity,price) VALUES(?,?,?,?,?,?)"
    insert,err :=  db.Prepare(queryinsert)
    PanicError(err)
    insert.Exec(info.MNAME,info.START_TIME,info.END_TIME,info.DATE,info.QUANTITY,info.PRICE)
    defer db.Close()
    return 200
  }

return 400
}
