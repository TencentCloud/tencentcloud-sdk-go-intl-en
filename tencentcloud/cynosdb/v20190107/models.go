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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type Ability struct {
	// Whether secondary AZ is supported
	IsSupportSlaveZone *string `json:"IsSupportSlaveZone,omitnil,omitempty" name:"IsSupportSlaveZone"`

	// The reason why secondary AZ is not supported
	// Note: This field may return null, indicating that no valid values can be obtained.
	NonsupportSlaveZoneReason *string `json:"NonsupportSlaveZoneReason,omitnil,omitempty" name:"NonsupportSlaveZoneReason"`

	// Whether read-only instance is supported
	IsSupportRo *string `json:"IsSupportRo,omitnil,omitempty" name:"IsSupportRo"`

	// The reason why read-only instance is not supported
	// Note: This field may return null, indicating that no valid values can be obtained.
	NonsupportRoReason *string `json:"NonsupportRoReason,omitnil,omitempty" name:"NonsupportRoReason"`
}

type Account struct {
	// Database account name
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// Database account description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Update time
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Host
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// The max connections
	MaxUserConnections *int64 `json:"MaxUserConnections,omitnil,omitempty" name:"MaxUserConnections"`
}

// Predefined struct for user
type ActivateInstanceRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// List of instance IDs in the format of `cynosdbmysql-ins-n7ocdslw` as displayed in the TDSQL-C for MySQL console. You can use the instance list querying API to query the ID, i.e., the `InstanceId` value in the output parameters.
	InstanceIdList []*string `json:"InstanceIdList,omitnil,omitempty" name:"InstanceIdList"`
}

type ActivateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// List of instance IDs in the format of `cynosdbmysql-ins-n7ocdslw` as displayed in the TDSQL-C for MySQL console. You can use the instance list querying API to query the ID, i.e., the `InstanceId` value in the output parameters.
	InstanceIdList []*string `json:"InstanceIdList,omitnil,omitempty" name:"InstanceIdList"`
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
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Replica AZ
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`
}

type AddClusterSlaveZoneRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Replica AZ
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`
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
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Number of CPU cores
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Memory in GB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Number of added read-only instances. Value range: (0,16].
	ReadOnlyCount *int64 `json:"ReadOnlyCount,omitnil,omitempty" name:"ReadOnlyCount"`

	// Instance group ID, which will be used when you add an instance in an existing RO group. If this parameter is left empty, an RO group will be created. But it is not recommended to pass in this parameter for the current version, as this version has been disused.
	InstanceGrpId *string `json:"InstanceGrpId,omitnil,omitempty" name:"InstanceGrpId"`

	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID. If `VpcId` is set, `SubnetId` is required.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// The port used when adding an RO group. Value range: [0,65535).
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Instance name. String length range: [0,64).
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Whether to automatically select a voucher. 1: yes; 0: no. Default value: 0
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// Database type. Valid values: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Order source. String length range: [0,64).
	OrderSource *string `json:"OrderSource,omitnil,omitempty" name:"OrderSource"`

	// Transaction mode. Valid values: `0` (place and pay for an order), `1` (place an order)
	DealMode *int64 `json:"DealMode,omitnil,omitempty" name:"DealMode"`

	// Parameter template ID
	ParamTemplateId *int64 `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// Parameter list, which is valid only if `InstanceParams` is passed in to `ParamTemplateId`.
	InstanceParams []*ModifyParamItem `json:"InstanceParams,omitnil,omitempty" name:"InstanceParams"`

	// Security group ID. You can specify an security group when creating a read-only instance.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

type AddInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Number of CPU cores
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Memory in GB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Number of added read-only instances. Value range: (0,16].
	ReadOnlyCount *int64 `json:"ReadOnlyCount,omitnil,omitempty" name:"ReadOnlyCount"`

	// Instance group ID, which will be used when you add an instance in an existing RO group. If this parameter is left empty, an RO group will be created. But it is not recommended to pass in this parameter for the current version, as this version has been disused.
	InstanceGrpId *string `json:"InstanceGrpId,omitnil,omitempty" name:"InstanceGrpId"`

	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID. If `VpcId` is set, `SubnetId` is required.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// The port used when adding an RO group. Value range: [0,65535).
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Instance name. String length range: [0,64).
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Whether to automatically select a voucher. 1: yes; 0: no. Default value: 0
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// Database type. Valid values: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Order source. String length range: [0,64).
	OrderSource *string `json:"OrderSource,omitnil,omitempty" name:"OrderSource"`

	// Transaction mode. Valid values: `0` (place and pay for an order), `1` (place an order)
	DealMode *int64 `json:"DealMode,omitnil,omitempty" name:"DealMode"`

	// Parameter template ID
	ParamTemplateId *int64 `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// Parameter list, which is valid only if `InstanceParams` is passed in to `ParamTemplateId`.
	InstanceParams []*ModifyParamItem `json:"InstanceParams,omitnil,omitempty" name:"InstanceParams"`

	// Security group ID. You can specify an security group when creating a read-only instance.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
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
	TranId *string `json:"TranId,omitnil,omitempty" name:"TranId"`

	// Pay-as-You-Go order ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// List of IDs of delivered resources
	// Note: this field may return null, indicating that no valid values can be obtained.
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// Big order ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	BigDealIds []*string `json:"BigDealIds,omitnil,omitempty" name:"BigDealIds"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// IP address
	IP *string `json:"IP,omitnil,omitempty" name:"IP"`

	// Port
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`
}

type AuditRuleFilters struct {
	// Audit rule
	RuleFilters []*RuleFilters `json:"RuleFilters,omitnil,omitempty" name:"RuleFilters"`
}

type AuditRuleTemplateInfo struct {
	// Rule template ID
	RuleTemplateId *string `json:"RuleTemplateId,omitnil,omitempty" name:"RuleTemplateId"`

	// Rule template name
	RuleTemplateName *string `json:"RuleTemplateName,omitnil,omitempty" name:"RuleTemplateName"`

	// Filter of the rule template
	RuleFilters []*RuleFilters `json:"RuleFilters,omitnil,omitempty" name:"RuleFilters"`

	// Description of a rule template
	// Note: This field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Creation time of a rule template
	CreateAt *string `json:"CreateAt,omitnil,omitempty" name:"CreateAt"`
}

type BackupFileInfo struct {
	// Snapshot file ID, which is deprecated. You need to use `BackupId`.
	SnapshotId *uint64 `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// Backup file name
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// Backup file size
	FileSize *uint64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// Backup start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Backup end time
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`

	// Backup type. Valid values: `snapshot` (snapshot backup), `logic` (logic backup).
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// Back mode. auto: auto backup; manual: manual backup
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// Backup file status. success: backup succeeded; fail: backup failed; creating: creating backup file; deleting: deleting backup file
	BackupStatus *string `json:"BackupStatus,omitnil,omitempty" name:"BackupStatus"`

	// Backup file time
	SnapshotTime *string `json:"SnapshotTime,omitnil,omitempty" name:"SnapshotTime"`

	// Backup ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupId *int64 `json:"BackupId,omitnil,omitempty" name:"BackupId"`


	SnapShotType *string `json:"SnapShotType,omitnil,omitempty" name:"SnapShotType"`

	// Backup file alias
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`
}

type BillingResourceInfo struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Instance ID list
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Order ID
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`
}

// Predefined struct for user
type BindClusterResourcePackagesRequestParams struct {
	// The unique ID of a resource pack
	PackageIds []*string `json:"PackageIds,omitnil,omitempty" name:"PackageIds"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type BindClusterResourcePackagesRequest struct {
	*tchttp.BaseRequest
	
	// The unique ID of a resource pack
	PackageIds []*string `json:"PackageIds,omitnil,omitempty" name:"PackageIds"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *BindClusterResourcePackagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindClusterResourcePackagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PackageIds")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindClusterResourcePackagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindClusterResourcePackagesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BindClusterResourcePackagesResponse struct {
	*tchttp.BaseResponse
	Response *BindClusterResourcePackagesResponseParams `json:"Response"`
}

func (r *BindClusterResourcePackagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindClusterResourcePackagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindInstanceInfo struct {
	// ID of the bound cluster
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Region of the instance bound Note: This field may return null, indicating that no valid values can be obtained.
	InstanceRegion *string `json:"InstanceRegion,omitnil,omitempty" name:"InstanceRegion"`

	// Type of the instance bound Note: This field may return null, indicating that no valid values can be obtained.
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// ID of the instance in the bound cluster
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExtendIds []*string `json:"ExtendIds,omitnil,omitempty" name:"ExtendIds"`
}

type BinlogItem struct {
	// Binlog filename
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// File size in bytes
	FileSize *int64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// Transaction start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Transaction end time
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`

	// Binlog file ID
	BinlogId *int64 `json:"BinlogId,omitnil,omitempty" name:"BinlogId"`
}

// Predefined struct for user
type CloseAuditServiceRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type CloseAuditServiceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type CloseClusterPasswordComplexityRequestParams struct {
	// Cluster IDs in array
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`
}

type CloseClusterPasswordComplexityRequest struct {
	*tchttp.BaseRequest
	
	// Cluster IDs in array
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`
}

func (r *CloseClusterPasswordComplexityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseClusterPasswordComplexityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseClusterPasswordComplexityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseClusterPasswordComplexityResponseParams struct {
	// Task flow ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloseClusterPasswordComplexityResponse struct {
	*tchttp.BaseResponse
	Response *CloseClusterPasswordComplexityResponseParams `json:"Response"`
}

func (r *CloseClusterPasswordComplexityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseClusterPasswordComplexityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseProxyRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Whether only to disable read/write separation. Valid values: `true`, `false`.
	OnlyCloseRW *bool `json:"OnlyCloseRW,omitnil,omitempty" name:"OnlyCloseRW"`
}

type CloseProxyRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Whether only to disable read/write separation. Valid values: `true`, `false`.
	OnlyCloseRW *bool `json:"OnlyCloseRW,omitnil,omitempty" name:"OnlyCloseRW"`
}

func (r *CloseProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ProxyGroupId")
	delete(f, "OnlyCloseRW")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseProxyResponseParams struct {
	// Async flow ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Async task ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloseProxyResponse struct {
	*tchttp.BaseResponse
	Response *CloseProxyResponseParams `json:"Response"`
}

func (r *CloseProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseWanRequestParams struct {
	// Instance group ID
	InstanceGrpId *string `json:"InstanceGrpId,omitnil,omitempty" name:"InstanceGrpId"`
}

type CloseWanRequest struct {
	*tchttp.BaseRequest
	
	// Instance group ID
	InstanceGrpId *string `json:"InstanceGrpId,omitnil,omitempty" name:"InstanceGrpId"`
}

func (r *CloseWanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseWanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceGrpId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseWanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseWanResponseParams struct {
	// Task flow ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloseWanResponse struct {
	*tchttp.BaseResponse
	Response *CloseWanResponseParams `json:"Response"`
}

func (r *CloseWanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseWanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterInstanceDetail struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Engine type
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Instance status
	InstanceStatus *string `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// Instance status description
	InstanceStatusDesc *string `json:"InstanceStatusDesc,omitnil,omitempty" name:"InstanceStatusDesc"`

	// Number of CPU cores
	InstanceCpu *int64 `json:"InstanceCpu,omitnil,omitempty" name:"InstanceCpu"`

	// Memory
	InstanceMemory *int64 `json:"InstanceMemory,omitnil,omitempty" name:"InstanceMemory"`

	// Disk
	InstanceStorage *int64 `json:"InstanceStorage,omitnil,omitempty" name:"InstanceStorage"`

	// Instance role
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// Execution start time in seconds from 0:00	
	// Note: This field may return null, indicating that no valid values can be obtained.
	MaintainStartTime *int64 `json:"MaintainStartTime,omitnil,omitempty" name:"MaintainStartTime"`

	// Duration in seconds	
	// Note: This field may return null, indicating that no valid values can be obtained.
	MaintainDuration *int64 `json:"MaintainDuration,omitnil,omitempty" name:"MaintainDuration"`

	// Execution time. Valid values: `Mon`, `Tue`, `Wed`, `Thu`, `Fri`, Sat`, `Sun`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MaintainWeekDays []*string `json:"MaintainWeekDays,omitnil,omitempty" name:"MaintainWeekDays"`

	// Serverless instance enablement status
	// Note: This field may return null, indicating that no valid values can be obtained.
	ServerlessStatus *string `json:"ServerlessStatus,omitnil,omitempty" name:"ServerlessStatus"`


	InstanceTasks []*ObjectTask `json:"InstanceTasks,omitnil,omitempty" name:"InstanceTasks"`


	InstanceDeviceType *string `json:"InstanceDeviceType,omitnil,omitempty" name:"InstanceDeviceType"`
}

// Predefined struct for user
type CopyClusterPasswordComplexityRequestParams struct {
	// A parameter used to replicate the array of cluster IDs
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// Cluster ID
	SourceClusterId *string `json:"SourceClusterId,omitnil,omitempty" name:"SourceClusterId"`
}

type CopyClusterPasswordComplexityRequest struct {
	*tchttp.BaseRequest
	
	// A parameter used to replicate the array of cluster IDs
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// Cluster ID
	SourceClusterId *string `json:"SourceClusterId,omitnil,omitempty" name:"SourceClusterId"`
}

func (r *CopyClusterPasswordComplexityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyClusterPasswordComplexityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterIds")
	delete(f, "SourceClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CopyClusterPasswordComplexityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CopyClusterPasswordComplexityResponseParams struct {
	// Task flow ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CopyClusterPasswordComplexityResponse struct {
	*tchttp.BaseResponse
	Response *CopyClusterPasswordComplexityResponseParams `json:"Response"`
}

func (r *CopyClusterPasswordComplexityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyClusterPasswordComplexityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccountsRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// List of new accounts
	Accounts []*NewAccount `json:"Accounts,omitnil,omitempty" name:"Accounts"`
}

type CreateAccountsRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// List of new accounts
	Accounts []*NewAccount `json:"Accounts,omitnil,omitempty" name:"Accounts"`
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
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	RuleFilters []*RuleFilters `json:"RuleFilters,omitnil,omitempty" name:"RuleFilters"`

	// Rule template name
	RuleTemplateName *string `json:"RuleTemplateName,omitnil,omitempty" name:"RuleTemplateName"`

	// Rule template description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateAuditRuleTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Audit rule
	RuleFilters []*RuleFilters `json:"RuleFilters,omitnil,omitempty" name:"RuleFilters"`

	// Rule template name
	RuleTemplateName *string `json:"RuleTemplateName,omitnil,omitempty" name:"RuleTemplateName"`

	// Rule template description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
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
	RuleTemplateId *string `json:"RuleTemplateId,omitnil,omitempty" name:"RuleTemplateId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Backup type. Valid values: `logic` (logic backup), `snapshot` (physical backup)
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// Backup database, which is valid when `BackupType` is `logic`.
	BackupDatabases []*string `json:"BackupDatabases,omitnil,omitempty" name:"BackupDatabases"`

	// Backup table, which is valid when `BackupType` is `logic`.
	BackupTables []*DatabaseTables `json:"BackupTables,omitnil,omitempty" name:"BackupTables"`

	// Backup name
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`
}

type CreateBackupRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Backup type. Valid values: `logic` (logic backup), `snapshot` (physical backup)
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// Backup database, which is valid when `BackupType` is `logic`.
	BackupDatabases []*string `json:"BackupDatabases,omitnil,omitempty" name:"BackupDatabases"`

	// Backup table, which is valid when `BackupType` is `logic`.
	BackupTables []*DatabaseTables `json:"BackupTables,omitnil,omitempty" name:"BackupTables"`

	// Backup name
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`
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
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateClusterDatabaseRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database name
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Character set
	CharacterSet *string `json:"CharacterSet,omitnil,omitempty" name:"CharacterSet"`

	// Collation
	CollateRule *string `json:"CollateRule,omitnil,omitempty" name:"CollateRule"`

	// Host permissions of the authorized user
	UserHostPrivileges []*UserHostPrivilege `json:"UserHostPrivileges,omitnil,omitempty" name:"UserHostPrivileges"`

	// Remarks
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateClusterDatabaseRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database name
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Character set
	CharacterSet *string `json:"CharacterSet,omitnil,omitempty" name:"CharacterSet"`

	// Collation
	CollateRule *string `json:"CollateRule,omitnil,omitempty" name:"CollateRule"`

	// Host permissions of the authorized user
	UserHostPrivileges []*UserHostPrivilege `json:"UserHostPrivileges,omitnil,omitempty" name:"UserHostPrivileges"`

	// Remarks
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateClusterDatabaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterDatabaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "DbName")
	delete(f, "CharacterSet")
	delete(f, "CollateRule")
	delete(f, "UserHostPrivileges")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClusterDatabaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterDatabaseResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateClusterDatabaseResponse struct {
	*tchttp.BaseResponse
	Response *CreateClusterDatabaseResponseParams `json:"Response"`
}

func (r *CreateClusterDatabaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterDatabaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClustersRequestParams struct {
	// AZ
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Database type. Valid values: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Database version. Valid values: 
	// <li> Valid values for `MYSQL`: 5.7 and 8.0 </li>
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`

	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// It is required when `DbMode` is set to `NORMAL` or left empty.
	// Number of CPU cores of normal instance
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// It is required when `DbMode` is set to `NORMAL` or left empty.
	// Memory of a non-serverless instance in GB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// This parameter has been deprecated.
	// Storage capacity in GB
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// Cluster name, which can contain less than 64 letters, digits, or symbols (-_.).
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// Account password, which must contain 8-64 characters in at least three of the following four types: uppercase letters, lowercase letters, digits, and symbols (~!@#$%^&*_-+=`|\(){}[]:;'<>,.?/).
	AdminPassword *string `json:"AdminPassword,omitnil,omitempty" name:"AdminPassword"`

	// Port. Valid range: [0, 65535). Default value: 3306
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Billing mode. `0`: pay-as-you-go; `1`: monthly subscription. Default value: `0`
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Number of purchased clusters. Valid range: [1,50]. Default value: 1
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// Rollback type:
	// noneRollback: no rollback;
	// snapRollback: rollback by snapshot;
	// timeRollback: rollback by time point
	RollbackStrategy *string `json:"RollbackStrategy,omitnil,omitempty" name:"RollbackStrategy"`

	// `snapshotId` for snapshot rollback or `queryId` for time point rollback. 0 indicates to determine whether the time point is valid
	RollbackId *uint64 `json:"RollbackId,omitnil,omitempty" name:"RollbackId"`

	// The source cluster ID passed in during rollback to find the source `poolId`
	OriginalClusterId *string `json:"OriginalClusterId,omitnil,omitempty" name:"OriginalClusterId"`

	// Specified time for time point rollback or snapshot time for snapshot rollback
	ExpectTime *string `json:"ExpectTime,omitnil,omitempty" name:"ExpectTime"`

	// This parameter has been deprecated.
	// Specified allowed time range for time point rollback
	ExpectTimeThresh *uint64 `json:"ExpectTimeThresh,omitnil,omitempty" name:"ExpectTimeThresh"`

	// Storage upper limit of normal instance in GB
	// If `DbType` is `MYSQL` and the storage billing mode is monthly subscription, the parameter value can’t exceed the maximum storage corresponding to the CPU and memory specifications.
	StorageLimit *int64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`

	// Number of Instances. Valid range: (0,16]
	InstanceCount *int64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// Purchase duration of monthly subscription plan
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// Duration unit of monthly subscription. Valid values: `s`, `d`, `m`, `y`
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// Whether auto-renewal is enabled for monthly subscription plan. Default value: `0`
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// Whether to automatically select a voucher. `1`: yes; `0`: no. Default value: `0`
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// Number of instances (this parameter has been disused and is retained only for compatibility with existing instances)
	HaCount *int64 `json:"HaCount,omitnil,omitempty" name:"HaCount"`

	// Order source
	OrderSource *string `json:"OrderSource,omitnil,omitempty" name:"OrderSource"`

	// Array of tags to be bound to the created cluster
	ResourceTags []*Tag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// Database type
	// Valid values when `DbType` is `MYSQL` (default value: `NORMAL`):
	// <li>NORMAL</li>
	// <li>SERVERLESS</li>
	DbMode *string `json:"DbMode,omitnil,omitempty" name:"DbMode"`

	// This parameter is required if `DbMode` is `SERVERLESS`.
	// Minimum number of CPU cores. For the value range, see the returned result of `DescribeServerlessInstanceSpecs`.
	MinCpu *float64 `json:"MinCpu,omitnil,omitempty" name:"MinCpu"`

	// This parameter is required if `DbMode` is `SERVERLESS`.
	// Maximum number of CPU cores. For the value range, see the returned result of `DescribeServerlessInstanceSpecs`.
	MaxCpu *float64 `json:"MaxCpu,omitnil,omitempty" name:"MaxCpu"`

	// This parameter specifies whether the cluster will be automatically paused if `DbMode` is `SERVERLESS`. Valid values:
	// <li>yes</li>
	// <li>no</li>
	// Default value: yes
	AutoPause *string `json:"AutoPause,omitnil,omitempty" name:"AutoPause"`

	// This parameter specifies the delay for automatic cluster pause in seconds if `DbMode` is `SERVERLESS`. Value range: [600,691200]
	// Default value: `600`
	AutoPauseDelay *int64 `json:"AutoPauseDelay,omitnil,omitempty" name:"AutoPauseDelay"`

	// The billing mode of cluster storage. Valid values: `0` (pay-as-you-go), `1` (monthly subscription). Default value: `0`.
	// If `DbType` is `MYSQL` and the billing mode of cluster compute is pay-as-you-go (or the `DbMode` is `SERVERLESS`), the billing mode of cluster storage must be pay-as-you-go.
	// Clusters with storage billed in monthly subscription can’t be cloned or rolled back.
	StoragePayMode *int64 `json:"StoragePayMode,omitnil,omitempty" name:"StoragePayMode"`

	// Array of security group IDs
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Array of alarm policy IDs
	AlarmPolicyIds []*string `json:"AlarmPolicyIds,omitnil,omitempty" name:"AlarmPolicyIds"`

	// Array of parameters. Valid values: `character_set_server` (utf8｜latin1｜gbk｜utf8mb4), `lower_case_table_names`. 0: case-sensitive; 1: case-insensitive).
	ClusterParams []*ParamItem `json:"ClusterParams,omitnil,omitempty" name:"ClusterParams"`

	// Transaction mode. Valid values: `0` (place and pay for an order), `1` (place an order)
	DealMode *int64 `json:"DealMode,omitnil,omitempty" name:"DealMode"`

	// Parameter template ID, which can be obtained by querying parameter template information “DescribeParamTemplates”
	ParamTemplateId *int64 `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// Multi-AZ address
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// Instance initialization configuration information, which is used to select instances with different specifications when purchasing a cluster.
	InstanceInitInfos []*InstanceInitInfo `json:"InstanceInitInfos,omitnil,omitempty" name:"InstanceInitInfos"`
}

type CreateClustersRequest struct {
	*tchttp.BaseRequest
	
	// AZ
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Database type. Valid values: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Database version. Valid values: 
	// <li> Valid values for `MYSQL`: 5.7 and 8.0 </li>
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`

	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// It is required when `DbMode` is set to `NORMAL` or left empty.
	// Number of CPU cores of normal instance
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// It is required when `DbMode` is set to `NORMAL` or left empty.
	// Memory of a non-serverless instance in GB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// This parameter has been deprecated.
	// Storage capacity in GB
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// Cluster name, which can contain less than 64 letters, digits, or symbols (-_.).
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// Account password, which must contain 8-64 characters in at least three of the following four types: uppercase letters, lowercase letters, digits, and symbols (~!@#$%^&*_-+=`|\(){}[]:;'<>,.?/).
	AdminPassword *string `json:"AdminPassword,omitnil,omitempty" name:"AdminPassword"`

	// Port. Valid range: [0, 65535). Default value: 3306
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Billing mode. `0`: pay-as-you-go; `1`: monthly subscription. Default value: `0`
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Number of purchased clusters. Valid range: [1,50]. Default value: 1
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// Rollback type:
	// noneRollback: no rollback;
	// snapRollback: rollback by snapshot;
	// timeRollback: rollback by time point
	RollbackStrategy *string `json:"RollbackStrategy,omitnil,omitempty" name:"RollbackStrategy"`

	// `snapshotId` for snapshot rollback or `queryId` for time point rollback. 0 indicates to determine whether the time point is valid
	RollbackId *uint64 `json:"RollbackId,omitnil,omitempty" name:"RollbackId"`

	// The source cluster ID passed in during rollback to find the source `poolId`
	OriginalClusterId *string `json:"OriginalClusterId,omitnil,omitempty" name:"OriginalClusterId"`

	// Specified time for time point rollback or snapshot time for snapshot rollback
	ExpectTime *string `json:"ExpectTime,omitnil,omitempty" name:"ExpectTime"`

	// This parameter has been deprecated.
	// Specified allowed time range for time point rollback
	ExpectTimeThresh *uint64 `json:"ExpectTimeThresh,omitnil,omitempty" name:"ExpectTimeThresh"`

	// Storage upper limit of normal instance in GB
	// If `DbType` is `MYSQL` and the storage billing mode is monthly subscription, the parameter value can’t exceed the maximum storage corresponding to the CPU and memory specifications.
	StorageLimit *int64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`

	// Number of Instances. Valid range: (0,16]
	InstanceCount *int64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// Purchase duration of monthly subscription plan
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// Duration unit of monthly subscription. Valid values: `s`, `d`, `m`, `y`
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// Whether auto-renewal is enabled for monthly subscription plan. Default value: `0`
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// Whether to automatically select a voucher. `1`: yes; `0`: no. Default value: `0`
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// Number of instances (this parameter has been disused and is retained only for compatibility with existing instances)
	HaCount *int64 `json:"HaCount,omitnil,omitempty" name:"HaCount"`

	// Order source
	OrderSource *string `json:"OrderSource,omitnil,omitempty" name:"OrderSource"`

	// Array of tags to be bound to the created cluster
	ResourceTags []*Tag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// Database type
	// Valid values when `DbType` is `MYSQL` (default value: `NORMAL`):
	// <li>NORMAL</li>
	// <li>SERVERLESS</li>
	DbMode *string `json:"DbMode,omitnil,omitempty" name:"DbMode"`

	// This parameter is required if `DbMode` is `SERVERLESS`.
	// Minimum number of CPU cores. For the value range, see the returned result of `DescribeServerlessInstanceSpecs`.
	MinCpu *float64 `json:"MinCpu,omitnil,omitempty" name:"MinCpu"`

	// This parameter is required if `DbMode` is `SERVERLESS`.
	// Maximum number of CPU cores. For the value range, see the returned result of `DescribeServerlessInstanceSpecs`.
	MaxCpu *float64 `json:"MaxCpu,omitnil,omitempty" name:"MaxCpu"`

	// This parameter specifies whether the cluster will be automatically paused if `DbMode` is `SERVERLESS`. Valid values:
	// <li>yes</li>
	// <li>no</li>
	// Default value: yes
	AutoPause *string `json:"AutoPause,omitnil,omitempty" name:"AutoPause"`

	// This parameter specifies the delay for automatic cluster pause in seconds if `DbMode` is `SERVERLESS`. Value range: [600,691200]
	// Default value: `600`
	AutoPauseDelay *int64 `json:"AutoPauseDelay,omitnil,omitempty" name:"AutoPauseDelay"`

	// The billing mode of cluster storage. Valid values: `0` (pay-as-you-go), `1` (monthly subscription). Default value: `0`.
	// If `DbType` is `MYSQL` and the billing mode of cluster compute is pay-as-you-go (or the `DbMode` is `SERVERLESS`), the billing mode of cluster storage must be pay-as-you-go.
	// Clusters with storage billed in monthly subscription can’t be cloned or rolled back.
	StoragePayMode *int64 `json:"StoragePayMode,omitnil,omitempty" name:"StoragePayMode"`

	// Array of security group IDs
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Array of alarm policy IDs
	AlarmPolicyIds []*string `json:"AlarmPolicyIds,omitnil,omitempty" name:"AlarmPolicyIds"`

	// Array of parameters. Valid values: `character_set_server` (utf8｜latin1｜gbk｜utf8mb4), `lower_case_table_names`. 0: case-sensitive; 1: case-insensitive).
	ClusterParams []*ParamItem `json:"ClusterParams,omitnil,omitempty" name:"ClusterParams"`

	// Transaction mode. Valid values: `0` (place and pay for an order), `1` (place an order)
	DealMode *int64 `json:"DealMode,omitnil,omitempty" name:"DealMode"`

	// Parameter template ID, which can be obtained by querying parameter template information “DescribeParamTemplates”
	ParamTemplateId *int64 `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// Multi-AZ address
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// Instance initialization configuration information, which is used to select instances with different specifications when purchasing a cluster.
	InstanceInitInfos []*InstanceInitInfo `json:"InstanceInitInfos,omitnil,omitempty" name:"InstanceInitInfos"`
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
	TranId *string `json:"TranId,omitnil,omitempty" name:"TranId"`

	// Order ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// List of resource IDs (This field has been deprecated. You need to use `dealNames` in the `DescribeResourcesByDealName` API to get resource IDs.)
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// List of cluster IDs (This field has been deprecated. You need to use `dealNames` in the `DescribeResourcesByDealName` API to get cluster IDs.)
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// Big order ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	BigDealIds []*string `json:"BigDealIds,omitnil,omitempty" name:"BigDealIds"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type CreateParamTemplateRequestParams struct {
	// Template name
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// MySQL version number
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// Template description
	TemplateDescription *string `json:"TemplateDescription,omitnil,omitempty" name:"TemplateDescription"`

	// ID of the template to be copied
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Database type. Valid values:  `NORMAL` (default), `SERVERLESS`.
	DbMode *string `json:"DbMode,omitnil,omitempty" name:"DbMode"`

	// List of the parameters
	ParamList []*ParamItem `json:"ParamList,omitnil,omitempty" name:"ParamList"`
}

type CreateParamTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Template name
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// MySQL version number
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// Template description
	TemplateDescription *string `json:"TemplateDescription,omitnil,omitempty" name:"TemplateDescription"`

	// ID of the template to be copied
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Database type. Valid values:  `NORMAL` (default), `SERVERLESS`.
	DbMode *string `json:"DbMode,omitnil,omitempty" name:"DbMode"`

	// List of the parameters
	ParamList []*ParamItem `json:"ParamList,omitnil,omitempty" name:"ParamList"`
}

func (r *CreateParamTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateParamTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateName")
	delete(f, "EngineVersion")
	delete(f, "TemplateDescription")
	delete(f, "TemplateId")
	delete(f, "DbMode")
	delete(f, "ParamList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateParamTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateParamTemplateResponseParams struct {
	// Template ID
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateParamTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateParamTemplateResponseParams `json:"Response"`
}

func (r *CreateParamTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateParamTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProxyEndPointRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// VPC ID, which is the same as that of the cluster by default.
	UniqueVpcId *string `json:"UniqueVpcId,omitnil,omitempty" name:"UniqueVpcId"`

	// VPCe subnet ID, which is the same as that of the cluster by default.
	UniqueSubnetId *string `json:"UniqueSubnetId,omitnil,omitempty" name:"UniqueSubnetId"`

	// Connection pool type. Valid value: `SessionConnectionPool` (session-level connection pool)
	ConnectionPoolType *string `json:"ConnectionPoolType,omitnil,omitempty" name:"ConnectionPoolType"`

	// Whether to enable connection pool. Valid value: `yes` (enable), `no` (disable).
	OpenConnectionPool *string `json:"OpenConnectionPool,omitnil,omitempty" name:"OpenConnectionPool"`

	// Connection pool threshold in seconds
	ConnectionPoolTimeOut *int64 `json:"ConnectionPoolTimeOut,omitnil,omitempty" name:"ConnectionPoolTimeOut"`

	// Array of security group IDs
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// VIP information
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Weight mode. 
	// Valid values: `system` (system-assigned), `custom` (custom).
	WeightMode *string `json:"WeightMode,omitnil,omitempty" name:"WeightMode"`

	// Whether to automatically add read-only instance. Valid value: `yes`, `no`.
	AutoAddRo *string `json:"AutoAddRo,omitnil,omitempty" name:"AutoAddRo"`

	// Whether to enable failover
	FailOver *string `json:"FailOver,omitnil,omitempty" name:"FailOver"`

	// Consistency type. Valid values: 
	// `eventual`, `global`, `session`.
	ConsistencyType *string `json:"ConsistencyType,omitnil,omitempty" name:"ConsistencyType"`

	// Read-write attribute. Valid values: 
	// `READWRITE`, `READONLY`.
	RwType *string `json:"RwType,omitnil,omitempty" name:"RwType"`

	// Consistency timeout period
	ConsistencyTimeOut *int64 `json:"ConsistencyTimeOut,omitnil,omitempty" name:"ConsistencyTimeOut"`

	// Transaction split
	TransSplit *bool `json:"TransSplit,omitnil,omitempty" name:"TransSplit"`

	// Connection mode. Valid values:
	// `nearby`, `balance`.
	AccessMode *string `json:"AccessMode,omitnil,omitempty" name:"AccessMode"`

	// Instance weight
	InstanceWeights []*ProxyInstanceWeight `json:"InstanceWeights,omitnil,omitempty" name:"InstanceWeights"`
}

type CreateProxyEndPointRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// VPC ID, which is the same as that of the cluster by default.
	UniqueVpcId *string `json:"UniqueVpcId,omitnil,omitempty" name:"UniqueVpcId"`

	// VPCe subnet ID, which is the same as that of the cluster by default.
	UniqueSubnetId *string `json:"UniqueSubnetId,omitnil,omitempty" name:"UniqueSubnetId"`

	// Connection pool type. Valid value: `SessionConnectionPool` (session-level connection pool)
	ConnectionPoolType *string `json:"ConnectionPoolType,omitnil,omitempty" name:"ConnectionPoolType"`

	// Whether to enable connection pool. Valid value: `yes` (enable), `no` (disable).
	OpenConnectionPool *string `json:"OpenConnectionPool,omitnil,omitempty" name:"OpenConnectionPool"`

	// Connection pool threshold in seconds
	ConnectionPoolTimeOut *int64 `json:"ConnectionPoolTimeOut,omitnil,omitempty" name:"ConnectionPoolTimeOut"`

	// Array of security group IDs
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// VIP information
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Weight mode. 
	// Valid values: `system` (system-assigned), `custom` (custom).
	WeightMode *string `json:"WeightMode,omitnil,omitempty" name:"WeightMode"`

	// Whether to automatically add read-only instance. Valid value: `yes`, `no`.
	AutoAddRo *string `json:"AutoAddRo,omitnil,omitempty" name:"AutoAddRo"`

	// Whether to enable failover
	FailOver *string `json:"FailOver,omitnil,omitempty" name:"FailOver"`

	// Consistency type. Valid values: 
	// `eventual`, `global`, `session`.
	ConsistencyType *string `json:"ConsistencyType,omitnil,omitempty" name:"ConsistencyType"`

	// Read-write attribute. Valid values: 
	// `READWRITE`, `READONLY`.
	RwType *string `json:"RwType,omitnil,omitempty" name:"RwType"`

	// Consistency timeout period
	ConsistencyTimeOut *int64 `json:"ConsistencyTimeOut,omitnil,omitempty" name:"ConsistencyTimeOut"`

	// Transaction split
	TransSplit *bool `json:"TransSplit,omitnil,omitempty" name:"TransSplit"`

	// Connection mode. Valid values:
	// `nearby`, `balance`.
	AccessMode *string `json:"AccessMode,omitnil,omitempty" name:"AccessMode"`

	// Instance weight
	InstanceWeights []*ProxyInstanceWeight `json:"InstanceWeights,omitnil,omitempty" name:"InstanceWeights"`
}

func (r *CreateProxyEndPointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProxyEndPointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "UniqueVpcId")
	delete(f, "UniqueSubnetId")
	delete(f, "ConnectionPoolType")
	delete(f, "OpenConnectionPool")
	delete(f, "ConnectionPoolTimeOut")
	delete(f, "SecurityGroupIds")
	delete(f, "Description")
	delete(f, "Vip")
	delete(f, "WeightMode")
	delete(f, "AutoAddRo")
	delete(f, "FailOver")
	delete(f, "ConsistencyType")
	delete(f, "RwType")
	delete(f, "ConsistencyTimeOut")
	delete(f, "TransSplit")
	delete(f, "AccessMode")
	delete(f, "InstanceWeights")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProxyEndPointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProxyEndPointResponseParams struct {
	// Async flow ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Async task ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Database proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateProxyEndPointResponse struct {
	*tchttp.BaseResponse
	Response *CreateProxyEndPointResponseParams `json:"Response"`
}

func (r *CreateProxyEndPointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProxyEndPointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProxyRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Number of CPU cores
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Memory
	Mem *int64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// VPC ID, which is the same as that of the cluster by default.
	UniqueVpcId *string `json:"UniqueVpcId,omitnil,omitempty" name:"UniqueVpcId"`

	// VPC subnet ID, which is the same as that of the cluster by default.
	UniqueSubnetId *string `json:"UniqueSubnetId,omitnil,omitempty" name:"UniqueSubnetId"`

	// Number of nodes in the proxy group
	ProxyCount *int64 `json:"ProxyCount,omitnil,omitempty" name:"ProxyCount"`

	// Connection pool type. Valid value: `SessionConnectionPool` (session-level connection pool)
	ConnectionPoolType *string `json:"ConnectionPoolType,omitnil,omitempty" name:"ConnectionPoolType"`

	// Whether to enable connection pool. Valid value: `yes` (enable), `no` (disable).
	OpenConnectionPool *string `json:"OpenConnectionPool,omitnil,omitempty" name:"OpenConnectionPool"`

	// Connection pool threshold in seconds
	ConnectionPoolTimeOut *int64 `json:"ConnectionPoolTimeOut,omitnil,omitempty" name:"ConnectionPoolTimeOut"`

	// Array of security group IDs
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Database node information
	ProxyZones []*ProxyZone `json:"ProxyZones,omitnil,omitempty" name:"ProxyZones"`
}

type CreateProxyRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Number of CPU cores
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Memory
	Mem *int64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// VPC ID, which is the same as that of the cluster by default.
	UniqueVpcId *string `json:"UniqueVpcId,omitnil,omitempty" name:"UniqueVpcId"`

	// VPC subnet ID, which is the same as that of the cluster by default.
	UniqueSubnetId *string `json:"UniqueSubnetId,omitnil,omitempty" name:"UniqueSubnetId"`

	// Number of nodes in the proxy group
	ProxyCount *int64 `json:"ProxyCount,omitnil,omitempty" name:"ProxyCount"`

	// Connection pool type. Valid value: `SessionConnectionPool` (session-level connection pool)
	ConnectionPoolType *string `json:"ConnectionPoolType,omitnil,omitempty" name:"ConnectionPoolType"`

	// Whether to enable connection pool. Valid value: `yes` (enable), `no` (disable).
	OpenConnectionPool *string `json:"OpenConnectionPool,omitnil,omitempty" name:"OpenConnectionPool"`

	// Connection pool threshold in seconds
	ConnectionPoolTimeOut *int64 `json:"ConnectionPoolTimeOut,omitnil,omitempty" name:"ConnectionPoolTimeOut"`

	// Array of security group IDs
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Database node information
	ProxyZones []*ProxyZone `json:"ProxyZones,omitnil,omitempty" name:"ProxyZones"`
}

func (r *CreateProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Cpu")
	delete(f, "Mem")
	delete(f, "UniqueVpcId")
	delete(f, "UniqueSubnetId")
	delete(f, "ProxyCount")
	delete(f, "ConnectionPoolType")
	delete(f, "OpenConnectionPool")
	delete(f, "ConnectionPoolTimeOut")
	delete(f, "SecurityGroupIds")
	delete(f, "Description")
	delete(f, "ProxyZones")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProxyResponseParams struct {
	// Async flow ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Async task ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Database proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateProxyResponse struct {
	*tchttp.BaseResponse
	Response *CreateProxyResponseParams `json:"Response"`
}

func (r *CreateProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateResourcePackageRequestParams struct {
	// Instance type
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Region of the resource pack. Valid values: `China` (Chinese mainland), `overseas` (outside Chinese mainland).
	PackageRegion *string `json:"PackageRegion,omitnil,omitempty" name:"PackageRegion"`

	// Resource pack type. Valid values: `CCU` (compute resource pack), `DISK` (storage resource pack).
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// Resource pack edition. Valid values: `base` (basic edition), `common` (general edition), `enterprise` (enterprise edition).
	PackageVersion *string `json:"PackageVersion,omitnil,omitempty" name:"PackageVersion"`

	// Resource pack size. Unit of the compute resource pack: Ten thousand.  Unit of the storage resource pack:  GB
	PackageSpec *float64 `json:"PackageSpec,omitnil,omitempty" name:"PackageSpec"`

	// Validity period of a resource pack in days
	ExpireDay *int64 `json:"ExpireDay,omitnil,omitempty" name:"ExpireDay"`

	// Number of the resource packs to be purchased
	PackageCount *int64 `json:"PackageCount,omitnil,omitempty" name:"PackageCount"`

	// Resource pack name
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`
}

type CreateResourcePackageRequest struct {
	*tchttp.BaseRequest
	
	// Instance type
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Region of the resource pack. Valid values: `China` (Chinese mainland), `overseas` (outside Chinese mainland).
	PackageRegion *string `json:"PackageRegion,omitnil,omitempty" name:"PackageRegion"`

	// Resource pack type. Valid values: `CCU` (compute resource pack), `DISK` (storage resource pack).
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// Resource pack edition. Valid values: `base` (basic edition), `common` (general edition), `enterprise` (enterprise edition).
	PackageVersion *string `json:"PackageVersion,omitnil,omitempty" name:"PackageVersion"`

	// Resource pack size. Unit of the compute resource pack: Ten thousand.  Unit of the storage resource pack:  GB
	PackageSpec *float64 `json:"PackageSpec,omitnil,omitempty" name:"PackageSpec"`

	// Validity period of a resource pack in days
	ExpireDay *int64 `json:"ExpireDay,omitnil,omitempty" name:"ExpireDay"`

	// Number of the resource packs to be purchased
	PackageCount *int64 `json:"PackageCount,omitnil,omitempty" name:"PackageCount"`

	// Resource pack name
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`
}

func (r *CreateResourcePackageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateResourcePackageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceType")
	delete(f, "PackageRegion")
	delete(f, "PackageType")
	delete(f, "PackageVersion")
	delete(f, "PackageSpec")
	delete(f, "ExpireDay")
	delete(f, "PackageCount")
	delete(f, "PackageName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateResourcePackageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateResourcePackageResponseParams struct {
	// Big order number
	BigDealIds []*string `json:"BigDealIds,omitnil,omitempty" name:"BigDealIds"`

	// Each item has only one `dealName`, through which you need to ensure the idempotency of the delivery API.
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateResourcePackageResponse struct {
	*tchttp.BaseResponse
	Response *CreateResourcePackageResponseParams `json:"Response"`
}

func (r *CreateResourcePackageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateResourcePackageResponse) FromJsonString(s string) error {
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
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Update time
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// AZ
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Cluster name
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// Region
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Database version
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Number of instances
	InstanceNum *int64 `json:"InstanceNum,omitnil,omitempty" name:"InstanceNum"`

	// User UIN
	// Note: This field may return null, indicating that no valid values can be obtained.
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Engine type
	// Note: This field may return null, indicating that no valid values can be obtained.
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// User `appid`
	// Note: This field may return null, indicating that no valid values can be obtained.
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Cluster status description
	// Note: This field may return null, indicating that no valid values can be obtained.
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// Cluster creation time
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Billing mode. `0`: Pay-as-you-go; `1`: Monthly subscription.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// End time
	// Note: This field may return null, indicating that no valid values can be obtained.
	PeriodEndTime *string `json:"PeriodEndTime,omitnil,omitempty" name:"PeriodEndTime"`

	// Cluster read-write VIP
	// Note: This field may return null, indicating that no valid values can be obtained.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Cluster read-write vport
	// Note: This field may return null, indicating that no valid values can be obtained.
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// Project ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// VPC ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// TDSQL-C kernel version
	// Note: This field may return null, indicating that no valid values can be obtained.
	CynosVersion *string `json:"CynosVersion,omitnil,omitempty" name:"CynosVersion"`

	// Storage capacity
	// Note: This field may return null, indicating that no valid values can be obtained.
	StorageLimit *int64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`

	// Renewal flag
	// Note: This field may return null, indicating that no valid values can be obtained.
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// Task in progress
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProcessingTask *string `json:"ProcessingTask,omitnil,omitempty" name:"ProcessingTask"`

	// Array of tasks in the cluster
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tasks []*ObjectTask `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// Array of tags bound to the cluster
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResourceTags []*Tag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// Database type. Valid values: `NORMAL`, `SERVERLESS`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DbMode *string `json:"DbMode,omitnil,omitempty" name:"DbMode"`

	// Serverless cluster status when the database type is `SERVERLESS`. Valid values:
	// `resume`
	// `pause`
	// Note: This field may return null, indicating that no valid values can be obtained.
	ServerlessStatus *string `json:"ServerlessStatus,omitnil,omitempty" name:"ServerlessStatus"`

	// Prepaid cluster storage capacity
	// Note: This field may return null, indicating that no valid values can be obtained.
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// Cluster storage ID used in prepaid storage modification
	// Note: This field may return null, indicating that no valid values can be obtained.
	StorageId *string `json:"StorageId,omitnil,omitempty" name:"StorageId"`

	// Billing mode of cluster storage. Valid values: `0` (pay-as-you-go), `1` (monthly subscription).
	// Note: This field may return null, indicating that no valid values can be obtained.
	StoragePayMode *int64 `json:"StoragePayMode,omitnil,omitempty" name:"StoragePayMode"`

	// The minimum storage corresponding to the compute specification of the cluster
	// Note: This field may return null, indicating that no valid values can be obtained.
	MinStorageSize *int64 `json:"MinStorageSize,omitnil,omitempty" name:"MinStorageSize"`

	// The maximum storage corresponding to the compute specification of the cluster
	// Note: This field may return null, indicating that no valid values can be obtained.
	MaxStorageSize *int64 `json:"MaxStorageSize,omitnil,omitempty" name:"MaxStorageSize"`

	// Network information of the cluster
	// Note: This field may return null, indicating that no valid values can be obtained.
	NetAddrs []*NetAddr `json:"NetAddrs,omitnil,omitempty" name:"NetAddrs"`

	// Physical AZ
	// Note: This field may return null, indicating that no valid values can be obtained.
	PhysicalZone *string `json:"PhysicalZone,omitnil,omitempty" name:"PhysicalZone"`

	// Primary AZ
	// Note: This field may return null, indicating that no valid values can be obtained.
	MasterZone *string `json:"MasterZone,omitnil,omitempty" name:"MasterZone"`

	// Whether there is a secondary AZ
	// Note: This field may return null, indicating that no valid values can be obtained.
	HasSlaveZone *string `json:"HasSlaveZone,omitnil,omitempty" name:"HasSlaveZone"`

	// Secondary AZ
	// Note: This field may return null, indicating that no valid values can be obtained.
	SlaveZones []*string `json:"SlaveZones,omitnil,omitempty" name:"SlaveZones"`

	// Business type
	// Note: This field may return null, indicating that no valid values can be obtained.
	BusinessType *string `json:"BusinessType,omitnil,omitempty" name:"BusinessType"`

	// Whether to freeze
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsFreeze *string `json:"IsFreeze,omitnil,omitempty" name:"IsFreeze"`

	// Order source
	// Note: This field may return null, indicating that no valid values can be obtained.
	OrderSource *string `json:"OrderSource,omitnil,omitempty" name:"OrderSource"`

	// Capability
	// Note: This field may return null, indicating that no valid values can be obtained.
	Ability *Ability `json:"Ability,omitnil,omitempty" name:"Ability"`

	// Information of the resource pack bound to an instance when `packageType` is `DISK`. Note: This field may return null, indicating that no valid values can be obtained.
	ResourcePackages []*ResourcePackage `json:"ResourcePackages,omitnil,omitempty" name:"ResourcePackages"`
}

type CynosdbClusterDetail struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Cluster name
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// Region
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// AZ
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Physical AZ
	// Note: This field may return null, indicating that no valid values can be obtained.
	PhysicalZone *string `json:"PhysicalZone,omitnil,omitempty" name:"PhysicalZone"`

	// Status
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Status description
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// Serverless cluster status when the database type is `SERVERLESS`. Valid values:
	// resume
	// resuming
	// pause
	// pausing
	ServerlessStatus *string `json:"ServerlessStatus,omitnil,omitempty" name:"ServerlessStatus"`

	// Storage ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	StorageId *string `json:"StorageId,omitnil,omitempty" name:"StorageId"`

	// Storage capacity in GB
	// Note: This field may return null, indicating that no valid values can be obtained.
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// Maximum storage specification in GB
	// Note: This field may return null, indicating that no valid values can be obtained.
	MaxStorageSize *int64 `json:"MaxStorageSize,omitnil,omitempty" name:"MaxStorageSize"`

	// Minimum storage specification in GB
	// Note: This field may return null, indicating that no valid values can be obtained.
	MinStorageSize *int64 `json:"MinStorageSize,omitnil,omitempty" name:"MinStorageSize"`

	// Storage billing mode. Valid values: `1` (monthly subscription), `0` (pay-as-you-go).
	// Note: This field may return null, indicating that no valid values can be obtained.
	StoragePayMode *int64 `json:"StoragePayMode,omitnil,omitempty" name:"StoragePayMode"`

	// VPC name
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// Unique VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet name
	SubnetName *string `json:"SubnetName,omitnil,omitempty" name:"SubnetName"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Character set
	Charset *string `json:"Charset,omitnil,omitempty" name:"Charset"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Database type
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Database type. Valid values: `normal`, `serverless`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DbMode *string `json:"DbMode,omitnil,omitempty" name:"DbMode"`

	// Database version
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`

	// Maximum storage space
	// Note: This field may return null, indicating that no valid values can be obtained.
	StorageLimit *int64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`

	// Used capacity
	UsedStorage *int64 `json:"UsedStorage,omitnil,omitempty" name:"UsedStorage"`

	// VIP
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// vport
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// VIP and vport of the read-only instance in a cluster
	RoAddr []*Addr `json:"RoAddr,omitnil,omitempty" name:"RoAddr"`

	// Features supported by the cluster
	// Note: This field may return null, indicating that no valid values can be obtained.
	Ability *Ability `json:"Ability,omitnil,omitempty" name:"Ability"`

	// TDSQL-C version
	// Note: This field may return null, indicating that no valid values can be obtained.
	CynosVersion *string `json:"CynosVersion,omitnil,omitempty" name:"CynosVersion"`

	// Business type
	// Note: This field may return null, indicating that no valid values can be obtained.
	BusinessType *string `json:"BusinessType,omitnil,omitempty" name:"BusinessType"`

	// Whether there is a secondary AZ
	// Note: This field may return null, indicating that no valid values can be obtained.
	HasSlaveZone *string `json:"HasSlaveZone,omitnil,omitempty" name:"HasSlaveZone"`

	// Whether to freeze
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsFreeze *string `json:"IsFreeze,omitnil,omitempty" name:"IsFreeze"`

	// Task list
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tasks []*ObjectTask `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// Primary AZ
	// Note: This field may return null, indicating that no valid values can be obtained.
	MasterZone *string `json:"MasterZone,omitnil,omitempty" name:"MasterZone"`

	// Secondary AZ list
	// Note: This field may return null, indicating that no valid values can be obtained.
	SlaveZones []*string `json:"SlaveZones,omitnil,omitempty" name:"SlaveZones"`

	// Instance information
	InstanceSet []*ClusterInstanceDetail `json:"InstanceSet,omitnil,omitempty" name:"InstanceSet"`

	// Billing mode
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Expiration time
	PeriodEndTime *string `json:"PeriodEndTime,omitnil,omitempty" name:"PeriodEndTime"`

	// Project ID
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// Array of tags bound to instance
	ResourceTags []*Tag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// Proxy status
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProxyStatus *string `json:"ProxyStatus,omitnil,omitempty" name:"ProxyStatus"`

	// Binlog switch. Valid values: `ON`, `OFF`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LogBin *string `json:"LogBin,omitnil,omitempty" name:"LogBin"`

	// Whether to skip the transaction
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsSkipTrade *string `json:"IsSkipTrade,omitnil,omitempty" name:"IsSkipTrade"`

	// PITR type. Valid values: `normal`, `redo_pitr`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PitrType *string `json:"PitrType,omitnil,omitempty" name:"PitrType"`

	// Whether to enable password complexity
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsOpenPasswordComplexity *string `json:"IsOpenPasswordComplexity,omitnil,omitempty" name:"IsOpenPasswordComplexity"`

	// Network type
	// Note: This field may return null, indicating that no valid values can be obtained.
	NetworkStatus *string `json:"NetworkStatus,omitnil,omitempty" name:"NetworkStatus"`

	// Information of the resource pack bound to a cluster Note: This field may return null, indicating that no valid values can be obtained.
	ResourcePackages []*ResourcePackage `json:"ResourcePackages,omitnil,omitempty" name:"ResourcePackages"`

	// The auto-renewal flag. Valid values: `0`: (manual renewal, default), `1` (auto-renewal). Note: This field may return null, indicating that no valid values can be obtained.
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`


	NetworkType *string `json:"NetworkType,omitnil,omitempty" name:"NetworkType"`


	SlaveZoneAttr []*SlaveZoneAttrItem `json:"SlaveZoneAttr,omitnil,omitempty" name:"SlaveZoneAttr"`
}

type CynosdbErrorLogItem struct {
	// Log timestamp Note: This field may return null, indicating that no valid values can be obtained.
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// Log level Note: This field may return null, indicating that no valid values can be obtained.
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// Log content Note: This field may return null, indicating that no valid values can be obtained.
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type CynosdbInstance struct {
	// User `Uin`
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// User `AppId`
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Cluster name
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Region
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// AZ
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Instance status
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Instance status description
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// Instance type, which is used to indicate whether it is a serverless instance.
	DbMode *string `json:"DbMode,omitnil,omitempty" name:"DbMode"`

	// Database type
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Database version
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`

	// Number of CPU cores
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Memory in GB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Storage capacity in GB
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// Instance type
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Current instance role
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// Update time
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Private IP of instance
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Private port of instance
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// Billing mode
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Instance expiration time
	PeriodEndTime *string `json:"PeriodEndTime,omitnil,omitempty" name:"PeriodEndTime"`

	// Termination deadline
	DestroyDeadlineText *string `json:"DestroyDeadlineText,omitnil,omitempty" name:"DestroyDeadlineText"`

	// Isolation time
	IsolateTime *string `json:"IsolateTime,omitnil,omitempty" name:"IsolateTime"`

	// Network type
	NetType *int64 `json:"NetType,omitnil,omitempty" name:"NetType"`

	// Public domain name
	WanDomain *string `json:"WanDomain,omitnil,omitempty" name:"WanDomain"`

	// Public IP
	WanIP *string `json:"WanIP,omitnil,omitempty" name:"WanIP"`

	// Public port
	WanPort *int64 `json:"WanPort,omitnil,omitempty" name:"WanPort"`

	// Public network status
	WanStatus *string `json:"WanStatus,omitnil,omitempty" name:"WanStatus"`

	// Instance termination time
	DestroyTime *string `json:"DestroyTime,omitnil,omitempty" name:"DestroyTime"`

	// TDSQL-C kernel version
	CynosVersion *string `json:"CynosVersion,omitnil,omitempty" name:"CynosVersion"`

	// Task in progress
	ProcessingTask *string `json:"ProcessingTask,omitnil,omitempty" name:"ProcessingTask"`

	// Renewal flag
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// Minimum number of CPU cores for serverless instance
	MinCpu *float64 `json:"MinCpu,omitnil,omitempty" name:"MinCpu"`

	// Maximum number of CPU cores for serverless instance
	MaxCpu *float64 `json:"MaxCpu,omitnil,omitempty" name:"MaxCpu"`

	// Serverless instance status. Valid values:
	// resume
	// pause
	ServerlessStatus *string `json:"ServerlessStatus,omitnil,omitempty" name:"ServerlessStatus"`

	// Prepaid storage ID
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	StorageId *string `json:"StorageId,omitnil,omitempty" name:"StorageId"`

	// Storage billing mode
	StoragePayMode *int64 `json:"StoragePayMode,omitnil,omitempty" name:"StoragePayMode"`

	// Physical zone
	PhysicalZone *string `json:"PhysicalZone,omitnil,omitempty" name:"PhysicalZone"`

	// Business type
	// Note: This field may return null, indicating that no valid value can be obtained.
	BusinessType *string `json:"BusinessType,omitnil,omitempty" name:"BusinessType"`

	// Task
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tasks []*ObjectTask `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// Whether to freeze
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsFreeze *string `json:"IsFreeze,omitnil,omitempty" name:"IsFreeze"`

	// The resource tag
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResourceTags []*Tag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// Source AZ
	// Note: This field may return null, indicating that no valid value can be obtained.
	MasterZone *string `json:"MasterZone,omitnil,omitempty" name:"MasterZone"`

	// Replica AZ
	// Note: This field may return null, indicating that no valid value can be obtained.
	SlaveZones []*string `json:"SlaveZones,omitnil,omitempty" name:"SlaveZones"`

	// Instance network information
	// Note: This field may return null, indicating that no valid value can be obtained.
	InstanceNetInfo []*InstanceNetInfo `json:"InstanceNetInfo,omitnil,omitempty" name:"InstanceNetInfo"`

	// Information of the resource pack bound to an instance when `packageType` is `CCU`. Note: This field may return null, indicating that no valid values can be obtained.
	ResourcePackages []*ResourcePackage `json:"ResourcePackages,omitnil,omitempty" name:"ResourcePackages"`
}

type CynosdbInstanceDetail struct {
	// User `Uin`
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// User `AppId`
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Cluster name
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Region
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// AZ
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Instance status
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Instance status description
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// Database type
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Database version
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`

	// Number of CPU cores
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Memory in GB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Storage capacity in GB
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// Instance type
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Current instance role
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// Update time
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Billing mode
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Instance expiration time
	PeriodEndTime *string `json:"PeriodEndTime,omitnil,omitempty" name:"PeriodEndTime"`

	// Network type
	NetType *int64 `json:"NetType,omitnil,omitempty" name:"NetType"`

	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Private IP of instance
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Private port of instance
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// Public domain name of instance
	WanDomain *string `json:"WanDomain,omitnil,omitempty" name:"WanDomain"`

	// Character set
	Charset *string `json:"Charset,omitnil,omitempty" name:"Charset"`

	// TDSQL-C kernel version
	CynosVersion *string `json:"CynosVersion,omitnil,omitempty" name:"CynosVersion"`

	// Renewal flag
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// The minimum number of CPU cores for a serverless instance
	MinCpu *float64 `json:"MinCpu,omitnil,omitempty" name:"MinCpu"`

	// The maximum number of CPU cores for a serverless instance
	MaxCpu *float64 `json:"MaxCpu,omitnil,omitempty" name:"MaxCpu"`

	// Serverless instance status. Valid values:
	// resume
	// pause
	ServerlessStatus *string `json:"ServerlessStatus,omitnil,omitempty" name:"ServerlessStatus"`
}

type CynosdbInstanceGrp struct {
	// User `appId`
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Creation time
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// Deletion time
	DeletedTime *string `json:"DeletedTime,omitnil,omitempty" name:"DeletedTime"`

	// Instance group ID
	InstanceGrpId *string `json:"InstanceGrpId,omitnil,omitempty" name:"InstanceGrpId"`

	// Status
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Instance group type. ha: HA group; ro: RO group
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Update time
	UpdatedTime *string `json:"UpdatedTime,omitnil,omitempty" name:"UpdatedTime"`

	// Private IP
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Private port
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// Public domain name
	WanDomain *string `json:"WanDomain,omitnil,omitempty" name:"WanDomain"`

	// Public IP
	WanIP *string `json:"WanIP,omitnil,omitempty" name:"WanIP"`

	// Public port
	WanPort *int64 `json:"WanPort,omitnil,omitempty" name:"WanPort"`

	// Public network status
	WanStatus *string `json:"WanStatus,omitnil,omitempty" name:"WanStatus"`

	// Information of instances contained in instance group
	InstanceSet []*CynosdbInstance `json:"InstanceSet,omitnil,omitempty" name:"InstanceSet"`

	// VPC ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// Subnet ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// Information of the old IP
	// Note: This field may return null, indicating that no valid values can be obtained.
	OldAddrInfo *OldAddrInfo `json:"OldAddrInfo,omitnil,omitempty" name:"OldAddrInfo"`

	// Task in progress
	ProcessingTasks []*string `json:"ProcessingTasks,omitnil,omitempty" name:"ProcessingTasks"`

	// Task list
	Tasks []*ObjectTask `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// biz_net_service table ID
	NetServiceId *int64 `json:"NetServiceId,omitnil,omitempty" name:"NetServiceId"`
}

type DatabasePrivileges struct {
	// Database
	Db *string `json:"Db,omitnil,omitempty" name:"Db"`

	// Permission list
	Privileges []*string `json:"Privileges,omitnil,omitempty" name:"Privileges"`
}

type DatabaseTables struct {
	// Database name
	// Note: This field may return null, indicating that no valid values can be obtained.
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// Table name list
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tables []*string `json:"Tables,omitnil,omitempty" name:"Tables"`
}

type DbInfo struct {
	// Database name
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Character set
	CharacterSet *string `json:"CharacterSet,omitnil,omitempty" name:"CharacterSet"`

	// Database status
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Collation
	CollateRule *string `json:"CollateRule,omitnil,omitempty" name:"CollateRule"`

	// Database remarks Note: This field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// User permissions Note: This field may return null, indicating that no valid values can be obtained.
	UserHostPrivileges []*UserHostPrivilege `json:"UserHostPrivileges,omitnil,omitempty" name:"UserHostPrivileges"`

	// Database ID Note: This field may return null, indicating that no valid values can be obtained.
	DbId *int64 `json:"DbId,omitnil,omitempty" name:"DbId"`

	// Creation time Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Update time Note: This field may return null, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// User appid Note: This field may return null, indicating that no valid values can be obtained.
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// User Uin Note: This field may return null, indicating that no valid values can be obtained.
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Cluster ID Note: This field may return null, indicating that no valid values can be obtained.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

// Predefined struct for user
type DeleteAccountsRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Accounts in array, which contains `account` and `host`.
	Accounts []*InputAccount `json:"Accounts,omitnil,omitempty" name:"Accounts"`
}

type DeleteAccountsRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Accounts in array, which contains `account` and `host`.
	Accounts []*InputAccount `json:"Accounts,omitnil,omitempty" name:"Accounts"`
}

func (r *DeleteAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Accounts")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAccountsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAccountsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAccountsResponseParams `json:"Response"`
}

func (r *DeleteAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAuditRuleTemplatesRequestParams struct {
	// Audit rule template ID
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`
}

type DeleteAuditRuleTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// Audit rule template ID
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`
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
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Backup file ID. This field is used by legacy versions and thus not recommended.
	SnapshotIdList []*int64 `json:"SnapshotIdList,omitnil,omitempty" name:"SnapshotIdList"`

	// Backup file ID. This field is recommended.
	BackupIds []*int64 `json:"BackupIds,omitnil,omitempty" name:"BackupIds"`
}

type DeleteBackupRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Backup file ID. This field is used by legacy versions and thus not recommended.
	SnapshotIdList []*int64 `json:"SnapshotIdList,omitnil,omitempty" name:"SnapshotIdList"`

	// Backup file ID. This field is recommended.
	BackupIds []*int64 `json:"BackupIds,omitnil,omitempty" name:"BackupIds"`
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
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteClusterDatabaseRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`


	DbNames []*string `json:"DbNames,omitnil,omitempty" name:"DbNames"`
}

type DeleteClusterDatabaseRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	DbNames []*string `json:"DbNames,omitnil,omitempty" name:"DbNames"`
}

func (r *DeleteClusterDatabaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterDatabaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "DbNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteClusterDatabaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterDatabaseResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteClusterDatabaseResponse struct {
	*tchttp.BaseResponse
	Response *DeleteClusterDatabaseResponseParams `json:"Response"`
}

func (r *DeleteClusterDatabaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterDatabaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteParamTemplateRequestParams struct {
	// Parameter template ID
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DeleteParamTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Parameter template ID
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *DeleteParamTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteParamTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteParamTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteParamTemplateResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteParamTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteParamTemplateResponseParams `json:"Response"`
}

func (r *DeleteParamTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteParamTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountPrivilegesRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Account name
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// Host
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// When the database name is “*”, the value specified in `Type` and `TableName` will be ignored, indicating that the user's global permissions are being modified.
	Db *string `json:"Db,omitnil,omitempty" name:"Db"`

	// Object type in a specified database. Valid values: `table`, `*`.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// The database name can be specified when `Type` is 'table'.
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`
}

type DescribeAccountPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Account name
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// Host
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// When the database name is “*”, the value specified in `Type` and `TableName` will be ignored, indicating that the user's global permissions are being modified.
	Db *string `json:"Db,omitnil,omitempty" name:"Db"`

	// Object type in a specified database. Valid values: `table`, `*`.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// The database name can be specified when `Type` is 'table'.
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`
}

func (r *DescribeAccountPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountPrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "AccountName")
	delete(f, "Host")
	delete(f, "Db")
	delete(f, "Type")
	delete(f, "TableName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountPrivilegesResponseParams struct {
	// The list of permissions, such as  ["select","update","delete","create","drop","references","index","alter","show_db","create_tmp_table","lock_tables","execute","create_view","show_view","create_routine","alter_routine","event","trigger"]
	Privileges []*string `json:"Privileges,omitnil,omitempty" name:"Privileges"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccountPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccountPrivilegesResponseParams `json:"Response"`
}

func (r *DescribeAccountPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountsRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// List of accounts to be filtered
	AccountNames []*string `json:"AccountNames,omitnil,omitempty" name:"AccountNames"`

	// Database type. Valid values: 
	// <li> MYSQL </li>
	// This parameter has been disused.
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// List of accounts to be filtered
	Hosts []*string `json:"Hosts,omitnil,omitempty" name:"Hosts"`

	// Maximum entries returned per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Keywords for fuzzy search (match `AccountName` and `AccountHost` at the same time), which supports regex. The union results will be returned.
	AccountRegular *string `json:"AccountRegular,omitnil,omitempty" name:"AccountRegular"`
}

type DescribeAccountsRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// List of accounts to be filtered
	AccountNames []*string `json:"AccountNames,omitnil,omitempty" name:"AccountNames"`

	// Database type. Valid values: 
	// <li> MYSQL </li>
	// This parameter has been disused.
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// List of accounts to be filtered
	Hosts []*string `json:"Hosts,omitnil,omitempty" name:"Hosts"`

	// Maximum entries returned per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Keywords for fuzzy search (match `AccountName` and `AccountHost` at the same time), which supports regex. The union results will be returned.
	AccountRegular *string `json:"AccountRegular,omitnil,omitempty" name:"AccountRegular"`
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
	delete(f, "AccountRegular")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountsResponseParams struct {
	// Database account list
	// Note: This field may return null, indicating that no valid values can be obtained.
	AccountSet []*Account `json:"AccountSet,omitnil,omitempty" name:"AccountSet"`

	// Total number of accounts
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`

	// Rule template name
	RuleTemplateNames []*string `json:"RuleTemplateNames,omitnil,omitempty" name:"RuleTemplateNames"`

	// Number of results returned per request. Default value: `20`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Default value: `0`.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeAuditRuleTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// Rule template ID
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`

	// Rule template name
	RuleTemplateNames []*string `json:"RuleTemplateNames,omitnil,omitempty" name:"RuleTemplateNames"`

	// Number of results returned per request. Default value: `20`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Default value: `0`.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of rule template details
	// Note: This field may return null, indicating that no valid values can be obtained.
	Items []*AuditRuleTemplateInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type DescribeAuditRuleWithInstanceIdsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. Currently, only one single instance can be queried.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Audit rule information of the instance
	// Note: This field may return null, indicating that no valid values can be obtained.
	Items []*InstanceAuditRule `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeBackupConfigRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
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
	BackupTimeBeg *uint64 `json:"BackupTimeBeg,omitnil,omitempty" name:"BackupTimeBeg"`

	// Full backup end time. Value range: [0-24*3600]. For example, 0:00 AM, 1:00 AM, and 2:00 AM are represented by 0, 3600, and 7200, respectively
	BackupTimeEnd *uint64 `json:"BackupTimeEnd,omitnil,omitempty" name:"BackupTimeEnd"`

	// Backup retention period in seconds. Backups will be cleared after this period elapses. 7 days is represented by 3600*24*7 = 604800
	ReserveDuration *uint64 `json:"ReserveDuration,omitnil,omitempty" name:"ReserveDuration"`

	// Backup frequency. It is an array of 7 elements corresponding to Monday through Sunday. full: full backup; increment: incremental backup
	// Note: this field may return null, indicating that no valid values can be obtained.
	BackupFreq []*string `json:"BackupFreq,omitnil,omitempty" name:"BackupFreq"`

	// Backup mode. logic: logic backup; snapshot: snapshot backup
	// Note: this field may return null, indicating that no valid values can be obtained.
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Backup ID
	BackupId *int64 `json:"BackupId,omitnil,omitempty" name:"BackupId"`
}

type DescribeBackupDownloadUrlRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Backup ID
	BackupId *int64 `json:"BackupId,omitnil,omitempty" name:"BackupId"`
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
	DownloadUrl *string `json:"DownloadUrl,omitnil,omitempty" name:"DownloadUrl"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// The number of results to be returned. Value range: (0,100]
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Record offset. Value range: [0,INF)
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Database type. Valid values: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Backup ID
	BackupIds []*int64 `json:"BackupIds,omitnil,omitempty" name:"BackupIds"`

	// Backup type. Valid values: `snapshot` (snapshot backup), `logic` (logic backup).
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// Back mode. Valid values: `auto` (automatic backup), `manual` (manual backup)
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`


	SnapShotType *string `json:"SnapShotType,omitnil,omitempty" name:"SnapShotType"`

	// Backup start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Backup end time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`


	FileNames []*string `json:"FileNames,omitnil,omitempty" name:"FileNames"`

	// Backup alias, which supports fuzzy query.
	BackupNames []*string `json:"BackupNames,omitnil,omitempty" name:"BackupNames"`

	// ID list of the snapshot backup
	SnapshotIdList []*int64 `json:"SnapshotIdList,omitnil,omitempty" name:"SnapshotIdList"`
}

type DescribeBackupListRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// The number of results to be returned. Value range: (0,100]
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Record offset. Value range: [0,INF)
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Database type. Valid values: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Backup ID
	BackupIds []*int64 `json:"BackupIds,omitnil,omitempty" name:"BackupIds"`

	// Backup type. Valid values: `snapshot` (snapshot backup), `logic` (logic backup).
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// Back mode. Valid values: `auto` (automatic backup), `manual` (manual backup)
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	SnapShotType *string `json:"SnapShotType,omitnil,omitempty" name:"SnapShotType"`

	// Backup start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Backup end time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	FileNames []*string `json:"FileNames,omitnil,omitempty" name:"FileNames"`

	// Backup alias, which supports fuzzy query.
	BackupNames []*string `json:"BackupNames,omitnil,omitempty" name:"BackupNames"`

	// ID list of the snapshot backup
	SnapshotIdList []*int64 `json:"SnapshotIdList,omitnil,omitempty" name:"SnapshotIdList"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Backup file list
	BackupList []*BackupFileInfo `json:"BackupList,omitnil,omitempty" name:"BackupList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Binlog file ID
	BinlogId *int64 `json:"BinlogId,omitnil,omitempty" name:"BinlogId"`
}

type DescribeBinlogDownloadUrlRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Binlog file ID
	BinlogId *int64 `json:"BinlogId,omitnil,omitempty" name:"BinlogId"`
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
	DownloadUrl *string `json:"DownloadUrl,omitnil,omitempty" name:"DownloadUrl"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeBinlogSaveDaysRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
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
	BinlogSaveDays *int64 `json:"BinlogSaveDays,omitnil,omitempty" name:"BinlogSaveDays"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeBinlogsRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Binlog list
	// Note: This field may return null, indicating that no valid values can be obtained.
	Binlogs []*BinlogItem `json:"Binlogs,omitnil,omitempty" name:"Binlogs"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeClusterDetailDatabasesRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Offset. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: `20`. Maximum value: `100`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Database name
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`
}

type DescribeClusterDetailDatabasesRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Offset. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: `20`. Maximum value: `100`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Database name
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`
}

func (r *DescribeClusterDetailDatabasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterDetailDatabasesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "DbName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterDetailDatabasesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterDetailDatabasesResponseParams struct {
	// Database information Note: This field may return null, indicating that no valid values can be obtained.
	DbInfos []*DbInfo `json:"DbInfos,omitnil,omitempty" name:"DbInfos"`

	// The total count
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterDetailDatabasesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterDetailDatabasesResponseParams `json:"Response"`
}

func (r *DescribeClusterDetailDatabasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterDetailDatabasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterDetailRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeClusterDetailRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
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
	Detail *CynosdbClusterDetail `json:"Detail,omitnil,omitempty" name:"Detail"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeClusterInstanceGrpsRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Instance group list
	InstanceGrpInfoList []*CynosdbInstanceGrp `json:"InstanceGrpInfoList,omitnil,omitempty" name:"InstanceGrpInfoList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Parameter name
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`
}

type DescribeClusterParamsRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Parameter name
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Instance parameter list
	// Note: This field may return null, indicating that no valid values can be obtained.
	Items []*ParamInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeClusterPasswordComplexityRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeClusterPasswordComplexityRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterPasswordComplexityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterPasswordComplexityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterPasswordComplexityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterPasswordComplexityResponseParams struct {
	// Data dictionary parameter Note: This field may return null, indicating that no valid values can be obtained.
	ValidatePasswordDictionary *ParamInfo `json:"ValidatePasswordDictionary,omitnil,omitempty" name:"ValidatePasswordDictionary"`

	// The length of the password Note: This field may return null, indicating that no valid values can be obtained.
	ValidatePasswordLength *ParamInfo `json:"ValidatePasswordLength,omitnil,omitempty" name:"ValidatePasswordLength"`

	// Number of case-sensitive characters Note: This field may return null, indicating that no valid values can be obtained.
	ValidatePasswordMixedCaseCount *ParamInfo `json:"ValidatePasswordMixedCaseCount,omitnil,omitempty" name:"ValidatePasswordMixedCaseCount"`

	// Number of digits Note: This field may return null, indicating that no valid values can be obtained.
	ValidatePasswordNumberCount *ParamInfo `json:"ValidatePasswordNumberCount,omitnil,omitempty" name:"ValidatePasswordNumberCount"`

	// Password level Note: This field may return null, indicating that no valid values can be obtained.
	ValidatePasswordPolicy *ParamInfo `json:"ValidatePasswordPolicy,omitnil,omitempty" name:"ValidatePasswordPolicy"`

	// Number of symbols Note: This field may return null, indicating that no valid values can be obtained.
	ValidatePasswordSpecialCharCount *ParamInfo `json:"ValidatePasswordSpecialCharCount,omitnil,omitempty" name:"ValidatePasswordSpecialCharCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterPasswordComplexityResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterPasswordComplexityResponseParams `json:"Response"`
}

func (r *DescribeClusterPasswordComplexityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterPasswordComplexityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClustersRequestParams struct {
	// Engine type. Currently, `MYSQL` is supported.
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Number of returned results. Default value: 20. Maximum value: 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Record offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sort by field. Valid values:
	// <li> CREATETIME: creation time</li>
	// <li> PERIODENDTIME: expiration time</li>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values:
	// <li> ASC: ascending</li>
	// <li> DESC: descending</li>
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Filter. If more than one filter exists, the logical relationship between these filters is `AND`.
	Filters []*QueryFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeClustersRequest struct {
	*tchttp.BaseRequest
	
	// Engine type. Currently, `MYSQL` is supported.
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Number of returned results. Default value: 20. Maximum value: 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Record offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sort by field. Valid values:
	// <li> CREATETIME: creation time</li>
	// <li> PERIODENDTIME: expiration time</li>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values:
	// <li> ASC: ascending</li>
	// <li> DESC: descending</li>
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Filter. If more than one filter exists, the logical relationship between these filters is `AND`.
	Filters []*QueryFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Cluster list
	ClusterSet []*CynosdbCluster `json:"ClusterSet,omitnil,omitempty" name:"ClusterSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDBSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Instance group ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	Groups []*SecurityGroup `json:"Groups,omitnil,omitempty" name:"Groups"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`
}

type DescribeFlowRequest struct {
	*tchttp.BaseRequest
	
	// Task flow ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`
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
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceDetailRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	Detail *CynosdbInstanceDetail `json:"Detail,omitnil,omitempty" name:"Detail"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeInstanceErrorLogsRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Limit on the number of logs
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset of the log number
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Sorting field. Valid value: 'Timestamp'.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: `ASC`, `DESC`.
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Log level, which supports combo search by multiple levels. Valid values: `error`, `warning`, `note`.
	LogLevels []*string `json:"LogLevels,omitnil,omitempty" name:"LogLevels"`


	KeyWords []*string `json:"KeyWords,omitnil,omitempty" name:"KeyWords"`
}

type DescribeInstanceErrorLogsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Limit on the number of logs
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset of the log number
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Sorting field. Valid value: 'Timestamp'.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: `ASC`, `DESC`.
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Log level, which supports combo search by multiple levels. Valid values: `error`, `warning`, `note`.
	LogLevels []*string `json:"LogLevels,omitnil,omitempty" name:"LogLevels"`

	KeyWords []*string `json:"KeyWords,omitnil,omitempty" name:"KeyWords"`
}

func (r *DescribeInstanceErrorLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceErrorLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "LogLevels")
	delete(f, "KeyWords")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceErrorLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceErrorLogsResponseParams struct {
	// Number of logs Note: This field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Error log list Note: This field may return null, indicating that no valid values can be obtained.
	ErrorLogs []*CynosdbErrorLogItem `json:"ErrorLogs,omitnil,omitempty" name:"ErrorLogs"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceErrorLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceErrorLogsResponseParams `json:"Response"`
}

func (r *DescribeInstanceErrorLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceErrorLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceParamsRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Instance ID, which supports batch query.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Search condition for a parameter name, which supports fuzzy search.
	ParamKeyword *string `json:"ParamKeyword,omitnil,omitempty" name:"ParamKeyword"`
}

type DescribeInstanceParamsRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Instance ID, which supports batch query.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Search condition for a parameter name, which supports fuzzy search.
	ParamKeyword *string `json:"ParamKeyword,omitnil,omitempty" name:"ParamKeyword"`
}

func (r *DescribeInstanceParamsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceParamsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceIds")
	delete(f, "ParamKeyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceParamsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceParamsResponseParams struct {
	// List of instance parameters
	Items []*InstanceParamItem `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceParamsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceParamsResponseParams `json:"Response"`
}

func (r *DescribeInstanceParamsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceParamsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceSlowQueriesRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Transaction start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Transaction end time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Maximum number
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Username
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// Client host
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// Database name
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// Sorting field. Valid values: QueryTime, LockTime, RowsExamined, RowsSent.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: asc, desc.
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

type DescribeInstanceSlowQueriesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Transaction start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Transaction end time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Maximum number
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Username
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// Client host
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// Database name
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// Sorting field. Valid values: QueryTime, LockTime, RowsExamined, RowsSent.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: asc, desc.
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Slow query record
	SlowQueries []*SlowQueriesItem `json:"SlowQueries,omitnil,omitempty" name:"SlowQueries"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Whether to return the AZ information.
	IncludeZoneStocks *bool `json:"IncludeZoneStocks,omitnil,omitempty" name:"IncludeZoneStocks"`
}

type DescribeInstanceSpecsRequest struct {
	*tchttp.BaseRequest
	
	// Database type. Valid values: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Whether to return the AZ information.
	IncludeZoneStocks *bool `json:"IncludeZoneStocks,omitnil,omitempty" name:"IncludeZoneStocks"`
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
	InstanceSpecSet []*InstanceSpec `json:"InstanceSpecSet,omitnil,omitempty" name:"InstanceSpecSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Record offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sort by field. Valid values:
	// <li> CREATETIME: creation time</li>
	// <li> PERIODENDTIME: expiration time</li>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values:
	// <li> ASC: ascending</li>
	// <li> DESC: descending</li>
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Filter. If more than one filter exists, the logical relationship between these filters is `AND`.
	Filters []*QueryFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Engine type. Currently, `MYSQL` is supported.
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Instance status. Valid values:
	// creating
	// running
	// isolating
	// isolated
	// activating: Removing the instance from isolation
	// offlining: Eliminating the instance
	// offlined: Instance eliminated
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Instance ID list
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Number of returned results. Default value: 20. Maximum value: 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Record offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sort by field. Valid values:
	// <li> CREATETIME: creation time</li>
	// <li> PERIODENDTIME: expiration time</li>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values:
	// <li> ASC: ascending</li>
	// <li> DESC: descending</li>
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Filter. If more than one filter exists, the logical relationship between these filters is `AND`.
	Filters []*QueryFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Engine type. Currently, `MYSQL` is supported.
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Instance status. Valid values:
	// creating
	// running
	// isolating
	// isolated
	// activating: Removing the instance from isolation
	// offlining: Eliminating the instance
	// offlined: Instance eliminated
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Instance ID list
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Instance list
	InstanceSet []*CynosdbInstance `json:"InstanceSet,omitnil,omitempty" name:"InstanceSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
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
type DescribeMaintainPeriodRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeMaintainPeriodRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	MaintainWeekDays []*string `json:"MaintainWeekDays,omitnil,omitempty" name:"MaintainWeekDays"`

	// Maintenance start time in seconds
	MaintainStartTime *int64 `json:"MaintainStartTime,omitnil,omitempty" name:"MaintainStartTime"`

	// Maintenance duration in seconds
	MaintainDuration *int64 `json:"MaintainDuration,omitnil,omitempty" name:"MaintainDuration"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeParamTemplateDetailRequestParams struct {
	// Parameter template ID
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DescribeParamTemplateDetailRequest struct {
	*tchttp.BaseRequest
	
	// Parameter template ID
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *DescribeParamTemplateDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParamTemplateDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeParamTemplateDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeParamTemplateDetailResponseParams struct {
	// Parameter template ID
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Parameter template name
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// Parameter template description
	TemplateDescription *string `json:"TemplateDescription,omitnil,omitempty" name:"TemplateDescription"`

	// Engine version
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// Total number of parameters
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of parameters
	Items []*ParamDetail `json:"Items,omitnil,omitempty" name:"Items"`

	// Database type. Valid values:  `NORMAL`, `SERVERLESS`.
	DbMode *string `json:"DbMode,omitnil,omitempty" name:"DbMode"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeParamTemplateDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeParamTemplateDetailResponseParams `json:"Response"`
}

func (r *DescribeParamTemplateDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParamTemplateDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeParamTemplatesRequestParams struct {
	// Database engine version number
	EngineVersions []*string `json:"EngineVersions,omitnil,omitempty" name:"EngineVersions"`

	// Template name
	TemplateNames []*string `json:"TemplateNames,omitnil,omitempty" name:"TemplateNames"`

	// Template ID
	TemplateIds []*int64 `json:"TemplateIds,omitnil,omitempty" name:"TemplateIds"`

	// Database Type. Valid values: `NORMAL`, `SERVERLESS`.
	DbModes []*string `json:"DbModes,omitnil,omitempty" name:"DbModes"`

	// Offset for query
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit on queries
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Product type of the queried template
	Products []*string `json:"Products,omitnil,omitempty" name:"Products"`

	// Template type
	TemplateTypes []*string `json:"TemplateTypes,omitnil,omitempty" name:"TemplateTypes"`

	// Version type
	EngineTypes []*string `json:"EngineTypes,omitnil,omitempty" name:"EngineTypes"`

	// The sorting order of the returned results
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: `desc`, `asc `.
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`
}

type DescribeParamTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// Database engine version number
	EngineVersions []*string `json:"EngineVersions,omitnil,omitempty" name:"EngineVersions"`

	// Template name
	TemplateNames []*string `json:"TemplateNames,omitnil,omitempty" name:"TemplateNames"`

	// Template ID
	TemplateIds []*int64 `json:"TemplateIds,omitnil,omitempty" name:"TemplateIds"`

	// Database Type. Valid values: `NORMAL`, `SERVERLESS`.
	DbModes []*string `json:"DbModes,omitnil,omitempty" name:"DbModes"`

	// Offset for query
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit on queries
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Product type of the queried template
	Products []*string `json:"Products,omitnil,omitempty" name:"Products"`

	// Template type
	TemplateTypes []*string `json:"TemplateTypes,omitnil,omitempty" name:"TemplateTypes"`

	// Version type
	EngineTypes []*string `json:"EngineTypes,omitnil,omitempty" name:"EngineTypes"`

	// The sorting order of the returned results
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: `desc`, `asc `.
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Parameter template information
	Items []*ParamTemplateListInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Maximum entries returned per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Search by keyword
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`
}

type DescribeProjectSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Project ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Maximum entries returned per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Search by keyword
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`
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
	Groups []*SecurityGroup `json:"Groups,omitnil,omitempty" name:"Groups"`

	// The total number of groups
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeProxiesRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Number of returned results. Default value: `20`. Maximum value: `100`,
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Record offset. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting field. Valid values:
	// <li> CREATETIME: Creation time</li>
	// <li> PERIODENDTIME: Expiration time</li>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values:
	// <li> `ASC`: Ascending.</li>
	// <li> `DESC`: Descending</li>
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Filter. If there are more than one filter, the logical relationship between these filters is `AND`.
	Filters []*QueryParamFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeProxiesRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Number of returned results. Default value: `20`. Maximum value: `100`,
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Record offset. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting field. Valid values:
	// <li> CREATETIME: Creation time</li>
	// <li> PERIODENDTIME: Expiration time</li>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values:
	// <li> `ASC`: Ascending.</li>
	// <li> `DESC`: Descending</li>
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Filter. If there are more than one filter, the logical relationship between these filters is `AND`.
	Filters []*QueryParamFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeProxiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxiesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProxiesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProxiesResponseParams struct {
	// Number of database proxy groups
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of database proxy groups
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProxyGroupInfos []*ProxyGroupInfo `json:"ProxyGroupInfos,omitnil,omitempty" name:"ProxyGroupInfos"`

	// Database proxy node
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProxyNodeInfos []*ProxyNodeInfo `json:"ProxyNodeInfos,omitnil,omitempty" name:"ProxyNodeInfos"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProxiesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProxiesResponseParams `json:"Response"`
}

func (r *DescribeProxiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProxyNodesRequestParams struct {
	// Number of returned results. Default value: `20`. Maximum value: `100`,
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Record offset. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting field. Valid values:
	// <li> CREATETIME: Creation time</li>
	// <li> PERIODENDTIME: Expiration time</li>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values:
	// <li> `ASC`: Ascending.</li>
	// <li> `DESC`: Descending</li>
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Filter. If there are more than one filter, the logical relationship between these filters is `AND`.
	Filters []*QueryFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeProxyNodesRequest struct {
	*tchttp.BaseRequest
	
	// Number of returned results. Default value: `20`. Maximum value: `100`,
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Record offset. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting field. Valid values:
	// <li> CREATETIME: Creation time</li>
	// <li> PERIODENDTIME: Expiration time</li>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values:
	// <li> `ASC`: Ascending.</li>
	// <li> `DESC`: Descending</li>
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Filter. If there are more than one filter, the logical relationship between these filters is `AND`.
	Filters []*QueryFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeProxyNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxyNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProxyNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProxyNodesResponseParams struct {
	// Number of the database proxy nodes
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of the database proxy nodes
	ProxyNodeInfos []*ProxyNodeInfo `json:"ProxyNodeInfos,omitnil,omitempty" name:"ProxyNodeInfos"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProxyNodesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProxyNodesResponseParams `json:"Response"`
}

func (r *DescribeProxyNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxyNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProxySpecsRequestParams struct {

}

type DescribeProxySpecsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeProxySpecsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxySpecsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProxySpecsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProxySpecsResponseParams struct {
	// List of database proxyspecifications
	ProxySpecs []*ProxySpec `json:"ProxySpecs,omitnil,omitempty" name:"ProxySpecs"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProxySpecsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProxySpecsResponseParams `json:"Response"`
}

func (r *DescribeProxySpecsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxySpecsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourcePackageDetailRequestParams struct {
	// The unique ID of a resource pack
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// Cluster ID
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// Start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Offset
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Instance D
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type DescribeResourcePackageDetailRequest struct {
	*tchttp.BaseRequest
	
	// The unique ID of a resource pack
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// Cluster ID
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// Start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Offset
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Instance D
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *DescribeResourcePackageDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcePackageDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PackageId")
	delete(f, "ClusterIds")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourcePackageDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourcePackageDetailResponseParams struct {
	// Total number of deducted resource packs
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Resource pack details
	Detail []*PackageDetail `json:"Detail,omitnil,omitempty" name:"Detail"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeResourcePackageDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourcePackageDetailResponseParams `json:"Response"`
}

func (r *DescribeResourcePackageDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcePackageDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourcePackageListRequestParams struct {
	// The unique ID of a resource pack
	PackageId []*string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// Resource pack name
	PackageName []*string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// Resource pack type. Valid values: `CCU` (compute resource pack), `DISK` (storage resource pack).
	PackageType []*string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// Region of the resource pack. Valid values: `China` (Chinese mainland), `overseas` (outside Chinese mainland).
	PackageRegion []*string `json:"PackageRegion,omitnil,omitempty" name:"PackageRegion"`

	// Resource pack status. Valid values: `using`, `expired`, `normal_finish` (used up), `apply_refund` (requesting a refund), `refund` (refunded).
	Status []*string `json:"Status,omitnil,omitempty" name:"Status"`

	// Sorting conditions. Valid values: `startTime` (effective time), `expireTime` (expiration date), `packageUsedSpec` (used capacity), `packageTotalSpec` (total storage capacity). 
	// Sorting by array order
	OrderBy []*string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: `ASC` (ascending), `DESC` (descending).
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeResourcePackageListRequest struct {
	*tchttp.BaseRequest
	
	// The unique ID of a resource pack
	PackageId []*string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// Resource pack name
	PackageName []*string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// Resource pack type. Valid values: `CCU` (compute resource pack), `DISK` (storage resource pack).
	PackageType []*string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// Region of the resource pack. Valid values: `China` (Chinese mainland), `overseas` (outside Chinese mainland).
	PackageRegion []*string `json:"PackageRegion,omitnil,omitempty" name:"PackageRegion"`

	// Resource pack status. Valid values: `using`, `expired`, `normal_finish` (used up), `apply_refund` (requesting a refund), `refund` (refunded).
	Status []*string `json:"Status,omitnil,omitempty" name:"Status"`

	// Sorting conditions. Valid values: `startTime` (effective time), `expireTime` (expiration date), `packageUsedSpec` (used capacity), `packageTotalSpec` (total storage capacity). 
	// Sorting by array order
	OrderBy []*string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: `ASC` (ascending), `DESC` (descending).
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeResourcePackageListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcePackageListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PackageId")
	delete(f, "PackageName")
	delete(f, "PackageType")
	delete(f, "PackageRegion")
	delete(f, "Status")
	delete(f, "OrderBy")
	delete(f, "OrderDirection")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourcePackageListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourcePackageListResponseParams struct {
	// Total number of resource packs
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Resource pack details Note: This field may return null, indicating that no valid values can be obtained.
	Detail []*Package `json:"Detail,omitnil,omitempty" name:"Detail"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeResourcePackageListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourcePackageListResponseParams `json:"Response"`
}

func (r *DescribeResourcePackageListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcePackageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourcePackageSaleSpecRequestParams struct {
	// Instance type
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Region of the resource pack. Valid values: `China` (Chinese mainland), `overseas` (outside Chinese mainland).
	PackageRegion *string `json:"PackageRegion,omitnil,omitempty" name:"PackageRegion"`

	// Resource pack type. Valid values: `CCU` (compute resource pack, `DISK` (storage resource pack).
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeResourcePackageSaleSpecRequest struct {
	*tchttp.BaseRequest
	
	// Instance type
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Region of the resource pack. Valid values: `China` (Chinese mainland), `overseas` (outside Chinese mainland).
	PackageRegion *string `json:"PackageRegion,omitnil,omitempty" name:"PackageRegion"`

	// Resource pack type. Valid values: `CCU` (compute resource pack, `DISK` (storage resource pack).
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeResourcePackageSaleSpecRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcePackageSaleSpecRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceType")
	delete(f, "PackageRegion")
	delete(f, "PackageType")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourcePackageSaleSpecRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourcePackageSaleSpecResponseParams struct {
	// Total number of available resource packs
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Resource pack details Note: This field may return null, indicating that no valid values can be obtained.
	Detail []*SalePackageSpec `json:"Detail,omitnil,omitempty" name:"Detail"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeResourcePackageSaleSpecResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourcePackageSaleSpecResponseParams `json:"Response"`
}

func (r *DescribeResourcePackageSaleSpecResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcePackageSaleSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourcesByDealNameRequestParams struct {
	// Order ID. (If the cluster is not delivered yet, the `DescribeResourcesByDealName` API may return the `InvalidParameterValue.DealNameNotFound` error. Call the API again until it succeeds.)
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// Order ID, which can be used to query the resource information of multiple orders ID (If the cluster is not delivered yet, the `DescribeResourcesByDealName` API may return the `InvalidParameterValue.DealNameNotFound` error. Call the API again until it succeeds.)
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`
}

type DescribeResourcesByDealNameRequest struct {
	*tchttp.BaseRequest
	
	// Order ID. (If the cluster is not delivered yet, the `DescribeResourcesByDealName` API may return the `InvalidParameterValue.DealNameNotFound` error. Call the API again until it succeeds.)
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// Order ID, which can be used to query the resource information of multiple orders ID (If the cluster is not delivered yet, the `DescribeResourcesByDealName` API may return the `InvalidParameterValue.DealNameNotFound` error. Call the API again until it succeeds.)
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`
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
	BillingResourceInfos []*BillingResourceInfo `json:"BillingResourceInfos,omitnil,omitempty" name:"BillingResourceInfos"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeRollbackTimeRangeRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
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
	TimeRangeStart *string `json:"TimeRangeStart,omitnil,omitempty" name:"TimeRangeStart"`

	// End time of valid rollback time range (disused)
	// Note: This field may return null, indicating that no valid values can be obtained.
	TimeRangeEnd *string `json:"TimeRangeEnd,omitnil,omitempty" name:"TimeRangeEnd"`

	// Time range available for rollback
	RollbackTimeRanges []*RollbackTimeRange `json:"RollbackTimeRanges,omitnil,omitempty" name:"RollbackTimeRanges"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Expected time point to roll back to
	ExpectTime *string `json:"ExpectTime,omitnil,omitempty" name:"ExpectTime"`

	// Error tolerance range for rollback time point
	ExpectTimeThresh *uint64 `json:"ExpectTimeThresh,omitnil,omitempty" name:"ExpectTimeThresh"`
}

type DescribeRollbackTimeValidityRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Expected time point to roll back to
	ExpectTime *string `json:"ExpectTime,omitnil,omitempty" name:"ExpectTime"`

	// Error tolerance range for rollback time point
	ExpectTimeThresh *uint64 `json:"ExpectTimeThresh,omitnil,omitempty" name:"ExpectTimeThresh"`
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
	PoolId *uint64 `json:"PoolId,omitnil,omitempty" name:"PoolId"`

	// Rollback task ID, which needs to be passed in when rolling back to this time point
	QueryId *uint64 `json:"QueryId,omitnil,omitempty" name:"QueryId"`

	// Whether the time point is valid. pass: check passed; fail: check failed
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Suggested time point. This value takes effect only if `Status` is `fail`
	SuggestTime *string `json:"SuggestTime,omitnil,omitempty" name:"SuggestTime"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeSupportProxyVersionRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`
}

type DescribeSupportProxyVersionRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`
}

func (r *DescribeSupportProxyVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSupportProxyVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ProxyGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSupportProxyVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSupportProxyVersionResponseParams struct {
	// Collection of supported database proxy versions
	// Note: This field may return null, indicating that no valid values can be obtained.
	SupportProxyVersions []*string `json:"SupportProxyVersions,omitnil,omitempty" name:"SupportProxyVersions"`

	// The current proxy version
	// Note: This field may return null, indicating that no valid values can be obtained.
	CurrentProxyVersion *string `json:"CurrentProxyVersion,omitnil,omitempty" name:"CurrentProxyVersion"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSupportProxyVersionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSupportProxyVersionResponseParams `json:"Response"`
}

func (r *DescribeSupportProxyVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSupportProxyVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZonesRequestParams struct {
	// Whether the virtual zone is included.–
	IncludeVirtualZones *bool `json:"IncludeVirtualZones,omitnil,omitempty" name:"IncludeVirtualZones"`

	// Whether to display all AZs in a region and the user’s permissions in each AZ.
	ShowPermission *bool `json:"ShowPermission,omitnil,omitempty" name:"ShowPermission"`
}

type DescribeZonesRequest struct {
	*tchttp.BaseRequest
	
	// Whether the virtual zone is included.–
	IncludeVirtualZones *bool `json:"IncludeVirtualZones,omitnil,omitempty" name:"IncludeVirtualZones"`

	// Whether to display all AZs in a region and the user’s permissions in each AZ.
	ShowPermission *bool `json:"ShowPermission,omitnil,omitempty" name:"ShowPermission"`
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
	RegionSet []*SaleRegion `json:"RegionSet,omitnil,omitempty" name:"RegionSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type ErrorLogItemExport struct {
	// Time Note: This field may return null, indicating that no valid values can be obtained.
	Timestamp *string `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// Log level. Valid values: `error`, `warning`, `note`. Note: This field may return null, indicating that no valid values can be obtained.
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// Log content Note: This field may return null, indicating that no valid values can be obtained.
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

// Predefined struct for user
type ExportInstanceErrorLogsRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Log start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Log end time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// The max number of returned results
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Log level
	LogLevels []*string `json:"LogLevels,omitnil,omitempty" name:"LogLevels"`


	KeyWords []*string `json:"KeyWords,omitnil,omitempty" name:"KeyWords"`

	// The template type. Valid values: `csv`, `original`.
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// Valid value: `Timestamp`
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Valid values: `ASC` or `DESC`.
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

type ExportInstanceErrorLogsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Log start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Log end time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// The max number of returned results
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Log level
	LogLevels []*string `json:"LogLevels,omitnil,omitempty" name:"LogLevels"`

	KeyWords []*string `json:"KeyWords,omitnil,omitempty" name:"KeyWords"`

	// The template type. Valid values: `csv`, `original`.
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// Valid value: `Timestamp`
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Valid values: `ASC` or `DESC`.
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

func (r *ExportInstanceErrorLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportInstanceErrorLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "LogLevels")
	delete(f, "KeyWords")
	delete(f, "FileType")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportInstanceErrorLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportInstanceErrorLogsResponseParams struct {
	// Export content of the error log Note: This field may return null, indicating that no valid values can be obtained.
	ErrorLogItems []*ErrorLogItemExport `json:"ErrorLogItems,omitnil,omitempty" name:"ErrorLogItems"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ExportInstanceErrorLogsResponse struct {
	*tchttp.BaseResponse
	Response *ExportInstanceErrorLogsResponseParams `json:"Response"`
}

func (r *ExportInstanceErrorLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportInstanceErrorLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportInstanceSlowQueriesRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Transaction start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Transaction end time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Maximum number
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Username
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// Client host
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// Database name
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// File type. Valid values: csv, original.
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// Sorting field. Valid values: `QueryTime`, `LockTime`, `RowsExamined`, and `RowsSent`.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: `asc`, `desc`.
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

type ExportInstanceSlowQueriesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Transaction start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Transaction end time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Maximum number
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Username
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// Client host
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// Database name
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// File type. Valid values: csv, original.
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// Sorting field. Valid values: `QueryTime`, `LockTime`, `RowsExamined`, and `RowsSent`.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: `asc`, `desc`.
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
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
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportInstanceSlowQueriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportInstanceSlowQueriesResponseParams struct {
	// Slow query export content
	FileContent *string `json:"FileContent,omitnil,omitempty" name:"FileContent"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type InputAccount struct {
	// Account
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// Host. Default value: `%`
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

// Predefined struct for user
type InquirePriceCreateRequestParams struct {
	// AZ
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Number of compute node to purchase
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// Instance type for purchase. Valid values: `PREPAID`, `POSTPAID`, `SERVERLESS`.
	InstancePayMode *string `json:"InstancePayMode,omitnil,omitempty" name:"InstancePayMode"`

	// Storage type for purchase. Valid values: `PREPAID`, `POSTPAID`.
	StoragePayMode *string `json:"StoragePayMode,omitnil,omitempty" name:"StoragePayMode"`

	// device type:common, exclusive
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// Number of CPU cores, which is required when `InstancePayMode` is `PREPAID` or `POSTPAID`.
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Memory size in GB, which is required when `InstancePayMode` is `PREPAID` or `POSTPAID`.
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// CCU size, which is required when `InstancePayMode` is `SERVERLESS`.
	Ccu *float64 `json:"Ccu,omitnil,omitempty" name:"Ccu"`

	// Storage size, which is required when `StoragePayMode` is `PREPAID`.
	StorageLimit *int64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`

	// Validity period, which is required when `InstancePayMode` is `PREPAID`.
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// Duration unit, which is required when `InstancePayMode` is `PREPAID`. Valid values: `m` (month), `d` (day).
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`
}

type InquirePriceCreateRequest struct {
	*tchttp.BaseRequest
	
	// AZ
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Number of compute node to purchase
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// Instance type for purchase. Valid values: `PREPAID`, `POSTPAID`, `SERVERLESS`.
	InstancePayMode *string `json:"InstancePayMode,omitnil,omitempty" name:"InstancePayMode"`

	// Storage type for purchase. Valid values: `PREPAID`, `POSTPAID`.
	StoragePayMode *string `json:"StoragePayMode,omitnil,omitempty" name:"StoragePayMode"`

	// device type:common, exclusive
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// Number of CPU cores, which is required when `InstancePayMode` is `PREPAID` or `POSTPAID`.
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Memory size in GB, which is required when `InstancePayMode` is `PREPAID` or `POSTPAID`.
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// CCU size, which is required when `InstancePayMode` is `SERVERLESS`.
	Ccu *float64 `json:"Ccu,omitnil,omitempty" name:"Ccu"`

	// Storage size, which is required when `StoragePayMode` is `PREPAID`.
	StorageLimit *int64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`

	// Validity period, which is required when `InstancePayMode` is `PREPAID`.
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// Duration unit, which is required when `InstancePayMode` is `PREPAID`. Valid values: `m` (month), `d` (day).
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`
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
	delete(f, "DeviceType")
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
	InstancePrice *TradePrice `json:"InstancePrice,omitnil,omitempty" name:"InstancePrice"`

	// Storage price
	StoragePrice *TradePrice `json:"StoragePrice,omitnil,omitempty" name:"StoragePrice"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Validity period, which needs to be used together with `TimeUnit`.
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// Unit of validity period, which needs to be used together with `TimeSpan`. Valid values: `d` (day), `m` (month).
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`
}

type InquirePriceRenewRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Validity period, which needs to be used together with `TimeUnit`.
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// Unit of validity period, which needs to be used together with `TimeSpan`. Valid values: `d` (day), `m` (month).
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Instance ID list
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Price of instance specification in array
	Prices []*TradePrice `json:"Prices,omitnil,omitempty" name:"Prices"`

	// Total renewal price of compute node
	InstanceRealTotalPrice *int64 `json:"InstanceRealTotalPrice,omitnil,omitempty" name:"InstanceRealTotalPrice"`

	// Total renewal price of storage node
	StorageRealTotalPrice *int64 `json:"StorageRealTotalPrice,omitnil,omitempty" name:"StorageRealTotalPrice"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Whether the audit is rule audit. Valid values: `true` (rule audit), `false` (full audit).
	// Note: This field may return null, indicating that no valid values can be obtained.
	AuditRule *bool `json:"AuditRule,omitnil,omitempty" name:"AuditRule"`

	// Audit rule details, which is valid only when `AuditRule` is `true`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AuditRuleFilters []*AuditRuleFilters `json:"AuditRuleFilters,omitnil,omitempty" name:"AuditRuleFilters"`
}

type InstanceInitInfo struct {
	// Instance CPU
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Instance memory
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Instance type. Valid values:`rw`, `ro`.
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Number of the instances. Value range: 1-15.
	InstanceCount *int64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// Minimum number of serverless instances. Value range: 1-15.
	MinRoCount *int64 `json:"MinRoCount,omitnil,omitempty" name:"MinRoCount"`

	// Maximum number of serverless instances. Value range: 1-15.
	MaxRoCount *int64 `json:"MaxRoCount,omitnil,omitempty" name:"MaxRoCount"`

	// Minimum specifications for serverless instance
	MinRoCpu *float64 `json:"MinRoCpu,omitnil,omitempty" name:"MinRoCpu"`

	// Maximum specifications for serverless instance
	MaxRoCpu *float64 `json:"MaxRoCpu,omitnil,omitempty" name:"MaxRoCpu"`
}

type InstanceNetInfo struct {
	// Network type
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceGroupType *string `json:"InstanceGroupType,omitnil,omitempty" name:"InstanceGroupType"`

	// Instance group ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceGroupId *string `json:"InstanceGroupId,omitnil,omitempty" name:"InstanceGroupId"`

	// VPC ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Network type
	// Note: This field may return null, indicating that no valid values can be obtained.
	NetType *int64 `json:"NetType,omitnil,omitempty" name:"NetType"`

	// VPC IP
	// Note: This field may return null, indicating that no valid values can be obtained.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// VPC port
	// Note: This field may return null, indicating that no valid values can be obtained.
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// Public network domain name
	// Note: This field may return null, indicating that no valid values can be obtained.
	WanDomain *string `json:"WanDomain,omitnil,omitempty" name:"WanDomain"`


	WanIP *string `json:"WanIP,omitnil,omitempty" name:"WanIP"`

	// Public network port
	// Note: This field may return null, indicating that no valid values can be obtained.
	WanPort *int64 `json:"WanPort,omitnil,omitempty" name:"WanPort"`

	// Public network status
	// Note: This field may return null, indicating that no valid values can be obtained.
	WanStatus *string `json:"WanStatus,omitnil,omitempty" name:"WanStatus"`
}

type InstanceParamItem struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// List of instance parameters
	ParamsItems []*ParamItemDetail `json:"ParamsItems,omitnil,omitempty" name:"ParamsItems"`
}

type InstanceSpec struct {
	// Number of instance CPU cores
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Instance memory in GB
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Maximum instance storage capacity GB
	MaxStorageSize *uint64 `json:"MaxStorageSize,omitnil,omitempty" name:"MaxStorageSize"`

	// Minimum instance storage capacity GB
	MinStorageSize *uint64 `json:"MinStorageSize,omitnil,omitempty" name:"MinStorageSize"`

	// Whether there is an inventory.
	HasStock *bool `json:"HasStock,omitnil,omitempty" name:"HasStock"`

	// Machine type
	MachineType *string `json:"MachineType,omitnil,omitempty" name:"MachineType"`

	// Maximum IOPS
	MaxIops *int64 `json:"MaxIops,omitnil,omitempty" name:"MaxIops"`

	// Maximum bandwidth
	MaxIoBandWidth *int64 `json:"MaxIoBandWidth,omitnil,omitempty" name:"MaxIoBandWidth"`

	// Inventory information in a region
	// Note: This field may return null, indicating that no valid values can be obtained.
	ZoneStockInfos []*ZoneStockInfo `json:"ZoneStockInfos,omitnil,omitempty" name:"ZoneStockInfos"`

	// Quantity in stock
	// Note: This field may return null, indicating that no valid values can be obtained.
	StockCount *int64 `json:"StockCount,omitnil,omitempty" name:"StockCount"`
}

// Predefined struct for user
type IsolateClusterRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// This parameter has been disused.
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`
}

type IsolateClusterRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// This parameter has been disused.
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`
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
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Refund order ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Instance ID array
	InstanceIdList []*string `json:"InstanceIdList,omitnil,omitempty" name:"InstanceIdList"`

	// This parameter has been disused.
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`
}

type IsolateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Instance ID array
	InstanceIdList []*string `json:"InstanceIdList,omitnil,omitempty" name:"InstanceIdList"`

	// This parameter has been disused.
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`
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
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Order ID for isolated instance (prepaid instance)
	// Note: this field may return null, indicating that no valid values can be obtained.
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyAccountDescriptionRequestParams struct {
	// Database account name
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// Database account description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Host. Default value: `%`
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

type ModifyAccountDescriptionRequest struct {
	*tchttp.BaseRequest
	
	// Database account name
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// Database account description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Host. Default value: `%`
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
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
	delete(f, "AccountName")
	delete(f, "Description")
	delete(f, "ClusterId")
	delete(f, "Host")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccountDescriptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountDescriptionResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAccountDescriptionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccountDescriptionResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyAccountHostRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// New host
	NewHost *string `json:"NewHost,omitnil,omitempty" name:"NewHost"`

	// Account infomation
	Account *InputAccount `json:"Account,omitnil,omitempty" name:"Account"`
}

type ModifyAccountHostRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// New host
	NewHost *string `json:"NewHost,omitnil,omitempty" name:"NewHost"`

	// Account infomation
	Account *InputAccount `json:"Account,omitnil,omitempty" name:"Account"`
}

func (r *ModifyAccountHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountHostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NewHost")
	delete(f, "Account")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccountHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountHostResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAccountHostResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccountHostResponseParams `json:"Response"`
}

func (r *ModifyAccountHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountPrivilegesRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Account infomation
	Account *InputAccount `json:"Account,omitnil,omitempty" name:"Account"`

	// Array of global permissions
	GlobalPrivileges []*string `json:"GlobalPrivileges,omitnil,omitempty" name:"GlobalPrivileges"`

	// Array of database permissions
	DatabasePrivileges []*DatabasePrivileges `json:"DatabasePrivileges,omitnil,omitempty" name:"DatabasePrivileges"`

	// Array of table permissions
	TablePrivileges []*TablePrivileges `json:"TablePrivileges,omitnil,omitempty" name:"TablePrivileges"`
}

type ModifyAccountPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Account infomation
	Account *InputAccount `json:"Account,omitnil,omitempty" name:"Account"`

	// Array of global permissions
	GlobalPrivileges []*string `json:"GlobalPrivileges,omitnil,omitempty" name:"GlobalPrivileges"`

	// Array of database permissions
	DatabasePrivileges []*DatabasePrivileges `json:"DatabasePrivileges,omitnil,omitempty" name:"DatabasePrivileges"`

	// Array of table permissions
	TablePrivileges []*TablePrivileges `json:"TablePrivileges,omitnil,omitempty" name:"TablePrivileges"`
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
	delete(f, "ClusterId")
	delete(f, "Account")
	delete(f, "GlobalPrivileges")
	delete(f, "DatabasePrivileges")
	delete(f, "TablePrivileges")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccountPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountPrivilegesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAccountPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccountPrivilegesResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyAuditRuleTemplatesRequestParams struct {
	// Audit rule template ID
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`

	// Audit rule after modification
	RuleFilters []*RuleFilters `json:"RuleFilters,omitnil,omitempty" name:"RuleFilters"`

	// New name of the rule template
	RuleTemplateName *string `json:"RuleTemplateName,omitnil,omitempty" name:"RuleTemplateName"`

	// New description of the rule template
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ModifyAuditRuleTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// Audit rule template ID
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`

	// Audit rule after modification
	RuleFilters []*RuleFilters `json:"RuleFilters,omitnil,omitempty" name:"RuleFilters"`

	// New name of the rule template
	RuleTemplateName *string `json:"RuleTemplateName,omitnil,omitempty" name:"RuleTemplateName"`

	// New description of the rule template
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
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
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Log retention period
	LogExpireDay *uint64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// Frequent log retention period
	HighLogExpireDay *uint64 `json:"HighLogExpireDay,omitnil,omitempty" name:"HighLogExpireDay"`

	// The parameter used to change the audit rule of the instance to full audit
	AuditAll *bool `json:"AuditAll,omitnil,omitempty" name:"AuditAll"`

	// Rule audit
	AuditRuleFilters []*AuditRuleFilters `json:"AuditRuleFilters,omitnil,omitempty" name:"AuditRuleFilters"`

	// Rule template ID
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`
}

type ModifyAuditServiceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Log retention period
	LogExpireDay *uint64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// Frequent log retention period
	HighLogExpireDay *uint64 `json:"HighLogExpireDay,omitnil,omitempty" name:"HighLogExpireDay"`

	// The parameter used to change the audit rule of the instance to full audit
	AuditAll *bool `json:"AuditAll,omitnil,omitempty" name:"AuditAll"`

	// Rule audit
	AuditRuleFilters []*AuditRuleFilters `json:"AuditRuleFilters,omitnil,omitempty" name:"AuditRuleFilters"`

	// Rule template ID
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`
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
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Full backup start time. Value range: [0-24*3600]. For example, 0:00 AM, 1:00 AM, and 2:00 AM are represented by 0, 3600, and 7200, respectively
	BackupTimeBeg *uint64 `json:"BackupTimeBeg,omitnil,omitempty" name:"BackupTimeBeg"`

	// Full backup end time. Value range: [0-24*3600]. For example, 0:00 AM, 1:00 AM, and 2:00 AM are represented by 0, 3600, and 7200, respectively.
	BackupTimeEnd *uint64 `json:"BackupTimeEnd,omitnil,omitempty" name:"BackupTimeEnd"`

	// Backup retention period in seconds. Backups will be cleared after this period elapses. 7 days is represented by 3600*24*7 = 604800. Maximum value: 158112000.
	ReserveDuration *uint64 `json:"ReserveDuration,omitnil,omitempty" name:"ReserveDuration"`

	// Backup frequency. It is an array of 7 elements corresponding to Monday through Sunday. full: full backup; increment: incremental backup. This parameter cannot be modified currently and doesn't need to be entered.
	BackupFreq []*string `json:"BackupFreq,omitnil,omitempty" name:"BackupFreq"`

	// Backup mode. logic: logic backup; snapshot: snapshot backup. This parameter cannot be modified currently and doesn't need to be entered.
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`
}

type ModifyBackupConfigRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Full backup start time. Value range: [0-24*3600]. For example, 0:00 AM, 1:00 AM, and 2:00 AM are represented by 0, 3600, and 7200, respectively
	BackupTimeBeg *uint64 `json:"BackupTimeBeg,omitnil,omitempty" name:"BackupTimeBeg"`

	// Full backup end time. Value range: [0-24*3600]. For example, 0:00 AM, 1:00 AM, and 2:00 AM are represented by 0, 3600, and 7200, respectively.
	BackupTimeEnd *uint64 `json:"BackupTimeEnd,omitnil,omitempty" name:"BackupTimeEnd"`

	// Backup retention period in seconds. Backups will be cleared after this period elapses. 7 days is represented by 3600*24*7 = 604800. Maximum value: 158112000.
	ReserveDuration *uint64 `json:"ReserveDuration,omitnil,omitempty" name:"ReserveDuration"`

	// Backup frequency. It is an array of 7 elements corresponding to Monday through Sunday. full: full backup; increment: incremental backup. This parameter cannot be modified currently and doesn't need to be entered.
	BackupFreq []*string `json:"BackupFreq,omitnil,omitempty" name:"BackupFreq"`

	// Backup mode. logic: logic backup; snapshot: snapshot backup. This parameter cannot be modified currently and doesn't need to be entered.
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`
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
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Backup file ID
	BackupId *int64 `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// Backup name, which can contain up to 60 characters.
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`
}

type ModifyBackupNameRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Backup file ID
	BackupId *int64 `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// Backup name, which can contain up to 60 characters.
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`
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
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyBinlogSaveDaysRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Binlog retention period in days
	BinlogSaveDays *int64 `json:"BinlogSaveDays,omitnil,omitempty" name:"BinlogSaveDays"`
}

type ModifyBinlogSaveDaysRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Binlog retention period in days
	BinlogSaveDays *int64 `json:"BinlogSaveDays,omitnil,omitempty" name:"BinlogSaveDays"`
}

func (r *ModifyBinlogSaveDaysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBinlogSaveDaysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "BinlogSaveDays")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBinlogSaveDaysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBinlogSaveDaysResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyBinlogSaveDaysResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBinlogSaveDaysResponseParams `json:"Response"`
}

func (r *ModifyBinlogSaveDaysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBinlogSaveDaysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterDatabaseRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database name
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Host permissions of the new authorized user
	NewUserHostPrivileges []*UserHostPrivilege `json:"NewUserHostPrivileges,omitnil,omitempty" name:"NewUserHostPrivileges"`

	// Remarks
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Host permissions of previously authorized users
	OldUserHostPrivileges []*UserHostPrivilege `json:"OldUserHostPrivileges,omitnil,omitempty" name:"OldUserHostPrivileges"`
}

type ModifyClusterDatabaseRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database name
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Host permissions of the new authorized user
	NewUserHostPrivileges []*UserHostPrivilege `json:"NewUserHostPrivileges,omitnil,omitempty" name:"NewUserHostPrivileges"`

	// Remarks
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Host permissions of previously authorized users
	OldUserHostPrivileges []*UserHostPrivilege `json:"OldUserHostPrivileges,omitnil,omitempty" name:"OldUserHostPrivileges"`
}

func (r *ModifyClusterDatabaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterDatabaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "DbName")
	delete(f, "NewUserHostPrivileges")
	delete(f, "Description")
	delete(f, "OldUserHostPrivileges")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterDatabaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterDatabaseResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyClusterDatabaseResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterDatabaseResponseParams `json:"Response"`
}

func (r *ModifyClusterDatabaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterDatabaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterNameRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Cluster name
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`
}

type ModifyClusterNameRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Cluster name
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`
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
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// List of the parameters to be modified. Each element in the list is a combination of `ParamName`, `CurrentValue`, and `OldValue`. `ParamName` is the parameter name; `CurrentValue` is the current value; `OldValue` is the old value that doesn’t need to be verified.
	ParamList []*ParamItem `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// Valid values: `yes` (execute during maintenance time), `no` (execute now)
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
}

type ModifyClusterParamRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// List of the parameters to be modified. Each element in the list is a combination of `ParamName`, `CurrentValue`, and `OldValue`. `ParamName` is the parameter name; `CurrentValue` is the current value; `OldValue` is the old value that doesn’t need to be verified.
	ParamList []*ParamItem `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// Valid values: `yes` (execute during maintenance time), `no` (execute now)
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
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
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyClusterPasswordComplexityRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Password length
	ValidatePasswordLength *int64 `json:"ValidatePasswordLength,omitnil,omitempty" name:"ValidatePasswordLength"`

	// Number of letters
	ValidatePasswordMixedCaseCount *int64 `json:"ValidatePasswordMixedCaseCount,omitnil,omitempty" name:"ValidatePasswordMixedCaseCount"`

	// Number of symbols
	ValidatePasswordSpecialCharCount *int64 `json:"ValidatePasswordSpecialCharCount,omitnil,omitempty" name:"ValidatePasswordSpecialCharCount"`

	// Number of digits
	ValidatePasswordNumberCount *int64 `json:"ValidatePasswordNumberCount,omitnil,omitempty" name:"ValidatePasswordNumberCount"`

	// Password strength. Valid values: `MEDIUM`, `STRONG`.
	ValidatePasswordPolicy *string `json:"ValidatePasswordPolicy,omitnil,omitempty" name:"ValidatePasswordPolicy"`

	// Data dictionary
	ValidatePasswordDictionary []*string `json:"ValidatePasswordDictionary,omitnil,omitempty" name:"ValidatePasswordDictionary"`
}

type ModifyClusterPasswordComplexityRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Password length
	ValidatePasswordLength *int64 `json:"ValidatePasswordLength,omitnil,omitempty" name:"ValidatePasswordLength"`

	// Number of letters
	ValidatePasswordMixedCaseCount *int64 `json:"ValidatePasswordMixedCaseCount,omitnil,omitempty" name:"ValidatePasswordMixedCaseCount"`

	// Number of symbols
	ValidatePasswordSpecialCharCount *int64 `json:"ValidatePasswordSpecialCharCount,omitnil,omitempty" name:"ValidatePasswordSpecialCharCount"`

	// Number of digits
	ValidatePasswordNumberCount *int64 `json:"ValidatePasswordNumberCount,omitnil,omitempty" name:"ValidatePasswordNumberCount"`

	// Password strength. Valid values: `MEDIUM`, `STRONG`.
	ValidatePasswordPolicy *string `json:"ValidatePasswordPolicy,omitnil,omitempty" name:"ValidatePasswordPolicy"`

	// Data dictionary
	ValidatePasswordDictionary []*string `json:"ValidatePasswordDictionary,omitnil,omitempty" name:"ValidatePasswordDictionary"`
}

func (r *ModifyClusterPasswordComplexityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterPasswordComplexityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ValidatePasswordLength")
	delete(f, "ValidatePasswordMixedCaseCount")
	delete(f, "ValidatePasswordSpecialCharCount")
	delete(f, "ValidatePasswordNumberCount")
	delete(f, "ValidatePasswordPolicy")
	delete(f, "ValidatePasswordDictionary")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterPasswordComplexityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterPasswordComplexityResponseParams struct {
	// Task flow ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyClusterPasswordComplexityResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterPasswordComplexityResponseParams `json:"Response"`
}

func (r *ModifyClusterPasswordComplexityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterPasswordComplexityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterSlaveZoneRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Old replica AZ
	OldSlaveZone *string `json:"OldSlaveZone,omitnil,omitempty" name:"OldSlaveZone"`

	// New replica AZ
	NewSlaveZone *string `json:"NewSlaveZone,omitnil,omitempty" name:"NewSlaveZone"`
}

type ModifyClusterSlaveZoneRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Old replica AZ
	OldSlaveZone *string `json:"OldSlaveZone,omitnil,omitempty" name:"OldSlaveZone"`

	// New replica AZ
	NewSlaveZone *string `json:"NewSlaveZone,omitnil,omitempty" name:"NewSlaveZone"`
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
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// List of IDs of security groups to be modified, which is an array of one or more security group IDs.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// AZ
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

type ModifyDBInstanceSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Instance group ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// List of IDs of security groups to be modified, which is an array of one or more security group IDs.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// AZ
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
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
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

type ModifyInstanceNameRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
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
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyInstanceParamRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Instance ID
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// List of cluster parameters
	ClusterParamList []*ModifyParamItem `json:"ClusterParamList,omitnil,omitempty" name:"ClusterParamList"`

	// List of the instance parameters
	InstanceParamList []*ModifyParamItem `json:"InstanceParamList,omitnil,omitempty" name:"InstanceParamList"`

	// Valid values: `yes` (modify in maintenance window),  `no`  (execute now by default).
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
}

type ModifyInstanceParamRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Instance ID
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// List of cluster parameters
	ClusterParamList []*ModifyParamItem `json:"ClusterParamList,omitnil,omitempty" name:"ClusterParamList"`

	// List of the instance parameters
	InstanceParamList []*ModifyParamItem `json:"InstanceParamList,omitnil,omitempty" name:"InstanceParamList"`

	// Valid values: `yes` (modify in maintenance window),  `no`  (execute now by default).
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
}

func (r *ModifyInstanceParamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceParamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceIds")
	delete(f, "ClusterParamList")
	delete(f, "InstanceParamList")
	delete(f, "IsInMaintainPeriod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceParamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceParamResponseParams struct {
	// Task ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceParamResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceParamResponseParams `json:"Response"`
}

func (r *ModifyInstanceParamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceParamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMaintainPeriodConfigRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Maintenance start time in seconds. For example, 03:00 AM is represented by 10800
	MaintainStartTime *int64 `json:"MaintainStartTime,omitnil,omitempty" name:"MaintainStartTime"`

	// Maintenance duration in seconds. For example, one hour is represented by 3600
	MaintainDuration *int64 `json:"MaintainDuration,omitnil,omitempty" name:"MaintainDuration"`

	// Maintenance days of the week. Valid values: [Mon, Tue, Wed, Thu, Fri, Sat, Sun].
	MaintainWeekDays []*string `json:"MaintainWeekDays,omitnil,omitempty" name:"MaintainWeekDays"`
}

type ModifyMaintainPeriodConfigRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Maintenance start time in seconds. For example, 03:00 AM is represented by 10800
	MaintainStartTime *int64 `json:"MaintainStartTime,omitnil,omitempty" name:"MaintainStartTime"`

	// Maintenance duration in seconds. For example, one hour is represented by 3600
	MaintainDuration *int64 `json:"MaintainDuration,omitnil,omitempty" name:"MaintainDuration"`

	// Maintenance days of the week. Valid values: [Mon, Tue, Wed, Thu, Fri, Sat, Sun].
	MaintainWeekDays []*string `json:"MaintainWeekDays,omitnil,omitempty" name:"MaintainWeekDays"`
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
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// Current parameter value
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`

	// Old parameter value, which is used only in output parameters.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OldValue *string `json:"OldValue,omitnil,omitempty" name:"OldValue"`
}

// Predefined struct for user
type ModifyParamTemplateRequestParams struct {
	// Template ID
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Template name
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// Template description
	TemplateDescription *string `json:"TemplateDescription,omitnil,omitempty" name:"TemplateDescription"`

	// Parameter list
	ParamList []*ModifyParamItem `json:"ParamList,omitnil,omitempty" name:"ParamList"`
}

type ModifyParamTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Template ID
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Template name
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// Template description
	TemplateDescription *string `json:"TemplateDescription,omitnil,omitempty" name:"TemplateDescription"`

	// Parameter list
	ParamList []*ModifyParamItem `json:"ParamList,omitnil,omitempty" name:"ParamList"`
}

func (r *ModifyParamTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyParamTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "TemplateName")
	delete(f, "TemplateDescription")
	delete(f, "ParamList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyParamTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyParamTemplateResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyParamTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyParamTemplateResponseParams `json:"Response"`
}

func (r *ModifyParamTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyParamTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProxyDescRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Database proxy description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ModifyProxyDescRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Database proxy description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *ModifyProxyDescRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProxyDescRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ProxyGroupId")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProxyDescRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProxyDescResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyProxyDescResponse struct {
	*tchttp.BaseResponse
	Response *ModifyProxyDescResponseParams `json:"Response"`
}

func (r *ModifyProxyDescResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProxyDescResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProxyRwSplitRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Consistency type. Valid values: `eventual` (eventual consistency), `session` (session consistency), `global` (global consistency).
	ConsistencyType *string `json:"ConsistencyType,omitnil,omitempty" name:"ConsistencyType"`

	// Consistency timeout period
	ConsistencyTimeOut *string `json:"ConsistencyTimeOut,omitnil,omitempty" name:"ConsistencyTimeOut"`

	// Assignment mode of read/write weights. Valid values: `system` (auto-assigned), `custom`
	WeightMode *string `json:"WeightMode,omitnil,omitempty" name:"WeightMode"`

	// Read-Only weight of an instance
	InstanceWeights []*ProxyInstanceWeight `json:"InstanceWeights,omitnil,omitempty" name:"InstanceWeights"`

	// Whether to enable failover. If it is enabled, the connection address will route requests to the source instance in case of proxy failure. Valid values: `true`, `false`.
	FailOver *string `json:"FailOver,omitnil,omitempty" name:"FailOver"`

	// Whether to automatically add read-only instances. Valid values: `true`, `false`
	AutoAddRo *string `json:"AutoAddRo,omitnil,omitempty" name:"AutoAddRo"`

	// Whether to enable read/write separation
	OpenRw *string `json:"OpenRw,omitnil,omitempty" name:"OpenRw"`

	// Read/Write type. Valid values:
	// `READWRITE`, `READONLY`.
	RwType *string `json:"RwType,omitnil,omitempty" name:"RwType"`

	// Transaction split
	TransSplit *bool `json:"TransSplit,omitnil,omitempty" name:"TransSplit"`

	// Connection mode. Valid values:
	// `nearby`, `balance`.
	AccessMode *string `json:"AccessMode,omitnil,omitempty" name:"AccessMode"`

	// Whether to enable the connection pool. Valid values: 
	// `yes`, `no`.
	OpenConnectionPool *string `json:"OpenConnectionPool,omitnil,omitempty" name:"OpenConnectionPool"`

	// Connection pool type. Valid values:
	// `ConnectionPoolType`, `SessionConnectionPool`.
	ConnectionPoolType *string `json:"ConnectionPoolType,omitnil,omitempty" name:"ConnectionPoolType"`

	// Connection persistence timeout
	ConnectionPoolTimeOut *int64 `json:"ConnectionPoolTimeOut,omitnil,omitempty" name:"ConnectionPoolTimeOut"`
}

type ModifyProxyRwSplitRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Consistency type. Valid values: `eventual` (eventual consistency), `session` (session consistency), `global` (global consistency).
	ConsistencyType *string `json:"ConsistencyType,omitnil,omitempty" name:"ConsistencyType"`

	// Consistency timeout period
	ConsistencyTimeOut *string `json:"ConsistencyTimeOut,omitnil,omitempty" name:"ConsistencyTimeOut"`

	// Assignment mode of read/write weights. Valid values: `system` (auto-assigned), `custom`
	WeightMode *string `json:"WeightMode,omitnil,omitempty" name:"WeightMode"`

	// Read-Only weight of an instance
	InstanceWeights []*ProxyInstanceWeight `json:"InstanceWeights,omitnil,omitempty" name:"InstanceWeights"`

	// Whether to enable failover. If it is enabled, the connection address will route requests to the source instance in case of proxy failure. Valid values: `true`, `false`.
	FailOver *string `json:"FailOver,omitnil,omitempty" name:"FailOver"`

	// Whether to automatically add read-only instances. Valid values: `true`, `false`
	AutoAddRo *string `json:"AutoAddRo,omitnil,omitempty" name:"AutoAddRo"`

	// Whether to enable read/write separation
	OpenRw *string `json:"OpenRw,omitnil,omitempty" name:"OpenRw"`

	// Read/Write type. Valid values:
	// `READWRITE`, `READONLY`.
	RwType *string `json:"RwType,omitnil,omitempty" name:"RwType"`

	// Transaction split
	TransSplit *bool `json:"TransSplit,omitnil,omitempty" name:"TransSplit"`

	// Connection mode. Valid values:
	// `nearby`, `balance`.
	AccessMode *string `json:"AccessMode,omitnil,omitempty" name:"AccessMode"`

	// Whether to enable the connection pool. Valid values: 
	// `yes`, `no`.
	OpenConnectionPool *string `json:"OpenConnectionPool,omitnil,omitempty" name:"OpenConnectionPool"`

	// Connection pool type. Valid values:
	// `ConnectionPoolType`, `SessionConnectionPool`.
	ConnectionPoolType *string `json:"ConnectionPoolType,omitnil,omitempty" name:"ConnectionPoolType"`

	// Connection persistence timeout
	ConnectionPoolTimeOut *int64 `json:"ConnectionPoolTimeOut,omitnil,omitempty" name:"ConnectionPoolTimeOut"`
}

func (r *ModifyProxyRwSplitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProxyRwSplitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ProxyGroupId")
	delete(f, "ConsistencyType")
	delete(f, "ConsistencyTimeOut")
	delete(f, "WeightMode")
	delete(f, "InstanceWeights")
	delete(f, "FailOver")
	delete(f, "AutoAddRo")
	delete(f, "OpenRw")
	delete(f, "RwType")
	delete(f, "TransSplit")
	delete(f, "AccessMode")
	delete(f, "OpenConnectionPool")
	delete(f, "ConnectionPoolType")
	delete(f, "ConnectionPoolTimeOut")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProxyRwSplitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProxyRwSplitResponseParams struct {
	// Async FlowId
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Async task ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyProxyRwSplitResponse struct {
	*tchttp.BaseResponse
	Response *ModifyProxyRwSplitResponseParams `json:"Response"`
}

func (r *ModifyProxyRwSplitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProxyRwSplitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourcePackageClustersRequestParams struct {
	// The unique ID of a resource pack
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// ID of the cluster to be bound
	BindClusterIds []*string `json:"BindClusterIds,omitnil,omitempty" name:"BindClusterIds"`

	// ID of the cluster to be unbound
	UnbindClusterIds []*string `json:"UnbindClusterIds,omitnil,omitempty" name:"UnbindClusterIds"`
}

type ModifyResourcePackageClustersRequest struct {
	*tchttp.BaseRequest
	
	// The unique ID of a resource pack
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// ID of the cluster to be bound
	BindClusterIds []*string `json:"BindClusterIds,omitnil,omitempty" name:"BindClusterIds"`

	// ID of the cluster to be unbound
	UnbindClusterIds []*string `json:"UnbindClusterIds,omitnil,omitempty" name:"UnbindClusterIds"`
}

func (r *ModifyResourcePackageClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourcePackageClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PackageId")
	delete(f, "BindClusterIds")
	delete(f, "UnbindClusterIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyResourcePackageClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourcePackageClustersResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyResourcePackageClustersResponse struct {
	*tchttp.BaseResponse
	Response *ModifyResourcePackageClustersResponseParams `json:"Response"`
}

func (r *ModifyResourcePackageClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourcePackageClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourcePackageNameRequestParams struct {
	// The unique ID of a resource pack
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// Custom resource pack name, which can contains up to 120 characters.
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`
}

type ModifyResourcePackageNameRequest struct {
	*tchttp.BaseRequest
	
	// The unique ID of a resource pack
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// Custom resource pack name, which can contains up to 120 characters.
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`
}

func (r *ModifyResourcePackageNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourcePackageNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PackageId")
	delete(f, "PackageName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyResourcePackageNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourcePackageNameResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyResourcePackageNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyResourcePackageNameResponseParams `json:"Response"`
}

func (r *ModifyResourcePackageNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourcePackageNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVipVportRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Instance group ID
	InstanceGrpId *string `json:"InstanceGrpId,omitnil,omitempty" name:"InstanceGrpId"`

	// Target IP to be modified
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Target port to be modified
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// Database type. Valid values: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Valid hours of old IPs. If it is set to `0` hours, the IPs will be released immediately.
	OldIpReserveHours *int64 `json:"OldIpReserveHours,omitnil,omitempty" name:"OldIpReserveHours"`
}

type ModifyVipVportRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Instance group ID
	InstanceGrpId *string `json:"InstanceGrpId,omitnil,omitempty" name:"InstanceGrpId"`

	// Target IP to be modified
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Target port to be modified
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// Database type. Valid values: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Valid hours of old IPs. If it is set to `0` hours, the IPs will be released immediately.
	OldIpReserveHours *int64 `json:"OldIpReserveHours,omitnil,omitempty" name:"OldIpReserveHours"`
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
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	IsDisable *string `json:"IsDisable,omitnil,omitempty" name:"IsDisable"`

	// Module name
	ModuleName *string `json:"ModuleName,omitnil,omitempty" name:"ModuleName"`
}

type NetAddr struct {
	// Private network IP
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Private network port number
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// Public network domain name
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	WanDomain *string `json:"WanDomain,omitnil,omitempty" name:"WanDomain"`

	// Public network port number
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	WanPort *int64 `json:"WanPort,omitnil,omitempty" name:"WanPort"`

	// Network type. Valid values: `ro` (read-only), `rw` or `ha` (read-write)
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// Subnet ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// VPC ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// Description
	// Note: This field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Public IP
	// Note: This field may return null, indicating that no valid values can be obtained.
	WanIP *string `json:"WanIP,omitnil,omitempty" name:"WanIP"`

	// Public network status
	// Note: This field may return null, indicating that no valid values can be obtained.
	WanStatus *string `json:"WanStatus,omitnil,omitempty" name:"WanStatus"`

	// Instance group ID Note: This field may return null, indicating that no valid values can be obtained.
	InstanceGroupId *string `json:"InstanceGroupId,omitnil,omitempty" name:"InstanceGroupId"`
}

type NewAccount struct {
	// Account name, which can contain 1-16 letters, digits, and underscores. It must begin with a letter and end with a letter or digit.
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// Password, which can contain 8-64 characters.
	AccountPassword *string `json:"AccountPassword,omitnil,omitempty" name:"AccountPassword"`

	// Host
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// Description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Maximum number of user connections, which cannot be above 10,240.
	MaxUserConnections *int64 `json:"MaxUserConnections,omitnil,omitempty" name:"MaxUserConnections"`
}

type ObjectTask struct {
	// Auto-Incrementing task ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Task type
	// Note: this field may return null, indicating that no valid values can be obtained.
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// Task status
	// Note: this field may return null, indicating that no valid values can be obtained.
	TaskStatus *string `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// Task ID (cluster ID | instance group ID | instance ID)
	// Note: this field may return null, indicating that no valid values can be obtained.
	ObjectId *string `json:"ObjectId,omitnil,omitempty" name:"ObjectId"`

	// Task type
	// Note: this field may return null, indicating that no valid values can be obtained.
	ObjectType *string `json:"ObjectType,omitnil,omitempty" name:"ObjectType"`
}

// Predefined struct for user
type OfflineClusterRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type OfflineClusterRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
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
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Instance ID array
	InstanceIdList []*string `json:"InstanceIdList,omitnil,omitempty" name:"InstanceIdList"`
}

type OfflineInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Instance ID array
	InstanceIdList []*string `json:"InstanceIdList,omitnil,omitempty" name:"InstanceIdList"`
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
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Port
	// Note: This field may return null, indicating that no valid values can be obtained.
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// Expected valid hours of old IPs
	// Note: This field may return null, indicating that no valid values can be obtained.
	ReturnTime *string `json:"ReturnTime,omitnil,omitempty" name:"ReturnTime"`
}

// Predefined struct for user
type OpenAuditServiceRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Log retention period
	LogExpireDay *uint64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// Frequent log retention period
	HighLogExpireDay *uint64 `json:"HighLogExpireDay,omitnil,omitempty" name:"HighLogExpireDay"`

	// Audit rule. If both this parameter and `RuleTemplateIds` are left empty, full audit will be applied.
	AuditRuleFilters []*AuditRuleFilters `json:"AuditRuleFilters,omitnil,omitempty" name:"AuditRuleFilters"`

	// Rule template ID. If both this parameter and `AuditRuleFilters` are left empty, full audit will be applied.
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`
}

type OpenAuditServiceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Log retention period
	LogExpireDay *uint64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// Frequent log retention period
	HighLogExpireDay *uint64 `json:"HighLogExpireDay,omitnil,omitempty" name:"HighLogExpireDay"`

	// Audit rule. If both this parameter and `RuleTemplateIds` are left empty, full audit will be applied.
	AuditRuleFilters []*AuditRuleFilters `json:"AuditRuleFilters,omitnil,omitempty" name:"AuditRuleFilters"`

	// Rule template ID. If both this parameter and `AuditRuleFilters` are left empty, full audit will be applied.
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`
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
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type OpenClusterPasswordComplexityRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Password length
	ValidatePasswordLength *int64 `json:"ValidatePasswordLength,omitnil,omitempty" name:"ValidatePasswordLength"`

	// Number of letters
	ValidatePasswordMixedCaseCount *int64 `json:"ValidatePasswordMixedCaseCount,omitnil,omitempty" name:"ValidatePasswordMixedCaseCount"`

	// Number of symbols
	ValidatePasswordSpecialCharCount *int64 `json:"ValidatePasswordSpecialCharCount,omitnil,omitempty" name:"ValidatePasswordSpecialCharCount"`

	// Number of digits
	ValidatePasswordNumberCount *int64 `json:"ValidatePasswordNumberCount,omitnil,omitempty" name:"ValidatePasswordNumberCount"`

	// Password strength. Valid values: `MEDIUM`, `STRONG`.
	ValidatePasswordPolicy *string `json:"ValidatePasswordPolicy,omitnil,omitempty" name:"ValidatePasswordPolicy"`

	// Data dictionary
	ValidatePasswordDictionary []*string `json:"ValidatePasswordDictionary,omitnil,omitempty" name:"ValidatePasswordDictionary"`
}

type OpenClusterPasswordComplexityRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Password length
	ValidatePasswordLength *int64 `json:"ValidatePasswordLength,omitnil,omitempty" name:"ValidatePasswordLength"`

	// Number of letters
	ValidatePasswordMixedCaseCount *int64 `json:"ValidatePasswordMixedCaseCount,omitnil,omitempty" name:"ValidatePasswordMixedCaseCount"`

	// Number of symbols
	ValidatePasswordSpecialCharCount *int64 `json:"ValidatePasswordSpecialCharCount,omitnil,omitempty" name:"ValidatePasswordSpecialCharCount"`

	// Number of digits
	ValidatePasswordNumberCount *int64 `json:"ValidatePasswordNumberCount,omitnil,omitempty" name:"ValidatePasswordNumberCount"`

	// Password strength. Valid values: `MEDIUM`, `STRONG`.
	ValidatePasswordPolicy *string `json:"ValidatePasswordPolicy,omitnil,omitempty" name:"ValidatePasswordPolicy"`

	// Data dictionary
	ValidatePasswordDictionary []*string `json:"ValidatePasswordDictionary,omitnil,omitempty" name:"ValidatePasswordDictionary"`
}

func (r *OpenClusterPasswordComplexityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenClusterPasswordComplexityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ValidatePasswordLength")
	delete(f, "ValidatePasswordMixedCaseCount")
	delete(f, "ValidatePasswordSpecialCharCount")
	delete(f, "ValidatePasswordNumberCount")
	delete(f, "ValidatePasswordPolicy")
	delete(f, "ValidatePasswordDictionary")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenClusterPasswordComplexityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenClusterPasswordComplexityResponseParams struct {
	// Task flow ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OpenClusterPasswordComplexityResponse struct {
	*tchttp.BaseResponse
	Response *OpenClusterPasswordComplexityResponseParams `json:"Response"`
}

func (r *OpenClusterPasswordComplexityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenClusterPasswordComplexityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenClusterReadOnlyInstanceGroupAccessRequestParams struct {

}

type OpenClusterReadOnlyInstanceGroupAccessRequest struct {
	*tchttp.BaseRequest
	
}

func (r *OpenClusterReadOnlyInstanceGroupAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenClusterReadOnlyInstanceGroupAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenClusterReadOnlyInstanceGroupAccessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenClusterReadOnlyInstanceGroupAccessResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OpenClusterReadOnlyInstanceGroupAccessResponse struct {
	*tchttp.BaseResponse
	Response *OpenClusterReadOnlyInstanceGroupAccessResponseParams `json:"Response"`
}

func (r *OpenClusterReadOnlyInstanceGroupAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenClusterReadOnlyInstanceGroupAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenReadOnlyInstanceExclusiveAccessRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// ID of the read-only instance with dedicated access to be enabled
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Specified VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Specified subnet ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Port
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Security group
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

type OpenReadOnlyInstanceExclusiveAccessRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// ID of the read-only instance with dedicated access to be enabled
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Specified VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Specified subnet ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Port
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Security group
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

func (r *OpenReadOnlyInstanceExclusiveAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenReadOnlyInstanceExclusiveAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "Port")
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenReadOnlyInstanceExclusiveAccessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenReadOnlyInstanceExclusiveAccessResponseParams struct {
	// Activation process ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OpenReadOnlyInstanceExclusiveAccessResponse struct {
	*tchttp.BaseResponse
	Response *OpenReadOnlyInstanceExclusiveAccessResponseParams `json:"Response"`
}

func (r *OpenReadOnlyInstanceExclusiveAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenReadOnlyInstanceExclusiveAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenWanRequestParams struct {
	// Instance group ID
	InstanceGrpId *string `json:"InstanceGrpId,omitnil,omitempty" name:"InstanceGrpId"`
}

type OpenWanRequest struct {
	*tchttp.BaseRequest
	
	// Instance group ID
	InstanceGrpId *string `json:"InstanceGrpId,omitnil,omitempty" name:"InstanceGrpId"`
}

func (r *OpenWanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenWanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceGrpId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenWanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenWanResponseParams struct {
	// Task flow ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OpenWanResponse struct {
	*tchttp.BaseResponse
	Response *OpenWanResponseParams `json:"Response"`
}

func (r *OpenWanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenWanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Package struct {
	// AppID Note: This field may return null, indicating that no valid values can be obtained.
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// The unique ID of a resource pack Note: This field may return null, indicating that no valid values can be obtained.
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// Resource pack name Note: This field may return null, indicating that no valid values can be obtained.
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// Resource pack type. Valid values: `CCU` (compute resource pack), `DISK` (storage resource pack). Note: This field may return null, indicating that no valid values can be obtained.
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// Region of the resource pack. Valid values: `China` (Chinese mainland), `overseas` (outside Chinese mainland). Note: This field may return null, indicating that no valid values can be obtained.
	PackageRegion *string `json:"PackageRegion,omitnil,omitempty" name:"PackageRegion"`

	// Resource pack status. Valid values: `creating`, `using`, `expired`, `normal_finish` (used up), `apply_refund` (requesting a refund), `refund` (refunded). 
	// Note:  This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Total number of resource packs Note: This field may return null, indicating that no valid values can be obtained.
	PackageTotalSpec *float64 `json:"PackageTotalSpec,omitnil,omitempty" name:"PackageTotalSpec"`

	// Consumed usage of resource packs Note: This field may return null, indicating that no valid values can be obtained.
	PackageUsedSpec *float64 `json:"PackageUsedSpec,omitnil,omitempty" name:"PackageUsedSpec"`

	// Remaining usage of resource packs Note: This field may return null, indicating that no valid values can be obtained.
	HasQuota *bool `json:"HasQuota,omitnil,omitempty" name:"HasQuota"`

	// Information of the instance bound Note: This field may return null, indicating that no valid values can be obtained.
	BindInstanceInfos []*BindInstanceInfo `json:"BindInstanceInfos,omitnil,omitempty" name:"BindInstanceInfos"`

	// Validity time:  2022-07-01 00:00:00 Note: This field may return null, indicating that no valid values can be obtained.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Validity time:  2022-08-01 00:00:00 Note: This field may return null, indicating that no valid values can be obtained.
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`
}

type PackageDetail struct {
	// Account ID of `AppId` Note: This field may return null, indicating that no valid values can be obtained.
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// The unique ID of a resource pack Note: This field may return null, indicating that no valid values can be obtained.
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// Instance ID Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The successfully deducted capacity Note: This field may return null, indicating that no valid values can be obtained.
	SuccessDeductSpec *float64 `json:"SuccessDeductSpec,omitnil,omitempty" name:"SuccessDeductSpec"`

	// Used capacity of a resource pack as of now Note: This field may return null, indicating that no valid values can be obtained.
	PackageTotalUsedSpec *float64 `json:"PackageTotalUsedSpec,omitnil,omitempty" name:"PackageTotalUsedSpec"`

	// Deduction start time Note: This field may return null, indicating that no valid values can be obtained.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Deduction end time Note: This field may return null, indicating that no valid values can be obtained.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Extended information Note: This field may return null, indicating that no valid values can be obtained.
	ExtendInfo *string `json:"ExtendInfo,omitnil,omitempty" name:"ExtendInfo"`
}

type ParamDetail struct {
	// Parameter name
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// Parameter type. Valid values:  `integer`, `enum`, `float`, `string`, `func`.
	ParamType *string `json:"ParamType,omitnil,omitempty" name:"ParamType"`

	// Whether `func` is supported. Valid values: `true` (supported), `false` (not supported).
	SupportFunc *bool `json:"SupportFunc,omitnil,omitempty" name:"SupportFunc"`

	// Default value
	Default *string `json:"Default,omitnil,omitempty" name:"Default"`

	// Parameter description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Current value of the parameter
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`

	// Whether to restart the database for the modified parameters to take effect. Valid values:  `0` (no), `1` (yes).
	NeedReboot *int64 `json:"NeedReboot,omitnil,omitempty" name:"NeedReboot"`

	// Maximum value of the parameter
	Max *string `json:"Max,omitnil,omitempty" name:"Max"`

	// Minimum value of the parameter
	Min *string `json:"Min,omitnil,omitempty" name:"Min"`

	// Enumerated values of the parameter.  It is null if the parameter is non-enumerated. Note: This field may return null, indicating that no valid values can be obtained.
	EnumValue []*string `json:"EnumValue,omitnil,omitempty" name:"EnumValue"`

	// Valid values: `1` (global parameter),  `0`  (non-global parameter).
	IsGlobal *int64 `json:"IsGlobal,omitnil,omitempty" name:"IsGlobal"`

	// The match type. Valid value: `multiVal`.
	MatchType *string `json:"MatchType,omitnil,omitempty" name:"MatchType"`

	// Match values, which will be separated by comma when `MatchType` is `multiVal`.
	MatchValue *string `json:"MatchValue,omitnil,omitempty" name:"MatchValue"`

	// Whether it is a `func` type. Valid values: `true` (yes), `false` (no). Note: This field may return null, indicating that no valid values can be obtained.
	IsFunc *bool `json:"IsFunc,omitnil,omitempty" name:"IsFunc"`

	// Formula content returned when `ParamType` is `func`. Note: This field may return null, indicating that no valid values can be obtained.
	Func *string `json:"Func,omitnil,omitempty" name:"Func"`

	// Whether the parameter can be modified Note: This field may return null, indicating that no valid values can be obtained.
	ModifiableInfo *ModifiableInfo `json:"ModifiableInfo,omitnil,omitempty" name:"ModifiableInfo"`
}

type ParamInfo struct {
	// Current value
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`

	// Default value
	Default *string `json:"Default,omitnil,omitempty" name:"Default"`

	// List of valid values when parameter type is `enum`, `string` or `bool`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EnumValue []*string `json:"EnumValue,omitnil,omitempty" name:"EnumValue"`

	// Maximum value when parameter type is `float` or `integer`.
	Max *string `json:"Max,omitnil,omitempty" name:"Max"`

	// Minimum value when parameter type is `float` or `integer`.
	Min *string `json:"Min,omitnil,omitempty" name:"Min"`

	// Parameter name
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// Whether to restart the instance for the modified parameters to take effect.
	NeedReboot *int64 `json:"NeedReboot,omitnil,omitempty" name:"NeedReboot"`

	// Parameter type: `integer`, `float`, `string`, `enum`, `bool`.
	ParamType *string `json:"ParamType,omitnil,omitempty" name:"ParamType"`

	// Match type. Regex can be used when parameter type is `string`. Valid value: `multiVal`.
	MatchType *string `json:"MatchType,omitnil,omitempty" name:"MatchType"`

	// Match values, which will be separated by semicolon when match type is `multiVal`.
	MatchValue *string `json:"MatchValue,omitnil,omitempty" name:"MatchValue"`

	// Parameter description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Whether it is global parameter
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsGlobal *int64 `json:"IsGlobal,omitnil,omitempty" name:"IsGlobal"`

	// Whether the parameter can be modified
	// Note: This field may return null, indicating that no valid values can be obtained.
	ModifiableInfo *ModifiableInfo `json:"ModifiableInfo,omitnil,omitempty" name:"ModifiableInfo"`

	// Whether it is a function
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsFunc *bool `json:"IsFunc,omitnil,omitempty" name:"IsFunc"`

	// Function
	// Note: This field may return null, indicating that no valid values can be obtained.
	Func *string `json:"Func,omitnil,omitempty" name:"Func"`
}

type ParamItem struct {
	// Parameter name
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// New value
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`

	// Original value
	OldValue *string `json:"OldValue,omitnil,omitempty" name:"OldValue"`
}

type ParamItemDetail struct {
	// Current value
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`

	// Default value
	Default *string `json:"Default,omitnil,omitempty" name:"Default"`

	// Enumerated values of the parameter It is null if the parameter is non-enumerated.
	EnumValue []*string `json:"EnumValue,omitnil,omitempty" name:"EnumValue"`

	// Valid values: `1` (global parameter),  `0`  (non-global parameter).
	IsGlobal *int64 `json:"IsGlobal,omitnil,omitempty" name:"IsGlobal"`

	// Maximum value
	Max *string `json:"Max,omitnil,omitempty" name:"Max"`

	// Minimum value
	Min *string `json:"Min,omitnil,omitempty" name:"Min"`

	// Whether to restart the database for the modified parameters to take effect. Valid values:  `0` (no), `1` (yes)
	NeedReboot *int64 `json:"NeedReboot,omitnil,omitempty" name:"NeedReboot"`

	// Parameter name
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// Parameter type. Valid values:  `integer`, `enum`, `float`, `string`, `func`.
	ParamType *string `json:"ParamType,omitnil,omitempty" name:"ParamType"`

	// Parameter description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Whether `ParamType` is a `func` Note: This field may return null, indicating that no valid values can be obtained.
	IsFunc *bool `json:"IsFunc,omitnil,omitempty" name:"IsFunc"`

	// Parameter configuration formula Note: This field may return null, indicating that no valid values can be obtained.
	Func *string `json:"Func,omitnil,omitempty" name:"Func"`
}

type ParamTemplateListInfo struct {
	// Parameter template ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Parameter template name
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// Parameter template description
	TemplateDescription *string `json:"TemplateDescription,omitnil,omitempty" name:"TemplateDescription"`

	// Engine version
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// Database Type. Valid values: `NORMAL`, `SERVERLESS`.
	DbMode *string `json:"DbMode,omitnil,omitempty" name:"DbMode"`

	// Parameter template details
	// Note: This field may return null, indicating that no valid values can be obtained.
	ParamInfoSet []*TemplateParamInfo `json:"ParamInfoSet,omitnil,omitempty" name:"ParamInfoSet"`
}

// Predefined struct for user
type PauseServerlessRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Whether to pause forcibly and ignore the current user connections. Valid values: `0` (no), `1` (yes). Default value: `1`
	ForcePause *int64 `json:"ForcePause,omitnil,omitempty" name:"ForcePause"`
}

type PauseServerlessRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Whether to pause forcibly and ignore the current user connections. Valid values: `0` (no), `1` (yes). Default value: `1`
	ForcePause *int64 `json:"ForcePause,omitnil,omitempty" name:"ForcePause"`
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
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// Source IP or source IP range, such as 192.168.0.0/16
	CidrIp *string `json:"CidrIp,omitnil,omitempty" name:"CidrIp"`

	// Port
	PortRange *string `json:"PortRange,omitnil,omitempty" name:"PortRange"`

	// Network protocol, such as UDP and TCP
	IpProtocol *string `json:"IpProtocol,omitnil,omitempty" name:"IpProtocol"`

	// Protocol port ID or protocol port group ID.
	ServiceModule *string `json:"ServiceModule,omitnil,omitempty" name:"ServiceModule"`

	// IP address ID or IP address group ID.
	AddressModule *string `json:"AddressModule,omitnil,omitempty" name:"AddressModule"`

	// id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Description
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`
}

type ProxyConnectionPoolInfo struct {
	// Connection persistence timeout in seconds
	// Note: This field may return null, indicating that no valid values can be obtained.
	ConnectionPoolTimeOut *int64 `json:"ConnectionPoolTimeOut,omitnil,omitempty" name:"ConnectionPoolTimeOut"`

	// Whether the connection pool is enabled
	// Note: This field may return null, indicating that no valid values can be obtained.
	OpenConnectionPool *string `json:"OpenConnectionPool,omitnil,omitempty" name:"OpenConnectionPool"`

	// Connection pool type. Valid value: `SessionConnectionPool` (session-level).
	// Note: This field may return null, indicating that no valid values can be obtained.
	ConnectionPoolType *string `json:"ConnectionPoolType,omitnil,omitempty" name:"ConnectionPoolType"`
}

type ProxyGroup struct {
	// Database proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Number of nodes in the proxy group
	ProxyNodeCount *int64 `json:"ProxyNodeCount,omitnil,omitempty" name:"ProxyNodeCount"`

	// Database proxy group status
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Region
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// AZ
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Current proxy version
	CurrentProxyVersion *string `json:"CurrentProxyVersion,omitnil,omitempty" name:"CurrentProxyVersion"`

	// Cluster ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// User `AppId`
	// Note: This field may return null, indicating that no valid values can be obtained.
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Enabling read/write separation for database proxy
	// Note: This field may return null, indicating that no valid values can be obtained.
	OpenRw *string `json:"OpenRw,omitnil,omitempty" name:"OpenRw"`
}

type ProxyGroupInfo struct {
	// Database proxy group
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProxyGroup *ProxyGroup `json:"ProxyGroup,omitnil,omitempty" name:"ProxyGroup"`

	// Read/write separation information of the database proxy group
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProxyGroupRwInfo *ProxyGroupRwInfo `json:"ProxyGroupRwInfo,omitnil,omitempty" name:"ProxyGroupRwInfo"`

	// Node information of the database proxy
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProxyNodes []*ProxyNodeInfo `json:"ProxyNodes,omitnil,omitempty" name:"ProxyNodes"`

	// Connection pool information for the database proxy
	// Note: This field may return null, indicating that no valid values can be obtained.
	ConnectionPool *ProxyConnectionPoolInfo `json:"ConnectionPool,omitnil,omitempty" name:"ConnectionPool"`

	// Network information for database proxy
	// Note: This field may return null, indicating that no valid values can be obtained.
	NetAddrInfos []*NetAddr `json:"NetAddrInfos,omitnil,omitempty" name:"NetAddrInfos"`

	// Task set of the database proxy
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tasks []*ObjectTask `json:"Tasks,omitnil,omitempty" name:"Tasks"`
}

type ProxyGroupRwInfo struct {
	// Consistency type. Valid values: `eventual` (eventual consistency), `session` (session consistency), `global` (global consistency).
	ConsistencyType *string `json:"ConsistencyType,omitnil,omitempty" name:"ConsistencyType"`

	// Consistency timeout period
	ConsistencyTimeOut *int64 `json:"ConsistencyTimeOut,omitnil,omitempty" name:"ConsistencyTimeOut"`

	// Weight mode. Valid values: `system` (auto-assigned), `custom`.
	WeightMode *string `json:"WeightMode,omitnil,omitempty" name:"WeightMode"`

	// Whether to enable failover
	FailOver *string `json:"FailOver,omitnil,omitempty" name:"FailOver"`

	// Whether to automatically add read-only instance. Valid value: `yes`, `no`.
	AutoAddRo *string `json:"AutoAddRo,omitnil,omitempty" name:"AutoAddRo"`

	// Instance weight array
	InstanceWeights []*ProxyInstanceWeight `json:"InstanceWeights,omitnil,omitempty" name:"InstanceWeights"`

	// Whether to enable read-write node. Valid values: `yes`, `no`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OpenRw *string `json:"OpenRw,omitnil,omitempty" name:"OpenRw"`

	// Read/write attribute. Valid values: `READWRITE`, `READONLY`.
	RwType *string `json:"RwType,omitnil,omitempty" name:"RwType"`

	// Transaction split
	TransSplit *bool `json:"TransSplit,omitnil,omitempty" name:"TransSplit"`

	// Connection mode. Valid values: `balance`, `nearby`.
	AccessMode *string `json:"AccessMode,omitnil,omitempty" name:"AccessMode"`
}

type ProxyInstanceWeight struct {
	// InstanID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance weight
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`
}

type ProxyNodeInfo struct {
	// Database proxy node ID
	ProxyNodeId *string `json:"ProxyNodeId,omitnil,omitempty" name:"ProxyNodeId"`

	// Current node connections, which is not returned by the `DescribeProxyNodes` API.
	ProxyNodeConnections *int64 `json:"ProxyNodeConnections,omitnil,omitempty" name:"ProxyNodeConnections"`

	// CPU of the database proxy node
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Memory of the database proxy node
	Mem *int64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// Status of the database proxy node
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Database proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// User AppID
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Region
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// AZ
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

type ProxySpec struct {
	// Number of database proxy CPU cores
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Database proxy memory
	Mem *int64 `json:"Mem,omitnil,omitempty" name:"Mem"`
}

type ProxyZone struct {
	// AZ of the proxy node
	ProxyNodeZone *string `json:"ProxyNodeZone,omitnil,omitempty" name:"ProxyNodeZone"`

	// The number of proxy nodes
	ProxyNodeCount *int64 `json:"ProxyNodeCount,omitnil,omitempty" name:"ProxyNodeCount"`
}

type QueryFilter struct {
	// Search field. Valid values: "InstanceId", "ProjectId", "InstanceName", "Vip"
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`

	// Search string
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// Whether to use exact match
	ExactMatch *bool `json:"ExactMatch,omitnil,omitempty" name:"ExactMatch"`

	// Search field
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Operator
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`
}

type QueryParamFilter struct {
	// Search field. Valid values: "InstanceId", "ProjectId", "InstanceName", "Vip"
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`

	// Search string
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// Whether to use exact match
	ExactMatch *bool `json:"ExactMatch,omitnil,omitempty" name:"ExactMatch"`
}

// Predefined struct for user
type RefundResourcePackageRequestParams struct {
	// The unique ID of a resource pack
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`
}

type RefundResourcePackageRequest struct {
	*tchttp.BaseRequest
	
	// The unique ID of a resource pack
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`
}

func (r *RefundResourcePackageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefundResourcePackageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PackageId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RefundResourcePackageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RefundResourcePackageResponseParams struct {
	// Each item has only one `dealName`, through which you can ensure the idempotency of the delivery API.
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RefundResourcePackageResponse struct {
	*tchttp.BaseResponse
	Response *RefundResourcePackageResponseParams `json:"Response"`
}

func (r *RefundResourcePackageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefundResourcePackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReloadBalanceProxyNodeRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`
}

type ReloadBalanceProxyNodeRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`
}

func (r *ReloadBalanceProxyNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReloadBalanceProxyNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ProxyGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReloadBalanceProxyNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReloadBalanceProxyNodeResponseParams struct {
	// Async flow ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Async task ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReloadBalanceProxyNodeResponse struct {
	*tchttp.BaseResponse
	Response *ReloadBalanceProxyNodeResponseParams `json:"Response"`
}

func (r *ReloadBalanceProxyNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReloadBalanceProxyNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveClusterSlaveZoneRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Replica AZ
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`
}

type RemoveClusterSlaveZoneRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Replica AZ
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`
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
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// New password of the database account
	AccountPassword *string `json:"AccountPassword,omitnil,omitempty" name:"AccountPassword"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Host. Default value: `%`
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

type ResetAccountPasswordRequest struct {
	*tchttp.BaseRequest
	
	// Database account name
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// New password of the database account
	AccountPassword *string `json:"AccountPassword,omitnil,omitempty" name:"AccountPassword"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Host. Default value: `%`
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
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
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type ResourcePackage struct {
	// The unique ID of a resource pack Note: This field may return null, indicating that no valid values can be obtained.
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// Resource pack type. Valid values:  `CCU` (compute resource pack),  `DISK`  (storage resource pack). Note: This field may return null, indicating that no valid values can be obtained.
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`
}

// Predefined struct for user
type RestartInstanceRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type RestartInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type ResumeServerlessRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
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
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	TimeRangeStart *string `json:"TimeRangeStart,omitnil,omitempty" name:"TimeRangeStart"`

	// End time
	TimeRangeEnd *string `json:"TimeRangeEnd,omitnil,omitempty" name:"TimeRangeEnd"`
}

type RuleFilters struct {
	// Filter parameter name of the audit rule. Valid values: `host` (client IP), `user` (database account), `dbName` (database name), `sqlType` (SQL type), `sql` (SQL statement).
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Filter match type of the audit rule. Valid values: `INC` (including), `EXC` (excluding), `EQS` (equal to), `NEQ` (not equal to).
	Compare *string `json:"Compare,omitnil,omitempty" name:"Compare"`

	// Filter match value of the audit rule
	Value []*string `json:"Value,omitnil,omitempty" name:"Value"`
}

type SalePackageSpec struct {
	// Region of the resource pack Note: This field may return null, indicating that no valid values can be obtained.
	PackageRegion *string `json:"PackageRegion,omitnil,omitempty" name:"PackageRegion"`

	// Resource pack type. Valid values: `CCU` (compute resource pack), `DISK` (storage resource pack). Note: This field may return null, indicating that no valid values can be obtained.
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// Resource pack edition. Valid values: `base` (basic edition), `common` (general edition), `enterprise` (enterprise edition). Note: This field may return null, indicating that no valid values can be obtained.
	PackageVersion *string `json:"PackageVersion,omitnil,omitempty" name:"PackageVersion"`

	// Minimum number of resources for the current edition of the resource pack.  Unit of the compute resources: pcs.  Unit of the storage resources: GB. Note: This field may return null, indicating that no valid values can be obtained.
	MinPackageSpec *float64 `json:"MinPackageSpec,omitnil,omitempty" name:"MinPackageSpec"`

	// Maximum number of resources for the current edition of the resource pack.  Unit of the compute resources: pcs.  Unit of the storage resources: GB. Note: This field may return null, indicating that no valid values can be obtained.
	MaxPackageSpec *float64 `json:"MaxPackageSpec,omitnil,omitempty" name:"MaxPackageSpec"`

	// Validity period of a resource pack in days Note: This field may return null, indicating that no valid values can be obtained.
	ExpireDay *int64 `json:"ExpireDay,omitnil,omitempty" name:"ExpireDay"`
}

type SaleRegion struct {
	// Region name
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Numeric ID of a region
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Region name
	RegionZh *string `json:"RegionZh,omitnil,omitempty" name:"RegionZh"`

	// List of purchasable AZs
	ZoneSet []*SaleZone `json:"ZoneSet,omitnil,omitempty" name:"ZoneSet"`

	// Engine type
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Supported modules in a region
	Modules []*Module `json:"Modules,omitnil,omitempty" name:"Modules"`
}

type SaleZone struct {
	// AZ name
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Numeric ID of an AZ
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// AZ name
	ZoneZh *string `json:"ZoneZh,omitnil,omitempty" name:"ZoneZh"`

	// Whether serverless cluster is supported. Valid values: <br>
	// `0`: No<br>
	// `1`: Yes
	IsSupportServerless *int64 `json:"IsSupportServerless,omitnil,omitempty" name:"IsSupportServerless"`

	// Whether standard cluster is supported. Valid values: <br>
	// `0`: No<br>
	// `1`: Yes
	IsSupportNormal *int64 `json:"IsSupportNormal,omitnil,omitempty" name:"IsSupportNormal"`

	// Physical zone
	PhysicalZone *string `json:"PhysicalZone,omitnil,omitempty" name:"PhysicalZone"`

	// Whether the user has AZ permission
	// Note: This field may return null, indicating that no valid values can be obtained.
	HasPermission *bool `json:"HasPermission,omitnil,omitempty" name:"HasPermission"`

	// Whether it is a full-linkage RDMA AZ.
	IsWholeRdmaZone *string `json:"IsWholeRdmaZone,omitnil,omitempty" name:"IsWholeRdmaZone"`
}

// Predefined struct for user
type SearchClusterDatabasesRequestParams struct {
	// The cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database name
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// Whether to search exactly
	// Valid values: `0` (fuzzy search), `1` (exact search). 
	// Default value: `0`.
	MatchType *int64 `json:"MatchType,omitnil,omitempty" name:"MatchType"`
}

type SearchClusterDatabasesRequest struct {
	*tchttp.BaseRequest
	
	// The cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database name
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// Whether to search exactly
	// Valid values: `0` (fuzzy search), `1` (exact search). 
	// Default value: `0`.
	MatchType *int64 `json:"MatchType,omitnil,omitempty" name:"MatchType"`
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
	Databases []*string `json:"Databases,omitnil,omitempty" name:"Databases"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database name
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// Data table name
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// Data table type. Valid values:
	// `view`: Only return to view,
	// `base_table`: Only return to basic table,
	// `all`: Return to view and table.
	TableType *string `json:"TableType,omitnil,omitempty" name:"TableType"`
}

type SearchClusterTablesRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database name
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// Data table name
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// Data table type. Valid values:
	// `view`: Only return to view,
	// `base_table`: Only return to basic table,
	// `all`: Return to view and table.
	TableType *string `json:"TableType,omitnil,omitempty" name:"TableType"`
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
	Tables []*DatabaseTables `json:"Tables,omitnil,omitempty" name:"Tables"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Creation time in the format of yyyy-mm-dd hh:mm:ss
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Inbound rule
	Inbound []*PolicyRule `json:"Inbound,omitnil,omitempty" name:"Inbound"`

	// Outbound rule
	Outbound []*PolicyRule `json:"Outbound,omitnil,omitempty" name:"Outbound"`

	// Security group ID
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// Security group name
	SecurityGroupName *string `json:"SecurityGroupName,omitnil,omitempty" name:"SecurityGroupName"`

	// Security group remarks
	SecurityGroupRemark *string `json:"SecurityGroupRemark,omitnil,omitempty" name:"SecurityGroupRemark"`
}

// Predefined struct for user
type SetRenewFlagRequestParams struct {
	// ID of the instance to be manipulated
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// Auto-renewal flag. 0: normal renewal, 1: auto-renewal, 2: no renewal.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`
}

type SetRenewFlagRequest struct {
	*tchttp.BaseRequest
	
	// ID of the instance to be manipulated
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// Auto-renewal flag. 0: normal renewal, 1: auto-renewal, 2: no renewal.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`
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
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type SlaveZoneAttrItem struct {

	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`


	BinlogSyncWay *string `json:"BinlogSyncWay,omitnil,omitempty" name:"BinlogSyncWay"`
}

type SlowQueriesItem struct {
	// Execution timestamp
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// Execution duration in seconds
	QueryTime *float64 `json:"QueryTime,omitnil,omitempty" name:"QueryTime"`

	// SQL statement
	SqlText *string `json:"SqlText,omitnil,omitempty" name:"SqlText"`

	// Client host
	UserHost *string `json:"UserHost,omitnil,omitempty" name:"UserHost"`

	// Username
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Database name
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// Lock duration in seconds
	LockTime *float64 `json:"LockTime,omitnil,omitempty" name:"LockTime"`

	// Number of scanned rows
	RowsExamined *int64 `json:"RowsExamined,omitnil,omitempty" name:"RowsExamined"`

	// Number of returned rows
	RowsSent *int64 `json:"RowsSent,omitnil,omitempty" name:"RowsSent"`

	// SQL template
	SqlTemplate *string `json:"SqlTemplate,omitnil,omitempty" name:"SqlTemplate"`

	// MD5 value of the SQL statement
	SqlMd5 *string `json:"SqlMd5,omitnil,omitempty" name:"SqlMd5"`
}

// Predefined struct for user
type SwitchClusterVpcRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// VPC ID in string
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// Subnet ID in string
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// Valid hours of old IP
	OldIpReserveHours *int64 `json:"OldIpReserveHours,omitnil,omitempty" name:"OldIpReserveHours"`
}

type SwitchClusterVpcRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// VPC ID in string
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// Subnet ID in string
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// Valid hours of old IP
	OldIpReserveHours *int64 `json:"OldIpReserveHours,omitnil,omitempty" name:"OldIpReserveHours"`
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
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// The current AZ
	OldZone *string `json:"OldZone,omitnil,omitempty" name:"OldZone"`

	// New AZ
	NewZone *string `json:"NewZone,omitnil,omitempty" name:"NewZone"`

	// Valid values: `yes` (execute during maintenance time), `no` (execute now)
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
}

type SwitchClusterZoneRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// The current AZ
	OldZone *string `json:"OldZone,omitnil,omitempty" name:"OldZone"`

	// New AZ
	NewZone *string `json:"NewZone,omitnil,omitempty" name:"NewZone"`

	// Valid values: `yes` (execute during maintenance time), `no` (execute now)
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
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
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// VPC ID in string
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// Subnet ID in string
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// Valid hours of old IP
	OldIpReserveHours *int64 `json:"OldIpReserveHours,omitnil,omitempty" name:"OldIpReserveHours"`

	// Database proxy group ID (required), which can be obtained through the `DescribeProxies` API.
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`
}

type SwitchProxyVpcRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// VPC ID in string
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// Subnet ID in string
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// Valid hours of old IP
	OldIpReserveHours *int64 `json:"OldIpReserveHours,omitnil,omitempty" name:"OldIpReserveHours"`

	// Database proxy group ID (required), which can be obtained through the `DescribeProxies` API.
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`
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
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type TablePrivileges struct {
	// Database name
	Db *string `json:"Db,omitnil,omitempty" name:"Db"`

	// Table name
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// Permission list
	Privileges []*string `json:"Privileges,omitnil,omitempty" name:"Privileges"`
}

type Tag struct {
	// Tag key
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TemplateParamInfo struct {
	// Current value
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`

	// Default value
	Default *string `json:"Default,omitnil,omitempty" name:"Default"`

	// The collection of valid value types when parameter type is `enum`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EnumValue []*string `json:"EnumValue,omitnil,omitempty" name:"EnumValue"`

	// Maximum value when parameter type is `float` or `integer`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Max *string `json:"Max,omitnil,omitempty" name:"Max"`

	// Minimum value when parameter type is `float` or `integer`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Min *string `json:"Min,omitnil,omitempty" name:"Min"`

	// Parameter name
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// Whether to restart the instance for the parameter to take effect
	NeedReboot *int64 `json:"NeedReboot,omitnil,omitempty" name:"NeedReboot"`

	// Parameter description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Parameter type. Valid value: `integer`, `float`, `string`, `enum`.
	ParamType *string `json:"ParamType,omitnil,omitempty" name:"ParamType"`
}

type TradePrice struct {
	// The non-discounted total price of monthly subscribed resources (unit: 0.000001 cent)
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalPrice *int64 `json:"TotalPrice,omitnil,omitempty" name:"TotalPrice"`

	// Total discount. `100` means no discount.
	Discount *float64 `json:"Discount,omitnil,omitempty" name:"Discount"`

	// The discounted total price of monthly subscribed resources (unit: 0.000001 cent). If a discount is applied, `TotalPriceDiscount` will be the product of `TotalPrice` and `Discount`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalPriceDiscount *int64 `json:"TotalPriceDiscount,omitnil,omitempty" name:"TotalPriceDiscount"`

	// The non-discounted unit price of pay-as-you-go resources (unit: 0.000001 cent)
	// Note: This field may return null, indicating that no valid values can be obtained.
	UnitPrice *int64 `json:"UnitPrice,omitnil,omitempty" name:"UnitPrice"`

	// The discounted unit price of pay-as-you-go resources (unit: 0.000001 cent). If a discount is applied, `UnitPriceDiscount` will be the product of `UnitPrice` and `Discount`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UnitPriceDiscount *int64 `json:"UnitPriceDiscount,omitnil,omitempty" name:"UnitPriceDiscount"`

	// Price unit
	ChargeUnit *string `json:"ChargeUnit,omitnil,omitempty" name:"ChargeUnit"`
}

// Predefined struct for user
type UnbindClusterResourcePackagesRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// The unique ID of a resource pack. If this parameter is left empty, all resource packs of the instance will be unbound.
	PackageIds []*string `json:"PackageIds,omitnil,omitempty" name:"PackageIds"`
}

type UnbindClusterResourcePackagesRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// The unique ID of a resource pack. If this parameter is left empty, all resource packs of the instance will be unbound.
	PackageIds []*string `json:"PackageIds,omitnil,omitempty" name:"PackageIds"`
}

func (r *UnbindClusterResourcePackagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindClusterResourcePackagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "PackageIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindClusterResourcePackagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindClusterResourcePackagesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UnbindClusterResourcePackagesResponse struct {
	*tchttp.BaseResponse
	Response *UnbindClusterResourcePackagesResponseParams `json:"Response"`
}

func (r *UnbindClusterResourcePackagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindClusterResourcePackagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeClusterVersionRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Kernel version
	CynosVersion *string `json:"CynosVersion,omitnil,omitempty" name:"CynosVersion"`

	// Upgrade time type. Valid values: `upgradeImmediate`, `upgradeInMaintain`.
	UpgradeType *string `json:"UpgradeType,omitnil,omitempty" name:"UpgradeType"`
}

type UpgradeClusterVersionRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Kernel version
	CynosVersion *string `json:"CynosVersion,omitnil,omitempty" name:"CynosVersion"`

	// Upgrade time type. Valid values: `upgradeImmediate`, `upgradeInMaintain`.
	UpgradeType *string `json:"UpgradeType,omitnil,omitempty" name:"UpgradeType"`
}

func (r *UpgradeClusterVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeClusterVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "CynosVersion")
	delete(f, "UpgradeType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeClusterVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeClusterVersionResponseParams struct {
	// Async task ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeClusterVersionResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeClusterVersionResponseParams `json:"Response"`
}

func (r *UpgradeClusterVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeClusterVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeInstanceRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Database CPU
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Database memory in GB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Upgrade type. Valid values: upgradeImmediate, upgradeInMaintain
	UpgradeType *string `json:"UpgradeType,omitnil,omitempty" name:"UpgradeType"`

	// This parameter has been disused.
	StorageLimit *uint64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`

	// Whether to automatically select a voucher. 1: yes; 0: no. Default value: 0
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// This parameter has been disused.
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Transaction mode. Valid values: `0` (place and pay for an order), `1` (place an order)
	DealMode *int64 `json:"DealMode,omitnil,omitempty" name:"DealMode"`

	// Valid values: `NormalUpgrade` (Normal mode), `FastUpgrade` (QuickChange). If the system detects that the configuration modification process will cause a momentary disconnection, the process will be terminated.
	UpgradeMode *string `json:"UpgradeMode,omitnil,omitempty" name:"UpgradeMode"`
}

type UpgradeInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Database CPU
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Database memory in GB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Upgrade type. Valid values: upgradeImmediate, upgradeInMaintain
	UpgradeType *string `json:"UpgradeType,omitnil,omitempty" name:"UpgradeType"`

	// This parameter has been disused.
	StorageLimit *uint64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`

	// Whether to automatically select a voucher. 1: yes; 0: no. Default value: 0
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// This parameter has been disused.
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Transaction mode. Valid values: `0` (place and pay for an order), `1` (place an order)
	DealMode *int64 `json:"DealMode,omitnil,omitempty" name:"DealMode"`

	// Valid values: `NormalUpgrade` (Normal mode), `FastUpgrade` (QuickChange). If the system detects that the configuration modification process will cause a momentary disconnection, the process will be terminated.
	UpgradeMode *string `json:"UpgradeMode,omitnil,omitempty" name:"UpgradeMode"`
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
	TranId *string `json:"TranId,omitnil,omitempty" name:"TranId"`

	// Big order ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	BigDealIds []*string `json:"BigDealIds,omitnil,omitempty" name:"BigDealIds"`

	// Order ID
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type UpgradeProxyRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Number of CPU cores
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Memory
	Mem *int64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// Number of nodes in the proxy group
	ProxyCount *int64 `json:"ProxyCount,omitnil,omitempty" name:"ProxyCount"`

	// ID of the database proxy group (disused)
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Load rebalance mode. Valid values: `auto`, `manual`
	ReloadBalance *string `json:"ReloadBalance,omitnil,omitempty" name:"ReloadBalance"`

	// Upgrade time. Valid values: `no` (upon upgrade completion), `timeWindow` (upgrade during instance maintenance time)
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`

	// Node information of the atabase proxy
	ProxyZones []*ProxyZone `json:"ProxyZones,omitnil,omitempty" name:"ProxyZones"`
}

type UpgradeProxyRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Number of CPU cores
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Memory
	Mem *int64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// Number of nodes in the proxy group
	ProxyCount *int64 `json:"ProxyCount,omitnil,omitempty" name:"ProxyCount"`

	// ID of the database proxy group (disused)
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Load rebalance mode. Valid values: `auto`, `manual`
	ReloadBalance *string `json:"ReloadBalance,omitnil,omitempty" name:"ReloadBalance"`

	// Upgrade time. Valid values: `no` (upon upgrade completion), `timeWindow` (upgrade during instance maintenance time)
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`

	// Node information of the atabase proxy
	ProxyZones []*ProxyZone `json:"ProxyZones,omitnil,omitempty" name:"ProxyZones"`
}

func (r *UpgradeProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Cpu")
	delete(f, "Mem")
	delete(f, "ProxyCount")
	delete(f, "ProxyGroupId")
	delete(f, "ReloadBalance")
	delete(f, "IsInMaintainPeriod")
	delete(f, "ProxyZones")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeProxyResponseParams struct {
	// Async flow ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Async task ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeProxyResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeProxyResponseParams `json:"Response"`
}

func (r *UpgradeProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeProxyVersionRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Current version of database proxy
	SrcProxyVersion *string `json:"SrcProxyVersion,omitnil,omitempty" name:"SrcProxyVersion"`

	// Target version of database proxy
	DstProxyVersion *string `json:"DstProxyVersion,omitnil,omitempty" name:"DstProxyVersion"`

	// Database proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Upgrade time. Valid values: `no` (upon upgrade completion), `yes` (upgrade during instance maintenance time)
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
}

type UpgradeProxyVersionRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Current version of database proxy
	SrcProxyVersion *string `json:"SrcProxyVersion,omitnil,omitempty" name:"SrcProxyVersion"`

	// Target version of database proxy
	DstProxyVersion *string `json:"DstProxyVersion,omitnil,omitempty" name:"DstProxyVersion"`

	// Database proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Upgrade time. Valid values: `no` (upon upgrade completion), `yes` (upgrade during instance maintenance time)
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
}

func (r *UpgradeProxyVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeProxyVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SrcProxyVersion")
	delete(f, "DstProxyVersion")
	delete(f, "ProxyGroupId")
	delete(f, "IsInMaintainPeriod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeProxyVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeProxyVersionResponseParams struct {
	// Async flow ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Async task ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeProxyVersionResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeProxyVersionResponseParams `json:"Response"`
}

func (r *UpgradeProxyVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeProxyVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserHostPrivilege struct {
	// Authorized user
	DbUserName *string `json:"DbUserName,omitnil,omitempty" name:"DbUserName"`

	// Client IP Note: This field may return null, indicating that no valid values can be obtained.
	DbHost *string `json:"DbHost,omitnil,omitempty" name:"DbHost"`

	// User permissions Note: This field may return null, indicating that no valid values can be obtained.
	DbPrivilege *string `json:"DbPrivilege,omitnil,omitempty" name:"DbPrivilege"`
}

type ZoneStockInfo struct {
	// AZ
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Whether there is an inventory.
	HasStock *bool `json:"HasStock,omitnil,omitempty" name:"HasStock"`

	// Quantity in stock
	StockCount *int64 `json:"StockCount,omitnil,omitempty" name:"StockCount"`
}