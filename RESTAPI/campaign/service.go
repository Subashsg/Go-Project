package campaign

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
	GetCampaignByID(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	UpdateCampaign(inputID GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error)
	SaveCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(userID int) ([]Campaign, error) {
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

func (s *service) GetCampaignByID(input GetCampaignDetailInput) (Campaign, error) {

	camp, err := s.repository.FindByID(input.ID)

	if err != nil {

		return camp, err
	}

	return camp, nil
}

func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {

	camp := Campaign{}

	camp.Name = input.Name
	camp.ShortDescription = input.Description
	camp.Description = input.Description
	camp.GoalAmount = input.GoalAmount
	camp.Perks = input.Perks
	camp.User.ID = input.User.ID

	slugCandidate := fmt.Sprintf("%s %d", input.Name, input.User.ID)

	camp.Slug = slug.Make(slugCandidate)

	campaigns, err := s.repository.Save(camp)

	if err != nil {

		return campaigns, err
	}

	return campaigns, nil
}

func (s *service) UpdateCampaign(inputID GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error) {

	camp, err := s.repository.FindByID(inputID.ID)

	if err != nil {

		return camp, err
	}

	if camp.UserID != inputData.User.ID {

		return camp, errors.New("not an owner of the campaign")
	}

	camp.Name = inputData.Name
	camp.ShortDescription = inputData.ShortDescription
	camp.GoalAmount = inputData.GoalAmount
	camp.Description = inputData.Description
	camp.Perks = inputData.Perks

	newcamp, err := s.repository.Update(camp)

	if err != nil {

		return newcamp, err
	}
	return newcamp, nil
}

func (s *service) SaveCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error) {

	campaign, err := s.repository.FindByID(input.User.ID)

	if err != nil {

		return CampaignImage{}, err
	}

	if campaign.UserID != input.User.ID {
		return CampaignImage{}, errors.New("not an owner of the campaign")
	}

	isPrimary := 0

	if input.IsPrimary { // this will work when  isPrimary = 0 (isPrimary false )

		isPrimary = 1

		_, err := s.repository.MarkAllImagesAsNonPrimary(input.CampaignID)
		if err != nil {
			return CampaignImage{}, err
		}
	}

	campaignImage := CampaignImage{}

	campaignImage.CampaignID = input.CampaignID
	campaignImage.IsPrimary = isPrimary
	campaignImage.FileName = fileLocation

	newCampaignImage, err := s.repository.CreateImage(campaignImage)

	if err != nil {

		return newCampaignImage, err
	}

	return newCampaignImage, nil
}
