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

package v20170320

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type Account struct {

	// New account name
	User *string `json:"User,omitempty" name:"User"`

	// New account domain name
	Host *string `json:"Host,omitempty" name:"Host"`
}

type AccountInfo struct {

	// Account remarks
	Notes *string `json:"Notes,omitempty" name:"Notes"`

	// Account domain name
	Host *string `json:"Host,omitempty" name:"Host"`

	// Account name
	User *string `json:"User,omitempty" name:"User"`

	// Account information modification time
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// Password modification time
	ModifyPasswordTime *string `json:"ModifyPasswordTime,omitempty" name:"ModifyPasswordTime"`

	// Account creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type AddTimeWindowRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Time period available for maintenance on Monday in the format of 10:00-12:00. Each period lasts from half an hour to three hours, with the start time and end time aligned by half-hour. Up to two time periods can be set. The same rule applies below.
	Monday []*string `json:"Monday,omitempty" name:"Monday" list`

	// Maintenance time window on Tuesday
	Tuesday []*string `json:"Tuesday,omitempty" name:"Tuesday" list`

	// Maintenance time window on Wednesday
	Wednesday []*string `json:"Wednesday,omitempty" name:"Wednesday" list`

	// Maintenance time window on Thursday
	Thursday []*string `json:"Thursday,omitempty" name:"Thursday" list`

	// Maintenance time window on Friday
	Friday []*string `json:"Friday,omitempty" name:"Friday" list`

	// Maintenance time window on Saturday
	Saturday []*string `json:"Saturday,omitempty" name:"Saturday" list`

	// Maintenance time window on Sunday
	Sunday []*string `json:"Sunday,omitempty" name:"Sunday" list`
}

func (r *AddTimeWindowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddTimeWindowRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddTimeWindowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddTimeWindowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddTimeWindowResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// Security group ID.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// List of instance IDs, which is an array of one or more instance IDs.
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

type BackupConfig struct {

	// Replication mode of slave database 2. Value range: async, semi-sync
	ReplicationMode *string `json:"ReplicationMode,omitempty" name:"ReplicationMode"`

	// Name of the AZ of slave database 2, such as ap-shanghai-1
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Private IP address of slave database 2
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Access port of slave database 2
	Vport *uint64 `json:"Vport,omitempty" name:"Vport"`
}

type BackupInfo struct {

	// Backup filename
	Name *string `json:"Name,omitempty" name:"Name"`

	// Backup file size in bytes
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// Backup snapshot time in the format of yyyy-MM-dd HH:mm:ss, such as 2016-03-17 02:10:37
	Date *string `json:"Date,omitempty" name:"Date"`

	// Download address on the private network
	IntranetUrl *string `json:"IntranetUrl,omitempty" name:"IntranetUrl"`

	// Download address on the public network
	InternetUrl *string `json:"InternetUrl,omitempty" name:"InternetUrl"`

	// Log type. Valid values: `logical` (logical cold backup), `physical` (physical cold backup).
	Type *string `json:"Type,omitempty" name:"Type"`

	// Backup subtask ID, which is used when backup files are deleted
	BackupId *int64 `json:"BackupId,omitempty" name:"BackupId"`

	// Backup task status. Valid values: `SUCCESS` (backup succeeded), `FAILED` (backup failed), `RUNNING` (backup is in progress).
	Status *string `json:"Status,omitempty" name:"Status"`

	// Backup task completion time
	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`

	// (This field will be disused and is thus not recommended) backup creator. Valid values: `SYSTEM` (created by system), `Uin` (initiator's `Uin` value).
	Creator *string `json:"Creator,omitempty" name:"Creator"`

	// Backup task start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Backup method. Valid values: `full` (full backup), `partial` (partial backup).
	Method *string `json:"Method,omitempty" name:"Method"`

	// Backup mode. Valid values: `manual` (manual backup), `automatic` (automatic backup).
	Way *string `json:"Way,omitempty" name:"Way"`
}

type BackupItem struct {

	// Name of the database to be backed up
	Db *string `json:"Db,omitempty" name:"Db"`

	// Name of the table to be backed up. If this parameter is passed in, the specified table in the database will be backed up; otherwise, the database will be backed up.
	Table *string `json:"Table,omitempty" name:"Table"`
}

type BackupSummaryItem struct {

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Number of automatic data backups of an instance.
	AutoBackupCount *int64 `json:"AutoBackupCount,omitempty" name:"AutoBackupCount"`

	// Capacity of automatic data backups of an instance.
	AutoBackupVolume *int64 `json:"AutoBackupVolume,omitempty" name:"AutoBackupVolume"`

	// Number of manual data backups of an instance.
	ManualBackupCount *int64 `json:"ManualBackupCount,omitempty" name:"ManualBackupCount"`

	// Capacity of manual data backups of an instance.
	ManualBackupVolume *int64 `json:"ManualBackupVolume,omitempty" name:"ManualBackupVolume"`

	// Total number of data backups of an instance (including automatic backups and manual backups).
	DataBackupCount *int64 `json:"DataBackupCount,omitempty" name:"DataBackupCount"`

	// Total capacity of data backups of an instance.
	DataBackupVolume *int64 `json:"DataBackupVolume,omitempty" name:"DataBackupVolume"`

	// Number of log backups of an instance.
	BinlogBackupCount *int64 `json:"BinlogBackupCount,omitempty" name:"BinlogBackupCount"`

	// Capacity of log backups of an instance.
	BinlogBackupVolume *int64 `json:"BinlogBackupVolume,omitempty" name:"BinlogBackupVolume"`

	// Total capacity of backups of an instance (including data backups and log backups).
	BackupVolume *int64 `json:"BackupVolume,omitempty" name:"BackupVolume"`
}

type BalanceRoGroupLoadRequest struct {
	*tchttp.BaseRequest

	// RO group ID in the format of `cdbrg-c1nl9rpv`.
	RoGroupId *string `json:"RoGroupId,omitempty" name:"RoGroupId"`
}

func (r *BalanceRoGroupLoadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BalanceRoGroupLoadRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BalanceRoGroupLoadResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BalanceRoGroupLoadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BalanceRoGroupLoadResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BinlogInfo struct {

	// Binlog backup filename
	Name *string `json:"Name,omitempty" name:"Name"`

	// Backup file size in bytes
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// File stored time in the format of 2016-03-17 02:10:37
	Date *string `json:"Date,omitempty" name:"Date"`

	// Download address on the private network
	IntranetUrl *string `json:"IntranetUrl,omitempty" name:"IntranetUrl"`

	// Download address on the public network
	InternetUrl *string `json:"InternetUrl,omitempty" name:"InternetUrl"`

	// Log type. Value range: binlog
	Type *string `json:"Type,omitempty" name:"Type"`

	// Binlog file start file
	BinlogStartTime *string `json:"BinlogStartTime,omitempty" name:"BinlogStartTime"`

	// Binlog file end time
	BinlogFinishTime *string `json:"BinlogFinishTime,omitempty" name:"BinlogFinishTime"`
}

type CloseWanServiceRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://cloud.tencent.com/document/api/236/15872) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *CloseWanServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CloseWanServiceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CloseWanServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task request ID, which can be used to query the execution result of an async task.
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CloseWanServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CloseWanServiceResponse) FromJsonString(s string) error {
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
	Privileges []*string `json:"Privileges,omitempty" name:"Privileges" list`
}

type CommonTimeWindow struct {

	// Time window on Monday in the format of 02:00–06:00
	Monday *string `json:"Monday,omitempty" name:"Monday"`

	// Time window on Tuesday in the format of 02:00–06:00
	Tuesday *string `json:"Tuesday,omitempty" name:"Tuesday"`

	// Time window on Wednesday in the format of 02:00–06:00
	Wednesday *string `json:"Wednesday,omitempty" name:"Wednesday"`

	// Time window on Thursday in the format of 02:00–06:00
	Thursday *string `json:"Thursday,omitempty" name:"Thursday"`

	// Time window on Friday in the format of 02:00–06:00
	Friday *string `json:"Friday,omitempty" name:"Friday"`

	// Time window on Saturday in the format of 02:00–06:00
	Saturday *string `json:"Saturday,omitempty" name:"Saturday"`

	// Time window on Sunday in the format of 02:00–06:00
	Sunday *string `json:"Sunday,omitempty" name:"Sunday"`
}

type CreateAccountsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// TencentDB account.
	Accounts []*Account `json:"Accounts,omitempty" name:"Accounts" list`

	// Password of the new account
	Password *string `json:"Password,omitempty" name:"Password"`

	// Remarks
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAccountsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAccountsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task request ID, which can be used to query the execution result of an async task.
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAccountsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateBackupRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Backup method of target instance. Value range: logical (logical cold backup), physical (physical cold backup).
	BackupMethod *string `json:"BackupMethod,omitempty" name:"BackupMethod"`

	// Information of the table to be backed up. If this parameter is not set, the entire instance will be backed up by default. It can be set only in logical backup (i.e., BackupMethod = logical). The specified table must exist; otherwise, backup may fail.
	// For example, if you want to backup tb1 and tb2 in db1 and the entire db2, you should set the parameter as [{"Db": "db1", "Table": "tb1"}, {"Db": "db1", "Table": "tb2"}, {"Db": "db2"} ].
	BackupDBTableList []*BackupItem `json:"BackupDBTableList,omitempty" name:"BackupDBTableList" list`
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

		// Backup task ID
		BackupId *uint64 `json:"BackupId,omitempty" name:"BackupId"`

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

type CreateDBImportJobRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Filename. The file should have already been uploaded to Tencent Cloud.
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// TencentDB username
	User *string `json:"User,omitempty" name:"User"`

	// Password of a TencentDB instance user account
	Password *string `json:"Password,omitempty" name:"Password"`

	// Name of the target database. If this parameter is not passed in, no database is specified.
	DbName *string `json:"DbName,omitempty" name:"DbName"`
}

func (r *CreateDBImportJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDBImportJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDBImportJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task request ID, which can be used to query the execution result of an async task.
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDBImportJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDBImportJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDBInstanceHourRequest struct {
	*tchttp.BaseRequest

	// Number of instances. Value range: 1–100. Default value: 1.
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// Instance memory size in MB. Please use the [DescribeDBZoneConfig](https://cloud.tencent.com/document/api/236/17229) API to query the supported memory specifications.
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Instance disk size in GB. Please use the [DescribeDBZoneConfig](https://cloud.tencent.com/document/api/236/17229) API to query the supported disk specifications.
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`

	// MySQL version. Valid values: 5.5, 5.6, 5.7. Please use the [DescribeDBZoneConfig](https://cloud.tencent.com/document/api/236/17229) API to query the supported instance versions.
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// VPC ID. If this parameter is not passed in, the basic network will be selected by default. Please use the [DescribeVpcs](/document/api/215/15778) API to query the VPCs.
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// VPC subnet ID. If `UniqVpcId` is set, then `UniqSubnetId` will be required. Please use the [DescribeSubnets](/document/api/215/15784) API to query the subnet lists.
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// Project ID. If this is left empty, the default project will be used. Please use the [DescribeProject](https://cloud.tencent.com/document/product/378/4400) API to get the project ID.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// AZ information. By default, the system will automatically select an AZ. Please use the [DescribeDBZoneConfig](https://cloud.tencent.com/document/api/236/17229) API to query the supported AZs.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Instance ID, which is required and the same as the master instance ID when purchasing read-only or disaster recovery instances. Please use the [DescribeDBInstances](https://cloud.tencent.com/document/api/236/15872) API to query the instance IDs.
	MasterInstanceId *string `json:"MasterInstanceId,omitempty" name:"MasterInstanceId"`

	// Instance type. Valid values: master (master instance), dr (disaster recovery instance), ro (read-only instance). Default value: master.
	InstanceRole *string `json:"InstanceRole,omitempty" name:"InstanceRole"`

	// AZ information of the master instance, which is required for purchasing disaster recovery instances.
	MasterRegion *string `json:"MasterRegion,omitempty" name:"MasterRegion"`

	// Custom port. Value range: [1024-65535].
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Sets the root account password. Rule: the password can contain 8–64 characters and must contain at least two of the following types of characters: letters, digits, and special symbols (_+-&=!@#$%^*()). This parameter can be specified when purchasing master instances and is meaningless for read-only or disaster recovery instances.
	Password *string `json:"Password,omitempty" name:"Password"`

	// List of parameters in the format of `ParamList.0.Name=auto_increment&ParamList.0.Value=1`. You can use the [DescribeDefaultParams](https://cloud.tencent.com/document/api/236/32662) API to query the configurable parameters.
	ParamList []*ParamInfo `json:"ParamList,omitempty" name:"ParamList" list`

	// Data replication mode. Valid values: 0 (async), 1 (semi-sync), 2 (strong sync). Default value: 0. This parameter can be specified when purchasing master instances and is meaningless for read-only or disaster recovery instances.
	ProtectMode *int64 `json:"ProtectMode,omitempty" name:"ProtectMode"`

	// Multi-AZ. Valid value: 0 (single-AZ), 1 (multi-AZ). Default value: 0. This parameter can be specified when purchasing master instances and is meaningless for read-only or disaster recovery instances.
	DeployMode *int64 `json:"DeployMode,omitempty" name:"DeployMode"`

	// AZ information of slave database 1, which is the `Zone` value by default. This parameter can be specified when purchasing master instances and is meaningless for read-only or disaster recovery instances.
	SlaveZone *string `json:"SlaveZone,omitempty" name:"SlaveZone"`

	// AZ information of slave database 2, which is empty by default. This parameter can be specified when purchasing strong sync master instances and is meaningless for other types of instances.
	BackupZone *string `json:"BackupZone,omitempty" name:"BackupZone"`

	// Security group parameter. You can use the [DescribeProjectSecurityGroups](https://cloud.tencent.com/document/api/236/15850) API to query the security group details of a project.
	SecurityGroup []*string `json:"SecurityGroup,omitempty" name:"SecurityGroup" list`

	// Read-only instance information. This parameter must be passed in when purchasing read-only instances.
	RoGroup *RoGroup `json:"RoGroup,omitempty" name:"RoGroup"`

	// This field is meaningless when purchasing pay-as-you-go instances.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// Instance name.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Instance tag information.
	ResourceTags []*TagInfo `json:"ResourceTags,omitempty" name:"ResourceTags" list`

	// Placement group ID.
	DeployGroupId *string `json:"DeployGroupId,omitempty" name:"DeployGroupId"`

	// A string that is used to guarantee the idempotency of the request, which is generated by the user and must be unique in each request on the same day. The maximum length is 64 ASCII characters. If this parameter is not specified, the idempotency of the request cannot be guaranteed.
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// Instance type. Valid values: HA (High-Availability Edition), BASIC (Basic Edition). If this parameter is not specified, High-Availability Edition will be used by default.
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`
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

		// Short order ID.
		DealIds []*string `json:"DealIds,omitempty" name:"DealIds" list`

		// Instance ID list
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

type CreateDeployGroupRequest struct {
	*tchttp.BaseRequest

	// Name of a placement group, which can contain up to 60 characters.
	DeployGroupName *string `json:"DeployGroupName,omitempty" name:"DeployGroupName"`

	// Description of a placement group, which can contain up to 200 characters.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Affinity policy of placement group. Currently, the value of this parameter can only be 1. Policy 1 indicates the upper limit of instances on one physical machine.
	Affinity []*int64 `json:"Affinity,omitempty" name:"Affinity" list`

	// Upper limit of instances on one physical machine as defined in affinity policy 1 of placement group.
	LimitNum *int64 `json:"LimitNum,omitempty" name:"LimitNum"`

	// Model attribute of placement group. Valid values: SH12+SH02, TS85.
	DevClass []*string `json:"DevClass,omitempty" name:"DevClass" list`
}

func (r *CreateDeployGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDeployGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDeployGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Placement group ID.
		DeployGroupId *string `json:"DeployGroupId,omitempty" name:"DeployGroupId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDeployGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDeployGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateParamTemplateRequest struct {
	*tchttp.BaseRequest

	// Parameter template name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Parameter template description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// MySQL version number.
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// Source parameter template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// List of parameters.
	ParamList []*Parameter `json:"ParamList,omitempty" name:"ParamList" list`
}

func (r *CreateParamTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateParamTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateParamTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Parameter template ID.
		TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateParamTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateParamTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DBSwitchInfo struct {

	// Switch time in the format of yyyy-MM-dd HH:mm:ss, such as 2017-09-03 01:34:31
	SwitchTime *string `json:"SwitchTime,omitempty" name:"SwitchTime"`

	// Switch type. Value range: TRANSFER (data migration), MASTER2SLAVE (master/slave switch), RECOVERY (master/slave recovery)
	SwitchType *string `json:"SwitchType,omitempty" name:"SwitchType"`
}

type DatabaseName struct {

	// Name of a database
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`
}

type DatabasePrivilege struct {

	// Permission information
	Privileges []*string `json:"Privileges,omitempty" name:"Privileges" list`

	// Database name
	Database *string `json:"Database,omitempty" name:"Database"`
}

type DeleteAccountsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// TencentDB account.
	Accounts []*Account `json:"Accounts,omitempty" name:"Accounts" list`
}

func (r *DeleteAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAccountsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAccountsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task request ID, which can be used to query the execution result of an async task.
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAccountsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteBackupRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Backup task ID, which is the task ID returned by the [TencentDB instance backup creating API](https://cloud.tencent.com/document/api/236/15844).
	BackupId *int64 `json:"BackupId,omitempty" name:"BackupId"`
}

func (r *DeleteBackupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteBackupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteBackupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBackupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteBackupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteDeployGroupsRequest struct {
	*tchttp.BaseRequest

	// List of IDs of placement groups to be deleted.
	DeployGroupIds []*string `json:"DeployGroupIds,omitempty" name:"DeployGroupIds" list`
}

func (r *DeleteDeployGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteDeployGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteDeployGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDeployGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteDeployGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteParamTemplateRequest struct {
	*tchttp.BaseRequest

	// Parameter template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DeleteParamTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteParamTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteParamTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteParamTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteParamTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTimeWindowRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DeleteTimeWindowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTimeWindowRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTimeWindowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTimeWindowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTimeWindowResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeployGroupInfo struct {

	// ID of a placement group.
	DeployGroupId *string `json:"DeployGroupId,omitempty" name:"DeployGroupId"`

	// Name of a placement group.
	DeployGroupName *string `json:"DeployGroupName,omitempty" name:"DeployGroupName"`

	// Creation time.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Instance quota of placement group, indicating the maximum number of instances that can be placed in one placement group.
	Quota *int64 `json:"Quota,omitempty" name:"Quota"`

	// Affinity policy of placement group. Currently, only policy 1 is supported, indicating to distribute instances across physical machines.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Affinity *string `json:"Affinity,omitempty" name:"Affinity"`

	// Upper limit of instances in one placement group on one physical machine as defined in affinity policy 1 of placement group.
	// Note: this field may return null, indicating that no valid values can be obtained.
	LimitNum *int64 `json:"LimitNum,omitempty" name:"LimitNum"`

	// Placement group details.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Physical model attribute of placement group.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DevClass *string `json:"DevClass,omitempty" name:"DevClass"`
}

type DescribeAccountPrivilegesRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Database user account.
	User *string `json:"User,omitempty" name:"User"`

	// Database account domain name.
	Host *string `json:"Host,omitempty" name:"Host"`
}

func (r *DescribeAccountPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccountPrivilegesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Array of global permissions.
		GlobalPrivileges []*string `json:"GlobalPrivileges,omitempty" name:"GlobalPrivileges" list`

		// Array of database permissions.
		DatabasePrivileges []*DatabasePrivilege `json:"DatabasePrivileges,omitempty" name:"DatabasePrivileges" list`

		// Array of table permissions in the database.
		TablePrivileges []*TablePrivilege `json:"TablePrivileges,omitempty" name:"TablePrivileges" list`

		// Array of column permissions in the table.
		ColumnPrivileges []*ColumnPrivilege `json:"ColumnPrivileges,omitempty" name:"ColumnPrivileges" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccountPrivilegesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Record offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results to be returned for a single request. Value range: 1-100. Default value: 20.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Regular expression for matching account names, which complies with the rules at MySQL official website.
	AccountRegexp *string `json:"AccountRegexp,omitempty" name:"AccountRegexp"`
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

		// Number of eligible accounts.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Details of eligible accounts.
		Items []*AccountInfo `json:"Items,omitempty" name:"Items" list`

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

type DescribeAsyncRequestInfoRequest struct {
	*tchttp.BaseRequest

	// Async task request ID.
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`
}

func (r *DescribeAsyncRequestInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAsyncRequestInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAsyncRequestInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task execution result. Valid values: INITIAL, RUNNING, SUCCESS, FAILED, KILLED, REMOVED, PAUSED.
	// Note: This field may return null, indicating that no valid values can be obtained.
		Status *string `json:"Status,omitempty" name:"Status"`

		// Task execution information.
	// Note: This field may return null, indicating that no valid values can be obtained.
		Info *string `json:"Info,omitempty" name:"Info"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAsyncRequestInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAsyncRequestInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupConfigRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeBackupConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBackupConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Earliest start time point of automatic backup, such as 2 (for 2:00 AM). (This field has been disused. You are recommended to use the `BackupTimeWindow` field)
		StartTimeMin *int64 `json:"StartTimeMin,omitempty" name:"StartTimeMin"`

		// Latest start time point of automatic backup, such as 6 (for 6:00 AM). (This field has been disused. You are recommended to use the `BackupTimeWindow` field)
		StartTimeMax *int64 `json:"StartTimeMax,omitempty" name:"StartTimeMax"`

		// Backup file retention period in days.
		BackupExpireDays *int64 `json:"BackupExpireDays,omitempty" name:"BackupExpireDays"`

		// Backup mode. Value range: physical, logical
		BackupMethod *string `json:"BackupMethod,omitempty" name:"BackupMethod"`

		// Binlog file retention period in days.
		BinlogExpireDays *int64 `json:"BinlogExpireDays,omitempty" name:"BinlogExpireDays"`

		// Time window for automatic instance backup.
		BackupTimeWindow *CommonTimeWindow `json:"BackupTimeWindow,omitempty" name:"BackupTimeWindow"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBackupConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupDatabasesRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Start time in the format of yyyy-MM-dd HH:mm:ss, such as 2017-07-12 10:29:20.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Prefix of the database to be queried.
	SearchDatabase *string `json:"SearchDatabase,omitempty" name:"SearchDatabase"`

	// Pagination offset.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page. Value range: 1-2,000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeBackupDatabasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBackupDatabasesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupDatabasesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of the returned data entries.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Array of eligible databases.
		Items []*DatabaseName `json:"Items,omitempty" name:"Items" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupDatabasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBackupDatabasesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupOverviewRequest struct {
	*tchttp.BaseRequest

	// TencentDB product type to be queried. Currently, only `mysql` is supported.
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeBackupOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBackupOverviewRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupOverviewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of backups of a user in the current region (including data backups and log backups).
		BackupCount *int64 `json:"BackupCount,omitempty" name:"BackupCount"`

		// Total capacity of backups of a user in the current region.
		BackupVolume *int64 `json:"BackupVolume,omitempty" name:"BackupVolume"`

		// Paid capacity of backups of a user in the current region, i.e., capacity that exceeds the free tier.
		BillingVolume *int64 `json:"BillingVolume,omitempty" name:"BillingVolume"`

		// Backup capacity in the free tier of a user in the current region.
		FreeVolume *int64 `json:"FreeVolume,omitempty" name:"FreeVolume"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBackupOverviewResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupSummariesRequest struct {
	*tchttp.BaseRequest

	// TencentDB product type to be queried. Currently, only `mysql` is supported.
	Product *string `json:"Product,omitempty" name:"Product"`

	// Pagination offset.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Paginated query limit. Default value: 20.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Sorting criterion. Valid values: BackupVolume (backup capacity), DataBackupVolume (data backup capacity), BinlogBackupVolume (log backup capacity), AutoBackupVolume (automatic backup capacity), ManualBackupVolume (manual backup capacity).
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: ASC (ascending), DESC (descending).
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

func (r *DescribeBackupSummariesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBackupSummariesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupSummariesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Statistical items of instance backup.
		Items []*BackupSummaryItem `json:"Items,omitempty" name:"Items" list`

		// Total number of instance backups.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupSummariesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBackupSummariesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupTablesRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Start time in the format of yyyy-MM-dd HH:mm:ss, such as 2017-07-12 10:29:20.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Specified database name.
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// Prefix of the table to be queried.
	SearchTable *string `json:"SearchTable,omitempty" name:"SearchTable"`

	// Pagination offset.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page. Value range: 1-2,000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeBackupTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBackupTablesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupTablesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of the returned data entries.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Array of eligible tables.
		Items []*TableName `json:"Items,omitempty" name:"Items" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBackupTablesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Offset. Minimum value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page. Value range: 1-100. Default value: 20.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
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

		// Number of eligible instances.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Details of eligible backups.
		Items []*BackupInfo `json:"Items,omitempty" name:"Items" list`

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

type DescribeBinlogBackupOverviewRequest struct {
	*tchttp.BaseRequest

	// TencentDB product type to be queried. Currently, only `mysql` is supported.
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeBinlogBackupOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBinlogBackupOverviewRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBinlogBackupOverviewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total capacity of log backups in bytes.
		BinlogBackupVolume *int64 `json:"BinlogBackupVolume,omitempty" name:"BinlogBackupVolume"`

		// Total number of log backups.
		BinlogBackupCount *int64 `json:"BinlogBackupCount,omitempty" name:"BinlogBackupCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBinlogBackupOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBinlogBackupOverviewResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBinlogsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Offset. Minimum value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page. Value range: 1-100. Default value: 20.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeBinlogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBinlogsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBinlogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible log files.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Number of eligible binlog files.
		Items []*BinlogInfo `json:"Items,omitempty" name:"Items" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBinlogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBinlogsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBImportRecordsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Start time in the format of yyyy-MM-dd HH:mm:ss, such as 2016-01-01 00:00:01.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time in the format of yyyy-MM-dd HH:mm:ss, such as 2016-01-01 23:59:59.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Pagination parameter indicating the offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Pagination parameter indicating the number of results to be returned for a single request. Value range: 1-100. Default value: 20.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDBImportRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBImportRecordsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBImportRecordsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible import task operation logs.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of import operation records.
		Items []*ImportRecord `json:"Items,omitempty" name:"Items" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBImportRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBImportRecordsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceCharsetRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://cloud.tencent.com/document/api/236/15872) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBInstanceCharsetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstanceCharsetRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceCharsetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Default character set of the instance, such as "latin1" and "utf8".
		Charset *string `json:"Charset,omitempty" name:"Charset"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBInstanceCharsetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstanceCharsetResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceConfigRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBInstanceConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstanceConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Data protection mode of the master instance. Value range: 0 (async replication), 1 (semi-sync replication), 2 (strong sync replication).
		ProtectMode *int64 `json:"ProtectMode,omitempty" name:"ProtectMode"`

		// Master instance deployment mode. Value range: 0 (single-AZ), 1 (multi-AZ)
		DeployMode *int64 `json:"DeployMode,omitempty" name:"DeployMode"`

		// Instance AZ information in the format of "ap-shanghai-1".
		Zone *string `json:"Zone,omitempty" name:"Zone"`

		// Configuration information of the slave database.
		SlaveConfig *SlaveConfig `json:"SlaveConfig,omitempty" name:"SlaveConfig"`

		// Configuration information of slave database 2 of a strong sync instance.
		BackupConfig *BackupConfig `json:"BackupConfig,omitempty" name:"BackupConfig"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBInstanceConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstanceConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceGTIDRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://cloud.tencent.com/document/api/236/15872) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBInstanceGTIDRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstanceGTIDRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceGTIDResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// GTID enablement flag. Value range: 0 (not enabled), 1 (enabled).
		IsGTIDOpen *int64 `json:"IsGTIDOpen,omitempty" name:"IsGTIDOpen"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBInstanceGTIDResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstanceGTIDResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceRebootTimeRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *DescribeDBInstanceRebootTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstanceRebootTimeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceRebootTimeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible instances.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Returned parameter information.
		Items []*InstanceRebootTime `json:"Items,omitempty" name:"Items" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBInstanceRebootTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstanceRebootTimeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstancesRequest struct {
	*tchttp.BaseRequest

	// Project ID. You can use the [project list querying API](https://cloud.tencent.com/document/product/378/4400) to query the project ID.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Instance type. Value range: 1 (master), 2 (disaster recovery), 3 (read-only).
	InstanceTypes []*uint64 `json:"InstanceTypes,omitempty" name:"InstanceTypes" list`

	// Private IP address of the instance.
	Vips []*string `json:"Vips,omitempty" name:"Vips" list`

	// Instance status. Value range: <br>0 - creating <br>1 - running <br>4 - isolating <br>5 - isolated (the instance can be restored and started in the recycle bin)
	Status []*uint64 `json:"Status,omitempty" name:"Status" list`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results to be returned for a single request. Default value: 20. Maximum value: 2,000.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Security group ID. When it is used as a filter, the `WithSecurityGroup` parameter should be set to 1.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// Billing method. Value range: 0 (monthly subscribed), 1 (hourly).
	PayTypes []*uint64 `json:"PayTypes,omitempty" name:"PayTypes" list`

	// Instance name.
	InstanceNames []*string `json:"InstanceNames,omitempty" name:"InstanceNames" list`

	// Instance task status. Value range: <br>0 - no task <br>1 - upgrading <br>2 - importing data <br>3 - activating slave <br>4 - public network access enabled <br>5 - batch operation in progress <br>6 - rolling back <br>7 - public network access not enabled <br>8 - modifying password <br>9 - renaming instance <br>10 - restarting <br>12 - migrating self-built instance <br>13 - dropping table <br>14 - creating and syncing disaster recovery instance <br>15 - pending upgrade and switch <br>16 - upgrade and switch in progress <br>17 - upgrade and switch completed
	TaskStatus []*uint64 `json:"TaskStatus,omitempty" name:"TaskStatus" list`

	// Version of the instance database engine. Value range: 5.1, 5.5, 5.6, 5.7.
	EngineVersions []*string `json:"EngineVersions,omitempty" name:"EngineVersions" list`

	// VPC ID.
	VpcIds []*uint64 `json:"VpcIds,omitempty" name:"VpcIds" list`

	// AZ ID.
	ZoneIds []*uint64 `json:"ZoneIds,omitempty" name:"ZoneIds" list`

	// Subnet ID.
	SubnetIds []*uint64 `json:"SubnetIds,omitempty" name:"SubnetIds" list`

	// Lock flag.
	CdbErrors []*int64 `json:"CdbErrors,omitempty" name:"CdbErrors" list`

	// Sort by field of the returned result set. Currently, supported values include "InstanceId", "InstanceName", "CreateTime", and "DeadlineTime".
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Sorting method of the returned result set. Currently, "ASC" or "DESC" is supported.
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`

	// Whether security group ID is used as a filter
	WithSecurityGroup *int64 `json:"WithSecurityGroup,omitempty" name:"WithSecurityGroup"`

	// Whether dedicated cluster details are included. Value range: 0 (not included), 1 (included)
	WithExCluster *int64 `json:"WithExCluster,omitempty" name:"WithExCluster"`

	// Exclusive cluster ID.
	ExClusterId *string `json:"ExClusterId,omitempty" name:"ExClusterId"`

	// Instance ID.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// Initialization flag. Value range: 0 (not initialized), 1 (initialized).
	InitFlag *int64 `json:"InitFlag,omitempty" name:"InitFlag"`

	// Whether instances corresponding to the disaster recovery relationship are included. Valid values: 0 (not included), 1 (included). Default value: 1. If a master instance is pulled, the data of the disaster recovery relationship will be in the `DrInfo` field. If a disaster recovery instance is pulled, the data of the disaster recovery relationship will be in the `MasterInfo` field. The disaster recovery relationship contains only partial basic data. To get the detailed data, you need to call an API to pull it.
	WithDr *int64 `json:"WithDr,omitempty" name:"WithDr"`

	// Whether read-only instances are included. Valid values: 0 (not included), 1 (included). Default value: 1.
	WithRo *int64 `json:"WithRo,omitempty" name:"WithRo"`

	// Whether master instances are included. Valid values: 0 (not included), 1 (included). Default value: 1.
	WithMaster *int64 `json:"WithMaster,omitempty" name:"WithMaster"`

	// Placement group ID list.
	DeployGroupIds []*string `json:"DeployGroupIds,omitempty" name:"DeployGroupIds" list`
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
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Instance details.
		Items []*InstanceInfo `json:"Items,omitempty" name:"Items" list`

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

type DescribeDBSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
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

		// Security group details.
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

type DescribeDBSwitchRecordsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Pagination offset.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page. Value range: 1-2,000. Default value: 50.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDBSwitchRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBSwitchRecordsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSwitchRecordsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of instance switches.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Details of instance switches.
		Items []*DBSwitchInfo `json:"Items,omitempty" name:"Items" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBSwitchRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBSwitchRecordsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBZoneConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDBZoneConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBZoneConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBZoneConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of configurations in purchasable regions
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Details of configurations in purchasable regions
		Items []*RegionSellConf `json:"Items,omitempty" name:"Items" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBZoneConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBZoneConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDataBackupOverviewRequest struct {
	*tchttp.BaseRequest

	// TencentDB product type to be queried. Currently, only `mysql` is supported.
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeDataBackupOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDataBackupOverviewRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDataBackupOverviewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total capacity of data backups in bytes in the current region (including automatic backups and manual backups).
		DataBackupVolume *int64 `json:"DataBackupVolume,omitempty" name:"DataBackupVolume"`

		// Total number of data backups in the current region.
		DataBackupCount *int64 `json:"DataBackupCount,omitempty" name:"DataBackupCount"`

		// Total capacity of automatic backups in the current region.
		AutoBackupVolume *int64 `json:"AutoBackupVolume,omitempty" name:"AutoBackupVolume"`

		// Total number of automatic backups in the current region.
		AutoBackupCount *int64 `json:"AutoBackupCount,omitempty" name:"AutoBackupCount"`

		// Total capacity of manual backups in the current region.
		ManualBackupVolume *int64 `json:"ManualBackupVolume,omitempty" name:"ManualBackupVolume"`

		// Total number of manual backups in the current region.
		ManualBackupCount *int64 `json:"ManualBackupCount,omitempty" name:"ManualBackupCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDataBackupOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDataBackupOverviewResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabasesRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Offset. Minimum value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results to be returned for a single request. Value range: 1-100. Maximum value: 20.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Regular expression for matching database names.
	DatabaseRegexp *string `json:"DatabaseRegexp,omitempty" name:"DatabaseRegexp"`
}

func (r *DescribeDatabasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDatabasesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabasesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible instances.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Information of an instance.
		Items []*string `json:"Items,omitempty" name:"Items" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDatabasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDatabasesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDefaultParamsRequest struct {
	*tchttp.BaseRequest

	// MySQL version. Currently, the supported versions are ["5.1", "5.5", "5.6", "5.7"].
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`
}

func (r *DescribeDefaultParamsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDefaultParamsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDefaultParamsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of parameters
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Parameter details.
		Items []*ParameterDetail `json:"Items,omitempty" name:"Items" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDefaultParamsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDefaultParamsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDeployGroupListRequest struct {
	*tchttp.BaseRequest

	// ID of a placement group.
	DeployGroupId *string `json:"DeployGroupId,omitempty" name:"DeployGroupId"`

	// Name of a placement group.
	DeployGroupName *string `json:"DeployGroupName,omitempty" name:"DeployGroupName"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeDeployGroupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDeployGroupListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDeployGroupListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible entries.
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// List of returned results.
	// Note: This field may return null, indicating that no valid values can be obtained.
		Items []*DeployGroupInfo `json:"Items,omitempty" name:"Items" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeployGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDeployGroupListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceMonitorInfoRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// This parameter is used to return the monitoring data of Count 5-minute time periods on the day. Value range: 1-288. If this parameter is not passed in, all monitoring data in a 5-minute granularity on the day will be returned by default.
	Count *uint64 `json:"Count,omitempty" name:"Count"`
}

func (r *DescribeDeviceMonitorInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDeviceMonitorInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceMonitorInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// CPU monitoring data of the instance
		Cpu *DeviceCpuInfo `json:"Cpu,omitempty" name:"Cpu"`

		// Memory monitoring data of the instance
		Mem *DeviceMemInfo `json:"Mem,omitempty" name:"Mem"`

		// Network monitoring data of the instance
		Net *DeviceNetInfo `json:"Net,omitempty" name:"Net"`

		// Disk monitoring data of the instance
		Disk *DeviceDiskInfo `json:"Disk,omitempty" name:"Disk"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeviceMonitorInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDeviceMonitorInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeErrorLogDataRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Start timestamp.
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// End timestamp.
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// List of keywords to match. Up to 15 keywords are supported.
	KeyWords []*string `json:"KeyWords,omitempty" name:"KeyWords" list`

	// Number of results to be returned per page. Maximum value: 400.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeErrorLogDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeErrorLogDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeErrorLogDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible entries.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Returned result.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Items []*ErrlogItem `json:"Items,omitempty" name:"Items" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeErrorLogDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeErrorLogDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceParamRecordsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://cloud.tencent.com/document/api/236/15872) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Pagination offset.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
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

		// Number of eligible records.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Parameter modification records.
		Items []*ParamRecord `json:"Items,omitempty" name:"Items" list`

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

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://cloud.tencent.com/document/api/236/15872) to query the ID, whose value is the `InstanceId` value in output parameters.
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

		// Number of instance parameters.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Parameter details.
		Items []*ParameterDetail `json:"Items,omitempty" name:"Items" list`

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

type DescribeParamTemplateInfoRequest struct {
	*tchttp.BaseRequest

	// Parameter template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DescribeParamTemplateInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeParamTemplateInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeParamTemplateInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Parameter template ID.
		TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

		// Parameter template name.
		Name *string `json:"Name,omitempty" name:"Name"`

		// Parameter template description
		EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

		// Number of parameters in the parameter template
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Parameter details
		Items []*ParameterDetail `json:"Items,omitempty" name:"Items" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeParamTemplateInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeParamTemplateInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeParamTemplatesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeParamTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeParamTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeParamTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of parameter templates of the user.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Parameter template details.
		Items []*ParamTemplateInfo `json:"Items,omitempty" name:"Items" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeParamTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeParamTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
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

		// Security group details.
		Groups []*SecurityGroup `json:"Groups,omitempty" name:"Groups" list`

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

type DescribeRoGroupsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of `cdb-c1nl9rpv` or `cdb-c1nl9rpv`. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeRoGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRoGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRoGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// RO group information array. An instance can be associated with multiple RO groups.
		RoGroups []*RoGroup `json:"RoGroups,omitempty" name:"RoGroups" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRoGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRoGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRollbackRangeTimeRequest struct {
	*tchttp.BaseRequest

	// Instance ID list. An instance ID is in the format of cdb-c1nl9rpv, which is the same as the instance ID displayed on the TencentDB Console page.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *DescribeRollbackRangeTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRollbackRangeTimeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRollbackRangeTimeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible instances.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Returned parameter information.
		Items []*InstanceRollbackRangeTime `json:"Items,omitempty" name:"Items" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRollbackRangeTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRollbackRangeTimeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogDataRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Start timestamp.
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// End timestamp.
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// Client `Host` list.
	UserHosts []*string `json:"UserHosts,omitempty" name:"UserHosts" list`

	// Client username list.
	UserNames []*string `json:"UserNames,omitempty" name:"UserNames" list`

	// Accessed database list.
	DataBases []*string `json:"DataBases,omitempty" name:"DataBases" list`

	// Sort by field. Valid values: Timestamp, QueryTime, LockTime, RowsExamined, RowsSent.
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// Sorting order. Valid values: ASC (ascending), DESC (descending).
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results to be returned at a time. Maximum value: 400.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSlowLogDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSlowLogDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible entries.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Queried results.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Items []*SlowLogItem `json:"Items,omitempty" name:"Items" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSlowLogDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSlowLogDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Offset. Minimum value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page. Value range: 1-100. Default value: 20.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
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

		// Number of eligible slow logs.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Details of eligible slow logs.
		Items []*SlowLogInfo `json:"Items,omitempty" name:"Items" list`

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

type DescribeSupportedPrivilegesRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeSupportedPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSupportedPrivilegesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSupportedPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Global permissions supported by the instance
		GlobalSupportedPrivileges []*string `json:"GlobalSupportedPrivileges,omitempty" name:"GlobalSupportedPrivileges" list`

		// Database permissions supported by the instance.
		DatabaseSupportedPrivileges []*string `json:"DatabaseSupportedPrivileges,omitempty" name:"DatabaseSupportedPrivileges" list`

		// Table permissions supported by the instance.
		TableSupportedPrivileges []*string `json:"TableSupportedPrivileges,omitempty" name:"TableSupportedPrivileges" list`

		// Column permissions supported by the instance.
		ColumnSupportedPrivileges []*string `json:"ColumnSupportedPrivileges,omitempty" name:"ColumnSupportedPrivileges" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSupportedPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSupportedPrivilegesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTablesRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Database name.
	Database *string `json:"Database,omitempty" name:"Database"`

	// Record offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results to be returned for a single request. Default value: 20. Maximum value: 2,000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Regular expression for matching table names, which complies with the rules at MySQL's official website
	TableRegexp *string `json:"TableRegexp,omitempty" name:"TableRegexp"`
}

func (r *DescribeTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTablesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTablesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible tables.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Information of a table.
		Items []*string `json:"Items,omitempty" name:"Items" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTablesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTagsOfInstanceIdsRequest struct {
	*tchttp.BaseRequest

	// List of instances.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// Pagination offset.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTagsOfInstanceIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTagsOfInstanceIdsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTagsOfInstanceIdsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Pagination offset.
		Offset *int64 `json:"Offset,omitempty" name:"Offset"`

		// Number of entries per page.
		Limit *int64 `json:"Limit,omitempty" name:"Limit"`

		// Instance tag information.
		Rows []*TagsInfoOfInstance `json:"Rows,omitempty" name:"Rows" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTagsOfInstanceIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTagsOfInstanceIdsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTasksRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://cloud.tencent.com/document/api/236/15872) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// ID of an async task request, i.e., `AsyncRequestId` returned by relevant TencentDB operations.
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// Task type. If no value is passed in, all task types will be queried. Valid values:
	// 1 - rolling back a database;
	// 2 - performing an SQL operation;
	// 3 - importing data;
	// 5 - setting a parameter;
	// 6 - initializing a TencentDB instance;
	// 7 - restarting a TencentDB instance;
	// 8 - enabling GTID of a TencentDB instance;
	// 9 - upgrading a read-only instance;
	// 10 - rolling back databases in batches;
	// 11 - upgrading a master instance;
	// 12 - deleting a TencentDB table;
	// 13 - promoting a disaster recovery instance.
	TaskTypes []*int64 `json:"TaskTypes,omitempty" name:"TaskTypes" list`

	// Task status. If no value is passed in, all task statuses will be queried. Valid values:
	// -1 - undefined;
	// 0 - initializing;
	// 1 - running;
	// 2 - succeeded;
	// 3 - failed;
	// 4 - terminated;
	// 5 - deleted;
	// 6 - paused.
	TaskStatus []*int64 `json:"TaskStatus,omitempty" name:"TaskStatus" list`

	// Start time of the first task in the format of yyyy-MM-dd HH:mm:ss, such as 2017-12-31 10:40:01. It is used for queries by time range.
	StartTimeBegin *string `json:"StartTimeBegin,omitempty" name:"StartTimeBegin"`

	// End time of the last task in the format of yyyy-MM-dd HH:mm:ss, such as 2017-12-31 10:40:01. It is used for queries by time range.
	StartTimeEnd *string `json:"StartTimeEnd,omitempty" name:"StartTimeEnd"`

	// Record offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results to be returned for a single request. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTasksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible instances.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Information of an instance task.
		Items []*TaskDetail `json:"Items,omitempty" name:"Items" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTasksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTimeWindowRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeTimeWindowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTimeWindowRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTimeWindowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of maintenance time windows on Monday.
		Monday []*string `json:"Monday,omitempty" name:"Monday" list`

		// List of maintenance time windows on Tuesday.
		Tuesday []*string `json:"Tuesday,omitempty" name:"Tuesday" list`

		// List of maintenance time windows on Wednesday.
		Wednesday []*string `json:"Wednesday,omitempty" name:"Wednesday" list`

		// List of maintenance time windows on Thursday.
		Thursday []*string `json:"Thursday,omitempty" name:"Thursday" list`

		// List of maintenance time windows on Friday.
		Friday []*string `json:"Friday,omitempty" name:"Friday" list`

		// List of maintenance time windows on Saturday.
		Saturday []*string `json:"Saturday,omitempty" name:"Saturday" list`

		// List of maintenance time windows on Sunday.
		Sunday []*string `json:"Sunday,omitempty" name:"Sunday" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTimeWindowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTimeWindowResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUploadedFilesRequest struct {
	*tchttp.BaseRequest

	// File path. `OwnerUin` information of the root account should be entered in this field.
	Path *string `json:"Path,omitempty" name:"Path"`

	// Record offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results to be returned for a single request. Default value: 20.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeUploadedFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUploadedFilesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUploadedFilesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible SQL files.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of returned SQL files.
		Items []*SqlFileInfo `json:"Items,omitempty" name:"Items" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUploadedFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUploadedFilesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeviceCpuInfo struct {

	// Average instance CPU utilization
	Rate []*DeviceCpuRateInfo `json:"Rate,omitempty" name:"Rate" list`

	// CPU monitoring data of the instance
	Load []*int64 `json:"Load,omitempty" name:"Load" list`
}

type DeviceCpuRateInfo struct {

	// CPU core number
	CpuCore *int64 `json:"CpuCore,omitempty" name:"CpuCore"`

	// CPU utilization
	Rate []*int64 `json:"Rate,omitempty" name:"Rate" list`
}

type DeviceDiskInfo struct {

	// Time percentage of IO operations per second
	IoRatioPerSec []*int64 `json:"IoRatioPerSec,omitempty" name:"IoRatioPerSec" list`

	// Average wait time of device I/O operations * 100 in milliseconds. For example, if the value is 201, the average wait time of I/O operations is 201/100 = 2.1 milliseconds.
	IoWaitTime []*int64 `json:"IoWaitTime,omitempty" name:"IoWaitTime" list`

	// Average number of read operations completed by the disk per second * 100. For example, if the value is 2,002, the average number of read operations completed by the disk per second is 2,002/100=20.2.
	Read []*int64 `json:"Read,omitempty" name:"Read" list`

	// Average number of write operations completed by the disk per second * 100. For example, if the value is 30,001, the average number of write operations completed by the disk per second is 30,001/100=300.01.
	Write []*int64 `json:"Write,omitempty" name:"Write" list`
}

type DeviceMemInfo struct {

	// Total memory size in KB, which is the value of `total` in the `Mem:` in the `free` command
	Total []*int64 `json:"Total,omitempty" name:"Total" list`

	// Used memory size in KB, which is the value of `used` in the `Mem:` row in the `free` command
	Used []*int64 `json:"Used,omitempty" name:"Used" list`
}

type DeviceNetInfo struct {

	// Number of TCP connections
	Conn []*int64 `json:"Conn,omitempty" name:"Conn" list`

	// ENI inbound packets per second
	PackageIn []*int64 `json:"PackageIn,omitempty" name:"PackageIn" list`

	// ENI outbound packets per second
	PackageOut []*int64 `json:"PackageOut,omitempty" name:"PackageOut" list`

	// Inbound traffic in Kbps
	FlowIn []*int64 `json:"FlowIn,omitempty" name:"FlowIn" list`

	// Outbound traffic in Kbps
	FlowOut []*int64 `json:"FlowOut,omitempty" name:"FlowOut" list`
}

type DisassociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// Security group ID.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// List of instance IDs, which is an array of one or more instance IDs.
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

type DrInfo struct {

	// Disaster recovery instance status
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// AZ information
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Region information
	Region *string `json:"Region,omitempty" name:"Region"`

	// Instance sync status. Possible returned values include:
	// 0 - disaster recovery not synced;
	// 1 - disaster recovery syncing;
	// 2 - disaster recovery synced successfully;
	// 3 - disaster recovery sync failed;
	// 4 - repairing disaster recovery sync;
	SyncStatus *int64 `json:"SyncStatus,omitempty" name:"SyncStatus"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Instance type
	InstanceType *int64 `json:"InstanceType,omitempty" name:"InstanceType"`
}

type ErrlogItem struct {

	// Error occurrence time.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// Error details
	// Note: this field may return null, indicating that no valid values can be obtained.
	Content *string `json:"Content,omitempty" name:"Content"`
}

type ImportRecord struct {

	// Status value
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Status value
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// Execution duration
	CostTime *int64 `json:"CostTime,omitempty" name:"CostTime"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Backend task ID
	WorkId *string `json:"WorkId,omitempty" name:"WorkId"`

	// Name of the file to be imported
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// Execution progress
	Process *int64 `json:"Process,omitempty" name:"Process"`

	// Task creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// File size
	FileSize *string `json:"FileSize,omitempty" name:"FileSize"`

	// Task execution information
	Message *string `json:"Message,omitempty" name:"Message"`

	// Task ID
	JobId *int64 `json:"JobId,omitempty" name:"JobId"`

	// Name of the table to be imported
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// Async task request ID
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`
}

type Inbound struct {

	// Policy, which can be ACCEPT or DROP
	Action *string `json:"Action,omitempty" name:"Action"`

	// Source IP or IP range, such as 192.168.0.0/16
	CidrIp *string `json:"CidrIp,omitempty" name:"CidrIp"`

	// Port
	PortRange *string `json:"PortRange,omitempty" name:"PortRange"`

	// Network protocol. UDP and TCP are supported.
	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`

	// The direction of the rule, which is INPUT for inbound rules
	Dir *string `json:"Dir,omitempty" name:"Dir"`
}

type InitDBInstancesRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://cloud.tencent.com/document/api/236/15872) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// New password of the instance. Rule: It can only contain 8-64 characters and must contain at least two of the following types of characters: letters, digits, and special characters (!@#$%^*()).
	NewPassword *string `json:"NewPassword,omitempty" name:"NewPassword"`

	// List of instance parameters. Currently, "character_set_server" and "lower_case_table_names" are supported, whose value ranges are ["utf8","latin1","gbk","utf8mb4"] and ["0","1"], respectively.
	Parameters []*ParamInfo `json:"Parameters,omitempty" name:"Parameters" list`

	// Instance port. Value range: [1024, 65535].
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
}

func (r *InitDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InitDBInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InitDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Array of async task request IDs, which can be used to query the execution results of async tasks.
		AsyncRequestIds []*string `json:"AsyncRequestIds,omitempty" name:"AsyncRequestIds" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InitDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InitDBInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InstanceInfo struct {

	// Public network access status. Value range: 0 (not enabled), 1 (enabled), 2 (disabled)
	WanStatus *int64 `json:"WanStatus,omitempty" name:"WanStatus"`

	// AZ information
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Initialization flag. Value range: 0 (not initialized), 1 (initialized)
	InitFlag *int64 `json:"InitFlag,omitempty" name:"InitFlag"`

	// VIP information of a read-only instance. This field is exclusive to read-only instances where read-only access is enabled separately
	// Note: This field may return null, indicating that no valid values can be obtained.
	RoVipInfo *RoVipInfo `json:"RoVipInfo,omitempty" name:"RoVipInfo"`

	// Memory capacity in MB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Instance status. Value range: 0 (creating), 1 (running), 4 (isolating), 5 (isolated)
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// VPC ID, such as 51102
	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`

	// Information of a slave server
	// Note: This field may return null, indicating that no valid values can be obtained.
	SlaveInfo *SlaveInfo `json:"SlaveInfo,omitempty" name:"SlaveInfo"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Disk capacity in GB
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`

	// Auto-renewal flag. Value range: 0 (auto-renewal not enabled), 1 (auto-renewal enabled), 2 (auto-renewal disabled)
	AutoRenew *int64 `json:"AutoRenew,omitempty" name:"AutoRenew"`

	// Data replication mode. Valid values: 0 (async), 1 (semi-sync), 2 (strong sync)
	ProtectMode *int64 `json:"ProtectMode,omitempty" name:"ProtectMode"`

	// Details of a read-only group
	// Note: This field may return null, indicating that no valid values can be obtained.
	RoGroups []*RoGroup `json:"RoGroups,omitempty" name:"RoGroups" list`

	// Subnet ID, such as 2333
	SubnetId *int64 `json:"SubnetId,omitempty" name:"SubnetId"`

	// Instance type. Value range: 1 (master), 2 (disaster recovery), 3 (read-only)
	InstanceType *int64 `json:"InstanceType,omitempty" name:"InstanceType"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Region information
	Region *string `json:"Region,omitempty" name:"Region"`

	// Instance expiration time
	DeadlineTime *string `json:"DeadlineTime,omitempty" name:"DeadlineTime"`

	// AZ deployment mode. Valid values: 0 (single-AZ), 1 (multi-AZ)
	DeployMode *int64 `json:"DeployMode,omitempty" name:"DeployMode"`

	// Instance task status. 0 - no task; 1 - upgrading; 2 - importing data; 3 - activating slave; 4 - enabling public network access; 5 - batch operation in progress; 6 - rolling back; 7 - disabling public network access; 8 - changing password; 9 - renaming instance; 10 - restarting; 12 - migrating self-built instance; 13 - dropping table; 14 - creating and syncing disaster recovery instance; 15 - pending upgrade and switch; 16 - upgrade and switch in progress; 17 - upgrade and switch completed
	TaskStatus *int64 `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// Details of a master instance
	// Note: This field may return null, indicating that no valid values can be obtained.
	MasterInfo *MasterInfo `json:"MasterInfo,omitempty" name:"MasterInfo"`

	// Instance type. Value range: HA (High-Availability Edition), FE (Finance Edition), BASIC (Basic Edition)
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// Kernel version
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Details of a disaster recovery instance
	// Note: This field may return null, indicating that no valid values can be obtained.
	DrInfo []*DrInfo `json:"DrInfo,omitempty" name:"DrInfo" list`

	// Public domain name
	WanDomain *string `json:"WanDomain,omitempty" name:"WanDomain"`

	// Public network port number
	WanPort *int64 `json:"WanPort,omitempty" name:"WanPort"`

	// Billing type
	PayType *int64 `json:"PayType,omitempty" name:"PayType"`

	// Instance creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Instance IP
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Port number
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// Lock flag
	CdbError *int64 `json:"CdbError,omitempty" name:"CdbError"`

	// VPC descriptor, such as "vpc-5v8wn9mg"
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// Subnet descriptor, such as "subnet-1typ0s7d"
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// Physical ID
	PhysicalId *string `json:"PhysicalId,omitempty" name:"PhysicalId"`

	// Number of cores
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// Queries per second
	Qps *int64 `json:"Qps,omitempty" name:"Qps"`

	// AZ name
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// Physical machine model
	// Note: This field may return null, indicating that no valid values can be obtained.
	DeviceClass *string `json:"DeviceClass,omitempty" name:"DeviceClass"`

	// Placement group ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	DeployGroupId *string `json:"DeployGroupId,omitempty" name:"DeployGroupId"`

	// AZ ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
}

type InstanceRebootTime struct {

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Estimated restart time
	TimeInSeconds *int64 `json:"TimeInSeconds,omitempty" name:"TimeInSeconds"`
}

type InstanceRollbackRangeTime struct {

	// Queries database error code
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// Queries database error message
	Message *string `json:"Message,omitempty" name:"Message"`

	// List of instance IDs. An instance ID is in the format of cdb-c1nl9rpv, which is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Time range available for rollback
	Times []*RollbackTimeRange `json:"Times,omitempty" name:"Times" list`
}

type IsolateDBInstanceRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://cloud.tencent.com/document/api/236/15872) to query the ID, whose value is the `InstanceId` value in output parameters.
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

		// Async task request ID, which can be used to query the execution result of an async task. (This returned field has been disused. You can query the isolation status of an instance through the `DescribeDBInstances` API.)
	// Note: this field may return null, indicating that no valid values can be obtained.
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

type MasterInfo struct {

	// Region information
	Region *string `json:"Region,omitempty" name:"Region"`

	// Region ID
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// AZ ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// AZ information
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Long instance ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// Instance status
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Instance type
	InstanceType *int64 `json:"InstanceType,omitempty" name:"InstanceType"`

	// Task status
	TaskStatus *int64 `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// Memory capacity
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Disk capacity
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`

	// Instance model
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// Queries per second
	Qps *int64 `json:"Qps,omitempty" name:"Qps"`

	// VPC ID
	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *int64 `json:"SubnetId,omitempty" name:"SubnetId"`

	// Dedicated cluster ID
	ExClusterId *string `json:"ExClusterId,omitempty" name:"ExClusterId"`

	// Dedicated cluster name
	ExClusterName *string `json:"ExClusterName,omitempty" name:"ExClusterName"`
}

type ModifyAccountDescriptionRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// TencentDB account
	Accounts []*Account `json:"Accounts,omitempty" name:"Accounts" list`

	// Database account remarks
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyAccountDescriptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAccountDescriptionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAccountDescriptionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task request ID, which can be used to query the execution result of an async task.
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccountDescriptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAccountDescriptionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAccountPasswordRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// New password of the database account. It can only contain 8-64 characters and must contain at least two of the following types of characters: letters, digits, and special characters (_+-&=!@#$%^*()).
	NewPassword *string `json:"NewPassword,omitempty" name:"NewPassword"`

	// TencentDB account
	Accounts []*Account `json:"Accounts,omitempty" name:"Accounts" list`
}

func (r *ModifyAccountPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAccountPasswordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAccountPasswordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task request ID, which can be used to query the execution result of an async task.
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccountPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAccountPasswordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAccountPrivilegesRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Database account, including username and domain name.
	Accounts []*Account `json:"Accounts,omitempty" name:"Accounts" list`

	// Global permission. Valid values: "SELECT", "INSERT", "UPDATE", "DELETE", "CREATE", "PROCESS", "DROP", "REFERENCES", "INDEX", "ALTER", "SHOW DATABASES", "CREATE TEMPORARY TABLES", "LOCK TABLES", "EXECUTE", "CREATE VIEW", "SHOW VIEW", "CREATE ROUTINE", "ALTER ROUTINE", "EVENT", "TRIGGER".
	// Note: if this parameter is not passed in, it means to clear the permission.
	GlobalPrivileges []*string `json:"GlobalPrivileges,omitempty" name:"GlobalPrivileges" list`

	// Database permission. Valid values: "SELECT", "INSERT", "UPDATE", "DELETE", "CREATE", 	"DROP", "REFERENCES", "INDEX", "ALTER", "CREATE TEMPORARY TABLES", "LOCK TABLES", "EXECUTE", "CREATE VIEW", "SHOW VIEW", "CREATE ROUTINE", "ALTER ROUTINE", "EVENT", "TRIGGER".
	// Note: if this parameter is not passed in, it means to clear the permission.
	DatabasePrivileges []*DatabasePrivilege `json:"DatabasePrivileges,omitempty" name:"DatabasePrivileges" list`

	// Table permission in the database. Valid values: "SELECT", "INSERT", "UPDATE", "DELETE", "CREATE", 	"DROP", "REFERENCES", "INDEX", "ALTER", "CREATE VIEW", "SHOW VIEW", "TRIGGER".
	// Note: if this parameter is not passed in, it means to clear the permission.
	TablePrivileges []*TablePrivilege `json:"TablePrivileges,omitempty" name:"TablePrivileges" list`

	// Column permission in table. Valid values: "SELECT", "INSERT", "UPDATE", "REFERENCES".
	// Note: if this parameter is not passed in, it means to clear the permission.
	ColumnPrivileges []*ColumnPrivilege `json:"ColumnPrivileges,omitempty" name:"ColumnPrivileges" list`
}

func (r *ModifyAccountPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAccountPrivilegesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAccountPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task request ID, which can be used to query the execution result of an async task.
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccountPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAccountPrivilegesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAutoRenewFlagRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// Auto-renewal flag. Value range: 0 (auto-renewal not enabled), 1 (auto-renewal enabled).
	AutoRenew *int64 `json:"AutoRenew,omitempty" name:"AutoRenew"`
}

func (r *ModifyAutoRenewFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAutoRenewFlagRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAutoRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAutoRenewFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAutoRenewFlagResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyBackupConfigRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Backup file retention period in days. Value range: 7-732.
	ExpireDays *int64 `json:"ExpireDays,omitempty" name:"ExpireDays"`

	// (This parameter will be disused. The `BackupTimeWindow` parameter is recommended.) Backup time range in the format of 02:00–06:00, with the start time and end time on the hour. Valid values: 00:00–12:00, 02:00–06:00, 06:00–10:00, 10:00–14:00, 14:00–18:00, 18:00–22:00, 22:00–02:00.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Automatic backup mode. Only `physical` (physical cold backup) is supported
	BackupMethod *string `json:"BackupMethod,omitempty" name:"BackupMethod"`

	// Binlog retention period in days. Value range: 7-732. It cannot be greater than the retention period of backup files.
	BinlogExpireDays *int64 `json:"BinlogExpireDays,omitempty" name:"BinlogExpireDays"`

	// Backup time window; for example, to set up backup between 10:00 and 14:00 on every Tuesday and Sunday, you should set this parameter as follows: {"Monday": "", "Tuesday": "10:00-14:00", "Wednesday": "", "Thursday": "", "Friday": "", "Saturday": "", "Sunday": "10:00-14:00"} (Note: You can set up backup on different days, but the backup time windows need to be the same. If this field is set, the `StartTime` field will be ignored)
	BackupTimeWindow *CommonTimeWindow `json:"BackupTimeWindow,omitempty" name:"BackupTimeWindow"`
}

func (r *ModifyBackupConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyBackupConfigRequest) FromJsonString(s string) error {
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

func (r *ModifyBackupConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceNameRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://cloud.tencent.com/document/api/236/15872) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance name.
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

	// Array of instance IDs in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://cloud.tencent.com/document/api/236/15872) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// Project ID.
	NewProjectId *int64 `json:"NewProjectId,omitempty" name:"NewProjectId"`
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

type ModifyDBInstanceSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// List of IDs of security groups to be modified, which is an array of one or more security group IDs.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`
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

type ModifyDBInstanceVipVportRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://cloud.tencent.com/document/api/236/15872) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Destination IP. Either this parameter or `DstPort` must be passed in.
	DstIp *string `json:"DstIp,omitempty" name:"DstIp"`

	// Destination port number. Value range: [1024-65535]. Either this parameter or `DstIp` must be passed in.
	DstPort *int64 `json:"DstPort,omitempty" name:"DstPort"`

	// Unified VPC ID
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// Unified subnet ID.
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// Repossession duration in hours for old IP in the original network when changing from the basic network to VPC or changing the VPC subnet. Value range: 0–168 hours. Default value: 24 hours.
	ReleaseDuration *int64 `json:"ReleaseDuration,omitempty" name:"ReleaseDuration"`
}

func (r *ModifyDBInstanceVipVportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBInstanceVipVportRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceVipVportResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task ID. (This returned field has been disused)
	// Note: this field may return null, indicating that no valid values can be obtained.
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBInstanceVipVportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBInstanceVipVportResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceParamRequest struct {
	*tchttp.BaseRequest

	// List of short instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// List of parameters to be modified. Every element is a combination of `Name` (parameter name) and `CurrentValue` (new value).
	ParamList []*Parameter `json:"ParamList,omitempty" name:"ParamList" list`
}

func (r *ModifyInstanceParamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstanceParamRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceParamResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task ID, which can be used to query task progress.
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceParamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstanceParamResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceTagRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Tag to be added or modified.
	ReplaceTags []*TagInfo `json:"ReplaceTags,omitempty" name:"ReplaceTags" list`

	// Tag to be deleted.
	DeleteTags []*TagInfo `json:"DeleteTags,omitempty" name:"DeleteTags" list`
}

func (r *ModifyInstanceTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstanceTagRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceTagResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstanceTagResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNameOrDescByDpIdRequest struct {
	*tchttp.BaseRequest

	// ID of a placement group.
	DeployGroupId *string `json:"DeployGroupId,omitempty" name:"DeployGroupId"`

	// Name of a placement group, which can contain up to 60 characters. The placement group name and description cannot both be empty.
	DeployGroupName *string `json:"DeployGroupName,omitempty" name:"DeployGroupName"`

	// Description of a placement group, which can contain up to 200 characters. The placement group name and description cannot both be empty.
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyNameOrDescByDpIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNameOrDescByDpIdRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNameOrDescByDpIdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNameOrDescByDpIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNameOrDescByDpIdResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyParamTemplateRequest struct {
	*tchttp.BaseRequest

	// Template ID.
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// Template name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Template description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// List of parameters.
	ParamList []*Parameter `json:"ParamList,omitempty" name:"ParamList" list`
}

func (r *ModifyParamTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyParamTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyParamTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyParamTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyParamTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyRoGroupInfoRequest struct {
	*tchttp.BaseRequest

	// RO group ID.
	RoGroupId *string `json:"RoGroupId,omitempty" name:"RoGroupId"`

	// RO group details.
	RoGroupInfo *RoGroupAttr `json:"RoGroupInfo,omitempty" name:"RoGroupInfo"`

	// Weights of instances in RO group. If the weighting mode of an RO group is changed to custom mode, this parameter must be set, and a weight value needs to be set for each RO instance.
	RoWeightValues []*RoWeightValue `json:"RoWeightValues,omitempty" name:"RoWeightValues" list`

	// Whether to rebalance the loads of RO instances in the RO group. Supported values include `1` (yes) and `0` (no). The default value is `0`. Please note that if this value is set to `1`, connections to the RO instances in the RO group will be interrupted transiently; therefore, you should ensure that your application can reconnect to the databases.
	IsBalanceRoLoad *int64 `json:"IsBalanceRoLoad,omitempty" name:"IsBalanceRoLoad"`
}

func (r *ModifyRoGroupInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyRoGroupInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyRoGroupInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRoGroupInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyRoGroupInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTimeWindowRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Time period available for maintenance after modification in the format of 10:00-12:00. Each period lasts from half an hour to three hours, with the start time and end time aligned by half-hour. Up to two time periods can be set. Start and end time range: [00:00, 24:00].
	TimeRanges []*string `json:"TimeRanges,omitempty" name:"TimeRanges" list`

	// Specifies for which day to modify the time period. Value range: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday. If it is not specified or is left blank, the time period will be modified for every day by default.
	Weekdays []*string `json:"Weekdays,omitempty" name:"Weekdays" list`
}

func (r *ModifyTimeWindowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTimeWindowRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTimeWindowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTimeWindowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTimeWindowResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OfflineIsolatedInstancesRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *OfflineIsolatedInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OfflineIsolatedInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OfflineIsolatedInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OfflineIsolatedInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OfflineIsolatedInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OpenDBInstanceGTIDRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *OpenDBInstanceGTIDRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OpenDBInstanceGTIDRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OpenDBInstanceGTIDResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task request ID, which can be used to query the execution result of an async task.
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenDBInstanceGTIDResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OpenDBInstanceGTIDResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OpenWanServiceRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://cloud.tencent.com/document/api/236/15872) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *OpenWanServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OpenWanServiceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OpenWanServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task request ID, which can be used to query the execution result of an async task.
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenWanServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OpenWanServiceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Outbound struct {

	// Policy, which can be ACCEPT or DROP
	Action *string `json:"Action,omitempty" name:"Action"`

	// Destination IP or IP range, such as 172.16.0.0/12
	CidrIp *string `json:"CidrIp,omitempty" name:"CidrIp"`

	// Port or port range
	PortRange *string `json:"PortRange,omitempty" name:"PortRange"`

	// Network protocol. UDP and TCP are supported
	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`

	// The direction of the rule, which is OUTPUT for inbound rules
	Dir *string `json:"Dir,omitempty" name:"Dir"`
}

type ParamInfo struct {

	// Parameter name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Parameter value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type ParamRecord struct {

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Parameter name
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`

	// Parameter value before modification
	OldValue *string `json:"OldValue,omitempty" name:"OldValue"`

	// Parameter value after modification
	NewValue *string `json:"NewValue,omitempty" name:"NewValue"`

	// Whether the parameter is modified successfully
	IsSucess *bool `json:"IsSucess,omitempty" name:"IsSucess"`

	// Modification time
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type ParamTemplateInfo struct {

	// Parameter template ID
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// Parameter template name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Parameter template description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Instance engine version
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`
}

type Parameter struct {

	// Parameter name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Parameter value
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`
}

type ParameterDetail struct {

	// Parameter name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Parameter type
	ParamType *string `json:"ParamType,omitempty" name:"ParamType"`

	// Default value of the parameter
	Default *string `json:"Default,omitempty" name:"Default"`

	// Parameter description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Current value of the parameter
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`

	// Whether the database needs to be restarted for the modified parameter to take effect. Value range: 0 (no); 1 (yes)
	NeedReboot *int64 `json:"NeedReboot,omitempty" name:"NeedReboot"`

	// Maximum value of the parameter
	Max *int64 `json:"Max,omitempty" name:"Max"`

	// Minimum value of the parameter
	Min *int64 `json:"Min,omitempty" name:"Min"`

	// Enumerated values of the parameter. It is null if the parameter is non-enumerated
	EnumValue []*string `json:"EnumValue,omitempty" name:"EnumValue" list`
}

type RegionSellConf struct {

	// Region name
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// Area
	Area *string `json:"Area,omitempty" name:"Area"`

	// Whether it is a default region
	IsDefaultRegion *int64 `json:"IsDefaultRegion,omitempty" name:"IsDefaultRegion"`

	// Region name
	Region *string `json:"Region,omitempty" name:"Region"`

	// Sale configuration of the AZ
	ZonesConf []*ZoneSellConf `json:"ZonesConf,omitempty" name:"ZonesConf" list`
}

type ReleaseIsolatedDBInstancesRequest struct {
	*tchttp.BaseRequest

	// Array of instance IDs in the format of `cdb-c1nl9rpv`. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [DescribeDBInstances](https://cloud.tencent.com/document/api/236/15872) API to query the ID, whose value is the `InstanceId` value in the output parameters.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *ReleaseIsolatedDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReleaseIsolatedDBInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReleaseIsolatedDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Deisolation result set.
		Items []*ReleaseResult `json:"Items,omitempty" name:"Items" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReleaseIsolatedDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReleaseIsolatedDBInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReleaseResult struct {

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Result value of instance deisolation. A returned value of 0 indicates success.
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// Error message for instance deisolation.
	Message *string `json:"Message,omitempty" name:"Message"`
}

type RestartDBInstancesRequest struct {
	*tchttp.BaseRequest

	// Array of instance IDs in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *RestartDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RestartDBInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RestartDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task request ID, which can be used to query the execution result of an async task.
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RestartDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RestartDBInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RoGroup struct {

	// Read-only group mode. Valid values: `alone` (the system assigns a read-only group automatically), `allinone` (a new read-only group will be created), `join` (an existing read-only group will be used).
	RoGroupMode *string `json:"RoGroupMode,omitempty" name:"RoGroupMode"`

	// Read-only group ID.
	RoGroupId *string `json:"RoGroupId,omitempty" name:"RoGroupId"`

	// Read-only group name.
	RoGroupName *string `json:"RoGroupName,omitempty" name:"RoGroupName"`

	// Whether to enable the function of isolating an instance that exceeds the latency threshold. If it is enabled, when the latency between the read-only instance and the master instance exceeds the latency threshold, the read-only instance will be isolated. Valid values: 1 (enabled), 0 (not enabled)
	RoOfflineDelay *int64 `json:"RoOfflineDelay,omitempty" name:"RoOfflineDelay"`

	// Latency threshold
	RoMaxDelayTime *int64 `json:"RoMaxDelayTime,omitempty" name:"RoMaxDelayTime"`

	// Minimum number of instances to be retained. If the number of the purchased read-only instances is smaller than the set value, they will not be removed.
	MinRoInGroup *int64 `json:"MinRoInGroup,omitempty" name:"MinRoInGroup"`

	// Read/write weight distribution mode. Valid values: `system` (weights are assigned by the system automatically), `custom` (weights are customized)
	WeightMode *string `json:"WeightMode,omitempty" name:"WeightMode"`

	// Weight value.
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// Details of read-only instances in read-only group
	RoInstances []*RoInstanceInfo `json:"RoInstances,omitempty" name:"RoInstances" list`

	// Private IP of read-only group.
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Private network port number of read-only group.
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// VPC ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// Subnet ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// Read-only group region.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RoGroupRegion *string `json:"RoGroupRegion,omitempty" name:"RoGroupRegion"`

	// Read-only group AZ.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RoGroupZone *string `json:"RoGroupZone,omitempty" name:"RoGroupZone"`
}

type RoGroupAttr struct {

	// RO group name.
	RoGroupName *string `json:"RoGroupName,omitempty" name:"RoGroupName"`

	// Maximum delay threshold for RO instances in seconds. Minimum value: 1. Please note that this value will take effect only if an instance removal policy is enabled in the RO group.
	RoMaxDelayTime *int64 `json:"RoMaxDelayTime,omitempty" name:"RoMaxDelayTime"`

	// Whether to enable instance removal. Valid values: 1 (enabled), 0 (not enabled). Please note that if instance removal is enabled, the delay threshold parameter (`RoMaxDelayTime`) must be set.
	RoOfflineDelay *int64 `json:"RoOfflineDelay,omitempty" name:"RoOfflineDelay"`

	// Minimum number of instances to be retained, which can be set to any value less than or equal to the number of RO instances in the RO group. Please note that if this value is set to be greater than the number of RO instances, no removal will be performed, and if it is set to 0, all instances with an excessive delay will be removed.
	MinRoInGroup *int64 `json:"MinRoInGroup,omitempty" name:"MinRoInGroup"`

	// Weighting mode. Supported values include `system` (automatically assigned by the system) and `custom` (defined by user). Please note that if the `custom` mode is selected, the RO instance weight configuration parameter (RoWeightValues) must be set.
	WeightMode *string `json:"WeightMode,omitempty" name:"WeightMode"`
}

type RoInstanceInfo struct {

	// Master instance ID corresponding to the RO group
	MasterInstanceId *string `json:"MasterInstanceId,omitempty" name:"MasterInstanceId"`

	// RO instance status in the RO group. Value range: online, offline
	RoStatus *string `json:"RoStatus,omitempty" name:"RoStatus"`

	// Last deactivation time of a RO instance in the RO group
	OfflineTime *string `json:"OfflineTime,omitempty" name:"OfflineTime"`

	// RO instance weight in the RO group
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// RO instance region name, such as ap-shanghai
	Region *string `json:"Region,omitempty" name:"Region"`

	// Name of RO AZ, such as ap-shanghai-1
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// RO instance ID in the format of cdbro-c1nl9rpv
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// RO instance status. Value range: 0 (creating), 1 (running), 4 (deleting)
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Instance type. Value range: 1 (master), 2 (disaster recovery), 3 (read-only)
	InstanceType *int64 `json:"InstanceType,omitempty" name:"InstanceType"`

	// RO instance name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Pay-as-you-go billing status. Value range: 1 (normal), 2 (in arrears)
	HourFeeStatus *int64 `json:"HourFeeStatus,omitempty" name:"HourFeeStatus"`

	// RO instance task status. Value range: <br>0 - no task <br>1 - upgrading <br>2 - importing data <br>3 - activating slave <br>4 - public network access enabled <br>5 - batch operation in progress <br>6 - rolling back <br>7 - public network access not enabled <br>8 - modifying password <br>9 - renaming instance <br>10 - restarting <br>12 - migrating self-built instance <br>13 - dropping table <br>14 - creating and syncing disaster recovery instance
	TaskStatus *int64 `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// RO instance memory size in MB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// RO instance disk size in GB
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`

	// Queries per second
	Qps *int64 `json:"Qps,omitempty" name:"Qps"`

	// Private IP address of the RO instance
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Access port of the RO instance
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// VPC ID of the RO instance
	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`

	// VPC subnet ID of the RO instance
	SubnetId *int64 `json:"SubnetId,omitempty" name:"SubnetId"`

	// RO instance specification description. Value range: CUSTOM
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// Database engine version of the RO instance. Value range: 5.1, 5.5, 5.6, 5.7
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// RO instance expiration time in the format of yyyy-mm-dd hh:mm:ss. If it is a pay-as-you-go instance, the value of this field is 0000-00-00 00:00:00
	DeadlineTime *string `json:"DeadlineTime,omitempty" name:"DeadlineTime"`

	// RO instance billing method. Value range: 0 (monthly subscribed), 1 (pay-as-you-go), 2 (monthly postpaid)
	PayType *int64 `json:"PayType,omitempty" name:"PayType"`
}

type RoVipInfo struct {

	// VIP status of the read-only instance
	RoVipStatus *int64 `json:"RoVipStatus,omitempty" name:"RoVipStatus"`

	// VPC subnet of the read-only instance
	RoSubnetId *int64 `json:"RoSubnetId,omitempty" name:"RoSubnetId"`

	// VPC of the read-only instance
	RoVpcId *int64 `json:"RoVpcId,omitempty" name:"RoVpcId"`

	// VIP port number of the read-only instance
	RoVport *int64 `json:"RoVport,omitempty" name:"RoVport"`

	// VIP of the read-only instance
	RoVip *string `json:"RoVip,omitempty" name:"RoVip"`
}

type RoWeightValue struct {

	// RO instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Weight value. Value range: [0, 100].
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
}

type RollbackDBName struct {

	// Original database name before rollback
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// New database name after rollback
	NewDatabaseName *string `json:"NewDatabaseName,omitempty" name:"NewDatabaseName"`
}

type RollbackInstancesInfo struct {

	// TencentDB instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Rollback policy. Value range: table, db, full. Default value: full. Table: expedited rollback mode, where only the selected table-level backups and binlogs are imported; for cross-table rollback, if the associated tables are not selected simultaneously, the rollback will fail; the parameter `Databases` must be empty under this mode. db: fast rollback mode, where only the selected database-level backups and binlogs are imported; for cross-database rollback, if the associated databases are not selected simultaneously, the rollback will fail. full: ordinary rollback mode, which imports all the backups and binlogs of the instance at a relatively low speed.
	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`

	// Database rollback time in the format of yyyy-mm-dd hh:mm:ss
	RollbackTime *string `json:"RollbackTime,omitempty" name:"RollbackTime"`

	// Information of the databases to be rolled back, which means rollback at the database level
	Databases []*RollbackDBName `json:"Databases,omitempty" name:"Databases" list`

	// Information of the tables to be rolled back, which means rollback at the table level
	Tables []*RollbackTables `json:"Tables,omitempty" name:"Tables" list`
}

type RollbackTableName struct {

	// Original table name before rollback
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// New table name after rollback
	NewTableName *string `json:"NewTableName,omitempty" name:"NewTableName"`
}

type RollbackTables struct {

	// Database name
	Database *string `json:"Database,omitempty" name:"Database"`

	// Table details
	Table []*RollbackTableName `json:"Table,omitempty" name:"Table" list`
}

type RollbackTimeRange struct {

	// Start time available for rollback in the format of yyyy-MM-dd HH:mm:ss, such as 2016-10-29 01:06:04
	Begin *string `json:"Begin,omitempty" name:"Begin"`

	// End time available for rollback in the format of yyyy-MM-dd HH:mm:ss, such as 2016-11-02 11:44:47
	End *string `json:"End,omitempty" name:"End"`
}

type SecurityGroup struct {

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Creation time in the format of yyyy-mm-dd hh:mm:ss
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Inbound rule
	Inbound []*Inbound `json:"Inbound,omitempty" name:"Inbound" list`

	// Outbound rule
	Outbound []*Outbound `json:"Outbound,omitempty" name:"Outbound" list`

	// Security group ID
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// Security group name
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`

	// Security group remarks
	SecurityGroupRemark *string `json:"SecurityGroupRemark,omitempty" name:"SecurityGroupRemark"`
}

type SellConfig struct {

	// Device type
	Device *string `json:"Device,omitempty" name:"Device"`

	// Purchasable specification description
	Type *string `json:"Type,omitempty" name:"Type"`

	// Instance type
	CdbType *string `json:"CdbType,omitempty" name:"CdbType"`

	// Memory size in MB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// CPU core count
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// Minimum disk size in GB
	VolumeMin *int64 `json:"VolumeMin,omitempty" name:"VolumeMin"`

	// Maximum disk size in GB
	VolumeMax *int64 `json:"VolumeMax,omitempty" name:"VolumeMax"`

	// Disk increment in GB
	VolumeStep *int64 `json:"VolumeStep,omitempty" name:"VolumeStep"`

	// Number of connections
	Connection *int64 `json:"Connection,omitempty" name:"Connection"`

	// Queries per second
	Qps *int64 `json:"Qps,omitempty" name:"Qps"`

	// IOs per second
	Iops *int64 `json:"Iops,omitempty" name:"Iops"`

	// Application scenario description
	Info *string `json:"Info,omitempty" name:"Info"`

	// Status value
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Tag value
	Tag *int64 `json:"Tag,omitempty" name:"Tag"`
}

type SellType struct {

	// Name of the purchasable instance
	TypeName *string `json:"TypeName,omitempty" name:"TypeName"`

	// Kernel version number
	EngineVersion []*string `json:"EngineVersion,omitempty" name:"EngineVersion" list`

	// Configuration details of a purchasable specification
	Configs []*SellConfig `json:"Configs,omitempty" name:"Configs" list`
}

type SlaveConfig struct {

	// Replication mode of the slave database. Value range: async, semi-sync
	ReplicationMode *string `json:"ReplicationMode,omitempty" name:"ReplicationMode"`

	// AZ name of the slave database, such as ap-shanghai-1
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type SlaveInfo struct {

	// Information of slave server 1
	First *SlaveInstanceInfo `json:"First,omitempty" name:"First"`

	// Information of slave server 2
	// Note: This field may return null, indicating that no valid values can be obtained.
	Second *SlaveInstanceInfo `json:"Second,omitempty" name:"Second"`
}

type SlaveInstanceInfo struct {

	// Port number
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// Region information
	Region *string `json:"Region,omitempty" name:"Region"`

	// Virtual IP information
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// AZ information
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type SlowLogInfo struct {

	// Backup filename
	Name *string `json:"Name,omitempty" name:"Name"`

	// Backup file size in bytes
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// Backup snapshot time in the format of yyyy-MM-dd HH:mm:ss, such as 2016-03-17 02:10:37
	Date *string `json:"Date,omitempty" name:"Date"`

	// Download address on the private network
	IntranetUrl *string `json:"IntranetUrl,omitempty" name:"IntranetUrl"`

	// Download address on the public network
	InternetUrl *string `json:"InternetUrl,omitempty" name:"InternetUrl"`

	// Log type. Value range: slowlog (slow log)
	Type *string `json:"Type,omitempty" name:"Type"`
}

type SlowLogItem struct {

	// SQL execution time.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// SQL execution duration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	QueryTime *float64 `json:"QueryTime,omitempty" name:"QueryTime"`

	// SQL statement.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SqlText *string `json:"SqlText,omitempty" name:"SqlText"`

	// Client address.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UserHost *string `json:"UserHost,omitempty" name:"UserHost"`

	// Username.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Database name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Database *string `json:"Database,omitempty" name:"Database"`

	// Lock duration.
	// Note: this field may return null, indicating that no valid values can be obtained.
	LockTime *float64 `json:"LockTime,omitempty" name:"LockTime"`

	// Number of scanned rows.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RowsExamined *int64 `json:"RowsExamined,omitempty" name:"RowsExamined"`

	// Number of rows in result set.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RowsSent *int64 `json:"RowsSent,omitempty" name:"RowsSent"`

	// SQL template.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SqlTemplate *string `json:"SqlTemplate,omitempty" name:"SqlTemplate"`

	// SQL statement MD5.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Md5 *string `json:"Md5,omitempty" name:"Md5"`
}

type SqlFileInfo struct {

	// Upload time
	UploadTime *string `json:"UploadTime,omitempty" name:"UploadTime"`

	// Upload progress
	UploadInfo *UploadInfo `json:"UploadInfo,omitempty" name:"UploadInfo"`

	// Filename
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// File size in bytes
	FileSize *int64 `json:"FileSize,omitempty" name:"FileSize"`

	// Whether upload is finished. Valid values: 0 (not completed), 1 (completed)
	IsUploadFinished *int64 `json:"IsUploadFinished,omitempty" name:"IsUploadFinished"`

	// File ID
	FileId *string `json:"FileId,omitempty" name:"FileId"`
}

type StartBatchRollbackRequest struct {
	*tchttp.BaseRequest

	// Details of the instance for rollback
	Instances []*RollbackInstancesInfo `json:"Instances,omitempty" name:"Instances" list`
}

func (r *StartBatchRollbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartBatchRollbackRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartBatchRollbackResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task request ID, which can be used to query the execution result of an async task.
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartBatchRollbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartBatchRollbackResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopDBImportJobRequest struct {
	*tchttp.BaseRequest

	// Async task request ID.
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`
}

func (r *StopDBImportJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopDBImportJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopDBImportJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopDBImportJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopDBImportJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SwitchForUpgradeRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *SwitchForUpgradeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SwitchForUpgradeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SwitchForUpgradeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchForUpgradeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SwitchForUpgradeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TableName struct {

	// Table name
	TableName *string `json:"TableName,omitempty" name:"TableName"`
}

type TablePrivilege struct {

	// Database name
	Database *string `json:"Database,omitempty" name:"Database"`

	// Table name
	Table *string `json:"Table,omitempty" name:"Table"`

	// Permission information
	Privileges []*string `json:"Privileges,omitempty" name:"Privileges" list`
}

type TagInfo struct {

	// Tag key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value
	TagValue []*string `json:"TagValue,omitempty" name:"TagValue" list`
}

type TagInfoUnit struct {

	// Tag key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type TagsInfoOfInstance struct {

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Tag information
	Tags []*TagInfoUnit `json:"Tags,omitempty" name:"Tags" list`
}

type TaskDetail struct {

	// Error code.
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// Error message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// ID of an instance task.
	JobId *int64 `json:"JobId,omitempty" name:"JobId"`

	// Instance task progress.
	Progress *int64 `json:"Progress,omitempty" name:"Progress"`

	// Instance task status. Valid values:
	// "UNDEFINED" - undefined;
	// "INITIAL" - initializing;
	// "RUNNING" - running;
	// "SUCCEED" - succeeded;
	// "FAILED" - failed;
	// "KILLED" - terminated;
	// "REMOVED" - deleted;
	// "PAUSED" - paused.
	TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// Instance task type. Valid values:
	// "ROLLBACK" - rolling back a database;
	// "SQL OPERATION" - performing an SQL operation;
	// "IMPORT DATA" - importing data;
	// "MODIFY PARAM" - setting a parameter;
	// "INITIAL" - initializing a TencentDB instance;
	// "REBOOT" - restarting a TencentDB instance;
	// "OPEN GTID" - enabling GTID of a TencentDB instance;
	// "UPGRADE RO" - upgrading a read-only instance;
	// "BATCH ROLLBACK" - rolling back databases in batches;
	// "UPGRADE MASTER" - upgrading a master instance;
	// "DROP TABLES" - dropping a TencentDB table;
	// "SWITCH DR TO MASTER" - promoting a disaster recovery instance.
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// Instance task start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Instance task end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// ID of an instance associated with a task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// Async task request ID.
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`
}

type UpgradeDBInstanceEngineVersionRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://cloud.tencent.com/document/api/236/15872) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Version of master instance database engine. Value range: 5.6, 5.7
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// Mode of switch to a new instance. Value range: 0 (switch immediately), 1 (switch within a time window). Default value: 0. If the value is 1, the switch process will be performed within a time window. Or, you can call the [switching to new instance API](https://cloud.tencent.com/document/product/236/15864) to trigger the process.
	WaitSwitch *int64 `json:"WaitSwitch,omitempty" name:"WaitSwitch"`
}

func (r *UpgradeDBInstanceEngineVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeDBInstanceEngineVersionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeDBInstanceEngineVersionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task ID. The task execution result can be queried using the [async task execution result querying API](https://cloud.tencent.com/document/api/236/20410).
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpgradeDBInstanceEngineVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeDBInstanceEngineVersionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeDBInstanceRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of `cdb-c1nl9rpv` or `cdbro-c1nl9rpv`. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [DescribeDBInstances](https://cloud.tencent.com/document/api/236/15872) API to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Memory size in MB after upgrade. To ensure that the `Memory` value to be passed in is valid, please use the [DescribeDBZoneConfig](https://cloud.tencent.com/document/product/236/17229) API to query the specifications of the memory that can be upgraded to.
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Disk size in GB after upgrade. To ensure that the `Volume` value to be passed in is valid, please use the [DescribeDBZoneConfig](https://cloud.tencent.com/document/product/236/17229) API to query the specifications of the disk that can be upgraded to.
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`

	// Data replication mode. Valid values: 0 (async), 1 (semi-sync), 2 (strong sync). This parameter can be specified when upgrading master instances and is meaningless for read-only or disaster recovery instances.
	ProtectMode *int64 `json:"ProtectMode,omitempty" name:"ProtectMode"`

	// Deployment mode. Valid values: 0 (single-AZ), 1 (multi-AZ). Default value: 0. This parameter can be specified when upgrading master instances and is meaningless for read-only or disaster recovery instances.
	DeployMode *int64 `json:"DeployMode,omitempty" name:"DeployMode"`

	// AZ information of slave database 1, which is the `Zone` value of the instance by default. This parameter can be specified when upgrading master instances in multi-AZ mode and is meaningless for read-only or disaster recovery instances. You can use the [DescribeDBZoneConfig](https://cloud.tencent.com/document/product/236/17229) API to query the supported AZs.
	SlaveZone *string `json:"SlaveZone,omitempty" name:"SlaveZone"`

	// Version of master instance database engine. Valid values: 5.5, 5.6, 5.7.
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// Mode of switch to new instance. Valid values: 0 (switch immediately), 1 (switch within a time window). Default value: 0. If the value is 1, the switch process will be performed within a time window. Or, you can call the [SwitchForUpgrade](https://cloud.tencent.com/document/product/236/15864) API to trigger the process.
	WaitSwitch *int64 `json:"WaitSwitch,omitempty" name:"WaitSwitch"`

	// AZ information of slave database 2, which is empty by default. This parameter can be specified when upgrading master instances and is meaningless for read-only or disaster recovery instances.
	BackupZone *string `json:"BackupZone,omitempty" name:"BackupZone"`

	// Instance type. Valid values: master (master instance), dr (disaster recovery instance), ro (read-only instance). Default value: master.
	InstanceRole *string `json:"InstanceRole,omitempty" name:"InstanceRole"`
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

		// Order ID.
		DealIds []*string `json:"DealIds,omitempty" name:"DealIds" list`

		// Async task request ID, which can be used to query the execution result of an async task.
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

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

type UploadInfo struct {

	// Number of parts of file
	AllSliceNum *int64 `json:"AllSliceNum,omitempty" name:"AllSliceNum"`

	// Number of completed parts
	CompleteNum *int64 `json:"CompleteNum,omitempty" name:"CompleteNum"`
}

type ZoneConf struct {

	// AZ deployment mode. Value range: 0 (single-AZ), 1 (multi-AZ)
	DeployMode []*int64 `json:"DeployMode,omitempty" name:"DeployMode" list`

	// AZ where the master instance is located
	MasterZone []*string `json:"MasterZone,omitempty" name:"MasterZone" list`

	// AZ where salve database 1 is located when the instance is deployed in multi-AZ mode
	SlaveZone []*string `json:"SlaveZone,omitempty" name:"SlaveZone" list`

	// AZ where salve database 2 is located when the instance is deployed in multi-AZ mode
	BackupZone []*string `json:"BackupZone,omitempty" name:"BackupZone" list`
}

type ZoneSellConf struct {

	// AZ status. Value range: 0 (not available), 1 (available), 2 (purchasable), 3 (not purchasable), 4 (not displayed)
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// AZ name
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// Whether it is a custom instance type
	IsCustom *bool `json:"IsCustom,omitempty" name:"IsCustom"`

	// Whether disaster recovery is supported
	IsSupportDr *bool `json:"IsSupportDr,omitempty" name:"IsSupportDr"`

	// Whether VPC is supported
	IsSupportVpc *bool `json:"IsSupportVpc,omitempty" name:"IsSupportVpc"`

	// Maximum purchasable quantity of hourly billed instances
	HourInstanceSaleMaxNum *int64 `json:"HourInstanceSaleMaxNum,omitempty" name:"HourInstanceSaleMaxNum"`

	// Whether it is a default AZ
	IsDefaultZone *bool `json:"IsDefaultZone,omitempty" name:"IsDefaultZone"`

	// Whether it is a BM zone
	IsBm *bool `json:"IsBm,omitempty" name:"IsBm"`

	// Supported billing method. Value range: 0 (monthly subscribed), 1 (hourly), 2 (postpaid)
	PayType []*string `json:"PayType,omitempty" name:"PayType" list`

	// Data replication type. Value range: 0 (async), 1 (semi-sync), 2 (strong sync)
	ProtectMode []*string `json:"ProtectMode,omitempty" name:"ProtectMode" list`

	// AZ name
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Array of purchasable instance types
	SellType []*SellType `json:"SellType,omitempty" name:"SellType" list`

	// Multi-AZ information
	ZoneConf *ZoneConf `json:"ZoneConf,omitempty" name:"ZoneConf"`

	// Information of the supported disaster recovery AZ
	DrZone []*string `json:"DrZone,omitempty" name:"DrZone" list`

	// Whether cross-AZ read-only access is supported
	IsSupportRemoteRo *bool `json:"IsSupportRemoteRo,omitempty" name:"IsSupportRemoteRo"`

	// Information of supported cross-AZ read-only zone
	// Note: this field may return null, indicating that no valid values can be obtained.
	RemoteRoZone []*string `json:"RemoteRoZone,omitempty" name:"RemoteRoZone" list`
}
