package service

import (
	"change_money/dto"
	"change_money/errs"
	"change_money/repo"
)

type UserService interface {
	ExchangeService(dto.UserResquestExchange) (dto.UserResponseExchange, *errs.AppError)
	ReceiveService(dto.UserResquestReceive) (dto.UserResponseReceive, *errs.AppError)
}

type DefaultUserService struct {
	repo repo.UserRepo
}

func NewUserService(repo repo.UserRepo) UserService {
	return DefaultUserService{
		repo: repo,
	}
}

// func (s DefaultUserService) SendBSCToContract(req dto.UserResquestExchange) (dto.UserResponseExchange, *errs.AppError) {
// 	response, err := s.repo.SendBSCToContract(req)
// 	if err != nil {

// 		return dto.UserResponseExchange{}, err
// 	}
// 	return response, nil
// }

func (s DefaultUserService) ExchangeService(req dto.UserResquestExchange) (dto.UserResponseExchange, *errs.AppError) {
	response, err := s.repo.ExchangeController(req)
	if err != nil {
		return dto.UserResponseExchange{}, err
	}
	return response, nil
}

func (s DefaultUserService) ReceiveService(req dto.UserResquestReceive) (dto.UserResponseReceive, *errs.AppError) {
	response, err := s.repo.ReceiveController(req)
	if err != nil {

		return dto.UserResponseReceive{}, err
	}
	return response, nil
}

// func (s DefaultUserService) ReceiveBSC(req dto.UserResquestReceive) (dto.UserResponseReceive, *errs.AppError) {

// 	response, err := s.repo.ReceiveBSC(req)
// 	if err != nil {

// 		return dto.UserResponseReceive{}, err
// 	}
// 	return response, nil
// }
