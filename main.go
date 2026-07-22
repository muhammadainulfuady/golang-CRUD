package main

import (
	"database/sql"
	"fmt"
	"log"

	"golang_standart_project/repositories"
	"golang_standart_project/services"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Koneksi ke Database MySQL
	dsn := "root:@tcp(localhost:3306)/belajar_golang_database_v2"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Inisialisasi layer
	repo := repositories.NewProductRepository(db)
	service := services.NewProductService(repo)

	// --- Simulasi Penggunaan ---

	// 1. Tambah Data
	id, err := service.AddProduct("Laptop Gaming", "Laptop dengan RTX 4060", 15000000, 10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Produk berhasil ditambahkan dengan ID: %d\n", id)

	// 2. Ambil Semua Data
	products, err := service.GetAllProducts()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Daftar Produk:")
	for _, p := range products {
		fmt.Printf("- %s | Harga: %.2f | Stok: %d\n", p.ProductName, p.Price, p.StockQuantity)
	}
}
