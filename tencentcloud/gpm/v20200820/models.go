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

package v20200820

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type AttributeMap struct {

	// Map key, supporting [a-zA-Z0-9-\.]*
	Key *string `json:"Key,omitempty" name:"Key"`

	// Map value
	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type CancelMatchingRequest struct {
	*tchttp.BaseRequest

	// MatchCode
	MatchCode *string `json:"MatchCode,omitempty" name:"MatchCode"`

	// The MatchTicket ID of the matchmaking to cancel
	MatchTicketId *string `json:"MatchTicketId,omitempty" name:"MatchTicketId"`
}

func (r *CancelMatchingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelMatchingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MatchCode")
	delete(f, "MatchTicketId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelMatchingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CancelMatchingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Error code
		ErrCode *uint64 `json:"ErrCode,omitempty" name:"ErrCode"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelMatchingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelMatchingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMatchRequest struct {
	*tchttp.BaseRequest

	// Match name. It can contain up to 128 bytes, supporting [a-zA-Z0-9-\.]*.
	MatchName *string `json:"MatchName,omitempty" name:"MatchName"`

	// RuleCode
	RuleCode *string `json:"RuleCode,omitempty" name:"RuleCode"`

	// Timeout period in seconds. Value range: 1 600
	Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`

	// Whether to request server resources for the matchmaking results. 0: no, 1: request GSE resources
	ServerType *int64 `json:"ServerType,omitempty" name:"ServerType"`

	// Matchmaking description. Up to 1024 bytes are allowed.
	MatchDesc *string `json:"MatchDesc,omitempty" name:"MatchDesc"`

	// Only HTTP and HTTPS protocols are supported.
	NotifyUrl *string `json:"NotifyUrl,omitempty" name:"NotifyUrl"`

	// Region of the game server queue
	ServerRegion *string `json:"ServerRegion,omitempty" name:"ServerRegion"`

	// Game server queue
	ServerQueue *string `json:"ServerQueue,omitempty" name:"ServerQueue"`

	// Custom push data
	CustomPushData *string `json:"CustomPushData,omitempty" name:"CustomPushData"`

	// Game server session data
	ServerSessionData *string `json:"ServerSessionData,omitempty" name:"ServerSessionData"`

	// Game attribute. It is an array of key-value structure.
	GameProperties []*StringKV `json:"GameProperties,omitempty" name:"GameProperties"`

	// Enable or disable the log. 0: disable, 1: enable
	LogSwitch *int64 `json:"LogSwitch,omitempty" name:"LogSwitch"`

	// Tag. It is an array of key-value structure.
	Tags []*StringKV `json:"Tags,omitempty" name:"Tags"`
}

func (r *CreateMatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MatchName")
	delete(f, "RuleCode")
	delete(f, "Timeout")
	delete(f, "ServerType")
	delete(f, "MatchDesc")
	delete(f, "NotifyUrl")
	delete(f, "ServerRegion")
	delete(f, "ServerQueue")
	delete(f, "CustomPushData")
	delete(f, "ServerSessionData")
	delete(f, "GameProperties")
	delete(f, "LogSwitch")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateMatchResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Matchmaking information
		MatchInfo *MatchInfo `json:"MatchInfo,omitempty" name:"MatchInfo"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRuleRequest struct {
	*tchttp.BaseRequest

	// Rule name. It can contain up to 128 bytes, supporting [a-zA-Z0-9-\.]*.
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// Rule script. Up to 65535 bytes are allowed.
	RuleScript *string `json:"RuleScript,omitempty" name:"RuleScript"`

	// Rule description. Up to 1024 bytes are allowed.
	RuleDesc *string `json:"RuleDesc,omitempty" name:"RuleDesc"`

	// Tag. It is an array of key-value structure. Up to 50 tags can be associated.
	Tags []*StringKV `json:"Tags,omitempty" name:"Tags"`
}

func (r *CreateRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleName")
	delete(f, "RuleScript")
	delete(f, "RuleDesc")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Rule information
		RuleInfo *RuleInfo `json:"RuleInfo,omitempty" name:"RuleInfo"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMatchRequest struct {
	*tchttp.BaseRequest

	// MatchCode
	MatchCode *string `json:"MatchCode,omitempty" name:"MatchCode"`
}

func (r *DeleteMatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MatchCode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMatchResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRuleRequest struct {
	*tchttp.BaseRequest

	// RuleCode
	RuleCode *string `json:"RuleCode,omitempty" name:"RuleCode"`
}

func (r *DeleteRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleCode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataRequest struct {
	*tchttp.BaseRequest

	// Start time in seconds
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// End time in seconds
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Time granularity. Valid values: 1: 1 day, 2: 1 hour, 3: 1 minute, 4: 10 minutes, 5: 30 minutes
	TimeType *int64 `json:"TimeType,omitempty" name:"TimeType"`

	// MatchCode
	MatchCode *string `json:"MatchCode,omitempty" name:"MatchCode"`
}

func (r *DescribeDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "TimeType")
	delete(f, "MatchCode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Matchmaking statistics overview
	// Note: this field may return `null`, indicating that no valid value is obtained.
		OverviewData *ReportOverviewData `json:"OverviewData,omitempty" name:"OverviewData"`

		// Trend data of the number of matchmaking requests
	// Note: this field may return `null`, indicating that no valid value is obtained.
		TrendData *ReportTrendData `json:"TrendData,omitempty" name:"TrendData"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMatchCodesRequest struct {
	*tchttp.BaseRequest

	// Offset, number of pages.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// The number of MatchCodes per page
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Query by the MatchCode value (a string).
	MatchCode *string `json:"MatchCode,omitempty" name:"MatchCode"`
}

func (r *DescribeMatchCodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMatchCodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "MatchCode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMatchCodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMatchCodesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// MatchCode
	// Note: this field may return `null`, indicating that no valid value is obtained.
		MatchCodes []*MatchCodeAttr `json:"MatchCodes,omitempty" name:"MatchCodes"`

		// The total number of queried MatchCodes
	// Note: this field may return `null`, indicating that no valid value is obtained.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMatchCodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMatchCodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMatchRequest struct {
	*tchttp.BaseRequest

	// MatchCode
	MatchCode *string `json:"MatchCode,omitempty" name:"MatchCode"`
}

func (r *DescribeMatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MatchCode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMatchResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Matchmaking information
	// Note: this field may return `null`, indicating that no valid value is obtained.
		MatchInfo *MatchInfo `json:"MatchInfo,omitempty" name:"MatchInfo"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMatchesRequest struct {
	*tchttp.BaseRequest

	// The current page number. If this parameter is left empty, all queried matches will be obtained.
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// Number of matchmaking lists per page. If this parameter is left empty, all queried matches will be obtained.
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// Query type (optional). Valid values: match (query by matchCode or matchName), rule (query by ruleCode or ruleName), and other types (not filtered)
	SearchType *string `json:"SearchType,omitempty" name:"SearchType"`

	// Keyword. Enter a keyword about SearchType to query.
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// Tags. Enter a tag for querying.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

func (r *DescribeMatchesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMatchesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "SearchType")
	delete(f, "Keyword")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMatchesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMatchesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Matchmaking information list
	// Note: this field may return `null`, indicating that no valid value is obtained.
		MatchInfoList []*MatchInfo `json:"MatchInfoList,omitempty" name:"MatchInfoList"`

		// Total number of results
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The current page number. The first page will be returned by default if this parameter is left empty.
		PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

		// The number of matches per page. If this parameter is left empty, 30 matches are displayed per page by default. Maximum value: 30
		PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

		// Query type (optional). Valid values: matchName (query by match name), matchCode (query by matchCode), ruleName (query by rule name), and tag (query by tag key/value)
		SearchType *string `json:"SearchType,omitempty" name:"SearchType"`

		// Keyword for querying (optional)
		Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMatchesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMatchesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMatchingProgressRequest struct {
	*tchttp.BaseRequest

	// List of MatchTicket IDs. It can contain up to 12 IDs.
	MatchTicketIds []*MTicket `json:"MatchTicketIds,omitempty" name:"MatchTicketIds"`
}

func (r *DescribeMatchingProgressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMatchingProgressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MatchTicketIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMatchingProgressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMatchingProgressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// MatchTicket list
	// Note: this field may return `null`, indicating that no valid value is obtained.
		MatchTickets []*MatchTicket `json:"MatchTickets,omitempty" name:"MatchTickets"`

		// Error code
	// Note: this field may return `null`, indicating that no valid value is obtained.
		ErrCode *uint64 `json:"ErrCode,omitempty" name:"ErrCode"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMatchingProgressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMatchingProgressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRuleRequest struct {
	*tchttp.BaseRequest

	// RuleCode
	RuleCode *string `json:"RuleCode,omitempty" name:"RuleCode"`
}

func (r *DescribeRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleCode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Rule information
	// Note: this field may return `null`, indicating that no valid value is obtained.
		RuleInfo *RuleInfo `json:"RuleInfo,omitempty" name:"RuleInfo"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRulesRequest struct {
	*tchttp.BaseRequest

	// The current page number. The first page will be returned if this parameter is left empty.
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// The number of rules per page. If this parameter is left empty, 30 rules are displayed per page by default. Maximum value: 30
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// Query type (optional). Valid values: match (query by matchCode or matchName), rule (query by ruleCode or ruleName), and other types (not filtered)
	SearchType *string `json:"SearchType,omitempty" name:"SearchType"`

	// Keyword. Enter a keyword about SearchType to query.
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// Tags. Enter a tag for querying.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

func (r *DescribeRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "SearchType")
	delete(f, "Keyword")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Rule information list
	// Note: this field may return `null`, indicating that no valid value is obtained.
		RuleInfoList []*RuleBriefInfo `json:"RuleInfoList,omitempty" name:"RuleInfoList"`

		// Total number of results
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The current page number
		PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

		// Number of rules per page
		PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

		// Query type (optional). Valid values: matchName (query by match name), matchCode (query by matchCode), ruleName (query by rule name), and tag (query by tag key/value)
		SearchType *string `json:"SearchType,omitempty" name:"SearchType"`

		// Keyword for querying (optional)
		Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTokenRequest struct {
	*tchttp.BaseRequest

	// MatchCode
	MatchCode *string `json:"MatchCode,omitempty" name:"MatchCode"`
}

func (r *DescribeTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MatchCode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTokenResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The token corresponding to the current MatchCode. If the current MatchCode does not have a token, this parameter may not obtain a valid value.
	// Note: this field may return `null`, indicating that no valid value is obtained.
		MatchToken *string `json:"MatchToken,omitempty" name:"MatchToken"`

		// The time period during which GPM will continuously push the original token in seconds when the token is replaced.
	// Note: this field may return `null`, indicating that no valid value is obtained.
		CompatibleSpan *uint64 `json:"CompatibleSpan,omitempty" name:"CompatibleSpan"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MTicket struct {

	// MatchCode
	MatchCode *string `json:"MatchCode,omitempty" name:"MatchCode"`

	// MatchTicket ID
	MatchTicketId *string `json:"MatchTicketId,omitempty" name:"MatchTicketId"`
}

type MatchAttribute struct {

	// Attribute name. It can contain up to 128 characters, supporting [a-zA-Z0-9-\.]*.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Attribute type. 0: number, 1: string. Default value: 0
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// Numeric attribute value. Default value: 0.0
	NumberValue *float64 `json:"NumberValue,omitempty" name:"NumberValue"`

	// String attribute value. Up to 128 characters are allowed. Default value: ""
	StringValue *string `json:"StringValue,omitempty" name:"StringValue"`

	// List attribute value
	ListValue []*string `json:"ListValue,omitempty" name:"ListValue"`

	// Map attribute value
	MapValue []*AttributeMap `json:"MapValue,omitempty" name:"MapValue"`
}

type MatchCodeAttr struct {

	// MatchCode
	// Note: this field may return `null`, indicating that no valid value is obtained.
	MatchCode *string `json:"MatchCode,omitempty" name:"MatchCode"`
}

type MatchInfo struct {

	// MatchCode
	MatchCode *string `json:"MatchCode,omitempty" name:"MatchCode"`

	// Match name
	MatchName *string `json:"MatchName,omitempty" name:"MatchName"`

	// Matchmaking description
	// Note: this field may return `null`, indicating that no valid value is obtained.
	MatchDesc *string `json:"MatchDesc,omitempty" name:"MatchDesc"`

	// RuleCode
	RuleCode *string `json:"RuleCode,omitempty" name:"RuleCode"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Timeout period
	Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`

	// Receiving notification address
	NotifyUrl *string `json:"NotifyUrl,omitempty" name:"NotifyUrl"`

	// Whether to request server resources for the match results. 0: no, 1: request GSE resources
	ServerType *int64 `json:"ServerType,omitempty" name:"ServerType"`

	// Region of the server queue
	// Note: this field may return `null`, indicating that no valid value is obtained.
	ServerRegion *string `json:"ServerRegion,omitempty" name:"ServerRegion"`

	// Server queue
	// Note: this field may return `null`, indicating that no valid value is obtained.
	ServerQueue *string `json:"ServerQueue,omitempty" name:"ServerQueue"`

	// Custom push data
	// Note: this field may return `null`, indicating that no valid value is obtained.
	CustomPushData *string `json:"CustomPushData,omitempty" name:"CustomPushData"`

	// Game server session data
	// Note: this field may return `null`, indicating that no valid value is obtained.
	ServerSessionData *string `json:"ServerSessionData,omitempty" name:"ServerSessionData"`

	// Game attributes
	// Note: this field may return `null`, indicating that no valid value is obtained.
	GameProperties []*StringKV `json:"GameProperties,omitempty" name:"GameProperties"`

	// Enable or disable the log. 0: disable, 1: enable
	LogSwitch *int64 `json:"LogSwitch,omitempty" name:"LogSwitch"`

	// Logset ID
	// Note: this field may return `null`, indicating that no valid value is obtained.
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// Logset name
	// Note: this field may return `null`, indicating that no valid value is obtained.
	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`

	// Log topic ID
	// Note: this field may return `null`, indicating that no valid value is obtained.
	LogTopicId *string `json:"LogTopicId,omitempty" name:"LogTopicId"`

	// Log topic name
	// Note: this field may return `null`, indicating that no valid value is obtained.
	LogTopicName *string `json:"LogTopicName,omitempty" name:"LogTopicName"`

	// Tag
	// Note: this field may return `null`, indicating that no valid value is obtained.
	Tags []*StringKV `json:"Tags,omitempty" name:"Tags"`

	// Region
	// Note: this field may return `null`, indicating that no valid value is obtained.
	Region *string `json:"Region,omitempty" name:"Region"`

	// User AppId
	// Note: this field may return `null`, indicating that no valid value is obtained.
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// User root account UIN
	// Note: this field may return `null`, indicating that no valid value is obtained.
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// Create user account UIN
	// Note: this field may return `null`, indicating that no valid value is obtained.
	CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`

	// Rule Name
	// Note: this field may return `null`, indicating that no valid value is obtained.
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// Log status. 0: normal, 1: the log set does not exist, 2: the log topic does not exist, 3: neither the log set nor the log topic exists.
	// Note: this field may return `null`, indicating that no valid value is obtained.
	LogStatus *int64 `json:"LogStatus,omitempty" name:"LogStatus"`
}

type MatchTicket struct {

	// MatchTicket ID. It can contain up to 128 characters, supporting [a-zA-Z0-9-\.]*.
	Id *string `json:"Id,omitempty" name:"Id"`

	// MatchCode
	MatchCode *string `json:"MatchCode,omitempty" name:"MatchCode"`

	// Different structure serialized results will be returned according to the MatchType.
	// Note: this field may return `null`, indicating that no valid value is obtained.
	MatchResult *string `json:"MatchResult,omitempty" name:"MatchResult"`

	// Matchmaking type. Valid values: NORMAL, GSE
	// Note: this field may return `null`, indicating that no valid value is obtained.
	MatchType *string `json:"MatchType,omitempty" name:"MatchType"`

	// Player information list
	Players []*Player `json:"Players,omitempty" name:"Players"`

	// Matching status. Valid values: SEARCHING (matching), PLACING (pending match), COMPLETED (match completed), CANCELLED (match cancelled), TIMEDOUT (match timeouts), FAILED (match failed)
	Status *string `json:"Status,omitempty" name:"Status"`

	// Matching status descriptions
	// Note: this field may return `null`, indicating that no valid value is obtained.
	StatusMessage *string `json:"StatusMessage,omitempty" name:"StatusMessage"`

	// Reason for the matching status
	// Note: this field may return `null`, indicating that no valid value is obtained.
	StatusReason *string `json:"StatusReason,omitempty" name:"StatusReason"`

	// The time when the GPM received the matchmaking request, for example "2020-08-17T08:14:38.077Z".
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The time when the matchmaking request stopped executing due to the completion, failure, timeout, or cancellation, for example "2020-08-17T08:14:38.077Z".
	// Note: this field may return `null`, indicating that no valid value is obtained.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type ModifyMatchRequest struct {
	*tchttp.BaseRequest

	// Match name. It can contain up to 128 bytes, supporting [a-zA-Z0-9-\.]*.
	MatchName *string `json:"MatchName,omitempty" name:"MatchName"`

	// RuleCode
	RuleCode *string `json:"RuleCode,omitempty" name:"RuleCode"`

	// Timeout period in seconds. Value range: 1 600
	Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`

	// Whether to request server resources for the matchmaking results. 0: no, 1: request GSE resources
	ServerType *int64 `json:"ServerType,omitempty" name:"ServerType"`

	// MatchCode
	MatchCode *string `json:"MatchCode,omitempty" name:"MatchCode"`

	// Matchmaking description. Up to 1024 bytes are allowed.
	MatchDesc *string `json:"MatchDesc,omitempty" name:"MatchDesc"`

	// Only HTTP and HTTPS protocols are supported.
	NotifyUrl *string `json:"NotifyUrl,omitempty" name:"NotifyUrl"`

	// Region of the game server queue
	ServerRegion *string `json:"ServerRegion,omitempty" name:"ServerRegion"`

	// Game server queue
	ServerQueue *string `json:"ServerQueue,omitempty" name:"ServerQueue"`

	// Custom push data
	CustomPushData *string `json:"CustomPushData,omitempty" name:"CustomPushData"`

	// Game server session data
	ServerSessionData *string `json:"ServerSessionData,omitempty" name:"ServerSessionData"`

	// Game attribute. It is an array of key-value structure.
	GameProperties []*StringKV `json:"GameProperties,omitempty" name:"GameProperties"`

	// Enable or disable the log. 0: disable, 1: enable
	LogSwitch *int64 `json:"LogSwitch,omitempty" name:"LogSwitch"`

	// Tag. It is an array of key-value structure.
	Tags []*StringKV `json:"Tags,omitempty" name:"Tags"`
}

func (r *ModifyMatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MatchName")
	delete(f, "RuleCode")
	delete(f, "Timeout")
	delete(f, "ServerType")
	delete(f, "MatchCode")
	delete(f, "MatchDesc")
	delete(f, "NotifyUrl")
	delete(f, "ServerRegion")
	delete(f, "ServerQueue")
	delete(f, "CustomPushData")
	delete(f, "ServerSessionData")
	delete(f, "GameProperties")
	delete(f, "LogSwitch")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMatchResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Matchmaking information
		MatchInfo *MatchInfo `json:"MatchInfo,omitempty" name:"MatchInfo"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRuleRequest struct {
	*tchttp.BaseRequest

	// RuleCode
	RuleCode *string `json:"RuleCode,omitempty" name:"RuleCode"`

	// Rule name. It can only contain numbers, letters, ".", and "-".
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// Rule description. Up to 1024 bytes are allowed.
	RuleDesc *string `json:"RuleDesc,omitempty" name:"RuleDesc"`

	// Tag. It is an array of key-value structure. Up to 50 tags can be associated.
	Tags []*StringKV `json:"Tags,omitempty" name:"Tags"`
}

func (r *ModifyRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleCode")
	delete(f, "RuleName")
	delete(f, "RuleDesc")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Rule information
		RuleInfo *RuleInfo `json:"RuleInfo,omitempty" name:"RuleInfo"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTokenRequest struct {
	*tchttp.BaseRequest

	// MatchCode
	MatchCode *string `json:"MatchCode,omitempty" name:"MatchCode"`

	// The time period during which GPM will continuously push the original token in seconds to the user when the token is replaced. Value range: 0 - 1800. Within the CompatibleSpan time period, user will receive the current and original token in the event notification.
	CompatibleSpan *uint64 `json:"CompatibleSpan,omitempty" name:"CompatibleSpan"`

	// Token. It can contain 0 - 64 characters, supporting [a-zA-Z0-9-_.]. If this parameter is left empty, the token will be randomly generated by GPM.
	MatchToken *string `json:"MatchToken,omitempty" name:"MatchToken"`
}

func (r *ModifyTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MatchCode")
	delete(f, "CompatibleSpan")
	delete(f, "MatchToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTokenResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Token that has been set successfully.
		MatchToken *string `json:"MatchToken,omitempty" name:"MatchToken"`

		// The time period during which GPM will continuously push the original token in seconds to the user when the token is replaced.
	// Note: this field may return `null`, indicating that no valid value is obtained.
		CompatibleSpan *uint64 `json:"CompatibleSpan,omitempty" name:"CompatibleSpan"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Player struct {

	// Player ID. It can contain up to 128 characters, supporting [a-zA-Z\d-\._]*.
	Id *string `json:"Id,omitempty" name:"Id"`

	// Player nickname. Up to 128 characters are allowed.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Player attribute for matching. Up to 10 attributes are supported.
	MatchAttributes []*MatchAttribute `json:"MatchAttributes,omitempty" name:"MatchAttributes"`

	// Team name. A player can pass in a different team name, which can contain up to 128 characters, and support [a-zA-Z0-9-\.]*.
	Team *string `json:"Team,omitempty" name:"Team"`

	// Custom player status. This parameter will be passed through. Value range: [0, 99999]
	CustomPlayerStatus *uint64 `json:"CustomPlayerStatus,omitempty" name:"CustomPlayerStatus"`

	// Custom player information. Up to 1024 characters are allowed. This parameter will be passed through.
	CustomProfile *string `json:"CustomProfile,omitempty" name:"CustomProfile"`

	// Number of delays in each area. Up to 20 delays are supported.
	RegionLatencies []*RegionLatency `json:"RegionLatencies,omitempty" name:"RegionLatencies"`
}

type RegionLatency struct {

	// Region
	// ap-beijing          North China (Beijing)
	// ap-chengdu          Southwest China (Chengdu)
	// ap-guangzhou           South China (Guangzhou)
	// ap-hongkong           Hong Kong/Macao/Taiwan (Hong Kong, China)
	// ap-seoul          Asia Pacific (Seoul)
	// ap-shanghai          East China (Shanghai)
	// ap-singapore          Southeast Asia (Singapore)
	// eu-frankfurt          Europe (Frankfurt)
	// na-siliconvalley          Western US (Silicon Valley)
	// na-toronto           North America (Toronto)
	// ap-mumbai           Asia Pacific (Mumbai)
	// na-ashburn          Eastern US (Virginia)
	// ap-bangkok           Asia Pacific (Bangkok)
	// eu-moscow           Europe (Moscow)
	// ap-tokyo           Asia Pacific (Tokyo)
	Region *string `json:"Region,omitempty" name:"Region"`

	// Delay time in ms. Value range: 0 - 999999
	Latency *uint64 `json:"Latency,omitempty" name:"Latency"`
}

type ReportOverviewData struct {

	// Total count
	TotalTimes *string `json:"TotalTimes,omitempty" name:"TotalTimes"`

	// Success rate
	SuccessPercent *float64 `json:"SuccessPercent,omitempty" name:"SuccessPercent"`

	// Timeout rate
	TimeoutPercent *float64 `json:"TimeoutPercent,omitempty" name:"TimeoutPercent"`

	// Failure rate
	FailPercent *float64 `json:"FailPercent,omitempty" name:"FailPercent"`

	// Average matching time
	AverageSec *float64 `json:"AverageSec,omitempty" name:"AverageSec"`
}

type ReportTrendData struct {

	// Total count
	TotalList []*string `json:"TotalList,omitempty" name:"TotalList"`

	// Number of requests cancelled
	CancelList []*string `json:"CancelList,omitempty" name:"CancelList"`

	// Number of successes
	SuccessList []*string `json:"SuccessList,omitempty" name:"SuccessList"`

	// Number of failures
	FailList []*string `json:"FailList,omitempty" name:"FailList"`

	// Number of request timeout
	TimeoutList []*string `json:"TimeoutList,omitempty" name:"TimeoutList"`

	// Time array in seconds
	TimeList []*string `json:"TimeList,omitempty" name:"TimeList"`
}

type RuleBriefInfo struct {

	// Rule name. It supports [a-zA-Z\d-\.]*.
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// The associated match
	MatchCodeList []*StringKV `json:"MatchCodeList,omitempty" name:"MatchCodeList"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// RuleCode
	RuleCode *string `json:"RuleCode,omitempty" name:"RuleCode"`
}

type RuleInfo struct {

	// Rule name. It supports [a-zA-Z0-9-\.]*.
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Rule description
	// Note: this field may return `null`, indicating that no valid value is obtained.
	RuleDesc *string `json:"RuleDesc,omitempty" name:"RuleDesc"`

	// Rule script
	RuleScript *string `json:"RuleScript,omitempty" name:"RuleScript"`

	// Tag
	// Note: this field may return `null`, indicating that no valid value is obtained.
	Tags []*StringKV `json:"Tags,omitempty" name:"Tags"`

	// The associated match
	// Note: this field may return `null`, indicating that no valid value is obtained.
	MatchCodeList []*StringKV `json:"MatchCodeList,omitempty" name:"MatchCodeList"`

	// RuleCode
	RuleCode *string `json:"RuleCode,omitempty" name:"RuleCode"`

	// Region
	// Note: this field may return `null`, indicating that no valid value is obtained.
	Region *string `json:"Region,omitempty" name:"Region"`

	// User AppId
	// Note: this field may return `null`, indicating that no valid value is obtained.
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// User UIN
	// Note: this field may return `null`, indicating that no valid value is obtained.
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// User OwnerUin
	// Note: this field may return `null`, indicating that no valid value is obtained.
	CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`
}

type StartMatchingBackfillRequest struct {
	*tchttp.BaseRequest

	// MatchCode
	MatchCode *string `json:"MatchCode,omitempty" name:"MatchCode"`

	// Player information
	Players []*Player `json:"Players,omitempty" name:"Players"`

	// Game server session ID. It should contain 1 to 256 ASCII characters.
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// MatchTicket ID, which can contain 1 to 128 characters. This parameter is left empty by default, and in this case, MatchTicket ID will be automatically generated by GPM.
	MatchTicketId *string `json:"MatchTicketId,omitempty" name:"MatchTicketId"`
}

func (r *StartMatchingBackfillRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartMatchingBackfillRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MatchCode")
	delete(f, "Players")
	delete(f, "GameServerSessionId")
	delete(f, "MatchTicketId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartMatchingBackfillRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type StartMatchingBackfillResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// MatchTicket
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		MatchTicket *MatchTicket `json:"MatchTicket,omitempty" name:"MatchTicket"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartMatchingBackfillResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartMatchingBackfillResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartMatchingRequest struct {
	*tchttp.BaseRequest

	// MatchCode
	MatchCode *string `json:"MatchCode,omitempty" name:"MatchCode"`

	// Player information. Up to 200 entries can be entered.
	Players []*Player `json:"Players,omitempty" name:"Players"`

	// MatchTicket ID, which can contain up to 128 characters and can only contain numbers, letters, “.”, and “-”. This parameter is left empty by default. When it is empty, the MatchTicket ID will be automatically generated by GPM.
	MatchTicketId *string `json:"MatchTicketId,omitempty" name:"MatchTicketId"`
}

func (r *StartMatchingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartMatchingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MatchCode")
	delete(f, "Players")
	delete(f, "MatchTicketId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartMatchingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type StartMatchingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Error code
		ErrCode *uint64 `json:"ErrCode,omitempty" name:"ErrCode"`

		// MatchTicket ID. Up to 128 characters are allowed.
		MatchTicketId *string `json:"MatchTicketId,omitempty" name:"MatchTicketId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartMatchingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartMatchingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StringKV struct {

	// Key
	Key *string `json:"Key,omitempty" name:"Key"`

	// Value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type Tag struct {

	// Tag key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}
