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

package v20170312

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type Account struct {

	// Account name
	User *string `json:"User,omitempty" name:"User"`

	// Host address
	Host *string `json:"Host,omitempty" name:"Host"`
}

type AssociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// Database engine name. Valid value: `mariadb`.
	Product *string `json:"Product,omitempty" name:"Product"`

	// ID of the security group to be associated in the format of sg-efil73jd.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// ID(s) of the instance(s) to be associated in the format of tdsql-lesecurk. You can specify multiple instances.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *AssociateSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "SecurityGroupId")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateSecurityGroupsRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CancelDcnJobRequest struct {
	*tchttp.BaseRequest

	// Disaster recovery instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *CancelDcnJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelDcnJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelDcnJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CancelDcnJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task ID
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelDcnJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelDcnJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloseDBExtranetAccessRequest struct {
	*tchttp.BaseRequest

	// ID of instance for which to disable public network access. The ID is in the format of `tdsql-ow728lmc` and can be obtained through the `DescribeDBInstances` API.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Whether IPv6 is used. Default value: 0
	Ipv6Flag *int64 `json:"Ipv6Flag,omitempty" name:"Ipv6Flag"`
}

func (r *CloseDBExtranetAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseDBExtranetAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Ipv6Flag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseDBExtranetAccessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CloseDBExtranetAccessResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task ID. The task status can be queried through the `DescribeFlow` API.
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CloseDBExtranetAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseDBExtranetAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ColumnPrivilege struct {

	// Database name
	Database *string `json:"Database,omitempty" name:"Database"`

	// Table name
	Table *string `json:"Table,omitempty" name:"Table"`

	// Column name
	Column *string `json:"Column,omitempty" name:"Column"`

	// Permission information
	Privileges []*string `json:"Privileges,omitempty" name:"Privileges"`
}

type CopyAccountPrivilegesRequest struct {
	*tchttp.BaseRequest

	// Instance ID, which is in the format of `tdsql-ow728lmc` and can be obtained through the `DescribeDBInstances` API.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Source username
	SrcUserName *string `json:"SrcUserName,omitempty" name:"SrcUserName"`

	// Access host allowed for source user
	SrcHost *string `json:"SrcHost,omitempty" name:"SrcHost"`

	// Target username
	DstUserName *string `json:"DstUserName,omitempty" name:"DstUserName"`

	// Access host allowed for target user
	DstHost *string `json:"DstHost,omitempty" name:"DstHost"`

	// `ReadOnly` attribute of source account
	SrcReadOnly *string `json:"SrcReadOnly,omitempty" name:"SrcReadOnly"`

	// `ReadOnly` attribute of target account
	DstReadOnly *string `json:"DstReadOnly,omitempty" name:"DstReadOnly"`
}

func (r *CopyAccountPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyAccountPrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SrcUserName")
	delete(f, "SrcHost")
	delete(f, "DstUserName")
	delete(f, "DstHost")
	delete(f, "SrcReadOnly")
	delete(f, "DstReadOnly")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CopyAccountPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CopyAccountPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CopyAccountPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyAccountPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAccountRequest struct {
	*tchttp.BaseRequest

	// Instance ID, which is in the format of `tdsql-ow728lmc` and can be obtained through the `DescribeDBInstances` API.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Login username, which can contain 1-32 letters, digits, underscores, and hyphens.
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Host that can be logged in to, which is in the same format as the host of the MySQL account and supports wildcards, such as %, 10.%, and 10.20.%.
	Host *string `json:"Host,omitempty" name:"Host"`

	// Account password. It must contain 8-32 characters in all of the following four types: lowercase letters, uppercase letters, digits, and symbols (()~!@#$%^&*-+=_|{}[]:<>,.?/), and cannot start with a slash (/).
	Password *string `json:"Password,omitempty" name:"Password"`

	// Whether to create a read-only account. 0: no; 1: for the account's SQL requests, the secondary will be used first, and if it is unavailable, the primary will be used; 2: the secondary will be used first, and if it is unavailable, the operation will fail.
	ReadOnly *int64 `json:"ReadOnly,omitempty" name:"ReadOnly"`

	// Account remarks, which can contain 0-256 letters, digits, and common symbols.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Determines whether the secondary is unavailable based on the passed-in time
	DelayThresh *int64 `json:"DelayThresh,omitempty" name:"DelayThresh"`
}

func (r *CreateAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "Host")
	delete(f, "Password")
	delete(f, "ReadOnly")
	delete(f, "Description")
	delete(f, "DelayThresh")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAccountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Instance ID, which is passed through from the input parameters.
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// Username, which is passed through from the input parameters.
		UserName *string `json:"UserName,omitempty" name:"UserName"`

		// Host allowed for access, which is passed through from the input parameters.
		Host *string `json:"Host,omitempty" name:"Host"`

		// Passed through from the input parameters.
		ReadOnly *int64 `json:"ReadOnly,omitempty" name:"ReadOnly"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateHourDBInstanceRequest struct {
	*tchttp.BaseRequest

	// AZs to deploy instance nodes. You can specify up to two AZs.
	Zones []*string `json:"Zones,omitempty" name:"Zones"`

	// Number of nodes.
	NodeCount *int64 `json:"NodeCount,omitempty" name:"NodeCount"`

	// Memory size in GB.
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Storage size in GB.
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// Number of instances to purchase.
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// Project ID. If this parameter is not passed in, the default project will be used.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Unique ID of the network. If this parameter is not passed in, the classic network will be used.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Unique ID of the subnet. If `VpcId` is specified, this parameter is required.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Database engine version. Valid values: 10.0.10, 10.1.9, 5.7.17.
	// If this parameter is left empty, `10.1.9` will be used.
	DbVersionId *string `json:"DbVersionId,omitempty" name:"DbVersionId"`

	// Custom name of the instance.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Security group ID. If this parameter is not passed in, no security groups will be associated when the instance is created.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// Whether IPv6 is supported.
	Ipv6Flag *int64 `json:"Ipv6Flag,omitempty" name:"Ipv6Flag"`

	// Array of tag key-value pairs.
	ResourceTags []*ResourceTag `json:"ResourceTags,omitempty" name:"ResourceTags"`

	// If you create a disaster recovery instance, you need to use this parameter to specify the region of the associated primary instance so that the disaster recovery instance can sync data with the primary instance over the Data Communication Network (DCN).
	DcnRegion *string `json:"DcnRegion,omitempty" name:"DcnRegion"`

	// If you create a disaster recovery instance, you need to use this parameter to specify the ID of the associated primary instance so that the disaster recovery instance can sync data with the primary instance over the Data Communication Network (DCN).
	DcnInstanceId *string `json:"DcnInstanceId,omitempty" name:"DcnInstanceId"`

	// List of parameters. Valid values: 
	// `character_set_server` (character set; required); `lower_case_table_names` (table name case sensitivity; required; 0: case-sensitive; 1: case-insensitive);
	// `innodb_page_size` (innoDB data page size; default size: 16 KB); `sync_mode` (sync mode; 0: async; 1: strong sync; 2: downgradable strong sync; default value: 2).
	InitParams []*DBParamValue `json:"InitParams,omitempty" name:"InitParams"`

	// ID of the instance whose backup data will be rolled back to the new instance you create.
	RollbackInstanceId *string `json:"RollbackInstanceId,omitempty" name:"RollbackInstanceId"`

	// Rollback time.
	RollbackTime *string `json:"RollbackTime,omitempty" name:"RollbackTime"`
}

func (r *CreateHourDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHourDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zones")
	delete(f, "NodeCount")
	delete(f, "Memory")
	delete(f, "Storage")
	delete(f, "Count")
	delete(f, "ProjectId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "DbVersionId")
	delete(f, "InstanceName")
	delete(f, "SecurityGroupIds")
	delete(f, "Ipv6Flag")
	delete(f, "ResourceTags")
	delete(f, "DcnRegion")
	delete(f, "DcnInstanceId")
	delete(f, "InitParams")
	delete(f, "RollbackInstanceId")
	delete(f, "RollbackTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateHourDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateHourDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Order ID, which is used for calling the `DescribeOrders` API.
	//  The parameter can be used to either query order details or call the user account APIs to make another payment when this payment fails.
		DealName *string `json:"DealName,omitempty" name:"DealName"`

		// IDs of the instances you have purchased in this order. If no instance IDs are returned, you can query them with the `DescribeOrders` API. You can also use the `DescribeDBInstances` API to check whether an instance has been created successfully.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateHourDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHourDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DBInstance struct {

	// Instance ID, which uniquely identifies a TDSQL instance
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Customizable instance name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Application ID of instance
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// Project ID of instance
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Instance region name, such as ap-shanghai
	Region *string `json:"Region,omitempty" name:"Region"`

	// Instance AZ name, such as ap-guangzhou-1
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// VPC ID, which is 0 if the basic network is used
	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID, which is 0 if the basic network is used
	SubnetId *int64 `json:"SubnetId,omitempty" name:"SubnetId"`

	// Instance status. Valid values: `0` (creating), `1` (running task), `2` (running), `3` (uninitialized), `-1` (isolated), `4` (initializing), `5` (eliminating), `6` (restarting), `7` (migrating data)
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Private IP address
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Private network port
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// Domain name for public network access, which can be resolved by the public network
	WanDomain *string `json:"WanDomain,omitempty" name:"WanDomain"`

	// Public IP address, which can be accessed over the public network
	WanVip *string `json:"WanVip,omitempty" name:"WanVip"`

	// Public network port
	WanPort *int64 `json:"WanPort,omitempty" name:"WanPort"`

	// Instance creation time in the format of `2006-01-02 15:04:05`
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last updated time of instance in the format of `2006-01-02 15:04:05`
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Auto-renewal flag. 0: no, 1: yes
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// Instance expiration time in the format of `2006-01-02 15:04:05`
	PeriodEndTime *string `json:"PeriodEndTime,omitempty" name:"PeriodEndTime"`

	// Instance account
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// TDSQL version information
	TdsqlVersion *string `json:"TdsqlVersion,omitempty" name:"TdsqlVersion"`

	// Instance memory size in GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Instance storage capacity in GB
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// VPC ID in string type
	UniqueVpcId *string `json:"UniqueVpcId,omitempty" name:"UniqueVpcId"`

	// VPC subnet ID in string type
	UniqueSubnetId *string `json:"UniqueSubnetId,omitempty" name:"UniqueSubnetId"`

	// Original ID of instance (this field is obsolete and should not be depended on)
	OriginSerialId *string `json:"OriginSerialId,omitempty" name:"OriginSerialId"`

	// Number of nodes. 2: one master and one slave, 3: one master and two slaves
	NodeCount *uint64 `json:"NodeCount,omitempty" name:"NodeCount"`

	// Whether it is a temp instance. 0: no, non-zero value: yes
	IsTmp *uint64 `json:"IsTmp,omitempty" name:"IsTmp"`

	// Dedicated cluster ID. If this parameter is empty, the instance is a general instance
	ExclusterId *string `json:"ExclusterId,omitempty" name:"ExclusterId"`

	// Numeric ID of instance (this field is obsolete and should not be depended on)
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// Product type ID
	Pid *int64 `json:"Pid,omitempty" name:"Pid"`

	// Maximum QPS value
	Qps *int64 `json:"Qps,omitempty" name:"Qps"`

	// Billing mode
	// Note: this field may return null, indicating that no valid values can be obtained.
	Paymode *string `json:"Paymode,omitempty" name:"Paymode"`

	// Async task flow ID when an async task is in progress on an instance
	// Note: this field may return null, indicating that no valid values can be obtained.
	Locker *int64 `json:"Locker,omitempty" name:"Locker"`

	// Current instance running status description
	// Note: this field may return null, indicating that no valid values can be obtained.
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

	// Public network access status. 0: not enabled, 1: enabled, 2: disabled, 3: enabling
	WanStatus *int64 `json:"WanStatus,omitempty" name:"WanStatus"`

	// Whether the instance supports audit. 1: yes, 0: no
	IsAuditSupported *uint64 `json:"IsAuditSupported,omitempty" name:"IsAuditSupported"`

	// Model
	Machine *string `json:"Machine,omitempty" name:"Machine"`

	// Whether data encryption is supported. 1: yes, 0: no
	IsEncryptSupported *int64 `json:"IsEncryptSupported,omitempty" name:"IsEncryptSupported"`

	// Number of CPU cores of instance
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// Indicates whether the instance uses IPv6
	// Note: this field may return null, indicating that no valid values can be obtained.
	Ipv6Flag *uint64 `json:"Ipv6Flag,omitempty" name:"Ipv6Flag"`

	// Private network IPv6 address
	// Note: this field may return null, indicating that no valid values can be obtained.
	Vipv6 *string `json:"Vipv6,omitempty" name:"Vipv6"`

	// Public network IPv6 address
	// Note: this field may return null, indicating that no valid values can be obtained.
	WanVipv6 *string `json:"WanVipv6,omitempty" name:"WanVipv6"`

	// Public network IPv6 port
	// Note: this field may return null, indicating that no valid values can be obtained.
	WanPortIpv6 *uint64 `json:"WanPortIpv6,omitempty" name:"WanPortIpv6"`

	// Public network IPv6 status
	// Note: this field may return null, indicating that no valid values can be obtained.
	WanStatusIpv6 *uint64 `json:"WanStatusIpv6,omitempty" name:"WanStatusIpv6"`

	// Database engine
	// Note: this field may return null, indicating that no valid values can be obtained.
	DbEngine *string `json:"DbEngine,omitempty" name:"DbEngine"`

	// Database version
	// Note: this field may return null, indicating that no valid values can be obtained.
	DbVersion *string `json:"DbVersion,omitempty" name:"DbVersion"`

	// DCN type. Valid values: 0 (null), 1 (primary instance), 2 (disaster recovery instance)
	// Note: this field may return null, indicating that no valid values can be obtained.
	DcnFlag *int64 `json:"DcnFlag,omitempty" name:"DcnFlag"`

	// DCN status. Valid values: 0 (null), 1 (creating), 2 (syncing), 3 (disconnected)
	// Note: this field may return null, indicating that no valid values can be obtained.
	DcnStatus *int64 `json:"DcnStatus,omitempty" name:"DcnStatus"`

	// The number of DCN disaster recovery instances
	// Note: this field may return null, indicating that no valid values can be obtained.
	DcnDstNum *int64 `json:"DcnDstNum,omitempty" name:"DcnDstNum"`

	// Instance type. Valid values: `1` (dedicated primary instance), `2` (primary instance), `3` (disaster recovery instance), and `4` (dedicated disaster recovery instance).
	// Note: this field may return null, indicating that no valid values can be obtained.
	InstanceType *int64 `json:"InstanceType,omitempty" name:"InstanceType"`

	// Instance tag information
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ResourceTags []*ResourceTag `json:"ResourceTags,omitempty" name:"ResourceTags"`
}

type DBParamValue struct {

	// Parameter name
	Param *string `json:"Param,omitempty" name:"Param"`

	// Parameter value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type DatabasePrivilege struct {

	// Permission information
	Privileges []*string `json:"Privileges,omitempty" name:"Privileges"`

	// Database name
	Database *string `json:"Database,omitempty" name:"Database"`
}

type DcnDetailItem struct {

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Region where the instance resides
	Region *string `json:"Region,omitempty" name:"Region"`

	// Availability zone where the instance resides
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Instance IP address
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Instance IPv6 address
	Vipv6 *string `json:"Vipv6,omitempty" name:"Vipv6"`

	// Instance port
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// Instance status
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Instance status description
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

	// DCN flag. Valid values: `1` (primary), `2` (disaster recovery)
	DcnFlag *int64 `json:"DcnFlag,omitempty" name:"DcnFlag"`

	// DCN status. Valid values: `0` (none), `1` (creating), `2` (syncing), `3` (disconnected)
	DcnStatus *int64 `json:"DcnStatus,omitempty" name:"DcnStatus"`

	// Number of CPU cores of the instance
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// Instance memory capacity in GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Instance storage capacity in GB
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// Billing mode
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// Creation time of the instance in the format of 2006-01-02 15:04:05
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Expiration time of the instance in the format of 2006-01-02 15:04:05
	PeriodEndTime *string `json:"PeriodEndTime,omitempty" name:"PeriodEndTime"`

	// Instance type. Valid values: `1` (dedicated primary instance), `2` (non-dedicated primary instance), `3` (non-dedicated disaster recovery instance), `4` (dedicated disaster recovery instance)
	InstanceType *int64 `json:"InstanceType,omitempty" name:"InstanceType"`
}

type DeleteAccountRequest struct {
	*tchttp.BaseRequest

	// Instance ID, which is in the format of `tdsql-ow728lmc` and can be obtained through the `DescribeDBInstances` API.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Username
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Access host allowed for user
	Host *string `json:"Host,omitempty" name:"Host"`
}

func (r *DeleteAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "Host")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAccountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstancesRequest struct {
	*tchttp.BaseRequest

	// Queries by instance ID or IDs. Instance ID is in the format of `tdsql-ow728lmc`. Up to 100 instances can be queried in one request.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Search field name. Valid values: instancename (search by instance name), vip (search by private IP), all (search by instance ID, instance name, and private IP).
	SearchName *string `json:"SearchName,omitempty" name:"SearchName"`

	// Search keyword. Fuzzy search is supported. Multiple keywords should be separated by line breaks (`\n`).
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// Queries by project ID
	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds"`

	// Whether to search by VPC
	IsFilterVpc *bool `json:"IsFilterVpc,omitempty" name:"IsFilterVpc"`

	// VPC ID, which is valid when `IsFilterVpc` is 1
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// VPC subnet ID, which is valid when `IsFilterVpc` is 1
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Sort by field. Valid values: projectId, createtime, instancename
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: desc, asc
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// Offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Queries by `OriginSerialId`
	OriginSerialIds []*string `json:"OriginSerialIds,omitempty" name:"OriginSerialIds"`

	// Identifies whether to use the `ExclusterType` field. false: no, true: yes
	IsFilterExcluster *bool `json:"IsFilterExcluster,omitempty" name:"IsFilterExcluster"`

	// Instance cluster type. 1: non-dedicated cluster, 2: dedicated cluster, 0: all
	ExclusterType *int64 `json:"ExclusterType,omitempty" name:"ExclusterType"`

	// Filters instances by dedicated cluster ID in the format of `dbdc-4ih6uct9`
	ExclusterIds []*string `json:"ExclusterIds,omitempty" name:"ExclusterIds"`

	// Tag key used in queries
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`

	// Instance types used in filtering. Valid values: 1 (dedicated instance), 2 (primary instance), 3 (disaster recovery instance). Multiple values should be separated by commas.
	FilterInstanceType *string `json:"FilterInstanceType,omitempty" name:"FilterInstanceType"`

	// Use this filter to include instances in specific statuses
	Status []*int64 `json:"Status,omitempty" name:"Status"`

	// Use this filter to exclude instances in specific statuses
	ExcludeStatus []*int64 `json:"ExcludeStatus,omitempty" name:"ExcludeStatus"`
}

func (r *DescribeDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "SearchName")
	delete(f, "SearchKey")
	delete(f, "ProjectIds")
	delete(f, "IsFilterVpc")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OriginSerialIds")
	delete(f, "IsFilterExcluster")
	delete(f, "ExclusterType")
	delete(f, "ExclusterIds")
	delete(f, "TagKeys")
	delete(f, "FilterInstanceType")
	delete(f, "Status")
	delete(f, "ExcludeStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible instances
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Instance details list
		Instances []*DBInstance `json:"Instances,omitempty" name:"Instances"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBLogFilesRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of `tdsql-ow728lmc`.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Requested log type. Valid values: 1 (binlog), 2 (cold backup), 3 (errlog), 4 (slowlog).
	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeDBLogFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBLogFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBLogFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBLogFilesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Instance ID in the format of `tdsql-ow728lmc`.
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// Requested log type. Valid values: 1 (binlog), 2 (cold backup), 3 (errlog), 4 (slowlog).
		Type *uint64 `json:"Type,omitempty" name:"Type"`

		// Total number of requested logs
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// Information such as `uri`, `length`, and `mtime` (modification time)
		Files []*LogFileInfo `json:"Files,omitempty" name:"Files"`

		// For an instance in a VPC, this prefix plus URI can be used as the download address
		VpcPrefix *string `json:"VpcPrefix,omitempty" name:"VpcPrefix"`

		// For an instance in a common network, this prefix plus URI can be used as the download address
		NormalPrefix *string `json:"NormalPrefix,omitempty" name:"NormalPrefix"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBLogFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBLogFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// Database engine name. Valid value: `mariadb`.
	Product *string `json:"Product,omitempty" name:"Product"`

	// Instance ID
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
	delete(f, "Product")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Security group details
		Groups []*SecurityGroup `json:"Groups,omitempty" name:"Groups"`

		// Instance VIP
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		VIP *string `json:"VIP,omitempty" name:"VIP"`

		// Instance port
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		VPort *int64 `json:"VPort,omitempty" name:"VPort"`

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

type DescribeDcnDetailRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDcnDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDcnDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDcnDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDcnDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// DCN synchronization details
		DcnDetails []*DcnDetailItem `json:"DcnDetails,omitempty" name:"DcnDetails"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDcnDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDcnDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFileDownloadUrlRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Unsigned file path
	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`
}

func (r *DescribeFileDownloadUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileDownloadUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "FilePath")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFileDownloadUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFileDownloadUrlResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Signed download URL
		PreSignedUrl *string `json:"PreSignedUrl,omitempty" name:"PreSignedUrl"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFileDownloadUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileDownloadUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceNodeInfoRequest struct {
	*tchttp.BaseRequest

	// Instance ID, such as tdsql-6ltok4u9
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// The maximum number of results returned at a time. By default, there is no upper limit to this value, that is, all results can be returned.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset of the returned results. Default value: `0`.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeInstanceNodeInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceNodeInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceNodeInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceNodeInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of nodes
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Node information
		NodesInfo []*NodeInfo `json:"NodesInfo,omitempty" name:"NodesInfo"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceNodeInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceNodeInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// Database engine name. Valid value: `mariadb`.
	Product *string `json:"Product,omitempty" name:"Product"`

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
	delete(f, "Product")
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

		// Total number of security groups.
		Total *uint64 `json:"Total,omitempty" name:"Total"`

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

type DestroyHourDBInstanceRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of tdsql-avw0207d. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DestroyHourDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyHourDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyHourDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DestroyHourDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task ID, which can be used in the [DescribeFlow](https://intl.cloud.tencent.com/document/product/237/16177?from_cn_redirect=1) API to query the async task result.
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// Instance ID, which is the same as the request parameter `InstanceId`.
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DestroyHourDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyHourDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisassociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// Database engine name. Valid value: `mariadb`.
	Product *string `json:"Product,omitempty" name:"Product"`

	// Security group ID
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// Instance ID list, which is an array of one or more instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DisassociateSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "SecurityGroupId")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateSecurityGroupsRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FunctionPrivilege struct {

	// Database name
	Database *string `json:"Database,omitempty" name:"Database"`

	// Function name
	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`

	// Permission information
	Privileges []*string `json:"Privileges,omitempty" name:"Privileges"`
}

type GrantAccountPrivilegesRequest struct {
	*tchttp.BaseRequest

	// Instance ID, which is in the format of `tdsql-ow728lmc` and can be obtained through the `DescribeDBInstances` API.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Login username.
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Access host allowed for user. An account is uniquely identified by username and host.
	Host *string `json:"Host,omitempty" name:"Host"`

	// Database name. `\*` indicates that global permissions will be set (i.e., `\*.\*`), in which case the `Type` and `Object ` parameters will be ignored. If `DbName` is not `\*`, the input parameter `Type` is required.
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// Global permission. Valid values: `SELECT`, `INSERT`, `UPDATE`, `DELETE`, `CREATE`, `DROP`, `REFERENCES`, `INDEX`, `ALTER`, `CREATE TEMPORARY TABLES`, `LOCK TABLES`, `EXECUTE`, `CREATE VIEW`, `SHOW VIEW`, `CREATE ROUTINE`, `ALTER ROUTINE`, `EVENT`, `TRIGGER`, `SHOW DATABASES`, `REPLICATION CLIENT`, `REPLICATION SLAVE`. 
	// Database permission. Valid values: `SELECT`, `INSERT`, `UPDATE`, `DELETE`, `CREATE`, `DROP`, `REFERENCES`, `INDEX`, `ALTER`, `CREATE TEMPORARY TABLES`, `LOCK TABLES`, `EXECUTE`, `CREATE VIEW`, `SHOW VIEW`, `CREATE ROUTINE`, `ALTER ROUTINE`, `EVENT`, `TRIGGER`. 
	// Table/View permission. Valid values: `SELECT`, `INSERT`, `UPDATE`, `DELETE`, `CREATE`, `DROP`, `REFERENCES`, `INDEX`, `ALTER`, `CREATE VIEW`, `SHOW VIEW`, `TRIGGER`. 
	// Stored procedure/function permission. Valid values: `ALTER ROUTINE`, `EXECUTE`. 
	// Field permission. Valid values: `INSERT`, `REFERENCES`, `SELECT`, `UPDATE`.
	Privileges []*string `json:"Privileges,omitempty" name:"Privileges"`

	// Type. Valid values: table, view, proc, func, \*. If `DbName` is a specific database name and `Type` is `\*`, the permissions of the database will be set (i.e., `db.\*`), in which case the `Object` parameter will be ignored
	Type *string `json:"Type,omitempty" name:"Type"`

	// Type name. For example, if `Type` is `table`, `Object` indicates a specific table name; if both `DbName` and `Type` are specific names, it indicates a specific object name and cannot be `\*` or empty
	Object *string `json:"Object,omitempty" name:"Object"`

	// If `Type` is `table` and `ColName` is `\*`, the permissions will be granted to the table; if `ColName` is a specific field name, the permissions will be granted to the field
	ColName *string `json:"ColName,omitempty" name:"ColName"`
}

func (r *GrantAccountPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GrantAccountPrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "Host")
	delete(f, "DbName")
	delete(f, "Privileges")
	delete(f, "Type")
	delete(f, "Object")
	delete(f, "ColName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GrantAccountPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GrantAccountPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GrantAccountPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GrantAccountPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogFileInfo struct {

	// Last modified time of log
	Mtime *uint64 `json:"Mtime,omitempty" name:"Mtime"`

	// File length
	Length *uint64 `json:"Length,omitempty" name:"Length"`

	// Uniform resource identifier (URI) used during log download
	Uri *string `json:"Uri,omitempty" name:"Uri"`

	// Filename
	FileName *string `json:"FileName,omitempty" name:"FileName"`
}

type ModifyAccountDescriptionRequest struct {
	*tchttp.BaseRequest

	// Instance ID, which is in the format of `tdsql-ow728lmc` and can be obtained through the `DescribeDBInstances` API.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Login username.
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Access host allowed for user. An account is uniquely identified by username and host.
	Host *string `json:"Host,omitempty" name:"Host"`

	// New account remarks, which can contain 0-256 characters.
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyAccountDescriptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountDescriptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "Host")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccountDescriptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccountDescriptionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccountDescriptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountDescriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccountPrivilegesRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of tdsql-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Database account, including username and host address.
	Accounts []*Account `json:"Accounts,omitempty" name:"Accounts"`

	// Global permission. Valid values of `GlobalPrivileges`: `"SELECT"`, `"INSERT"`, `"UPDATE"`, `"DELETE"`, `"CREATE"`, `"PROCESS"`, `"DROP"`, `"REFERENCES"`, `"INDEX"`, `"ALTER"`, `"SHOW DATABASES"`, `"CREATE TEMPORARY TABLES"`, `"LOCK TABLES"`, `"EXECUTE"`, `"CREATE VIEW"`, `"SHOW VIEW"`, `"CREATE ROUTINE"`, `"ALTER ROUTINE"`, `"EVENT"`, `"TRIGGER"`.
	// Note: if the parameter is left empty, no change will be made to the granted global permissions. To clear the granted global permissions, set the parameter to an empty array.
	GlobalPrivileges []*string `json:"GlobalPrivileges,omitempty" name:"GlobalPrivileges"`

	// Database permission. Valid values of `Privileges`: `"SELECT"`, `"INSERT"`, `"UPDATE"`, `"DELETE"`, `"CREATE"`, `"DROP"`, `"REFERENCES"`, `"INDEX"`, `"ALTER"`, `"CREATE TEMPORARY TABLES"`, `"LOCK TABLES"`, `"EXECUTE"`, `"CREATE VIEW"`, `"SHOW VIEW"`, `"CREATE ROUTINE"`, `"ALTER ROUTINE"`, `"EVENT"`, `"TRIGGER"`.
	// Note: if the parameter is left empty, no change will be made to the granted database permissions. To clear the granted database permissions, set `Privileges` to an empty array.
	DatabasePrivileges []*DatabasePrivilege `json:"DatabasePrivileges,omitempty" name:"DatabasePrivileges"`

	// Database table permission. Valid values of `Privileges`: `SELECT`, `INSERT`, `UPDATE`, `DELETE`, `CREATE`, `DROP`, `REFERENCES`, `INDEX`, `ALTER`, `CREATE VIEW`, `SHOW VIEW`, `TRIGGER`.
	// Note: if the parameter is not passed in, no change will be made to the granted table permissions. To clear the granted table permissions, set `Privileges` to an empty array.
	TablePrivileges []*TablePrivilege `json:"TablePrivileges,omitempty" name:"TablePrivileges"`

	// Column permission. Valid values of `Privileges`: `"SELECT"`, `"INSERT"`, `"UPDATE"`, `"REFERENCES"`.
	// Note: if the parameter is left empty, no change will be made to the granted column permissions. To clear the granted column permissions, set `Privileges` to an empty array.
	ColumnPrivileges []*ColumnPrivilege `json:"ColumnPrivileges,omitempty" name:"ColumnPrivileges"`

	// Database view permission. Valid values of `Privileges`: `SELECT`, `INSERT`, `UPDATE`, `DELETE`, `CREATE`, `DROP`, `REFERENCES`, `INDEX`, `ALTER`, `CREATE VIEW`, `SHOW VIEW`, `TRIGGER`.
	// Note: if the parameter is not passed in, no change will be made to the granted view permissions. To clear the granted view permissions, set `Privileges` to an empty array.
	ViewPrivileges []*ViewPrivileges `json:"ViewPrivileges,omitempty" name:"ViewPrivileges"`

	// Database function permissions. Valid values of `Privileges`: `ALTER ROUTINE`, `EXECUTE`.
	// Note: if the parameter is not passed in, no change will be made to the granted function permissions. To clear the granted function permissions, set `Privileges` to an empty array.
	FunctionPrivileges []*FunctionPrivilege `json:"FunctionPrivileges,omitempty" name:"FunctionPrivileges"`

	// Database stored procedure permission. Valid values of `Privileges`: `ALTER ROUTINE`, `EXECUTE`.
	// Note: if the parameter is not passed in, no change will be made to the granted stored procedure permissions. To clear the granted stored procedure permissions, set `Privileges` to an empty array.
	ProcedurePrivileges []*ProcedurePrivilege `json:"ProcedurePrivileges,omitempty" name:"ProcedurePrivileges"`
}

func (r *ModifyAccountPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountPrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Accounts")
	delete(f, "GlobalPrivileges")
	delete(f, "DatabasePrivileges")
	delete(f, "TablePrivileges")
	delete(f, "ColumnPrivileges")
	delete(f, "ViewPrivileges")
	delete(f, "FunctionPrivileges")
	delete(f, "ProcedurePrivileges")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccountPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccountPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task ID, which can be used in the [DescribeFlow](https://intl.cloud.tencent.com/document/product/237/16177?from_cn_redirect=1) API to query the async task result.
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccountPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstancesProjectRequest struct {
	*tchttp.BaseRequest

	// List of IDs of instances to be modified. The ID is in the format of `tdsql-ow728lmc` and can be obtained through the `DescribeDBInstances` API.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// ID of the project to be assigned, which can be obtained through the `DescribeProjects` API.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *ModifyDBInstancesProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstancesProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstancesProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstancesProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBInstancesProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstancesProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySyncTaskAttributeRequest struct {
	*tchttp.BaseRequest

	// IDs of tasks for which to modify the attributes. The IDs can be obtained by the return value of the `DescribeSyncTasks` API. Up to 100 tasks can be operated at a time.
	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds"`

	// Task name. You can specify any name you like, but its length cannot exceed 100 characters.
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
}

func (r *ModifySyncTaskAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySyncTaskAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskIds")
	delete(f, "TaskName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySyncTaskAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifySyncTaskAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySyncTaskAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySyncTaskAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NodeInfo struct {

	// Node ID
	NodeId *string `json:"NodeId,omitempty" name:"NodeId"`

	// Node role. Valid values: `master`, `slave`
	Role *string `json:"Role,omitempty" name:"Role"`
}

type ProcedurePrivilege struct {

	// Database name
	Database *string `json:"Database,omitempty" name:"Database"`

	// Stored procedure name
	Procedure *string `json:"Procedure,omitempty" name:"Procedure"`

	// Permission information
	Privileges []*string `json:"Privileges,omitempty" name:"Privileges"`
}

type ResetAccountPasswordRequest struct {
	*tchttp.BaseRequest

	// Instance ID, which is in the format of `tdsql-ow728lmc` and can be obtained through the `DescribeDBInstances` API.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Login username.
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Access host allowed for user. An account is uniquely identified by username and host.
	Host *string `json:"Host,omitempty" name:"Host"`

	// New password, which can contain 6-32 letters, digits, and common symbols but not semicolons, single quotation marks, and double quotation marks.
	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *ResetAccountPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetAccountPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "Host")
	delete(f, "Password")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetAccountPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ResetAccountPasswordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetAccountPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetAccountPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceTag struct {

	// Tag key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type SecurityGroup struct {

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Creation time in the format of yyyy-mm-dd hh:mm:ss
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Security group ID
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// Security group name
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`

	// Security group remarks
	SecurityGroupRemark *string `json:"SecurityGroupRemark,omitempty" name:"SecurityGroupRemark"`

	// Inbound rule
	Inbound []*SecurityGroupBound `json:"Inbound,omitempty" name:"Inbound"`

	// Outbound rule
	Outbound []*SecurityGroupBound `json:"Outbound,omitempty" name:"Outbound"`
}

type SecurityGroupBound struct {

	// Policy, which can be `ACCEPT` or `DROP`
	Action *string `json:"Action,omitempty" name:"Action"`

	// Source IP or source IP range, such as 192.168.0.0/16
	CidrIp *string `json:"CidrIp,omitempty" name:"CidrIp"`

	// Port
	PortRange *string `json:"PortRange,omitempty" name:"PortRange"`

	// Network protocol. UDP and TCP are supported.
	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`
}

type TablePrivilege struct {

	// Database name
	Database *string `json:"Database,omitempty" name:"Database"`

	// Table name
	Table *string `json:"Table,omitempty" name:"Table"`

	// Permission information
	Privileges []*string `json:"Privileges,omitempty" name:"Privileges"`
}

type ViewPrivileges struct {

	// Database name
	Database *string `json:"Database,omitempty" name:"Database"`

	// View name
	View *string `json:"View,omitempty" name:"View"`

	// Permission information
	Privileges []*string `json:"Privileges,omitempty" name:"Privileges"`
}
