package service


type YSToken struct {
        Code    string `json:"code"`
        Message string `json:"message"`
        Data    struct {
                AccessToken string `json:"access_token"`
                Expire      int    `json:"expire"`
        } `json:"data"`
}


type Merchantrole struct {
	Businessrole       string `json:"businessRole,omitempty"`
	Tobimmigrationmode string `json:"toBImmigrationMode,omitempty"`
	Settlementmethod   string `json:"settlementMethod,omitempty"`
	Cardtype           string `json:"cardType,omitempty"`
	Status             string `json:"_status,omitempty"`
} 

type Merchantapplieddetail  struct {
	Searchcode               string `json:"searchcode,omitempty"`
	Merchantapplyrangeid     int64 `json:"merchantApplyRangeId,omitempty`
	Merchantid               string `json:"merchantId,omitempty"`
	Customerlevel            int64  `json:"customerLevel,omitempty"`
	CustomerlevelCode        string `json:"customerLevel_code,omitempty"`
	Specialmanagementdep     string `json:"specialManagementDep,omitempty"`
	CustomerlevelName        string `json:"customerLevel_Name,omitempty"`
	SpecialmanagementdepName string `json:"specialManagementDep_Name,omitempty"`
	SpecialmanagementdepCode string `json:"specialManagementDep_code,omitempty"`
	ProfessSalesman          string `json:"professSalesman,omitempty"`
	ProfesssalesmanName      string `json:"professSalesman_Name,omitempty"`
	Deliverywarehouse        int64  `json:"deliveryWarehouse,omitempty"`
	DeliverywarehouseName    string `json:"deliveryWarehouse_Name,omitempty"`
	Transactioncurrency      string `json:"transactionCurrency,omitempty"`
	Exchangeratetype         string `json:"exchangeratetype,omitempty"`
	TransactioncurrencyName  string `json:"transactionCurrency_Name,omitempty"`
	ExchangeratetypeName     string `json:"exchangeratetype_Name,omitempty"`
	Taxrate                  string `json:"taxRate,omitempty"`
	TaxrateName              float64    `json:"taxRate_Name,omitempty"`
	Payway                   int    `json:"payway,omitempty"`
	Creditserviceday         int    `json:"creditServiceDay,omitempty"`
	Settlementmethod         int64  `json:"settlementMethod,omitempty"`
	SettlementmethodName     string `json:"settlementMethod_Name,omitempty"`
	Shipmentmethod           int64  `json:"shipmentMethod,omitempty"`
	ShipmentmethodName       string `json:"shipmentMethod_Name,omitempty"`
	Signback                 bool   `json:"signBack,omitempty"`
	Stopstatus               bool   `json:"stopstatus,omitempty"`
	Status                   string `json:"_status"`
} 

type Merchantsmanager struct {
	Username    string `json:"userName,omitempty"`
	Countrycode string `json:"countryCode,omitempty"`
	Mobile      string `json:"mobile,omitempty"`
	Fullname    string `json:"fullName,omitempty"`
	Email       string `json:"email,omitempty"`
	Qq          string `json:"qq,omitempty"`
	Expiredate  string `json:"expireDate,omitempty"`
	Status      string `json:"_status,omitempty"`
}

type Customerdefine struct {
	Customerdefine22 string `json:"customerDefine22,omitempty"`
	Customerdefine23 string `json:"customerDefine23,omitempty"`
	Customerdefine24 string `json:"customerDefine24,omitempty"`
	Status           string `json:"_status"`
}

type Merchantattachment struct {
	Folder   string `json:"folder,omitempty"`
	Type     string `json:"type,omitempty"`
	Size     int    `json:"size,omitempty"`
	Filename string `json:"fileName,omitempty"`
	Status   string `json:"_status"`
} 

type Merchantcorpimage  struct {
	Folder  string `json:"folder,omitempty"`
	Type    string `json:"type,omitempty"`
	Size    int    `json:"size,omitempty"`
	Imgname string `json:"imgName,omitempty"`
	Sort    int    `json:"sort,omitempty"`
	Status  string `json:"_status"`
} 

type Merchantaddressinfo  struct {
	Isdefault      string `json:"isDefault,omitempty"`
	Hasdefaultinit bool   `json:"hasDefaultInit,omitempty"`
	Addresscode    string `json:"addressCode,omitempty"`
	Receiver       string `json:"receiver,omitempty"`
	Mobile         string `json:"mobile,omitempty"`
	Address        string `json:"address,omitempty"`
	Zipcode        string `json:"zipCode,omitempty"`
	Regioncode     string `json:"regionCode,omitempty"`
	Status         string `json:"_status"`
}

type  FullName struct {
	ZhCn  string `json:"zh_CN,omitempty"`

}
type Merchantcontacterinfo struct {
	Isdefault      string `json:"isDefault"`
	Hasdefaultinit bool   `json:"hasDefaultInit"`
	PositionName   string  `json:"positionName,omitempty"`
	Fullname       FullName `json:"fullName,omitempty"`
	Mobile string `json:"mobile,omitempty"`
	Email  string `json:"email,omitempty"`
	Wechat string `json:"weChat,omitempty"`
	Status string `json:"_status"`
	ID     string `json:"id,omitempty"`
	Pubts  string `json:"pubts,omitempty"`
	Merchantid string `json:"merchantId,omitempty"`
	Gender string `json:"gender,omitempty"`
	Yhtuserid string `json:"yhtUserId,omitempty"`
	Cuscontact string `json:"cusContact,omitempty"`
	Ordercontact string `json:"orderContact,omitempty"`
	Mallcontact  string `json:"mallContact,omitempty"`
	TableDisplayOutlineAll  bool `json:"_tableDisplayOutlineAll,omitempty"`
	Remarks      string   `json:"remarks,,omitempty"`
}

type Merchantagentfinancialinfo  struct {
	Accounttype     string `json:"accountType,omitempty"`
	Isdefault       string `json:"isDefault,omitempty"`
	Stopstatus      string `json:"stopstatus,omitempty"`
	Hasdefaultinit  bool   `json:"hasDefaultInit,omitempty"`
	CurrencyName    string `json:"currency_name,omitempty"`
	Currency        string `json:"currency,omitempty"`
	BankName        string `json:"bank_name,omitempty"`
	Bank            string `json:"bank,omitempty"`
	OpenbankName    string `json:"openBank_name,omitempty"`
	Openbank        string `json:"openBank,omitempty"`
	Jointlineno     string `json:"jointLineNo,omitempty"`
	Bankaccount     string `json:"bankAccount,omitempty"`
	Bankaccountname string `json:"bankAccountName,omitempty"`
	CountryName     string `json:"country_name,omitempty"`
	Country         string `json:"country,omitempty"`
	ProvinceName    string `json:"province_Name,omitempty"`
	Province        int64  `json:"province,omitempty"`
	Status          string `json:"_status,omitempty"`
}


type Merchantagentinvoiceinfo struct {
	Isdefault            string `json:"isDefault,omitempty"`
	Hasdefaultinit       bool   `json:"hasDefaultInit,omitempty"`
	Billingtype          string `json:"billingType,omitempty"`
	Title                string `json:"title,omitempty"`
	Taxno                string `json:"taxNo,omitempty"`
	Receievinvoicemobile string `json:"receievInvoiceMobile,omitempty"`
	Status               string `json:"_status,omitempty"`
} 

type Merchantapplyrange  struct {
	Rangetype      int    `json:"rangeType"`
	Iscreator      string `json:"isCreator"`
	Isapplied      string `json:"isApplied"`
	Hasdefaultinit bool   `json:"hasDefaultInit"`
	Orgid          string `json:"orgId"`
	Tenant         int64  `json:"tenant,omitempty"`
	Orgcode        string `json:"orgCode"`
	Orgname        string `json:"orgName"`
	Status         string `json:"_status"`
	Id             string `json:"_id,omitempty"`
}
type Name    struct {
	ZhCn string `json:"zh_CN"`
} 

type Shortname struct {
	ZhCn string `json:"zh_CN"`
} 
//设定专管业务员
type Principal struct {
	SpecialManagementDepCode string `json:"specialManagementDep_code,omitempty"`
	IsDefault                bool   `json:"isDefault"`
	HasDefaultInit           bool   `json:"hasDefaultInit"`
	ID                       string `json:"_id,omitempty"`
	ProfessSalesmanName      string `json:"professSalesman_Name"`
	ProfessSalesman          string `json:"professSalesman"`
	SpecialManagementDep     string `json:"specialManagementDep,omitempty"`
	SpecialManagementDepName string `json:"specialManagementDep_Name,omitempty"`
	Status                   string `json:"_status"`
}

type MerchantData struct {
	Data struct {
		Createorg     string `json:"createOrg"`
		CreateorgName string `json:"createOrg_name"`
		CreateorgCode string `json:"createOrg_code"`
		Code          string `json:"code"`
		Name          Name  `json:"name"`
		Shortname     Shortname `json:"shortname"`
		Principals    []Principal `json:"principals"`
		Country                string      `json:"country"`
		Belongorg                string      `json:"belongOrg"`
		Timezone               string      `json:"timeZone,omitempty"`
		Language               string      `json:"language,omitempty"`
		CountryName            string      `json:"country_name,omitempty"`
		TimezoneName           string      `json:"timeZone_Name,omitempty"`
		LanguageName           string      `json:"language_Name,omitempty"`
		CountryCode            string      `json:"country_code,omitempty"`
		Invoicingcustomers     int64       `json:"invoicingCustomers,omitempty"`
		InvoicingcustomersName string      `json:"invoicingCustomers_Name,omitempty"`
		ParentcustomerName     string      `json:"parentCustomer_Name,omitempty"`
		Suppliers              int64       `json:"suppliers,omitempty"`
		SuppliersName          string      `json:"suppliers_Name,omitempty"`
		Retailinvestors        bool        `json:"retailInvestors,omitempty"`
		Internalorg            bool        `json:"internalOrg,omitempty"`
		Internalorgid          interface{} `json:"internalOrgId,omitempty"`
		InternalorgidName      string      `json:"internalOrgId_Name,omitempty"`
		Taxpayingcategories    int         `json:"taxPayingCategories,omitempty"`
		Customerclass          int64       `json:"customerClass,omitempty"`
		CustomerclassName      string      `json:"customerClass_Name,omitempty"`
		CustomerclassCode      string      `json:"customerClass_code,omitempty"`
		CustomerclassPath      string      `json:"customerClass_path,omitempty"`
		Customerarea           int64       `json:"customerArea,omitempty"`
		CustomerareaName       string      `json:"customerArea_Name,omitempty"`
		Customerindustry       int64       `json:"customerIndustry,omitempty"`
		CustomerindustryName   string      `json:"customerIndustry_Name,omitempty"`
		CustomerindustryCode   string      `json:"customerIndustry_code,omitempty"`
		Merchantapplieddetail  Merchantapplieddetail  `json:"merchantAppliedDetail,omitempty"`
		Merchantsmanager       Merchantsmanager `json:"merchantsManager,omitempty"`
		Merchantrole           Merchantrole `json:"merchantRole,omitempty"`
		Customerdefine         Customerdefine  `json:"customerDefine,omitempty"`
		Enterprisenature  string `json:"enterpriseNature"`
		Enterprisename    string `json:"enterpriseName"`
		Leadername        string `json:"leaderName"`
		Leadernameidno    string `json:"leaderNameIdNo"`
		Personname        string `json:"personName,omitempty"`
		Idno              string `json:"idNo,omitempty"`
		Orgname           string `json:"orgName"`
		Creditcode        string `json:"creditCode"`
		Personcharge      string `json:"personCharge,omitempty"`
		Personchargeidno  string `json:"personChargeIdNo,omitempty"`
		Frontidno         string `json:"frontIdNo,omitempty"`
		Backidno          string `json:"backIdNo,omitempty"`
		Businesslicense   string `json:"businessLicense,omitempty"`
		Orgidno           string `json:"orgIdNo,omitempty"`
		Businesslicenseno string `json:"businessLicenseNo,omitempty"`
		Regioncode        string `json:"regionCode,omitempty"`
		Address           struct {
			ZhCn string `json:"zh_CN"`
		} `json:"address"`
		Postcode     string `json:"postCode,omitempty"`
		Contactname  string `json:"contactName,omitempty"`
		Contacttel   string `json:"contactTel,omitempty"`
		Email        string `json:"email,omitempty"`
		Buildtime    string `json:"buildTime"`
		Currencycode string `json:"currencyCode,omitempty"`
		Money        int    `json:"money"`
		Peoplenum    int    `json:"peopleNum,omitempty"`
		Scopemodel   int    `json:"scopeModel,omitempty"`
		Yearmoney    int    `json:"yearMoney,omitempty"`
		Scope        struct {
			ZhCn string `json:"zh_CN"`
		} `json:"scope"`
		Website             string `json:"website,omitempty"`
		Wid                 string `json:"wid,omitempty"`
		Createtime          string `json:"createTime"`
		Verfymark           int    `json:"verfyMark"`
		Merchantattachments []Merchantattachment `json:"merchantAttachments,omitempty"`
		Merchantcorpimages []Merchantcorpimage   `json:"merchantCorpImages,omitempty"`
		Merchantaddressinfos []Merchantaddressinfo   `json:"merchantAddressInfos,omitempty"`
		Merchantcontacterinfos []Merchantcontacterinfo   `json:"merchantContacterInfos"`
		Merchantagentfinancialinfos []Merchantagentfinancialinfo  `json:"merchantAgentFinancialInfos,omitempty"`
		Merchantagentinvoiceinfos []Merchantagentinvoiceinfo   `json:"merchantAgentInvoiceInfos,omitempty"`
		Merchantapplyranges []Merchantapplyrange  `json:"merchantApplyRanges"`
		Status string `json:"_status"`
		// Source int    `json:"source"`
	} `json:"data"`
}

type MerchantResult struct {
	Code    string    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Createorg     string `json:"createOrg"`
		CreateorgName string `json:"createOrg_name"`
		CreateorgCode string `json:"createOrg_code"`
		Code          string `json:"code"`
		Name          struct {
			ZhCn string `json:"zh_CN"`
		} `json:"name"`
		Shortname struct {
			ZhCn string `json:"zh_CN"`
		} `json:"shortname"`
		Invoicingcustomers     int64  `json:"invoicingCustomers"`
		InvoicingcustomersName string `json:"invoicingCustomers_Name"`
		ParentcustomerName     string `json:"parentCustomer_Name"`
		Suppliers              int64  `json:"suppliers"`
		SuppliersName          string `json:"suppliers_Name"`
		Retailinvestors        bool   `json:"retailInvestors"`
		Internalorg            bool   `json:"internalOrg"`
		InternalorgidName      string `json:"internalOrgId_Name"`
		Taxpayingcategories    int    `json:"taxPayingCategories"`
		Customerclass          int64  `json:"customerClass"`
		CustomerclassName      string `json:"customerClass_Name"`
		CustomerclassCode      string `json:"customerClass_code"`
		Customerarea           int64  `json:"customerArea"`
		CustomerareaName       string `json:"customerArea_Name"`
		Customerindustry       int64  `json:"customerIndustry"`
		CustomerindustryName   string `json:"customerIndustry_Name"`
		CustomerindustryCode   string `json:"customerIndustry_code"`
		Merchantsmanager       struct {
			Countrycode string `json:"countryCode"`
			Mobile      string `json:"mobile"`
			Fullname    string `json:"fullName"`
			Email       string `json:"email"`
			Qq          string `json:"qq"`
			Expiredate  string `json:"expireDate"`
			ID          int64  `json:"id"`
			Merchantid  int64  `json:"merchantId"`
			Tenant      int64  `json:"tenant"`
		} `json:"merchantsManager"`
		Customerdefine struct {
			Customerdefine22 string `json:"customerDefine22"`
			Customerdefine23 string `json:"customerDefine23"`
			Customerdefine24 string `json:"customerDefine24"`
			ID               int64  `json:"id"`
			Merchantid       int64  `json:"merchantId"`
			Tenant           int64  `json:"tenant"`
		} `json:"customerDefine"`
		Enterprisenature  int    `json:"enterpriseNature"`
		Enterprisename    string `json:"enterpriseName"`
		Leadername        string `json:"leaderName"`
		Leadernameidno    string `json:"leaderNameIdNo"`
		Orgname           string `json:"orgName"`
		Creditcode        string `json:"creditCode"`
		Frontidno         string `json:"frontIdNo"`
		Businesslicense   string `json:"businessLicense"`
		Businesslicenseno string `json:"businessLicenseNo"`
		Regioncode        string `json:"regionCode"`
		Address           struct {
			ZhCn string `json:"zh_CN"`
		} `json:"address"`
		Postcode     string `json:"postCode"`
		Contactname  string `json:"contactName"`
		Contacttel   string `json:"contactTel"`
		Email        string `json:"email"`
		Buildtime    string `json:"buildTime"`
		Currencycode string `json:"currencyCode"`
		Money        string `json:"money"`
		Peoplenum    int    `json:"peopleNum"`
		Scopemodel   int    `json:"scopeModel"`
		Yearmoney    string `json:"yearMoney"`
		Scope        struct {
			ZhCn string `json:"zh_CN"`
		} `json:"scope"`
		Website             string `json:"website"`
		Wid                 string `json:"wid"`
		Createtime          string `json:"createTime"`
		Verfymark           int    `json:"verfyMark"`
		Merchantattachments []struct {
			Folder     string `json:"folder"`
			Type       string `json:"type"`
			Size       int    `json:"size"`
			Filename   string `json:"fileName"`
			ID         int64  `json:"id"`
			Merchantid int64  `json:"merchantId"`
			Tenant     int64  `json:"tenant"`
		} `json:"merchantAttachments"`
		Merchantcorpimages []struct {
			Folder     string `json:"folder"`
			Type       string `json:"type"`
			Size       int    `json:"size"`
			Imgname    string `json:"imgName"`
			Sort       int    `json:"sort"`
			ID         int64  `json:"id"`
			Merchantid int64  `json:"merchantId"`
			Tenant     int64  `json:"tenant"`
		} `json:"merchantCorpImages"`
		Merchantaddressinfos []struct {
			Isdefault      bool   `json:"isDefault"`
			Hasdefaultinit bool   `json:"hasDefaultInit"`
			Addresscode    string `json:"addressCode"`
			Receiver       string `json:"receiver"`
			Mobile         string `json:"mobile"`
			Address        string `json:"address"`
			Zipcode        string `json:"zipCode"`
			Regioncode     string `json:"regionCode"`
			ID             int64  `json:"id"`
			Merchantid     int64  `json:"merchantId"`
			Mergername     string `json:"mergerName"`
			Tenant         int64  `json:"tenant"`
		} `json:"merchantAddressInfos"`
		Merchantcontacterinfos []struct {
			Isdefault      bool   `json:"isDefault"`
			Hasdefaultinit bool   `json:"hasDefaultInit"`
			Fullname       FullName `json:"fullName"`
			Mobile         string `json:"mobile"`
			ID             int64  `json:"id"`
			Merchantid     int64  `json:"merchantId"`
			Tenant         int64  `json:"tenant"`
		} `json:"merchantContacterInfos"`
		Merchantagentfinancialinfos []struct {
			Accounttype     int    `json:"accountType"`
			Isdefault       bool   `json:"isDefault"`
			Stopstatus      bool   `json:"stopstatus"`
			Hasdefaultinit  bool   `json:"hasDefaultInit"`
			CurrencyName    string `json:"currency_name"`
			Currency        string `json:"currency"`
			BankName        string `json:"bank_name"`
			Bank            string `json:"bank"`
			OpenbankName    string `json:"openBank_name"`
			Openbank        string `json:"openBank"`
			Jointlineno     string `json:"jointLineNo"`
			Bankaccount     string `json:"bankAccount"`
			Bankaccountname string `json:"bankAccountName"`
			CountryName     string `json:"country_name"`
			Country         string `json:"country"`
			ProvinceName    string `json:"province_Name"`
			Province        int64  `json:"province"`
			ID              int64  `json:"id"`
			Merchantid      int64  `json:"merchantId"`
			Tenant          int64  `json:"tenant"`
		} `json:"merchantAgentFinancialInfos"`
		Merchantagentinvoiceinfos []struct {
			Isdefault            bool   `json:"isDefault"`
			Hasdefaultinit       bool   `json:"hasDefaultInit"`
			Billingtype          int    `json:"billingType"`
			Title                string `json:"title"`
			Taxno                string `json:"taxNo"`
			Receievinvoicemobile string `json:"receievInvoiceMobile"`
			ID                   int64  `json:"id"`
			Merchantid           int64  `json:"merchantId"`
			Tenant               int64  `json:"tenant"`
		} `json:"merchantAgentInvoiceInfos"`
		Merchantapplyranges []struct {
			Rangetype                                 int    `json:"rangeType"`
			Iscreator                                 bool   `json:"isCreator"`
			Isapplied                                 bool   `json:"isApplied"`
			Hasdefaultinit                            bool   `json:"hasDefaultInit"`
			Orgid                                     string `json:"orgId"`
			Tenant                                    int64  `json:"tenant"`
			Orgcode                                   string `json:"orgCode"`
			Orgname                                   string `json:"orgName"`
			ID                                        int64  `json:"id"`
			Merchantid                                int64  `json:"merchantId"`
			MerchantapplieddetailSearchcode           string `json:"merchantAppliedDetail!searchcode"`
			MerchantapplieddetailCreditserviceday     int    `json:"merchantAppliedDetail!creditServiceDay"`
			MerchantapplieddetailStopstatus           bool   `json:"merchantAppliedDetail!stopstatus"`
			MerchantapplieddetailID                   int64  `json:"merchantAppliedDetail!id"`
			MerchantapplieddetailModifytime           string `json:"merchantAppliedDetail!modifyTime"`
			MerchantapplieddetailModifier             string `json:"merchantAppliedDetail!modifier"`
			MerchantapplieddetailCreatetime           string `json:"merchantAppliedDetail!createTime"`
			MerchantapplieddetailCreator              string `json:"merchantAppliedDetail!creator"`
			MerchantapplieddetailMerchantapplyrangeid int64  `json:"merchantAppliedDetail!merchantApplyRangeId"`
			MerchantapplieddetailTenant               int64  `json:"merchantAppliedDetail!tenant"`
		} `json:"merchantApplyRanges"`
		Source       string `json:"source"`
		Merchantrole struct {
			ID         int64 `json:"id"`
			Merchantid int64 `json:"merchantId"`
			Tenant     int64 `json:"tenant"`
		} `json:"merchantRole"`
		ID                   int64  `json:"id"`
		Isenabled            bool   `json:"isEnabled"`
		Entitystatus         string `json:"entityStatus"`
		Businessrole         string `json:"businessRole"`
		Creatortype          int    `json:"creatorType"`
		Isstop               bool   `json:"isStop"`
		Iscreator            bool   `json:"isCreator"`
		Merchantapplyrangeid int64    `json:"merchantApplyRangeId"`
		Creator              string `json:"creator"`
		Creatorid            int64  `json:"creatorId"`
		Createdate           string `json:"createDate"`
		Tenant               int64  `json:"tenant"`
		Pubts                string `json:"pubts"`
	} `json:"data"`
}


//update客户联系人
type UpdateMerchantcontacterinfo struct {
	Isdefault      string `json:"isDefault"`
	Hasdefaultinit bool   `json:"hasDefaultInit"`
	PositionName   string  `json:"positionName"`
	Fullname       FullName `json:"fullName"`
	Mobile string `json:"mobile,omitempty"`
	Email string `json:"email,omitempty"`
	Wechat string `json:"weChat,omitempty"`
	Status string `json:"_status"`
	// ID     string `json:"_id"`
	TableDisplayOutlineAll  bool `json:"_tableDisplayOutlineAll,omitempty"`
}

//update设定专管业务员
type UpdatePrincipal struct {
	SpecialManagementDepCode string `json:"specialManagementDep_code"`
	IsDefault                bool   `json:"isDefault"`
	HasDefaultInit           bool   `json:"hasDefaultInit"`
	ID                       string `json:"_id,omitempty"`
	ProfessSalesmanName      string `json:"professSalesman_Name"`
	ProfessSalesman          string `json:"professSalesman"`
	SpecialManagementDep     string `json:"specialManagementDep"`
	SpecialManagementDepName string `json:"specialManagementDep_Name"`
	Status                   string `json:"_status"`
}

type UpdateMerchantrole struct {
	Status             string `json:"_status"`
	ID                 int64  `json:"id"`   
} 
type UpdateMerchantapplieddetail  struct {
	// Searchcode               string `json:"searchcode"`
	Customerlevel            int64  `json:"customerLevel"`
	CustomerlevelCode        string `json:"customerLevel_code"`
	MerchantApplyRangeId     int64     `json:"merchantApplyRangeId"` 
	// Specialmanagementdep     string `json:"specialManagementDep"`
	CustomerlevelName        string `json:"customerLevel_Name"`
	// SpecialmanagementdepName string `json:"specialManagementDep_Name"`
	// SpecialmanagementdepCode string `json:"specialManagementDep_code"`
	// ProfessSalesman          string `json:"professSalesman`
	// ProfesssalesmanName      string `json:"professSalesman_Name"`
	// Deliverywarehouse        int64  `json:"deliveryWarehouse"`
	// DeliverywarehouseName    string `json:"deliveryWarehouse_Name"`
	// Transactioncurrency      string `json:"transactionCurrency"`
	// Exchangeratetype         string `json:"exchangeratetype"`
	// TransactioncurrencyName  string `json:"transactionCurrency_Name"`
	// ExchangeratetypeName     string `json:"exchangeratetype_Name"`
	// Taxrate                  string `json:"taxRate"`
	// TaxrateName              float64    `json:"taxRate_Name"`
	Payway                   int    `json:"payway"`
	// Creditserviceday         int    `json:"creditServiceDay"`
	// Settlementmethod         int64  `json:"settlementMethod"`
	// SettlementmethodName     string `json:"settlementMethod_Name"`
	// Shipmentmethod           int64  `json:"shipmentMethod"`
	// ShipmentmethodName       string `json:"shipmentMethod_Name"`
	Signback                 bool   `json:"signBack"`
	Stopstatus               bool   `json:"stopstatus"`
	Frozenstate              int    `json:"frozenState"`
	Status                   string `json:"_status"`
	ID                       int64  `json:"id"`       
} 

type MerchantUpdateData struct {
	Data struct{
		ID                       int64       `json:"id"`
		Status                   string      `json:"_status"`
		MerchantApplyRangeID     int64       `json:"merchantApplyRangeId"` 
		CreateOrg                string      `json:"createOrg"` 
		CreateorgName            string      `json:"createOrg_name"`
		Code                     string       `json:"code"`
		Verfymark                int    `json:"verfyMark"`
		Iscreator                bool   `json:"isCreator"`
		Isapplied                bool   `json:"isApplied"`
		Belongorg                string      `json:"belongOrg"`
		Customerclass            int64       `json:"customerClass"`
		CustomerclassPath        string      `json:"customerClass_path"`
		Merchantcontacterinfos   []UpdateMerchantcontacterinfo            `json:"merchantContacterInfos"`   //联系人
		Leadername        string `json:"leaderName"`
		Creditcode        string `json:"creditCode"`
		Address           struct {
			ZhCn string `json:"zh_CN"`
		} `json:"address"`            
		Buildtime    string `json:"buildTime"`
		Money        int    `json:"money"`
		// Peoplenum    int    `json:"peopleNum"`
		// Scopemodel   int    `json:"scopeModel"`
		Scope        struct {
			ZhCn string `json:"zh_CN"`
		} `json:"scope"`             
		Taxpayingcategories    int    `json:"taxPayingCategories"`
		Enterprisenature  string `json:"enterpriseNature"`
		Enterprisename    string `json:"enterpriseName"`
		Principals               []UpdatePrincipal     `json:"principals"`         //专管人员
		Merchantapplieddetail   UpdateMerchantapplieddetail  `json:"merchantAppliedDetail"`    //客户级别
		Merchantrole           UpdateMerchantrole `json:"merchantRole"`
		Source int    `json:"source"`
	}`json:"data"`
}



type MerchantListData struct {
	PageIndex                                      int     `json:"pageIndex"`
	PageSize                                       int     `json:"pageSize"`
	Name                                           string  `json:"name"`
}
// type MerchantListData struct {
// 	MerchantapplieddetailStopstatus                bool    `json:"merchantAppliedDetail.stopstatus"`
// 	Pageindex                                      int     `json:"pageIndex"`
// 	Pagesize                                       int     `json:"pageSize"`
// 	Createorg                                      string  `json:"createOrg"`
// 	MerchantapplieddetailMerchantapplyrangeidOrgid string  `json:"merchantAppliedDetail.merchantApplyRangeId.orgId"`
// 	Code                                           string  `json:"code"`
// 	Name                                           string  `json:"name"`
// 	Shortname                                      string  `json:"shortname"`
// 	Customerclass                                  []int64 `json:"customerClass"`
// 	MerchantapplieddetailCustomerlevel             []int64 `json:"merchantAppliedDetail.customerLevel"`
// 	Customerarea                                   []int64 `json:"customerArea"`
// 	Customerindustry                               []int64 `json:"customerIndustry"`
// 	Simple                                         struct {
// 		Pubts string `json:"pubts"`
// 		Name  string `json:"name"`
// 	} `json:"simple"`
// }

type MerchantListResult struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		EndPageIndex   int    `json:"endPageIndex"`
		Pubts          string `json:"pubts"`
		BeginPageIndex int    `json:"beginPageIndex"`
		PageSize       int    `json:"pageSize"`
		RecordCount    int    `json:"recordCount"`
		PageIndex      int    `json:"pageIndex"`
		RecordList     []struct {
			Shortname struct {
				ZhCN string `json:"zh_CN"`
				EnUS string `json:"en_US"`
				ZhTW string `json:"zh_TW"`
			} `json:"shortname"`
			Money         string `json:"money"`
			BelongOrgName string `json:"belongOrg_Name"`
			CountryName   string `json:"country_name"`
			Scope         struct {
				ZhCN string `json:"zh_CN"`
				EnUS string `json:"en_US"`
				ZhTW string `json:"zh_TW"`
			} `json:"scope"`
			RetailInvestors bool `json:"retailInvestors"`
			Name            struct {
				ZhCN string `json:"zh_CN"`
				EnUS string `json:"en_US"`
				ZhTW string `json:"zh_TW"`
			} `json:"name"`
			Suppliers              int64  `json:"suppliers"`
			InvoicingCustomersName string `json:"invoicingCustomers_Name"`
			LeaderNameIDNo         string `json:"leaderNameIdNo"`
			Language               string `json:"language"`
			ContactName            string `json:"contactName"`
			CreateOrg              string `json:"createOrg"`
			FrontIDNo              string `json:"frontIdNo"`
			IsCreator              bool   `json:"isCreator"`
			OrgName                string `json:"orgName"`
			CustomerArea           int  `json:"customerArea"`
			MerchantRole           struct {
				MerchantID         int64  `json:"merchantId"`
				ToBImmigrationMode int    `json:"toBImmigrationMode"`
				CardType           int    `json:"cardType"`
				ID                 int64  `json:"id"`
				BusinessRole       string `json:"businessRole"`
				SettlementMethod   int    `json:"settlementMethod"`
			} `json:"merchantRole"`
			CreditCode     string `json:"creditCode"`
			ID             int64  `json:"id"`
			CustomerDefine struct {
				MerchantID       int64  `json:"merchantId"`
				CustomerDefine23 string `json:"customerDefine23"`
				CustomerDefine22 string `json:"customerDefine22"`
				ID               int64  `json:"id"`
				CustomerDefine24 string `json:"customerDefine24"`
			} `json:"customerDefine"`
			EnterpriseName string `json:"enterpriseName"`
			Address        struct {
				ZhCN string `json:"zh_CN"`
				EnUS string `json:"en_US"`
				ZhTW string `json:"zh_TW"`
			} `json:"address"`
			YearMoney             string `json:"yearMoney"`
			CustomerAreaName      string `json:"customerArea_Name"`
			LeaderName            string `json:"leaderName"`
			CurrencyCode          string `json:"currencyCode"`
			Stopstatus            bool   `json:"stopstatus"`
			PeopleNum             int    `json:"peopleNum"`
			MerchantApplyRangeID  int64  `json:"merchantApplyRangeId"`
			BusinessLicenseNo     string `json:"businessLicenseNo"`
			ScopeModel            int    `json:"scopeModel"`
			Code                  string `json:"code"`
			TaxPayingCategories   int    `json:"taxPayingCategories"`
			InvoicingCustomers    int64  `json:"invoicingCustomers"`
			VerfyMark             int    `json:"verfyMark"`
			InternalOrg           bool   `json:"internalOrg"`
			SuppliersName         string `json:"suppliers_Name"`
			Wid                   string `json:"wid"`
			ContactTel            string `json:"contactTel"`
			CustomerIndustry      int64  `json:"customerIndustry"`
			TimeZoneName          string `json:"timeZone_Name"`
			BelongOrg             string `json:"belongOrg"`
			LanguageName          string `json:"language_Name"`
			IsApplied             bool   `json:"isApplied"`
			TimeZone              string `json:"timeZone"`
			CustomerClassName     string `json:"customerClass_Name"`
			BusinessLicense       string `json:"businessLicense"`
			EnterpriseNature      int    `json:"enterpriseNature"`
			Country               string `json:"country"`
			MerchantAppliedDetail struct {
				MerchantID               int64  `json:"merchantId"`
				Payway                   int    `json:"payway"`
				CreditServiceDay         int    `json:"creditServiceDay"`
				CustomerLevelName        string `json:"customerLevel_Name"`
				TaxRateName              int    `json:"taxRate_Name"`
				ExchangeratetypeName     string `json:"exchangeratetype_Name"`
				ShipmentMethodName       string `json:"shipmentMethod_Name"`
				Searchcode               string `json:"searchcode"`
				ShipmentMethod           int64  `json:"shipmentMethod"`
				Creator                  string `json:"creator"`
				TransactionCurrencyName  string `json:"transactionCurrency_Name"`
				SpecialManagementDep     string `json:"specialManagementDep"`
				TransactionCurrency      string `json:"transactionCurrency"`
				SettlementMethodName     string `json:"settlementMethod_Name"`
				SpecialManagementDepName string `json:"specialManagementDep_Name"`
				DeliveryWarehouse        int64  `json:"deliveryWarehouse"`
				ModifyTime               string `json:"modifyTime"`
				SignBack                 bool   `json:"signBack"`
				SettlementMethod         int64  `json:"settlementMethod"`
				Exchangeratetype         string `json:"exchangeratetype"`
				Modifier                 string `json:"modifier"`
				DeliveryWarehouseName    string `json:"deliveryWarehouse_Name"`
				TaxRate                  string `json:"taxRate"`
				CustomerLevel            int64  `json:"customerLevel"`
			} `json:"merchantAppliedDetail"`
			CreateTime           string `json:"createTime"`
			BuildTime            string `json:"buildTime"`
			CreateOrgName        string `json:"createOrg_name"`
			CustomerIndustryName string `json:"customerIndustry_Name"`
			MerchantsManager     struct {
				ExpireDate string `json:"expireDate"`
				MerchantID int64  `json:"merchantId"`
				Email      string `json:"email"`
				Qq         string `json:"qq"`
				FullName   string `json:"fullName"`
				ID         int64  `json:"id"`
				Mobile     string `json:"mobile"`
				UserName   string `json:"userName"`
			} `json:"merchantsManager"`
			RegionCode    string `json:"regionCode"`
			CustomerClass int64  `json:"customerClass"`
			PostCode      string `json:"postCode"`
			Pubts         string `json:"pubts"`
		} `json:"recordList"`
		PageCount int `json:"pageCount"`
	} `json:"data"`
}

type MerchantDetailResult struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Country             string `json:"country"`
		InvoicingCustomerss []struct {
			IsDefault                bool   `json:"isDefault"`
			InvoicingCustomersID     int64  `json:"invoicingCustomersId"`
			MerchantID               int64  `json:"merchantId"`
			MerchantApplyRangeID     int64  `json:"merchantApplyRangeId"`
			InvoicingCustomersIDName string `json:"invoicingCustomersId_name"`
			ID                       int64  `json:"id"`
			InvoicingCustomersIDCode string `json:"invoicingCustomersId_code"`
			Pubts                    string `json:"pubts"`
		} `json:"invoicingCustomerss"`
		Suppliers              int    `json:"suppliers"`
		MasterOrgKeyField      string `json:"masterOrgKeyField"`
		ChannCustomerIndustry  int    `json:"channCustomerIndustry"`
		CustomerClassCode      string `json:"customerClass_code"`
		ParentManageOrg        string `json:"parentManageOrg"`
		ParentManageOrgName    string `json:"parentManageOrg_Name"`
		MerchantContacterInfos []struct {
			PositionName string `json:"positionName"`
			IsDefault    bool   `json:"isDefault"`
			MerchantID   int64  `json:"merchantId"`
			OrderContact bool   `json:"orderContact"`
			Mobile       string `json:"mobile"`
			FullName     struct {
				ZhCN string `json:"zh_CN"`
			} `json:"fullName"`
			CustContact bool   `json:"custContact"`
			ID          int64  `json:"id"`
			Pubts       string `json:"pubts"`
			MallContact bool   `json:"mallContact"`
		} `json:"merchantContacterInfos"`
		TaxPayingCategories    int    `json:"taxPayingCategories"`
		CustomerClass          int64  `json:"customerClass"`
		ID                     int64  `json:"id"`
		IsCreator              bool   `json:"isCreator"`
		InternalOrg            bool   `json:"internalOrg"`
		OrgName                string `json:"orgName"`
		CustomerIndustry       int    `json:"customerIndustry"`
		BuildTime              string `json:"buildTime"`
		InvoicingCustomersName string `json:"invoicingCustomers_Name"`
		CustomerArea           int    `json:"customerArea"`
		PeopleNum              int    `json:"peopleNum"`
		Shortname              struct {
			ZhCN string `json:"zh_CN"`
		} `json:"shortname"`
		ScopeModel         int    `json:"scopeModel"`
		ChannCustomerClass int64  `json:"channCustomerClass"`
		CountryCode        string `json:"country_code"`
		LeaderName         string `json:"leaderName"`
		Name               struct {
			ZhCN string `json:"zh_CN"`
		} `json:"name"`
		Code                   string `json:"code"`
		ChannCustomerClassName string `json:"channCustomerClass_Name"`
		CustomerClassName      string `json:"customerClass_Name"`
		ChannCustomerLevel     int    `json:"channCustomerLevel"`
		Principals             []struct {
			IsDefault                bool   `json:"isDefault"`
			ProfessSalesman          string `json:"professSalesman"`
			MerchantID               int64  `json:"merchantId"`
			MerchantApplyRangeID     int64  `json:"merchantApplyRangeId"`
			SpecialManagementDep     string `json:"specialManagementDep"`
			ProfessSalesmanName      string `json:"professSalesman_Name"`
			ID                       int64  `json:"id"`
			SpecialManagementDepCode string `json:"specialManagementDep_code"`
			Pubts                    string `json:"pubts"`
			SpecialManagementDepName string `json:"specialManagementDep_Name"`
		} `json:"principals"`
		InvoicingCustomers         int64  `json:"invoicingCustomers"`
		RegionCode                 string `json:"regionCode"`
		ChannelCertificationStatus int    `json:"channelCertificationStatus"`
		CreditCode                 string `json:"creditCode"`
		CreateOrgCode              string `json:"createOrg_code"`
		Scope                      struct {
			ZhCN string `json:"zh_CN"`
		} `json:"scope"`
		CountryName    string `json:"country_name"`
		IsApplied      bool   `json:"isApplied"`
		EnterpriseName string `json:"enterpriseName"`
		CreateOrg      string `json:"createOrg"`
		Address        struct {
			ZhCN string `json:"zh_CN"`
		} `json:"address"`
		YearMoney           string `json:"yearMoney"`
		VerfyMark           int    `json:"verfyMark"`
		RetailInvestors     bool   `json:"retailInvestors"`
		CustomerClassPath   string `json:"customerClass_path"`
		EnterpriseNature    int    `json:"enterpriseNature"`
		MerchantApplyRanges []struct {
			OrgName    string `json:"orgName"`
			RangeType  int    `json:"rangeType"`
			MerchantID int64  `json:"merchantId"`
			OrgCode    string `json:"orgCode"`
			IsApplied  bool   `json:"isApplied"`
			ID         int64  `json:"id"`
			Pubts      string `json:"pubts"`
			IsCreator  bool   `json:"isCreator"`
			OrgID      string `json:"orgId"`
		} `json:"merchantApplyRanges"`
		Money                 string `json:"money"`
		BelongOrgName         string `json:"belongOrg_Name"`
		CreateTime            string `json:"createTime"`
		CreateOrgName         string `json:"createOrg_name"`
		BelongOrg             string `json:"belongOrg"`
		MerchantAppliedDetail struct {
			SettlementMethod         int    `json:"settlementMethod"`
			Exchangeratetype         string `json:"exchangeratetype"`
			ModifyTime               string `json:"modifyTime"`
			Modifier                 string `json:"modifier"`
			Creator                  string `json:"creator"`
			Stopstatus               bool   `json:"stopstatus"`
			FrozenState              int    `json:"frozenState"`
			ShipmentMethod           int    `json:"shipmentMethod"`
			BelongMerchantName       string `json:"belongMerchant_Name"`
			ProfessSalesmanYhtUserID string `json:"professSalesman_yhtUserId"`
			MerchantApplyRangeID     int64  `json:"merchantApplyRangeId"`
			DeliveryWarehouse        int    `json:"deliveryWarehouse"`
			SpecialManagementDepCode string `json:"specialManagementDep_code"`
			BelongMerchant           string `json:"belongMerchant"`
			SpecialManagementDepName string `json:"specialManagementDep_Name"`
			ProfessSalesmanName      string `json:"professSalesman_Name"`
			Payway                   int    `json:"payway"`
			SignBack                 bool   `json:"signBack"`
			ID                       int64  `json:"id"`
			CreditServiceDay         int    `json:"creditServiceDay"`
			SpecialManagementDep     string `json:"specialManagementDep"`
			ProfessSalesman          string `json:"professSalesman"`
			CustomerLevel            int    `json:"customerLevel"`
			ExchangeratetypeName     string `json:"exchangeratetype_Name"`
		} `json:"merchantAppliedDetail"`
		MerchantRole struct {
			SettlementMethod   int    `json:"settlementMethod"`
			MerchantOptions    bool   `json:"merchantOptions"`
			MerchantID         int64  `json:"merchantId"`
			BusinessRole       string `json:"businessRole"`
			ToBImmigrationMode int    `json:"toBImmigrationMode"`
			ID                 int64  `json:"id"`
			CardType           int    `json:"cardType"`
		} `json:"merchantRole"`
		MerchantsManager struct {
			MerchantID int64 `json:"merchantId"`
		} `json:"merchantsManager"`
		CustomerDefine struct {
			ID         int64 `json:"id"`
			MerchantID int64 `json:"merchantId"`
		} `json:"customerDefine"`
	} `json:"data"`
}


	
type QueryTreeData struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Enable string `json:"enable"`
}

type QueryTreeResult struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		Parent        string `json:"parent"`
		Code          string `json:"code"`
		Exchangerate  string `json:"exchangerate"`
		Modifier      string `json:"modifier"`
		Effectivedate string `json:"effectivedate"`
		Innercode     string `json:"innercode"`
		Dr            int    `json:"dr"`
		Parentid      string `json:"parentid"`
		Modifiedtime  string `json:"modifiedtime"`
		Children      []struct {
			Parent        string `json:"parent"`
			Code          string `json:"code"`
			Exchangerate  string `json:"exchangerate"`
			Modifier      string `json:"modifier,omitempty"`
			Effectivedate string `json:"effectivedate,omitempty"`
			Innercode     string `json:"innercode"`
			Dr            int    `json:"dr"`
			Parentid      string `json:"parentid"`
			Modifiedtime  string `json:"modifiedtime,omitempty"`
			Enable        int    `json:"enable"`
			Tenantid      string `json:"tenantid"`
			ID            string `json:"id"`
			Isdefault     int    `json:"isdefault"`
			Isbizunit     int    `json:"isbizunit"`
			Pubts         string `json:"pubts"`
			Tenant        string `json:"tenant"`
			Orgtype       int    `json:"orgtype"`
			Creator       string `json:"creator"`
			Sysid         string `json:"sysid"`
			Level         int    `json:"level"`
			Displayorder  int    `json:"displayorder"`
			Sort          int    `json:"sort"`
			Shortname     string `json:"shortname"`
			Orgid         string `json:"orgid"`
			Companytype   string `json:"companytype"`
			Name          string `json:"name"`
			Creationtime  string `json:"creationtime"`
			Taxpayerid    string `json:"taxpayerid,omitempty"`
			Ts            string `json:"ts"`
			Children      []struct {
				Parent        string `json:"parent"`
				Code          string `json:"code"`
				Exchangerate  string `json:"exchangerate"`
				Modifier      string `json:"modifier"`
				Effectivedate string `json:"effectivedate"`
				Innercode     string `json:"innercode"`
				Dr            int    `json:"dr"`
				Parentid      string `json:"parentid"`
				Modifiedtime  string `json:"modifiedtime"`
				Enable        int    `json:"enable"`
				Tenantid      string `json:"tenantid"`
				ID            string `json:"id"`
				Isdefault     int    `json:"isdefault"`
				Isbizunit     int    `json:"isbizunit"`
				Pubts         string `json:"pubts"`
				Tenant        string `json:"tenant"`
				Orgtype       int    `json:"orgtype"`
				Creator       string `json:"creator"`
				Sysid         string `json:"sysid"`
				Level         int    `json:"level"`
				Displayorder  int    `json:"displayorder"`
				Sort          int    `json:"sort"`
				Shortname     string `json:"shortname"`
				Orgid         string `json:"orgid"`
				Companytype   string `json:"companytype"`
				Parentorgid   string `json:"parentorgid"`
				Name          string `json:"name"`
				Creationtime  string `json:"creationtime"`
				Ts            string `json:"ts"`
			} `json:"children,omitempty"`
			Parentorgid      string `json:"parentorgid,omitempty"`
			Corpid           string `json:"corpid,omitempty"`
			Timezone         string `json:"timezone,omitempty"`
			Description      string `json:"description,omitempty"`
			Language         string `json:"language,omitempty"`
			Principal        string `json:"principal,omitempty"`
			Branchleader     string `json:"branchleader,omitempty"`
			Contact          string `json:"contact,omitempty"`
			Usedtaxpayerid   string `json:"usedtaxpayerid,omitempty"`
			Address          string `json:"address,omitempty"`
			Telephone        string `json:"telephone,omitempty"`
			Taxpayername     string `json:"taxpayername,omitempty"`
			Locationid       string `json:"locationid,omitempty"`
			Usedtaxpayername string `json:"usedtaxpayername,omitempty"`
			Taxpayertype     int    `json:"taxpayertype,omitempty"`
			Depttype         string `json:"depttype,omitempty"`
		} `json:"children"`
		Enable       int    `json:"enable"`
		Tenantid     string `json:"tenantid"`
		ID           string `json:"id"`
		Isdefault    int    `json:"isdefault"`
		Isbizunit    int    `json:"isbizunit"`
		Pubts        string `json:"pubts"`
		Tenant       string `json:"tenant"`
		Orgtype      int    `json:"orgtype"`
		Creator      string `json:"creator"`
		Address      string `json:"address,omitempty"`
		Sysid        string `json:"sysid"`
		Level        int    `json:"level"`
		Telephone    string `json:"telephone,omitempty"`
		Displayorder int    `json:"displayorder"`
		Sort         int    `json:"sort"`
		Shortname    string `json:"shortname,omitempty"`
		Orgid        string `json:"orgid"`
		Taxpayername string `json:"taxpayername"`
		Companytype  string `json:"companytype"`
		Name         string `json:"name"`
		Creationtime string `json:"creationtime"`
		Taxpayerid   string `json:"taxpayerid"`
		Ts           string `json:"ts"`
		Taxpayertype int    `json:"taxpayertype,omitempty"`
	} `json:"data"`
}

type UnitResult struct {
	Code    string    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Pageindex   int `json:"pageIndex"`
		Pagesize    int `json:"pageSize"`
		Recordcount int `json:"recordCount"`
		Recordlist  []struct {
			Code          string `json:"code"`
			Modifier      string `json:"modifier"`
			Effectivedate string `json:"effectivedate"`
			Dr            int    `json:"dr"`
			Parentid      string `json:"parentid"`
			Path          string `json:"path"`
			Modifiedtime  string `json:"modifiedtime"`
			Enable        int    `json:"enable"`
			ID            string `json:"id"`
			Isdefault     int    `json:"isdefault"`
			Isbizunit     int    `json:"isbizunit"`
			Pubts         string `json:"pubts"`
			Orgtype       int    `json:"orgtype"`
			Creator       string `json:"creator"`
			Level         int    `json:"level"`
			Isend         int    `json:"isEnd"`
			Orgid         string `json:"orgid"`
			Companytype   string `json:"companytype"`
			Name          string `json:"name"`
			Parentorgid   string `json:"parentorgid"`
			Creationtime  string `json:"creationtime"`
			Externalorg   int    `json:"externalorg"`
		} `json:"recordList"`
		Pagecount      int `json:"pageCount"`
		Beginpageindex int `json:"beginPageIndex"`
		Endpageindex   int `json:"endPageIndex"`
	} `json:"data"`
}

type CusLevelData struct{
	Pageindex     int  `json:"pageIndex"`
	Pagesize      int  `json:"pageSize"`
}

type CusLevelResult struct {
	Code    string    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Pageindex   int `json:"pageIndex"`
		Pagesize    int `json:"pageSize"`
		Recordcount int `json:"recordCount"`
		Recordlist  []struct {
			Creator    string `json:"creator"`
			Code       string `json:"code"`
			Createtime string `json:"createTime"`
			Isenabled  bool   `json:"isEnabled"`
			Name       string `json:"name"`
			ID         int64  `json:"id"`
			Pubts      string `json:"pubts"`
			Order      int    `json:"order"`
		} `json:"recordList"`
		Pagecount      int    `json:"pageCount"`
		Beginpageindex int    `json:"beginPageIndex"`
		Endpageindex   int    `json:"endPageIndex"`
		Pubts          string `json:"pubts"`
	} `json:"data"`
}

type  CommenInput struct {
    PageIndex int  `json:"pageIndex"`
    PageSize  int  `json:"pageSize"`
}


type  YSProjectListData struct {
    PageIndex int  `json:"pageIndex"`
    PageSize  int  `json:"pageSize"`
}


type YSProjectListResult struct {
    Code string `json:"code"`
    Message string `json:"message"`
    Data struct {
        PageIndex int `json:"pageIndex"`
        PageSize int `json:"pageSize"`
        RecordCount int `json:"recordCount"`
        RecordList []struct {
            OrgidName string `json:"orgid_name"`
            Code string `json:"code"`
            Enable int `json:"enable"`
            Name string `json:"name"`
            Classifyid string `json:"classifyid"`
            ClassifyidName string `json:"classifyid_name"`
            Person string `json:"person"`
            PersonName string `json:"person_name"`
            Description string `json:"description"`
            ID string `json:"id"`
            Pubts string `json:"pubts"`
            Orgid string `json:"orgid"`
            Dr int `json:"dr"`
        } `json:"recordList"`
        PageCount int `json:"pageCount"`
        BeginPageIndex int `json:"beginPageIndex"`
        EndPageIndex int `json:"endPageIndex"`
    } `json:"data"`
}

type StaffListResult struct {
    Code    string `json:"code"`
    Message string `json:"message"`
    Data    struct {
        PageIndex   int `json:"pageIndex"`
        PageSize    int `json:"pageSize"`
        RecordCount int `json:"recordCount"`
        RecordList  []struct {
            Code        string `json:"code"`
            UserID      string `json:"user_id"`
            Enable      int    `json:"enable"`
            OrgID       string `json:"org_id"`
            Name        string `json:"name"`
            Mobile      string `json:"mobile"`
            Ordernumber int    `json:"ordernumber"`
            ID          string `json:"id"`
            Pubts       string `json:"pubts"`
            DeptID      string `json:"dept_id"`
            Email       string `json:"email,omitempty"`
            Dr          int    `json:"dr"`
            OrgIDName   string `json:"org_id_name,omitempty"`
            DeptIDName  string `json:"dept_id_name,omitempty"`
        } `json:"recordList"`
        PageCount      int `json:"pageCount"`
        BeginPageIndex int `json:"beginPageIndex"`
        EndPageIndex   int `json:"endPageIndex"`
    } `json:"data"`
}

type StaffDetailResult struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		CertType    string `json:"cert_type"`
		Code        string `json:"code"`
		Birthdate   string `json:"birthdate"`
		MainJobList []struct {
			JobIDName        string `json:"job_id_name"`
			JobgradeIDName   string `json:"jobgrade_id_name"`
			PostIDName       string `json:"post_id_name"`
			OrgIDName        string `json:"org_id_name"`
			Begindate        string `json:"begindate"`
			DeptIDName       string `json:"dept_id_name"`
			Director         string `json:"director"`
			PsnclID          string `json:"psncl_id"`
			Dr               int    `json:"dr"`
			JobgradeID       string `json:"jobgrade_id"`
			Enddate          string `json:"enddate"`
			Responsibilities string `json:"responsibilities"`
			PostID           string `json:"post_id"`
			OrgID            string `json:"org_id"`
			JobID            string `json:"job_id"`
			StaffID          string `json:"staff_id"`
			DirectorName     string `json:"director_name"`
			ID               string `json:"id"`
			Pubts            string `json:"pubts"`
			DeptID           string `json:"dept_id"`
			PsnclIDName      string `json:"psncl_id_name"`
		} `json:"mainJobList"`
		Remark       string `json:"remark"`
		BankAcctList []struct {
			CurrencyName string `json:"currency_name"`
			Memo         string `json:"memo"`
			Accttype     string `json:"accttype"`
			Dr           int    `json:"dr"`
			Bank         string `json:"bank"`
			BanknameName string `json:"bankname_name"`
			StaffID      string `json:"staff_id"`
			BankName     string `json:"bank_name"`
			Currency     string `json:"currency"`
			ID           string `json:"id"`
			Bankname     string `json:"bankname"`
			Isdefault    int    `json:"isdefault"`
			Pubts        string `json:"pubts"`
			Account      string `json:"account"`
		} `json:"bankAcctList"`
		Enable       int    `json:"enable"`
		Ordernumber  int    `json:"ordernumber"`
		ID           string `json:"id"`
		Pubts        string `json:"pubts"`
		Email        string `json:"email"`
		ShopAssisTag int    `json:"shop_assis_tag"`
		CertNo       string `json:"cert_no"`
		Sex          int    `json:"sex"`
		Mobile       string `json:"mobile"`
		Photo        string `json:"photo"`
		BizManTag    int    `json:"biz_man_tag"`
		PtJobList    []struct {
			JobIDName        string `json:"job_id_name"`
			JobgradeIDName   string `json:"jobgrade_id_name"`
			PostIDName       string `json:"post_id_name"`
			OrgIDName        string `json:"org_id_name"`
			Begindate        string `json:"begindate"`
			DeptIDName       string `json:"dept_id_name"`
			Director         string `json:"director"`
			PsnclID          string `json:"psncl_id"`
			Dr               int    `json:"dr"`
			JobgradeID       string `json:"jobgrade_id"`
			Enddate          string `json:"enddate"`
			Responsibilities string `json:"responsibilities"`
			PostID           string `json:"post_id"`
			OrgID            string `json:"org_id"`
			JobID            string `json:"job_id"`
			StaffID          string `json:"staff_id"`
			DirectorName     string `json:"director_name"`
			ID               string `json:"id"`
			Pubts            string `json:"pubts"`
			DeptID           string `json:"dept_id"`
			PsnclIDName      string `json:"psncl_id_name"`
		} `json:"ptJobList"`
		Officetel    string `json:"officetel"`
		CertTypeName string `json:"cert_type_name"`
		Name         string `json:"name"`
		Objid        string `json:"objid"`
	} `json:"data"`
}