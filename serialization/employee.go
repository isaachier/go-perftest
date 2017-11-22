package serialization

import (
	"math/rand"
	"time"

	"github.com/Pallinder/go-randomdata"
)

type employee struct {
	firstName  string
	lastName   string
	position   string
	salary     float64
	id         int64
	managerID  *int64
	updateTime int64
}

var dataset []employee
var positions = [...]string{"engineer", "janitor", "manager"}

func init() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100000; i++ {
		dataset = append(dataset, employee{
			firstName:  randomdata.FirstName(randomdata.RandomGender),
			lastName:   randomdata.LastName(),
			position:   positions[rand.Intn(len(positions))],
			salary:     float64(rand.Intn(1000000)),
			id:         rand.Int63(),
			updateTime: time.Now().Unix(),
		})
		if randomdata.Boolean() {
			managerID := rand.Int63()
			dataset[len(dataset)-1].managerID = &managerID
		}
	}
}
