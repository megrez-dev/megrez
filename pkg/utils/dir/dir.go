package dir

import (
	"github.com/mitchellh/go-homedir"
	"os"
	"path"
)

const MegrezDir = ".megrez"

func GetOrCreateMegrezHome() (string, error) {
	homeDir, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	megrezHome := path.Join(homeDir, MegrezDir)
	_, err = os.Stat(megrezHome)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.Mkdir(megrezHome, os.ModePerm); err != nil {
				return "", err
			}
		} else {
			return "", err
		}
	}
	return megrezHome, nil
}
