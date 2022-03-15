package common
import (
        "gopkg.in/ini.v1"
        "os"
        "wxcrm/pkg/common/log"
        "github.com/jpillora/opts"
)

var (
  version = "0.5.2"
)

type Opts struct{
        Port               int      `opts:"help=listening port, default=8089, env=PORT"`
        ServerName         string   `opts:"help=Access server dns name, default=configuration file set"`
        RedisHost          string   `opts:"help=redis server addresss, default=configuration file set"`
        MysqlHost          string   `opts:"help=mysql server addresss, default=configuration file set"`
        MysqlUser          string   `opts:"help=mysql server username, default=configuration file set"`
        MysqlPassword      string   `opts:"help=mysql server user password, default=configuration file set"`
        DatabaseName       string   `opts:"help=mysql server database name,default=configuration file set"`
        WXCorpId           string   `opts:"help=wx corp id, default=configuration file set"`
        WXAppSecret        string   `opts:"help=wx app secret, default=configuration file set"`
        WXAgentId          string   `opts:"help=wx app agent id, default=configuration file set"`
        WXContactSecret    string   `opts:"help=wx contact secret, default=configuration file set"`
        WXRecycleEventsToken         string   `opts:"help=wx customer change recycle event token, default=configuration file set"`
        WXRecycleEventsAesKey        string   `opts:"help=wx customer change recycle event aes key, default=configuration file set"`
        YSAppKey           string   `opts:"help=ys app key, default=configuration file set"`
        YSAppSecret        string   `opts:"help=ys app secret, default=configuration file set"`
        QCCAppKey          string   `opts:"help=qcc app key, default=configuration file set"`
        QCCAppSecret       string   `opts:"help=qcc app secret, default=configuration file set"`
        AdminUser          string   `opts:"help=admin user name, default=configuration file set"`
        LogFile            string   `opts:"help=application log file, default=configuration file set"`
        LogLevel           string   `opts:"help=application log level, default=configuration file set"`
        ConfigFile         string   `opts:"help=application configuration file, default=wxcrm.conf"`
        SrcDir             string   `opts:"help=application source directory, default=src"`
        NotificationUsers  string   `opts:"help=notification users, default=configuration file set"`
}

//Configuration fields
type Obj struct {
  Mysql  Mysql
  Redis  Redisc
  WX     WX
  YS     YS
  QCC    QCC
  Admin  string
  NotificationUsers string
  Port   int
  ServerName    string
  LogFile string
  LogLevel string
  SrcDir   string 
}

type Mysql struct{
  User   string
  Password string
  Host    string
  DBname  string
}

type Redisc struct{
  Host   string
}

type WX struct{
  CorpId  string
  AppSecret string
  ContactSecret  string
  AgentId      string 
  RecycleEventsToken    string  
  RecycleEventsAesKey  string
}

type YS struct{
  AppKey  string
  AppSecret string
}

type QCC struct{
  AppKey  string
  AppSecret  string
}

func NewLogger(logfile,LogLevel string)*log.Logger{
    var Logger *log.Logger
    file, err := os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
            panic("Failed to open configuration file.")
    }
    switch LogLevel{
    case "debug":
            Logger = log.New(file,"",log.Ldate|log.Ltime|log.Lshortfile,0)
    case "info":
            Logger = log.New(file,"",log.Ldate|log.Ltime|log.Lshortfile,1)
    case "warning":
            Logger = log.New(file,"",log.Ldate|log.Ltime|log.Lshortfile,2)
    case "error":
            Logger = log.New(file,"",log.Ldate|log.Ltime|log.Lshortfile,3)
    case "fatal":
            Logger = log.New(file,"",log.Ldate|log.Ltime|log.Lshortfile,4)
    default:
            Logger = log.New(file,"",log.Ldate|log.Ltime|log.Lshortfile,1)
    }
    return Logger
}

func ConfigParse(ConfigFile string) *Obj {
        configuration := &Obj{}
        cfg, err := ini.Load(ConfigFile)
        if err != nil {
                panic("Fail to read configuration file.")
        }
        configuration.Port,_ = cfg.Section("").Key("port").Int()
        configuration.ServerName = cfg.Section("").Key("name").String()
        configuration.SrcDir = cfg.Section("").Key("srcdir").String()
        configuration.NotificationUsers = cfg.Section("notification").Key("users").String()
        configuration.LogFile = cfg.Section("log").Key("logfile").String()
        configuration.LogLevel = cfg.Section("log").Key("level").In("info", []string{"debug", "info", "warning","error","fatal"})

        configuration.Mysql.User = cfg.Section("mysql").Key("username").String()
        configuration.Mysql.Password = cfg.Section("mysql").Key("password").String()
        configuration.Mysql.Host = cfg.Section("mysql").Key("host").String()
        configuration.Mysql.DBname = cfg.Section("mysql").Key("dbname").String()
        configuration.Redis.Host = cfg.Section("redis").Key("host").String()

        configuration.WX.CorpId = cfg.Section("wx").Key("wxcorpid").String()
        configuration.WX.AppSecret = cfg.Section("wx").Key("wxappsecret").String()
        configuration.WX.ContactSecret = cfg.Section("wx").Key("wxcontactsecret").String()
        configuration.WX.AgentId = cfg.Section("wx").Key("agentid").String()
        configuration.WX.RecycleEventsToken = cfg.Section("wx").Key("token").String()
        configuration.WX.RecycleEventsAesKey = cfg.Section("wx").Key("aeskey").String()

        configuration.YS.AppKey = cfg.Section("ys").Key("appkey").String()
        configuration.YS.AppSecret = cfg.Section("ys").Key("appsecret").String()

        configuration.QCC.AppKey = cfg.Section("qcc").Key("appkey").String()
        configuration.QCC.AppSecret = cfg.Section("qcc").Key("appsecret").String()

        configuration.Admin = cfg.Section("admin").Key("username").String()
        return configuration
}

func NewOpts()*Opts{
  opt := Opts{}
  opts.New(&opt).Name("wxcrm").PkgRepo().Version(version).Parse()
  ValidateOpts(&opt)
  return &opt
}


func ValidateOpts(c *Opts){
    var cfg *Obj
    if c.ConfigFile == ""{
       c.ConfigFile = "wxcrm.conf"
       cfg = ConfigParse(c.ConfigFile)
    }else{
       cfg = ConfigParse(c.ConfigFile)
    }
	if c.Port == 0 {
		c.Port = cfg.Port
	}
	if c.ServerName == "" {
		c.ServerName = cfg.ServerName
	}
        if c.NotificationUsers == ""{
                c.NotificationUsers = cfg.NotificationUsers
        }           
        if c.RedisHost == ""{
                c.RedisHost = cfg.Redis.Host
        }
        if c.LogFile == ""{
                c.LogFile = cfg.LogFile
        }
        if c.LogLevel == ""{
                c.LogLevel = cfg.LogLevel
        }
        if c.MysqlHost == ""{
                c.MysqlHost = cfg.Mysql.Host
        }
        if c.MysqlUser == ""{
                c.MysqlUser = cfg.Mysql.User
        }
        if c.MysqlPassword == ""{
                c.MysqlPassword = cfg.Mysql.Password
        }
        if c.DatabaseName == ""{
                c.DatabaseName = cfg.Mysql.DBname
        }
        if c.WXCorpId == ""{
                c.WXCorpId = cfg.WX.CorpId
        }
        if c.WXAppSecret == ""{
                c.WXAppSecret =  cfg.WX.AppSecret
        }    
        if c.WXContactSecret == ""{
                c.WXContactSecret = cfg.WX.ContactSecret
        }
        if c.WXAgentId == ""{
                c.WXAgentId = cfg.WX.AgentId
        }
        if c.WXRecycleEventsToken == ""{
                c.WXRecycleEventsToken = cfg.WX.RecycleEventsToken
        }
        if c.WXRecycleEventsAesKey == ""{
                c.WXRecycleEventsAesKey = cfg.WX.RecycleEventsAesKey
        }
        if c.YSAppKey == ""{
                c.YSAppKey = cfg.YS.AppKey
        }
        if c.YSAppSecret == ""{
                c.YSAppSecret = cfg.YS.AppSecret
        }
        if c.QCCAppKey == ""{
                c.QCCAppKey =  cfg.QCC.AppKey
        }
        if c.QCCAppSecret == ""{
                c.QCCAppSecret = cfg.QCC.AppSecret
        }
        if c.AdminUser == ""{
                c.AdminUser = cfg.Admin
        }  
        if c.SrcDir == ""{
                c.SrcDir = cfg.SrcDir
        }     
}
