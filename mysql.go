package main

import (
    "github.com/astaxie/beego/logs"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    logs.Debug("main()")
    db, err := sql.Open("mysql", "ailumiyana:qwedsa@tcp(127.0.0.1:3306)/mysql")
    if err != nil {
      logs.Error("sql Open() err", err)
    }

    stmt, err := db.Prepare("Insert user_info set id=?,name=?")
    if err != nil {
      logs.Error("sql Prepare() err", err)
    }

    stmt.Exec(1, "sola")
    stmt.Exec(2, "ailumiyana")

    rows, err :=db.Query("SELECT * FROM user_info")
    if err != nil {
      logs.Error("sql Query() err", err)
    }

    for rows.Next() {
        var uid int
        var username string

        err = rows.Scan(&uid, &username)
        if err != nil {
          logs.Error("sql rows.Scan() err", err)
        }
        logs.Debug(uid, username)
    }

}