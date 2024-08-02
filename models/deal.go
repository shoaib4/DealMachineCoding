package models

import (
	"sync"
	"time"
)

type Deal struct {
	Mx          sync.Mutex
	UsersBought []*User
	ItemName    string
	Count       int
	Start       time.Time
	End         time.Time
}
