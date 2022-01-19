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

package v20190819

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type Acl struct {

	// ACL resource type. 0: UNKNOWN, 1: ANY, 2: TOPIC, 3: GROUP, 4: CLUSTER, 5: TRANSACTIONAL_ID. Currently, only `TOPIC` is available,
	ResourceType *int64 `json:"ResourceType,omitempty" name:"ResourceType"`

	// Resource name, which is related to `resourceType`. For example, if `resourceType` is `TOPIC`, this field indicates the topic name; if `resourceType` is `GROUP`, this field indicates the group name
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// User list. The default value is `User:*`, which means that any user can access. The current user can only be one included in the user list
	// Note: this field may return null, indicating that no valid values can be obtained.
	Principal *string `json:"Principal,omitempty" name:"Principal"`

	// The default value is `*`, which means that any host can access. Currently, CKafka does not support the host as `*`, but the future product based on the open-source Kafka will directly support this
	// Note: this field may return null, indicating that no valid values can be obtained.
	Host *string `json:"Host,omitempty" name:"Host"`

	// ACL operation mode. 0: UNKNOWN, 1: ANY, 2: ALL, 3: READ, 4: WRITE, 5: CREATE, 6: DELETE, 7: ALTER, 8: DESCRIBE, 9: CLUSTER_ACTION, 10: DESCRIBE_CONFIGS, 11: ALTER_CONFIGS, 12: IDEMPOTEN_WRITE
	Operation *int64 `json:"Operation,omitempty" name:"Operation"`

	// Permission type. 0: UNKNOWN, 1: ANY, 2: DENY, 3: ALLOW
	PermissionType *int64 `json:"PermissionType,omitempty" name:"PermissionType"`
}

type AclResponse struct {

	// Number of eligible data entries
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// ACL list
	// Note: this field may return null, indicating that no valid values can be obtained.
	AclList []*Acl `json:"AclList,omitempty" name:"AclList"`
}

type AclRule struct {

	// ACL rule name.
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// Instance ID.
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Matching type. Currently, only prefix match is supported. Enumerated value list: PREFIXED
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	PatternType *string `json:"PatternType,omitempty" name:"PatternType"`

	// Prefix value for prefix match.
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	Pattern *string `json:"Pattern,omitempty" name:"Pattern"`

	// ACL resource type. Only “Topic” is supported. Enumerated value list: Topic.
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// ACL information contained in the rule.
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	AclList *string `json:"AclList,omitempty" name:"AclList"`

	// Creation time of the rule.
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	CreateTimeStamp *string `json:"CreateTimeStamp,omitempty" name:"CreateTimeStamp"`

	// A parameter used to specify whether the preset ACL rule is applied to new topics.
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	IsApplied *int64 `json:"IsApplied,omitempty" name:"IsApplied"`

	// Rule update time.
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	UpdateTimeStamp *string `json:"UpdateTimeStamp,omitempty" name:"UpdateTimeStamp"`

	// Remarks of the rule.
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// One of the corresponding topic names that is displayed.
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// The number of topics that apply this ACL rule.
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	TopicCount *int64 `json:"TopicCount,omitempty" name:"TopicCount"`

	// Name of rule type.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	PatternTypeTitle *string `json:"PatternTypeTitle,omitempty" name:"PatternTypeTitle"`
}

type AclRuleInfo struct {

	// ACL operation types. Enumerated values: `All` (all operations), `Read` (read), `Write` (write).
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// Permission types: `Deny`, `Allow`.
	PermissionType *string `json:"PermissionType,omitempty" name:"PermissionType"`

	// The default value is `*`, which means that any host can access the topic. CKafka currently does not support specifying a host value of * or an IP range.
	Host *string `json:"Host,omitempty" name:"Host"`

	// The list of users allowed to access the topic. Default value: `User:*`, which means all users. The current user must be in the user list. Add the prefix `User:` before the user name (`User:A`, for example).
	Principal *string `json:"Principal,omitempty" name:"Principal"`
}

type AppIdResponse struct {

	// Number of eligible `AppId`
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of eligible `AppId`
	// Note: this field may return null, indicating that no valid values can be obtained.
	AppIdList []*int64 `json:"AppIdList,omitempty" name:"AppIdList"`
}

type Assignment struct {

	// Assignment version information
	Version *int64 `json:"Version,omitempty" name:"Version"`

	// Topic information list
	// Note: this field may return null, indicating that no valid values can be obtained.
	Topics []*GroupInfoTopics `json:"Topics,omitempty" name:"Topics"`
}

type BatchContent struct {

	// Message body that is sent.
	Body *string `json:"Body,omitempty" name:"Body"`

	// Message sending key name.
	Key *string `json:"Key,omitempty" name:"Key"`
}

type BatchCreateAclRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// ACL resource type. Default value: `2` (topic).
	ResourceType *int64 `json:"ResourceType,omitempty" name:"ResourceType"`

	// Resource list array.
	ResourceNames []*string `json:"ResourceNames,omitempty" name:"ResourceNames"`

	// ACL rule list.
	RuleList []*AclRuleInfo `json:"RuleList,omitempty" name:"RuleList"`
}

func (r *BatchCreateAclRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchCreateAclRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ResourceType")
	delete(f, "ResourceNames")
	delete(f, "RuleList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchCreateAclRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type BatchCreateAclResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Status code.
		Result *int64 `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BatchCreateAclResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchCreateAclResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchModifyGroupOffsetsRequest struct {
	*tchttp.BaseRequest

	// Consumer group name.
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// Instance name.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Partition information.
	Partitions []*Partitions `json:"Partitions,omitempty" name:"Partitions"`

	// Name of the specified topic. Default value: names of all topics.
	TopicName []*string `json:"TopicName,omitempty" name:"TopicName"`
}

func (r *BatchModifyGroupOffsetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchModifyGroupOffsetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupName")
	delete(f, "InstanceId")
	delete(f, "Partitions")
	delete(f, "TopicName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchModifyGroupOffsetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type BatchModifyGroupOffsetsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Returned result.
		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BatchModifyGroupOffsetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchModifyGroupOffsetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchModifyTopicAttributesRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Topic attribute list
	Topic []*BatchModifyTopicInfo `json:"Topic,omitempty" name:"Topic"`
}

func (r *BatchModifyTopicAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchModifyTopicAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Topic")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchModifyTopicAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type BatchModifyTopicAttributesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Returned result.
		Result []*BatchModifyTopicResultDTO `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BatchModifyTopicAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchModifyTopicAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchModifyTopicInfo struct {

	// Topic name.
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// The number of partitions.
	PartitionNum *int64 `json:"PartitionNum,omitempty" name:"PartitionNum"`

	// Remarks.
	Note *string `json:"Note,omitempty" name:"Note"`

	// Number of replicas.
	ReplicaNum *int64 `json:"ReplicaNum,omitempty" name:"ReplicaNum"`

	// Message deletion policy. Valid values: `delete`, `compact`.
	CleanUpPolicy *string `json:"CleanUpPolicy,omitempty" name:"CleanUpPolicy"`

	// The minimum number of replicas specified by `min.insync.replicas` when the producer sets `request.required.acks` to `-1`.
	MinInsyncReplicas *int64 `json:"MinInsyncReplicas,omitempty" name:"MinInsyncReplicas"`

	// Whether to allow a non-ISR replica to be the leader.
	UncleanLeaderElectionEnable *bool `json:"UncleanLeaderElectionEnable,omitempty" name:"UncleanLeaderElectionEnable"`

	// Message retention period in topic dimension in milliseconds. Value range: 1 minute to 90 days.
	RetentionMs *int64 `json:"RetentionMs,omitempty" name:"RetentionMs"`

	// Message retention size in topic dimension. Value range: 1 MB - 1024 GB.
	RetentionBytes *int64 `json:"RetentionBytes,omitempty" name:"RetentionBytes"`

	// Segment rolling duration in milliseconds. Value range: 1-90 days.
	SegmentMs *int64 `json:"SegmentMs,omitempty" name:"SegmentMs"`

	// Message size per batch. Value range: 1 KB - 12 MB.
	MaxMessageBytes *int64 `json:"MaxMessageBytes,omitempty" name:"MaxMessageBytes"`
}

type BatchModifyTopicResultDTO struct {

	// Instance ID.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Topic name.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Status code.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ReturnCode *string `json:"ReturnCode,omitempty" name:"ReturnCode"`

	// Message status.
	Message *string `json:"Message,omitempty" name:"Message"`
}

type ClusterInfo struct {

	// Cluster ID
	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`

	// Cluster name
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// The cluster’s maximum disk capacity in GB
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	MaxDiskSize *int64 `json:"MaxDiskSize,omitempty" name:"MaxDiskSize"`

	// The cluster’s maximum bandwidth in MB/s
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	MaxBandWidth *int64 `json:"MaxBandWidth,omitempty" name:"MaxBandWidth"`

	// The cluster’s available disk capacity in GB
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	AvailableDiskSize *int64 `json:"AvailableDiskSize,omitempty" name:"AvailableDiskSize"`

	// The cluster’s available bandwidth in MB/s
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	AvailableBandWidth *int64 `json:"AvailableBandWidth,omitempty" name:"AvailableBandWidth"`

	// The AZ where the cluster resides
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// The AZ where the cluster nodes reside. If the cluster is a multi-AZ cluster, this field means multiple AZs where the cluster nodes reside.
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	ZoneIds []*int64 `json:"ZoneIds,omitempty" name:"ZoneIds"`
}

type Config struct {

	// Message retention period
	// Note: this field may return null, indicating that no valid values can be obtained.
	Retention *int64 `json:"Retention,omitempty" name:"Retention"`

	// Minimum number of sync replications
	// Note: this field may return null, indicating that no valid values can be obtained.
	MinInsyncReplicas *int64 `json:"MinInsyncReplicas,omitempty" name:"MinInsyncReplicas"`

	// Log cleanup mode. Default value: delete.
	// delete: logs will be deleted by save time; compact: logs will be compressed by key; compact, delete: logs will be compressed by key and deleted by save time.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CleanUpPolicy *string `json:"CleanUpPolicy,omitempty" name:"CleanUpPolicy"`

	// Segment rolling duration
	// Note: this field may return null, indicating that no valid values can be obtained.
	SegmentMs *int64 `json:"SegmentMs,omitempty" name:"SegmentMs"`

	// 0: false, 1: true.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UncleanLeaderElectionEnable *int64 `json:"UncleanLeaderElectionEnable,omitempty" name:"UncleanLeaderElectionEnable"`

	// Number of bytes for segment rolling
	// Note: this field may return null, indicating that no valid values can be obtained.
	SegmentBytes *int64 `json:"SegmentBytes,omitempty" name:"SegmentBytes"`

	// Maximum number of message bytes
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxMessageBytes *int64 `json:"MaxMessageBytes,omitempty" name:"MaxMessageBytes"`

	// Message retention file size.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	RetentionBytes *int64 `json:"RetentionBytes,omitempty" name:"RetentionBytes"`
}

type ConsumerGroup struct {

	// User group name
	ConsumerGroupName *string `json:"ConsumerGroupName,omitempty" name:"ConsumerGroupName"`

	// Subscribed message entity
	SubscribedInfo []*SubscribedInfo `json:"SubscribedInfo,omitempty" name:"SubscribedInfo"`
}

type ConsumerGroupResponse struct {

	// Number of eligible consumer groups
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Topic list
	// Note: this field may return null, indicating that no valid values can be obtained.
	TopicList []*ConsumerGroupTopic `json:"TopicList,omitempty" name:"TopicList"`

	// Consumer group list
	// Note: this field may return null, indicating that no valid values can be obtained.
	GroupList []*ConsumerGroup `json:"GroupList,omitempty" name:"GroupList"`

	// Total number of partitions
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalPartition *int64 `json:"TotalPartition,omitempty" name:"TotalPartition"`

	// List of monitored partitions
	// Note: this field may return null, indicating that no valid values can be obtained.
	PartitionListForMonitor []*Partition `json:"PartitionListForMonitor,omitempty" name:"PartitionListForMonitor"`

	// Total number of topics
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalTopic *int64 `json:"TotalTopic,omitempty" name:"TotalTopic"`

	// List of monitored topics
	// Note: this field may return null, indicating that no valid values can be obtained.
	TopicListForMonitor []*ConsumerGroupTopic `json:"TopicListForMonitor,omitempty" name:"TopicListForMonitor"`

	// List of monitored groups
	// Note: this field may return null, indicating that no valid values can be obtained.
	GroupListForMonitor []*Group `json:"GroupListForMonitor,omitempty" name:"GroupListForMonitor"`
}

type ConsumerGroupTopic struct {

	// Topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Topic name
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

type ConsumerRecord struct {

	// Topic name
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// Partition ID
	Partition *int64 `json:"Partition,omitempty" name:"Partition"`

	// Offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Message key
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Key *string `json:"Key,omitempty" name:"Key"`

	// Message value
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Value *string `json:"Value,omitempty" name:"Value"`

	// Message timestamp
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`
}

type CreateAclRequest struct {
	*tchttp.BaseRequest

	// Instance ID information
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// ACL resource type (`2`: TOPIC, `3`: GROUP, `4`: CLUSTER).
	ResourceType *int64 `json:"ResourceType,omitempty" name:"ResourceType"`

	// ACL operation type (`2`: ALL, `3`: READ, `4`: WRITE, `5`: CREATE, `6`: DELETE, `7`: ALTER, `8`: DESCRIBE, `9`: CLUSTER_ACTION, `10`: DESCRIBE_CONFIGS, `11`: ALTER_CONFIGS, `12`: IDEMPOTENT_WRITE).
	Operation *int64 `json:"Operation,omitempty" name:"Operation"`

	// Permission type (`2`: DENY, `3`: ALLOW). CKafka currently supports `ALLOW`, which is equivalent to allowlist. `DENY` will be supported for ACLs compatible with open-source Kafka.
	PermissionType *int64 `json:"PermissionType,omitempty" name:"PermissionType"`

	// Resource name, which is related to `resourceType`. For example, if `resourceType` is `TOPIC`, this field indicates the topic name; if `resourceType` is `GROUP`, this field indicates the group name; if `resourceType` is `CLUSTER`, this field can be left empty.
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// The default value is `*`, which means that any host can access. Currently, CKafka does not support the host as `*`, but the future product based on the open-source Kafka will directly support this
	Host *string `json:"Host,omitempty" name:"Host"`

	// The list of users allowed to access the topic. Default: User:*, meaning all users. The current user must be in the user list. Add `User:` before the user name (`User:A` for example).
	Principal *string `json:"Principal,omitempty" name:"Principal"`

	// The resource name list, which is in JSON string format. Either `ResourceName` or `resourceNameList` can be specified.
	ResourceNameList *string `json:"ResourceNameList,omitempty" name:"ResourceNameList"`
}

func (r *CreateAclRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAclRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ResourceType")
	delete(f, "Operation")
	delete(f, "PermissionType")
	delete(f, "ResourceName")
	delete(f, "Host")
	delete(f, "Principal")
	delete(f, "ResourceNameList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAclRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAclResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Returned result
		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAclResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAclResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePartitionRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Topic name
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Number of topic partitions
	PartitionNum *int64 `json:"PartitionNum,omitempty" name:"PartitionNum"`
}

func (r *CreatePartitionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePartitionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "TopicName")
	delete(f, "PartitionNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePartitionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreatePartitionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Returned result set
		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePartitionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePartitionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTopicIpWhiteListRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Topic name
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// IP allowlist list
	IpWhiteList []*string `json:"IpWhiteList,omitempty" name:"IpWhiteList"`
}

func (r *CreateTopicIpWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTopicIpWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "TopicName")
	delete(f, "IpWhiteList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTopicIpWhiteListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateTopicIpWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result of deleting topic IP allowlist
		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTopicIpWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTopicIpWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTopicRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Topic name, which is a string of up to 128 characters. It can contain letters, digits, and hyphens (-) and must start with a letter.
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Number of partitions, which should be greater than 0
	PartitionNum *int64 `json:"PartitionNum,omitempty" name:"PartitionNum"`

	// Number of replicas, which cannot be higher than the number of brokers. Maximum value: 3
	ReplicaNum *int64 `json:"ReplicaNum,omitempty" name:"ReplicaNum"`

	// IP allowlist switch. 1: enabled, 0: disabled. Default value: 0
	EnableWhiteList *int64 `json:"EnableWhiteList,omitempty" name:"EnableWhiteList"`

	// IP allowlist list for quota limit, which is required if `enableWhileList` is 1
	IpWhiteList []*string `json:"IpWhiteList,omitempty" name:"IpWhiteList"`

	// Log cleanup policy, which is `delete` by default. `delete`: logs will be deleted by save time; `compact`: logs will be compressed by key; `compact, delete`: logs will be compressed by key and deleted by save time.
	CleanUpPolicy *string `json:"CleanUpPolicy,omitempty" name:"CleanUpPolicy"`

	// Topic remarks string of up to 64 characters, which must begin with a letter and can contain letters, digits, and dashes (`-`)
	Note *string `json:"Note,omitempty" name:"Note"`

	// Default value: 1
	MinInsyncReplicas *int64 `json:"MinInsyncReplicas,omitempty" name:"MinInsyncReplicas"`

	// Whether to allow an unsynced replica to be elected as leader. false: no, true: yes. Default value: false
	UncleanLeaderElectionEnable *int64 `json:"UncleanLeaderElectionEnable,omitempty" name:"UncleanLeaderElectionEnable"`

	// Message retention period in ms, which is optional. The current minimum value is 60,000 ms
	RetentionMs *int64 `json:"RetentionMs,omitempty" name:"RetentionMs"`

	// Segment rolling duration in ms. The current minimum value is 3,600,000 ms
	SegmentMs *int64 `json:"SegmentMs,omitempty" name:"SegmentMs"`

	// Preset ACL rule. `1`: enable, `0`: disable. Default value: `0`.
	EnableAclRule *int64 `json:"EnableAclRule,omitempty" name:"EnableAclRule"`

	// Name of the preset ACL rule.
	AclRuleName *string `json:"AclRuleName,omitempty" name:"AclRuleName"`

	// Message retention file size in bytes, which is an optional parameter. Default value: -1. Currently, the min value that can be entered is 1,048,576 B.
	RetentionBytes *int64 `json:"RetentionBytes,omitempty" name:"RetentionBytes"`

	// Tag list.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
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
	delete(f, "InstanceId")
	delete(f, "TopicName")
	delete(f, "PartitionNum")
	delete(f, "ReplicaNum")
	delete(f, "EnableWhiteList")
	delete(f, "IpWhiteList")
	delete(f, "CleanUpPolicy")
	delete(f, "Note")
	delete(f, "MinInsyncReplicas")
	delete(f, "UncleanLeaderElectionEnable")
	delete(f, "RetentionMs")
	delete(f, "SegmentMs")
	delete(f, "EnableAclRule")
	delete(f, "AclRuleName")
	delete(f, "RetentionBytes")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateTopicResp struct {

	// Topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

type CreateTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Returned creation result
		Result *CreateTopicResp `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateUserRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Username
	Name *string `json:"Name,omitempty" name:"Name"`

	// User password
	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *CreateUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Name")
	delete(f, "Password")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Returned result
		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAclRequest struct {
	*tchttp.BaseRequest

	// Instance ID information
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// ACL resource type (`2`: TOPIC, `3`: GROUP, `4`: CLUSTER).
	ResourceType *int64 `json:"ResourceType,omitempty" name:"ResourceType"`

	// Resource name, which is related to `resourceType`. For example, if `resourceType` is `TOPIC`, this field indicates the topic name; if `resourceType` is `GROUP`, this field indicates the group name; if `resourceType` is `CLUSTER`, this field can be left empty.
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// ACL operation type (`2`: ALL, `3`: READ, `4`: WRITE, `5`: CREATE, `6`: DELETE, `7`: ALTER, `8`: DESCRIBE, `9`: CLUSTER_ACTION, `10`: DESCRIBE_CONFIGS, `11`: ALTER_CONFIGS, `12`: IDEMPOTENT_WRITE).
	Operation *int64 `json:"Operation,omitempty" name:"Operation"`

	// Permission type (`2`: DENY, `3`: ALLOW). CKafka currently supports `ALLOW`, which is equivalent to allowlist. `DENY` will be supported for ACLs compatible with open-source Kafka.
	PermissionType *int64 `json:"PermissionType,omitempty" name:"PermissionType"`

	// The default value is `*`, which means that any host can access. Currently, CKafka does not support the host as `*`, but the future product based on the open-source Kafka will directly support this
	Host *string `json:"Host,omitempty" name:"Host"`

	// User list. The default value is `*`, which means that any user can access. The current user can only be one included in the user list
	Principal *string `json:"Principal,omitempty" name:"Principal"`
}

func (r *DeleteAclRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAclRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ResourceType")
	delete(f, "ResourceName")
	delete(f, "Operation")
	delete(f, "PermissionType")
	delete(f, "Host")
	delete(f, "Principal")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAclRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAclResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Returned result
		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAclResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAclResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRouteTriggerTimeRequest struct {
	*tchttp.BaseRequest

	// Modification time.
	DelayTime *string `json:"DelayTime,omitempty" name:"DelayTime"`
}

func (r *DeleteRouteTriggerTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRouteTriggerTimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DelayTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRouteTriggerTimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRouteTriggerTimeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRouteTriggerTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRouteTriggerTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTopicIpWhiteListRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Topic name
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// IP allowlist list
	IpWhiteList []*string `json:"IpWhiteList,omitempty" name:"IpWhiteList"`
}

func (r *DeleteTopicIpWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTopicIpWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "TopicName")
	delete(f, "IpWhiteList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTopicIpWhiteListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTopicIpWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result of deleting topic IP allowlist
		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTopicIpWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTopicIpWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTopicRequest struct {
	*tchttp.BaseRequest

	// CKafka instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// CKafka topic name
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
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
	delete(f, "InstanceId")
	delete(f, "TopicName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Returned result set
		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeleteUserRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Username
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DeleteUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Returned result
		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeACLRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// ACL resource type (`2`: TOPIC, `3`: GROUP, `4`: CLUSTER).
	ResourceType *int64 `json:"ResourceType,omitempty" name:"ResourceType"`

	// Resource name, which is related to `resourceType`. For example, if `resourceType` is `TOPIC`, this field indicates the topic name; if `resourceType` is `GROUP`, this field indicates the group name; if `resourceType` is `CLUSTER`, this field can be left empty.
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// Offset position
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Quantity limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Keyword match
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
}

func (r *DescribeACLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeACLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ResourceType")
	delete(f, "ResourceName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchWord")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeACLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeACLResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Returned ACL result set object
		Result *AclResponse `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeACLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeACLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAppInfoRequest struct {
	*tchttp.BaseRequest

	// Offset position
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of users to be queried in this request. Maximum value: 50. Default value: 50
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAppInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAppInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAppInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Returned list of eligible `AppId`
		Result *AppIdResponse `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAppInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCkafkaZoneRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeCkafkaZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCkafkaZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCkafkaZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCkafkaZoneResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Returned results for the query
		Result *ZoneResponse `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCkafkaZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCkafkaZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConsumerGroupRequest struct {
	*tchttp.BaseRequest

	// CKafka instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Name of the group to be queried, which is optional.
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// Name of the corresponding topic in the group to be queried, which is optional. If this parameter is specified but `group` is not specified, this parameter will be ignored.
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Number of results to be returned in this request
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset position
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeConsumerGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConsumerGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "GroupName")
	delete(f, "TopicName")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConsumerGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConsumerGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Returned consumer group information
		Result *ConsumerGroupResponse `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConsumerGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConsumerGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGroup struct {

	// groupId
	Group *string `json:"Group,omitempty" name:"Group"`

	// Protocol used by the group.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

type DescribeGroupInfoRequest struct {
	*tchttp.BaseRequest

	// (Filter) filter by instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Kafka consumer group (`Consumer-group`), which is an array in the format of `GroupList.0=xxx&GroupList.1=yyy`.
	GroupList []*string `json:"GroupList,omitempty" name:"GroupList"`
}

func (r *DescribeGroupInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "GroupList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGroupInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Returned result
	// Note: this field may return null, indicating that no valid values can be obtained.
		Result []*GroupInfoResponse `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGroupInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupOffsetsRequest struct {
	*tchttp.BaseRequest

	// (Filter) filter by instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Kafka consumer group
	Group *string `json:"Group,omitempty" name:"Group"`

	// Array of the names of topics subscribed to by a group. If there is no such array, this parameter means the information of all topics in the specified group
	Topics []*string `json:"Topics,omitempty" name:"Topics"`

	// Fuzzy match by `topicName`
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// Offset position of this query. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of results to be returned in this request. Default value: 50. Maximum value: 50
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeGroupOffsetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupOffsetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Group")
	delete(f, "Topics")
	delete(f, "SearchWord")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGroupOffsetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupOffsetsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Returned result object
		Result *GroupOffsetResponse `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGroupOffsetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupOffsetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Search keyword
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// Offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of results to be returned
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SearchWord")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of returned results
		Result *GroupResponse `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceAttributesRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceAttributesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Returned result object of instance attributes
		Result *InstanceAttributesResponse `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesDetailRequest struct {
	*tchttp.BaseRequest

	// (Filter) filter by instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// (Filter) filter by instance name. Fuzzy search is supported
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// (Filter) instance status. 0: creating, 1: running, 2: deleting. If this parameter is left empty, all instances will be returned by default
	Status []*int64 `json:"Status,omitempty" name:"Status"`

	// Offset. If this parameter is left empty, 0 will be used by default
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results to be returned. If this parameter is left empty, 10 will be used by default. The maximum value is 20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Tag key match.
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Filter
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeInstancesDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SearchWord")
	delete(f, "Status")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "TagKey")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Returned result object of instance details
		Result *InstanceDetailResponse `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest

	// (Filter) filter by instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// (Filter) filter by instance name. Fuzzy search is supported
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// (Filter) instance status. 0: creating, 1: running, 2: deleting. If this parameter is left empty, all instances will be returned by default
	Status []*int64 `json:"Status,omitempty" name:"Status"`

	// Offset. If this parameter is left empty, 0 will be used by default
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results to be returned. If this parameter is left empty, 10 will be used by default. The maximum value is 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Tag key value (this field has been deprecated).
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SearchWord")
	delete(f, "Status")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "TagKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Returned result
		Result *InstanceResponse `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionRequest struct {
	*tchttp.BaseRequest

	// The offset value
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// The maximum number of results returned
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Business field, which can be ignored.
	Business *string `json:"Business,omitempty" name:"Business"`
}

func (r *DescribeRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Business")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRegionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of the returned results of enumerated regions
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
		Result []*Region `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRouteRequest struct {
	*tchttp.BaseRequest

	// Unique instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeRouteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRouteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRouteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRouteResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Returned result set of route information
		Result *RouteResponse `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRouteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicAttributesRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Topic name
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

func (r *DescribeTopicAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "TopicName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopicAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicAttributesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Returned result object
		Result *TopicAttributesResponse `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopicAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicDetailRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// (Filter) filter by `topicName`. Fuzzy search is supported
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// Offset. If this parameter is left empty, 0 will be used by default
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results to be returned. If this parameter is left empty, 10 will be used by default. The maximum value is 20. This value must be greater than 0
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Name of the preset ACL rule.
	AclRuleName *string `json:"AclRuleName,omitempty" name:"AclRuleName"`
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
	delete(f, "InstanceId")
	delete(f, "SearchWord")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "AclRuleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopicDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Returned entity of topic details
		Result *TopicDetailResponse `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeTopicRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Filter by `topicName`. Fuzzy search is supported
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// Offset. If this parameter is left empty, 0 will be used by default
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results to be returned. If this parameter is left empty, 10 will be used by default. The maximum value is 20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Name of the preset ACL rule.
	AclRuleName *string `json:"AclRuleName,omitempty" name:"AclRuleName"`
}

func (r *DescribeTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SearchWord")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "AclRuleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Returned result
	// Note: this field may return null, indicating that no valid values can be obtained.
		Result *TopicResult `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicSubscribeGroupRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Topic name
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Starting position of paging
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results per page
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTopicSubscribeGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicSubscribeGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "TopicName")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopicSubscribeGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicSubscribeGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Returned results
		Result *TopicSubscribeGroup `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopicSubscribeGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicSubscribeGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicSyncReplicaRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Topic name
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Offset. If this parameter is left empty, 0 will be used by default.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results to be returned. If this parameter is left empty, 10 will be used by default. The maximum value is 20.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Filters unsynced replicas only
	OutOfSyncReplicaOnly *bool `json:"OutOfSyncReplicaOnly,omitempty" name:"OutOfSyncReplicaOnly"`
}

func (r *DescribeTopicSyncReplicaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicSyncReplicaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "TopicName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OutOfSyncReplicaOnly")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopicSyncReplicaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicSyncReplicaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Returns topic replica details
		Result *TopicInSyncReplicaResult `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopicSyncReplicaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicSyncReplicaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Filter by name
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// Offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results to be returned in this request
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SearchWord")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Returned result list
		Result *UserResponse `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DynamicDiskConfig struct {

	// Whether to enable dynamic disk expansion configuration. `0`: disable, `1`: enable.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// Percentage of dynamic disk expansion each time.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	StepForwardPercentage *int64 `json:"StepForwardPercentage,omitempty" name:"StepForwardPercentage"`

	// Disk quota threshold (in percentage) for triggering the automatic disk expansion event.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DiskQuotaPercentage *int64 `json:"DiskQuotaPercentage,omitempty" name:"DiskQuotaPercentage"`

	// Max disk space in GB.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	MaxDiskSpace *int64 `json:"MaxDiskSpace,omitempty" name:"MaxDiskSpace"`
}

type DynamicRetentionTime struct {

	// Whether the dynamic message retention time configuration is enabled. 0: disabled; 1: enabled
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// Disk quota threshold (in percentage) for triggering the message retention time change event
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	DiskQuotaPercentage *int64 `json:"DiskQuotaPercentage,omitempty" name:"DiskQuotaPercentage"`

	// Percentage by which the message retention time is shortened each time
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	StepForwardPercentage *int64 `json:"StepForwardPercentage,omitempty" name:"StepForwardPercentage"`

	// Minimum retention time, in minutes
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	BottomRetention *int64 `json:"BottomRetention,omitempty" name:"BottomRetention"`
}

type FetchMessageByOffsetRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Topic name
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// Partition ID
	Partition *int64 `json:"Partition,omitempty" name:"Partition"`

	// Offset information
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *FetchMessageByOffsetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FetchMessageByOffsetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Topic")
	delete(f, "Partition")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FetchMessageByOffsetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type FetchMessageByOffsetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Returned results
		Result *ConsumerRecord `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *FetchMessageByOffsetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FetchMessageByOffsetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// Field to be filtered.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Filter value of field.
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type Group struct {

	// Group name
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

type GroupInfoMember struct {

	// Unique ID generated for consumer in consumer group by coordinator
	MemberId *string `json:"MemberId,omitempty" name:"MemberId"`

	// `client.id` information by the client consumer SDK
	ClientId *string `json:"ClientId,omitempty" name:"ClientId"`

	// Generally stores client IP address
	ClientHost *string `json:"ClientHost,omitempty" name:"ClientHost"`

	// Stores the information of partition assigned to this consumer
	Assignment *Assignment `json:"Assignment,omitempty" name:"Assignment"`
}

type GroupInfoResponse struct {

	// Error code. 0: success
	ErrorCode *string `json:"ErrorCode,omitempty" name:"ErrorCode"`

	// Group status description (common valid values: Empty, Stable, Dead):
	// Dead: the consumer group does not exist
	// Empty: there are currently no consumer subscriptions in the consumer group
	// PreparingRebalance: the consumer group is currently in `rebalance` state
	// CompletingRebalance: the consumer group is currently in `rebalance` state
	// Stable: each consumer in the consumer group has joined and is in stable state
	State *string `json:"State,omitempty" name:"State"`

	// The type of protocol selected by the consumer group, which is `consumer` for common consumers. However, some systems use their own protocols; for example, the protocol used by kafka-connect is `connect`. Only with the standard `consumer` protocol can this API get to know the specific assigning method and parse the specific partition assignment
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

	// Consumer partition assignment algorithm, such as `range` (which is the default value for the Kafka consumer SDK), `roundrobin`, and `sticky`
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// This array contains information only if `state` is `Stable` and `protocol_type` is `consumer`
	Members []*GroupInfoMember `json:"Members,omitempty" name:"Members"`

	// Kafka consumer group
	Group *string `json:"Group,omitempty" name:"Group"`
}

type GroupInfoTopics struct {

	// Name of assigned topics
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// Information of assigned partition
	// Note: this field may return null, indicating that no valid values can be obtained.
	Partitions []*int64 `json:"Partitions,omitempty" name:"Partitions"`
}

type GroupOffsetPartition struct {

	// Topic `partitionId`
	Partition *int64 `json:"Partition,omitempty" name:"Partition"`

	// Offset position submitted by consumer
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Metadata can be passed in for other purposes when the consumer submits messages. Currently, this parameter is usually an empty string
	// Note: this field may return null, indicating that no valid values can be obtained.
	Metadata *string `json:"Metadata,omitempty" name:"Metadata"`

	// Error code
	ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`

	// Latest offset of current partition
	LogEndOffset *int64 `json:"LogEndOffset,omitempty" name:"LogEndOffset"`

	// Number of unconsumed messages
	Lag *int64 `json:"Lag,omitempty" name:"Lag"`
}

type GroupOffsetResponse struct {

	// Total number of eligible results
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Array of partitions in the topic, where each element is a JSON object
	// Note: this field may return null, indicating that no valid values can be obtained.
	TopicList []*GroupOffsetTopic `json:"TopicList,omitempty" name:"TopicList"`
}

type GroupOffsetTopic struct {

	// Topic name
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// Array of partitions in the topic, where each element is a JSON object
	// Note: this field may return null, indicating that no valid values can be obtained.
	Partitions []*GroupOffsetPartition `json:"Partitions,omitempty" name:"Partitions"`
}

type GroupResponse struct {

	// Count
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// GroupList
	// Note: this field may return null, indicating that no valid values can be obtained.
	GroupList []*DescribeGroup `json:"GroupList,omitempty" name:"GroupList"`
}

type Instance struct {

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Instance status. 0: creating, 1: running, 2: deleting, 5: isolated, -1: creation failed
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Whether it is an open-source instance. true: yes, false: no
	// Note: this field may return null, indicating that no valid values can be obtained.
	IfCommunity *bool `json:"IfCommunity,omitempty" name:"IfCommunity"`
}

type InstanceAttributesResponse struct {

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// VIP list information of access point
	VipList []*VipEntity `json:"VipList,omitempty" name:"VipList"`

	// Virtual IP
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Virtual port
	Vport *string `json:"Vport,omitempty" name:"Vport"`

	// Instance status. 0: creating, 1: running, 2: deleting
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Instance bandwidth in Mbps
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// Instance storage capacity in GB
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// AZ
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// VPC ID. If this parameter is empty, it means the basic network
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID. If this parameter is empty, it means the basic network
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Instance health status. 1: healthy, 2: alarmed, 3: exceptional
	Healthy *int64 `json:"Healthy,omitempty" name:"Healthy"`

	// Instance health information. Currently, the disk utilization is displayed with a maximum length of 256
	HealthyMessage *string `json:"HealthyMessage,omitempty" name:"HealthyMessage"`

	// Creation time
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// Message retention period in minutes
	MsgRetentionTime *int64 `json:"MsgRetentionTime,omitempty" name:"MsgRetentionTime"`

	// Configuration for automatic topic creation. If this field is empty, it means that automatic creation is not enabled
	Config *InstanceConfigDO `json:"Config,omitempty" name:"Config"`

	// Number of remaining creatable partitions
	RemainderPartitions *int64 `json:"RemainderPartitions,omitempty" name:"RemainderPartitions"`

	// Number of remaining creatable topics
	RemainderTopics *int64 `json:"RemainderTopics,omitempty" name:"RemainderTopics"`

	// Number of partitions already created
	CreatedPartitions *int64 `json:"CreatedPartitions,omitempty" name:"CreatedPartitions"`

	// Number of topics already created
	CreatedTopics *int64 `json:"CreatedTopics,omitempty" name:"CreatedTopics"`

	// Tag array
	// Note: this field may return null, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Expiration time
	// Note: this field may return null, indicating that no valid values can be obtained.
	ExpireTime *uint64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// Cross-AZ
	// Note: this field may return null, indicating that no valid values can be obtained.
	ZoneIds []*int64 `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// Kafka version information
	// Note: this field may return null, indicating that no valid values can be obtained.
	Version *string `json:"Version,omitempty" name:"Version"`

	// Maximum number of groups
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxGroupNum *int64 `json:"MaxGroupNum,omitempty" name:"MaxGroupNum"`

	// Offering type. `0`: Standard Edition; `1`: Professional Edition
	// Note: this field may return `null`, indicating that no valid value was found.
	Cvm *int64 `json:"Cvm,omitempty" name:"Cvm"`

	// Type.
	// Note: this field may return `null`, indicating that no valid value was found.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Features supported by the instance. `FEATURE_SUBNET_ACL` indicates that the ACL policy supports setting subnets. 
	// Note: this field may return null, indicating that no valid values can be obtained.
	Features []*string `json:"Features,omitempty" name:"Features"`

	// Dynamic message retention policy
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	RetentionTimeConfig *DynamicRetentionTime `json:"RetentionTimeConfig,omitempty" name:"RetentionTimeConfig"`

	// Maximum number of connections
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxConnection *uint64 `json:"MaxConnection,omitempty" name:"MaxConnection"`

	// Public network bandwidth
	// Note: this field may return null, indicating that no valid values can be obtained.
	PublicNetwork *int64 `json:"PublicNetwork,omitempty" name:"PublicNetwork"`

	// Time
	// Note: this field may return null, indicating that no valid values can be obtained.
	DeleteRouteTimestamp *string `json:"DeleteRouteTimestamp,omitempty" name:"DeleteRouteTimestamp"`

	// Number of remaining creatable partitions
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	RemainingPartitions *int64 `json:"RemainingPartitions,omitempty" name:"RemainingPartitions"`

	// Number of remaining creatable topics
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	RemainingTopics *int64 `json:"RemainingTopics,omitempty" name:"RemainingTopics"`

	// Dynamic disk expansion policy.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DynamicDiskConfig *DynamicDiskConfig `json:"DynamicDiskConfig,omitempty" name:"DynamicDiskConfig"`
}

type InstanceConfigDO struct {

	// Whether to create topics automatically
	AutoCreateTopicsEnable *bool `json:"AutoCreateTopicsEnable,omitempty" name:"AutoCreateTopicsEnable"`

	// Number of partitions
	DefaultNumPartitions *int64 `json:"DefaultNumPartitions,omitempty" name:"DefaultNumPartitions"`

	// Default replication factor
	DefaultReplicationFactor *int64 `json:"DefaultReplicationFactor,omitempty" name:"DefaultReplicationFactor"`
}

type InstanceDetail struct {

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Instance VIP information
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Instance port information
	Vport *string `json:"Vport,omitempty" name:"Vport"`

	// Virtual IP list
	VipList []*VipEntity `json:"VipList,omitempty" name:"VipList"`

	// Instance status. 0: creating, 1: running, 2: deleting, 5: isolated, -1: creation failed
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Instance bandwidth in Mbps
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// Instance storage capacity in GB
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// AZ ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// vpcId. If this parameter is empty, it means the basic network
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Whether to renew the instance automatically, which is an int-type enumerated value. 1: yes, 2: no
	RenewFlag *int64 `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// Instance status, which is an int-type value. 0: healthy, 1: alarmed, 2: exceptional
	Healthy *int64 `json:"Healthy,omitempty" name:"Healthy"`

	// Instance status information
	HealthyMessage *string `json:"HealthyMessage,omitempty" name:"HealthyMessage"`

	// Instance creation time
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// Instance expiration time
	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// Whether it is an internal customer. 1: yes
	IsInternal *int64 `json:"IsInternal,omitempty" name:"IsInternal"`

	// Number of topics
	TopicNum *int64 `json:"TopicNum,omitempty" name:"TopicNum"`

	// Tag
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Kafka version information
	// Note: this field may return null, indicating that no valid values can be obtained.
	Version *string `json:"Version,omitempty" name:"Version"`

	// Cross-AZ
	// Note: this field may return null, indicating that no valid values can be obtained.
	ZoneIds []*int64 `json:"ZoneIds,omitempty" name:"ZoneIds"`

	// CKafka sale type
	// Note: this field may return null, indicating that no valid values can be obtained.
	Cvm *int64 `json:"Cvm,omitempty" name:"Cvm"`

	// CKafka instance type
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Disk type
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// Maximum number of topics for the current instance
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	MaxTopicNumber *int64 `json:"MaxTopicNumber,omitempty" name:"MaxTopicNumber"`

	// Maximum number of partitions for the current instance
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	MaxPartitionNumber *int64 `json:"MaxPartitionNumber,omitempty" name:"MaxPartitionNumber"`

	// Time of scheduled upgrade
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	RebalanceTime *string `json:"RebalanceTime,omitempty" name:"RebalanceTime"`
}

type InstanceDetailResponse struct {

	// Total number of eligible instances
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of eligible instance details
	InstanceList []*InstanceDetail `json:"InstanceList,omitempty" name:"InstanceList"`
}

type InstanceResponse struct {

	// List of eligible instances
	// Note: this field may return null, indicating that no valid values can be obtained.
	InstanceList []*Instance `json:"InstanceList,omitempty" name:"InstanceList"`

	// Total number of eligible results
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type JgwOperateResponse struct {

	// Returned code. 0: normal, other values: error
	ReturnCode *string `json:"ReturnCode,omitempty" name:"ReturnCode"`

	// Success message
	ReturnMessage *string `json:"ReturnMessage,omitempty" name:"ReturnMessage"`

	// Data returned by an operation, which may contain `flowId`, etc.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Data *OperateResponseData `json:"Data,omitempty" name:"Data"`
}

type ModifyGroupOffsetsRequest struct {
	*tchttp.BaseRequest

	// Kafka instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Kafka consumer group
	Group *string `json:"Group,omitempty" name:"Group"`

	// Offset resetting policy. Meanings of the input parameters: 0: equivalent to the `shift-by` parameter, which indicates to shift the offset forward or backward by the value of the `shift`. 1: equivalent to `by-duration`, `to-datetime`, `to-earliest`, or `to-latest`, which indicates to move the offset to the specified timestamp. 2: equivalent to `to-offset`, which indicates to move the offset to the specified offset position
	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`

	// Indicates the topics to be reset. If this parameter is left empty, all topics will be reset
	Topics []*string `json:"Topics,omitempty" name:"Topics"`

	// When `strategy` is 0, this field is required. If it is above zero, the offset will be shifted backward by the value of the `shift`. If it is below zero, the offset will be shifted forward by the value of the `shift`. After a correct reset, the new offset should be (old_offset + shift). Note that if the new offset is smaller than the `earliest` parameter of the partition, it will be set to `earliest`, and if it is greater than the `latest` parameter of the partition, it will be set to `latest`
	Shift *int64 `json:"Shift,omitempty" name:"Shift"`

	// Unit: ms. When `strategy` is 1, this field is required, where -2 indicates to reset the offset to the initial position, -1 indicates to reset to the latest position (equivalent to emptying), and other values represent the specified time, i.e., the offset of the topic at the specified time will be obtained and then reset. Note that if there is no message at the specified time, the last offset will be obtained
	ShiftTimestamp *int64 `json:"ShiftTimestamp,omitempty" name:"ShiftTimestamp"`

	// Position of the offset that needs to be reset. When `strategy` is 2, this field is required
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// List of partitions that need to be reset. If the topics parameter is not specified, reset partitions in the corresponding partition list of all topics. If the topics parameter is specified, reset partitions of the corresponding partition list of the specified topic list.
	Partitions []*int64 `json:"Partitions,omitempty" name:"Partitions"`
}

func (r *ModifyGroupOffsetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGroupOffsetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Group")
	delete(f, "Strategy")
	delete(f, "Topics")
	delete(f, "Shift")
	delete(f, "ShiftTimestamp")
	delete(f, "Offset")
	delete(f, "Partitions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyGroupOffsetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyGroupOffsetsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Returned result
		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyGroupOffsetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGroupOffsetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceAttributesConfig struct {

	// Automatic creation. true: enabled, false: not enabled
	AutoCreateTopicEnable *bool `json:"AutoCreateTopicEnable,omitempty" name:"AutoCreateTopicEnable"`

	// Optional. If `auto.create.topic.enable` is set to `true` and this value is not set, 3 will be used by default
	DefaultNumPartitions *int64 `json:"DefaultNumPartitions,omitempty" name:"DefaultNumPartitions"`

	// If `auto.create.topic.enable` is set to `true` but this value is not set, 2 will be used by default
	DefaultReplicationFactor *int64 `json:"DefaultReplicationFactor,omitempty" name:"DefaultReplicationFactor"`
}

type ModifyInstanceAttributesRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Maximum retention period in minutes for instance log, which can be up to 30 days. 0 indicates not to enable the log retention period policy
	MsgRetentionTime *int64 `json:"MsgRetentionTime,omitempty" name:"MsgRetentionTime"`

	// Instance name string of up to 64 characters, which must begin with a letter and can contain letters, digits, and dashes (`-`)
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Instance configuration
	Config *ModifyInstanceAttributesConfig `json:"Config,omitempty" name:"Config"`

	// Dynamic message retention policy configuration
	DynamicRetentionConfig *DynamicRetentionTime `json:"DynamicRetentionConfig,omitempty" name:"DynamicRetentionConfig"`

	// Modification of the rebalancing time after upgrade
	RebalanceTime *int64 `json:"RebalanceTime,omitempty" name:"RebalanceTime"`

	// Timestamp
	PublicNetwork *int64 `json:"PublicNetwork,omitempty" name:"PublicNetwork"`

	// Dynamic disk expansion policy configuration.
	DynamicDiskConfig *DynamicDiskConfig `json:"DynamicDiskConfig,omitempty" name:"DynamicDiskConfig"`
}

func (r *ModifyInstanceAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "MsgRetentionTime")
	delete(f, "InstanceName")
	delete(f, "Config")
	delete(f, "DynamicRetentionConfig")
	delete(f, "RebalanceTime")
	delete(f, "PublicNetwork")
	delete(f, "DynamicDiskConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceAttributesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Returned result
		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPasswordRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Username
	Name *string `json:"Name,omitempty" name:"Name"`

	// Current user password
	Password *string `json:"Password,omitempty" name:"Password"`

	// New user password
	PasswordNew *string `json:"PasswordNew,omitempty" name:"PasswordNew"`
}

func (r *ModifyPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Name")
	delete(f, "Password")
	delete(f, "PasswordNew")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPasswordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Returned result
		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTopicAttributesRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Topic name.
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Topic remarks string of up to 64 characters, which must begin with a letter and can contain letters, digits, and dashes (`-`).
	Note *string `json:"Note,omitempty" name:"Note"`

	// IP allowlist switch. 1: enabled, 0: disabled.
	EnableWhiteList *int64 `json:"EnableWhiteList,omitempty" name:"EnableWhiteList"`

	// Default value: 1.
	MinInsyncReplicas *int64 `json:"MinInsyncReplicas,omitempty" name:"MinInsyncReplicas"`

	// 0: false, 1: true. Default value: 0.
	UncleanLeaderElectionEnable *int64 `json:"UncleanLeaderElectionEnable,omitempty" name:"UncleanLeaderElectionEnable"`

	// Message retention period in ms. The current minimum value is 60,000 ms.
	RetentionMs *int64 `json:"RetentionMs,omitempty" name:"RetentionMs"`

	// Segment rolling duration in ms. The current minimum value is 86,400,000 ms.
	SegmentMs *int64 `json:"SegmentMs,omitempty" name:"SegmentMs"`

	// Maximum topic message length in bytes. The maximum value is 8,388,608 bytes (i.e., 8 MB).
	MaxMessageBytes *int64 `json:"MaxMessageBytes,omitempty" name:"MaxMessageBytes"`

	// Message deletion policy. Valid values: delete, compact
	CleanUpPolicy *string `json:"CleanUpPolicy,omitempty" name:"CleanUpPolicy"`

	// IP allowlist, which is required if the value of `enableWhileList` is 1.
	IpWhiteList []*string `json:"IpWhiteList,omitempty" name:"IpWhiteList"`

	// Preset ACL rule. `1`: enable, `0`: disable. Default value: `0`.
	EnableAclRule *int64 `json:"EnableAclRule,omitempty" name:"EnableAclRule"`

	// Name of the preset ACL rule.
	AclRuleName *string `json:"AclRuleName,omitempty" name:"AclRuleName"`

	// Message retention file size in bytes, which is an optional parameter. Default value: -1. Currently, the min value that can be entered is 1,048,576 B.
	RetentionBytes *int64 `json:"RetentionBytes,omitempty" name:"RetentionBytes"`

	// Tag list.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

func (r *ModifyTopicAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTopicAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "TopicName")
	delete(f, "Note")
	delete(f, "EnableWhiteList")
	delete(f, "MinInsyncReplicas")
	delete(f, "UncleanLeaderElectionEnable")
	delete(f, "RetentionMs")
	delete(f, "SegmentMs")
	delete(f, "MaxMessageBytes")
	delete(f, "CleanUpPolicy")
	delete(f, "IpWhiteList")
	delete(f, "EnableAclRule")
	delete(f, "AclRuleName")
	delete(f, "RetentionBytes")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTopicAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTopicAttributesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Returned result set
		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTopicAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTopicAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperateResponseData struct {

	// FlowId11
	// Note: this field may return null, indicating that no valid values can be obtained.
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
}

type Partition struct {

	// Partition ID
	PartitionId *int64 `json:"PartitionId,omitempty" name:"PartitionId"`
}

type PartitionOffset struct {

	// Partition, such as "0" or "1"
	// Note: this field may return null, indicating that no valid values can be obtained.
	Partition *string `json:"Partition,omitempty" name:"Partition"`

	// Offset, such as 100
	// Note: this field may return null, indicating that no valid values can be obtained.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type Partitions struct {

	// Partition.
	Partition *int64 `json:"Partition,omitempty" name:"Partition"`

	// Partition consumption offset.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type Price struct {

	// Discounted price
	RealTotalCost *float64 `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// Original price
	TotalCost *float64 `json:"TotalCost,omitempty" name:"TotalCost"`
}

type Region struct {

	// Region ID
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// Region name
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// Area name
	AreaName *string `json:"AreaName,omitempty" name:"AreaName"`

	// Region code
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	RegionCode *string `json:"RegionCode,omitempty" name:"RegionCode"`

	// Region code (v3)
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	RegionCodeV3 *string `json:"RegionCodeV3,omitempty" name:"RegionCodeV3"`

	// NONE: no special models are supported by default.\nCVM: the CVM type is supported.
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	Support *string `json:"Support,omitempty" name:"Support"`

	// Whether IPv6 is supported. `0` indicates no, and `1` indicates yes.
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	Ipv6 *int64 `json:"Ipv6,omitempty" name:"Ipv6"`

	// Whether cross-AZ clusters are supported.`0` indicates no, and `1` indicates yes.
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	MultiZone *int64 `json:"MultiZone,omitempty" name:"MultiZone"`
}

type Route struct {

	// Instance connection method
	// 0: PLAINTEXT (plaintext method, which does not carry user information and is supported for legacy versions and Community Edition)
	// 1: SASL_PLAINTEXT (plaintext method, which authenticates the login through SASL before data start and is supported only for Community Edition)
	// 2: SSL (SSL-encrypted communication, which does not carry user information and is supported for legacy versions and Community Edition)
	// 3: SASL_SSL (SSL-encrypted communication, which authenticates the login through SASL before data start and is supported only for Community Edition)
	AccessType *int64 `json:"AccessType,omitempty" name:"AccessType"`

	// Route ID
	RouteId *int64 `json:"RouteId,omitempty" name:"RouteId"`

	// VIP network type (1: public network TGW; 2: classic network; 3: VPC; 4: supporting network (Standard Edition); 5: SSL public network access; 6: BM VPC; 7: supporting network (Pro Edition))
	VipType *int64 `json:"VipType,omitempty" name:"VipType"`

	// Virtual IP list
	VipList []*VipEntity `json:"VipList,omitempty" name:"VipList"`

	// Domain name
	// Note: this field may return null, indicating that no valid values can be obtained.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Domain name port
	// Note: this field may return null, indicating that no valid values can be obtained.
	DomainPort *int64 `json:"DomainPort,omitempty" name:"DomainPort"`

	// Timestamp
	// Note: this field may return null, indicating that no valid values can be obtained.
	DeleteTimestamp *string `json:"DeleteTimestamp,omitempty" name:"DeleteTimestamp"`
}

type RouteResponse struct {

	// Route information list
	// Note: this field may return null, indicating that no valid values can be obtained.
	Routers []*Route `json:"Routers,omitempty" name:"Routers"`
}

type SaleInfo struct {

	// Manually set flag.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Flag *bool `json:"Flag,omitempty" name:"Flag"`

	// CKafka version (v1.1.1/2.4.2/0.10.2）
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Version *string `json:"Version,omitempty" name:"Version"`

	// Whether it is Pro Edition or Standard Edition.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// Whether it has been sold out. `true`: sold out.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SoldOut *bool `json:"SoldOut,omitempty" name:"SoldOut"`
}

type SendMessageRequest struct {
	*tchttp.BaseRequest

	// Datahub access ID.
	DataHubId *string `json:"DataHubId,omitempty" name:"DataHubId"`

	// Message content that is sent.
	Message []*BatchContent `json:"Message,omitempty" name:"Message"`
}

func (r *SendMessageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendMessageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataHubId")
	delete(f, "Message")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendMessageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SendMessageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Message ID list.
		MessageId []*string `json:"MessageId,omitempty" name:"MessageId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendMessageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubscribedInfo struct {

	// Subscribed topic name
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Subscribed partition
	// Note: this field may return null, indicating that no valid values can be obtained.
	Partition []*int64 `json:"Partition,omitempty" name:"Partition"`

	// Partition offset information
	// Note: this field may return null, indicating that no valid values can be obtained.
	PartitionOffset []*PartitionOffset `json:"PartitionOffset,omitempty" name:"PartitionOffset"`

	// ID of the subscribed topic. 
	// Note: this field may return null, indicating that no valid values can be obtained.
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

type Tag struct {

	// Tag key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type Topic struct {

	// Topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Topic name
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Remarks
	// Note: this field may return null, indicating that no valid values can be obtained.
	Note *string `json:"Note,omitempty" name:"Note"`
}

type TopicAttributesResponse struct {

	// Topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Creation time
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// Topic remarks
	// Note: this field may return null, indicating that no valid values can be obtained.
	Note *string `json:"Note,omitempty" name:"Note"`

	// Number of partitions
	PartitionNum *int64 `json:"PartitionNum,omitempty" name:"PartitionNum"`

	// IP allowlist switch. 1: enabled, 0: disabled
	EnableWhiteList *int64 `json:"EnableWhiteList,omitempty" name:"EnableWhiteList"`

	// IP allowlist list
	IpWhiteList []*string `json:"IpWhiteList,omitempty" name:"IpWhiteList"`

	// Topic configuration array
	Config *Config `json:"Config,omitempty" name:"Config"`

	// Partition details
	Partitions []*TopicPartitionDO `json:"Partitions,omitempty" name:"Partitions"`

	// Switch of the preset ACL rule. `1`: enable, `0`: disable.
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	EnableAclRule *int64 `json:"EnableAclRule,omitempty" name:"EnableAclRule"`

	// Preset ACL rule list.
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	AclRuleList []*AclRule `json:"AclRuleList,omitempty" name:"AclRuleList"`
}

type TopicDetail struct {

	// Topic name
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Number of partitions
	PartitionNum *int64 `json:"PartitionNum,omitempty" name:"PartitionNum"`

	// Number of replicas
	ReplicaNum *int64 `json:"ReplicaNum,omitempty" name:"ReplicaNum"`

	// Remarks
	// Note: this field may return null, indicating that no valid values can be obtained.
	Note *string `json:"Note,omitempty" name:"Note"`

	// Creation time
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// Whether to enable IP authentication allowlist. true: yes, false: no
	EnableWhiteList *bool `json:"EnableWhiteList,omitempty" name:"EnableWhiteList"`

	// Number of IPs in IP allowlist
	IpWhiteListCount *int64 `json:"IpWhiteListCount,omitempty" name:"IpWhiteListCount"`

	// COS bucket for data backup: address of the destination COS bucket
	// Note: this field may return null, indicating that no valid values can be obtained.
	ForwardCosBucket *string `json:"ForwardCosBucket,omitempty" name:"ForwardCosBucket"`

	// Status of data backup to COS. 1: not enabled, 0: enabled
	ForwardStatus *int64 `json:"ForwardStatus,omitempty" name:"ForwardStatus"`

	// Frequency of data backup to COS
	ForwardInterval *int64 `json:"ForwardInterval,omitempty" name:"ForwardInterval"`

	// Advanced configuration
	// Note: this field may return null, indicating that no valid values can be obtained.
	Config *Config `json:"Config,omitempty" name:"Config"`

	// Message retention time configuration (for recording the latest retention time)
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	RetentionTimeConfig *TopicRetentionTimeConfigRsp `json:"RetentionTimeConfig,omitempty" name:"RetentionTimeConfig"`

	// `0`: normal, `1`: deleted, `2`: deleting
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type TopicDetailResponse struct {

	// List of returned topic details
	// Note: this field may return null, indicating that no valid values can be obtained.
	TopicList []*TopicDetail `json:"TopicList,omitempty" name:"TopicList"`

	// Number of all eligible topic details
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type TopicInSyncReplicaInfo struct {

	// Partition name
	Partition *string `json:"Partition,omitempty" name:"Partition"`

	// Leader ID
	Leader *uint64 `json:"Leader,omitempty" name:"Leader"`

	// Replica set
	Replica *string `json:"Replica,omitempty" name:"Replica"`

	// ISR
	InSyncReplica *string `json:"InSyncReplica,omitempty" name:"InSyncReplica"`

	// Starting offset
	// Note: this field may return null, indicating that no valid values can be obtained.
	BeginOffset *uint64 `json:"BeginOffset,omitempty" name:"BeginOffset"`

	// Ending offset
	// Note: this field may return null, indicating that no valid values can be obtained.
	EndOffset *uint64 `json:"EndOffset,omitempty" name:"EndOffset"`

	// Number of messages
	// Note: this field may return null, indicating that no valid values can be obtained.
	MessageCount *uint64 `json:"MessageCount,omitempty" name:"MessageCount"`

	// Unsynced replica set
	// Note: this field may return null, indicating that no valid values can be obtained.
	OutOfSyncReplica *string `json:"OutOfSyncReplica,omitempty" name:"OutOfSyncReplica"`
}

type TopicInSyncReplicaResult struct {

	// Set of topic details and replicas
	TopicInSyncReplicaList []*TopicInSyncReplicaInfo `json:"TopicInSyncReplicaList,omitempty" name:"TopicInSyncReplicaList"`

	// Total number
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type TopicPartitionDO struct {

	// Partition ID
	Partition *int64 `json:"Partition,omitempty" name:"Partition"`

	// Leader running status
	LeaderStatus *int64 `json:"LeaderStatus,omitempty" name:"LeaderStatus"`

	// ISR quantity
	IsrNum *int64 `json:"IsrNum,omitempty" name:"IsrNum"`

	// Number of replicas
	ReplicaNum *int64 `json:"ReplicaNum,omitempty" name:"ReplicaNum"`
}

type TopicResult struct {

	// List of returned topic information
	// Note: this field may return null, indicating that no valid values can be obtained.
	TopicList []*Topic `json:"TopicList,omitempty" name:"TopicList"`

	// Number of eligible topics
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type TopicRetentionTimeConfigRsp struct {

	// Expected value, i.e., the topic message retention time (min) configured
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	Expect *int64 `json:"Expect,omitempty" name:"Expect"`

	// Current value (min), i.e., the retention time currently in effect, which may be dynamically adjusted
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	Current *int64 `json:"Current,omitempty" name:"Current"`

	// Last modified time
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	ModTimeStamp *int64 `json:"ModTimeStamp,omitempty" name:"ModTimeStamp"`
}

type TopicSubscribeGroup struct {

	// Total number
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Number of consumer group status
	StatusCountInfo *string `json:"StatusCountInfo,omitempty" name:"StatusCountInfo"`

	// Consumer group information
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	GroupsInfo []*GroupInfoResponse `json:"GroupsInfo,omitempty" name:"GroupsInfo"`

	// Whether a request is asynchronous. If there are fewer consumer groups in the instances, the result will be returned directly, and status code is 1. When there are many consumer groups in the instances, cache will be updated asynchronously. When status code is 0, grouping information will not be returned until cache update is completed and status code becomes 1.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type User struct {

	// User ID
	UserId *int64 `json:"UserId,omitempty" name:"UserId"`

	// Username
	Name *string `json:"Name,omitempty" name:"Name"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last updated time
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type UserResponse struct {

	// List of eligible users
	// Note: this field may return null, indicating that no valid values can be obtained.
	Users []*User `json:"Users,omitempty" name:"Users"`

	// Total number of eligible users
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type VipEntity struct {

	// Virtual IP
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Virtual port
	Vport *string `json:"Vport,omitempty" name:"Vport"`
}

type ZoneInfo struct {

	// Zone ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Whether it is an internal application.
	IsInternalApp *int64 `json:"IsInternalApp,omitempty" name:"IsInternalApp"`

	// Application ID
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// Flag
	Flag *bool `json:"Flag,omitempty" name:"Flag"`

	// Zone name
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// Zone status
	ZoneStatus *int64 `json:"ZoneStatus,omitempty" name:"ZoneStatus"`

	// Extra flag
	Exflag *string `json:"Exflag,omitempty" name:"Exflag"`

	// JSON object. The key is the model. The value `true` means “sold out”, and `false` means “not sold out”.
	SoldOut *string `json:"SoldOut,omitempty" name:"SoldOut"`

	// Information on whether Standard Edition has been sold out.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SalesInfo []*SaleInfo `json:"SalesInfo,omitempty" name:"SalesInfo"`
}

type ZoneResponse struct {

	// Zone list
	ZoneList []*ZoneInfo `json:"ZoneList,omitempty" name:"ZoneList"`

	// Maximum number of instances to be purchased
	MaxBuyInstanceNum *int64 `json:"MaxBuyInstanceNum,omitempty" name:"MaxBuyInstanceNum"`

	// Maximum bandwidth in MB/S
	MaxBandwidth *int64 `json:"MaxBandwidth,omitempty" name:"MaxBandwidth"`

	// Pay-as-you-go unit price
	UnitPrice *Price `json:"UnitPrice,omitempty" name:"UnitPrice"`

	// Pay-as-you-go unit message price
	MessagePrice *Price `json:"MessagePrice,omitempty" name:"MessagePrice"`

	// Cluster information dedicated to a user
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	ClusterInfo []*ClusterInfo `json:"ClusterInfo,omitempty" name:"ClusterInfo"`

	// Purchase of Standard Edition configurations
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	Standard *string `json:"Standard,omitempty" name:"Standard"`

	// Purchase of Standard S2 Edition configurations
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	StandardS2 *string `json:"StandardS2,omitempty" name:"StandardS2"`

	// Purchase of Pro Edition configurations
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	Profession *string `json:"Profession,omitempty" name:"Profession"`

	// Purchase of Physical Dedicated Edition configurations
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	Physical *string `json:"Physical,omitempty" name:"Physical"`

	// Public network bandwidth.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	PublicNetwork *string `json:"PublicNetwork,omitempty" name:"PublicNetwork"`

	// Public network bandwidth configuration.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	PublicNetworkLimit *string `json:"PublicNetworkLimit,omitempty" name:"PublicNetworkLimit"`
}
