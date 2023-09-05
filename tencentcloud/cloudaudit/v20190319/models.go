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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type AttributeKeyDetail struct {
	// Input box type
	LabelType *string `json:"LabelType,omitnil" name:"LabelType"`

	// Initial display
	Starter *string `json:"Starter,omitnil" name:"Starter"`

	// Display sort order
	Order *int64 `json:"Order,omitnil" name:"Order"`

	// `AttributeKey` value
	Value *string `json:"Value,omitnil" name:"Value"`

	// Tag
	Label *string `json:"Label,omitnil" name:"Label"`
}

type AuditSummary struct {
	// Tracking set status. 1: enabled, 0: disabled
	AuditStatus *int64 `json:"AuditStatus,omitnil" name:"AuditStatus"`

	// COS bucket name
	CosBucketName *string `json:"CosBucketName,omitnil" name:"CosBucketName"`

	// Tracking set name
	AuditName *string `json:"AuditName,omitnil" name:"AuditName"`

	// Log prefix
	LogFilePrefix *string `json:"LogFilePrefix,omitnil" name:"LogFilePrefix"`
}

type CmqRegionInfo struct {
	// Region description
	CmqRegionName *string `json:"CmqRegionName,omitnil" name:"CmqRegionName"`

	// CMQ region
	CmqRegion *string `json:"CmqRegion,omitnil" name:"CmqRegion"`
}

type CosRegionInfo struct {
	// COS region
	CosRegion *string `json:"CosRegion,omitnil" name:"CosRegion"`

	// Region description
	CosRegionName *string `json:"CosRegionName,omitnil" name:"CosRegionName"`
}

// Predefined struct for user
type CreateAuditRequestParams struct {
	// Whether to enable CMQ message notification. 1: Yes; 0: No. Only CMQ queue service is currently supported. If CMQ message notification is enabled, CloudAudit will deliver your log contents to the designated queue in the specified region in real time.
	IsEnableCmqNotify *int64 `json:"IsEnableCmqNotify,omitnil" name:"IsEnableCmqNotify"`

	// Manages the read/write attribute of event. Valid values: 1 (read-only), 2 (write-only), 3 (read/write).
	ReadWriteAttribute *int64 `json:"ReadWriteAttribute,omitnil" name:"ReadWriteAttribute"`

	// Tracking set name, which can contain 3–128 ASCII letters (a–z; A–Z), digits (0–9), and underscores (_).
	AuditName *string `json:"AuditName,omitnil" name:"AuditName"`

	// COS region. Supported regions can be queried through the `ListCosEnableRegion` API.
	CosRegion *string `json:"CosRegion,omitnil" name:"CosRegion"`

	// Whether to create a COS bucket. 1: Yes; 0: No.
	IsCreateNewBucket *int64 `json:"IsCreateNewBucket,omitnil" name:"IsCreateNewBucket"`

	// User-defined COS bucket name, which can only contain 1–40 lowercase letters (a–z), digits (0–9), and dashes (-) and cannot begin or end with "-". If a bucket is not newly created, CloudAudit will not verify whether it actually exists. Enter the name with caution to avoid log delivery failure and consequent data loss.
	CosBucketName *string `json:"CosBucketName,omitnil" name:"CosBucketName"`

	// Globally unique ID of the CMK. This value is required if it is not a newly created KMS element. It can be obtained through `ListKeyAliasByRegion`. CloudAudit will not verify the validity of the `KeyId`. Enter it carefully to avoid data loss.
	KeyId *string `json:"KeyId,omitnil" name:"KeyId"`

	// Queue name, which must begin with a letter and can contain up to 64 letters, digits, and dashes (-). This field is required if the value of `IsEnableCmqNotify` is 1. If a queue is not newly created, CloudAudit will not verify whether it actually exists. Enter the name with caution to avoid log delivery failure and consequent data loss.
	CmqQueueName *string `json:"CmqQueueName,omitnil" name:"CmqQueueName"`

	// KMS region. Currently supported regions can be obtained through `ListKmsEnableRegion`. This must be the same as the COS region.
	KmsRegion *string `json:"KmsRegion,omitnil" name:"KmsRegion"`

	// Whether to enable KMS encryption. 1: Yes, 0: No. If KMS encryption is enabled, the data will be encrypted when delivered to COS.
	IsEnableKmsEncry *int64 `json:"IsEnableKmsEncry,omitnil" name:"IsEnableKmsEncry"`

	// Region where the queue is located. Supported CMQ regions can be queried through the `ListCmqEnableRegion` API. This field is required if the value of `IsEnableCmqNotify` is 1.
	CmqRegion *string `json:"CmqRegion,omitnil" name:"CmqRegion"`

	// Log file prefix, which can only contain 3–40 ASCII letters (a–z; A–Z) and digits (0–9). It can be left empty and is the account ID by default.
	LogFilePrefix *string `json:"LogFilePrefix,omitnil" name:"LogFilePrefix"`

	// Whether to create a queue. 1: Yes; 0: No. This field is required if the value of `IsEnableCmqNotify` is 1.
	IsCreateNewQueue *int64 `json:"IsCreateNewQueue,omitnil" name:"IsCreateNewQueue"`
}

type CreateAuditRequest struct {
	*tchttp.BaseRequest
	
	// Whether to enable CMQ message notification. 1: Yes; 0: No. Only CMQ queue service is currently supported. If CMQ message notification is enabled, CloudAudit will deliver your log contents to the designated queue in the specified region in real time.
	IsEnableCmqNotify *int64 `json:"IsEnableCmqNotify,omitnil" name:"IsEnableCmqNotify"`

	// Manages the read/write attribute of event. Valid values: 1 (read-only), 2 (write-only), 3 (read/write).
	ReadWriteAttribute *int64 `json:"ReadWriteAttribute,omitnil" name:"ReadWriteAttribute"`

	// Tracking set name, which can contain 3–128 ASCII letters (a–z; A–Z), digits (0–9), and underscores (_).
	AuditName *string `json:"AuditName,omitnil" name:"AuditName"`

	// COS region. Supported regions can be queried through the `ListCosEnableRegion` API.
	CosRegion *string `json:"CosRegion,omitnil" name:"CosRegion"`

	// Whether to create a COS bucket. 1: Yes; 0: No.
	IsCreateNewBucket *int64 `json:"IsCreateNewBucket,omitnil" name:"IsCreateNewBucket"`

	// User-defined COS bucket name, which can only contain 1–40 lowercase letters (a–z), digits (0–9), and dashes (-) and cannot begin or end with "-". If a bucket is not newly created, CloudAudit will not verify whether it actually exists. Enter the name with caution to avoid log delivery failure and consequent data loss.
	CosBucketName *string `json:"CosBucketName,omitnil" name:"CosBucketName"`

	// Globally unique ID of the CMK. This value is required if it is not a newly created KMS element. It can be obtained through `ListKeyAliasByRegion`. CloudAudit will not verify the validity of the `KeyId`. Enter it carefully to avoid data loss.
	KeyId *string `json:"KeyId,omitnil" name:"KeyId"`

	// Queue name, which must begin with a letter and can contain up to 64 letters, digits, and dashes (-). This field is required if the value of `IsEnableCmqNotify` is 1. If a queue is not newly created, CloudAudit will not verify whether it actually exists. Enter the name with caution to avoid log delivery failure and consequent data loss.
	CmqQueueName *string `json:"CmqQueueName,omitnil" name:"CmqQueueName"`

	// KMS region. Currently supported regions can be obtained through `ListKmsEnableRegion`. This must be the same as the COS region.
	KmsRegion *string `json:"KmsRegion,omitnil" name:"KmsRegion"`

	// Whether to enable KMS encryption. 1: Yes, 0: No. If KMS encryption is enabled, the data will be encrypted when delivered to COS.
	IsEnableKmsEncry *int64 `json:"IsEnableKmsEncry,omitnil" name:"IsEnableKmsEncry"`

	// Region where the queue is located. Supported CMQ regions can be queried through the `ListCmqEnableRegion` API. This field is required if the value of `IsEnableCmqNotify` is 1.
	CmqRegion *string `json:"CmqRegion,omitnil" name:"CmqRegion"`

	// Log file prefix, which can only contain 3–40 ASCII letters (a–z; A–Z) and digits (0–9). It can be left empty and is the account ID by default.
	LogFilePrefix *string `json:"LogFilePrefix,omitnil" name:"LogFilePrefix"`

	// Whether to create a queue. 1: Yes; 0: No. This field is required if the value of `IsEnableCmqNotify` is 1.
	IsCreateNewQueue *int64 `json:"IsCreateNewQueue,omitnil" name:"IsCreateNewQueue"`
}

func (r *CreateAuditRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAuditRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IsEnableCmqNotify")
	delete(f, "ReadWriteAttribute")
	delete(f, "AuditName")
	delete(f, "CosRegion")
	delete(f, "IsCreateNewBucket")
	delete(f, "CosBucketName")
	delete(f, "KeyId")
	delete(f, "CmqQueueName")
	delete(f, "KmsRegion")
	delete(f, "IsEnableKmsEncry")
	delete(f, "CmqRegion")
	delete(f, "LogFilePrefix")
	delete(f, "IsCreateNewQueue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAuditRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAuditResponseParams struct {
	// Whether creation succeeded.
	IsSuccess *int64 `json:"IsSuccess,omitnil" name:"IsSuccess"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateAuditResponse struct {
	*tchttp.BaseResponse
	Response *CreateAuditResponseParams `json:"Response"`
}

func (r *CreateAuditResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAuditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAuditTrackRequestParams struct {
	// Tracking set name, which can only contain 3-48 letters, digits, hyphens, and underscores.
	Name *string `json:"Name,omitnil" name:"Name"`

	// Tracking set event type (`Read`: Read; `Write`: Write; `*`: All)
	ActionType *string `json:"ActionType,omitnil" name:"ActionType"`

	// The product to which the tracking set event belongs. The value can be a single product such as `cos`, or `*` that indicates all products.
	ResourceType *string `json:"ResourceType,omitnil" name:"ResourceType"`

	// Tracking set status (0: Not enabled; 1: Enabled)
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// The list of API names of tracking set events. When `ResourceType` is `*`, the value of `EventNames` must be `*`. When `ResourceType` is a specified product, the value of `EventNames` can be `*`. When `ResourceType` is `cos` or `cls`, up to 10 APIs are supported.
	EventNames []*string `json:"EventNames,omitnil" name:"EventNames"`

	// Storage type of shipped data. Valid values: `cos`, `cls`.
	Storage *Storage `json:"Storage,omitnil" name:"Storage"`

	// Whether to enable the feature of shipping organization members’ operation logs to the organization admin account or the trusted service admin account (0: Not enabled; 1: Enabled. This feature can only be enabled by the organization admin account or the trusted service admin account)
	TrackForAllMembers *uint64 `json:"TrackForAllMembers,omitnil" name:"TrackForAllMembers"`
}

type CreateAuditTrackRequest struct {
	*tchttp.BaseRequest
	
	// Tracking set name, which can only contain 3-48 letters, digits, hyphens, and underscores.
	Name *string `json:"Name,omitnil" name:"Name"`

	// Tracking set event type (`Read`: Read; `Write`: Write; `*`: All)
	ActionType *string `json:"ActionType,omitnil" name:"ActionType"`

	// The product to which the tracking set event belongs. The value can be a single product such as `cos`, or `*` that indicates all products.
	ResourceType *string `json:"ResourceType,omitnil" name:"ResourceType"`

	// Tracking set status (0: Not enabled; 1: Enabled)
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// The list of API names of tracking set events. When `ResourceType` is `*`, the value of `EventNames` must be `*`. When `ResourceType` is a specified product, the value of `EventNames` can be `*`. When `ResourceType` is `cos` or `cls`, up to 10 APIs are supported.
	EventNames []*string `json:"EventNames,omitnil" name:"EventNames"`

	// Storage type of shipped data. Valid values: `cos`, `cls`.
	Storage *Storage `json:"Storage,omitnil" name:"Storage"`

	// Whether to enable the feature of shipping organization members’ operation logs to the organization admin account or the trusted service admin account (0: Not enabled; 1: Enabled. This feature can only be enabled by the organization admin account or the trusted service admin account)
	TrackForAllMembers *uint64 `json:"TrackForAllMembers,omitnil" name:"TrackForAllMembers"`
}

func (r *CreateAuditTrackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAuditTrackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "ActionType")
	delete(f, "ResourceType")
	delete(f, "Status")
	delete(f, "EventNames")
	delete(f, "Storage")
	delete(f, "TrackForAllMembers")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAuditTrackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAuditTrackResponseParams struct {
	// Tracking set ID
	TrackId *uint64 `json:"TrackId,omitnil" name:"TrackId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateAuditTrackResponse struct {
	*tchttp.BaseResponse
	Response *CreateAuditTrackResponseParams `json:"Response"`
}

func (r *CreateAuditTrackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAuditTrackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAuditRequestParams struct {
	// Tracking set name
	AuditName *string `json:"AuditName,omitnil" name:"AuditName"`
}

type DeleteAuditRequest struct {
	*tchttp.BaseRequest
	
	// Tracking set name
	AuditName *string `json:"AuditName,omitnil" name:"AuditName"`
}

func (r *DeleteAuditRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAuditRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AuditName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAuditRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAuditResponseParams struct {
	// Whether deletion succeeded
	IsSuccess *int64 `json:"IsSuccess,omitnil" name:"IsSuccess"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteAuditResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAuditResponseParams `json:"Response"`
}

func (r *DeleteAuditResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAuditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAuditTrackRequestParams struct {
	// Tracking set ID
	TrackId *uint64 `json:"TrackId,omitnil" name:"TrackId"`
}

type DeleteAuditTrackRequest struct {
	*tchttp.BaseRequest
	
	// Tracking set ID
	TrackId *uint64 `json:"TrackId,omitnil" name:"TrackId"`
}

func (r *DeleteAuditTrackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAuditTrackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TrackId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAuditTrackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAuditTrackResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteAuditTrackResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAuditTrackResponseParams `json:"Response"`
}

func (r *DeleteAuditTrackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAuditTrackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditRequestParams struct {
	// Tracking set name
	AuditName *string `json:"AuditName,omitnil" name:"AuditName"`
}

type DescribeAuditRequest struct {
	*tchttp.BaseRequest
	
	// Tracking set name
	AuditName *string `json:"AuditName,omitnil" name:"AuditName"`
}

func (r *DescribeAuditRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AuditName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuditRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditResponseParams struct {
	// Whether to enable CMQ message notification. 1: Yes; 0: No.
	IsEnableCmqNotify *int64 `json:"IsEnableCmqNotify,omitnil" name:"IsEnableCmqNotify"`

	// Manages the read/write attribute of event. Valid values: 1 (read-only), 2 (write-only), 3 (read/write)
	ReadWriteAttribute *int64 `json:"ReadWriteAttribute,omitnil" name:"ReadWriteAttribute"`

	// Globally unique CMK ID.
	KeyId *string `json:"KeyId,omitnil" name:"KeyId"`

	// Tracking set status. 1: enabled, 0: disabled.
	AuditStatus *int64 `json:"AuditStatus,omitnil" name:"AuditStatus"`

	// Tracking set name.
	AuditName *string `json:"AuditName,omitnil" name:"AuditName"`

	// COS bucket region.
	CosRegion *string `json:"CosRegion,omitnil" name:"CosRegion"`

	// Queue name.
	CmqQueueName *string `json:"CmqQueueName,omitnil" name:"CmqQueueName"`

	// CMK alias.
	KmsAlias *string `json:"KmsAlias,omitnil" name:"KmsAlias"`

	// KMS region.
	KmsRegion *string `json:"KmsRegion,omitnil" name:"KmsRegion"`

	// Whether to enable KMS encryption. 1: Yes, 0: No. If KMS encryption is enabled, the data will be encrypted when it is delivered to COS.
	IsEnableKmsEncry *int64 `json:"IsEnableKmsEncry,omitnil" name:"IsEnableKmsEncry"`

	// COS bucket name.
	CosBucketName *string `json:"CosBucketName,omitnil" name:"CosBucketName"`

	// Queue region.
	CmqRegion *string `json:"CmqRegion,omitnil" name:"CmqRegion"`

	// Log prefix.
	LogFilePrefix *string `json:"LogFilePrefix,omitnil" name:"LogFilePrefix"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAuditResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuditResponseParams `json:"Response"`
}

func (r *DescribeAuditResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditTrackRequestParams struct {
	// Tracking set ID
	TrackId *uint64 `json:"TrackId,omitnil" name:"TrackId"`
}

type DescribeAuditTrackRequest struct {
	*tchttp.BaseRequest
	
	// Tracking set ID
	TrackId *uint64 `json:"TrackId,omitnil" name:"TrackId"`
}

func (r *DescribeAuditTrackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditTrackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TrackId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuditTrackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditTrackResponseParams struct {
	// Tracking set name
	Name *string `json:"Name,omitnil" name:"Name"`

	// Tracking set event type (`Read`: Read; `Write`: Write; `*`: All)
	ActionType *string `json:"ActionType,omitnil" name:"ActionType"`

	// The product to which the tracking set event belongs, such as `cos`, or `*` that indicates all products
	ResourceType *string `json:"ResourceType,omitnil" name:"ResourceType"`

	// Tracking set status (0: Not enabled; 1: Enabled)
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// The list of API names of tracking set events (`*`: All)
	EventNames []*string `json:"EventNames,omitnil" name:"EventNames"`

	// Storage type of shipped data. Valid values: `cos`, `cls`.
	Storage *Storage `json:"Storage,omitnil" name:"Storage"`

	// Creation time of the tracking set
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// Whether to enable the feature of shipping organization members’ operation logs to the organization admin account or the trusted service admin account
	// Note: This field may return null, indicating that no valid values can be obtained.
	TrackForAllMembers *uint64 `json:"TrackForAllMembers,omitnil" name:"TrackForAllMembers"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAuditTrackResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuditTrackResponseParams `json:"Response"`
}

func (r *DescribeAuditTrackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditTrackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditTracksRequestParams struct {
	// Page number
	PageNumber *uint64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// The number of tracking sets per page
	PageSize *uint64 `json:"PageSize,omitnil" name:"PageSize"`
}

type DescribeAuditTracksRequest struct {
	*tchttp.BaseRequest
	
	// Page number
	PageNumber *uint64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// The number of tracking sets per page
	PageSize *uint64 `json:"PageSize,omitnil" name:"PageSize"`
}

func (r *DescribeAuditTracksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditTracksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuditTracksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditTracksResponseParams struct {
	// Tracking set list
	Tracks []*Tracks `json:"Tracks,omitnil" name:"Tracks"`

	// Total number of tracking sets
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAuditTracksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuditTracksResponseParams `json:"Response"`
}

func (r *DescribeAuditTracksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditTracksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEventsRequestParams struct {
	// Start timestamp in seconds (cannot be 90 days after the current time).
	StartTime *uint64 `json:"StartTime,omitnil" name:"StartTime"`

	// End timestamp in seconds (the time range for query is less than 30 days).
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`

	// Credential for viewing more logs.
	NextToken *uint64 `json:"NextToken,omitnil" name:"NextToken"`

	// Max number of returned logs (up to 50).
	MaxResults *uint64 `json:"MaxResults,omitnil" name:"MaxResults"`

	// Search condition. Valid values: `RequestId`, `EventName`, `ActionType` (write/read), `PrincipalId` (sub-account), `ResourceType`, `ResourceName`, `AccessKeyId`, `SensitiveAction`, `ApiErrorCode`, `CamErrorCode`, and `Tags` (Format of AttributeValue: [{"key":"*","value":"*"}])
	LookupAttributes []*LookupAttribute `json:"LookupAttributes,omitnil" name:"LookupAttributes"`

	// Whether to return the IP location. `1`: yes, `0`: no.
	IsReturnLocation *uint64 `json:"IsReturnLocation,omitnil" name:"IsReturnLocation"`
}

type DescribeEventsRequest struct {
	*tchttp.BaseRequest
	
	// Start timestamp in seconds (cannot be 90 days after the current time).
	StartTime *uint64 `json:"StartTime,omitnil" name:"StartTime"`

	// End timestamp in seconds (the time range for query is less than 30 days).
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`

	// Credential for viewing more logs.
	NextToken *uint64 `json:"NextToken,omitnil" name:"NextToken"`

	// Max number of returned logs (up to 50).
	MaxResults *uint64 `json:"MaxResults,omitnil" name:"MaxResults"`

	// Search condition. Valid values: `RequestId`, `EventName`, `ActionType` (write/read), `PrincipalId` (sub-account), `ResourceType`, `ResourceName`, `AccessKeyId`, `SensitiveAction`, `ApiErrorCode`, `CamErrorCode`, and `Tags` (Format of AttributeValue: [{"key":"*","value":"*"}])
	LookupAttributes []*LookupAttribute `json:"LookupAttributes,omitnil" name:"LookupAttributes"`

	// Whether to return the IP location. `1`: yes, `0`: no.
	IsReturnLocation *uint64 `json:"IsReturnLocation,omitnil" name:"IsReturnLocation"`
}

func (r *DescribeEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "NextToken")
	delete(f, "MaxResults")
	delete(f, "LookupAttributes")
	delete(f, "IsReturnLocation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEventsResponseParams struct {
	// Whether the log list has come to an end. `true`: Yes. Pagination is not required.
	ListOver *bool `json:"ListOver,omitnil" name:"ListOver"`

	// Credential for viewing more logs.
	NextToken *uint64 `json:"NextToken,omitnil" name:"NextToken"`

	// Logset.
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	Events []*Event `json:"Events,omitnil" name:"Events"`

	// This parameter has been deprecated. Please use `ListOver` and `NextToken` for pagination, and read data of the next page when the value of `ListOver` is `false`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEventsResponseParams `json:"Response"`
}

func (r *DescribeEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Event struct {
	// Log ID
	EventId *string `json:"EventId,omitnil" name:"EventId"`

	// Username
	Username *string `json:"Username,omitnil" name:"Username"`

	// Event Time
	EventTime *string `json:"EventTime,omitnil" name:"EventTime"`

	// Log details
	CloudAuditEvent *string `json:"CloudAuditEvent,omitnil" name:"CloudAuditEvent"`

	// Description of resource type in Chinese (please use this field as required; if you are using other languages, ignore this field)
	ResourceTypeCn *string `json:"ResourceTypeCn,omitnil" name:"ResourceTypeCn"`

	// Authentication error code
	ErrorCode *int64 `json:"ErrorCode,omitnil" name:"ErrorCode"`

	// Event name
	EventName *string `json:"EventName,omitnil" name:"EventName"`

	// Certificate ID
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	SecretId *string `json:"SecretId,omitnil" name:"SecretId"`

	// Request source
	EventSource *string `json:"EventSource,omitnil" name:"EventSource"`

	// Request ID
	RequestID *string `json:"RequestID,omitnil" name:"RequestID"`

	// Resource region
	ResourceRegion *string `json:"ResourceRegion,omitnil" name:"ResourceRegion"`

	// Root account ID
	AccountID *int64 `json:"AccountID,omitnil" name:"AccountID"`

	// Source IP
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	SourceIPAddress *string `json:"SourceIPAddress,omitnil" name:"SourceIPAddress"`

	// Description of event name in Chinese (please use this field as required; if you are using other languages, ignore this field)
	EventNameCn *string `json:"EventNameCn,omitnil" name:"EventNameCn"`

	// Resource pair
	Resources *Resource `json:"Resources,omitnil" name:"Resources"`

	// Event region
	EventRegion *string `json:"EventRegion,omitnil" name:"EventRegion"`

	// IP location
	Location *string `json:"Location,omitnil" name:"Location"`
}

// Predefined struct for user
type GetAttributeKeyRequestParams struct {
	// Website type. Valid values: zh, en. If this parameter is left empty, `zh` will be used by default
	WebsiteType *string `json:"WebsiteType,omitnil" name:"WebsiteType"`
}

type GetAttributeKeyRequest struct {
	*tchttp.BaseRequest
	
	// Website type. Valid values: zh, en. If this parameter is left empty, `zh` will be used by default
	WebsiteType *string `json:"WebsiteType,omitnil" name:"WebsiteType"`
}

func (r *GetAttributeKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAttributeKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WebsiteType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAttributeKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAttributeKeyResponseParams struct {
	// Valid values of `AttributeKey`
	AttributeKeyDetails []*AttributeKeyDetail `json:"AttributeKeyDetails,omitnil" name:"AttributeKeyDetails"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetAttributeKeyResponse struct {
	*tchttp.BaseResponse
	Response *GetAttributeKeyResponseParams `json:"Response"`
}

func (r *GetAttributeKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAttributeKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquireAuditCreditRequestParams struct {

}

type InquireAuditCreditRequest struct {
	*tchttp.BaseRequest
	
}

func (r *InquireAuditCreditRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquireAuditCreditRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquireAuditCreditRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquireAuditCreditResponseParams struct {
	// Number of tracking sets that can be created
	AuditAmount *int64 `json:"AuditAmount,omitnil" name:"AuditAmount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type InquireAuditCreditResponse struct {
	*tchttp.BaseResponse
	Response *InquireAuditCreditResponseParams `json:"Response"`
}

func (r *InquireAuditCreditResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquireAuditCreditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAuditsRequestParams struct {

}

type ListAuditsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ListAuditsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAuditsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAuditsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAuditsResponseParams struct {
	// Set of queried tracking set summaries
	// Note: This field may return null, indicating that no valid values can be obtained.
	AuditSummarys []*AuditSummary `json:"AuditSummarys,omitnil" name:"AuditSummarys"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ListAuditsResponse struct {
	*tchttp.BaseResponse
	Response *ListAuditsResponseParams `json:"Response"`
}

func (r *ListAuditsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAuditsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListCmqEnableRegionRequestParams struct {
	// Website type. zh: Chinese mainland (default); en: outside Chinese mainland.
	WebsiteType *string `json:"WebsiteType,omitnil" name:"WebsiteType"`
}

type ListCmqEnableRegionRequest struct {
	*tchttp.BaseRequest
	
	// Website type. zh: Chinese mainland (default); en: outside Chinese mainland.
	WebsiteType *string `json:"WebsiteType,omitnil" name:"WebsiteType"`
}

func (r *ListCmqEnableRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListCmqEnableRegionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WebsiteType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListCmqEnableRegionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListCmqEnableRegionResponseParams struct {
	// CloudAudit-enabled CMQ AZs
	EnableRegions []*CmqRegionInfo `json:"EnableRegions,omitnil" name:"EnableRegions"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ListCmqEnableRegionResponse struct {
	*tchttp.BaseResponse
	Response *ListCmqEnableRegionResponseParams `json:"Response"`
}

func (r *ListCmqEnableRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListCmqEnableRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListCosEnableRegionRequestParams struct {
	// Website type. zh: Chinese mainland (default); en: outside Chinese mainland.
	WebsiteType *string `json:"WebsiteType,omitnil" name:"WebsiteType"`
}

type ListCosEnableRegionRequest struct {
	*tchttp.BaseRequest
	
	// Website type. zh: Chinese mainland (default); en: outside Chinese mainland.
	WebsiteType *string `json:"WebsiteType,omitnil" name:"WebsiteType"`
}

func (r *ListCosEnableRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListCosEnableRegionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WebsiteType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListCosEnableRegionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListCosEnableRegionResponseParams struct {
	// CloudAudit-enabled COS AZs
	EnableRegions []*CosRegionInfo `json:"EnableRegions,omitnil" name:"EnableRegions"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ListCosEnableRegionResponse struct {
	*tchttp.BaseResponse
	Response *ListCosEnableRegionResponseParams `json:"Response"`
}

func (r *ListCosEnableRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListCosEnableRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type LookUpEventsRequestParams struct {
	// Start time
	StartTime *int64 `json:"StartTime,omitnil" name:"StartTime"`

	// End time
	EndTime *int64 `json:"EndTime,omitnil" name:"EndTime"`

	// Search criteria
	LookupAttributes []*LookupAttribute `json:"LookupAttributes,omitnil" name:"LookupAttributes"`

	// Credential for viewing more logs
	NextToken *string `json:"NextToken,omitnil" name:"NextToken"`

	// Maximum number of logs to be returned
	MaxResults *int64 `json:"MaxResults,omitnil" name:"MaxResults"`

	// CloudAudit mode. Valid values: standard, quick. Default value: standard
	Mode *string `json:"Mode,omitnil" name:"Mode"`
}

type LookUpEventsRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	StartTime *int64 `json:"StartTime,omitnil" name:"StartTime"`

	// End time
	EndTime *int64 `json:"EndTime,omitnil" name:"EndTime"`

	// Search criteria
	LookupAttributes []*LookupAttribute `json:"LookupAttributes,omitnil" name:"LookupAttributes"`

	// Credential for viewing more logs
	NextToken *string `json:"NextToken,omitnil" name:"NextToken"`

	// Maximum number of logs to be returned
	MaxResults *int64 `json:"MaxResults,omitnil" name:"MaxResults"`

	// CloudAudit mode. Valid values: standard, quick. Default value: standard
	Mode *string `json:"Mode,omitnil" name:"Mode"`
}

func (r *LookUpEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LookUpEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "LookupAttributes")
	delete(f, "NextToken")
	delete(f, "MaxResults")
	delete(f, "Mode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "LookUpEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type LookUpEventsResponseParams struct {
	// Credential for viewing more logs
	// Note: This field may return null, indicating that no valid values can be obtained.
	NextToken *string `json:"NextToken,omitnil" name:"NextToken"`

	// Logset
	// Note: This field may return null, indicating that no valid values can be obtained.
	Events []*Event `json:"Events,omitnil" name:"Events"`

	// Whether the logset ends
	// Note: This field may return null, indicating that no valid values can be obtained.
	ListOver *bool `json:"ListOver,omitnil" name:"ListOver"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type LookUpEventsResponse struct {
	*tchttp.BaseResponse
	Response *LookUpEventsResponseParams `json:"Response"`
}

func (r *LookUpEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LookUpEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LookupAttribute struct {
	// Valid values: RequestId, EventName, ReadOnly, Username, ResourceType, ResourceName, AccessKeyId, and EventId
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	AttributeKey *string `json:"AttributeKey,omitnil" name:"AttributeKey"`

	// Value of `AttributeValue`
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	AttributeValue *string `json:"AttributeValue,omitnil" name:"AttributeValue"`
}

// Predefined struct for user
type ModifyAuditTrackRequestParams struct {
	// Tracking set ID
	TrackId *uint64 `json:"TrackId,omitnil" name:"TrackId"`

	// Tracking set name, which can only contain 3-48 letters, digits, hyphens, and underscores.
	Name *string `json:"Name,omitnil" name:"Name"`

	// Tracking set event type (`Read`: Read; `Write`: Write; `*`: All)
	ActionType *string `json:"ActionType,omitnil" name:"ActionType"`

	// The product to which the tracking set event belongs. The value can be a single product such as `cos`, or `*` that indicates all products.
	ResourceType *string `json:"ResourceType,omitnil" name:"ResourceType"`

	// Tracking set status (0: Not enabled; 1: Enabled)
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// The list of API names of tracking set events. When `ResourceType` is `*`, the value of `EventNames` must be `*`. When `ResourceType` is a specified product, the value of `EventNames` can be `*`. When `ResourceType` is `cos` or `cls`, up to 10 APIs are supported.
	EventNames []*string `json:"EventNames,omitnil" name:"EventNames"`

	// Storage type of shipped data. Valid values: `cos`, `cls`.
	Storage *Storage `json:"Storage,omitnil" name:"Storage"`

	// Whether to enable the feature of shipping organization members’ operation logs to the organization admin account or the trusted service admin account (0: Not enabled; 1: Enabled. This feature can only be enabled by the organization admin account or the trusted service admin account)
	TrackForAllMembers *uint64 `json:"TrackForAllMembers,omitnil" name:"TrackForAllMembers"`
}

type ModifyAuditTrackRequest struct {
	*tchttp.BaseRequest
	
	// Tracking set ID
	TrackId *uint64 `json:"TrackId,omitnil" name:"TrackId"`

	// Tracking set name, which can only contain 3-48 letters, digits, hyphens, and underscores.
	Name *string `json:"Name,omitnil" name:"Name"`

	// Tracking set event type (`Read`: Read; `Write`: Write; `*`: All)
	ActionType *string `json:"ActionType,omitnil" name:"ActionType"`

	// The product to which the tracking set event belongs. The value can be a single product such as `cos`, or `*` that indicates all products.
	ResourceType *string `json:"ResourceType,omitnil" name:"ResourceType"`

	// Tracking set status (0: Not enabled; 1: Enabled)
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// The list of API names of tracking set events. When `ResourceType` is `*`, the value of `EventNames` must be `*`. When `ResourceType` is a specified product, the value of `EventNames` can be `*`. When `ResourceType` is `cos` or `cls`, up to 10 APIs are supported.
	EventNames []*string `json:"EventNames,omitnil" name:"EventNames"`

	// Storage type of shipped data. Valid values: `cos`, `cls`.
	Storage *Storage `json:"Storage,omitnil" name:"Storage"`

	// Whether to enable the feature of shipping organization members’ operation logs to the organization admin account or the trusted service admin account (0: Not enabled; 1: Enabled. This feature can only be enabled by the organization admin account or the trusted service admin account)
	TrackForAllMembers *uint64 `json:"TrackForAllMembers,omitnil" name:"TrackForAllMembers"`
}

func (r *ModifyAuditTrackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAuditTrackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TrackId")
	delete(f, "Name")
	delete(f, "ActionType")
	delete(f, "ResourceType")
	delete(f, "Status")
	delete(f, "EventNames")
	delete(f, "Storage")
	delete(f, "TrackForAllMembers")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAuditTrackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAuditTrackResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyAuditTrackResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAuditTrackResponseParams `json:"Response"`
}

func (r *ModifyAuditTrackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAuditTrackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Resource struct {
	// Resource type
	ResourceType *string `json:"ResourceType,omitnil" name:"ResourceType"`

	// Resource name
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	ResourceName *string `json:"ResourceName,omitnil" name:"ResourceName"`
}

// Predefined struct for user
type StartLoggingRequestParams struct {
	// Tracking set name
	AuditName *string `json:"AuditName,omitnil" name:"AuditName"`
}

type StartLoggingRequest struct {
	*tchttp.BaseRequest
	
	// Tracking set name
	AuditName *string `json:"AuditName,omitnil" name:"AuditName"`
}

func (r *StartLoggingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartLoggingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AuditName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartLoggingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartLoggingResponseParams struct {
	// Whether enablement succeeded
	IsSuccess *int64 `json:"IsSuccess,omitnil" name:"IsSuccess"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type StartLoggingResponse struct {
	*tchttp.BaseResponse
	Response *StartLoggingResponseParams `json:"Response"`
}

func (r *StartLoggingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartLoggingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopLoggingRequestParams struct {
	// Tracking set name
	AuditName *string `json:"AuditName,omitnil" name:"AuditName"`
}

type StopLoggingRequest struct {
	*tchttp.BaseRequest
	
	// Tracking set name
	AuditName *string `json:"AuditName,omitnil" name:"AuditName"`
}

func (r *StopLoggingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopLoggingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AuditName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopLoggingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopLoggingResponseParams struct {
	// Whether disablement succeeded
	IsSuccess *int64 `json:"IsSuccess,omitnil" name:"IsSuccess"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type StopLoggingResponse struct {
	*tchttp.BaseResponse
	Response *StopLoggingResponseParams `json:"Response"`
}

func (r *StopLoggingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopLoggingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Storage struct {
	// Storage type (Valid values: cos, cls)
	StorageType *string `json:"StorageType,omitnil" name:"StorageType"`

	// Storage region
	StorageRegion *string `json:"StorageRegion,omitnil" name:"StorageRegion"`

	// Storage name. For COS, the storage name is the custom bucket name, which can contain up to 50 lowercase letters, digits, and hyphens. It cannot contain "-APPID" and cannot start or end with a hyphen. For CLS, the storage name is the log topic ID, which can contain 1-50 characters.
	StorageName *string `json:"StorageName,omitnil" name:"StorageName"`

	// Storage directory prefix. The COS log file prefix can only contain 3-40 letters and digits.
	StoragePrefix *string `json:"StoragePrefix,omitnil" name:"StoragePrefix"`
}

type Tracks struct {
	// Tracking set name
	Name *string `json:"Name,omitnil" name:"Name"`

	// Tracking set event type (`Read`: Read; `Write`: Write; `*`: All)
	ActionType *string `json:"ActionType,omitnil" name:"ActionType"`

	// The product to which the tracking set event belongs, such as `cos`, or `*` that indicates all products
	ResourceType *string `json:"ResourceType,omitnil" name:"ResourceType"`

	// Tracking set status (0: Not enabled; 1: Enabled)
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// The list of API names of tracking set events (`*`: All)
	EventNames []*string `json:"EventNames,omitnil" name:"EventNames"`

	// Storage type of shipped data. Valid values: `cos`, `cls`.
	Storage *Storage `json:"Storage,omitnil" name:"Storage"`

	// Creation time of the tracking set
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// Tracking set ID
	TrackId *uint64 `json:"TrackId,omitnil" name:"TrackId"`
}

// Predefined struct for user
type UpdateAuditRequestParams struct {
	// Tracking set name
	AuditName *string `json:"AuditName,omitnil" name:"AuditName"`

	// Whether to enable CMQ message notification. 1: Yes; 0: No. Only CMQ queue service is currently supported. If CMQ message notification is enabled, CloudAudit will deliver your log contents to the designated queue in the specified region in real time.
	IsEnableCmqNotify *int64 `json:"IsEnableCmqNotify,omitnil" name:"IsEnableCmqNotify"`

	// Manages the read/write attribute of event. Valid values: 1 (read-only), 2 (write-only), 3 (read/write).
	ReadWriteAttribute *int64 `json:"ReadWriteAttribute,omitnil" name:"ReadWriteAttribute"`

	// Globally unique ID of the CMK. This value is required if it is not a newly created KMS element. It can be obtained through `ListKeyAliasByRegion`. CloudAudit will not verify the validity of the `KeyId`. Enter it carefully to avoid data loss.
	KeyId *string `json:"KeyId,omitnil" name:"KeyId"`

	// COS region. Supported regions can be queried through the `ListCosEnableRegion` API.
	CosRegion *string `json:"CosRegion,omitnil" name:"CosRegion"`

	// Queue name, which must begin with a letter and can contain up to 64 letters, digits, and dashes (-). This field is required if the value of `IsEnableCmqNotify` is 1. If a queue is not newly created, CloudAudit will not verify whether it actually exists. Enter the name with caution to avoid log delivery failure and consequent data loss.
	CmqQueueName *string `json:"CmqQueueName,omitnil" name:"CmqQueueName"`

	// Whether to create a COS bucket. 1: Yes; 0: No.
	IsCreateNewBucket *int64 `json:"IsCreateNewBucket,omitnil" name:"IsCreateNewBucket"`

	// KMS region. Currently supported regions can be obtained through `ListKmsEnableRegion`. This must be the same as the COS region.
	KmsRegion *string `json:"KmsRegion,omitnil" name:"KmsRegion"`

	// Whether to enable KMS encryption. 1: Yes, 0: No. If KMS encryption is enabled, the data will be encrypted when delivered to COS.
	IsEnableKmsEncry *int64 `json:"IsEnableKmsEncry,omitnil" name:"IsEnableKmsEncry"`

	// User-defined COS bucket name, which can only contain 1–40 lowercase letters (a–z), digits (0–9), and dashes (-) and cannot begin or end with "-". If a bucket is not newly created, CloudAudit will not verify whether it actually exists. Enter the name with caution to avoid log delivery failure and consequent data loss.
	CosBucketName *string `json:"CosBucketName,omitnil" name:"CosBucketName"`

	// Region where the queue is located. Supported CMQ regions can be queried through the `ListCmqEnableRegion` API. This field is required if the value of `IsEnableCmqNotify` is 1.
	CmqRegion *string `json:"CmqRegion,omitnil" name:"CmqRegion"`

	// Log file prefix, which can only contain 3–40 ASCII letters (a–z; A–Z) and digits (0–9).
	LogFilePrefix *string `json:"LogFilePrefix,omitnil" name:"LogFilePrefix"`

	// Whether to create a queue. 1: Yes; 0: No. This field is required if the value of `IsEnableCmqNotify` is 1.
	IsCreateNewQueue *int64 `json:"IsCreateNewQueue,omitnil" name:"IsCreateNewQueue"`
}

type UpdateAuditRequest struct {
	*tchttp.BaseRequest
	
	// Tracking set name
	AuditName *string `json:"AuditName,omitnil" name:"AuditName"`

	// Whether to enable CMQ message notification. 1: Yes; 0: No. Only CMQ queue service is currently supported. If CMQ message notification is enabled, CloudAudit will deliver your log contents to the designated queue in the specified region in real time.
	IsEnableCmqNotify *int64 `json:"IsEnableCmqNotify,omitnil" name:"IsEnableCmqNotify"`

	// Manages the read/write attribute of event. Valid values: 1 (read-only), 2 (write-only), 3 (read/write).
	ReadWriteAttribute *int64 `json:"ReadWriteAttribute,omitnil" name:"ReadWriteAttribute"`

	// Globally unique ID of the CMK. This value is required if it is not a newly created KMS element. It can be obtained through `ListKeyAliasByRegion`. CloudAudit will not verify the validity of the `KeyId`. Enter it carefully to avoid data loss.
	KeyId *string `json:"KeyId,omitnil" name:"KeyId"`

	// COS region. Supported regions can be queried through the `ListCosEnableRegion` API.
	CosRegion *string `json:"CosRegion,omitnil" name:"CosRegion"`

	// Queue name, which must begin with a letter and can contain up to 64 letters, digits, and dashes (-). This field is required if the value of `IsEnableCmqNotify` is 1. If a queue is not newly created, CloudAudit will not verify whether it actually exists. Enter the name with caution to avoid log delivery failure and consequent data loss.
	CmqQueueName *string `json:"CmqQueueName,omitnil" name:"CmqQueueName"`

	// Whether to create a COS bucket. 1: Yes; 0: No.
	IsCreateNewBucket *int64 `json:"IsCreateNewBucket,omitnil" name:"IsCreateNewBucket"`

	// KMS region. Currently supported regions can be obtained through `ListKmsEnableRegion`. This must be the same as the COS region.
	KmsRegion *string `json:"KmsRegion,omitnil" name:"KmsRegion"`

	// Whether to enable KMS encryption. 1: Yes, 0: No. If KMS encryption is enabled, the data will be encrypted when delivered to COS.
	IsEnableKmsEncry *int64 `json:"IsEnableKmsEncry,omitnil" name:"IsEnableKmsEncry"`

	// User-defined COS bucket name, which can only contain 1–40 lowercase letters (a–z), digits (0–9), and dashes (-) and cannot begin or end with "-". If a bucket is not newly created, CloudAudit will not verify whether it actually exists. Enter the name with caution to avoid log delivery failure and consequent data loss.
	CosBucketName *string `json:"CosBucketName,omitnil" name:"CosBucketName"`

	// Region where the queue is located. Supported CMQ regions can be queried through the `ListCmqEnableRegion` API. This field is required if the value of `IsEnableCmqNotify` is 1.
	CmqRegion *string `json:"CmqRegion,omitnil" name:"CmqRegion"`

	// Log file prefix, which can only contain 3–40 ASCII letters (a–z; A–Z) and digits (0–9).
	LogFilePrefix *string `json:"LogFilePrefix,omitnil" name:"LogFilePrefix"`

	// Whether to create a queue. 1: Yes; 0: No. This field is required if the value of `IsEnableCmqNotify` is 1.
	IsCreateNewQueue *int64 `json:"IsCreateNewQueue,omitnil" name:"IsCreateNewQueue"`
}

func (r *UpdateAuditRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAuditRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AuditName")
	delete(f, "IsEnableCmqNotify")
	delete(f, "ReadWriteAttribute")
	delete(f, "KeyId")
	delete(f, "CosRegion")
	delete(f, "CmqQueueName")
	delete(f, "IsCreateNewBucket")
	delete(f, "KmsRegion")
	delete(f, "IsEnableKmsEncry")
	delete(f, "CosBucketName")
	delete(f, "CmqRegion")
	delete(f, "LogFilePrefix")
	delete(f, "IsCreateNewQueue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAuditRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAuditResponseParams struct {
	// Whether update succeeded
	IsSuccess *int64 `json:"IsSuccess,omitnil" name:"IsSuccess"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpdateAuditResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAuditResponseParams `json:"Response"`
}

func (r *UpdateAuditResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAuditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}