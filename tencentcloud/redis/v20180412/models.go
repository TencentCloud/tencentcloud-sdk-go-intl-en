// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type Account struct {
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Account name.</p>
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// <p>Account description.</p>
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// <p>Read/write permission policy. - r: read-only. - w: write-only. - rw: read-write.</p>
	Privilege *string `json:"Privilege,omitnil,omitempty" name:"Privilege"`

	// <p>Read-only Routing Policy. - master: Master node. - replication: Replica node.</p>
	ReadonlyPolicy []*string `json:"ReadonlyPolicy,omitnil,omitempty" name:"ReadonlyPolicy"`

	// <p>Sub-account status. - 1: Account change in progress. - 2: Valid. - 4: Deleted.</p>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Account creation time.</p><p>If the parameter is an empty string, the account was created in an earlier version where the recording feature was not supported.</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>The time when the account last changed the password.</p><p>If the parameter is an empty string, it means the account was created in an earlier version that did not support the password modification time recording feature.</p>
	PasswordLastModifiedTime *string `json:"PasswordLastModifiedTime,omitnil,omitempty" name:"PasswordLastModifiedTime"`
}

// Predefined struct for user
type AddReplicationInstanceRequestParams struct {
	// Replication group ID. Log in to the [global replication](https://console.tencentcloud.com/redis/replication) page of the Redis console and obtain it.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Instance ID.
	// - There are region and AZ limitations for adding a replication group instance. For detailed information, see [Use Limits](https://www.tencentcloud.com/document/product/239/71934?from_cn_redirect=1).
	// - Currently, only Redis 4.0 and 5.0 cluster architecture instances support being added to the replication groups.
	// - Log in to the [Redis console](https://console.cloud.tencent.com/redis/instance/list), and copy the ID of the instance that needs to be added to the replication group in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Assigns roles to instances added to the replication group. <ul><li>rw: read-write;</li> <li>r: read-only.</li></ul>
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`
}

type AddReplicationInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Replication group ID. Log in to the [global replication](https://console.tencentcloud.com/redis/replication) page of the Redis console and obtain it.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Instance ID.
	// - There are region and AZ limitations for adding a replication group instance. For detailed information, see [Use Limits](https://www.tencentcloud.com/document/product/239/71934?from_cn_redirect=1).
	// - Currently, only Redis 4.0 and 5.0 cluster architecture instances support being added to the replication groups.
	// - Log in to the [Redis console](https://console.cloud.tencent.com/redis/instance/list), and copy the ID of the instance that needs to be added to the replication group in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Assigns roles to instances added to the replication group. <ul><li>rw: read-write;</li> <li>r: read-only.</li></ul>
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`
}

func (r *AddReplicationInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddReplicationInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "InstanceId")
	delete(f, "InstanceRole")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddReplicationInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddReplicationInstanceResponseParams struct {
	// Asynchronous process ID.
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddReplicationInstanceResponse struct {
	*tchttp.BaseResponse
	Response *AddReplicationInstanceResponseParams `json:"Response"`
}

func (r *AddReplicationInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddReplicationInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AllocateWanAddressRequestParams struct {
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type AllocateWanAddressRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *AllocateWanAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AllocateWanAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AllocateWanAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AllocateWanAddressResponseParams struct {
	// Async task ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Status of enabling public network access
	WanStatus *string `json:"WanStatus,omitnil,omitempty" name:"WanStatus"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AllocateWanAddressResponse struct {
	*tchttp.BaseResponse
	Response *AllocateWanAddressResponseParams `json:"Response"`
}

func (r *AllocateWanAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AllocateWanAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyParamsTemplateRequestParams struct {
	// Instance ID list. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy the instance ID in the instance list.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// ID of the applied parameter template.
	// - The parameter template ID can be obtained through the response parameter **TemplateId** of the API [DescribeParamTemplateInfo](https://www.tencentcloud.com/document/product/239/58748?from_cn_redirect=1).
	// - The operation can only be successfully performed when the parameter template version matches the architecture version of the target instance. A version mismatch will trigger an execution error.
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type ApplyParamsTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID list. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy the instance ID in the instance list.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// ID of the applied parameter template.
	// - The parameter template ID can be obtained through the response parameter **TemplateId** of the API [DescribeParamTemplateInfo](https://www.tencentcloud.com/document/product/239/58748?from_cn_redirect=1).
	// - The operation can only be successfully performed when the parameter template version matches the architecture version of the target instance. A version mismatch will trigger an execution error.
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *ApplyParamsTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyParamsTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyParamsTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyParamsTemplateResponseParams struct {
	// Task ID
	TaskIds []*int64 `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ApplyParamsTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ApplyParamsTemplateResponseParams `json:"Response"`
}

func (r *ApplyParamsTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyParamsTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateSecurityGroupsRequestParams struct {
	// Database engine name, which is `redis` for this API.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// ID of the security group to be bound. Obtain it on the [security group](https://console.tencentcloud.com/vpc/security-group) page of the console.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// ID of the bound instance. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list. You can specify multiple instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type AssociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Database engine name, which is `redis` for this API.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// ID of the security group to be bound. Obtain it on the [security group](https://console.tencentcloud.com/vpc/security-group) page of the console.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// ID of the bound instance. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list. You can specify multiple instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
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

// Predefined struct for user
type AssociateSecurityGroupsResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AssociateSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *AssociateSecurityGroupsResponseParams `json:"Response"`
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

type AvailableRegion struct {
	// Region
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// AZ information
	AvailableZones []*string `json:"AvailableZones,omitnil,omitempty" name:"AvailableZones"`
}

type BackupDownloadInfo struct {
	// Backup file name
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// Backup file size in bytes. If the parameter value is `0`, the backup file size is unknown.
	FileSize *int64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// Address (valid for six hours) used to download the backup file over the public network
	DownloadUrl *string `json:"DownloadUrl,omitnil,omitempty" name:"DownloadUrl"`

	// Address (valid for six hours) used to download the backup file over the private network
	InnerDownloadUrl *string `json:"InnerDownloadUrl,omitnil,omitempty" name:"InnerDownloadUrl"`
}

type BackupLimitVpcItem struct {
	// The region of the VPC that corresponds to the download address of the backup file
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// The list of VPCs that correspond to the download addresses of the backup files
	VpcList []*string `json:"VpcList,omitnil,omitempty" name:"VpcList"`
}

type BigKeyInfo struct {
	// Database
	DB *int64 `json:"DB,omitnil,omitempty" name:"DB"`

	// Big key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Type
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Size
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// Data timestamp
	Updatetime *int64 `json:"Updatetime,omitnil,omitempty" name:"Updatetime"`
}

type BigKeyTypeInfo struct {
	// Type
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Count
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// Size
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// Timestamp
	Updatetime *int64 `json:"Updatetime,omitnil,omitempty" name:"Updatetime"`
}

type CDCResource struct {
	// App ID of a user.
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Region ID.
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// AZ ID.
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// ID of the dedicated Redis cluster.
	RedisClusterId *string `json:"RedisClusterId,omitnil,omitempty" name:"RedisClusterId"`

	// Billing mode. 1: monthly subscription; 0: pay-as-you-go.
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Automatic renewal flag. 0: default status (manual renewal); 1: automatic renewal; 2: no automatic renewal.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// Dedicated cluster name.
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// Instance creation time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Instance expiration time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Cluster status. 1: in process; 2: running; 3: isolated.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Basic control resource package.
	BaseBundles []*ResourceBundle `json:"BaseBundles,omitnil,omitempty" name:"BaseBundles"`

	// Resource package list.
	ResourceBundles []*ResourceBundle `json:"ResourceBundles,omitnil,omitempty" name:"ResourceBundles"`

	// Local dedicated cluster ID.
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`
}

// Predefined struct for user
type ChangeInstanceRoleRequestParams struct {
	// Replication group ID. Log in to the [global replication](https://console.tencentcloud.com/redis/replication) page of the Redis console and obtain it.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance role.
	// - rw: read/write.
	// - r: read-only.
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`
}

type ChangeInstanceRoleRequest struct {
	*tchttp.BaseRequest
	
	// Replication group ID. Log in to the [global replication](https://console.tencentcloud.com/redis/replication) page of the Redis console and obtain it.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance role.
	// - rw: read/write.
	// - r: read-only.
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`
}

func (r *ChangeInstanceRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeInstanceRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "InstanceId")
	delete(f, "InstanceRole")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChangeInstanceRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChangeInstanceRoleResponseParams struct {
	// Async task ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChangeInstanceRoleResponse struct {
	*tchttp.BaseResponse
	Response *ChangeInstanceRoleResponseParams `json:"Response"`
}

func (r *ChangeInstanceRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeInstanceRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChangeMasterInstanceRequestParams struct {
	// Replication group ID, such as `crs-rpl-m3zt****`. It is the unique identifier automatically assigned by the system when creating a replication group. Log in to the [TencentDB for Redis console](https://console.cloud.tencent.com/redis/replication), and get it in the global replication list.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// ID of the read-only instance to be promoted to the master instance, such as `crs-xjhsdj****`. Log in to the the [TencentDB for Redis console](https://console.cloud.tencent.com/redis), and copy it in the instance list.
	// 
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Whether to promote the read-only instance to the master instance forcibly. Valid values:
	// - `true`: Yes
	// - `false`: No
	ForceSwitch *bool `json:"ForceSwitch,omitnil,omitempty" name:"ForceSwitch"`
}

type ChangeMasterInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Replication group ID, such as `crs-rpl-m3zt****`. It is the unique identifier automatically assigned by the system when creating a replication group. Log in to the [TencentDB for Redis console](https://console.cloud.tencent.com/redis/replication), and get it in the global replication list.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// ID of the read-only instance to be promoted to the master instance, such as `crs-xjhsdj****`. Log in to the the [TencentDB for Redis console](https://console.cloud.tencent.com/redis), and copy it in the instance list.
	// 
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Whether to promote the read-only instance to the master instance forcibly. Valid values:
	// - `true`: Yes
	// - `false`: No
	ForceSwitch *bool `json:"ForceSwitch,omitnil,omitempty" name:"ForceSwitch"`
}

func (r *ChangeMasterInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeMasterInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "InstanceId")
	delete(f, "ForceSwitch")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChangeMasterInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChangeMasterInstanceResponseParams struct {
	// Async task ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChangeMasterInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ChangeMasterInstanceResponseParams `json:"Response"`
}

func (r *ChangeMasterInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeMasterInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChangeReplicaToMasterRequestParams struct {
	// <p>Specifies the instance ID. Example: crs-xjhsdj****. Log in to the <a href="https://console.cloud.tencent.com/redis">Redis console</a> and copy the instance ID from the instance list.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Replica node group ID. Use the interface <a href="https://www.tencentcloud.com/document/product/239/50312?from_cn_redirect=1">DescribeInstanceZoneInfo</a> to obtain the id information of the multi-AZ replica node group. For a single AZ, no need to configure this parameter.</p>
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// <p>Emergency mode.</p><p>Enumeration values:</p><ul><li>false: Standard mode (Recommended for security)</li><li>true: Speed mode: (High-risk acceleration) Skip verification, speed up primary node promotion. High-level operation, highly likely to cause a single primary node in abnormal situations.</li></ul><p>Default value: false</p>
	Emergency *bool `json:"Emergency,omitnil,omitempty" name:"Emergency"`
}

type ChangeReplicaToMasterRequest struct {
	*tchttp.BaseRequest
	
	// <p>Specifies the instance ID. Example: crs-xjhsdj****. Log in to the <a href="https://console.cloud.tencent.com/redis">Redis console</a> and copy the instance ID from the instance list.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Replica node group ID. Use the interface <a href="https://www.tencentcloud.com/document/product/239/50312?from_cn_redirect=1">DescribeInstanceZoneInfo</a> to obtain the id information of the multi-AZ replica node group. For a single AZ, no need to configure this parameter.</p>
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// <p>Emergency mode.</p><p>Enumeration values:</p><ul><li>false: Standard mode (Recommended for security)</li><li>true: Speed mode: (High-risk acceleration) Skip verification, speed up primary node promotion. High-level operation, highly likely to cause a single primary node in abnormal situations.</li></ul><p>Default value: false</p>
	Emergency *bool `json:"Emergency,omitnil,omitempty" name:"Emergency"`
}

func (r *ChangeReplicaToMasterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeReplicaToMasterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "GroupId")
	delete(f, "Emergency")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChangeReplicaToMasterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChangeReplicaToMasterResponseParams struct {
	// <p>Asynchronous task ID.</p>
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChangeReplicaToMasterResponse struct {
	*tchttp.BaseResponse
	Response *ChangeReplicaToMasterResponseParams `json:"Response"`
}

func (r *ChangeReplicaToMasterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeReplicaToMasterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CleanUpInstanceRequestParams struct {
	// Instance ID. Log in to the [Redis console recycle bin](https://console.cloud.tencent.com/redis/recycle), and copy it from the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type CleanUpInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. Log in to the [Redis console recycle bin](https://console.cloud.tencent.com/redis/recycle), and copy it from the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *CleanUpInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CleanUpInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CleanUpInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CleanUpInstanceResponseParams struct {
	// Task ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CleanUpInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CleanUpInstanceResponseParams `json:"Response"`
}

func (r *CleanUpInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CleanUpInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ClearInstanceRequestParams struct {
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance access password.
	// - Password-free access: No configuration is required.
	// - Password authentication: The password is required. It cannot start with a forward slash (/) and should contain 8 to 64 characters, including at least two of the following types: lowercase letters, uppercase letters, digits, and special characters (such as ()`~!@#$%^&*-+=_|{}[]:;<>,.?/).
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Whether to encrypt the password.
	EncryptPassword *bool `json:"EncryptPassword,omitnil,omitempty" name:"EncryptPassword"`
}

type ClearInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance access password.
	// - Password-free access: No configuration is required.
	// - Password authentication: The password is required. It cannot start with a forward slash (/) and should contain 8 to 64 characters, including at least two of the following types: lowercase letters, uppercase letters, digits, and special characters (such as ()`~!@#$%^&*-+=_|{}[]:;<>,.?/).
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Whether to encrypt the password.
	EncryptPassword *bool `json:"EncryptPassword,omitnil,omitempty" name:"EncryptPassword"`
}

func (r *ClearInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClearInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Password")
	delete(f, "EncryptPassword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ClearInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ClearInstanceResponseParams struct {
	// Task ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ClearInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ClearInstanceResponseParams `json:"Response"`
}

func (r *ClearInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClearInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloneInstancesRequestParams struct {
	// <p>Specify the source instance ID to be cloned. Example: crs-xjhsdj****. Log in to the <a href="https://console.cloud.tencent.com/redis">Redis console</a> and copy the instance ID from the instance list.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>The number of clone instances per operation.</p><ul><li>The maximum allowed number for each Monthly Subscription purchase is 100.</li><li>The maximum allowed number for each Pay-As-You-Go purchase is 30.</li></ul>
	GoodsNum *uint64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// <p>Availability zone ID of the cloned instance. For supported AZ IDs, see <a href="https://www.tencentcloud.com/document/product/239/4106?from_cn_redirect=1">Regions and Availability Zones</a>.</p>
	ZoneId *uint64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>Payment method.<ul><li>0: Pay-As-You-Go.</li><li>1: Monthly Subscription.</li></ul></p>
	BillingMode *int64 `json:"BillingMode,omitnil,omitempty" name:"BillingMode"`

	// <p>Instance Purchase Duration.<ul><li>Unit: month.</li><li>When the payment method is set to Monthly Subscription, the value range is [1,2,3,4,5,6,7,8,9,10,11,12,24,36,48,60].</li><li>When the payment method is set to Pay-As-You-Go, it is set to 1.</li></ul></p>
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// <p>Security group ID. Call the <a href="https://www.tencentcloud.com/document/product/239/34447?from_cn_redirect=1">DescribeInstanceSecurityGroup</a> API to obtain the security group ID for the instance.</p>
	SecurityGroupIdList []*string `json:"SecurityGroupIdList,omitnil,omitempty" name:"SecurityGroupIdList"`

	// <p>Backup ID used to clone an instance. Use the interface <a href="https://www.tencentcloud.com/document/product/239/20011?from_cn_redirect=1">DescribeInstanceBackups</a> to obtain the backup ID.</p>
	BackupId *string `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// <p>Configure whether the cloned instance supports password-free access. Enabling SSL or public network does not support password-free access.<ul><li>true: Password-free instance,</li><li>false: Non-password-free instance. Default for non-passwordless instance.</li></ul></p>
	NoAuth *bool `json:"NoAuth,omitnil,omitempty" name:"NoAuth"`

	// <p>Configure the VPC ID for the clone instance. If not configured, the basic network is selected by default.</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>Configure the subnet of the private network for the cloned instance. This parameter requires no configuration for the basic network.</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>Name of the cloned instance.<br>Only Chinese characters, English letters, numbers, dashes ("-"), or underscores ("_") are allowed, with a length of less than 60.<br></p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>The access password of the clone instance.<ul><li>When the input parameter <b>NoAuth</b> is <b>true</b>, setting this parameter is optional.</li><li>For Redis 2.8, 4.0, and 5.0 instances, the password format is: 8-30 characters, containing at least lowercase letters, uppercase letters, digits, and 2 types of characters from ()`~!@#$%^&amp;*-+=_|{}[]:;&lt;&gt;,.?/, and cannot start with "/".</li><li>For CKV 3.2 instances, the password format is: 8-30 characters, must include letters and digits, and exclude other characters.</li></ul></p>
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// <p>Automatic renewal flag.<ul><li>0: default status, manual renewal.</li><li>1: automatic renewal.</li><li>2: no automatic renewal, auto-isolation upon expiration.</li></ul></p>
	AutoRenew *uint64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// <p>User-defined port, defaults to 6379, in the range of [1024,65535].</p>
	VPort *uint64 `json:"VPort,omitnil,omitempty" name:"VPort"`

	// <p>Node information of instance.<ul><li>Currently supports configuring node type (primary node or replica node) and its availability zone info. For details, please refer to <a href="https://www.tencentcloud.com/document/product/239/20022?from_cn_redirect=1#RedisNodeInfo">RedisNodeInfo</a>.</li><li>This parameter can be left blank for single-AZ deployment.</li></ul></p>
	NodeSet []*RedisNodeInfo `json:"NodeSet,omitnil,omitempty" name:"NodeSet"`

	// <p>Project ID. Log in to the <a href="https://console.cloud.tencent.com/redis#/">Redis console</a>. You can find the project ID in the <b>Account Center</b> &gt; <b>Project Management</b> at the top-right corner.</p>
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Tag bound to the clone instance.</p>
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// <p>Specify the parameter template ID related to the cloned instance.</p><ul><li>If this parameter is not configured, the system will automatically adapt the corresponding default template based on the selected compatible version and architecture.</li><li>Query the parameter template list of the instance through the <a href="https://www.tencentcloud.com/document/product/239/58750?from_cn_redirect=1">DescribeParamTemplates</a> API to obtain the template ID number.</li></ul>
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// <p>Specify the alarm policy ID of the clone instance. Log in to the <a href="https://console.cloud.tencent.com/monitor/alarm2/policy">Tencent Cloud observability platform console</a>, and get policy ID information on the <b>alarm management</b> &gt; <b>policy management</b> page.</p>
	AlarmPolicyList []*string `json:"AlarmPolicyList,omitnil,omitempty" name:"AlarmPolicyList"`

	// <p>Clone the time when data is recovered.<br>Only instances with second-level backup enabled are supported.</p>
	CloneTime *string `json:"CloneTime,omitnil,omitempty" name:"CloneTime"`

	// <p>Whether to encrypt the password</p>
	EncryptPassword *bool `json:"EncryptPassword,omitnil,omitempty" name:"EncryptPassword"`

	// <p>Instance password complexity policy</p><p>Input parameter limit: If not passed or Enabled=false, deem as not enabled and verify by default rule.</p>
	PasswordPolicy *PasswordPolicy `json:"PasswordPolicy,omitnil,omitempty" name:"PasswordPolicy"`

	// <p>Whether to enable SSL encryption.</p><p>Enumeration value:</p><ul><li>true: Enable.</li><li>false: Disable (default value).</li></ul><p>Default value: false</p>
	EnableSSL *bool `json:"EnableSSL,omitnil,omitempty" name:"EnableSSL"`

	// <p>Whether to write the private IPv4 address of the instance to the domain alias (SAN) of the certificate when SSL is enabled. This parameter is valid only when EnableSSL is true.</p><p>Enumeration value:</p><ul><li>true: The private IP is allowed for SSL certificate verification.</li><li>false: The SAN extended information of the certificate is not added.</li></ul><p>Default value: false</p>
	SSLBindPrivateIPv4 *bool `json:"SSLBindPrivateIPv4,omitnil,omitempty" name:"SSLBindPrivateIPv4"`

	// <p>Indicates the instance type.</p><p>Enumeration value:</p><ul><li>local: Common I</li><li>localv2: Common II</li></ul><p>If not passed, it remains the same as the original instance type by default.</p>
	ProductVersion *string `json:"ProductVersion,omitnil,omitempty" name:"ProductVersion"`
}

type CloneInstancesRequest struct {
	*tchttp.BaseRequest
	
	// <p>Specify the source instance ID to be cloned. Example: crs-xjhsdj****. Log in to the <a href="https://console.cloud.tencent.com/redis">Redis console</a> and copy the instance ID from the instance list.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>The number of clone instances per operation.</p><ul><li>The maximum allowed number for each Monthly Subscription purchase is 100.</li><li>The maximum allowed number for each Pay-As-You-Go purchase is 30.</li></ul>
	GoodsNum *uint64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// <p>Availability zone ID of the cloned instance. For supported AZ IDs, see <a href="https://www.tencentcloud.com/document/product/239/4106?from_cn_redirect=1">Regions and Availability Zones</a>.</p>
	ZoneId *uint64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>Payment method.<ul><li>0: Pay-As-You-Go.</li><li>1: Monthly Subscription.</li></ul></p>
	BillingMode *int64 `json:"BillingMode,omitnil,omitempty" name:"BillingMode"`

	// <p>Instance Purchase Duration.<ul><li>Unit: month.</li><li>When the payment method is set to Monthly Subscription, the value range is [1,2,3,4,5,6,7,8,9,10,11,12,24,36,48,60].</li><li>When the payment method is set to Pay-As-You-Go, it is set to 1.</li></ul></p>
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// <p>Security group ID. Call the <a href="https://www.tencentcloud.com/document/product/239/34447?from_cn_redirect=1">DescribeInstanceSecurityGroup</a> API to obtain the security group ID for the instance.</p>
	SecurityGroupIdList []*string `json:"SecurityGroupIdList,omitnil,omitempty" name:"SecurityGroupIdList"`

	// <p>Backup ID used to clone an instance. Use the interface <a href="https://www.tencentcloud.com/document/product/239/20011?from_cn_redirect=1">DescribeInstanceBackups</a> to obtain the backup ID.</p>
	BackupId *string `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// <p>Configure whether the cloned instance supports password-free access. Enabling SSL or public network does not support password-free access.<ul><li>true: Password-free instance,</li><li>false: Non-password-free instance. Default for non-passwordless instance.</li></ul></p>
	NoAuth *bool `json:"NoAuth,omitnil,omitempty" name:"NoAuth"`

	// <p>Configure the VPC ID for the clone instance. If not configured, the basic network is selected by default.</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>Configure the subnet of the private network for the cloned instance. This parameter requires no configuration for the basic network.</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>Name of the cloned instance.<br>Only Chinese characters, English letters, numbers, dashes ("-"), or underscores ("_") are allowed, with a length of less than 60.<br></p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>The access password of the clone instance.<ul><li>When the input parameter <b>NoAuth</b> is <b>true</b>, setting this parameter is optional.</li><li>For Redis 2.8, 4.0, and 5.0 instances, the password format is: 8-30 characters, containing at least lowercase letters, uppercase letters, digits, and 2 types of characters from ()`~!@#$%^&amp;*-+=_|{}[]:;&lt;&gt;,.?/, and cannot start with "/".</li><li>For CKV 3.2 instances, the password format is: 8-30 characters, must include letters and digits, and exclude other characters.</li></ul></p>
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// <p>Automatic renewal flag.<ul><li>0: default status, manual renewal.</li><li>1: automatic renewal.</li><li>2: no automatic renewal, auto-isolation upon expiration.</li></ul></p>
	AutoRenew *uint64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// <p>User-defined port, defaults to 6379, in the range of [1024,65535].</p>
	VPort *uint64 `json:"VPort,omitnil,omitempty" name:"VPort"`

	// <p>Node information of instance.<ul><li>Currently supports configuring node type (primary node or replica node) and its availability zone info. For details, please refer to <a href="https://www.tencentcloud.com/document/product/239/20022?from_cn_redirect=1#RedisNodeInfo">RedisNodeInfo</a>.</li><li>This parameter can be left blank for single-AZ deployment.</li></ul></p>
	NodeSet []*RedisNodeInfo `json:"NodeSet,omitnil,omitempty" name:"NodeSet"`

	// <p>Project ID. Log in to the <a href="https://console.cloud.tencent.com/redis#/">Redis console</a>. You can find the project ID in the <b>Account Center</b> &gt; <b>Project Management</b> at the top-right corner.</p>
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Tag bound to the clone instance.</p>
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// <p>Specify the parameter template ID related to the cloned instance.</p><ul><li>If this parameter is not configured, the system will automatically adapt the corresponding default template based on the selected compatible version and architecture.</li><li>Query the parameter template list of the instance through the <a href="https://www.tencentcloud.com/document/product/239/58750?from_cn_redirect=1">DescribeParamTemplates</a> API to obtain the template ID number.</li></ul>
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// <p>Specify the alarm policy ID of the clone instance. Log in to the <a href="https://console.cloud.tencent.com/monitor/alarm2/policy">Tencent Cloud observability platform console</a>, and get policy ID information on the <b>alarm management</b> &gt; <b>policy management</b> page.</p>
	AlarmPolicyList []*string `json:"AlarmPolicyList,omitnil,omitempty" name:"AlarmPolicyList"`

	// <p>Clone the time when data is recovered.<br>Only instances with second-level backup enabled are supported.</p>
	CloneTime *string `json:"CloneTime,omitnil,omitempty" name:"CloneTime"`

	// <p>Whether to encrypt the password</p>
	EncryptPassword *bool `json:"EncryptPassword,omitnil,omitempty" name:"EncryptPassword"`

	// <p>Instance password complexity policy</p><p>Input parameter limit: If not passed or Enabled=false, deem as not enabled and verify by default rule.</p>
	PasswordPolicy *PasswordPolicy `json:"PasswordPolicy,omitnil,omitempty" name:"PasswordPolicy"`

	// <p>Whether to enable SSL encryption.</p><p>Enumeration value:</p><ul><li>true: Enable.</li><li>false: Disable (default value).</li></ul><p>Default value: false</p>
	EnableSSL *bool `json:"EnableSSL,omitnil,omitempty" name:"EnableSSL"`

	// <p>Whether to write the private IPv4 address of the instance to the domain alias (SAN) of the certificate when SSL is enabled. This parameter is valid only when EnableSSL is true.</p><p>Enumeration value:</p><ul><li>true: The private IP is allowed for SSL certificate verification.</li><li>false: The SAN extended information of the certificate is not added.</li></ul><p>Default value: false</p>
	SSLBindPrivateIPv4 *bool `json:"SSLBindPrivateIPv4,omitnil,omitempty" name:"SSLBindPrivateIPv4"`

	// <p>Indicates the instance type.</p><p>Enumeration value:</p><ul><li>local: Common I</li><li>localv2: Common II</li></ul><p>If not passed, it remains the same as the original instance type by default.</p>
	ProductVersion *string `json:"ProductVersion,omitnil,omitempty" name:"ProductVersion"`
}

func (r *CloneInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloneInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "GoodsNum")
	delete(f, "ZoneId")
	delete(f, "BillingMode")
	delete(f, "Period")
	delete(f, "SecurityGroupIdList")
	delete(f, "BackupId")
	delete(f, "NoAuth")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "InstanceName")
	delete(f, "Password")
	delete(f, "AutoRenew")
	delete(f, "VPort")
	delete(f, "NodeSet")
	delete(f, "ProjectId")
	delete(f, "ResourceTags")
	delete(f, "TemplateId")
	delete(f, "AlarmPolicyList")
	delete(f, "CloneTime")
	delete(f, "EncryptPassword")
	delete(f, "PasswordPolicy")
	delete(f, "EnableSSL")
	delete(f, "SSLBindPrivateIPv4")
	delete(f, "ProductVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloneInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloneInstancesResponseParams struct {
	// <p>Transaction ID.</p>
	//
	// Deprecated: DealId is deprecated.
	DealId *string `json:"DealId,omitnil,omitempty" name:"DealId"`

	// <p>ID of the clone instance.</p>
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// <p>Order ID.</p>
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloneInstancesResponse struct {
	*tchttp.BaseResponse
	Response *CloneInstancesResponseParams `json:"Response"`
}

func (r *CloneInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloneInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseLogRequestParams struct {
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Log type.</p>
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`
}

type CloseLogRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Log type.</p>
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`
}

func (r *CloseLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "LogType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseLogResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloseLogResponse struct {
	*tchttp.BaseResponse
	Response *CloseLogResponseParams `json:"Response"`
}

func (r *CloseLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseSSLRequestParams struct {
	// <p>Instance ID. Log in to the <a href="https://console.cloud.tencent.com/redis">Redis console</a> and copy it in the instance list.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>SSL address type.</p><p>Enumeration value:</p><ul><li>0: Unlimited.</li><li>1: Private IPv4 address.</li><li>2: Private IPv6 address.</li><li>3: Public network.</li><li>-1: Unspecified.</li></ul><p>Default value: 0</p>
	AddressType *int64 `json:"AddressType,omitnil,omitempty" name:"AddressType"`
}

type CloseSSLRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance ID. Log in to the <a href="https://console.cloud.tencent.com/redis">Redis console</a> and copy it in the instance list.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>SSL address type.</p><p>Enumeration value:</p><ul><li>0: Unlimited.</li><li>1: Private IPv4 address.</li><li>2: Private IPv6 address.</li><li>3: Public network.</li><li>-1: Unspecified.</li></ul><p>Default value: 0</p>
	AddressType *int64 `json:"AddressType,omitnil,omitempty" name:"AddressType"`
}

func (r *CloseSSLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseSSLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AddressType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseSSLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseSSLResponseParams struct {
	// <p>Task ID.</p>
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloseSSLResponse struct {
	*tchttp.BaseResponse
	Response *CloseSSLResponseParams `json:"Response"`
}

func (r *CloseSSLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseSSLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CommandTake struct {
	// Command name.
	Cmd *string `json:"Cmd,omitnil,omitempty" name:"Cmd"`

	// Time consumed. Unit: ms.
	Took *int64 `json:"Took,omitnil,omitempty" name:"Took"`
}

// Predefined struct for user
type CreateExportTaskRequestParams struct {
	// <p>Specify the instance ID. Example: crs-xjhsdj****. Log in to the <a href="https://console.cloud.tencent.com/redis">Redis console</a> and copy the instance ID from the instance list.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Log type.</p><p>Enumeration value:</p><ul><li>auditLog: Audit log.</li></ul>
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// <p>Start time of retrieval.</p><p>Parameter format: YYYY-MM-DD HH:mm:ss, for example 2026-03-06 00:00:00. The returned result contains only the logs at this time point and afterward.</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>End time of log retrieval.</p><p>Parameter format: YYYY-MM-DD HH:mm:ss, for example 2026-03-06 23:59:59. The returned result contains only the logs at this time point and earlier.</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>Set the log filter field to filter and download qualified logs.</p>
	LogFilter []*LogFilter `json:"LogFilter,omitnil,omitempty" name:"LogFilter"`

	// <p>Custom log fields for download, multiple fields separated by commas, such as "timestamp,operation,user". Only the data of selected fields will be downloaded when specified. The parameter defaults to downloading all fields when not passed.</p>
	ColumnFilter []*string `json:"ColumnFilter,omitnil,omitempty" name:"ColumnFilter"`
}

type CreateExportTaskRequest struct {
	*tchttp.BaseRequest
	
	// <p>Specify the instance ID. Example: crs-xjhsdj****. Log in to the <a href="https://console.cloud.tencent.com/redis">Redis console</a> and copy the instance ID from the instance list.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Log type.</p><p>Enumeration value:</p><ul><li>auditLog: Audit log.</li></ul>
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// <p>Start time of retrieval.</p><p>Parameter format: YYYY-MM-DD HH:mm:ss, for example 2026-03-06 00:00:00. The returned result contains only the logs at this time point and afterward.</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>End time of log retrieval.</p><p>Parameter format: YYYY-MM-DD HH:mm:ss, for example 2026-03-06 23:59:59. The returned result contains only the logs at this time point and earlier.</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>Set the log filter field to filter and download qualified logs.</p>
	LogFilter []*LogFilter `json:"LogFilter,omitnil,omitempty" name:"LogFilter"`

	// <p>Custom log fields for download, multiple fields separated by commas, such as "timestamp,operation,user". Only the data of selected fields will be downloaded when specified. The parameter defaults to downloading all fields when not passed.</p>
	ColumnFilter []*string `json:"ColumnFilter,omitnil,omitempty" name:"ColumnFilter"`
}

func (r *CreateExportTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateExportTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "LogType")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "LogFilter")
	delete(f, "ColumnFilter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateExportTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateExportTaskResponseParams struct {
	// <p>File name.</p>
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateExportTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateExportTaskResponseParams `json:"Response"`
}

func (r *CreateExportTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateExportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceAccountRequestParams struct {
	// <p>Instance ID. Log in to the <a href="https://console.cloud.tencent.com/redis">Redis console</a> and copy the instance ID from the instance list.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Custom account name for accessing the database.</p><ul><li>Consist of letters, digits, underscores, and hyphens only.</li><li>Length cannot be greater than 32.</li></ul>
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// <p>Set a password for the customized account. The password complexity requirements are as follows:</p><ul><li>Character count: [8,64].</li><li>Contain at least two kinds of lowercase letters, uppercase letters, digits and characters ()`~!@#$%^&amp;*-+=_|{}[]:;&lt;&gt;,.?/.</li><li>Cannot start with "/".</li></ul>
	AccountPassword *string `json:"AccountPassword,omitnil,omitempty" name:"AccountPassword"`

	// <p>Read requests for the designated account are routed to the primary node or replica node. Read-only replica is not enabled, and selection of replica nodes is not supported.</p><ul><li>master: primary node</li><li>replication: replica node</li></ul>
	ReadonlyPolicy []*string `json:"ReadonlyPolicy,omitnil,omitempty" name:"ReadonlyPolicy"`

	// <p>Account read/write permission supports selecting read-only or read-write permission.</p><ul><li>r: Read-only.</li><li>rw: Read-write.</li></ul>
	Privilege *string `json:"Privilege,omitnil,omitempty" name:"Privilege"`

	// <p>Description information about account remarks, with a length of [0, 64] bytes.</p>
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// <p>Whether to enable password encryption for transmission.</p><ul><li>true: Encrypted.</li><li>false: Not encrypted (default value).</li></ul>
	EncryptPassword *bool `json:"EncryptPassword,omitnil,omitempty" name:"EncryptPassword"`
}

type CreateInstanceAccountRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance ID. Log in to the <a href="https://console.cloud.tencent.com/redis">Redis console</a> and copy the instance ID from the instance list.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Custom account name for accessing the database.</p><ul><li>Consist of letters, digits, underscores, and hyphens only.</li><li>Length cannot be greater than 32.</li></ul>
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// <p>Set a password for the customized account. The password complexity requirements are as follows:</p><ul><li>Character count: [8,64].</li><li>Contain at least two kinds of lowercase letters, uppercase letters, digits and characters ()`~!@#$%^&amp;*-+=_|{}[]:;&lt;&gt;,.?/.</li><li>Cannot start with "/".</li></ul>
	AccountPassword *string `json:"AccountPassword,omitnil,omitempty" name:"AccountPassword"`

	// <p>Read requests for the designated account are routed to the primary node or replica node. Read-only replica is not enabled, and selection of replica nodes is not supported.</p><ul><li>master: primary node</li><li>replication: replica node</li></ul>
	ReadonlyPolicy []*string `json:"ReadonlyPolicy,omitnil,omitempty" name:"ReadonlyPolicy"`

	// <p>Account read/write permission supports selecting read-only or read-write permission.</p><ul><li>r: Read-only.</li><li>rw: Read-write.</li></ul>
	Privilege *string `json:"Privilege,omitnil,omitempty" name:"Privilege"`

	// <p>Description information about account remarks, with a length of [0, 64] bytes.</p>
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// <p>Whether to enable password encryption for transmission.</p><ul><li>true: Encrypted.</li><li>false: Not encrypted (default value).</li></ul>
	EncryptPassword *bool `json:"EncryptPassword,omitnil,omitempty" name:"EncryptPassword"`
}

func (r *CreateInstanceAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AccountName")
	delete(f, "AccountPassword")
	delete(f, "ReadonlyPolicy")
	delete(f, "Privilege")
	delete(f, "Remark")
	delete(f, "EncryptPassword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstanceAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceAccountResponseParams struct {
	// <p>Task ID.</p>
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateInstanceAccountResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstanceAccountResponseParams `json:"Response"`
}

func (r *CreateInstanceAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstancesRequestParams struct {
	// <p>Instance type.</p><ul><li>2: Redis 2.8 memory edition (standard architecture).</li><li>3: CKV 3.2 memory edition (standard architecture).</li><li>4: CKV 3.2 memory edition (cluster architecture).</li><li>6: Redis 4.0 memory edition (standard architecture).</li><li>7: Redis 4.0 memory edition (cluster architecture).</li><li>8: Redis 5.0 memory edition (standard architecture).</li><li>9: Redis 5.0 memory edition (cluster architecture).</li><li>15: Redis 6.2 memory edition (standard architecture).</li><li>16: Redis 6.2 memory edition (cluster architecture).</li><li>17: Redis 7.0 memory edition (standard architecture).</li><li>18: Redis 7.0 memory edition (cluster architecture).</li><li>19: Valkey 8.0 memory edition (standard architecture).</li><li>20: Valkey 8.0 memory edition (cluster architecture).</li><li>21: Valkey 9.0 memory edition (standard architecture).</li><li>22: Valkey 9.0 memory edition (cluster architecture).</li><li>200: Memcached 1.6 memory edition (cluster architecture).<br><strong>Note</strong>: CKV editions are currently used by some users and are temporarily retained.</li></ul>
	TypeId *uint64 `json:"TypeId,omitnil,omitempty" name:"TypeId"`

	// <p>Memory capacity, measured in MB, must be a multiple of 1024. For specific specifications, query the sales specifications for all regions via the <a href="https://www.tencentcloud.com/document/api/239/30600?from_cn_redirect=1">DescribeProductInfo</a> API.</p><ul><li>When <strong>TypeId</strong> is standard architecture, <strong>MemSize</strong> is the total memory capacity of the instance.</li><li>When <strong>TypeId</strong> is cluster architecture, <strong>MemSize</strong> is the sharded memory capacity.</li></ul>
	MemSize *uint64 `json:"MemSize,omitnil,omitempty" name:"MemSize"`

	// <p>Instance count. Number of instances to purchase at a time. For details, query sales specifications in all regions via the <a href="https://www.tencentcloud.com/document/api/239/30600?from_cn_redirect=1">DescribeProductInfo</a> API.</p>
	GoodsNum *uint64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// <p>Duration of instance purchase.</p><ul><li>If <strong>BillingMode</strong> is <strong>1</strong>, that is, the billing mode is monthly subscription, you need to set this parameter to specify the duration of instance purchase. Measurement unit: month, permissible range [1,2,3,4,5,6,7,8,9,10,11,12,24,36].</li><li>If <strong>BillingMode</strong> is <strong>0</strong>, that is, the billing mode is pay-as-you-go, set this parameter to 1.</li></ul>
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// <p>Billing mode.</p><ul><li>0: Pay-As-You-Go.</li><li>1: Monthly Subscription.</li></ul>
	BillingMode *int64 `json:"BillingMode,omitnil,omitempty" name:"BillingMode"`

	// <p>ID of the AZ to which the instance belongs. See <a href="https://www.tencentcloud.com/document/product/239/4106?from_cn_redirect=1">Regions and AZs</a>.</p>
	ZoneId *uint64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>Password for accessing instances.</p><ul><li>When the input parameter <strong>NoAuth</strong> is <strong>true</strong>, it means setting instances to Password-free access, and Password does not need to be configured. Otherwise, Password is required.</li><li>When the instance type <strong>TypeId</strong> is Redis 2.8 memory edition standard architecture, Redis 4.0, 5.0, 6.2, or 7.0 memory edition standard architecture or cluster architecture, the Password complexity requirements are: 8-64 characters, containing at least lowercase letters, uppercase letters, digits, and 2 of the following characters: ()`~!@#$%^&amp;*-+=_|{}[]:;&lt;&gt;,.?/, and cannot start with "/".</li><li>When the instance type <strong>TypeId</strong> is CKV 3.2 memory edition standard architecture or cluster architecture, the Password complexity is: 8-30 characters, must include letters and digits, and exclude other characters.</li></ul>
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// <p>VPC ID. If you do not configure this parameter, the basic network is selected by default. Log in to the <a href="https://console.cloud.tencent.com/vpc">private network</a> console to query the specific ID.</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>Subnet of the private network VPC. In the basic network, this parameter requires no configuration. Log in to the <a href="https://console.cloud.tencent.com/vpc">Private Network</a> console to query the subnet list and obtain the specific ID.</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>Project ID. Log in to the <a href="https://console.cloud.tencent.com/redis#/">Redis console</a>, select <strong>Project Management</strong> from the account information menu in the top-right corner, and query the project ID.</p>
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Automatic renewal flag.</p><ul><li>0: default status (manual renewal).</li><li>1: automatic renewal.</li><li>2: non-renewal upon expiration.</li></ul>
	AutoRenew *uint64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// <p>Security group ID array.</p><ul><li>A security group is a virtual firewall that controls network access to a cloud database instance. When creating an instance, it is recommended to bind the corresponding security group.</li><li>Get the security group ID of an instance through the <a href="https://www.tencentcloud.com/document/product/239/34447?from_cn_redirect=1">DescribeInstanceSecurityGroup</a> API.</li></ul>
	SecurityGroupIdList []*string `json:"SecurityGroupIdList,omitnil,omitempty" name:"SecurityGroupIdList"`

	// <p>User-defined network port. Defaults to 6379, range [1024,65535].</p>
	VPort *uint64 `json:"VPort,omitnil,omitempty" name:"VPort"`

	// <p>Number of instance shards.</p><ul><li>No need to configure this parameter for standard edition instances.</li><li>For cluster edition instances, the shard quantity range is: [1, 3, 5, 8, 12, 16, 24, 32, 40, 48, 64, 80, 96, 128].</li></ul>
	RedisShardNum *int64 `json:"RedisShardNum,omitnil,omitempty" name:"RedisShardNum"`

	// <p>Instance replica number.</p><ul><li>Redis Memory Edition 4.0, 5.0, 6.2, and 7.0 cluster architecture supports a replica quantity range of [1,5].</li><li>Redis 2.8 Standard Edition and CKV Standard Edition support only 1 replica.</li></ul>
	RedisReplicasNum *int64 `json:"RedisReplicasNum,omitnil,omitempty" name:"RedisReplicasNum"`

	// <p>Flag whether the instance needs to support read-only replica.</p><ul><li>Redis 2.8 Standard Edition and CKV Standard Edition do not support read-only replica.</li><li>If read-only replica is enabled, the instance will automatically separate read and write operations, with write requests routed to the primary node and read requests routed to replica nodes.</li><li>If needed to enable read-only replica, it is recommended to have at least 2 replicas.</li></ul>
	ReplicasReadonly *bool `json:"ReplicasReadonly,omitnil,omitempty" name:"ReplicasReadonly"`

	// <p>Instance name. Naming requirement: It only supports Chinese characters, letters, numbers, hyphens ("-"), or underscores ("_"), with a length of less than 60.</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>Configure whether the instance supports password-free access.</p><ul><li>true: Access the instance without a password.</li><li>false: Access the instance with a password. By default, password access is enabled. Only instances in a VPC network support password-free access.</li></ul>
	NoAuth *bool `json:"NoAuth,omitnil,omitempty" name:"NoAuth"`

	// <p>Node information of instance, including node ID, node type, node availability zone ID. For details, please see <a href="https://www.tencentcloud.com/document/product/239/20022?from_cn_redirect=1">RedisNodeInfo</a>.<br>Currently support inputting node type (primary node or replica node) and availability zone. When this parameter is not specified, in regions that support multi-availability zone deployment, the system defaults to creating instances with multi-availability zone architecture.</p>
	NodeSet []*RedisNodeInfo `json:"NodeSet,omitnil,omitempty" name:"NodeSet"`

	// <p>Set a tag for the instance.</p>
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// <p>Specify the name of the AZ to which the instance belongs. For details, see <a href="https://www.tencentcloud.com/document/product/239/4106?from_cn_redirect=1">Regions and AZs</a>.</p>
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// <p>Parameter template ID of the specified instance.</p><ul><li>If this parameter is not configured, the system will automatically adapt to the corresponding default template based on the selected compatible version and architecture.</li><li>Query the parameter template list of the instance through the <a href="https://www.tencentcloud.com/document/product/239/58750?from_cn_redirect=1">DescribeParamTemplates</a> API to obtain the template ID number.</li></ul>
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// <p>Internal parameter to identify whether creating an instance needs to check.</p><ul><li>false: Default value. Send a normal request and create the instance directly after passing the check.</li><li>true: Send a check request without creating an instance.</li></ul>
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// <p>Indicates the instance deployment mode.</p><ul><li>local: traditional architecture, defaults to local.</li><li>cdc: dedicated cluster.</li><li>cloud: cloud native, currently not available for sale.</li></ul>
	ProductVersion *string `json:"ProductVersion,omitnil,omitempty" name:"ProductVersion"`

	// <p>Dedicated cluster ID.</p><ul><li>When <strong>ProductVersion</strong> is set to <strong>cdc</strong>, this parameter must be set.</li><li>Get cluster ID through the API <a href="https://www.tencentcloud.com/document/product/239/109628?from_cn_redirect=1">DescribeRedisClusters</a>.</li></ul>
	RedisClusterId *string `json:"RedisClusterId,omitnil,omitempty" name:"RedisClusterId"`

	// <p>Alarm policy ID array.</p><ul><li>Log in to <a href="https://console.cloud.tencent.com/monitor/alarm/policy">Tencent Cloud Observability Platform - Alarm Management - Policy Management</a> to get alarm policy ID.</li><li>If this parameter is not configured, the default alarm policy will be bound. For the default alarm policy details, log in to <a href="https://console.cloud.tencent.com/monitor/alarm/policy">Tencent Cloud Observability Platform - Alarm Management - Policy Management</a> to view.</li></ul>
	AlarmPolicyList []*string `json:"AlarmPolicyList,omitnil,omitempty" name:"AlarmPolicyList"`

	// <p>Whether to enable password encryption for transmission.</p><ul><li>true: Encrypted.</li><li>false: Not encrypted (default value).</li></ul>
	EncryptPassword *bool `json:"EncryptPassword,omitnil,omitempty" name:"EncryptPassword"`

	// <p>Instance-level password complexity policy. When not passed in or Enabled=false, deem as not enabling policy, validate by system default rule.</p>
	PasswordPolicy *PasswordPolicy `json:"PasswordPolicy,omitnil,omitempty" name:"PasswordPolicy"`

	// <p>Whether to enable SSL encryption.</p><ul><li>true: Enable.</li><li>false: Disable (default value).</li></ul>
	EnableSSL *bool `json:"EnableSSL,omitnil,omitempty" name:"EnableSSL"`

	// <p>Whether to write the private IPv4 address of an instance to the domain alias (SAN) of the certificate when SSL is enabled. This parameter is valid only when EnableSSL is true.</p><ul><li>true: Allows using private IP to perform SSL certificate verification.</li><li>false: Does not add the SAN extended information to the certificate.</li></ul>
	SSLBindPrivateIPv4 *bool `json:"SSLBindPrivateIPv4,omitnil,omitempty" name:"SSLBindPrivateIPv4"`

	// <p>Instance connectivity access Mode.</p><ul><li>0: Proxy Mode (default value).</li><li>1: Direct access Mode.</li></ul>
	ConnectionMode *int64 `json:"ConnectionMode,omitnil,omitempty" name:"ConnectionMode"`
}

type CreateInstancesRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance type.</p><ul><li>2: Redis 2.8 memory edition (standard architecture).</li><li>3: CKV 3.2 memory edition (standard architecture).</li><li>4: CKV 3.2 memory edition (cluster architecture).</li><li>6: Redis 4.0 memory edition (standard architecture).</li><li>7: Redis 4.0 memory edition (cluster architecture).</li><li>8: Redis 5.0 memory edition (standard architecture).</li><li>9: Redis 5.0 memory edition (cluster architecture).</li><li>15: Redis 6.2 memory edition (standard architecture).</li><li>16: Redis 6.2 memory edition (cluster architecture).</li><li>17: Redis 7.0 memory edition (standard architecture).</li><li>18: Redis 7.0 memory edition (cluster architecture).</li><li>19: Valkey 8.0 memory edition (standard architecture).</li><li>20: Valkey 8.0 memory edition (cluster architecture).</li><li>21: Valkey 9.0 memory edition (standard architecture).</li><li>22: Valkey 9.0 memory edition (cluster architecture).</li><li>200: Memcached 1.6 memory edition (cluster architecture).<br><strong>Note</strong>: CKV editions are currently used by some users and are temporarily retained.</li></ul>
	TypeId *uint64 `json:"TypeId,omitnil,omitempty" name:"TypeId"`

	// <p>Memory capacity, measured in MB, must be a multiple of 1024. For specific specifications, query the sales specifications for all regions via the <a href="https://www.tencentcloud.com/document/api/239/30600?from_cn_redirect=1">DescribeProductInfo</a> API.</p><ul><li>When <strong>TypeId</strong> is standard architecture, <strong>MemSize</strong> is the total memory capacity of the instance.</li><li>When <strong>TypeId</strong> is cluster architecture, <strong>MemSize</strong> is the sharded memory capacity.</li></ul>
	MemSize *uint64 `json:"MemSize,omitnil,omitempty" name:"MemSize"`

	// <p>Instance count. Number of instances to purchase at a time. For details, query sales specifications in all regions via the <a href="https://www.tencentcloud.com/document/api/239/30600?from_cn_redirect=1">DescribeProductInfo</a> API.</p>
	GoodsNum *uint64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// <p>Duration of instance purchase.</p><ul><li>If <strong>BillingMode</strong> is <strong>1</strong>, that is, the billing mode is monthly subscription, you need to set this parameter to specify the duration of instance purchase. Measurement unit: month, permissible range [1,2,3,4,5,6,7,8,9,10,11,12,24,36].</li><li>If <strong>BillingMode</strong> is <strong>0</strong>, that is, the billing mode is pay-as-you-go, set this parameter to 1.</li></ul>
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// <p>Billing mode.</p><ul><li>0: Pay-As-You-Go.</li><li>1: Monthly Subscription.</li></ul>
	BillingMode *int64 `json:"BillingMode,omitnil,omitempty" name:"BillingMode"`

	// <p>ID of the AZ to which the instance belongs. See <a href="https://www.tencentcloud.com/document/product/239/4106?from_cn_redirect=1">Regions and AZs</a>.</p>
	ZoneId *uint64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>Password for accessing instances.</p><ul><li>When the input parameter <strong>NoAuth</strong> is <strong>true</strong>, it means setting instances to Password-free access, and Password does not need to be configured. Otherwise, Password is required.</li><li>When the instance type <strong>TypeId</strong> is Redis 2.8 memory edition standard architecture, Redis 4.0, 5.0, 6.2, or 7.0 memory edition standard architecture or cluster architecture, the Password complexity requirements are: 8-64 characters, containing at least lowercase letters, uppercase letters, digits, and 2 of the following characters: ()`~!@#$%^&amp;*-+=_|{}[]:;&lt;&gt;,.?/, and cannot start with "/".</li><li>When the instance type <strong>TypeId</strong> is CKV 3.2 memory edition standard architecture or cluster architecture, the Password complexity is: 8-30 characters, must include letters and digits, and exclude other characters.</li></ul>
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// <p>VPC ID. If you do not configure this parameter, the basic network is selected by default. Log in to the <a href="https://console.cloud.tencent.com/vpc">private network</a> console to query the specific ID.</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>Subnet of the private network VPC. In the basic network, this parameter requires no configuration. Log in to the <a href="https://console.cloud.tencent.com/vpc">Private Network</a> console to query the subnet list and obtain the specific ID.</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>Project ID. Log in to the <a href="https://console.cloud.tencent.com/redis#/">Redis console</a>, select <strong>Project Management</strong> from the account information menu in the top-right corner, and query the project ID.</p>
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Automatic renewal flag.</p><ul><li>0: default status (manual renewal).</li><li>1: automatic renewal.</li><li>2: non-renewal upon expiration.</li></ul>
	AutoRenew *uint64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// <p>Security group ID array.</p><ul><li>A security group is a virtual firewall that controls network access to a cloud database instance. When creating an instance, it is recommended to bind the corresponding security group.</li><li>Get the security group ID of an instance through the <a href="https://www.tencentcloud.com/document/product/239/34447?from_cn_redirect=1">DescribeInstanceSecurityGroup</a> API.</li></ul>
	SecurityGroupIdList []*string `json:"SecurityGroupIdList,omitnil,omitempty" name:"SecurityGroupIdList"`

	// <p>User-defined network port. Defaults to 6379, range [1024,65535].</p>
	VPort *uint64 `json:"VPort,omitnil,omitempty" name:"VPort"`

	// <p>Number of instance shards.</p><ul><li>No need to configure this parameter for standard edition instances.</li><li>For cluster edition instances, the shard quantity range is: [1, 3, 5, 8, 12, 16, 24, 32, 40, 48, 64, 80, 96, 128].</li></ul>
	RedisShardNum *int64 `json:"RedisShardNum,omitnil,omitempty" name:"RedisShardNum"`

	// <p>Instance replica number.</p><ul><li>Redis Memory Edition 4.0, 5.0, 6.2, and 7.0 cluster architecture supports a replica quantity range of [1,5].</li><li>Redis 2.8 Standard Edition and CKV Standard Edition support only 1 replica.</li></ul>
	RedisReplicasNum *int64 `json:"RedisReplicasNum,omitnil,omitempty" name:"RedisReplicasNum"`

	// <p>Flag whether the instance needs to support read-only replica.</p><ul><li>Redis 2.8 Standard Edition and CKV Standard Edition do not support read-only replica.</li><li>If read-only replica is enabled, the instance will automatically separate read and write operations, with write requests routed to the primary node and read requests routed to replica nodes.</li><li>If needed to enable read-only replica, it is recommended to have at least 2 replicas.</li></ul>
	ReplicasReadonly *bool `json:"ReplicasReadonly,omitnil,omitempty" name:"ReplicasReadonly"`

	// <p>Instance name. Naming requirement: It only supports Chinese characters, letters, numbers, hyphens ("-"), or underscores ("_"), with a length of less than 60.</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>Configure whether the instance supports password-free access.</p><ul><li>true: Access the instance without a password.</li><li>false: Access the instance with a password. By default, password access is enabled. Only instances in a VPC network support password-free access.</li></ul>
	NoAuth *bool `json:"NoAuth,omitnil,omitempty" name:"NoAuth"`

	// <p>Node information of instance, including node ID, node type, node availability zone ID. For details, please see <a href="https://www.tencentcloud.com/document/product/239/20022?from_cn_redirect=1">RedisNodeInfo</a>.<br>Currently support inputting node type (primary node or replica node) and availability zone. When this parameter is not specified, in regions that support multi-availability zone deployment, the system defaults to creating instances with multi-availability zone architecture.</p>
	NodeSet []*RedisNodeInfo `json:"NodeSet,omitnil,omitempty" name:"NodeSet"`

	// <p>Set a tag for the instance.</p>
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// <p>Specify the name of the AZ to which the instance belongs. For details, see <a href="https://www.tencentcloud.com/document/product/239/4106?from_cn_redirect=1">Regions and AZs</a>.</p>
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// <p>Parameter template ID of the specified instance.</p><ul><li>If this parameter is not configured, the system will automatically adapt to the corresponding default template based on the selected compatible version and architecture.</li><li>Query the parameter template list of the instance through the <a href="https://www.tencentcloud.com/document/product/239/58750?from_cn_redirect=1">DescribeParamTemplates</a> API to obtain the template ID number.</li></ul>
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// <p>Internal parameter to identify whether creating an instance needs to check.</p><ul><li>false: Default value. Send a normal request and create the instance directly after passing the check.</li><li>true: Send a check request without creating an instance.</li></ul>
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// <p>Indicates the instance deployment mode.</p><ul><li>local: traditional architecture, defaults to local.</li><li>cdc: dedicated cluster.</li><li>cloud: cloud native, currently not available for sale.</li></ul>
	ProductVersion *string `json:"ProductVersion,omitnil,omitempty" name:"ProductVersion"`

	// <p>Dedicated cluster ID.</p><ul><li>When <strong>ProductVersion</strong> is set to <strong>cdc</strong>, this parameter must be set.</li><li>Get cluster ID through the API <a href="https://www.tencentcloud.com/document/product/239/109628?from_cn_redirect=1">DescribeRedisClusters</a>.</li></ul>
	RedisClusterId *string `json:"RedisClusterId,omitnil,omitempty" name:"RedisClusterId"`

	// <p>Alarm policy ID array.</p><ul><li>Log in to <a href="https://console.cloud.tencent.com/monitor/alarm/policy">Tencent Cloud Observability Platform - Alarm Management - Policy Management</a> to get alarm policy ID.</li><li>If this parameter is not configured, the default alarm policy will be bound. For the default alarm policy details, log in to <a href="https://console.cloud.tencent.com/monitor/alarm/policy">Tencent Cloud Observability Platform - Alarm Management - Policy Management</a> to view.</li></ul>
	AlarmPolicyList []*string `json:"AlarmPolicyList,omitnil,omitempty" name:"AlarmPolicyList"`

	// <p>Whether to enable password encryption for transmission.</p><ul><li>true: Encrypted.</li><li>false: Not encrypted (default value).</li></ul>
	EncryptPassword *bool `json:"EncryptPassword,omitnil,omitempty" name:"EncryptPassword"`

	// <p>Instance-level password complexity policy. When not passed in or Enabled=false, deem as not enabling policy, validate by system default rule.</p>
	PasswordPolicy *PasswordPolicy `json:"PasswordPolicy,omitnil,omitempty" name:"PasswordPolicy"`

	// <p>Whether to enable SSL encryption.</p><ul><li>true: Enable.</li><li>false: Disable (default value).</li></ul>
	EnableSSL *bool `json:"EnableSSL,omitnil,omitempty" name:"EnableSSL"`

	// <p>Whether to write the private IPv4 address of an instance to the domain alias (SAN) of the certificate when SSL is enabled. This parameter is valid only when EnableSSL is true.</p><ul><li>true: Allows using private IP to perform SSL certificate verification.</li><li>false: Does not add the SAN extended information to the certificate.</li></ul>
	SSLBindPrivateIPv4 *bool `json:"SSLBindPrivateIPv4,omitnil,omitempty" name:"SSLBindPrivateIPv4"`

	// <p>Instance connectivity access Mode.</p><ul><li>0: Proxy Mode (default value).</li><li>1: Direct access Mode.</li></ul>
	ConnectionMode *int64 `json:"ConnectionMode,omitnil,omitempty" name:"ConnectionMode"`
}

func (r *CreateInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TypeId")
	delete(f, "MemSize")
	delete(f, "GoodsNum")
	delete(f, "Period")
	delete(f, "BillingMode")
	delete(f, "ZoneId")
	delete(f, "Password")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "ProjectId")
	delete(f, "AutoRenew")
	delete(f, "SecurityGroupIdList")
	delete(f, "VPort")
	delete(f, "RedisShardNum")
	delete(f, "RedisReplicasNum")
	delete(f, "ReplicasReadonly")
	delete(f, "InstanceName")
	delete(f, "NoAuth")
	delete(f, "NodeSet")
	delete(f, "ResourceTags")
	delete(f, "ZoneName")
	delete(f, "TemplateId")
	delete(f, "DryRun")
	delete(f, "ProductVersion")
	delete(f, "RedisClusterId")
	delete(f, "AlarmPolicyList")
	delete(f, "EncryptPassword")
	delete(f, "PasswordPolicy")
	delete(f, "EnableSSL")
	delete(f, "SSLBindPrivateIPv4")
	delete(f, "ConnectionMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstancesResponseParams struct {
	// <p>Transaction ID.</p>
	//
	// Deprecated: DealId is deprecated.
	DealId *string `json:"DealId,omitnil,omitempty" name:"DealId"`

	// <p>Instance ID.</p>
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// <p>Order ID.</p>
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateInstancesResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstancesResponseParams `json:"Response"`
}

func (r *CreateInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateParamTemplateRequestParams struct {
	// Parameter template name, which can contain [2, 64] characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Parameter template description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Product type.
	// - 6: Redis 4.0 Memory Edition (standard architecture).
	// - 7: Redis 4.0 Memory Edition (cluster architecture).
	// - 8: Redis 5.0 Memory Edition (standard architecture).
	// - 9: Redis 5.0 Memory Edition (cluster architecture).
	// - 15: Redis 6.2 Memory Edition (standard architecture).
	// - 16: Redis 6.2 Memory Edition (cluster architecture).
	// - 17: Redis 7.0 Memory Edition (standard architecture).
	// - 18: Redis 7.0 Memory Edition (cluster architecture).
	ProductType *uint64 `json:"ProductType,omitnil,omitempty" name:"ProductType"`

	// ID of the source parameter template.
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Parameter list.
	ParamList []*InstanceParam `json:"ParamList,omitnil,omitempty" name:"ParamList"`
}

type CreateParamTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Parameter template name, which can contain [2, 64] characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Parameter template description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Product type.
	// - 6: Redis 4.0 Memory Edition (standard architecture).
	// - 7: Redis 4.0 Memory Edition (cluster architecture).
	// - 8: Redis 5.0 Memory Edition (standard architecture).
	// - 9: Redis 5.0 Memory Edition (cluster architecture).
	// - 15: Redis 6.2 Memory Edition (standard architecture).
	// - 16: Redis 6.2 Memory Edition (cluster architecture).
	// - 17: Redis 7.0 Memory Edition (standard architecture).
	// - 18: Redis 7.0 Memory Edition (cluster architecture).
	ProductType *uint64 `json:"ProductType,omitnil,omitempty" name:"ProductType"`

	// ID of the source parameter template.
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Parameter list.
	ParamList []*InstanceParam `json:"ParamList,omitnil,omitempty" name:"ParamList"`
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
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "ProductType")
	delete(f, "TemplateId")
	delete(f, "ParamList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateParamTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateParamTemplateResponseParams struct {
	// Parameter template ID.
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
type CreateReplicationGroupRequestParams struct {
	// ID of the primary instance in the replication group. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Replication group name. The name should contain 2 to 64 characters, including only letters, digits, underscores (_), and hyphens (-).
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Remark information.
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type CreateReplicationGroupRequest struct {
	*tchttp.BaseRequest
	
	// ID of the primary instance in the replication group. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Replication group name. The name should contain 2 to 64 characters, including only letters, digits, underscores (_), and hyphens (-).
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Remark information.
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *CreateReplicationGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReplicationGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "GroupName")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateReplicationGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReplicationGroupResponseParams struct {
	// Asynchronous process ID.
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Replication group ID of the string type.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateReplicationGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateReplicationGroupResponseParams `json:"Response"`
}

func (r *CreateReplicationGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReplicationGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DelayDistribution struct {
	// The delay distribution. The mapping between delay range and `Ladder` value is as follows:  - `1`: [0ms,1ms]. - `5`: [1ms,5ms]. - `10`: [5ms,10ms]. - `50`: [10ms,50ms]. - `200`:  [50ms,200ms]. - `-1`: [200ms,∞].
	Ladder *int64 `json:"Ladder,omitnil,omitempty" name:"Ladder"`

	// The number of commands with delay falling within the current delay range -
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// Modification time
	Updatetime *int64 `json:"Updatetime,omitnil,omitempty" name:"Updatetime"`
}

// Predefined struct for user
type DeleteExportTaskRequestParams struct {
	// <p>Log type to specify deletion.</p><p>Enumeration value:</p><ul><li>auditLog: Audit log.</li></ul>
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// <p>Specify deletion of the log filename.</p>
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// <p>Specify the instance ID. Example: crs-xjhsdj****. Log in to the <a href="https://console.cloud.tencent.com/redis">Redis console</a> and copy the instance ID from the instance list.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DeleteExportTaskRequest struct {
	*tchttp.BaseRequest
	
	// <p>Log type to specify deletion.</p><p>Enumeration value:</p><ul><li>auditLog: Audit log.</li></ul>
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// <p>Specify deletion of the log filename.</p>
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// <p>Specify the instance ID. Example: crs-xjhsdj****. Log in to the <a href="https://console.cloud.tencent.com/redis">Redis console</a> and copy the instance ID from the instance list.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DeleteExportTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteExportTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogType")
	delete(f, "FileName")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteExportTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteExportTaskResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteExportTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteExportTaskResponseParams `json:"Response"`
}

func (r *DeleteExportTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteExportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstanceAccountRequestParams struct {
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Sub-account name. Log in to the [Redis console](https://console.tencentcloud.com/redis) and switch to the **Account Management** page to obtain it. For details, see [Managing Account](https://www.tencentcloud.com/document/product/239/34590).
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`
}

type DeleteInstanceAccountRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Sub-account name. Log in to the [Redis console](https://console.tencentcloud.com/redis) and switch to the **Account Management** page to obtain it. For details, see [Managing Account](https://www.tencentcloud.com/document/product/239/34590).
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`
}

func (r *DeleteInstanceAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstanceAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AccountName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteInstanceAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstanceAccountResponseParams struct {
	// Task ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteInstanceAccountResponse struct {
	*tchttp.BaseResponse
	Response *DeleteInstanceAccountResponseParams `json:"Response"`
}

func (r *DeleteInstanceAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstanceAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteParamTemplateRequestParams struct {
	// Parameter template ID. Log in to the [Redis console and go to the parameter template page](https://console.cloud.tencent.com/redis/templates) to obtain the template ID.
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DeleteParamTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Parameter template ID. Log in to the [Redis console and go to the parameter template page](https://console.cloud.tencent.com/redis/templates) to obtain the template ID.
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
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
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
type DeleteReplicationInstanceRequestParams struct {
	// Replication group ID. Log in to the [global replication](https://console.tencentcloud.com/redis/replication) page of the Redis console and obtain it.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Data synchronization type.
	// - true: Strong synchronization is required.
	// - false: Strong synchronization is not required, and only the primary instance can be deleted.
	SyncType *bool `json:"SyncType,omitnil,omitempty" name:"SyncType"`
}

type DeleteReplicationInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Replication group ID. Log in to the [global replication](https://console.tencentcloud.com/redis/replication) page of the Redis console and obtain it.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Data synchronization type.
	// - true: Strong synchronization is required.
	// - false: Strong synchronization is not required, and only the primary instance can be deleted.
	SyncType *bool `json:"SyncType,omitnil,omitempty" name:"SyncType"`
}

func (r *DeleteReplicationInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReplicationInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "InstanceId")
	delete(f, "SyncType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteReplicationInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReplicationInstanceResponseParams struct {
	// Asynchronous task ID.
	TaskId *float64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteReplicationInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteReplicationInstanceResponseParams `json:"Response"`
}

func (r *DeleteReplicationInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReplicationInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeliverSummary struct {
	// <p>Delivery Type, store (storage class), mq (message channel)</p>
	DeliverType *string `json:"DeliverType,omitnil,omitempty" name:"DeliverType"`

	// <p>Delivery subtype: cls, ckafka.</p>
	DeliverSubType *string `json:"DeliverSubType,omitnil,omitempty" name:"DeliverSubType"`

	// <p>Subscriber</p>
	DeliverConsumer *string `json:"DeliverConsumer,omitnil,omitempty" name:"DeliverConsumer"`

	// <p>Subscriber name</p>
	DeliverConsumerName *string `json:"DeliverConsumerName,omitnil,omitempty" name:"DeliverConsumerName"`

	// <p>Delivery</p>
	DeliverError *string `json:"DeliverError,omitnil,omitempty" name:"DeliverError"`
}

// Predefined struct for user
type DescribeAutoBackupConfigRequestParams struct {
	// ID of the specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeAutoBackupConfigRequest struct {
	*tchttp.BaseRequest
	
	// ID of the specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeAutoBackupConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoBackupConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAutoBackupConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoBackupConfigResponseParams struct {
	// This parameter is retained due to compatibility and can be ignored.
	AutoBackupType *int64 `json:"AutoBackupType,omitnil,omitempty" name:"AutoBackupType"`

	// Backup cycle, which will be daily by default. Valid values: `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, `Sunday`.
	WeekDays []*string `json:"WeekDays,omitnil,omitempty" name:"WeekDays"`

	// Time period for backup task initialization
	TimePeriod *string `json:"TimePeriod,omitnil,omitempty" name:"TimePeriod"`

	// Retention time of full backup files in days.  Default value: `7`.  To retain the files for more days, [submit a ticket](https://console.cloud.tencent.com/workorder/category) for application.
	BackupStorageDays *int64 `json:"BackupStorageDays,omitnil,omitempty" name:"BackupStorageDays"`

	// This parameter has been disused.
	BinlogStorageDays *int64 `json:"BinlogStorageDays,omitnil,omitempty" name:"BinlogStorageDays"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAutoBackupConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAutoBackupConfigResponseParams `json:"Response"`
}

func (r *DescribeAutoBackupConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoBackupConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupDetailRequestParams struct {
	// <p>Instance ID. Log in to the <a href="https://console.cloud.tencent.com/redis/instance/list">Redis console</a> and copy it from the instance list.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Backup ID, which can be obtained through the response parameter <strong>RedisBackupSet</strong> of the API <a href="https://www.tencentcloud.com/document/product/239/20011?from_cn_redirect=1">DescribeInstanceBackups</a>.</p>
	BackupId *string `json:"BackupId,omitnil,omitempty" name:"BackupId"`
}

type DescribeBackupDetailRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance ID. Log in to the <a href="https://console.cloud.tencent.com/redis/instance/list">Redis console</a> and copy it from the instance list.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Backup ID, which can be obtained through the response parameter <strong>RedisBackupSet</strong> of the API <a href="https://www.tencentcloud.com/document/product/239/20011?from_cn_redirect=1">DescribeInstanceBackups</a>.</p>
	BackupId *string `json:"BackupId,omitnil,omitempty" name:"BackupId"`
}

func (r *DescribeBackupDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupDetailResponseParams struct {
	// <p>Backup ID.</p>
	BackupId *string `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// <p>Backup start time.</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>Backup end time.</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>Backup method. </p><ul><li>1: Manual backup.</li><li>0: Auto-backup.</li></ul>
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// <p>Backup status.</p><ul><li>1: Backup is locked by other processes.</li><li>2: Backup is normal, not locked by any processes.</li><li>-1: Backup has expired.</li><li>3: Backup is being exported.</li><li>4: Backup export successful.</li></ul>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Backup remarks.</p>
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// <p>Backup lock status.</p><ul><li>0: Unlocked.</li><li>1: Has been locked.</li></ul>
	Locked *int64 `json:"Locked,omitnil,omitempty" name:"Locked"`

	// <p>Backup file size. Measurement unit: Byte.</p>
	BackupSize *int64 `json:"BackupSize,omitnil,omitempty" name:"BackupSize"`

	// <p>Instance type.</p>
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// <p>Memory size of a single shard. Unit: MB.</p>
	MemSize *int64 `json:"MemSize,omitnil,omitempty" name:"MemSize"`

	// <p>Number of shards.</p>
	ShardNum *int64 `json:"ShardNum,omitnil,omitempty" name:"ShardNum"`

	// <p>Number of replicas.</p>
	ReplicasNum *int64 `json:"ReplicasNum,omitnil,omitempty" name:"ReplicasNum"`

	// <p>Whether it is encrypted or not.</p><p>Enumeration value:</p><ul><li>true: Encrypted</li><li>false: Unencrypted</li></ul>
	Encrypted *bool `json:"Encrypted,omitnil,omitempty" name:"Encrypted"`

	// <p>Decryption key.</p>
	DecryptKey *string `json:"DecryptKey,omitnil,omitempty" name:"DecryptKey"`

	// <p>Key ID of the key in KMS.</p>
	KmsKeyId *string `json:"KmsKeyId,omitnil,omitempty" name:"KmsKeyId"`

	// <p>Encryption algorithm used to encrypt the backup file.</p><p>Enumeration value:</p><ul><li>AES-256-CBC: Currently only support AES-256-CBC.</li></ul>
	KeyAlgorithm *string `json:"KeyAlgorithm,omitnil,omitempty" name:"KeyAlgorithm"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupDetailResponseParams `json:"Response"`
}

func (r *DescribeBackupDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupDownloadRestrictionRequestParams struct {

}

type DescribeBackupDownloadRestrictionRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeBackupDownloadRestrictionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDownloadRestrictionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupDownloadRestrictionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupDownloadRestrictionResponseParams struct {
	// Type of the network restrictions for downloading backup files. Valid values:
	// 
	// - `NoLimit`: Backup files can be downloaded over both public and private networks.
	// - `LimitOnlyIntranet`: Backup files can be downloaded only at private network addresses auto-assigned by Tencent Cloud.
	// - `Customize`: Backup files can be downloaded only in the customized VPC.
	LimitType *string `json:"LimitType,omitnil,omitempty" name:"LimitType"`

	// Only `In` can be passed in for this parameter, indicating that backup files can be downloaded in the custom `LimitVpc`.
	VpcComparisonSymbol *string `json:"VpcComparisonSymbol,omitnil,omitempty" name:"VpcComparisonSymbol"`

	// Whether backups can be downloaded at the custom `LimitIp` address.
	// 
	// - `In`: Download is allowed for the custom IP.
	// - `NotIn`: Download is not allowed for the custom IP.
	IpComparisonSymbol *string `json:"IpComparisonSymbol,omitnil,omitempty" name:"IpComparisonSymbol"`

	// VPC ID of the custom backup file download address, which will be displayed if `LimitType` is `Customize`.
	LimitVpc []*BackupLimitVpcItem `json:"LimitVpc,omitnil,omitempty" name:"LimitVpc"`

	// Custom VPC IP address for downloading backup files.
	//  This parameter is displayed when the **LimitType** parameter is set to **Customize**.
	LimitIp []*string `json:"LimitIp,omitnil,omitempty" name:"LimitIp"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupDownloadRestrictionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupDownloadRestrictionResponseParams `json:"Response"`
}

func (r *DescribeBackupDownloadRestrictionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDownloadRestrictionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupUrlRequestParams struct {
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup ID, which can be obtained through the `RedisBackupSet` parameter returned by the [DescribeInstanceBackups](https://intl.cloud.tencent.com/document/product/239/20011?from_cn_redirect=1) API.
	BackupId *string `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// Limit type of the network from which you can download backup files. If this parameter is not configured, the user-defined configuration will be used.
	// - NoLimit: There is no limit. Backup files can be downloaded from both Tencent Cloud private and public networks.
	// - LimitOnlyIntranet: Backup files can be downloaded 
	//  only from the private IP address automatically assigned by Tencent Cloud.
	// - Customize: Backup files can be downloaded from the user-defined VPC.
	LimitType *string `json:"LimitType,omitnil,omitempty" name:"LimitType"`

	// Only `In` can be passed in for this parameter, indicating that backup files can be downloaded in the custom `LimitVpc`.
	VpcComparisonSymbol *string `json:"VpcComparisonSymbol,omitnil,omitempty" name:"VpcComparisonSymbol"`

	// Whether backup files can be downloaded from the custom IP address specified by LimitIp.
	// - In: yes. This is the default value.
	// - NotIn: no.
	IpComparisonSymbol *string `json:"IpComparisonSymbol,omitnil,omitempty" name:"IpComparisonSymbol"`

	// VPC ID of the custom backup file download address, which is required if `LimitType` is `Customize`.
	LimitVpc []*BackupLimitVpcItem `json:"LimitVpc,omitnil,omitempty" name:"LimitVpc"`

	// VPC IP of the custom backup file download address, which is required if `LimitType` is `Customize`.
	LimitIp []*string `json:"LimitIp,omitnil,omitempty" name:"LimitIp"`
}

type DescribeBackupUrlRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup ID, which can be obtained through the `RedisBackupSet` parameter returned by the [DescribeInstanceBackups](https://intl.cloud.tencent.com/document/product/239/20011?from_cn_redirect=1) API.
	BackupId *string `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// Limit type of the network from which you can download backup files. If this parameter is not configured, the user-defined configuration will be used.
	// - NoLimit: There is no limit. Backup files can be downloaded from both Tencent Cloud private and public networks.
	// - LimitOnlyIntranet: Backup files can be downloaded 
	//  only from the private IP address automatically assigned by Tencent Cloud.
	// - Customize: Backup files can be downloaded from the user-defined VPC.
	LimitType *string `json:"LimitType,omitnil,omitempty" name:"LimitType"`

	// Only `In` can be passed in for this parameter, indicating that backup files can be downloaded in the custom `LimitVpc`.
	VpcComparisonSymbol *string `json:"VpcComparisonSymbol,omitnil,omitempty" name:"VpcComparisonSymbol"`

	// Whether backup files can be downloaded from the custom IP address specified by LimitIp.
	// - In: yes. This is the default value.
	// - NotIn: no.
	IpComparisonSymbol *string `json:"IpComparisonSymbol,omitnil,omitempty" name:"IpComparisonSymbol"`

	// VPC ID of the custom backup file download address, which is required if `LimitType` is `Customize`.
	LimitVpc []*BackupLimitVpcItem `json:"LimitVpc,omitnil,omitempty" name:"LimitVpc"`

	// VPC IP of the custom backup file download address, which is required if `LimitType` is `Customize`.
	LimitIp []*string `json:"LimitIp,omitnil,omitempty" name:"LimitIp"`
}

func (r *DescribeBackupUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupId")
	delete(f, "LimitType")
	delete(f, "VpcComparisonSymbol")
	delete(f, "IpComparisonSymbol")
	delete(f, "LimitVpc")
	delete(f, "LimitIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupUrlResponseParams struct {
	// Public network download address (valid for six hours). This field will be disused soon.
	//
	// Deprecated: DownloadUrl is deprecated.
	DownloadUrl []*string `json:"DownloadUrl,omitnil,omitempty" name:"DownloadUrl"`

	// Private network download address (valid for six hours). This field will be disused soon.
	//
	// Deprecated: InnerDownloadUrl is deprecated.
	InnerDownloadUrl []*string `json:"InnerDownloadUrl,omitnil,omitempty" name:"InnerDownloadUrl"`

	// Filename, this field is gradually being deprecated.
	//
	// Deprecated: Filenames is deprecated.
	Filenames []*string `json:"Filenames,omitnil,omitempty" name:"Filenames"`

	// Backup file information list.
	BackupInfos []*BackupDownloadInfo `json:"BackupInfos,omitnil,omitempty" name:"BackupInfos"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupUrlResponseParams `json:"Response"`
}

func (r *DescribeBackupUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBandwidthRangeRequestParams struct {
	// <p>Instance ID. Log in to the <a href="https://console.cloud.tencent.com/redis">Redis console</a> and copy it from the instance list.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeBandwidthRangeRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance ID. Log in to the <a href="https://console.cloud.tencent.com/redis">Redis console</a> and copy it from the instance list.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeBandwidthRangeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBandwidthRangeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBandwidthRangeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBandwidthRangeResponseParams struct {
	// <p>Standard bandwidth. The bandwidth allocated by the system to each node when purchasing instances.</p>Measurement unit: MB/s.
	BaseBandwidth *int64 `json:"BaseBandwidth,omitnil,omitempty" name:"BaseBandwidth"`

	// <p>Refers to the additional bandwidth of the instance. When the standard bandwidth cannot meet needs, users can manually add bandwidth.</p><ul><li>When read-only replicas are enabled, total instance bandwidth = additional bandwidth * shard quantity + standard bandwidth * shard quantity * Max ([number of read-only replicas, 1]). Shard quantity in standard architecture = 1.</li><li>When read-only replicas are not enabled, total instance bandwidth = additional bandwidth * shard quantity + standard bandwidth * shard quantity. Shard quantity in standard architecture = 1.</li></ul>Unit: MB/s.
	AddBandwidth *int64 `json:"AddBandwidth,omitnil,omitempty" name:"AddBandwidth"`

	// <p>Minimum set for additional bandwidth.</p> Unit: MB/s.
	MinAddBandwidth *int64 `json:"MinAddBandwidth,omitnil,omitempty" name:"MinAddBandwidth"`

	// <p>Set upper limit for additional bandwidth.</p> Measurement unit: MB/s.
	MaxAddBandwidth *int64 `json:"MaxAddBandwidth,omitnil,omitempty" name:"MaxAddBandwidth"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBandwidthRangeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBandwidthRangeResponseParams `json:"Response"`
}

func (r *DescribeBandwidthRangeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBandwidthRangeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCommonDBInstancesRequestParams struct {
	// List of VPC IDs
	VpcIds []*int64 `json:"VpcIds,omitnil,omitempty" name:"VpcIds"`

	// List of subnet IDs
	SubnetIds []*int64 `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// List of billing modes. 0: monthly subscription; 1: pay-as-you-go
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Instance ID filter information list, with a maximum array length of 100.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// List of instance names
	InstanceNames []*string `json:"InstanceNames,omitnil,omitempty" name:"InstanceNames"`

	// List of instance status
	Status []*string `json:"Status,omitnil,omitempty" name:"Status"`

	// Sorting field
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// List of instance VIPs
	Vips []*string `json:"Vips,omitnil,omitempty" name:"Vips"`

	// List of VPC IDs
	UniqVpcIds []*string `json:"UniqVpcIds,omitnil,omitempty" name:"UniqVpcIds"`

	// List of unique subnet IDs
	UniqSubnetIds []*string `json:"UniqSubnetIds,omitnil,omitempty" name:"UniqSubnetIds"`

	// Quantity limit. Recommended default value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeCommonDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// List of VPC IDs
	VpcIds []*int64 `json:"VpcIds,omitnil,omitempty" name:"VpcIds"`

	// List of subnet IDs
	SubnetIds []*int64 `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// List of billing modes. 0: monthly subscription; 1: pay-as-you-go
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Instance ID filter information list, with a maximum array length of 100.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// List of instance names
	InstanceNames []*string `json:"InstanceNames,omitnil,omitempty" name:"InstanceNames"`

	// List of instance status
	Status []*string `json:"Status,omitnil,omitempty" name:"Status"`

	// Sorting field
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// List of instance VIPs
	Vips []*string `json:"Vips,omitnil,omitempty" name:"Vips"`

	// List of VPC IDs
	UniqVpcIds []*string `json:"UniqVpcIds,omitnil,omitempty" name:"UniqVpcIds"`

	// List of unique subnet IDs
	UniqSubnetIds []*string `json:"UniqSubnetIds,omitnil,omitempty" name:"UniqSubnetIds"`

	// Quantity limit. Recommended default value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeCommonDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCommonDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcIds")
	delete(f, "SubnetIds")
	delete(f, "PayMode")
	delete(f, "InstanceIds")
	delete(f, "InstanceNames")
	delete(f, "Status")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "Vips")
	delete(f, "UniqVpcIds")
	delete(f, "UniqSubnetIds")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCommonDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCommonDBInstancesResponseParams struct {
	// Number of instances
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Instance information
	InstanceDetails []*RedisCommonInstanceList `json:"InstanceDetails,omitnil,omitempty" name:"InstanceDetails"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCommonDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCommonDBInstancesResponseParams `json:"Response"`
}

func (r *DescribeCommonDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCommonDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSecurityGroupsRequestParams struct {
	// Database engine name, which is `redis` for this API.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// ID of the specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDBSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Database engine name, which is `redis` for this API.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// ID of the specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
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
	delete(f, "Product")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSecurityGroupsResponseParams struct {
	// Security group rules
	Groups []*SecurityGroup `json:"Groups,omitnil,omitempty" name:"Groups"`

	// Private IPv4 address of an instance
	VIP *string `json:"VIP,omitnil,omitempty" name:"VIP"`

	// Private network port
	VPort *string `json:"VPort,omitnil,omitempty" name:"VPort"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
type DescribeExportTasksRequestParams struct {
	// <p>Log type.</p><p>Enumeration value:</p><ul><li>auditLog: Audit log.</li></ul>
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// <p>Size of the output task list per page.</p><ul><li>Default value: 20.</li><li>Value ranges from 1 to 100.</li></ul>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Pagination offset.</p><ul><li>Default value: 0.</li><li>Value: Multiple of Limit. Calculation formula: offset=limit*(page number-1).</li></ul>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Specified query instance ID. Log in to the <a href="https://console.cloud.tencent.com/redis">Redis console</a> and copy the instance ID from the instance list.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeExportTasksRequest struct {
	*tchttp.BaseRequest
	
	// <p>Log type.</p><p>Enumeration value:</p><ul><li>auditLog: Audit log.</li></ul>
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// <p>Size of the output task list per page.</p><ul><li>Default value: 20.</li><li>Value ranges from 1 to 100.</li></ul>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Pagination offset.</p><ul><li>Default value: 0.</li><li>Value: Multiple of Limit. Calculation formula: offset=limit*(page number-1).</li></ul>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Specified query instance ID. Log in to the <a href="https://console.cloud.tencent.com/redis">Redis console</a> and copy the instance ID from the instance list.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeExportTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExportTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogType")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExportTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExportTasksResponseParams struct {
	// <p>The total number of log entries.</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>Log file attribute information, including file name, file size, download link, etc.</p>
	Items []*ExportFile `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeExportTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExportTasksResponseParams `json:"Response"`
}

func (r *DescribeExportTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExportTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGlobalReplicationAreaRequestParams struct {

}

type DescribeGlobalReplicationAreaRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeGlobalReplicationAreaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGlobalReplicationAreaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGlobalReplicationAreaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGlobalReplicationAreaResponseParams struct {
	// Available region information.
	AvailableRegions []*AvailableRegion `json:"AvailableRegions,omitnil,omitempty" name:"AvailableRegions"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGlobalReplicationAreaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGlobalReplicationAreaResponseParams `json:"Response"`
}

func (r *DescribeGlobalReplicationAreaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGlobalReplicationAreaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceAccountRequestParams struct {
	// <p>Specify the instance ID. Example: crs-xjhsdj****. Log in to the <a href="https://console.cloud.tencent.com/redis">Redis console</a> and copy the instance ID from the instance list.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Page size. Default value: 20; minimum value: 1; maximum value: 100.</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Pagination offset.</p><ul><li>Parameter value: Multiple of Limit, offset=limit*(page number-1).</li><li>Default value: 0.</li></ul>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeInstanceAccountRequest struct {
	*tchttp.BaseRequest
	
	// <p>Specify the instance ID. Example: crs-xjhsdj****. Log in to the <a href="https://console.cloud.tencent.com/redis">Redis console</a> and copy the instance ID from the instance list.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Page size. Default value: 20; minimum value: 1; maximum value: 100.</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Pagination offset.</p><ul><li>Parameter value: Multiple of Limit, offset=limit*(page number-1).</li><li>Default value: 0.</li></ul>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeInstanceAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceAccountResponseParams struct {
	// <p>Account detailed information.</p>
	Accounts []*Account `json:"Accounts,omitnil,omitempty" name:"Accounts"`

	// <p>Account count.</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceAccountResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceAccountResponseParams `json:"Response"`
}

func (r *DescribeInstanceAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceBackupsRequestParams struct {
	// <p>List size of output backup per page. Default size is 20, maximum value is 100.</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Pagination offset, integer multiple of Limit. Calculation formula: offset=limit*(page number-1).</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Instance ID to be operated. You can get it from the InstanceId in the return value from the DescribeInstance API.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Start time, for example, in the format of 2017-02-08 16:46:34. Query the backup list of instances that started backup during the [beginTime, endTime] period, with a maximum query span of 30 days.</p>
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// <p>End time, in the format of 2017-02-08 19:09:26. Query the backup list of instances that started backup within the period of [beginTime, endTime]. The maximum query time span is 30 days.</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>Backup task status:<br>1: Backup is in progress.<br>2: Backup is normal.<br>3: Backup is switching to RDB file processing.<br>4: RDB switch completed.<br>-1: Backup has expired.<br>-2: Backup has been deleted.</p>
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Instance name, supports name fuzzy search based on instance name.</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

type DescribeInstanceBackupsRequest struct {
	*tchttp.BaseRequest
	
	// <p>List size of output backup per page. Default size is 20, maximum value is 100.</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Pagination offset, integer multiple of Limit. Calculation formula: offset=limit*(page number-1).</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Instance ID to be operated. You can get it from the InstanceId in the return value from the DescribeInstance API.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Start time, for example, in the format of 2017-02-08 16:46:34. Query the backup list of instances that started backup during the [beginTime, endTime] period, with a maximum query span of 30 days.</p>
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// <p>End time, in the format of 2017-02-08 19:09:26. Query the backup list of instances that started backup within the period of [beginTime, endTime]. The maximum query time span is 30 days.</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>Backup task status:<br>1: Backup is in progress.<br>2: Backup is normal.<br>3: Backup is switching to RDB file processing.<br>4: RDB switch completed.<br>-1: Backup has expired.<br>-2: Backup has been deleted.</p>
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Instance name, supports name fuzzy search based on instance name.</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

func (r *DescribeInstanceBackupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceBackupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "InstanceId")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "Status")
	delete(f, "InstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceBackupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceBackupsResponseParams struct {
	// <p>Total number of backups.</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>Backup array of the instance.</p>
	BackupSet []*RedisBackupSet `json:"BackupSet,omitnil,omitempty" name:"BackupSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceBackupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceBackupsResponseParams `json:"Response"`
}

func (r *DescribeInstanceBackupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceBackupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceDTSInfoRequestParams struct {
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceDTSInfoRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceDTSInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceDTSInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceDTSInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceDTSInfoResponseParams struct {
	// DTS task ID.
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// DTS task name.
	JobName *string `json:"JobName,omitnil,omitempty" name:"JobName"`

	// Task status. 1: creating (Creating); 3: checking (Checking); 4: check successful (CheckPass); 5: check failed (CheckNotPass); 7: task running (Running); 8: preparation completed (ReadyComplete); 9: task successful (Success); 10: task failed (Failed); 11: stopping (Stopping); 12: completing (Completing).
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Status description.
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// Synchronization delay. Unit: bytes.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Disconnection time.
	CutDownTime *string `json:"CutDownTime,omitnil,omitempty" name:"CutDownTime"`

	// Source instance information.
	SrcInfo *DescribeInstanceDTSInstanceInfo `json:"SrcInfo,omitnil,omitempty" name:"SrcInfo"`

	// Destination instance information.
	DstInfo *DescribeInstanceDTSInstanceInfo `json:"DstInfo,omitnil,omitempty" name:"DstInfo"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceDTSInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceDTSInfoResponseParams `json:"Response"`
}

func (r *DescribeInstanceDTSInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceDTSInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceDTSInstanceInfo struct {
	// Region ID.
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Repository ID.
	SetId *int64 `json:"SetId,omitnil,omitempty" name:"SetId"`

	// AZ ID.
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Instance type.
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Instance name.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Instance access address.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Status.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type DescribeInstanceDealDetailRequestParams struct {
	// Order number, which is the output parameter DealId of [CreateInstances](https://www.tencentcloud.com/document/api/239/20026?from_cn_redirect=1), with the maximum array length of 10.
	//
	// Deprecated: DealIds is deprecated.
	DealIds []*string `json:"DealIds,omitnil,omitempty" name:"DealIds"`

	// Order number, which is the output parameter DealName of [CreateInstances](https://www.tencentcloud.com/document/api/239/20026?from_cn_redirect=1), with the maximum array length of 10.
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`
}

type DescribeInstanceDealDetailRequest struct {
	*tchttp.BaseRequest
	
	// Order number, which is the output parameter DealId of [CreateInstances](https://www.tencentcloud.com/document/api/239/20026?from_cn_redirect=1), with the maximum array length of 10.
	DealIds []*string `json:"DealIds,omitnil,omitempty" name:"DealIds"`

	// Order number, which is the output parameter DealName of [CreateInstances](https://www.tencentcloud.com/document/api/239/20026?from_cn_redirect=1), with the maximum array length of 10.
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`
}

func (r *DescribeInstanceDealDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceDealDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DealIds")
	delete(f, "DealName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceDealDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceDealDetailResponseParams struct {
	// Order details
	DealDetails []*TradeDealDetail `json:"DealDetails,omitnil,omitempty" name:"DealDetails"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceDealDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceDealDetailResponseParams `json:"Response"`
}

func (r *DescribeInstanceDealDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceDealDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceEventsRequestParams struct {
	// Start date for querying the event execution plan, with a maximum query span of 30 days.
	ExecutionStartDate *string `json:"ExecutionStartDate,omitnil,omitempty" name:"ExecutionStartDate"`

	// End date for querying the event execution plan, with a maximum query span of 30 days.
	ExecutionEndDate *string `json:"ExecutionEndDate,omitnil,omitempty" name:"ExecutionEndDate"`

	// Specifies the instance ID. Example: crs-xjhsdj****. Log in to the [TencentDB for Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Outputs the number of events displayed per page.
	// - Default value: 10.
	// - Value range: [1, 100].
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Configures the output page number for querying events. You can query events on a certain page by specifying PageNo (page number) and PageSize (number of output results per page).
	// - Default value: 1.
	// - Value range: positive integers greater than 0.
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// Current status of the event.- Waiting: The event is waiting for execution on the execution date or during the operations period.- Running: The event is being executed during the operations period.- Finished: Execution of the event is completed.- Canceled: Execution of the event is canceled.
	Status []*string `json:"Status,omitnil,omitempty" name:"Status"`

	// Event type. Currently, the type can only be related to instance migration, resource movement, and IDC deletion. This parameter can be set only to **InstanceMigration**.
	EventTypes []*string `json:"EventTypes,omitnil,omitempty" name:"EventTypes"`

	// Configures the level of the queried event. Events are divided into Critical, High, Medium, and Low events according to the severity and urgency.- Critical- High- Medium- Low
	Grades []*string `json:"Grades,omitnil,omitempty" name:"Grades"`
}

type DescribeInstanceEventsRequest struct {
	*tchttp.BaseRequest
	
	// Start date for querying the event execution plan, with a maximum query span of 30 days.
	ExecutionStartDate *string `json:"ExecutionStartDate,omitnil,omitempty" name:"ExecutionStartDate"`

	// End date for querying the event execution plan, with a maximum query span of 30 days.
	ExecutionEndDate *string `json:"ExecutionEndDate,omitnil,omitempty" name:"ExecutionEndDate"`

	// Specifies the instance ID. Example: crs-xjhsdj****. Log in to the [TencentDB for Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Outputs the number of events displayed per page.
	// - Default value: 10.
	// - Value range: [1, 100].
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Configures the output page number for querying events. You can query events on a certain page by specifying PageNo (page number) and PageSize (number of output results per page).
	// - Default value: 1.
	// - Value range: positive integers greater than 0.
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// Current status of the event.- Waiting: The event is waiting for execution on the execution date or during the operations period.- Running: The event is being executed during the operations period.- Finished: Execution of the event is completed.- Canceled: Execution of the event is canceled.
	Status []*string `json:"Status,omitnil,omitempty" name:"Status"`

	// Event type. Currently, the type can only be related to instance migration, resource movement, and IDC deletion. This parameter can be set only to **InstanceMigration**.
	EventTypes []*string `json:"EventTypes,omitnil,omitempty" name:"EventTypes"`

	// Configures the level of the queried event. Events are divided into Critical, High, Medium, and Low events according to the severity and urgency.- Critical- High- Medium- Low
	Grades []*string `json:"Grades,omitnil,omitempty" name:"Grades"`
}

func (r *DescribeInstanceEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ExecutionStartDate")
	delete(f, "ExecutionEndDate")
	delete(f, "InstanceId")
	delete(f, "PageSize")
	delete(f, "PageNo")
	delete(f, "Status")
	delete(f, "EventTypes")
	delete(f, "Grades")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceEventsResponseParams struct {
	// Total number of events.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Instance event information.
	RedisInstanceEvents []*RedisInstanceEvent `json:"RedisInstanceEvents,omitnil,omitempty" name:"RedisInstanceEvents"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceEventsResponseParams `json:"Response"`
}

func (r *DescribeInstanceEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceLogDeliveryRequestParams struct {
	// Instance ID. Log in to the [Redis console](https://console.cloud.tencent.com/redis/instance), and copy it from the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceLogDeliveryRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. Log in to the [Redis console](https://console.cloud.tencent.com/redis/instance), and copy it from the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceLogDeliveryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceLogDeliveryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceLogDeliveryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceLogDeliveryResponseParams struct {
	// Slow query log shipping information of the instance.
	SlowLog *LogDeliveryInfo `json:"SlowLog,omitnil,omitempty" name:"SlowLog"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceLogDeliveryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceLogDeliveryResponseParams `json:"Response"`
}

func (r *DescribeInstanceLogDeliveryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceLogDeliveryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorBigKeyRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Request type. 1: string type; 2: all types
	ReqType *int64 `json:"ReqType,omitnil,omitempty" name:"ReqType"`

	// Time, such as "20190219"
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`
}

type DescribeInstanceMonitorBigKeyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Request type. 1: string type; 2: all types
	ReqType *int64 `json:"ReqType,omitnil,omitempty" name:"ReqType"`

	// Time, such as "20190219"
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`
}

func (r *DescribeInstanceMonitorBigKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorBigKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ReqType")
	delete(f, "Date")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceMonitorBigKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorBigKeyResponseParams struct {
	// Big key details
	Data []*BigKeyInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceMonitorBigKeyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceMonitorBigKeyResponseParams `json:"Response"`
}

func (r *DescribeInstanceMonitorBigKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorBigKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorBigKeySizeDistRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Time, such as "20190219"
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`
}

type DescribeInstanceMonitorBigKeySizeDistRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Time, such as "20190219"
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`
}

func (r *DescribeInstanceMonitorBigKeySizeDistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorBigKeySizeDistRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Date")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceMonitorBigKeySizeDistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorBigKeySizeDistResponseParams struct {
	// Big key size distribution details
	Data []*DelayDistribution `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceMonitorBigKeySizeDistResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceMonitorBigKeySizeDistResponseParams `json:"Response"`
}

func (r *DescribeInstanceMonitorBigKeySizeDistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorBigKeySizeDistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorBigKeyTypeDistRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Time, such as "20190219"
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`
}

type DescribeInstanceMonitorBigKeyTypeDistRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Time, such as "20190219"
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`
}

func (r *DescribeInstanceMonitorBigKeyTypeDistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorBigKeyTypeDistRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Date")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceMonitorBigKeyTypeDistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorBigKeyTypeDistResponseParams struct {
	// Big key type distribution details
	Data []*BigKeyTypeInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceMonitorBigKeyTypeDistResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceMonitorBigKeyTypeDistResponseParams `json:"Response"`
}

func (r *DescribeInstanceMonitorBigKeyTypeDistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorBigKeyTypeDistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorHotKeyRequestParams struct {
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Query time range.
	// - 1: real-time.
	// - 2: last 30 minutes.
	// - 3: last 6 hours.
	// - 4: last 24 hours.
	SpanType *int64 `json:"SpanType,omitnil,omitempty" name:"SpanType"`
}

type DescribeInstanceMonitorHotKeyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Query time range.
	// - 1: real-time.
	// - 2: last 30 minutes.
	// - 3: last 6 hours.
	// - 4: last 24 hours.
	SpanType *int64 `json:"SpanType,omitnil,omitempty" name:"SpanType"`
}

func (r *DescribeInstanceMonitorHotKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorHotKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SpanType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceMonitorHotKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorHotKeyResponseParams struct {
	// Hot key details.
	Data []*HotKeyInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceMonitorHotKeyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceMonitorHotKeyResponseParams `json:"Response"`
}

func (r *DescribeInstanceMonitorHotKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorHotKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorSIPRequestParams struct {
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceMonitorSIPRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceMonitorSIPRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorSIPRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceMonitorSIPRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorSIPResponseParams struct {
	// Access source information
	Data []*SourceInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceMonitorSIPResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceMonitorSIPResponseParams `json:"Response"`
}

func (r *DescribeInstanceMonitorSIPResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorSIPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorTookDistRequestParams struct {
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Query date.
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// Time range.
	// - 1: real-time.
	// - 2: last 30 minutes.
	// - 3: last 6 hours.
	// - 4: last 24 hours.
	SpanType *int64 `json:"SpanType,omitnil,omitempty" name:"SpanType"`
}

type DescribeInstanceMonitorTookDistRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Query date.
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// Time range.
	// - 1: real-time.
	// - 2: last 30 minutes.
	// - 3: last 6 hours.
	// - 4: last 24 hours.
	SpanType *int64 `json:"SpanType,omitnil,omitempty" name:"SpanType"`
}

func (r *DescribeInstanceMonitorTookDistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorTookDistRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Date")
	delete(f, "SpanType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceMonitorTookDistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorTookDistResponseParams struct {
	// Latency distribution information.
	Data []*DelayDistribution `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceMonitorTookDistResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceMonitorTookDistResponseParams `json:"Response"`
}

func (r *DescribeInstanceMonitorTookDistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorTookDistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorTopNCmdRequestParams struct {
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Time range.
	// - 1: real-time.
	// - 2: last 30 minutes.
	// - 3: last 6 hours.
	// - 4: last 24 hours.
	SpanType *int64 `json:"SpanType,omitnil,omitempty" name:"SpanType"`
}

type DescribeInstanceMonitorTopNCmdRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Time range.
	// - 1: real-time.
	// - 2: last 30 minutes.
	// - 3: last 6 hours.
	// - 4: last 24 hours.
	SpanType *int64 `json:"SpanType,omitnil,omitempty" name:"SpanType"`
}

func (r *DescribeInstanceMonitorTopNCmdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorTopNCmdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SpanType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceMonitorTopNCmdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorTopNCmdResponseParams struct {
	// Access command information
	Data []*SourceCommand `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceMonitorTopNCmdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceMonitorTopNCmdResponseParams `json:"Response"`
}

func (r *DescribeInstanceMonitorTopNCmdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorTopNCmdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorTopNCmdTookRequestParams struct {
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Query time range.
	// - 1: real-time.
	// - 2: last 30 minutes.
	// - 3: last 6 hours.
	// - 4: last 24 hours.
	SpanType *int64 `json:"SpanType,omitnil,omitempty" name:"SpanType"`
}

type DescribeInstanceMonitorTopNCmdTookRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Query time range.
	// - 1: real-time.
	// - 2: last 30 minutes.
	// - 3: last 6 hours.
	// - 4: last 24 hours.
	SpanType *int64 `json:"SpanType,omitnil,omitempty" name:"SpanType"`
}

func (r *DescribeInstanceMonitorTopNCmdTookRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorTopNCmdTookRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SpanType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceMonitorTopNCmdTookRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorTopNCmdTookResponseParams struct {
	// Duration details
	Data []*CommandTake `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceMonitorTopNCmdTookResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceMonitorTopNCmdTookResponseParams `json:"Response"`
}

func (r *DescribeInstanceMonitorTopNCmdTookResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorTopNCmdTookResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceNodeInfoRequestParams struct {
	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// List size Size of node information returned per page.  Default value: `20`. Maximum value: `1000`.  This field has been disused.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset, which is an integral multiple of `Limit`. Calculation formula:  `offset` = `limit` * (page number - 1). This field has been disused.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeInstanceNodeInfoRequest struct {
	*tchttp.BaseRequest
	
	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// List size Size of node information returned per page.  Default value: `20`. Maximum value: `1000`.  This field has been disused.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset, which is an integral multiple of `Limit`. Calculation formula:  `offset` = `limit` * (page number - 1). This field has been disused.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
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

// Predefined struct for user
type DescribeInstanceNodeInfoResponseParams struct {
	// The number of proxy nodes
	ProxyCount *int64 `json:"ProxyCount,omitnil,omitempty" name:"ProxyCount"`

	// Proxy node information.
	Proxy []*ProxyNodes `json:"Proxy,omitnil,omitempty" name:"Proxy"`

	// The number of Redis nodes
	RedisCount *int64 `json:"RedisCount,omitnil,omitempty" name:"RedisCount"`

	// TencentDB for Redis® node information.
	Redis []*RedisNodes `json:"Redis,omitnil,omitempty" name:"Redis"`

	// This parameter has been disused.
	TendisCount *int64 `json:"TendisCount,omitnil,omitempty" name:"TendisCount"`

	// This parameter is no longer used. Please ignore it.
	Tendis []*TendisNodes `json:"Tendis,omitnil,omitempty" name:"Tendis"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceNodeInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceNodeInfoResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeInstanceParamRecordsRequestParams struct {
	// Instance ID. Log in to the [Redis console](https://console.cloud.tencent.com/redis/instance/list), and copy it from the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Pagination size. The default value is 100, and the maximum value is 200.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset, which is an integer multiple of Limit. Calculation formula: Offset = Limit x (Page number – 1). The default value is 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeInstanceParamRecordsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. Log in to the [Redis console](https://console.cloud.tencent.com/redis/instance/list), and copy it from the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Pagination size. The default value is 100, and the maximum value is 200.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset, which is an integer multiple of Limit. Calculation formula: Offset = Limit x (Page number – 1). The default value is 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeInstanceParamRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceParamRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceParamRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceParamRecordsResponseParams struct {
	// Total number of modifications.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Information of modifications.
	InstanceParamHistory []*InstanceParamHistory `json:"InstanceParamHistory,omitnil,omitempty" name:"InstanceParamHistory"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceParamRecordsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceParamRecordsResponseParams `json:"Response"`
}

func (r *DescribeInstanceParamRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceParamRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceParamsRequestParams struct {
	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceParamsRequest struct {
	*tchttp.BaseRequest
	
	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceParamsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceParamsResponseParams struct {
	// Total number of the parameter lists
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Instance parameter in Enum type
	InstanceEnumParam []*InstanceEnumParam `json:"InstanceEnumParam,omitnil,omitempty" name:"InstanceEnumParam"`

	// Instance parameter in Integer type
	InstanceIntegerParam []*InstanceIntegerParam `json:"InstanceIntegerParam,omitnil,omitempty" name:"InstanceIntegerParam"`

	// Instance parameter in Char type
	InstanceTextParam []*InstanceTextParam `json:"InstanceTextParam,omitnil,omitempty" name:"InstanceTextParam"`

	// Instance parameter in Multi type
	InstanceMultiParam []*InstanceMultiParam `json:"InstanceMultiParam,omitnil,omitempty" name:"InstanceMultiParam"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
type DescribeInstanceSecurityGroupRequestParams struct {
	// Instance ID list. The array length limit is [0, 100]. Log in to the [Redis console](https://console.cloud.tencent.com/redis/instance/list), and copy the instance ID from the instance list.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type DescribeInstanceSecurityGroupRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID list. The array length limit is [0, 100]. Log in to the [Redis console](https://console.cloud.tencent.com/redis/instance/list), and copy the instance ID from the instance list.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *DescribeInstanceSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceSecurityGroupResponseParams struct {
	// Security group information of an instance
	InstanceSecurityGroupsDetail []*InstanceSecurityGroupDetail `json:"InstanceSecurityGroupsDetail,omitnil,omitempty" name:"InstanceSecurityGroupsDetail"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceSecurityGroupResponseParams `json:"Response"`
}

func (r *DescribeInstanceSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceShardsRequestParams struct {
	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Specifies whether to filter out secondary node information.
	// - true: filter out secondary nodes.
	// - false: filtering not required. The default value is false.
	FilterSlave *bool `json:"FilterSlave,omitnil,omitempty" name:"FilterSlave"`
}

type DescribeInstanceShardsRequest struct {
	*tchttp.BaseRequest
	
	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Specifies whether to filter out secondary node information.
	// - true: filter out secondary nodes.
	// - false: filtering not required. The default value is false.
	FilterSlave *bool `json:"FilterSlave,omitnil,omitempty" name:"FilterSlave"`
}

func (r *DescribeInstanceShardsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceShardsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "FilterSlave")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceShardsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceShardsResponseParams struct {
	// List information of the instance shards, which includes  node information, node ID, key count, used capacity, and capacity slope.
	InstanceShards []*InstanceClusterShard `json:"InstanceShards,omitnil,omitempty" name:"InstanceShards"`

	// Number of instance shard nodes
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceShardsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceShardsResponseParams `json:"Response"`
}

func (r *DescribeInstanceShardsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceShardsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceSpecBandwidthRequestParams struct {
	// <p>Specify the instance ID. Log in to the <a href="https://console.cloud.tencent.com/redis">Redis console</a> and copy the instance ID from the instance list. Meanwhile, InstanceId and specification parameters cannot be empty at the same time. Provide at least one.</p><ul><li>If only InstanceId is specified: Query the bandwidth of the current instance.</li><li>If InstanceId and at least one specification parameter (ShardSize, ShardNum, or ReplicateNum) are specified: Calculate the bandwidth after specification modification.</li><li>If partial or all specification parameters (ShardSize, ShardNum, ReplicateNum, and Type) are specified without InstanceId: Query the theoretical bandwidth based on the combined query of specifications.</li></ul>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Shard size. Unit: MB.</p>
	ShardSize *int64 `json:"ShardSize,omitnil,omitempty" name:"ShardSize"`

	// <p>Number of shards.</p>
	ShardNum *int64 `json:"ShardNum,omitnil,omitempty" name:"ShardNum"`

	// <p>Number of replication groups.</p>
	ReplicateNum *int64 `json:"ReplicateNum,omitnil,omitempty" name:"ReplicateNum"`

	// <p>Read-only weight. - 100: Enable read-only slave. - 0: Disable read-only slave.</p>
	ReadOnlyWeight *int64 `json:"ReadOnlyWeight,omitnil,omitempty" name:"ReadOnlyWeight"`

	// <p>Instance type, same as Type in <a href="https://www.tencentcloud.com/document/api/239/20026?from_cn_redirect=1">CreateInstances</a>.</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`
}

type DescribeInstanceSpecBandwidthRequest struct {
	*tchttp.BaseRequest
	
	// <p>Specify the instance ID. Log in to the <a href="https://console.cloud.tencent.com/redis">Redis console</a> and copy the instance ID from the instance list. Meanwhile, InstanceId and specification parameters cannot be empty at the same time. Provide at least one.</p><ul><li>If only InstanceId is specified: Query the bandwidth of the current instance.</li><li>If InstanceId and at least one specification parameter (ShardSize, ShardNum, or ReplicateNum) are specified: Calculate the bandwidth after specification modification.</li><li>If partial or all specification parameters (ShardSize, ShardNum, ReplicateNum, and Type) are specified without InstanceId: Query the theoretical bandwidth based on the combined query of specifications.</li></ul>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Shard size. Unit: MB.</p>
	ShardSize *int64 `json:"ShardSize,omitnil,omitempty" name:"ShardSize"`

	// <p>Number of shards.</p>
	ShardNum *int64 `json:"ShardNum,omitnil,omitempty" name:"ShardNum"`

	// <p>Number of replication groups.</p>
	ReplicateNum *int64 `json:"ReplicateNum,omitnil,omitempty" name:"ReplicateNum"`

	// <p>Read-only weight. - 100: Enable read-only slave. - 0: Disable read-only slave.</p>
	ReadOnlyWeight *int64 `json:"ReadOnlyWeight,omitnil,omitempty" name:"ReadOnlyWeight"`

	// <p>Instance type, same as Type in <a href="https://www.tencentcloud.com/document/api/239/20026?from_cn_redirect=1">CreateInstances</a>.</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *DescribeInstanceSpecBandwidthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceSpecBandwidthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ShardSize")
	delete(f, "ShardNum")
	delete(f, "ReplicateNum")
	delete(f, "ReadOnlyWeight")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceSpecBandwidthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceSpecBandwidthResponseParams struct {
	// <p>Basic bandwidth.</p>
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// <p>Connection limit.</p>
	ClientLimit *int64 `json:"ClientLimit,omitnil,omitempty" name:"ClientLimit"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceSpecBandwidthResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceSpecBandwidthResponseParams `json:"Response"`
}

func (r *DescribeInstanceSpecBandwidthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceSpecBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceSupportFeatureRequestParams struct {
	// Specifies the instance ID. Log in to the [Redis console](https://console.cloud.tencent.com/redis#/), and copy it from the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The features that support queries are as follows.
	// - read-local-node-only: nearby access.
	// - multi-account: multi-account management.
	// - auto-failback: fault recovery scenario, such as whether automatic failback is enabled for the primary node.
	FeatureName *string `json:"FeatureName,omitnil,omitempty" name:"FeatureName"`
}

type DescribeInstanceSupportFeatureRequest struct {
	*tchttp.BaseRequest
	
	// Specifies the instance ID. Log in to the [Redis console](https://console.cloud.tencent.com/redis#/), and copy it from the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The features that support queries are as follows.
	// - read-local-node-only: nearby access.
	// - multi-account: multi-account management.
	// - auto-failback: fault recovery scenario, such as whether automatic failback is enabled for the primary node.
	FeatureName *string `json:"FeatureName,omitnil,omitempty" name:"FeatureName"`
}

func (r *DescribeInstanceSupportFeatureRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceSupportFeatureRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "FeatureName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceSupportFeatureRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceSupportFeatureResponseParams struct {
	// Whether to support.
	Support *bool `json:"Support,omitnil,omitempty" name:"Support"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceSupportFeatureResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceSupportFeatureResponseParams `json:"Response"`
}

func (r *DescribeInstanceSupportFeatureResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceSupportFeatureResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceZoneInfoRequestParams struct {
	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceZoneInfoRequest struct {
	*tchttp.BaseRequest
	
	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceZoneInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceZoneInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceZoneInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceZoneInfoResponseParams struct {
	// Number of instance node groups
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of instance node groups
	ReplicaGroups []*ReplicaGroup `json:"ReplicaGroups,omitnil,omitempty" name:"ReplicaGroups"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceZoneInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceZoneInfoResponseParams `json:"Response"`
}

func (r *DescribeInstanceZoneInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceZoneInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesRequestParams struct {
	// Number of instances returned per page. Default value: `20`. Maximum value: `1000`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset, which is an integral multiple of `Limit`. Calculation formula:  `offset` = `limit` * (page number - 1).
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	// 
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The instance list is sorted according to the following enumeration valid values:
	// - projectId: By project ID.- createtime: By the creation time of instances.- instancename: By the name of instances.- type: By the type of instances.- curDeadline: By the expiration time of instances.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// For instance sorting order, the default is descending order.
	// - 1: Descending order.
	// - 0: Ascending order.
	OrderType *int64 `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// Array of VPC IDs such as 47525. If this parameter is not passed in or the array is empty, the classic network will be selected by default. This parameter is retained and can be ignored. It is set based on `UniqVpcIds` parameter format.
	VpcIds []*string `json:"VpcIds,omitnil,omitempty" name:"VpcIds"`

	// Array of VPC subnet IDs such as 56854. This parameter is retained and can be ignored. It is set based on `UniqSubnetIds` parameter format.
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// Setting keywords field for fuzzy query, only instance names support fuzzy query.
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// Array of project IDs
	ProjectIds []*int64 `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Array of VPC IDs such as vpc-sad23jfdfk. If this parameter is not passed in or or the array is empty, the classic network will be selected by default.
	UniqVpcIds []*string `json:"UniqVpcIds,omitnil,omitempty" name:"UniqVpcIds"`

	// Array of VPC subnet IDs such as subnet-fdj24n34j2
	UniqSubnetIds []*string `json:"UniqSubnetIds,omitnil,omitempty" name:"UniqSubnetIds"`

	// Array of region IDs (disused). The corresponding region can be queried through the common parameter `Region`.
	RegionIds []*int64 `json:"RegionIds,omitnil,omitempty" name:"RegionIds"`

	// Instance status.
	// - 0: To be initialized.
	// - 1: In process.
	// - 2: Running.
	// - -2: Isolated.
	// - -3: Pending Delete.
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Instance architecture version.
	// - 1: Single-node edition.
	// - 2: Master-replica edition.- 3: Cluster edition.
	TypeVersion *int64 `json:"TypeVersion,omitnil,omitempty" name:"TypeVersion"`

	// Storage engine information. Valid values: `Redis-2.8`, `Redis-4.0`, `Redis-5.0`, `Redis-6.0` or `CKV`.
	EngineName *string `json:"EngineName,omitnil,omitempty" name:"EngineName"`

	// Renewal pattern.
	// - 0: Manual renewal.
	// - 1: Automatic renewal.
	// - 2: No renewal after expiry.
	AutoRenew []*int64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// Billing mode.
	// - postpaid: pay-as-you-go.
	// - prepaid: monthly subscription.
	BillingMode *string `json:"BillingMode,omitnil,omitempty" name:"BillingMode"`

	// Instance type.
	// 
	// - 2: Redis 2.8 memory edition (standard architecture).
	// - 3: CKV 3.2 memory edition (standard architecture).
	// - 4: CKV 3.2 memory edition (cluster architecture).
	// - 5: Redis 2.8 memory edition (standalone).
	// - 6: Redis 4.0 memory edition (standard architecture).
	// - 7: Redis 4.0 memory edition (cluster architecture).
	// - 8: Redis 5.0 memory edition (standard architecture).
	// - 9: Redis 5.0 memory edition (cluster architecture).
	// - 15: Redis 6.2 memory edition (standard architecture).
	// - 16: Redis 6.2 memory edition (cluster architecture).
	// - 17: Redis 7.0 memory edition (standard architecture).
	// - 18: Redis 7.0 memory edition (cluster architecture).
	// - 200: Memcached 1.6 memory edition (cluster architecture).
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// This parameter is of array type and supports the configuration of instance names, instance IDs, and IP addresses. Among these, the instance name is fuzzy matching while the instance ID and IP address are precise matching.
	// - Each element in the array is used for a union-based matching query.- When both **InstanceId** and **SearchKeys** are configured simultaneously, their intersection will be used for the matching query.
	SearchKeys []*string `json:"SearchKeys,omitnil,omitempty" name:"SearchKeys"`

	// Internal parameter, which can be ignored.
	TypeList []*int64 `json:"TypeList,omitnil,omitempty" name:"TypeList"`

	// Internal parameter, which can be ignored.
	MonitorVersion *string `json:"MonitorVersion,omitnil,omitempty" name:"MonitorVersion"`

	// Resources filter by tag key and value. If this parameter is not specified or is an empty array, resources will not be filtered.
	InstanceTags []*InstanceTagInfo `json:"InstanceTags,omitnil,omitempty" name:"InstanceTags"`

	// Resources filter by tag key. If this parameter is not specified or is an empty array, resources will not be filtered.
	TagKeys []*string `json:"TagKeys,omitnil,omitempty" name:"TagKeys"`

	// The product version of the instance. If this parameter is not configured or the array is set to empty, instances will not be filtered based on this parameter by default.
	// - local: local Disk Edition.- cdc: Cluster dedicated edition.
	ProductVersions []*string `json:"ProductVersions,omitnil,omitempty" name:"ProductVersions"`

	// Batch query of the specified instances ID. The number of results returned is based on `Limit`.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Availability zone mode.
	// - singleaz: Single availability zone.- multiaz: Multiple availability zones.
	AzMode *string `json:"AzMode,omitnil,omitempty" name:"AzMode"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Number of instances returned per page. Default value: `20`. Maximum value: `1000`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset, which is an integral multiple of `Limit`. Calculation formula:  `offset` = `limit` * (page number - 1).
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	// 
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The instance list is sorted according to the following enumeration valid values:
	// - projectId: By project ID.- createtime: By the creation time of instances.- instancename: By the name of instances.- type: By the type of instances.- curDeadline: By the expiration time of instances.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// For instance sorting order, the default is descending order.
	// - 1: Descending order.
	// - 0: Ascending order.
	OrderType *int64 `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// Array of VPC IDs such as 47525. If this parameter is not passed in or the array is empty, the classic network will be selected by default. This parameter is retained and can be ignored. It is set based on `UniqVpcIds` parameter format.
	VpcIds []*string `json:"VpcIds,omitnil,omitempty" name:"VpcIds"`

	// Array of VPC subnet IDs such as 56854. This parameter is retained and can be ignored. It is set based on `UniqSubnetIds` parameter format.
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// Setting keywords field for fuzzy query, only instance names support fuzzy query.
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// Array of project IDs
	ProjectIds []*int64 `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Array of VPC IDs such as vpc-sad23jfdfk. If this parameter is not passed in or or the array is empty, the classic network will be selected by default.
	UniqVpcIds []*string `json:"UniqVpcIds,omitnil,omitempty" name:"UniqVpcIds"`

	// Array of VPC subnet IDs such as subnet-fdj24n34j2
	UniqSubnetIds []*string `json:"UniqSubnetIds,omitnil,omitempty" name:"UniqSubnetIds"`

	// Array of region IDs (disused). The corresponding region can be queried through the common parameter `Region`.
	RegionIds []*int64 `json:"RegionIds,omitnil,omitempty" name:"RegionIds"`

	// Instance status.
	// - 0: To be initialized.
	// - 1: In process.
	// - 2: Running.
	// - -2: Isolated.
	// - -3: Pending Delete.
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Instance architecture version.
	// - 1: Single-node edition.
	// - 2: Master-replica edition.- 3: Cluster edition.
	TypeVersion *int64 `json:"TypeVersion,omitnil,omitempty" name:"TypeVersion"`

	// Storage engine information. Valid values: `Redis-2.8`, `Redis-4.0`, `Redis-5.0`, `Redis-6.0` or `CKV`.
	EngineName *string `json:"EngineName,omitnil,omitempty" name:"EngineName"`

	// Renewal pattern.
	// - 0: Manual renewal.
	// - 1: Automatic renewal.
	// - 2: No renewal after expiry.
	AutoRenew []*int64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// Billing mode.
	// - postpaid: pay-as-you-go.
	// - prepaid: monthly subscription.
	BillingMode *string `json:"BillingMode,omitnil,omitempty" name:"BillingMode"`

	// Instance type.
	// 
	// - 2: Redis 2.8 memory edition (standard architecture).
	// - 3: CKV 3.2 memory edition (standard architecture).
	// - 4: CKV 3.2 memory edition (cluster architecture).
	// - 5: Redis 2.8 memory edition (standalone).
	// - 6: Redis 4.0 memory edition (standard architecture).
	// - 7: Redis 4.0 memory edition (cluster architecture).
	// - 8: Redis 5.0 memory edition (standard architecture).
	// - 9: Redis 5.0 memory edition (cluster architecture).
	// - 15: Redis 6.2 memory edition (standard architecture).
	// - 16: Redis 6.2 memory edition (cluster architecture).
	// - 17: Redis 7.0 memory edition (standard architecture).
	// - 18: Redis 7.0 memory edition (cluster architecture).
	// - 200: Memcached 1.6 memory edition (cluster architecture).
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// This parameter is of array type and supports the configuration of instance names, instance IDs, and IP addresses. Among these, the instance name is fuzzy matching while the instance ID and IP address are precise matching.
	// - Each element in the array is used for a union-based matching query.- When both **InstanceId** and **SearchKeys** are configured simultaneously, their intersection will be used for the matching query.
	SearchKeys []*string `json:"SearchKeys,omitnil,omitempty" name:"SearchKeys"`

	// Internal parameter, which can be ignored.
	TypeList []*int64 `json:"TypeList,omitnil,omitempty" name:"TypeList"`

	// Internal parameter, which can be ignored.
	MonitorVersion *string `json:"MonitorVersion,omitnil,omitempty" name:"MonitorVersion"`

	// Resources filter by tag key and value. If this parameter is not specified or is an empty array, resources will not be filtered.
	InstanceTags []*InstanceTagInfo `json:"InstanceTags,omitnil,omitempty" name:"InstanceTags"`

	// Resources filter by tag key. If this parameter is not specified or is an empty array, resources will not be filtered.
	TagKeys []*string `json:"TagKeys,omitnil,omitempty" name:"TagKeys"`

	// The product version of the instance. If this parameter is not configured or the array is set to empty, instances will not be filtered based on this parameter by default.
	// - local: local Disk Edition.- cdc: Cluster dedicated edition.
	ProductVersions []*string `json:"ProductVersions,omitnil,omitempty" name:"ProductVersions"`

	// Batch query of the specified instances ID. The number of results returned is based on `Limit`.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Availability zone mode.
	// - singleaz: Single availability zone.- multiaz: Multiple availability zones.
	AzMode *string `json:"AzMode,omitnil,omitempty" name:"AzMode"`
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
	delete(f, "InstanceId")
	delete(f, "OrderBy")
	delete(f, "OrderType")
	delete(f, "VpcIds")
	delete(f, "SubnetIds")
	delete(f, "SearchKey")
	delete(f, "ProjectIds")
	delete(f, "InstanceName")
	delete(f, "UniqVpcIds")
	delete(f, "UniqSubnetIds")
	delete(f, "RegionIds")
	delete(f, "Status")
	delete(f, "TypeVersion")
	delete(f, "EngineName")
	delete(f, "AutoRenew")
	delete(f, "BillingMode")
	delete(f, "Type")
	delete(f, "SearchKeys")
	delete(f, "TypeList")
	delete(f, "MonitorVersion")
	delete(f, "InstanceTags")
	delete(f, "TagKeys")
	delete(f, "ProductVersions")
	delete(f, "InstanceIds")
	delete(f, "AzMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// Total number of instances
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of instance details
	InstanceSet []*InstanceSet `json:"InstanceSet,omitnil,omitempty" name:"InstanceSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
type DescribeLogInstanceListRequestParams struct {
	// <p>Log type.</p><p>Enumeration value:</p><ul><li>auditLog: Audit log.</li></ul>
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// <p>Size of the output task list per page.</p><ul><li>Value ranges from 1 to 100.</li><li>Default value: 20.</li></ul>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Pagination offset. Default value: 0. Value: Multiple of Limit. Calculation formula: offset=limit*(page number-1).</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Set the log filtering field to filter and return logs that meet a specified condition.</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>Log subcategory.</p><p>Enumeration value:</p><ul><li>write: Write logs.</li><li>read: Read logs.</li><li>all: Read/write logs.</li></ul>
	LogSubType *string `json:"LogSubType,omitnil,omitempty" name:"LogSubType"`

	// <p>Log switch.</p><p>Enumeration value:</p><ul><li>on: Enable</li><li>off: Disable</li></ul><p>Default value: off</p>
	LogSwitch *string `json:"LogSwitch,omitnil,omitempty" name:"LogSwitch"`
}

type DescribeLogInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// <p>Log type.</p><p>Enumeration value:</p><ul><li>auditLog: Audit log.</li></ul>
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// <p>Size of the output task list per page.</p><ul><li>Value ranges from 1 to 100.</li><li>Default value: 20.</li></ul>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Pagination offset. Default value: 0. Value: Multiple of Limit. Calculation formula: offset=limit*(page number-1).</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Set the log filtering field to filter and return logs that meet a specified condition.</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>Log subcategory.</p><p>Enumeration value:</p><ul><li>write: Write logs.</li><li>read: Read logs.</li><li>all: Read/write logs.</li></ul>
	LogSubType *string `json:"LogSubType,omitnil,omitempty" name:"LogSubType"`

	// <p>Log switch.</p><p>Enumeration value:</p><ul><li>on: Enable</li><li>off: Disable</li></ul><p>Default value: off</p>
	LogSwitch *string `json:"LogSwitch,omitnil,omitempty" name:"LogSwitch"`
}

func (r *DescribeLogInstanceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogInstanceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogType")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "LogSubType")
	delete(f, "LogSwitch")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogInstanceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogInstanceListResponseParams struct {
	// <p>The number of queried logs.</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>Log platform instance information.</p>
	Items []*LogInstance `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLogInstanceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLogInstanceListResponseParams `json:"Response"`
}

func (r *DescribeLogInstanceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogsRequestParams struct {
	// <p>Specify the instance ID. Example: crs-xjhsdj****. Log in to the <a href="https://console.cloud.tencent.com/redis">Redis console</a> and copy the instance ID from the instance list.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Start time of retrieval.</p><p>Parameter format: YYYY-MM-DD HH:mm:ss, for example 2026-03-06 00:00:00. The returned result contains only the logs at this time point and afterward.</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>End time of log retrieval.</p><p>Parameter format: YYYY-MM-DD HH:mm:ss, for example 2026-03-06 23:59:59. The returned result contains only the logs at this time point and earlier.</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>Log type.</p><p>Enumeration value:</p><ul><li>auditLog: Audit log.</li></ul>
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// <p>Filter conditions.</p>
	LogFilter []*LogFilter `json:"LogFilter,omitnil,omitempty" name:"LogFilter"`

	// <p>List size of returned logs per page.</p><ul><li>Default value: 20.</li><li>Value ranges from 1 to 100.</li></ul>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Starting offset amount of pagination.</p><ul><li>Default: 0.</li><li>Value: Multiple of Limit. Calculation formula: offset=limit*(page number-1).</li></ul>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Log sorting method. Default value is DESC. Values are as follows:</p><ul><li>ASC: Sort in ascending order by time with the earliest log first.</li><li>DESC: Sort in descending order with the latest log first.</li></ul>
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// <p>Sorting field. Specifies the field used to sort logs.</p>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`
}

type DescribeLogsRequest struct {
	*tchttp.BaseRequest
	
	// <p>Specify the instance ID. Example: crs-xjhsdj****. Log in to the <a href="https://console.cloud.tencent.com/redis">Redis console</a> and copy the instance ID from the instance list.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Start time of retrieval.</p><p>Parameter format: YYYY-MM-DD HH:mm:ss, for example 2026-03-06 00:00:00. The returned result contains only the logs at this time point and afterward.</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>End time of log retrieval.</p><p>Parameter format: YYYY-MM-DD HH:mm:ss, for example 2026-03-06 23:59:59. The returned result contains only the logs at this time point and earlier.</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>Log type.</p><p>Enumeration value:</p><ul><li>auditLog: Audit log.</li></ul>
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// <p>Filter conditions.</p>
	LogFilter []*LogFilter `json:"LogFilter,omitnil,omitempty" name:"LogFilter"`

	// <p>List size of returned logs per page.</p><ul><li>Default value: 20.</li><li>Value ranges from 1 to 100.</li></ul>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Starting offset amount of pagination.</p><ul><li>Default: 0.</li><li>Value: Multiple of Limit. Calculation formula: offset=limit*(page number-1).</li></ul>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Log sorting method. Default value is DESC. Values are as follows:</p><ul><li>ASC: Sort in ascending order by time with the earliest log first.</li><li>DESC: Sort in descending order with the latest log first.</li></ul>
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// <p>Sorting field. Specifies the field used to sort logs.</p>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`
}

func (r *DescribeLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "LogType")
	delete(f, "LogFilter")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "OrderBy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogsResponseParams struct {
	// <p>Total log quantity of the query.</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>Log details.</p>
	Items []*LogResult `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLogsResponseParams `json:"Response"`
}

func (r *DescribeLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMaintenanceWindowRequestParams struct {
	// Specifies the instance ID. Example: crs-xjhsdj****. Log in to the [TencentDB for Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeMaintenanceWindowRequest struct {
	*tchttp.BaseRequest
	
	// Specifies the instance ID. Example: crs-xjhsdj****. Log in to the [TencentDB for Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeMaintenanceWindowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMaintenanceWindowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMaintenanceWindowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMaintenanceWindowResponseParams struct {
	// Start time of the maintenance window. Value range: any time point between 00:00 and 23:00, for example, 03:24.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time of the maintenance window.
	// - Value range: any time point between 00:00 and 23:00, for example, 04:24.
	// - The minimum maintenance duration is 30 minutes and the maximum is 3 hours.
	// - The end time should be later than the start time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMaintenanceWindowResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMaintenanceWindowResponseParams `json:"Response"`
}

func (r *DescribeMaintenanceWindowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMaintenanceWindowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeParamTemplateInfoRequestParams struct {
	// The parameter template ID for query. Get parameter template list information through the [DescribeParamTemplates](https://intl.cloud.tencent.com/document/product/239/58750?from_cn_redirect=1) API.
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DescribeParamTemplateInfoRequest struct {
	*tchttp.BaseRequest
	
	// The parameter template ID for query. Get parameter template list information through the [DescribeParamTemplates](https://intl.cloud.tencent.com/document/product/239/58750?from_cn_redirect=1) API.
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *DescribeParamTemplateInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParamTemplateInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeParamTemplateInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeParamTemplateInfoResponseParams struct {
	// Quantity of parameters in the parameter template
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Parameter template ID.
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Parameter template name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Product type.
	// - 2: Redis 2.8 Memory Edition (standard architecture).
	// - 3: CKV 3.2 Memory Edition (standard architecture).
	// - 4: CKV 3.2 Memory Edition (cluster architecture).
	// - 5: Redis 2.8 Memory Edition (stand-alone).
	// - 6: Redis 4.0 Memory Edition (standard architecture).
	// - 7: Redis 4.0 Memory Edition (cluster architecture).
	// - 8: Redis 5.0 Memory Edition (standard architecture).
	// - 9: Redis 5.0 Memory Edition (cluster architecture).
	// - 15: Redis 6.2 Memory Edition (standard architecture).
	// - 16: Redis 6.2 Memory Edition (cluster architecture).
	// - 17: Redis 7.0 Memory Edition (standard architecture).
	// - 18: Redis 7.0 Memory Edition (cluster architecture).
	ProductType *uint64 `json:"ProductType,omitnil,omitempty" name:"ProductType"`

	// Parameter template description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Parameter details, including parameter name, current value, default value, maximum value, minimum value, enumeration value and other information.
	Items []*ParameterDetail `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeParamTemplateInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeParamTemplateInfoResponseParams `json:"Response"`
}

func (r *DescribeParamTemplateInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParamTemplateInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeParamTemplatesRequestParams struct {
	// <p>Specified query for product version and architecture.</p><ul><li>6: Redis 4.0 standard architecture;</li><li>7: Redis 4.0 cluster architecture;</li><li>8: Redis 5.0 standard architecture;</li><li>9: Redis 5.0 cluster architecture;</li><li>15: Redis 6.2 standard architecture;</li><li>16: Redis 6.2 cluster architecture;</li><li>17: Redis 7.0 standard architecture;</li><li>18: Redis 7.0 cluster architecture;</li><li>19: ValKey 8.0 standard architecture;</li><li>20: ValKey 8.0 cluster architecture.</li></ul>
	ProductTypes []*int64 `json:"ProductTypes,omitnil,omitempty" name:"ProductTypes"`

	// <p>Specify the parameter template name for the query.</p><ul><li>Data type: string array, with a maximum length limit of 50.</li><li>Method for obtaining: Copy the Template name of a custom template or system default template on the <a href="https://console.cloud.tencent.com/redis/templates">parameter template page in the Redis console</a>.</li></ul>
	TemplateNames []*string `json:"TemplateNames,omitnil,omitempty" name:"TemplateNames"`

	// <p>Parameter template ID specified for query.</p><ul><li>Data type: string array, with a maximum length limit of 50.</li><li>Method for obtaining: Copy the template ID of a custom template or system default template on the <a href="https://console.cloud.tencent.com/redis/templates">parameter template page in the Redis console</a>.</li></ul>
	TemplateIds []*string `json:"TemplateIds,omitnil,omitempty" name:"TemplateIds"`

	// <p>Specify the pagination size of the query result, which is the number of records returned per page.</p><ul><li>Value ranges from 0–200.</li><li>Default value: 200.</li></ul>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Pagination offset, used to specify the starting position of the query result.</p><ul><li>Value: Must be an integral multiple of Limit. Default value is 0.</li><li>Calculation formula: offset=limit*(page number-1).</li></ul>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeParamTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// <p>Specified query for product version and architecture.</p><ul><li>6: Redis 4.0 standard architecture;</li><li>7: Redis 4.0 cluster architecture;</li><li>8: Redis 5.0 standard architecture;</li><li>9: Redis 5.0 cluster architecture;</li><li>15: Redis 6.2 standard architecture;</li><li>16: Redis 6.2 cluster architecture;</li><li>17: Redis 7.0 standard architecture;</li><li>18: Redis 7.0 cluster architecture;</li><li>19: ValKey 8.0 standard architecture;</li><li>20: ValKey 8.0 cluster architecture.</li></ul>
	ProductTypes []*int64 `json:"ProductTypes,omitnil,omitempty" name:"ProductTypes"`

	// <p>Specify the parameter template name for the query.</p><ul><li>Data type: string array, with a maximum length limit of 50.</li><li>Method for obtaining: Copy the Template name of a custom template or system default template on the <a href="https://console.cloud.tencent.com/redis/templates">parameter template page in the Redis console</a>.</li></ul>
	TemplateNames []*string `json:"TemplateNames,omitnil,omitempty" name:"TemplateNames"`

	// <p>Parameter template ID specified for query.</p><ul><li>Data type: string array, with a maximum length limit of 50.</li><li>Method for obtaining: Copy the template ID of a custom template or system default template on the <a href="https://console.cloud.tencent.com/redis/templates">parameter template page in the Redis console</a>.</li></ul>
	TemplateIds []*string `json:"TemplateIds,omitnil,omitempty" name:"TemplateIds"`

	// <p>Specify the pagination size of the query result, which is the number of records returned per page.</p><ul><li>Value ranges from 0–200.</li><li>Default value: 200.</li></ul>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Pagination offset, used to specify the starting position of the query result.</p><ul><li>Value: Must be an integral multiple of Limit. Default value is 0.</li><li>Calculation formula: offset=limit*(page number-1).</li></ul>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
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
	delete(f, "ProductTypes")
	delete(f, "TemplateNames")
	delete(f, "TemplateIds")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeParamTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeParamTemplatesResponseParams struct {
	// <p>Number of parameter templates for this user.</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>Parameter template details.</p>
	Items []*ParamTemplateInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
type DescribeProductInfoRequestParams struct {

}

type DescribeProductInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeProductInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProductInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductInfoResponseParams struct {
	// Selling information on the region. The selling information on all regions is returned even if a region is specified.
	RegionSet []*RegionConf `json:"RegionSet,omitnil,omitempty" name:"RegionSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProductInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProductInfoResponseParams `json:"Response"`
}

func (r *DescribeProductInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectSecurityGroupRequestParams struct {
	// Project ID for query.
	// - 0: default project.
	// - -1: all projects.
	// - Greater than 0: specific project. Log in to the [project management](https://console.tencentcloud.com/project) page of the Redis console and copy the project ID in **Project Name**.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Security group ID, which can be obtained through the sub-parameter **SecurityGroupId** of the response parameter **InstanceSecurityGroupsDetail** of the API [DescribeInstanceSecurityGroup](https://intl.cloud.tencent.com/document/product/239/34447?from_cn_redirect=1).
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`
}

type DescribeProjectSecurityGroupRequest struct {
	*tchttp.BaseRequest
	
	// Project ID for query.
	// - 0: default project.
	// - -1: all projects.
	// - Greater than 0: specific project. Log in to the [project management](https://console.tencentcloud.com/project) page of the Redis console and copy the project ID in **Project Name**.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Security group ID, which can be obtained through the sub-parameter **SecurityGroupId** of the response parameter **InstanceSecurityGroupsDetail** of the API [DescribeInstanceSecurityGroup](https://intl.cloud.tencent.com/document/product/239/34447?from_cn_redirect=1).
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`
}

func (r *DescribeProjectSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "SecurityGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectSecurityGroupResponseParams struct {
	// Security group of the project
	SecurityGroupDetails []*SecurityGroupDetail `json:"SecurityGroupDetails,omitnil,omitempty" name:"SecurityGroupDetails"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProjectSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProjectSecurityGroupResponseParams `json:"Response"`
}

func (r *DescribeProjectSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectSecurityGroupsRequestParams struct {
	// Database engine name, which is `redis` for this API.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// Project ID. Log in to the [Project Management](https://console.tencentcloud.com/project) page of the Redis console and copy the project ID in **Project Name**.
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Offset, which is an integral multiple of `Limit`.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of security groups to be pulled. Default value: `20`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Search criteria. You can enter a security group ID or name.
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`
}

type DescribeProjectSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Database engine name, which is `redis` for this API.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// Project ID. Log in to the [Project Management](https://console.tencentcloud.com/project) page of the Redis console and copy the project ID in **Project Name**.
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Offset, which is an integral multiple of `Limit`.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of security groups to be pulled. Default value: `20`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Search criteria. You can enter a security group ID or name.
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
	delete(f, "Product")
	delete(f, "ProjectId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectSecurityGroupsResponseParams struct {
	// Security group rules.
	Groups []*SecurityGroup `json:"Groups,omitnil,omitempty" name:"Groups"`

	// Total number of eligible security groups.
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
type DescribeProxySlowLogRequestParams struct {
	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time of a slow query, with a maximum query span of 30 days.
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// End time of a slow query, with a maximum query span of 30 days.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Slow query threshold, in milliseconds. The value is a positive integer greater than 0.
	MinQueryTime *int64 `json:"MinQueryTime,omitnil,omitempty" name:"MinQueryTime"`

	// Size of the output task list per page. The default value is 20, the minimum value is 1, and the maximum value is 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset, which is an integer multiple of Limit. Calculation formula: Offset = Limit x (Page number – 1). The default value is 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeProxySlowLogRequest struct {
	*tchttp.BaseRequest
	
	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time of a slow query, with a maximum query span of 30 days.
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// End time of a slow query, with a maximum query span of 30 days.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Slow query threshold, in milliseconds. The value is a positive integer greater than 0.
	MinQueryTime *int64 `json:"MinQueryTime,omitnil,omitempty" name:"MinQueryTime"`

	// Size of the output task list per page. The default value is 20, the minimum value is 1, and the maximum value is 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset, which is an integer multiple of Limit. Calculation formula: Offset = Limit x (Page number – 1). The default value is 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeProxySlowLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxySlowLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "MinQueryTime")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProxySlowLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProxySlowLogResponseParams struct {
	// Total number of slow queries
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Slow query details. Note: If the value of TotalCount is greater than 10,000, indicating that the number of slow logs exceeds 10,000, log details cannot be returned. Instead, the returned data is empty. It is recommended to reduce the interval between BeginTime and EndTime and perform multiple queries.
	InstanceProxySlowLogDetail []*InstanceProxySlowlogDetail `json:"InstanceProxySlowLogDetail,omitnil,omitempty" name:"InstanceProxySlowLogDetail"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProxySlowLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProxySlowLogResponseParams `json:"Response"`
}

func (r *DescribeProxySlowLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxySlowLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRedisClusterOverviewRequestParams struct {
	// CDC ID. Log in to the [CDC console](https://console.cloud.tencent.com/cdc/dedicatedcluster/index?rid=1)
	// and obtain the cluster ID in the instance list.
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`
}

type DescribeRedisClusterOverviewRequest struct {
	*tchttp.BaseRequest
	
	// CDC ID. Log in to the [CDC console](https://console.cloud.tencent.com/cdc/dedicatedcluster/index?rid=1)
	// and obtain the cluster ID in the instance list.
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`
}

func (r *DescribeRedisClusterOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRedisClusterOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DedicatedClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRedisClusterOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRedisClusterOverviewResponseParams struct {
	// Total number of resource packages.
	TotalBundle *int64 `json:"TotalBundle,omitnil,omitempty" name:"TotalBundle"`

	// Total memory size occupied by resource packages. Unit: GB.
	TotalMemory *int64 `json:"TotalMemory,omitnil,omitempty" name:"TotalMemory"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRedisClusterOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRedisClusterOverviewResponseParams `json:"Response"`
}

func (r *DescribeRedisClusterOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRedisClusterOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRedisClustersRequestParams struct {
	// Dedicated Redis cluster ID. Log in to the [CDC console](https://console.cloud.tencent.com/cdc/dedicatedcluster/index?rid=1),
	// switch to the **Cloud Service Management** page, select **TencentDB for Redis** from the drop-down list, and obtain the dedicated cluster ID.
	RedisClusterIds []*string `json:"RedisClusterIds,omitnil,omitempty" name:"RedisClusterIds"`

	// Cluster status.
	// - 1: in process.
	// - 2: running.
	// - 3: isolated.
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Project ID array. Log in to the [project management](https://console.tencentcloud.com/project) page and copy the project ID in **Project Name**.
	ProjectIds []*int64 `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// Renewal mode.
	// - 0: default status (manual renewal).
	// - 1: automatic renewal.
	// - 2: no automatic renewal.
	AutoRenewFlag []*int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// Dedicated Redis cluster name.
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// Search keyword. Valid values: cluster ID and cluster name.
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// Limit on the number of records returned in pagination mode. If this parameter is not specified, the value 20 will be used by default.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset, which is an integer multiple of Limit.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// CDC ID. Log in to the [CDC console](https://console.cloud.tencent.com/cdc/dedicatedcluster/index?rid=1)
	// and obtain the cluster ID in the instance list.
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`
}

type DescribeRedisClustersRequest struct {
	*tchttp.BaseRequest
	
	// Dedicated Redis cluster ID. Log in to the [CDC console](https://console.cloud.tencent.com/cdc/dedicatedcluster/index?rid=1),
	// switch to the **Cloud Service Management** page, select **TencentDB for Redis** from the drop-down list, and obtain the dedicated cluster ID.
	RedisClusterIds []*string `json:"RedisClusterIds,omitnil,omitempty" name:"RedisClusterIds"`

	// Cluster status.
	// - 1: in process.
	// - 2: running.
	// - 3: isolated.
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Project ID array. Log in to the [project management](https://console.tencentcloud.com/project) page and copy the project ID in **Project Name**.
	ProjectIds []*int64 `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// Renewal mode.
	// - 0: default status (manual renewal).
	// - 1: automatic renewal.
	// - 2: no automatic renewal.
	AutoRenewFlag []*int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// Dedicated Redis cluster name.
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// Search keyword. Valid values: cluster ID and cluster name.
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// Limit on the number of records returned in pagination mode. If this parameter is not specified, the value 20 will be used by default.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset, which is an integer multiple of Limit.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// CDC ID. Log in to the [CDC console](https://console.cloud.tencent.com/cdc/dedicatedcluster/index?rid=1)
	// and obtain the cluster ID in the instance list.
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`
}

func (r *DescribeRedisClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRedisClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RedisClusterIds")
	delete(f, "Status")
	delete(f, "ProjectIds")
	delete(f, "AutoRenewFlag")
	delete(f, "ClusterName")
	delete(f, "SearchKey")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "DedicatedClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRedisClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRedisClustersResponseParams struct {
	// Total number of clusters.
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// CDC cluster resource list.
	Resources []*CDCResource `json:"Resources,omitnil,omitempty" name:"Resources"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRedisClustersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRedisClustersResponseParams `json:"Response"`
}

func (r *DescribeRedisClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRedisClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReplicationGroupInstanceRequestParams struct {
	// Specifies the instance ID. Example: crs-xjhsdj****. Log in to the [TencentDB for Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeReplicationGroupInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Specifies the instance ID. Example: crs-xjhsdj****. Log in to the [TencentDB for Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeReplicationGroupInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReplicationGroupInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReplicationGroupInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReplicationGroupInstanceResponseParams struct {
	// AppID。
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Numerical code of a region.
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// String ID of a replication group.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Replication group name.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Instance replication group role.
	// - r: secondary instance.
	// - rw: primary instance.
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeReplicationGroupInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReplicationGroupInstanceResponseParams `json:"Response"`
}

func (r *DescribeReplicationGroupInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReplicationGroupInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReplicationGroupRequestParams struct {
	// Size of the output instance list per page. The value is a positive integer greater than 0, and the default value is 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset, which is an integral multiple of `Limit`. `offset` = `limit` * (page number - 1).
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// ID of the specified replication group, such as `crs-rpl-m3zt****`. Log in to the [TencentDB for Redis console](https://console.cloud.tencent.com/redis/replication), and get it in the global replication group list.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Keyword for fuzzy query, which can be a replication group ID or name. Log in to the [TencentDB for Redis console](https://console.cloud.tencent.com/redis/replication), and get them in the global replication group list.
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`
}

type DescribeReplicationGroupRequest struct {
	*tchttp.BaseRequest
	
	// Size of the output instance list per page. The value is a positive integer greater than 0, and the default value is 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset, which is an integral multiple of `Limit`. `offset` = `limit` * (page number - 1).
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// ID of the specified replication group, such as `crs-rpl-m3zt****`. Log in to the [TencentDB for Redis console](https://console.cloud.tencent.com/redis/replication), and get it in the global replication group list.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Keyword for fuzzy query, which can be a replication group ID or name. Log in to the [TencentDB for Redis console](https://console.cloud.tencent.com/redis/replication), and get them in the global replication group list.
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`
}

func (r *DescribeReplicationGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReplicationGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "GroupId")
	delete(f, "SearchKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReplicationGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReplicationGroupResponseParams struct {
	// Number of replication groups
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Replication group information
	Groups []*Groups `json:"Groups,omitnil,omitempty" name:"Groups"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeReplicationGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReplicationGroupResponseParams `json:"Response"`
}

func (r *DescribeReplicationGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReplicationGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSSLStatusRequestParams struct {
	// <p>Instance ID. Log in to the <a href="https://console.cloud.tencent.com/redis/instance/list">Redis console</a> and copy it from the instance list.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeSSLStatusRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance ID. Log in to the <a href="https://console.cloud.tencent.com/redis/instance/list">Redis console</a> and copy it from the instance list.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeSSLStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSSLStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSSLStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSSLStatusResponseParams struct {
	// <p>SSL certificate download address.</p>
	CertDownloadUrl *string `json:"CertDownloadUrl,omitnil,omitempty" name:"CertDownloadUrl"`

	// <p>Certificate download link expiration time.</p>
	UrlExpiredTime *string `json:"UrlExpiredTime,omitnil,omitempty" name:"UrlExpiredTime"`

	// <p>Flag whether the instance enables SSL feature.</p><ul><li>true: enable.</li><li>false: disable.</li></ul>
	SSLConfig *bool `json:"SSLConfig,omitnil,omitempty" name:"SSLConfig"`

	// <p>Flag whether the instance supports SSL feature.</p><ul><li>true: support.</li><li>false: unsupported.</li></ul>
	FeatureSupport *bool `json:"FeatureSupport,omitnil,omitempty" name:"FeatureSupport"`

	// <p>Describes the SSL configuration status.</p><ul><li>1: Configuration in progress.</li><li>2: Configured successfully.</li></ul>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Address type.</p><p>Enumeration value:</p><ul><li>0: Unlimited.</li><li>1: Private IPv4 address.</li><li>2: Private IPv6 address.</li><li>3: Public network.</li><li>-1: Unspecified.</li></ul>
	AddressType *int64 `json:"AddressType,omitnil,omitempty" name:"AddressType"`

	// <p>Current encrypted connection address</p>
	EncryptAddress *string `json:"EncryptAddress,omitnil,omitempty" name:"EncryptAddress"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSSLStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSSLStatusResponseParams `json:"Response"`
}

func (r *DescribeSSLStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSSLStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecondLevelBackupInfoRequestParams struct {
	// Specifies the instance ID. Log in to the [Redis console](https://console.cloud.tencent.com/redis/instance/list), and copy it from the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Second-level backup timestamp.
	// - Setting range: support any second-level time point within 7 days.
	// - Timestamp format: UNIX timestamp.
	BackupTimestamp *int64 `json:"BackupTimestamp,omitnil,omitempty" name:"BackupTimestamp"`
}

type DescribeSecondLevelBackupInfoRequest struct {
	*tchttp.BaseRequest
	
	// Specifies the instance ID. Log in to the [Redis console](https://console.cloud.tencent.com/redis/instance/list), and copy it from the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Second-level backup timestamp.
	// - Setting range: support any second-level time point within 7 days.
	// - Timestamp format: UNIX timestamp.
	BackupTimestamp *int64 `json:"BackupTimestamp,omitnil,omitempty" name:"BackupTimestamp"`
}

func (r *DescribeSecondLevelBackupInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecondLevelBackupInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupTimestamp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecondLevelBackupInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecondLevelBackupInfoResponseParams struct {
	// Backup record ID.
	BackupId *string `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// Backup timestamp.
	BackupTimestamp *int64 `json:"BackupTimestamp,omitnil,omitempty" name:"BackupTimestamp"`

	// Timestamp range within which backup is missing.
	MissingTimestamps []*SecondLevelBackupMissingTimestamps `json:"MissingTimestamps,omitnil,omitempty" name:"MissingTimestamps"`

	// Timestamp when second-level backup is enabled for the instance.
	StartTimestamp *int64 `json:"StartTimestamp,omitnil,omitempty" name:"StartTimestamp"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSecondLevelBackupInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecondLevelBackupInfoResponseParams `json:"Response"`
}

func (r *DescribeSecondLevelBackupInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecondLevelBackupInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogRequestParams struct {
	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time for pre-querying slow query logs, with a maximum query span of 30 days.
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// End time for pre-querying slow query logs, with a maximum query span of 30 days.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Average execution time threshold for slow queries, in milliseconds. The value is a positive integer greater than 0
	MinQueryTime *int64 `json:"MinQueryTime,omitnil,omitempty" name:"MinQueryTime"`

	// Number of slow query results displayed per page. The default value is 20, the minimum value is 1, and the maximum value is 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset of the number of slow queries. The value is an integer multiple of Limit. Calculation formula: Offset = Limit x (Page number – 1). The default value is 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Node role.
	// - master: Master node.- slave: Replica node.
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`
}

type DescribeSlowLogRequest struct {
	*tchttp.BaseRequest
	
	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time for pre-querying slow query logs, with a maximum query span of 30 days.
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// End time for pre-querying slow query logs, with a maximum query span of 30 days.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Average execution time threshold for slow queries, in milliseconds. The value is a positive integer greater than 0
	MinQueryTime *int64 `json:"MinQueryTime,omitnil,omitempty" name:"MinQueryTime"`

	// Number of slow query results displayed per page. The default value is 20, the minimum value is 1, and the maximum value is 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset of the number of slow queries. The value is an integer multiple of Limit. Calculation formula: Offset = Limit x (Page number – 1). The default value is 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Node role.
	// - master: Master node.- slave: Replica node.
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`
}

func (r *DescribeSlowLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "MinQueryTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Role")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogResponseParams struct {
	// Total number of slow queries
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Slow query log details. This parameter has been deprecated and will be replaced by InstanceSlowLogDetail because it is not properly named.
	//
	// Deprecated: InstanceSlowlogDetail is deprecated.
	InstanceSlowlogDetail []*InstanceSlowlogDetail `json:"InstanceSlowlogDetail,omitnil,omitempty" name:"InstanceSlowlogDetail"`

	// Slow query details. Note: If the value of TotalCount is greater than 10,000, indicating that the number of slow logs exceeds 10,000, log details cannot be returned. Instead, the returned data is empty. It is recommended to reduce the interval between BeginTime and EndTime and perform multiple queries.
	InstanceSlowLogDetail []*InstanceSlowlogDetail `json:"InstanceSlowLogDetail,omitnil,omitempty" name:"InstanceSlowLogDetail"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSlowLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowLogResponseParams `json:"Response"`
}

func (r *DescribeSlowLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskInfoRequestParams struct {
	// Task ID, which can be obtained through the sub-parameter **TaskId** of the response parameter **Tasks** of the API [DescribeTaskList](https://intl.cloud.tencent.com/document/product/239/39374?from_cn_redirect=1).
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeTaskInfoRequest struct {
	*tchttp.BaseRequest
	
	// Task ID, which can be obtained through the sub-parameter **TaskId** of the response parameter **Tasks** of the API [DescribeTaskList](https://intl.cloud.tencent.com/document/product/239/39374?from_cn_redirect=1).
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DescribeTaskInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskInfoResponseParams struct {
	// Task status. Valid values: 
	// - `preparing`: To be run
	// - `running`: Running
	// - `succeed`: Succeeded
	// - `failed`: Failed
	// - `Error`: Error occurred while running
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Task start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Task type, including `Create`, `Configure`, `Disable Instance`, `Clear Instance`, `Reset Password`, `Upgrade Version`, `Back up Instance`, `Modify Network`, `Migrate to New AZ` and `Promote to Master`.
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Message returned by task execution, which will be an error message when execution fails or be empty when the status is `running `or `succeed-`.
	TaskMessage *string `json:"TaskMessage,omitnil,omitempty" name:"TaskMessage"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTaskInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskInfoResponseParams `json:"Response"`
}

func (r *DescribeTaskInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskListRequestParams struct {
	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name. Log in to the [Redis console](https://console.cloud.tencent.com/redis/instance/list), and copy it from the instance list.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Number of taskss returned per page.  Default value: `20`. Maximum value: `100`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset, which is an integer multiple of Limit. Offset = Limit x (Page number – 1). The default value is 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Project ID. This field has been deprecated. Please ignore it.
	//
	// Deprecated: ProjectIds is deprecated.
	ProjectIds []*int64 `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// Task type.
	// 
	// 
	// 
	// - FLOW_CREATE: "001", indicating instance creation.
	// - FLOW_RESIZE: "002", indicating configuration modification.
	// - FLOW_CLOSE: "003", indicating instance disabling.
	// - FLOW_CLEAN: "004", indicating instance cleanup.
	// - FLOW_STARTUP: "005", indicating instance enabling.
	// - FLOW_DELETE: "006", indicating instance deletion.
	// - FLOW_SETPWD: "007", indicating password reset.
	// - FLOW_EXPORTBACKUP: "009", indicating backup file export.
	// - FLOW_RESTOREBACKUP: "010", indicating backup restoration.
	// - FLOW_BACKUPINSTANCE: "012", indicating instance backup.
	// - FLOW_MIGRATEINSTANCE: "013", indicating instance migration.
	// - FLOW_DELBACKUP: "014", indicating backup deletion.
	// - FLOW_EXCHANGEINSTANCE: "016", indicating instance switch.
	// - FLOW_AUTOBACKUP: "017", indicating automatic instance backup.
	// - FLOW_MIGRATECHECK: "022", indicating migration parameter verification.
	// - FLOW_MIGRATETASK: "023", indicating that data migration is in progress.
	// - FLOW_CLEANDB: "025", indicating database cleanup.
	// - FLOW_CLONEBACKUP: "026": indicating backup cloning.
	// - FLOW_CHANGEVIP: "027", indicating VIP address modification.
	// - FLOW_EXPORSHR: "028", indicating scaling.
	// - FLOW_ADDNODES: "029", indicating node addition (removal).
	// - FLOW_CHANGENET: "031", indicating network type modification.
	// - FLOW_MODIFYINSTACEREADONLY: "033": indicating read-only policy modification.
	// - FLOW_MODIFYINSTANCEPARAMS: "034", indicating instance parameter modification.
	// - FLOW_MODIFYINSTANCEPASSWORDFREE: "035", indicating password-free access settings.
	// - FLOW_SWITCHINSTANCEVIP: "036", indicating instance VIP address switch.
	// - FLOW_MODIFYINSTANCEACCOUNT: "037", indicating instance account modification.
	// - FLOW_MODIFYINSTANCEBANDWIDTH: "038", indicating instance bandwidth modification.
	// - FLOW_ENABLEINSTANCE_REPLICATE: "039", indicating enabling of read-only replica.
	// - FLOW_DISABLEINSTANCE_REPLICATE: "040", indicating disabling of read-only replica.
	// - FLOW_UpgradeArch: "041", indicating instance architecture upgrade from the standard architecture to the cluster architecture.
	// - FLOW_DowngradeArch: "042", indicating instance architecture downgrade from the cluster architecture to the standard architecture.
	// - FLOW_UpgradeVersion: "043", indicating version upgrade.
	// - FLOW_MODIFYCONNECTIONCONFIG: "044", indicating adjustment of the bandwidth and the number of connections.
	// - FLOW_CLEARNETWORK: "045", indicating network change.
	// - FLOW_REMOVE_BACKUP_FILE: "046", indicating backup deletion.
	// - FLOW_UPGRADE_SUPPORT_MULTI_AZ: "047", indicating instance upgrade to multi-AZ deployment.
	// - FLOW_SHUTDOWN_MASTER: "048", indicating fault simulation.
	// - FLOW_CHANGE_REPLICA_TO_MASTER: "049", indicating manual promotion to the primary node.
	// - FLOW_CODE_ADD_REPLICATION_INSTANCE: "050", indicating replication group addition.
	// - FLOW_OPEN_WAN: "052", indicating enabling of public network access.
	// - FLOW_CLOSE_WAN: "053", indicating disabling of public network access.
	//  - FLOW_UPDATE_WAN: "054", indicating update of the public network access configuration.
	// - FLOW_CODE_DELETE_REPLICATION_INSTANCE: "055", indicating replication group unbinding.
	// - FLOW_CODE_CHANGE_MASTER_INSTANCE: "056", indicating switching a replication group instance to the primary instance.
	// - FLOW_CODE_CHANGE_INSTANCE_ROLE: "057", indicating modification of the replication group instance role.
	// - FLOW_MIGRATE_NODE: "058", indicating node migration.
	// - FLOW_SWITCH_NODE: "059", indicating node switch.
	// - FLOW_UPGRADE_SMALL_VERSION: "060", indicating Redis version upgrade.
	// - FLOW_UPGRADE_PROXY_VERSION: "061", indicating proxy version upgrade.
	// - FLOW_MODIFY_INSTANCE_NETWORK: "062", indicating instance network modification.
	// - FLOW_MIGRATE_PROXY_NODE: "063", indicating proxy node migration.
	// - FLOW_MIGRATION_INSTANCE_ZONE: "066", indicating that instance migration to another AZ is in progress.
	// - FLOW_UPGRADE_INSTANCE_CACHE_AND_PROXY: "067", indicating that instance version upgrade is in progress.
	// - FLOW_MODIFY_PROXY_NUM: "069", indicating proxy node addition (removal).
	// - FLOW_MODIFYBACKUPMOD: "070", indicating instance backup mode modification.
	TaskTypes []*string `json:"TaskTypes,omitnil,omitempty" name:"TaskTypes"`

	// Start time of the task, for example, in the format of 2021-12-30 00:00:00. Data in the last 30 days can be queried.
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// End time of the task, for example, in the format of 2021-12-30 20:59:35. Data in the last 30 days can be queried.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// This parameter is only for internal use and can be ignored.
	TaskStatus []*int64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// Task execution status. Valid values: - `0` (initilized) - `1` (executing) - `2` (completed) - `4` (failed)
	Result []*int64 `json:"Result,omitnil,omitempty" name:"Result"`

	// The field `OperatorUin` has been disused and replaced by `OperateUin`.
	//
	// Deprecated: OperatorUin is deprecated.
	OperatorUin []*int64 `json:"OperatorUin,omitnil,omitempty" name:"OperatorUin"`

	// Operator account ID or UIN
	OperateUin []*string `json:"OperateUin,omitnil,omitempty" name:"OperateUin"`
}

type DescribeTaskListRequest struct {
	*tchttp.BaseRequest
	
	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name. Log in to the [Redis console](https://console.cloud.tencent.com/redis/instance/list), and copy it from the instance list.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Number of taskss returned per page.  Default value: `20`. Maximum value: `100`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset, which is an integer multiple of Limit. Offset = Limit x (Page number – 1). The default value is 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Project ID. This field has been deprecated. Please ignore it.
	ProjectIds []*int64 `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// Task type.
	// 
	// 
	// 
	// - FLOW_CREATE: "001", indicating instance creation.
	// - FLOW_RESIZE: "002", indicating configuration modification.
	// - FLOW_CLOSE: "003", indicating instance disabling.
	// - FLOW_CLEAN: "004", indicating instance cleanup.
	// - FLOW_STARTUP: "005", indicating instance enabling.
	// - FLOW_DELETE: "006", indicating instance deletion.
	// - FLOW_SETPWD: "007", indicating password reset.
	// - FLOW_EXPORTBACKUP: "009", indicating backup file export.
	// - FLOW_RESTOREBACKUP: "010", indicating backup restoration.
	// - FLOW_BACKUPINSTANCE: "012", indicating instance backup.
	// - FLOW_MIGRATEINSTANCE: "013", indicating instance migration.
	// - FLOW_DELBACKUP: "014", indicating backup deletion.
	// - FLOW_EXCHANGEINSTANCE: "016", indicating instance switch.
	// - FLOW_AUTOBACKUP: "017", indicating automatic instance backup.
	// - FLOW_MIGRATECHECK: "022", indicating migration parameter verification.
	// - FLOW_MIGRATETASK: "023", indicating that data migration is in progress.
	// - FLOW_CLEANDB: "025", indicating database cleanup.
	// - FLOW_CLONEBACKUP: "026": indicating backup cloning.
	// - FLOW_CHANGEVIP: "027", indicating VIP address modification.
	// - FLOW_EXPORSHR: "028", indicating scaling.
	// - FLOW_ADDNODES: "029", indicating node addition (removal).
	// - FLOW_CHANGENET: "031", indicating network type modification.
	// - FLOW_MODIFYINSTACEREADONLY: "033": indicating read-only policy modification.
	// - FLOW_MODIFYINSTANCEPARAMS: "034", indicating instance parameter modification.
	// - FLOW_MODIFYINSTANCEPASSWORDFREE: "035", indicating password-free access settings.
	// - FLOW_SWITCHINSTANCEVIP: "036", indicating instance VIP address switch.
	// - FLOW_MODIFYINSTANCEACCOUNT: "037", indicating instance account modification.
	// - FLOW_MODIFYINSTANCEBANDWIDTH: "038", indicating instance bandwidth modification.
	// - FLOW_ENABLEINSTANCE_REPLICATE: "039", indicating enabling of read-only replica.
	// - FLOW_DISABLEINSTANCE_REPLICATE: "040", indicating disabling of read-only replica.
	// - FLOW_UpgradeArch: "041", indicating instance architecture upgrade from the standard architecture to the cluster architecture.
	// - FLOW_DowngradeArch: "042", indicating instance architecture downgrade from the cluster architecture to the standard architecture.
	// - FLOW_UpgradeVersion: "043", indicating version upgrade.
	// - FLOW_MODIFYCONNECTIONCONFIG: "044", indicating adjustment of the bandwidth and the number of connections.
	// - FLOW_CLEARNETWORK: "045", indicating network change.
	// - FLOW_REMOVE_BACKUP_FILE: "046", indicating backup deletion.
	// - FLOW_UPGRADE_SUPPORT_MULTI_AZ: "047", indicating instance upgrade to multi-AZ deployment.
	// - FLOW_SHUTDOWN_MASTER: "048", indicating fault simulation.
	// - FLOW_CHANGE_REPLICA_TO_MASTER: "049", indicating manual promotion to the primary node.
	// - FLOW_CODE_ADD_REPLICATION_INSTANCE: "050", indicating replication group addition.
	// - FLOW_OPEN_WAN: "052", indicating enabling of public network access.
	// - FLOW_CLOSE_WAN: "053", indicating disabling of public network access.
	//  - FLOW_UPDATE_WAN: "054", indicating update of the public network access configuration.
	// - FLOW_CODE_DELETE_REPLICATION_INSTANCE: "055", indicating replication group unbinding.
	// - FLOW_CODE_CHANGE_MASTER_INSTANCE: "056", indicating switching a replication group instance to the primary instance.
	// - FLOW_CODE_CHANGE_INSTANCE_ROLE: "057", indicating modification of the replication group instance role.
	// - FLOW_MIGRATE_NODE: "058", indicating node migration.
	// - FLOW_SWITCH_NODE: "059", indicating node switch.
	// - FLOW_UPGRADE_SMALL_VERSION: "060", indicating Redis version upgrade.
	// - FLOW_UPGRADE_PROXY_VERSION: "061", indicating proxy version upgrade.
	// - FLOW_MODIFY_INSTANCE_NETWORK: "062", indicating instance network modification.
	// - FLOW_MIGRATE_PROXY_NODE: "063", indicating proxy node migration.
	// - FLOW_MIGRATION_INSTANCE_ZONE: "066", indicating that instance migration to another AZ is in progress.
	// - FLOW_UPGRADE_INSTANCE_CACHE_AND_PROXY: "067", indicating that instance version upgrade is in progress.
	// - FLOW_MODIFY_PROXY_NUM: "069", indicating proxy node addition (removal).
	// - FLOW_MODIFYBACKUPMOD: "070", indicating instance backup mode modification.
	TaskTypes []*string `json:"TaskTypes,omitnil,omitempty" name:"TaskTypes"`

	// Start time of the task, for example, in the format of 2021-12-30 00:00:00. Data in the last 30 days can be queried.
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// End time of the task, for example, in the format of 2021-12-30 20:59:35. Data in the last 30 days can be queried.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// This parameter is only for internal use and can be ignored.
	TaskStatus []*int64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// Task execution status. Valid values: - `0` (initilized) - `1` (executing) - `2` (completed) - `4` (failed)
	Result []*int64 `json:"Result,omitnil,omitempty" name:"Result"`

	// The field `OperatorUin` has been disused and replaced by `OperateUin`.
	OperatorUin []*int64 `json:"OperatorUin,omitnil,omitempty" name:"OperatorUin"`

	// Operator account ID or UIN
	OperateUin []*string `json:"OperateUin,omitnil,omitempty" name:"OperateUin"`
}

func (r *DescribeTaskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InstanceName")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "ProjectIds")
	delete(f, "TaskTypes")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "TaskStatus")
	delete(f, "Result")
	delete(f, "OperatorUin")
	delete(f, "OperateUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskListResponseParams struct {
	// Total number of tasks
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Task details
	Tasks []*TaskInfoDetail `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTaskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskListResponseParams `json:"Response"`
}

func (r *DescribeTaskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTendisSlowLogRequestParams struct {
	// Instance ID. Log in to the [Tendis console](https://console.cloud.tencent.com/tendis) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time for a query, for example, 2019-09-08 12:12:41, with a maximum query span of 30 days.
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// End time for a query, for example, 2019-09-09 12:12:41, with a maximum query span of 30 days.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Slow query threshold, in milliseconds. The value is a positive integer greater than 0.
	MinQueryTime *int64 `json:"MinQueryTime,omitnil,omitempty" name:"MinQueryTime"`

	// Page size. The default value is 20, the minimum value is 1, and the maximum value is 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset, which is an integer multiple of Limit. Calculation formula: Offset = Limit x (Page number – 1). The default value is 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeTendisSlowLogRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. Log in to the [Tendis console](https://console.cloud.tencent.com/tendis) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time for a query, for example, 2019-09-08 12:12:41, with a maximum query span of 30 days.
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// End time for a query, for example, 2019-09-09 12:12:41, with a maximum query span of 30 days.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Slow query threshold, in milliseconds. The value is a positive integer greater than 0.
	MinQueryTime *int64 `json:"MinQueryTime,omitnil,omitempty" name:"MinQueryTime"`

	// Page size. The default value is 20, the minimum value is 1, and the maximum value is 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset, which is an integer multiple of Limit. Calculation formula: Offset = Limit x (Page number – 1). The default value is 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeTendisSlowLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTendisSlowLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "MinQueryTime")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTendisSlowLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTendisSlowLogResponseParams struct {
	// Total number of slow queries
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Slow query details. Note: If the value of TotalCount is greater than 10,000, indicating that the number of slow logs exceeds 10,000, log details cannot be returned. Instead, the returned data is empty. It is recommended to reduce the interval between BeginTime and EndTime and perform multiple queries.
	TendisSlowLogDetail []*TendisSlowLogDetail `json:"TendisSlowLogDetail,omitnil,omitempty" name:"TendisSlowLogDetail"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTendisSlowLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTendisSlowLogResponseParams `json:"Response"`
}

func (r *DescribeTendisSlowLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTendisSlowLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyPostpaidInstanceRequestParams struct {
	// Instance ID. Log in to the [Redis console](https://console.cloud.tencent.com/redis/instance/list), and copy the pay-as-you-go instance ID from the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DestroyPostpaidInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. Log in to the [Redis console](https://console.cloud.tencent.com/redis/instance/list), and copy the pay-as-you-go instance ID from the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DestroyPostpaidInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyPostpaidInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyPostpaidInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyPostpaidInstanceResponseParams struct {
	// Task ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DestroyPostpaidInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DestroyPostpaidInstanceResponseParams `json:"Response"`
}

func (r *DestroyPostpaidInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyPostpaidInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyPrepaidInstanceRequestParams struct {
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DestroyPrepaidInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DestroyPrepaidInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyPrepaidInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyPrepaidInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyPrepaidInstanceResponseParams struct {
	// Order ID
	//
	// Deprecated: DealId is deprecated.
	DealId *string `json:"DealId,omitnil,omitempty" name:"DealId"`

	// Order number.
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DestroyPrepaidInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DestroyPrepaidInstanceResponseParams `json:"Response"`
}

func (r *DestroyPrepaidInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyPrepaidInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableReplicaReadonlyRequestParams struct {
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DisableReplicaReadonlyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DisableReplicaReadonlyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableReplicaReadonlyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableReplicaReadonlyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableReplicaReadonlyResponseParams struct {
	// Task ID.
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisableReplicaReadonlyResponse struct {
	*tchttp.BaseResponse
	Response *DisableReplicaReadonlyResponseParams `json:"Response"`
}

func (r *DisableReplicaReadonlyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableReplicaReadonlyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateSecurityGroupsRequestParams struct {
	// Database engine name, which is `redis` for this API.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// Security group ID, which can be obtained through the sub-parameter **SecurityGroupId** of the response parameter InstanceSecurityGroupsDetail of the API [DescribeInstanceSecurityGroup](https://intl.cloud.tencent.com/document/product/239/34447?from_cn_redirect=1).
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// Instance ID list, which is an array of one or more instance IDs. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type DisassociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Database engine name, which is `redis` for this API.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// Security group ID, which can be obtained through the sub-parameter **SecurityGroupId** of the response parameter InstanceSecurityGroupsDetail of the API [DescribeInstanceSecurityGroup](https://intl.cloud.tencent.com/document/product/239/34447?from_cn_redirect=1).
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// Instance ID list, which is an array of one or more instance IDs. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
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

// Predefined struct for user
type DisassociateSecurityGroupsResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisassociateSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateSecurityGroupsResponseParams `json:"Response"`
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

// Predefined struct for user
type EnableReplicaReadonlyRequestParams struct {
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Read-only routing policy.
	// - master: read-only routing to the primary node.
	// - replication: read-only routing to the secondary node. The default value is replication.
	ReadonlyPolicy []*string `json:"ReadonlyPolicy,omitnil,omitempty" name:"ReadonlyPolicy"`
}

type EnableReplicaReadonlyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Read-only routing policy.
	// - master: read-only routing to the primary node.
	// - replication: read-only routing to the secondary node. The default value is replication.
	ReadonlyPolicy []*string `json:"ReadonlyPolicy,omitnil,omitempty" name:"ReadonlyPolicy"`
}

func (r *EnableReplicaReadonlyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableReplicaReadonlyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ReadonlyPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableReplicaReadonlyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableReplicaReadonlyResponseParams struct {
	// ERROR: incorrect; OK: correct (discarded).
	//
	// Deprecated: Status is deprecated.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Task ID.
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EnableReplicaReadonlyResponse struct {
	*tchttp.BaseResponse
	Response *EnableReplicaReadonlyResponseParams `json:"Response"`
}

func (r *EnableReplicaReadonlyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableReplicaReadonlyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportFile struct {
	// <p>File name.</p>
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// <p>Status value.</p><p>Enumeration values: </p><ul><li>creating: File creation in progress, </li><li>success: File generated, </li><li>failed: File generation failed, </li><li>deleted: File deleted.</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>File size. Measurement unit: byte.</p>
	FileSize *int64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// <p>File creation time.</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>Download link.</p>
	DownloadUrl *string `json:"DownloadUrl,omitnil,omitempty" name:"DownloadUrl"`

	// <p>Error information of the exported file.</p>
	ErrMsg *string `json:"ErrMsg,omitnil,omitempty" name:"ErrMsg"`

	// <p>Progress of the exported file.</p>
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// <p>Completion time of the exported file.</p>
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`

	// <p>Asynchronous request ID.</p>
	AsyncRequestId *int64 `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`
}

type FailedInstance struct {
	// <p>Failed instance ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Failure information</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

type Filter struct {
	// <p>Filter field.</p><p>Enumeration value:</p><ul><li>InstanceId: Instance ID.</li><li>InstanceName: Instance name.</li><li>TagKey: Tag key.</li><li>InstanceTags: Instance tag key-value, tag key & tag value.</li></ul>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Value of the filter field.</p>
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// <p>Exact match switch.</p><ul><li>false: Turn off.</li><li>true: Turn on.</li></ul>
	ExactMatch *bool `json:"ExactMatch,omitnil,omitempty" name:"ExactMatch"`
}

type Groups struct {
	// User APPID, which is the unique application ID that matches an account. Some Tencent Cloud products use this APPID.
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Region ID.
	//  - 1: Guangzhou.
	//  - 4: Shanghai.
	//  - 5: Hong Kong (China).
	//  - 7: Shanghai Finance.
	//  - 8: Beijing.
	//  - 9: Singapore.
	//  - 11: Shenzhen Finance.
	//  - 15: Western US (Silicon Valley).
	//  - 16: Chengdu.
	//  - 17: Germany.
	//  - 18: South Korea.
	//  - 19: Chongqing.
	//  - 22: Eastern US (Virginia).
	//  - 23: Thailand.
	// - 25: Japan.
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Replication group ID in the format of "crs-rpl-deind****"
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Replication group name.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Status of replication group
	// - `37`: Associating replication group
	// - `38`: Reconnecting to replication group
	// - `51`: Disassociating replication group
	// - `52`: Switching with master instance in replication group
	// - `53`: Modifying the roles
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Number of replication groups
	InstanceCount *int64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// Instance information on the replication group.
	Instances []*Instances `json:"Instances,omitnil,omitempty" name:"Instances"`

	// Remark information.
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type HotKeyInfo struct {
	// The name of the hot key.
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Key type.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Number of accesses for the hot key in a specified time period.
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type Inbound struct {
	// Policy. Valid values: ACCEPT, DROP.
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// All the addresses that the address group ID represents.
	AddressModule *string `json:"AddressModule,omitnil,omitempty" name:"AddressModule"`

	// Source IP or IP address range, such as 192.168.0.0/16.
	CidrIp *string `json:"CidrIp,omitnil,omitempty" name:"CidrIp"`

	// Description.
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// Network protocol, such as UDP and TCP.
	IpProtocol *string `json:"IpProtocol,omitnil,omitempty" name:"IpProtocol"`

	// Port.
	PortRange *string `json:"PortRange,omitnil,omitempty" name:"PortRange"`

	// All the protocols and ports that the service group ID represents.
	ServiceModule *string `json:"ServiceModule,omitnil,omitempty" name:"ServiceModule"`

	// All the addresses that the security group ID represents.
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

// Predefined struct for user
type InquiryPriceCreateInstanceRequestParams struct {
	// <p>Instance type. - 2: Redis 2.8 memory edition (standard architecture). - 6: Redis 4.0 memory edition (standard architecture). - 7: Redis 4.0 memory edition (cluster architecture). - 8: Redis 5.0 memory edition (standard architecture). - 9: Redis 5.0 memory edition (cluster architecture). - 15: Redis 6.2 memory edition (standard architecture). - 16: Redis 6.2 memory edition (cluster architecture). - 17: Redis 7.0 memory edition (standard architecture). - 18: Redis 7.0 memory edition (cluster architecture). - 200: Memcached 1.6 memory edition (cluster architecture).</p>
	TypeId *uint64 `json:"TypeId,omitnil,omitempty" name:"TypeId"`

	// <p>Memory capacity, measured in MB, must be a multiple of 1024. For specific specifications, refer to the specifications returned by <a href="https://www.tencentcloud.com/document/api/239/30600?from_cn_redirect=1">query product sales specifications</a>. When TypeId is standard architecture, MemSize is the total memory capacity of the instance. When TypeId is cluster architecture, MemSize is the sharded memory capacity.</p>
	MemSize *uint64 `json:"MemSize,omitnil,omitempty" name:"MemSize"`

	// <p>Instance count. The number of instances to purchase at a time is subject to the specifications returned by <a href="https://www.tencentcloud.com/document/api/239/30600?from_cn_redirect=1">query product sales specifications</a>.</p>
	GoodsNum *uint64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// <p>Purchase period needs to be filled in when creating an annual and monthly subscription instance. For pay-as-you-go instances, just fill in 1. Unit: month. Value ranges from 1 to 36 [1,2,3,4,5,6,7,8,9,10,11,12,24,36].</p>
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// <p>Payment method. - 0: Pay-As-You-Go. - 1: Monthly Subscription.</p>
	BillingMode *int64 `json:"BillingMode,omitnil,omitempty" name:"BillingMode"`

	// <p>ID of the AZ to which the instance belongs. See <a href="https://www.tencentcloud.com/document/product/239/4106?from_cn_redirect=1">Regions and AZs</a>.<strong>Note</strong>: Please specify at least one parameter in <strong>ZoneId</strong> and <strong>ZoneName</strong>.</p>
	ZoneId *uint64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>Number of instance shards. - The shard number should be set to 1 for the standard architecture. - The number of shards can be set to 1, 3, 5, 8, 12, 16, 24, 32, 40, 48, 64, 80, 96, or 128 for the cluster architecture.</p>
	RedisShardNum *int64 `json:"RedisShardNum,omitnil,omitempty" name:"RedisShardNum"`

	// <p>Number of instance replicas. Valid values: 1, 2, 3, 4, and 5.</p>
	RedisReplicasNum *int64 `json:"RedisReplicasNum,omitnil,omitempty" name:"RedisReplicasNum"`

	// <p>Whether replica read-only is supported. For Redis 2.8 standard architecture and CKV standard architecture, this parameter is not required. - true: replica read-only not required. - false: read-only replica supported.</p>
	ReplicasReadonly *bool `json:"ReplicasReadonly,omitnil,omitempty" name:"ReplicasReadonly"`

	// <p>Name of the availability zone to which the instance belongs. See <a href="https://www.tencentcloud.com/document/product/239/4106?from_cn_redirect=1">Regions and Availability Zones</a>. <strong>Description</strong>: Please specify at least one parameter in <strong>ZoneId</strong> and <strong>ZoneName</strong>.</p>
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// <p>Deployment method. - local: local disk, defaults to local. - cloud: cloud disk. - cdc: dedicated cluster edition.</p>
	ProductVersion *string `json:"ProductVersion,omitnil,omitempty" name:"ProductVersion"`
}

type InquiryPriceCreateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance type. - 2: Redis 2.8 memory edition (standard architecture). - 6: Redis 4.0 memory edition (standard architecture). - 7: Redis 4.0 memory edition (cluster architecture). - 8: Redis 5.0 memory edition (standard architecture). - 9: Redis 5.0 memory edition (cluster architecture). - 15: Redis 6.2 memory edition (standard architecture). - 16: Redis 6.2 memory edition (cluster architecture). - 17: Redis 7.0 memory edition (standard architecture). - 18: Redis 7.0 memory edition (cluster architecture). - 200: Memcached 1.6 memory edition (cluster architecture).</p>
	TypeId *uint64 `json:"TypeId,omitnil,omitempty" name:"TypeId"`

	// <p>Memory capacity, measured in MB, must be a multiple of 1024. For specific specifications, refer to the specifications returned by <a href="https://www.tencentcloud.com/document/api/239/30600?from_cn_redirect=1">query product sales specifications</a>. When TypeId is standard architecture, MemSize is the total memory capacity of the instance. When TypeId is cluster architecture, MemSize is the sharded memory capacity.</p>
	MemSize *uint64 `json:"MemSize,omitnil,omitempty" name:"MemSize"`

	// <p>Instance count. The number of instances to purchase at a time is subject to the specifications returned by <a href="https://www.tencentcloud.com/document/api/239/30600?from_cn_redirect=1">query product sales specifications</a>.</p>
	GoodsNum *uint64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// <p>Purchase period needs to be filled in when creating an annual and monthly subscription instance. For pay-as-you-go instances, just fill in 1. Unit: month. Value ranges from 1 to 36 [1,2,3,4,5,6,7,8,9,10,11,12,24,36].</p>
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// <p>Payment method. - 0: Pay-As-You-Go. - 1: Monthly Subscription.</p>
	BillingMode *int64 `json:"BillingMode,omitnil,omitempty" name:"BillingMode"`

	// <p>ID of the AZ to which the instance belongs. See <a href="https://www.tencentcloud.com/document/product/239/4106?from_cn_redirect=1">Regions and AZs</a>.<strong>Note</strong>: Please specify at least one parameter in <strong>ZoneId</strong> and <strong>ZoneName</strong>.</p>
	ZoneId *uint64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>Number of instance shards. - The shard number should be set to 1 for the standard architecture. - The number of shards can be set to 1, 3, 5, 8, 12, 16, 24, 32, 40, 48, 64, 80, 96, or 128 for the cluster architecture.</p>
	RedisShardNum *int64 `json:"RedisShardNum,omitnil,omitempty" name:"RedisShardNum"`

	// <p>Number of instance replicas. Valid values: 1, 2, 3, 4, and 5.</p>
	RedisReplicasNum *int64 `json:"RedisReplicasNum,omitnil,omitempty" name:"RedisReplicasNum"`

	// <p>Whether replica read-only is supported. For Redis 2.8 standard architecture and CKV standard architecture, this parameter is not required. - true: replica read-only not required. - false: read-only replica supported.</p>
	ReplicasReadonly *bool `json:"ReplicasReadonly,omitnil,omitempty" name:"ReplicasReadonly"`

	// <p>Name of the availability zone to which the instance belongs. See <a href="https://www.tencentcloud.com/document/product/239/4106?from_cn_redirect=1">Regions and Availability Zones</a>. <strong>Description</strong>: Please specify at least one parameter in <strong>ZoneId</strong> and <strong>ZoneName</strong>.</p>
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// <p>Deployment method. - local: local disk, defaults to local. - cloud: cloud disk. - cdc: dedicated cluster edition.</p>
	ProductVersion *string `json:"ProductVersion,omitnil,omitempty" name:"ProductVersion"`
}

func (r *InquiryPriceCreateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceCreateInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TypeId")
	delete(f, "MemSize")
	delete(f, "GoodsNum")
	delete(f, "Period")
	delete(f, "BillingMode")
	delete(f, "ZoneId")
	delete(f, "RedisShardNum")
	delete(f, "RedisReplicasNum")
	delete(f, "ReplicasReadonly")
	delete(f, "ZoneName")
	delete(f, "ProductVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceCreateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceCreateInstanceResponseParams struct {
	// <p>Discounted price.</p>
	Price *float64 `json:"Price,omitnil,omitempty" name:"Price"`

	// <p>High-precision discounted price</p>
	HighPrecisionPrice *float64 `json:"HighPrecisionPrice,omitnil,omitempty" name:"HighPrecisionPrice"`

	// <p>Original price</p>
	OriginalPrice *float64 `json:"OriginalPrice,omitnil,omitempty" name:"OriginalPrice"`

	// <p>High-precision original price</p>
	HighPrecisionOriginalPrice *float64 `json:"HighPrecisionOriginalPrice,omitnil,omitempty" name:"HighPrecisionOriginalPrice"`

	// <p>Currency</p>
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`

	// <p>Price amount unit</p><ul><li>pent: cent</li><li>microPent: microcent</li></ul>
	AmountUnit *string `json:"AmountUnit,omitnil,omitempty" name:"AmountUnit"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquiryPriceCreateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceCreateInstanceResponseParams `json:"Response"`
}

func (r *InquiryPriceCreateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceCreateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceUpgradeInstanceRequestParams struct {
	// <p>Instance ID. Log in to the <a href="https://console.cloud.tencent.com/redis/instance/list">Redis console</a> and copy the instance ID from the instance list.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Shard size. Unit: MB.</p>
	MemSize *uint64 `json:"MemSize,omitnil,omitempty" name:"MemSize"`

	// <p>Number of shards. - For instances with standard architecture, RedisShardNum defaults to 1. - This parameter is not required for Redis 2.8 Primary-Secondary Edition, CKV Primary-Secondary Edition, and Redis 2.8 Standalone Edition.</p>
	RedisShardNum *uint64 `json:"RedisShardNum,omitnil,omitempty" name:"RedisShardNum"`

	// <p>Number of replicas. This parameter is not required for Redis 2.8 Primary-Secondary Edition, CKV Primary-Secondary Edition, and Redis 2.8 Single-node Edition.</p>
	RedisReplicasNum *uint64 `json:"RedisReplicasNum,omitnil,omitempty" name:"RedisReplicasNum"`
}

type InquiryPriceUpgradeInstanceRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance ID. Log in to the <a href="https://console.cloud.tencent.com/redis/instance/list">Redis console</a> and copy the instance ID from the instance list.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Shard size. Unit: MB.</p>
	MemSize *uint64 `json:"MemSize,omitnil,omitempty" name:"MemSize"`

	// <p>Number of shards. - For instances with standard architecture, RedisShardNum defaults to 1. - This parameter is not required for Redis 2.8 Primary-Secondary Edition, CKV Primary-Secondary Edition, and Redis 2.8 Standalone Edition.</p>
	RedisShardNum *uint64 `json:"RedisShardNum,omitnil,omitempty" name:"RedisShardNum"`

	// <p>Number of replicas. This parameter is not required for Redis 2.8 Primary-Secondary Edition, CKV Primary-Secondary Edition, and Redis 2.8 Single-node Edition.</p>
	RedisReplicasNum *uint64 `json:"RedisReplicasNum,omitnil,omitempty" name:"RedisReplicasNum"`
}

func (r *InquiryPriceUpgradeInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceUpgradeInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "MemSize")
	delete(f, "RedisShardNum")
	delete(f, "RedisReplicasNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceUpgradeInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceUpgradeInstanceResponseParams struct {
	// <p>Discounted price.</p>
	Price *float64 `json:"Price,omitnil,omitempty" name:"Price"`

	// <p>High-precision discounted price</p>
	HighPrecisionPrice *float64 `json:"HighPrecisionPrice,omitnil,omitempty" name:"HighPrecisionPrice"`

	// <p>Original price</p>
	OriginalPrice *float64 `json:"OriginalPrice,omitnil,omitempty" name:"OriginalPrice"`

	// <p>High-precision original price</p>
	HighPrecisionOriginalPrice *float64 `json:"HighPrecisionOriginalPrice,omitnil,omitempty" name:"HighPrecisionOriginalPrice"`

	// <p>Currency</p>
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`

	// <p>Price amount unit</p><ul><li>pent: cent</li><li>microPent: microcent</li></ul>
	AmountUnit *string `json:"AmountUnit,omitnil,omitempty" name:"AmountUnit"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquiryPriceUpgradeInstanceResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceUpgradeInstanceResponseParams `json:"Response"`
}

func (r *InquiryPriceUpgradeInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceUpgradeInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceClusterNode struct {
	// Node group name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// ID of the runtime node of an instance
	RunId *string `json:"RunId,omitnil,omitempty" name:"RunId"`

	// Cluster role. Valid values:  - `0` (master) - `1` (replica)
	Role *int64 `json:"Role,omitnil,omitempty" name:"Role"`

	// Node status. Valid values:  - `0` (read/write) - `1` (read) - `2` (backup)
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Service status. Valid values: `0` (down), `1` (on).
	Connected *int64 `json:"Connected,omitnil,omitempty" name:"Connected"`

	// Node creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Node elimination time
	DownTime *string `json:"DownTime,omitnil,omitempty" name:"DownTime"`

	// Node slot distribution range
	Slots *string `json:"Slots,omitnil,omitempty" name:"Slots"`

	// Distribution of node keys
	Keys *int64 `json:"Keys,omitnil,omitempty" name:"Keys"`

	// Node QPS Number of executions per second on sharded nodes Unit: Counts/sec
	Qps *int64 `json:"Qps,omitnil,omitempty" name:"Qps"`

	// QPS slope of a node
	QpsSlope *float64 `json:"QpsSlope,omitnil,omitempty" name:"QpsSlope"`

	// Node storage
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// Node storage slope
	StorageSlope *float64 `json:"StorageSlope,omitnil,omitempty" name:"StorageSlope"`
}

type InstanceClusterShard struct {
	// The name of a shard node
	ShardName *string `json:"ShardName,omitnil,omitempty" name:"ShardName"`

	// The serial number of a shard node
	ShardId *string `json:"ShardId,omitnil,omitempty" name:"ShardId"`

	// The role of a shard node
	// - `0`: Master node.
	// - `1`: Replica node.
	Role *int64 `json:"Role,omitnil,omitempty" name:"Role"`

	// Number of keys
	Keys *int64 `json:"Keys,omitnil,omitempty" name:"Keys"`

	// Slot information
	Slots *string `json:"Slots,omitnil,omitempty" name:"Slots"`

	// Used Capacity
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// Capacity slope
	StorageSlope *float64 `json:"StorageSlope,omitnil,omitempty" name:"StorageSlope"`

	// This field is recommended to use the RunId instead due to spelling inconsistency.
	//  Meaning: The node ID during instance runtime.
	Runid *string `json:"Runid,omitnil,omitempty" name:"Runid"`

	// The node ID during instance runtime.
	RunId *string `json:"RunId,omitnil,omitempty" name:"RunId"`

	// Service status
	// - `0`: Down.
	// - `1`: On.
	Connected *int64 `json:"Connected,omitnil,omitempty" name:"Connected"`

	// AZ information.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Node group ID.
	ReplicasNodeId *int64 `json:"ReplicasNodeId,omitnil,omitempty" name:"ReplicasNodeId"`
}

type InstanceEnumParam struct {
	// Parameter name
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// Parameter type, such as `Enum`.
	ValueType *string `json:"ValueType,omitnil,omitempty" name:"ValueType"`

	// Whether to restart the database after modifying the parameters. Valid values: - `true` (required) - `false` (not required)
	NeedRestart *string `json:"NeedRestart,omitnil,omitempty" name:"NeedRestart"`

	// Default value of the parameter
	DefaultValue *string `json:"DefaultValue,omitnil,omitempty" name:"DefaultValue"`

	// Current value
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`

	// Description
	Tips *string `json:"Tips,omitnil,omitempty" name:"Tips"`

	// Acceptable values for the parameter
	EnumValue []*string `json:"EnumValue,omitnil,omitempty" name:"EnumValue"`

	// Parameter modification status. Valid values: - `1` (modifying) - `2` (modified)
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type InstanceInfo struct {
	// <p>Instance name.</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>Project ID</p>
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Instance status
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Substatus of the instance in process returned.</p><p>Enumeration value:</p><ul><li>0: Read-only disk.</li></ul>
	SubStatus *int64 `json:"SubStatus,omitnil,omitempty" name:"SubStatus"`

	// <p>Region.</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>Zone</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>Degradation policy, measurement unit: ms. After the instance P99 reaches the degradation policy, audit data is automatically discarded to prioritize business availability. Default value: 500 ms. Range value: 300-1000 ms.</p>
	DegradeStrategy *int64 `json:"DegradeStrategy,omitnil,omitempty" name:"DegradeStrategy"`

	// <p>Tag information</p>
	InstanceTags []*InstanceTagInfo `json:"InstanceTags,omitnil,omitempty" name:"InstanceTags"`

	// <p>Architecture edition</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`
}

type InstanceIntegerParam struct {
	// Parameter name
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// Parameter type: Integer
	ValueType *string `json:"ValueType,omitnil,omitempty" name:"ValueType"`

	// Whether restart is required after a modification is made. Valid values: true, false
	NeedRestart *string `json:"NeedRestart,omitnil,omitempty" name:"NeedRestart"`

	// Default value of the parameter
	DefaultValue *string `json:"DefaultValue,omitnil,omitempty" name:"DefaultValue"`

	// Current value
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`

	// Parameter description
	Tips *string `json:"Tips,omitnil,omitempty" name:"Tips"`

	// Minimum value of the parameter
	Min *string `json:"Min,omitnil,omitempty" name:"Min"`

	// Maximum value of the parameter
	Max *string `json:"Max,omitnil,omitempty" name:"Max"`

	// Parameter status. 1: modifying; 2: modified
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Parameter unit.
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`
}

type InstanceMultiParam struct {
	// Parameter name
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// Parameter Type such as  `MULTI`
	ValueType *string `json:"ValueType,omitnil,omitempty" name:"ValueType"`

	// Whether to restart the database after modifying the parameter. Valid values:  - `true` (required) - `false` (not required)
	NeedRestart *string `json:"NeedRestart,omitnil,omitempty" name:"NeedRestart"`

	// Default value of the parameter
	DefaultValue *string `json:"DefaultValue,omitnil,omitempty" name:"DefaultValue"`

	// Current value
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`

	// Description
	Tips *string `json:"Tips,omitnil,omitempty" name:"Tips"`

	// Parameter enumeration value.
	EnumValue []*string `json:"EnumValue,omitnil,omitempty" name:"EnumValue"`

	// Parameter modification status. Valid values: - `1` (modifying) - `2` (modified)
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type InstanceNode struct {
	// Instance ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Node details
	InstanceClusterNode []*InstanceClusterNode `json:"InstanceClusterNode,omitnil,omitempty" name:"InstanceClusterNode"`
}

type InstanceParam struct {
	// Parameter name, such as “timeout”. For supported custom parameters, see <a href="https://www.tencentcloud.com/document/product/239/39796">Setting Instance Parameters</a>
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Current parameter value. For example, if you set the current value of “timeout” to 120 (in seconds), the client connections that remain idle longer than 120 seconds will be closed. For more information on parameter values, see <a href="https://www.tencentcloud.com/document/product/239/39796">Setting Instance Parameters</a>
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type InstanceParamHistory struct {
	// Parameter name
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// The value of the parameter before modification
	PreValue *string `json:"PreValue,omitnil,omitempty" name:"PreValue"`

	// The value of the parameter after modification
	NewValue *string `json:"NewValue,omitnil,omitempty" name:"NewValue"`

	// Parameter configuration status
	// - `1`: The parameter configuration is being modified.
	// - `2`: The parameter configuration has been modified successfully.
	// - `3`: Failed to modify the parameter configuration.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Modification time
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`
}

type InstanceProxySlowlogDetail struct {
	// <p>Slow query duration. Measurement unit: ms.</p>
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// <p>Client address.</p>
	Client *string `json:"Client,omitnil,omitempty" name:"Client"`

	// <p>Slow query command.</p>
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// <p>Slow query command line information.</p>
	CommandLine *string `json:"CommandLine,omitnil,omitempty" name:"CommandLine"`

	// <p>Execution time.</p>
	ExecuteTime *string `json:"ExecuteTime,omitnil,omitempty" name:"ExecuteTime"`

	// <p>Duration of receiving client requests (ms)</p>
	RecvClientEnd *int64 `json:"RecvClientEnd,omitnil,omitempty" name:"RecvClientEnd"`

	// <p>Duration of sending client requests (ms)</p>
	SendClientEnd *int64 `json:"SendClientEnd,omitnil,omitempty" name:"SendClientEnd"`

	// <p>Proxy node ID.</p>
	Node *string `json:"Node,omitnil,omitempty" name:"Node"`
}

type InstanceSecurityGroupDetail struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Security group information, which includes  security group ID, name, outbound and inbound rules.
	SecurityGroupDetails []*SecurityGroupDetail `json:"SecurityGroupDetails,omitnil,omitempty" name:"SecurityGroupDetails"`
}

type InstanceSet struct {
	// <p>Instance name.</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>User AppId. AppId is an application ID with a one-to-one correspondence to the account ID. Some Tencent Cloud products use this AppId.</p>
	Appid *int64 `json:"Appid,omitnil,omitempty" name:"Appid"`

	// <p>Project ID.</p>
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Region ID.<ul><li>1: Guangzhou.</li><li>4: Shanghai.</li><li>5: Hong Kong (China).</li><li>7: Shanghai Finance.</li><li>8: Beijing.</li><li>9: Singapore.</li><li>11: Shenzhen Finance.</li><li>15: Western US (Silicon Valley).</li><li>16: Chengdu.</li><li>17: Frankfurt.</li><li>18: Seoul.</li><li>19: Chongqing.</li><li>22: Eastern US (Virginia).</li><li>23: Bangkok.</li><li>25: Tokyo.</li></ul></p>
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// <p>Region ID.</p>
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>vpc network ID, such as 75101.</p>
	VpcId *int64 `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>Subnet ID under vpc, for example: 46315.</p>
	SubnetId *int64 `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>Current instance status. <ul><li>0: to be initialized;</li> <li>1: in process;</li> <li>2: running;</li> <li>-2: isolated;</li> <li>-3: to be deleted.</li></ul></p>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Instance VIP.</p>
	WanIp *string `json:"WanIp,omitnil,omitempty" name:"WanIp"`

	// <p>Instance port number.</p>
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>Instance creation time, for example, in the format of 2020-01-15 10:20:00.</p>
	Createtime *string `json:"Createtime,omitnil,omitempty" name:"Createtime"`

	// <p>Instance memory capacity. Unit: MB (1 MB = 1024 KB).</p>
	Size *float64 `json:"Size,omitnil,omitempty" name:"Size"`

	// <p>This field is deprecated. Please use the Tencent Cloud observability platform API interface <a href="https://www.tencentcloud.com/document/product/248/31014?from_cn_redirect=1">GetMonitorData</a> to obtain the memory capacity used by the instance.</p>
	//
	// Deprecated: SizeUsed is deprecated.
	SizeUsed *float64 `json:"SizeUsed,omitnil,omitempty" name:"SizeUsed"`

	// <p>Instance type.</p><p>Enumeration value:</p><ul><li>2: Redis 2.8 memory edition (standard architecture).</li><li>3: CKV 3.2 memory edition (standard architecture).</li><li>4: CKV 3.2 memory edition (cluster architecture).</li><li>5: Redis 2.8 memory edition (standalone).</li><li>6: Redis 4.0 memory edition (standard architecture).</li><li>7: Redis 4.0 memory edition (cluster architecture).</li><li>8: Redis 5.0 memory edition (standard architecture).</li><li>9: Redis 5.0 memory edition (cluster architecture).</li><li>15: Redis 6.2 memory edition (standard architecture).</li><li>16: Redis 6.2 memory edition (cluster architecture).</li><li>17: Redis 7.0 memory edition (standard architecture).</li><li>18: Redis 7.0 memory edition (cluster architecture).</li><li>19: Valkey 8.0 memory edition (standard architecture).</li><li>20: Valkey 8.0 memory edition (cluster architecture).</li><li>21: Valkey 8.0 memory edition (standard architecture).</li><li>22: Valkey 8.0 memory edition (cluster architecture).</li><li>200: Memcached 1.6 memory edition (cluster architecture).</li></ul>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Whether the automatic renewal flag is set for an instance.</p><ul><li>1: set auto-renewal.</li><li>0: automatic renewal flag not set.</li></ul>
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// <p>Expiration time of a monthly subscription instance.</p>
	DeadlineTime *string `json:"DeadlineTime,omitnil,omitempty" name:"DeadlineTime"`

	// <p>Engine. Valid values: Redis Community Edition and Tencent Cloud CKV.</p>
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// <p>Product type.<ul><li>standalone: standard version.</li><li>cluster: cluster version.</li></ul></p>
	ProductType *string `json:"ProductType,omitnil,omitempty" name:"ProductType"`

	// <p>vpc Network id, such as vpc-fk33jsf43kgv.</p>
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// <p>subnet id under vpc, for example: subnet-fd3j6l35mm0.</p>
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// <p>Billing mode.<ul><li>0: Pay-As-You-Go.</li><li>1: Monthly Subscription.</li></ul></p>
	BillingMode *int64 `json:"BillingMode,omitnil,omitempty" name:"BillingMode"`

	// <p>Description of instance running status: for example "instance running".</p>
	InstanceTitle *string `json:"InstanceTitle,omitnil,omitempty" name:"InstanceTitle"`

	// <p>Default termination time of isolated instances. Pay-as-you-go instance offline after isolation. Monthly Subscription instance offline after 7 days. In the format of: 2020-02-15 10:20:00.</p>
	OfflineTime *string `json:"OfflineTime,omitnil,omitempty" name:"OfflineTime"`

	// <p>Substatus of the instance in process returned.</p><ul><li>0: Read and write status of the disk.</li><li>1: Read-only status of the disk due to exceeding limit.</li></ul>
	SubStatus *int64 `json:"SubStatus,omitnil,omitempty" name:"SubStatus"`

	// <p>Anti-affinity tag.</p>
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>Instance node information.</p>
	InstanceNode []*InstanceNode `json:"InstanceNode,omitnil,omitempty" name:"InstanceNode"`

	// <p>Shard size.</p>
	RedisShardSize *int64 `json:"RedisShardSize,omitnil,omitempty" name:"RedisShardSize"`

	// <p>Number of shards.</p>
	RedisShardNum *int64 `json:"RedisShardNum,omitnil,omitempty" name:"RedisShardNum"`

	// <p>Number of replicas.</p>
	RedisReplicasNum *int64 `json:"RedisReplicasNum,omitnil,omitempty" name:"RedisReplicasNum"`

	// <p>Billing ID.</p>
	PriceId *int64 `json:"PriceId,omitnil,omitempty" name:"PriceId"`

	// <p>Time when an instance starts to be isolated.</p>
	CloseTime *string `json:"CloseTime,omitnil,omitempty" name:"CloseTime"`

	// <p>Read weight of the secondary node.</p><ul><li>0: means disable read-only replica.</li><li>100: means enable read-only replica.</li></ul>
	SlaveReadWeight *int64 `json:"SlaveReadWeight,omitnil,omitempty" name:"SlaveReadWeight"`

	// <p>Tag information associated with an instance.</p>
	InstanceTags []*InstanceTagInfo `json:"InstanceTags,omitnil,omitempty" name:"InstanceTags"`

	// <p>Project name.</p>
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// <p>Whether an instance is a password-free instance. <ul><li>true: yes;</li> <li>false: no.</li></ul></p>
	NoAuth *bool `json:"NoAuth,omitnil,omitempty" name:"NoAuth"`

	// <p>Number of client connections.</p>
	ClientLimit *int64 `json:"ClientLimit,omitnil,omitempty" name:"ClientLimit"`

	// <p>DTS status (internal parameter, can be ignored by users).</p>
	DtsStatus *int64 `json:"DtsStatus,omitnil,omitempty" name:"DtsStatus"`

	// <p>Upper limit of the shard bandwidth. Unit: MB.</p>
	NetLimit *int64 `json:"NetLimit,omitnil,omitempty" name:"NetLimit"`

	// <p>Password-free instance flag (internal parameter, which can be ignored).</p>
	PasswordFree *int64 `json:"PasswordFree,omitnil,omitempty" name:"PasswordFree"`

	// <p>Internal parameter, which can be ignored. This parameter is not properly named. It is recommended to use the IPv6 parameter to replace it.</p>
	Vip6 *string `json:"Vip6,omitnil,omitempty" name:"Vip6"`

	// <p>Internal parameter, which can be ignored.</p>
	IPv6 *string `json:"IPv6,omitnil,omitempty" name:"IPv6"`

	// <p>Instance read-only flag (internal parameter, which can be ignored).</p>
	ReadOnly *int64 `json:"ReadOnly,omitnil,omitempty" name:"ReadOnly"`

	// <p>Internal parameter, which can be ignored.</p>
	RemainBandwidthDuration *string `json:"RemainBandwidthDuration,omitnil,omitempty" name:"RemainBandwidthDuration"`

	// <p>For Redis instances, ignore this parameter.</p>
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// <p>Monitoring version.<ul><li>1m: 1-minute granularity monitoring. This monitoring granularity is deprecated. For details, please see <a href="https://www.tencentcloud.com/document/product/239/80653?from_cn_redirect=1">TencentDB for Redis 1-minute granularity deprecation notice</a>.</li><li>5s: 5-second granularity monitoring.</li></ul></p>
	MonitorVersion *string `json:"MonitorVersion,omitnil,omitempty" name:"MonitorVersion"`

	// <p>Minimum value that can be set for the maximum number of client connections.</p>
	ClientLimitMin *int64 `json:"ClientLimitMin,omitnil,omitempty" name:"ClientLimitMin"`

	// <p>Maximum value that can be set for the maximum number of client connections.</p>
	ClientLimitMax *int64 `json:"ClientLimitMax,omitnil,omitempty" name:"ClientLimitMax"`

	// <p>Detailed node information of the instance.<br>Only multi-AZ instances will be returned.</p>
	NodeSet []*RedisNodeInfo `json:"NodeSet,omitnil,omitempty" name:"NodeSet"`

	// <p>Region information of an instance, for example, ap-guangzhou.</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>Public network address.</p>
	WanAddress *string `json:"WanAddress,omitnil,omitempty" name:"WanAddress"`

	// <p>Polaris service address for internal use.</p>
	PolarisServer *string `json:"PolarisServer,omitnil,omitempty" name:"PolarisServer"`

	// <p>CDC Redis cluster ID.</p>
	RedisClusterId *string `json:"RedisClusterId,omitnil,omitempty" name:"RedisClusterId"`

	// <p>CDC cluster ID.</p>
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`

	// <p>Product edition. <ul><li>local: local disk;</li> <li>cloud: cloud disk;</li> <li>cdc: CDC cluster edition.</li></ul></p>
	ProductVersion *string `json:"ProductVersion,omitnil,omitempty" name:"ProductVersion"`

	// <p>Current Proxy version of the instance.</p>
	CurrentProxyVersion *string `json:"CurrentProxyVersion,omitnil,omitempty" name:"CurrentProxyVersion"`

	// <p>Current Cache minor version of an instance. If the instance joins a global replication group, the kernel version of the global replication group will be displayed.</p>
	CurrentRedisVersion *string `json:"CurrentRedisVersion,omitnil,omitempty" name:"CurrentRedisVersion"`

	// <p>Upgradable Proxy version of an instance.</p>
	UpgradeProxyVersion *string `json:"UpgradeProxyVersion,omitnil,omitempty" name:"UpgradeProxyVersion"`

	// <p>Upgradable Cache minor version of an instance.</p>
	UpgradeRedisVersion *string `json:"UpgradeRedisVersion,omitnil,omitempty" name:"UpgradeRedisVersion"`

	// <p>Backup mode.</p><ul><li>SecondLevelBackup: second-level backup.</li><li>NormalLevelBackup: normal backup.</li></ul>
	BackupMode *string `json:"BackupMode,omitnil,omitempty" name:"BackupMode"`

	// <p>Instance destruction protection switch.</p><ul><li>0: disabled.</li><li>1: enabled.</li></ul>
	DeleteProtectionSwitch *int64 `json:"DeleteProtectionSwitch,omitnil,omitempty" name:"DeleteProtectionSwitch"`
}

type InstanceSlowlogDetail struct {
	// Slow log duration
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// Client address
	Client *string `json:"Client,omitnil,omitempty" name:"Client"`

	// Command
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// Command line details
	CommandLine *string `json:"CommandLine,omitnil,omitempty" name:"CommandLine"`

	// Execution duration
	ExecuteTime *string `json:"ExecuteTime,omitnil,omitempty" name:"ExecuteTime"`

	// Node ID
	Node *string `json:"Node,omitnil,omitempty" name:"Node"`
}

type InstanceTagInfo struct {
	// Tag key
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type InstanceTextParam struct {
	// Parameter name
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// Parameter type such as  `Text`.
	ValueType *string `json:"ValueType,omitnil,omitempty" name:"ValueType"`

	// Whether to restart the database after modifying the parameter. Valid values:  - `true` (required) - `false` (not required)
	NeedRestart *string `json:"NeedRestart,omitnil,omitempty" name:"NeedRestart"`

	// Default value of the parameter
	DefaultValue *string `json:"DefaultValue,omitnil,omitempty" name:"DefaultValue"`

	// Current value
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`

	// Description
	Tips *string `json:"Tips,omitnil,omitempty" name:"Tips"`

	// Acceptable values of the parameter
	TextValue []*string `json:"TextValue,omitnil,omitempty" name:"TextValue"`

	// Parameter modification status. Valid values: - `1` (modifying) - `2` (modified)
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type Instances struct {
	// User APPID, which is the unique application ID that matches an account. Some Tencent Cloud products use this APPID.
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Region ID. <ul><li>1: Guangzhou;</li> <li>4: Shanghai;</li> <li>5: Hong Kong (China);</li> <li>7: Shanghai Finance;</li> <li>8: Beijing;</li> <li>9: Singapore;</li> <li>11: Shenzhen Finance;</li> <li>15: Western United States (Silicon Valley).</li></ul>
	RegionId *uint64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Region ID
	ZoneId *uint64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Number of replicas
	RedisReplicasNum *uint64 `json:"RedisReplicasNum,omitnil,omitempty" name:"RedisReplicasNum"`

	// Number of shards
	RedisShardNum *int64 `json:"RedisShardNum,omitnil,omitempty" name:"RedisShardNum"`

	// Shard memory size.
	RedisShardSize *int64 `json:"RedisShardSize,omitnil,omitempty" name:"RedisShardSize"`

	// Instance disk size.
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// Engine: Redis Community Edition, Tencent Cloud CKV.
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// Read-write permission of the instance. <ul><li>`rw`: Read/Write. </li><li>`r`: Read-only. </li></ul>
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// Instance VIP
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// This parameter encounters a naming issue. It is recommended to use the parameter IPv6 instead. It is an internal parameter and can be ignored.
	Vip6 *string `json:"Vip6,omitnil,omitempty" name:"Vip6"`

	// Internal parameter, which can be ignored.
	IPv6 *string `json:"IPv6,omitnil,omitempty" name:"IPv6"`

	// VPC ID, such as `75101`.
	VpcID *int64 `json:"VpcID,omitnil,omitempty" name:"VpcID"`

	// Instance port
	VPort *int64 `json:"VPort,omitnil,omitempty" name:"VPort"`

	// Instance status. <ul><li>`0`: Uninitialized. </li><li>`1`: In the process. </li><li>`2`: Running. </li><li>`-2`: Isolated. </li><li>`-3`: To be deleted. </li></ul>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Repository ID
	GrocerySysId *int64 `json:"GrocerySysId,omitnil,omitempty" name:"GrocerySysId"`

	// Instance type
	// - `2`: Redis 2.8 Memory Edition (Standard Architecture).
	// - `3`: CKV 3.2 Memory Edition (Standard Architecture).
	// - `4`: CKV 3.2 Memory Edition (Cluster Architecture)
	// - `5`: Redis 2.8 Memory Edition (Standalone)
	// - `6`: Redis 4.0 Memory Edition (Standard Architecture).
	// - `7`: Redis 4.0 Memory Edition (Cluster Architecture)
	// - `8`: Redis 5.0 Memory Edition (Standard Architecture).
	// - `9`: Redis 5.0 Memory Edition (Cluster Architecture)
	// - `15`: Redis 6.2 Memory Edition (Standard Architecture).
	// - `16`: Redis 6.2 Memory Edition (Cluster Architecture)
	ProductType *int64 `json:"ProductType,omitnil,omitempty" name:"ProductType"`

	// The time when the instance was added to the replication group.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// The time when instances in the replication group were updated.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

// Predefined struct for user
type KillMasterGroupRequestParams struct {
	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// A parameter used to configure the access password for a specified instance. If password-free authentication is enabled, this parameter will not be required. Required password strength. - It must contains 8-30 characters. We recommend that you use a password of more than 12 characters. - It must contain at least two of the following types: lowercase letters, uppercase letters, digits, and symbols (()`~!@#$%^&*-+=_|{}[]:;<>,.?/), and it cannot start with a slash (/).
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Sharded cluster ID, which can be obtained through **ClusterId** of the response parameter 
	//  **Redis** of the API [DescribeInstanceNodeInfo](https://intl.cloud.tencent.com/document/product/239/48603?from_cn_redirect=1).
	ShardIds []*int64 `json:"ShardIds,omitnil,omitempty" name:"ShardIds"`
}

type KillMasterGroupRequest struct {
	*tchttp.BaseRequest
	
	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// A parameter used to configure the access password for a specified instance. If password-free authentication is enabled, this parameter will not be required. Required password strength. - It must contains 8-30 characters. We recommend that you use a password of more than 12 characters. - It must contain at least two of the following types: lowercase letters, uppercase letters, digits, and symbols (()`~!@#$%^&*-+=_|{}[]:;<>,.?/), and it cannot start with a slash (/).
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Sharded cluster ID, which can be obtained through **ClusterId** of the response parameter 
	//  **Redis** of the API [DescribeInstanceNodeInfo](https://intl.cloud.tencent.com/document/product/239/48603?from_cn_redirect=1).
	ShardIds []*int64 `json:"ShardIds,omitnil,omitempty" name:"ShardIds"`
}

func (r *KillMasterGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KillMasterGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Password")
	delete(f, "ShardIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "KillMasterGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type KillMasterGroupResponseParams struct {
	// Async task ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type KillMasterGroupResponse struct {
	*tchttp.BaseResponse
	Response *KillMasterGroupResponseParams `json:"Response"`
}

func (r *KillMasterGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KillMasterGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogDeliveryInfo struct {
	// Enabling status of log shipping. true: enabled; false: disabled.
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// Log set ID.
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// Log topic ID.
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Logset region
	LogRegion *string `json:"LogRegion,omitnil,omitempty" name:"LogRegion"`
}

type LogFilter struct {
	// <p>Filter criterion name.</p><p>Enumeration value:</p><ul><li>Timestamp: Creation time (format: 2006-01-02 15:04:05.000)</li><li>UserName: User name</li><li>CacheCode: Cache code, backend redis node</li><li>ClientAddr: Client IP address</li><li>CommandDetail: Command details</li><li>CommandLatency: Command delay (ms)</li><li>CommandType: Command type</li><li>DBId: Database ID</li><li>ErrMsg: Error information</li></ul>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Condition match type.</p><p>Enumeration value:</p><ul><li>INC: Include, multiple values have a || relationship before</li><li>EXC: Exclude, multiple values have a || relationship before</li><li>EQS: Equal, multiple values have a || relationship before</li><li>NEQ: Not equal, multiple values have a && relationship before</li><li>RA: Range</li></ul>
	Compare *string `json:"Compare,omitnil,omitempty" name:"Compare"`

	// <p>Matching value of the filter condition. When Compare=RA, for example: ["1-100","200-300"].</p>
	Value []*string `json:"Value,omitnil,omitempty" name:"Value"`
}

type LogInstance struct {
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Log status, create: creating; normal: enabled; close: turning off.</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Is it possible to switch log query - Value: yes - allowed, no - not allowed. This parameter mainly controls migration of existing logs to the log platform for query usage. Only when the status is yes can you call the log API.</p>
	EnableQuery *string `json:"EnableQuery,omitnil,omitempty" name:"EnableQuery"`

	// <p>Start time</p>
	CreateAt *string `json:"CreateAt,omitnil,omitempty" name:"CreateAt"`

	// <p>High frequency storage days</p>
	HighLogExpireDay *int64 `json:"HighLogExpireDay,omitnil,omitempty" name:"HighLogExpireDay"`

	// <p>Infrequent storage days</p>
	LowLogExpireDay *int64 `json:"LowLogExpireDay,omitnil,omitempty" name:"LowLogExpireDay"`

	// <p>Total storage duration</p>
	LogExpireDay *int64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// <p>High-frequency storage volume (in MB)</p>
	HighStorage *float64 `json:"HighStorage,omitnil,omitempty" name:"HighStorage"`

	// <p>Infrequent access storage, unit: MB</p>
	LowStorage *float64 `json:"LowStorage,omitnil,omitempty" name:"LowStorage"`

	// <p>Total storage</p>
	LogStorage *float64 `json:"LogStorage,omitnil,omitempty" name:"LogStorage"`

	// <p>Whether to enable delivery: ON, OFF</p>
	Deliver *string `json:"Deliver,omitnil,omitempty" name:"Deliver"`

	// <p>Log shipping information</p>
	DeliverSummary []*DeliverSummary `json:"DeliverSummary,omitnil,omitempty" name:"DeliverSummary"`

	// <p>Instance-related information on the business side varies according to business and returns different information.</p>
	InstanceInfo *InstanceInfo `json:"InstanceInfo,omitnil,omitempty" name:"InstanceInfo"`

	// <p>Audit sub-type.</p>
	LogSubType *string `json:"LogSubType,omitnil,omitempty" name:"LogSubType"`
}

type LogResult struct {
	// <p>Database ID</p>
	DBId *int64 `json:"DBId,omitnil,omitempty" name:"DBId"`

	// <p>Command delay (ms)</p>
	CommandLatency *int64 `json:"CommandLatency,omitnil,omitempty" name:"CommandLatency"`

	// <p>Creation time (Format: 2006-01-02 15:04:05.000)</p>
	Timestamp *string `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// <p>client address</p>
	ClientAddr *string `json:"ClientAddr,omitnil,omitempty" name:"ClientAddr"`

	// <p>Username.</p>
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// <p>Command type</p>
	CommandType *string `json:"CommandType,omitnil,omitempty" name:"CommandType"`

	// <p>Cache code, backend redis node</p>
	CacheCode *string `json:"CacheCode,omitnil,omitempty" name:"CacheCode"`

	// <p>Command details</p>
	CommandDetail *string `json:"CommandDetail,omitnil,omitempty" name:"CommandDetail"`

	// <p>Error information</p>
	ErrMsg *string `json:"ErrMsg,omitnil,omitempty" name:"ErrMsg"`
}

// Predefined struct for user
type ManualBackupInstanceRequestParams struct {
	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Remarks for manual backup task
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Retention period of backup data in days.  Default value: 7 days.  Value range: [0,1825].  If the value exceeds 7 days, [submit a ticket](https://console.cloud.tencent.com/workorder/category) for application. - If this parameter is not configured, it will set to be the same as the period of automatic backup retention. - If automatic backup is not set, the default value will be 7 days.
	StorageDays *int64 `json:"StorageDays,omitnil,omitempty" name:"StorageDays"`
}

type ManualBackupInstanceRequest struct {
	*tchttp.BaseRequest
	
	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Remarks for manual backup task
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Retention period of backup data in days.  Default value: 7 days.  Value range: [0,1825].  If the value exceeds 7 days, [submit a ticket](https://console.cloud.tencent.com/workorder/category) for application. - If this parameter is not configured, it will set to be the same as the period of automatic backup retention. - If automatic backup is not set, the default value will be 7 days.
	StorageDays *int64 `json:"StorageDays,omitnil,omitempty" name:"StorageDays"`
}

func (r *ManualBackupInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManualBackupInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Remark")
	delete(f, "StorageDays")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ManualBackupInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ManualBackupInstanceResponseParams struct {
	// Task ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ManualBackupInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ManualBackupInstanceResponseParams `json:"Response"`
}

func (r *ManualBackupInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManualBackupInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModfiyInstancePasswordRequestParams struct {
	// Instance ID, such as "crs-xjhsdj****". Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Old password of an instance
	OldPassword *string `json:"OldPassword,omitnil,omitempty" name:"OldPassword"`

	// New instance password. The password complexity requirements are as follows:
	//  - It should contain 8 to 64 characters. 12 or more characters are recommended.
	//  - It cannot start with a forward slash (/).
	//  - It should contain at least two of the following types: lowercase letters (a–z), uppercase letters (A–Z), digits (0–9), and special characters (such as ()~!@#$%^&*-+=_|{}[]:;<>,.?/).
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Whether to encrypt the password.
	EncryptPassword *bool `json:"EncryptPassword,omitnil,omitempty" name:"EncryptPassword"`
}

type ModfiyInstancePasswordRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, such as "crs-xjhsdj****". Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Old password of an instance
	OldPassword *string `json:"OldPassword,omitnil,omitempty" name:"OldPassword"`

	// New instance password. The password complexity requirements are as follows:
	//  - It should contain 8 to 64 characters. 12 or more characters are recommended.
	//  - It cannot start with a forward slash (/).
	//  - It should contain at least two of the following types: lowercase letters (a–z), uppercase letters (A–Z), digits (0–9), and special characters (such as ()~!@#$%^&*-+=_|{}[]:;<>,.?/).
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Whether to encrypt the password.
	EncryptPassword *bool `json:"EncryptPassword,omitnil,omitempty" name:"EncryptPassword"`
}

func (r *ModfiyInstancePasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModfiyInstancePasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "OldPassword")
	delete(f, "Password")
	delete(f, "EncryptPassword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModfiyInstancePasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModfiyInstancePasswordResponseParams struct {
	// Task ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModfiyInstancePasswordResponse struct {
	*tchttp.BaseResponse
	Response *ModfiyInstancePasswordResponseParams `json:"Response"`
}

func (r *ModfiyInstancePasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModfiyInstancePasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAutoBackupConfigRequestParams struct {
	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Automatic backup cycle. Valid values: `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, `Sunday`. This parameter currently cannot be modified.
	WeekDays []*string `json:"WeekDays,omitnil,omitempty" name:"WeekDays"`

	// Automatic backup time in the format of 00:00-01:00, 01:00-02:00... 23:00-00:00.
	TimePeriod *string `json:"TimePeriod,omitnil,omitempty" name:"TimePeriod"`

	// Automatic backup type.  Valid value:  `1` (scheduled backup).
	AutoBackupType *int64 `json:"AutoBackupType,omitnil,omitempty" name:"AutoBackupType"`

	// Retention days for full backup files. Only support setting to 7, unit: day. If needed for longer period, please submit a ticket (https://console.cloud.tencent.com/workorder/category) to apply.
	BackupStorageDays *int64 `json:"BackupStorageDays,omitnil,omitempty" name:"BackupStorageDays"`
}

type ModifyAutoBackupConfigRequest struct {
	*tchttp.BaseRequest
	
	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Automatic backup cycle. Valid values: `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, `Sunday`. This parameter currently cannot be modified.
	WeekDays []*string `json:"WeekDays,omitnil,omitempty" name:"WeekDays"`

	// Automatic backup time in the format of 00:00-01:00, 01:00-02:00... 23:00-00:00.
	TimePeriod *string `json:"TimePeriod,omitnil,omitempty" name:"TimePeriod"`

	// Automatic backup type.  Valid value:  `1` (scheduled backup).
	AutoBackupType *int64 `json:"AutoBackupType,omitnil,omitempty" name:"AutoBackupType"`

	// Retention days for full backup files. Only support setting to 7, unit: day. If needed for longer period, please submit a ticket (https://console.cloud.tencent.com/workorder/category) to apply.
	BackupStorageDays *int64 `json:"BackupStorageDays,omitnil,omitempty" name:"BackupStorageDays"`
}

func (r *ModifyAutoBackupConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoBackupConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "WeekDays")
	delete(f, "TimePeriod")
	delete(f, "AutoBackupType")
	delete(f, "BackupStorageDays")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAutoBackupConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAutoBackupConfigResponseParams struct {
	// Automatic backup type.  Valid value:  `1` (scheduled backup).
	AutoBackupType *int64 `json:"AutoBackupType,omitnil,omitempty" name:"AutoBackupType"`

	// Automatic backup cycle. Valid values: `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, `Sunday`.
	WeekDays []*string `json:"WeekDays,omitnil,omitempty" name:"WeekDays"`

	// Time period for automatic scheduled backup  in the format of  “00:00-01:00, 01:00-02:00...... 23:00-00:00”.
	TimePeriod *string `json:"TimePeriod,omitnil,omitempty" name:"TimePeriod"`

	// Retention time of full backup files in days
	BackupStorageDays *int64 `json:"BackupStorageDays,omitnil,omitempty" name:"BackupStorageDays"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAutoBackupConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAutoBackupConfigResponseParams `json:"Response"`
}

func (r *ModifyAutoBackupConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoBackupConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupDownloadRestrictionRequestParams struct {
	// Type of the network restrictions for downloading backup files. Valid values:
	// 
	// - `NoLimit`: Backup files can be downloaded over both public and private networks.
	// - `LimitOnlyIntranet`: Backup files can be downloaded only at private network addresses auto-assigned by Tencent Cloud.
	// - `Customize`: Backup files can be downloaded only in the customized VPC.
	LimitType *string `json:"LimitType,omitnil,omitempty" name:"LimitType"`

	// Only `In` can be passed in for this parameter, indicating that backup files can be downloaded in the custom `LimitVpc`.
	VpcComparisonSymbol *string `json:"VpcComparisonSymbol,omitnil,omitempty" name:"VpcComparisonSymbol"`

	// Whether backups can be downloaded at the custom `LimitIp` address.
	// 
	// - `In`: Download is allowed for the custom IP.
	// - `NotIn`: Download is not allowed for the custom IP.
	IpComparisonSymbol *string `json:"IpComparisonSymbol,omitnil,omitempty" name:"IpComparisonSymbol"`

	// VPC ID of the custom backup file download address, which is required if `LimitType` is `Customize`.
	LimitVpc []*BackupLimitVpcItem `json:"LimitVpc,omitnil,omitempty" name:"LimitVpc"`

	// VPC IP of the custom backup file download address, which is required if `LimitType` is `Customize`.
	LimitIp []*string `json:"LimitIp,omitnil,omitempty" name:"LimitIp"`
}

type ModifyBackupDownloadRestrictionRequest struct {
	*tchttp.BaseRequest
	
	// Type of the network restrictions for downloading backup files. Valid values:
	// 
	// - `NoLimit`: Backup files can be downloaded over both public and private networks.
	// - `LimitOnlyIntranet`: Backup files can be downloaded only at private network addresses auto-assigned by Tencent Cloud.
	// - `Customize`: Backup files can be downloaded only in the customized VPC.
	LimitType *string `json:"LimitType,omitnil,omitempty" name:"LimitType"`

	// Only `In` can be passed in for this parameter, indicating that backup files can be downloaded in the custom `LimitVpc`.
	VpcComparisonSymbol *string `json:"VpcComparisonSymbol,omitnil,omitempty" name:"VpcComparisonSymbol"`

	// Whether backups can be downloaded at the custom `LimitIp` address.
	// 
	// - `In`: Download is allowed for the custom IP.
	// - `NotIn`: Download is not allowed for the custom IP.
	IpComparisonSymbol *string `json:"IpComparisonSymbol,omitnil,omitempty" name:"IpComparisonSymbol"`

	// VPC ID of the custom backup file download address, which is required if `LimitType` is `Customize`.
	LimitVpc []*BackupLimitVpcItem `json:"LimitVpc,omitnil,omitempty" name:"LimitVpc"`

	// VPC IP of the custom backup file download address, which is required if `LimitType` is `Customize`.
	LimitIp []*string `json:"LimitIp,omitnil,omitempty" name:"LimitIp"`
}

func (r *ModifyBackupDownloadRestrictionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupDownloadRestrictionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LimitType")
	delete(f, "VpcComparisonSymbol")
	delete(f, "IpComparisonSymbol")
	delete(f, "LimitVpc")
	delete(f, "LimitIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBackupDownloadRestrictionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupDownloadRestrictionResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyBackupDownloadRestrictionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBackupDownloadRestrictionResponseParams `json:"Response"`
}

func (r *ModifyBackupDownloadRestrictionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupDownloadRestrictionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyConnectionConfigRequestParams struct {
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Additional bandwidth, in MB, which should be greater than 0.
	// **Note**: The Bandwidth and ClientLimit parameters cannot be empty simultaneously. You should select at least one of them for configuration.
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// Total number of connections per shard.
	// - When read-only replicas are not enabled, the lower limit is 10,000 and the upper limit is 40,000.
	// - When read-only replicas are enabled, the lower limit is 10,000, and the upper limit is calculated as follows: 10,000 x (Number of read-only replicas + 3).
	// **Note**: The Bandwidth and ClientLimit parameters cannot be empty simultaneously. You should select at least one of them for configuration.
	ClientLimit *int64 `json:"ClientLimit,omitnil,omitempty" name:"ClientLimit"`
}

type ModifyConnectionConfigRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Additional bandwidth, in MB, which should be greater than 0.
	// **Note**: The Bandwidth and ClientLimit parameters cannot be empty simultaneously. You should select at least one of them for configuration.
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// Total number of connections per shard.
	// - When read-only replicas are not enabled, the lower limit is 10,000 and the upper limit is 40,000.
	// - When read-only replicas are enabled, the lower limit is 10,000, and the upper limit is calculated as follows: 10,000 x (Number of read-only replicas + 3).
	// **Note**: The Bandwidth and ClientLimit parameters cannot be empty simultaneously. You should select at least one of them for configuration.
	ClientLimit *int64 `json:"ClientLimit,omitnil,omitempty" name:"ClientLimit"`
}

func (r *ModifyConnectionConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConnectionConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Bandwidth")
	delete(f, "ClientLimit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyConnectionConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyConnectionConfigResponseParams struct {
	// Task ID.
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyConnectionConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyConnectionConfigResponseParams `json:"Response"`
}

func (r *ModifyConnectionConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConnectionConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSecurityGroupsRequestParams struct {
	// Database engine name, which is `redis` for this API.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// Replaces with the new security group ID list, which is an array of one or more security group IDs.
	// - To configure a security group for an instance for the first time, call the API [AssociateSecurityGroups](https://www.tencentcloud.com/document/product/239/41260?from_cn_redirect=1) to bind a security group first.
	// - To replace the security group, obtain the security group ID on the [security group](https://console.tencentcloud.com/vpc/security-group) page of the console.
	// 
	// **Note:** This input parameter performs a full replacement on all existing collections but not an incremental update. To modify it, import the expected full collections.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type ModifyDBInstanceSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Database engine name, which is `redis` for this API.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// Replaces with the new security group ID list, which is an array of one or more security group IDs.
	// - To configure a security group for an instance for the first time, call the API [AssociateSecurityGroups](https://www.tencentcloud.com/document/product/239/41260?from_cn_redirect=1) to bind a security group first.
	// - To replace the security group, obtain the security group ID on the [security group](https://console.tencentcloud.com/vpc/security-group) page of the console.
	// 
	// **Note:** This input parameter performs a full replacement on all existing collections but not an incremental update. To modify it, import the expected full collections.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	delete(f, "Product")
	delete(f, "SecurityGroupIds")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSecurityGroupsResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
type ModifyInstanceAccountRequestParams struct {
	// <p>Instance ID. Log in to the <a href="https://console.cloud.tencent.com/redis/instance/list">Redis console</a> and copy the instance ID from the instance list.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Specify the account that needs modification.</p><ul><li>root: refers to the automatically generated account when a Redis Database Instance is created. Users cannot modify its read-write permissions, but can only modify its request routing strategy.</li><li>Custom account: an account manually created by users once an instance is created successfully. Users can modify its read and write permissions and request routing strategy at any time.</li></ul>
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// <p>Specifies the access password for the account to be modified.</p>
	AccountPassword *string `json:"AccountPassword,omitnil,omitempty" name:"AccountPassword"`

	// <p>Account description.</p>
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// <p>Specify the policy for request routing of read-write requests for the modified account.</p><ul><li>master: means read-write requests are routed to the primary node.</li><li>replication: means read-write requests are routed to the secondary node.</li></ul>
	ReadonlyPolicy []*string `json:"ReadonlyPolicy,omitnil,omitempty" name:"ReadonlyPolicy"`

	// <p>Specify the read/write permission of the account to be modified.</p><ul><li>r: Read-only.</li><li>w: Write-only.</li><li>rw: Read-write.</li></ul>
	Privilege *string `json:"Privilege,omitnil,omitempty" name:"Privilege"`

	// <p>Specifies whether to set the default account (root) to a password-free account. Custom accounts do not support password-free access.</p><ul><li>true: The default account (root) is set to a password-free account.</li><li>false: The default account (root) is not set to a password-free account.</li></ul>
	NoAuth *bool `json:"NoAuth,omitnil,omitempty" name:"NoAuth"`

	// <p>Whether to enable password encryption for transmission.</p><ul><li>true: Encrypted.</li><li>false: Not encrypted (default value).</li></ul>
	EncryptPassword *bool `json:"EncryptPassword,omitnil,omitempty" name:"EncryptPassword"`
}

type ModifyInstanceAccountRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance ID. Log in to the <a href="https://console.cloud.tencent.com/redis/instance/list">Redis console</a> and copy the instance ID from the instance list.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Specify the account that needs modification.</p><ul><li>root: refers to the automatically generated account when a Redis Database Instance is created. Users cannot modify its read-write permissions, but can only modify its request routing strategy.</li><li>Custom account: an account manually created by users once an instance is created successfully. Users can modify its read and write permissions and request routing strategy at any time.</li></ul>
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// <p>Specifies the access password for the account to be modified.</p>
	AccountPassword *string `json:"AccountPassword,omitnil,omitempty" name:"AccountPassword"`

	// <p>Account description.</p>
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// <p>Specify the policy for request routing of read-write requests for the modified account.</p><ul><li>master: means read-write requests are routed to the primary node.</li><li>replication: means read-write requests are routed to the secondary node.</li></ul>
	ReadonlyPolicy []*string `json:"ReadonlyPolicy,omitnil,omitempty" name:"ReadonlyPolicy"`

	// <p>Specify the read/write permission of the account to be modified.</p><ul><li>r: Read-only.</li><li>w: Write-only.</li><li>rw: Read-write.</li></ul>
	Privilege *string `json:"Privilege,omitnil,omitempty" name:"Privilege"`

	// <p>Specifies whether to set the default account (root) to a password-free account. Custom accounts do not support password-free access.</p><ul><li>true: The default account (root) is set to a password-free account.</li><li>false: The default account (root) is not set to a password-free account.</li></ul>
	NoAuth *bool `json:"NoAuth,omitnil,omitempty" name:"NoAuth"`

	// <p>Whether to enable password encryption for transmission.</p><ul><li>true: Encrypted.</li><li>false: Not encrypted (default value).</li></ul>
	EncryptPassword *bool `json:"EncryptPassword,omitnil,omitempty" name:"EncryptPassword"`
}

func (r *ModifyInstanceAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AccountName")
	delete(f, "AccountPassword")
	delete(f, "Remark")
	delete(f, "ReadonlyPolicy")
	delete(f, "Privilege")
	delete(f, "NoAuth")
	delete(f, "EncryptPassword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceAccountResponseParams struct {
	// <p>Task ID.</p>
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceAccountResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceAccountResponseParams `json:"Response"`
}

func (r *ModifyInstanceAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceAvailabilityZonesRequestParams struct {
	// Specify the instance ID.
	//  For example: crs-xjhsdj****, please log in to the [Redis Console](https://console.tencentcloud.com/redis) and copy the instance ID from the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Switch time.
	// - 1: Switch during the maintenance window.
	// - 2: Switch immediately.
	SwitchOption *int64 `json:"SwitchOption,omitnil,omitempty" name:"SwitchOption"`

	// Instance node information includes the node ID, node type, and node availability zone ID, and so on. For specific information, please see [RedisNodeInfo ](https://intl.cloud.tencent.com/document/product/239/20022?from_cn_redirect=1).
	// For instances in a single availability zone, there is no need to configure the NodeId. For instances in multiple availability zones, the NodeId is required to configure.
	NodeSet []*RedisNodeInfo `json:"NodeSet,omitnil,omitempty" name:"NodeSet"`
}

type ModifyInstanceAvailabilityZonesRequest struct {
	*tchttp.BaseRequest
	
	// Specify the instance ID.
	//  For example: crs-xjhsdj****, please log in to the [Redis Console](https://console.tencentcloud.com/redis) and copy the instance ID from the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Switch time.
	// - 1: Switch during the maintenance window.
	// - 2: Switch immediately.
	SwitchOption *int64 `json:"SwitchOption,omitnil,omitempty" name:"SwitchOption"`

	// Instance node information includes the node ID, node type, and node availability zone ID, and so on. For specific information, please see [RedisNodeInfo ](https://intl.cloud.tencent.com/document/product/239/20022?from_cn_redirect=1).
	// For instances in a single availability zone, there is no need to configure the NodeId. For instances in multiple availability zones, the NodeId is required to configure.
	NodeSet []*RedisNodeInfo `json:"NodeSet,omitnil,omitempty" name:"NodeSet"`
}

func (r *ModifyInstanceAvailabilityZonesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceAvailabilityZonesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SwitchOption")
	delete(f, "NodeSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceAvailabilityZonesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceAvailabilityZonesResponseParams struct {
	// Task ID.
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceAvailabilityZonesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceAvailabilityZonesResponseParams `json:"Response"`
}

func (r *ModifyInstanceAvailabilityZonesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceAvailabilityZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceBackupModeRequestParams struct {
	// Instance ID. Log in to the [Redis console](https://console.cloud.tencent.com/redis/instance/list), and copy it from the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup mode:
	// - SecondLevelBackup: second-level backup.
	// - NormalLevelBackup: ordinary backup.
	BackupMode *string `json:"BackupMode,omitnil,omitempty" name:"BackupMode"`
}

type ModifyInstanceBackupModeRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. Log in to the [Redis console](https://console.cloud.tencent.com/redis/instance/list), and copy it from the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup mode:
	// - SecondLevelBackup: second-level backup.
	// - NormalLevelBackup: ordinary backup.
	BackupMode *string `json:"BackupMode,omitnil,omitempty" name:"BackupMode"`
}

func (r *ModifyInstanceBackupModeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceBackupModeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceBackupModeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceBackupModeResponseParams struct {
	// Task ID.
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceBackupModeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceBackupModeResponseParams `json:"Response"`
}

func (r *ModifyInstanceBackupModeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceBackupModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceChargeTypeRequestParams struct {
	// <p>Instance ID Array</p><p>Input parameter limitation: Length of batch operation array not exceeding 20</p>
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// <p>Operation type for billing mode change</p><p>Enumeration value:</p><ul><li>PREPAID: Transition from pay-as-you-go to monthly subscription</li><li>POSTPAID: Monthly subscription to pay-as-you-go</li></ul>
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// <p>Purchase duration is required only when InstanceChargeType=PREPAID.</p><p>Valid values: 1 to 36.</p><p>Unit: months.</p>
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`
}

type ModifyInstanceChargeTypeRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance ID Array</p><p>Input parameter limitation: Length of batch operation array not exceeding 20</p>
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// <p>Operation type for billing mode change</p><p>Enumeration value:</p><ul><li>PREPAID: Transition from pay-as-you-go to monthly subscription</li><li>POSTPAID: Monthly subscription to pay-as-you-go</li></ul>
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// <p>Purchase duration is required only when InstanceChargeType=PREPAID.</p><p>Valid values: 1 to 36.</p><p>Unit: months.</p>
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`
}

func (r *ModifyInstanceChargeTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceChargeTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "InstanceChargeType")
	delete(f, "Period")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceChargeTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceChargeTypeResponseParams struct {
	// <p>Summary of failed instance modification info</p>
	FailedInstanceIds []*FailedInstance `json:"FailedInstanceIds,omitnil,omitempty" name:"FailedInstanceIds"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceChargeTypeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceChargeTypeResponseParams `json:"Response"`
}

func (r *ModifyInstanceChargeTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceChargeTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceEventRequestParams struct {
	// Specifies the instance ID. Example: crs-xjhsdj****. Log in to the [TencentDB for Redis console](https://console.cloud.tencent.com/redis#/) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Event ID. Call the [DescribeInstanceEvents](https://www.tencentcloud.com/document/product/239/104779?from_cn_redirect=1) API to obtain the ID of the event to be modified.
	EventId *int64 `json:"EventId,omitnil,omitempty" name:"EventId"`

	// Modifies the scheduled start time of event execution.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Modifies the scheduled end time of event execution. After the start time is configured, the end time can only be 30 minutes, 1 hour, 1.5 hours, 2 hours, or 3 hours later than the start time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Modifies the start date of the event execution schedule.
	ExecutionDate *string `json:"ExecutionDate,omitnil,omitempty" name:"ExecutionDate"`

	// Modifies the running status of the event. Currently, this parameter can be set only to **Canceled**, indicating that the execution of the current event is canceled. You can query the running status and level of the current event using DescribeInstanceEvents.- Critical or High events cannot be canceled, which means that they must be executed.- Only events in the Waiting state (to be executed) can be canceled.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type ModifyInstanceEventRequest struct {
	*tchttp.BaseRequest
	
	// Specifies the instance ID. Example: crs-xjhsdj****. Log in to the [TencentDB for Redis console](https://console.cloud.tencent.com/redis#/) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Event ID. Call the [DescribeInstanceEvents](https://www.tencentcloud.com/document/product/239/104779?from_cn_redirect=1) API to obtain the ID of the event to be modified.
	EventId *int64 `json:"EventId,omitnil,omitempty" name:"EventId"`

	// Modifies the scheduled start time of event execution.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Modifies the scheduled end time of event execution. After the start time is configured, the end time can only be 30 minutes, 1 hour, 1.5 hours, 2 hours, or 3 hours later than the start time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Modifies the start date of the event execution schedule.
	ExecutionDate *string `json:"ExecutionDate,omitnil,omitempty" name:"ExecutionDate"`

	// Modifies the running status of the event. Currently, this parameter can be set only to **Canceled**, indicating that the execution of the current event is canceled. You can query the running status and level of the current event using DescribeInstanceEvents.- Critical or High events cannot be canceled, which means that they must be executed.- Only events in the Waiting state (to be executed) can be canceled.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *ModifyInstanceEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "EventId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ExecutionDate")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceEventResponseParams struct {
	// Event ID.
	EventId *int64 `json:"EventId,omitnil,omitempty" name:"EventId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceEventResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceEventResponseParams `json:"Response"`
}

func (r *ModifyInstanceEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceLogDeliveryRequestParams struct {
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Log type. Currently, only slowlog is supported, indicating the slow query log.
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// Whether log shopping is enabled.
	// - true: enabled.
	// - false: disabled.
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// ID of the shipped logset. It can be obtained through the API [DescribeLogsets](https://intl.cloud.tencent.com/document/api/614/58624?from_cn_redirect=1).
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// ID of the shipped log topic. It can be obtained through the API [DescribeTopics](https://intl.cloud.tencent.com/document/api/614/56454?from_cn_redirect=1).
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Logset name. It is required when **LogsetId** is left blank. The system will create a logset with the value of LogsetName and ship logs.
	LogsetName *string `json:"LogsetName,omitnil,omitempty" name:"LogsetName"`

	// Log topic name. It is required when **TopicId** is left blank. The system will create a log topic with the value of TopicName and ship logs.
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Region where the logset is located. If it is not provided, the region where the instance is located will be used by default.
	LogRegion *string `json:"LogRegion,omitnil,omitempty" name:"LogRegion"`

	// Log storage duration. Default value: 30 days. Value range: 1 to 3600 days.
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Whether to create an index when creating a log topic.
	CreateIndex *bool `json:"CreateIndex,omitnil,omitempty" name:"CreateIndex"`
}

type ModifyInstanceLogDeliveryRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Log type. Currently, only slowlog is supported, indicating the slow query log.
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// Whether log shopping is enabled.
	// - true: enabled.
	// - false: disabled.
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// ID of the shipped logset. It can be obtained through the API [DescribeLogsets](https://intl.cloud.tencent.com/document/api/614/58624?from_cn_redirect=1).
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// ID of the shipped log topic. It can be obtained through the API [DescribeTopics](https://intl.cloud.tencent.com/document/api/614/56454?from_cn_redirect=1).
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Logset name. It is required when **LogsetId** is left blank. The system will create a logset with the value of LogsetName and ship logs.
	LogsetName *string `json:"LogsetName,omitnil,omitempty" name:"LogsetName"`

	// Log topic name. It is required when **TopicId** is left blank. The system will create a log topic with the value of TopicName and ship logs.
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Region where the logset is located. If it is not provided, the region where the instance is located will be used by default.
	LogRegion *string `json:"LogRegion,omitnil,omitempty" name:"LogRegion"`

	// Log storage duration. Default value: 30 days. Value range: 1 to 3600 days.
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Whether to create an index when creating a log topic.
	CreateIndex *bool `json:"CreateIndex,omitnil,omitempty" name:"CreateIndex"`
}

func (r *ModifyInstanceLogDeliveryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceLogDeliveryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "LogType")
	delete(f, "Enabled")
	delete(f, "LogsetId")
	delete(f, "TopicId")
	delete(f, "LogsetName")
	delete(f, "TopicName")
	delete(f, "LogRegion")
	delete(f, "Period")
	delete(f, "CreateIndex")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceLogDeliveryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceLogDeliveryResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceLogDeliveryResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceLogDeliveryResponseParams `json:"Response"`
}

func (r *ModifyInstanceLogDeliveryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceLogDeliveryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceParamsRequestParams struct {
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// List of instance parameters modified
	InstanceParams []*InstanceParam `json:"InstanceParams,omitnil,omitempty" name:"InstanceParams"`
}

type ModifyInstanceParamsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// List of instance parameters modified
	InstanceParams []*InstanceParam `json:"InstanceParams,omitnil,omitempty" name:"InstanceParams"`
}

func (r *ModifyInstanceParamsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceParamsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InstanceParams")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceParamsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceParamsResponseParams struct {
	// Whether the parameter configuration is successfully modified.<br> <li>true: successful;</li> <li>false: failed.</li>
	Changed *bool `json:"Changed,omitnil,omitempty" name:"Changed"`

	// ID of the task
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceParamsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceParamsResponseParams `json:"Response"`
}

func (r *ModifyInstanceParamsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceParamsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstancePasswordRequestParams struct {
	// Specifies the instance ID. Example: crs-xjhsdj****. Log in to the [TencentDB for Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Old password of the instance.
	OldPassword *string `json:"OldPassword,omitnil,omitempty" name:"OldPassword"`

	// New password of the instance. The password complexity requirements are as follows:
	// - It can contain 8 to 30 characters. It is recommended to use a password of more than 12 characters.
	// - It cannot start with a forward slash (/).
	// - It should contain at least two of the following types: lowercase letters, uppercase letters, digits, and special characters (such as ()~!@#$%^&*-+=_|{}[]:;<>,.?/).
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

type ModifyInstancePasswordRequest struct {
	*tchttp.BaseRequest
	
	// Specifies the instance ID. Example: crs-xjhsdj****. Log in to the [TencentDB for Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Old password of the instance.
	OldPassword *string `json:"OldPassword,omitnil,omitempty" name:"OldPassword"`

	// New password of the instance. The password complexity requirements are as follows:
	// - It can contain 8 to 30 characters. It is recommended to use a password of more than 12 characters.
	// - It cannot start with a forward slash (/).
	// - It should contain at least two of the following types: lowercase letters, uppercase letters, digits, and special characters (such as ()~!@#$%^&*-+=_|{}[]:;<>,.?/).
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

func (r *ModifyInstancePasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancePasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "OldPassword")
	delete(f, "Password")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstancePasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstancePasswordResponseParams struct {
	// Task ID.
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstancePasswordResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstancePasswordResponseParams `json:"Response"`
}

func (r *ModifyInstancePasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancePasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceReadOnlyRequestParams struct {
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance input mode.
	// - 0: read/write.
	// - 1: read-only.
	InputMode *string `json:"InputMode,omitnil,omitempty" name:"InputMode"`
}

type ModifyInstanceReadOnlyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance input mode.
	// - 0: read/write.
	// - 1: read-only.
	InputMode *string `json:"InputMode,omitnil,omitempty" name:"InputMode"`
}

func (r *ModifyInstanceReadOnlyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceReadOnlyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InputMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceReadOnlyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceReadOnlyResponseParams struct {
	// Task ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceReadOnlyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceReadOnlyResponseParams `json:"Response"`
}

func (r *ModifyInstanceReadOnlyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceReadOnlyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceRequestParams struct {
	// Instance modification operation. Valid values:
	// - rename: rename an instance.
	// - modifyProject: modify the project to which the instance belongs.
	// - modifyAutoRenew: modify the instance renewal flag.
	// - modifyDeleteProtectionSwitch: modify the instance deletion protection switch status.
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list. The maximum number of instances per request is 10.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// New name of the instance. Only Chinese characters, letters, digits, underscores (_), and delimiters (-) are supported. The length can be up to 60 characters.
	InstanceNames []*string `json:"InstanceNames,omitnil,omitempty" name:"InstanceNames"`

	// Project ID. Log in to the [Project Management](https://console.tencentcloud.com/project) page of the Redis console and copy the project ID in **Project Name**.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Auto-renewal flag.
	// 
	// - 0: default status (manual renewal).
	// - 1: automatic renewal.
	// - 2: no automatic renewal.
	AutoRenews []*int64 `json:"AutoRenews,omitnil,omitempty" name:"AutoRenews"`

	// Deletion protection switch. - 0: disabled by default; - 1: enabled.
	DeleteProtectionSwitches []*int64 `json:"DeleteProtectionSwitches,omitnil,omitempty" name:"DeleteProtectionSwitches"`

	// This parameter is currently being deprecated and can still be used by existing users. It is recommended that new users use InstanceIds.
	//
	// Deprecated: InstanceId is deprecated.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Disused
	//
	// Deprecated: InstanceName is deprecated.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// This parameter has been deprecated.
	//
	// Deprecated: AutoRenew is deprecated.
	AutoRenew *int64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`
}

type ModifyInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance modification operation. Valid values:
	// - rename: rename an instance.
	// - modifyProject: modify the project to which the instance belongs.
	// - modifyAutoRenew: modify the instance renewal flag.
	// - modifyDeleteProtectionSwitch: modify the instance deletion protection switch status.
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list. The maximum number of instances per request is 10.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// New name of the instance. Only Chinese characters, letters, digits, underscores (_), and delimiters (-) are supported. The length can be up to 60 characters.
	InstanceNames []*string `json:"InstanceNames,omitnil,omitempty" name:"InstanceNames"`

	// Project ID. Log in to the [Project Management](https://console.tencentcloud.com/project) page of the Redis console and copy the project ID in **Project Name**.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Auto-renewal flag.
	// 
	// - 0: default status (manual renewal).
	// - 1: automatic renewal.
	// - 2: no automatic renewal.
	AutoRenews []*int64 `json:"AutoRenews,omitnil,omitempty" name:"AutoRenews"`

	// Deletion protection switch. - 0: disabled by default; - 1: enabled.
	DeleteProtectionSwitches []*int64 `json:"DeleteProtectionSwitches,omitnil,omitempty" name:"DeleteProtectionSwitches"`

	// This parameter is currently being deprecated and can still be used by existing users. It is recommended that new users use InstanceIds.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Disused
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// This parameter has been deprecated.
	AutoRenew *int64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`
}

func (r *ModifyInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operation")
	delete(f, "InstanceIds")
	delete(f, "InstanceNames")
	delete(f, "ProjectId")
	delete(f, "AutoRenews")
	delete(f, "DeleteProtectionSwitches")
	delete(f, "InstanceId")
	delete(f, "InstanceName")
	delete(f, "AutoRenew")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceResponseParams `json:"Response"`
}

func (r *ModifyInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLogRequestParams struct {
	// <p>Specify the instance ID. Example: crs-xjhsdj****. Log in to the <a href="https://console.cloud.tencent.com/redis">Redis console</a> and copy the instance ID from the instance list.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Log type.</p><p>Enumeration value:</p><ul><li>auditLog: Audit log.</li></ul>
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// <p>Log subcategory.</p><p>Enumeration value:</p><ul><li>write: Write command.</li><li>read: Read command.</li><li>all: All commands.</li></ul>
	LogSubType *string `json:"LogSubType,omitnil,omitempty" name:"LogSubType"`

	// <p>Log expiration time, unit: day.</p><p>Enumeration value:</p><ul><li>7: 7 days</li><li>30: 30 days</li></ul>
	LogExpireDay *int64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// <p>High-frequency log expiration time, unit: day.</p><p>Enumeration value:</p><ul><li>7: 7 days</li></ul><p>Default value: 7</p>
	HighLogExpireDay *int64 `json:"HighLogExpireDay,omitnil,omitempty" name:"HighLogExpireDay"`

	// <p>Downgrade policy, unit: ms. When the instance P99 reaches the downgrade policy, audit data is automatically discarded to prioritize business availability. Default value: 500 ms.</p><p>Value ranges from 300 to 1000.</p>
	DegradeStrategy *int64 `json:"DegradeStrategy,omitnil,omitempty" name:"DegradeStrategy"`
}

type ModifyLogRequest struct {
	*tchttp.BaseRequest
	
	// <p>Specify the instance ID. Example: crs-xjhsdj****. Log in to the <a href="https://console.cloud.tencent.com/redis">Redis console</a> and copy the instance ID from the instance list.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Log type.</p><p>Enumeration value:</p><ul><li>auditLog: Audit log.</li></ul>
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// <p>Log subcategory.</p><p>Enumeration value:</p><ul><li>write: Write command.</li><li>read: Read command.</li><li>all: All commands.</li></ul>
	LogSubType *string `json:"LogSubType,omitnil,omitempty" name:"LogSubType"`

	// <p>Log expiration time, unit: day.</p><p>Enumeration value:</p><ul><li>7: 7 days</li><li>30: 30 days</li></ul>
	LogExpireDay *int64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// <p>High-frequency log expiration time, unit: day.</p><p>Enumeration value:</p><ul><li>7: 7 days</li></ul><p>Default value: 7</p>
	HighLogExpireDay *int64 `json:"HighLogExpireDay,omitnil,omitempty" name:"HighLogExpireDay"`

	// <p>Downgrade policy, unit: ms. When the instance P99 reaches the downgrade policy, audit data is automatically discarded to prioritize business availability. Default value: 500 ms.</p><p>Value ranges from 300 to 1000.</p>
	DegradeStrategy *int64 `json:"DegradeStrategy,omitnil,omitempty" name:"DegradeStrategy"`
}

func (r *ModifyLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "LogType")
	delete(f, "LogSubType")
	delete(f, "LogExpireDay")
	delete(f, "HighLogExpireDay")
	delete(f, "DegradeStrategy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLogResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLogResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLogResponseParams `json:"Response"`
}

func (r *ModifyLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMaintenanceWindowRequestParams struct {
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time of the maintenance window, for example, 17:00.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time of the maintenance window, for example, 19:00.
	// **Note:** Maintenance window duration. Valid values: 30 minutes, 1 hour, 1.5 hours, 2 hours, and 3 hours.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type ModifyMaintenanceWindowRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time of the maintenance window, for example, 17:00.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time of the maintenance window, for example, 19:00.
	// **Note:** Maintenance window duration. Valid values: 30 minutes, 1 hour, 1.5 hours, 2 hours, and 3 hours.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *ModifyMaintenanceWindowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMaintenanceWindowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMaintenanceWindowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMaintenanceWindowResponseParams struct {
	// Modification status. Valid values: success, failed.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyMaintenanceWindowResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMaintenanceWindowResponseParams `json:"Response"`
}

func (r *ModifyMaintenanceWindowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMaintenanceWindowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNetworkConfigRequestParams struct {
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Network change type. Valid values:
	// - `changeVip`: VPC change, including the private IPv4 address and port.
	// - `changeVpc`: Subnet change.
	// - `changeBaseToVpc`: Change from classic network to VPC.
	// - `changeVPort`: Port change.
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// Private IPv4 address of the instance, which is required if `Operation` is `changeVip`.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// VPC ID after the modification.
	// - Configure this parameter when **Operation** is set to **changeVpc** or **changeBaseToVpc**.
	// - Log in to the [Redis console](https://console.tencentcloud.com/redis/instance), switch to the **Instance Details** page, and click the VPC name next to the associated network in the **Network Information** area to obtain the VPC ID.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// ID of the subnet to which the VPC belongs after the modification.
	// - Configure this parameter when **Operation** is set to **changeVpc** or **changeBaseToVpc**.
	// - Log in to the [Redis console](https://console.tencentcloud.com/redis/instance), switch to the **Instance Details** page, and click the subnet name next to the associated network in the **Network Information** area to obtain the subnet ID.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Retention duration of the original private IPv4 address.
	// - Unit: day.
	// - Valid values: 0, 1, 2, 3, 7, and 15.
	// **Note**: If the retention duration is not set or set to 0, the original network address will be released immediately.
	Recycle *int64 `json:"Recycle,omitnil,omitempty" name:"Recycle"`

	// Network port after the change, which is required if `Operation` is `changeVPort` or `changeVip`. Value range: [1024,65535].
	VPort *int64 `json:"VPort,omitnil,omitempty" name:"VPort"`
}

type ModifyNetworkConfigRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Network change type. Valid values:
	// - `changeVip`: VPC change, including the private IPv4 address and port.
	// - `changeVpc`: Subnet change.
	// - `changeBaseToVpc`: Change from classic network to VPC.
	// - `changeVPort`: Port change.
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// Private IPv4 address of the instance, which is required if `Operation` is `changeVip`.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// VPC ID after the modification.
	// - Configure this parameter when **Operation** is set to **changeVpc** or **changeBaseToVpc**.
	// - Log in to the [Redis console](https://console.tencentcloud.com/redis/instance), switch to the **Instance Details** page, and click the VPC name next to the associated network in the **Network Information** area to obtain the VPC ID.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// ID of the subnet to which the VPC belongs after the modification.
	// - Configure this parameter when **Operation** is set to **changeVpc** or **changeBaseToVpc**.
	// - Log in to the [Redis console](https://console.tencentcloud.com/redis/instance), switch to the **Instance Details** page, and click the subnet name next to the associated network in the **Network Information** area to obtain the subnet ID.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Retention duration of the original private IPv4 address.
	// - Unit: day.
	// - Valid values: 0, 1, 2, 3, 7, and 15.
	// **Note**: If the retention duration is not set or set to 0, the original network address will be released immediately.
	Recycle *int64 `json:"Recycle,omitnil,omitempty" name:"Recycle"`

	// Network port after the change, which is required if `Operation` is `changeVPort` or `changeVip`. Value range: [1024,65535].
	VPort *int64 `json:"VPort,omitnil,omitempty" name:"VPort"`
}

func (r *ModifyNetworkConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNetworkConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Operation")
	delete(f, "Vip")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "Recycle")
	delete(f, "VPort")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNetworkConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNetworkConfigResponseParams struct {
	// Execution status. Ignore this parameter.
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// New subnet ID of the instance
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// New VPC ID of the instance
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// New private IPv4 address of the instance
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Task ID. Obtain **taskId** and query the task execution status through the API [DescribeTaskInfo](https://intl.cloud.tencent.com/document/product/239/30601?from_cn_redirect=1).
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNetworkConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNetworkConfigResponseParams `json:"Response"`
}

func (r *ModifyNetworkConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNetworkConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyParamTemplateRequestParams struct {
	// Source parameter template ID, which can be obtained through the response parameter **TemplateId** of the API [DescribeParamTemplateInfo](https://intl.cloud.tencent.com/document/product/239/58748?from_cn_redirect=1).
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// New name after the parameter template is modified.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// New description after the parameter template is modified.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// New parameter list after the parameter template is modified.
	ParamList []*InstanceParam `json:"ParamList,omitnil,omitempty" name:"ParamList"`
}

type ModifyParamTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Source parameter template ID, which can be obtained through the response parameter **TemplateId** of the API [DescribeParamTemplateInfo](https://intl.cloud.tencent.com/document/product/239/58748?from_cn_redirect=1).
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// New name after the parameter template is modified.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// New description after the parameter template is modified.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// New parameter list after the parameter template is modified.
	ParamList []*InstanceParam `json:"ParamList,omitnil,omitempty" name:"ParamList"`
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
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "ParamList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyParamTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyParamTemplateResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
type ModifyReplicationGroupRequestParams struct {
	// Replication group ID. Log in to the [global replication](https://console.tencentcloud.com/redis/replication) page of the Redis console and obtain it.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Replication group name after the modification.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Description of remarks.
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type ModifyReplicationGroupRequest struct {
	*tchttp.BaseRequest
	
	// Replication group ID. Log in to the [global replication](https://console.tencentcloud.com/redis/replication) page of the Redis console and obtain it.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Replication group name after the modification.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Description of remarks.
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *ModifyReplicationGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyReplicationGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "GroupName")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyReplicationGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyReplicationGroupResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyReplicationGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyReplicationGroupResponseParams `json:"Response"`
}

func (r *ModifyReplicationGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyReplicationGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenLogRequestParams struct {
	// <p>Specify the instance ID. Example: crs-xjhsdj****. Log in to the <a href="https://console.cloud.tencent.com/redis">Redis console</a> and copy the instance ID from the instance list.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Log type.</p><p>Enumeration value:</p><ul><li>auditLog: Audit log.</li></ul>
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// <p>Log subcategory.</p><p>Enumeration value:</p><ul><li>write: Write command.</li><li>read: Read command.</li><li>all: Read/write commands.</li></ul>
	LogSubType *string `json:"LogSubType,omitnil,omitempty" name:"LogSubType"`

	// <p>Log valid period, unit: day.</p><p>Enumeration value:</p><ul><li>7: 7 days</li><li>30: 30 days</li></ul><p>Default value: 7</p>
	LogExpireDay *int64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// <p>High-frequency log valid period, unit: day.</p><p>Enumeration value:</p><ul><li>7: 7 days</li></ul><p>Default value: 7</p>
	HighLogExpireDay *int64 `json:"HighLogExpireDay,omitnil,omitempty" name:"HighLogExpireDay"`

	// <p>Log degradation policy threshold. When the P99 latency of an instance reaches this threshold, the system will automatically discard audit log data to ensure service availability.</p><ul><li>Measurement unit: ms.</li><li>Default value: 500.</li><li>Value ranges from 300 to 1000.</li></ul>
	DegradeStrategy *int64 `json:"DegradeStrategy,omitnil,omitempty" name:"DegradeStrategy"`
}

type OpenLogRequest struct {
	*tchttp.BaseRequest
	
	// <p>Specify the instance ID. Example: crs-xjhsdj****. Log in to the <a href="https://console.cloud.tencent.com/redis">Redis console</a> and copy the instance ID from the instance list.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Log type.</p><p>Enumeration value:</p><ul><li>auditLog: Audit log.</li></ul>
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// <p>Log subcategory.</p><p>Enumeration value:</p><ul><li>write: Write command.</li><li>read: Read command.</li><li>all: Read/write commands.</li></ul>
	LogSubType *string `json:"LogSubType,omitnil,omitempty" name:"LogSubType"`

	// <p>Log valid period, unit: day.</p><p>Enumeration value:</p><ul><li>7: 7 days</li><li>30: 30 days</li></ul><p>Default value: 7</p>
	LogExpireDay *int64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// <p>High-frequency log valid period, unit: day.</p><p>Enumeration value:</p><ul><li>7: 7 days</li></ul><p>Default value: 7</p>
	HighLogExpireDay *int64 `json:"HighLogExpireDay,omitnil,omitempty" name:"HighLogExpireDay"`

	// <p>Log degradation policy threshold. When the P99 latency of an instance reaches this threshold, the system will automatically discard audit log data to ensure service availability.</p><ul><li>Measurement unit: ms.</li><li>Default value: 500.</li><li>Value ranges from 300 to 1000.</li></ul>
	DegradeStrategy *int64 `json:"DegradeStrategy,omitnil,omitempty" name:"DegradeStrategy"`
}

func (r *OpenLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "LogType")
	delete(f, "LogSubType")
	delete(f, "LogExpireDay")
	delete(f, "HighLogExpireDay")
	delete(f, "DegradeStrategy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenLogResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OpenLogResponse struct {
	*tchttp.BaseResponse
	Response *OpenLogResponseParams `json:"Response"`
}

func (r *OpenLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenSSLRequestParams struct {
	// <p>Instance ID. Log in to the <a href="https://console.cloud.tencent.com/redis/instance/list">Redis console</a> and copy it from the instance list.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>SSL address type.</p><p>Enumeration values:</p><ul><li>0: Unlimited.</li><li>1: Private IPv4 address.</li><li>2: Private IPv6 address.</li><li>3: Public network.</li><li>-1: Unspecified.</li></ul><p>Default value: 0</p>
	AddressType *int64 `json:"AddressType,omitnil,omitempty" name:"AddressType"`
}

type OpenSSLRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance ID. Log in to the <a href="https://console.cloud.tencent.com/redis/instance/list">Redis console</a> and copy it from the instance list.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>SSL address type.</p><p>Enumeration values:</p><ul><li>0: Unlimited.</li><li>1: Private IPv4 address.</li><li>2: Private IPv6 address.</li><li>3: Public network.</li><li>-1: Unspecified.</li></ul><p>Default value: 0</p>
	AddressType *int64 `json:"AddressType,omitnil,omitempty" name:"AddressType"`
}

func (r *OpenSSLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenSSLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AddressType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenSSLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenSSLResponseParams struct {
	// <p>Task ID.</p>
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OpenSSLResponse struct {
	*tchttp.BaseResponse
	Response *OpenSSLResponseParams `json:"Response"`
}

func (r *OpenSSLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenSSLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Outbound struct {
	// Policy. Valid values: ACCEPT, DROP.
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// All the addresses that the address group ID represents.
	AddressModule *string `json:"AddressModule,omitnil,omitempty" name:"AddressModule"`

	// Source IP or IP address range, such as 192.168.0.0/16.
	CidrIp *string `json:"CidrIp,omitnil,omitempty" name:"CidrIp"`

	// Description.
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// Network protocol, such as UDP and TCP.
	IpProtocol *string `json:"IpProtocol,omitnil,omitempty" name:"IpProtocol"`

	// Port.
	PortRange *string `json:"PortRange,omitnil,omitempty" name:"PortRange"`

	// All the protocols and ports that the service group ID represents.
	ServiceModule *string `json:"ServiceModule,omitnil,omitempty" name:"ServiceModule"`

	// All the addresses that the security group ID represents.
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type ParamTemplateInfo struct {
	// Parameter template ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Parameter template name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Parameter template description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Instance type
	// - `2`: Redis 2.8 Memory Edition (Standard Architecture).
	// - `3`: CKV 3.2 Memory Edition (Standard Architecture).
	// - `4`: CKV 3.2 Memory Edition (Cluster Architecture).
	// - `5`: Redis 2.8 Memory Edition (Standalone).
	// - `6`: Redis 4.0 Memory Edition (Standard Architecture).
	// - `7`: Redis 4.0 Memory Edition (Cluster Architecture).
	// - `8`: Redis 5.0 Memory Edition (Standard Architecture).
	// - `9`: Redis 5.0 Memory Edition (Cluster Architecture).
	// - `15`: Redis 6.2 Memory Edition (Standard Architecture).
	// - `16`: Redis 6.2 Memory Edition (Cluster Architecture).
	ProductType *uint64 `json:"ProductType,omitnil,omitempty" name:"ProductType"`
}

type ParameterDetail struct {
	// Parameter name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Parameter Type
	ParamType *string `json:"ParamType,omitnil,omitempty" name:"ParamType"`

	// Default value of the parameter
	Default *string `json:"Default,omitnil,omitempty" name:"Default"`

	// Parameter description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Current value of the parameter
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`

	// Whether to restart the database for the modified parameters to take effect
	// - `0`: No restart.
	// - `1`: Restart required.
	NeedReboot *int64 `json:"NeedReboot,omitnil,omitempty" name:"NeedReboot"`

	// Maximum parameter value allowed.
	Max *string `json:"Max,omitnil,omitempty" name:"Max"`

	// Minimum parameter value allowed.
	Min *string `json:"Min,omitnil,omitempty" name:"Min"`

	// Optional enumeration values of a parameter. For non-enumeration parameters, it is empty.
	EnumValue []*string `json:"EnumValue,omitnil,omitempty" name:"EnumValue"`
}

type PasswordPolicy struct {
	// <p>Whether to enable the instance-level password complexity policy.</p><ul><li>true: Enable. ALL password changes (create/reset) must pass the complexity verification defined below.</li><li>false: Disable. No complexity filtering is performed.</li></ul><p>Default value: false</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// <p>Minimum number of uppercase and lowercase letters.</p><ul><li>Value ranges from 1 to 16.</li><li>Default value: 1.</li></ul>
	MinLetterCount *int64 `json:"MinLetterCount,omitnil,omitempty" name:"MinLetterCount"`

	// <p>Minimum number of numeric characters.</p><ul><li>Value ranges from 1 to 16.</li><li>Default value: 1.</li></ul>
	MinDigitCount *int64 `json:"MinDigitCount,omitnil,omitempty" name:"MinDigitCount"`

	// <p>Minimum number of special characters.</p><ul><li>Value ranges from 1 to 16.</li><li>Default value: 1.</li></ul>
	MinSpecialCount *int64 `json:"MinSpecialCount,omitnil,omitempty" name:"MinSpecialCount"`

	// <p>Minimum total length of the password (number of characters).</p><ul><li>Value ranges from 8 to 64.</li><li>Default value: 8.</li><li>Constraints and limitations: The minimum total length of the password must be at least the sum of three parameters: MinLetterCount, MinDigitCount, and MinSpecialCount.</li></ul>
	MinLength *int64 `json:"MinLength,omitnil,omitempty" name:"MinLength"`
}

type ProductConf struct {
	// Product type
	// - `2`: Redis 2.8 Memory Edition (Standard Architecture).
	// - `3`: CKV 3.2 Memory Edition (Standard Architecture).
	// - `4`: CKV 3.2 Memory Edition (Cluster Architecture).
	// - `5`: Redis 2.8 Memory Edition (Standalone).
	// - `6`: Redis 4.0 Memory Edition (Standard Architecture).
	// - `7`: Redis 4.0 Memory Edition (Cluster Architecture).
	// - `8`: Redis 5.0 Memory Edition (Standard Architecture).
	// - `9`: Redis 5.0 Memory Edition (Cluster Architecture).
	// - `15`: Redis 6.2 Memory Edition (Standard Architecture).
	// - `16`: Redis 6.2 Memory Edition (Cluster Architecture).
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Product names, including Redis Master-Replica Edition, Redis Standalone Edition, Redis 4.0 Cluster Edition, CKV Master-Replica Edition, and CKV Standalone Edition.
	TypeName *string `json:"TypeName,omitnil,omitempty" name:"TypeName"`

	// Minimum purchasable quantity
	MinBuyNum *int64 `json:"MinBuyNum,omitnil,omitempty" name:"MinBuyNum"`

	// Maximum purchasable quantity
	MaxBuyNum *int64 `json:"MaxBuyNum,omitnil,omitempty" name:"MaxBuyNum"`

	// Whether a product is sold out
	// - `true`: Sold out.
	// - `false`: Not sold out.
	Saleout *bool `json:"Saleout,omitnil,omitempty" name:"Saleout"`

	// Product engine. Valid values: Redis and CKV.
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// Compatible versions, including Redis 2.8, 3.2, 4.0, 5.0, and 6.2.
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// Total capacity in GB
	TotalSize []*string `json:"TotalSize,omitnil,omitempty" name:"TotalSize"`

	// Shard size in GB
	ShardSize []*string `json:"ShardSize,omitnil,omitempty" name:"ShardSize"`

	// Quantity of replicas
	ReplicaNum []*string `json:"ReplicaNum,omitnil,omitempty" name:"ReplicaNum"`

	// Quantity of shards
	ShardNum []*string `json:"ShardNum,omitnil,omitempty" name:"ShardNum"`

	// Supported billing modes
	// - `1`: Monthly subscription.
	// - `0`: Pay-as-you-go.
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Due to spelling inconsistency in this parameter name, it is recommended to use the **EnableReplicaReadOnly** parameter instead. Its meaning refers to whether the Read-Only Replica is supported.
	// - true: Supported.
	// - false: Not supported.
	//
	// Deprecated: EnableRepicaReadOnly is deprecated.
	EnableRepicaReadOnly *bool `json:"EnableRepicaReadOnly,omitnil,omitempty" name:"EnableRepicaReadOnly"`

	// Whether read-only replica is supported.
	//  - true: read-only replica supported.
	//  - false: not supported.
	EnableReplicaReadOnly *bool `json:"EnableReplicaReadOnly,omitnil,omitempty" name:"EnableReplicaReadOnly"`
}

type ProxyNodes struct {
	// Node ID.
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// AZ ID.
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

type RedisBackupSet struct {
	// <p>Backup start time.</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>Backup task ID.</p>
	BackupId *string `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// <p>Backup type.</p><ul><li>1: Automatic backup initiated by the system in the wee hours.</li><li>0: Manual backup initiated by the user.</li></ul>
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// <p>Backup status.</p><ul><li>1: Backup is locked by other processes.</li><li>2: Backup is normal, not locked by any processes.</li><li>-1: Backup has expired.</li><li>3: Backup is being exported.</li><li>4: Backup export successful.</li></ul>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Backup remarks.</p>
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// <p>Backup lock status.</p><ul><li>0: Unlocked.</li><li>1: Has been locked.</li></ul>
	Locked *int64 `json:"Locked,omitnil,omitempty" name:"Locked"`

	// <p>Internal field, which can be ignored.</p>
	BackupSize *int64 `json:"BackupSize,omitnil,omitempty" name:"BackupSize"`

	// <p>Internal field, which can be ignored.</p>
	FullBackup *int64 `json:"FullBackup,omitnil,omitempty" name:"FullBackup"`

	// <p>Internal field, which can be ignored.</p>
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Instance name.</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>Local backup region.</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>Backup end time.</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>Backup file type.</p>
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// <p>Backup file expiration time.</p>
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// <p>Whether the backup file is encrypted</p>
	Encrypted *bool `json:"Encrypted,omitnil,omitempty" name:"Encrypted"`
}

type RedisCommonInstanceList struct {
	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// User APPID, which is the unique application ID that matches an account. Some Tencent Cloud products use this APPID.
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Project ID of the instance
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Instance region
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Instance AZ
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Instance VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// VPC subnet ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Instance status information
	// - `1`: Task running.
	// - `2`: Instance running.
	// - `-2`: Instance isolated.
	// - `-3`: Instance being eliminated.
	// - `-4`: Instance eliminated.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Private network IP address of an instance
	Vips []*string `json:"Vips,omitnil,omitempty" name:"Vips"`

	// Instance network port
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// Instance creation time
	Createtime *string `json:"Createtime,omitnil,omitempty" name:"Createtime"`

	// Billing type
	// - `0`: Pay-as-you-go.
	// - `1`: Monthly subscription.
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Network Type
	// - `0`: Classic network.
	// - `1`: VPC.
	NetType *int64 `json:"NetType,omitnil,omitempty" name:"NetType"`
}

type RedisInstanceEvent struct {
	// Event ID.
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Event type. Currently, the type can only be related to instance migration, resource movement, and IDC deletion. This parameter can be set only to **InstanceMigration**.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Event level. The levels are divided into critical, important, medium, and general based on severity and urgency.
	//  - Critical: critical.
	//  - High: important.
	//  - Middle: medium.
	//  - Low.
	Grade *string `json:"Grade,omitnil,omitempty" name:"Grade"`

	// Scheduled event execution date.
	ExecutionDate *string `json:"ExecutionDate,omitnil,omitempty" name:"ExecutionDate"`

	// Start date of scheduled event execution.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End date of scheduled event execution.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Latest execution date of the Ops event. The event should be completed before this date. Otherwise, the business may be affected.
	LatestExecutionDate *string `json:"LatestExecutionDate,omitnil,omitempty" name:"LatestExecutionDate"`

	// Current event status.
	//  - Waiting: event not reached the execution date or not within the maintenance window.
	//  - Running: event within the maintenance window and under maintenance execution.
	//  - Finished: event with maintenance completed.
	// - Canceled: Execution of the event is canceled.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Completion time of the event execution task.
	TaskEndTime *string `json:"TaskEndTime,omitnil,omitempty" name:"TaskEndTime"`

	// Event impact information.
	EffectInfo *string `json:"EffectInfo,omitnil,omitempty" name:"EffectInfo"`

	// Initial scheduled event execution date.
	InitialExecutionDate *string `json:"InitialExecutionDate,omitnil,omitempty" name:"InitialExecutionDate"`
}

type RedisNode struct {
	// Number of keys on Redis nodes
	Keys *int64 `json:"Keys,omitnil,omitempty" name:"Keys"`

	// Slot distribution range for Redis node.  Value range:  0-5460.
	Slot *string `json:"Slot,omitnil,omitempty" name:"Slot"`

	// Node sequence ID
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// Node status
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Node role
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`
}

type RedisNodeInfo struct {
	// Node type. <ul><li>`0`: Master node.</li><li>`1`: Replica node.</li></ul>
	NodeType *int64 `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// ID of the master or replica node <ul><li>This parameter is optional when the [CreateInstances](https://intl.cloud.tencent.com/document/product/239/20026?from_cn_redirect=1) API is used to create a TencentDB for Redis instance, but it is required when the [UpgradeInstance](https://intl.cloud.tencent.com/document/product/239/20013?from_cn_redirect=1) API is used to adjust the configuration of an instance by deleting a replica.  </li><li>You can use the [DescribeInstances](https://intl.cloud.tencent.com/document/product/239/20018?from_cn_redirect=1) API to get the node ID of integer type. </li></ul> </li></ul>
	NodeId *int64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// ID of the AZ of the master or replica node
	ZoneId *uint64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Name of the AZ of the master or replica node
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`
}

type RedisNodes struct {
	// Node ID
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// Node role
	NodeRole *string `json:"NodeRole,omitnil,omitempty" name:"NodeRole"`

	// Shard ID
	ClusterId *int64 `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// AZ ID
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

type RegionConf struct {
	// Region ID
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Region name
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// Region abbreviation
	RegionShortName *string `json:"RegionShortName,omitnil,omitempty" name:"RegionShortName"`

	// Name of the area where a region is located
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// AZ information
	ZoneSet []*ZoneCapacityConf `json:"ZoneSet,omitnil,omitempty" name:"ZoneSet"`
}

// Predefined struct for user
type ReleaseWanAddressRequestParams struct {
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type ReleaseWanAddressRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *ReleaseWanAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseWanAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReleaseWanAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReleaseWanAddressResponseParams struct {
	// Async task ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Status of disabling public network access
	WanStatus *string `json:"WanStatus,omitnil,omitempty" name:"WanStatus"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReleaseWanAddressResponse struct {
	*tchttp.BaseResponse
	Response *ReleaseWanAddressResponseParams `json:"Response"`
}

func (r *ReleaseWanAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseWanAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveReplicationGroupRequestParams struct {
	// Replication group ID. Log in to the [Redis console and go to the global replication](https://console.cloud.tencent.com/redis/replication) page to obtain the ID.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

type RemoveReplicationGroupRequest struct {
	*tchttp.BaseRequest
	
	// Replication group ID. Log in to the [Redis console and go to the global replication](https://console.cloud.tencent.com/redis/replication) page to obtain the ID.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

func (r *RemoveReplicationGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveReplicationGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveReplicationGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveReplicationGroupResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemoveReplicationGroupResponse struct {
	*tchttp.BaseResponse
	Response *RemoveReplicationGroupResponseParams `json:"Response"`
}

func (r *RemoveReplicationGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveReplicationGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveReplicationInstanceRequestParams struct {
	// Replication group ID, for example, crs-rpl-m3zt****. Log in to the [Redis console](https://console.tencentcloud.com/redis/replication) and obtain it in the global replication group list.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Specifies the instance ID. Example: crs-xjhsdj****. Log in to the [TencentDB for Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Data synchronization type.
	// - true: Strong synchronization is required.
	// - false: Strong synchronization is not required, and only the primary instance can be deleted.
	SyncType *bool `json:"SyncType,omitnil,omitempty" name:"SyncType"`
}

type RemoveReplicationInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Replication group ID, for example, crs-rpl-m3zt****. Log in to the [Redis console](https://console.tencentcloud.com/redis/replication) and obtain it in the global replication group list.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Specifies the instance ID. Example: crs-xjhsdj****. Log in to the [TencentDB for Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Data synchronization type.
	// - true: Strong synchronization is required.
	// - false: Strong synchronization is not required, and only the primary instance can be deleted.
	SyncType *bool `json:"SyncType,omitnil,omitempty" name:"SyncType"`
}

func (r *RemoveReplicationInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveReplicationInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "InstanceId")
	delete(f, "SyncType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveReplicationInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveReplicationInstanceResponseParams struct {
	// Asynchronous task ID.
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemoveReplicationInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RemoveReplicationInstanceResponseParams `json:"Response"`
}

func (r *RemoveReplicationInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveReplicationInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewInstanceRequestParams struct {
	// Purchase duration.
	// -Unit: month.
	// - Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, and 36.
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Instance ID. Log in to the [Redis console](https://console.cloud.tencent.com/redis/instance/list) and copy it from the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Flag whether to change billing mode.
	// -The billing mode of the current instance is pay-as-you-go. To convert to yearly/monthly subscription and renew, specify this parameter as <b>prepaid</b>.
	// -The current instance billing mode is yearly/monthly subscription, so this parameter can be left unset.
	ModifyPayMode *string `json:"ModifyPayMode,omitnil,omitempty" name:"ModifyPayMode"`
}

type RenewInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Purchase duration.
	// -Unit: month.
	// - Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, and 36.
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Instance ID. Log in to the [Redis console](https://console.cloud.tencent.com/redis/instance/list) and copy it from the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Flag whether to change billing mode.
	// -The billing mode of the current instance is pay-as-you-go. To convert to yearly/monthly subscription and renew, specify this parameter as <b>prepaid</b>.
	// -The current instance billing mode is yearly/monthly subscription, so this parameter can be left unset.
	ModifyPayMode *string `json:"ModifyPayMode,omitnil,omitempty" name:"ModifyPayMode"`
}

func (r *RenewInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Period")
	delete(f, "InstanceId")
	delete(f, "ModifyPayMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewInstanceResponseParams struct {
	// Transaction ID
	//
	// Deprecated: DealId is deprecated.
	DealId *string `json:"DealId,omitnil,omitempty" name:"DealId"`

	// Order number.
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RenewInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RenewInstanceResponseParams `json:"Response"`
}

func (r *RenewInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplicaGroup struct {
	// Node group ID
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Node group name, which is empty for the master node
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Node AZ ID, such as ap-guangzhou-1
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Node group type. master: primary node; replica: replica node.
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// List of nodes in the node group
	RedisNodes []*RedisNode `json:"RedisNodes,omitnil,omitempty" name:"RedisNodes"`
}

// Predefined struct for user
type ResetPasswordRequestParams struct {
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Reset password. This parameter can be left unspecified when a password-free instance is used.
	// - It should contain 8 to 32 characters. 12 or more characters are recommended.
	// - It cannot start with a forward slash (/).
	// - It should contain at least two of the following types: lowercase letters, uppercase letters, digits, and special characters (such as ()~!@#$%^&*-+=_|{}[]:;<>,.?/).
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Whether to switch to a password-free instance.
	// - false: switch to an instance that requires a password. The default value is false.
	// - true: switch to a password-free instance.
	NoAuth *bool `json:"NoAuth,omitnil,omitempty" name:"NoAuth"`

	// Whether to encrypt the password.
	// - false: non-encrypted password. The default value is false.
	// - true: encrypted password.
	EncryptPassword *bool `json:"EncryptPassword,omitnil,omitempty" name:"EncryptPassword"`
}

type ResetPasswordRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Reset password. This parameter can be left unspecified when a password-free instance is used.
	// - It should contain 8 to 32 characters. 12 or more characters are recommended.
	// - It cannot start with a forward slash (/).
	// - It should contain at least two of the following types: lowercase letters, uppercase letters, digits, and special characters (such as ()~!@#$%^&*-+=_|{}[]:;<>,.?/).
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Whether to switch to a password-free instance.
	// - false: switch to an instance that requires a password. The default value is false.
	// - true: switch to a password-free instance.
	NoAuth *bool `json:"NoAuth,omitnil,omitempty" name:"NoAuth"`

	// Whether to encrypt the password.
	// - false: non-encrypted password. The default value is false.
	// - true: encrypted password.
	EncryptPassword *bool `json:"EncryptPassword,omitnil,omitempty" name:"EncryptPassword"`
}

func (r *ResetPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Password")
	delete(f, "NoAuth")
	delete(f, "EncryptPassword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetPasswordResponseParams struct {
	// Task ID (this parameter is the task ID when changing a password. If you are switching between password-free and password-enabled instance mode, you can leave this parameter alone)
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResetPasswordResponse struct {
	*tchttp.BaseResponse
	Response *ResetPasswordResponseParams `json:"Response"`
}

func (r *ResetPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceBundle struct {
	// Resource package name.
	ResourceBundleName *string `json:"ResourceBundleName,omitnil,omitempty" name:"ResourceBundleName"`

	// Saleable memory. Unit: GB.
	AvailableMemory *int64 `json:"AvailableMemory,omitnil,omitempty" name:"AvailableMemory"`

	// Number of resource packages.
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type ResourceTag struct {
	// Tag key
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// The value corresponding to the tag key
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

// Predefined struct for user
type RestoreInstanceRequestParams struct {
	// ID of the instance to be operated, which can be obtained through the response parameter InstanceId of the [DescribeInstances](https://www.tencentcloud.com/document/product/239/20018?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup ID, which can be obtained through the response parameter RedisBackupSet of the [DescribeInstanceBackups](https://www.tencentcloud.com/document/product/239/20011?from_cn_redirect=1) API.
	BackupId *string `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// Instance password, which needs to be validated during instance restoration (this parameter is not required for password-free instances)
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

type RestoreInstanceRequest struct {
	*tchttp.BaseRequest
	
	// ID of the instance to be operated, which can be obtained through the response parameter InstanceId of the [DescribeInstances](https://www.tencentcloud.com/document/product/239/20018?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup ID, which can be obtained through the response parameter RedisBackupSet of the [DescribeInstanceBackups](https://www.tencentcloud.com/document/product/239/20011?from_cn_redirect=1) API.
	BackupId *string `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// Instance password, which needs to be validated during instance restoration (this parameter is not required for password-free instances)
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

func (r *RestoreInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestoreInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupId")
	delete(f, "Password")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestoreInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestoreInstanceResponseParams struct {
	// Task ID, which can be used to query the task execution status through the `DescribeTaskInfo` API.
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RestoreInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RestoreInstanceResponseParams `json:"Response"`
}

func (r *RestoreInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestoreInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecondLevelBackupMissingTimestamps struct {
	// Start timestamp.
	StartTimeStamp *int64 `json:"StartTimeStamp,omitnil,omitempty" name:"StartTimeStamp"`

	// End timestamp.
	EndTimeStamp *int64 `json:"EndTimeStamp,omitnil,omitempty" name:"EndTimeStamp"`
}

type SecurityGroup struct {
	// Creation time in the format of yyyy-mm-dd hh:mm:ss.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Project ID.
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Security group ID.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// Security group name.
	SecurityGroupName *string `json:"SecurityGroupName,omitnil,omitempty" name:"SecurityGroupName"`

	// Security group remarks.
	SecurityGroupRemark *string `json:"SecurityGroupRemark,omitnil,omitempty" name:"SecurityGroupRemark"`

	// Outbound rule.
	Outbound []*Outbound `json:"Outbound,omitnil,omitempty" name:"Outbound"`

	// Inbound rule.
	Inbound []*Inbound `json:"Inbound,omitnil,omitempty" name:"Inbound"`
}

type SecurityGroupDetail struct {
	// Project ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Security group creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Security group ID
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// Security group name
	SecurityGroupName *string `json:"SecurityGroupName,omitnil,omitempty" name:"SecurityGroupName"`

	// Security group remarks
	SecurityGroupRemark *string `json:"SecurityGroupRemark,omitnil,omitempty" name:"SecurityGroupRemark"`

	// Inbound rules of the security group, which control the access source to the database.
	InboundRule []*SecurityGroupsInboundAndOutbound `json:"InboundRule,omitnil,omitempty" name:"InboundRule"`

	// Security group outbound rule
	OutboundRule []*SecurityGroupsInboundAndOutbound `json:"OutboundRule,omitnil,omitempty" name:"OutboundRule"`
}

type SecurityGroupsInboundAndOutbound struct {
	// Whether the inbound and outbound IP addresses and ports of the database are allowed.
	// - ACCEPT: allowed.
	// - DROP: disallowed.
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// IP address for accessing the database
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// Port number
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// Protocol type
	Proto *string `json:"Proto,omitnil,omitempty" name:"Proto"`
}

type SourceCommand struct {
	// Command name.
	Cmd *string `json:"Cmd,omitnil,omitempty" name:"Cmd"`

	// Number of executions.
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type SourceInfo struct {
	// Source IP
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// Number of connections
	Conn *int64 `json:"Conn,omitnil,omitempty" name:"Conn"`

	// Command
	Cmd *int64 `json:"Cmd,omitnil,omitempty" name:"Cmd"`
}

// Predefined struct for user
type StartupInstanceRequestParams struct {
	// Instance ID. Log in to the [Redis console](https://console.cloud.tencent.com/redis/instance/list), and copy the ID of the instance to be deisolated from the recycle bin.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type StartupInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. Log in to the [Redis console](https://console.cloud.tencent.com/redis/instance/list), and copy the ID of the instance to be deisolated from the recycle bin.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *StartupInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartupInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartupInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartupInstanceResponseParams struct {
	// This parameter has been deprecated. Determine whether the instance has been deisolated based on the status obtained through the instance query API.
	//
	// Deprecated: TaskId is deprecated.
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartupInstanceResponse struct {
	*tchttp.BaseResponse
	Response *StartupInstanceResponseParams `json:"Response"`
}

func (r *StartupInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartupInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchAccessNewInstanceRequestParams struct {
	// Specify the instance ID.
	//  For example: crs-xjhsdj****. Please log in to the[ Redis Colose](https://console.tencentcloud.com/redis) and copy the instance ID from the instance list.
	// Sample value: crs-asdasdas.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type SwitchAccessNewInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Specify the instance ID.
	//  For example: crs-xjhsdj****. Please log in to the[ Redis Colose](https://console.tencentcloud.com/redis) and copy the instance ID from the instance list.
	// Sample value: crs-asdasdas.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *SwitchAccessNewInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchAccessNewInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchAccessNewInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchAccessNewInstanceResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SwitchAccessNewInstanceResponse struct {
	*tchttp.BaseResponse
	Response *SwitchAccessNewInstanceResponseParams `json:"Response"`
}

func (r *SwitchAccessNewInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchAccessNewInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchInstanceVipRequestParams struct {
	// Source instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	SrcInstanceId *string `json:"SrcInstanceId,omitnil,omitempty" name:"SrcInstanceId"`

	// Target instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	DstInstanceId *string `json:"DstInstanceId,omitnil,omitempty" name:"DstInstanceId"`

	// DTS disconnection time between the source instance and target instance. Unit: second. If the DTS disconnection time exceeds TimeDelay, the VIP will not be switched. It is recommended to set an acceptable value based on business needs.
	TimeDelay *int64 `json:"TimeDelay,omitnil,omitempty" name:"TimeDelay"`

	// Whether to force a switch in the case of a DTS disconnection.
	// - 1: Force the switch.
	// - 0: Do not force the switch.
	ForceSwitch *int64 `json:"ForceSwitch,omitnil,omitempty" name:"ForceSwitch"`

	// now: switch now; syncComplete: switch after sync is completed
	SwitchTime *string `json:"SwitchTime,omitnil,omitempty" name:"SwitchTime"`
}

type SwitchInstanceVipRequest struct {
	*tchttp.BaseRequest
	
	// Source instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	SrcInstanceId *string `json:"SrcInstanceId,omitnil,omitempty" name:"SrcInstanceId"`

	// Target instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	DstInstanceId *string `json:"DstInstanceId,omitnil,omitempty" name:"DstInstanceId"`

	// DTS disconnection time between the source instance and target instance. Unit: second. If the DTS disconnection time exceeds TimeDelay, the VIP will not be switched. It is recommended to set an acceptable value based on business needs.
	TimeDelay *int64 `json:"TimeDelay,omitnil,omitempty" name:"TimeDelay"`

	// Whether to force a switch in the case of a DTS disconnection.
	// - 1: Force the switch.
	// - 0: Do not force the switch.
	ForceSwitch *int64 `json:"ForceSwitch,omitnil,omitempty" name:"ForceSwitch"`

	// now: switch now; syncComplete: switch after sync is completed
	SwitchTime *string `json:"SwitchTime,omitnil,omitempty" name:"SwitchTime"`
}

func (r *SwitchInstanceVipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchInstanceVipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SrcInstanceId")
	delete(f, "DstInstanceId")
	delete(f, "TimeDelay")
	delete(f, "ForceSwitch")
	delete(f, "SwitchTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchInstanceVipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchInstanceVipResponseParams struct {
	// Task ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SwitchInstanceVipResponse struct {
	*tchttp.BaseResponse
	Response *SwitchInstanceVipResponseParams `json:"Response"`
}

func (r *SwitchInstanceVipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchInstanceVipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchProxyRequestParams struct {
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Proxy ID of an instance, which can be obtained through NodeId of the response parameter **Proxy** of the API [DescribeInstanceNodeInfo](https://intl.cloud.tencent.com/document/product/239/48603?from_cn_redirect=1).
	ProxyID *string `json:"ProxyID,omitnil,omitempty" name:"ProxyID"`

	// Instance proxy ID list. Call the API [DescribeInstanceNodeInfo](https://www.tencentcloud.com/document/product/239/48603?from_cn_redirect=1) to obtain IDs from **NodeId** in the **Proxy** response parameter.
	ProxyIDList []*string `json:"ProxyIDList,omitnil,omitempty" name:"ProxyIDList"`
}

type SwitchProxyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Proxy ID of an instance, which can be obtained through NodeId of the response parameter **Proxy** of the API [DescribeInstanceNodeInfo](https://intl.cloud.tencent.com/document/product/239/48603?from_cn_redirect=1).
	ProxyID *string `json:"ProxyID,omitnil,omitempty" name:"ProxyID"`

	// Instance proxy ID list. Call the API [DescribeInstanceNodeInfo](https://www.tencentcloud.com/document/product/239/48603?from_cn_redirect=1) to obtain IDs from **NodeId** in the **Proxy** response parameter.
	ProxyIDList []*string `json:"ProxyIDList,omitnil,omitempty" name:"ProxyIDList"`
}

func (r *SwitchProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ProxyID")
	delete(f, "ProxyIDList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchProxyResponseParams struct {
	// Task ID.
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SwitchProxyResponse struct {
	*tchttp.BaseResponse
	Response *SwitchProxyResponseParams `json:"Response"`
}

func (r *SwitchProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TaskInfoDetail struct {
	// Task ID.
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Task start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Task type.
	// 
	// - FLOW_CREATE: "001", indicating instance creation.
	// - FLOW_RESIZE: "002", indicating configuration modification.
	// - FLOW_CLOSE: "003", indicating instance disabling.
	// - FLOW_CLEAN: "004", indicating instance cleanup.
	// - FLOW_STARTUP: "005", indicating instance enabling.
	// - FLOW_DELETE: "006", indicating instance deletion.
	// - FLOW_SETPWD: "007", indicating password reset.
	// - FLOW_EXPORTBACKUP: "009", indicating backup file export.
	// - FLOW_RESTOREBACKUP: "010", indicating backup restoration.
	// - FLOW_BACKUPINSTANCE: "012", indicating instance backup.
	// - FLOW_MIGRATEINSTANCE: "013", indicating instance migration.
	// - FLOW_DELBACKUP: "014", indicating backup deletion.
	// - FLOW_EXCHANGEINSTANCE: "016", indicating instance switch.
	// - FLOW_AUTOBACKUP: "017", indicating automatic instance backup.
	// - FLOW_MIGRATECHECK: "022", indicating migration parameter verification.
	// - FLOW_MIGRATETASK: "023", indicating that data migration is in progress.
	// - FLOW_CLEANDB: "025", indicating database cleanup.
	// - FLOW_CLONEBACKUP: "026": indicating backup cloning.
	// - FLOW_CHANGEVIP: "027", indicating VIP address modification.
	// - FLOW_EXPORSHR: "028", indicating scaling.
	// - FLOW_ADDNODES: "029", indicating node addition (removal).
	// - FLOW_CHANGENET: "031", indicating network type modification.
	// - FLOW_MODIFYINSTACEREADONLY: "033": indicating read-only policy modification.
	// - FLOW_MODIFYINSTANCEPARAMS: "034", indicating instance parameter modification.
	// - FLOW_MODIFYINSTANCEPASSWORDFREE: "035", indicating password-free access settings.
	// - FLOW_SWITCHINSTANCEVIP: "036", indicating instance VIP address switch.
	// - FLOW_MODIFYINSTANCEACCOUNT: "037", indicating instance account modification.
	// - FLOW_MODIFYINSTANCEBANDWIDTH: "038", indicating instance bandwidth modification.
	// - FLOW_ENABLEINSTANCE_REPLICATE: "039", indicating enabling of read-only replica.
	// - FLOW_DISABLEINSTANCE_REPLICATE: "040", indicating disabling of read-only replica.
	// - FLOW_UpgradeArch: "041", indicating instance architecture upgrade from the standard architecture to the cluster architecture.
	// - FLOW_DowngradeArch: "042", indicating instance architecture downgrade from the cluster architecture to the standard architecture.
	// - FLOW_UpgradeVersion: "043", indicating version upgrade.
	// - FLOW_MODIFYCONNECTIONCONFIG: "044", indicating adjustment of the bandwidth and the number of connections.
	// - FLOW_CLEARNETWORK: "045", indicating network change.
	// - FLOW_REMOVE_BACKUP_FILE: "046", indicating backup deletion.
	// - FLOW_UPGRADE_SUPPORT_MULTI_AZ: "047", indicating instance upgrade to multi-AZ deployment.
	// - FLOW_SHUTDOWN_MASTER: "048", indicating fault simulation.
	// - FLOW_CHANGE_REPLICA_TO_MASTER: "049", indicating manual promotion to the primary node.
	// - FLOW_CODE_ADD_REPLICATION_INSTANCE: "050", indicating replication group addition.
	// - FLOW_OPEN_WAN: "052", indicating enabling of public network access.
	// - FLOW_CLOSE_WAN: "053", indicating disabling of public network access.
	// - FLOW_CODE_DELETE_REPLICATION_INSTANCE: "055", indicating replication group unbinding.
	// - FLOW_CODE_CHANGE_MASTER_INSTANCE: "056", indicating switching a replication group instance to the primary instance.
	// - FLOW_CODE_CHANGE_INSTANCE_ROLE: "057", indicating modification of the replication group instance role.
	// - FLOW_MIGRATE_NODE: "058", indicating node migration.
	// - FLOW_SWITCH_NODE: "059", indicating node switch.
	// - FLOW_UPGRADE_SMALL_VERSION: "060", indicating Redis version upgrade.
	// - FLOW_UPGRADE_PROXY_VERSION: "061", indicating proxy version upgrade.
	// - FLOW_MODIFY_INSTANCE_NETWORK: "062", indicating instance network modification.
	// - FLOW_MIGRATE_PROXY_NODE: "063", indicating proxy node migration.
	// - FLOW_MIGRATION_INSTANCE_ZONE: "066", indicating that instance migration to another AZ is in progress.
	// - FLOW_UPGRADE_INSTANCE_CACHE_AND_PROXY: "067", indicating that instance version upgrade is in progress.
	// - FLOW_MODIFY_PROXY_NUM: "069", indicating proxy node addition (removal).
	// - FLOW_MODIFYBACKUPMOD: "070", indicating instance backup mode modification.
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// Instance name.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Task progress.
	Progress *float64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// Task execution end time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Task execution status.
	// 
	// 
	// 
	// 0: initializing the task.
	// 1: executing.
	// 2. completed.
	// 4: failed.
	Result *int64 `json:"Result,omitnil,omitempty" name:"Result"`
}

type TendisNodes struct {
	// Node ID
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// Node role
	NodeRole *string `json:"NodeRole,omitnil,omitempty" name:"NodeRole"`

	// AZ ID.	
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

type TendisSlowLogDetail struct {
	// Execution time
	ExecuteTime *string `json:"ExecuteTime,omitnil,omitempty" name:"ExecuteTime"`

	// Duration of the slow query (ms)
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// Command
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// Command line details
	CommandLine *string `json:"CommandLine,omitnil,omitempty" name:"CommandLine"`

	// Node ID
	Node *string `json:"Node,omitnil,omitempty" name:"Node"`
}

type TradeDealDetail struct {
	// Order ID, which is used when a TencentCloud API is called.
	//
	// Deprecated: DealId is deprecated.
	DealId *string `json:"DealId,omitnil,omitempty" name:"DealId"`

	// Long order ID, which is used when an order issue is submitted for assistance.
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// AZ ID
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Number of instances associated with the order
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// Creator UIN
	Creater *string `json:"Creater,omitnil,omitempty" name:"Creater"`

	// Order creation time
	CreatTime *string `json:"CreatTime,omitnil,omitempty" name:"CreatTime"`

	// Order timeout period
	OverdueTime *string `json:"OverdueTime,omitnil,omitempty" name:"OverdueTime"`

	// Order completion time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Order status. 1: unpaid; 2: paid but not delivered; 3: In delivery; 4: successfully delivered; 5: delivery failed; 6: refunded; 7: order closed; 8: order expired; 9: order invalidated; 10: product invalidated; 11: requested payment rejected; 12: paying
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Order status description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Order actual total price (in cents)
	Price *int64 `json:"Price,omitnil,omitempty" name:"Price"`

	// Instance ID
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

// Predefined struct for user
type UpgradeInstanceRequestParams struct {
	// ID of the instance whose configuration is to be modified. Log in to the [TencentDB for Redis® console](https://console.cloud.tencent.com/Redis/instance/list) and copy the instance ID from the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Refers to the size of memory in each shard after the change.
	// -Unit: MB.
	// -You can only modify one of the parameters MemSize, RedisShardNum, and RedisReplicasNum each time, and cannot include both. When modifying one parameter, you need to manually input the original instance configuration specification for the other two parameters.
	// -When scaling down, the new specifications must be equal to or greater than 1.3 times the used capacity, otherwise execution will fail.
	MemSize *uint64 `json:"MemSize,omitnil,omitempty" name:"MemSize"`

	// Specifies the number of shards after the instance specification modification.
	// -Standard architecture does not require this parameter, while cluster architecture is mandatory.
	// -Cluster architecture. Every time, you can only modify one of the parameters RedisShardNum, MemSize, and RedisReplicasNum. You cannot modify them simultaneously. When modifying one parameter, you need to manually input the original instance configuration specification for the other two parameters.
	RedisShardNum *uint64 `json:"RedisShardNum,omitnil,omitempty" name:"RedisShardNum"`

	// Specifies the number of replicas after the instance specification modification.
	// -Only one of the parameters RedisReplicasNum, MemSize, and RedisShardNum can be modified each time. You cannot include both. When modifying one parameter, the other two parameters require the original instance configuration specification.
	// -When modifying a replica of a multi-AZ instance, you must input NodeSet.
	RedisReplicasNum *uint64 `json:"RedisReplicasNum,omitnil,omitempty" name:"RedisReplicasNum"`

	// Node information set when you add a replica for multi-AZ instances, including the ID and AZ information of the replica. This parameter is not required for non-multi-AZ instances.
	NodeSet []*RedisNodeInfo `json:"NodeSet,omitnil,omitempty" name:"NodeSet"`

	// Switch time.
	//  - 1: Perform the operation within the maintenance window: Specification upgrade is executed within the set maintenance window. Use the API [DescribeMaintenanceWindow](https://intl.cloud.tencent.com/document/product/239/46336?from_cn_redirect=1) to query the time period of the set maintenance window. Replica addition/removal, shard addition/removal, and memory capacity expansion/shrinkage are supported within the maintenance window. Specification upgrade within the maintenance window is being gradually tested and published by region. It is already supported in some regions. For urgent integration in regions that do not support it, [submit a ticket](https://console.cloud.tencent.com/workorder/category) to apply for an allowlist.
	//  -2: Perform the operation immediately: The operation will be performed immediately, without the need to wait for the maintenance window. Operations will be performed immediately by default for the system.
	SwitchOption *uint64 `json:"SwitchOption,omitnil,omitempty" name:"SwitchOption"`
}

type UpgradeInstanceRequest struct {
	*tchttp.BaseRequest
	
	// ID of the instance whose configuration is to be modified. Log in to the [TencentDB for Redis® console](https://console.cloud.tencent.com/Redis/instance/list) and copy the instance ID from the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Refers to the size of memory in each shard after the change.
	// -Unit: MB.
	// -You can only modify one of the parameters MemSize, RedisShardNum, and RedisReplicasNum each time, and cannot include both. When modifying one parameter, you need to manually input the original instance configuration specification for the other two parameters.
	// -When scaling down, the new specifications must be equal to or greater than 1.3 times the used capacity, otherwise execution will fail.
	MemSize *uint64 `json:"MemSize,omitnil,omitempty" name:"MemSize"`

	// Specifies the number of shards after the instance specification modification.
	// -Standard architecture does not require this parameter, while cluster architecture is mandatory.
	// -Cluster architecture. Every time, you can only modify one of the parameters RedisShardNum, MemSize, and RedisReplicasNum. You cannot modify them simultaneously. When modifying one parameter, you need to manually input the original instance configuration specification for the other two parameters.
	RedisShardNum *uint64 `json:"RedisShardNum,omitnil,omitempty" name:"RedisShardNum"`

	// Specifies the number of replicas after the instance specification modification.
	// -Only one of the parameters RedisReplicasNum, MemSize, and RedisShardNum can be modified each time. You cannot include both. When modifying one parameter, the other two parameters require the original instance configuration specification.
	// -When modifying a replica of a multi-AZ instance, you must input NodeSet.
	RedisReplicasNum *uint64 `json:"RedisReplicasNum,omitnil,omitempty" name:"RedisReplicasNum"`

	// Node information set when you add a replica for multi-AZ instances, including the ID and AZ information of the replica. This parameter is not required for non-multi-AZ instances.
	NodeSet []*RedisNodeInfo `json:"NodeSet,omitnil,omitempty" name:"NodeSet"`

	// Switch time.
	//  - 1: Perform the operation within the maintenance window: Specification upgrade is executed within the set maintenance window. Use the API [DescribeMaintenanceWindow](https://intl.cloud.tencent.com/document/product/239/46336?from_cn_redirect=1) to query the time period of the set maintenance window. Replica addition/removal, shard addition/removal, and memory capacity expansion/shrinkage are supported within the maintenance window. Specification upgrade within the maintenance window is being gradually tested and published by region. It is already supported in some regions. For urgent integration in regions that do not support it, [submit a ticket](https://console.cloud.tencent.com/workorder/category) to apply for an allowlist.
	//  -2: Perform the operation immediately: The operation will be performed immediately, without the need to wait for the maintenance window. Operations will be performed immediately by default for the system.
	SwitchOption *uint64 `json:"SwitchOption,omitnil,omitempty" name:"SwitchOption"`
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
	delete(f, "MemSize")
	delete(f, "RedisShardNum")
	delete(f, "RedisReplicasNum")
	delete(f, "NodeSet")
	delete(f, "SwitchOption")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeInstanceResponseParams struct {
	// Order ID
	//
	// Deprecated: DealId is deprecated.
	DealId *string `json:"DealId,omitnil,omitempty" name:"DealId"`

	// Order number.
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
type UpgradeInstanceVersionRequestParams struct {
	// Target instance type after the change, which is the same as the `TypeId` of the [CreateInstances](https://intl.cloud.tencent.com/document/api/239/20026?from_cn_redirect=1) API.
	// - For Redis 4.0 or later, a standard architecture instance can be upgraded to a cluster architecture instance on the same version; for example, you can upgrade from Redis 4.0 Standard Architecture to Redis 4.0 Cluster Architecture.
	// - Cross-version architecture upgrade is not supported; for example, you cannot upgrade from Redis 4.0 Standard Architecture to Redis 5.0 Cluster Architecture.
	// - The architecture of Redis 2.8 cannot be upgraded.
	// - Cluster architecture cannot be downgraded to standard architecture.
	TargetInstanceType *string `json:"TargetInstanceType,omitnil,omitempty" name:"TargetInstanceType"`

	// Switch time. Valid values:
	// - `1`: Switch in the maintenance time.
	// - `2` (default value): Switch now.
	SwitchOption *int64 `json:"SwitchOption,omitnil,omitempty" name:"SwitchOption"`

	// ID of the specified instance, such as `crs-xjhsdj****`. Log in to the [Redis console](https://console.cloud.tencent.com/redis#/), and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type UpgradeInstanceVersionRequest struct {
	*tchttp.BaseRequest
	
	// Target instance type after the change, which is the same as the `TypeId` of the [CreateInstances](https://intl.cloud.tencent.com/document/api/239/20026?from_cn_redirect=1) API.
	// - For Redis 4.0 or later, a standard architecture instance can be upgraded to a cluster architecture instance on the same version; for example, you can upgrade from Redis 4.0 Standard Architecture to Redis 4.0 Cluster Architecture.
	// - Cross-version architecture upgrade is not supported; for example, you cannot upgrade from Redis 4.0 Standard Architecture to Redis 5.0 Cluster Architecture.
	// - The architecture of Redis 2.8 cannot be upgraded.
	// - Cluster architecture cannot be downgraded to standard architecture.
	TargetInstanceType *string `json:"TargetInstanceType,omitnil,omitempty" name:"TargetInstanceType"`

	// Switch time. Valid values:
	// - `1`: Switch in the maintenance time.
	// - `2` (default value): Switch now.
	SwitchOption *int64 `json:"SwitchOption,omitnil,omitempty" name:"SwitchOption"`

	// ID of the specified instance, such as `crs-xjhsdj****`. Log in to the [Redis console](https://console.cloud.tencent.com/redis#/), and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *UpgradeInstanceVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeInstanceVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetInstanceType")
	delete(f, "SwitchOption")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeInstanceVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeInstanceVersionResponseParams struct {
	// Order ID
	//
	// Deprecated: DealId is deprecated.
	DealId *string `json:"DealId,omitnil,omitempty" name:"DealId"`

	// Order number.
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeInstanceVersionResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeInstanceVersionResponseParams `json:"Response"`
}

func (r *UpgradeInstanceVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeInstanceVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeProxyVersionRequestParams struct {
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Current proxy version. Call the [DescribeInstances](https://www.tencentcloud.com/document/product/239/20018?from_cn_redirect=1) API to obtain the current proxy version for the instance.
	CurrentProxyVersion *string `json:"CurrentProxyVersion,omitnil,omitempty" name:"CurrentProxyVersion"`

	// Upgradable Redis version. Call the [DescribeInstances](https://www.tencentcloud.com/document/product/239/20018?from_cn_redirect=1) API to obtain the upgradable Redis version for the instance.
	UpgradeProxyVersion *string `json:"UpgradeProxyVersion,omitnil,omitempty" name:"UpgradeProxyVersion"`

	// Specifies whether to upgrade immediately.
	// - 1: Upgrade immediately.
	// - 0: Upgrade during the maintenance window.
	InstanceTypeUpgradeNow *int64 `json:"InstanceTypeUpgradeNow,omitnil,omitempty" name:"InstanceTypeUpgradeNow"`
}

type UpgradeProxyVersionRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Current proxy version. Call the [DescribeInstances](https://www.tencentcloud.com/document/product/239/20018?from_cn_redirect=1) API to obtain the current proxy version for the instance.
	CurrentProxyVersion *string `json:"CurrentProxyVersion,omitnil,omitempty" name:"CurrentProxyVersion"`

	// Upgradable Redis version. Call the [DescribeInstances](https://www.tencentcloud.com/document/product/239/20018?from_cn_redirect=1) API to obtain the upgradable Redis version for the instance.
	UpgradeProxyVersion *string `json:"UpgradeProxyVersion,omitnil,omitempty" name:"UpgradeProxyVersion"`

	// Specifies whether to upgrade immediately.
	// - 1: Upgrade immediately.
	// - 0: Upgrade during the maintenance window.
	InstanceTypeUpgradeNow *int64 `json:"InstanceTypeUpgradeNow,omitnil,omitempty" name:"InstanceTypeUpgradeNow"`
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
	delete(f, "InstanceId")
	delete(f, "CurrentProxyVersion")
	delete(f, "UpgradeProxyVersion")
	delete(f, "InstanceTypeUpgradeNow")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeProxyVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeProxyVersionResponseParams struct {
	// Async task ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

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

// Predefined struct for user
type UpgradeSmallVersionRequestParams struct {
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Current Redis minor version. For minor version information, see [Upgrading Instance Version](https://www.tencentcloud.com/document/product/239/37710).
	CurrentRedisVersion *string `json:"CurrentRedisVersion,omitnil,omitempty" name:"CurrentRedisVersion"`

	// Upgraded Redis minor version. For minor version information, see [Upgrading Instance Version](https://www.tencentcloud.com/document/product/239/37710).
	UpgradeRedisVersion *string `json:"UpgradeRedisVersion,omitnil,omitempty" name:"UpgradeRedisVersion"`

	// Whether to upgrade immediately.
	// - 1: Upgrade immediately.
	// - 0: Upgrade during the maintenance window.
	InstanceTypeUpgradeNow *int64 `json:"InstanceTypeUpgradeNow,omitnil,omitempty" name:"InstanceTypeUpgradeNow"`
}

type UpgradeSmallVersionRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. Log in to the [Redis console](https://console.tencentcloud.com/redis/instance) and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Current Redis minor version. For minor version information, see [Upgrading Instance Version](https://www.tencentcloud.com/document/product/239/37710).
	CurrentRedisVersion *string `json:"CurrentRedisVersion,omitnil,omitempty" name:"CurrentRedisVersion"`

	// Upgraded Redis minor version. For minor version information, see [Upgrading Instance Version](https://www.tencentcloud.com/document/product/239/37710).
	UpgradeRedisVersion *string `json:"UpgradeRedisVersion,omitnil,omitempty" name:"UpgradeRedisVersion"`

	// Whether to upgrade immediately.
	// - 1: Upgrade immediately.
	// - 0: Upgrade during the maintenance window.
	InstanceTypeUpgradeNow *int64 `json:"InstanceTypeUpgradeNow,omitnil,omitempty" name:"InstanceTypeUpgradeNow"`
}

func (r *UpgradeSmallVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeSmallVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "CurrentRedisVersion")
	delete(f, "UpgradeRedisVersion")
	delete(f, "InstanceTypeUpgradeNow")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeSmallVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeSmallVersionResponseParams struct {
	// Async task ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeSmallVersionResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeSmallVersionResponseParams `json:"Response"`
}

func (r *UpgradeSmallVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeSmallVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeVersionToMultiAvailabilityZonesRequestParams struct {
	// <p>Instance ID. Log in to the <a href="https://console.cloud.tencent.com/redis/instance/list">Redis console</a> and copy it from the instance list.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Whether nearest access is supported after upgrading to multiple availability zones. - true: Supports nearest access. The upgrade process involves upgrading the Proxy version and Redis kernel minor version simultaneously, which may require data migration and take up to several hours. - false: No need to support nearest access. Upgrading to multiple availability zones only involves metadata management migration, with no impact on the service. The upgrade process is usually completed within 3 minutes. Defaults to false.</p>
	UpgradeProxyAndRedisServer *bool `json:"UpgradeProxyAndRedisServer,omitnil,omitempty" name:"UpgradeProxyAndRedisServer"`
}

type UpgradeVersionToMultiAvailabilityZonesRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance ID. Log in to the <a href="https://console.cloud.tencent.com/redis/instance/list">Redis console</a> and copy it from the instance list.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Whether nearest access is supported after upgrading to multiple availability zones. - true: Supports nearest access. The upgrade process involves upgrading the Proxy version and Redis kernel minor version simultaneously, which may require data migration and take up to several hours. - false: No need to support nearest access. Upgrading to multiple availability zones only involves metadata management migration, with no impact on the service. The upgrade process is usually completed within 3 minutes. Defaults to false.</p>
	UpgradeProxyAndRedisServer *bool `json:"UpgradeProxyAndRedisServer,omitnil,omitempty" name:"UpgradeProxyAndRedisServer"`
}

func (r *UpgradeVersionToMultiAvailabilityZonesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeVersionToMultiAvailabilityZonesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UpgradeProxyAndRedisServer")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeVersionToMultiAvailabilityZonesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeVersionToMultiAvailabilityZonesResponseParams struct {
	// <p>Task ID.</p>
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeVersionToMultiAvailabilityZonesResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeVersionToMultiAvailabilityZonesResponseParams `json:"Response"`
}

func (r *UpgradeVersionToMultiAvailabilityZonesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeVersionToMultiAvailabilityZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ZoneCapacityConf struct {
	// AZ ID, such as ap-guangzhou-3
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Availability zone name.
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// Whether the AZ is sold out.
	IsSaleout *bool `json:"IsSaleout,omitnil,omitempty" name:"IsSaleout"`

	// Whether the default AZ is used.
	IsDefault *bool `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`

	// Network type.
	// 
	// - basenet: basic network.
	// - vpcnet: VPC.
	NetWorkType []*string `json:"NetWorkType,omitnil,omitempty" name:"NetWorkType"`

	// Information of an AZ, such as product specifications in it
	ProductSet []*ProductConf `json:"ProductSet,omitnil,omitempty" name:"ProductSet"`

	// AZ ID, such as 100003
	OldZoneId *int64 `json:"OldZoneId,omitnil,omitempty" name:"OldZoneId"`
}