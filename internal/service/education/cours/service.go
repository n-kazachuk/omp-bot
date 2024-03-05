package cours

import "github.com/ozonmp/omp-bot/internal/model/education"

var allEntities = make([]education.Cours, 0)

type CoursService interface {
	Describe(coursID uint64) (*education.Cours, error)
	List(cursor uint64, limit uint64) ([]education.Cours, error)
	Create(cours education.Cours) (uint64, error)
	Update(coursID uint64, cours education.Cours) error
	Remove(coursID uint64) (bool, error)
}

type DummyCourseService struct{}

func NewDummyCourseService() *DummyCourseService {
	return &DummyCourseService{}
}

func (s *DummyCourseService) List(cursor uint64, limit uint64) ([]education.Cours, error) {
	return allEntities, nil
}

func (s *DummyCourseService) Describe(coursID uint64) (*education.Cours, error) {
	return &allEntities[coursID], nil
}

func (s *DummyCourseService) Create(cours education.Cours) (uint64, error) {
	allEntities = append(allEntities, cours)

	return uint64(len(allEntities) - 1), nil
}

func (s *DummyCourseService) Update(coursID uint64, cours education.Cours) error {
	allEntities[coursID] = cours
	return nil
}

func (s *DummyCourseService) Remove(coursID uint64) (bool, error) {
	allEntities = append(allEntities[:coursID], allEntities[coursID+1:]...)
	return true, nil
}
