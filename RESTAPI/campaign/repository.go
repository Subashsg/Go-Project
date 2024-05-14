package campaign

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Campaign, error)
	FindByUserID(userID int) ([]Campaign, error)
	FindByID(ID int) (Campaign, error)
	Save(campaign Campaign) (Campaign, error)
	Update(campaign Campaign) (Campaign, error)
	CreateImage(campaignImage CampaignImage) (CampaignImage, error)
	MarkAllImagesAsNonPrimary(campaignID int) (bool, error)
}

type Reposit struct {
	db *gorm.DB
}

func NewRepository(DB *gorm.DB) *Reposit {

	return &Reposit{DB}
}

func (r *Reposit) FindAll() ([]Campaign, error) {

	var camp []Campaign

	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&camp).Error

	if err != nil {

		return camp, err
	}
	return camp, nil
}

func (r *Reposit) FindByUserID(userID int) ([]Campaign, error) {

	var camp []Campaign

	err := r.db.Where("id =?", userID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&camp).Error

	if err != nil {

		return camp, err
	}
	return camp, nil
}

func (r *Reposit) FindByID(ID int) (Campaign, error) {

	var camp Campaign

	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&camp).Error

	if err != nil {

		return camp, err
	}

	return camp, nil
}

func (r *Reposit) Save(campaign Campaign) (Campaign, error) {

	err := r.db.Create(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *Reposit) Update(campaign Campaign) (Campaign, error) {

	err := r.db.Save(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil

}

func (r *Reposit) CreateImage(campaignImage CampaignImage) (CampaignImage, error) {

	err := r.db.Create(&campaignImage).Error

	if err != nil {

		return campaignImage, err
	}

	return campaignImage, nil
}

func (r *Reposit) MarkAllImagesAsNonPrimary(campaignID int) (bool, error) {

	// campaignID is matching with return true

	err := r.db.Model(&CampaignImage{}).Where("campaign_id = ?", campaignID).Update("is_primary", false).Error

	if err != nil {

		return false, err
	}
	return true, nil
}
