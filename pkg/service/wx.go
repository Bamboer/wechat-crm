package service
import (
  "io"
  "fmt"
  "time"
  "bytes"
  "strconv"
  "net/url"
  "net/http"
  "io/ioutil"
  "mime/multipart"
  "encoding/json"
  "wxcrm/pkg/common"
  "wxcrm/pkg/common/log"
)


type WX struct{
  Redis   *common.Redis
  Logger  *log.Logger
  CorpID  string 
  AppSecret string
  ContactSecret string  
  AgentId       string 
  RecycleEventsToken    string  
  RecycleEventsAesKey  string

  ContactToken string //WX 通讯录token
  Token        string
  JSTicket     string 
  TokenSetTime int64
  TicketSetTime int64
}

const (
  WXTokenName   = "wxtoken"
  JSTicketName  = "jsticket"
  WXContactTokenName = "contactoken"
  // YYTokenName   = "yytoken"
  TokenTTL    = "7200"
)

func NewWX(cfg *common.Opts,redis *common.Redis,logger *log.Logger)*WX{
  return &WX{
    Redis: redis,
    Logger: logger,
    CorpID: cfg.WXCorpId,
    AppSecret: cfg.WXAppSecret,
    ContactSecret: cfg.WXContactSecret,
    AgentId:  cfg.WXAgentId,
    RecycleEventsToken: cfg.WXRecycleEventsToken,   
    RecycleEventsAesKey: cfg.WXRecycleEventsAesKey,  
  }
}

// 获取应用access token
func (wx *WX)GetWXToken()(string,error){
  tnow := time.Now().Unix()
  if tnow - wx.TokenSetTime < -10 && wx.Token != ""{
      // wx.Logger.Debugln("WXToken set time is not timeout and return this token value.")
      // wx.Logger.Debugln("WX token: ",wx.Token)
      return wx.Token,nil
  }    

  t1,err := wx.Redis.GetKeyTTL(WXTokenName)
  // wx.Logger.Debugln("Get token :",WXTokenName,"ttl: ",t1)
  if err != nil || t1 == -2{
    wxrt := &WXRT{}
    client := &http.Client{}

    // https://qyapi.weixin.qq.com/cgi-bin/gettoken

    uri,_ := url.Parse("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=" + wx.CorpID +"&corpsecret=" +wx.AppSecret )
    req,err := http.NewRequest("GET", uri.String(), nil)
    if err != nil{
      wx.Logger.Errorln(err)
      return "",err
    }

    wx.Logger.Debugln("GET WX Token url: ",req.URL.String())
  
    resq,err := client.Do(req)
    defer resq.Body.Close()
  
    b,err := ioutil.ReadAll(resq.Body)
    if err != nil {
      wx.Logger.Errorln(err)
      return "",err
    }
    if err := json.Unmarshal(b,wxrt);err != nil{
      wx.Logger.Errorln(err)
      return "",err
    }
    if wxrt.Errcode == 0 {
      if err := wx.Redis.SetKey(WXTokenName,strconv.Itoa(wxrt.Expires_in),wxrt.Access_token);err != nil{
         wx.Logger.Errorln(err)
      }

      tnow := time.Now().Unix()
      wx.Token = wxrt.Access_token
      wx.TokenSetTime = tnow + int64(wxrt.Expires_in)

      return wxrt.Access_token,nil
    }else{
      wx.Logger.Errorln(wxrt.Errmsg)
      return "",fmt.Errorf(wxrt.Errmsg)
   }
  }else{
    token,err := wx.Redis.GetKey(WXTokenName)
    if err != nil{
      wx.Logger.Errorln(err)
       return "",err
    }
    // wx.Logger.Debugln("WX token: ",token)
    return token,nil
  }
}



//获取企业微信当前用户
func (wx *WX)GetWXAccessUser(code string)(string,error){
  //https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo?access_token=ACCESS_TOKEN&code=CODE
  accessuser := &AccessUser{}
  client := &http.Client{}

  token,err := wx.GetWXToken()
  if err != nil{
    wx.Logger.Errorln(err)
    return "",err
  }
  uri,_ := url.Parse("https://qyapi.weixin.qq.com"+"/cgi-bin/user/getuserinfo?access_token="+token+"&code="+code+"&debug=1")
  req,err := http.NewRequest("GET", uri.String(), nil)
  if err != nil{
     wx.Logger.Errorln(err)
     return "",err
  }

  resq,err := client.Do(req)
  if err != nil{
     wx.Logger.Errorln(err)
     return "",err
  }

  defer resq.Body.Close()

  b,err := ioutil.ReadAll(resq.Body)
  if err != nil {
     wx.Logger.Errorln(err)
     return "",err
  }
  if err := json.Unmarshal(b,accessuser);err != nil{
     wx.Logger.Errorln(err)
     return "",err
  }
  if accessuser.Errcode == 0 {
     wx.Logger.Debugln("Current userID: ",accessuser.UserId)
     return accessuser.UserId,nil
  }else{
     wx.Logger.Errorln("Get userid Error code: ",accessuser.Errcode," msg: ",accessuser)
     return "",fmt.Errorf("get wx exuserid error.")
  }
}

//读取成员详细信息

func (wx *WX)GetWXAccessUserDetail(userid string)(*WXUserDetail,error){
  //GET https://qyapi.weixin.qq.com/cgi-bin/user/get?access_token=ACCESS_TOKEN&userid=USERID
  userdetail := &WXUserDetail{}
  client := &http.Client{}

  token,err := wx.GetWXToken()
  if err != nil{
    wx.Logger.Errorln(err)
    return nil,err
  }
  uri,_ := url.Parse("https://qyapi.weixin.qq.com"+"/cgi-bin/user/get?access_token="+token+"&userid="+userid)
  req,err := http.NewRequest("GET", uri.String(), nil)
  if err != nil{
     wx.Logger.Errorln(err)
     return nil,err
  }

  resq,err := client.Do(req)
  if err != nil{
     wx.Logger.Errorln(err)
     return nil,err
  }

  defer resq.Body.Close()

  b,err := ioutil.ReadAll(resq.Body)
  if err != nil {
     wx.Logger.Errorln(err)
     return nil,err
  }
  if err := json.Unmarshal(b,userdetail);err != nil{
     wx.Logger.Errorln(err)
     return nil,err
  }
  if userdetail.Errcode == 0 {
     wx.Logger.Debugln("Current user name: ",userdetail.Name)
     return userdetail,nil
  }else{
     wx.Logger.Errorln(userdetail.Errmsg)
     return nil,fmt.Errorf(userdetail.Errmsg)
  }
}


//批量获取企业微信 配置了客户联系功能的员工列表
func (wx *WX)GetFollowUsers()([]string,error){
  //https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_follow_user_list?access_token=ACCESS_TOKEN
  excontactusers := &ExContactUsers{}
  client := &http.Client{}

  token,err := wx.GetWXContactToken()
  if err != nil{
    wx.Logger.Errorln(err)
    return nil,err
  }

  uri,_ := url.Parse("https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_follow_user_list?access_token="+token)
  req,err := http.NewRequest("GET", uri.String(), nil)
  if err != nil{
     wx.Logger.Errorln(err)
     return nil,err
  }
  wx.Logger.Debugln("GetFollowUsers: ",uri.String())

  resq,err := client.Do(req)
  defer resq.Body.Close()

  b,err := ioutil.ReadAll(resq.Body)
  if err != nil {
     wx.Logger.Errorln(err)
     return nil,err
  }
  if err := json.Unmarshal(b,excontactusers);err != nil{
     wx.Logger.Errorln(err)
     return nil,err
  }
  wx.Logger.Debugln("excontactusers: ",excontactusers)
  if excontactusers.Errcode == 0{
     wx.Logger.Debugln("Follow_user: ",excontactusers.Follow_user)
     return excontactusers.Follow_user,nil
  }else{
     wx.Logger.Errorln(excontactusers.Errmsg)
     return nil,fmt.Errorf(excontactusers.Errmsg)
  }
}

//+++获取企业微信联系人 token
func  (wx *WX)GetWXContactToken()(string,error){
  tnow := time.Now().Unix()
  if tnow - wx.TokenSetTime < -10 && wx.ContactToken != ""{
      return wx.ContactToken,nil
  }    

  t1,err := wx.Redis.GetKeyTTL(WXContactTokenName)
  wx.Logger.Debugln("Get token :",WXContactTokenName,"ttl: ",t1)
  if err != nil || t1 == -2 {
    wxrt := &WXRT{}
    client := &http.Client{}

    // https://qyapi.weixin.qq.com/cgi-bin/gettoken

    uri,_ := url.Parse("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=" + wx.CorpID +"&corpsecret=" +wx.ContactSecret)
    req,err := http.NewRequest("GET", uri.String(), nil)
    if err != nil{
      wx.Logger.Errorln(err)
      return "",err
    }
    wx.Logger.Debugln("GET WX Token: ",req.URL.String())
  
    resq,err := client.Do(req)
    defer resq.Body.Close()
  
    b,err := ioutil.ReadAll(resq.Body)
    if err != nil {
      wx.Logger.Errorln(err)
      return "",err
    }
    if err := json.Unmarshal(b,wxrt);err != nil{
      wx.Logger.Errorln(err)
      return "",err
    }
    if wxrt.Errcode == 0 {
      if err := wx.Redis.SetKey(WXContactTokenName,strconv.Itoa(wxrt.Expires_in),wxrt.Access_token);err != nil{
         wx.Logger.Errorln(err)
      }

      tnow := time.Now().Unix()
      wx.ContactToken = wxrt.Access_token
      wx.TokenSetTime = tnow + int64(wxrt.Expires_in)
      return wxrt.Access_token,nil
    }else{
      wx.Logger.Errorln(wxrt.Errmsg)
      return "",fmt.Errorf(wxrt.Errmsg)
   }
  }else{
    token,err := wx.Redis.GetKey(WXContactTokenName)
    if err != nil{
       wx.Logger.Errorln(err)
       return "",err
    }
    return token,nil
  }
}

//获取当前用户 所属的外部客户id列表
func  (wx *WX)GetExUserList(userid string)([]string,error){
  // https://qyapi.weixin.qq.com/cgi-bin/externalcontact/list?access_token=ACCESS_TOKEN&userid=USERID
  external_list := &External_list{}
  client := &http.Client{}

  token,err := wx.GetWXToken()
  if err != nil{
    wx.Logger.Errorln(err)
    return nil,err
  }

  uri,_ := url.Parse("https://qyapi.weixin.qq.com")
  uri.Path = "/cgi-bin/externalcontact/list"
  req,err := http.NewRequest("GET", uri.String(), nil)
  if err != nil{
     wx.Logger.Errorln(err)
     return nil,err
  }
  q := req.URL.Query()
  q.Add("access_token",token)
  q.Add("userid",userid)
  req.URL.RawQuery = q.Encode()
  wx.Logger.Debugln(req.URL.String())

  resq,err := client.Do(req)
  defer resq.Body.Close()

  b,err := ioutil.ReadAll(resq.Body)
  if err != nil {
     wx.Logger.Errorln(err)
     return nil,err
  }
  if err := json.Unmarshal(b,external_list);err != nil{
     wx.Logger.Errorln(err)
     return nil,err
  }
  if external_list.Errcode == 0 {
     return external_list.External_userid,nil
  }else{
     wx.Logger.Errorln(external_list.Errmsg)
     return nil,fmt.Errorf(external_list.Errmsg)
  }
}

//获取外部客户详细信息 单个客户
func  (wx *WX)GetExUserInfo(exuserid string)(*ExUserInfo,error){
  // https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get?access_token=ACCESS_TOKEN&external_userid=EXTERNAL_USERID&cursor=CURSOR
  exuserinfo:= &ExUserInfo{}
  client := &http.Client{}

  token,err := wx.GetWXToken()
  if err != nil{
    wx.Logger.Errorln(err)
    return nil,err
  }

  uri,_ := url.Parse("https://qyapi.weixin.qq.com")
  uri.Path = "/cgi-bin/externalcontact/get"
  req,err := http.NewRequest("GET", uri.String(), nil)
  if err != nil{
     wx.Logger.Errorln(err)
     return nil,err
  }
  q := req.URL.Query()
  q.Add("access_token",token)
  q.Add("external_userid",exuserid)
  req.URL.RawQuery = q.Encode()
  // wx.Logger.Debugln(req.URL.String())

  resq,err := client.Do(req)
  defer resq.Body.Close()

  b,err := ioutil.ReadAll(resq.Body)
  if err != nil {
     wx.Logger.Errorln(err)
     return nil,err
  }
  if err := json.Unmarshal(b,exuserinfo);err != nil{
     wx.Logger.Errorln(err)
     return nil,err
  }
  if exuserinfo.Errcode == 0 {
     return exuserinfo,nil
  }else{
     wx.Logger.Errorln(exuserinfo.Errmsg)
     return nil,fmt.Errorf(exuserinfo.Errmsg)
  }
}



//修改客户联系人描述 备注 公司信息
//POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/remark?access_token=ACCESS_TOKEN
func(wx *WX)UpdateExUserRemark(data io.Reader)error{
  var result UpdateExUserRemarkResult
  client := &http.Client{}
  token,err := wx.GetWXToken()
  if err != nil{
    wx.Logger.Errorln(err)
    return err
  }

  uri,_ := url.Parse("https://qyapi.weixin.qq.com")
  uri.Path = "/cgi-bin/externalcontact/remark"
  req,err := http.NewRequest("POST", uri.String(), data)
  if err != nil{
     wx.Logger.Errorln(err)
     return err
  }
  q := req.URL.Query()
  q.Add("access_token",token)
  req.URL.RawQuery = q.Encode()
  resq,err := client.Do(req)
  defer resq.Body.Close()

  b,err := ioutil.ReadAll(resq.Body)
  if err != nil {
     wx.Logger.Errorln(err)
     return err
  }
  if err := json.Unmarshal(b,&result);err != nil{
     wx.Logger.Errorln(err)
     return err
  }
  if result.Errcode == 0 {
     return nil
  }else{
     wx.Logger.Errorln(result.Errmsg)
     return  fmt.Errorf(result.Errmsg)
  }
}



//在职分配 客户
//POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/transfer_customer?access_token=ACCESS_TOKEN
func (wx *WX)TransferCustomer(data io.Reader)(*TransferCustomerResult,error){
  var result TransferCustomerResult
  client := &http.Client{}
  token,err := wx.GetWXToken()
  if err != nil{
    wx.Logger.Errorln(err)
    return &result,err
  }

  uri,_ := url.Parse("https://qyapi.weixin.qq.com")
  uri.Path = "/cgi-bin/externalcontact/transfer_customer"
  req,err := http.NewRequest("POST", uri.String(), data)
  if err != nil{
    wx.Logger.Errorln(err)
     return &result,err
  }
  q := req.URL.Query()
  q.Add("access_token",token)
  req.URL.RawQuery = q.Encode()
  resq,err := client.Do(req)
  defer resq.Body.Close()

  b,err := ioutil.ReadAll(resq.Body)
  if err != nil {
     wx.Logger.Errorln(err)
     return &result,err
  }
  if err := json.Unmarshal(b,&result);err != nil{
     wx.Logger.Errorln(err)
     return &result,err
  }
  if result.Errcode == 0 {
     return &result,nil
  }else{
     wx.Logger.Errorln(result.Errmsg)
     return  nil,fmt.Errorf(result.Errmsg)
  }
}



//查询客户 在职分配的接替情况
//POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/transfer_result?access_token=ACCESS_TOKEN
func (wx *WX)TransferResult(data io.Reader)(*TransferResultLookResult,error){
  var result TransferResultLookResult
  client := &http.Client{}
  token,err := wx.GetWXToken()
  if err != nil{
    wx.Logger.Errorln(err)
    return nil,err
  }

  uri,_ := url.Parse("https://qyapi.weixin.qq.com")
  uri.Path = "/cgi-bin/externalcontact/transfer_result"
  req,err := http.NewRequest("POST", uri.String(), data)
  if err != nil{
     wx.Logger.Errorln(err)
     return nil,err
  }
  q := req.URL.Query()
  q.Add("access_token",token)
  req.URL.RawQuery = q.Encode()
  resq,err := client.Do(req)
  defer resq.Body.Close()

  b,err := ioutil.ReadAll(resq.Body)
  if err != nil {
     wx.Logger.Errorln(err)
     return nil,err
  }
  if err := json.Unmarshal(b,&result);err != nil{
     wx.Logger.Errorln(err)
     return nil,err
  }
  if result.Errcode == 0 {
     return &result,nil
  }else{
     wx.Logger.Errorln(result.Errmsg)
     return  nil,fmt.Errorf(result.Errmsg)
  }
}



//获取待分配的离职成员列表
//POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_unassigned_list?access_token=ACCESS_TOKEN
func (wx *WX)GetUnassignedList(data io.Reader)(*UnassignedList,error){
  var result UnassignedList
  client := &http.Client{}
  token,err := wx.GetWXToken()
  if err != nil{
    wx.Logger.Errorln(err)
    return nil,err
  }

  uri,_ := url.Parse("https://qyapi.weixin.qq.com")
  uri.Path = "/cgi-bin/externalcontact/get_unassigned_list"
  req,err := http.NewRequest("POST", uri.String(), data)
  if err != nil{
     wx.Logger.Errorln(err)
     return nil,err
  }
  q := req.URL.Query()
  q.Add("access_token",token)
  req.URL.RawQuery = q.Encode()
  resq,err := client.Do(req)
  defer resq.Body.Close()

  b,err := ioutil.ReadAll(resq.Body)
  if err != nil {
     wx.Logger.Errorln(err)
     return nil,err
  }
  if err := json.Unmarshal(b,&result);err != nil{
     wx.Logger.Errorln(err)
     return nil,err
  }
  if result.Errcode == 0 {
     wx.Logger.Errorln(err)
     return &result,nil
  }else{
     wx.Logger.Errorln(result.Errmsg)
     return  nil,fmt.Errorf(result.Errmsg)
  }
}



//分配离职成员的客户
//POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/resigned/transfer_customer?access_token=ACCESS_TOKEN
func (wx *WX)LiZiTransferCustomer(data io.Reader)(*LiZiTransferCustomerResult,error){
  var result LiZiTransferCustomerResult
  client := &http.Client{}
  token,err := wx.GetWXToken()
  if err != nil{
    wx.Logger.Errorln(err)
    return nil,err
  }

  uri,_ := url.Parse("https://qyapi.weixin.qq.com")
  uri.Path = "/cgi-bin/externalcontact/resigned/transfer_customer"
  req,err := http.NewRequest("POST", uri.String(), data)
  if err != nil{
    wx.Logger.Errorln(err)
     return nil,err
  }
  q := req.URL.Query()
  q.Add("access_token",token)
  req.URL.RawQuery = q.Encode()
  resq,err := client.Do(req)
  defer resq.Body.Close()

  b,err := ioutil.ReadAll(resq.Body)
  if err != nil {
    wx.Logger.Errorln(err)
     return nil,err
  }
  if err := json.Unmarshal(b,&result);err != nil{
    wx.Logger.Errorln(err)
     return nil,err
  }
  if result.Errcode == 0 {
     return &result,nil
  }else{
    wx.Logger.Errorln(result.Errmsg)
     return  nil,fmt.Errorf(result.Errmsg)
  }
}

//查询 离职继承的客户接替状态
//POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/resigned/transfer_result?access_token=ACCESS_TOKEN
func (wx *WX)LiZiTransferResult(data io.Reader)(*TransferResultLookResult,error){
  var result TransferResultLookResult
  client := &http.Client{}
  token,err := wx.GetWXToken()
  if err != nil{
    wx.Logger.Errorln(err)
    return nil,err
  }

  uri,_ := url.Parse("https://qyapi.weixin.qq.com")
  uri.Path = "/cgi-bin/externalcontact/resigned/transfer_result"
  req,err := http.NewRequest("POST", uri.String(), data)
  if err != nil{
    wx.Logger.Errorln(err)
     return nil,err
  }
  q := req.URL.Query()
  q.Add("access_token",token)
  req.URL.RawQuery = q.Encode()
  resq,err := client.Do(req)
  defer resq.Body.Close()

  b,err := ioutil.ReadAll(resq.Body)
  if err != nil {
    wx.Logger.Errorln(err)
     return nil,err
  }
  if err := json.Unmarshal(b,&result);err != nil{
    wx.Logger.Errorln(err)
     return nil,err
  }
  if result.Errcode == 0 {
     return &result,nil
  }else{
    wx.Logger.Errorln(result.Errmsg)
     return  nil,fmt.Errorf(result.Errmsg)
  }
}



//客户群 管理

//分配离职成员的客户群
//POSt https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/transfer?access_token=ACCESS_TOKEN
func (wx *WX)LiZiGroupChatTransfer(data io.Reader)(*LiZiGroupChatTransferResult,error){
  var result LiZiGroupChatTransferResult
  client := &http.Client{}
  token,err := wx.GetWXToken()
  if err != nil{
    wx.Logger.Errorln(err)
    return nil,err
  }

  uri,_ := url.Parse("https://qyapi.weixin.qq.com")
  uri.Path = "/cgi-bin/externalcontact/groupchat/transfer"
  req,err := http.NewRequest("POST", uri.String(), data)
  if err != nil{
    wx.Logger.Errorln(err)
     return nil,err
  }
  q := req.URL.Query()
  q.Add("access_token",token)
  req.URL.RawQuery = q.Encode()
  resq,err := client.Do(req)
  defer resq.Body.Close()

  b,err := ioutil.ReadAll(resq.Body)
  if err != nil {
    wx.Logger.Errorln(err)
     return nil,err
  }
  if err := json.Unmarshal(b,&result);err != nil{
    wx.Logger.Errorln(err)
     return nil,err
  }
  if result.Errcode == 0 {
     return &result,nil
  }else{
     wx.Logger.Errorln(result.Errmsg)
     return  nil,fmt.Errorf(result.Errmsg)
  }
}


  
//获取客户群列表
//POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/list?access_token=ACCESS_TOKEN
func (wx *WX)GroupChatList(data io.Reader)(*GroupChatListResult,error){
  var result GroupChatListResult
  client := &http.Client{}
  token,err := wx.GetWXToken()
  if err != nil{
    wx.Logger.Errorln(err)
    return nil,err
  }

  uri,_ := url.Parse("https://qyapi.weixin.qq.com")
  uri.Path = "/cgi-bin/externalcontact/groupchat/list"
  req,err := http.NewRequest("POST", uri.String(), data)
  if err != nil{
    wx.Logger.Errorln(err)
     return nil,err
  }
  q := req.URL.Query()
  q.Add("access_token",token)
  req.URL.RawQuery = q.Encode()
  resq,err := client.Do(req)
  defer resq.Body.Close()

  b,err := ioutil.ReadAll(resq.Body)
  if err != nil {
    wx.Logger.Errorln(err)
     return nil,err
  }
  if err := json.Unmarshal(b,&result);err != nil{
    wx.Logger.Errorln(err)
     return nil,err
  }
  if result.Errcode == 0 {
     return &result,nil
  }else{
    wx.Logger.Errorln(result.Errmsg)
     return  nil,fmt.Errorf(result.Errmsg)
  }
}


//获取客户群详情
//POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/get?access_token=ACCESS_TOKEN
func (wx *WX)GropuChatDetail(data io.Reader)(*GroupChatDetailResult,error){
  var result GroupChatDetailResult
  client := &http.Client{}
  token,err := wx.GetWXToken()
  if err != nil{
    wx.Logger.Errorln(err)
    return nil,err
  }

  uri,_ := url.Parse("https://qyapi.weixin.qq.com")
  uri.Path = "/cgi-bin/externalcontact/groupchat/get"
  req,err := http.NewRequest("POST", uri.String(), data)
  if err != nil{
    wx.Logger.Errorln(err)
     return nil,err
  }
  q := req.URL.Query()
  q.Add("access_token",token)
  req.URL.RawQuery = q.Encode()
  resq,err := client.Do(req)
  defer resq.Body.Close()

  b,err := ioutil.ReadAll(resq.Body)
  if err != nil {
    wx.Logger.Errorln(err)
     return nil,err
  }
  if err := json.Unmarshal(b,&result);err != nil{
    wx.Logger.Errorln(err)
     return nil,err
  }
  if result.Errcode == 0 {
     return &result,nil
  }else{
    wx.Logger.Errorln(result.Errmsg)
     return  nil,fmt.Errorf(result.Errmsg)
  }
}



//统计管理

//获取成员联系客户的数据，包括发起申请数、新增客户数、聊天数、发送消息数和删除/拉黑成员的客户数等指标 https://work.weixin.qq.com/api/doc/90000/90135/92132
//POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_user_behavior_data?access_token=ACCESS_TOKEN
func (wx *WX)UserBehaviorDataGet(data io.Reader)(*UserBehaviorDataResult,error){
  var result UserBehaviorDataResult
  client := &http.Client{}
  token,err := wx.GetWXToken()
  if err != nil{
    wx.Logger.Errorln(err)
    return nil,err
  }

  uri,_ := url.Parse("https://qyapi.weixin.qq.com")
  uri.Path = "/cgi-bin/externalcontact/get_user_behavior_data"
  req,err := http.NewRequest("POST", uri.String(), data)
  if err != nil{
    wx.Logger.Errorln(err)
     return nil,err
  }
  q := req.URL.Query()
  q.Add("access_token",token)
  req.URL.RawQuery = q.Encode()
  resq,err := client.Do(req)
  defer resq.Body.Close()

  b,err := ioutil.ReadAll(resq.Body)
  if err != nil {
    wx.Logger.Errorln(err)
     return nil,err
  }
  if err := json.Unmarshal(b,&result);err != nil{
    wx.Logger.Errorln(err)
     return nil,err
  }
  if result.Errcode == 0 {
     return &result,nil
  }else{
    wx.Logger.Errorln(result.Errmsg)
     return  nil,fmt.Errorf(result.Errmsg)
  }
}


//群聊数据统计

//按群主聚合的方式
//POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/statistic?access_token=ACCESS_TOKEN
func (wx *WX)GroupChatStatisticByPerson(data io.Reader)(*GroupChatStatisticByPersonResult,error){
  var result GroupChatStatisticByPersonResult
  client := &http.Client{}
  token,err := wx.GetWXToken()
  if err != nil{
    wx.Logger.Errorln(err)
    return nil,err
  }

  uri,_ := url.Parse("https://qyapi.weixin.qq.com")
  uri.Path = "/cgi-bin/externalcontact/groupchat/statistic"
  req,err := http.NewRequest("POST", uri.String(), data)
  if err != nil{
    wx.Logger.Errorln(err)
     return nil,err
  }
  q := req.URL.Query()
  q.Add("access_token",token)
  req.URL.RawQuery = q.Encode()
  resq,err := client.Do(req)
  defer resq.Body.Close()

  b,err := ioutil.ReadAll(resq.Body)
  if err != nil {
    wx.Logger.Errorln(err)
     return nil,err
  }
  if err := json.Unmarshal(b,&result);err != nil{
    wx.Logger.Errorln(err)
     return nil,err
  }
  if result.Errcode == 0 {
     return &result,nil
  }else{
    wx.Logger.Errorln(result.Errmsg)
     return  nil,fmt.Errorf(result.Errmsg)
  }
}

  
//按自然日聚合的方式
//POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/statistic_group_by_day?access_token=ACCESS_TOKEN
func (wx *WX)GroupChatStatisticByDay(data io.Reader)(*GroupChatStatisticByDayResult,error){
  var result GroupChatStatisticByDayResult
  client := &http.Client{}
  token,err := wx.GetWXToken()
  if err != nil{
    wx.Logger.Errorln(err)
    return nil,err
  }

  uri,_ := url.Parse("https://qyapi.weixin.qq.com")
  uri.Path = "/cgi-bin/externalcontact/groupchat/statistic_group_by_day"
  req,err := http.NewRequest("POST", uri.String(), data)
  if err != nil{
    wx.Logger.Errorln(err)
     return nil,err
  }
  q := req.URL.Query()
  q.Add("access_token",token)
  req.URL.RawQuery = q.Encode()
  resq,err := client.Do(req)
  defer resq.Body.Close()

  b,err := ioutil.ReadAll(resq.Body)
  if err != nil {
    wx.Logger.Errorln(err)
     return nil,err
  }
  if err := json.Unmarshal(b,&result);err != nil{
    wx.Logger.Errorln(err)
     return nil,err
  }
  if result.Errcode == 0 {
     return &result,nil
  }else{
    wx.Logger.Errorln(result.Errmsg)
     return  nil,fmt.Errorf(result.Errmsg)
  }
}



//jsapi ticket 获取
//GET https://qyapi.weixin.qq.com/cgi-bin/get_jsapi_ticket?access_token=ACCESS_TOKEN
func (wx *WX)GetJSTicket()(string,error){
  tnow := time.Now().Unix()
  if  wx.TicketSetTime - tnow > 5 && wx.JSTicket != ""{
      wx.Logger.Debugln("WX JSTicket set time is not timeout and return this token value.")
      return wx.JSTicket,nil
  }    

  t1,err := wx.Redis.GetKeyTTL(JSTicketName)
  wx.Logger.Debugln("Get JSTicket :",JSTicketName,"ttl: ",t1)
  if err != nil || t1 == -2{
      ticket := &JSAPITicket{}
      client := &http.Client{}
      token,err := wx.GetWXToken()
      if err != nil{
         wx.Logger.Errorln(err)
         return "",err
      }
   uri,_ := url.Parse("https://qyapi.weixin.qq.com")
   uri.Path = "/cgi-bin/get_jsapi_ticket"
   req,err := http.NewRequest("GET", uri.String(), nil)
   if err != nil{
     wx.Logger.Errorln(err)
     return "",err
  }
  q := req.URL.Query()
  q.Add("access_token",token)
  req.URL.RawQuery = q.Encode()

  resq,err := client.Do(req)
  defer resq.Body.Close()

  b,err := ioutil.ReadAll(resq.Body)
  if err != nil {
     wx.Logger.Errorln(err)
     return "",err
  }
  if err := json.Unmarshal(b,ticket);err != nil{
     wx.Logger.Errorln(err)
     return "",err
  }
  if ticket.Errcode == 0 {
      if err := wx.Redis.SetKey(JSTicketName,strconv.Itoa(ticket.ExpiresIn),ticket.Ticket);err != nil{
         wx.Logger.Errorln(err)
      }
      tnow := time.Now().Unix()
      wx.TicketSetTime = tnow + int64(ticket.ExpiresIn)
      wx.JSTicket = ticket.Ticket
      return ticket.Ticket,nil
   }else{
      wx.Logger.Errorln(ticket.Errmsg)
      return "",fmt.Errorf(ticket.Errmsg)
   }
  }else{
    ticket,err := wx.Redis.GetKey(JSTicketName)
    if err != nil{
      wx.Logger.Errorln(err)
       return "",err
    }
    return ticket,nil
 }
}


//获取部门信息
//GET https://qyapi.weixin.qq.com/cgi-bin/department/list?access_token=ACCESS_TOKEN&id=ID
func  (wx *WX)GetDepartmentInfo(departmentid string)(*DepartmentInfo,error){
  departmentinfo:= &DepartmentInfo{}
  client := &http.Client{}

  token,err := wx.GetWXToken()
  if err != nil{
    wx.Logger.Errorln(err)
    return nil,err
  }

  uri,_ := url.Parse("https://qyapi.weixin.qq.com")
  uri.Path = "/cgi-bin/department/list"
  req,err := http.NewRequest("GET", uri.String(), nil)
  if err != nil{
     wx.Logger.Errorln(err)
     return nil,err
  }
  q := req.URL.Query()
  q.Add("access_token",token)
  q.Add("id",departmentid)
  req.URL.RawQuery = q.Encode()
  wx.Logger.Debugln(req.URL.String())

  resq,err := client.Do(req)
  defer resq.Body.Close()

  b,err := ioutil.ReadAll(resq.Body)
  if err != nil {
     wx.Logger.Errorln(err)
     return nil,err
  }
  if err := json.Unmarshal(b,departmentinfo);err != nil{
     wx.Logger.Errorln(err)
     return nil,err
  }
  if departmentinfo.Errcode == 0 {
     return departmentinfo,nil
  }else{
     wx.Logger.Errorln(departmentinfo.Errmsg)
     return nil,fmt.Errorf(departmentinfo.Errmsg)
  }
}


//add user
//GET https://qyapi.weixin.qq.com/cgi-bin/user/simplelist?access_token=ACCESS_TOKEN&department_id=DEPARTMENT_ID&fetch_child=FETCH_CHILD
func (wx *WX)GetUserlist()(*WXUsers,error){
  users := &WXUsers{}
  client := &http.Client{}

  token,err := wx.GetWXToken()
  if err != nil{
    wx.Logger.Errorln(err)
    return nil,err
  }

  uri,_ := url.Parse("https://qyapi.weixin.qq.com")
  uri.Path = "/cgi-bin/user/simplelist"
  req,err := http.NewRequest("GET", uri.String(), nil)
  if err != nil{
     wx.Logger.Errorln(err)
     return nil,err
  }
  q := req.URL.Query()
  q.Add("access_token",token)
  q.Add("department_id","1")
   q.Add("fetch_child","1")
  req.URL.RawQuery = q.Encode()
  wx.Logger.Debugln(req.URL.String())

  resq,err := client.Do(req)
  defer resq.Body.Close()

  b,err := ioutil.ReadAll(resq.Body)
  if err != nil {
     wx.Logger.Errorln(err)
     return nil,err
  }
  if err := json.Unmarshal(b,users);err != nil{
     wx.Logger.Errorln(err)
     return nil,err
  }
  if users.Errcode == 0 {
     return users,nil
  }else{
     wx.Logger.Errorln(users.Errmsg)
     return nil,fmt.Errorf(users.Errmsg)
  }
}


//上传临时文件
// POST https://qyapi.weixin.qq.com/cgi-bin/media/upload?access_token=ACCESS_TOKEN&type=TYPE
func (wx *WX)UploadFile(file *bytes.Buffer,filename string)(*UploadFileResult,error){
  var result UploadFileResult
  client := &http.Client{}
  token,err := wx.GetWXToken()
  if err != nil{
    wx.Logger.Errorln(err)
    return nil,err
  }

   // 实例化multipart
   body := &bytes.Buffer{}                                            
   writer := multipart.NewWriter(body)       
   
   // 创建multipart 文件字段
   if filename ==""{
    filename = common.GenUid()+".xlsx"
   }
   part, err := writer.CreateFormFile("file",filename ) 
   if err != nil {
      return nil, err
   }
   // 写入文件数据到multipart
   _, err = io.Copy(part, file) 
   //将额外参数也写入到multipart
   // _ = writer.WriteField("filelength", file.Len()) 

   err = writer.Close()
   if err != nil {
      return nil, err
   }


  uri,_ := url.Parse("https://qyapi.weixin.qq.com")
  uri.Path = "/cgi-bin/media/upload"
  req,err := http.NewRequest("POST", uri.String(), body)
  if err != nil{
     wx.Logger.Errorln(err)
     return nil,err
  }
  req.Header.Add("Content-Type", writer.FormDataContentType())
  q := req.URL.Query()
  q.Add("access_token",token)
  q.Add("type","file")
  req.URL.RawQuery = q.Encode()
  resq,err := client.Do(req)
  defer resq.Body.Close()

  b,err := ioutil.ReadAll(resq.Body)
  if err != nil {
     wx.Logger.Errorln(err)
     return nil,err
  }
  if err := json.Unmarshal(b,&result);err != nil{
     wx.Logger.Errorln(err)
     return nil,err
  }
  if result.Errcode == 0 {
     return &result,nil
  }else{
     wx.Logger.Errorln(result.Errmsg)
     return  nil,fmt.Errorf(result.Errmsg)
  }
}



//发送消息
//POST https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=ACCESS_TOKEN
func (wx *WX)SendFileMsg(mediaid,touser string)(*MsgResult,error){
   agentid,err := strconv.Atoi(wx.AgentId)
   if err != nil{
      wx.Logger.Errorln(err)
   }
   filemsg := MsgData{
   Touser: touser,  
   Msgtype: "file", 
   Agentid: agentid,   
   File:  File{MediaID: mediaid},
   }
  var result MsgResult
  client := &http.Client{}
  token,err := wx.GetWXToken()
  if err != nil{
    wx.Logger.Errorln(err)
    return nil,err
  }

  data,err := json.Marshal(filemsg)
  if err != nil{
   wx.Logger.Errorln(err)
   return nil,err 
  }


  uri,_ := url.Parse("https://qyapi.weixin.qq.com")
  uri.Path = "/cgi-bin/message/send"
  req,err := http.NewRequest("POST", uri.String(), bytes.NewReader(data))
  if err != nil{
     wx.Logger.Errorln(err)
     return nil,err
  }
  q := req.URL.Query()
  q.Add("access_token",token)
  req.URL.RawQuery = q.Encode()
  resq,err := client.Do(req)
  defer resq.Body.Close()

  b,err := ioutil.ReadAll(resq.Body)
  if err != nil {
     wx.Logger.Errorln(err)
     return nil,err
  }
  if err := json.Unmarshal(b,&result);err != nil{
     wx.Logger.Errorln(err)
     return nil,err
  }
  if result.Errcode == 0 {
     return &result,nil
  }else{
     wx.Logger.Errorln(result.Errmsg)
     return  nil,fmt.Errorf(result.Errmsg)
  }
}


//发送text消息
//POST https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=ACCESS_TOKEN
func (wx *WX)SendTextMsg(touser,msgs string)(*MsgResult,error){
   agentid,err := strconv.Atoi(wx.AgentId)
   if err != nil{
      wx.Logger.Errorln(err)
   }
   textmsg := MsgData{
   Touser: touser,  
   Msgtype: "text", 
   Agentid: agentid,   
   Text:  Text{Content: msgs},
   }
  var result MsgResult
  client := &http.Client{}
  token,err := wx.GetWXToken()
  if err != nil{
    wx.Logger.Errorln(err)
    return nil,err
  }

  data,err := json.Marshal(textmsg)
  if err != nil{
   wx.Logger.Errorln(err)
   return nil,err 
  }


  uri,_ := url.Parse("https://qyapi.weixin.qq.com")
  uri.Path = "/cgi-bin/message/send"
  req,err := http.NewRequest("POST", uri.String(), bytes.NewReader(data))
  if err != nil{
     wx.Logger.Errorln(err)
     return nil,err
  }
  q := req.URL.Query()
  q.Add("access_token",token)
  req.URL.RawQuery = q.Encode()
  resq,err := client.Do(req)
  defer resq.Body.Close()

  b,err := ioutil.ReadAll(resq.Body)
  if err != nil {
     wx.Logger.Errorln(err)
     return nil,err
  }
  if err := json.Unmarshal(b,&result);err != nil{
     wx.Logger.Errorln(err)
     return nil,err
  }
  if result.Errcode == 0 {
     return &result,nil
  }else{
     wx.Logger.Errorln(result.Errmsg)
     return  nil,fmt.Errorf(result.Errmsg)
  }
}

//发送textcard消息
//POST https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=ACCESS_TOKEN
func (wx *WX)SendTextCardMsg(touser,msgs,titile string)(*MsgResult,error){
   agentid,err := strconv.Atoi(wx.AgentId)
   if err != nil{
      wx.Logger.Errorln(err)
   }
   textmsg := MsgData{
     Touser: touser,  
     Msgtype: "textcard", 
     Agentid: agentid,   
     Textcard:  Textcard{    
       Title: titile +"通知",
       Description: msgs,
       URL: "https://crm-iti.dm-ai.com",
       Btntxt: "详情",
      }, 
   }
  var result MsgResult
  client := &http.Client{}
  token,err := wx.GetWXToken()
  if err != nil{
    wx.Logger.Errorln(err)
    return nil,err
  }

  data,err := json.Marshal(textmsg)
  if err != nil{
   wx.Logger.Errorln(err)
   return nil,err 
  }


  uri,_ := url.Parse("https://qyapi.weixin.qq.com")
  uri.Path = "/cgi-bin/message/send"
  req,err := http.NewRequest("POST", uri.String(), bytes.NewReader(data))
  if err != nil{
     wx.Logger.Errorln(err)
     return nil,err
  }
  q := req.URL.Query()
  q.Add("access_token",token)
  req.URL.RawQuery = q.Encode()
  resq,err := client.Do(req)
  defer resq.Body.Close()

  b,err := ioutil.ReadAll(resq.Body)
  if err != nil {
     wx.Logger.Errorln(err)
     return nil,err
  }
  if err := json.Unmarshal(b,&result);err != nil{
     wx.Logger.Errorln(err)
     return nil,err
  }
  if result.Errcode == 0 {
     return &result,nil
  }else{
     wx.Logger.Errorln(result.Errmsg)
     return  nil,fmt.Errorf(result.Errmsg)
  }
}




