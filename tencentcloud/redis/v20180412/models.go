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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
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
	ReadonlyPolicy []*string `json:"ReadonlyPolicy,omitempty" name:"ReadonlyPolicy"`

	// Sub-account status. 1: account is being changed; 2: account is valid; -4: account has been deleted
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

// Predefined struct for user
type AllocateWanAddressRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type AllocateWanAddressRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// Status of enabling public network access
	WanStatus *string `json:"WanStatus,omitempty" name:"WanStatus"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// List of instance IDs
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// ID of the parameter template to be applied
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
}

type ApplyParamsTemplateRequest struct {
	*tchttp.BaseRequest
	
	// List of instance IDs
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// ID of the parameter template to be applied
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
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
	TaskIds []*int64 `json:"TaskIds,omitempty" name:"TaskIds"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Product *string `json:"Product,omitempty" name:"Product"`

	// ID of the security group to be associated in the format of sg-efil73jd.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// ID(s) of the instance(s) to be associated in the format of ins-lesecurk. You can specify multiple instances.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type AssociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Database engine name, which is `redis` for this API.
	Product *string `json:"Product,omitempty" name:"Product"`

	// ID of the security group to be associated in the format of sg-efil73jd.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// ID(s) of the instance(s) to be associated in the format of ins-lesecurk. You can specify multiple instances.
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

// Predefined struct for user
type AssociateSecurityGroupsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type BackupDownloadInfo struct {
	// Backup file name
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// Backup file size in bytes. If the parameter value is `0`, the backup file size is unknown.
	FileSize *uint64 `json:"FileSize,omitempty" name:"FileSize"`

	// Address (valid for six hours) used to download the backup file over the public network
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// Address (valid for six hours) used to download the backup file over the private network
	InnerDownloadUrl *string `json:"InnerDownloadUrl,omitempty" name:"InnerDownloadUrl"`
}

type BackupLimitVpcItem struct {
	// Region of the VPC of the custom backup file download address.
	Region *string `json:"Region,omitempty" name:"Region"`

	// VPC list of the custom backup file download address.
	VpcList []*string `json:"VpcList,omitempty" name:"VpcList"`
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

// Predefined struct for user
type ChangeInstanceRoleRequestParams struct {
	// Replication group ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance role. Valid values: `rw` (read-write), `r`( read-only).
	InstanceRole *string `json:"InstanceRole,omitempty" name:"InstanceRole"`
}

type ChangeInstanceRoleRequest struct {
	*tchttp.BaseRequest
	
	// Replication group ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance role. Valid values: `rw` (read-write), `r`( read-only).
	InstanceRole *string `json:"InstanceRole,omitempty" name:"InstanceRole"`
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
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Replication group ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type ChangeMasterInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Replication group ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChangeMasterInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChangeMasterInstanceResponseParams struct {
	// Async task ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Replica group ID, which is required for multi-AZ instances.
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
}

type ChangeReplicaToMasterRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Replica group ID, which is required for multi-AZ instances.
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChangeReplicaToMasterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChangeReplicaToMasterResponseParams struct {
	// Async task ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Redis instance password (this parameter is required for password-enabled instances but not for password-free instances)
	Password *string `json:"Password,omitempty" name:"Password"`
}

type ClearInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Redis instance password (this parameter is required for password-enabled instances but not for password-free instances)
	Password *string `json:"Password,omitempty" name:"Password"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ClearInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ClearInstanceResponseParams struct {
	// Task ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type CloseSSLRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type CloseSSLRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseSSLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseSSLResponseParams struct {
	// Task ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Command
	Cmd *string `json:"Cmd,omitempty" name:"Cmd"`

	// Duration
	Took *int64 `json:"Took,omitempty" name:"Took"`
}

// Predefined struct for user
type CreateInstanceAccountRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Sub-account name
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// 1. The password must contain 8–30 characters. A password of 12 or more characters is recommended.
	// 2. It cannot start with a slash (/).
	// 3. It must contain characters in at least two of the following types:
	//     a. Lowercase letters (a–z)
	//     b. Uppercase letters (A–Z)
	//     c. Digits (0–9)
	//     d. ()`~!@#$%^&*-+=_|{}[]:;<>,.?/
	AccountPassword *string `json:"AccountPassword,omitempty" name:"AccountPassword"`

	// Routing policy. Valid values: master (master node); replication (replica node)
	ReadonlyPolicy []*string `json:"ReadonlyPolicy,omitempty" name:"ReadonlyPolicy"`

	// Read/Write policy. Valid values: r (read-only); rw (read/write).
	Privilege *string `json:"Privilege,omitempty" name:"Privilege"`

	// Sub-account description information
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type CreateInstanceAccountRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Sub-account name
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// 1. The password must contain 8–30 characters. A password of 12 or more characters is recommended.
	// 2. It cannot start with a slash (/).
	// 3. It must contain characters in at least two of the following types:
	//     a. Lowercase letters (a–z)
	//     b. Uppercase letters (A–Z)
	//     c. Digits (0–9)
	//     d. ()`~!@#$%^&*-+=_|{}[]:;<>,.?/
	AccountPassword *string `json:"AccountPassword,omitempty" name:"AccountPassword"`

	// Routing policy. Valid values: master (master node); replication (replica node)
	ReadonlyPolicy []*string `json:"ReadonlyPolicy,omitempty" name:"ReadonlyPolicy"`

	// Read/Write policy. Valid values: r (read-only); rw (read/write).
	Privilege *string `json:"Privilege,omitempty" name:"Privilege"`

	// Sub-account description information
	Remark *string `json:"Remark,omitempty" name:"Remark"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstanceAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceAccountResponseParams struct {
	// Task ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Instance type. Valid values: `2` (Redis 2.8 Memory Edition in standard architecture), `3` (CKV 3.2 Memory Edition in standard architecture), `4` (CKV 3.2 Memory Edition in cluster architecture), `6` (Redis 4.0 Memory Edition in standard architecture), `7` (Redis 4.0 Memory Edition in cluster architecture), `8` (Redis 5.0 Memory Edition in standard architecture), `9` (Redis 5.0 Memory Edition in cluster architecture), `15` (Redis 6.0 Memory Edition in standard architecture), `16` (Redis 6.0 Memory Edition in cluster architecture)
	TypeId *uint64 `json:"TypeId,omitempty" name:"TypeId"`

	// Memory capacity in MB, which must be a multiple of 1,024. It is subject to the purchasable specifications returned by the [DescribeProductInfo API](https://intl.cloud.tencent.com/document/api/239/30600?from_cn_redirect=1).
	// If `TypeId` is the standard architecture, `MemSize` indicates the total memory capacity of the instance; if `TypeId` is the cluster architecture, `MemSize` indicates the memory capacity per shard.
	MemSize *uint64 `json:"MemSize,omitempty" name:"MemSize"`

	// Number of instances. The actual quantity purchasable at a time is subject to the specifications returned by the [DescribeProductInfo API](https://intl.cloud.tencent.com/document/api/239/30600?from_cn_redirect=1).
	GoodsNum *uint64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// Length of purchase in months, which is required when creating a monthly subscribed instance. Valid values: [1,2,3,4,5,6,7,8,9,10,11,12,24,36]. For pay-as-you-go instances, set the parameter to `1`.
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// Billing mode. 0: pay-as-you-go
	BillingMode *int64 `json:"BillingMode,omitempty" name:"BillingMode"`

	// ID of the AZ where the instance resides. For more information, see [Regions and AZs](https://intl.cloud.tencent.com/document/product/239/4106?from_cn_redirect=1).
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// Instance password. If the input parameter `NoAuth` is `true` and a VPC is used, the `Password` is optional; otherwise, it is required.
	// If the instance `TypeId` is Redis 2.8, 4.0, or 5.0, the password cannot start with "/" and must contain 8–30 characters in at least two of the following character types: lowercase letters, uppercase letters, digits, and special symbols (()`~!@#$%^&*-+=_|{}[]:;<>,.?/).
	// If the instance `TypeId` is CKV 3.2, the password can contain 8–30 letters and digits.
	Password *string `json:"Password,omitempty" name:"Password"`

	// VPC ID such as vpc-sad23jfdfk. If this parameter is not passed in, the classic network will be selected by default. Use the VPC list querying API to query.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// In the classic network, `subnetId` is invalid. In a VPC subnet, the value is the subnet ID, such as subnet-fdj24n34j2.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Project ID. The value is subject to the `projectId` returned by user account > user account querying APIs > project list.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Auto-renewal flag. 0: default status (manual renewal); 1: auto-renewal enabled; 2: auto-renewal disabled
	AutoRenew *uint64 `json:"AutoRenew,omitempty" name:"AutoRenew"`

	// Array of security group IDs.
	SecurityGroupIdList []*string `json:"SecurityGroupIdList,omitempty" name:"SecurityGroupIdList"`

	// User-defined port. If this parameter is left empty, 6379 will be used by default. Value range: [1024,65535].
	VPort *uint64 `json:"VPort,omitempty" name:"VPort"`

	// Number of shards in an instance. This parameter is required for Cluster Edition instances. Valid values: [3,5,8,12,16,24,32,64,96,128].
	RedisShardNum *int64 `json:"RedisShardNum,omitempty" name:"RedisShardNum"`

	// Number of replicas in the instance. Redis 2.8 Standard Edition and CKV Standard Edition support 1 replica. Standard/Cluster Edition 4.0 and 5.0 support 1–5 replicas.
	RedisReplicasNum *int64 `json:"RedisReplicasNum,omitempty" name:"RedisReplicasNum"`

	// Whether to support read-only replicas. Neither Redis 2.8 Standard Edition nor CKV Standard Edition supports read-only replicas. Read/write separation will be automatically enabled for an instance after it enables read-only replicas. Write requests will be directed to the master node and read requests will be distributed to replica nodes. To enable read-only replicas, we recommend you create two or more replicas.
	ReplicasReadonly *bool `json:"ReplicasReadonly,omitempty" name:"ReplicasReadonly"`

	// Instance name, which can contain up to 60 letters, digits, underscores, or hyphens.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Whether to support the password-free feature. Valid values: true (password-free instance), false (password-enabled instance). Default value: false. Only instances in a VPC support the password-free access.
	NoAuth *bool `json:"NoAuth,omitempty" name:"NoAuth"`

	// Node information of the instance. Currently, information about the node type (master or replica) and node AZ can be passed in. This parameter is not required for single-AZ deployed instances.
	NodeSet []*RedisNodeInfo `json:"NodeSet,omitempty" name:"NodeSet"`

	// Tag bound to the instance to be purchased
	ResourceTags []*ResourceTag `json:"ResourceTags,omitempty" name:"ResourceTags"`

	// Name of the AZ where the instance resides. For more information, see [Regions and AZs](https://intl.cloud.tencent.com/document/product/239/4106?from_cn_redirect=1).
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// ID of the parameter template applied to the created instance. If this parameter is left blank, the default parameter template will be applied.
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// false: send a normal request and create an instance directly after the check is passed (default value); true: send a check request without creating an instance.
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`

	// Valid values: `local` (local disk edition), `cloud` (cloud disk edition), `cdc` (dedicated cluster edition). Default value: `local` (local disk edition)
	ProductVersion *string `json:"ProductVersion,omitempty" name:"ProductVersion"`

	// Dedicated cluster ID, which is required when `ProductVersion` is "cdc".
	RedisClusterId *string `json:"RedisClusterId,omitempty" name:"RedisClusterId"`
}

type CreateInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Instance type. Valid values: `2` (Redis 2.8 Memory Edition in standard architecture), `3` (CKV 3.2 Memory Edition in standard architecture), `4` (CKV 3.2 Memory Edition in cluster architecture), `6` (Redis 4.0 Memory Edition in standard architecture), `7` (Redis 4.0 Memory Edition in cluster architecture), `8` (Redis 5.0 Memory Edition in standard architecture), `9` (Redis 5.0 Memory Edition in cluster architecture), `15` (Redis 6.0 Memory Edition in standard architecture), `16` (Redis 6.0 Memory Edition in cluster architecture)
	TypeId *uint64 `json:"TypeId,omitempty" name:"TypeId"`

	// Memory capacity in MB, which must be a multiple of 1,024. It is subject to the purchasable specifications returned by the [DescribeProductInfo API](https://intl.cloud.tencent.com/document/api/239/30600?from_cn_redirect=1).
	// If `TypeId` is the standard architecture, `MemSize` indicates the total memory capacity of the instance; if `TypeId` is the cluster architecture, `MemSize` indicates the memory capacity per shard.
	MemSize *uint64 `json:"MemSize,omitempty" name:"MemSize"`

	// Number of instances. The actual quantity purchasable at a time is subject to the specifications returned by the [DescribeProductInfo API](https://intl.cloud.tencent.com/document/api/239/30600?from_cn_redirect=1).
	GoodsNum *uint64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// Length of purchase in months, which is required when creating a monthly subscribed instance. Valid values: [1,2,3,4,5,6,7,8,9,10,11,12,24,36]. For pay-as-you-go instances, set the parameter to `1`.
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// Billing mode. 0: pay-as-you-go
	BillingMode *int64 `json:"BillingMode,omitempty" name:"BillingMode"`

	// ID of the AZ where the instance resides. For more information, see [Regions and AZs](https://intl.cloud.tencent.com/document/product/239/4106?from_cn_redirect=1).
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// Instance password. If the input parameter `NoAuth` is `true` and a VPC is used, the `Password` is optional; otherwise, it is required.
	// If the instance `TypeId` is Redis 2.8, 4.0, or 5.0, the password cannot start with "/" and must contain 8–30 characters in at least two of the following character types: lowercase letters, uppercase letters, digits, and special symbols (()`~!@#$%^&*-+=_|{}[]:;<>,.?/).
	// If the instance `TypeId` is CKV 3.2, the password can contain 8–30 letters and digits.
	Password *string `json:"Password,omitempty" name:"Password"`

	// VPC ID such as vpc-sad23jfdfk. If this parameter is not passed in, the classic network will be selected by default. Use the VPC list querying API to query.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// In the classic network, `subnetId` is invalid. In a VPC subnet, the value is the subnet ID, such as subnet-fdj24n34j2.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Project ID. The value is subject to the `projectId` returned by user account > user account querying APIs > project list.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Auto-renewal flag. 0: default status (manual renewal); 1: auto-renewal enabled; 2: auto-renewal disabled
	AutoRenew *uint64 `json:"AutoRenew,omitempty" name:"AutoRenew"`

	// Array of security group IDs.
	SecurityGroupIdList []*string `json:"SecurityGroupIdList,omitempty" name:"SecurityGroupIdList"`

	// User-defined port. If this parameter is left empty, 6379 will be used by default. Value range: [1024,65535].
	VPort *uint64 `json:"VPort,omitempty" name:"VPort"`

	// Number of shards in an instance. This parameter is required for Cluster Edition instances. Valid values: [3,5,8,12,16,24,32,64,96,128].
	RedisShardNum *int64 `json:"RedisShardNum,omitempty" name:"RedisShardNum"`

	// Number of replicas in the instance. Redis 2.8 Standard Edition and CKV Standard Edition support 1 replica. Standard/Cluster Edition 4.0 and 5.0 support 1–5 replicas.
	RedisReplicasNum *int64 `json:"RedisReplicasNum,omitempty" name:"RedisReplicasNum"`

	// Whether to support read-only replicas. Neither Redis 2.8 Standard Edition nor CKV Standard Edition supports read-only replicas. Read/write separation will be automatically enabled for an instance after it enables read-only replicas. Write requests will be directed to the master node and read requests will be distributed to replica nodes. To enable read-only replicas, we recommend you create two or more replicas.
	ReplicasReadonly *bool `json:"ReplicasReadonly,omitempty" name:"ReplicasReadonly"`

	// Instance name, which can contain up to 60 letters, digits, underscores, or hyphens.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Whether to support the password-free feature. Valid values: true (password-free instance), false (password-enabled instance). Default value: false. Only instances in a VPC support the password-free access.
	NoAuth *bool `json:"NoAuth,omitempty" name:"NoAuth"`

	// Node information of the instance. Currently, information about the node type (master or replica) and node AZ can be passed in. This parameter is not required for single-AZ deployed instances.
	NodeSet []*RedisNodeInfo `json:"NodeSet,omitempty" name:"NodeSet"`

	// Tag bound to the instance to be purchased
	ResourceTags []*ResourceTag `json:"ResourceTags,omitempty" name:"ResourceTags"`

	// Name of the AZ where the instance resides. For more information, see [Regions and AZs](https://intl.cloud.tencent.com/document/product/239/4106?from_cn_redirect=1).
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// ID of the parameter template applied to the created instance. If this parameter is left blank, the default parameter template will be applied.
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// false: send a normal request and create an instance directly after the check is passed (default value); true: send a check request without creating an instance.
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`

	// Valid values: `local` (local disk edition), `cloud` (cloud disk edition), `cdc` (dedicated cluster edition). Default value: `local` (local disk edition)
	ProductVersion *string `json:"ProductVersion,omitempty" name:"ProductVersion"`

	// Dedicated cluster ID, which is required when `ProductVersion` is "cdc".
	RedisClusterId *string `json:"RedisClusterId,omitempty" name:"RedisClusterId"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstancesResponseParams struct {
	// Transaction ID
	DealId *string `json:"DealId,omitempty" name:"DealId"`

	// Instance ID
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Parameter template name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Parameter template description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Instance type. Valid values: `1` (Redis 2.8 Memory Edition in cluster architecture), `2` (Redis 2.8 Memory Edition in standard architecture), `3` (CKV 3.2 Memory Edition in standard architecture), `4` (CKV 3.2 Memory Edition in cluster architecture), `5` (Redis 2.8 Memory Edition in standalone architecture), `6` (Redis 4.0 Memory Edition in standard architecture), `7` (Redis 4.0 Memory Edition in cluster architecture), `8` (Redis 5.0 Memory Edition in standard architecture), `9` (Redis 5.0 Memory Edition in cluster architecture). If `TempateId` is specified, this parameter can be left blank; otherwise, it is required.
	ProductType *uint64 `json:"ProductType,omitempty" name:"ProductType"`

	// ID of the source parameter template.
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// Parameter list.
	ParamList []*InstanceParam `json:"ParamList,omitempty" name:"ParamList"`
}

type CreateParamTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Parameter template name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Parameter template description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Instance type. Valid values: `1` (Redis 2.8 Memory Edition in cluster architecture), `2` (Redis 2.8 Memory Edition in standard architecture), `3` (CKV 3.2 Memory Edition in standard architecture), `4` (CKV 3.2 Memory Edition in cluster architecture), `5` (Redis 2.8 Memory Edition in standalone architecture), `6` (Redis 4.0 Memory Edition in standard architecture), `7` (Redis 4.0 Memory Edition in cluster architecture), `8` (Redis 5.0 Memory Edition in standard architecture), `9` (Redis 5.0 Memory Edition in cluster architecture). If `TempateId` is specified, this parameter can be left blank; otherwise, it is required.
	ProductType *uint64 `json:"ProductType,omitempty" name:"ProductType"`

	// ID of the source parameter template.
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// Parameter list.
	ParamList []*InstanceParam `json:"ParamList,omitempty" name:"ParamList"`
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
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type DelayDistribution struct {
	// Delay distribution. The mapping between delay range and `Ladder` value is as follows:
	// [0ms,1ms]: 1;
	// [1ms,5ms]: 5;
	// [5ms,10ms]: 10;
	// [10ms,50ms]: 50;
	// [50ms,200ms]: 200;
	// [200ms,∞]: -1.
	Ladder *int64 `json:"Ladder,omitempty" name:"Ladder"`

	// The number of commands whose delay falls within the current delay range
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// Modification time
	Updatetime *int64 `json:"Updatetime,omitempty" name:"Updatetime"`
}

// Predefined struct for user
type DeleteInstanceAccountRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Sub-account name
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`
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
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Parameter template ID.
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
}

type DeleteParamTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Parameter template ID.
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeAutoBackupConfigRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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
	// Backup type. Automatic backup type: 1 (scheduled rollback)
	AutoBackupType *int64 `json:"AutoBackupType,omitempty" name:"AutoBackupType"`

	// Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday.
	WeekDays []*string `json:"WeekDays,omitempty" name:"WeekDays"`

	// Time period.
	TimePeriod *string `json:"TimePeriod,omitempty" name:"TimePeriod"`

	// Number of days to retain full backup files
	BackupStorageDays *int64 `json:"BackupStorageDays,omitempty" name:"BackupStorageDays"`

	// Number of days to retain Tendis binlog backup files
	BinlogStorageDays *int64 `json:"BinlogStorageDays,omitempty" name:"BinlogStorageDays"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	LimitType *string `json:"LimitType,omitempty" name:"LimitType"`

	// Only `In` can be passed in for this parameter, indicating that backup files can be downloaded in the custom `LimitVpc`.
	VpcComparisonSymbol *string `json:"VpcComparisonSymbol,omitempty" name:"VpcComparisonSymbol"`

	// Whether backups can be downloaded at the custom `LimitIp` address.
	// 
	// - `In`: Download is allowed for the custom IP.
	// - `NotIn`: Download is not allowed for the custom IP.
	IpComparisonSymbol *string `json:"IpComparisonSymbol,omitempty" name:"IpComparisonSymbol"`

	// VPC ID of the custom backup file download address, which will be displayed if `LimitType` is `Customize`.
	LimitVpc []*BackupLimitVpcItem `json:"LimitVpc,omitempty" name:"LimitVpc"`

	// VPC ID of the custom backup file download address, which will be displayed if `LimitType` is `Customize`.
	LimitIp []*string `json:"LimitIp,omitempty" name:"LimitIp"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Backup ID, which can be obtained through the `RedisBackupSet` parameter returned by the [DescribeInstanceBackups](https://intl.cloud.tencent.com/document/product/239/20011?from_cn_redirect=1) API.
	BackupId *string `json:"BackupId,omitempty" name:"BackupId"`

	// Type of the network restriction for downloading backup files. If this parameter is not configured, the user-defined configuration will be used.
	// 
	// - `NoLimit`: Backup files can be downloaded over both public and private networks.
	// - `LimitOnlyIntranet`: Backup files can be downloaded only at private network addresses auto-assigned by Tencent Cloud.
	// - `Customize`: Backup files can be downloaded only in the customized VPC.
	LimitType *string `json:"LimitType,omitempty" name:"LimitType"`

	// Only `In` can be passed in for this parameter, indicating that backup files can be downloaded in the custom `LimitVpc`.
	VpcComparisonSymbol *string `json:"VpcComparisonSymbol,omitempty" name:"VpcComparisonSymbol"`

	// Whether backups can be downloaded at the custom `LimitIp` address.
	// 
	// - `In` (default value): Download is allowed for the custom IP.
	// - `NotIn`: Download is not allowed for the custom IP.
	IpComparisonSymbol *string `json:"IpComparisonSymbol,omitempty" name:"IpComparisonSymbol"`

	// VPC ID of the custom backup file download address, which is required if `LimitType` is `Customize`.
	LimitVpc []*BackupLimitVpcItem `json:"LimitVpc,omitempty" name:"LimitVpc"`

	// VPC IP of the custom backup file download address, which is required if `LimitType` is `Customize`.
	LimitIp []*string `json:"LimitIp,omitempty" name:"LimitIp"`
}

type DescribeBackupUrlRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Backup ID, which can be obtained through the `RedisBackupSet` parameter returned by the [DescribeInstanceBackups](https://intl.cloud.tencent.com/document/product/239/20011?from_cn_redirect=1) API.
	BackupId *string `json:"BackupId,omitempty" name:"BackupId"`

	// Type of the network restriction for downloading backup files. If this parameter is not configured, the user-defined configuration will be used.
	// 
	// - `NoLimit`: Backup files can be downloaded over both public and private networks.
	// - `LimitOnlyIntranet`: Backup files can be downloaded only at private network addresses auto-assigned by Tencent Cloud.
	// - `Customize`: Backup files can be downloaded only in the customized VPC.
	LimitType *string `json:"LimitType,omitempty" name:"LimitType"`

	// Only `In` can be passed in for this parameter, indicating that backup files can be downloaded in the custom `LimitVpc`.
	VpcComparisonSymbol *string `json:"VpcComparisonSymbol,omitempty" name:"VpcComparisonSymbol"`

	// Whether backups can be downloaded at the custom `LimitIp` address.
	// 
	// - `In` (default value): Download is allowed for the custom IP.
	// - `NotIn`: Download is not allowed for the custom IP.
	IpComparisonSymbol *string `json:"IpComparisonSymbol,omitempty" name:"IpComparisonSymbol"`

	// VPC ID of the custom backup file download address, which is required if `LimitType` is `Customize`.
	LimitVpc []*BackupLimitVpcItem `json:"LimitVpc,omitempty" name:"LimitVpc"`

	// VPC IP of the custom backup file download address, which is required if `LimitType` is `Customize`.
	LimitIp []*string `json:"LimitIp,omitempty" name:"LimitIp"`
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
	DownloadUrl []*string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// Private network download address (valid for six hours). This field will be disused soon.
	InnerDownloadUrl []*string `json:"InnerDownloadUrl,omitempty" name:"InnerDownloadUrl"`

	// Filename. This field will be disused soon.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Filenames []*string `json:"Filenames,omitempty" name:"Filenames"`

	// List of backup file information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupInfos []*BackupDownloadInfo `json:"BackupInfos,omitempty" name:"BackupInfos"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeCommonDBInstancesRequestParams struct {
	// List of VPC IDs
	VpcIds []*int64 `json:"VpcIds,omitempty" name:"VpcIds"`

	// List of subnet IDs
	SubnetIds []*int64 `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// List of billing modes. 0: monthly subscription; 1: pay-as-you-go
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// List of instance IDs
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// List of instance names
	InstanceNames []*string `json:"InstanceNames,omitempty" name:"InstanceNames"`

	// List of instance status
	Status []*string `json:"Status,omitempty" name:"Status"`

	// Sorting field
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Sorting order
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// List of instance VIPs
	Vips []*string `json:"Vips,omitempty" name:"Vips"`

	// List of VPC IDs
	UniqVpcIds []*string `json:"UniqVpcIds,omitempty" name:"UniqVpcIds"`

	// List of unique subnet IDs
	UniqSubnetIds []*string `json:"UniqSubnetIds,omitempty" name:"UniqSubnetIds"`

	// Quantity limit. Recommended default value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeCommonDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// List of VPC IDs
	VpcIds []*int64 `json:"VpcIds,omitempty" name:"VpcIds"`

	// List of subnet IDs
	SubnetIds []*int64 `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// List of billing modes. 0: monthly subscription; 1: pay-as-you-go
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// List of instance IDs
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// List of instance names
	InstanceNames []*string `json:"InstanceNames,omitempty" name:"InstanceNames"`

	// List of instance status
	Status []*string `json:"Status,omitempty" name:"Status"`

	// Sorting field
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Sorting order
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// List of instance VIPs
	Vips []*string `json:"Vips,omitempty" name:"Vips"`

	// List of VPC IDs
	UniqVpcIds []*string `json:"UniqVpcIds,omitempty" name:"UniqVpcIds"`

	// List of unique subnet IDs
	UniqSubnetIds []*string `json:"UniqSubnetIds,omitempty" name:"UniqSubnetIds"`

	// Quantity limit. Recommended default value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
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
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Instance information
	InstanceDetails []*RedisCommonInstanceList `json:"InstanceDetails,omitempty" name:"InstanceDetails"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Product *string `json:"Product,omitempty" name:"Product"`

	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeDBSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Database engine name, which is `redis` for this API.
	Product *string `json:"Product,omitempty" name:"Product"`

	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
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

// Predefined struct for user
type DescribeDBSecurityGroupsResponseParams struct {
	// Security group rules
	Groups []*SecurityGroup `json:"Groups,omitempty" name:"Groups"`

	// Private IP for which the security group takes effect
	VIP *string `json:"VIP,omitempty" name:"VIP"`

	// Private port for which the security group takes effect
	VPort *string `json:"VPort,omitempty" name:"VPort"`

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
type DescribeInstanceAccountRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Number of entries per page
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Page offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
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
	// Account details
	// Note: This field may return null, indicating that no valid values can be obtained.
	Accounts []*Account `json:"Accounts,omitempty" name:"Accounts"`

	// Number of accounts
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Number of backups returned per page. Default value: `20`. Maximum value: `100`.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Pagination offset, which is an integral multiple of `Limit`. `offset` = `limit` * (page number - 1).
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// ID of the instance to be operated on, which can be obtained through the `InstanceId` field in the return value of the `DescribeInstance` API.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Start time in the format of yyyy-MM-dd HH:mm:ss, such as 2017-02-08 16:46:34. This parameter is used to query the list of instance backups started during the [beginTime, endTime] range.
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// End time in the format of yyyy-MM-dd HH:mm:ss, such as 2017-02-08 19:09:26. This parameter is used to query the list of instance backups started during the [beginTime, endTime] range.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Backup task status:
	// `1`: The backup is in the process.
	// `2`: The backup is normal.
	// `3`: The backup is being converted to an RDB file.
	// `4`: Conversion to RDB has been completed.
	// `-1`: The backup expired.
	// `-2`: The backup has been deleted.
	Status []*int64 `json:"Status,omitempty" name:"Status"`

	// Instance name, which can be fuzzily searched.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

type DescribeInstanceBackupsRequest struct {
	*tchttp.BaseRequest
	
	// Number of backups returned per page. Default value: `20`. Maximum value: `100`.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Pagination offset, which is an integral multiple of `Limit`. `offset` = `limit` * (page number - 1).
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// ID of the instance to be operated on, which can be obtained through the `InstanceId` field in the return value of the `DescribeInstance` API.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Start time in the format of yyyy-MM-dd HH:mm:ss, such as 2017-02-08 16:46:34. This parameter is used to query the list of instance backups started during the [beginTime, endTime] range.
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// End time in the format of yyyy-MM-dd HH:mm:ss, such as 2017-02-08 19:09:26. This parameter is used to query the list of instance backups started during the [beginTime, endTime] range.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Backup task status:
	// `1`: The backup is in the process.
	// `2`: The backup is normal.
	// `3`: The backup is being converted to an RDB file.
	// `4`: Conversion to RDB has been completed.
	// `-1`: The backup expired.
	// `-2`: The backup has been deleted.
	Status []*int64 `json:"Status,omitempty" name:"Status"`

	// Instance name, which can be fuzzily searched.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
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
	// Total number of backups.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Array of instance backups.
	BackupSet []*RedisBackupSet `json:"BackupSet,omitempty" name:"BackupSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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
	// DTS task ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// DTS task name
	// Note: This field may return null, indicating that no valid values can be obtained.
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// Task status. Valid values: 1 (Creating), 3 (Checking), 4 (CheckPass), 5 (CheckNotPass), 7 (Running), 8 (ReadyComplete), 9 (Success), 10 (Failed), 11 (Stopping), 12 (Completing)
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Status description
	// Note: This field may return null, indicating that no valid values can be obtained.
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

	// Sync latency in bytes
	// Note: This field may return null, indicating that no valid values can be obtained.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Disconnection time
	// Note: This field may return null, indicating that no valid values can be obtained.
	CutDownTime *string `json:"CutDownTime,omitempty" name:"CutDownTime"`

	// Source instance information
	// Note: This field may return null, indicating that no valid values can be obtained.
	SrcInfo *DescribeInstanceDTSInstanceInfo `json:"SrcInfo,omitempty" name:"SrcInfo"`

	// Target instance information
	// Note: This field may return null, indicating that no valid values can be obtained.
	DstInfo *DescribeInstanceDTSInstanceInfo `json:"DstInfo,omitempty" name:"DstInfo"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Region ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// Instance ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Repository ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	SetId *int64 `json:"SetId,omitempty" name:"SetId"`

	// AZ ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// Instance type
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// Instance name
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Instance access address
	// Note: This field may return null, indicating that no valid values can be obtained.
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Status
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

// Predefined struct for user
type DescribeInstanceDealDetailRequestParams struct {
	// Array of order transaction IDs, i.e., the `DealId` output parameter of the [CreateInstances](https://intl.cloud.tencent.com/document/api/239/20026?from_cn_redirect=1) API.
	DealIds []*string `json:"DealIds,omitempty" name:"DealIds"`
}

type DescribeInstanceDealDetailRequest struct {
	*tchttp.BaseRequest
	
	// Array of order transaction IDs, i.e., the `DealId` output parameter of the [CreateInstances](https://intl.cloud.tencent.com/document/api/239/20026?from_cn_redirect=1) API.
	DealIds []*string `json:"DealIds,omitempty" name:"DealIds"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceDealDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceDealDetailResponseParams struct {
	// Order details
	DealDetails []*TradeDealDetail `json:"DealDetails,omitempty" name:"DealDetails"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeInstanceMonitorBigKeyRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Request type. 1: string type; 2: all types
	ReqType *int64 `json:"ReqType,omitempty" name:"ReqType"`

	// Time, such as "20190219"
	Date *string `json:"Date,omitempty" name:"Date"`
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
	Data []*BigKeyInfo `json:"Data,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Time, such as "20190219"
	Date *string `json:"Date,omitempty" name:"Date"`
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
	Data []*DelayDistribution `json:"Data,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Time, such as "20190219"
	Date *string `json:"Date,omitempty" name:"Date"`
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
	Data []*BigKeyTypeInfo `json:"Data,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Time span. 1: real time; 2: past 30 minutes; 3: past 6 hours; 4: past 24 hours
	SpanType *int64 `json:"SpanType,omitempty" name:"SpanType"`
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
	// Hot key details
	Data []*HotKeyInfo `json:"Data,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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
	Data []*SourceInfo `json:"Data,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Time, such as "20190219"
	Date *string `json:"Date,omitempty" name:"Date"`

	// Time span. 1: real time; 2: last 30 minutes; 3: last 6 hours; 4: last 24 hours
	SpanType *int64 `json:"SpanType,omitempty" name:"SpanType"`
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
	// Latency distribution information
	Data []*DelayDistribution `json:"Data,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Time span. 1: real time; 2: last 30 minutes; 3: last 6 hours; 4: last 24 hours
	SpanType *int64 `json:"SpanType,omitempty" name:"SpanType"`
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
	Data []*SourceCommand `json:"Data,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Time span. 1: real time; 2: last 30 minutes; 3: last 6 hours; 4: last 24 hours
	SpanType *int64 `json:"SpanType,omitempty" name:"SpanType"`
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
	Data []*CommandTake `json:"Data,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// List size
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// The offset value
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeInstanceNodeInfoRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// List size
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// The offset value
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
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
	ProxyCount *int64 `json:"ProxyCount,omitempty" name:"ProxyCount"`

	// Proxy node information
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Proxy []*ProxyNodes `json:"Proxy,omitempty" name:"Proxy"`

	// The number of redis nodes
	RedisCount *int64 `json:"RedisCount,omitempty" name:"RedisCount"`

	// Redis node information
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Redis []*RedisNodes `json:"Redis,omitempty" name:"Redis"`

	// The number of tendis nodes
	TendisCount *int64 `json:"TendisCount,omitempty" name:"TendisCount"`

	// Tendis node information
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Tendis []*TendisNodes `json:"Tendis,omitempty" name:"Tendis"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Maximum number of results returned per page
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset, which is an integral multiple of `Limit`.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeInstanceParamRecordsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Maximum number of results returned per page
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset, which is an integral multiple of `Limit`.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
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
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Information of modifications.
	InstanceParamHistory []*InstanceParamHistory `json:"InstanceParamHistory,omitempty" name:"InstanceParamHistory"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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
	// Number of instance parameters
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Instance parameter in Enum type
	InstanceEnumParam []*InstanceEnumParam `json:"InstanceEnumParam,omitempty" name:"InstanceEnumParam"`

	// Instance parameter in Integer type
	InstanceIntegerParam []*InstanceIntegerParam `json:"InstanceIntegerParam,omitempty" name:"InstanceIntegerParam"`

	// Instance parameter in Char type
	InstanceTextParam []*InstanceTextParam `json:"InstanceTextParam,omitempty" name:"InstanceTextParam"`

	// Instance parameter in Multi type
	InstanceMultiParam []*InstanceMultiParam `json:"InstanceMultiParam,omitempty" name:"InstanceMultiParam"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// List of instances
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type DescribeInstanceSecurityGroupRequest struct {
	*tchttp.BaseRequest
	
	// List of instances
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
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
	// Security group information of the instance
	InstanceSecurityGroupsDetail []*InstanceSecurityGroupDetail `json:"InstanceSecurityGroupsDetail,omitempty" name:"InstanceSecurityGroupsDetail"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Whether to filter out the replica node information
	FilterSlave *bool `json:"FilterSlave,omitempty" name:"FilterSlave"`
}

type DescribeInstanceShardsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Whether to filter out the replica node information
	FilterSlave *bool `json:"FilterSlave,omitempty" name:"FilterSlave"`
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
	// Information list of instance shards
	InstanceShards []*InstanceClusterShard `json:"InstanceShards,omitempty" name:"InstanceShards"`

	// Total number of instance shard nodes
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeInstanceZoneInfoRequestParams struct {
	// Instance ID, such as crs-6ubhgouj
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeInstanceZoneInfoRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, such as crs-6ubhgouj
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of instance node groups
	ReplicaGroups []*ReplicaGroup `json:"ReplicaGroups,omitempty" name:"ReplicaGroups"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Pagination offset, which is an integral multiple of `Limit`.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Instance ID, such as crs-6ubhgouj.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance sorting criteria. The enumerated values are as listed below: <ul><li>projectId: Project ID. </li><li>createtime: Instance creation time. </li><li>instancename: Instance name. </li><li>type: Instance type. </li><li>curDeadline: Instance expiration time. </li></ul>
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Instance sorting order. <ul><li>`1`: Descending. </li><li>`0`: Ascending. Default value: `1`.</li></ul>
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// Array of VPC IDs such as 47525. If this parameter is not passed in or the array is empty, the classic network will be selected by default. This parameter is retained and can be ignored. It is set based on `UniqVpcIds` parameter format.
	VpcIds []*string `json:"VpcIds,omitempty" name:"VpcIds"`

	// Array of VPC subnet IDs such as 56854. This parameter is retained and can be ignored. It is set based on `UniqSubnetIds` parameter format.
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// Keywords for fuzzy query. which can be used to fuzzy query an instance by its ID or name.
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// Array of project IDs
	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Array of VPC IDs such as vpc-sad23jfdfk. If this parameter is not passed in or or the array is empty, the classic network will be selected by default.
	UniqVpcIds []*string `json:"UniqVpcIds,omitempty" name:"UniqVpcIds"`

	// Array of VPC subnet IDs such as subnet-fdj24n34j2
	UniqSubnetIds []*string `json:"UniqSubnetIds,omitempty" name:"UniqSubnetIds"`

	// Array of region IDs (disused). The corresponding region can be queried through the common parameter `Region`.
	RegionIds []*int64 `json:"RegionIds,omitempty" name:"RegionIds"`

	// Instance status. <ul><li>`0`: Uninitialized. </li><li>`1`: In the process. </li><li>`2`: Running. </li><li>`-2`: Isolated. </li><li>`-3`: To be deleted. </li></ul>
	Status []*int64 `json:"Status,omitempty" name:"Status"`

	// Instance architecture. <ul><li>`1`: Standalone edition. </li><li>`2`: Master-replica edition. </li><li>`3`: Cluster edition. </li></ul>
	TypeVersion *int64 `json:"TypeVersion,omitempty" name:"TypeVersion"`

	// Storage engine information. Valid values: `Redis-2.8`, `Redis-4.0`, `Redis-5.0`, `Redis-6.0` or `CKV`.
	EngineName *string `json:"EngineName,omitempty" name:"EngineName"`

	// Renewal mode. <ul><li>`0`: Manual renewal (default). </li><li>`1`: Auto-renewal. </li><li>`2`: No auto-renewal (set by user)</ul>
	AutoRenew []*int64 `json:"AutoRenew,omitempty" name:"AutoRenew"`

	// Billing mode. Only pay-as-you-go billing is supported.
	BillingMode *string `json:"BillingMode,omitempty" name:"BillingMode"`

	// Instance type. <ul><li>`1`: Legacy Redis cluster edition. </li><li>`2`: Redis 2.8 master-replica edition. </li><li>`3`: CKV master-replica edition. </li><li>`4`: CKV cluster edition. </li><li>`5`: Redis 2.8 standalone edition. </li><li>`6`: Redis 4.0 master-replica edition. </li><li>`7`: Redis 4.0 cluster edition. </li><li>8: Redis 5.0 master-replica edition. </li><li>`9`: Redis 5.0 cluster edition. </li></ul>
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// Array of the search keywords, which can query the instance by its ID, name, IP address.
	SearchKeys []*string `json:"SearchKeys,omitempty" name:"SearchKeys"`

	// Internal parameter, which can be ignored.
	TypeList []*int64 `json:"TypeList,omitempty" name:"TypeList"`

	// Internal parameter, which can be ignored.
	MonitorVersion *string `json:"MonitorVersion,omitempty" name:"MonitorVersion"`

	// Resources filter by tag key and value. If this parameter is not specified or is an empty array, resources will not be filtered.
	InstanceTags []*InstanceTagInfo `json:"InstanceTags,omitempty" name:"InstanceTags"`

	// Resources filter by tag key. If this parameter is not specified or is an empty array, resources will not be filtered.
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`

	// Instance product version. If this parameter is not passed in or the array is empty, the instances will not be filtered based this parameter by default. <ul><li>`local`: local disk edition. </li><li>`cloud`: Cloud disk edition. </li><li>`cdc`: Dedicated cluster edition. </li></ul>
	ProductVersions []*string `json:"ProductVersions,omitempty" name:"ProductVersions"`

	// Batch query of the specified instances ID. The number of results returned is based on `Limit`.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// AZ deployment mode. <ul><li>`singleaz`: Single-AZ. </li><li>`1`: Multi-AZ. </li></ul>
	AzMode *string `json:"AzMode,omitempty" name:"AzMode"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Number of instances returned per page. Default value: `20`. Maximum value: `1000`.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Pagination offset, which is an integral multiple of `Limit`.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Instance ID, such as crs-6ubhgouj.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance sorting criteria. The enumerated values are as listed below: <ul><li>projectId: Project ID. </li><li>createtime: Instance creation time. </li><li>instancename: Instance name. </li><li>type: Instance type. </li><li>curDeadline: Instance expiration time. </li></ul>
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Instance sorting order. <ul><li>`1`: Descending. </li><li>`0`: Ascending. Default value: `1`.</li></ul>
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// Array of VPC IDs such as 47525. If this parameter is not passed in or the array is empty, the classic network will be selected by default. This parameter is retained and can be ignored. It is set based on `UniqVpcIds` parameter format.
	VpcIds []*string `json:"VpcIds,omitempty" name:"VpcIds"`

	// Array of VPC subnet IDs such as 56854. This parameter is retained and can be ignored. It is set based on `UniqSubnetIds` parameter format.
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// Keywords for fuzzy query. which can be used to fuzzy query an instance by its ID or name.
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// Array of project IDs
	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Array of VPC IDs such as vpc-sad23jfdfk. If this parameter is not passed in or or the array is empty, the classic network will be selected by default.
	UniqVpcIds []*string `json:"UniqVpcIds,omitempty" name:"UniqVpcIds"`

	// Array of VPC subnet IDs such as subnet-fdj24n34j2
	UniqSubnetIds []*string `json:"UniqSubnetIds,omitempty" name:"UniqSubnetIds"`

	// Array of region IDs (disused). The corresponding region can be queried through the common parameter `Region`.
	RegionIds []*int64 `json:"RegionIds,omitempty" name:"RegionIds"`

	// Instance status. <ul><li>`0`: Uninitialized. </li><li>`1`: In the process. </li><li>`2`: Running. </li><li>`-2`: Isolated. </li><li>`-3`: To be deleted. </li></ul>
	Status []*int64 `json:"Status,omitempty" name:"Status"`

	// Instance architecture. <ul><li>`1`: Standalone edition. </li><li>`2`: Master-replica edition. </li><li>`3`: Cluster edition. </li></ul>
	TypeVersion *int64 `json:"TypeVersion,omitempty" name:"TypeVersion"`

	// Storage engine information. Valid values: `Redis-2.8`, `Redis-4.0`, `Redis-5.0`, `Redis-6.0` or `CKV`.
	EngineName *string `json:"EngineName,omitempty" name:"EngineName"`

	// Renewal mode. <ul><li>`0`: Manual renewal (default). </li><li>`1`: Auto-renewal. </li><li>`2`: No auto-renewal (set by user)</ul>
	AutoRenew []*int64 `json:"AutoRenew,omitempty" name:"AutoRenew"`

	// Billing mode. Only pay-as-you-go billing is supported.
	BillingMode *string `json:"BillingMode,omitempty" name:"BillingMode"`

	// Instance type. <ul><li>`1`: Legacy Redis cluster edition. </li><li>`2`: Redis 2.8 master-replica edition. </li><li>`3`: CKV master-replica edition. </li><li>`4`: CKV cluster edition. </li><li>`5`: Redis 2.8 standalone edition. </li><li>`6`: Redis 4.0 master-replica edition. </li><li>`7`: Redis 4.0 cluster edition. </li><li>8: Redis 5.0 master-replica edition. </li><li>`9`: Redis 5.0 cluster edition. </li></ul>
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// Array of the search keywords, which can query the instance by its ID, name, IP address.
	SearchKeys []*string `json:"SearchKeys,omitempty" name:"SearchKeys"`

	// Internal parameter, which can be ignored.
	TypeList []*int64 `json:"TypeList,omitempty" name:"TypeList"`

	// Internal parameter, which can be ignored.
	MonitorVersion *string `json:"MonitorVersion,omitempty" name:"MonitorVersion"`

	// Resources filter by tag key and value. If this parameter is not specified or is an empty array, resources will not be filtered.
	InstanceTags []*InstanceTagInfo `json:"InstanceTags,omitempty" name:"InstanceTags"`

	// Resources filter by tag key. If this parameter is not specified or is an empty array, resources will not be filtered.
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`

	// Instance product version. If this parameter is not passed in or the array is empty, the instances will not be filtered based this parameter by default. <ul><li>`local`: local disk edition. </li><li>`cloud`: Cloud disk edition. </li><li>`cdc`: Dedicated cluster edition. </li></ul>
	ProductVersions []*string `json:"ProductVersions,omitempty" name:"ProductVersions"`

	// Batch query of the specified instances ID. The number of results returned is based on `Limit`.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// AZ deployment mode. <ul><li>`singleaz`: Single-AZ. </li><li>`1`: Multi-AZ. </li></ul>
	AzMode *string `json:"AzMode,omitempty" name:"AzMode"`
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
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of instance details
	InstanceSet []*InstanceSet `json:"InstanceSet,omitempty" name:"InstanceSet"`

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
type DescribeMaintenanceWindowRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeMaintenanceWindowRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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
	// Start time of the maintenance window, such as 17:00.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of the maintenance window, such as 19:00.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Parameter template ID.
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
}

type DescribeParamTemplateInfoRequest struct {
	*tchttp.BaseRequest
	
	// Parameter template ID.
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
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
	// Number of instance parameters
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Parameter template ID.
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// Parameter template name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Instance type. Valid values: `1` (Redis 2.8 Memory Edition in cluster architecture), `2` (Redis 2.8 Memory Edition in standard architecture), `3` (CKV 3.2 Memory Edition in standard architecture), `4` (CKV 3.2 Memory Edition in cluster architecture), `5` (Redis 2.8 Memory Edition in standalone architecture), `6` (Redis 4.0 Memory Edition in standard architecture), `7` (Redis 4.0 Memory Edition in cluster architecture), `8` (Redis 5.0 Memory Edition in standard architecture), `9` (Redis 5.0 Memory Edition in cluster architecture)
	ProductType *uint64 `json:"ProductType,omitempty" name:"ProductType"`

	// Parameter template description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Parameter details
	Items []*ParameterDetail `json:"Items,omitempty" name:"Items"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Array of instance types. Valid values: `1` (Redis 2.8 Memory Edition in cluster architecture), `2` (Redis 2.8 Memory Edition in standard architecture), `3` (CKV 3.2 Memory Edition in standard architecture), `4` (CKV 3.2 Memory Edition in cluster architecture), `5` (Redis 2.8 Memory Edition in standalone architecture), `6` (Redis 4.0 Memory Edition in standard architecture), `7` (Redis 4.0 Memory Edition in cluster architecture), `8` (Redis 5.0 Memory Edition in standard architecture), `9` (Redis 5.0 Memory Edition in cluster architecture).
	ProductTypes []*int64 `json:"ProductTypes,omitempty" name:"ProductTypes"`

	// Array of template names.
	TemplateNames []*string `json:"TemplateNames,omitempty" name:"TemplateNames"`

	// Array of template IDs.
	TemplateIds []*string `json:"TemplateIds,omitempty" name:"TemplateIds"`
}

type DescribeParamTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// Array of instance types. Valid values: `1` (Redis 2.8 Memory Edition in cluster architecture), `2` (Redis 2.8 Memory Edition in standard architecture), `3` (CKV 3.2 Memory Edition in standard architecture), `4` (CKV 3.2 Memory Edition in cluster architecture), `5` (Redis 2.8 Memory Edition in standalone architecture), `6` (Redis 4.0 Memory Edition in standard architecture), `7` (Redis 4.0 Memory Edition in cluster architecture), `8` (Redis 5.0 Memory Edition in standard architecture), `9` (Redis 5.0 Memory Edition in cluster architecture).
	ProductTypes []*int64 `json:"ProductTypes,omitempty" name:"ProductTypes"`

	// Array of template names.
	TemplateNames []*string `json:"TemplateNames,omitempty" name:"TemplateNames"`

	// Array of template IDs.
	TemplateIds []*string `json:"TemplateIds,omitempty" name:"TemplateIds"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeParamTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeParamTemplatesResponseParams struct {
	// Number of parameter templates of the user.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Parameter template details.
	Items []*ParamTemplateInfo `json:"Items,omitempty" name:"Items"`

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
	// Sale information of a region.
	RegionSet []*RegionConf `json:"RegionSet,omitempty" name:"RegionSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 0: default project; -1: all projects; >0: specified project
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Security group ID
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
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
	SecurityGroupDetails []*SecurityGroupDetail `json:"SecurityGroupDetails,omitempty" name:"SecurityGroupDetails"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Product *string `json:"Product,omitempty" name:"Product"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Offset, which is an integral multiple of `Limit`.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// The number of security groups to be pulled. Default value: `20`.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Search criteria. You can enter a security group ID or name.
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
}

type DescribeProjectSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Database engine name, which is `redis` for this API.
	Product *string `json:"Product,omitempty" name:"Product"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Offset, which is an integral multiple of `Limit`.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// The number of security groups to be pulled. Default value: `20`.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Search criteria. You can enter a security group ID or name.
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
	Groups []*SecurityGroup `json:"Groups,omitempty" name:"Groups"`

	// Total number of eligible security groups.
	Total *uint64 `json:"Total,omitempty" name:"Total"`

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
type DescribeProxySlowLogRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Start time
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Slow query threshold in milliseconds
	MinQueryTime *int64 `json:"MinQueryTime,omitempty" name:"MinQueryTime"`

	// Maximum number of results returned per page
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset, which is an integral multiple of `Limit`.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeProxySlowLogRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Start time
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Slow query threshold in milliseconds
	MinQueryTime *int64 `json:"MinQueryTime,omitempty" name:"MinQueryTime"`

	// Maximum number of results returned per page
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset, which is an integral multiple of `Limit`.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
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
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Slow query details
	InstanceProxySlowLogDetail []*InstanceProxySlowlogDetail `json:"InstanceProxySlowLogDetail,omitempty" name:"InstanceProxySlowLogDetail"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeReplicationGroupRequestParams struct {
	// Number of instances returned per page. Default value: `20`.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Pagination offset, which is an integral multiple of `Limit`. `offset` = `limit` * (page number - 1).
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Replication group ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Keyword for fuzzy search, which can be an instance name or instance ID.
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
}

type DescribeReplicationGroupRequest struct {
	*tchttp.BaseRequest
	
	// Number of instances returned per page. Default value: `20`.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Pagination offset, which is an integral multiple of `Limit`. `offset` = `limit` * (page number - 1).
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Replication group ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Keyword for fuzzy search, which can be an instance name or instance ID.
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
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
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Replication group information
	Groups []*Groups `json:"Groups,omitempty" name:"Groups"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeSSLStatusRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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
	// Certificate download address
	CertDownloadUrl *string `json:"CertDownloadUrl,omitempty" name:"CertDownloadUrl"`

	// Expiration time of the certificate download address
	UrlExpiredTime *string `json:"UrlExpiredTime,omitempty" name:"UrlExpiredTime"`

	// SSL configuration status of an instance. Valid values: `true` (enable), `false` (disable).
	SSLConfig *bool `json:"SSLConfig,omitempty" name:"SSLConfig"`

	// Whether the instance supports SSL. Valid values: `true` (Yes. When minor version is upgraded.), `false` (No).
	FeatureSupport *bool `json:"FeatureSupport,omitempty" name:"FeatureSupport"`

	// SSL configuration status. Valid values: `1`(Configuring), `2` (Configured).
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeSlowLogRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// The start time
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// The end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// The average execution time threshold of slow query in ms.
	MinQueryTime *int64 `json:"MinQueryTime,omitempty" name:"MinQueryTime"`

	// Number of slow queries displayed per page. Default value: `20`.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Slow query offset, which is an integral multiple of `Limit`.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Node role. <ul><li>`Master`: Master node</li><li>`Slave`: Replica node</li></ul>
	Role *string `json:"Role,omitempty" name:"Role"`
}

type DescribeSlowLogRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// The start time
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// The end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// The average execution time threshold of slow query in ms.
	MinQueryTime *int64 `json:"MinQueryTime,omitempty" name:"MinQueryTime"`

	// Number of slow queries displayed per page. Default value: `20`.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Slow query offset, which is an integral multiple of `Limit`.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Node role. <ul><li>`Master`: Master node</li><li>`Slave`: Replica node</li></ul>
	Role *string `json:"Role,omitempty" name:"Role"`
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
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Slow query details
	InstanceSlowlogDetail []*InstanceSlowlogDetail `json:"InstanceSlowlogDetail,omitempty" name:"InstanceSlowlogDetail"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Task ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
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
	// Task status. preparing: to be run; running: running; succeed: succeeded; failed: failed; error: running error.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Task start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Task type
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Task message, which is displayed in case of an error. It will be blank if the status is running or succeeded.
	TaskMessage *string `json:"TaskMessage,omitempty" name:"TaskMessage"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Maximum number of results returned per page. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset, which is an integral multiple of `Limit` (rounded down automatically).
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Project ID
	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds"`

	// Task type
	TaskTypes []*string `json:"TaskTypes,omitempty" name:"TaskTypes"`

	// Start time
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Task status
	TaskStatus []*int64 `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// Task status
	Result []*int64 `json:"Result,omitempty" name:"Result"`

	// The field `OperatorUin` has been disused and replaced by `OperateUin`.
	OperatorUin []*int64 `json:"OperatorUin,omitempty" name:"OperatorUin"`

	// Operator Uin
	OperateUin []*string `json:"OperateUin,omitempty" name:"OperateUin"`
}

type DescribeTaskListRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Maximum number of results returned per page. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset, which is an integral multiple of `Limit` (rounded down automatically).
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Project ID
	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds"`

	// Task type
	TaskTypes []*string `json:"TaskTypes,omitempty" name:"TaskTypes"`

	// Start time
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Task status
	TaskStatus []*int64 `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// Task status
	Result []*int64 `json:"Result,omitempty" name:"Result"`

	// The field `OperatorUin` has been disused and replaced by `OperateUin`.
	OperatorUin []*int64 `json:"OperatorUin,omitempty" name:"OperatorUin"`

	// Operator Uin
	OperateUin []*string `json:"OperateUin,omitempty" name:"OperateUin"`
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
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Task details
	Tasks []*TaskInfoDetail `json:"Tasks,omitempty" name:"Tasks"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Instance ID in the format of crs-ngvou0i1
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Start time in the format of 2019-09-08 12:12:41
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// End time in the format of 2019-09-09 12:12:41
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Slow query threshold in ms
	MinQueryTime *int64 `json:"MinQueryTime,omitempty" name:"MinQueryTime"`

	// Maximum number of results returned per page. Default value: 20.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset, which is an integral multiple of `Limit`.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeTendisSlowLogRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of crs-ngvou0i1
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Start time in the format of 2019-09-08 12:12:41
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// End time in the format of 2019-09-09 12:12:41
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Slow query threshold in ms
	MinQueryTime *int64 `json:"MinQueryTime,omitempty" name:"MinQueryTime"`

	// Maximum number of results returned per page. Default value: 20.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset, which is an integral multiple of `Limit`.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
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
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Slow query details
	TendisSlowLogDetail []*TendisSlowLogDetail `json:"TendisSlowLogDetail,omitempty" name:"TendisSlowLogDetail"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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
	DealId *string `json:"DealId,omitempty" name:"DealId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DisableReplicaReadonlyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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
	// Task ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Product *string `json:"Product,omitempty" name:"Product"`

	// Security group ID
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// List of instance IDs, which is an array of one or more instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type DisassociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Database engine name, which is `redis` for this API.
	Product *string `json:"Product,omitempty" name:"Product"`

	// Security group ID
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// List of instance IDs, which is an array of one or more instance IDs.
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

// Predefined struct for user
type DisassociateSecurityGroupsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Account routing policy. If `master` or `replication` is entered, it means to route to the master or replica node; if this parameter is left empty, it means to write into the master node and read from the replica node by default.
	ReadonlyPolicy []*string `json:"ReadonlyPolicy,omitempty" name:"ReadonlyPolicy"`
}

type EnableReplicaReadonlyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Account routing policy. If `master` or `replication` is entered, it means to route to the master or replica node; if this parameter is left empty, it means to write into the master node and read from the replica node by default.
	ReadonlyPolicy []*string `json:"ReadonlyPolicy,omitempty" name:"ReadonlyPolicy"`
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
	// Valid values: `ERROR`, `OK`. This field has been disused.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Task ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type Groups struct {
	// User App ID
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// Region ID. 1: Guangzhou; 4: Shanghai; 5: Hong Kong (China); 6: Toronto; 7: Shanghai Finance; 8: Beijing; 9: Singapore; 11: Shenzhen Finance; 15: West US (Silicon Valley); 16: Chengdu; 17: Germany; 18: South Korea; 19: Chongqing; 21: India; 22: East US (Virginia); 23: Thailand; 24: Russia; 25: Japan
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// Replication group info
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Replication group name
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// Replication group status. Valid values: 37 (Associating replication group), 38 (Reconnecting to replication group), 51 (Disassociating the replication group), 52 (Switching primary instance in the replication group), 53 (Modifying roles)
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Number of replication groups
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// Instances in replication group
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Instances []*Instances `json:"Instances,omitempty" name:"Instances"`

	// Remarks
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Remark *string `json:"Remark,omitempty" name:"Remark"`
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

	// Network protocol, such as UDP and TCP.
	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`

	// Port.
	PortRange *string `json:"PortRange,omitempty" name:"PortRange"`

	// All the protocols and ports that the service group ID represents.
	ServiceModule *string `json:"ServiceModule,omitempty" name:"ServiceModule"`

	// All the addresses that the security group ID represents.
	Id *string `json:"Id,omitempty" name:"Id"`
}

// Predefined struct for user
type InquiryPriceCreateInstanceRequestParams struct {
	// Instance type. Valid values: `2` (Redis 2.8 memory edition in standard architecture), `3` (CKV 3.2 memory edition in standard architecture), `4` (CKV 3.2 memory edition in cluster architecture), `6` (Redis 4.0 memory edition in standard architecture), `7` (Redis 4.0 memory edition in cluster architecture), `8` (Redis 5.0 memory edition in standard architecture), `9` (Redis 5.0 memory edition in cluster architecture).
	TypeId *uint64 `json:"TypeId,omitempty" name:"TypeId"`

	// Memory capacity in MB, which must be a multiple of 1,024. It is subject to the purchasable specifications returned by the [DescribeProductInfo API](https://intl.cloud.tencent.com/document/api/239/30600?from_cn_redirect=1).
	// If `TypeId` indicates the standard architecture, `MemSize` indicates the total memory capacity of an instance; if `TypeId` indicates the cluster architecture, `MemSize` indicates the memory capacity per shard.
	MemSize *uint64 `json:"MemSize,omitempty" name:"MemSize"`

	// Number of instances. The actual quantity purchasable at a time is subject to the specifications returned by the [DescribeProductInfo API](https://intl.cloud.tencent.com/document/api/239/30600?from_cn_redirect=1).
	GoodsNum *uint64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// Length of purchase in months, which is required when creating a monthly-subscribed instance. Value range: [1,2,3,4,5,6,7,8,9,10,11,12,24,36]. For pay-as-you-go instances, set the parameter to `1`.
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// Billing mode. Valid values: `0` (pay-as-you-go), `1` (monthly subscription).
	BillingMode *int64 `json:"BillingMode,omitempty" name:"BillingMode"`

	// ID of the AZ where the instance resides. For more information, see [Regions and AZs](https://intl.cloud.tencent.com/document/product/239/4106?from_cn_redirect=1).
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// Instance shard quantity. This field is not required by Redis 2.8 standard architecture, CKV standard architecture, Redis 2.8 standalone edition, and Redis 4.0 standard architecture.
	RedisShardNum *int64 `json:"RedisShardNum,omitempty" name:"RedisShardNum"`

	// Instance replica quantity. This field is not required by Redis 2.8 standard architecture, CKV standard architecture, and Redis 2.8 standalone edition.
	RedisReplicasNum *int64 `json:"RedisReplicasNum,omitempty" name:"RedisReplicasNum"`

	// Whether to support read-only replicas. This field is not required by Redis 2.8 standard architecture, CKV standard architecture, and Redis 2.8 standalone edition.
	ReplicasReadonly *bool `json:"ReplicasReadonly,omitempty" name:"ReplicasReadonly"`

	// Name of the AZ where the instance resides. For more information, see [Regions and AZs](https://intl.cloud.tencent.com/document/product/239/4106?from_cn_redirect=1).
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// Valid values: `local` (local disk edition), `cloud` (cloud disk edition), `cdc` (dedicated cluster edition). Default value: `local` (local disk edition)
	ProductVersion *string `json:"ProductVersion,omitempty" name:"ProductVersion"`
}

type InquiryPriceCreateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance type. Valid values: `2` (Redis 2.8 memory edition in standard architecture), `3` (CKV 3.2 memory edition in standard architecture), `4` (CKV 3.2 memory edition in cluster architecture), `6` (Redis 4.0 memory edition in standard architecture), `7` (Redis 4.0 memory edition in cluster architecture), `8` (Redis 5.0 memory edition in standard architecture), `9` (Redis 5.0 memory edition in cluster architecture).
	TypeId *uint64 `json:"TypeId,omitempty" name:"TypeId"`

	// Memory capacity in MB, which must be a multiple of 1,024. It is subject to the purchasable specifications returned by the [DescribeProductInfo API](https://intl.cloud.tencent.com/document/api/239/30600?from_cn_redirect=1).
	// If `TypeId` indicates the standard architecture, `MemSize` indicates the total memory capacity of an instance; if `TypeId` indicates the cluster architecture, `MemSize` indicates the memory capacity per shard.
	MemSize *uint64 `json:"MemSize,omitempty" name:"MemSize"`

	// Number of instances. The actual quantity purchasable at a time is subject to the specifications returned by the [DescribeProductInfo API](https://intl.cloud.tencent.com/document/api/239/30600?from_cn_redirect=1).
	GoodsNum *uint64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// Length of purchase in months, which is required when creating a monthly-subscribed instance. Value range: [1,2,3,4,5,6,7,8,9,10,11,12,24,36]. For pay-as-you-go instances, set the parameter to `1`.
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// Billing mode. Valid values: `0` (pay-as-you-go), `1` (monthly subscription).
	BillingMode *int64 `json:"BillingMode,omitempty" name:"BillingMode"`

	// ID of the AZ where the instance resides. For more information, see [Regions and AZs](https://intl.cloud.tencent.com/document/product/239/4106?from_cn_redirect=1).
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// Instance shard quantity. This field is not required by Redis 2.8 standard architecture, CKV standard architecture, Redis 2.8 standalone edition, and Redis 4.0 standard architecture.
	RedisShardNum *int64 `json:"RedisShardNum,omitempty" name:"RedisShardNum"`

	// Instance replica quantity. This field is not required by Redis 2.8 standard architecture, CKV standard architecture, and Redis 2.8 standalone edition.
	RedisReplicasNum *int64 `json:"RedisReplicasNum,omitempty" name:"RedisReplicasNum"`

	// Whether to support read-only replicas. This field is not required by Redis 2.8 standard architecture, CKV standard architecture, and Redis 2.8 standalone edition.
	ReplicasReadonly *bool `json:"ReplicasReadonly,omitempty" name:"ReplicasReadonly"`

	// Name of the AZ where the instance resides. For more information, see [Regions and AZs](https://intl.cloud.tencent.com/document/product/239/4106?from_cn_redirect=1).
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// Valid values: `local` (local disk edition), `cloud` (cloud disk edition), `cdc` (dedicated cluster edition). Default value: `local` (local disk edition)
	ProductVersion *string `json:"ProductVersion,omitempty" name:"ProductVersion"`
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
	// Price. Unit: USD (accurate down to the cent)
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Price *float64 `json:"Price,omitempty" name:"Price"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Shard size in MB.
	MemSize *uint64 `json:"MemSize,omitempty" name:"MemSize"`

	// Number of shards. This parameter can be left blank for Redis 2.8 in standard architecture, CKV in standard architecture, and Redis 2.8 in standalone architecture.
	RedisShardNum *uint64 `json:"RedisShardNum,omitempty" name:"RedisShardNum"`

	// Number of replicas. This parameter can be left blank for Redis 2.8 in standard architecture, CKV in standard architecture, and Redis 2.8 in standalone architecture.
	RedisReplicasNum *uint64 `json:"RedisReplicasNum,omitempty" name:"RedisReplicasNum"`
}

type InquiryPriceUpgradeInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Shard size in MB.
	MemSize *uint64 `json:"MemSize,omitempty" name:"MemSize"`

	// Number of shards. This parameter can be left blank for Redis 2.8 in standard architecture, CKV in standard architecture, and Redis 2.8 in standalone architecture.
	RedisShardNum *uint64 `json:"RedisShardNum,omitempty" name:"RedisShardNum"`

	// Number of replicas. This parameter can be left blank for Redis 2.8 in standard architecture, CKV in standard architecture, and Redis 2.8 in standalone architecture.
	RedisReplicasNum *uint64 `json:"RedisReplicasNum,omitempty" name:"RedisReplicasNum"`
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
	// Price. Unit: USD
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Price *float64 `json:"Price,omitempty" name:"Price"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Node name
	Name *string `json:"Name,omitempty" name:"Name"`

	// ID of the runtime node of the instance
	RunId *string `json:"RunId,omitempty" name:"RunId"`

	// Cluster role. 0: master; 1: replica
	Role *int64 `json:"Role,omitempty" name:"Role"`

	// Node status. 0: readwrite; 1: read; 2: backup
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Service status. 0: down; 1: on
	Connected *int64 `json:"Connected,omitempty" name:"Connected"`

	// Node creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Node elimination time
	DownTime *string `json:"DownTime,omitempty" name:"DownTime"`

	// Distribution of node slots
	Slots *string `json:"Slots,omitempty" name:"Slots"`

	// Distribution of node keys
	Keys *int64 `json:"Keys,omitempty" name:"Keys"`

	// Node QPS
	Qps *int64 `json:"Qps,omitempty" name:"Qps"`

	// Node QPS slope
	QpsSlope *float64 `json:"QpsSlope,omitempty" name:"QpsSlope"`

	// Node storage
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// Node storage slope
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

	// ID of the runtime node of the instance
	Runid *string `json:"Runid,omitempty" name:"Runid"`

	// Service status. 0: down; 1: on
	Connected *int64 `json:"Connected,omitempty" name:"Connected"`
}

type InstanceEnumParam struct {
	// Parameter name
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`

	// Parameter type: Enum
	ValueType *string `json:"ValueType,omitempty" name:"ValueType"`

	// Whether restart is required after a modification is made. Valid values: true, false
	NeedRestart *string `json:"NeedRestart,omitempty" name:"NeedRestart"`

	// Default value of the parameter
	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`

	// Current value
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`

	// Parameter description
	Tips *string `json:"Tips,omitempty" name:"Tips"`

	// Valid values of the parameter
	EnumValue []*string `json:"EnumValue,omitempty" name:"EnumValue"`

	// Parameter status. 1: modifying; 2: modified
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type InstanceIntegerParam struct {
	// Parameter name
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`

	// Parameter type: Integer
	ValueType *string `json:"ValueType,omitempty" name:"ValueType"`

	// Whether restart is required after a modification is made. Valid values: true, false
	NeedRestart *string `json:"NeedRestart,omitempty" name:"NeedRestart"`

	// Default value of the parameter
	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`

	// Current value
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`

	// Parameter description
	Tips *string `json:"Tips,omitempty" name:"Tips"`

	// Minimum value of the parameter
	Min *string `json:"Min,omitempty" name:"Min"`

	// Maximum value of the parameter
	Max *string `json:"Max,omitempty" name:"Max"`

	// Parameter status. 1: modifying; 2: modified
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Parameter unit
	// Note: This field may return null, indicating that no valid values can be obtained.
	Unit *string `json:"Unit,omitempty" name:"Unit"`
}

type InstanceMultiParam struct {
	// Parameter name
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`

	// Parameter type: Multi
	ValueType *string `json:"ValueType,omitempty" name:"ValueType"`

	// Whether restart is required after a modification is made. Valid values: true, false
	NeedRestart *string `json:"NeedRestart,omitempty" name:"NeedRestart"`

	// Default value of the parameter
	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`

	// Current value
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`

	// Parameter description
	Tips *string `json:"Tips,omitempty" name:"Tips"`

	// Parameter description
	EnumValue []*string `json:"EnumValue,omitempty" name:"EnumValue"`

	// Parameter status. 1: modifying; 2: modified
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type InstanceNode struct {
	// Id
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// Node details
	InstanceClusterNode []*InstanceClusterNode `json:"InstanceClusterNode,omitempty" name:"InstanceClusterNode"`
}

type InstanceParam struct {
	// Parameter name, such as “timeout”. For supported custom parameters, see <a href="https://www.tencentcloud.com/document/product/239/39796">Setting Instance Parameters</a>
	Key *string `json:"Key,omitempty" name:"Key"`

	// Current parameter value. For example, if you set the current value of “timeout” to 120 (in seconds), the client connections that remain idle longer than 120 seconds will be closed. For more information on parameter values, see <a href="https://www.tencentcloud.com/document/product/239/39796">Setting Instance Parameters</a>
	Value *string `json:"Value,omitempty" name:"Value"`
}

type InstanceParamHistory struct {
	// Parameter name
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`

	// Value before modification
	PreValue *string `json:"PreValue,omitempty" name:"PreValue"`

	// Value after modification
	NewValue *string `json:"NewValue,omitempty" name:"NewValue"`

	// Status. 1: modifying the parameter configuration; 2: modified the parameter configuration successfully; 3: failed to modify the parameter configuration
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Modification time
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type InstanceProxySlowlogDetail struct {
	// Duration of the slow query in ms.
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`

	// Client address
	Client *string `json:"Client,omitempty" name:"Client"`

	// Command
	Command *string `json:"Command,omitempty" name:"Command"`

	// Command line details
	CommandLine *string `json:"CommandLine,omitempty" name:"CommandLine"`

	// Execution time
	ExecuteTime *string `json:"ExecuteTime,omitempty" name:"ExecuteTime"`
}

type InstanceSecurityGroupDetail struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Security group information
	SecurityGroupDetails []*SecurityGroupDetail `json:"SecurityGroupDetails,omitempty" name:"SecurityGroupDetails"`
}

type InstanceSet struct {
	// Instance name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// User AppID
	Appid *int64 `json:"Appid,omitempty" name:"Appid"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Region ID. <ul><li>`1`: Guangzhou. </li><li>`4`: Shanghai. </li><li>`5`: Hong Kong (China). </li><li>`6`: Toronto. </li> <li>`7`: Shanghai Finance. </li> <li>`8`: Beijing. </li> <li>`9`: Singapore. </li> <li>`11`: Shenzhen Finance. </li> <li>`15`: West US (Silicon Valley). </li><li>`16`: Chengdu. </li><li>`17`: Frankfurt. </li><li>`18`: Seoul. </li><li>`19`: Chongqing. </li><li>`21`: Mumbai. </li><li>`22`: East US (Virginia). </li><li>`23`: Bangkok. </li><li>`24`: Moscow. </li><li>`25`: Tokyo. </li></ul>
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// Region ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// VPC ID, such as `75101`.
	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID, such as `46315`.
	SubnetId *int64 `json:"SubnetId,omitempty" name:"SubnetId"`

	// Current instance status. <ul><li>`0`: To be initialized. </li><li>`1`: In the process. </li><li>`2`: Running. </li><li>`-2`: Isolated. </li><li>`-3`: To be deleted. </li></ul>
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Instance VIP
	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`

	// Port number of an instance
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Instance creation time
	Createtime *string `json:"Createtime,omitempty" name:"Createtime"`

	// Instance capacity in MB
	Size *float64 `json:"Size,omitempty" name:"Size"`

	// This field has been disused. You can use the [GetMonitorData](https://intl.cloud.tencent.com/document/product/248/31014?from_cn_redirect=1) API to query the capacity used by the instance.
	SizeUsed *float64 `json:"SizeUsed,omitempty" name:"SizeUsed"`

	// Instance type. <ul><li>`1`: Redis 2.8 memory edition in cluster architecture. </li><li>`2`: Redis 2.8 memory edition in standard architecture. </li><li>`3`: CKV 3.2 memory edition in standard architecture. </li><li>`4`: CKV 3.2 memory edition in cluster architecture. </li><li>`5`: Redis 2.8 memory edition in standalone architecture. </li></li><li>`6`: Redis 4.0 memory edition in standard architecture. </li></li><li>`7`: Redis 4.0 memory edition in cluster architecture. </li></li><li>`8`: Redis 5.0 memory edition in standard architecture. </li></li><li>`9`: Redis 5.0 memory edition in cluster architecture. </li></ul>
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// Whether to set the auto-renewal flag for an instance. <ul><li>`1`: Auto-renewal set. </li><li>`0`: Auto-renewal not set.</li></ul>
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// Instance expiration time
	DeadlineTime *string `json:"DeadlineTime,omitempty" name:"DeadlineTime"`

	// Engine: Redis community edition, Tencent Cloud CKV
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// Product type. <ul><li>`standalone`: Standard edition. </li><li>`cluster`: Cluster edition. </li></ul>
	ProductType *string `json:"ProductType,omitempty" name:"ProductType"`

	// VPC ID, such as vpc-fk33jsf43kgv.
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// VPC subnet ID, such as subnet-fd3j6l35mm0.
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// Billing mode. Only pay-as-you-go billing is supported.
	BillingMode *int64 `json:"BillingMode,omitempty" name:"BillingMode"`

	// Description of an instance status, such as "Running".
	InstanceTitle *string `json:"InstanceTitle,omitempty" name:"InstanceTitle"`

	// Scheduled deactivation time
	OfflineTime *string `json:"OfflineTime,omitempty" name:"OfflineTime"`

	// Sub-status returned for an instance in process
	SubStatus *int64 `json:"SubStatus,omitempty" name:"SubStatus"`

	// Anti-affinity tag
	Tags []*string `json:"Tags,omitempty" name:"Tags"`

	// Instance node information
	InstanceNode []*InstanceNode `json:"InstanceNode,omitempty" name:"InstanceNode"`

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

	// Read weight of a replica node
	SlaveReadWeight *int64 `json:"SlaveReadWeight,omitempty" name:"SlaveReadWeight"`

	// Instance tag information
	// Note: This field may return null, indicating that no valid value can be obtained.
	InstanceTags []*InstanceTagInfo `json:"InstanceTags,omitempty" name:"InstanceTags"`

	// Project name
	// Note: This field may return null, indicating that no valid value can be obtained.
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// Whether an instance is password-free. <ul><li>`true`: Yes. </li><li>`false`: No. </li></ul>
	// Note: This field may return null, indicating that no valid value can be obtained.
	NoAuth *bool `json:"NoAuth,omitempty" name:"NoAuth"`

	// Number of client connections
	// Note: This field may return null, indicating that no valid value can be obtained.
	ClientLimit *int64 `json:"ClientLimit,omitempty" name:"ClientLimit"`

	// DTS status (internal parameter, which can be ignored)
	// Note: This field may return null, indicating that no valid value can be obtained.
	DtsStatus *int64 `json:"DtsStatus,omitempty" name:"DtsStatus"`

	// Upper shard bandwidth limit in MB
	// Note: This field may return null, indicating that no valid value can be obtained.
	NetLimit *int64 `json:"NetLimit,omitempty" name:"NetLimit"`

	// Password-free instance flag (internal parameter, which can be ignored)
	// Note: This field may return null, indicating that no valid value can be obtained.
	PasswordFree *int64 `json:"PasswordFree,omitempty" name:"PasswordFree"`

	// Internal parameter, which can be ignored.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Vip6 *string `json:"Vip6,omitempty" name:"Vip6"`

	// Read-only instance flag (internal parameter, which can be ignored)
	// Note: This field may return null, indicating that no valid value can be obtained.
	ReadOnly *int64 `json:"ReadOnly,omitempty" name:"ReadOnly"`

	// Internal parameter, which can be ignored.
	// Note: This field may return null, indicating that no valid value can be obtained.
	RemainBandwidthDuration *string `json:"RemainBandwidthDuration,omitempty" name:"RemainBandwidthDuration"`

	// This parameter can be ignored for Redis instance.
	// Note: This field may return null, indicating that no valid value can be obtained.
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// Monitoring granularity type. <ul><li>`1m`: Monitoring at 1-minute granularity. </li><li>`5s`: Monitoring at 5-second granularity. </li></ul>
	// Note: This field may return null, indicating that no valid values can be obtained.
	MonitorVersion *string `json:"MonitorVersion,omitempty" name:"MonitorVersion"`

	// The minimum number of max client connections
	// Note: This field may return null, indicating that no valid value can be obtained.
	ClientLimitMin *int64 `json:"ClientLimitMin,omitempty" name:"ClientLimitMin"`

	// The maximum number of max client connections
	// Note: This field may return null, indicating that no valid value can be obtained.
	ClientLimitMax *int64 `json:"ClientLimitMax,omitempty" name:"ClientLimitMax"`

	// Instance node details
	// Note: This field may return null, indicating that no valid value can be obtained.
	NodeSet []*RedisNodeInfo `json:"NodeSet,omitempty" name:"NodeSet"`

	// Information of the region where the instance is deployed, such as `ap-guangzhou`.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Region *string `json:"Region,omitempty" name:"Region"`

	// Public IP
	// Note: This field may return null, indicating that no valid value can be obtained.
	WanAddress *string `json:"WanAddress,omitempty" name:"WanAddress"`

	// Polaris service address
	// Note: This field may return null, indicating that no valid value can be obtained.
	PolarisServer *string `json:"PolarisServer,omitempty" name:"PolarisServer"`

	// The current proxy version of an instance
	// Note: This field may return null, indicating that no valid value can be obtained.
	CurrentProxyVersion *string `json:"CurrentProxyVersion,omitempty" name:"CurrentProxyVersion"`

	// The current cache minor version of an instance
	// Note: This field may return null, indicating that no valid value can be obtained.
	CurrentRedisVersion *string `json:"CurrentRedisVersion,omitempty" name:"CurrentRedisVersion"`

	// Proxy version, which can be upgraded for the instance
	// Note: This field may return null, indicating that no valid value can be obtained.
	UpgradeProxyVersion *string `json:"UpgradeProxyVersion,omitempty" name:"UpgradeProxyVersion"`

	// Cache minor version, which can be upgraded for the instance
	// Note: This field may return null, indicating that no valid value can be obtained.
	UpgradeRedisVersion *string `json:"UpgradeRedisVersion,omitempty" name:"UpgradeRedisVersion"`
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

	// Node ID
	Node *string `json:"Node,omitempty" name:"Node"`
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

	// Whether restart is required after a modification is made. Valid values: true, false
	NeedRestart *string `json:"NeedRestart,omitempty" name:"NeedRestart"`

	// Default value of the parameter
	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`

	// Current value
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`

	// Parameter description
	Tips *string `json:"Tips,omitempty" name:"Tips"`

	// Valid values of the parameter
	TextValue []*string `json:"TextValue,omitempty" name:"TextValue"`

	// Parameter status. 1: modifying; 2: modified
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type Instances struct {
	// User AppID
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Region ID. <ul><li>`1`: Guangzhou. </li><li>`4`: Shanghai. </li><li>`5`: Hong Kong (China). </li> <li>`6`: Toronto. </li> <li>`7`: Shanghai Finance. </li> <li>`8`: Beijing. </li> <li>`9`: Singapore. </li> <li>`11`: Shenzhen Finance. </li> <li>`15`: West US (Silicon Valley). </li> </ul>
	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`

	// Region ID
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// Number of replicas
	RedisReplicasNum *uint64 `json:"RedisReplicasNum,omitempty" name:"RedisReplicasNum"`

	// Number of shards
	RedisShardNum *int64 `json:"RedisShardNum,omitempty" name:"RedisShardNum"`

	// Shard memory size.
	RedisShardSize *int64 `json:"RedisShardSize,omitempty" name:"RedisShardSize"`

	// Instance disk size
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// Engine: Redis Community Edition, Tencent Cloud CKV.
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// Read-write permission of the instance. <ul><li>`rw`: Read/Write. </li><li>`r`: Read-only. </li></ul>
	Role *string `json:"Role,omitempty" name:"Role"`

	// Instance VIP
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Internal parameter, which can be ignored.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Vip6 *string `json:"Vip6,omitempty" name:"Vip6"`

	// VPC ID, such as `75101`.
	VpcID *int64 `json:"VpcID,omitempty" name:"VpcID"`

	// Instance port
	VPort *int64 `json:"VPort,omitempty" name:"VPort"`

	// Instance status. <ul><li>`0`: Uninitialized. </li><li>`1`: In the process. </li><li>`2`: Running. </li><li>`-2`: Isolated. </li><li>`-3`: To be deleted. </li></ul>
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Repository ID
	GrocerySysId *int64 `json:"GrocerySysId,omitempty" name:"GrocerySysId"`

	// Instance type. <ul><li>`1`: Redis 2.8 Memory Edition (Cluster Architecture). </li><li>`2`: Redis 2.8 Memory Edition (Standard Architecture). </li><li>`3`: CKV 3.2 Memory Edition (Standard Architecture). </li><li>`4`: CKV 3.2 Memory Edition (Cluster Architecture). </li><li>`5`: Redis 2.8 Standalone Edition. </li><li>`6`: Redis 4.0 Memory Edition (Standard Architecture). </li><li>`7`: Redis 4.0 Memory Edition (Cluster Architecture). </li><li>`8`: Redis 5.0 Memory Edition (Standard Architecture). </li><li>`9`: Redis 5.0 Memory Edition (Cluster Architecture). </li></ul>
	ProductType *int64 `json:"ProductType,omitempty" name:"ProductType"`

	// The time when the instance was added to the replication group.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// The time when instances in the replication group were updated.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

// Predefined struct for user
type KillMasterGroupRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 1. The password must contain 8–30 characters. A password of 12 or more characters is recommended.
	// 2. It cannot start with a slash (/).
	// 3. It must contain characters in at least two of the following types:
	//     a. Lowercase letters (a–z)
	//     b. Uppercase letters (A–Z)
	//     c. Digits (0–9)
	//     d. ()`~!@#$%^&*-+=_|{}[]:;<>,.?/
	Password *string `json:"Password,omitempty" name:"Password"`

	// Node information of a single-AZ deployed instance
	ShardIds []*int64 `json:"ShardIds,omitempty" name:"ShardIds"`
}

type KillMasterGroupRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 1. The password must contain 8–30 characters. A password of 12 or more characters is recommended.
	// 2. It cannot start with a slash (/).
	// 3. It must contain characters in at least two of the following types:
	//     a. Lowercase letters (a–z)
	//     b. Uppercase letters (A–Z)
	//     c. Digits (0–9)
	//     d. ()`~!@#$%^&*-+=_|{}[]:;<>,.?/
	Password *string `json:"Password,omitempty" name:"Password"`

	// Node information of a single-AZ deployed instance
	ShardIds []*int64 `json:"ShardIds,omitempty" name:"ShardIds"`
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
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

// Predefined struct for user
type ManualBackupInstanceRequestParams struct {
	// ID of the instance to be operated on, which can be obtained through the `InstanceId` field in the return value of the `DescribeInstance` API.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Backup remarks
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Retention time in days. 0 indicates the default retention time.
	StorageDays *int64 `json:"StorageDays,omitempty" name:"StorageDays"`
}

type ManualBackupInstanceRequest struct {
	*tchttp.BaseRequest
	
	// ID of the instance to be operated on, which can be obtained through the `InstanceId` field in the return value of the `DescribeInstance` API.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Backup remarks
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Retention time in days. 0 indicates the default retention time.
	StorageDays *int64 `json:"StorageDays,omitempty" name:"StorageDays"`
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
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Old password of an instance
	OldPassword *string `json:"OldPassword,omitempty" name:"OldPassword"`

	// New password of an instance
	Password *string `json:"Password,omitempty" name:"Password"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModfiyInstancePasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModfiyInstancePasswordResponseParams struct {
	// Task ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Automatic backup cycle. Valid values: `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, `Sunday`. This parameter currently cannot be modified.
	WeekDays []*string `json:"WeekDays,omitempty" name:"WeekDays"`

	// Automatic backup time in the format of 00:00-01:00, 01:00-02:00... 23:00-00:00.
	TimePeriod *string `json:"TimePeriod,omitempty" name:"TimePeriod"`

	// Automatic backup type. `1`: Scheduled rollback.
	AutoBackupType *int64 `json:"AutoBackupType,omitempty" name:"AutoBackupType"`
}

type ModifyAutoBackupConfigRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Automatic backup cycle. Valid values: `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, `Sunday`. This parameter currently cannot be modified.
	WeekDays []*string `json:"WeekDays,omitempty" name:"WeekDays"`

	// Automatic backup time in the format of 00:00-01:00, 01:00-02:00... 23:00-00:00.
	TimePeriod *string `json:"TimePeriod,omitempty" name:"TimePeriod"`

	// Automatic backup type. `1`: Scheduled rollback.
	AutoBackupType *int64 `json:"AutoBackupType,omitempty" name:"AutoBackupType"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAutoBackupConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAutoBackupConfigResponseParams struct {
	// Automatic backup type: 1 (scheduled rollback)
	AutoBackupType *int64 `json:"AutoBackupType,omitempty" name:"AutoBackupType"`

	// Automatic backup cycle. Valid values: `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, `Sunday`.
	WeekDays []*string `json:"WeekDays,omitempty" name:"WeekDays"`

	// Automatic backup time in the format of 00:00-01:00, 01:00-02:00... 23:00-00:00.
	TimePeriod *string `json:"TimePeriod,omitempty" name:"TimePeriod"`

	// Retention time of full backup files in days
	BackupStorageDays *int64 `json:"BackupStorageDays,omitempty" name:"BackupStorageDays"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	LimitType *string `json:"LimitType,omitempty" name:"LimitType"`

	// Only `In` can be passed in for this parameter, indicating that backup files can be downloaded in the custom `LimitVpc`.
	VpcComparisonSymbol *string `json:"VpcComparisonSymbol,omitempty" name:"VpcComparisonSymbol"`

	// Whether backups can be downloaded at the custom `LimitIp` address.
	// 
	// - `In`: Download is allowed for the custom IP.
	// - `NotIn`: Download is not allowed for the custom IP.
	IpComparisonSymbol *string `json:"IpComparisonSymbol,omitempty" name:"IpComparisonSymbol"`

	// VPC ID of the custom backup file download address, which is required if `LimitType` is `Customize`.
	LimitVpc []*BackupLimitVpcItem `json:"LimitVpc,omitempty" name:"LimitVpc"`

	// VPC IP of the custom backup file download address, which is required if `LimitType` is `Customize`.
	LimitIp []*string `json:"LimitIp,omitempty" name:"LimitIp"`
}

type ModifyBackupDownloadRestrictionRequest struct {
	*tchttp.BaseRequest
	
	// Type of the network restrictions for downloading backup files. Valid values:
	// 
	// - `NoLimit`: Backup files can be downloaded over both public and private networks.
	// - `LimitOnlyIntranet`: Backup files can be downloaded only at private network addresses auto-assigned by Tencent Cloud.
	// - `Customize`: Backup files can be downloaded only in the customized VPC.
	LimitType *string `json:"LimitType,omitempty" name:"LimitType"`

	// Only `In` can be passed in for this parameter, indicating that backup files can be downloaded in the custom `LimitVpc`.
	VpcComparisonSymbol *string `json:"VpcComparisonSymbol,omitempty" name:"VpcComparisonSymbol"`

	// Whether backups can be downloaded at the custom `LimitIp` address.
	// 
	// - `In`: Download is allowed for the custom IP.
	// - `NotIn`: Download is not allowed for the custom IP.
	IpComparisonSymbol *string `json:"IpComparisonSymbol,omitempty" name:"IpComparisonSymbol"`

	// VPC ID of the custom backup file download address, which is required if `LimitType` is `Customize`.
	LimitVpc []*BackupLimitVpcItem `json:"LimitVpc,omitempty" name:"LimitVpc"`

	// VPC IP of the custom backup file download address, which is required if `LimitType` is `Customize`.
	LimitIp []*string `json:"LimitIp,omitempty" name:"LimitIp"`
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
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ModifyDBInstanceSecurityGroupsRequestParams struct {
	// Database engine name, which is `redis` for this API.
	Product *string `json:"Product,omitempty" name:"Product"`

	// List of IDs of security groups to be modified, which is an array of one or more security group IDs.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type ModifyDBInstanceSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Database engine name, which is `redis` for this API.
	Product *string `json:"Product,omitempty" name:"Product"`

	// List of IDs of security groups to be modified, which is an array of one or more security group IDs.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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
type ModifyInstanceAccountRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Sub-account name. If the root account is to be modified, enter `root`.
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// Sub-account password
	AccountPassword *string `json:"AccountPassword,omitempty" name:"AccountPassword"`

	// Sub-account description information
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Routing policy. Valid values: master (master node); replication (replica node)
	ReadonlyPolicy []*string `json:"ReadonlyPolicy,omitempty" name:"ReadonlyPolicy"`

	// Sub-account read/write policy. Valid values: r (read-only); w (write-only); rw (read/write).
	Privilege *string `json:"Privilege,omitempty" name:"Privilege"`

	// true: make the root account password-free. This is applicable to root accounts only. Sub-accounts cannot be made password-free.
	NoAuth *bool `json:"NoAuth,omitempty" name:"NoAuth"`
}

type ModifyInstanceAccountRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Sub-account name. If the root account is to be modified, enter `root`.
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// Sub-account password
	AccountPassword *string `json:"AccountPassword,omitempty" name:"AccountPassword"`

	// Sub-account description information
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Routing policy. Valid values: master (master node); replication (replica node)
	ReadonlyPolicy []*string `json:"ReadonlyPolicy,omitempty" name:"ReadonlyPolicy"`

	// Sub-account read/write policy. Valid values: r (read-only); w (write-only); rw (read/write).
	Privilege *string `json:"Privilege,omitempty" name:"Privilege"`

	// true: make the root account password-free. This is applicable to root accounts only. Sub-accounts cannot be made password-free.
	NoAuth *bool `json:"NoAuth,omitempty" name:"NoAuth"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceAccountResponseParams struct {
	// Task ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ModifyInstanceParamsRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// List of instance parameters modified
	InstanceParams []*InstanceParam `json:"InstanceParams,omitempty" name:"InstanceParams"`
}

type ModifyInstanceParamsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// List of instance parameters modified
	InstanceParams []*InstanceParam `json:"InstanceParams,omitempty" name:"InstanceParams"`
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
	// Whether the parameter is modified successfully. <br><li>`True`: Yes<br><li>`False`: No<br>
	Changed *bool `json:"Changed,omitempty" name:"Changed"`

	// ID of the task
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ModifyInstanceReadOnlyRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance input mode. Valid values: `0` (read/write), `1` (read-only)
	InputMode *string `json:"InputMode,omitempty" name:"InputMode"`
}

type ModifyInstanceReadOnlyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance input mode. Valid values: `0` (read/write), `1` (read-only)
	InputMode *string `json:"InputMode,omitempty" name:"InputMode"`
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
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Instance modification type. rename: rename an instance; modifyProject: modify the project of an instance; modifyAutoRenew: modify the auto-renewal flag of an instance.
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// Instance ID
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// New name of the instance
	InstanceNames []*string `json:"InstanceNames,omitempty" name:"InstanceNames"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Auto-renewal flag. 0: default status (manual renewal); 1: auto-renewal enabled; 2: auto-renewal disabled
	AutoRenews []*int64 `json:"AutoRenews,omitempty" name:"AutoRenews"`

	// Disused
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Disused
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Disused
	AutoRenew *int64 `json:"AutoRenew,omitempty" name:"AutoRenew"`
}

type ModifyInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance modification type. rename: rename an instance; modifyProject: modify the project of an instance; modifyAutoRenew: modify the auto-renewal flag of an instance.
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// Instance ID
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// New name of the instance
	InstanceNames []*string `json:"InstanceNames,omitempty" name:"InstanceNames"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Auto-renewal flag. 0: default status (manual renewal); 1: auto-renewal enabled; 2: auto-renewal disabled
	AutoRenews []*int64 `json:"AutoRenews,omitempty" name:"AutoRenews"`

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
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ModifyMaintenanceWindowRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Maintenance start time, such as 17:00
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Maintenance end time, such as 19:00
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type ModifyMaintenanceWindowRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Maintenance start time, such as 17:00
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Maintenance end time, such as 19:00
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
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
	Status *string `json:"Status,omitempty" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Network change type. Valid values:
	// - `changeVip`: VPC change, including the private IPv4 address and port.
	// - `changeVpc`: Subnet change.
	// - `changeBaseToVpc`: Change from classic network to VPC.
	// - `changeVPort`: Port change.
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// Private IPv4 address of the instance, which is required if `Operation` is `changeVip`.
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// VPC ID after the change, which is required if `Operation` is `changeVpc` or `changeBaseToVpc`.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID after the change, which is required if `Operation` is `changeVpc` or `changeBaseToVpc`.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Retention period of the original private IPv4 address
	// - Unit: Days.
	// - Valid values: `0`, `1`, `2`, `3`, `7`, `15`.
	// 
	// **Note**: You can set the retention period of the original address only in the latest SDK. In earlier SDKs, the original address is released immediately. To view the SDK version, go to [SDK Center](https://intl.cloud.tencent.com/document/sdk?from_cn_redirect=1).
	Recycle *int64 `json:"Recycle,omitempty" name:"Recycle"`

	// Network port after the change, which is required if `Operation` is `changeVPort` or `changeVip`. Value range: [1024,65535].
	VPort *int64 `json:"VPort,omitempty" name:"VPort"`
}

type ModifyNetworkConfigRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Network change type. Valid values:
	// - `changeVip`: VPC change, including the private IPv4 address and port.
	// - `changeVpc`: Subnet change.
	// - `changeBaseToVpc`: Change from classic network to VPC.
	// - `changeVPort`: Port change.
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// Private IPv4 address of the instance, which is required if `Operation` is `changeVip`.
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// VPC ID after the change, which is required if `Operation` is `changeVpc` or `changeBaseToVpc`.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID after the change, which is required if `Operation` is `changeVpc` or `changeBaseToVpc`.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Retention period of the original private IPv4 address
	// - Unit: Days.
	// - Valid values: `0`, `1`, `2`, `3`, `7`, `15`.
	// 
	// **Note**: You can set the retention period of the original address only in the latest SDK. In earlier SDKs, the original address is released immediately. To view the SDK version, go to [SDK Center](https://intl.cloud.tencent.com/document/sdk?from_cn_redirect=1).
	Recycle *int64 `json:"Recycle,omitempty" name:"Recycle"`

	// Network port after the change, which is required if `Operation` is `changeVPort` or `changeVip`. Value range: [1024,65535].
	VPort *int64 `json:"VPort,omitempty" name:"VPort"`
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
	Status *bool `json:"Status,omitempty" name:"Status"`

	// New subnet ID of the instance
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// New VPC ID of the instance
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// New private IPv4 address of the instance
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Task ID, which can be used to query the task execution status through the `DescribeTaskInfo` API.
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// ID of the source parameter template.
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// New name after the parameter template is modified.
	Name *string `json:"Name,omitempty" name:"Name"`

	// New description after the parameter template is modified.
	Description *string `json:"Description,omitempty" name:"Description"`

	// New parameter list after the parameter template is modified.
	ParamList []*InstanceParam `json:"ParamList,omitempty" name:"ParamList"`
}

type ModifyParamTemplateRequest struct {
	*tchttp.BaseRequest
	
	// ID of the source parameter template.
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// New name after the parameter template is modified.
	Name *string `json:"Name,omitempty" name:"Name"`

	// New description after the parameter template is modified.
	Description *string `json:"Description,omitempty" name:"Description"`

	// New parameter list after the parameter template is modified.
	ParamList []*InstanceParam `json:"ParamList,omitempty" name:"ParamList"`
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
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type OpenSSLRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type OpenSSLRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenSSLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenSSLResponseParams struct {
	// Task ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Action *string `json:"Action,omitempty" name:"Action"`

	// All the addresses that the address group ID represents.
	AddressModule *string `json:"AddressModule,omitempty" name:"AddressModule"`

	// Source IP or IP address range, such as 192.168.0.0/16.
	CidrIp *string `json:"CidrIp,omitempty" name:"CidrIp"`

	// Description.
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// Network protocol, such as UDP and TCP.
	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`

	// Port.
	PortRange *string `json:"PortRange,omitempty" name:"PortRange"`

	// All the protocols and ports that the service group ID represents.
	ServiceModule *string `json:"ServiceModule,omitempty" name:"ServiceModule"`

	// All the addresses that the security group ID represents.
	Id *string `json:"Id,omitempty" name:"Id"`
}

type ParamTemplateInfo struct {
	// Parameter template ID
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// Parameter template name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Parameter template description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Instance type. Valid values: `1` (Redis 2.8 Memory Edition in cluster architecture), `2` (Redis 2.8 Memory Edition in standard architecture), `3` (CKV 3.2 Memory Edition in standard architecture), `4` (CKV 3.2 Memory Edition in cluster architecture), `5` (Redis 2.8 Memory Edition in standalone architecture), `6` (Redis 4.0 Memory Edition in standard architecture), `7` (Redis 4.0 Memory Edition in cluster architecture), `8` (Redis 5.0 Memory Edition in standard architecture), `9` (Redis 5.0 Memory Edition in cluster architecture)
	ProductType *uint64 `json:"ProductType,omitempty" name:"ProductType"`
}

type ParameterDetail struct {
	// Parameter name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Data type of the parameter
	ParamType *string `json:"ParamType,omitempty" name:"ParamType"`

	// Default value of the parameter
	Default *string `json:"Default,omitempty" name:"Default"`

	// Parameter description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Current value
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`

	// Whether the database needs to be restarted for the modified parameter to take effect. Valid values: 0 (no); 1 (yes).
	NeedReboot *int64 `json:"NeedReboot,omitempty" name:"NeedReboot"`

	// Maximum value of the parameter
	// Note: This field may return null, indicating that no valid values can be obtained.
	Max *string `json:"Max,omitempty" name:"Max"`

	// Minimum value of the parameter
	// Note: This field may return null, indicating that no valid values can be obtained.
	Min *string `json:"Min,omitempty" name:"Min"`

	// Enumerated values of the parameter. It is null if the parameter is non-enumerated.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EnumValue []*string `json:"EnumValue,omitempty" name:"EnumValue"`
}

type ProductConf struct {
	// Product type. Valid values: `2` (Redis 2.8 Memory Edition in standard architecture), `3` (CKV 3.2 Memory Edition in standard architecture), `4` (CKV 3.2 Memory Edition in cluster architecture), `5` (Redis 2.8 Memory Edition in standalone architecture), `6` (Redis 4.0 Memory Edition in standard architecture), `7` (Redis 4.0 Memory Edition in cluster architecture), `8` (Redis 5.0 Memory Edition in standard architecture), `9` (Redis 5.0 Memory Edition in cluster architecture), `10` (Redis 4.0 Hybrid Storage Edition (Tendis)).
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// Product name: Redis Master-Replica Edition, CKV Master-Replica Edition, CKV Cluster Edition, Redis Standalone Edition, Redis Cluster Edition, Tendis Hybrid Storage Edition
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
	TotalSize []*string `json:"TotalSize,omitempty" name:"TotalSize"`

	// Shard size in GB
	ShardSize []*string `json:"ShardSize,omitempty" name:"ShardSize"`

	// Number of replicas
	ReplicaNum []*string `json:"ReplicaNum,omitempty" name:"ReplicaNum"`

	// Number of shards
	ShardNum []*string `json:"ShardNum,omitempty" name:"ShardNum"`

	// Supported billing method. 1: monthly subscription; 0: pay-as-you-go
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// Whether to support read-only replicas
	EnableRepicaReadOnly *bool `json:"EnableRepicaReadOnly,omitempty" name:"EnableRepicaReadOnly"`
}

type ProxyNodes struct {
	// Node ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	NodeId *string `json:"NodeId,omitempty" name:"NodeId"`

	// AZ ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
}

type RedisBackupSet struct {
	// Backup start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Backup ID
	BackupId *string `json:"BackupId,omitempty" name:"BackupId"`

	// Backup type
	// 
	// - `1`: Manual backup initiated by the user.
	// - `0`: Automatic backup in the early morning initiated by the system.
	BackupType *string `json:"BackupType,omitempty" name:"BackupType"`

	// Backup status 
	// 
	// - `1`: The backup is locked by another process.
	// - `2`: The backup is normal and not locked by any process.
	// - `-1`: The backup expired.
	// - `3`: The backup is being exported.
	// - `4`: The backup was exported successfully.
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Backup remarks
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Whether the backup is locked
	// 
	// - `0`: Not locked.
	// - `1`: Locked.
	Locked *int64 `json:"Locked,omitempty" name:"Locked"`

	// Internal field, which can be ignored.
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupSize *int64 `json:"BackupSize,omitempty" name:"BackupSize"`

	// Internal field, which can be ignored.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FullBackup *int64 `json:"FullBackup,omitempty" name:"FullBackup"`

	// Internal field, which can be ignored.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceType *int64 `json:"InstanceType,omitempty" name:"InstanceType"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// The region where the local backup resides.
	Region *string `json:"Region,omitempty" name:"Region"`

	// Backup end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Backup file type
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// Backup file expiration time
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

type RedisCommonInstanceList struct {
	// Instance name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// User ID
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// Instance project ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Instance region
	Region *string `json:"Region,omitempty" name:"Region"`

	// Instance AZ
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Instance network ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Instance status. 1: task running; 2: instance running; -2: instance isolated; -3: instance being eliminated; -4: instance eliminated
	Status *string `json:"Status,omitempty" name:"Status"`

	// Instance network IP
	Vips []*string `json:"Vips,omitempty" name:"Vips"`

	// Instance network port
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// Instance creation time
	Createtime *string `json:"Createtime,omitempty" name:"Createtime"`

	// Billing mode. 0: pay-as-you-go; 1: monthly subscription
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// Network type. Valid values: 0 (classic network); 1 (VPC).
	NetType *int64 `json:"NetType,omitempty" name:"NetType"`
}

type RedisNode struct {
	// Number of keys on the node
	Keys *int64 `json:"Keys,omitempty" name:"Keys"`

	// Distribution of node slots
	Slot *string `json:"Slot,omitempty" name:"Slot"`

	// Node ID
	NodeId *string `json:"NodeId,omitempty" name:"NodeId"`

	// Node status
	Status *string `json:"Status,omitempty" name:"Status"`

	// Node role
	Role *string `json:"Role,omitempty" name:"Role"`
}

type RedisNodeInfo struct {
	// Node type. <ul><li>`0`: Master node.</li><li>`1`: Replica node.</li></ul>
	NodeType *int64 `json:"NodeType,omitempty" name:"NodeType"`

	// Master or replica node ID. <ul><li>This parameter is optional when the [CreateInstances](https://intl.cloud.tencent.com/document/product/239/20026?from_cn_redirect=1) API is used to create a TencentDB for Redis instance, but it is required when the [UpgradeInstance](https://intl.cloud.tencent.com/document/product/239/20013?from_cn_redirect=1) API is used to adjust the configuration of an instance. </li><li>You can use the [DescribeInstances](https://intl.cloud.tencent.com/document/product/239/20018?from_cn_redirect=1) API to get the node ID of integer type. </li></ul>
	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`

	// ID of the AZ of the master or replica node
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// Name of the AZ of the master or replica node
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}

type RedisNodes struct {
	// Node ID
	NodeId *string `json:"NodeId,omitempty" name:"NodeId"`

	// Node role
	NodeRole *string `json:"NodeRole,omitempty" name:"NodeRole"`

	// Shard ID
	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`

	// AZ ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
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
	ZoneSet []*ZoneCapacityConf `json:"ZoneSet,omitempty" name:"ZoneSet"`
}

// Predefined struct for user
type ReleaseWanAddressRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type ReleaseWanAddressRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// Status of disabling public network access
	WanStatus *string `json:"WanStatus,omitempty" name:"WanStatus"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type RenewInstanceRequestParams struct {
	// Validity period in months
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// The parameter used to determine whether to modify the billing mode. <ul><li>If you want to change the billing mode from pay-as-you-go to monthly subscription, specify this parameter as <b>prepaid</b>. </li><li>If the current instance is monthly subscribed, this parameter is not required. </li></ul>
	ModifyPayMode *string `json:"ModifyPayMode,omitempty" name:"ModifyPayMode"`
}

type RenewInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Validity period in months
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// The parameter used to determine whether to modify the billing mode. <ul><li>If you want to change the billing mode from pay-as-you-go to monthly subscription, specify this parameter as <b>prepaid</b>. </li><li>If the current instance is monthly subscribed, this parameter is not required. </li></ul>
	ModifyPayMode *string `json:"ModifyPayMode,omitempty" name:"ModifyPayMode"`
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
	DealId *string `json:"DealId,omitempty" name:"DealId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// Node group name, which is empty for the master node
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// Node AZ ID, such as ap-guangzhou-1
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Node group type. Valid values: master (master node group); replica (replica node group)
	Role *string `json:"Role,omitempty" name:"Role"`

	// List of nodes in the node group
	RedisNodes []*RedisNode `json:"RedisNodes,omitempty" name:"RedisNodes"`
}

// Predefined struct for user
type ResetPasswordRequestParams struct {
	// Redis instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Password reset (this parameter can be left blank when switching to password-free instance mode and is required in other cases)
	Password *string `json:"Password,omitempty" name:"Password"`

	// Whether to switch to password-free instance mode. false: switch to password-enabled instance mode; true: switch to password-free instance mode. Default value: false.
	NoAuth *bool `json:"NoAuth,omitempty" name:"NoAuth"`
}

type ResetPasswordRequest struct {
	*tchttp.BaseRequest
	
	// Redis instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Password reset (this parameter can be left blank when switching to password-free instance mode and is required in other cases)
	Password *string `json:"Password,omitempty" name:"Password"`

	// Whether to switch to password-free instance mode. false: switch to password-enabled instance mode; true: switch to password-free instance mode. Default value: false.
	NoAuth *bool `json:"NoAuth,omitempty" name:"NoAuth"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetPasswordResponseParams struct {
	// Task ID (this parameter is the task ID when changing a password. If you are switching between password-free and password-enabled instance mode, you can leave this parameter alone)
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type ResourceTag struct {
	// Tag key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

// Predefined struct for user
type RestoreInstanceRequestParams struct {
	// ID of the instance to be operated on, which can be obtained through the `InstanceId` field in the return value of the `DescribeInstances` API.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Backup ID, which can be obtained through the `backupId` field in the return value of the `GetRedisBackupList` API.
	BackupId *string `json:"BackupId,omitempty" name:"BackupId"`

	// Instance password, which needs to be validated during instance restoration (this parameter is not required for password-free instances)
	Password *string `json:"Password,omitempty" name:"Password"`
}

type RestoreInstanceRequest struct {
	*tchttp.BaseRequest
	
	// ID of the instance to be operated on, which can be obtained through the `InstanceId` field in the return value of the `DescribeInstances` API.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Backup ID, which can be obtained through the `backupId` field in the return value of the `GetRedisBackupList` API.
	BackupId *string `json:"BackupId,omitempty" name:"BackupId"`

	// Instance password, which needs to be validated during instance restoration (this parameter is not required for password-free instances)
	Password *string `json:"Password,omitempty" name:"Password"`
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
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Outbound []*Outbound `json:"Outbound,omitempty" name:"Outbound"`

	// Inbound rule.
	Inbound []*Inbound `json:"Inbound,omitempty" name:"Inbound"`
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
	InboundRule []*SecurityGroupsInboundAndOutbound `json:"InboundRule,omitempty" name:"InboundRule"`

	// Security group outbound rule
	OutboundRule []*SecurityGroupsInboundAndOutbound `json:"OutboundRule,omitempty" name:"OutboundRule"`
}

type SecurityGroupsInboundAndOutbound struct {
	// Action to be executed
	Action *string `json:"Action,omitempty" name:"Action"`

	// IP addresses
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

// Predefined struct for user
type StartupInstanceRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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
	// Task ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type SwitchInstanceVipRequestParams struct {
	// Source instance ID
	SrcInstanceId *string `json:"SrcInstanceId,omitempty" name:"SrcInstanceId"`

	// Target instance ID
	DstInstanceId *string `json:"DstInstanceId,omitempty" name:"DstInstanceId"`

	// The time that lapses in seconds since DTS is disconnected between the source instance and the target instance. If the DTS disconnection time period is greater than TimeDelay, the VIP will not be switched. We recommend you set an acceptable value based on the actual business conditions.
	TimeDelay *int64 `json:"TimeDelay,omitempty" name:"TimeDelay"`

	// Whether to force the switch when DTS is disconnected. 1: yes; 0: no.
	ForceSwitch *int64 `json:"ForceSwitch,omitempty" name:"ForceSwitch"`

	// now: switch now; syncComplete: switch after sync is completed
	SwitchTime *string `json:"SwitchTime,omitempty" name:"SwitchTime"`
}

type SwitchInstanceVipRequest struct {
	*tchttp.BaseRequest
	
	// Source instance ID
	SrcInstanceId *string `json:"SrcInstanceId,omitempty" name:"SrcInstanceId"`

	// Target instance ID
	DstInstanceId *string `json:"DstInstanceId,omitempty" name:"DstInstanceId"`

	// The time that lapses in seconds since DTS is disconnected between the source instance and the target instance. If the DTS disconnection time period is greater than TimeDelay, the VIP will not be switched. We recommend you set an acceptable value based on the actual business conditions.
	TimeDelay *int64 `json:"TimeDelay,omitempty" name:"TimeDelay"`

	// Whether to force the switch when DTS is disconnected. 1: yes; 0: no.
	ForceSwitch *int64 `json:"ForceSwitch,omitempty" name:"ForceSwitch"`

	// now: switch now; syncComplete: switch after sync is completed
	SwitchTime *string `json:"SwitchTime,omitempty" name:"SwitchTime"`
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
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance ProxyID
	ProxyID *string `json:"ProxyID,omitempty" name:"ProxyID"`
}

type SwitchProxyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance ProxyID
	ProxyID *string `json:"ProxyID,omitempty" name:"ProxyID"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchProxyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type TendisNodes struct {
	// Node ID
	NodeId *string `json:"NodeId,omitempty" name:"NodeId"`

	// Node role
	NodeRole *string `json:"NodeRole,omitempty" name:"NodeRole"`
}

type TendisSlowLogDetail struct {
	// Execution time
	ExecuteTime *string `json:"ExecuteTime,omitempty" name:"ExecuteTime"`

	// Duration of the slow query (ms)
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`

	// Command
	Command *string `json:"Command,omitempty" name:"Command"`

	// Command line details
	CommandLine *string `json:"CommandLine,omitempty" name:"CommandLine"`

	// Node ID
	Node *string `json:"Node,omitempty" name:"Node"`
}

type TradeDealDetail struct {
	// Order ID, which is used when a TencentCloud API is called.
	DealId *string `json:"DealId,omitempty" name:"DealId"`

	// Long order ID, which is used when an order issue is submitted for assistance.
	DealName *string `json:"DealName,omitempty" name:"DealName"`

	// AZ ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// Number of instances associated with the order
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// Creator UIN
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

	// Actual total price of the order in 0.01 CNY
	Price *int64 `json:"Price,omitempty" name:"Price"`

	// Instance ID
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

// Predefined struct for user
type UpgradeInstanceRequestParams struct {
	// The ID of instance to be modified.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// New memory size of an instance shard. <ul><li>Unit: MB. </li><li>You can only modify one of the three parameters at a time: `MemSize`, `RedisShardNum`, and `RedisReplicasNum`. To modify one of them, you need to enter the other two, which are consistent with the original configuration specifications of the instance. </li><li>In case of capacity reduction, the new specification must be at least 1.3 times the used capacity; otherwise, the operation will fail.</li></ul>
	MemSize *uint64 `json:"MemSize,omitempty" name:"MemSize"`

	// New number of instance shards. <ul><li>This parameter is not required for standard architecture instances, but for cluster architecture instances. </li><li>For cluster architecture, you can only modify one of the three parameters at a time: `MemSize`, `RedisShardNum`, and `RedisReplicasNum`. To modify one of them, you need to enter the other two, which are consistent with the original configuration specifications of the instance. </li></ul>
	RedisShardNum *uint64 `json:"RedisShardNum,omitempty" name:"RedisShardNum"`

	// New replica quantity. <ul><li>You can only modify one of the three parameters at a time: `MemSize`, `RedisShardNum`, and `RedisReplicasNum`. To modify one of them, you need to enter the other two, which are consistent with the original configuration specifications of the instance. </li></ul>To modify the number of replicas in a multi-AZ instance, `NodeSet` must be passed in.</li></ul>
	RedisReplicasNum *uint64 `json:"RedisReplicasNum,omitempty" name:"RedisReplicasNum"`

	// Additional information for adding replicas for multi-AZ instances, including replica AZ and type (`NodeType` is `1`). This parameter is not required for single-AZ instances.
	NodeSet []*RedisNodeInfo `json:"NodeSet,omitempty" name:"NodeSet"`
}

type UpgradeInstanceRequest struct {
	*tchttp.BaseRequest
	
	// The ID of instance to be modified.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// New memory size of an instance shard. <ul><li>Unit: MB. </li><li>You can only modify one of the three parameters at a time: `MemSize`, `RedisShardNum`, and `RedisReplicasNum`. To modify one of them, you need to enter the other two, which are consistent with the original configuration specifications of the instance. </li><li>In case of capacity reduction, the new specification must be at least 1.3 times the used capacity; otherwise, the operation will fail.</li></ul>
	MemSize *uint64 `json:"MemSize,omitempty" name:"MemSize"`

	// New number of instance shards. <ul><li>This parameter is not required for standard architecture instances, but for cluster architecture instances. </li><li>For cluster architecture, you can only modify one of the three parameters at a time: `MemSize`, `RedisShardNum`, and `RedisReplicasNum`. To modify one of them, you need to enter the other two, which are consistent with the original configuration specifications of the instance. </li></ul>
	RedisShardNum *uint64 `json:"RedisShardNum,omitempty" name:"RedisShardNum"`

	// New replica quantity. <ul><li>You can only modify one of the three parameters at a time: `MemSize`, `RedisShardNum`, and `RedisReplicasNum`. To modify one of them, you need to enter the other two, which are consistent with the original configuration specifications of the instance. </li></ul>To modify the number of replicas in a multi-AZ instance, `NodeSet` must be passed in.</li></ul>
	RedisReplicasNum *uint64 `json:"RedisReplicasNum,omitempty" name:"RedisReplicasNum"`

	// Additional information for adding replicas for multi-AZ instances, including replica AZ and type (`NodeType` is `1`). This parameter is not required for single-AZ instances.
	NodeSet []*RedisNodeInfo `json:"NodeSet,omitempty" name:"NodeSet"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeInstanceResponseParams struct {
	// Order ID
	DealId *string `json:"DealId,omitempty" name:"DealId"`

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

// Predefined struct for user
type UpgradeInstanceVersionRequestParams struct {
	// Target instance type after the change, which is the same as the `Type` of the [CreateInstances](https://intl.cloud.tencent.com/document/api/239/20026?from_cn_redirect=1) API.
	TargetInstanceType *string `json:"TargetInstanceType,omitempty" name:"TargetInstanceType"`

	// Switch mode. Valid values: 1 (switch during the maintenance time), 2 (switch now).
	SwitchOption *int64 `json:"SwitchOption,omitempty" name:"SwitchOption"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type UpgradeInstanceVersionRequest struct {
	*tchttp.BaseRequest
	
	// Target instance type after the change, which is the same as the `Type` of the [CreateInstances](https://intl.cloud.tencent.com/document/api/239/20026?from_cn_redirect=1) API.
	TargetInstanceType *string `json:"TargetInstanceType,omitempty" name:"TargetInstanceType"`

	// Switch mode. Valid values: 1 (switch during the maintenance time), 2 (switch now).
	SwitchOption *int64 `json:"SwitchOption,omitempty" name:"SwitchOption"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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
	DealId *string `json:"DealId,omitempty" name:"DealId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// The current proxy version
	CurrentProxyVersion *string `json:"CurrentProxyVersion,omitempty" name:"CurrentProxyVersion"`

	// Upgradeable redis version
	UpgradeProxyVersion *string `json:"UpgradeProxyVersion,omitempty" name:"UpgradeProxyVersion"`

	// `1` (upgrade immediately), `0` (upgrade during maintenance time)
	InstanceTypeUpgradeNow *int64 `json:"InstanceTypeUpgradeNow,omitempty" name:"InstanceTypeUpgradeNow"`
}

type UpgradeProxyVersionRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// The current proxy version
	CurrentProxyVersion *string `json:"CurrentProxyVersion,omitempty" name:"CurrentProxyVersion"`

	// Upgradeable redis version
	UpgradeProxyVersion *string `json:"UpgradeProxyVersion,omitempty" name:"UpgradeProxyVersion"`

	// `1` (upgrade immediately), `0` (upgrade during maintenance time)
	InstanceTypeUpgradeNow *int64 `json:"InstanceTypeUpgradeNow,omitempty" name:"InstanceTypeUpgradeNow"`
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
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// The current redis version
	CurrentRedisVersion *string `json:"CurrentRedisVersion,omitempty" name:"CurrentRedisVersion"`

	// Upgradeable redis version
	UpgradeRedisVersion *string `json:"UpgradeRedisVersion,omitempty" name:"UpgradeRedisVersion"`

	// `1` (upgrade immediately), `0` (upgrade during maintenance time)
	InstanceTypeUpgradeNow *int64 `json:"InstanceTypeUpgradeNow,omitempty" name:"InstanceTypeUpgradeNow"`
}

type UpgradeSmallVersionRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// The current redis version
	CurrentRedisVersion *string `json:"CurrentRedisVersion,omitempty" name:"CurrentRedisVersion"`

	// Upgradeable redis version
	UpgradeRedisVersion *string `json:"UpgradeRedisVersion,omitempty" name:"UpgradeRedisVersion"`

	// `1` (upgrade immediately), `0` (upgrade during maintenance time)
	InstanceTypeUpgradeNow *int64 `json:"InstanceTypeUpgradeNow,omitempty" name:"InstanceTypeUpgradeNow"`
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
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Whether to support “Reading Local Nodes Only” feature after upgrading to multi-AZ deployment.
	// ul><li>`true`: The “Read Local Nodes Only” feature is supported. During the upgrade, you need to upgrade the proxy version and Redis kernel minor version simultaneously, which will involve data migration and may take hours to complete. </li><li>`false`: The “Read Local Nodes Only” feature is not supported. Upgrading to multi-AZ deployment will involve metadata migration only without affecting the service, which generally take less than three minutes to complete.</li></ul>
	UpgradeProxyAndRedisServer *bool `json:"UpgradeProxyAndRedisServer,omitempty" name:"UpgradeProxyAndRedisServer"`
}

type UpgradeVersionToMultiAvailabilityZonesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Whether to support “Reading Local Nodes Only” feature after upgrading to multi-AZ deployment.
	// ul><li>`true`: The “Read Local Nodes Only” feature is supported. During the upgrade, you need to upgrade the proxy version and Redis kernel minor version simultaneously, which will involve data migration and may take hours to complete. </li><li>`false`: The “Read Local Nodes Only” feature is not supported. Upgrading to multi-AZ deployment will involve metadata migration only without affecting the service, which generally take less than three minutes to complete.</li></ul>
	UpgradeProxyAndRedisServer *bool `json:"UpgradeProxyAndRedisServer,omitempty" name:"UpgradeProxyAndRedisServer"`
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
	// Task ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// AZ name
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// Whether a product is sold out in an AZ
	IsSaleout *bool `json:"IsSaleout,omitempty" name:"IsSaleout"`

	// Whether it is a default AZ
	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`

	// Network type. basenet: basic network; vpcnet: VPC
	NetWorkType []*string `json:"NetWorkType,omitempty" name:"NetWorkType"`

	// Information of an AZ, such as product specifications in it
	ProductSet []*ProductConf `json:"ProductSet,omitempty" name:"ProductSet"`

	// AZ ID, such as 100003
	OldZoneId *int64 `json:"OldZoneId,omitempty" name:"OldZoneId"`
}