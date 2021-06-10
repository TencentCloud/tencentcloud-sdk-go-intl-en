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

package v20190319

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type ConfigurationItems struct {

	// Time of getting a configuration item
	ConfigurationItemCaptureTime *string `json:"ConfigurationItemCaptureTime,omitempty" name:"ConfigurationItemCaptureTime"`

	// Resource relationship list
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Relationships *string `json:"Relationships,omitempty" name:"Relationships"`

	// This parameter takes effect only when `DiffMode` is set to `true`. When the input parameter `ChronologicalOrder` of the `GetConfigurationItems` API is set to `Forward`, details of the configuration item before the first one (if not a creation configuration item) will be returned. When this parameter is set to `Reverse`, details of the configuration item after the last one (if not a resource deletion configuration item) will be returned.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	LastItemInfo *string `json:"LastItemInfo,omitempty" name:"LastItemInfo"`

	// List of events associated with the configuration changes
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	RelatedEvents []*RelatedEvent `json:"RelatedEvents,omitempty" name:"RelatedEvents"`

	// Resource type
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// Resource ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// Configuration item ID
	ConfigurationStateId *string `json:"ConfigurationStateId,omitempty" name:"ConfigurationStateId"`

	// Resource creation time
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ResourceCreateTime *string `json:"ResourceCreateTime,omitempty" name:"ResourceCreateTime"`

	// CFA version
	Version *string `json:"Version,omitempty" name:"Version"`

	// Resource region
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 
	Configuration *string `json:"Configuration,omitempty" name:"Configuration"`

	// Resource name
	ResourceAlias *string `json:"ResourceAlias,omitempty" name:"ResourceAlias"`

	// Configuration item status. Valid values: OK, ResourceDiscovered, ResourceNotRecorded, ResourceDeleted, ResourceDeletedNotRecorded.
	ConfigurationItemStatus *string `json:"ConfigurationItemStatus,omitempty" name:"ConfigurationItemStatus"`
}

type CreateRecorderRequest struct {
	*tchttp.BaseRequest

	// Role name authorized to CFA
	Role *string `json:"Role,omitempty" name:"Role"`

	// Whether to select all supported resource types. Valid values: true (default), false.
	AllSupported *bool `json:"AllSupported,omitempty" name:"AllSupported"`

	// Whether to enable the resource recorder. Valid values: true (default), false.
	Enable *bool `json:"Enable,omitempty" name:"Enable"`

	// Resource recorder name. Default name: default.
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *CreateRecorderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRecorderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Role")
	delete(f, "AllSupported")
	delete(f, "Enable")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRecorderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateRecorderResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Whether the recorder was created successfully
		IsSuccess *bool `json:"IsSuccess,omitempty" name:"IsSuccess"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRecorderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRecorderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRecorderRequest struct {
	*tchttp.BaseRequest
}

func (r *DeleteRecorderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecorderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRecorderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRecorderResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Whether the recorder was deleted successfully
		IsSuccess *bool `json:"IsSuccess,omitempty" name:"IsSuccess"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRecorderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecorderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDiscoveredResourceRequest struct {
	*tchttp.BaseRequest

	// Request ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

func (r *DescribeDiscoveredResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiscoveredResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDiscoveredResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDiscoveredResourceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Last update time
		LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`

		// Resource type
		ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

		// Resource ID
		ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

		// Resource creation time
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// Tag details
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		Tag *string `json:"Tag,omitempty" name:"Tag"`

		// Current resource configuration details
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		ResourceInfo *string `json:"ResourceInfo,omitempty" name:"ResourceInfo"`

		// Resource region
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

		// Resource alias
		ResourceAlias *string `json:"ResourceAlias,omitempty" name:"ResourceAlias"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDiscoveredResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiscoveredResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecorderRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRecorderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecorderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecorderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecorderResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Whether to enable the recorder. Valid values: true (enable), false (disable).
		Enable *bool `json:"Enable,omitempty" name:"Enable"`

		// Recorder name
		Name *string `json:"Name,omitempty" name:"Name"`

		// Last error message of the recorder, which corresponds to `LastErrorCode`.
		LastErrorMessage *string `json:"LastErrorMessage,omitempty" name:"LastErrorMessage"`

		// The status of the recorder when it recorded information last time. Valid values: PENDING, OK, FAILED.
		LastStatus *string `json:"LastStatus,omitempty" name:"LastStatus"`

		// List of the resource types monitored by the recorder
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		ResourceTypes []*RecordResourceType `json:"ResourceTypes,omitempty" name:"ResourceTypes"`

		// Time when the recorder was enabled last time
		LastStartTime *string `json:"LastStartTime,omitempty" name:"LastStartTime"`

		// Last error code of the recorder
		LastErrorCode *string `json:"LastErrorCode,omitempty" name:"LastErrorCode"`

		// Time when the recorder was disabled last time
		LastStopTime *string `json:"LastStopTime,omitempty" name:"LastStopTime"`

		// Whether to monitor all currently supported resource types. Valid values: true (yes), false (no).
		AllSupported *bool `json:"AllSupported,omitempty" name:"AllSupported"`

		// Recorder creation time
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// Role name authorized to CFA
		Role *string `json:"Role,omitempty" name:"Role"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRecorderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecorderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetConfigurationItemsRequest struct {
	*tchttp.BaseRequest

	// Resource ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// Chronological order. Valid values: Reverse, Forward (default).
	ChronologicalOrder *string `json:"ChronologicalOrder,omitempty" name:"ChronologicalOrder"`

	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Whether to enable `DiffMode`. Valid values: true, false (default).
	DiffMode *bool `json:"DiffMode,omitempty" name:"DiffMode"`

	// Returned number. default: 10, maximum: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *GetConfigurationItemsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetConfigurationItemsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceId")
	delete(f, "ChronologicalOrder")
	delete(f, "StartTime")
	delete(f, "Offset")
	delete(f, "DiffMode")
	delete(f, "Limit")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetConfigurationItemsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetConfigurationItemsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Resource configuration item list
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		ConfigurationItems []*ConfigurationItems `json:"ConfigurationItems,omitempty" name:"ConfigurationItems"`

		// Total number
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetConfigurationItemsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetConfigurationItemsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListDiscoveredResourcesRequest struct {
	*tchttp.BaseRequest

	// Resource type
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// Resource ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// Returned number. default: 20, maximum: 200.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Resource region
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// Offset. Default: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Whether the resource is deleted
	IsDeleted *bool `json:"IsDeleted,omitempty" name:"IsDeleted"`
}

func (r *ListDiscoveredResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDiscoveredResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceType")
	delete(f, "ResourceId")
	delete(f, "Limit")
	delete(f, "ResourceRegion")
	delete(f, "Offset")
	delete(f, "IsDeleted")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListDiscoveredResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListDiscoveredResourcesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Resource list
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		Resources []*Resources `json:"Resources,omitempty" name:"Resources"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListDiscoveredResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDiscoveredResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSupportResourceTypesRequest struct {
	*tchttp.BaseRequest
}

func (r *ListSupportResourceTypesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSupportResourceTypesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListSupportResourceTypesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListSupportResourceTypesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of supported resource types
		ResourceTypes []*SupportResourceType `json:"ResourceTypes,omitempty" name:"ResourceTypes"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListSupportResourceTypesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSupportResourceTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecordResourceType struct {

	// CAM policy name
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// Modification time of resource types for monitoring
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Service
	Service *string `json:"Service,omitempty" name:"Service"`

	// Resource type
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// Service name
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// Resource type name
	ResourceTypeName *string `json:"ResourceTypeName,omitempty" name:"ResourceTypeName"`
}

type RelatedEvent struct {

	// Event name
	EventName *string `json:"EventName,omitempty" name:"EventName"`

	// Operation time
	EventTime *string `json:"EventTime,omitempty" name:"EventTime"`

	// ID of the operator account
	OperateUin *uint64 `json:"OperateUin,omitempty" name:"OperateUin"`

	// CloudAudit event ID
	EventReqId *string `json:"EventReqId,omitempty" name:"EventReqId"`
}

type Resources struct {

	// Resource type
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// Resource ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// Resource creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Resource region
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// Resource alias
	ResourceAlias *string `json:"ResourceAlias,omitempty" name:"ResourceAlias"`

	// Whether the resource is deleted
	IsDeleted *bool `json:"IsDeleted,omitempty" name:"IsDeleted"`
}

type SupportResourceType struct {

	// Resource type
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// CAM policy name
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// Service name
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// Resource type name in Chinese
	ResourceTypeName *string `json:"ResourceTypeName,omitempty" name:"ResourceTypeName"`

	// Service
	Service *string `json:"Service,omitempty" name:"Service"`
}

type UpdateRecorderRequest struct {
	*tchttp.BaseRequest

	// Whether to select all currently supported resource types
	AllSupported *bool `json:"AllSupported,omitempty" name:"AllSupported"`

	// Whether to enable the recorder. Valid values: true (enable), false (disable).
	Enable *bool `json:"Enable,omitempty" name:"Enable"`

	// Recorder name after modification
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *UpdateRecorderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRecorderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AllSupported")
	delete(f, "Enable")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateRecorderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpdateRecorderResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Whether the modification is successful
		IsSuccess *bool `json:"IsSuccess,omitempty" name:"IsSuccess"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateRecorderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRecorderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
