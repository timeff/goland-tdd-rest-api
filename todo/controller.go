package todo

type Controller interface {
	Get() ([]*Todo, error)
	Create(*Todo) (int64, error)
	Update(*Todo) error
	Delete(int64) error
}
