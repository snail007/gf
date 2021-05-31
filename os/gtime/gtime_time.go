// Copyright 2018 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/snail007/gf.

package gtime

import (
	"time"
)

type Time struct {
	time.Time
}

// 创建一个空的时间对象，参数可以是标准库时间对象，可选
func New(t ...time.Time) *Time {
	if len(t) > 0 {
		return NewFromTime(t[0])
	}
	return &Time{
		time.Time{},
	}
}

// 当前时间对象
func Now() *Time {
	return &Time{
		time.Now(),
	}
}

// 标准时间对象转换为自定义的时间对象
func NewFromTime(t time.Time) *Time {
	return &Time{
		t,
	}
}

// 从字符串转换为时间对象，复杂的时间字符串需要给定格式
func NewFromStr(str string) *Time {
	if t, err := StrToTime(str); err == nil {
		return t
	}
	return nil
}

// 从字符串转换为时间对象，指定字符串时间格式，format格式形如：Y-m-d H:i:s
func NewFromStrFormat(str string, format string) *Time {
	if t, err := StrToTimeFormat(str, format); err == nil {
		return t
	}
	return nil
}

// 从字符串转换为时间对象，通过标准库layout格式进行解析，layout格式形如：2006-01-02 15:04:05
func NewFromStrLayout(str string, layout string) *Time {
	if t, err := StrToTimeLayout(str, layout); err == nil {
		return t
	}
	return nil
}

// 时间戳转换为时间对象，时间戳支持到纳秒的数值
func NewFromTimeStamp(timestamp int64) *Time {
	if timestamp == 0 {
		return &Time{}
	}
	for timestamp < 1e18 {
		timestamp *= 10
	}
	return &Time{
		time.Unix(int64(timestamp/1e9), timestamp%1e9),
	}
}

// 秒数(时间戳)
func (t *Time) Second() int64 {
	return t.UnixNano() / 1e9
}

// 纳秒数
func (t *Time) Nanosecond() int64 {
	return t.UnixNano()
}

// 微秒数
func (t *Time) Microsecond() int64 {
	return t.UnixNano() / 1e3
}

// 毫秒数
func (t *Time) Millisecond() int64 {
	return t.UnixNano() / 1e6
}

// 转换为字符串
func (t *Time) String() string {
	return t.Format("Y-m-d H:i:s")
}

// 复制当前时间对象
func (t *Time) Clone() *Time {
	return New(t.Time)
}

// 当前时间加上指定时间段
func (t *Time) Add(d time.Duration) *Time {
	t.Time = t.Time.Add(d)
	return t
}

// 当前时间加上指定时间段(使用字符串格式)
func (t *Time) AddStr(duration string) error {
	if d, err := time.ParseDuration(duration); err != nil {
		return err
	} else {
		t.Time = t.Time.Add(d)
	}
	return nil
}

// 时区转换为指定的时区(通过time.Location)
func (t *Time) ToLocation(location *time.Location) *Time {
	t.Time = t.Time.In(location)
	return t
}

// 时区转换为指定的时区(通过时区名称，如：Asia/Shanghai)
func (t *Time) ToZone(zone string) (*Time, error) {
	if l, err := time.LoadLocation(zone); err == nil {
		t.Time = t.Time.In(l)
		return t, nil
	} else {
		return nil, err
	}
}

// 时区转换为UTC时区
func (t *Time) UTC() *Time {
	t.Time = t.Time.UTC()
	return t
}

// 时间输出为ISO8601格式
func (t *Time) ISO8601() string {
	return t.Layout("2006-01-02T15:04:05-07:00")
}

// 时间输出为RFC822格式
func (t *Time) RFC822() string {
	return t.Layout("Mon, 02 Jan 06 15:04 MST")
}

// 时区转换为当前设定的Local时区
func (t *Time) Local() *Time {
	t.Time = t.Time.Local()
	return t
}

// 时间日期计算
func (t *Time) AddDate(years int, months int, days int) *Time {
	t.Time = t.Time.AddDate(years, months, days)
	return t
}

// Round将舍入t的结果返回到d的最接近的倍数(从零时间开始)。
// 中间值的舍入行为是向上舍入。 如果d <= 0，Round返回t剥离任何单调时钟读数但不改变。
// Round作为零时间以来的绝对持续时间运行; 它不适用于当时的演示形式。
// 因此，Round(Hour)可能会返回非零分钟的时间，具体取决于时间的位置。
func (t *Time) Round(d time.Duration) *Time {
	t.Time = t.Time.Round(d)
	return t
}

// Truncate将舍入t的结果返回到d的倍数(从零时间开始)。 如果d <= 0，则Truncate返回t剥离任何单调时钟读数但不改变。
// 截断时间作为零时间以来的绝对持续时间运行; 它不适用于当时的演示形式。
// 因此，截断（小时）可能会返回非零分钟的时间，具体取决于时间的位置。
func (t *Time) Truncate(d time.Duration) *Time {
	t.Time = t.Time.Truncate(d)
	return t
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (t *Time) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.String() + `"`), nil
}
