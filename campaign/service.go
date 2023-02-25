package campaign

import "errors"

type Service interface {
	Campaigns(userID int) ([]Campaign, error)
	CampaignDetail(ID CampaignDetailInput) (Campaign, error)
}

type service struct {
	repository Repository
}

func AssignService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Campaigns(userID int) ([]Campaign, error) {
	if userID != 0 {
		campaigns, err := s.repository.FindByUserID(userID)
		if err != nil {
			return campaigns, err
		}

		return campaigns, nil
	}

	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (s *service) CampaignDetail(input CampaignDetailInput) (Campaign, error) {
	campaign, err := s.repository.FindByID(input.ID)
	if err != nil {
		return campaign, err
	}

	if campaign.ID == 0 {
		return campaign, errors.New("Campaign doesn't exist")
	}

	return campaign, nil
}
