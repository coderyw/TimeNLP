// Package timenlp
// @Author: yinwei
// @File: time_unit_test.go
// @Version: 1.0.0
// @Date: 2023/11/12 20:09

package timenlp

import (
	"testing"
	"time"
)

func TestTimeUnit_normSetCurRelatedWeek(t1 *testing.T) {
	type fields struct {
		expTime                 string
		normalizer              *TimeNormalizer
		tp                      TimePoint
		tpOrigin                TimePoint
		noYear                  bool
		isMorning               bool
		isAllDayTime            bool
		isFirstTimeSolveContext bool
		pos                     int
		length                  int
		ts                      time.Time
	}
	type args struct {
		cur time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		//want   time.Time
		//want1  bool
	}{
		{name: "下周五", fields: fields{expTime: "下周5", normalizer: NewTimeNormalizer(true)}, args: args{time.Now().AddDate(0, 0, 0)}},
		{name: "下周日", fields: fields{expTime: "下周7"}, args: args{time.Now().AddDate(0, 0, 0)}},
		{name: "上周一", fields: fields{expTime: "上周1"}, args: args{time.Now().AddDate(0, 0, 0)}},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TimeUnit{
				expTime:                 tt.fields.expTime,
				normalizer:              tt.fields.normalizer,
				tp:                      tt.fields.tp,
				tpOrigin:                tt.fields.tpOrigin,
				noYear:                  tt.fields.noYear,
				isMorning:               tt.fields.isMorning,
				isAllDayTime:            tt.fields.isAllDayTime,
				isFirstTimeSolveContext: tt.fields.isFirstTimeSolveContext,
				pos:                     tt.fields.pos,
				length:                  tt.fields.length,
				ts:                      tt.fields.ts,
			}
			got, _ := t.normSetCurRelatedWeek(tt.args.cur)
			//if !reflect.DeepEqual(got, tt.want) {
			//	t1.Errorf("normSetCurRelatedWeek() got = %v, want %v", got, tt.want)
			//}
			//if got1 != tt.want1 {
			//	t1.Errorf("normSetCurRelatedWeek() got1 = %v, want %v", got1, tt.want1)
			//}
			t1.Log(got)
		})
	}
}
