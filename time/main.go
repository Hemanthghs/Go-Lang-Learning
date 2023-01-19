package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Format("01-02-2006 15:04:05 Monday"))

	date := time.Date(2020, time.December, 13, 10, 0, 0, 0, time.Local)

	fmt.Println(date)
	fmt.Println(date.Format("01-02-2006 Monday 15:04:05"))
}

func UserActivityAdd(c activity_pb.UserServiceClient) {
	t := time.Now()
	ts := t.Format("01-02-2006 15:04:05 Monday")
	userActivityAddRequest := activity_pb.ActivityRequest{
		Activity: &activity_pb.Activity{
			ActivityType: 0,
			Timestamp:    ts,
			Duration:     10,
			Label:        "Working",
			Email:        "hemanth@gmail.com",
		},
	}
	res, err := c.UserActivityAdd(context.Background(), &userActivityAddRequest)
	handleError(err)
	fmt.Println(res)
}
