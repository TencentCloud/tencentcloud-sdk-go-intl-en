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

package v20190107

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type Account struct {

	// Database account name
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// Database account description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Update time
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Host
	Host *string `json:"Host,omitempty" name:"Host"`
}

type AddInstancesRequest struct {
	*tchttp.BaseRequest

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Number of CPU cores
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// Memory
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Number of added read-only instances
	ReadOnlyCount *int64 `json:"ReadOnlyCount,omitempty" name:"ReadOnlyCount"`

	// Instance group ID, which is used when you add an instance in an existing RO group. If this parameter is left empty, an RO group will be created.
	InstanceGrpId *string `json:"InstanceGrpId,omitempty" name:"InstanceGrpId"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Port used when adding RO group
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Whether to automatically select a voucher. 1: yes; 0: no. Default value: 0
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// Database type. Valid values: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// Order source
	OrderSource *string `json:"OrderSource,omitempty" name:"OrderSource"`

	// Transaction mode. Valid values: `0` (place and pay for an order), `1` (place an order)
	DealMode *int64 `json:"DealMode,omitempty" name:"DealMode"`
}

func (r *AddInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Cpu")
	delete(f, "Memory")
	delete(f, "ReadOnlyCount")
	delete(f, "InstanceGrpId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "Port")
	delete(f, "InstanceName")
	delete(f, "AutoVoucher")
	delete(f, "DbType")
	delete(f, "OrderSource")
	delete(f, "DealMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AddInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Freezing transaction. One freezing transaction ID is generated each time an instance is added.
	// Note: this field may return null, indicating that no valid values can be obtained.
		TranId *string `json:"TranId,omitempty" name:"TranId"`

		// Pay-as-You-Go order ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
		DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`

		// List of IDs of delivered resources
	// Note: this field may return null, indicating that no valid values can be obtained.
		ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

		// Big order ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
		BigDealIds []*string `json:"BigDealIds,omitempty" name:"BigDealIds"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Addr struct {

	// IP
	IP *string `json:"IP,omitempty" name:"IP"`

	// Port
	Port *int64 `json:"Port,omitempty" name:"Port"`
}

type BackupFileInfo struct {

	// Snapshot file ID used for rollback
	SnapshotId *uint64 `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// Snapshot file name
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// Snapshot file size
	FileSize *uint64 `json:"FileSize,omitempty" name:"FileSize"`

	// Snapshot backup start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Snapshot backup end time
	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`

	// Backup type. snapshot: snapshot backup; timepoint: time point backup
	BackupType *string `json:"BackupType,omitempty" name:"BackupType"`

	// Back mode. auto: auto backup; manual: manual backup
	BackupMethod *string `json:"BackupMethod,omitempty" name:"BackupMethod"`

	// Backup file status. success: backup succeeded; fail: backup failed; creating: creating backup file; deleting: deleting backup file
	BackupStatus *string `json:"BackupStatus,omitempty" name:"BackupStatus"`

	// Backup file time
	SnapshotTime *string `json:"SnapshotTime,omitempty" name:"SnapshotTime"`
}

type BillingResourceInfo struct {

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Instance ID list
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type ClusterInstanceDetail struct {

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Engine type
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Instance status
	InstanceStatus *string `json:"InstanceStatus,omitempty" name:"InstanceStatus"`

	// Instance status description
	InstanceStatusDesc *string `json:"InstanceStatusDesc,omitempty" name:"InstanceStatusDesc"`

	// Number of CPU cores
	InstanceCpu *int64 `json:"InstanceCpu,omitempty" name:"InstanceCpu"`

	// Memory
	InstanceMemory *int64 `json:"InstanceMemory,omitempty" name:"InstanceMemory"`

	// Disk
	InstanceStorage *int64 `json:"InstanceStorage,omitempty" name:"InstanceStorage"`
}

type CreateClustersRequest struct {
	*tchttp.BaseRequest

	// AZ
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Database type. Valid values: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// Database version. Valid values: 
	// <li> Valid values for `MYSQL`: 5.7 </li>
	DbVersion *string `json:"DbVersion,omitempty" name:"DbVersion"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Number of CPU cores of normal instance
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// Memory of a non-serverless instance in GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Storage capacity in GB
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// Cluster name
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// Account password (it must contain 8-64 characters in at least three of the following four types: uppercase letters, lowercase letters, digits, and symbols (~!@#$%^&*_-+=`|\(){}[]:;'<>,.?/).)
	AdminPassword *string `json:"AdminPassword,omitempty" name:"AdminPassword"`

	// Port. Default value: 5432
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Billing mode. 0: pay-as-you-go; 1: monthly subscription. Default value: 0
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// Number of purchased items. Currently, only 1 can be passed in. If this parameter is left empty, 1 will be used by default.
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// Rollback type:
	// noneRollback: no rollback
	// snapRollback: rollback by snapshot
	// timeRollback: rollback by time point
	RollbackStrategy *string `json:"RollbackStrategy,omitempty" name:"RollbackStrategy"`

	// `snapshotId` for snapshot rollback or `queryId` for time point rollback. 0 indicates to determine whether the time point is valid
	RollbackId *uint64 `json:"RollbackId,omitempty" name:"RollbackId"`

	// Pass in the source cluster ID during rollback to find the source `poolId`
	OriginalClusterId *string `json:"OriginalClusterId,omitempty" name:"OriginalClusterId"`

	// Specified time for time point rollback or snapshot time for snapshot rollback
	ExpectTime *string `json:"ExpectTime,omitempty" name:"ExpectTime"`

	// Specified allowed time range for time point rollback
	ExpectTimeThresh *uint64 `json:"ExpectTimeThresh,omitempty" name:"ExpectTimeThresh"`

	// The maximum storage of a non-serverless instance in GB
	// If `DbType` is `MYSQL` and the storage billing mode is prepaid, the parameter value cannot exceed the maximum storage corresponding to the CPU and memory specifications.
	StorageLimit *int64 `json:"StorageLimit,omitempty" name:"StorageLimit"`

	// Number of instances
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// Purchase duration of monthly subscription plan
	TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// Purchase duration unit of monthly subscription plan
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// Whether auto-renewal is enabled for monthly subscription plan
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// Whether to automatically select a voucher. 1: yes; 0: no. Default value: 0
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// Number of instances (this parameter has been disused and is retained only for compatibility with existing instances)
	HaCount *int64 `json:"HaCount,omitempty" name:"HaCount"`

	// Order source
	OrderSource *string `json:"OrderSource,omitempty" name:"OrderSource"`

	// Array of tags to be bound to the created cluster
	ResourceTags []*Tag `json:"ResourceTags,omitempty" name:"ResourceTags"`

	// Database type
	// Valid values when `DbType` is `MYSQL` (default value: NORMAL):
	// <li>NORMAL</li>
	// <li>SERVERLESS</li>
	DbMode *string `json:"DbMode,omitempty" name:"DbMode"`

	// This parameter is required if `DbMode` is `SERVERLESS`
	// Minimum number of CPU cores. For the value range, please see the returned result of `DescribeServerlessInstanceSpecs`
	MinCpu *float64 `json:"MinCpu,omitempty" name:"MinCpu"`

	// This parameter is required if `DbMode` is `SERVERLESS`:
	// Maximum number of CPU cores. For the value range, please see the returned result of `DescribeServerlessInstanceSpecs`
	MaxCpu *float64 `json:"MaxCpu,omitempty" name:"MaxCpu"`

	// This parameter specifies whether the cluster will be automatically paused if `DbMode` is `SERVERLESS`. Valid values:
	// <li>yes</li>
	// <li>no</li>
	// Default value: yes
	AutoPause *string `json:"AutoPause,omitempty" name:"AutoPause"`

	// This parameter specifies the delay for automatic cluster pause in seconds if `DbMode` is `SERVERLESS`. Value range: [600,691200]
	// Default value: 600
	AutoPauseDelay *int64 `json:"AutoPauseDelay,omitempty" name:"AutoPauseDelay"`

	// The billing mode of cluster storage. Valid values: `0` (postpaid), `1` (prepaid). Default value: `0`.
	// If `DbType` is `MYSQL` and the billing mode of cluster compute is pay-as-you-go (or the `DbMode` is `SERVERLESS`), the billing mode of cluster storage must be postpaid.
	// Clusters with storage billed in prepaid mode cannot be cloned or rolled back.
	StoragePayMode *int64 `json:"StoragePayMode,omitempty" name:"StoragePayMode"`

	// Array of security group IDs
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// Array of alarm policy IDs
	AlarmPolicyIds []*string `json:"AlarmPolicyIds,omitempty" name:"AlarmPolicyIds"`

	// Array of parameters
	ClusterParams []*ParamItem `json:"ClusterParams,omitempty" name:"ClusterParams"`

	// Transaction mode. Valid values: `0` (place and pay for an order), `1` (place an order)
	DealMode *int64 `json:"DealMode,omitempty" name:"DealMode"`

	// Parameter template ID
	ParamTemplateId *int64 `json:"ParamTemplateId,omitempty" name:"ParamTemplateId"`
}

func (r *CreateClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "DbType")
	delete(f, "DbVersion")
	delete(f, "ProjectId")
	delete(f, "Cpu")
	delete(f, "Memory")
	delete(f, "Storage")
	delete(f, "ClusterName")
	delete(f, "AdminPassword")
	delete(f, "Port")
	delete(f, "PayMode")
	delete(f, "Count")
	delete(f, "RollbackStrategy")
	delete(f, "RollbackId")
	delete(f, "OriginalClusterId")
	delete(f, "ExpectTime")
	delete(f, "ExpectTimeThresh")
	delete(f, "StorageLimit")
	delete(f, "InstanceCount")
	delete(f, "TimeSpan")
	delete(f, "TimeUnit")
	delete(f, "AutoRenewFlag")
	delete(f, "AutoVoucher")
	delete(f, "HaCount")
	delete(f, "OrderSource")
	delete(f, "ResourceTags")
	delete(f, "DbMode")
	delete(f, "MinCpu")
	delete(f, "MaxCpu")
	delete(f, "AutoPause")
	delete(f, "AutoPauseDelay")
	delete(f, "StoragePayMode")
	delete(f, "SecurityGroupIds")
	delete(f, "AlarmPolicyIds")
	delete(f, "ClusterParams")
	delete(f, "DealMode")
	delete(f, "ParamTemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateClustersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Freezing transaction ID
	// Note: this field may return null, indicating that no valid values can be obtained.
		TranId *string `json:"TranId,omitempty" name:"TranId"`

		// Order ID
	// Note: this field may return null, indicating that no valid values can be obtained.
		DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`

		// List of resource IDs (This field has been deprecated. Please use `dealNames` in the `DescribeResourcesByDealName` API to get resource IDs.)
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

		// List of cluster IDs (This field has been deprecated. Please use `dealNames` in the `DescribeResourcesByDealName` API to get cluster IDs.)
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`

		// Big order ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
		BigDealIds []*string `json:"BigDealIds,omitempty" name:"BigDealIds"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CynosdbCluster struct {

	// Cluster status
	Status *string `json:"Status,omitempty" name:"Status"`

	// Update time
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// AZ
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Cluster name
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// Region
	Region *string `json:"Region,omitempty" name:"Region"`

	// Database version
	DbVersion *string `json:"DbVersion,omitempty" name:"DbVersion"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Number of instances
	InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`

	// User `uin`
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// Engine type
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// User `appid`
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// Cluster status description
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

	// Cluster creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Billing mode. 0: pay-as-you-go; 1: monthly subscription
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// End time
	PeriodEndTime *string `json:"PeriodEndTime,omitempty" name:"PeriodEndTime"`

	// Cluster read-write VIP
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Cluster read-write vport
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// Project ID
	ProjectID *int64 `json:"ProjectID,omitempty" name:"ProjectID"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// TDSQL-C kernel version
	CynosVersion *string `json:"CynosVersion,omitempty" name:"CynosVersion"`

	// Storage capacity
	StorageLimit *int64 `json:"StorageLimit,omitempty" name:"StorageLimit"`

	// Renewal flag
	RenewFlag *int64 `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// Task in progress
	ProcessingTask *string `json:"ProcessingTask,omitempty" name:"ProcessingTask"`

	// Array of tasks in cluster
	Tasks []*ObjectTask `json:"Tasks,omitempty" name:"Tasks"`

	// Array of tags bound to cluster
	ResourceTags []*Tag `json:"ResourceTags,omitempty" name:"ResourceTags"`

	// Database type (`NORMAL` or `SERVERLESS`)
	DbMode *string `json:"DbMode,omitempty" name:"DbMode"`

	// Serverless cluster status when the database type is `SERVERLESS`. Valid values:
	// resume
	// pause
	ServerlessStatus *string `json:"ServerlessStatus,omitempty" name:"ServerlessStatus"`

	// Prepaid cluster storage
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// Cluster storage ID used in prepaid storage modification
	StorageId *string `json:"StorageId,omitempty" name:"StorageId"`

	// Billing mode of cluster storage. Valid values: `0` (postpaid), `1` (prepaid)
	StoragePayMode *int64 `json:"StoragePayMode,omitempty" name:"StoragePayMode"`

	// The minimum storage corresponding to the compute specifications of the cluster
	MinStorageSize *int64 `json:"MinStorageSize,omitempty" name:"MinStorageSize"`

	// The maximum storage corresponding to the compute specifications of the cluster
	MaxStorageSize *int64 `json:"MaxStorageSize,omitempty" name:"MaxStorageSize"`
}

type CynosdbClusterDetail struct {

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Cluster name
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// Region
	Region *string `json:"Region,omitempty" name:"Region"`

	// Status
	Status *string `json:"Status,omitempty" name:"Status"`

	// Status description
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

	// VPC name
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// Unique VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet name
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Character set
	Charset *string `json:"Charset,omitempty" name:"Charset"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Database type
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// Database version
	DbVersion *string `json:"DbVersion,omitempty" name:"DbVersion"`

	// Used capacity
	UsedStorage *int64 `json:"UsedStorage,omitempty" name:"UsedStorage"`

	// vport for read/write separation
	RoAddr []*Addr `json:"RoAddr,omitempty" name:"RoAddr"`

	// Instance information
	InstanceSet []*ClusterInstanceDetail `json:"InstanceSet,omitempty" name:"InstanceSet"`

	// Billing mode
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// Expiration time
	PeriodEndTime *string `json:"PeriodEndTime,omitempty" name:"PeriodEndTime"`

	// VIP
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// vport
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// Project ID
	ProjectID *int64 `json:"ProjectID,omitempty" name:"ProjectID"`

	// AZ
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Array of tags bound to instance
	ResourceTags []*Tag `json:"ResourceTags,omitempty" name:"ResourceTags"`

	// Serverless cluster status when the database type is `SERVERLESS`. Valid values:
	// resume
	// resuming
	// pause
	// pausing
	ServerlessStatus *string `json:"ServerlessStatus,omitempty" name:"ServerlessStatus"`
}

type CynosdbInstance struct {

	// User `Uin`
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// User `AppId`
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Cluster name
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Region
	Region *string `json:"Region,omitempty" name:"Region"`

	// AZ
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Instance status
	Status *string `json:"Status,omitempty" name:"Status"`

	// Instance status description
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

	// Database type
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// Database version
	DbVersion *string `json:"DbVersion,omitempty" name:"DbVersion"`

	// Number of CPU cores
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// Memory in GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Storage capacity in GB
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// Instance type
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Current instance role
	InstanceRole *string `json:"InstanceRole,omitempty" name:"InstanceRole"`

	// Update time
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Private IP of instance
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Private port of instance
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// Billing mode
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// Instance expiration time
	PeriodEndTime *string `json:"PeriodEndTime,omitempty" name:"PeriodEndTime"`

	// Termination deadline
	DestroyDeadlineText *string `json:"DestroyDeadlineText,omitempty" name:"DestroyDeadlineText"`

	// Isolation time
	IsolateTime *string `json:"IsolateTime,omitempty" name:"IsolateTime"`

	// Network type
	NetType *int64 `json:"NetType,omitempty" name:"NetType"`

	// Public domain name
	WanDomain *string `json:"WanDomain,omitempty" name:"WanDomain"`

	// Public IP
	WanIP *string `json:"WanIP,omitempty" name:"WanIP"`

	// Public port
	WanPort *int64 `json:"WanPort,omitempty" name:"WanPort"`

	// Public network status
	WanStatus *string `json:"WanStatus,omitempty" name:"WanStatus"`

	// Instance termination time
	DestroyTime *string `json:"DestroyTime,omitempty" name:"DestroyTime"`

	// TDSQL-C kernel version
	CynosVersion *string `json:"CynosVersion,omitempty" name:"CynosVersion"`

	// Task in progress
	ProcessingTask *string `json:"ProcessingTask,omitempty" name:"ProcessingTask"`

	// Renewal flag
	RenewFlag *int64 `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// Minimum number of CPU cores for serverless instance
	MinCpu *float64 `json:"MinCpu,omitempty" name:"MinCpu"`

	// Maximum number of CPU cores for serverless instance
	MaxCpu *float64 `json:"MaxCpu,omitempty" name:"MaxCpu"`

	// Serverless instance status. Valid values:
	// resume
	// pause
	ServerlessStatus *string `json:"ServerlessStatus,omitempty" name:"ServerlessStatus"`

	// Prepaid storage ID
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	StorageId *string `json:"StorageId,omitempty" name:"StorageId"`

	// Storage billing mode
	StoragePayMode *int64 `json:"StoragePayMode,omitempty" name:"StoragePayMode"`
}

type CynosdbInstanceDetail struct {

	// User `Uin`
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// User `AppId`
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Cluster name
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Region
	Region *string `json:"Region,omitempty" name:"Region"`

	// AZ
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Instance status
	Status *string `json:"Status,omitempty" name:"Status"`

	// Instance status description
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

	// Database type
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// Database version
	DbVersion *string `json:"DbVersion,omitempty" name:"DbVersion"`

	// Number of CPU cores
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// Memory in GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Storage capacity in GB
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// Instance type
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Current instance role
	InstanceRole *string `json:"InstanceRole,omitempty" name:"InstanceRole"`

	// Update time
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Billing mode
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// Instance expiration time
	PeriodEndTime *string `json:"PeriodEndTime,omitempty" name:"PeriodEndTime"`

	// Network type
	NetType *int64 `json:"NetType,omitempty" name:"NetType"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Private IP of instance
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Private port of instance
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// Public domain name of instance
	WanDomain *string `json:"WanDomain,omitempty" name:"WanDomain"`

	// Character set
	Charset *string `json:"Charset,omitempty" name:"Charset"`

	// TDSQL-C kernel version
	CynosVersion *string `json:"CynosVersion,omitempty" name:"CynosVersion"`

	// Renewal flag
	RenewFlag *int64 `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// The minimum number of CPU cores for a serverless instance
	MinCpu *float64 `json:"MinCpu,omitempty" name:"MinCpu"`

	// The maximum number of CPU cores for a serverless instance
	MaxCpu *float64 `json:"MaxCpu,omitempty" name:"MaxCpu"`

	// Serverless instance status. Valid values:
	// resume
	// pause
	ServerlessStatus *string `json:"ServerlessStatus,omitempty" name:"ServerlessStatus"`
}

type CynosdbInstanceGrp struct {

	// appId
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Creation time
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Deletion time
	DeletedTime *string `json:"DeletedTime,omitempty" name:"DeletedTime"`

	// Instance group ID
	InstanceGrpId *string `json:"InstanceGrpId,omitempty" name:"InstanceGrpId"`

	// Status
	Status *string `json:"Status,omitempty" name:"Status"`

	// Instance group type. ha: HA group; ro: RO group
	Type *string `json:"Type,omitempty" name:"Type"`

	// Update time
	UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`

	// Private IP
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Private port
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// Public domain name
	WanDomain *string `json:"WanDomain,omitempty" name:"WanDomain"`

	// Public IP
	WanIP *string `json:"WanIP,omitempty" name:"WanIP"`

	// Public port
	WanPort *int64 `json:"WanPort,omitempty" name:"WanPort"`

	// Public network status
	WanStatus *string `json:"WanStatus,omitempty" name:"WanStatus"`

	// Information of instances contained in instance group
	InstanceSet []*CynosdbInstance `json:"InstanceSet,omitempty" name:"InstanceSet"`
}

type DescribeAccountsRequest struct {
	*tchttp.BaseRequest

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// List of accounts to be filtered
	AccountNames []*string `json:"AccountNames,omitempty" name:"AccountNames"`

	// Database type. Valid values: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitempty" name:"DbType"`
}

func (r *DescribeAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "AccountNames")
	delete(f, "DbType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Database account list
		AccountSet []*Account `json:"AccountSet,omitempty" name:"AccountSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupConfigRequest struct {
	*tchttp.BaseRequest

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeBackupConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Full backup start time. Value range: [0-24*3600]. For example, 0:00 AM, 1:00 AM, and 2:00 AM are represented by 0, 3600, and 7200, respectively
		BackupTimeBeg *uint64 `json:"BackupTimeBeg,omitempty" name:"BackupTimeBeg"`

		// Full backup end time. Value range: [0-24*3600]. For example, 0:00 AM, 1:00 AM, and 2:00 AM are represented by 0, 3600, and 7200, respectively
		BackupTimeEnd *uint64 `json:"BackupTimeEnd,omitempty" name:"BackupTimeEnd"`

		// Backup retention period in seconds. Backups will be cleared after this period elapses. 7 days is represented by 3600*24*7 = 604800
		ReserveDuration *uint64 `json:"ReserveDuration,omitempty" name:"ReserveDuration"`

		// Backup frequency. It is an array of 7 elements corresponding to Monday through Sunday. full: full backup; increment: incremental backup
	// Note: this field may return null, indicating that no valid values can be obtained.
		BackupFreq []*string `json:"BackupFreq,omitempty" name:"BackupFreq"`

		// Backup mode. logic: logic backup; snapshot: snapshot backup
	// Note: this field may return null, indicating that no valid values can be obtained.
		BackupType *string `json:"BackupType,omitempty" name:"BackupType"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupListRequest struct {
	*tchttp.BaseRequest

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Backup file list offset
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Backup file list start
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Database type. Valid values: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitempty" name:"DbType"`
}

func (r *DescribeBackupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "DbType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of backup files
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Backup file list
		BackupList []*BackupFileInfo `json:"BackupList,omitempty" name:"BackupList"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterDetailRequest struct {
	*tchttp.BaseRequest

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Cluster details
		Detail *CynosdbClusterDetail `json:"Detail,omitempty" name:"Detail"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterInstanceGrpsRequest struct {
	*tchttp.BaseRequest

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterInstanceGrpsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterInstanceGrpsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterInstanceGrpsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterInstanceGrpsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of instance groups
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Instance group list
		InstanceGrpInfoList []*CynosdbInstanceGrp `json:"InstanceGrpInfoList,omitempty" name:"InstanceGrpInfoList"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterInstanceGrpsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterInstanceGrpsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClustersRequest struct {
	*tchttp.BaseRequest

	// Engine type. Valid values: MYSQL, POSTGRESQL
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// Number of returned results. Default value: 20. Maximum value: 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Record offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Sort by field. Valid values:
	// <li> CREATETIME: creation time</li>
	// <li> PERIODENDTIME: expiration time</li>
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Sorting order. Valid values:
	// <li> ASC: ascending</li>
	// <li> DESC: descending</li>
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// Filter. If more than one filter exists, the logical relationship between these filters is `AND`.
	Filters []*QueryFilter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DbType")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClustersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of clusters
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Cluster list
		ClusterSet []*CynosdbCluster `json:"ClusterSet,omitempty" name:"ClusterSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// Instance group ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Security group information
		Groups []*SecurityGroup `json:"Groups,omitempty" name:"Groups"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceDetailRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Instance details
		Detail *CynosdbInstanceDetail `json:"Detail,omitempty" name:"Detail"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceSpecsRequest struct {
	*tchttp.BaseRequest

	// Database type. Valid values: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitempty" name:"DbType"`
}

func (r *DescribeInstanceSpecsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceSpecsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DbType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceSpecsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceSpecsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Specification information
		InstanceSpecSet []*InstanceSpec `json:"InstanceSpecSet,omitempty" name:"InstanceSpecSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceSpecsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceSpecsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest

	// Number of returned results. Default value: 20. Maximum value: 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Record offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Sort by field. Valid values:
	// <li> CREATETIME: creation time</li>
	// <li> PERIODENDTIME: expiration time</li>
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Sorting order. Valid values:
	// <li> ASC: ascending</li>
	// <li> DESC: descending</li>
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// Filter. If more than one filter exists, the logical relationship between these filters is `AND`.
	Filters []*QueryFilter `json:"Filters,omitempty" name:"Filters"`

	// Engine type. Valid values: MYSQL, POSTGRESQL
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// Instance status
	Status *string `json:"Status,omitempty" name:"Status"`

	// Instance ID list
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
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
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "Filters")
	delete(f, "DbType")
	delete(f, "Status")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of instances
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Instance list
		InstanceSet []*CynosdbInstance `json:"InstanceSet,omitempty" name:"InstanceSet"`

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

type DescribeMaintainPeriodRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeMaintainPeriodRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMaintainPeriodRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMaintainPeriodRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMaintainPeriodResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Maintenance days of the week
		MaintainWeekDays []*string `json:"MaintainWeekDays,omitempty" name:"MaintainWeekDays"`

		// Maintenance start time in seconds
		MaintainStartTime *int64 `json:"MaintainStartTime,omitempty" name:"MaintainStartTime"`

		// Maintenance duration in seconds
		MaintainDuration *int64 `json:"MaintainDuration,omitempty" name:"MaintainDuration"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMaintainPeriodResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMaintainPeriodResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DescribeProjectSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Security group details
		Groups []*SecurityGroup `json:"Groups,omitempty" name:"Groups"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProjectSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourcesByDealNameRequest struct {
	*tchttp.BaseRequest

	// Order ID. (If the cluster is not delivered yet, the `DescribeResourcesByDealName` API may return the `InvalidParameterValue.DealNameNotFound` error. Please call the API again until it succeeds.)
	DealName *string `json:"DealName,omitempty" name:"DealName"`
}

func (r *DescribeResourcesByDealNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcesByDealNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DealName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourcesByDealNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourcesByDealNameResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Billable resource ID information array
		BillingResourceInfos []*BillingResourceInfo `json:"BillingResourceInfos,omitempty" name:"BillingResourceInfos"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourcesByDealNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcesByDealNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRollbackTimeRangeRequest struct {
	*tchttp.BaseRequest

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeRollbackTimeRangeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRollbackTimeRangeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRollbackTimeRangeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRollbackTimeRangeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Start time point of valid rollback time range
		TimeRangeStart *string `json:"TimeRangeStart,omitempty" name:"TimeRangeStart"`

		// End time point of valid rollback time range
		TimeRangeEnd *string `json:"TimeRangeEnd,omitempty" name:"TimeRangeEnd"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRollbackTimeRangeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRollbackTimeRangeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRollbackTimeValidityRequest struct {
	*tchttp.BaseRequest

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Expected time point to roll back to
	ExpectTime *string `json:"ExpectTime,omitempty" name:"ExpectTime"`

	// Error tolerance range for rollback time point
	ExpectTimeThresh *uint64 `json:"ExpectTimeThresh,omitempty" name:"ExpectTimeThresh"`
}

func (r *DescribeRollbackTimeValidityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRollbackTimeValidityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ExpectTime")
	delete(f, "ExpectTimeThresh")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRollbackTimeValidityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRollbackTimeValidityResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Storage `poolID`
		PoolId *uint64 `json:"PoolId,omitempty" name:"PoolId"`

		// Rollback task ID, which needs to be passed in when rolling back to this time point
		QueryId *uint64 `json:"QueryId,omitempty" name:"QueryId"`

		// Whether the time point is valid. pass: check passed; fail: check failed
		Status *string `json:"Status,omitempty" name:"Status"`

		// Suggested time point. This value takes effect only if `Status` is `fail`
		SuggestTime *string `json:"SuggestTime,omitempty" name:"SuggestTime"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRollbackTimeValidityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRollbackTimeValidityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceSpec struct {

	// Number of instance CPU cores
	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`

	// Instance memory in GB
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// Maximum instance storage capacity GB
	MaxStorageSize *uint64 `json:"MaxStorageSize,omitempty" name:"MaxStorageSize"`

	// Minimum instance storage capacity GB
	MinStorageSize *uint64 `json:"MinStorageSize,omitempty" name:"MinStorageSize"`
}

type IsolateClusterRequest struct {
	*tchttp.BaseRequest

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Database type. Valid values: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitempty" name:"DbType"`
}

func (r *IsolateClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "DbType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsolateClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type IsolateClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task flow ID
	// Note: this field may return null, indicating that no valid values can be obtained.
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// Refund order ID
	// Note: this field may return null, indicating that no valid values can be obtained.
		DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *IsolateClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IsolateInstanceRequest struct {
	*tchttp.BaseRequest

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Instance ID array
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList"`

	// Database type. Valid values: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitempty" name:"DbType"`
}

func (r *IsolateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceIdList")
	delete(f, "DbType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsolateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type IsolateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task flow ID
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// Order ID for isolated instance (prepaid instance)
	// Note: this field may return null, indicating that no valid values can be obtained.
		DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *IsolateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBackupConfigRequest struct {
	*tchttp.BaseRequest

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Full backup start time. Value range: [0-24*3600]. For example, 0:00 AM, 1:00 AM, and 2:00 AM are represented by 0, 3600, and 7200, respectively
	BackupTimeBeg *uint64 `json:"BackupTimeBeg,omitempty" name:"BackupTimeBeg"`

	// Full backup start time. Value range: [0-24*3600]. For example, 0:00 AM, 1:00 AM, and 2:00 AM are represented by 0, 3600, and 7200, respectively
	BackupTimeEnd *uint64 `json:"BackupTimeEnd,omitempty" name:"BackupTimeEnd"`

	// Backup retention period in seconds. Backups will be cleared after this period elapses. 7 days is represented by 3600*24*7 = 604800
	ReserveDuration *uint64 `json:"ReserveDuration,omitempty" name:"ReserveDuration"`

	// Backup frequency. It is an array of 7 elements corresponding to Monday through Sunday. full: full backup; increment: incremental backup
	BackupFreq []*string `json:"BackupFreq,omitempty" name:"BackupFreq"`

	// Backup mode. logic: logic backup; snapshot: snapshot backup
	BackupType *string `json:"BackupType,omitempty" name:"BackupType"`
}

func (r *ModifyBackupConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "BackupTimeBeg")
	delete(f, "BackupTimeEnd")
	delete(f, "ReserveDuration")
	delete(f, "BackupFreq")
	delete(f, "BackupType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBackupConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBackupConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBackupConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// Instance group ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// List of IDs of the security groups to be modified, which is an array of one or more security group IDs.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// AZ
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

func (r *ModifyDBInstanceSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SecurityGroupIds")
	delete(f, "Zone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceSecurityGroupsRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMaintainPeriodConfigRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Maintenance start time in seconds. For example, 03:00 AM is represented by 10800
	MaintainStartTime *int64 `json:"MaintainStartTime,omitempty" name:"MaintainStartTime"`

	// Maintenance duration in seconds. For example, one hour is represented by 3600
	MaintainDuration *int64 `json:"MaintainDuration,omitempty" name:"MaintainDuration"`

	// Maintenance days of the week
	MaintainWeekDays []*string `json:"MaintainWeekDays,omitempty" name:"MaintainWeekDays"`
}

func (r *ModifyMaintainPeriodConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMaintainPeriodConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "MaintainStartTime")
	delete(f, "MaintainDuration")
	delete(f, "MaintainWeekDays")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMaintainPeriodConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMaintainPeriodConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMaintainPeriodConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMaintainPeriodConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ObjectTask struct {

	// Auto-Incrementing task ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// Task type
	// Note: this field may return null, indicating that no valid values can be obtained.
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// Task status
	// Note: this field may return null, indicating that no valid values can be obtained.
	TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// Task ID (cluster ID | instance group ID | instance ID)
	// Note: this field may return null, indicating that no valid values can be obtained.
	ObjectId *string `json:"ObjectId,omitempty" name:"ObjectId"`

	// Task type
	// Note: this field may return null, indicating that no valid values can be obtained.
	ObjectType *string `json:"ObjectType,omitempty" name:"ObjectType"`
}

type OfflineClusterRequest struct {
	*tchttp.BaseRequest

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *OfflineClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OfflineClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OfflineClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type OfflineClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task flow ID
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OfflineClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OfflineClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OfflineInstanceRequest struct {
	*tchttp.BaseRequest

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Instance ID array
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList"`
}

func (r *OfflineInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OfflineInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OfflineInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type OfflineInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task flow ID
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OfflineInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OfflineInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ParamItem struct {

	// Parameter name
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`

	// New value
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`

	// Original value
	OldValue *string `json:"OldValue,omitempty" name:"OldValue"`
}

type PolicyRule struct {

	// Policy, which can be `ACCEPT` or `DROP`
	Action *string `json:"Action,omitempty" name:"Action"`

	// Source IP or IP range, such as 192.168.0.0/16
	CidrIp *string `json:"CidrIp,omitempty" name:"CidrIp"`

	// Port
	PortRange *string `json:"PortRange,omitempty" name:"PortRange"`

	// Network protocol, such as UDP and TCP
	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`

	// Protocol port ID or protocol port group ID.
	ServiceModule *string `json:"ServiceModule,omitempty" name:"ServiceModule"`

	// IP address ID or IP address group ID.
	AddressModule *string `json:"AddressModule,omitempty" name:"AddressModule"`

	// id
	Id *string `json:"Id,omitempty" name:"Id"`

	// Description
	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

type QueryFilter struct {

	// Search field. Valid values: "InstanceId", "ProjectId", "InstanceName", "Vip"
	Names []*string `json:"Names,omitempty" name:"Names"`

	// Search string
	Values []*string `json:"Values,omitempty" name:"Values"`

	// Whether to use exact match
	ExactMatch *bool `json:"ExactMatch,omitempty" name:"ExactMatch"`

	// Search field
	Name *string `json:"Name,omitempty" name:"Name"`
}

type SecurityGroup struct {

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Creation time in the format of yyyy-mm-dd hh:mm:ss
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Inbound rule
	Inbound []*PolicyRule `json:"Inbound,omitempty" name:"Inbound"`

	// Outbound rule
	Outbound []*PolicyRule `json:"Outbound,omitempty" name:"Outbound"`

	// Security group ID
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// Security group name
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`

	// Security group remarks
	SecurityGroupRemark *string `json:"SecurityGroupRemark,omitempty" name:"SecurityGroupRemark"`
}

type SetRenewFlagRequest struct {
	*tchttp.BaseRequest

	// ID of the instance to be manipulated
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// Auto-Renewal flag
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
}

func (r *SetRenewFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetRenewFlagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceIds")
	delete(f, "AutoRenewFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetRenewFlagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SetRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of successfully manipulated instances
		Count *int64 `json:"Count,omitempty" name:"Count"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetRenewFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetRenewFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {

	// Tag key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type UpgradeInstanceRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Database CPU
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// Database memory
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Upgrade type. Valid values: upgradeImmediate, upgradeInMaintain
	UpgradeType *string `json:"UpgradeType,omitempty" name:"UpgradeType"`

	// Storage upper limit. 0 indicates to use the standard configuration
	StorageLimit *uint64 `json:"StorageLimit,omitempty" name:"StorageLimit"`

	// Whether to automatically select a voucher. 1: yes; 0: no. Default value: 0
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// Database type. Valid values: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// Transaction mode. Valid values: `0` (place and pay for an order), `1` (place an order)
	DealMode *int64 `json:"DealMode,omitempty" name:"DealMode"`
}

func (r *UpgradeInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Cpu")
	delete(f, "Memory")
	delete(f, "UpgradeType")
	delete(f, "StorageLimit")
	delete(f, "AutoVoucher")
	delete(f, "DbType")
	delete(f, "DealMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Freezing transaction ID
	// Note: this field may return null, indicating that no valid values can be obtained.
		TranId *string `json:"TranId,omitempty" name:"TranId"`

		// Big order ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
		BigDealIds []*string `json:"BigDealIds,omitempty" name:"BigDealIds"`

		// Order ID
		DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpgradeInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
