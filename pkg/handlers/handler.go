package handlers
import(
  "strings"
  "wxcrm/pkg/backend"
  "wxcrm/pkg/common"
  "wxcrm/pkg/common/log"
  "wxcrm/pkg/service"
)


type HandlerVar struct{
    Opts     *common.Opts
    Logger   *log.Logger 
    DB       *backend.DB
    QCC      *service.QCC  
    WX       *service.WX  
    YS       *service.YS  
    Reporter *service.Reporter
    Redis    *common.Redis
    CustomerCheckChann chan  string   //定期检查客户认证信息
    CustomerSyncChann  chan  string    //同步客户到用友系统
    // UpdateCustomerCheckInfoChann  chan  string 
    UserInfoSyncToRD chan  string  //缓存到redis
    // UpdateYSCustomerchan chan string  //更新ERP customer 信息
    CodeRecords map[string]string
}

var (
    CheckChann = make(chan string,10)
    SyncChann =  make(chan string,100)
    // UpdateCheckInfoChann =  make(chan string,10)
    UserSyncRD = make(chan string,10)
    // UpdateYSCustomerchan = make(chan string,20)
)
func NewHandlerVar(cfg *common.Opts,logger *log.Logger)*HandlerVar{
    redis := common.NewRedis(cfg.RedisHost,logger)
    database  := backend.NewDB(cfg.MysqlHost,cfg.MysqlUser,cfg.MysqlPassword,cfg.DatabaseName,logger)
    qcc  := service.NewQCC(cfg.QCCAppKey,cfg.QCCAppSecret,logger)
    wx   := service.NewWX(cfg,redis,logger)
    ys   := service.NewYS(cfg.YSAppKey,cfg.YSAppSecret,logger)
    reporter := service.NewReporter(logger,database,cfg,ys,qcc,SyncChann)
    userids := strings.Split(cfg.AdminUser," ")
    code := map[string]string{}
    database.InitAdmin(userids)
    return &HandlerVar{
        Opts: cfg,
        Logger: logger,
        DB: database,
        QCC: qcc,
        WX: wx,
        YS: ys,
        Reporter: reporter,
        Redis: redis,
        CustomerCheckChann: CheckChann,
        CustomerSyncChann:  SyncChann,
        // UpdateCustomerCheckInfoChann:  UpdateCheckInfoChann,
        UserInfoSyncToRD: UserSyncRD,
        CodeRecords: code,
        // UpdateYSCustomerchan: UpdateYSCustomerchan,
    }
}


