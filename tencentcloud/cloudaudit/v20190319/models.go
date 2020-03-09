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

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type AttributeKeyDetail struct {

	// Chinese label
	Label *string `json:"Label,omitempty" name:"Label"`

	// Input box type
	LabelType *string `json:"LabelType,omitempty" name:"LabelType"`

	// Display sort order
	Order *int64 `json:"Order,omitempty" name:"Order"`

	// Initial display
	Starter *string `json:"Starter,omitempty" name:"Starter"`

	// AttributeKey value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type AuditSummary struct {

	// Tracking set name
	AuditName *string `json:"AuditName,omitempty" name:"AuditName"`

	// Tracking set status. Value range: 1 (enabled), 0 (disabled)
	AuditStatus *int64 `json:"AuditStatus,omitempty" name:"AuditStatus"`

	// COS bucket name
	CosBucketName *string `json:"CosBucketName,omitempty" name:"CosBucketName"`

	// Log prefix
	LogFilePrefix *string `json:"LogFilePrefix,omitempty" name:"LogFilePrefix"`
}

type CmqRegionInfo struct {

	// CMQ region
	CmqRegion *string `json:"CmqRegion,omitempty" name:"CmqRegion"`

	// Region description
	CmqRegionName *string `json:"CmqRegionName,omitempty" name:"CmqRegionName"`
}

type CosRegionInfo struct {

	// COS region
	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`

	// Region description
	CosRegionName *string `json:"CosRegionName,omitempty" name:"CosRegionName"`
}

type CreateAuditRequest struct {
	*tchttp.BaseRequest

	// Tracking set name, which can contain 3-128 ASCII letters (a-z; A-Z), digits (0-9), and underscores (_).
	AuditName *string `json:"AuditName,omitempty" name:"AuditName"`

	// User-defined COS bucket name, which can only contain 1-40 lowercase letters (a-z), digits (0-9), and dashes (-) and cannot begin or end with "-". If a bucket is not newly created, CloudAudit will not verify whether it actually exists. Please enter the name with caution so as to avoid log delivery failure and consequent data loss.
	CosBucketName *string `json:"CosBucketName,omitempty" name:"CosBucketName"`

	// COS region. Supported regions can be queried using the ListCosEnableRegion API.
	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`

	// Whether to create a COS bucket. 1: yes; 0: no.
	IsCreateNewBucket *int64 `json:"IsCreateNewBucket,omitempty" name:"IsCreateNewBucket"`

	// Whether to enable CMQ message notification. 1: yes; 0: no. Only CMQ queue service is currently supported. If CMQ message notification is enabled, CloudAudit will deliver your log contents to the designated queue in the specified region in real time.
	IsEnableCmqNotify *int64 `json:"IsEnableCmqNotify,omitempty" name:"IsEnableCmqNotify"`

	// Manages the read/write attribute of an event. Value range: 1 (read-only), 2 (write-only), 3 (read/write).
	ReadWriteAttribute *int64 `json:"ReadWriteAttribute,omitempty" name:"ReadWriteAttribute"`

	// Queue name, which must begin with a letter and can contain up to 64 letters, digits, and dashes (-). This field is required if the value of IsEnableCmqNotify is 1. If a queue is not newly created, CloudAudit will not verify whether it actually exists. Please enter the name with caution so as to avoid log delivery failure and consequent data loss.
	CmqQueueName *string `json:"CmqQueueName,omitempty" name:"CmqQueueName"`

	// Region where the queue is located. Supported CMQ regions can be queried using the ListCmqEnableRegion API. This field is required if the value of IsEnableCmqNotify is 1.
	CmqRegion *string `json:"CmqRegion,omitempty" name:"CmqRegion"`

	// Whether to create a queue. 1: yes; 0: no. This field is required if the value of IsEnableCmqNotify is 1.
	IsCreateNewQueue *int64 `json:"IsCreateNewQueue,omitempty" name:"IsCreateNewQueue"`

	// Prefix of a log file, which can only contain 3-40 ASCII letters (a-z; A-Z) and digits (0-9). It can be left empty and is the account ID by default.
	LogFilePrefix *string `json:"LogFilePrefix,omitempty" name:"LogFilePrefix"`
}

func (r *CreateAuditRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAuditRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAuditResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Whether creation is successful.
		IsSuccess *int64 `json:"IsSuccess,omitempty" name:"IsSuccess"`

		// Unique ID of the request. Each request returns a unique ID. The RequestId is required to troubleshoot issues.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAuditResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAuditResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DeleteAuditRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAuditResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Whether deletion is successful
		IsSuccess *int64 `json:"IsSuccess,omitempty" name:"IsSuccess"`

		// Unique ID of the request. Each request returns a unique ID. The RequestId is required to troubleshoot issues.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAuditResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAuditResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeAuditRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAuditResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Tracking set name.
		AuditName *string `json:"AuditName,omitempty" name:"AuditName"`

		// Tracking set status. Value range: 1 (enabled), 0 (disabled).
		AuditStatus *int64 `json:"AuditStatus,omitempty" name:"AuditStatus"`

		// Queue name.
		CmqQueueName *string `json:"CmqQueueName,omitempty" name:"CmqQueueName"`

		// Region where the queue is located.
		CmqRegion *string `json:"CmqRegion,omitempty" name:"CmqRegion"`

		// COS bucket name.
		CosBucketName *string `json:"CosBucketName,omitempty" name:"CosBucketName"`

		// Region where the COS bucket is located.
		CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`

		// Whether to enable CMQ message notification. 1: yes; 0: no.
		IsEnableCmqNotify *int64 `json:"IsEnableCmqNotify,omitempty" name:"IsEnableCmqNotify"`

		// Log prefix.
		LogFilePrefix *string `json:"LogFilePrefix,omitempty" name:"LogFilePrefix"`

		// Manages the read/write attribute of an event. Value range: 1 (read-only), 2 (write-only), 3 (read/write)
		ReadWriteAttribute *int64 `json:"ReadWriteAttribute,omitempty" name:"ReadWriteAttribute"`

		// Unique ID of the request. Each request returns a unique ID. The RequestId is required to troubleshoot issues.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAuditResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAuditResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Event struct {

	// Resource pair
	Resources *Resource `json:"Resources,omitempty" name:"Resources"`

	// Root account ID
	AccountID *int64 `json:"AccountID,omitempty" name:"AccountID"`

	// Log details
	CloudAuditEvent *string `json:"CloudAuditEvent,omitempty" name:"CloudAuditEvent"`

	// Authentication error code
	ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`

	// Log ID
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// Event name
	EventName *string `json:"EventName,omitempty" name:"EventName"`

	// Description of event name
	EventNameCn *string `json:"EventNameCn,omitempty" name:"EventNameCn"`

	// Event region
	EventRegion *string `json:"EventRegion,omitempty" name:"EventRegion"`

	// Request source
	EventSource *string `json:"EventSource,omitempty" name:"EventSource"`

	// Event time
	EventTime *string `json:"EventTime,omitempty" name:"EventTime"`

	// Request ID
	RequestID *string `json:"RequestID,omitempty" name:"RequestID"`

	// Description of resource type
	ResourceTypeCn *string `json:"ResourceTypeCn,omitempty" name:"ResourceTypeCn"`

	// Certificate ID
	SecretId *string `json:"SecretId,omitempty" name:"SecretId"`

	// Source IP
	SourceIPAddress *string `json:"SourceIPAddress,omitempty" name:"SourceIPAddress"`

	// Username
	Username *string `json:"Username,omitempty" name:"Username"`
}

type GetAttributeKeyRequest struct {
	*tchttp.BaseRequest

	// Website type. Value range: zh, en. Default value: zh
	WebsiteType *string `json:"WebsiteType,omitempty" name:"WebsiteType"`
}

func (r *GetAttributeKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetAttributeKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetAttributeKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// AttributeKey value range
		AttributeKeyDetails []*AttributeKeyDetail `json:"AttributeKeyDetails,omitempty" name:"AttributeKeyDetails" list`

		// Unique ID of the request. Each request returns a unique ID. The RequestId is required to troubleshoot issues.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAttributeKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetAttributeKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquireAuditCreditRequest struct {
	*tchttp.BaseRequest
}

func (r *InquireAuditCreditRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquireAuditCreditRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquireAuditCreditResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of tracking sets that can be created
		AuditAmount *int64 `json:"AuditAmount,omitempty" name:"AuditAmount"`

		// Unique ID of the request. Each request returns a unique ID. The RequestId is required to troubleshoot issues.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquireAuditCreditResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquireAuditCreditResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListAuditsRequest struct {
	*tchttp.BaseRequest
}

func (r *ListAuditsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListAuditsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListAuditsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Queries the summary set of tracking sets
		AuditSummarys []*AuditSummary `json:"AuditSummarys,omitempty" name:"AuditSummarys" list`

		// Unique ID of the request. Each request returns a unique ID. The RequestId is required to troubleshoot issues.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListAuditsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListAuditsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListCmqEnableRegionRequest struct {
	*tchttp.BaseRequest

	// Website type. zh: Mainland China (default); en: outside Mainland China.
	WebsiteType *string `json:"WebsiteType,omitempty" name:"WebsiteType"`
}

func (r *ListCmqEnableRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListCmqEnableRegionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListCmqEnableRegionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// CloudAudit-enabled CMQ AZs
		EnableRegions []*CmqRegionInfo `json:"EnableRegions,omitempty" name:"EnableRegions" list`

		// Unique ID of the request. Each request returns a unique ID. The RequestId is required to troubleshoot issues.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListCmqEnableRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListCmqEnableRegionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListCosEnableRegionRequest struct {
	*tchttp.BaseRequest

	// Website type. zh: Mainland China (default); en: outside Mainland China.
	WebsiteType *string `json:"WebsiteType,omitempty" name:"WebsiteType"`
}

func (r *ListCosEnableRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListCosEnableRegionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListCosEnableRegionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// CloudAudit-enabled COS AZs
		EnableRegions []*CosRegionInfo `json:"EnableRegions,omitempty" name:"EnableRegions" list`

		// Unique ID of the request. Each request returns a unique ID. The RequestId is required to troubleshoot issues.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListCosEnableRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListCosEnableRegionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LookUpEventsRequest struct {
	*tchttp.BaseRequest

	// End time
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Start time
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// Search criteria
	LookupAttributes []*LookupAttribute `json:"LookupAttributes,omitempty" name:"LookupAttributes" list`

	// Maximum number of logs that can be returned
	MaxResults *int64 `json:"MaxResults,omitempty" name:"MaxResults"`

	// Credential for viewing more logs
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *LookUpEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *LookUpEventsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LookUpEventsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Logset
		Events []*Event `json:"Events,omitempty" name:"Events" list`

		// Whether the logset ends
		ListOver *bool `json:"ListOver,omitempty" name:"ListOver"`

		// Credential for viewing more logs
		NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

		// Unique ID of the request. Each request returns a unique ID. The RequestId is required to troubleshoot issues.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LookUpEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *LookUpEventsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LookupAttribute struct {

	// AttributeKey value range: RequestId, EventName, ReadOnly, Username, ResourceType, ResourceName, AccessKeyId, EventId
	AttributeKey *string `json:"AttributeKey,omitempty" name:"AttributeKey"`

	// AttributeValue
	AttributeValue *string `json:"AttributeValue,omitempty" name:"AttributeValue"`
}

type Resource struct {

	// Resource name
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// Resource type
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
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

func (r *StartLoggingRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartLoggingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Whether it is successfully enabled
		IsSuccess *int64 `json:"IsSuccess,omitempty" name:"IsSuccess"`

		// Unique ID of the request. Each request returns a unique ID. The RequestId is required to troubleshoot issues.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartLoggingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartLoggingResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *StopLoggingRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopLoggingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Whether it is successfully disabled
		IsSuccess *int64 `json:"IsSuccess,omitempty" name:"IsSuccess"`

		// Unique ID of the request. Each request returns a unique ID. The RequestId is required to troubleshoot issues.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopLoggingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopLoggingResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateAuditRequest struct {
	*tchttp.BaseRequest

	// Tracking set name
	AuditName *string `json:"AuditName,omitempty" name:"AuditName"`

	// Queue name, which must begin with a letter and can contain up to 64 letters, digits, and dashes (-). This field is required if the value of IsEnableCmqNotify is 1. If a queue is not newly created, CloudAudit will not verify whether it actually exists. Please enter the name with caution so as to avoid log delivery failure and consequent data loss.
	CmqQueueName *string `json:"CmqQueueName,omitempty" name:"CmqQueueName"`

	// Region where the queue is located. Supported CMQ regions can be queried using the ListCmqEnableRegion API. This field is required if the value of IsEnableCmqNotify is 1.
	CmqRegion *string `json:"CmqRegion,omitempty" name:"CmqRegion"`

	// User-defined COS bucket name, which can only contain 1-40 lowercase letters (a-z), digits (0-9), and dashes (-) and cannot begin or end with "-". If a bucket is not newly created, CloudAudit will not verify whether it actually exists. Please enter the name with caution so as to avoid log delivery failure and consequent data loss.
	CosBucketName *string `json:"CosBucketName,omitempty" name:"CosBucketName"`

	// COS region. Supported regions can be queried using the ListCosEnableRegion API.
	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`

	// Whether to create a COS bucket. 1: yes; 0: no.
	IsCreateNewBucket *int64 `json:"IsCreateNewBucket,omitempty" name:"IsCreateNewBucket"`

	// Whether to create a queue. 1: yes; 0: no. This field is required if the value of IsEnableCmqNotify is 1.
	IsCreateNewQueue *int64 `json:"IsCreateNewQueue,omitempty" name:"IsCreateNewQueue"`

	// Whether to enable CMQ message notification. 1: yes; 0: no. Only CMQ queue service is currently supported. If CMQ message notification is enabled, CloudAudit will deliver your log contents to the designated queue in the specified region in real time.
	IsEnableCmqNotify *int64 `json:"IsEnableCmqNotify,omitempty" name:"IsEnableCmqNotify"`

	// Prefix of a log file, which can only contain 3-40 ASCII letters (a-z; A-Z) and digits (0-9).
	LogFilePrefix *string `json:"LogFilePrefix,omitempty" name:"LogFilePrefix"`

	// Manages the read/write attribute of an event. Value range: 1 (read-only), 2 (write-only), 3 (read/write).
	ReadWriteAttribute *int64 `json:"ReadWriteAttribute,omitempty" name:"ReadWriteAttribute"`
}

func (r *UpdateAuditRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateAuditRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateAuditResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Whether update is successful
		IsSuccess *int64 `json:"IsSuccess,omitempty" name:"IsSuccess"`

		// Unique ID of the request. Each request returns a unique ID. The RequestId is required to troubleshoot issues.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateAuditResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateAuditResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
