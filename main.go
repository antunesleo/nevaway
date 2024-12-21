package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)


func activateWindow(name string) error {
	pids, err := robotgo.FindIds(name)
	if err == nil && len(pids) > 0 {
		succeeded := false
		for pid := range pids {
			err = robotgo.ActivePid(pid)
			if err == nil {
				succeeded = true
				break
			}
		}
		if !succeeded {
			return fmt.Errorf("we find pids but failed to activate the window")
		}
		return nil
	}

	return err
}




func main() {
	err := activateWindow("slack")
	if err != nil {
		fmt.Println(err)
	}

	for {
		time.Sleep(5 * time.Second)
		x, y := robotgo.Location()
		robotgo.MoveSmooth(x+1, y+1)
	}
}