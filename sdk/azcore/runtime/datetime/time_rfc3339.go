// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package datetime

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

const (
	utcTimeJSON = `"15:04:05.999999999"`
	utcTime     = "15:04:05.999999999"
	timeFormat  = "15:04:05.999999999Z07:00"
)

type TimeRFC3339 time.Time

func (t TimeRFC3339) MarshalJSON() ([]byte, error) {
	s, _ := t.MarshalText()
	return []byte(fmt.Sprintf("\"%s\"", s)), nil
}

func (t TimeRFC3339) MarshalText() ([]byte, error) {
	tt := time.Time(t)
	return []byte(tt.Format(timeFormat)), nil
}

func (t *TimeRFC3339) UnmarshalJSON(data []byte) error {
	layout := utcTimeJSON
	if tzOffsetRegex.Match(data) {
		layout = timeFormat
	}
	return t.Parse(layout, string(data))
}

func (t *TimeRFC3339) UnmarshalText(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	layout := utcTime
	if tzOffsetRegex.Match(data) {
		layout = timeFormat
	}
	return t.Parse(layout, string(data))
}

func (t *TimeRFC3339) Parse(layout, value string) error {
	p, err := time.Parse(layout, strings.ToUpper(value))
	*t = TimeRFC3339(p)
	return err
}

func (t TimeRFC3339) String() string {
	tt := time.Time(t)
	return tt.Format(timeFormat)
}

func PopulateTimeRFC3339(m map[string]any, k string, t *time.Time) {
	if t == nil {
		return
	} else if azcore.IsNullValue(t) {
		m[k] = nil
		return
	} else if reflect.ValueOf(t).IsNil() {
		return
	}
	m[k] = (*TimeRFC3339)(t)
}

func UnpopulateTimeRFC3339(data json.RawMessage, fn string, t **time.Time) error {
	if data == nil || string(data) == "null" {
		return nil
	}
	var aux TimeRFC3339
	if err := json.Unmarshal(data, &aux); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	*t = (*time.Time)(&aux)
	return nil
}
