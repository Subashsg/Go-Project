package models

import 	"time"


type Employee struct {
	ID          int        `json:"id" gorm:"primary_key"`
	Name        string     `json:"name"`
	CheckIn     time.Time  `json:"check_in"`
	CheckOut    *time.Time `json:"check_out"`
	HoursWorked int        `json:"hours_worked"`
}
