package zip

import (
	"archive/zip"
	"io"
	"os"
	"path"
)

func UnZip(reader *zip.Reader, dstPath string) error {
	// create dest dir
	err := os.MkdirAll(dstPath, 0755)
	if err != nil {
		return err
	}
	// 遍历压缩包的内容
	// 注意：reader.File获取到的是压缩包内的所有文件，包括子文件夹下的文件
	for _, file := range reader.File {
		// 文件夹就不解压出来了
		if file.FileInfo().IsDir() { // 文件夹
			// 不管三七二十一，先创建目标文件夹
			err := os.MkdirAll(path.Join(dstPath, file.Name), 0755)
			if err != nil {
				return err
			}
		} else { // 文件
			// 打开压缩包内的文件
			srcFile, err := file.Open()
			if err != nil {
				return err
			}
			defer srcFile.Close()
			// 在文件夹内创建这个文件
			destFile, err := os.Create(path.Join(dstPath, file.Name))
			if err != nil {
				return err
			}
			defer destFile.Close()
			// 执行文件拷贝
			_, err = io.Copy(destFile, srcFile)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
