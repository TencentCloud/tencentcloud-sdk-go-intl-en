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

package v20221128

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type ApplicationVersion struct {
	// Version type
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Version ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	ApplicationVersionId *string `json:"ApplicationVersionId,omitnil,omitempty" name:"ApplicationVersionId"`

	// Release name
	// Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Release description
	// Note: This field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Entry file
	// Note: This field may return null, indicating that no valid values can be obtained.
	Entrypoint *string `json:"Entrypoint,omitnil,omitempty" name:"Entrypoint"`

	// Creation time
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Creator name
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreatorName *string `json:"CreatorName,omitnil,omitempty" name:"CreatorName"`

	// Creator ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreatorId *string `json:"CreatorId,omitnil,omitempty" name:"CreatorId"`

	// Git information
	// Note: This field may return null, indicating that no valid values can be obtained.
	GitInfo *string `json:"GitInfo,omitnil,omitempty" name:"GitInfo"`
}

type CVMOption struct {
	// CVM availability zone
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// CVM instance specifications
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`
}

type CacheInfo struct {
	// Cache cleanup time (hours)
	// Note: This field may return null, indicating that no valid values can be obtained.
	CacheClearDelay *uint64 `json:"CacheClearDelay,omitnil,omitempty" name:"CacheClearDelay"`

	// Cache cleanup schedule time
	// Note: This field may return null, indicating that no valid values can be obtained.
	CacheClearTime *string `json:"CacheClearTime,omitnil,omitempty" name:"CacheClearTime"`

	// Whether the cache has been cleaned up
	// Note: This field may return null, indicating that no valid values can be obtained.
	CacheCleared *bool `json:"CacheCleared,omitnil,omitempty" name:"CacheCleared"`
}

type ClusterOption struct {
	// Computing cluster availability zone
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Computing cluster type. Valid values:
	// - KUBERNETES
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Computing cluster Service CIDR. It must not overlap with the VPC IP range.
	ServiceCidr *string `json:"ServiceCidr,omitnil,omitempty" name:"ServiceCidr"`

	// Resource quota
	ResourceQuota *ResourceQuota `json:"ResourceQuota,omitnil,omitempty" name:"ResourceQuota"`

	// Limit scope
	LimitRange *LimitRange `json:"LimitRange,omitnil,omitempty" name:"LimitRange"`
}

// Predefined struct for user
type CreateEnvironmentRequestParams struct {
	// Environment name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Environment configuration information
	Config *EnvironmentConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// Environment description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Whether it is the default environment.
	IsDefault *bool `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`
}

type CreateEnvironmentRequest struct {
	*tchttp.BaseRequest
	
	// Environment name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Environment configuration information
	Config *EnvironmentConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// Environment description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Whether it is the default environment.
	IsDefault *bool `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`
}

func (r *CreateEnvironmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEnvironmentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Config")
	delete(f, "Description")
	delete(f, "IsDefault")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEnvironmentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEnvironmentResponseParams struct {
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// Workflow UUID
	WorkflowUuid *string `json:"WorkflowUuid,omitnil,omitempty" name:"WorkflowUuid"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateEnvironmentResponse struct {
	*tchttp.BaseResponse
	Response *CreateEnvironmentResponseParams `json:"Response"`
}

func (r *CreateEnvironmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEnvironmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVolumeRequestParams struct {
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// Name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Volume type. Valid values:
	// * SHARED: Multi-point mount shared storage
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Volume specifications. Valid values:
	// 
	// - SD: standard
	// - HP: high-performance
	// - TB: standard Turbo
	// - TP: high-performance Turbo
	Spec *string `json:"Spec,omitnil,omitempty" name:"Spec"`

	// Description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Volume size (GB), which is required to be specified for the Turbo series.
	Capacity *uint64 `json:"Capacity,omitnil,omitempty" name:"Capacity"`
}

type CreateVolumeRequest struct {
	*tchttp.BaseRequest
	
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// Name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Volume type. Valid values:
	// * SHARED: Multi-point mount shared storage
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Volume specifications. Valid values:
	// 
	// - SD: standard
	// - HP: high-performance
	// - TB: standard Turbo
	// - TP: high-performance Turbo
	Spec *string `json:"Spec,omitnil,omitempty" name:"Spec"`

	// Description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Volume size (GB), which is required to be specified for the Turbo series.
	Capacity *uint64 `json:"Capacity,omitnil,omitempty" name:"Capacity"`
}

func (r *CreateVolumeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVolumeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "Spec")
	delete(f, "Description")
	delete(f, "Capacity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVolumeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVolumeResponseParams struct {
	// Volume ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	VolumeId *string `json:"VolumeId,omitnil,omitempty" name:"VolumeId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateVolumeResponse struct {
	*tchttp.BaseResponse
	Response *CreateVolumeResponseParams `json:"Response"`
}

func (r *CreateVolumeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVolumeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DatabaseOption struct {
	// Database availability zone
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

// Predefined struct for user
type DeleteEnvironmentRequestParams struct {
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`
}

type DeleteEnvironmentRequest struct {
	*tchttp.BaseRequest
	
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`
}

func (r *DeleteEnvironmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEnvironmentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteEnvironmentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEnvironmentResponseParams struct {
	// Workflow UUID
	WorkflowUuid *string `json:"WorkflowUuid,omitnil,omitempty" name:"WorkflowUuid"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteEnvironmentResponse struct {
	*tchttp.BaseResponse
	Response *DeleteEnvironmentResponseParams `json:"Response"`
}

func (r *DeleteEnvironmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEnvironmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVolumeDataRequestParams struct {
	// Volume ID
	VolumeId *string `json:"VolumeId,omitnil,omitempty" name:"VolumeId"`

	// Path to be deleted
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`
}

type DeleteVolumeDataRequest struct {
	*tchttp.BaseRequest
	
	// Volume ID
	VolumeId *string `json:"VolumeId,omitnil,omitempty" name:"VolumeId"`

	// Path to be deleted
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`
}

func (r *DeleteVolumeDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVolumeDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VolumeId")
	delete(f, "Path")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVolumeDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVolumeDataResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteVolumeDataResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVolumeDataResponseParams `json:"Response"`
}

func (r *DeleteVolumeDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVolumeDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVolumeRequestParams struct {
	// Volume ID
	VolumeId *string `json:"VolumeId,omitnil,omitempty" name:"VolumeId"`
}

type DeleteVolumeRequest struct {
	*tchttp.BaseRequest
	
	// Volume ID
	VolumeId *string `json:"VolumeId,omitnil,omitempty" name:"VolumeId"`
}

func (r *DeleteVolumeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVolumeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VolumeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVolumeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVolumeResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteVolumeResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVolumeResponseParams `json:"Response"`
}

func (r *DeleteVolumeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVolumeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvironmentsRequestParams struct {
	// Offset, which defaults to 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Quantity of returns. It is 20 by default, and the maximum value is 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter, which supports filtering fields:
	// - EnvironmentId: Environment ID
	// - Name: Name
	// - Status: Environmental status
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeEnvironmentsRequest struct {
	*tchttp.BaseRequest
	
	// Offset, which defaults to 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Quantity of returns. It is 20 by default, and the maximum value is 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filter, which supports filtering fields:
	// - EnvironmentId: Environment ID
	// - Name: Name
	// - Status: Environmental status
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeEnvironmentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvironmentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEnvironmentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvironmentsResponseParams struct {
	// Number of qualified items
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of Environment details
	Environments []*Environment `json:"Environments,omitnil,omitempty" name:"Environments"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEnvironmentsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEnvironmentsResponseParams `json:"Response"`
}

func (r *DescribeEnvironmentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvironmentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRunGroupsRequestParams struct {
	// Project ID
	// (If you leave it blank, the default item in the specified region will be used.)
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Quantity of returns. It is 10 by default, and the maximum value is 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset, which defaults to 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Filter, which supports filtering fields:
	// - Name: Run group name
	// - RunGroupId: Run group ID
	// - Status: Run group status
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeRunGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Project ID
	// (If you leave it blank, the default item in the specified region will be used.)
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Quantity of returns. It is 10 by default, and the maximum value is 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset, which defaults to 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Filter, which supports filtering fields:
	// - Name: Run group name
	// - RunGroupId: Run group ID
	// - Status: Run group status
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeRunGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRunGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRunGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRunGroupsResponseParams struct {
	// Number of qualified items
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Run group list
	RunGroups []*RunGroup `json:"RunGroups,omitnil,omitempty" name:"RunGroups"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRunGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRunGroupsResponseParams `json:"Response"`
}

func (r *DescribeRunGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRunGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRunsRequestParams struct {
	// Project ID
	// (If you leave it blank, the default item in the specified region will be used.)
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Quantity of returns. It is 10 by default, and the maximum value is 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset, which defaults to 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Filter, which supports filtering fields:
	// - RunGroupId: run group ID
	// - Status: run status
	// - RunUuid: run UUID
	// - UserDefinedId: user-defined ID
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeRunsRequest struct {
	*tchttp.BaseRequest
	
	// Project ID
	// (If you leave it blank, the default item in the specified region will be used.)
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Quantity of returns. It is 10 by default, and the maximum value is 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset, which defaults to 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Filter, which supports filtering fields:
	// - RunGroupId: run group ID
	// - Status: run status
	// - RunUuid: run UUID
	// - UserDefinedId: user-defined ID
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeRunsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRunsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRunsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRunsResponseParams struct {
	// Number of qualified items
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Run list
	Runs []*Run `json:"Runs,omitnil,omitempty" name:"Runs"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRunsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRunsResponseParams `json:"Response"`
}

func (r *DescribeRunsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRunsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTablesRequestParams struct {
	// Project ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Quantity of returns. It is 10 by default, and the maximum value is 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset, which defaults to 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Filter, which supports filtering fields:
	// - Name: Table name
	// - TableId: Table ID
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeTablesRequest struct {
	*tchttp.BaseRequest
	
	// Project ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Quantity of returns. It is 10 by default, and the maximum value is 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset, which defaults to 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Filter, which supports filtering fields:
	// - Name: Table name
	// - TableId: Table ID
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	delete(f, "ProjectId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTablesResponseParams struct {
	// Total number of results
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Table list
	Tables []*Table `json:"Tables,omitnil,omitempty" name:"Tables"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeTablesRowsRequestParams struct {
	// Project ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Table ID
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// Quantity of returns. It is 10 by default, and the maximum value is 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset, which defaults to 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Filter, which supports filtering fields:
	// - Tr: Table data, which supports fuzzy query.
	// - TableRowUuid: table row UUID
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeTablesRowsRequest struct {
	*tchttp.BaseRequest
	
	// Project ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Table ID
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// Quantity of returns. It is 10 by default, and the maximum value is 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset, which defaults to 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Filter, which supports filtering fields:
	// - Tr: Table data, which supports fuzzy query.
	// - TableRowUuid: table row UUID
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeTablesRowsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTablesRowsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TableId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTablesRowsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTablesRowsResponseParams struct {
	// Total number of results
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Table row list
	Rows []*TableRow `json:"Rows,omitnil,omitempty" name:"Rows"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTablesRowsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTablesRowsResponseParams `json:"Response"`
}

func (r *DescribeTablesRowsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTablesRowsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVolumesRequestParams struct {
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// Quantity of returns. It is 20 by default, and the maximum value is 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset, defaults to 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Filter, supports filtering fields:
	// - Name: Name
	// - IsDefault: Whether it is the default.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeVolumesRequest struct {
	*tchttp.BaseRequest
	
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// Quantity of returns. It is 20 by default, and the maximum value is 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset, defaults to 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Filter, supports filtering fields:
	// - Name: Name
	// - IsDefault: Whether it is the default.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeVolumesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVolumesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVolumesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVolumesResponseParams struct {
	// Volume
	// Note: This field may return null, indicating that no valid values can be obtained.
	Volumes []*Volume `json:"Volumes,omitnil,omitempty" name:"Volumes"`

	// Number of qualified items
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVolumesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVolumesResponseParams `json:"Response"`
}

func (r *DescribeVolumesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVolumesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Environment struct {
	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// Environment name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Environment description information
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Environment region
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Environment type. Valid values:
	// - KUBERNETES: Kubernetes container cluster
	// - HPC:HPC HCC 
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Environment status. Valid values:
	// - INITIALIZING: Creating
	// - INITIALIZATION_ERROR: Creation failed
	// - RUNNING: Running
	// - ERROR: Exceptional
	// - DELETING: Deleting
	// - DELETE_ERROR: Deletion failed.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Whether the environment is available. The environment needs to be available before computing runs can be delivered.
	Available *bool `json:"Available,omitnil,omitempty" name:"Available"`

	// Whether the environment is the default environment.
	IsDefault *bool `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`

	// Whether the environment is a managed environment.
	IsManaged *bool `json:"IsManaged,omitnil,omitempty" name:"IsManaged"`

	// Environment information
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Cloud resource ID
	ResourceIds *ResourceIds `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// The UUID of the previous workflow
	// Note: This field may return null, indicating that no valid values can be obtained.
	LastWorkflowUuid *string `json:"LastWorkflowUuid,omitnil,omitempty" name:"LastWorkflowUuid"`

	// Creation time
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreationTime *string `json:"CreationTime,omitnil,omitempty" name:"CreationTime"`
}

type EnvironmentConfig struct {
	// VPC configuration
	VPCOption *VPCOption `json:"VPCOption,omitnil,omitempty" name:"VPCOption"`

	// Computing cluster configuration
	ClusterOption *ClusterOption `json:"ClusterOption,omitnil,omitempty" name:"ClusterOption"`

	// Database configuration
	DatabaseOption *DatabaseOption `json:"DatabaseOption,omitnil,omitempty" name:"DatabaseOption"`

	// Storage configuration
	StorageOption *StorageOption `json:"StorageOption,omitnil,omitempty" name:"StorageOption"`

	// CVM configuration
	CVMOption *CVMOption `json:"CVMOption,omitnil,omitempty" name:"CVMOption"`

	// Security group configuration
	SecurityGroupOption *SecurityGroupOption `json:"SecurityGroupOption,omitnil,omitempty" name:"SecurityGroupOption"`
}

type ExecutionTime struct {
	// Submission time
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubmitTime *string `json:"SubmitTime,omitnil,omitempty" name:"SubmitTime"`

	// Start time
	// Note: This field may return null, indicating that no valid values can be obtained.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	// Note: This field may return null, indicating that no valid values can be obtained.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type Filter struct {
	// Filtering fields
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Filtering field values
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

// Predefined struct for user
type GetRunCallsRequestParams struct {
	// Run UUID
	RunUuid *string `json:"RunUuid,omitnil,omitempty" name:"RunUuid"`

	// Job path
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// Project ID
	// (If you leave it blank, the default item in the specified region will be used.)
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type GetRunCallsRequest struct {
	*tchttp.BaseRequest
	
	// Run UUID
	RunUuid *string `json:"RunUuid,omitnil,omitempty" name:"RunUuid"`

	// Job path
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// Project ID
	// (If you leave it blank, the default item in the specified region will be used.)
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *GetRunCallsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRunCallsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RunUuid")
	delete(f, "Path")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetRunCallsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRunCallsResponseParams struct {
	// Job details
	Calls []*RunMetadata `json:"Calls,omitnil,omitempty" name:"Calls"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetRunCallsResponse struct {
	*tchttp.BaseResponse
	Response *GetRunCallsResponseParams `json:"Response"`
}

func (r *GetRunCallsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRunCallsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRunMetadataFileRequestParams struct {
	// Run UUID
	RunUuid *string `json:"RunUuid,omitnil,omitempty" name:"RunUuid"`

	// Project ID
	// (If you leave it blank, the default item in the specified region will be used.)
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// File names to be get
	// 
	// The following files are supported by default:
	// - nextflow.log
	// 
	// When report is specified as true in NFOption during submission, the following files are additionally supported:
	// - execution_report.html
	// - execution_timeline.html
	// - execution_trace.txt
	// - pipeline_dag.html
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// File names to be get in batch
	// 
	// The following files are supported by default:
	// - nextflow.log
	// 
	// When report is specified as true in NFOption during submission, the following files are additionally supported:
	// - execution_report.html
	// - execution_timeline.html
	// - execution_trace.txt
	// - pipeline_dag.html
	Keys []*string `json:"Keys,omitnil,omitempty" name:"Keys"`
}

type GetRunMetadataFileRequest struct {
	*tchttp.BaseRequest
	
	// Run UUID
	RunUuid *string `json:"RunUuid,omitnil,omitempty" name:"RunUuid"`

	// Project ID
	// (If you leave it blank, the default item in the specified region will be used.)
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// File names to be get
	// 
	// The following files are supported by default:
	// - nextflow.log
	// 
	// When report is specified as true in NFOption during submission, the following files are additionally supported:
	// - execution_report.html
	// - execution_timeline.html
	// - execution_trace.txt
	// - pipeline_dag.html
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// File names to be get in batch
	// 
	// The following files are supported by default:
	// - nextflow.log
	// 
	// When report is specified as true in NFOption during submission, the following files are additionally supported:
	// - execution_report.html
	// - execution_timeline.html
	// - execution_trace.txt
	// - pipeline_dag.html
	Keys []*string `json:"Keys,omitnil,omitempty" name:"Keys"`
}

func (r *GetRunMetadataFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRunMetadataFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RunUuid")
	delete(f, "ProjectId")
	delete(f, "Key")
	delete(f, "Keys")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetRunMetadataFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRunMetadataFileResponseParams struct {
	// Document pre-signed link that works in a minute
	CosSignedUrl *string `json:"CosSignedUrl,omitnil,omitempty" name:"CosSignedUrl"`

	// Batch document pre-signed link that works in a minute
	// Note: This field may return null, indicating that no valid values can be obtained.
	CosSignedUrls []*string `json:"CosSignedUrls,omitnil,omitempty" name:"CosSignedUrls"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetRunMetadataFileResponse struct {
	*tchttp.BaseResponse
	Response *GetRunMetadataFileResponseParams `json:"Response"`
}

func (r *GetRunMetadataFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRunMetadataFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRunStatusRequestParams struct {
	// Run UUID
	RunUuid *string `json:"RunUuid,omitnil,omitempty" name:"RunUuid"`

	// Project ID
	// (If you leave it blank, the default item in the specified region will be used.)
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type GetRunStatusRequest struct {
	*tchttp.BaseRequest
	
	// Run UUID
	RunUuid *string `json:"RunUuid,omitnil,omitempty" name:"RunUuid"`

	// Project ID
	// (If you leave it blank, the default item in the specified region will be used.)
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *GetRunStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRunStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RunUuid")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetRunStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRunStatusResponseParams struct {
	// Job details
	Metadata *RunMetadata `json:"Metadata,omitnil,omitempty" name:"Metadata"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetRunStatusResponse struct {
	*tchttp.BaseResponse
	Response *GetRunStatusResponseParams `json:"Response"`
}

func (r *GetRunStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRunStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GitInfo struct {
	// Git URL
	GitHttpPath *string `json:"GitHttpPath,omitnil,omitempty" name:"GitHttpPath"`

	// Git username .
	GitUserName *string `json:"GitUserName,omitnil,omitempty" name:"GitUserName"`

	// Git password or Token
	GitTokenOrPassword *string `json:"GitTokenOrPassword,omitnil,omitempty" name:"GitTokenOrPassword"`

	// Branch
	Branch *string `json:"Branch,omitnil,omitempty" name:"Branch"`

	// Tag
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`
}

// Predefined struct for user
type ImportTableFileRequestParams struct {
	// Project ID associated with the table
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Table name: Up to 200 characters in length is supported.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Table file COS object path
	CosUri *string `json:"CosUri,omitnil,omitempty" name:"CosUri"`

	// Data type of each column in the table file. Supported types include Int, Float, String, File, Boolean, Array[Int], Array[Float], Array[String], Array[File], and Array[Boolean].
	DataType []*string `json:"DataType,omitnil,omitempty" name:"DataType"`

	// Table description: Up to 500 characters in length is supported.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ImportTableFileRequest struct {
	*tchttp.BaseRequest
	
	// Project ID associated with the table
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Table name: Up to 200 characters in length is supported.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Table file COS object path
	CosUri *string `json:"CosUri,omitnil,omitempty" name:"CosUri"`

	// Data type of each column in the table file. Supported types include Int, Float, String, File, Boolean, Array[Int], Array[Float], Array[String], Array[File], and Array[Boolean].
	DataType []*string `json:"DataType,omitnil,omitempty" name:"DataType"`

	// Table description: Up to 500 characters in length is supported.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *ImportTableFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportTableFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "Name")
	delete(f, "CosUri")
	delete(f, "DataType")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImportTableFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportTableFileResponseParams struct {
	// Table ID
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ImportTableFileResponse struct {
	*tchttp.BaseResponse
	Response *ImportTableFileResponseParams `json:"Response"`
}

func (r *ImportTableFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportTableFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LimitRange struct {
	// Maximum CPU setting
	// Note: This field may return null, indicating that no valid values can be obtained.
	MaxCPU *string `json:"MaxCPU,omitnil,omitempty" name:"MaxCPU"`

	// Maximum memory setting (unit: Mi, Gi, Ti, M, G, and T)
	// Note: This field may return null, indicating that no valid values can be obtained.
	MaxMemory *string `json:"MaxMemory,omitnil,omitempty" name:"MaxMemory"`
}

// Predefined struct for user
type ModifyVolumeRequestParams struct {
	// Volume ID
	VolumeId *string `json:"VolumeId,omitnil,omitempty" name:"VolumeId"`

	// Name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ModifyVolumeRequest struct {
	*tchttp.BaseRequest
	
	// Volume ID
	VolumeId *string `json:"VolumeId,omitnil,omitempty" name:"VolumeId"`

	// Name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *ModifyVolumeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVolumeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VolumeId")
	delete(f, "Name")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVolumeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVolumeResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyVolumeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVolumeResponseParams `json:"Response"`
}

func (r *ModifyVolumeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVolumeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NFOption struct {
	// Config.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`

	// Profile.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Profile *string `json:"Profile,omitnil,omitempty" name:"Profile"`

	// Report.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Report *bool `json:"Report,omitnil,omitempty" name:"Report"`

	// Resume.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Resume *bool `json:"Resume,omitnil,omitempty" name:"Resume"`

	// Nextflow engine version. Valid values:
	// - 22.10.4
	// - 22.10.8 
	// - 23.10.1
	// Note: This field may return null, indicating that no valid values can be obtained.
	NFVersion *string `json:"NFVersion,omitnil,omitempty" name:"NFVersion"`
}

type ResourceIds struct {
	// VPC ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	VPCId *string `json:"VPCId,omitnil,omitempty" name:"VPCId"`

	// Subnet ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Security group ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// TDSQL-C for MySQL database ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	TDSQLCId *string `json:"TDSQLCId,omitnil,omitempty" name:"TDSQLCId"`

	//  CFS ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CFSId *string `json:"CFSId,omitnil,omitempty" name:"CFSId"`

	// CFS type. Valid values:
	// - SD: standard
	// - HP: high-performance
	// - TB: standard Turbo
	// - TP: high-performance Turbo
	// Note: This field may return null, indicating that no valid values can be obtained.
	CFSStorageType *string `json:"CFSStorageType,omitnil,omitempty" name:"CFSStorageType"`

	//  Cloud Virtual Machine ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CVMId *string `json:"CVMId,omitnil,omitempty" name:"CVMId"`

	// Elastic container cluster ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	EKSId *string `json:"EKSId,omitnil,omitempty" name:"EKSId"`
}

type ResourceQuota struct {
	// CPU limit setting
	// Note: This field may return null, indicating that no valid values can be obtained.
	CPULimit *string `json:"CPULimit,omitnil,omitempty" name:"CPULimit"`

	// Memory limit setting (Unit: Mi, Gi, Ti, M, G, and T)
	// Note: This field may return null, indicating that no valid values can be obtained.
	MemoryLimit *string `json:"MemoryLimit,omitnil,omitempty" name:"MemoryLimit"`

	// Pod quantity setting
	// Note: This field may return null, indicating that no valid values can be obtained.
	Pods *string `json:"Pods,omitnil,omitempty" name:"Pods"`
}

// Predefined struct for user
type RetryRunsRequestParams struct {
	// Project ID. (If you leave it blank, the default item in the specified region will be used.)
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// The run group ID that needs to be retried
	RunGroupId *string `json:"RunGroupId,omitnil,omitempty" name:"RunGroupId"`

	// The run UUID that needs to be retried
	RunUuids []*string `json:"RunUuids,omitnil,omitempty" name:"RunUuids"`

	// WDL running option. If you leave it blank, the retried run group running option will be used.
	WDLOption *RunOption `json:"WDLOption,omitnil,omitempty" name:"WDLOption"`

	// Nextflow running option. If you leave it blank, the retried run group running option will be used.
	NFOption *NFOption `json:"NFOption,omitnil,omitempty" name:"NFOption"`
}

type RetryRunsRequest struct {
	*tchttp.BaseRequest
	
	// Project ID. (If you leave it blank, the default item in the specified region will be used.)
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// The run group ID that needs to be retried
	RunGroupId *string `json:"RunGroupId,omitnil,omitempty" name:"RunGroupId"`

	// The run UUID that needs to be retried
	RunUuids []*string `json:"RunUuids,omitnil,omitempty" name:"RunUuids"`

	// WDL running option. If you leave it blank, the retried run group running option will be used.
	WDLOption *RunOption `json:"WDLOption,omitnil,omitempty" name:"WDLOption"`

	// Nextflow running option. If you leave it blank, the retried run group running option will be used.
	NFOption *NFOption `json:"NFOption,omitnil,omitempty" name:"NFOption"`
}

func (r *RetryRunsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetryRunsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "RunGroupId")
	delete(f, "RunUuids")
	delete(f, "WDLOption")
	delete(f, "NFOption")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RetryRunsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RetryRunsResponseParams struct {
	// New run group ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	RunGroupId *string `json:"RunGroupId,omitnil,omitempty" name:"RunGroupId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RetryRunsResponse struct {
	*tchttp.BaseResponse
	Response *RetryRunsResponseParams `json:"Response"`
}

func (r *RetryRunsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetryRunsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Run struct {
	// Run UUID
	RunUuid *string `json:"RunUuid,omitnil,omitempty" name:"RunUuid"`

	// Project ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Application ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Run group ID
	RunGroupId *string `json:"RunGroupId,omitnil,omitempty" name:"RunGroupId"`

	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// User-defined ID. Null for running in singleton mode.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserDefinedId *string `json:"UserDefinedId,omitnil,omitempty" name:"UserDefinedId"`

	// Table ID. Null for running in singleton mode.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// Table row UUID. Null for running in singleton mode.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TableRowUuid *string `json:"TableRowUuid,omitnil,omitempty" name:"TableRowUuid"`

	// Run status
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Run input
	Input *string `json:"Input,omitnil,omitempty" name:"Input"`

	// Running option
	//
	// Deprecated: Option is deprecated.
	Option *RunOption `json:"Option,omitnil,omitempty" name:"Option"`

	// Execution time
	ExecutionTime *ExecutionTime `json:"ExecutionTime,omitnil,omitempty" name:"ExecutionTime"`

	// Cache information
	// Note: This field may return null, indicating that no valid values can be obtained.
	Cache *CacheInfo `json:"Cache,omitnil,omitempty" name:"Cache"`

	// Error message
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Update time
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

// Predefined struct for user
type RunApplicationRequestParams struct {
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Run group name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Delivery environment ID
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// Project ID. (If you leave it blank, the default item in the specified region will be used.)
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Run group description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Run input COS path. (Either InputBase64 or InputCosUri must be selected.)
	InputCosUri *string `json:"InputCosUri,omitnil,omitempty" name:"InputCosUri"`

	// Run input JSON. Base64 encoding is required. (Either InputBase64 or InputCosUri must be selected.)
	InputBase64 *string `json:"InputBase64,omitnil,omitempty" name:"InputBase64"`

	// Batch deliver table ID. Leaving it blank indicates delivery in singleton mode.
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// Batch deliver table row UUID. Leaving it blank indicates all rows of the table.
	TableRowUuids []*string `json:"TableRowUuids,omitnil,omitempty" name:"TableRowUuids"`

	// Run cache cleanup time (hours). Leaving it blank or entering 0 indicates no cleanup.
	CacheClearDelay *uint64 `json:"CacheClearDelay,omitnil,omitempty" name:"CacheClearDelay"`

	// Application version ID. Leaving it blank indicates that the latest version is used.
	ApplicationVersionId *string `json:"ApplicationVersionId,omitnil,omitempty" name:"ApplicationVersionId"`

	// WDL running option
	Option *RunOption `json:"Option,omitnil,omitempty" name:"Option"`

	// Nextflow running option
	NFOption *NFOption `json:"NFOption,omitnil,omitempty" name:"NFOption"`

	// Working directory. You can fill in the absolute path in the specified volume. If you leave it blank, the default path in the default volume will be used. Currently, only Nextflow is supported.
	WorkDir *string `json:"WorkDir,omitnil,omitempty" name:"WorkDir"`

	// Access mode. Leaving it blank indicates it is private by default. Valid values:
	// - PRIVATE: Private application
	// - PUBLIC: Public application
	AccessMode *string `json:"AccessMode,omitnil,omitempty" name:"AccessMode"`

	// Volume ID. If you leave it blank, the default volume will be used. Currently, only Nextflow is supported.
	VolumeIds []*string `json:"VolumeIds,omitnil,omitempty" name:"VolumeIds"`
}

type RunApplicationRequest struct {
	*tchttp.BaseRequest
	
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Run group name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Delivery environment ID
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// Project ID. (If you leave it blank, the default item in the specified region will be used.)
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Run group description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Run input COS path. (Either InputBase64 or InputCosUri must be selected.)
	InputCosUri *string `json:"InputCosUri,omitnil,omitempty" name:"InputCosUri"`

	// Run input JSON. Base64 encoding is required. (Either InputBase64 or InputCosUri must be selected.)
	InputBase64 *string `json:"InputBase64,omitnil,omitempty" name:"InputBase64"`

	// Batch deliver table ID. Leaving it blank indicates delivery in singleton mode.
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// Batch deliver table row UUID. Leaving it blank indicates all rows of the table.
	TableRowUuids []*string `json:"TableRowUuids,omitnil,omitempty" name:"TableRowUuids"`

	// Run cache cleanup time (hours). Leaving it blank or entering 0 indicates no cleanup.
	CacheClearDelay *uint64 `json:"CacheClearDelay,omitnil,omitempty" name:"CacheClearDelay"`

	// Application version ID. Leaving it blank indicates that the latest version is used.
	ApplicationVersionId *string `json:"ApplicationVersionId,omitnil,omitempty" name:"ApplicationVersionId"`

	// WDL running option
	Option *RunOption `json:"Option,omitnil,omitempty" name:"Option"`

	// Nextflow running option
	NFOption *NFOption `json:"NFOption,omitnil,omitempty" name:"NFOption"`

	// Working directory. You can fill in the absolute path in the specified volume. If you leave it blank, the default path in the default volume will be used. Currently, only Nextflow is supported.
	WorkDir *string `json:"WorkDir,omitnil,omitempty" name:"WorkDir"`

	// Access mode. Leaving it blank indicates it is private by default. Valid values:
	// - PRIVATE: Private application
	// - PUBLIC: Public application
	AccessMode *string `json:"AccessMode,omitnil,omitempty" name:"AccessMode"`

	// Volume ID. If you leave it blank, the default volume will be used. Currently, only Nextflow is supported.
	VolumeIds []*string `json:"VolumeIds,omitnil,omitempty" name:"VolumeIds"`
}

func (r *RunApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunApplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "Name")
	delete(f, "EnvironmentId")
	delete(f, "ProjectId")
	delete(f, "Description")
	delete(f, "InputCosUri")
	delete(f, "InputBase64")
	delete(f, "TableId")
	delete(f, "TableRowUuids")
	delete(f, "CacheClearDelay")
	delete(f, "ApplicationVersionId")
	delete(f, "Option")
	delete(f, "NFOption")
	delete(f, "WorkDir")
	delete(f, "AccessMode")
	delete(f, "VolumeIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunApplicationResponseParams struct {
	// Run group ID
	RunGroupId *string `json:"RunGroupId,omitnil,omitempty" name:"RunGroupId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RunApplicationResponse struct {
	*tchttp.BaseResponse
	Response *RunApplicationResponseParams `json:"Response"`
}

func (r *RunApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunGroup struct {
	// Run group ID
	RunGroupId *string `json:"RunGroupId,omitnil,omitempty" name:"RunGroupId"`

	// Project ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Project name
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// Application ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// Application name
	ApplicationName *string `json:"ApplicationName,omitnil,omitempty" name:"ApplicationName"`

	// Application type
	ApplicationType *string `json:"ApplicationType,omitnil,omitempty" name:"ApplicationType"`

	// Environment ID
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// Environment name
	EnvironmentName *string `json:"EnvironmentName,omitnil,omitempty" name:"EnvironmentName"`

	// Table ID. Null for running in singleton mode.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// Run name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Run description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Run status
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Run input
	Input *string `json:"Input,omitnil,omitempty" name:"Input"`

	// WDL running option
	Option *RunOption `json:"Option,omitnil,omitempty" name:"Option"`

	// Nextflow running option
	// Note: This field may return null, indicating that no valid values can be obtained.
	NFOption *NFOption `json:"NFOption,omitnil,omitempty" name:"NFOption"`

	// Total number of runs
	TotalRun *uint64 `json:"TotalRun,omitnil,omitempty" name:"TotalRun"`

	// Number of runs in various status
	RunStatusCounts []*RunStatusCount `json:"RunStatusCounts,omitnil,omitempty" name:"RunStatusCounts"`

	// Execution time
	ExecutionTime *ExecutionTime `json:"ExecutionTime,omitnil,omitempty" name:"ExecutionTime"`

	// Error message
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Update time
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Creator
	// Note: This field may return null, indicating that no valid values can be obtained.
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// Creator ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreatorId *string `json:"CreatorId,omitnil,omitempty" name:"CreatorId"`

	// Running result notification method
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResultNotify *string `json:"ResultNotify,omitnil,omitempty" name:"ResultNotify"`

	// Application version
	// Note: This field may return null, indicating that no valid values can be obtained.
	ApplicationVersion *ApplicationVersion `json:"ApplicationVersion,omitnil,omitempty" name:"ApplicationVersion"`
}

type RunMetadata struct {
	// Run type
	// Note: This field may return null, indicating that no valid values can be obtained.
	RunType *string `json:"RunType,omitnil,omitempty" name:"RunType"`

	// Run ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	RunId *string `json:"RunId,omitnil,omitempty" name:"RunId"`

	// Parent layer ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	ParentId *string `json:"ParentId,omitnil,omitempty" name:"ParentId"`

	// Job ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Job name
	// Note: This field may return null, indicating that no valid values can be obtained.
	CallName *string `json:"CallName,omitnil,omitempty" name:"CallName"`

	// Scatter index
	// Note: This field may return null, indicating that no valid values can be obtained.
	ScatterIndex *string `json:"ScatterIndex,omitnil,omitempty" name:"ScatterIndex"`

	// Input
	// Note: This field may return null, indicating that no valid values can be obtained.
	Input *string `json:"Input,omitnil,omitempty" name:"Input"`

	// Output
	// Note: This field may return null, indicating that no valid values can be obtained.
	Output *string `json:"Output,omitnil,omitempty" name:"Output"`

	// Status
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Error message
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// Start time
	// Note: This field may return null, indicating that no valid values can be obtained.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Submission time
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubmitTime *string `json:"SubmitTime,omitnil,omitempty" name:"SubmitTime"`

	// End time
	// Note: This field may return null, indicating that no valid values can be obtained.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Command Line
	// Note: This field may return null, indicating that no valid values can be obtained.
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// Runtime
	// Note: This field may return null, indicating that no valid values can be obtained.
	Runtime *string `json:"Runtime,omitnil,omitempty" name:"Runtime"`

	// Preprocessing
	// Note: This field may return null, indicating that no valid values can be obtained.
	Preprocess *bool `json:"Preprocess,omitnil,omitempty" name:"Preprocess"`

	// Post-processing
	// Note: This field may return null, indicating that no valid values can be obtained.
	PostProcess *bool `json:"PostProcess,omitnil,omitempty" name:"PostProcess"`

	// Cache hit
	// Note: This field may return null, indicating that no valid values can be obtained.
	CallCached *bool `json:"CallCached,omitnil,omitempty" name:"CallCached"`

	// Standard output
	// Note: This field may return null, indicating that no valid values can be obtained.
	Stdout *string `json:"Stdout,omitnil,omitempty" name:"Stdout"`

	// Error output
	// Note: This field may return null, indicating that no valid values can be obtained.
	Stderr *string `json:"Stderr,omitnil,omitempty" name:"Stderr"`

	// Other information
	// Note: This field may return null, indicating that no valid values can be obtained.
	Meta *string `json:"Meta,omitnil,omitempty" name:"Meta"`
}

type RunOption struct {
	// Operation failure mode. Valid values:
	// - ContinueWhilePossible
	// - NoNewCalls
	FailureMode *string `json:"FailureMode,omitnil,omitempty" name:"FailureMode"`

	// Whether to use the Call-Caching feature.
	UseCallCache *bool `json:"UseCallCache,omitnil,omitempty" name:"UseCallCache"`

	// Whether to use the error suspension feature.
	UseErrorOnHold *bool `json:"UseErrorOnHold,omitnil,omitempty" name:"UseErrorOnHold"`

	// Output archive COS path
	// Note: This field may return null, indicating that no valid values can be obtained.
	FinalWorkflowOutputsDir *string `json:"FinalWorkflowOutputsDir,omitnil,omitempty" name:"FinalWorkflowOutputsDir"`

	// Whether to use the relative directory archive output.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UseRelativeOutputPaths *bool `json:"UseRelativeOutputPaths,omitnil,omitempty" name:"UseRelativeOutputPaths"`
}

type RunStatusCount struct {
	// Status
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Quantity
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`
}

// Predefined struct for user
type RunWorkflowRequestParams struct {
	// Run group name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Delivery environment ID
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// Workflow Git repository information
	GitSource *GitInfo `json:"GitSource,omitnil,omitempty" name:"GitSource"`

	// Workflow type
	// 
	// Supported type:
	// - NEXTFLOW
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Nextflow option
	NFOption *NFOption `json:"NFOption,omitnil,omitempty" name:"NFOption"`

	// Project ID
	// (If you leave it blank, the default item in the specified region will be used.)
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Run group description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Run input JSON. Base64 encoding is required.
	// (Either InputBase64 or InputCosUri must be selected.)
	InputBase64 *string `json:"InputBase64,omitnil,omitempty" name:"InputBase64"`

	// Run input COS path
	// (Either InputBase64 or InputCosUri must be selected.)
	InputCosUri *string `json:"InputCosUri,omitnil,omitempty" name:"InputCosUri"`

	// Run cache cleanup time (hours). Leaving it blank or entering 0 indicates no cleanup.
	CacheClearDelay *uint64 `json:"CacheClearDelay,omitnil,omitempty" name:"CacheClearDelay"`

	// Working directory. You can fill in the absolute path in the specified volume. If you leave it blank, the default path in the default volume will be used. Currently, only Nextflow is supported.
	WorkDir *string `json:"WorkDir,omitnil,omitempty" name:"WorkDir"`

	// Volume ID. If you leave it blank, the default volume will be used. Currently, only Nextflow is supported.
	VolumeIds []*string `json:"VolumeIds,omitnil,omitempty" name:"VolumeIds"`
}

type RunWorkflowRequest struct {
	*tchttp.BaseRequest
	
	// Run group name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Delivery environment ID
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// Workflow Git repository information
	GitSource *GitInfo `json:"GitSource,omitnil,omitempty" name:"GitSource"`

	// Workflow type
	// 
	// Supported type:
	// - NEXTFLOW
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Nextflow option
	NFOption *NFOption `json:"NFOption,omitnil,omitempty" name:"NFOption"`

	// Project ID
	// (If you leave it blank, the default item in the specified region will be used.)
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Run group description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Run input JSON. Base64 encoding is required.
	// (Either InputBase64 or InputCosUri must be selected.)
	InputBase64 *string `json:"InputBase64,omitnil,omitempty" name:"InputBase64"`

	// Run input COS path
	// (Either InputBase64 or InputCosUri must be selected.)
	InputCosUri *string `json:"InputCosUri,omitnil,omitempty" name:"InputCosUri"`

	// Run cache cleanup time (hours). Leaving it blank or entering 0 indicates no cleanup.
	CacheClearDelay *uint64 `json:"CacheClearDelay,omitnil,omitempty" name:"CacheClearDelay"`

	// Working directory. You can fill in the absolute path in the specified volume. If you leave it blank, the default path in the default volume will be used. Currently, only Nextflow is supported.
	WorkDir *string `json:"WorkDir,omitnil,omitempty" name:"WorkDir"`

	// Volume ID. If you leave it blank, the default volume will be used. Currently, only Nextflow is supported.
	VolumeIds []*string `json:"VolumeIds,omitnil,omitempty" name:"VolumeIds"`
}

func (r *RunWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunWorkflowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "EnvironmentId")
	delete(f, "GitSource")
	delete(f, "Type")
	delete(f, "NFOption")
	delete(f, "ProjectId")
	delete(f, "Description")
	delete(f, "InputBase64")
	delete(f, "InputCosUri")
	delete(f, "CacheClearDelay")
	delete(f, "WorkDir")
	delete(f, "VolumeIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunWorkflowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunWorkflowResponseParams struct {
	// Run group ID
	RunGroupId *string `json:"RunGroupId,omitnil,omitempty" name:"RunGroupId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RunWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *RunWorkflowResponseParams `json:"Response"`
}

func (r *RunWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunWorkflowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityGroupOption struct {
	// Security group ID
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`
}

type StorageOption struct {
	// CFS type. Valid values:
	// - SD: standard
	// - HP: high-performance
	// - TB: standard Turbo
	// - TP: high-performance Turbo
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// CFS availability zone
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// CFS capacity in GiB, required for the Turbo series
	// - Standard Turbo has a minimum capacity of 40 TiB, or 40,960 GiB; the capacity expansion step is 20 TiB, or 20,480 GiB.
	// - High-performance Turbo has a minimum capacity of 20 TiB, or 20,480 GiB; the capacity expansion step is 10 TiB, or 10,240 GiB.
	Capacity *uint64 `json:"Capacity,omitnil,omitempty" name:"Capacity"`
}

type Table struct {
	// Table ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// Associated project ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Table name
	// Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Table description
	// Note: This field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Table column
	// Note: This field may return null, indicating that no valid values can be obtained.
	Columns []*TableColumn `json:"Columns,omitnil,omitempty" name:"Columns"`

	// Creation time
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Creator
	// Note: This field may return null, indicating that no valid values can be obtained.
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`
}

type TableColumn struct {
	// Column name
	// Note: This field may return null, indicating that no valid values can be obtained.
	Header *string `json:"Header,omitnil,omitempty" name:"Header"`

	// Column data type
	// Note: This field may return null, indicating that no valid values can be obtained.
	DataType *string `json:"DataType,omitnil,omitempty" name:"DataType"`
}

type TableRow struct {
	// Table row UUID
	// Note: This field may return null, indicating that no valid values can be obtained.
	TableRowUuid *string `json:"TableRowUuid,omitnil,omitempty" name:"TableRowUuid"`

	// Table row content
	// Note: This field may return null, indicating that no valid values can be obtained.
	Content []*string `json:"Content,omitnil,omitempty" name:"Content"`
}

// Predefined struct for user
type TerminateRunGroupRequestParams struct {
	// Run group ID
	RunGroupId *string `json:"RunGroupId,omitnil,omitempty" name:"RunGroupId"`

	// Project ID
	// (If you leave it blank, the default item in the specified region will be used.)
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type TerminateRunGroupRequest struct {
	*tchttp.BaseRequest
	
	// Run group ID
	RunGroupId *string `json:"RunGroupId,omitnil,omitempty" name:"RunGroupId"`

	// Project ID
	// (If you leave it blank, the default item in the specified region will be used.)
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *TerminateRunGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateRunGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RunGroupId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateRunGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateRunGroupResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TerminateRunGroupResponse struct {
	*tchttp.BaseResponse
	Response *TerminateRunGroupResponseParams `json:"Response"`
}

func (r *TerminateRunGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateRunGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VPCOption struct {
	// VPC ID (Either VPCId or VPCCIDRBlock must be selected. If VPCId is selected, the existing VPCs will be used; if VPCCIDRBlock is selected, a new VPC will be created.)
	VPCId *string `json:"VPCId,omitnil,omitempty" name:"VPCId"`

	// Subnet ID (Either SubnetId or SubnetZone&SubnetCIDRBlock must be selected. If SubnetId is selected, the existing subnet will be used; if SubnetZone&SubnetCIDRBlock is selected, a new subnet will be created.)
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Subnet availability zone
	SubnetZone *string `json:"SubnetZone,omitnil,omitempty" name:"SubnetZone"`

	//  VPC CIDR.
	VPCCIDRBlock *string `json:"VPCCIDRBlock,omitnil,omitempty" name:"VPCCIDRBlock"`

	// Subnet CIDR
	SubnetCIDRBlock *string `json:"SubnetCIDRBlock,omitnil,omitempty" name:"SubnetCIDRBlock"`
}

type Volume struct {
	// Volume ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	VolumeId *string `json:"VolumeId,omitnil,omitempty" name:"VolumeId"`

	// Name
	// Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Description
	// Note: This field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Environment ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// Volume type. Valid values:
	// * SHARED: Multi-point mount shared storage
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Volume specifications. Valid values:
	// 
	// - SD: standard
	// - HP: high-performance
	// - TB: standard Turbo
	// - TP: high-performance Turbo
	// Note: This field may return null, indicating that no valid values can be obtained.
	Spec *string `json:"Spec,omitnil,omitempty" name:"Spec"`

	// Volume size (GB)
	// Note: This field may return null, indicating that no valid values can be obtained.
	Capacity *uint64 `json:"Capacity,omitnil,omitempty" name:"Capacity"`

	// Volume usage (Byte)
	// Note: This field may return null, indicating that no valid values can be obtained.
	Usage *uint64 `json:"Usage,omitnil,omitempty" name:"Usage"`

	// Volume throughput upper limit (MiB/s)
	// Note: This field may return null, indicating that no valid values can be obtained.
	BandwidthLimit *float64 `json:"BandwidthLimit,omitnil,omitempty" name:"BandwidthLimit"`

	// Default mount path
	// Note: This field may return null, indicating that no valid values can be obtained.
	DefaultMountPath *string `json:"DefaultMountPath,omitnil,omitempty" name:"DefaultMountPath"`

	// Whether it is the default volume.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsDefault *bool `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`

	// Status
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}