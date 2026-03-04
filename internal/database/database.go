package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/CornWithMint/TelegramBot-Washing/internal/entity"
)

func CreateTable(db *sql.DB) {
	create := `CREATE TABLE IF NOT EXISTS Clothes (
		User_id INTEGER,
		Thing TEXT,
		Color TEXT,
		Number INTEGER,
		UNIQUE(Thing, Color, Number)
	)`
	_, err := db.Exec(create)
	if err != nil {
		log.Fatal("Ошибка создания бд", err)
	}
	fmt.Println("Таблица создана")
}

func UpdateTable(db *sql.DB, u *entity.User) {
	insert := `INSERT OR IGNORE INTO Clothes (User_id, Thing, Color, Number) VALUES (?, ?, ?, ?)`
	_, err := db.Exec(insert, u.Id, u.Thing, u.Color, u.Number)
	if err != nil {
		log.Fatal("Ошибка вставки данных", err)
	}
}

func ReadValues(id int, db *sql.DB) {
	get := `SELECT * FROM Clothes WHERE User_id = ?`
	_, err := db.Query(get, id)
	if err != nil {
		log.Fatal("", err)
	}
}

func DeleteValues() {
}
