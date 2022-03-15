package service


//企业微信access token 返回data struct
type WXRT struct{
  Errcode      int          `json:"errcode,required"`
  Errmsg       string       `json:"errmsq,required"`
  Access_token string       `json:"access_token,required"`
  Expires_in   int          `json:"expires_in,required"`
}

//企业微信当前使用用户ID struct
type AccessUser struct{
  Errcode int          `json:"errcode,required"`
  Errmsg  string       `json:"errmsq,required"`
  UserId  string       `json:"userid,required"`
  DeviceId string      `json:"deviceid,required"`
}

//企业微信 配置了客户联系功能的员工列表
type ExContactUsers struct{
  Errcode int          `json:"errcode,required"`
  Errmsg  string       `json:"errmsq,required"`
  Follow_user  []string `json:"follow_user,required"`
}

//获取外部客户id列表
type External_list struct{
  Errcode             int          `json:"errcode,required"`
  Errmsg              string       `json:"errmsq,required"`
  External_userid     []string     `json:"external_userid,required"`
}

//获取外部客户详细信息 单个客户
type ExUserInfo struct{
	Errcode int `json:"errcode"`
	Errmsg string `json:"errmsg"`
	ExternalContact struct {
		ExternalUserid string `json:"external_userid"`
		Name string `json:"name"`
		Position string `json:"position"`
		Avatar string `json:"avatar"`
		CorpName string `json:"corp_name"`
		CorpFullName string `json:"corp_full_name"`
		Type int `json:"type"`
		Gender int `json:"gender"`
		Unionid string `json:"unionid"`
		ExternalProfile struct {
			ExternalAttr []struct {
				Type int `json:"type"`
				Name string `json:"name"`
				Text struct {
					Value string `json:"value"`
				} `json:"text,omitempty"`
				Web struct {
					URL string `json:"url"`
					Title string `json:"title"`
				} `json:"web,omitempty"`
				Miniprogram struct {
					Appid string `json:"appid"`
					Pagepath string `json:"pagepath"`
					Title string `json:"title"`
				} `json:"miniprogram,omitempty"`
			} `json:"external_attr"`
		} `json:"external_profile"`
	} `json:"external_contact"`
	FollowUser []struct {
		Userid string `json:"userid"`
		Remark string `json:"remark"`
		Description string `json:"description"`
		Createtime int `json:"createtime"`
		Tags []struct {
			GroupName string `json:"group_name"`
			TagName string `json:"tag_name"`
			TagID string `json:"tag_id"`
			Type int `json:"type"`
		} `json:"tags,omitempty"`
		RemarkCorpName string `json:"remark_corp_name,omitempty"`
		RemarkMobiles []string `json:"remark_mobiles,omitempty"`
		OperUserid string `json:"oper_userid"`
		AddWay int `json:"add_way"`
		State string `json:"state,omitempty"`
	} `json:"follow_user"`
	NextCursor string `json:"next_cursor"`
}

//修改客户联系人描述 备注 公司信息
type UpdateExUserRemarkData struct {
  Userid           string   `json:"userid"`
  ExternalUserid   string   `json:"external_userid"`
  Remark           string   `json:"remark"`
  Description      string   `json:"description"`
  RemarkCompany    string   `json:"remark_company"`
  RemarkMobiles    []string `json:"remark_mobiles"`
  RemarkPicMediaid string   `json:"remark_pic_mediaid"`
}
type UpdateExUserRemarkResult struct{
  Errcode int 
  Errmsg    string 
}


//在职分配 客户
type TransferCustomerData struct {
  HandoverUserid     string   `json:"handover_userid"`
  TakeoverUserid     string   `json:"takeover_userid"`
  ExternalUserid     []string `json:"external_userid"`
  TransferSuccessMsg string   `json:"transfer_success_msg"`
}
type TransferCustomerResult struct {
  Errcode  int    `json:"errcode"`
  Errmsg   string `json:"errmsg"`
  Customer []struct {
    ExternalUserid string `json:"external_userid"`
    Errcode        int    `json:"errcode"`
  } `json:"customer"`
}

//查询客户 在职分配的接替情况
type TransferResultLookData struct {
  HandoverUserid string `json:"handover_userid"`
  TakeoverUserid string `json:"takeover_userid"`
  Cursor         string `json:"cursor"`
}
type TransferResultLookResult struct {
  Errcode  int    `json:"errcode"`
  Errmsg   string `json:"errmsg"`
  Customer []struct {
    ExternalUserid string `json:"external_userid"`
    Status         int    `json:"status"`
    TakeoverTime   int    `json:"takeover_time"`
  } `json:"customer"`
  NextCursor string `json:"next_cursor"`
}

//获取待分配的离职成员列表
type GetUnassignedList struct{
  PageId int   `json:"page_id"`
  Cursor string `json:"cursor"`
  PageSize int   `json:"page_size"`
}

type UnassignedList struct {
  Errcode int    `json:"errcode"`
  Errmsg  string `json:"errmsg"`
  Info    []struct {
    HandoverUserid string `json:"handover_userid"`
    ExternalUserid string `json:"external_userid"`
    DimissionTime  int    `json:"dimission_time"`
  } `json:"info"`
  IsLast     bool   `json:"is_last"`
  NextCursor string `json:"next_cursor"`
}


//分配离职成员的客户
type LiZiTransferCustomerData struct {
  HandoverUserid string   `json:"handover_userid"`
  TakeoverUserid string   `json:"takeover_userid"`
  ExternalUserid []string `json:"external_userid"`
}
type LiZiTransferCustomerResult struct {
  Errcode  int    `json:"errcode"`
  Errmsg   string `json:"errmsg"`
  Customer []struct {
    ExternalUserid string `json:"external_userid"`
    Errcode        int    `json:"errcode"`
  } `json:"customer"`
}


//客户群 管理

//分配离职成员的客户群
type LiZiGroupChatTransferData struct {
  ChatIDList []string `json:"chat_id_list"`
  NewOwner   string   `json:"new_owner"`
}

type LiZiGroupChatTransferResult struct {
  Errcode        int    `json:"errcode"`
  Errmsg         string `json:"errmsg"`
  FailedChatList []struct {
    ChatID  string `json:"chat_id"`
    Errcode int    `json:"errcode"`
    Errmsg  string `json:"errmsg"`
  } `json:"failed_chat_list"`
}

//获取客户群列表 
type GroupChatListData struct {
  StatusFilter int `json:"status_filter"`
  OwnerFilter  struct {
    UseridList []string `json:"userid_list"`
  } `json:"owner_filter"`
  Cursor string `json:"cursor"`
  Limit  int    `json:"limit"`
}
  
type GroupChatListResult struct {
  Errcode       int    `json:"errcode"`
  Errmsg        string `json:"errmsg"`
  GroupChatList []struct {
    ChatID string `json:"chat_id"`
    Status int    `json:"status"`
  } `json:"group_chat_list"`
  NextCursor string `json:"next_cursor"`
}


//获取客户群详情
type GroupChatDetailData struct{
  ChatID string `json:"chat_id"`
}
type GroupChatDetailResult struct {
  Errcode   int    `json:"errcode"`
  Errmsg    string `json:"errmsg"`
  GroupChat struct {
    ChatID     string `json:"chat_id"`
    Name       string `json:"name"`
    Owner      string `json:"owner"`
    CreateTime int    `json:"create_time"`
    Notice     string `json:"notice"`
    MemberList []struct {
      Userid    string `json:"userid"`
      Type      int    `json:"type"`
      JoinTime  int    `json:"join_time"`
      JoinScene int    `json:"join_scene"`
      Invitor   struct {
        Userid string `json:"userid"`
      } `json:"invitor,omitempty"`
      Unionid string `json:"unionid,omitempty"`
    } `json:"member_list"`
    AdminList []struct {
      Userid string `json:"userid"`
    } `json:"admin_list"`
  } `json:"group_chat"`
}


//统计管理
//获取成员联系客户的数据，包括发起申请数、新增客户数、聊天数、发送消息数和删除/拉黑成员的客户数等指标 https://work.weixin.qq.com/api/doc/90000/90135/92132
  
type UserBehaviorData struct {
  Userid    []string `json:"userid"`
  Partyid   []int    `json:"partyid"`
  StartTime int      `json:"start_time"`
  EndTime   int      `json:"end_time"`
}
type UserBehaviorDataResult struct {
  Errcode      int    `json:"errcode"`
  Errmsg       string `json:"errmsg"`
  BehaviorData []struct {
    StatTime            int     `json:"stat_time"`
    ChatCnt             int     `json:"chat_cnt"`
    MessageCnt          int     `json:"message_cnt"`
    ReplyPercentage     float64 `json:"reply_percentage"`
    AvgReplyTime        int     `json:"avg_reply_time"`
    NegativeFeedbackCnt int     `json:"negative_feedback_cnt"`
    NewApplyCnt         int     `json:"new_apply_cnt"`
    NewContactCnt       int     `json:"new_contact_cnt"`
  } `json:"behavior_data"`
}

//群聊数据统计


//按群主聚合的方式
type GroupChatStatisticByPersonData struct {
  DayBeginTime int `json:"day_begin_time"`
  DayEndTime   int `json:"day_end_time"`
  OwnerFilter  struct {
    UseridList []string `json:"userid_list"`
  } `json:"owner_filter"`
  OrderBy  int `json:"order_by"`
  OrderAsc int `json:"order_asc"`
  Offset   int `json:"offset"`
  Limit    int `json:"limit"`
}
type GroupChatStatisticByPersonResult struct {
  Errcode    int    `json:"errcode"`
  Errmsg     string `json:"errmsg"`
  Total      int    `json:"total"`
  NextOffset int    `json:"next_offset"`
  Items      []struct {
    Owner string `json:"owner"`
    Data  struct {
      NewChatCnt   int `json:"new_chat_cnt"`
      ChatTotal    int `json:"chat_total"`
      ChatHasMsg   int `json:"chat_has_msg"`
      NewMemberCnt int `json:"new_member_cnt"`
      MemberTotal  int `json:"member_total"`
      MemberHasMsg int `json:"member_has_msg"`
      MsgTotal     int `json:"msg_total"`
    } `json:"data"`
  } `json:"items"`
}

//按自然日聚合的方式
type GroupChatStatisticByDayData  struct {
  DayBeginTime int `json:"day_begin_time"`
  DayEndTime   int `json:"day_end_time"`
  OwnerFilter  struct {
    UseridList []string `json:"userid_list"`
  } `json:"owner_filter"`
}

type GroupChatStatisticByDayResult struct {
  Errcode int    `json:"errcode"`
  Errmsg  string `json:"errmsg"`
  Items   []struct {
    StatTime int `json:"stat_time"`
    Data     struct {
      NewChatCnt   int `json:"new_chat_cnt"`
      ChatTotal    int `json:"chat_total"`
      ChatHasMsg   int `json:"chat_has_msg"`
      NewMemberCnt int `json:"new_member_cnt"`
      MemberTotal  int `json:"member_total"`
      MemberHasMsg int `json:"member_has_msg"`
      MsgTotal     int `json:"msg_total"`
    } `json:"data"`
  } `json:"items"`
}

//获取企业微信员工据体信息
type WXUserDetail struct {
  Errcode        int    `json:"errcode"`
  Errmsg         string `json:"errmsg"`
  Userid         string `json:"userid"`
  Name           string `json:"name"`
  Department     []int  `json:"department"`
  Order          []int  `json:"order"`
  Position       string `json:"position"`
  Mobile         string `json:"mobile"`
  Gender         string `json:"gender"`
  Email          string `json:"email"`
  IsLeaderInDept []int  `json:"is_leader_in_dept"`
  Avatar         string `json:"avatar"`
  ThumbAvatar    string `json:"thumb_avatar"`
  Telephone      string `json:"telephone"`
  Alias          string `json:"alias"`
  Address        string `json:"address"`
  OpenUserid     string `json:"open_userid"`
  MainDepartment int    `json:"main_department"`
  Extattr        struct {
    Attrs []struct {
      Type int    `json:"type"`
      Name string `json:"name"`
      Text struct {
        Value string `json:"value"`
      } `json:"text,omitempty"`
      Web struct {
        URL   string `json:"url"`
        Title string `json:"title"`
      } `json:"web,omitempty"`
    } `json:"attrs"`
  } `json:"extattr"`
  Status           int    `json:"status"`
  QrCode           string `json:"qr_code"`
  ExternalPosition string `json:"external_position"`
  ExternalProfile  struct {
    ExternalCorpName string `json:"external_corp_name"`
    ExternalAttr     []struct {
      Type int    `json:"type"`
      Name string `json:"name"`
      Text struct {
        Value string `json:"value"`
      } `json:"text,omitempty"`
      Web struct {
        URL   string `json:"url"`
        Title string `json:"title"`
      } `json:"web,omitempty"`
      Miniprogram struct {
        Appid    string `json:"appid"`
        Pagepath string `json:"pagepath"`
        Title    string `json:"title"`
      } `json:"miniprogram,omitempty"`
    } `json:"external_attr"`
  } `json:"external_profile"`
}

//获取jsapi_ticket
type JSAPITicket struct {
  Errcode   int    `json:"errcode"`
  Errmsg    string `json:"errmsg"`
  Ticket    string `json:"ticket"`
  ExpiresIn int    `json:"expires_in"`
}


type DepartmentInfo struct {
  Errcode    int    `json:"errcode"`
  Errmsg     string `json:"errmsg"`
  Department []struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    NameEn   string `json:"name_en"`
    Parentid int    `json:"parentid"`
    Order    int    `json:"order"`
  } `json:"department"`
}


type WXUsers struct {
  Errcode  int    `json:"errcode"`
  Errmsg   string `json:"errmsg"`
  Userlist []struct {
    Userid     string `json:"userid"`
    Name       string `json:"name"`
    Department []int  `json:"department"`
    OpenUserid string `json:"open_userid"`
  } `json:"userlist"`
}


type UploadFileResult struct {
  Errcode   int    `json:"errcode"`
  Errmsg    string `json:"errmsg"`
  Type      string `json:"type"`
  MediaID   string `json:"media_id"`
  CreatedAt string `json:"created_at"`
}

type File struct{
    MediaID string `json:"media_id,omitempty"`
} 

type Text struct{
  Content  string `json:"content,omitempty"`
}


type Textcard struct {
    Title       string `json:"title"`
    Description string `json:"description"`
    URL         string `json:"url"`
    Btntxt      string `json:"btntxt"`
  }

type MsgData struct {
  Touser  string `json:"touser"`
  Msgtype string `json:"msgtype"`
  Agentid int    `json:"agentid"`
  File    File   `json:"file,omitempty"`
  Text    Text   `json:"text,omitempty"`
  Textcard  Textcard  `json:"textcard,omitempty"`
  Safe    int `json:"safe,omitempty"`
  EnableIDTrans int `json:"enable_id_trans,omitempty"`
  EnableDuplicateCheck   int `json:"enable_duplicate_check"`
  DuplicateCheckInterval int `json:"duplicate_check_interval"`
}

type MsgResult struct {
  Errcode      int    `json:"errcode"`
  Errmsg       string `json:"errmsg"`
  Invaliduser  string `json:"invaliduser"`
  Invalidparty string `json:"invalidparty"`
  Invalidtag   string `json:"invalidtag"`
}

