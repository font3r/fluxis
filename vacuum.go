package main

import (
	"fmt"
	"time"
)

func StartVacuum(st *Storage, now func() time.Time) {
	fmt.Println("INFO: starting vacuum")

	timer := time.NewTicker(5 * time.Second)

	for {
		vacuumTime := <-timer.C

		fmt.Println("INFO: vacuum cleanup start at", vacuumTime.Unix())
		for key, entry := range st.entries {
			if vacuumTime.After(time.Unix(entry.TTL, 0)) {
				delete(st.entries, key)
				fmt.Printf("INFO: vacuum cleaned key: %s\n", key)
			}
		}

		vacuumEnd := now()
		fmt.Printf("INFO: vacuum cleanup end at %d, took %d\n",
			vacuumEnd.Unix(), int(vacuumEnd.Sub(vacuumTime).Seconds()))
	}
}
