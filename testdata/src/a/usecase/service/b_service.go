package service

import "a/infrastructure/repository"

type bService struct {
	aRepository         repository.ARepository
}

func (bService *bService) a(){
	bService.aRepository.AMethod()
}
