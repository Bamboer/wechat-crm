package backend

//=========================================客户联系人变更动作记录=============================
func(db *DB)AddWxlog(action *Wxlog)error{
    db.DB.Create(action)
    db.Logger.Debugln("add a action log: ",*action)
    return db.DB.Error
}
func(db *DB)UpdateWxlog(action *Wxlog)error{
    db.DB.Where("id = ?",action.Model.ID).Updates(action)
    db.Logger.Debugln("update member: ",*action)
    return db.DB.Error
}

func(db *DB)ViewWxlog(action string)*Wxlog{
    wxlog := Wxlog{}
    db.DB.Model(&Wxlog{}).Where("id = ?",action).Find(&wxlog)
    db.Logger.Errorln(db.DB.Error)
    return &wxlog
}
//查看所有
func(db *DB)ViewWxlogs()[]Wxlog{
    wxlogs := []Wxlog{}
    db.DB.Model(&Wxlog{}).Find(&wxlogs)
    db.Logger.Errorln(db.DB.Error)
    return wxlogs
}

//查看某类变更动作
func(db *DB)ViewChangeTypeActions(action_type string)[]Wxlog{
    wxlogs := []Wxlog{}
    db.DB.Model(&Wxlog{}).Where("change_type = ?",action_type).Find(&wxlogs)
    db.Logger.Errorln(db.DB.Error)
    return wxlogs
}
//查看某人的变更记录
func(db *DB)ViewUserActions(userid string)[]Wxlog{
    wxlogs := []Wxlog{}
    db.DB.Model(&Wxlog{}).Where("user_id = ?",userid).Find(&wxlogs)
    db.Logger.Errorln(db.DB.Error)
    return wxlogs
}
//查看某人某类变更记录
func(db *DB)ViewTypeUserActions(userid ,chang_type string)[]Wxlog{
    wxlogs := []Wxlog{}
    db.DB.Model(&Wxlog{}).Where("user_id = ? and change_type",userid,chang_type).Find(&wxlogs)
    db.Logger.Errorln(db.DB.Error)
    return wxlogs
}


//================================行业信息====================================
func(db *DB)AddTrade(t *Trade)error{
    db.DB.Create(t)
    db.Logger.Debugln("add a new trade member: ",*t)
    return db.DB.Error
}
func(db *DB)UpdateTrade(t *Trade)error{
    db.DB.Where("name = ?",t.Name).Updates(t)
    db.Logger.Debugln("update member: ",*t)
    return db.DB.Error
}

func(db *DB)ViewTrade(tname string)*Trade{
    t := Trade{}
    db.DB.Model(&Trade{}).Where("name = ?",tname).Find(&t)
    db.Logger.Errorln(db.DB.Error)
    return &t
}
//查看所有
func(db *DB)ViewTrades()[]Trade{
    ts := []Trade{}
    db.DB.Model(&Trade{}).Find(&ts)
    db.Logger.Errorln(db.DB.Error)
    return ts
}

func(db *DB)RMTrade(name string)error{
    t := Trade{}
    db.DB.Model(&Trade{}).Where("name = ?",name).Find(&t)
    // db.DB.Where("name = ?",name).Delete(&Trade{})
    db.DB.Delete(&Trade{},t.Model.ID)
    db.Logger.Debugln("delete a trade: ",name)
    return db.DB.Error
}



// ==============================Operation操作记录===========================
func(db *DB)AddOperation(t *Operation)error{
    db.DB.Create(t)
    db.Logger.Debugln("add a new operation member: ",*t)
    return db.DB.Error
}

func(db *DB)UpdateOperation(t *Operation)error{
    db.DB.Model(&Operation{}).Where("id = ?",t.Model.ID).Updates(t)
    db.Logger.Debugln("update member: ",*t)
    return db.DB.Error
}

//查看某人的操作记录
func(db *DB)ViewManOperation(tname string)[]Operation{
    t := []Operation{}
    db.DB.Model(&Operation{}).Where("fman = ?",tname).Find(&t)
    db.Logger.Errorln(db.DB.Error)
    return t
}

//查看某人某段时间的操作记录
func(db *DB)ViewManTimeOperation(tname,start,end string)[]Operation{
    t := []Operation{}
    db.DB.Model(&Operation{}).Where("fman = ? and BETWEEN ? and ?",tname,start,end).Find(&t)
    db.Logger.Errorln(db.DB.Error)
    return t
}

//查看某客户 某个用户 的操作记录
func(db *DB)ViewManCusOperations(tname,corpcode string)[]Operation{
    t := []Operation{}
    db.DB.Model(&Operation{}).Where("fman = ? and corpcode=?",tname,corpcode).Find(&t)
    db.Logger.Errorln(db.DB.Error)
    return t
}

//查看某客户 某个用户 某时间内的操作记录
func(db *DB)ViewManCusTimeOperations(tname,corpcode,start,end string)[]Operation{
    t := []Operation{}
    db.DB.Model(&Operation{}).Where("fman = ? and corpcode=? and unix_timestamp(created_at) BETWEEN ? and ?",tname,corpcode,start,end).Find(&t)
    db.Logger.Errorln(db.DB.Error)
    return t
}

//查看某客户的操作记录
func(db *DB)ViewCusOperations(corpcode string)[]Operation{
    t := []Operation{}
    db.DB.Model(&Operation{}).Where("corpcode=?",corpcode).Find(&t)
    db.Logger.Errorln(db.DB.Error)
    return t
}

//查看某客户的操作记录
func(db *DB)ViewCusTimeOperation(corpcode,start,end string)[]Operation{
    t := []Operation{}
    db.DB.Model(&Operation{}).Where("corpcode=? and unix_timestamp(created_at) BETWEEN ? and ?",corpcode,start,end).Find(&t)
    db.Logger.Errorln(db.DB.Error)
    return t
}

//查看所有操作记录
func(db *DB)ViewOperations()[]Operation{
    ts := []Operation{}
    db.DB.Model(&Operation{}).Find(&ts)
    db.Logger.Errorln(db.DB.Error)
    return ts
}


//查看某段时间的操作记录
func(db *DB)ViewTimeOperations(start,end string)[]Operation{
    t := []Operation{}
    db.DB.Model(&Operation{}).Where("unix_timestamp(created_at)  BETWEEN ? AND ?",start,end).Find(&t)
    db.Logger.Errorln(db.DB.Error)
    return t
}



func(db *DB)RMOperation(id string)error{
    db.DB.Delete(&Operation{},id)
    db.Logger.Debugln("delete a operation: ",id)
    return db.DB.Error
}
