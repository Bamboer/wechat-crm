package common
import(
  "fmt"
  "time"
  "sync"
  "strings"
)

var (
    Regions = map[string]string{}
    M sync.Mutex
)
//basic time uid generator
func GenUid()string{
  M.Lock()
  t := int(time.Now().UnixNano() / 1e6)
  time.Sleep(1 * time.Millisecond )
  defer M.Unlock()
  return  fmt.Sprintf("%X",t)
}


//区域判断
func RegionCheck(name string)string{
    // 1、东北（黑龙江省、吉林省、辽宁省）；
    Regions["东北"] = "HLJ,JL,LN"
    // 2、华东（上海市、江苏省、浙江省、安徽省、福建省、江西省、山东省、台湾省）；
    Regions["华东"] = "SD,JS,AH,ZJ,FJ,SH,JX,TW"
    // 3、华北（北京市、天津市、山西省、河北省、内蒙古自治区）；
    Regions["华北"] = "BJ,TJ,HEB,SX,NMG"
    // 4、华中（河南省、湖北省、湖南省）；
    Regions["华中"] ="HEN,HUB,HUN"
    // 5、华南（广东省、广西壮族自治区、海南省、香港特别行政区、澳门特别行政区）；
    Regions["华南"] = "GD,GX,HN,XG,AM"
    // 6、西南（四川省、贵州省、云南省、重庆市、西藏自治区）；
    Regions["西南"] = "SC,GZ,YN,CQ,XZ"
    // 7、西北（陕西省、甘肃省、青海省、宁夏回族自治区、新疆维吾尔自治区）
    Regions["西北"] = "NX,XJ,QH,SAX,GS"
    for k,v := range Regions{
        if strings.Index(v,name) != -1{
            return k  
        }
    }
    return ""    
}

//行业判断
func TradeCheck(corpname,scope string)string{
    if strings.Index(corpname,"金融")!= -1||strings.Index(corpname,"投资")!= -1||strings.Index(corpname,"证券")!= -1||strings.Index(corpname,"资产")!= -1||strings.Index(corpname,"银行")!= -1||strings.Index(corpname,"保险")!= -1{
        return "金融"
    }
    if strings.Index(corpname,"软件")!= -1{return "计算机软件信息"}
    if strings.Index(corpname,"健康") != -1||strings.Index(corpname,"医") != -1{return "医药医疗"}
    if strings.Index(corpname,"学")!= -1 ||strings.Index(corpname,"教育")!= -1{return "教育"}
    if strings.Index(corpname,"贸") != -1||strings.Index(scope,"零售")!= -1{return "批发零售"}
    if strings.Index(corpname,"科技")!= -1{
        if strings.Index(scope,"动漫")!= -1{return "游戏"}
        if strings.Index(scope,"AR")!= -1{
            return "AR"
        }
        if strings.Index(scope,"教育")!= -1{
            return "教育"
        }
        if strings.Index(scope,"医")!= -1{
            return "医药医疗"
        }
        return "计算机软硬件信息服务"
    }
    if strings.Index(scope,"生产")!= -1&&strings.Index(scope,"销售")!= -1{return "制造"}
    return "未知行业"
}