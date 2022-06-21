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

package v20220106

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

// Predefined struct for user
type CreatePrefetchTaskRequestParams struct {
	// ID of the site
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// List of resources to be pre-warmed, for example:
	// http://www.example.com/example.txt
	Targets []*string `json:"Targets,omitempty" name:"Targets"`

	// Specifies whether to encode the URL
	// Note that if it’s enabled, the purging is based on the converted URLs.
	EncodeUrl *bool `json:"EncodeUrl,omitempty" name:"EncodeUrl"`

	// HTTP header information
	Headers []*Header `json:"Headers,omitempty" name:"Headers"`
}

type CreatePrefetchTaskRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// List of resources to be pre-warmed, for example:
	// http://www.example.com/example.txt
	Targets []*string `json:"Targets,omitempty" name:"Targets"`

	// Specifies whether to encode the URL
	// Note that if it’s enabled, the purging is based on the converted URLs.
	EncodeUrl *bool `json:"EncodeUrl,omitempty" name:"EncodeUrl"`

	// HTTP header information
	Headers []*Header `json:"Headers,omitempty" name:"Headers"`
}

func (r *CreatePrefetchTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrefetchTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Targets")
	delete(f, "EncodeUrl")
	delete(f, "Headers")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrefetchTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrefetchTaskResponseParams struct {
	// Task ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// List of failed tasks
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	FailedList []*FailReason `json:"FailedList,omitempty" name:"FailedList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePrefetchTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrefetchTaskResponseParams `json:"Response"`
}

func (r *CreatePrefetchTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrefetchTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePurgeTaskRequestParams struct {
	// ID of the site
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Type of the purging task. Values:
	// - `purge_url`: Purge by the URL
	// - `purge_prefix`: Purge by the prefix
	// - `purge_host`: Purge by the Hostname
	// - `purge_all`: Purge all cached contents
	Type *string `json:"Type,omitempty" name:"Type"`

	// Target resource to be purged, which depends on the `Type` field.
	// 1. When `Type = purge_host`:
	// Hostnames are purged, such as www.example.com and foo.bar.example.com.
	// 2. When `Type = purge_prefix`:
	// Prefixes are purged, such as http://www.example.com/example.
	// 3. When `Type = purge_url`:
	// URLs are purged, such as https://www.example.com/example.jpg.
	// 4. When `Type = purge_all`: All types of resources are purged.
	// `Targets` is not a required field.
	Targets []*string `json:"Targets,omitempty" name:"Targets"`

	// Specifies whether to transcode non-ASCII URLs according to RFC3986.
	// Note that if it’s enabled, the purging is based on the converted URLs.
	EncodeUrl *bool `json:"EncodeUrl,omitempty" name:"EncodeUrl"`
}

type CreatePurgeTaskRequest struct {
	*tchttp.BaseRequest
	
	// ID of the site
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Type of the purging task. Values:
	// - `purge_url`: Purge by the URL
	// - `purge_prefix`: Purge by the prefix
	// - `purge_host`: Purge by the Hostname
	// - `purge_all`: Purge all cached contents
	Type *string `json:"Type,omitempty" name:"Type"`

	// Target resource to be purged, which depends on the `Type` field.
	// 1. When `Type = purge_host`:
	// Hostnames are purged, such as www.example.com and foo.bar.example.com.
	// 2. When `Type = purge_prefix`:
	// Prefixes are purged, such as http://www.example.com/example.
	// 3. When `Type = purge_url`:
	// URLs are purged, such as https://www.example.com/example.jpg.
	// 4. When `Type = purge_all`: All types of resources are purged.
	// `Targets` is not a required field.
	Targets []*string `json:"Targets,omitempty" name:"Targets"`

	// Specifies whether to transcode non-ASCII URLs according to RFC3986.
	// Note that if it’s enabled, the purging is based on the converted URLs.
	EncodeUrl *bool `json:"EncodeUrl,omitempty" name:"EncodeUrl"`
}

func (r *CreatePurgeTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePurgeTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Type")
	delete(f, "Targets")
	delete(f, "EncodeUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePurgeTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePurgeTaskResponseParams struct {
	// Task ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// List of failed tasks and reasons
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	FailedList []*FailReason `json:"FailedList,omitempty" name:"FailedList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePurgeTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreatePurgeTaskResponseParams `json:"Response"`
}

func (r *CreatePurgeTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePurgeTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrefetchTasksRequestParams struct {
	// Task ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// Start time of the query
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of the query
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Offset of the query
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of results returned
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Statuses of tasks to be queried. Values:
	// `processing`, `success`, `failed`, `timeout` and `invalid`
	Statuses []*string `json:"Statuses,omitempty" name:"Statuses"`

	// ID of the site
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// List of domain names queried
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Resources queried
	Target *string `json:"Target,omitempty" name:"Target"`
}

type DescribePrefetchTasksRequest struct {
	*tchttp.BaseRequest
	
	// Task ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// Start time of the query
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of the query
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Offset of the query
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of results returned
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Statuses of tasks to be queried. Values:
	// `processing`, `success`, `failed`, `timeout` and `invalid`
	Statuses []*string `json:"Statuses,omitempty" name:"Statuses"`

	// ID of the site
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// List of domain names queried
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Resources queried
	Target *string `json:"Target,omitempty" name:"Target"`
}

func (r *DescribePrefetchTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrefetchTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Statuses")
	delete(f, "ZoneId")
	delete(f, "Domains")
	delete(f, "Target")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrefetchTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrefetchTasksResponseParams struct {
	// Total entries that match the specified query condition
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of tasks returned
	Tasks []*Task `json:"Tasks,omitempty" name:"Tasks"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePrefetchTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrefetchTasksResponseParams `json:"Response"`
}

func (r *DescribePrefetchTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrefetchTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePurgeTasksRequestParams struct {
	// Task ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// Type of the purging task
	Type *string `json:"Type,omitempty" name:"Type"`

	// Start time of the query
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of the query
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Offset of the query
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of results returned
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Statuses of tasks to be queried. Values:
	// `processing`, `success`, `failed`, `timeout` and `invalid`
	Statuses []*string `json:"Statuses,omitempty" name:"Statuses"`

	// ID of the site
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// List of domain names queried
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Queries content
	Target *string `json:"Target,omitempty" name:"Target"`
}

type DescribePurgeTasksRequest struct {
	*tchttp.BaseRequest
	
	// Task ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// Type of the purging task
	Type *string `json:"Type,omitempty" name:"Type"`

	// Start time of the query
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of the query
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Offset of the query
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of results returned
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Statuses of tasks to be queried. Values:
	// `processing`, `success`, `failed`, `timeout` and `invalid`
	Statuses []*string `json:"Statuses,omitempty" name:"Statuses"`

	// ID of the site
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// List of domain names queried
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Queries content
	Target *string `json:"Target,omitempty" name:"Target"`
}

func (r *DescribePurgeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePurgeTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "Type")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Statuses")
	delete(f, "ZoneId")
	delete(f, "Domains")
	delete(f, "Target")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePurgeTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePurgeTasksResponseParams struct {
	// Total entries that match the specified query condition
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of tasks returned
	Tasks []*Task `json:"Tasks,omitempty" name:"Tasks"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePurgeTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribePurgeTasksResponseParams `json:"Response"`
}

func (r *DescribePurgeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePurgeTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZonesRequestParams struct {
	// Pagination parameter, which specifies the offset.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Pagination parameter, which specifies the number of sites returned in each page.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Query condition filter, which supports complex type.
	Filters []*ZoneFilter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeZonesRequest struct {
	*tchttp.BaseRequest
	
	// Pagination parameter, which specifies the offset.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Pagination parameter, which specifies the number of sites returned in each page.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Query condition filter, which supports complex type.
	Filters []*ZoneFilter `json:"Filters,omitempty" name:"Filters"`
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
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZonesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZonesResponseParams struct {
	// Number of sites that match the specified conditions
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Details of sites
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Zones []*Zone `json:"Zones,omitempty" name:"Zones"`

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
type DownloadL7LogsRequestParams struct {
	// Start time. It must conform to the RFC3339 standard.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time. It must conform to the RFC3339 standard.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Number of entries per page
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// Current page
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// List of sites
	Zones []*string `json:"Zones,omitempty" name:"Zones"`

	// List of domain names
	Domains []*string `json:"Domains,omitempty" name:"Domains"`
}

type DownloadL7LogsRequest struct {
	*tchttp.BaseRequest
	
	// Start time. It must conform to the RFC3339 standard.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time. It must conform to the RFC3339 standard.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Number of entries per page
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// Current page
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// List of sites
	Zones []*string `json:"Zones,omitempty" name:"Zones"`

	// List of domain names
	Domains []*string `json:"Domains,omitempty" name:"Domains"`
}

func (r *DownloadL7LogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadL7LogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageSize")
	delete(f, "PageNo")
	delete(f, "Zones")
	delete(f, "Domains")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DownloadL7LogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadL7LogsResponseParams struct {
	// Layer-7 offline log data
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Data []*L7OfflineLog `json:"Data,omitempty" name:"Data"`

	// Page size
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// Page number
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// Total number of pages
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Pages *int64 `json:"Pages,omitempty" name:"Pages"`

	// Total number of entries
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	TotalSize *int64 `json:"TotalSize,omitempty" name:"TotalSize"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DownloadL7LogsResponse struct {
	*tchttp.BaseResponse
	Response *DownloadL7LogsResponseParams `json:"Response"`
}

func (r *DownloadL7LogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadL7LogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FailReason struct {
	// Failure reason
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// List of resources failed to be processed. 
	//  
	Targets []*string `json:"Targets,omitempty" name:"Targets"`
}

type Header struct {
	// HTTP header name
	Name *string `json:"Name,omitempty" name:"Name"`

	// HTTP header value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type L7OfflineLog struct {
	// Start time of the log packaging
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	LogTime *int64 `json:"LogTime,omitempty" name:"LogTime"`

	// Site name
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Log size, in bytes
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// Download address
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Log package name
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	LogPacketName *string `json:"LogPacketName,omitempty" name:"LogPacketName"`
}

type Task struct {
	// Task ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// Status of the task
	Status *string `json:"Status,omitempty" name:"Status"`

	// Resource
	Target *string `json:"Target,omitempty" name:"Target"`

	// Task type
	Type *string `json:"Type,omitempty" name:"Type"`

	// Task creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Task completion time
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type Zone struct {
	// Site ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Site name
	Name *string `json:"Name,omitempty" name:"Name"`

	// List of name servers used by the site
	OriginalNameServers []*string `json:"OriginalNameServers,omitempty" name:"OriginalNameServers"`

	// List of name servers assigned by Tencent Cloud
	NameServers []*string `json:"NameServers,omitempty" name:"NameServers"`

	// Site status
	// - `active`: The name server is switched.
	// - `pending`: The name server is not switched.
	// - `moved`: The name server is moved.
	// - `deactivated`: The name server is blocked.
	Status *string `json:"Status,omitempty" name:"Status"`

	// How the site is connected to EdgeOne.
	// - `full`: The site is connected via name server.
	// - `partial`: The site is connected via CNAME.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Indicates whether the site is disabled
	Paused *bool `json:"Paused,omitempty" name:"Paused"`

	// Site creation date
	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// Site modification date
	ModifiedOn *string `json:"ModifiedOn,omitempty" name:"ModifiedOn"`
}

type ZoneFilter struct {
	// Filters by the field name. Vaules:
	// - `name`: Site name.
	// - `status`: Site status.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Filters by the field value
	Values []*string `json:"Values,omitempty" name:"Values"`

	// Specifies whether to enable fuzzy query. It’s only available when filter name is `name`. If it’s enabled, the length of `Values` must be 1.
	Fuzzy *bool `json:"Fuzzy,omitempty" name:"Fuzzy"`
}