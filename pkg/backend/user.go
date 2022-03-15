package backend

//员工信息
func(db *DB)AddUser(user *DmaiUser)error{
    db.DB.Create(user)
    db.Logger.Debugln("add a new member: ",*user)
    return db.DB.Error
}
func(db *DB)UpdateUser(user *DmaiUser)error{
    db.DB.Where("userid = ?",user.Userid).Updates(user)
    db.Logger.Debugln("update member: ",*user)
    return db.DB.Error
}

func(db *DB)ViewUser(user string)*DmaiUser{
	dmaiuser := DmaiUser{}
    db.DB.Model(&DmaiUser{}).Where("userid = ?",user).Find(&dmaiuser)
    if db.DB.Error != nil{
       db.Logger.Errorln(db.DB.Error) 
    }
    
    return &dmaiuser
}
//查看所有员工
func(db *DB)ViewUsers()[]DmaiUser{
    dmaiusers := []DmaiUser{}
    db.DB.Model(&DmaiUser{}).Find(&dmaiusers)
    if db.DB.Error != nil{
      db.Logger.Errorln(db.DB.Error)
    }

    return dmaiusers
}


//基于用户中文名查询
func(db *DB)ViewUserName(username string)*DmaiUser{
    dmaiuser := DmaiUser{}
    db.DB.Model(&DmaiUser{}).Where("name like ?",username).Find(&dmaiuser)
    if db.DB.Error != nil{
       db.Logger.Errorln(db.DB.Error) 
    }
    db.Logger.Debugln(dmaiuser)
    return &dmaiuser
}