package databaseconn
import(
    "fmt"
    "encoding/json"

    //"database/sql"
    //R "book_show/backend/auth"
  //_ "github.com/go-sql-driver/mysql"
  //d "book_show/backend/databaseconn"
)

func Connmovies() []string {
  var (
    Mname string
    Mdesc string
    Image []byte
  )
  type Movie struct {
    Mname string
    Mdesc string
    Image []byte
  }
  db:=Connect()
  querycheck := "SELECT mname,mdesc,image FROM movie"
  rows,err :=db.Query(querycheck)
  PanicError(err)
  defer rows.Close()
  var str []string
  defer db.Close()
  for rows.Next() {
        err := rows.Scan(&Mname, &Mdesc, &Image)
        PanicError(err)
        m := Movie {
          Mname: Mname,
          Mdesc: Mdesc,
          Image:Image,
        }
        b, err := json.Marshal(m)
        PanicError(err)
        str = append(str,string(b))
      }

      // fmt.Println(str)
      return str
}


func Purchase(user string ) []string  {
  var (
    Mname string
    start_time string
    date string
    quantity int
  )
type Userbuy struct {
  Mname string
  Start_time string
  Date string
  Quantity int
  }
  db:=Connect()
  rows,err := db.Query("SELECT mname,start_time,date,quantity FROM userbuy where user=?",user)
  PanicError(err)
  defer rows.Close()
  defer db.Close()
  var str []string
  for rows.Next() {
    err := rows.Scan(&Mname, &start_time, &date, &quantity)
    PanicError(err)
    c := Userbuy {
      Mname :Mname,
      Start_time :start_time,
      Date :date,
      Quantity :quantity,
    }
    fmt.Println(c)
    b, err := json.Marshal(c)
    PanicError(err)
    str = append(str,string(b))
  }
  fmt.Println(str)
  return str
}
