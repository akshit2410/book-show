package databaseconn
import (
  "fmt"
  R "book_show/backend/auth"
  "strings"
  "time"
   "strconv"
)

func  Buytickets(buy R.Buy,user string) int {
  var (
      mname string
      start_time string
      end_time string
      date string
      quantity int
      price int
  )
db :=  Connect()
query := "SELECT mname,start_time,end_time,date,quantity,price FROM timing"
queryinsert := "INSERT INTO userbuy(user,mname,start_time,date,quantity) VALUES(?,?,?,?,?)"
queryupdate := "UPDATE timing SET quantity=? WHERE date = ? AND mname = ?"
rows,err := db.Query(query)
PanicError(err)
defer rows.Close()
for rows.Next() {
      err := rows.Scan(&mname, &start_time, &end_time, &date, &quantity, &price)
      PanicError(err)
    }
    starttime := strings.Replace(start_time, ":", "",-1)
    endtime := strings.Replace(end_time, ":", "",-1)
    t := buy.MTIME+":00"
    rtime := strings.Replace(t, ":", "",-1)
    start,err := time.Parse("150405",starttime)
    end,err := time.Parse("150405",endtime)
    real,err := time.Parse("150405",rtime)
    PanicError(err)
     quant,err := strconv.Atoi(buy.QUANTITY)
       PanicError(err)
    if buy.MNAME == mname && buy.MDATE == date && quantity !=0 && quantity > quant{
      if real.After(start) && real.Before(end){
        quantnow := quantity - quant
        fmt.Println(quantity,buy.QUANTITY,quantnow)
        insert,err :=  db.Prepare(queryinsert)
        PanicError(err)
        insert.Exec(user,buy.MNAME,buy.MTIME,buy.MDATE,buy.QUANTITY)
        update,err :=  db.Prepare(queryupdate)
        PanicError(err)
        update.Exec(quantnow,buy.MDATE,buy.MNAME)
        return 200
      }
    }else{
      return 201
    }
    defer db.Close()
    return 0
}
