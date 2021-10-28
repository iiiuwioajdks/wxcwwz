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
			if t == 8 || t == 12 || t == 18 {
				reqHandler.WxReq()
				time.Sleep(time.Hour)
			}
			if t < 8 {
				time.Sleep(time.Duration(8-t) * time.Hour)
			} else if t > 8 && t < 12 {
				time.Sleep(time.Duration(15-t) * time.Hour)
			} else if t > 12 && t < 18 {
				time.Sleep(time.Duration(18-t) * time.Hour)
			} else if t > 18 {
				time.Sleep(time.Duration(24-t) * time.Hour)
			}
		}
		quit <- 1
	}()
	r := router.InitRouter()
	r.Run(":9000")
}
