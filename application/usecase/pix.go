package usecase

import (
	"github.com/leoluzh/codepix-go/domain/model"
)

type PixUseCase struct {
	PixKeyRepository model.PixKeyRepository
}

func (p *PixUseCase) RegisterKey(key string, kind string, accountId string) (*model.PixKey, error) {
	account, err := p.PixKeyRepository.FindAccount(key)
	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(kind, account, key)

	if err != nil {
		return nil, err
	}

	_, err = p.PixKeyRepository.RegisterKey(pixKey)

	if pixKey.ID == "" {
		return nil, err
	}

	return pixKey, nil
}

func (p *PixUseCase) FindKey(key string, kind string) (*model.PixKey, error) {
	pixKey, err := p.PixKeyRepository.FindKeyByKind(key, kind)
	if err != nil {
		return nil, err
	}
	return pixKey, nil
}
