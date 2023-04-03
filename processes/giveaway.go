package processes

import (
	"EdwardBot_LITE/modules"
	"time"
)

func Giveaway() {
	var tick = time.Tick(30000 * time.Millisecond)
	for range tick {
		modules.GiveawayModule()
	}
}
