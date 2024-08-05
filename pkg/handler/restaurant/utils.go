package restaurant

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func generateFileName(originalFileName string) string {
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	extension := filepath.Ext(originalFileName)
	nameWithoutExt := strings.TrimSuffix(originalFileName, extension)
	newFileName := fmt.Sprintf("%s_%s%s", timestamp, nameWithoutExt, extension)
	return newFileName
}
