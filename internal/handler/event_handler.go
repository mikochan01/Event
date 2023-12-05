package handler

import (
	"Event/internal/entity"
	"fmt"
	"time"
)

func StartServer() {
	events := []entity.Event{
		{
			Nama:       "Konser Maher Zain",
			Tanggal:    time.Date(2022, time.June, 22, 0, 0, 0, 0, time.UTC),
			Venue:      "Bandung",
			Harga:      1299999,
			Keterangan: "",
		},
		{
			Nama:       "Festival Musik 2023",
			Tanggal:    time.Date(2023, time.April, 2, 0, 0, 0, 0, time.UTC),
			Venue:      "Semarang",
			Harga:      1099999,
			Keterangan: "",
		},
		{
			Nama:       "Charity Concert",
			Tanggal:    time.Date(2023, time.July, 31, 0, 0, 0, 0, time.UTC),
			Venue:      "Manado",
			Harga:      800000,
			Keterangan: "",
		},
	}

	for _, event := range events {
		printEventDetails(event)
	}
}

func printEventDetails(e entity.Event) {
	fmt.Printf("Nama Event: %s\n", e.Nama)
	fmt.Printf("Hari dan Tanggal: %s\n", e.Tanggal.Format("Monday, 2 January 06"))
	fmt.Printf("Venue: %s\n", e.Venue)
	fmt.Printf("Harga: IDR %.0f\n", e.Harga)
	fmt.Printf("Keterangan: %s\n", e.Keterangan)
	fmt.Println("---------------------------------------------------")
}
