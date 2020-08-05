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

package v20190725

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type AssignProjectRequest struct {
	*tchttp.BaseRequest

	// List of instance IDs in the format of cmgo-p8vnipr5. It is the same as the instance ID displayed on the TencentDB Console page
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *AssignProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssignProjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssignProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of the returned async task IDs
		FlowIds []*uint64 `json:"FlowIds,omitempty" name:"FlowIds" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssignProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssignProjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BackupFile struct {

	// ID of the replica set/shard to which a backup file belongs
	ReplicateSetId *string `json:"ReplicateSetId,omitempty" name:"ReplicateSetId"`

	// Path to a backup file
	File *string `json:"File,omitempty" name:"File"`
}

type BackupInfo struct {

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Backup mode. 0: automatic backup; 1: manual backup
	BackupType *uint64 `json:"BackupType,omitempty" name:"BackupType"`

	// Backup name
	BackupName *string `json:"BackupName,omitempty" name:"BackupName"`

	// Backup remarks
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupDesc *string `json:"BackupDesc,omitempty" name:"BackupDesc"`

	// Backup file size in KB
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupSize *uint64 `json:"BackupSize,omitempty" name:"BackupSize"`

	// Backup start time
	// Note: This field may return null, indicating that no valid values can be obtained.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Backup end time
	// Note: This field may return null, indicating that no valid values can be obtained.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Backup status. 1: backing up; 2: backed up successful
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// Backup method. 0: logical backup; 1: physical backup
	BackupMethod *uint64 `json:"BackupMethod,omitempty" name:"BackupMethod"`
}

type ClientConnection struct {

	// Client IP of a connection
	IP *string `json:"IP,omitempty" name:"IP"`

	// Number of connections corresponding to a client IP
	Count *uint64 `json:"Count,omitempty" name:"Count"`
}

type CreateDBInstanceHourRequest struct {
	*tchttp.BaseRequest

	// Instance memory size in GB
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// Instance disk size in GB
	Volume *uint64 `json:"Volume,omitempty" name:"Volume"`

	// Number of replica sets. When a replica set instance is created, this parameter must be set to 1. When a sharding instance is created, please see the parameters returned by the DescribeSpecInfo API
	ReplicateSetNum *uint64 `json:"ReplicateSetNum,omitempty" name:"ReplicateSetNum"`

	// Number of nodes in each replica set. Currently, the number of nodes in a replica set is fixed at 3, while the number of shards is customizable. For more information, please see the parameter returned by the DescribeSpecInfo API
	NodeNum *uint64 `json:"NodeNum,omitempty" name:"NodeNum"`

	// Version number. For the specific purchasable versions supported, please see the return result of the DescribeSpecInfo API. The correspondences between parameters and versions are as follows: MONGO_3_WT: MongoDB 3.2 WiredTiger Edition; MONGO_3_ROCKS: MongoDB 3.2 RocksDB Edition; MONGO_36_WT: MongoDB 3.6 WiredTiger Edition
	MongoVersion *string `json:"MongoVersion,omitempty" name:"MongoVersion"`

	// Server type. HIO: high IO; HIO10G: 10-Gigabit high IO
	MachineCode *string `json:"MachineCode,omitempty" name:"MachineCode"`

	// Number of instances. Minimum value: 1. Maximum value: 10
	GoodsNum *uint64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// AZ information in the format of ap-guangzhou-2
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Instance type. REPLSET: replica set; SHARD: sharding cluster
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// VPC ID. If this parameter is not set, the basic network will be selected by default
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// VPC subnet ID. If VpcId is set, then SubnetId will be required
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Instance password. If this parameter is not set, you need to set an instance password through the password setting API after creating an instance. The password can only contain 8-16 characters and must contain at least two of the following types of characters: letters, digits, and special characters `!@#%^*()` |
	Password *string `json:"Password,omitempty" name:"Password"`

	// Project ID. If this parameter is not set, the default project will be used
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Instance tag information
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags" list`

	// 
	Clone *int64 `json:"Clone,omitempty" name:"Clone"`

	// 
	Father *string `json:"Father,omitempty" name:"Father"`

	// 
	SecurityGroup []*string `json:"SecurityGroup,omitempty" name:"SecurityGroup" list`
}

func (r *CreateDBInstanceHourRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDBInstanceHourRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDBInstanceHourResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Order ID
		DealId *string `json:"DealId,omitempty" name:"DealId"`

		// List of IDs of created instances
		InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDBInstanceHourResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDBInstanceHourResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDBInstanceRequest struct {
	*tchttp.BaseRequest

	// Number of nodes in each replica set. Currently, the number of nodes per replica set is fixed at 3, while the number of secondary nodes per shard is customizable. For more information, please see the parameter returned by the `DescribeSpecInfo` API.
	NodeNum *uint64 `json:"NodeNum,omitempty" name:"NodeNum"`

	// Instance memory size in GB.
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// Instance disk size in GB.
	Volume *uint64 `json:"Volume,omitempty" name:"Volume"`

	// Version number. For the specific purchasable versions supported, please see the return result of the `DescribeSpecInfo` API. The correspondences between parameters and versions are as follows: MONGO_3_WT: MongoDB 3.2 WiredTiger Edition; MONGO_3_ROCKS: MongoDB 3.2 RocksDB Edition; MONGO_36_WT: MongoDB 3.6 WiredTiger Edition.
	MongoVersion *string `json:"MongoVersion,omitempty" name:"MongoVersion"`

	// Number of instances. Minimum value: 1. Maximum value: 10.
	GoodsNum *uint64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// Instance region name in the format of ap-guangzhou-2.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Instance validity period in months. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36.
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// Server type. Valid values: HIO (high IO), HIO10G (10-gigabit high IO).
	MachineCode *string `json:"MachineCode,omitempty" name:"MachineCode"`

	// Instance type. Valid values: REPLSET (replica set), SHARD (sharded cluster).
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// Number of replica sets. To create a replica set instance, set this parameter to 1; to create a shard instance, see the parameters returned by the `DescribeSpecInfo` API.
	ReplicateSetNum *uint64 `json:"ReplicateSetNum,omitempty" name:"ReplicateSetNum"`

	// Project ID. If this parameter is not set, the default project will be used.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// VPC ID. If this parameter is not set, the classic network will be used. Please use the `DescribeVpcs` API to query the VPC list.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// VPC subnet ID. If `UniqVpcId` is set, then `UniqSubnetId` will be required. Please use the `DescribeSubnets` API to query the subnet list.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Instance password. If this parameter is not set, you need to set an instance password through the `SetPassword` API after creating an instance. The password can only contain 8-16 characters and must contain at least two of the following types of characters: letters, digits, and special characters `!@#%^*()`.
	Password *string `json:"Password,omitempty" name:"Password"`

	// Instance tag information.
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags" list`

	// Auto-renewal flag. Valid values: 0 (auto-renewal not enabled), 1 (auto-renewal enabled). Default value: 0.
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 
	AutoVoucher *uint64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 
	Clone *int64 `json:"Clone,omitempty" name:"Clone"`

	// 
	Father *string `json:"Father,omitempty" name:"Father"`

	// 
	SecurityGroup []*string `json:"SecurityGroup,omitempty" name:"SecurityGroup" list`
}

func (r *CreateDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Order ID.
		DealId *string `json:"DealId,omitempty" name:"DealId"`

		// List of IDs of created instances.
		InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DBInstanceInfo struct {

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Region information
	Region *string `json:"Region,omitempty" name:"Region"`
}

type DBInstancePrice struct {

	// Unit price.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UnitPrice *float64 `json:"UnitPrice,omitempty" name:"UnitPrice"`

	// Original price.
	OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

	// Discounted price.
	DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`
}

type DescribeBackupAccessRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cmgo-p8vnipr5. It is the same as the instance ID displayed on the TencentDB Console page
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Name of the backup file for which to get the download permission
	BackupName *string `json:"BackupName,omitempty" name:"BackupName"`
}

func (r *DescribeBackupAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBackupAccessRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupAccessResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Instance region
		Region *string `json:"Region,omitempty" name:"Region"`

		// The bucket where a backup file is located
		Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

		// Storage information of a backup file
		Files []*BackupFile `json:"Files,omitempty" name:"Files" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBackupAccessResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClientConnectionsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cmgo-p8vnipr5. It is the same as the instance ID displayed on the TencentDB Console page
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeClientConnectionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClientConnectionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClientConnectionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Client connection information, including client IP and number of connections
		Clients []*ClientConnection `json:"Clients,omitempty" name:"Clients" list`

		// 
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClientConnectionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClientConnectionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBBackupsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cmgo-p8vnipr5. It is the same as the instance ID displayed on the TencentDB Console page
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBBackupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBBackupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBBackupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Backup list
		BackupList []*BackupInfo `json:"BackupList,omitempty" name:"BackupList" list`

		// Total number of backups
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBBackupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBBackupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceDealRequest struct {
	*tchttp.BaseRequest

	// Order ID. It is returned by the `CreateDBInstance` and other APIs.
	DealId *string `json:"DealId,omitempty" name:"DealId"`
}

func (r *DescribeDBInstanceDealRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstanceDealRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceDealResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Order status. Valid values: 1 (unpaid), 2 (paid), 3 (delivering), 4 (delivered), 5 (delivery failed), 6 (refunded), 7 (order closed), 8 (order closed because it failed to be paid within timeout period).
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// Original price of the order.
		OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

		// Discounted price of the order.
		DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`

		// Operation performed by the order. Valid values: purchase, renew, upgrade, downgrade, refund.
		Action *string `json:"Action,omitempty" name:"Action"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBInstanceDealResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstanceDealResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstancesRequest struct {
	*tchttp.BaseRequest

	// List of instance IDs in the format of cmgo-p8vnipr5. It is the same as the instance ID displayed on the TencentDB Console page
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// Instance type. Valid values: 0 (all instances), 1 (promoted), 2 (temp), 3 (read-only), -1 (promoted + read-only + disaster recovery)
	InstanceType *int64 `json:"InstanceType,omitempty" name:"InstanceType"`

	// Cluster type. Valid values: 0 (replica set instance), 1 (sharding instance), -1 (all instances)
	ClusterType *int64 `json:"ClusterType,omitempty" name:"ClusterType"`

	// Instance status. Valid values: 0 (to be initialized), 1 (in process), 2 (valid), -2 (expired)
	Status []*int64 `json:"Status,omitempty" name:"Status" list`

	// VPC ID. This parameter can be left empty for the basic network
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID of VPC. This parameter can be left empty for the basic network. If it is passed in as an input parameter, the corresponding VpcId must be set
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Billing type. Valid value: 0 (pay-as-you-go)
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// Number of results to be returned for a single request. Valid values: 1-100. Default value: 20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Sort by field of the returned result set. Currently, supported values include "ProjectId", "InstanceName", and "CreateTime". The return results are sorted in ascending order by default.
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Sorting method of the return result set. Currently, "ASC" or "DESC" is supported
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// Project ID
	ProjectIds []*uint64 `json:"ProjectIds,omitempty" name:"ProjectIds" list`

	// Search keyword, which can be instance ID, instance name, or complete IP
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
}

func (r *DescribeDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible instances.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of instance details
		InstanceDetails []*InstanceDetail `json:"InstanceDetails,omitempty" name:"InstanceDetails" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogPatternsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of `cmgo-p8vnipr5`, which is the same as the instance ID displayed on the TencentDB Console page
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Start time of slow log in the format of `yyyy-mm-dd hh:mm:ss`, such as 2019-06-01 10:00:00. The query time range cannot exceed 24 hours. Only slow logs for the last 7 days can be queried.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of slow log in the format of `yyyy-mm-dd hh:mm:ss`, such as 2019-06-02 12:00:00. The query time range cannot exceed 24 hours. Only slow logs for the last 7 days can be queried.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Threshold of slow log execution time in milliseconds. Minimum value: 100. Slow logs whose execution time exceeds the threshold will be returned.
	SlowMS *uint64 `json:"SlowMS,omitempty" name:"SlowMS"`

	// Offset. Minimum value: 0. Maximum value: 10000. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page. Minimum value: 1. Maximum value: 100. Default value: 20.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSlowLogPatternsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSlowLogPatternsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogPatternsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of slow logs
		Count *uint64 `json:"Count,omitempty" name:"Count"`

		// Slow log statistics
		SlowLogPatterns []*SlowLogPattern `json:"SlowLogPatterns,omitempty" name:"SlowLogPatterns" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSlowLogPatternsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSlowLogPatternsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of `cmgo-p8vnipr5`, which is the same as the instance ID displayed on the TencentDB Console page
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Start time of slow log in the format of `yyyy-mm-dd hh:mm:ss`, such as 2019-06-01 10:00:00. The query time range cannot exceed 24 hours. Only slow logs for the last 7 days can be queried.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of slow log in the format of `yyyy-mm-dd hh:mm:ss`, such as 2019-06-02 12:00:00. The query time range cannot exceed 24 hours. Only slow logs for the last 7 days can be queried.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Threshold of slow log execution time in milliseconds. Minimum value: 100. Slow logs whose execution time exceeds the threshold will be returned.
	SlowMS *uint64 `json:"SlowMS,omitempty" name:"SlowMS"`

	// Offset. Minimum value: 0. Maximum value: 10000. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page. Minimum value: 1. Maximum value: 100. Default value: 20.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSlowLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSlowLogsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of slow logs
		Count *uint64 `json:"Count,omitempty" name:"Count"`

		// Slow log details
		SlowLogs []*string `json:"SlowLogs,omitempty" name:"SlowLogs" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSlowLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSlowLogsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSpecInfoRequest struct {
	*tchttp.BaseRequest

	// AZ to be queried
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

func (r *DescribeSpecInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSpecInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSpecInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of purchasable instance specifications
		SpecInfoList []*SpecificationInfo `json:"SpecInfoList,omitempty" name:"SpecInfoList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSpecInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSpecInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type FlushInstanceRouterConfigRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *FlushInstanceRouterConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *FlushInstanceRouterConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type FlushInstanceRouterConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *FlushInstanceRouterConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *FlushInstanceRouterConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquirePriceCreateDBInstancesRequest struct {
	*tchttp.BaseRequest

	// Instance region name in the format of ap-guangzhou-2.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Number of nodes in each replica set. Currently, the number of nodes per replica set is fixed at 3, while the number of secondary nodes per shard is customizable. For more information, please see the parameter returned by the `DescribeSpecInfo` API.
	NodeNum *int64 `json:"NodeNum,omitempty" name:"NodeNum"`

	// Instance memory size in GB.
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Instance disk size in GB.
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`

	// Version number. For the specific purchasable versions supported, please see the return result of the `DescribeSpecInfo` API. The correspondences between parameters and versions are as follows: MONGO_3_WT: MongoDB 3.2 WiredTiger Edition; MONGO_3_ROCKS: MongoDB 3.2 RocksDB Edition; MONGO_36_WT: MongoDB 3.6 WiredTiger Edition; MONGO_40_WT: MongoDB 4.0 WiredTiger Edition.
	MongoVersion *string `json:"MongoVersion,omitempty" name:"MongoVersion"`

	// Server type. Valid values: HIO (high IO), HIO10G (10-gigabit high IO), STDS5 (standard).
	MachineCode *string `json:"MachineCode,omitempty" name:"MachineCode"`

	// Number of instances. Minimum value: 1. Maximum value: 10.
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// Instance validity period in months. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36.
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// Instance type. Valid values: REPLSET (replica set), SHARD (sharded cluster), STANDALONE (single-node).
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// Number of replica sets. To create a replica set instance, set this parameter to 1; to create a shard instance, see the parameters returned by the `DescribeSpecInfo` API; to create a single-node instance, set this parameter to 0.
	ReplicateSetNum *int64 `json:"ReplicateSetNum,omitempty" name:"ReplicateSetNum"`
}

func (r *InquirePriceCreateDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquirePriceCreateDBInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquirePriceCreateDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Price.
		Price *DBInstancePrice `json:"Price,omitempty" name:"Price"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquirePriceCreateDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquirePriceCreateDBInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquirePriceModifyDBInstanceSpecRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cmgo-p8vnipr5. It is the same as the instance ID displayed in the TencentDB Console.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance memory size in GB after specification adjustment.
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Instance disk size in GB after specification adjustment.
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`
}

func (r *InquirePriceModifyDBInstanceSpecRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquirePriceModifyDBInstanceSpecRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquirePriceModifyDBInstanceSpecResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Price.
		Price *DBInstancePrice `json:"Price,omitempty" name:"Price"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquirePriceModifyDBInstanceSpecResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquirePriceModifyDBInstanceSpecResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquirePriceRenewDBInstancesRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cmgo-p8vnipr5. It is the same as the instance ID displayed in the TencentDB Console. This API supports operations on up to 5 instances at a time.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// The parameter setting for the prepaid mode (monthly subscription mode). This parameter can specify the renewal period, whether to set automatic renewal, and other attributes of the monthly subscription instance.
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
}

func (r *InquirePriceRenewDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquirePriceRenewDBInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquirePriceRenewDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Price.
		Price *DBInstancePrice `json:"Price,omitempty" name:"Price"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquirePriceRenewDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquirePriceRenewDBInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InstanceChargePrepaid struct {

	// Purchased usage period (in month). Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36. Default value: 1.
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// Auto-renewal flag. Valid values:
	// NOTIFY_AND_AUTO_RENEW: notify expiration and renew automatically
	// NOTIFY_AND_MANUAL_RENEW: notify expiration but not renew automatically
	// DISABLE_NOTIFY_AND_MANUAL_RENEW: neither notify expiration nor renew automatically
	// 
	// Default value: NOTIFY_AND_MANUAL_RENEW. If this parameter is specified as NOTIFY_AND_AUTO_RENEW, the instance will be automatically renewed on a monthly basis when the account balance is sufficient.
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type InstanceDetail struct {

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Billing type. Valid value: 0 (pay-as-you-go)
	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Cluster type. Valid values: 0 (replica set instance), 1 (sharding instance),
	ClusterType *uint64 `json:"ClusterType,omitempty" name:"ClusterType"`

	// Region information
	Region *string `json:"Region,omitempty" name:"Region"`

	// AZ information
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Network type. Valid values: 0 (basic network), 1 (VPC)
	NetType *uint64 `json:"NetType,omitempty" name:"NetType"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID of VPC
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Instance status. Valid values: 0 (to be initialized), 1 (in process), 2 (running), -2 (expired)
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Instance IP
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Port number
	Vport *uint64 `json:"Vport,omitempty" name:"Vport"`

	// Instance creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Instance expiration time
	DeadLine *string `json:"DeadLine,omitempty" name:"DeadLine"`

	// Instance version information
	MongoVersion *string `json:"MongoVersion,omitempty" name:"MongoVersion"`

	// Instance memory size in MB
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// Instance disk size in MB
	Volume *uint64 `json:"Volume,omitempty" name:"Volume"`

	// Number of CPU cores of an instance
	CpuNum *uint64 `json:"CpuNum,omitempty" name:"CpuNum"`

	// Instance machine type
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// Number of secondary nodes of an instance
	SecondaryNum *uint64 `json:"SecondaryNum,omitempty" name:"SecondaryNum"`

	// Number of instance shards
	ReplicationSetNum *uint64 `json:"ReplicationSetNum,omitempty" name:"ReplicationSetNum"`

	// Instance auto-renewal flag. Valid values: 0 (manual renewal), 1 (auto-renewal), 2 (no renewal upon expiration)
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// Used capacity in MB
	UsedVolume *uint64 `json:"UsedVolume,omitempty" name:"UsedVolume"`

	// Start time of the maintenance time window
	MaintenanceStart *string `json:"MaintenanceStart,omitempty" name:"MaintenanceStart"`

	// End time of the maintenance time window
	MaintenanceEnd *string `json:"MaintenanceEnd,omitempty" name:"MaintenanceEnd"`

	// Shard information
	ReplicaSets []*ShardInfo `json:"ReplicaSets,omitempty" name:"ReplicaSets" list`

	// Information of read-only instances
	ReadonlyInstances []*DBInstanceInfo `json:"ReadonlyInstances,omitempty" name:"ReadonlyInstances" list`

	// Information of disaster recovery instances
	StandbyInstances []*DBInstanceInfo `json:"StandbyInstances,omitempty" name:"StandbyInstances" list`

	// Information of temp instances
	CloneInstances []*DBInstanceInfo `json:"CloneInstances,omitempty" name:"CloneInstances" list`

	// Information of associated instances. For a promoted instance, this field represents information of its temp instance; for a temp instance, this field represents information of its promoted instance; and for a read-only/disaster recovery instance, this field represents information of its primary instance
	RelatedInstance *DBInstanceInfo `json:"RelatedInstance,omitempty" name:"RelatedInstance"`

	// Instance tag information set
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags" list`

	// Instance version tag
	InstanceVer *uint64 `json:"InstanceVer,omitempty" name:"InstanceVer"`

	// Instance version tag
	ClusterVer *uint64 `json:"ClusterVer,omitempty" name:"ClusterVer"`

	// Protocol information. Valid values: 1 (mongodb), 2 (dynamodb)
	Protocol *uint64 `json:"Protocol,omitempty" name:"Protocol"`

	// Instance type. Valid values: 1 (promoted instance), 2 (temp instance), 3 (read-only instance), 4 (disaster recovery instance)
	InstanceType *uint64 `json:"InstanceType,omitempty" name:"InstanceType"`

	// Instance status description
	InstanceStatusDesc *string `json:"InstanceStatusDesc,omitempty" name:"InstanceStatusDesc"`

	// Physical instance ID. For an instance that has been rolled back and replaced, its InstanceId and RealInstanceId are different. The physical instance ID is needed in such scenarios as getting monitoring data from Barad
	RealInstanceId *string `json:"RealInstanceId,omitempty" name:"RealInstanceId"`
}

type IsolateDBInstanceRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cmgo-p8vnipr5. It is the same as the instance ID displayed on the TencentDB Console page
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *IsolateDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *IsolateDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type IsolateDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task request ID, which can be used to query the execution result of an async task.
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *IsolateDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *IsolateDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceSpecRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cmgo-p8vnipr5. It is the same as the instance ID displayed on the TencentDB Console page
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Memory size after instance configuration change in GB. Memory and disk must be upgraded or degraded simultaneously
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// Disk size after instance configuration change in GB. Memory and disk must be upgraded or degraded simultaneously. For degradation, the new disk capacity must be greater than 1.2 times the used disk capacity
	Volume *uint64 `json:"Volume,omitempty" name:"Volume"`

	// Oplog size after instance configuration change in GB, which ranges from 10% to 90% of the disk capacity and is 10% of the disk capacity by default
	OplogSize *uint64 `json:"OplogSize,omitempty" name:"OplogSize"`
}

func (r *ModifyDBInstanceSpecRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBInstanceSpecRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceSpecResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Order ID
		DealId *string `json:"DealId,omitempty" name:"DealId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBInstanceSpecResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBInstanceSpecResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OfflineIsolatedDBInstanceRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cmgo-p8vnipr5. It is the same as the instance ID displayed on the TencentDB Console page
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *OfflineIsolatedDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OfflineIsolatedDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OfflineIsolatedDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task request ID, which can be used to query the execution result of an async task.
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OfflineIsolatedDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OfflineIsolatedDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RenameInstanceRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cmgo-p8vnipr5. It is the same as the instance ID displayed on the TencentDB Console page
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance name
	NewName *string `json:"NewName,omitempty" name:"NewName"`
}

func (r *RenameInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenameInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RenameInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RenameInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenameInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RenewDBInstancesRequest struct {
	*tchttp.BaseRequest

	// IDs of one or more instances to be operated. The value can be obtained from the `InstanceId` parameter returned by the `DescribeInstances` API. Up to 100 instances can be requested at a time.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// The parameter setting for the prepaid mode (monthly subscription mode). This parameter can specify the renewal period, whether to set automatic renewal, and other attributes of the monthly subscription instance. This parameter is mandatory in monthly subscription.
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
}

func (r *RenewDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenewDBInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RenewDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RenewDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenewDBInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ShardInfo struct {

	// Used shard capacity
	UsedVolume *float64 `json:"UsedVolume,omitempty" name:"UsedVolume"`

	// Shard ID
	ReplicaSetId *string `json:"ReplicaSetId,omitempty" name:"ReplicaSetId"`

	// Shard name
	ReplicaSetName *string `json:"ReplicaSetName,omitempty" name:"ReplicaSetName"`

	// Shard memory size in MB
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// Shard disk size in MB
	Volume *uint64 `json:"Volume,omitempty" name:"Volume"`

	// Shard oplog size in MB
	OplogSize *uint64 `json:"OplogSize,omitempty" name:"OplogSize"`

	// Number of secondary nodes of a shard
	SecondaryNum *uint64 `json:"SecondaryNum,omitempty" name:"SecondaryNum"`

	// Shard physical ID
	RealReplicaSetId *string `json:"RealReplicaSetId,omitempty" name:"RealReplicaSetId"`
}

type SlowLogPattern struct {

	// Slow log pattern
	Pattern *string `json:"Pattern,omitempty" name:"Pattern"`

	// Maximum execution time
	MaxTime *uint64 `json:"MaxTime,omitempty" name:"MaxTime"`

	// Average execution time
	AverageTime *uint64 `json:"AverageTime,omitempty" name:"AverageTime"`

	// Number of slow logs in this pattern
	Total *uint64 `json:"Total,omitempty" name:"Total"`
}

type SpecItem struct {

	// Specification information identifier
	SpecCode *string `json:"SpecCode,omitempty" name:"SpecCode"`

	// Specification purchasable flag. Valid values: 0 (not purchasable), 1 (purchasable)
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// Specification purchasable flag. Valid values: 0 (not purchasable), 1 (purchasable)
	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`

	// Memory size in MB
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// Default disk size in MB
	DefaultStorage *uint64 `json:"DefaultStorage,omitempty" name:"DefaultStorage"`

	// Maximum disk size in MB
	MaxStorage *uint64 `json:"MaxStorage,omitempty" name:"MaxStorage"`

	// Minimum disk size in MB
	MinStorage *uint64 `json:"MinStorage,omitempty" name:"MinStorage"`

	// Maximum QPS
	Qps *uint64 `json:"Qps,omitempty" name:"Qps"`

	// Maximum number of connections
	Conns *uint64 `json:"Conns,omitempty" name:"Conns"`

	// MongoDB version information of an instance
	MongoVersionCode *string `json:"MongoVersionCode,omitempty" name:"MongoVersionCode"`

	// MongoDB version number of an instance
	MongoVersionValue *uint64 `json:"MongoVersionValue,omitempty" name:"MongoVersionValue"`

	// MongoDB version number of an instance (short)
	Version *string `json:"Version,omitempty" name:"Version"`

	// Storage engine
	EngineName *string `json:"EngineName,omitempty" name:"EngineName"`

	// Cluster type. Valid values: 1 (sharding cluster), 0 (replica set cluster)
	ClusterType *uint64 `json:"ClusterType,omitempty" name:"ClusterType"`

	// Minimum number of secondary nodes in a replica set
	MinNodeNum *uint64 `json:"MinNodeNum,omitempty" name:"MinNodeNum"`

	// Maximum number of secondary nodes in a replica set
	MaxNodeNum *uint64 `json:"MaxNodeNum,omitempty" name:"MaxNodeNum"`

	// Minimum number of shards
	MinReplicateSetNum *uint64 `json:"MinReplicateSetNum,omitempty" name:"MinReplicateSetNum"`

	// Maximum number of shards
	MaxReplicateSetNum *uint64 `json:"MaxReplicateSetNum,omitempty" name:"MaxReplicateSetNum"`

	// Minimum number of secondary nodes in a shard
	MinReplicateSetNodeNum *uint64 `json:"MinReplicateSetNodeNum,omitempty" name:"MinReplicateSetNodeNum"`

	// Maximum number of secondary nodes in a shard
	MaxReplicateSetNodeNum *uint64 `json:"MaxReplicateSetNodeNum,omitempty" name:"MaxReplicateSetNodeNum"`

	// Server type. Valid values: 0 (HIO), 4 (HIO10G)
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
}

type SpecificationInfo struct {

	// Region information
	Region *string `json:"Region,omitempty" name:"Region"`

	// AZ information
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Purchasable specification information
	SpecItems []*SpecItem `json:"SpecItems,omitempty" name:"SpecItems" list`
}

type TagInfo struct {

	// Tag key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}
