package services

import (
	"github.com/Parovozzzik/real-estate-portfolio/internal/repositories"
	"github.com/Parovozzzik/real-estate-portfolio/internal/utils"
)

type EstateValuesService struct {
	estateValueRepository *repositories.EstateValueRepository
}

func NewEstateValuesService(
	estateValueRepository *repositories.EstateValueRepository,
) *EstateValuesService {
	return &EstateValuesService{
		estateValueRepository: estateValueRepository,
	}
}

type RecalculateEstateValues struct {
	EstateId  int64            `json:"estate_id", db:"estate_id"`
	DateStart utils.CustomTime `json:"date_start", db:"date_start"`
	DateEnd   utils.CustomTime `json:"date_start", db:"date_start"`
}

func (s *EstateValuesService) Recalculate(recalculateEstateValues *RecalculateEstateValues) {

}
