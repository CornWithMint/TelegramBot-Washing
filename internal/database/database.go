package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/CornWithMint/TelegramBot-Washing/config"

	"github.com/CornWithMint/TelegramBot-Washing/internal/entity"
)

type SqliteRepo struct {
	db *sql.DB
}

func NewSqliteRepo(cfg *config.Config) (*SqliteRepo, error) {
	db, err := sql.Open("sqlite3", cfg.BdPath)
	if err != nil {
		return nil, fmt.Errorf("Ошибка открытия бд", err)

	}

	repo := &SqliteRepo{db: db}
	repo.CreateTable()

	return repo, nil
}

func (r *SqliteRepo) CreateTable() {
	create := `CREATE TABLE IF NOT EXISTS Clothes (
		User_id INTEGER NOT NULL,
		Thing TEXT ,
		Color TEXT,
		Number INTEGER,
		UNIQUE(Thing, Color, Number)
	)`
	_, err := r.db.Exec(create)
	if err != nil {
		log.Fatal("Ошибка создания бд", err)
	}
	fmt.Println("Таблица создана")
}

func (r *SqliteRepo) UpdateTable(u *entity.User, id int64) {
	insert := `INSERT INTO Clothes (User_id, Thing, Color, Number) VALUES (?, ?, ?, ?)`
	_, err := r.db.Exec(insert, id, u.Thing, u.Color, u.Number)
	if err != nil {
		log.Fatal("Ошибка вставки данных", err)
	}
	fmt.Printf("Вставлено %s с цветом %s \n", u.Thing, u.Color)
}

func (r *SqliteRepo) ReadValues(id int64) []entity.User {
	get := `SELECT * FROM Clothes WHERE User_id = ?`
	rows, err := r.db.Query(get, id)
	if err != nil {
		log.Fatal("", err)
	}

	defer rows.Close()

	res := make([]entity.User, 0)

	for rows.Next() {
		var number, id int
		var thing, color string
		err = rows.Scan(&id, &thing, &color, &number)
		if err != nil {
			log.Fatal("Ошибка сканирования строки:", err)
		}
		res = append(res, entity.User{Thing: thing, Color: color, Number: number})
		fmt.Printf("В БД Вещь: %s, Цвет: %s, Количество: %d\n", thing, color, number)

	}
	return res
}

func (r *SqliteRepo) DeleteValues() {
}
