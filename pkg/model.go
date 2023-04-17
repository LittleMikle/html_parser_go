package model

type Money struct {
	Name  string `header:"Валюта"`
	Value string `header:"Курс"`
}
