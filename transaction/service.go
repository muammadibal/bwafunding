package transaction

import (
	"bwafunding/campaign"
	"errors"
)

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
}

type Service interface {
	TransactionsByCampaignID(input TransactionDetailInput) ([]Transaction, error)
	TransactionsByUserID(userID int) ([]Transaction, error)
}

func AssignService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository, campaignRepository}
}

func (s *service) TransactionsByCampaignID(input TransactionDetailInput) ([]Transaction, error) {
	campaign, err := s.campaignRepository.FindByID(input.ID)
	if err != nil {
		return []Transaction{}, err
	}

	if campaign.UserID != int32(input.User.ID) {
		return []Transaction{}, errors.New("Not an owner of the campaign")
	}

	transactions, err := s.repository.FindByID(input.ID)
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (s service) TransactionsByUserID(userID int) ([]Transaction, error) {
	transactions, err := s.repository.FindByUserID(userID)

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
