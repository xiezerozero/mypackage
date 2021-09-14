package domain

import (
	"github.com/xiezerozero/mypackage/solid/dependency"
	"strconv"
	"time"
)

type Reply struct {
	userInfoRepo dependency.UserInfoRepor
	cacheRepo    dependency.Cacher
}

//sendVip  返回是否赠送了vip, 以及vip过期时间(毫秒)
func (r Reply) sendVip(userid string, reply *dependency.WxReply) (isSend bool, days int, vipExpire int64) {
	defer func() {
		// 如果没有赠送, 没有更新vipExpire, 手动查询下会员到期时间
		if vipExpire == 0 {
			userInfo, _ := r.userInfoRepo.GetUserInfo(userid)
			if userInfo != nil {
				vipExpire = userInfo.AdEquityTime
			}
		}
	}()
	if userid == "" || reply == nil {
		return
	}
	days, cycle := reply.GetActivityConfigDays()
	if days <= 0 || cycle <= 0 {
		return
	}
	key := "REPLY_VIP_" + userid + "_" + strconv.Itoa(reply.Id)
	re, e := r.cacheRepo.Get(key)
	if e != nil {
		return
	}
	if re != "" { //存在值,解析出上次赠送时间,判断是否该送了
		lastSend, e := strconv.Atoi(re)
		if e != nil {
			r.cacheRepo.Del(key)
			return
		}
		// 领取周期还未过
		if int64(lastSend+cycle*24*3600) >= time.Now().Unix() {
			return
		}
	}
	// 没有缓存值或者周期已过, 直接送
	r.cacheRepo.Set(key, time.Now().Unix(), (time.Duration(cycle))*24*time.Hour)
	userInfo, _ := r.userInfoRepo.GetUserInfo(userid)
	return true, days, userInfo.AdEquityTime
}
