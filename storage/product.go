package storage

import (
	"database/sql"
	"fmt"

	"apiv0.1/model"
)

/*
!
*/
type scanner interface {
	Scan(dest ...interface{}) error
}

const (
	MigrateTable = `
		CREATE TABLE IF NOT EXISTS products(
		id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(100) NOT NULL,
		body TEXT NOT NULL,
		create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		update_at TIMESTAMP NULL
		)ENGINE = InnoDB;
	
	`
	CreateProduct    = `INSERT INTO products(title,body) VALUES (?,?);`
	QueryProduct     = `SELECT * FROM products;`
	QueryByIdProduct = "SELECT * FROM products Where id = ?;"
	UpdateProduct    = `UPDATE products SET title = ?, body =?, update_at=? WHERE id= ?;`
	DeleteProduct    = `DELETE FROM products WHERE id = ?;`
)

type MySQLProduct struct {
	db *sql.DB
}

func NewMySQL(db *sql.DB) *MySQLProduct {
	return &MySQLProduct{db}
}

func (m *MySQLProduct) Migrate() error {
	stm, err := m.db.Prepare(MigrateTable)
	if err != nil {
		return err
	}
	defer stm.Close()
	_, err = stm.Exec()
	if err != nil {
		return err
	}
	fmt.Println("Procees Create table was successfully!!")
	return nil
}

func (m *MySQLProduct) Create(p *model.Product) error {

	stm, err := db.Prepare(CreateProduct)
	if err != nil {
		return err
	}
	defer stm.Close()

	result, err := stm.Exec(
		p.Title,
		p.Body,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	p.ID = uint(id)
	fmt.Println("Data was inserted successfully!!")
	return nil

}

func (m *MySQLProduct) Update(p *model.Product) error {
	stm, err := db.Prepare(UpdateProduct)
	if err != nil {
		return err
	}
	defer stm.Close()
	result, err := stm.Exec(
		p.Title,
		p.Body,
		p.Update_At,
		p.ID,
	)
	if err != nil {
		return err
	}
	rowsAf, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAf == 0 {
		return fmt.Errorf("the field not exists witht id: %d", p.ID)
	}
	fmt.Println("Data was updated successfully!!")
	return nil
}

func (m *MySQLProduct) Delete(id uint) error {
	stm, err := db.Prepare(DeleteProduct)
	if err != nil {
		return err
	}
	defer stm.Close()
	_, err = stm.Exec(id)
	if err != nil {
		return err
	}
	fmt.Println("Data was destroy of the database successfully!!")
	return nil
}

func (m *MySQLProduct) GetAll() (model.Products, error) {
	stm, err := db.Prepare(QueryProduct)
	if err != nil {
		return nil, err
	}
	defer stm.Close()

	rows, err := stm.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	pro := make(model.Products, 0)
	for rows.Next() {
		m, err := scanRowProduct(rows)
		if err != nil {
			return nil, err
		}
		pro = append(pro, m)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	fmt.Println("Query data was successfully!!")
	return pro, nil
}

func (m *MySQLProduct) GetById(id uint) (*model.Product, error) {
	stm, err := db.Prepare(QueryByIdProduct)
	if err != nil {
		return &model.Product{}, err
	}
	defer stm.Close()
	return scanRowProduct(stm.QueryRow(id))
}

/*
? helper for scanRow
*/
func scanRowProduct(s scanner) (*model.Product, error) {

	m := &model.Product{}
	err := s.Scan(
		&m.ID,
		&m.Title,
		&m.Body,
		&m.Create_At,
		&m.Update_At,
	)

	if err != nil {
		return &model.Product{}, err
	}

	return m, nil

}
