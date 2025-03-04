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

package v20201230

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type AccessInfo struct {
	// Address.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// Protocol.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`
}

type CBSSpec struct {
	// Disk type.
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// Size.
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// Number.
	DiskCount *int64 `json:"DiskCount,omitnil,omitempty" name:"DiskCount"`
}

type CBSSpecInfo struct {
	// Disk type.Note: This field may return null, indicating that no valid values can be obtained.
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// Size.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// Number.Note: This field may return null, indicating that no valid values can be obtained.
	DiskCount *int64 `json:"DiskCount,omitnil,omitempty" name:"DiskCount"`
}

type ChargeProperties struct {
	// 1: requires auto-renewal.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// Order time range.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// Time unit. Valid values: h and m.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// Billing type: 0 indicates pay-as-you-go and 1 indicates monthly subscription.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// PREPAID and POSTPAID_BY_HOUR
	// Note: This field may return null, indicating that no valid values can be obtained.
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`
}

// Predefined struct for user
type CreateInstanceByApiRequestParams struct {
	// Instance name.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Availability zone.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Virtual Private Cloud (VPC).
	UserVPCId *string `json:"UserVPCId,omitnil,omitempty" name:"UserVPCId"`

	// Subnet.
	UserSubnetId *string `json:"UserSubnetId,omitnil,omitempty" name:"UserSubnetId"`

	// Billing method.
	ChargeProperties *ChargeProperties `json:"ChargeProperties,omitnil,omitempty" name:"ChargeProperties"`

	// Instance password.
	AdminPassword *string `json:"AdminPassword,omitnil,omitempty" name:"AdminPassword"`

	// Resource information.
	Resources []*ResourceSpecNew `json:"Resources,omitnil,omitempty" name:"Resources"`

	// Tag list.
	Tags *Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Version.
	ProductVersion *string `json:"ProductVersion,omitnil,omitempty" name:"ProductVersion"`
}

type CreateInstanceByApiRequest struct {
	*tchttp.BaseRequest
	
	// Instance name.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Availability zone.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Virtual Private Cloud (VPC).
	UserVPCId *string `json:"UserVPCId,omitnil,omitempty" name:"UserVPCId"`

	// Subnet.
	UserSubnetId *string `json:"UserSubnetId,omitnil,omitempty" name:"UserSubnetId"`

	// Billing method.
	ChargeProperties *ChargeProperties `json:"ChargeProperties,omitnil,omitempty" name:"ChargeProperties"`

	// Instance password.
	AdminPassword *string `json:"AdminPassword,omitnil,omitempty" name:"AdminPassword"`

	// Resource information.
	Resources []*ResourceSpecNew `json:"Resources,omitnil,omitempty" name:"Resources"`

	// Tag list.
	Tags *Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Version.
	ProductVersion *string `json:"ProductVersion,omitnil,omitempty" name:"ProductVersion"`
}

func (r *CreateInstanceByApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceByApiRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceName")
	delete(f, "Zone")
	delete(f, "UserVPCId")
	delete(f, "UserSubnetId")
	delete(f, "ChargeProperties")
	delete(f, "AdminPassword")
	delete(f, "Resources")
	delete(f, "Tags")
	delete(f, "ProductVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstanceByApiRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceByApiResponseParams struct {
	// Process ID.Note: This field may return null, indicating that no valid values can be obtained.
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Instance ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Error message.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateInstanceByApiResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstanceByApiResponseParams `json:"Response"`
}

func (r *CreateInstanceByApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceByApiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceInfoRequestParams struct {
	// instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceInfoRequest struct {
	*tchttp.BaseRequest
	
	// instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceInfoResponseParams struct {
	// Instance simple information
	SimpleInstanceInfo *SimpleInstanceInfo `json:"SimpleInstanceInfo,omitnil,omitempty" name:"SimpleInstanceInfo"`

	// Error Message
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceInfoResponseParams `json:"Response"`
}

func (r *DescribeInstanceInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceRequestParams struct {
	// instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceRequest struct {
	*tchttp.BaseRequest
	
	// instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceResponseParams struct {
	// Instance description information.
	InstanceInfo *InstanceInfo `json:"InstanceInfo,omitnil,omitempty" name:"InstanceInfo"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceResponseParams `json:"Response"`
}

func (r *DescribeInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceStateRequestParams struct {
	//  Instance id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceStateRequest struct {
	*tchttp.BaseRequest
	
	//  Instance id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceStateResponseParams struct {
	// Instance status. Example: serving.
	InstanceState *string `json:"InstanceState,omitnil,omitempty" name:"InstanceState"`

	// Creation time of instance operation.Note: This field may return null, indicating that no valid values can be obtained.
	FlowCreateTime *string `json:"FlowCreateTime,omitnil,omitempty" name:"FlowCreateTime"`

	// Instance operation name.Note: This field may return null, indicating that no valid values can be obtained.
	FlowName *string `json:"FlowName,omitnil,omitempty" name:"FlowName"`

	// Instance operation progress.Note: This field may return null, indicating that no valid values can be obtained.
	FlowProgress *float64 `json:"FlowProgress,omitnil,omitempty" name:"FlowProgress"`

	// Cluster status description. Example: running.Note: This field may return null, indicating that no valid values can be obtained.
	InstanceStateDesc *string `json:"InstanceStateDesc,omitnil,omitempty" name:"InstanceStateDesc"`

	// Instance process error messages. Example: "Creation failed, insufficient resources."
	// Note: This field may return null, indicating that no valid values can be obtained.
	FlowMsg *string `json:"FlowMsg,omitnil,omitempty" name:"FlowMsg"`

	// The name of the current step. Example: "Purchasing resources."Note: This field may return null, indicating that no valid values can be obtained.
	ProcessName *string `json:"ProcessName,omitnil,omitempty" name:"ProcessName"`

	// Enabling status of the instance backup task.Note: This field may return null, indicating that no valid values can be obtained.
	BackupStatus *int64 `json:"BackupStatus,omitnil,omitempty" name:"BackupStatus"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceStateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceStateResponseParams `json:"Response"`
}

func (r *DescribeInstanceStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesRequestParams struct {
	// The name of the Instance ID for the search.
	SearchInstanceId *string `json:"SearchInstanceId,omitnil,omitempty" name:"SearchInstanceId"`

	// The instance name for the search.
	SearchInstanceName *string `json:"SearchInstanceName,omitnil,omitempty" name:"SearchInstanceName"`

	// Pagination parameter. The first page is 0, and the second page is 10.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Pagination parameter. Pagination step length. It is 10 by default.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Searched tag list.
	SearchTags []*SearchTags `json:"SearchTags,omitnil,omitempty" name:"SearchTags"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// The name of the Instance ID for the search.
	SearchInstanceId *string `json:"SearchInstanceId,omitnil,omitempty" name:"SearchInstanceId"`

	// The instance name for the search.
	SearchInstanceName *string `json:"SearchInstanceName,omitnil,omitempty" name:"SearchInstanceName"`

	// Pagination parameter. The first page is 0, and the second page is 10.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Pagination parameter. Pagination step length. It is 10 by default.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Searched tag list.
	SearchTags []*SearchTags `json:"SearchTags,omitnil,omitempty" name:"SearchTags"`
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
	delete(f, "SearchInstanceId")
	delete(f, "SearchInstanceName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchTags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// Total number of instances.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Instance array.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstancesList []*InstanceInfo `json:"InstancesList,omitnil,omitempty" name:"InstancesList"`

	//  -Note: This field may return null, indicating that no valid values can be obtained.
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

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
type DescribeSimpleInstancesRequestParams struct {
	// Search Instance Id
	SearchInstanceId *string `json:"SearchInstanceId,omitnil,omitempty" name:"SearchInstanceId"`

	// Search Instance Name
	SearchInstanceName *string `json:"SearchInstanceName,omitnil,omitempty" name:"SearchInstanceName"`

	// 	
	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Search Tags
	SearchTags []*string `json:"SearchTags,omitnil,omitempty" name:"SearchTags"`
}

type DescribeSimpleInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Search Instance Id
	SearchInstanceId *string `json:"SearchInstanceId,omitnil,omitempty" name:"SearchInstanceId"`

	// Search Instance Name
	SearchInstanceName *string `json:"SearchInstanceName,omitnil,omitempty" name:"SearchInstanceName"`

	// 	
	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Search Tags
	SearchTags []*string `json:"SearchTags,omitnil,omitempty" name:"SearchTags"`
}

func (r *DescribeSimpleInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSimpleInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SearchInstanceId")
	delete(f, "SearchInstanceName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchTags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSimpleInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSimpleInstancesResponseParams struct {
	// Total Count
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Instances List
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstancesList []*InstanceSimpleInfoNew `json:"InstancesList,omitnil,omitempty" name:"InstancesList"`

	//  Error Message
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSimpleInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSimpleInstancesResponseParams `json:"Response"`
}

func (r *DescribeSimpleInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSimpleInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyInstanceByApiRequestParams struct {
	// Instance id. Example: "cdwpg-xxxx".
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DestroyInstanceByApiRequest struct {
	*tchttp.BaseRequest
	
	// Instance id. Example: "cdwpg-xxxx".
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DestroyInstanceByApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyInstanceByApiRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyInstanceByApiRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyInstanceByApiResponseParams struct {
	// Destroy  process ID.
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Error message.
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DestroyInstanceByApiResponse struct {
	*tchttp.BaseResponse
	Response *DestroyInstanceByApiResponseParams `json:"Response"`
}

func (r *DestroyInstanceByApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyInstanceByApiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiskSpecPlus struct {
	// Disk Count
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskCount *int64 `json:"DiskCount,omitnil,omitempty" name:"DiskCount"`

	// Max Disk Size
	// Note: This field may return null, indicating that no valid values can be obtained.
	MaxDiskSize *int64 `json:"MaxDiskSize,omitnil,omitempty" name:"MaxDiskSize"`

	// Min Disk Size
	// Note: This field may return null, indicating that no valid values can be obtained.
	MinDiskSize *int64 `json:"MinDiskSize,omitnil,omitempty" name:"MinDiskSize"`

	// Disk Type
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// Disk Description
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskDesc *string `json:"DiskDesc,omitnil,omitempty" name:"DiskDesc"`

	// Cvm Class
	// Note: This field may return null, indicating that no valid values can be obtained.
	CvmClass *string `json:"CvmClass,omitnil,omitempty" name:"CvmClass"`
}

type InstanceInfo struct {
	// Instance ID 
	// Note: This field may return null, indicating that no valid values can be obtained.
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// cn/dn or others.Note: This field may return null, indicating that no valid values can be obtained.
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Instance Name
	// .Note: This field may return null, indicating that no valid values can be obtained.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Instance Status ,such as  2
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Instance Status  Description ,such as Running.
	// Note: This field may return null, indicating that no valid values can be obtained.
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// Instance State Infomation
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceStateInfo *InstanceStateInfo `json:"InstanceStateInfo,omitnil,omitempty" name:"InstanceStateInfo"`

	// Instance ID 
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// CreateTime ,such as 2022-09-05 20:00:01
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Region ,such as ap-chongqing.Note: This field may return null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Zone : This field may return null, indicating that no valid values can be obtained.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Region.Note: This field may return null, indicating that no valid values can be obtained.
	RegionDesc *string `json:"RegionDesc,omitnil,omitempty" name:"RegionDesc"`

	// Zone.Note: This field may return null, indicating that no valid values can be obtained.
	ZoneDesc *string `json:"ZoneDesc,omitnil,omitempty" name:"ZoneDesc"`

	// Tag.Note: This field may return null, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Version.Note: This field may return null, indicating that no valid values can be obtained.
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// Character set.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Charset *string `json:"Charset,omitnil,omitempty" name:"Charset"`

	// CN node list.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CNNodes []*InstanceNodeGroup `json:"CNNodes,omitnil,omitempty" name:"CNNodes"`

	// DN node list.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DNNodes []*InstanceNodeGroup `json:"DNNodes,omitnil,omitempty" name:"DNNodes"`

	//  Region Id
	// Note: This field may return null, indicating that no valid values can be obtained.
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Zone Id
	// Note: This field may return null, indicating that no valid values can be obtained.
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Vpc Id
	// Note: This field may return null, indicating that no valid values can be obtained.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet Id
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Expire Time
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// Pay Mode
	// Note: This field may return null, indicating that no valid values can be obtained.
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Renew Flag
	// Note: This field may return null, indicating that no valid values can be obtained.
	RenewFlag *bool `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// Instance Id
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Access information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AccessDetails []*AccessInfo `json:"AccessDetails,omitnil,omitempty" name:"AccessDetails"`
}

type InstanceNodeGroup struct {
	// Spec Name
	// Note: This field may return null, indicating that no valid values can be obtained.
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// Data Disk
	// Note: This field may return null, indicating that no valid values can be obtained.
	DataDisk *DiskSpecPlus `json:"DataDisk,omitnil,omitempty" name:"DataDisk"`

	// Cvm Count
	// Note: This field may return null, indicating that no valid values can be obtained.
	CvmCount *int64 `json:"CvmCount,omitnil,omitempty" name:"CvmCount"`
}

type InstanceSimpleInfoNew struct {
	// ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// Instance Id
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance Name
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Version
	// Note: This field may return null, indicating that no valid values can be obtained.
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// Region
	// Note: This field may return null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Region Id
	// Note: This field may return null, indicating that no valid values can be obtained.
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Region Desc
	// Note: This field may return null, indicating that no valid values can be obtained.
	RegionDesc *string `json:"RegionDesc,omitnil,omitempty" name:"RegionDesc"`

	// Zone
	// Note: This field may return null, indicating that no valid values can be obtained.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Zone Id
	// Note: This field may return null, indicating that no valid values can be obtained.
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Zone Desc
	// Note: This field may return null, indicating that no valid values can be obtained.
	ZoneDesc *string `json:"ZoneDesc,omitnil,omitempty" name:"ZoneDesc"`

	// Vpc Id
	// Note: This field may return null, indicating that no valid values can be obtained.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet Id
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// CreateTime
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// ExpireTime
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// Access Infomation
	// Note: This field may return null, indicating that no valid values can be obtained.
	AccessInfo *string `json:"AccessInfo,omitnil,omitempty" name:"AccessInfo"`

	// Pay Mode
	// Note: This field may return null, indicating that no valid values can be obtained.
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Renew Flag
	// Note: This field may return null, indicating that no valid values can be obtained.
	RenewFlag *bool `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

type InstanceStateInfo struct {
	// Instance status. Example: serving.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceState *string `json:"InstanceState,omitnil,omitempty" name:"InstanceState"`

	// Creation time of instance operation.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FlowCreateTime *string `json:"FlowCreateTime,omitnil,omitempty" name:"FlowCreateTime"`

	// Instance operation name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FlowName *string `json:"FlowName,omitnil,omitempty" name:"FlowName"`

	// Instance operation progress.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FlowProgress *int64 `json:"FlowProgress,omitnil,omitempty" name:"FlowProgress"`

	// Instance status description. Example: running.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceStateDesc *string `json:"InstanceStateDesc,omitnil,omitempty" name:"InstanceStateDesc"`

	// Instance process error messages. Example: "Creation failed, insufficient resources."
	// Note: This field may return null, indicating that no valid values can be obtained.
	FlowMsg *string `json:"FlowMsg,omitnil,omitempty" name:"FlowMsg"`

	// The name of the current step. Example: "Purchasing resources."
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProcessName *string `json:"ProcessName,omitnil,omitempty" name:"ProcessName"`

	// Indicates whether there is a backup task in the instance. 1 indicates yes and 0 indicates no.
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupStatus *int64 `json:"BackupStatus,omitnil,omitempty" name:"BackupStatus"`

	// Request Id
	// Note: This field may return null, indicating that no valid values can be obtained.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`

	// Backup Open Status
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupOpenStatus *int64 `json:"BackupOpenStatus,omitnil,omitempty" name:"BackupOpenStatus"`
}

// Predefined struct for user
type ModifyInstanceRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Name of the newly modified instance.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

type ModifyInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Name of the newly modified instance.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
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
	delete(f, "InstanceId")
	delete(f, "InstanceName")
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

type ResourceInfo struct {
	// Resource name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// Resource count.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// Disk information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskSpec *CBSSpecInfo `json:"DiskSpec,omitnil,omitempty" name:"DiskSpec"`

	// Resource type.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type ResourceSpecNew struct {
	// Resource name.
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// Resource count.
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// Disk information.
	DiskSpec *CBSSpec `json:"DiskSpec,omitnil,omitempty" name:"DiskSpec"`

	// Resource type, DATA.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type SearchTags struct {
	// Tag key.
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag value.
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`

	// 1 means only the Tag key is entered without a value, and 0 means both the key and the value are entered.
	AllValue *int64 `json:"AllValue,omitnil,omitempty" name:"AllValue"`
}

type SimpleInstanceInfo struct {
	// ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// Instance Id
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance Name
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Version
	// Note: This field may return null, indicating that no valid values can be obtained.
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// Region
	// Note: This field may return null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Zone
	// Note: This field may return null, indicating that no valid values can be obtained.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// User VPC ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserVPCID *string `json:"UserVPCID,omitnil,omitempty" name:"UserVPCID"`

	// User Subnet ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserSubnetID *string `json:"UserSubnetID,omitnil,omitempty" name:"UserSubnetID"`

	// CreateTime
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// ExpireTime
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// Access Infomation
	// Note: This field may return null, indicating that no valid values can be obtained.
	AccessInfo *string `json:"AccessInfo,omitnil,omitempty" name:"AccessInfo"`

	// Renew Flag
	// Note: This field may return null, indicating that no valid values can be obtained.
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// Charge Properties
	// Note: This field may return null, indicating that no valid values can be obtained.
	ChargeProperties *ChargeProperties `json:"ChargeProperties,omitnil,omitempty" name:"ChargeProperties"`

	// Resources
	// Note: This field may return null, indicating that no valid values can be obtained.
	Resources []*ResourceInfo `json:"Resources,omitnil,omitempty" name:"Resources"`

	// Tags
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Status
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type Tag struct {
	// Tag key.
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag value.
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}