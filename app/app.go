package app

import (
	"fmt"
	"github.com/megrez/pkg/config"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/log"
	"github.com/megrez/pkg/model"
	"github.com/megrez/pkg/router"
	dirUtils "github.com/megrez/pkg/utils/dir"
	"gorm.io/gorm"
)

// Megrez application
type Megrez struct {
	config *config.Config
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
	logger, err := log.New(true)
	if err != nil {
		return err
	}
	m.Logger = logger
	return nil
}

func (m *Megrez) initConfig() error {
	var cfg = &config.Config{}
	c, err := ioutil.ReadFile(path.Join(m.Home, "config.yaml"))
	if err != nil {
		log.Error("read config file failed, ", err.Error())
		return m.initDefaultConfig()
	}
	err = yaml.Unmarshal(c, cfg)
	if err != nil {
		log.Error("unmarshal config failed, ", err.Error())
		return m.initDefaultConfig()
	}
	m.config = cfg
	return nil
}

func (m *Megrez) initDefaultConfig() error {
	var cfg = &config.Config{}
	cfg.Database.SQLite.Path = path.Join(m.Home, "megrez.db")
	cfg.Debug = false
	m.config = cfg
	return nil
}

func (m *Megrez) initDAO() error {
	var db *gorm.DB
	var err error
	if m.config.Database.SQLite.Path != "" {
		log.Info("Use SQLite as database")
		db, err = model.NewSQLite(m.config.Database.SQLite.Path)
		if err != nil {
			return err
		}
	} else if m.config.Database.MySQL.Host != "" {
		log.Info("Use MySQL as database")
		host := m.config.Database.MySQL.Host
		port := m.config.Database.MySQL.Port
		name := m.config.Database.MySQL.Name
		user := m.config.Database.MySQL.User
		pwd := m.config.Database.MySQL.Password
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=30s",
			user, pwd, host, port, name)
		db, err = model.NewMySQL(dsn)
		if err != nil {
			log.Error("connect db failed, ", err)
			return err
		}
	} else {
		log.Error("Please check database config correctly")
		return fmt.Errorf("database config error")
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
