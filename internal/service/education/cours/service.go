package cours

import (
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/education"
)

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
	if len(allEntities) == 0 {
		return []education.Cours{}, nil
	}

	from := cursor * limit
	if from > uint64(len(allEntities)) {
		return nil, fmt.Errorf("❌ Invalid data for paginate")
	}

	to := from + limit
	if to > uint64(len(allEntities)) {
		to = uint64(len(allEntities))
	}

	return allEntities[from:to], nil
}

func (s *DummyCourseService) Describe(coursID uint64) (*education.Cours, error) {
	if int(coursID) > len(allEntities)-1 {
		return nil, fmt.Errorf("❌ Invalid cours ID")
	}

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
	if int(coursID) > len(allEntities)-1 {
		return false, fmt.Errorf("❌ Invalid cours ID")
	}

	allEntities = append(allEntities[:coursID], allEntities[coursID+1:]...)

	return true, nil
}

func (s *DummyCourseService) GetCoursesCount() int {
	return len(allEntities)
}
