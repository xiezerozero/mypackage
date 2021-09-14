package domain

import (
	"github.com/xiezerozero/mypackage/solid/apdater/repository"
	"github.com/xiezerozero/mypackage/solid/dependency"
	"strconv"
	"testing"
	"time"
)

func TestReply_sendVip(t *testing.T) {
	tests := []struct {
		name          string
		adEquity      int64
		lastsend      int64
		userinfoId    int
		replyid       int
		replyActivity string
		cycle         int
		expectsend    bool
		expectdsys    int
		expectexpire  int64
	}{
		{
			name:          "赠送10天会员",
			adEquity:      time.Now().Unix() * 1000,
			lastsend:      time.Now().Unix() - 24*60*60*3 - 1,
			userinfoId:    1,
			replyid:       1,
			replyActivity: "{\"activityType\": 1, \"activityDayType\": 1, \"activityDays\":10, \"activityCycle\": 3}",
			cycle:         3,
			expectsend:    true,
			expectdsys:    10,
			expectexpire:  time.Now().Unix() * 1000,
		},
		{
			name:          "赠送10天会员01",
			adEquity:      time.Now().Unix() * 1000,
			lastsend:      time.Now().Unix() - 24*60*60*2, //上次赠送时间还没有3天,不赠送
			userinfoId:    1,
			replyid:       1,
			replyActivity: "{\"activityType\": 1, \"activityDayType\": 1, \"activityDays\":10, \"activityCycle\": 3}",
			cycle:         3,
			expectsend:    false,
			expectdsys:    10,
			expectexpire:  time.Now().Unix() * 1000,
		},
	}
	for _, tt := range tests {
		r := Reply{
			userInfoRepo: repository.MockUserinfo{
				Id:           tt.userinfoId,
				AdEquityTime: tt.adEquity,
			},
			cacheRepo: repository.MockCache{TimeUnix: strconv.FormatInt(tt.lastsend, 10)},
		}
		issend, days, expire := r.sendVip("userid", &dependency.WxReply{
			Id:            tt.replyid,
			ReplyActivity: tt.replyActivity,
		})
		t.Run(tt.name, func(t *testing.T) {
			if tt.expectsend != issend {
				t.Errorf("send 10 days 3 days cycle issend, got = %v, expect = %v", issend, tt.expectsend)
			}
			if tt.expectdsys != days {
				t.Errorf("send 10 days 3 days cycle days, got = %d, expect = %d", days, tt.expectdsys)
			}
			if tt.expectexpire != expire {
				t.Errorf("send 10 days 3 days cycle expire, got = %d, expect = %d", tt.expectexpire, expire)
			}
		})
	}
}
