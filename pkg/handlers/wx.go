
package handlers
import(
	"encoding/json"
  "net/http"
  "io/ioutil"
  "encoding/xml"
  "wxcrm/pkg/common/wxbizmsgcrypt"
  "github.com/gin-gonic/gin"
  "wxcrm/pkg/backend"
  "wxcrm/pkg/service"
)

type WXExMember struct{
   Name  string  `json:"text"`
   ExUserId  string 
}

func (h *HandlerVar)WXUserinfo(c *gin.Context){
	exuserid  := c.Query("exuserid")
	var user string
	if cookie, err := c.Cookie("wxcrm_username"); err == nil {
		user = cookie
	}else{
		user =  c.GetString("user")
		h.Logger.Debugln("WX get user from map:",user)
	}
	if user == "" {
		h.Logger.Errorln("user id  is nil.---GetCustomer")
		c.JSON(http.StatusOK, gin.H{"status": 400,"ErrMsg": "userid id  is null"})
		c.AbortWithStatus(http.StatusOK)
		return
	}
	if exuserid !="" {
		customerInfo, err := h.WX.GetExUserInfo(exuserid)
		h.Logger.Debugln("WX Exuser Info: ",customerInfo)
		if err != nil{
			h.Logger.Errorln("Get Exuser info error:",err)
			c.JSON(http.StatusOK,gin.H{"status":400,"ErrMsg": err})
			c.AbortWithStatus(http.StatusOK)
		}else{
			c.JSON(http.StatusOK,gin.H{
              "ExternalUserid": customerInfo.ExternalContact.ExternalUserid,
              "text": customerInfo.ExternalContact.Name ,
              "Avatar": customerInfo.ExternalContact.Avatar,   //头像url
              "Type": customerInfo.ExternalContact.Type ,
              "Gender": customerInfo.ExternalContact.Gender,
              "FollowUser": customerInfo.FollowUser,
              "status": 200,
			})
		}
		c.AbortWithStatus(http.StatusOK)
		return
	}
	WXMemberList := []WXExMember{}
	wxlist := h.UserWXCustomer_list(user)
	h.Logger.Debugln("External WX list: ", wxlist)
	for _,v := range wxlist{
		WXMemberList = append(WXMemberList,WXExMember{Name: v.ExternalContact.Name,ExUserId: v.ExternalContact.ExternalUserid})
	}
	c.JSON(http.StatusOK, gin.H{"status": 200,"EXUsers": WXMemberList})
	c.AbortWithStatus(http.StatusOK)
	return
}


// //企业微信 客户联系人变更监控
// type MsgContent struct {
// 	ToUsername   string `xml:"ToUserName";gorm:"size:255;comment:操作对象"`
// 	FromUsername string `xml:"FromUserName";gorm:"size:255;comment:操作者"`
// 	CreateTime   uint32 `xml:"CreateTime";gorm:"size:255;comment:创建时间"`
// 	MsgType      string `xml:"MsgType";gorm:"size:255;comment:信息类型"`
// 	Event        string `xml:"Event";gorm:"size:255;comment:什么事件"`
// 	ChangeType   string `xml:"ChangeType";gorm:"size:255;comment:改变类型id"`
// 	UserID       string `xml:"UserID";gorm:"size:255;comment:用户id"`
// 	ExternalUserID   string `xml:"ExternalUserID";gorm:"size:255;comment:外部联系人id"`
//    State        string `xml:"State";gorm:"size:255;comment:交接状态"`
//    WelcomeCode  string `xml:"WelcomeCode";gorm:"size:255;comment:欢迎码"`
// //客户联系人交接	
// 	FailReason   string `xml:"FailReason";gorm:"size:255;comment:客户联系人交接失败原因"`
// 	ChatId       string `xml:"ChatId";gorm:"size:255;comment:群聊id"`
// 	UpdateDetail string `xml:"UpdateDetail";gorm:"size:255;comment:改变详情"`
// 	JoinScene    string `xml:"JoinScene";gorm:"size:255;comment:加入"`
// 	QuitScene    string `xml:"QuitScene";gorm:"size:255;comment:推出"`
// 	MemChangeCnt string `xml:"MemChangeCnt";gorm:"size:255;comment:成员改变多少"`
// }

func (h *HandlerVar)WXExUserRecv(c *gin.Context){
	wxcpt := h.Wxcpt()
  verifyMsgSign := c.Query("msg_signature")

  verifyTimestamp := c.Query("timestamp")
  verifyNonce := c.Query("nonce")
  h.Logger.Debugln("WX EXuser notification events, Method:",c.Request.Method,verifyMsgSign,verifyTimestamp,verifyNonce)
	if c.Request.Method == "GET"{
	  verifyEchoStr := c.Query("echostr")
	  echoStr, cryptErr := wxcpt.VerifyURL(verifyMsgSign, verifyTimestamp, verifyNonce, verifyEchoStr)
	  if nil != cryptErr{
	   	h.Logger.Errorln("verifyUrl fail", cryptErr)
	   	c.String(http.StatusNotAcceptable,"None")
	  }else{
	  	h.Logger.Debugln("verify url success echostr: ",string(echoStr))
	  	c.String(http.StatusOK,string(echoStr))
	  }
	  c.AbortWithStatus(http.StatusOK)
	}else{
	 data,err := ioutil.ReadAll(c.Request.Body)
    if err != nil{
       h.Logger.Errorln("read notify events error:",err)
    }
    msg,cryptErr := wxcpt.DecryptMsg(verifyMsgSign, verifyTimestamp, verifyNonce, data)
	 if nil != cryptErr {
			h.Logger.Errorln("DecryptMsg fail", cryptErr)
	 }
		h.Logger.Debugln("Not Decode msg: ",msg)
		//解析
		var msgContent backend.Wxlog
		if err := xml.Unmarshal(msg, &msgContent);nil != err {
		    h.Logger.Errorln("Unmarshal fail")
		}else {
		 	h.Logger.Debugln("struct", msgContent)
		 	if err := h.DB.AddWxlog(&msgContent);err != nil{
		 		h.Logger.Errorln(err)
		 	}

		 	if msgContent.Event == "change_external_contact"{
		 		go h.UserWXCustomer_listUpdate(msgContent.UserID)
		 	}
	  }
		c.String(http.StatusOK,"okay")
		c.AbortWithStatus(http.StatusOK)
	}
}

func(h *HandlerVar)Wxcpt()*wxbizmsgcrypt.WXBizMsgCrypt{
   return wxbizmsgcrypt.NewWXBizMsgCrypt(h.WX.RecycleEventsToken, h.WX.RecycleEventsAesKey, h.WX.CorpID, wxbizmsgcrypt.XmlType)
}


func (h *HandlerVar)GetJsticket(c *gin.Context){
	if jsticket,err := h.WX.GetJSTicket();err != nil{
		h.Logger.Errorln(err)
		c.JSON(http.StatusOK,gin.H{"jsticket":"","status":400,"ErrMsg": err})
	}else{
		c.JSON(http.StatusOK,gin.H{"jsticket":jsticket,"status": 200})
	}
	c.AbortWithStatus(http.StatusOK)
}


func (h *HandlerVar)UserWXCustomer_list(username string) []*service.ExUserInfo {
	var WXCustomerList []*service.ExUserInfo
	if result := h.Redis.CacheRe(username+"contact");result != nil{
      if v,ok := result.([]byte);ok{
            json.Unmarshal(v,&WXCustomerList)
      }
      h.Logger.Debugln("contact list 从redis中拿数据！")
   }
   if len(WXCustomerList) >0 {
   	return WXCustomerList
   }
	data, _ := h.WX.GetExUserList(username)
	for _, v := range data {
		customerInfo, err := h.WX.GetExUserInfo(v)
		if err == nil {
			WXCustomerList = append(WXCustomerList, customerInfo)
		} else {
			h.Logger.Errorln(err)
		}
	}
   if err := h.Redis.CacheSet(username+"contact",WXCustomerList);err != nil{
      h.Logger.Errorln(err)
   }
	return WXCustomerList
}

func (h *HandlerVar)UserWXCustomer_listUpdate(username string){
	if userinfo :=h.DB.ViewUser(username);userinfo.Name != ""{
	   var WXCustomerList []*service.ExUserInfo
	   data, _ := h.WX.GetExUserList(username)
	   for _, v := range data {
		  customerInfo, err := h.WX.GetExUserInfo(v)
		  if err == nil {
			WXCustomerList = append(WXCustomerList, customerInfo)
		  } else {
			h.Logger.Errorln(err)
		  }
	   }
  if err := h.Redis.CacheSet(username+"contact",WXCustomerList);err != nil{
      h.Logger.Errorln(err)
  }
}
}


