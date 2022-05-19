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

	// The target resource to be purged. One target per line.
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

type CreatePurgeTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task ID
		JobId *string `json:"JobId,omitempty" name:"JobId"`

		// List of failed tasks and reasons
	// Note: This field may return `null`, indicating that no valid value can be obtained.
		FailedList []*FailReason `json:"FailedList,omitempty" name:"FailedList"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribePurgeTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total entries that match the specified query condition
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of tasks returned
		Tasks []*Task `json:"Tasks,omitempty" name:"Tasks"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeZonesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of sites that match the specified conditions
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Details of sites
	// Note: This field may return `null`, indicating that no valid value can be obtained.
		Zones []*Zone `json:"Zones,omitempty" name:"Zones"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type FailReason struct {

	// Failure reason
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// List of resources failed to be purged
	Targets []*string `json:"Targets,omitempty" name:"Targets"`
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
