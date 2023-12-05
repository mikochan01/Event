package entity

import (
	"fmt"
	"time"
)

// Event adalah struktur untuk menyimpan detail event
type Event struct {
	Nama       string
	Tanggal    time.Time
	Venue      string
	Harga      float64
	Keterangan string
}

// Membuat metode untuk mencetak detail event
func (e Event) PrintEventDetails() {
	fmt.Printf("Nama Event: %s\n", e.Nama)
	fmt.Printf("Hari dan Tanggal: %s\n", e.Tanggal.Format("Monday, 2 January 06"))
	fmt.Printf("Venue: %s\n", e.Venue)
	fmt.Printf("Harga: IDR %.0f\n", e.Harga)
	fmt.Printf("Keterangan: %s\n", e.Keterangan)
	fmt.Println("---------------------------------------------------")
}
