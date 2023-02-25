package campaign

import "time"

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
	// CampaignImages   []CampaignImageFormatter
}

type CampaignImageFormatter struct {
	ID         int       `json:"id"`
	CampaignId int32     `json:"campaign_id"`
	FileName   string    `json:"filename"`
	IsPrimary  int8      `json:"is_primary"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
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
