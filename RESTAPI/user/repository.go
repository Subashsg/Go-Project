package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error)
	FindByemail(email string) (User, error)
	FindById(ID int) (User, error)
	Update(user User) (User, error)
	Findall() ([]User, error)
}

type Reposit struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Reposit {

	return &Reposit{db}

}

func (r *Reposit) Save(user User) (User, error) {

	err := r.db.Create(&user).Error // is used to insert a new record into the database.

	if err != nil {

		return user, err
	}
	return user, nil
}

func (r *Reposit) FindByemail(email string) (User, error) {

	var user User
	err := r.db.Where("email = ?", email).Find(&user).Error

	if err != nil {

		return user, err
	}
	return user, nil
}

func (r *Reposit) FindById(ID int) (User, error) {

	var user User

	err := r.db.Where("id = ?", ID).Find(&user).Error

	if err != nil {

		return user, err
	}

	return user, nil
}

func (r *Reposit) Update(user User) (User, error) {

	err := r.db.Save(&user).Error

	if err != nil {

		return user, err
	}
	return user, nil
}

func (r *Reposit) Findall() ([]User, error) {

	var users []User

	err := r.db.Find(&users).Error

	if err != nil {

		return users, err
	}
	return users, nil
}
