package services

import (
	"DealSystem/models"
	"errors"
	"fmt"
	"sync"
	"time"
)

type DealService struct {
	mx    sync.Mutex
	deals []*models.Deal
}

func (ds *DealService) CreateDeal(name string, count int, start time.Time, end time.Time) *models.Deal {
	return &models.Deal{
		ItemName: name,
		Count:    count,
		Start:    start,
		End:      end,
	}
}

func isOverLap(a *models.Deal, b *models.Deal) bool {
	if a.Start.After(b.Start) {
		return isOverLap(b, a)
	}
	if a.End.After(b.Start) {
		return true
	}
	return false
}

func (ds *DealService) Create(cd *models.Deal) error {
	ds.mx.Lock()
	defer ds.mx.Unlock()

	for _, d := range ds.deals {
		if d.ItemName == cd.ItemName {
			if isOverLap(d, cd) {
				fmt.Println("error : duplicate deals at the overlaping time")
				return errors.New("duplicate deals at the overlaping time")
			}
		}
	}
	ds.deals = append(ds.deals, cd)
	fmt.Println("Successfully Created deal")
	return nil
}

func (ds *DealService) End(cd *models.Deal) error {
	ds.mx.Lock()
	defer ds.mx.Unlock()
	index := -1
	for i, d := range ds.deals {
		if cd == d {
			if cd.End.Before(time.Now()) {
				index = i
			}
		}
	}
	if index != -1 {
		fmt.Println("Successfully ended deal")
		ds.deals = append(ds.deals[:index], ds.deals[index+1:]...)
		return nil
	} else {
		fmt.Println("error : no valid deals found")
		return errors.New("no valid deals found")
	}
}

func (ds *DealService) Update(cd *models.Deal, endTime time.Time, updatedCount int) error {
	cd.Mx.Lock()
	defer cd.Mx.Unlock()
	if endTime.Before(time.Now()) {
		return errors.New("time has passed")
	}
	if endTime.Before(cd.End) {
		return errors.New("end time will decrease")
	}
	if updatedCount < cd.Count {
		return errors.New("count cant decrease")
	}
	cd.End = endTime
	cd.Count = updatedCount
	fmt.Println("Successfully updated deal")
	return nil
}

func (ds *DealService) Claim(currUser *models.User, name string, currentTime time.Time) error {
	var validDeal *models.Deal
	for _, d := range ds.deals {
		if d.ItemName == name && d.Start.Before(currentTime) && d.End.After(currentTime) {
			validDeal = d
		}
	}
	if validDeal == nil {
		return errors.New("no valid deal found")
	}
	validDeal.Mx.Lock()
	defer validDeal.Mx.Unlock()
	if validDeal.Count == 0 {
		return errors.New("all deals claimed")
	}
	for _, u := range validDeal.UsersBought {
		if u == currUser {
			return errors.New("user can not claim a deal twice")
		}
	}
	validDeal.Count--
	validDeal.UsersBought = append(validDeal.UsersBought, currUser)
	fmt.Println("Successfully Claimed deal")
	return nil
}
