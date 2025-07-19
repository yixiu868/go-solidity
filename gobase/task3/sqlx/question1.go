package sqlx

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/yixiu868/go-solidity/gobase/task3/sqlx/utils"
	"log"
)

type Employee struct {
	Id         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

func searchTechDepart(db *sqlx.DB, department string) ([]Employee, error) {
	var employees []Employee
	sqlStr := `SELECT * FROM employees WHERE department = ?`
	err := db.Select(&employees, sqlStr, department)
	if err != nil {
		return nil, err
	}
	return employees, nil
}

func queryHighestSalaryEmp(db *sqlx.DB) (*Employee, error) {
	var employee Employee
	sqlStr := `SELECT * FROM employees ORDER BY salary DESC limit 1;`
	err := db.Get(&employee, sqlStr)
	if err != nil {
		return nil, err
	}
	return &employee, nil
}

func Run() {
	err := utils.InitDB()
	if err != nil {
		log.Println("init db fail", err.Error())
		return
	}
	defer utils.Db.Close()

	employees, err := searchTechDepart(utils.Db, "技术部")
	if err != nil {
		log.Println("search db fail", err.Error())
		return
	}
	for _, employee := range employees {
		fmt.Printf("employee: %v\n", employee)
	}

	employee, err := queryHighestSalaryEmp(utils.Db)
	if err != nil {
		log.Println("search db fail", err.Error())
		return
	}
	fmt.Printf("employee: %v\n", *employee)
}
