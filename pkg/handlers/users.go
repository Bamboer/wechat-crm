package handlers
import(
  "fmt"
  "net/http"
  "strings"
  "strconv"
  "wxcrm/pkg/backend"
  "github.com/gin-gonic/gin"
)

func (h *HandlerVar)StartPage(c *gin.Context) {
 if name := c.Param("name");name != ""{
     c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", name))
     c.Writer.Header().Add("Content-Type", "application/octet-stream")
     c.File("./"+name)
     // c.AbortWithStatus(http.StatusOK)
  } 
//   var user string 
//   if cookie, err := c.Cookie("wxcrm_username"); err == nil {
// 	user = cookie
// 	h.Logger.Debugln("StartPage cookie: ", cookie)
//   }else{
// 	user =  c.MustGet("user").(string)
//   }

//   if user != "" {
// 	userinfo := h.DB.ViewUser(user)
// 	products := h.DB.ViewProducts()
// 	c.HTML(http.StatusOK, "index.html", gin.H{"username": userinfo.Name,"Products": products})
// //    c.AbortWithStatus(http.StatusOK)
// 	} else {
// 		h.Logger.Debugln("cookie not set: ", user)
// 		c.HTML(http.StatusOK, "index.html", "")
//      } 
}

func(h *HandlerVar)User(c *gin.Context){
	action := c.Query("action")
	userid := c.Query("userid")
	cookie, err := c.Cookie("wxcrm_username")
	dmaiuser := &backend.DmaiUser{}
	if err != nil{
		h.Logger.Debugln("Not get wxcrm_username and try to get user key that from context map.")
		cookie =  c.GetString("user")
	}

	if userid != ""{
		dmaiuser := h.DB.ViewUser(userid)
    h.Logger.Debugln("get a dmai user info :",dmaiuser)
    c.JSON(http.StatusOK,gin.H{"Dmaiuser": dmaiuser, "status":200})
    return
	}

	if action == "update"{
		c.ShouldBindJSON(dmaiuser)
    h.Logger.Debugln("add dmai user :",dmaiuser)
		err := h.DB.UpdateUser(dmaiuser)
    if err == nil{
    	c.JSON(http.StatusOK,gin.H{"status": 200})
    }else{
    	c.JSON(http.StatusOK,gin.H{"status": 500,"ErrMsg": err})
    }
    return 
	}else if action == "view"{
		var dmaiusers []backend.DmaiUser
		dmaiusers = h.DB.ViewUsers()
		// h.Logger.Debugln("list dmai users: ",dmaiusers)
    c.JSON(http.StatusOK,gin.H{"Dmaiusers": dmaiusers,"status": 200})
	}else{
	  dmaiuser := h.DB.ViewUser(cookie)
		if dmaiuser.Name == ""{
			c.JSON(http.StatusOK, gin.H{"status": 500,"CookieName": "wxcrm_username","Userinfo":"","ErrMsg": "server excute error"})
			return
		}
		h.Logger.Debugln("get a dmai user info :",dmaiuser)
	  c.JSON(http.StatusOK,gin.H{"status": 200,"CookieName": "wxcrm_username","Userinfo":dmaiuser})
	}
}

func (h *HandlerVar)Auth(c *gin.Context){
    eurl := c.Query("url")
    if eurl != ""{
        WXOAuth2URL := "https://open.weixin.qq.com/connect/oauth2/authorize?appid=" + h.Opts.WXCorpId + "&redirect_uri="+ eurl + "&response_type=code&scope=snsapi_base&state=#wechat_redirect"
        h.Logger.Infoln("JSON return","emit from: ",c.GetHeader("X-Requested-With"))
        h.Logger.Infoln("Auth URL: ",WXOAuth2URL)
        c.JSON(http.StatusOK,gin.H{"URL": WXOAuth2URL,"status":200,"ErrMsg":"not found cookie,try to authorization with oauth2 url"})
        c.AbortWithStatus(http.StatusOK)
        return
    }else{
        c.JSON(http.StatusOK,gin.H{"status":400,"ErrMsg":"请设定需要重定向的url"})
    }
}



//创建用户时，需要设置应用课可见范围为最大范围
func(h *HandlerVar)AddWXUser(username string){
   users,err := h.WX.GetUserlist()
   if err != nil{
      h.Logger.Errorln(err)
      return
   }
   h.Logger.Debugln("用户名: ",username)
   for _,v := range users.Userlist{
      if strings.HasPrefix(v.Name,username){
        h.Logger.Debugln("找到查询name: ",v.Name)
      	dmaiuser := &backend.DmaiUser{}
      	dmaiuser1 := h.DB.ViewUser(v.Userid)
        userinfo,err := h.WX.GetWXAccessUserDetail(v.Userid)
        if err != nil{
            h.Logger.Errorln(err)
            return
        }
        dmaiuser.Name = userinfo.Name 
        dmaiuser.Gender =  userinfo.Gender
        dmaiuser.OpenUserid =  userinfo.OpenUserid
        dmaiuser.Userid = v.Userid
        dmaiuser.Logo  = userinfo.Avatar 
        // dmaiuser.Department = userinfo.Department
        dmaiuser.Position =  userinfo.Position
        dmaiuser.Mobile =  userinfo.Mobile
        dmaiuser.Email =  userinfo.Email
        dmaiuser.Alias =  userinfo.Alias
        dmaiuser.Address =  userinfo.Address
        if len(userinfo.Department) >0{
            h.Logger.Debugln("User department: ",userinfo.Department)
            id := strconv.Itoa(userinfo.Department[len(userinfo.Department)-1])
            if err != nil{
                h.Logger.Errorln(err)
            }else{
                if departmentinfo,err := h.WX.GetDepartmentInfo(id);err == nil{
                    h.Logger.Debugln("Department info: ",departmentinfo)
                    dmaiuser.Department = departmentinfo.Department[len(departmentinfo.Department)-1].Name 
                }
            }
        }
        if dmaiuser1.Userid != ""{
            h.DB.UpdateUser(dmaiuser)
        }else{
            h.DB.AddUser(dmaiuser)
        }
        break
      }else{
        h.Logger.Debugln("未找到用户： ",username)
      }
   }
}


func(h *HandlerVar)AddWXUsers(){
    userlist := []string{"丘螣称","何保君","傅尧","刘可","庄晓新","张小俊","张文岳","彭龙","李礼","杨晓宾","梁康森","王芳媛","申磊","米俊杰","苗迪","邓超","郭岩","郭运祥","陈强","魏启伟","黄国安","黄榕莺"}
    for _,v := range userlist{
        h.AddWXUser(v)
    }
}















