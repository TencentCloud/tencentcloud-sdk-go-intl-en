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

package v20211228

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type AttachCBSSpec struct {
	// Node disk type, such as CLOUD_SSD"\"CLOUD_PREMIUM
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// Disk capacity, in GB
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// Total number of disks
	DiskCount *int64 `json:"DiskCount,omitnil,omitempty" name:"DiskCount"`

	// Description
	DiskDesc *string `json:"DiskDesc,omitnil,omitempty" name:"DiskDesc"`
}

type ChargeProperties struct {
	// Billing type: PREPAID for prepayment, and POSTPAID_BY_HOUR for postpayment.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// Whether to automatically renew. 1 means automatic renewal is enabled.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// Billing duration
	// Note: This field may return null, indicating that no valid values can be obtained.
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// Billing time unit, and "m" means month, etc.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`
}

type ClusterConfigsInfoFromEMR struct {
	// Configuration file's name
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// Related attribute information corresponding to the configuration files
	FileConf *string `json:"FileConf,omitnil,omitempty" name:"FileConf"`

	// Other attribute information corresponding to the configuration files
	KeyConf *string `json:"KeyConf,omitnil,omitempty" name:"KeyConf"`

	// Contents of the configuration files, base64 encoded
	OriParam *string `json:"OriParam,omitnil,omitempty" name:"OriParam"`

	// This is used to indicate whether the current configuration file has been modified without a restart, and reminds the user that a restart is needed.
	NeedRestart *int64 `json:"NeedRestart,omitnil,omitempty" name:"NeedRestart"`

	// Configuration file path
	// Note: This field may return null, indicating that no valid values can be obtained.
	FilePath *string `json:"FilePath,omitnil,omitempty" name:"FilePath"`

	// kv value of a configuration file
	// Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: FileKeyValues is deprecated.
	FileKeyValues *string `json:"FileKeyValues,omitnil,omitempty" name:"FileKeyValues"`

	// kv value of a configuration file
	// Note: This field may return null, indicating that no valid values can be obtained.
	FileKeyValuesNew []*ConfigKeyValue `json:"FileKeyValuesNew,omitnil,omitempty" name:"FileKeyValuesNew"`
}

type ConfigKeyValue struct {
	// key
	// Note: This field may return null, indicating that no valid values can be obtained.
	KeyName *string `json:"KeyName,omitnil,omitempty" name:"KeyName"`

	// Value
	// Note: This field may return null, indicating that no valid values can be obtained.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// Notes
	// Note: This field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 1 indicates read-only, 2 indicates editable but undeletable, and 3 indicates deletable.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Display *int64 `json:"Display,omitnil,omitempty" name:"Display"`

	// 0 means not supported, and 1 means hot update is supported.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SupportHotUpdate *int64 `json:"SupportHotUpdate,omitnil,omitempty" name:"SupportHotUpdate"`
}

// Predefined struct for user
type CreateInstanceNewRequestParams struct {
	// Availability zone
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// FE specifications
	FeSpec *CreateInstanceSpec `json:"FeSpec,omitnil,omitempty" name:"FeSpec"`

	// BE specifications.
	BeSpec *CreateInstanceSpec `json:"BeSpec,omitnil,omitempty" name:"BeSpec"`

	// Whether it is highly available.
	HaFlag *bool `json:"HaFlag,omitnil,omitempty" name:"HaFlag"`

	// User VPCID
	UserVPCId *string `json:"UserVPCId,omitnil,omitempty" name:"UserVPCId"`

	// User subnet ID
	UserSubnetId *string `json:"UserSubnetId,omitnil,omitempty" name:"UserSubnetId"`

	// Product version number
	ProductVersion *string `json:"ProductVersion,omitnil,omitempty" name:"ProductVersion"`

	// Payment type
	ChargeProperties *ChargeProperties `json:"ChargeProperties,omitnil,omitempty" name:"ChargeProperties"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Database password
	DorisUserPwd *string `json:"DorisUserPwd,omitnil,omitempty" name:"DorisUserPwd"`

	// Tag list
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// High availability type:
	// 0 indicates non-high availability (only one FE, FeSpec.CreateInstanceSpec.Count=1),
	// 1 indicates read high availability (at least 3 FEs must be deployed, FeSpec.CreateInstanceSpec.Count>=3, and it must be an odd number),
	// 2 indicates read and write high availability (at least 5 FEs must be deployed, FeSpec.CreateInstanceSpec.Count>=5, and it must be an odd number).
	HaType *int64 `json:"HaType,omitnil,omitempty" name:"HaType"`

	// Whether the table name is case sensitive, 0 refers to sensitive, 1 refers to insensitive, compared in lowercase; 2 refers to insensitive, and the table name is changed to lowercase for storage.
	CaseSensitive *int64 `json:"CaseSensitive,omitnil,omitempty" name:"CaseSensitive"`

	// Whether to enable multi-availability zone.
	EnableMultiZones *bool `json:"EnableMultiZones,omitnil,omitempty" name:"EnableMultiZones"`

	// After the Multi-AZ is enabled, all user's Availability Zones and Subnets information are shown.
	UserMultiZoneInfos *NetworkInfo `json:"UserMultiZoneInfos,omitnil,omitempty" name:"UserMultiZoneInfos"`
}

type CreateInstanceNewRequest struct {
	*tchttp.BaseRequest
	
	// Availability zone
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// FE specifications
	FeSpec *CreateInstanceSpec `json:"FeSpec,omitnil,omitempty" name:"FeSpec"`

	// BE specifications.
	BeSpec *CreateInstanceSpec `json:"BeSpec,omitnil,omitempty" name:"BeSpec"`

	// Whether it is highly available.
	HaFlag *bool `json:"HaFlag,omitnil,omitempty" name:"HaFlag"`

	// User VPCID
	UserVPCId *string `json:"UserVPCId,omitnil,omitempty" name:"UserVPCId"`

	// User subnet ID
	UserSubnetId *string `json:"UserSubnetId,omitnil,omitempty" name:"UserSubnetId"`

	// Product version number
	ProductVersion *string `json:"ProductVersion,omitnil,omitempty" name:"ProductVersion"`

	// Payment type
	ChargeProperties *ChargeProperties `json:"ChargeProperties,omitnil,omitempty" name:"ChargeProperties"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Database password
	DorisUserPwd *string `json:"DorisUserPwd,omitnil,omitempty" name:"DorisUserPwd"`

	// Tag list
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// High availability type:
	// 0 indicates non-high availability (only one FE, FeSpec.CreateInstanceSpec.Count=1),
	// 1 indicates read high availability (at least 3 FEs must be deployed, FeSpec.CreateInstanceSpec.Count>=3, and it must be an odd number),
	// 2 indicates read and write high availability (at least 5 FEs must be deployed, FeSpec.CreateInstanceSpec.Count>=5, and it must be an odd number).
	HaType *int64 `json:"HaType,omitnil,omitempty" name:"HaType"`

	// Whether the table name is case sensitive, 0 refers to sensitive, 1 refers to insensitive, compared in lowercase; 2 refers to insensitive, and the table name is changed to lowercase for storage.
	CaseSensitive *int64 `json:"CaseSensitive,omitnil,omitempty" name:"CaseSensitive"`

	// Whether to enable multi-availability zone.
	EnableMultiZones *bool `json:"EnableMultiZones,omitnil,omitempty" name:"EnableMultiZones"`

	// After the Multi-AZ is enabled, all user's Availability Zones and Subnets information are shown.
	UserMultiZoneInfos *NetworkInfo `json:"UserMultiZoneInfos,omitnil,omitempty" name:"UserMultiZoneInfos"`
}

func (r *CreateInstanceNewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceNewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "FeSpec")
	delete(f, "BeSpec")
	delete(f, "HaFlag")
	delete(f, "UserVPCId")
	delete(f, "UserSubnetId")
	delete(f, "ProductVersion")
	delete(f, "ChargeProperties")
	delete(f, "InstanceName")
	delete(f, "DorisUserPwd")
	delete(f, "Tags")
	delete(f, "HaType")
	delete(f, "CaseSensitive")
	delete(f, "EnableMultiZones")
	delete(f, "UserMultiZoneInfos")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstanceNewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceNewResponseParams struct {
	// Process ID
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Error message
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateInstanceNewResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstanceNewResponseParams `json:"Response"`
}

func (r *CreateInstanceNewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceSpec struct {
	// Specification name
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// Quantities
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// Cloud disk size
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`
}

type DataBaseAuditRecord struct {
	// Query user
	// Note: This field may return null, indicating that no valid values can be obtained.
	OsUser *string `json:"OsUser,omitnil,omitempty" name:"OsUser"`

	// Query ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	InitialQueryId *string `json:"InitialQueryId,omitnil,omitempty" name:"InitialQueryId"`

	// SQL statement
	// Note: This field may return null, indicating that no valid values can be obtained.
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// Start time
	// Note: This field may return null, indicating that no valid values can be obtained.
	QueryStartTime *string `json:"QueryStartTime,omitnil,omitempty" name:"QueryStartTime"`

	// Execution duration
	// Note: This field may return null, indicating that no valid values can be obtained.
	DurationMs *int64 `json:"DurationMs,omitnil,omitempty" name:"DurationMs"`

	// The number of read rows
	// Note: This field may return null, indicating that no valid values can be obtained.
	ReadRows *int64 `json:"ReadRows,omitnil,omitempty" name:"ReadRows"`

	// Total number of read bytes
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResultRows *int64 `json:"ResultRows,omitnil,omitempty" name:"ResultRows"`

	// Result bytes
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResultBytes *uint64 `json:"ResultBytes,omitnil,omitempty" name:"ResultBytes"`

	// Memory
	// Note: This field may return null, indicating that no valid values can be obtained.
	MemoryUsage *int64 `json:"MemoryUsage,omitnil,omitempty" name:"MemoryUsage"`

	// Initial query IP
	// Note: This field may return null, indicating that no valid values can be obtained.
	InitialAddress *string `json:"InitialAddress,omitnil,omitempty" name:"InitialAddress"`

	// Database
	// Note: This field may return null, indicating that no valid values can be obtained.
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// SQL type
	// Note: This field may return null, indicating that no valid values can be obtained.
	SqlType *string `json:"SqlType,omitnil,omitempty" name:"SqlType"`

	// Catalog name
	// Note: This field may return null, indicating that no valid values can be obtained.
	Catalog *string `json:"Catalog,omitnil,omitempty" name:"Catalog"`
}

// Predefined struct for user
type DescribeClusterConfigsRequestParams struct {
	// Cluster instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 0 indicates public cloud query, and 1 Qinge query. Qinge query shows all that needs to be displayed.
	ConfigType *int64 `json:"ConfigType,omitnil,omitempty" name:"ConfigType"`

	// Search for files with fuzzy keywords
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 0 indicates cluster dimension and 1 node dimension
	ClusterConfigType *int64 `json:"ClusterConfigType,omitnil,omitempty" name:"ClusterConfigType"`

	// eth0's IP address
	IPAddress *string `json:"IPAddress,omitnil,omitempty" name:"IPAddress"`
}

type DescribeClusterConfigsRequest struct {
	*tchttp.BaseRequest
	
	// Cluster instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 0 indicates public cloud query, and 1 Qinge query. Qinge query shows all that needs to be displayed.
	ConfigType *int64 `json:"ConfigType,omitnil,omitempty" name:"ConfigType"`

	// Search for files with fuzzy keywords
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 0 indicates cluster dimension and 1 node dimension
	ClusterConfigType *int64 `json:"ClusterConfigType,omitnil,omitempty" name:"ClusterConfigType"`

	// eth0's IP address
	IPAddress *string `json:"IPAddress,omitnil,omitempty" name:"IPAddress"`
}

func (r *DescribeClusterConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ConfigType")
	delete(f, "FileName")
	delete(f, "ClusterConfigType")
	delete(f, "IPAddress")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterConfigsResponseParams struct {
	// Return information about the instance's configuration file.
	ClusterConfList []*ClusterConfigsInfoFromEMR `json:"ClusterConfList,omitnil,omitempty" name:"ClusterConfList"`

	// Return the current kernel version. If it does not exist, a null character string is returned.
	BuildVersion *string `json:"BuildVersion,omitnil,omitempty" name:"BuildVersion"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterConfigsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterConfigsResponseParams `json:"Response"`
}

func (r *DescribeClusterConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseAuditDownloadRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Paging
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Paging
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// Sort parameters
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// User
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// Database
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// SQL type
	SqlType *string `json:"SqlType,omitnil,omitempty" name:"SqlType"`

	// SQL statement
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// Users (multiple selections)
	Users []*string `json:"Users,omitnil,omitempty" name:"Users"`

	// Databases (multiple selections)
	DbNames []*string `json:"DbNames,omitnil,omitempty" name:"DbNames"`

	// SQL types (multiple selections)
	SqlTypes []*string `json:"SqlTypes,omitnil,omitempty" name:"SqlTypes"`

	// Catalog names (multiple selections)
	Catalogs []*string `json:"Catalogs,omitnil,omitempty" name:"Catalogs"`
}

type DescribeDatabaseAuditDownloadRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Paging
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Paging
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// Sort parameters
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// User
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// Database
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// SQL type
	SqlType *string `json:"SqlType,omitnil,omitempty" name:"SqlType"`

	// SQL statement
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// Users (multiple selections)
	Users []*string `json:"Users,omitnil,omitempty" name:"Users"`

	// Databases (multiple selections)
	DbNames []*string `json:"DbNames,omitnil,omitempty" name:"DbNames"`

	// SQL types (multiple selections)
	SqlTypes []*string `json:"SqlTypes,omitnil,omitempty" name:"SqlTypes"`

	// Catalog names (multiple selections)
	Catalogs []*string `json:"Catalogs,omitnil,omitempty" name:"Catalogs"`
}

func (r *DescribeDatabaseAuditDownloadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseAuditDownloadRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageSize")
	delete(f, "PageNum")
	delete(f, "OrderType")
	delete(f, "User")
	delete(f, "DbName")
	delete(f, "SqlType")
	delete(f, "Sql")
	delete(f, "Users")
	delete(f, "DbNames")
	delete(f, "SqlTypes")
	delete(f, "Catalogs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatabaseAuditDownloadRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseAuditDownloadResponseParams struct {
	// The cos address of the log
	CosUrl *string `json:"CosUrl,omitnil,omitempty" name:"CosUrl"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDatabaseAuditDownloadResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatabaseAuditDownloadResponseParams `json:"Response"`
}

func (r *DescribeDatabaseAuditDownloadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseAuditDownloadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseAuditRecordsRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Paging
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Paging
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// Sort parameters
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// User
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// Database
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// SQL type
	SqlType *string `json:"SqlType,omitnil,omitempty" name:"SqlType"`

	// SQL statement
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// Users (multiple selections)
	Users []*string `json:"Users,omitnil,omitempty" name:"Users"`

	// Databases (multiple selections)
	DbNames []*string `json:"DbNames,omitnil,omitempty" name:"DbNames"`

	// SQL types (multiple selections)
	SqlTypes []*string `json:"SqlTypes,omitnil,omitempty" name:"SqlTypes"`

	// Catalog names (multiple selections)
	Catalogs []*string `json:"Catalogs,omitnil,omitempty" name:"Catalogs"`
}

type DescribeDatabaseAuditRecordsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Paging
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Paging
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// Sort parameters
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// User
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// Database
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// SQL type
	SqlType *string `json:"SqlType,omitnil,omitempty" name:"SqlType"`

	// SQL statement
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// Users (multiple selections)
	Users []*string `json:"Users,omitnil,omitempty" name:"Users"`

	// Databases (multiple selections)
	DbNames []*string `json:"DbNames,omitnil,omitempty" name:"DbNames"`

	// SQL types (multiple selections)
	SqlTypes []*string `json:"SqlTypes,omitnil,omitempty" name:"SqlTypes"`

	// Catalog names (multiple selections)
	Catalogs []*string `json:"Catalogs,omitnil,omitempty" name:"Catalogs"`
}

func (r *DescribeDatabaseAuditRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseAuditRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageSize")
	delete(f, "PageNum")
	delete(f, "OrderType")
	delete(f, "User")
	delete(f, "DbName")
	delete(f, "SqlType")
	delete(f, "Sql")
	delete(f, "Users")
	delete(f, "DbNames")
	delete(f, "SqlTypes")
	delete(f, "Catalogs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatabaseAuditRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseAuditRecordsResponseParams struct {
	// Total
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Record list
	SlowQueryRecords *DataBaseAuditRecord `json:"SlowQueryRecords,omitnil,omitempty" name:"SlowQueryRecords"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDatabaseAuditRecordsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatabaseAuditRecordsResponseParams `json:"Response"`
}

func (r *DescribeDatabaseAuditRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseAuditRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceNodesInfoRequestParams struct {
	// Cluster ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
}

type DescribeInstanceNodesInfoRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
}

func (r *DescribeInstanceNodesInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceNodesInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceNodesInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceNodesInfoResponseParams struct {
	// Be node
	// Note: This field may return null, indicating that no valid values can be obtained.
	BeNodes []*string `json:"BeNodes,omitnil,omitempty" name:"BeNodes"`

	// Fe node
	// Note: This field may return null, indicating that no valid values can be obtained.
	FeNodes []*string `json:"FeNodes,omitnil,omitempty" name:"FeNodes"`

	// Fe master node
	FeMaster *string `json:"FeMaster,omitnil,omitempty" name:"FeMaster"`

	// Be node information
	// Note: This field may return null, indicating that no valid values can be obtained.
	BeNodeInfos []*NodeInfo `json:"BeNodeInfos,omitnil,omitempty" name:"BeNodeInfos"`

	// Fe node information
	// Note: This field may return null, indicating that no valid values can be obtained.
	FeNodeInfos []*NodeInfo `json:"FeNodeInfos,omitnil,omitempty" name:"FeNodeInfos"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceNodesInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceNodesInfoResponseParams `json:"Response"`
}

func (r *DescribeInstanceNodesInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceNodesInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceNodesRequestParams struct {
	// Cluster instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Cluster role type, defaulted as "data node".
	NodeRole *string `json:"NodeRole,omitnil,omitempty" name:"NodeRole"`

	// Pagination parameters. The first page is 0, and the second page is 10.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Pagination parameters. The pagination step length is 10 by default.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Display policy, and all items are displayed when All is selected.
	DisplayPolicy *string `json:"DisplayPolicy,omitnil,omitempty" name:"DisplayPolicy"`
}

type DescribeInstanceNodesRequest struct {
	*tchttp.BaseRequest
	
	// Cluster instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Cluster role type, defaulted as "data node".
	NodeRole *string `json:"NodeRole,omitnil,omitempty" name:"NodeRole"`

	// Pagination parameters. The first page is 0, and the second page is 10.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Pagination parameters. The pagination step length is 10 by default.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Display policy, and all items are displayed when All is selected.
	DisplayPolicy *string `json:"DisplayPolicy,omitnil,omitempty" name:"DisplayPolicy"`
}

func (r *DescribeInstanceNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "NodeRole")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "DisplayPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceNodesResponseParams struct {
	// Total number
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Total number of instance nodes
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceNodesList []*InstanceNode `json:"InstanceNodesList,omitnil,omitempty" name:"InstanceNodesList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceNodesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceNodesResponseParams `json:"Response"`
}

func (r *DescribeInstanceNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceRequestParams struct {
	// Cluster instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Cluster instance ID
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
	// Instance description information
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
	// Cluster instance name
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceStateRequest struct {
	*tchttp.BaseRequest
	
	// Cluster instance name
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
	// Cluster status. Example: Serving
	InstanceState *string `json:"InstanceState,omitnil,omitempty" name:"InstanceState"`

	// Creation time of cluster operation
	// Note: This field may return null, indicating that no valid values can be obtained.
	FlowCreateTime *string `json:"FlowCreateTime,omitnil,omitempty" name:"FlowCreateTime"`

	// Cluster operation name
	// Note: This field may return null, indicating that no valid values can be obtained.
	FlowName *string `json:"FlowName,omitnil,omitempty" name:"FlowName"`

	// Cluster operation progress
	// Note: This field may return null, indicating that no valid values can be obtained.
	FlowProgress *float64 `json:"FlowProgress,omitnil,omitempty" name:"FlowProgress"`

	// Cluster status description. Example: running
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceStateDesc *string `json:"InstanceStateDesc,omitnil,omitempty" name:"InstanceStateDesc"`

	// Cluster process error messages, such as "Creation failed, insufficient resources"
	// Note: This field may return null, indicating that no valid values can be obtained.
	FlowMsg *string `json:"FlowMsg,omitnil,omitempty" name:"FlowMsg"`

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
	// The name of the cluster ID for the search
	SearchInstanceId *string `json:"SearchInstanceId,omitnil,omitempty" name:"SearchInstanceId"`

	// The cluster name for the search
	SearchInstanceName *string `json:"SearchInstanceName,omitnil,omitempty" name:"SearchInstanceName"`

	// Pagination parameters. The first page is 0, and the second page is 10.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Pagination parameters. The pagination step length is 10 by default.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Search tag list
	SearchTags []*SearchTags `json:"SearchTags,omitnil,omitempty" name:"SearchTags"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// The name of the cluster ID for the search
	SearchInstanceId *string `json:"SearchInstanceId,omitnil,omitempty" name:"SearchInstanceId"`

	// The cluster name for the search
	SearchInstanceName *string `json:"SearchInstanceName,omitnil,omitempty" name:"SearchInstanceName"`

	// Pagination parameters. The first page is 0, and the second page is 10.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Pagination parameters. The pagination step length is 10 by default.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Search tag list
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
	// Total Number of Instances
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Quantities of instances array
	InstancesList []*InstanceInfo `json:"InstancesList,omitnil,omitempty" name:"InstancesList"`

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
type DescribeSlowQueryRecordsDownloadRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Slow log time
	QueryDurationMs *int64 `json:"QueryDurationMs,omitnil,omitempty" name:"QueryDurationMs"`

	// Start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Sort parameters
	DurationMs *string `json:"DurationMs,omitnil,omitempty" name:"DurationMs"`

	// Query SQL
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// Sort parameters
	ReadRows *string `json:"ReadRows,omitnil,omitempty" name:"ReadRows"`

	// Sort parameters
	ResultBytes *string `json:"ResultBytes,omitnil,omitempty" name:"ResultBytes"`

	// Sort parameters
	MemoryUsage *string `json:"MemoryUsage,omitnil,omitempty" name:"MemoryUsage"`

	// IsQuery condition
	IsQuery *int64 `json:"IsQuery,omitnil,omitempty" name:"IsQuery"`

	// Database name
	DbName []*string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// catalog name
	CatalogName []*string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`
}

type DescribeSlowQueryRecordsDownloadRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Slow log time
	QueryDurationMs *int64 `json:"QueryDurationMs,omitnil,omitempty" name:"QueryDurationMs"`

	// Start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Sort parameters
	DurationMs *string `json:"DurationMs,omitnil,omitempty" name:"DurationMs"`

	// Query SQL
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// Sort parameters
	ReadRows *string `json:"ReadRows,omitnil,omitempty" name:"ReadRows"`

	// Sort parameters
	ResultBytes *string `json:"ResultBytes,omitnil,omitempty" name:"ResultBytes"`

	// Sort parameters
	MemoryUsage *string `json:"MemoryUsage,omitnil,omitempty" name:"MemoryUsage"`

	// IsQuery condition
	IsQuery *int64 `json:"IsQuery,omitnil,omitempty" name:"IsQuery"`

	// Database name
	DbName []*string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// catalog name
	CatalogName []*string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`
}

func (r *DescribeSlowQueryRecordsDownloadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowQueryRecordsDownloadRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "QueryDurationMs")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "DurationMs")
	delete(f, "Sql")
	delete(f, "ReadRows")
	delete(f, "ResultBytes")
	delete(f, "MemoryUsage")
	delete(f, "IsQuery")
	delete(f, "DbName")
	delete(f, "CatalogName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowQueryRecordsDownloadRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowQueryRecordsDownloadResponseParams struct {
	// cos address
	CosUrl *string `json:"CosUrl,omitnil,omitempty" name:"CosUrl"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSlowQueryRecordsDownloadResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowQueryRecordsDownloadResponseParams `json:"Response"`
}

func (r *DescribeSlowQueryRecordsDownloadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowQueryRecordsDownloadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowQueryRecordsRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Slow log time
	QueryDurationMs *int64 `json:"QueryDurationMs,omitnil,omitempty" name:"QueryDurationMs"`

	// Start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Paging
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Paging
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// Sort parameters
	DurationMs *string `json:"DurationMs,omitnil,omitempty" name:"DurationMs"`

	// Database name
	DbName []*string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Whether it is a query. 0 indicates no, and 1 indicates yes.
	IsQuery *int64 `json:"IsQuery,omitnil,omitempty" name:"IsQuery"`

	// catalog name
	CatalogName []*string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`

	// SQL name
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// ReadRows sort field
	ReadRows *string `json:"ReadRows,omitnil,omitempty" name:"ReadRows"`

	// ResultBytes sort field
	ResultBytes *string `json:"ResultBytes,omitnil,omitempty" name:"ResultBytes"`

	// MemoryUsage sort field
	MemoryUsage *string `json:"MemoryUsage,omitnil,omitempty" name:"MemoryUsage"`
}

type DescribeSlowQueryRecordsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Slow log time
	QueryDurationMs *int64 `json:"QueryDurationMs,omitnil,omitempty" name:"QueryDurationMs"`

	// Start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Paging
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Paging
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// Sort parameters
	DurationMs *string `json:"DurationMs,omitnil,omitempty" name:"DurationMs"`

	// Database name
	DbName []*string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Whether it is a query. 0 indicates no, and 1 indicates yes.
	IsQuery *int64 `json:"IsQuery,omitnil,omitempty" name:"IsQuery"`

	// catalog name
	CatalogName []*string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`

	// SQL name
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// ReadRows sort field
	ReadRows *string `json:"ReadRows,omitnil,omitempty" name:"ReadRows"`

	// ResultBytes sort field
	ResultBytes *string `json:"ResultBytes,omitnil,omitempty" name:"ResultBytes"`

	// MemoryUsage sort field
	MemoryUsage *string `json:"MemoryUsage,omitnil,omitempty" name:"MemoryUsage"`
}

func (r *DescribeSlowQueryRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowQueryRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "QueryDurationMs")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageSize")
	delete(f, "PageNum")
	delete(f, "DurationMs")
	delete(f, "DbName")
	delete(f, "IsQuery")
	delete(f, "CatalogName")
	delete(f, "Sql")
	delete(f, "ReadRows")
	delete(f, "ResultBytes")
	delete(f, "MemoryUsage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowQueryRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowQueryRecordsResponseParams struct {
	// Total
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Record list
	SlowQueryRecords []*SlowQueryRecord `json:"SlowQueryRecords,omitnil,omitempty" name:"SlowQueryRecords"`

	// All database names
	// Note: This field may return null, indicating that no valid values can be obtained.
	DBNameList []*string `json:"DBNameList,omitnil,omitempty" name:"DBNameList"`

	// All catalog names
	// Note: This field may return null, indicating that no valid values can be obtained.
	CatalogNameList []*string `json:"CatalogNameList,omitnil,omitempty" name:"CatalogNameList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSlowQueryRecordsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowQueryRecordsResponseParams `json:"Response"`
}

func (r *DescribeSlowQueryRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowQueryRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyInstanceRequestParams struct {
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DestroyInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DestroyInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyInstanceResponseParams struct {
	// Process ID
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Error message
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DestroyInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DestroyInstanceResponseParams `json:"Response"`
}

func (r *DestroyInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceInfo struct {
	// Cluster instance ID, "cdw-xxxx" string type
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Cluster instance name
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Status,
	// Init is being created. Serving is running. 
	// Deleted indicates the cluster has been terminated. Deleting indicates the cluster is being terminated.
	// Modify indicates the cluster is being changed.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Version
	// Note: This field may return null, indicating that no valid values can be obtained.
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// Region, ap-guangzhou
	// Note: This field may return null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Availability zone, ap-guangzhou-3
	// Note: This field may return null, indicating that no valid values can be obtained.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// VPC name
	// Note: This field may return null, indicating that no valid values can be obtained.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet name
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Payment type: hour and prepay
	// Note: This field may return null, indicating that no valid values can be obtained.
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Creation time
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Expiration time
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// Data node description information
	// Note: This field may return null, indicating that no valid values can be obtained.
	MasterSummary *NodesSummary `json:"MasterSummary,omitnil,omitempty" name:"MasterSummary"`

	// Zookeeper node description information
	// Note: This field may return null, indicating that no valid values can be obtained.
	CoreSummary *NodesSummary `json:"CoreSummary,omitnil,omitempty" name:"CoreSummary"`

	// High availability, being true or false
	// Note: This field may return null, indicating that no valid values can be obtained.
	HA *string `json:"HA,omitnil,omitempty" name:"HA"`

	// High availability type:
	// 0: non-high availability
	// 1: read high availability
	// 2: read-write high availability
	// Note: This field may return null, indicating that no valid values can be obtained.
	HaType *int64 `json:"HaType,omitnil,omitempty" name:"HaType"`

	// Access address. Example: 10.0.0.1:9000
	// Note: This field may return null, indicating that no valid values can be obtained.
	AccessInfo *string `json:"AccessInfo,omitnil,omitempty" name:"AccessInfo"`

	// Record ID, in numerical type
	// Note: This field may return null, indicating that no valid values can be obtained.
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Region ID, indicating the region
	// Note: This field may return null, indicating that no valid values can be obtained.
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Note about availability zone, such as Guangzhou Zone 2
	// Note: This field may return null, indicating that no valid values can be obtained.
	ZoneDesc *string `json:"ZoneDesc,omitnil,omitempty" name:"ZoneDesc"`

	// Error process description information
	// Note: This field may return null, indicating that no valid values can be obtained.
	FlowMsg *string `json:"FlowMsg,omitnil,omitempty" name:"FlowMsg"`

	// Status description, such as "running"
	// Note: This field may return null, indicating that no valid values can be obtained.
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// Automatic renewal marker
	// Note: This field may return null, indicating that no valid values can be obtained.
	RenewFlag *bool `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// Tag list
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Monitoring Information
	// Note: This field may return null, indicating that no valid values can be obtained.
	Monitor *string `json:"Monitor,omitnil,omitempty" name:"Monitor"`

	// Whether to enable logs.
	// Note: This field may return null, indicating that no valid values can be obtained.
	HasClsTopic *bool `json:"HasClsTopic,omitnil,omitempty" name:"HasClsTopic"`

	// Log Topic ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClsTopicId *string `json:"ClsTopicId,omitnil,omitempty" name:"ClsTopicId"`

	// Logset ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClsLogSetId *string `json:"ClsLogSetId,omitnil,omitempty" name:"ClsLogSetId"`

	// Whether to support XML configuration management.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EnableXMLConfig *int64 `json:"EnableXMLConfig,omitnil,omitempty" name:"EnableXMLConfig"`

	// Region
	// Note: This field may return null, indicating that no valid values can be obtained.
	RegionDesc *string `json:"RegionDesc,omitnil,omitempty" name:"RegionDesc"`

	// Elastic network interface address
	// Note: This field may return null, indicating that no valid values can be obtained.
	Eip *string `json:"Eip,omitnil,omitempty" name:"Eip"`

	// Cold and hot stratification coefficient
	// Note: This field may return null, indicating that no valid values can be obtained.
	CosMoveFactor *int64 `json:"CosMoveFactor,omitnil,omitempty" name:"CosMoveFactor"`

	// external/local/yunti
	// Note: This field may return null, indicating that no valid values can be obtained.
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// COS bucket
	// Note: This field may return null, indicating that no valid values can be obtained.
	CosBucketName *string `json:"CosBucketName,omitnil,omitempty" name:"CosBucketName"`

	// cbs
	// Note: This field may return null, indicating that no valid values can be obtained.
	CanAttachCbs *bool `json:"CanAttachCbs,omitnil,omitempty" name:"CanAttachCbs"`

	// Minor versions
	// Note: This field may return null, indicating that no valid values can be obtained.
	BuildVersion *string `json:"BuildVersion,omitnil,omitempty" name:"BuildVersion"`

	// Component Information
	// Note: The return type here is map[string]struct, not the string type displayed. You can refer to "Sample Value" to parse the data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Components *string `json:"Components,omitnil,omitempty" name:"Components"`

	// Determine whether the audit log table has a catalog field.
	// Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: IfExistCatalog is deprecated.
	IfExistCatalog *int64 `json:"IfExistCatalog,omitnil,omitempty" name:"IfExistCatalog"`

	// Page features, used to block some page entrances on the front end.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Characteristic []*string `json:"Characteristic,omitnil,omitempty" name:"Characteristic"`

	// Timeout period, in seconds
	// Note: This field may return null, indicating that no valid values can be obtained.
	RestartTimeout *string `json:"RestartTimeout,omitnil,omitempty" name:"RestartTimeout"`

	// The timeout time for the graceful restart of the kernel. If it is -1, it means it is not set.
	// Note: This field may return null, indicating that no valid values can be obtained.
	GraceShutdownWaitSeconds *string `json:"GraceShutdownWaitSeconds,omitnil,omitempty" name:"GraceShutdownWaitSeconds"`

	// Whether the table name is case sensitive, 0 refers to sensitive, 1 refers to insensitive, compared in lowercase; 2 refers to insensitive, and the table name is changed to lowercase for storage.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CaseSensitive *int64 `json:"CaseSensitive,omitnil,omitempty" name:"CaseSensitive"`

	// Whether users can bind security groups.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsWhiteSGs *bool `json:"IsWhiteSGs,omitnil,omitempty" name:"IsWhiteSGs"`

	// Bound security group information
	// Note: This field may return null, indicating that no valid values can be obtained.
	BindSGs []*string `json:"BindSGs,omitnil,omitempty" name:"BindSGs"`

	// Whether it is a multi-AZ.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EnableMultiZones *bool `json:"EnableMultiZones,omitnil,omitempty" name:"EnableMultiZones"`

	// User availability zone and subnet information
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserNetworkInfos *string `json:"UserNetworkInfos,omitnil,omitempty" name:"UserNetworkInfos"`

	// Whether to enable hot and cold stratification. 0 refers to disabled, and 1 refers to enabled.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EnableCoolDown *int64 `json:"EnableCoolDown,omitnil,omitempty" name:"EnableCoolDown"`

	// COS buckets are used for hot and cold stratification
	// Note: This field may return null, indicating that no valid values can be obtained.
	CoolDownBucket *string `json:"CoolDownBucket,omitnil,omitempty" name:"CoolDownBucket"`
}

type InstanceNode struct {
	// IP address
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// Model, such as S1
	Spec *string `json:"Spec,omitnil,omitempty" name:"Spec"`

	// Number of CPU cores
	Core *int64 `json:"Core,omitnil,omitempty" name:"Core"`

	// Memory size
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Disk type
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// Disk size
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// The name of the clickhouse cluster to which it belongs.
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// Status
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// rip
	// Note: This field may return null, indicating that no valid values can be obtained.
	Rip *string `json:"Rip,omitnil,omitempty" name:"Rip"`

	// FE node role
	// Note: This field may return null, indicating that no valid values can be obtained.
	FeRole *string `json:"FeRole,omitnil,omitempty" name:"FeRole"`

	// UUID
	// Note: This field may return null, indicating that no valid values can be obtained.
	UUID *string `json:"UUID,omitnil,omitempty" name:"UUID"`
}

// Predefined struct for user
type ModifyInstanceRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Newly modified instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

type ModifyInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Newly modified instance name
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

type NetworkInfo struct {
	// Availability zone
	// Note: This field may return null, indicating that no valid values can be obtained.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Subnet ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// The number of available IP addresses in the current subnet
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubnetIpNum *int64 `json:"SubnetIpNum,omitnil,omitempty" name:"SubnetIpNum"`
}

type NodeInfo struct {
	// User IP
	// Note: This field may return null, indicating that no valid values can be obtained.
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// Node status
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Node role name
	// Note: This field may return null, indicating that no valid values can be obtained.
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// Component name
	// Note: This field may return null, indicating that no valid values can be obtained.
	ComponentName *string `json:"ComponentName,omitnil,omitempty" name:"ComponentName"`

	// Node role
	// Note: This field may return null, indicating that no valid values can be obtained.
	NodeRole *string `json:"NodeRole,omitnil,omitempty" name:"NodeRole"`

	// The time when the node was last restarted
	// Note: This field may return null, indicating that no valid values can be obtained.
	LastRestartTime *string `json:"LastRestartTime,omitnil,omitempty" name:"LastRestartTime"`

	// The availability zone where the node is located
	// Note: This field may return null, indicating that no valid values can be obtained.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

type NodesSummary struct {
	// Model, such as S1
	Spec *string `json:"Spec,omitnil,omitempty" name:"Spec"`

	// Number of nodes
	NodeSize *int64 `json:"NodeSize,omitnil,omitempty" name:"NodeSize"`

	// Number of CPU cores, in counts
	Core *int64 `json:"Core,omitnil,omitempty" name:"Core"`

	// Memory size, in GB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Disk size, in GB
	Disk *int64 `json:"Disk,omitnil,omitempty" name:"Disk"`

	// Disk type
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// Disk description
	DiskDesc *string `json:"DiskDesc,omitnil,omitempty" name:"DiskDesc"`

	// Information of mounted cloud disks
	// Note: This field may return null, indicating that no valid values can be obtained.
	AttachCBSSpec *AttachCBSSpec `json:"AttachCBSSpec,omitnil,omitempty" name:"AttachCBSSpec"`

	// Sub-product name
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubProductType *string `json:"SubProductType,omitnil,omitempty" name:"SubProductType"`

	// Specified cores
	// Note: This field may return null, indicating that no valid values can be obtained.
	SpecCore *int64 `json:"SpecCore,omitnil,omitempty" name:"SpecCore"`

	// Specified memory
	// Note: This field may return null, indicating that no valid values can be obtained.
	SpecMemory *int64 `json:"SpecMemory,omitnil,omitempty" name:"SpecMemory"`

	// Disk size
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskCount *int64 `json:"DiskCount,omitnil,omitempty" name:"DiskCount"`

	// Whether it is encrypted.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Encrypt *int64 `json:"Encrypt,omitnil,omitempty" name:"Encrypt"`

	// Maximum disk
	// Note: This field may return null, indicating that no valid values can be obtained.
	MaxDiskSize *int64 `json:"MaxDiskSize,omitnil,omitempty" name:"MaxDiskSize"`
}

// Predefined struct for user
type ResizeDiskRequestParams struct {
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Role (MATER/CORE), MASTER corresponds to FE, CORE corresponds to BE.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Cloud disk size
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`
}

type ResizeDiskRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Role (MATER/CORE), MASTER corresponds to FE, CORE corresponds to BE.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Cloud disk size
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`
}

func (r *ResizeDiskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResizeDiskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Type")
	delete(f, "DiskSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResizeDiskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResizeDiskResponseParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Process ID
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Error message
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResizeDiskResponse struct {
	*tchttp.BaseResponse
	Response *ResizeDiskResponseParams `json:"Response"`
}

func (r *ResizeDiskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResizeDiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartClusterForNodeRequestParams struct {
	// Cluster ID, such as cdwch-xxxx
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Configuration file's name
	ConfigName *string `json:"ConfigName,omitnil,omitempty" name:"ConfigName"`

	// Each batch of restarts
	BatchSize *int64 `json:"BatchSize,omitnil,omitempty" name:"BatchSize"`

	// Restart node
	NodeList []*string `json:"NodeList,omitnil,omitempty" name:"NodeList"`

	// False means non-rolling restart, and true means rolling restart.
	RollingRestart *bool `json:"RollingRestart,omitnil,omitempty" name:"RollingRestart"`
}

type RestartClusterForNodeRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID, such as cdwch-xxxx
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Configuration file's name
	ConfigName *string `json:"ConfigName,omitnil,omitempty" name:"ConfigName"`

	// Each batch of restarts
	BatchSize *int64 `json:"BatchSize,omitnil,omitempty" name:"BatchSize"`

	// Restart node
	NodeList []*string `json:"NodeList,omitnil,omitempty" name:"NodeList"`

	// False means non-rolling restart, and true means rolling restart.
	RollingRestart *bool `json:"RollingRestart,omitnil,omitempty" name:"RollingRestart"`
}

func (r *RestartClusterForNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartClusterForNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ConfigName")
	delete(f, "BatchSize")
	delete(f, "NodeList")
	delete(f, "RollingRestart")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartClusterForNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartClusterForNodeResponseParams struct {
	// Process related information
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Error message
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RestartClusterForNodeResponse struct {
	*tchttp.BaseResponse
	Response *RestartClusterForNodeResponseParams `json:"Response"`
}

func (r *RestartClusterForNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartClusterForNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleOutInstanceRequestParams struct {
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Role (MATER/CORE), MASTER corresponds to FE, CORE corresponds to BE.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Number of nodes
	NodeCount *uint64 `json:"NodeCount,omitnil,omitempty" name:"NodeCount"`

	// Cluster high availability type after scaled out: 0 indicates non-high availability, 1 indicates read high availability, and 2 indicates read-write high availability.
	HaType *int64 `json:"HaType,omitnil,omitempty" name:"HaType"`
}

type ScaleOutInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Role (MATER/CORE), MASTER corresponds to FE, CORE corresponds to BE.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Number of nodes
	NodeCount *uint64 `json:"NodeCount,omitnil,omitempty" name:"NodeCount"`

	// Cluster high availability type after scaled out: 0 indicates non-high availability, 1 indicates read high availability, and 2 indicates read-write high availability.
	HaType *int64 `json:"HaType,omitnil,omitempty" name:"HaType"`
}

func (r *ScaleOutInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleOutInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Type")
	delete(f, "NodeCount")
	delete(f, "HaType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScaleOutInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleOutInstanceResponseParams struct {
	// Process ID
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Error message
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ScaleOutInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ScaleOutInstanceResponseParams `json:"Response"`
}

func (r *ScaleOutInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleOutInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleUpInstanceRequestParams struct {
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Node specifications
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// Role (MATER/CORE). MASTER corresponds to FE, and CORE corresponds to BE.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type ScaleUpInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Node specifications
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// Role (MATER/CORE). MASTER corresponds to FE, and CORE corresponds to BE.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *ScaleUpInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleUpInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SpecName")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScaleUpInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleUpInstanceResponseParams struct {
	// Process ID
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Error message
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ScaleUpInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ScaleUpInstanceResponseParams `json:"Response"`
}

func (r *ScaleUpInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleUpInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchTags struct {
	// Tag key
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`

	// 1 means only the tag key is entered without a value, and 0 means both the key and the value are entered.
	AllValue *int64 `json:"AllValue,omitnil,omitempty" name:"AllValue"`
}

type SlowQueryRecord struct {
	// User query 
	OsUser *string `json:"OsUser,omitnil,omitempty" name:"OsUser"`

	// ID query
	InitialQueryId *string `json:"InitialQueryId,omitnil,omitempty" name:"InitialQueryId"`

	// SQL statement
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// Start time
	QueryStartTime *string `json:"QueryStartTime,omitnil,omitempty" name:"QueryStartTime"`

	// Execution duration
	DurationMs *int64 `json:"DurationMs,omitnil,omitempty" name:"DurationMs"`

	// The number of read rows
	ReadRows *int64 `json:"ReadRows,omitnil,omitempty" name:"ReadRows"`

	// Total number of read bytes
	ResultRows *int64 `json:"ResultRows,omitnil,omitempty" name:"ResultRows"`

	// Result bytes
	ResultBytes *uint64 `json:"ResultBytes,omitnil,omitempty" name:"ResultBytes"`

	// Memory
	MemoryUsage *int64 `json:"MemoryUsage,omitnil,omitempty" name:"MemoryUsage"`

	// Initial query IP
	InitialAddress *string `json:"InitialAddress,omitnil,omitempty" name:"InitialAddress"`

	// Database name
	// Note: This field may return null, indicating that no valid values can be obtained.
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Whether it is a query. 0 indicates no, and 1 indicates query statement.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsQuery *int64 `json:"IsQuery,omitnil,omitempty" name:"IsQuery"`

	// MB format of ResultBytes
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResultBytesMB *float64 `json:"ResultBytesMB,omitnil,omitempty" name:"ResultBytesMB"`

	// MemoryUsage, in MB
	// Note: This field may return null, indicating that no valid values can be obtained.
	MemoryUsageMB *float64 `json:"MemoryUsageMB,omitnil,omitempty" name:"MemoryUsageMB"`

	// DurationMs, in seconds
	// Note: This field may return null, indicating that no valid values can be obtained.
	DurationSec *float64 `json:"DurationSec,omitnil,omitempty" name:"DurationSec"`
}

type Tag struct {
	// Tag key
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}