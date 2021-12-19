package processes

import (
	"fmt"
	"time"
)

func Giveaway() {
	var tick = time.Tick(60000 * time.Millisecond)
	for range tick {
		fmt.Println("I check it!")
	}
}