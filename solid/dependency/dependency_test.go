package dependency

import "testing"

func TestWxReply_GetActivityConfigDays(t *testing.T) {
	tests := []struct {
		name     string
		fields   string
		isRandom bool
		senddays int
		cycle    int
	}{
		{
			name:     "固定每3天赠送10天",
			fields:   "{\"activityType\": 1, \"activityDayType\": 1, \"activityDays\":10, \"activityCycle\": 3}",
			senddays: 10,
			cycle:    3,
		},
		{
			name:     "固定每3天赠送2天",
			fields:   "{\"activityType\": 1, \"activityDayType\": 1, \"activityDays\":2, \"activityCycle\": 3}",
			senddays: 2,
			cycle:    3,
		},
		{
			name:     "随机每3天赠送1-10天",
			fields:   "{\"activityType\": 1, \"activityDayType\": 2, \"activityDays\":10, \"activityCycle\": 3}",
			senddays: 10,
			cycle:    3,
			isRandom: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WxReply{
				ReplyActivity: tt.fields,
			}
			got, got1 := w.GetActivityConfigDays()
			if tt.isRandom {
				if got > tt.senddays || got <= 0 {
					t.Errorf("random GetActivityConfigDays() got = %v, senddays %v", got, tt.senddays)
				}
			} else {
				if got != tt.senddays {
					t.Errorf("GetActivityConfigDays() got = %v, senddays %v", got, tt.senddays)
				}
			}
			if got1 != tt.cycle {
				t.Errorf("GetActivityConfigDays() got1 = %v, senddays %v", got1, tt.cycle)
			}
		})
	}
}
