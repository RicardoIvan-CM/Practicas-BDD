package warehouse

import "github.com/RicardoIvan-CM/Practicas-BDD/Practica5/Productos/internal/domain"

type Service interface {
	//GetAll trae todos los warehouses
	GetAll() ([]domain.Warehouse, error)
	// GetByID busca un warehouse por su id
	GetByID(id int) (domain.Warehouse, error)
	//GetProductsByID trae los productos de un warehouse a partir de su id
	GetProductsByID(id int) ([]domain.Product, error)
	// Create agrega un nuevo warehouse
	Create(p domain.Warehouse) (domain.Warehouse, error)
	// Delete elimina un warehouse
	Delete(id int) error
	// Update actualiza un warehouse
	Update(id int, p domain.Warehouse) (domain.Warehouse, error)
}

type service struct {
	r Repository
}

// NewService crea un nuevo servicio
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetAll() ([]domain.Warehouse, error) {
	ws, err := s.r.GetAll()
	if err != nil {
		return []domain.Warehouse{}, err
	}
	return ws, nil
}

func (s *service) GetByID(id int) (domain.Warehouse, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Warehouse{}, err
	}
	return p, nil
}

func (s *service) GetProductsByID(id int) ([]domain.Product, error) {
	products, err := s.r.GetProductsByID(id)
	if err != nil {
		return []domain.Product{}, err
	}
	return products, nil
}

func (s *service) Create(p domain.Warehouse) (domain.Warehouse, error) {
	p, err := s.r.Create(p)
	if err != nil {
		return domain.Warehouse{}, err
	}
	return p, nil
}
func (s *service) Update(id int, u domain.Warehouse) (domain.Warehouse, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Warehouse{}, err
	}
	if u.Name != "" {
		p.Name = u.Name
	}

	if u.Address != "" {
		p.Address = u.Address
	}
	if u.Telephone != "" {
		p.Telephone = u.Telephone
	}
	if u.Capacity > 0 {
		p.Capacity = u.Capacity
	}
	p, err = s.r.Update(id, p)
	if err != nil {
		return domain.Warehouse{}, err
	}
	return p, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
