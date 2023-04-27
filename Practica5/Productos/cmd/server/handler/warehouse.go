package handler

import (
	"errors"
	"os"
	"strconv"

	"github.com/RicardoIvan-CM/Practicas-BDD/Practica5/Productos/internal/domain"
	"github.com/RicardoIvan-CM/Practicas-BDD/Practica5/Productos/internal/warehouse"
	"github.com/RicardoIvan-CM/Practicas-BDD/Practica5/Productos/pkg/web"
	"github.com/gin-gonic/gin"
)

type warehouseHandler struct {
	s warehouse.Service
}

// NewWarehouseHandler crea un nuevo controller de productos
func NewWarehouseHandler(s warehouse.Service) *warehouseHandler {
	return &warehouseHandler{
		s: s,
	}
}

// GetAll obtiene todos los productos
func (h *warehouseHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		warehouses, err := h.s.GetAll()
		if err != nil {
			web.Failure(c, 500, err)
			return
		}
		web.Success(c, 200, warehouses)
	}
}

// Get obtiene un producto por id
func (h *warehouseHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		product, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("warehouse not found"))
			return
		}
		web.Success(c, 200, product)
	}
}

func (h *warehouseHandler) GetProductsByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Query("id")
		id, err := strconv.Atoi(idParam)
		if err != nil || id == 0 {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		product, err := h.s.GetProductsByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("warehouse not found"))
			return
		}
		web.Success(c, 200, product)
	}
}

// validateEmptys valida que los campos no esten vacios
func validateEmptysWarehouse(warehouse *domain.Warehouse) (bool, error) {
	switch {
	case warehouse.Name == "" || warehouse.Address == "" || warehouse.Telephone == "":
		return false, errors.New("fields can't be empty")
	case warehouse.Capacity <= 0:
		return false, errors.New("quantity must be greater than 0")
	}
	return true, nil
}

// Post crea un nuevo producto
func (h *warehouseHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var warehouse domain.Warehouse
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
		err := c.ShouldBindJSON(&warehouse)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptysWarehouse(&warehouse)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.Create(warehouse)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p)
	}
}

// Delete elimina un producto
func (h *warehouseHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 204, nil)
	}
}

// Put actualiza un producto
func (h *warehouseHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("warehouse not found"))
			return
		}
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		var warehouse domain.Warehouse
		err = c.ShouldBindJSON(&warehouse)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptysWarehouse(&warehouse)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.Update(id, warehouse)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}

// Patch actualiza un producto o alguno de sus campos
func (h *warehouseHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Name      string  `json:"name,omitempty"`
		Address   string  `json:"address,omitempty"`
		Telephone string  `json:"telephone,omitempty"`
		Capacity  float64 `json:"capacity,omitempty"`
	}
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
		var r Request
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("warehouse not found"))
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		update := domain.Warehouse{
			Name:      r.Name,
			Address:   r.Address,
			Telephone: r.Telephone,
			Capacity:  int(r.Capacity),
		}
		p, err := h.s.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}
