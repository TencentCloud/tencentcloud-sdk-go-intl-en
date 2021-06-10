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

package v20180319

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type DeregisterMigrationTaskRequest struct {
	*tchttp.BaseRequest

	// Task ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DeregisterMigrationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeregisterMigrationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeregisterMigrationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeregisterMigrationTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeregisterMigrationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeregisterMigrationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMigrationTaskRequest struct {
	*tchttp.BaseRequest

	// Task ID, such as msp-jitoh33n
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeMigrationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMigrationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMigrationTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Migration details list
		TaskStatus []*TaskStatus `json:"TaskStatus,omitempty" name:"TaskStatus"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMigrationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DstInfo struct {

	// Migration destination region
	Region *string `json:"Region,omitempty" name:"Region"`

	// 
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Migration destination port
	Port *string `json:"Port,omitempty" name:"Port"`

	// Migration destination instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type ListMigrationProjectRequest struct {
	*tchttp.BaseRequest

	// The initial number of records, default value: 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// The number of records returned, default value: 500
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListMigrationProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListMigrationProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListMigrationProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListMigrationProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Project list
		Projects []*Project `json:"Projects,omitempty" name:"Projects"`

		// Total number of projects
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListMigrationProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListMigrationProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMigrationTaskRequest struct {
	*tchttp.BaseRequest

	// The initial number of records, default value: 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of records, default value: 10
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Project ID, the default value is empty.
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *ListMigrationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListMigrationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListMigrationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListMigrationTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of records
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Migration task list
		Tasks []*Task `json:"Tasks,omitempty" name:"Tasks"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListMigrationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListMigrationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMigrationTaskBelongToProjectRequest struct {
	*tchttp.BaseRequest

	// Task ID, such as msp-jitoh33n
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Project ID, such as 10005
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *ModifyMigrationTaskBelongToProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMigrationTaskBelongToProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMigrationTaskBelongToProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMigrationTaskBelongToProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMigrationTaskBelongToProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMigrationTaskBelongToProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMigrationTaskStatusRequest struct {
	*tchttp.BaseRequest

	// Task status, valid values include `unstart` (migration has not started), `migrating` (migration in progress), `finish` (migration completed) or `fail` (migration failed).
	Status *string `json:"Status,omitempty" name:"Status"`

	// Task ID, such as msp-jitoh33n
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *ModifyMigrationTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMigrationTaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMigrationTaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMigrationTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMigrationTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMigrationTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Project struct {

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Project name
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
}

type RegisterMigrationTaskRequest struct {
	*tchttp.BaseRequest

	// Task type, valid values include `database` (database migration), `file` (file migration) or `host` (host migration).
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// Task name
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// Service supplier name
	ServiceSupplier *string `json:"ServiceSupplier,omitempty" name:"ServiceSupplier"`

	// Migration task creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Migration task update time
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Migration type, for example `mysql:mysql` in database migration means migration from mysql to mysql and `oss:cos` in file migration means migration from Alibaba Cloud OSS to Tencent COS.
	MigrateClass *string `json:"MigrateClass,omitempty" name:"MigrateClass"`

	// Migration task source information
	SrcInfo *SrcInfo `json:"SrcInfo,omitempty" name:"SrcInfo"`

	// Migration task destination information
	DstInfo *DstInfo `json:"DstInfo,omitempty" name:"DstInfo"`

	// Source instance access type. Valid values for database migration include `extranet` (public network), `cvm` (CVM-created instance), `dcg` (Direct Connect-enabled instance), `vpncloud` (Tencent Cloud VPN-enabled instance), `vpnselfbuild` (self-built VPN-enabled instance), `cdb` (TencentDB instance)
	SrcAccessType *string `json:"SrcAccessType,omitempty" name:"SrcAccessType"`

	// Database type of the source instance. Valid values for database migration: `mysql`, `redis`, `percona`, `mongodb`, `postgresql`, `sqlserver`, `mariadb`
	SrcDatabaseType *string `json:"SrcDatabaseType,omitempty" name:"SrcDatabaseType"`

	// Target instance access type. Valid values for database migration include `extranet` (public network), `cvm` (CVM-created instance), `dcg` (Direct Connect-enabled instance), `vpncloud` (Tencent Cloud VPN-enabled instance), `vpnselfbuild` (self-built VPN-enabled instance), `cdb` (TencentDB instance)
	DstAccessType *string `json:"DstAccessType,omitempty" name:"DstAccessType"`

	// Database type of the target instance. Valid values for database migration: `mysql`, `redis`, `percona`, `mongodb`, `postgresql`, `sqlserver`, `mariadb`
	DstDatabaseType *string `json:"DstDatabaseType,omitempty" name:"DstDatabaseType"`
}

func (r *RegisterMigrationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterMigrationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskType")
	delete(f, "TaskName")
	delete(f, "ServiceSupplier")
	delete(f, "CreateTime")
	delete(f, "UpdateTime")
	delete(f, "MigrateClass")
	delete(f, "SrcInfo")
	delete(f, "DstInfo")
	delete(f, "SrcAccessType")
	delete(f, "SrcDatabaseType")
	delete(f, "DstAccessType")
	delete(f, "DstDatabaseType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RegisterMigrationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RegisterMigrationTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task ID
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RegisterMigrationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterMigrationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SrcInfo struct {

	// Migration source region
	Region *string `json:"Region,omitempty" name:"Region"`

	// 
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Migration source port
	Port *string `json:"Port,omitempty" name:"Port"`

	// Migration source instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type Task struct {

	// Task ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Task name
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// Migration type
	MigrationType *string `json:"MigrationType,omitempty" name:"MigrationType"`

	// Migration status
	Status *string `json:"Status,omitempty" name:"Status"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Project name
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// Migration source information
	SrcInfo *SrcInfo `json:"SrcInfo,omitempty" name:"SrcInfo"`

	// Migration time information
	MigrationTimeLine *TimeObj `json:"MigrationTimeLine,omitempty" name:"MigrationTimeLine"`

	// Status update time
	Updated *string `json:"Updated,omitempty" name:"Updated"`

	// Migration destination information
	DstInfo *DstInfo `json:"DstInfo,omitempty" name:"DstInfo"`
}

type TaskStatus struct {

	// Migration status
	Status *string `json:"Status,omitempty" name:"Status"`

	// Migration progress
	Progress *string `json:"Progress,omitempty" name:"Progress"`

	// Migration date
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type TimeObj struct {

	// The creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}
