package backend


//--------------------------Customer User Principal----------------------------------------
func(db *DB)ViewCustomerUserprincipal(corpcode string)[]CustomerUserprincipal{
   result := []CustomerUserprincipal{} 
   db.DB.Model(&CustomerUserprincipal{}).Where("corp_code = ? ",corpcode).Find(&result)
   // db.Logger.Debugln(corpcode, " view customer user principal: ",result)
   return result
}
//查看某个用户是否归属customer
func(db *DB)CheckCUP(userid,corpcode string)bool{
    p := &CustomerUserprincipal{}
    db.DB.Model(&CustomerUserprincipal{}).Where("corp_code=? and merchandiserid =?",corpcode,userid).Find(p)
    db.Logger.Debugln("view customer customer user principal: ",*p)
    if p.Merchandiserid  != ""{
        return true
    }else{
        return false
    }
}

//查看某条记录
func(db *DB)ViewSingleCP(userid,corpcode string)*CustomerUserprincipal{
    p := &CustomerUserprincipal{}
    db.DB.Model(&CustomerUserprincipal{}).Where("corp_code=? and merchandiserid =?",corpcode,userid).Find(p)
    db.Logger.Debugln("view customer customer user principal: ",*p)
    return p
}

func(db *DB)AddCustomerUserprincipal(info *CustomerUserprincipal)error{
    db.DB.Create(info)
    db.Logger.Debugln("add customer user principal: ",*info)
    return db.DB.Error
}

func(db *DB)DelCustomerUserprincipal(userid,corpcode string)error{
    p := &CustomerUserprincipal{}
    db.DB.Model(&CustomerUserprincipal{}).Where("corp_code=? and merchandiserid =?",corpcode,userid).Find(p)
    db.DB.Delete(&CustomerUserprincipal{},p.Model.ID)
    db.Logger.Debugln("delete customer user principal id:",userid)
    return db.DB.Error
}

func(db *DB)TruncateCUP(corpcode string)error{
    // db.DB.Where("corp_code=?",corpcode).Delete(&CustomerUserprincipal{})
    db.DB.Delete(&CustomerUserprincipal{}, "corp_code LIKE ?", corpcode)
    db.Logger.Debugln("delete customer user principal corpcode:",corpcode)
    return db.DB.Error
}

func(db *DB)UpdateCustomerUserprincipal(info *CustomerUserprincipal)error{
    db.DB.Where("id=?",info.Model.ID).Updates(info)
    db.Logger.Debugln("update customer user principal: ",*info)
    return db.DB.Error
}


//----------------------------Contact User Principal---------------------------------------
func(db *DB)AddContactUserprincipal(info *ContactUserprincipal)error{
    db.DB.Create(info)
    db.Logger.Debugln("add new contactor user principal: ",*info)
    return db.DB.Error
}

func(db *DB)DelContactUserprincipal(id string)error{
    db.DB.Delete(&ContactUserprincipal{}, id)
    db.Logger.Debugln("delete a contactor user principal id: ",id)
    return db.DB.Error
}

func(db *DB)UpdateContactprincipal(info *ContactUserprincipal)error{
    db.DB.Where("id = ?",info.Model.ID).Updates(info)
    db.Logger.Debugln("update contactor user principal: ",*info)
    return db.DB.Error
}

func(db *DB)ViewContactUserprincipal(corpcode string)[]CustomerUserprincipal{
   result := []CustomerUserprincipal{} 
   db.DB.Model(&CustomerUserprincipal{}).Where("corp_code = ? ",corpcode).Find(&result)
   db.Logger.Debugln(corpcode, " view customer user principal: ",result)
   return result
}




//----------------------------Productline---------------------------------------
// func(db *DB)ViewProductline(name string)*Productline{
//    productline := Productline{} 
//    db.DB.Model(&Productline{}).Where("name = ? ",name).Find(&productline)
//    db.Logger.Debugln(name, " view product line: ",productline)
//    return &productline
// }

// func(db *DB)ViewProductlines()[]Productline{
//    productlines := []Productline{} 
//    db.DB.Model(&Productline{}).Find(&productlines)
//    db.Logger.Debugln("view product lines: ",productlines)
//    return productlines
// }

// func(db *DB)AddProductline(info *Productline)error{
//     db.DB.Create(info)
//     db.Logger.Debugln("add new product line: ",*info)
//     return db.DB.Error
// }

// func(db *DB)DelProductline(id string)error{
//     db.DB.Where("id=?",id).Delete(&Productline{})
//     db.Logger.Debugln("delete a product line id: ",id)
//     return db.DB.Error
// }

// func(db *DB)UpdateProductline(info *Productline)error{
//     db.DB.Where("id = ?",info.Model.ID).Updates(info)
//     db.Logger.Debugln("update product line: ",*info)
//     return db.DB.Error
// }

// //----------------------------Productline Customer Principal---------------------------------------
// func(db *DB)ViewProductlineCustomerPrincipal(corpcode string)[]ProductlineCustomerprincipal{
//    result := []ProductlineCustomerprincipal{} 
//    db.DB.Model(&ProductlineCustomerprincipal{}).Where("corp_code = ? ",corpcode).Find(&result)
//    db.Logger.Debugln(corpcode, " view product line principal: ",result)
//    return result
// }

// func(db *DB)AddProductlineCustomerPrincipal(info *ProductlineCustomerprincipal)error{
//     db.DB.Create(info)
//     db.Logger.Debugln("add new product line customer principal: ",*info)
//     return db.DB.Error
// }

// func(db *DB)DelProductlineCustomerPrincipal(id string)error{
//     db.DB.Where("id=?",id).Delete(&ProductlineCustomerprincipal{})
//     db.Logger.Debugln("delete a product line principal id: ",id)
//     return db.DB.Error
// }

// func(db *DB)UpdateProductlineCustomerPrincipal(info *ProductlineCustomerprincipal)error{
//     db.DB.Where("id = ?",info.Model.ID).Updates(info)
//     db.Logger.Debugln("update product line principal: ",*info)
//     return db.DB.Error
// }


//----------------------------Agreement---------------------------------------
//查看某个客户的所有合同
func(db *DB)ViewAgreement(corpcode string)[]Agreement{
   result := []Agreement{} 
   db.DB.Model(&Agreement{}).Where("corp_code = ? ",corpcode).Find(&result)
   db.Logger.Debugln(corpcode, " view Agreement: ",result)
   return result
}

//查看某个客户的所有合同
func(db *DB)ViewSingleAgreement(id string)*Agreement{
   result  := &Agreement{}
   db.DB.Model(&Agreement{}).Where("id = ? ",id).Find(result)
   db.Logger.Debugln(id, " view Agreement: ",result)
   return result
}


func(db *DB)AddAgreement(info *Agreement)error{
    db.DB.Create(info)
    db.Logger.Debugln("add new Agreement: ",*info)
    return db.DB.Error
}

func(db *DB)DelAgreement(id string)error{
    db.DB.Delete(&Agreement{}, id)
    db.Logger.Debugln("delete a Agreement id: ",id)
    return db.DB.Error
}

func(db *DB)UpdateAgreement(info *Agreement)error{
    db.DB.Where("id = ?",info.Model.ID).Updates(info)
    db.Logger.Debugln("update Agreement: ",*info)
    return db.DB.Error
}


// ------------------Project 项目-----------------------------

func(db *DB)AddProject(p *Project)error{
    db.Logger.Debugln("add a new project: ",*p)
    db.DB.Create(p)
    return db.DB.Error
}

//删除项目
func(db *DB)DelProject(pid string)error{
    db.Logger.Debugln("delete a project,here is a project id: ", pid)
    db.DB.Delete(&Project{}, pid)
    return db.DB.Error
}


//更新项目
func(db *DB)UpdateProject(p *Project)error{
    db.DB.Model(&Project{}).Where("id = ?",p.Model.ID).Updates(p)
    db.Logger.Debugln("update project: ",*p)
    return db.DB.Error
}
//查询某个项目
func(db *DB)ViewSingleProject(id string)*Project{
    p := &Project{}
    db.DB.Model(&Project{}).Where("id = ?", id).Find(p)
    db.Logger.Debugln("view customer projects: ",p)
    return p
}
//查看是否存在 Name
func(db *DB)CheckProject(name string)*Project{
    p := &Project{}
    db.DB.Model(&Project{}).Where("name = ?",name).Find(p)
    db.Logger.Errorln(db.DB.Error)
    return p
}

//查询某个客户的项目
func(db *DB)ViewProject(corpcode string)[]Project{
    var p []Project
    db.DB.Model(&Project{}).Where("corp_code=?",corpcode).Find(&p)
    db.Logger.Debugln("view customer projects: ",p)
    return p
}

//查看所有项目
func(db *DB)ViewProjectsP()[]Project{
    var p []Project
    db.DB.Model(&Project{}).Find(&p)
    db.Logger.Errorln(db.DB.Error)
    return p
}


// ------------------product产品线-----------------------------

func(db *DB)AddProduct(p *Product)error{
    db.Logger.Debugln("add a new product: ",*p)
    db.DB.Create(p)
    return db.DB.Error
}

//删除产品线
func(db *DB)DelProduct(pid string)error{
    db.Logger.Debugln("delete a product,here is a product id: ", pid)
//    db.DB.Where("id=?",pid).Delete(Product{})
    db.DB.Delete(&Product{}, pid)
    return db.DB.Error
}


//更新产品线
func(db *DB)UpdateProduct(p *Product)error{
    db.DB.Model(&Product{}).Where("id = ?",p.Model.ID).Updates(p)
    db.Logger.Debugln("update product: ",*p)
    return db.DB.Error
}

//查看所有产品线
func(db *DB)ViewProducts()[]Product{
    var p []Product
    db.DB.Model(&Product{}).Find(&p)
    db.Logger.Errorln(db.DB.Error)
    return p
}
//查看是否存在 Name
func(db *DB)CheckProduct(name string)*Product{
    p := &Product{}
    db.DB.Model(&Product{}).Where("name = ?",name).Find(p)
    db.Logger.Errorln(db.DB.Error)
    return p
}

//查看产品
func(db *DB)ViewProduct(id string)*Product{
    p := &Product{}
    db.DB.Model(&Product{}).Where("id = ?",id).Find(p)
    db.Logger.Errorln(db.DB.Error)
    return p
}