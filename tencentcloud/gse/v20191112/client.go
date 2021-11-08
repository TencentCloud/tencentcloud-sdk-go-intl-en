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
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2019-11-12"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewCopyFleetRequest() (request *CopyFleetRequest) {
    request = &CopyFleetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "CopyFleet")
    
    return
}

func NewCopyFleetResponse() (response *CopyFleetResponse) {
    response = &CopyFleetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CopyFleet
// This API is used to replicate server fleet.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_FLEETLIMITEXCEEDED = "LimitExceeded.FleetLimitExceeded"
//  LIMITEXCEEDED_INSTANCELIMITEXCEEDED = "LimitExceeded.InstanceLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CopyFleet(request *CopyFleetRequest) (response *CopyFleetResponse, err error) {
    if request == nil {
        request = NewCopyFleetRequest()
    }
    response = NewCopyFleetResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGameServerSessionRequest() (request *CreateGameServerSessionRequest) {
    request = &CreateGameServerSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "CreateGameServerSession")
    
    return
}

func NewCreateGameServerSessionResponse() (response *CreateGameServerSessionResponse) {
    response = &CreateGameServerSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateGameServerSession
// This API is used to create a game server session.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) CreateGameServerSession(request *CreateGameServerSessionRequest) (response *CreateGameServerSessionResponse, err error) {
    if request == nil {
        request = NewCreateGameServerSessionRequest()
    }
    response = NewCreateGameServerSessionResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTimerScalingPolicyRequest() (request *DeleteTimerScalingPolicyRequest) {
    request = &DeleteTimerScalingPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DeleteTimerScalingPolicy")
    
    return
}

func NewDeleteTimerScalingPolicyResponse() (response *DeleteTimerScalingPolicyResponse) {
    response = &DeleteTimerScalingPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteTimerScalingPolicy
// This API (DeleteTimerScalingPolicy) is used to delete a scheduled scaling policy of a fleet.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) DeleteTimerScalingPolicy(request *DeleteTimerScalingPolicyRequest) (response *DeleteTimerScalingPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteTimerScalingPolicyRequest()
    }
    response = NewDeleteTimerScalingPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGameServerSessionDetailsRequest() (request *DescribeGameServerSessionDetailsRequest) {
    request = &DescribeGameServerSessionDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeGameServerSessionDetails")
    
    return
}

func NewDescribeGameServerSessionDetailsResponse() (response *DescribeGameServerSessionDetailsResponse) {
    response = &DescribeGameServerSessionDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGameServerSessionDetails
// This API is used to query the list of game server session details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) DescribeGameServerSessionDetails(request *DescribeGameServerSessionDetailsRequest) (response *DescribeGameServerSessionDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeGameServerSessionDetailsRequest()
    }
    response = NewDescribeGameServerSessionDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGameServerSessionPlacementRequest() (request *DescribeGameServerSessionPlacementRequest) {
    request = &DescribeGameServerSessionPlacementRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeGameServerSessionPlacement")
    
    return
}

func NewDescribeGameServerSessionPlacementResponse() (response *DescribeGameServerSessionPlacementResponse) {
    response = &DescribeGameServerSessionPlacementResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGameServerSessionPlacement
// This API is used to query the placement of a game server session.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) DescribeGameServerSessionPlacement(request *DescribeGameServerSessionPlacementRequest) (response *DescribeGameServerSessionPlacementResponse, err error) {
    if request == nil {
        request = NewDescribeGameServerSessionPlacementRequest()
    }
    response = NewDescribeGameServerSessionPlacementResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGameServerSessionsRequest() (request *DescribeGameServerSessionsRequest) {
    request = &DescribeGameServerSessionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeGameServerSessions")
    
    return
}

func NewDescribeGameServerSessionsResponse() (response *DescribeGameServerSessionsResponse) {
    response = &DescribeGameServerSessionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGameServerSessions
// This API is used to query the list of game server sessions.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) DescribeGameServerSessions(request *DescribeGameServerSessionsRequest) (response *DescribeGameServerSessionsResponse, err error) {
    if request == nil {
        request = NewDescribeGameServerSessionsRequest()
    }
    response = NewDescribeGameServerSessionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceTypesRequest() (request *DescribeInstanceTypesRequest) {
    request = &DescribeInstanceTypesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeInstanceTypes")
    
    return
}

func NewDescribeInstanceTypesResponse() (response *DescribeInstanceTypesResponse) {
    response = &DescribeInstanceTypesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceTypes
// This API is used to obtain the list of CVM types in the specified region.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInstanceTypes(request *DescribeInstanceTypesRequest) (response *DescribeInstanceTypesResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceTypesRequest()
    }
    response = NewDescribeInstanceTypesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePlayerSessionsRequest() (request *DescribePlayerSessionsRequest) {
    request = &DescribePlayerSessionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribePlayerSessions")
    
    return
}

func NewDescribePlayerSessionsResponse() (response *DescribePlayerSessionsResponse) {
    response = &DescribePlayerSessionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePlayerSessions
// This API is used to get the list of player sessions.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) DescribePlayerSessions(request *DescribePlayerSessionsRequest) (response *DescribePlayerSessionsResponse, err error) {
    if request == nil {
        request = NewDescribePlayerSessionsRequest()
    }
    response = NewDescribePlayerSessionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTimerScalingPoliciesRequest() (request *DescribeTimerScalingPoliciesRequest) {
    request = &DescribeTimerScalingPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeTimerScalingPolicies")
    
    return
}

func NewDescribeTimerScalingPoliciesResponse() (response *DescribeTimerScalingPoliciesResponse) {
    response = &DescribeTimerScalingPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTimerScalingPolicies
// This API (DescribeTimerScalingPolicies) is used to query the scheduled scaling policies of a fleet. You can query the policies by `fleetID` or the fleet name. The returned results are paged. 
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) DescribeTimerScalingPolicies(request *DescribeTimerScalingPoliciesRequest) (response *DescribeTimerScalingPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeTimerScalingPoliciesRequest()
    }
    response = NewDescribeTimerScalingPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewEndGameServerSessionAndProcessRequest() (request *EndGameServerSessionAndProcessRequest) {
    request = &EndGameServerSessionAndProcessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "EndGameServerSessionAndProcess")
    
    return
}

func NewEndGameServerSessionAndProcessResponse() (response *EndGameServerSessionAndProcessResponse) {
    response = &EndGameServerSessionAndProcessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EndGameServerSessionAndProcess
// This API is used to terminate the game server session and the corresponding process, which is applicable to time-limited protection and no protection.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) EndGameServerSessionAndProcess(request *EndGameServerSessionAndProcessRequest) (response *EndGameServerSessionAndProcessResponse, err error) {
    if request == nil {
        request = NewEndGameServerSessionAndProcessRequest()
    }
    response = NewEndGameServerSessionAndProcessResponse()
    err = c.Send(request, response)
    return
}

func NewGetGameServerSessionLogUrlRequest() (request *GetGameServerSessionLogUrlRequest) {
    request = &GetGameServerSessionLogUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "GetGameServerSessionLogUrl")
    
    return
}

func NewGetGameServerSessionLogUrlResponse() (response *GetGameServerSessionLogUrlResponse) {
    response = &GetGameServerSessionLogUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetGameServerSessionLogUrl
// This API is used to get the log URL of a game server session.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) GetGameServerSessionLogUrl(request *GetGameServerSessionLogUrlRequest) (response *GetGameServerSessionLogUrlResponse, err error) {
    if request == nil {
        request = NewGetGameServerSessionLogUrlRequest()
    }
    response = NewGetGameServerSessionLogUrlResponse()
    err = c.Send(request, response)
    return
}

func NewGetInstanceAccessRequest() (request *GetInstanceAccessRequest) {
    request = &GetInstanceAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "GetInstanceAccess")
    
    return
}

func NewGetInstanceAccessResponse() (response *GetInstanceAccessResponse) {
    response = &GetInstanceAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetInstanceAccess
// This API is used to get the credentials required for instance login.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetInstanceAccess(request *GetInstanceAccessRequest) (response *GetInstanceAccessResponse, err error) {
    if request == nil {
        request = NewGetInstanceAccessRequest()
    }
    response = NewGetInstanceAccessResponse()
    err = c.Send(request, response)
    return
}

func NewJoinGameServerSessionRequest() (request *JoinGameServerSessionRequest) {
    request = &JoinGameServerSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "JoinGameServerSession")
    
    return
}

func NewJoinGameServerSessionResponse() (response *JoinGameServerSessionResponse) {
    response = &JoinGameServerSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// JoinGameServerSession
// This API is used to join a game server session.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) JoinGameServerSession(request *JoinGameServerSessionRequest) (response *JoinGameServerSessionResponse, err error) {
    if request == nil {
        request = NewJoinGameServerSessionRequest()
    }
    response = NewJoinGameServerSessionResponse()
    err = c.Send(request, response)
    return
}

func NewJoinGameServerSessionBatchRequest() (request *JoinGameServerSessionBatchRequest) {
    request = &JoinGameServerSessionBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "JoinGameServerSessionBatch")
    
    return
}

func NewJoinGameServerSessionBatchResponse() (response *JoinGameServerSessionBatchResponse) {
    response = &JoinGameServerSessionBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// JoinGameServerSessionBatch
// This API is used to join game server sessions in batch.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) JoinGameServerSessionBatch(request *JoinGameServerSessionBatchRequest) (response *JoinGameServerSessionBatchResponse, err error) {
    if request == nil {
        request = NewJoinGameServerSessionBatchRequest()
    }
    response = NewJoinGameServerSessionBatchResponse()
    err = c.Send(request, response)
    return
}

func NewPutTimerScalingPolicyRequest() (request *PutTimerScalingPolicyRequest) {
    request = &PutTimerScalingPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "PutTimerScalingPolicy")
    
    return
}

func NewPutTimerScalingPolicyResponse() (response *PutTimerScalingPolicyResponse) {
    response = &PutTimerScalingPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PutTimerScalingPolicy
// This API (PutTimerScalingPolicy) is used to create or update a scheduled scaling policy for a fleet.
//
// 
//
// If the field `timerID` is filled in, the specified policy will be updated, and if `timerID` is left empty, a new policy will be created.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_INSTANCELIMITEXCEEDED = "LimitExceeded.InstanceLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) PutTimerScalingPolicy(request *PutTimerScalingPolicyRequest) (response *PutTimerScalingPolicyResponse, err error) {
    if request == nil {
        request = NewPutTimerScalingPolicyRequest()
    }
    response = NewPutTimerScalingPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewSearchGameServerSessionsRequest() (request *SearchGameServerSessionsRequest) {
    request = &SearchGameServerSessionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "SearchGameServerSessions")
    
    return
}

func NewSearchGameServerSessionsResponse() (response *SearchGameServerSessionsResponse) {
    response = &SearchGameServerSessionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SearchGameServerSessions
// This API is used to search in the list of game server sessions.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SearchGameServerSessions(request *SearchGameServerSessionsRequest) (response *SearchGameServerSessionsResponse, err error) {
    if request == nil {
        request = NewSearchGameServerSessionsRequest()
    }
    response = NewSearchGameServerSessionsResponse()
    err = c.Send(request, response)
    return
}

func NewSetServerReservedRequest() (request *SetServerReservedRequest) {
    request = &SetServerReservedRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "SetServerReserved")
    
    return
}

func NewSetServerReservedResponse() (response *SetServerReservedResponse) {
    response = &SetServerReservedResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetServerReserved
// This API (SetServerReserved) is used to mark the exceptional instances as retained for troubleshooting.
//
// 
//
// `ReserveValue`: specifies whether to retain the instance. Valid values: `0` (do not retain), `1` (retain). Default value: `0`.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) SetServerReserved(request *SetServerReservedRequest) (response *SetServerReservedResponse, err error) {
    if request == nil {
        request = NewSetServerReservedRequest()
    }
    response = NewSetServerReservedResponse()
    err = c.Send(request, response)
    return
}

func NewStartGameServerSessionPlacementRequest() (request *StartGameServerSessionPlacementRequest) {
    request = &StartGameServerSessionPlacementRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "StartGameServerSessionPlacement")
    
    return
}

func NewStartGameServerSessionPlacementResponse() (response *StartGameServerSessionPlacementResponse) {
    response = &StartGameServerSessionPlacementResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartGameServerSessionPlacement
// This API is used to start placing a game server session.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) StartGameServerSessionPlacement(request *StartGameServerSessionPlacementRequest) (response *StartGameServerSessionPlacementResponse, err error) {
    if request == nil {
        request = NewStartGameServerSessionPlacementRequest()
    }
    response = NewStartGameServerSessionPlacementResponse()
    err = c.Send(request, response)
    return
}

func NewStopGameServerSessionPlacementRequest() (request *StopGameServerSessionPlacementRequest) {
    request = &StopGameServerSessionPlacementRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "StopGameServerSessionPlacement")
    
    return
}

func NewStopGameServerSessionPlacementResponse() (response *StopGameServerSessionPlacementResponse) {
    response = &StopGameServerSessionPlacementResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopGameServerSessionPlacement
// This API is used to stop placing a game server session.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) StopGameServerSessionPlacement(request *StopGameServerSessionPlacementRequest) (response *StopGameServerSessionPlacementResponse, err error) {
    if request == nil {
        request = NewStopGameServerSessionPlacementRequest()
    }
    response = NewStopGameServerSessionPlacementResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateBucketAccelerateOptRequest() (request *UpdateBucketAccelerateOptRequest) {
    request = &UpdateBucketAccelerateOptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "UpdateBucketAccelerateOpt")
    
    return
}

func NewUpdateBucketAccelerateOptResponse() (response *UpdateBucketAccelerateOptResponse) {
    response = &UpdateBucketAccelerateOptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateBucketAccelerateOpt
// This API (UpdateBucketAccelerateOpt) is used to enable COS global acceleration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) UpdateBucketAccelerateOpt(request *UpdateBucketAccelerateOptRequest) (response *UpdateBucketAccelerateOptResponse, err error) {
    if request == nil {
        request = NewUpdateBucketAccelerateOptRequest()
    }
    response = NewUpdateBucketAccelerateOptResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateBucketCORSOptRequest() (request *UpdateBucketCORSOptRequest) {
    request = &UpdateBucketCORSOptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "UpdateBucketCORSOpt")
    
    return
}

func NewUpdateBucketCORSOptResponse() (response *UpdateBucketCORSOptResponse) {
    response = &UpdateBucketCORSOptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateBucketCORSOpt
// This API (UpdateBucketCORSOpt) is used to configure CORS for COS.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) UpdateBucketCORSOpt(request *UpdateBucketCORSOptRequest) (response *UpdateBucketCORSOptResponse, err error) {
    if request == nil {
        request = NewUpdateBucketCORSOptRequest()
    }
    response = NewUpdateBucketCORSOptResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateGameServerSessionRequest() (request *UpdateGameServerSessionRequest) {
    request = &UpdateGameServerSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "UpdateGameServerSession")
    
    return
}

func NewUpdateGameServerSessionResponse() (response *UpdateGameServerSessionResponse) {
    response = &UpdateGameServerSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateGameServerSession
// This API is used to update a game server session.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) UpdateGameServerSession(request *UpdateGameServerSessionRequest) (response *UpdateGameServerSessionResponse, err error) {
    if request == nil {
        request = NewUpdateGameServerSessionRequest()
    }
    response = NewUpdateGameServerSessionResponse()
    err = c.Send(request, response)
    return
}
