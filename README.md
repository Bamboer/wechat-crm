CRM-Backend
===========
WXCRM后端

# API信息
## 目录
- [本系统用户](#查看本系统用户)
  - [查看用户](#查看本系统用户)
  - [修改用户信息](#修改本系统用户信息)
- [客户](#查看当前用户权限范围内的客户列表)
  - [查看当前用户权限范围内的客户列表](#查看当前用户权限范围内的客户列表)
  - [添加客户](#添加客户)
  - [更新客户](#更新客户)
- [客户的联系人](#查看某客户的联系人)
  - [查看某客户的联系人](#查看某客户的联系人)
  - [查看某个联系人信息](#查看某客户的联系人)
  - [增加联系人](#增加联系人)
  - [修改联系人信息](#修改联系人信息)
  - [删除某个联系人信息](#删除某个联系人信息)

- [企查查](#企查查模糊查询)
  - [企查查模糊查询](#企查查模糊查询)

- [产品](#查看产品信息)
  - [查看产品列表](#查看产品列表) 
  - [增加产品信息](#增加产品信息)        
  - [删除产品信息](#删除产品信息)     

- [合同](#查看合同信息)
  - [查看合同信息](#查看合同信息)
  - [增加合同信息](#增加合同信息)
  - [更新合同信息](#更新合同信息)
  - [删除合同信息](#删除合同信息)      

- [项目](#查看项目信息)
  - [查看项目信息](#查看项目信息)
  - [增加项目信息](#增加项目信息)   
  - [更新项目信息](#更新项目信息)   
  - [删除项目信息](#删除项目信息)               

- [跟进记录](#查看跟进记录信息)
  - [查看跟进记录信息](#查看跟进记录信息)
  - [增加跟进记录信息](#增加跟进记录信息)   
  - [更新跟进记录信息](#更新跟进记录信息)   
  - [删除跟进记录信息](#删除跟进记录信息)  

- [客户跟进人](#查看客户-跟进人)
  - [查看客户 跟进人](#查看客户-跟进人)
  - [添加客户 跟进人](#添加客户-跟进人)   
  - [更新客户 跟进人](#更新客户-跟进人)   
  - [删除客户 跟进人](#删除客户-跟进人)   

- [查看用户操作记录](#查看用户操作记录)
- [获取当前企业微信联系人](#获取当前用户企业微信联系人)

- [获取客户级别](#获取客户级别)
- [check校验](#check-校验)
- [获取Jsticket](#获取-Jsticket)
- [获取销售报表下载链接](#获取销售报表下载链接)
- [同步客户信息到用友](#同步客户信息到用友)




```
Model:{
	CreateAt: time,
	UpdateAt: time,
	DeleteAt: time,
} 
```

### 获取当前用户登录状态
```
req: GET https://tcrm-iti.dm-ai.com/api/logincrm
res: {
    "status": "200",                     status: 200表示成功登录
    "CookieName": "wxcrm_cookie",     
    "Userinfo":dmaiuser                    用户信息
    }   

失败示例： {"status": "400","CookieName": "wxcrm_cookie","Userinfo":"","ErrMsg": err}
```


### 查看本系统用户
```
req: GET https://tcrm-iti.dm-ai.com/api/user
res: {
   Dmaiusers:[{
    Name :string           姓名
    Gender :string         性别 0 未知 1男 2女
    OpenUserid:string      微信id
    Userid :string         企业微信用户Id
    Position:string        职位
    Mobile :string         手机号
    Email :string          邮箱
    Alias :string          别称
    Logo :string           头像链接
    Address :string        地址
    Manager :string        管理员;default:false
    Enabled :string        员工状态;default:true
   },]
   "status": "200"
}
失败示例：{"status": "400","Dmaiusers":[],"ErrMsg": err}
```



### 修改本系统用户信息  
```
req: POST https://tcrm-iti.dm-ai.com/api/user?action=update
body:{
    Name :string           姓名
    Gender :string         性别 0 未知 1男 2女
    OpenUserid:string      微信id
    Userid :string         企业微信用户Id
    Position:string        职位
    Mobile :string         手机号
    Email :string          邮箱
    Alias :string          别称
    Logo :string           头像链接
    Address :string        地址
    Manager :string        管理员;default:false
    Enabled :string        员工状态;default:true
  }
res: {"status":  errocode ,"ErrMsg": errmessage}
```




### 查看当前用户权限范围内的客户列表
```
req: GET https://tcrm-iti.dm-ai.com/api/customer 
res: {
	customers:[{
    CorpCode:string       客户编码
    CorpName:string       客户名称
    Opername:string       法人名称
    Startdate:string      营业起始时间
    Province:string       省
    Address:string        详细地址
    Status:string         在营状态
    Registcapi:string     注册资金
    Creditcode:string     信用代码
    Scope:string          经营范围
    Econkind:string       公司类型
    Belongorg:string      归属管理组织
    Orgno:string          组织代码
    Logo:string           logo地址
    Customerlevel:string  客户级别
    CustomerlevelId:string  客户级别id
    SuperiorCustomer：string   上级客户
    Trade：string        客户行业
    Investor：string      投资者
    CompetingGoodsInfo：string        竞品情报
    BusinessRequirements：string      业务需求
    Comments：string        备注
    Telephone：string       固定电话
    Dndisturb：string       免打扰;default:false
    Model
  },]
  status: string  200正常返回
}

```

### 添加客户
```
req: POST https://tcrm-iti.dm-ai.com/api/customer?action=add
body: {
  	Customer:{
    CorpName:string       客户名称  *
    Customerlevel:string  客户级别*
    CustomerlevelId:string  客户级别id *
    SuperiorCustomer：string   上级客户
    Trade：string        客户行业
    CompetingGoodsInfo：string        竞品情报
    BusinessRequirements：string      业务需求
    Comments：string        备注
    Telephone：string       固定电话 
  }，
   Contactor: {            
  	Name: string   联系人名字 当联系人为输入联系人时必填*
  	Contactid：string 联系人ID   当联系人为企业微信联系人时，此id为WX ExternalID *
    Gender：string   性别 0 未知 1男 2女
    Logo:string    联系人头像
    Department:string   部门
    Postion:string    职务
    Phone1:string    手机号1
    Phone2:string   手机号2
    Email:string    邮箱
    QQ:string   QQ号
    Wechat:string   微信号
    Iswx:string     联系人是否默认为企业微信联系人
    Comments:string   备注
    Address:string   地址
    ParentDepartment:string 上级部门
    BirthDay:string   生日
    Role:string   联系人角色
    Dndisturb:string   免打扰;default:false
  }
}
res: {result: "true|false"}
```


### 更新客户
```
req: POST https://tcrm-iti.dm-ai.com/api/customer?action=update
body: {
  	CorpCode： string      客户编码*
    Customerlevel: string  客户级别
    CustomerlevelId: string  客户级别id
    SuperiorCustomer： string   上级客户
    Trade： string        客户行业
    CompetingGoodsInfo： string        竞品情报
    BusinessRequirements： string      业务需求
    Comments： string        备注
    Telephone： string       固定电话
  }
res: {result: "true|false"}
```



### 查看某客户的联系人
```
req: GET https://tcrm-iti.dm-ai.com/api/contact?action=view&corpcode=XXX
res: {
	Contacts:[{
    CorpCode:string     客户编码
    Name:string      姓名
    Contactid:string     联系人id 同时存放企业微信联系人id跟手动添加联系人id
    Gender:string      性别 0 未知 1男 2女
    Logo:string      联系人头像
    Department:string     部门
    Postion:string      职务
    Phone1:string      手机号
    Phone2:string      手机号
    Email:string      邮箱
    QQ:string      QQ号
    Wechat:string      微信号
    Iswx:string      默认为企业微信联系人
    Comments:string      备注
    Address:string      地址
    ParentDepartment:string      上级部门
    BirthDay:string      生日
    Role :string      角色 //联系人角色
    Dndisturb :string      免打扰;default:false
    Enabled  :string      默认是否有效;default:true
    Model            
  }],
  CustomerLog: string 客户logo url地址
}
```


### 查看某个联系人信息
```
req: GET https://tcrm-iti.dm-ai.com/api/contact?action=info&contactid=XXX
res: {
	Contacts:{
    CorpCode:string     客户编码
    Name:string      姓名
    Contactid:string     联系人id 同时存放企业微信联系人id跟手动添加联系人id
    Gender:string      性别 0 未知 1男 2女
    Logo:string      联系人头像
    Department:string     部门
    Postion:string      职务
    Phone1:string      手机号
    Phone2:string      手机号
    Email:string      邮箱
    QQ:string      QQ号
    Wechat:string      微信号
    Iswx:string      默认为企业微信联系人
    Comments:string      备注
    Address:string      地址
    ParentDepartment:string      上级部门
    BirthDay:string      生日
    Role :string      角色 //联系人角色
    Dndisturb :string      免打扰;default:false
    Enabled  :string      默认是否有效;default:true
    Model            
  }
}
```


### 增加联系人
```
req: POST https://tcrm-iti.dm-ai.com/api/contact?action=add
body:{
    CorpCode:string     客户编码 *
    Name:string      姓名  *
    Contactid:string     联系人id 同时存放企业微信联系人id跟手动添加联系人id(如果为企业微信联系人此字段为必须，反之不需要生成此字段) *
    Gender:string      性别 0 未知 1男 2女
    Logo:string      联系人头像
    Department:string     部门
    Postion:string      职务
    Phone1:string      手机号
    Phone2:string      手机号
    Email:string      邮箱
    QQ:string      QQ号
    Wechat:string      微信号
    Iswx:string      默认为企业微信联系人
    Comments:string      备注
    Address:string      地址
    ParentDepartment:string      上级部门
    BirthDay:string      生日
    Role :string      角色 //联系人角色
    Dndisturb :string      免打扰;default:false
    Enabled  :string      默认是否有效;default:true
    Model 
  }
res: {result: "false|true"}
```



### 修改联系人信息
```
req: POST https://tcrm-iti.dm-ai.com/api/contact?action=update 
body:{
    CorpCode:string     客户编码
    Name:string      姓名
    Contactid:string     联系人id 同时存放企业微信联系人id跟手动添加联系人id *
    Gender:string      性别 0 未知 1男 2女
    Logo:string      联系人头像
    Department:string     部门
    Postion:string      职务
    Phone1:string      手机号
    Phone2:string      手机号
    Email:string      邮箱
    QQ:string      QQ号
    Wechat:string      微信号
    Iswx:string      默认为企业微信联系人
    Comments:string      备注
    Address:string      地址
    ParentDepartment:string      上级部门
    BirthDay:string      生日
    Role :string      角色 //联系人角色
    Dndisturb :string      免打扰;default:false
    Enabled  :string      默认是否有效;default:true
  }
res: {result: "false|true"}
```



### 删除某个联系人信息
```
req: GET https://tcrm-iti.dm-ai.com/api/contact?action=delete&contactid=XXX
res: {result: "false|true"}
```



### 企查查模糊查询
```
req: GET https://tcrm-iti.dm-ai.com/api/qcc?action=fuzzy&keyword=XXX
res: {result: [{Keyno:string,Name:string 客户名称,Creditcode:string 信用代码,Opername:string 法人,Status:string 状态,No:string 编号},]	
```




### 查看产品列表
```
req: GET  https://tcrm-iti.dm-ai.com/api/product?action=list
res: {
    Products: [{
    Name: string       产品名称
    Comments: string       备注
    Createby: string       创建人
    Createbyname: string   创建人名字
    Model  
  },]  
}
```




### 增加产品信息
```
req: POST https://tcrm-iti.dm-ai.com/api/product?action=add
  body:{
    Name :string       产品名称*
    Comments : string  备注
  }
res: {result: "false|true"}
```



### 更新产品信息
```
req: POST https://tcrm-iti.dm-ai.com/api/product?action=update&id=XXX
body:{
    Name :string       产品名称*
    Comments : string  备注
  }
res: {result: "false|true"}
```



### 删除产品信息
```
req: GET https://tcrm-iti.dm-ai.com/api/product?action=delete&id=XXX
res: {result: "false|true"}	
```


### 查看合同信息
```
req: GET https://tcrm-iti.dm-ai.com/api/agreement?action=info&id=XXX
res: {
    CorpCode :string    客户编码
    Name : string    合同名称
    Signer:string    签订人
    Projectname:string    项目
    Money:string    金额
    Startdate:string    开始日期
    Enddate:string    截止日期
    Comments:string   备注及说明
    Createby:string   创建人
    CreatebyName:string    创建人
    Updateby:string       更新人
    UpdatebyName:string    更新人名字
    Model
}
```


### 增加合同信息
```
req: POST https://tcrm-iti.dm-ai.com/api/agreement?action=add
body:{
    CorpCode :string    客户编码*
    Name : string    合同名称*
    Signer:string    签订人*
    Projectname:string    项目*
    Money:string    金额
    Startdate:string    开始日期
    Enddate:string    截止日期
    Comments:string   备注及说明
  }
res: {result: "false|true"}
```



### 更新合同信息
```
req: POST https://tcrm-iti.dm-ai.com/api/agreement?action=update&id=XXX
body:{
    CorpCode :string    客户编码
    Name : string    合同名称
    Signer:string    签订人
    Projectname:string    项目
    Money:string    金额
    Startdate:string    开始日期
    Enddate:string    截止日期
    Comments:string   备注及说明
  }
res: {result: "false|true"}
```



### 删除合同信息
```
req: GET https://tcrm-iti.dm-ai.com/api/agreement?action=delete&id=XXX
res: {result: "false|true"}
```


### 查看项目列表
```
req: GET https://tcrm-iti.dm-ai.com/api/project?corpcode=XXX
res: {Projects:[{
    CorpCode:string       客户编码
    Name :string       项目名称
    Productid:string       产品id
    Productname:string     产品名称
    Timeplan:string       时间计划
    Budget:string       预计金额（元)
    Presaler:string      售前负责人
    Saler :string       销售负责人
    Comments :string     备注
    Createby :string      创建人
    CreatebyName :string      创建人名字
    Updateby :string       更新人
    UpdatebyName :string      更新人名字
    Model   
},]
}
```

### 查看项目信息
```
req: GET https://tcrm-iti.dm-ai.com/api/project?action=info&id=XXX
res: {
    CorpCode:string       客户编码
    Name :string       项目名称
    Productid:string       产品id
    Productname:string     产品名称
    Timeplan:string       时间计划
    Budget:string       预计金额（元)
    Presaler:string      售前负责人
    Saler :string       销售负责人
    Comments :string     备注
    Createby :string      创建人
    CreatebyName :string      创建人名字
    Updateby :string       更新人
    UpdatebyName :string      更新人名字
    Model   
}
```


### 增加项目信息
```
req: POST https://tcrm-iti.dm-ai.com/api/project?action=add
body:{
    CorpCode:string       客户编码*
    Name :string       项目名称*
    Productid:string       产品id*
    Productname:string     产品名称*
    Timeplan:string       时间计划*
    Budget:string       预计金额（元)*
    Presaler:string      售前负责人*
    Saler :string       销售负责人*
    Comments :string     备注
  }
res: {result: "false|true"}
```



### 更新项目信息
```
req: POST https://tcrm-iti.dm-ai.com/api/project?action=update&id=XXX
body:{
    CorpCode:string       客户编码
    Name :string       项目名称
    Productid:string       产品id
    Productname:string     产品名称
    Timeplan:string       时间计划
    Budget:string       预计金额（元) 
    Presaler:string      售前负责人
    Saler :string       销售负责人
    Comments :string     备注
  }
res: {result: "false|true"}

```


### 删除项目信息
```
req: GET https://tcrm-iti.dm-ai.com/api/project?action=delete&id=XXX
res:  {result: "false|true"}
```



### 查看跟进记录信息
```
req: GET https://tcrm-iti.dm-ai.com/api/record?action=view&corpcode=XXX
res: {
   Records:[{
    Name  :string    跟进记录名
    CorpCode :string    客户编码
    Projectid :string    产品线id
    Projectname :string   项目名
    Createby :string     员工ID
    CreatebyName: string   员工名字
    Updateby  :string    更新员工ID
    UpdatebyName :string    更新员工名字
    Followlog  :string    备注
    Model
},]
}
```


### 增加跟进记录信息
```
req: POST https://tcrm-iti.dm-ai.com/api/record?action=add
body:{
    Name  :string    跟进记录名*
    CorpCode :string    客户编码*
    Projectid :string    产品线id*
    Projectname :string   项目名*
    Followlog  :string    备注*
  }
res: {result: "false|true"}
```



### 更新跟进记录信息
```
req: POST https://tcrm-iti.dm-ai.com/api/record?action=update&id=XXX
body:{
    Name  :string    跟进记录名
    CorpCode :string    客户编码
    Projectid :string    产品线id
    Projectname :string   项目名
    Followlog  :string    备注
  }
res: {result: "false|true"}
```



### 删除跟进记录信息
```
req: GET https://tcrm-iti.dm-ai.com/api/record?action=delete&id=XXX
res: {result: "false|true"}
```


### 获取当前用户企业微信联系人
```
req: GET https://tcrm-iti.dm-ai.com/api/wx
res: {
  WXExusers:[{
  	Name: string
  	EXUserid: string
  },]
}
```

### 获取 Jsticket
```
req: GET https://tcrm-iti.dm-ai.com/api/jsticket
res:  {"jsticket": jsticket, "status": "200"}    
请求失败示例： {"jsticket":"","status":"400","ErrMsg": err}   
```



### 查看客户 跟进人 
```
req: GET https://tcrm-iti.dm-ai.com/api/principal?corpcode=XXX
res: {
    CustomerUserprincipal:[{
    CorpCode :string    客户编码
    Merchandiserid : string    专管业务员编码(为dmaiuser userid)
    MerchandiserName : string    专管业务员名字
    Logo : string    头像链接
    DepartmentCode :string    专管部门编码
    Default :string    是否默认负责人;default:false
    Createby : string    创建人编码
},]
}
```


### 添加客户 跟进人 
```
req: POST https://tcrm-iti.dm-ai.com/api/principal?action=add
body:{
    CorpCode :string    客户编码*
    Merchandiserid : string    专管业务员编码 *
    MerchandiserName : string    专管业务员名字
    Logo : string    头像链接
    DepartmentCode : string    专管部门编码
  }
res: {result: "false|true"}
```



###  更新客户 跟进人 
```
req: POST https://tcrm-iti.dm-ai.com/api/principal?action=update&id=XXX
body:{
    CorpCode :string    客户编码*
    Merchandiserid : string    专管业务员编码 *
    MerchandiserName : string  专管业务员名字
    Logo : string    头像链接
    DepartmentCode : string    专管部门编码
  }
res: {result: "false|true"}
```



### 删除客户 跟进人 
```
req: GET https://tcrm-iti.dm-ai.com/api/principal?action=delete&id=XXX
res: {result: "false|true"}
```





### 查看用户操作记录
```
req: GET https://tcrm-iti.dm-ai.com/api/operation&subject="customer|contact"
res: {
   Fman :string     动作发起人
   Action :string   增删改 delete,add,update
   Subject:string   对象主体
   Detail:string    详细信息
   Model
}
```



### 获取销售报表下载链接
```
req: GET https://tcrm-iti.dm-ai.com/api/report
res: {
  url: string      "excelfile download url"
}
```


### 获取客户级别
```
req: GET https://tcrm-iti.dm-ai.com/api/customer?type=level
res: {
    CustomerLevel: [{"Name":string,"Orgid":string},]  
}
```


### 同步客户信息到用友
```
req: POST https://tcrm-iti.dm-ai.com/api/customer?type=sync
body:{
    CorpCodes: ["customer1","customer2",...]       //customer1为corpcode string
}
res: {result: "false|true"}
```

### check 校验
```
req: GET https://tcrm-iti.dm-ai.com/api/check?check=XXXX&name=XXX&id=XXX
res: {
}
```

### xxxx
![dxxx](xxx "demo")  





    if action == "update"{
        h.Logger.Debugln("users跟进人: ",users)
        if len(users) == 0 && dmaiuser.Manager =="true"{
            if err := h.DB.TruncateCUP(corpcode);err == nil{

                h.Redis.DelKey("manager")
                h.UserInfoSyncToRD <- "manager"
                h.Redis.DelKey(cookie)
                h.UserInfoSyncToRD <- cookie

                c.JSON(http.StatusOK,gin.H{"status": 200})
                c.AbortWithStatus(http.StatusOK)
            }else{
                h.Logger.Errorln(err)
            }
        }else{
        cups := h.DB.ViewCustomerUserprincipal(corpcode)
        h.Logger.Debugln("现有跟进人： ",cups)
        tmp := []string{}
        if len(cups) != 0{ 
            for _,cup := range cups{
                flag := true
                for _,v1 := range users{
                    if v1 == cup.Merchandiserid {
                        flag = false 
                        continue
                    }
                }
                if flag{
                    tmp = append(tmp,cup.Merchandiserid)
                }
            }
            for _,v2 := range tmp{
                dmaiuser := h.DB.ViewUser(v2)
                if err := h.DB.DelCustomerUserprincipal(v2,corpcode);err != nil{
                    h.Logger.Errorln(err)
                }else{
                    h.AddOperation(cookie,corpcode,"删除","跟进人",dmaiuser.Name)
                    h.Logger.Debugln("delete user customer principal: ",v2)
                }
            }
        }
            for _,v := range users{
                if h.DB.CheckCUP(v,corpcode){
                    continue
                }else{
                    var info backend.CustomerUserprincipal
                    dmaiuser := h.DB.ViewUser(v)
                    info.CorpCode = corpcode
                    info.Merchandiserid  = v 
                    info.MerchandiserName = dmaiuser.Name
                    info.Createby = cookie 
                    if err := h.DB.AddCustomerUserprincipal(&info);err !=nil{
                        h.Logger.Errorln(err)
                    }else{

                        h.Redis.DelKey("manager")
                        h.UserInfoSyncToRD <- "manager"
                        h.Redis.DelKey(cookie)
                        h.UserInfoSyncToRD <- cookie
            
                        h.AddOperation(cookie,corpcode,"增加","跟进人",dmaiuser.Name)
                        h.Logger.Debugln("add a new customer user principal: ",v)
                    }
                }
            }

            c.JSON(http.StatusOK,gin.H{"status": 200})
            c.AbortWithStatus(http.StatusOK)
        }
    }




























































