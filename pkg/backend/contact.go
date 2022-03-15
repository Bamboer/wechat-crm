package backend


//------------------------------------------contactor--------------------------------------
//添加联系人
func(db *DB)AddContactor(contact *Contactor)error{
	db.Logger.Debugln("add a new contact: ",*contact)
    db.DB.Create(contact)
    return db.DB.Error
}
//删除联系人
func(db *DB)DelContactor(contactcode string)error{
	db.Logger.Debugln("delete a contact,here is a contact id: ",contactcode)
    contact := &Contactor{}
    db.DB.Where("contact_code=?",contactcode).Find(contact)
    db.DB.Delete(&Contactor{},contact.Model.ID)
    return db.DB.Error
}
//查询某个客户所有的联系人
func(db *DB)ViewCustomerContact(corpcode string)[]Contactor{
    var contacts []Contactor
    db.DB.Where("corp_code=?",corpcode).Find(&contacts)
    db.Logger.Debugln("view contact detail: ",contacts)
    return contacts
}
//查询单个联系人信息
func(db *DB)ViewWechatContactor(wechatid string)[]Contactor{
    contact := []Contactor{}
    if wechatid == ""{
        return  contact
    }
    db.DB.Model(&Contactor{}).Where("wechatid=?",wechatid).Find(&contact)
    if db.DB.Error != nil{
        db.Logger.Errorln("View Wechat Contactor error: ",db.DB.Error)
    }
    return contact
}

//更新联系人
func(db *DB)UpdateContactor(contact *Contactor)error{
    db.DB.Model(&Contactor{}).Where("contact_code = ?",contact.ContactCode).Updates(contact)
    db.Logger.Debugln("update contactor: ",*contact)
    return db.DB.Error
}

//查询某个销售人员的所有联系人
// type ResultUser struct{
//     CorpName    string 
//     ContactCode  string
//     Wechatid     string
//     Name        string
//     Logo        string
// }
// func(db *DB)ViewUserContacts(user string)[]ResultUser{
//     var contacts []ResultUser
//     db.DB.Model(&Contactor{}).Select("customers.corp_name,contactors.contact_code,contactors.wechatid,contactors.name,contactors.logo").Joins("JOIN contact_userprincipals cp  ON cp.corp_code = contactors.corp_code").Joins("JOIN customers  ON customers.corp_code = cp.corp_code").Where("cp.merchandiserid= ?", user).Find(&contacts)
//     db.Logger.Debugln("view contact detail: ",contacts)
//     return contacts
// }
//查看所有联系人
func(db *DB)ViewContacts()[]Contactor{
    var contacts []Contactor
    db.DB.Model(&Contactor{}).Find(&contacts)
    if db.DB.Error != nil{
        db.Logger.Errorln(db.DB.Error)
    }
    
    return contacts
}

//根据ID查询单个联系人信息
func(db *DB)ViewIdContactor(contactcode string)*Contactor{
    contact := &Contactor{}
    db.DB.Model(&Contactor{}).Where("contact_code=?",contactcode).Find(contact)
    if db.DB.Error != nil{
        db.Logger.Errorln("ViewSingleContactor error: ",db.DB.Error)
    }
    
    return contact
}
