package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/golang_restful_api_crud")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// - ConnMaxIdleTime = kalau taksi nganggur terlalu lama, disuruh pulang
	db.SetConnMaxIdleTime(5 * time.Minute)
	// - ConnMaxLifetime = setelah sekian jam, taksi harus masuk bengkel (ganti baru)
	db.SetConnMaxLifetime(10 * time.Minute)
	// - MaxIdleConns = jumlah taksi yang boleh nunggu (standby) di pangkalan
	db.SetMaxIdleConns(2)
	// - MaxOpenConns = jumlah taksi maksimal yang boleh beroperasi
	db.SetMaxOpenConns(10)

	return db
}
