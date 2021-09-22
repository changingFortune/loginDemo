package main

import (
	"github.com/sony/sonyflake"
	"time"
)

func main() {
	sonyflake.NewSonyflake(sonyflake.Settings{})
	var st sonyflake.Settings
	st.StartTime = time.Now()
	sf := sonyflake.NewSonyflake(st)
	f := func() {
		for i:=0;i<300;i++ {
			print(sf.NextID())
		}
	}
	go f()
	go f()
	go f()
	go f()
	


}
