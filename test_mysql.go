package main

import (
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
    "fmt"

    //"time"
)

func main() {
    db, err := sql.Open("mysql", "tom:tom@tcp(10.10.10.225:3306)/mysql?charset=utf8")
    checkErr(err)
    _,err =db.Exec("CREATE TABLE `userinfo` (`uid` INT(10) NOT NULL AUTO_INCREMENT,`username` VARCHAR(64) NULL DEFAULT NULL,`departname` VARCHAR(64) NULL DEFAULT NULL,`created` DATE NULL DEFAULT NULL,PRIMARY KEY (`uid`))")
    checkErr(err)
	fmt.Printf("create table finished!!!!!\n")
    db.Close()

}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}


