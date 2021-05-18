package repository

import (
	"a/usecase/repository"
	"fmt"
)

type ARepository interface{
	AMethod()
}

type aRepository struct {
}

func NewARepository() repository.ARepository {
	return &aRepository{}
}

func (aRepository *aRepository) AMethod(){
	fmt.Println("test")
}
