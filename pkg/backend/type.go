package backend
import(
    "gorm.io/gorm"
)
const (
    DBError = 3000
    RedisError = 6000
    WXError = 9000
    YSError = 5000
    QccError = 7000
)

type DmaiUser struct{
    Name           string           `gorm:"not null;size:255;comment:姓名"`
    Gender         string           `gorm:"type:enum('0','1','2');comment:"性别 0 未知 1男 2女"`
    OpenUserid     string           `gorm:"size:50;comment:微信id"`
    Userid         string           `gorm:"size:50;comment:企业微信用户Id "`
    Position       string           `gorm:"size:50;comment:职位"`
    Department     string           `gorm:"size:50;comment:部门"` 
    Mobile         string           `gorm:"size:50;comment:手机号"`
    Email          string           `gorm:"size:50;comment:邮箱"` 
    Alias          string           `gorm:"size:50;comment:别称"`
    Logo           string           `gorm:"size:255;comment:头像链接"`
    Address         string          `gorm:"size:255;comment:地址"`
    Manager        string           `gorm:"type:enum('false','true');comment:管理员;default:false"`
    Enabled         string          `gorm:"type:enum('false','true');comment:员工状态;default:true"` 
    gorm.Model
}

//客户公司名称
type Customer struct{
    CorpCode                string        `gorm:"size:255;unique;comment:客户编码"`
    CorpName                string        `gorm:"size:255;comment:客户名称"`
    Opername                string        `gorm:"size:50;comment:法人名称"`
    Startdate               string        `gorm:"size:50;comment:营业起始时间"`
    Province                string        `gorm:"size:50;comment:省"`
    Address                 string        `gorm:"size:255;comment:详细地址"`
    Status                  string        `gorm:"size:128;comment:在营状态"`
    Registcapi              string        `gorm:"size:255;comment:注册资金"`
    Creditcode              string        `gorm:"size:128;comment:信用代码"`
    Scope                   string        `gorm:"size:255;comment:经营范围"`
    Econkind                string        `gorm:"size:255;comment:公司类型"`
    Belongorg               string        `gorm:"size:255;comment:归属管理组织"`
    Orgno                   string        `gorm:"size:255;comment:组织代码"`
    Ordernumber             string        `gorm:"size:255;comment:工商快照"`
    Logo                    string        `gorm:"size:255;comment:logo地址"`

    Customerlevel           string        `gorm:"size:255;comment:客户级别"`
    CustomerlevelId         int64         `gorm:"size:255;comment:客户级别ID"`
    OwnerOrgid              string        `gorm:"size:255;comment:客户归属的分公司ID"`
    OwnerOrgname            string        `gorm:"size:255;comment:客户归属的分公司名称"`
    Ysid                    int64         `gorm:"size:255;comment:用友对应的客户id"` 
    Ysapplyid               int64         `gorm:"size:255;comment:用友对应的客户applyrangeid"` 
    IntentionScore          string        `gorm:"size:4;default:50;comment:"客户意向评分"`
    QualificationScore      string        `gorm:"size:4;default:52;comment:"客户资质评分"`
    Trade                   string        `gorm:"size:255;comment:客户行业"`   
    Region                  string        `gorm:"size:255;comment:区域"`   
    BaseReporturl           string        `gorm:"size:255;comment:企查查基础信用报告地址"`
    // CompetingGoodsInfo      string        `gorm:"size:255;comment:竞品情报"`
    BusinessRequirements    string        `gorm:"size:255;comment:业务需求"`
    Comments                string        `gorm:"size:255;comment:备注"`
    Staffs                  int           `gorm:"comment:员工人数"`
    // Telephone               string        `gorm:"size:50;comment:固定电话"`
    // SuperiorCustomer        string        `gorm:"size:255;comment:上级客户"`
    // Dndisturb               string        `gorm:"type:enum('false','true');comment:免打扰;default:false"`
    gorm.Model
}

type CustomerShiXin struct{
    CorpCode                string        `gorm:"size:255;comment:客户编码"`
    LianDate                string        `gorm:"size:50;comment:立案时间"`
    Anno                    string        `gorm:"size:50;comment:案号"`
    Executegov              string        `gorm:"size:50;comment:执行法院"`
    Executestatus           string        `gorm:"size:255;comment:被执行人的履行情况"`
    Publicdate              string        `gorm:"size:50;comment:发布时间"`
    Executeno               string        `gorm:"size:50;comment:执行依据文号"`   
    gorm.Model
}

type CustomerSeriousIllegal  struct{
    CorpCode                string        `gorm:"size:255;comment:客户编码"`
    Type                    string        `gorm:"size:50;comment:类型"`
    AddReason               string        `gorm:"size:255;comment:列入原因"`
    AddDate                 string        `gorm:"size:50;comment:列入日期"`
    AddOffice               string        `gorm:"size:50;comment:列入决定机关"`
    RemoveReason            string        `gorm:"size:255;comment:移除原因"`
    RemoveDate              string        `gorm:"size:50;comment:移除日期"`
    RemoveOffice            string        `gorm:"size:50;comment:移除决定机关"`
    gorm.Model
}

type CustomerTaxIllegal struct{
    CorpCode                string        `gorm:"size:255;comment:客户编码"`
    PublishDate             string        `gorm:"size:50;comment:发布日期"`
    CaseNature              string        `gorm:"size:255;comment:案件性质"`
    TaxGov                  string        `gorm:"size:50;comment:所属税务机关"`
    gorm.Model
}


type CustomerAdminPenalty struct{
    CorpCode                string        `gorm:"size:255;comment:客户编码"`
    DocNo                   string        `gorm:"size:50;comment:处罚决定文书"`
    PunishReason            string        `gorm:"size:255;comment:处罚事由"`
    PunishResult            string        `gorm:"size:1024;comment:处罚结果"`
    PunishOffice            string        `gorm:"size:50;comment:处罚单位"`
    PunishDate              string        `gorm:"size:50;comment:处罚日期"`
    Source                  string        `gorm:"size:50;comment:数据来源"`
    SourceCode              string        `gorm:"size:50;comment:数据来源Code，1,2-市场监督管理局，3-税务局，4-其他"` 
    gorm.Model
}



// type Product struct{
//     CorpCode                string       `gorm:"size:255;comment:客户编码"`
//     ProductName             string       `gorm:"size:255;comment:产品名称"`
//     BusinessType            string       `gorm:"size:50;comment:业务类型"`
//     NormalPrice             string       `gorm:"type:decimal(65,30);comment:常规价格"`
//     SellingUnit             string       `gorm:"type:enum('套','件','台','个','件');default:套;comment:销售单位"`
//     ProductImageUrl         string       `gorm:"size:255;comment:产品图片URL地址"`
//     ProductDescription      string       `gorm:"size:1024;comment:产品描述"`
//     Enabled                 string       `gorm:"type:enum('false','true');comment:产品状态;default:true"`
//     ProductOwner            string       `gorm:"size:255;comment:产品所有人"`
//     CreateBy                string       `gorm:"size:255;comment:创建者"`
//     ProductDir              string       `gorm:"size:255;comment:产品目录"`
//     gorm.Model
// }

//联系人
type Contactor struct{
    CorpCode               string      `gorm:"size:255;comment:客户编码"`
    Name                   string      `gorm:"size:100;comment:姓名" json:"text"`
    ContactCode           string      `gorm:"size:100;comment:联系人code"`
    Wechatid              string      `gorm:"size:100;comment:联系人id 同时存放企业微信联系人id"`
    Gender                 string      `gorm:"type:enum('0','1','2');default:0;comment:"性别 0 未知 1男 2女"`
    Logo                   string      `gorm:"size:255;comment:联系人头像"`
    // Department             string      `gorm:"size:50;comment:部门"`
    Postion                string      `gorm:"size:100;comment:职务"`
    Phone                 string      `gorm:"size:50;comment:手机号"`
    Phone2                 string      `gorm:"size:50;comment:手机号"`
    Email                  string      `gorm:"size:50;comment:邮箱"`
    Iswx                   string      `gorm:"type=enum('1','0');default:0;comment:默认不是企业微信联系人"`
    // Comments               string      `gorm:"size:255;comment:备注"`
    // Address                string      `gorm:"size:255;comment:地址"`
    // ParentDepartment       string      `gorm:"size:50;comment:上级部门"`
    // BirthDay               string      `gorm:"size:50;comment:生日"`
    // Role                   string      `gorm:"size:50;comment:角色"` //联系人角色
    Dndisturb              string      `gorm:"type:enum('false','true');comment:免打扰;default:false"`
    Enabled                string      `gorm:"type:enum('false','true');comment:默认是否有效;default:true"` 
    gorm.Model            
}


//CustomerStaff 客户销售对应
type CustomerUserprincipal struct{
    CorpCode                string    `gorm:"size:255;comment:客户编码"`
    Merchandiserid          string    `gorm:"size:100;comment:专管业务员编码"`
    MerchandiserName        string    `gorm:"size:100;comment:专管业务员名字"`
    Logo                    string    `gorm:"size:255;comment:头像链接"`
    DepartmentCode          string    `gorm:"size:50;comment:专管部门编码"`
    Presaler                string    `gorm:"type:enum('false','true');comment:是否售前;default:true"`   
    Saler                   string    `gorm:"type:enum('false','true');comment:是否销;default:false"`  
    Createby                string    `gorm:"size:50;comment:创建人编码"`
    // CustomerLevel           string    `gorm:"size:50;comment:客户级别"`
    gorm.Model 
}





//联系人 销售对应
type ContactUserprincipal struct{
    ContactCode          string   `gorm:"size:255;comment:联系人Id"`//联系人
    Ownerid            string   `gorm:"size:255;comment:联系人归属人userid"`// 联系人归属名字
    Createid           string   `gorm:"size:255;comment:联系人创建者userid"`//创建联系人名字
    gorm.Model 
}

//产品线
type Product struct{
    Name              string       `gorm:"size:50;comment:产品名称"`
    Comments          string       `gorm:"size:1024;comment:备注"`
    Createby          string       `gorm:"size:50;comment:创建人"` 
    Createbyname      string       `gorm:"size:50;comment:创建人名字"`
    // Imgurl            string       `gorm:"size:1024;comment: 头像图片地址"`
    gorm.Model    
}

//项目
type Project struct{
    CorpCode          string       `gorm:"size:255;comment:客户编码"`
    Name              string       `gorm:"size:50;comment:项目名称"`
    Productid         string       `gorm:"size:50;comment:产品id"`
    Productname       string       `gorm:"size:50;comment:产品名称"`
    Timeplan          string       `gorm:"size:50;comment:时间计划"`
    Budget            string       `gorm:"size:255;comment:预计金额（元）"`
    Presaler          string       `gorm:"size:50;comment:售前负责人"`
    Saler             string       `gorm:"size:50;comment:销售负责人"`
    Comments          string       `gorm:"size:1024;comment:备注"`
    YSpid             string       `gorm:"size:255;comment:用友projectid"`
    YSpname           string       `gorm:"size:255;comment:用友projectname"`

    Createby          string       `gorm:"size:255;comment:创建人"`
    CreatebyName      string       `gorm:"size:255;comment:创建人名字"`
    Updateby          string       `gorm:"size:255;comment:更新人"`
    UpdatebyName      string       `gorm:"size:255;comment:更新人名字"`
    // Imgurl            string       `gorm:"size:1024;comment: 头像图片地址"`
    gorm.Model    
}


//合同
type Agreement struct{
    CorpCode    string    `gorm:"size:255;comment:客户编码"`
    Name        string    `gorm:"size:255;comment:合同名称"`
    Signer      string    `gorm:"size:50;comment:签订人"`
    Projectname string    `gorm:"size:255;comment:项目"`
    Money       string    `gorm:"size:255;comment:金额"`
    Startdate    string    `gorm:"size:255;comment:开始日期"`
    Enddate     string    `gorm:"size:255;comment:截止日期"`
    Comments    string    `gorm:"size:255;comment:备注及说明"`

    Createby     string    `gorm:"size:255;comment:创建人"`
    CreatebyName     string    `gorm:"size:255;comment:创建人"`
    Updateby          string       `gorm:"size:255;comment:更新人"`
    UpdatebyName      string       `gorm:"size:255;comment:更新人名字"`
    gorm.Model 
}


//客户行业
type Trade struct{
    Name              string       `gorm:"size:50;comment:行业名称"`
    gorm.Model    
}

//跟进记录
type FollowRecord struct{
    Name              string    `gorm:"size:255;comment:跟进记录名" json:"text"`
    CorpCode          string    `gorm:"size:255;comment:客户编码"`
    Projectid         string    `gorm:"size:255;comment:产品线id"`
    Projectname       string    `gorm:"size:255;comment:项目名"`
    Timeplan          string    `gorm:"size:50;comment:时间计划"`
    Budget            string    `gorm:"size:255;comment:预计金额（元）"`
    Createby          string    `gorm:"size:255;comment:员工ID"` //销售人名字
    CreatebyName      string    `gorm:"size:255;comment:员工名字"` //销售人名字
    Updateby          string    `gorm:"size:255;comment:更新员工ID"` //销售人名字
    UpdatebyName      string    `gorm:"size:255;comment:更新员工名字"` //销售人名字
    Followlog         string    `gorm:"size:1024;comment:备注"`
    gorm.Model
}

//企业微信 客户联系人变更监控
type Wxlog struct {
    ToUsername   string `xml:"ToUserName";gorm:"size:255;comment:操作对象"`
    FromUsername string `xml:"FromUserName";gorm:"size:255;comment:操作者"`
    CreateTime   uint32 `xml:"CreateTime";gorm:"size:255;comment:创建时间"`
    MsgType      string `xml:"MsgType";gorm:"size:255;comment:信息类型"`
    Event        string `xml:"Event";gorm:"size:255;comment:什么事件"`
    ChangeType   string `xml:"ChangeType";gorm:"size:255;comment:改变类型id"`
    UserID       string `xml:"UserID";gorm:"size:255;comment:用户id"`
    Contact      string `xml:"Contact";gorm:"size:255;comment:联系人id"`
    ExternalUserID   string `xml:"ExternalUserID";gorm:"size:255;comment:外部联系人id"`
   State        string `xml:"State";gorm:"size:255;comment:交接状态"`
   WelcomeCode  string `xml:"WelcomeCode";gorm:"size:255;comment:欢迎码"`
//客户联系人交接   
    FailReason   string `xml:"FailReason";gorm:"size:255;comment:客户联系人交接失败原因"`
    ChatId       string `xml:"ChatId";gorm:"size:255;comment:群聊id"`
    UpdateDetail string `xml:"UpdateDetail";gorm:"size:255;comment:改变详情"`
    JoinScene    string `xml:"JoinScene";gorm:"size:255;comment:加入"`
    QuitScene    string `xml:"QuitScene";gorm:"size:255;comment:推出"`
    MemChangeCnt string `xml:"MemChangeCnt";gorm:"size:255;comment:成员改变多少"`
    gorm.Model
}

//操作日志
type Operation struct{
   Fman       string   `gorm:"size:50;comment:动作发起人id"`
   Fmanname   string   `gorm:"size:50;comment:动作发起人名字"`
   Corpcode   string   `gorm:"size:50;comment:客户corpcode"`
   Operation  string   `gorm:"size:50;comment:增删改"`
   Subject    string   `gorm:"size:50;comment:对象主体"`
   Object     string   `gorm:"size:50;comment:具体对象"`
   gorm.Model
}

//客户档案保存执行结果
type YsCustomer struct{
   Name       string   `gorm:"size:255;comment:用友客户档案名称"`
   Cid         int64    `gorm:"size:255;comment:用友客户档案id"`
   CreateOrg   string   `gorm:"size:255;comment:用友客户档案归属业务部门id"`
   Code        string   `gorm:"size:255;comment:客户档案code"`
   ChannCustomerClass  int64   `gorm:"size:255;comment:渠道客户分类id"`
   InvoicingCustomers  int64    `gorm:"size:255;comment:开票客户id"`
   MerchantAppliedDetail   int64     `gorm:"size:255;comment:客户档案code"`
   DetailMerchantApplyRangeId  int64   `gorm:"size:255;comment:业务信息merchantAppliedDetail applyrange id"`
   MerchantRole        int64         `gorm:"size:255;comment:MerchantRole id"`
   CustomerDefine   int64       `gorm:"size:255;comment:CustomerDefine id"`
   MerchantApplyRangeId  int64     `gorm:"size:255;comment:erchantApplyRangeId id"`
   gorm.Model
}

