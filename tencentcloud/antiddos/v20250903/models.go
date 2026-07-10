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

package v20250903

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type DDoSBlockRecord struct {
	// <p>Blocked resources, public IP address, for example:</p><ul><li>Public IP address: 117.175.94.231.</li></ul>
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// <p>The time when it was blocked.</p>
	BlockTime *string `json:"BlockTime,omitnil,omitempty" name:"BlockTime"`

	// <p>Blocking and unblocking status.</p><p>Enumeration value:</p><ul><li>Blocked: Blocked;</li><li>Unblocking: Unblocking;</li><li>Unblocked: Unblocked.</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type DDoSUnblockQuota struct {
	// <p>Total quota of the number of unblocking times.</p>
	TotalQuota *uint64 `json:"TotalQuota,omitnil,omitempty" name:"TotalQuota"`

	// <p>Total quota used.</p>
	UsedQuota *uint64 `json:"UsedQuota,omitnil,omitempty" name:"UsedQuota"`

	// <p>Start time when the quota takes effect.</p>
	QuotaStartTime *string `json:"QuotaStartTime,omitnil,omitempty" name:"QuotaStartTime"`

	// <p>End time when the quota takes effect.</p>
	QuotaEndTime *string `json:"QuotaEndTime,omitnil,omitempty" name:"QuotaEndTime"`
}

// Predefined struct for user
type DescribeDDoSBlockRecordsRequestParams struct {
	// <p>Start time of the query. Support up to data query for the past one year.</p><p>Parameter format: 2026-02-04T11:30:00+08:00.</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>End time of query. The query time range (EndTime - StartTime) must be less than or equal to 31 days.</p><p>Parameter format: 2026-03-04T11:30:00+08:00.</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>Filter criteria. The maximum number of Filters.Values is 20. If this parameter is left empty, return the current list of resources blocked under the appid. Detailed filter criteria:</p><ul><li> Resource: Filter by blocked IP or six-segment resource format;</li><li> Status: Filter by blocked resource status.</li></ul><p>When Filters.Name value is Status, Filters.Values valid values:</p><ul><li>Blocked: blocked;</li><li>Unblocking: unblocking;</li><li>Unblocked: unblocked.</li></ul>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>Paginated query limit count. Maximum value: 100.</p><p>Default value: 20</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Paginated query offset.</p><p>Default value: 0.</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeDDoSBlockRecordsRequest struct {
	*tchttp.BaseRequest
	
	// <p>Start time of the query. Support up to data query for the past one year.</p><p>Parameter format: 2026-02-04T11:30:00+08:00.</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>End time of query. The query time range (EndTime - StartTime) must be less than or equal to 31 days.</p><p>Parameter format: 2026-03-04T11:30:00+08:00.</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>Filter criteria. The maximum number of Filters.Values is 20. If this parameter is left empty, return the current list of resources blocked under the appid. Detailed filter criteria:</p><ul><li> Resource: Filter by blocked IP or six-segment resource format;</li><li> Status: Filter by blocked resource status.</li></ul><p>When Filters.Name value is Status, Filters.Values valid values:</p><ul><li>Blocked: blocked;</li><li>Unblocking: unblocking;</li><li>Unblocked: unblocked.</li></ul>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>Paginated query limit count. Maximum value: 100.</p><p>Default value: 20</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Paginated query offset.</p><p>Default value: 0.</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeDDoSBlockRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSBlockRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSBlockRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSBlockRecordsResponseParams struct {
	// <p>Total number of block and unblock records.</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>Unblock record.</p>
	BlockRecords []*DDoSBlockRecord `json:"BlockRecords,omitnil,omitempty" name:"BlockRecords"`

	// <p>Quota information of the number of unblocking times.</p>
	UnblockQuotaInfo *DDoSUnblockQuota `json:"UnblockQuotaInfo,omitnil,omitempty" name:"UnblockQuotaInfo"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDDoSBlockRecordsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDoSBlockRecordsResponseParams `json:"Response"`
}

func (r *DescribeDDoSBlockRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSBlockRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// <p>Fields to be filtered. Check corresponding API for specific available values.</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Field's filter value.</p>
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

// Predefined struct for user
type UnblockResourcesRequestParams struct {
	// <p>List of resources to apply for unblocking. Supports unblocking based on public IP. You can obtain detailed resource information of blocked resources through the DescribeDDoSBlockRecords API. Parameter example:</p><ul><li>Public IP: 117.175.94.230.</li></ul><p>Input parameter limit: Maximum list length is 10.</p>
	Resources []*string `json:"Resources,omitnil,omitempty" name:"Resources"`
}

type UnblockResourcesRequest struct {
	*tchttp.BaseRequest
	
	// <p>List of resources to apply for unblocking. Supports unblocking based on public IP. You can obtain detailed resource information of blocked resources through the DescribeDDoSBlockRecords API. Parameter example:</p><ul><li>Public IP: 117.175.94.230.</li></ul><p>Input parameter limit: Maximum list length is 10.</p>
	Resources []*string `json:"Resources,omitnil,omitempty" name:"Resources"`
}

func (r *UnblockResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnblockResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Resources")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnblockResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnblockResourcesResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UnblockResourcesResponse struct {
	*tchttp.BaseResponse
	Response *UnblockResourcesResponseParams `json:"Response"`
}

func (r *UnblockResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnblockResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}