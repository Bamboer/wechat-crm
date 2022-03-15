package backend
import(
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
  "wxcrm/pkg/common/log"
)

type DB struct{
    DB	   *gorm.DB
    Logger *log.Logger
} 

func NewDB(host,user,password,dbname string,logger *log.Logger)*DB{
	logger.Debugln("database init info: host: ",host,"user: ",user," password: ",password," dbname: ",dbname)
	dsn := user + ":" + password + "@tcp(" + host + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{PrepareStmt:false})
	if err != nil {
        logger.Fatalln("database init error: ",err)
	}
  db.Set("gorm:table_options", `ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT="DMAI用户表"`).AutoMigrate(&DmaiUser{})
  db.Set("gorm:table_options", `ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT="CRM客户表"`).AutoMigrate(&Customer{})
	db.Set("gorm:table_options", `ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT="客户联系人表"`).AutoMigrate(&Contactor{})
	db.Set("gorm:table_options", `ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT="客户销售表"`).AutoMigrate(&CustomerUserprincipal{})
	db.Set("gorm:table_options", `ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT="客户联系人负责人表"`).AutoMigrate(&ContactUserprincipal{})

	db.Set("gorm:table_options", `ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT="客户失信表"`).AutoMigrate(&CustomerShiXin{})
	db.Set("gorm:table_options", `ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT="客户严重违法表"`).AutoMigrate(&CustomerSeriousIllegal{})
	db.Set("gorm:table_options", `ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT="客户税收违法表"`).AutoMigrate(&CustomerTaxIllegal{})
	db.Set("gorm:table_options", `ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT="客户受到行政处罚表"`).AutoMigrate(&CustomerAdminPenalty{})


	db.Set("gorm:table_options", `ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT="产品线"`).AutoMigrate(&Product{})
	db.Set("gorm:table_options", `ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT="项目表"`).AutoMigrate(&Project{})
	db.Set("gorm:table_options", `ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT="合同"`).AutoMigrate(&Agreement{})
  db.Set("gorm:table_options", `ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT="客户档案保存结果存储表"`).AutoMigrate(&YsCustomer{})

	db.Set("gorm:table_options", `ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT="客户跟进表"`).AutoMigrate(&FollowRecord{})
	db.Set("gorm:table_options", `ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT="企业微信变化日志表"`).AutoMigrate(&Wxlog{})
	db.Set("gorm:table_options", `ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT="操作日志表"`).AutoMigrate(&Operation{})
	db.Set("gorm:table_options", `ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT="客户行业表"`).AutoMigrate(&Trade{})
	return &DB{DB: db,Logger:logger} 
}


func (db *DB)InitAdmin(userids []string){
	for _,userid := range userids{
		user := db.ViewUser(userid)
		if user.Userid == ""{
			db.Logger.Infoln(userid," does not exists.")
			continue
		}
		if user.Manager == "true"{
			db.Logger.Infoln(userid," is exists and setted.")
		}else{
			db.Logger.Infoln(userid," exists and set manager.")
			db.DB.Model(&DmaiUser{}).Where("userid = ?",userid).Update("manager", "true")
		}
	}
}
