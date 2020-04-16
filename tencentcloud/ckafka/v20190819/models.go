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
	AclList []*Acl `json:"AclList,omitempty" name:"AclList" list`
}

type AppIdResponse struct {

	// Number of eligible `AppId`
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of eligible `AppId`
	// Note: this field may return null, indicating that no valid values can be obtained.
	AppIdList []*int64 `json:"AppIdList,omitempty" name:"AppIdList" list`
}

type Assignment struct {

	// Assignment version information
	Version *int64 `json:"Version,omitempty" name:"Version"`

	// Topic information list
	// Note: this field may return null, indicating that no valid values can be obtained.
	Topics []*GroupInfoTopics `json:"Topics,omitempty" name:"Topics" list`
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
}

type ConsumerGroup struct {

	// User group name
	ConsumerGroupName *string `json:"ConsumerGroupName,omitempty" name:"ConsumerGroupName"`

	// Subscribed message entity
	SubscribedInfo []*SubscribedInfo `json:"SubscribedInfo,omitempty" name:"SubscribedInfo" list`
}

type ConsumerGroupResponse struct {

	// Number of eligible consumer groups
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Topic list
	// Note: this field may return null, indicating that no valid values can be obtained.
	TopicList []*ConsumerGroupTopic `json:"TopicList,omitempty" name:"TopicList" list`

	// Consumer group list
	// Note: this field may return null, indicating that no valid values can be obtained.
	GroupList []*ConsumerGroup `json:"GroupList,omitempty" name:"GroupList" list`

	// Total number of partitions
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalPartition *int64 `json:"TotalPartition,omitempty" name:"TotalPartition"`

	// List of monitored partitions
	// Note: this field may return null, indicating that no valid values can be obtained.
	PartitionListForMonitor []*Partition `json:"PartitionListForMonitor,omitempty" name:"PartitionListForMonitor" list`

	// Total number of topics
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalTopic *int64 `json:"TotalTopic,omitempty" name:"TotalTopic"`

	// List of monitored topics
	// Note: this field may return null, indicating that no valid values can be obtained.
	TopicListForMonitor []*ConsumerGroupTopic `json:"TopicListForMonitor,omitempty" name:"TopicListForMonitor" list`

	// List of monitored groups
	// Note: this field may return null, indicating that no valid values can be obtained.
	GroupListForMonitor []*Group `json:"GroupListForMonitor,omitempty" name:"GroupListForMonitor" list`
}

type ConsumerGroupTopic struct {

	// Topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Topic name
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

type CreateAclRequest struct {
	*tchttp.BaseRequest

	// Instance ID information
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// ACL resource type. 0: UNKNOWN, 1: ANY, 2: TOPIC, 3: GROUP, 4: CLUSTER, 5: TRANSACTIONAL_ID. Currently, only `TOPIC` is available, and other fields will be used for future ACLs compatible with open-source Kafka
	ResourceType *int64 `json:"ResourceType,omitempty" name:"ResourceType"`

	// Resource name, which is related to `resourceType`. For example, if `resourceType` is `TOPIC`, this field indicates the topic name; if `resourceType` is `GROUP`, this field indicates the group name
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// ACL operation mode. 0: UNKNOWN, 1: ANY, 2: ALL, 3: READ, 4: WRITE, 5: CREATE, 6: DELETE, 7: ALTER, 8: DESCRIBE, 9: CLUSTER_ACTION, 10: DESCRIBE_CONFIGS, 11: ALTER_CONFIGS
	Operation *int64 `json:"Operation,omitempty" name:"Operation"`

	// Permission type. 0: UNKNOWN, 1: ANY, 2: DENY, 3: ALLOW. Currently, CKafka supports `ALLOW` (equivalent to whitelist), and other fields will be used for future ACLs compatible with open-source Kafka
	PermissionType *int64 `json:"PermissionType,omitempty" name:"PermissionType"`

	// The default value is `*`, which means that any host can access. Currently, CKafka does not support the host as `*`, but the future product based on the open-source Kafka will directly support this
	Host *string `json:"Host,omitempty" name:"Host"`

	// User list. The default value is `*`, which means that any user can access. The current user can only be one included in the user list
	Principal *string `json:"Principal,omitempty" name:"Principal"`
}

func (r *CreateAclRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAclRequest) FromJsonString(s string) error {
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

func (r *CreatePartitionRequest) FromJsonString(s string) error {
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

func (r *CreatePartitionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTopicIpWhiteListRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Topic name
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// IP whitelist list
	IpWhiteList []*string `json:"IpWhiteList,omitempty" name:"IpWhiteList" list`
}

func (r *CreateTopicIpWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTopicIpWhiteListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTopicIpWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result of deleting topic IP whitelist
		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTopicIpWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTopicIpWhiteListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTopicRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Topic name string of up to 64 characters, which must begin with a letter and can contain letters, digits, and dashes (`-`)
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Number of partitions, which should be greater than 0
	PartitionNum *int64 `json:"PartitionNum,omitempty" name:"PartitionNum"`

	// Number of replicas, which cannot be higher than the number of brokers. Maximum value: 3
	ReplicaNum *int64 `json:"ReplicaNum,omitempty" name:"ReplicaNum"`

	// IP whitelist switch. 1: enabled, 0: disabled. Default value: 0
	EnableWhiteList *int64 `json:"EnableWhiteList,omitempty" name:"EnableWhiteList"`

	// IP whitelist list for quota limit, which is required if `enableWhileList` is 1
	IpWhiteList []*string `json:"IpWhiteList,omitempty" name:"IpWhiteList" list`

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
}

func (r *CreateTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTopicRequest) FromJsonString(s string) error {
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

func (r *CreateUserRequest) FromJsonString(s string) error {
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

func (r *CreateUserResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAclRequest struct {
	*tchttp.BaseRequest

	// Instance ID information
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// ACL resource type. 0: UNKNOWN, 1: ANY, 2: TOPIC, 3: GROUP, 4: CLUSTER, 5: TRANSACTIONAL_ID. Currently, only `TOPIC` is available, and other fields will be used for future ACLs compatible with open-source Kafka
	ResourceType *int64 `json:"ResourceType,omitempty" name:"ResourceType"`

	// Resource name, which is related to `resourceType`. For example, if `resourceType` is `TOPIC`, this field indicates the topic name; if `resourceType` is `GROUP`, this field indicates the group name
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// ACL operation mode. 0: UNKNOWN, 1: ANY, 2: ALL, 3: READ, 4: WRITE, 5: CREATE, 6: DELETE, 7: ALTER, 8: DESCRIBE, 9: CLUSTER_ACTION, 10: DESCRIBE_CONFIGS, 11: ALTER_CONFIGS, 12: IDEMPOTEN_WRITE. Currently, CKafka only supports `READ` and `WRITE`, and other values will be used for future ACLs compatible with open-source Kafka
	Operation *int64 `json:"Operation,omitempty" name:"Operation"`

	// Permission type. 0: UNKNOWN, 1: ANY, 2: DENY, 3: ALLOW. Currently, CKafka supports `ALLOW` (equivalent to whitelist), and other fields will be used for future ACLs compatible with open-source Kafka
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

func (r *DeleteAclRequest) FromJsonString(s string) error {
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

func (r *DeleteAclResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTopicIpWhiteListRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Topic name
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// IP whitelist list
	IpWhiteList []*string `json:"IpWhiteList,omitempty" name:"IpWhiteList" list`
}

func (r *DeleteTopicIpWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTopicIpWhiteListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTopicIpWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result of deleting topic IP whitelist
		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTopicIpWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

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

func (r *DeleteTopicRequest) FromJsonString(s string) error {
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

func (r *DeleteUserRequest) FromJsonString(s string) error {
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

func (r *DeleteUserResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeACLRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// ACL resource type. 0: UNKNOWN, 1: ANY, 2: TOPIC, 3: GROUP, 4: CLUSTER, 5: TRANSACTIONAL_ID. Currently, only `TOPIC` is available, and other fields will be used for future ACLs compatible with open-source Kafka
	ResourceType *int64 `json:"ResourceType,omitempty" name:"ResourceType"`

	// Resource name, which is related to `resourceType`. For example, if `resourceType` is `TOPIC`, this field indicates the topic name; if `resourceType` is `GROUP`, this field indicates the group name
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

func (r *DescribeACLRequest) FromJsonString(s string) error {
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

func (r *DescribeAppInfoRequest) FromJsonString(s string) error {
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

func (r *DescribeAppInfoResponse) FromJsonString(s string) error {
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

func (r *DescribeConsumerGroupRequest) FromJsonString(s string) error {
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
	GroupList []*string `json:"GroupList,omitempty" name:"GroupList" list`
}

func (r *DescribeGroupInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGroupInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Returned result
	// Note: this field may return null, indicating that no valid values can be obtained.
		Result []*GroupInfoResponse `json:"Result,omitempty" name:"Result" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGroupInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

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
	Topics []*string `json:"Topics,omitempty" name:"Topics" list`

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

func (r *DescribeGroupOffsetsRequest) FromJsonString(s string) error {
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

func (r *DescribeGroupRequest) FromJsonString(s string) error {
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

func (r *DescribeInstanceAttributesRequest) FromJsonString(s string) error {
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
	Status []*int64 `json:"Status,omitempty" name:"Status" list`

	// Offset. If this parameter is left empty, 0 will be used by default
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results to be returned. If this parameter is left empty, 10 will be used by default. The maximum value is 20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Tag key match.
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Filter
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeInstancesDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesDetailRequest) FromJsonString(s string) error {
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
	Status []*int64 `json:"Status,omitempty" name:"Status" list`

	// Offset. If this parameter is left empty, 0 will be used by default
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results to be returned. If this parameter is left empty, 10 will be used by default. The maximum value is 20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Tag key match.
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesRequest) FromJsonString(s string) error {
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

func (r *DescribeInstancesResponse) FromJsonString(s string) error {
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

func (r *DescribeTopicAttributesRequest) FromJsonString(s string) error {
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
}

func (r *DescribeTopicDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTopicDetailRequest) FromJsonString(s string) error {
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
}

func (r *DescribeTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTopicRequest) FromJsonString(s string) error {
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

func (r *DescribeTopicResponse) FromJsonString(s string) error {
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

func (r *DescribeUserRequest) FromJsonString(s string) error {
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

func (r *DescribeUserResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// Field to be filtered.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Filter value of field.
	Values []*string `json:"Values,omitempty" name:"Values" list`
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
	Members []*GroupInfoMember `json:"Members,omitempty" name:"Members" list`

	// Kafka consumer group
	Group *string `json:"Group,omitempty" name:"Group"`
}

type GroupInfoTopics struct {

	// Name of assigned topics
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// Information of assigned partition
	// Note: this field may return null, indicating that no valid values can be obtained.
	Partitions []*int64 `json:"Partitions,omitempty" name:"Partitions" list`
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
	TopicList []*GroupOffsetTopic `json:"TopicList,omitempty" name:"TopicList" list`
}

type GroupOffsetTopic struct {

	// Topic name
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// Array of partitions in the topic, where each element is a JSON object
	// Note: this field may return null, indicating that no valid values can be obtained.
	Partitions []*GroupOffsetPartition `json:"Partitions,omitempty" name:"Partitions" list`
}

type GroupResponse struct {

	// Count
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// GroupList
	// Note: this field may return null, indicating that no valid values can be obtained.
	GroupList []*DescribeGroup `json:"GroupList,omitempty" name:"GroupList" list`
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
	VipList []*VipEntity `json:"VipList,omitempty" name:"VipList" list`

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
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`

	// Expiration time
	// Note: this field may return null, indicating that no valid values can be obtained.
	ExpireTime *uint64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// Cross-AZ
	// Note: this field may return null, indicating that no valid values can be obtained.
	ZoneIds []*int64 `json:"ZoneIds,omitempty" name:"ZoneIds" list`

	// Kafka version information
	// Note: this field may return null, indicating that no valid values can be obtained.
	Version *string `json:"Version,omitempty" name:"Version"`

	// Maximum number of groups
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxGroupNum *int64 `json:"MaxGroupNum,omitempty" name:"MaxGroupNum"`

	// Sale type
	// Note: this field may return null, indicating that no valid values can be obtained.
	Cvm *int64 `json:"Cvm,omitempty" name:"Cvm"`
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
	VipList []*VipEntity `json:"VipList,omitempty" name:"VipList" list`

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
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`

	// Kafka version information
	// Note: this field may return null, indicating that no valid values can be obtained.
	Version *string `json:"Version,omitempty" name:"Version"`

	// Cross-AZ
	// Note: this field may return null, indicating that no valid values can be obtained.
	ZoneIds []*int64 `json:"ZoneIds,omitempty" name:"ZoneIds" list`

	// CKafka sale type
	// Note: this field may return null, indicating that no valid values can be obtained.
	Cvm *int64 `json:"Cvm,omitempty" name:"Cvm"`
}

type InstanceDetailResponse struct {

	// Total number of eligible instances
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of eligible instance details
	InstanceList []*InstanceDetail `json:"InstanceList,omitempty" name:"InstanceList" list`
}

type InstanceResponse struct {

	// List of eligible instances
	// Note: this field may return null, indicating that no valid values can be obtained.
	InstanceList []*Instance `json:"InstanceList,omitempty" name:"InstanceList" list`

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
	Topics []*string `json:"Topics,omitempty" name:"Topics" list`

	// When `strategy` is 0, this field is required. If it is above zero, the offset will be shifted backward by the value of the `shift`. If it is below zero, the offset will be shifted forward by the value of the `shift`. After a correct reset, the new offset should be (old_offset + shift). Note that if the new offset is smaller than the `earliest` parameter of the partition, it will be set to `earliest`, and if it is greater than the `latest` parameter of the partition, it will be set to `latest`
	Shift *int64 `json:"Shift,omitempty" name:"Shift"`

	// Unit: ms. When `strategy` is 1, this field is required, where -2 indicates to reset the offset to the initial position, -1 indicates to reset to the latest position (equivalent to emptying), and other values represent the specified time, i.e., the offset of the topic at the specified time will be obtained and then reset. Note that if there is no message at the specified time, the last offset will be obtained
	ShiftTimestamp *int64 `json:"ShiftTimestamp,omitempty" name:"ShiftTimestamp"`

	// Position of the offset that needs to be reset. When `strategy` is 2, this field is required
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *ModifyGroupOffsetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyGroupOffsetsRequest) FromJsonString(s string) error {
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
}

func (r *ModifyInstanceAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstanceAttributesRequest) FromJsonString(s string) error {
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

func (r *ModifyPasswordRequest) FromJsonString(s string) error {
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

	// IP whitelist switch. 1: enabled, 0: disabled.
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
}

func (r *ModifyTopicAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTopicAttributesRequest) FromJsonString(s string) error {
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

func (r *ModifyTopicAttributesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OperateResponseData struct {

	// FlowId
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

type SubscribedInfo struct {

	// Subscribed topic name
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Subscribed partition
	// Note: this field may return null, indicating that no valid values can be obtained.
	Partition []*int64 `json:"Partition,omitempty" name:"Partition" list`

	// Partition offset information
	// Note: this field may return null, indicating that no valid values can be obtained.
	PartitionOffset []*PartitionOffset `json:"PartitionOffset,omitempty" name:"PartitionOffset" list`
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

	// IP whitelist switch. 1: enabled, 0: disabled
	EnableWhiteList *int64 `json:"EnableWhiteList,omitempty" name:"EnableWhiteList"`

	// IP whitelist list
	IpWhiteList []*string `json:"IpWhiteList,omitempty" name:"IpWhiteList" list`

	// Topic configuration array
	Config *Config `json:"Config,omitempty" name:"Config"`

	// Partition details
	Partitions []*TopicPartitionDO `json:"Partitions,omitempty" name:"Partitions" list`
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

	// Whether to enable IP authentication whitelist. true: yes, false: no
	EnableWhiteList *bool `json:"EnableWhiteList,omitempty" name:"EnableWhiteList"`

	// Number of IPs in IP whitelist
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
}

type TopicDetailResponse struct {

	// List of returned topic details
	// Note: this field may return null, indicating that no valid values can be obtained.
	TopicList []*TopicDetail `json:"TopicList,omitempty" name:"TopicList" list`

	// Number of all eligible topic details
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
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
	TopicList []*Topic `json:"TopicList,omitempty" name:"TopicList" list`

	// Number of eligible topics
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
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
	Users []*User `json:"Users,omitempty" name:"Users" list`

	// Total number of eligible users
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type VipEntity struct {

	// Virtual IP
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Virtual port
	Vport *string `json:"Vport,omitempty" name:"Vport"`
}
