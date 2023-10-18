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
    "context"
    "errors"
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
    return c.CopyFleetWithContext(context.Background(), request)
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
func (c *Client) CopyFleetWithContext(ctx context.Context, request *CopyFleetRequest) (response *CopyFleetResponse, err error) {
    if request == nil {
        request = NewCopyFleetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CopyFleet require credential")
    }

    request.SetContext(ctx)
    
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
    return c.CreateGameServerSessionWithContext(context.Background(), request)
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
func (c *Client) CreateGameServerSessionWithContext(ctx context.Context, request *CreateGameServerSessionRequest) (response *CreateGameServerSessionResponse, err error) {
    if request == nil {
        request = NewCreateGameServerSessionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateGameServerSession require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DeleteTimerScalingPolicyWithContext(context.Background(), request)
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
func (c *Client) DeleteTimerScalingPolicyWithContext(ctx context.Context, request *DeleteTimerScalingPolicyRequest) (response *DeleteTimerScalingPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteTimerScalingPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTimerScalingPolicy require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeGameServerSessionDetailsWithContext(context.Background(), request)
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
func (c *Client) DescribeGameServerSessionDetailsWithContext(ctx context.Context, request *DescribeGameServerSessionDetailsRequest) (response *DescribeGameServerSessionDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeGameServerSessionDetailsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGameServerSessionDetails require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeGameServerSessionPlacementWithContext(context.Background(), request)
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
func (c *Client) DescribeGameServerSessionPlacementWithContext(ctx context.Context, request *DescribeGameServerSessionPlacementRequest) (response *DescribeGameServerSessionPlacementResponse, err error) {
    if request == nil {
        request = NewDescribeGameServerSessionPlacementRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGameServerSessionPlacement require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeGameServerSessionsWithContext(context.Background(), request)
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
func (c *Client) DescribeGameServerSessionsWithContext(ctx context.Context, request *DescribeGameServerSessionsRequest) (response *DescribeGameServerSessionsResponse, err error) {
    if request == nil {
        request = NewDescribeGameServerSessionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGameServerSessions require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeInstanceTypesWithContext(context.Background(), request)
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
func (c *Client) DescribeInstanceTypesWithContext(ctx context.Context, request *DescribeInstanceTypesRequest) (response *DescribeInstanceTypesResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceTypesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceTypes require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribePlayerSessionsWithContext(context.Background(), request)
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
func (c *Client) DescribePlayerSessionsWithContext(ctx context.Context, request *DescribePlayerSessionsRequest) (response *DescribePlayerSessionsResponse, err error) {
    if request == nil {
        request = NewDescribePlayerSessionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePlayerSessions require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeTimerScalingPoliciesWithContext(context.Background(), request)
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
func (c *Client) DescribeTimerScalingPoliciesWithContext(ctx context.Context, request *DescribeTimerScalingPoliciesRequest) (response *DescribeTimerScalingPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeTimerScalingPoliciesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTimerScalingPolicies require credential")
    }

    request.SetContext(ctx)
    
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
    return c.EndGameServerSessionAndProcessWithContext(context.Background(), request)
}

// EndGameServerSessionAndProcess
// This API is used to terminate the game server session and the corresponding process, which is applicable to time-limited protection and no protection.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) EndGameServerSessionAndProcessWithContext(ctx context.Context, request *EndGameServerSessionAndProcessRequest) (response *EndGameServerSessionAndProcessResponse, err error) {
    if request == nil {
        request = NewEndGameServerSessionAndProcessRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EndGameServerSessionAndProcess require credential")
    }

    request.SetContext(ctx)
    
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
    return c.GetGameServerSessionLogUrlWithContext(context.Background(), request)
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
func (c *Client) GetGameServerSessionLogUrlWithContext(ctx context.Context, request *GetGameServerSessionLogUrlRequest) (response *GetGameServerSessionLogUrlResponse, err error) {
    if request == nil {
        request = NewGetGameServerSessionLogUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetGameServerSessionLogUrl require credential")
    }

    request.SetContext(ctx)
    
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
    return c.GetInstanceAccessWithContext(context.Background(), request)
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
func (c *Client) GetInstanceAccessWithContext(ctx context.Context, request *GetInstanceAccessRequest) (response *GetInstanceAccessResponse, err error) {
    if request == nil {
        request = NewGetInstanceAccessRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetInstanceAccess require credential")
    }

    request.SetContext(ctx)
    
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
    return c.JoinGameServerSessionWithContext(context.Background(), request)
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
func (c *Client) JoinGameServerSessionWithContext(ctx context.Context, request *JoinGameServerSessionRequest) (response *JoinGameServerSessionResponse, err error) {
    if request == nil {
        request = NewJoinGameServerSessionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("JoinGameServerSession require credential")
    }

    request.SetContext(ctx)
    
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
    return c.JoinGameServerSessionBatchWithContext(context.Background(), request)
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
func (c *Client) JoinGameServerSessionBatchWithContext(ctx context.Context, request *JoinGameServerSessionBatchRequest) (response *JoinGameServerSessionBatchResponse, err error) {
    if request == nil {
        request = NewJoinGameServerSessionBatchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("JoinGameServerSessionBatch require credential")
    }

    request.SetContext(ctx)
    
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
    return c.PutTimerScalingPolicyWithContext(context.Background(), request)
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
func (c *Client) PutTimerScalingPolicyWithContext(ctx context.Context, request *PutTimerScalingPolicyRequest) (response *PutTimerScalingPolicyResponse, err error) {
    if request == nil {
        request = NewPutTimerScalingPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PutTimerScalingPolicy require credential")
    }

    request.SetContext(ctx)
    
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
    return c.SearchGameServerSessionsWithContext(context.Background(), request)
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
func (c *Client) SearchGameServerSessionsWithContext(ctx context.Context, request *SearchGameServerSessionsRequest) (response *SearchGameServerSessionsResponse, err error) {
    if request == nil {
        request = NewSearchGameServerSessionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchGameServerSessions require credential")
    }

    request.SetContext(ctx)
    
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
    return c.SetServerReservedWithContext(context.Background(), request)
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
func (c *Client) SetServerReservedWithContext(ctx context.Context, request *SetServerReservedRequest) (response *SetServerReservedResponse, err error) {
    if request == nil {
        request = NewSetServerReservedRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetServerReserved require credential")
    }

    request.SetContext(ctx)
    
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
    return c.StartGameServerSessionPlacementWithContext(context.Background(), request)
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
func (c *Client) StartGameServerSessionPlacementWithContext(ctx context.Context, request *StartGameServerSessionPlacementRequest) (response *StartGameServerSessionPlacementResponse, err error) {
    if request == nil {
        request = NewStartGameServerSessionPlacementRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartGameServerSessionPlacement require credential")
    }

    request.SetContext(ctx)
    
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
    return c.StopGameServerSessionPlacementWithContext(context.Background(), request)
}

// StopGameServerSessionPlacement
// This API is used to stop placing a game server session.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) StopGameServerSessionPlacementWithContext(ctx context.Context, request *StopGameServerSessionPlacementRequest) (response *StopGameServerSessionPlacementResponse, err error) {
    if request == nil {
        request = NewStopGameServerSessionPlacementRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopGameServerSessionPlacement require credential")
    }

    request.SetContext(ctx)
    
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
    return c.UpdateBucketAccelerateOptWithContext(context.Background(), request)
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
func (c *Client) UpdateBucketAccelerateOptWithContext(ctx context.Context, request *UpdateBucketAccelerateOptRequest) (response *UpdateBucketAccelerateOptResponse, err error) {
    if request == nil {
        request = NewUpdateBucketAccelerateOptRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateBucketAccelerateOpt require credential")
    }

    request.SetContext(ctx)
    
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
    return c.UpdateBucketCORSOptWithContext(context.Background(), request)
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
func (c *Client) UpdateBucketCORSOptWithContext(ctx context.Context, request *UpdateBucketCORSOptRequest) (response *UpdateBucketCORSOptResponse, err error) {
    if request == nil {
        request = NewUpdateBucketCORSOptRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateBucketCORSOpt require credential")
    }

    request.SetContext(ctx)
    
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
    return c.UpdateGameServerSessionWithContext(context.Background(), request)
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
func (c *Client) UpdateGameServerSessionWithContext(ctx context.Context, request *UpdateGameServerSessionRequest) (response *UpdateGameServerSessionResponse, err error) {
    if request == nil {
        request = NewUpdateGameServerSessionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateGameServerSession require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateGameServerSessionResponse()
    err = c.Send(request, response)
    return
}
