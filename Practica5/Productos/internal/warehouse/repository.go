package warehouse

import (
	"database/sql"
	"errors"

	"github.com/RicardoIvan-CM/Practicas-BDD/Practica5/Productos/internal/domain"
	"github.com/go-sql-driver/mysql"
)

type Repository interface {
	//GetAll trae todos los warehouses
	GetAll() ([]domain.Warehouse, error)
	// GetByID busca un warehouse por su id
	GetByID(id int) (domain.Warehouse, error)
	//GetProductsByID trae los productos de un warehouse a partir de su id
	GetProductsByID(id int) ([]domain.Product, error)
	// Create agrega un nuevo warehouse
	Create(p domain.Warehouse) (domain.Warehouse, error)
	// Update actualiza un warehouse
	Update(id int, p domain.Warehouse) (domain.Warehouse, error)
	// Delete elimina un warehouse
	Delete(id int) error
}

type repository struct {
	db *sql.DB
}

// NewRepository crea un nuevo repositorio
func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

var (
	ErrNotFound      = errors.New("The requested warehouse was not found")
	ErrAlreadyExists = errors.New("The warehouse already exists")
)

func (r *repository) GetAll() ([]domain.Warehouse, error) {
	rows, err := r.db.Query("select * from warehouses")
	if err != nil {
		return []domain.Warehouse{}, err
	}
	warehouses := make([]domain.Warehouse, 0)
	for rows.Next() {
		var warehouse domain.Warehouse
		err := rows.Scan(&warehouse.Id, &warehouse.Name, &warehouse.Address, &warehouse.Telephone, &warehouse.Capacity)
		if err != nil {
			return []domain.Warehouse{}, err
		}
		warehouses = append(warehouses, warehouse)
	}
	return warehouses, nil
}

func (r *repository) GetByID(id int) (domain.Warehouse, error) {
	var warehouse domain.Warehouse
	row := r.db.QueryRow("select * from warehouses where id = ?", id)
	err := row.Scan(&warehouse.Id, &warehouse.Name, &warehouse.Address, &warehouse.Telephone, &warehouse.Capacity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Warehouse{}, ErrNotFound
		}
	}
	return warehouse, nil
}

func (r *repository) GetProductsByID(id int) ([]domain.Product, error) {
	rows, err := r.db.Query("select p.id, p.name, p.quantity, p.code_value, p.is_published, p.expiration, p.price, p.id_warehouse from products p inner join warehouses w on p.id_warehouse = w.id where w.id = ?", id)
	if err != nil {
		return []domain.Product{}, err
	}
	products := make([]domain.Product, 0)
	for rows.Next() {
		var p domain.Product
		err = rows.Scan(&p.Id, &p.Name, &p.Quantity, &p.CodeValue, &p.IsPublished, &p.Expiration, &p.Price, &p.WarehouseId)
		if err != nil {
			return []domain.Product{}, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (r *repository) Create(warehouse domain.Warehouse) (domain.Warehouse, error) {
	stmt, err := r.db.Prepare(`
		insert into warehouses
		(name, address, telephone, capacity)
		values
		(?, ?, ?, ?)`)
	if err != nil {
		return domain.Warehouse{}, err
	}

	defer stmt.Close()
	result, err := stmt.Exec(
		warehouse.Name,
		warehouse.Address,
		warehouse.Telephone,
		warehouse.Capacity,
	)
	if err != nil {
		//Castear a error de MySQL
		mysqlError, ok := err.(*mysql.MySQLError)
		if !ok {
			return domain.Warehouse{}, err
		}
		switch mysqlError.Number {
		case 1062, 1586:
			return domain.Warehouse{}, ErrAlreadyExists
		}
		return domain.Warehouse{}, err
	}

	//Mapear resultado
	lastID, err := result.LastInsertId()
	if err != nil {
		return domain.Warehouse{}, err
	}
	warehouse.Id = int(lastID)
	return warehouse, nil
}

func (r *repository) Update(id int, warehouse domain.Warehouse) (domain.Warehouse, error) {
	stmt, err := r.db.Prepare(`
		update warehouses set name = ?, address = ?, telephone = ?, capacity = ? where id = ?
	`)
	if err != nil {
		return domain.Warehouse{}, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(warehouse.Name, warehouse.Address, warehouse.Telephone, warehouse.Capacity, id)
	if err != nil {
		return domain.Warehouse{}, err
	}
	warehouse.Id = id
	return warehouse, nil
}

func (r *repository) Delete(id int) error {
	stmt, err := r.db.Prepare(`
		delete from warehouses where id = ?
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
