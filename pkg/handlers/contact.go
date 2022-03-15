package handlers
import(
	"os"
  "bytes"
  "io/ioutil"
	"time"
	"strconv"
  "net/http"
  "wxcrm/pkg/backend"
  "wxcrm/pkg/common"
  // "strconv"
  // "wxcrm/pkg/service"
  "github.com/gin-gonic/gin"
)

//Here are Contacts,QCC search

func (h *HandlerVar)Contacts(c *gin.Context) {
	var recv backend.Contactor
    user := ""
	action := c.Query("action")

	corpcode := c.Query("corpcode")
	contactcode := c.Query("contactcode")
	c.ShouldBindJSON(&recv)
	h.Logger.Debugln("Contacts recv: ", recv)
	if cookie, err := c.Cookie("wxcrm_username");err == nil {
		user = cookie
	}else{
		user =  c.GetString("user")
	}
	if corpcode ==""{
		c.JSON(http.StatusOK,gin.H{"status": 400,"ErrMsg":"corpcode为空"})
		// c.AbortWithStatus(http.StatusOK)
		return
	}
	if action == "update"{
		// singlecontactor := h.DB.ViewWechatContactor(recv.Wechatid)
		// if  singlecontactor.ContactCode != recv.ContactCode && singlecontactor.Name != ""{
		// 	h.Logger.Infoln("企业微信联系人已经绑定到其它联系人了！")
		// 	c.JSON(http.StatusOK,gin.H{"status":400,"ErrMsg":"企业微信联系人已经绑定到其它联系人了！"})
		// 	return
		// }

		if recv.Wechatid !="" {
	    	wechatContactor := h.DB.ViewWechatContactor(recv.Wechatid)
	    	if len(wechatContactor) > 3{
	    	    h.Logger.Debugln("Name ",wechatContactor[0].Name)
	    		c.JSON(http.StatusOK,gin.H{"status":400,"ErrMsg":"企业微信联系人已添加3次！"})
	    		c.AbortWithStatus(http.StatusOK)
	    		return
	    	}
		}

		if err := h.DB.UpdateContactor(&recv); err == nil {
			h.CustomerSyncChann <- recv.CorpCode
			h.AddOperation(user,recv.CorpCode,"更新","联系人",recv.Name)
			c.JSON(http.StatusOK, gin.H{"status": 200})
		}else {
			h.Logger.Errorln(err)
			c.JSON(http.StatusOK, gin.H{"status": 400,"ErrMsg":err})
		}
		c.AbortWithStatus(http.StatusOK)
	}

	if action == "add" {
		recv.CorpCode = corpcode
		if recv.Wechatid !="" {
	    	wechatContactor := h.DB.ViewWechatContactor(recv.Wechatid)
	    	if len(wechatContactor) >= 3{
	    	    h.Logger.Debugln("Name ",wechatContactor[0].Name)
	    		c.JSON(http.StatusOK,gin.H{"status":400,"ErrMsg":"企业微信联系人已添加3次！"})
	    		c.AbortWithStatus(http.StatusOK)
	    		return
	    	}
		}


	    recv.ContactCode = "H" + common.GenUid()
	    if recv.Wechatid !=""{
	    	recv.Iswx = "1"
	    	exuserinfo,err := h.WX.GetExUserInfo(recv.Wechatid)
	        if err != nil{
	            h.Logger.Errorln(err)
	        }
	    	recv.Logo = exuserinfo.ExternalContact.Avatar
        }

        if err := h.DB.AddContactor(&recv); err == nil {
			h.AddOperation(user,corpcode,"增加","联系人",recv.Name)
     //添加相对应的 联系人 销售对应表
			var cop backend.ContactUserprincipal 
		    cop.ContactCode = recv.ContactCode
		    cop.Ownerid = user
		    cop.Createid = user
		    if err := h.DB.AddContactUserprincipal(&cop);err !=nil{
		    	h.Logger.Errorln(err)
		    }
		   h.CustomerSyncChann <- recv.CorpCode
			c.JSON(http.StatusOK, gin.H{"status": 200})
		}else {
			h.Logger.Errorln(err)
			c.JSON(http.StatusOK, gin.H{"status": 400,"ErrMsg":err})
		}
		c.AbortWithStatus(http.StatusOK)
	}
	if action == "view" {
		if corpcode ==""{
			c.JSON(http.StatusOK,gin.H{"status": 400,"ErrMsg":"corpcode为空"})
			// c.AbortWithStatus(http.StatusOK)
			return
		}
		contacts := h.DB.ViewCustomerContact(corpcode);  //id是customer id
		c.JSON(http.StatusOK,gin.H{"Contacts": contacts,"status": 200})
		c.AbortWithStatus(http.StatusOK)
	}
	if action == "info" {
		if  contactcode ==""{
			c.JSON(http.StatusOK,gin.H{"status": 400,"ErrMsg":"contactcode为空"})
			c.AbortWithStatus(http.StatusOK)
		}
		contact := h.DB.ViewIdContactor(contactcode);  
		c.JSON(http.StatusOK,gin.H{"Contact": contact,"status": 200})
		c.AbortWithStatus(http.StatusOK)
	}
	if action == "delete"{
		if  contactcode ==""{
			c.JSON(http.StatusOK,gin.H{"status": 400,"ErrMsg":"contaccode为空"})
			c.AbortWithStatus(http.StatusOK)
		}
		contactinfo := h.DB.ViewIdContactor(contactcode)
		if err := h.DB.DelContactor(contactcode); err != nil{
            h.Logger.Errorln(err)
            c.JSON(http.StatusOK, gin.H{"status": 400,"ErrMsg":err})
        }else{
        	h.CustomerSyncChann <- contactinfo.CorpCode
        	h.AddOperation(user,contactinfo.CorpCode,"删除","联系人",contactinfo.Name)
            c.JSON(http.StatusOK, gin.H{"status": 200})
        }
        c.AbortWithStatus(http.StatusOK)
	}
} 




//企查查查询
func (h *HandlerVar)QSearcher(c *gin.Context){
	action := c.Query("action")
	keyword := c.Query("keyword")
	corpcode := c.Query("corpcode")
   user,_ := c.Cookie("wxcrm_username")
   if user == ""{
    	user = c.GetString("user")
   }

	if action == "fuzzy"{
		if keyword == ""{
			c.JSON(http.StatusOK,gin.H{"status": 400,"ErrMsg":"keyword not exists"})
			c.AbortWithStatus(http.StatusNotAcceptable)
			return
		}
       recv,err := h.QCC.QFuzzySearch(keyword)
       if err != nil{
       	 h.Logger.Errorln(err)
       	 c.JSON(http.StatusOK,gin.H{"status": 400,"ErrMsg": err});
        }else{
       	if len(recv.Result) > 5{
       		c.JSON(http.StatusOK,gin.H{"result":recv.Result[0:5],"status": 200})
       	}else{
           	c.JSON(http.StatusOK,gin.H{"result":recv.Result,"status": 200})
           }
       }
       c.AbortWithStatus(http.StatusOK)
       return
	}
	if corpcode == ""{
	 	c.JSON(http.StatusOK,gin.H{"status":400,"ErrMsg":"参数错误！"})
	 	c.AbortWithStatus(http.StatusOK)
	 	return
	}
	//失信核查
	if action == "shixin"{
		shixins := h.DB.ViewCustomerShixin(corpcode)
		c.JSON(http.StatusOK,gin.H{"status":200,"shixins":shixins})
		c.AbortWithStatus(http.StatusOK)
		return
	}
    //税收核查
	if action == "tax"{
		taxs := h.DB.ViewCustomerTaxIllegal(corpcode)
		c.JSON(http.StatusOK,gin.H{"status":200,"taxs":taxs})
		c.AbortWithStatus(http.StatusOK)
		return 
	}
    // 重大违法核查
	if action == "serious"{
		serious := h.DB.ViewCustomerSeriousIllegal(corpcode)
		c.JSON(http.StatusOK,gin.H{"status":200,"serious":serious})
		c.AbortWithStatus(http.StatusOK)
		return
	}
    // 行政处罚核查
	if action == "penalty"{
		penalties := h.DB.ViewCustomerAdminPenalty(corpcode)
		c.JSON(http.StatusOK,gin.H{"status":200,"penalties":penalties})
		c.AbortWithStatus(http.StatusOK)
		return
	}
    //基础信用报告
	if action == "report"{
		customerinfo := h.DB.ViewCustomerInfo(corpcode)
		if data,err := os.ReadFile(h.Opts.SrcDir+"/"+corpcode+".pdf");err == nil{
			tmpdata := bytes.NewBuffer(data)
		    result,err := h.WX.UploadFile(tmpdata,customerinfo.CorpName+".pdf")
		    if err !=nil{
		        h.Logger.Errorln(err)
		        c.JSON(http.StatusOK,gin.H{"status":400,"ErrMsg":err})
		        return
		    }
		    if _,err := h.WX.SendFileMsg(result.MediaID,user);err != nil{
		        h.Logger.Errorln(err)
		        c.JSON(http.StatusOK,gin.H{"status":400,"ErrMsg":err})
		        return
		    }
		    c.JSON(http.StatusOK,gin.H{"status":200,"ErrMsg":"基础信用报告已发送到用户企业微信"})
		    return
		}

		go h.BasicReportSender(user,customerinfo)
		c.JSON(http.StatusOK,gin.H{"status":200,"ErrMsg":"基础信用报告正在生成中，稍后将发送到用户企业微信"})
		return
	}
	c.JSON(http.StatusOK,gin.H{"status":400,"ErrMsg":"参数错误！"})

}



//报表
func (h *HandlerVar)GenReporter(c *gin.Context){
	user := ""
	from := c.Query("from")
	to := c.Query("to")
	if cookie, err := c.Cookie("wxcrm_username");err == nil {
		user = cookie
	}else{
		user =  c.GetString("user")
	}

   tnow := time.Now().Unix()
	start_time,err := strconv.Atoi(from)
	if err != nil {
		// h.Logger.Errorln(err)
		start_time = int(tnow - 30 * 86400)
	}
	end_time,err := strconv.Atoi(to)
	if err != nil{
		// h.Logger.Errorln(err)
		end_time = int(tnow)
	}
	data,err := h.Reporter.GenReport(user,strconv.Itoa(start_time),strconv.Itoa(end_time))
	if err !=nil{
		h.Logger.Errorln(err)
		c.JSON(http.StatusOK,gin.H{"status":400,"ErrMsg":err})
		c.AbortWithStatus(http.StatusOK)
		return
	}
	result,err := h.WX.UploadFile(data,"")
	if err !=nil{
		h.Logger.Errorln(err)
		c.JSON(http.StatusOK,gin.H{"status":400,"ErrMsg":err})
		c.AbortWithStatus(http.StatusOK)
	}
	if _,err := h.WX.SendFileMsg(result.MediaID,user);err != nil{
		h.Logger.Errorln(err)
		c.JSON(http.StatusOK,gin.H{"status":400,"ErrMsg":err})
		c.AbortWithStatus(http.StatusOK)
	}
	h.Logger.Infoln("报表生成并已发送到用户: "+ user +" 企业微信！")
	c.JSON(http.StatusOK,gin.H{"status":200})
	c.AbortWithStatus(http.StatusOK)
	return
}

//批量导入客户
func (h *HandlerVar)ImportExcel(c *gin.Context){
	var user string
	if cookie, err := c.Cookie("wxcrm_username"); err == nil {
		user = cookie
	}else{
		user =  c.GetString("user")
		h.Logger.Debugln("GetCustomer get user from map:",user)
	}
	if c.Request.Method == "POST"{
			file, err := c.FormFile("file")
			tflag := c.Query("tflag")  // tflag = sea
			if err != nil{
				h.Logger.Errorln(err)
			}
			h.Logger.Debugln("上传文件名：",file.Filename)
			f,err := file.Open()
			defer f.Close()
			if err != nil{
				h.Logger.Errorln(err)
			}else{
				h.Logger.Debugln("user: ",user)
				if user == ""{
					h.Logger.Errorln("user is nil")
					return
				}
				go h.UploadXlsxFile(f,user,tflag)
			}

			// Upload the file to specific dst.
			// c.SaveUploadedFile(file, dst)

			c.JSON(http.StatusOK, gin.H{"status":200,"Msg":"文件已上传，正在处理中..."})
			return
	}
	if c.Request.Method == "GET"{
		//     file,err := os.Open("/mnt/example.xlsx")
		//     if err != nil{
		//     	h.Logger.Errorln(err)
		//     }
		//     defer file.Close()
		//     fstat,err := file.Stat()
		//     if err != nil{
		//     	h.Logger.Errorln(err)
		//     }
		// 	reader := file  
		// 	contentLength := fstat.Size()
		// 	contentType := "application/octet-stream"
		// 	extraHeaders := map[string]string{
		// 		"Content-Disposition": `attachment; filename="example.xlsx"`,
		// 	}
		// 	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
		// 	return
		// }
	    //  h.Logger.Debugln("UpLoad not get any file!")
		h.Logger.Debugln("upload get file!")
        tmpdata,err := ioutil.ReadFile("/mnt/example.xlsx")
        if err != nil{
        	h.Logger.Errorln(err)
        }
        data := bytes.NewBuffer(tmpdata)
		result,err := h.WX.UploadFile(data,"example.xlsx")
		if err !=nil{
			h.Logger.Errorln(err)
			c.JSON(http.StatusOK,gin.H{"status":400,"ErrMsg":err})
			return
		}
		if _,err := h.WX.SendFileMsg(result.MediaID,user);err != nil{
			h.Logger.Errorln(err)
			c.JSON(http.StatusOK,gin.H{"status":400,"ErrMsg":err})
			return
		}
		c.JSON(http.StatusOK,gin.H{"status":200})
		return
    }
}





