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
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AttributeKeyDetail struct {
	// Input box type
	LabelType *string `json:"LabelType,omitempty" name:"LabelType"`

	// Initial display
	Starter *string `json:"Starter,omitempty" name:"Starter"`

	// Display sort order
	Order *int64 `json:"Order,omitempty" name:"Order"`

	// `AttributeKey` value
	Value *string `json:"Value,omitempty" name:"Value"`

	// Tag
	Label *string `json:"Label,omitempty" name:"Label"`
}

type AuditSummary struct {
	// Tracking set status. 1: enabled, 0: disabled
	AuditStatus *int64 `json:"AuditStatus,omitempty" name:"AuditStatus"`

	// COS bucket name
	CosBucketName *string `json:"CosBucketName,omitempty" name:"CosBucketName"`

	// Tracking set name
	AuditName *string `json:"AuditName,omitempty" name:"AuditName"`

	// Log prefix
	LogFilePrefix *string `json:"LogFilePrefix,omitempty" name:"LogFilePrefix"`
}

type CmqRegionInfo struct {
	// Region description
	CmqRegionName *string `json:"CmqRegionName,omitempty" name:"CmqRegionName"`

	// CMQ region
	CmqRegion *string `json:"CmqRegion,omitempty" name:"CmqRegion"`
}

type CosRegionInfo struct {
	// COS region
	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`

	// Region description
	CosRegionName *string `json:"CosRegionName,omitempty" name:"CosRegionName"`
}

// Predefined struct for user
type CreateAuditRequestParams struct {
	// Whether to enable CMQ message notification. 1: Yes; 0: No. Only CMQ queue service is currently supported. If CMQ message notification is enabled, CloudAudit will deliver your log contents to the designated queue in the specified region in real time.
	IsEnableCmqNotify *int64 `json:"IsEnableCmqNotify,omitempty" name:"IsEnableCmqNotify"`

	// Manages the read/write attribute of event. Valid values: 1 (read-only), 2 (write-only), 3 (read/write).
	ReadWriteAttribute *int64 `json:"ReadWriteAttribute,omitempty" name:"ReadWriteAttribute"`

	// Tracking set name, which can contain 3–128 ASCII letters (a–z; A–Z), digits (0–9), and underscores (_).
	AuditName *string `json:"AuditName,omitempty" name:"AuditName"`

	// COS region. Supported regions can be queried through the `ListCosEnableRegion` API.
	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`

	// Whether to create a COS bucket. 1: Yes; 0: No.
	IsCreateNewBucket *int64 `json:"IsCreateNewBucket,omitempty" name:"IsCreateNewBucket"`

	// User-defined COS bucket name, which can only contain 1–40 lowercase letters (a–z), digits (0–9), and dashes (-) and cannot begin or end with "-". If a bucket is not newly created, CloudAudit will not verify whether it actually exists. Enter the name with caution to avoid log delivery failure and consequent data loss.
	CosBucketName *string `json:"CosBucketName,omitempty" name:"CosBucketName"`

	// Globally unique ID of the CMK. This value is required if it is not a newly created KMS element. It can be obtained through `ListKeyAliasByRegion`. CloudAudit will not verify the validity of the `KeyId`. Enter it carefully to avoid data loss.
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Queue name, which must begin with a letter and can contain up to 64 letters, digits, and dashes (-). This field is required if the value of `IsEnableCmqNotify` is 1. If a queue is not newly created, CloudAudit will not verify whether it actually exists. Enter the name with caution to avoid log delivery failure and consequent data loss.
	CmqQueueName *string `json:"CmqQueueName,omitempty" name:"CmqQueueName"`

	// KMS region. Currently supported regions can be obtained through `ListKmsEnableRegion`. This must be the same as the COS region.
	KmsRegion *string `json:"KmsRegion,omitempty" name:"KmsRegion"`

	// Whether to enable KMS encryption. 1: Yes, 0: No. If KMS encryption is enabled, the data will be encrypted when delivered to COS.
	IsEnableKmsEncry *int64 `json:"IsEnableKmsEncry,omitempty" name:"IsEnableKmsEncry"`

	// Region where the queue is located. Supported CMQ regions can be queried through the `ListCmqEnableRegion` API. This field is required if the value of `IsEnableCmqNotify` is 1.
	CmqRegion *string `json:"CmqRegion,omitempty" name:"CmqRegion"`

	// Log file prefix, which can only contain 3–40 ASCII letters (a–z; A–Z) and digits (0–9). It can be left empty and is the account ID by default.
	LogFilePrefix *string `json:"LogFilePrefix,omitempty" name:"LogFilePrefix"`

	// Whether to create a queue. 1: Yes; 0: No. This field is required if the value of `IsEnableCmqNotify` is 1.
	IsCreateNewQueue *int64 `json:"IsCreateNewQueue,omitempty" name:"IsCreateNewQueue"`
}

type CreateAuditRequest struct {
	*tchttp.BaseRequest
	
	// Whether to enable CMQ message notification. 1: Yes; 0: No. Only CMQ queue service is currently supported. If CMQ message notification is enabled, CloudAudit will deliver your log contents to the designated queue in the specified region in real time.
	IsEnableCmqNotify *int64 `json:"IsEnableCmqNotify,omitempty" name:"IsEnableCmqNotify"`

	// Manages the read/write attribute of event. Valid values: 1 (read-only), 2 (write-only), 3 (read/write).
	ReadWriteAttribute *int64 `json:"ReadWriteAttribute,omitempty" name:"ReadWriteAttribute"`

	// Tracking set name, which can contain 3–128 ASCII letters (a–z; A–Z), digits (0–9), and underscores (_).
	AuditName *string `json:"AuditName,omitempty" name:"AuditName"`

	// COS region. Supported regions can be queried through the `ListCosEnableRegion` API.
	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`

	// Whether to create a COS bucket. 1: Yes; 0: No.
	IsCreateNewBucket *int64 `json:"IsCreateNewBucket,omitempty" name:"IsCreateNewBucket"`

	// User-defined COS bucket name, which can only contain 1–40 lowercase letters (a–z), digits (0–9), and dashes (-) and cannot begin or end with "-". If a bucket is not newly created, CloudAudit will not verify whether it actually exists. Enter the name with caution to avoid log delivery failure and consequent data loss.
	CosBucketName *string `json:"CosBucketName,omitempty" name:"CosBucketName"`

	// Globally unique ID of the CMK. This value is required if it is not a newly created KMS element. It can be obtained through `ListKeyAliasByRegion`. CloudAudit will not verify the validity of the `KeyId`. Enter it carefully to avoid data loss.
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Queue name, which must begin with a letter and can contain up to 64 letters, digits, and dashes (-). This field is required if the value of `IsEnableCmqNotify` is 1. If a queue is not newly created, CloudAudit will not verify whether it actually exists. Enter the name with caution to avoid log delivery failure and consequent data loss.
	CmqQueueName *string `json:"CmqQueueName,omitempty" name:"CmqQueueName"`

	// KMS region. Currently supported regions can be obtained through `ListKmsEnableRegion`. This must be the same as the COS region.
	KmsRegion *string `json:"KmsRegion,omitempty" name:"KmsRegion"`

	// Whether to enable KMS encryption. 1: Yes, 0: No. If KMS encryption is enabled, the data will be encrypted when delivered to COS.
	IsEnableKmsEncry *int64 `json:"IsEnableKmsEncry,omitempty" name:"IsEnableKmsEncry"`

	// Region where the queue is located. Supported CMQ regions can be queried through the `ListCmqEnableRegion` API. This field is required if the value of `IsEnableCmqNotify` is 1.
	CmqRegion *string `json:"CmqRegion,omitempty" name:"CmqRegion"`

	// Log file prefix, which can only contain 3–40 ASCII letters (a–z; A–Z) and digits (0–9). It can be left empty and is the account ID by default.
	LogFilePrefix *string `json:"LogFilePrefix,omitempty" name:"LogFilePrefix"`

	// Whether to create a queue. 1: Yes; 0: No. This field is required if the value of `IsEnableCmqNotify` is 1.
	IsCreateNewQueue *int64 `json:"IsCreateNewQueue,omitempty" name:"IsCreateNewQueue"`
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
	IsSuccess *int64 `json:"IsSuccess,omitempty" name:"IsSuccess"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

}

type CreateAuditTrackRequest struct {
	*tchttp.BaseRequest
	
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
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAuditTrackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAuditTrackResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	AuditName *string `json:"AuditName,omitempty" name:"AuditName"`
}

type DeleteAuditRequest struct {
	*tchttp.BaseRequest
	
	// Tracking set name
	AuditName *string `json:"AuditName,omitempty" name:"AuditName"`
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
	IsSuccess *int64 `json:"IsSuccess,omitempty" name:"IsSuccess"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

}

type DeleteAuditTrackRequest struct {
	*tchttp.BaseRequest
	
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
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAuditTrackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAuditTrackResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	AuditName *string `json:"AuditName,omitempty" name:"AuditName"`
}

type DescribeAuditRequest struct {
	*tchttp.BaseRequest
	
	// Tracking set name
	AuditName *string `json:"AuditName,omitempty" name:"AuditName"`
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
	IsEnableCmqNotify *int64 `json:"IsEnableCmqNotify,omitempty" name:"IsEnableCmqNotify"`

	// Manages the read/write attribute of event. Valid values: 1 (read-only), 2 (write-only), 3 (read/write)
	ReadWriteAttribute *int64 `json:"ReadWriteAttribute,omitempty" name:"ReadWriteAttribute"`

	// Globally unique CMK ID.
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Tracking set status. 1: enabled, 0: disabled.
	AuditStatus *int64 `json:"AuditStatus,omitempty" name:"AuditStatus"`

	// Tracking set name.
	AuditName *string `json:"AuditName,omitempty" name:"AuditName"`

	// COS bucket region.
	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`

	// Queue name.
	CmqQueueName *string `json:"CmqQueueName,omitempty" name:"CmqQueueName"`

	// CMK alias.
	KmsAlias *string `json:"KmsAlias,omitempty" name:"KmsAlias"`

	// KMS region.
	KmsRegion *string `json:"KmsRegion,omitempty" name:"KmsRegion"`

	// Whether to enable KMS encryption. 1: Yes, 0: No. If KMS encryption is enabled, the data will be encrypted when it is delivered to COS.
	IsEnableKmsEncry *int64 `json:"IsEnableKmsEncry,omitempty" name:"IsEnableKmsEncry"`

	// COS bucket name.
	CosBucketName *string `json:"CosBucketName,omitempty" name:"CosBucketName"`

	// Queue region.
	CmqRegion *string `json:"CmqRegion,omitempty" name:"CmqRegion"`

	// Log prefix.
	LogFilePrefix *string `json:"LogFilePrefix,omitempty" name:"LogFilePrefix"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeAuditTracksRequestParams struct {

}

type DescribeAuditTracksRequest struct {
	*tchttp.BaseRequest
	
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
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuditTracksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditTracksResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// End timestamp in seconds (the time range for query is less than 30 days).
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// Credential for viewing more logs.
	NextToken *uint64 `json:"NextToken,omitempty" name:"NextToken"`

	// Max number of returned logs (up to 50).
	MaxResults *uint64 `json:"MaxResults,omitempty" name:"MaxResults"`

	// Search criterion. Valid values: RequestId, EventName, ActionType (write/read), PrincipalId (sub-account), ResourceType, ResourceName, AccessKeyId, SensitiveAction, ApiErrorCode, and CamErrorCode.
	LookupAttributes []*LookupAttribute `json:"LookupAttributes,omitempty" name:"LookupAttributes"`

	// Whether to return the IP location. `1`: yes, `0`: no.
	IsReturnLocation *uint64 `json:"IsReturnLocation,omitempty" name:"IsReturnLocation"`
}

type DescribeEventsRequest struct {
	*tchttp.BaseRequest
	
	// Start timestamp in seconds (cannot be 90 days after the current time).
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// End timestamp in seconds (the time range for query is less than 30 days).
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// Credential for viewing more logs.
	NextToken *uint64 `json:"NextToken,omitempty" name:"NextToken"`

	// Max number of returned logs (up to 50).
	MaxResults *uint64 `json:"MaxResults,omitempty" name:"MaxResults"`

	// Search criterion. Valid values: RequestId, EventName, ActionType (write/read), PrincipalId (sub-account), ResourceType, ResourceName, AccessKeyId, SensitiveAction, ApiErrorCode, and CamErrorCode.
	LookupAttributes []*LookupAttribute `json:"LookupAttributes,omitempty" name:"LookupAttributes"`

	// Whether to return the IP location. `1`: yes, `0`: no.
	IsReturnLocation *uint64 `json:"IsReturnLocation,omitempty" name:"IsReturnLocation"`
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
	// Whether the logset ends.
	ListOver *bool `json:"ListOver,omitempty" name:"ListOver"`

	// Credential for viewing more logs.
	NextToken *uint64 `json:"NextToken,omitempty" name:"NextToken"`

	// Logset.
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	Events []*Event `json:"Events,omitempty" name:"Events"`

	// Total number of events.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// Username
	Username *string `json:"Username,omitempty" name:"Username"`

	// Event Time
	EventTime *string `json:"EventTime,omitempty" name:"EventTime"`

	// Log details
	CloudAuditEvent *string `json:"CloudAuditEvent,omitempty" name:"CloudAuditEvent"`

	// Description of resource type in Chinese (please use this field as required; if you are using other languages, ignore this field)
	ResourceTypeCn *string `json:"ResourceTypeCn,omitempty" name:"ResourceTypeCn"`

	// Authentication error code
	ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`

	// Event name
	EventName *string `json:"EventName,omitempty" name:"EventName"`

	// Certificate ID
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	SecretId *string `json:"SecretId,omitempty" name:"SecretId"`

	// Request source
	EventSource *string `json:"EventSource,omitempty" name:"EventSource"`

	// Request ID
	RequestID *string `json:"RequestID,omitempty" name:"RequestID"`

	// Resource region
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// Root account ID
	AccountID *int64 `json:"AccountID,omitempty" name:"AccountID"`

	// Source IP
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	SourceIPAddress *string `json:"SourceIPAddress,omitempty" name:"SourceIPAddress"`

	// Description of event name in Chinese (please use this field as required; if you are using other languages, ignore this field)
	EventNameCn *string `json:"EventNameCn,omitempty" name:"EventNameCn"`

	// Resource pair
	Resources *Resource `json:"Resources,omitempty" name:"Resources"`

	// Event region
	EventRegion *string `json:"EventRegion,omitempty" name:"EventRegion"`

	// IP location
	Location *string `json:"Location,omitempty" name:"Location"`
}

// Predefined struct for user
type GetAttributeKeyRequestParams struct {
	// Website type. Valid values: zh, en. If this parameter is left empty, `zh` will be used by default
	WebsiteType *string `json:"WebsiteType,omitempty" name:"WebsiteType"`
}

type GetAttributeKeyRequest struct {
	*tchttp.BaseRequest
	
	// Website type. Valid values: zh, en. If this parameter is left empty, `zh` will be used by default
	WebsiteType *string `json:"WebsiteType,omitempty" name:"WebsiteType"`
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
	AttributeKeyDetails []*AttributeKeyDetail `json:"AttributeKeyDetails,omitempty" name:"AttributeKeyDetails"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	AuditAmount *int64 `json:"AuditAmount,omitempty" name:"AuditAmount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	AuditSummarys []*AuditSummary `json:"AuditSummarys,omitempty" name:"AuditSummarys"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	WebsiteType *string `json:"WebsiteType,omitempty" name:"WebsiteType"`
}

type ListCmqEnableRegionRequest struct {
	*tchttp.BaseRequest
	
	// Website type. zh: Chinese mainland (default); en: outside Chinese mainland.
	WebsiteType *string `json:"WebsiteType,omitempty" name:"WebsiteType"`
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
	EnableRegions []*CmqRegionInfo `json:"EnableRegions,omitempty" name:"EnableRegions"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	WebsiteType *string `json:"WebsiteType,omitempty" name:"WebsiteType"`
}

type ListCosEnableRegionRequest struct {
	*tchttp.BaseRequest
	
	// Website type. zh: Chinese mainland (default); en: outside Chinese mainland.
	WebsiteType *string `json:"WebsiteType,omitempty" name:"WebsiteType"`
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
	EnableRegions []*CosRegionInfo `json:"EnableRegions,omitempty" name:"EnableRegions"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Search criteria
	LookupAttributes []*LookupAttribute `json:"LookupAttributes,omitempty" name:"LookupAttributes"`

	// Credential for viewing more logs
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// Maximum number of logs to be returned
	MaxResults *int64 `json:"MaxResults,omitempty" name:"MaxResults"`

	// CloudAudit mode. Valid values: standard, quick. Default value: standard
	Mode *string `json:"Mode,omitempty" name:"Mode"`
}

type LookUpEventsRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Search criteria
	LookupAttributes []*LookupAttribute `json:"LookupAttributes,omitempty" name:"LookupAttributes"`

	// Credential for viewing more logs
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// Maximum number of logs to be returned
	MaxResults *int64 `json:"MaxResults,omitempty" name:"MaxResults"`

	// CloudAudit mode. Valid values: standard, quick. Default value: standard
	Mode *string `json:"Mode,omitempty" name:"Mode"`
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
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// Logset
	// Note: This field may return null, indicating that no valid values can be obtained.
	Events []*Event `json:"Events,omitempty" name:"Events"`

	// Whether the logset ends
	// Note: This field may return null, indicating that no valid values can be obtained.
	ListOver *bool `json:"ListOver,omitempty" name:"ListOver"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	AttributeKey *string `json:"AttributeKey,omitempty" name:"AttributeKey"`

	// Value of `AttributeValue`
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	AttributeValue *string `json:"AttributeValue,omitempty" name:"AttributeValue"`
}

// Predefined struct for user
type ModifyAuditTrackRequestParams struct {

}

type ModifyAuditTrackRequest struct {
	*tchttp.BaseRequest
	
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
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAuditTrackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAuditTrackResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// Resource name
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
}

// Predefined struct for user
type StartLoggingRequestParams struct {
	// Tracking set name
	AuditName *string `json:"AuditName,omitempty" name:"AuditName"`
}

type StartLoggingRequest struct {
	*tchttp.BaseRequest
	
	// Tracking set name
	AuditName *string `json:"AuditName,omitempty" name:"AuditName"`
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
	IsSuccess *int64 `json:"IsSuccess,omitempty" name:"IsSuccess"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	AuditName *string `json:"AuditName,omitempty" name:"AuditName"`
}

type StopLoggingRequest struct {
	*tchttp.BaseRequest
	
	// Tracking set name
	AuditName *string `json:"AuditName,omitempty" name:"AuditName"`
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
	IsSuccess *int64 `json:"IsSuccess,omitempty" name:"IsSuccess"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

// Predefined struct for user
type UpdateAuditRequestParams struct {
	// Tracking set name
	AuditName *string `json:"AuditName,omitempty" name:"AuditName"`

	// Whether to enable CMQ message notification. 1: Yes; 0: No. Only CMQ queue service is currently supported. If CMQ message notification is enabled, CloudAudit will deliver your log contents to the designated queue in the specified region in real time.
	IsEnableCmqNotify *int64 `json:"IsEnableCmqNotify,omitempty" name:"IsEnableCmqNotify"`

	// Manages the read/write attribute of event. Valid values: 1 (read-only), 2 (write-only), 3 (read/write).
	ReadWriteAttribute *int64 `json:"ReadWriteAttribute,omitempty" name:"ReadWriteAttribute"`

	// Globally unique ID of the CMK. This value is required if it is not a newly created KMS element. It can be obtained through `ListKeyAliasByRegion`. CloudAudit will not verify the validity of the `KeyId`. Enter it carefully to avoid data loss.
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// COS region. Supported regions can be queried through the `ListCosEnableRegion` API.
	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`

	// Queue name, which must begin with a letter and can contain up to 64 letters, digits, and dashes (-). This field is required if the value of `IsEnableCmqNotify` is 1. If a queue is not newly created, CloudAudit will not verify whether it actually exists. Enter the name with caution to avoid log delivery failure and consequent data loss.
	CmqQueueName *string `json:"CmqQueueName,omitempty" name:"CmqQueueName"`

	// Whether to create a COS bucket. 1: Yes; 0: No.
	IsCreateNewBucket *int64 `json:"IsCreateNewBucket,omitempty" name:"IsCreateNewBucket"`

	// KMS region. Currently supported regions can be obtained through `ListKmsEnableRegion`. This must be the same as the COS region.
	KmsRegion *string `json:"KmsRegion,omitempty" name:"KmsRegion"`

	// Whether to enable KMS encryption. 1: Yes, 0: No. If KMS encryption is enabled, the data will be encrypted when delivered to COS.
	IsEnableKmsEncry *int64 `json:"IsEnableKmsEncry,omitempty" name:"IsEnableKmsEncry"`

	// User-defined COS bucket name, which can only contain 1–40 lowercase letters (a–z), digits (0–9), and dashes (-) and cannot begin or end with "-". If a bucket is not newly created, CloudAudit will not verify whether it actually exists. Enter the name with caution to avoid log delivery failure and consequent data loss.
	CosBucketName *string `json:"CosBucketName,omitempty" name:"CosBucketName"`

	// Region where the queue is located. Supported CMQ regions can be queried through the `ListCmqEnableRegion` API. This field is required if the value of `IsEnableCmqNotify` is 1.
	CmqRegion *string `json:"CmqRegion,omitempty" name:"CmqRegion"`

	// Log file prefix, which can only contain 3–40 ASCII letters (a–z; A–Z) and digits (0–9).
	LogFilePrefix *string `json:"LogFilePrefix,omitempty" name:"LogFilePrefix"`

	// Whether to create a queue. 1: Yes; 0: No. This field is required if the value of `IsEnableCmqNotify` is 1.
	IsCreateNewQueue *int64 `json:"IsCreateNewQueue,omitempty" name:"IsCreateNewQueue"`
}

type UpdateAuditRequest struct {
	*tchttp.BaseRequest
	
	// Tracking set name
	AuditName *string `json:"AuditName,omitempty" name:"AuditName"`

	// Whether to enable CMQ message notification. 1: Yes; 0: No. Only CMQ queue service is currently supported. If CMQ message notification is enabled, CloudAudit will deliver your log contents to the designated queue in the specified region in real time.
	IsEnableCmqNotify *int64 `json:"IsEnableCmqNotify,omitempty" name:"IsEnableCmqNotify"`

	// Manages the read/write attribute of event. Valid values: 1 (read-only), 2 (write-only), 3 (read/write).
	ReadWriteAttribute *int64 `json:"ReadWriteAttribute,omitempty" name:"ReadWriteAttribute"`

	// Globally unique ID of the CMK. This value is required if it is not a newly created KMS element. It can be obtained through `ListKeyAliasByRegion`. CloudAudit will not verify the validity of the `KeyId`. Enter it carefully to avoid data loss.
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// COS region. Supported regions can be queried through the `ListCosEnableRegion` API.
	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`

	// Queue name, which must begin with a letter and can contain up to 64 letters, digits, and dashes (-). This field is required if the value of `IsEnableCmqNotify` is 1. If a queue is not newly created, CloudAudit will not verify whether it actually exists. Enter the name with caution to avoid log delivery failure and consequent data loss.
	CmqQueueName *string `json:"CmqQueueName,omitempty" name:"CmqQueueName"`

	// Whether to create a COS bucket. 1: Yes; 0: No.
	IsCreateNewBucket *int64 `json:"IsCreateNewBucket,omitempty" name:"IsCreateNewBucket"`

	// KMS region. Currently supported regions can be obtained through `ListKmsEnableRegion`. This must be the same as the COS region.
	KmsRegion *string `json:"KmsRegion,omitempty" name:"KmsRegion"`

	// Whether to enable KMS encryption. 1: Yes, 0: No. If KMS encryption is enabled, the data will be encrypted when delivered to COS.
	IsEnableKmsEncry *int64 `json:"IsEnableKmsEncry,omitempty" name:"IsEnableKmsEncry"`

	// User-defined COS bucket name, which can only contain 1–40 lowercase letters (a–z), digits (0–9), and dashes (-) and cannot begin or end with "-". If a bucket is not newly created, CloudAudit will not verify whether it actually exists. Enter the name with caution to avoid log delivery failure and consequent data loss.
	CosBucketName *string `json:"CosBucketName,omitempty" name:"CosBucketName"`

	// Region where the queue is located. Supported CMQ regions can be queried through the `ListCmqEnableRegion` API. This field is required if the value of `IsEnableCmqNotify` is 1.
	CmqRegion *string `json:"CmqRegion,omitempty" name:"CmqRegion"`

	// Log file prefix, which can only contain 3–40 ASCII letters (a–z; A–Z) and digits (0–9).
	LogFilePrefix *string `json:"LogFilePrefix,omitempty" name:"LogFilePrefix"`

	// Whether to create a queue. 1: Yes; 0: No. This field is required if the value of `IsEnableCmqNotify` is 1.
	IsCreateNewQueue *int64 `json:"IsCreateNewQueue,omitempty" name:"IsCreateNewQueue"`
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
	IsSuccess *int64 `json:"IsSuccess,omitempty" name:"IsSuccess"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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