package service
import(
    "io"
  "fmt"
  "time"
  "sort"
  "bytes"
  "strings"
  "strconv"
  // "context"
//  "gorm.io/driver/mysql"
//  "gorm.io/gorm"
  "wxcrm/pkg/backend"
  "wxcrm/pkg/common"
  "wxcrm/pkg/common/log"
  "github.com/360EntSecGroup-Skylar/excelize/v2"
)


type Reporter struct{
    Logger  *log.Logger
    DB      *backend.DB
    Opts    *common.Opts
    YS      *YS 
    QCC     *QCC
    C1      chan string 
}

type ProjectRecord struct{
	CustomerName    string 
	Level           string 
	Area            string 
	Contacter       string 
	ProjectName     string 
	Budget          int 
	Timeplan        string 
	Product         string 
	Presaler        string 
	Saler           string
	Logs            Logs     
}


type Log struct{
	Followlog       string 
	Time            time.Time 
}
type Logs []Log

func (s Logs) Len() int {
    return len(s)
}

func (s Logs) Less(i, j int) bool {
    return s[i].Time.Before(s[j].Time)
}

func (s Logs) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}




func NewReporter(log *log.Logger,db *backend.DB,opts *common.Opts,ys *YS,qcc *QCC,ch1 chan string)*Reporter{
    return &Reporter{
        Logger: log,
        DB: db,
        Opts: opts,
        QCC: qcc,
        YS: ys,
        C1: ch1,
    }
}


func (r *Reporter)GenReport(username,start,end string)(*bytes.Buffer,error){

    sheetData := map[string][]ProjectRecord{}
    f := excelize.NewFile()
    // f.DeleteSheet("Sheet1")
    f.SetRowHeight("Sheet1",1,46.25)
    f.SetDefaultFont("微软雅黑")
    genjin, _ := f.NewStyle(`{
        "border": [
        {
            "type": "left",
            "color": "000000",
            "style": 1
        },
        {
            "type": "top",
            "color": "000000",
            "style": 1
        },
        {
            "type": "bottom",
            "color": "000000",
            "style": 1
        },
        {
            "type": "right",
            "color": "000000",
            "style": 1
        }],
        "alignment": {
            "horizontal": "left",
            "ident": 0,
            "justify_last_line": false,
            "reading_order": 0,
            "relative_indent": 1,
            "shrink_to_fit": true,
            "text_rotation": 0,
            "vertical": "top",
            "wrap_text": true
        },
        }`)
       
    other, _ := f.NewStyle(`{
        "border": [
        {
            "type": "left",
            "color": "000000",
            "style": 1
        },
        {
            "type": "top",
            "color": "000000",
            "style": 1
        },
        {
            "type": "bottom",
            "color": "000000",
            "style": 1
        },
        {
            "type": "right",
            "color": "000000",
            "style": 1
        }],
        "alignment": {
            "horizontal": "center",
            "ident": 0,
            "justify_last_line": false,
            "reading_order": 0,
            "relative_indent": 1,
            "shrink_to_fit": true,
            "text_rotation": 0,
            "vertical": "center",
            "wrap_text": true
        },
        "fill":{"type":"pattern","color":["#FFFFFF"],"pattern":1},
        "font": {"size":11,"Family":"微软雅黑","bold": false}
        }`)    
     
    header, _ := f.NewStyle(`{
        "border": [
        {
            "type": "left",
            "color": "000000",
            "style": 1
        },
        {
            "type": "top",
            "color": "000000",
            "style": 1
        },
        {
            "type": "bottom",
            "color": "000000",
            "style": 1
        },
        {
            "type": "right",
            "color": "000000",
            "style": 1
        }],
        "alignment": {
            "horizontal": "center",
            "ident": 0,
            "justify_last_line": false,
            "reading_order": 0,
            "relative_indent": 1,
            "shrink_to_fit": true,
            "text_rotation": 0,
            "vertical": "center",
            "wrap_text": true
        },
        "fill":{"type":"pattern","color":["#CD5555"],"pattern":1},
        "font": {"size":14,"Family":"微软雅黑","bold":true}
        }`)

    sumer, _ := f.NewStyle(`{
        "border": [
        {
            "type": "left",
            "color": "000000",
            "style": 1
        },
        {
            "type": "top",
            "color": "000000",
            "style": 1
        },
        {
            "type": "bottom",
            "color": "000000",
            "style": 1
        },
        {
            "type": "right",
            "color": "000000",
            "style": 1
        }],
        "alignment": {
            "horizontal": "center",
            "ident": 0,
            "justify_last_line": false,
            "reading_order": 0,
            "relative_indent": 1,
            "shrink_to_fit": true,
            "text_rotation": 0,
            "vertical": "center",
            "wrap_text": true
        },
        "fill":{"type":"pattern","color":["#CD5555"],"pattern":1},
        "font": {"size":14,"Family":"微软雅黑","bold": true}
        }`)

    s := "序号 优先级 客户名 地区 联系 项目名 项目现状及后续跟进 预计金额(万) 时间计划 产品线 销售负责人 售前负责人"
    title := strings.Split(s," ")

    f.SetCellValue("Sheet1","A"+strconv.Itoa(1),title[0])
    f.SetCellValue("Sheet1","B"+strconv.Itoa(1),title[1])
    f.SetCellValue("Sheet1","C"+strconv.Itoa(1),title[2])
    f.SetCellValue("Sheet1","D"+strconv.Itoa(1),title[3])
    f.SetCellValue("Sheet1","E"+strconv.Itoa(1),title[4])
    f.SetCellValue("Sheet1","F"+strconv.Itoa(1),title[5])
    f.SetCellValue("Sheet1","G"+strconv.Itoa(1),title[6])
    f.SetCellValue("Sheet1","H"+strconv.Itoa(1),title[7])
    f.SetCellValue("Sheet1","I"+strconv.Itoa(1),title[8])
    f.SetCellValue("Sheet1","J"+strconv.Itoa(1),title[9])
    f.SetCellValue("Sheet1","K"+strconv.Itoa(1),title[10])
    f.SetCellValue("Sheet1","L"+strconv.Itoa(1),title[11])

    if err := f.SetCellStyle("Sheet1","A1","L1",header);err != nil{
    	r.Logger.Errorln(err)
    	return nil,err
    }
    if err := f.SetColWidth("Sheet1", "G", "G", 100);err != nil{
    	r.Logger.Errorln(err)
    	return nil,err
    } 

    projects := []backend.FollowRecord{}
    if username == "manager"{
        projects = r.DB.ViewProjects()
    }else{
        projects = r.DB.ViewUserProjects(username)
    }
   r.Logger.Debugln("Gen reporter projects: ",projects)

   ysprojects := map[string]string{} 
   input := &YSProjectListData{PageIndex:1,PageSize:10000}
   if ysprojectemps,err := r.YS.GetProjectList(input);err == nil{
     for _, class := range ysprojectemps.Data.RecordList{
        ysprojects[class.Name] = class.ClassifyidName
     }
   }
   


    for _,v := range projects{
        presaler := ""
        saler := ""
        contactname := ""
        records := r.DB.ViewProjectTimeFollowRecords(v.CorpCode,v.Projectid,start,end)
        customer := r.DB.ViewCustomerInfo(v.CorpCode)
        contacts := r.DB.ViewCustomerContact(v.CorpCode)
        timebudget := r.DB.ViewTimeBudget(v.CorpCode,v.Projectid)
        principals  :=  r.DB.ViewCustomerUserprincipal(v.CorpCode)
       
        if len(principals) >= 1{
            for _,v := range principals{
                if v.Presaler == "true"{
                    presaler = v.MerchandiserName
                }
                if v.Saler == "true"{
                    saler = v.MerchandiserName
                }
            }
            if presaler ==""{
                presaler = principals[0].MerchandiserName
            }
            if saler ==""{
                saler = principals[len(principals)-1].MerchandiserName
            }
        }

        if len(contacts) >0 {
            contactname = contacts[len(contacts)-1].Name
        }
       
       budget,err := strconv.Atoi(timebudget[len(timebudget)-1].Budget)
       if err != nil{
        r.Logger.Errorln(err)
       }
        tmp := ProjectRecord{
            CustomerName: customer.CorpName,
            Level: customer.Customerlevel,
            Area: customer.Province,
            Contacter: contactname,
            ProjectName: v.Projectname ,
            Budget: budget,
            Timeplan: timebudget[len(timebudget)-1].Timeplan,
            Product: ysprojects[v.Projectname],
            Presaler: presaler,
            Saler: saler, 
        }
        logs := Logs{}
        for _,log := range records{
            logs = append(logs,Log{Followlog:log.Followlog,Time: log.Model.UpdatedAt}) 
        }
        tmp.Logs = logs 
        if _,ok := sheetData[customer.Customerlevel];ok{
            sheetData[customer.Customerlevel] = append(sheetData[customer.Customerlevel],tmp)
        }else{
            sheetData[customer.Customerlevel] = []ProjectRecord{}
            sheetData[customer.Customerlevel] = append(sheetData[customer.Customerlevel],tmp)
        }
    }


    count := 2 
    leveslen := 0
    endtime,_ := strconv.Atoi(end)
    r.Logger.Debugln("Sheet Data: ",sheetData)
    for levelname,items := range sheetData{
        for i:=0;i<len(items);i++{
            sort.Sort(items[i].Logs)
            richtext := []excelize.RichTextRun{}
            if sort.IsSorted(items[i].Logs){
                // if t1 - items[i].Logs[len(items[i].Logs)-1].Time.Unix() > 86400*365{
                //         r.Logger.Errorln("latest followlogs  was pass one year  and skip: ",items[i].ProjectName)
                //         continue   
                // }
                loglength := len(items[i].Logs)
                for k,j := range items[i].Logs{
                    if int64(endtime) - j.Time.Unix() < 86400*7{
                        if k+1 == loglength{
                            s := excelize.RichTextRun{ Text: strconv.Itoa(k+1)+j.Followlog,Font: &excelize.Font{Bold:false,Color: "FF0000",Family: "微软雅黑",Size:11}}
                            richtext = append(richtext,s)
                        }else{
                            s := excelize.RichTextRun{ Text: strconv.Itoa(k+1)+j.Followlog +"\n",Font: &excelize.Font{Bold:false,Color: "FF0000",Family: "微软雅黑",Size:11}}
                            richtext = append(richtext,s)
                        }

                    }else{
                        if k+1 == loglength{
                            s := excelize.RichTextRun{ Text: strconv.Itoa(k+1)+j.Followlog,Font: &excelize.Font{Bold:false,Color: "000000",Family: "微软雅黑",Size:11}}
                            richtext = append(richtext,s)
                        }else{
                            s := excelize.RichTextRun{ Text: strconv.Itoa(k+1)+j.Followlog +"\n",Font: &excelize.Font{Bold:false,Color: "000000",Family: "微软雅黑",Size:11}}
                            richtext = append(richtext,s)
                        }
                    }
                }
            }else{
                r.Logger.Infoln("follow logs slice not sorted and break out reporter program.")
                return nil,fmt.Errorf("follow logs slice not sorted and break out reporter program.")
            }
            f.SetCellValue("Sheet1","A"+strconv.Itoa(leveslen + i+count), leveslen+i+1)
            f.SetCellValue("Sheet1","B"+strconv.Itoa(leveslen + i+count),items[i].Level)
            f.SetCellValue("Sheet1","C"+strconv.Itoa(leveslen + i+count),items[i].CustomerName)
            f.SetCellValue("Sheet1","D"+strconv.Itoa(leveslen + i+count),items[i].Area)
            f.SetCellValue("Sheet1","E"+strconv.Itoa(leveslen + i+count),items[i].Contacter)
            f.SetCellValue("Sheet1","F"+strconv.Itoa(leveslen + i+count),items[i].ProjectName)
            if err := f.SetCellStyle("Sheet1","A"+strconv.Itoa(leveslen+i+count),"F"+strconv.Itoa(leveslen+i+count),other);err != nil{
                r.Logger.Errorln(err)
                return nil,err
            }    
            f.SetCellRichText("Sheet1","G"+strconv.Itoa(leveslen+i+count),richtext)
            if err := f.SetCellStyle("Sheet1","G"+strconv.Itoa(leveslen+i+count),"G"+strconv.Itoa(leveslen+i+count),genjin);err != nil{
                r.Logger.Errorln(err)
                return nil,err
            }    
            f.SetCellValue("Sheet1","H"+strconv.Itoa(leveslen+i+count),items[i].Budget)
            f.SetCellValue("Sheet1","I"+strconv.Itoa(leveslen+i+count),items[i].Timeplan)
            f.SetCellValue("Sheet1","J"+strconv.Itoa(leveslen+i+count),items[i].Product)
            f.SetCellValue("Sheet1","K"+strconv.Itoa(leveslen+i+count),items[i].Saler)
            f.SetCellValue("Sheet1","L"+strconv.Itoa(leveslen+i+count),items[i].Presaler)
            if err := f.SetCellStyle("Sheet1","H"+strconv.Itoa(leveslen+i+count),"L"+strconv.Itoa(leveslen+i+count),other);err != nil{
                r.Logger.Errorln(err)
                return nil,err 
            } 
        }

        if len(items) > 0{
            f.SetCellValue("Sheet1","G"+strconv.Itoa(leveslen+len(items)+count),levelname +"类客户总计：")
            if len(items) == 1 {
                f.SetCellValue("Sheet1","H"+strconv.Itoa(leveslen+count+len(items)),items[0].Budget)
            }else{
                f.SetCellFormula("Sheet1","H"+strconv.Itoa(leveslen+count+len(items)),"SUM(Sheet1!H" +strconv.Itoa(count+leveslen) +",Sheet1!H"+strconv.Itoa(leveslen+len(items)+count-1)+")")
            }
            
            if err := f.SetCellStyle("Sheet1","A"+strconv.Itoa(leveslen+len(items)+count),"L"+strconv.Itoa(leveslen+len(items)+count),sumer);err != nil{
                r.Logger.Errorln(err)
                return nil,err 
            }  
        }
        leveslen = leveslen + len(items)
        if len(items) > 0{
            count++ 
        }
    }
        // if err := f.SaveAs(r.Opts.ExcelFile);err != nil{
        //   r.Logger.Errorln(err)
        // }
    data,err := f.WriteToBuffer()   //(*bytes.Buffer, error)
    if err != nil{
        return data,err 
    }
    return data,nil 
}

//导出客户资料




//Excel导入客户资料
func (r *Reporter)UpLoadWithExcel(f *excelize.File,tflag string)(result map[string][]string){
        // OpenReader(r io.Reader, opt ...Options) (*File, error)
        // excelize.OpenFile("Book1.xlsx")
            // 获取 Sheet1 上所有单元格
            result = map[string][]string{"success":[]string{},"failure":[]string{}}
            rows, err := f.GetRows("Sheet1")
            if err != nil {
                r.Logger.Debugln(err)
                return result
            }

            for _, row := range rows {
                Customer := backend.Customer{}
                Contact  :=  backend.Contactor{}
                cup      :=  backend.CustomerUserprincipal{}

                contactname := row[0]
                customername := row[1]
                Position := row[2]
                phone1 := row[3]
                phone2 := row[4]
                email := row[5]
                cupname := row[6]

                if customername=="客户名称" ||customername==""{
                    continue
                }
                r.Logger.Debugln("查询客户： ",customername)
                check :=  MerchantListData{PageIndex:1,PageSize:100,Name:customername}
                merchanlistresult,err := r.YS.MerchantList(&check)
                if err != nil{
                    r.Logger.Errorln(err)
                }

                customertmp := r.DB.CheckCustomer(customername)
                r.Logger.Debugln("返回客户列表为：",merchanlistresult)
                r.Logger.Debugln("客户检查结果为：",customertmp)
 
                if customertmp.CorpName != "" && len(merchanlistresult.Data.RecordList)>0{
                    r.Logger.Infoln(customername,"在ERP跟CRM中已经存在!")
                    if tflag != "sea"{
                        if cupname == ""{
                            r.Logger.Debugln("跟进人信息为空")
                        }else{
                            cupuserinfo :=r.DB.ViewUserName(cupname+"%")
                            r.Logger.Debugln("用户信息：",cupuserinfo)
                            cupinfo := r.DB.ViewSingleCP(cupuserinfo.Userid,customertmp.CorpCode)
                            r.Logger.Debugln("跟进用户信息：",cupinfo)
                            if cupinfo.Merchandiserid  =="" &&cupuserinfo.Userid !=""{
                                //添加相对应的销售 客户对应表
                                cup.CorpCode = customertmp.CorpCode
                                cup.Merchandiserid = cupuserinfo.Userid
                                cup.MerchandiserName = cupuserinfo.Name 
                                cup.Createby =  cupuserinfo.Userid
                                if err := r.DB.AddCustomerUserprincipal(&cup);err !=nil{
                                    r.Logger.Errorln(err)
                                }
                            }else{
                                r.Logger.Debugln("未添加跟进用户信息！")
                            }
                        }
                    }
                    contactmps := r.DB.ViewCustomerContact(customertmp.CorpCode)
                    Flag := true
                    for _,v := range contactmps{
                        if v.Name == contactname || phone1 == ""{
                            Flag = false
                            break
                        }
                    }
                    if Flag {
                       //添加对应的联系人
                        Contact.CorpCode =  customertmp.CorpCode
                        Contact.Name = contactname
                        Contact.ContactCode = "H" + common.GenUid()
                        Contact.Postion = Position
                        Contact.Phone = phone1
                        Contact.Phone2 = phone2
                        Contact.Email = email
                        if err := r.DB.AddContactor(&Contact); err != nil {
                            r.Logger.Errorln(err)
                        }
                    }

                    result["success"] = append(result["success"],customername)
                    r.C1 <- customertmp.CorpCode
                    continue
                }

                //判断ERP系统是否存在客户档案
                if customertmp.CorpName == "" && len(merchanlistresult.Data.RecordList)>0 && merchanlistresult.Data.RecordList[0].Name.ZhCN == customername{
                    r.Logger.Infoln(customername,"ERP中存在,但CRM中没有!")
                    Customer.CorpName = customername
                    Customer.CorpCode = "W" + common.GenUid()
                    //这里需要加企查查数据 或另外存放一份企查查数据
                    companyInfo,err := r.QCC.QGetBasicDetailsByName(customername)
                    if err == nil && companyInfo.Status == "200"{
                        Customer.Opername = companyInfo.Result.Opername
                        Customer.Startdate = companyInfo.Result.Startdate
                        Customer.Province = companyInfo.Result.Province
                        Customer.Address = companyInfo.Result.Address
                        Customer.Status = companyInfo.Result.Status 
                        Customer.Registcapi = companyInfo.Result.Registcapi
                        Customer.Creditcode = companyInfo.Result.Creditcode
                        Customer.Scope = companyInfo.Result.Scope
                        Customer.Econkind = companyInfo.Result.Econkind
                        Customer.Belongorg = companyInfo.Result.Belongorg 
                        Customer.Orgno = companyInfo.Result.Orgno 
                        Customer.Logo = companyInfo.Result.Imageurl 
                        Customer.Region = common.RegionCheck(companyInfo.Result.Province)
                        Customer.Trade = common.TradeCheck(customername,companyInfo.Result.Scope)
                    }else{
                        r.Logger.Errorln(companyInfo.Message)
                        result["failure"] = append(result["failure"],customername)
                        continue
                    }
                    if err := r.DB.AddCustomer(&Customer); err != nil {
                        r.Logger.Errorln(err)
                    }

                    Contact.CorpCode =  Customer.CorpCode
                    Contact.Name = contactname
                    Contact.ContactCode = "H" + common.GenUid()
                    Contact.Postion = Position
                    Contact.Phone = phone1
                    Contact.Phone2 = phone2
                    Contact.Email = email
                    if err := r.DB.AddContactor(&Contact); err != nil {
                        r.Logger.Errorln(err)
                    }

                    //添加相对应的销售 客户对应表
                    dmaiuser := r.DB.ViewUserName(cupname)
                    cup.CorpCode = Customer.CorpCode
                    cup.Merchandiserid = dmaiuser.Userid
                    cup.MerchandiserName = dmaiuser.Name 
                    cup.Createby =  dmaiuser.Userid
                    if err := r.DB.AddCustomerUserprincipal(&cup);err !=nil{
                        r.Logger.Errorln(err)
                    }
                    r.C1 <- Customer.CorpCode
                    result["success"] = append(result["success"],customername)
                    continue
                }

                //判断ERP系统是否存在客户档案
                if customertmp.CorpName != "" && len(merchanlistresult.Data.RecordList)==0{
                    r.Logger.Infoln(customername,"CRM中存在,但ERP中没有!")
                    if tflag != "sea"{
                        if cupname == ""{
                            r.Logger.Debugln("跟进人信息为空")
                        }else{
                            cupuserinfo :=r.DB.ViewUserName(cupname+"%")
                            cupinfo := r.DB.ViewSingleCP(cupuserinfo.Userid,customertmp.CorpCode)
                            if cupinfo.Merchandiserid  =="" &&cupuserinfo.Userid !=""{
                                //添加相对应的销售 客户对应表
                                cup.CorpCode = customertmp.CorpCode
                                cup.Merchandiserid = cupuserinfo.Userid
                                cup.MerchandiserName = cupuserinfo.Name 
                                cup.Createby =  cupuserinfo.Userid
                                if err := r.DB.AddCustomerUserprincipal(&cup);err !=nil{
                                    r.Logger.Errorln(err)
                                }
                            }else{
                                r.Logger.Errorln("未添加跟进用户信息！")
                            }
                        }
                    }

                    contactmps := r.DB.ViewCustomerContact(customertmp.CorpCode)
                    Flag := true
                    for _,v := range contactmps{
                        if v.Name == contactname || phone1 == ""{
                            Flag = false
                            break
                        }
                    }
                    if Flag {
                       //添加对应的联系人
                        Contact.CorpCode =  customertmp.CorpCode
                        Contact.Name = contactname
                        Contact.ContactCode = "H" + common.GenUid()
                        Contact.Postion = Position
                        Contact.Phone = phone1
                        Contact.Phone2 = phone2
                        Contact.Email = email
                        if err := r.DB.AddContactor(&Contact); err != nil {
                            r.Logger.Errorln(err)
                        }
                    }
                    r.C1 <- customertmp.CorpCode
                    result["success"] = append(result["success"],customername)
                    continue
                }

                Customer.CorpName = customername
                Customer.CorpCode = "W" + common.GenUid()
                //这里需要加企查查数据 或另外存放一份企查查数据
                companyInfo,err := r.QCC.QGetBasicDetailsByName(customername)
                if err == nil && companyInfo.Status == "200"{
                    Customer.Opername = companyInfo.Result.Opername
                    Customer.Startdate = companyInfo.Result.Startdate
                    Customer.Province = companyInfo.Result.Province
                    Customer.Address = companyInfo.Result.Address
                    Customer.Status = companyInfo.Result.Status 
                    Customer.Registcapi = companyInfo.Result.Registcapi
                    Customer.Creditcode = companyInfo.Result.Creditcode
                    Customer.Scope = companyInfo.Result.Scope
                    Customer.Econkind = companyInfo.Result.Econkind
                    Customer.Belongorg = companyInfo.Result.Belongorg 
                    Customer.Orgno = companyInfo.Result.Orgno 
                    Customer.Logo = companyInfo.Result.Imageurl 
                    Customer.Region = common.RegionCheck(companyInfo.Result.Province)
                    Customer.Trade = common.TradeCheck(customername,companyInfo.Result.Scope)
                }else{
                    r.Logger.Debugln(companyInfo.Message,customername)
                    result["failure"] = append(result["failure"],customername)
                    continue
                }
                if err := r.DB.AddCustomer(&Customer); err != nil {
                    r.Logger.Errorln(err)
                }
                //添加对应的联系人
                Contact.CorpCode =  Customer.CorpCode
                Contact.Name = contactname
                Contact.ContactCode = "H" + common.GenUid()
                Contact.Postion = Position
                Contact.Phone = phone1
                Contact.Phone2 = phone2
                Contact.Email = email
                if err := r.DB.AddContactor(&Contact); err != nil {
                    r.Logger.Errorln(err)
                }

                //添加相对应的销售 客户对应表
                if tflag != "sea"{
                    if cupname == ""{
                        r.Logger.Debugln("跟进人信息为空")
                        continue
                    }
                    dmaiuser := r.DB.ViewUserName(cupname+"%")
                    cup.CorpCode = Customer.CorpCode
                    cup.Merchandiserid = dmaiuser.Userid
                    cup.MerchandiserName = dmaiuser.Name 
                    cup.Createby =  dmaiuser.Userid
                    if err := r.DB.AddCustomerUserprincipal(&cup);err !=nil{
                        r.Logger.Errorln(err)
                    }
                }
                r.C1 <- Customer.CorpCode
                r.Logger.Debugln("import success : ",Customer.CorpName)
                result["success"] = append(result["success"],Customer.CorpName)
            }
            r.Logger.Debugln("客户导入状态: ",result)
            return result
}



func (r *Reporter)UpLoadWithHttp(reader io.Reader,tflag string)map[string][]string{
    f, err := excelize.OpenReader(reader)
    if err != nil {
        fmt.Println(err)
        return nil  
    }
    result := r.UpLoadWithExcel(f,tflag)
    return result
}

func (r *Reporter)UpLoadWithFile(filename string){
// OpenReader(r io.Reader, opt ...Options) (*File, error)
// excelize.OpenFile("Book1.xlsx")
    f, err := excelize.OpenFile(filename)
    if err != nil {
        fmt.Println(err)
        return
    }
    tflag := ""
    r.UpLoadWithExcel(f,tflag)
}
