package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load .env")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	fmt.Println("DSN:", dsn)  // Tambah ini untuk cek DSN di log container

	var db *gorm.DB
	var errDB error

	// Retry koneksi sampai max 10 kali, delay 3 detik
	for i := 0; i < 10; i++ {
		db, errDB = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if errDB == nil {
			break
		}
		fmt.Printf("Gagal konek ke DB, retry %d/10...\n", i+1)
		time.Sleep(3 * time.Second)
	}

	if errDB != nil {
		panic(errDB)
	}

	DB = db
}
