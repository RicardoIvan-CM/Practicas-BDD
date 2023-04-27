package product

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	"github.com/RicardoIvan-CM/Practicas-BDD/Practica5/Productos/internal/domain"
	"github.com/RicardoIvan-CM/Practicas-BDD/Practica5/Productos/pkg/store"
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

		store := store.NewRDBStore(db)
		rp := NewRepository(store)
		result, err := rp.GetAll()
		assert.NoError(t, err)
		assert.IsType(t, []domain.Product{}, result)
	})
}
