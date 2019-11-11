package position

type IPositionService interface {
	Save(p Position) error
}

type positionService struct {
	repo IPositionRepo
}

func NewPositionService(repo IPositionRepo) IPositionService {
	return &positionService{
		repo,
	}
}

func (ps positionService) Save(p Position) error {
	return ps.repo.Save(p)
}
