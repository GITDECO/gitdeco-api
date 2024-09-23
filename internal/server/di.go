package server

import (
	"gitdeco-api/internal/auth"
	ad "gitdeco-api/internal/auth/delivery"
	ar "gitdeco-api/internal/auth/repository"
	au "gitdeco-api/internal/auth/usecase"
	"gitdeco-api/internal/deco"
	dd "gitdeco-api/internal/deco/delivery"
	dr "gitdeco-api/internal/deco/repository"
	du "gitdeco-api/internal/deco/usecase"
	"gitdeco-api/internal/svg"
	sd "gitdeco-api/internal/svg/delivery"
	su "gitdeco-api/internal/svg/usecase"
	"gitdeco-api/internal/user"
	ud "gitdeco-api/internal/user/delivery"
	ur "gitdeco-api/internal/user/repository"
	uu "gitdeco-api/internal/user/usecase"

	"gorm.io/gorm"
)

func AuthDI(db *gorm.DB) auth.Handler {
	repo := ar.NewAuthRepository(db)
	usecase := au.NewAuthUsecase(repo)
	handler := ad.NewAuthHandler(usecase)
	return handler
}

func UserDI(db *gorm.DB) user.Handler {
	repo := ur.NewUserRepository(db)
	usecase := uu.NewUserUsecase(repo)
	handler := ud.NewUserHandler(usecase)
	return handler
}

func SvgDI(db *gorm.DB) svg.Handler {
	usecase := su.NewSvgUsecase()
	handler := sd.NewSvgHandler(usecase)
	return handler
}

func DecoDI(db *gorm.DB) deco.Handler {
	repo := dr.NewDecoRepository(db)
	usecase := du.NewDecoUsecase(repo)
	handler := dd.NewDecoHandler(usecase)
	return handler
}
