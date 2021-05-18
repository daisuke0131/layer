package service

import (
	"a/usecase/repository"
)

type aService struct {
	aRepository         repository.ARepository
}

func (aService *aService)a(){
	aService.aRepository.AMethod()
}