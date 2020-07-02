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

package v20191112

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type CreateGameServerSessionRequest struct {
	*tchttp.BaseRequest

	// Maximum number of players
	MaximumPlayerSessionCount *uint64 `json:"MaximumPlayerSessionCount,omitempty" name:"MaximumPlayerSessionCount"`

	// Alias ID
	AliasId *string `json:"AliasId,omitempty" name:"AliasId"`

	// Creator ID
	CreatorId *string `json:"CreatorId,omitempty" name:"CreatorId"`

	// Fleet ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// Game attributes
	GameProperties []*GameProperty `json:"GameProperties,omitempty" name:"GameProperties" list`

	// Game server session attribute details
	GameServerSessionData *string `json:"GameServerSessionData,omitempty" name:"GameServerSessionData"`

	// Custom ID of game server session
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// Idempotency token
	IdempotencyToken *string `json:"IdempotencyToken,omitempty" name:"IdempotencyToken"`

	// Game server session name
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *CreateGameServerSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateGameServerSessionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateGameServerSessionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Game server session
	// Note: this field may return null, indicating that no valid values can be obtained.
		GameServerSession *GameServerSession `json:"GameServerSession,omitempty" name:"GameServerSession"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateGameServerSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateGameServerSessionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Credentials struct {

	// SSH private key
	Secret *string `json:"Secret,omitempty" name:"Secret"`

	// Username
	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

type DescribeGameServerSessionDetailsRequest struct {
	*tchttp.BaseRequest

	// Alias ID
	AliasId *string `json:"AliasId,omitempty" name:"AliasId"`

	// Fleet ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// Game server session ID
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// Maximum number of entries in a single query
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Pagination offset, which is used for querying the next page
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// Game server session status. Valid values: ACTIVE, ACTIVATING, TERMINATED, TERMINATING, ERROR
	StatusFilter *string `json:"StatusFilter,omitempty" name:"StatusFilter"`
}

func (r *DescribeGameServerSessionDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGameServerSessionDetailsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGameServerSessionDetailsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of game server session details
	// Note: this field may return null, indicating that no valid values can be obtained.
		GameServerSessionDetails []*GameServerSessionDetail `json:"GameServerSessionDetails,omitempty" name:"GameServerSessionDetails" list`

		// Pagination offset, which is used for querying the next page
	// Note: this field may return null, indicating that no valid values can be obtained.
		NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGameServerSessionDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGameServerSessionDetailsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGameServerSessionPlacementRequest struct {
	*tchttp.BaseRequest

	// Unique ID of game server session placement
	PlacementId *string `json:"PlacementId,omitempty" name:"PlacementId"`
}

func (r *DescribeGameServerSessionPlacementRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGameServerSessionPlacementRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGameServerSessionPlacementResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Game server session placement
		GameServerSessionPlacement *GameServerSessionPlacement `json:"GameServerSessionPlacement,omitempty" name:"GameServerSessionPlacement"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGameServerSessionPlacementResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGameServerSessionPlacementResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGameServerSessionsRequest struct {
	*tchttp.BaseRequest

	// Alias ID
	AliasId *string `json:"AliasId,omitempty" name:"AliasId"`

	// Fleet ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// Game server session ID
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// Maximum number of entries in a single query
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Pagination offset, which is used for querying the next page
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// Game server session status. Valid values: ACTIVE, ACTIVATING, TERMINATED, TERMINATING, ERROR
	StatusFilter *string `json:"StatusFilter,omitempty" name:"StatusFilter"`
}

func (r *DescribeGameServerSessionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGameServerSessionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGameServerSessionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Game server session list
	// Note: this field may return null, indicating that no valid values can be obtained.
		GameServerSessions []*GameServerSession `json:"GameServerSessions,omitempty" name:"GameServerSessions" list`

		// Pagination offset, which is used for querying the next page
	// Note: this field may return null, indicating that no valid values can be obtained.
		NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGameServerSessionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGameServerSessionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePlayerSessionsRequest struct {
	*tchttp.BaseRequest

	// Game server session ID
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// Maximum number of entries in a single query
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Pagination offset, which is used for querying the next page
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// Player ID
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`

	// Player session ID
	PlayerSessionId *string `json:"PlayerSessionId,omitempty" name:"PlayerSessionId"`

	// Player session status. Valid values: RESERVED, ACTIVE, COMPLETED, TIMEDOUT
	PlayerSessionStatusFilter *string `json:"PlayerSessionStatusFilter,omitempty" name:"PlayerSessionStatusFilter"`
}

func (r *DescribePlayerSessionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePlayerSessionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePlayerSessionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Player session list
	// Note: this field may return null, indicating that no valid values can be obtained.
		PlayerSessions []*PlayerSession `json:"PlayerSessions,omitempty" name:"PlayerSessions" list`

		// Pagination offset
	// Note: this field may return null, indicating that no valid values can be obtained.
		NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePlayerSessionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePlayerSessionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DesiredPlayerSession struct {

	// Unique player ID associated with player session
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`

	// Developer-defined player data
	PlayerData *string `json:"PlayerData,omitempty" name:"PlayerData"`
}

type GameProperty struct {

	// Attribute name
	Key *string `json:"Key,omitempty" name:"Key"`

	// Attribute value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type GameServerSession struct {

	// Game server session creation time
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// Creator ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreatorId *string `json:"CreatorId,omitempty" name:"CreatorId"`

	// Current number of players
	CurrentPlayerSessionCount *uint64 `json:"CurrentPlayerSessionCount,omitempty" name:"CurrentPlayerSessionCount"`

	// CVM DNS ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	DnsName *string `json:"DnsName,omitempty" name:"DnsName"`

	// Fleet ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// Game attributes
	// Note: this field may return null, indicating that no valid values can be obtained.
	GameProperties []*GameProperty `json:"GameProperties,omitempty" name:"GameProperties" list`

	// Game server session attribute details
	// Note: this field may return null, indicating that no valid values can be obtained.
	GameServerSessionData *string `json:"GameServerSessionData,omitempty" name:"GameServerSessionData"`

	// Game server session ID
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// CVM IP address
	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`

	// Battle progress details
	// Note: this field may return null, indicating that no valid values can be obtained.
	MatchmakerData *string `json:"MatchmakerData,omitempty" name:"MatchmakerData"`

	// Maximum number of players
	MaximumPlayerSessionCount *uint64 `json:"MaximumPlayerSessionCount,omitempty" name:"MaximumPlayerSessionCount"`

	// Game server session name
	// Note: this field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Player session creation policy
	// Note: this field may return null, indicating that no valid values can be obtained.
	PlayerSessionCreationPolicy *string `json:"PlayerSessionCreationPolicy,omitempty" name:"PlayerSessionCreationPolicy"`

	// Port number
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// Game server session status
	Status *string `json:"Status,omitempty" name:"Status"`

	// Additional information of game server session status
	// Note: this field may return null, indicating that no valid values can be obtained.
	StatusReason *string `json:"StatusReason,omitempty" name:"StatusReason"`

	// Termination time
	// Note: this field may return null, indicating that no valid values can be obtained.
	TerminationTime *string `json:"TerminationTime,omitempty" name:"TerminationTime"`

	// Instance type
	// Note: this field may return null, indicating that no valid values can be obtained.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Current custom count
	// Note: this field may return null, indicating that no valid values can be obtained.
	CurrentCustomCount *int64 `json:"CurrentCustomCount,omitempty" name:"CurrentCustomCount"`

	// Maximum custom count
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxCustomCount *int64 `json:"MaxCustomCount,omitempty" name:"MaxCustomCount"`

	// Weight
	// Note: this field may return null, indicating that no valid values can be obtained.
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// Session availability status, i.e., whether it is blocked
	// Note: this field may return null, indicating that no valid values can be obtained.
	AvailabilityStatus *string `json:"AvailabilityStatus,omitempty" name:"AvailabilityStatus"`
}

type GameServerSessionDetail struct {

	// Game server session
	GameServerSession *GameServerSession `json:"GameServerSession,omitempty" name:"GameServerSession"`

	// Protection policy. Valid values: NoProtection, FullProtection
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProtectionPolicy *string `json:"ProtectionPolicy,omitempty" name:"ProtectionPolicy"`
}

type GameServerSessionPlacement struct {

	// Deployment ID
	PlacementId *string `json:"PlacementId,omitempty" name:"PlacementId"`

	// Service deployment group name
	GameServerSessionQueueName *string `json:"GameServerSessionQueueName,omitempty" name:"GameServerSessionQueueName"`

	// Player latency
	// Note: this field may return null, indicating that no valid values can be obtained.
	PlayerLatencies []*PlayerLatency `json:"PlayerLatencies,omitempty" name:"PlayerLatencies" list`

	// Service deployment status
	Status *string `json:"Status,omitempty" name:"Status"`

	// DNS ID assigned to the instance where the game session is running
	// Note: this field may return null, indicating that no valid values can be obtained.
	DnsName *string `json:"DnsName,omitempty" name:"DnsName"`

	// Game session ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// Game session name
	// Note: this field may return null, indicating that no valid values can be obtained.
	GameServerSessionName *string `json:"GameServerSessionName,omitempty" name:"GameServerSessionName"`

	// Service deployment region
	// Note: this field may return null, indicating that no valid values can be obtained.
	GameServerSessionRegion *string `json:"GameServerSessionRegion,omitempty" name:"GameServerSessionRegion"`

	// Game attributes
	// Note: this field may return null, indicating that no valid values can be obtained.
	GameProperties []*GameProperty `json:"GameProperties,omitempty" name:"GameProperties" list`

	// Maximum number of players
	MaximumPlayerSessionCount *uint64 `json:"MaximumPlayerSessionCount,omitempty" name:"MaximumPlayerSessionCount"`

	// Game session data
	// Note: this field may return null, indicating that no valid values can be obtained.
	GameServerSessionData *string `json:"GameServerSessionData,omitempty" name:"GameServerSessionData"`

	// IP address of the instance where the game session is running
	// Note: this field may return null, indicating that no valid values can be obtained.
	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`

	// Port number of the instance where the game session is running
	// Note: this field may return null, indicating that no valid values can be obtained.
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// Game match data
	// Note: this field may return null, indicating that no valid values can be obtained.
	MatchmakerData *string `json:"MatchmakerData,omitempty" name:"MatchmakerData"`

	// Deployed player game data
	// Note: this field may return null, indicating that no valid values can be obtained.
	PlacedPlayerSessions []*PlacedPlayerSession `json:"PlacedPlayerSessions,omitempty" name:"PlacedPlayerSessions" list`

	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	// Note: this field may return null, indicating that no valid values can be obtained.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type GetGameServerSessionLogUrlRequest struct {
	*tchttp.BaseRequest

	// Game server session ID
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`
}

func (r *GetGameServerSessionLogUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetGameServerSessionLogUrlRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetGameServerSessionLogUrlResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Log download URL
	// Note: this field may return null, indicating that no valid values can be obtained.
		PreSignedUrl *string `json:"PreSignedUrl,omitempty" name:"PreSignedUrl"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetGameServerSessionLogUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetGameServerSessionLogUrlResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetInstanceAccessRequest struct {
	*tchttp.BaseRequest

	// Service deployment ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *GetInstanceAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetInstanceAccessRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetInstanceAccessResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Credentials required for instance login
		InstanceAccess *InstanceAccess `json:"InstanceAccess,omitempty" name:"InstanceAccess"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetInstanceAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetInstanceAccessResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InstanceAccess struct {

	// Credentials required for instance access
	Credentials *Credentials `json:"Credentials,omitempty" name:"Credentials"`

	// Service deployment ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Public IP of instance
	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`

	// OS
	OperatingSystem *string `json:"OperatingSystem,omitempty" name:"OperatingSystem"`
}

type JoinGameServerSessionRequest struct {
	*tchttp.BaseRequest

	// Game server session ID
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// Player ID
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`

	// Custom player information
	PlayerData *string `json:"PlayerData,omitempty" name:"PlayerData"`
}

func (r *JoinGameServerSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *JoinGameServerSessionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type JoinGameServerSessionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Player session
	// Note: this field may return null, indicating that no valid values can be obtained.
		PlayerSession *PlayerSession `json:"PlayerSession,omitempty" name:"PlayerSession"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *JoinGameServerSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *JoinGameServerSessionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PlacedPlayerSession struct {

	// Player ID
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`

	// Player session ID
	PlayerSessionId *string `json:"PlayerSessionId,omitempty" name:"PlayerSessionId"`
}

type PlayerLatency struct {

	// Player ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`

	// Name of region corresponding to latency
	// Note: this field may return null, indicating that no valid values can be obtained.
	RegionIdentifier *string `json:"RegionIdentifier,omitempty" name:"RegionIdentifier"`

	// Latency in milliseconds
	LatencyInMilliseconds *uint64 `json:"LatencyInMilliseconds,omitempty" name:"LatencyInMilliseconds"`
}

type PlayerSession struct {

	// Player session creation time
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// ID of the DNS where the game server session is running
	// Note: this field may return null, indicating that no valid values can be obtained.
	DnsName *string `json:"DnsName,omitempty" name:"DnsName"`

	// Fleet ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// Game server session ID
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// Address of the CVM instance where the game server session is running
	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`

	// Player information
	// Note: this field may return null, indicating that no valid values can be obtained.
	PlayerData *string `json:"PlayerData,omitempty" name:"PlayerData"`

	// Player ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`

	// Player session ID
	PlayerSessionId *string `json:"PlayerSessionId,omitempty" name:"PlayerSessionId"`

	// Port number
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// Player session status
	Status *string `json:"Status,omitempty" name:"Status"`

	// Player session termination time
	// Note: this field may return null, indicating that no valid values can be obtained.
	TerminationTime *string `json:"TerminationTime,omitempty" name:"TerminationTime"`
}

type SearchGameServerSessionsRequest struct {
	*tchttp.BaseRequest

	// Alias ID
	AliasId *string `json:"AliasId,omitempty" name:"AliasId"`

	// Fleet ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// Maximum number of entries in a single query
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Pagination offset, which is used for querying the next page
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// Search filter expression. Valid values:
	// gameServerSessionName: game session name in `String` type
	// gameServerSessionId: game session ID in `String` type
	// maximumSessions: maximum number of player sessions in `Number` type
	// creationTimeMillis: creation time in milliseconds in `Number` type
	// playerSessionCount: current number of player sessions in `Number` type
	// hasAvailablePlayerSessions: whether there is available player session in `String` type. Valid values: true, false
	// gameServerSessionProperties: game session attributes in `String` type
	// 
	// Expressions in `String` type support = and <> for judgment
	// Expressions in `Number` type support =, <>, >, >=, <, and <= for judgment
	FilterExpression *string `json:"FilterExpression,omitempty" name:"FilterExpression"`

	// Sorting keyword
	// Valid values:
	// gameServerSessionName: game session name in `String` type
	// gameServerSessionId: game session ID in `String` type
	// maximumSessions: maximum number of player sessions in `Number` type
	// creationTimeMillis: creation time in milliseconds in `Number` type
	// playerSessionCount: current number of player sessions in `Number` type
	SortExpression *string `json:"SortExpression,omitempty" name:"SortExpression"`
}

func (r *SearchGameServerSessionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchGameServerSessionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchGameServerSessionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Game server session list
	// Note: this field may return null, indicating that no valid values can be obtained.
		GameServerSessions []*GameServerSession `json:"GameServerSessions,omitempty" name:"GameServerSessions" list`

		// Pagination offset, which is used for querying the next page
	// Note: this field may return null, indicating that no valid values can be obtained.
		NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchGameServerSessionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchGameServerSessionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartGameServerSessionPlacementRequest struct {
	*tchttp.BaseRequest

	// Unique ID of starting game server session placement
	PlacementId *string `json:"PlacementId,omitempty" name:"PlacementId"`

	// Game server session queue name
	GameServerSessionQueueName *string `json:"GameServerSessionQueueName,omitempty" name:"GameServerSessionQueueName"`

	// Maximum number of concurrent players allowed by the game server to connect to the game session
	MaximumPlayerSessionCount *uint64 `json:"MaximumPlayerSessionCount,omitempty" name:"MaximumPlayerSessionCount"`

	// Player game session information
	DesiredPlayerSessions []*DesiredPlayerSession `json:"DesiredPlayerSessions,omitempty" name:"DesiredPlayerSessions" list`

	// Player game session attributes
	GameProperties []*GameProperty `json:"GameProperties,omitempty" name:"GameProperties" list`

	// Game server session data
	GameServerSessionData *string `json:"GameServerSessionData,omitempty" name:"GameServerSessionData"`

	// Game server session name
	GameServerSessionName *string `json:"GameServerSessionName,omitempty" name:"GameServerSessionName"`

	// Player latency
	PlayerLatencies []*PlayerLatency `json:"PlayerLatencies,omitempty" name:"PlayerLatencies" list`
}

func (r *StartGameServerSessionPlacementRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartGameServerSessionPlacementRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartGameServerSessionPlacementResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Game server session placement
		GameServerSessionPlacement *GameServerSessionPlacement `json:"GameServerSessionPlacement,omitempty" name:"GameServerSessionPlacement"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartGameServerSessionPlacementResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartGameServerSessionPlacementResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopGameServerSessionPlacementRequest struct {
	*tchttp.BaseRequest

	// Unique ID of game server session placement
	PlacementId *string `json:"PlacementId,omitempty" name:"PlacementId"`
}

func (r *StopGameServerSessionPlacementRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopGameServerSessionPlacementRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopGameServerSessionPlacementResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Game server session placement
		GameServerSessionPlacement *GameServerSessionPlacement `json:"GameServerSessionPlacement,omitempty" name:"GameServerSessionPlacement"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopGameServerSessionPlacementResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopGameServerSessionPlacementResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateGameServerSessionRequest struct {
	*tchttp.BaseRequest

	// Game server session ID
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// Maximum number of players
	MaximumPlayerSessionCount *uint64 `json:"MaximumPlayerSessionCount,omitempty" name:"MaximumPlayerSessionCount"`

	// Game server session name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Player session creation policy. Valid values: ACCEPT_ALL, DENY_ALL
	PlayerSessionCreationPolicy *string `json:"PlayerSessionCreationPolicy,omitempty" name:"PlayerSessionCreationPolicy"`

	// Protection policy. Valid values: NoProtection, TimeLimitProtection, FullProtection
	ProtectionPolicy *string `json:"ProtectionPolicy,omitempty" name:"ProtectionPolicy"`
}

func (r *UpdateGameServerSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateGameServerSessionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateGameServerSessionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Updated game session
		GameServerSession *GameServerSession `json:"GameServerSession,omitempty" name:"GameServerSession"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateGameServerSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateGameServerSessionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
