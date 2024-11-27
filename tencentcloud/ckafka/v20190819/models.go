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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type Acl struct {
	// ACL resource type. 0: UNKNOWN, 1: ANY, 2: TOPIC, 3: GROUP, 4: CLUSTER, 5: TRANSACTIONAL_ID. Currently, only `TOPIC` is available,
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// Resource name, which is related to `resourceType`. For example, if `resourceType` is `TOPIC`, this field indicates the topic name; if `resourceType` is `GROUP`, this field indicates the group name
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// User list. The default value is `User:*`, which means that any user can access. The current user can only be one included in the user list
	// Note: this field may return null, indicating that no valid values can be obtained.
	Principal *string `json:"Principal,omitnil,omitempty" name:"Principal"`

	// The default value is `*`, which means that any host can access. Currently, CKafka does not support the host as `*`, but the future product based on the open-source Kafka will directly support this
	// Note: this field may return null, indicating that no valid values can be obtained.
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// ACL operation mode. 0: UNKNOWN, 1: ANY, 2: ALL, 3: READ, 4: WRITE, 5: CREATE, 6: DELETE, 7: ALTER, 8: DESCRIBE, 9: CLUSTER_ACTION, 10: DESCRIBE_CONFIGS, 11: ALTER_CONFIGS, 12: IDEMPOTEN_WRITE
	Operation *int64 `json:"Operation,omitnil,omitempty" name:"Operation"`

	// Permission type. 0: UNKNOWN, 1: ANY, 2: DENY, 3: ALLOW
	PermissionType *int64 `json:"PermissionType,omitnil,omitempty" name:"PermissionType"`
}

type AclResponse struct {
	// Number of eligible data entries
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// ACL list
	// Note: this field may return null, indicating that no valid values can be obtained.
	AclList []*Acl `json:"AclList,omitnil,omitempty" name:"AclList"`
}

type AclRule struct {
	// ACL rule name.
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// Instance ID.
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Matching type. Currently, only prefix match is supported. Enumerated value list: PREFIXED
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	PatternType *string `json:"PatternType,omitnil,omitempty" name:"PatternType"`

	// Prefix value for prefix match.
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	Pattern *string `json:"Pattern,omitnil,omitempty" name:"Pattern"`

	// ACL resource type. Only “Topic” is supported. Enumerated value list: Topic.
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// ACL information contained in the rule.
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	AclList *string `json:"AclList,omitnil,omitempty" name:"AclList"`

	// Creation time of the rule.
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	CreateTimeStamp *string `json:"CreateTimeStamp,omitnil,omitempty" name:"CreateTimeStamp"`

	// A parameter used to specify whether the preset ACL rule is applied to new topics.
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	IsApplied *int64 `json:"IsApplied,omitnil,omitempty" name:"IsApplied"`

	// Rule update time.
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	UpdateTimeStamp *string `json:"UpdateTimeStamp,omitnil,omitempty" name:"UpdateTimeStamp"`

	// Remarks of the rule.
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// One of the corresponding topic names that is displayed.
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// The number of topics that apply this ACL rule.
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	TopicCount *int64 `json:"TopicCount,omitnil,omitempty" name:"TopicCount"`

	// Name of rule type.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	PatternTypeTitle *string `json:"PatternTypeTitle,omitnil,omitempty" name:"PatternTypeTitle"`
}

type AclRuleInfo struct {
	// ACL operation types. Enumerated values: `All` (all operations), `Read` (read), `Write` (write).
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// Permission types: `Deny`, `Allow`.
	PermissionType *string `json:"PermissionType,omitnil,omitempty" name:"PermissionType"`

	// The default value is `*`, which means that any host can access the topic. CKafka currently does not support specifying a host value of * or an IP range.
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// The list of users allowed to access the topic. Default value: `User:*`, which means all users. The current user must be in the user list. Add the prefix `User:` before the user name (`User:A`, for example).
	Principal *string `json:"Principal,omitnil,omitempty" name:"Principal"`
}

type AclRuleResp struct {
	// Total number of data entries
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// ACL rule list
	// Note: This field may return null, indicating that no valid values can be obtained.
	AclRuleList []*AclRule `json:"AclRuleList,omitnil,omitempty" name:"AclRuleList"`
}

type AppIdResponse struct {
	// Number of eligible `AppId`
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of eligible `AppId`
	// Note: this field may return null, indicating that no valid values can be obtained.
	AppIdList []*int64 `json:"AppIdList,omitnil,omitempty" name:"AppIdList"`
}

type Assignment struct {
	// Assignment version information
	Version *int64 `json:"Version,omitnil,omitempty" name:"Version"`

	// Topic information list
	// Note: this field may return null, indicating that no valid values can be obtained.
	Topics []*GroupInfoTopics `json:"Topics,omitnil,omitempty" name:"Topics"`
}

type BatchContent struct {
	// Message body that is sent.
	Body *string `json:"Body,omitnil,omitempty" name:"Body"`

	// Message sending key name.
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`
}

// Predefined struct for user
type BatchCreateAclRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ACL resource type. Default value: `2` (topic).
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// Resource list array.
	ResourceNames []*string `json:"ResourceNames,omitnil,omitempty" name:"ResourceNames"`

	// ACL rule list.
	RuleList []*AclRuleInfo `json:"RuleList,omitnil,omitempty" name:"RuleList"`
}

type BatchCreateAclRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ACL resource type. Default value: `2` (topic).
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// Resource list array.
	ResourceNames []*string `json:"ResourceNames,omitnil,omitempty" name:"ResourceNames"`

	// ACL rule list.
	RuleList []*AclRuleInfo `json:"RuleList,omitnil,omitempty" name:"RuleList"`
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

// Predefined struct for user
type BatchCreateAclResponseParams struct {
	// Status code.
	Result *int64 `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BatchCreateAclResponse struct {
	*tchttp.BaseResponse
	Response *BatchCreateAclResponseParams `json:"Response"`
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

// Predefined struct for user
type BatchModifyGroupOffsetsRequestParams struct {
	// Consumer group name.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Instance name.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Partition information.
	Partitions []*Partitions `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// Name of the specified topic. Default value: names of all topics.
	TopicName []*string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

type BatchModifyGroupOffsetsRequest struct {
	*tchttp.BaseRequest
	
	// Consumer group name.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Instance name.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Partition information.
	Partitions []*Partitions `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// Name of the specified topic. Default value: names of all topics.
	TopicName []*string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
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

// Predefined struct for user
type BatchModifyGroupOffsetsResponseParams struct {
	// Returned result.
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BatchModifyGroupOffsetsResponse struct {
	*tchttp.BaseResponse
	Response *BatchModifyGroupOffsetsResponseParams `json:"Response"`
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

// Predefined struct for user
type BatchModifyTopicAttributesRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Topic attribute list
	Topic []*BatchModifyTopicInfo `json:"Topic,omitnil,omitempty" name:"Topic"`
}

type BatchModifyTopicAttributesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Topic attribute list
	Topic []*BatchModifyTopicInfo `json:"Topic,omitnil,omitempty" name:"Topic"`
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

// Predefined struct for user
type BatchModifyTopicAttributesResponseParams struct {
	// Returned result.
	Result []*BatchModifyTopicResultDTO `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BatchModifyTopicAttributesResponse struct {
	*tchttp.BaseResponse
	Response *BatchModifyTopicAttributesResponseParams `json:"Response"`
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
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// The number of partitions.
	PartitionNum *int64 `json:"PartitionNum,omitnil,omitempty" name:"PartitionNum"`

	// Remarks.
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// Number of replicas.
	ReplicaNum *int64 `json:"ReplicaNum,omitnil,omitempty" name:"ReplicaNum"`

	// Message deletion policy. Valid values: `delete`, `compact`.
	CleanUpPolicy *string `json:"CleanUpPolicy,omitnil,omitempty" name:"CleanUpPolicy"`

	// The minimum number of replicas specified by `min.insync.replicas` when the producer sets `request.required.acks` to `-1`.
	MinInsyncReplicas *int64 `json:"MinInsyncReplicas,omitnil,omitempty" name:"MinInsyncReplicas"`

	// Whether to allow a non-ISR replica to be the leader.
	UncleanLeaderElectionEnable *bool `json:"UncleanLeaderElectionEnable,omitnil,omitempty" name:"UncleanLeaderElectionEnable"`

	// Message retention period in topic dimension in milliseconds. Value range: 1 minute to 90 days.
	RetentionMs *int64 `json:"RetentionMs,omitnil,omitempty" name:"RetentionMs"`

	// Message retention size in topic dimension. Value range: 1 MB - 1024 GB.
	RetentionBytes *int64 `json:"RetentionBytes,omitnil,omitempty" name:"RetentionBytes"`

	// Segment rolling duration in milliseconds. Value range: 1-90 days.
	SegmentMs *int64 `json:"SegmentMs,omitnil,omitempty" name:"SegmentMs"`

	// Message size per batch. Value range: 1 KB - 12 MB.
	MaxMessageBytes *int64 `json:"MaxMessageBytes,omitnil,omitempty" name:"MaxMessageBytes"`
}

type BatchModifyTopicResultDTO struct {
	// Instance ID.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Topic name.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Status code.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ReturnCode *string `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// Message status.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

type ClusterInfo struct {
	// Cluster ID
	ClusterId *int64 `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Cluster name
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// The cluster’s maximum disk capacity in GB
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	MaxDiskSize *int64 `json:"MaxDiskSize,omitnil,omitempty" name:"MaxDiskSize"`

	// The cluster’s maximum bandwidth in MB/s
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	MaxBandWidth *int64 `json:"MaxBandWidth,omitnil,omitempty" name:"MaxBandWidth"`

	// The cluster’s available disk capacity in GB
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	AvailableDiskSize *int64 `json:"AvailableDiskSize,omitnil,omitempty" name:"AvailableDiskSize"`

	// The cluster’s available bandwidth in MB/s
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	AvailableBandWidth *int64 `json:"AvailableBandWidth,omitnil,omitempty" name:"AvailableBandWidth"`

	// The AZ where the cluster resides
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The AZ where the cluster nodes reside. If the cluster is a multi-AZ cluster, this field means multiple AZs where the cluster nodes reside.
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	ZoneIds []*int64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`
}

type Config struct {
	// Message retention period
	// Note: this field may return null, indicating that no valid values can be obtained.
	Retention *int64 `json:"Retention,omitnil,omitempty" name:"Retention"`

	// Minimum number of sync replications
	// Note: this field may return null, indicating that no valid values can be obtained.
	MinInsyncReplicas *int64 `json:"MinInsyncReplicas,omitnil,omitempty" name:"MinInsyncReplicas"`

	// Log cleanup mode. Default value: delete.
	// delete: logs will be deleted by save time; compact: logs will be compressed by key; compact, delete: logs will be compressed by key and deleted by save time.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CleanUpPolicy *string `json:"CleanUpPolicy,omitnil,omitempty" name:"CleanUpPolicy"`

	// Segment rolling duration
	// Note: this field may return null, indicating that no valid values can be obtained.
	SegmentMs *int64 `json:"SegmentMs,omitnil,omitempty" name:"SegmentMs"`

	// 0: false, 1: true.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UncleanLeaderElectionEnable *int64 `json:"UncleanLeaderElectionEnable,omitnil,omitempty" name:"UncleanLeaderElectionEnable"`

	// Number of bytes for segment rolling
	// Note: this field may return null, indicating that no valid values can be obtained.
	SegmentBytes *int64 `json:"SegmentBytes,omitnil,omitempty" name:"SegmentBytes"`

	// Maximum number of message bytes
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxMessageBytes *int64 `json:"MaxMessageBytes,omitnil,omitempty" name:"MaxMessageBytes"`

	// Message retention file size.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	RetentionBytes *int64 `json:"RetentionBytes,omitnil,omitempty" name:"RetentionBytes"`
}

type ConsumerGroup struct {
	// User group name
	ConsumerGroupName *string `json:"ConsumerGroupName,omitnil,omitempty" name:"ConsumerGroupName"`

	// Subscribed message entity
	SubscribedInfo []*SubscribedInfo `json:"SubscribedInfo,omitnil,omitempty" name:"SubscribedInfo"`
}

type ConsumerGroupResponse struct {
	// Number of eligible consumer groups
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Topic list
	// Note: this field may return null, indicating that no valid values can be obtained.
	TopicList []*ConsumerGroupTopic `json:"TopicList,omitnil,omitempty" name:"TopicList"`

	// Consumer group list
	// Note: this field may return null, indicating that no valid values can be obtained.
	GroupList []*ConsumerGroup `json:"GroupList,omitnil,omitempty" name:"GroupList"`

	// Total number of partitions
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalPartition *int64 `json:"TotalPartition,omitnil,omitempty" name:"TotalPartition"`

	// List of monitored partitions
	// Note: this field may return null, indicating that no valid values can be obtained.
	PartitionListForMonitor []*Partition `json:"PartitionListForMonitor,omitnil,omitempty" name:"PartitionListForMonitor"`

	// Total number of topics
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalTopic *int64 `json:"TotalTopic,omitnil,omitempty" name:"TotalTopic"`

	// List of monitored topics
	// Note: this field may return null, indicating that no valid values can be obtained.
	TopicListForMonitor []*ConsumerGroupTopic `json:"TopicListForMonitor,omitnil,omitempty" name:"TopicListForMonitor"`

	// List of monitored groups
	// Note: this field may return null, indicating that no valid values can be obtained.
	GroupListForMonitor []*Group `json:"GroupListForMonitor,omitnil,omitempty" name:"GroupListForMonitor"`
}

type ConsumerGroupTopic struct {
	// Topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Topic name
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

type ConsumerRecord struct {
	// Topic name
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// Partition ID
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Message key
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Message value
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// Message timestamp
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// Message headers
	// Note: This field may return null, indicating that no valid values can be obtained.
	Headers *string `json:"Headers,omitnil,omitempty" name:"Headers"`
}

// Predefined struct for user
type CreateAclRequestParams struct {
	// Instance ID information
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ACL resource type (`2`: TOPIC, `3`: GROUP, `4`: CLUSTER).
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// ACL operation type (`2`: ALL, `3`: READ, `4`: WRITE, `5`: CREATE, `6`: DELETE, `7`: ALTER, `8`: DESCRIBE, `9`: CLUSTER_ACTION, `10`: DESCRIBE_CONFIGS, `11`: ALTER_CONFIGS, `12`: IDEMPOTENT_WRITE).
	Operation *int64 `json:"Operation,omitnil,omitempty" name:"Operation"`

	// Permission type (`2`: DENY, `3`: ALLOW). CKafka currently supports `ALLOW`, which is equivalent to allowlist. `DENY` will be supported for ACLs compatible with open-source Kafka.
	PermissionType *int64 `json:"PermissionType,omitnil,omitempty" name:"PermissionType"`

	// Resource name, which is related to `resourceType`. For example, if `resourceType` is `TOPIC`, this field indicates the topic name; if `resourceType` is `GROUP`, this field indicates the group name; if `resourceType` is `CLUSTER`, this field can be left empty.
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// The default value is `*`, which means that any host can access. Currently, CKafka does not support the host as `*`, but the future product based on the open-source Kafka will directly support this
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// The list of users allowed to access the topic. Default: User:*, meaning all users. The current user must be in the user list. Add `User:` before the user name (`User:A` for example).
	Principal *string `json:"Principal,omitnil,omitempty" name:"Principal"`

	// The resource name list, which is in JSON string format. Either `ResourceName` or `resourceNameList` can be specified.
	ResourceNameList *string `json:"ResourceNameList,omitnil,omitempty" name:"ResourceNameList"`
}

type CreateAclRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID information
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ACL resource type (`2`: TOPIC, `3`: GROUP, `4`: CLUSTER).
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// ACL operation type (`2`: ALL, `3`: READ, `4`: WRITE, `5`: CREATE, `6`: DELETE, `7`: ALTER, `8`: DESCRIBE, `9`: CLUSTER_ACTION, `10`: DESCRIBE_CONFIGS, `11`: ALTER_CONFIGS, `12`: IDEMPOTENT_WRITE).
	Operation *int64 `json:"Operation,omitnil,omitempty" name:"Operation"`

	// Permission type (`2`: DENY, `3`: ALLOW). CKafka currently supports `ALLOW`, which is equivalent to allowlist. `DENY` will be supported for ACLs compatible with open-source Kafka.
	PermissionType *int64 `json:"PermissionType,omitnil,omitempty" name:"PermissionType"`

	// Resource name, which is related to `resourceType`. For example, if `resourceType` is `TOPIC`, this field indicates the topic name; if `resourceType` is `GROUP`, this field indicates the group name; if `resourceType` is `CLUSTER`, this field can be left empty.
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// The default value is `*`, which means that any host can access. Currently, CKafka does not support the host as `*`, but the future product based on the open-source Kafka will directly support this
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// The list of users allowed to access the topic. Default: User:*, meaning all users. The current user must be in the user list. Add `User:` before the user name (`User:A` for example).
	Principal *string `json:"Principal,omitnil,omitempty" name:"Principal"`

	// The resource name list, which is in JSON string format. Either `ResourceName` or `resourceNameList` can be specified.
	ResourceNameList *string `json:"ResourceNameList,omitnil,omitempty" name:"ResourceNameList"`
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

// Predefined struct for user
type CreateAclResponseParams struct {
	// Returned result
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAclResponse struct {
	*tchttp.BaseResponse
	Response *CreateAclResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateAclRuleRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ACL resource type. Currently, the only valid value is `Topic`.
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// Matching type. Valid values: `PREFIXED`(match by prefix), `PRESET` (match by preset policy).
	PatternType *string `json:"PatternType,omitnil,omitempty" name:"PatternType"`

	// Rule name
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// ACL rule list
	RuleList []*AclRuleInfo `json:"RuleList,omitnil,omitempty" name:"RuleList"`

	// Prefix value for prefix match
	Pattern *string `json:"Pattern,omitnil,omitempty" name:"Pattern"`

	// A parameter used to specify whether the preset ACL rule is applied to new topics
	IsApplied *int64 `json:"IsApplied,omitnil,omitempty" name:"IsApplied"`

	// Remarks for ACL rules
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

type CreateAclRuleRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ACL resource type. Currently, the only valid value is `Topic`.
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// Matching type. Valid values: `PREFIXED`(match by prefix), `PRESET` (match by preset policy).
	PatternType *string `json:"PatternType,omitnil,omitempty" name:"PatternType"`

	// Rule name
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// ACL rule list
	RuleList []*AclRuleInfo `json:"RuleList,omitnil,omitempty" name:"RuleList"`

	// Prefix value for prefix match
	Pattern *string `json:"Pattern,omitnil,omitempty" name:"Pattern"`

	// A parameter used to specify whether the preset ACL rule is applied to new topics
	IsApplied *int64 `json:"IsApplied,omitnil,omitempty" name:"IsApplied"`

	// Remarks for ACL rules
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

func (r *CreateAclRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAclRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ResourceType")
	delete(f, "PatternType")
	delete(f, "RuleName")
	delete(f, "RuleList")
	delete(f, "Pattern")
	delete(f, "IsApplied")
	delete(f, "Comment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAclRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAclRuleResponseParams struct {
	// Unique key of a rule
	Result *int64 `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAclRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateAclRuleResponseParams `json:"Response"`
}

func (r *CreateAclRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAclRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConsumerRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Group name.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Topic name. You must specify the name of an existing topic for either `TopicName` or `TopicNameList`.
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Topic name array.
	TopicNameList []*string `json:"TopicNameList,omitnil,omitempty" name:"TopicNameList"`
}

type CreateConsumerRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Group name.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Topic name. You must specify the name of an existing topic for either `TopicName` or `TopicNameList`.
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Topic name array.
	TopicNameList []*string `json:"TopicNameList,omitnil,omitempty" name:"TopicNameList"`
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
	delete(f, "InstanceId")
	delete(f, "GroupName")
	delete(f, "TopicName")
	delete(f, "TopicNameList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateConsumerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConsumerResponseParams struct {
	// Description of the created consumer group.
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

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
type CreateDatahubTopicRequestParams struct {
	// Topic name, which is a string of up to 128 characters. It can contain letters, digits, and hyphens (-) and must start with a letter.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Number of partitions, which should be greater than 0.
	PartitionNum *int64 `json:"PartitionNum,omitnil,omitempty" name:"PartitionNum"`

	// Message retention period in milliseconds. The current minimum value is 60,000 ms.
	RetentionMs *int64 `json:"RetentionMs,omitnil,omitempty" name:"RetentionMs"`

	// Topic remarks, which are a string of up to 128 characters. It can contain letters, digits, and hyphens (-) and must start with a letter.
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// Tag list
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateDatahubTopicRequest struct {
	*tchttp.BaseRequest
	
	// Topic name, which is a string of up to 128 characters. It can contain letters, digits, and hyphens (-) and must start with a letter.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Number of partitions, which should be greater than 0.
	PartitionNum *int64 `json:"PartitionNum,omitnil,omitempty" name:"PartitionNum"`

	// Message retention period in milliseconds. The current minimum value is 60,000 ms.
	RetentionMs *int64 `json:"RetentionMs,omitnil,omitempty" name:"RetentionMs"`

	// Topic remarks, which are a string of up to 128 characters. It can contain letters, digits, and hyphens (-) and must start with a letter.
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// Tag list
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreateDatahubTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDatahubTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "PartitionNum")
	delete(f, "RetentionMs")
	delete(f, "Note")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDatahubTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDatahubTopicResponseParams struct {
	// Returned creation result
	Result *DatahubTopicResp `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDatahubTopicResponse struct {
	*tchttp.BaseResponse
	Response *CreateDatahubTopicResponseParams `json:"Response"`
}

func (r *CreateDatahubTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDatahubTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstancePostData struct {
	// This parameter has a fixed value of 0 returned by `CreateInstancePre`. It is only used for backend data alignment  and cannot be used as the query condition for `CheckTaskStatus`. 
	// Note:  This field may return null, indicating that no valid values can be obtained.
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// List of order IDs Note: This field may return null, indicating that no valid values can be obtained.
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// Instance ID. When multiple instances are purchased, the ID of the first one is returned by default . Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Mapping between orders and the purchased instances.  Note: This field may return null, indicating that no valid values can be obtained.
	DealNameInstanceIdMapping []*DealInstanceDTO `json:"DealNameInstanceIdMapping,omitnil,omitempty" name:"DealNameInstanceIdMapping"`
}

// Predefined struct for user
type CreateInstancePostRequestParams struct {
	// Instance name, which is a string of up to 64 characters. It can contain letters, digits, and hyphens (-) and must start with a letter.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Private network peak bandwidth of an instance  in MB/sec.  If you create a Standard Edition instance, pass in the corresponding peak bandwidth for the current instance specification.  If you create a Pro Edition instance, configure the peak bandwidth, partition count, and other parameters as required by Pro Edition.
	BandWidth *int64 `json:"BandWidth,omitnil,omitempty" name:"BandWidth"`

	// ID of the VPC where the default access point of the created instance resides.  This parameter is required as instances cannot be created in the classic network currently.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// ID of the subnet  where the default access point of the created instance resides. 
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Instance specification.  This parameter is required for a Standard Edition instance but not for a Pro Edition instance.  Valid values:  `1` (Small),  `2` (Standard),  `3` (Advanced),  `4` (Large),  `5` (Xlarge L1),  `6` (Xlarge L2),  `7` (Xlarge L3),  `8` (Xlarge L4),  
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// The maximum instance log retention period in minutes by default.  If this parameter is left empty, the default retention period is 1,440 minutes (1 day) to 30 days.  If the message retention period of the topic is explicitly set, it will prevail.
	MsgRetentionTime *int64 `json:"MsgRetentionTime,omitnil,omitempty" name:"MsgRetentionTime"`

	// Cluster ID, which can be selected when you create an instance.  You don’t need to pass in this parameter if the cluster where the instance resides is not specified.
	ClusterId *int64 `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Instance version.  Valid values: `0.10.2`, `1.1.1`, `2.4.2`, and `2.8.1`.
	KafkaVersion *string `json:"KafkaVersion,omitnil,omitempty" name:"KafkaVersion"`

	// Instance type. Valid values: `standard` (Standard Edition),  `profession`  (Pro Edition)
	SpecificationsType *string `json:"SpecificationsType,omitnil,omitempty" name:"SpecificationsType"`

	// Instance disk type. Valid values:  `CLOUD_BASIC` (Premium Cloud Storage),  `CLOUD_SSD` (SSD).  If this parameter is left empty, the default value `CLOUD_BASIC` will be used.
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// Instance disk size, which must meet the requirement of the instance’s specification.
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// The maximum number of partitions of the instance, which must meet the requirement of the instance’s specification.
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// The maximum number of topics of the instance, which must meet the requirement of the instance’s specification.
	TopicNum *int64 `json:"TopicNum,omitnil,omitempty" name:"TopicNum"`

	// AZ of the instance.  When a multi-AZ instance is created, the value of this parameter is the AZ ID of the subnet where the instance’s default access point resides.
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Whether the current instance is a multi-AZ instance
	MultiZoneFlag *bool `json:"MultiZoneFlag,omitnil,omitempty" name:"MultiZoneFlag"`

	// This parameter indicates the list of AZ IDs when the instance is deployed in multiple AZs.  Note that `ZoneId` must be included in the array of this parameter.
	ZoneIds []*int64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// The number of purchased instances.  Default value: `1`. This parameter is optional.  If it is passed in, multiple instances will be created, with their names being `instanceName` plus different suffixes.
	InstanceNum *int64 `json:"InstanceNum,omitnil,omitempty" name:"InstanceNum"`

	// Public network bandwidth in Mbps.  The 3 Mbps of free bandwidth is not included here by default.  For example, if you need 3 Mbps of public network bandwidth, pass in `0`; if you need 6 Mbps, pass in `3`. The value must be an integer multiple of 3.
	PublicNetworkMonthly *int64 `json:"PublicNetworkMonthly,omitnil,omitempty" name:"PublicNetworkMonthly"`
}

type CreateInstancePostRequest struct {
	*tchttp.BaseRequest
	
	// Instance name, which is a string of up to 64 characters. It can contain letters, digits, and hyphens (-) and must start with a letter.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Private network peak bandwidth of an instance  in MB/sec.  If you create a Standard Edition instance, pass in the corresponding peak bandwidth for the current instance specification.  If you create a Pro Edition instance, configure the peak bandwidth, partition count, and other parameters as required by Pro Edition.
	BandWidth *int64 `json:"BandWidth,omitnil,omitempty" name:"BandWidth"`

	// ID of the VPC where the default access point of the created instance resides.  This parameter is required as instances cannot be created in the classic network currently.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// ID of the subnet  where the default access point of the created instance resides. 
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Instance specification.  This parameter is required for a Standard Edition instance but not for a Pro Edition instance.  Valid values:  `1` (Small),  `2` (Standard),  `3` (Advanced),  `4` (Large),  `5` (Xlarge L1),  `6` (Xlarge L2),  `7` (Xlarge L3),  `8` (Xlarge L4),  
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// The maximum instance log retention period in minutes by default.  If this parameter is left empty, the default retention period is 1,440 minutes (1 day) to 30 days.  If the message retention period of the topic is explicitly set, it will prevail.
	MsgRetentionTime *int64 `json:"MsgRetentionTime,omitnil,omitempty" name:"MsgRetentionTime"`

	// Cluster ID, which can be selected when you create an instance.  You don’t need to pass in this parameter if the cluster where the instance resides is not specified.
	ClusterId *int64 `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Instance version.  Valid values: `0.10.2`, `1.1.1`, `2.4.2`, and `2.8.1`.
	KafkaVersion *string `json:"KafkaVersion,omitnil,omitempty" name:"KafkaVersion"`

	// Instance type. Valid values: `standard` (Standard Edition),  `profession`  (Pro Edition)
	SpecificationsType *string `json:"SpecificationsType,omitnil,omitempty" name:"SpecificationsType"`

	// Instance disk type. Valid values:  `CLOUD_BASIC` (Premium Cloud Storage),  `CLOUD_SSD` (SSD).  If this parameter is left empty, the default value `CLOUD_BASIC` will be used.
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// Instance disk size, which must meet the requirement of the instance’s specification.
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// The maximum number of partitions of the instance, which must meet the requirement of the instance’s specification.
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// The maximum number of topics of the instance, which must meet the requirement of the instance’s specification.
	TopicNum *int64 `json:"TopicNum,omitnil,omitempty" name:"TopicNum"`

	// AZ of the instance.  When a multi-AZ instance is created, the value of this parameter is the AZ ID of the subnet where the instance’s default access point resides.
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Whether the current instance is a multi-AZ instance
	MultiZoneFlag *bool `json:"MultiZoneFlag,omitnil,omitempty" name:"MultiZoneFlag"`

	// This parameter indicates the list of AZ IDs when the instance is deployed in multiple AZs.  Note that `ZoneId` must be included in the array of this parameter.
	ZoneIds []*int64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// The number of purchased instances.  Default value: `1`. This parameter is optional.  If it is passed in, multiple instances will be created, with their names being `instanceName` plus different suffixes.
	InstanceNum *int64 `json:"InstanceNum,omitnil,omitempty" name:"InstanceNum"`

	// Public network bandwidth in Mbps.  The 3 Mbps of free bandwidth is not included here by default.  For example, if you need 3 Mbps of public network bandwidth, pass in `0`; if you need 6 Mbps, pass in `3`. The value must be an integer multiple of 3.
	PublicNetworkMonthly *int64 `json:"PublicNetworkMonthly,omitnil,omitempty" name:"PublicNetworkMonthly"`
}

func (r *CreateInstancePostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstancePostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceName")
	delete(f, "BandWidth")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "InstanceType")
	delete(f, "MsgRetentionTime")
	delete(f, "ClusterId")
	delete(f, "KafkaVersion")
	delete(f, "SpecificationsType")
	delete(f, "DiskType")
	delete(f, "DiskSize")
	delete(f, "Partition")
	delete(f, "TopicNum")
	delete(f, "ZoneId")
	delete(f, "MultiZoneFlag")
	delete(f, "ZoneIds")
	delete(f, "InstanceNum")
	delete(f, "PublicNetworkMonthly")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstancePostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstancePostResp struct {
	// Returned code. `0` indicates normal status while other codes indicate errors.
	ReturnCode *string `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// Message returned by the API. An error message will be returned if the API reports an error. 
	ReturnMessage *string `json:"ReturnMessage,omitnil,omitempty" name:"ReturnMessage"`

	// Returned data.  Note: This field may return null, indicating that no valid values can be obtained.
	Data *CreateInstancePostData `json:"Data,omitnil,omitempty" name:"Data"`
}

// Predefined struct for user
type CreateInstancePostResponseParams struct {
	// Returned result
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateInstancePostResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstancePostResponseParams `json:"Response"`
}

func (r *CreateInstancePostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstancePostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstancePreData struct {
	// The value returned by `CreateInstancePre` is 0, which is fixed and cannot be used as the query condition of `CheckTaskStatus`. It is only used to ensure the consistency with the backend data structure.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Order number list.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// Instance ID. When multiple instances are purchased, the ID of the first one is returned by default . Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Mapping between orders and the purchased instances.  Note: This field may return null, indicating that no valid values can be obtained.
	DealNameInstanceIdMapping []*DealInstanceDTO `json:"DealNameInstanceIdMapping,omitnil,omitempty" name:"DealNameInstanceIdMapping"`
}

type CreateInstancePreResp struct {
	// Returned code. 0: Normal; other values: Error.
	ReturnCode *string `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// The message indicating whether the operation is successful.
	ReturnMessage *string `json:"ReturnMessage,omitnil,omitempty" name:"ReturnMessage"`

	// Data returned by the operation.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Data *CreateInstancePreData `json:"Data,omitnil,omitempty" name:"Data"`

	// Deletion time.  This parameter has been deprecated and will be deleted.  Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: DeleteRouteTimestamp is deprecated.
	DeleteRouteTimestamp *string `json:"DeleteRouteTimestamp,omitnil,omitempty" name:"DeleteRouteTimestamp"`
}

// Predefined struct for user
type CreatePartitionRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Topic name
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Number of topic partitions
	PartitionNum *int64 `json:"PartitionNum,omitnil,omitempty" name:"PartitionNum"`
}

type CreatePartitionRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Topic name
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Number of topic partitions
	PartitionNum *int64 `json:"PartitionNum,omitnil,omitempty" name:"PartitionNum"`
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

// Predefined struct for user
type CreatePartitionResponseParams struct {
	// Returned result set
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePartitionResponse struct {
	*tchttp.BaseResponse
	Response *CreatePartitionResponseParams `json:"Response"`
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

// Predefined struct for user
type CreatePostPaidInstanceRequestParams struct {
	// Instance name, which is a string of up to 64 letters, digits, and hyphens (-). It must start with a letter.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// ID of the VPC where the default access point of the created instance resides.  This parameter is required as instances cannot be created in the classic network currently.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// ID of the subnet  where the default access point of the created instance resides.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Instance specification.  This parameter is required for a Standard Edition instance but not for a Pro Edition instance.  Valid values:  `1` (Small),  `2` (Standard),  `3` (Advanced),  `4` (Large),  `5` (Xlarge L1),  `6` (Xlarge L2),  `7` (Xlarge L3),  `8` (Xlarge L4),  
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// The maximum instance log retention period in minutes by default.  If this parameter is left empty, the default retention period is 1,440 minutes (1 day) to 30 days.  If the message retention period of the topic is explicitly set, it will prevail.
	MsgRetentionTime *int64 `json:"MsgRetentionTime,omitnil,omitempty" name:"MsgRetentionTime"`

	// Cluster ID, which can be selected when you create an instance.  You don’t need to pass in this parameter if the cluster where the instance resides is not specified.
	ClusterId *int64 `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Instance version.  Valid values: `0.10.2`, `1.1.1`, `2.4.2`, and `2.8.1`.
	KafkaVersion *string `json:"KafkaVersion,omitnil,omitempty" name:"KafkaVersion"`

	// Instance type. `standard` (Standard Edition),  `profession`  (Pro Edition)
	SpecificationsType *string `json:"SpecificationsType,omitnil,omitempty" name:"SpecificationsType"`

	// Instance disk type.  `CLOUD_BASIC` (Premium Cloud Storage),  `CLOUD_SSD` (SSD).  If this parameter is left empty, the default value `CLOUD_BASIC` will be used.
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// Private network peak bandwidth of an instance  in MB/sec.  If you create a Standard Edition instance, pass in the corresponding peak bandwidth for the current instance specification.  If you create a Pro Edition instance, configure the peak bandwidth, partition count, and other parameters as required by Pro Edition.
	BandWidth *int64 `json:"BandWidth,omitnil,omitempty" name:"BandWidth"`

	// Instance disk size, which must meet the requirement of the instance’s specification.
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// The maximum number of partitions of the instance, which must meet the requirement of the instance’s specification.
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// The maximum number of topics of the instance, which must meet the requirement of the instance’s specification.
	TopicNum *int64 `json:"TopicNum,omitnil,omitempty" name:"TopicNum"`

	// AZ of the instance.  When a multi-AZ instance is created, the value of this parameter is the AZ ID of the subnet where the instance’s default access point resides.
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Whether the current instance is a multi-AZ instance
	MultiZoneFlag *bool `json:"MultiZoneFlag,omitnil,omitempty" name:"MultiZoneFlag"`

	// This parameter indicates the list of AZ IDs when the instance is deployed in multiple AZs.  Note that `ZoneId` must be included in the array of this parameter.
	ZoneIds []*int64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// The number of purchased instances.  Default value: `1`. This parameter is optional.  If it is passed in, multiple instances will be created, with their names being `instanceName` plus different suffixes.
	InstanceNum *int64 `json:"InstanceNum,omitnil,omitempty" name:"InstanceNum"`

	// Public network bandwidth in Mbps.  The 3 Mbps of free bandwidth is not included here by default.  For example, if you need 3 Mbps of public network bandwidth, pass in `0`; if you need 6 Mbps, pass in `3`.  The value must be an integer multiple of 3.
	PublicNetworkMonthly *int64 `json:"PublicNetworkMonthly,omitnil,omitempty" name:"PublicNetworkMonthly"`
}

type CreatePostPaidInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance name, which is a string of up to 64 letters, digits, and hyphens (-). It must start with a letter.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// ID of the VPC where the default access point of the created instance resides.  This parameter is required as instances cannot be created in the classic network currently.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// ID of the subnet  where the default access point of the created instance resides.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Instance specification.  This parameter is required for a Standard Edition instance but not for a Pro Edition instance.  Valid values:  `1` (Small),  `2` (Standard),  `3` (Advanced),  `4` (Large),  `5` (Xlarge L1),  `6` (Xlarge L2),  `7` (Xlarge L3),  `8` (Xlarge L4),  
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// The maximum instance log retention period in minutes by default.  If this parameter is left empty, the default retention period is 1,440 minutes (1 day) to 30 days.  If the message retention period of the topic is explicitly set, it will prevail.
	MsgRetentionTime *int64 `json:"MsgRetentionTime,omitnil,omitempty" name:"MsgRetentionTime"`

	// Cluster ID, which can be selected when you create an instance.  You don’t need to pass in this parameter if the cluster where the instance resides is not specified.
	ClusterId *int64 `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Instance version.  Valid values: `0.10.2`, `1.1.1`, `2.4.2`, and `2.8.1`.
	KafkaVersion *string `json:"KafkaVersion,omitnil,omitempty" name:"KafkaVersion"`

	// Instance type. `standard` (Standard Edition),  `profession`  (Pro Edition)
	SpecificationsType *string `json:"SpecificationsType,omitnil,omitempty" name:"SpecificationsType"`

	// Instance disk type.  `CLOUD_BASIC` (Premium Cloud Storage),  `CLOUD_SSD` (SSD).  If this parameter is left empty, the default value `CLOUD_BASIC` will be used.
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// Private network peak bandwidth of an instance  in MB/sec.  If you create a Standard Edition instance, pass in the corresponding peak bandwidth for the current instance specification.  If you create a Pro Edition instance, configure the peak bandwidth, partition count, and other parameters as required by Pro Edition.
	BandWidth *int64 `json:"BandWidth,omitnil,omitempty" name:"BandWidth"`

	// Instance disk size, which must meet the requirement of the instance’s specification.
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// The maximum number of partitions of the instance, which must meet the requirement of the instance’s specification.
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// The maximum number of topics of the instance, which must meet the requirement of the instance’s specification.
	TopicNum *int64 `json:"TopicNum,omitnil,omitempty" name:"TopicNum"`

	// AZ of the instance.  When a multi-AZ instance is created, the value of this parameter is the AZ ID of the subnet where the instance’s default access point resides.
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Whether the current instance is a multi-AZ instance
	MultiZoneFlag *bool `json:"MultiZoneFlag,omitnil,omitempty" name:"MultiZoneFlag"`

	// This parameter indicates the list of AZ IDs when the instance is deployed in multiple AZs.  Note that `ZoneId` must be included in the array of this parameter.
	ZoneIds []*int64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// The number of purchased instances.  Default value: `1`. This parameter is optional.  If it is passed in, multiple instances will be created, with their names being `instanceName` plus different suffixes.
	InstanceNum *int64 `json:"InstanceNum,omitnil,omitempty" name:"InstanceNum"`

	// Public network bandwidth in Mbps.  The 3 Mbps of free bandwidth is not included here by default.  For example, if you need 3 Mbps of public network bandwidth, pass in `0`; if you need 6 Mbps, pass in `3`.  The value must be an integer multiple of 3.
	PublicNetworkMonthly *int64 `json:"PublicNetworkMonthly,omitnil,omitempty" name:"PublicNetworkMonthly"`
}

func (r *CreatePostPaidInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePostPaidInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceName")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "InstanceType")
	delete(f, "MsgRetentionTime")
	delete(f, "ClusterId")
	delete(f, "KafkaVersion")
	delete(f, "SpecificationsType")
	delete(f, "DiskType")
	delete(f, "BandWidth")
	delete(f, "DiskSize")
	delete(f, "Partition")
	delete(f, "TopicNum")
	delete(f, "ZoneId")
	delete(f, "MultiZoneFlag")
	delete(f, "ZoneIds")
	delete(f, "InstanceNum")
	delete(f, "PublicNetworkMonthly")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePostPaidInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePostPaidInstanceResponseParams struct {
	// Returned result
	Result *CreateInstancePostResp `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePostPaidInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreatePostPaidInstanceResponseParams `json:"Response"`
}

func (r *CreatePostPaidInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePostPaidInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTopicIpWhiteListRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Topic name
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// IP allowlist list
	IpWhiteList []*string `json:"IpWhiteList,omitnil,omitempty" name:"IpWhiteList"`
}

type CreateTopicIpWhiteListRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Topic name
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// IP allowlist list
	IpWhiteList []*string `json:"IpWhiteList,omitnil,omitempty" name:"IpWhiteList"`
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

// Predefined struct for user
type CreateTopicIpWhiteListResponseParams struct {
	// Result of deleting topic IP allowlist
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTopicIpWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *CreateTopicIpWhiteListResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateTopicRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Topic name, which is a string of up to 128 characters. It can contain letters, digits, and hyphens (-) and must start with a letter.
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Number of partitions, which should be greater than 0
	PartitionNum *int64 `json:"PartitionNum,omitnil,omitempty" name:"PartitionNum"`

	// Number of replicas, which cannot be higher than the number of brokers. Maximum value: 3
	ReplicaNum *int64 `json:"ReplicaNum,omitnil,omitempty" name:"ReplicaNum"`

	// IP allowlist switch. 1: enabled, 0: disabled. Default value: 0
	EnableWhiteList *int64 `json:"EnableWhiteList,omitnil,omitempty" name:"EnableWhiteList"`

	// IP allowlist list for quota limit, which is required if `enableWhileList` is 1
	IpWhiteList []*string `json:"IpWhiteList,omitnil,omitempty" name:"IpWhiteList"`

	// Log cleanup policy, which is `delete` by default. `delete`: logs will be deleted by save time; `compact`: logs will be compressed by key; `compact, delete`: logs will be compressed by key and deleted by save time.
	CleanUpPolicy *string `json:"CleanUpPolicy,omitnil,omitempty" name:"CleanUpPolicy"`

	// Topic remarks string of up to 64 characters, which must begin with a letter and can contain letters, digits, and dashes (`-`)
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// Default value: 1
	MinInsyncReplicas *int64 `json:"MinInsyncReplicas,omitnil,omitempty" name:"MinInsyncReplicas"`

	// Whether to allow an unsynced replica to be elected as leader. false: no, true: yes. Default value: false
	UncleanLeaderElectionEnable *int64 `json:"UncleanLeaderElectionEnable,omitnil,omitempty" name:"UncleanLeaderElectionEnable"`

	// Message retention period in milliseconds, which is optional. Min value: 60,000 ms.
	RetentionMs *int64 `json:"RetentionMs,omitnil,omitempty" name:"RetentionMs"`

	// Segment rolling duration in ms. The current minimum value is 3,600,000 ms
	SegmentMs *int64 `json:"SegmentMs,omitnil,omitempty" name:"SegmentMs"`

	// Max message size in bytes. Value range: 1,024 bytes (1 KB) to 8,388,608 bytes (8 MB).
	MaxMessageBytes *int64 `json:"MaxMessageBytes,omitnil,omitempty" name:"MaxMessageBytes"`

	// Preset ACL rule. `1`: enable, `0`: disable. Default value: `0`.
	EnableAclRule *int64 `json:"EnableAclRule,omitnil,omitempty" name:"EnableAclRule"`

	// Name of the preset ACL rule.
	AclRuleName *string `json:"AclRuleName,omitnil,omitempty" name:"AclRuleName"`

	// Message retention file size in bytes, which is an optional parameter. Default value: -1. Currently, the min value that can be entered is 1,048,576 B.
	RetentionBytes *int64 `json:"RetentionBytes,omitnil,omitempty" name:"RetentionBytes"`

	// Tag list.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateTopicRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Topic name, which is a string of up to 128 characters. It can contain letters, digits, and hyphens (-) and must start with a letter.
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Number of partitions, which should be greater than 0
	PartitionNum *int64 `json:"PartitionNum,omitnil,omitempty" name:"PartitionNum"`

	// Number of replicas, which cannot be higher than the number of brokers. Maximum value: 3
	ReplicaNum *int64 `json:"ReplicaNum,omitnil,omitempty" name:"ReplicaNum"`

	// IP allowlist switch. 1: enabled, 0: disabled. Default value: 0
	EnableWhiteList *int64 `json:"EnableWhiteList,omitnil,omitempty" name:"EnableWhiteList"`

	// IP allowlist list for quota limit, which is required if `enableWhileList` is 1
	IpWhiteList []*string `json:"IpWhiteList,omitnil,omitempty" name:"IpWhiteList"`

	// Log cleanup policy, which is `delete` by default. `delete`: logs will be deleted by save time; `compact`: logs will be compressed by key; `compact, delete`: logs will be compressed by key and deleted by save time.
	CleanUpPolicy *string `json:"CleanUpPolicy,omitnil,omitempty" name:"CleanUpPolicy"`

	// Topic remarks string of up to 64 characters, which must begin with a letter and can contain letters, digits, and dashes (`-`)
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// Default value: 1
	MinInsyncReplicas *int64 `json:"MinInsyncReplicas,omitnil,omitempty" name:"MinInsyncReplicas"`

	// Whether to allow an unsynced replica to be elected as leader. false: no, true: yes. Default value: false
	UncleanLeaderElectionEnable *int64 `json:"UncleanLeaderElectionEnable,omitnil,omitempty" name:"UncleanLeaderElectionEnable"`

	// Message retention period in milliseconds, which is optional. Min value: 60,000 ms.
	RetentionMs *int64 `json:"RetentionMs,omitnil,omitempty" name:"RetentionMs"`

	// Segment rolling duration in ms. The current minimum value is 3,600,000 ms
	SegmentMs *int64 `json:"SegmentMs,omitnil,omitempty" name:"SegmentMs"`

	// Max message size in bytes. Value range: 1,024 bytes (1 KB) to 8,388,608 bytes (8 MB).
	MaxMessageBytes *int64 `json:"MaxMessageBytes,omitnil,omitempty" name:"MaxMessageBytes"`

	// Preset ACL rule. `1`: enable, `0`: disable. Default value: `0`.
	EnableAclRule *int64 `json:"EnableAclRule,omitnil,omitempty" name:"EnableAclRule"`

	// Name of the preset ACL rule.
	AclRuleName *string `json:"AclRuleName,omitnil,omitempty" name:"AclRuleName"`

	// Message retention file size in bytes, which is an optional parameter. Default value: -1. Currently, the min value that can be entered is 1,048,576 B.
	RetentionBytes *int64 `json:"RetentionBytes,omitnil,omitempty" name:"RetentionBytes"`

	// Tag list.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
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
	delete(f, "MaxMessageBytes")
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
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

// Predefined struct for user
type CreateTopicResponseParams struct {
	// Returned creation result
	Result *CreateTopicResp `json:"Result,omitnil,omitempty" name:"Result"`

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

// Predefined struct for user
type CreateUserRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Username
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// User password
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

type CreateUserRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Username
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// User password
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
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

// Predefined struct for user
type CreateUserResponseParams struct {
	// Returned result
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateUserResponse struct {
	*tchttp.BaseResponse
	Response *CreateUserResponseParams `json:"Response"`
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

type DatahubTopicDTO struct {
	// Name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Topic name
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// The number of partitions
	PartitionNum *uint64 `json:"PartitionNum,omitnil,omitempty" name:"PartitionNum"`

	// Expiration time
	RetentionMs *uint64 `json:"RetentionMs,omitnil,omitempty" name:"RetentionMs"`

	// Remarks
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// Status (`1`: In use; `2`: Deleting)
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type DatahubTopicResp struct {
	// Topic name
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// TopicId
	// Note: This field may return null, indicating that no valid values can be obtained.
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

type DealInstanceDTO struct {
	// Order list.  Note: This field may return null, indicating that no valid values can be obtained.
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// ID list of the purchased CKafka instances corresponding to the order list.  Note: This field may return null, indicating that no valid values can be obtained.
	InstanceIdList []*string `json:"InstanceIdList,omitnil,omitempty" name:"InstanceIdList"`
}

// Predefined struct for user
type DeleteAclRequestParams struct {
	// Instance ID information
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ACL resource type (`2`: TOPIC, `3`: GROUP, `4`: CLUSTER).
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// Resource name, which is related to `resourceType`. For example, if `resourceType` is `TOPIC`, this field indicates the topic name; if `resourceType` is `GROUP`, this field indicates the group name; if `resourceType` is `CLUSTER`, this field can be left empty.
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// ACL operation type (`2`: ALL, `3`: READ, `4`: WRITE, `5`: CREATE, `6`: DELETE, `7`: ALTER, `8`: DESCRIBE, `9`: CLUSTER_ACTION, `10`: DESCRIBE_CONFIGS, `11`: ALTER_CONFIGS, `12`: IDEMPOTENT_WRITE).
	Operation *int64 `json:"Operation,omitnil,omitempty" name:"Operation"`

	// Permission type (`2`: DENY, `3`: ALLOW). CKafka currently supports `ALLOW`, which is equivalent to allowlist. `DENY` will be supported for ACLs compatible with open-source Kafka.
	PermissionType *int64 `json:"PermissionType,omitnil,omitempty" name:"PermissionType"`

	// The default value is `*`, which means that any host can access. Currently, CKafka does not support the host as `*`, but the future product based on the open-source Kafka will directly support this
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// User list. The default value is `*`, which means that any user can access. The current user can only be one included in the user list
	Principal *string `json:"Principal,omitnil,omitempty" name:"Principal"`
}

type DeleteAclRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID information
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ACL resource type (`2`: TOPIC, `3`: GROUP, `4`: CLUSTER).
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// Resource name, which is related to `resourceType`. For example, if `resourceType` is `TOPIC`, this field indicates the topic name; if `resourceType` is `GROUP`, this field indicates the group name; if `resourceType` is `CLUSTER`, this field can be left empty.
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// ACL operation type (`2`: ALL, `3`: READ, `4`: WRITE, `5`: CREATE, `6`: DELETE, `7`: ALTER, `8`: DESCRIBE, `9`: CLUSTER_ACTION, `10`: DESCRIBE_CONFIGS, `11`: ALTER_CONFIGS, `12`: IDEMPOTENT_WRITE).
	Operation *int64 `json:"Operation,omitnil,omitempty" name:"Operation"`

	// Permission type (`2`: DENY, `3`: ALLOW). CKafka currently supports `ALLOW`, which is equivalent to allowlist. `DENY` will be supported for ACLs compatible with open-source Kafka.
	PermissionType *int64 `json:"PermissionType,omitnil,omitempty" name:"PermissionType"`

	// The default value is `*`, which means that any host can access. Currently, CKafka does not support the host as `*`, but the future product based on the open-source Kafka will directly support this
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// User list. The default value is `*`, which means that any user can access. The current user can only be one included in the user list
	Principal *string `json:"Principal,omitnil,omitempty" name:"Principal"`
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

// Predefined struct for user
type DeleteAclResponseParams struct {
	// Returned result
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAclResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAclResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteInstancePreRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DeleteInstancePreRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DeleteInstancePreRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstancePreRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteInstancePreRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstancePreResponseParams struct {
	// Returned result
	Result *CreateInstancePreResp `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteInstancePreResponse struct {
	*tchttp.BaseResponse
	Response *DeleteInstancePreResponseParams `json:"Response"`
}

func (r *DeleteInstancePreResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstancePreResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRouteRequestParams struct {
	// Unique instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Route ID.
	RouteId *int64 `json:"RouteId,omitnil,omitempty" name:"RouteId"`

	// AppId of the caller.
	CallerAppid *int64 `json:"CallerAppid,omitnil,omitempty" name:"CallerAppid"`

	// The time when a route was deleted.
	DeleteRouteTime *string `json:"DeleteRouteTime,omitnil,omitempty" name:"DeleteRouteTime"`
}

type DeleteRouteRequest struct {
	*tchttp.BaseRequest
	
	// Unique instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Route ID.
	RouteId *int64 `json:"RouteId,omitnil,omitempty" name:"RouteId"`

	// AppId of the caller.
	CallerAppid *int64 `json:"CallerAppid,omitnil,omitempty" name:"CallerAppid"`

	// The time when a route was deleted.
	DeleteRouteTime *string `json:"DeleteRouteTime,omitnil,omitempty" name:"DeleteRouteTime"`
}

func (r *DeleteRouteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRouteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "RouteId")
	delete(f, "CallerAppid")
	delete(f, "DeleteRouteTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRouteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRouteResponseParams struct {
	// Returned result.
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRouteResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRouteResponseParams `json:"Response"`
}

func (r *DeleteRouteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRouteTriggerTimeRequestParams struct {
	// Modification time.
	DelayTime *string `json:"DelayTime,omitnil,omitempty" name:"DelayTime"`
}

type DeleteRouteTriggerTimeRequest struct {
	*tchttp.BaseRequest
	
	// Modification time.
	DelayTime *string `json:"DelayTime,omitnil,omitempty" name:"DelayTime"`
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

// Predefined struct for user
type DeleteRouteTriggerTimeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRouteTriggerTimeResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRouteTriggerTimeResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteTopicIpWhiteListRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Topic name
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// IP allowlist list
	IpWhiteList []*string `json:"IpWhiteList,omitnil,omitempty" name:"IpWhiteList"`
}

type DeleteTopicIpWhiteListRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Topic name
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// IP allowlist list
	IpWhiteList []*string `json:"IpWhiteList,omitnil,omitempty" name:"IpWhiteList"`
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

// Predefined struct for user
type DeleteTopicIpWhiteListResponseParams struct {
	// Result of deleting topic IP allowlist
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTopicIpWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTopicIpWhiteListResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteTopicRequestParams struct {
	// CKafka instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// CKafka topic name
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

type DeleteTopicRequest struct {
	*tchttp.BaseRequest
	
	// CKafka instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// CKafka topic name
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
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

// Predefined struct for user
type DeleteTopicResponseParams struct {
	// Returned result set
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

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
type DeleteUserRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Username
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DeleteUserRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Username
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
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

// Predefined struct for user
type DeleteUserResponseParams struct {
	// Returned result
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteUserResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUserResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeACLRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ACL resource type (`2`: TOPIC, `3`: GROUP, `4`: CLUSTER).
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// Resource name, which is related to `resourceType`. For example, if `resourceType` is `TOPIC`, this field indicates the topic name; if `resourceType` is `GROUP`, this field indicates the group name; if `resourceType` is `CLUSTER`, this field can be left empty.
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// Offset position
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Quantity limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Keyword match
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`
}

type DescribeACLRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ACL resource type (`2`: TOPIC, `3`: GROUP, `4`: CLUSTER).
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// Resource name, which is related to `resourceType`. For example, if `resourceType` is `TOPIC`, this field indicates the topic name; if `resourceType` is `GROUP`, this field indicates the group name; if `resourceType` is `CLUSTER`, this field can be left empty.
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// Offset position
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Quantity limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Keyword match
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`
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

// Predefined struct for user
type DescribeACLResponseParams struct {
	// Returned ACL result set object
	Result *AclResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeACLResponse struct {
	*tchttp.BaseResponse
	Response *DescribeACLResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeAclRuleRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ACL rule name
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// ACL rule matching type
	PatternType *string `json:"PatternType,omitnil,omitempty" name:"PatternType"`

	// Whether to read simplified ACL rules
	IsSimplified *bool `json:"IsSimplified,omitnil,omitempty" name:"IsSimplified"`
}

type DescribeAclRuleRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ACL rule name
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// ACL rule matching type
	PatternType *string `json:"PatternType,omitnil,omitempty" name:"PatternType"`

	// Whether to read simplified ACL rules
	IsSimplified *bool `json:"IsSimplified,omitnil,omitempty" name:"IsSimplified"`
}

func (r *DescribeAclRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAclRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "RuleName")
	delete(f, "PatternType")
	delete(f, "IsSimplified")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAclRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAclRuleResponseParams struct {
	// The set of returned ACL rules
	Result *AclRuleResp `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAclRuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAclRuleResponseParams `json:"Response"`
}

func (r *DescribeAclRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAclRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppInfoRequestParams struct {
	// Offset position
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of users to be queried in this request. Maximum value: 50. Default value: 50
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeAppInfoRequest struct {
	*tchttp.BaseRequest
	
	// Offset position
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of users to be queried in this request. Maximum value: 50. Default value: 50
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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

// Predefined struct for user
type DescribeAppInfoResponseParams struct {
	// Returned list of eligible `AppId`
	Result *AppIdResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAppInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAppInfoResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeCkafkaZoneRequestParams struct {
	// Cloud Dedicated Cluster (CDC) business parameter.
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`
}

type DescribeCkafkaZoneRequest struct {
	*tchttp.BaseRequest
	
	// Cloud Dedicated Cluster (CDC) business parameter.
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`
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
	delete(f, "CdcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCkafkaZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCkafkaZoneResponseParams struct {
	// Returned results for the query
	Result *ZoneResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCkafkaZoneResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCkafkaZoneResponseParams `json:"Response"`
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

type DescribeConnectInfoResultDTO struct {
	// IP address
	// Note: This field may return null, indicating that no valid values can be obtained.
	IpAddr *string `json:"IpAddr,omitnil,omitempty" name:"IpAddr"`

	// Connection time
	// Note: This field may return null, indicating that no valid values can be obtained.
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// Whether it is a supported version
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsUnSupportVersion *bool `json:"IsUnSupportVersion,omitnil,omitempty" name:"IsUnSupportVersion"`
}

// Predefined struct for user
type DescribeConsumerGroupRequestParams struct {
	// CKafka instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Name of the group to be queried, which is optional.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Name of the corresponding topic in the group to be queried, which is optional. If this parameter is specified but `group` is not specified, this parameter will be ignored.
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Number of results to be returned in this request
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset position
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeConsumerGroupRequest struct {
	*tchttp.BaseRequest
	
	// CKafka instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Name of the group to be queried, which is optional.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Name of the corresponding topic in the group to be queried, which is optional. If this parameter is specified but `group` is not specified, this parameter will be ignored.
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Number of results to be returned in this request
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset position
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
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

// Predefined struct for user
type DescribeConsumerGroupResponseParams struct {
	// Returned consumer group information
	Result *ConsumerGroupResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConsumerGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConsumerGroupResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeDatahubTopicRequestParams struct {
	// Name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DescribeDatahubTopicRequest struct {
	*tchttp.BaseRequest
	
	// Name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *DescribeDatahubTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatahubTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatahubTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDatahubTopicResp struct {
	// Name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Topic name
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// The number of partitions
	PartitionNum *uint64 `json:"PartitionNum,omitnil,omitempty" name:"PartitionNum"`

	// Expiration time
	RetentionMs *uint64 `json:"RetentionMs,omitnil,omitempty" name:"RetentionMs"`

	// Remarks
	// Note: This field may return null, indicating that no valid values can be obtained.
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// Username
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Password
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Status (`1`: In use; `2`: Deleting)
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Service routing address
	// Note: This field may return null, indicating that no valid values can be obtained.
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`
}

// Predefined struct for user
type DescribeDatahubTopicResponseParams struct {
	// Returned result object
	Result *DescribeDatahubTopicResp `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDatahubTopicResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatahubTopicResponseParams `json:"Response"`
}

func (r *DescribeDatahubTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatahubTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatahubTopicsRequestParams struct {
	// Keyword for query
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// Query offset, which defaults to `0`.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of results to be returned in this request. Default value: `50`. Maximum value: `50`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeDatahubTopicsRequest struct {
	*tchttp.BaseRequest
	
	// Keyword for query
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// Query offset, which defaults to `0`.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of results to be returned in this request. Default value: `50`. Maximum value: `50`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeDatahubTopicsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatahubTopicsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SearchWord")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatahubTopicsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDatahubTopicsResp struct {
	// Total count
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Topic list
	// Note: This field may return null, indicating that no valid values can be obtained.
	TopicList []*DatahubTopicDTO `json:"TopicList,omitnil,omitempty" name:"TopicList"`
}

// Predefined struct for user
type DescribeDatahubTopicsResponseParams struct {
	// Topic list
	// Note: This field may return null, indicating that no valid values can be obtained.
	Result *DescribeDatahubTopicsResp `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDatahubTopicsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatahubTopicsResponseParams `json:"Response"`
}

func (r *DescribeDatahubTopicsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatahubTopicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGroup struct {
	// groupId
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// Protocol used by the group.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`
}

// Predefined struct for user
type DescribeGroupInfoRequestParams struct {
	// (Filter) filter by instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Kafka consumer group (`Consumer-group`), which is an array in the format of `GroupList.0=xxx&GroupList.1=yyy`.
	GroupList []*string `json:"GroupList,omitnil,omitempty" name:"GroupList"`
}

type DescribeGroupInfoRequest struct {
	*tchttp.BaseRequest
	
	// (Filter) filter by instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Kafka consumer group (`Consumer-group`), which is an array in the format of `GroupList.0=xxx&GroupList.1=yyy`.
	GroupList []*string `json:"GroupList,omitnil,omitempty" name:"GroupList"`
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

// Predefined struct for user
type DescribeGroupInfoResponseParams struct {
	// Returned result
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result []*GroupInfoResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGroupInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGroupInfoResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeGroupOffsetsRequestParams struct {
	// (Filter) filter by instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Kafka consumer group
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// Array of the names of topics subscribed to by a group. If there is no such array, this parameter means the information of all topics in the specified group
	Topics []*string `json:"Topics,omitnil,omitempty" name:"Topics"`

	// Fuzzy match by `topicName`
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// Offset position of this query. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of results to be returned in this request. Default value: 50. Maximum value: 50
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeGroupOffsetsRequest struct {
	*tchttp.BaseRequest
	
	// (Filter) filter by instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Kafka consumer group
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// Array of the names of topics subscribed to by a group. If there is no such array, this parameter means the information of all topics in the specified group
	Topics []*string `json:"Topics,omitnil,omitempty" name:"Topics"`

	// Fuzzy match by `topicName`
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// Offset position of this query. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of results to be returned in this request. Default value: 50. Maximum value: 50
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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

// Predefined struct for user
type DescribeGroupOffsetsResponseParams struct {
	// Returned result object
	Result *GroupOffsetResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGroupOffsetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGroupOffsetsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeGroupRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Search keyword
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of results to be returned
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeGroupRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Search keyword
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of results to be returned
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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

// Predefined struct for user
type DescribeGroupResponseParams struct {
	// List of returned results
	Result *GroupResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGroupResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeInstanceAttributesRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceAttributesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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

// Predefined struct for user
type DescribeInstanceAttributesResponseParams struct {
	// Returned result object of instance attributes
	Result *InstanceAttributesResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceAttributesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceAttributesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeInstancesDetailRequestParams struct {
	// (Filter) filter by instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Filter by instance name, instance ID, AZ, VPC ID, or subnet ID. Fuzzy query is supported.
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// (Filter) instance status. 0: creating, 1: running, 2: deleting. If this parameter is left empty, all instances will be returned by default
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Offset. If this parameter is left empty, `0` will be used by default.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. If this parameter is left empty, `10` will be used by default. The maximum value is `20`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Tag key match.
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Filter. Valid values of `filter.Name` include `Ip`, `VpcId`, `SubNetId`, `InstanceType`, and `InstanceId`. Up to 10 values can be passed for `filter.Values`.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// This parameter has been deprecated and replaced with `InstanceIdList`.
	InstanceIds *string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Filter by instance ID.
	InstanceIdList []*string `json:"InstanceIdList,omitnil,omitempty" name:"InstanceIdList"`

	// Filter instances by a set of tags
	TagList []*Tag `json:"TagList,omitnil,omitempty" name:"TagList"`
}

type DescribeInstancesDetailRequest struct {
	*tchttp.BaseRequest
	
	// (Filter) filter by instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Filter by instance name, instance ID, AZ, VPC ID, or subnet ID. Fuzzy query is supported.
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// (Filter) instance status. 0: creating, 1: running, 2: deleting. If this parameter is left empty, all instances will be returned by default
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Offset. If this parameter is left empty, `0` will be used by default.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. If this parameter is left empty, `10` will be used by default. The maximum value is `20`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Tag key match.
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Filter. Valid values of `filter.Name` include `Ip`, `VpcId`, `SubNetId`, `InstanceType`, and `InstanceId`. Up to 10 values can be passed for `filter.Values`.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// This parameter has been deprecated and replaced with `InstanceIdList`.
	InstanceIds *string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Filter by instance ID.
	InstanceIdList []*string `json:"InstanceIdList,omitnil,omitempty" name:"InstanceIdList"`

	// Filter instances by a set of tags
	TagList []*Tag `json:"TagList,omitnil,omitempty" name:"TagList"`
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
	delete(f, "InstanceIds")
	delete(f, "InstanceIdList")
	delete(f, "TagList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesDetailResponseParams struct {
	// Returned result object of instance details
	Result *InstanceDetailResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstancesDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesDetailResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeInstancesRequestParams struct {
	// (Filter) filter by instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// (Filter) filter by instance name. Fuzzy search is supported
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// (Filter) instance status. 0: creating, 1: running, 2: deleting. If this parameter is left empty, all instances will be returned by default
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Offset. If this parameter is left empty, 0 will be used by default
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results to be returned. If this parameter is left empty, 10 will be used by default. The maximum value is 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Tag key value (this field has been deprecated).
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// VPC ID.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// (Filter) filter by instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// (Filter) filter by instance name. Fuzzy search is supported
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// (Filter) instance status. 0: creating, 1: running, 2: deleting. If this parameter is left empty, all instances will be returned by default
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Offset. If this parameter is left empty, 0 will be used by default
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results to be returned. If this parameter is left empty, 10 will be used by default. The maximum value is 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Tag key value (this field has been deprecated).
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// VPC ID.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`
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
	delete(f, "VpcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// Returned result
	Result *InstanceResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeRegionRequestParams struct {
	// The offset value
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The maximum number of results returned
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Business field, which can be ignored.
	Business *string `json:"Business,omitnil,omitempty" name:"Business"`

	// CDC business field, which can be ignored.
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`
}

type DescribeRegionRequest struct {
	*tchttp.BaseRequest
	
	// The offset value
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The maximum number of results returned
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Business field, which can be ignored.
	Business *string `json:"Business,omitnil,omitempty" name:"Business"`

	// CDC business field, which can be ignored.
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`
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
	delete(f, "CdcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRegionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegionResponseParams struct {
	// List of the returned results of enumerated regions
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	Result []*Region `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRegionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRegionResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeRouteRequestParams struct {
	// Unique instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Route ID
	RouteId *int64 `json:"RouteId,omitnil,omitempty" name:"RouteId"`
}

type DescribeRouteRequest struct {
	*tchttp.BaseRequest
	
	// Unique instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Route ID
	RouteId *int64 `json:"RouteId,omitnil,omitempty" name:"RouteId"`
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
	delete(f, "RouteId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRouteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRouteResponseParams struct {
	// Returned result set of route information
	Result *RouteResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRouteResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRouteResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeTaskStatusRequestParams struct {
	// Unique task ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`
}

type DescribeTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// Unique task ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`
}

func (r *DescribeTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskStatusResponseParams struct {
	// Returned result
	Result *TaskStatusResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskStatusResponseParams `json:"Response"`
}

func (r *DescribeTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicAttributesRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Topic name
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

type DescribeTopicAttributesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Topic name
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
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

// Predefined struct for user
type DescribeTopicAttributesResponseParams struct {
	// Returned result object
	Result *TopicAttributesResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTopicAttributesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopicAttributesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeTopicDetailRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// (Filter) filter by `topicName`. Fuzzy search is supported
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// Offset. If this parameter is left empty, 0 will be used by default
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results to be returned. If this parameter is left empty, 10 will be used by default. The maximum value is 20. This value must be greater than 0
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Name of the preset ACL rule.
	AclRuleName *string `json:"AclRuleName,omitnil,omitempty" name:"AclRuleName"`
}

type DescribeTopicDetailRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// (Filter) filter by `topicName`. Fuzzy search is supported
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// Offset. If this parameter is left empty, 0 will be used by default
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results to be returned. If this parameter is left empty, 10 will be used by default. The maximum value is 20. This value must be greater than 0
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Name of the preset ACL rule.
	AclRuleName *string `json:"AclRuleName,omitnil,omitempty" name:"AclRuleName"`
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

// Predefined struct for user
type DescribeTopicDetailResponseParams struct {
	// Returned entity of topic details
	Result *TopicDetailResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
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

// Predefined struct for user
type DescribeTopicProduceConnectionRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Topic name
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

type DescribeTopicProduceConnectionRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Topic name
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

func (r *DescribeTopicProduceConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicProduceConnectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "TopicName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopicProduceConnectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicProduceConnectionResponseParams struct {
	// Result set of returned connection information
	Result []*DescribeConnectInfoResultDTO `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTopicProduceConnectionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopicProduceConnectionResponseParams `json:"Response"`
}

func (r *DescribeTopicProduceConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicProduceConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Filter by `topicName`. Fuzzy search is supported
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// Offset. If this parameter is left empty, 0 will be used by default
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of results to be returned, which defaults to 20 if left empty. The maximum value is 50.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Name of the preset ACL rule.
	AclRuleName *string `json:"AclRuleName,omitnil,omitempty" name:"AclRuleName"`
}

type DescribeTopicRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Filter by `topicName`. Fuzzy search is supported
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// Offset. If this parameter is left empty, 0 will be used by default
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of results to be returned, which defaults to 20 if left empty. The maximum value is 50.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Name of the preset ACL rule.
	AclRuleName *string `json:"AclRuleName,omitnil,omitempty" name:"AclRuleName"`
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

// Predefined struct for user
type DescribeTopicResponseParams struct {
	// Returned result
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *TopicResult `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTopicResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopicResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeTopicSubscribeGroupRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Topic name
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Starting position of paging
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeTopicSubscribeGroupRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Topic name
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Starting position of paging
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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

// Predefined struct for user
type DescribeTopicSubscribeGroupResponseParams struct {
	// Returned results
	Result *TopicSubscribeGroup `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTopicSubscribeGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopicSubscribeGroupResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeTopicSyncReplicaRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Topic name
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Offset. If this parameter is left empty, 0 will be used by default.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results to be returned. If this parameter is left empty, 10 will be used by default. The maximum value is 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filters unsynced replicas only
	OutOfSyncReplicaOnly *bool `json:"OutOfSyncReplicaOnly,omitnil,omitempty" name:"OutOfSyncReplicaOnly"`
}

type DescribeTopicSyncReplicaRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Topic name
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Offset. If this parameter is left empty, 0 will be used by default.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results to be returned. If this parameter is left empty, 10 will be used by default. The maximum value is 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filters unsynced replicas only
	OutOfSyncReplicaOnly *bool `json:"OutOfSyncReplicaOnly,omitnil,omitempty" name:"OutOfSyncReplicaOnly"`
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

// Predefined struct for user
type DescribeTopicSyncReplicaResponseParams struct {
	// Returns topic replica details
	Result *TopicInSyncReplicaResult `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTopicSyncReplicaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopicSyncReplicaResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeUserRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Filter by name
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results to be returned in this request
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeUserRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Filter by name
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results to be returned in this request
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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

// Predefined struct for user
type DescribeUserResponseParams struct {
	// Returned result list
	Result *UserResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserResponseParams `json:"Response"`
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
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// Percentage of dynamic disk expansion each time.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	StepForwardPercentage *int64 `json:"StepForwardPercentage,omitnil,omitempty" name:"StepForwardPercentage"`

	// Disk quota threshold (in percentage) for triggering the automatic disk expansion event.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DiskQuotaPercentage *int64 `json:"DiskQuotaPercentage,omitnil,omitempty" name:"DiskQuotaPercentage"`

	// Max disk space in GB.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	MaxDiskSpace *int64 `json:"MaxDiskSpace,omitnil,omitempty" name:"MaxDiskSpace"`
}

type DynamicRetentionTime struct {
	// Whether the dynamic message retention time configuration is enabled. 0: disabled; 1: enabled
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// Disk quota threshold (in percentage) for triggering the message retention time change event
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	DiskQuotaPercentage *int64 `json:"DiskQuotaPercentage,omitnil,omitempty" name:"DiskQuotaPercentage"`

	// Percentage by which the message retention time is shortened each time
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	StepForwardPercentage *int64 `json:"StepForwardPercentage,omitnil,omitempty" name:"StepForwardPercentage"`

	// Minimum retention time, in minutes
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	BottomRetention *int64 `json:"BottomRetention,omitnil,omitempty" name:"BottomRetention"`
}

// Predefined struct for user
type FetchMessageByOffsetRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Topic name
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// Partition ID
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// Offset information, which is required.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type FetchMessageByOffsetRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Topic name
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// Partition ID
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// Offset information, which is required.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
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

// Predefined struct for user
type FetchMessageByOffsetResponseParams struct {
	// Returned results
	Result *ConsumerRecord `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type FetchMessageByOffsetResponse struct {
	*tchttp.BaseResponse
	Response *FetchMessageByOffsetResponseParams `json:"Response"`
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

// Predefined struct for user
type FetchMessageListByOffsetRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Topic name
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// Partition ID
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// Offset information
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The maximum number of messages that can be queried. Default value: 20. Maximum value: 20.
	SinglePartitionRecordNumber *int64 `json:"SinglePartitionRecordNumber,omitnil,omitempty" name:"SinglePartitionRecordNumber"`
}

type FetchMessageListByOffsetRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Topic name
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// Partition ID
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// Offset information
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The maximum number of messages that can be queried. Default value: 20. Maximum value: 20.
	SinglePartitionRecordNumber *int64 `json:"SinglePartitionRecordNumber,omitnil,omitempty" name:"SinglePartitionRecordNumber"`
}

func (r *FetchMessageListByOffsetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FetchMessageListByOffsetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Topic")
	delete(f, "Partition")
	delete(f, "Offset")
	delete(f, "SinglePartitionRecordNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FetchMessageListByOffsetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FetchMessageListByOffsetResponseParams struct {
	// Returned result. Note: The returned list does not display the message content (key and value). To query the message content, call the `FetchMessageByOffset` API.
	Result []*ConsumerRecord `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type FetchMessageListByOffsetResponse struct {
	*tchttp.BaseResponse
	Response *FetchMessageListByOffsetResponseParams `json:"Response"`
}

func (r *FetchMessageListByOffsetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FetchMessageListByOffsetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// Field to be filtered.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Filter value of field.
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type Group struct {
	// Group name
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`
}

type GroupInfoMember struct {
	// Unique ID generated for consumer in consumer group by coordinator
	MemberId *string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// `client.id` information by the client consumer SDK
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// Generally stores client IP address
	ClientHost *string `json:"ClientHost,omitnil,omitempty" name:"ClientHost"`

	// Stores the information of partition assigned to this consumer
	Assignment *Assignment `json:"Assignment,omitnil,omitempty" name:"Assignment"`
}

type GroupInfoResponse struct {
	// Error code. 0: success
	ErrorCode *string `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// Group status description (common valid values: Empty, Stable, Dead):
	// Dead: the consumer group does not exist
	// Empty: there are currently no consumer subscriptions in the consumer group
	// PreparingRebalance: the consumer group is currently in `rebalance` state
	// CompletingRebalance: the consumer group is currently in `rebalance` state
	// Stable: each consumer in the consumer group has joined and is in stable state
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// The type of protocol selected by the consumer group, which is `consumer` for common consumers. However, some systems use their own protocols; for example, the protocol used by kafka-connect is `connect`. Only with the standard `consumer` protocol can this API get to know the specific assigning method and parse the specific partition assignment
	ProtocolType *string `json:"ProtocolType,omitnil,omitempty" name:"ProtocolType"`

	// Consumer partition assignment algorithm, such as `range` (which is the default value for the Kafka consumer SDK), `roundrobin`, and `sticky`
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// This array contains information only if `state` is `Stable` and `protocol_type` is `consumer`
	Members []*GroupInfoMember `json:"Members,omitnil,omitempty" name:"Members"`

	// Kafka consumer group
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`
}

type GroupInfoTopics struct {
	// Name of assigned topics
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// Information of assigned partition
	// Note: this field may return null, indicating that no valid values can be obtained.
	Partitions []*int64 `json:"Partitions,omitnil,omitempty" name:"Partitions"`
}

type GroupOffsetPartition struct {
	// Topic `partitionId`
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// Offset position submitted by consumer
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Metadata can be passed in for other purposes when the consumer submits messages. Currently, this parameter is usually an empty string
	// Note: this field may return null, indicating that no valid values can be obtained.
	Metadata *string `json:"Metadata,omitnil,omitempty" name:"Metadata"`

	// Error code
	ErrorCode *int64 `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// Latest offset of current partition
	LogEndOffset *int64 `json:"LogEndOffset,omitnil,omitempty" name:"LogEndOffset"`

	// Number of unconsumed messages
	Lag *int64 `json:"Lag,omitnil,omitempty" name:"Lag"`
}

type GroupOffsetResponse struct {
	// Total number of eligible results
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Array of partitions in the topic, where each element is a JSON object
	// Note: this field may return null, indicating that no valid values can be obtained.
	TopicList []*GroupOffsetTopic `json:"TopicList,omitnil,omitempty" name:"TopicList"`
}

type GroupOffsetTopic struct {
	// Topic name
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// Array of partitions in the topic, where each element is a JSON object
	// Note: this field may return null, indicating that no valid values can be obtained.
	Partitions []*GroupOffsetPartition `json:"Partitions,omitnil,omitempty" name:"Partitions"`
}

type GroupResponse struct {
	// Count
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// GroupList
	// Note: this field may return null, indicating that no valid values can be obtained.
	GroupList []*DescribeGroup `json:"GroupList,omitnil,omitempty" name:"GroupList"`

	// Consumer group quota
	// Note: This field may return null, indicating that no valid values can be obtained.
	GroupCountQuota *uint64 `json:"GroupCountQuota,omitnil,omitempty" name:"GroupCountQuota"`
}

// Predefined struct for user
type InquireCkafkaPriceRequestParams struct {
	// `standard`: Standard Edition; `profession`: Pro Edition
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Billing mode for instance purchase/renewal. If this parameter is left empty when you purchase an instance, the fees for one month under the monthly subscription mode will be displayed by default.
	InstanceChargeParam *InstanceChargeParam `json:"InstanceChargeParam,omitnil,omitempty" name:"InstanceChargeParam"`

	// The number of instances to be purchased or renewed. If this parameter is left empty, the default value is `1`.
	InstanceNum *int64 `json:"InstanceNum,omitnil,omitempty" name:"InstanceNum"`

	// Private network bandwidth in MB/sec, which is required when you purchase an instance.
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// Disk type and size, which is required when you purchase an instance.
	InquiryDiskParam *InquiryDiskParam `json:"InquiryDiskParam,omitnil,omitempty" name:"InquiryDiskParam"`

	// Message retention period in hours, which is required when you purchase an instance.
	MessageRetention *int64 `json:"MessageRetention,omitnil,omitempty" name:"MessageRetention"`

	// The number of instance topics to be purchased, which is required when you purchase an instance.
	Topic *int64 `json:"Topic,omitnil,omitempty" name:"Topic"`

	// The number of instance partitions to be purchased, which is required when you purchase an instance.
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// The region for instance purchase, which can be obtained via the `DescribeCkafkaZone` API.
	ZoneIds []*int64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// Operation type flag. `purchase`: Making new purchases; `renew`: Renewing an instance. The default value is `purchase` if this parameter is left empty.
	CategoryAction *string `json:"CategoryAction,omitnil,omitempty" name:"CategoryAction"`

	// This field is not required.
	BillType *string `json:"BillType,omitnil,omitempty" name:"BillType"`

	// Billing mode for public network bandwidth, which is required when you purchase public network bandwidth. Currently, public network bandwidth is only supported for Pro Edition.
	PublicNetworkParam *InquiryPublicNetworkParam `json:"PublicNetworkParam,omitnil,omitempty" name:"PublicNetworkParam"`

	// ID of the instance to be renewed, which is required when you renew an instance.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type InquireCkafkaPriceRequest struct {
	*tchttp.BaseRequest
	
	// `standard`: Standard Edition; `profession`: Pro Edition
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Billing mode for instance purchase/renewal. If this parameter is left empty when you purchase an instance, the fees for one month under the monthly subscription mode will be displayed by default.
	InstanceChargeParam *InstanceChargeParam `json:"InstanceChargeParam,omitnil,omitempty" name:"InstanceChargeParam"`

	// The number of instances to be purchased or renewed. If this parameter is left empty, the default value is `1`.
	InstanceNum *int64 `json:"InstanceNum,omitnil,omitempty" name:"InstanceNum"`

	// Private network bandwidth in MB/sec, which is required when you purchase an instance.
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// Disk type and size, which is required when you purchase an instance.
	InquiryDiskParam *InquiryDiskParam `json:"InquiryDiskParam,omitnil,omitempty" name:"InquiryDiskParam"`

	// Message retention period in hours, which is required when you purchase an instance.
	MessageRetention *int64 `json:"MessageRetention,omitnil,omitempty" name:"MessageRetention"`

	// The number of instance topics to be purchased, which is required when you purchase an instance.
	Topic *int64 `json:"Topic,omitnil,omitempty" name:"Topic"`

	// The number of instance partitions to be purchased, which is required when you purchase an instance.
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// The region for instance purchase, which can be obtained via the `DescribeCkafkaZone` API.
	ZoneIds []*int64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// Operation type flag. `purchase`: Making new purchases; `renew`: Renewing an instance. The default value is `purchase` if this parameter is left empty.
	CategoryAction *string `json:"CategoryAction,omitnil,omitempty" name:"CategoryAction"`

	// This field is not required.
	BillType *string `json:"BillType,omitnil,omitempty" name:"BillType"`

	// Billing mode for public network bandwidth, which is required when you purchase public network bandwidth. Currently, public network bandwidth is only supported for Pro Edition.
	PublicNetworkParam *InquiryPublicNetworkParam `json:"PublicNetworkParam,omitnil,omitempty" name:"PublicNetworkParam"`

	// ID of the instance to be renewed, which is required when you renew an instance.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *InquireCkafkaPriceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquireCkafkaPriceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceType")
	delete(f, "InstanceChargeParam")
	delete(f, "InstanceNum")
	delete(f, "Bandwidth")
	delete(f, "InquiryDiskParam")
	delete(f, "MessageRetention")
	delete(f, "Topic")
	delete(f, "Partition")
	delete(f, "ZoneIds")
	delete(f, "CategoryAction")
	delete(f, "BillType")
	delete(f, "PublicNetworkParam")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquireCkafkaPriceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type InquireCkafkaPriceResp struct {
	// Instance price
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstancePrice *InquiryPrice `json:"InstancePrice,omitnil,omitempty" name:"InstancePrice"`

	// Public network bandwidth price
	// Note: This field may return null, indicating that no valid values can be obtained.
	PublicNetworkBandwidthPrice *InquiryPrice `json:"PublicNetworkBandwidthPrice,omitnil,omitempty" name:"PublicNetworkBandwidthPrice"`
}

// Predefined struct for user
type InquireCkafkaPriceResponseParams struct {
	// Output parameters
	Result *InquireCkafkaPriceResp `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquireCkafkaPriceResponse struct {
	*tchttp.BaseResponse
	Response *InquireCkafkaPriceResponseParams `json:"Response"`
}

func (r *InquireCkafkaPriceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquireCkafkaPriceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryBasePrice struct {
	// Original unit price
	// Note: This field may return null, indicating that no valid values can be obtained.
	UnitPrice *float64 `json:"UnitPrice,omitnil,omitempty" name:"UnitPrice"`

	// Discounted unit price
	// Note: This field may return null, indicating that no valid values can be obtained.
	UnitPriceDiscount *float64 `json:"UnitPriceDiscount,omitnil,omitempty" name:"UnitPriceDiscount"`

	// Original price in total
	// Note: This field may return null, indicating that no valid values can be obtained.
	OriginalPrice *float64 `json:"OriginalPrice,omitnil,omitempty" name:"OriginalPrice"`

	// Discounted price in total
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiscountPrice *float64 `json:"DiscountPrice,omitnil,omitempty" name:"DiscountPrice"`

	// Discount (%)
	// Note: This field may return null, indicating that no valid values can be obtained.
	Discount *float64 `json:"Discount,omitnil,omitempty" name:"Discount"`

	// Number of purchased items
	// Note: This field may return null, indicating that no valid values can be obtained.
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// Currency for payment
	// Note: This field may return null, indicating that no valid values can be obtained.
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`

	// Dedicated disk response parameter
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// Validity period
	// Note: This field may return null, indicating that no valid values can be obtained.
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// Unit of the validity period (`m`: Month; `h`: Hour)
	// Note: This field may return null, indicating that no valid values can be obtained.
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// Purchase quantity
	// Note: This field may return null, indicating that no valid values can be obtained.
	Value *int64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type InquiryDetailPrice struct {
	// Price of additional private network bandwidth
	// Note: This field may return null, indicating that no valid values can be obtained.
	BandwidthPrice *InquiryBasePrice `json:"BandwidthPrice,omitnil,omitempty" name:"BandwidthPrice"`

	// Disk price
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskPrice *InquiryBasePrice `json:"DiskPrice,omitnil,omitempty" name:"DiskPrice"`

	// Price of additional partitions
	// Note: This field may return null, indicating that no valid values can be obtained.
	PartitionPrice *InquiryBasePrice `json:"PartitionPrice,omitnil,omitempty" name:"PartitionPrice"`

	// Price of additional topics
	// Note: This field may return null, indicating that no valid values can be obtained.
	TopicPrice *InquiryBasePrice `json:"TopicPrice,omitnil,omitempty" name:"TopicPrice"`

	// Instance package price
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceTypePrice *InquiryBasePrice `json:"InstanceTypePrice,omitnil,omitempty" name:"InstanceTypePrice"`
}

type InquiryDiskParam struct {
	// Disk type. Valid values: `SSD` (SSD), `CLOUD_SSD` (SSD cloud disk), `CLOUD_PREMIUM` (Premium cloud disk), `CLOUD_BASIC` (Cloud disk).
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// Size of the purchased disk in GB
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`
}

type InquiryPrice struct {
	// Original unit price
	// Note: This field may return null, indicating that no valid values can be obtained.
	UnitPrice *float64 `json:"UnitPrice,omitnil,omitempty" name:"UnitPrice"`

	// Discounted unit price
	// Note: This field may return null, indicating that no valid values can be obtained.
	UnitPriceDiscount *float64 `json:"UnitPriceDiscount,omitnil,omitempty" name:"UnitPriceDiscount"`

	// Original price in total
	// Note: This field may return null, indicating that no valid values can be obtained.
	OriginalPrice *float64 `json:"OriginalPrice,omitnil,omitempty" name:"OriginalPrice"`

	// Discounted price in total
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiscountPrice *float64 `json:"DiscountPrice,omitnil,omitempty" name:"DiscountPrice"`

	// Discount (%)
	// Note: This field may return null, indicating that no valid values can be obtained.
	Discount *float64 `json:"Discount,omitnil,omitempty" name:"Discount"`

	// Number of purchased items
	// Note: This field may return null, indicating that no valid values can be obtained.
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// Currency for payment
	// Note: This field may return null, indicating that no valid values can be obtained.
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`

	// Dedicated disk response parameter
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// Validity period
	// Note: This field may return null, indicating that no valid values can be obtained.
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// Unit of the validity period (`m`: Month; `h`: Hour)
	// Note: This field may return null, indicating that no valid values can be obtained.
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// Purchase quantity
	// Note: This field may return null, indicating that no valid values can be obtained.
	Value *int64 `json:"Value,omitnil,omitempty" name:"Value"`

	// Prices of different purchased items
	// Note: This field may return null, indicating that no valid values can be obtained.
	DetailPrices *InquiryDetailPrice `json:"DetailPrices,omitnil,omitempty" name:"DetailPrices"`
}

type InquiryPublicNetworkParam struct {
	// Public network bandwidth billing mode (`BANDWIDTH_PREPAID`: Monthly subscription; `BANDWIDTH_POSTPAID_BY_HOUR`: Bill-by-hour)
	PublicNetworkChargeType *string `json:"PublicNetworkChargeType,omitnil,omitempty" name:"PublicNetworkChargeType"`

	// Public network bandwidth in MB
	PublicNetworkMonthly *int64 `json:"PublicNetworkMonthly,omitnil,omitempty" name:"PublicNetworkMonthly"`
}

type Instance struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Instance status. 0: creating, 1: running, 2: deleting, 5: isolated, -1: creation failed
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Whether it is an open-source instance. true: yes, false: no
	// Note: this field may return null, indicating that no valid values can be obtained.
	IfCommunity *bool `json:"IfCommunity,omitnil,omitempty" name:"IfCommunity"`
}

type InstanceAttributesResponse struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// VIP list information of access point
	VipList []*VipEntity `json:"VipList,omitnil,omitempty" name:"VipList"`

	// Virtual IP
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Virtual port
	Vport *string `json:"Vport,omitnil,omitempty" name:"Vport"`

	// Instance status. 0: creating, 1: running, 2: deleting
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Instance bandwidth in Mbps
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// Instance storage capacity in GB
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// AZ
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// VPC ID. If this parameter is empty, it means the basic network
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID. If this parameter is empty, it means the basic network
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Instance health status. 1: healthy, 2: alarmed, 3: exceptional
	Healthy *int64 `json:"Healthy,omitnil,omitempty" name:"Healthy"`

	// Instance health information. Currently, the disk utilization is displayed with a maximum length of 256
	HealthyMessage *string `json:"HealthyMessage,omitnil,omitempty" name:"HealthyMessage"`

	// Creation time
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Message retention period in minutes
	MsgRetentionTime *int64 `json:"MsgRetentionTime,omitnil,omitempty" name:"MsgRetentionTime"`

	// Configuration for automatic topic creation. If this field is empty, it means that automatic creation is not enabled
	Config *InstanceConfigDO `json:"Config,omitnil,omitempty" name:"Config"`

	// Number of remaining creatable partitions
	RemainderPartitions *int64 `json:"RemainderPartitions,omitnil,omitempty" name:"RemainderPartitions"`

	// Number of remaining creatable topics
	RemainderTopics *int64 `json:"RemainderTopics,omitnil,omitempty" name:"RemainderTopics"`

	// Number of partitions already created
	CreatedPartitions *int64 `json:"CreatedPartitions,omitnil,omitempty" name:"CreatedPartitions"`

	// Number of topics already created
	CreatedTopics *int64 `json:"CreatedTopics,omitnil,omitempty" name:"CreatedTopics"`

	// Tag array
	// Note: this field may return null, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Expiration time
	// Note: this field may return null, indicating that no valid values can be obtained.
	ExpireTime *uint64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// Cross-AZ
	// Note: this field may return null, indicating that no valid values can be obtained.
	ZoneIds []*int64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// Kafka version information
	// Note: this field may return null, indicating that no valid values can be obtained.
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// Maximum number of groups
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxGroupNum *int64 `json:"MaxGroupNum,omitnil,omitempty" name:"MaxGroupNum"`

	// Offering type. `0`: Standard Edition; `1`: Professional Edition
	// Note: this field may return `null`, indicating that no valid value was found.
	Cvm *int64 `json:"Cvm,omitnil,omitempty" name:"Cvm"`

	// Type.
	// Note: this field may return `null`, indicating that no valid value was found.
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Features supported by the instance. `FEATURE_SUBNET_ACL` indicates that the ACL policy supports setting subnets. 
	// Note: this field may return null, indicating that no valid values can be obtained.
	Features []*string `json:"Features,omitnil,omitempty" name:"Features"`

	// Dynamic message retention policy
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	RetentionTimeConfig *DynamicRetentionTime `json:"RetentionTimeConfig,omitnil,omitempty" name:"RetentionTimeConfig"`

	// Maximum number of connections
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxConnection *uint64 `json:"MaxConnection,omitnil,omitempty" name:"MaxConnection"`

	// Public network bandwidth
	// Note: this field may return null, indicating that no valid values can be obtained.
	PublicNetwork *int64 `json:"PublicNetwork,omitnil,omitempty" name:"PublicNetwork"`

	// Time
	// Note: this field may return null, indicating that no valid values can be obtained.
	DeleteRouteTimestamp *string `json:"DeleteRouteTimestamp,omitnil,omitempty" name:"DeleteRouteTimestamp"`

	// Number of remaining creatable partitions
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	RemainingPartitions *int64 `json:"RemainingPartitions,omitnil,omitempty" name:"RemainingPartitions"`

	// Number of remaining creatable topics
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	RemainingTopics *int64 `json:"RemainingTopics,omitnil,omitempty" name:"RemainingTopics"`

	// Dynamic disk expansion policy.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DynamicDiskConfig *DynamicDiskConfig `json:"DynamicDiskConfig,omitnil,omitempty" name:"DynamicDiskConfig"`
}

type InstanceChargeParam struct {
	// Instance billing mode (`PREPAID`: Monthly subscription; `POSTPAID_BY_HOUR`: Pay-as-you-go)
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// Validity period, which is only required for the monthly subscription billing mode
	InstanceChargePeriod *int64 `json:"InstanceChargePeriod,omitnil,omitempty" name:"InstanceChargePeriod"`
}

type InstanceConfigDO struct {
	// Whether to create topics automatically
	AutoCreateTopicsEnable *bool `json:"AutoCreateTopicsEnable,omitnil,omitempty" name:"AutoCreateTopicsEnable"`

	// Number of partitions
	DefaultNumPartitions *int64 `json:"DefaultNumPartitions,omitnil,omitempty" name:"DefaultNumPartitions"`

	// Default replication factor
	DefaultReplicationFactor *int64 `json:"DefaultReplicationFactor,omitnil,omitempty" name:"DefaultReplicationFactor"`
}

type InstanceDetail struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Instance VIP information
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Instance port information
	Vport *string `json:"Vport,omitnil,omitempty" name:"Vport"`

	// Virtual IP list
	VipList []*VipEntity `json:"VipList,omitnil,omitempty" name:"VipList"`

	// Instance status. 0: creating, 1: running, 2: deleting, 5: isolated, -1: creation failed
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Instance bandwidth in Mbps
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// Instance storage capacity in GB
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// AZ ID
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// vpcId. If this parameter is empty, it means the basic network
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Whether to renew the instance automatically, which is an int-type enumerated value. 1: yes, 2: no
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// Instance status. An int-type value will be returned. `0`: Healthy, `1`: Alarmed, `2`: Exceptional
	Healthy *int64 `json:"Healthy,omitnil,omitempty" name:"Healthy"`

	// Instance status information
	HealthyMessage *string `json:"HealthyMessage,omitnil,omitempty" name:"HealthyMessage"`

	// Instance creation time
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Instance expiration time
	ExpireTime *int64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// Whether it is an internal customer. 1: yes
	IsInternal *int64 `json:"IsInternal,omitnil,omitempty" name:"IsInternal"`

	// Number of topics
	TopicNum *int64 `json:"TopicNum,omitnil,omitempty" name:"TopicNum"`

	// Tag
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Kafka version information
	// Note: this field may return null, indicating that no valid values can be obtained.
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// Cross-AZ
	// Note: this field may return null, indicating that no valid values can be obtained.
	ZoneIds []*int64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// CKafka sale type
	// Note: this field may return null, indicating that no valid values can be obtained.
	Cvm *int64 `json:"Cvm,omitnil,omitempty" name:"Cvm"`

	// CKafka instance type
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Disk type
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// Maximum number of topics for the current instance
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	MaxTopicNumber *int64 `json:"MaxTopicNumber,omitnil,omitempty" name:"MaxTopicNumber"`

	// Maximum number of partitions for the current instance
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	MaxPartitionNumber *int64 `json:"MaxPartitionNumber,omitnil,omitempty" name:"MaxPartitionNumber"`

	// Time of scheduled upgrade
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	RebalanceTime *string `json:"RebalanceTime,omitnil,omitempty" name:"RebalanceTime"`

	// Number of partitions in the current instance.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	PartitionNumber *uint64 `json:"PartitionNumber,omitnil,omitempty" name:"PartitionNumber"`

	// Public network bandwidth type.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	PublicNetworkChargeType *string `json:"PublicNetworkChargeType,omitnil,omitempty" name:"PublicNetworkChargeType"`

	// Public network bandwidth.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	PublicNetwork *int64 `json:"PublicNetwork,omitnil,omitempty" name:"PublicNetwork"`

	// Instance type.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// Instance feature list.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Features []*string `json:"Features,omitnil,omitempty" name:"Features"`
}

type InstanceDetailResponse struct {
	// Total number of eligible instances
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of eligible instance details
	InstanceList []*InstanceDetail `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`
}

type InstanceQuotaConfigResp struct {
	// Production throttling in MB/sec.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	QuotaProducerByteRate *int64 `json:"QuotaProducerByteRate,omitnil,omitempty" name:"QuotaProducerByteRate"`

	// Consumption throttling in MB/sec.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	QuotaConsumerByteRate *int64 `json:"QuotaConsumerByteRate,omitnil,omitempty" name:"QuotaConsumerByteRate"`
}

type InstanceResponse struct {
	// List of eligible instances
	// Note: this field may return null, indicating that no valid values can be obtained.
	InstanceList []*Instance `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// Total number of eligible results
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type JgwOperateResponse struct {
	// Returned code. 0: normal, other values: error
	ReturnCode *string `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// Success message
	ReturnMessage *string `json:"ReturnMessage,omitnil,omitempty" name:"ReturnMessage"`

	// Data returned by an operation, which may contain `flowId`, etc.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Data *OperateResponseData `json:"Data,omitnil,omitempty" name:"Data"`
}

// Predefined struct for user
type ModifyAclRuleRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ACL policy name
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// Whether to be applied to new topics
	IsApplied *int64 `json:"IsApplied,omitnil,omitempty" name:"IsApplied"`
}

type ModifyAclRuleRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ACL policy name
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// Whether to be applied to new topics
	IsApplied *int64 `json:"IsApplied,omitnil,omitempty" name:"IsApplied"`
}

func (r *ModifyAclRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAclRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "RuleName")
	delete(f, "IsApplied")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAclRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAclRuleResponseParams struct {
	// Unique key of a rule
	Result *int64 `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAclRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAclRuleResponseParams `json:"Response"`
}

func (r *ModifyAclRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAclRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDatahubTopicRequestParams struct {
	// Name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Message retention period in ms. The current minimum value is 60,000 ms.
	RetentionMs *int64 `json:"RetentionMs,omitnil,omitempty" name:"RetentionMs"`

	// Topic remarks, which are a string of up to 64 characters. It can contain letters, digits, and hyphens (-) and must start with a letter.
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// Tag list
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type ModifyDatahubTopicRequest struct {
	*tchttp.BaseRequest
	
	// Name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Message retention period in ms. The current minimum value is 60,000 ms.
	RetentionMs *int64 `json:"RetentionMs,omitnil,omitempty" name:"RetentionMs"`

	// Topic remarks, which are a string of up to 64 characters. It can contain letters, digits, and hyphens (-) and must start with a letter.
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// Tag list
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *ModifyDatahubTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDatahubTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "RetentionMs")
	delete(f, "Note")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDatahubTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDatahubTopicResponseParams struct {
	// Returned result set
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDatahubTopicResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDatahubTopicResponseParams `json:"Response"`
}

func (r *ModifyDatahubTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDatahubTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGroupOffsetsRequestParams struct {
	// Kafka instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Kafka consumer group
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// Offset resetting policy. Meanings of the input parameters: 0: equivalent to the `shift-by` parameter, which indicates to shift the offset forward or backward by the value of the `shift`. 1: equivalent to `by-duration`, `to-datetime`, `to-earliest`, or `to-latest`, which indicates to move the offset to the specified timestamp. 2: equivalent to `to-offset`, which indicates to move the offset to the specified offset position
	Strategy *int64 `json:"Strategy,omitnil,omitempty" name:"Strategy"`

	// Indicates the topics to be reset. If this parameter is left empty, all topics will be reset
	Topics []*string `json:"Topics,omitnil,omitempty" name:"Topics"`

	// When `strategy` is 0, this field is required. If it is above zero, the offset will be shifted backward by the value of the `shift`. If it is below zero, the offset will be shifted forward by the value of the `shift`. After a correct reset, the new offset should be (old_offset + shift). Note that if the new offset is smaller than the `earliest` parameter of the partition, it will be set to `earliest`, and if it is greater than the `latest` parameter of the partition, it will be set to `latest`
	Shift *int64 `json:"Shift,omitnil,omitempty" name:"Shift"`

	// Unit: ms. When `strategy` is 1, this field is required, where -2 indicates to reset the offset to the initial position, -1 indicates to reset to the latest position (equivalent to emptying), and other values represent the specified time, i.e., the offset of the topic at the specified time will be obtained and then reset. Note that if there is no message at the specified time, the last offset will be obtained
	ShiftTimestamp *int64 `json:"ShiftTimestamp,omitnil,omitempty" name:"ShiftTimestamp"`

	// Position of the offset that needs to be reset. When `strategy` is 2, this field is required
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// List of partitions that need to be reset. If the topics parameter is not specified, reset partitions in the corresponding partition list of all topics. If the topics parameter is specified, reset partitions of the corresponding partition list of the specified topic list.
	Partitions []*int64 `json:"Partitions,omitnil,omitempty" name:"Partitions"`
}

type ModifyGroupOffsetsRequest struct {
	*tchttp.BaseRequest
	
	// Kafka instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Kafka consumer group
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// Offset resetting policy. Meanings of the input parameters: 0: equivalent to the `shift-by` parameter, which indicates to shift the offset forward or backward by the value of the `shift`. 1: equivalent to `by-duration`, `to-datetime`, `to-earliest`, or `to-latest`, which indicates to move the offset to the specified timestamp. 2: equivalent to `to-offset`, which indicates to move the offset to the specified offset position
	Strategy *int64 `json:"Strategy,omitnil,omitempty" name:"Strategy"`

	// Indicates the topics to be reset. If this parameter is left empty, all topics will be reset
	Topics []*string `json:"Topics,omitnil,omitempty" name:"Topics"`

	// When `strategy` is 0, this field is required. If it is above zero, the offset will be shifted backward by the value of the `shift`. If it is below zero, the offset will be shifted forward by the value of the `shift`. After a correct reset, the new offset should be (old_offset + shift). Note that if the new offset is smaller than the `earliest` parameter of the partition, it will be set to `earliest`, and if it is greater than the `latest` parameter of the partition, it will be set to `latest`
	Shift *int64 `json:"Shift,omitnil,omitempty" name:"Shift"`

	// Unit: ms. When `strategy` is 1, this field is required, where -2 indicates to reset the offset to the initial position, -1 indicates to reset to the latest position (equivalent to emptying), and other values represent the specified time, i.e., the offset of the topic at the specified time will be obtained and then reset. Note that if there is no message at the specified time, the last offset will be obtained
	ShiftTimestamp *int64 `json:"ShiftTimestamp,omitnil,omitempty" name:"ShiftTimestamp"`

	// Position of the offset that needs to be reset. When `strategy` is 2, this field is required
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// List of partitions that need to be reset. If the topics parameter is not specified, reset partitions in the corresponding partition list of all topics. If the topics parameter is specified, reset partitions of the corresponding partition list of the specified topic list.
	Partitions []*int64 `json:"Partitions,omitnil,omitempty" name:"Partitions"`
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

// Predefined struct for user
type ModifyGroupOffsetsResponseParams struct {
	// Returned result
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyGroupOffsetsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyGroupOffsetsResponseParams `json:"Response"`
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
	AutoCreateTopicEnable *bool `json:"AutoCreateTopicEnable,omitnil,omitempty" name:"AutoCreateTopicEnable"`

	// Optional. If `auto.create.topic.enable` is set to `true` and this value is not set, 3 will be used by default
	DefaultNumPartitions *int64 `json:"DefaultNumPartitions,omitnil,omitempty" name:"DefaultNumPartitions"`

	// If `auto.create.topic.enable` is set to `true` but this value is not set, 2 will be used by default
	DefaultReplicationFactor *int64 `json:"DefaultReplicationFactor,omitnil,omitempty" name:"DefaultReplicationFactor"`
}

// Predefined struct for user
type ModifyInstanceAttributesRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Maximum retention period in minutes for instance log, which can be up to 30 days. 0 indicates not to enable the log retention period policy
	MsgRetentionTime *int64 `json:"MsgRetentionTime,omitnil,omitempty" name:"MsgRetentionTime"`

	// Instance name string of up to 64 characters, which must begin with a letter and can contain letters, digits, and dashes (`-`)
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Instance configuration
	Config *ModifyInstanceAttributesConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// Dynamic message retention policy configuration
	DynamicRetentionConfig *DynamicRetentionTime `json:"DynamicRetentionConfig,omitnil,omitempty" name:"DynamicRetentionConfig"`

	// Modification of the rebalancing time after upgrade
	RebalanceTime *int64 `json:"RebalanceTime,omitnil,omitempty" name:"RebalanceTime"`

	// Public network bandwidth
	PublicNetwork *int64 `json:"PublicNetwork,omitnil,omitempty" name:"PublicNetwork"`

	// Dynamic disk expansion policy configuration.
	DynamicDiskConfig *DynamicDiskConfig `json:"DynamicDiskConfig,omitnil,omitempty" name:"DynamicDiskConfig"`

	// The size of a single message in bytes at the instance level.
	MaxMessageByte *uint64 `json:"MaxMessageByte,omitnil,omitempty" name:"MaxMessageByte"`
}

type ModifyInstanceAttributesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Maximum retention period in minutes for instance log, which can be up to 30 days. 0 indicates not to enable the log retention period policy
	MsgRetentionTime *int64 `json:"MsgRetentionTime,omitnil,omitempty" name:"MsgRetentionTime"`

	// Instance name string of up to 64 characters, which must begin with a letter and can contain letters, digits, and dashes (`-`)
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Instance configuration
	Config *ModifyInstanceAttributesConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// Dynamic message retention policy configuration
	DynamicRetentionConfig *DynamicRetentionTime `json:"DynamicRetentionConfig,omitnil,omitempty" name:"DynamicRetentionConfig"`

	// Modification of the rebalancing time after upgrade
	RebalanceTime *int64 `json:"RebalanceTime,omitnil,omitempty" name:"RebalanceTime"`

	// Public network bandwidth
	PublicNetwork *int64 `json:"PublicNetwork,omitnil,omitempty" name:"PublicNetwork"`

	// Dynamic disk expansion policy configuration.
	DynamicDiskConfig *DynamicDiskConfig `json:"DynamicDiskConfig,omitnil,omitempty" name:"DynamicDiskConfig"`

	// The size of a single message in bytes at the instance level.
	MaxMessageByte *uint64 `json:"MaxMessageByte,omitnil,omitempty" name:"MaxMessageByte"`
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
	delete(f, "MaxMessageByte")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceAttributesResponseParams struct {
	// Returned result
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceAttributesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceAttributesResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyInstancePreRequestParams struct {
	// Instance name.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Estimated disk capacity, which can be increased by increment.
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// Estimated bandwidth, which can be increased by increment.
	BandWidth *int64 `json:"BandWidth,omitnil,omitempty" name:"BandWidth"`

	// Estimated partition count, which can be increased by increment.
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`
}

type ModifyInstancePreRequest struct {
	*tchttp.BaseRequest
	
	// Instance name.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Estimated disk capacity, which can be increased by increment.
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// Estimated bandwidth, which can be increased by increment.
	BandWidth *int64 `json:"BandWidth,omitnil,omitempty" name:"BandWidth"`

	// Estimated partition count, which can be increased by increment.
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`
}

func (r *ModifyInstancePreRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancePreRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DiskSize")
	delete(f, "BandWidth")
	delete(f, "Partition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstancePreRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstancePreResponseParams struct {
	// Response structure of modifying the configurations of a prepaid instance.
	Result *CreateInstancePreResp `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstancePreResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstancePreResponseParams `json:"Response"`
}

func (r *ModifyInstancePreResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancePreResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPasswordRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Username
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Current user password
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// New user password
	PasswordNew *string `json:"PasswordNew,omitnil,omitempty" name:"PasswordNew"`
}

type ModifyPasswordRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Username
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Current user password
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// New user password
	PasswordNew *string `json:"PasswordNew,omitnil,omitempty" name:"PasswordNew"`
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

// Predefined struct for user
type ModifyPasswordResponseParams struct {
	// Returned result
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyPasswordResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPasswordResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyTopicAttributesRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Topic name.
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Topic remarks string of up to 64 characters, which must begin with a letter and can contain letters, digits, and dashes (`-`).
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// IP allowlist switch. 1: enabled, 0: disabled.
	EnableWhiteList *int64 `json:"EnableWhiteList,omitnil,omitempty" name:"EnableWhiteList"`

	// Default value: 1.
	MinInsyncReplicas *int64 `json:"MinInsyncReplicas,omitnil,omitempty" name:"MinInsyncReplicas"`

	// 0: false, 1: true. Default value: 0.
	UncleanLeaderElectionEnable *int64 `json:"UncleanLeaderElectionEnable,omitnil,omitempty" name:"UncleanLeaderElectionEnable"`

	// Message retention period in ms. The current minimum value is 60,000 ms.
	RetentionMs *int64 `json:"RetentionMs,omitnil,omitempty" name:"RetentionMs"`

	// Segment rolling duration in ms. The current minimum value is 86,400,000 ms.
	SegmentMs *int64 `json:"SegmentMs,omitnil,omitempty" name:"SegmentMs"`

	// Max message size in bytes. Max value: 8,388,608 bytes (8 MB).
	MaxMessageBytes *int64 `json:"MaxMessageBytes,omitnil,omitempty" name:"MaxMessageBytes"`

	// Message deletion policy. Valid values: delete, compact
	CleanUpPolicy *string `json:"CleanUpPolicy,omitnil,omitempty" name:"CleanUpPolicy"`

	// IP allowlist, which is required if the value of `enableWhileList` is 1.
	IpWhiteList []*string `json:"IpWhiteList,omitnil,omitempty" name:"IpWhiteList"`

	// Preset ACL rule. `1`: enable, `0`: disable. Default value: `0`.
	EnableAclRule *int64 `json:"EnableAclRule,omitnil,omitempty" name:"EnableAclRule"`

	// Name of the preset ACL rule.
	AclRuleName *string `json:"AclRuleName,omitnil,omitempty" name:"AclRuleName"`

	// Message retention file size in bytes, which is an optional parameter. Default value: -1. Currently, the min value that can be entered is 1,048,576 B.
	RetentionBytes *int64 `json:"RetentionBytes,omitnil,omitempty" name:"RetentionBytes"`

	// Tag list.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Production throttling in MB/sec.
	QuotaProducerByteRate *int64 `json:"QuotaProducerByteRate,omitnil,omitempty" name:"QuotaProducerByteRate"`

	// Consumption throttling in MB/sec.
	QuotaConsumerByteRate *int64 `json:"QuotaConsumerByteRate,omitnil,omitempty" name:"QuotaConsumerByteRate"`

	// The number of topic replicas.
	ReplicaNum *int64 `json:"ReplicaNum,omitnil,omitempty" name:"ReplicaNum"`
}

type ModifyTopicAttributesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Topic name.
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Topic remarks string of up to 64 characters, which must begin with a letter and can contain letters, digits, and dashes (`-`).
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// IP allowlist switch. 1: enabled, 0: disabled.
	EnableWhiteList *int64 `json:"EnableWhiteList,omitnil,omitempty" name:"EnableWhiteList"`

	// Default value: 1.
	MinInsyncReplicas *int64 `json:"MinInsyncReplicas,omitnil,omitempty" name:"MinInsyncReplicas"`

	// 0: false, 1: true. Default value: 0.
	UncleanLeaderElectionEnable *int64 `json:"UncleanLeaderElectionEnable,omitnil,omitempty" name:"UncleanLeaderElectionEnable"`

	// Message retention period in ms. The current minimum value is 60,000 ms.
	RetentionMs *int64 `json:"RetentionMs,omitnil,omitempty" name:"RetentionMs"`

	// Segment rolling duration in ms. The current minimum value is 86,400,000 ms.
	SegmentMs *int64 `json:"SegmentMs,omitnil,omitempty" name:"SegmentMs"`

	// Max message size in bytes. Max value: 8,388,608 bytes (8 MB).
	MaxMessageBytes *int64 `json:"MaxMessageBytes,omitnil,omitempty" name:"MaxMessageBytes"`

	// Message deletion policy. Valid values: delete, compact
	CleanUpPolicy *string `json:"CleanUpPolicy,omitnil,omitempty" name:"CleanUpPolicy"`

	// IP allowlist, which is required if the value of `enableWhileList` is 1.
	IpWhiteList []*string `json:"IpWhiteList,omitnil,omitempty" name:"IpWhiteList"`

	// Preset ACL rule. `1`: enable, `0`: disable. Default value: `0`.
	EnableAclRule *int64 `json:"EnableAclRule,omitnil,omitempty" name:"EnableAclRule"`

	// Name of the preset ACL rule.
	AclRuleName *string `json:"AclRuleName,omitnil,omitempty" name:"AclRuleName"`

	// Message retention file size in bytes, which is an optional parameter. Default value: -1. Currently, the min value that can be entered is 1,048,576 B.
	RetentionBytes *int64 `json:"RetentionBytes,omitnil,omitempty" name:"RetentionBytes"`

	// Tag list.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Production throttling in MB/sec.
	QuotaProducerByteRate *int64 `json:"QuotaProducerByteRate,omitnil,omitempty" name:"QuotaProducerByteRate"`

	// Consumption throttling in MB/sec.
	QuotaConsumerByteRate *int64 `json:"QuotaConsumerByteRate,omitnil,omitempty" name:"QuotaConsumerByteRate"`

	// The number of topic replicas.
	ReplicaNum *int64 `json:"ReplicaNum,omitnil,omitempty" name:"ReplicaNum"`
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
	delete(f, "QuotaProducerByteRate")
	delete(f, "QuotaConsumerByteRate")
	delete(f, "ReplicaNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTopicAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTopicAttributesResponseParams struct {
	// Returned result set
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTopicAttributesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTopicAttributesResponseParams `json:"Response"`
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
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// RouteIdDto Note: This field may return null, indicating that no valid values can be obtained.
	RouteDTO *RouteDTO `json:"RouteDTO,omitnil,omitempty" name:"RouteDTO"`
}

type Partition struct {
	// Partition ID
	PartitionId *int64 `json:"PartitionId,omitnil,omitempty" name:"PartitionId"`
}

type PartitionOffset struct {
	// Partition, such as "0" or "1"
	// Note: this field may return null, indicating that no valid values can be obtained.
	Partition *string `json:"Partition,omitnil,omitempty" name:"Partition"`

	// Offset, such as 100
	// Note: this field may return null, indicating that no valid values can be obtained.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type Partitions struct {
	// Partition.
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// Partition consumption offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type Price struct {
	// Discounted price
	RealTotalCost *float64 `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// Original price
	TotalCost *float64 `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`
}

type Region struct {
	// Region ID
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Region name
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// Area name
	AreaName *string `json:"AreaName,omitnil,omitempty" name:"AreaName"`

	// Region code
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	RegionCode *string `json:"RegionCode,omitnil,omitempty" name:"RegionCode"`

	// Region code (v3)
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	RegionCodeV3 *string `json:"RegionCodeV3,omitnil,omitempty" name:"RegionCodeV3"`

	// NONE: no special models are supported by default.\nCVM: the CVM type is supported.
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	Support *string `json:"Support,omitnil,omitempty" name:"Support"`

	// Whether IPv6 is supported. `0` indicates no, and `1` indicates yes.
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	Ipv6 *int64 `json:"Ipv6,omitnil,omitempty" name:"Ipv6"`

	// Whether cross-AZ clusters are supported.`0` indicates no, and `1` indicates yes.
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	MultiZone *int64 `json:"MultiZone,omitnil,omitempty" name:"MultiZone"`
}

type Route struct {
	// Instance connection method
	// 0: PLAINTEXT (plaintext method, which does not carry user information and is supported for legacy versions and Community Edition)
	// 1: SASL_PLAINTEXT (plaintext method, which authenticates the login through SASL before data start and is supported only for Community Edition)
	// 2: SSL (SSL-encrypted communication, which does not carry user information and is supported for legacy versions and Community Edition)
	// 3: SASL_SSL (SSL-encrypted communication, which authenticates the login through SASL before data start and is supported only for Community Edition)
	AccessType *int64 `json:"AccessType,omitnil,omitempty" name:"AccessType"`

	// Route ID
	RouteId *int64 `json:"RouteId,omitnil,omitempty" name:"RouteId"`

	// VIP network type (1: Public network TGW; 2: Classic network; 3: VPC; 4: Supporting network (IDC environment); 5: SSL public network access; 6: BM VPC; 7: Supporting network (CVM environment)).
	VipType *int64 `json:"VipType,omitnil,omitempty" name:"VipType"`

	// Virtual IP list
	VipList []*VipEntity `json:"VipList,omitnil,omitempty" name:"VipList"`

	// Domain name
	// Note: this field may return null, indicating that no valid values can be obtained.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Domain name port
	// Note: this field may return null, indicating that no valid values can be obtained.
	DomainPort *int64 `json:"DomainPort,omitnil,omitempty" name:"DomainPort"`

	// Timestamp
	// Note: this field may return null, indicating that no valid values can be obtained.
	DeleteTimestamp *string `json:"DeleteTimestamp,omitnil,omitempty" name:"DeleteTimestamp"`
}

type RouteDTO struct {
	// RouteId11 Note: This field may return null, indicating that no valid values can be obtained.
	RouteId *int64 `json:"RouteId,omitnil,omitempty" name:"RouteId"`
}

type RouteResponse struct {
	// Route information list
	// Note: this field may return null, indicating that no valid values can be obtained.
	Routers []*Route `json:"Routers,omitnil,omitempty" name:"Routers"`
}

type SaleInfo struct {
	// Manually set flag.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Flag *bool `json:"Flag,omitnil,omitempty" name:"Flag"`

	// CKafka version (v1.1.1/2.4.2/0.10.2）
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// Whether it is Pro Edition or Standard Edition.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// Whether it has been sold out. `true`: sold out.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SoldOut *bool `json:"SoldOut,omitnil,omitempty" name:"SoldOut"`
}

// Predefined struct for user
type SendMessageRequestParams struct {
	// Datahub access ID.
	DataHubId *string `json:"DataHubId,omitnil,omitempty" name:"DataHubId"`

	// Content of the message that has been sent. Up to 500 messages can be sent in a single request.
	Message []*BatchContent `json:"Message,omitnil,omitempty" name:"Message"`
}

type SendMessageRequest struct {
	*tchttp.BaseRequest
	
	// Datahub access ID.
	DataHubId *string `json:"DataHubId,omitnil,omitempty" name:"DataHubId"`

	// Content of the message that has been sent. Up to 500 messages can be sent in a single request.
	Message []*BatchContent `json:"Message,omitnil,omitempty" name:"Message"`
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

// Predefined struct for user
type SendMessageResponseParams struct {
	// Message ID list.
	MessageId []*string `json:"MessageId,omitnil,omitempty" name:"MessageId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SendMessageResponse struct {
	*tchttp.BaseResponse
	Response *SendMessageResponseParams `json:"Response"`
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
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Subscribed partition
	// Note: this field may return null, indicating that no valid values can be obtained.
	Partition []*int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// Partition offset information
	// Note: this field may return null, indicating that no valid values can be obtained.
	PartitionOffset []*PartitionOffset `json:"PartitionOffset,omitnil,omitempty" name:"PartitionOffset"`

	// ID of the subscribed topic. 
	// Note: this field may return null, indicating that no valid values can be obtained.
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

type Tag struct {
	// Tag key
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TaskStatusResponse struct {
	// Task status. `0` (Successful), `1` (Failed), `2` ( Running)
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Output information Note: This field may return null, indicating that no valid values can be obtained.
	Output *string `json:"Output,omitnil,omitempty" name:"Output"`
}

type Topic struct {
	// Topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Topic name
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Remarks
	// Note: this field may return null, indicating that no valid values can be obtained.
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`
}

type TopicAttributesResponse struct {
	// Topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Creation time
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Topic remarks
	// Note: this field may return null, indicating that no valid values can be obtained.
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// Number of partitions
	PartitionNum *int64 `json:"PartitionNum,omitnil,omitempty" name:"PartitionNum"`

	// IP allowlist switch. 1: enabled, 0: disabled
	EnableWhiteList *int64 `json:"EnableWhiteList,omitnil,omitempty" name:"EnableWhiteList"`

	// IP allowlist list
	IpWhiteList []*string `json:"IpWhiteList,omitnil,omitempty" name:"IpWhiteList"`

	// Topic configuration array
	Config *Config `json:"Config,omitnil,omitempty" name:"Config"`

	// Partition details
	Partitions []*TopicPartitionDO `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// Switch of the preset ACL rule. `1`: enable, `0`: disable.
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	EnableAclRule *int64 `json:"EnableAclRule,omitnil,omitempty" name:"EnableAclRule"`

	// Preset ACL rule list.
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	AclRuleList []*AclRule `json:"AclRuleList,omitnil,omitempty" name:"AclRuleList"`

	// Traffic throttling policy in topic dimension.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	QuotaConfig *InstanceQuotaConfigResp `json:"QuotaConfig,omitnil,omitempty" name:"QuotaConfig"`

	// Number of replicas
	// Note: This field may return null, indicating that no valid values can be obtained.
	ReplicaNum *int64 `json:"ReplicaNum,omitnil,omitempty" name:"ReplicaNum"`
}

type TopicDetail struct {
	// Topic name
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Number of partitions
	PartitionNum *int64 `json:"PartitionNum,omitnil,omitempty" name:"PartitionNum"`

	// Number of replicas
	ReplicaNum *int64 `json:"ReplicaNum,omitnil,omitempty" name:"ReplicaNum"`

	// Remarks
	// Note: this field may return null, indicating that no valid values can be obtained.
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// Creation time
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Whether to enable IP authentication allowlist. true: yes, false: no
	EnableWhiteList *bool `json:"EnableWhiteList,omitnil,omitempty" name:"EnableWhiteList"`

	// Number of IPs in IP allowlist
	IpWhiteListCount *int64 `json:"IpWhiteListCount,omitnil,omitempty" name:"IpWhiteListCount"`

	// COS bucket for data backup: address of the destination COS bucket
	// Note: this field may return null, indicating that no valid values can be obtained.
	ForwardCosBucket *string `json:"ForwardCosBucket,omitnil,omitempty" name:"ForwardCosBucket"`

	// Status of data backup to COS. 1: not enabled, 0: enabled
	ForwardStatus *int64 `json:"ForwardStatus,omitnil,omitempty" name:"ForwardStatus"`

	// Frequency of data backup to COS
	ForwardInterval *int64 `json:"ForwardInterval,omitnil,omitempty" name:"ForwardInterval"`

	// Advanced configuration
	// Note: this field may return null, indicating that no valid values can be obtained.
	Config *Config `json:"Config,omitnil,omitempty" name:"Config"`

	// Message retention time configuration (for recording the latest retention time)
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	RetentionTimeConfig *TopicRetentionTimeConfigRsp `json:"RetentionTimeConfig,omitnil,omitempty" name:"RetentionTimeConfig"`

	// `0`: normal, `1`: deleted, `2`: deleting
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Tag list
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type TopicDetailResponse struct {
	// List of returned topic details
	// Note: this field may return null, indicating that no valid values can be obtained.
	TopicList []*TopicDetail `json:"TopicList,omitnil,omitempty" name:"TopicList"`

	// Number of all eligible topic details
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type TopicInSyncReplicaInfo struct {
	// Partition name
	Partition *string `json:"Partition,omitnil,omitempty" name:"Partition"`

	// Leader ID
	Leader *uint64 `json:"Leader,omitnil,omitempty" name:"Leader"`

	// Replica set
	Replica *string `json:"Replica,omitnil,omitempty" name:"Replica"`

	// ISR
	InSyncReplica *string `json:"InSyncReplica,omitnil,omitempty" name:"InSyncReplica"`

	// Starting offset
	// Note: this field may return null, indicating that no valid values can be obtained.
	BeginOffset *uint64 `json:"BeginOffset,omitnil,omitempty" name:"BeginOffset"`

	// Ending offset
	// Note: this field may return null, indicating that no valid values can be obtained.
	EndOffset *uint64 `json:"EndOffset,omitnil,omitempty" name:"EndOffset"`

	// Number of messages
	// Note: this field may return null, indicating that no valid values can be obtained.
	MessageCount *uint64 `json:"MessageCount,omitnil,omitempty" name:"MessageCount"`

	// Unsynced replica set
	// Note: this field may return null, indicating that no valid values can be obtained.
	OutOfSyncReplica *string `json:"OutOfSyncReplica,omitnil,omitempty" name:"OutOfSyncReplica"`
}

type TopicInSyncReplicaResult struct {
	// Set of topic details and replicas
	TopicInSyncReplicaList []*TopicInSyncReplicaInfo `json:"TopicInSyncReplicaList,omitnil,omitempty" name:"TopicInSyncReplicaList"`

	// Total number
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type TopicPartitionDO struct {
	// Partition ID
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// Leader running status
	LeaderStatus *int64 `json:"LeaderStatus,omitnil,omitempty" name:"LeaderStatus"`

	// ISR quantity
	IsrNum *int64 `json:"IsrNum,omitnil,omitempty" name:"IsrNum"`

	// Number of replicas
	ReplicaNum *int64 `json:"ReplicaNum,omitnil,omitempty" name:"ReplicaNum"`
}

type TopicResult struct {
	// List of returned topic information
	// Note: this field may return null, indicating that no valid values can be obtained.
	TopicList []*Topic `json:"TopicList,omitnil,omitempty" name:"TopicList"`

	// Number of eligible topics
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type TopicRetentionTimeConfigRsp struct {
	// Expected value, i.e., the topic message retention time (min) configured
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	Expect *int64 `json:"Expect,omitnil,omitempty" name:"Expect"`

	// Current value (min), i.e., the retention time currently in effect, which may be dynamically adjusted
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	Current *int64 `json:"Current,omitnil,omitempty" name:"Current"`

	// Last modified time
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	ModTimeStamp *int64 `json:"ModTimeStamp,omitnil,omitempty" name:"ModTimeStamp"`
}

type TopicSubscribeGroup struct {
	// Total number
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Number of consumer group status
	StatusCountInfo *string `json:"StatusCountInfo,omitnil,omitempty" name:"StatusCountInfo"`

	// Consumer group information
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	GroupsInfo []*GroupInfoResponse `json:"GroupsInfo,omitnil,omitempty" name:"GroupsInfo"`

	// Whether a request is asynchronous. If there are fewer consumer groups in the instances, the result will be returned directly, and status code is 1. When there are many consumer groups in the instances, cache will be updated asynchronously. When status code is 0, grouping information will not be returned until cache update is completed and status code becomes 1.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type User struct {
	// User ID
	UserId *int64 `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Username
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Last updated time
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type UserResponse struct {
	// List of eligible users
	// Note: this field may return null, indicating that no valid values can be obtained.
	Users []*User `json:"Users,omitnil,omitempty" name:"Users"`

	// Total number of eligible users
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type VipEntity struct {
	// Virtual IP
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Virtual port
	Vport *string `json:"Vport,omitnil,omitempty" name:"Vport"`
}

type ZoneInfo struct {
	// Zone ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Whether it is an internal application.
	IsInternalApp *int64 `json:"IsInternalApp,omitnil,omitempty" name:"IsInternalApp"`

	// Application ID
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Flag
	Flag *bool `json:"Flag,omitnil,omitempty" name:"Flag"`

	// Zone name
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// Zone status
	ZoneStatus *int64 `json:"ZoneStatus,omitnil,omitempty" name:"ZoneStatus"`

	// Extra flag
	Exflag *string `json:"Exflag,omitnil,omitempty" name:"Exflag"`

	// JSON object. The key is the model. The value `true` means “sold out”, and `false` means “not sold out”.
	SoldOut *string `json:"SoldOut,omitnil,omitempty" name:"SoldOut"`

	// Information on whether Standard Edition has been sold out.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SalesInfo []*SaleInfo `json:"SalesInfo,omitnil,omitempty" name:"SalesInfo"`
}

type ZoneResponse struct {
	// Zone list
	ZoneList []*ZoneInfo `json:"ZoneList,omitnil,omitempty" name:"ZoneList"`

	// Maximum number of instances to be purchased
	MaxBuyInstanceNum *int64 `json:"MaxBuyInstanceNum,omitnil,omitempty" name:"MaxBuyInstanceNum"`

	// Maximum bandwidth in MB/S
	MaxBandwidth *int64 `json:"MaxBandwidth,omitnil,omitempty" name:"MaxBandwidth"`

	// Pay-as-you-go unit price
	UnitPrice *Price `json:"UnitPrice,omitnil,omitempty" name:"UnitPrice"`

	// Pay-as-you-go unit message price
	MessagePrice *Price `json:"MessagePrice,omitnil,omitempty" name:"MessagePrice"`

	// Cluster information dedicated to a user
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	ClusterInfo []*ClusterInfo `json:"ClusterInfo,omitnil,omitempty" name:"ClusterInfo"`

	// Purchase of Standard Edition configurations
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	Standard *string `json:"Standard,omitnil,omitempty" name:"Standard"`

	// Purchase of Standard S2 Edition configurations
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	StandardS2 *string `json:"StandardS2,omitnil,omitempty" name:"StandardS2"`

	// Purchase of Pro Edition configurations
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	Profession *string `json:"Profession,omitnil,omitempty" name:"Profession"`

	// Purchase of Physical Dedicated Edition configurations
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	Physical *string `json:"Physical,omitnil,omitempty" name:"Physical"`

	// Public network bandwidth.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	PublicNetwork *string `json:"PublicNetwork,omitnil,omitempty" name:"PublicNetwork"`

	// Public network bandwidth configuration.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	PublicNetworkLimit *string `json:"PublicNetworkLimit,omitnil,omitempty" name:"PublicNetworkLimit"`
}