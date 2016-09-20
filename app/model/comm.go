package model

import (
	"errors"
	"strings"
	"time"

	"github.com/admpub/confl"
	"github.com/admpub/log"
	"github.com/webx-top/db/lib/factory"
	"github.com/webx-top/db/mongo"
	"github.com/webx-top/db/mysql"
)

type CmdLineConfig struct {
	Port       *int
	Conf       *string
	LogTargets *string
}

type ConfigData struct {
	DB struct {
		Type     string            `json:"type"`
		User     string            `json:"user"`
		Password string            `json:"password"`
		Host     string            `json:"host"`
		Database string            `json:"database"`
		Prefix   string            `json:"prefix"`
		Options  map[string]string `json:"options"`
		Debug    bool              `json:"debug"`
	} `json:"db"`

	Cron struct {
		Interval time.Duration `json:"interval"`
	} `json:"cron"`

	Log struct {
		Debug        bool   `json:"debug"`
		Colorable    bool   `json:"colorable"`    // for console
		SaveFile     string `json:"saveFile"`     // for file
		FileMaxBytes int64  `json:"fileMaxBytes"` // for file
		Targets      string `json:"targets"`
	} `json:"log"`

	Sys struct {
		Accounts map[string]string `json:"accounts"`
	} `json:"sys"`
}

var (
	Config                = &ConfigData{}
	CmdConfig             = &CmdLineConfig{}
	ErrUnknowDatabaseType = errors.New(`unkown database type`)
)

func ParseConfig() error {
	_, err := confl.DecodeFile(*CmdConfig.Conf, Config)
	if err != nil {
		return err
	}
	InitLog()
	return ConnectDB()
}

func ConnectDB() error {
	factory.CloseAll()
	switch Config.DB.Type {
	case `mysql`:
		settings := mysql.ConnectionURL{
			Host:     Config.DB.Host,
			Database: Config.DB.Database,
			User:     Config.DB.User,
			Password: Config.DB.Password,
			Options:  Config.DB.Options,
		}
		database, err := mysql.Open(settings)
		if err != nil {
			return err
		}
		factory.SetDebug(Config.DB.Debug)
		cluster := factory.NewCluster().AddW(database)
		factory.SetCluster(0, cluster).Cluster(0).SetPrefix(Config.DB.Prefix)
	case `mongo`:
		settings := mongo.ConnectionURL{
			Host:     Config.DB.Host,
			Database: Config.DB.Database,
			User:     Config.DB.User,
			Password: Config.DB.Password,
			Options:  Config.DB.Options,
		}
		database, err := mongo.Open(settings)
		if err != nil {
			return err
		}
		factory.SetDebug(Config.DB.Debug)
		cluster := factory.NewCluster().AddW(database)
		factory.SetCluster(0, cluster).Cluster(0).SetPrefix(Config.DB.Prefix)
	default:
		return ErrUnknowDatabaseType
	}
	return nil
}

func InitLog() {

	//======================================================
	// 配置日志
	//======================================================
	if Config.Log.Debug {
		log.DefaultLog.MaxLevel = log.LevelDebug
	} else {
		log.DefaultLog.MaxLevel = log.LevelInfo
	}
	targets := []log.Target{}

	for _, targetName := range strings.Split(Config.Log.Targets, `,`) {
		switch targetName {
		case "console":
			//输出到命令行
			consoleTarget := log.NewConsoleTarget()
			consoleTarget.ColorMode = Config.Log.Colorable
			targets = append(targets, consoleTarget)

		case "file":
			//输出到文件
			fileTarget := log.NewFileTarget()
			fileTarget.FileName = Config.Log.SaveFile
			fileTarget.Filter.MaxLevel = log.LevelInfo
			fileTarget.MaxBytes = Config.Log.FileMaxBytes
			targets = append(targets, fileTarget)
		}
	}

	log.SetTarget(targets...)
	log.SetFatalAction(log.ActionExit)
}

func MustOK(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
