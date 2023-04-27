package store

import "github.com/RicardoIvan-CM/Practicas-BDD/Practica5/Productos/internal/domain"

type StoreInterface interface {
	//ReadAll devuelve todos los productos
	ReadAll() ([]domain.Product, error)
	// Read devuelve un producto por su id
	Read(id int) (domain.Product, error)
	// Create agrega un nuevo producto
	Create(product domain.Product) error
	// Update actualiza un producto
	Update(product domain.Product) error
	// Delete elimina un producto
	Delete(id int) error
	// Exists verifica si un producto existe
	Exists(codeValue string) bool
}
