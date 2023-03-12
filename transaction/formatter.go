package transaction

import "time"

type CampaignTransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int32     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type CampaignTransactionUserFormatter struct {
	ID        int               `json:"id"`
	Amount    int               `json:"amount"`
	Status    string            `json:"status"`
	CreatedAt time.Time         `json:"created_at"`
	Campaign  CampaignFormatter `json:"campaign"`
}

type CampaignFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func FormatCampaignTransactions(transactions []Transaction) []CampaignTransactionFormatter {
	if len(transactions) == 0 {
		return []CampaignTransactionFormatter{}
	}

	var transactionFormatter []CampaignTransactionFormatter
	for _, transaction := range transactions {
		transactionFormatter = append(transactionFormatter, FormatCampaignTransaction(transaction))
	}
	return transactionFormatter
}

func FormatCampaignTransaction(transaction Transaction) CampaignTransactionFormatter {
	formatter := CampaignTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.Name = transaction.User.Name
	formatter.Amount = transaction.Amount
	formatter.CreatedAt = transaction.CreatedAt
	return formatter
}

func FormatCampaignTransactionUsers(transactions []Transaction) []CampaignTransactionUserFormatter {
	if len(transactions) == 0 {
		return []CampaignTransactionUserFormatter{}
	}

	var transactionFormatter []CampaignTransactionUserFormatter
	for _, transaction := range transactions {
		transactionFormatter = append(transactionFormatter, FormatCampaignTransactionUser(transaction))
	}
	return transactionFormatter
}

func FormatCampaignTransactionUser(transaction Transaction) CampaignTransactionUserFormatter {
	formatter := CampaignTransactionUserFormatter{}
	formatter.ID = transaction.ID
	formatter.Amount = int(transaction.Amount)
	formatter.Status = transaction.Status
	formatter.CreatedAt = transaction.CreatedAt

	campaignFormatter := CampaignFormatter{}
	campaignFormatter.Name = transaction.Campaign.Name
	campaignFormatter.ImageURL = ""

	if len(transaction.Campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = transaction.Campaign.CampaignImages[0].FileName
	}

	formatter.Campaign = campaignFormatter
	return formatter
}
