package file

import (
	"fmt"
)

func FormatSize(fileSize int64) string {
	if fileSize < 1024 {
		return fmt.Sprintf("%d B", fileSize)
	} else if fileSize < 1024*1024 {
		return fmt.Sprintf("%.2f KB", float64(fileSize)/1024)
	} else if fileSize < 1024*1024*1024 {
		return fmt.Sprintf("%.2f MB", float64(fileSize)/1024/1024)
	} else {
		return fmt.Sprintf("%.2f GB", float64(fileSize)/1024/1024/1024)
	}
}
