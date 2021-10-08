// Name, phone number, address

package models

type Contact struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	PhoneNumber string `json:"PhoneNumber"`
	Address     string `json:"Address"`
}
