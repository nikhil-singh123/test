package models

type Users struct {
	ID            int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name          string `json:"name"`
	Email         string `json:"email" gorm:"unique;not null;"`
	ContactNumber int    `json:"contactnumber"`
	Role          string `json:"role"`
	LibID         int    `json:"libid"`
}

func (u *Users) SaveUser() (*Users, error) {

	if err := DB.Create(&u).Error; err != nil {
		return &Users{}, err
	}
	return u, nil
}
