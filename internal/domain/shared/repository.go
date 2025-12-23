package shared

type Repository interface {
	GetByID(id string) (BaseEntity, error)
	GetAll() ([]BaseEntity, error)
	Create(entity BaseEntity) error
	Update(entity BaseEntity) error
	Delete(id string) error
}
