package main

import (
	"database/sql"
	"time"
)


const(
  maxOpenDbConn = 25
  maxIdleDbConn = 25
  maxDBLifetime = 5 * time.Minute 
)


func initDb(dsn string) (*sql.DB, error){
  db, err := sql.Open("mysql", dsn)

  if err != nil {
    return nil, err
  }


  if err = db.Ping(); err != nil{
    return nil , err
  }

  db.SetMaxOpenConns(maxOpenDbConn)
  db.SetMaxIdleConns(maxIdleDbConn)
  db.SetConnMaxLifetime(maxDBLifetime)


  return db, nil

}
