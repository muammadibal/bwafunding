package campaign

import (
	"strings"
	"time"
)

type CampaignFormatter struct {
	ID               int       `json:"id"`
	UserID           int32     `json:"user_id"`
	Name             string    `json:"name"`
	ShortDescription string    `json:"short_description"`
	Description      string    `json:"description"`
	Perks            string    `json:"perks"`
	BackerAmount     int32     `json:"backer_amount"`
	GoalAmount       int32     `json:"goal_amount"`
	CurrentAmount    int32     `json:"current_amount"`
	ImageURL         string    `json:"image_url"`
	Slug             string    `json:"slug"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type CampaignDetailFormatter struct {
	ID               int                         `json:"id"`
	Name             string                      `json:"name"`
	ShortDescription string                      `json:"short_description"`
	Description      string                      `json:"description"`
	Perks            []string                    `json:"perks"`
	BackerAmount     int32                       `json:"backer_amount"`
	GoalAmount       int32                       `json:"goal_amount"`
	CurrentAmount    int32                       `json:"current_amount"`
	ImageURL         string                      `json:"image_url"`
	Slug             string                      `json:"slug"`
	User             CampaignDetailUserFormatter `json:"user"`
	CampaignImages   []CampaignImageFormatter
}

type CampaignDetailUserFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type CampaignImageFormatter struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	if len(campaigns) == 0 {
		return []CampaignFormatter{}
	}

	var campaignsData []CampaignFormatter

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsData = append(campaignsData, campaignFormatter)
	}

	return campaignsData
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	campaignFormatter := CampaignFormatter{}
	campaignFormatter.ID = campaign.ID
	campaignFormatter.UserID = campaign.UserID
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.ShortDescription
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.CurrentAmount = campaign.CurrentAmount
	campaignFormatter.Slug = campaign.Slug
	campaignFormatter.ImageURL = ""

	if len(campaign.CampaignImages) > 0 {
		var filename string
		for _, image := range campaign.CampaignImages {
			if image.IsPrimary == 1 {
				filename = image.FileName
			}
		}
		campaignFormatter.ImageURL = filename
	}

	return campaignFormatter
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
	campaignFormatter := CampaignDetailFormatter{}
	campaignFormatter.ID = campaign.ID
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.ShortDescription
	campaignFormatter.Description = campaign.Description
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.BackerAmount = campaign.BackerAmount
	campaignFormatter.CurrentAmount = campaign.CurrentAmount

	var perks []string
	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	campaignFormatter.Perks = perks
	campaignFormatter.Slug = campaign.Slug
	campaignFormatter.ImageURL = ""

	if len(campaign.CampaignImages) > 0 {
		var filename string
		for _, image := range campaign.CampaignImages {
			if image.IsPrimary == 1 {
				filename = image.FileName
			}
		}
		campaignFormatter.ImageURL = filename
	}

	user := campaign.User
	campaignUserFormatter := CampaignDetailUserFormatter{}
	campaignUserFormatter.Name = user.Name
	campaignUserFormatter.ImageURL = user.AvatarFileName
	campaignFormatter.User = campaignUserFormatter

	var campaignImagesData []CampaignImageFormatter
	for _, campaignImage := range campaign.CampaignImages {
		isPrimary := false
		if campaignImage.IsPrimary == 1 {
			isPrimary = true
		}

		valueImage := CampaignImageFormatter{
			ImageURL:  campaignImage.FileName,
			IsPrimary: isPrimary,
		}

		campaignImagesData = append(campaignImagesData, valueImage)
	}
	campaignFormatter.CampaignImages = campaignImagesData
	return campaignFormatter
}
