package usecase

import (
	"gitdeco-api/internal/deco"
	"gitdeco-api/internal/deco/dto"
	"gitdeco-api/internal/exception"
	"gitdeco-api/internal/models"
)

type DecoUsecase struct {
	decoRepository deco.Repository
}

func NewDecoUsecase(decoRepository deco.Repository) deco.Usecase {
	return &DecoUsecase{decoRepository: decoRepository}
}

func (du *DecoUsecase) Create(username string, request *dto.DecoRequest) (*models.Deco, error) {
	deco, err := du.decoRepository.Create(request.ToEntity(username))
	if err != nil {
		panic(&exception.Error{Key: "REPO_ERROR", Data: err.Error()})
	}
	return deco, nil
}

func (du *DecoUsecase) Read(ID uint) (*models.Deco, error) {
	deco, err := du.decoRepository.Read(ID)
	if err != nil {
		panic(&exception.Error{Key: "REPO_ERROR", Data: err.Error()})
	}
	return deco, nil
}

func (du *DecoUsecase) ReadAll() ([]*models.Deco, error) {
	decos, err := du.decoRepository.ReadAll()
	if err != nil {
		panic(&exception.Error{Key: "REPO_ERROR", Data: err.Error()})
	}
	return decos, nil
}

func (du *DecoUsecase) Update(ID uint, request *dto.DecoRequest) (*models.Deco, error) {
	deco, err := du.decoRepository.Read(ID)
	if err != nil {
		panic(&exception.Error{Key: "REPO_ERROR", Data: err.Error()})
	}
	return du.decoRepository.Update(deco.Update(request.Title, request.Markdown))
}

func (du *DecoUsecase) Delete(ID uint) error {
	deco, err := du.decoRepository.Read(ID)
	if err != nil {
		panic(&exception.Error{Key: "REPO_ERROR", Data: err.Error()})
	}
	return du.decoRepository.Delete(deco)
}
