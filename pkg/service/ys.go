package service
import (
        "crypto/hmac"
        "crypto/sha256"
        "encoding/base64"
        "encoding/json"
        "io"
        "fmt"
        "bytes"
        "strings"
        "net/http"
        "net/url"
        "strconv"
        "time"
//        "wxcrm/pkg/common"
        "wxcrm/pkg/common/log"
)

type YS struct{
    AppKey string 
    AppSecret string 
    Logger    *log.Logger
    Token    string 
    SetTime  int64 
}

func NewYS(appkey , appsecret string,logger *log.Logger)*YS{
    return &YS{
        AppKey: appkey,
        AppSecret: appsecret,
        Logger: logger,
    }
}

func (y *YS)Post(data interface{},url string,result interface{})error{
    client := &http.Client{}
    token,err := y.GetYSToken()
    if err != nil{
        y.Logger.Errorln(err)
        return err
    }

    v,err := json.Marshal(data)
    if err != nil{
        y.Logger.Errorln(err)
        return err
    }
    req, err := http.NewRequest("POST",url +"?access_token=" + token , bytes.NewReader(v))
    if err != nil{
        y.Logger.Errorln(err)
        return err
    }
    req.Header.Add("Content-Type","application/json")
    resp, err := client.Do(req)
    if err != nil {
        y.Logger.Errorln(err)
        return  err
    }
    defer resp.Body.Close()
    if err = json.NewDecoder(resp.Body).Decode(result); err != nil {
        y.Logger.Errorln(err)
        return  err
    }
    return  nil 
}


//用友token获取
//URLEncode( Base64( HmacSHA256( parameterMap ) ) )
//appKey41832a3d2df94989b500da6a22268747timestamp1568098531823
//https://api.diwork.com/open-auth/selfAppAuth/getAccessToken?appKey=xxx&timestamp=xxx&signature=xxx

func (y *YS)GetYSToken() (string, error) {
        client := &http.Client{}
        var result YSToken
        // y.Logger.Debugln("AppSecret: ",y.AppSecret)
        key := []byte(y.AppSecret)
        t := strconv.Itoa(int(time.Now().UnixNano() / 1e6))
        mac := hmac.New(sha256.New, key)
        io.WriteString(mac, "appKey"+ y.AppKey + "timestamp"+t)
        expectedMAC := mac.Sum(nil)
        str := base64.StdEncoding.EncodeToString(expectedMAC)
        signature := url.QueryEscape(str)
        req, err := http.NewRequest("GET", "https://api.diwork.com/open-auth/selfAppAuth/getAccessToken?appKey=" + y.AppKey +"&timestamp="+t+"&signature=" + signature, nil)
        if err != nil {
                y.Logger.Errorln(err)
                return "", err
        }

        req.Header.Add("Content-Type","application/json")
        resp, err := client.Do(req)
        if err != nil {
            y.Logger.Errorln(err)
                return "", err
        }
        defer resp.Body.Close()
        if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
            y.Logger.Errorln(err)
                return "", err
        }
        if result.Code != "00000"{
            y.Logger.Errorln(result.Message)
            return "",fmt.Errorf(result.Message)
        }
        return result.Data.AccessToken, nil
}


//用友 单个客户档案上传
//POST https://api.diwork.com/yonbip/digitalModel/merchant/save
func (y *YS)MerchantPush(data *MerchantData)(*MerchantResult,error){
    client := &http.Client{}
    var result MerchantResult
    token,err := y.GetYSToken()
    if err != nil{
        y.Logger.Errorln(err)
        return &result,err
    }
    v,err := json.Marshal(data)
    if err != nil{
        y.Logger.Errorln(err)
        return &result,err
    }
    req, err := http.NewRequest("POST","https://api.diwork.com/yonbip/digitalModel/merchant/save?access_token=" + token , bytes.NewReader(v))
    if err != nil{
        y.Logger.Errorln(err)
        return &result,err
    }
    req.Header.Add("Content-Type","application/json")
    resp, err := client.Do(req)
    if err != nil {
            y.Logger.Errorln(err)
            return  &result,err
    }
    defer resp.Body.Close()
    if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
        y.Logger.Errorln(err)
            return  &result,err
    }
    if result.Code != "200"{
        y.Logger.Errorln(result.Message)
        return &result,fmt.Errorf(result.Message)
    }
    return  &result,nil
}



//用友 单个客户档案update上传
//POST https://api.diwork.com/yonbip/digitalModel/merchant/save
func (y *YS)MerchantUpdatePush(data *MerchantUpdateData)(*MerchantResult,error){
    client := &http.Client{}
    var result MerchantResult
    token,err := y.GetYSToken()
    if err != nil{
        y.Logger.Errorln(err)
        return &result,err
    }
    v,err := json.Marshal(data)
    if err != nil{
        y.Logger.Errorln(err)
        return &result,err
    }
    req, err := http.NewRequest("POST","https://api.diwork.com/yonbip/digitalModel/merchant/save?access_token=" + token , bytes.NewReader(v))
    if err != nil{
        y.Logger.Errorln(err)
        return &result,err
    }
    req.Header.Add("Content-Type","application/json")
    resp, err := client.Do(req)
    if err != nil {
            y.Logger.Errorln(err)
            return  &result,err
    }
    defer resp.Body.Close()
    if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
        y.Logger.Errorln(err)
            return  &result,err
    }
    if result.Code != "200"{
        y.Logger.Errorln(result.Message)
        return &result,fmt.Errorf(result.Message)
    }
    return  &result,nil
}



//用友客户档案列表查询
//POST https://api.diwork.com/yonbip/digitalModel/merchant/list
func (y *YS)MerchantList(data *MerchantListData)(*MerchantListResult,error){
    client := &http.Client{}
    result := MerchantListResult{}
    token,err := y.GetYSToken()
    if err != nil{
        y.Logger.Errorln(err)
        return nil,err
    }
    v,err := json.Marshal(data)
    if err != nil{
        y.Logger.Errorln(err)
        return nil,err
    }
    req, err := http.NewRequest("POST","https://api.diwork.com/yonbip/digitalModel/merchant/list?access_token=" + token , bytes.NewReader(v))
    if err != nil{
        y.Logger.Errorln(err)
        return nil,err
    }
    req.Header.Add("Content-Type","application/json")
    resp, err := client.Do(req)
    if err != nil {
        y.Logger.Errorln(err)
            return  nil,err
    }
    defer resp.Body.Close()
    if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
        y.Logger.Errorln(err)
            return  nil,err
    }
    if result.Code != "200"{
        y.Logger.Errorln(result.Message)
        return nil,fmt.Errorf(result.Message)
    }
    return  &result,nil 
}

//客户档案详情查询
// GET https://api.diwork.com/yonbip/digitalModel/merchant/detail
func(y *YS)MerchantDetail(id,merchantApplyRangeId string)(*MerchantDetailResult ,error){
        client := &http.Client{}
        result := &MerchantDetailResult{}
        token,err := y.GetYSToken()
        if err != nil{
            y.Logger.Errorln(err)
            return nil,err
        }
        req, err := http.NewRequest("GET", "https://api.diwork.com/yonbip/digitalModel/merchant/detail?access_token="+ token +"&id="+ id + "&merchantApplyRangeId=" + merchantApplyRangeId, nil)
        if err != nil {
                y.Logger.Errorln(err)
                return nil, err
        }

        req.Header.Add("Content-Type","application/json")
        resp, err := client.Do(req)
        if err != nil {
            y.Logger.Errorln(err)
                return nil, err
        }
        defer resp.Body.Close()
        if err = json.NewDecoder(resp.Body).Decode(result); err != nil {
            y.Logger.Errorln(err)
                return nil, err
        }
        if result.Code != "200"{
            y.Logger.Errorln(result.Message)
            return nil,fmt.Errorf(result.Message)
        }
        return result, nil
}


//组织单元列表查询
//POST https://api.diwork.com/yonbip/digitalModel/orgunit/querytree
func (y *YS)QueryTree(data *QueryTreeData)(*QueryTreeResult,error){
    client := &http.Client{}
    var result QueryTreeResult
    token,err := y.GetYSToken()
    if err != nil{
        y.Logger.Errorln(err)
        return nil,err
    }
    v,err := json.Marshal(data)
    if err != nil{
        y.Logger.Errorln(err)
        return nil,err
    }
    req, err := http.NewRequest("POST","https://api.diwork.com/yonbip/digitalModel/orgunit/querytree?access_token=" + token , bytes.NewReader(v))
    if err != nil{
        y.Logger.Errorln(err)
        return nil,err
    }
    req.Header.Add("Content-Type","application/json")
    resp, err := client.Do(req)
    if err != nil {
        y.Logger.Errorln(err)
            return  nil,err
    }
    defer resp.Body.Close()
    if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
        y.Logger.Errorln(err)
            return  nil,err
    }
    if result.Code != "200"{
        y.Logger.Errorln(result.Message)
        return nil,fmt.Errorf(result.Message)
    }
    return  &result,nil 
}

//业务单元同步
//GET https://api.diwork.com/yonbip/digitalModel/OrgUnitSync/doSync
func (y *YS)GetUnits() (*UnitResult, error) {
    client := &http.Client{}
    var result UnitResult
    token,err := y.GetYSToken()
    if err != nil{
        y.Logger.Errorln(err)
        return nil,err
    }

    req, err := http.NewRequest("GET","https://api.diwork.com/yonbip/digitalModel/OrgUnitSync/doSync?access_token=" + token , nil)
    if err != nil{
        y.Logger.Errorln(err)
        return nil,err
    }
    req.Header.Add("Content-Type","application/json")
    resp, err := client.Do(req)
    if err != nil {
        y.Logger.Errorln(err)
            return  nil,err
    }
    defer resp.Body.Close()
    if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
        y.Logger.Errorln(err)
            return  nil,err
    }
    if result.Code != "200"{
        y.Logger.Errorln(result.Message)
        return nil,fmt.Errorf(result.Message)
    }
    return  &result,nil 
}


//获取组织单元
type Ur struct{
   Text string   `json:"text"`       //组织单元名称
   Orgid string  `json:"orgid"`
}

func (y *YS)Componets()[]Ur{
    s := []Ur{}
    units,err := y.GetUnits()
    if err != nil{
        y.Logger.Errorln("Error: ",err)
    }
        
    for _,v := range units.Data.Recordlist{
           if strings.Index(v.Name,"暗物") != -1{
           s = append(s,Ur{Text: v.Name, Orgid: v.Orgid})
           }
    }
    y.Logger.Debugln("Organization units: ",s)
    return s 
}


//客户级别查询
//POST https://api.diwork.com/yonbip/digitalModel/cuslevel/list
func (y *YS)GetCuslevel(data *CusLevelData)(*CusLevelResult,error){
    client := &http.Client{}
    var result CusLevelResult
    token,err := y.GetYSToken()
    if err != nil{
        y.Logger.Errorln(err)
        return &result,err
    }
    v,err := json.Marshal(data)
    if err != nil{
        y.Logger.Errorln(err)
        return &result,err
    }
    req, err := http.NewRequest("POST","https://api.diwork.com/yonbip/digitalModel/cuslevel/list?access_token=" + token , bytes.NewReader(v))
    if err != nil{
        y.Logger.Errorln(err)
        return &result,err
    }
    req.Header.Add("Content-Type","application/json")
    resp, err := client.Do(req)
    if err != nil {
        y.Logger.Errorln(err)
            return  &result,err
    }
    defer resp.Body.Close()
    if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
        y.Logger.Errorln(err)
            return  &result,err
    }
    if result.Code != "200"{
        y.Logger.Errorln(result.Message)
        return &result,fmt.Errorf(result.Message)
    }
    return  &result,nil 
}



//项目列表查询
//POST https://api.diwork.com/yonbip/digitalModel/project/list
func (y *YS)GetProjectList(data *YSProjectListData)(*YSProjectListResult,error){
    client := &http.Client{}
    var result YSProjectListResult
    token,err := y.GetYSToken()
    if err != nil{
        y.Logger.Errorln(err)
        return &result,err
    }
    v,err := json.Marshal(data)
    if err != nil{
        y.Logger.Errorln(err)
        return &result,err
    }
    req, err := http.NewRequest("POST","https://api.diwork.com/yonbip/digitalModel/project/list?access_token=" + token , bytes.NewReader(v))
    if err != nil{
        y.Logger.Errorln(err)
        return &result,err
    }
    req.Header.Add("Content-Type","application/json")
    resp, err := client.Do(req)
    if err != nil {
        y.Logger.Errorln(err)
            return  &result,err
    }
    defer resp.Body.Close()
    if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
        y.Logger.Errorln(err)
            return  &result,err
    }
    if result.Code != "200"{
        y.Logger.Errorln(result.Message)
        return &result,fmt.Errorf(result.Message)
    }
    return  &result,nil 
}



//员工列表查询
//POST https://api.diwork.com/yonbip/digitalModel/staff/list
func (y *YS)GetStaffList(data *CommenInput)(*StaffListResult,error){
    client := &http.Client{}
    var result StaffListResult
    token,err := y.GetYSToken()
    if err != nil{
        y.Logger.Errorln(err)
        return &result,err
    }
    v,err := json.Marshal(data)
    if err != nil{
        y.Logger.Errorln(err)
        return &result,err
    }
    req, err := http.NewRequest("POST","https://api.diwork.com/yonbip/digitalModel/staff/list?access_token=" + token , bytes.NewReader(v))
    if err != nil{
        y.Logger.Errorln(err)
        return &result,err
    }
    req.Header.Add("Content-Type","application/json")
    resp, err := client.Do(req)
    if err != nil {
        y.Logger.Errorln(err)
            return  &result,err
    }
    defer resp.Body.Close()
    if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
        y.Logger.Errorln(err)
            return  &result,err
    }
    if result.Code != "200"{
        y.Logger.Errorln(result.Message)
        return &result,fmt.Errorf(result.Message)
    }
    return  &result,nil 
}



//员工详情查询
//GET https://api.diwork.com/yonbip/digitalModel/staff/detail
func (y *YS)GetStaffDetail(id,code string) (*StaffDetailResult, error) {
    client := &http.Client{}
    var result StaffDetailResult
    token,err := y.GetYSToken()
    if err != nil{
        y.Logger.Errorln(err)
        return nil,err
    }

    req, err := http.NewRequest("GET","https://api.diwork.com/yonbip/digitalModel/OrgUnitSync/doSync?access_token=" + token+"&id="+id+"&code="+code , nil)
    if err != nil{
        y.Logger.Errorln(err)
        return nil,err
    }
    req.Header.Add("Content-Type","application/json")
    resp, err := client.Do(req)
    if err != nil {
        y.Logger.Errorln(err)
            return  nil,err
    }
    defer resp.Body.Close()
    if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
        y.Logger.Errorln(err)
            return  nil,err
    }
    if result.Code != "200"{
        y.Logger.Errorln(result.Message)
        return nil,fmt.Errorf(result.Message)
    }
    return  &result,nil 
}


type CheckResult struct {
    ID  string 
    // Code string 
    DeptID string 
    DpName string 
}
//check 用户在erp中的信息
func (y *YS)CheckERPUserInfo(name string)(name1,id,dptname,dptid string,e1 error){
    y.Logger.Debugln("查询ERP用户名为: ",name)
    num := 2 
    result := []CheckResult{}
    input := CommenInput{PageIndex:1,PageSize: 1000}
    data,err := y.GetStaffList(&input)
    if err != nil{
        y.Logger.Errorln(err)
        return "","","","",err 
    }
    for _,v := range data.Data.RecordList{
        if name == v.Name {
            y.Logger.Debugln("Found person: ",v.Name)
            result = append(result,CheckResult{ID: v.ID, DpName: v.DeptIDName, DeptID:v.DeptID})
        }
    }
    for num <= data.Data.PageCount{
        input := CommenInput{PageIndex:num,PageSize: 1000}
        data,err := y.GetStaffList(&input)
        if err != nil{
            y.Logger.Errorln(err)
            return "","","","",err 
        }
        for _,v := range data.Data.RecordList{
            if name == v.Name {
                y.Logger.Debugln("Found person: ",v.Name)
                result = append(result,CheckResult{ID:v.ID, DeptID:v.DeptID, DpName: v.DeptIDName})
            }
        }
        num++
    }

    if len(result) == 0{
        return "","","","",fmt.Errorf("Not found")
    }
    for _, v := range result{
        if v.DpName == "销售" || v.DpName == "产品"|| v.DpName == "运营"{
            return name,v.ID,v.DpName,v.DeptID,nil 
        }
    }

    return name,result[len(result)-1].ID,result[len(result)-1].DpName,result[len(result)-1].DeptID,nil 
}














