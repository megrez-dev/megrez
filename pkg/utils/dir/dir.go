package dir

import (
	"bytes"
	"github.com/megrez/pkg/log"
	"github.com/mitchellh/go-homedir"
	"io"
	"io/fs"
	"os"
	"path"
)

const MegrezDir = ".megrez"

func GetOrCreateMegrezHome() (string, error) {
	homeDir, err := homedir.Dir()
	if err != nil {
		log.Error(err)
		return "", err
	}
	megrezHome := path.Join(homeDir, MegrezDir)
	_, err = os.Stat(megrezHome)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.Mkdir(megrezHome, os.ModePerm); err != nil {
				log.Error(err)
				return "", err
			}
		} else {
			log.Error(err)
			return "", err
		}
	}
	return megrezHome, nil
}

// CopyDirFromFS
// @param static: static/default/css/xxx.css
// @src: default
// @dst: ./megrez/themes
// @return: relative path
func CopyDirFromFS(static fs.FS, src string, dst string) error {
	// dirs: default/...
	err := CreateDir(path.Join(dst, src))
	if err != nil {
		log.Error(err)
		return err
	}
	dirs, err := fs.ReadDir(static, src)
	if err != nil {
		log.Error(err)
		return err
	}
	defaultDir, err := fs.Sub(static, src)
	if err != nil {
		log.Error(err)
		return err
	}
	for _, dir := range dirs {
		if dir.IsDir() {
			err := CopyDirFromFS(defaultDir, dir.Name(), path.Join(dst, src))
			if err != nil {
				log.Error(err)
				return err
			}
		} else {
			file, err := fs.ReadFile(defaultDir, dir.Name())
			if err != nil {
				log.Error(err)
				return err
			}
			created, err := CreateFile(path.Join(dst, src, dir.Name()))
			if err != nil {
				log.Error(err)
				return err
			}
			_, err = io.Copy(created, bytes.NewReader(file))
			if err != nil {
				log.Error(err)
				return err
			}
		}
	}
	return nil
}

func CreateDir(dir string) error {
	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return os.Mkdir(dir, os.ModePerm)
		} else {
			log.Error(err)
			return err
		}
	}
	return nil
}

func CreateFile(path string) (*os.File, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return os.Create(path)
		} else {
			log.Error(err)
			return nil, err
		}
	}
	return os.Open(path)
}
