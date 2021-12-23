package app

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/dao"
	"github.com/megrez/pkg/router"
	"github.com/spf13/viper"
)

// Megrez application
type Megrez struct {
	config *viper.Viper
	dao    *dao.DAO
	server *gin.Engine
}

// NewMegrez create an instance of Megrez
func NewMegrez() *Megrez {
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

func (m *Megrez) initConfig() error {
	v := viper.New()
	// TODO: 通过 NewMegrez 传递 flag 作为配置文件路径
	v.SetConfigFile("./config.yaml")

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
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pwd, host, port, name)
	dao, err := dao.New(dsn)
	if err != nil {
		log.Println("connect db failed, ", err)
		return err
	}
	m.dao = dao
	return nil
}

func (m *Megrez) initRouter() error {
	m.server = router.NewRouter(m.dao)
	return nil
}
