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

package v20210622

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

// Predefined struct for user
type CreateLogExportRequestParams struct {
	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Log export start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Log export end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Log export search statement
	Query *string `json:"Query,omitempty" name:"Query"`

	// Number of logs to be exported. Maximum value: 10 million
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// Exported log sorting order by time. Valid values: asc: ascending; desc: descending. Default value: desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// Exported log data format. Valid values: json, csv. Default value: json
	Format *string `json:"Format,omitempty" name:"Format"`
}

type CreateLogExportRequest struct {
	*tchttp.BaseRequest
	
	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Log export start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Log export end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Log export search statement
	Query *string `json:"Query,omitempty" name:"Query"`

	// Number of logs to be exported. Maximum value: 10 million
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// Exported log sorting order by time. Valid values: asc: ascending; desc: descending. Default value: desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// Exported log data format. Valid values: json, csv. Default value: json
	Format *string `json:"Format,omitempty" name:"Format"`
}

func (r *CreateLogExportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLogExportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Query")
	delete(f, "Count")
	delete(f, "Order")
	delete(f, "Format")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLogExportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLogExportResponseParams struct {
	// Log export ID
	ExportID *string `json:"ExportID,omitempty" name:"ExportID"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateLogExportResponse struct {
	*tchttp.BaseResponse
	Response *CreateLogExportResponseParams `json:"Response"`
}

func (r *CreateLogExportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLogExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOfflineLogConfigRequestParams struct {
	// Unique project key for reporting
	ProjectKey *string `json:"ProjectKey,omitempty" name:"ProjectKey"`

	// Unique identifier of the user to be listened on (`aid` or `uin`)
	UniqueID *string `json:"UniqueID,omitempty" name:"UniqueID"`
}

type CreateOfflineLogConfigRequest struct {
	*tchttp.BaseRequest
	
	// Unique project key for reporting
	ProjectKey *string `json:"ProjectKey,omitempty" name:"ProjectKey"`

	// Unique identifier of the user to be listened on (`aid` or `uin`)
	UniqueID *string `json:"UniqueID,omitempty" name:"UniqueID"`
}

func (r *CreateOfflineLogConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOfflineLogConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectKey")
	delete(f, "UniqueID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOfflineLogConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOfflineLogConfigResponseParams struct {
	// API response information
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateOfflineLogConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateOfflineLogConfigResponseParams `json:"Response"`
}

func (r *CreateOfflineLogConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOfflineLogConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProjectRequestParams struct {
	// Name of the created project (required and up to 200 characters)
	Name *string `json:"Name,omitempty" name:"Name"`

	// Business system ID
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`

	// Project sampling rate (greater than or equal to 0)
	Rate *string `json:"Rate,omitempty" name:"Rate"`

	// Whether to enable aggregation
	EnableURLGroup *uint64 `json:"EnableURLGroup,omitempty" name:"EnableURLGroup"`

	// Project type (valid values: "web", "mp", "android", "ios", "node", "hippy", "weex", "viola", "rn")
	Type *string `json:"Type,omitempty" name:"Type"`

	// Repository address of the project (optional and up to 256 characters)
	Repo *string `json:"Repo,omitempty" name:"Repo"`

	// Webpage address of the project (optional and up to 256 characters)
	URL *string `json:"URL,omitempty" name:"URL"`

	// Description of the created project (optional and up to 1,000 characters)
	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

type CreateProjectRequest struct {
	*tchttp.BaseRequest
	
	// Name of the created project (required and up to 200 characters)
	Name *string `json:"Name,omitempty" name:"Name"`

	// Business system ID
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`

	// Project sampling rate (greater than or equal to 0)
	Rate *string `json:"Rate,omitempty" name:"Rate"`

	// Whether to enable aggregation
	EnableURLGroup *uint64 `json:"EnableURLGroup,omitempty" name:"EnableURLGroup"`

	// Project type (valid values: "web", "mp", "android", "ios", "node", "hippy", "weex", "viola", "rn")
	Type *string `json:"Type,omitempty" name:"Type"`

	// Repository address of the project (optional and up to 256 characters)
	Repo *string `json:"Repo,omitempty" name:"Repo"`

	// Webpage address of the project (optional and up to 256 characters)
	URL *string `json:"URL,omitempty" name:"URL"`

	// Description of the created project (optional and up to 1,000 characters)
	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

func (r *CreateProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "InstanceID")
	delete(f, "Rate")
	delete(f, "EnableURLGroup")
	delete(f, "Type")
	delete(f, "Repo")
	delete(f, "URL")
	delete(f, "Desc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProjectResponseParams struct {
	// Project ID
	ID *uint64 `json:"ID,omitempty" name:"ID"`

	// Unique project key
	Key *string `json:"Key,omitempty" name:"Key"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateProjectResponse struct {
	*tchttp.BaseResponse
	Response *CreateProjectResponseParams `json:"Response"`
}

func (r *CreateProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReleaseFileRequestParams struct {
	// Project ID
	ProjectID *int64 `json:"ProjectID,omitempty" name:"ProjectID"`

	// File information list
	Files []*ReleaseFile `json:"Files,omitempty" name:"Files"`
}

type CreateReleaseFileRequest struct {
	*tchttp.BaseRequest
	
	// Project ID
	ProjectID *int64 `json:"ProjectID,omitempty" name:"ProjectID"`

	// File information list
	Files []*ReleaseFile `json:"Files,omitempty" name:"Files"`
}

func (r *CreateReleaseFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReleaseFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectID")
	delete(f, "Files")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateReleaseFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReleaseFileResponseParams struct {
	// Call result
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateReleaseFileResponse struct {
	*tchttp.BaseResponse
	Response *CreateReleaseFileResponseParams `json:"Response"`
}

func (r *CreateReleaseFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReleaseFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStarProjectRequestParams struct {
	// Instance ID, such as taw-123
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`

	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`
}

type CreateStarProjectRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, such as taw-123
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`

	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`
}

func (r *CreateStarProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStarProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStarProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStarProjectResponseParams struct {
	// API response information
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateStarProjectResponse struct {
	*tchttp.BaseResponse
	Response *CreateStarProjectResponseParams `json:"Response"`
}

func (r *CreateStarProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStarProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTawInstanceRequestParams struct {
	// Region ID (at least greater than 0)
	AreaId *int64 `json:"AreaId,omitempty" name:"AreaId"`

	// Billing type (1: Pay-as-you-go).
	ChargeType *int64 `json:"ChargeType,omitempty" name:"ChargeType"`

	// Data retention period (at least greater than 0)
	DataRetentionDays *int64 `json:"DataRetentionDays,omitempty" name:"DataRetentionDays"`

	// Instance name (up to 255 bytes)
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Tag list
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Instance description (up to 1,024 bytes)
	InstanceDesc *string `json:"InstanceDesc,omitempty" name:"InstanceDesc"`

	// Number of data entries reported per day
	CountNum *string `json:"CountNum,omitempty" name:"CountNum"`

	// Billing for data storage
	PeriodRetain *string `json:"PeriodRetain,omitempty" name:"PeriodRetain"`

	// Instance purchase channel. Valid value: `cdn`.
	BuyingChannel *string `json:"BuyingChannel,omitempty" name:"BuyingChannel"`

	// Type of prepaid resource pack (only required for prepaid mode)
	ResourcePackageType *uint64 `json:"ResourcePackageType,omitempty" name:"ResourcePackageType"`

	// The number of prepaid resource packs (only required for prepaid mode)
	ResourcePackageNum *uint64 `json:"ResourcePackageNum,omitempty" name:"ResourcePackageNum"`
}

type CreateTawInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Region ID (at least greater than 0)
	AreaId *int64 `json:"AreaId,omitempty" name:"AreaId"`

	// Billing type (1: Pay-as-you-go).
	ChargeType *int64 `json:"ChargeType,omitempty" name:"ChargeType"`

	// Data retention period (at least greater than 0)
	DataRetentionDays *int64 `json:"DataRetentionDays,omitempty" name:"DataRetentionDays"`

	// Instance name (up to 255 bytes)
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Tag list
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Instance description (up to 1,024 bytes)
	InstanceDesc *string `json:"InstanceDesc,omitempty" name:"InstanceDesc"`

	// Number of data entries reported per day
	CountNum *string `json:"CountNum,omitempty" name:"CountNum"`

	// Billing for data storage
	PeriodRetain *string `json:"PeriodRetain,omitempty" name:"PeriodRetain"`

	// Instance purchase channel. Valid value: `cdn`.
	BuyingChannel *string `json:"BuyingChannel,omitempty" name:"BuyingChannel"`

	// Type of prepaid resource pack (only required for prepaid mode)
	ResourcePackageType *uint64 `json:"ResourcePackageType,omitempty" name:"ResourcePackageType"`

	// The number of prepaid resource packs (only required for prepaid mode)
	ResourcePackageNum *uint64 `json:"ResourcePackageNum,omitempty" name:"ResourcePackageNum"`
}

func (r *CreateTawInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTawInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AreaId")
	delete(f, "ChargeType")
	delete(f, "DataRetentionDays")
	delete(f, "InstanceName")
	delete(f, "Tags")
	delete(f, "InstanceDesc")
	delete(f, "CountNum")
	delete(f, "PeriodRetain")
	delete(f, "BuyingChannel")
	delete(f, "ResourcePackageType")
	delete(f, "ResourcePackageNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTawInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTawInstanceResponseParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// ID of prepaid order
	// Note: This field may return null, indicating that no valid values can be obtained.
	DealName *string `json:"DealName,omitempty" name:"DealName"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateTawInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateTawInstanceResponseParams `json:"Response"`
}

func (r *CreateTawInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTawInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWhitelistRequestParams struct {
	// Instance ID, such as taw-123
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`

	// Remarks
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// uin: business identifier
	WhitelistUin *string `json:"WhitelistUin,omitempty" name:"WhitelistUin"`

	// Business identifier
	Aid *string `json:"Aid,omitempty" name:"Aid"`
}

type CreateWhitelistRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, such as taw-123
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`

	// Remarks
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// uin: business identifier
	WhitelistUin *string `json:"WhitelistUin,omitempty" name:"WhitelistUin"`

	// Business identifier
	Aid *string `json:"Aid,omitempty" name:"Aid"`
}

func (r *CreateWhitelistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWhitelistRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	delete(f, "Remark")
	delete(f, "WhitelistUin")
	delete(f, "Aid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWhitelistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWhitelistResponseParams struct {
	// Message
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// Allowlist ID
	ID *uint64 `json:"ID,omitempty" name:"ID"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateWhitelistResponse struct {
	*tchttp.BaseResponse
	Response *CreateWhitelistResponseParams `json:"Response"`
}

func (r *CreateWhitelistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWhitelistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstanceRequestParams struct {
	// ID of the instance to be deleted
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DeleteInstanceRequest struct {
	*tchttp.BaseRequest
	
	// ID of the instance to be deleted
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DeleteInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstanceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteInstanceResponseParams `json:"Response"`
}

func (r *DeleteInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLogExportRequestParams struct {
	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Log export ID
	ExportID *string `json:"ExportID,omitempty" name:"ExportID"`
}

type DeleteLogExportRequest struct {
	*tchttp.BaseRequest
	
	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Log export ID
	ExportID *string `json:"ExportID,omitempty" name:"ExportID"`
}

func (r *DeleteLogExportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLogExportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	delete(f, "ExportID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLogExportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLogExportResponseParams struct {
	// Whether it is successful. If so, `success` will be returned; otherwise, `Error` rather than this parameter will be returned
	// Note: this field may return null, indicating that no valid values can be obtained.
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteLogExportResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLogExportResponseParams `json:"Response"`
}

func (r *DeleteLogExportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLogExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOfflineLogConfigRequestParams struct {
	// Unique project key for reporting
	ProjectKey *string `json:"ProjectKey,omitempty" name:"ProjectKey"`

	// Unique user identifier (uin or aid)
	UniqueID *string `json:"UniqueID,omitempty" name:"UniqueID"`
}

type DeleteOfflineLogConfigRequest struct {
	*tchttp.BaseRequest
	
	// Unique project key for reporting
	ProjectKey *string `json:"ProjectKey,omitempty" name:"ProjectKey"`

	// Unique user identifier (uin or aid)
	UniqueID *string `json:"UniqueID,omitempty" name:"UniqueID"`
}

func (r *DeleteOfflineLogConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOfflineLogConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectKey")
	delete(f, "UniqueID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteOfflineLogConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOfflineLogConfigResponseParams struct {
	// API call information
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteOfflineLogConfigResponse struct {
	*tchttp.BaseResponse
	Response *DeleteOfflineLogConfigResponseParams `json:"Response"`
}

func (r *DeleteOfflineLogConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOfflineLogConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOfflineLogRecordRequestParams struct {
	// Unique project key for reporting
	ProjectKey *string `json:"ProjectKey,omitempty" name:"ProjectKey"`

	// Offline log file ID
	FileID *string `json:"FileID,omitempty" name:"FileID"`
}

type DeleteOfflineLogRecordRequest struct {
	*tchttp.BaseRequest
	
	// Unique project key for reporting
	ProjectKey *string `json:"ProjectKey,omitempty" name:"ProjectKey"`

	// Offline log file ID
	FileID *string `json:"FileID,omitempty" name:"FileID"`
}

func (r *DeleteOfflineLogRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOfflineLogRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectKey")
	delete(f, "FileID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteOfflineLogRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOfflineLogRecordResponseParams struct {
	// API call information
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteOfflineLogRecordResponse struct {
	*tchttp.BaseResponse
	Response *DeleteOfflineLogRecordResponseParams `json:"Response"`
}

func (r *DeleteOfflineLogRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOfflineLogRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProjectRequestParams struct {
	// ID of the project to be deleted
	ID *uint64 `json:"ID,omitempty" name:"ID"`
}

type DeleteProjectRequest struct {
	*tchttp.BaseRequest
	
	// ID of the project to be deleted
	ID *uint64 `json:"ID,omitempty" name:"ID"`
}

func (r *DeleteProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProjectResponseParams struct {
	// Operation information
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteProjectResponse struct {
	*tchttp.BaseResponse
	Response *DeleteProjectResponseParams `json:"Response"`
}

func (r *DeleteProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReleaseFileRequestParams struct {
	// File ID
	ID *int64 `json:"ID,omitempty" name:"ID"`
}

type DeleteReleaseFileRequest struct {
	*tchttp.BaseRequest
	
	// File ID
	ID *int64 `json:"ID,omitempty" name:"ID"`
}

func (r *DeleteReleaseFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReleaseFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteReleaseFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReleaseFileResponseParams struct {
	// Response string of the API request
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteReleaseFileResponse struct {
	*tchttp.BaseResponse
	Response *DeleteReleaseFileResponseParams `json:"Response"`
}

func (r *DeleteReleaseFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReleaseFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStarProjectRequestParams struct {
	// Instance ID, such as taw-123
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`

	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`
}

type DeleteStarProjectRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, such as taw-123
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`

	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`
}

func (r *DeleteStarProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStarProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteStarProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStarProjectResponseParams struct {
	// Response message
	// Note: this field may return null, indicating that no valid values can be obtained.
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteStarProjectResponse struct {
	*tchttp.BaseResponse
	Response *DeleteStarProjectResponseParams `json:"Response"`
}

func (r *DeleteStarProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStarProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWhitelistRequestParams struct {
	// Instance ID
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`

	// List ID
	ID *string `json:"ID,omitempty" name:"ID"`
}

type DeleteWhitelistRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`

	// List ID
	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *DeleteWhitelistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWhitelistRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWhitelistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWhitelistResponseParams struct {
	// Success message
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteWhitelistResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWhitelistResponseParams `json:"Response"`
}

func (r *DeleteWhitelistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWhitelistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataCustomUrlRequestParams struct {
	// Start time
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// `top`: top resources view; `allcount`: performance view; `day`: 14-day data; `condition`: condition list; `pagepv`: PV view; `area`: request speed distribution; `nettype`/`version`/`platform`/`isp`/`region`/`device`/`browser`/`ext1`/`ext2`/`ext3`/`ret`/`status`/`from`/`url`/`env`: network/platform view; version view; device view; ISP view; region view; browser view; custom view, and so on.
	Type *string `json:"Type,omitempty" name:"Type"`

	// End time
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Custom 2
	ExtSecond *string `json:"ExtSecond,omitempty" name:"ExtSecond"`

	// Browser engine
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// ISP
	Isp *string `json:"Isp,omitempty" name:"Isp"`

	// Source page
	From *string `json:"From,omitempty" name:"From"`

	// Log level
	Level *string `json:"Level,omitempty" name:"Level"`

	// Brand
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// Region
	Area *string `json:"Area,omitempty" name:"Area"`

	// Version
	VersionNum *string `json:"VersionNum,omitempty" name:"VersionNum"`

	// Platform
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// Custom 3
	ExtThird *string `json:"ExtThird,omitempty" name:"ExtThird"`

	// Custom 1
	ExtFirst *string `json:"ExtFirst,omitempty" name:"ExtFirst"`

	// Network type
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// Model
	Device *string `json:"Device,omitempty" name:"Device"`

	// Whether it is outside the Chinese mainland
	IsAbroad *string `json:"IsAbroad,omitempty" name:"IsAbroad"`

	// OS
	Os *string `json:"Os,omitempty" name:"Os"`

	// Browser
	Browser *string `json:"Browser,omitempty" name:"Browser"`

	// Duration calculation method
	CostType *string `json:"CostType,omitempty" name:"CostType"`

	// Key value of the custom speed test
	Url *string `json:"Url,omitempty" name:"Url"`

	// Environment
	Env *string `json:"Env,omitempty" name:"Env"`
}

type DescribeDataCustomUrlRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// `top`: top resources view; `allcount`: performance view; `day`: 14-day data; `condition`: condition list; `pagepv`: PV view; `area`: request speed distribution; `nettype`/`version`/`platform`/`isp`/`region`/`device`/`browser`/`ext1`/`ext2`/`ext3`/`ret`/`status`/`from`/`url`/`env`: network/platform view; version view; device view; ISP view; region view; browser view; custom view, and so on.
	Type *string `json:"Type,omitempty" name:"Type"`

	// End time
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Custom 2
	ExtSecond *string `json:"ExtSecond,omitempty" name:"ExtSecond"`

	// Browser engine
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// ISP
	Isp *string `json:"Isp,omitempty" name:"Isp"`

	// Source page
	From *string `json:"From,omitempty" name:"From"`

	// Log level
	Level *string `json:"Level,omitempty" name:"Level"`

	// Brand
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// Region
	Area *string `json:"Area,omitempty" name:"Area"`

	// Version
	VersionNum *string `json:"VersionNum,omitempty" name:"VersionNum"`

	// Platform
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// Custom 3
	ExtThird *string `json:"ExtThird,omitempty" name:"ExtThird"`

	// Custom 1
	ExtFirst *string `json:"ExtFirst,omitempty" name:"ExtFirst"`

	// Network type
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// Model
	Device *string `json:"Device,omitempty" name:"Device"`

	// Whether it is outside the Chinese mainland
	IsAbroad *string `json:"IsAbroad,omitempty" name:"IsAbroad"`

	// OS
	Os *string `json:"Os,omitempty" name:"Os"`

	// Browser
	Browser *string `json:"Browser,omitempty" name:"Browser"`

	// Duration calculation method
	CostType *string `json:"CostType,omitempty" name:"CostType"`

	// Key value of the custom speed test
	Url *string `json:"Url,omitempty" name:"Url"`

	// Environment
	Env *string `json:"Env,omitempty" name:"Env"`
}

func (r *DescribeDataCustomUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataCustomUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "CostType")
	delete(f, "Url")
	delete(f, "Env")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataCustomUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataCustomUrlResponseParams struct {
	// Returned value
	Result *string `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDataCustomUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataCustomUrlResponseParams `json:"Response"`
}

func (r *DescribeDataCustomUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataCustomUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataEventUrlRequestParams struct {
	// Start time
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// `allcount`: performance view; `day`: 14-day data; `condition`: condition list; `ckuv`: UV trend; `ckpv`: PV trend; `nettype`/`version`/`platform`/`isp`/`region`/`device`/`browser`/`ext1`/`ext2`/`ext3`/`ret`/`status`/`from`/`url`/`env`: network/platform view; version view; device view; ISP view; region view; browser view; custom view, and so on.
	Type *string `json:"Type,omitempty" name:"Type"`

	// End time
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Custom 2
	ExtSecond *string `json:"ExtSecond,omitempty" name:"ExtSecond"`

	// Browser engine
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// ISP
	Isp *string `json:"Isp,omitempty" name:"Isp"`

	// Source page
	From *string `json:"From,omitempty" name:"From"`

	// Log level
	Level *string `json:"Level,omitempty" name:"Level"`

	// Brand
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// Region
	Area *string `json:"Area,omitempty" name:"Area"`

	// Version
	VersionNum *string `json:"VersionNum,omitempty" name:"VersionNum"`

	// Platform
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// Custom 3
	ExtThird *string `json:"ExtThird,omitempty" name:"ExtThird"`

	// Custom 1
	ExtFirst *string `json:"ExtFirst,omitempty" name:"ExtFirst"`

	// Network type
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// Model
	Device *string `json:"Device,omitempty" name:"Device"`

	// Whether it is outside the Chinese mainland
	IsAbroad *string `json:"IsAbroad,omitempty" name:"IsAbroad"`

	// OS
	Os *string `json:"Os,omitempty" name:"Os"`

	// Browser
	Browser *string `json:"Browser,omitempty" name:"Browser"`

	// Filter
	Name *string `json:"Name,omitempty" name:"Name"`

	// Environment
	Env *string `json:"Env,omitempty" name:"Env"`
}

type DescribeDataEventUrlRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// `allcount`: performance view; `day`: 14-day data; `condition`: condition list; `ckuv`: UV trend; `ckpv`: PV trend; `nettype`/`version`/`platform`/`isp`/`region`/`device`/`browser`/`ext1`/`ext2`/`ext3`/`ret`/`status`/`from`/`url`/`env`: network/platform view; version view; device view; ISP view; region view; browser view; custom view, and so on.
	Type *string `json:"Type,omitempty" name:"Type"`

	// End time
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Custom 2
	ExtSecond *string `json:"ExtSecond,omitempty" name:"ExtSecond"`

	// Browser engine
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// ISP
	Isp *string `json:"Isp,omitempty" name:"Isp"`

	// Source page
	From *string `json:"From,omitempty" name:"From"`

	// Log level
	Level *string `json:"Level,omitempty" name:"Level"`

	// Brand
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// Region
	Area *string `json:"Area,omitempty" name:"Area"`

	// Version
	VersionNum *string `json:"VersionNum,omitempty" name:"VersionNum"`

	// Platform
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// Custom 3
	ExtThird *string `json:"ExtThird,omitempty" name:"ExtThird"`

	// Custom 1
	ExtFirst *string `json:"ExtFirst,omitempty" name:"ExtFirst"`

	// Network type
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// Model
	Device *string `json:"Device,omitempty" name:"Device"`

	// Whether it is outside the Chinese mainland
	IsAbroad *string `json:"IsAbroad,omitempty" name:"IsAbroad"`

	// OS
	Os *string `json:"Os,omitempty" name:"Os"`

	// Browser
	Browser *string `json:"Browser,omitempty" name:"Browser"`

	// Filter
	Name *string `json:"Name,omitempty" name:"Name"`

	// Environment
	Env *string `json:"Env,omitempty" name:"Env"`
}

func (r *DescribeDataEventUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataEventUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "Name")
	delete(f, "Env")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataEventUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataEventUrlResponseParams struct {
	// Returned value
	Result *string `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDataEventUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataEventUrlResponseParams `json:"Response"`
}

func (r *DescribeDataEventUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataEventUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataFetchProjectRequestParams struct {
	// Start time
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// `allcount`: performance view; `day`: 14-day data; `condition`: condition list; `area`: request speed distribution; `nettype`/`version`/`platform`/`isp`/`region`/`device`/`browser`/`ext1`/`ext2`/`ext3`/`ret`/`status`/`from`/`url`/`env`: network/platform view; version view; device view; ISP view; region view; browser view; custom view, and so on.
	Type *string `json:"Type,omitempty" name:"Type"`

	// End time
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Custom 2
	ExtSecond *string `json:"ExtSecond,omitempty" name:"ExtSecond"`

	// Browser engine
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// ISP
	Isp *string `json:"Isp,omitempty" name:"Isp"`

	// Source page
	From *string `json:"From,omitempty" name:"From"`

	// Log level
	Level *string `json:"Level,omitempty" name:"Level"`

	// Brand
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// Region
	Area *string `json:"Area,omitempty" name:"Area"`

	// Version
	VersionNum *string `json:"VersionNum,omitempty" name:"VersionNum"`

	// Platform
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// Custom 3
	ExtThird *string `json:"ExtThird,omitempty" name:"ExtThird"`

	// Custom 1
	ExtFirst *string `json:"ExtFirst,omitempty" name:"ExtFirst"`

	// Network type
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// Model
	Device *string `json:"Device,omitempty" name:"Device"`

	// Whether it is outside the Chinese mainland
	IsAbroad *string `json:"IsAbroad,omitempty" name:"IsAbroad"`

	// OS
	Os *string `json:"Os,omitempty" name:"Os"`

	// Browser
	Browser *string `json:"Browser,omitempty" name:"Browser"`

	// Duration calculation method
	CostType *string `json:"CostType,omitempty" name:"CostType"`

	// Source
	Url *string `json:"Url,omitempty" name:"Url"`

	// Environment
	Env *string `json:"Env,omitempty" name:"Env"`

	// HTTP status code.
	Status *string `json:"Status,omitempty" name:"Status"`

	// retcode
	Ret *string `json:"Ret,omitempty" name:"Ret"`
}

type DescribeDataFetchProjectRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// `allcount`: performance view; `day`: 14-day data; `condition`: condition list; `area`: request speed distribution; `nettype`/`version`/`platform`/`isp`/`region`/`device`/`browser`/`ext1`/`ext2`/`ext3`/`ret`/`status`/`from`/`url`/`env`: network/platform view; version view; device view; ISP view; region view; browser view; custom view, and so on.
	Type *string `json:"Type,omitempty" name:"Type"`

	// End time
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Custom 2
	ExtSecond *string `json:"ExtSecond,omitempty" name:"ExtSecond"`

	// Browser engine
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// ISP
	Isp *string `json:"Isp,omitempty" name:"Isp"`

	// Source page
	From *string `json:"From,omitempty" name:"From"`

	// Log level
	Level *string `json:"Level,omitempty" name:"Level"`

	// Brand
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// Region
	Area *string `json:"Area,omitempty" name:"Area"`

	// Version
	VersionNum *string `json:"VersionNum,omitempty" name:"VersionNum"`

	// Platform
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// Custom 3
	ExtThird *string `json:"ExtThird,omitempty" name:"ExtThird"`

	// Custom 1
	ExtFirst *string `json:"ExtFirst,omitempty" name:"ExtFirst"`

	// Network type
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// Model
	Device *string `json:"Device,omitempty" name:"Device"`

	// Whether it is outside the Chinese mainland
	IsAbroad *string `json:"IsAbroad,omitempty" name:"IsAbroad"`

	// OS
	Os *string `json:"Os,omitempty" name:"Os"`

	// Browser
	Browser *string `json:"Browser,omitempty" name:"Browser"`

	// Duration calculation method
	CostType *string `json:"CostType,omitempty" name:"CostType"`

	// Source
	Url *string `json:"Url,omitempty" name:"Url"`

	// Environment
	Env *string `json:"Env,omitempty" name:"Env"`

	// HTTP status code.
	Status *string `json:"Status,omitempty" name:"Status"`

	// retcode
	Ret *string `json:"Ret,omitempty" name:"Ret"`
}

func (r *DescribeDataFetchProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataFetchProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "CostType")
	delete(f, "Url")
	delete(f, "Env")
	delete(f, "Status")
	delete(f, "Ret")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataFetchProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataFetchProjectResponseParams struct {
	// Returned value
	Result *string `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDataFetchProjectResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataFetchProjectResponseParams `json:"Response"`
}

func (r *DescribeDataFetchProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataFetchProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataFetchUrlInfoRequestParams struct {
	// Start time
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// Type
	Type *string `json:"Type,omitempty" name:"Type"`

	// End time
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Custom 2
	ExtSecond *string `json:"ExtSecond,omitempty" name:"ExtSecond"`

	// Browser engine
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// ISP
	Isp *string `json:"Isp,omitempty" name:"Isp"`

	// Source page
	From *string `json:"From,omitempty" name:"From"`

	// Log level
	Level *string `json:"Level,omitempty" name:"Level"`

	// Brand
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// Region
	Area *string `json:"Area,omitempty" name:"Area"`

	// Version
	VersionNum *string `json:"VersionNum,omitempty" name:"VersionNum"`

	// Platform
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// Custom 3
	ExtThird *string `json:"ExtThird,omitempty" name:"ExtThird"`

	// Custom 1
	ExtFirst *string `json:"ExtFirst,omitempty" name:"ExtFirst"`

	// Network type
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// Model
	Device *string `json:"Device,omitempty" name:"Device"`

	// Whether it is outside the Chinese mainland
	IsAbroad *string `json:"IsAbroad,omitempty" name:"IsAbroad"`

	// OS
	Os *string `json:"Os,omitempty" name:"Os"`

	// Browser
	Browser *string `json:"Browser,omitempty" name:"Browser"`

	// Duration calculation method
	CostType *string `json:"CostType,omitempty" name:"CostType"`

	// Source
	Url *string `json:"Url,omitempty" name:"Url"`

	// Environment
	Env *string `json:"Env,omitempty" name:"Env"`
}

type DescribeDataFetchUrlInfoRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// Type
	Type *string `json:"Type,omitempty" name:"Type"`

	// End time
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Custom 2
	ExtSecond *string `json:"ExtSecond,omitempty" name:"ExtSecond"`

	// Browser engine
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// ISP
	Isp *string `json:"Isp,omitempty" name:"Isp"`

	// Source page
	From *string `json:"From,omitempty" name:"From"`

	// Log level
	Level *string `json:"Level,omitempty" name:"Level"`

	// Brand
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// Region
	Area *string `json:"Area,omitempty" name:"Area"`

	// Version
	VersionNum *string `json:"VersionNum,omitempty" name:"VersionNum"`

	// Platform
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// Custom 3
	ExtThird *string `json:"ExtThird,omitempty" name:"ExtThird"`

	// Custom 1
	ExtFirst *string `json:"ExtFirst,omitempty" name:"ExtFirst"`

	// Network type
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// Model
	Device *string `json:"Device,omitempty" name:"Device"`

	// Whether it is outside the Chinese mainland
	IsAbroad *string `json:"IsAbroad,omitempty" name:"IsAbroad"`

	// OS
	Os *string `json:"Os,omitempty" name:"Os"`

	// Browser
	Browser *string `json:"Browser,omitempty" name:"Browser"`

	// Duration calculation method
	CostType *string `json:"CostType,omitempty" name:"CostType"`

	// Source
	Url *string `json:"Url,omitempty" name:"Url"`

	// Environment
	Env *string `json:"Env,omitempty" name:"Env"`
}

func (r *DescribeDataFetchUrlInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataFetchUrlInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "CostType")
	delete(f, "Url")
	delete(f, "Env")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataFetchUrlInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataFetchUrlInfoResponseParams struct {
	// Returned value
	Result *string `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDataFetchUrlInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataFetchUrlInfoResponseParams `json:"Response"`
}

func (r *DescribeDataFetchUrlInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataFetchUrlInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataFetchUrlRequestParams struct {
	// Start time
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// `allcount`: performance view; `day`: 14-day data; `count40x`: HTTP status codes 40X view; `count50x`: HTTP status codes 50X view; `count5xand4x`: HTTP status codes 40X∑50X view; `top`: top resources view; `nettype`/`version`/`platform`/`isp`/`region`/`device`/`browser`/`ext1`/`ext2`/`ext3`/`ret`/`status`/`from`/`url`/`env`: network/platform view; version view; device view; ISP view; region view; browser view; custom view, and so on.
	Type *string `json:"Type,omitempty" name:"Type"`

	// End time
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Custom 2
	ExtSecond *string `json:"ExtSecond,omitempty" name:"ExtSecond"`

	// Browser engine
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// ISP
	Isp *string `json:"Isp,omitempty" name:"Isp"`

	// Source page
	From *string `json:"From,omitempty" name:"From"`

	// Log level
	Level *string `json:"Level,omitempty" name:"Level"`

	// Brand
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// Region
	Area *string `json:"Area,omitempty" name:"Area"`

	// Version
	VersionNum *string `json:"VersionNum,omitempty" name:"VersionNum"`

	// Platform
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// Custom 3
	ExtThird *string `json:"ExtThird,omitempty" name:"ExtThird"`

	// Custom 1
	ExtFirst *string `json:"ExtFirst,omitempty" name:"ExtFirst"`

	// Network type
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// Model
	Device *string `json:"Device,omitempty" name:"Device"`

	// Whether it is outside the Chinese mainland
	IsAbroad *string `json:"IsAbroad,omitempty" name:"IsAbroad"`

	// OS
	Os *string `json:"Os,omitempty" name:"Os"`

	// Browser
	Browser *string `json:"Browser,omitempty" name:"Browser"`

	// Duration calculation method
	CostType *string `json:"CostType,omitempty" name:"CostType"`

	// Source
	Url *string `json:"Url,omitempty" name:"Url"`

	// Environment
	Env *string `json:"Env,omitempty" name:"Env"`

	// HTTP status code.
	Status *string `json:"Status,omitempty" name:"Status"`

	// retcode
	Ret *string `json:"Ret,omitempty" name:"Ret"`

	// Network status
	NetStatus *string `json:"NetStatus,omitempty" name:"NetStatus"`
}

type DescribeDataFetchUrlRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// `allcount`: performance view; `day`: 14-day data; `count40x`: HTTP status codes 40X view; `count50x`: HTTP status codes 50X view; `count5xand4x`: HTTP status codes 40X∑50X view; `top`: top resources view; `nettype`/`version`/`platform`/`isp`/`region`/`device`/`browser`/`ext1`/`ext2`/`ext3`/`ret`/`status`/`from`/`url`/`env`: network/platform view; version view; device view; ISP view; region view; browser view; custom view, and so on.
	Type *string `json:"Type,omitempty" name:"Type"`

	// End time
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Custom 2
	ExtSecond *string `json:"ExtSecond,omitempty" name:"ExtSecond"`

	// Browser engine
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// ISP
	Isp *string `json:"Isp,omitempty" name:"Isp"`

	// Source page
	From *string `json:"From,omitempty" name:"From"`

	// Log level
	Level *string `json:"Level,omitempty" name:"Level"`

	// Brand
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// Region
	Area *string `json:"Area,omitempty" name:"Area"`

	// Version
	VersionNum *string `json:"VersionNum,omitempty" name:"VersionNum"`

	// Platform
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// Custom 3
	ExtThird *string `json:"ExtThird,omitempty" name:"ExtThird"`

	// Custom 1
	ExtFirst *string `json:"ExtFirst,omitempty" name:"ExtFirst"`

	// Network type
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// Model
	Device *string `json:"Device,omitempty" name:"Device"`

	// Whether it is outside the Chinese mainland
	IsAbroad *string `json:"IsAbroad,omitempty" name:"IsAbroad"`

	// OS
	Os *string `json:"Os,omitempty" name:"Os"`

	// Browser
	Browser *string `json:"Browser,omitempty" name:"Browser"`

	// Duration calculation method
	CostType *string `json:"CostType,omitempty" name:"CostType"`

	// Source
	Url *string `json:"Url,omitempty" name:"Url"`

	// Environment
	Env *string `json:"Env,omitempty" name:"Env"`

	// HTTP status code.
	Status *string `json:"Status,omitempty" name:"Status"`

	// retcode
	Ret *string `json:"Ret,omitempty" name:"Ret"`

	// Network status
	NetStatus *string `json:"NetStatus,omitempty" name:"NetStatus"`
}

func (r *DescribeDataFetchUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataFetchUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "CostType")
	delete(f, "Url")
	delete(f, "Env")
	delete(f, "Status")
	delete(f, "Ret")
	delete(f, "NetStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataFetchUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataFetchUrlResponseParams struct {
	// Returned value
	Result *string `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDataFetchUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataFetchUrlResponseParams `json:"Response"`
}

func (r *DescribeDataFetchUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataFetchUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataLogUrlInfoRequestParams struct {
	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Timestamp
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// Timestamp
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeDataLogUrlInfoRequest struct {
	*tchttp.BaseRequest
	
	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Timestamp
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// Timestamp
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeDataLogUrlInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataLogUrlInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataLogUrlInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataLogUrlInfoResponseParams struct {
	// Response string
	Result *string `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDataLogUrlInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataLogUrlInfoResponseParams `json:"Response"`
}

func (r *DescribeDataLogUrlInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataLogUrlInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataLogUrlStatisticsRequestParams struct {
	// Start time
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// `analysis`: exception analysis; `compare`: exception list comparison; `allcount`: performance view; `condition`: condition list; `nettype`/`version`/`platform`/`isp`/`region`/`device`/`browser`/`ext1`/`ext2`/`ext3`/`ret`/`status`/`from`/`url`/`env`: network/platform view; version view; device view; ISP view; region view; browser view; custom view, and so on.
	Type *string `json:"Type,omitempty" name:"Type"`

	// End time
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Custom 2
	ExtSecond *string `json:"ExtSecond,omitempty" name:"ExtSecond"`

	// Browser engine
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// ISP
	Isp *string `json:"Isp,omitempty" name:"Isp"`

	// Source page
	From *string `json:"From,omitempty" name:"From"`

	// Log level
	Level *string `json:"Level,omitempty" name:"Level"`

	// Brand
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// Region
	Area *string `json:"Area,omitempty" name:"Area"`

	// Version
	VersionNum *string `json:"VersionNum,omitempty" name:"VersionNum"`

	// Platform
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// Custom 3
	ExtThird *string `json:"ExtThird,omitempty" name:"ExtThird"`

	// Custom 1
	ExtFirst *string `json:"ExtFirst,omitempty" name:"ExtFirst"`

	// Network type
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// Model
	Device *string `json:"Device,omitempty" name:"Device"`

	// Whether it is outside the Chinese mainland
	IsAbroad *string `json:"IsAbroad,omitempty" name:"IsAbroad"`

	// OS
	Os *string `json:"Os,omitempty" name:"Os"`

	// Browser
	Browser *string `json:"Browser,omitempty" name:"Browser"`

	// Environment
	Env *string `json:"Env,omitempty" name:"Env"`
}

type DescribeDataLogUrlStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// `analysis`: exception analysis; `compare`: exception list comparison; `allcount`: performance view; `condition`: condition list; `nettype`/`version`/`platform`/`isp`/`region`/`device`/`browser`/`ext1`/`ext2`/`ext3`/`ret`/`status`/`from`/`url`/`env`: network/platform view; version view; device view; ISP view; region view; browser view; custom view, and so on.
	Type *string `json:"Type,omitempty" name:"Type"`

	// End time
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Custom 2
	ExtSecond *string `json:"ExtSecond,omitempty" name:"ExtSecond"`

	// Browser engine
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// ISP
	Isp *string `json:"Isp,omitempty" name:"Isp"`

	// Source page
	From *string `json:"From,omitempty" name:"From"`

	// Log level
	Level *string `json:"Level,omitempty" name:"Level"`

	// Brand
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// Region
	Area *string `json:"Area,omitempty" name:"Area"`

	// Version
	VersionNum *string `json:"VersionNum,omitempty" name:"VersionNum"`

	// Platform
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// Custom 3
	ExtThird *string `json:"ExtThird,omitempty" name:"ExtThird"`

	// Custom 1
	ExtFirst *string `json:"ExtFirst,omitempty" name:"ExtFirst"`

	// Network type
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// Model
	Device *string `json:"Device,omitempty" name:"Device"`

	// Whether it is outside the Chinese mainland
	IsAbroad *string `json:"IsAbroad,omitempty" name:"IsAbroad"`

	// OS
	Os *string `json:"Os,omitempty" name:"Os"`

	// Browser
	Browser *string `json:"Browser,omitempty" name:"Browser"`

	// Environment
	Env *string `json:"Env,omitempty" name:"Env"`
}

func (r *DescribeDataLogUrlStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataLogUrlStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "Env")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataLogUrlStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataLogUrlStatisticsResponseParams struct {
	// Returned value
	Result *string `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDataLogUrlStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataLogUrlStatisticsResponseParams `json:"Response"`
}

func (r *DescribeDataLogUrlStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataLogUrlStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataPerformancePageRequestParams struct {
	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Start time
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// `pagepv`: PV view; `allcount`: performance view; `falls`: page loading waterfall plot; `samp`: FMP, `day`: 14-day data, `nettype`: network/platform view; `performance`: top underperformed pages view; `version`/`platform`/`isp`/`region`/`device`/`browser`/`ext1`/`ext2`/`ext3`/`ret`/`status`/`from`/`url`/`env`: version view; device view; ISP view; region view; browser view, and so on.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Log level
	Level *string `json:"Level,omitempty" name:"Level"`

	// ISP
	Isp *string `json:"Isp,omitempty" name:"Isp"`

	// Region
	Area *string `json:"Area,omitempty" name:"Area"`

	// Network type
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// Platform
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// Model
	Device *string `json:"Device,omitempty" name:"Device"`

	// Version
	VersionNum *string `json:"VersionNum,omitempty" name:"VersionNum"`

	// Custom 1
	ExtFirst *string `json:"ExtFirst,omitempty" name:"ExtFirst"`

	// Custom 2
	ExtSecond *string `json:"ExtSecond,omitempty" name:"ExtSecond"`

	// Custom 3
	ExtThird *string `json:"ExtThird,omitempty" name:"ExtThird"`

	// Whether it is outside the Chinese mainland
	IsAbroad *string `json:"IsAbroad,omitempty" name:"IsAbroad"`

	// Browser
	Browser *string `json:"Browser,omitempty" name:"Browser"`

	// OS
	Os *string `json:"Os,omitempty" name:"Os"`

	// Browser engine
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// Brand
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// Source page
	From *string `json:"From,omitempty" name:"From"`

	// Duration calculation method
	CostType *string `json:"CostType,omitempty" name:"CostType"`

	// Environment variable
	Env *string `json:"Env,omitempty" name:"Env"`

	// Network status
	NetStatus *string `json:"NetStatus,omitempty" name:"NetStatus"`
}

type DescribeDataPerformancePageRequest struct {
	*tchttp.BaseRequest
	
	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Start time
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// `pagepv`: PV view; `allcount`: performance view; `falls`: page loading waterfall plot; `samp`: FMP, `day`: 14-day data, `nettype`: network/platform view; `performance`: top underperformed pages view; `version`/`platform`/`isp`/`region`/`device`/`browser`/`ext1`/`ext2`/`ext3`/`ret`/`status`/`from`/`url`/`env`: version view; device view; ISP view; region view; browser view, and so on.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Log level
	Level *string `json:"Level,omitempty" name:"Level"`

	// ISP
	Isp *string `json:"Isp,omitempty" name:"Isp"`

	// Region
	Area *string `json:"Area,omitempty" name:"Area"`

	// Network type
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// Platform
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// Model
	Device *string `json:"Device,omitempty" name:"Device"`

	// Version
	VersionNum *string `json:"VersionNum,omitempty" name:"VersionNum"`

	// Custom 1
	ExtFirst *string `json:"ExtFirst,omitempty" name:"ExtFirst"`

	// Custom 2
	ExtSecond *string `json:"ExtSecond,omitempty" name:"ExtSecond"`

	// Custom 3
	ExtThird *string `json:"ExtThird,omitempty" name:"ExtThird"`

	// Whether it is outside the Chinese mainland
	IsAbroad *string `json:"IsAbroad,omitempty" name:"IsAbroad"`

	// Browser
	Browser *string `json:"Browser,omitempty" name:"Browser"`

	// OS
	Os *string `json:"Os,omitempty" name:"Os"`

	// Browser engine
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// Brand
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// Source page
	From *string `json:"From,omitempty" name:"From"`

	// Duration calculation method
	CostType *string `json:"CostType,omitempty" name:"CostType"`

	// Environment variable
	Env *string `json:"Env,omitempty" name:"Env"`

	// Network status
	NetStatus *string `json:"NetStatus,omitempty" name:"NetStatus"`
}

func (r *DescribeDataPerformancePageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataPerformancePageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Type")
	delete(f, "Level")
	delete(f, "Isp")
	delete(f, "Area")
	delete(f, "NetType")
	delete(f, "Platform")
	delete(f, "Device")
	delete(f, "VersionNum")
	delete(f, "ExtFirst")
	delete(f, "ExtSecond")
	delete(f, "ExtThird")
	delete(f, "IsAbroad")
	delete(f, "Browser")
	delete(f, "Os")
	delete(f, "Engine")
	delete(f, "Brand")
	delete(f, "From")
	delete(f, "CostType")
	delete(f, "Env")
	delete(f, "NetStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataPerformancePageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataPerformancePageResponseParams struct {
	// Returned value
	Result *string `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDataPerformancePageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataPerformancePageResponseParams `json:"Response"`
}

func (r *DescribeDataPerformancePageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataPerformancePageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataPerformanceProjectRequestParams struct {
	// Start time
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// `allcount`; performance view; `falls`: page loading waterfall plot; `samp`: FMP, `day`: 14-day data, `nettype`: network/platform view; `performance`: top underperformed pages view; `condition`: condition list; `area`: request speed distribution; `version`/`platform`/`isp`/`region`/`device`/`browser`/`ext1`/`ext2`/`ext3`/`ret`/`status`/`from`/`url`/`env`: version view; device view; ISP view; region view; browser view, and so on.
	Type *string `json:"Type,omitempty" name:"Type"`

	// End time
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Custom 2
	ExtSecond *string `json:"ExtSecond,omitempty" name:"ExtSecond"`

	// Browser engine
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// ISP
	Isp *string `json:"Isp,omitempty" name:"Isp"`

	// Source page
	From *string `json:"From,omitempty" name:"From"`

	// Log level
	Level *string `json:"Level,omitempty" name:"Level"`

	// Brand
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// Region
	Area *string `json:"Area,omitempty" name:"Area"`

	// Version
	VersionNum *string `json:"VersionNum,omitempty" name:"VersionNum"`

	// Platform
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// Custom 3
	ExtThird *string `json:"ExtThird,omitempty" name:"ExtThird"`

	// Custom 1
	ExtFirst *string `json:"ExtFirst,omitempty" name:"ExtFirst"`

	// Network type
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// Model
	Device *string `json:"Device,omitempty" name:"Device"`

	// Whether it is outside the Chinese mainland
	IsAbroad *string `json:"IsAbroad,omitempty" name:"IsAbroad"`

	// OS
	Os *string `json:"Os,omitempty" name:"Os"`

	// Browser
	Browser *string `json:"Browser,omitempty" name:"Browser"`

	// Duration calculation
	CostType *string `json:"CostType,omitempty" name:"CostType"`

	// Environment
	Env *string `json:"Env,omitempty" name:"Env"`
}

type DescribeDataPerformanceProjectRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// `allcount`; performance view; `falls`: page loading waterfall plot; `samp`: FMP, `day`: 14-day data, `nettype`: network/platform view; `performance`: top underperformed pages view; `condition`: condition list; `area`: request speed distribution; `version`/`platform`/`isp`/`region`/`device`/`browser`/`ext1`/`ext2`/`ext3`/`ret`/`status`/`from`/`url`/`env`: version view; device view; ISP view; region view; browser view, and so on.
	Type *string `json:"Type,omitempty" name:"Type"`

	// End time
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Custom 2
	ExtSecond *string `json:"ExtSecond,omitempty" name:"ExtSecond"`

	// Browser engine
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// ISP
	Isp *string `json:"Isp,omitempty" name:"Isp"`

	// Source page
	From *string `json:"From,omitempty" name:"From"`

	// Log level
	Level *string `json:"Level,omitempty" name:"Level"`

	// Brand
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// Region
	Area *string `json:"Area,omitempty" name:"Area"`

	// Version
	VersionNum *string `json:"VersionNum,omitempty" name:"VersionNum"`

	// Platform
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// Custom 3
	ExtThird *string `json:"ExtThird,omitempty" name:"ExtThird"`

	// Custom 1
	ExtFirst *string `json:"ExtFirst,omitempty" name:"ExtFirst"`

	// Network type
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// Model
	Device *string `json:"Device,omitempty" name:"Device"`

	// Whether it is outside the Chinese mainland
	IsAbroad *string `json:"IsAbroad,omitempty" name:"IsAbroad"`

	// OS
	Os *string `json:"Os,omitempty" name:"Os"`

	// Browser
	Browser *string `json:"Browser,omitempty" name:"Browser"`

	// Duration calculation
	CostType *string `json:"CostType,omitempty" name:"CostType"`

	// Environment
	Env *string `json:"Env,omitempty" name:"Env"`
}

func (r *DescribeDataPerformanceProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataPerformanceProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "CostType")
	delete(f, "Env")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataPerformanceProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataPerformanceProjectResponseParams struct {
	// Returned value
	Result *string `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDataPerformanceProjectResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataPerformanceProjectResponseParams `json:"Response"`
}

func (r *DescribeDataPerformanceProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataPerformanceProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataPvUrlInfoRequestParams struct {
	// Start time
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// Type
	Type *string `json:"Type,omitempty" name:"Type"`

	// End time
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Custom 2
	ExtSecond *string `json:"ExtSecond,omitempty" name:"ExtSecond"`

	// Browser engine
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// ISP
	Isp *string `json:"Isp,omitempty" name:"Isp"`

	// Source page
	From *string `json:"From,omitempty" name:"From"`

	// Log level
	Level *string `json:"Level,omitempty" name:"Level"`

	// Brand
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// Region
	Area *string `json:"Area,omitempty" name:"Area"`

	// Version
	VersionNum *string `json:"VersionNum,omitempty" name:"VersionNum"`

	// Platform
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// Custom 3
	ExtThird *string `json:"ExtThird,omitempty" name:"ExtThird"`

	// Custom 1
	ExtFirst *string `json:"ExtFirst,omitempty" name:"ExtFirst"`

	// Network type
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// Model
	Device *string `json:"Device,omitempty" name:"Device"`

	// Whether it is outside the Chinese mainland
	IsAbroad *string `json:"IsAbroad,omitempty" name:"IsAbroad"`

	// OS
	Os *string `json:"Os,omitempty" name:"Os"`

	// Browser
	Browser *string `json:"Browser,omitempty" name:"Browser"`

	// Environment
	Env *string `json:"Env,omitempty" name:"Env"`
}

type DescribeDataPvUrlInfoRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// Type
	Type *string `json:"Type,omitempty" name:"Type"`

	// End time
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Custom 2
	ExtSecond *string `json:"ExtSecond,omitempty" name:"ExtSecond"`

	// Browser engine
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// ISP
	Isp *string `json:"Isp,omitempty" name:"Isp"`

	// Source page
	From *string `json:"From,omitempty" name:"From"`

	// Log level
	Level *string `json:"Level,omitempty" name:"Level"`

	// Brand
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// Region
	Area *string `json:"Area,omitempty" name:"Area"`

	// Version
	VersionNum *string `json:"VersionNum,omitempty" name:"VersionNum"`

	// Platform
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// Custom 3
	ExtThird *string `json:"ExtThird,omitempty" name:"ExtThird"`

	// Custom 1
	ExtFirst *string `json:"ExtFirst,omitempty" name:"ExtFirst"`

	// Network type
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// Model
	Device *string `json:"Device,omitempty" name:"Device"`

	// Whether it is outside the Chinese mainland
	IsAbroad *string `json:"IsAbroad,omitempty" name:"IsAbroad"`

	// OS
	Os *string `json:"Os,omitempty" name:"Os"`

	// Browser
	Browser *string `json:"Browser,omitempty" name:"Browser"`

	// Environment
	Env *string `json:"Env,omitempty" name:"Env"`
}

func (r *DescribeDataPvUrlInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataPvUrlInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "Env")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataPvUrlInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataPvUrlInfoResponseParams struct {
	// Returned value
	Result *string `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDataPvUrlInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataPvUrlInfoResponseParams `json:"Response"`
}

func (r *DescribeDataPvUrlInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataPvUrlInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataPvUrlStatisticsRequestParams struct {
	// Start time
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// `allcount`; performance view; `day`: 14-day data, `vp`: performance; `ckuv`: UV; `ckpv`: PV; `condition`: condition list; `nettype`/`version`/`platform`/`isp`/`region`/`device`/`browser`/`ext1`/`ext2`/`ext3`/`ret`/`status`/`from`/`url`/`env`: network/platform view; version view; device view; ISP view; region view; browser view; custom view, and so on.
	Type *string `json:"Type,omitempty" name:"Type"`

	// End time
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Custom 2
	ExtSecond *string `json:"ExtSecond,omitempty" name:"ExtSecond"`

	// Browser engine
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// ISP
	Isp *string `json:"Isp,omitempty" name:"Isp"`

	// Source page
	From *string `json:"From,omitempty" name:"From"`

	// Log level
	Level *string `json:"Level,omitempty" name:"Level"`

	// Brand
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// Region
	Area *string `json:"Area,omitempty" name:"Area"`

	// Version
	VersionNum *string `json:"VersionNum,omitempty" name:"VersionNum"`

	// Platform
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// Custom 3
	ExtThird *string `json:"ExtThird,omitempty" name:"ExtThird"`

	// Custom 1
	ExtFirst *string `json:"ExtFirst,omitempty" name:"ExtFirst"`

	// Network type
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// Model
	Device *string `json:"Device,omitempty" name:"Device"`

	// Whether it is outside the Chinese mainland
	IsAbroad *string `json:"IsAbroad,omitempty" name:"IsAbroad"`

	// OS
	Os *string `json:"Os,omitempty" name:"Os"`

	// Browser
	Browser *string `json:"Browser,omitempty" name:"Browser"`

	// Environment
	Env *string `json:"Env,omitempty" name:"Env"`
}

type DescribeDataPvUrlStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// `allcount`; performance view; `day`: 14-day data, `vp`: performance; `ckuv`: UV; `ckpv`: PV; `condition`: condition list; `nettype`/`version`/`platform`/`isp`/`region`/`device`/`browser`/`ext1`/`ext2`/`ext3`/`ret`/`status`/`from`/`url`/`env`: network/platform view; version view; device view; ISP view; region view; browser view; custom view, and so on.
	Type *string `json:"Type,omitempty" name:"Type"`

	// End time
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Custom 2
	ExtSecond *string `json:"ExtSecond,omitempty" name:"ExtSecond"`

	// Browser engine
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// ISP
	Isp *string `json:"Isp,omitempty" name:"Isp"`

	// Source page
	From *string `json:"From,omitempty" name:"From"`

	// Log level
	Level *string `json:"Level,omitempty" name:"Level"`

	// Brand
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// Region
	Area *string `json:"Area,omitempty" name:"Area"`

	// Version
	VersionNum *string `json:"VersionNum,omitempty" name:"VersionNum"`

	// Platform
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// Custom 3
	ExtThird *string `json:"ExtThird,omitempty" name:"ExtThird"`

	// Custom 1
	ExtFirst *string `json:"ExtFirst,omitempty" name:"ExtFirst"`

	// Network type
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// Model
	Device *string `json:"Device,omitempty" name:"Device"`

	// Whether it is outside the Chinese mainland
	IsAbroad *string `json:"IsAbroad,omitempty" name:"IsAbroad"`

	// OS
	Os *string `json:"Os,omitempty" name:"Os"`

	// Browser
	Browser *string `json:"Browser,omitempty" name:"Browser"`

	// Environment
	Env *string `json:"Env,omitempty" name:"Env"`
}

func (r *DescribeDataPvUrlStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataPvUrlStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "Env")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataPvUrlStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataPvUrlStatisticsResponseParams struct {
	// Returned value
	Result *string `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDataPvUrlStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataPvUrlStatisticsResponseParams `json:"Response"`
}

func (r *DescribeDataPvUrlStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataPvUrlStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataReportCountRequestParams struct {
	// Start time
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Reporting type
	ReportType *string `json:"ReportType,omitempty" name:"ReportType"`

	// Instance ID
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

type DescribeDataReportCountRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Reporting type
	ReportType *string `json:"ReportType,omitempty" name:"ReportType"`

	// Instance ID
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

func (r *DescribeDataReportCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataReportCountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ReportType")
	delete(f, "InstanceID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataReportCountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataReportCountResponseParams struct {
	// Returned value
	Result *string `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDataReportCountResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataReportCountResponseParams `json:"Response"`
}

func (r *DescribeDataReportCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataReportCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataRequestParams struct {
	// Query string
	Query *string `json:"Query,omitempty" name:"Query"`

	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`
}

type DescribeDataRequest struct {
	*tchttp.BaseRequest
	
	// Query string
	Query *string `json:"Query,omitempty" name:"Query"`

	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`
}

func (r *DescribeDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Query")
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataResponseParams struct {
	// Response string
	Result *string `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataResponseParams `json:"Response"`
}

func (r *DescribeDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataSetUrlStatisticsRequestParams struct {
	// Start time
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// `allcount`: performance view; `data`: mini program; `component`: mini program-related components; `day`: 14-day data; `nettype`: network/platform view; `performance`: top underperformed pages view; `version`/`platform`/`isp`/`region`/`device`/`browser`/`ext1`/`ext2`/`ext3`/`ret`/`status`/`from`/`url`/`env`: version view; device view; ISP view; region view; browser view, and so on.
	Type *string `json:"Type,omitempty" name:"Type"`

	// End time
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Custom 2
	ExtSecond *string `json:"ExtSecond,omitempty" name:"ExtSecond"`

	// Browser engine
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// ISP
	Isp *string `json:"Isp,omitempty" name:"Isp"`

	// Source page
	From *string `json:"From,omitempty" name:"From"`

	// Log level
	Level *string `json:"Level,omitempty" name:"Level"`

	// Brand
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// Region
	Area *string `json:"Area,omitempty" name:"Area"`

	// Version
	VersionNum *string `json:"VersionNum,omitempty" name:"VersionNum"`

	// Platform
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// Custom 3
	ExtThird *string `json:"ExtThird,omitempty" name:"ExtThird"`

	// Custom 1
	ExtFirst *string `json:"ExtFirst,omitempty" name:"ExtFirst"`

	// Network type
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// Model
	Device *string `json:"Device,omitempty" name:"Device"`

	// Whether it is outside the Chinese mainland
	IsAbroad *string `json:"IsAbroad,omitempty" name:"IsAbroad"`

	// OS
	Os *string `json:"Os,omitempty" name:"Os"`

	// Browser
	Browser *string `json:"Browser,omitempty" name:"Browser"`

	// Duration calculation
	CostType *string `json:"CostType,omitempty" name:"CostType"`

	// Environment
	Env *string `json:"Env,omitempty" name:"Env"`

	// The obtained package.
	PackageType *string `json:"PackageType,omitempty" name:"PackageType"`
}

type DescribeDataSetUrlStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// `allcount`: performance view; `data`: mini program; `component`: mini program-related components; `day`: 14-day data; `nettype`: network/platform view; `performance`: top underperformed pages view; `version`/`platform`/`isp`/`region`/`device`/`browser`/`ext1`/`ext2`/`ext3`/`ret`/`status`/`from`/`url`/`env`: version view; device view; ISP view; region view; browser view, and so on.
	Type *string `json:"Type,omitempty" name:"Type"`

	// End time
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Custom 2
	ExtSecond *string `json:"ExtSecond,omitempty" name:"ExtSecond"`

	// Browser engine
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// ISP
	Isp *string `json:"Isp,omitempty" name:"Isp"`

	// Source page
	From *string `json:"From,omitempty" name:"From"`

	// Log level
	Level *string `json:"Level,omitempty" name:"Level"`

	// Brand
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// Region
	Area *string `json:"Area,omitempty" name:"Area"`

	// Version
	VersionNum *string `json:"VersionNum,omitempty" name:"VersionNum"`

	// Platform
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// Custom 3
	ExtThird *string `json:"ExtThird,omitempty" name:"ExtThird"`

	// Custom 1
	ExtFirst *string `json:"ExtFirst,omitempty" name:"ExtFirst"`

	// Network type
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// Model
	Device *string `json:"Device,omitempty" name:"Device"`

	// Whether it is outside the Chinese mainland
	IsAbroad *string `json:"IsAbroad,omitempty" name:"IsAbroad"`

	// OS
	Os *string `json:"Os,omitempty" name:"Os"`

	// Browser
	Browser *string `json:"Browser,omitempty" name:"Browser"`

	// Duration calculation
	CostType *string `json:"CostType,omitempty" name:"CostType"`

	// Environment
	Env *string `json:"Env,omitempty" name:"Env"`

	// The obtained package.
	PackageType *string `json:"PackageType,omitempty" name:"PackageType"`
}

func (r *DescribeDataSetUrlStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataSetUrlStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "CostType")
	delete(f, "Env")
	delete(f, "PackageType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataSetUrlStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataSetUrlStatisticsResponseParams struct {
	// Returned value
	Result *string `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDataSetUrlStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataSetUrlStatisticsResponseParams `json:"Response"`
}

func (r *DescribeDataSetUrlStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataSetUrlStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataStaticProjectRequestParams struct {
	// Start time
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// `allcount`: performance view; `day`: 14-day data; `condition`: condition list; `area`: request speed distribution; `nettype`/`version`/`platform`/`isp`/`region`/`device`/`browser`/`ext1`/`ext2`/`ext3`/`ret`/`status`/`from`/`url`/`env`: network/platform view; version view; device view; ISP view; region view; browser view; custom view, and so on.
	Type *string `json:"Type,omitempty" name:"Type"`

	// End time
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Custom 2
	ExtSecond *string `json:"ExtSecond,omitempty" name:"ExtSecond"`

	// Browser engine
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// ISP
	Isp *string `json:"Isp,omitempty" name:"Isp"`

	// Source page
	From *string `json:"From,omitempty" name:"From"`

	// Log level
	Level *string `json:"Level,omitempty" name:"Level"`

	// Brand
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// Region
	Area *string `json:"Area,omitempty" name:"Area"`

	// Version
	VersionNum *string `json:"VersionNum,omitempty" name:"VersionNum"`

	// Platform
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// Custom 3
	ExtThird *string `json:"ExtThird,omitempty" name:"ExtThird"`

	// Custom 1
	ExtFirst *string `json:"ExtFirst,omitempty" name:"ExtFirst"`

	// Network type
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// Model
	Device *string `json:"Device,omitempty" name:"Device"`

	// Whether it is outside the Chinese mainland
	IsAbroad *string `json:"IsAbroad,omitempty" name:"IsAbroad"`

	// OS
	Os *string `json:"Os,omitempty" name:"Os"`

	// Browser
	Browser *string `json:"Browser,omitempty" name:"Browser"`

	// Duration calculation
	CostType *string `json:"CostType,omitempty" name:"CostType"`

	// Source
	Url []*string `json:"Url,omitempty" name:"Url"`

	// Environment
	Env *string `json:"Env,omitempty" name:"Env"`
}

type DescribeDataStaticProjectRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// `allcount`: performance view; `day`: 14-day data; `condition`: condition list; `area`: request speed distribution; `nettype`/`version`/`platform`/`isp`/`region`/`device`/`browser`/`ext1`/`ext2`/`ext3`/`ret`/`status`/`from`/`url`/`env`: network/platform view; version view; device view; ISP view; region view; browser view; custom view, and so on.
	Type *string `json:"Type,omitempty" name:"Type"`

	// End time
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Custom 2
	ExtSecond *string `json:"ExtSecond,omitempty" name:"ExtSecond"`

	// Browser engine
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// ISP
	Isp *string `json:"Isp,omitempty" name:"Isp"`

	// Source page
	From *string `json:"From,omitempty" name:"From"`

	// Log level
	Level *string `json:"Level,omitempty" name:"Level"`

	// Brand
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// Region
	Area *string `json:"Area,omitempty" name:"Area"`

	// Version
	VersionNum *string `json:"VersionNum,omitempty" name:"VersionNum"`

	// Platform
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// Custom 3
	ExtThird *string `json:"ExtThird,omitempty" name:"ExtThird"`

	// Custom 1
	ExtFirst *string `json:"ExtFirst,omitempty" name:"ExtFirst"`

	// Network type
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// Model
	Device *string `json:"Device,omitempty" name:"Device"`

	// Whether it is outside the Chinese mainland
	IsAbroad *string `json:"IsAbroad,omitempty" name:"IsAbroad"`

	// OS
	Os *string `json:"Os,omitempty" name:"Os"`

	// Browser
	Browser *string `json:"Browser,omitempty" name:"Browser"`

	// Duration calculation
	CostType *string `json:"CostType,omitempty" name:"CostType"`

	// Source
	Url []*string `json:"Url,omitempty" name:"Url"`

	// Environment
	Env *string `json:"Env,omitempty" name:"Env"`
}

func (r *DescribeDataStaticProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataStaticProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "CostType")
	delete(f, "Url")
	delete(f, "Env")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataStaticProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataStaticProjectResponseParams struct {
	// Returned value
	Result *string `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDataStaticProjectResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataStaticProjectResponseParams `json:"Response"`
}

func (r *DescribeDataStaticProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataStaticProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataStaticResourceRequestParams struct {
	// Start time
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// `top`: top resources view; `count40x`: HTTP status codes 40X view; `nettype`/`version`/`platform`/`isp`/`region`/`device`/`browser`/`ext1`/`ext2`/`ext3`/`ret`/`status`/`from`/`url`/`env`: network/platform view; version view; device view; ISP view; region view; browser view; custom view, and so on.
	Type *string `json:"Type,omitempty" name:"Type"`

	// End time
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Custom 2
	ExtSecond *string `json:"ExtSecond,omitempty" name:"ExtSecond"`

	// Browser engine
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// ISP
	Isp *string `json:"Isp,omitempty" name:"Isp"`

	// Source page
	From *string `json:"From,omitempty" name:"From"`

	// Log level
	Level *string `json:"Level,omitempty" name:"Level"`

	// Brand
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// Region
	Area *string `json:"Area,omitempty" name:"Area"`

	// Version
	VersionNum *string `json:"VersionNum,omitempty" name:"VersionNum"`

	// Platform
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// Custom 3
	ExtThird *string `json:"ExtThird,omitempty" name:"ExtThird"`

	// Custom 1
	ExtFirst *string `json:"ExtFirst,omitempty" name:"ExtFirst"`

	// Network type
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// Model
	Device *string `json:"Device,omitempty" name:"Device"`

	// Whether it is outside the Chinese mainland
	IsAbroad *string `json:"IsAbroad,omitempty" name:"IsAbroad"`

	// OS
	Os *string `json:"Os,omitempty" name:"Os"`

	// Browser
	Browser *string `json:"Browser,omitempty" name:"Browser"`

	// Duration calculation method
	CostType *string `json:"CostType,omitempty" name:"CostType"`

	// Source
	Url *string `json:"Url,omitempty" name:"Url"`

	// Environment
	Env *string `json:"Env,omitempty" name:"Env"`
}

type DescribeDataStaticResourceRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// `top`: top resources view; `count40x`: HTTP status codes 40X view; `nettype`/`version`/`platform`/`isp`/`region`/`device`/`browser`/`ext1`/`ext2`/`ext3`/`ret`/`status`/`from`/`url`/`env`: network/platform view; version view; device view; ISP view; region view; browser view; custom view, and so on.
	Type *string `json:"Type,omitempty" name:"Type"`

	// End time
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Custom 2
	ExtSecond *string `json:"ExtSecond,omitempty" name:"ExtSecond"`

	// Browser engine
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// ISP
	Isp *string `json:"Isp,omitempty" name:"Isp"`

	// Source page
	From *string `json:"From,omitempty" name:"From"`

	// Log level
	Level *string `json:"Level,omitempty" name:"Level"`

	// Brand
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// Region
	Area *string `json:"Area,omitempty" name:"Area"`

	// Version
	VersionNum *string `json:"VersionNum,omitempty" name:"VersionNum"`

	// Platform
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// Custom 3
	ExtThird *string `json:"ExtThird,omitempty" name:"ExtThird"`

	// Custom 1
	ExtFirst *string `json:"ExtFirst,omitempty" name:"ExtFirst"`

	// Network type
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// Model
	Device *string `json:"Device,omitempty" name:"Device"`

	// Whether it is outside the Chinese mainland
	IsAbroad *string `json:"IsAbroad,omitempty" name:"IsAbroad"`

	// OS
	Os *string `json:"Os,omitempty" name:"Os"`

	// Browser
	Browser *string `json:"Browser,omitempty" name:"Browser"`

	// Duration calculation method
	CostType *string `json:"CostType,omitempty" name:"CostType"`

	// Source
	Url *string `json:"Url,omitempty" name:"Url"`

	// Environment
	Env *string `json:"Env,omitempty" name:"Env"`
}

func (r *DescribeDataStaticResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataStaticResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "CostType")
	delete(f, "Url")
	delete(f, "Env")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataStaticResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataStaticResourceResponseParams struct {
	// Returned value
	Result *string `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDataStaticResourceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataStaticResourceResponseParams `json:"Response"`
}

func (r *DescribeDataStaticResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataStaticResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataStaticUrlRequestParams struct {
	// Start time
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// `pagepv`: page view; `nettype`/`version`/`platform`/`isp`/`region`/`device`/`browser`/`ext1`/`ext2`/`ext3`/`ret`/`status`/`from`/`url`/`env`: network/platform view; version view; device view; ISP view; region view; browser view; custom view, and so on.
	Type *string `json:"Type,omitempty" name:"Type"`

	// End time
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Custom 2
	ExtSecond *string `json:"ExtSecond,omitempty" name:"ExtSecond"`

	// Browser engine
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// ISP
	Isp *string `json:"Isp,omitempty" name:"Isp"`

	// Source page
	From *string `json:"From,omitempty" name:"From"`

	// Log level
	Level *string `json:"Level,omitempty" name:"Level"`

	// Brand
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// Region
	Area *string `json:"Area,omitempty" name:"Area"`

	// Version
	VersionNum *string `json:"VersionNum,omitempty" name:"VersionNum"`

	// Platform
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// Custom 3
	ExtThird *string `json:"ExtThird,omitempty" name:"ExtThird"`

	// Custom 1
	ExtFirst *string `json:"ExtFirst,omitempty" name:"ExtFirst"`

	// Network type
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// Model
	Device *string `json:"Device,omitempty" name:"Device"`

	// Whether it is outside the Chinese mainland
	IsAbroad *string `json:"IsAbroad,omitempty" name:"IsAbroad"`

	// OS
	Os *string `json:"Os,omitempty" name:"Os"`

	// Browser
	Browser *string `json:"Browser,omitempty" name:"Browser"`

	// Duration calculation method
	CostType *string `json:"CostType,omitempty" name:"CostType"`

	// Source
	Url *string `json:"Url,omitempty" name:"Url"`

	// Environment
	Env *string `json:"Env,omitempty" name:"Env"`
}

type DescribeDataStaticUrlRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// `pagepv`: page view; `nettype`/`version`/`platform`/`isp`/`region`/`device`/`browser`/`ext1`/`ext2`/`ext3`/`ret`/`status`/`from`/`url`/`env`: network/platform view; version view; device view; ISP view; region view; browser view; custom view, and so on.
	Type *string `json:"Type,omitempty" name:"Type"`

	// End time
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Custom 2
	ExtSecond *string `json:"ExtSecond,omitempty" name:"ExtSecond"`

	// Browser engine
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// ISP
	Isp *string `json:"Isp,omitempty" name:"Isp"`

	// Source page
	From *string `json:"From,omitempty" name:"From"`

	// Log level
	Level *string `json:"Level,omitempty" name:"Level"`

	// Brand
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// Region
	Area *string `json:"Area,omitempty" name:"Area"`

	// Version
	VersionNum *string `json:"VersionNum,omitempty" name:"VersionNum"`

	// Platform
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// Custom 3
	ExtThird *string `json:"ExtThird,omitempty" name:"ExtThird"`

	// Custom 1
	ExtFirst *string `json:"ExtFirst,omitempty" name:"ExtFirst"`

	// Network type
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// Model
	Device *string `json:"Device,omitempty" name:"Device"`

	// Whether it is outside the Chinese mainland
	IsAbroad *string `json:"IsAbroad,omitempty" name:"IsAbroad"`

	// OS
	Os *string `json:"Os,omitempty" name:"Os"`

	// Browser
	Browser *string `json:"Browser,omitempty" name:"Browser"`

	// Duration calculation method
	CostType *string `json:"CostType,omitempty" name:"CostType"`

	// Source
	Url *string `json:"Url,omitempty" name:"Url"`

	// Environment
	Env *string `json:"Env,omitempty" name:"Env"`
}

func (r *DescribeDataStaticUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataStaticUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "CostType")
	delete(f, "Url")
	delete(f, "Env")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataStaticUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataStaticUrlResponseParams struct {
	// Returned value
	Result *string `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDataStaticUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataStaticUrlResponseParams `json:"Response"`
}

func (r *DescribeDataStaticUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataStaticUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataWebVitalsPageRequestParams struct {
	// Start time
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Custom 2
	ExtSecond *string `json:"ExtSecond,omitempty" name:"ExtSecond"`

	// Browser engine
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// ISP
	Isp *string `json:"Isp,omitempty" name:"Isp"`

	// Source page
	From *string `json:"From,omitempty" name:"From"`

	// Log level
	Level *string `json:"Level,omitempty" name:"Level"`

	// No type yet
	Type *string `json:"Type,omitempty" name:"Type"`

	// Brand
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// Region
	Area *string `json:"Area,omitempty" name:"Area"`

	// Version
	VersionNum *string `json:"VersionNum,omitempty" name:"VersionNum"`

	// Platform
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// Custom 3
	ExtThird *string `json:"ExtThird,omitempty" name:"ExtThird"`

	// Custom 1
	ExtFirst *string `json:"ExtFirst,omitempty" name:"ExtFirst"`

	// Network type
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// Model
	Device *string `json:"Device,omitempty" name:"Device"`

	// Whether it is outside the Chinese mainland
	IsAbroad *string `json:"IsAbroad,omitempty" name:"IsAbroad"`

	// OS
	Os *string `json:"Os,omitempty" name:"Os"`

	// Browser
	Browser *string `json:"Browser,omitempty" name:"Browser"`

	// Duration calculation
	CostType *string `json:"CostType,omitempty" name:"CostType"`

	// Environment
	Env *string `json:"Env,omitempty" name:"Env"`
}

type DescribeDataWebVitalsPageRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Custom 2
	ExtSecond *string `json:"ExtSecond,omitempty" name:"ExtSecond"`

	// Browser engine
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// ISP
	Isp *string `json:"Isp,omitempty" name:"Isp"`

	// Source page
	From *string `json:"From,omitempty" name:"From"`

	// Log level
	Level *string `json:"Level,omitempty" name:"Level"`

	// No type yet
	Type *string `json:"Type,omitempty" name:"Type"`

	// Brand
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// Region
	Area *string `json:"Area,omitempty" name:"Area"`

	// Version
	VersionNum *string `json:"VersionNum,omitempty" name:"VersionNum"`

	// Platform
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// Custom 3
	ExtThird *string `json:"ExtThird,omitempty" name:"ExtThird"`

	// Custom 1
	ExtFirst *string `json:"ExtFirst,omitempty" name:"ExtFirst"`

	// Network type
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// Model
	Device *string `json:"Device,omitempty" name:"Device"`

	// Whether it is outside the Chinese mainland
	IsAbroad *string `json:"IsAbroad,omitempty" name:"IsAbroad"`

	// OS
	Os *string `json:"Os,omitempty" name:"Os"`

	// Browser
	Browser *string `json:"Browser,omitempty" name:"Browser"`

	// Duration calculation
	CostType *string `json:"CostType,omitempty" name:"CostType"`

	// Environment
	Env *string `json:"Env,omitempty" name:"Env"`
}

func (r *DescribeDataWebVitalsPageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataWebVitalsPageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Type")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "CostType")
	delete(f, "Env")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataWebVitalsPageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataWebVitalsPageResponseParams struct {
	// Returned value
	Result *string `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDataWebVitalsPageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataWebVitalsPageResponseParams `json:"Response"`
}

func (r *DescribeDataWebVitalsPageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataWebVitalsPageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeErrorRequestParams struct {
	// Date
	Date *string `json:"Date,omitempty" name:"Date"`

	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`
}

type DescribeErrorRequest struct {
	*tchttp.BaseRequest
	
	// Date
	Date *string `json:"Date,omitempty" name:"Date"`

	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`
}

func (r *DescribeErrorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeErrorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Date")
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeErrorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeErrorResponseParams struct {
	// Content
	Content *string `json:"Content,omitempty" name:"Content"`

	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeErrorResponse struct {
	*tchttp.BaseResponse
	Response *DescribeErrorResponseParams `json:"Response"`
}

func (r *DescribeErrorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeErrorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogExportsRequestParams struct {
	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`
}

type DescribeLogExportsRequest struct {
	*tchttp.BaseRequest
	
	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`
}

func (r *DescribeLogExportsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogExportsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogExportsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogExportsResponseParams struct {
	// List of log export records
	LogExportSet []*LogExport `json:"LogExportSet,omitempty" name:"LogExportSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLogExportsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLogExportsResponseParams `json:"Response"`
}

func (r *DescribeLogExportsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogExportsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogListRequestParams struct {
	// Sorting order (required). Valid values: `desc`, `asc`.
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// This parameter is required. Valid values: `searchlog`, `histogram`.
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// Project ID (required)
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Start time (required)
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The number of raw logs returned for a single query. This parameter is required. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Context, which is used to load more logs. Pass through the last `Context` value returned to get more log content (up to 10,000 raw logs). It will expire after 1 hour
	Context *string `json:"Context,omitempty" name:"Context"`

	// Query statement, which is required and can contain up to 4,096 characters, such as "id:120001 AND type:\"log\"".
	Query *string `json:"Query,omitempty" name:"Query"`

	// End time (required)
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeLogListRequest struct {
	*tchttp.BaseRequest
	
	// Sorting order (required). Valid values: `desc`, `asc`.
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// This parameter is required. Valid values: `searchlog`, `histogram`.
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// Project ID (required)
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Start time (required)
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The number of raw logs returned for a single query. This parameter is required. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Context, which is used to load more logs. Pass through the last `Context` value returned to get more log content (up to 10,000 raw logs). It will expire after 1 hour
	Context *string `json:"Context,omitempty" name:"Context"`

	// Query statement, which is required and can contain up to 4,096 characters, such as "id:120001 AND type:\"log\"".
	Query *string `json:"Query,omitempty" name:"Query"`

	// End time (required)
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeLogListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Sort")
	delete(f, "ActionType")
	delete(f, "ID")
	delete(f, "StartTime")
	delete(f, "Limit")
	delete(f, "Context")
	delete(f, "Query")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogListResponseParams struct {
	// Response string
	Result *string `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLogListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLogListResponseParams `json:"Response"`
}

func (r *DescribeLogListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOfflineLogConfigsRequestParams struct {
	// Unique project key for reporting
	ProjectKey *string `json:"ProjectKey,omitempty" name:"ProjectKey"`
}

type DescribeOfflineLogConfigsRequest struct {
	*tchttp.BaseRequest
	
	// Unique project key for reporting
	ProjectKey *string `json:"ProjectKey,omitempty" name:"ProjectKey"`
}

func (r *DescribeOfflineLogConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOfflineLogConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOfflineLogConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOfflineLogConfigsResponseParams struct {
	// API call information
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// Array of unique user identifiers
	UniqueIDSet []*string `json:"UniqueIDSet,omitempty" name:"UniqueIDSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeOfflineLogConfigsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOfflineLogConfigsResponseParams `json:"Response"`
}

func (r *DescribeOfflineLogConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOfflineLogConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOfflineLogRecordsRequestParams struct {
	// Unique project key for reporting
	ProjectKey *string `json:"ProjectKey,omitempty" name:"ProjectKey"`
}

type DescribeOfflineLogRecordsRequest struct {
	*tchttp.BaseRequest
	
	// Unique project key for reporting
	ProjectKey *string `json:"ProjectKey,omitempty" name:"ProjectKey"`
}

func (r *DescribeOfflineLogRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOfflineLogRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOfflineLogRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOfflineLogRecordsResponseParams struct {
	// API call information
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// Array of record IDs
	RecordSet []*string `json:"RecordSet,omitempty" name:"RecordSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeOfflineLogRecordsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOfflineLogRecordsResponseParams `json:"Response"`
}

func (r *DescribeOfflineLogRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOfflineLogRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOfflineLogsRequestParams struct {
	// Unique project key for reporting
	ProjectKey *string `json:"ProjectKey,omitempty" name:"ProjectKey"`

	// List of offline log file IDs
	FileIDs []*string `json:"FileIDs,omitempty" name:"FileIDs"`
}

type DescribeOfflineLogsRequest struct {
	*tchttp.BaseRequest
	
	// Unique project key for reporting
	ProjectKey *string `json:"ProjectKey,omitempty" name:"ProjectKey"`

	// List of offline log file IDs
	FileIDs []*string `json:"FileIDs,omitempty" name:"FileIDs"`
}

func (r *DescribeOfflineLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOfflineLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectKey")
	delete(f, "FileIDs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOfflineLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOfflineLogsResponseParams struct {
	// API call response
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// Log list
	LogSet []*string `json:"LogSet,omitempty" name:"LogSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeOfflineLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOfflineLogsResponseParams `json:"Response"`
}

func (r *DescribeOfflineLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOfflineLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectLimitsRequestParams struct {
	// Project ID
	ProjectID *int64 `json:"ProjectID,omitempty" name:"ProjectID"`
}

type DescribeProjectLimitsRequest struct {
	*tchttp.BaseRequest
	
	// Project ID
	ProjectID *int64 `json:"ProjectID,omitempty" name:"ProjectID"`
}

func (r *DescribeProjectLimitsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectLimitsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectLimitsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectLimitsResponseParams struct {
	// Array of reporting rates
	ProjectLimitSet []*ProjectLimit `json:"ProjectLimitSet,omitempty" name:"ProjectLimitSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeProjectLimitsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProjectLimitsResponseParams `json:"Response"`
}

func (r *DescribeProjectLimitsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectLimitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectsRequestParams struct {
	// Number of items per page (integer)
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Page number (integer)
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Filter parameter. Pass in {"Name": "IsDemo", "Values":["1"]} for the demo mode.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// This parameter has been disused. You need to indicate whether the demo mode is used in `Filters`.
	IsDemo *int64 `json:"IsDemo,omitempty" name:"IsDemo"`
}

type DescribeProjectsRequest struct {
	*tchttp.BaseRequest
	
	// Number of items per page (integer)
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Page number (integer)
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Filter parameter. Pass in {"Name": "IsDemo", "Values":["1"]} for the demo mode.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// This parameter has been disused. You need to indicate whether the demo mode is used in `Filters`.
	IsDemo *int64 `json:"IsDemo,omitempty" name:"IsDemo"`
}

func (r *DescribeProjectsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "IsDemo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectsResponseParams struct {
	// Total number of items in the list
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Project list
	ProjectSet []*RumProject `json:"ProjectSet,omitempty" name:"ProjectSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeProjectsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProjectsResponseParams `json:"Response"`
}

func (r *DescribeProjectsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePvListRequestParams struct {
	// ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Get day:d (leave this parameter empty if to get min)
	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`
}

type DescribePvListRequest struct {
	*tchttp.BaseRequest
	
	// ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Get day:d (leave this parameter empty if to get min)
	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`
}

func (r *DescribePvListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePvListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "EndTime")
	delete(f, "StartTime")
	delete(f, "Dimension")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePvListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePvListResponseParams struct {
	// PV list
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProjectPvSet []*RumPvInfo `json:"ProjectPvSet,omitempty" name:"ProjectPvSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePvListResponse struct {
	*tchttp.BaseResponse
	Response *DescribePvListResponseParams `json:"Response"`
}

func (r *DescribePvListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePvListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReleaseFileSignRequestParams struct {
	// Timeout period. If it is not set, it will be 5 minutes by default
	Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`
}

type DescribeReleaseFileSignRequest struct {
	*tchttp.BaseRequest
	
	// Timeout period. If it is not set, it will be 5 minutes by default
	Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`
}

func (r *DescribeReleaseFileSignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReleaseFileSignRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Timeout")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReleaseFileSignRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReleaseFileSignResponseParams struct {
	// Temporary key
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`

	// Temporary key ID
	SecretID *string `json:"SecretID,omitempty" name:"SecretID"`

	// Temporary key token
	SessionToken *string `json:"SessionToken,omitempty" name:"SessionToken"`

	// Start timestamp
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// Expiration timestamp
	ExpiredTime *int64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeReleaseFileSignResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReleaseFileSignResponseParams `json:"Response"`
}

func (r *DescribeReleaseFileSignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReleaseFileSignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReleaseFilesRequestParams struct {
	// Project ID
	ProjectID *int64 `json:"ProjectID,omitempty" name:"ProjectID"`

	// File version
	FileVersion *string `json:"FileVersion,omitempty" name:"FileVersion"`
}

type DescribeReleaseFilesRequest struct {
	*tchttp.BaseRequest
	
	// Project ID
	ProjectID *int64 `json:"ProjectID,omitempty" name:"ProjectID"`

	// File version
	FileVersion *string `json:"FileVersion,omitempty" name:"FileVersion"`
}

func (r *DescribeReleaseFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReleaseFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectID")
	delete(f, "FileVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReleaseFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReleaseFilesResponseParams struct {
	// File information list
	Files []*ReleaseFile `json:"Files,omitempty" name:"Files"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeReleaseFilesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReleaseFilesResponseParams `json:"Response"`
}

func (r *DescribeReleaseFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReleaseFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRumGroupLogRequestParams struct {
	// Sorting order (required). Valid values: `desc`, `asc`.
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Start time (required)
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The number of raw logs returned for a single query. This parameter is required. Maximum value: `100`.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Page number
	Page *int64 `json:"Page,omitempty" name:"Page"`

	// Query statement, which is required and can contain up to 4,096 characters.
	Query *string `json:"Query,omitempty" name:"Query"`

	// End time (required)
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Project ID (required)
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Aggregate field
	GroupField *string `json:"GroupField,omitempty" name:"GroupField"`
}

type DescribeRumGroupLogRequest struct {
	*tchttp.BaseRequest
	
	// Sorting order (required). Valid values: `desc`, `asc`.
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Start time (required)
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The number of raw logs returned for a single query. This parameter is required. Maximum value: `100`.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Page number
	Page *int64 `json:"Page,omitempty" name:"Page"`

	// Query statement, which is required and can contain up to 4,096 characters.
	Query *string `json:"Query,omitempty" name:"Query"`

	// End time (required)
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Project ID (required)
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Aggregate field
	GroupField *string `json:"GroupField,omitempty" name:"GroupField"`
}

func (r *DescribeRumGroupLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRumGroupLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrderBy")
	delete(f, "StartTime")
	delete(f, "Limit")
	delete(f, "Page")
	delete(f, "Query")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "GroupField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRumGroupLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRumGroupLogResponseParams struct {
	// Response string
	Result *string `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRumGroupLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRumGroupLogResponseParams `json:"Response"`
}

func (r *DescribeRumGroupLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRumGroupLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRumLogListRequestParams struct {
	// Sorting order (required). Valid values: `desc`, `asc`.
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Start time (required)
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The number of raw logs returned for a single query. This parameter is required. Maximum value: `100`.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Page number
	Page *int64 `json:"Page,omitempty" name:"Page"`

	// Query statement, which is required and can contain up to 4,096 characters.
	Query *string `json:"Query,omitempty" name:"Query"`

	// End time (required)
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Project ID (required)
	ID *int64 `json:"ID,omitempty" name:"ID"`
}

type DescribeRumLogListRequest struct {
	*tchttp.BaseRequest
	
	// Sorting order (required). Valid values: `desc`, `asc`.
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Start time (required)
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The number of raw logs returned for a single query. This parameter is required. Maximum value: `100`.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Page number
	Page *int64 `json:"Page,omitempty" name:"Page"`

	// Query statement, which is required and can contain up to 4,096 characters.
	Query *string `json:"Query,omitempty" name:"Query"`

	// End time (required)
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Project ID (required)
	ID *int64 `json:"ID,omitempty" name:"ID"`
}

func (r *DescribeRumLogListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRumLogListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrderBy")
	delete(f, "StartTime")
	delete(f, "Limit")
	delete(f, "Page")
	delete(f, "Query")
	delete(f, "EndTime")
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRumLogListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRumLogListResponseParams struct {
	// Response string
	Result *string `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRumLogListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRumLogListResponseParams `json:"Response"`
}

func (r *DescribeRumLogListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRumLogListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRumStatsLogListRequestParams struct {
	// Start time (required)
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The number of raw logs returned for a single query. This parameter is required. Maximum value: `100`.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Query statement, which is required and can contain up to 4,096 characters.
	Query *string `json:"Query,omitempty" name:"Query"`

	// End time (required)
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Project ID (required)
	ID *int64 `json:"ID,omitempty" name:"ID"`
}

type DescribeRumStatsLogListRequest struct {
	*tchttp.BaseRequest
	
	// Start time (required)
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The number of raw logs returned for a single query. This parameter is required. Maximum value: `100`.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Query statement, which is required and can contain up to 4,096 characters.
	Query *string `json:"Query,omitempty" name:"Query"`

	// End time (required)
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Project ID (required)
	ID *int64 `json:"ID,omitempty" name:"ID"`
}

func (r *DescribeRumStatsLogListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRumStatsLogListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Limit")
	delete(f, "Query")
	delete(f, "EndTime")
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRumStatsLogListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRumStatsLogListResponseParams struct {
	// Response string
	Result *string `json:"Result,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRumStatsLogListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRumStatsLogListResponseParams `json:"Response"`
}

func (r *DescribeRumStatsLogListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRumStatsLogListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScoresRequestParams struct {
	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// This parameter has been disused.
	IsDemo *int64 `json:"IsDemo,omitempty" name:"IsDemo"`
}

type DescribeScoresRequest struct {
	*tchttp.BaseRequest
	
	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// This parameter has been disused.
	IsDemo *int64 `json:"IsDemo,omitempty" name:"IsDemo"`
}

func (r *DescribeScoresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScoresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EndTime")
	delete(f, "StartTime")
	delete(f, "ID")
	delete(f, "IsDemo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScoresRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScoresResponseParams struct {
	// Array
	ScoreSet []*ScoreInfo `json:"ScoreSet,omitempty" name:"ScoreSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeScoresResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScoresResponseParams `json:"Response"`
}

func (r *DescribeScoresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScoresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTawAreasRequestParams struct {
	// Region ID
	AreaIds []*int64 `json:"AreaIds,omitempty" name:"AreaIds"`

	// Region key
	AreaKeys []*string `json:"AreaKeys,omitempty" name:"AreaKeys"`

	// Pagination limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Region status (1: valid; 2: invalid)
	AreaStatuses []*int64 `json:"AreaStatuses,omitempty" name:"AreaStatuses"`

	// Pagination offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeTawAreasRequest struct {
	*tchttp.BaseRequest
	
	// Region ID
	AreaIds []*int64 `json:"AreaIds,omitempty" name:"AreaIds"`

	// Region key
	AreaKeys []*string `json:"AreaKeys,omitempty" name:"AreaKeys"`

	// Pagination limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Region status (1: valid; 2: invalid)
	AreaStatuses []*int64 `json:"AreaStatuses,omitempty" name:"AreaStatuses"`

	// Pagination offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeTawAreasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTawAreasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AreaIds")
	delete(f, "AreaKeys")
	delete(f, "Limit")
	delete(f, "AreaStatuses")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTawAreasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTawAreasResponseParams struct {
	// Total number of regions
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Region list
	AreaSet []*RumAreaInfo `json:"AreaSet,omitempty" name:"AreaSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTawAreasResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTawAreasResponseParams `json:"Response"`
}

func (r *DescribeTawAreasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTawAreasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUvListRequestParams struct {
	// ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Get day:d min:m
	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`
}

type DescribeUvListRequest struct {
	*tchttp.BaseRequest
	
	// ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Get day:d min:m
	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`
}

func (r *DescribeUvListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUvListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "EndTime")
	delete(f, "StartTime")
	delete(f, "Dimension")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUvListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUvListResponseParams struct {
	// UV list
	ProjectUvSet []*RumUvInfo `json:"ProjectUvSet,omitempty" name:"ProjectUvSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUvListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUvListResponseParams `json:"Response"`
}

func (r *DescribeUvListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUvListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWhitelistsRequestParams struct {
	// Instance ID
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

type DescribeWhitelistsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

func (r *DescribeWhitelistsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWhitelistsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWhitelistsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWhitelistsResponseParams struct {
	// Allowlist list
	WhitelistSet []*Whitelist `json:"WhitelistSet,omitempty" name:"WhitelistSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWhitelistsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWhitelistsResponseParams `json:"Response"`
}

func (r *DescribeWhitelistsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWhitelistsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// One or more filter values.
	Values []*string `json:"Values,omitempty" name:"Values"`

	// Filter name.
	Name *string `json:"Name,omitempty" name:"Name"`
}

type LogExport struct {
	// Log export path
	CosPath *string `json:"CosPath,omitempty" name:"CosPath"`

	// Number of logs to be exported
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// Log export task creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Log export task ID
	ExportID *string `json:"ExportID,omitempty" name:"ExportID"`

	// Log export filename
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// Log file size
	FileSize *int64 `json:"FileSize,omitempty" name:"FileSize"`

	// Log export format
	Format *string `json:"Format,omitempty" name:"Format"`

	// Log export time sorting
	Order *string `json:"Order,omitempty" name:"Order"`

	// Log export query statement
	Query *string `json:"Query,omitempty" name:"Query"`

	// Log export start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Log export end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Log download status. Valid values: Queuing: queuing; Processing: exporting; Complete: completed; Failed: failed; Expired: expired (3-day validity period).
	Status *string `json:"Status,omitempty" name:"Status"`
}

// Predefined struct for user
type ModifyInstanceRequestParams struct {
	// ID of the instance to be modified
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// New instance name (up to 255 characters)
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// New instance description (up to 1,024 characters)
	InstanceDesc *string `json:"InstanceDesc,omitempty" name:"InstanceDesc"`
}

type ModifyInstanceRequest struct {
	*tchttp.BaseRequest
	
	// ID of the instance to be modified
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// New instance name (up to 255 characters)
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// New instance description (up to 1,024 characters)
	InstanceDesc *string `json:"InstanceDesc,omitempty" name:"InstanceDesc"`
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
	delete(f, "InstanceDesc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

// Predefined struct for user
type ModifyProjectLimitRequestParams struct {
	// Project ID
	ProjectID *int64 `json:"ProjectID,omitempty" name:"ProjectID"`

	// Project API
	ProjectInterface *string `json:"ProjectInterface,omitempty" name:"ProjectInterface"`

	// Reporting rate. 10 means 10%
	ReportRate *int64 `json:"ReportRate,omitempty" name:"ReportRate"`

	// Reporting type. 1: rate; 2: number of reported data entries
	ReportType *int64 `json:"ReportType,omitempty" name:"ReportType"`

	// Primary key ID
	ID *int64 `json:"ID,omitempty" name:"ID"`
}

type ModifyProjectLimitRequest struct {
	*tchttp.BaseRequest
	
	// Project ID
	ProjectID *int64 `json:"ProjectID,omitempty" name:"ProjectID"`

	// Project API
	ProjectInterface *string `json:"ProjectInterface,omitempty" name:"ProjectInterface"`

	// Reporting rate. 10 means 10%
	ReportRate *int64 `json:"ReportRate,omitempty" name:"ReportRate"`

	// Reporting type. 1: rate; 2: number of reported data entries
	ReportType *int64 `json:"ReportType,omitempty" name:"ReportType"`

	// Primary key ID
	ID *int64 `json:"ID,omitempty" name:"ID"`
}

func (r *ModifyProjectLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProjectLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectID")
	delete(f, "ProjectInterface")
	delete(f, "ReportRate")
	delete(f, "ReportType")
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProjectLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProjectLimitResponseParams struct {
	// Returned message
	// Note: this field may return null, indicating that no valid values can be obtained.
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyProjectLimitResponse struct {
	*tchttp.BaseResponse
	Response *ModifyProjectLimitResponseParams `json:"Response"`
}

func (r *ModifyProjectLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProjectLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProjectRequestParams struct {
	// Project ID
	ID *uint64 `json:"ID,omitempty" name:"ID"`

	// Project name (optional, not empty, and up to 200 characters)
	Name *string `json:"Name,omitempty" name:"Name"`

	// Project webpage URL (optional and up to 256 characters)
	URL *string `json:"URL,omitempty" name:"URL"`

	// Project repository address (optional and up to 256 characters)
	Repo *string `json:"Repo,omitempty" name:"Repo"`

	// ID of the instance to which to move the project (optional)
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`

	// Project sample rate (optional)
	Rate *string `json:"Rate,omitempty" name:"Rate"`

	// Whether to enable aggregation (optional)
	EnableURLGroup *uint64 `json:"EnableURLGroup,omitempty" name:"EnableURLGroup"`

	// Project type (valid values: "web", "mp", "android", "ios", "node", "hippy", "weex", "viola", "rn")
	Type *string `json:"Type,omitempty" name:"Type"`

	// Project description (optional and up to 1,000 characters)
	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

type ModifyProjectRequest struct {
	*tchttp.BaseRequest
	
	// Project ID
	ID *uint64 `json:"ID,omitempty" name:"ID"`

	// Project name (optional, not empty, and up to 200 characters)
	Name *string `json:"Name,omitempty" name:"Name"`

	// Project webpage URL (optional and up to 256 characters)
	URL *string `json:"URL,omitempty" name:"URL"`

	// Project repository address (optional and up to 256 characters)
	Repo *string `json:"Repo,omitempty" name:"Repo"`

	// ID of the instance to which to move the project (optional)
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`

	// Project sample rate (optional)
	Rate *string `json:"Rate,omitempty" name:"Rate"`

	// Whether to enable aggregation (optional)
	EnableURLGroup *uint64 `json:"EnableURLGroup,omitempty" name:"EnableURLGroup"`

	// Project type (valid values: "web", "mp", "android", "ios", "node", "hippy", "weex", "viola", "rn")
	Type *string `json:"Type,omitempty" name:"Type"`

	// Project description (optional and up to 1,000 characters)
	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

func (r *ModifyProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	delete(f, "Name")
	delete(f, "URL")
	delete(f, "Repo")
	delete(f, "InstanceID")
	delete(f, "Rate")
	delete(f, "EnableURLGroup")
	delete(f, "Type")
	delete(f, "Desc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProjectResponseParams struct {
	// Operation information
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// Project ID
	ID *uint64 `json:"ID,omitempty" name:"ID"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyProjectResponse struct {
	*tchttp.BaseResponse
	Response *ModifyProjectResponseParams `json:"Response"`
}

func (r *ModifyProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProjectLimit struct {
	// Primary key ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Project ID
	ProjectID *int64 `json:"ProjectID,omitempty" name:"ProjectID"`

	// API
	ProjectInterface *string `json:"ProjectInterface,omitempty" name:"ProjectInterface"`

	// Reporting rate
	ReportRate *int64 `json:"ReportRate,omitempty" name:"ReportRate"`

	// Reporting type. 1: reporting rate; 2: reporting count limit
	ReportType *int64 `json:"ReportType,omitempty" name:"ReportType"`
}

type ReleaseFile struct {
	// File version
	Version *string `json:"Version,omitempty" name:"Version"`

	// Unique file key
	FileKey *string `json:"FileKey,omitempty" name:"FileKey"`

	// Filename
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// File hash
	FileHash *string `json:"FileHash,omitempty" name:"FileHash"`

	// File ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ID *int64 `json:"ID,omitempty" name:"ID"`
}

// Predefined struct for user
type ResumeInstanceRequestParams struct {
	// ID of the instance to be resumed
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type ResumeInstanceRequest struct {
	*tchttp.BaseRequest
	
	// ID of the instance to be resumed
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ResumeInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResumeInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeInstanceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ResumeInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ResumeInstanceResponseParams `json:"Response"`
}

func (r *ResumeInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RumAreaInfo struct {
	// Region ID
	AreaId *int64 `json:"AreaId,omitempty" name:"AreaId"`

	// Region status (1: valid; 2: invalid)
	AreaStatus *int64 `json:"AreaStatus,omitempty" name:"AreaStatus"`

	// Region name
	AreaName *string `json:"AreaName,omitempty" name:"AreaName"`

	// Region key
	AreaKey *string `json:"AreaKey,omitempty" name:"AreaKey"`

	// Region ID.
	AreaRegionID *string `json:"AreaRegionID,omitempty" name:"AreaRegionID"`

	// Region code, such as “ap-xxx” (xxx is the region name).
	AreaRegionCode *string `json:"AreaRegionCode,omitempty" name:"AreaRegionCode"`

	// Region abbreviation.
	AreaAbbr *string `json:"AreaAbbr,omitempty" name:"AreaAbbr"`
}

type RumProject struct {
	// Project name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Creator ID
	Creator *string `json:"Creator,omitempty" name:"Creator"`

	// Instance ID
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`

	// Project type
	Type *string `json:"Type,omitempty" name:"Type"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Project repository address
	// Note: this field may return null, indicating that no valid values can be obtained.
	Repo *string `json:"Repo,omitempty" name:"Repo"`

	// Project URL
	// Note: this field may return null, indicating that no valid values can be obtained.
	URL *string `json:"URL,omitempty" name:"URL"`

	// Project sample rate
	Rate *string `json:"Rate,omitempty" name:"Rate"`

	// Unique project key (12 characters)
	Key *string `json:"Key,omitempty" name:"Key"`

	// Whether to enable URL aggregation
	EnableURLGroup *int64 `json:"EnableURLGroup,omitempty" name:"EnableURLGroup"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Project ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// Instance key
	InstanceKey *string `json:"InstanceKey,omitempty" name:"InstanceKey"`

	// Project description
	// Note: this field may return null, indicating that no valid values can be obtained.
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// Starred status. 1: yes; 0: no
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsStar *int64 `json:"IsStar,omitempty" name:"IsStar"`

	// Project status (`1`: Creating; `2`: Running; `3`: Abnormal; `4`: Restarting; `5`: Stopping; `6`: Stopped; `7`: Terminating; `8`: Terminated)
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ProjectStatus *int64 `json:"ProjectStatus,omitempty" name:"ProjectStatus"`
}

type RumPvInfo struct {
	// Project ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Number of PVs
	// Note: this field may return null, indicating that no valid values can be obtained.
	Pv *string `json:"Pv,omitempty" name:"Pv"`

	// Time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type RumUvInfo struct {
	// Project ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Number of UVs
	Uv *string `json:"Uv,omitempty" name:"Uv"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type ScoreInfo struct {
	// duration
	StaticDuration *string `json:"StaticDuration,omitempty" name:"StaticDuration"`

	// pv
	PagePv *string `json:"PagePv,omitempty" name:"PagePv"`

	// Failure
	ApiFail *string `json:"ApiFail,omitempty" name:"ApiFail"`

	// Request
	ApiNum *string `json:"ApiNum,omitempty" name:"ApiNum"`

	// fail
	StaticFail *string `json:"StaticFail,omitempty" name:"StaticFail"`

	// Project ID
	ProjectID *int64 `json:"ProjectID,omitempty" name:"ProjectID"`

	// uv
	PageUv *string `json:"PageUv,omitempty" name:"PageUv"`

	// Number of requests
	ApiDuration *string `json:"ApiDuration,omitempty" name:"ApiDuration"`

	// Score
	Score *string `json:"Score,omitempty" name:"Score"`

	// error
	PageError *string `json:"PageError,omitempty" name:"PageError"`

	// num
	StaticNum *string `json:"StaticNum,omitempty" name:"StaticNum"`

	// num
	RecordNum *int64 `json:"RecordNum,omitempty" name:"RecordNum"`

	// Duration
	PageDuration *string `json:"PageDuration,omitempty" name:"PageDuration"`

	// Time
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

// Predefined struct for user
type StopInstanceRequestParams struct {
	// ID of the instance to be stopped
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type StopInstanceRequest struct {
	*tchttp.BaseRequest
	
	// ID of the instance to be stopped
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *StopInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopInstanceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StopInstanceResponse struct {
	*tchttp.BaseResponse
	Response *StopInstanceResponseParams `json:"Response"`
}

func (r *StopInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopProjectRequestParams struct {
	// Project ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

type StopProjectRequest struct {
	*tchttp.BaseRequest
	
	// Project ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *StopProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopProjectResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StopProjectResponse struct {
	*tchttp.BaseResponse
	Response *StopProjectResponseParams `json:"Response"`
}

func (r *StopProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {
	// Tag key
	Key *string `json:"Key,omitempty" name:"Key"`

	// Tag value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type Whitelist struct {
	// Remarks
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Instance ID
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`

	// End time
	Ttl *string `json:"Ttl,omitempty" name:"Ttl"`

	// Auto-Increment allowlist ID
	ID *string `json:"ID,omitempty" name:"ID"`

	// Unique business identifier
	WhitelistUin *string `json:"WhitelistUin,omitempty" name:"WhitelistUin"`

	// Creator ID
	CreateUser *string `json:"CreateUser,omitempty" name:"CreateUser"`

	// aid
	Aid *string `json:"Aid,omitempty" name:"Aid"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}