package test

import (
	"DealSystem/models"
	"DealSystem/services"
	"testing"
	"time"
)

func TestClaim(t *testing.T) {
	ds := services.DealService{}
	deal1 := ds.CreateDeal("fan", 5, time.Now(), time.Now().Add(10*time.Hour))
	err := ds.Create(deal1)
	if err != nil {
		println(err.Error())
	}
	user1 := &models.User{Name: "User1"}
	err = ds.Claim(user1, "fan", time.Now().Add(5*time.Hour))
	if err != nil {
		t.Errorf("error in Claim ::" + err.Error())
	}
}
