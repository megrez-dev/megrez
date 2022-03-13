package app

import (
	"fmt"
	"log"
	"path"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/model"
	"github.com/megrez/pkg/router"
	dirUtils "github.com/megrez/pkg/utils/dir"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// Megrez application
type Megrez struct {
	config *viper.Viper
	db     *gorm.DB
	server *gin.Engine
	home   string
}

// New create an instance of Megrez
func New() *Megrez {
	return &Megrez{}
}

// Run run blog application
func (m *Megrez) Run() error {
	err := m.server.Run()
	if err != nil {
		return err
	}
	return nil
}

// Init init application
func (m *Megrez) Init() error {
	log.SetPrefix("[MEGREZ-debug] ")
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	for _, f := range []func() error{
		m.initDir,
		m.initConfig,
		m.initDAO,
		m.initRouter,
	} {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}

func (m *Megrez) initDir() error {
	megrezHome, err := dirUtils.GetOrCreateMegrezHome()
	if err != nil {
		return err
	}
	m.home = megrezHome
	return nil
}

func (m *Megrez) initConfig() error {
	v := viper.New()
	v.SetConfigFile(path.Join(m.home, "config.yaml"))

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("Config file not found; ignore error if desired")
			return err
		}
		log.Println("Config file was found but another error was produced")
		return err
	}

	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name, " | ", e.Op.String())
	})
	m.config = v
	return nil
}

func (m *Megrez) initDAO() error {
	dbconfig := m.config.GetStringMapString("mysql")
	host := dbconfig["host"]
	port := dbconfig["port"]
	name := dbconfig["database"]
	user := dbconfig["username"]
	pwd := dbconfig["password"]
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=30s",
		user, pwd, host, port, name)
	db, err := model.New(dsn)
	if err != nil {
		log.Println("connect db failed, ", err)
		return err
	}
	m.db = db
	return nil
}

func (m *Megrez) initRouter() error {
	server, err := router.NewRouter(m.home)
	if err != nil {
		log.Println("init router failed, ", err)
		return err
	}
	m.server = server
	return nil
}
