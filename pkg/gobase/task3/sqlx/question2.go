package sqlx

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/yixiu868/go-solidity/pkg/gobase/task3/sqlx/utils"
	"log"
)

type Book struct {
	Id     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func searchBooksByCondition(db *sqlx.DB, title string, price float64) ([]Book, error) {
	var books []Book
	sqlStr := `SELECT * FROM books WHERE title LIKE ? AND price > ?`
	err := db.Select(&books, sqlStr, title, price)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func RunBook() {
	err := utils.InitDB()
	if err != nil {
		log.Println("init db fail", err.Error())
		return
	}
	defer utils.Db.Close()

	books, err := searchBooksByCondition(utils.Db, "%Go%", 50)
	if err != nil {
		log.Println("search db fail", err.Error())
		return
	}
	for _, book := range books {
		fmt.Println(book)
	}
}
