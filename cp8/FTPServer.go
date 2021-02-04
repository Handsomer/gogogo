package cp8

import (
	"fmt"
	"sync"
)

func SyncMapUse() {
	var syncMap sync.Map
	// Load use
	value, ok := syncMap.Load("user_one")
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("no this user")
	}
	// Store use
	syncMap.Store("user_007", "/home/user_007")
	value, ok = syncMap.Load("user_007")
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("no this user")
	}
	// Delete use
	if ok {
		syncMap.Delete("user_007")
	}
	value, ok = syncMap.Load("user_007")
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("no this user")
	}
	// range Use
	syncMap.Store("user_007", "/home/user_007")
	syncMap.Store("user_008", "/home/user_008")
	syncMap.Store("user_009", "/home/user_009")
	syncMap.Store("user_010", "/home/user_010")
	f := func(k, v interface{}) bool {
		fmt.Println(k, "   :   ", v)
		return true
	}
	syncMap.Range(f)
}
