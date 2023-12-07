package models

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	Nome          string         `json:"nome"`
	Usuario       string         `json:"usuario"`
	Email         string         `json:"email"`
	Senha         string         `json:"senha"`
	Investimentos []Investimento `json:"investimentos"`
}
