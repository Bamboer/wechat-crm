package service



//企查查模糊搜索 
type FuzzySearch struct {
	Paging struct {
		Pagesize     int `json:"PageSize"`
		Pageindex    int `json:"PageIndex"`
		Totalrecords int `json:"TotalRecords"`
	} `json:"Paging"`
	Result []struct {
		Keyno      string `json:"KeyNo"`
		Name       string `json:"Name"`
		Creditcode string `json:"CreditCode"`
		Startdate  string `json:"StartDate"`
		Opername   string `json:"OperName"`
		Status     string `json:"Status"`
		No         string `json:"No"`
	} `json:"Result"`
	Status      string `json:"Status"`
	Message     string `json:"Message"`
	Ordernumber string `json:"OrderNumber"`
}

//工商数据查询
type SearchWide struct {
	Ordernumber string `json:"OrderNumber"`
	Paging      struct {
		Pagesize     int `json:"PageSize"`
		Pageindex    int `json:"PageIndex"`
		Totalrecords int `json:"TotalRecords"`
	} `json:"Paging"`
	Result []struct {
		Dimension  string `json:"Dimension"`
		Keyno      string `json:"KeyNo"`
		Name       string `json:"Name"`
		Opername   string `json:"OperName"`
		Startdate  string `json:"StartDate"`
		Status     string `json:"Status"`
		No         string `json:"No"`
		Creditcode string `json:"CreditCode"`
	} `json:"Result"`
	Status  string `json:"Status"`
	Message string `json:"Message"`
}

//企查查工商详情 查询
type GetBasicDetailsByName struct {
	Ordernumber string `json:"OrderNumber"`
	Result      struct {
		Keyno        string        `json:"KeyNo"`
		Name         string        `json:"Name"`
		No           string        `json:"No"`
		Belongorg    string        `json:"BelongOrg"`
		Operid       string        `json:"OperId"`
		Opername     string        `json:"OperName"`
		Startdate    string        `json:"StartDate"`
		Enddate      string        `json:"EndDate"`
		Status       string        `json:"Status"`
		Province     string        `json:"Province"`
		Updateddate  string        `json:"UpdatedDate"`
		Creditcode   string        `json:"CreditCode"`
		Registcapi   string        `json:"RegistCapi"`
		Econkind     string        `json:"EconKind"`
		Address      string        `json:"Address"`
		Scope        string        `json:"Scope"`
		Termstart    string        `json:"TermStart"`
		Teamend      string        `json:"TeamEnd"`
		Checkdate    string        `json:"CheckDate"`
		Orgno        string        `json:"OrgNo"`
		Isonstock    string        `json:"IsOnStock"`
		Stocknumber  interface{}   `json:"StockNumber"`
		Stocktype    interface{}   `json:"StockType"`
		Originalname []interface{} `json:"OriginalName"`
		Imageurl     string        `json:"ImageUrl"`
		Enttype      string        `json:"EntType"`
		Reccap       string        `json:"RecCap"`
	} `json:"Result"`
	Status  string `json:"Status"`
	Message string `json:"Message"`
}


// http://api.qichacha.com/ShixinCheck/GetList   失信检查    
type ShiXinCheck struct {
	Paging struct {
		PageSize     int `json:"PageSize"`
		PageIndex    int `json:"PageIndex"`
		TotalRecords int `json:"TotalRecords"`
	} `json:"Paging"`
	Result struct {
		VerifyResult int `json:"VerifyResult"`
		Data         []struct {
			ID            string `json:"Id"`
			Liandate      string `json:"Liandate"`
			Anno          string `json:"Anno"`
			Executegov    string `json:"Executegov"`
			Executestatus string `json:"Executestatus"`
			Publicdate    string `json:"Publicdate"`
			Executeno     string `json:"Executeno"`
		} `json:"Data"`
	} `json:"Result"`
	Status      string `json:"Status"`
	Message     string `json:"Message"`
	OrderNumber string `json:"OrderNumber"`
}

//严重违法核查
//http://api.qichacha.com/SeriousIllegalCheck/GetList
type SeriousIllegalcheck struct {
	Result struct {
		VerifyResult int `json:"VerifyResult"`
		Data         []struct {
			Type         string `json:"Type"`
			AddReason    string `json:"AddReason"`
			AddDate      string `json:"AddDate"`
			AddOffice    string `json:"AddOffice"`
			RemoveReason string `json:"RemoveReason"`
			RemoveDate   string `json:"RemoveDate"`
			RemoveOffice string `json:"RemoveOffice"`
		} `json:"Data"`
	} `json:"Result"`
	Status      string `json:"Status"`
	Message     string `json:"Message"`
	OrderNumber string `json:"OrderNumber"`
}

//税收违法核查
// http://api.qichacha.com/TaxIllegalCheck/GetList
type TaxIllegalCheck struct {
	Paging struct {
		PageSize     int `json:"PageSize"`
		PageIndex    int `json:"PageIndex"`
		TotalRecords int `json:"TotalRecords"`
	} `json:"Paging"`
	Result struct {
		VerifyResult int `json:"VerifyResult"`
		Data         []struct {
			ID          string `json:"Id"`
			PublishDate string `json:"PublishDate"`
			CaseNature  string `json:"CaseNature"`
			TaxGov      string `json:"TaxGov"`
		} `json:"Data"`
	} `json:"Result"`
	Status      string `json:"Status"`
	Message     string `json:"Message"`
	OrderNumber string `json:"OrderNumber"`
}


//行政处罚核查
// http://api.qichacha.com/AdminPenaltyCheck/GetList?key=AppKey&searchKey=XXXXXX
type AdminPenalty struct {
	Paging struct {
		PageSize     int `json:"PageSize"`
		PageIndex    int `json:"PageIndex"`
		TotalRecords int `json:"TotalRecords"`
	} `json:"Paging"`
	Result struct {
		VerifyResult int `json:"VerifyResult"`
		Data         []struct {
			ID           string `json:"Id"`
			DocNo        string `json:"DocNo"`
			PunishReason string `json:"PunishReason"`
			PunishResult string `json:"PunishResult"`
			PunishOffice string `json:"PunishOffice"`
			PunishDate   string `json:"PunishDate"`
			Source       string `json:"Source"`
			SourceCode   string `json:"SourceCode"`
		} `json:"Data"`
	} `json:"Result"`
	Status      string `json:"Status"`
	Message     string `json:"Message"`
	OrderNumber string `json:"OrderNumber"`
}



//基础信用报告下单
//POST http://api.qichacha.com/ReportBase/CreateReport?key=AppKey&searchKey={searchKey}
type ReportBase struct {
	Result struct {
		OrderNo string `json:"OrderNo"`
	} `json:"Result"`
	Status  string `json:"Status"`
	Message string `json:"Message"`
}

//获取基础信用报告
//GET http://api.qichacha.com/ReportBase/GetReportInfo?key=AppKey&orderNo={orderNo}
type ReportBaseResult struct {
	Result struct {
		ReportStatus string `json:"ReportStatus"`
		ReportURL    string `json:"ReportUrl"`
	} `json:"Result"`
	Status  string `json:"Status"`
	Message string `json:"Message"`
}






