package djusb

import (
	"loopbox/utils/logger"
	"os"
	"path/filepath"
)

func Scan(root string) {
	if fErr := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		logger.Info("[djusb] opened " + root)
		return nil
	}); fErr != nil {
		logger.Error("[djusb] Error opening usb: " + fErr.Error())
	}
}
