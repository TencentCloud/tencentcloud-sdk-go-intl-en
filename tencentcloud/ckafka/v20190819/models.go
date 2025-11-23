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

	// List of users, defaults to User:*, means any User is accessible in the entire region. the current User can only be the User in the list of users.
	Principal *string `json:"Principal,omitnil,omitempty" name:"Principal"`

	// Defaults to *, indicating any host is accessible in the entire region. currently, ckafka does not support * as the host, however, the following open-source kafka productization will directly support it.
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// ACL operation mode. 0: UNKNOWN, 1: ANY, 2: ALL, 3: READ, 4: WRITE, 5: CREATE, 6: DELETE, 7: ALTER, 8: DESCRIBE, 9: CLUSTER_ACTION, 10: DESCRIBE_CONFIGS, 11: ALTER_CONFIGS, 12: IDEMPOTEN_WRITE
	Operation *int64 `json:"Operation,omitnil,omitempty" name:"Operation"`

	// Permission type. 0: UNKNOWN, 1: ANY, 2: DENY, 3: ALLOW
	PermissionType *int64 `json:"PermissionType,omitnil,omitempty" name:"PermissionType"`
}

type AclResponse struct {
	// Number of eligible data entries
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// ACL list.
	AclList []*Acl `json:"AclList,omitnil,omitempty" name:"AclList"`
}

type AclRule struct {
	// ACL rule name.
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// The ckafka cluster instance Id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ACL rule-based matching type. currently only supports prefix match. valid values: PREFIXED.
	PatternType *string `json:"PatternType,omitnil,omitempty" name:"PatternType"`

	// Indicates the prefix value for prefix match.
	Pattern *string `json:"Pattern,omitnil,omitempty" name:"Pattern"`

	// Acl resource type, currently only support Topic. valid values: Topic.
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// Specifies the ACL information contained in the rule.
	AclList *string `json:"AclList,omitnil,omitempty" name:"AclList"`

	// Specifies the time when the rule was created.
	CreateTimeStamp *string `json:"CreateTimeStamp,omitnil,omitempty" name:"CreateTimeStamp"`

	// Specifies whether to apply the preset ACL rule to newly-added topics.
	IsApplied *int64 `json:"IsApplied,omitnil,omitempty" name:"IsApplied"`

	// Rule update time.
	UpdateTimeStamp *string `json:"UpdateTimeStamp,omitnil,omitempty" name:"UpdateTimeStamp"`

	// Specifies the remark of the rule.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// One of the displayed corresponding TopicName.
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Number of topics to which the ACL rule is applied.
	TopicCount *int64 `json:"TopicCount,omitnil,omitempty" name:"TopicCount"`

	// Specifies the pattern type.
	PatternTypeTitle *string `json:"PatternTypeTitle,omitnil,omitempty" name:"PatternTypeTitle"`
}

type AclRuleInfo struct {
	// ACL operation types. Enumerated values: `All` (all operations), `Read` (read), `Write` (write).
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// Permission type. Deny: Deny. Allow: permission.
	PermissionType *string `json:"PermissionType,omitnil,omitempty" name:"PermissionType"`

	// Indicates any host is accessible in the entire region.
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// The User. User:* means any User is accessible in the entire region. the current User can only be the User in the list of users. the input format requires the [User:] prefix. for example, for User A, input User:A.
	Principal *string `json:"Principal,omitnil,omitempty" name:"Principal"`
}

type AclRuleResp struct {
	// Total number of data entries
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// AclRule list.
	AclRuleList []*AclRule `json:"AclRuleList,omitnil,omitempty" name:"AclRuleList"`
}

type Assignment struct {
	// Assignment version information
	Version *int64 `json:"Version,omitnil,omitempty" name:"Version"`

	// topic information list.
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
	// The ckafka cluster instance Id, which can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ACL resource type. Default value: `2` (topic).
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// Resource list array, obtainable through the DescribeTopic API (https://www.tencentcloud.comom/document/product/597/40847?from_cn_redirect=1).
	ResourceNames []*string `json:"ResourceNames,omitnil,omitempty" name:"ResourceNames"`

	// Specifies the set ACL rule list, which can be obtained through the DescribeAclRule API (https://www.tencentcloud.comom/document/product/597/89217?from_cn_redirect=1).
	RuleList []*AclRuleInfo `json:"RuleList,omitnil,omitempty" name:"RuleList"`
}

type BatchCreateAclRequest struct {
	*tchttp.BaseRequest
	
	// The ckafka cluster instance Id, which can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ACL resource type. Default value: `2` (topic).
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// Resource list array, obtainable through the DescribeTopic API (https://www.tencentcloud.comom/document/product/597/40847?from_cn_redirect=1).
	ResourceNames []*string `json:"ResourceNames,omitnil,omitempty" name:"ResourceNames"`

	// Specifies the set ACL rule list, which can be obtained through the DescribeAclRule API (https://www.tencentcloud.comom/document/product/597/89217?from_cn_redirect=1).
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
	// Status code: 0 - modification succeeded, otherwise modification failed.
	Result *int64 `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// The ckafka cluster instance Id.
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

	// The ckafka cluster instance Id.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// The ckafka cluster instance Id, which can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Specifies the topic attribute list (a maximum of 10 per batch), which can be obtained through the [DescribeTopic](https://www.tencentcloud.comom/document/product/597/40847?from_cn_redirect=1) api.
	Topic []*BatchModifyTopicInfo `json:"Topic,omitnil,omitempty" name:"Topic"`
}

type BatchModifyTopicAttributesRequest struct {
	*tchttp.BaseRequest
	
	// The ckafka cluster instance Id, which can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Specifies the topic attribute list (a maximum of 10 per batch), which can be obtained through the [DescribeTopic](https://www.tencentcloud.comom/document/product/597/40847?from_cn_redirect=1) api.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// Topic name
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

	// Specifies the message retention size in the topic dimension in bytes. value range: 1 GB to 1024 GB.
	RetentionBytes *int64 `json:"RetentionBytes,omitnil,omitempty" name:"RetentionBytes"`

	// Duration of Segment shard scrolling in milliseconds. value range: 1 day to 90 days.
	SegmentMs *int64 `json:"SegmentMs,omitnil,omitempty" name:"SegmentMs"`

	// Message size per batch. Value range: 1 KB - 12 MB.
	MaxMessageBytes *int64 `json:"MaxMessageBytes,omitnil,omitempty" name:"MaxMessageBytes"`

	// Specifies the time type for message storage: CreateTime/LogAppendTime.
	LogMsgTimestampType *string `json:"LogMsgTimestampType,omitnil,omitempty" name:"LogMsgTimestampType"`
}

type BatchModifyTopicResultDTO struct {
	// The ckafka cluster instance Id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Topic name
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Operation return code.
	ReturnCode *string `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// Returned information.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

type ClusterInfo struct {
	// Cluster ID
	ClusterId *int64 `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Cluster name
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// Maximum disk of the cluster (unit: GB).
	MaxDiskSize *int64 `json:"MaxDiskSize,omitnil,omitempty" name:"MaxDiskSize"`

	// Maximum bandwidth of the cluster. unit: MB/s.
	MaxBandWidth *int64 `json:"MaxBandWidth,omitnil,omitempty" name:"MaxBandWidth"`

	// Current availability of cluster disk (unit: GB).
	AvailableDiskSize *int64 `json:"AvailableDiskSize,omitnil,omitempty" name:"AvailableDiskSize"`

	// Available bandwidth of the cluster. unit: MB/s.
	AvailableBandWidth *int64 `json:"AvailableBandWidth,omitnil,omitempty" name:"AvailableBandWidth"`

	// Indicates the AZ to which the cluster belongs.
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The AZ where the cluster nodes are located. If the cluster is a cross-AZ cluster, it includes multiple AZs where the cluster nodes are located.
	ZoneIds []*int64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`
}

type Config struct {
	// Message retention period in milliseconds.
	Retention *int64 `json:"Retention,omitnil,omitempty" name:"Retention"`

	// Minimum number of sync replications
	// Note: this field may return null, indicating that no valid values can be obtained.
	MinInsyncReplicas *int64 `json:"MinInsyncReplicas,omitnil,omitempty" name:"MinInsyncReplicas"`

	// Log cleanup mode. Default value: delete.
	// delete: logs will be deleted by save time; compact: logs will be compressed by key; compact, delete: logs will be compressed by key and deleted by save time.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CleanUpPolicy *string `json:"CleanUpPolicy,omitnil,omitempty" name:"CleanUpPolicy"`

	// Duration of Segment shard scrolling in milliseconds.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SegmentMs *int64 `json:"SegmentMs,omitnil,omitempty" name:"SegmentMs"`

	// 0: false, 1: true.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UncleanLeaderElectionEnable *int64 `json:"UncleanLeaderElectionEnable,omitnil,omitempty" name:"UncleanLeaderElectionEnable"`

	// Segment specifies the number of bytes for sharding scroll. unit: bytes.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SegmentBytes *int64 `json:"SegmentBytes,omitnil,omitempty" name:"SegmentBytes"`

	// Maximum message byte size. unit: bytes.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MaxMessageBytes *int64 `json:"MaxMessageBytes,omitnil,omitempty" name:"MaxMessageBytes"`

	// Specifies the message retention file size in Bytes.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RetentionBytes *int64 `json:"RetentionBytes,omitnil,omitempty" name:"RetentionBytes"`

	// The time type for message saving. CreateTime means the time when the producer created this message. LogAppendTime means the time when the broker received the message.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LogMsgTimestampType *string `json:"LogMsgTimestampType,omitnil,omitempty" name:"LogMsgTimestampType"`
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

	// Topic list.
	TopicList []*ConsumerGroupTopic `json:"TopicList,omitnil,omitempty" name:"TopicList"`

	// Specifies the consumption group List.
	GroupList []*ConsumerGroup `json:"GroupList,omitnil,omitempty" name:"GroupList"`

	// Total number of partitions.
	TotalPartition *int64 `json:"TotalPartition,omitnil,omitempty" name:"TotalPartition"`

	// Monitored partition list.
	PartitionListForMonitor []*Partition `json:"PartitionListForMonitor,omitnil,omitempty" name:"PartitionListForMonitor"`

	// Total number of topics.
	TotalTopic *int64 `json:"TotalTopic,omitnil,omitempty" name:"TotalTopic"`

	// Monitored topic list.
	TopicListForMonitor []*ConsumerGroupTopic `json:"TopicListForMonitor,omitnil,omitempty" name:"TopicListForMonitor"`

	// Monitored group list.
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

	// Message timestamp.
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// Message headers
	// Note: This field may return null, indicating that no valid values can be obtained.
	Headers *string `json:"Headers,omitnil,omitempty" name:"Headers"`
}

// Predefined struct for user
type CreateAclRequestParams struct {
	// The ckafka cluster instance Id, which can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ACL resource type (`2`: TOPIC, `3`: GROUP, `4`: CLUSTER).
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// ACL operation type (`2`: ALL, `3`: READ, `4`: WRITE, `5`: CREATE, `6`: DELETE, `7`: ALTER, `8`: DESCRIBE, `9`: CLUSTER_ACTION, `10`: DESCRIBE_CONFIGS, `11`: ALTER_CONFIGS, `12`: IDEMPOTENT_WRITE).
	Operation *int64 `json:"Operation,omitnil,omitempty" name:"Operation"`

	// Permission type (2:DENY, 3:ALLOW). currently ckafka supports ALLOW (equivalent to allowlist), others used when compatible with open-source kafka acl.
	PermissionType *int64 `json:"PermissionType,omitnil,omitempty" name:"PermissionType"`

	// Resource name, which is related to `resourceType`. For example, if `resourceType` is `TOPIC`, this field indicates the topic name; if `resourceType` is `GROUP`, this field indicates the group name; if `resourceType` is `CLUSTER`, this field can be left empty.
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// Defaults to *, indicating any host is accessible in the entire region. supports filling in ips or ranges, and uses ";" for separation.
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// The list of users allowed to access the topic. Default: User:*, meaning all users. The current user must be in the user list. Add `User:` before the user name (`User:A` for example).
	Principal *string `json:"Principal,omitnil,omitempty" name:"Principal"`

	// The resource name list, which is in JSON string format. Either `ResourceName` or `resourceNameList` can be specified.
	ResourceNameList *string `json:"ResourceNameList,omitnil,omitempty" name:"ResourceNameList"`
}

type CreateAclRequest struct {
	*tchttp.BaseRequest
	
	// The ckafka cluster instance Id, which can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ACL resource type (`2`: TOPIC, `3`: GROUP, `4`: CLUSTER).
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// ACL operation type (`2`: ALL, `3`: READ, `4`: WRITE, `5`: CREATE, `6`: DELETE, `7`: ALTER, `8`: DESCRIBE, `9`: CLUSTER_ACTION, `10`: DESCRIBE_CONFIGS, `11`: ALTER_CONFIGS, `12`: IDEMPOTENT_WRITE).
	Operation *int64 `json:"Operation,omitnil,omitempty" name:"Operation"`

	// Permission type (2:DENY, 3:ALLOW). currently ckafka supports ALLOW (equivalent to allowlist), others used when compatible with open-source kafka acl.
	PermissionType *int64 `json:"PermissionType,omitnil,omitempty" name:"PermissionType"`

	// Resource name, which is related to `resourceType`. For example, if `resourceType` is `TOPIC`, this field indicates the topic name; if `resourceType` is `GROUP`, this field indicates the group name; if `resourceType` is `CLUSTER`, this field can be left empty.
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// Defaults to *, indicating any host is accessible in the entire region. supports filling in ips or ranges, and uses ";" for separation.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// The ckafka cluster instance Id, which can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ACL resource type. Currently, the only valid value is `Topic`.
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// ACL rule-based matching type. currently supports prefix match and PRESET policy. valid values: PREFIXED/PRESET.
	PatternType *string `json:"PatternType,omitnil,omitempty" name:"PatternType"`

	// Rule name
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// ACL rule list
	RuleList []*AclRuleInfo `json:"RuleList,omitnil,omitempty" name:"RuleList"`

	// Indicates the prefix for prefix match. this parameter is required when PatternType value is PREFIXED.
	Pattern *string `json:"Pattern,omitnil,omitempty" name:"Pattern"`

	// Specifies whether to apply the preset ACL rule to newly-added topics. defaults to 0, which means no. a value of 1 means yes.
	IsApplied *int64 `json:"IsApplied,omitnil,omitempty" name:"IsApplied"`

	// Remarks for ACL rules
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

type CreateAclRuleRequest struct {
	*tchttp.BaseRequest
	
	// The ckafka cluster instance Id, which can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ACL resource type. Currently, the only valid value is `Topic`.
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// ACL rule-based matching type. currently supports prefix match and PRESET policy. valid values: PREFIXED/PRESET.
	PatternType *string `json:"PatternType,omitnil,omitempty" name:"PatternType"`

	// Rule name
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// ACL rule list
	RuleList []*AclRuleInfo `json:"RuleList,omitnil,omitempty" name:"RuleList"`

	// Indicates the prefix for prefix match. this parameter is required when PatternType value is PREFIXED.
	Pattern *string `json:"Pattern,omitnil,omitempty" name:"Pattern"`

	// Specifies whether to apply the preset ACL rule to newly-added topics. defaults to 0, which means no. a value of 1 means yes.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// The ckafka cluster instance Id, which can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Consumer group name.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Topic name. one of TopicName or TopicNameList must display a specified existing topic name.
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Topic name list.
	TopicNameList []*string `json:"TopicNameList,omitnil,omitempty" name:"TopicNameList"`
}

type CreateConsumerRequest struct {
	*tchttp.BaseRequest
	
	// The ckafka cluster instance Id, which can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Consumer group name.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Topic name. one of TopicName or TopicNameList must display a specified existing topic name.
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Topic name list.
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
	// Create consumer group returned results.
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// Name, a string of no more than 128 characters, must start with "AppId-" and can contain letters, digits, and hyphens (-).
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
	
	// Name, a string of no more than 128 characters, must start with "AppId-" and can contain letters, digits, and hyphens (-).
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// CreateInstancePre returns fixed as 0. it cannot be used as a query condition for CheckTaskStatus. this is merely to ensure alignment with the backend data structure.
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Order ID list
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// The ckafka cluster instance Id. by default, returns the Id of the first purchased instance when purchasing multiple instances.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Order and purchase mapping list corresponding to the instance.
	DealNameInstanceIdMapping []*DealInstanceDTO `json:"DealNameInstanceIdMapping,omitnil,omitempty" name:"DealNameInstanceIdMapping"`
}

type CreateInstancePostResp struct {
	// Returned code. `0` indicates normal status while other codes indicate errors.
	ReturnCode *string `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// Message returned by the API. An error message will be returned if the API reports an error. 
	ReturnMessage *string `json:"ReturnMessage,omitnil,omitempty" name:"ReturnMessage"`

	// Specifies the Data returned.
	Data *CreateInstancePostData `json:"Data,omitnil,omitempty" name:"Data"`
}

type CreateInstancePreData struct {
	// CreateInstancePre returns fixed as 0. it cannot be used as a query condition for CheckTaskStatus. this is merely to ensure alignment with the backend data structure.
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Order ID list
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// The ckafka cluster instance Id. by default, returns the Id of the first purchased instance when purchasing multiple instances.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Order and purchase mapping list corresponding to the instance.
	DealNameInstanceIdMapping []*DealInstanceDTO `json:"DealNameInstanceIdMapping,omitnil,omitempty" name:"DealNameInstanceIdMapping"`
}

// Predefined struct for user
type CreateInstancePreRequestParams struct {
	// Specifies the ckafka cluster instance Name, an arbitrary string with length no more than 128 characters.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Availability zone. when purchasing a multi-availability zone instance, this parameter specifies the primary az. [view availability zones](https://www.tencentcloud.comom/document/product/597/55246?from_cn_redirect=1).
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Prepaid purchase duration, such as "1m", exactly one month. value ranges from 1m to 36m.
	Period *string `json:"Period,omitnil,omitempty" name:"Period"`

	// Specifies the standard edition instance specification for the international site. currently only the international site standard edition uses the current field to distinguish specifications, while the domestic site standard edition distinguishes specifications by peak bandwidth. fill in 1 for all instances except the international site standard edition. for international site standard edition instances: [entry-level (general)] fill 1; [standard type (standard)] fill 2; [advanced] fill 3; [capacity type (capacity)] fill 4; [advanced type 1 (specialized-1)] fill 5; [advanced type 2 (specialized-2)] fill 6; [advanced type 3 (specialized-3)] fill 7; [advanced type 4 (specialized-4)] fill 8.
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// VPC Id.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Optional. maximum retention time of instance logs, in minutes. default value: 1440 (1 day). value range: 1 minute to 90 days.
	MsgRetentionTime *int64 `json:"MsgRetentionTime,omitnil,omitempty" name:"MsgRetentionTime"`

	// Specifies the cluster Id when creating an instance.
	ClusterId *int64 `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Auto-Renewal tag for prepaid services. valid values: 0 (default state, not set by the user, initial status), 1 (auto-renew), 2 (explicitly no auto-renew, set by the user).
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// Specifies the CKafka version number. valid values: 2.4.1, 2.4.2, 2.8.1, 3.2.3. default value 2.4.1. 2.4.1 and 2.4.2 belong to the same version. any can be passed.
	KafkaVersion *string `json:"KafkaVersion,omitnil,omitempty" name:"KafkaVersion"`

	// Specifies the instance type. valid values: standard (default), profession, premium.
	SpecificationsType *string `json:"SpecificationsType,omitnil,omitempty" name:"SpecificationsType"`

	// Disk size. if it does not match the console specification ratio, the creation cannot succeed. default value is 500. step length is set to 100. can be accessed through the following link: https://www.tencentcloud.comom/document/product/597/122562.?from_cn_redirect=1
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// Instance bandwidth. default value: 40 MB/s. minimum value: 20 MB/s. maximum value for advanced edition: 360 MB/s. maximum value for pro edition: 100000 MB/s. standard version fixed bandwidth specifications: 40 MB/s, 100 MB/s, 150 MB/s. view billing specifications through the following link: https://www.tencentcloud.comom/document/product/597/11745.?from_cn_redirect=1
	BandWidth *int64 `json:"BandWidth,omitnil,omitempty" name:"BandWidth"`

	// Partition size. if it does not match the console specification ratio, creation will fail. default value is 800, step length is 100. billing specifications can be viewed through the following link: https://www.tencentcloud.comom/document/product/597/122563.?from_cn_redirect=1
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// Tag.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Specifies the disk type for a pro/advanced edition instance. you do not need to fill it in for a standard edition instance. valid values: "CLOUD_SSD" for SSD CLOUD disk; "CLOUD_BASIC" for high-performance CLOUD block storage. default value: "CLOUD_BASIC".
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// Specifies whether to create a cross-az instance. when the current parameter is true, zoneIds is required.
	MultiZoneFlag *bool `json:"MultiZoneFlag,omitnil,omitempty" name:"MultiZoneFlag"`

	// Availability zone list. required item when purchasing a multi-availability zone instance.
	ZoneIds []*int64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// Public network bandwidth size, in Mbps. the default is no free 3 Mbps bandwidth. for example, for a total of 3 Mbps public network bandwidth, pass 0 here; for a total of 6 Mbps public network bandwidth, pass 3 here. default value is 0. ensure the input parameter is a multiple of 3.
	PublicNetworkMonthly *int64 `json:"PublicNetworkMonthly,omitnil,omitempty" name:"PublicNetworkMonthly"`

	// Number of instances to purchase. optional. default value is 1. when you input this parameter, it enables the creation of multiple instances with case-sensitive suffixes added to instanceName.
	InstanceNum *int64 `json:"InstanceNum,omitnil,omitempty" name:"InstanceNum"`

	// Whether to automatically select a voucher. valid values: 1 (yes), 0 (no). default is 0.
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// Elastic bandwidth switch. specifies whether to enable elastic bandwidth. valid values: 0 (not enabled, default), 1 (enabled).
	ElasticBandwidthSwitch *int64 `json:"ElasticBandwidthSwitch,omitnil,omitempty" name:"ElasticBandwidthSwitch"`
}

type CreateInstancePreRequest struct {
	*tchttp.BaseRequest
	
	// Specifies the ckafka cluster instance Name, an arbitrary string with length no more than 128 characters.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Availability zone. when purchasing a multi-availability zone instance, this parameter specifies the primary az. [view availability zones](https://www.tencentcloud.comom/document/product/597/55246?from_cn_redirect=1).
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Prepaid purchase duration, such as "1m", exactly one month. value ranges from 1m to 36m.
	Period *string `json:"Period,omitnil,omitempty" name:"Period"`

	// Specifies the standard edition instance specification for the international site. currently only the international site standard edition uses the current field to distinguish specifications, while the domestic site standard edition distinguishes specifications by peak bandwidth. fill in 1 for all instances except the international site standard edition. for international site standard edition instances: [entry-level (general)] fill 1; [standard type (standard)] fill 2; [advanced] fill 3; [capacity type (capacity)] fill 4; [advanced type 1 (specialized-1)] fill 5; [advanced type 2 (specialized-2)] fill 6; [advanced type 3 (specialized-3)] fill 7; [advanced type 4 (specialized-4)] fill 8.
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// VPC Id.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Optional. maximum retention time of instance logs, in minutes. default value: 1440 (1 day). value range: 1 minute to 90 days.
	MsgRetentionTime *int64 `json:"MsgRetentionTime,omitnil,omitempty" name:"MsgRetentionTime"`

	// Specifies the cluster Id when creating an instance.
	ClusterId *int64 `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Auto-Renewal tag for prepaid services. valid values: 0 (default state, not set by the user, initial status), 1 (auto-renew), 2 (explicitly no auto-renew, set by the user).
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// Specifies the CKafka version number. valid values: 2.4.1, 2.4.2, 2.8.1, 3.2.3. default value 2.4.1. 2.4.1 and 2.4.2 belong to the same version. any can be passed.
	KafkaVersion *string `json:"KafkaVersion,omitnil,omitempty" name:"KafkaVersion"`

	// Specifies the instance type. valid values: standard (default), profession, premium.
	SpecificationsType *string `json:"SpecificationsType,omitnil,omitempty" name:"SpecificationsType"`

	// Disk size. if it does not match the console specification ratio, the creation cannot succeed. default value is 500. step length is set to 100. can be accessed through the following link: https://www.tencentcloud.comom/document/product/597/122562.?from_cn_redirect=1
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// Instance bandwidth. default value: 40 MB/s. minimum value: 20 MB/s. maximum value for advanced edition: 360 MB/s. maximum value for pro edition: 100000 MB/s. standard version fixed bandwidth specifications: 40 MB/s, 100 MB/s, 150 MB/s. view billing specifications through the following link: https://www.tencentcloud.comom/document/product/597/11745.?from_cn_redirect=1
	BandWidth *int64 `json:"BandWidth,omitnil,omitempty" name:"BandWidth"`

	// Partition size. if it does not match the console specification ratio, creation will fail. default value is 800, step length is 100. billing specifications can be viewed through the following link: https://www.tencentcloud.comom/document/product/597/122563.?from_cn_redirect=1
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// Tag.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Specifies the disk type for a pro/advanced edition instance. you do not need to fill it in for a standard edition instance. valid values: "CLOUD_SSD" for SSD CLOUD disk; "CLOUD_BASIC" for high-performance CLOUD block storage. default value: "CLOUD_BASIC".
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// Specifies whether to create a cross-az instance. when the current parameter is true, zoneIds is required.
	MultiZoneFlag *bool `json:"MultiZoneFlag,omitnil,omitempty" name:"MultiZoneFlag"`

	// Availability zone list. required item when purchasing a multi-availability zone instance.
	ZoneIds []*int64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// Public network bandwidth size, in Mbps. the default is no free 3 Mbps bandwidth. for example, for a total of 3 Mbps public network bandwidth, pass 0 here; for a total of 6 Mbps public network bandwidth, pass 3 here. default value is 0. ensure the input parameter is a multiple of 3.
	PublicNetworkMonthly *int64 `json:"PublicNetworkMonthly,omitnil,omitempty" name:"PublicNetworkMonthly"`

	// Number of instances to purchase. optional. default value is 1. when you input this parameter, it enables the creation of multiple instances with case-sensitive suffixes added to instanceName.
	InstanceNum *int64 `json:"InstanceNum,omitnil,omitempty" name:"InstanceNum"`

	// Whether to automatically select a voucher. valid values: 1 (yes), 0 (no). default is 0.
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// Elastic bandwidth switch. specifies whether to enable elastic bandwidth. valid values: 0 (not enabled, default), 1 (enabled).
	ElasticBandwidthSwitch *int64 `json:"ElasticBandwidthSwitch,omitnil,omitempty" name:"ElasticBandwidthSwitch"`
}

func (r *CreateInstancePreRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstancePreRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceName")
	delete(f, "ZoneId")
	delete(f, "Period")
	delete(f, "InstanceType")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "MsgRetentionTime")
	delete(f, "ClusterId")
	delete(f, "RenewFlag")
	delete(f, "KafkaVersion")
	delete(f, "SpecificationsType")
	delete(f, "DiskSize")
	delete(f, "BandWidth")
	delete(f, "Partition")
	delete(f, "Tags")
	delete(f, "DiskType")
	delete(f, "MultiZoneFlag")
	delete(f, "ZoneIds")
	delete(f, "PublicNetworkMonthly")
	delete(f, "InstanceNum")
	delete(f, "AutoVoucher")
	delete(f, "ElasticBandwidthSwitch")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstancePreRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstancePreResp struct {
	// Returned code. 0: Normal; other values: Error.
	ReturnCode *string `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// The message indicating whether the operation is successful.
	ReturnMessage *string `json:"ReturnMessage,omitnil,omitempty" name:"ReturnMessage"`

	// Specifies the Data returned by the operation.
	Data *CreateInstancePreData `json:"Data,omitnil,omitempty" name:"Data"`

	// Deletion time.  This parameter has been deprecated and will be deleted.  Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: DeleteRouteTimestamp is deprecated.
	DeleteRouteTimestamp *string `json:"DeleteRouteTimestamp,omitnil,omitempty" name:"DeleteRouteTimestamp"`
}

// Predefined struct for user
type CreateInstancePreResponseParams struct {
	// Returned result.
	Result *CreateInstancePreResp `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateInstancePreResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstancePreResponseParams `json:"Response"`
}

func (r *CreateInstancePreResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstancePreResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePartitionRequestParams struct {
	// The ckafka cluster instance Id, which can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Specifies the topic name, which can be obtained through the [DescribeTopic](https://www.tencentcloud.comom/document/product/597/40847?from_cn_redirect=1) api.
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Topic partition count. the input parameter is the number of partitions after modification rather than adding partitions. therefore, the input parameter must exceed the current topic partition count.
	PartitionNum *int64 `json:"PartitionNum,omitnil,omitempty" name:"PartitionNum"`
}

type CreatePartitionRequest struct {
	*tchttp.BaseRequest
	
	// The ckafka cluster instance Id, which can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Specifies the topic name, which can be obtained through the [DescribeTopic](https://www.tencentcloud.comom/document/product/597/40847?from_cn_redirect=1) api.
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Topic partition count. the input parameter is the number of partitions after modification rather than adding partitions. therefore, the input parameter must exceed the current topic partition count.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// VPC Id, obtain through the API [DescribeVpcs](https://www.tencentcloud.comom/document/product/215/15778?from_cn_redirect=1).
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet Id. can be obtained through the [DescribeSubnets](https://www.tencentcloud.comom/document/product/215/15784?from_cn_redirect=1) api.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Specifies the cluster instance name of ckafka, an arbitrary character with length not exceeding 128.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Specifies the standard edition instance specification for the international site. currently only the international site standard edition uses the current field to distinguish specifications, while the domestic site standard edition distinguishes specifications by peak bandwidth. fill in 1 for all instances except the international site standard edition. for international site standard edition instances: [entry-level (general)] fill 1; [standard type (standard)] fill 2; [advanced] fill 3; [capacity type (capacity)] fill 4; [advanced type 1 (specialized-1)] fill 5; [advanced type 2 (specialized-2)] fill 6; [advanced type 3 (specialized-3)] fill 7; [advanced type 4 (specialized-4)] fill 8.
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// The maximum instance log retention period in minutes by default.  If this parameter is left empty, the default retention period is 1,440 minutes (1 day) to 30 days.  If the message retention period of the topic is explicitly set, it will prevail.
	MsgRetentionTime *int64 `json:"MsgRetentionTime,omitnil,omitempty" name:"MsgRetentionTime"`

	// Cluster ID, which can be selected when you create an instance.  You don’t need to pass in this parameter if the cluster where the instance resides is not specified.
	ClusterId *int64 `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Instance version. currently supports "2.4.1", "2.4.2", "2.8.1", "3.2.3". default value "2.4.1". "2.4.1" and "2.4.2" belong to the same version. any one can be passed.
	KafkaVersion *string `json:"KafkaVersion,omitnil,omitempty" name:"KafkaVersion"`

	// Instance type. "standard": standard version. "profession": pro edition. (standard version is only supported on the international site. currently, the chinese site supports pro edition.).
	SpecificationsType *string `json:"SpecificationsType,omitnil,omitempty" name:"SpecificationsType"`

	// Specifies the disk type for a pro edition instance. you do not need to fill it in for a standard edition instance. valid values: "CLOUD_SSD" for SSD CLOUD disk; "CLOUD_BASIC" for high-performance CLOUD block storage. default value: "CLOUD_BASIC".
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// Specifies the peak bandwidth of the instance private network, with a default value of 40 MB/s. for standard version, input the peak bandwidth corresponding to the current instance specifications. note that if the instance created is a pro edition instance, parameter configuration such as peak bandwidth and number of partitions should meet the billing specification of the professional edition. view billing specifications through the following link: https://www.tencentcloud.comom/document/product/597/11745.?from_cn_redirect=1
	BandWidth *int64 `json:"BandWidth,omitnil,omitempty" name:"BandWidth"`

	// Instance disk size. default value is 500. step length is set to 100. should meet the billing specification of the current instance. can be accessed through the following link: https://www.tencentcloud.comom/document/product/597/122562.?from_cn_redirect=1
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// Specifies the maximum number of partitions for the instance, which should meet the billing specification of the current instance. default value is 800 with a step length of 100. the billing specification can be viewed through the following link: https://www.tencentcloud.comom/document/product/597/122563.?from_cn_redirect=1
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// Maximum number of topics for the instance should meet the billing specification of the current instance. default value is 800, step length is set to 100.
	TopicNum *int64 `json:"TopicNum,omitnil,omitempty" name:"TopicNum"`

	// Specifies the availability zone of the instance. when creating a multi-az instance, this parameter is the availability zone id of the subnet where the default access point is created. ZoneId and ZoneIds cannot be empty at the same time. obtain through the API [DescribeCkafkaZone](https://www.tencentcloud.comom/document/product/597/55246?from_cn_redirect=1).
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Whether the current instance is a multi-AZ instance
	MultiZoneFlag *bool `json:"MultiZoneFlag,omitnil,omitempty" name:"MultiZoneFlag"`

	// Specifies the multi-az id list when the instance is a multi-az instance. note that the multi-az corresponding to parameter ZoneId must be included in this parameter array. ZoneId and ZoneIds cannot be empty at the same time. you can obtain this information through the [DescribeCkafkaZone](https://www.tencentcloud.comom/document/product/597/55246?from_cn_redirect=1) api.
	ZoneIds []*int64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// The number of purchased instances.  Default value: `1`. This parameter is optional.  If it is passed in, multiple instances will be created, with their names being `instanceName` plus different suffixes.
	InstanceNum *int64 `json:"InstanceNum,omitnil,omitempty" name:"InstanceNum"`

	// Public network bandwidth in Mbps.  The 3 Mbps of free bandwidth is not included here by default.  For example, if you need 3 Mbps of public network bandwidth, pass in `0`; if you need 6 Mbps, pass in `3`.  The value must be an integer multiple of 3.
	PublicNetworkMonthly *int64 `json:"PublicNetworkMonthly,omitnil,omitempty" name:"PublicNetworkMonthly"`

	// Tag.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Elastic bandwidth switch. valid values: 0 (disable, default), 1 (enable).
	ElasticBandwidthSwitch *int64 `json:"ElasticBandwidthSwitch,omitnil,omitempty" name:"ElasticBandwidthSwitch"`
}

type CreatePostPaidInstanceRequest struct {
	*tchttp.BaseRequest
	
	// VPC Id, obtain through the API [DescribeVpcs](https://www.tencentcloud.comom/document/product/215/15778?from_cn_redirect=1).
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet Id. can be obtained through the [DescribeSubnets](https://www.tencentcloud.comom/document/product/215/15784?from_cn_redirect=1) api.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Specifies the cluster instance name of ckafka, an arbitrary character with length not exceeding 128.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Specifies the standard edition instance specification for the international site. currently only the international site standard edition uses the current field to distinguish specifications, while the domestic site standard edition distinguishes specifications by peak bandwidth. fill in 1 for all instances except the international site standard edition. for international site standard edition instances: [entry-level (general)] fill 1; [standard type (standard)] fill 2; [advanced] fill 3; [capacity type (capacity)] fill 4; [advanced type 1 (specialized-1)] fill 5; [advanced type 2 (specialized-2)] fill 6; [advanced type 3 (specialized-3)] fill 7; [advanced type 4 (specialized-4)] fill 8.
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// The maximum instance log retention period in minutes by default.  If this parameter is left empty, the default retention period is 1,440 minutes (1 day) to 30 days.  If the message retention period of the topic is explicitly set, it will prevail.
	MsgRetentionTime *int64 `json:"MsgRetentionTime,omitnil,omitempty" name:"MsgRetentionTime"`

	// Cluster ID, which can be selected when you create an instance.  You don’t need to pass in this parameter if the cluster where the instance resides is not specified.
	ClusterId *int64 `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Instance version. currently supports "2.4.1", "2.4.2", "2.8.1", "3.2.3". default value "2.4.1". "2.4.1" and "2.4.2" belong to the same version. any one can be passed.
	KafkaVersion *string `json:"KafkaVersion,omitnil,omitempty" name:"KafkaVersion"`

	// Instance type. "standard": standard version. "profession": pro edition. (standard version is only supported on the international site. currently, the chinese site supports pro edition.).
	SpecificationsType *string `json:"SpecificationsType,omitnil,omitempty" name:"SpecificationsType"`

	// Specifies the disk type for a pro edition instance. you do not need to fill it in for a standard edition instance. valid values: "CLOUD_SSD" for SSD CLOUD disk; "CLOUD_BASIC" for high-performance CLOUD block storage. default value: "CLOUD_BASIC".
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// Specifies the peak bandwidth of the instance private network, with a default value of 40 MB/s. for standard version, input the peak bandwidth corresponding to the current instance specifications. note that if the instance created is a pro edition instance, parameter configuration such as peak bandwidth and number of partitions should meet the billing specification of the professional edition. view billing specifications through the following link: https://www.tencentcloud.comom/document/product/597/11745.?from_cn_redirect=1
	BandWidth *int64 `json:"BandWidth,omitnil,omitempty" name:"BandWidth"`

	// Instance disk size. default value is 500. step length is set to 100. should meet the billing specification of the current instance. can be accessed through the following link: https://www.tencentcloud.comom/document/product/597/122562.?from_cn_redirect=1
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// Specifies the maximum number of partitions for the instance, which should meet the billing specification of the current instance. default value is 800 with a step length of 100. the billing specification can be viewed through the following link: https://www.tencentcloud.comom/document/product/597/122563.?from_cn_redirect=1
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// Maximum number of topics for the instance should meet the billing specification of the current instance. default value is 800, step length is set to 100.
	TopicNum *int64 `json:"TopicNum,omitnil,omitempty" name:"TopicNum"`

	// Specifies the availability zone of the instance. when creating a multi-az instance, this parameter is the availability zone id of the subnet where the default access point is created. ZoneId and ZoneIds cannot be empty at the same time. obtain through the API [DescribeCkafkaZone](https://www.tencentcloud.comom/document/product/597/55246?from_cn_redirect=1).
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Whether the current instance is a multi-AZ instance
	MultiZoneFlag *bool `json:"MultiZoneFlag,omitnil,omitempty" name:"MultiZoneFlag"`

	// Specifies the multi-az id list when the instance is a multi-az instance. note that the multi-az corresponding to parameter ZoneId must be included in this parameter array. ZoneId and ZoneIds cannot be empty at the same time. you can obtain this information through the [DescribeCkafkaZone](https://www.tencentcloud.comom/document/product/597/55246?from_cn_redirect=1) api.
	ZoneIds []*int64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// The number of purchased instances.  Default value: `1`. This parameter is optional.  If it is passed in, multiple instances will be created, with their names being `instanceName` plus different suffixes.
	InstanceNum *int64 `json:"InstanceNum,omitnil,omitempty" name:"InstanceNum"`

	// Public network bandwidth in Mbps.  The 3 Mbps of free bandwidth is not included here by default.  For example, if you need 3 Mbps of public network bandwidth, pass in `0`; if you need 6 Mbps, pass in `3`.  The value must be an integer multiple of 3.
	PublicNetworkMonthly *int64 `json:"PublicNetworkMonthly,omitnil,omitempty" name:"PublicNetworkMonthly"`

	// Tag.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Elastic bandwidth switch. valid values: 0 (disable, default), 1 (enable).
	ElasticBandwidthSwitch *int64 `json:"ElasticBandwidthSwitch,omitnil,omitempty" name:"ElasticBandwidthSwitch"`
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
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "InstanceName")
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
	delete(f, "Tags")
	delete(f, "ElasticBandwidthSwitch")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePostPaidInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePostPaidInstanceResponseParams struct {
	// Returned result
	Result *CreateInstancePostResp `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
type CreateRouteRequestParams struct {
	// <p>Specifies the ckafka cluster instance id. obtain through the API <a href="https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1">DescribeInstances</a>.</p>.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <P>Specifies the network type of the route (3: vpc routing; 7: internal support route; 1: public network route).</p>.
	VipType *int64 `json:"VipType,omitnil,omitempty" name:"VipType"`

	// <p>vpc network Id. required when vipType is 3.</p>.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>Specifies the vpc subnet id. required when vipType is 3.</p>.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>Access type: 0-plaintext; 1-sasl_plaintext; 3-sasl_ssl; 4-sasl_scram_sha_256; 5-sasl_scram_sha_512. defaults to 0. when vipType=3, supports 0,1,3,4,5. when vipType=7, supports 0,1,3. when vipType=1, supports 1,3.</p>.
	AccessType *int64 `json:"AccessType,omitnil,omitempty" name:"AccessType"`

	// <P>Specifies whether access management is required. this field has been deprecated.</p>.
	AuthFlag *int64 `json:"AuthFlag,omitnil,omitempty" name:"AuthFlag"`

	// <p>Specifies the caller appId.</p>.
	CallerAppid *int64 `json:"CallerAppid,omitnil,omitempty" name:"CallerAppid"`

	// <P>Public network bandwidth. required for public network route. must be a multiple of 3. no default value.</p>.
	PublicNetwork *int64 `json:"PublicNetwork,omitnil,omitempty" name:"PublicNetwork"`

	// <p>vip address.</p>.
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// <P>Specifies the remark information.</p>.
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// <P>Specifies the ordered list of security group associations.</p>.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

type CreateRouteRequest struct {
	*tchttp.BaseRequest
	
	// <p>Specifies the ckafka cluster instance id. obtain through the API <a href="https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1">DescribeInstances</a>.</p>.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <P>Specifies the network type of the route (3: vpc routing; 7: internal support route; 1: public network route).</p>.
	VipType *int64 `json:"VipType,omitnil,omitempty" name:"VipType"`

	// <p>vpc network Id. required when vipType is 3.</p>.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>Specifies the vpc subnet id. required when vipType is 3.</p>.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>Access type: 0-plaintext; 1-sasl_plaintext; 3-sasl_ssl; 4-sasl_scram_sha_256; 5-sasl_scram_sha_512. defaults to 0. when vipType=3, supports 0,1,3,4,5. when vipType=7, supports 0,1,3. when vipType=1, supports 1,3.</p>.
	AccessType *int64 `json:"AccessType,omitnil,omitempty" name:"AccessType"`

	// <P>Specifies whether access management is required. this field has been deprecated.</p>.
	AuthFlag *int64 `json:"AuthFlag,omitnil,omitempty" name:"AuthFlag"`

	// <p>Specifies the caller appId.</p>.
	CallerAppid *int64 `json:"CallerAppid,omitnil,omitempty" name:"CallerAppid"`

	// <P>Public network bandwidth. required for public network route. must be a multiple of 3. no default value.</p>.
	PublicNetwork *int64 `json:"PublicNetwork,omitnil,omitempty" name:"PublicNetwork"`

	// <p>vip address.</p>.
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// <P>Specifies the remark information.</p>.
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// <P>Specifies the ordered list of security group associations.</p>.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

func (r *CreateRouteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRouteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "VipType")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "AccessType")
	delete(f, "AuthFlag")
	delete(f, "CallerAppid")
	delete(f, "PublicNetwork")
	delete(f, "Ip")
	delete(f, "Note")
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRouteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRouteResponseParams struct {
	// <P>Returned result.</p>.
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRouteResponse struct {
	*tchttp.BaseResponse
	Response *CreateRouteResponseParams `json:"Response"`
}

func (r *CreateRouteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTopicIpWhiteListRequestParams struct {
	// The ckafka cluster instance Id, which can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Specifies the topic name, which can be obtained through the [DescribeTopic](https://www.tencentcloud.comom/document/product/597/40847?from_cn_redirect=1) api.
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Allowlist list. maximum value is 512. upper limit for incoming ips is 512.
	IpWhiteList []*string `json:"IpWhiteList,omitnil,omitempty" name:"IpWhiteList"`
}

type CreateTopicIpWhiteListRequest struct {
	*tchttp.BaseRequest
	
	// The ckafka cluster instance Id, which can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Specifies the topic name, which can be obtained through the [DescribeTopic](https://www.tencentcloud.comom/document/product/597/40847?from_cn_redirect=1) api.
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Allowlist list. maximum value is 512. upper limit for incoming ips is 512.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// Instance Id. you can obtain it by calling the DescribeInstances api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Can only contain letters, digits, underscores, "-", or ".".
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

	// Topic remark is a string of no more than 64 characters. the first character can be a letter or digit, and the remaining part can contain letters, digits, and hyphens (-).
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// Minimum number of synchronous replicas, defaults to 1.
	MinInsyncReplicas *int64 `json:"MinInsyncReplicas,omitnil,omitempty" name:"MinInsyncReplicas"`

	// Whether to allow unsynchronized replicas to be elected as leader. valid values: 0 (not allowed), 1 (allowed). default: not allowed.
	UncleanLeaderElectionEnable *int64 `json:"UncleanLeaderElectionEnable,omitnil,omitempty" name:"UncleanLeaderElectionEnable"`

	// Optional parameter. specifies the message retention period in milliseconds. current min value is 60000. default value is 7200000 ms (2 hours). maximum value is 7776000000 ms (90 days).
	RetentionMs *int64 `json:"RetentionMs,omitnil,omitempty" name:"RetentionMs"`

	// Duration of Segment shard scrolling in milliseconds. minimum value is 86400000 ms (1 day).
	SegmentMs *int64 `json:"SegmentMs,omitnil,omitempty" name:"SegmentMs"`

	// Maximum topic messages in Bytes. value range: 1024 (1 KB) to 12582912 (12 MB).
	MaxMessageBytes *int64 `json:"MaxMessageBytes,omitnil,omitempty" name:"MaxMessageBytes"`

	// Preset ACL rule. `1`: enable, `0`: disable. Default value: `0`.
	EnableAclRule *int64 `json:"EnableAclRule,omitnil,omitempty" name:"EnableAclRule"`

	// Name of the preset ACL rule.
	AclRuleName *string `json:"AclRuleName,omitnil,omitempty" name:"AclRuleName"`

	// Optional. retain file size. defaults to -1, unit Byte. current min value is 1073741824.
	RetentionBytes *int64 `json:"RetentionBytes,omitnil,omitempty" name:"RetentionBytes"`

	// Tag list.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Time type for message saving. valid values: CreateTime/LogAppendTime.
	LogMsgTimestampType *string `json:"LogMsgTimestampType,omitnil,omitempty" name:"LogMsgTimestampType"`
}

type CreateTopicRequest struct {
	*tchttp.BaseRequest
	
	// Instance Id. you can obtain it by calling the DescribeInstances api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Can only contain letters, digits, underscores, "-", or ".".
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

	// Topic remark is a string of no more than 64 characters. the first character can be a letter or digit, and the remaining part can contain letters, digits, and hyphens (-).
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// Minimum number of synchronous replicas, defaults to 1.
	MinInsyncReplicas *int64 `json:"MinInsyncReplicas,omitnil,omitempty" name:"MinInsyncReplicas"`

	// Whether to allow unsynchronized replicas to be elected as leader. valid values: 0 (not allowed), 1 (allowed). default: not allowed.
	UncleanLeaderElectionEnable *int64 `json:"UncleanLeaderElectionEnable,omitnil,omitempty" name:"UncleanLeaderElectionEnable"`

	// Optional parameter. specifies the message retention period in milliseconds. current min value is 60000. default value is 7200000 ms (2 hours). maximum value is 7776000000 ms (90 days).
	RetentionMs *int64 `json:"RetentionMs,omitnil,omitempty" name:"RetentionMs"`

	// Duration of Segment shard scrolling in milliseconds. minimum value is 86400000 ms (1 day).
	SegmentMs *int64 `json:"SegmentMs,omitnil,omitempty" name:"SegmentMs"`

	// Maximum topic messages in Bytes. value range: 1024 (1 KB) to 12582912 (12 MB).
	MaxMessageBytes *int64 `json:"MaxMessageBytes,omitnil,omitempty" name:"MaxMessageBytes"`

	// Preset ACL rule. `1`: enable, `0`: disable. Default value: `0`.
	EnableAclRule *int64 `json:"EnableAclRule,omitnil,omitempty" name:"EnableAclRule"`

	// Name of the preset ACL rule.
	AclRuleName *string `json:"AclRuleName,omitnil,omitempty" name:"AclRuleName"`

	// Optional. retain file size. defaults to -1, unit Byte. current min value is 1073741824.
	RetentionBytes *int64 `json:"RetentionBytes,omitnil,omitempty" name:"RetentionBytes"`

	// Tag list.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Time type for message saving. valid values: CreateTime/LogAppendTime.
	LogMsgTimestampType *string `json:"LogMsgTimestampType,omitnil,omitempty" name:"LogMsgTimestampType"`
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
	delete(f, "LogMsgTimestampType")
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// The ckafka cluster instance Id, which can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Username
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// User password
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

type CreateUserRequest struct {
	*tchttp.BaseRequest
	
	// The ckafka cluster instance Id, which can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
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
	// Returned result.
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

type CvmAndIpInfo struct {
	// The ckafka cluster instance Id.
	CkafkaInstanceId *string `json:"CkafkaInstanceId,omitnil,omitempty" name:"CkafkaInstanceId"`

	// CVM instance ID (ins-test) or POD IP (10.0.0.30).
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// IP address.
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`
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

	// Expiration time in milliseconds.
	RetentionMs *uint64 `json:"RetentionMs,omitnil,omitempty" name:"RetentionMs"`

	// Remarks
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// Status (`1`: In use; `2`: Deleting)
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type DatahubTopicResp struct {
	// Topic name
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Topic Id.
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

type DealInstanceDTO struct {
	// Order transaction.
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// Order transaction corresponds to the list of purchased CKafka instance ids.
	InstanceIdList []*string `json:"InstanceIdList,omitnil,omitempty" name:"InstanceIdList"`
}

// Predefined struct for user
type DeleteAclRequestParams struct {
	// The ckafka cluster instance Id, which can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ACL resource type (`2`: TOPIC, `3`: GROUP, `4`: CLUSTER).
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// Resource name, which is related to `resourceType`. For example, if `resourceType` is `TOPIC`, this field indicates the topic name; if `resourceType` is `GROUP`, this field indicates the group name; if `resourceType` is `CLUSTER`, this field can be left empty.
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// ACL operation type (`2`: ALL, `3`: READ, `4`: WRITE, `5`: CREATE, `6`: DELETE, `7`: ALTER, `8`: DESCRIBE, `9`: CLUSTER_ACTION, `10`: DESCRIBE_CONFIGS, `11`: ALTER_CONFIGS, `12`: IDEMPOTENT_WRITE).
	Operation *int64 `json:"Operation,omitnil,omitempty" name:"Operation"`

	// Permission type (2:DENY, 3:ALLOW). currently ckafka supports ALLOW (equivalent to allowlist), others used when compatible with open-source kafka acl.
	PermissionType *int64 `json:"PermissionType,omitnil,omitempty" name:"PermissionType"`

	// The default value is `*`, which means that any host can access. Currently, CKafka does not support the host as `*`, but the future product based on the open-source Kafka will directly support this
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// List of users, defaults to User:*, means any User is accessible in the entire region. the current User can only be the User in the list of users.
	Principal *string `json:"Principal,omitnil,omitempty" name:"Principal"`
}

type DeleteAclRequest struct {
	*tchttp.BaseRequest
	
	// The ckafka cluster instance Id, which can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ACL resource type (`2`: TOPIC, `3`: GROUP, `4`: CLUSTER).
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// Resource name, which is related to `resourceType`. For example, if `resourceType` is `TOPIC`, this field indicates the topic name; if `resourceType` is `GROUP`, this field indicates the group name; if `resourceType` is `CLUSTER`, this field can be left empty.
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// ACL operation type (`2`: ALL, `3`: READ, `4`: WRITE, `5`: CREATE, `6`: DELETE, `7`: ALTER, `8`: DESCRIBE, `9`: CLUSTER_ACTION, `10`: DESCRIBE_CONFIGS, `11`: ALTER_CONFIGS, `12`: IDEMPOTENT_WRITE).
	Operation *int64 `json:"Operation,omitnil,omitempty" name:"Operation"`

	// Permission type (2:DENY, 3:ALLOW). currently ckafka supports ALLOW (equivalent to allowlist), others used when compatible with open-source kafka acl.
	PermissionType *int64 `json:"PermissionType,omitnil,omitempty" name:"PermissionType"`

	// The default value is `*`, which means that any host can access. Currently, CKafka does not support the host as `*`, but the future product based on the open-source Kafka will directly support this
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// List of users, defaults to User:*, means any User is accessible in the entire region. the current User can only be the User in the list of users.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
type DeleteAclRuleRequestParams struct {
	// Instance id information. you can obtain it through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// acl rule name, obtain through the API DescribeAclRule.
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`
}

type DeleteAclRuleRequest struct {
	*tchttp.BaseRequest
	
	// Instance id information. you can obtain it through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// acl rule name, obtain through the API DescribeAclRule.
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`
}

func (r *DeleteAclRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAclRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "RuleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAclRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAclRuleResponseParams struct {
	// Returns the rule ID of the deleted rule.
	Result *int64 `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAclRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAclRuleResponseParams `json:"Response"`
}

func (r *DeleteAclRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAclRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGroupRequestParams struct {
	// The ckafka cluster instance Id. can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Consumer group name, which can be obtained through the DescribeConsumerGroup API (https://www.tencentcloud.comom/document/product/597/40841?from_cn_redirect=1).
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`
}

type DeleteGroupRequest struct {
	*tchttp.BaseRequest
	
	// The ckafka cluster instance Id. can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Consumer group name, which can be obtained through the DescribeConsumerGroup API (https://www.tencentcloud.comom/document/product/597/40841?from_cn_redirect=1).
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`
}

func (r *DeleteGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Group")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGroupResponseParams struct {
	// Returned results.
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteGroupResponseParams `json:"Response"`
}

func (r *DeleteGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstancePostRequestParams struct {
	// The ckafka cluster instance Id, which can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DeleteInstancePostRequest struct {
	*tchttp.BaseRequest
	
	// The ckafka cluster instance Id, which can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DeleteInstancePostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstancePostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteInstancePostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstancePostResponseParams struct {
	// Returned result.
	Result *InstanceDeleteResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteInstancePostResponse struct {
	*tchttp.BaseResponse
	Response *DeleteInstancePostResponseParams `json:"Response"`
}

func (r *DeleteInstancePostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstancePostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstancePreRequestParams struct {
	// The ckafka cluster instance Id, which can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DeleteInstancePreRequest struct {
	*tchttp.BaseRequest
	
	// The ckafka cluster instance Id, which can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// ckafka cluster instance Id. obtain through the API [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1).
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Route id, obtain through the API [DescribeRoute](https://www.tencentcloud.comom/document/product/597/45484?from_cn_redirect=1).
	RouteId *int64 `json:"RouteId,omitnil,omitempty" name:"RouteId"`

	// AppId of the caller.
	CallerAppid *int64 `json:"CallerAppid,omitnil,omitempty" name:"CallerAppid"`

	// Sets the scheduled deletion time for routes. only public network routes support scheduled deletion. available for any time within the next 24 hours.
	DeleteRouteTime *string `json:"DeleteRouteTime,omitnil,omitempty" name:"DeleteRouteTime"`
}

type DeleteRouteRequest struct {
	*tchttp.BaseRequest
	
	// ckafka cluster instance Id. obtain through the API [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1).
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Route id, obtain through the API [DescribeRoute](https://www.tencentcloud.comom/document/product/597/45484?from_cn_redirect=1).
	RouteId *int64 `json:"RouteId,omitnil,omitempty" name:"RouteId"`

	// AppId of the caller.
	CallerAppid *int64 `json:"CallerAppid,omitnil,omitempty" name:"CallerAppid"`

	// Sets the scheduled deletion time for routes. only public network routes support scheduled deletion. available for any time within the next 24 hours.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// ckafka cluster instance Id. obtain through the API [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1).
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Modifies the scheduled time for deleting routes.
	DelayTime *string `json:"DelayTime,omitnil,omitempty" name:"DelayTime"`
}

type DeleteRouteTriggerTimeRequest struct {
	*tchttp.BaseRequest
	
	// ckafka cluster instance Id. obtain through the API [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1).
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Modifies the scheduled time for deleting routes.
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
	delete(f, "InstanceId")
	delete(f, "DelayTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRouteTriggerTimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRouteTriggerTimeResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// The ckafka cluster instance Id, which can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Specifies the topic name, which can be obtained through the [DescribeTopic](https://www.tencentcloud.comom/document/product/597/40847?from_cn_redirect=1) api.
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// IP allowlist list
	IpWhiteList []*string `json:"IpWhiteList,omitnil,omitempty" name:"IpWhiteList"`
}

type DeleteTopicIpWhiteListRequest struct {
	*tchttp.BaseRequest
	
	// The ckafka cluster instance Id, which can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Specifies the topic name, which can be obtained through the [DescribeTopic](https://www.tencentcloud.comom/document/product/597/40847?from_cn_redirect=1) api.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// The ckafka cluster instance Id, which can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Specifies the username, which can be obtained through the [DescribeUser](https://www.tencentcloud.comom/document/product/597/40855?from_cn_redirect=1) api.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DeleteUserRequest struct {
	*tchttp.BaseRequest
	
	// The ckafka cluster instance Id, which can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Specifies the username, which can be obtained through the [DescribeUser](https://www.tencentcloud.comom/document/product/597/40855?from_cn_redirect=1) api.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// The ckafka cluster instance Id, which can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ACL resource type (`2`: TOPIC, `3`: GROUP, `4`: CLUSTER).
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// Resource name, which is related to `resourceType`. For example, if `resourceType` is `TOPIC`, this field indicates the topic name; if `resourceType` is `GROUP`, this field indicates the group name; if `resourceType` is `CLUSTER`, this field can be left empty.
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// Offset position
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number limit. default value is 50. maximum value is 50.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Keyword match
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`
}

type DescribeACLRequest struct {
	*tchttp.BaseRequest
	
	// The ckafka cluster instance Id, which can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ACL resource type (`2`: TOPIC, `3`: GROUP, `4`: CLUSTER).
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// Resource name, which is related to `resourceType`. For example, if `resourceType` is `TOPIC`, this field indicates the topic name; if `resourceType` is `GROUP`, this field indicates the group name; if `resourceType` is `CLUSTER`, this field can be left empty.
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// Offset position
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number limit. default value is 50. maximum value is 50.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// The ckafka cluster instance Id, which can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ACL rule name
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// ACL rule-based matching type (PREFIXED: prefix match, PRESET: PRESET policy).
	PatternType *string `json:"PatternType,omitnil,omitempty" name:"PatternType"`

	// Specifies whether to read the simplified ACL rule. default value is false, which means not to read the simplified ACL rule.
	IsSimplified *bool `json:"IsSimplified,omitnil,omitempty" name:"IsSimplified"`
}

type DescribeAclRuleRequest struct {
	*tchttp.BaseRequest
	
	// The ckafka cluster instance Id, which can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ACL rule name
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// ACL rule-based matching type (PREFIXED: prefix match, PRESET: PRESET policy).
	PatternType *string `json:"PatternType,omitnil,omitempty" name:"PatternType"`

	// Specifies whether to read the simplified ACL rule. default value is false, which means not to read the simplified ACL rule.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
type DescribeCkafkaVersionRequestParams struct {
	// The ckafka cluster instance Id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeCkafkaVersionRequest struct {
	*tchttp.BaseRequest
	
	// The ckafka cluster instance Id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeCkafkaVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCkafkaVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCkafkaVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCkafkaVersionResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCkafkaVersionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCkafkaVersionResponseParams `json:"Response"`
}

func (r *DescribeCkafkaVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCkafkaVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCkafkaZoneRequestParams struct {
	// cdc cluster Id.
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`
}

type DescribeCkafkaZoneRequest struct {
	*tchttp.BaseRequest
	
	// cdc cluster Id.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// Specifies whether supported versions are required or not.
	IsUnSupportVersion *bool `json:"IsUnSupportVersion,omitnil,omitempty" name:"IsUnSupportVersion"`
}

// Predefined struct for user
type DescribeConsumerGroupRequestParams struct {
	// ckafka cluster instance Id. obtain through the API [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1).
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Specifies the group name you want to query.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Specifies the corresponding topic name in the group to be queried by the user. if this parameter is specified while the group is unspecified, ignore this parameter.
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Returns the limit quantity of the consumption group. supports a maximum of 50.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Specifies the starting offset amount of the consumer group list.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeConsumerGroupRequest struct {
	*tchttp.BaseRequest
	
	// ckafka cluster instance Id. obtain through the API [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1).
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Specifies the group name you want to query.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Specifies the corresponding topic name in the group to be queried by the user. if this parameter is specified while the group is unspecified, ignore this parameter.
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Returns the limit quantity of the consumption group. supports a maximum of 50.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Specifies the starting offset amount of the consumer group list.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
type DescribeCvmInfoRequestParams struct {
	// ckafka cluster instance Id. obtain through the API [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1).
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeCvmInfoRequest struct {
	*tchttp.BaseRequest
	
	// ckafka cluster instance Id. obtain through the API [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1).
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeCvmInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCvmInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCvmInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCvmInfoResponseParams struct {
	// Returned result.
	Result *ListCvmAndIpInfoRsp `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCvmInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCvmInfoResponseParams `json:"Response"`
}

func (r *DescribeCvmInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCvmInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatahubTopicRequestParams struct {
	// Elastic topic name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DescribeDatahubTopicRequest struct {
	*tchttp.BaseRequest
	
	// Elastic topic name.
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

	// Expiration time in milliseconds.
	RetentionMs *uint64 `json:"RetentionMs,omitnil,omitempty" name:"RetentionMs"`

	// Remarks.
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// Username
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Password
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Status (`1`: In use; `2`: Deleting)
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Specifies the service routing address.
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`
}

// Predefined struct for user
type DescribeDatahubTopicResponseParams struct {
	// Returned result object
	Result *DescribeDatahubTopicResp `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// Search term.
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// Query offset, which defaults to `0`.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of results to be returned in this request. Default value: `50`. Maximum value: `50`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Specifies whether to query the topic list from the connection.
	QueryFromConnectResource *bool `json:"QueryFromConnectResource,omitnil,omitempty" name:"QueryFromConnectResource"`

	// Connection ID.
	ConnectResourceId *string `json:"ConnectResourceId,omitnil,omitempty" name:"ConnectResourceId"`

	// topic resource expression.
	TopicRegularExpression *string `json:"TopicRegularExpression,omitnil,omitempty" name:"TopicRegularExpression"`
}

type DescribeDatahubTopicsRequest struct {
	*tchttp.BaseRequest
	
	// Search term.
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// Query offset, which defaults to `0`.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of results to be returned in this request. Default value: `50`. Maximum value: `50`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Specifies whether to query the topic list from the connection.
	QueryFromConnectResource *bool `json:"QueryFromConnectResource,omitnil,omitempty" name:"QueryFromConnectResource"`

	// Connection ID.
	ConnectResourceId *string `json:"ConnectResourceId,omitnil,omitempty" name:"ConnectResourceId"`

	// topic resource expression.
	TopicRegularExpression *string `json:"TopicRegularExpression,omitnil,omitempty" name:"TopicRegularExpression"`
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
	delete(f, "QueryFromConnectResource")
	delete(f, "ConnectResourceId")
	delete(f, "TopicRegularExpression")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatahubTopicsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDatahubTopicsResp struct {
	// Total count
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Topic list.
	TopicList []*DatahubTopicDTO `json:"TopicList,omitnil,omitempty" name:"TopicList"`
}

// Predefined struct for user
type DescribeDatahubTopicsResponseParams struct {
	// Topic list.
	Result *DescribeDatahubTopicsResp `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// Consumer group name.
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// Protocol used by the group.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`
}

// Predefined struct for user
type DescribeGroupInfoRequestParams struct {
	// ckafka cluster instance Id. obtain through the API [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1).
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Kafka group list. obtain through the API [DescribeConsumerGroup](https://www.tencentcloud.comom/document/product/597/40841?from_cn_redirect=1).
	GroupList []*string `json:"GroupList,omitnil,omitempty" name:"GroupList"`
}

type DescribeGroupInfoRequest struct {
	*tchttp.BaseRequest
	
	// ckafka cluster instance Id. obtain through the API [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1).
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Kafka group list. obtain through the API [DescribeConsumerGroup](https://www.tencentcloud.comom/document/product/597/40841?from_cn_redirect=1).
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
	// Returned result.
	Result []*GroupInfoResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// The ckafka cluster instance Id.
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
	
	// The ckafka cluster instance Id.
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
	// Returned result.
	Result *GroupOffsetResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// The ckafka cluster instance Id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Search keyword
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of results to be returned
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Only supported for GroupState filter criteria. valid values: Empty, Stable. note: this parameter can only be accessed in versions 2.8/3.2.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeGroupRequest struct {
	*tchttp.BaseRequest
	
	// The ckafka cluster instance Id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Search keyword
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of results to be returned
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Only supported for GroupState filter criteria. valid values: Empty, Stable. note: this parameter can only be accessed in versions 2.8/3.2.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGroupResponseParams struct {
	// Returned result.
	Result *GroupResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// ckafka cluster instance Id. obtain through the API [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1).
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceAttributesRequest struct {
	*tchttp.BaseRequest
	
	// ckafka cluster instance Id. obtain through the API [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1).
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	//
	// Deprecated: InstanceIds is deprecated.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// (Query condition) filter by the ckafka cluster instance Id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Search term. example: (query condition) filter by instance name. fuzzy query is supported.
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// Instance status (query condition). valid values: 0: creating, 1: running, 2: deleting, 5: isolated, 7: upgrading. default return: all.
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Offset. If this parameter is left empty, 0 will be used by default
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results to be returned. If this parameter is left empty, 10 will be used by default. The maximum value is 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Tag key value (this field has been deprecated).
	//
	// Deprecated: TagKey is deprecated.
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// (Query condition) VPC Id.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// (Query condition) filter by the ckafka cluster instance Id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Search term. example: (query condition) filter by instance name. fuzzy query is supported.
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// Instance status (query condition). valid values: 0: creating, 1: running, 2: deleting, 5: isolated, 7: upgrading. default return: all.
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Offset. If this parameter is left empty, 0 will be used by default
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results to be returned. If this parameter is left empty, 10 will be used by default. The maximum value is 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Tag key value (this field has been deprecated).
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// (Query condition) VPC Id.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// Returns the region enumeration result list.
	Result []*Region `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// The ckafka cluster instance Id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Route ID
	RouteId *int64 `json:"RouteId,omitnil,omitempty" name:"RouteId"`

	// Specifies whether to display the primary route. when true, the routing list will additionally display the primary route information during instance creation (not affected by InternalFlag or UsedFor parameter filtering).	
	MainRouteFlag *bool `json:"MainRouteFlag,omitnil,omitempty" name:"MainRouteFlag"`
}

type DescribeRouteRequest struct {
	*tchttp.BaseRequest
	
	// The ckafka cluster instance Id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Route ID
	RouteId *int64 `json:"RouteId,omitnil,omitempty" name:"RouteId"`

	// Specifies whether to display the primary route. when true, the routing list will additionally display the primary route information during instance creation (not affected by InternalFlag or UsedFor parameter filtering).	
	MainRouteFlag *bool `json:"MainRouteFlag,omitnil,omitempty" name:"MainRouteFlag"`
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
	delete(f, "MainRouteFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRouteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRouteResponseParams struct {
	// Returned result set of route information
	Result *RouteResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
type DescribeSecurityGroupRoutesRequestParams struct {
	// Specifies the routing information.
	InstanceRoute *InstanceRoute `json:"InstanceRoute,omitnil,omitempty" name:"InstanceRoute"`

	// Filter.
	Filters []*RouteFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Specifies the pagination Offset. default is 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Pagination Limit. default: 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Keyword. specifies fuzzy search by instance id, instance name, or vip.
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`
}

type DescribeSecurityGroupRoutesRequest struct {
	*tchttp.BaseRequest
	
	// Specifies the routing information.
	InstanceRoute *InstanceRoute `json:"InstanceRoute,omitnil,omitempty" name:"InstanceRoute"`

	// Filter.
	Filters []*RouteFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Specifies the pagination Offset. default is 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Pagination Limit. default: 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Keyword. specifies fuzzy search by instance id, instance name, or vip.
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`
}

func (r *DescribeSecurityGroupRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceRoute")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchWord")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityGroupRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupRoutesResponseParams struct {
	// Returns the security group routing information result object.
	Result *SecurityGroupRouteResp `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSecurityGroupRoutesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityGroupRoutesResponseParams `json:"Response"`
}

func (r *DescribeSecurityGroupRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskStatusRequestParams struct {
	// Flow ID.
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`
}

type DescribeTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// Flow ID.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// The ckafka cluster instance Id, which can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Specifies the topic name, which can be obtained through the [DescribeTopic](https://www.tencentcloud.comom/document/product/597/40847?from_cn_redirect=1) api.
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

type DescribeTopicAttributesRequest struct {
	*tchttp.BaseRequest
	
	// The ckafka cluster instance Id, which can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Specifies the topic name, which can be obtained through the [DescribeTopic](https://www.tencentcloud.comom/document/product/597/40847?from_cn_redirect=1) api.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// The ckafka cluster instance Id, which can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// (Filter) filter by `topicName`. Fuzzy search is supported
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// Offset. If this parameter is left empty, 0 will be used by default
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. default: 20. value must be above 0.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Name of the preset ACL rule.
	AclRuleName *string `json:"AclRuleName,omitnil,omitempty" name:"AclRuleName"`

	// Sorts based on specific attributes (currently supports PartitionNum/CreateTime). default value: CreateTime.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 0 - sequential, 1 - reverse order. default value: 0.
	OrderType *int64 `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// Currently supports ReplicaNum (number of replicas) filter criteria.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeTopicDetailRequest struct {
	*tchttp.BaseRequest
	
	// The ckafka cluster instance Id, which can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// (Filter) filter by `topicName`. Fuzzy search is supported
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// Offset. If this parameter is left empty, 0 will be used by default
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. default: 20. value must be above 0.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Name of the preset ACL rule.
	AclRuleName *string `json:"AclRuleName,omitnil,omitempty" name:"AclRuleName"`

	// Sorts based on specific attributes (currently supports PartitionNum/CreateTime). default value: CreateTime.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 0 - sequential, 1 - reverse order. default value: 0.
	OrderType *int64 `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// Currently supports ReplicaNum (number of replicas) filter criteria.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	delete(f, "OrderBy")
	delete(f, "OrderType")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopicDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicDetailResponseParams struct {
	// Returned entity of topic details
	Result *TopicDetailResponse `json:"Result,omitnil,omitempty" name:"Result"`

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

// Predefined struct for user
type DescribeTopicProduceConnectionRequestParams struct {
	// The ckafka cluster instance Id, which can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Specifies the topic name, which can be obtained through the [DescribeTopic](https://www.tencentcloud.comom/document/product/597/40847?from_cn_redirect=1) api.
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

type DescribeTopicProduceConnectionRequest struct {
	*tchttp.BaseRequest
	
	// The ckafka cluster instance Id, which can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Specifies the topic name, which can be obtained through the [DescribeTopic](https://www.tencentcloud.comom/document/product/597/40847?from_cn_redirect=1) api.
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
	// The ckafka cluster instance Id.
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
	
	// The ckafka cluster instance Id.
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
	// Returned result.
	Result *TopicResult `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// The ckafka cluster instance Id.
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
	
	// The ckafka cluster instance Id.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// Number of returned results. default value: 20. must be greater than 0.
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

	// Number of returned results. default value: 20. must be greater than 0.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
type DescribeTypeInstancesRequestParams struct {
	// (Filter condition) filter by instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// (Filter condition) filter by instance name. fuzzy query is supported.
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// Instance status (filter condition). valid values: 0: creating, 1: running, 2: deleting. default return: all.
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Offset. default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. default: 10. maximum value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Matches the Tag key.
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`
}

type DescribeTypeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// (Filter condition) filter by instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// (Filter condition) filter by instance name. fuzzy query is supported.
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// Instance status (filter condition). valid values: 0: creating, 1: running, 2: deleting. default return: all.
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Offset. default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. default: 10. maximum value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Matches the Tag key.
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`
}

func (r *DescribeTypeInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTypeInstancesRequest) FromJsonString(s string) error {
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTypeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTypeInstancesResponseParams struct {
	// Returned result.
	Result *InstanceResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTypeInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTypeInstancesResponseParams `json:"Response"`
}

func (r *DescribeTypeInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTypeInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserRequestParams struct {
	// The ckafka cluster instance Id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Filter by name
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of returns.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeUserRequest struct {
	*tchttp.BaseRequest
	
	// The ckafka cluster instance Id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Filter by name
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of returns.
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
	// Returned result.
	Result *UserResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// The ckafka cluster instance Id, which can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Specifies the topic name, which can be obtained through the [DescribeTopic](https://www.tencentcloud.comom/document/product/597/40847?from_cn_redirect=1) api.
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// Partition ID
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// Specifies the position information.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type FetchMessageByOffsetRequest struct {
	*tchttp.BaseRequest
	
	// The ckafka cluster instance Id, which can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Specifies the topic name, which can be obtained through the [DescribeTopic](https://www.tencentcloud.comom/document/product/597/40847?from_cn_redirect=1) api.
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// Partition ID
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// Specifies the position information.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// The ckafka cluster instance Id.
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
	
	// The ckafka cluster instance Id.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

// Predefined struct for user
type FetchMessageListByTimestampRequestParams struct {
	// The ckafka cluster instance Id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Topic name
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// Partition ID
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// Query start time, a timestamp.
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Maximum number of query results. default: 20. value range: 1-20.
	SinglePartitionRecordNumber *int64 `json:"SinglePartitionRecordNumber,omitnil,omitempty" name:"SinglePartitionRecordNumber"`
}

type FetchMessageListByTimestampRequest struct {
	*tchttp.BaseRequest
	
	// The ckafka cluster instance Id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Topic name
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// Partition ID
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// Query start time, a timestamp.
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Maximum number of query results. default: 20. value range: 1-20.
	SinglePartitionRecordNumber *int64 `json:"SinglePartitionRecordNumber,omitnil,omitempty" name:"SinglePartitionRecordNumber"`
}

func (r *FetchMessageListByTimestampRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FetchMessageListByTimestampRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Topic")
	delete(f, "Partition")
	delete(f, "StartTime")
	delete(f, "SinglePartitionRecordNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FetchMessageListByTimestampRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FetchMessageListByTimestampResponseParams struct {
	// Returned results. note that the list does not return specific message content (key, value). if necessary, please use the FetchMessageByOffset API to query specific message content.
	Result []*ConsumerRecord `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type FetchMessageListByTimestampResponse struct {
	*tchttp.BaseResponse
	Response *FetchMessageListByTimestampResponseParams `json:"Response"`
}

func (r *FetchMessageListByTimestampResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FetchMessageListByTimestampResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// Field to be filtered.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Filter value of field.
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type Group struct {
	// Consumer group name.
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

	// Consumer group name.
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`
}

type GroupInfoTopics struct {
	// Name of assigned topics
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// Allocates partition info.
	Partitions []*int64 `json:"Partitions,omitnil,omitempty" name:"Partitions"`
}

type GroupOffsetPartition struct {
	// Topic `partitionId`
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// Offset position submitted by consumer
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Supports consumers to submit messages with imported metadata for other purposes, currently an empty string.
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

	// The topic partition array, where each element is a json object.
	TopicList []*GroupOffsetTopic `json:"TopicList,omitnil,omitempty" name:"TopicList"`
}

type GroupOffsetTopic struct {
	// Topic name
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// The topic partition array, where each element is a json object.
	Partitions []*GroupOffsetPartition `json:"Partitions,omitnil,omitempty" name:"Partitions"`
}

type GroupResponse struct {
	// Counting.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// GroupList
	GroupList []*DescribeGroup `json:"GroupList,omitnil,omitempty" name:"GroupList"`

	// Specifies the consumer group quota.
	GroupCountQuota *uint64 `json:"GroupCountQuota,omitnil,omitempty" name:"GroupCountQuota"`
}

// Predefined struct for user
type InquireCkafkaPriceRequestParams struct {
	// Chinese site standard version fill in standards2 international site standard version fill in standard pro edition fill in profession advanced edition fill in premium.
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Billing mode for instance purchase/renewal. If this parameter is left empty when you purchase an instance, the fees for one month under the monthly subscription mode will be displayed by default.
	InstanceChargeParam *InstanceChargeParam `json:"InstanceChargeParam,omitnil,omitempty" name:"InstanceChargeParam"`

	// The number of instances to be purchased or renewed. If this parameter is left empty, the default value is `1`.
	InstanceNum *int64 `json:"InstanceNum,omitnil,omitempty" name:"InstanceNum"`

	// Specifies the internal network bandwidth size of the instance, in MB/s (required when purchased; bandwidth information is required for pro edition/advanced edition inquiries).
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// Specifies the purchase type and size of the hard disk of the instance. required when purchased. disk information is required for pro edition or advanced edition inquiries.
	InquiryDiskParam *InquiryDiskParam `json:"InquiryDiskParam,omitnil,omitempty" name:"InquiryDiskParam"`

	// Message retention period in hours, which is required when you purchase an instance.
	MessageRetention *int64 `json:"MessageRetention,omitnil,omitempty" name:"MessageRetention"`

	// The number of instance topics to be purchased, which is required when you purchase an instance.
	Topic *int64 `json:"Topic,omitnil,omitempty" name:"Topic"`

	// Number of partitions for instance purchase, unit: unit (required when purchased; bandwidth information required for pro edition/advanced edition inquiry).
	// Partition upper limit. maximum value of 40000. step length of 100.
	// Specifies the specifications and limits that can be viewed through the following link: https://www.tencentcloud.comom/document/product/597/122563.?from_cn_redirect=1
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// The region for instance purchase, which can be obtained via the `DescribeCkafkaZone` API.
	ZoneIds []*int64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// Operation type flag. `purchase`: Making new purchases; `renew`: Renewing an instance. The default value is `purchase` if this parameter is left empty.
	CategoryAction *string `json:"CategoryAction,omitnil,omitempty" name:"CategoryAction"`

	// This field is not required.
	BillType *string `json:"BillType,omitnil,omitempty" name:"BillType"`

	// Public network bandwidth billing mode. currently only the pro edition supports public network bandwidth. required when purchasing public network bandwidth. value must be a multiple of 3.
	PublicNetworkParam *InquiryPublicNetworkParam `json:"PublicNetworkParam,omitnil,omitempty" name:"PublicNetworkParam"`

	// ID of the instance to be renewed, which is required when you renew an instance.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type InquireCkafkaPriceRequest struct {
	*tchttp.BaseRequest
	
	// Chinese site standard version fill in standards2 international site standard version fill in standard pro edition fill in profession advanced edition fill in premium.
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Billing mode for instance purchase/renewal. If this parameter is left empty when you purchase an instance, the fees for one month under the monthly subscription mode will be displayed by default.
	InstanceChargeParam *InstanceChargeParam `json:"InstanceChargeParam,omitnil,omitempty" name:"InstanceChargeParam"`

	// The number of instances to be purchased or renewed. If this parameter is left empty, the default value is `1`.
	InstanceNum *int64 `json:"InstanceNum,omitnil,omitempty" name:"InstanceNum"`

	// Specifies the internal network bandwidth size of the instance, in MB/s (required when purchased; bandwidth information is required for pro edition/advanced edition inquiries).
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// Specifies the purchase type and size of the hard disk of the instance. required when purchased. disk information is required for pro edition or advanced edition inquiries.
	InquiryDiskParam *InquiryDiskParam `json:"InquiryDiskParam,omitnil,omitempty" name:"InquiryDiskParam"`

	// Message retention period in hours, which is required when you purchase an instance.
	MessageRetention *int64 `json:"MessageRetention,omitnil,omitempty" name:"MessageRetention"`

	// The number of instance topics to be purchased, which is required when you purchase an instance.
	Topic *int64 `json:"Topic,omitnil,omitempty" name:"Topic"`

	// Number of partitions for instance purchase, unit: unit (required when purchased; bandwidth information required for pro edition/advanced edition inquiry).
	// Partition upper limit. maximum value of 40000. step length of 100.
	// Specifies the specifications and limits that can be viewed through the following link: https://www.tencentcloud.comom/document/product/597/122563.?from_cn_redirect=1
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// The region for instance purchase, which can be obtained via the `DescribeCkafkaZone` API.
	ZoneIds []*int64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// Operation type flag. `purchase`: Making new purchases; `renew`: Renewing an instance. The default value is `purchase` if this parameter is left empty.
	CategoryAction *string `json:"CategoryAction,omitnil,omitempty" name:"CategoryAction"`

	// This field is not required.
	BillType *string `json:"BillType,omitnil,omitempty" name:"BillType"`

	// Public network bandwidth billing mode. currently only the pro edition supports public network bandwidth. required when purchasing public network bandwidth. value must be a multiple of 3.
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
	// Specifies the instance price.
	InstancePrice *InquiryPrice `json:"InstancePrice,omitnil,omitempty" name:"InstancePrice"`

	// Public network bandwidth price
	// Note: This field may return null, indicating that no valid values can be obtained.
	PublicNetworkBandwidthPrice *InquiryPrice `json:"PublicNetworkBandwidthPrice,omitnil,omitempty" name:"PublicNetworkBandwidthPrice"`
}

// Predefined struct for user
type InquireCkafkaPriceResponseParams struct {
	// Returned result.
	Result *InquireCkafkaPriceResp `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// Original price unit.
	UnitPrice *float64 `json:"UnitPrice,omitnil,omitempty" name:"UnitPrice"`

	// Discount unit price.
	UnitPriceDiscount *float64 `json:"UnitPriceDiscount,omitnil,omitempty" name:"UnitPriceDiscount"`

	// Total original price.
	OriginalPrice *float64 `json:"OriginalPrice,omitnil,omitempty" name:"OriginalPrice"`

	// Total discount price.
	DiscountPrice *float64 `json:"DiscountPrice,omitnil,omitempty" name:"DiscountPrice"`

	// Discount (unit: %).
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

	// Purchase quantity.
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

	// Instance package price.
	InstanceTypePrice *InquiryBasePrice `json:"InstanceTypePrice,omitnil,omitempty" name:"InstanceTypePrice"`
}

type InquiryDiskParam struct {
	// Disk type. Valid values: `SSD` (SSD), `CLOUD_SSD` (SSD cloud disk), `CLOUD_PREMIUM` (Premium cloud disk), `CLOUD_BASIC` (Cloud disk).
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// Size of the purchased disk in GB
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`
}

type InquiryPrice struct {
	// Original price unit.
	UnitPrice *float64 `json:"UnitPrice,omitnil,omitempty" name:"UnitPrice"`

	// Discount unit price.
	UnitPriceDiscount *float64 `json:"UnitPriceDiscount,omitnil,omitempty" name:"UnitPriceDiscount"`

	// Total original price.
	OriginalPrice *float64 `json:"OriginalPrice,omitnil,omitempty" name:"OriginalPrice"`

	// Total discount price.
	DiscountPrice *float64 `json:"DiscountPrice,omitnil,omitempty" name:"DiscountPrice"`

	// Discount (unit: %).
	Discount *float64 `json:"Discount,omitnil,omitempty" name:"Discount"`

	// Number of products
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// Specifies the payment currency.
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`

	// Dedicated disk response parameter
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// Purchase duration.
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// Specifies the purchase duration unit ("m" for monthly, "h" for hourly).
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

	// Public network bandwidth, in MB. value must be 0 or a multiple of 3.
	PublicNetworkMonthly *int64 `json:"PublicNetworkMonthly,omitnil,omitempty" name:"PublicNetworkMonthly"`
}

type Instance struct {
	// The ckafka cluster instance Id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Specifies the Name of the ckafka cluster instance.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Instance status. 0: creating, 1: running, 2: deleting, 3: deleted, 5: isolated, 7: upgrading, -1: creation failed.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Specifies whether the instance is open-source. valid values: true (open-source), false (not open-source).
	IfCommunity *bool `json:"IfCommunity,omitnil,omitempty" name:"IfCommunity"`
}

type InstanceAttributesResponse struct {
	// The ckafka cluster instance Id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Specifies the Name of the ckafka cluster instance.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// VIP list information of access point
	VipList []*VipEntity `json:"VipList,omitnil,omitempty" name:"VipList"`

	// Virtual IP
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Virtual port
	Vport *string `json:"Vport,omitnil,omitempty" name:"Vport"`

	// Instance status. 0: creating, 1: running, 2: deleting, 3: deleted, 5: isolated, 7: upgrading, -1: creation failed.
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
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Expiration time
	ExpireTime *uint64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// Availability Zone List
	ZoneIds []*int64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// Specifies the ckafka cluster instance version.
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// Maximum number of groups.
	MaxGroupNum *int64 `json:"MaxGroupNum,omitnil,omitempty" name:"MaxGroupNum"`

	// Sale type. valid values: 0 (standard version), 1 (pro edition).
	Cvm *int64 `json:"Cvm,omitnil,omitempty" name:"Cvm"`

	// Instance type. valid values:. 
	// Specifies the pro edition.    
	// Standard version.
	// premium. specifies the advanced edition.
	// Specifies the serverless version.
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Indicates the characteristics supported by the instance. FEATURE_SUBNET_ACL means the policy support for configuring subnets.
	Features []*string `json:"Features,omitnil,omitempty" name:"Features"`

	// Dynamic message retention policy.
	RetentionTimeConfig *DynamicRetentionTime `json:"RetentionTimeConfig,omitnil,omitempty" name:"RetentionTimeConfig"`

	// Maximum number of connections.
	MaxConnection *uint64 `json:"MaxConnection,omitnil,omitempty" name:"MaxConnection"`

	// Public network bandwidth
	PublicNetwork *int64 `json:"PublicNetwork,omitnil,omitempty" name:"PublicNetwork"`

	// Specifies the deprecated field with no actual meaning.
	DeleteRouteTimestamp *string `json:"DeleteRouteTimestamp,omitnil,omitempty" name:"DeleteRouteTimestamp"`

	// Number of remaining creatable partitions.
	RemainingPartitions *int64 `json:"RemainingPartitions,omitnil,omitempty" name:"RemainingPartitions"`

	// Number of remaining creatable topics.
	RemainingTopics *int64 `json:"RemainingTopics,omitnil,omitempty" name:"RemainingTopics"`

	// Scaling policy for dynamic disk.
	DynamicDiskConfig *DynamicDiskConfig `json:"DynamicDiskConfig,omitnil,omitempty" name:"DynamicDiskConfig"`

	// Specifies the system maintenance time.
	SystemMaintenanceTime *string `json:"SystemMaintenanceTime,omitnil,omitempty" name:"SystemMaintenanceTime"`

	// Specifies the maximum size of messages at the instance level.
	MaxMessageByte *uint64 `json:"MaxMessageByte,omitnil,omitempty" name:"MaxMessageByte"`

	// Specifies the instance billing type. POSTPAID_BY_HOUR: hourly billing; PREPAID: annual/monthly package.
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// Whether to enable the elastic bandwidth allowlist.   
	// Indicates the allowlist feature with elastic bandwidth enabled.
	// 0: elastic bandwidth allowlist feature is disabled.
	ElasticBandwidthSwitch *int64 `json:"ElasticBandwidthSwitch,omitnil,omitempty" name:"ElasticBandwidthSwitch"`

	// Indicates the elastic bandwidth activation status.
	// 1: indicates elastic bandwidth is disabled.
	// Enable elastic bandwidth.
	// Enable elastic bandwidth successfully.
	// 33: disabling elastic bandwidth.
	// Indicates that the elastic bandwidth is successfully disabled.
	// Enable elastic bandwidth failed.
	// Bandwidth failure.
	ElasticBandwidthOpenStatus *int64 `json:"ElasticBandwidthOpenStatus,omitnil,omitempty" name:"ElasticBandwidthOpenStatus"`

	// Cluster type.  
	// CLOUD_IDC idc cluster.
	// CLOUD_CVM_SHARE shared cluster.
	// CLOUD_CVM_YUNTI yunti cvm cluster.
	// CLOUD_CVM. specifies the cvm cluster.
	// CLOUD_CDC cdc cluster.
	// CLOUD_EKS_TSE eks cluster.
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// Number of free partitions.
	FreePartitionNumber *int64 `json:"FreePartitionNumber,omitnil,omitempty" name:"FreePartitionNumber"`

	// Specifies the elastic bandwidth upper limit.
	ElasticFloatBandwidth *int64 `json:"ElasticFloatBandwidth,omitnil,omitempty" name:"ElasticFloatBandwidth"`

	// ssl custom certificate id. only returned for instance clusters with custom certificates.
	CustomCertId *string `json:"CustomCertId,omitnil,omitempty" name:"CustomCertId"`

	// Default unclean.leader.election.enable configuration for cluster topic: 1 enable 0 disable.
	UncleanLeaderElectionEnable *int64 `json:"UncleanLeaderElectionEnable,omitnil,omitempty" name:"UncleanLeaderElectionEnable"`

	// Instance deletion protection switch. 1: enabled; 0: disabled.
	DeleteProtectionEnable *int64 `json:"DeleteProtectionEnable,omitnil,omitempty" name:"DeleteProtectionEnable"`
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

type InstanceDeleteResponse struct {
	// Specifies the task Id returned after deleting an instance.
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`
}

type InstanceDetail struct {
	// The ckafka cluster instance Id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// CKafka cluster instance name.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Instance VIP information
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Instance port information
	Vport *string `json:"Vport,omitnil,omitempty" name:"Vport"`

	// Virtual IP list
	VipList []*VipEntity `json:"VipList,omitnil,omitempty" name:"VipList"`

	// Instance status. 0: creating, 1: running, 2: deleting, 3: deleted, 5: isolated, 7: upgrading, -1: creation failed.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Instance bandwidth in Mbps
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// Specifies the ckafka cluster instance disk size in gb.
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

	// kafka version information.
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// Cross-Availability zone.
	ZoneIds []*int64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// ckafka sales type.
	Cvm *int64 `json:"Cvm,omitnil,omitempty" name:"Cvm"`

	// Specifies the cluster instance type of ckafka.
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Specifies the ckafka cluster instance disk type.
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// Maximum number of topics for current specifications.
	MaxTopicNumber *int64 `json:"MaxTopicNumber,omitnil,omitempty" name:"MaxTopicNumber"`

	// Maximum number of partitions for current specifications.
	MaxPartitionNumber *int64 `json:"MaxPartitionNumber,omitnil,omitempty" name:"MaxPartitionNumber"`

	// Scheduled configuration upgrade time.
	RebalanceTime *string `json:"RebalanceTime,omitnil,omitempty" name:"RebalanceTime"`

	// Specifies the number of partitions in the current instance.
	PartitionNumber *uint64 `json:"PartitionNumber,omitnil,omitempty" name:"PartitionNumber"`

	// Specifies the public network bandwidth type of the ckafka cluster instance.
	PublicNetworkChargeType *string `json:"PublicNetworkChargeType,omitnil,omitempty" name:"PublicNetworkChargeType"`

	// Public network bandwidth. minimum 3 Mbps. maximum 999 Mbps. only the pro edition supports filling in.
	PublicNetwork *int64 `json:"PublicNetwork,omitnil,omitempty" name:"PublicNetwork"`

	// Specifies the underlying cluster type of the ckafka cluster instance.
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// Instance feature list.
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
	// Specifies the list of instances meeting the conditions.
	InstanceList []*Instance `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// Total results that meet the conditions.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type InstanceRoute struct {
	// The ckafka cluster instance Id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Route ID
	RouteId *int64 `json:"RouteId,omitnil,omitempty" name:"RouteId"`
}

// Predefined struct for user
type InstanceScalingDownRequestParams struct {
	// ckafka cluster instance Id. can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Shrink mode. 1: stable mode. 
	// 2. specifies high-speed configuration change.
	UpgradeStrategy *int64 `json:"UpgradeStrategy,omitnil,omitempty" name:"UpgradeStrategy"`

	// Specifies the disk capacity in GB. value range: maximum value 500000, step length 100.
	// The specifications and limitations can be viewed through the following link: https://www.tencentcloud.comom/document/product/597/122562.?from_cn_redirect=1
	// 
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// Peak bandwidth in MB/s.
	// Specifies the url (https://www.tencentcloud.comom/document/product/597/11745?from_cn_redirect=1) to view specification limits and corresponding step length.
	BandWidth *int64 `json:"BandWidth,omitnil,omitempty" name:"BandWidth"`

	// Partition upper limit maximum value of 40000, step length 100.
	// Specification limits can be viewed through the following link: https://www.tencentcloud.comom/document/product/597/122563.?from_cn_redirect=1
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`
}

type InstanceScalingDownRequest struct {
	*tchttp.BaseRequest
	
	// ckafka cluster instance Id. can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Shrink mode. 1: stable mode. 
	// 2. specifies high-speed configuration change.
	UpgradeStrategy *int64 `json:"UpgradeStrategy,omitnil,omitempty" name:"UpgradeStrategy"`

	// Specifies the disk capacity in GB. value range: maximum value 500000, step length 100.
	// The specifications and limitations can be viewed through the following link: https://www.tencentcloud.comom/document/product/597/122562.?from_cn_redirect=1
	// 
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// Peak bandwidth in MB/s.
	// Specifies the url (https://www.tencentcloud.comom/document/product/597/11745?from_cn_redirect=1) to view specification limits and corresponding step length.
	BandWidth *int64 `json:"BandWidth,omitnil,omitempty" name:"BandWidth"`

	// Partition upper limit maximum value of 40000, step length 100.
	// Specification limits can be viewed through the following link: https://www.tencentcloud.comom/document/product/597/122563.?from_cn_redirect=1
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`
}

func (r *InstanceScalingDownRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InstanceScalingDownRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UpgradeStrategy")
	delete(f, "DiskSize")
	delete(f, "BandWidth")
	delete(f, "Partition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InstanceScalingDownRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InstanceScalingDownResponseParams struct {
	// Returned results.
	Result *ScalingDownResp `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InstanceScalingDownResponse struct {
	*tchttp.BaseResponse
	Response *InstanceScalingDownResponseParams `json:"Response"`
}

func (r *InstanceScalingDownResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InstanceScalingDownResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type ListCvmAndIpInfoRsp struct {
	// cvm and IP list.
	CvmList []*CvmAndIpInfo `json:"CvmList,omitnil,omitempty" name:"CvmList"`

	// Specifies the instance data volume.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

// Predefined struct for user
type ModifyAclRuleRequestParams struct {
	// The ckafka cluster instance Id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ACL rule name.
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// Specifies whether to apply to newly-added topics when importing predefined rule modifications.
	IsApplied *int64 `json:"IsApplied,omitnil,omitempty" name:"IsApplied"`
}

type ModifyAclRuleRequest struct {
	*tchttp.BaseRequest
	
	// The ckafka cluster instance Id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ACL rule name.
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// Specifies whether to apply to newly-added topics when importing predefined rule modifications.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// Elastic topic name.
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
	
	// Elastic topic name.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// ckafka cluster instance Id. obtain through the API [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1).
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Consumer group name. obtain through the API [DescribeConsumerGroup](https://www.tencentcloud.comom/document/product/597/40841?from_cn_redirect=1).
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// Reset offset strategy. parameter meaning: 0. align with the shift-by parameter, move the offset forward or backward by shift entries. 1. alignment reference (by-duration, to-datetime, to-earliest, to-latest), move the offset to the specified timestamp position. 2. alignment reference (to-offset), move the offset to the specified offset position.
	Strategy *int64 `json:"Strategy,omitnil,omitempty" name:"Strategy"`

	// Specifies the topic name list that needs to reset.
	Topics []*string `json:"Topics,omitnil,omitempty" name:"Topics"`

	// When `strategy` is 0, this field is required. If it is above zero, the offset will be shifted backward by the value of the `shift`. If it is below zero, the offset will be shifted forward by the value of the `shift`. After a correct reset, the new offset should be (old_offset + shift). Note that if the new offset is smaller than the `earliest` parameter of the partition, it will be set to `earliest`, and if it is greater than the `latest` parameter of the partition, it will be set to `latest`
	Shift *int64 `json:"Shift,omitnil,omitempty" name:"Shift"`

	// In milliseconds. when strategy is 1, must include this field. among them, -2 means reset offset to the start position, -1 means reset to the latest position (equivalent to clearing), other values represent the specified time. obtain the offset at the specified time in the topic and reset. notably, if no message exists at the specified time, get the last offset.
	ShiftTimestamp *int64 `json:"ShiftTimestamp,omitnil,omitempty" name:"ShiftTimestamp"`

	// Position of the offset that needs to be reset. When `strategy` is 2, this field is required
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// List of partitions that need to be reset. If the topics parameter is not specified, reset partitions in the corresponding partition list of all topics. If the topics parameter is specified, reset partitions of the corresponding partition list of the specified topic list.
	Partitions []*int64 `json:"Partitions,omitnil,omitempty" name:"Partitions"`
}

type ModifyGroupOffsetsRequest struct {
	*tchttp.BaseRequest
	
	// ckafka cluster instance Id. obtain through the API [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1).
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Consumer group name. obtain through the API [DescribeConsumerGroup](https://www.tencentcloud.comom/document/product/597/40841?from_cn_redirect=1).
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// Reset offset strategy. parameter meaning: 0. align with the shift-by parameter, move the offset forward or backward by shift entries. 1. alignment reference (by-duration, to-datetime, to-earliest, to-latest), move the offset to the specified timestamp position. 2. alignment reference (to-offset), move the offset to the specified offset position.
	Strategy *int64 `json:"Strategy,omitnil,omitempty" name:"Strategy"`

	// Specifies the topic name list that needs to reset.
	Topics []*string `json:"Topics,omitnil,omitempty" name:"Topics"`

	// When `strategy` is 0, this field is required. If it is above zero, the offset will be shifted backward by the value of the `shift`. If it is below zero, the offset will be shifted forward by the value of the `shift`. After a correct reset, the new offset should be (old_offset + shift). Note that if the new offset is smaller than the `earliest` parameter of the partition, it will be set to `earliest`, and if it is greater than the `latest` parameter of the partition, it will be set to `latest`
	Shift *int64 `json:"Shift,omitnil,omitempty" name:"Shift"`

	// In milliseconds. when strategy is 1, must include this field. among them, -2 means reset offset to the start position, -1 means reset to the latest position (equivalent to clearing), other values represent the specified time. obtain the offset at the specified time in the topic and reset. notably, if no message exists at the specified time, get the last offset.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// Default number of partitions for a newly created topic. if AutoCreateTopicEnable is set to true and no value is set, defaults to 3.
	DefaultNumPartitions *int64 `json:"DefaultNumPartitions,omitnil,omitempty" name:"DefaultNumPartitions"`

	// Default number of replicas for a newly created topic. if AutoCreateTopicEnable is set to true and not specified, defaults to 2.
	DefaultReplicationFactor *int64 `json:"DefaultReplicationFactor,omitnil,omitempty" name:"DefaultReplicationFactor"`
}

// Predefined struct for user
type ModifyInstanceAttributesRequestParams struct {
	// ckafka cluster instance Id. obtain through the API [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1).
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Maximum retention time of instance logs, in minutes, with a value range of 1min to 90 days.
	MsgRetentionTime *int64 `json:"MsgRetentionTime,omitnil,omitempty" name:"MsgRetentionTime"`

	// Specifies the Name of the ckafka cluster instance.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Instance configuration
	Config *ModifyInstanceAttributesConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// Dynamic message retention policy configuration
	DynamicRetentionConfig *DynamicRetentionTime `json:"DynamicRetentionConfig,omitnil,omitempty" name:"DynamicRetentionConfig"`

	// Specifies the execution time of a scheduled task for edition upgrade or configuration upgrade in Unix timestamp, accurate to the second.
	RebalanceTime *int64 `json:"RebalanceTime,omitnil,omitempty" name:"RebalanceTime"`

	// Public network bandwidth. minimum 3 Mbps. maximum 999 Mbps. only the pro edition supports filling in.
	PublicNetwork *int64 `json:"PublicNetwork,omitnil,omitempty" name:"PublicNetwork"`

	// Dynamic disk expansion policy configuration.
	//
	// Deprecated: DynamicDiskConfig is deprecated.
	DynamicDiskConfig *DynamicDiskConfig `json:"DynamicDiskConfig,omitnil,omitempty" name:"DynamicDiskConfig"`

	// Single message size at the instance level (unit: byte). value range: 1024 (excluding) to 12582912 (excluding).
	MaxMessageByte *uint64 `json:"MaxMessageByte,omitnil,omitempty" name:"MaxMessageByte"`

	// Whether to allow unsynchronized replicas to be elected as leader. valid values: 1 (enable), 0 (disable).
	UncleanLeaderElectionEnable *int64 `json:"UncleanLeaderElectionEnable,omitnil,omitempty" name:"UncleanLeaderElectionEnable"`

	// Instance deletion protection switch. 1: enabled; 0: disabled.
	DeleteProtectionEnable *int64 `json:"DeleteProtectionEnable,omitnil,omitempty" name:"DeleteProtectionEnable"`
}

type ModifyInstanceAttributesRequest struct {
	*tchttp.BaseRequest
	
	// ckafka cluster instance Id. obtain through the API [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1).
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Maximum retention time of instance logs, in minutes, with a value range of 1min to 90 days.
	MsgRetentionTime *int64 `json:"MsgRetentionTime,omitnil,omitempty" name:"MsgRetentionTime"`

	// Specifies the Name of the ckafka cluster instance.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Instance configuration
	Config *ModifyInstanceAttributesConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// Dynamic message retention policy configuration
	DynamicRetentionConfig *DynamicRetentionTime `json:"DynamicRetentionConfig,omitnil,omitempty" name:"DynamicRetentionConfig"`

	// Specifies the execution time of a scheduled task for edition upgrade or configuration upgrade in Unix timestamp, accurate to the second.
	RebalanceTime *int64 `json:"RebalanceTime,omitnil,omitempty" name:"RebalanceTime"`

	// Public network bandwidth. minimum 3 Mbps. maximum 999 Mbps. only the pro edition supports filling in.
	PublicNetwork *int64 `json:"PublicNetwork,omitnil,omitempty" name:"PublicNetwork"`

	// Dynamic disk expansion policy configuration.
	DynamicDiskConfig *DynamicDiskConfig `json:"DynamicDiskConfig,omitnil,omitempty" name:"DynamicDiskConfig"`

	// Single message size at the instance level (unit: byte). value range: 1024 (excluding) to 12582912 (excluding).
	MaxMessageByte *uint64 `json:"MaxMessageByte,omitnil,omitempty" name:"MaxMessageByte"`

	// Whether to allow unsynchronized replicas to be elected as leader. valid values: 1 (enable), 0 (disable).
	UncleanLeaderElectionEnable *int64 `json:"UncleanLeaderElectionEnable,omitnil,omitempty" name:"UncleanLeaderElectionEnable"`

	// Instance deletion protection switch. 1: enabled; 0: disabled.
	DeleteProtectionEnable *int64 `json:"DeleteProtectionEnable,omitnil,omitempty" name:"DeleteProtectionEnable"`
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
	delete(f, "UncleanLeaderElectionEnable")
	delete(f, "DeleteProtectionEnable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceAttributesResponseParams struct {
	// Returned result
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// ckafka cluster instance Id. obtain through the API [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1).
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Specifies the disk capacity in GB. value range: 100 to 500000 with a step length of 100.
	// Specification limits can be viewed through the following link: https://www.tencentcloud.comom/document/product/597/122562.?from_cn_redirect=1
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// Peak bandwidth in MB/s.
	// Specifies the specification limits and corresponding step length through the following link: https://www.tencentcloud.comom/document/product/597/11745.?from_cn_redirect=1
	BandWidth *int64 `json:"BandWidth,omitnil,omitempty" name:"BandWidth"`

	// Partition upper bound. maximum value of 40000. step length of 100.
	// Specifies the specifications and limits that can be viewed through the following link: https://www.tencentcloud.comom/document/product/597/122563.?from_cn_redirect=1
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`
}

type ModifyInstancePreRequest struct {
	*tchttp.BaseRequest
	
	// ckafka cluster instance Id. obtain through the API [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1).
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Specifies the disk capacity in GB. value range: 100 to 500000 with a step length of 100.
	// Specification limits can be viewed through the following link: https://www.tencentcloud.comom/document/product/597/122562.?from_cn_redirect=1
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// Peak bandwidth in MB/s.
	// Specifies the specification limits and corresponding step length through the following link: https://www.tencentcloud.comom/document/product/597/11745.?from_cn_redirect=1
	BandWidth *int64 `json:"BandWidth,omitnil,omitempty" name:"BandWidth"`

	// Partition upper bound. maximum value of 40000. step length of 100.
	// Specifies the specifications and limits that can be viewed through the following link: https://www.tencentcloud.comom/document/product/597/122563.?from_cn_redirect=1
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// Instance Id. you can obtain it by calling the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Specifies the username, which can be obtained through the [DescribeUser](https://www.tencentcloud.comom/document/product/597/40855?from_cn_redirect=1) api.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Current user password
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// New user password
	PasswordNew *string `json:"PasswordNew,omitnil,omitempty" name:"PasswordNew"`
}

type ModifyPasswordRequest struct {
	*tchttp.BaseRequest
	
	// Instance Id. you can obtain it by calling the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Specifies the username, which can be obtained through the [DescribeUser](https://www.tencentcloud.comom/document/product/597/40855?from_cn_redirect=1) api.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
type ModifyRoutineMaintenanceTaskRequestParams struct {
	// Specifies the ckafka cluster instance id. can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Automated operation and maintenance category. valid values: QUOTA, ANALYSIS, RE_BALANCE, ELASTIC_BANDWIDTH.
	MaintenanceType *string `json:"MaintenanceType,omitnil,omitempty" name:"MaintenanceType"`

	// INSTANCE_STORAGE_CAPACITY (automatic disk scale-out)/MESSAGE_RETENTION_PERIOD (dynamic MESSAGE RETENTION policy).
	MaintenanceSubtype *string `json:"MaintenanceSubtype,omitnil,omitempty" name:"MaintenanceSubtype"`

	// Topic name
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Task trigger threshold.
	ConfigureThreshold *int64 `json:"ConfigureThreshold,omitnil,omitempty" name:"ConfigureThreshold"`

	// Specifies the step length for task adjustment.
	ConfigureStepSize *int64 `json:"ConfigureStepSize,omitnil,omitempty" name:"ConfigureStepSize"`

	// Task adjustment upper limit.
	ConfigureLimit *int64 `json:"ConfigureLimit,omitnil,omitempty" name:"ConfigureLimit"`

	// Specifies the expected trigger time of the task, storing the offset in seconds from 0 AM of the current day.
	PlannedTime *int64 `json:"PlannedTime,omitnil,omitempty" name:"PlannedTime"`

	// Additional task information.
	ExtraConfig *string `json:"ExtraConfig,omitnil,omitempty" name:"ExtraConfig"`

	// Task status. 0: enabled, 1: disabled.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Specifies the day of the week.
	Week *string `json:"Week,omitnil,omitempty" name:"Week"`
}

type ModifyRoutineMaintenanceTaskRequest struct {
	*tchttp.BaseRequest
	
	// Specifies the ckafka cluster instance id. can be obtained through the [DescribeInstances](https://www.tencentcloud.comom/document/product/597/40835?from_cn_redirect=1) api.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Automated operation and maintenance category. valid values: QUOTA, ANALYSIS, RE_BALANCE, ELASTIC_BANDWIDTH.
	MaintenanceType *string `json:"MaintenanceType,omitnil,omitempty" name:"MaintenanceType"`

	// INSTANCE_STORAGE_CAPACITY (automatic disk scale-out)/MESSAGE_RETENTION_PERIOD (dynamic MESSAGE RETENTION policy).
	MaintenanceSubtype *string `json:"MaintenanceSubtype,omitnil,omitempty" name:"MaintenanceSubtype"`

	// Topic name
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Task trigger threshold.
	ConfigureThreshold *int64 `json:"ConfigureThreshold,omitnil,omitempty" name:"ConfigureThreshold"`

	// Specifies the step length for task adjustment.
	ConfigureStepSize *int64 `json:"ConfigureStepSize,omitnil,omitempty" name:"ConfigureStepSize"`

	// Task adjustment upper limit.
	ConfigureLimit *int64 `json:"ConfigureLimit,omitnil,omitempty" name:"ConfigureLimit"`

	// Specifies the expected trigger time of the task, storing the offset in seconds from 0 AM of the current day.
	PlannedTime *int64 `json:"PlannedTime,omitnil,omitempty" name:"PlannedTime"`

	// Additional task information.
	ExtraConfig *string `json:"ExtraConfig,omitnil,omitempty" name:"ExtraConfig"`

	// Task status. 0: enabled, 1: disabled.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Specifies the day of the week.
	Week *string `json:"Week,omitnil,omitempty" name:"Week"`
}

func (r *ModifyRoutineMaintenanceTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRoutineMaintenanceTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "MaintenanceType")
	delete(f, "MaintenanceSubtype")
	delete(f, "TopicName")
	delete(f, "ConfigureThreshold")
	delete(f, "ConfigureStepSize")
	delete(f, "ConfigureLimit")
	delete(f, "PlannedTime")
	delete(f, "ExtraConfig")
	delete(f, "Status")
	delete(f, "Week")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRoutineMaintenanceTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRoutineMaintenanceTaskResponseParams struct {
	// Returned results.
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRoutineMaintenanceTaskResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRoutineMaintenanceTaskResponseParams `json:"Response"`
}

func (r *ModifyRoutineMaintenanceTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRoutineMaintenanceTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTopicAttributesRequestParams struct {
	// The ckafka cluster instance Id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Topic name
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

	// Max message size in bytes. Max value: 8,388,608 bytes (8 MB).
	MaxMessageBytes *int64 `json:"MaxMessageBytes,omitnil,omitempty" name:"MaxMessageBytes"`

	// Duration of Segment shard scrolling in milliseconds. current min value is 86400000 ms.
	SegmentMs *int64 `json:"SegmentMs,omitnil,omitempty" name:"SegmentMs"`

	// Message deletion policy. Valid values: delete, compact
	CleanUpPolicy *string `json:"CleanUpPolicy,omitnil,omitempty" name:"CleanUpPolicy"`

	// IP allowlist, which is required if the value of `enableWhileList` is 1.
	IpWhiteList []*string `json:"IpWhiteList,omitnil,omitempty" name:"IpWhiteList"`

	// Preset ACL rule. `1`: enable, `0`: disable. Default value: `0`.
	EnableAclRule *int64 `json:"EnableAclRule,omitnil,omitempty" name:"EnableAclRule"`

	// ACL rule name.
	AclRuleName *string `json:"AclRuleName,omitnil,omitempty" name:"AclRuleName"`

	// Message retention file size in bytes, which is an optional parameter. Default value: -1. Currently, the min value that can be entered is 1,048,576 B.
	RetentionBytes *int64 `json:"RetentionBytes,omitnil,omitempty" name:"RetentionBytes"`

	// Tag list.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Production traffic throttling in MB/s. set to -1 to disable throttling.
	QuotaProducerByteRate *int64 `json:"QuotaProducerByteRate,omitnil,omitempty" name:"QuotaProducerByteRate"`

	// Consumption traffic throttling in MB/s. set to -1 for unlimited consumption.
	QuotaConsumerByteRate *int64 `json:"QuotaConsumerByteRate,omitnil,omitempty" name:"QuotaConsumerByteRate"`

	// Number of topic replicas. valid values: 1, 3.
	ReplicaNum *int64 `json:"ReplicaNum,omitnil,omitempty" name:"ReplicaNum"`

	// Specifies the time type for message saving: CreateTime/LogAppendTime.
	LogMsgTimestampType *string `json:"LogMsgTimestampType,omitnil,omitempty" name:"LogMsgTimestampType"`
}

type ModifyTopicAttributesRequest struct {
	*tchttp.BaseRequest
	
	// The ckafka cluster instance Id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Topic name
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

	// Max message size in bytes. Max value: 8,388,608 bytes (8 MB).
	MaxMessageBytes *int64 `json:"MaxMessageBytes,omitnil,omitempty" name:"MaxMessageBytes"`

	// Duration of Segment shard scrolling in milliseconds. current min value is 86400000 ms.
	SegmentMs *int64 `json:"SegmentMs,omitnil,omitempty" name:"SegmentMs"`

	// Message deletion policy. Valid values: delete, compact
	CleanUpPolicy *string `json:"CleanUpPolicy,omitnil,omitempty" name:"CleanUpPolicy"`

	// IP allowlist, which is required if the value of `enableWhileList` is 1.
	IpWhiteList []*string `json:"IpWhiteList,omitnil,omitempty" name:"IpWhiteList"`

	// Preset ACL rule. `1`: enable, `0`: disable. Default value: `0`.
	EnableAclRule *int64 `json:"EnableAclRule,omitnil,omitempty" name:"EnableAclRule"`

	// ACL rule name.
	AclRuleName *string `json:"AclRuleName,omitnil,omitempty" name:"AclRuleName"`

	// Message retention file size in bytes, which is an optional parameter. Default value: -1. Currently, the min value that can be entered is 1,048,576 B.
	RetentionBytes *int64 `json:"RetentionBytes,omitnil,omitempty" name:"RetentionBytes"`

	// Tag list.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Production traffic throttling in MB/s. set to -1 to disable throttling.
	QuotaProducerByteRate *int64 `json:"QuotaProducerByteRate,omitnil,omitempty" name:"QuotaProducerByteRate"`

	// Consumption traffic throttling in MB/s. set to -1 for unlimited consumption.
	QuotaConsumerByteRate *int64 `json:"QuotaConsumerByteRate,omitnil,omitempty" name:"QuotaConsumerByteRate"`

	// Number of topic replicas. valid values: 1, 3.
	ReplicaNum *int64 `json:"ReplicaNum,omitnil,omitempty" name:"ReplicaNum"`

	// Specifies the time type for message saving: CreateTime/LogAppendTime.
	LogMsgTimestampType *string `json:"LogMsgTimestampType,omitnil,omitempty" name:"LogMsgTimestampType"`
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
	delete(f, "MaxMessageBytes")
	delete(f, "SegmentMs")
	delete(f, "CleanUpPolicy")
	delete(f, "IpWhiteList")
	delete(f, "EnableAclRule")
	delete(f, "AclRuleName")
	delete(f, "RetentionBytes")
	delete(f, "Tags")
	delete(f, "QuotaProducerByteRate")
	delete(f, "QuotaConsumerByteRate")
	delete(f, "ReplicaNum")
	delete(f, "LogMsgTimestampType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTopicAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTopicAttributesResponseParams struct {
	// Returned result.
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// Flow ID.
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// RouteIdDto
	RouteDTO *RouteDTO `json:"RouteDTO,omitnil,omitempty" name:"RouteDTO"`
}

type Partition struct {
	// Partition ID
	PartitionId *int64 `json:"PartitionId,omitnil,omitempty" name:"PartitionId"`
}

type PartitionOffset struct {
	// Partition
	Partition *string `json:"Partition,omitnil,omitempty" name:"Partition"`

	// Specifies the offset.
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

	// Region code.
	RegionCode *string `json:"RegionCode,omitnil,omitempty" name:"RegionCode"`

	// Region code (V3 version).
	RegionCodeV3 *string `json:"RegionCodeV3,omitnil,omitempty" name:"RegionCodeV3"`

	// Specifies the default value does not support any special type instance type.
	Support *string `json:"Support,omitnil,omitempty" name:"Support"`

	// Whether ipv6 is supported. 0: indicates no support. 1: indicates support.
	Ipv6 *int64 `json:"Ipv6,omitnil,omitempty" name:"Ipv6"`

	// Whether cross-az is supported. valid values: 0 (unsupported), 1 (supported).
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

	// Specifies the network type of the route (3: vpc routing; 7: internal support route; 1: public network route).
	VipType *int64 `json:"VipType,omitnil,omitempty" name:"VipType"`

	// Virtual IP list
	VipList []*VipEntity `json:"VipList,omitnil,omitempty" name:"VipList"`

	// Domain name
	// Note: this field may return null, indicating that no valid values can be obtained.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Domain name port
	// Note: this field may return null, indicating that no valid values can be obtained.
	DomainPort *int64 `json:"DomainPort,omitnil,omitempty" name:"DomainPort"`

	// Timestamp.
	DeleteTimestamp *string `json:"DeleteTimestamp,omitnil,omitempty" name:"DeleteTimestamp"`

	// Specifies the subnet Id.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Subnet *string `json:"Subnet,omitnil,omitempty" name:"Subnet"`

	// Virtual IP list (1:1 broker node).
	BrokerVipList []*VipEntity `json:"BrokerVipList,omitnil,omitempty" name:"BrokerVipList"`

	// VPC Id. specifies the Id of the vpc.
	// Note: This field may return null, indicating that no valid values can be obtained.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Remarks
	// Note: This field may return null, indicating that no valid values can be obtained.
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// Route status. 1: creating, 2: creation succeeded, 3: creation failed, 4: deleting, 6: deletion failed.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type RouteDTO struct {
	// Route ID
	RouteId *int64 `json:"RouteId,omitnil,omitempty" name:"RouteId"`
}

type RouteFilter struct {
	// Filters by name. currently supports security-group-id. filters by security group association.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Filter value. when the filter name is security-group-id, only supports transmission of one value.
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// Filter relationship. supports IN and NOT_IN. default is IN.
	Relation *string `json:"Relation,omitnil,omitempty" name:"Relation"`
}

type RouteResponse struct {
	// Route Information List
	Routers []*Route `json:"Routers,omitnil,omitempty" name:"Routers"`
}

type SaleInfo struct {
	// The manually configured flag. valid values: true (sold-out), false (available).
	Flag *bool `json:"Flag,omitnil,omitempty" name:"Flag"`

	// Specifies the ckafka version number (1.1.1/2.4.2/0.10.2).
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// Pro edition, standard version flag.
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// Specifies whether the item is sold-out. valid values: true (sold-out).
	SoldOut *bool `json:"SoldOut,omitnil,omitempty" name:"SoldOut"`
}

type ScalingDownResp struct {
	// Order ID list
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`
}

type SecurityGroupRoute struct {
	// Specifies the routing information.
	InstanceRoute *InstanceRoute `json:"InstanceRoute,omitnil,omitempty" name:"InstanceRoute"`

	// Specifies the security group list to associate.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// CKafka cluster instance name.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Specifies the route vpcId.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Route vip.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`
}

type SecurityGroupRouteResp struct {
	// Total number of eligible security group routes.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Eligible security group route information list.
	SecurityGroupRoutes []*SecurityGroupRoute `json:"SecurityGroupRoutes,omitnil,omitempty" name:"SecurityGroupRoutes"`
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// Specifies the subscription partition.
	Partition []*int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// Specifies the partition offset information.
	PartitionOffset []*PartitionOffset `json:"PartitionOffset,omitnil,omitempty" name:"PartitionOffset"`

	// Subscribed topic ID.
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

	// Output information.
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

	// Specifies the unix second-level timestamp of the creation time.
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Describes the topic remark.
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

	// ACL preset policy switch. valid values: 1 (on); 0 (off).
	EnableAclRule *int64 `json:"EnableAclRule,omitnil,omitempty" name:"EnableAclRule"`

	// Preset policy list.
	AclRuleList []*AclRule `json:"AclRuleList,omitnil,omitempty" name:"AclRuleList"`

	// topic throttling policy.
	QuotaConfig *InstanceQuotaConfigResp `json:"QuotaConfig,omitnil,omitempty" name:"QuotaConfig"`

	// Number of replicas
	ReplicaNum *int64 `json:"ReplicaNum,omitnil,omitempty" name:"ReplicaNum"`
}

type TopicDetail struct {
	// Topic name
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Topic Id.
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Number of partitions
	PartitionNum *int64 `json:"PartitionNum,omitnil,omitempty" name:"PartitionNum"`

	// Number of topic replicas. valid values: 1, 3.
	ReplicaNum *int64 `json:"ReplicaNum,omitnil,omitempty" name:"ReplicaNum"`

	// Remarks.
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// Creation time
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Whether to enable IP authentication allowlist. true: yes, false: no
	EnableWhiteList *bool `json:"EnableWhiteList,omitnil,omitempty" name:"EnableWhiteList"`

	// Number of IPs in IP allowlist
	IpWhiteListCount *int64 `json:"IpWhiteListCount,omitnil,omitempty" name:"IpWhiteListCount"`

	// Data backup cos bucket. specifies the bucket address for archiving to cos.
	ForwardCosBucket *string `json:"ForwardCosBucket,omitnil,omitempty" name:"ForwardCosBucket"`

	// Status of data backup to COS. 1: not enabled, 0: enabled
	ForwardStatus *int64 `json:"ForwardStatus,omitnil,omitempty" name:"ForwardStatus"`

	// Frequency of data backup to COS
	ForwardInterval *int64 `json:"ForwardInterval,omitnil,omitempty" name:"ForwardInterval"`

	// Advanced configuration.
	Config *Config `json:"Config,omitnil,omitempty" name:"Config"`

	// Message retention period configuration (used for dynamic configuration change records).
	RetentionTimeConfig *TopicRetentionTimeConfigRsp `json:"RetentionTimeConfig,omitnil,omitempty" name:"RetentionTimeConfig"`

	// 0: normal. 1: deleted. 2: deleting.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Tag list
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type TopicDetailResponse struct {
	// List of returned topic details.
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

	// Start Offset.
	BeginOffset *uint64 `json:"BeginOffset,omitnil,omitempty" name:"BeginOffset"`

	// End Offset.
	EndOffset *uint64 `json:"EndOffset,omitnil,omitempty" name:"EndOffset"`

	// Message count.
	MessageCount *uint64 `json:"MessageCount,omitnil,omitempty" name:"MessageCount"`

	// Unsynced replica.
	OutOfSyncReplica *string `json:"OutOfSyncReplica,omitnil,omitempty" name:"OutOfSyncReplica"`
}

type TopicInSyncReplicaResult struct {
	// Set of topic details and replicas
	TopicInSyncReplicaList []*TopicInSyncReplicaInfo `json:"TopicInSyncReplicaList,omitnil,omitempty" name:"TopicInSyncReplicaList"`

	// Total number
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type TopicPartitionDO struct {
	// Partition ID. specifies the Partition ID.
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`

	// Leader running status. 0 means running normally.
	LeaderStatus *int64 `json:"LeaderStatus,omitnil,omitempty" name:"LeaderStatus"`

	// ISR quantity
	IsrNum *int64 `json:"IsrNum,omitnil,omitempty" name:"IsrNum"`

	// Number of replicas
	ReplicaNum *int64 `json:"ReplicaNum,omitnil,omitempty" name:"ReplicaNum"`
}

type TopicResult struct {
	// List of returned topic information.
	TopicList []*Topic `json:"TopicList,omitnil,omitempty" name:"TopicList"`

	// Number of eligible topics.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type TopicRetentionTimeConfigRsp struct {
	// Expected value, the message retention period (in minutes) set by user configuration.
	Expect *int64 `json:"Expect,omitnil,omitempty" name:"Expect"`

	// Current value, which is the current effective value (may contain dynamic adjustment in minutes).
	Current *int64 `json:"Current,omitnil,omitempty" name:"Current"`

	// Last modified time.
	ModTimeStamp *int64 `json:"ModTimeStamp,omitnil,omitempty" name:"ModTimeStamp"`
}

type TopicSubscribeGroup struct {
	// Total number
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Number of consumer group status
	StatusCountInfo *string `json:"StatusCountInfo,omitnil,omitempty" name:"StatusCountInfo"`

	// Consumer group information.
	GroupsInfo []*GroupInfoResponse `json:"GroupsInfo,omitnil,omitempty" name:"GroupsInfo"`

	// Indicates whether the request is asynchronous. instances with fewer groups will return results directly with Status as 1. when there are more groups, the cache will be updated asynchronously. no group information will be returned when Status is 0 until the update is complete and results are returned with Status as 1.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type UpgradeBrokerVersionRequestParams struct {
	// The ckafka cluster instance Id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 1. smooth configuration upgrade 2. vertical configuration upgrade.
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Version number
	SourceVersion *string `json:"SourceVersion,omitnil,omitempty" name:"SourceVersion"`

	// Version number
	TargetVersion *string `json:"TargetVersion,omitnil,omitempty" name:"TargetVersion"`

	// Delay time.
	DelayTimeStamp *string `json:"DelayTimeStamp,omitnil,omitempty" name:"DelayTimeStamp"`
}

type UpgradeBrokerVersionRequest struct {
	*tchttp.BaseRequest
	
	// The ckafka cluster instance Id.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 1. smooth configuration upgrade 2. vertical configuration upgrade.
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Version number
	SourceVersion *string `json:"SourceVersion,omitnil,omitempty" name:"SourceVersion"`

	// Version number
	TargetVersion *string `json:"TargetVersion,omitnil,omitempty" name:"TargetVersion"`

	// Delay time.
	DelayTimeStamp *string `json:"DelayTimeStamp,omitnil,omitempty" name:"DelayTimeStamp"`
}

func (r *UpgradeBrokerVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeBrokerVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Type")
	delete(f, "SourceVersion")
	delete(f, "TargetVersion")
	delete(f, "DelayTimeStamp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeBrokerVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeBrokerVersionResponseParams struct {
	// Upgrade result.
	Result *JgwOperateResponse `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeBrokerVersionResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeBrokerVersionResponseParams `json:"Response"`
}

func (r *UpgradeBrokerVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeBrokerVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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
	// Specifies the eligible users list.
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
	// Availability zone
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Whether it is an internal application.
	IsInternalApp *int64 `json:"IsInternalApp,omitnil,omitempty" name:"IsInternalApp"`

	// Application identifier
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Indicates whether the AZ is sold out. true indicates sold out. false indicates not sold out.
	Flag *bool `json:"Flag,omitnil,omitempty" name:"Flag"`

	// Availability zone name.
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// Availability zone status. enumerates example: 3: enable, 4: disable. availability zone status is subject to SoldOut.
	ZoneStatus *int64 `json:"ZoneStatus,omitnil,omitempty" name:"ZoneStatus"`

	// Extra flag
	//
	// Deprecated: Exflag is deprecated.
	Exflag *string `json:"Exflag,omitnil,omitempty" name:"Exflag"`

	// Specifies whether the item is sold-out. valid values: true (sold-out), false (not sold out).
	SoldOut *string `json:"SoldOut,omitnil,omitempty" name:"SoldOut"`

	// Specifies the sell-out information of the standard version.
	SalesInfo []*SaleInfo `json:"SalesInfo,omitnil,omitempty" name:"SalesInfo"`

	// Additional flag.
	ExtraFlag *string `json:"ExtraFlag,omitnil,omitempty" name:"ExtraFlag"`
}

type ZoneResponse struct {
	// <P>Specifies the zone list.</p>.
	ZoneList []*ZoneInfo `json:"ZoneList,omitnil,omitempty" name:"ZoneList"`

	// <P>Maximum number of instances that can be purchased.</p>.
	MaxBuyInstanceNum *int64 `json:"MaxBuyInstanceNum,omitnil,omitempty" name:"MaxBuyInstanceNum"`

	// <p>Maximum purchase bandwidth in Mb/s.</p>.
	MaxBandwidth *int64 `json:"MaxBandwidth,omitnil,omitempty" name:"MaxBandwidth"`

	// <P>Unit price for postpayment.</p>.
	UnitPrice *Price `json:"UnitPrice,omitnil,omitempty" name:"UnitPrice"`

	// <P>Message unit price for postpayment.</p>.
	MessagePrice *Price `json:"MessagePrice,omitnil,omitempty" name:"MessagePrice"`

	// <P>User-Exclusive cluster information.</p>.
	ClusterInfo []*ClusterInfo `json:"ClusterInfo,omitnil,omitempty" name:"ClusterInfo"`

	// <P>Specifies the standard version configuration to purchase.</p>.
	Standard *string `json:"Standard,omitnil,omitempty" name:"Standard"`

	// <P>Specifies the purchase of standard version s2 configuration.</p>.
	StandardS2 *string `json:"StandardS2,omitnil,omitempty" name:"StandardS2"`

	// <P>Specifies the configuration for purchasing professional edition.</p>.
	Profession *string `json:"Profession,omitnil,omitempty" name:"Profession"`

	// <P>Purchase physical dedicated edition configuration.</p>.
	Physical *string `json:"Physical,omitnil,omitempty" name:"Physical"`

	// <p>Specifies the public network bandwidth. valid values: 3Mbps to 999Mbps. only supported in pro edition. abandoned, meaningless.</p>.
	PublicNetwork *string `json:"PublicNetwork,omitnil,omitempty" name:"PublicNetwork"`

	// <P>Public network bandwidth configuration.</p>.
	PublicNetworkLimit *string `json:"PublicNetworkLimit,omitnil,omitempty" name:"PublicNetworkLimit"`

	// <p>Request Id.</p>.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`

	// <P>Specifies the pagination offset.</p>.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <P>Specifies the pagination limit.</p>.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <P>Specifies whether the tag is mandatory.</p>.
	ForceCheckTag *bool `json:"ForceCheckTag,omitnil,omitempty" name:"ForceCheckTag"`
}