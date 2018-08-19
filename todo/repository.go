package todo

type Repository interface {
	Get() ([]*Todo, error)
	Create(*Todo) (int64, error)
	Update(*Todo) (int64, error)
	Delete(id int64) error
}
