package main

import (
	"DealSystem/models"
	"DealSystem/services"
	"fmt"
	"time"
)

//ou happen to be a budding entrepreneur and you have come up with an idea to build an e-commerce giant like Amazon, Flipkart, Walmart, etc. As part of this ambition, you want to build a platform to duplicate the concept of Limited Time Deals.
//
//
//
//Limited Time Deals
//
//A limited-time deal implies that a seller will put up an item on sale for a limited time period, say, 2 hours, and will keep a maximum limit on the number of items that would be sold as part of that deal.
//than one item as part of the deal.
//
//
//
//The task is to create APIs to enable the following operations
//
//Create a deal with the price and number of items to be sold as part of the deal
//
//End a deal
//
//Update a deal to increase the number of items or end time
//
//Claim a deal
//
//Guidelines
//
//Document and communicate your assumptions in README.
//
//Create a working solution with production-quality code.
//
//Define and
//
//Create APIs to support the operations mentioned above
//
//Write a few unit tests for the most important code
//Users cannot buy the deal if the deal time is over
//
//Users cannot buy if the maximum allowed deal has already been bought by other users.

// One user cannot buy more
func main() {
	ds := services.DealService{}

	deal1 := ds.CreateDeal("fan", 5, time.Now(), time.Now().Add(10*time.Hour))

	err := ds.Create(deal1)
	if err != nil {
		println(err.Error())
	}
	user1 := &models.User{Name: "User1"}

	err = ds.Claim(user1, "fan", time.Now().Add(50*time.Hour))
	if err != nil {
		println(err.Error())
	}

	err = ds.Claim(user1, "fan", time.Now().Add(5*time.Hour))
	if err != nil {
		println(err.Error())
	}

	user2 := &models.User{Name: "User2"}
	err = ds.Claim(user2, "fan", time.Now().Add(5*time.Hour))
	if err != nil {
		println(err.Error())
	}
	user3 := &models.User{Name: "User3"}
	err = ds.Claim(user3, "fan", time.Now().Add(5*time.Hour))
	if err != nil {
		println(err.Error())
	}
	// fails for count = 3
	user4 := &models.User{Name: "User4"}
	err = ds.Claim(user4, "fan", time.Now().Add(5*time.Hour))
	if err != nil {
		println(err.Error())
	}
	// fails for time
	user5 := &models.User{Name: "User5"}
	err = ds.Claim(user5, "fan", time.Now().Add(15*time.Hour))
	if err != nil {
		println("time over shoot : ", err.Error())
	}

	err = ds.Update(deal1, time.Now().Add(16*time.Hour), 1)
	if err != nil {
		println("update time error : ", err.Error())
	}
	fmt.Println(deal1.Start, deal1.End)

	err = ds.Claim(user5, "fan", time.Now().Add(15*time.Hour))
	if err != nil {
		println(err.Error())
	}

}
