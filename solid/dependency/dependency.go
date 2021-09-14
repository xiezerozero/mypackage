package dependency

import (
	"encoding/json"
	"math/rand"
	"time"
)

type UserInfoRepor interface {
	GetUserInfo(string) (*UserInfoEntity, error)
}

type Cacher interface {
	Get(key string) (string, error)
	Set(key string, value interface{}, expire time.Duration)
	Del(key string) error
}

type UserInfoEntity struct {
	Id           int   `json:"id,omitempty"`
	AdEquityTime int64 `json:"adEquityTime,omitempty"` // 广告权益时间
}

type WxReply struct {
	Id            int    `json:"id"`
	ReplyActivity string `json:"replyActivity"`
}

type ReplyActivity struct {
	ActivityType    int `json:"activityType"`    //1vip活动...
	ActivityDayType int `json:"activityDayType"` //1固定2随机
	ActivityDays    int `json:"activityDays"`    //天数, 随机的话是0-配置的天数
	ActivityCycle   int `json:"activityCycle"`   //领取周期(100000表示永久, 小于100000表示配置天数)
}

// 获取回复的配置活动,返回赠送天数和赠送周期(0,0 表示不应该赠送)
func (w *WxReply) GetActivityConfigDays() (int, int) {
	if w.ReplyActivity == "" {
		return 0, 0
	}
	var ra = &ReplyActivity{}
	e := json.Unmarshal([]byte(w.ReplyActivity), ra)
	if e != nil {
		return 0, 0
	}
	if ra.ActivityDays <= 0 {
		return 0, 0
	}
	if ra.ActivityDayType == 1 { //1固定
		return ra.ActivityDays, ra.ActivityCycle
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(ra.ActivityDays) + 1, ra.ActivityCycle
}
