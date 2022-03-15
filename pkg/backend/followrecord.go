package backend


//--------------------------------------followerThroughs-------------------------------------
//跟进记录
func(db *DB)ViewCustomerFollowRecords(corpcode string)[]FollowRecord{
   follothroughs := []FollowRecord{} 
   db.DB.Model(&FollowRecord{}).Where("corp_code = ? ",corpcode).Find(&follothroughs)
   // db.Logger.Debugln(corpcode, " view FollowRecords: ",follothroughs)
   return follothroughs
}
//最近一条跟进记录
func(db *DB)ViewSingleCustomerFollowRecord(corpcode string)*FollowRecord{
   follothroughs := FollowRecord{} 
   db.DB.Where("corp_code = ?",corpcode).Order("created_at desc").Limit(1).Find(&follothroughs)
   db.Logger.Debugln(corpcode, " view FollowRecords: ",follothroughs)
   return &follothroughs
}


func(db *DB)UpdateFollowRecord(follothrough *FollowRecord)error{
   db.DB.Model(&FollowRecord{}).Where("id = ?",follothrough.Model.ID).Updates(follothrough)
   db.Logger.Debugln("update followlog :",follothrough)
   return db.DB.Error
}

func(db *DB)AddFollowRecord(follothrough *FollowRecord)error{
   db.DB.Create(follothrough)
   db.Logger.Debugln("add a new FollowRecord: ",follothrough)
   return db.DB.Error
}

func(db *DB)DelFollowRecord(id string)error{
	// db.DB.Where("id = ?",id).Delete(&FollowRecord{})
   db.DB.Delete(&FollowRecord{},id)
	db.Logger.Debugln("delete a follow record: ",id)

	return db.DB.Error
}

//查看某用户跟进记录
func(db *DB)ViewUserFollowRecords(userid string)[]FollowRecord{
   follothroughs := []FollowRecord{} 
   db.DB.Model(&FollowRecord{}).Where("updateby = ? ",userid).Find(&follothroughs)
   db.Logger.Debugln(userid, " view FollowRecords: ",follothroughs)
   return follothroughs
}
//查看一条跟进记录
func(db *DB)ViewSingleFollowRecord(id string)*FollowRecord{
   follothrough := FollowRecord{} 
   db.DB.Model(&FollowRecord{}).Where("id = ? ", id).Find(&follothrough)
   db.Logger.Debugln(id, " view FollowRecord: ",follothrough)
   return &follothrough
}


//查看特定客户的项目
func(db *DB)ViewCustomerRecords(corpcode string)[]FollowRecord{
   follothroughs := []FollowRecord{} 
   db.DB.Model(&FollowRecord{}).Distinct("projectid").Where("corp_code=?",corpcode).Find(&follothroughs)
   db.Logger.Debugln(corpcode, " view FollowRecords: ",follothroughs)
   return follothroughs
}


//查看某个用户的跟进项目
func(db *DB)ViewUserProjects(userid string)[]FollowRecord{
   follothroughs := []FollowRecord{} 
   db.DB.Model(&FollowRecord{}).Distinct("corp_code","projectid","projectname").Where("updateby = ? ",userid).Find(&follothroughs)
   db.Logger.Debugln(userid, " view FollowRecords: ",follothroughs)
   return follothroughs
}



//查看某个项目的跟进记录
func(db *DB)ViewProjectsFollowRecords(corpcode,project string)[]FollowRecord{
   follothroughs := []FollowRecord{} 
   db.DB.Model(&FollowRecord{}).Where("corp_code = ? and projectid = ? ",corpcode,project).Find(&follothroughs)
   db.Logger.Debugln(corpcode, " view FollowRecords: ",follothroughs)
   return follothroughs
}

//查看某个项目某时间内的跟进记录
func(db *DB)ViewProjectTimeFollowRecords(corpcode,project,start,end string)[]FollowRecord{
   follothroughs := []FollowRecord{} 
   db.DB.Model(&FollowRecord{}).Where("corp_code = ? and projectid = ? and unix_timestamp(updated_at) BETWEEN ? and ?",corpcode,project,start,end).Find(&follothroughs)
   db.Logger.Debugln(corpcode, " view FollowRecords: ",follothroughs)
   return follothroughs
}


//查看所有客户跟进项目
func(db *DB)ViewProjects()[]FollowRecord {
   follothroughs := []FollowRecord{} 
   db.DB.Model(&FollowRecord{}).Distinct("corp_code", "projectid","projectname").Find(&follothroughs)
   db.Logger.Debugln(" view FollowRecords: ",follothroughs)
   return follothroughs
}


//查看特定客户项目的时间规划及预算
func(db *DB)ViewTimeBudget(corpcode,projectid string)[]FollowRecord{
   follothroughs := []FollowRecord{} 
   db.DB.Model(&FollowRecord{}).Distinct("corp_code", "projectid","timeplan","budget").Where("corp_code=? and projectid=?",corpcode,projectid).Find(&follothroughs)
   // db.DB.Model(&FollowRecord{}).Select("projectid,timeplan,budget").Group("projectid").Find(&follothroughs)
   db.Logger.Debugln(corpcode, " view FollowRecords: ",follothroughs)
   return follothroughs
}














