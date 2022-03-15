package backend

import(
   "time"
   // "sort"
)

type Result struct{
    CorpCode   string
    CorpName   string
    Logo       string 
    Address    string
    Trade      string
    Region     string
    UpdatedAt   time.Time
    CreatedAt   time.Time
    Genjin     string   //跟进人
    Genjins    []string  `gorm:"type:text"` //跟进人
    Contact    string    `gorm:"omitempty"`  //联系人
    RecordTime  time.Time  //跟进记录创建时间
}

type Results  []Result

func (s Results) Len() int {
    return len(s)
}

func (s Results) Less(i, j int) bool {
   if s[i].RecordTime.Year() > 1973 && s[j].RecordTime.Year() > 1973 {
      if  s[i].UpdatedAt.After(s[j].RecordTime){
         return s[i].UpdatedAt.After(s[j].RecordTime)
      }
    return s[i].RecordTime.After(s[j].RecordTime)
   }
   return s[i].UpdatedAt.After(s[j].UpdatedAt)
}

func (s Results) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}


//--------------------------------Customers------------------------------------------------------
//查看销售人员的客户  list
func(db *DB)ViewUserCustomers(user string)Results{
     //CustomerPrincipal  客户销售对应
   var ab  Results
   db.Logger.Debugln("View User: ",user,"'s customers.")
	db.DB.Model(&CustomerUserprincipal{}).Select("customer_userprincipals.corp_code,customers.corp_name,customers.logo,customers.address,customers.trade,customers.region,customers.updated_at,customers.created_at").Joins("JOIN customers  ON customers.corp_code = customer_userprincipals.corp_code").Joins("JOIN dmai_users d ON d.userid = customer_userprincipals.merchandiserid").Where("d.userid = ?", user).Find(&ab)
   return ab
}

//查看销售人员昨天新增的客户  list
func(db *DB)ViewUserYesterdayCustomers(user,yesterday string)Results{
     //CustomerPrincipal  客户销售对应
    var ab   Results
   db.Logger.Debugln("View User: ",user,"'s customers.")
   db.DB.Model(&CustomerUserprincipal{}).Select("customer_userprincipals.corp_code,customers.corp_name,customers.logo,customers.address,customers.updated_at,customers.created_at").Joins("JOIN customers  ON customers.corp_code = customer_userprincipals.corp_code").Joins("JOIN dmai_users d ON d.userid = customer_userprincipals.merchandiserid").Where("d.userid = ? and unix_timestamp(customers.created_at) > ?", user,yesterday).Find(&ab)
   return ab
}


//查看某一个客户的详情  map
func (db *DB)ViewCustomerInfo(corpcode string)*Customer{
   var customer Customer
   db.DB.Where("corp_code=?",corpcode).Find(&customer)
   db.Logger.Debugln("view customer details: ",customer)
   return &customer
}

//查看某一个客户的详情  map
func (db *DB)CheckCustomer(name string)*Customer{
   customer := Customer{}
   db.DB.Where("corp_name = ?",name).Find(&customer)
   db.Logger.Debugln("view customer details: ",customer)
   return &customer
}


//更新客户信息
func (db *DB)UpdateCustomer(customer *Customer)error{
	db.Logger.Debugln("update customer: ",*customer)
    db.DB.Model(&Customer{}).Where("corp_code = ?",customer.CorpCode).Updates(customer)
    return db.DB.Error
}
//添加客户
func (db *DB)AddCustomer(customer *Customer)error{
	db.Logger.Debugln("add a new customer: ",*customer)
    db.DB.Create(customer)
    return db.DB.Error
}
//删除客户
func (db *DB)RMCustomer(corpcode string)error{
  db.Logger.Debugln("delete customer, here is a corpcode: ",corpcode)
  var customer Customer
  db.DB.Where("corp_code=?",corpcode).Find(&customer)
  db.DB.Delete(&Customer{},customer.Model.ID)
  return db.DB.Error
}

//查看所有客户list
func (db *DB)ViewCustomers()Results{
   var ab Results
   db.DB.Model(&Customer{}).Select("customers.corp_code,customers.corp_name,customers.logo,customers.address,customers.trade,customers.region,customers.updated_at,customers.created_at").Find(&ab)
   db.Logger.Debugln("view customer details: ",ab)
   return ab
}



//查看昨天新增的客户list
func (db *DB)ViewYesterdayCustomers(yesterday string)Results{
   var ab Results
   db.DB.Model(&Customer{}).Select("customers.corp_code,customers.corp_name,customers.logo,customers.address,customers.updated_at,customers.created_at").Where("unix_timestamp(customers.created_at) > ?",yesterday).Find(&ab)
   db.Logger.Debugln("查看昨天新增的客户list : ",ab)
   return ab
}

//--------------------------------Customers 失信验证信息------------------------------------------------------
//更新一条失信核验信息
func (db *DB)UpdateCustomerShixin(shixin *CustomerShiXin)error{
   db.Logger.Debugln("update customer shixin: ",*shixin)
   db.DB.Model(&CustomerShiXin{}).Where("id = ?",shixin.Model.ID).Updates(shixin)
   return db.DB.Error
}
//添加一条失信核验信息
func (db *DB)AddCustomerShixin(shixin *CustomerShiXin)error{
   db.Logger.Debugln("add a new customer shixin: ",*shixin)
   db.DB.Create(shixin)
   return db.DB.Error
}
//删除一条失信核验信息
func (db *DB)RMCustomerShixin(id string)error{
  db.Logger.Debugln("delete customer shixin cheker, here is a id: ",id)
  shixin := CustomerShiXin{}
  db.DB.Where("id=?",id).Find(&shixin)
  db.DB.Delete(&CustomerShiXin{},shixin.Model.ID)
  return db.DB.Error
}
//查看某一个失信核验信息
func (db *DB)ViewCustomerShixin(corpcode string)[]CustomerShiXin{
   shixins := []CustomerShiXin{}
   db.DB.Where("corp_code=?",corpcode).Find(&shixins)
   // db.Logger.Debugln("view customer shixin details: ",shixins)
   return shixins
}

func(db *DB)TruncateShixin(corpcode string)error{
   db.DB.Delete(&CustomerShiXin{}, "corp_code LIKE ?", corpcode)
   db.Logger.Debugln("delete customer shixin corpcode:",corpcode)
   return db.DB.Error
}




//--------------------------------Customers 严重违法验证信息------------------------------------------------------
//更新严重违法信息
func (db *DB)UpdateCustomerSeriousIllegal(serious *CustomerSeriousIllegal)error{
   db.Logger.Debugln("update customer serious illegal: ",*serious)
   db.DB.Model(&CustomerSeriousIllegal{}).Where("id = ?",serious.Model.ID).Updates(serious)
   return db.DB.Error
}
//添加一条严重违法信息
func (db *DB)AddCustomerSeriousIllegal(serious *CustomerSeriousIllegal)error{
   db.Logger.Debugln("add a new customer serious illegal: ",*serious)
   db.DB.Create(serious)
   return db.DB.Error
}
//删除一条严重违法信息
func (db *DB)RMCustomerSeriousIllegal(id string)error{
  db.Logger.Debugln("delete customer, here is a id: ",id)
  serious := CustomerSeriousIllegal{}
  db.DB.Where("id=?",id).Find(&serious)
  db.DB.Delete(&CustomerSeriousIllegal{},serious.Model.ID)
  return db.DB.Error
}
//查看某一个客户严重违法信息
func (db *DB)ViewCustomerSeriousIllegal(corpcode string)[]CustomerSeriousIllegal{
   serious := []CustomerSeriousIllegal{}
   db.DB.Where("corp_code=?",corpcode).Find(&serious)
   // db.Logger.Debugln("view customer SeriousIllegals: ",serious)
   return serious
}

func(db *DB)TruncateSerious(corpcode string)error{
   db.DB.Delete(&CustomerSeriousIllegal{}, "corp_code LIKE ?", corpcode)
   db.Logger.Debugln("delete customer serious corpcode:",corpcode)
   return db.DB.Error
}

//--------------------------------Customers 税收验证信息------------------------------------------------------
//更新客户税收处罚信息
func (db *DB)UpdateCustomerTaxIllegal(tax *CustomerTaxIllegal)error{
   db.Logger.Debugln("update customer tax : ",*tax)
   db.DB.Model(&CustomerTaxIllegal{}).Where("id = ?",tax.Model.ID).Updates(tax)
   return db.DB.Error
}
//添加一条客户税收处罚信息
func (db *DB)AddCustomerTaxIllegal(tax *CustomerTaxIllegal)error{
   db.Logger.Debugln("add a new customer: ",*tax)
   db.DB.Create(tax)
   return db.DB.Error
}
//删除一条税收处罚信息
func (db *DB)RMCustomerTaxIllegal(id string)error{
  db.Logger.Debugln("delete customer tax, here is a id: ",id)
  tax := CustomerTaxIllegal{}
  db.DB.Where("id=?",id).Find(&tax)
  db.DB.Delete(&CustomerTaxIllegal{},tax.Model.ID)
  return db.DB.Error
}
//查看某一个客户税收核验信息
func (db *DB)ViewCustomerTaxIllegal(corpcode string)[]CustomerTaxIllegal{
   taxs := []CustomerTaxIllegal{}
   db.DB.Where("corp_code=?",corpcode).Find(&taxs)
   // db.Logger.Debugln("view customer TaxIllegals: ",taxs)
   return taxs
}

func(db *DB)TruncateTax(corpcode string)error{
   db.DB.Delete(&CustomerTaxIllegal{}, "corp_code LIKE ?", corpcode)
   db.Logger.Debugln("delete customer serious corpcode:",corpcode)
   return db.DB.Error
}

//--------------------------------Customers 行政处罚验证信息------------------------------------------------------
//更新客户行政处罚信息
func (db *DB)UpdateCustomerAdminPenalty (penalty *CustomerAdminPenalty )error{
   db.Logger.Debugln("update customer admin penalty: ",*penalty)
   db.DB.Model(&CustomerAdminPenalty {}).Where("id = ?",penalty.Model.ID).Updates(penalty)
   return db.DB.Error
}
//添加一条客户行政处罚
func (db *DB)AddCustomerAdminPenalty(penalty *CustomerAdminPenalty )error{
   db.Logger.Debugln("add a new customer admin penalty: ",*penalty)
   db.DB.Create(penalty)
   return db.DB.Error
}
//删除一个客户行政处罚信息
func (db *DB)RMCustomerAdminPenalty(id string)error{
  db.Logger.Debugln("delete customer admin penalty, here is a id: ",id)
  penalty := CustomerAdminPenalty {}
  db.DB.Where("id=?",id).Find(&penalty)
  db.DB.Delete(&CustomerAdminPenalty {},penalty.Model.ID)
  return db.DB.Error
}
//查看某一个客户相关行政处罚信息
func (db *DB)ViewCustomerAdminPenalty(corpcode string)[]CustomerAdminPenalty{
   penalties := []CustomerAdminPenalty {}
   db.DB.Where("corp_code=?",corpcode).Find(&penalties)
   // db.Logger.Debugln("view customer AdminPenaltys: ",penalties)
   return penalties
}

func(db *DB)TruncatePenalty(corpcode string)error{
   db.DB.Delete(&CustomerAdminPenalty{}, "corp_code LIKE ?", corpcode)
   db.Logger.Debugln("delete customer serious corpcode:",corpcode)
   return db.DB.Error
}

//--------------------------------Customers 保存结果表------------------------------------------------------
//更新用友客户保存执行结果
func (db *DB)UpdateCustomerSave(yscustomer *YsCustomer)error{
   db.Logger.Debugln("update customer: ",*yscustomer)
    db.DB.Model(&YsCustomer{}).Where("cid = ?",yscustomer.Cid).Updates(yscustomer)
    return db.DB.Error
}
//添加用友客户保存执行结果
func (db *DB)AddCustomerSave(yscustomer *YsCustomer)error{
   db.Logger.Debugln("add a new customer: ",*yscustomer)
   db.DB.Create(yscustomer)
   return db.DB.Error
}
//删除用友客户保存执行结果
func (db *DB)RMCustomerSave(cid string)error{
  db.Logger.Debugln("delete customer, here is a corpcode: ",cid)
  yscustomer := YsCustomer{}
  db.DB.Where("cid=?",cid).Find(&yscustomer)
  db.DB.Delete(&YsCustomer{},yscustomer.Model.ID)
  return db.DB.Error
}
//查看某一个客户保存执行结果
func (db *DB)ViewCustomerInfoSave(cid string)*YsCustomer{
   yscustomer := YsCustomer{}
   db.DB.Where("cid=?",cid).Find(&yscustomer)
   db.Logger.Debugln("view customer details: ",yscustomer)
   return &yscustomer
}