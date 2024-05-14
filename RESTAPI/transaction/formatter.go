package transaction

import "time"

type CampaignTransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatCampaignTransaction(t Transaction) CampaignTransactionFormatter {

	formatter := CampaignTransactionFormatter{

		ID:        t.ID,
		Name:      t.User.Name,
		Amount:    t.Amount,
		CreatedAt: t.CreatedAt,
	}
	return formatter
}

func FormatCampaignTransactions(t []Transaction) []CampaignTransactionFormatter {

	if len(t) == 0 {

		return []CampaignTransactionFormatter{}
	}

	var transactionsFormatter []CampaignTransactionFormatter

	for _, value := range t {

		formater := FormatCampaignTransaction(value)

		transactionsFormatter = append(transactionsFormatter, formater)
	}
	return transactionsFormatter
}

type UserTransactionFormatter struct {
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

func FormatUserTransaction(t Transaction) UserTransactionFormatter {

	formatter := UserTransactionFormatter{}

	formatter.ID = t.ID
	formatter.Amount = t.Amount
	formatter.Status = t.Status
	formatter.CreatedAt = t.CreatedAt

	campformatter := CampaignFormatter{}

	campformatter.Name = t.Campaign.Name
	campformatter.ImageURL = ""

	if len(t.Campaign.CampaignImages) > 0 {

		campformatter.ImageURL = t.Campaign.CampaignImages[0].FileName
	}

	formatter.Campaign = campformatter //  formatter is object of UserTransactionFormatter{} struct
	// campformatter is object of  CampaignFormatter{} struct
	return formatter // formatter.Campaign is filed = value is campformatter
}

func FormatUserTransactions(transactions []Transaction) []UserTransactionFormatter {
	if len(transactions) == 0 {
		return []UserTransactionFormatter{}
	}

	var transactionsFormatter []UserTransactionFormatter

	for _, transaction := range transactions {
		formatter := FormatUserTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, formatter)
	}

	return transactionsFormatter
}

type TransactionFormatter struct {
	ID         int    `json:"id"`
	CampaignID int    `json:"campaign_id"`
	UserID     int    `json:"user_id"`
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
	Code       string `json:"code"`
	PaymentURL string `json:"payment_url"`
}

func FormatTransaction(t Transaction) TransactionFormatter {

	formater := TransactionFormatter{}

	formater.ID = t.ID
	formater.CampaignID = t.CampaignID
	formater.UserID = t.UserID
	formater.Amount = t.Amount
	formater.Status = t.Status
	formater.Code = t.Code
	formater.PaymentURL = t.PaymentURL
	return formater
}
