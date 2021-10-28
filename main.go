package main

import (
	"time"
	"wvCheck/dbHandler"
	"wvCheck/reqHandler"
	"wvCheck/router"
)

func main() {
	err := dbHandler.InitConfig()
	if err != nil {
		panic(err)
	}
	dbHandler.InitDB()
	quit := make(chan int)
	go func() {
		for {
			t := time.Now().Hour()
			if t == 9 || t == 13 || t == 19 {
				reqHandler.WxReq()
				minute := time.Now().Minute()
				m := time.Duration(minute) * time.Minute
				time.Sleep(time.Hour - m)
			}
			if t < 9 {
				time.Sleep(time.Duration(8-t) * time.Hour)
			} else if t > 9 && t < 13 {
				time.Sleep(time.Duration(15-t) * time.Hour)
			} else if t > 13 && t < 19 {
				time.Sleep(time.Duration(18-t) * time.Hour)
			} else if t > 19 {
				time.Sleep(time.Duration(24-t) * time.Hour)
			}
		}
		quit <- 1
	}()
	r := router.InitRouter()
	r.Run(":9000")
}
