package position

type IPositionRepo interface {
	Save(p Position) error
}
