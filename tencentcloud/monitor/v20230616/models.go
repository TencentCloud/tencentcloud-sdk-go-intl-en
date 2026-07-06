// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20230616

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type AlarmLable struct {
	// label name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// label value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type AlarmNotifyHistory struct {
	// Unique notification ID.
	NotifyId *string `json:"NotifyId,omitnil,omitempty" name:"NotifyId"`

	// Alert policy ID
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// Alarm cycle iD
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// Notification time in Unix timestamp (in seconds).
	NotifyTime *int64 `json:"NotifyTime,omitnil,omitempty" name:"NotifyTime"`

	// Trigger time in Unix timestamp (in seconds).
	TriggerTime *int64 `json:"TriggerTime,omitnil,omitempty" name:"TriggerTime"`

	// Alarm severity level. Valid values: None, Note, Warn, and Serious.
	TriggerLevel *string `json:"TriggerLevel,omitnil,omitempty" name:"TriggerLevel"`

	// alert content
	AlarmContent *string `json:"AlarmContent,omitnil,omitempty" name:"AlarmContent"`

	// Alarm object
	AlarmObject *string `json:"AlarmObject,omitnil,omitempty" name:"AlarmObject"`

	// Alarm notification channel collection involved this time
	// Note: This field may return null, indicating that no valid values can be obtained.
	ChannelSet []*string `json:"ChannelSet,omitnil,omitempty" name:"ChannelSet"`

	// Recipient information of the channel
	ChannelsReceivers []*ChannelsReceivers `json:"ChannelsReceivers,omitnil,omitempty" name:"ChannelsReceivers"`

	// Alarm policy name
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// Prometheus Instance ID, valid only when MT_PROME
	PromeInstanceID *string `json:"PromeInstanceID,omitnil,omitempty" name:"PromeInstanceID"`

	// Region of the Prometheus Instance. Valid at that time only for MT_PROME.
	PromeInstanceRegion *string `json:"PromeInstanceRegion,omitnil,omitempty" name:"PromeInstanceRegion"`

	// Notification template related configuration information
	Notices []*NotifyRelatedNotice `json:"Notices,omitnil,omitempty" name:"Notices"`

	// Alarm trigger status. Valid values: Trigger and Recovery.
	TriggerStatus *string `json:"TriggerStatus,omitnil,omitempty" name:"TriggerStatus"`

	// Console page address related to the present Prometheus notification history, valid only when MR_PROME
	PromeConsoleURL *string `json:"PromeConsoleURL,omitnil,omitempty" name:"PromeConsoleURL"`

	// Alarm label
	Labels []*AlarmLable `json:"Labels,omitnil,omitempty" name:"Labels"`
}

type ChannelsReceivers struct {
	// Notification channel name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// Recipient.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Receivers []*string `json:"Receivers,omitnil,omitempty" name:"Receivers"`

	// Sending result. Valid values: 0, (invalid), 1 (successful), 2 (failed), and 3 (no sending required).
	// Note: This field may return null, indicating that no valid values can be obtained.
	SendStatus *string `json:"SendStatus,omitnil,omitempty" name:"SendStatus"`
}

// Predefined struct for user
type DescribeAlarmNotifyHistoriesRequestParams struct {
	// Monitoring type
	MonitorType *string `json:"MonitorType,omitnil,omitempty" name:"MonitorType"`

	// Start time, used as a Unix timestamp in seconds.
	QueryBaseTime *int64 `json:"QueryBaseTime,omitnil,omitempty" name:"QueryBaseTime"`

	// Period to query before QueryBaseTime, in seconds.
	QueryBeforeSeconds *int64 `json:"QueryBeforeSeconds,omitnil,omitempty" name:"QueryBeforeSeconds"`

	// Pagination parameter.
	PageParams *PageByNoParams `json:"PageParams,omitnil,omitempty" name:"PageParams"`

	// Fill in when the monitoring type is MT_QCE. Namespace of the affiliation.
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// Fill in when the monitoring type is MT_QCE. Alarm policy type
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// Query the notification history of a policy
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

type DescribeAlarmNotifyHistoriesRequest struct {
	*tchttp.BaseRequest
	
	// Monitoring type
	MonitorType *string `json:"MonitorType,omitnil,omitempty" name:"MonitorType"`

	// Start time, used as a Unix timestamp in seconds.
	QueryBaseTime *int64 `json:"QueryBaseTime,omitnil,omitempty" name:"QueryBaseTime"`

	// Period to query before QueryBaseTime, in seconds.
	QueryBeforeSeconds *int64 `json:"QueryBeforeSeconds,omitnil,omitempty" name:"QueryBeforeSeconds"`

	// Pagination parameter.
	PageParams *PageByNoParams `json:"PageParams,omitnil,omitempty" name:"PageParams"`

	// Fill in when the monitoring type is MT_QCE. Namespace of the affiliation.
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// Fill in when the monitoring type is MT_QCE. Alarm policy type
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// Query the notification history of a policy
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

func (r *DescribeAlarmNotifyHistoriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmNotifyHistoriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MonitorType")
	delete(f, "QueryBaseTime")
	delete(f, "QueryBeforeSeconds")
	delete(f, "PageParams")
	delete(f, "Namespace")
	delete(f, "ModelName")
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmNotifyHistoriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmNotifyHistoriesResponseParams struct {
	// Alarm history
	AlarmNotifyHistoryList []*AlarmNotifyHistory `json:"AlarmNotifyHistoryList,omitnil,omitempty" name:"AlarmNotifyHistoryList"`

	// Pagination condition
	PageResult *PageByNoResult `json:"PageResult,omitnil,omitempty" name:"PageResult"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAlarmNotifyHistoriesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlarmNotifyHistoriesResponseParams `json:"Response"`
}

func (r *DescribeAlarmNotifyHistoriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmNotifyHistoriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NotifyRelatedNotice struct {
	// Notification template ID
	NoticeId *string `json:"NoticeId,omitnil,omitempty" name:"NoticeId"`

	// Name of the notification template
	NoticeName *string `json:"NoticeName,omitnil,omitempty" name:"NoticeName"`
}

type PageByNoParams struct {
	// Number of items per page.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PerPage *int64 `json:"PerPage,omitnil,omitempty" name:"PerPage"`

	// Page number, starting from 1.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PageNo *string `json:"PageNo,omitnil,omitempty" name:"PageNo"`
}

type PageByNoResult struct {
	// Total data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Total number of pages.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalPage *int64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// Current page number.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CurrentPageNo *int64 `json:"CurrentPageNo,omitnil,omitempty" name:"CurrentPageNo"`

	// [Deprecated] Whether it has reached the end.
	// Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: IsEnd is deprecated.
	IsEnd *bool `json:"IsEnd,omitnil,omitempty" name:"IsEnd"`

	// Whether it has traversed to the end.
	End *bool `json:"End,omitnil,omitempty" name:"End"`
}