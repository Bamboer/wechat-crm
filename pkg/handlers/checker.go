package handlers
import(

  "net/http"
  "strconv"
  "wxcrm/pkg/backend"
  "github.com/gin-gonic/gin"
)

//Here are Checker,Product,project,agreement

func (h *HandlerVar)Checker(c *gin.Context) {
	check := c.Query("check")
	id := c.Query("id")
	name := c.Query("name")
	// id is contactid 
	if check == "contact" && id != ""{
		if contact := h.DB.ViewWechatContactor(id); len(contact) == 0 {
			h.Logger.Debugln("not found contactor :",id)
			c.JSON(http.StatusOK, gin.H{"status": 200})
		} else {
			h.Logger.Debugln("found contactor :",id)
			c.JSON(http.StatusOK, gin.H{"status": 407,"ErrMsg":"found contactor "+ id})
		}
		c.AbortWithStatus(http.StatusOK)
	}
	//customer name check
	if check == "customer" && name != "" {
		if customer := h.DB.CheckCustomer(name); customer.CorpName == "" {
			h.Logger.Debugln("not found customer :", name)
			c.JSON(http.StatusOK, gin.H{"status": 200})
		} else {
			h.Logger.Debugln("Found customer :",name)
			cuperson := h.DB.ViewCustomerUserprincipal(customer.CorpCode)
			if len(cuperson) > 0{
				c.JSON(http.StatusOK, gin.H{"status": 407,"ErrMsg": "客户已存在！","Cuperson": cuperson[len(cuperson)-1].MerchandiserName})
			}else{
				c.JSON(http.StatusOK, gin.H{"status": 407,"ErrMsg": "客户已存在！","Cuperson": ""})
			}

		}
		c.AbortWithStatus(http.StatusOK)
	}
	//product name check
	if check == "product" && name != "" {
		if product := h.DB.CheckProduct(name); product.Name == "" {
			h.Logger.Debugln("not found product :",name)
			c.JSON(http.StatusOK, gin.H{"status": 200})
		} else {
			h.Logger.Debugln("found product :",name)
			c.JSON(http.StatusOK, gin.H{"status": 407,"ErrMsg":"found product "+ name})
		}
		c.AbortWithStatus(http.StatusOK)
	}
	//project name  check
	if check == "project" && name != "" {
		if project := h.DB.CheckProject(name); project.Name == "" {
			h.Logger.Debugln("not found project :",name)
			c.JSON(http.StatusOK, gin.H{"status": 200})
		}else {
			h.Logger.Debugln("found project :",name)
			c.JSON(http.StatusOK, gin.H{"status": 407,"ErrMsg":"found project "+ name})
		}
		c.AbortWithStatus(http.StatusOK)
	}
	//trade name  check
	// if check == "trade" && name != "" {
	// 	if trade := h.DB.ViewTrade(name); trade.Name == "" {
	// 		h.Logger.Debugln("not found trade :",name)
	// 		c.JSON(http.StatusOK, gin.H{"status": 200})

	// 	} else {
	// 		h.Logger.Debugln("found trade :",name)
	// 		c.JSON(http.StatusOK, gin.H{"status": 400,"ErrMsg":"found trade "+name})
	// 	}
	// 	c.AbortWithStatus(http.StatusOK)
	// }
}


func (h *HandlerVar)Product(c *gin.Context){
	var recv backend.Product
	action := c.Query("action")
	id := c.Query("id")
  c.ShouldBindJSON(&recv)
  user := c.Query("wxcrm_username")
	if action == "add" {
		if err := h.DB.AddProduct(&recv);err != nil{
			h.Logger.Errorln(err)
			c.JSON(http.StatusOK,gin.H{"status": 400,"ErrMsg": err})
		}else{
			h.Logger.Debugln("add product :",recv.Name)
			h.AddOperation(user,"","增加","产品",recv.Name)
			c.JSON(http.StatusOK,gin.H{"status": 200})
		}
		c.AbortWithStatus(http.StatusOK)
	}
	if action == "delete" && id != ""{
		productinfo := h.DB.ViewProduct(id)
		if err := h.DB.DelProduct(id);err == nil{
			h.Logger.Debugln("delete product :",id)
			h.AddOperation(user,"","删除","产品",productinfo.Name)
			c.JSON(http.StatusOK,gin.H{"status": 200})
		}else{
			c.JSON(http.StatusOK,gin.H{"status": 400,"ErrMsg": err})
		}
		c.AbortWithStatus(http.StatusOK)
	}
	if action == "update"{
		if err := h.DB.UpdateProduct(&recv);err ==nil{
			h.Logger.Debugln("update product :",recv.Name)
			h.AddOperation(user,"","更新","产品",recv.Name)
			c.JSON(http.StatusOK,gin.H{"status": 200})
		}else{
			c.JSON(http.StatusOK,gin.H{"status": 400,"ErrMsg": err})
		}
		c.AbortWithStatus(http.StatusOK)
	}	
}


func (h *HandlerVar)Project(c *gin.Context){
	var recv backend.Project 
	action := c.Query("action")
	id := c.Query("id")
	corpcode :=c.Query("corpcode")
	cookie, _ := c.Cookie("wxcrm_username")
	c.ShouldBindJSON(&recv)
	if action == "add"{
		if recv.CorpCode == ""{
			c.JSON(http.StatusOK,gin.H{"status": 400,"ErrMsg": "CorpCode为空!"})
			c.AbortWithStatus(http.StatusOK)
			return
		}
		dmaiuser := h.DB.ViewUser(cookie)
		recv.Createby = cookie
		recv.CreatebyName = dmaiuser.Name 
		if err := h.DB.AddProject(&recv);err != nil{
			h.Logger.Errorln(err)
			c.JSON(http.StatusOK,gin.H{"status": 400,"ErrMsg": err})
		}else{
			h.Logger.Debugln("add project :",recv.Name)
			h.AddOperation(cookie,corpcode,"增加","项目",recv.Name)
			c.JSON(http.StatusOK,gin.H{"status": 200})
		}
		c.AbortWithStatus(http.StatusOK)
	}
	if action == "delete" && id != ""{
		projectinfo := h.DB.ViewSingleProject(id)
		if err := h.DB.DelProject(id);err == nil{
			h.Logger.Debugln("delete project :",id)
			h.AddOperation(cookie,corpcode,"删除","项目",projectinfo.Name)
			c.JSON(http.StatusOK,gin.H{"status": 200})
		}else{
			c.JSON(http.StatusOK,gin.H{"status": 400,"ErrMsg": err})
		}
		c.AbortWithStatus(http.StatusOK)
	}
	if action == "update"{
		tid,_ := strconv.Atoi(id)
		recv.Model.ID = uint(tid)
		dmaiuser := h.DB.ViewUser(cookie)
		recv.Updateby = cookie
		recv.UpdatebyName = dmaiuser.Name 
		if err := h.DB.UpdateProject(&recv);err ==nil{
			h.Logger.Debugln("update project :",recv.Name)
			h.AddOperation(cookie,corpcode,"更新","项目",recv.Name)
			c.JSON(http.StatusOK,gin.H{"status": 200})
		}else{
			c.JSON(http.StatusOK,gin.H{"status": 400,"ErrMsg": err})
		}
		c.AbortWithStatus(http.StatusOK)
	}
	if action == "info" && id != ""{
		project := h.DB.ViewSingleProject(id)
		h.Logger.Debugln("project corpcode:",project.CorpCode)
		c.JSON(http.StatusOK,gin.H{"Project":project,"status": 200})
		c.AbortWithStatus(http.StatusOK)
	}
	if action == "list" && corpcode != "" {
		projects := h.DB.ViewProject(corpcode)
		h.Logger.Debugln("projects:",projects)
		c.JSON(http.StatusOK,gin.H{"Projects":projects,"status": 200})
		c.AbortWithStatus(http.StatusOK)
	}
}

func (h *HandlerVar)Agreement(c *gin.Context){
	var recv backend.Agreement
	action := c.Query("action")
	id := c.Query("id")
	cookie, _ := c.Cookie("wxcrm_username")
	corpcode := c.Query("corpcode")
	c.ShouldBindJSON(&recv)
	if action == "add"{
		if recv.CorpCode == ""{
			c.JSON(http.StatusOK,gin.H{"status": 400,"ErrMsg": "CorpCode为空!"})
			c.AbortWithStatus(http.StatusOK)
			return
		}
		if err := h.DB.AddAgreement(&recv);err != nil{
			h.Logger.Errorln(err)
			c.JSON(http.StatusOK,gin.H{"status": 400,"ErrMsg": err})
		}else{
			h.Logger.Debugln("add agreement :",recv.Name)
			h.AddOperation(cookie,corpcode,"增加","合同",recv.Name)
			c.JSON(http.StatusOK,gin.H{"status": 200})
		}
		c.AbortWithStatus(http.StatusOK)
	}
	if action == "delete" && id != ""{
		agreementinfo := h.DB.ViewSingleAgreement(id)
		if err := h.DB.DelAgreement(id);err == nil{
			h.Logger.Debugln("delete agreement :",id)
			h.AddOperation(cookie,corpcode,"删除","合同",agreementinfo.Name)
			c.JSON(http.StatusOK,gin.H{"status": 200})
		}else{
			c.JSON(http.StatusOK,gin.H{"status": 400,"ErrMsg": err})
		}
		c.AbortWithStatus(http.StatusOK)
	}
	if action == "update" && id != ""{
		if recv.CorpCode == ""{
			c.JSON(http.StatusOK,gin.H{"status": 400,"ErrMsg": "CorpCode为空!"})
			c.AbortWithStatus(http.StatusOK)
			return
		}
    tid,_ := strconv.Atoi(id)
		recv.Model.ID = uint(tid)
		if err := h.DB.UpdateAgreement(&recv);err ==nil{
			h.Logger.Debugln("update agreement :",recv.Name)
			h.AddOperation(cookie,corpcode,"更新","合同",recv.Name)
			c.JSON(http.StatusOK,gin.H{"status": 200})
		}else{
			c.JSON(http.StatusOK,gin.H{"status": 400,"ErrMsg": err})
		}
		c.AbortWithStatus(http.StatusOK)
	}
	if action == "info" && id != ""{
		agreement := h.DB.ViewSingleAgreement(id)
		h.Logger.Debugln("agreement corpcode:",agreement.CorpCode)
		c.JSON(http.StatusOK,gin.H{"Agreement": agreement,"status": 200})
		c.AbortWithStatus(http.StatusOK)
	}
	if action == "list" && corpcode != ""{
		agreements := h.DB.ViewAgreement(corpcode)
		h.Logger.Debugln("agreement corpcode:", corpcode)
		c.JSON(http.StatusOK,gin.H{"Agreements": agreements,"status": 200})
		c.AbortWithStatus(http.StatusOK)
	}
}

func (h *HandlerVar)Operation(c *gin.Context){
	user := c.Query("user")
	corpcode := c.Query("corpcode")
	from_time := c.Query("from_time")
	end_time := c.Query("end_time")

  if user == "" && corpcode == "" && from_time == "" && end_time == ""{
  	result := h.DB.ViewOperations()
  	c.JSON(http.StatusOK,gin.H{"Operations": result,"status":200})
  	c.AbortWithStatus(http.StatusOK)
  	return
  }
  if user != "" && corpcode == "" && from_time == "" && end_time == ""{
  	result := h.DB.ViewManOperation(user)
  	c.JSON(http.StatusOK,gin.H{"Operations": result,"status":200})
  	c.AbortWithStatus(http.StatusOK)
  	return
  }
  if user == "" && corpcode != "" && from_time == "" && end_time == ""{
  	result := h.DB.ViewCusOperations(corpcode)
  	c.JSON(http.StatusOK,gin.H{"Operations": result,"status":200})
  	c.AbortWithStatus(http.StatusOK)
  	return
  }
  if user == "" && corpcode == "" && from_time != "" && end_time != ""{
  	result := h.DB.ViewTimeOperations(from_time,end_time)
  	c.JSON(http.StatusOK,gin.H{"Operations": result,"status":200})
  	c.AbortWithStatus(http.StatusOK)
  	return
  }
  if user == "" && corpcode != "" && from_time != "" && end_time != ""{
  	result := h.DB.ViewCusTimeOperation(corpcode,from_time,end_time)
  	c.JSON(http.StatusOK,gin.H{"Operations": result,"status":200})
  	c.AbortWithStatus(http.StatusOK)
  	return
  }
  if user != "" && corpcode == "" && from_time != "" && end_time != ""{
  	result := h.DB.ViewManTimeOperation(user,from_time,end_time)
  	c.JSON(http.StatusOK,gin.H{"Operations": result,"status":200})
  	c.AbortWithStatus(http.StatusOK)
  	return
  }
  if user != "" && corpcode != "" && from_time != "" && end_time != ""{
  	result := h.DB.ViewManCusTimeOperations(user,corpcode,from_time,end_time)
  	c.JSON(http.StatusOK,gin.H{"Operations": result,"status":200})
  	c.AbortWithStatus(http.StatusOK)
  	return
  }
  if user != "" && corpcode != "" && from_time == "" && end_time == ""{
  	result := h.DB.ViewManCusOperations(user,corpcode)
  	c.JSON(http.StatusOK,gin.H{"Operations": result,"status":200})
  	c.AbortWithStatus(http.StatusOK)
  	return
  }
  c.JSON(http.StatusOK,gin.H{"status":400,"ErrMsg":"参数错误！"})
  c.AbortWithStatus(http.StatusOK)
}


type ProjectsInfo struct{
    CorpCode  string `json:"corpcode"`
    Projectid string `json:"projectid"`
    Timeplan  string `json:"timeplan"`
    Budget    string `json:"budget"`
}

type ProjectsInfos []ProjectsInfo

func (h *HandlerVar)CustomerProjectsInfo(c *gin.Context){
	result := ProjectsInfos{}
	corpcode :=c.Query("corpcode")
	if corpcode == ""{
		c.JSON(http.StatusOK,gin.H{"status": 400,"ErrMsg": "CorpCode为空!"})
		c.AbortWithStatus(http.StatusOK)
		return
	}
  projects := h.DB.ViewCustomerRecords(corpcode)
  for _,v := range projects{
  	timebudgets := h.DB.ViewTimeBudget(corpcode,v.Projectid)
  	result = append(result,ProjectsInfo{CorpCode: timebudgets[len(timebudgets)-1].CorpCode,Projectid:timebudgets[len(timebudgets)-1].Projectid,Timeplan:timebudgets[len(timebudgets)-1].Timeplan,Budget:timebudgets[len(timebudgets)-1].Budget})
  }
  c.JSON(http.StatusOK,gin.H{"status":200,"ProjectInfos":result})
  c.AbortWithStatus(http.StatusOK)
  return
}














