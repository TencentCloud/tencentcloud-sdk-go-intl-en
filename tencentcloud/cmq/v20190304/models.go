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

package v20190304

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type DeadLetterPolicy struct {
	// DeadLetterQueueName
	// Note: this field may return null, indicating that no valid values can be obtained.
	DeadLetterQueueName *string `json:"DeadLetterQueueName,omitnil,omitempty" name:"DeadLetterQueueName"`

	// DeadLetterQueue
	// Note: this field may return null, indicating that no valid values can be obtained.
	DeadLetterQueue *string `json:"DeadLetterQueue,omitnil,omitempty" name:"DeadLetterQueue"`

	// Policy
	// Note: this field may return null, indicating that no valid values can be obtained.
	Policy *uint64 `json:"Policy,omitnil,omitempty" name:"Policy"`

	// MaxTimeToLive
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxTimeToLive *uint64 `json:"MaxTimeToLive,omitnil,omitempty" name:"MaxTimeToLive"`

	// MaxReceiveCount
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxReceiveCount *uint64 `json:"MaxReceiveCount,omitnil,omitempty" name:"MaxReceiveCount"`
}

type DeadLetterSource struct {
	// QueueId
	// Note: this field may return null, indicating that no valid values can be obtained.
	QueueId *string `json:"QueueId,omitnil,omitempty" name:"QueueId"`

	// QueueName
	// Note: this field may return null, indicating that no valid values can be obtained.
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`
}

// Predefined struct for user
type DescribeQueueDetailRequestParams struct {
	// Starting position of queue list to be returned on the current page in case of paginated return. If a value is entered, `limit` is required. If this parameter is left empty, 0 will be used by default
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of queues to be returned per page in case of paginated return. If this parameter is not passed in, 20 will be used by default. Maximum value: 50.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter parameter. Currently, filtering by `QueueName` is supported, and only one keyword is allowed
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Tag search
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Exact match by `QueueName`
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`
}

type DescribeQueueDetailRequest struct {
	*tchttp.BaseRequest
	
	// Starting position of queue list to be returned on the current page in case of paginated return. If a value is entered, `limit` is required. If this parameter is left empty, 0 will be used by default
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of queues to be returned per page in case of paginated return. If this parameter is not passed in, 20 will be used by default. Maximum value: 50.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter parameter. Currently, filtering by `QueueName` is supported, and only one keyword is allowed
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Tag search
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Exact match by `QueueName`
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`
}

func (r *DescribeQueueDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQueueDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "TagKey")
	delete(f, "QueueName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeQueueDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQueueDetailResponseParams struct {
	// Total number of queues
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Queue list
	QueueSet []*QueueSet `json:"QueueSet,omitnil,omitempty" name:"QueueSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeQueueDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeQueueDetailResponseParams `json:"Response"`
}

func (r *DescribeQueueDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQueueDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicDetailRequestParams struct {
	// Starting position of queue list to be returned on the current page in case of paginated return. If a value is entered, `limit` is required. If this parameter is left empty, 0 will be used by default.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of queues to be returned per page in case of paginated return. If this parameter is not passed in, 20 will be used by default. Maximum value: 50.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Currently, only filtering by `TopicName` is supported, and only one filter value can be entered
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Tag match
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Exact match by `TopicName`
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

type DescribeTopicDetailRequest struct {
	*tchttp.BaseRequest
	
	// Starting position of queue list to be returned on the current page in case of paginated return. If a value is entered, `limit` is required. If this parameter is left empty, 0 will be used by default.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of queues to be returned per page in case of paginated return. If this parameter is not passed in, 20 will be used by default. Maximum value: 50.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Currently, only filtering by `TopicName` is supported, and only one filter value can be entered
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Tag match
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Exact match by `TopicName`
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

func (r *DescribeTopicDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "TagKey")
	delete(f, "TopicName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopicDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicDetailResponseParams struct {
	// TotalCount
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// TopicSet
	TopicSet []*TopicSet `json:"TopicSet,omitnil,omitempty" name:"TopicSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTopicDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopicDetailResponseParams `json:"Response"`
}

func (r *DescribeTopicDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// Filter parameter name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Value
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type QueueSet struct {
	// QueueId
	QueueId *string `json:"QueueId,omitnil,omitempty" name:"QueueId"`

	// QueueName
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// Qps
	// Note: this field may return null, indicating that no valid values can be obtained.
	Qps *uint64 `json:"Qps,omitnil,omitempty" name:"Qps"`

	// Bps
	// Note: this field may return null, indicating that no valid values can be obtained.
	Bps *uint64 `json:"Bps,omitnil,omitempty" name:"Bps"`

	// MaxDelaySeconds
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxDelaySeconds *uint64 `json:"MaxDelaySeconds,omitnil,omitempty" name:"MaxDelaySeconds"`

	// MaxMsgHeapNum
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxMsgHeapNum *uint64 `json:"MaxMsgHeapNum,omitnil,omitempty" name:"MaxMsgHeapNum"`

	// PollingWaitSeconds
	// Note: this field may return null, indicating that no valid values can be obtained.
	PollingWaitSeconds *uint64 `json:"PollingWaitSeconds,omitnil,omitempty" name:"PollingWaitSeconds"`

	// MsgRetentionSeconds
	// Note: this field may return null, indicating that no valid values can be obtained.
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitnil,omitempty" name:"MsgRetentionSeconds"`

	// VisibilityTimeout
	// Note: this field may return null, indicating that no valid values can be obtained.
	VisibilityTimeout *uint64 `json:"VisibilityTimeout,omitnil,omitempty" name:"VisibilityTimeout"`

	// MaxMsgSize
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitnil,omitempty" name:"MaxMsgSize"`

	// RewindSeconds
	// Note: this field may return null, indicating that no valid values can be obtained.
	RewindSeconds *uint64 `json:"RewindSeconds,omitnil,omitempty" name:"RewindSeconds"`

	// CreateTime
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// LastModifyTime
	// Note: this field may return null, indicating that no valid values can be obtained.
	LastModifyTime *uint64 `json:"LastModifyTime,omitnil,omitempty" name:"LastModifyTime"`

	// ActiveMsgNum
	// Note: this field may return null, indicating that no valid values can be obtained.
	ActiveMsgNum *uint64 `json:"ActiveMsgNum,omitnil,omitempty" name:"ActiveMsgNum"`

	// InactiveMsgNum
	// Note: this field may return null, indicating that no valid values can be obtained.
	InactiveMsgNum *uint64 `json:"InactiveMsgNum,omitnil,omitempty" name:"InactiveMsgNum"`

	// DelayMsgNum
	// Note: this field may return null, indicating that no valid values can be obtained.
	DelayMsgNum *uint64 `json:"DelayMsgNum,omitnil,omitempty" name:"DelayMsgNum"`

	// RewindMsgNum
	// Note: this field may return null, indicating that no valid values can be obtained.
	RewindMsgNum *uint64 `json:"RewindMsgNum,omitnil,omitempty" name:"RewindMsgNum"`

	// MinMsgTime
	// Note: this field may return null, indicating that no valid values can be obtained.
	MinMsgTime *uint64 `json:"MinMsgTime,omitnil,omitempty" name:"MinMsgTime"`

	// Transaction
	// Note: this field may return null, indicating that no valid values can be obtained.
	Transaction *bool `json:"Transaction,omitnil,omitempty" name:"Transaction"`

	// DeadLetterSource
	// Note: this field may return null, indicating that no valid values can be obtained.
	DeadLetterSource []*DeadLetterSource `json:"DeadLetterSource,omitnil,omitempty" name:"DeadLetterSource"`

	// DeadLetterPolicy
	// Note: this field may return null, indicating that no valid values can be obtained.
	DeadLetterPolicy *DeadLetterPolicy `json:"DeadLetterPolicy,omitnil,omitempty" name:"DeadLetterPolicy"`

	// TransactionPolicy
	// Note: this field may return null, indicating that no valid values can be obtained.
	TransactionPolicy *TransactionPolicy `json:"TransactionPolicy,omitnil,omitempty" name:"TransactionPolicy"`

	// Creator `uin`
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreateUin *uint64 `json:"CreateUin,omitnil,omitempty" name:"CreateUin"`

	// Tag
	// Note: this field may return null, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Message trace flag. true: enabled, false: not enabled
	// Note: this field may return null, indicating that no valid values can be obtained.
	Trace *bool `json:"Trace,omitnil,omitempty" name:"Trace"`
}

type Tag struct {
	// Tag key
	// Note: this field may return null, indicating that no valid values can be obtained.
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag value
	// Note: this field may return null, indicating that no valid values can be obtained.
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TopicSet struct {
	// TopicId
	// Note: this field may return null, indicating that no valid values can be obtained.
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// TopicName
	// Note: this field may return null, indicating that no valid values can be obtained.
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// MsgRetentionSeconds
	// Note: this field may return null, indicating that no valid values can be obtained.
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitnil,omitempty" name:"MsgRetentionSeconds"`

	// MaxMsgSize
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitnil,omitempty" name:"MaxMsgSize"`

	// Qps
	// Note: this field may return null, indicating that no valid values can be obtained.
	Qps *uint64 `json:"Qps,omitnil,omitempty" name:"Qps"`

	// FilterType
	// Note: this field may return null, indicating that no valid values can be obtained.
	FilterType *uint64 `json:"FilterType,omitnil,omitempty" name:"FilterType"`

	// CreateTime
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// LastModifyTime
	// Note: this field may return null, indicating that no valid values can be obtained.
	LastModifyTime *uint64 `json:"LastModifyTime,omitnil,omitempty" name:"LastModifyTime"`

	// MsgCount
	// Note: this field may return null, indicating that no valid values can be obtained.
	MsgCount *uint64 `json:"MsgCount,omitnil,omitempty" name:"MsgCount"`

	// CreateUin
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreateUin *uint64 `json:"CreateUin,omitnil,omitempty" name:"CreateUin"`

	// Tags
	// Note: this field may return null, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Whether to enable message trace for a topic. true: yes, false: no
	// Note: this field may return null, indicating that no valid values can be obtained.
	Trace *bool `json:"Trace,omitnil,omitempty" name:"Trace"`
}

type TransactionPolicy struct {
	// FirstQueryInterval
	// Note: this field may return null, indicating that no valid values can be obtained.
	FirstQueryInterval *uint64 `json:"FirstQueryInterval,omitnil,omitempty" name:"FirstQueryInterval"`

	// MaxQueryCount
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxQueryCount *uint64 `json:"MaxQueryCount,omitnil,omitempty" name:"MaxQueryCount"`
}