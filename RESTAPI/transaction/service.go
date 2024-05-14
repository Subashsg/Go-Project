package transaction

import (
	"User/campaign"
	"User/payment"
	"errors"
	"strconv"
)

type service struct {
	repository         Repository          // interface
	campaignRepository campaign.Repository // interface
	paymentservice     payment.Service
}

type Service interface {
	GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error)
	GetTransactionsByUserID(userID int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
	ProcessPayment(input TransactionNotificationInput) error
	GetAllTransactions() ([]Transaction, error)
}

func NewService(repository Repository, campaignRepository campaign.Repository, paymentservice payment.Service) *service {

	return &service{repository, campaignRepository, paymentservice}
}

func (s *service) GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error) {

	campaign, err := s.campaignRepository.FindByID(input.ID)

	if err != nil {
		return []Transaction{}, err
	}

	if campaign.UserID != input.User.ID {
		return []Transaction{}, errors.New("Not an owner of the campaign")
	}

	transactions, err := s.repository.GetByCampaignID(input.ID)

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) GetTransactionsByUserID(userID int) ([]Transaction, error) {

	transactions, err := s.repository.GetByUserID(userID)

	if err != nil {

		return transactions, err
	}
	return transactions, nil
}

func (s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error) {

	Transactions := Transaction{}

	Transactions.Amount = input.Amount
	Transactions.CampaignID = input.CampaignID
	Transactions.UserID = input.User.ID
	Transactions.Status = "pending"

	newTransaction, err := s.repository.Save(Transactions)

	if err != nil {

		return newTransaction, err
	}

	paymentTransacation := payment.Transaction{

		ID:     newTransaction.ID,
		Amount: newTransaction.Amount,
	}

	paymentURL, err := s.paymentservice.GetPaymentURL(paymentTransacation, input.User)

	if err != nil {
		return newTransaction, err
	}

	newTransaction.PaymentURL = paymentURL

	newtrans, err := s.repository.Update(newTransaction)

	if err != nil {
		return newtrans, err
	}

	return newtrans, nil
}

func (s *service) ProcessPayment(input TransactionNotificationInput) error {

	transaction_id, _ := strconv.Atoi(input.OrderID) // converting string to int

	transaction, err := s.repository.GetByID(transaction_id)

	if err != nil {
		return err
	}

	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {

		transaction.Status = "paid"
	} else if input.TransactionStatus == "settlement" {

		transaction.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {

		transaction.Status = "cancelled"
	}

	updatedTransaction, err := s.repository.Update(transaction)

	if err != nil {

		return err
	}

	campaign, err := s.campaignRepository.FindByID(updatedTransaction.CampaignID)

	if err != nil {

		return err
	}

	if updatedTransaction.Status == "paid" {

		campaign.BackerCount = campaign.BackerCount + 1
		campaign.CurrentAmount = campaign.CurrentAmount + updatedTransaction.Amount

		_, err := s.campaignRepository.Update(campaign)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *service) GetAllTransactions() ([]Transaction, error) {

	transactions, err := s.repository.FindAll()

	if err != nil {

		return transactions, err
	}
	return transactions, nil
}
