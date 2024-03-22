package cours

import (
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/education"
)

var entitiesPrimaryKey = 0
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
	for index, cours := range allEntities {
		if cours.ID == int(coursID) {
			return &allEntities[index], nil
		}
	}

	return nil, fmt.Errorf("❌ Invalid ID")
}

func (s *DummyCourseService) Create(cours education.Cours) (uint64, error) {
	cours.ID = s.generateID()

	allEntities = append(allEntities, cours)

	return uint64(cours.ID), nil
}

func (s *DummyCourseService) Update(coursID uint64, cours education.Cours) error {
	for index, entity := range allEntities {
		if entity.ID == int(coursID) {
			allEntities[index] = cours
			return nil
		}
	}

	return fmt.Errorf("❌ Invalid ID")
}

func (s *DummyCourseService) Remove(coursID uint64) (bool, error) {
	for index, cours := range allEntities {
		if cours.ID == int(coursID) {
			allEntities = append(allEntities[:index], allEntities[index+1:]...)
			return true, nil
		}
	}

	return false, fmt.Errorf("❌ Invalid ID")
}

func (s *DummyCourseService) GetCoursesCount() int {
	return len(allEntities)
}

func (s *DummyCourseService) generateID() int {
	entitiesPrimaryKey += 1
	return entitiesPrimaryKey
}
