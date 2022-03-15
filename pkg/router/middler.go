package router
import (
  "strings"
  "strconv"
  "net/http"
  "wxcrm/pkg/backend"
  "wxcrm/pkg/handlers"
  "github.com/gin-gonic/gin"
)
 

//用户验证
func Middler(handler *handlers.HandlerVar) gin.HandlerFunc {
	return func(c *gin.Context) {
		//"https://open.weixin.qq.com/connect/oauth2/authorize?appid=ww84222d5b0d1744ac&redirect_uri=https://tcrm-iti.dm-ai.com/customer&response_type=code&scope=snsapi_base&state=abc123#wechat_redirect"
        code := c.Query("code")
        if strings.HasPrefix(c.Request.URL.Path,"/api/test"){
            c.Next()
            c.AbortWithStatus(http.StatusOK)
            return
        }

        // if strings.HasPrefix(c.Request.URL.Path,"/api/admin"){
        //     c.Next()
        //     c.AbortWithStatus(http.StatusOK)
        //     return
        // }
        if strings.HasPrefix(c.Request.URL.Path,"/api/auth"){
            c.Next()
            c.AbortWithStatus(http.StatusOK)
            return
        }

		var dmaiuser *backend.DmaiUser
		cookie, err := c.Cookie("wxcrm_username")
		if err != nil {
			handler.Logger.Infoln("Cookie get error try to wx authorize...")
			if code != "" {
                    handler.Logger.Infoln("code not nil. code is:", code)
                    handler.Logger.Infoln("Before set cookie and rewrite Request: ",c.Request.URL.String())
                    if _, ok := handler.CodeRecords[code]; ok {
                        handler.Logger.Errorln("Code重复消费！")
                        c.JSON(http.StatusOK,gin.H{"status":405,"ErrMsg":"Code重复消费！"})
                        c.AbortWithStatus(http.StatusNotAcceptable)
                        return
                    }

                    var err error
                    cookie, err = handler.WX.GetWXAccessUser(code)
                    handler.CodeRecords[code]=""
                    if cookie == ""{
                        handler.Logger.Errorln(err.Error()+"非企业成员！")
                        c.JSON(http.StatusOK,gin.H{"status":402,"ErrMsg":"非企业成员！"})
                        handler.Logger.Infoln("非企业成员！")
                        c.AbortWithStatus(http.StatusNotAcceptable)
                        return
                    }
                    handler.Logger.Infoln("user: ", cookie, "error: ", err)
                    if err == nil {
                        c.SetCookie("wxcrm_username", cookie, 7200, "/", handler.Opts.ServerName, false, true)
                        // c.SetCookie("wxcrm_username", cookie, 7200, "/", "192.168.4.126", false, true)
                        c.Set("user",cookie) 
                    }
			}else {
                c.JSON(http.StatusOK,gin.H{"status":401,"ErrMsg":"未认证用户，请验证用户！"})
                c.AbortWithStatus(http.StatusOK)
                return
			}
		}else{
            c.SetCookie("wxcrm_username", cookie, 7200, "/", handler.Opts.ServerName, false, true)
            // c.SetCookie("wxcrm_username", cookie, 7200, "/", "192.168.4.126", false, true)
        }

        handler.Logger.Infoln("cookie is: ", cookie)
    	dmaiuser = handler.DB.ViewUser(cookie)
        userid_judge := dmaiuser.Userid
    	handler.Logger.Infoln("check dmaiuser user info from db and get user id: ", dmaiuser.Userid)
    	if len(dmaiuser.Department) == 0{
    		userinfo,err := handler.WX.GetWXAccessUserDetail(cookie)
    		if err != nil{
    		  handler.Logger.Errorln(err)
    		  c.AbortWithStatus(http.StatusNotAcceptable)
    		  return
    		}
            handler.Logger.Infoln(dmaiuser.Userid," not in database and fill user info to database which come from wx")
            dmaiuser.Name = userinfo.Name 
            dmaiuser.Gender =  userinfo.Gender
            dmaiuser.OpenUserid =  userinfo.OpenUserid
            dmaiuser.Userid = cookie
            dmaiuser.Logo  = userinfo.Avatar 
            // dmaiuser.Department = userinfo.Department
            dmaiuser.Position =  userinfo.Position
            dmaiuser.Mobile =  userinfo.Mobile
            dmaiuser.Email =  userinfo.Email
            dmaiuser.Alias =  userinfo.Alias
            dmaiuser.Address =  userinfo.Address
            if len(userinfo.Department) >0{
                handler.Logger.Debugln("User department: ",userinfo.Department)
                id := strconv.Itoa(userinfo.Department[len(userinfo.Department)-1])
                if err != nil{
                    handler.Logger.Errorln(err)
                }else{
                    if departmentinfo,err := handler.WX.GetDepartmentInfo(id);err == nil{
                        handler.Logger.Debugln("Department info: ",departmentinfo)
                        dmaiuser.Department = departmentinfo.Department[len(departmentinfo.Department)-1].Name 
                    }
                }
            }
            if managerJudge(cookie,handler){
            	dmaiuser.Manager = "true"
                
            }
            if userid_judge != ""{
                handler.DB.UpdateUser(dmaiuser)
            }else{
                handler.DB.AddUser(dmaiuser)
            }

    //admin auth		
        	if dmaiuser.Manager == "true"{
        		_,err := c.Cookie("manager_cookie")
        		if err != nil{
        			handler.Logger.Warningln("set a manager cookie for user: ",cookie)
        			c.SetCookie("manager_cookie", "manager", 7200, "/", handler.Opts.ServerName, false, true)
        		}
        	}
        }
	c.Next()
   }
}

func managerJudge(name string,handler *handlers.HandlerVar)bool{
     mangers := strings.Split(handler.Opts.AdminUser," ")
     for _,v := range mangers{
     	if name == v{
     		handler.Logger.Infoln("manager: ",v)
     		return true
     	}
     }
     return false 
}


func EndErr(handler *handlers.HandlerVar)gin.HandlerFunc{
    return func(c *gin.Context){
        c.JSON(http.StatusOK,gin.H{"status":400,"ErrMsg":"参数错误"})
    }
}


//跨域访问中间件
func Cors(handler *handlers.HandlerVar) gin.HandlerFunc {
    return func(c *gin.Context) {
        method := c.Request.Method
        origin := c.Request.Header.Get("Origin") //请求头部
        if origin != "" {
            //接收客户端发送的origin （重要！）
            c.Writer.Header().Set("Access-Control-Allow-Origin", origin) 
            //服务器支持的所有跨域请求的方法
            c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") 
            //允许跨域设置可以返回其他子段，可以自定义字段
            c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
            // 允许浏览器（客户端）可以解析的头部 （重要）
            c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers") 
            //设置缓存时间
            c.Header("Access-Control-Max-Age", "172800") 
            //允许客户端传递校验信息比如 cookie (重要)
            c.Header("Access-Control-Allow-Credentials", "true")                                                                                                                                                                                                                          
        }

        //允许类型校验 
        if method == "OPTIONS" {
            c.JSON(http.StatusOK, "ok!")
        }

        defer func() {
            if err := recover(); err != nil {
                handler.Logger.Debugln("Panic info is: %v", err)
            }
        }()

        c.Next()
    }
}











