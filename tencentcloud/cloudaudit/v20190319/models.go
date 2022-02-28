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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

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

type DescribeAuditTracksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeEventsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type LookupAttribute struct {

	// Valid values: RequestId, EventName, ReadOnly, Username, ResourceType, ResourceName, AccessKeyId, and EventId
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	AttributeKey *string `json:"AttributeKey,omitempty" name:"AttributeKey"`

	// Value of `AttributeValue`
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	AttributeValue *string `json:"AttributeValue,omitempty" name:"AttributeValue"`
}

type Resource struct {

	// Resource type
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// Resource name
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
}
