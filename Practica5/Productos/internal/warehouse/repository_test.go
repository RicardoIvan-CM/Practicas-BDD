package warehouse

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	"github.com/RicardoIvan-CM/Practicas-BDD/Practica5/Productos/internal/domain"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func init() {
	dsn := mysql.Config{
		User:   "user1",
		Passwd: "secret_password",
		Addr:   "127.0.0.1:3306",
		DBName: "my_db",
	}
	txdb.Register("txdb", "mysql", dsn.FormatDSN())
}

func TestRepository_GetAll(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		//Arrange
		db, err := sql.Open("txdb", "indentifier")
		assert.NoError(t, err)
		defer db.Close()

		expectedResult := []domain.Warehouse{
			{
				Id:        1,
				Name:      "Main Warehouse",
				Address:   "221 Baker Street",
				Telephone: "4555666",
				Capacity:  100,
			},
			{
				Id:        2,
				Name:      "Other Warehouse",
				Address:   "123 Fake Street",
				Telephone: "5511223344",
				Capacity:  300,
			},
		}

		rp := NewRepository(db)
		result, err := rp.GetAll()
		assert.NoError(t, err)
		assert.Equal(t, expectedResult, result)
	})
}

func TestRepository_GetByID(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		//Arrange
		id := 1
		db, err := sql.Open("txdb", "indentifier")
		assert.NoError(t, err)
		defer db.Close()

		expectedResult := domain.Warehouse{
			Id:        1,
			Name:      "Main Warehouse",
			Address:   "221 Baker Street",
			Telephone: "4555666",
			Capacity:  100,
		}

		rp := NewRepository(db)
		result, err := rp.GetByID(id)
		assert.NoError(t, err)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("Not Found", func(t *testing.T) {
		//Arrange
		id := 7777
		db, err := sql.Open("txdb", "indentifier")
		assert.NoError(t, err)
		defer db.Close()

		expectedResult := domain.Warehouse{}

		rp := NewRepository(db)
		result, err := rp.GetByID(id)
		assert.ErrorIs(t, err, ErrNotFound)
		assert.Equal(t, expectedResult, result)
	})
}

func TestRepository_Create(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		//Arrange
		db, err := sql.Open("txdb", "indentifier")
		assert.NoError(t, err)
		defer db.Close()

		warehouseToInsert := domain.Warehouse{
			Name:      "My Warehouse",
			Address:   "123 Mock Street",
			Telephone: "11223344",
			Capacity:  150,
		}

		expectedResult := domain.Warehouse{
			Id:        3,
			Name:      "My Warehouse",
			Address:   "123 Mock Street",
			Telephone: "11223344",
			Capacity:  150,
		}

		rp := NewRepository(db)
		result, err := rp.Create(warehouseToInsert)
		assert.NoError(t, err)
		assert.Equal(t, expectedResult, result)
	})
}
