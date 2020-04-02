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

package v20180328

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type AccountCreateInfo struct {

	// Instance username
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Instance password
	Password *string `json:"Password,omitempty" name:"Password"`

	// List of database permissions
	DBPrivileges []*DBPrivilege `json:"DBPrivileges,omitempty" name:"DBPrivileges" list`

	// Account remarks
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Whether it is an admin account. Default value: no
	IsAdmin *bool `json:"IsAdmin,omitempty" name:"IsAdmin"`
}

type AccountDetail struct {

	// Account name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Account remarks
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Account creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Account status. 1: creating, 2: normal, 3: modifying, 4: resetting password, -1: deleting
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Account update time
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Password update time
	PassTime *string `json:"PassTime,omitempty" name:"PassTime"`

	// Internal account status, which should be `enable` normally
	InternalStatus *string `json:"InternalStatus,omitempty" name:"InternalStatus"`

	// Information of read and write permissions of this account on relevant databases
	Dbs []*DBPrivilege `json:"Dbs,omitempty" name:"Dbs" list`
}

type AccountPassword struct {

	// Username
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Password
	Password *string `json:"Password,omitempty" name:"Password"`
}

type AccountPrivilege struct {

	// Database username
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Database permissions. ReadWrite: read/write, ReadOnly: read-only
	Privilege *string `json:"Privilege,omitempty" name:"Privilege"`
}

type AccountPrivilegeModifyInfo struct {

	// Database username
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Account permission change information
	DBPrivileges []*DBPrivilegeModifyInfo `json:"DBPrivileges,omitempty" name:"DBPrivileges" list`
}

type AccountRemark struct {

	// Account name
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// New remarks of account
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type Backup struct {

	// Filename
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// File size in KB
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// Backup start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Backup end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Download address for private network
	InternalAddr *string `json:"InternalAddr,omitempty" name:"InternalAddr"`

	// Download address for public network
	ExternalAddr *string `json:"ExternalAddr,omitempty" name:"ExternalAddr"`

	// Unique ID of backup file, which will be used by the `RestoreInstance` API
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// Backup file status (0: creating, 1: succeeded, 2: failed)
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// List of databases for multi-database backup
	DBs []*string `json:"DBs,omitempty" name:"DBs" list`

	// Backup policy (0: instance backup, 1: multi-database backup)
	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`

	// Backup mode. 0: scheduled, 1: manual
	BackupWay *int64 `json:"BackupWay,omitempty" name:"BackupWay"`
}

type CreateAccountRequest struct {
	*tchttp.BaseRequest

	// Database instance ID in the format of mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Database instance account information
	Accounts []*AccountCreateInfo `json:"Accounts,omitempty" name:"Accounts" list`
}

func (r *CreateAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAccountRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAccountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task flow ID
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAccountResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateBackupRequest struct {
	*tchttp.BaseRequest

	// Backup policy (0: instance backup, 1: multi-database backup)
	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`

	// List of names of databases to be backed up (required only for multi-database backup)
	DBNames []*string `json:"DBNames,omitempty" name:"DBNames" list`

	// Instance ID in the format of mssql-i1z41iwd
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *CreateBackupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateBackupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateBackupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The async job ID
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateBackupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateBackupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDBInstancesRequest struct {
	*tchttp.BaseRequest

	// Instance AZ, such as ap-guangzhou-1 (Guangzhou Zone 1). Purchasable AZs for an instance can be obtained through the `DescribeZones` API
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Instance memory size in GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Instance storage capacity in GB
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// Billing mode. Valid value: POSTPAID (pay-as-you-go).
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Number of instances purchased this time. Default value: 1. Maximum value: 10
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// VPC subnet ID in the format of subnet-bdoe83fa. `SubnetId` and `VpcId` should be set or ignored simultaneously
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// VPC ID in the format of vpc-dsp338hz. `SubnetId` and `VpcId` should be set or ignored simultaneously
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Length of purchase of instance. The default value is 1, indicating one month. The value cannot exceed 48
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// Whether to automatically use voucher. 0: no, 1: yes. Default value: no
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// Array of voucher IDs (currently, only one voucher can be used per order)
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds" list`

	// SQL Server version. Valid values: 2008R2 (SQL Server 2008 Enterprise), 2012SP3 (SQL Server 2012 Enterprise), 2016SP1 (SQL Server 2016 Enterprise), 201602 (SQL Server 2016 Standard), 2017 (SQL Server 2017 Enterprise). The version purchasable varies by region and can be queried by calling the `DescribeProductConfig` API. If this parameter is left empty, 2008R2 will be used by default.
	DBVersion *string `json:"DBVersion,omitempty" name:"DBVersion"`

	// Auto-renewal flag. 0: normal renewal, 1: auto-renewal. Default value: 1.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
}

func (r *CreateDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDBInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Order name
		DealName *string `json:"DealName,omitempty" name:"DealName"`

		// Order name array
		DealNames []*string `json:"DealNames,omitempty" name:"DealNames" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDBInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDBRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Database creation information
	DBs []*DBCreateInfo `json:"DBs,omitempty" name:"DBs" list`
}

func (r *CreateDBRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDBRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDBResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task flow ID
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDBResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDBResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateMigrationRequest struct {
	*tchttp.BaseRequest

	// Migration task name
	MigrateName *string `json:"MigrateName,omitempty" name:"MigrateName"`

	// Migration type (1: structure migration, 2: data migration, 3: incremental sync)
	MigrateType *uint64 `json:"MigrateType,omitempty" name:"MigrateType"`

	// Migration source type. 1: TencentDB for SQL Server, 2: CVM-based self-created SQL Server database; 3: SQL Server backup restoration, 4: SQL Server backup restoration (in COS mode)
	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`

	// Migration source
	Source *MigrateSource `json:"Source,omitempty" name:"Source"`

	// Migration target
	Target *MigrateTarget `json:"Target,omitempty" name:"Target"`

	// Database objects to be migrated. This parameter is not used for offline migration (SourceType=4 or SourceType=5)
	MigrateDBSet []*MigrateDB `json:"MigrateDBSet,omitempty" name:"MigrateDBSet" list`
}

func (r *CreateMigrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateMigrationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateMigrationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Migration task ID
		MigrateId *int64 `json:"MigrateId,omitempty" name:"MigrateId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMigrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateMigrationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DBCreateInfo struct {

	// Database name
	DBName *string `json:"DBName,omitempty" name:"DBName"`

	// Character set. Valid values: Chinese_PRC_CI_AS, Chinese_PRC_CS_AS, Chinese_PRC_BIN, Chinese_Taiwan_Stroke_CI_AS, SQL_Latin1_General_CP1_CI_AS, and SQL_Latin1_General_CP1_CS_AS. If this parameter is left empty, `Chinese_PRC_CI_AS` will be used by default
	Charset *string `json:"Charset,omitempty" name:"Charset"`

	// Database account permission information
	Accounts []*AccountPrivilege `json:"Accounts,omitempty" name:"Accounts" list`

	// Remarks
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type DBDetail struct {

	// Database name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Character set
	Charset *string `json:"Charset,omitempty" name:"Charset"`

	// Remarks
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Database creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Database status. 1: creating, 2: running, 3: modifying, -1: dropping
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Database account permission information
	Accounts []*AccountPrivilege `json:"Accounts,omitempty" name:"Accounts" list`

	// Internal status. ONLINE: running
	InternalStatus *string `json:"InternalStatus,omitempty" name:"InternalStatus"`
}

type DBInstance struct {

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Project ID of instance
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Instance region ID
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// Instance AZ ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// Instance VPC ID, which will be 0 if the basic network is used
	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`

	// Instance VPC subnet ID, which will be 0 if the basic network is used
	SubnetId *int64 `json:"SubnetId,omitempty" name:"SubnetId"`

	// Instance status. Valid values: <li>1: applying </li> <li>2: running </li> <li>3: restrictedly running (master/slave switching) </li> <li>4: isolated </li> <li>5: repossessing </li> <li>6: repossessed </li> <li>7: task running (e.g., backing up or rolling back the instance) </li> <li>8: decommissioned </li> <li>9: scaling </li> <li>10: migrating </li> <li>11: read-only </li> <li>12: restarting </li>
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Instance access IP
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Instance access port
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// Instance creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Instance update time
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Instance billing start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Instance billing end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Instance isolation time
	IsolateTime *string `json:"IsolateTime,omitempty" name:"IsolateTime"`

	// Instance memory size in GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Used storage capacity of instance in GB
	UsedStorage *int64 `json:"UsedStorage,omitempty" name:"UsedStorage"`

	// Instance storage capacity in GB
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// Instance version
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// Instance renewal flag
	RenewFlag *int64 `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// Instance high availability status. 1: dual-server high-availability, 2: single-server
	Model *int64 `json:"Model,omitempty" name:"Model"`

	// Instance region name, such as ap-guangzhou
	Region *string `json:"Region,omitempty" name:"Region"`

	// Instance AZ name, such as ap-guangzhou-1
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Backup time point
	BackupTime *string `json:"BackupTime,omitempty" name:"BackupTime"`

	// Instance billing mode. 0: pay-as-you-go
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// Instance UID
	Uid *string `json:"Uid,omitempty" name:"Uid"`

	// Number of CPU cores of instance
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// Instance version code
	Version *string `json:"Version,omitempty" name:"Version"`

	// Physical server code
	Type *string `json:"Type,omitempty" name:"Type"`

	// Billing ID
	Pid *int64 `json:"Pid,omitempty" name:"Pid"`

	// Unique string-type ID of instance VPC in the format of `vpc-xxx`, which is an empty string if the basic network is used
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// Unique string-type ID of instance subnet in the format of `subnet-xxx`, which is an empty string if the basic network is used
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`
}

type DBPrivilege struct {

	// Database name
	DBName *string `json:"DBName,omitempty" name:"DBName"`

	// Database permissions. ReadWrite: read/write, ReadOnly: read-only
	Privilege *string `json:"Privilege,omitempty" name:"Privilege"`
}

type DBPrivilegeModifyInfo struct {

	// Database name
	DBName *string `json:"DBName,omitempty" name:"DBName"`

	// Permission change information. ReadWrite: read/write, ReadOnly: read-only, Delete: the account has the permission to delete this database
	Privilege *string `json:"Privilege,omitempty" name:"Privilege"`
}

type DBRemark struct {

	// Database name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Remarks
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type DbRollbackTimeInfo struct {

	// Database name
	DBName *string `json:"DBName,omitempty" name:"DBName"`

	// Start time of time range available for rollback
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of time range available for rollback
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DealInfo struct {

	// Order name
	DealName *string `json:"DealName,omitempty" name:"DealName"`

	// Number of items
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// ID of associated flow, which can be used to query the flow execution status
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// This field is required only for an order that creates an instance, indicating the ID of the instance created by the order
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet" list`

	// Account
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// Instance billing type
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
}

type DeleteAccountRequest struct {
	*tchttp.BaseRequest

	// Database instance ID in the format of mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Array of instance usernames
	UserNames []*string `json:"UserNames,omitempty" name:"UserNames" list`
}

func (r *DeleteAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAccountRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAccountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task flow ID
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAccountResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteDBRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of mssql-rljoi3bf
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Array of database names
	Names []*string `json:"Names,omitempty" name:"Names" list`
}

func (r *DeleteDBRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteDBRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteDBResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task flow ID
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDBResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteDBResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMigrationRequest struct {
	*tchttp.BaseRequest

	// Migration task ID
	MigrateId *uint64 `json:"MigrateId,omitempty" name:"MigrateId"`
}

func (r *DeleteMigrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMigrationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMigrationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMigrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMigrationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountsRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Number of results per page. Value range: 1–100. Default value: 20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Page number. Default value: 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccountsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Instance ID
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// Account information list
		Accounts []*AccountDetail `json:"Accounts,omitempty" name:"Accounts" list`

		// Total number
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccountsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupsRequest struct {
	*tchttp.BaseRequest

	// Start name (yyyy-MM-dd HH:mm:ss)
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time (yyyy-MM-dd HH:mm:ss)
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Instance ID in the format of mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Number of results per page. Value range: 1–100. Default value: 20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Page number. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeBackupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBackupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of backups
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Backup list details
		Backups []*Backup `json:"Backups,omitempty" name:"Backups" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBackupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstancesRequest struct {
	*tchttp.BaseRequest

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Instance status. Valid values:
	// <li>1: applying</li>
	// <li>2: running</li>
	// <li>3: running restrictedly (master/slave switching)</li>
	// <li>4: isolated</li>
	// <li>5: repossessing</li>
	// <li>6: repossessed</li>
	// <li>7: executing task (e.g., backing up or rolling back instance)</li>
	// <li>8: deactivated</li>
	// <li>9: scaling out instance</li>
	// <li>10: migrating instance</li>
	// <li>11: read-only</li>
	// <li>12: restarting</li>
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Page number. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results per page. Value range: 1–100. Default value: 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// One or more instance IDs in the format of mssql-si2823jyl
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet" list`

	// Retrieves billing type. 0: pay-as-you-go
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// Unique string-type ID of instance VPC in the format of `vpc-xxx`. If an empty string ("") is passed in, filtering will be made by basic network.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Unique string-type ID of instance subnet in the format of `subnet-xxx`. If an empty string ("") is passed in, filtering will be made by basic network.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
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

		// Total number of eligible instances. If the results are returned in multiple pages, this value will be the number of all eligible instances but not the number of instances returned according to the current values of `Limit` and `Offset`
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Instance list
		DBInstances []*DBInstance `json:"DBInstances,omitempty" name:"DBInstances" list`

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

type DescribeDBsRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet" list`

	// Number of results per page. Value range: 1–100. Default value: 20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Page number. Default value: 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeDBsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of databases
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of instance databases
		DBInstances []*InstanceDBDetail `json:"DBInstances,omitempty" name:"DBInstances" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowStatusRequest struct {
	*tchttp.BaseRequest

	// Flow ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
}

func (r *DescribeFlowStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFlowStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Flow status. 0: succeeded, 1: failed, 2: running
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFlowStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFlowStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMigrationDetailRequest struct {
	*tchttp.BaseRequest

	// Migration task ID
	MigrateId *uint64 `json:"MigrateId,omitempty" name:"MigrateId"`
}

func (r *DescribeMigrationDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMigrationDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMigrationDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Migration task ID
		MigrateId *uint64 `json:"MigrateId,omitempty" name:"MigrateId"`

		// Migration task name
		MigrateName *string `json:"MigrateName,omitempty" name:"MigrateName"`

		// User ID of migration task
		AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

		// Migration task region
		Region *string `json:"Region,omitempty" name:"Region"`

		// Migration source type. 1: TencentDB for SQL Server, 2: CVM-based self-created SQL Server database; 3: SQL Server backup restoration, 4: SQL Server backup restoration (in COS mode)
		SourceType *int64 `json:"SourceType,omitempty" name:"SourceType"`

		// Migration task creation time
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// Migration task start time
		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

		// Migration task end time
		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

		// Migration task status (1: initializing, 4: migrating, 5: migration failed, 6: migration succeeded)
		Status *uint64 `json:"Status,omitempty" name:"Status"`

		// Migration task progress
		Progress *int64 `json:"Progress,omitempty" name:"Progress"`

		// Migration type (1: structure migration, 2: data migration, 3: incremental sync)
		MigrateType *int64 `json:"MigrateType,omitempty" name:"MigrateType"`

		// Migration source
		Source *MigrateSource `json:"Source,omitempty" name:"Source"`

		// Migration target
		Target *MigrateTarget `json:"Target,omitempty" name:"Target"`

		// Database objects to be migrated. This parameter is not used for offline migration (SourceType=4 or SourceType=5)
		MigrateDBSet []*MigrateDB `json:"MigrateDBSet,omitempty" name:"MigrateDBSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMigrationDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMigrationDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMigrationsRequest struct {
	*tchttp.BaseRequest

	// Status set. As long as a migration task is in a status therein, it will be listed
	StatusSet []*int64 `json:"StatusSet,omitempty" name:"StatusSet" list`

	// Migration task name (fuzzy match)
	MigrateName *string `json:"MigrateName,omitempty" name:"MigrateName"`

	// Number of results per page. Value range: 1–100. Default value: 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Page number. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// The query results are sorted by keyword. Valid values: name, createTime, startTime, endTime, status
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: desc, asc
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
}

func (r *DescribeMigrationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMigrationsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMigrationsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of query results
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of query results
		MigrateTaskSet []*MigrateTask `json:"MigrateTaskSet,omitempty" name:"MigrateTaskSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMigrationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMigrationsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeOrdersRequest struct {
	*tchttp.BaseRequest

	// Order array. The order name will be returned upon shipping, which can be used to call the `DescribeOrders` API to query shipment status
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames" list`
}

func (r *DescribeOrdersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeOrdersRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeOrdersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Order information array
		Deals []*DealInfo `json:"Deals,omitempty" name:"Deals" list`

		// Number of orders returned
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrdersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeOrdersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProductConfigRequest struct {
	*tchttp.BaseRequest

	// AZ ID in the format of ap-guangzhou-1
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

func (r *DescribeProductConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProductConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProductConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Specification information array
		SpecInfoList []*SpecInfo `json:"SpecInfoList,omitempty" name:"SpecInfoList" list`

		// Number of date entries returned
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProductConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRegionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of regions returned
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Region information array
		RegionSet []*RegionInfo `json:"RegionSet,omitempty" name:"RegionSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRegionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRollbackTimeRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// List of databases to be queried
	DBs []*string `json:"DBs,omitempty" name:"DBs" list`
}

func (r *DescribeRollbackTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRollbackTimeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRollbackTimeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Information of time range available for database rollback
		Details []*DbRollbackTimeInfo `json:"Details,omitempty" name:"Details" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRollbackTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRollbackTimeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowlogsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of mssql-k8voqdlz
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Query start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Query end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Number of results per page. Value range: 1–100. Default value: 20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Page number. Default value: 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeSlowlogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSlowlogsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowlogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of queries
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Information list of slow query logs
		Slowlogs []*SlowlogInfo `json:"Slowlogs,omitempty" name:"Slowlogs" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSlowlogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSlowlogsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeZonesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeZonesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeZonesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeZonesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of AZs returned
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Array of AZs
		ZoneSet []*ZoneInfo `json:"ZoneSet,omitempty" name:"ZoneSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZonesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeZonesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceCreateDBInstancesRequest struct {
	*tchttp.BaseRequest

	// AZ ID, which can be obtained through the `Zone` field in the returned value of the `DescribeZones` API
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Memory size in GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Instance capacity in GB
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// Billing type. Valid value: POSTPAID.
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// Length of purchase in months. Value range: 1–48. Default value: 1
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// Number of instances purchased at a time. Value range: 1–100. Default value: 1
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// SQL Server version. Valid values: 2008R2 (SQL Server 2008 Enterprise), 2012SP3 (SQL Server 2012 Enterprise), 2016SP1 (SQL Server 2016 Enterprise), 201602 (SQL Server 2016 Standard), 2017 (SQL Server 2017 Enterprise). Default value: 2008R2.
	DBVersion *string `json:"DBVersion,omitempty" name:"DBVersion"`
}

func (r *InquiryPriceCreateDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceCreateDBInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceCreateDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Price before discount. This value divided by 100 indicates the price; for example, 10010 means 100.10 USD
		OriginalPrice *int64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

		// The actual price to be paid. This value divided by 100 indicates the price; for example, 10010 means 100.10 USD
		Price *int64 `json:"Price,omitempty" name:"Price"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceCreateDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceCreateDBInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceUpgradeDBInstanceRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Memory size after instance upgrade in GB, which cannot be smaller than the current instance memory size
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Storage capacity after instance upgrade in GB, which cannot be smaller than the current instance storage capacity
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`
}

func (r *InquiryPriceUpgradeDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceUpgradeDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceUpgradeDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Price before discount. This value divided by 100 indicates the price; for example, 10094 means 100.94 USD
		OriginalPrice *int64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

		// The actual price to be paid. This value divided by 100 indicates the price; for example, 10094 means 100.94 USD
		Price *int64 `json:"Price,omitempty" name:"Price"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceUpgradeDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceUpgradeDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InstanceDBDetail struct {

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Database information list
	DBDetails []*DBDetail `json:"DBDetails,omitempty" name:"DBDetails" list`
}

type MigrateDB struct {

	// Name of migrated database
	DBName *string `json:"DBName,omitempty" name:"DBName"`
}

type MigrateDetail struct {

	// Name of current step
	StepName *string `json:"StepName,omitempty" name:"StepName"`

	// Progress of current step in %
	Progress *int64 `json:"Progress,omitempty" name:"Progress"`
}

type MigrateSource struct {

	// Source instance ID in the format of `mssql-si2823jyl`, which is used when `MigrateType` is 1 (TencentDB for SQL Server)
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// ID of source CVM instance, which is used when `MigrateType` is 2 (CVM-based self-created SQL Server database)
	CvmId *string `json:"CvmId,omitempty" name:"CvmId"`

	// VPC ID of source CVM instance in the format of vpc-6ys9ont9, which is used when `MigrateType` is 2 (CVM-based self-created SQL Server database)
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// VPC subnet ID of source CVM instance in the format of subnet-h9extioi, which is used when `MigrateType` is 2 (CVM-based self-created SQL Server database)
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Username, which is used when `MigrateType` is 1 or 2
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Password, which is used when `MigrateType` is 1 or 2
	Password *string `json:"Password,omitempty" name:"Password"`

	// Private IP of source CVM database, which is used when `MigrateType` is 2 (CVM-based self-created SQL Server database)
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Port number of source CVM database, which is used when `MigrateType` is 2 (CVM-based self-created SQL Server database)
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// Source backup address for offline migration, which is used when `MigrateType` is 4 or 5
	Url []*string `json:"Url,omitempty" name:"Url" list`

	// Source backup password for offline migration, which is used when `MigrateType` is 4 or 5
	UrlPassword *string `json:"UrlPassword,omitempty" name:"UrlPassword"`
}

type MigrateTarget struct {

	// ID of target instance in the format of mssql-si2823jyl
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Username of migration target instance
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Password of migration target instance
	Password *string `json:"Password,omitempty" name:"Password"`
}

type MigrateTask struct {

	// Migration task ID
	MigrateId *uint64 `json:"MigrateId,omitempty" name:"MigrateId"`

	// Migration task name
	MigrateName *string `json:"MigrateName,omitempty" name:"MigrateName"`

	// User ID of migration task
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// Migration task region
	Region *string `json:"Region,omitempty" name:"Region"`

	// Migration source type. 1: TencentDB for SQL Server, 2: CVM-based self-created SQL Server database; 3: SQL Server backup restoration, 4: SQL Server backup restoration (in COS mode)
	SourceType *int64 `json:"SourceType,omitempty" name:"SourceType"`

	// Migration task creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Migration task start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Migration task end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Migration task status (1: initializing, 4: migrating, 5: migration failed, 6: migration succeeded)
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// Information
	Message *string `json:"Message,omitempty" name:"Message"`

	// Whether migration task has been checked (0: not checked, 1: check succeeded, 2: check failed, 3: checking)
	CheckFlag *uint64 `json:"CheckFlag,omitempty" name:"CheckFlag"`

	// Migration task progress in %
	Progress *int64 `json:"Progress,omitempty" name:"Progress"`

	// Migration task progress details
	MigrateDetail *MigrateDetail `json:"MigrateDetail,omitempty" name:"MigrateDetail"`
}

type ModifyAccountPrivilegeRequest struct {
	*tchttp.BaseRequest

	// Database instance ID in the format of mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Account permission change information
	Accounts []*AccountPrivilegeModifyInfo `json:"Accounts,omitempty" name:"Accounts" list`
}

func (r *ModifyAccountPrivilegeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAccountPrivilegeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAccountPrivilegeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task flow ID
		FlowId *uint64 `json:"FlowId,omitempty" name:"FlowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccountPrivilegeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAccountPrivilegeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAccountRemarkRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of mssql-j8kv137v
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Information of account for which to modify remarks
	Accounts []*AccountRemark `json:"Accounts,omitempty" name:"Accounts" list`
}

func (r *ModifyAccountRemarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAccountRemarkRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAccountRemarkResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccountRemarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAccountRemarkResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceNameRequest struct {
	*tchttp.BaseRequest

	// Database instance ID in the format of mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// New name of database instance
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

func (r *ModifyDBInstanceNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBInstanceNameRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceNameResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBInstanceNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBInstanceNameResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceProjectRequest struct {
	*tchttp.BaseRequest

	// Array of instance IDs in the format of mssql-j8kv137v
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet" list`

	// Project ID. If this parameter is 0, the default project will be used
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *ModifyDBInstanceProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBInstanceProjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of successfully modified instances
		Count *int64 `json:"Count,omitempty" name:"Count"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBInstanceProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBInstanceProjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBNameRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Old database name
	OldDBName *string `json:"OldDBName,omitempty" name:"OldDBName"`

	// New name of database
	NewDBName *string `json:"NewDBName,omitempty" name:"NewDBName"`
}

func (r *ModifyDBNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBNameRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBNameResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task flow ID
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBNameResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBRemarkRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of mssql-rljoi3bf
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Array of database names and remarks, where each element contains a database name and the corresponding remarks
	DBRemarks []*DBRemark `json:"DBRemarks,omitempty" name:"DBRemarks" list`
}

func (r *ModifyDBRemarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBRemarkRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBRemarkResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBRemarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBRemarkResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyMigrationRequest struct {
	*tchttp.BaseRequest

	// Migration task ID
	MigrateId *uint64 `json:"MigrateId,omitempty" name:"MigrateId"`

	// New name of migration task. If this parameter is left empty, no modification will be made
	MigrateName *string `json:"MigrateName,omitempty" name:"MigrateName"`

	// New migration type (1: structure migration, 2: data migration, 3: incremental sync). If this parameter is left empty, no modification will be made
	MigrateType *uint64 `json:"MigrateType,omitempty" name:"MigrateType"`

	// Migration source type. 1: TencentDB for SQL Server, 2: CVM-based self-created SQL Server database; 3: SQL Server backup restoration, 4: SQL Server backup restoration (in COS mode). If this parameter is left empty, no modification will be made
	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`

	// Migration source. If this parameter is left empty, no modification will be made
	Source *MigrateSource `json:"Source,omitempty" name:"Source"`

	// Migration target. If this parameter is left empty, no modification will be made
	Target *MigrateTarget `json:"Target,omitempty" name:"Target"`

	// Database objects to be migrated. This parameter is not used for offline migration (SourceType=4 or SourceType=5). If it left empty, no modification will be made
	MigrateDBSet []*MigrateDB `json:"MigrateDBSet,omitempty" name:"MigrateDBSet" list`
}

func (r *ModifyMigrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMigrationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyMigrationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Migration task ID
		MigrateId *uint64 `json:"MigrateId,omitempty" name:"MigrateId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMigrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMigrationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RegionInfo struct {

	// Region ID in the format of ap-guangzhou
	Region *string `json:"Region,omitempty" name:"Region"`

	// Region name
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// Numeric ID of region
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// Current purchasability of this region. UNAVAILABLE: not purchasable, AVAILABLE: purchasable
	RegionState *string `json:"RegionState,omitempty" name:"RegionState"`
}

type ResetAccountPasswordRequest struct {
	*tchttp.BaseRequest

	// Database instance ID in the format of mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Updated account password information array
	Accounts []*AccountPassword `json:"Accounts,omitempty" name:"Accounts" list`
}

func (r *ResetAccountPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetAccountPasswordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetAccountPasswordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// ID of async task flow for account password change
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetAccountPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetAccountPasswordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RestartDBInstanceRequest struct {
	*tchttp.BaseRequest

	// Database instance ID in the format of mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *RestartDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RestartDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RestartDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task flow ID
		FlowId *uint64 `json:"FlowId,omitempty" name:"FlowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RestartDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RestartDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RestoreInstanceRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of mssql-j8kv137v
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Backup file ID, which can be obtained through the `Id` field in the returned value of the `DescribeBackups` API
	BackupId *int64 `json:"BackupId,omitempty" name:"BackupId"`
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

		// Async flow task ID, which can be used to call the `DescribeFlowStatus` API to get the task execution status
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

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

type RollbackInstanceRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Rollback type. 0: the database rolled back overwrites the original database; 1: the database rolled back is renamed and does not overwrite the original database
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// Database to be rolled back
	DBs []*string `json:"DBs,omitempty" name:"DBs" list`

	// Target time point for rollback
	Time *string `json:"Time,omitempty" name:"Time"`
}

func (r *RollbackInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RollbackInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RollbackInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The async job ID
		FlowId *uint64 `json:"FlowId,omitempty" name:"FlowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RollbackInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RollbackInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RunMigrationRequest struct {
	*tchttp.BaseRequest

	// Migration task ID
	MigrateId *uint64 `json:"MigrateId,omitempty" name:"MigrateId"`
}

func (r *RunMigrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RunMigrationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RunMigrationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// After the migration task starts, the flow ID will be returned
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RunMigrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RunMigrationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SlowlogInfo struct {

	// Unique ID of slow query log file
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// File generation start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// File generation end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// File size in KB
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// Number of logs in file
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// Download address for private network
	InternalAddr *string `json:"InternalAddr,omitempty" name:"InternalAddr"`

	// Download address for public network
	ExternalAddr *string `json:"ExternalAddr,omitempty" name:"ExternalAddr"`

	// Status (1: success, 2: failure)
	// Note: this field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type SpecInfo struct {

	// Instance specification ID. The `SpecId` returned by `DescribeZones` together with the purchasable specification information returned by `DescribeProductConfig` can be used to find out what specifications can be purchased in a specified AZ
	SpecId *int64 `json:"SpecId,omitempty" name:"SpecId"`

	// Model ID
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// Model name
	MachineTypeName *string `json:"MachineTypeName,omitempty" name:"MachineTypeName"`

	// Database version information. Valid values: 2008R2 (SQL Server 2008 Enterprise), 2012SP3 (SQL Server 2012 Enterprise), 2016SP1 (SQL Server 2016 Enterprise), 201602 (SQL Server 2016 Standard), 2017 (SQL Server 2017 Enterprise)
	Version *string `json:"Version,omitempty" name:"Version"`

	// Version name corresponding to the `Version` field
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// Memory size in GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Number of CPU cores
	CPU *int64 `json:"CPU,omitempty" name:"CPU"`

	// Minimum disk size under this specification in GB
	MinStorage *int64 `json:"MinStorage,omitempty" name:"MinStorage"`

	// Maximum disk size under this specification in GB
	MaxStorage *int64 `json:"MaxStorage,omitempty" name:"MaxStorage"`

	// QPS of this specification
	QPS *int64 `json:"QPS,omitempty" name:"QPS"`

	// Description of this specification
	SuitInfo *string `json:"SuitInfo,omitempty" name:"SuitInfo"`

	// Pid of this specification
	Pid *int64 `json:"Pid,omitempty" name:"Pid"`

	// Pay-as-you-go Pid list corresponding to this specification
	// Note: this field may return null, indicating that no valid values can be obtained.
	PostPid []*int64 `json:"PostPid,omitempty" name:"PostPid" list`

	// Billing mode under this specification. POST: pay-as-you-go
	PayModeStatus *string `json:"PayModeStatus,omitempty" name:"PayModeStatus"`
}

type TerminateDBInstanceRequest struct {
	*tchttp.BaseRequest

	// List of instance IDs manually terminated in the format of [mssql-3l3fgqn7], which are the same as the instance IDs displayed on the TencentDB Console page
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet" list`
}

func (r *TerminateDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TerminateDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TerminateDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TerminateDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TerminateDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeDBInstanceRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of mssql-j8kv137v
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Memory size after instance upgrade in GB, which cannot be smaller than the current instance memory size
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Storage capacity after instance upgrade in GB, which cannot be smaller than the current instance storage capacity
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// Whether to automatically use vouchers. 0: no, 1: yes. Default value: 0
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// Voucher ID (currently, only one voucher can be used per order)
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds" list`
}

func (r *UpgradeDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Order name
		DealName *string `json:"DealName,omitempty" name:"DealName"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpgradeDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ZoneInfo struct {

	// AZ ID in the format of ap-guangzhou-1 (i.e., Guangzhou Zone 1)
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// AZ name
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// Numeric ID of AZ
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// ID of specification purchasable in this AZ, which, together with the returned value of the `DescribeProductConfig` API, can be used to find out the specifications currently purchasable in the AZ
	SpecId *int64 `json:"SpecId,omitempty" name:"SpecId"`

	// Information of database versions purchasable under the current AZ and specification. Valid values: 2008R2 (SQL Server 2008 Enterprise), 2012SP3 (SQL Server 2012 Enterprise), 2016SP1 (SQL Server 2016 Enterprise), 201602 (SQL Server 2016 Standard), 2017 (SQL Server 2017 Enterprise)
	Version *string `json:"Version,omitempty" name:"Version"`
}
