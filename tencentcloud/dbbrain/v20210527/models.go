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

package v20210527

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

// Predefined struct for user
type CreateProxySessionKillTaskRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Service type. Valid value: `redis` (TencentDB for Redis).
	Product *string `json:"Product,omitempty" name:"Product"`
}

type CreateProxySessionKillTaskRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Service type. Valid value: `redis` (TencentDB for Redis).
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *CreateProxySessionKillTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProxySessionKillTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProxySessionKillTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProxySessionKillTaskResponseParams struct {
	// Async task ID that is returned after the session killing task is created.
	AsyncRequestId *int64 `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateProxySessionKillTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateProxySessionKillTaskResponseParams `json:"Response"`
}

func (r *CreateProxySessionKillTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProxySessionKillTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBDiagEventsRequestParams struct {
	// Start time in the format of “2021-05-27 00:00:00”. The earliest time that can be queried is 30 days before the current time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time in the format of "2021-05-27 01:00:00". The interval between the end time and the start time can be up to 7 days.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Risk level list. Valid values in descending order of severity: `1` (critical), `2` (serious), `3` (alarm), `4` (warning), `5` (healthy).
	Severities []*int64 `json:"Severities,omitempty" name:"Severities"`

	// Instance ID list.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 50.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeDBDiagEventsRequest struct {
	*tchttp.BaseRequest
	
	// Start time in the format of “2021-05-27 00:00:00”. The earliest time that can be queried is 30 days before the current time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time in the format of "2021-05-27 01:00:00". The interval between the end time and the start time can be up to 7 days.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Risk level list. Valid values in descending order of severity: `1` (critical), `2` (serious), `3` (alarm), `4` (warning), `5` (healthy).
	Severities []*int64 `json:"Severities,omitempty" name:"Severities"`

	// Instance ID list.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 50.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDBDiagEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBDiagEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Severities")
	delete(f, "InstanceIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBDiagEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBDiagEventsResponseParams struct {
	// Total number of diagnosis events.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Diagnosis event list.
	Items []*DiagHistoryEventItem `json:"Items,omitempty" name:"Items"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDBDiagEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBDiagEventsResponseParams `json:"Response"`
}

func (r *DescribeDBDiagEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBDiagEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDiagDBInstancesRequestParams struct {
	// Whether it is an instance supported by DBbrain. It is fixed to `true`.
	IsSupported *bool `json:"IsSupported,omitempty" name:"IsSupported"`

	// Service type. Valid values: mysql (TencentDB for MySQL), cynosdb (TDSQL-C for MySQL). Default value: mysql.
	Product *string `json:"Product,omitempty" name:"Product"`

	// Pagination parameter indicating the offset.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Pagination parameter. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Query by instance name.
	InstanceNames []*string `json:"InstanceNames,omitempty" name:"InstanceNames"`

	// Query by instance ID.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Query by region.
	Regions []*string `json:"Regions,omitempty" name:"Regions"`
}

type DescribeDiagDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Whether it is an instance supported by DBbrain. It is fixed to `true`.
	IsSupported *bool `json:"IsSupported,omitempty" name:"IsSupported"`

	// Service type. Valid values: mysql (TencentDB for MySQL), cynosdb (TDSQL-C for MySQL). Default value: mysql.
	Product *string `json:"Product,omitempty" name:"Product"`

	// Pagination parameter indicating the offset.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Pagination parameter. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Query by instance name.
	InstanceNames []*string `json:"InstanceNames,omitempty" name:"InstanceNames"`

	// Query by instance ID.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Query by region.
	Regions []*string `json:"Regions,omitempty" name:"Regions"`
}

func (r *DescribeDiagDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiagDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IsSupported")
	delete(f, "Product")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "InstanceNames")
	delete(f, "InstanceIds")
	delete(f, "Regions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDiagDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDiagDBInstancesResponseParams struct {
	// Total number of instances.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Status of all instance inspection. 0: all instance inspection enabled, 1: all instance inspection disabled.
	DbScanStatus *int64 `json:"DbScanStatus,omitempty" name:"DbScanStatus"`

	// Instance information.
	Items []*InstanceInfo `json:"Items,omitempty" name:"Items"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDiagDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDiagDBInstancesResponseParams `json:"Response"`
}

func (r *DescribeDiagDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiagDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMySqlProcessListRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Thread ID, which is used to filter the thread list.
	ID *uint64 `json:"ID,omitempty" name:"ID"`

	// Thread operation account name, which is used to filter the thread list.
	User *string `json:"User,omitempty" name:"User"`

	// Thread operation host address, which is used to filter the thread list.
	Host *string `json:"Host,omitempty" name:"Host"`

	// Thread operation database, which is used to filter the thread list.
	DB *string `json:"DB,omitempty" name:"DB"`

	// Thread operation status, which is used to filter the thread list.
	State *string `json:"State,omitempty" name:"State"`

	// Thread execution type, which is used to filter the thread list.
	Command *string `json:"Command,omitempty" name:"Command"`

	// Minimum operation duration of the thread in seconds, which is used to filter the list of threads whose operation duration is greater than this value.
	Time *uint64 `json:"Time,omitempty" name:"Time"`

	// Thread operation statement, which is used to filter the thread list.
	Info *string `json:"Info,omitempty" name:"Info"`

	// Number of returned results. Default value: 20.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Service type. Valid values: mysql (TencentDB for MySQL), cynosdb (TDSQL-C for MySQL). Default value: mysql.
	Product *string `json:"Product,omitempty" name:"Product"`
}

type DescribeMySqlProcessListRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Thread ID, which is used to filter the thread list.
	ID *uint64 `json:"ID,omitempty" name:"ID"`

	// Thread operation account name, which is used to filter the thread list.
	User *string `json:"User,omitempty" name:"User"`

	// Thread operation host address, which is used to filter the thread list.
	Host *string `json:"Host,omitempty" name:"Host"`

	// Thread operation database, which is used to filter the thread list.
	DB *string `json:"DB,omitempty" name:"DB"`

	// Thread operation status, which is used to filter the thread list.
	State *string `json:"State,omitempty" name:"State"`

	// Thread execution type, which is used to filter the thread list.
	Command *string `json:"Command,omitempty" name:"Command"`

	// Minimum operation duration of the thread in seconds, which is used to filter the list of threads whose operation duration is greater than this value.
	Time *uint64 `json:"Time,omitempty" name:"Time"`

	// Thread operation statement, which is used to filter the thread list.
	Info *string `json:"Info,omitempty" name:"Info"`

	// Number of returned results. Default value: 20.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Service type. Valid values: mysql (TencentDB for MySQL), cynosdb (TDSQL-C for MySQL). Default value: mysql.
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeMySqlProcessListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMySqlProcessListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ID")
	delete(f, "User")
	delete(f, "Host")
	delete(f, "DB")
	delete(f, "State")
	delete(f, "Command")
	delete(f, "Time")
	delete(f, "Info")
	delete(f, "Limit")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMySqlProcessListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMySqlProcessListResponseParams struct {
	// List of real-time threads.
	ProcessList []*MySqlProcess `json:"ProcessList,omitempty" name:"ProcessList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMySqlProcessListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMySqlProcessListResponseParams `json:"Response"`
}

func (r *DescribeMySqlProcessListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMySqlProcessListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProxySessionKillTasksRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// The async session killing task ID, which is obtained after the API `CreateProxySessionKillTask` is successfully called.
	AsyncRequestIds []*int64 `json:"AsyncRequestIds,omitempty" name:"AsyncRequestIds"`

	// Service type. Valid value: `redis` (TencentDB for Redis).
	Product *string `json:"Product,omitempty" name:"Product"`
}

type DescribeProxySessionKillTasksRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// The async session killing task ID, which is obtained after the API `CreateProxySessionKillTask` is successfully called.
	AsyncRequestIds []*int64 `json:"AsyncRequestIds,omitempty" name:"AsyncRequestIds"`

	// Service type. Valid value: `redis` (TencentDB for Redis).
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeProxySessionKillTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxySessionKillTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AsyncRequestIds")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProxySessionKillTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProxySessionKillTasksResponseParams struct {
	// Session killing task details.
	Tasks []*TaskInfo `json:"Tasks,omitempty" name:"Tasks"`

	// Total number of tasks.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeProxySessionKillTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProxySessionKillTasksResponseParams `json:"Response"`
}

func (r *DescribeProxySessionKillTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxySessionKillTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRedisTopKeyPrefixListRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Date for query, such as `2021-05-27`. You can select a date as early as in the last 30 days for query.
	Date *string `json:"Date,omitempty" name:"Date"`

	// Service type. Valid value: `redis` (TencentDB for Redis).
	Product *string `json:"Product,omitempty" name:"Product"`

	// The number of queried items. Default value: `20`. Max value: `100`.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeRedisTopKeyPrefixListRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Date for query, such as `2021-05-27`. You can select a date as early as in the last 30 days for query.
	Date *string `json:"Date,omitempty" name:"Date"`

	// Service type. Valid value: `redis` (TencentDB for Redis).
	Product *string `json:"Product,omitempty" name:"Product"`

	// The number of queried items. Default value: `20`. Max value: `100`.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeRedisTopKeyPrefixListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRedisTopKeyPrefixListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Date")
	delete(f, "Product")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRedisTopKeyPrefixListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRedisTopKeyPrefixListResponseParams struct {
	// List of top key prefixes
	Items []*RedisPreKeySpaceData `json:"Items,omitempty" name:"Items"`

	// Data collection timestamp in seconds
	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRedisTopKeyPrefixListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRedisTopKeyPrefixListResponseParams `json:"Response"`
}

func (r *DescribeRedisTopKeyPrefixListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRedisTopKeyPrefixListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogUserHostStatsRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Start time of the time range in the format of yyyy-MM-dd HH:mm:ss, such as 2019-09-10 12:13:14.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of the time range in the format of yyyy-MM-dd HH:mm:ss, such as 2019-09-10 12:13:14.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Service type. Valid values: mysql (TencentDB for MySQL), cynosdb (TDSQL-C for MySQL). Default value: mysql.
	Product *string `json:"Product,omitempty" name:"Product"`

	// MD5 value of SOL template
	Md5 *string `json:"Md5,omitempty" name:"Md5"`
}

type DescribeSlowLogUserHostStatsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Start time of the time range in the format of yyyy-MM-dd HH:mm:ss, such as 2019-09-10 12:13:14.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of the time range in the format of yyyy-MM-dd HH:mm:ss, such as 2019-09-10 12:13:14.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Service type. Valid values: mysql (TencentDB for MySQL), cynosdb (TDSQL-C for MySQL). Default value: mysql.
	Product *string `json:"Product,omitempty" name:"Product"`

	// MD5 value of SOL template
	Md5 *string `json:"Md5,omitempty" name:"Md5"`
}

func (r *DescribeSlowLogUserHostStatsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogUserHostStatsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Product")
	delete(f, "Md5")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowLogUserHostStatsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogUserHostStatsResponseParams struct {
	// Total number of source addresses.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Detailed list of the proportion of slow logs from each source address.
	Items []*SlowLogHost `json:"Items,omitempty" name:"Items"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSlowLogUserHostStatsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowLogUserHostStatsResponseParams `json:"Response"`
}

func (r *DescribeSlowLogUserHostStatsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogUserHostStatsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopSpaceSchemasRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Number of returned top databases. Maximum value: 100. Default value: 20.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Field used to sort top databases. Valid values: DataLength, IndexLength, TotalLength, DataFree, FragRatio, TableRows, PhysicalFileSize (supported only by TencentDB for MySQL instances). For TencentDB for MySQL instances, the default value is `PhysicalFileSize`. For other database instances, the default value is `TotalLength`.
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// Service type. Valid values: mysql (TencentDB for MySQL), cynosdb (TDSQL-C for MySQL). Default value: mysql.
	Product *string `json:"Product,omitempty" name:"Product"`
}

type DescribeTopSpaceSchemasRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Number of returned top databases. Maximum value: 100. Default value: 20.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Field used to sort top databases. Valid values: DataLength, IndexLength, TotalLength, DataFree, FragRatio, TableRows, PhysicalFileSize (supported only by TencentDB for MySQL instances). For TencentDB for MySQL instances, the default value is `PhysicalFileSize`. For other database instances, the default value is `TotalLength`.
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// Service type. Valid values: mysql (TencentDB for MySQL), cynosdb (TDSQL-C for MySQL). Default value: mysql.
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeTopSpaceSchemasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopSpaceSchemasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "SortBy")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopSpaceSchemasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopSpaceSchemasResponseParams struct {
	// List of the returned space statistics of top databases.
	TopSpaceSchemas []*SchemaSpaceData `json:"TopSpaceSchemas,omitempty" name:"TopSpaceSchemas"`

	// Timestamp (in seconds) of database space data collection points
	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTopSpaceSchemasResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopSpaceSchemasResponseParams `json:"Response"`
}

func (r *DescribeTopSpaceSchemasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopSpaceSchemasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserSqlAdviceRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// SQL statement.
	SqlText *string `json:"SqlText,omitempty" name:"SqlText"`

	// Database name.
	Schema *string `json:"Schema,omitempty" name:"Schema"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL), `dbbrain-mysql` (self-built MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitempty" name:"Product"`
}

type DescribeUserSqlAdviceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// SQL statement.
	SqlText *string `json:"SqlText,omitempty" name:"SqlText"`

	// Database name.
	Schema *string `json:"Schema,omitempty" name:"Schema"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL), `dbbrain-mysql` (self-built MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeUserSqlAdviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserSqlAdviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SqlText")
	delete(f, "Schema")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserSqlAdviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserSqlAdviceResponseParams struct {
	// SQL statement optimization suggestions, which can be parsed into JSON arrays. If there is no need for optimization, the output will be empty.
	Advices *string `json:"Advices,omitempty" name:"Advices"`

	// Notes of SQL statement optimization suggestions, which can be parsed into String arrays. If there is no need for optimization, the output will be empty.
	Comments *string `json:"Comments,omitempty" name:"Comments"`

	// SQL statement.
	SqlText *string `json:"SqlText,omitempty" name:"SqlText"`

	// Database name.
	Schema *string `json:"Schema,omitempty" name:"Schema"`

	// DDL information of related tables, which can be parsed into JSON arrays.
	Tables *string `json:"Tables,omitempty" name:"Tables"`

	// SQL execution plan, which can be parsed into JSON arrays. If there is no need for optimization, the output will be empty.
	SqlPlan *string `json:"SqlPlan,omitempty" name:"SqlPlan"`

	// Cost saving details after SQL statement optimization, which can be parsed into JSON arrays. If there is no need for optimization, the output will be empty.
	Cost *string `json:"Cost,omitempty" name:"Cost"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUserSqlAdviceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserSqlAdviceResponseParams `json:"Response"`
}

func (r *DescribeUserSqlAdviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserSqlAdviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiagHistoryEventItem struct {
	// Diagnosis type.
	DiagType *string `json:"DiagType,omitempty" name:"DiagType"`

	// End time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Unique event ID.
	EventId *int64 `json:"EventId,omitempty" name:"EventId"`

	// Severity, which can be divided into 5 levels: 1: fatal, 2: severe, 3: warning, 4: notice, 5: healthy.
	Severity *int64 `json:"Severity,omitempty" name:"Severity"`

	// Diagnosis summary.
	Outline *string `json:"Outline,omitempty" name:"Outline"`

	// Diagnosis item description.
	DiagItem *string `json:"DiagItem,omitempty" name:"DiagItem"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Reserved field.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// Region.
	Region *string `json:"Region,omitempty" name:"Region"`
}

type InstanceConfs struct {
	// Whether to enable database inspection. Valid values: Yes, No.
	DailyInspection *string `json:"DailyInspection,omitempty" name:"DailyInspection"`

	// Whether to enable instance overview. Valid values: Yes, No.
	OverviewDisplay *string `json:"OverviewDisplay,omitempty" name:"OverviewDisplay"`
}

type InstanceInfo struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance name.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Instance region.
	Region *string `json:"Region,omitempty" name:"Region"`

	// Health score.
	HealthScore *int64 `json:"HealthScore,omitempty" name:"HealthScore"`

	// Service.
	Product *string `json:"Product,omitempty" name:"Product"`

	// Number of exceptions.
	EventCount *int64 `json:"EventCount,omitempty" name:"EventCount"`

	// Instance type. Valid values: 1 (MASTER), 2 (DR), 3 (RO), 4 (SDR)
	InstanceType *int64 `json:"InstanceType,omitempty" name:"InstanceType"`

	// Number of cores.
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// Memory in MB.
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Disk storage in GB.
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`

	// Database version.
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// Private network address.
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Private network port.
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// Access source.
	Source *string `json:"Source,omitempty" name:"Source"`

	// Group ID.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Group name.
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// Instance status. Valid values: 0 (delivering), 1 (running), 4 (terminating), 5 (isolated)
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Unified subnet ID.
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// TencentDB instance type.
	DeployMode *string `json:"DeployMode,omitempty" name:"DeployMode"`

	// TencentDB instance initialization flag. Valid values: 0 (not initialized), 1 (initialized).
	InitFlag *int64 `json:"InitFlag,omitempty" name:"InitFlag"`

	// Task status.
	TaskStatus *int64 `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// Unified VPC ID.
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// Instance inspection/overview status.
	InstanceConf *InstanceConfs `json:"InstanceConf,omitempty" name:"InstanceConf"`

	// Resource expiration time.
	DeadlineTime *string `json:"DeadlineTime,omitempty" name:"DeadlineTime"`

	// Whether it is an instance supported by DBbrain.
	IsSupported *bool `json:"IsSupported,omitempty" name:"IsSupported"`

	// Status of instance security audit log. Valid values: ON (enabled), OFF (disabled).
	SecAuditStatus *string `json:"SecAuditStatus,omitempty" name:"SecAuditStatus"`

	// Status of instance audit log. Valid values: ALL_AUDIT (full audit is enabled), RULE_AUDIT (rule audit is enabled), UNBOUND (audit is disabled).
	AuditPolicyStatus *string `json:"AuditPolicyStatus,omitempty" name:"AuditPolicyStatus"`

	// Running status of instance audit log. Valid values: normal (running), paused (suspension due to overdue payment).
	AuditRunningStatus *string `json:"AuditRunningStatus,omitempty" name:"AuditRunningStatus"`
}

type MySqlProcess struct {
	// Thread ID.
	ID *string `json:"ID,omitempty" name:"ID"`

	// Thread operation account name.
	User *string `json:"User,omitempty" name:"User"`

	// Thread operation host address.
	Host *string `json:"Host,omitempty" name:"Host"`

	// Thread operation database.
	DB *string `json:"DB,omitempty" name:"DB"`

	// Thread operation status.
	State *string `json:"State,omitempty" name:"State"`

	// Thread execution type.
	Command *string `json:"Command,omitempty" name:"Command"`

	// Thread operation duration in seconds.
	Time *string `json:"Time,omitempty" name:"Time"`

	// Thread operation statement.
	Info *string `json:"Info,omitempty" name:"Info"`
}

type RedisPreKeySpaceData struct {
	// Average element length
	AveElementSize *int64 `json:"AveElementSize,omitempty" name:"AveElementSize"`

	// Total memory usage in bytes
	Length *int64 `json:"Length,omitempty" name:"Length"`

	// Key prefix
	KeyPreIndex *string `json:"KeyPreIndex,omitempty" name:"KeyPreIndex"`

	// The number of elements
	ItemCount *int64 `json:"ItemCount,omitempty" name:"ItemCount"`

	// The number of keys
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// The max element length
	MaxElementSize *int64 `json:"MaxElementSize,omitempty" name:"MaxElementSize"`
}

type SchemaSpaceData struct {
	// Database name.
	TableSchema *string `json:"TableSchema,omitempty" name:"TableSchema"`

	// Data space in MB.
	DataLength *float64 `json:"DataLength,omitempty" name:"DataLength"`

	// Index space in MB.
	IndexLength *float64 `json:"IndexLength,omitempty" name:"IndexLength"`

	// Fragmented space in MB.
	DataFree *float64 `json:"DataFree,omitempty" name:"DataFree"`

	// Total space usage in MB.
	TotalLength *float64 `json:"TotalLength,omitempty" name:"TotalLength"`

	// Fragmentation rate in %.
	FragRatio *float64 `json:"FragRatio,omitempty" name:"FragRatio"`

	// Number of rows.
	TableRows *int64 `json:"TableRows,omitempty" name:"TableRows"`

	// Total size in MB of physical files exclusive to all tables in the database.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PhysicalFileSize *float64 `json:"PhysicalFileSize,omitempty" name:"PhysicalFileSize"`
}

type SlowLogHost struct {
	// Source addresses.
	UserHost *string `json:"UserHost,omitempty" name:"UserHost"`

	// Proportion (in %) of slow logs from this source address to the total number of slow logs.
	Ratio *float64 `json:"Ratio,omitempty" name:"Ratio"`

	// Number of slow logs from this source address.
	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type TaskInfo struct {
	// Async task ID.
	AsyncRequestId *int64 `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// List of all proxies of the current instance.
	InstProxyList []*string `json:"InstProxyList,omitempty" name:"InstProxyList"`

	// Total number of proxies of the current instance.
	InstProxyCount *int64 `json:"InstProxyCount,omitempty" name:"InstProxyCount"`

	// Task creation time.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Task start time.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Task status. Valid values: `created` (create), `chosen` (to be executed), `running` (being executed), `failed` (failed), and `finished` (completed).
	TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// IDs of the proxies that have completed the session killing tasks.
	FinishedProxyList []*string `json:"FinishedProxyList,omitempty" name:"FinishedProxyList"`

	// IDs of the proxies that failed to execute the session killing tasks.
	FailedProxyList []*string `json:"FailedProxyList,omitempty" name:"FailedProxyList"`

	// Task end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Task progress.
	Progress *int64 `json:"Progress,omitempty" name:"Progress"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}