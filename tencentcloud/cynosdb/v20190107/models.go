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

type Ability struct {
	// Whether secondary AZ is supported
	IsSupportSlaveZone *string `json:"IsSupportSlaveZone,omitempty" name:"IsSupportSlaveZone"`

	// The reason why secondary AZ is not supported
	// Note: This field may return null, indicating that no valid values can be obtained.
	NonsupportSlaveZoneReason *string `json:"NonsupportSlaveZoneReason,omitempty" name:"NonsupportSlaveZoneReason"`

	// Whether read-only instance is supported
	IsSupportRo *string `json:"IsSupportRo,omitempty" name:"IsSupportRo"`

	// The reason why read-only instance is not supported
	// Note: This field may return null, indicating that no valid values can be obtained.
	NonsupportRoReason *string `json:"NonsupportRoReason,omitempty" name:"NonsupportRoReason"`
}

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

	// The max connections
	MaxUserConnections *int64 `json:"MaxUserConnections,omitempty" name:"MaxUserConnections"`
}

// Predefined struct for user
type ActivateInstanceRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// List of instance IDs in the format of `cynosdbmysql-ins-n7ocdslw` as displayed in the TDSQL-C for MySQL console. You can use the instance list querying API to query the ID, i.e., the `InstanceId` value in the output parameters.
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList"`
}

type ActivateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// List of instance IDs in the format of `cynosdbmysql-ins-n7ocdslw` as displayed in the TDSQL-C for MySQL console. You can use the instance list querying API to query the ID, i.e., the `InstanceId` value in the output parameters.
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList"`
}

func (r *ActivateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ActivateInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ActivateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ActivateInstanceResponseParams struct {
	// Task flow ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ActivateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ActivateInstanceResponseParams `json:"Response"`
}

func (r *ActivateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ActivateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddClusterSlaveZoneRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Replica AZ
	SlaveZone *string `json:"SlaveZone,omitempty" name:"SlaveZone"`
}

type AddClusterSlaveZoneRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Replica AZ
	SlaveZone *string `json:"SlaveZone,omitempty" name:"SlaveZone"`
}

func (r *AddClusterSlaveZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddClusterSlaveZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SlaveZone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddClusterSlaveZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddClusterSlaveZoneResponseParams struct {
	// Async FlowId
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddClusterSlaveZoneResponse struct {
	*tchttp.BaseResponse
	Response *AddClusterSlaveZoneResponseParams `json:"Response"`
}

func (r *AddClusterSlaveZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddClusterSlaveZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddInstancesRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Number of CPU cores
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// Memory in GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Number of added read-only instances. Value range: (0,16].
	ReadOnlyCount *int64 `json:"ReadOnlyCount,omitempty" name:"ReadOnlyCount"`

	// Instance group ID, which will be used when you add an instance in an existing RO group. If this parameter is left empty, an RO group will be created. But it is not recommended to pass in this parameter for the current version, as this version has been disused.
	InstanceGrpId *string `json:"InstanceGrpId,omitempty" name:"InstanceGrpId"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID. If `VpcId` is set, `SubnetId` is required.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// The port used when adding an RO group. Value range: [0,65535).
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Instance name. String length range: [0,64).
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Whether to automatically select a voucher. 1: yes; 0: no. Default value: 0
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// Database type. Valid values: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// Order source. String length range: [0,64).
	OrderSource *string `json:"OrderSource,omitempty" name:"OrderSource"`

	// Transaction mode. Valid values: `0` (place and pay for an order), `1` (place an order)
	DealMode *int64 `json:"DealMode,omitempty" name:"DealMode"`

	// Parameter template ID
	ParamTemplateId *int64 `json:"ParamTemplateId,omitempty" name:"ParamTemplateId"`

	// Parameter list, which is valid only if `InstanceParams` is passed in to `ParamTemplateId`.
	InstanceParams []*ModifyParamItem `json:"InstanceParams,omitempty" name:"InstanceParams"`

	// Security group ID. You can specify an security group when creating a read-only instance.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
}

type AddInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Number of CPU cores
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// Memory in GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Number of added read-only instances. Value range: (0,16].
	ReadOnlyCount *int64 `json:"ReadOnlyCount,omitempty" name:"ReadOnlyCount"`

	// Instance group ID, which will be used when you add an instance in an existing RO group. If this parameter is left empty, an RO group will be created. But it is not recommended to pass in this parameter for the current version, as this version has been disused.
	InstanceGrpId *string `json:"InstanceGrpId,omitempty" name:"InstanceGrpId"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID. If `VpcId` is set, `SubnetId` is required.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// The port used when adding an RO group. Value range: [0,65535).
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Instance name. String length range: [0,64).
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Whether to automatically select a voucher. 1: yes; 0: no. Default value: 0
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// Database type. Valid values: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// Order source. String length range: [0,64).
	OrderSource *string `json:"OrderSource,omitempty" name:"OrderSource"`

	// Transaction mode. Valid values: `0` (place and pay for an order), `1` (place an order)
	DealMode *int64 `json:"DealMode,omitempty" name:"DealMode"`

	// Parameter template ID
	ParamTemplateId *int64 `json:"ParamTemplateId,omitempty" name:"ParamTemplateId"`

	// Parameter list, which is valid only if `InstanceParams` is passed in to `ParamTemplateId`.
	InstanceParams []*ModifyParamItem `json:"InstanceParams,omitempty" name:"InstanceParams"`

	// Security group ID. You can specify an security group when creating a read-only instance.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
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
	delete(f, "ParamTemplateId")
	delete(f, "InstanceParams")
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddInstancesResponseParams struct {
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
}

type AddInstancesResponse struct {
	*tchttp.BaseResponse
	Response *AddInstancesResponseParams `json:"Response"`
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

type AuditRuleFilters struct {
	// Audit rule
	RuleFilters []*RuleFilters `json:"RuleFilters,omitempty" name:"RuleFilters"`
}

type AuditRuleTemplateInfo struct {
	// Rule template ID
	RuleTemplateId *string `json:"RuleTemplateId,omitempty" name:"RuleTemplateId"`

	// Rule template name
	RuleTemplateName *string `json:"RuleTemplateName,omitempty" name:"RuleTemplateName"`

	// Filter of the rule template
	RuleFilters []*RuleFilters `json:"RuleFilters,omitempty" name:"RuleFilters"`

	// Description of a rule template
	// Note: This field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Creation time of a rule template
	CreateAt *string `json:"CreateAt,omitempty" name:"CreateAt"`
}

type BackupFileInfo struct {
	// Snapshot file ID, which is deprecated. You need to use `BackupId`.
	SnapshotId *uint64 `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// Backup file name
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// Backup file size
	FileSize *uint64 `json:"FileSize,omitempty" name:"FileSize"`

	// Backup start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Backup end time
	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`

	// Backup type. Valid values: `snapshot` (snapshot backup), `logic` (logic backup).
	BackupType *string `json:"BackupType,omitempty" name:"BackupType"`

	// Back mode. auto: auto backup; manual: manual backup
	BackupMethod *string `json:"BackupMethod,omitempty" name:"BackupMethod"`

	// Backup file status. success: backup succeeded; fail: backup failed; creating: creating backup file; deleting: deleting backup file
	BackupStatus *string `json:"BackupStatus,omitempty" name:"BackupStatus"`

	// Backup file time
	SnapshotTime *string `json:"SnapshotTime,omitempty" name:"SnapshotTime"`

	// Backup ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupId *int64 `json:"BackupId,omitempty" name:"BackupId"`


	SnapShotType *string `json:"SnapShotType,omitempty" name:"SnapShotType"`

	// Backup file alias
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupName *string `json:"BackupName,omitempty" name:"BackupName"`
}

type BillingResourceInfo struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Instance ID list
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Order ID
	DealName *string `json:"DealName,omitempty" name:"DealName"`
}

type BinlogItem struct {
	// Binlog filename
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// File size in bytes
	FileSize *int64 `json:"FileSize,omitempty" name:"FileSize"`

	// Transaction start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Transaction end time
	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`

	// Binlog file ID
	BinlogId *int64 `json:"BinlogId,omitempty" name:"BinlogId"`
}

// Predefined struct for user
type CloseAuditServiceRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type CloseAuditServiceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *CloseAuditServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseAuditServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseAuditServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseAuditServiceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CloseAuditServiceResponse struct {
	*tchttp.BaseResponse
	Response *CloseAuditServiceResponseParams `json:"Response"`
}

func (r *CloseAuditServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseAuditServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

	// Instance role
	InstanceRole *string `json:"InstanceRole,omitempty" name:"InstanceRole"`
}

// Predefined struct for user
type CreateAccountsRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// List of new accounts
	Accounts []*NewAccount `json:"Accounts,omitempty" name:"Accounts"`
}

type CreateAccountsRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// List of new accounts
	Accounts []*NewAccount `json:"Accounts,omitempty" name:"Accounts"`
}

func (r *CreateAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Accounts")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccountsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAccountsResponse struct {
	*tchttp.BaseResponse
	Response *CreateAccountsResponseParams `json:"Response"`
}

func (r *CreateAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAuditRuleTemplateRequestParams struct {
	// Audit rule
	RuleFilters []*RuleFilters `json:"RuleFilters,omitempty" name:"RuleFilters"`

	// Rule template name
	RuleTemplateName *string `json:"RuleTemplateName,omitempty" name:"RuleTemplateName"`

	// Rule template description.
	Description *string `json:"Description,omitempty" name:"Description"`
}

type CreateAuditRuleTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Audit rule
	RuleFilters []*RuleFilters `json:"RuleFilters,omitempty" name:"RuleFilters"`

	// Rule template name
	RuleTemplateName *string `json:"RuleTemplateName,omitempty" name:"RuleTemplateName"`

	// Rule template description.
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateAuditRuleTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAuditRuleTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleFilters")
	delete(f, "RuleTemplateName")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAuditRuleTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAuditRuleTemplateResponseParams struct {
	// The generated rule template ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	RuleTemplateId *string `json:"RuleTemplateId,omitempty" name:"RuleTemplateId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAuditRuleTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateAuditRuleTemplateResponseParams `json:"Response"`
}

func (r *CreateAuditRuleTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAuditRuleTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBackupRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Backup type. Valid values: `logic` (logic backup), `snapshot` (physical backup)
	BackupType *string `json:"BackupType,omitempty" name:"BackupType"`

	// Backup database, which is valid when `BackupType` is `logic`.
	BackupDatabases []*string `json:"BackupDatabases,omitempty" name:"BackupDatabases"`

	// Backup table, which is valid when `BackupType` is `logic`.
	BackupTables []*DatabaseTables `json:"BackupTables,omitempty" name:"BackupTables"`

	// Backup name
	BackupName *string `json:"BackupName,omitempty" name:"BackupName"`
}

type CreateBackupRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Backup type. Valid values: `logic` (logic backup), `snapshot` (physical backup)
	BackupType *string `json:"BackupType,omitempty" name:"BackupType"`

	// Backup database, which is valid when `BackupType` is `logic`.
	BackupDatabases []*string `json:"BackupDatabases,omitempty" name:"BackupDatabases"`

	// Backup table, which is valid when `BackupType` is `logic`.
	BackupTables []*DatabaseTables `json:"BackupTables,omitempty" name:"BackupTables"`

	// Backup name
	BackupName *string `json:"BackupName,omitempty" name:"BackupName"`
}

func (r *CreateBackupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "BackupType")
	delete(f, "BackupDatabases")
	delete(f, "BackupTables")
	delete(f, "BackupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBackupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBackupResponseParams struct {
	// Async task flow ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateBackupResponse struct {
	*tchttp.BaseResponse
	Response *CreateBackupResponseParams `json:"Response"`
}

func (r *CreateBackupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClustersRequestParams struct {
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
	// <li> Valid values for `MYSQL`: 5.7 and 8.0 </li>
	DbVersion *string `json:"DbVersion,omitempty" name:"DbVersion"`

	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// It is required when `DbMode` is set to `NORMAL` or left empty.
	// Number of CPU cores of normal instance
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// It is required when `DbMode` is set to `NORMAL` or left empty.
	// Memory of a non-serverless instance in GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// This parameter has been deprecated.
	// Storage capacity in GB
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// Cluster name, which can contain less than 64 letters, digits, or symbols (-_.).
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// Account password, which must contain 8-64 characters in at least three of the following four types: uppercase letters, lowercase letters, digits, and symbols (~!@#$%^&*_-+=`|\(){}[]:;'<>,.?/).
	AdminPassword *string `json:"AdminPassword,omitempty" name:"AdminPassword"`

	// Port. Valid range: [0, 65535). Default value: 3306
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Billing mode. `0`: pay-as-you-go; `1`: monthly subscription. Default value: `0`
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// Number of purchased clusters. Valid range: [1,50]. Default value: 1
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// Rollback type:
	// noneRollback: no rollback;
	// snapRollback: rollback by snapshot;
	// timeRollback: rollback by time point
	RollbackStrategy *string `json:"RollbackStrategy,omitempty" name:"RollbackStrategy"`

	// `snapshotId` for snapshot rollback or `queryId` for time point rollback. 0 indicates to determine whether the time point is valid
	RollbackId *uint64 `json:"RollbackId,omitempty" name:"RollbackId"`

	// The source cluster ID passed in during rollback to find the source `poolId`
	OriginalClusterId *string `json:"OriginalClusterId,omitempty" name:"OriginalClusterId"`

	// Specified time for time point rollback or snapshot time for snapshot rollback
	ExpectTime *string `json:"ExpectTime,omitempty" name:"ExpectTime"`

	// This parameter has been deprecated.
	// Specified allowed time range for time point rollback
	ExpectTimeThresh *uint64 `json:"ExpectTimeThresh,omitempty" name:"ExpectTimeThresh"`

	// Storage upper limit of normal instance in GB
	// If `DbType` is `MYSQL` and the storage billing mode is monthly subscription, the parameter value can’t exceed the maximum storage corresponding to the CPU and memory specifications.
	StorageLimit *int64 `json:"StorageLimit,omitempty" name:"StorageLimit"`

	// Number of Instances. Valid range: (0,16]
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// Purchase duration of monthly subscription plan
	TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// Duration unit of monthly subscription. Valid values: `s`, `d`, `m`, `y`
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// Whether auto-renewal is enabled for monthly subscription plan. Default value: `0`
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// Whether to automatically select a voucher. `1`: yes; `0`: no. Default value: `0`
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// Number of instances (this parameter has been disused and is retained only for compatibility with existing instances)
	HaCount *int64 `json:"HaCount,omitempty" name:"HaCount"`

	// Order source
	OrderSource *string `json:"OrderSource,omitempty" name:"OrderSource"`

	// Array of tags to be bound to the created cluster
	ResourceTags []*Tag `json:"ResourceTags,omitempty" name:"ResourceTags"`

	// Database type
	// Valid values when `DbType` is `MYSQL` (default value: `NORMAL`):
	// <li>NORMAL</li>
	// <li>SERVERLESS</li>
	DbMode *string `json:"DbMode,omitempty" name:"DbMode"`

	// This parameter is required if `DbMode` is `SERVERLESS`.
	// Minimum number of CPU cores. For the value range, see the returned result of `DescribeServerlessInstanceSpecs`.
	MinCpu *float64 `json:"MinCpu,omitempty" name:"MinCpu"`

	// This parameter is required if `DbMode` is `SERVERLESS`.
	// Maximum number of CPU cores. For the value range, see the returned result of `DescribeServerlessInstanceSpecs`.
	MaxCpu *float64 `json:"MaxCpu,omitempty" name:"MaxCpu"`

	// This parameter specifies whether the cluster will be automatically paused if `DbMode` is `SERVERLESS`. Valid values:
	// <li>yes</li>
	// <li>no</li>
	// Default value: yes
	AutoPause *string `json:"AutoPause,omitempty" name:"AutoPause"`

	// This parameter specifies the delay for automatic cluster pause in seconds if `DbMode` is `SERVERLESS`. Value range: [600,691200]
	// Default value: `600`
	AutoPauseDelay *int64 `json:"AutoPauseDelay,omitempty" name:"AutoPauseDelay"`

	// The billing mode of cluster storage. Valid values: `0` (pay-as-you-go), `1` (monthly subscription). Default value: `0`.
	// If `DbType` is `MYSQL` and the billing mode of cluster compute is pay-as-you-go (or the `DbMode` is `SERVERLESS`), the billing mode of cluster storage must be pay-as-you-go.
	// Clusters with storage billed in monthly subscription can’t be cloned or rolled back.
	StoragePayMode *int64 `json:"StoragePayMode,omitempty" name:"StoragePayMode"`

	// Array of security group IDs
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// Array of alarm policy IDs
	AlarmPolicyIds []*string `json:"AlarmPolicyIds,omitempty" name:"AlarmPolicyIds"`

	// Array of parameters. Valid values: `character_set_server` (utf8｜latin1｜gbk｜utf8mb4), `lower_case_table_names`. 0: case-sensitive; 1: case-insensitive).
	ClusterParams []*ParamItem `json:"ClusterParams,omitempty" name:"ClusterParams"`

	// Transaction mode. Valid values: `0` (place and pay for an order), `1` (place an order)
	DealMode *int64 `json:"DealMode,omitempty" name:"DealMode"`

	// Parameter template ID, which can be obtained by querying parameter template information “DescribeParamTemplates”
	ParamTemplateId *int64 `json:"ParamTemplateId,omitempty" name:"ParamTemplateId"`

	// Multi-AZ address
	SlaveZone *string `json:"SlaveZone,omitempty" name:"SlaveZone"`

	// Instance initialization configuration information, which is used to select instances with different specifications when purchasing a cluster.
	InstanceInitInfos []*InstanceInitInfo `json:"InstanceInitInfos,omitempty" name:"InstanceInitInfos"`
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
	// <li> Valid values for `MYSQL`: 5.7 and 8.0 </li>
	DbVersion *string `json:"DbVersion,omitempty" name:"DbVersion"`

	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// It is required when `DbMode` is set to `NORMAL` or left empty.
	// Number of CPU cores of normal instance
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// It is required when `DbMode` is set to `NORMAL` or left empty.
	// Memory of a non-serverless instance in GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// This parameter has been deprecated.
	// Storage capacity in GB
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// Cluster name, which can contain less than 64 letters, digits, or symbols (-_.).
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// Account password, which must contain 8-64 characters in at least three of the following four types: uppercase letters, lowercase letters, digits, and symbols (~!@#$%^&*_-+=`|\(){}[]:;'<>,.?/).
	AdminPassword *string `json:"AdminPassword,omitempty" name:"AdminPassword"`

	// Port. Valid range: [0, 65535). Default value: 3306
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Billing mode. `0`: pay-as-you-go; `1`: monthly subscription. Default value: `0`
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// Number of purchased clusters. Valid range: [1,50]. Default value: 1
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// Rollback type:
	// noneRollback: no rollback;
	// snapRollback: rollback by snapshot;
	// timeRollback: rollback by time point
	RollbackStrategy *string `json:"RollbackStrategy,omitempty" name:"RollbackStrategy"`

	// `snapshotId` for snapshot rollback or `queryId` for time point rollback. 0 indicates to determine whether the time point is valid
	RollbackId *uint64 `json:"RollbackId,omitempty" name:"RollbackId"`

	// The source cluster ID passed in during rollback to find the source `poolId`
	OriginalClusterId *string `json:"OriginalClusterId,omitempty" name:"OriginalClusterId"`

	// Specified time for time point rollback or snapshot time for snapshot rollback
	ExpectTime *string `json:"ExpectTime,omitempty" name:"ExpectTime"`

	// This parameter has been deprecated.
	// Specified allowed time range for time point rollback
	ExpectTimeThresh *uint64 `json:"ExpectTimeThresh,omitempty" name:"ExpectTimeThresh"`

	// Storage upper limit of normal instance in GB
	// If `DbType` is `MYSQL` and the storage billing mode is monthly subscription, the parameter value can’t exceed the maximum storage corresponding to the CPU and memory specifications.
	StorageLimit *int64 `json:"StorageLimit,omitempty" name:"StorageLimit"`

	// Number of Instances. Valid range: (0,16]
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// Purchase duration of monthly subscription plan
	TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// Duration unit of monthly subscription. Valid values: `s`, `d`, `m`, `y`
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// Whether auto-renewal is enabled for monthly subscription plan. Default value: `0`
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// Whether to automatically select a voucher. `1`: yes; `0`: no. Default value: `0`
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// Number of instances (this parameter has been disused and is retained only for compatibility with existing instances)
	HaCount *int64 `json:"HaCount,omitempty" name:"HaCount"`

	// Order source
	OrderSource *string `json:"OrderSource,omitempty" name:"OrderSource"`

	// Array of tags to be bound to the created cluster
	ResourceTags []*Tag `json:"ResourceTags,omitempty" name:"ResourceTags"`

	// Database type
	// Valid values when `DbType` is `MYSQL` (default value: `NORMAL`):
	// <li>NORMAL</li>
	// <li>SERVERLESS</li>
	DbMode *string `json:"DbMode,omitempty" name:"DbMode"`

	// This parameter is required if `DbMode` is `SERVERLESS`.
	// Minimum number of CPU cores. For the value range, see the returned result of `DescribeServerlessInstanceSpecs`.
	MinCpu *float64 `json:"MinCpu,omitempty" name:"MinCpu"`

	// This parameter is required if `DbMode` is `SERVERLESS`.
	// Maximum number of CPU cores. For the value range, see the returned result of `DescribeServerlessInstanceSpecs`.
	MaxCpu *float64 `json:"MaxCpu,omitempty" name:"MaxCpu"`

	// This parameter specifies whether the cluster will be automatically paused if `DbMode` is `SERVERLESS`. Valid values:
	// <li>yes</li>
	// <li>no</li>
	// Default value: yes
	AutoPause *string `json:"AutoPause,omitempty" name:"AutoPause"`

	// This parameter specifies the delay for automatic cluster pause in seconds if `DbMode` is `SERVERLESS`. Value range: [600,691200]
	// Default value: `600`
	AutoPauseDelay *int64 `json:"AutoPauseDelay,omitempty" name:"AutoPauseDelay"`

	// The billing mode of cluster storage. Valid values: `0` (pay-as-you-go), `1` (monthly subscription). Default value: `0`.
	// If `DbType` is `MYSQL` and the billing mode of cluster compute is pay-as-you-go (or the `DbMode` is `SERVERLESS`), the billing mode of cluster storage must be pay-as-you-go.
	// Clusters with storage billed in monthly subscription can’t be cloned or rolled back.
	StoragePayMode *int64 `json:"StoragePayMode,omitempty" name:"StoragePayMode"`

	// Array of security group IDs
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// Array of alarm policy IDs
	AlarmPolicyIds []*string `json:"AlarmPolicyIds,omitempty" name:"AlarmPolicyIds"`

	// Array of parameters. Valid values: `character_set_server` (utf8｜latin1｜gbk｜utf8mb4), `lower_case_table_names`. 0: case-sensitive; 1: case-insensitive).
	ClusterParams []*ParamItem `json:"ClusterParams,omitempty" name:"ClusterParams"`

	// Transaction mode. Valid values: `0` (place and pay for an order), `1` (place an order)
	DealMode *int64 `json:"DealMode,omitempty" name:"DealMode"`

	// Parameter template ID, which can be obtained by querying parameter template information “DescribeParamTemplates”
	ParamTemplateId *int64 `json:"ParamTemplateId,omitempty" name:"ParamTemplateId"`

	// Multi-AZ address
	SlaveZone *string `json:"SlaveZone,omitempty" name:"SlaveZone"`

	// Instance initialization configuration information, which is used to select instances with different specifications when purchasing a cluster.
	InstanceInitInfos []*InstanceInitInfo `json:"InstanceInitInfos,omitempty" name:"InstanceInitInfos"`
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
	delete(f, "SlaveZone")
	delete(f, "InstanceInitInfos")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClustersResponseParams struct {
	// Freezing transaction ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	TranId *string `json:"TranId,omitempty" name:"TranId"`

	// Order ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`

	// List of resource IDs (This field has been deprecated. You need to use `dealNames` in the `DescribeResourcesByDealName` API to get resource IDs.)
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// List of cluster IDs (This field has been deprecated. You need to use `dealNames` in the `DescribeResourcesByDealName` API to get cluster IDs.)
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`

	// Big order ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	BigDealIds []*string `json:"BigDealIds,omitempty" name:"BigDealIds"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateClustersResponse struct {
	*tchttp.BaseResponse
	Response *CreateClustersResponseParams `json:"Response"`
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
	// Cluster status. Valid values are as follows:
	// creating
	// running
	// isolating
	// isolated
	// activating (removing isolation)
	// offlining (deactivating)
	// offlined (deactivated)
	// deleting
	// deleted
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

	// User UIN
	// Note: This field may return null, indicating that no valid values can be obtained.
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// Engine type
	// Note: This field may return null, indicating that no valid values can be obtained.
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// User `appid`
	// Note: This field may return null, indicating that no valid values can be obtained.
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// Cluster status description
	// Note: This field may return null, indicating that no valid values can be obtained.
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

	// Cluster creation time
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Billing mode. `0`: Pay-as-you-go; `1`: Monthly subscription.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// End time
	// Note: This field may return null, indicating that no valid values can be obtained.
	PeriodEndTime *string `json:"PeriodEndTime,omitempty" name:"PeriodEndTime"`

	// Cluster read-write VIP
	// Note: This field may return null, indicating that no valid values can be obtained.
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Cluster read-write vport
	// Note: This field may return null, indicating that no valid values can be obtained.
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// Project ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProjectID *int64 `json:"ProjectID,omitempty" name:"ProjectID"`

	// VPC ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// TDSQL-C kernel version
	// Note: This field may return null, indicating that no valid values can be obtained.
	CynosVersion *string `json:"CynosVersion,omitempty" name:"CynosVersion"`

	// Storage capacity
	// Note: This field may return null, indicating that no valid values can be obtained.
	StorageLimit *int64 `json:"StorageLimit,omitempty" name:"StorageLimit"`

	// Renewal flag
	// Note: This field may return null, indicating that no valid values can be obtained.
	RenewFlag *int64 `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// Task in progress
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProcessingTask *string `json:"ProcessingTask,omitempty" name:"ProcessingTask"`

	// Array of tasks in the cluster
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tasks []*ObjectTask `json:"Tasks,omitempty" name:"Tasks"`

	// Array of tags bound to the cluster
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResourceTags []*Tag `json:"ResourceTags,omitempty" name:"ResourceTags"`

	// Database type. Valid values: `NORMAL`, `SERVERLESS`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DbMode *string `json:"DbMode,omitempty" name:"DbMode"`

	// Serverless cluster status when the database type is `SERVERLESS`. Valid values:
	// `resume`
	// `pause`
	// Note: This field may return null, indicating that no valid values can be obtained.
	ServerlessStatus *string `json:"ServerlessStatus,omitempty" name:"ServerlessStatus"`

	// Prepaid cluster storage capacity
	// Note: This field may return null, indicating that no valid values can be obtained.
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// Cluster storage ID used in prepaid storage modification
	// Note: This field may return null, indicating that no valid values can be obtained.
	StorageId *string `json:"StorageId,omitempty" name:"StorageId"`

	// Billing mode of cluster storage. Valid values: `0` (pay-as-you-go), `1` (monthly subscription).
	// Note: This field may return null, indicating that no valid values can be obtained.
	StoragePayMode *int64 `json:"StoragePayMode,omitempty" name:"StoragePayMode"`

	// The minimum storage corresponding to the compute specification of the cluster
	// Note: This field may return null, indicating that no valid values can be obtained.
	MinStorageSize *int64 `json:"MinStorageSize,omitempty" name:"MinStorageSize"`

	// The maximum storage corresponding to the compute specification of the cluster
	// Note: This field may return null, indicating that no valid values can be obtained.
	MaxStorageSize *int64 `json:"MaxStorageSize,omitempty" name:"MaxStorageSize"`

	// Network information of the cluster
	// Note: This field may return null, indicating that no valid values can be obtained.
	NetAddrs []*NetAddr `json:"NetAddrs,omitempty" name:"NetAddrs"`

	// Physical AZ
	// Note: This field may return null, indicating that no valid values can be obtained.
	PhysicalZone *string `json:"PhysicalZone,omitempty" name:"PhysicalZone"`

	// Primary AZ
	// Note: This field may return null, indicating that no valid values can be obtained.
	MasterZone *string `json:"MasterZone,omitempty" name:"MasterZone"`

	// Whether there is a secondary AZ
	// Note: This field may return null, indicating that no valid values can be obtained.
	HasSlaveZone *string `json:"HasSlaveZone,omitempty" name:"HasSlaveZone"`

	// Secondary AZ
	// Note: This field may return null, indicating that no valid values can be obtained.
	SlaveZones []*string `json:"SlaveZones,omitempty" name:"SlaveZones"`

	// Business type
	// Note: This field may return null, indicating that no valid values can be obtained.
	BusinessType *string `json:"BusinessType,omitempty" name:"BusinessType"`

	// Whether to freeze
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsFreeze *string `json:"IsFreeze,omitempty" name:"IsFreeze"`

	// Order source
	// Note: This field may return null, indicating that no valid values can be obtained.
	OrderSource *string `json:"OrderSource,omitempty" name:"OrderSource"`

	// Capability
	// Note: This field may return null, indicating that no valid values can be obtained.
	Ability *Ability `json:"Ability,omitempty" name:"Ability"`
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

	// Binlog switch. Valid values: `ON`, `OFF`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LogBin *string `json:"LogBin,omitempty" name:"LogBin"`

	// PITR type. Valid values: `normal`, `redo_pitr`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PitrType *string `json:"PitrType,omitempty" name:"PitrType"`

	// Physical AZ
	// Note: This field may return null, indicating that no valid values can be obtained.
	PhysicalZone *string `json:"PhysicalZone,omitempty" name:"PhysicalZone"`

	// Storage ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	StorageId *string `json:"StorageId,omitempty" name:"StorageId"`

	// Storage capacity in GB
	// Note: This field may return null, indicating that no valid values can be obtained.
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// Maximum storage specification in GB
	// Note: This field may return null, indicating that no valid values can be obtained.
	MaxStorageSize *int64 `json:"MaxStorageSize,omitempty" name:"MaxStorageSize"`

	// Minimum storage specification in GB
	// Note: This field may return null, indicating that no valid values can be obtained.
	MinStorageSize *int64 `json:"MinStorageSize,omitempty" name:"MinStorageSize"`

	// Storage billing mode. Valid values: `1` (monthly subscription), `0` (pay-as-you-go).
	// Note: This field may return null, indicating that no valid values can be obtained.
	StoragePayMode *int64 `json:"StoragePayMode,omitempty" name:"StoragePayMode"`

	// Database type. Valid values: `normal`, `serverless`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DbMode *string `json:"DbMode,omitempty" name:"DbMode"`

	// Maximum storage space
	// Note: This field may return null, indicating that no valid values can be obtained.
	StorageLimit *int64 `json:"StorageLimit,omitempty" name:"StorageLimit"`

	// Features supported by the cluster
	// Note: This field may return null, indicating that no valid values can be obtained.
	Ability *Ability `json:"Ability,omitempty" name:"Ability"`

	// TDSQL-C version
	// Note: This field may return null, indicating that no valid values can be obtained.
	CynosVersion *string `json:"CynosVersion,omitempty" name:"CynosVersion"`

	// Business type
	// Note: This field may return null, indicating that no valid values can be obtained.
	BusinessType *string `json:"BusinessType,omitempty" name:"BusinessType"`

	// Whether there is a secondary AZ
	// Note: This field may return null, indicating that no valid values can be obtained.
	HasSlaveZone *string `json:"HasSlaveZone,omitempty" name:"HasSlaveZone"`

	// Whether to freeze
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsFreeze *string `json:"IsFreeze,omitempty" name:"IsFreeze"`

	// Task list
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tasks []*ObjectTask `json:"Tasks,omitempty" name:"Tasks"`

	// Primary AZ
	// Note: This field may return null, indicating that no valid values can be obtained.
	MasterZone *string `json:"MasterZone,omitempty" name:"MasterZone"`

	// Secondary AZ list
	// Note: This field may return null, indicating that no valid values can be obtained.
	SlaveZones []*string `json:"SlaveZones,omitempty" name:"SlaveZones"`

	// Proxy status
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProxyStatus *string `json:"ProxyStatus,omitempty" name:"ProxyStatus"`

	// Whether to skip the transaction
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsSkipTrade *string `json:"IsSkipTrade,omitempty" name:"IsSkipTrade"`

	// Whether to enable password complexity
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsOpenPasswordComplexity *string `json:"IsOpenPasswordComplexity,omitempty" name:"IsOpenPasswordComplexity"`

	// Network type
	// Note: This field may return null, indicating that no valid values can be obtained.
	NetworkStatus *string `json:"NetworkStatus,omitempty" name:"NetworkStatus"`
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

	// Physical zone
	PhysicalZone *string `json:"PhysicalZone,omitempty" name:"PhysicalZone"`

	// Business type
	// Note: This field may return null, indicating that no valid value can be obtained.
	BusinessType *string `json:"BusinessType,omitempty" name:"BusinessType"`

	// Task
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tasks []*ObjectTask `json:"Tasks,omitempty" name:"Tasks"`

	// Whether to freeze
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsFreeze *string `json:"IsFreeze,omitempty" name:"IsFreeze"`

	// The resource tag
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResourceTags []*Tag `json:"ResourceTags,omitempty" name:"ResourceTags"`

	// Source AZ
	// Note: This field may return null, indicating that no valid value can be obtained.
	MasterZone *string `json:"MasterZone,omitempty" name:"MasterZone"`

	// Replica AZ
	// Note: This field may return null, indicating that no valid value can be obtained.
	SlaveZones []*string `json:"SlaveZones,omitempty" name:"SlaveZones"`

	// Instance network information
	// Note: This field may return null, indicating that no valid value can be obtained.
	InstanceNetInfo []*InstanceNetInfo `json:"InstanceNetInfo,omitempty" name:"InstanceNetInfo"`
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
	// User `appId`
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

	// VPC ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// Subnet ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// Information of the old IP
	// Note: This field may return null, indicating that no valid values can be obtained.
	OldAddrInfo *OldAddrInfo `json:"OldAddrInfo,omitempty" name:"OldAddrInfo"`

	// Task in progress
	ProcessingTasks []*string `json:"ProcessingTasks,omitempty" name:"ProcessingTasks"`

	// Task list
	Tasks []*ObjectTask `json:"Tasks,omitempty" name:"Tasks"`

	// biz_net_service table ID
	NetServiceId *int64 `json:"NetServiceId,omitempty" name:"NetServiceId"`
}

type DatabaseTables struct {
	// Database name
	// Note: This field may return null, indicating that no valid values can be obtained.
	Database *string `json:"Database,omitempty" name:"Database"`

	// Table name list
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tables []*string `json:"Tables,omitempty" name:"Tables"`
}

// Predefined struct for user
type DeleteAuditRuleTemplatesRequestParams struct {
	// Audit rule template ID
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitempty" name:"RuleTemplateIds"`
}

type DeleteAuditRuleTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// Audit rule template ID
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitempty" name:"RuleTemplateIds"`
}

func (r *DeleteAuditRuleTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAuditRuleTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleTemplateIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAuditRuleTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAuditRuleTemplatesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteAuditRuleTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAuditRuleTemplatesResponseParams `json:"Response"`
}

func (r *DeleteAuditRuleTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAuditRuleTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBackupRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Backup file ID. This field is used by legacy versions and thus not recommended.
	SnapshotIdList []*int64 `json:"SnapshotIdList,omitempty" name:"SnapshotIdList"`

	// Backup file ID. This field is recommended.
	BackupIds []*int64 `json:"BackupIds,omitempty" name:"BackupIds"`
}

type DeleteBackupRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Backup file ID. This field is used by legacy versions and thus not recommended.
	SnapshotIdList []*int64 `json:"SnapshotIdList,omitempty" name:"SnapshotIdList"`

	// Backup file ID. This field is recommended.
	BackupIds []*int64 `json:"BackupIds,omitempty" name:"BackupIds"`
}

func (r *DeleteBackupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBackupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SnapshotIdList")
	delete(f, "BackupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBackupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBackupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteBackupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteBackupResponseParams `json:"Response"`
}

func (r *DeleteBackupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountsRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// List of accounts to be filtered
	AccountNames []*string `json:"AccountNames,omitempty" name:"AccountNames"`

	// Database type. Valid values: 
	// <li> MYSQL </li>
	// This parameter has been disused.
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// List of accounts to be filtered
	Hosts []*string `json:"Hosts,omitempty" name:"Hosts"`

	// Maximum entries returned per page
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeAccountsRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// List of accounts to be filtered
	AccountNames []*string `json:"AccountNames,omitempty" name:"AccountNames"`

	// Database type. Valid values: 
	// <li> MYSQL </li>
	// This parameter has been disused.
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// List of accounts to be filtered
	Hosts []*string `json:"Hosts,omitempty" name:"Hosts"`

	// Maximum entries returned per page
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
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
	delete(f, "Hosts")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountsResponseParams struct {
	// Database account list
	// Note: This field may return null, indicating that no valid values can be obtained.
	AccountSet []*Account `json:"AccountSet,omitempty" name:"AccountSet"`

	// Total number of accounts
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAccountsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccountsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeAuditRuleTemplatesRequestParams struct {
	// Rule template ID
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitempty" name:"RuleTemplateIds"`

	// Rule template name
	RuleTemplateNames []*string `json:"RuleTemplateNames,omitempty" name:"RuleTemplateNames"`

	// Number of results returned per request. Default value: `20`.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: `0`.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeAuditRuleTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// Rule template ID
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitempty" name:"RuleTemplateIds"`

	// Rule template name
	RuleTemplateNames []*string `json:"RuleTemplateNames,omitempty" name:"RuleTemplateNames"`

	// Number of results returned per request. Default value: `20`.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: `0`.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeAuditRuleTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditRuleTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleTemplateIds")
	delete(f, "RuleTemplateNames")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuditRuleTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditRuleTemplatesResponseParams struct {
	// Number of eligible instances
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of rule template details
	// Note: This field may return null, indicating that no valid values can be obtained.
	Items []*AuditRuleTemplateInfo `json:"Items,omitempty" name:"Items"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAuditRuleTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuditRuleTemplatesResponseParams `json:"Response"`
}

func (r *DescribeAuditRuleTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditRuleTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditRuleWithInstanceIdsRequestParams struct {
	// Instance ID. Currently, only one single instance can be queried.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type DescribeAuditRuleWithInstanceIdsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. Currently, only one single instance can be queried.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DescribeAuditRuleWithInstanceIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditRuleWithInstanceIdsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuditRuleWithInstanceIdsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditRuleWithInstanceIdsResponseParams struct {
	// None
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Audit rule information of the instance
	// Note: This field may return null, indicating that no valid values can be obtained.
	Items []*InstanceAuditRule `json:"Items,omitempty" name:"Items"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAuditRuleWithInstanceIdsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuditRuleWithInstanceIdsResponseParams `json:"Response"`
}

func (r *DescribeAuditRuleWithInstanceIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditRuleWithInstanceIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupConfigRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
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

// Predefined struct for user
type DescribeBackupConfigResponseParams struct {
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
}

type DescribeBackupConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupConfigResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeBackupDownloadUrlRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Backup ID
	BackupId *int64 `json:"BackupId,omitempty" name:"BackupId"`
}

type DescribeBackupDownloadUrlRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Backup ID
	BackupId *int64 `json:"BackupId,omitempty" name:"BackupId"`
}

func (r *DescribeBackupDownloadUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDownloadUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "BackupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupDownloadUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupDownloadUrlResponseParams struct {
	// Backup download address
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBackupDownloadUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupDownloadUrlResponseParams `json:"Response"`
}

func (r *DescribeBackupDownloadUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDownloadUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupListRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The number of results to be returned. Value range: (0,100]
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Record offset. Value range: [0,INF)
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Database type. Valid values: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// Backup ID
	BackupIds []*int64 `json:"BackupIds,omitempty" name:"BackupIds"`

	// Backup type. Valid values: `snapshot` (snapshot backup), `logic` (logic backup).
	BackupType *string `json:"BackupType,omitempty" name:"BackupType"`

	// Back mode. Valid values: `auto` (automatic backup), `manual` (manual backup)
	BackupMethod *string `json:"BackupMethod,omitempty" name:"BackupMethod"`


	SnapShotType *string `json:"SnapShotType,omitempty" name:"SnapShotType"`

	// Backup start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Backup end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`


	FileNames []*string `json:"FileNames,omitempty" name:"FileNames"`

	// Backup alias, which supports fuzzy query.
	BackupNames []*string `json:"BackupNames,omitempty" name:"BackupNames"`

	// ID list of the snapshot backup
	SnapshotIdList []*int64 `json:"SnapshotIdList,omitempty" name:"SnapshotIdList"`
}

type DescribeBackupListRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The number of results to be returned. Value range: (0,100]
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Record offset. Value range: [0,INF)
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Database type. Valid values: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// Backup ID
	BackupIds []*int64 `json:"BackupIds,omitempty" name:"BackupIds"`

	// Backup type. Valid values: `snapshot` (snapshot backup), `logic` (logic backup).
	BackupType *string `json:"BackupType,omitempty" name:"BackupType"`

	// Back mode. Valid values: `auto` (automatic backup), `manual` (manual backup)
	BackupMethod *string `json:"BackupMethod,omitempty" name:"BackupMethod"`

	SnapShotType *string `json:"SnapShotType,omitempty" name:"SnapShotType"`

	// Backup start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Backup end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	FileNames []*string `json:"FileNames,omitempty" name:"FileNames"`

	// Backup alias, which supports fuzzy query.
	BackupNames []*string `json:"BackupNames,omitempty" name:"BackupNames"`

	// ID list of the snapshot backup
	SnapshotIdList []*int64 `json:"SnapshotIdList,omitempty" name:"SnapshotIdList"`
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
	delete(f, "BackupIds")
	delete(f, "BackupType")
	delete(f, "BackupMethod")
	delete(f, "SnapShotType")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "FileNames")
	delete(f, "BackupNames")
	delete(f, "SnapshotIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupListResponseParams struct {
	// Total number of backup files
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Backup file list
	BackupList []*BackupFileInfo `json:"BackupList,omitempty" name:"BackupList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBackupListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeBinlogDownloadUrlRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Binlog file ID
	BinlogId *int64 `json:"BinlogId,omitempty" name:"BinlogId"`
}

type DescribeBinlogDownloadUrlRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Binlog file ID
	BinlogId *int64 `json:"BinlogId,omitempty" name:"BinlogId"`
}

func (r *DescribeBinlogDownloadUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBinlogDownloadUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "BinlogId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBinlogDownloadUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBinlogDownloadUrlResponseParams struct {
	// Download address
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBinlogDownloadUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBinlogDownloadUrlResponseParams `json:"Response"`
}

func (r *DescribeBinlogDownloadUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBinlogDownloadUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBinlogSaveDaysRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type DescribeBinlogSaveDaysRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeBinlogSaveDaysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBinlogSaveDaysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBinlogSaveDaysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBinlogSaveDaysResponseParams struct {
	// Binlog retention period in days
	BinlogSaveDays *int64 `json:"BinlogSaveDays,omitempty" name:"BinlogSaveDays"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBinlogSaveDaysResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBinlogSaveDaysResponseParams `json:"Response"`
}

func (r *DescribeBinlogSaveDaysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBinlogSaveDaysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBinlogsRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeBinlogsRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeBinlogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBinlogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBinlogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBinlogsResponseParams struct {
	// Total number of records
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Binlog list
	// Note: This field may return null, indicating that no valid values can be obtained.
	Binlogs []*BinlogItem `json:"Binlogs,omitempty" name:"Binlogs"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBinlogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBinlogsResponseParams `json:"Response"`
}

func (r *DescribeBinlogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBinlogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterDetailRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
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

// Predefined struct for user
type DescribeClusterDetailResponseParams struct {
	// Cluster details
	Detail *CynosdbClusterDetail `json:"Detail,omitempty" name:"Detail"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClusterDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterDetailResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeClusterInstanceGrpsRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
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

// Predefined struct for user
type DescribeClusterInstanceGrpsResponseParams struct {
	// Number of instance groups
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Instance group list
	InstanceGrpInfoList []*CynosdbInstanceGrp `json:"InstanceGrpInfoList,omitempty" name:"InstanceGrpInfoList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClusterInstanceGrpsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterInstanceGrpsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeClusterParamsRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Parameter name
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`
}

type DescribeClusterParamsRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Parameter name
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`
}

func (r *DescribeClusterParamsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterParamsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ParamName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterParamsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterParamsResponseParams struct {
	// Number of parameters
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Instance parameter list
	// Note: This field may return null, indicating that no valid values can be obtained.
	Items []*ParamInfo `json:"Items,omitempty" name:"Items"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClusterParamsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterParamsResponseParams `json:"Response"`
}

func (r *DescribeClusterParamsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterParamsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClustersRequestParams struct {
	// Engine type. Currently, `MYSQL` is supported.
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

type DescribeClustersRequest struct {
	*tchttp.BaseRequest
	
	// Engine type. Currently, `MYSQL` is supported.
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

// Predefined struct for user
type DescribeClustersResponseParams struct {
	// Number of clusters
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Cluster list
	ClusterSet []*CynosdbCluster `json:"ClusterSet,omitempty" name:"ClusterSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClustersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClustersResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeDBSecurityGroupsRequestParams struct {
	// Instance group ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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

// Predefined struct for user
type DescribeDBSecurityGroupsResponseParams struct {
	// Security group information
	Groups []*SecurityGroup `json:"Groups,omitempty" name:"Groups"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDBSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBSecurityGroupsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeFlowRequestParams struct {
	// Task flow ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
}

type DescribeFlowRequest struct {
	*tchttp.BaseRequest
	
	// Task flow ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
}

func (r *DescribeFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowResponseParams struct {
	// Task flow status. Valid values: `0` (succeeded), `1` (failed), `2` (Processing).
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFlowResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFlowResponseParams `json:"Response"`
}

func (r *DescribeFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceDetailRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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

// Predefined struct for user
type DescribeInstanceDetailResponseParams struct {
	// Instance details
	Detail *CynosdbInstanceDetail `json:"Detail,omitempty" name:"Detail"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstanceDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceDetailResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeInstanceSlowQueriesRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Transaction start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Transaction end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Maximum number
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Username
	Username *string `json:"Username,omitempty" name:"Username"`

	// Client host
	Host *string `json:"Host,omitempty" name:"Host"`

	// Database name
	Database *string `json:"Database,omitempty" name:"Database"`

	// Sorting field. Valid values: QueryTime, LockTime, RowsExamined, RowsSent.
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: asc, desc.
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
}

type DescribeInstanceSlowQueriesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Transaction start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Transaction end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Maximum number
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Username
	Username *string `json:"Username,omitempty" name:"Username"`

	// Client host
	Host *string `json:"Host,omitempty" name:"Host"`

	// Database name
	Database *string `json:"Database,omitempty" name:"Database"`

	// Sorting field. Valid values: QueryTime, LockTime, RowsExamined, RowsSent.
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: asc, desc.
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
}

func (r *DescribeInstanceSlowQueriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceSlowQueriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Username")
	delete(f, "Host")
	delete(f, "Database")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceSlowQueriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceSlowQueriesResponseParams struct {
	// Total number
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Slow query record
	SlowQueries []*SlowQueriesItem `json:"SlowQueries,omitempty" name:"SlowQueries"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstanceSlowQueriesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceSlowQueriesResponseParams `json:"Response"`
}

func (r *DescribeInstanceSlowQueriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceSlowQueriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceSpecsRequestParams struct {
	// Database type. Valid values: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// Whether to return the AZ information.
	IncludeZoneStocks *bool `json:"IncludeZoneStocks,omitempty" name:"IncludeZoneStocks"`
}

type DescribeInstanceSpecsRequest struct {
	*tchttp.BaseRequest
	
	// Database type. Valid values: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// Whether to return the AZ information.
	IncludeZoneStocks *bool `json:"IncludeZoneStocks,omitempty" name:"IncludeZoneStocks"`
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
	delete(f, "IncludeZoneStocks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceSpecsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceSpecsResponseParams struct {
	// Specification information
	InstanceSpecSet []*InstanceSpec `json:"InstanceSpecSet,omitempty" name:"InstanceSpecSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstanceSpecsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceSpecsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeInstancesRequestParams struct {
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

	// Engine type. Currently, `MYSQL` is supported.
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// Instance status. Valid values:
	// creating
	// running
	// isolating
	// isolated
	// activating: Removing the instance from isolation
	// offlining: Eliminating the instance
	// offlined: Instance eliminated
	Status *string `json:"Status,omitempty" name:"Status"`

	// Instance ID list
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
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

	// Engine type. Currently, `MYSQL` is supported.
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// Instance status. Valid values:
	// creating
	// running
	// isolating
	// isolated
	// activating: Removing the instance from isolation
	// offlining: Eliminating the instance
	// offlined: Instance eliminated
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

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// Number of instances
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Instance list
	InstanceSet []*CynosdbInstance `json:"InstanceSet,omitempty" name:"InstanceSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeMaintainPeriodRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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

// Predefined struct for user
type DescribeMaintainPeriodResponseParams struct {
	// Maintenance days of the week
	MaintainWeekDays []*string `json:"MaintainWeekDays,omitempty" name:"MaintainWeekDays"`

	// Maintenance start time in seconds
	MaintainStartTime *int64 `json:"MaintainStartTime,omitempty" name:"MaintainStartTime"`

	// Maintenance duration in seconds
	MaintainDuration *int64 `json:"MaintainDuration,omitempty" name:"MaintainDuration"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMaintainPeriodResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMaintainPeriodResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeParamTemplatesRequestParams struct {
	// Database engine version number
	EngineVersions []*string `json:"EngineVersions,omitempty" name:"EngineVersions"`

	// Template name
	TemplateNames []*string `json:"TemplateNames,omitempty" name:"TemplateNames"`

	// Template ID
	TemplateIds []*int64 `json:"TemplateIds,omitempty" name:"TemplateIds"`

	// Database Type. Valid values: `NORMAL`, `SERVERLESS`.
	DbModes []*string `json:"DbModes,omitempty" name:"DbModes"`

	// Offset for query
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Limit on queries
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Product type of the queried template
	Products []*string `json:"Products,omitempty" name:"Products"`

	// Template type
	TemplateTypes []*string `json:"TemplateTypes,omitempty" name:"TemplateTypes"`

	// Version type
	EngineTypes []*string `json:"EngineTypes,omitempty" name:"EngineTypes"`

	// The sorting order of the returned results
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: `desc`, `asc `.
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

type DescribeParamTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// Database engine version number
	EngineVersions []*string `json:"EngineVersions,omitempty" name:"EngineVersions"`

	// Template name
	TemplateNames []*string `json:"TemplateNames,omitempty" name:"TemplateNames"`

	// Template ID
	TemplateIds []*int64 `json:"TemplateIds,omitempty" name:"TemplateIds"`

	// Database Type. Valid values: `NORMAL`, `SERVERLESS`.
	DbModes []*string `json:"DbModes,omitempty" name:"DbModes"`

	// Offset for query
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Limit on queries
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Product type of the queried template
	Products []*string `json:"Products,omitempty" name:"Products"`

	// Template type
	TemplateTypes []*string `json:"TemplateTypes,omitempty" name:"TemplateTypes"`

	// Version type
	EngineTypes []*string `json:"EngineTypes,omitempty" name:"EngineTypes"`

	// The sorting order of the returned results
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: `desc`, `asc `.
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

func (r *DescribeParamTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParamTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EngineVersions")
	delete(f, "TemplateNames")
	delete(f, "TemplateIds")
	delete(f, "DbModes")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Products")
	delete(f, "TemplateTypes")
	delete(f, "EngineTypes")
	delete(f, "OrderBy")
	delete(f, "OrderDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeParamTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeParamTemplatesResponseParams struct {
	// Number of parameter templates
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Parameter template information
	Items []*ParamTemplateListInfo `json:"Items,omitempty" name:"Items"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeParamTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeParamTemplatesResponseParams `json:"Response"`
}

func (r *DescribeParamTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParamTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectSecurityGroupsRequestParams struct {
	// Project ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Maximum entries returned per page
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Search by keyword
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
}

type DescribeProjectSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Project ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Maximum entries returned per page
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Search by keyword
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
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
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "SearchKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectSecurityGroupsResponseParams struct {
	// Security group details
	Groups []*SecurityGroup `json:"Groups,omitempty" name:"Groups"`

	// The total number of groups
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeProjectSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProjectSecurityGroupsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeResourcesByDealNameRequestParams struct {
	// Order ID. (If the cluster is not delivered yet, the `DescribeResourcesByDealName` API may return the `InvalidParameterValue.DealNameNotFound` error. Call the API again until it succeeds.)
	DealName *string `json:"DealName,omitempty" name:"DealName"`

	// Order ID, which can be used to query the resource information of multiple orders ID (If the cluster is not delivered yet, the `DescribeResourcesByDealName` API may return the `InvalidParameterValue.DealNameNotFound` error. Call the API again until it succeeds.)
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`
}

type DescribeResourcesByDealNameRequest struct {
	*tchttp.BaseRequest
	
	// Order ID. (If the cluster is not delivered yet, the `DescribeResourcesByDealName` API may return the `InvalidParameterValue.DealNameNotFound` error. Call the API again until it succeeds.)
	DealName *string `json:"DealName,omitempty" name:"DealName"`

	// Order ID, which can be used to query the resource information of multiple orders ID (If the cluster is not delivered yet, the `DescribeResourcesByDealName` API may return the `InvalidParameterValue.DealNameNotFound` error. Call the API again until it succeeds.)
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`
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
	delete(f, "DealNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourcesByDealNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourcesByDealNameResponseParams struct {
	// Billable resource ID information array
	BillingResourceInfos []*BillingResourceInfo `json:"BillingResourceInfos,omitempty" name:"BillingResourceInfos"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeResourcesByDealNameResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourcesByDealNameResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeRollbackTimeRangeRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
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

// Predefined struct for user
type DescribeRollbackTimeRangeResponseParams struct {
	// Start time of valid rollback time range (disused)
	// Note: This field may return null, indicating that no valid values can be obtained.
	TimeRangeStart *string `json:"TimeRangeStart,omitempty" name:"TimeRangeStart"`

	// End time of valid rollback time range (disused)
	// Note: This field may return null, indicating that no valid values can be obtained.
	TimeRangeEnd *string `json:"TimeRangeEnd,omitempty" name:"TimeRangeEnd"`

	// Time range available for rollback
	RollbackTimeRanges []*RollbackTimeRange `json:"RollbackTimeRanges,omitempty" name:"RollbackTimeRanges"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRollbackTimeRangeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRollbackTimeRangeResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeRollbackTimeValidityRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Expected time point to roll back to
	ExpectTime *string `json:"ExpectTime,omitempty" name:"ExpectTime"`

	// Error tolerance range for rollback time point
	ExpectTimeThresh *uint64 `json:"ExpectTimeThresh,omitempty" name:"ExpectTimeThresh"`
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

// Predefined struct for user
type DescribeRollbackTimeValidityResponseParams struct {
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
}

type DescribeRollbackTimeValidityResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRollbackTimeValidityResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeZonesRequestParams struct {
	// Whether the virtual zone is included.–
	IncludeVirtualZones *bool `json:"IncludeVirtualZones,omitempty" name:"IncludeVirtualZones"`

	// Whether to display all AZs in a region and the user’s permissions in each AZ.
	ShowPermission *bool `json:"ShowPermission,omitempty" name:"ShowPermission"`
}

type DescribeZonesRequest struct {
	*tchttp.BaseRequest
	
	// Whether the virtual zone is included.–
	IncludeVirtualZones *bool `json:"IncludeVirtualZones,omitempty" name:"IncludeVirtualZones"`

	// Whether to display all AZs in a region and the user’s permissions in each AZ.
	ShowPermission *bool `json:"ShowPermission,omitempty" name:"ShowPermission"`
}

func (r *DescribeZonesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZonesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IncludeVirtualZones")
	delete(f, "ShowPermission")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZonesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZonesResponseParams struct {
	// Region information
	RegionSet []*SaleRegion `json:"RegionSet,omitempty" name:"RegionSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeZonesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeZonesResponseParams `json:"Response"`
}

func (r *DescribeZonesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportInstanceSlowQueriesRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Transaction start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Transaction end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Maximum number
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Username
	Username *string `json:"Username,omitempty" name:"Username"`

	// Client host
	Host *string `json:"Host,omitempty" name:"Host"`

	// Database name
	Database *string `json:"Database,omitempty" name:"Database"`

	// File type. Valid values: csv, original.
	FileType *string `json:"FileType,omitempty" name:"FileType"`
}

type ExportInstanceSlowQueriesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Transaction start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Transaction end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Maximum number
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Username
	Username *string `json:"Username,omitempty" name:"Username"`

	// Client host
	Host *string `json:"Host,omitempty" name:"Host"`

	// Database name
	Database *string `json:"Database,omitempty" name:"Database"`

	// File type. Valid values: csv, original.
	FileType *string `json:"FileType,omitempty" name:"FileType"`
}

func (r *ExportInstanceSlowQueriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportInstanceSlowQueriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Username")
	delete(f, "Host")
	delete(f, "Database")
	delete(f, "FileType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportInstanceSlowQueriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportInstanceSlowQueriesResponseParams struct {
	// Slow query export content
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ExportInstanceSlowQueriesResponse struct {
	*tchttp.BaseResponse
	Response *ExportInstanceSlowQueriesResponseParams `json:"Response"`
}

func (r *ExportInstanceSlowQueriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportInstanceSlowQueriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceCreateRequestParams struct {
	// AZ
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Number of compute node to purchase
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// Instance type for purchase. Valid values: `PREPAID`, `POSTPAID`, `SERVERLESS`.
	InstancePayMode *string `json:"InstancePayMode,omitempty" name:"InstancePayMode"`

	// Storage type for purchase. Valid values: `PREPAID`, `POSTPAID`.
	StoragePayMode *string `json:"StoragePayMode,omitempty" name:"StoragePayMode"`

	// Number of CPU cores, which is required when `InstancePayMode` is `PREPAID` or `POSTPAID`.
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// Memory size in GB, which is required when `InstancePayMode` is `PREPAID` or `POSTPAID`.
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// CCU size, which is required when `InstancePayMode` is `SERVERLESS`.
	Ccu *float64 `json:"Ccu,omitempty" name:"Ccu"`

	// Storage size, which is required when `StoragePayMode` is `PREPAID`.
	StorageLimit *int64 `json:"StorageLimit,omitempty" name:"StorageLimit"`

	// Validity period, which is required when `InstancePayMode` is `PREPAID`.
	TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// Duration unit, which is required when `InstancePayMode` is `PREPAID`. Valid values: `m` (month), `d` (day).
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
}

type InquirePriceCreateRequest struct {
	*tchttp.BaseRequest
	
	// AZ
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Number of compute node to purchase
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// Instance type for purchase. Valid values: `PREPAID`, `POSTPAID`, `SERVERLESS`.
	InstancePayMode *string `json:"InstancePayMode,omitempty" name:"InstancePayMode"`

	// Storage type for purchase. Valid values: `PREPAID`, `POSTPAID`.
	StoragePayMode *string `json:"StoragePayMode,omitempty" name:"StoragePayMode"`

	// Number of CPU cores, which is required when `InstancePayMode` is `PREPAID` or `POSTPAID`.
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// Memory size in GB, which is required when `InstancePayMode` is `PREPAID` or `POSTPAID`.
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// CCU size, which is required when `InstancePayMode` is `SERVERLESS`.
	Ccu *float64 `json:"Ccu,omitempty" name:"Ccu"`

	// Storage size, which is required when `StoragePayMode` is `PREPAID`.
	StorageLimit *int64 `json:"StorageLimit,omitempty" name:"StorageLimit"`

	// Validity period, which is required when `InstancePayMode` is `PREPAID`.
	TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// Duration unit, which is required when `InstancePayMode` is `PREPAID`. Valid values: `m` (month), `d` (day).
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
}

func (r *InquirePriceCreateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceCreateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "GoodsNum")
	delete(f, "InstancePayMode")
	delete(f, "StoragePayMode")
	delete(f, "Cpu")
	delete(f, "Memory")
	delete(f, "Ccu")
	delete(f, "StorageLimit")
	delete(f, "TimeSpan")
	delete(f, "TimeUnit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceCreateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceCreateResponseParams struct {
	// Instance price
	InstancePrice *TradePrice `json:"InstancePrice,omitempty" name:"InstancePrice"`

	// Storage price
	StoragePrice *TradePrice `json:"StoragePrice,omitempty" name:"StoragePrice"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type InquirePriceCreateResponse struct {
	*tchttp.BaseResponse
	Response *InquirePriceCreateResponseParams `json:"Response"`
}

func (r *InquirePriceCreateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceCreateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceRenewRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Validity period, which needs to be used together with `TimeUnit`.
	TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// Unit of validity period, which needs to be used together with `TimeSpan`. Valid values: `d` (day), `m` (month).
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
}

type InquirePriceRenewRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Validity period, which needs to be used together with `TimeUnit`.
	TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// Unit of validity period, which needs to be used together with `TimeSpan`. Valid values: `d` (day), `m` (month).
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
}

func (r *InquirePriceRenewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceRenewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "TimeSpan")
	delete(f, "TimeUnit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceRenewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceRenewResponseParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Instance ID list
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Price of instance specification in array
	Prices []*TradePrice `json:"Prices,omitempty" name:"Prices"`

	// Total renewal price of compute node
	InstanceRealTotalPrice *int64 `json:"InstanceRealTotalPrice,omitempty" name:"InstanceRealTotalPrice"`

	// Total renewal price of storage node
	StorageRealTotalPrice *int64 `json:"StorageRealTotalPrice,omitempty" name:"StorageRealTotalPrice"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type InquirePriceRenewResponse struct {
	*tchttp.BaseResponse
	Response *InquirePriceRenewResponseParams `json:"Response"`
}

func (r *InquirePriceRenewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceRenewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceAuditRule struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Whether the audit is rule audit. Valid values: `true` (rule audit), `false` (full audit).
	// Note: This field may return null, indicating that no valid values can be obtained.
	AuditRule *bool `json:"AuditRule,omitempty" name:"AuditRule"`

	// Audit rule details, which is valid only when `AuditRule` is `true`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AuditRuleFilters []*AuditRuleFilters `json:"AuditRuleFilters,omitempty" name:"AuditRuleFilters"`
}

type InstanceInitInfo struct {
	// Instance CPU
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// Instance memory
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Instance type. Valid values:`rw`, `ro`.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Number of the instances. Value range: 1-15.
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`
}

type InstanceNetInfo struct {
	// Network type
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceGroupType *string `json:"InstanceGroupType,omitempty" name:"InstanceGroupType"`

	// Instance group ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceGroupId *string `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`

	// VPC ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Network type
	// Note: This field may return null, indicating that no valid values can be obtained.
	NetType *int64 `json:"NetType,omitempty" name:"NetType"`

	// VPC IP
	// Note: This field may return null, indicating that no valid values can be obtained.
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// VPC port
	// Note: This field may return null, indicating that no valid values can be obtained.
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// Public network domain name
	// Note: This field may return null, indicating that no valid values can be obtained.
	WanDomain *string `json:"WanDomain,omitempty" name:"WanDomain"`


	WanIP *string `json:"WanIP,omitempty" name:"WanIP"`

	// Public network port
	// Note: This field may return null, indicating that no valid values can be obtained.
	WanPort *int64 `json:"WanPort,omitempty" name:"WanPort"`

	// Public network status
	// Note: This field may return null, indicating that no valid values can be obtained.
	WanStatus *string `json:"WanStatus,omitempty" name:"WanStatus"`
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

	// Whether there is an inventory.
	HasStock *bool `json:"HasStock,omitempty" name:"HasStock"`

	// Machine type
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// Maximum IOPS
	MaxIops *int64 `json:"MaxIops,omitempty" name:"MaxIops"`

	// Maximum bandwidth
	MaxIoBandWidth *int64 `json:"MaxIoBandWidth,omitempty" name:"MaxIoBandWidth"`

	// Inventory information in a region
	// Note: This field may return null, indicating that no valid values can be obtained.
	ZoneStockInfos []*ZoneStockInfo `json:"ZoneStockInfos,omitempty" name:"ZoneStockInfos"`

	// Quantity in stock
	// Note: This field may return null, indicating that no valid values can be obtained.
	StockCount *int64 `json:"StockCount,omitempty" name:"StockCount"`
}

// Predefined struct for user
type IsolateClusterRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// This parameter has been disused.
	DbType *string `json:"DbType,omitempty" name:"DbType"`
}

type IsolateClusterRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// This parameter has been disused.
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

// Predefined struct for user
type IsolateClusterResponseParams struct {
	// Task flow ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// Refund order ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type IsolateClusterResponse struct {
	*tchttp.BaseResponse
	Response *IsolateClusterResponseParams `json:"Response"`
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

// Predefined struct for user
type IsolateInstanceRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Instance ID array
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList"`

	// This parameter has been disused.
	DbType *string `json:"DbType,omitempty" name:"DbType"`
}

type IsolateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Instance ID array
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList"`

	// This parameter has been disused.
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

// Predefined struct for user
type IsolateInstanceResponseParams struct {
	// Task flow ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// Order ID for isolated instance (prepaid instance)
	// Note: this field may return null, indicating that no valid values can be obtained.
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type IsolateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *IsolateInstanceResponseParams `json:"Response"`
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

type ModifiableInfo struct {

}

// Predefined struct for user
type ModifyAuditRuleTemplatesRequestParams struct {
	// Audit rule template ID
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitempty" name:"RuleTemplateIds"`

	// Audit rule after modification
	RuleFilters []*RuleFilters `json:"RuleFilters,omitempty" name:"RuleFilters"`

	// New name of the rule template
	RuleTemplateName *string `json:"RuleTemplateName,omitempty" name:"RuleTemplateName"`

	// New description of the rule template
	Description *string `json:"Description,omitempty" name:"Description"`
}

type ModifyAuditRuleTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// Audit rule template ID
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitempty" name:"RuleTemplateIds"`

	// Audit rule after modification
	RuleFilters []*RuleFilters `json:"RuleFilters,omitempty" name:"RuleFilters"`

	// New name of the rule template
	RuleTemplateName *string `json:"RuleTemplateName,omitempty" name:"RuleTemplateName"`

	// New description of the rule template
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyAuditRuleTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAuditRuleTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleTemplateIds")
	delete(f, "RuleFilters")
	delete(f, "RuleTemplateName")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAuditRuleTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAuditRuleTemplatesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAuditRuleTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAuditRuleTemplatesResponseParams `json:"Response"`
}

func (r *ModifyAuditRuleTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAuditRuleTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAuditServiceRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Log retention period
	LogExpireDay *uint64 `json:"LogExpireDay,omitempty" name:"LogExpireDay"`

	// Frequent log retention period
	HighLogExpireDay *uint64 `json:"HighLogExpireDay,omitempty" name:"HighLogExpireDay"`

	// The parameter used to change the audit rule of the instance to full audit
	AuditAll *bool `json:"AuditAll,omitempty" name:"AuditAll"`

	// Rule audit
	AuditRuleFilters []*AuditRuleFilters `json:"AuditRuleFilters,omitempty" name:"AuditRuleFilters"`

	// Rule template ID
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitempty" name:"RuleTemplateIds"`
}

type ModifyAuditServiceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Log retention period
	LogExpireDay *uint64 `json:"LogExpireDay,omitempty" name:"LogExpireDay"`

	// Frequent log retention period
	HighLogExpireDay *uint64 `json:"HighLogExpireDay,omitempty" name:"HighLogExpireDay"`

	// The parameter used to change the audit rule of the instance to full audit
	AuditAll *bool `json:"AuditAll,omitempty" name:"AuditAll"`

	// Rule audit
	AuditRuleFilters []*AuditRuleFilters `json:"AuditRuleFilters,omitempty" name:"AuditRuleFilters"`

	// Rule template ID
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitempty" name:"RuleTemplateIds"`
}

func (r *ModifyAuditServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAuditServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "LogExpireDay")
	delete(f, "HighLogExpireDay")
	delete(f, "AuditAll")
	delete(f, "AuditRuleFilters")
	delete(f, "RuleTemplateIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAuditServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAuditServiceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAuditServiceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAuditServiceResponseParams `json:"Response"`
}

func (r *ModifyAuditServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAuditServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupConfigRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Full backup start time. Value range: [0-24*3600]. For example, 0:00 AM, 1:00 AM, and 2:00 AM are represented by 0, 3600, and 7200, respectively
	BackupTimeBeg *uint64 `json:"BackupTimeBeg,omitempty" name:"BackupTimeBeg"`

	// Full backup end time. Value range: [0-24*3600]. For example, 0:00 AM, 1:00 AM, and 2:00 AM are represented by 0, 3600, and 7200, respectively.
	BackupTimeEnd *uint64 `json:"BackupTimeEnd,omitempty" name:"BackupTimeEnd"`

	// Backup retention period in seconds. Backups will be cleared after this period elapses. 7 days is represented by 3600*24*7 = 604800. Maximum value: 158112000.
	ReserveDuration *uint64 `json:"ReserveDuration,omitempty" name:"ReserveDuration"`

	// Backup frequency. It is an array of 7 elements corresponding to Monday through Sunday. full: full backup; increment: incremental backup. This parameter cannot be modified currently and doesn't need to be entered.
	BackupFreq []*string `json:"BackupFreq,omitempty" name:"BackupFreq"`

	// Backup mode. logic: logic backup; snapshot: snapshot backup. This parameter cannot be modified currently and doesn't need to be entered.
	BackupType *string `json:"BackupType,omitempty" name:"BackupType"`
}

type ModifyBackupConfigRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Full backup start time. Value range: [0-24*3600]. For example, 0:00 AM, 1:00 AM, and 2:00 AM are represented by 0, 3600, and 7200, respectively
	BackupTimeBeg *uint64 `json:"BackupTimeBeg,omitempty" name:"BackupTimeBeg"`

	// Full backup end time. Value range: [0-24*3600]. For example, 0:00 AM, 1:00 AM, and 2:00 AM are represented by 0, 3600, and 7200, respectively.
	BackupTimeEnd *uint64 `json:"BackupTimeEnd,omitempty" name:"BackupTimeEnd"`

	// Backup retention period in seconds. Backups will be cleared after this period elapses. 7 days is represented by 3600*24*7 = 604800. Maximum value: 158112000.
	ReserveDuration *uint64 `json:"ReserveDuration,omitempty" name:"ReserveDuration"`

	// Backup frequency. It is an array of 7 elements corresponding to Monday through Sunday. full: full backup; increment: incremental backup. This parameter cannot be modified currently and doesn't need to be entered.
	BackupFreq []*string `json:"BackupFreq,omitempty" name:"BackupFreq"`

	// Backup mode. logic: logic backup; snapshot: snapshot backup. This parameter cannot be modified currently and doesn't need to be entered.
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

// Predefined struct for user
type ModifyBackupConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyBackupConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBackupConfigResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyBackupNameRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Backup file ID
	BackupId *int64 `json:"BackupId,omitempty" name:"BackupId"`

	// Backup name, which can contain up to 60 characters.
	BackupName *string `json:"BackupName,omitempty" name:"BackupName"`
}

type ModifyBackupNameRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Backup file ID
	BackupId *int64 `json:"BackupId,omitempty" name:"BackupId"`

	// Backup name, which can contain up to 60 characters.
	BackupName *string `json:"BackupName,omitempty" name:"BackupName"`
}

func (r *ModifyBackupNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "BackupId")
	delete(f, "BackupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBackupNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupNameResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyBackupNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBackupNameResponseParams `json:"Response"`
}

func (r *ModifyBackupNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterNameRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Cluster name
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

type ModifyClusterNameRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Cluster name
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

func (r *ModifyClusterNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ClusterName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterNameResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyClusterNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterNameResponseParams `json:"Response"`
}

func (r *ModifyClusterNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterParamRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// List of the parameters to be modified. Each element in the list is a combination of `ParamName`, `CurrentValue`, and `OldValue`. `ParamName` is the parameter name; `CurrentValue` is the current value; `OldValue` is the old value that doesn’t need to be verified.
	ParamList []*ParamItem `json:"ParamList,omitempty" name:"ParamList"`

	// Valid values: `yes` (execute during maintenance time), `no` (execute now)
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitempty" name:"IsInMaintainPeriod"`
}

type ModifyClusterParamRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// List of the parameters to be modified. Each element in the list is a combination of `ParamName`, `CurrentValue`, and `OldValue`. `ParamName` is the parameter name; `CurrentValue` is the current value; `OldValue` is the old value that doesn’t need to be verified.
	ParamList []*ParamItem `json:"ParamList,omitempty" name:"ParamList"`

	// Valid values: `yes` (execute during maintenance time), `no` (execute now)
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitempty" name:"IsInMaintainPeriod"`
}

func (r *ModifyClusterParamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterParamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ParamList")
	delete(f, "IsInMaintainPeriod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterParamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterParamResponseParams struct {
	// Async request ID used to query the result
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyClusterParamResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterParamResponseParams `json:"Response"`
}

func (r *ModifyClusterParamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterParamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterSlaveZoneRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Old replica AZ
	OldSlaveZone *string `json:"OldSlaveZone,omitempty" name:"OldSlaveZone"`

	// New replica AZ
	NewSlaveZone *string `json:"NewSlaveZone,omitempty" name:"NewSlaveZone"`
}

type ModifyClusterSlaveZoneRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Old replica AZ
	OldSlaveZone *string `json:"OldSlaveZone,omitempty" name:"OldSlaveZone"`

	// New replica AZ
	NewSlaveZone *string `json:"NewSlaveZone,omitempty" name:"NewSlaveZone"`
}

func (r *ModifyClusterSlaveZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterSlaveZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "OldSlaveZone")
	delete(f, "NewSlaveZone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterSlaveZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterSlaveZoneResponseParams struct {
	// Async FlowId
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyClusterSlaveZoneResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterSlaveZoneResponseParams `json:"Response"`
}

func (r *ModifyClusterSlaveZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterSlaveZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSecurityGroupsRequestParams struct {
	// Instance group ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// List of IDs of security groups to be modified, which is an array of one or more security group IDs.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// AZ
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type ModifyDBInstanceSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Instance group ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// List of IDs of security groups to be modified, which is an array of one or more security group IDs.
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

// Predefined struct for user
type ModifyDBInstanceSecurityGroupsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDBInstanceSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceSecurityGroupsResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyInstanceNameRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

type ModifyInstanceNameRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

func (r *ModifyInstanceNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceNameResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyInstanceNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceNameResponseParams `json:"Response"`
}

func (r *ModifyInstanceNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMaintainPeriodConfigRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Maintenance start time in seconds. For example, 03:00 AM is represented by 10800
	MaintainStartTime *int64 `json:"MaintainStartTime,omitempty" name:"MaintainStartTime"`

	// Maintenance duration in seconds. For example, one hour is represented by 3600
	MaintainDuration *int64 `json:"MaintainDuration,omitempty" name:"MaintainDuration"`

	// Maintenance days of the week. Valid values: [Mon, Tue, Wed, Thu, Fri, Sat, Sun].
	MaintainWeekDays []*string `json:"MaintainWeekDays,omitempty" name:"MaintainWeekDays"`
}

type ModifyMaintainPeriodConfigRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Maintenance start time in seconds. For example, 03:00 AM is represented by 10800
	MaintainStartTime *int64 `json:"MaintainStartTime,omitempty" name:"MaintainStartTime"`

	// Maintenance duration in seconds. For example, one hour is represented by 3600
	MaintainDuration *int64 `json:"MaintainDuration,omitempty" name:"MaintainDuration"`

	// Maintenance days of the week. Valid values: [Mon, Tue, Wed, Thu, Fri, Sat, Sun].
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

// Predefined struct for user
type ModifyMaintainPeriodConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyMaintainPeriodConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMaintainPeriodConfigResponseParams `json:"Response"`
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

type ModifyParamItem struct {
	// Parameter name
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`

	// Current parameter value
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`

	// Old parameter value, which is used only in output parameters.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OldValue *string `json:"OldValue,omitempty" name:"OldValue"`
}

// Predefined struct for user
type ModifyVipVportRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Instance group ID
	InstanceGrpId *string `json:"InstanceGrpId,omitempty" name:"InstanceGrpId"`

	// Target IP to be modified
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Target port to be modified
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// Database type. Valid values: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// Valid hours of old IPs. If it is set to `0` hours, the IPs will be released immediately.
	OldIpReserveHours *int64 `json:"OldIpReserveHours,omitempty" name:"OldIpReserveHours"`
}

type ModifyVipVportRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Instance group ID
	InstanceGrpId *string `json:"InstanceGrpId,omitempty" name:"InstanceGrpId"`

	// Target IP to be modified
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Target port to be modified
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// Database type. Valid values: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// Valid hours of old IPs. If it is set to `0` hours, the IPs will be released immediately.
	OldIpReserveHours *int64 `json:"OldIpReserveHours,omitempty" name:"OldIpReserveHours"`
}

func (r *ModifyVipVportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVipVportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceGrpId")
	delete(f, "Vip")
	delete(f, "Vport")
	delete(f, "DbType")
	delete(f, "OldIpReserveHours")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVipVportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVipVportResponseParams struct {
	// Async task ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyVipVportResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVipVportResponseParams `json:"Response"`
}

func (r *ModifyVipVportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVipVportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Module struct {
	// Whether it is supported. Valid values: `yes`, `no`.
	IsDisable *string `json:"IsDisable,omitempty" name:"IsDisable"`

	// Module name
	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
}

type NetAddr struct {
	// Private network IP
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Private network port number
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// Public network domain name
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	WanDomain *string `json:"WanDomain,omitempty" name:"WanDomain"`

	// Public network port number
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	WanPort *int64 `json:"WanPort,omitempty" name:"WanPort"`

	// Network type. Valid values: `ro` (read-only), `rw` or `ha` (read-write)
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// Subnet ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// VPC ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// Description
	// Note: This field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Public IP
	// Note: This field may return null, indicating that no valid values can be obtained.
	WanIP *string `json:"WanIP,omitempty" name:"WanIP"`

	// Public network status
	// Note: This field may return null, indicating that no valid values can be obtained.
	WanStatus *string `json:"WanStatus,omitempty" name:"WanStatus"`
}

type NewAccount struct {
	// Account name, which can contain 1-16 letters, digits, and underscores. It must begin with a letter and end with a letter or digit.
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// Password, which can contain 8-64 characters.
	AccountPassword *string `json:"AccountPassword,omitempty" name:"AccountPassword"`

	// Host
	Host *string `json:"Host,omitempty" name:"Host"`

	// Description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Maximum number of user connections, which cannot be above 10,240.
	MaxUserConnections *int64 `json:"MaxUserConnections,omitempty" name:"MaxUserConnections"`
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

// Predefined struct for user
type OfflineClusterRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
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

// Predefined struct for user
type OfflineClusterResponseParams struct {
	// Task flow ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type OfflineClusterResponse struct {
	*tchttp.BaseResponse
	Response *OfflineClusterResponseParams `json:"Response"`
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

// Predefined struct for user
type OfflineInstanceRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Instance ID array
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList"`
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

// Predefined struct for user
type OfflineInstanceResponseParams struct {
	// Task flow ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type OfflineInstanceResponse struct {
	*tchttp.BaseResponse
	Response *OfflineInstanceResponseParams `json:"Response"`
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

type OldAddrInfo struct {
	// IP
	// Note: This field may return null, indicating that no valid values can be obtained.
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Port
	// Note: This field may return null, indicating that no valid values can be obtained.
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// Expected valid hours of old IPs
	// Note: This field may return null, indicating that no valid values can be obtained.
	ReturnTime *string `json:"ReturnTime,omitempty" name:"ReturnTime"`
}

// Predefined struct for user
type OpenAuditServiceRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Log retention period
	LogExpireDay *uint64 `json:"LogExpireDay,omitempty" name:"LogExpireDay"`

	// Frequent log retention period
	HighLogExpireDay *uint64 `json:"HighLogExpireDay,omitempty" name:"HighLogExpireDay"`

	// Audit rule. If both this parameter and `RuleTemplateIds` are left empty, full audit will be applied.
	AuditRuleFilters []*AuditRuleFilters `json:"AuditRuleFilters,omitempty" name:"AuditRuleFilters"`

	// Rule template ID. If both this parameter and `AuditRuleFilters` are left empty, full audit will be applied.
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitempty" name:"RuleTemplateIds"`
}

type OpenAuditServiceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Log retention period
	LogExpireDay *uint64 `json:"LogExpireDay,omitempty" name:"LogExpireDay"`

	// Frequent log retention period
	HighLogExpireDay *uint64 `json:"HighLogExpireDay,omitempty" name:"HighLogExpireDay"`

	// Audit rule. If both this parameter and `RuleTemplateIds` are left empty, full audit will be applied.
	AuditRuleFilters []*AuditRuleFilters `json:"AuditRuleFilters,omitempty" name:"AuditRuleFilters"`

	// Rule template ID. If both this parameter and `AuditRuleFilters` are left empty, full audit will be applied.
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitempty" name:"RuleTemplateIds"`
}

func (r *OpenAuditServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenAuditServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "LogExpireDay")
	delete(f, "HighLogExpireDay")
	delete(f, "AuditRuleFilters")
	delete(f, "RuleTemplateIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenAuditServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenAuditServiceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type OpenAuditServiceResponse struct {
	*tchttp.BaseResponse
	Response *OpenAuditServiceResponseParams `json:"Response"`
}

func (r *OpenAuditServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenAuditServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ParamInfo struct {
	// Current value
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`

	// Default value
	Default *string `json:"Default,omitempty" name:"Default"`

	// List of valid values when parameter type is `enum`, `string` or `bool`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EnumValue []*string `json:"EnumValue,omitempty" name:"EnumValue"`

	// Maximum value when parameter type is `float` or `integer`.
	Max *string `json:"Max,omitempty" name:"Max"`

	// Minimum value when parameter type is `float` or `integer`.
	Min *string `json:"Min,omitempty" name:"Min"`

	// Parameter name
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`

	// Whether to restart the instance for the modified parameters to take effect.
	NeedReboot *int64 `json:"NeedReboot,omitempty" name:"NeedReboot"`

	// Parameter type: `integer`, `float`, `string`, `enum`, `bool`.
	ParamType *string `json:"ParamType,omitempty" name:"ParamType"`

	// Match type. Regex can be used when parameter type is `string`. Valid value: `multiVal`.
	MatchType *string `json:"MatchType,omitempty" name:"MatchType"`

	// Match values, which will be separated by semicolon when match type is `multiVal`.
	MatchValue *string `json:"MatchValue,omitempty" name:"MatchValue"`

	// Parameter description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Whether it is global parameter
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsGlobal *int64 `json:"IsGlobal,omitempty" name:"IsGlobal"`

	// Whether the parameter can be modified
	// Note: This field may return null, indicating that no valid values can be obtained.
	ModifiableInfo *ModifiableInfo `json:"ModifiableInfo,omitempty" name:"ModifiableInfo"`

	// Whether it is a function
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsFunc *bool `json:"IsFunc,omitempty" name:"IsFunc"`

	// Function
	// Note: This field may return null, indicating that no valid values can be obtained.
	Func *string `json:"Func,omitempty" name:"Func"`
}

type ParamItem struct {
	// Parameter name
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`

	// New value
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`

	// Original value
	OldValue *string `json:"OldValue,omitempty" name:"OldValue"`
}

type ParamTemplateListInfo struct {
	// Parameter template ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// Parameter template name
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// Parameter template description
	TemplateDescription *string `json:"TemplateDescription,omitempty" name:"TemplateDescription"`

	// Engine version
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// Database Type. Valid values: `NORMAL`, `SERVERLESS`.
	DbMode *string `json:"DbMode,omitempty" name:"DbMode"`

	// Parameter template details
	// Note: This field may return null, indicating that no valid values can be obtained.
	ParamInfoSet []*TemplateParamInfo `json:"ParamInfoSet,omitempty" name:"ParamInfoSet"`
}

// Predefined struct for user
type PauseServerlessRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Whether to pause forcibly and ignore the current user connections. Valid values: `0` (no), `1` (yes). Default value: `1`
	ForcePause *int64 `json:"ForcePause,omitempty" name:"ForcePause"`
}

type PauseServerlessRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Whether to pause forcibly and ignore the current user connections. Valid values: `0` (no), `1` (yes). Default value: `1`
	ForcePause *int64 `json:"ForcePause,omitempty" name:"ForcePause"`
}

func (r *PauseServerlessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PauseServerlessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ForcePause")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PauseServerlessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PauseServerlessResponseParams struct {
	// Async task ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type PauseServerlessResponse struct {
	*tchttp.BaseResponse
	Response *PauseServerlessResponseParams `json:"Response"`
}

func (r *PauseServerlessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PauseServerlessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

	// Operator
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

// Predefined struct for user
type RemoveClusterSlaveZoneRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Replica AZ
	SlaveZone *string `json:"SlaveZone,omitempty" name:"SlaveZone"`
}

type RemoveClusterSlaveZoneRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Replica AZ
	SlaveZone *string `json:"SlaveZone,omitempty" name:"SlaveZone"`
}

func (r *RemoveClusterSlaveZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveClusterSlaveZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SlaveZone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveClusterSlaveZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveClusterSlaveZoneResponseParams struct {
	// Async FlowId
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RemoveClusterSlaveZoneResponse struct {
	*tchttp.BaseResponse
	Response *RemoveClusterSlaveZoneResponseParams `json:"Response"`
}

func (r *RemoveClusterSlaveZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveClusterSlaveZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetAccountPasswordRequestParams struct {
	// Database account name
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// New password of the database account
	AccountPassword *string `json:"AccountPassword,omitempty" name:"AccountPassword"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Host. Default value: `%`
	Host *string `json:"Host,omitempty" name:"Host"`
}

type ResetAccountPasswordRequest struct {
	*tchttp.BaseRequest
	
	// Database account name
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// New password of the database account
	AccountPassword *string `json:"AccountPassword,omitempty" name:"AccountPassword"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Host. Default value: `%`
	Host *string `json:"Host,omitempty" name:"Host"`
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
	delete(f, "AccountName")
	delete(f, "AccountPassword")
	delete(f, "ClusterId")
	delete(f, "Host")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetAccountPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetAccountPasswordResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ResetAccountPasswordResponse struct {
	*tchttp.BaseResponse
	Response *ResetAccountPasswordResponseParams `json:"Response"`
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

// Predefined struct for user
type RestartInstanceRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type RestartInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *RestartInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartInstanceResponseParams struct {
	// Async task ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RestartInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RestartInstanceResponseParams `json:"Response"`
}

func (r *RestartInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeServerlessRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type ResumeServerlessRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *ResumeServerlessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeServerlessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResumeServerlessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeServerlessResponseParams struct {
	// Async task ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ResumeServerlessResponse struct {
	*tchttp.BaseResponse
	Response *ResumeServerlessResponseParams `json:"Response"`
}

func (r *ResumeServerlessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeServerlessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RollbackTimeRange struct {
	// Start time
	TimeRangeStart *string `json:"TimeRangeStart,omitempty" name:"TimeRangeStart"`

	// End time
	TimeRangeEnd *string `json:"TimeRangeEnd,omitempty" name:"TimeRangeEnd"`
}

type RuleFilters struct {
	// Filter parameter name of the audit rule. Valid values: `host` (client IP), `user` (database account), `dbName` (database name), `sqlType` (SQL type), `sql` (SQL statement).
	Type *string `json:"Type,omitempty" name:"Type"`

	// Filter match type of the audit rule. Valid values: `INC` (including), `EXC` (excluding), `EQS` (equal to), `NEQ` (not equal to).
	Compare *string `json:"Compare,omitempty" name:"Compare"`

	// Filter match value of the audit rule
	Value []*string `json:"Value,omitempty" name:"Value"`
}

type SaleRegion struct {
	// Region name
	Region *string `json:"Region,omitempty" name:"Region"`

	// Numeric ID of a region
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// Region name
	RegionZh *string `json:"RegionZh,omitempty" name:"RegionZh"`

	// List of purchasable AZs
	ZoneSet []*SaleZone `json:"ZoneSet,omitempty" name:"ZoneSet"`

	// Engine type
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// Supported modules in a region
	Modules []*Module `json:"Modules,omitempty" name:"Modules"`
}

type SaleZone struct {
	// AZ name
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Numeric ID of an AZ
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// AZ name
	ZoneZh *string `json:"ZoneZh,omitempty" name:"ZoneZh"`

	// Whether serverless cluster is supported. Valid values: <br>
	// `0`: No<br>
	// `1`: Yes
	IsSupportServerless *int64 `json:"IsSupportServerless,omitempty" name:"IsSupportServerless"`

	// Whether standard cluster is supported. Valid values: <br>
	// `0`: No<br>
	// `1`: Yes
	IsSupportNormal *int64 `json:"IsSupportNormal,omitempty" name:"IsSupportNormal"`

	// Physical zone
	PhysicalZone *string `json:"PhysicalZone,omitempty" name:"PhysicalZone"`

	// Whether the user has AZ permission
	// Note: This field may return null, indicating that no valid values can be obtained.
	HasPermission *bool `json:"HasPermission,omitempty" name:"HasPermission"`

	// Whether it is a full-linkage RDMA AZ.
	IsWholeRdmaZone *string `json:"IsWholeRdmaZone,omitempty" name:"IsWholeRdmaZone"`
}

// Predefined struct for user
type SearchClusterDatabasesRequestParams struct {
	// The cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Database name
	Database *string `json:"Database,omitempty" name:"Database"`

	// Whether to search exactly
	// Valid values: `0` (fuzzy search), `1` (exact search). 
	// Default value: `0`.
	MatchType *int64 `json:"MatchType,omitempty" name:"MatchType"`
}

type SearchClusterDatabasesRequest struct {
	*tchttp.BaseRequest
	
	// The cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Database name
	Database *string `json:"Database,omitempty" name:"Database"`

	// Whether to search exactly
	// Valid values: `0` (fuzzy search), `1` (exact search). 
	// Default value: `0`.
	MatchType *int64 `json:"MatchType,omitempty" name:"MatchType"`
}

func (r *SearchClusterDatabasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchClusterDatabasesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Database")
	delete(f, "MatchType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchClusterDatabasesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchClusterDatabasesResponseParams struct {
	// Database List
	// Note: This field may return null, indicating that no valid values can be obtained.
	Databases []*string `json:"Databases,omitempty" name:"Databases"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SearchClusterDatabasesResponse struct {
	*tchttp.BaseResponse
	Response *SearchClusterDatabasesResponseParams `json:"Response"`
}

func (r *SearchClusterDatabasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchClusterDatabasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchClusterTablesRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Database name
	Database *string `json:"Database,omitempty" name:"Database"`

	// Data table name
	Table *string `json:"Table,omitempty" name:"Table"`

	// Data table type. Valid values:
	// `view`: Only return to view,
	// `base_table`: Only return to basic table,
	// `all`: Return to view and table.
	TableType *string `json:"TableType,omitempty" name:"TableType"`
}

type SearchClusterTablesRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Database name
	Database *string `json:"Database,omitempty" name:"Database"`

	// Data table name
	Table *string `json:"Table,omitempty" name:"Table"`

	// Data table type. Valid values:
	// `view`: Only return to view,
	// `base_table`: Only return to basic table,
	// `all`: Return to view and table.
	TableType *string `json:"TableType,omitempty" name:"TableType"`
}

func (r *SearchClusterTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchClusterTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Database")
	delete(f, "Table")
	delete(f, "TableType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchClusterTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchClusterTablesResponseParams struct {
	// Data table list
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tables []*DatabaseTables `json:"Tables,omitempty" name:"Tables"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SearchClusterTablesResponse struct {
	*tchttp.BaseResponse
	Response *SearchClusterTablesResponseParams `json:"Response"`
}

func (r *SearchClusterTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchClusterTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

// Predefined struct for user
type SetRenewFlagRequestParams struct {
	// ID of the instance to be manipulated
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// Auto-renewal flag. 0: normal renewal, 1: auto-renewal, 2: no renewal.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
}

type SetRenewFlagRequest struct {
	*tchttp.BaseRequest
	
	// ID of the instance to be manipulated
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// Auto-renewal flag. 0: normal renewal, 1: auto-renewal, 2: no renewal.
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

// Predefined struct for user
type SetRenewFlagResponseParams struct {
	// Number of successfully manipulated instances
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *SetRenewFlagResponseParams `json:"Response"`
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

type SlowQueriesItem struct {
	// Execution timestamp
	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// Execution duration in seconds
	QueryTime *float64 `json:"QueryTime,omitempty" name:"QueryTime"`

	// SQL statement
	SqlText *string `json:"SqlText,omitempty" name:"SqlText"`

	// Client host
	UserHost *string `json:"UserHost,omitempty" name:"UserHost"`

	// Username
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Database name
	Database *string `json:"Database,omitempty" name:"Database"`

	// Lock duration in seconds
	LockTime *float64 `json:"LockTime,omitempty" name:"LockTime"`

	// Number of scanned rows
	RowsExamined *int64 `json:"RowsExamined,omitempty" name:"RowsExamined"`

	// Number of returned rows
	RowsSent *int64 `json:"RowsSent,omitempty" name:"RowsSent"`

	// SQL template
	SqlTemplate *string `json:"SqlTemplate,omitempty" name:"SqlTemplate"`

	// MD5 value of the SQL statement
	SqlMd5 *string `json:"SqlMd5,omitempty" name:"SqlMd5"`
}

// Predefined struct for user
type SwitchClusterVpcRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// VPC ID in string
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// Subnet ID in string
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// Valid hours of old IP
	OldIpReserveHours *int64 `json:"OldIpReserveHours,omitempty" name:"OldIpReserveHours"`
}

type SwitchClusterVpcRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// VPC ID in string
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// Subnet ID in string
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// Valid hours of old IP
	OldIpReserveHours *int64 `json:"OldIpReserveHours,omitempty" name:"OldIpReserveHours"`
}

func (r *SwitchClusterVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchClusterVpcRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "UniqVpcId")
	delete(f, "UniqSubnetId")
	delete(f, "OldIpReserveHours")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchClusterVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchClusterVpcResponseParams struct {
	// Async task ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SwitchClusterVpcResponse struct {
	*tchttp.BaseResponse
	Response *SwitchClusterVpcResponseParams `json:"Response"`
}

func (r *SwitchClusterVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchClusterVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchClusterZoneRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The current AZ
	OldZone *string `json:"OldZone,omitempty" name:"OldZone"`

	// New AZ
	NewZone *string `json:"NewZone,omitempty" name:"NewZone"`

	// Valid values: `yes` (execute during maintenance time), `no` (execute now)
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitempty" name:"IsInMaintainPeriod"`
}

type SwitchClusterZoneRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The current AZ
	OldZone *string `json:"OldZone,omitempty" name:"OldZone"`

	// New AZ
	NewZone *string `json:"NewZone,omitempty" name:"NewZone"`

	// Valid values: `yes` (execute during maintenance time), `no` (execute now)
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitempty" name:"IsInMaintainPeriod"`
}

func (r *SwitchClusterZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchClusterZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "OldZone")
	delete(f, "NewZone")
	delete(f, "IsInMaintainPeriod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchClusterZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchClusterZoneResponseParams struct {
	// Async FlowId
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SwitchClusterZoneResponse struct {
	*tchttp.BaseResponse
	Response *SwitchClusterZoneResponseParams `json:"Response"`
}

func (r *SwitchClusterZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchClusterZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchProxyVpcRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// VPC ID in string
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// Subnet ID in string
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// Valid hours of old IP
	OldIpReserveHours *int64 `json:"OldIpReserveHours,omitempty" name:"OldIpReserveHours"`

	// Database proxy group ID (required), which can be obtained through the `DescribeProxies` API.
	ProxyGroupId *string `json:"ProxyGroupId,omitempty" name:"ProxyGroupId"`
}

type SwitchProxyVpcRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// VPC ID in string
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// Subnet ID in string
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// Valid hours of old IP
	OldIpReserveHours *int64 `json:"OldIpReserveHours,omitempty" name:"OldIpReserveHours"`

	// Database proxy group ID (required), which can be obtained through the `DescribeProxies` API.
	ProxyGroupId *string `json:"ProxyGroupId,omitempty" name:"ProxyGroupId"`
}

func (r *SwitchProxyVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchProxyVpcRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "UniqVpcId")
	delete(f, "UniqSubnetId")
	delete(f, "OldIpReserveHours")
	delete(f, "ProxyGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchProxyVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchProxyVpcResponseParams struct {
	// Async task ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SwitchProxyVpcResponse struct {
	*tchttp.BaseResponse
	Response *SwitchProxyVpcResponseParams `json:"Response"`
}

func (r *SwitchProxyVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchProxyVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {
	// Tag key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type TemplateParamInfo struct {
	// Current value
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`

	// Default value
	Default *string `json:"Default,omitempty" name:"Default"`

	// The collection of valid value types when parameter type is `enum`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EnumValue []*string `json:"EnumValue,omitempty" name:"EnumValue"`

	// Maximum value when parameter type is `float` or `integer`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Max *string `json:"Max,omitempty" name:"Max"`

	// Minimum value when parameter type is `float` or `integer`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Min *string `json:"Min,omitempty" name:"Min"`

	// Parameter name
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`

	// Whether to restart the instance for the parameter to take effect
	NeedReboot *int64 `json:"NeedReboot,omitempty" name:"NeedReboot"`

	// Parameter description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Parameter type. Valid value: `integer`, `float`, `string`, `enum`.
	ParamType *string `json:"ParamType,omitempty" name:"ParamType"`
}

type TradePrice struct {
	// The non-discounted total price of monthly subscribed resources (unit: US cent)
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalPrice *int64 `json:"TotalPrice,omitempty" name:"TotalPrice"`

	// Total discount. `100` means no discount.
	Discount *float64 `json:"Discount,omitempty" name:"Discount"`

	// The discounted total price of monthly subscribed resources (unit: US cent). If a discount is applied, `TotalPriceDiscount` will be the product of `TotalPrice` and `Discount`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalPriceDiscount *int64 `json:"TotalPriceDiscount,omitempty" name:"TotalPriceDiscount"`

	// The non-discounted unit price of pay-as-you-go resources (unit: US cent)
	// Note: This field may return null, indicating that no valid values can be obtained.
	UnitPrice *int64 `json:"UnitPrice,omitempty" name:"UnitPrice"`

	// The discounted unit price of pay-as-you-go resources (unit: US cent). If a discount is applied, `UnitPriceDiscount` will be the product of `UnitPrice` and `Discount`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UnitPriceDiscount *int64 `json:"UnitPriceDiscount,omitempty" name:"UnitPriceDiscount"`

	// Price unit
	ChargeUnit *string `json:"ChargeUnit,omitempty" name:"ChargeUnit"`
}

// Predefined struct for user
type UpgradeInstanceRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Database CPU
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// Database memory in GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Upgrade type. Valid values: upgradeImmediate, upgradeInMaintain
	UpgradeType *string `json:"UpgradeType,omitempty" name:"UpgradeType"`

	// This parameter has been disused.
	StorageLimit *uint64 `json:"StorageLimit,omitempty" name:"StorageLimit"`

	// Whether to automatically select a voucher. 1: yes; 0: no. Default value: 0
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// This parameter has been disused.
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// Transaction mode. Valid values: `0` (place and pay for an order), `1` (place an order)
	DealMode *int64 `json:"DealMode,omitempty" name:"DealMode"`

	// Valid values: `NormalUpgrade` (Normal mode), `FastUpgrade` (QuickChange). If the system detects that the configuration modification process will cause a momentary disconnection, the process will be terminated.
	UpgradeMode *string `json:"UpgradeMode,omitempty" name:"UpgradeMode"`
}

type UpgradeInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Database CPU
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// Database memory in GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Upgrade type. Valid values: upgradeImmediate, upgradeInMaintain
	UpgradeType *string `json:"UpgradeType,omitempty" name:"UpgradeType"`

	// This parameter has been disused.
	StorageLimit *uint64 `json:"StorageLimit,omitempty" name:"StorageLimit"`

	// Whether to automatically select a voucher. 1: yes; 0: no. Default value: 0
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// This parameter has been disused.
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// Transaction mode. Valid values: `0` (place and pay for an order), `1` (place an order)
	DealMode *int64 `json:"DealMode,omitempty" name:"DealMode"`

	// Valid values: `NormalUpgrade` (Normal mode), `FastUpgrade` (QuickChange). If the system detects that the configuration modification process will cause a momentary disconnection, the process will be terminated.
	UpgradeMode *string `json:"UpgradeMode,omitempty" name:"UpgradeMode"`
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
	delete(f, "UpgradeMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeInstanceResponseParams struct {
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
}

type UpgradeInstanceResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeInstanceResponseParams `json:"Response"`
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

type ZoneStockInfo struct {
	// AZ
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Whether there is an inventory.
	HasStock *bool `json:"HasStock,omitempty" name:"HasStock"`

	// Quantity in stock
	StockCount *int64 `json:"StockCount,omitempty" name:"StockCount"`
}