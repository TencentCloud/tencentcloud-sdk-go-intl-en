// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
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
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type ClearQueueRequest struct {
	*tchttp.BaseRequest

	// Queue name, which is unique under the same account in an individual region. It is a string of up to 64 characters, which must begin with a letter and can contain letters, digits, and dashes (`-`).
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
}

func (r *ClearQueueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClearQueueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "QueueName")
	if len(f) > 0 {
		return errors.New("ClearQueueRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ClearQueueResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ClearQueueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClearQueueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClearSubscriptionFilterTagsRequest struct {
	*tchttp.BaseRequest

	// Topic name, which is unique under the same account in an individual region. It is a string of up to 64 characters, which must begin with a letter and can contain letters, digits, and dashes (`-`).
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Subscription name, which is unique in the same topic under the same account in an individual region. It is a string of up to 64 characters, which must begin with a letter and can contain letters, digits, and dashes (`-`).
	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`
}

func (r *ClearSubscriptionFilterTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClearSubscriptionFilterTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicName")
	delete(f, "SubscriptionName")
	if len(f) > 0 {
		return errors.New("ClearSubscriptionFilterTagsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ClearSubscriptionFilterTagsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ClearSubscriptionFilterTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClearSubscriptionFilterTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateQueueRequest struct {
	*tchttp.BaseRequest

	// Queue name, which is unique under the same account in an individual region. It is a string of up to 64 characters, which must begin with a letter and can contain letters, digits, and dashes (`-`).
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`

	// Maximum number of heaped messages. The value range is 1,000,000–10,000,000 during the beta test and can be 1,000,000–1,000,000,000 after the product is officially released. The default value is 10,000,000 during the beta test and will be 100,000,000 after the product is officially released.
	MaxMsgHeapNum *uint64 `json:"MaxMsgHeapNum,omitempty" name:"MaxMsgHeapNum"`

	// Long polling wait time for message reception. Value range: 0–30 seconds. Default value: 0.
	PollingWaitSeconds *uint64 `json:"PollingWaitSeconds,omitempty" name:"PollingWaitSeconds"`

	// Message visibility timeout period. Value range: 1–43200 seconds (i.e., 12 hours). Default value: 30.
	VisibilityTimeout *uint64 `json:"VisibilityTimeout,omitempty" name:"VisibilityTimeout"`

	// Maximum message length. Value range: 1024–65536 bytes (i.e., 1–64 KB). Default value: 65536.
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitempty" name:"MaxMsgSize"`

	// Message retention period. Value range: 60–1296000 seconds (i.e., 1 minute–15 days). Default value: 345600 (i.e., 4 days).
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitempty" name:"MsgRetentionSeconds"`

	// Whether to enable the message rewinding feature for a queue. Value range: 0–msgRetentionSeconds, where 0 means not to enable this feature, while `msgRetentionSeconds` indicates that the maximum rewindable period is the message retention period of the queue.
	RewindSeconds *uint64 `json:"RewindSeconds,omitempty" name:"RewindSeconds"`

	// 1: transaction queue, 0: general queue
	Transaction *uint64 `json:"Transaction,omitempty" name:"Transaction"`

	// First lookback interval
	FirstQueryInterval *uint64 `json:"FirstQueryInterval,omitempty" name:"FirstQueryInterval"`

	// Maximum number of lookbacks
	MaxQueryCount *uint64 `json:"MaxQueryCount,omitempty" name:"MaxQueryCount"`

	// Dead letter queue name
	DeadLetterQueueName *string `json:"DeadLetterQueueName,omitempty" name:"DeadLetterQueueName"`

	// Dead letter policy. 0: message has been consumed multiple times but not deleted, 1: `Time-To-Live` has elapsed
	Policy *uint64 `json:"Policy,omitempty" name:"Policy"`

	// Maximum receipt times. Value range: 1–1000
	MaxReceiveCount *uint64 `json:"MaxReceiveCount,omitempty" name:"MaxReceiveCount"`

	// Maximum period in seconds before an unconsumed message expires, which is required if `policy` is 1. Value range: 300–43200. This value should be smaller than `msgRetentionSeconds` (maximum message retention period)
	MaxTimeToLive *uint64 `json:"MaxTimeToLive,omitempty" name:"MaxTimeToLive"`

	// Whether to enable message trace. true: yes, false: no. If this field is not set, the feature will not be enabled
	Trace *bool `json:"Trace,omitempty" name:"Trace"`
}

func (r *CreateQueueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateQueueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "QueueName")
	delete(f, "MaxMsgHeapNum")
	delete(f, "PollingWaitSeconds")
	delete(f, "VisibilityTimeout")
	delete(f, "MaxMsgSize")
	delete(f, "MsgRetentionSeconds")
	delete(f, "RewindSeconds")
	delete(f, "Transaction")
	delete(f, "FirstQueryInterval")
	delete(f, "MaxQueryCount")
	delete(f, "DeadLetterQueueName")
	delete(f, "Policy")
	delete(f, "MaxReceiveCount")
	delete(f, "MaxTimeToLive")
	delete(f, "Trace")
	if len(f) > 0 {
		return errors.New("CreateQueueRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateQueueResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// `queueId` of a successfully created queue
		QueueId *string `json:"QueueId,omitempty" name:"QueueId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateQueueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateQueueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSubscribeRequest struct {
	*tchttp.BaseRequest

	// Topic name, which is unique under the same account in an individual region. It is a string of up to 64 characters, which must begin with a letter and can contain letters, digits, and dashes (`-`).
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Subscription name, which is unique in the same topic under the same account in an individual region. It is a string of up to 64 characters, which must begin with a letter and can contain letters, digits, and dashes (`-`).
	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`

	// Subscription protocol. Currently, two protocols are supported: http and queue. To use the `http` protocol, you need to build your own web server to receive messages. With the `queue` protocol, messages are automatically pushed to a CMQ queue and you can pull them concurrently.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// `Endpoint` for notification receipt, which is distinguished by `Protocol`. For `http`, `Endpoint` must begin with `http://` and `host` can be a domain name or IP. For `Queue`, enter `QueueName`. Please note that currently the push service cannot push messages to a VPC; therefore, if a VPC domain name or address is entered for `Endpoint`, pushed messages will not be received. Currently, messages can be pushed only to the public network and basic network.
	Endpoint *string `json:"Endpoint,omitempty" name:"Endpoint"`

	// CMQ push server retry policy in case an error occurs while pushing a message to `Endpoint`. Valid values: 1. BACKOFF_RETRY: backoff retry, which is to retry at a fixed interval, discard the message after a certain number of retries, and continue to push the next message; 2. EXPONENTIAL_DECAY_RETRY: exponential decay retry, which is to retry at an exponentially increasing interval, such as 1s, 2s, 4s, 8s, and so on. As a message can be retained in a topic for one day, failed messages will be discarded at most after one day of retry. Default value: EXPONENTIAL_DECAY_RETRY.
	NotifyStrategy *string `json:"NotifyStrategy,omitempty" name:"NotifyStrategy"`

	// Message body tag (used for message filtering). The number of tags cannot exceed 5, and each tag can contain up to 16 characters. It is used in conjunction with the `MsgTag` parameter of `(Batch)PublishMessage`. Rules: 1. If `FilterTag` is not set, no matter whether `MsgTag` is set, the subscription will receive all messages published to the topic; 2. If the `FilterTag` array has a value, only when at least one of the values in the array also exists in the `MsgTag` array (i.e., `FilterTag` and `MsgTag` have an intersection) can the subscription receive messages published to the topic; 3. If the `FilterTag` array has a value, but `MsgTag` is not set, then no message published to the topic will be received, which can be considered as a special case of rule 2 as `FilterTag` and `MsgTag` do not intersect in this case. The overall design idea of rules is based on the intention of the subscriber.
	FilterTag []*string `json:"FilterTag,omitempty" name:"FilterTag" list`

	// The number of `BindingKey` cannot exceed 5, and the length of each `BindingKey` cannot exceed 64 bytes. This field indicates the filtering policy for subscribing to and receiving messages. Each `BindingKey` can contain up to 15 `.`, i.e., up to 16 phrases.
	BindingKey []*string `json:"BindingKey,omitempty" name:"BindingKey" list`

	// Push content format. Valid values: 1. JSON, 2. SIMPLIFIED, i.e., the raw format. If `Protocol` is `queue`, this value must be `SIMPLIFIED`. If `Protocol` is `http`, both options are acceptable, and the default value is `JSON`.
	NotifyContentFormat *string `json:"NotifyContentFormat,omitempty" name:"NotifyContentFormat"`
}

func (r *CreateSubscribeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSubscribeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicName")
	delete(f, "SubscriptionName")
	delete(f, "Protocol")
	delete(f, "Endpoint")
	delete(f, "NotifyStrategy")
	delete(f, "FilterTag")
	delete(f, "BindingKey")
	delete(f, "NotifyContentFormat")
	if len(f) > 0 {
		return errors.New("CreateSubscribeRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateSubscribeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// SubscriptionId
		SubscriptionId *string `json:"SubscriptionId,omitempty" name:"SubscriptionId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSubscribeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTopicRequest struct {
	*tchttp.BaseRequest

	// Topic name, which is unique under the same account in an individual region. It is a string of up to 64 characters, which must begin with a letter and can contain letters, digits, and dashes (`-`).
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Maximum message length. Value range: 1024–65536 bytes (i.e., 1–64 KB). Default value: 65536.
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitempty" name:"MaxMsgSize"`

	// Message match policy for a specified topic.
	FilterType *uint64 `json:"FilterType,omitempty" name:"FilterType"`

	// Message retention period. Value range: 60–86400 seconds (i.e., 1 minute–1 day). Default value: 86400.
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitempty" name:"MsgRetentionSeconds"`

	// Whether to enable message trace. true: yes, false: no. If this field is left empty, the feature will not be enabled.
	Trace *bool `json:"Trace,omitempty" name:"Trace"`
}

func (r *CreateTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicName")
	delete(f, "MaxMsgSize")
	delete(f, "FilterType")
	delete(f, "MsgRetentionSeconds")
	delete(f, "Trace")
	if len(f) > 0 {
		return errors.New("CreateTopicRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// TopicName
		TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeadLetterPolicy struct {

	// DeadLetterQueueName
	// Note: this field may return null, indicating that no valid values can be obtained.
	DeadLetterQueueName *string `json:"DeadLetterQueueName,omitempty" name:"DeadLetterQueueName"`

	// DeadLetterQueue
	// Note: this field may return null, indicating that no valid values can be obtained.
	DeadLetterQueue *string `json:"DeadLetterQueue,omitempty" name:"DeadLetterQueue"`

	// Policy
	// Note: this field may return null, indicating that no valid values can be obtained.
	Policy *uint64 `json:"Policy,omitempty" name:"Policy"`

	// MaxTimeToLive
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxTimeToLive *uint64 `json:"MaxTimeToLive,omitempty" name:"MaxTimeToLive"`

	// MaxReceiveCount
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxReceiveCount *uint64 `json:"MaxReceiveCount,omitempty" name:"MaxReceiveCount"`
}

type DeadLetterSource struct {

	// QueueId
	// Note: this field may return null, indicating that no valid values can be obtained.
	QueueId *string `json:"QueueId,omitempty" name:"QueueId"`

	// QueueName
	// Note: this field may return null, indicating that no valid values can be obtained.
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
}

type DeleteQueueRequest struct {
	*tchttp.BaseRequest

	// Queue name, which is unique under the same account in an individual region. It is a string of up to 64 characters, which must begin with a letter and can contain letters, digits, and dashes (`-`).
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
}

func (r *DeleteQueueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteQueueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "QueueName")
	if len(f) > 0 {
		return errors.New("DeleteQueueRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteQueueResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteQueueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteQueueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSubscribeRequest struct {
	*tchttp.BaseRequest

	// Topic name, which is unique under the same account in an individual region. It is a string of up to 64 characters, which must begin with a letter and can contain letters, digits, and dashes (`-`).
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Subscription name, which is unique in the same topic under the same account in an individual region. It is a string of up to 64 characters, which must begin with a letter and can contain letters, digits, and dashes (`-`).
	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`
}

func (r *DeleteSubscribeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSubscribeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicName")
	delete(f, "SubscriptionName")
	if len(f) > 0 {
		return errors.New("DeleteSubscribeRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSubscribeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSubscribeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTopicRequest struct {
	*tchttp.BaseRequest

	// Topic name, which is unique under the same account in an individual region. It is a string of up to 64 characters, which must begin with a letter and can contain letters, digits, and dashes (`-`).
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

func (r *DeleteTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicName")
	if len(f) > 0 {
		return errors.New("DeleteTopicRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeadLetterSourceQueuesRequest struct {
	*tchttp.BaseRequest

	// Dead letter queue name
	DeadLetterQueueName *string `json:"DeadLetterQueueName,omitempty" name:"DeadLetterQueueName"`

	// Starting position of topic list to be returned on the current page in case of paginated return. If a value is entered, `limit` is required. If this parameter is left empty, 0 will be used by default.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Number of topics to be returned per page in case of paginated return. If this parameter is not passed in, 20 will be used by default. Maximum value: 50.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Filters source queue name of dead letter queue. Currently, only filtering by `SourceQueueName` is supported
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeDeadLetterSourceQueuesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeadLetterSourceQueuesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeadLetterQueueName")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return errors.New("DescribeDeadLetterSourceQueuesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeadLetterSourceQueuesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible queues
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Source queues of dead letter queue
		QueueSet []*DeadLetterSource `json:"QueueSet,omitempty" name:"QueueSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeadLetterSourceQueuesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeadLetterSourceQueuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeQueueDetailRequest struct {
	*tchttp.BaseRequest

	// Starting position of queue list to be returned on the current page in case of paginated return. If a value is entered, `limit` is required. If this parameter is left empty, 0 will be used by default
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of queues to be returned per page in case of paginated return. If this parameter is not passed in, 20 will be used by default. Maximum value: 50.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Filter parameter. Currently, filtering by `QueueName` is supported, and only one keyword is allowed
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Tag search
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Exact match by `QueueName`
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
}

func (r *DescribeQueueDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("DescribeQueueDetailRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeQueueDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of queues
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Queue list
		QueueSet []*QueueSet `json:"QueueSet,omitempty" name:"QueueSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeQueueDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQueueDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscriptionDetailRequest struct {
	*tchttp.BaseRequest

	// Topic name, which is unique under the same account in an individual region. It is a string of up to 64 characters, which must begin with a letter and can contain letters, digits, and dashes (`-`).
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Starting position of topic list to be returned on the current page in case of paginated return. If a value is entered, `limit` is required. If this parameter is left empty, 0 will be used by default
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of topics to be returned per page in case of paginated return. If this parameter is not passed in, 20 will be used by default. Maximum value: 50.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Filter parameter. Currently, only filtering by `SubscriptionName` is supported, and only one keyword is allowed.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeSubscriptionDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubscriptionDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return errors.New("DescribeSubscriptionDetailRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscriptionDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Subscription attribute set
	// Note: this field may return null, indicating that no valid values can be obtained.
		SubscriptionSet []*Subscription `json:"SubscriptionSet,omitempty" name:"SubscriptionSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubscriptionDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubscriptionDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicDetailRequest struct {
	*tchttp.BaseRequest

	// Starting position of queue list to be returned on the current page in case of paginated return. If a value is entered, `limit` is required. If this parameter is left empty, 0 will be used by default.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of queues to be returned per page in case of paginated return. If this parameter is not passed in, 20 will be used by default. Maximum value: 50.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Currently, only filtering by `TopicName` is supported, and only one filter value can be entered
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Tag match
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Exact match by `TopicName`
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

func (r *DescribeTopicDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("DescribeTopicDetailRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// TotalCount
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// TopicSet
		TopicSet []*TopicSet `json:"TopicSet,omitempty" name:"TopicSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopicDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// Filter parameter name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Value
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type ModifyQueueAttributeRequest struct {
	*tchttp.BaseRequest

	// Queue name, which is unique under the same account in an individual region. It is a string of up to 64 characters, which must begin with a letter and can contain letters, digits, and dashes (`-`).
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`

	// Maximum number of heaped messages. The value range is 1,000,000–10,000,000 during the beta test and can be 1,000,000–1,000,000,000 after the product is officially released. The default value is 10,000,000 during the beta test and will be 100,000,000 after the product is officially released.
	MaxMsgHeapNum *uint64 `json:"MaxMsgHeapNum,omitempty" name:"MaxMsgHeapNum"`

	// Long polling wait time for message reception. Value range: 0–30 seconds. Default value: 0.
	PollingWaitSeconds *uint64 `json:"PollingWaitSeconds,omitempty" name:"PollingWaitSeconds"`

	// Message visibility timeout period. Value range: 1–43200 seconds (i.e., 12 hours). Default value: 30.
	VisibilityTimeout *uint64 `json:"VisibilityTimeout,omitempty" name:"VisibilityTimeout"`

	// Maximum message length. Value range: 1024–65536 bytes (i.e., 1–64 KB). Default value: 65536.
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitempty" name:"MaxMsgSize"`

	// Message retention period. Value range: 60–1296000 seconds (i.e., 1 minute–15 days). Default value: 345600 (i.e., 4 days).
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitempty" name:"MsgRetentionSeconds"`

	// Maximum message rewindable period. Value range: 0–msgRetentionSeconds (maximum message retention period of a queue). 0 means not to enable message rewinding.
	RewindSeconds *uint64 `json:"RewindSeconds,omitempty" name:"RewindSeconds"`

	// First query time
	FirstQueryInterval *uint64 `json:"FirstQueryInterval,omitempty" name:"FirstQueryInterval"`

	// Maximum number of queries
	MaxQueryCount *uint64 `json:"MaxQueryCount,omitempty" name:"MaxQueryCount"`

	// Dead letter queue name
	DeadLetterQueueName *string `json:"DeadLetterQueueName,omitempty" name:"DeadLetterQueueName"`

	// Maximum period in seconds before an unconsumed message expires, which is required if `MaxTimeToLivepolicy` is 1. Value range: 300–43200. This value should be smaller than `MsgRetentionSeconds` (maximum message retention period)
	MaxTimeToLive *uint64 `json:"MaxTimeToLive,omitempty" name:"MaxTimeToLive"`

	// Maximum number of receipts
	MaxReceiveCount *uint64 `json:"MaxReceiveCount,omitempty" name:"MaxReceiveCount"`

	// Dead letter queue policy
	Policy *uint64 `json:"Policy,omitempty" name:"Policy"`

	// Whether to enable message trace. true: yes, false: no. If this field is left empty, the feature will not be enabled.
	Trace *bool `json:"Trace,omitempty" name:"Trace"`
}

func (r *ModifyQueueAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyQueueAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "QueueName")
	delete(f, "MaxMsgHeapNum")
	delete(f, "PollingWaitSeconds")
	delete(f, "VisibilityTimeout")
	delete(f, "MaxMsgSize")
	delete(f, "MsgRetentionSeconds")
	delete(f, "RewindSeconds")
	delete(f, "FirstQueryInterval")
	delete(f, "MaxQueryCount")
	delete(f, "DeadLetterQueueName")
	delete(f, "MaxTimeToLive")
	delete(f, "MaxReceiveCount")
	delete(f, "Policy")
	delete(f, "Trace")
	if len(f) > 0 {
		return errors.New("ModifyQueueAttributeRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyQueueAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyQueueAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyQueueAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySubscriptionAttributeRequest struct {
	*tchttp.BaseRequest

	// Topic name, which is unique under the same account in an individual region. It is a string of up to 64 characters, which must begin with a letter and can contain letters, digits, and dashes (`-`).
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Subscription name, which is unique in the same topic under the same account in an individual region. It is a string of up to 64 characters, which must begin with a letter and can contain letters, digits, and dashes (`-`).
	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`

	// CMQ push server retry policy in case an error occurs while pushing a message to `Endpoint`. Valid values:
	// 1. BACKOFF_RETRY: backoff retry, which is to retry at a fixed interval, discard the message after a certain number of retries, and continue to push the next message.
	// 2. EXPONENTIAL_DECAY_RETRY: exponential decay retry, which is to retry at an exponentially increasing interval, such as 1s, 2s, 4s, 8s, and so on. As a message can be retained in a topic for one day, failed messages will be discarded at most after one day of retry. Default value: EXPONENTIAL_DECAY_RETRY.
	NotifyStrategy *string `json:"NotifyStrategy,omitempty" name:"NotifyStrategy"`

	// Push content format. Valid values: 1. JSON, 2. SIMPLIFIED, i.e., the raw format. If `Protocol` is `queue`, this value must be `SIMPLIFIED`. If `Protocol` is `HTTP`, both options are acceptable, and the default value is `JSON`.
	NotifyContentFormat *string `json:"NotifyContentFormat,omitempty" name:"NotifyContentFormat"`

	// Message body tag (used for message filtering). The number of tags cannot exceed 5, and each tag can contain up to 16 characters. It is used in conjunction with the `MsgTag` parameter of `(Batch)PublishMessage`. Rules: 1. If `FilterTag` is not set, no matter whether `MsgTag` is set, the subscription will receive all messages published to the topic; 2. If the `FilterTag` array has a value, only when at least one of the values in the array also exists in the `MsgTag` array (i.e., `FilterTag` and `MsgTag` have an intersection) can the subscription receive messages published to the topic; 3. If the `FilterTag` array has a value, but `MsgTag` is not set, then no message published to the topic will be received, which can be considered as a special case of rule 2 as `FilterTag` and `MsgTag` do not intersect in this case. The overall design idea of rules is based on the intention of the subscriber.
	FilterTags []*string `json:"FilterTags,omitempty" name:"FilterTags" list`

	// The number of `BindingKey` cannot exceed 5, and the length of each `BindingKey` cannot exceed 64 bytes. This field indicates the filtering policy for subscribing to and receiving messages. Each `BindingKey` can contain up to 15 `.`, i.e., up to 16 phrases.
	BindingKey []*string `json:"BindingKey,omitempty" name:"BindingKey" list`
}

func (r *ModifySubscriptionAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubscriptionAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicName")
	delete(f, "SubscriptionName")
	delete(f, "NotifyStrategy")
	delete(f, "NotifyContentFormat")
	delete(f, "FilterTags")
	delete(f, "BindingKey")
	if len(f) > 0 {
		return errors.New("ModifySubscriptionAttributeRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifySubscriptionAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySubscriptionAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubscriptionAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTopicAttributeRequest struct {
	*tchttp.BaseRequest

	// Topic name, which is unique under the same account in an individual region. It is a string of up to 64 characters, which must begin with a letter and can contain letters, digits, and dashes (`-`).
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Maximum message length. Value range: 1024–65536 bytes (i.e., 1–64 KB). Default value: 65536.
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitempty" name:"MaxMsgSize"`

	// Message retention period. Value range: 60–86400 seconds (i.e., 1 minute–1 day). Default value: 86400.
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitempty" name:"MsgRetentionSeconds"`

	// Whether to enable message trace. true: yes, false: no. If this field is left empty, the feature will not be enabled.
	Trace *bool `json:"Trace,omitempty" name:"Trace"`
}

func (r *ModifyTopicAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTopicAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicName")
	delete(f, "MaxMsgSize")
	delete(f, "MsgRetentionSeconds")
	delete(f, "Trace")
	if len(f) > 0 {
		return errors.New("ModifyTopicAttributeRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTopicAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTopicAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTopicAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueueSet struct {

	// QueueId
	QueueId *string `json:"QueueId,omitempty" name:"QueueId"`

	// QueueName
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`

	// Qps
	// Note: this field may return null, indicating that no valid values can be obtained.
	Qps *uint64 `json:"Qps,omitempty" name:"Qps"`

	// Bps
	// Note: this field may return null, indicating that no valid values can be obtained.
	Bps *uint64 `json:"Bps,omitempty" name:"Bps"`

	// MaxDelaySeconds
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxDelaySeconds *uint64 `json:"MaxDelaySeconds,omitempty" name:"MaxDelaySeconds"`

	// MaxMsgHeapNum
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxMsgHeapNum *uint64 `json:"MaxMsgHeapNum,omitempty" name:"MaxMsgHeapNum"`

	// PollingWaitSeconds
	// Note: this field may return null, indicating that no valid values can be obtained.
	PollingWaitSeconds *uint64 `json:"PollingWaitSeconds,omitempty" name:"PollingWaitSeconds"`

	// MsgRetentionSeconds
	// Note: this field may return null, indicating that no valid values can be obtained.
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitempty" name:"MsgRetentionSeconds"`

	// VisibilityTimeout
	// Note: this field may return null, indicating that no valid values can be obtained.
	VisibilityTimeout *uint64 `json:"VisibilityTimeout,omitempty" name:"VisibilityTimeout"`

	// MaxMsgSize
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitempty" name:"MaxMsgSize"`

	// RewindSeconds
	// Note: this field may return null, indicating that no valid values can be obtained.
	RewindSeconds *uint64 `json:"RewindSeconds,omitempty" name:"RewindSeconds"`

	// CreateTime
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// LastModifyTime
	// Note: this field may return null, indicating that no valid values can be obtained.
	LastModifyTime *uint64 `json:"LastModifyTime,omitempty" name:"LastModifyTime"`

	// ActiveMsgNum
	// Note: this field may return null, indicating that no valid values can be obtained.
	ActiveMsgNum *uint64 `json:"ActiveMsgNum,omitempty" name:"ActiveMsgNum"`

	// InactiveMsgNum
	// Note: this field may return null, indicating that no valid values can be obtained.
	InactiveMsgNum *uint64 `json:"InactiveMsgNum,omitempty" name:"InactiveMsgNum"`

	// DelayMsgNum
	// Note: this field may return null, indicating that no valid values can be obtained.
	DelayMsgNum *uint64 `json:"DelayMsgNum,omitempty" name:"DelayMsgNum"`

	// RewindMsgNum
	// Note: this field may return null, indicating that no valid values can be obtained.
	RewindMsgNum *uint64 `json:"RewindMsgNum,omitempty" name:"RewindMsgNum"`

	// MinMsgTime
	// Note: this field may return null, indicating that no valid values can be obtained.
	MinMsgTime *uint64 `json:"MinMsgTime,omitempty" name:"MinMsgTime"`

	// Transaction
	// Note: this field may return null, indicating that no valid values can be obtained.
	Transaction *bool `json:"Transaction,omitempty" name:"Transaction"`

	// DeadLetterSource
	// Note: this field may return null, indicating that no valid values can be obtained.
	DeadLetterSource []*DeadLetterSource `json:"DeadLetterSource,omitempty" name:"DeadLetterSource" list`

	// DeadLetterPolicy
	// Note: this field may return null, indicating that no valid values can be obtained.
	DeadLetterPolicy *DeadLetterPolicy `json:"DeadLetterPolicy,omitempty" name:"DeadLetterPolicy"`

	// TransactionPolicy
	// Note: this field may return null, indicating that no valid values can be obtained.
	TransactionPolicy *TransactionPolicy `json:"TransactionPolicy,omitempty" name:"TransactionPolicy"`

	// Creator `uin`
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// Tag
	// Note: this field may return null, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`

	// Message trace flag. true: enabled, false: not enabled
	// Note: this field may return null, indicating that no valid values can be obtained.
	Trace *bool `json:"Trace,omitempty" name:"Trace"`
}

type RewindQueueRequest struct {
	*tchttp.BaseRequest

	// Queue name, which is unique under the same account in an individual region. It is a string of up to 64 characters, which must begin with a letter and can contain letters, digits, and dashes (`-`).
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`

	// After this time is set, the `(Batch)receiveMessage` API will consume the messages received after this timestamp in the order in which they are produced.
	StartConsumeTime *uint64 `json:"StartConsumeTime,omitempty" name:"StartConsumeTime"`
}

func (r *RewindQueueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RewindQueueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "QueueName")
	delete(f, "StartConsumeTime")
	if len(f) > 0 {
		return errors.New("RewindQueueRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RewindQueueResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RewindQueueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RewindQueueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Subscription struct {

	// SubscriptionName
	// Note: this field may return null, indicating that no valid values can be obtained.
	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`

	// SubscriptionId
	// Note: this field may return null, indicating that no valid values can be obtained.
	SubscriptionId *string `json:"SubscriptionId,omitempty" name:"SubscriptionId"`

	// TopicOwner
	// Note: this field may return null, indicating that no valid values can be obtained.
	TopicOwner *uint64 `json:"TopicOwner,omitempty" name:"TopicOwner"`

	// MsgCount
	// Note: this field may return null, indicating that no valid values can be obtained.
	MsgCount *uint64 `json:"MsgCount,omitempty" name:"MsgCount"`

	// LastModifyTime
	// Note: this field may return null, indicating that no valid values can be obtained.
	LastModifyTime *uint64 `json:"LastModifyTime,omitempty" name:"LastModifyTime"`

	// CreateTime
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// BindingKey
	// Note: this field may return null, indicating that no valid values can be obtained.
	BindingKey []*string `json:"BindingKey,omitempty" name:"BindingKey" list`

	// Endpoint
	// Note: this field may return null, indicating that no valid values can be obtained.
	Endpoint *string `json:"Endpoint,omitempty" name:"Endpoint"`

	// FilterTags
	// Note: this field may return null, indicating that no valid values can be obtained.
	FilterTags []*string `json:"FilterTags,omitempty" name:"FilterTags" list`

	// Protocol
	// Note: this field may return null, indicating that no valid values can be obtained.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// NotifyStrategy
	// Note: this field may return null, indicating that no valid values can be obtained.
	NotifyStrategy *string `json:"NotifyStrategy,omitempty" name:"NotifyStrategy"`

	// NotifyContentFormat
	// Note: this field may return null, indicating that no valid values can be obtained.
	NotifyContentFormat *string `json:"NotifyContentFormat,omitempty" name:"NotifyContentFormat"`
}

type Tag struct {

	// Tag key
	// Note: this field may return null, indicating that no valid values can be obtained.
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value
	// Note: this field may return null, indicating that no valid values can be obtained.
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type TopicSet struct {

	// TopicId
	// Note: this field may return null, indicating that no valid values can be obtained.
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// TopicName
	// Note: this field may return null, indicating that no valid values can be obtained.
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// MsgRetentionSeconds
	// Note: this field may return null, indicating that no valid values can be obtained.
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitempty" name:"MsgRetentionSeconds"`

	// MaxMsgSize
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitempty" name:"MaxMsgSize"`

	// Qps
	// Note: this field may return null, indicating that no valid values can be obtained.
	Qps *uint64 `json:"Qps,omitempty" name:"Qps"`

	// FilterType
	// Note: this field may return null, indicating that no valid values can be obtained.
	FilterType *uint64 `json:"FilterType,omitempty" name:"FilterType"`

	// CreateTime
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// LastModifyTime
	// Note: this field may return null, indicating that no valid values can be obtained.
	LastModifyTime *uint64 `json:"LastModifyTime,omitempty" name:"LastModifyTime"`

	// MsgCount
	// Note: this field may return null, indicating that no valid values can be obtained.
	MsgCount *uint64 `json:"MsgCount,omitempty" name:"MsgCount"`

	// CreateUin
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// Tags
	// Note: this field may return null, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`

	// Whether to enable message trace for a topic. true: yes, false: no
	// Note: this field may return null, indicating that no valid values can be obtained.
	Trace *bool `json:"Trace,omitempty" name:"Trace"`
}

type TransactionPolicy struct {

	// FirstQueryInterval
	// Note: this field may return null, indicating that no valid values can be obtained.
	FirstQueryInterval *uint64 `json:"FirstQueryInterval,omitempty" name:"FirstQueryInterval"`

	// MaxQueryCount
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxQueryCount *uint64 `json:"MaxQueryCount,omitempty" name:"MaxQueryCount"`
}

type UnbindDeadLetterRequest struct {
	*tchttp.BaseRequest

	// Source queue name of dead letter policy. Calling this API will clear the dead letter queue policy of this queue.
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
}

func (r *UnbindDeadLetterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindDeadLetterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "QueueName")
	if len(f) > 0 {
		return errors.New("UnbindDeadLetterRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UnbindDeadLetterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnbindDeadLetterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindDeadLetterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
