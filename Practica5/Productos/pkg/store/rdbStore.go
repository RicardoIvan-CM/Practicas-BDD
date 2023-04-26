package store

import (
	"database/sql"
	"errors"

	"github.com/RicardoIvan-CM/Practicas-BDD/Practica5/Productos/internal/domain"
	"github.com/go-sql-driver/mysql"
)

type rdbStore struct {
	db *sql.DB
}

func NewRDBStore(db sql.DB) StoreInterface {
	return &rdbStore{
		db: &db,
	}
}

var (
	ErrNotFound      = errors.New("The requested product was not found")
	ErrAlreadyExists = errors.New("The product already exists")
)

func (s *rdbStore) Read(id int) (domain.Product, error) {
	var product domain.Product
	row := s.db.QueryRow("select * from products where id = ?", id)
	err := row.Scan(&product.Id, &product.Name, &product.Quantity, &product.CodeValue, &product.IsPublished, &product.Expiration, &product.Price)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Product{}, ErrNotFound
		}
	}
	return product, nil
}

func (s *rdbStore) Create(product domain.Product) error {
	stmt, err := s.db.Prepare(`
		insert into products
		(id, name, quantity, code_value, is_published, expiration, price)
		values
		((select max(id)+1 from products prods),?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(
		product.Name,
		product.Quantity,
		product.CodeValue,
		product.IsPublished,
		product.Expiration,
		product.Price,
	)
	if err != nil {
		//Castear a error de MySQL
		mysqlError, ok := err.(*mysql.MySQLError)
		if !ok {
			return err
		}
		switch mysqlError.Number {
		case 1062, 1586:
			return ErrAlreadyExists
		}
		return err
	}

	//Mapear resultado
	lastID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	product.Id = int(lastID)
	return nil
}

func (s *rdbStore) Update(product domain.Product) error {
	stmt, err := s.db.Prepare(`
		update products set name = ?, quantity = ?, expiration = ?, price = ? where id = ?
	`)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Quantity, product.Expiration, product.Price, product.Id)
	if err != nil {
		return err
	}
	return nil
}

func (s *rdbStore) Delete(id int) error {
	stmt, err := s.db.Prepare(`
		delete from products where id = ?
	`)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *rdbStore) Exists(codeValue string) bool {
	rows, err := s.db.Query("select id from products where code_value = ?", codeValue)
	defer rows.Close()
	ids := make([]int, 0)

	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			panic(err)
		}
		ids = append(ids, id)
	}

	if rows.Err(); err != nil {
		panic(err)
	}

	if len(ids) > 0 {
		return true
	}
	return false
}
