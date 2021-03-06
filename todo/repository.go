package todo

type Repository interface {
	Get() ([]*Todo, error)
	GetByID(int64) (*Todo, error)
	Create(*Todo) (int64, error)
	Update(*Todo) error
	Delete(int64) error
}
