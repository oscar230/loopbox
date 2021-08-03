// Entrypoint for loopbox

package main

import (
	"fmt"
	"loopbox/utils/env"
	"loopbox/utils/logger"
	"os"
)

func main() {
	env.Load()
	logger.Info("Starting " + fmt.Sprint(os.Getenv("APP_NAME")) + " CLI.")
}
