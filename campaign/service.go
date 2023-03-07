package campaign

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	CreateCampaign(input CampaignCreateInput) (Campaign, error)
	UpdateCampaign(ID CampaignDetailInput, input CampaignCreateInput) (Campaign, error)
	Campaigns(userID int) ([]Campaign, error)
	CampaignDetail(ID CampaignDetailInput) (Campaign, error)
	CreateCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error)
}

type service struct {
	repository Repository
}

func AssignService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateCampaign(input CampaignCreateInput) (Campaign, error) {
	campaign := Campaign{}
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.Perks = input.Perks
	campaign.GoalAmount = input.GoalAmount
	campaign.UserID = int32(input.User.ID)

	slugTitle := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	campaign.Slug = slug.Make(slugTitle)

	campaignData, err := s.repository.Save(campaign)
	if err != nil {
		return campaignData, err
	}
	return campaignData, nil
}

func (s *service) UpdateCampaign(input CampaignDetailInput, inputData CampaignCreateInput) (Campaign, error) {
	campaign, err := s.repository.FindByID(input.ID)
	if err != nil {
		return campaign, err
	}

	if int32(inputData.User.ID) != campaign.UserID {
		return campaign, errors.New("Not an owner of the campaign")
	}

	campaign.Name = inputData.Name
	campaign.ShortDescription = inputData.ShortDescription
	campaign.Description = inputData.Description
	campaign.Perks = inputData.Perks
	campaign.GoalAmount = inputData.GoalAmount

	// slugTitle := fmt.Sprintf("%s %d", inputData.Name, inputData.User.ID)
	// campaign.Slug = slug.Make(slugTitle)

	campaignData, err := s.repository.Update(campaign)
	if err != nil {
		return campaign, err
	}
	return campaignData, nil
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

func (s *service) CreateCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error) {
	campaign, err := s.repository.FindByID(input.CampaignID)
	if err != nil {
		return CampaignImage{}, err
	}

	if int32(input.User.ID) != campaign.UserID {
		return CampaignImage{}, errors.New("Not an owner of the campaign")
	}

	isPrimary := 0
	if input.IsPrimary {
		isPrimary = 1
		_, err := s.repository.MarkAllImagesAsNonPrimary(input.CampaignID)
		if err != nil {
			return CampaignImage{}, err
		}
	}

	campaignImage := CampaignImage{}
	campaignImage.CampaignID = int32(input.CampaignID)
	campaignImage.IsPrimary = int8(isPrimary)
	campaignImage.FileName = fileLocation

	newCampaignImage, err := s.repository.CreateImage(campaignImage)
	if err != nil {
		return newCampaignImage, err
	}

	return newCampaignImage, nil
}
