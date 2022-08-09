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

package v20190823

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type ApplyResult struct {
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Application type
	ApplicationType *int64 `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// Status. Valid values: `0` (pending approval), `1` (application approved and task submitted), `2` (rejected)
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	ApplicationStatus *int64 `json:"ApplicationStatus,omitempty" name:"ApplicationStatus"`

	// ID of the submitted task
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Error information
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	Error *ErrorInfo `json:"Error,omitempty" name:"Error"`
}

type ApplyStatus struct {
	// Value format: cluster ID-application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Status. Valid values: `-1` (canceled), `0` (pending approval), `1` (application approved and task submitted), `2` (rejected). Only applications in the pending approval status can be updated.
	ApplicationStatus *int64 `json:"ApplicationStatus,omitempty" name:"ApplicationStatus"`

	// Application type
	ApplicationType *int64 `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

// Predefined struct for user
type ClearTablesRequestParams struct {
	// ID of the cluster instance where a table resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// List of information of tables to be cleared
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

type ClearTablesRequest struct {
	*tchttp.BaseRequest
	
	// ID of the cluster instance where a table resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// List of information of tables to be cleared
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

func (r *ClearTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClearTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SelectedTables")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ClearTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ClearTablesResponseParams struct {
	// Number of cleared tables
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of table clearing results
	TableResults []*TableResultNew `json:"TableResults,omitempty" name:"TableResults"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ClearTablesResponse struct {
	*tchttp.BaseResponse
	Response *ClearTablesResponseParams `json:"Response"`
}

func (r *ClearTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClearTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterInfo struct {
	// Cluster name
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Cluster region
	Region *string `json:"Region,omitempty" name:"Region"`

	// Cluster data description language type, such as `PROTO`, `TDR`, or `MIX`
	IdlType *string `json:"IdlType,omitempty" name:"IdlType"`

	// Network type
	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`

	// ID of the VPC instance with which a cluster is associated
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// ID of the subnet instance with which a cluster is associated
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Creation time
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Cluster password
	Password *string `json:"Password,omitempty" name:"Password"`

	// Password status
	PasswordStatus *string `json:"PasswordStatus,omitempty" name:"PasswordStatus"`

	// TcaplusDB SDK connection parameter: access ID
	ApiAccessId *string `json:"ApiAccessId,omitempty" name:"ApiAccessId"`

	// TcaplusDB SDK connection parameter: access address
	ApiAccessIp *string `json:"ApiAccessIp,omitempty" name:"ApiAccessIp"`

	// TcaplusDB SDK connection parameter: access port
	ApiAccessPort *int64 `json:"ApiAccessPort,omitempty" name:"ApiAccessPort"`

	// If `PasswordStatus` is `unmodifiable`, the old password has not expired, and this field will display its expiration time; otherwise, this field will be empty
	// Note: this field may return null, indicating that no valid values can be obtained.
	OldPasswordExpireTime *string `json:"OldPasswordExpireTime,omitempty" name:"OldPasswordExpireTime"`

	// TcaplusDB SDK connection parameter for accessing IPv6 addresses
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiAccessIpv6 *string `json:"ApiAccessIpv6,omitempty" name:"ApiAccessIpv6"`

	// Cluster type
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ClusterType *int64 `json:"ClusterType,omitempty" name:"ClusterType"`

	// Cluster status
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ClusterStatus *int64 `json:"ClusterStatus,omitempty" name:"ClusterStatus"`

	// Read CU
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ReadCapacityUnit *int64 `json:"ReadCapacityUnit,omitempty" name:"ReadCapacityUnit"`

	// Write CU
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	WriteCapacityUnit *int64 `json:"WriteCapacityUnit,omitempty" name:"WriteCapacityUnit"`

	// Disk capacity
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DiskVolume *int64 `json:"DiskVolume,omitempty" name:"DiskVolume"`

	// Information of the machine at the storage layer (tcapsvr) in a dedicated cluster
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ServerList []*ServerDetailInfo `json:"ServerList,omitempty" name:"ServerList"`

	// Information of the machine at the access layer (tcaproxy) in a dedicated cluster
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ProxyList []*ProxyDetailInfo `json:"ProxyList,omitempty" name:"ProxyList"`

	// Whether the cluster operation approval feature is enabled. Valid values: `0` (disabled), `1` (enabled)
	Censorship *int64 `json:"Censorship,omitempty" name:"Censorship"`

	// Approver UIN list
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	DbaUins []*string `json:"DbaUins,omitempty" name:"DbaUins"`

	// Whether data subscription is enabled
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DataFlowStatus *int64 `json:"DataFlowStatus,omitempty" name:"DataFlowStatus"`

	// CKafka information when data subscription is enabled
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	KafkaInfo *KafkaInfo `json:"KafkaInfo,omitempty" name:"KafkaInfo"`

	// The number of days after which the cluster Txh backup file will expire and be deleted.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	TxhBackupExpireDay *uint64 `json:"TxhBackupExpireDay,omitempty" name:"TxhBackupExpireDay"`

	// The number of days after which the cluster Ulog backup file will expire and be deleted.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	UlogBackupExpireDay *uint64 `json:"UlogBackupExpireDay,omitempty" name:"UlogBackupExpireDay"`

	// Whether the expiration policy of cluster Ulog backup file is read-only. `0`: Yes; `1`: No.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	IsReadOnlyUlogBackupExpireDay *uint64 `json:"IsReadOnlyUlogBackupExpireDay,omitempty" name:"IsReadOnlyUlogBackupExpireDay"`
}

// Predefined struct for user
type CompareIdlFilesRequestParams struct {
	// ID of the cluster where the table to be modified resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// List of tables to be modified
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`

	// Selected list of uploaded IDL files. Either this parameter or `NewIdlFiles` must be selected
	ExistingIdlFiles []*IdlFileInfo `json:"ExistingIdlFiles,omitempty" name:"ExistingIdlFiles"`

	// List of IDL files to be uploaded. Either this parameter or `ExistingIdlFiles` must be selected
	NewIdlFiles []*IdlFileInfo `json:"NewIdlFiles,omitempty" name:"NewIdlFiles"`
}

type CompareIdlFilesRequest struct {
	*tchttp.BaseRequest
	
	// ID of the cluster where the table to be modified resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// List of tables to be modified
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`

	// Selected list of uploaded IDL files. Either this parameter or `NewIdlFiles` must be selected
	ExistingIdlFiles []*IdlFileInfo `json:"ExistingIdlFiles,omitempty" name:"ExistingIdlFiles"`

	// List of IDL files to be uploaded. Either this parameter or `ExistingIdlFiles` must be selected
	NewIdlFiles []*IdlFileInfo `json:"NewIdlFiles,omitempty" name:"NewIdlFiles"`
}

func (r *CompareIdlFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CompareIdlFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SelectedTables")
	delete(f, "ExistingIdlFiles")
	delete(f, "NewIdlFiles")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CompareIdlFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CompareIdlFilesResponseParams struct {
	// Information list of all IDL files uploaded and verified in this request
	IdlFiles []*IdlFileInfo `json:"IdlFiles,omitempty" name:"IdlFiles"`

	// Number of tables verified to be valid in this request
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Verification result parsed from the selected table after the IDL description file is read
	TableInfos []*ParsedTableInfoNew `json:"TableInfos,omitempty" name:"TableInfos"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CompareIdlFilesResponse struct {
	*tchttp.BaseResponse
	Response *CompareIdlFilesResponseParams `json:"Response"`
}

func (r *CompareIdlFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CompareIdlFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CompareTablesInfo struct {
	// Cluster ID of the source table
	SrcTableClusterId *string `json:"SrcTableClusterId,omitempty" name:"SrcTableClusterId"`

	// Table group ID of the source table
	SrcTableGroupId *string `json:"SrcTableGroupId,omitempty" name:"SrcTableGroupId"`

	// Source table name
	SrcTableName *string `json:"SrcTableName,omitempty" name:"SrcTableName"`

	// Cluster ID of the target table
	DstTableClusterId *string `json:"DstTableClusterId,omitempty" name:"DstTableClusterId"`

	// Table group ID of the target table
	DstTableGroupId *string `json:"DstTableGroupId,omitempty" name:"DstTableGroupId"`

	// Target table name
	DstTableName *string `json:"DstTableName,omitempty" name:"DstTableName"`

	// Source table ID
	SrcTableInstanceId *string `json:"SrcTableInstanceId,omitempty" name:"SrcTableInstanceId"`

	// Target table ID
	DstTableInstanceId *string `json:"DstTableInstanceId,omitempty" name:"DstTableInstanceId"`
}

// Predefined struct for user
type CreateBackupRequestParams struct {
	// ID of the cluster where the table to be backed up resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Information list of tables to be backed up
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`

	// Remarks
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type CreateBackupRequest struct {
	*tchttp.BaseRequest
	
	// ID of the cluster where the table to be backed up resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Information list of tables to be backed up
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`

	// Remarks
	Remark *string `json:"Remark,omitempty" name:"Remark"`
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
	delete(f, "SelectedTables")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBackupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBackupResponseParams struct {
	// List of backup creation task IDs
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds"`

	// List of backup creation application IDs
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	ApplicationIds []*string `json:"ApplicationIds,omitempty" name:"ApplicationIds"`

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
type CreateClusterRequestParams struct {
	// Cluster data description language type, such as `PROTO`, `TDR`, or `MIX`
	IdlType *string `json:"IdlType,omitempty" name:"IdlType"`

	// Cluster name, which can contain up to 32 letters and digits
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// ID of the VPC instance bound to a cluster in the format of `vpc-f49l6u0z`
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// ID of the subnet instance bound to a cluster in the format of `subnet-pxir56ns`
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Cluster access password, which must contain lowercase letters (a-z), uppercase letters (A-Z), and digits (0-9).
	Password *string `json:"Password,omitempty" name:"Password"`

	// Cluster tag list
	ResourceTags []*TagInfoUnit `json:"ResourceTags,omitempty" name:"ResourceTags"`

	// Whether to enable IPv6 address access for clusters
	Ipv6Enable *int64 `json:"Ipv6Enable,omitempty" name:"Ipv6Enable"`

	// Information of the machine at the storage layer (tcapsvr) in a dedicated cluster
	ServerList []*MachineInfo `json:"ServerList,omitempty" name:"ServerList"`

	// Information of the machine at the access layer (tcaproxy) in a dedicated cluster
	ProxyList []*MachineInfo `json:"ProxyList,omitempty" name:"ProxyList"`

	// Cluster type. Valid values: `1` (standard), `2` (dedicated)
	ClusterType *int64 `json:"ClusterType,omitempty" name:"ClusterType"`

	// Authentication type. Valid values: `0` (static password), `1` (signature)
	AuthType *int64 `json:"AuthType,omitempty" name:"AuthType"`
}

type CreateClusterRequest struct {
	*tchttp.BaseRequest
	
	// Cluster data description language type, such as `PROTO`, `TDR`, or `MIX`
	IdlType *string `json:"IdlType,omitempty" name:"IdlType"`

	// Cluster name, which can contain up to 32 letters and digits
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// ID of the VPC instance bound to a cluster in the format of `vpc-f49l6u0z`
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// ID of the subnet instance bound to a cluster in the format of `subnet-pxir56ns`
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Cluster access password, which must contain lowercase letters (a-z), uppercase letters (A-Z), and digits (0-9).
	Password *string `json:"Password,omitempty" name:"Password"`

	// Cluster tag list
	ResourceTags []*TagInfoUnit `json:"ResourceTags,omitempty" name:"ResourceTags"`

	// Whether to enable IPv6 address access for clusters
	Ipv6Enable *int64 `json:"Ipv6Enable,omitempty" name:"Ipv6Enable"`

	// Information of the machine at the storage layer (tcapsvr) in a dedicated cluster
	ServerList []*MachineInfo `json:"ServerList,omitempty" name:"ServerList"`

	// Information of the machine at the access layer (tcaproxy) in a dedicated cluster
	ProxyList []*MachineInfo `json:"ProxyList,omitempty" name:"ProxyList"`

	// Cluster type. Valid values: `1` (standard), `2` (dedicated)
	ClusterType *int64 `json:"ClusterType,omitempty" name:"ClusterType"`

	// Authentication type. Valid values: `0` (static password), `1` (signature)
	AuthType *int64 `json:"AuthType,omitempty" name:"AuthType"`
}

func (r *CreateClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdlType")
	delete(f, "ClusterName")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "Password")
	delete(f, "ResourceTags")
	delete(f, "Ipv6Enable")
	delete(f, "ServerList")
	delete(f, "ProxyList")
	delete(f, "ClusterType")
	delete(f, "AuthType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterResponseParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateClusterResponse struct {
	*tchttp.BaseResponse
	Response *CreateClusterResponseParams `json:"Response"`
}

func (r *CreateClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSnapshotsRequestParams struct {
	// The ID of the cluster where the table resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Snapshot list
	SelectedTables []*SnapshotInfo `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

type CreateSnapshotsRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the cluster where the table resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Snapshot list
	SelectedTables []*SnapshotInfo `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

func (r *CreateSnapshotsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSnapshotsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SelectedTables")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSnapshotsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSnapshotsResponseParams struct {
	// The number of snapshots created in batches
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The result list of snapshots created in batches
	TableResults []*SnapshotResult `json:"TableResults,omitempty" name:"TableResults"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSnapshotsResponse struct {
	*tchttp.BaseResponse
	Response *CreateSnapshotsResponseParams `json:"Response"`
}

func (r *CreateSnapshotsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSnapshotsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTableGroupRequestParams struct {
	// ID of the cluster where a table group resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Table group name, which can contain up to 32 letters and digits
	TableGroupName *string `json:"TableGroupName,omitempty" name:"TableGroupName"`

	// Table group ID, which can be customized but must be unique in one cluster. If it is not specified, the auto-increment mode will be used.
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// Table group tag list
	ResourceTags []*TagInfoUnit `json:"ResourceTags,omitempty" name:"ResourceTags"`
}

type CreateTableGroupRequest struct {
	*tchttp.BaseRequest
	
	// ID of the cluster where a table group resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Table group name, which can contain up to 32 letters and digits
	TableGroupName *string `json:"TableGroupName,omitempty" name:"TableGroupName"`

	// Table group ID, which can be customized but must be unique in one cluster. If it is not specified, the auto-increment mode will be used.
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// Table group tag list
	ResourceTags []*TagInfoUnit `json:"ResourceTags,omitempty" name:"ResourceTags"`
}

func (r *CreateTableGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTableGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "TableGroupName")
	delete(f, "TableGroupId")
	delete(f, "ResourceTags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTableGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTableGroupResponseParams struct {
	// ID of table group successfully created
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateTableGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateTableGroupResponseParams `json:"Response"`
}

func (r *CreateTableGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTableGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTablesRequestParams struct {
	// ID of the cluster where to create a table
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Table creation IDL file list selected by user
	IdlFiles []*IdlFileInfo `json:"IdlFiles,omitempty" name:"IdlFiles"`

	// Information list of tables to be created
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`

	// Table tag list
	ResourceTags []*TagInfoUnit `json:"ResourceTags,omitempty" name:"ResourceTags"`
}

type CreateTablesRequest struct {
	*tchttp.BaseRequest
	
	// ID of the cluster where to create a table
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Table creation IDL file list selected by user
	IdlFiles []*IdlFileInfo `json:"IdlFiles,omitempty" name:"IdlFiles"`

	// Information list of tables to be created
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`

	// Table tag list
	ResourceTags []*TagInfoUnit `json:"ResourceTags,omitempty" name:"ResourceTags"`
}

func (r *CreateTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "IdlFiles")
	delete(f, "SelectedTables")
	delete(f, "ResourceTags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTablesResponseParams struct {
	// Number of tables created in batches
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of tables created in batches
	TableResults []*TableResultNew `json:"TableResults,omitempty" name:"TableResults"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateTablesResponse struct {
	*tchttp.BaseResponse
	Response *CreateTablesResponseParams `json:"Response"`
}

func (r *CreateTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterRequestParams struct {
	// ID of cluster to be deleted
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type DeleteClusterRequest struct {
	*tchttp.BaseRequest
	
	// ID of cluster to be deleted
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DeleteClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterResponseParams struct {
	// Task ID generated by cluster deletion
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteClusterResponse struct {
	*tchttp.BaseResponse
	Response *DeleteClusterResponseParams `json:"Response"`
}

func (r *DeleteClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIdlFilesRequestParams struct {
	// ID of the cluster where IDL resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// List of information of IDL files to be deleted
	IdlFiles []*IdlFileInfo `json:"IdlFiles,omitempty" name:"IdlFiles"`
}

type DeleteIdlFilesRequest struct {
	*tchttp.BaseRequest
	
	// ID of the cluster where IDL resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// List of information of IDL files to be deleted
	IdlFiles []*IdlFileInfo `json:"IdlFiles,omitempty" name:"IdlFiles"`
}

func (r *DeleteIdlFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIdlFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "IdlFiles")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIdlFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIdlFilesResponseParams struct {
	// Number of returned results
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Deletion result
	IdlFileInfos []*IdlFileInfoWithoutContent `json:"IdlFileInfos,omitempty" name:"IdlFileInfos"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteIdlFilesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteIdlFilesResponseParams `json:"Response"`
}

func (r *DeleteIdlFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIdlFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSnapshotsRequestParams struct {
	// The ID of the cluster where the table resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The list of snapshots to delete
	SelectedTables []*SnapshotInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

type DeleteSnapshotsRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the cluster where the table resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The list of snapshots to delete
	SelectedTables []*SnapshotInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

func (r *DeleteSnapshotsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSnapshotsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SelectedTables")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSnapshotsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSnapshotsResponseParams struct {
	// The number of snapshots deleted in batches
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The result list of snapshots deleted in batches
	TableResults []*SnapshotResult `json:"TableResults,omitempty" name:"TableResults"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteSnapshotsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSnapshotsResponseParams `json:"Response"`
}

func (r *DeleteSnapshotsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSnapshotsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTableDataFlowRequestParams struct {
	// The ID of the cluster where the tables reside
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The list of tables for which data subscription will be disabled
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

type DeleteTableDataFlowRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the cluster where the tables reside
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The list of tables for which data subscription will be disabled
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

func (r *DeleteTableDataFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTableDataFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SelectedTables")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTableDataFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTableDataFlowResponseParams struct {
	// The number of tables for which data subscription has been disabled
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The result list of tables for which data subscription has been disabled
	TableResults []*TableResultNew `json:"TableResults,omitempty" name:"TableResults"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteTableDataFlowResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTableDataFlowResponseParams `json:"Response"`
}

func (r *DeleteTableDataFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTableDataFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTableGroupRequestParams struct {
	// ID of the cluster where a table group resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Table group ID
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`
}

type DeleteTableGroupRequest struct {
	*tchttp.BaseRequest
	
	// ID of the cluster where a table group resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Table group ID
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`
}

func (r *DeleteTableGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTableGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "TableGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTableGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTableGroupResponseParams struct {
	// Task ID generated by table group deletion
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteTableGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTableGroupResponseParams `json:"Response"`
}

func (r *DeleteTableGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTableGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTableIndexRequestParams struct {
	// ID of the cluster where the table resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The list of tables whose global indexes need to be deleted
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

type DeleteTableIndexRequest struct {
	*tchttp.BaseRequest
	
	// ID of the cluster where the table resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The list of tables whose global indexes need to be deleted
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

func (r *DeleteTableIndexRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTableIndexRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SelectedTables")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTableIndexRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTableIndexResponseParams struct {
	// The number of tables whose global indexes are deleted
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The list of global index deletion results
	TableResults []*TableResultNew `json:"TableResults,omitempty" name:"TableResults"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteTableIndexResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTableIndexResponseParams `json:"Response"`
}

func (r *DeleteTableIndexResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTableIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTablesRequestParams struct {
	// ID of the cluster where the table to be dropped resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// List of information of tables to be dropped
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

type DeleteTablesRequest struct {
	*tchttp.BaseRequest
	
	// ID of the cluster where the table to be dropped resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// List of information of tables to be dropped
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

func (r *DeleteTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SelectedTables")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTablesResponseParams struct {
	// Number of dropped tables
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of details of dropped tables
	TableResults []*TableResultNew `json:"TableResults,omitempty" name:"TableResults"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteTablesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTablesResponseParams `json:"Response"`
}

func (r *DeleteTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterTagsRequestParams struct {
	// The list of cluster IDs
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`
}

type DescribeClusterTagsRequest struct {
	*tchttp.BaseRequest
	
	// The list of cluster IDs
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`
}

func (r *DescribeClusterTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterTagsResponseParams struct {
	// The information list of cluster tags
	Rows []*TagsInfoOfCluster `json:"Rows,omitempty" name:"Rows"`

	// The number of returned results
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClusterTagsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterTagsResponseParams `json:"Response"`
}

func (r *DescribeClusterTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClustersRequestParams struct {
	// List of IDs of clusters to be queried
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`

	// Query filter
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Query list offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results in query list. Default value: 20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Whether to enable IPv6 address access
	Ipv6Enable *int64 `json:"Ipv6Enable,omitempty" name:"Ipv6Enable"`
}

type DescribeClustersRequest struct {
	*tchttp.BaseRequest
	
	// List of IDs of clusters to be queried
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`

	// Query filter
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Query list offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results in query list. Default value: 20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Whether to enable IPv6 address access
	Ipv6Enable *int64 `json:"Ipv6Enable,omitempty" name:"Ipv6Enable"`
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
	delete(f, "ClusterIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Ipv6Enable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClustersResponseParams struct {
	// Number of cluster instances
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Cluster instance list
	Clusters []*ClusterInfo `json:"Clusters,omitempty" name:"Clusters"`

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
type DescribeIdlFileInfosRequestParams struct {
	// ID of the cluster where a file resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// ID of the table group where a file resides
	TableGroupIds []*string `json:"TableGroupIds,omitempty" name:"TableGroupIds"`

	// File ID list
	IdlFileIds []*string `json:"IdlFileIds,omitempty" name:"IdlFileIds"`

	// Query list offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results in query list
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeIdlFileInfosRequest struct {
	*tchttp.BaseRequest
	
	// ID of the cluster where a file resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// ID of the table group where a file resides
	TableGroupIds []*string `json:"TableGroupIds,omitempty" name:"TableGroupIds"`

	// File ID list
	IdlFileIds []*string `json:"IdlFileIds,omitempty" name:"IdlFileIds"`

	// Query list offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results in query list
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeIdlFileInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIdlFileInfosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "TableGroupIds")
	delete(f, "IdlFileIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIdlFileInfosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIdlFileInfosResponseParams struct {
	// Number of files
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of file details
	IdlFileInfos []*IdlFileInfo `json:"IdlFileInfos,omitempty" name:"IdlFileInfos"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIdlFileInfosResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIdlFileInfosResponseParams `json:"Response"`
}

func (r *DescribeIdlFileInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIdlFileInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMachineRequestParams struct {
	// If this parameter is not `0`, machines supporting IPv6 will be queried.
	Ipv6Enable *int64 `json:"Ipv6Enable,omitempty" name:"Ipv6Enable"`
}

type DescribeMachineRequest struct {
	*tchttp.BaseRequest
	
	// If this parameter is not `0`, machines supporting IPv6 will be queried.
	Ipv6Enable *int64 `json:"Ipv6Enable,omitempty" name:"Ipv6Enable"`
}

func (r *DescribeMachineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ipv6Enable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMachineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMachineResponseParams struct {
	// The list of dedicated machine resources
	PoolList []*PoolInfo `json:"PoolList,omitempty" name:"PoolList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMachineResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMachineResponseParams `json:"Response"`
}

func (r *DescribeMachineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegionsRequestParams struct {

}

type DescribeRegionsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegionsResponseParams struct {
	// Number of queried AZs
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of AZ query results
	RegionInfos []*RegionInfo `json:"RegionInfos,omitempty" name:"RegionInfos"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRegionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRegionsResponseParams `json:"Response"`
}

func (r *DescribeRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotsRequestParams struct {
	// The ID of the cluster where the table resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The ID of the table group where the table resides
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// Table name
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// Snapshot name
	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`
}

type DescribeSnapshotsRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the cluster where the table resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The ID of the table group where the table resides
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// Table name
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// Snapshot name
	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`
}

func (r *DescribeSnapshotsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "TableGroupId")
	delete(f, "TableName")
	delete(f, "SnapshotName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSnapshotsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotsResponseParams struct {
	// The number of snapshots
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The result list of snapshots
	TableResults []*SnapshotResult `json:"TableResults,omitempty" name:"TableResults"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSnapshotsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSnapshotsResponseParams `json:"Response"`
}

func (r *DescribeSnapshotsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableGroupTagsRequestParams struct {
	// The ID of the cluster where table group tags need to be queried
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The list of IDs of the table groups whose tags need to be queried
	TableGroupIds []*string `json:"TableGroupIds,omitempty" name:"TableGroupIds"`
}

type DescribeTableGroupTagsRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the cluster where table group tags need to be queried
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The list of IDs of the table groups whose tags need to be queried
	TableGroupIds []*string `json:"TableGroupIds,omitempty" name:"TableGroupIds"`
}

func (r *DescribeTableGroupTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableGroupTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "TableGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTableGroupTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableGroupTagsResponseParams struct {
	// The information list of table group tags
	Rows []*TagsInfoOfTableGroup `json:"Rows,omitempty" name:"Rows"`

	// The number of returned results
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTableGroupTagsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTableGroupTagsResponseParams `json:"Response"`
}

func (r *DescribeTableGroupTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableGroupTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableGroupsRequestParams struct {
	// ID of the cluster where a table group resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Table group ID list
	TableGroupIds []*string `json:"TableGroupIds,omitempty" name:"TableGroupIds"`

	// Filter. Valid values: TableGroupName, TableGroupId
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Query list offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results in query list
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeTableGroupsRequest struct {
	*tchttp.BaseRequest
	
	// ID of the cluster where a table group resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Table group ID list
	TableGroupIds []*string `json:"TableGroupIds,omitempty" name:"TableGroupIds"`

	// Filter. Valid values: TableGroupName, TableGroupId
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Query list offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results in query list
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTableGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "TableGroupIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTableGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableGroupsResponseParams struct {
	// Number of table groups
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Table group information list
	TableGroups []*TableGroupInfo `json:"TableGroups,omitempty" name:"TableGroups"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTableGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTableGroupsResponseParams `json:"Response"`
}

func (r *DescribeTableGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableTagsRequestParams struct {
	// The ID of the cluster where a table resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Table list
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

type DescribeTableTagsRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the cluster where a table resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Table list
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

func (r *DescribeTableTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SelectedTables")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTableTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableTagsResponseParams struct {
	// The total number of returned results
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The information list of table tags
	Rows []*TagsInfoOfTable `json:"Rows,omitempty" name:"Rows"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTableTagsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTableTagsResponseParams `json:"Response"`
}

func (r *DescribeTableTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTablesInRecycleRequestParams struct {
	// ID of the cluster where the table to be queried resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// List of IDs of the table groups where the table to be queried resides
	TableGroupIds []*string `json:"TableGroupIds,omitempty" name:"TableGroupIds"`

	// Filter. Valid values: TableName, TableInstanceId
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Query result offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned query results
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeTablesInRecycleRequest struct {
	*tchttp.BaseRequest
	
	// ID of the cluster where the table to be queried resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// List of IDs of the table groups where the table to be queried resides
	TableGroupIds []*string `json:"TableGroupIds,omitempty" name:"TableGroupIds"`

	// Filter. Valid values: TableName, TableInstanceId
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Query result offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned query results
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTablesInRecycleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTablesInRecycleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "TableGroupIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTablesInRecycleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTablesInRecycleResponseParams struct {
	// Number of tables
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Table details result list
	TableInfos []*TableInfoNew `json:"TableInfos,omitempty" name:"TableInfos"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTablesInRecycleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTablesInRecycleResponseParams `json:"Response"`
}

func (r *DescribeTablesInRecycleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTablesInRecycleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTablesRequestParams struct {
	// ID of the cluster where the table to be queried resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// List of IDs of the table groups where the table to be queried resides
	TableGroupIds []*string `json:"TableGroupIds,omitempty" name:"TableGroupIds"`

	// Information list of tables to be queried
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`

	// Filter. Valid values: TableName, TableInstanceId
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Query result offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned query results
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeTablesRequest struct {
	*tchttp.BaseRequest
	
	// ID of the cluster where the table to be queried resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// List of IDs of the table groups where the table to be queried resides
	TableGroupIds []*string `json:"TableGroupIds,omitempty" name:"TableGroupIds"`

	// Information list of tables to be queried
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`

	// Filter. Valid values: TableName, TableInstanceId
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Query result offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned query results
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "TableGroupIds")
	delete(f, "SelectedTables")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTablesResponseParams struct {
	// Number of tables
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Table details result list
	TableInfos []*TableInfoNew `json:"TableInfos,omitempty" name:"TableInfos"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTablesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTablesResponseParams `json:"Response"`
}

func (r *DescribeTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTasksRequestParams struct {
	// List of IDs of clusters where the tasks to be queried reside
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`

	// List of IDs of tasks to be queried
	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds"`

	// Filter. Valid values: Content, TaskType, Operator, Time
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Query list offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results in query list
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeTasksRequest struct {
	*tchttp.BaseRequest
	
	// List of IDs of clusters where the tasks to be queried reside
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`

	// List of IDs of tasks to be queried
	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds"`

	// Filter. Valid values: Content, TaskType, Operator, Time
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Query list offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results in query list
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterIds")
	delete(f, "TaskIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTasksResponseParams struct {
	// Number of tasks
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of details of queried tasks
	TaskInfos []*TaskInfoNew `json:"TaskInfos,omitempty" name:"TaskInfos"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTasksResponseParams `json:"Response"`
}

func (r *DescribeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUinInWhitelistRequestParams struct {

}

type DescribeUinInWhitelistRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeUinInWhitelistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUinInWhitelistRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUinInWhitelistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUinInWhitelistResponseParams struct {
	// Query result. FALSE: yes, TRUE: no
	Result *string `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUinInWhitelistResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUinInWhitelistResponseParams `json:"Response"`
}

func (r *DescribeUinInWhitelistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUinInWhitelistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableRestProxyRequestParams struct {
	// The value is the same as `appid`.
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type DisableRestProxyRequest struct {
	*tchttp.BaseRequest
	
	// The value is the same as `appid`.
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DisableRestProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableRestProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableRestProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableRestProxyResponseParams struct {
	// RestProxy status. Valid values: 0 (disabled), 1 (enabling), 2 (enabled), 3 (disabling).
	RestProxyStatus *uint64 `json:"RestProxyStatus,omitempty" name:"RestProxyStatus"`

	// `TaskId` is in the format of `AppInstanceId-taskId`, used to identify tasks of different clusters.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DisableRestProxyResponse struct {
	*tchttp.BaseResponse
	Response *DisableRestProxyResponseParams `json:"Response"`
}

func (r *DisableRestProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableRestProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableRestProxyRequestParams struct {
	// The value is the same as `appid`.
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type EnableRestProxyRequest struct {
	*tchttp.BaseRequest
	
	// The value is the same as `appid`.
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *EnableRestProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableRestProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableRestProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableRestProxyResponseParams struct {
	// RestProxy status. Valid values: 0 (disabled), 1 (enabling), 2 (enabled), 3 (disabling).
	RestProxyStatus *uint64 `json:"RestProxyStatus,omitempty" name:"RestProxyStatus"`

	// `TaskId` is in the format of `AppInstanceId-taskId`, used to identify tasks of different clusters.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type EnableRestProxyResponse struct {
	*tchttp.BaseResponse
	Response *EnableRestProxyResponseParams `json:"Response"`
}

func (r *EnableRestProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableRestProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ErrorInfo struct {
	// Error code
	Code *string `json:"Code,omitempty" name:"Code"`

	// Error message
	Message *string `json:"Message,omitempty" name:"Message"`
}

type FieldInfo struct {
	// Table field name
	FieldName *string `json:"FieldName,omitempty" name:"FieldName"`

	// Whether it is a primary key field
	IsPrimaryKey *string `json:"IsPrimaryKey,omitempty" name:"IsPrimaryKey"`

	// Field type
	FieldType *string `json:"FieldType,omitempty" name:"FieldType"`

	// Field length
	FieldSize *int64 `json:"FieldSize,omitempty" name:"FieldSize"`
}

type Filter struct {
	// Filter field name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Filter field value
	Value *string `json:"Value,omitempty" name:"Value"`

	// Filter field value
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type IdlFileInfo struct {
	// Filename excluding extension
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// Data interface description language (IDL) type
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// File extension
	FileExtType *string `json:"FileExtType,omitempty" name:"FileExtType"`

	// File size in bytes
	FileSize *int64 `json:"FileSize,omitempty" name:"FileSize"`

	// File ID, which is meaningful for files already uploaded
	// Note: this field may return null, indicating that no valid values can be obtained.
	FileId *int64 `json:"FileId,omitempty" name:"FileId"`

	// File content, which is meaningful for files to be uploaded in this request
	// Note: this field may return null, indicating that no valid values can be obtained.
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`
}

type IdlFileInfoWithoutContent struct {
	// Filename excluding extension
	// Note: this field may return null, indicating that no valid values can be obtained.
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// Data interface description language (IDL) type
	// Note: this field may return null, indicating that no valid values can be obtained.
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// File extension
	// Note: this field may return null, indicating that no valid values can be obtained.
	FileExtType *string `json:"FileExtType,omitempty" name:"FileExtType"`

	// File size in bytes
	// Note: this field may return null, indicating that no valid values can be obtained.
	FileSize *int64 `json:"FileSize,omitempty" name:"FileSize"`

	// File ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	FileId *int64 `json:"FileId,omitempty" name:"FileId"`

	// Error message
	// Note: this field may return null, indicating that no valid values can be obtained.
	Error *ErrorInfo `json:"Error,omitempty" name:"Error"`
}

// Predefined struct for user
type ImportSnapshotsRequestParams struct {
	// The ID of the cluster where the original table (from which the snapshot was created) resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The information of the snapshot to import
	Snapshots *SnapshotInfo `json:"Snapshots,omitempty" name:"Snapshots"`

	// Whether to import partial data of the snapshot. Valid values: `TRUE` (import partial data), `FALSE` (import all data).
	ImportSpecialKey *string `json:"ImportSpecialKey,omitempty" name:"ImportSpecialKey"`

	// Whether to import to the original table. Valid values: `TRUE` (import to the original table), `FALSE` (import to a new table).
	ImportOriginTable *string `json:"ImportOriginTable,omitempty" name:"ImportOriginTable"`

	// The file of the keys of the partial data
	KeyFile *KeyFile `json:"KeyFile,omitempty" name:"KeyFile"`

	// The ID of the table group where the new table resides, which is valid only when `ImportOriginTable` is set to `FALSE`
	NewTableGroupId *string `json:"NewTableGroupId,omitempty" name:"NewTableGroupId"`

	// The name of the new table, which is valid only when `ImportOriginTable` is set to `FALSE`. TcaplusDB will automatically create a table named `NewTableName` of the same structure as that of the original table.
	NewTableName *string `json:"NewTableName,omitempty" name:"NewTableName"`
}

type ImportSnapshotsRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the cluster where the original table (from which the snapshot was created) resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The information of the snapshot to import
	Snapshots *SnapshotInfo `json:"Snapshots,omitempty" name:"Snapshots"`

	// Whether to import partial data of the snapshot. Valid values: `TRUE` (import partial data), `FALSE` (import all data).
	ImportSpecialKey *string `json:"ImportSpecialKey,omitempty" name:"ImportSpecialKey"`

	// Whether to import to the original table. Valid values: `TRUE` (import to the original table), `FALSE` (import to a new table).
	ImportOriginTable *string `json:"ImportOriginTable,omitempty" name:"ImportOriginTable"`

	// The file of the keys of the partial data
	KeyFile *KeyFile `json:"KeyFile,omitempty" name:"KeyFile"`

	// The ID of the table group where the new table resides, which is valid only when `ImportOriginTable` is set to `FALSE`
	NewTableGroupId *string `json:"NewTableGroupId,omitempty" name:"NewTableGroupId"`

	// The name of the new table, which is valid only when `ImportOriginTable` is set to `FALSE`. TcaplusDB will automatically create a table named `NewTableName` of the same structure as that of the original table.
	NewTableName *string `json:"NewTableName,omitempty" name:"NewTableName"`
}

func (r *ImportSnapshotsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportSnapshotsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Snapshots")
	delete(f, "ImportSpecialKey")
	delete(f, "ImportOriginTable")
	delete(f, "KeyFile")
	delete(f, "NewTableGroupId")
	delete(f, "NewTableName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImportSnapshotsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportSnapshotsResponseParams struct {
	// `TaskId` is in the format of `AppInstanceId-taskId`, used to identify tasks of different clusters.
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ImportSnapshotsResponse struct {
	*tchttp.BaseResponse
	Response *ImportSnapshotsResponseParams `json:"Response"`
}

func (r *ImportSnapshotsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportSnapshotsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KafkaInfo struct {
	// CKafka address
	Address *string `json:"Address,omitempty" name:"Address"`

	// CKafka topic
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// CKafka username
	User *string `json:"User,omitempty" name:"User"`

	// CKafka password
	Password *string `json:"Password,omitempty" name:"Password"`

	// CKafka instance
	Instance *string `json:"Instance,omitempty" name:"Instance"`

	// Whether VPC access is enabled
	IsVpc *int64 `json:"IsVpc,omitempty" name:"IsVpc"`
}

type KeyFile struct {
	// Key file name
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// Key file extension
	FileExtType *string `json:"FileExtType,omitempty" name:"FileExtType"`

	// Key file content
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`

	// Key file size
	FileSize *int64 `json:"FileSize,omitempty" name:"FileSize"`
}

type MachineInfo struct {
	// Machine type
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// Machine quantity
	MachineNum *int64 `json:"MachineNum,omitempty" name:"MachineNum"`
}

type MergeTableResult struct {
	// Task ID
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// If table merging is successful, `null` will be returned
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	Error *ErrorInfo `json:"Error,omitempty" name:"Error"`

	// Comparison results of tables
	Table *CompareTablesInfo `json:"Table,omitempty" name:"Table"`

	// Application ID
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

// Predefined struct for user
type MergeTablesDataRequestParams struct {
	// Tables to be merged
	SelectedTables []*MergeTablesInfo `json:"SelectedTables,omitempty" name:"SelectedTables"`

	// Valid values: `true` (only compare), `false` (compare and merge)
	IsOnlyCompare *bool `json:"IsOnlyCompare,omitempty" name:"IsOnlyCompare"`
}

type MergeTablesDataRequest struct {
	*tchttp.BaseRequest
	
	// Tables to be merged
	SelectedTables []*MergeTablesInfo `json:"SelectedTables,omitempty" name:"SelectedTables"`

	// Valid values: `true` (only compare), `false` (compare and merge)
	IsOnlyCompare *bool `json:"IsOnlyCompare,omitempty" name:"IsOnlyCompare"`
}

func (r *MergeTablesDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MergeTablesDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SelectedTables")
	delete(f, "IsOnlyCompare")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MergeTablesDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MergeTablesDataResponseParams struct {
	// Table merging results
	Results []*MergeTableResult `json:"Results,omitempty" name:"Results"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type MergeTablesDataResponse struct {
	*tchttp.BaseResponse
	Response *MergeTablesDataResponseParams `json:"Response"`
}

func (r *MergeTablesDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MergeTablesDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MergeTablesInfo struct {
	// Information of tables to be merged
	MergeTables *CompareTablesInfo `json:"MergeTables,omitempty" name:"MergeTables"`

	// Whether to check indexes
	CheckIndex *bool `json:"CheckIndex,omitempty" name:"CheckIndex"`
}

// Predefined struct for user
type ModifyCensorshipRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Whether to enable the operation approval feature for this cluster. Valid values: `0` (disable), `1` (enable)
	Censorship *int64 `json:"Censorship,omitempty" name:"Censorship"`

	// Approver UIN list
	Uins []*string `json:"Uins,omitempty" name:"Uins"`
}

type ModifyCensorshipRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Whether to enable the operation approval feature for this cluster. Valid values: `0` (disable), `1` (enable)
	Censorship *int64 `json:"Censorship,omitempty" name:"Censorship"`

	// Approver UIN list
	Uins []*string `json:"Uins,omitempty" name:"Uins"`
}

func (r *ModifyCensorshipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCensorshipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Censorship")
	delete(f, "Uins")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCensorshipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCensorshipResponseParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Approver UIN list
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	Uins []*string `json:"Uins,omitempty" name:"Uins"`

	// Whether the operation approval feature is enabled for this cluster. Valid values: `0` (disabled), `1` (enabled)
	Censorship *int64 `json:"Censorship,omitempty" name:"Censorship"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyCensorshipResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCensorshipResponseParams `json:"Response"`
}

func (r *ModifyCensorshipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCensorshipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterMachineRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Information of the machines at the storage layer (tcapsvr)
	ServerList []*MachineInfo `json:"ServerList,omitempty" name:"ServerList"`

	// Information of the machines at the access layer (tcaproxy)
	ProxyList []*MachineInfo `json:"ProxyList,omitempty" name:"ProxyList"`

	// Cluster type. Valid values: `1` (standard), `2` (dedicated)
	ClusterType *int64 `json:"ClusterType,omitempty" name:"ClusterType"`
}

type ModifyClusterMachineRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Information of the machines at the storage layer (tcapsvr)
	ServerList []*MachineInfo `json:"ServerList,omitempty" name:"ServerList"`

	// Information of the machines at the access layer (tcaproxy)
	ProxyList []*MachineInfo `json:"ProxyList,omitempty" name:"ProxyList"`

	// Cluster type. Valid values: `1` (standard), `2` (dedicated)
	ClusterType *int64 `json:"ClusterType,omitempty" name:"ClusterType"`
}

func (r *ModifyClusterMachineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterMachineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ServerList")
	delete(f, "ProxyList")
	delete(f, "ClusterType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterMachineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterMachineResponseParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyClusterMachineResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterMachineResponseParams `json:"Response"`
}

func (r *ModifyClusterMachineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterMachineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterNameRequestParams struct {
	// ID of the cluster to be renamed
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Cluster name to be changed to, which can contain up to 32 letters and digits
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

type ModifyClusterNameRequest struct {
	*tchttp.BaseRequest
	
	// ID of the cluster to be renamed
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Cluster name to be changed to, which can contain up to 32 letters and digits
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
type ModifyClusterPasswordRequestParams struct {
	// ID of the cluster for which to modify the password
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Old cluster password
	OldPassword *string `json:"OldPassword,omitempty" name:"OldPassword"`

	// Expected expiration time of old cluster password
	OldPasswordExpireTime *string `json:"OldPasswordExpireTime,omitempty" name:"OldPasswordExpireTime"`

	// New cluster password, which must contain lowercase letters (a-z), uppercase letters (A-Z), and digits (0-9).
	NewPassword *string `json:"NewPassword,omitempty" name:"NewPassword"`

	// Update mode. 1: updates password, 2: updates old password expiration time. Default value: 1
	Mode *string `json:"Mode,omitempty" name:"Mode"`
}

type ModifyClusterPasswordRequest struct {
	*tchttp.BaseRequest
	
	// ID of the cluster for which to modify the password
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Old cluster password
	OldPassword *string `json:"OldPassword,omitempty" name:"OldPassword"`

	// Expected expiration time of old cluster password
	OldPasswordExpireTime *string `json:"OldPasswordExpireTime,omitempty" name:"OldPasswordExpireTime"`

	// New cluster password, which must contain lowercase letters (a-z), uppercase letters (A-Z), and digits (0-9).
	NewPassword *string `json:"NewPassword,omitempty" name:"NewPassword"`

	// Update mode. 1: updates password, 2: updates old password expiration time. Default value: 1
	Mode *string `json:"Mode,omitempty" name:"Mode"`
}

func (r *ModifyClusterPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "OldPassword")
	delete(f, "OldPasswordExpireTime")
	delete(f, "NewPassword")
	delete(f, "Mode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterPasswordResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyClusterPasswordResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterPasswordResponseParams `json:"Response"`
}

func (r *ModifyClusterPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterTagsRequestParams struct {
	// The ID of the cluster whose tags need to be modified
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The list of tags to add or modify
	ReplaceTags []*TagInfoUnit `json:"ReplaceTags,omitempty" name:"ReplaceTags"`

	// Tags to delete
	DeleteTags []*TagInfoUnit `json:"DeleteTags,omitempty" name:"DeleteTags"`
}

type ModifyClusterTagsRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the cluster whose tags need to be modified
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The list of tags to add or modify
	ReplaceTags []*TagInfoUnit `json:"ReplaceTags,omitempty" name:"ReplaceTags"`

	// Tags to delete
	DeleteTags []*TagInfoUnit `json:"DeleteTags,omitempty" name:"DeleteTags"`
}

func (r *ModifyClusterTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ReplaceTags")
	delete(f, "DeleteTags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterTagsResponseParams struct {
	// Task ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyClusterTagsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterTagsResponseParams `json:"Response"`
}

func (r *ModifyClusterTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySnapshotsRequestParams struct {
	// The ID of the cluster where the table resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Snapshot list
	SelectedTables []*SnapshotInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

type ModifySnapshotsRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the cluster where the table resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Snapshot list
	SelectedTables []*SnapshotInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

func (r *ModifySnapshotsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySnapshotsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SelectedTables")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySnapshotsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySnapshotsResponseParams struct {
	// The number of snapshots modified in batches
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The result list of snapshots modified in batches
	TableResults []*SnapshotResult `json:"TableResults,omitempty" name:"TableResults"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySnapshotsResponse struct {
	*tchttp.BaseResponse
	Response *ModifySnapshotsResponseParams `json:"Response"`
}

func (r *ModifySnapshotsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySnapshotsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTableGroupNameRequestParams struct {
	// ID of the cluster where a table group resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// ID of the table group to be renamed
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// New table group name, which can contain letters and symbols
	TableGroupName *string `json:"TableGroupName,omitempty" name:"TableGroupName"`
}

type ModifyTableGroupNameRequest struct {
	*tchttp.BaseRequest
	
	// ID of the cluster where a table group resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// ID of the table group to be renamed
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// New table group name, which can contain letters and symbols
	TableGroupName *string `json:"TableGroupName,omitempty" name:"TableGroupName"`
}

func (r *ModifyTableGroupNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTableGroupNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "TableGroupId")
	delete(f, "TableGroupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTableGroupNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTableGroupNameResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyTableGroupNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTableGroupNameResponseParams `json:"Response"`
}

func (r *ModifyTableGroupNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTableGroupNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTableGroupTagsRequestParams struct {
	// The ID of the cluster where table group tags need to be modified
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The ID of the table group whose tags need to be modified
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// The list of tags to add or modify
	ReplaceTags []*TagInfoUnit `json:"ReplaceTags,omitempty" name:"ReplaceTags"`

	// Tags to delete
	DeleteTags []*TagInfoUnit `json:"DeleteTags,omitempty" name:"DeleteTags"`
}

type ModifyTableGroupTagsRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the cluster where table group tags need to be modified
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The ID of the table group whose tags need to be modified
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// The list of tags to add or modify
	ReplaceTags []*TagInfoUnit `json:"ReplaceTags,omitempty" name:"ReplaceTags"`

	// Tags to delete
	DeleteTags []*TagInfoUnit `json:"DeleteTags,omitempty" name:"DeleteTags"`
}

func (r *ModifyTableGroupTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTableGroupTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "TableGroupId")
	delete(f, "ReplaceTags")
	delete(f, "DeleteTags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTableGroupTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTableGroupTagsResponseParams struct {
	// Task ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyTableGroupTagsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTableGroupTagsResponseParams `json:"Response"`
}

func (r *ModifyTableGroupTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTableGroupTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTableMemosRequestParams struct {
	// ID of the cluster instance where a table resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// List of details of selected tables
	TableMemos []*SelectedTableInfoNew `json:"TableMemos,omitempty" name:"TableMemos"`
}

type ModifyTableMemosRequest struct {
	*tchttp.BaseRequest
	
	// ID of the cluster instance where a table resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// List of details of selected tables
	TableMemos []*SelectedTableInfoNew `json:"TableMemos,omitempty" name:"TableMemos"`
}

func (r *ModifyTableMemosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTableMemosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "TableMemos")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTableMemosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTableMemosResponseParams struct {
	// Number of tables modified for remarks
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of table remarks modification results
	TableResults []*TableResultNew `json:"TableResults,omitempty" name:"TableResults"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyTableMemosResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTableMemosResponseParams `json:"Response"`
}

func (r *ModifyTableMemosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTableMemosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTableQuotasRequestParams struct {
	// ID of the cluster where the table to be scaled resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// List of quotas of tables selected for modification
	TableQuotas []*SelectedTableInfoNew `json:"TableQuotas,omitempty" name:"TableQuotas"`
}

type ModifyTableQuotasRequest struct {
	*tchttp.BaseRequest
	
	// ID of the cluster where the table to be scaled resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// List of quotas of tables selected for modification
	TableQuotas []*SelectedTableInfoNew `json:"TableQuotas,omitempty" name:"TableQuotas"`
}

func (r *ModifyTableQuotasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTableQuotasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "TableQuotas")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTableQuotasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTableQuotasResponseParams struct {
	// Number of scaled tables
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of table scaling results
	TableResults []*TableResultNew `json:"TableResults,omitempty" name:"TableResults"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyTableQuotasResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTableQuotasResponseParams `json:"Response"`
}

func (r *ModifyTableQuotasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTableQuotasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTableTagsRequestParams struct {
	// The ID of the cluster where table tags need to be modified
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The list of tables whose tags need to be modified
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`

	// The list of tags to add or modify
	ReplaceTags []*TagInfoUnit `json:"ReplaceTags,omitempty" name:"ReplaceTags"`

	// The list of tags to delete
	DeleteTags []*TagInfoUnit `json:"DeleteTags,omitempty" name:"DeleteTags"`
}

type ModifyTableTagsRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the cluster where table tags need to be modified
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The list of tables whose tags need to be modified
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`

	// The list of tags to add or modify
	ReplaceTags []*TagInfoUnit `json:"ReplaceTags,omitempty" name:"ReplaceTags"`

	// The list of tags to delete
	DeleteTags []*TagInfoUnit `json:"DeleteTags,omitempty" name:"DeleteTags"`
}

func (r *ModifyTableTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTableTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SelectedTables")
	delete(f, "ReplaceTags")
	delete(f, "DeleteTags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTableTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTableTagsResponseParams struct {
	// The total number of returned results
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Returned results
	TableResults []*TableResultNew `json:"TableResults,omitempty" name:"TableResults"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyTableTagsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTableTagsResponseParams `json:"Response"`
}

func (r *ModifyTableTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTableTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTablesRequestParams struct {
	// ID of the cluster where the table to be modified resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Selected table modification IDL files
	IdlFiles []*IdlFileInfo `json:"IdlFiles,omitempty" name:"IdlFiles"`

	// List of tables to be modified
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

type ModifyTablesRequest struct {
	*tchttp.BaseRequest
	
	// ID of the cluster where the table to be modified resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Selected table modification IDL files
	IdlFiles []*IdlFileInfo `json:"IdlFiles,omitempty" name:"IdlFiles"`

	// List of tables to be modified
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

func (r *ModifyTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "IdlFiles")
	delete(f, "SelectedTables")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTablesResponseParams struct {
	// Number of modified tables
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of table modification results
	TableResults []*TableResultNew `json:"TableResults,omitempty" name:"TableResults"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyTablesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTablesResponseParams `json:"Response"`
}

func (r *ModifyTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ParsedTableInfoNew struct {
	// Table description language type. Valid values: PROTO, TDR
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableIdlType *string `json:"TableIdlType,omitempty" name:"TableIdlType"`

	// Table instance ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableInstanceId *string `json:"TableInstanceId,omitempty" name:"TableInstanceId"`

	// Table name
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// Table data structure type. Valid values: GENERIC, LIST
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableType *string `json:"TableType,omitempty" name:"TableType"`

	// Primary key field information
	// Note: this field may return null, indicating that no valid values can be obtained.
	KeyFields *string `json:"KeyFields,omitempty" name:"KeyFields"`

	// Old primary key field information, which is valid during verification of table modification
	// Note: this field may return null, indicating that no valid values can be obtained.
	OldKeyFields *string `json:"OldKeyFields,omitempty" name:"OldKeyFields"`

	// Non-primary key field information
	// Note: this field may return null, indicating that no valid values can be obtained.
	ValueFields *string `json:"ValueFields,omitempty" name:"ValueFields"`

	// Old non-primary key field information, which is valid during verification of table modification
	// Note: this field may return null, indicating that no valid values can be obtained.
	OldValueFields *string `json:"OldValueFields,omitempty" name:"OldValueFields"`

	// Table group ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// Total size of primary key field
	// Note: this field may return null, indicating that no valid values can be obtained.
	SumKeyFieldSize *int64 `json:"SumKeyFieldSize,omitempty" name:"SumKeyFieldSize"`

	// Total size of non-primary key fields
	// Note: this field may return null, indicating that no valid values can be obtained.
	SumValueFieldSize *int64 `json:"SumValueFieldSize,omitempty" name:"SumValueFieldSize"`

	// Index key set
	// Note: this field may return null, indicating that no valid values can be obtained.
	IndexKeySet *string `json:"IndexKeySet,omitempty" name:"IndexKeySet"`

	// Shardkey set
	// Note: this field may return null, indicating that no valid values can be obtained.
	ShardingKeySet *string `json:"ShardingKeySet,omitempty" name:"ShardingKeySet"`

	// TDR version number
	// Note: this field may return null, indicating that no valid values can be obtained.
	TdrVersion *int64 `json:"TdrVersion,omitempty" name:"TdrVersion"`

	// Error message
	// Note: this field may return null, indicating that no valid values can be obtained.
	Error *ErrorInfo `json:"Error,omitempty" name:"Error"`

	// Number of LIST-type table elements
	// Note: this field may return null, indicating that no valid values can be obtained.
	ListElementNum *int64 `json:"ListElementNum,omitempty" name:"ListElementNum"`

	// Number of SORTLIST-type table sort fields
	// Note: this field may return null, indicating that no valid values can be obtained.
	SortFieldNum *int64 `json:"SortFieldNum,omitempty" name:"SortFieldNum"`

	// Sort order of SORTLIST-type tables
	// Note: this field may return null, indicating that no valid values can be obtained.
	SortRule *int64 `json:"SortRule,omitempty" name:"SortRule"`
}

type PoolInfo struct {
	// Unique ID
	PoolUid *int64 `json:"PoolUid,omitempty" name:"PoolUid"`

	// Whether IPv6 is supported
	Ipv6Enable *int64 `json:"Ipv6Enable,omitempty" name:"Ipv6Enable"`

	// Remaining available cluster resources
	AvailableAppCount *int64 `json:"AvailableAppCount,omitempty" name:"AvailableAppCount"`

	// The list of machines at the storage layer (tcapsvr)
	ServerList []*ServerMachineInfo `json:"ServerList,omitempty" name:"ServerList"`

	// The list of machines at the access layer (tcaproxy)
	ProxyList []*ProxyMachineInfo `json:"ProxyList,omitempty" name:"ProxyList"`
}

type ProxyDetailInfo struct {
	// The unique ID of the access layer (tcaproxy)
	ProxyUid *string `json:"ProxyUid,omitempty" name:"ProxyUid"`

	// Machine type
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// The speed of processing request packets
	ProcessSpeed *int64 `json:"ProcessSpeed,omitempty" name:"ProcessSpeed"`

	// Request packet delay
	AverageProcessDelay *int64 `json:"AverageProcessDelay,omitempty" name:"AverageProcessDelay"`

	// The speed of processing delayed request packets
	SlowProcessSpeed *int64 `json:"SlowProcessSpeed,omitempty" name:"SlowProcessSpeed"`
}

type ProxyMachineInfo struct {
	// Unique ID
	ProxyUid *string `json:"ProxyUid,omitempty" name:"ProxyUid"`

	// Machine type
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// The number of proxy resources to be assigned
	AvailableCount *int64 `json:"AvailableCount,omitempty" name:"AvailableCount"`
}

// Predefined struct for user
type RecoverRecycleTablesRequestParams struct {
	// ID of the cluster where a table resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Information of tables to be recovered
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

type RecoverRecycleTablesRequest struct {
	*tchttp.BaseRequest
	
	// ID of the cluster where a table resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Information of tables to be recovered
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

func (r *RecoverRecycleTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecoverRecycleTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SelectedTables")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecoverRecycleTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecoverRecycleTablesResponseParams struct {
	// Number of recovered tables
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of information of recovered tables
	TableResults []*TableResultNew `json:"TableResults,omitempty" name:"TableResults"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RecoverRecycleTablesResponse struct {
	*tchttp.BaseResponse
	Response *RecoverRecycleTablesResponseParams `json:"Response"`
}

func (r *RecoverRecycleTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecoverRecycleTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionInfo struct {
	// Region `Ap-code`
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// Region abbreviation
	RegionAbbr *string `json:"RegionAbbr,omitempty" name:"RegionAbbr"`

	// Region ID
	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`

	// Whether to support IPv6 address access. Valid values: 0 (support), 1 (not support)
	Ipv6Enable *uint64 `json:"Ipv6Enable,omitempty" name:"Ipv6Enable"`
}

// Predefined struct for user
type RollbackTablesRequestParams struct {
	// ID of the cluster where the table to be rolled back resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// List of tables to be rolled back
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`

	// Time to roll back to
	RollbackTime *string `json:"RollbackTime,omitempty" name:"RollbackTime"`

	// Rollback mode. `KEYS` is supported
	Mode *string `json:"Mode,omitempty" name:"Mode"`
}

type RollbackTablesRequest struct {
	*tchttp.BaseRequest
	
	// ID of the cluster where the table to be rolled back resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// List of tables to be rolled back
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`

	// Time to roll back to
	RollbackTime *string `json:"RollbackTime,omitempty" name:"RollbackTime"`

	// Rollback mode. `KEYS` is supported
	Mode *string `json:"Mode,omitempty" name:"Mode"`
}

func (r *RollbackTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollbackTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SelectedTables")
	delete(f, "RollbackTime")
	delete(f, "Mode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RollbackTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RollbackTablesResponseParams struct {
	// Number of table rollback task results
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Table rollback task result list
	TableResults []*TableRollbackResultNew `json:"TableResults,omitempty" name:"TableResults"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RollbackTablesResponse struct {
	*tchttp.BaseResponse
	Response *RollbackTablesResponseParams `json:"Response"`
}

func (r *RollbackTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollbackTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SelectedTableInfoNew struct {
	// ID of the table group where a table resides
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// Table name
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// Table instance ID
	TableInstanceId *string `json:"TableInstanceId,omitempty" name:"TableInstanceId"`

	// Table description language type. Valid values: PROTO, TDR
	TableIdlType *string `json:"TableIdlType,omitempty" name:"TableIdlType"`

	// Table data structure type. Valid values: GENERIC, LIST
	TableType *string `json:"TableType,omitempty" name:"TableType"`

	// Number of LIST-type table elements
	ListElementNum *int64 `json:"ListElementNum,omitempty" name:"ListElementNum"`

	// Reserved table capacity in GB
	ReservedVolume *int64 `json:"ReservedVolume,omitempty" name:"ReservedVolume"`

	// Reserved table read QPS
	ReservedReadQps *int64 `json:"ReservedReadQps,omitempty" name:"ReservedReadQps"`

	// Reserved table write QPS
	ReservedWriteQps *int64 `json:"ReservedWriteQps,omitempty" name:"ReservedWriteQps"`

	// Table remarks
	Memo *string `json:"Memo,omitempty" name:"Memo"`

	// Key rollback filename, which is only used for rollback
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// Key rollback file extension, which is only used for rollback
	FileExtType *string `json:"FileExtType,omitempty" name:"FileExtType"`

	// Key rollback file size, which is only used for rollback
	FileSize *int64 `json:"FileSize,omitempty" name:"FileSize"`

	// Key rollback file content, which is only used for rollback
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`
}

type SelectedTableWithField struct {
	// ID of the table group where the table resides
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// Table name
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// Table ID
	TableInstanceId *string `json:"TableInstanceId,omitempty" name:"TableInstanceId"`

	// Table description language. Valid values: `PROTO`, `TDR`
	TableIdlType *string `json:"TableIdlType,omitempty" name:"TableIdlType"`

	// Table data structure. Valid values: `GENERIC`, `LIST`
	TableType *string `json:"TableType,omitempty" name:"TableType"`

	// The list of fields on which indexes will be created, table caching enabled, or data subscription enabled
	SelectedFields []*FieldInfo `json:"SelectedFields,omitempty" name:"SelectedFields"`

	// The number of index shards
	ShardNum *uint64 `json:"ShardNum,omitempty" name:"ShardNum"`

	// CKafka instance information
	KafkaInfo *KafkaInfo `json:"KafkaInfo,omitempty" name:"KafkaInfo"`
}

type ServerDetailInfo struct {
	// The unique ID of the storage layer (tcapsvr)
	ServerUid *string `json:"ServerUid,omitempty" name:"ServerUid"`

	// Machine type
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// Memory utilization
	MemoryRate *int64 `json:"MemoryRate,omitempty" name:"MemoryRate"`

	// Disk utilization
	DiskRate *int64 `json:"DiskRate,omitempty" name:"DiskRate"`

	// The number of reads
	ReadNum *int64 `json:"ReadNum,omitempty" name:"ReadNum"`

	// The number of writes
	WriteNum *int64 `json:"WriteNum,omitempty" name:"WriteNum"`
}

type ServerMachineInfo struct {
	// The unique ID of the machine
	ServerUid *string `json:"ServerUid,omitempty" name:"ServerUid"`

	// Machine type
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
}

// Predefined struct for user
type SetTableDataFlowRequestParams struct {
	// The ID of the cluster where the tables reside
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The list of tables for which data subscription will be enabled
	SelectedTables []*SelectedTableWithField `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

type SetTableDataFlowRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the cluster where the tables reside
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The list of tables for which data subscription will be enabled
	SelectedTables []*SelectedTableWithField `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

func (r *SetTableDataFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetTableDataFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SelectedTables")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetTableDataFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetTableDataFlowResponseParams struct {
	// The number of tables for which data subscription has been enabled
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The result list of tables for which data subscription has been enabled
	TableResults []*TableResultNew `json:"TableResults,omitempty" name:"TableResults"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetTableDataFlowResponse struct {
	*tchttp.BaseResponse
	Response *SetTableDataFlowResponseParams `json:"Response"`
}

func (r *SetTableDataFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetTableDataFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetTableIndexRequestParams struct {
	// ID of the cluster where the table resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The list of tables that need to create global indexes
	SelectedTables []*SelectedTableWithField `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

type SetTableIndexRequest struct {
	*tchttp.BaseRequest
	
	// ID of the cluster where the table resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The list of tables that need to create global indexes
	SelectedTables []*SelectedTableWithField `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

func (r *SetTableIndexRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetTableIndexRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SelectedTables")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetTableIndexRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetTableIndexResponseParams struct {
	// The number of tables whose global indexes are created
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The list of global index creation results
	TableResults []*TableResultNew `json:"TableResults,omitempty" name:"TableResults"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetTableIndexResponse struct {
	*tchttp.BaseResponse
	Response *SetTableIndexResponseParams `json:"Response"`
}

func (r *SetTableIndexResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetTableIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SnapshotInfo struct {
	// The ID of the table group where the table resides
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// Table name
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// Snapshot name
	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`

	// The time of the data from which the snapshot was created
	SnapshotTime *string `json:"SnapshotTime,omitempty" name:"SnapshotTime"`

	// Snapshot expiration time
	SnapshotDeadTime *string `json:"SnapshotDeadTime,omitempty" name:"SnapshotDeadTime"`
}

type SnapshotInfoNew struct {
	// The ID of the table group where the table resides
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// Table name
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// Snapshot name
	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`

	// Snapshot expiration time
	SnapshotDeadTime *string `json:"SnapshotDeadTime,omitempty" name:"SnapshotDeadTime"`
}

type SnapshotResult struct {
	// The ID of the table group where the table resides
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// Table name
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// Task ID, which is valid for the API that creates one task at a time
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Error information
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	Error *ErrorInfo `json:"Error,omitempty" name:"Error"`

	// Snapshot name
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`

	// The time of the data from which the snapshot was created
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	SnapshotTime *string `json:"SnapshotTime,omitempty" name:"SnapshotTime"`

	// When the snapshot expires
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	SnapshotDeadTime *string `json:"SnapshotDeadTime,omitempty" name:"SnapshotDeadTime"`

	// When the snapshot was created
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	SnapshotCreateTime *string `json:"SnapshotCreateTime,omitempty" name:"SnapshotCreateTime"`

	// Snapshot size
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	SnapshotSize *uint64 `json:"SnapshotSize,omitempty" name:"SnapshotSize"`

	// Snapshot status. Valid values: `0` (creating), `1` (normal), `2` (deleting), `3` (expired), `4` (rolling back).
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	SnapshotStatus *uint64 `json:"SnapshotStatus,omitempty" name:"SnapshotStatus"`
}

type TableGroupInfo struct {
	// Table group ID
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// Table group name
	TableGroupName *string `json:"TableGroupName,omitempty" name:"TableGroupName"`

	// Table group creation time
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Number of tables in table group
	TableCount *uint64 `json:"TableCount,omitempty" name:"TableCount"`

	// Total table storage capacity in MB in table group
	TotalSize *uint64 `json:"TotalSize,omitempty" name:"TotalSize"`
}

type TableInfoNew struct {
	// Table name
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// Table instance ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableInstanceId *string `json:"TableInstanceId,omitempty" name:"TableInstanceId"`

	// Table data structure type, such as `GENERIC` or `LIST`
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableType *string `json:"TableType,omitempty" name:"TableType"`

	// Table data interface description language (IDL) type, such as `PROTO` or `TDR`
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableIdlType *string `json:"TableIdlType,omitempty" name:"TableIdlType"`

	// ID of the cluster where a table resides
	// Note: this field may return null, indicating that no valid values can be obtained.
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Name of the cluster where a table resides
	// Note: this field may return null, indicating that no valid values can be obtained.
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// ID of the table group where a table resides
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// Name of the table group where a table resides
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableGroupName *string `json:"TableGroupName,omitempty" name:"TableGroupName"`

	// JSON string of table's primary key field structure
	// Note: this field may return null, indicating that no valid values can be obtained.
	KeyStruct *string `json:"KeyStruct,omitempty" name:"KeyStruct"`

	// JSON string of table non-primary key field structure
	// Note: this field may return null, indicating that no valid values can be obtained.
	ValueStruct *string `json:"ValueStruct,omitempty" name:"ValueStruct"`

	// Table shardkey set, which is valid for PROTO-type tables
	// Note: this field may return null, indicating that no valid values can be obtained.
	ShardingKeySet *string `json:"ShardingKeySet,omitempty" name:"ShardingKeySet"`

	// Table index key field set, which is valid for PROTO-type tables
	// Note: this field may return null, indicating that no valid values can be obtained.
	IndexStruct *string `json:"IndexStruct,omitempty" name:"IndexStruct"`

	// Number of LIST-type table elements
	// Note: this field may return null, indicating that no valid values can be obtained.
	ListElementNum *uint64 `json:"ListElementNum,omitempty" name:"ListElementNum"`

	// Information list of IDL files associated with table
	// Note: this field may return null, indicating that no valid values can be obtained.
	IdlFiles []*IdlFileInfo `json:"IdlFiles,omitempty" name:"IdlFiles"`

	// Reserved table capacity in GB
	// Note: this field may return null, indicating that no valid values can be obtained.
	ReservedVolume *int64 `json:"ReservedVolume,omitempty" name:"ReservedVolume"`

	// Reserved table read QPS
	// Note: this field may return null, indicating that no valid values can be obtained.
	ReservedReadQps *int64 `json:"ReservedReadQps,omitempty" name:"ReservedReadQps"`

	// Reserved table write QPS
	// Note: this field may return null, indicating that no valid values can be obtained.
	ReservedWriteQps *int64 `json:"ReservedWriteQps,omitempty" name:"ReservedWriteQps"`

	// Actual table data size in MB
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableSize *int64 `json:"TableSize,omitempty" name:"TableSize"`

	// Table status
	// Note: this field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Table creation time
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Table's last modified time
	// Note: this field may return null, indicating that no valid values can be obtained.
	UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`

	// Table remarks
	// Note: this field may return null, indicating that no valid values can be obtained.
	Memo *string `json:"Memo,omitempty" name:"Memo"`

	// Error message
	// Note: this field may return null, indicating that no valid values can be obtained.
	Error *ErrorInfo `json:"Error,omitempty" name:"Error"`

	// TcaplusDB SDK data access ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiAccessId *string `json:"ApiAccessId,omitempty" name:"ApiAccessId"`

	// Number of SORTLIST-type table sort fields
	// Note: this field may return null, indicating that no valid values can be obtained.
	SortFieldNum *int64 `json:"SortFieldNum,omitempty" name:"SortFieldNum"`

	// Sort order of SORTLIST-type tables
	// Note: this field may return null, indicating that no valid values can be obtained.
	SortRule *int64 `json:"SortRule,omitempty" name:"SortRule"`

	// Information about global indexes, table caching, or data subscription
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DbClusterInfoStruct *string `json:"DbClusterInfoStruct,omitempty" name:"DbClusterInfoStruct"`

	// The number of days after which the table Txh backup files will be expire and deleted.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	TxhBackupExpireDay *uint64 `json:"TxhBackupExpireDay,omitempty" name:"TxhBackupExpireDay"`
}

type TableResultNew struct {
	// Table instance ID in the format of `tcaplus-3be64cbb`
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableInstanceId *string `json:"TableInstanceId,omitempty" name:"TableInstanceId"`

	// Task ID, which is valid for the API that creates one task
	// Note: this field may return null, indicating that no valid values can be obtained.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Table name
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// Table data structure type, such as `GENERIC` or `LIST`
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableType *string `json:"TableType,omitempty" name:"TableType"`

	// Table data interface description language (IDL) type, such as `PROTO` or `TDR`
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableIdlType *string `json:"TableIdlType,omitempty" name:"TableIdlType"`

	// ID of the table group where a table resides
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// Error message
	// Note: this field may return null, indicating that no valid values can be obtained.
	Error *ErrorInfo `json:"Error,omitempty" name:"Error"`

	// Task ID list, which is valid for the API that creates multiple tasks
	// Note: this field may return null, indicating that no valid values can be obtained.
	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds"`

	// Cluster operation application ID
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

type TableRollbackResultNew struct {
	// Table instance ID in the format of `tcaplus-3be64cbb`
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableInstanceId *string `json:"TableInstanceId,omitempty" name:"TableInstanceId"`

	// Task ID, which is valid for the API that creates one task
	// Note: this field may return null, indicating that no valid values can be obtained.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Table name
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// Table data structure type, such as `GENERIC` or `LIST`
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableType *string `json:"TableType,omitempty" name:"TableType"`

	// Table data interface description language (IDL) type, such as `PROTO` or `TDR`
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableIdlType *string `json:"TableIdlType,omitempty" name:"TableIdlType"`

	// ID of the table group where a table resides
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// Error message
	// Note: this field may return null, indicating that no valid values can be obtained.
	Error *ErrorInfo `json:"Error,omitempty" name:"Error"`

	// Task ID list, which is valid for the API that creates multiple tasks
	// Note: this field may return null, indicating that no valid values can be obtained.
	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds"`

	// ID of uploaded key file
	// Note: this field may return null, indicating that no valid values can be obtained.
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// Number of keys successfully verified
	// Note: this field may return null, indicating that no valid values can be obtained.
	SuccKeyNum *uint64 `json:"SuccKeyNum,omitempty" name:"SuccKeyNum"`

	// Total number of keys contained in key file
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalKeyNum *uint64 `json:"TotalKeyNum,omitempty" name:"TotalKeyNum"`
}

type TagInfoUnit struct {
	// Tag key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type TagsInfoOfCluster struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Tag information
	Tags []*TagInfoUnit `json:"Tags,omitempty" name:"Tags"`

	// Error message
	Error *ErrorInfo `json:"Error,omitempty" name:"Error"`
}

type TagsInfoOfTable struct {
	// Table instance ID
	TableInstanceId *string `json:"TableInstanceId,omitempty" name:"TableInstanceId"`

	// Table name
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// Table group ID
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// Tag information
	Tags []*TagInfoUnit `json:"Tags,omitempty" name:"Tags"`

	// Error message
	Error *ErrorInfo `json:"Error,omitempty" name:"Error"`
}

type TagsInfoOfTableGroup struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Table group ID
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// Tag information
	Tags []*TagInfoUnit `json:"Tags,omitempty" name:"Tags"`

	// Error message
	Error *ErrorInfo `json:"Error,omitempty" name:"Error"`
}

type TaskInfoNew struct {
	// Task ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Task type
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// ID of TcaplusDB internal transaction associated with task
	TransId *string `json:"TransId,omitempty" name:"TransId"`

	// ID of the cluster where a task resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Name of the cluster where a task resides
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// Task progress
	Progress *int64 `json:"Progress,omitempty" name:"Progress"`

	// Task creation time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Task last modified time
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Operator
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// Task details
	Content *string `json:"Content,omitempty" name:"Content"`
}

// Predefined struct for user
type UpdateApplyRequestParams struct {
	// Application status
	ApplyStatus []*ApplyStatus `json:"ApplyStatus,omitempty" name:"ApplyStatus"`
}

type UpdateApplyRequest struct {
	*tchttp.BaseRequest
	
	// Application status
	ApplyStatus []*ApplyStatus `json:"ApplyStatus,omitempty" name:"ApplyStatus"`
}

func (r *UpdateApplyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateApplyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplyStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateApplyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateApplyResponseParams struct {
	// List of updated applications
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	ApplyResults []*ApplyResult `json:"ApplyResults,omitempty" name:"ApplyResults"`

	// Total number of updated applications
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateApplyResponse struct {
	*tchttp.BaseResponse
	Response *UpdateApplyResponseParams `json:"Response"`
}

func (r *UpdateApplyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateApplyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyIdlFilesRequestParams struct {
	// ID of the cluster where to create a table
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// ID of the table group where to create a table
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// List of information of uploaded IDL files. Either this parameter or `NewIdlFiles` must be present
	ExistingIdlFiles []*IdlFileInfo `json:"ExistingIdlFiles,omitempty" name:"ExistingIdlFiles"`

	// List of information of IDL files to be uploaded. Either this parameter or `ExistingIdlFiles` must be present
	NewIdlFiles []*IdlFileInfo `json:"NewIdlFiles,omitempty" name:"NewIdlFiles"`
}

type VerifyIdlFilesRequest struct {
	*tchttp.BaseRequest
	
	// ID of the cluster where to create a table
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// ID of the table group where to create a table
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// List of information of uploaded IDL files. Either this parameter or `NewIdlFiles` must be present
	ExistingIdlFiles []*IdlFileInfo `json:"ExistingIdlFiles,omitempty" name:"ExistingIdlFiles"`

	// List of information of IDL files to be uploaded. Either this parameter or `ExistingIdlFiles` must be present
	NewIdlFiles []*IdlFileInfo `json:"NewIdlFiles,omitempty" name:"NewIdlFiles"`
}

func (r *VerifyIdlFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyIdlFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "TableGroupId")
	delete(f, "ExistingIdlFiles")
	delete(f, "NewIdlFiles")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyIdlFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyIdlFilesResponseParams struct {
	// Information list of all IDL files uploaded and verified in this request
	IdlFiles []*IdlFileInfo `json:"IdlFiles,omitempty" name:"IdlFiles"`

	// Number of valid tables parsed by reading IDL description file, excluding tables already created
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of valid tables parsed by reading IDL description file, excluding tables already created
	TableInfos []*ParsedTableInfoNew `json:"TableInfos,omitempty" name:"TableInfos"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type VerifyIdlFilesResponse struct {
	*tchttp.BaseResponse
	Response *VerifyIdlFilesResponseParams `json:"Response"`
}

func (r *VerifyIdlFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyIdlFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}