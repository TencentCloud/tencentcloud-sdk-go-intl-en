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

package v20201016

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

// Predefined struct for user
type AddMachineGroupInfoRequestParams struct {
	// Machine group ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Machine group type
	// Supported types: `ip` and `label`
	MachineGroupType *MachineGroupTypeInfo `json:"MachineGroupType,omitnil,omitempty" name:"MachineGroupType"`
}

type AddMachineGroupInfoRequest struct {
	*tchttp.BaseRequest
	
	// Machine group ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Machine group type
	// Supported types: `ip` and `label`
	MachineGroupType *MachineGroupTypeInfo `json:"MachineGroupType,omitnil,omitempty" name:"MachineGroupType"`
}

func (r *AddMachineGroupInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddMachineGroupInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "MachineGroupType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddMachineGroupInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddMachineGroupInfoResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddMachineGroupInfoResponse struct {
	*tchttp.BaseResponse
	Response *AddMachineGroupInfoResponseParams `json:"Response"`
}

func (r *AddMachineGroupInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddMachineGroupInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmAnalysisConfig struct {
	// Key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type AlarmInfo struct {
	// Alarm policy name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Monitoring object list
	AlarmTargets []*AlarmTargetInfo `json:"AlarmTargets,omitnil,omitempty" name:"AlarmTargets"`

	// Monitoring task running time point
	MonitorTime *MonitorTime `json:"MonitorTime,omitnil,omitempty" name:"MonitorTime"`

	// Trigger condition
	Condition *string `json:"Condition,omitnil,omitempty" name:"Condition"`

	// Alarm persistence cycle. An alarm will be triggered only after the corresponding trigger condition is met for the number of times specified by `TriggerCount`. Value range: 1–10.
	TriggerCount *int64 `json:"TriggerCount,omitnil,omitempty" name:"TriggerCount"`

	// Repeated alarm interval in minutes. Value range: 0–1440.
	AlarmPeriod *int64 `json:"AlarmPeriod,omitnil,omitempty" name:"AlarmPeriod"`

	// List of associated alarm notification templates
	AlarmNoticeIds []*string `json:"AlarmNoticeIds,omitnil,omitempty" name:"AlarmNoticeIds"`

	// Enablement status
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// Alarm policy ID
	AlarmId *string `json:"AlarmId,omitnil,omitempty" name:"AlarmId"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Last update time
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Custom notification template
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	MessageTemplate *string `json:"MessageTemplate,omitnil,omitempty" name:"MessageTemplate"`

	// Custom callback template
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	CallBack *CallBackInfo `json:"CallBack,omitnil,omitempty" name:"CallBack"`

	// Multi-Dimensional analysis settings
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Analysis []*AnalysisDimensional `json:"Analysis,omitnil,omitempty" name:"Analysis"`
}

type AlarmNotice struct {
	// Alarm notification template name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Alarm template type. Valid values:
	// <br><li> `Trigger`: alarm triggered
	// <br><li> `Recovery`: alarm cleared
	// <br><li> `All`: alarm triggered and alarm cleared
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Information of the recipient in alarm notification template
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	NoticeReceivers []*NoticeReceiver `json:"NoticeReceivers,omitnil,omitempty" name:"NoticeReceivers"`

	// Callback information of alarm notification template
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	WebCallbacks []*WebCallback `json:"WebCallbacks,omitnil,omitempty" name:"WebCallbacks"`

	// Alarm notification template ID
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	AlarmNoticeId *string `json:"AlarmNoticeId,omitnil,omitempty" name:"AlarmNoticeId"`

	// Creation time
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Last update time
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type AlarmTarget struct {
	// Log topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Query statement
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// Monitoring object number, which is incremental from 1.
	Number *int64 `json:"Number,omitnil,omitempty" name:"Number"`

	// Offset of the query start time from the alarm execution time in minutes. The value cannot be positive. Value range: -1440–0.
	StartTimeOffset *int64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// Offset of the query end time from the alarm execution time in minutes. The value cannot be positive and must be greater than `StartTimeOffset`. Value range: -1440–0.
	EndTimeOffset *int64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// Logset ID
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// Search syntax. Valid values:
	// `0` (default): Lucene; `1`: CQL
	// For more information, see <a href="https://intl.cloud.tencent.com/document/product/614/47044?from_cn_redirect=1#RetrievesConditionalRules" target="_blank">Search Syntax</a>.
	SyntaxRule *uint64 `json:"SyntaxRule,omitnil,omitempty" name:"SyntaxRule"`
}

type AlarmTargetInfo struct {
	// Logset ID
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// Logset name
	LogsetName *string `json:"LogsetName,omitnil,omitempty" name:"LogsetName"`

	// Log topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Log topic name
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Query statement
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// Monitoring object number
	Number *int64 `json:"Number,omitnil,omitempty" name:"Number"`

	// Offset of the query start time from the alarm execution time in minutes. The value cannot be positive. Value range: -1440–0.
	StartTimeOffset *int64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// Offset of the query end time from the alarm execution time in minutes. The value cannot be positive and must be greater than `StartTimeOffset`. Value range: -1440–0.
	EndTimeOffset *int64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`
}

type AlertHistoryNotice struct {
	// Notification group name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Notification group ID
	AlarmNoticeId *string `json:"AlarmNoticeId,omitnil,omitempty" name:"AlarmNoticeId"`
}

type AlertHistoryRecord struct {
	// Alarm record ID
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// Alarm policy ID
	AlarmId *string `json:"AlarmId,omitnil,omitempty" name:"AlarmId"`

	// Alarm policy name
	AlarmName *string `json:"AlarmName,omitnil,omitempty" name:"AlarmName"`

	// ID of the monitored object
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Name of the monitored object
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Region of the monitored object
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Trigger condition
	Trigger *string `json:"Trigger,omitnil,omitempty" name:"Trigger"`

	// Number of cycles for which the alarm lasts. An alarm will be triggered only after the trigger condition is met for the number of cycles specified by `TriggerCount`.
	TriggerCount *int64 `json:"TriggerCount,omitnil,omitempty" name:"TriggerCount"`

	// Alarm notification frequency (minutes)
	AlarmPeriod *int64 `json:"AlarmPeriod,omitnil,omitempty" name:"AlarmPeriod"`

	// Notification group
	Notices []*AlertHistoryNotice `json:"Notices,omitnil,omitempty" name:"Notices"`

	// Alarm duration (minutes)
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// Alarm status. Valid values: `0` (uncleared), `1` (cleared), `2` (expired)
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Alarm generation time, which is a Unix timestamp in ms
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Group information corresponding to triggering by group
	// Note: This field may return null, indicating that no valid values can be obtained.
	GroupTriggerCondition []*GroupTriggerConditionInfo `json:"GroupTriggerCondition,omitnil,omitempty" name:"GroupTriggerCondition"`

	// Alarm severity. Valid values: `0` (Warn), `1` (Info), `2` (Critical)
	// Note: This field may return null, indicating that no valid values can be obtained.
	AlarmLevel *uint64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// Type of the monitored object
	// `0`: The same object is specified for all statements. `1`: An object is separately specified for each statement. 
	// Note: This field may return null, indicating that no valid values can be obtained.
	MonitorObjectType *uint64 `json:"MonitorObjectType,omitnil,omitempty" name:"MonitorObjectType"`
}

type AnalysisDimensional struct {
	// Analysis name
	// Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Type of data being analyzed. Valid values: `query`, `field`, `original`
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Analysis content
	// Note: This field may return null, indicating that no valid values can be obtained.
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// Configuration
	// Note: This field may return null, indicating that no valid values can be obtained.
	ConfigInfo []*AlarmAnalysisConfig `json:"ConfigInfo,omitnil,omitempty" name:"ConfigInfo"`
}

// Predefined struct for user
type ApplyConfigToMachineGroupRequestParams struct {
	// Collection configuration ID
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`

	// Machine group ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

type ApplyConfigToMachineGroupRequest struct {
	*tchttp.BaseRequest
	
	// Collection configuration ID
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`

	// Machine group ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

func (r *ApplyConfigToMachineGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyConfigToMachineGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigId")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyConfigToMachineGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyConfigToMachineGroupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ApplyConfigToMachineGroupResponse struct {
	*tchttp.BaseResponse
	Response *ApplyConfigToMachineGroupResponseParams `json:"Response"`
}

func (r *ApplyConfigToMachineGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyConfigToMachineGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CallBackInfo struct {
	// `Body` during callback
	Body *string `json:"Body,omitnil,omitempty" name:"Body"`

	// `Headers` during callback
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Headers []*string `json:"Headers,omitnil,omitempty" name:"Headers"`
}

// Predefined struct for user
type CheckRechargeKafkaServerRequestParams struct {
	// Kafka type. Valid values: 0 (Tencent Cloud CKafka) and 1 (customer's Kafka).
	KafkaType *uint64 `json:"KafkaType,omitnil,omitempty" name:"KafkaType"`

	// CKafka instance ID, which is required when `KafkaType` is set to `0`
	KafkaInstance *string `json:"KafkaInstance,omitnil,omitempty" name:"KafkaInstance"`

	// Service address
	ServerAddr *string `json:"ServerAddr,omitnil,omitempty" name:"ServerAddr"`

	// Whether the service address uses an encrypted connection
	IsEncryptionAddr *bool `json:"IsEncryptionAddr,omitnil,omitempty" name:"IsEncryptionAddr"`

	// Encryption access protocol, which is required when `IsEncryptionAddr` is set to `true`
	Protocol *KafkaProtocolInfo `json:"Protocol,omitnil,omitempty" name:"Protocol"`
}

type CheckRechargeKafkaServerRequest struct {
	*tchttp.BaseRequest
	
	// Kafka type. Valid values: 0 (Tencent Cloud CKafka) and 1 (customer's Kafka).
	KafkaType *uint64 `json:"KafkaType,omitnil,omitempty" name:"KafkaType"`

	// CKafka instance ID, which is required when `KafkaType` is set to `0`
	KafkaInstance *string `json:"KafkaInstance,omitnil,omitempty" name:"KafkaInstance"`

	// Service address
	ServerAddr *string `json:"ServerAddr,omitnil,omitempty" name:"ServerAddr"`

	// Whether the service address uses an encrypted connection
	IsEncryptionAddr *bool `json:"IsEncryptionAddr,omitnil,omitempty" name:"IsEncryptionAddr"`

	// Encryption access protocol, which is required when `IsEncryptionAddr` is set to `true`
	Protocol *KafkaProtocolInfo `json:"Protocol,omitnil,omitempty" name:"Protocol"`
}

func (r *CheckRechargeKafkaServerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckRechargeKafkaServerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KafkaType")
	delete(f, "KafkaInstance")
	delete(f, "ServerAddr")
	delete(f, "IsEncryptionAddr")
	delete(f, "Protocol")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckRechargeKafkaServerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckRechargeKafkaServerResponseParams struct {
	// Kafka cluster accessibility. 0: Accessible.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckRechargeKafkaServerResponse struct {
	*tchttp.BaseResponse
	Response *CheckRechargeKafkaServerResponseParams `json:"Response"`
}

func (r *CheckRechargeKafkaServerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckRechargeKafkaServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Ckafka struct {
	// CKafka VIP
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// CKafka Vport
	Vport *string `json:"Vport,omitnil,omitempty" name:"Vport"`

	// CKafka instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// CKafka instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// CKafka topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// CKafka topic name
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

// Predefined struct for user
type CloseKafkaConsumerRequestParams struct {
	// CLS topic identifier
	FromTopicId *string `json:"FromTopicId,omitnil,omitempty" name:"FromTopicId"`
}

type CloseKafkaConsumerRequest struct {
	*tchttp.BaseRequest
	
	// CLS topic identifier
	FromTopicId *string `json:"FromTopicId,omitnil,omitempty" name:"FromTopicId"`
}

func (r *CloseKafkaConsumerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseKafkaConsumerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FromTopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseKafkaConsumerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseKafkaConsumerResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloseKafkaConsumerResponse struct {
	*tchttp.BaseResponse
	Response *CloseKafkaConsumerResponseParams `json:"Response"`
}

func (r *CloseKafkaConsumerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseKafkaConsumerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Column struct {
	// Column name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Column attribute
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type CompressInfo struct {
	// Compression format. Valid values: `gzip`; `lzop`; `snappy`; `none` (no compression)
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`
}

type ConfigInfo struct {
	// Collection rule configuration ID
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`

	// Name of the collection rule configuration
	// Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Log formatting method
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	LogFormat *string `json:"LogFormat,omitnil,omitempty" name:"LogFormat"`

	// Log collection path
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// Type of the log to be collected. Valid values: `json_log`: log in JSON format; `delimiter_log`: log in delimited format; `minimalist_log`: minimalist log; `multiline_log`: log in multi-line format; `fullregex_log`: log in full regex format. Default value: `minimalist_log`
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// Extraction rule. If `ExtractRule` is set, `LogType` must be set
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ExtractRule *ExtractRuleInfo `json:"ExtractRule,omitnil,omitempty" name:"ExtractRule"`

	// Collection path blocklist
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ExcludePaths []*ExcludePathInfo `json:"ExcludePaths,omitnil,omitempty" name:"ExcludePaths"`

	// Log topic ID (TopicId) of collection configuration
	Output *string `json:"Output,omitnil,omitempty" name:"Output"`

	// Update time
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Custom parsing string
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	UserDefineRule *string `json:"UserDefineRule,omitnil,omitempty" name:"UserDefineRule"`
}

type ConsumerContent struct {
	// Whether to ship tag information
	// Note: This field may return `null`, indicating that no valid value was found.
	EnableTag *bool `json:"EnableTag,omitnil,omitempty" name:"EnableTag"`

	// List of metadata to ship. Supported metadata types: \_\_SOURCE\_\_, \_\_FILENAME\_\_, \_\_TIMESTAMP\_\_, \_\_HOSTNAME\_\_, and \_\_PKGID\_\_.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MetaFields []*string `json:"MetaFields,omitnil,omitempty" name:"MetaFields"`

	// This parameter is required if `EnableTag` is `true`, and is used to specify whether the tag information is JSON tiled. Valid values: `true` (not tiled); `false` (tiled)
	// Note: This field may return `null`, indicating that no valid value was found.
	TagJsonNotTiled *bool `json:"TagJsonNotTiled,omitnil,omitempty" name:"TagJsonNotTiled"`

	// Shipping timestamp precision in seconds (default) or milliseconds
	// Note: This field may return null, indicating that no valid values can be obtained.
	TimestampAccuracy *int64 `json:"TimestampAccuracy,omitnil,omitempty" name:"TimestampAccuracy"`
}

type ContentInfo struct {
	// Content format. Valid values: `json`, `csv`
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// CSV format content description
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Csv *CsvInfo `json:"Csv,omitnil,omitempty" name:"Csv"`

	// JSON format content description
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Json *JsonInfo `json:"Json,omitnil,omitempty" name:"Json"`

	// `Parquet` format description
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Parquet *ParquetInfo `json:"Parquet,omitnil,omitempty" name:"Parquet"`
}

type CosRechargeInfo struct {
	// COS import configuration ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// ID of the log topic.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// ID of the logset.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// COS import task name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// COS bucket.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// Region where the COS bucket is located.
	// Note: This field may return null, indicating that no valid values can be obtained.
	BucketRegion *string `json:"BucketRegion,omitnil,omitempty" name:"BucketRegion"`

	// The prefix of the folder where COS files are located.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Prefix *string `json:"Prefix,omitnil,omitempty" name:"Prefix"`

	// The type of log collected. `json_log`: JSON logs; `delimiter_log`: separator logs; `minimalist_log`: full text in a single line
	// Default value: `minimalist_log`
	// Note: This field may return null, indicating that no valid values can be obtained.
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// Status. `0`: Created, `1`: Running, `2`: Stopped, `3`: Completed, `4`: Run failed
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Whether the configuration is enabled. `0`: Not enabled, `1`: Enabled
	// Note: This field may return null, indicating that no valid values can be obtained.
	Enable *uint64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// Creation time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Update time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Progress in percentage.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Progress *uint64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// Valid values: "" (default), "gzip", "lzop", "snappy"
	// Note: This field may return null, indicating that no valid values can be obtained.
	Compress *string `json:"Compress,omitnil,omitempty" name:"Compress"`

	// See the description of the `ExtractRuleInfo` structure.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExtractRuleInfo *ExtractRuleInfo `json:"ExtractRuleInfo,omitnil,omitempty" name:"ExtractRuleInfo"`
}

// Predefined struct for user
type CreateAlarmNoticeRequestParams struct {
	// Notification group name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Notification type. Valid values:
	// <li> `Trigger`: alarm triggered
	// <li> `Recovery`: alarm cleared
	// <li> `All`: alarm triggered and alarm cleared
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Notification recipient
	NoticeReceivers []*NoticeReceiver `json:"NoticeReceivers,omitnil,omitempty" name:"NoticeReceivers"`

	// API callback information (including WeCom)
	WebCallbacks []*WebCallback `json:"WebCallbacks,omitnil,omitempty" name:"WebCallbacks"`
}

type CreateAlarmNoticeRequest struct {
	*tchttp.BaseRequest
	
	// Notification group name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Notification type. Valid values:
	// <li> `Trigger`: alarm triggered
	// <li> `Recovery`: alarm cleared
	// <li> `All`: alarm triggered and alarm cleared
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Notification recipient
	NoticeReceivers []*NoticeReceiver `json:"NoticeReceivers,omitnil,omitempty" name:"NoticeReceivers"`

	// API callback information (including WeCom)
	WebCallbacks []*WebCallback `json:"WebCallbacks,omitnil,omitempty" name:"WebCallbacks"`
}

func (r *CreateAlarmNoticeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAlarmNoticeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "NoticeReceivers")
	delete(f, "WebCallbacks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAlarmNoticeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAlarmNoticeResponseParams struct {
	// Alarm template ID
	AlarmNoticeId *string `json:"AlarmNoticeId,omitnil,omitempty" name:"AlarmNoticeId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAlarmNoticeResponse struct {
	*tchttp.BaseResponse
	Response *CreateAlarmNoticeResponseParams `json:"Response"`
}

func (r *CreateAlarmNoticeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAlarmNoticeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAlarmRequestParams struct {
	// Alarm policy name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Monitoring object list
	AlarmTargets []*AlarmTarget `json:"AlarmTargets,omitnil,omitempty" name:"AlarmTargets"`

	// Monitoring task running time point
	MonitorTime *MonitorTime `json:"MonitorTime,omitnil,omitempty" name:"MonitorTime"`

	// Trigger condition
	Condition *string `json:"Condition,omitnil,omitempty" name:"Condition"`

	// Alarm persistence cycle. An alarm will be triggered only after the corresponding trigger condition is met for the number of times specified by `TriggerCount`. Value range: 1–10.
	TriggerCount *int64 `json:"TriggerCount,omitnil,omitempty" name:"TriggerCount"`

	// Repeated alarm interval in minutes. Value range: 0–1440.
	AlarmPeriod *int64 `json:"AlarmPeriod,omitnil,omitempty" name:"AlarmPeriod"`

	// List of associated alarm notification templates
	AlarmNoticeIds []*string `json:"AlarmNoticeIds,omitnil,omitempty" name:"AlarmNoticeIds"`

	// Whether to enable the alarm policy. Default value: true
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// Custom alarm content
	MessageTemplate *string `json:"MessageTemplate,omitnil,omitempty" name:"MessageTemplate"`

	// Custom callback
	CallBack *CallBackInfo `json:"CallBack,omitnil,omitempty" name:"CallBack"`

	// Multi-Dimensional analysis
	Analysis []*AnalysisDimensional `json:"Analysis,omitnil,omitempty" name:"Analysis"`
}

type CreateAlarmRequest struct {
	*tchttp.BaseRequest
	
	// Alarm policy name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Monitoring object list
	AlarmTargets []*AlarmTarget `json:"AlarmTargets,omitnil,omitempty" name:"AlarmTargets"`

	// Monitoring task running time point
	MonitorTime *MonitorTime `json:"MonitorTime,omitnil,omitempty" name:"MonitorTime"`

	// Trigger condition
	Condition *string `json:"Condition,omitnil,omitempty" name:"Condition"`

	// Alarm persistence cycle. An alarm will be triggered only after the corresponding trigger condition is met for the number of times specified by `TriggerCount`. Value range: 1–10.
	TriggerCount *int64 `json:"TriggerCount,omitnil,omitempty" name:"TriggerCount"`

	// Repeated alarm interval in minutes. Value range: 0–1440.
	AlarmPeriod *int64 `json:"AlarmPeriod,omitnil,omitempty" name:"AlarmPeriod"`

	// List of associated alarm notification templates
	AlarmNoticeIds []*string `json:"AlarmNoticeIds,omitnil,omitempty" name:"AlarmNoticeIds"`

	// Whether to enable the alarm policy. Default value: true
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// Custom alarm content
	MessageTemplate *string `json:"MessageTemplate,omitnil,omitempty" name:"MessageTemplate"`

	// Custom callback
	CallBack *CallBackInfo `json:"CallBack,omitnil,omitempty" name:"CallBack"`

	// Multi-Dimensional analysis
	Analysis []*AnalysisDimensional `json:"Analysis,omitnil,omitempty" name:"Analysis"`
}

func (r *CreateAlarmRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAlarmRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "AlarmTargets")
	delete(f, "MonitorTime")
	delete(f, "Condition")
	delete(f, "TriggerCount")
	delete(f, "AlarmPeriod")
	delete(f, "AlarmNoticeIds")
	delete(f, "Status")
	delete(f, "MessageTemplate")
	delete(f, "CallBack")
	delete(f, "Analysis")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAlarmRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAlarmResponseParams struct {
	// Alarm policy ID.
	AlarmId *string `json:"AlarmId,omitnil,omitempty" name:"AlarmId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAlarmResponse struct {
	*tchttp.BaseResponse
	Response *CreateAlarmResponseParams `json:"Response"`
}

func (r *CreateAlarmResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAlarmResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConfigRequestParams struct {
	// Collection configuration name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Log topic ID (TopicId) of collection configuration
	Output *string `json:"Output,omitnil,omitempty" name:"Output"`

	// Log collection path containing the filename
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// Type of the log to be collected. Valid values: `json_log`: log in JSON format; `delimiter_log`: log in delimited format; `minimalist_log`: minimalist log; `multiline_log`: log in multi-line format; `fullregex_log`: log in full regex format. Default value: `minimalist_log`
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// Extraction rule. If `ExtractRule` is set, `LogType` must be set.
	ExtractRule *ExtractRuleInfo `json:"ExtractRule,omitnil,omitempty" name:"ExtractRule"`

	// Collection path blocklist
	ExcludePaths []*ExcludePathInfo `json:"ExcludePaths,omitnil,omitempty" name:"ExcludePaths"`

	// Custom collection rule, which is a serialized JSON string
	UserDefineRule *string `json:"UserDefineRule,omitnil,omitempty" name:"UserDefineRule"`

	// Advanced collection configuration
	AdvancedConfig *string `json:"AdvancedConfig,omitnil,omitempty" name:"AdvancedConfig"`
}

type CreateConfigRequest struct {
	*tchttp.BaseRequest
	
	// Collection configuration name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Log topic ID (TopicId) of collection configuration
	Output *string `json:"Output,omitnil,omitempty" name:"Output"`

	// Log collection path containing the filename
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// Type of the log to be collected. Valid values: `json_log`: log in JSON format; `delimiter_log`: log in delimited format; `minimalist_log`: minimalist log; `multiline_log`: log in multi-line format; `fullregex_log`: log in full regex format. Default value: `minimalist_log`
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// Extraction rule. If `ExtractRule` is set, `LogType` must be set.
	ExtractRule *ExtractRuleInfo `json:"ExtractRule,omitnil,omitempty" name:"ExtractRule"`

	// Collection path blocklist
	ExcludePaths []*ExcludePathInfo `json:"ExcludePaths,omitnil,omitempty" name:"ExcludePaths"`

	// Custom collection rule, which is a serialized JSON string
	UserDefineRule *string `json:"UserDefineRule,omitnil,omitempty" name:"UserDefineRule"`

	// Advanced collection configuration
	AdvancedConfig *string `json:"AdvancedConfig,omitnil,omitempty" name:"AdvancedConfig"`
}

func (r *CreateConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Output")
	delete(f, "Path")
	delete(f, "LogType")
	delete(f, "ExtractRule")
	delete(f, "ExcludePaths")
	delete(f, "UserDefineRule")
	delete(f, "AdvancedConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConfigResponseParams struct {
	// Collection configuration ID
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateConfigResponseParams `json:"Response"`
}

func (r *CreateConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConsumerRequestParams struct {
	// Log topic ID to bind
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Whether to ship log metadata. Default value: `true`
	NeedContent *bool `json:"NeedContent,omitnil,omitempty" name:"NeedContent"`

	// Metadata to ship if `NeedContent` is `true`
	Content *ConsumerContent `json:"Content,omitnil,omitempty" name:"Content"`

	// CKafka information
	Ckafka *Ckafka `json:"Ckafka,omitnil,omitempty" name:"Ckafka"`

	// Compression mode. Valid values: `0` (no compression), `2` (snappy), `3` (LZ4).
	Compression *int64 `json:"Compression,omitnil,omitempty" name:"Compression"`
}

type CreateConsumerRequest struct {
	*tchttp.BaseRequest
	
	// Log topic ID to bind
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Whether to ship log metadata. Default value: `true`
	NeedContent *bool `json:"NeedContent,omitnil,omitempty" name:"NeedContent"`

	// Metadata to ship if `NeedContent` is `true`
	Content *ConsumerContent `json:"Content,omitnil,omitempty" name:"Content"`

	// CKafka information
	Ckafka *Ckafka `json:"Ckafka,omitnil,omitempty" name:"Ckafka"`

	// Compression mode. Valid values: `0` (no compression), `2` (snappy), `3` (LZ4).
	Compression *int64 `json:"Compression,omitnil,omitempty" name:"Compression"`
}

func (r *CreateConsumerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConsumerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "NeedContent")
	delete(f, "Content")
	delete(f, "Ckafka")
	delete(f, "Compression")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateConsumerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConsumerResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateConsumerResponse struct {
	*tchttp.BaseResponse
	Response *CreateConsumerResponseParams `json:"Response"`
}

func (r *CreateConsumerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConsumerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCosRechargeRequestParams struct {
	// ID of the log topic.
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// ID of the logset.
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// Shipping task name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// COS bucket.
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// Region where the COS bucket is located.
	BucketRegion *string `json:"BucketRegion,omitnil,omitempty" name:"BucketRegion"`

	// The prefix of the folder where COS files are located.
	Prefix *string `json:"Prefix,omitnil,omitempty" name:"Prefix"`

	// The type of log collected. `json_log`: JSON logs; `delimiter_log`: separator logs; `minimalist_log`: full text in a single line
	// Default value: `minimalist_log`
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// Valid values: "" (default), "gzip", "lzop", "snappy"
	Compress *string `json:"Compress,omitnil,omitempty" name:"Compress"`

	// Extraction rule. If `ExtractRule` is set, `LogType` must be set.
	ExtractRuleInfo *ExtractRuleInfo `json:"ExtractRuleInfo,omitnil,omitempty" name:"ExtractRuleInfo"`
}

type CreateCosRechargeRequest struct {
	*tchttp.BaseRequest
	
	// ID of the log topic.
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// ID of the logset.
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// Shipping task name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// COS bucket.
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// Region where the COS bucket is located.
	BucketRegion *string `json:"BucketRegion,omitnil,omitempty" name:"BucketRegion"`

	// The prefix of the folder where COS files are located.
	Prefix *string `json:"Prefix,omitnil,omitempty" name:"Prefix"`

	// The type of log collected. `json_log`: JSON logs; `delimiter_log`: separator logs; `minimalist_log`: full text in a single line
	// Default value: `minimalist_log`
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// Valid values: "" (default), "gzip", "lzop", "snappy"
	Compress *string `json:"Compress,omitnil,omitempty" name:"Compress"`

	// Extraction rule. If `ExtractRule` is set, `LogType` must be set.
	ExtractRuleInfo *ExtractRuleInfo `json:"ExtractRuleInfo,omitnil,omitempty" name:"ExtractRuleInfo"`
}

func (r *CreateCosRechargeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCosRechargeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "LogsetId")
	delete(f, "Name")
	delete(f, "Bucket")
	delete(f, "BucketRegion")
	delete(f, "Prefix")
	delete(f, "LogType")
	delete(f, "Compress")
	delete(f, "ExtractRuleInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCosRechargeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCosRechargeResponseParams struct {
	// cos_recharge record ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCosRechargeResponse struct {
	*tchttp.BaseResponse
	Response *CreateCosRechargeResponseParams `json:"Response"`
}

func (r *CreateCosRechargeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCosRechargeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDataTransformRequestParams struct {
	// Task type. Valid values: 1 (specified topic) and 2 (dynamically created).
	FuncType *int64 `json:"FuncType,omitnil,omitempty" name:"FuncType"`

	// Source log topic
	SrcTopicId *string `json:"SrcTopicId,omitnil,omitempty" name:"SrcTopicId"`

	// Data processing task name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Data processing statement
	EtlContent *string `json:"EtlContent,omitnil,omitempty" name:"EtlContent"`

	// Data processing type. Valid values: `1`: Use random data from the source log topic for processing preview. `2`: Use user-defined test data for processing preview. `3`: Create a real processing task.
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// Task status. Valid values: 1 (enabled) and 2 (disabled).
	EnableFlag *int64 `json:"EnableFlag,omitnil,omitempty" name:"EnableFlag"`

	// Target topic ID and alias of the data processing task
	DstResources []*DataTransformResouceInfo `json:"DstResources,omitnil,omitempty" name:"DstResources"`

	// Test data used for previewing the processing result
	PreviewLogStatistics []*PreviewLogStatistic `json:"PreviewLogStatistics,omitnil,omitempty" name:"PreviewLogStatistics"`
}

type CreateDataTransformRequest struct {
	*tchttp.BaseRequest
	
	// Task type. Valid values: 1 (specified topic) and 2 (dynamically created).
	FuncType *int64 `json:"FuncType,omitnil,omitempty" name:"FuncType"`

	// Source log topic
	SrcTopicId *string `json:"SrcTopicId,omitnil,omitempty" name:"SrcTopicId"`

	// Data processing task name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Data processing statement
	EtlContent *string `json:"EtlContent,omitnil,omitempty" name:"EtlContent"`

	// Data processing type. Valid values: `1`: Use random data from the source log topic for processing preview. `2`: Use user-defined test data for processing preview. `3`: Create a real processing task.
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// Task status. Valid values: 1 (enabled) and 2 (disabled).
	EnableFlag *int64 `json:"EnableFlag,omitnil,omitempty" name:"EnableFlag"`

	// Target topic ID and alias of the data processing task
	DstResources []*DataTransformResouceInfo `json:"DstResources,omitnil,omitempty" name:"DstResources"`

	// Test data used for previewing the processing result
	PreviewLogStatistics []*PreviewLogStatistic `json:"PreviewLogStatistics,omitnil,omitempty" name:"PreviewLogStatistics"`
}

func (r *CreateDataTransformRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataTransformRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FuncType")
	delete(f, "SrcTopicId")
	delete(f, "Name")
	delete(f, "EtlContent")
	delete(f, "TaskType")
	delete(f, "EnableFlag")
	delete(f, "DstResources")
	delete(f, "PreviewLogStatistics")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDataTransformRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDataTransformResponseParams struct {
	// Task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDataTransformResponse struct {
	*tchttp.BaseResponse
	Response *CreateDataTransformResponseParams `json:"Response"`
}

func (r *CreateDataTransformResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataTransformResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateExportRequestParams struct {
	// Log topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Number of logs to be exported. Maximum value: 50 million
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// Search statements for log export. <a href="https://intl.cloud.tencent.com/document/product/614/44061?from_cn_redirect=1" target="_blank">[SQL statements]</a> are not supported.
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// Start time of the log to be exported, which is a timestamp in milliseconds
	From *int64 `json:"From,omitnil,omitempty" name:"From"`

	// End time of the log to be exported, which is a timestamp in milliseconds
	To *int64 `json:"To,omitnil,omitempty" name:"To"`

	// Exported log sorting order by time. Valid values: `asc`: ascending; `desc`: descending. Default value: `desc`
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Exported log data format. Valid values: `json`, `csv`. Default value: `json`
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`
}

type CreateExportRequest struct {
	*tchttp.BaseRequest
	
	// Log topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Number of logs to be exported. Maximum value: 50 million
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// Search statements for log export. <a href="https://intl.cloud.tencent.com/document/product/614/44061?from_cn_redirect=1" target="_blank">[SQL statements]</a> are not supported.
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// Start time of the log to be exported, which is a timestamp in milliseconds
	From *int64 `json:"From,omitnil,omitempty" name:"From"`

	// End time of the log to be exported, which is a timestamp in milliseconds
	To *int64 `json:"To,omitnil,omitempty" name:"To"`

	// Exported log sorting order by time. Valid values: `asc`: ascending; `desc`: descending. Default value: `desc`
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Exported log data format. Valid values: `json`, `csv`. Default value: `json`
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`
}

func (r *CreateExportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateExportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "Count")
	delete(f, "Query")
	delete(f, "From")
	delete(f, "To")
	delete(f, "Order")
	delete(f, "Format")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateExportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateExportResponseParams struct {
	// Log export ID.
	ExportId *string `json:"ExportId,omitnil,omitempty" name:"ExportId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateExportResponse struct {
	*tchttp.BaseResponse
	Response *CreateExportResponseParams `json:"Response"`
}

func (r *CreateExportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIndexRequestParams struct {
	// Log topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Index rule
	Rule *RuleInfo `json:"Rule,omitnil,omitempty" name:"Rule"`

	// Whether to take effect. Default value: true
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// Whether full-text indexing includes internal fields (`__FILENAME__`, `__HOSTNAME__`, and `__SOURCE__`). Default value: `false`. Recommended value: `true`.
	// * `false`: Full-text indexing does not include internal fields.
	// * `true`: Full-text indexing includes internal fields.
	IncludeInternalFields *bool `json:"IncludeInternalFields,omitnil,omitempty" name:"IncludeInternalFields"`

	// Whether full-text indexing includes metadata fields (which are prefixed with `__TAG__`). Default value: `0`. Recommended value: `1`.
	// * `0`: Full-text indexing includes only the metadata fields with key-value indexing enabled.
	// * `1`: Full-text indexing includes all metadata fields.
	// * `2`: Full-text indexing does not include metadata fields.
	MetadataFlag *uint64 `json:"MetadataFlag,omitnil,omitempty" name:"MetadataFlag"`
}

type CreateIndexRequest struct {
	*tchttp.BaseRequest
	
	// Log topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Index rule
	Rule *RuleInfo `json:"Rule,omitnil,omitempty" name:"Rule"`

	// Whether to take effect. Default value: true
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// Whether full-text indexing includes internal fields (`__FILENAME__`, `__HOSTNAME__`, and `__SOURCE__`). Default value: `false`. Recommended value: `true`.
	// * `false`: Full-text indexing does not include internal fields.
	// * `true`: Full-text indexing includes internal fields.
	IncludeInternalFields *bool `json:"IncludeInternalFields,omitnil,omitempty" name:"IncludeInternalFields"`

	// Whether full-text indexing includes metadata fields (which are prefixed with `__TAG__`). Default value: `0`. Recommended value: `1`.
	// * `0`: Full-text indexing includes only the metadata fields with key-value indexing enabled.
	// * `1`: Full-text indexing includes all metadata fields.
	// * `2`: Full-text indexing does not include metadata fields.
	MetadataFlag *uint64 `json:"MetadataFlag,omitnil,omitempty" name:"MetadataFlag"`
}

func (r *CreateIndexRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIndexRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "Rule")
	delete(f, "Status")
	delete(f, "IncludeInternalFields")
	delete(f, "MetadataFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIndexRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIndexResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateIndexResponse struct {
	*tchttp.BaseResponse
	Response *CreateIndexResponseParams `json:"Response"`
}

func (r *CreateIndexResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateKafkaRechargeRequestParams struct {
	// Target topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Kafka data import configuration name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Kafka type. Valid values: 0 (Tencent Cloud CKafka) and 1 (customer's Kafka).
	KafkaType *uint64 `json:"KafkaType,omitnil,omitempty" name:"KafkaType"`

	// List of Kafka topics to import data from. Separate multiple topics with commas (,).
	UserKafkaTopics *string `json:"UserKafkaTopics,omitnil,omitempty" name:"UserKafkaTopics"`

	// Position for data import. Valid values: -2 (earliest, default) and -1 (latest).
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// CKafka instance ID, which is required when `KafkaType` is set to `0`
	KafkaInstance *string `json:"KafkaInstance,omitnil,omitempty" name:"KafkaInstance"`

	// Service address, which is required when `KafkaType` is set to `1`
	ServerAddr *string `json:"ServerAddr,omitnil,omitempty" name:"ServerAddr"`

	// Whether the service address uses an encrypted connection, which is required when `KafkaType` is set to `1`
	IsEncryptionAddr *bool `json:"IsEncryptionAddr,omitnil,omitempty" name:"IsEncryptionAddr"`

	// Encryption access protocol, which is required when `IsEncryptionAddr` is set to `true`
	Protocol *KafkaProtocolInfo `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Kafka consumer group name
	ConsumerGroupName *string `json:"ConsumerGroupName,omitnil,omitempty" name:"ConsumerGroupName"`

	// Log import rule
	LogRechargeRule *LogRechargeRuleInfo `json:"LogRechargeRule,omitnil,omitempty" name:"LogRechargeRule"`
}

type CreateKafkaRechargeRequest struct {
	*tchttp.BaseRequest
	
	// Target topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Kafka data import configuration name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Kafka type. Valid values: 0 (Tencent Cloud CKafka) and 1 (customer's Kafka).
	KafkaType *uint64 `json:"KafkaType,omitnil,omitempty" name:"KafkaType"`

	// List of Kafka topics to import data from. Separate multiple topics with commas (,).
	UserKafkaTopics *string `json:"UserKafkaTopics,omitnil,omitempty" name:"UserKafkaTopics"`

	// Position for data import. Valid values: -2 (earliest, default) and -1 (latest).
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// CKafka instance ID, which is required when `KafkaType` is set to `0`
	KafkaInstance *string `json:"KafkaInstance,omitnil,omitempty" name:"KafkaInstance"`

	// Service address, which is required when `KafkaType` is set to `1`
	ServerAddr *string `json:"ServerAddr,omitnil,omitempty" name:"ServerAddr"`

	// Whether the service address uses an encrypted connection, which is required when `KafkaType` is set to `1`
	IsEncryptionAddr *bool `json:"IsEncryptionAddr,omitnil,omitempty" name:"IsEncryptionAddr"`

	// Encryption access protocol, which is required when `IsEncryptionAddr` is set to `true`
	Protocol *KafkaProtocolInfo `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Kafka consumer group name
	ConsumerGroupName *string `json:"ConsumerGroupName,omitnil,omitempty" name:"ConsumerGroupName"`

	// Log import rule
	LogRechargeRule *LogRechargeRuleInfo `json:"LogRechargeRule,omitnil,omitempty" name:"LogRechargeRule"`
}

func (r *CreateKafkaRechargeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateKafkaRechargeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "Name")
	delete(f, "KafkaType")
	delete(f, "UserKafkaTopics")
	delete(f, "Offset")
	delete(f, "KafkaInstance")
	delete(f, "ServerAddr")
	delete(f, "IsEncryptionAddr")
	delete(f, "Protocol")
	delete(f, "ConsumerGroupName")
	delete(f, "LogRechargeRule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateKafkaRechargeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateKafkaRechargeResponseParams struct {
	// Kafka data import configuration ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateKafkaRechargeResponse struct {
	*tchttp.BaseResponse
	Response *CreateKafkaRechargeResponseParams `json:"Response"`
}

func (r *CreateKafkaRechargeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateKafkaRechargeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLogsetRequestParams struct {
	// Logset name, which must be unique
	LogsetName *string `json:"LogsetName,omitnil,omitempty" name:"LogsetName"`

	// Tag description list. Up to 10 tag key-value pairs are supported and must be unique.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateLogsetRequest struct {
	*tchttp.BaseRequest
	
	// Logset name, which must be unique
	LogsetName *string `json:"LogsetName,omitnil,omitempty" name:"LogsetName"`

	// Tag description list. Up to 10 tag key-value pairs are supported and must be unique.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreateLogsetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLogsetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogsetName")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLogsetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLogsetResponseParams struct {
	// Logset ID
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLogsetResponse struct {
	*tchttp.BaseResponse
	Response *CreateLogsetResponseParams `json:"Response"`
}

func (r *CreateLogsetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLogsetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMachineGroupRequestParams struct {
	// Machine group name, which must be unique
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Type of the machine group to be created. Valid values: `ip`: use the IP string list in `Values` to create a machine group; `label`: use the tag string list in `Values` to create a machine group
	MachineGroupType *MachineGroupTypeInfo `json:"MachineGroupType,omitnil,omitempty" name:"MachineGroupType"`

	// Tag description list. This parameter is used to bind a tag to a machine group. Up to 10 tag key-value pairs are supported, and a resource can be bound to only one tag key.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Whether to enable automatic update for the machine group
	AutoUpdate *bool `json:"AutoUpdate,omitnil,omitempty" name:"AutoUpdate"`

	// Update start time. We recommend you update LogListener during off-peak hours.
	UpdateStartTime *string `json:"UpdateStartTime,omitnil,omitempty" name:"UpdateStartTime"`

	// Update end time. We recommend you update LogListener during off-peak hours.
	UpdateEndTime *string `json:"UpdateEndTime,omitnil,omitempty" name:"UpdateEndTime"`

	// Whether to enable the service log to record the logs generated by the LogListener service itself. After it is enabled, the internal logset `cls_service_logging` and the `loglistener_status`, `loglistener_alarm`, and `loglistener_business` log topics will be created, which will not incur fees
	ServiceLogging *bool `json:"ServiceLogging,omitnil,omitempty" name:"ServiceLogging"`

	// Metadata information list of a machine group
	MetaTags []*MetaTagInfo `json:"MetaTags,omitnil,omitempty" name:"MetaTags"`
}

type CreateMachineGroupRequest struct {
	*tchttp.BaseRequest
	
	// Machine group name, which must be unique
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Type of the machine group to be created. Valid values: `ip`: use the IP string list in `Values` to create a machine group; `label`: use the tag string list in `Values` to create a machine group
	MachineGroupType *MachineGroupTypeInfo `json:"MachineGroupType,omitnil,omitempty" name:"MachineGroupType"`

	// Tag description list. This parameter is used to bind a tag to a machine group. Up to 10 tag key-value pairs are supported, and a resource can be bound to only one tag key.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Whether to enable automatic update for the machine group
	AutoUpdate *bool `json:"AutoUpdate,omitnil,omitempty" name:"AutoUpdate"`

	// Update start time. We recommend you update LogListener during off-peak hours.
	UpdateStartTime *string `json:"UpdateStartTime,omitnil,omitempty" name:"UpdateStartTime"`

	// Update end time. We recommend you update LogListener during off-peak hours.
	UpdateEndTime *string `json:"UpdateEndTime,omitnil,omitempty" name:"UpdateEndTime"`

	// Whether to enable the service log to record the logs generated by the LogListener service itself. After it is enabled, the internal logset `cls_service_logging` and the `loglistener_status`, `loglistener_alarm`, and `loglistener_business` log topics will be created, which will not incur fees
	ServiceLogging *bool `json:"ServiceLogging,omitnil,omitempty" name:"ServiceLogging"`

	// Metadata information list of a machine group
	MetaTags []*MetaTagInfo `json:"MetaTags,omitnil,omitempty" name:"MetaTags"`
}

func (r *CreateMachineGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMachineGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupName")
	delete(f, "MachineGroupType")
	delete(f, "Tags")
	delete(f, "AutoUpdate")
	delete(f, "UpdateStartTime")
	delete(f, "UpdateEndTime")
	delete(f, "ServiceLogging")
	delete(f, "MetaTags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMachineGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMachineGroupResponseParams struct {
	// Machine group ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateMachineGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateMachineGroupResponseParams `json:"Response"`
}

func (r *CreateMachineGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMachineGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateShipperRequestParams struct {
	// ID of the log topic to which the shipping rule to be created belongs
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Destination bucket in the shipping rule to be created
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// Prefix of the shipping directory in the shipping rule to be created
	Prefix *string `json:"Prefix,omitnil,omitempty" name:"Prefix"`

	// Shipping rule name
	ShipperName *string `json:"ShipperName,omitnil,omitempty" name:"ShipperName"`

	// Interval between shipping tasks (in sec). Default value: 300. Value range: 300-900
	Interval *uint64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// Maximum size of a file to be shipped, in MB. Default value: 256. Value range: 5-256
	MaxSize *uint64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`

	// Filter rules for shipped logs. Only logs matching the rules can be shipped. All rules are in the AND relationship, and up to five rules can be added. If the array is empty, no filtering will be performed, and all logs will be shipped.
	FilterRules []*FilterRuleInfo `json:"FilterRules,omitnil,omitempty" name:"FilterRules"`

	// Rules for partitioning logs to be shipped. `strftime` can be used to define the presentation of time format.
	Partition *string `json:"Partition,omitnil,omitempty" name:"Partition"`

	// Compression configuration of shipped log
	Compress *CompressInfo `json:"Compress,omitnil,omitempty" name:"Compress"`

	// Format configuration of shipped log content
	Content *ContentInfo `json:"Content,omitnil,omitempty" name:"Content"`

	// Naming a shipping file. Valid values: `0` (by random number); `1` (by shipping time). Default value: `0`.
	FilenameMode *uint64 `json:"FilenameMode,omitnil,omitempty" name:"FilenameMode"`

	// Start time for data shipping, which cannot be earlier than the lifecycle start time of the log topic. If you do not specify this parameter, it will be set to the time when you create the data shipping task.
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time for data shipping, which cannot be set to a future time. If you do not specify this parameter, it indicates continuous data shipping.
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type CreateShipperRequest struct {
	*tchttp.BaseRequest
	
	// ID of the log topic to which the shipping rule to be created belongs
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Destination bucket in the shipping rule to be created
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// Prefix of the shipping directory in the shipping rule to be created
	Prefix *string `json:"Prefix,omitnil,omitempty" name:"Prefix"`

	// Shipping rule name
	ShipperName *string `json:"ShipperName,omitnil,omitempty" name:"ShipperName"`

	// Interval between shipping tasks (in sec). Default value: 300. Value range: 300-900
	Interval *uint64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// Maximum size of a file to be shipped, in MB. Default value: 256. Value range: 5-256
	MaxSize *uint64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`

	// Filter rules for shipped logs. Only logs matching the rules can be shipped. All rules are in the AND relationship, and up to five rules can be added. If the array is empty, no filtering will be performed, and all logs will be shipped.
	FilterRules []*FilterRuleInfo `json:"FilterRules,omitnil,omitempty" name:"FilterRules"`

	// Rules for partitioning logs to be shipped. `strftime` can be used to define the presentation of time format.
	Partition *string `json:"Partition,omitnil,omitempty" name:"Partition"`

	// Compression configuration of shipped log
	Compress *CompressInfo `json:"Compress,omitnil,omitempty" name:"Compress"`

	// Format configuration of shipped log content
	Content *ContentInfo `json:"Content,omitnil,omitempty" name:"Content"`

	// Naming a shipping file. Valid values: `0` (by random number); `1` (by shipping time). Default value: `0`.
	FilenameMode *uint64 `json:"FilenameMode,omitnil,omitempty" name:"FilenameMode"`

	// Start time for data shipping, which cannot be earlier than the lifecycle start time of the log topic. If you do not specify this parameter, it will be set to the time when you create the data shipping task.
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time for data shipping, which cannot be set to a future time. If you do not specify this parameter, it indicates continuous data shipping.
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *CreateShipperRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateShipperRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "Bucket")
	delete(f, "Prefix")
	delete(f, "ShipperName")
	delete(f, "Interval")
	delete(f, "MaxSize")
	delete(f, "FilterRules")
	delete(f, "Partition")
	delete(f, "Compress")
	delete(f, "Content")
	delete(f, "FilenameMode")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateShipperRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateShipperResponseParams struct {
	// Shipping task ID.
	ShipperId *string `json:"ShipperId,omitnil,omitempty" name:"ShipperId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateShipperResponse struct {
	*tchttp.BaseResponse
	Response *CreateShipperResponseParams `json:"Response"`
}

func (r *CreateShipperResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateShipperResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTopicRequestParams struct {
	// Logset ID
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// Log topic name
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Number of log topic partitions. Default value: 1. Maximum value: 10
	PartitionCount *int64 `json:"PartitionCount,omitnil,omitempty" name:"PartitionCount"`

	// Tag description list. This parameter is used to bind a tag to a log topic. Up to 10 tag key-value pairs are supported, and a resource can be bound to only one tag key.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Whether to enable automatic split. Default value: true
	AutoSplit *bool `json:"AutoSplit,omitnil,omitempty" name:"AutoSplit"`

	// Maximum number of partitions to split into for this topic if automatic split is enabled. Default value: 50
	MaxSplitPartitions *int64 `json:"MaxSplitPartitions,omitnil,omitempty" name:"MaxSplitPartitions"`

	// Log topic storage type. Valid values: `hot` (STANDARD storage); `cold` (IA storage). Default value: `hot`.
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// Lifecycle in days. Value range: 1–3600 (STANDARD storage); 7–3600 (IA storage). `3640` indicates permanent retention.
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Log topic description
	Describes *string `json:"Describes,omitnil,omitempty" name:"Describes"`

	// `0`: Disable log transitioning.
	// A value other than `0`: The number of STANDARD storage days after log transitioning is enabled (valid only if `StorageType` is `hot`). Note: `HotPeriod` should be greater than or equal to `7` and less than `Period`.
	HotPeriod *uint64 `json:"HotPeriod,omitnil,omitempty" name:"HotPeriod"`

	// Whether to enable web tracking. Valid values: `false` (disable); `true` (enable)
	IsWebTracking *bool `json:"IsWebTracking,omitnil,omitempty" name:"IsWebTracking"`
}

type CreateTopicRequest struct {
	*tchttp.BaseRequest
	
	// Logset ID
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// Log topic name
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Number of log topic partitions. Default value: 1. Maximum value: 10
	PartitionCount *int64 `json:"PartitionCount,omitnil,omitempty" name:"PartitionCount"`

	// Tag description list. This parameter is used to bind a tag to a log topic. Up to 10 tag key-value pairs are supported, and a resource can be bound to only one tag key.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Whether to enable automatic split. Default value: true
	AutoSplit *bool `json:"AutoSplit,omitnil,omitempty" name:"AutoSplit"`

	// Maximum number of partitions to split into for this topic if automatic split is enabled. Default value: 50
	MaxSplitPartitions *int64 `json:"MaxSplitPartitions,omitnil,omitempty" name:"MaxSplitPartitions"`

	// Log topic storage type. Valid values: `hot` (STANDARD storage); `cold` (IA storage). Default value: `hot`.
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// Lifecycle in days. Value range: 1–3600 (STANDARD storage); 7–3600 (IA storage). `3640` indicates permanent retention.
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Log topic description
	Describes *string `json:"Describes,omitnil,omitempty" name:"Describes"`

	// `0`: Disable log transitioning.
	// A value other than `0`: The number of STANDARD storage days after log transitioning is enabled (valid only if `StorageType` is `hot`). Note: `HotPeriod` should be greater than or equal to `7` and less than `Period`.
	HotPeriod *uint64 `json:"HotPeriod,omitnil,omitempty" name:"HotPeriod"`

	// Whether to enable web tracking. Valid values: `false` (disable); `true` (enable)
	IsWebTracking *bool `json:"IsWebTracking,omitnil,omitempty" name:"IsWebTracking"`
}

func (r *CreateTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogsetId")
	delete(f, "TopicName")
	delete(f, "PartitionCount")
	delete(f, "Tags")
	delete(f, "AutoSplit")
	delete(f, "MaxSplitPartitions")
	delete(f, "StorageType")
	delete(f, "Period")
	delete(f, "Describes")
	delete(f, "HotPeriod")
	delete(f, "IsWebTracking")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTopicResponseParams struct {
	// Log topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTopicResponse struct {
	*tchttp.BaseResponse
	Response *CreateTopicResponseParams `json:"Response"`
}

func (r *CreateTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CsvInfo struct {
	// Whether to print `key` on the first row of the CSV file
	PrintKey *bool `json:"PrintKey,omitnil,omitempty" name:"PrintKey"`

	// Names of keys
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Keys []*string `json:"Keys,omitnil,omitempty" name:"Keys"`

	// Field delimiter
	Delimiter *string `json:"Delimiter,omitnil,omitempty" name:"Delimiter"`

	// Escape character used to enclose any field delimiter in field content. You can enter only a single quotation mark, double quotation mark, or an empty string.
	EscapeChar *string `json:"EscapeChar,omitnil,omitempty" name:"EscapeChar"`

	// Content used to populate non-existing fields
	NonExistingField *string `json:"NonExistingField,omitnil,omitempty" name:"NonExistingField"`
}

type DataTransformResouceInfo struct {
	// Target topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Alias
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`
}

type DataTransformTaskInfo struct {
	// Data processing task name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Data processing task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Task status. Valid values: 1 (enabled) and 2 (disabled).
	EnableFlag *int64 `json:"EnableFlag,omitnil,omitempty" name:"EnableFlag"`

	// Task type. Valid values: 1 (DSL) and 2 (SQL).
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Source log topic
	SrcTopicId *string `json:"SrcTopicId,omitnil,omitempty" name:"SrcTopicId"`

	// Current task status. Valid values: 1 (preparing), 2 (in progress), 3 (being stopped), and 4 (stopped).
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Task creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Last modified time
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Last enabled time. If you need to rebuild a cluster, modify this time.
	LastEnableTime *string `json:"LastEnableTime,omitnil,omitempty" name:"LastEnableTime"`

	// Log topic name
	SrcTopicName *string `json:"SrcTopicName,omitnil,omitempty" name:"SrcTopicName"`

	// Logset ID
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// Target topic ID and alias of the data processing task
	DstResources []*DataTransformResouceInfo `json:"DstResources,omitnil,omitempty" name:"DstResources"`

	// Logical function for data processing
	EtlContent *string `json:"EtlContent,omitnil,omitempty" name:"EtlContent"`
}

// Predefined struct for user
type DeleteAlarmNoticeRequestParams struct {
	// Notification group ID
	AlarmNoticeId *string `json:"AlarmNoticeId,omitnil,omitempty" name:"AlarmNoticeId"`
}

type DeleteAlarmNoticeRequest struct {
	*tchttp.BaseRequest
	
	// Notification group ID
	AlarmNoticeId *string `json:"AlarmNoticeId,omitnil,omitempty" name:"AlarmNoticeId"`
}

func (r *DeleteAlarmNoticeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAlarmNoticeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AlarmNoticeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAlarmNoticeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAlarmNoticeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAlarmNoticeResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAlarmNoticeResponseParams `json:"Response"`
}

func (r *DeleteAlarmNoticeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAlarmNoticeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAlarmRequestParams struct {
	// Alarm policy ID.
	AlarmId *string `json:"AlarmId,omitnil,omitempty" name:"AlarmId"`
}

type DeleteAlarmRequest struct {
	*tchttp.BaseRequest
	
	// Alarm policy ID.
	AlarmId *string `json:"AlarmId,omitnil,omitempty" name:"AlarmId"`
}

func (r *DeleteAlarmRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAlarmRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AlarmId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAlarmRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAlarmResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAlarmResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAlarmResponseParams `json:"Response"`
}

func (r *DeleteAlarmResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAlarmResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteConfigFromMachineGroupRequestParams struct {
	// Machine group ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Collection configuration ID
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`
}

type DeleteConfigFromMachineGroupRequest struct {
	*tchttp.BaseRequest
	
	// Machine group ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Collection configuration ID
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`
}

func (r *DeleteConfigFromMachineGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConfigFromMachineGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "ConfigId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteConfigFromMachineGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteConfigFromMachineGroupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteConfigFromMachineGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteConfigFromMachineGroupResponseParams `json:"Response"`
}

func (r *DeleteConfigFromMachineGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConfigFromMachineGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteConfigRequestParams struct {
	// Collection rule configuration ID
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`
}

type DeleteConfigRequest struct {
	*tchttp.BaseRequest
	
	// Collection rule configuration ID
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`
}

func (r *DeleteConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteConfigResponse struct {
	*tchttp.BaseResponse
	Response *DeleteConfigResponseParams `json:"Response"`
}

func (r *DeleteConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteConsumerRequestParams struct {
	// Log topic ID bound to the task
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

type DeleteConsumerRequest struct {
	*tchttp.BaseRequest
	
	// Log topic ID bound to the task
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

func (r *DeleteConsumerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConsumerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteConsumerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteConsumerResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteConsumerResponse struct {
	*tchttp.BaseResponse
	Response *DeleteConsumerResponseParams `json:"Response"`
}

func (r *DeleteConsumerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConsumerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDataTransformRequestParams struct {
	// Data processing task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DeleteDataTransformRequest struct {
	*tchttp.BaseRequest
	
	// Data processing task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DeleteDataTransformRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDataTransformRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDataTransformRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDataTransformResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDataTransformResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDataTransformResponseParams `json:"Response"`
}

func (r *DeleteDataTransformResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDataTransformResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteExportRequestParams struct {
	// Log export ID
	ExportId *string `json:"ExportId,omitnil,omitempty" name:"ExportId"`
}

type DeleteExportRequest struct {
	*tchttp.BaseRequest
	
	// Log export ID
	ExportId *string `json:"ExportId,omitnil,omitempty" name:"ExportId"`
}

func (r *DeleteExportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteExportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ExportId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteExportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteExportResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteExportResponse struct {
	*tchttp.BaseResponse
	Response *DeleteExportResponseParams `json:"Response"`
}

func (r *DeleteExportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIndexRequestParams struct {
	// Log topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

type DeleteIndexRequest struct {
	*tchttp.BaseRequest
	
	// Log topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

func (r *DeleteIndexRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIndexRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIndexRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIndexResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteIndexResponse struct {
	*tchttp.BaseResponse
	Response *DeleteIndexResponseParams `json:"Response"`
}

func (r *DeleteIndexResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteKafkaRechargeRequestParams struct {
	// Kafka data import configuration ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Target CLS log topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

type DeleteKafkaRechargeRequest struct {
	*tchttp.BaseRequest
	
	// Kafka data import configuration ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Target CLS log topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

func (r *DeleteKafkaRechargeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteKafkaRechargeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "TopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteKafkaRechargeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteKafkaRechargeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteKafkaRechargeResponse struct {
	*tchttp.BaseResponse
	Response *DeleteKafkaRechargeResponseParams `json:"Response"`
}

func (r *DeleteKafkaRechargeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteKafkaRechargeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLogsetRequestParams struct {
	// Logset ID
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`
}

type DeleteLogsetRequest struct {
	*tchttp.BaseRequest
	
	// Logset ID
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`
}

func (r *DeleteLogsetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLogsetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogsetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLogsetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLogsetResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLogsetResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLogsetResponseParams `json:"Response"`
}

func (r *DeleteLogsetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLogsetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMachineGroupInfoRequestParams struct {
	// Machine group ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Machine group type
	// Supported types: `ip` and `label`
	MachineGroupType *MachineGroupTypeInfo `json:"MachineGroupType,omitnil,omitempty" name:"MachineGroupType"`
}

type DeleteMachineGroupInfoRequest struct {
	*tchttp.BaseRequest
	
	// Machine group ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Machine group type
	// Supported types: `ip` and `label`
	MachineGroupType *MachineGroupTypeInfo `json:"MachineGroupType,omitnil,omitempty" name:"MachineGroupType"`
}

func (r *DeleteMachineGroupInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMachineGroupInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "MachineGroupType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMachineGroupInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMachineGroupInfoResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteMachineGroupInfoResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMachineGroupInfoResponseParams `json:"Response"`
}

func (r *DeleteMachineGroupInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMachineGroupInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMachineGroupRequestParams struct {
	// Machine group ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

type DeleteMachineGroupRequest struct {
	*tchttp.BaseRequest
	
	// Machine group ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

func (r *DeleteMachineGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMachineGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMachineGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMachineGroupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteMachineGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMachineGroupResponseParams `json:"Response"`
}

func (r *DeleteMachineGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMachineGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteShipperRequestParams struct {
	// Shipping rule ID
	ShipperId *string `json:"ShipperId,omitnil,omitempty" name:"ShipperId"`
}

type DeleteShipperRequest struct {
	*tchttp.BaseRequest
	
	// Shipping rule ID
	ShipperId *string `json:"ShipperId,omitnil,omitempty" name:"ShipperId"`
}

func (r *DeleteShipperRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteShipperRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ShipperId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteShipperRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteShipperResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteShipperResponse struct {
	*tchttp.BaseResponse
	Response *DeleteShipperResponseParams `json:"Response"`
}

func (r *DeleteShipperResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteShipperResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTopicRequestParams struct {
	// Log topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

type DeleteTopicRequest struct {
	*tchttp.BaseRequest
	
	// Log topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

func (r *DeleteTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTopicResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTopicResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTopicResponseParams `json:"Response"`
}

func (r *DeleteTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmNoticesRequestParams struct {
	// <li> name
	// Filter by **notification group name**.
	// Type: String
	// Required: No
	// <li> alarmNoticeId
	// Filter by **notification group ID**.
	// Type: String
	// Required: No
	// <li> uid
	// Filter by **recipient ID**.
	// Type: String
	// Required: No
	// <li> groupId
	// Filter by **recipient ID**.
	// Type: String
	// Required: No
	// 
	// Each request can have up to 10 `Filters` and 5 `Filter.Values`.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Page offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of entries per page. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeAlarmNoticesRequest struct {
	*tchttp.BaseRequest
	
	// <li> name
	// Filter by **notification group name**.
	// Type: String
	// Required: No
	// <li> alarmNoticeId
	// Filter by **notification group ID**.
	// Type: String
	// Required: No
	// <li> uid
	// Filter by **recipient ID**.
	// Type: String
	// Required: No
	// <li> groupId
	// Filter by **recipient ID**.
	// Type: String
	// Required: No
	// 
	// Each request can have up to 10 `Filters` and 5 `Filter.Values`.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Page offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of entries per page. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeAlarmNoticesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmNoticesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmNoticesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmNoticesResponseParams struct {
	// Alarm notification template list
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	AlarmNotices []*AlarmNotice `json:"AlarmNotices,omitnil,omitempty" name:"AlarmNotices"`

	// Total number of eligible alarm notification templates
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAlarmNoticesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlarmNoticesResponseParams `json:"Response"`
}

func (r *DescribeAlarmNoticesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmNoticesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmsRequestParams struct {
	// name
	// - Filter by **alarm policy name**
	// - Type: String
	// - Required: No
	// 
	// alarmId
	// - Filter by **alarm policy ID**
	// - Type: String
	// - Required: No
	// 
	// topicId
	// - Filter by **log topic ID**
	// - Type: String
	// - Required: No
	// 
	// enable
	// - Filter by **enablement status**
	// - Type: String
	// - Note: The valid values of `enable` include `1`, `t`, `T`, `TRUE`, `true`, `True`, `0`, `f`, `F`, `FALSE`, `false`, and `False`. If other values are entered, an "invalid parameter" error will be returned.
	// - Required: No
	// 
	// Each request can contain up to 10 `Filters` and 5 `Filter.Values`.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Page offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of entries per page. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeAlarmsRequest struct {
	*tchttp.BaseRequest
	
	// name
	// - Filter by **alarm policy name**
	// - Type: String
	// - Required: No
	// 
	// alarmId
	// - Filter by **alarm policy ID**
	// - Type: String
	// - Required: No
	// 
	// topicId
	// - Filter by **log topic ID**
	// - Type: String
	// - Required: No
	// 
	// enable
	// - Filter by **enablement status**
	// - Type: String
	// - Note: The valid values of `enable` include `1`, `t`, `T`, `TRUE`, `true`, `True`, `0`, `f`, `F`, `FALSE`, `false`, and `False`. If other values are entered, an "invalid parameter" error will be returned.
	// - Required: No
	// 
	// Each request can contain up to 10 `Filters` and 5 `Filter.Values`.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Page offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of entries per page. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeAlarmsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmsResponseParams struct {
	// Alarm policy list
	Alarms []*AlarmInfo `json:"Alarms,omitnil,omitempty" name:"Alarms"`

	// Number of eligible alarm policies
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAlarmsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlarmsResponseParams `json:"Response"`
}

func (r *DescribeAlarmsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlertRecordHistoryRequestParams struct {
	// Start time of the query range, which is a Unix timestamp in ms
	From *uint64 `json:"From,omitnil,omitempty" name:"From"`

	// End time of the query range, which is a Unix timestamp in ms
	To *uint64 `json:"To,omitnil,omitempty" name:"To"`

	// Page offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of entries per page. Maximum value: 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// - alertId: Filter by alarm policy ID. Type: String; optional
	// - topicId: Filter by ID of monitored object. Type: String; optional
	// - status: Filter by alarm status. Type: String, optional. Valid values: `0` (uncleared), `1` (cleared), `2` (expired)
	// - alarmLevel: Filter by alarm severity. Type: String, optional. Valid values: `0` (Warn), `1` (Info), `2` (Critical)
	// 
	// Each request can have up to 10 `Filters` and 100 `Filter.Values`.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeAlertRecordHistoryRequest struct {
	*tchttp.BaseRequest
	
	// Start time of the query range, which is a Unix timestamp in ms
	From *uint64 `json:"From,omitnil,omitempty" name:"From"`

	// End time of the query range, which is a Unix timestamp in ms
	To *uint64 `json:"To,omitnil,omitempty" name:"To"`

	// Page offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of entries per page. Maximum value: 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// - alertId: Filter by alarm policy ID. Type: String; optional
	// - topicId: Filter by ID of monitored object. Type: String; optional
	// - status: Filter by alarm status. Type: String, optional. Valid values: `0` (uncleared), `1` (cleared), `2` (expired)
	// - alarmLevel: Filter by alarm severity. Type: String, optional. Valid values: `0` (Warn), `1` (Info), `2` (Critical)
	// 
	// Each request can have up to 10 `Filters` and 100 `Filter.Values`.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeAlertRecordHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlertRecordHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "From")
	delete(f, "To")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlertRecordHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlertRecordHistoryResponseParams struct {
	// Total alarm records
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Alarm record details
	Records []*AlertHistoryRecord `json:"Records,omitnil,omitempty" name:"Records"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAlertRecordHistoryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlertRecordHistoryResponseParams `json:"Response"`
}

func (r *DescribeAlertRecordHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlertRecordHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigMachineGroupsRequestParams struct {
	// Collection configuration ID
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`
}

type DescribeConfigMachineGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Collection configuration ID
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`
}

func (r *DescribeConfigMachineGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigMachineGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigMachineGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigMachineGroupsResponseParams struct {
	// List of machine groups bound to the collection rule configuration
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	MachineGroups []*MachineGroupInfo `json:"MachineGroups,omitnil,omitempty" name:"MachineGroups"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConfigMachineGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConfigMachineGroupsResponseParams `json:"Response"`
}

func (r *DescribeConfigMachineGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigMachineGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigsRequestParams struct {
	// configName
	// - Filter by fuzzy match of **collection configuration name**
	// - Type: String
	// - Required: No
	// 
	// configId
	// - Filter by **collection configuration ID**
	// - Type: String
	// - Required: No
	// 
	// topicId
	// - Filter by **log topic**
	// - Type: String
	// - Required: No
	// 
	// Each request can contain up to 10 `Filters` and 5 `Filter.Values`.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Page offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of entries per page. Default value: 20. Maximum value: 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeConfigsRequest struct {
	*tchttp.BaseRequest
	
	// configName
	// - Filter by fuzzy match of **collection configuration name**
	// - Type: String
	// - Required: No
	// 
	// configId
	// - Filter by **collection configuration ID**
	// - Type: String
	// - Required: No
	// 
	// topicId
	// - Filter by **log topic**
	// - Type: String
	// - Required: No
	// 
	// Each request can contain up to 10 `Filters` and 5 `Filter.Values`.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Page offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of entries per page. Default value: 20. Maximum value: 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigsResponseParams struct {
	// Collection configuration list
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Configs []*ConfigInfo `json:"Configs,omitnil,omitempty" name:"Configs"`

	// Total number of filtered items
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConfigsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConfigsResponseParams `json:"Response"`
}

func (r *DescribeConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConsumerRequestParams struct {
	// Log topic ID bound to the task
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

type DescribeConsumerRequest struct {
	*tchttp.BaseRequest
	
	// Log topic ID bound to the task
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

func (r *DescribeConsumerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConsumerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConsumerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConsumerResponseParams struct {
	// Whether the shipping task is effective
	Effective *bool `json:"Effective,omitnil,omitempty" name:"Effective"`

	// Whether log metadata is shipped
	NeedContent *bool `json:"NeedContent,omitnil,omitempty" name:"NeedContent"`

	// Metadata shipped if `NeedContent` is `true`
	// Note: This field may return `null`, indicating that no valid value was found.
	Content *ConsumerContent `json:"Content,omitnil,omitempty" name:"Content"`

	// CKafka information
	Ckafka *Ckafka `json:"Ckafka,omitnil,omitempty" name:"Ckafka"`

	// Compression mode. Valid values: `0` (no compression), `2` (snappy), `3` (LZ4).
	// Note: This field may return null, indicating that no valid values can be obtained.
	Compression *int64 `json:"Compression,omitnil,omitempty" name:"Compression"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConsumerResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConsumerResponseParams `json:"Response"`
}

func (r *DescribeConsumerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConsumerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosRechargesRequestParams struct {
	// ID of the log topic.
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Status. `0`: Created, `1`: Running, `2`: Stopped, `3`: Completed, `4`: Run failed
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Whether the configuration is enabled. `0`: Not enabled, `1`: Enabled
	Enable *uint64 `json:"Enable,omitnil,omitempty" name:"Enable"`
}

type DescribeCosRechargesRequest struct {
	*tchttp.BaseRequest
	
	// ID of the log topic.
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Status. `0`: Created, `1`: Running, `2`: Stopped, `3`: Completed, `4`: Run failed
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Whether the configuration is enabled. `0`: Not enabled, `1`: Enabled
	Enable *uint64 `json:"Enable,omitnil,omitempty" name:"Enable"`
}

func (r *DescribeCosRechargesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosRechargesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "Status")
	delete(f, "Enable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCosRechargesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosRechargesResponseParams struct {
	// See the description of the `CosRechargeInfo` structure.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*CosRechargeInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCosRechargesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCosRechargesResponseParams `json:"Response"`
}

func (r *DescribeCosRechargesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosRechargesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataTransformInfoRequestParams struct {
	// <br><li>taskName
	// 
	// Filter by **processing task name**.
	// Type: String
	// 
	// Required: No
	// 
	// <br><li>taskId
	// 
	// Filter by **processing task ID**.
	// Type: String
	// 
	// Required: No
	// 
	// <br><li>srctopicId
	// 
	// Filter by **source topic ID**.
	// Type: String
	// 
	// Required: No
	// 
	// Each request can have up to 10 `Filters` and 100 `Filter.Values`.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The pagination offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of entries per page. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Task type. Valid values: 1: Get the details of a single task. 2 (default): Get the task list.
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Task ID, which is required when `Type` is set to `1`
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeDataTransformInfoRequest struct {
	*tchttp.BaseRequest
	
	// <br><li>taskName
	// 
	// Filter by **processing task name**.
	// Type: String
	// 
	// Required: No
	// 
	// <br><li>taskId
	// 
	// Filter by **processing task ID**.
	// Type: String
	// 
	// Required: No
	// 
	// <br><li>srctopicId
	// 
	// Filter by **source topic ID**.
	// Type: String
	// 
	// Required: No
	// 
	// Each request can have up to 10 `Filters` and 100 `Filter.Values`.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The pagination offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of entries per page. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Task type. Valid values: 1: Get the details of a single task. 2 (default): Get the task list.
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Task ID, which is required when `Type` is set to `1`
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DescribeDataTransformInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataTransformInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Type")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataTransformInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataTransformInfoResponseParams struct {
	// List of data processing tasks
	DataTransformTaskInfos []*DataTransformTaskInfo `json:"DataTransformTaskInfos,omitnil,omitempty" name:"DataTransformTaskInfos"`

	// Total tasks
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataTransformInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataTransformInfoResponseParams `json:"Response"`
}

func (r *DescribeDataTransformInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataTransformInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExportsRequestParams struct {
	// Log topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Page offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of entries per page. Default value: 20. Maximum value: 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeExportsRequest struct {
	*tchttp.BaseRequest
	
	// Log topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Page offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of entries per page. Default value: 20. Maximum value: 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeExportsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExportsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExportsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExportsResponseParams struct {
	// List of exported logs
	Exports []*ExportInfo `json:"Exports,omitnil,omitempty" name:"Exports"`

	// Total number
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeExportsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExportsResponseParams `json:"Response"`
}

func (r *DescribeExportsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExportsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIndexRequestParams struct {
	// Log topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

type DescribeIndexRequest struct {
	*tchttp.BaseRequest
	
	// Log topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

func (r *DescribeIndexRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIndexRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIndexRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIndexResponseParams struct {
	// Log topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Whether it takes effect
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// Index configuration information
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Rule *RuleInfo `json:"Rule,omitnil,omitempty" name:"Rule"`

	// Index modification time. The default value is the index creation time.
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// Whether full-text indexing includes internal fields (`__FILENAME__`, `__HOSTNAME__`, and `__SOURCE__`)
	// * `false`: Full-text indexing does not include internal fields.
	// * `true`: Full-text indexing includes internal fields.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IncludeInternalFields *bool `json:"IncludeInternalFields,omitnil,omitempty" name:"IncludeInternalFields"`

	// Whether full-text indexing includes metadata fields (which are prefixed with `__TAG__`)
	// * `0`: Full-text indexing includes only the metadata fields with key-value indexing enabled.
	// * `1`: Full-text indexing includes all metadata fields.
	// * `2`: Full-text indexing does not include metadata fields.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MetadataFlag *uint64 `json:"MetadataFlag,omitnil,omitempty" name:"MetadataFlag"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIndexResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIndexResponseParams `json:"Response"`
}

func (r *DescribeIndexResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKafkaRechargesRequestParams struct {
	// Log topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Import configuration ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Status. Valid values: 1 (running) and 2 (suspended).
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type DescribeKafkaRechargesRequest struct {
	*tchttp.BaseRequest
	
	// Log topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Import configuration ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Status. Valid values: 1 (running) and 2 (suspended).
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *DescribeKafkaRechargesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKafkaRechargesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "Id")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKafkaRechargesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKafkaRechargesResponseParams struct {
	// KafkaRechargeInfo list
	Infos []*KafkaRechargeInfo `json:"Infos,omitnil,omitempty" name:"Infos"`

	// Total Kafka data records imported
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeKafkaRechargesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKafkaRechargesResponseParams `json:"Response"`
}

func (r *DescribeKafkaRechargesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKafkaRechargesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogContextRequestParams struct {
	// Log topic ID to be queried
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Log time in the format of YYYY-mm-dd HH:MM:SS.FFF
	BTime *string `json:"BTime,omitnil,omitempty" name:"BTime"`

	// Log package number
	PkgId *string `json:"PkgId,omitnil,omitempty" name:"PkgId"`

	// Log number in log package
	PkgLogId *int64 `json:"PkgLogId,omitnil,omitempty" name:"PkgLogId"`

	// Number of previous logs. Default value: 10
	PrevLogs *int64 `json:"PrevLogs,omitnil,omitempty" name:"PrevLogs"`

	// Number of next logs. Default value: 10
	NextLogs *int64 `json:"NextLogs,omitnil,omitempty" name:"NextLogs"`
}

type DescribeLogContextRequest struct {
	*tchttp.BaseRequest
	
	// Log topic ID to be queried
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Log time in the format of YYYY-mm-dd HH:MM:SS.FFF
	BTime *string `json:"BTime,omitnil,omitempty" name:"BTime"`

	// Log package number
	PkgId *string `json:"PkgId,omitnil,omitempty" name:"PkgId"`

	// Log number in log package
	PkgLogId *int64 `json:"PkgLogId,omitnil,omitempty" name:"PkgLogId"`

	// Number of previous logs. Default value: 10
	PrevLogs *int64 `json:"PrevLogs,omitnil,omitempty" name:"PrevLogs"`

	// Number of next logs. Default value: 10
	NextLogs *int64 `json:"NextLogs,omitnil,omitempty" name:"NextLogs"`
}

func (r *DescribeLogContextRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogContextRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "BTime")
	delete(f, "PkgId")
	delete(f, "PkgLogId")
	delete(f, "PrevLogs")
	delete(f, "NextLogs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogContextRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogContextResponseParams struct {
	// Log context information set
	LogContextInfos []*LogContextInfo `json:"LogContextInfos,omitnil,omitempty" name:"LogContextInfos"`

	// Whether the previous logs have been returned
	PrevOver *bool `json:"PrevOver,omitnil,omitempty" name:"PrevOver"`

	// Whether the next logs have been returned
	NextOver *bool `json:"NextOver,omitnil,omitempty" name:"NextOver"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLogContextResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLogContextResponseParams `json:"Response"`
}

func (r *DescribeLogContextResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogContextResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogHistogramRequestParams struct {
	// Start time of the log to be queried, which is a Unix timestamp in milliseconds
	From *int64 `json:"From,omitnil,omitempty" name:"From"`

	// End time of the log to be queried, which is a Unix timestamp in milliseconds
	To *int64 `json:"To,omitnil,omitempty" name:"To"`

	// Query statement
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// ID of the log topic to be queried
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Interval in milliseconds. Condition: (To – From) / Interval ≤ 200
	Interval *int64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// Search syntax. Valid values:
	// `0` (default): Lucene; `1`: CQL
	// For more information, see <a href="https://intl.cloud.tencent.com/document/product/614/47044?from_cn_redirect=1#RetrievesConditionalRules" target="_blank">Search Syntax</a>.
	SyntaxRule *uint64 `json:"SyntaxRule,omitnil,omitempty" name:"SyntaxRule"`
}

type DescribeLogHistogramRequest struct {
	*tchttp.BaseRequest
	
	// Start time of the log to be queried, which is a Unix timestamp in milliseconds
	From *int64 `json:"From,omitnil,omitempty" name:"From"`

	// End time of the log to be queried, which is a Unix timestamp in milliseconds
	To *int64 `json:"To,omitnil,omitempty" name:"To"`

	// Query statement
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// ID of the log topic to be queried
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Interval in milliseconds. Condition: (To – From) / Interval ≤ 200
	Interval *int64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// Search syntax. Valid values:
	// `0` (default): Lucene; `1`: CQL
	// For more information, see <a href="https://intl.cloud.tencent.com/document/product/614/47044?from_cn_redirect=1#RetrievesConditionalRules" target="_blank">Search Syntax</a>.
	SyntaxRule *uint64 `json:"SyntaxRule,omitnil,omitempty" name:"SyntaxRule"`
}

func (r *DescribeLogHistogramRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogHistogramRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "From")
	delete(f, "To")
	delete(f, "Query")
	delete(f, "TopicId")
	delete(f, "Interval")
	delete(f, "SyntaxRule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogHistogramRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogHistogramResponseParams struct {
	// Statistical period in milliseconds
	Interval *int64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// The number of logs that hit the keywords
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Statistical result details within the period
	HistogramInfos []*HistogramInfo `json:"HistogramInfos,omitnil,omitempty" name:"HistogramInfos"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLogHistogramResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLogHistogramResponseParams `json:"Response"`
}

func (r *DescribeLogHistogramResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogHistogramResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogsetsRequestParams struct {
	// logsetName
	// - Filter by **logset name**
	// - Type: String
	// - Required: No
	// 
	// logsetId
	// - Filter by **logset ID**
	// - Type: String
	// - Required: No
	// 
	// tagKey
	// - Filter by **tag key**
	// - Type: String
	// - Required: No
	// 
	// tag:tagKey
	// - Filter by **tag key-value pair**. The `tagKey` should be replaced with a specified tag key.
	// - Type: String
	// - Required: No
	// 
	// Each request can have up to 10 `Filters` and 5 `Filter.Values`.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Page offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of entries per page. Default value: 20. Maximum value: 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeLogsetsRequest struct {
	*tchttp.BaseRequest
	
	// logsetName
	// - Filter by **logset name**
	// - Type: String
	// - Required: No
	// 
	// logsetId
	// - Filter by **logset ID**
	// - Type: String
	// - Required: No
	// 
	// tagKey
	// - Filter by **tag key**
	// - Type: String
	// - Required: No
	// 
	// tag:tagKey
	// - Filter by **tag key-value pair**. The `tagKey` should be replaced with a specified tag key.
	// - Type: String
	// - Required: No
	// 
	// Each request can have up to 10 `Filters` and 5 `Filter.Values`.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Page offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of entries per page. Default value: 20. Maximum value: 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeLogsetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogsetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogsetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogsetsResponseParams struct {
	// Total number of pages
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Logset list
	Logsets []*LogsetInfo `json:"Logsets,omitnil,omitempty" name:"Logsets"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLogsetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLogsetsResponseParams `json:"Response"`
}

func (r *DescribeLogsetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogsetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMachineGroupConfigsRequestParams struct {
	// Machine group ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

type DescribeMachineGroupConfigsRequest struct {
	*tchttp.BaseRequest
	
	// Machine group ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

func (r *DescribeMachineGroupConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineGroupConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMachineGroupConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMachineGroupConfigsResponseParams struct {
	// Collection rule configuration list
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Configs []*ConfigInfo `json:"Configs,omitnil,omitempty" name:"Configs"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMachineGroupConfigsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMachineGroupConfigsResponseParams `json:"Response"`
}

func (r *DescribeMachineGroupConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineGroupConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMachineGroupsRequestParams struct {
	// machineGroupName
	// - Filter by **machine group name**
	// - Type: String
	// - Required: No
	// 
	// machineGroupId
	// - Filter by **machine group ID**
	// - Type: String
	// - Required: No
	// 
	// tagKey
	// - Filter by **tag key**
	// - Type: String
	// - Required: No
	// 
	// tag:tagKey
	// - Filter by **tag key-value pair**. The `tagKey` should be replaced with a specified tag key.
	// - Type: String
	// - Required: No
	// 
	// Each request can have up to 10 `Filters` and 100 `Filter.Values`.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Page offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of entries per page. Default value: 20. Maximum value: 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeMachineGroupsRequest struct {
	*tchttp.BaseRequest
	
	// machineGroupName
	// - Filter by **machine group name**
	// - Type: String
	// - Required: No
	// 
	// machineGroupId
	// - Filter by **machine group ID**
	// - Type: String
	// - Required: No
	// 
	// tagKey
	// - Filter by **tag key**
	// - Type: String
	// - Required: No
	// 
	// tag:tagKey
	// - Filter by **tag key-value pair**. The `tagKey` should be replaced with a specified tag key.
	// - Type: String
	// - Required: No
	// 
	// Each request can have up to 10 `Filters` and 100 `Filter.Values`.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Page offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of entries per page. Default value: 20. Maximum value: 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeMachineGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMachineGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMachineGroupsResponseParams struct {
	// Machine group information list
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	MachineGroups []*MachineGroupInfo `json:"MachineGroups,omitnil,omitempty" name:"MachineGroups"`

	// Total number of pages
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMachineGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMachineGroupsResponseParams `json:"Response"`
}

func (r *DescribeMachineGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMachinesRequestParams struct {
	// ID of the machine group to be queried
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

type DescribeMachinesRequest struct {
	*tchttp.BaseRequest
	
	// ID of the machine group to be queried
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

func (r *DescribeMachinesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachinesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMachinesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMachinesResponseParams struct {
	// Group of machine status information
	Machines []*MachineInfo `json:"Machines,omitnil,omitempty" name:"Machines"`

	// Whether to enable the automatic update feature for the machine group
	AutoUpdate *int64 `json:"AutoUpdate,omitnil,omitempty" name:"AutoUpdate"`

	// Preset start time of automatic update of machine group
	UpdateStartTime *string `json:"UpdateStartTime,omitnil,omitempty" name:"UpdateStartTime"`

	// Preset end time of automatic update of machine group
	UpdateEndTime *string `json:"UpdateEndTime,omitnil,omitempty" name:"UpdateEndTime"`

	// Latest LogListener version available to the current user
	LatestAgentVersion *string `json:"LatestAgentVersion,omitnil,omitempty" name:"LatestAgentVersion"`

	// Whether to enable the service log
	ServiceLogging *bool `json:"ServiceLogging,omitnil,omitempty" name:"ServiceLogging"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMachinesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMachinesResponseParams `json:"Response"`
}

func (r *DescribeMachinesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachinesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePartitionsRequestParams struct {
	// Log topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

type DescribePartitionsRequest struct {
	*tchttp.BaseRequest
	
	// Log topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

func (r *DescribePartitionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePartitionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePartitionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePartitionsResponseParams struct {
	// Partition list
	Partitions []*PartitionInfo `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePartitionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePartitionsResponseParams `json:"Response"`
}

func (r *DescribePartitionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePartitionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShipperTasksRequestParams struct {
	// Shipping rule ID
	ShipperId *string `json:"ShipperId,omitnil,omitempty" name:"ShipperId"`

	// Query start timestamp in milliseconds, which can be within the last three days
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Query end timestamp in milliseconds
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeShipperTasksRequest struct {
	*tchttp.BaseRequest
	
	// Shipping rule ID
	ShipperId *string `json:"ShipperId,omitnil,omitempty" name:"ShipperId"`

	// Query start timestamp in milliseconds, which can be within the last three days
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Query end timestamp in milliseconds
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *DescribeShipperTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShipperTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ShipperId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeShipperTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShipperTasksResponseParams struct {
	// Shipping task list
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Tasks []*ShipperTaskInfo `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeShipperTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeShipperTasksResponseParams `json:"Response"`
}

func (r *DescribeShipperTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShipperTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShippersRequestParams struct {
	// - shipperName: Filter by **shipping rule name**. Type: String. Required: No.
	// - shipperId: Filter by **shipping rule ID**. Type: String. Required: No.
	// - topicId: Filter by **log topic**. Type: String. Required: No.
	// 
	// Each request can have up to 10 `Filters` and 100 `Filter.Values`.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Page offset. Default value: 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of entries per page. Default value: 20. Maximum value: 100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeShippersRequest struct {
	*tchttp.BaseRequest
	
	// - shipperName: Filter by **shipping rule name**. Type: String. Required: No.
	// - shipperId: Filter by **shipping rule ID**. Type: String. Required: No.
	// - topicId: Filter by **log topic**. Type: String. Required: No.
	// 
	// Each request can have up to 10 `Filters` and 100 `Filter.Values`.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Page offset. Default value: 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of entries per page. Default value: 20. Maximum value: 100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeShippersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShippersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeShippersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShippersResponseParams struct {
	// Shipping rule list
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Shippers []*ShipperInfo `json:"Shippers,omitnil,omitempty" name:"Shippers"`

	// Total number of results obtained in this query
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeShippersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeShippersResponseParams `json:"Response"`
}

func (r *DescribeShippersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShippersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicsRequestParams struct {
	// <li>topicName: Filter by **log topic name**. Fuzzy match is implemented by default. You can use the `PreciseSearch` parameter to set exact match. Type: String. Required. No. <br><li>logsetName: Filter by **logset name**. Fuzzy match is implemented by default. You can use the `PreciseSearch` parameter to set exact match. Type: String. Required: No. <br><li>topicId: Filter by **log topic ID**. Type: String. Required: No. <br><li>logsetId: Filter by **logset ID**. You can call `DescribeLogsets` to query the list of created logsets or log in to the console to view them. You can also call `CreateLogset` to create a logset. Type: String. Required: No. <br><li>tagKey: Filter by **tag key**. Type: String. Required: No. <br><li>tag:tagKey: Filter by **tag key-value pair**. The `tagKey` should be replaced with a specified tag key, such as `tag:exampleKey`. Type: String. Required: No. <br><li>storageType: Filter by **log topic storage type**. Valid values: `hot` (standard storage) and `cold` (IA storage). Type: String. Required: No. Each request can have up to 10 `Filters` and 100 `Filter.Values`.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Page offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of entries per page. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Match mode for `Filters` fields.
	// - 0: Fuzzy match for `topicName` and `logsetName`. This is the default value.
	// - 1: Exact match for `topicName`.
	// - 2: Exact match for `logsetName`.
	// - 3: Exact match for `topicName` and `logsetName`.
	PreciseSearch *uint64 `json:"PreciseSearch,omitnil,omitempty" name:"PreciseSearch"`

	// Topic type
	// - 0 (default): Log topic.
	// - 1: Metric topic.
	BizType *uint64 `json:"BizType,omitnil,omitempty" name:"BizType"`
}

type DescribeTopicsRequest struct {
	*tchttp.BaseRequest
	
	// <li>topicName: Filter by **log topic name**. Fuzzy match is implemented by default. You can use the `PreciseSearch` parameter to set exact match. Type: String. Required. No. <br><li>logsetName: Filter by **logset name**. Fuzzy match is implemented by default. You can use the `PreciseSearch` parameter to set exact match. Type: String. Required: No. <br><li>topicId: Filter by **log topic ID**. Type: String. Required: No. <br><li>logsetId: Filter by **logset ID**. You can call `DescribeLogsets` to query the list of created logsets or log in to the console to view them. You can also call `CreateLogset` to create a logset. Type: String. Required: No. <br><li>tagKey: Filter by **tag key**. Type: String. Required: No. <br><li>tag:tagKey: Filter by **tag key-value pair**. The `tagKey` should be replaced with a specified tag key, such as `tag:exampleKey`. Type: String. Required: No. <br><li>storageType: Filter by **log topic storage type**. Valid values: `hot` (standard storage) and `cold` (IA storage). Type: String. Required: No. Each request can have up to 10 `Filters` and 100 `Filter.Values`.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Page offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of entries per page. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Match mode for `Filters` fields.
	// - 0: Fuzzy match for `topicName` and `logsetName`. This is the default value.
	// - 1: Exact match for `topicName`.
	// - 2: Exact match for `logsetName`.
	// - 3: Exact match for `topicName` and `logsetName`.
	PreciseSearch *uint64 `json:"PreciseSearch,omitnil,omitempty" name:"PreciseSearch"`

	// Topic type
	// - 0 (default): Log topic.
	// - 1: Metric topic.
	BizType *uint64 `json:"BizType,omitnil,omitempty" name:"BizType"`
}

func (r *DescribeTopicsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "PreciseSearch")
	delete(f, "BizType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopicsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicsResponseParams struct {
	// Log topic list
	Topics []*TopicInfo `json:"Topics,omitnil,omitempty" name:"Topics"`

	// Total number
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTopicsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopicsResponseParams `json:"Response"`
}

func (r *DescribeTopicsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DynamicIndex struct {
	// Dynamic index configuration status
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

type ExcludePathInfo struct {
	// Type. Valid values: `File`, `Path`
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Specific content corresponding to `Type`
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type ExportInfo struct {
	// Log topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Log export task ID
	ExportId *string `json:"ExportId,omitnil,omitempty" name:"ExportId"`

	// Log export query statement
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// Log export filename
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// Log file size
	FileSize *uint64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// Log export time sorting
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Log export format
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// Number of logs to be exported
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// Log download status. Valid values: `Processing`, `Completed`, `Failed`, `Expired` (three-day validity period), and `Queuing`.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Log export start time
	From *int64 `json:"From,omitnil,omitempty" name:"From"`

	// Log export end time
	To *int64 `json:"To,omitnil,omitempty" name:"To"`

	// Log export path
	CosPath *string `json:"CosPath,omitnil,omitempty" name:"CosPath"`

	// Log export creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type ExtractRuleInfo struct {
	// Time field key name. `time_key` and `time_format` must appear in pairs
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TimeKey *string `json:"TimeKey,omitnil,omitempty" name:"TimeKey"`

	// Time field format. For more information, please see the output parameters of the time format description of the `strftime` function in C language
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TimeFormat *string `json:"TimeFormat,omitnil,omitempty" name:"TimeFormat"`

	// Delimiter for delimited log, which is valid only if `log_type` is `delimiter_log`
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Delimiter *string `json:"Delimiter,omitnil,omitempty" name:"Delimiter"`

	// Full log matching rule, which is valid only if `log_type` is `fullregex_log`
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	LogRegex *string `json:"LogRegex,omitnil,omitempty" name:"LogRegex"`

	// First-Line matching rule, which is valid only if `log_type` is `multiline_log` or `fullregex_log`
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	BeginRegex *string `json:"BeginRegex,omitnil,omitempty" name:"BeginRegex"`

	// Key name of each extracted field. An empty key indicates to discard the field. This parameter is valid only if `log_type` is `delimiter_log`. `json_log` logs use the key of JSON itself. A maximum of 100 keys are supported.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Keys []*string `json:"Keys,omitnil,omitempty" name:"Keys"`

	// Log keys to be filtered and the corresponding regex
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	FilterKeyRegex []*KeyRegexInfo `json:"FilterKeyRegex,omitnil,omitempty" name:"FilterKeyRegex"`

	// Whether to upload the logs that failed to be parsed. Valid values: `true`: yes; `false`: no
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	UnMatchUpLoadSwitch *bool `json:"UnMatchUpLoadSwitch,omitnil,omitempty" name:"UnMatchUpLoadSwitch"`

	// Unmatched log key
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	UnMatchLogKey *string `json:"UnMatchLogKey,omitnil,omitempty" name:"UnMatchLogKey"`

	// Size of the data to be rewound in incremental collection mode. Default value: -1 (full collection)
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Backtracking *int64 `json:"Backtracking,omitnil,omitempty" name:"Backtracking"`

	// Whether to be encoded in GBK format. Valid values: `0` (No) and `1` (Yes).
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsGBK *int64 `json:"IsGBK,omitnil,omitempty" name:"IsGBK"`

	// Whether to be formatted as JSON (standard). Valid values: `0` (No) and `1` (Yes).
	// Note: This field may return null, indicating that no valid values can be obtained.
	JsonStandard *int64 `json:"JsonStandard,omitnil,omitempty" name:"JsonStandard"`

	// Syslog protocol. Valid values: `tcp`, `udp`.
	// This field can be used when you create or modify collection rule configurations.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Listening address and port specified by the syslog collection. Format: [ip]:[port]. Example: 127.0.0.1:9000.
	// This field can be used when you create or modify collection rule configurations.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// `rfc3164`: Resolve logs by using the RFC 3164 protocol during the syslog collection.
	// `rfc5424`: Resolve logs by using the RFC 5424 protocol during the syslog collection.
	// `auto`: Automatically match either the RFC 3164 or RFC 5424 protocol.
	// This field can be used when you create or modify collection rule configurations.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ParseProtocol *string `json:"ParseProtocol,omitnil,omitempty" name:"ParseProtocol"`

	// Metadata type. Valid values:
	// 0: Do not use metadata.
	// 1: Use machine group metadata.
	// 2: Use user-defined metadata.
	// 3: Use the collection path to extract metadata.
	MetadataType *int64 `json:"MetadataType,omitnil,omitempty" name:"MetadataType"`

	// Regular expression of the collection configuration path, which is required when `MetadataType` is set to `3`
	// Note: This field may return null, indicating that no valid values can be obtained.
	PathRegex *string `json:"PathRegex,omitnil,omitempty" name:"PathRegex"`

	// User-defined metadata, which is required when `MetadataType` is set to `2`.
	MetaTags []*MetaTagInfo `json:"MetaTags,omitnil,omitempty" name:"MetaTags"`
}

type Filter struct {
	// Field to be filtered
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Value to be filtered
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type FilterRuleInfo struct {
	// Filter rule key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Filter rule
	Regex *string `json:"Regex,omitnil,omitempty" name:"Regex"`

	// Filter rule value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type FullTextInfo struct {
	// Case sensitivity
	CaseSensitive *bool `json:"CaseSensitive,omitnil,omitempty" name:"CaseSensitive"`

	// Separator of the full-text index. Each character represents a separator.
	// Only symbols, \n\t\r, and escape character \ are supported.
	// Note: \n\t\r can be directly enclosed in double quotes as the input parameter without escaping. When debugging with API Explorer, use the JSON parameter input method to avoid repeated escaping of \n\t\r.
	Tokenizer *string `json:"Tokenizer,omitnil,omitempty" name:"Tokenizer"`

	// Whether Chinese characters are contained
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ContainZH *bool `json:"ContainZH,omitnil,omitempty" name:"ContainZH"`
}

// Predefined struct for user
type GetAlarmLogRequestParams struct {
	// Start time of the log to be queried, which is a Unix timestamp in milliseconds
	From *int64 `json:"From,omitnil,omitempty" name:"From"`

	// End time of the log to be queried, which is a Unix timestamp in milliseconds
	To *int64 `json:"To,omitnil,omitempty" name:"To"`

	// Query statement. Maximum length: 1024
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// Number of logs returned in a single query. Maximum value: 1000
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// This field is used to load more logs. Pass through the last `Context` value returned to get more log content.
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// Order of the logs sorted by time returned by the log API. Valid values: `asc`: ascending; `desc`: descending. Default value: `desc`
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// If the value is `true`, the new search method will be used, and the response parameters `AnalysisRecords` and `Columns` will be valid. If the value is `false`, the old search method will be used, and `AnalysisResults` and `ColNames` will be valid.
	UseNewAnalysis *bool `json:"UseNewAnalysis,omitnil,omitempty" name:"UseNewAnalysis"`
}

type GetAlarmLogRequest struct {
	*tchttp.BaseRequest
	
	// Start time of the log to be queried, which is a Unix timestamp in milliseconds
	From *int64 `json:"From,omitnil,omitempty" name:"From"`

	// End time of the log to be queried, which is a Unix timestamp in milliseconds
	To *int64 `json:"To,omitnil,omitempty" name:"To"`

	// Query statement. Maximum length: 1024
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// Number of logs returned in a single query. Maximum value: 1000
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// This field is used to load more logs. Pass through the last `Context` value returned to get more log content.
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// Order of the logs sorted by time returned by the log API. Valid values: `asc`: ascending; `desc`: descending. Default value: `desc`
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// If the value is `true`, the new search method will be used, and the response parameters `AnalysisRecords` and `Columns` will be valid. If the value is `false`, the old search method will be used, and `AnalysisResults` and `ColNames` will be valid.
	UseNewAnalysis *bool `json:"UseNewAnalysis,omitnil,omitempty" name:"UseNewAnalysis"`
}

func (r *GetAlarmLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAlarmLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "From")
	delete(f, "To")
	delete(f, "Query")
	delete(f, "Limit")
	delete(f, "Context")
	delete(f, "Sort")
	delete(f, "UseNewAnalysis")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAlarmLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAlarmLogResponseParams struct {
	// `Context` for loading subsequent content
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// Whether all log query results are returned
	ListOver *bool `json:"ListOver,omitnil,omitempty" name:"ListOver"`

	// Whether the return is the analysis result
	Analysis *bool `json:"Analysis,omitnil,omitempty" name:"Analysis"`

	// If `Analysis` is `true`, column name of the analysis result will be returned; otherwise, empty content will be returned.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ColNames []*string `json:"ColNames,omitnil,omitempty" name:"ColNames"`

	// Log query result. If `Analysis` is `True`, `null` may be returned
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Results []*LogInfo `json:"Results,omitnil,omitempty" name:"Results"`

	// Log analysis result. If `Analysis` is `False`, `null` may be returned
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	AnalysisResults []*LogItems `json:"AnalysisResults,omitnil,omitempty" name:"AnalysisResults"`

	// New log analysis result, which will be valid if `UseNewAnalysis` is `true`
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	AnalysisRecords []*string `json:"AnalysisRecords,omitnil,omitempty" name:"AnalysisRecords"`

	// Column attribute of log analysis, which will be valid if `UseNewAnalysis` is `true`
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Columns []*Column `json:"Columns,omitnil,omitempty" name:"Columns"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetAlarmLogResponse struct {
	*tchttp.BaseResponse
	Response *GetAlarmLogResponseParams `json:"Response"`
}

func (r *GetAlarmLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAlarmLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GroupTriggerConditionInfo struct {
	// Name of the field for triggering by group
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Value of the field for triggering by group
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type HistogramInfo struct {
	// The number of logs within the statistical period
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// Unix timestamp rounded by `period`, in milliseconds
	BTime *int64 `json:"BTime,omitnil,omitempty" name:"BTime"`
}

type JsonInfo struct {
	// Enablement flag
	EnableTag *bool `json:"EnableTag,omitnil,omitempty" name:"EnableTag"`

	// List of metadata. Supported metadata types: __SOURCE__, __FILENAME__, __TIMESTAMP__, __HOSTNAME__.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MetaFields []*string `json:"MetaFields,omitnil,omitempty" name:"MetaFields"`

	// JSON format for shipping. `0`: String format; `1`: Structured format.
	// Note: This field may return null, indicating that no valid values can be obtained.
	JsonType *int64 `json:"JsonType,omitnil,omitempty" name:"JsonType"`
}

type KafkaConsumerContent struct {
	// Format. Valid values: 0 (full-text) and 1 (JSON).
	Format *int64 `json:"Format,omitnil,omitempty" name:"Format"`

	// Whether to ship tag information
	// This parameter does not need to be set when `Format` is set to `0`.
	EnableTag *bool `json:"EnableTag,omitnil,omitempty" name:"EnableTag"`

	// Metadata information list. Valid values: \_\_SOURCE\_\_, \_\_FILENAME\_\_,
	// \_\_TIMESTAMP\_\_, \_\_HOSTNAME\_\_, and \_\_PKGID\_\_.
	// This parameter does not need to be set when `Format` is set to `0`.
	MetaFields []*string `json:"MetaFields,omitnil,omitempty" name:"MetaFields"`

	// Tag data processing mode. Valid values:
	// 1 (default): Do not tile data.
	// 2: Tile data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TagTransaction *int64 `json:"TagTransaction,omitnil,omitempty" name:"TagTransaction"`

	// JSON data format. Valid values:
	// 1 (default): Not escaped.
	// 2: Escaped.
	JsonType *int64 `json:"JsonType,omitnil,omitempty" name:"JsonType"`
}

type KafkaProtocolInfo struct {
	// Protocol type. Valid values: `plaintext`, `sasl_plaintext`, and `sasl_ssl`. `sasl_ssl` is recommended. Using this protocol will encrypt the connection and implement user authentication.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Encryption type. Valid values: `PLAIN`, `SCRAM-SHA-256`, and SCRAM-SHA-512`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Mechanism *string `json:"Mechanism,omitnil,omitempty" name:"Mechanism"`

	// Username
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// User password
	// Note: This field may return null, indicating that no valid values can be obtained.
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

type KafkaRechargeInfo struct {
	// Primary key ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Log topic ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Kafka data import task name
	// Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Kafka type. Valid values: 0 (Tencent Cloud CKafka) and 1 (customer's Kafka).
	// Note: This field may return null, indicating that no valid values can be obtained.
	KafkaType *uint64 `json:"KafkaType,omitnil,omitempty" name:"KafkaType"`

	// CKafka instance ID, which is required when `KafkaType` is set to `0`
	// Note: This field may return null, indicating that no valid values can be obtained.
	KafkaInstance *string `json:"KafkaInstance,omitnil,omitempty" name:"KafkaInstance"`

	// Service address
	// Note: This field may return null, indicating that no valid values can be obtained.
	ServerAddr *string `json:"ServerAddr,omitnil,omitempty" name:"ServerAddr"`

	// Whether the service address uses an encrypted connection	
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsEncryptionAddr *bool `json:"IsEncryptionAddr,omitnil,omitempty" name:"IsEncryptionAddr"`

	// Encryption access protocol, which is required when `IsEncryptionAddr` is set to `true`
	Protocol *KafkaProtocolInfo `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// List of Kafka topics to import data from. Separate multiple topics with commas (,).
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserKafkaTopics *string `json:"UserKafkaTopics,omitnil,omitempty" name:"UserKafkaTopics"`

	// Kafka consumer group name	
	// Note: This field may return null, indicating that no valid values can be obtained.
	ConsumerGroupName *string `json:"ConsumerGroupName,omitnil,omitempty" name:"ConsumerGroupName"`

	// Status. Valid values: 1 (running) and 2 (suspended).
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Position for data import. Valid values: -2 (earliest, default) and -1 (latest).  
	// Note: This field may return null, indicating that no valid values can be obtained.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Creation time
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Update time
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Log import rule
	// Note: This field may return null, indicating that no valid values can be obtained.
	LogRechargeRule *LogRechargeRuleInfo `json:"LogRechargeRule,omitnil,omitempty" name:"LogRechargeRule"`
}

type KeyRegexInfo struct {
	// Log key to be filtered
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Filter rule regex corresponding to key
	Regex *string `json:"Regex,omitnil,omitempty" name:"Regex"`
}

type KeyValueInfo struct {
	// Name of the field for which you want to configure a key-value or metadata field index. The name can contain letters, digits, underscores, and symbols -./@ and cannot start with an underscore.
	// 
	// Note:
	// For a metadata field, set its `Key` to be consistent with the one for log uploading, without prefixing it with `__TAG__.`. `__TAG__.` will be prefixed automatically for display in the console.
	// 2. The total number of keys in key-value indexes (`KeyValue`) and metadata field indexes (`Tag`) cannot exceed 300.
	// 3. The number of levels in `Key` cannot exceed 10. Example: a.b.c.d.e.f.g.h.j.k
	// 4. JSON parent and child fields (such as “a” and “a.b”) cannot be contained at the same time.
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Field index description information
	Value *ValueInfo `json:"Value,omitnil,omitempty" name:"Value"`
}

type LogContextInfo struct {
	// Log source device
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// Collection path
	Filename *string `json:"Filename,omitnil,omitempty" name:"Filename"`

	// Log content
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// Log package number
	PkgId *string `json:"PkgId,omitnil,omitempty" name:"PkgId"`

	// Log number in log package
	PkgLogId *int64 `json:"PkgLogId,omitnil,omitempty" name:"PkgLogId"`

	// Log timestamp
	BTime *int64 `json:"BTime,omitnil,omitempty" name:"BTime"`

	// Source host name of logs
	// Note: This field may return `null`, indicating that no valid value was found.
	HostName *string `json:"HostName,omitnil,omitempty" name:"HostName"`

	// Raw log (this parameter has a value only when an exception occurred while creating indexes for logs).
	// Note: This field may return null, indicating that no valid values can be obtained.
	RawLog *string `json:"RawLog,omitnil,omitempty" name:"RawLog"`

	// The cause of index creation exception (this parameter has a value only when an exception occurred while creating indexes for logs).
	// Note: This field may return null, indicating that no valid values can be obtained.
	IndexStatus *string `json:"IndexStatus,omitnil,omitempty" name:"IndexStatus"`
}

type LogInfo struct {
	// Log time in milliseconds
	Time *int64 `json:"Time,omitnil,omitempty" name:"Time"`

	// Log topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Log topic name
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Log source IP
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// Log filename
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// ID of the request package for log reporting
	PkgId *string `json:"PkgId,omitnil,omitempty" name:"PkgId"`

	// Log ID in request package
	PkgLogId *string `json:"PkgLogId,omitnil,omitempty" name:"PkgLogId"`

	// Serialized JSON string of log content
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	LogJson *string `json:"LogJson,omitnil,omitempty" name:"LogJson"`

	// Source host name of logs
	// Note: This field may return `null`, indicating that no valid value was found.
	HostName *string `json:"HostName,omitnil,omitempty" name:"HostName"`

	// Raw log (this parameter has a value only when an exception occurred while creating indexes for logs).
	// Note: This field may return null, indicating that no valid values can be obtained.
	RawLog *string `json:"RawLog,omitnil,omitempty" name:"RawLog"`

	// The cause of index creation exception (this parameter has a value only when an exception occurred while creating indexes for logs).
	// Note: This field may return null, indicating that no valid values can be obtained.
	IndexStatus *string `json:"IndexStatus,omitnil,omitempty" name:"IndexStatus"`
}

type LogItem struct {
	// Log key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Log value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type LogItems struct {
	// Key-Value pair returned in analysis result
	Data []*LogItem `json:"Data,omitnil,omitempty" name:"Data"`
}

type LogRechargeRuleInfo struct {
	// Import type. Valid values: `json_log` (JSON logs), `minimalist_log` (single-line full text), and fullregex_log u200d(single-line full regex)
	RechargeType *string `json:"RechargeType,omitnil,omitempty" name:"RechargeType"`

	// Encoding format. Valid values: 0 (default, UTF-8) and 1 GBK).
	EncodingFormat *uint64 `json:"EncodingFormat,omitnil,omitempty" name:"EncodingFormat"`

	// Whether to use the default time. Valid values: `true` (default) and `false`.
	DefaultTimeSwitch *bool `json:"DefaultTimeSwitch,omitnil,omitempty" name:"DefaultTimeSwitch"`

	// Full log matching rule, which is valid only if `RechargeType` is `fullregex_log`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LogRegex *string `json:"LogRegex,omitnil,omitempty" name:"LogRegex"`

	// Whether to upload the logs that failed to be parsed. Valid values: `true` and `false`.
	UnMatchLogSwitch *bool `json:"UnMatchLogSwitch,omitnil,omitempty" name:"UnMatchLogSwitch"`

	// Key of the log that failed to be parsed
	// Note: This field may return null, indicating that no valid values can be obtained.
	UnMatchLogKey *string `json:"UnMatchLogKey,omitnil,omitempty" name:"UnMatchLogKey"`

	// Time source of the log that failed to be parsed. Valid values: 0 (current system time) and 1 (Kafka message timestamp).
	// Note: This field may return null, indicating that no valid values can be obtained.
	UnMatchLogTimeSrc *uint64 `json:"UnMatchLogTimeSrc,omitnil,omitempty" name:"UnMatchLogTimeSrc"`

	// Default time source. Valid values: 0 (current system time) and 1 (Kafka message timestamp).
	// Note: This field may return null, indicating that no valid values can be obtained.
	DefaultTimeSrc *uint64 `json:"DefaultTimeSrc,omitnil,omitempty" name:"DefaultTimeSrc"`

	// Time field
	// Note: This field may return null, indicating that no valid values can be obtained.
	TimeKey *string `json:"TimeKey,omitnil,omitempty" name:"TimeKey"`

	// Time regular expression
	// Note: This field may return null, indicating that no valid values can be obtained.
	TimeRegex *string `json:"TimeRegex,omitnil,omitempty" name:"TimeRegex"`

	// Time field format
	// Note: This field may return null, indicating that no valid values can be obtained.
	TimeFormat *string `json:"TimeFormat,omitnil,omitempty" name:"TimeFormat"`

	// Time zone
	// Note: This field may return null, indicating that no valid values can be obtained.
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`

	// Metadata information. Kafka supports import of kafka_topic, kafka_partition, kafka_offset, and kafka_timestamp.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Metadata []*string `json:"Metadata,omitnil,omitempty" name:"Metadata"`

	// List of log keys, which is required when `RechargeType` is set to `full_regex_log`
	// Note: This field may return null, indicating that no valid values can be obtained.
	Keys []*string `json:"Keys,omitnil,omitempty" name:"Keys"`
}

type LogsetInfo struct {
	// Logset ID
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// Logset name
	LogsetName *string `json:"LogsetName,omitnil,omitempty" name:"LogsetName"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Cloud product identifier. If the logset is created by another cloud product, this field returns the name of the cloud product, such as `CDN` or `TKE`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AssumerName *string `json:"AssumerName,omitnil,omitempty" name:"AssumerName"`

	// Tag bound to logset
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Number of log topics in logset
	TopicCount *int64 `json:"TopicCount,omitnil,omitempty" name:"TopicCount"`

	// If `AssumerName` is not empty, it indicates the service provider who creates the logset.
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`
}

type MachineGroupInfo struct {
	// Machine group ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Machine group name
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Machine group type
	MachineGroupType *MachineGroupTypeInfo `json:"MachineGroupType,omitnil,omitempty" name:"MachineGroupType"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// List of tags bound to machine group
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Whether to enable automatic update for the machine group
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	AutoUpdate *string `json:"AutoUpdate,omitnil,omitempty" name:"AutoUpdate"`

	// Update start time. We recommend you update LogListener during off-peak hours.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	UpdateStartTime *string `json:"UpdateStartTime,omitnil,omitempty" name:"UpdateStartTime"`

	// Update end time. We recommend you update LogListener during off-peak hours.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	UpdateEndTime *string `json:"UpdateEndTime,omitnil,omitempty" name:"UpdateEndTime"`

	// Whether to enable the service log to record the logs generated by the LogListener service itself. After it is enabled, the internal logset `cls_service_logging` and the `loglistener_status`, `loglistener_alarm`, and `loglistener_business` log topics will be created, which will not incur fees.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ServiceLogging *bool `json:"ServiceLogging,omitnil,omitempty" name:"ServiceLogging"`

	// Metadata information list of a machine group
	MetaTags []*MetaTagInfo `json:"MetaTags,omitnil,omitempty" name:"MetaTags"`
}

type MachineGroupTypeInfo struct {
	// Machine group type. Valid values: `ip`: the IP addresses of collection machines are stored in `Values` of the machine group; `label`: the tags of the machines are stored in `Values` of the machine group
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Machine description list
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type MachineInfo struct {
	// Machine IP
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// Machine status. Valid values: `0`: exceptional; `1`: normal
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Machine disconnection time. If the value is empty, the machine is normal. If the machine is exceptional, a specific value will be returned.
	OfflineTime *string `json:"OfflineTime,omitnil,omitempty" name:"OfflineTime"`

	// Whether to enable automatic update for the machine. Valid values: `0`: no; `1`: yes
	AutoUpdate *int64 `json:"AutoUpdate,omitnil,omitempty" name:"AutoUpdate"`

	// Current machine version number
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// Machine update feature status
	UpdateStatus *int64 `json:"UpdateStatus,omitnil,omitempty" name:"UpdateStatus"`

	// Machine update result flag
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// Machine update result information
	ErrMsg *string `json:"ErrMsg,omitnil,omitempty" name:"ErrMsg"`
}

// Predefined struct for user
type MergePartitionRequestParams struct {
	// Log topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Merged `PartitionId`
	PartitionId *int64 `json:"PartitionId,omitnil,omitempty" name:"PartitionId"`
}

type MergePartitionRequest struct {
	*tchttp.BaseRequest
	
	// Log topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Merged `PartitionId`
	PartitionId *int64 `json:"PartitionId,omitnil,omitempty" name:"PartitionId"`
}

func (r *MergePartitionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MergePartitionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "PartitionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MergePartitionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MergePartitionResponseParams struct {
	// Merge result set
	Partitions []*PartitionInfo `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type MergePartitionResponse struct {
	*tchttp.BaseResponse
	Response *MergePartitionResponseParams `json:"Response"`
}

func (r *MergePartitionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MergePartitionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MetaTagInfo struct {
	// Metadata key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Metadata value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type ModifyAlarmNoticeRequestParams struct {
	// Notification group ID
	AlarmNoticeId *string `json:"AlarmNoticeId,omitnil,omitempty" name:"AlarmNoticeId"`

	// Notification group name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Notification type. Valid values:
	// <li> `Trigger`: alarm triggered
	// <li> `Recovery`: alarm cleared
	// <li> `All`: alarm triggered and alarm cleared
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Notification recipient
	NoticeReceivers []*NoticeReceiver `json:"NoticeReceivers,omitnil,omitempty" name:"NoticeReceivers"`

	// API callback information (including WeCom)
	WebCallbacks []*WebCallback `json:"WebCallbacks,omitnil,omitempty" name:"WebCallbacks"`
}

type ModifyAlarmNoticeRequest struct {
	*tchttp.BaseRequest
	
	// Notification group ID
	AlarmNoticeId *string `json:"AlarmNoticeId,omitnil,omitempty" name:"AlarmNoticeId"`

	// Notification group name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Notification type. Valid values:
	// <li> `Trigger`: alarm triggered
	// <li> `Recovery`: alarm cleared
	// <li> `All`: alarm triggered and alarm cleared
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Notification recipient
	NoticeReceivers []*NoticeReceiver `json:"NoticeReceivers,omitnil,omitempty" name:"NoticeReceivers"`

	// API callback information (including WeCom)
	WebCallbacks []*WebCallback `json:"WebCallbacks,omitnil,omitempty" name:"WebCallbacks"`
}

func (r *ModifyAlarmNoticeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmNoticeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AlarmNoticeId")
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "NoticeReceivers")
	delete(f, "WebCallbacks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAlarmNoticeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAlarmNoticeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAlarmNoticeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAlarmNoticeResponseParams `json:"Response"`
}

func (r *ModifyAlarmNoticeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmNoticeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAlarmRequestParams struct {
	// Alarm policy ID
	AlarmId *string `json:"AlarmId,omitnil,omitempty" name:"AlarmId"`

	// Alarm policy name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Monitoring task running time point
	MonitorTime *MonitorTime `json:"MonitorTime,omitnil,omitempty" name:"MonitorTime"`

	// Trigger condition
	Condition *string `json:"Condition,omitnil,omitempty" name:"Condition"`

	// Alarm persistence cycle. An alarm will be triggered only after the corresponding trigger condition is met for the number of times specified by `TriggerCount`. Value range: 1–10.
	TriggerCount *int64 `json:"TriggerCount,omitnil,omitempty" name:"TriggerCount"`

	// Repeated alarm interval in minutes. Value range: 0–1440.
	AlarmPeriod *int64 `json:"AlarmPeriod,omitnil,omitempty" name:"AlarmPeriod"`

	// List of associated alarm notification templates
	AlarmNoticeIds []*string `json:"AlarmNoticeIds,omitnil,omitempty" name:"AlarmNoticeIds"`

	// Monitoring object list
	AlarmTargets []*AlarmTarget `json:"AlarmTargets,omitnil,omitempty" name:"AlarmTargets"`

	// Whether to enable the alarm policy
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// Custom alarm content
	MessageTemplate *string `json:"MessageTemplate,omitnil,omitempty" name:"MessageTemplate"`

	// Custom callback
	CallBack *CallBackInfo `json:"CallBack,omitnil,omitempty" name:"CallBack"`

	// Multi-Dimensional analysis
	Analysis []*AnalysisDimensional `json:"Analysis,omitnil,omitempty" name:"Analysis"`
}

type ModifyAlarmRequest struct {
	*tchttp.BaseRequest
	
	// Alarm policy ID
	AlarmId *string `json:"AlarmId,omitnil,omitempty" name:"AlarmId"`

	// Alarm policy name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Monitoring task running time point
	MonitorTime *MonitorTime `json:"MonitorTime,omitnil,omitempty" name:"MonitorTime"`

	// Trigger condition
	Condition *string `json:"Condition,omitnil,omitempty" name:"Condition"`

	// Alarm persistence cycle. An alarm will be triggered only after the corresponding trigger condition is met for the number of times specified by `TriggerCount`. Value range: 1–10.
	TriggerCount *int64 `json:"TriggerCount,omitnil,omitempty" name:"TriggerCount"`

	// Repeated alarm interval in minutes. Value range: 0–1440.
	AlarmPeriod *int64 `json:"AlarmPeriod,omitnil,omitempty" name:"AlarmPeriod"`

	// List of associated alarm notification templates
	AlarmNoticeIds []*string `json:"AlarmNoticeIds,omitnil,omitempty" name:"AlarmNoticeIds"`

	// Monitoring object list
	AlarmTargets []*AlarmTarget `json:"AlarmTargets,omitnil,omitempty" name:"AlarmTargets"`

	// Whether to enable the alarm policy
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// Custom alarm content
	MessageTemplate *string `json:"MessageTemplate,omitnil,omitempty" name:"MessageTemplate"`

	// Custom callback
	CallBack *CallBackInfo `json:"CallBack,omitnil,omitempty" name:"CallBack"`

	// Multi-Dimensional analysis
	Analysis []*AnalysisDimensional `json:"Analysis,omitnil,omitempty" name:"Analysis"`
}

func (r *ModifyAlarmRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AlarmId")
	delete(f, "Name")
	delete(f, "MonitorTime")
	delete(f, "Condition")
	delete(f, "TriggerCount")
	delete(f, "AlarmPeriod")
	delete(f, "AlarmNoticeIds")
	delete(f, "AlarmTargets")
	delete(f, "Status")
	delete(f, "MessageTemplate")
	delete(f, "CallBack")
	delete(f, "Analysis")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAlarmRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAlarmResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAlarmResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAlarmResponseParams `json:"Response"`
}

func (r *ModifyAlarmResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyConfigRequestParams struct {
	// Collection rule configuration ID
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`

	// Collection rule configuration name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Log collection path containing the filename
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// Type of the log to be collected. Valid values: `json_log`: log in JSON format; `delimiter_log`: log in delimited format; `minimalist_log`: minimalist log; `multiline_log`: log in multi-line format; `fullregex_log`: log in full regex format. Default value: `minimalist_log`
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// Extraction rule. If `ExtractRule` is set, `LogType` must be set.
	ExtractRule *ExtractRuleInfo `json:"ExtractRule,omitnil,omitempty" name:"ExtractRule"`

	// Collection path blocklist
	ExcludePaths []*ExcludePathInfo `json:"ExcludePaths,omitnil,omitempty" name:"ExcludePaths"`

	// Log topic (TopicId) associated with collection configuration
	Output *string `json:"Output,omitnil,omitempty" name:"Output"`

	// Custom parsing string, which is a serialized JSON string
	UserDefineRule *string `json:"UserDefineRule,omitnil,omitempty" name:"UserDefineRule"`
}

type ModifyConfigRequest struct {
	*tchttp.BaseRequest
	
	// Collection rule configuration ID
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`

	// Collection rule configuration name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Log collection path containing the filename
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// Type of the log to be collected. Valid values: `json_log`: log in JSON format; `delimiter_log`: log in delimited format; `minimalist_log`: minimalist log; `multiline_log`: log in multi-line format; `fullregex_log`: log in full regex format. Default value: `minimalist_log`
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// Extraction rule. If `ExtractRule` is set, `LogType` must be set.
	ExtractRule *ExtractRuleInfo `json:"ExtractRule,omitnil,omitempty" name:"ExtractRule"`

	// Collection path blocklist
	ExcludePaths []*ExcludePathInfo `json:"ExcludePaths,omitnil,omitempty" name:"ExcludePaths"`

	// Log topic (TopicId) associated with collection configuration
	Output *string `json:"Output,omitnil,omitempty" name:"Output"`

	// Custom parsing string, which is a serialized JSON string
	UserDefineRule *string `json:"UserDefineRule,omitnil,omitempty" name:"UserDefineRule"`
}

func (r *ModifyConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigId")
	delete(f, "Name")
	delete(f, "Path")
	delete(f, "LogType")
	delete(f, "ExtractRule")
	delete(f, "ExcludePaths")
	delete(f, "Output")
	delete(f, "UserDefineRule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyConfigResponseParams `json:"Response"`
}

func (r *ModifyConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyConsumerRequestParams struct {
	// Log topic ID bound to the task
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Whether the shipping task takes effect (default: no)
	Effective *bool `json:"Effective,omitnil,omitempty" name:"Effective"`

	// Whether to ship metadata. Default value: `false`
	NeedContent *bool `json:"NeedContent,omitnil,omitempty" name:"NeedContent"`

	// Metadata to ship if `NeedContent` is `true`
	Content *ConsumerContent `json:"Content,omitnil,omitempty" name:"Content"`

	// CKafka information
	Ckafka *Ckafka `json:"Ckafka,omitnil,omitempty" name:"Ckafka"`

	// Compression mode. Valid values: `0` (no compression), `2` (snappy), `3` (LZ4).
	Compression *int64 `json:"Compression,omitnil,omitempty" name:"Compression"`
}

type ModifyConsumerRequest struct {
	*tchttp.BaseRequest
	
	// Log topic ID bound to the task
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Whether the shipping task takes effect (default: no)
	Effective *bool `json:"Effective,omitnil,omitempty" name:"Effective"`

	// Whether to ship metadata. Default value: `false`
	NeedContent *bool `json:"NeedContent,omitnil,omitempty" name:"NeedContent"`

	// Metadata to ship if `NeedContent` is `true`
	Content *ConsumerContent `json:"Content,omitnil,omitempty" name:"Content"`

	// CKafka information
	Ckafka *Ckafka `json:"Ckafka,omitnil,omitempty" name:"Ckafka"`

	// Compression mode. Valid values: `0` (no compression), `2` (snappy), `3` (LZ4).
	Compression *int64 `json:"Compression,omitnil,omitempty" name:"Compression"`
}

func (r *ModifyConsumerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConsumerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "Effective")
	delete(f, "NeedContent")
	delete(f, "Content")
	delete(f, "Ckafka")
	delete(f, "Compression")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyConsumerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyConsumerResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyConsumerResponse struct {
	*tchttp.BaseResponse
	Response *ModifyConsumerResponseParams `json:"Response"`
}

func (r *ModifyConsumerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConsumerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCosRechargeRequestParams struct {
	// COS import configuration ID.
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// ID of the log topic.
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// COS import task name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Whether the configuration is enabled. `0`: Not enabled, `1`: Enabled
	Enable *uint64 `json:"Enable,omitnil,omitempty" name:"Enable"`
}

type ModifyCosRechargeRequest struct {
	*tchttp.BaseRequest
	
	// COS import configuration ID.
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// ID of the log topic.
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// COS import task name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Whether the configuration is enabled. `0`: Not enabled, `1`: Enabled
	Enable *uint64 `json:"Enable,omitnil,omitempty" name:"Enable"`
}

func (r *ModifyCosRechargeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCosRechargeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "TopicId")
	delete(f, "Name")
	delete(f, "Enable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCosRechargeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCosRechargeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCosRechargeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCosRechargeResponseParams `json:"Response"`
}

func (r *ModifyCosRechargeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCosRechargeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDataTransformRequestParams struct {
	// Data processing task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Data processing task name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Data processing statement
	EtlContent *string `json:"EtlContent,omitnil,omitempty" name:"EtlContent"`

	// Task status. Valid values: 1 (enabled) and 2 (disabled).
	EnableFlag *int64 `json:"EnableFlag,omitnil,omitempty" name:"EnableFlag"`

	// Destination topic ID and alias of the data processing task
	DstResources []*DataTransformResouceInfo `json:"DstResources,omitnil,omitempty" name:"DstResources"`
}

type ModifyDataTransformRequest struct {
	*tchttp.BaseRequest
	
	// Data processing task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Data processing task name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Data processing statement
	EtlContent *string `json:"EtlContent,omitnil,omitempty" name:"EtlContent"`

	// Task status. Valid values: 1 (enabled) and 2 (disabled).
	EnableFlag *int64 `json:"EnableFlag,omitnil,omitempty" name:"EnableFlag"`

	// Destination topic ID and alias of the data processing task
	DstResources []*DataTransformResouceInfo `json:"DstResources,omitnil,omitempty" name:"DstResources"`
}

func (r *ModifyDataTransformRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDataTransformRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Name")
	delete(f, "EtlContent")
	delete(f, "EnableFlag")
	delete(f, "DstResources")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDataTransformRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDataTransformResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDataTransformResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDataTransformResponseParams `json:"Response"`
}

func (r *ModifyDataTransformResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDataTransformResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIndexRequestParams struct {
	// Log topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// It does not take effect by default
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// Index rule
	Rule *RuleInfo `json:"Rule,omitnil,omitempty" name:"Rule"`

	// Whether full-text indexing includes internal fields (`__FILENAME__`, `__HOSTNAME__`, and `__SOURCE__`). Default value: `false`. Recommended value: `true`.
	// * `false`: Full-text indexing does not include internal fields.
	// * `true`: Full-text indexing includes internal fields.
	IncludeInternalFields *bool `json:"IncludeInternalFields,omitnil,omitempty" name:"IncludeInternalFields"`

	// Whether full-text indexing includes metadata fields (which are prefixed with `__TAG__`). Default value: `0`. Recommended value: `1`.
	// * `0`: Full-text indexing includes only metadata fields with key-value indexing enabled.
	// * `1`: Full-text indexing includes all metadata fields.
	// * `2`: Full-text indexing does not include metadata fields.
	MetadataFlag *uint64 `json:"MetadataFlag,omitnil,omitempty" name:"MetadataFlag"`
}

type ModifyIndexRequest struct {
	*tchttp.BaseRequest
	
	// Log topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// It does not take effect by default
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// Index rule
	Rule *RuleInfo `json:"Rule,omitnil,omitempty" name:"Rule"`

	// Whether full-text indexing includes internal fields (`__FILENAME__`, `__HOSTNAME__`, and `__SOURCE__`). Default value: `false`. Recommended value: `true`.
	// * `false`: Full-text indexing does not include internal fields.
	// * `true`: Full-text indexing includes internal fields.
	IncludeInternalFields *bool `json:"IncludeInternalFields,omitnil,omitempty" name:"IncludeInternalFields"`

	// Whether full-text indexing includes metadata fields (which are prefixed with `__TAG__`). Default value: `0`. Recommended value: `1`.
	// * `0`: Full-text indexing includes only metadata fields with key-value indexing enabled.
	// * `1`: Full-text indexing includes all metadata fields.
	// * `2`: Full-text indexing does not include metadata fields.
	MetadataFlag *uint64 `json:"MetadataFlag,omitnil,omitempty" name:"MetadataFlag"`
}

func (r *ModifyIndexRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIndexRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "Status")
	delete(f, "Rule")
	delete(f, "IncludeInternalFields")
	delete(f, "MetadataFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyIndexRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIndexResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyIndexResponse struct {
	*tchttp.BaseResponse
	Response *ModifyIndexResponseParams `json:"Response"`
}

func (r *ModifyIndexResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyKafkaRechargeRequestParams struct {
	// Kafka data import configuration ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Target topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Kafka data import configuration name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Kafka type. Valid values: 0 (Tencent Cloud CKafka) and 1 (customer's Kafka)
	KafkaType *uint64 `json:"KafkaType,omitnil,omitempty" name:"KafkaType"`

	// CKafka instance ID, which is required when `KafkaType` is set to `0`
	KafkaInstance *string `json:"KafkaInstance,omitnil,omitempty" name:"KafkaInstance"`

	// Service address
	ServerAddr *string `json:"ServerAddr,omitnil,omitempty" name:"ServerAddr"`

	// Whether the service address uses an encrypted connection
	IsEncryptionAddr *bool `json:"IsEncryptionAddr,omitnil,omitempty" name:"IsEncryptionAddr"`

	// Encryption access protocol, which is required when IsEncryptionAddr` is set to `true`
	Protocol *KafkaProtocolInfo `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// List of Kafka topics to import data from. Separate multiple topics with commas (,).
	UserKafkaTopics *string `json:"UserKafkaTopics,omitnil,omitempty" name:"UserKafkaTopics"`

	// Kafka consumer group name
	ConsumerGroupName *string `json:"ConsumerGroupName,omitnil,omitempty" name:"ConsumerGroupName"`

	// Log import rule
	LogRechargeRule *LogRechargeRuleInfo `json:"LogRechargeRule,omitnil,omitempty" name:"LogRechargeRule"`

	// Import control. Valid values: 1 (suspend) and 2 (resume).
	StatusControl *uint64 `json:"StatusControl,omitnil,omitempty" name:"StatusControl"`
}

type ModifyKafkaRechargeRequest struct {
	*tchttp.BaseRequest
	
	// Kafka data import configuration ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Target topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Kafka data import configuration name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Kafka type. Valid values: 0 (Tencent Cloud CKafka) and 1 (customer's Kafka)
	KafkaType *uint64 `json:"KafkaType,omitnil,omitempty" name:"KafkaType"`

	// CKafka instance ID, which is required when `KafkaType` is set to `0`
	KafkaInstance *string `json:"KafkaInstance,omitnil,omitempty" name:"KafkaInstance"`

	// Service address
	ServerAddr *string `json:"ServerAddr,omitnil,omitempty" name:"ServerAddr"`

	// Whether the service address uses an encrypted connection
	IsEncryptionAddr *bool `json:"IsEncryptionAddr,omitnil,omitempty" name:"IsEncryptionAddr"`

	// Encryption access protocol, which is required when IsEncryptionAddr` is set to `true`
	Protocol *KafkaProtocolInfo `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// List of Kafka topics to import data from. Separate multiple topics with commas (,).
	UserKafkaTopics *string `json:"UserKafkaTopics,omitnil,omitempty" name:"UserKafkaTopics"`

	// Kafka consumer group name
	ConsumerGroupName *string `json:"ConsumerGroupName,omitnil,omitempty" name:"ConsumerGroupName"`

	// Log import rule
	LogRechargeRule *LogRechargeRuleInfo `json:"LogRechargeRule,omitnil,omitempty" name:"LogRechargeRule"`

	// Import control. Valid values: 1 (suspend) and 2 (resume).
	StatusControl *uint64 `json:"StatusControl,omitnil,omitempty" name:"StatusControl"`
}

func (r *ModifyKafkaRechargeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyKafkaRechargeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "TopicId")
	delete(f, "Name")
	delete(f, "KafkaType")
	delete(f, "KafkaInstance")
	delete(f, "ServerAddr")
	delete(f, "IsEncryptionAddr")
	delete(f, "Protocol")
	delete(f, "UserKafkaTopics")
	delete(f, "ConsumerGroupName")
	delete(f, "LogRechargeRule")
	delete(f, "StatusControl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyKafkaRechargeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyKafkaRechargeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyKafkaRechargeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyKafkaRechargeResponseParams `json:"Response"`
}

func (r *ModifyKafkaRechargeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyKafkaRechargeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLogsetRequestParams struct {
	// Logset ID
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// Logset name
	LogsetName *string `json:"LogsetName,omitnil,omitempty" name:"LogsetName"`

	// Tag key-value pair bound to logset. Up to 10 tag key-value pairs are supported, and a resource can be bound to only one tag key at any time.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type ModifyLogsetRequest struct {
	*tchttp.BaseRequest
	
	// Logset ID
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// Logset name
	LogsetName *string `json:"LogsetName,omitnil,omitempty" name:"LogsetName"`

	// Tag key-value pair bound to logset. Up to 10 tag key-value pairs are supported, and a resource can be bound to only one tag key at any time.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *ModifyLogsetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLogsetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogsetId")
	delete(f, "LogsetName")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLogsetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLogsetResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLogsetResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLogsetResponseParams `json:"Response"`
}

func (r *ModifyLogsetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLogsetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMachineGroupRequestParams struct {
	// Machine group ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Machine group name
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Machine group type
	MachineGroupType *MachineGroupTypeInfo `json:"MachineGroupType,omitnil,omitempty" name:"MachineGroupType"`

	// Tag list
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Whether to enable automatic update for the machine group
	AutoUpdate *bool `json:"AutoUpdate,omitnil,omitempty" name:"AutoUpdate"`

	// Update start time. We recommend you update LogListener during off-peak hours.
	UpdateStartTime *string `json:"UpdateStartTime,omitnil,omitempty" name:"UpdateStartTime"`

	// Update end time. We recommend you update LogListener during off-peak hours.
	UpdateEndTime *string `json:"UpdateEndTime,omitnil,omitempty" name:"UpdateEndTime"`

	// Whether to enable the service log to record the logs generated by the LogListener service itself. After it is enabled, the internal logset `cls_service_logging` and the `loglistener_status`, `loglistener_alarm`, and `loglistener_business` log topics will be created, which will not incur fees.
	ServiceLogging *bool `json:"ServiceLogging,omitnil,omitempty" name:"ServiceLogging"`

	// Metadata information list of a machine group
	MetaTags []*MetaTagInfo `json:"MetaTags,omitnil,omitempty" name:"MetaTags"`
}

type ModifyMachineGroupRequest struct {
	*tchttp.BaseRequest
	
	// Machine group ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Machine group name
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Machine group type
	MachineGroupType *MachineGroupTypeInfo `json:"MachineGroupType,omitnil,omitempty" name:"MachineGroupType"`

	// Tag list
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Whether to enable automatic update for the machine group
	AutoUpdate *bool `json:"AutoUpdate,omitnil,omitempty" name:"AutoUpdate"`

	// Update start time. We recommend you update LogListener during off-peak hours.
	UpdateStartTime *string `json:"UpdateStartTime,omitnil,omitempty" name:"UpdateStartTime"`

	// Update end time. We recommend you update LogListener during off-peak hours.
	UpdateEndTime *string `json:"UpdateEndTime,omitnil,omitempty" name:"UpdateEndTime"`

	// Whether to enable the service log to record the logs generated by the LogListener service itself. After it is enabled, the internal logset `cls_service_logging` and the `loglistener_status`, `loglistener_alarm`, and `loglistener_business` log topics will be created, which will not incur fees.
	ServiceLogging *bool `json:"ServiceLogging,omitnil,omitempty" name:"ServiceLogging"`

	// Metadata information list of a machine group
	MetaTags []*MetaTagInfo `json:"MetaTags,omitnil,omitempty" name:"MetaTags"`
}

func (r *ModifyMachineGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMachineGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "GroupName")
	delete(f, "MachineGroupType")
	delete(f, "Tags")
	delete(f, "AutoUpdate")
	delete(f, "UpdateStartTime")
	delete(f, "UpdateEndTime")
	delete(f, "ServiceLogging")
	delete(f, "MetaTags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMachineGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMachineGroupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyMachineGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMachineGroupResponseParams `json:"Response"`
}

func (r *ModifyMachineGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMachineGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyShipperRequestParams struct {
	// Shipping rule ID
	ShipperId *string `json:"ShipperId,omitnil,omitempty" name:"ShipperId"`

	// New destination bucket in shipping rule
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// New destination directory prefix in shipping rule
	Prefix *string `json:"Prefix,omitnil,omitempty" name:"Prefix"`

	// Shipping rule status
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// Shipping rule name
	ShipperName *string `json:"ShipperName,omitnil,omitempty" name:"ShipperName"`

	// Shipping time interval in seconds. Default value: 300. Value range: 300–900
	Interval *uint64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// Maximum size of a file to be shipped, in MB. Default value: 256. Value range: 5-256
	MaxSize *uint64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`

	// Filter rules for shipped logs. Only logs matching the rules can be shipped. All rules are in the AND relationship, and up to five rules can be added. If the array is empty, no filtering will be performed, and all logs will be shipped.
	FilterRules []*FilterRuleInfo `json:"FilterRules,omitnil,omitempty" name:"FilterRules"`

	// Partition rule of shipped log, which can be represented in `strftime` time format
	Partition *string `json:"Partition,omitnil,omitempty" name:"Partition"`

	// Compression configuration of shipped log
	Compress *CompressInfo `json:"Compress,omitnil,omitempty" name:"Compress"`

	// Format configuration of shipped log content
	Content *ContentInfo `json:"Content,omitnil,omitempty" name:"Content"`

	// Naming a shipping file. Valid values: `0` (by random number), `1` (by shipping time). Default value: `0`.
	FilenameMode *uint64 `json:"FilenameMode,omitnil,omitempty" name:"FilenameMode"`
}

type ModifyShipperRequest struct {
	*tchttp.BaseRequest
	
	// Shipping rule ID
	ShipperId *string `json:"ShipperId,omitnil,omitempty" name:"ShipperId"`

	// New destination bucket in shipping rule
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// New destination directory prefix in shipping rule
	Prefix *string `json:"Prefix,omitnil,omitempty" name:"Prefix"`

	// Shipping rule status
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// Shipping rule name
	ShipperName *string `json:"ShipperName,omitnil,omitempty" name:"ShipperName"`

	// Shipping time interval in seconds. Default value: 300. Value range: 300–900
	Interval *uint64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// Maximum size of a file to be shipped, in MB. Default value: 256. Value range: 5-256
	MaxSize *uint64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`

	// Filter rules for shipped logs. Only logs matching the rules can be shipped. All rules are in the AND relationship, and up to five rules can be added. If the array is empty, no filtering will be performed, and all logs will be shipped.
	FilterRules []*FilterRuleInfo `json:"FilterRules,omitnil,omitempty" name:"FilterRules"`

	// Partition rule of shipped log, which can be represented in `strftime` time format
	Partition *string `json:"Partition,omitnil,omitempty" name:"Partition"`

	// Compression configuration of shipped log
	Compress *CompressInfo `json:"Compress,omitnil,omitempty" name:"Compress"`

	// Format configuration of shipped log content
	Content *ContentInfo `json:"Content,omitnil,omitempty" name:"Content"`

	// Naming a shipping file. Valid values: `0` (by random number), `1` (by shipping time). Default value: `0`.
	FilenameMode *uint64 `json:"FilenameMode,omitnil,omitempty" name:"FilenameMode"`
}

func (r *ModifyShipperRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyShipperRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ShipperId")
	delete(f, "Bucket")
	delete(f, "Prefix")
	delete(f, "Status")
	delete(f, "ShipperName")
	delete(f, "Interval")
	delete(f, "MaxSize")
	delete(f, "FilterRules")
	delete(f, "Partition")
	delete(f, "Compress")
	delete(f, "Content")
	delete(f, "FilenameMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyShipperRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyShipperResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyShipperResponse struct {
	*tchttp.BaseResponse
	Response *ModifyShipperResponseParams `json:"Response"`
}

func (r *ModifyShipperResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyShipperResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTopicRequestParams struct {
	// Log topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Log topic name
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Tag description list. This parameter is used to bind a tag to a log topic. Up to 10 tag key-value pairs are supported, and they must be unique.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Whether to start collection for this log topic
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// Whether to enable automatic split
	AutoSplit *bool `json:"AutoSplit,omitnil,omitempty" name:"AutoSplit"`

	// Maximum number of partitions to split into for this topic if automatic split is enabled
	MaxSplitPartitions *int64 `json:"MaxSplitPartitions,omitnil,omitempty" name:"MaxSplitPartitions"`

	// Lifecycle in days. Value range: 1–3600 (STANDARD storage); 7–3600 (IA storage). `3640` indicates permanent retention.
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Log topic description
	Describes *string `json:"Describes,omitnil,omitempty" name:"Describes"`

	// `0`: Disable log transitioning.
	// A value other than `0`: The number of STANDARD storage days after log transitioning is enabled (valid only if `StorageType` is `hot`). Note: `HotPeriod` should be greater than or equal to `7` and less than `Period`.
	HotPeriod *uint64 `json:"HotPeriod,omitnil,omitempty" name:"HotPeriod"`

	// Whether to enable web tracking. Valid values: `false` (disable); `true` (enable)
	IsWebTracking *bool `json:"IsWebTracking,omitnil,omitempty" name:"IsWebTracking"`
}

type ModifyTopicRequest struct {
	*tchttp.BaseRequest
	
	// Log topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Log topic name
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Tag description list. This parameter is used to bind a tag to a log topic. Up to 10 tag key-value pairs are supported, and they must be unique.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Whether to start collection for this log topic
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// Whether to enable automatic split
	AutoSplit *bool `json:"AutoSplit,omitnil,omitempty" name:"AutoSplit"`

	// Maximum number of partitions to split into for this topic if automatic split is enabled
	MaxSplitPartitions *int64 `json:"MaxSplitPartitions,omitnil,omitempty" name:"MaxSplitPartitions"`

	// Lifecycle in days. Value range: 1–3600 (STANDARD storage); 7–3600 (IA storage). `3640` indicates permanent retention.
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Log topic description
	Describes *string `json:"Describes,omitnil,omitempty" name:"Describes"`

	// `0`: Disable log transitioning.
	// A value other than `0`: The number of STANDARD storage days after log transitioning is enabled (valid only if `StorageType` is `hot`). Note: `HotPeriod` should be greater than or equal to `7` and less than `Period`.
	HotPeriod *uint64 `json:"HotPeriod,omitnil,omitempty" name:"HotPeriod"`

	// Whether to enable web tracking. Valid values: `false` (disable); `true` (enable)
	IsWebTracking *bool `json:"IsWebTracking,omitnil,omitempty" name:"IsWebTracking"`
}

func (r *ModifyTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "TopicName")
	delete(f, "Tags")
	delete(f, "Status")
	delete(f, "AutoSplit")
	delete(f, "MaxSplitPartitions")
	delete(f, "Period")
	delete(f, "Describes")
	delete(f, "HotPeriod")
	delete(f, "IsWebTracking")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTopicResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTopicResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTopicResponseParams `json:"Response"`
}

func (r *ModifyTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MonitorTime struct {
	// Valid values:
	// <br><li> `Period`: periodic execution
	// <br><li> `Fixed`: scheduled execution
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Execution interval or scheduled time point in minutes. Value range: 1–1440.
	Time *int64 `json:"Time,omitnil,omitempty" name:"Time"`
}

type MultiTopicSearchInformation struct {
	// ID of the log topic to be searched for
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// You can pass through the `Context` value (validity: 1 hour) returned by the last API to continue to get logs, which can get up to 10,000 raw logs.
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`
}

type NoticeReceiver struct {
	// Recipient type. Valid values:
	// <br><li> `Uin`: user ID
	// <br><li> `Group`: user group ID
	// Currently, other recipient types are not supported.
	ReceiverType *string `json:"ReceiverType,omitnil,omitempty" name:"ReceiverType"`

	// Recipient
	ReceiverIds []*int64 `json:"ReceiverIds,omitnil,omitempty" name:"ReceiverIds"`

	// Notification method
	// <br><li> `Email`: email
	// <br><li> `Sms`: SMS
	// <br><li> `WeChat`: WeChat
	// <br><li> `Phone`: phone
	ReceiverChannels []*string `json:"ReceiverChannels,omitnil,omitempty" name:"ReceiverChannels"`

	// Start time for allowed message receipt
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time for allowed message receipt
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Index
	Index *int64 `json:"Index,omitnil,omitempty" name:"Index"`
}

// Predefined struct for user
type OpenKafkaConsumerRequestParams struct {
	// `TopicId` created by the CLS console
	FromTopicId *string `json:"FromTopicId,omitnil,omitempty" name:"FromTopicId"`

	// Compression mode. Valid values: `0` (no compression); `2` (snappy); `3` (LZ4)
	Compression *int64 `json:"Compression,omitnil,omitempty" name:"Compression"`

	// Kafka consumer data format
	ConsumerContent *KafkaConsumerContent `json:"ConsumerContent,omitnil,omitempty" name:"ConsumerContent"`
}

type OpenKafkaConsumerRequest struct {
	*tchttp.BaseRequest
	
	// `TopicId` created by the CLS console
	FromTopicId *string `json:"FromTopicId,omitnil,omitempty" name:"FromTopicId"`

	// Compression mode. Valid values: `0` (no compression); `2` (snappy); `3` (LZ4)
	Compression *int64 `json:"Compression,omitnil,omitempty" name:"Compression"`

	// Kafka consumer data format
	ConsumerContent *KafkaConsumerContent `json:"ConsumerContent,omitnil,omitempty" name:"ConsumerContent"`
}

func (r *OpenKafkaConsumerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenKafkaConsumerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FromTopicId")
	delete(f, "Compression")
	delete(f, "ConsumerContent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenKafkaConsumerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenKafkaConsumerResponseParams struct {
	// `TopicId` to be consumed
	TopicID *string `json:"TopicID,omitnil,omitempty" name:"TopicID"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OpenKafkaConsumerResponse struct {
	*tchttp.BaseResponse
	Response *OpenKafkaConsumerResponseParams `json:"Response"`
}

func (r *OpenKafkaConsumerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenKafkaConsumerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ParquetInfo struct {
	// `ParquetKeyInfo` array
	ParquetKeyInfo []*ParquetKeyInfo `json:"ParquetKeyInfo,omitnil,omitempty" name:"ParquetKeyInfo"`
}

type ParquetKeyInfo struct {
	// Key name
	KeyName *string `json:"KeyName,omitnil,omitempty" name:"KeyName"`

	// Supported data types: string, boolean, int32, int64, float, and double
	KeyType *string `json:"KeyType,omitnil,omitempty" name:"KeyType"`

	// Assignment information returned upon resolution failure
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	KeyNonExistingField *string `json:"KeyNonExistingField,omitnil,omitempty" name:"KeyNonExistingField"`
}

type PartitionInfo struct {
	// Partition ID
	PartitionId *int64 `json:"PartitionId,omitnil,omitempty" name:"PartitionId"`

	// Partition status. Valid values: `readwrite`, `readonly`
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Partition hash start key
	InclusiveBeginKey *string `json:"InclusiveBeginKey,omitnil,omitempty" name:"InclusiveBeginKey"`

	// Partition hash end key
	ExclusiveEndKey *string `json:"ExclusiveEndKey,omitnil,omitempty" name:"ExclusiveEndKey"`

	// Partition creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Last modified of read-only partition
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	LastWriteTime *string `json:"LastWriteTime,omitnil,omitempty" name:"LastWriteTime"`
}

// Predefined struct for user
type PreviewKafkaRechargeRequestParams struct {
	// Preview type. Valid values: 1 (source data preview) and 2 (result preview).
	PreviewType *uint64 `json:"PreviewType,omitnil,omitempty" name:"PreviewType"`

	// Kafka type. Valid values: 0 (Tencent Cloud CKafka) and 1 (customer's Kafka)
	KafkaType *uint64 `json:"KafkaType,omitnil,omitempty" name:"KafkaType"`

	// List of Kafka topics to import data from. Separate multiple topics with commas (,).
	UserKafkaTopics *string `json:"UserKafkaTopics,omitnil,omitempty" name:"UserKafkaTopics"`

	// Position for data import. Valid values: -2 (earliest, default) and -1 (latest).
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// CKafka instance ID, which is required when `KafkaType` is set to `0`
	KafkaInstance *string `json:"KafkaInstance,omitnil,omitempty" name:"KafkaInstance"`

	// Service address
	ServerAddr *string `json:"ServerAddr,omitnil,omitempty" name:"ServerAddr"`

	// Whether the service address uses an encrypted connection
	IsEncryptionAddr *bool `json:"IsEncryptionAddr,omitnil,omitempty" name:"IsEncryptionAddr"`

	// Encryption access protocol, which is required when `IsEncryptionAddr` is set to `true`
	Protocol *KafkaProtocolInfo `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Kafka consumer group name
	ConsumerGroupName *string `json:"ConsumerGroupName,omitnil,omitempty" name:"ConsumerGroupName"`

	// Log import rule
	LogRechargeRule *LogRechargeRuleInfo `json:"LogRechargeRule,omitnil,omitempty" name:"LogRechargeRule"`
}

type PreviewKafkaRechargeRequest struct {
	*tchttp.BaseRequest
	
	// Preview type. Valid values: 1 (source data preview) and 2 (result preview).
	PreviewType *uint64 `json:"PreviewType,omitnil,omitempty" name:"PreviewType"`

	// Kafka type. Valid values: 0 (Tencent Cloud CKafka) and 1 (customer's Kafka)
	KafkaType *uint64 `json:"KafkaType,omitnil,omitempty" name:"KafkaType"`

	// List of Kafka topics to import data from. Separate multiple topics with commas (,).
	UserKafkaTopics *string `json:"UserKafkaTopics,omitnil,omitempty" name:"UserKafkaTopics"`

	// Position for data import. Valid values: -2 (earliest, default) and -1 (latest).
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// CKafka instance ID, which is required when `KafkaType` is set to `0`
	KafkaInstance *string `json:"KafkaInstance,omitnil,omitempty" name:"KafkaInstance"`

	// Service address
	ServerAddr *string `json:"ServerAddr,omitnil,omitempty" name:"ServerAddr"`

	// Whether the service address uses an encrypted connection
	IsEncryptionAddr *bool `json:"IsEncryptionAddr,omitnil,omitempty" name:"IsEncryptionAddr"`

	// Encryption access protocol, which is required when `IsEncryptionAddr` is set to `true`
	Protocol *KafkaProtocolInfo `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Kafka consumer group name
	ConsumerGroupName *string `json:"ConsumerGroupName,omitnil,omitempty" name:"ConsumerGroupName"`

	// Log import rule
	LogRechargeRule *LogRechargeRuleInfo `json:"LogRechargeRule,omitnil,omitempty" name:"LogRechargeRule"`
}

func (r *PreviewKafkaRechargeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PreviewKafkaRechargeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PreviewType")
	delete(f, "KafkaType")
	delete(f, "UserKafkaTopics")
	delete(f, "Offset")
	delete(f, "KafkaInstance")
	delete(f, "ServerAddr")
	delete(f, "IsEncryptionAddr")
	delete(f, "Protocol")
	delete(f, "ConsumerGroupName")
	delete(f, "LogRechargeRule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PreviewKafkaRechargeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PreviewKafkaRechargeResponseParams struct {
	// Log sample, which is returned when `PreviewType` is set to `2`
	LogSample *string `json:"LogSample,omitnil,omitempty" name:"LogSample"`

	// Log preview result
	// Note: This field may return null, indicating that no valid values can be obtained.
	LogData *string `json:"LogData,omitnil,omitempty" name:"LogData"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PreviewKafkaRechargeResponse struct {
	*tchttp.BaseResponse
	Response *PreviewKafkaRechargeResponseParams `json:"Response"`
}

func (r *PreviewKafkaRechargeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PreviewKafkaRechargeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PreviewLogStatistic struct {
	// Log content
	LogContent *string `json:"LogContent,omitnil,omitempty" name:"LogContent"`

	// Line number
	LineNum *int64 `json:"LineNum,omitnil,omitempty" name:"LineNum"`

	// Target log topic
	DstTopicId *string `json:"DstTopicId,omitnil,omitempty" name:"DstTopicId"`

	// Error code. An empty string "" indicates no error.
	FailReason *string `json:"FailReason,omitnil,omitempty" name:"FailReason"`

	// Log timestamp
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// Target topic name
	// Note: This field may return null, indicating that no valid values can be obtained.
	DstTopicName *string `json:"DstTopicName,omitnil,omitempty" name:"DstTopicName"`
}

// Predefined struct for user
type RetryShipperTaskRequestParams struct {
	// Shipping rule ID
	ShipperId *string `json:"ShipperId,omitnil,omitempty" name:"ShipperId"`

	// Shipping task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type RetryShipperTaskRequest struct {
	*tchttp.BaseRequest
	
	// Shipping rule ID
	ShipperId *string `json:"ShipperId,omitnil,omitempty" name:"ShipperId"`

	// Shipping task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *RetryShipperTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetryShipperTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ShipperId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RetryShipperTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RetryShipperTaskResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RetryShipperTaskResponse struct {
	*tchttp.BaseResponse
	Response *RetryShipperTaskResponseParams `json:"Response"`
}

func (r *RetryShipperTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetryShipperTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleInfo struct {
	// Full-text index configuration. If the configuration is left empty, full-text indexing is not enabled.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FullText *FullTextInfo `json:"FullText,omitnil,omitempty" name:"FullText"`

	// Key-value index configuration. If the configuration is left empty, key-value indexing is not enabled.
	// Note: This field may return null, indicating that no valid values can be obtained.
	KeyValue *RuleKeyValueInfo `json:"KeyValue,omitnil,omitempty" name:"KeyValue"`

	// Metadata field index configuration. If the configuration is left empty, metadata field indexing is not enabled.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tag *RuleTagInfo `json:"Tag,omitnil,omitempty" name:"Tag"`

	// Dynamic index configuration. If the configuration is empty, dynamic indexing is not enabled.
	// 
	// Note: This feature is currently in a beta test. To use it, please contact technical support.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DynamicIndex *DynamicIndex `json:"DynamicIndex,omitnil,omitempty" name:"DynamicIndex"`
}

type RuleKeyValueInfo struct {
	// Case sensitivity
	CaseSensitive *bool `json:"CaseSensitive,omitnil,omitempty" name:"CaseSensitive"`

	// Key-value pair information of the index to be created
	KeyValues []*KeyValueInfo `json:"KeyValues,omitnil,omitempty" name:"KeyValues"`
}

type RuleTagInfo struct {
	// Case sensitivity
	CaseSensitive *bool `json:"CaseSensitive,omitnil,omitempty" name:"CaseSensitive"`

	// Field information in the metafield index configuration
	KeyValues []*KeyValueInfo `json:"KeyValues,omitnil,omitempty" name:"KeyValues"`
}

// Predefined struct for user
type SearchLogRequestParams struct {
	// Start time of the log to be searched, which is a Unix timestamp in milliseconds
	From *int64 `json:"From,omitnil,omitempty" name:"From"`

	// End time of the log to be searched, which is a Unix timestamp in milliseconds
	To *int64 `json:"To,omitnil,omitempty" name:"To"`

	// Search and analysis statement. Maximum length: 12 KB
	// A statement is in the format of <a href="https://intl.cloud.tencent.com/document/product/614/47044?from_cn_redirect=1" target="_blank">[search criteria]</a> | <a href="https://intl.cloud.tencent.com/document/product/614/44061?from_cn_redirect=1" target="_blank">[SQL statement]</a>. You can omit the pipe symbol <code> | </code> and SQL statement when log analysis is not required.
	// Queries all logs using * or an empty string
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// - The ID of the log topic to be searched for. Only one log topic can be specified.
	// - To search for multiple log topics at a time, use the `Topics` parameter.
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// The number of raw logs returned by a single query. Maximum value: 1000. You need to use `Context` to continue to get logs.
	// Notes:
	// * This parameter is valid only when the query statement (`Query`) does not contain an SQL statement.
	// * To limit the number of analysis results, see <a href="https://intl.cloud.tencent.com/document/product/614/58977?from_cn_redirect=1" target="_blank">SQL LIMIT Syntax</a>.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// You can pass through the `Context` value (validity: an hour) returned by the API last time to continue to get logs (up to 10,000 raw logs).
	// Notes:
	// * Do not modify any other parameters while passing through the `Context` parameter.
	// * This parameter is valid only when the query statement (`Query`) does not contain an SQL statement.
	// * To continue to get analysis results, see <a href="https://intl.cloud.tencent.com/document/product/614/58977?from_cn_redirect=1" target="_blank">SQL LIMIT Syntax</a>.
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// Time order of the logs returned. Valid values: `asc` (ascending); `desc`: (descending). Default value: `desc`
	// Notes:
	// * This parameter is valid only when the query statement (`Query`) does not contain an SQL statement.
	// * To sort the analysis results, see <a href="https://intl.cloud.tencent.com/document/product/614/58978?from_cn_redirect=1" target="_blank">SQL ORDER BY Syntax</a>.
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// If the value is `true`, the new response method will be used, and the output parameters `AnalysisRecords` and `Columns` will be valid.
	// If the value is `false`, the old response method will be used, and the output parameters `AnalysisResults` and `ColNames` will be valid.
	// The two response methods differ slightly in terms of encoding format. You are advised to use the new method (`true`).
	UseNewAnalysis *bool `json:"UseNewAnalysis,omitnil,omitempty" name:"UseNewAnalysis"`

	// Indicates whether to sample raw logs before statistical analysis (`Query` includes SQL statements).
	// `0`: Auto-sample.
	// `0–1`: Sample by the specified sample rate, such as `0.02`.
	// `1`: Precise analysis without sampling.
	// Default value: `1`
	SamplingRate *float64 `json:"SamplingRate,omitnil,omitempty" name:"SamplingRate"`

	// Search syntax
	// `0` (default): Lucene; `1`: CQL.
	// For more information, see <a href="https://intl.cloud.tencent.com/document/product/614/47044?from_cn_redirect=1#RetrievesConditionalRules" target="_blank">Syntax Rules</a>
	SyntaxRule *uint64 `json:"SyntaxRule,omitnil,omitempty" name:"SyntaxRule"`

	// - The IDs of the log topics (up to 20) to be searched for.
	// - To search for a single log topic, use the `TopicId` parameter.
	// - You cannot use both `TopicId` and `Topics`.
	Topics []*MultiTopicSearchInformation `json:"Topics,omitnil,omitempty" name:"Topics"`
}

type SearchLogRequest struct {
	*tchttp.BaseRequest
	
	// Start time of the log to be searched, which is a Unix timestamp in milliseconds
	From *int64 `json:"From,omitnil,omitempty" name:"From"`

	// End time of the log to be searched, which is a Unix timestamp in milliseconds
	To *int64 `json:"To,omitnil,omitempty" name:"To"`

	// Search and analysis statement. Maximum length: 12 KB
	// A statement is in the format of <a href="https://intl.cloud.tencent.com/document/product/614/47044?from_cn_redirect=1" target="_blank">[search criteria]</a> | <a href="https://intl.cloud.tencent.com/document/product/614/44061?from_cn_redirect=1" target="_blank">[SQL statement]</a>. You can omit the pipe symbol <code> | </code> and SQL statement when log analysis is not required.
	// Queries all logs using * or an empty string
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// - The ID of the log topic to be searched for. Only one log topic can be specified.
	// - To search for multiple log topics at a time, use the `Topics` parameter.
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// The number of raw logs returned by a single query. Maximum value: 1000. You need to use `Context` to continue to get logs.
	// Notes:
	// * This parameter is valid only when the query statement (`Query`) does not contain an SQL statement.
	// * To limit the number of analysis results, see <a href="https://intl.cloud.tencent.com/document/product/614/58977?from_cn_redirect=1" target="_blank">SQL LIMIT Syntax</a>.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// You can pass through the `Context` value (validity: an hour) returned by the API last time to continue to get logs (up to 10,000 raw logs).
	// Notes:
	// * Do not modify any other parameters while passing through the `Context` parameter.
	// * This parameter is valid only when the query statement (`Query`) does not contain an SQL statement.
	// * To continue to get analysis results, see <a href="https://intl.cloud.tencent.com/document/product/614/58977?from_cn_redirect=1" target="_blank">SQL LIMIT Syntax</a>.
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// Time order of the logs returned. Valid values: `asc` (ascending); `desc`: (descending). Default value: `desc`
	// Notes:
	// * This parameter is valid only when the query statement (`Query`) does not contain an SQL statement.
	// * To sort the analysis results, see <a href="https://intl.cloud.tencent.com/document/product/614/58978?from_cn_redirect=1" target="_blank">SQL ORDER BY Syntax</a>.
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// If the value is `true`, the new response method will be used, and the output parameters `AnalysisRecords` and `Columns` will be valid.
	// If the value is `false`, the old response method will be used, and the output parameters `AnalysisResults` and `ColNames` will be valid.
	// The two response methods differ slightly in terms of encoding format. You are advised to use the new method (`true`).
	UseNewAnalysis *bool `json:"UseNewAnalysis,omitnil,omitempty" name:"UseNewAnalysis"`

	// Indicates whether to sample raw logs before statistical analysis (`Query` includes SQL statements).
	// `0`: Auto-sample.
	// `0–1`: Sample by the specified sample rate, such as `0.02`.
	// `1`: Precise analysis without sampling.
	// Default value: `1`
	SamplingRate *float64 `json:"SamplingRate,omitnil,omitempty" name:"SamplingRate"`

	// Search syntax
	// `0` (default): Lucene; `1`: CQL.
	// For more information, see <a href="https://intl.cloud.tencent.com/document/product/614/47044?from_cn_redirect=1#RetrievesConditionalRules" target="_blank">Syntax Rules</a>
	SyntaxRule *uint64 `json:"SyntaxRule,omitnil,omitempty" name:"SyntaxRule"`

	// - The IDs of the log topics (up to 20) to be searched for.
	// - To search for a single log topic, use the `TopicId` parameter.
	// - You cannot use both `TopicId` and `Topics`.
	Topics []*MultiTopicSearchInformation `json:"Topics,omitnil,omitempty" name:"Topics"`
}

func (r *SearchLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "From")
	delete(f, "To")
	delete(f, "Query")
	delete(f, "TopicId")
	delete(f, "Limit")
	delete(f, "Context")
	delete(f, "Sort")
	delete(f, "UseNewAnalysis")
	delete(f, "SamplingRate")
	delete(f, "SyntaxRule")
	delete(f, "Topics")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchLogResponseParams struct {
	// You can pass through the `Context` value (validity: 1 hour) returned by this API to continue to get more logs.
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// Whether to return all raw log query results. If not, you can use `Context` to continue to get logs.
	// Note: This parameter is valid only when the query statement (`Query`) does not contain an SQL statement.
	ListOver *bool `json:"ListOver,omitnil,omitempty" name:"ListOver"`

	// Whether the returned data is the analysis (SQL) result
	Analysis *bool `json:"Analysis,omitnil,omitempty" name:"Analysis"`

	// Raw logs that meet the search conditions
	// Note: This field may return `null`, indicating that no valid value was found.
	Results []*LogInfo `json:"Results,omitnil,omitempty" name:"Results"`

	// Column names of log analysis
	// This parameter is valid only when `UseNewAnalysis` is `false`.
	// Note: This field may return `null`, indicating that no valid value was found.
	ColNames []*string `json:"ColNames,omitnil,omitempty" name:"ColNames"`

	// Log analysis result
	// This parameter is valid only when `UseNewAnalysis` is `false`.
	// Note: This field may return `null`, indicating that no valid value was found.
	AnalysisResults []*LogItems `json:"AnalysisResults,omitnil,omitempty" name:"AnalysisResults"`

	// Log analysis result
	// This parameter is valid only when `UseNewAnalysis` is `true`.
	// Note: This field may return `null`, indicating that no valid value was found.
	AnalysisRecords []*string `json:"AnalysisRecords,omitnil,omitempty" name:"AnalysisRecords"`

	// Column attributes of log analysis
	// This parameter is valid only when `UseNewAnalysis` is `true`.
	// Note: This field may return `null`, indicating that no valid value was found.
	Columns []*Column `json:"Columns,omitnil,omitempty" name:"Columns"`

	// Sample rate used in this statistical analysis
	// Note: This field may return null, indicating that no valid values can be obtained.
	SamplingRate *float64 `json:"SamplingRate,omitnil,omitempty" name:"SamplingRate"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SearchLogResponse struct {
	*tchttp.BaseResponse
	Response *SearchLogResponseParams `json:"Response"`
}

func (r *SearchLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ShipperInfo struct {
	// Shipping rule ID
	ShipperId *string `json:"ShipperId,omitnil,omitempty" name:"ShipperId"`

	// Log topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Bucket address shipped to
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// Shipping prefix directory
	Prefix *string `json:"Prefix,omitnil,omitempty" name:"Prefix"`

	// Shipping rule name
	ShipperName *string `json:"ShipperName,omitnil,omitempty" name:"ShipperName"`

	// Shipping time interval in seconds
	Interval *uint64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// Maximum size of shipped file in MB
	MaxSize *uint64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`

	// Whether it takes effect
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// Filter rule for shipped log
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	FilterRules []*FilterRuleInfo `json:"FilterRules,omitnil,omitempty" name:"FilterRules"`

	// Partition rule of shipped log, which can be represented in `strftime` time format
	Partition *string `json:"Partition,omitnil,omitempty" name:"Partition"`

	// Compression configuration of shipped log
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Compress *CompressInfo `json:"Compress,omitnil,omitempty" name:"Compress"`

	// Format configuration of shipped log content
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Content *ContentInfo `json:"Content,omitnil,omitempty" name:"Content"`

	// Creation time of shipped log
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Shipping file naming configuration. Valid values: `0` (by random number); `1` (by shipping time). Default value: `0`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FilenameMode *uint64 `json:"FilenameMode,omitnil,omitempty" name:"FilenameMode"`

	// Start time for data shipping
	// Note: This field may return null, indicating that no valid values can be obtained.
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time for data shipping
	// Note: This field may return null, indicating that no valid values can be obtained.
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Progress of historical data shipping (valid only when the selected data scope contains historical data)
	// Note: This field may return null, indicating that no valid values can be obtained.
	Progress *float64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// Remaining time required for shipping all historical data (valid only when the selected data scope contains historical data)
	// Note: This field may return null, indicating that no valid values can be obtained.
	RemainTime *int64 `json:"RemainTime,omitnil,omitempty" name:"RemainTime"`

	// Status of historical data shipping. Valid values:
	// 0: Real-time data is being shipped.
	// 1: The system is preparing for historical data shipping.
	// 2: Historical data is being shipped.
	// 3: An error occurred while shipping historical data.
	// 4: Historical data shipping ended.
	// Note: This field may return null, indicating that no valid values can be obtained.
	HistoryStatus *int64 `json:"HistoryStatus,omitnil,omitempty" name:"HistoryStatus"`
}

type ShipperTaskInfo struct {
	// Shipping task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Shipping information ID
	ShipperId *string `json:"ShipperId,omitnil,omitempty" name:"ShipperId"`

	// Log topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Start timestamp of the current batch of shipped logs in milliseconds
	RangeStart *int64 `json:"RangeStart,omitnil,omitempty" name:"RangeStart"`

	// End timestamp of the current batch of shipped logs in milliseconds
	RangeEnd *int64 `json:"RangeEnd,omitnil,omitempty" name:"RangeEnd"`

	// Start timestamp of the current shipping task in milliseconds
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End timestamp of the current shipping task in milliseconds
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Result of the current shipping task. Valid values: `success`, `running`, `failed`
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Result details
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

// Predefined struct for user
type SplitPartitionRequestParams struct {
	// Log topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// ID of the partition to be split
	PartitionId *int64 `json:"PartitionId,omitnil,omitempty" name:"PartitionId"`

	// Partition split hash key position, which is meaningful only if `Number=2` is set
	SplitKey *string `json:"SplitKey,omitnil,omitempty" name:"SplitKey"`

	// Number of partitions to split into, which is optional. Default value: 2
	Number *int64 `json:"Number,omitnil,omitempty" name:"Number"`
}

type SplitPartitionRequest struct {
	*tchttp.BaseRequest
	
	// Log topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// ID of the partition to be split
	PartitionId *int64 `json:"PartitionId,omitnil,omitempty" name:"PartitionId"`

	// Partition split hash key position, which is meaningful only if `Number=2` is set
	SplitKey *string `json:"SplitKey,omitnil,omitempty" name:"SplitKey"`

	// Number of partitions to split into, which is optional. Default value: 2
	Number *int64 `json:"Number,omitnil,omitempty" name:"Number"`
}

func (r *SplitPartitionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SplitPartitionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "PartitionId")
	delete(f, "SplitKey")
	delete(f, "Number")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SplitPartitionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SplitPartitionResponseParams struct {
	// Split result set
	Partitions []*PartitionInfo `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SplitPartitionResponse struct {
	*tchttp.BaseResponse
	Response *SplitPartitionResponseParams `json:"Response"`
}

func (r *SplitPartitionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SplitPartitionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {
	// The tag key.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// The tag value.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type TopicInfo struct {
	// Logset ID
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// Log topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Log topic name
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Number of topic partitions
	PartitionCount *int64 `json:"PartitionCount,omitnil,omitempty" name:"PartitionCount"`

	// Whether index is enabled
	Index *bool `json:"Index,omitnil,omitempty" name:"Index"`

	// Cloud product identifier. If the log topic is created by another cloud product, this field returns the name of the cloud product, such as `CDN` or `TKE`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AssumerName *string `json:"AssumerName,omitnil,omitempty" name:"AssumerName"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Whether collection is enabled in the log topic
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// Information of tags bound to log topic
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Whether automatic split is enabled for this topic
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	AutoSplit *bool `json:"AutoSplit,omitnil,omitempty" name:"AutoSplit"`

	// Maximum number of partitions to split into for this topic if automatic split is enabled
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	MaxSplitPartitions *int64 `json:"MaxSplitPartitions,omitnil,omitempty" name:"MaxSplitPartitions"`

	// Log topic storage class
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// Lifecycle in days. Value range: 1-3600 (3640 indicates permanent retention)
	// Note: This field may return `null`, indicating that no valid value was found.
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Cloud product sub-identifier. If the log topic is created by another cloud product, this field returns the name of the cloud product and its log type, such as `TKE-Audit` or `TKE-Event`. Some products only return the cloud product identifier (`AssumerName`), without this field.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubAssumerName *string `json:"SubAssumerName,omitnil,omitempty" name:"SubAssumerName"`

	// Log topic description
	// Note: This field may return null, indicating that no valid values can be obtained.
	Describes *string `json:"Describes,omitnil,omitempty" name:"Describes"`

	// The lifecycle of hot storage when log transitioning is enabled. The value of `hotPeriod` is smaller than that of `Period`.
	// The hot storage period is the value of `hotPeriod`, and the cold storage period is the value of `Period` minus the value of `hotPeriod`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	HotPeriod *uint64 `json:"HotPeriod,omitnil,omitempty" name:"HotPeriod"`
}

// Predefined struct for user
type UploadLogRequestParams struct {
	// Topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Topic partition where data will be written into by `HashKey` 
	HashKey *string `json:"HashKey,omitnil,omitempty" name:"HashKey"`

	// Compression type
	CompressType *string `json:"CompressType,omitnil,omitempty" name:"CompressType"`
}

type UploadLogRequest struct {
	*tchttp.BaseRequest
	
	// Topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Topic partition where data will be written into by `HashKey` 
	HashKey *string `json:"HashKey,omitnil,omitempty" name:"HashKey"`

	// Compression type
	CompressType *string `json:"CompressType,omitnil,omitempty" name:"CompressType"`
}

func (r *UploadLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "HashKey")
	delete(f, "CompressType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadLogResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UploadLogResponse struct {
	*tchttp.BaseResponse
	Response *UploadLogResponseParams `json:"Response"`
}

func (r *UploadLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ValueInfo struct {
	// Field type. Valid values: `long`, `text`, `double`
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Separator of fields. Each character represents a separator.
	// Only symbols, \n\t\r, and escape character \ are supported.
	// `long` and `double` fields need to be null.
	// Note: \n\t\r can be directly enclosed in double quotes as the input parameter without escaping. When debugging with API Explorer, use the JSON parameter input method to avoid repeated escaping of \n\t\r.
	Tokenizer *string `json:"Tokenizer,omitnil,omitempty" name:"Tokenizer"`

	// Whether the analysis feature is enabled for the field
	SqlFlag *bool `json:"SqlFlag,omitnil,omitempty" name:"SqlFlag"`

	// Whether Chinese characters are contained. For `long` and `double` fields, set them to `false`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ContainZH *bool `json:"ContainZH,omitnil,omitempty" name:"ContainZH"`
}

type WebCallback struct {
	// Callback address
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Callback type. Valid values:
	// <li> WeCom
	// <li> Http
	CallbackType *string `json:"CallbackType,omitnil,omitempty" name:"CallbackType"`

	// Callback method. Valid values:
	// <li> POST
	// <li> PUT
	// Default value: `POST`. This parameter is required if `CallbackType` is `Http`.
	// Note: This field may return `null`, indicating that no valid value was found.
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// Request header
	// Note: This parameter is disused. To specify request headers, see `CallBack` in <a href="https://intl.cloud.tencent.com/document/product/614/56466?from_cn_redirect=1">CreateAlarmNotice</a>.
	// Note: This field may return `null`, indicating that no valid value was found.
	Headers []*string `json:"Headers,omitnil,omitempty" name:"Headers"`

	// Request content
	// Note: This parameter is disused. To specify request content, see `CallBack` in <a href="https://intl.cloud.tencent.com/document/product/614/56466?from_cn_redirect=1">CreateAlarmNotice</a>.
	// Note: This field may return `null`, indicating that no valid value was found.
	Body *string `json:"Body,omitnil,omitempty" name:"Body"`

	// Number
	Index *int64 `json:"Index,omitnil,omitempty" name:"Index"`
}