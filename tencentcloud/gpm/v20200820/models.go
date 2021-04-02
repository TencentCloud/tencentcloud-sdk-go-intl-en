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

func (r *CancelMatchingRequest) FromJsonString(s string) error {
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

func (r *CancelMatchingResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMatchingProgressRequest struct {
	*tchttp.BaseRequest

	// List of MatchTicket IDs. It can contain up to 12 IDs.
	MatchTicketIds []*MTicket `json:"MatchTicketIds,omitempty" name:"MatchTicketIds" list`
}

func (r *DescribeMatchingProgressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMatchingProgressRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMatchingProgressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// MatchTicket list
	// Note: this field may return `null`, indicating that no valid value is obtained.
		MatchTickets []*MatchTicket `json:"MatchTickets,omitempty" name:"MatchTickets" list`

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

func (r *DescribeMatchingProgressResponse) FromJsonString(s string) error {
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

func (r *DescribeTokenRequest) FromJsonString(s string) error {
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
	ListValue []*string `json:"ListValue,omitempty" name:"ListValue" list`

	// Map attribute value
	MapValue []*AttributeMap `json:"MapValue,omitempty" name:"MapValue" list`
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
	Players []*Player `json:"Players,omitempty" name:"Players" list`

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

func (r *ModifyTokenRequest) FromJsonString(s string) error {
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

func (r *ModifyTokenResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Player struct {

	// Player ID. It can contain up to 128 characters, supporting [a-zA-Z\d-\._]*.
	Id *string `json:"Id,omitempty" name:"Id"`

	// Player nickname. Up to 128 characters are allowed.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Player attribute for matching. Up to 10 attributes are supported.
	MatchAttributes []*MatchAttribute `json:"MatchAttributes,omitempty" name:"MatchAttributes" list`

	// Team name. A player can pass in a different team name, which can contain up to 128 characters, and support [a-zA-Z0-9-\.]*.
	Team *string `json:"Team,omitempty" name:"Team"`

	// Custom player status. This parameter will be passed through. Value range: [0, 99999]
	CustomPlayerStatus *uint64 `json:"CustomPlayerStatus,omitempty" name:"CustomPlayerStatus"`

	// Custom player information. Up to 1024 characters are allowed. This parameter will be passed through.
	CustomProfile *string `json:"CustomProfile,omitempty" name:"CustomProfile"`

	// Number of delays in each area. Up to 20 delays are supported.
	RegionLatencies []*RegionLatency `json:"RegionLatencies,omitempty" name:"RegionLatencies" list`
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

type StartMatchingRequest struct {
	*tchttp.BaseRequest

	// MatchCode
	MatchCode *string `json:"MatchCode,omitempty" name:"MatchCode"`

	// Player information. Up to 200 entries can be entered.
	Players []*Player `json:"Players,omitempty" name:"Players" list`

	// MatchTicket ID, which can contain up to 128 characters and can only contain numbers, letters, “.”, and “-”. This parameter is left empty by default. When it is empty, the MatchTicket ID will be automatically generated by GPM.
	MatchTicketId *string `json:"MatchTicketId,omitempty" name:"MatchTicketId"`
}

func (r *StartMatchingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartMatchingRequest) FromJsonString(s string) error {
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

func (r *StartMatchingResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
