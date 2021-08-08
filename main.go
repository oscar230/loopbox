// Entrypoint for loopbox

package main

import (
	"fmt"
	"loopbox/lib/rbdj/backupxml"
	"loopbox/utils/env"
	"loopbox/utils/logger"
	"os"
)

func main() {
	env.Load()
	logger.Info("Starting " + fmt.Sprint(os.Getenv("APP_NAME")) + " CLI.")
	b := new(backupxml.Backupxml)
	b.Parse(os.Getenv("RB_BACKUPXML_PATH"))
}
