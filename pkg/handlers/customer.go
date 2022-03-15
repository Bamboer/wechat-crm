package handlers
import(
  "fmt"
  "sort"
  "time"
  "net/http"
  "strconv"
  "sync"
  // "strings"
  "encoding/json"
  "wxcrm/pkg/backend"
  "github.com/gin-gonic/gin"
  "wxcrm/pkg/common"
  "wxcrm/pkg/service"
)
//Here are customer,followrecords,customer user principal

var SeaCustomers = backend.Results{}
var mu sync.Mutex


func (h *HandlerVar)GetCustomer(c *gin.Context) {
	var user string
	action := c.Query("action")
	corpcode := c.Query("corpcode")

	pageindex := c.Query("pageindex")
	pagesize := c.Query("pagesize")
	if cookie, err := c.Cookie("wxcrm_username"); err == nil {
		user = cookie
	}else{
		user =  c.GetString("user")
		h.Logger.Debugln("GetCustomer get user from map:",user)
	}
	h.Logger.Debugln("Get customer Request: ",c.Request.URL.String())
	if user == "" {
		// h.Logger.Errorln("user id  is nil.---GetCustomer")
		c.JSON(http.StatusOK, gin.H{"status": 400,"ErrMsg": "userid id  is null"})
		c.AbortWithStatus(http.StatusOK)
	}

	if action == "info" && corpcode != "" {
		//返回与客户公司相关的信息
		customer:= h.DB.ViewCustomerInfo(corpcode)

		h.CustomerCheckChann <- corpcode
		// h.UserInfoSyncToRD <- user
		// h.UserInfoSyncToRD <- "manager"

		if customer.CorpName != "" {
			c.JSON(http.StatusOK,gin.H{"Customer": customer,"status": 200})
		}else{
			// h.Logger.Errorln("customer name is null ---GetCustomer")
			c.JSON(http.StatusOK, gin.H{"status": 400,"ErrMsg": "customer name is null"})
		}
		c.AbortWithStatus(http.StatusOK)
	}else if action=="list"{
		//返回客户列表
		dmaiuser := h.DB.ViewUser(user)
		var customers backend.Results
		scope := c.Query("scope")
		yesterday_increase := backend.Results{}

    tnow := time.Now()
    year := tnow.Year()
    month := tnow.Month()
    day  := tnow.Day()
		secondsEastOfUTC := int((8 * time.Hour).Seconds())
		beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)
    time1 := time.Date(year,month,day,0,0,0,0,beijing)
		yesterday := strconv.Itoa(int(time1.Unix()))
		Flag := false
		if  _,err := c.Cookie("manager_cookie");err == nil || dmaiuser.Manager == "true"{
			Flag = true
		}
		if  Flag && scope != "" && scope =="false"{
			h.Logger.Debugln("我要查看所有客户信息")
			yesterday_increase = h.DB.ViewYesterdayCustomers(yesterday)
			if redis_customers := h.Redis.CacheRe("scope");redis_customers != nil{
		      if v,ok := redis_customers.([]byte);ok{
		            json.Unmarshal(v,&customers)
		      }
		    h.Logger.Debugln("Big customers 从redis中拿数据！")
			}else{
		    customers = h.DB.ViewCustomers()
		    for k,v := range customers{
		     	tmp := h.DB.ViewCustomerUserprincipal(v.CorpCode)
		     	records := h.DB.ViewCustomerFollowRecords(v.CorpCode)
		     	if len(tmp) >0 && tmp[len(tmp)-1].Model.UpdatedAt.Year() > 1973{
		        	customers[k].Genjin = tmp[len(tmp)-1].MerchandiserName
		     	}
		     	if len(records)>0 && records[len(records)-1].Model.UpdatedAt.Year() > 1973{
		        	customers[k].RecordTime = records[len(records)-1].Model.UpdatedAt
		        }
		     	h.Logger.Debugln("Genjin: ",customers[k].Genjin)
		    }
		    sort.Sort(customers)
				if sort.IsSorted(customers){
				    h.Logger.Debugln("customers sorted....")
				}else{
				    h.Logger.Debugln("customers not sorted....")
				}
        if err := h.Redis.CacheSet("scope",customers);err != nil{
            h.Logger.Errorln(err)
        }
			}
		}else if Flag {
			h.Logger.Debugln("我的manager 能获取所有客户信息")
			yesterday_increase = h.DB.ViewYesterdayCustomers(yesterday)
			if redis_customers := h.Redis.CacheRe("manager");redis_customers != nil{
		      if v,ok := redis_customers.([]byte);ok{
		            json.Unmarshal(v,&customers)
		      }
		      h.Logger.Debugln("Manager 从redis中拿数据！")
			}else{
	         tmpcustomers := h.DB.ViewCustomers()
	         for _,v := range tmpcustomers{
	         	tmp := h.DB.ViewCustomerUserprincipal(v.CorpCode)
	         	records := h.DB.ViewCustomerFollowRecords(v.CorpCode)
	         	if len(records)>0 && records[len(records)-1].Model.UpdatedAt.Year() > 1973{
	            	v.RecordTime = records[len(records)-1].Model.UpdatedAt
	           }
	         	if len(tmp) >0 && tmp[len(tmp)-1].MerchandiserName !=""{
	            v.Genjin = tmp[len(tmp)-1].MerchandiserName
	            for _,n  := range tmp{
	            	v.Genjins = append(v.Genjins,n.MerchandiserName)
	            }
	            customers = append(customers,v)
	         	}
	         	// h.Logger.Debugln("Genjin: ",customers[k].Genjin)
	         }

					sort.Sort(customers)
					if sort.IsSorted(customers){
					      h.Logger.Debugln("customers sorted....")
					}else{
					    h.Logger.Debugln("customers not sorted....")
					}
          if err := h.Redis.CacheSet("manager",customers);err != nil{
            h.Logger.Errorln(err)
          }
		 }
		}else{
			yesterday_increase = h.DB.ViewUserYesterdayCustomers(user,yesterday)
			if redis_customers := h.Redis.CacheRe(user);redis_customers != nil{
		      if v,ok := redis_customers.([]byte);ok{
		            json.Unmarshal(v,&customers)
		      }
		      h.Logger.Debugln("user 从redis中拿数据！")
		  }else{
			  customers = h.DB.ViewUserCustomers(user)
			  h.Logger.Debugln("我是普通用户，只能看到我自己的客户")
	      for k,v := range customers{
	      	tmp := h.DB.ViewCustomerContact(v.CorpCode)
	      	records := h.DB.ViewCustomerFollowRecords(v.CorpCode)

	      	if len(tmp) >0 && tmp[0].Model.UpdatedAt.Year() > 1973{
	        	customers[k].Contact = tmp[len(tmp)-1].Name
	      	}
	      	if len(records)>0 && records[0].Model.UpdatedAt.Year() > 1973{
	        	customers[k].RecordTime = records[len(records)-1].Model.UpdatedAt
	        }
	      	h.Logger.Debugln("Contact: ",customers[k].Contact)
	      }
				sort.Sort(customers)
			  if sort.IsSorted(customers){
			      h.Logger.Debugln("customers sorted....")
			  }else{
			      h.Logger.Debugln("customers not sorted....")
			  }
				h.Redis.DelKey(user)
				h.UserInfoSyncToRD <- user
		  }
		}

		cuslen := len(customers)
		index,err := strconv.Atoi(pageindex)
		if err != nil{
			index = 1 
		}
		size,err := strconv.Atoi(pagesize)
		if err != nil {
			size = cuslen 
		}
		increasernum := 0
		if len(yesterday_increase) == 0{
			increasernum = 0
		}else{
			increasernum = len(yesterday_increase)
		}
		if index > 0 && size >0{
			if  index ==1 && size > cuslen{
				c.JSON(http.StatusOK,gin.H{"customers":customers,"status":200,"pagesize":size,"pageindex":index,"total":cuslen,"yesterday_increase": increasernum})
			}else if index*size <= cuslen{
				c.JSON(http.StatusOK,gin.H{"customers": customers[(index-1)*size:index*size],"status": 200,"pagesize":size,"pageindex":index,"total":cuslen,"yesterday_increase": increasernum})
				c.AbortWithStatus(http.StatusOK)
			}else if index*size > cuslen && cuslen > (index-1)*size{
				c.JSON(http.StatusOK,gin.H{"customers":customers[(index-1)*size:cuslen],"status":200,"pagesize":size,"pageindex":index,"total":cuslen,"yesterday_increase": increasernum})
			}else{
				c.JSON(http.StatusOK,gin.H{"customers":[]string{},"status":200,"pagesize":size,"pageindex":index,"total":cuslen})
			}
		}else{
		h.Logger.Errorln("参数错误！")
		c.JSON(http.StatusOK,gin.H{"status":400,"ErrMsg":"参数错误！"})
	}}else{
		c.JSON(http.StatusOK,gin.H{"status":400,"ErrMsg":"参数错误!"})
	}
}

type AddCustomer struct {
	Customer  backend.Customer
	Contactor backend.Contactor
}

func (h *HandlerVar)PostCustomer(c *gin.Context) {
	var user string
	action := c.Query("action")
	pageindex := c.Query("pageindex")
	pagesize := c.Query("pagesize")
	if cookie, err := c.Cookie("wxcrm_username"); err == nil {
		user  = cookie
	}else{
		user =  c.GetString("user")
	}

	if user == "" {
		h.Logger.Errorln("cookie and context map key: user is nil.")

		c.JSON(http.StatusOK, gin.H{"status": 500,"ErrMsg": "userid获取失败!"})
		c.AbortWithStatus(http.StatusOK)
	}
	if action == "update" {
		var recv backend.Customer
		c.ShouldBindJSON(&recv)
		if recv.CorpCode == ""{
			c.JSON(http.StatusOK,gin.H{"status":400,"ErrMsg":"corpcode为空"})
			return
		}
		h.Logger.Debugln("Post Update  Customer recv: ", recv)
		if err := h.DB.UpdateCustomer(&recv); err == nil {

			h.Redis.DelKey("manager")
			h.UserInfoSyncToRD <- "manager"
			h.Redis.DelKey(user)
			h.UserInfoSyncToRD <- user
			h.Redis.DelKey("scope")
			h.UserInfoSyncToRD <- "scope"
			h.CustomerSyncChann <-  recv.CorpCode


			udpate_customer_info := h.DB.ViewCustomerInfo(recv.CorpCode)
			h.AddOperation(user,recv.CorpCode,"更新","客户",udpate_customer_info.CorpName)
			c.JSON(http.StatusOK, gin.H{"status": 200})
		} else {
			h.Logger.Errorln(err)
			c.JSON(http.StatusOK, gin.H{"status": 400,"ErrMsg": err})
		}
		c.AbortWithStatus(http.StatusOK)
	}
  if action=="search"{
		//返回客户列表
		dmaiuser := h.DB.ViewUser(user)
		var searchdata map[string]string
		var customers backend.Results
		c.ShouldBindJSON(&searchdata)
		Flag := false 

		if _,err := c.Cookie("manager_cookie");err == nil || dmaiuser.Manager == "true" {
			Flag = true
		}

		if Flag {
			h.Logger.Debugln("我的manager 能获取所有客户信息")
			if redis_customers := h.Redis.CacheRe("manager");redis_customers != nil{
		      if v,ok := redis_customers.([]byte);ok{
		            json.Unmarshal(v,&customers)
		      }
		      h.Logger.Debugln("Manager 从redis中拿数据！")
		  }else{
	         tmpcustomers := h.DB.ViewCustomers()
	         for _,v := range tmpcustomers{
	         	tmp := h.DB.ViewCustomerUserprincipal(v.CorpCode)
	         	records := h.DB.ViewCustomerFollowRecords(v.CorpCode)
	         	if len(records)>0 && records[0].Model.UpdatedAt.Year() > 1973 {
	            	v.RecordTime = records[len(records)-1].Model.UpdatedAt
	           }
	         	if len(tmp) >0 && tmp[0].Model.UpdatedAt.Year() > 1973{
	            	v.Genjin = tmp[len(tmp)-1].MerchandiserName
	            	for _,j := range tmp{
	            		v.Genjins = append(v.Genjins,j.MerchandiserName) 
	            	}	
	            	customers = append(customers,v)
	         	}
	         	// h.Logger.Debugln("Genjin: ",customers[k].Genjin)
	         }
				  sort.Sort(customers)
			    if sort.IsSorted(customers){
			      h.Logger.Debugln("customers sorted....")
			    }else{
			      h.Logger.Debugln("customers not sorted....")
			    }
        if err := h.Redis.CacheSet("manager",customers);err != nil{
            h.Logger.Errorln(err)
        }
		  }
		}else{
			if redis_customers := h.Redis.CacheRe(user);redis_customers != nil{
		      if v,ok := redis_customers.([]byte);ok{
		            json.Unmarshal(v,&customers)
		      }
		      h.Logger.Debugln("user 从redis中拿数据！")
		  }else{
			  customers = h.DB.ViewUserCustomers(user)
			  h.Logger.Debugln("我是普通用户，只能看到我自己的客户")
	      for k,v := range customers{
	      	genjintmp := h.DB.ViewCustomerUserprincipal(v.CorpCode)
	      	tmp := h.DB.ViewCustomerContact(v.CorpCode)
	      	records := h.DB.ViewCustomerFollowRecords(v.CorpCode)

	      	if len(tmp) >0 &&tmp[len(tmp)-1].Model.UpdatedAt.Year() > 1973{
	      	  customers[k].Contact = tmp[len(tmp)-1].Name
	      	}
	      	for _,j := range genjintmp{
	      		customers[k].Genjins = append(customers[k].Genjins,j.MerchandiserName)
	      	}
	      	if len(records)>0 && records[len(records)-1].Model.UpdatedAt.Year() > 1973{
	        	customers[k].RecordTime = records[len(records)-1].Model.UpdatedAt
	        }
	      	h.Logger.Debugln("Contact: ",customers[k].Contact)
	      }
				sort.Sort(customers)
			  if sort.IsSorted(customers){
			      h.Logger.Debugln("customers sorted....")
			  }else{
			      h.Logger.Debugln("customers not sorted....")
			  }
				h.Redis.DelKey(user)
				h.UserInfoSyncToRD <- user
		  }
		}
		if customers != nil {
			// tmp := backend.Results{}
			// // for k,v := range searchdata{

			// // }
			tmp := h.BasisSearch(customers,searchdata)
			// for _,v := range customers{
			// 	h.Logger.Debugln("string index: ",strings.Index(v.CorpName,keyword),"corpname: ",v.CorpName," key: ",keyword)
			// 	if strings.Index(v.CorpName,keyword) != -1{
			// 		tmp = append(tmp,v)
			// 		h.Logger.Debugln("Match: ",v)
			// 	}
			// }
			cuslen := len(tmp)		
			index,err := strconv.Atoi(pageindex)
			if err != nil{
				index = 1 
			}
			size,err := strconv.Atoi(pagesize)
			if err != nil{
				size = cuslen 
			}
			if index > 0 && size>0{
				if index == 1 && size > cuslen{
					c.JSON(http.StatusOK,gin.H{"customers":tmp,"status":200,"pagesize":size,"pageindex":index,"total":cuslen})
				}else if index*size <=cuslen{
					c.JSON(http.StatusOK,gin.H{"customers": tmp[(index-1)*size:index*size],"status": 200,"pagesize":size,"pageindex":index,"total":cuslen})
				}else if index*size >cuslen && cuslen > (index-1)*size{
					c.JSON(http.StatusOK,gin.H{"customers":tmp[(index-1)*size:cuslen],"status":200,"pagesize":size,"pageindex":index,"total":cuslen})
				}else{
					c.JSON(http.StatusOK,gin.H{"customers":[]string{},"status":200,"pagesize":size,"pageindex":index,"total":cuslen})
				}
			}else{
				h.Logger.Errorln("参数错误！")
				c.JSON(http.StatusOK,gin.H{"status":400,"ErrMsg":"参数错误！"})
			}
		}
		// c.AbortWithStatus(http.StatusOK)
		return
	}
	if action == "add" {
		var recv AddCustomer
		var cup backend.CustomerUserprincipal
		var cop backend.ContactUserprincipal

		// dmaiuser := h.DB.ViewUser(user)

		c.ShouldBindJSON(&recv)
		tmpresult := fmt.Sprintf("%+v",recv)
		h.Logger.Debugln("Post Add Customer recv: ", tmpresult)


		if recv.Contactor.Wechatid !="" {
    	wechatContactor := h.DB.ViewWechatContactor(recv.Contactor.Wechatid)
    	if len(wechatContactor) >= 3{
    		c.JSON(http.StatusOK,gin.H{"status":400,"ErrMsg":"企业微信联系人已添加3次！"})
    		c.AbortWithStatus(http.StatusOK)
    		return
    	}
		} 
		if recv.Customer.CorpName == ""{
    		c.JSON(http.StatusOK,gin.H{"status":400,"ErrMsg":"客户名称为空！"})
    		c.AbortWithStatus(http.StatusOK)
    		return
		}


		tmpcustomer := h.DB.CheckCustomer(recv.Customer.CorpName)
		if tmpcustomer.CorpName != ""{
			h.Logger.Infoln(tmpcustomer.CorpName," 已存在！")
			recv.Contactor.CorpCode = tmpcustomer.CorpCode
	    if recv.Contactor.Wechatid !=""{
		    	exuserinfo,err := h.WX.GetExUserInfo(recv.Contactor.Wechatid)
		    	recv.Contactor.Logo = exuserinfo.ExternalContact.Avatar
		    	recv.Contactor.ContactCode = "H" + common.GenUid()
		    	recv.Contactor.Iswx = "1"
		    	if err != nil{
		    		h.Logger.Errorln(err)
		    	}
		  }  	
			if err := h.DB.AddContactor(&recv.Contactor); err != nil {
			    h.Logger.Errorln(err)
			}
			h.AddOperation(user,tmpcustomer.CorpCode,"增加","联系人",recv.Contactor.Name)
			
			h.Redis.DelKey("manager")
			h.UserInfoSyncToRD <- "manager"
			h.Redis.DelKey(user)
			h.UserInfoSyncToRD <- user
			h.Redis.DelKey("scope")
			h.UserInfoSyncToRD <- "scope"
        //添加相对应的 联系人 销售对应表
		    cop.ContactCode = recv.Contactor.ContactCode
		    cop.Ownerid = user
		    cop.Createid = user
		    if err := h.DB.AddContactUserprincipal(&cop);err !=nil{
		    	h.Logger.Errorln(err)
		    }

	    //添加相对应的销售 客户对应表
	    if  ! h.DB.CheckCUP(user,tmpcustomer.CorpCode){
		    dmaiuser := h.DB.ViewUser(user)
		    cup.CorpCode = tmpcustomer.CorpCode
		    cup.Merchandiserid = user
		    cup.MerchandiserName = dmaiuser.Name 
		    cup.Createby = user 
		    if err := h.DB.AddCustomerUserprincipal(&cup);err !=nil{
		        h.Logger.Errorln(err)
		    }
		    h.AddOperation(user,tmpcustomer.CorpCode,"增加","跟进人",dmaiuser.Name)
	    }
		  h.CustomerSyncChann <-  tmpcustomer.CorpCode
		  c.JSON(200, gin.H{"status": 200})
			c.AbortWithStatus(http.StatusOK)
			return
		}

		recv.Customer.CorpCode = "W" + common.GenUid()
		//这里需要加企查查数据 或另外存放一份企查查数据
		companyInfo,err := h.QCC.QGetBasicDetailsByName(recv.Customer.CorpName)
		if err == nil &&companyInfo.Status =="200"{
			recv.Customer.Opername = companyInfo.Result.Opername
			recv.Customer.Startdate = companyInfo.Result.Startdate
			recv.Customer.Province = companyInfo.Result.Province
			recv.Customer.Address = companyInfo.Result.Address
			recv.Customer.Status = companyInfo.Result.Status 
			recv.Customer.Registcapi = companyInfo.Result.Registcapi
			recv.Customer.Creditcode = companyInfo.Result.Creditcode
			recv.Customer.Scope = companyInfo.Result.Scope
			recv.Customer.Econkind = companyInfo.Result.Econkind
			recv.Customer.Belongorg = companyInfo.Result.Belongorg 
			recv.Customer.Orgno = companyInfo.Result.Orgno 
			recv.Customer.Logo = companyInfo.Result.Imageurl 
		}else{
			h.Logger.Errorln(err)
			c.JSON(http.StatusOK,gin.H{"status":400,"ErrMsg":"客户不存在工商信息，添加失败！"})
			return
		}
    recv.Customer.Region = common.RegionCheck(recv.Customer.Province)
    recv.Customer.Trade = common.TradeCheck(recv.Customer.CorpName,recv.Customer.Scope)
    if recv.Customer.Logo == ""{
    	if err := common.DrawLogo(recv.Customer.CorpName,h.Opts.SrcDir,h.Opts.SrcDir+"/"+recv.Customer.CorpCode+".png");err == nil{
    		recv.Customer.Logo = "https://"+ h.Opts.ServerName +"/api/src/"+ recv.Customer.CorpCode+".png"
    	}
    }

		if err := h.DB.AddCustomer(&recv.Customer); err != nil {
			h.Logger.Errorln(err)
		} 
		h.AddOperation(user,recv.Customer.CorpCode,"增加","客户",recv.Customer.CorpName)

		//添加相对应的客户联系人
		recv.Contactor.CorpCode = recv.Customer.CorpCode
		recv.Contactor.ContactCode = "H" + common.GenUid()
        if recv.Contactor.Wechatid !=""{
	    	recv.Contactor.Iswx = "1"
	    	exuserinfo,err := h.WX.GetExUserInfo(recv.Contactor.Wechatid)
	        if err != nil{
	            h.Logger.Errorln(err)
	        }
	    	// recv.Contactor.Name = exuserinfo.ExternalContact.Name
	    	// recv.Contactor.Gender = strconv.Itoa(exuserinfo.ExternalContact.Gender)
	    	recv.Contactor.Logo = exuserinfo.ExternalContact.Avatar
	    	if recv.Contactor.Postion == ""{
		    	recv.Contactor.Postion  = exuserinfo.ExternalContact.Position
	    	}
      }
    
	if err := h.DB.AddContactor(&recv.Contactor); err != nil {
	    h.Logger.Errorln(err)
	}
	h.AddOperation(user,recv.Customer.CorpCode,"增加","联系人",recv.Contactor.Name)
		// contact := h.DB.ViewWechatContactor(recv.Contactor.Contactid)
    //添加相对应的销售 客户对应表
    dmaiuser := h.DB.ViewUser(user)
    cup.CorpCode = recv.Customer.CorpCode
    cup.Merchandiserid = user
    cup.MerchandiserName = dmaiuser.Name 
    cup.Createby = user 
    cup.Presaler = "true"
    if err := h.DB.AddCustomerUserprincipal(&cup);err !=nil{
        h.Logger.Errorln(err)
    }
    h.AddOperation(user,recv.Customer.CorpCode,"增加","跟进人",dmaiuser.Name)

		h.Redis.DelKey("manager")
		h.UserInfoSyncToRD <- "manager"
		h.Redis.DelKey(user)
		h.UserInfoSyncToRD <- user
		h.Redis.DelKey("scope")
		h.UserInfoSyncToRD <- "scope"
     //添加相对应的 联系人 销售对应表
     cop.ContactCode= recv.Contactor.ContactCode
     cop.Ownerid = user
     cop.Createid = user
    if err := h.DB.AddContactUserprincipal(&cop);err !=nil{
    	h.Logger.Errorln(err)
    }
		//同步客户到用友
		h.CustomerSyncChann <- recv.Customer.CorpCode
		h.CustomerCheckChann <- recv.Customer.CorpCode
		c.JSON(200, gin.H{"status": 200})
		c.AbortWithStatus(http.StatusOK)
	}
}



//跟进记录
func (h *HandlerVar)FollowRecords(c *gin.Context) {
	var data backend.FollowRecord
	action := c.Query("action")
	id := c.Query("id")
	var user string
	if cookie, err := c.Cookie("wxcrm_username"); err == nil {
		user  = cookie
	}else{
		user =  c.GetString("user")
	}

	if action == "update" &&id != "" {
		c.ShouldBindJSON(&data) 
		data.Updateby = user 
		dmaiuser := h.DB.ViewUser(user)
		data.UpdatebyName = dmaiuser.Name
		tid,_ := strconv.Atoi(id)
                data.Model.ID = uint(tid)
		err := h.DB.UpdateFollowRecord(&data)
		if err == nil {

			h.Redis.DelKey("manager")
			h.UserInfoSyncToRD <- "manager"
			h.Redis.DelKey(user)
			h.UserInfoSyncToRD <- user
			
			h.AddOperation(user,data.CorpCode,"更新","跟进记录",data.Name)
			c.JSON(http.StatusOK, gin.H{"status": 200})
		} else {
			h.Logger.Errorln(err)
			c.JSON(http.StatusOK, gin.H{"status": 400,"ErrMsg": err})
		}
		c.AbortWithStatus(http.StatusOK)
	}
	if action == "add" {
		dmaiuser := h.DB.ViewUser(user)
		c.ShouldBindJSON(&data)
		if data.CorpCode == ""{
			h.Logger.Errorln("Add Follow Record get CorpCode is nil.")
			c.JSON(http.StatusOK,gin.H{"status":400,"ErrMsg":"corpcode为空"})
			c.AbortWithStatus(http.StatusNotAcceptable)
			return
		}
		data.Createby = user
		data.CreatebyName = dmaiuser.Name
		data.Updateby = user
		data.UpdatebyName = dmaiuser.Name
		h.Logger.Debugln("Follow Record add data: ",data)
		err := h.DB.AddFollowRecord(&data)
		if err == nil {
			h.Redis.DelKey("manager")
			h.UserInfoSyncToRD <- "manager"
			h.Redis.DelKey(user)
			h.UserInfoSyncToRD <- user

			h.AddOperation(user,data.CorpCode,"增加","跟进记录",data.Name)
			c.JSON(http.StatusOK, gin.H{"status": 200})
		} else {
			h.Logger.Errorln(err)
			c.JSON(http.StatusOK, gin.H{"status": 400,"ErrMsg": err})
		}
		c.AbortWithStatus(http.StatusOK)
	}
	if action == "delete" && id != "" {
		fl := h.DB.ViewSingleFollowRecord(id) 
		err := h.DB.DelFollowRecord(id)
		if err == nil {
			h.AddOperation(user,fl.CorpCode,"删除","跟进记录",fl.Name)
			c.JSON(http.StatusOK, gin.H{"status": 200})
			return
		} else {
			h.Logger.Errorln(err)
			c.JSON(http.StatusOK, gin.H{"status": 400,"ErrMsg": err})
		}
		c.AbortWithStatus(http.StatusOK)
	}
	if action == "view"{
		records := []backend.FollowRecord{} 
		corpcode := c.Query("corpcode")
		projectname := c.Query("project")
		if corpcode == ""{
			c.JSON(http.StatusOK, gin.H{"status": 400,"ErrMsg": "CorpCode为空!"})
			// c.AbortWithStatus(http.StatusNotAcceptable)
			return
		}
		if projectname != ""{
			records = h.DB.ViewProjectsFollowRecords(corpcode,projectname)
		}else{
		    records = h.DB.ViewCustomerFollowRecords(corpcode)
		}

		// c.JSON(http.StatusOK, gin.H{"FollowerThroughs": followthroughs})
		c.JSON(http.StatusOK, gin.H{"Records": records,"status": 200})
		c.AbortWithStatus(http.StatusOK)
	}
	if action == "single" {
		corpcode := c.Query("corpcode")
		if corpcode == ""{
			c.JSON(http.StatusOK, gin.H{"status": 400,"ErrMsg": "CorpCode为空!"})
			// c.AbortWithStatus(http.StatusNotAcceptable)
			return
		}
		record := h.DB.ViewSingleCustomerFollowRecord(corpcode)
		// c.JSON(http.StatusOK, gin.H{"FollowerThroughs": followthroughs})
		if record.Name ==""{
			c.JSON(http.StatusOK, gin.H{"Record": "null","status": 200})
		}else{
		c.JSON(http.StatusOK, gin.H{"Record": record,"status": 200})	
		}

		c.AbortWithStatus(http.StatusOK)
	}
}

type CUP struct{
	Presaler  string  
	Saler     string  
	CorpCode  string  
	Deletes    []string
}
type CUPs []CUP

type  Assign map[string] []string

func (h *HandlerVar)CustomerUserPrincipal(c *gin.Context) { 
	var users CUPs  
	c.ShouldBindJSON(&users)
	action := c.Query("action")
	cookie, err := c.Cookie("wxcrm_username")
	h.Logger.Debugln("STart principal.....")
  if err != nil{
  	cookie = c.GetString("user")
  }

  if action == "view"{
  	corpcode := c.Query("corpcode")

		if corpcode == ""{
			c.JSON(http.StatusNotAcceptable,gin.H{"status": 400,"ErrMsg": "CorpCode为空!"})
			c.AbortWithStatus(http.StatusOK)
			return
		}
		cups := h.DB.ViewCustomerUserprincipal(corpcode)
		h.Logger.Debugln("view现有跟进人： ",cups)
		c.JSON(http.StatusOK,gin.H{"status":200,"CustomerPrincipalUsers": cups})
		c.AbortWithStatus(http.StatusOK)
	}
	if action == "add"{
		fairluretmp := Assign{}
		addtmp := Assign{}
		h.Logger.Debugln("user principals:",users)
		if len(users) >0{
			for _,v := range users{
				// judge := h.checkin(cookie,v.CorpCode)
				// cur_dmaiuser := h.DB.ViewUser(cookie)
			 //  if cur_dmaiuser.Manager =="false"{
				//   if !judge {
				// 	  h.Logger.Infoln(cookie," have not customer: ",v.CorpCode," privileges...")
				//   	continue
				//   }
			 //  }

				if v.Presaler != ""{
					if tmpcup := h.DB.ViewSingleCP(v.Presaler,v.CorpCode);tmpcup.MerchandiserName != ""{
				  	continue
					}

					customerinfotmp := h.DB.ViewCustomerInfo(v.CorpCode)
					cup := backend.CustomerUserprincipal{}
					cup.CorpCode = v.CorpCode
					cup.Createby = cookie
					predmaiuser := h.DB.ViewUser(v.Presaler)
					cup.Merchandiserid = v.Presaler
					cup.MerchandiserName = predmaiuser.Name
					cup.Presaler = "true"
					cup.Saler = "false"
					cup.Logo = predmaiuser.Logo
					if err := h.DB.AddCustomerUserprincipal(&cup);err != nil{
						h.Logger.Errorln(err)
						fairluretmp[v.Presaler] = append(fairluretmp[v.Presaler],customerinfotmp.CorpName)
					}else{
						addtmp[v.Presaler] = append(addtmp[v.Presaler],customerinfotmp.CorpName)
						h.AddOperation(cookie,v.CorpCode,"添加","跟进人",v.Presaler)
					}
				}
				if v.Saler !=""{
					if tmpcup := h.DB.ViewSingleCP(v.Saler,v.CorpCode);tmpcup.MerchandiserName != ""{
				  	continue
					}

					customerinfotmp := h.DB.ViewCustomerInfo(v.CorpCode)
					cup := backend.CustomerUserprincipal{}
					cup.CorpCode = v.CorpCode
					cup.Createby = cookie
					salerdmaiuser := h.DB.ViewUser(v.Saler)
					cup.Merchandiserid = v.Saler
					cup.MerchandiserName = salerdmaiuser.Name
					cup.Saler = "true"
					cup.Presaler = "false"
					cup.Logo = salerdmaiuser.Logo
					if err := h.DB.AddCustomerUserprincipal(&cup);err != nil{
						h.Logger.Errorln(err)
						fairluretmp[v.Saler] = append(fairluretmp[v.Saler],customerinfotmp.CorpName)
					}else{
						addtmp[v.Saler] = append(addtmp[v.Saler],customerinfotmp.CorpName)
						h.AddOperation(cookie,v.CorpCode,"添加","跟进人",v.Saler)
					}
				}
        h.CustomerSyncChann <- v.CorpCode
        h.Redis.DelKey("manager")
		h.UserInfoSyncToRD <- "manager"
		h.Redis.DelKey(cookie)
		h.UserInfoSyncToRD <- cookie
		h.Redis.DelKey("scope")
		h.UserInfoSyncToRD <- "scope"

		for j,h := range SeaCustomers{
			if v.CorpCode == h.CorpCode{
				mu.Lock()
				SeaCustomers = append(SeaCustomers[:j],SeaCustomers[j+1:]...)
	      mu.Unlock()
				break
			} 
		}

	}
			h.GenJinNotification(addtmp)
		}else{
			h.Logger.Debugln("没有处理！")
		}
		c.JSON(http.StatusOK,gin.H{"status":200,"customer_failures":fairluretmp})
		return
	}
	if action == "delete"{
		fairluretmp := Assign{}
		if len(users) >0{
			for _,v := range users{
				for _,j := range v.Deletes{
					judge := h.checkin(cookie,v.CorpCode)
					cur_dmaiuser := h.DB.ViewUser(cookie)
					customerinfotmp := h.DB.ViewCustomerInfo(v.CorpCode)
				  if cur_dmaiuser.Manager =="false"{
				  	if  !judge {
					  	h.Logger.Infoln(cookie," have not customer: ",v.CorpCode," privileges...")
					  	continue
				  	}
				  }
				  if cur_dmaiuser.Manager == "true"{
				  	if err := h.DB.DelCustomerUserprincipal(j,v.CorpCode);err != nil{
				  		h.Logger.Errorln(err)
				  		fairluretmp[j] = append(fairluretmp[j],customerinfotmp.CorpName)
				  		continue
				  	}else{
					  	h.AddOperation(cookie,v.CorpCode,"删除","跟进人",j)
				  	}
				  }
				  singlecup := h.DB.ViewSingleCP(j,v.CorpCode)
				  if singlecup.MerchandiserName == ""{
				  	continue
				  }
				  if singlecup.Presaler == "true"{
				  	h.Logger.Debugln("不能删除售前！")
				  	fairluretmp[j] = append(fairluretmp[j],customerinfotmp.CorpName)
				  	continue
				  }
				  if cookie == singlecup.Merchandiserid{
				  	h.Logger.Debugln("不能删除自身！")
				  	fairluretmp[j] = append(fairluretmp[j],customerinfotmp.CorpName)
				  	continue
				  }
				  usercup := h.DB.ViewSingleCP(cookie,v.CorpCode)
				  if usercup.Model.CreatedAt.Unix() > singlecup.Model.CreatedAt.Unix(){
				  		h.Logger.Debugln("不能删除之前添加的跟进人！")
				  		fairluretmp[j] = append(fairluretmp[j],customerinfotmp.CorpName)
				  		continue
				  	}
			  	if err := h.DB.DelCustomerUserprincipal(j,v.CorpCode);err != nil{
			  		h.Logger.Errorln(err)
			  		fairluretmp[j] = append(fairluretmp[j],customerinfotmp.CorpName)
			  	}else{
				  	h.AddOperation(cookie,v.CorpCode,"删除","跟进人",j)
			  	}	
				}
				h.CustomerSyncChann <- v.CorpCode
        h.Redis.DelKey("manager")
			  h.UserInfoSyncToRD <- "manager"
			  h.Redis.DelKey(cookie)
				h.UserInfoSyncToRD <- cookie
        cups := h.DB.ViewCustomerUserprincipal(v.CorpCode)
				if len(cups)< 0 || cups[0].MerchandiserName == ""{
					customerinfo := h.DB.ViewCustomerInfo(v.CorpCode)
					if customerinfo.CorpName != ""{
						tmpcustomer := backend.Result{}
						tmpcustomer.CorpCode = customerinfo.CorpCode
						tmpcustomer.CorpName = customerinfo.CorpName
						tmpcustomer.Logo = customerinfo.Logo
						tmpcustomer.Address = customerinfo.Address
						tmpcustomer.Trade = customerinfo.Trade 
						tmpcustomer.Region = customerinfo.Region 
						tmpcustomer.UpdatedAt = customerinfo.Model.UpdatedAt 
						tmpcustomer.CreatedAt = customerinfo.Model.CreatedAt
						mu.Lock()
						SeaCustomers = append(backend.Results{tmpcustomer},SeaCustomers...)
						mu.Unlock()
					}
				}
			}
		}
		c.JSON(http.StatusOK,gin.H{"status":200,"customer_failures":fairluretmp})
		c.AbortWithStatus(http.StatusOK)
		return
	}

}

func (h *HandlerVar)SeaCustomer(c *gin.Context){
	// var SeaCustomers backend.Results
	var user string
	action := c.Query("action")
  cookie,err := c.Cookie("wxcrm_username")
  if err != nil{
  	user = c.GetString("user")
  }else{
  	user = cookie
  }
	if action == "sea"{
		corpcode := c.Query("corpcode")
		if corpcode == ""{
			c.JSON(http.StatusNotAcceptable,gin.H{"status": 400,"ErrMsg": "CorpCode为空!"})
			c.AbortWithStatus(http.StatusOK)
			return
		}

		if err := h.DB.TruncateCUP(corpcode);err != nil{
			c.JSON(http.StatusNotAcceptable,gin.H{"status": 400,"ErrMsg": err})
			return
		}else{
			h.Redis.DelKey(user)
			h.UserInfoSyncToRD <- user
			h.Redis.DelKey("manager")
			h.UserInfoSyncToRD <- "manager"
		  h.Redis.DelKey("scope")
			h.UserInfoSyncToRD <- "scope"
			customerinfo := h.DB.ViewCustomerInfo(corpcode)
			if customerinfo.CorpName != ""{
				tmpcustomer := backend.Result{}
				tmpcustomer.CorpCode = corpcode
				tmpcustomer.CorpName = customerinfo.CorpName
				tmpcustomer.Logo = customerinfo.Logo
				tmpcustomer.Address = customerinfo.Address
				tmpcustomer.Trade = customerinfo.Trade 
				tmpcustomer.Region = customerinfo.Region 
				tmpcustomer.UpdatedAt = customerinfo.Model.UpdatedAt 
				tmpcustomer.CreatedAt = customerinfo.Model.CreatedAt

	    	contactmp := h.DB.ViewCustomerContact(corpcode)
	    	records := h.DB.ViewCustomerFollowRecords(corpcode)
	    	if len(contactmp) >0 &&contactmp[len(contactmp)-1].Model.UpdatedAt.Year() > 1973{
	    	  tmpcustomer.Contact = contactmp[len(contactmp)-1].Name
	    	}
	    	if len(records)>0 && records[len(records)-1].Model.UpdatedAt.Year() > 1973{
	      	tmpcustomer.RecordTime = records[len(records)-1].Model.UpdatedAt
	      }

				mu.Lock()
				SeaCustomers = append(backend.Results{tmpcustomer},SeaCustomers...)
				mu.Unlock()
			}
			c.JSON(http.StatusOK,gin.H{"status": 200})
			return
		}
	}
	if action == "list"{
		pageindex := c.Query("pageindex")
		pagesize := c.Query("pagesize")
		// yesterday_increase := backend.Results{}
		// yesterday := strconv.Itoa(int(time.Now().Unix()-86400))
		// yesterday_increase = h.DB.ViewYesterdayCustomers(yesterday)
    mu.Lock()
		if len(SeaCustomers)==0{
			customers := h.DB.ViewCustomers()
			for _,v := range customers{
				tmp := h.DB.ViewCustomerUserprincipal(v.CorpCode)
				if len(tmp) >= 1 && tmp[len(tmp) - 1].MerchandiserName != ""{
					continue
				}
				SeaCustomers = append(SeaCustomers,v)
			}
		    sort.Sort(SeaCustomers)
			if sort.IsSorted(SeaCustomers){
			    h.Logger.Debugln("SeaCustomers sorted....")
			}else{
			    h.Logger.Debugln("SeaCustomers not sorted....")
			}
		}
    mu.Unlock()

		cuslen := len(SeaCustomers)	
		index,err := strconv.Atoi(pageindex)
		if err != nil{
			index = 1 
		}
		size,err := strconv.Atoi(pagesize)
		if err != nil {
			size = cuslen 
		}
		if index > 0 && size >0{
			if  index ==1 && size > cuslen{
				c.JSON(http.StatusOK,gin.H{"customers":SeaCustomers,"status":200,"pagesize":size,"pageindex":index,"total":cuslen})
				return
			}else if index*size <= cuslen{
				c.JSON(http.StatusOK,gin.H{"customers": SeaCustomers[(index-1)*size:index*size],"status": 200,"pagesize":size,"pageindex":index,"total":cuslen})
				c.AbortWithStatus(http.StatusOK)
				return
			}else if index*size > cuslen && cuslen > (index-1)*size{
				c.JSON(http.StatusOK,gin.H{"customers":SeaCustomers[(index-1)*size:cuslen],"status":200,"pagesize":size,"pageindex":index,"total":cuslen})
				return
			}else{
				c.JSON(http.StatusOK,gin.H{"customers":[]string{},"status":200,"pagesize":size,"pageindex":index,"total":cuslen})
				return
			}
    }

		c.JSON(http.StatusOK,gin.H{"status":200,"Seacustomers":SeaCustomers}) 
		return
	}

  if action == "search"{
		pageindex := c.Query("pageindex")
		pagesize := c.Query("pagesize")
		var searchdata map[string]string
		c.ShouldBindJSON(&searchdata)
		// yesterday_increase := backend.Results{}
		// yesterday := strconv.Itoa(int(time.Now().Unix()-86400))
		// yesterday_increase = h.DB.ViewYesterdayCustomers(yesterday)
    mu.Lock()
		if len(SeaCustomers)==0{
			customers := h.DB.ViewCustomers()
			for _,v := range customers{
				tmp := h.DB.ViewCustomerUserprincipal(v.CorpCode)
				if len(tmp) >= 1 && tmp[len(tmp) - 1].MerchandiserName != ""{
					continue
				}
      	contactmp := h.DB.ViewCustomerContact(v.CorpCode)
      	records := h.DB.ViewCustomerFollowRecords(v.CorpCode)
      	if len(contactmp) >0 &&contactmp[len(contactmp)-1].Model.UpdatedAt.Year() > 1973{
      	  v.Contact = contactmp[len(contactmp)-1].Name
      	}
      	if len(records)>0 && records[len(records)-1].Model.UpdatedAt.Year() > 1973{
        	v.RecordTime = records[len(records)-1].Model.UpdatedAt
        }
				SeaCustomers = append(SeaCustomers,v)
			}
			sort.Sort(SeaCustomers)
			if sort.IsSorted(SeaCustomers){
			      h.Logger.Debugln("SeaCustomers sorted....")
			}else{
			      h.Logger.Debugln("SeaCustomers not sorted....")
			}
		}
    mu.Unlock()
    if len(SeaCustomers) > 0{
        tmp := h.BasisSearch(SeaCustomers,searchdata)
				cuslen := len(tmp)	
				index,err := strconv.Atoi(pageindex)
				if err != nil{
					index = 1 
				}
				size,err := strconv.Atoi(pagesize)
				if err != nil {
					size = cuslen 
				}

				if index > 0 && size >0{
					if  index ==1 && size > cuslen{
						c.JSON(http.StatusOK,gin.H{"customers":tmp,"status":200,"pagesize":size,"pageindex":index,"total":cuslen})
						return
					}else if index*size <= cuslen{
						c.JSON(http.StatusOK,gin.H{"customers": tmp[(index-1)*size:index*size],"status": 200,"pagesize":size,"pageindex":index,"total":cuslen})
						return
					}else if index*size > cuslen && cuslen > (index-1)*size{
						c.JSON(http.StatusOK,gin.H{"customers":tmp[(index-1)*size:cuslen],"status":200,"pagesize":size,"pageindex":index,"total":cuslen})
						return
					}else{
						c.JSON(http.StatusOK,gin.H{"customers":[]string{},"status":200,"pagesize":size,"pageindex":index,"total":cuslen})
						return 
					}
		    }

		c.JSON(http.StatusOK,gin.H{"status":200,"Seacustomers":tmp}) 
		return
  }
  }
}


type Cuslevel struct{
	ID  int64     `json:"id"`
	Text string    `json:"text"`//名称
}
type Projects struct{
	ID  string    `json:"id"`
	Text  string    `json:"text"`
}
//获取组织单元
func (h *HandlerVar)YSInfo(c *gin.Context){
	typ := c.Query("type")
	if typ == "component"{
		tmp := h.YS.Componets()
		if tmp == nil{
			c.JSON(http.StatusOK,gin.H{"status":400,"ErrMsg":"components not found"})
			return
		}
		c.JSON(http.StatusOK,gin.H{"status":200,"Components":tmp})
		c.AbortWithStatus(http.StatusOK)
	}else if typ == "level"{
		data := &service.CusLevelData{Pageindex:0,Pagesize:1000}
		result := []Cuslevel{} 
		tmp,err := h.YS.GetCuslevel(data)
		for _,v := range tmp.Data.Recordlist{
			result = append(result,Cuslevel{ID: v.ID,Text: v.Name})
		}
		if err != nil{
			c.JSON(http.StatusOK,gin.H{"status":400,"ErrMsg": err,"Default": []Cuslevel{{Text:"核心客户",ID:123},{Text:"重要客户",ID:223},{Text:"一般客户",ID:333}}})
		}else{
		c.JSON(http.StatusOK,gin.H{"status":200,"CusLevels": result})
		}
		c.AbortWithStatus(http.StatusOK)
	}else if typ =="project"{
		index,err := strconv.Atoi(c.Query("pageindex"))
		if err != nil{
			// h.Logger.Errorln(err)
			index = 1
		}
		size,err := strconv.Atoi(c.Query("pagesize"))
		if err != nil{
			// h.Logger.Errorln(err)
			size = 100
		}
		if index >0 && size >0{
		data := &service.YSProjectListData{
			PageIndex:index,
			PageSize: size,
		    }
		if result,err := h.YS.GetProjectList(data);err != nil{
			c.JSON(http.StatusOK,gin.H{"status":400,"ErrMsg":err})
			c.AbortWithStatus(http.StatusOK)
		   }else{
		   	tmp := []Projects{}
		   	for _,v := range result.Data.RecordList{
		   		tmp = append(tmp,Projects{ID:v.ID,Text:v.Name})
		   	}
			c.JSON(http.StatusOK,gin.H{"status":200,"projects":tmp,"PageCount":result.Data.PageCount,"BeginPageIndex":result.Data.BeginPageIndex,"EndPageIndex": result.Data.EndPageIndex})
		   } 			
		}else{
		  data := &service.YSProjectListData{
			PageIndex:1,
			PageSize: 100,
		    }
		if result,err := h.YS.GetProjectList(data);err != nil{
			c.JSON(http.StatusOK,gin.H{"status":400,"ErrMsg":err})
			c.AbortWithStatus(http.StatusOK)
		   }else{
		   	tmp := []Projects{}
		   	for _,v := range result.Data.RecordList{
		   		tmp = append(tmp,Projects{ID:v.ID,Text:v.Name})
		   	}
			c.JSON(http.StatusOK,gin.H{"status":200,"projects":result.Data.RecordList,"PageCount":result.Data.PageCount,"BeginPageIndex":result.Data.BeginPageIndex,"EndPageIndex": result.Data.EndPageIndex})
		   }	
		}

	}else if typ=="sync"{
		corpcode := c.Query("corpcode")
		tmpcode := []string{corpcode}
		if err := h.DoSync(tmpcode);err !=nil{
			c.JSON(http.StatusOK,gin.H{"status":400,"ErrMsg":err})
		}else{
			c.JSON(http.StatusOK,gin.H{"status":200})
		}
	}else{
		c.JSON(http.StatusOK,gin.H{"status":400,"ErrMsg":"参数错误"})
	}
}

func (h *HandlerVar)checkin(user,corpcode string)bool{
	cups := h.DB.ViewCustomerUserprincipal(corpcode)
	for _,v := range cups{
		if user == v.Merchandiserid{
			return true
		}
	}
	return false
}
