// Name, phone number, address

package models

type Contact struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
}
