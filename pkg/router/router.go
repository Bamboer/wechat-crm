package router
import (
  "io"
  "os"
 "net/http"
  "strconv"
  "context"
  // "wxcrm/pkg/backend"
  "wxcrm/pkg/common"
  "wxcrm/pkg/common/log"
  "wxcrm/pkg/handlers"
  "github.com/gin-gonic/gin"
  // "wxcrm/pkg/service"
)

func Run(cfg *common.Opts,logger *log.Logger){
    f,_ := os.OpenFile(cfg.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    gin.DefaultWriter = io.MultiWriter(f)
    handlervar := handlers.NewHandlerVar(cfg,logger)
    ctx := context.Background()
    go handlervar.OnSync(ctx)
    go handlervar.Notficationer(ctx)
    // handlervar.AddWXUsers()
    // handlervar.Reporter.UpLoadWithFile("/mnt/import_excel.xlsx")
    r := gin.New()
//    r.LoadHTMLGlob("templates/*")
   logger.Debugln("source directory: ",cfg.SrcDir)
   r.StaticFS("/api/src", http.Dir(cfg.SrcDir))
    r.Use(gin.Recovery())
//    r.Use(Cors(handlervar))
    r.GET("/api/test/user",handlervar.UserTest)

    r.GET("/api/wxevents",handlervar.WXExUserRecv)
    r.POST("/api/wxevents",handlervar.WXExUserRecv)
//home page
    r.GET("/", handlervar.StartPage) 
//    r.GET("/:name", handlervar.StartPage)   
    r.Use(Middler(handlervar))
   
    r.GET("/api/customer",handlervar.GetCustomer)
    r.POST("/api/customer",handlervar.PostCustomer)

    r.GET("/api/user",handlervar.User)
    r.POST("/api/user",handlervar.User)

    r.GET("/api/wx",handlervar.WXUserinfo)
    r.GET("/api/qcc",handlervar.QSearcher)
    r.GET("/api/check",handlervar.Checker)
    r.POST("/api/principal",handlervar.CustomerUserPrincipal)
    r.GET("/api/principal",handlervar.CustomerUserPrincipal)
    r.GET("/api/jsticket",handlervar.GetJsticket)
    r.GET("/api/auth",handlervar.Auth) 
    r.GET("/api/ysinfo",handlervar.YSInfo) 
    r.GET("/api/operation",handlervar.Operation)
    r.GET("/api/report",handlervar.GenReporter)
    r.GET("/api/customerprojectinfo",handlervar.CustomerProjectsInfo)
    r.POST("/api/upload",handlervar.ImportExcel)
    r.GET("/api/upload",handlervar.ImportExcel)
    r.GET("/api/seacustomer",handlervar.SeaCustomer)
    r.POST("/api/seacustomer",handlervar.SeaCustomer)

    r.GET("/api/record",handlervar.FollowRecords)
    r.POST("/api/record",handlervar.FollowRecords)

    r.GET("/api/contact",handlervar.Contacts)
    r.POST("/api/contact",handlervar.Contacts)

    r.GET("/api/product",handlervar.Product)
    r.POST("/api/product",handlervar.Product)


    
    r.GET("/api/project",handlervar.Project)
    r.POST("/api/project",handlervar.Project)

    r.GET("/api/agreement",handlervar.Agreement)
    r.POST("/api/agreement",handlervar.Agreement)           
  
    r.Use(EndErr(handlervar))

    logger.Infoln("Start server...")
    r.Run( ":" + strconv.Itoa(cfg.Port))
}

