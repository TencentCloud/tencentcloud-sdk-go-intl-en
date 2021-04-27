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

package v20180330

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type ActivateSubscribeRequest struct {
	*tchttp.BaseRequest

	// Subscription instance ID.
	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`

	// Database Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Data subscription type. 0: full instance subscription, 1: data subscription, 2: structure subscription, 3: data subscription and structure subscription
	SubscribeObjectType *int64 `json:"SubscribeObjectType,omitempty" name:"SubscribeObjectType"`

	// Subscription object
	Objects *SubscribeObject `json:"Objects,omitempty" name:"Objects"`

	// Subnet of data subscription service, which is the subnet of the database instance by default.
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// Subscription service port. Default value: 7507
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ActivateSubscribeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubscribeId")
	delete(f, "InstanceId")
	delete(f, "SubscribeObjectType")
	delete(f, "Objects")
	delete(f, "UniqSubnetId")
	delete(f, "Vport")
	if len(f) > 0 {
		return errors.New("ActivateSubscribeRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ActivateSubscribeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Data subscription configuration task ID.
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ActivateSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CompleteMigrateJobRequest struct {
	*tchttp.BaseRequest

	// Data migration task ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CompleteMigrateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return errors.New("CompleteMigrateJobRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CompleteMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CompleteMigrateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConsistencyParams struct {

	// Data content check parameter, which refers to the proportion of the rows selected for data comparison in all the rows of the table. Value: an integer between 1 and 100.
	SelectRowsPerTable *int64 `json:"SelectRowsPerTable,omitempty" name:"SelectRowsPerTable"`

	// Data content check parameter, which refers to the proportion of the tables selected for data detection in all the tables. Value: an integer between 1 and 100.
	TablesSelectAll *int64 `json:"TablesSelectAll,omitempty" name:"TablesSelectAll"`

	// Data quantity check parameter, which checks whether the numbers of rows are identical. It refers to the proportion of the tables selected for quantity check in all the tables. Value: an integer between 1 and 100.
	TablesSelectCount *int64 `json:"TablesSelectCount,omitempty" name:"TablesSelectCount"`
}

type CreateMigrateCheckJobRequest struct {
	*tchttp.BaseRequest

	// Data migration task ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMigrateCheckJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return errors.New("CreateMigrateCheckJobRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateMigrateCheckJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMigrateCheckJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMigrateJobRequest struct {
	*tchttp.BaseRequest

	// Data migration task name
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// Migration task configuration options
	MigrateOption *MigrateOption `json:"MigrateOption,omitempty" name:"MigrateOption"`

	// Source instance database type, which currently supports MySQL, Redis, MongoDB, PostgreSQL, MariaDB, and Percona. For more information on the supported types in a specific region, see the migration task creation page in the console.
	SrcDatabaseType *string `json:"SrcDatabaseType,omitempty" name:"SrcDatabaseType"`

	// Source instance access type. Valid values: extranet (public network), cvm (CVM-based self-created instance), dcg (Direct Connect-enabled instance), vpncloud (Tencent Cloud VPN-enabled instance), cdb (TencentDB instance), ccn (CCN instance)
	SrcAccessType *string `json:"SrcAccessType,omitempty" name:"SrcAccessType"`

	// Source instance information, which is correlated with the migration task type
	SrcInfo *SrcInfo `json:"SrcInfo,omitempty" name:"SrcInfo"`

	// Target instance access type, which currently supports MySQL, Redis, MongoDB, PostgreSQL, MariaDB, and Percona. For more information on the supported types in a specific region, see the migration task creation page in the console.
	DstDatabaseType *string `json:"DstDatabaseType,omitempty" name:"DstDatabaseType"`

	// Target instance access type, which currently only supports cdb (TencentDB instance)
	DstAccessType *string `json:"DstAccessType,omitempty" name:"DstAccessType"`

	// Destination instance information
	DstInfo *DstInfo `json:"DstInfo,omitempty" name:"DstInfo"`

	// Information of the source table to be migrated, which is described in JSON string format. It is required if MigrateOption.MigrateObject is 2 (migrating the specified table).
	// For databases with a database-table structure:
	// [{Database:db1,Table:[table1,table2]},{Database:db2}]
	// For databases with a database-schema-table structure:
	// [{Database:db1,Schema:s1
	// Table:[table1,table2]},{Database:db1,Schema:s2
	// Table:[table1,table2]},{Database:db2,Schema:s1
	// Table:[table1,table2]},{Database:db3},{Database:db4
	// Schema:s1}]
	DatabaseInfo *string `json:"DatabaseInfo,omitempty" name:"DatabaseInfo"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMigrateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobName")
	delete(f, "MigrateOption")
	delete(f, "SrcDatabaseType")
	delete(f, "SrcAccessType")
	delete(f, "SrcInfo")
	delete(f, "DstDatabaseType")
	delete(f, "DstAccessType")
	delete(f, "DstInfo")
	delete(f, "DatabaseInfo")
	if len(f) > 0 {
		return errors.New("CreateMigrateJobRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Data migration task ID
		JobId *string `json:"JobId,omitempty" name:"JobId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMigrateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSubscribeRequest struct {
	*tchttp.BaseRequest

	// Subscribed database type. Currently, MySQL is supported
	Product *string `json:"Product,omitempty" name:"Product"`

	// Instance billing mode, which is always 1 (hourly billing),
	PayType *int64 `json:"PayType,omitempty" name:"PayType"`

	// Purchase duration in months, which is required if `PayType` is 0. Maximum value: 120 (this field is not required of global site users and is better to be hidden)
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`

	// Quantity. Default value: 1. Maximum value: 10
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// Whether to auto-renew. Default value: 0. This flag does not take effect for hourly billed instances (this field should be hidden from global site users)
	AutoRenew *int64 `json:"AutoRenew,omitempty" name:"AutoRenew"`

	// Instance resource tags
	Tags []*TagItem `json:"Tags,omitempty" name:"Tags" list`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSubscribeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "PayType")
	delete(f, "Duration")
	delete(f, "Count")
	delete(f, "AutoRenew")
	delete(f, "Tags")
	if len(f) > 0 {
		return errors.New("CreateSubscribeRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateSubscribeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Data subscription instance ID array
	// Note: this field may return null, indicating that no valid values can be obtained.
		SubscribeIds []*string `json:"SubscribeIds,omitempty" name:"SubscribeIds" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSyncCheckJobRequest struct {
	*tchttp.BaseRequest

	// Disaster recovery sync task ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSyncCheckJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return errors.New("CreateSyncCheckJobRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateSyncCheckJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSyncCheckJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSyncJobRequest struct {
	*tchttp.BaseRequest

	// Disaster recovery sync task name
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// Configuration options of a disaster recovery sync task
	SyncOption *SyncOption `json:"SyncOption,omitempty" name:"SyncOption"`

	// Source instance database type, which currently only supports mysql
	SrcDatabaseType *string `json:"SrcDatabaseType,omitempty" name:"SrcDatabaseType"`

	// Source instance access type, which currently only supports cdb (TencentDB instances)
	SrcAccessType *string `json:"SrcAccessType,omitempty" name:"SrcAccessType"`

	// Source instance information
	SrcInfo *SyncInstanceInfo `json:"SrcInfo,omitempty" name:"SrcInfo"`

	// Target instance access type, which currently only supports mysql
	DstDatabaseType *string `json:"DstDatabaseType,omitempty" name:"DstDatabaseType"`

	// Target instance access type, which currently only supports cdb (TencentDB instances)
	DstAccessType *string `json:"DstAccessType,omitempty" name:"DstAccessType"`

	// Target instance information
	DstInfo *SyncInstanceInfo `json:"DstInfo,omitempty" name:"DstInfo"`

	// Information of the source table to be synced, which is described in JSON string format.
	// For databases with a database-table structure:
	// [{Database:db1,Table:[table1,table2]},{Database:db2}]
	DatabaseInfo *string `json:"DatabaseInfo,omitempty" name:"DatabaseInfo"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSyncJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobName")
	delete(f, "SyncOption")
	delete(f, "SrcDatabaseType")
	delete(f, "SrcAccessType")
	delete(f, "SrcInfo")
	delete(f, "DstDatabaseType")
	delete(f, "DstAccessType")
	delete(f, "DstInfo")
	delete(f, "DatabaseInfo")
	if len(f) > 0 {
		return errors.New("CreateSyncJobRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateSyncJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Disaster recovery sync task ID
		JobId *string `json:"JobId,omitempty" name:"JobId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMigrateJobRequest struct {
	*tchttp.BaseRequest

	// Data migration task ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMigrateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return errors.New("DeleteMigrateJobRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMigrateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSyncJobRequest struct {
	*tchttp.BaseRequest

	// ID of the disaster recovery sync task to be deleted
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSyncJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return errors.New("DeleteSyncJobRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSyncJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAsyncRequestInfoRequest struct {
	*tchttp.BaseRequest

	// Task ID
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAsyncRequestInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AsyncRequestId")
	if len(f) > 0 {
		return errors.New("DescribeAsyncRequestInfoRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAsyncRequestInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task execution result information
		Info *string `json:"Info,omitempty" name:"Info"`

		// Task execution status. Valid values: success, failed, running
		Status *string `json:"Status,omitempty" name:"Status"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAsyncRequestInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMigrateCheckJobRequest struct {
	*tchttp.BaseRequest

	// Data migration task ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrateCheckJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return errors.New("DescribeMigrateCheckJobRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMigrateCheckJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Check task status: unavailable, starting, running, finished
		Status *string `json:"Status,omitempty" name:"Status"`

		// Task error code
		ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`

		// Task error message
		ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`

		// Check task progress. For example, "30" means 30% completed
		Progress *string `json:"Progress,omitempty" name:"Progress"`

		// Whether the check succeeds. 0: no; 1: yes; 3: not checked
		CheckFlag *int64 `json:"CheckFlag,omitempty" name:"CheckFlag"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrateCheckJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMigrateJobsRequest struct {
	*tchttp.BaseRequest

	// Data migration task ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// Data migration task name
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// Sort by field. Value range: JobId, Status, JobName, MigrateType, RunMode, CreateTime
	Order *string `json:"Order,omitempty" name:"Order"`

	// Sorting order. Value range: ASC (ascending), DESC (descending)
	OrderSeq *string `json:"OrderSeq,omitempty" name:"OrderSeq"`

	// Offset. Default value: 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of the returned instances. Value range: [1, 100]. Default value: 20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrateJobsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "JobName")
	delete(f, "Order")
	delete(f, "OrderSeq")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return errors.New("DescribeMigrateJobsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMigrateJobsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of tasks
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Array of task details
		JobList []*MigrateJobInfo `json:"JobList,omitempty" name:"JobList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrateJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionConfRequest struct {
	*tchttp.BaseRequest
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionConfRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.New("DescribeRegionConfRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionConfResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of purchasable regions
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Purchasable region details
		Items []*SubscribeRegionConf `json:"Items,omitempty" name:"Items" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionConfResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscribeConfRequest struct {
	*tchttp.BaseRequest

	// Subscription instance ID
	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubscribeConfRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubscribeId")
	if len(f) > 0 {
		return errors.New("DescribeSubscribeConfRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscribeConfResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Subscription instance ID
		SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`

		// Subscription instance name
		SubscribeName *string `json:"SubscribeName,omitempty" name:"SubscribeName"`

		// Subscription channel
		ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

		// Subscribed database type
		Product *string `json:"Product,omitempty" name:"Product"`

		// Subscribed instance
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// Subscribed instance status. Valid values: running, offline, isolate
		InstanceStatus *string `json:"InstanceStatus,omitempty" name:"InstanceStatus"`

		// Subscription instance status. Valid values: unconfigure, configuring, configured
		SubsStatus *string `json:"SubsStatus,omitempty" name:"SubsStatus"`

		// Subscription instance lifecycle status. Valid values: normal, isolating, isolated, offlining
		Status *string `json:"Status,omitempty" name:"Status"`

		// Subscription instance creation time
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// Subscription instance isolation time
		IsolateTime *string `json:"IsolateTime,omitempty" name:"IsolateTime"`

		// Subscription instance expiration time
		ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

		// Subscription instance deactivation time
		OfflineTime *string `json:"OfflineTime,omitempty" name:"OfflineTime"`

		// Consumption starting time point of subscription instance
		ConsumeStartTime *string `json:"ConsumeStartTime,omitempty" name:"ConsumeStartTime"`

		// Subscription instance billing mode. 1: hourly billing
		PayType *int64 `json:"PayType,omitempty" name:"PayType"`

		// Subscription channel VIP
		Vip *string `json:"Vip,omitempty" name:"Vip"`

		// Subscription channel port
		Vport *int64 `json:"Vport,omitempty" name:"Vport"`

		// Subscription channel `VpcId`
		UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

		// Subscription channel `SubnetId`
		UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

		// Current SDK consumption time point
		SdkConsumedTime *string `json:"SdkConsumedTime,omitempty" name:"SdkConsumedTime"`

		// Subscription SDK IP address
		SdkHost *string `json:"SdkHost,omitempty" name:"SdkHost"`

		// Subscription object type. 0: full instance subscription, 1: DDL data subscription, 2: DML structure subscription, 3: DDL data subscription + DML structure subscription
		SubscribeObjectType *int64 `json:"SubscribeObjectType,omitempty" name:"SubscribeObjectType"`

		// Subscription object, which is an empty array if `SubscribeObjectType` is 0
		SubscribeObjects []*SubscribeObject `json:"SubscribeObjects,omitempty" name:"SubscribeObjects" list`

		// Modification time
		ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

		// Region
		Region *string `json:"Region,omitempty" name:"Region"`

		// Tags of the subscription
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		Tags []*TagItem `json:"Tags,omitempty" name:"Tags" list`

		// Whether auto-renewal is enabled. 0: do not enable, 1: enable
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubscribeConfResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscribesRequest struct {
	*tchttp.BaseRequest

	// Data subscription instance ID
	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`

	// Data subscription instance name
	SubscribeName *string `json:"SubscribeName,omitempty" name:"SubscribeName"`

	// ID of bound database instance
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Data subscription instance channel ID
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// Billing mode filter. Default value: 1 (pay-as-you-go)
	PayType *string `json:"PayType,omitempty" name:"PayType"`

	// Subscribed database product, such as MySQL
	Product *string `json:"Product,omitempty" name:"Product"`

	// Data subscription instance status. Valid values: creating, normal, isolating, isolated, offlining
	Status []*string `json:"Status,omitempty" name:"Status" list`

	// Data subscription instance configuration status. Valid values: unconfigure, configuring, configured
	SubsStatus []*string `json:"SubsStatus,omitempty" name:"SubsStatus" list`

	// Starting offset of returned results
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results to be returned at a time
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Sorting order. Valid values: DESC, ASC. Default value: DESC, indicating descending by creation time
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`

	// Tag filtering condition
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters" list`

	// Subscription instance edition. `txdts`: legacy data subscription; `kafka`: data subscription in Kafka edition
	SubscribeVersion *string `json:"SubscribeVersion,omitempty" name:"SubscribeVersion"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubscribesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubscribeId")
	delete(f, "SubscribeName")
	delete(f, "InstanceId")
	delete(f, "ChannelId")
	delete(f, "PayType")
	delete(f, "Product")
	delete(f, "Status")
	delete(f, "SubsStatus")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderDirection")
	delete(f, "TagFilters")
	delete(f, "SubscribeVersion")
	if len(f) > 0 {
		return errors.New("DescribeSubscribesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscribesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible instances.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Information list of data subscription instances
		Items []*SubscribeInfo `json:"Items,omitempty" name:"Items" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubscribesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSyncCheckJobRequest struct {
	*tchttp.BaseRequest

	// ID of the disaster recovery sync task to be queried
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSyncCheckJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return errors.New("DescribeSyncCheckJobRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSyncCheckJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task check status: starting, running, finished
		Status *string `json:"Status,omitempty" name:"Status"`

		// Code of the task check result
		ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`

		// Prompt message
		ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`

		// Description of a task execution step
		StepInfo []*SyncCheckStepInfo `json:"StepInfo,omitempty" name:"StepInfo" list`

		// Check flag. 0: checking; 1: successfully checked
		CheckFlag *int64 `json:"CheckFlag,omitempty" name:"CheckFlag"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSyncCheckJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSyncJobsRequest struct {
	*tchttp.BaseRequest

	// Disaster recovery sync task ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// Disaster recovery sync task name
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// Sort by field. Value range: JobId, Status, JobName, CreateTime
	Order *string `json:"Order,omitempty" name:"Order"`

	// Sorting order. Value range: ASC (ascending), DESC (descending)
	OrderSeq *string `json:"OrderSeq,omitempty" name:"OrderSeq"`

	// Offset. Default value: 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of the returned instances. Value range: [1, 100]. Default value: 20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSyncJobsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "JobName")
	delete(f, "Order")
	delete(f, "OrderSeq")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return errors.New("DescribeSyncJobsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSyncJobsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of tasks
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Array of task details
		JobList []*SyncJobInfo `json:"JobList,omitempty" name:"JobList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSyncJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DstInfo struct {

	// Target instance ID, such as cdb-jd92ijd8
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Target instance region, such as ap-guangzhou
	Region *string `json:"Region,omitempty" name:"Region"`

	// Target instance VIP, which has been disused and does not need to be entered
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Target instance Vport, which has been disused and does not need to be entered
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Only valid for MySQL currently. For instance-level migration, the value range is: 1 (read-only), 0 (read/write)
	ReadOnly *int64 `json:"ReadOnly,omitempty" name:"ReadOnly"`

	// Target database account
	User *string `json:"User,omitempty" name:"User"`

	// Target database password
	Password *string `json:"Password,omitempty" name:"Password"`
}

type ErrorInfo struct {

	// Specific error log, including error code and error message
	ErrorLog *string `json:"ErrorLog,omitempty" name:"ErrorLog"`

	// Help document URL corresponding to error
	HelpDoc *string `json:"HelpDoc,omitempty" name:"HelpDoc"`
}

type IsolateSubscribeRequest struct {
	*tchttp.BaseRequest

	// Subscription instance ID
	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateSubscribeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubscribeId")
	if len(f) > 0 {
		return errors.New("IsolateSubscribeRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type IsolateSubscribeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MigrateDetailInfo struct {

	// Total number of steps
	StepAll *int64 `json:"StepAll,omitempty" name:"StepAll"`

	// Current step
	StepNow *int64 `json:"StepNow,omitempty" name:"StepNow"`

	// Overall progress, such as "10"
	Progress *string `json:"Progress,omitempty" name:"Progress"`

	// Progress of current step, such as "1"
	CurrentStepProgress *string `json:"CurrentStepProgress,omitempty" name:"CurrentStepProgress"`

	// Master/slave lag in MB, which is valid during incremental sync and currently supported by TencentDB for Redis and MySQL
	MasterSlaveDistance *int64 `json:"MasterSlaveDistance,omitempty" name:"MasterSlaveDistance"`

	// Master/slave lag in seconds, which is valid during incremental sync and currently supported by TencentDB for MySQL
	SecondsBehindMaster *int64 `json:"SecondsBehindMaster,omitempty" name:"SecondsBehindMaster"`

	// Step information
	StepInfo []*MigrateStepDetailInfo `json:"StepInfo,omitempty" name:"StepInfo" list`
}

type MigrateJobInfo struct {

	// Data migration task ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// Data migration task name
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// Migration task configuration options
	MigrateOption *MigrateOption `json:"MigrateOption,omitempty" name:"MigrateOption"`

	// Source instance database type: MySQL, Redis, MongoDB, PostgreSQL, MariaDB, Percona
	SrcDatabaseType *string `json:"SrcDatabaseType,omitempty" name:"SrcDatabaseType"`

	// Source instance access type. Value range: extranet (public network), cvm (CVM-created instance), dcg (Direct Connect-enabled instance), vpncloud (Tencent Cloud VPN-enabled instance), cdb (TencentDB instance), ccn (CCN instances)
	SrcAccessType *string `json:"SrcAccessType,omitempty" name:"SrcAccessType"`

	// Source instance information, which is correlated with the migration task type
	SrcInfo *SrcInfo `json:"SrcInfo,omitempty" name:"SrcInfo"`

	// Target instance access type: MySQL, Redis, MongoDB, PostgreSQL, MariaDB, Percona
	DstDatabaseType *string `json:"DstDatabaseType,omitempty" name:"DstDatabaseType"`

	// Target instance access type, which currently only supports cdb (TencentDB instance)
	DstAccessType *string `json:"DstAccessType,omitempty" name:"DstAccessType"`

	// Target instance information
	DstInfo *DstInfo `json:"DstInfo,omitempty" name:"DstInfo"`

	// Information of the source table to be migrated. If the entire instance is to be migrated, this field should be []
	DatabaseInfo *string `json:"DatabaseInfo,omitempty" name:"DatabaseInfo"`

	// Task creation/submission time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Task start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Task end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Task status. Value range: 1 (Creating), 3 (Checking), 4 (CheckPass), 5 (CheckNotPass), 7 (Running), 8 (ReadyComplete), 9 (Success), 10 (Failed), 11 (Stopping), 12 (Completing)
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Task details
	Detail *MigrateDetailInfo `json:"Detail,omitempty" name:"Detail"`

	// Prompt message for task error, which is not null or empty when an error occurs with the task
	ErrorInfo []*ErrorInfo `json:"ErrorInfo,omitempty" name:"ErrorInfo" list`
}

type MigrateOption struct {

	// Task operation mode. Value range: 1 (immediate execution), 2 (scheduled execution)
	RunMode *int64 `json:"RunMode,omitempty" name:"RunMode"`

	// Expected execution time in the format of yyyy-mm-dd hh:mm:ss. If runMode=2, this field is required
	ExpectTime *string `json:"ExpectTime,omitempty" name:"ExpectTime"`

	// Data migration type. Value range: 1 (structural migration), 2 (full migration), 3 (full + incremental migration)
	MigrateType *int64 `json:"MigrateType,omitempty" name:"MigrateType"`

	// Migration subject. 1: entire instance; 2: specified table
	MigrateObject *int64 `json:"MigrateObject,omitempty" name:"MigrateObject"`

	// Parameter of spot data consistency check. 1: not configured; 2: full check; 3: spot check; 4: check inconsistent tables only; 5: no check
	ConsistencyType *int64 `json:"ConsistencyType,omitempty" name:"ConsistencyType"`

	// Whether to overwrite the target database with the root account of the source database. Value range: 0 (no), 1 (yes). This value should be 0 for table or structural migration
	IsOverrideRoot *int64 `json:"IsOverrideRoot,omitempty" name:"IsOverrideRoot"`

	// Additional parameters for different databases, which are described in JSON format. 
	// The following parameters can be defined for Redis: 
	// { 
	// 	"ClientOutputBufferHardLimit":512, 	Hard capacity limit of slave buffer (MB) 
	// 	"ClientOutputBufferSoftLimit":512, 	Soft capacity limit of slave buffer (MB) 
	// 	"ClientOutputBufferPersistTime":60, Soft limit duration of slave buffer (s) 
	// 	"ReplBacklogSize":512, 	Circular buffer capacity limit (MB) 
	// 	"ReplTimeout":120, 		Replication timeout period (s) 
	// }
	// The following parameters can be defined for MongoDB: 
	// {
	// 	'SrcAuthDatabase':'admin', 
	// 	'SrcAuthFlag': "1", 
	// 	'SrcAuthMechanism':"SCRAM-SHA-1"
	// }
	// MySQL currently does not support configuring additional parameters.
	ExternParams *string `json:"ExternParams,omitempty" name:"ExternParams"`

	// Only used for "spot data consistency check". It is required if ConsistencyType is spot check
	ConsistencyParams *ConsistencyParams `json:"ConsistencyParams,omitempty" name:"ConsistencyParams"`
}

type MigrateStepDetailInfo struct {

	// Step number
	StepNo *int64 `json:"StepNo,omitempty" name:"StepNo"`

	// Step name
	StepName *string `json:"StepName,omitempty" name:"StepName"`

	// Step ID
	StepId *string `json:"StepId,omitempty" name:"StepId"`

	// Step status. Value range: 0 (default), 1 (succeeded), 2 (failed), 3 (in progress), 4 (not started)
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Start time of current step in the format of `yyyy-mm-dd hh:mm:ss`. This field is meaningless if it does not exist or is empty
	// Note: this field may return null, indicating that no valid values can be obtained.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

type ModifyMigrateJobRequest struct {
	*tchttp.BaseRequest

	// ID of the data migration task to be modified
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// Data migration task name
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// Migration task configuration options
	MigrateOption *MigrateOption `json:"MigrateOption,omitempty" name:"MigrateOption"`

	// Source instance access type. Valid values: extranet (public network), cvm (CVM-based self-created instance), dcg (Direct Connect-enabled instance), vpncloud (Tencent Cloud VPN-enabled instance), cdb (TencentDB instance)
	SrcAccessType *string `json:"SrcAccessType,omitempty" name:"SrcAccessType"`

	// Source instance information, which is correlated with the migration task type
	SrcInfo *SrcInfo `json:"SrcInfo,omitempty" name:"SrcInfo"`

	// Target instance access type. Valid values: extranet (public network), cvm (CVM-based self-created instance), dcg (Direct Connect-enabled instance), vpncloud (Tencent Cloud VPN-enabled instance), cdb (TencentDB instance). Currently, only `cdb` is supported
	DstAccessType *string `json:"DstAccessType,omitempty" name:"DstAccessType"`

	// Target instance information. The region where the target instance is located cannot be modified.
	DstInfo *DstInfo `json:"DstInfo,omitempty" name:"DstInfo"`

	// When migrating the specified table, you need to set the information of the source database table to be migrated, which should be described in JSON string format. Below are examples.
	// 
	// For databases with a database-table structure:
	// [{"Database":"db1","Table":["table1","table2"]},{"Database":"db2"}]
	// For databases with a database-schema-table structure:
	// [{"Database":"db1","Schema":"s1","Table":["table1","table2"]},{"Database":"db1","Schema":"s2","Table":["table1","table2"]},{"Database":"db2","Schema":"s1","Table":["table1","table2"]},{"Database":"db3"},{"Database":"db4","Schema":"s1"}]
	// 
	// This field does not need to be set when the entire instance is to be migrated
	DatabaseInfo *string `json:"DatabaseInfo,omitempty" name:"DatabaseInfo"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMigrateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "JobName")
	delete(f, "MigrateOption")
	delete(f, "SrcAccessType")
	delete(f, "SrcInfo")
	delete(f, "DstAccessType")
	delete(f, "DstInfo")
	delete(f, "DatabaseInfo")
	if len(f) > 0 {
		return errors.New("ModifyMigrateJobRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMigrateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySubscribeConsumeTimeRequest struct {
	*tchttp.BaseRequest

	// Data subscription instance ID
	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`

	// Consumption starting time point in the format of `Y-m-d h:m:s`, i.e., the starting time point for data subscription. Value range: within the last 24 hours
	ConsumeStartTime *string `json:"ConsumeStartTime,omitempty" name:"ConsumeStartTime"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubscribeConsumeTimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubscribeId")
	delete(f, "ConsumeStartTime")
	if len(f) > 0 {
		return errors.New("ModifySubscribeConsumeTimeRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifySubscribeConsumeTimeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubscribeConsumeTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySubscribeNameRequest struct {
	*tchttp.BaseRequest

	// Data subscription instance ID
	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`

	// Data subscription instance name. Length limit: [1,60]
	SubscribeName *string `json:"SubscribeName,omitempty" name:"SubscribeName"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubscribeNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubscribeId")
	delete(f, "SubscribeName")
	if len(f) > 0 {
		return errors.New("ModifySubscribeNameRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifySubscribeNameResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubscribeNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySubscribeObjectsRequest struct {
	*tchttp.BaseRequest

	// Data subscription instance ID
	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`

	// Data subscription type. Valid values: 0 (full instance subscription), 1 (data subscription), 2 (structure subscription), 3 (data subscription + structure subscription)
	SubscribeObjectType *int64 `json:"SubscribeObjectType,omitempty" name:"SubscribeObjectType"`

	// Information of subscribed table
	Objects []*SubscribeObject `json:"Objects,omitempty" name:"Objects" list`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubscribeObjectsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubscribeId")
	delete(f, "SubscribeObjectType")
	delete(f, "Objects")
	if len(f) > 0 {
		return errors.New("ModifySubscribeObjectsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifySubscribeObjectsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task ID
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubscribeObjectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySubscribeVipVportRequest struct {
	*tchttp.BaseRequest

	// Data subscription instance ID
	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`

	// Specified destination subnet. If this parameter is passed in, `DstIp` must be in the destination subnet
	DstUniqSubnetId *string `json:"DstUniqSubnetId,omitempty" name:"DstUniqSubnetId"`

	// Target IP. Either this field or `DstPort` must be passed in
	DstIp *string `json:"DstIp,omitempty" name:"DstIp"`

	// Target port. Value range: [1025-65535]
	DstPort *int64 `json:"DstPort,omitempty" name:"DstPort"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubscribeVipVportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubscribeId")
	delete(f, "DstUniqSubnetId")
	delete(f, "DstIp")
	delete(f, "DstPort")
	if len(f) > 0 {
		return errors.New("ModifySubscribeVipVportRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifySubscribeVipVportResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubscribeVipVportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySyncJobRequest struct {
	*tchttp.BaseRequest

	// ID of the disaster recovery sync task to be modified
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// Name of the disaster recovery sync task
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// Configuration options of a disaster recovery sync task
	SyncOption *SyncOption `json:"SyncOption,omitempty" name:"SyncOption"`

	// When syncing the specified table, you need to set the information of the source table to be synced, which should be described in JSON string format. Below are examples.
	// For databases with a database-table structure:
	// [{"Database":"db1","Table":["table1","table2"]},{"Database":"db2"}]
	DatabaseInfo *string `json:"DatabaseInfo,omitempty" name:"DatabaseInfo"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySyncJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "JobName")
	delete(f, "SyncOption")
	delete(f, "DatabaseInfo")
	if len(f) > 0 {
		return errors.New("ModifySyncJobRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifySyncJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OfflineIsolatedSubscribeRequest struct {
	*tchttp.BaseRequest

	// Data subscription instance ID
	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OfflineIsolatedSubscribeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubscribeId")
	if len(f) > 0 {
		return errors.New("OfflineIsolatedSubscribeRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type OfflineIsolatedSubscribeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OfflineIsolatedSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetSubscribeRequest struct {
	*tchttp.BaseRequest

	// Data subscription instance ID
	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetSubscribeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubscribeId")
	if len(f) > 0 {
		return errors.New("ResetSubscribeRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ResetSubscribeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SrcInfo struct {

	// Alibaba Cloud AccessKey, which is applicable if the source database is an Alibaba Cloud ApsaraDB for RDS 5.6 instance
	AccessKey *string `json:"AccessKey,omitempty" name:"AccessKey"`

	// Instance IP address
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Instance port
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Instance username
	User *string `json:"User,omitempty" name:"User"`

	// Instance password
	Password *string `json:"Password,omitempty" name:"Password"`

	// Alibaba Cloud ApsaraDB for RDS instance ID, which is applicable if the source database is an Alibaba Cloud ApsaraDB for RDS 5.6/5.7 instance
	RdsInstanceId *string `json:"RdsInstanceId,omitempty" name:"RdsInstanceId"`

	// Short CVM instance ID in the format of `ins-olgl39y8`. It is the same as the instance ID displayed on the CVM Console page. For CVM-based self-created instances, this field needs to be passed in
	CvmInstanceId *string `json:"CvmInstanceId,omitempty" name:"CvmInstanceId"`

	// Direct Connect gateway ID in the format of dcg-0rxtqqxb
	UniqDcgId *string `json:"UniqDcgId,omitempty" name:"UniqDcgId"`

	// VPC ID in the format of vpc-92jblxto
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// VPC Subnet ID in the format of subnet-3paxmkdz
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// VPN gateway ID in the format of vpngw-9ghexg7q
	UniqVpnGwId *string `json:"UniqVpnGwId,omitempty" name:"UniqVpnGwId"`

	// Database instance ID in the format of cdb-powiqx8q
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Region name, such as ap-guangzhou
	Region *string `json:"Region,omitempty" name:"Region"`

	// For Alibaba Cloud ApsaraDB for RDS instances, enter "aliyun"; otherwise, enter "others"
	Supplier *string `json:"Supplier,omitempty" name:"Supplier"`

	// CCN instance ID, such as ccn-afp6kltc
	// Note: This field may return null, indicating that no valid values can be obtained.
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`

	// Database version. This parameter is valid only when the instance is an RDS instance. Value: 5.6 or 5.7. Default value: 5.6
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`
}

type StartMigrateJobRequest struct {
	*tchttp.BaseRequest

	// Data migration task ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartMigrateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return errors.New("StartMigrateJobRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type StartMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartMigrateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartSyncJobRequest struct {
	*tchttp.BaseRequest

	// Disaster recovery sync task ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartSyncJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return errors.New("StartSyncJobRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type StartSyncJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartSyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopMigrateJobRequest struct {
	*tchttp.BaseRequest

	// Data migration task ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopMigrateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return errors.New("StopMigrateJobRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type StopMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopMigrateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubscribeInfo struct {

	// Data subscription instance ID
	SubscribeId *string `json:"SubscribeId,omitempty" name:"SubscribeId"`

	// Data subscription instance name
	SubscribeName *string `json:"SubscribeName,omitempty" name:"SubscribeName"`

	// ID of channel bound to data subscription instance
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// Name of product bound to data subscription instance
	Product *string `json:"Product,omitempty" name:"Product"`

	// ID of database instance bound to data subscription instance
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Status of database instance bound to data subscription instance
	InstanceStatus *string `json:"InstanceStatus,omitempty" name:"InstanceStatus"`

	// Data subscription instance configuration status. Valid values: unconfigure, configuring, configured
	SubsStatus *string `json:"SubsStatus,omitempty" name:"SubsStatus"`

	// Last modified time
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Isolation time
	IsolateTime *string `json:"IsolateTime,omitempty" name:"IsolateTime"`

	// Expiration time
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// Deactivation time
	OfflineTime *string `json:"OfflineTime,omitempty" name:"OfflineTime"`

	// Last modified consumption starting time point. If it has never been modified, this field is 0
	ConsumeStartTime *string `json:"ConsumeStartTime,omitempty" name:"ConsumeStartTime"`

	// Data subscription instance region
	Region *string `json:"Region,omitempty" name:"Region"`

	// Billing mode. 1: pay-as-you-go
	PayType *int64 `json:"PayType,omitempty" name:"PayType"`

	// Data subscription instance VIP
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Data subscription instance Vport
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// Unique ID of the VPC where the data subscription instance VIP resides
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// Unique ID of the subnet where the data subscription instance VIP resides
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// Data subscription instance status. Valid values: creating, normal, isolating, isolated, offlining, offline
	Status *string `json:"Status,omitempty" name:"Status"`

	// Timestamp of the last message confirmed by the SDK. If the SDK keeps consuming, this field can also be used as the current consumption time point of the SDK
	SdkConsumedTime *string `json:"SdkConsumedTime,omitempty" name:"SdkConsumedTime"`

	// Tag
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Tags []*TagItem `json:"Tags,omitempty" name:"Tags" list`

	// Whether auto-renewal is enabled. 0: do not enable; 1: enable
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// Subscription instance edition. ·`txdts`: legacy data subscription; `kafka`: data subscription in Kafka edition
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SubscribeVersion *string `json:"SubscribeVersion,omitempty" name:"SubscribeVersion"`
}

type SubscribeObject struct {

	// Data subscription object type. 0: database, 1: database table
	// Note: this field may return null, indicating that no valid values can be obtained.
	ObjectsType *int64 `json:"ObjectsType,omitempty" name:"ObjectsType"`

	// Name of subscribed database
	// Note: this field may return null, indicating that no valid values can be obtained.
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// Array of table names in subscribed database
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableNames []*string `json:"TableNames,omitempty" name:"TableNames" list`
}

type SubscribeRegionConf struct {

	// Region name, such as Guangzhou
	// Note: this field may return null, indicating that no valid values can be obtained.
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// Region ID, such as ap-guangzhou
	// Note: this field may return null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitempty" name:"Region"`

	// Region name, such as South China
	// Note: this field may return null, indicating that no valid values can be obtained.
	Area *string `json:"Area,omitempty" name:"Area"`

	// Whether it is the default region. 0: no, 1: yes
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsDefaultRegion *int64 `json:"IsDefaultRegion,omitempty" name:"IsDefaultRegion"`

	// Purchasable status of current region. 1: normal, 2: beta test, 3: not purchasable
	// Note: this field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type SwitchDrToMasterRequest struct {
	*tchttp.BaseRequest

	// Disaster recovery instance information
	DstInfo *SyncInstanceInfo `json:"DstInfo,omitempty" name:"DstInfo"`

	// Database type (such as MySQL)
	DatabaseType *string `json:"DatabaseType,omitempty" name:"DatabaseType"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchDrToMasterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DstInfo")
	delete(f, "DatabaseType")
	if len(f) > 0 {
		return errors.New("SwitchDrToMasterRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SwitchDrToMasterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Backend async task request ID
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchDrToMasterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncCheckStepInfo struct {

	// Step number
	StepNo *uint64 `json:"StepNo,omitempty" name:"StepNo"`

	// Step name
	StepName *string `json:"StepName,omitempty" name:"StepName"`

	// Code of the step execution result
	StepCode *int64 `json:"StepCode,omitempty" name:"StepCode"`

	// Message about the step execution result
	StepMessage *string `json:"StepMessage,omitempty" name:"StepMessage"`
}

type SyncDetailInfo struct {

	// Total number of steps
	StepAll *int64 `json:"StepAll,omitempty" name:"StepAll"`

	// Current step
	StepNow *int64 `json:"StepNow,omitempty" name:"StepNow"`

	// Overall progress
	Progress *string `json:"Progress,omitempty" name:"Progress"`

	// Progress of the current step
	CurrentStepProgress *string `json:"CurrentStepProgress,omitempty" name:"CurrentStepProgress"`

	// Master/slave delay in MB
	MasterSlaveDistance *int64 `json:"MasterSlaveDistance,omitempty" name:"MasterSlaveDistance"`

	// Master/slave delay in seconds
	SecondsBehindMaster *int64 `json:"SecondsBehindMaster,omitempty" name:"SecondsBehindMaster"`

	// Step information
	StepInfo []*SyncStepDetailInfo `json:"StepInfo,omitempty" name:"StepInfo" list`
}

type SyncInstanceInfo struct {

	// Region name, such as ap-guangzhou
	Region *string `json:"Region,omitempty" name:"Region"`

	// Short instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type SyncJobInfo struct {

	// Disaster recovery task ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// Disaster recovery task name
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// Task sync
	SyncOption *SyncOption `json:"SyncOption,omitempty" name:"SyncOption"`

	// Source access type
	SrcAccessType *string `json:"SrcAccessType,omitempty" name:"SrcAccessType"`

	// Source data type
	SrcDatabaseType *string `json:"SrcDatabaseType,omitempty" name:"SrcDatabaseType"`

	// Source instance information
	SrcInfo *SyncInstanceInfo `json:"SrcInfo,omitempty" name:"SrcInfo"`

	// Disaster recovery access type
	DstAccessType *string `json:"DstAccessType,omitempty" name:"DstAccessType"`

	// Disaster recovery data type
	DstDatabaseType *string `json:"DstDatabaseType,omitempty" name:"DstDatabaseType"`

	// Disaster recovery instance information
	DstInfo *SyncInstanceInfo `json:"DstInfo,omitempty" name:"DstInfo"`

	// Task information
	Detail *SyncDetailInfo `json:"Detail,omitempty" name:"Detail"`

	// Task status
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Table to be migrated
	DatabaseInfo *string `json:"DatabaseInfo,omitempty" name:"DatabaseInfo"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type SyncOption struct {

	// Sync object. 1: entire instance; 2: specified table
	SyncObject *uint64 `json:"SyncObject,omitempty" name:"SyncObject"`

	// Sync start configuration. 1: start immediately
	RunMode *uint64 `json:"RunMode,omitempty" name:"RunMode"`

	// Sync mode. 3: full + incremental sync
	SyncType *uint64 `json:"SyncType,omitempty" name:"SyncType"`

	// Data consistency check. 1: no configuration required
	ConsistencyType *uint64 `json:"ConsistencyType,omitempty" name:"ConsistencyType"`
}

type SyncStepDetailInfo struct {

	// Step number
	StepNo *uint64 `json:"StepNo,omitempty" name:"StepNo"`

	// Step name
	StepName *string `json:"StepName,omitempty" name:"StepName"`

	// Whether it can be stopped
	CanStop *int64 `json:"CanStop,omitempty" name:"CanStop"`

	// Step ID
	StepId *int64 `json:"StepId,omitempty" name:"StepId"`
}

type TagFilter struct {

	// Tag key value
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value
	TagValue []*string `json:"TagValue,omitempty" name:"TagValue" list`
}

type TagItem struct {

	// Tag key value
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}
