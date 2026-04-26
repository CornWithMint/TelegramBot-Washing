package database

import (
	"database/sql"
	"log/slog"
	"os"

	"github.com/CornWithMint/TelegramBot-Washing/config"
	"github.com/CornWithMint/TelegramBot-Washing/internal/entity"
	_ "github.com/mattn/go-sqlite3"
)

type SqliteRepo struct {
	db *sql.DB
}

func NewSqliteRepo(cfg *config.Config, db *sql.DB) *SqliteRepo {
	slog.Debug("Запуск функции NewSqliteRepo")

	repo := &SqliteRepo{db: db}
	repo.CreateTable()

	slog.Debug("Завершение функции NewSqliteRepo")
	return repo
}

func (r *SqliteRepo) CreateTable() {
	slog.Debug("Запуск функции CreateTable")

	// del := `DROP TABLE IF EXISTS Clothes`

	create := `CREATE TABLE IF NOT EXISTS Clothes (
		Thing_id INTEGER NOT NULL,
		Thing TEXT ,
		Color TEXT,
		Number INTEGER,
		DateOfWashing  TEXT,
		UNIQUE(Thing, Color, Number)
	)`

	_, err := r.db.Exec(create)
	if err != nil {
		slog.Error("Ошибка создания таблицы", "error", err)
		os.Exit(1)
	}
	slog.Debug("Завершение функции CreateTable")
}

func (r *SqliteRepo) InsertTable(u *entity.Thing, id int64) {
	slog.Debug("Запуск функции InsertTable")

	// !!!!! ВМЕСТО ИНСЕРТ ИСПОЛЬЗОВАТЬ UPDATE
	insert := `INSERT INTO Clothes (Thing_id, Thing, Color, Number, DateOfWashing) VALUES (?, ?, ?, ?, ?)`
	_, err := r.db.Exec(insert, id, u.Thing, u.Color, u.Number, u.DateOfWashing)
	if err != nil {
		slog.Warn("Не удалось вставить данные в таблицу", "warn", err)
	}
	//fmt.Printf("Вставлено %s с цветом %s \n", u.Thing, u.Color)
	slog.Debug("Завершение функции InsertTable")
}

func (r *SqliteRepo) UpdateTable(u *entity.Thing, id int64) {
	update := `UPDATE Clothes SET DateOfWashing = ? WHERE Thing_id = ? AND Thing = ? AND Color = ?`

	_, err := r.db.Exec(update, entity.TimeNow, id, u.Thing, u.Color)
	if err != nil {
		slog.Warn("Не удалось вставить данные в таблицу", "warn", err)
	}
}

func (r *SqliteRepo) ReadValues(id int64) []entity.Thing {
	slog.Debug("Запуск функции ReadValues")

	get := `SELECT * FROM Clothes WHERE Thing_id = ?`
	rows, err := r.db.Query(get, id)
	if err != nil {
		slog.Warn("Не удалось получить данные из БД", "warn", err)
	}

	defer rows.Close()

	res := make([]entity.Thing, 0)

	for rows.Next() {
		var number, id int
		var thing, color, DateOfWashing string
		err = rows.Scan(&id, &thing, &color, &number, &DateOfWashing)

		if err != nil {
			slog.Warn("Не удалось прочитать данные", "warn", err)
		}
		res = append(res, entity.Thing{Thing: thing, Color: color, Number: number, DateOfWashing: DateOfWashing})
		//fmt.Printf("В БД Вещь: %s, Цвет: %s, Количество: %d\n", thing, color, number)
	}
	slog.Debug("Завершение функции ReadValues")
	return res
}

func (r *SqliteRepo) DeleteValues(u *entity.Thing, id int64) {
	del := `DELETE FROM Clothes WHERE Thing_id = ? AND Thing = ? AND Color = ?`
	_, err := r.db.Exec(del, id, u.Thing, u.Color)
	if err != nil {
		slog.Warn("Не удалось удалить данные из таблицы", "warn", err)
	}
}
