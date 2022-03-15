package handlers
import (
	"fmt"
    "sort"
    "bytes"
   "time"
   "bufio"
   "os"
   "io"
   "io/ioutil"
   "context"
   "strings"
   "strconv"
   "net/http"
  "wxcrm/pkg/backend"
  // "wxcrm/pkg/service"
  "wxcrm/pkg/common"
  "wxcrm/pkg/service"
)





//添加操作记录
func (h *HandlerVar)AddOperation(fman,corpcode,operation,subject,object string)error{
	userinfo  := h.DB.ViewUser(fman)
	input := backend.Operation{Fman:fman,Fmanname:userinfo.Name,Corpcode:corpcode,Operation:operation,Subject:subject,Object:object}
	if err := h.DB.AddOperation(&input);err != nil{
		h.Logger.Errorln(err)
		return err 
	}
	return nil 
}

//接收通道信息
func(h *HandlerVar)OnSync(ctx context.Context){
	for{
		select{
		case <- ctx.Done():
			h.Logger.Errorln(ctx.Err())
		case corpcode := <- h.CustomerCheckChann:
            h.Logger.Debugln("get corpcode: ",corpcode)
			h.CustomerCheckInfo(corpcode)
	    case corpcode := <- h.CustomerSyncChann:
            h.Logger.Debugln("customer sync channel length: ",len(h.CustomerSyncChann))
	    	h.SyncToYS(corpcode) 
        case identify := <- h.UserInfoSyncToRD:
            h.Logger.Debugln("start cache to redis server.")
            h.SyncToRD(identify)    
		default:
		    time.Sleep(1 * time.Millisecond )
            tnow := time.Now().Unix() % 1628942400    
            if tnow == 0{
                h.CodeRecords = map[string]string{}
            }     	
		}
	}
}


//添加客户验证信息
func (h *HandlerVar)AddCustomerCheck(corpcode string){
	customerinfo := h.DB.ViewCustomerInfo(corpcode)
	keyword  := customerinfo.CorpName

	shixin,err := h.QCC.QShiXinChecker(keyword)
	if err != nil || shixin.Status != "200" || shixin.Result.VerifyResult == 0{
        if shixin.Result.VerifyResult == 0{
            h.Logger.Errorln(keyword,shixin)
        }else{
            h.Logger.Errorln(err)
        }
	}else{
		for _,v := range shixin.Result.Data{
            Customershixin := backend.CustomerShiXin{}
			Customershixin.CorpCode = corpcode
			Customershixin.LianDate = v.Liandate
			Customershixin.Anno = v.Anno
			Customershixin.Executegov = v.Executegov 
			Customershixin.Executestatus = v.Executestatus
			Customershixin.Publicdate = v.Publicdate
			Customershixin.Executeno = v.Executeno
			if err := h.DB.AddCustomerShixin(&Customershixin);err != nil{
				h.Logger.Errorln(err)
			}  
		}
	}
	serious,err := h.QCC.QSeriousIllegalChecker(keyword)
	if err != nil ||serious.Status != "200" || serious.Result.VerifyResult == 0{
        if serious.Result.VerifyResult == 0{
            h.Logger.Errorln(keyword,serious)
        }else{
            h.Logger.Errorln(err)
        }
	}else{
		for _,v := range serious.Result.Data{
            CustomerSerious := backend.CustomerSeriousIllegal{}
			CustomerSerious.CorpCode = corpcode
			CustomerSerious.Type = v.Type 
			CustomerSerious.AddReason = v.AddReason
			CustomerSerious.AddDate = v.AddDate
			CustomerSerious.AddOffice = v.AddOffice
			CustomerSerious.RemoveReason = v.RemoveReason
			CustomerSerious.RemoveDate = v.RemoveDate 
			CustomerSerious.RemoveOffice = v.RemoveOffice
			if err := h.DB.AddCustomerSeriousIllegal(&CustomerSerious);err != nil{
				h.Logger.Errorln(err)
			}
		}
	}
	tax,err := h.QCC.QTaxIllegalChecker(keyword)
	if err != nil ||tax.Status != "200" || tax.Result.VerifyResult == 0{
        if tax.Result.VerifyResult == 0{
            h.Logger.Errorln(keyword,tax)
        }else{
            h.Logger.Errorln(err)
        }
	}else{
		for _,v := range tax.Result.Data{
            CustomerTax := backend.CustomerTaxIllegal{}
			CustomerTax.CorpCode = corpcode
			CustomerTax.PublishDate = v.PublishDate
			CustomerTax.CaseNature = v.CaseNature
			CustomerTax.TaxGov = v.TaxGov 
			if err := h.DB.AddCustomerTaxIllegal(&CustomerTax);err != nil{
				h.Logger.Errorln(err)
			}
		}
	}
	adminpenalty,err := h.QCC.QAdminPenaltyChecker(keyword)
	if err != nil || adminpenalty.Status != "200" || adminpenalty.Result.VerifyResult == 0{
        if adminpenalty.Result.VerifyResult == 0{
            h.Logger.Errorln(keyword,adminpenalty)
        }else{
            h.Logger.Errorln(err)
        }
	}else{
		for _,v := range adminpenalty.Result.Data{
            CustomerAdmin := backend.CustomerAdminPenalty{}
			CustomerAdmin.CorpCode = corpcode 
			CustomerAdmin.DocNo = v.DocNo  
			CustomerAdmin.PunishReason = v.PunishReason
			CustomerAdmin.PunishResult = v.PunishResult
			CustomerAdmin.PunishOffice = v.PunishOffice
			CustomerAdmin.PunishDate = v.PunishDate
			CustomerAdmin.Source = v.Source
			CustomerAdmin.SourceCode = v.SourceCode
			if err := h.DB.AddCustomerAdminPenalty(&CustomerAdmin);err != nil{
				h.Logger.Errorln(err)
			}
		}
	}
}

func(h *HandlerVar)CustomerCheckInfo(corpcode string){
    h.Logger.Debugln("检查验证有效期: ",corpcode)
    shixins := h.DB.ViewCustomerShixin(corpcode)
    tnow := time.Now().Unix()
    if len(shixins) > 0{
        if tnow - shixins[len(shixins)-1].Model.UpdatedAt.Unix() > 86400*30{
            h.CustomerBasicInfo(corpcode)
            h.AddCustomerCheck(corpcode)
        }
        return
    }
    serious := h.DB.ViewCustomerSeriousIllegal(corpcode)
    if len(serious) > 0{
        if tnow - serious[len(serious)-1].Model.UpdatedAt.Unix() > 86400*30{
            h.CustomerBasicInfo(corpcode)
            h.AddCustomerCheck(corpcode)
        }
        return
    }
    taxs   := h.DB.ViewCustomerTaxIllegal(corpcode)
    if len(taxs) > 0{
        if tnow - taxs[len(taxs)-1].Model.UpdatedAt.Unix() > 86400*30{
            h.CustomerBasicInfo(corpcode)
            h.AddCustomerCheck(corpcode)
        }
        return
    }
    penalties := h.DB.ViewCustomerAdminPenalty(corpcode)
    if len(penalties) > 0{
        if tnow - penalties[len(penalties)-1].Model.UpdatedAt.Unix() > 86400*30{
            h.CustomerBasicInfo(corpcode)
            h.AddCustomerCheck(corpcode)
        }
        return
    }
    if timeout,_ := h.Redis.GetKeyTTL(corpcode);timeout == -1 {  //关闭 -1  需要开启 -2
        h.Logger.Debugln("两周休眠期：",corpcode)
        if err := h.Redis.SetKey(corpcode,"2592000","checked");err != nil{
            h.Logger.Errorln(err)
        }
        h.AddCustomerCheck(corpcode)
        h.CustomerBasicInfo(corpcode)
    }
    //意向分设定
    customerinfo := h.DB.ViewCustomerInfo(corpcode)
    // customerinfo.Region = h.RegionCheck(customerinfo.Province)
    // customerinfo.Trade = h.TradeCheck(customerinfo)
    tmpflogs := h.DB.ViewCustomerFollowRecords(corpcode)
    customerinfo.IntentionScore = strconv.Itoa(50+len(tmpflogs))
    if customerinfo.Customerlevel == "重要客户" {
        customerinfo.IntentionScore = strconv.Itoa(60+len(tmpflogs))
    }
    if customerinfo.Customerlevel == "核心客户" {
        customerinfo.IntentionScore = strconv.Itoa(65+len(tmpflogs))
    }
    customerinfo.QualificationScore = "50"
    if 50+len(tmpflogs) > 100{
        customerinfo.IntentionScore = "99.99"
    }
    //资质分

    money := 0                                                    
    if strings.Index(customerinfo.Registcapi,"万") != -1{
        smon := strings.Split(customerinfo.Registcapi,"万")[0]
        if strings.Index(smon,".") != -1{
            smon = strings.Split(smon,".")[0]
        }
        money,_ = strconv.Atoi(smon)
    }else if strings.Index(customerinfo.Registcapi,"亿") != -1 {
        smon := strings.Split(customerinfo.Registcapi,"亿")[0] 
        if strings.Index(smon,".") != -1{
            smon = strings.Split(smon,".")[0]
        }
        money,_ = strconv.Atoi(smon)
        money = money *1e8
    } 

    if money > 1000 && money < 10000 {
        customerinfo.QualificationScore = strconv.Itoa(50+10)
    }else if money>10000 || strings.Index(customerinfo.CorpName,"集团") != -1{
        customerinfo.QualificationScore = strconv.Itoa(50+12)
    }else{
        customerinfo.QualificationScore = strconv.Itoa(50+2)
    }
    //
    customerinfo.Region = common.RegionCheck(customerinfo.Province)
    customerinfo.Trade = common.TradeCheck(customerinfo.CorpName,customerinfo.Scope)
    if customerinfo.Logo == ""{
        if err := common.DrawLogo(customerinfo.CorpName,h.Opts.SrcDir,h.Opts.SrcDir+"/"+customerinfo.CorpCode+".png");err == nil{
            customerinfo.Logo = "https://"+ h.Opts.ServerName +"/api/src/"+ customerinfo.CorpCode+".png"
            h.UserInfoSyncToRD <- "manager"
            h.UserInfoSyncToRD <- "scope"
        }
    }

    if err := h.DB.UpdateCustomer(customerinfo);err != nil{
        h.Logger.Errorln(err)
    }

}


//手动同步客户到用友
func (h *HandlerVar)DoSync(corpcodes []string)error{
  if len(corpcodes) > 0{
    for _,v := range corpcodes{
        h.CustomerSyncChann <- v 
        // h.UpdateYSCustomerchan <- v  
    }
    return nil  
  }else{
    return fmt.Errorf("参数错误！")
  }
} 

//同步客户到ERP
func(h *HandlerVar)SyncToYS(corpcode string){
    if corpcode==""{
        h.Logger.Debugln("corpcode 为空")
        return
    }
    customertmp := h.DB.ViewCustomerInfo(corpcode)
    check :=  service.MerchantListData{PageIndex:1,PageSize:100,Name: customertmp.CorpName}
    merchanlistresult,err := h.YS.MerchantList(&check)
    if err != nil{
        h.Logger.Errorln(err)
    }

    h.Logger.Debugln("返回客户列表为：",merchanlistresult)
    h.Logger.Debugln("客户检查结果为：",customertmp)

    if customertmp.CorpName != "" && len(merchanlistresult.Data.RecordList)>0{
        h.Logger.Infoln(customertmp.CorpName,"在ERP跟CRM中已经存在!执行同步操作")
        h.UpdateToYS(corpcode)
        return
    }

    //判断ERP系统是否存在客户档案
    if customertmp.CorpName != "" && len(merchanlistresult.Data.RecordList)==0{
        h.Logger.Infoln(customertmp.CorpName ,"CRM中存在,但ERP中没有!执行ERP客户档案创建操作")
        h.NewToYS(corpcode)
        return
    }
    h.Logger.Errorln("ERP客户档案同步失败")
}


//新增时同步crm客户到ERP 
func(h *HandlerVar)NewToYS(corpcode string){
    //判断用友接口是否正常，失败时睡眠10min后再检查
    for {
        if _,err := h.YS.GetUnits();err != nil{
            h.Logger.Errorln("用友检查错误： ",err," 睡眠10分钟后同步: ",corpcode)
            time.Sleep(time.Minute * 10)
        }else{
            time.Sleep(time.Second * 5)
            break
        }
    }
    cup := h.DB.ViewCustomerUserprincipal(corpcode)
    if len(cup)<1{
        h.Logger.Errorln("客户没有跟进人！")
    }

    h.Logger.Debugln("start syncer...")
    gid := ""
    gname := ""
    for _,v := range h.YS.Componets(){
        h.Logger.Debugln("components: ",v,"gname: ",v.Text)
        if strings.Index(v.Text,"广州") != -1{
            gid = v.Orgid
            gname = v.Text
            break   
        }
    } 
    data := &service.MerchantData{}
    customer := h.DB.ViewCustomerInfo(corpcode)

    data.Data.Createorg =  gid 
    data.Data.Code = common.GenUid()
    data.Data.CreateorgName = gname
    // data.Data.Country = "0040be98-735b-44c0-afe5-54d11a96037b"
    // data.Data.CountryName = 
    data.Data.Customerclass = 2174258698604800
    data.Data.CustomerclassPath = "2174256368997120|2174258698604800|"
    // data.Data.CustomerclassCode = "0101"
    data.Data.Taxpayingcategories= 0
    data.Data.Enterprisenature = "0"
    data.Data.Internalorg = false
    data.Data.InternalorgidName = ""
    data.Data.CreateorgCode = "DMAI01"            
    data.Data.InvoicingcustomersName = customer.OwnerOrgname        //开票客户
    data.Data.Enterprisename = customer.CorpName         //企业名称
    data.Data.Orgname = customer.CorpName
    data.Data.Merchantsmanager.Countrycode ="86"
    data.Data.Merchantrole = service.Merchantrole{Businessrole: "1",Tobimmigrationmode: "0",Settlementmethod: "0",Cardtype: "0",Status: "Insert"}
    data.Data.Status = "Insert"
    data.Data.Name = service.Name{ZhCn: customer.CorpName}
    data.Data.Shortname = service.Shortname{ZhCn: customer.CorpName}
    data.Data.Belongorg = customer.OwnerOrgid
    data.Data.Leadername = customer.Opername          //法人代表
    //  data.Data.Leadernameidno =                        //法人身份证
    data.Data.Creditcode = customer.Creditcode        //信用代码
    data.Data.Address.ZhCn = customer.Address              //地址
    customerlevel := customer.CustomerlevelId
    h.Logger.Debugln("customerlevel: ",customerlevel)
    data.Data.Merchantapplieddetail =  service.Merchantapplieddetail{Signback:false,
                                                        Payway:99,
                                                        Searchcode:"",
                                                        Customerlevel: customerlevel,
                                                        CustomerlevelCode: "xvsf",
                                                        CustomerlevelName: customer.Customerlevel,
                                                        Status: "Insert",
                                                        }

    tmprincipals := []service.Principal{}                                                      
    for _,v := range cup{
        pname,pid,dpname,dpid,err := h.YS.CheckERPUserInfo(strings.Split(v.MerchandiserName,"(")[0])
        if err != nil{
            h.Logger.Errorln("ERP查询失败: ",err)
            continue
        }
        h.Logger.Debugln("Pid: ",pid," pname: ",pname,"dpname: ",dpname," dpid: ",dpid)
        // if len(tmprincipals) == 0{
        //     tmprincipals = append(tmprincipals,service.Principal{SpecialManagementDepCode:"false",IsDefault:true,HasDefaultInit:true,ProfessSalesmanName:pname,ProfessSalesman:pid,SpecialManagementDep:dpid,SpecialManagementDepName:dpname,Status:"Insert"})
        // }else{
            tmprincipals = append(tmprincipals,service.Principal{SpecialManagementDepCode:"false",IsDefault:false,HasDefaultInit:true,ProfessSalesmanName:pname,ProfessSalesman:pid,SpecialManagementDep:dpid,SpecialManagementDepName:dpname,Status:"Insert"})
        // }
    }

    data.Data.Principals = tmprincipals
    contacts := h.DB.ViewCustomerContact(corpcode)                                                    
    tmpcontactinfo := []service.Merchantcontacterinfo{}  
    for k,v := range contacts{
        if k == 0{
            tmpcontactinfo = append(tmpcontactinfo,service.Merchantcontacterinfo{Isdefault:"true",Hasdefaultinit:true,Mobile: v.Phone,Fullname:service.FullName{ZhCn:v.Name},PositionName: v.Postion,Status:"Insert",TableDisplayOutlineAll:false})    
        }else{
            tmpcontactinfo = append(tmpcontactinfo,service.Merchantcontacterinfo{Isdefault:"false",Hasdefaultinit:true,Mobile: v.Phone,Fullname:service.FullName{ZhCn:v.Name},PositionName: v.Postion,Status:"Insert",TableDisplayOutlineAll:false})
        }
    }                                                  
    data.Data.Merchantcontacterinfos = tmpcontactinfo

    money := 0                                                    
    if strings.Index(customer.Registcapi,"万") != -1{
        smon := strings.Split(customer.Registcapi,"万")[0]
        if strings.Index(smon,".") != -1{
            smon = strings.Split(smon,".")[0]
        }
        money,_ = strconv.Atoi(smon)
    }else if strings.Index(customer.Registcapi,"亿") != -1 {
        smon := strings.Split(customer.Registcapi,"亿")[0] 
        if strings.Index(smon,".") != -1{
            smon = strings.Split(smon,".")[0]
        }
        money,_ = strconv.Atoi(smon)
        money = money *1e8
    }                                                
    
    h.Logger.Debugln("注册资金： ",customer.Registcapi,money)
    data.Data.Money =  money            //注册资金
    data.Data.Buildtime =  customer.Startdate         //成立时间
    data.Data.Scope.ZhCn = customer.Scope 
    data.Data.Regioncode = customer.Belongorg         //注册地区
    // data.Data.Businesslicenseno =                     //经营许可证号
    data.Data.Peoplenum =  customer.Staffs            //员工人数
    data.Data.Verfymark = 1
    data.Data.Contactname =      ""                     //联系人
    data.Data.Contacttel =    ""                        //联系电话


    tmprange := []service.Merchantapplyrange{}
    tmprange = append(tmprange,service.Merchantapplyrange{Rangetype:1,Orgname: gname,Status:"Insert",Iscreator:"true",Isapplied:"true",Hasdefaultinit:true,Orgid: gid})
    data.Data.Merchantapplyranges = tmprange 
    if result,err := h.YS.MerchantPush(data);err != nil{
        h.Logger.Errorln("Error:",err)
    }else{
        h.Logger.Debugln("Result: ",result)
        // customer.Ysid = result.Data.ID 
        customer.Ysapplyid = result.Data.Merchantapplyrangeid
        if err := h.DB.UpdateCustomer(customer);err != nil{
            h.Logger.Errorln(err)
        }
    }	
}


//更新crm客户到erp
func(h *HandlerVar)UpdateToYS(corpcode string){
    //判断用友接口是否正常，失败时睡眠10min后再检查
    for {
        if _,err := h.YS.GetUnits();err != nil{
            h.Logger.Errorln("用友检查错误： ",err," 睡眠10分钟后同步: ",corpcode)
            time.Sleep(time.Minute * 10)
        }else{
            time.Sleep(time.Second * 1)
            break
        }
    }


    cup := h.DB.ViewCustomerUserprincipal(corpcode)
    customer := h.DB.ViewCustomerInfo(corpcode)
    check :=  service.MerchantListData{PageIndex:1,PageSize:100,Name: customer.CorpName}
    merchanlistresult,err := h.YS.MerchantList(&check)
    if err != nil{
        h.Logger.Errorln(err)
    }
    if len(merchanlistresult.Data.RecordList) >= 1{
        for _,v := range merchanlistresult.Data.RecordList{
            merchandetails,err := h.YS.MerchantDetail(strconv.Itoa(int(v.ID)),strconv.Itoa(int(v.MerchantApplyRangeID)))
            if err == nil{
                data := &service.MerchantUpdateData{}
                data.Data.MerchantApplyRangeID = v.MerchantApplyRangeID
                data.Data.Code = v.Code
                data.Data.ID = v.ID 
                data.Data.Iscreator = true
                data.Data.Isapplied = true
                data.Data.CreateOrg =  v.CreateOrg  
                data.Data.CreateorgName = v.CreateOrgName
                data.Data.Belongorg =  v.BelongOrg
                data.Data.Customerclass = 2174258698604800
                data.Data.CustomerclassPath = "2174256368997120|2174258698604800|"
                data.Data.Enterprisename =  customer.CorpName 
                data.Data.Enterprisenature = "0"
                data.Data.Taxpayingcategories = 0
                data.Data.Status = "Update"
                data.Data.Source =  0 

                data.Data.Buildtime =  customer.Startdate         //成立时间
                data.Data.Scope.ZhCn = customer.Scope 
                // data.Data.Regioncode = customer.Belongorg         //注册地区
                data.Data.Verfymark = 1    

                data.Data.Merchantrole = service.UpdateMerchantrole{Status:"Update",ID: v.ID}
                data.Data.Leadername = customer.Opername          //法人代表
                data.Data.Creditcode = customer.Creditcode        //信用代码
                data.Data.Address.ZhCn = customer.Address              //地址
                data.Data.Merchantapplieddetail =  service.UpdateMerchantapplieddetail{Signback:false,
                                                                    Payway:99,
                                                                    // Searchcode:"",
                                                                    MerchantApplyRangeId: v.MerchantApplyRangeID,
                                                                    Customerlevel: customer.CustomerlevelId,
                                                                    // CustomerlevelCode: "xvsf",
                                                                    CustomerlevelName: customer.Customerlevel,
                                                                    Status: "Update",
                                                                    ID: merchandetails.Data.MerchantAppliedDetail.ID,
                                                                    Stopstatus: false,
                                                                    Frozenstate:0,
                                                                    }
                tmprincipals := []service.UpdatePrincipal{}    
                for _,v := range cup{
                    pname,pid,dpname,dpid,err := h.YS.CheckERPUserInfo(strings.Split(v.MerchandiserName,"(")[0])
                    if err != nil{
                        h.Logger.Infoln("ERP查询人员失败: ",err)
                        continue
                    }
                    Flag := true
                    for _,j := range merchandetails.Data.Principals{
                        if pname == j.ProfessSalesmanName{
                            Flag = false
                            break
                        }
                    }
                    h.Logger.Debugln("Pid: ",pid," pname: ",pname,"dpname: ",dpname," dpid: ",dpid)
                    if Flag{
                        // if len(tmprincipals) == 0 {
                        //     tmprincipals = append(tmprincipals,service.UpdatePrincipal{SpecialManagementDepCode:"false",IsDefault:true,HasDefaultInit:true,ProfessSalesmanName:pname,ProfessSalesman:pid,SpecialManagementDep:dpid,SpecialManagementDepName:dpname,Status:"Insert"})
                        // }else{
                            tmprincipals = append(tmprincipals,service.UpdatePrincipal{SpecialManagementDepCode:"false",IsDefault:false,HasDefaultInit:true,ProfessSalesmanName:pname,ProfessSalesman:pid,SpecialManagementDep:dpid,SpecialManagementDepName:dpname,Status:"Insert"})
                        // }
                    }
                }

                // if len(merchandetails.Data.Principals) <= 0{
                //     data.Data.Principals = tmprincipals
                // }
                data.Data.Principals = tmprincipals
                //
                //联系人
                //
                
                contacts := h.DB.ViewCustomerContact(corpcode)                                                    
                tmpcontactinfo := []service.UpdateMerchantcontacterinfo{}  
                for k,v := range contacts{
                    Flag := true
                    for _,j := range merchandetails.Data.MerchantContacterInfos{
                        if j.FullName.ZhCN == v.Name && j.Mobile == v.Phone{
                            Flag = false
                        }
                    }
                    if Flag {
                        if k == 0{
                            tmpcontactinfo = append(tmpcontactinfo,service.UpdateMerchantcontacterinfo{Isdefault:"true",Hasdefaultinit:true,Mobile: v.Phone,Fullname:service.FullName{ZhCn:v.Name},PositionName: v.Postion,Status:"Insert",TableDisplayOutlineAll:false})    
                        }else{
                            tmpcontactinfo = append(tmpcontactinfo,service.UpdateMerchantcontacterinfo{Isdefault:"false",Hasdefaultinit:true,Mobile: v.Phone,Fullname:service.FullName{ZhCn:v.Name},PositionName: v.Postion,Status:"Insert",TableDisplayOutlineAll:false})
                        }
                    }
                }                                                  
                data.Data.Merchantcontacterinfos = tmpcontactinfo 

                money := 0                                                    
                if strings.Index(customer.Registcapi,"万") != -1{
                    smon := strings.Split(customer.Registcapi,"万")[0]
                    if strings.Index(smon,".") != -1{
                        smon = strings.Split(smon,".")[0]
                    }
                    money,_ = strconv.Atoi(smon)
                }else if strings.Index(customer.Registcapi,"亿") != -1 {
                    smon := strings.Split(customer.Registcapi,"亿")[0] 
                    if strings.Index(smon,".") != -1{
                        smon = strings.Split(smon,".")[0]
                    }
                    money,_ = strconv.Atoi(smon)
                    money = money *1e8
                }                                                

                h.Logger.Debugln("注册资金： ",customer.Registcapi,money)
                data.Data.Money =  money            //注册资金

                if _,err := h.YS.MerchantUpdatePush(data);err != nil{
                    h.Logger.Errorln("Error:",err)
                }else{
                    h.Logger.Infoln("成功更新客户!",corpcode,customer.CorpName)  
                }   

            }
        } 
    }
}



func (h *HandlerVar)SyncToRD(identify string){
    h.Logger.Debugln("identify: ",identify)
    if identify == "manager"{
        var customers backend.Results
        tmpcustomers := h.DB.ViewCustomers()
        for _,v := range tmpcustomers{
            tmp := h.DB.ViewCustomerUserprincipal(v.CorpCode)
            records := h.DB.ViewCustomerFollowRecords(v.CorpCode)
            if len(records)>0 && records[len(records)-1].Model.UpdatedAt.Year() > 1973{
                v.RecordTime = records[len(records)-1].Model.UpdatedAt

            }
            if len(tmp) >0 &&  tmp[len(tmp)-1].MerchandiserName !=""{
                v.Genjin = tmp[len(tmp)-1].MerchandiserName
                for _,j := range tmp{
                    v.Genjins = append(v.Genjins,j.MerchandiserName) 
                }
                customers = append(customers,v)
            }
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
        return
    }
    if identify == "scope"{
        customers := h.DB.ViewCustomers()
        for k,v := range customers{
            tmp := h.DB.ViewCustomerUserprincipal(v.CorpCode)
            records := h.DB.ViewCustomerFollowRecords(v.CorpCode)
            if len(tmp) >0 && tmp[len(tmp)-1].Model.UpdatedAt.Year() > 1973{
                customers[k].Genjin = tmp[len(tmp)-1].MerchandiserName
                for _,j := range tmp{
                    customers[k].Genjins = append(customers[k].Genjins,j.MerchandiserName) 
                }
            }
            if len(records)>0 && records[len(records)-1].Model.UpdatedAt.Year() > 1973{
                customers[k].RecordTime = records[len(records)-1].Model.UpdatedAt
            }
            // h.Logger.Debugln("Genjin: ",customers[k].Genjin)
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
        return
    }
    dmaiuser := h.DB.ViewUser(identify)
    if dmaiuser.Name != ""{
        customers := h.DB.ViewUserCustomers(identify)
        for k,v := range customers{
            tmp := h.DB.ViewCustomerContact(v.CorpCode)
            records := h.DB.ViewCustomerFollowRecords(v.CorpCode)
            genjintmp := h.DB.ViewCustomerUserprincipal(v.CorpCode)
            if len(tmp) >0 && tmp[len(tmp)-1].Model.UpdatedAt.Year() > 1973 {
                
                customers[k].Contact = tmp[len(tmp)-1].Name
            }
            if len(records)>0 && records[len(records)-1].Model.UpdatedAt.Year() > 1973{
                customers[k].RecordTime = records[len(records)-1].Model.UpdatedAt
            }
            for _,j := range genjintmp{
                customers[k].Genjins = append(customers[k].Genjins,j.MerchandiserName)
            }
            h.Logger.Debugln("Contact: ",customers[k].Contact)
        }
        sort.Sort(customers)
        if sort.IsSorted(customers){
          h.Logger.Debugln("customers sorted....")
        }else{
          h.Logger.Debugln("customers not sorted....")
        }
        if err := h.Redis.SetKey(identify,"300",customers);err != nil{
            h.Logger.Errorln(err)
        }
    }else{
        h.Logger.Errorln("user not exists!")
    }
}

//客户基本信息企查查查询
func (h *HandlerVar)CustomerBasicInfo(corpcode string){
    customerinfo := h.DB.ViewCustomerInfo(corpcode)
    if customerinfo.CorpName !=  ""{
        companyInfo,err := h.QCC.QGetBasicDetailsByName(customerinfo.CorpName)
        if err == nil{
            customerinfo.Opername = companyInfo.Result.Opername
            customerinfo.Startdate = companyInfo.Result.Startdate
            customerinfo.Province = companyInfo.Result.Province
            customerinfo.Address = companyInfo.Result.Address
            customerinfo.Status = companyInfo.Result.Status 
            customerinfo.Registcapi = companyInfo.Result.Registcapi
            customerinfo.Creditcode = companyInfo.Result.Creditcode
            customerinfo.Scope = companyInfo.Result.Scope
            customerinfo.Econkind = companyInfo.Result.Econkind
            customerinfo.Belongorg = companyInfo.Result.Belongorg 
            customerinfo.Orgno = companyInfo.Result.Orgno 
            customerinfo.Logo = companyInfo.Result.Imageurl 
            if err := h.DB.UpdateCustomer(customerinfo);err != nil{
                h.Logger.Errorln(err)
            }
        }else{
            h.Logger.Errorln(err)
        }
    }
}


//日志文件Error时通知管理人员
func (h *HandlerVar)Notficationer(ctx  context.Context){
    f, err := os.Open(h.Opts.LogFile)
    if err != nil {
        h.Logger.Errorln(err)
    }
    defer f.Close()
    fileinfo,err := f.Stat()
    if err != nil{
        h.Logger.Errorln(err)
    }
    fsize := int(fileinfo.Size())

    rd := bufio.NewReader(f)
    rdsize := rd.Size()
    if fsize < rdsize{
        h.Logger.Debugln("文件比缓存小")
        _,err := rd.Discard(fsize)
        if err !=nil{
            h.Logger.Errorln(err)
        }
    }
    if fsize >= rdsize{
        h.Logger.Debugln("文件比缓存大")
        for i := 1;i <= fsize/rdsize;i++ {
            _,err := rd.Discard(rdsize)
            if err !=nil{
                h.Logger.Errorln(err)
            }
        }
    }

    h.Logger.Debugln("开始监控日志信息！")
    for {
        line, err := rd.ReadString('\n') 
        
        if io.EOF == err {
            time.Sleep(time.Second*1)
        }
        if err != nil && err != io.EOF{
            h.Logger.Errorln(err)
            h.Logger.Debugln("文件监控退出")
            break
        }
       if strings.Index(line,"Error") != -1 {
          _,err := h.WX.SendTextMsg(h.Opts.NotificationUsers,line)
          if err != nil{
            h.Logger.Errorln(err)
          }
       }
       if err := ctx.Done();err != nil{
          h.Logger.Errorln(err)
          return
       } 
    }      
}


func(h *HandlerVar)UploadXlsxFile(f io.Reader,user ,tflag string){
    result := h.Reporter.UpLoadWithHttp(f,tflag)
    success := ""
    failure := ""
    for _,v := range result["success"]{
        success += " " + v
    }
    for _,v := range result["failure"]{
        failure += " " + v
    }
    msgs := ""
    if len(result["failure"]) > 0{
        msgs = fmt.Sprintf("<div class=\"normal\">成功客户:%s</div>失败客户:%s<div class=\"highlight\">因未检测到工商信息，拒绝导入。</div>",success,failure)
    }else{
        msgs = fmt.Sprintf("<div class=\"normal\">成功客户:%s</div>",success)
    }

    _,err := h.WX.SendTextCardMsg(user,msgs,"客户批量上传处理")
    if err != nil{
        h.Logger.Errorln(err)
    }
    h.Redis.DelKey("manager")
    h.UserInfoSyncToRD <- "manager"

    h.Redis.DelKey("scope")
    h.UserInfoSyncToRD <- "scope"
    h.Redis.DelKey(user)
    h.UserInfoSyncToRD <- user
    SeaCustomers = backend.Results{}
}

func(h *HandlerVar)GenJinNotification(data Assign){
    for k,v := range data{
        tmp  := ""
        for _,j := range v{
            if len(tmp) >0{
                tmp += " " + j 
            }else{
                tmp += j 
            }
        }
        msgs := "新客户： "+ tmp
        _,err := h.WX.SendTextCardMsg(k,msgs,"新客户分配")
        if err != nil{
            h.Logger.Errorln(err)
        }
    }
}



func(h *HandlerVar)BasisSearch(customers backend.Results,searchdata map[string]string)backend.Results{
    // basis,keyword string
    tmp := backend.Results{}
    customerstmp := backend.Results{}
    for basis,keyword := range searchdata{
        if len(tmp) == 0{
            customerstmp = customers
        }else{
            customerstmp = tmp 
        }
        tmp = backend.Results{}
        if basis == "region"{
            h.Logger.Debugln("basis is region...",searchdata,"customers tmp: ",customerstmp)
            for _,v := range customerstmp{
                if strings.Index(v.Region,keyword) != -1{
                    tmp = append(tmp,v)
                    h.Logger.Debugln("string index: ",strings.Index(v.Region,keyword),"corpname: ",v.CorpName," key: ",keyword)
                }
            }
        }
        if basis == "customername"{
            h.Logger.Debugln("basis is customer name...",searchdata,"customers tmp: ",customerstmp)
            for _,v := range customerstmp{
                if strings.Index(v.CorpName,keyword) != -1{
                    tmp = append(tmp,v)
                    // h.Logger.Debugln("Match: ",v)
                    h.Logger.Debugln("string index: ",strings.Index(v.CorpName,keyword),"corpname: ",v.CorpName," key: ",keyword)
                }
            }
        }
        if basis == "genjin"{
            h.Logger.Debugln("basis is genjin...",searchdata,"customers tmp: ",customerstmp)
            for _,v := range customerstmp{
                for _,j := range v.Genjins{
                    if strings.Index(j,keyword) != -1{
                        tmp = append(tmp,v)
                        h.Logger.Debugln("string index: ",strings.Index(v.CorpName,keyword),"corpname: ",v.CorpName," key: ",keyword)
                        break
                    }
                }
            }
        }
        if basis == "trade"{
            h.Logger.Debugln("basis is trade...",searchdata,"customers tmp: ",customerstmp)
            for _,v := range customerstmp{
                if strings.Index(v.Trade,keyword) != -1{
                    tmp = append(tmp,v)
                    h.Logger.Debugln("string index: ",strings.Index(v.Region,keyword),"corpname: ",v.CorpName," key: ",keyword)
                }
            }
        }
        if len(tmp) == 0{
            return tmp 
        }
    }
    return tmp 
}


func(h *HandlerVar)BasicReportSender(user string,customerinfo *backend.Customer){
    url := ""
    tnow := time.Now().Unix()
    if customerinfo.BaseReporturl == ""{
        if result,err := h.QCC.QReportBaseEmit(customerinfo.CorpName);err == nil && result.Status =="200"{
            h.Logger.Debugln("order number: ",result.Result.OrderNo)
            for {
                reportresult,_ := h.QCC.QGetReportBase(result.Result.OrderNo)
                if reportresult.Result.ReportStatus == "S" && reportresult.Status =="200"{
                    h.Logger.Debugln("reporter url: ",reportresult.Result.ReportURL)
                    url = reportresult.Result.ReportURL
                    customerinfo.BaseReporturl = url 
                    if err := h.DB.UpdateCustomer(customerinfo);err != nil{
                        h.Logger.Errorln(err)
                    }
                    break
                }
                time.Sleep(5*time.Second)
                if time.Now().Unix() - tnow > 3600{
                    h.Logger.Errorln("企查查客户基础信用报告生成超时: ",customerinfo.CorpName,": ",customerinfo.CorpCode)
                    return
                }
            }
        }else{
            h.Logger.Errorln(err,"result: ",result)
        }
    }else{
        url = customerinfo.BaseReporturl
    }
    if len(url) == 0{
        h.Logger.Errorln("企查查客户基础信用报告生成错误: ",customerinfo.CorpName,": ",customerinfo.CorpCode)
        return
    }
    resp,err := http.Get(url)
    if err != nil{
        h.Logger.Errorln(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil{
        h.Logger.Errorln(err)
    }
    if err := ioutil.WriteFile(h.Opts.SrcDir+"/"+customerinfo.CorpCode+".pdf", body, 0644);err != nil {
        h.Logger.Errorln(err)
    }

    tmpdata := bytes.NewBuffer(body)
    result,err := h.WX.UploadFile(tmpdata,customerinfo.CorpName+".pdf")
    if err !=nil{
        h.Logger.Errorln(err)
    }
    if _,err := h.WX.SendFileMsg(result.MediaID,user);err != nil{
        h.Logger.Errorln(err)
    }
    h.Logger.Infoln("报表生成并已发送到用户: "+ user +" 企业微信！")
}
















