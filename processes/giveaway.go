package processes

import (
	"time"
	"EdwardBot_LITE/modules"
)

func Giveaway() {
	var tick = time.Tick(30000 * time.Millisecond)
	for range tick {
		modules.GiveawayModule()
	}
}