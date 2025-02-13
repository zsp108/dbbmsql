package zmysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Options struct {
	Host                  string
	Username              string
	Password              string
	Port                  int
	Database              string
	MaxIdleConnections    int
	MaxOpenConnections    int
	MaxConnectionLifeTime time.Duration
}

func Connect(opts Options) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		opts.Username,
		opts.Password,
		opts.Host,
		opts.Port,
		opts.Database,
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	if err = db.Ping(); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	db.SetMaxIdleConns(opts.MaxIdleConnections)
	db.SetMaxOpenConns(opts.MaxOpenConnections)
	db.SetConnMaxLifetime(opts.MaxConnectionLifeTime)
	return db, nil
}

func Select(db *sql.DB, sql string) (rows *sql.Rows, err error) {
	rows, err = db.Query(sql)
	if err != nil {
		return nil, err
	}
	return rows, nil

}

func Execute(db *sql.DB, sql string) (result sql.Result, err error) {
	result, err = db.Exec(sql)
	if err != nil {
		return nil, err
	}
	return result, nil
}
