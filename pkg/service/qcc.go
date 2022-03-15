package service
import (
        "time"
        "strconv"
        "crypto/md5"
        "fmt"
        "io"
        "net/http"
        "encoding/json"
        "wxcrm/pkg/common/log"
)

type QCC struct{
	AppKey string
	SecretKey string
	Logger    *log.Logger 
}

func NewQCC(appkey,secretkey string,logger *log.Logger)*QCC{
     return &QCC{
     	AppKey: appkey, 
     	SecretKey: secretkey, 
     	Logger: logger,
     }
}


//获取基础信用报告
//GET http://api.qichacha.com/ReportBase/GetReportInfo?key=AppKey&orderNo={orderNo}
func (q *QCC)QGetReportBase(orderNo string)(*ReportBaseResult,error){
        var result ReportBaseResult 
        reqInterNme := "http://api.qichacha.com/ReportBase/GetReportInfo"
        paramStr := "orderNo="+ orderNo
        url := reqInterNme + "?key=" + q.AppKey + "&" + paramStr
        err := q.get(url,&result)
        if err != nil{
                q.Logger.Errorln(err)
                return nil,err
        }
        return &result,nil 
}



//基础信用报告下单
//POST http://api.qichacha.com/ReportBase/CreateReport?key=AppKey&searchKey={searchKey}
func (q *QCC)QReportBaseEmit(keyword string)(*ReportBase,error){
        var result ReportBase  
        reqInterNme := "http://api.qichacha.com/ReportBase/CreateReport"
        paramStr := "searchKey="+ keyword +"&reportFormat=1"
        url := reqInterNme + "?key=" + q.AppKey + "&" + paramStr
        err := q.BasicPost(url,&result)
        if err != nil{
                q.Logger.Errorln(err)
                return nil,err
        }
        return &result,nil 
}

//行政处罚核查
// http://api.qichacha.com/AdminPenaltyCheck/GetList?key=AppKey&searchKey=XXXXXX
func (q *QCC)QAdminPenaltyChecker(keyword string)(*AdminPenalty,error){
        var result AdminPenalty 
        reqInterNme := "http://api.qichacha.com/AdminPenaltyCheck/GetList"
        paramStr := "searchKey="+ keyword +"&pageSize=20"
        url := reqInterNme + "?key=" + q.AppKey + "&" + paramStr
        err := q.get(url,&result)
        if err != nil{
                q.Logger.Errorln(err)
                return nil,err
        }
        return &result,nil 
}



//税收违法核查
// http://api.qichacha.com/TaxIllegalCheck/GetList?key=AppKey&searchKey=XXXXXX
func (q *QCC)QTaxIllegalChecker(keyword string)(*TaxIllegalCheck,error){
        var result TaxIllegalCheck  
        reqInterNme := "http://api.qichacha.com/TaxIllegalCheck/GetList"
        paramStr := "searchKey="+ keyword +"&pageSize=20"
        url := reqInterNme + "?key=" + q.AppKey + "&" + paramStr
        err := q.get(url,&result)
        if err != nil{
                q.Logger.Errorln(err)
                return nil,err
        }
        return &result,nil 
}




//严重违法查询
//GET http://api.qichacha.com/SeriousIllegalCheck/GetList?key=AppKey&searchKey=XXXXXX 
func (q *QCC)QSeriousIllegalChecker(keyword string)(*SeriousIllegalcheck,error){
        var result SeriousIllegalcheck  
        reqInterNme := "http://api.qichacha.com/SeriousIllegalCheck/GetList"
        paramStr := "searchKey="+ keyword 
        url := reqInterNme + "?key=" + q.AppKey + "&" + paramStr
        err := q.get(url,&result)
        if err != nil{
                q.Logger.Errorln(err)
                return nil,err
        }
        return &result,nil 
}



//失信查询
//http://api.qichacha.com/ShixinCheck/GetList?key=AppKey&searchKey=乐视网信息技术（北京）股份有限公司
func (q *QCC)QShiXinChecker(keyword string)(*ShiXinCheck,error){
        var result ShiXinCheck  
        reqInterNme := "http://api.qichacha.com/ShixinCheck/GetList"
        paramStr := "searchKey="+ keyword +"&pageSize=20"
        url := reqInterNme + "?key=" + q.AppKey + "&" + paramStr
        err := q.get(url,&result)
        if err != nil{
                q.Logger.Errorln(err)
                return nil,err
        }
        return &result,nil 
}

//企查查工商数据查询  多维度
//http://api.qichacha.com/ECIV4/SearchWide?key=AppKey&keyword=小桔科技
func (q *QCC)QSearchWide(keyword string)(*SearchWide,error){
	var result SearchWide
        reqInterNme := "http://api.qichacha.com/ECIV4/Search"
        paramStr := "keyword="+ keyword
        url := reqInterNme + "?key=" + q.AppKey + "&" + paramStr
        err := q.get(url,&result)
        if err != nil{
        	q.Logger.Errorln(err)
        	return nil,err
        }
        return &result,nil 
}

//企查查工商详情 查询
//http://api.qichacha.com/ECIV4/GetBasicDetailsByName?key=AppKey&keyword=北京小桔科技有限公司
func (q *QCC)QGetBasicDetailsByName(keyword string)(*GetBasicDetailsByName,error){
	var result GetBasicDetailsByName
        reqInterNme := "http://api.qichacha.com/ECIV4/GetBasicDetailsByName"
        paramStr := "keyword="+ keyword
        url := reqInterNme + "?key=" + q.AppKey + "&" + paramStr
        err := q.get(url,&result)
        if err != nil{
        	q.Logger.Errorln(err)
        	return nil,err
        }
        return &result,nil 
}


//企查查模糊搜索
//http://api.qichacha.com/FuzzySearch/GetList?key=AppKey&searchKey=北京小桔科技有限公司
func (q *QCC)QFuzzySearch(keyword string)(*FuzzySearch,error){
	var result FuzzySearch
        reqInterNme := "http://api.qichacha.com/FuzzySearch/GetList"
        paramStr := "searchkey="+ keyword + "&pageSize=10"
        url := reqInterNme + "?key=" + q.AppKey + "&" + paramStr
        err := q.get(url,&result)
        if err != nil{
        	q.Logger.Errorln(err)
        	return nil,err
        }
        return &result,nil 	
}


func (q *QCC)get(url string, v interface{}) error {
        h := md5.New()
        timespan := strconv.Itoa(int(time.Now().Unix()))
        io.WriteString(h, q.AppKey)
        io.WriteString(h, timespan)
        io.WriteString(h, q.SecretKey)
        token := fmt.Sprintf("%X",h.Sum(nil))

        client := &http.Client{}
        req, err := http.NewRequest("GET", url, nil)
        if err != nil {
                q.Logger.Errorln(err)
                return err
        }
        req.Header.Add("Token", token)
        req.Header.Add("Timespan", timespan)
        resp, err := client.Do(req)
        if err != nil {
        	q.Logger.Errorln(err)
                return err
        }
        defer resp.Body.Close()
        if err = json.NewDecoder(resp.Body).Decode(v); err != nil {
                q.Logger.Errorln(err)
                return err
        }
        return nil
}


func (q *QCC)BasicPost(url string, v interface{}) error {
        h := md5.New()
        timespan := strconv.Itoa(int(time.Now().Unix()))
        io.WriteString(h, q.AppKey)
        io.WriteString(h, timespan)
        io.WriteString(h, q.SecretKey)
        token := fmt.Sprintf("%X",h.Sum(nil))

        client := &http.Client{}
        req, err := http.NewRequest("POST", url, nil)
        if err != nil {
                q.Logger.Errorln(err)
                return err
        }
        req.Header.Add("Token", token)
        req.Header.Add("Timespan", timespan)
        resp, err := client.Do(req)
        if err != nil {
                q.Logger.Errorln(err)
                return err
        }
        defer resp.Body.Close()
        if err = json.NewDecoder(resp.Body).Decode(v); err != nil {
                q.Logger.Errorln(err)
                return err
        }
        return nil
}



