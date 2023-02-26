package campaign

import "bwafunding/user"

type CampaignDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CampaignCreateInput struct {
	Name             string `json:"name" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description      string `json:"description" binding:"required"`
	GoalAmount       int32  `json:"goal_amount" binding:"required"`
	Perks            string `json:"perks" binding:"required"`
	User             user.User
}
