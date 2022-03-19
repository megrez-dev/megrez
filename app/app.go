package app

import (
	"fmt"
	"go.uber.org/zap"
	"path"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/log"
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
	Home   string
	Logger *zap.Logger
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
	for _, f := range []func() error{
		m.initDir,
		m.initLogger,
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
	m.Home = megrezHome
	return nil
}

func (m *Megrez) initLogger() error {
	logger, err := log.New()
	if err != nil {
		return err
	}
	m.Logger = logger
	return nil
}

func (m *Megrez) initConfig() error {
	v := viper.New()
	v.SetConfigFile(path.Join(m.Home, "config.yaml"))

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Error("Config file not found; ignore error if desired")
			return err
		}
		log.Error("Config file was found but another error was produced")
		return err
	}

	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		log.Info("Config file changed:", e.Name, " | ", e.Op.String())
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
		log.Error("connect db failed, ", err)
		return err
	}
	m.db = db
	return nil
}

func (m *Megrez) initRouter() error {
	server, err := router.NewRouter(m.Logger)
	if err != nil {
		log.Error("init router failed, ", err)
		return err
	}
	m.server = server
	return nil
}
