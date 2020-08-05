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

package v20180412

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type Account struct {

	// Instance ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Account name (`root` for a root account)
	// Note: This field may return null, indicating that no valid values can be obtained.
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// Account description information
	// Note: This field may return null, indicating that no valid values can be obtained.
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Read/write policy. r: read-only; w: write-only; rw: read/write
	// Note: This field may return null, indicating that no valid values can be obtained.
	Privilege *string `json:"Privilege,omitempty" name:"Privilege"`

	// Routing policy. master: master node; replication: secondary node
	// Note: This field may return null, indicating that no valid values can be obtained.
	ReadonlyPolicy []*string `json:"ReadonlyPolicy,omitempty" name:"ReadonlyPolicy" list`

	// Sub-account status. 1: account is being changed; 2: account is valid; -4: account has been deleted
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type AssociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// Database engine name: mariadb, cdb, cynosdb, dcdb, redis, mongodb, etc.
	Product *string `json:"Product,omitempty" name:"Product"`

	// ID of the security group to be associated in the format of sg-efil73jd.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// ID(s) of the instance(s) to be associated in the format of ins-lesecurk. You can specify multiple instances.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *AssociateSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssociateSecurityGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssociateSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssociateSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssociateSecurityGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BigKeyInfo struct {

	// Database
	DB *int64 `json:"DB,omitempty" name:"DB"`

	// Big key
	Key *string `json:"Key,omitempty" name:"Key"`

	// Type
	Type *string `json:"Type,omitempty" name:"Type"`

	// Size
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// Data timestamp
	Updatetime *int64 `json:"Updatetime,omitempty" name:"Updatetime"`
}

type BigKeyTypeInfo struct {

	// Type
	Type *string `json:"Type,omitempty" name:"Type"`

	// Count
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// Size
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// Timestamp
	Updatetime *int64 `json:"Updatetime,omitempty" name:"Updatetime"`
}

type CleanUpInstanceRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *CleanUpInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CleanUpInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CleanUpInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task ID
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CleanUpInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CleanUpInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ClearInstanceRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Redis instance password (this parameter is not required for password-free instances but for password-enabled instances)
	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *ClearInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ClearInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ClearInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task ID
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ClearInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ClearInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CommandTake struct {

	// Command
	Cmd *string `json:"Cmd,omitempty" name:"Cmd"`

	// Duration
	Took *int64 `json:"Took,omitempty" name:"Took"`
}

type CreateInstanceAccountRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Sub-account name
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// Sub-account password
	AccountPassword *string `json:"AccountPassword,omitempty" name:"AccountPassword"`

	// Routing policy. Enter `master` for primary node or `replication` for secondary node
	ReadonlyPolicy []*string `json:"ReadonlyPolicy,omitempty" name:"ReadonlyPolicy" list`

	// Read/write policy. Valid values: r (read-only), rw (read/write).
	Privilege *string `json:"Privilege,omitempty" name:"Privilege"`

	// Sub-account description information
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *CreateInstanceAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateInstanceAccountRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceAccountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task ID
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateInstanceAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateInstanceAccountResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateInstancesRequest struct {
	*tchttp.BaseRequest

	// AZ ID of instance
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// Instance type. Valid values: 2 (Redis 2.8 memory edition in standard architecture), 3 (Redis 3.2 memory edition in standard architecture), 4 (CKV 3.2 memory edition in standard architecture), 6 (Redis 4.0 memory edition in standard architecture), 7 (Redis 4.0 memory edition in cluster architecture), 8 (Redis 5.0 memory edition in standard architecture), 9 (Redis 5.0 memory edition in cluster architecture).
	TypeId *uint64 `json:"TypeId,omitempty" name:"TypeId"`

	// Instance capacity in MB. The actual value is subject to the specifications returned by the purchasable specification querying API |
	MemSize *uint64 `json:"MemSize,omitempty" name:"MemSize"`

	// Number of instances. The actual quantity purchasable at a time is subject to the specifications returned by the purchasable specification querying API
	GoodsNum *uint64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// Length of purchase in months, which is required when creating a monthly subscribed instances. Value range: [1,2,3,4,5,6,7,8,9,10,11,12,24,36]. For pay-as-you-go instances, enter 1
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// Billing method. 0: pay as you go
	BillingMode *int64 `json:"BillingMode,omitempty" name:"BillingMode"`

	// Instance password. It can contain 8-30 characters and must contain at least two of the following types of characters: lowercase letters, uppercase letters, digits, and special symbols (()`~!@#$%^&*-+=_|{}[]:;<>,.?/). It cannot stat with the symbol (/).
	Password *string `json:"Password,omitempty" name:"Password"`

	// VPC ID such as vpc-sad23jfdfk. If this parameter is not passed in, the basic network will be selected by default. Please use the VPC list querying API to query.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// In a basic network, subnetId is invalid. In a VPC subnet, the value is the subnet ID, such as subnet-fdj24n34j2
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Project ID. The value is subject to the projectId returned by user account > user account-related querying APIs > project list
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Auto-renewal flag. 0: default status (manual renewal); 1: auto-renewal enabled; 2: auto-renewal disabled
	AutoRenew *uint64 `json:"AutoRenew,omitempty" name:"AutoRenew"`

	// Array of security group IDs
	SecurityGroupIdList []*string `json:"SecurityGroupIdList,omitempty" name:"SecurityGroupIdList" list`

	// User-defined port. If this parameter is left empty, 6379 will be used by default. Value range: [1024,65535]
	VPort *uint64 `json:"VPort,omitempty" name:"VPort"`

	// Number of shards in an instance. This parameter is required for cluster edition instances. Valid values: 3, 5, 8, 12, 16, 24, 32, 64, 96, 128.
	RedisShardNum *int64 `json:"RedisShardNum,omitempty" name:"RedisShardNum"`

	// Number of replicas in an instance. Redis 2.8 standard edition and CKV standard edition support 1 replica. Standard/cluster edition 4.0 and 5.0 support 1-5 replicas.
	RedisReplicasNum *int64 `json:"RedisReplicasNum,omitempty" name:"RedisReplicasNum"`

	// Whether to support read-only replicas. Neither Redis 2.8 standard edition nor CKV standard edition supports read-only replicas. Read/write separation will be automatically enabled for an instance after it enables read-only replicas. Write requests will be directed to the primary node and read requests will be distributed on secondary nodes. To enable read-only replicas, we recommend you create 2 or more replicas.
	ReplicasReadonly *bool `json:"ReplicasReadonly,omitempty" name:"ReplicasReadonly"`

	// Instance name. It contains only letters, digits, underscores, and dashes with a length of up to 60 characters.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Whether to support the password-free feature. Valid values: true (password-free instance), false (password-enabled instance). Default value: false. Only instances in a VPC support the password-free access.
	NoAuth *bool `json:"NoAuth,omitempty" name:"NoAuth"`
}

func (r *CreateInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Transaction ID
		DealId *string `json:"DealId,omitempty" name:"DealId"`

		// Instance ID (this field is during beta test and is not displayed in some regions)
		InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DelayDistribution struct {

	// Distribution ladder
	Ladder *int64 `json:"Ladder,omitempty" name:"Ladder"`

	// Size
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// Modification time
	Updatetime *int64 `json:"Updatetime,omitempty" name:"Updatetime"`
}

type DeleteInstanceAccountRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Sub-account name
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`
}

func (r *DeleteInstanceAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteInstanceAccountRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteInstanceAccountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task ID
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteInstanceAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteInstanceAccountResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAutoBackupConfigRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeAutoBackupConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAutoBackupConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAutoBackupConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Backup type. Auto backup type: 1 "scheduled rollback"
		AutoBackupType *int64 `json:"AutoBackupType,omitempty" name:"AutoBackupType"`

		// Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday.
		WeekDays []*string `json:"WeekDays,omitempty" name:"WeekDays" list`

		// Time period.
		TimePeriod *string `json:"TimePeriod,omitempty" name:"TimePeriod"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAutoBackupConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAutoBackupConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupUrlRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Backup ID, which can be queried through the `DescribeInstanceBackups` API
	BackupId *string `json:"BackupId,omitempty" name:"BackupId"`
}

func (r *DescribeBackupUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBackupUrlRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupUrlResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Download address on the public network (valid for 6 hours)
		DownloadUrl []*string `json:"DownloadUrl,omitempty" name:"DownloadUrl" list`

		// Download address on the private network (valid for 6 hours)
		InnerDownloadUrl []*string `json:"InnerDownloadUrl,omitempty" name:"InnerDownloadUrl" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBackupUrlResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// Database engine name: mariadb, cdb, cynosdb, dcdb, redis, mongodb, etc.
	Product *string `json:"Product,omitempty" name:"Product"`

	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB Console.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBSecurityGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Security group rules.
		Groups []*SecurityGroup `json:"Groups,omitempty" name:"Groups" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBSecurityGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceAccountRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Number of entries per page
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Page offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeInstanceAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceAccountRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceAccountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Account details
	// Note: This field may return null, indicating that no valid values can be obtained.
		Accounts []*Account `json:"Accounts,omitempty" name:"Accounts" list`

		// Number of accounts
	// Note: This field may return null, indicating that no valid values can be obtained.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceAccountResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceBackupsRequest struct {
	*tchttp.BaseRequest

	// ID of the instance to be operated on, which can be obtained through the `InstanceId` field in the return value of the DescribeInstance API.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance list size. Default value: 20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset, which is an integral multiple of `Limit`
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Start time in the format of yyyy-MM-dd HH:mm:ss, such as 2017-02-08 16:46:34. This parameter is used to query the list of instance backups started during the [beginTime, endTime] range.
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// End time in the format of yyyy-MM-dd HH:mm:ss, such as 2017-02-08 19:09:26. This parameter is used to query the list of instance backups started during the [beginTime, endTime] range.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 1: backup in process; 2: backing up normally; 3: converting from backup to RDB file; 4: RDB conversion completed; -1: backup expired; -2: backup deleted.
	Status []*int64 `json:"Status,omitempty" name:"Status" list`
}

func (r *DescribeInstanceBackupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceBackupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceBackupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of backups
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Array of instance backups
		BackupSet []*RedisBackupSet `json:"BackupSet,omitempty" name:"BackupSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceBackupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceBackupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceDTSInfoRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceDTSInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceDTSInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceDTSInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// DTS task ID
	// Note: this field may return null, indicating that no valid values can be obtained.
		JobId *string `json:"JobId,omitempty" name:"JobId"`

		// DTS task name
	// Note: this field may return null, indicating that no valid values can be obtained.
		JobName *string `json:"JobName,omitempty" name:"JobName"`

		// Task status. Valid values: 1 (Creating), 3 (Checking), 4 (CheckPass), 5 (CheckNotPass), 7 (Running), 8 (ReadyComplete), 9 (Success), 10 (Failed), 11 (Stopping), 12 (Completing)
	// Note: this field may return null, indicating that no valid values can be obtained.
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// Status description
	// Note: this field may return null, indicating that no valid values can be obtained.
		StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

		// Synchronization latency in bytes
	// Note: this field may return null, indicating that no valid values can be obtained.
		Offset *int64 `json:"Offset,omitempty" name:"Offset"`

		// Disconnection time
	// Note: this field may return null, indicating that no valid values can be obtained.
		CutDownTime *string `json:"CutDownTime,omitempty" name:"CutDownTime"`

		// Source instance information
	// Note: this field may return null, indicating that no valid values can be obtained.
		SrcInfo *DescribeInstanceDTSInstanceInfo `json:"SrcInfo,omitempty" name:"SrcInfo"`

		// Target instance information
	// Note: this field may return null, indicating that no valid values can be obtained.
		DstInfo *DescribeInstanceDTSInstanceInfo `json:"DstInfo,omitempty" name:"DstInfo"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceDTSInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceDTSInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceDTSInstanceInfo struct {

	// Region ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// Instance ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Repository ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	SetId *int64 `json:"SetId,omitempty" name:"SetId"`

	// Availability zone ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// Instance type
	// Note: this field may return null, indicating that no valid values can be obtained.
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// Instance name
	// Note: this field may return null, indicating that no valid values can be obtained.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Instance access address
	// Note: this field may return null, indicating that no valid values can be obtained.
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Status
	// Note: this field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type DescribeInstanceDealDetailRequest struct {
	*tchttp.BaseRequest

	// Array of order IDs
	DealIds []*string `json:"DealIds,omitempty" name:"DealIds" list`
}

func (r *DescribeInstanceDealDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceDealDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceDealDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Order details
		DealDetails []*TradeDealDetail `json:"DealDetails,omitempty" name:"DealDetails" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceDealDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceDealDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorBigKeyRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Request type. 1: string type; 2: all types
	ReqType *int64 `json:"ReqType,omitempty" name:"ReqType"`

	// Time, such as "20190219"
	Date *string `json:"Date,omitempty" name:"Date"`
}

func (r *DescribeInstanceMonitorBigKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceMonitorBigKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorBigKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Big key details
		Data []*BigKeyInfo `json:"Data,omitempty" name:"Data" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceMonitorBigKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceMonitorBigKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorBigKeySizeDistRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Time, such as "20190219"
	Date *string `json:"Date,omitempty" name:"Date"`
}

func (r *DescribeInstanceMonitorBigKeySizeDistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceMonitorBigKeySizeDistRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorBigKeySizeDistResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Big key size distribution details
		Data []*DelayDistribution `json:"Data,omitempty" name:"Data" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceMonitorBigKeySizeDistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceMonitorBigKeySizeDistResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorBigKeyTypeDistRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Time, such as "20190219"
	Date *string `json:"Date,omitempty" name:"Date"`
}

func (r *DescribeInstanceMonitorBigKeyTypeDistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceMonitorBigKeyTypeDistRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorBigKeyTypeDistResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Big key type distribution details
		Data []*BigKeyTypeInfo `json:"Data,omitempty" name:"Data" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceMonitorBigKeyTypeDistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceMonitorBigKeyTypeDistResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorHotKeyRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Time span. 1: real time; 2: past 30 minutes; 3: past 6 hours; 4: past 24 hours
	SpanType *int64 `json:"SpanType,omitempty" name:"SpanType"`
}

func (r *DescribeInstanceMonitorHotKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceMonitorHotKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorHotKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Hot key details
		Data []*HotKeyInfo `json:"Data,omitempty" name:"Data" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceMonitorHotKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceMonitorHotKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorSIPRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceMonitorSIPRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceMonitorSIPRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorSIPResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Access source information
		Data []*SourceInfo `json:"Data,omitempty" name:"Data" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceMonitorSIPResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceMonitorSIPResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorTookDistRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Time, such as "20190219"
	Date *string `json:"Date,omitempty" name:"Date"`

	// Time span. 1: real time; 2: last 30 minutes; 3: last 6 hours; 4: last 24 hours
	SpanType *int64 `json:"SpanType,omitempty" name:"SpanType"`
}

func (r *DescribeInstanceMonitorTookDistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceMonitorTookDistRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorTookDistResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Latency distribution information
		Data []*DelayDistribution `json:"Data,omitempty" name:"Data" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceMonitorTookDistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceMonitorTookDistResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorTopNCmdRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Time span. 1: real time; 2: last 30 minutes; 3: last 6 hours; 4: last 24 hours
	SpanType *int64 `json:"SpanType,omitempty" name:"SpanType"`
}

func (r *DescribeInstanceMonitorTopNCmdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceMonitorTopNCmdRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorTopNCmdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Access command information
		Data []*SourceCommand `json:"Data,omitempty" name:"Data" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceMonitorTopNCmdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceMonitorTopNCmdResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorTopNCmdTookRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Time span. 1: real time; 2: last 30 minutes; 3: last 6 hours; 4: last 24 hours
	SpanType *int64 `json:"SpanType,omitempty" name:"SpanType"`
}

func (r *DescribeInstanceMonitorTopNCmdTookRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceMonitorTopNCmdTookRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorTopNCmdTookResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Duration details
		Data []*CommandTake `json:"Data,omitempty" name:"Data" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceMonitorTopNCmdTookResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceMonitorTopNCmdTookResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceParamRecordsRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Number of entries per page
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset, which is an integral multiple of `Limit`
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeInstanceParamRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceParamRecordsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceParamRecordsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of modifications.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Information of modifications.
		InstanceParamHistory []*InstanceParamHistory `json:"InstanceParamHistory,omitempty" name:"InstanceParamHistory" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceParamRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceParamRecordsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceParamsRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceParamsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceParamsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceParamsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of instance parameters
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Instance parameter in Enum type
		InstanceEnumParam []*InstanceEnumParam `json:"InstanceEnumParam,omitempty" name:"InstanceEnumParam" list`

		// Instance parameter in Integer type
		InstanceIntegerParam []*InstanceIntegerParam `json:"InstanceIntegerParam,omitempty" name:"InstanceIntegerParam" list`

		// Instance parameter in Char type
		InstanceTextParam []*InstanceTextParam `json:"InstanceTextParam,omitempty" name:"InstanceTextParam" list`

		// Instance parameter in Multi type
		InstanceMultiParam []*InstanceMultiParam `json:"InstanceMultiParam,omitempty" name:"InstanceMultiParam" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceParamsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceParamsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceSecurityGroupRequest struct {
	*tchttp.BaseRequest

	// Instance list
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *DescribeInstanceSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceSecurityGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Security group information of an instance
		InstanceSecurityGroupsDetail []*InstanceSecurityGroupDetail `json:"InstanceSecurityGroupsDetail,omitempty" name:"InstanceSecurityGroupsDetail" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceSecurityGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceShardsRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Whether to filter out the secondary node information
	FilterSlave *bool `json:"FilterSlave,omitempty" name:"FilterSlave"`
}

func (r *DescribeInstanceShardsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceShardsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceShardsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Information list of instance shards
		InstanceShards []*InstanceClusterShard `json:"InstanceShards,omitempty" name:"InstanceShards" list`

		// Total number of instance shard nodes
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceShardsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceShardsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest

	// Instance list size. Default value: 20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset, which is an integral multiple of `Limit`
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Instance ID, such as crs-6ubhgouj
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Enumerated values: projectId, createtime, instancename, type, curDeadline
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 1: reverse; 0: sequential; reverse by default
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// Array of VPC IDs such as 47525. The array subscript starts from 0. If this parameter is not passed in, the basic network will be selected by default
	VpcIds []*string `json:"VpcIds,omitempty" name:"VpcIds" list`

	// Array of subnet IDs such as 56854. The array subscript starts from 0
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds" list`

	// Array of project IDs. The array subscript starts from 0
	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds" list`

	// ID of the instance to be searched for.
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Array of VPC IDs such as vpc-sad23jfdfk. The array subscript starts from 0. If this parameter is not passed in, the basic network will be selected by default
	UniqVpcIds []*string `json:"UniqVpcIds,omitempty" name:"UniqVpcIds" list`

	// Array of subnet IDs such as subnet-fdj24n34j2. The array subscript starts from 0
	UniqSubnetIds []*string `json:"UniqSubnetIds,omitempty" name:"UniqSubnetIds" list`

	// Region ID, which has already been disused. The corresponding region can be queried through the common parameter `Region`
	RegionIds []*int64 `json:"RegionIds,omitempty" name:"RegionIds" list`

	// Instance status. 0: to be initialized; 1: in process; 2: running; -2: isolated; -3: to be deleted
	Status []*int64 `json:"Status,omitempty" name:"Status" list`

	// Type edition. 1: standalone edition; 2: primary-secondary edition; 3: cluster edition
	TypeVersion *int64 `json:"TypeVersion,omitempty" name:"TypeVersion"`

	// Engine information: Redis-2.8, Redis-4.0, CKV
	EngineName *string `json:"EngineName,omitempty" name:"EngineName"`

	// Renewal mode. 0: default status (manual renewal); 1: auto-renewal enabled; 2: auto-renewal disabled
	AutoRenew []*int64 `json:"AutoRenew,omitempty" name:"AutoRenew" list`

	// Billing method. postpaid: pay-as-you-go; prepaid: monthly subscription
	BillingMode *string `json:"BillingMode,omitempty" name:"BillingMode"`

	// Instance type. 1: legacy Redis Cluster Edition, 2: Redis 2.8 Master-Slave Edition, 3: CKV Master-Slave Edition, 4: CKV Cluster Edition, 5: Redis 2.8 Standalone Edition, 6: Redis 4.0 Master-Slave Edition, 7: Redis 4.0 Cluster Edition, 8: Redis 5.0 Master-Slave Edition, 9: Redis 5.0 Cluster Edition,
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// Search keywords, which can be instance ID, instance name, or complete IP
	SearchKeys []*string `json:"SearchKeys,omitempty" name:"SearchKeys" list`

	// Internal parameter, which can be ignored
	TypeList []*int64 `json:"TypeList,omitempty" name:"TypeList" list`
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

		// Number of instances
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of instance details
		InstanceSet []*InstanceSet `json:"InstanceSet,omitempty" name:"InstanceSet" list`

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

type DescribeProductInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeProductInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProductInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProductInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Sale information of a region
		RegionSet []*RegionConf `json:"RegionSet,omitempty" name:"RegionSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProductInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectSecurityGroupRequest struct {
	*tchttp.BaseRequest

	// 0: default project; -1: all projects; >0: specified project
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Security group ID
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
}

func (r *DescribeProjectSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProjectSecurityGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Security group of a project
		SecurityGroupDetails []*SecurityGroupDetail `json:"SecurityGroupDetails,omitempty" name:"SecurityGroupDetails" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProjectSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProjectSecurityGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// Database engine name: mariadb, cdb, cynosdb, dcdb, redis, mongodb.
	Product *string `json:"Product,omitempty" name:"Product"`

	// Project ID.
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Offset.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// The number of security groups to be pulled.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Search criteria. You can enter a security group ID or name.
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
}

func (r *DescribeProjectSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProjectSecurityGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Security group rules.
		Groups []*SecurityGroup `json:"Groups,omitempty" name:"Groups" list`

		// Total number of the security groups meeting the condition.
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProjectSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProjectSecurityGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Start time
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Slow log threshold in microseconds
	MinQueryTime *int64 `json:"MinQueryTime,omitempty" name:"MinQueryTime"`

	// Number of entries per page
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset, which is an integral multiple of `Limit`
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeSlowLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSlowLogRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of slow logs
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Slow log details
		InstanceSlowlogDetail []*InstanceSlowlogDetail `json:"InstanceSlowlogDetail,omitempty" name:"InstanceSlowlogDetail" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSlowLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSlowLogResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskInfoRequest struct {
	*tchttp.BaseRequest

	// Task ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeTaskInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task status. preparing: to be run; running: running; succeed: succeeded; failed: failed; error: running error.
		Status *string `json:"Status,omitempty" name:"Status"`

		// Task start time
		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

		// Task type
		TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

		// Instance ID
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// Task message, which is displayed in case of an error. It will be blank if the status is running or succeeded
		TaskMessage *string `json:"TaskMessage,omitempty" name:"TaskMessage"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskListRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Number of entries per page
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset, which is an integral multiple of `Limit` (rounded down automatically)
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Project ID
	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds" list`

	// Task type
	TaskTypes []*string `json:"TaskTypes,omitempty" name:"TaskTypes" list`

	// Start time
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Task status
	TaskStatus []*int64 `json:"TaskStatus,omitempty" name:"TaskStatus" list`
}

func (r *DescribeTaskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of tasks
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Task details
		Tasks []*TaskInfoDetail `json:"Tasks,omitempty" name:"Tasks" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DestroyPostpaidInstanceRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DestroyPostpaidInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DestroyPostpaidInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DestroyPostpaidInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task ID
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DestroyPostpaidInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DestroyPostpaidInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DestroyPrepaidInstanceRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DestroyPrepaidInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DestroyPrepaidInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DestroyPrepaidInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Order ID
		DealId *string `json:"DealId,omitempty" name:"DealId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DestroyPrepaidInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DestroyPrepaidInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableReplicaReadonlyRequest struct {
	*tchttp.BaseRequest

	// Serial ID of an instance
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DisableReplicaReadonlyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableReplicaReadonlyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableReplicaReadonlyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// ERROR: failure; OK: success
		Status *string `json:"Status,omitempty" name:"Status"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableReplicaReadonlyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableReplicaReadonlyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisassociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// Database engine name: mariadb, cdb, cynosdb, dcdb, redis, mongodb, etc.
	Product *string `json:"Product,omitempty" name:"Product"`

	// Security group ID.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// Instance ID list, which is an array of one or more instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *DisassociateSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisassociateSecurityGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisassociateSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisassociateSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisassociateSecurityGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableReplicaReadonlyRequest struct {
	*tchttp.BaseRequest

	// Serial ID of an instance
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Account routing policy. If `master` or `replication` is entered, it means to route to the primary or secondary node; if this is left blank, it means to write into the primary node and read from the secondary node by default
	ReadonlyPolicy []*string `json:"ReadonlyPolicy,omitempty" name:"ReadonlyPolicy" list`
}

func (r *EnableReplicaReadonlyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableReplicaReadonlyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableReplicaReadonlyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// ERROR: erroneous; OK: correct.
		Status *string `json:"Status,omitempty" name:"Status"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableReplicaReadonlyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableReplicaReadonlyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type HotKeyInfo struct {

	// Hot key
	Key *string `json:"Key,omitempty" name:"Key"`

	// Type
	Type *string `json:"Type,omitempty" name:"Type"`

	// Count
	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type Inbound struct {

	// Policy. Valid values: ACCEPT, DROP.
	Action *string `json:"Action,omitempty" name:"Action"`

	// All the addresses that the address group ID represents.
	AddressModule *string `json:"AddressModule,omitempty" name:"AddressModule"`

	// Source IP or IP address range, such as 192.168.0.0/16.
	CidrIp *string `json:"CidrIp,omitempty" name:"CidrIp"`

	// Description.
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// Network protocol, such as UDP and TCP, etc.
	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`

	// Port.
	PortRange *string `json:"PortRange,omitempty" name:"PortRange"`

	// All the protocols and ports that the service group ID represents.
	ServiceModule *string `json:"ServiceModule,omitempty" name:"ServiceModule"`

	// All the addresses that the security group ID represents.
	Id *string `json:"Id,omitempty" name:"Id"`
}

type InstanceClusterNode struct {

	// Node name
	Name *string `json:"Name,omitempty" name:"Name"`

	// ID of the runtime node of an instance
	RunId *string `json:"RunId,omitempty" name:"RunId"`

	// Cluster role. 0: primary; 1: secondary
	Role *int64 `json:"Role,omitempty" name:"Role"`

	// Node status. 0: readwrite; 1: read; 2: backup
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Service status. 0: down; 1: on
	Connected *int64 `json:"Connected,omitempty" name:"Connected"`

	// Node creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Node deactivation time
	DownTime *string `json:"DownTime,omitempty" name:"DownTime"`

	// Distribution of node slots
	Slots *string `json:"Slots,omitempty" name:"Slots"`

	// Distribution of node keys
	Keys *int64 `json:"Keys,omitempty" name:"Keys"`

	// Node QPS
	Qps *int64 `json:"Qps,omitempty" name:"Qps"`

	// QPS slope of a node
	QpsSlope *float64 `json:"QpsSlope,omitempty" name:"QpsSlope"`

	// Node storage
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// Storage slope of a node
	StorageSlope *float64 `json:"StorageSlope,omitempty" name:"StorageSlope"`
}

type InstanceClusterShard struct {

	// Shard node name
	ShardName *string `json:"ShardName,omitempty" name:"ShardName"`

	// Shard node ID
	ShardId *string `json:"ShardId,omitempty" name:"ShardId"`

	// Role
	Role *int64 `json:"Role,omitempty" name:"Role"`

	// Number of keys
	Keys *int64 `json:"Keys,omitempty" name:"Keys"`

	// Slot information
	Slots *string `json:"Slots,omitempty" name:"Slots"`

	// Storage capacity
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// Capacity slope
	StorageSlope *float64 `json:"StorageSlope,omitempty" name:"StorageSlope"`

	// ID of the runtime node of an instance
	Runid *string `json:"Runid,omitempty" name:"Runid"`

	// Service status. 0: down; 1: on
	Connected *int64 `json:"Connected,omitempty" name:"Connected"`
}

type InstanceEnumParam struct {

	// Parameter name
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`

	// Parameter type: Enum
	ValueType *string `json:"ValueType,omitempty" name:"ValueType"`

	// Whether restart is required after a modification is made. Value range: true, false
	NeedRestart *string `json:"NeedRestart,omitempty" name:"NeedRestart"`

	// Default value of the parameter
	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`

	// Current value of a parameter
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`

	// Parameter description
	Tips *string `json:"Tips,omitempty" name:"Tips"`

	// Value range of a parameter
	EnumValue []*string `json:"EnumValue,omitempty" name:"EnumValue" list`

	// Parameter status. 1: modifying; 2: modified
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type InstanceIntegerParam struct {

	// Parameter name
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`

	// Parameter type: Integer
	ValueType *string `json:"ValueType,omitempty" name:"ValueType"`

	// Whether restart is required after a modification is made. Value range: true, false
	NeedRestart *string `json:"NeedRestart,omitempty" name:"NeedRestart"`

	// Default value of the parameter
	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`

	// Current value of a parameter
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`

	// Parameter description
	Tips *string `json:"Tips,omitempty" name:"Tips"`

	// Minimum value of a parameter
	Min *string `json:"Min,omitempty" name:"Min"`

	// Maximum value of a parameter
	Max *string `json:"Max,omitempty" name:"Max"`

	// Parameter status. 1: modifying; 2: modified
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type InstanceMultiParam struct {

	// Parameter name
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`

	// Parameter type: Multi
	ValueType *string `json:"ValueType,omitempty" name:"ValueType"`

	// Whether restart is required after a modification is made. Value range: true, false
	NeedRestart *string `json:"NeedRestart,omitempty" name:"NeedRestart"`

	// Default value of the parameter
	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`

	// Current value of a parameter
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`

	// Parameter description
	Tips *string `json:"Tips,omitempty" name:"Tips"`

	// Parameter description
	EnumValue []*string `json:"EnumValue,omitempty" name:"EnumValue" list`

	// Parameter status. 1: modifying; 2: modified
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type InstanceNode struct {

	// Id
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// Node details
	InstanceClusterNode []*InstanceClusterNode `json:"InstanceClusterNode,omitempty" name:"InstanceClusterNode" list`
}

type InstanceParam struct {

	// Sets a parameter name
	Key *string `json:"Key,omitempty" name:"Key"`

	// Sets a parameter value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type InstanceParamHistory struct {

	// Parameter name
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`

	// Value before modification
	PreValue *string `json:"PreValue,omitempty" name:"PreValue"`

	// Value after modification
	NewValue *string `json:"NewValue,omitempty" name:"NewValue"`

	// Status. 1: modifying parameter configuration; 2: parameter configuration modified successfully; 3: failed to modify parameter configuration
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Modification time
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type InstanceSecurityGroupDetail struct {

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Security group information
	SecurityGroupDetails []*SecurityGroupDetail `json:"SecurityGroupDetails,omitempty" name:"SecurityGroupDetails" list`
}

type InstanceSet struct {

	// Instance name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// User's Appid
	Appid *int64 `json:"Appid,omitempty" name:"Appid"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Region ID. 1: Guangzhou; 4: Shanghai; 5: Hong Kong, China; 6: Toronto; 7: Shanghai Finance; 8: Beijing; 9: Singapore; 11: Shenzhen Finance; 15: West US (Silicon Valley); 16: Chengdu; 17: Germany; 18: South Korea; 19: Chongqing; 21: India; 22: East US (Virginia); 23: Thailand; 24: Russia; 25: Japan
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// Region ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// VPC ID, such as 75101
	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`

	// VPC subnet ID, such as 46315
	SubnetId *int64 `json:"SubnetId,omitempty" name:"SubnetId"`

	// Current instance status. 0: to be initialized; 1: instance in process; 2: instance running; -2: instance isolated; -3: instance to be deleted
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Instance VIP
	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`

	// Port number of an instance
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Instance creation time
	Createtime *string `json:"Createtime,omitempty" name:"Createtime"`

	// Instance capacity in MB
	Size *float64 `json:"Size,omitempty" name:"Size"`

	// This field has been disused
	SizeUsed *float64 `json:"SizeUsed,omitempty" name:"SizeUsed"`

	// Instance type. 1: Redis 2.8 cluster edition; 2: Redis 2.8 primary-secondary edition; 3: CKV primary-secondary edition (Redis 3.2); 4: CKV cluster edition (Redis 3.2); 5: Redis 2.8 standalone edition; 6: Redis 4.0 primary-secondary edition; 7: Redis 4.0 cluster edition
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// Whether to set the auto-renewal flag for an instance. 1: auto-renewal set; 0: auto-renewal not set
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// Instance expiration time
	DeadlineTime *string `json:"DeadlineTime,omitempty" name:"DeadlineTime"`

	// Engine: Redis community edition, Tencent Cloud CKV
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// Product type: Redis 2.8 cluster edition, Redis 2.8 primary-secondary edition, Redis 3.2 primary-secondary edition (CKV primary-secondary edition), Redis 3.2 cluster edition (CKV cluster edition), Redis 2.8 standalone edition, Redis 4.0 cluster edition
	ProductType *string `json:"ProductType,omitempty" name:"ProductType"`

	// VPC ID, such as vpc-fk33jsf43kgv
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// VPC subnet ID, such as subnet-fd3j6l35mm0
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// Billing method. 0: pay-as-you-go; 1: monthly subscription
	BillingMode *int64 `json:"BillingMode,omitempty" name:"BillingMode"`

	// Description of an instance status, such as "instance running"
	InstanceTitle *string `json:"InstanceTitle,omitempty" name:"InstanceTitle"`

	// Scheduled deactivation time
	OfflineTime *string `json:"OfflineTime,omitempty" name:"OfflineTime"`

	// Sub-status returned for an instance in process
	SubStatus *int64 `json:"SubStatus,omitempty" name:"SubStatus"`

	// Anti-affinity tag
	Tags []*string `json:"Tags,omitempty" name:"Tags" list`

	// Instance node information
	InstanceNode []*InstanceNode `json:"InstanceNode,omitempty" name:"InstanceNode" list`

	// Shard size
	RedisShardSize *int64 `json:"RedisShardSize,omitempty" name:"RedisShardSize"`

	// Number of shards
	RedisShardNum *int64 `json:"RedisShardNum,omitempty" name:"RedisShardNum"`

	// Number of replicas
	RedisReplicasNum *int64 `json:"RedisReplicasNum,omitempty" name:"RedisReplicasNum"`

	// Billing ID
	PriceId *int64 `json:"PriceId,omitempty" name:"PriceId"`

	// Isolation time
	CloseTime *string `json:"CloseTime,omitempty" name:"CloseTime"`

	// Read weight of a secondary node
	SlaveReadWeight *int64 `json:"SlaveReadWeight,omitempty" name:"SlaveReadWeight"`

	// Instance tag information
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceTags []*InstanceTagInfo `json:"InstanceTags,omitempty" name:"InstanceTags" list`

	// Project name
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// Whether an instance is password-free. true: yes; false: no
	// Note: This field may return null, indicating that no valid values can be obtained.
	NoAuth *bool `json:"NoAuth,omitempty" name:"NoAuth"`

	// Number of client connections
	// Note: this field may return null, indicating that no valid values can be obtained.
	ClientLimit *int64 `json:"ClientLimit,omitempty" name:"ClientLimit"`

	// DTS status (internal parameter, which can be ignored)
	// Note: this field may return null, indicating that no valid values can be obtained.
	DtsStatus *int64 `json:"DtsStatus,omitempty" name:"DtsStatus"`

	// Upper shard bandwidth limit in MB
	// Note: this field may return null, indicating that no valid values can be obtained.
	NetLimit *int64 `json:"NetLimit,omitempty" name:"NetLimit"`

	// Password-free instance flag (internal parameter, which can be ignored)
	// Note: this field may return null, indicating that no valid values can be obtained.
	PasswordFree *int64 `json:"PasswordFree,omitempty" name:"PasswordFree"`

	// Read-only instance flag (internal parameter, which can be ignored)
	// Note: this field may return null, indicating that no valid values can be obtained.
	ReadOnly *int64 `json:"ReadOnly,omitempty" name:"ReadOnly"`

	// Internal parameter, which can be ignored
	// Note: this field may return null, indicating that no valid values can be obtained.
	Vip6 *string `json:"Vip6,omitempty" name:"Vip6"`

	// Internal parameter, which can be ignored
	// Note: this field may return null, indicating that no valid values can be obtained.
	RemainBandwidthDuration *string `json:"RemainBandwidthDuration,omitempty" name:"RemainBandwidthDuration"`
}

type InstanceSlowlogDetail struct {

	// Slow log duration
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`

	// Client address
	Client *string `json:"Client,omitempty" name:"Client"`

	// Command
	Command *string `json:"Command,omitempty" name:"Command"`

	// Command line details
	CommandLine *string `json:"CommandLine,omitempty" name:"CommandLine"`

	// Execution duration
	ExecuteTime *string `json:"ExecuteTime,omitempty" name:"ExecuteTime"`
}

type InstanceTagInfo struct {

	// Tag key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type InstanceTextParam struct {

	// Parameter name
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`

	// Parameter type: Text
	ValueType *string `json:"ValueType,omitempty" name:"ValueType"`

	// Whether restart is required after a modification is made. Value range: true, false
	NeedRestart *string `json:"NeedRestart,omitempty" name:"NeedRestart"`

	// Default value of the parameter
	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`

	// Current value of a parameter
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`

	// Parameter description
	Tips *string `json:"Tips,omitempty" name:"Tips"`

	// Value range of a parameter
	TextValue []*string `json:"TextValue,omitempty" name:"TextValue" list`

	// Parameter status. 1: modifying; 2: modified
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type ManualBackupInstanceRequest struct {
	*tchttp.BaseRequest

	// ID of the instance to be operated on, which can be obtained through the `InstanceId` field in the return value of the DescribeInstance API.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Backup remarks
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ManualBackupInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ManualBackupInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ManualBackupInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task ID
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ManualBackupInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ManualBackupInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModfiyInstancePasswordRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Old password of an instance
	OldPassword *string `json:"OldPassword,omitempty" name:"OldPassword"`

	// New password of an instance
	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *ModfiyInstancePasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModfiyInstancePasswordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModfiyInstancePasswordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task ID
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModfiyInstancePasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModfiyInstancePasswordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAutoBackupConfigRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Date. Value range: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday
	WeekDays []*string `json:"WeekDays,omitempty" name:"WeekDays" list`

	// Time period. Value range: 00:00-01:00, 01:00-02:00...... 23:00-00:00
	TimePeriod *string `json:"TimePeriod,omitempty" name:"TimePeriod"`

	// Auto backup type: 1 "scheduled rollback"
	AutoBackupType *int64 `json:"AutoBackupType,omitempty" name:"AutoBackupType"`
}

func (r *ModifyAutoBackupConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAutoBackupConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAutoBackupConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Auto backup type: 1 "scheduled rollback"
		AutoBackupType *int64 `json:"AutoBackupType,omitempty" name:"AutoBackupType"`

		// Date. Value range: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday.
		WeekDays []*string `json:"WeekDays,omitempty" name:"WeekDays" list`

		// Time period. Value range: 00:00-01:00, 01:00-02:00...... 23:00-00:00
		TimePeriod *string `json:"TimePeriod,omitempty" name:"TimePeriod"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAutoBackupConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAutoBackupConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// Database engine name: mariadb, cdb, cynosdb, dcdb, redis, mongodb, etc.
	Product *string `json:"Product,omitempty" name:"Product"`

	// The ID list of the security groups to be modified, which is an array of one or more security group IDs.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`

	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB Console.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ModifyDBInstanceSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBInstanceSecurityGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBInstanceSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBInstanceSecurityGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceAccountRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Sub-account name. If the root account is to be modified, enter `root`
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// Sub-account password
	AccountPassword *string `json:"AccountPassword,omitempty" name:"AccountPassword"`

	// Sub-account description information
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Sub-account routing policy. Enter `master` to route to the primary node or `slave` to route to the secondary node
	ReadonlyPolicy []*string `json:"ReadonlyPolicy,omitempty" name:"ReadonlyPolicy" list`

	// Sub-account read/write policy. Enter `r` for read-only, `w` for write-only, or `rw` for read/write
	Privilege *string `json:"Privilege,omitempty" name:"Privilege"`

	// true: make the root account password-free. This is applicable to root accounts only; sub-accounts cannot be made password-free
	NoAuth *bool `json:"NoAuth,omitempty" name:"NoAuth"`
}

func (r *ModifyInstanceAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstanceAccountRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceAccountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task ID
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstanceAccountResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceParamsRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// List of instance parameters modified
	InstanceParams []*InstanceParam `json:"InstanceParams,omitempty" name:"InstanceParams" list`
}

func (r *ModifyInstanceParamsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstanceParamsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceParamsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Whether a modification is successfully made.
		Changed *bool `json:"Changed,omitempty" name:"Changed"`

		// Task ID
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceParamsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstanceParamsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceRequest struct {
	*tchttp.BaseRequest

	// Instance modification type. rename: rename an instance; modifyProject: modify the project of an instance; modifyAutoRenew: modify the auto-renewal flag of an instance
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// Instance ID
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// New name of instance
	InstanceNames []*string `json:"InstanceNames,omitempty" name:"InstanceNames" list`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Auto-renewal flag. 0: default status (manual renewal), 1: auto-renewal enabled, 2: auto-renewal disabled
	AutoRenews []*int64 `json:"AutoRenews,omitempty" name:"AutoRenews" list`

	// Disused
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Disused
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Disused
	AutoRenew *int64 `json:"AutoRenew,omitempty" name:"AutoRenew"`
}

func (r *ModifyInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNetworkConfigRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Operation type. changeVip: modify the VIP of an instance; changeVpc: modify the subnet of an instance; changeBaseToVpc: change from basic network to VPC
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// VIP address, which is required for the `changeVip` operation. If this parameter is left blank, a random one will be assigned by default
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// VPC ID, which is required for `changeVpc` and `changeBaseToVpc` operations
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID, which is required for `changeVpc` and `changeBaseToVpc` operations
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

func (r *ModifyNetworkConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNetworkConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNetworkConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Execution status: true or false
		Status *bool `json:"Status,omitempty" name:"Status"`

		// Subnet ID
		SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

		// VPC ID
		VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

		// VIP address
		Vip *string `json:"Vip,omitempty" name:"Vip"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNetworkConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNetworkConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Outbound struct {

	// Policy. Valid values: ACCEPT, DROP.
	Action *string `json:"Action,omitempty" name:"Action"`

	// All the addresses that the address group ID represents.
	AddressModule *string `json:"AddressModule,omitempty" name:"AddressModule"`

	// Source IP or IP address range, such as 192.168.0.0/16.
	CidrIp *string `json:"CidrIp,omitempty" name:"CidrIp"`

	// Description.
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// Network protocol, such as UDP and TCP, etc.
	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`

	// Port.
	PortRange *string `json:"PortRange,omitempty" name:"PortRange"`

	// All the protocols and ports that the service group ID represents.
	ServiceModule *string `json:"ServiceModule,omitempty" name:"ServiceModule"`

	// All the addresses that the security group ID represents.
	Id *string `json:"Id,omitempty" name:"Id"`
}

type ProductConf struct {

	// Product type. 2: Redis primary-secondary edition; 3: CKV primary-secondary edition; 4: CKV cluster edition; 5: Redis standalone edition; 7: Redis cluster edition
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// Product name: Redis primary-secondary edition, CKV primary-secondary edition, CKV cluster edition, Redis standalone edition, or Redis cluster edition
	TypeName *string `json:"TypeName,omitempty" name:"TypeName"`

	// Minimum purchasable quantity
	MinBuyNum *int64 `json:"MinBuyNum,omitempty" name:"MinBuyNum"`

	// Maximum purchasable quantity
	MaxBuyNum *int64 `json:"MaxBuyNum,omitempty" name:"MaxBuyNum"`

	// Whether a product is sold out
	Saleout *bool `json:"Saleout,omitempty" name:"Saleout"`

	// Product engine: Tencent Cloud CKV or Redis community edition
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// Compatible version: Redis 2.8, Redis 3.2, or Redis 4.0
	Version *string `json:"Version,omitempty" name:"Version"`

	// Total capacity in GB
	TotalSize []*string `json:"TotalSize,omitempty" name:"TotalSize" list`

	// Shard size in GB
	ShardSize []*string `json:"ShardSize,omitempty" name:"ShardSize" list`

	// Number of replicas
	ReplicaNum []*string `json:"ReplicaNum,omitempty" name:"ReplicaNum" list`

	// Number of shards
	ShardNum []*string `json:"ShardNum,omitempty" name:"ShardNum" list`

	// Supported billing method. 1: monthly subscription; 0: pay-as-you-go
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// Whether to support read-only replicas
	EnableRepicaReadOnly *bool `json:"EnableRepicaReadOnly,omitempty" name:"EnableRepicaReadOnly"`
}

type RedisBackupSet struct {

	// Backup start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Backup ID
	BackupId *string `json:"BackupId,omitempty" name:"BackupId"`

	// Backup type. manualBackupInstance: manual backup initiated by user; systemBackupInstance: midnight backup initiated by system
	BackupType *string `json:"BackupType,omitempty" name:"BackupType"`

	// Backup status. 1: backup is locked by another process; 2: backup is normal and not locked by any process; -1: backup has expired; 3: backup is being exported; 4: backup is exported successfully
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Backup remarks
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Whether a backup is locked. 0: no; 1: yes
	Locked *int64 `json:"Locked,omitempty" name:"Locked"`
}

type RegionConf struct {

	// Region ID
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// Region name
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// Region abbreviation
	RegionShortName *string `json:"RegionShortName,omitempty" name:"RegionShortName"`

	// Name of the area where a region is located
	Area *string `json:"Area,omitempty" name:"Area"`

	// AZ information
	ZoneSet []*ZoneCapacityConf `json:"ZoneSet,omitempty" name:"ZoneSet" list`
}

type RenewInstanceRequest struct {
	*tchttp.BaseRequest

	// Length of purchase in months
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *RenewInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenewInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RenewInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Transaction ID
		DealId *string `json:"DealId,omitempty" name:"DealId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RenewInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenewInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetPasswordRequest struct {
	*tchttp.BaseRequest

	// Redis instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Password reset (this parameter can be left blank when switching to password-free instance mode and is required in other cases)
	Password *string `json:"Password,omitempty" name:"Password"`

	// Whether to switch to password-free instance mode. false: switch to password-enabled instance mode; true: switch to password-free instance mode. Default value: false
	NoAuth *bool `json:"NoAuth,omitempty" name:"NoAuth"`
}

func (r *ResetPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetPasswordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetPasswordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task ID (this parameter is the task ID when changing a password. If you are switching between password-free and password-enabled instance mode, you can leave this parameter alone)
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetPasswordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RestoreInstanceRequest struct {
	*tchttp.BaseRequest

	// ID of the instance to be operated on, which can be obtained through the `redisId` field in the return value of the DescribeRedis API.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Backup ID, which can be obtained through the `backupId` field in the return value of the GetRedisBackupList API
	BackupId *string `json:"BackupId,omitempty" name:"BackupId"`

	// Instance password, which needs to be validated during instance restoration (this parameter is not required for password-free instances)
	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *RestoreInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RestoreInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RestoreInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task ID, which can be used to query the task execution status through the DescribeTaskInfo API
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RestoreInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RestoreInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SecurityGroup struct {

	// Creation time in the format of yyyy-mm-dd hh:mm:ss.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Project ID.
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Security group ID.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// Security group name.
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`

	// Security group remarks.
	SecurityGroupRemark *string `json:"SecurityGroupRemark,omitempty" name:"SecurityGroupRemark"`

	// Outbound rule.
	Outbound []*Outbound `json:"Outbound,omitempty" name:"Outbound" list`

	// Inbound rule.
	Inbound []*Inbound `json:"Inbound,omitempty" name:"Inbound" list`
}

type SecurityGroupDetail struct {

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Security group ID
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// Security group name
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`

	// Security group remarks
	SecurityGroupRemark *string `json:"SecurityGroupRemark,omitempty" name:"SecurityGroupRemark"`

	// Security group inbound rule
	InboundRule []*SecurityGroupsInboundAndOutbound `json:"InboundRule,omitempty" name:"InboundRule" list`

	// Security group outbound rule
	OutboundRule []*SecurityGroupsInboundAndOutbound `json:"OutboundRule,omitempty" name:"OutboundRule" list`
}

type SecurityGroupsInboundAndOutbound struct {

	// Action to be executed
	Action *string `json:"Action,omitempty" name:"Action"`

	// IP address
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Port number
	Port *string `json:"Port,omitempty" name:"Port"`

	// Protocol type
	Proto *string `json:"Proto,omitempty" name:"Proto"`
}

type SourceCommand struct {

	// Command
	Cmd *string `json:"Cmd,omitempty" name:"Cmd"`

	// Number of executions
	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type SourceInfo struct {

	// Source IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Number of connections
	Conn *int64 `json:"Conn,omitempty" name:"Conn"`

	// Command
	Cmd *int64 `json:"Cmd,omitempty" name:"Cmd"`
}

type StartupInstanceRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *StartupInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartupInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartupInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task ID
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartupInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartupInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SwitchInstanceVipRequest struct {
	*tchttp.BaseRequest

	// Source instance ID.
	SrcInstanceId *string `json:"SrcInstanceId,omitempty" name:"SrcInstanceId"`

	// Target instance ID.
	DstInstanceId *string `json:"DstInstanceId,omitempty" name:"DstInstanceId"`

	// The time that lapses in seconds since DTS is disconnected between the source instance and the target instance. If the DTS disconnection time period is greater than `TimeDelay`, the VIP will not be switched. We recommend setting an acceptable value based on the actual business conditions.
	TimeDelay *int64 `json:"TimeDelay,omitempty" name:"TimeDelay"`

	// Whether to force the switch when DTS is disconnected. Valid values: 1 (yes), 0 (no).
	ForceSwitch *int64 `json:"ForceSwitch,omitempty" name:"ForceSwitch"`

	// Valid values: now (switch now), syncComplete (switch after sync is completed).
	SwitchTime *string `json:"SwitchTime,omitempty" name:"SwitchTime"`
}

func (r *SwitchInstanceVipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SwitchInstanceVipRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SwitchInstanceVipResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task ID
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchInstanceVipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SwitchInstanceVipResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TaskInfoDetail struct {

	// Task ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// Start time
	// Note: This field may return null, indicating that no valid values can be obtained.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Task type
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// Instance name
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Instance ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Project ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Task progress
	// Note: This field may return null, indicating that no valid values can be obtained.
	Progress *float64 `json:"Progress,omitempty" name:"Progress"`

	// End time
	// Note: This field may return null, indicating that no valid values can be obtained.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Task status
	// Note: This field may return null, indicating that no valid values can be obtained.
	Result *int64 `json:"Result,omitempty" name:"Result"`
}

type TradeDealDetail struct {

	// Order ID, which is used when a TencentCloud API is called
	DealId *string `json:"DealId,omitempty" name:"DealId"`

	// Long order ID, which is used when an order issue is submitted for assistance
	DealName *string `json:"DealName,omitempty" name:"DealName"`

	// AZ ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// Number of instances associated with an order
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// Creates a user uin
	Creater *string `json:"Creater,omitempty" name:"Creater"`

	// Order creation time
	CreatTime *string `json:"CreatTime,omitempty" name:"CreatTime"`

	// Order timeout period
	OverdueTime *string `json:"OverdueTime,omitempty" name:"OverdueTime"`

	// Order completion time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Order status. 1: unpaid; 2: paid but not delivered; 3: In delivery; 4: successfully delivered; 5: delivery failed; 6: refunded; 7: order closed; 8: order expired; 9: order invalidated; 10: product invalidated; 11: requested payment rejected; 12: paying
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Order status description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Actual total price of an order in 0.01 CNY
	Price *int64 `json:"Price,omitempty" name:"Price"`

	// Instance ID
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

type UpgradeInstanceRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Shard size in MB
	MemSize *uint64 `json:"MemSize,omitempty" name:"MemSize"`

	// Number of shards. This parameter can be left blank for Redis 2.8 primary-secondary edition, CKV primary-secondary edition, and Redis 2.8 standalone edition
	RedisShardNum *uint64 `json:"RedisShardNum,omitempty" name:"RedisShardNum"`

	// Number of replicas. This parameter can be left blank for Redis 2.8 primary-secondary edition, CKV primary-secondary edition, and Redis 2.8 standalone edition
	RedisReplicasNum *uint64 `json:"RedisReplicasNum,omitempty" name:"RedisReplicasNum"`
}

func (r *UpgradeInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Order ID
		DealId *string `json:"DealId,omitempty" name:"DealId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpgradeInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ZoneCapacityConf struct {

	// AZ ID, such as ap-guangzhou-3
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// AZ name
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// Whether a product is sold out in an AZ
	IsSaleout *bool `json:"IsSaleout,omitempty" name:"IsSaleout"`

	// Whether it is a default AZ
	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`

	// Network type. basenet: basic network; vpcnet: VPC
	NetWorkType []*string `json:"NetWorkType,omitempty" name:"NetWorkType" list`

	// Information of an AZ, such as product specifications in it
	ProductSet []*ProductConf `json:"ProductSet,omitempty" name:"ProductSet" list`

	// AZ ID, such as 100003
	OldZoneId *int64 `json:"OldZoneId,omitempty" name:"OldZoneId"`
}
