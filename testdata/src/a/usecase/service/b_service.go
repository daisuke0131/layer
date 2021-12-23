package service

import "a/infrastructure/repository" // want "NG"

type bService struct {
	aRepository         repository.ARepository
}

func (bService *bService) a(){
	bService.aRepository.AMethod()
}
