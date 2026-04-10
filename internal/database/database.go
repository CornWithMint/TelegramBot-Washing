package database

import (
	"database/sql"
	"log/slog"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/CornWithMint/TelegramBot-Washing/config"

	"github.com/CornWithMint/TelegramBot-Washing/internal/entity"
)

type SqliteRepo struct {
	db *sql.DB
}

func NewSqliteRepo(cfg *config.Config) *SqliteRepo {
	slog.Debug("Запуск функции NewSqliteRepo")

	db, err := sql.Open("sqlite3", cfg.BdPath)
	if err != nil {
		slog.Error("Ошибка открытия БД", "error", err)
		os.Exit(1)
	}

	repo := &SqliteRepo{db: db}
	repo.CreateTable()

	slog.Debug("Завершение функции NewSqliteRepo")
	return repo
}

func (r *SqliteRepo) CreateTable() {
	slog.Debug("Запуск функции CreateTable")

	create := `CREATE TABLE IF NOT EXISTS Clothes (
		User_id INTEGER NOT NULL,
		Thing TEXT ,
		Color TEXT,
		Number INTEGER,
		UNIQUE(Thing, Color, Number)
	)`
	_, err := r.db.Exec(create)
	if err != nil {
		slog.Error("Ошибка создания таблицы", "error", err)
		os.Exit(1)
	}
	slog.Debug("Завершение функции CreateTable")
}

func (r *SqliteRepo) UpdateTable(u *entity.User, id int64) {
	slog.Debug("Запуск функции UpdateTable")

	insert := `INSERT INTO Clothes (User_id, Thing, Color, Number) VALUES (?, ?, ?, ?)`
	_, err := r.db.Exec(insert, id, u.Thing, u.Color, u.Number)
	if err != nil {
		slog.Warn("Не удалось вставить данные в таблицу", "warn", err)
	}
	//fmt.Printf("Вставлено %s с цветом %s \n", u.Thing, u.Color)
	slog.Debug("Завершение функции UpdateTable")
}

func (r *SqliteRepo) ReadValues(id int64) []entity.User {
	slog.Debug("Запуск функции ReadValues")

	get := `SELECT * FROM Clothes WHERE User_id = ?`
	rows, err := r.db.Query(get, id)
	if err != nil {
		slog.Warn("Не удалось", "warn", err)
	}

	defer rows.Close()

	res := make([]entity.User, 0)

	for rows.Next() {
		var number, id int
		var thing, color string
		err = rows.Scan(&id, &thing, &color, &number)
		if err != nil {
			slog.Warn("Не удалось прочитать данные", "warn", err)
		}
		res = append(res, entity.User{Thing: thing, Color: color, Number: number})
		//fmt.Printf("В БД Вещь: %s, Цвет: %s, Количество: %d\n", thing, color, number)
	}
	slog.Debug("Завершение функции ReadValues")
	return res
}

func (r *SqliteRepo) DeleteValues() {
	slog.Warn("Ничего не реализует")
}
