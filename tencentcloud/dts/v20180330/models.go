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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

// Predefined struct for user
type ActivateSubscribeRequestParams struct {
	// Subscription instance ID.
	SubscribeId *string `json:"SubscribeId,omitnil" name:"SubscribeId"`

	// Database Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Data subscription type. 0: full instance subscription, 1: data subscription, 2: structure subscription, 3: data subscription and structure subscription
	SubscribeObjectType *int64 `json:"SubscribeObjectType,omitnil" name:"SubscribeObjectType"`

	// Subscription object
	Objects *SubscribeObject `json:"Objects,omitnil" name:"Objects"`

	// Subnet of data subscription service, which is the subnet of the database instance by default.
	UniqSubnetId *string `json:"UniqSubnetId,omitnil" name:"UniqSubnetId"`

	// Subscription service port. Default value: 7507
	Vport *int64 `json:"Vport,omitnil" name:"Vport"`
}

type ActivateSubscribeRequest struct {
	*tchttp.BaseRequest
	
	// Subscription instance ID.
	SubscribeId *string `json:"SubscribeId,omitnil" name:"SubscribeId"`

	// Database Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Data subscription type. 0: full instance subscription, 1: data subscription, 2: structure subscription, 3: data subscription and structure subscription
	SubscribeObjectType *int64 `json:"SubscribeObjectType,omitnil" name:"SubscribeObjectType"`

	// Subscription object
	Objects *SubscribeObject `json:"Objects,omitnil" name:"Objects"`

	// Subnet of data subscription service, which is the subnet of the database instance by default.
	UniqSubnetId *string `json:"UniqSubnetId,omitnil" name:"UniqSubnetId"`

	// Subscription service port. Default value: 7507
	Vport *int64 `json:"Vport,omitnil" name:"Vport"`
}

func (r *ActivateSubscribeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ActivateSubscribeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ActivateSubscribeResponseParams struct {
	// Data subscription configuration task ID.
	AsyncRequestId *string `json:"AsyncRequestId,omitnil" name:"AsyncRequestId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ActivateSubscribeResponse struct {
	*tchttp.BaseResponse
	Response *ActivateSubscribeResponseParams `json:"Response"`
}

func (r *ActivateSubscribeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ActivateSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CompleteMigrateJobRequestParams struct {
	// Data migration task ID
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// The way to complete the task, which is supported only for legacy MySQL migration task. waitForSync: wait for the source-replica lag to become 0 before stopping; immediately: complete immediately without waiting for source-replica sync. Default value: waitForSync
	CompleteMode *string `json:"CompleteMode,omitnil" name:"CompleteMode"`
}

type CompleteMigrateJobRequest struct {
	*tchttp.BaseRequest
	
	// Data migration task ID
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// The way to complete the task, which is supported only for legacy MySQL migration task. waitForSync: wait for the source-replica lag to become 0 before stopping; immediately: complete immediately without waiting for source-replica sync. Default value: waitForSync
	CompleteMode *string `json:"CompleteMode,omitnil" name:"CompleteMode"`
}

func (r *CompleteMigrateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CompleteMigrateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "CompleteMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CompleteMigrateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CompleteMigrateJobResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CompleteMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *CompleteMigrateJobResponseParams `json:"Response"`
}

func (r *CompleteMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CompleteMigrateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConsistencyParams struct {
	// Data content check parameter, which refers to the proportion of the rows selected for data comparison in all the rows of the table. Value: an integer between 1 and 100.
	SelectRowsPerTable *int64 `json:"SelectRowsPerTable,omitnil" name:"SelectRowsPerTable"`

	// Data content check parameter, which refers to the proportion of the tables selected for data detection in all the tables. Value: an integer between 1 and 100.
	TablesSelectAll *int64 `json:"TablesSelectAll,omitnil" name:"TablesSelectAll"`

	// Data quantity check parameter, which checks whether the numbers of rows are identical. It refers to the proportion of the tables selected for quantity check in all the tables. Value: an integer between 1 and 100.
	TablesSelectCount *int64 `json:"TablesSelectCount,omitnil" name:"TablesSelectCount"`
}

// Predefined struct for user
type CreateMigrateCheckJobRequestParams struct {
	// Data migration task ID
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

type CreateMigrateCheckJobRequest struct {
	*tchttp.BaseRequest
	
	// Data migration task ID
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

func (r *CreateMigrateCheckJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMigrateCheckJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMigrateCheckJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMigrateCheckJobResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateMigrateCheckJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateMigrateCheckJobResponseParams `json:"Response"`
}

func (r *CreateMigrateCheckJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMigrateCheckJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMigrateJobRequestParams struct {
	// Data migration task name
	JobName *string `json:"JobName,omitnil" name:"JobName"`

	// Migration task configuration options
	MigrateOption *MigrateOption `json:"MigrateOption,omitnil" name:"MigrateOption"`

	// Source instance database type, which currently supports MySQL, Redis, MongoDB, PostgreSQL, MariaDB, Percona, and SQL Server. For more information on the supported types in a specific region, see the migration task creation page in the console.
	SrcDatabaseType *string `json:"SrcDatabaseType,omitnil" name:"SrcDatabaseType"`

	// Source instance access type. Valid values: extranet (public network), cvm (CVM-based self-created instance), dcg (Direct Connect-enabled instance), vpncloud (Tencent Cloud VPN-enabled instance), cdb (TencentDB instance), ccn (CCN instance)
	SrcAccessType *string `json:"SrcAccessType,omitnil" name:"SrcAccessType"`

	// Source instance information, which is correlated with the migration task type
	SrcInfo *SrcInfo `json:"SrcInfo,omitnil" name:"SrcInfo"`

	// Target instance access type, which currently supports MySQL, Redis, MongoDB, PostgreSQL, MariaDB, and Percona, SQL Server, and TDSQL-C for MySQL. For more information on the supported types in a specific region, see the migration task creation page in the console.
	DstDatabaseType *string `json:"DstDatabaseType,omitnil" name:"DstDatabaseType"`

	// Target instance access type, which currently only supports cdb (TencentDB instance)
	DstAccessType *string `json:"DstAccessType,omitnil" name:"DstAccessType"`

	// Destination instance information
	DstInfo *DstInfo `json:"DstInfo,omitnil" name:"DstInfo"`

	// Information of the source table to be migrated, which is described in JSON string format. It is required if MigrateOption.MigrateObject is 2 (migrating the specified table).
	// For databases with a database-table structure:
	// [{"Database":"db1","Table":["table1","table2"]},{"Database":"db2"}]
	// For databases with a database-schema-table structure:
	// [{"Database":"db1","Schema":"s1","Table":["table1","table2"]},{"Database":"db1","Schema":"s2","Table":["table1","table2"]},{"Database":"db2","Schema":"s1","Table":["table1","table2"]},{"Database":"db3"},{"Database":"db4","Schema":"s1"}]
	DatabaseInfo *string `json:"DatabaseInfo,omitnil" name:"DatabaseInfo"`

	// Tag of the instance to be migrated.
	Tags []*TagItem `json:"Tags,omitnil" name:"Tags"`

	// Source instance type. `simple`: Primary/Secondary node; `cluster`: Cluster node. If this field is left empty, the value defaults to primary/secondary node.
	SrcNodeType *string `json:"SrcNodeType,omitnil" name:"SrcNodeType"`

	// Source instance information, which is correlated with the migration task type.
	SrcInfoMulti []*SrcInfo `json:"SrcInfoMulti,omitnil" name:"SrcInfoMulti"`
}

type CreateMigrateJobRequest struct {
	*tchttp.BaseRequest
	
	// Data migration task name
	JobName *string `json:"JobName,omitnil" name:"JobName"`

	// Migration task configuration options
	MigrateOption *MigrateOption `json:"MigrateOption,omitnil" name:"MigrateOption"`

	// Source instance database type, which currently supports MySQL, Redis, MongoDB, PostgreSQL, MariaDB, Percona, and SQL Server. For more information on the supported types in a specific region, see the migration task creation page in the console.
	SrcDatabaseType *string `json:"SrcDatabaseType,omitnil" name:"SrcDatabaseType"`

	// Source instance access type. Valid values: extranet (public network), cvm (CVM-based self-created instance), dcg (Direct Connect-enabled instance), vpncloud (Tencent Cloud VPN-enabled instance), cdb (TencentDB instance), ccn (CCN instance)
	SrcAccessType *string `json:"SrcAccessType,omitnil" name:"SrcAccessType"`

	// Source instance information, which is correlated with the migration task type
	SrcInfo *SrcInfo `json:"SrcInfo,omitnil" name:"SrcInfo"`

	// Target instance access type, which currently supports MySQL, Redis, MongoDB, PostgreSQL, MariaDB, and Percona, SQL Server, and TDSQL-C for MySQL. For more information on the supported types in a specific region, see the migration task creation page in the console.
	DstDatabaseType *string `json:"DstDatabaseType,omitnil" name:"DstDatabaseType"`

	// Target instance access type, which currently only supports cdb (TencentDB instance)
	DstAccessType *string `json:"DstAccessType,omitnil" name:"DstAccessType"`

	// Destination instance information
	DstInfo *DstInfo `json:"DstInfo,omitnil" name:"DstInfo"`

	// Information of the source table to be migrated, which is described in JSON string format. It is required if MigrateOption.MigrateObject is 2 (migrating the specified table).
	// For databases with a database-table structure:
	// [{"Database":"db1","Table":["table1","table2"]},{"Database":"db2"}]
	// For databases with a database-schema-table structure:
	// [{"Database":"db1","Schema":"s1","Table":["table1","table2"]},{"Database":"db1","Schema":"s2","Table":["table1","table2"]},{"Database":"db2","Schema":"s1","Table":["table1","table2"]},{"Database":"db3"},{"Database":"db4","Schema":"s1"}]
	DatabaseInfo *string `json:"DatabaseInfo,omitnil" name:"DatabaseInfo"`

	// Tag of the instance to be migrated.
	Tags []*TagItem `json:"Tags,omitnil" name:"Tags"`

	// Source instance type. `simple`: Primary/Secondary node; `cluster`: Cluster node. If this field is left empty, the value defaults to primary/secondary node.
	SrcNodeType *string `json:"SrcNodeType,omitnil" name:"SrcNodeType"`

	// Source instance information, which is correlated with the migration task type.
	SrcInfoMulti []*SrcInfo `json:"SrcInfoMulti,omitnil" name:"SrcInfoMulti"`
}

func (r *CreateMigrateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
	delete(f, "Tags")
	delete(f, "SrcNodeType")
	delete(f, "SrcInfoMulti")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMigrateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMigrateJobResponseParams struct {
	// Data migration task ID
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateMigrateJobResponseParams `json:"Response"`
}

func (r *CreateMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMigrateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSubscribeRequestParams struct {
	// Subscribed database type. Currently, MySQL is supported
	Product *string `json:"Product,omitnil" name:"Product"`

	// Instance billing mode, which is always 1 (hourly billing),
	PayType *int64 `json:"PayType,omitnil" name:"PayType"`

	// Purchase duration in months, which is required if `PayType` is 0. Maximum value: 120 (this field is not required of global site users and is better to be hidden)
	Duration *int64 `json:"Duration,omitnil" name:"Duration"`

	// Quantity. Default value: 1. Maximum value: 10
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// Whether to auto-renew. Default value: 0. This flag does not take effect for hourly billed instances (this field should be hidden from global site users)
	AutoRenew *int64 `json:"AutoRenew,omitnil" name:"AutoRenew"`

	// Instance resource tags
	Tags []*TagItem `json:"Tags,omitnil" name:"Tags"`

	// A custom instance name.
	Name *string `json:"Name,omitnil" name:"Name"`
}

type CreateSubscribeRequest struct {
	*tchttp.BaseRequest
	
	// Subscribed database type. Currently, MySQL is supported
	Product *string `json:"Product,omitnil" name:"Product"`

	// Instance billing mode, which is always 1 (hourly billing),
	PayType *int64 `json:"PayType,omitnil" name:"PayType"`

	// Purchase duration in months, which is required if `PayType` is 0. Maximum value: 120 (this field is not required of global site users and is better to be hidden)
	Duration *int64 `json:"Duration,omitnil" name:"Duration"`

	// Quantity. Default value: 1. Maximum value: 10
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// Whether to auto-renew. Default value: 0. This flag does not take effect for hourly billed instances (this field should be hidden from global site users)
	AutoRenew *int64 `json:"AutoRenew,omitnil" name:"AutoRenew"`

	// Instance resource tags
	Tags []*TagItem `json:"Tags,omitnil" name:"Tags"`

	// A custom instance name.
	Name *string `json:"Name,omitnil" name:"Name"`
}

func (r *CreateSubscribeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSubscribeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSubscribeResponseParams struct {
	// Data subscription instance ID array
	// Note: this field may return null, indicating that no valid values can be obtained.
	SubscribeIds []*string `json:"SubscribeIds,omitnil" name:"SubscribeIds"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateSubscribeResponse struct {
	*tchttp.BaseResponse
	Response *CreateSubscribeResponseParams `json:"Response"`
}

func (r *CreateSubscribeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMigrateJobRequestParams struct {
	// Data migration task ID
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

type DeleteMigrateJobRequest struct {
	*tchttp.BaseRequest
	
	// Data migration task ID
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

func (r *DeleteMigrateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMigrateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMigrateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMigrateJobResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMigrateJobResponseParams `json:"Response"`
}

func (r *DeleteMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMigrateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAsyncRequestInfoRequestParams struct {
	// Task ID
	AsyncRequestId *string `json:"AsyncRequestId,omitnil" name:"AsyncRequestId"`
}

type DescribeAsyncRequestInfoRequest struct {
	*tchttp.BaseRequest
	
	// Task ID
	AsyncRequestId *string `json:"AsyncRequestId,omitnil" name:"AsyncRequestId"`
}

func (r *DescribeAsyncRequestInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAsyncRequestInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AsyncRequestId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAsyncRequestInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAsyncRequestInfoResponseParams struct {
	// Task execution result information
	Info *string `json:"Info,omitnil" name:"Info"`

	// Task execution status. Valid values: success, failed, running
	Status *string `json:"Status,omitnil" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAsyncRequestInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAsyncRequestInfoResponseParams `json:"Response"`
}

func (r *DescribeAsyncRequestInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAsyncRequestInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigrateCheckJobRequestParams struct {
	// Data migration task ID
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

type DescribeMigrateCheckJobRequest struct {
	*tchttp.BaseRequest
	
	// Data migration task ID
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

func (r *DescribeMigrateCheckJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrateCheckJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMigrateCheckJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigrateCheckJobResponseParams struct {
	// Check task status: unavailable, starting, running, finished
	Status *string `json:"Status,omitnil" name:"Status"`

	// Task error code
	ErrorCode *int64 `json:"ErrorCode,omitnil" name:"ErrorCode"`

	// Task error message
	ErrorMessage *string `json:"ErrorMessage,omitnil" name:"ErrorMessage"`

	// Check task progress. For example, "30" means 30% completed
	Progress *string `json:"Progress,omitnil" name:"Progress"`

	// Whether the check succeeds. 0: no; 1: yes; 3: not checked
	CheckFlag *int64 `json:"CheckFlag,omitnil" name:"CheckFlag"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeMigrateCheckJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMigrateCheckJobResponseParams `json:"Response"`
}

func (r *DescribeMigrateCheckJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrateCheckJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigrateJobsRequestParams struct {
	// Data migration task ID
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// Data migration task name
	JobName *string `json:"JobName,omitnil" name:"JobName"`

	// Sort by field. Value range: JobId, Status, JobName, MigrateType, RunMode, CreateTime
	Order *string `json:"Order,omitnil" name:"Order"`

	// Sorting order. Value range: ASC (ascending), DESC (descending)
	OrderSeq *string `json:"OrderSeq,omitnil" name:"OrderSeq"`

	// Offset. Default value: 0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Number of the returned instances. Value range: [1, 100]. Default value: 20
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Tag filter.
	TagFilters []*TagFilter `json:"TagFilters,omitnil" name:"TagFilters"`
}

type DescribeMigrateJobsRequest struct {
	*tchttp.BaseRequest
	
	// Data migration task ID
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// Data migration task name
	JobName *string `json:"JobName,omitnil" name:"JobName"`

	// Sort by field. Value range: JobId, Status, JobName, MigrateType, RunMode, CreateTime
	Order *string `json:"Order,omitnil" name:"Order"`

	// Sorting order. Value range: ASC (ascending), DESC (descending)
	OrderSeq *string `json:"OrderSeq,omitnil" name:"OrderSeq"`

	// Offset. Default value: 0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Number of the returned instances. Value range: [1, 100]. Default value: 20
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Tag filter.
	TagFilters []*TagFilter `json:"TagFilters,omitnil" name:"TagFilters"`
}

func (r *DescribeMigrateJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
	delete(f, "TagFilters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMigrateJobsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigrateJobsResponseParams struct {
	// Number of tasks
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Array of task details
	JobList []*MigrateJobInfo `json:"JobList,omitnil" name:"JobList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeMigrateJobsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMigrateJobsResponseParams `json:"Response"`
}

func (r *DescribeMigrateJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrateJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegionConfRequestParams struct {

}

type DescribeRegionConfRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeRegionConfRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionConfRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRegionConfRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegionConfResponseParams struct {
	// Number of purchasable regions
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Purchasable region details
	Items []*SubscribeRegionConf `json:"Items,omitnil" name:"Items"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRegionConfResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRegionConfResponseParams `json:"Response"`
}

func (r *DescribeRegionConfResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionConfResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubscribeConfRequestParams struct {
	// Subscription instance ID
	SubscribeId *string `json:"SubscribeId,omitnil" name:"SubscribeId"`
}

type DescribeSubscribeConfRequest struct {
	*tchttp.BaseRequest
	
	// Subscription instance ID
	SubscribeId *string `json:"SubscribeId,omitnil" name:"SubscribeId"`
}

func (r *DescribeSubscribeConfRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubscribeConfRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubscribeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSubscribeConfRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubscribeConfResponseParams struct {
	// Subscription instance ID
	SubscribeId *string `json:"SubscribeId,omitnil" name:"SubscribeId"`

	// Subscription instance name
	SubscribeName *string `json:"SubscribeName,omitnil" name:"SubscribeName"`

	// Subscription channel
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// Subscribed database type
	Product *string `json:"Product,omitnil" name:"Product"`

	// Subscribed instance
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Subscribed instance status. Valid values: running, offline, isolate
	InstanceStatus *string `json:"InstanceStatus,omitnil" name:"InstanceStatus"`

	// Subscription instance status. Valid values: unconfigure, configuring, configured
	SubsStatus *string `json:"SubsStatus,omitnil" name:"SubsStatus"`

	// Subscription instance lifecycle status. Valid values: normal, isolating, isolated, offlining
	Status *string `json:"Status,omitnil" name:"Status"`

	// Subscription instance creation time
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// Subscription instance isolation time
	IsolateTime *string `json:"IsolateTime,omitnil" name:"IsolateTime"`

	// Subscription instance expiration time
	ExpireTime *string `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// Subscription instance deactivation time
	OfflineTime *string `json:"OfflineTime,omitnil" name:"OfflineTime"`

	// Consumption starting time point of subscription instance
	ConsumeStartTime *string `json:"ConsumeStartTime,omitnil" name:"ConsumeStartTime"`

	// Subscription instance billing mode. 1: hourly billing
	PayType *int64 `json:"PayType,omitnil" name:"PayType"`

	// Subscription channel VIP
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// Subscription channel port
	Vport *int64 `json:"Vport,omitnil" name:"Vport"`

	// Subscription channel `VpcId`
	UniqVpcId *string `json:"UniqVpcId,omitnil" name:"UniqVpcId"`

	// Subscription channel `SubnetId`
	UniqSubnetId *string `json:"UniqSubnetId,omitnil" name:"UniqSubnetId"`

	// Current SDK consumption time point
	SdkConsumedTime *string `json:"SdkConsumedTime,omitnil" name:"SdkConsumedTime"`

	// Subscription SDK IP address
	SdkHost *string `json:"SdkHost,omitnil" name:"SdkHost"`

	// Subscription object type. 0: full instance subscription, 1: DDL data subscription, 2: DML structure subscription, 3: DDL data subscription + DML structure subscription
	SubscribeObjectType *int64 `json:"SubscribeObjectType,omitnil" name:"SubscribeObjectType"`

	// Subscription object, which is an empty array if `SubscribeObjectType` is 0
	SubscribeObjects []*SubscribeObject `json:"SubscribeObjects,omitnil" name:"SubscribeObjects"`

	// Modification time
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`

	// Region
	Region *string `json:"Region,omitnil" name:"Region"`

	// Tags of the subscription
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Tags []*TagItem `json:"Tags,omitnil" name:"Tags"`

	// Whether auto-renewal is enabled. 0: do not enable, 1: enable
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// Data subscription edition. `txdts`: Legacy edition; `kafka`: Kafka edition.
	SubscribeVersion *string `json:"SubscribeVersion,omitnil" name:"SubscribeVersion"`

	// Error message.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Errors []*SubsErr `json:"Errors,omitnil" name:"Errors"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSubscribeConfResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSubscribeConfResponseParams `json:"Response"`
}

func (r *DescribeSubscribeConfResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubscribeConfResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubscribesRequestParams struct {
	// Data subscription instance ID
	SubscribeId *string `json:"SubscribeId,omitnil" name:"SubscribeId"`

	// Data subscription instance name
	SubscribeName *string `json:"SubscribeName,omitnil" name:"SubscribeName"`

	// ID of bound database instance
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Data subscription instance channel ID
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// Billing mode filter. Default value: 1 (pay-as-you-go)
	PayType *string `json:"PayType,omitnil" name:"PayType"`

	// Subscribed database product, such as MySQL
	Product *string `json:"Product,omitnil" name:"Product"`

	// Data subscription instance status. Valid values: creating, normal, isolating, isolated, offlining
	Status []*string `json:"Status,omitnil" name:"Status"`

	// Data subscription instance configuration status. Valid values: unconfigure, configuring, configured
	SubsStatus []*string `json:"SubsStatus,omitnil" name:"SubsStatus"`

	// Starting offset of returned results
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of results to be returned at a time
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Sorting order. Valid values: DESC, ASC. Default value: DESC, indicating descending by creation time
	OrderDirection *string `json:"OrderDirection,omitnil" name:"OrderDirection"`

	// Tag filtering condition
	TagFilters []*TagFilter `json:"TagFilters,omitnil" name:"TagFilters"`

	// Subscription instance edition. `txdts`: legacy data subscription; `kafka`: data subscription in Kafka edition
	SubscribeVersion *string `json:"SubscribeVersion,omitnil" name:"SubscribeVersion"`
}

type DescribeSubscribesRequest struct {
	*tchttp.BaseRequest
	
	// Data subscription instance ID
	SubscribeId *string `json:"SubscribeId,omitnil" name:"SubscribeId"`

	// Data subscription instance name
	SubscribeName *string `json:"SubscribeName,omitnil" name:"SubscribeName"`

	// ID of bound database instance
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Data subscription instance channel ID
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// Billing mode filter. Default value: 1 (pay-as-you-go)
	PayType *string `json:"PayType,omitnil" name:"PayType"`

	// Subscribed database product, such as MySQL
	Product *string `json:"Product,omitnil" name:"Product"`

	// Data subscription instance status. Valid values: creating, normal, isolating, isolated, offlining
	Status []*string `json:"Status,omitnil" name:"Status"`

	// Data subscription instance configuration status. Valid values: unconfigure, configuring, configured
	SubsStatus []*string `json:"SubsStatus,omitnil" name:"SubsStatus"`

	// Starting offset of returned results
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of results to be returned at a time
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Sorting order. Valid values: DESC, ASC. Default value: DESC, indicating descending by creation time
	OrderDirection *string `json:"OrderDirection,omitnil" name:"OrderDirection"`

	// Tag filtering condition
	TagFilters []*TagFilter `json:"TagFilters,omitnil" name:"TagFilters"`

	// Subscription instance edition. `txdts`: legacy data subscription; `kafka`: data subscription in Kafka edition
	SubscribeVersion *string `json:"SubscribeVersion,omitnil" name:"SubscribeVersion"`
}

func (r *DescribeSubscribesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSubscribesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubscribesResponseParams struct {
	// Number of eligible instances.
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Information list of data subscription instances
	Items []*SubscribeInfo `json:"Items,omitnil" name:"Items"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSubscribesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSubscribesResponseParams `json:"Response"`
}

func (r *DescribeSubscribesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubscribesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DstInfo struct {
	// Target instance ID, such as cdb-jd92ijd8
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Target instance region, such as ap-guangzhou
	Region *string `json:"Region,omitnil" name:"Region"`

	// Target instance VIP, which has been disused and does not need to be entered
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// Target instance Vport, which has been disused and does not need to be entered
	Port *int64 `json:"Port,omitnil" name:"Port"`

	// Only valid for MySQL currently. For instance-level migration, the value range is: 1 (read-only), 0 (read/write)
	ReadOnly *int64 `json:"ReadOnly,omitnil" name:"ReadOnly"`

	// Target database account
	User *string `json:"User,omitnil" name:"User"`

	// Target database password
	Password *string `json:"Password,omitnil" name:"Password"`
}

type ErrorInfo struct {
	// Specific error log, including error code and error message
	ErrorLog *string `json:"ErrorLog,omitnil" name:"ErrorLog"`

	// Help document URL corresponding to error
	HelpDoc *string `json:"HelpDoc,omitnil" name:"HelpDoc"`
}

// Predefined struct for user
type IsolateSubscribeRequestParams struct {
	// Subscription instance ID
	SubscribeId *string `json:"SubscribeId,omitnil" name:"SubscribeId"`
}

type IsolateSubscribeRequest struct {
	*tchttp.BaseRequest
	
	// Subscription instance ID
	SubscribeId *string `json:"SubscribeId,omitnil" name:"SubscribeId"`
}

func (r *IsolateSubscribeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateSubscribeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubscribeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsolateSubscribeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateSubscribeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type IsolateSubscribeResponse struct {
	*tchttp.BaseResponse
	Response *IsolateSubscribeResponseParams `json:"Response"`
}

func (r *IsolateSubscribeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MigrateDetailInfo struct {
	// Total number of steps
	StepAll *int64 `json:"StepAll,omitnil" name:"StepAll"`

	// Current step
	StepNow *int64 `json:"StepNow,omitnil" name:"StepNow"`

	// Overall progress, such as "10"
	Progress *string `json:"Progress,omitnil" name:"Progress"`

	// Progress of current step, such as "1"
	CurrentStepProgress *string `json:"CurrentStepProgress,omitnil" name:"CurrentStepProgress"`

	// Master/slave lag in MB, which is valid during incremental sync and currently supported by TencentDB for Redis and MySQL
	MasterSlaveDistance *int64 `json:"MasterSlaveDistance,omitnil" name:"MasterSlaveDistance"`

	// Master/slave lag in seconds, which is valid during incremental sync and currently supported by TencentDB for MySQL
	SecondsBehindMaster *int64 `json:"SecondsBehindMaster,omitnil" name:"SecondsBehindMaster"`

	// Step information
	StepInfo []*MigrateStepDetailInfo `json:"StepInfo,omitnil" name:"StepInfo"`
}

type MigrateJobInfo struct {
	// Data migration task ID
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// Data migration task name
	JobName *string `json:"JobName,omitnil" name:"JobName"`

	// Migration task configuration options
	MigrateOption *MigrateOption `json:"MigrateOption,omitnil" name:"MigrateOption"`

	// Source instance database type: MySQL, Redis, MongoDB, PostgreSQL, MariaDB, Percona
	SrcDatabaseType *string `json:"SrcDatabaseType,omitnil" name:"SrcDatabaseType"`

	// Source instance access type. Value range: extranet (public network), cvm (CVM-created instance), dcg (Direct Connect-enabled instance), vpncloud (Tencent Cloud VPN-enabled instance), cdb (TencentDB instance), ccn (CCN instances)
	SrcAccessType *string `json:"SrcAccessType,omitnil" name:"SrcAccessType"`

	// Source instance information, which is correlated with the migration task type
	SrcInfo *SrcInfo `json:"SrcInfo,omitnil" name:"SrcInfo"`

	// Target instance access type: MySQL, Redis, MongoDB, PostgreSQL, MariaDB, Percona
	DstDatabaseType *string `json:"DstDatabaseType,omitnil" name:"DstDatabaseType"`

	// Target instance access type, which currently only supports cdb (TencentDB instance)
	DstAccessType *string `json:"DstAccessType,omitnil" name:"DstAccessType"`

	// Target instance information
	DstInfo *DstInfo `json:"DstInfo,omitnil" name:"DstInfo"`

	// Information of the source table to be migrated. If the entire instance is to be migrated, this field should be []
	DatabaseInfo *string `json:"DatabaseInfo,omitnil" name:"DatabaseInfo"`

	// Task creation/submission time
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// Task start time
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// Task end time
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Task status. Value range: 1 (Creating), 3 (Checking), 4 (CheckPass), 5 (CheckNotPass), 7 (Running), 8 (ReadyComplete), 9 (Success), 10 (Failed), 11 (Stopping), 12 (Completing)
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// Task details
	Detail *MigrateDetailInfo `json:"Detail,omitnil" name:"Detail"`

	// Prompt message for task error, which is not null or empty when an error occurs with the task
	ErrorInfo []*ErrorInfo `json:"ErrorInfo,omitnil" name:"ErrorInfo"`

	// Tag
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Tags []*TagItem `json:"Tags,omitnil" name:"Tags"`

	// Information of the source instance, a cluster edition instance whose access type is not `cdb`.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	SrcInfoMulti []*SrcInfo `json:"SrcInfoMulti,omitnil" name:"SrcInfoMulti"`
}

type MigrateOption struct {
	// Task operation mode. Value range: 1 (immediate execution), 2 (scheduled execution)
	RunMode *int64 `json:"RunMode,omitnil" name:"RunMode"`

	// Expected execution time in the format of yyyy-mm-dd hh:mm:ss. If runMode=2, this field is required
	ExpectTime *string `json:"ExpectTime,omitnil" name:"ExpectTime"`

	// Data migration type. Value range: 1 (structural migration), 2 (full migration), 3 (full + incremental migration)
	MigrateType *int64 `json:"MigrateType,omitnil" name:"MigrateType"`

	// Migration subject. 1: entire instance; 2: specified table
	MigrateObject *int64 `json:"MigrateObject,omitnil" name:"MigrateObject"`

	// Parameter of spot data consistency check. 1: not configured; 2: full check; 3: spot check; 4: check inconsistent tables only; 5: no check
	ConsistencyType *int64 `json:"ConsistencyType,omitnil" name:"ConsistencyType"`

	// Whether to overwrite the target database with the root account of the source database. Value range: 0 (no), 1 (yes). This value should be 0 for table or structural migration
	IsOverrideRoot *int64 `json:"IsOverrideRoot,omitnil" name:"IsOverrideRoot"`

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
	ExternParams *string `json:"ExternParams,omitnil" name:"ExternParams"`

	// Only used for "spot data consistency check". It is required if ConsistencyType is spot check
	ConsistencyParams *ConsistencyParams `json:"ConsistencyParams,omitnil" name:"ConsistencyParams"`
}

type MigrateStepDetailInfo struct {
	// Step number
	StepNo *int64 `json:"StepNo,omitnil" name:"StepNo"`

	// Step name
	StepName *string `json:"StepName,omitnil" name:"StepName"`

	// Step ID
	StepId *string `json:"StepId,omitnil" name:"StepId"`

	// Step status. Value range: 0 (default), 1 (succeeded), 2 (failed), 3 (in progress), 4 (not started)
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// Start time of current step in the format of `yyyy-mm-dd hh:mm:ss`. This field is meaningless if it does not exist or is empty
	// Note: this field may return null, indicating that no valid values can be obtained.
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`
}

// Predefined struct for user
type ModifyMigrateJobRequestParams struct {
	// ID of the data migration task to be modified
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// Data migration task name
	JobName *string `json:"JobName,omitnil" name:"JobName"`

	// Migration task configuration options
	MigrateOption *MigrateOption `json:"MigrateOption,omitnil" name:"MigrateOption"`

	// Source instance access type. Valid values: extranet (public network), cvm (CVM-based self-created instance), dcg (Direct Connect-enabled instance), vpncloud (Tencent Cloud VPN-enabled instance), cdb (TencentDB instance)
	SrcAccessType *string `json:"SrcAccessType,omitnil" name:"SrcAccessType"`

	// Source instance information, which is correlated with the migration task type
	SrcInfo *SrcInfo `json:"SrcInfo,omitnil" name:"SrcInfo"`

	// Target instance access type. Valid values: extranet (public network), cvm (CVM-based self-created instance), dcg (Direct Connect-enabled instance), vpncloud (Tencent Cloud VPN-enabled instance), cdb (TencentDB instance). Currently, only `cdb` is supported
	DstAccessType *string `json:"DstAccessType,omitnil" name:"DstAccessType"`

	// Target instance information. The region where the target instance is located cannot be modified.
	DstInfo *DstInfo `json:"DstInfo,omitnil" name:"DstInfo"`

	// When migrating the specified table, you need to set the information of the source database table to be migrated, which should be described in JSON string format. Below are examples.
	// 
	// For databases with a database-table structure:
	// [{"Database":"db1","Table":["table1","table2"]},{"Database":"db2"}]
	// For databases with a database-schema-table structure:
	// [{"Database":"db1","Schema":"s1","Table":["table1","table2"]},{"Database":"db1","Schema":"s2","Table":["table1","table2"]},{"Database":"db2","Schema":"s1","Table":["table1","table2"]},{"Database":"db3"},{"Database":"db4","Schema":"s1"}]
	// 
	// This field does not need to be set when the entire instance is to be migrated
	DatabaseInfo *string `json:"DatabaseInfo,omitnil" name:"DatabaseInfo"`

	// Source instance type. `simple`: Primary/Secondary node; `cluster`: Cluster node. If this field is left empty, the value defaults to primary/secondary node.
	SrcNodeType *string `json:"SrcNodeType,omitnil" name:"SrcNodeType"`

	// Source instance information, which is correlated with the migration task type.
	SrcInfoMulti []*SrcInfo `json:"SrcInfoMulti,omitnil" name:"SrcInfoMulti"`
}

type ModifyMigrateJobRequest struct {
	*tchttp.BaseRequest
	
	// ID of the data migration task to be modified
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// Data migration task name
	JobName *string `json:"JobName,omitnil" name:"JobName"`

	// Migration task configuration options
	MigrateOption *MigrateOption `json:"MigrateOption,omitnil" name:"MigrateOption"`

	// Source instance access type. Valid values: extranet (public network), cvm (CVM-based self-created instance), dcg (Direct Connect-enabled instance), vpncloud (Tencent Cloud VPN-enabled instance), cdb (TencentDB instance)
	SrcAccessType *string `json:"SrcAccessType,omitnil" name:"SrcAccessType"`

	// Source instance information, which is correlated with the migration task type
	SrcInfo *SrcInfo `json:"SrcInfo,omitnil" name:"SrcInfo"`

	// Target instance access type. Valid values: extranet (public network), cvm (CVM-based self-created instance), dcg (Direct Connect-enabled instance), vpncloud (Tencent Cloud VPN-enabled instance), cdb (TencentDB instance). Currently, only `cdb` is supported
	DstAccessType *string `json:"DstAccessType,omitnil" name:"DstAccessType"`

	// Target instance information. The region where the target instance is located cannot be modified.
	DstInfo *DstInfo `json:"DstInfo,omitnil" name:"DstInfo"`

	// When migrating the specified table, you need to set the information of the source database table to be migrated, which should be described in JSON string format. Below are examples.
	// 
	// For databases with a database-table structure:
	// [{"Database":"db1","Table":["table1","table2"]},{"Database":"db2"}]
	// For databases with a database-schema-table structure:
	// [{"Database":"db1","Schema":"s1","Table":["table1","table2"]},{"Database":"db1","Schema":"s2","Table":["table1","table2"]},{"Database":"db2","Schema":"s1","Table":["table1","table2"]},{"Database":"db3"},{"Database":"db4","Schema":"s1"}]
	// 
	// This field does not need to be set when the entire instance is to be migrated
	DatabaseInfo *string `json:"DatabaseInfo,omitnil" name:"DatabaseInfo"`

	// Source instance type. `simple`: Primary/Secondary node; `cluster`: Cluster node. If this field is left empty, the value defaults to primary/secondary node.
	SrcNodeType *string `json:"SrcNodeType,omitnil" name:"SrcNodeType"`

	// Source instance information, which is correlated with the migration task type.
	SrcInfoMulti []*SrcInfo `json:"SrcInfoMulti,omitnil" name:"SrcInfoMulti"`
}

func (r *ModifyMigrateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
	delete(f, "SrcNodeType")
	delete(f, "SrcInfoMulti")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMigrateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMigrateJobResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMigrateJobResponseParams `json:"Response"`
}

func (r *ModifyMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMigrateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubscribeConsumeTimeRequestParams struct {
	// Data subscription instance ID
	SubscribeId *string `json:"SubscribeId,omitnil" name:"SubscribeId"`

	// Consumption starting time point in the format of `Y-m-d h:m:s`, i.e., the starting time point for data subscription. Value range: within the last 24 hours
	ConsumeStartTime *string `json:"ConsumeStartTime,omitnil" name:"ConsumeStartTime"`
}

type ModifySubscribeConsumeTimeRequest struct {
	*tchttp.BaseRequest
	
	// Data subscription instance ID
	SubscribeId *string `json:"SubscribeId,omitnil" name:"SubscribeId"`

	// Consumption starting time point in the format of `Y-m-d h:m:s`, i.e., the starting time point for data subscription. Value range: within the last 24 hours
	ConsumeStartTime *string `json:"ConsumeStartTime,omitnil" name:"ConsumeStartTime"`
}

func (r *ModifySubscribeConsumeTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubscribeConsumeTimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubscribeId")
	delete(f, "ConsumeStartTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySubscribeConsumeTimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubscribeConsumeTimeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifySubscribeConsumeTimeResponse struct {
	*tchttp.BaseResponse
	Response *ModifySubscribeConsumeTimeResponseParams `json:"Response"`
}

func (r *ModifySubscribeConsumeTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubscribeConsumeTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubscribeNameRequestParams struct {
	// Data subscription instance ID
	SubscribeId *string `json:"SubscribeId,omitnil" name:"SubscribeId"`

	// Data subscription instance name. Length limit: [1,60]
	SubscribeName *string `json:"SubscribeName,omitnil" name:"SubscribeName"`
}

type ModifySubscribeNameRequest struct {
	*tchttp.BaseRequest
	
	// Data subscription instance ID
	SubscribeId *string `json:"SubscribeId,omitnil" name:"SubscribeId"`

	// Data subscription instance name. Length limit: [1,60]
	SubscribeName *string `json:"SubscribeName,omitnil" name:"SubscribeName"`
}

func (r *ModifySubscribeNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubscribeNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubscribeId")
	delete(f, "SubscribeName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySubscribeNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubscribeNameResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifySubscribeNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifySubscribeNameResponseParams `json:"Response"`
}

func (r *ModifySubscribeNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubscribeNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubscribeObjectsRequestParams struct {
	// Data subscription instance ID
	SubscribeId *string `json:"SubscribeId,omitnil" name:"SubscribeId"`

	// Data subscription type. Valid values: 0 (full instance subscription), 1 (data subscription), 2 (structure subscription), 3 (data subscription + structure subscription)
	SubscribeObjectType *int64 `json:"SubscribeObjectType,omitnil" name:"SubscribeObjectType"`

	// Information of subscribed table
	Objects []*SubscribeObject `json:"Objects,omitnil" name:"Objects"`
}

type ModifySubscribeObjectsRequest struct {
	*tchttp.BaseRequest
	
	// Data subscription instance ID
	SubscribeId *string `json:"SubscribeId,omitnil" name:"SubscribeId"`

	// Data subscription type. Valid values: 0 (full instance subscription), 1 (data subscription), 2 (structure subscription), 3 (data subscription + structure subscription)
	SubscribeObjectType *int64 `json:"SubscribeObjectType,omitnil" name:"SubscribeObjectType"`

	// Information of subscribed table
	Objects []*SubscribeObject `json:"Objects,omitnil" name:"Objects"`
}

func (r *ModifySubscribeObjectsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySubscribeObjectsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubscribeObjectsResponseParams struct {
	// Async task ID
	AsyncRequestId *string `json:"AsyncRequestId,omitnil" name:"AsyncRequestId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifySubscribeObjectsResponse struct {
	*tchttp.BaseResponse
	Response *ModifySubscribeObjectsResponseParams `json:"Response"`
}

func (r *ModifySubscribeObjectsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubscribeObjectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubscribeVipVportRequestParams struct {
	// Data subscription instance ID
	SubscribeId *string `json:"SubscribeId,omitnil" name:"SubscribeId"`

	// Specified destination subnet. If this parameter is passed in, `DstIp` must be in the destination subnet
	DstUniqSubnetId *string `json:"DstUniqSubnetId,omitnil" name:"DstUniqSubnetId"`

	// Target IP. Either this field or `DstPort` must be passed in
	DstIp *string `json:"DstIp,omitnil" name:"DstIp"`

	// Target port. Value range: [1025-65535]
	DstPort *int64 `json:"DstPort,omitnil" name:"DstPort"`
}

type ModifySubscribeVipVportRequest struct {
	*tchttp.BaseRequest
	
	// Data subscription instance ID
	SubscribeId *string `json:"SubscribeId,omitnil" name:"SubscribeId"`

	// Specified destination subnet. If this parameter is passed in, `DstIp` must be in the destination subnet
	DstUniqSubnetId *string `json:"DstUniqSubnetId,omitnil" name:"DstUniqSubnetId"`

	// Target IP. Either this field or `DstPort` must be passed in
	DstIp *string `json:"DstIp,omitnil" name:"DstIp"`

	// Target port. Value range: [1025-65535]
	DstPort *int64 `json:"DstPort,omitnil" name:"DstPort"`
}

func (r *ModifySubscribeVipVportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySubscribeVipVportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubscribeVipVportResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifySubscribeVipVportResponse struct {
	*tchttp.BaseResponse
	Response *ModifySubscribeVipVportResponseParams `json:"Response"`
}

func (r *ModifySubscribeVipVportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubscribeVipVportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OfflineIsolatedSubscribeRequestParams struct {
	// Data subscription instance ID
	SubscribeId *string `json:"SubscribeId,omitnil" name:"SubscribeId"`
}

type OfflineIsolatedSubscribeRequest struct {
	*tchttp.BaseRequest
	
	// Data subscription instance ID
	SubscribeId *string `json:"SubscribeId,omitnil" name:"SubscribeId"`
}

func (r *OfflineIsolatedSubscribeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OfflineIsolatedSubscribeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubscribeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OfflineIsolatedSubscribeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OfflineIsolatedSubscribeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type OfflineIsolatedSubscribeResponse struct {
	*tchttp.BaseResponse
	Response *OfflineIsolatedSubscribeResponseParams `json:"Response"`
}

func (r *OfflineIsolatedSubscribeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OfflineIsolatedSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetSubscribeRequestParams struct {
	// Data subscription instance ID
	SubscribeId *string `json:"SubscribeId,omitnil" name:"SubscribeId"`
}

type ResetSubscribeRequest struct {
	*tchttp.BaseRequest
	
	// Data subscription instance ID
	SubscribeId *string `json:"SubscribeId,omitnil" name:"SubscribeId"`
}

func (r *ResetSubscribeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetSubscribeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubscribeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetSubscribeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetSubscribeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ResetSubscribeResponse struct {
	*tchttp.BaseResponse
	Response *ResetSubscribeResponseParams `json:"Response"`
}

func (r *ResetSubscribeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SrcInfo struct {
	// Alibaba Cloud AccessKey, which is applicable if the source database is an Alibaba Cloud ApsaraDB for RDS 5.6 instance
	AccessKey *string `json:"AccessKey,omitnil" name:"AccessKey"`

	// Instance IP address
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// Instance port
	Port *int64 `json:"Port,omitnil" name:"Port"`

	// Instance username
	User *string `json:"User,omitnil" name:"User"`

	// Instance password
	Password *string `json:"Password,omitnil" name:"Password"`

	// Alibaba Cloud ApsaraDB for RDS instance ID, which is applicable if the source database is an Alibaba Cloud ApsaraDB for RDS 5.6/5.7 instance
	RdsInstanceId *string `json:"RdsInstanceId,omitnil" name:"RdsInstanceId"`

	// Short CVM instance ID in the format of `ins-olgl39y8`. It is the same as the instance ID displayed on the CVM Console page. For CVM-based self-created instances, this field needs to be passed in
	CvmInstanceId *string `json:"CvmInstanceId,omitnil" name:"CvmInstanceId"`

	// Direct Connect gateway ID in the format of dcg-0rxtqqxb
	UniqDcgId *string `json:"UniqDcgId,omitnil" name:"UniqDcgId"`

	// VPC ID in the format of vpc-92jblxto
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// VPC Subnet ID in the format of subnet-3paxmkdz
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// VPN gateway ID in the format of vpngw-9ghexg7q
	UniqVpnGwId *string `json:"UniqVpnGwId,omitnil" name:"UniqVpnGwId"`

	// Database instance ID in the format of cdb-powiqx8q
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Region name, such as ap-guangzhou
	Region *string `json:"Region,omitnil" name:"Region"`

	// For Alibaba Cloud ApsaraDB for RDS instances, enter "aliyun"; otherwise, enter "others"
	Supplier *string `json:"Supplier,omitnil" name:"Supplier"`

	// CCN instance ID, such as ccn-afp6kltc
	// Note: This field may return null, indicating that no valid values can be obtained.
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`

	// Database version. This parameter is valid only when the instance is an RDS instance. Value: 5.6 or 5.7. Default value: 5.6
	EngineVersion *string `json:"EngineVersion,omitnil" name:"EngineVersion"`
}

// Predefined struct for user
type StartMigrateJobRequestParams struct {
	// Data migration task ID
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

type StartMigrateJobRequest struct {
	*tchttp.BaseRequest
	
	// Data migration task ID
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

func (r *StartMigrateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartMigrateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartMigrateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartMigrateJobResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type StartMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *StartMigrateJobResponseParams `json:"Response"`
}

func (r *StartMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartMigrateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopMigrateJobRequestParams struct {
	// Data migration task ID
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

type StopMigrateJobRequest struct {
	*tchttp.BaseRequest
	
	// Data migration task ID
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

func (r *StopMigrateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopMigrateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopMigrateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopMigrateJobResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type StopMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *StopMigrateJobResponseParams `json:"Response"`
}

func (r *StopMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopMigrateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubsErr struct {
	// Error message.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitnil" name:"Message"`
}

type SubscribeInfo struct {
	// Data subscription instance ID
	SubscribeId *string `json:"SubscribeId,omitnil" name:"SubscribeId"`

	// Data subscription instance name
	SubscribeName *string `json:"SubscribeName,omitnil" name:"SubscribeName"`

	// ID of channel bound to data subscription instance
	ChannelId *string `json:"ChannelId,omitnil" name:"ChannelId"`

	// Name of product bound to data subscription instance
	Product *string `json:"Product,omitnil" name:"Product"`

	// ID of database instance bound to data subscription instance
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Status of database instance bound to data subscription instance
	InstanceStatus *string `json:"InstanceStatus,omitnil" name:"InstanceStatus"`

	// Data subscription instance configuration status. Valid values: unconfigure, configuring, configured
	SubsStatus *string `json:"SubsStatus,omitnil" name:"SubsStatus"`

	// Last modified time
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// Isolation time
	IsolateTime *string `json:"IsolateTime,omitnil" name:"IsolateTime"`

	// Expiration time
	ExpireTime *string `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// Deactivation time
	OfflineTime *string `json:"OfflineTime,omitnil" name:"OfflineTime"`

	// Last modified consumption starting time point. If it has never been modified, this field is 0
	ConsumeStartTime *string `json:"ConsumeStartTime,omitnil" name:"ConsumeStartTime"`

	// Data subscription instance region
	Region *string `json:"Region,omitnil" name:"Region"`

	// Billing mode. 1: pay-as-you-go
	PayType *int64 `json:"PayType,omitnil" name:"PayType"`

	// Data subscription instance VIP
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// Data subscription instance Vport
	Vport *int64 `json:"Vport,omitnil" name:"Vport"`

	// Unique ID of the VPC where the data subscription instance VIP resides
	UniqVpcId *string `json:"UniqVpcId,omitnil" name:"UniqVpcId"`

	// Unique ID of the subnet where the data subscription instance VIP resides
	UniqSubnetId *string `json:"UniqSubnetId,omitnil" name:"UniqSubnetId"`

	// Data subscription instance status. Valid values: creating, normal, isolating, isolated, offlining, offline
	Status *string `json:"Status,omitnil" name:"Status"`

	// Timestamp of the last message confirmed by the SDK. If the SDK keeps consuming, this field can also be used as the current consumption time point of the SDK
	SdkConsumedTime *string `json:"SdkConsumedTime,omitnil" name:"SdkConsumedTime"`

	// Tag
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Tags []*TagItem `json:"Tags,omitnil" name:"Tags"`

	// Whether auto-renewal is enabled. 0: do not enable; 1: enable
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// Subscription instance edition. ·`txdts`: legacy data subscription; `kafka`: data subscription in Kafka edition
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SubscribeVersion *string `json:"SubscribeVersion,omitnil" name:"SubscribeVersion"`
}

type SubscribeObject struct {
	// Data subscription object type. 0: database, 1: database table
	// Note: this field may return null, indicating that no valid values can be obtained.
	ObjectsType *int64 `json:"ObjectsType,omitnil" name:"ObjectsType"`

	// Name of subscribed database
	// Note: this field may return null, indicating that no valid values can be obtained.
	DatabaseName *string `json:"DatabaseName,omitnil" name:"DatabaseName"`

	// Array of table names in subscribed database
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableNames []*string `json:"TableNames,omitnil" name:"TableNames"`
}

type SubscribeRegionConf struct {
	// Region name, such as Guangzhou
	// Note: this field may return null, indicating that no valid values can be obtained.
	RegionName *string `json:"RegionName,omitnil" name:"RegionName"`

	// Region ID, such as ap-guangzhou
	// Note: this field may return null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil" name:"Region"`

	// Region name, such as South China
	// Note: this field may return null, indicating that no valid values can be obtained.
	Area *string `json:"Area,omitnil" name:"Area"`

	// Whether it is the default region. 0: no, 1: yes
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsDefaultRegion *int64 `json:"IsDefaultRegion,omitnil" name:"IsDefaultRegion"`

	// Purchasable status of current region. 1: normal, 2: beta test, 3: not purchasable
	// Note: this field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitnil" name:"Status"`
}

type TagFilter struct {
	// Tag key value
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// Tag value
	TagValue []*string `json:"TagValue,omitnil" name:"TagValue"`
}

type TagItem struct {
	// Tag key value
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// Tag value
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TagValue *string `json:"TagValue,omitnil" name:"TagValue"`
}