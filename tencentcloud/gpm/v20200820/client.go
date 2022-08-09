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
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-08-20"

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


func NewCancelMatchingRequest() (request *CancelMatchingRequest) {
    request = &CancelMatchingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gpm", APIVersion, "CancelMatching")
    
    
    return
}

func NewCancelMatchingResponse() (response *CancelMatchingResponse) {
    response = &CancelMatchingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CancelMatching
// This API is used to cancel matching.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICEUNAVAILABLE = "FailedOperation.ServiceUnavailable"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MATCHCODENOTFOUND = "InvalidParameterValue.MatchCodeNotFound"
//  INVALIDPARAMETERVALUE_MATCHSTATUSNOTPERMITCANCEL = "InvalidParameterValue.MatchStatusNotPermitCancel"
//  INVALIDPARAMETERVALUE_MATCHTICKETIDNOTFOUND = "InvalidParameterValue.MatchTicketIdNotFound"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACTION = "UnauthorizedOperation.UnauthorizedAction"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) CancelMatching(request *CancelMatchingRequest) (response *CancelMatchingResponse, err error) {
    return c.CancelMatchingWithContext(context.Background(), request)
}

// CancelMatching
// This API is used to cancel matching.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICEUNAVAILABLE = "FailedOperation.ServiceUnavailable"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MATCHCODENOTFOUND = "InvalidParameterValue.MatchCodeNotFound"
//  INVALIDPARAMETERVALUE_MATCHSTATUSNOTPERMITCANCEL = "InvalidParameterValue.MatchStatusNotPermitCancel"
//  INVALIDPARAMETERVALUE_MATCHTICKETIDNOTFOUND = "InvalidParameterValue.MatchTicketIdNotFound"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACTION = "UnauthorizedOperation.UnauthorizedAction"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) CancelMatchingWithContext(ctx context.Context, request *CancelMatchingRequest) (response *CancelMatchingResponse, err error) {
    if request == nil {
        request = NewCancelMatchingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelMatching require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelMatchingResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMatchRequest() (request *CreateMatchRequest) {
    request = &CreateMatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gpm", APIVersion, "CreateMatch")
    
    
    return
}

func NewCreateMatchResponse() (response *CreateMatchResponse) {
    response = &CreateMatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateMatch
// This API is used to create a match.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LIMITMATCHCOUNT = "FailedOperation.LimitMatchCount"
//  FAILEDOPERATION_SERVICEUNAVAILABLE = "FailedOperation.ServiceUnavailable"
//  FAILEDOPERATION_TAGSOPEARTIONFAILED = "FailedOperation.TagsOpeartionFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_RULENOTFOUND = "InvalidParameterValue.RuleNotFound"
//  INVALIDPARAMETERVALUE_TAGSLIMITPERMISSION = "InvalidParameterValue.TagsLimitPermission"
//  INVALIDPARAMETERVALUE_VALUERANGELIMIT = "InvalidParameterValue.ValueRangeLimit"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZEDOPERATION = "UnauthorizedOperation.CAMUnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACTION = "UnauthorizedOperation.UnauthorizedAction"
//  UNAUTHORIZEDOPERATION_USERUNAUTH = "UnauthorizedOperation.UserUnAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) CreateMatch(request *CreateMatchRequest) (response *CreateMatchResponse, err error) {
    return c.CreateMatchWithContext(context.Background(), request)
}

// CreateMatch
// This API is used to create a match.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LIMITMATCHCOUNT = "FailedOperation.LimitMatchCount"
//  FAILEDOPERATION_SERVICEUNAVAILABLE = "FailedOperation.ServiceUnavailable"
//  FAILEDOPERATION_TAGSOPEARTIONFAILED = "FailedOperation.TagsOpeartionFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_RULENOTFOUND = "InvalidParameterValue.RuleNotFound"
//  INVALIDPARAMETERVALUE_TAGSLIMITPERMISSION = "InvalidParameterValue.TagsLimitPermission"
//  INVALIDPARAMETERVALUE_VALUERANGELIMIT = "InvalidParameterValue.ValueRangeLimit"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZEDOPERATION = "UnauthorizedOperation.CAMUnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACTION = "UnauthorizedOperation.UnauthorizedAction"
//  UNAUTHORIZEDOPERATION_USERUNAUTH = "UnauthorizedOperation.UserUnAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) CreateMatchWithContext(ctx context.Context, request *CreateMatchRequest) (response *CreateMatchResponse, err error) {
    if request == nil {
        request = NewCreateMatchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMatch require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMatchResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRuleRequest() (request *CreateRuleRequest) {
    request = &CreateRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gpm", APIVersion, "CreateRule")
    
    
    return
}

func NewCreateRuleResponse() (response *CreateRuleResponse) {
    response = &CreateRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRule
// This API is used to create a rule.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LIMITRULECOUNT = "FailedOperation.LimitRuleCount"
//  FAILEDOPERATION_SERVICEUNAVAILABLE = "FailedOperation.ServiceUnavailable"
//  FAILEDOPERATION_TAGSOPEARTIONFAILED = "FailedOperation.TagsOpeartionFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDRULESCRIPT = "InvalidParameterValue.InvalidRuleScript"
//  INVALIDPARAMETERVALUE_RULENAMEDUPLICATED = "InvalidParameterValue.RuleNameDuplicated"
//  INVALIDPARAMETERVALUE_TAGSLIMITPERMISSION = "InvalidParameterValue.TagsLimitPermission"
//  INVALIDPARAMETERVALUE_VALUERANGELIMIT = "InvalidParameterValue.ValueRangeLimit"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) CreateRule(request *CreateRuleRequest) (response *CreateRuleResponse, err error) {
    return c.CreateRuleWithContext(context.Background(), request)
}

// CreateRule
// This API is used to create a rule.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LIMITRULECOUNT = "FailedOperation.LimitRuleCount"
//  FAILEDOPERATION_SERVICEUNAVAILABLE = "FailedOperation.ServiceUnavailable"
//  FAILEDOPERATION_TAGSOPEARTIONFAILED = "FailedOperation.TagsOpeartionFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDRULESCRIPT = "InvalidParameterValue.InvalidRuleScript"
//  INVALIDPARAMETERVALUE_RULENAMEDUPLICATED = "InvalidParameterValue.RuleNameDuplicated"
//  INVALIDPARAMETERVALUE_TAGSLIMITPERMISSION = "InvalidParameterValue.TagsLimitPermission"
//  INVALIDPARAMETERVALUE_VALUERANGELIMIT = "InvalidParameterValue.ValueRangeLimit"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) CreateRuleWithContext(ctx context.Context, request *CreateRuleRequest) (response *CreateRuleResponse, err error) {
    if request == nil {
        request = NewCreateRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMatchRequest() (request *DeleteMatchRequest) {
    request = &DeleteMatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gpm", APIVersion, "DeleteMatch")
    
    
    return
}

func NewDeleteMatchResponse() (response *DeleteMatchResponse) {
    response = &DeleteMatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteMatch
// This API is used to delete a match.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSFAILEDOPERATION = "FailedOperation.CLSFailedOperation"
//  FAILEDOPERATION_SERVICEUNAVAILABLE = "FailedOperation.ServiceUnavailable"
//  FAILEDOPERATION_TAGSOPEARTIONFAILED = "FailedOperation.TagsOpeartionFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MATCHNOTFOUND = "InvalidParameterValue.MatchNotFound"
//  INVALIDPARAMETERVALUE_TAGSLIMITPERMISSION = "InvalidParameterValue.TagsLimitPermission"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZEDOPERATION = "UnauthorizedOperation.CAMUnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) DeleteMatch(request *DeleteMatchRequest) (response *DeleteMatchResponse, err error) {
    return c.DeleteMatchWithContext(context.Background(), request)
}

// DeleteMatch
// This API is used to delete a match.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSFAILEDOPERATION = "FailedOperation.CLSFailedOperation"
//  FAILEDOPERATION_SERVICEUNAVAILABLE = "FailedOperation.ServiceUnavailable"
//  FAILEDOPERATION_TAGSOPEARTIONFAILED = "FailedOperation.TagsOpeartionFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MATCHNOTFOUND = "InvalidParameterValue.MatchNotFound"
//  INVALIDPARAMETERVALUE_TAGSLIMITPERMISSION = "InvalidParameterValue.TagsLimitPermission"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZEDOPERATION = "UnauthorizedOperation.CAMUnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) DeleteMatchWithContext(ctx context.Context, request *DeleteMatchRequest) (response *DeleteMatchResponse, err error) {
    if request == nil {
        request = NewDeleteMatchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMatch require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMatchResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRuleRequest() (request *DeleteRuleRequest) {
    request = &DeleteRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gpm", APIVersion, "DeleteRule")
    
    
    return
}

func NewDeleteRuleResponse() (response *DeleteRuleResponse) {
    response = &DeleteRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRule
// This API is used to delete a rule.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICEUNAVAILABLE = "FailedOperation.ServiceUnavailable"
//  FAILEDOPERATION_TAGSOPEARTIONFAILED = "FailedOperation.TagsOpeartionFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_RULEMATCHEXISTENT = "InvalidParameterValue.RuleMatchExistent"
//  INVALIDPARAMETERVALUE_RULENOTFOUND = "InvalidParameterValue.RuleNotFound"
//  INVALIDPARAMETERVALUE_TAGSLIMITPERMISSION = "InvalidParameterValue.TagsLimitPermission"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) DeleteRule(request *DeleteRuleRequest) (response *DeleteRuleResponse, err error) {
    return c.DeleteRuleWithContext(context.Background(), request)
}

// DeleteRule
// This API is used to delete a rule.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICEUNAVAILABLE = "FailedOperation.ServiceUnavailable"
//  FAILEDOPERATION_TAGSOPEARTIONFAILED = "FailedOperation.TagsOpeartionFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_RULEMATCHEXISTENT = "InvalidParameterValue.RuleMatchExistent"
//  INVALIDPARAMETERVALUE_RULENOTFOUND = "InvalidParameterValue.RuleNotFound"
//  INVALIDPARAMETERVALUE_TAGSLIMITPERMISSION = "InvalidParameterValue.TagsLimitPermission"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) DeleteRuleWithContext(ctx context.Context, request *DeleteRuleRequest) (response *DeleteRuleResponse, err error) {
    if request == nil {
        request = NewDeleteRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataRequest() (request *DescribeDataRequest) {
    request = &DescribeDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gpm", APIVersion, "DescribeData")
    
    
    return
}

func NewDescribeDataResponse() (response *DescribeDataResponse) {
    response = &DescribeDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeData
// This API is used to view statistics.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICEUNAVAILABLE = "FailedOperation.ServiceUnavailable"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MATCHCODENOTFOUND = "InvalidParameterValue.MatchCodeNotFound"
//  INVALIDPARAMETERVALUE_MATCHNOTFOUND = "InvalidParameterValue.MatchNotFound"
//  INVALIDPARAMETERVALUE_VALUERANGELIMIT = "InvalidParameterValue.ValueRangeLimit"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACTION = "UnauthorizedOperation.UnauthorizedAction"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) DescribeData(request *DescribeDataRequest) (response *DescribeDataResponse, err error) {
    return c.DescribeDataWithContext(context.Background(), request)
}

// DescribeData
// This API is used to view statistics.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICEUNAVAILABLE = "FailedOperation.ServiceUnavailable"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MATCHCODENOTFOUND = "InvalidParameterValue.MatchCodeNotFound"
//  INVALIDPARAMETERVALUE_MATCHNOTFOUND = "InvalidParameterValue.MatchNotFound"
//  INVALIDPARAMETERVALUE_VALUERANGELIMIT = "InvalidParameterValue.ValueRangeLimit"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACTION = "UnauthorizedOperation.UnauthorizedAction"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) DescribeDataWithContext(ctx context.Context, request *DescribeDataRequest) (response *DescribeDataResponse, err error) {
    if request == nil {
        request = NewDescribeDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMatchRequest() (request *DescribeMatchRequest) {
    request = &DescribeMatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gpm", APIVersion, "DescribeMatch")
    
    
    return
}

func NewDescribeMatchResponse() (response *DescribeMatchResponse) {
    response = &DescribeMatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMatch
// This API is used to query the matchmaking details.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICEUNAVAILABLE = "FailedOperation.ServiceUnavailable"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MATCHNOTFOUND = "InvalidParameterValue.MatchNotFound"
//  INVALIDPARAMETERVALUE_RULENOTFOUND = "InvalidParameterValue.RuleNotFound"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) DescribeMatch(request *DescribeMatchRequest) (response *DescribeMatchResponse, err error) {
    return c.DescribeMatchWithContext(context.Background(), request)
}

// DescribeMatch
// This API is used to query the matchmaking details.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICEUNAVAILABLE = "FailedOperation.ServiceUnavailable"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MATCHNOTFOUND = "InvalidParameterValue.MatchNotFound"
//  INVALIDPARAMETERVALUE_RULENOTFOUND = "InvalidParameterValue.RuleNotFound"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) DescribeMatchWithContext(ctx context.Context, request *DescribeMatchRequest) (response *DescribeMatchResponse, err error) {
    if request == nil {
        request = NewDescribeMatchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMatch require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMatchResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMatchCodesRequest() (request *DescribeMatchCodesRequest) {
    request = &DescribeMatchCodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gpm", APIVersion, "DescribeMatchCodes")
    
    
    return
}

func NewDescribeMatchCodesResponse() (response *DescribeMatchCodesResponse) {
    response = &DescribeMatchCodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMatchCodes
// This API is used to query the created MatchCodes and paginate the query result.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICEUNAVAILABLE = "FailedOperation.ServiceUnavailable"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACTION = "UnauthorizedOperation.UnauthorizedAction"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) DescribeMatchCodes(request *DescribeMatchCodesRequest) (response *DescribeMatchCodesResponse, err error) {
    return c.DescribeMatchCodesWithContext(context.Background(), request)
}

// DescribeMatchCodes
// This API is used to query the created MatchCodes and paginate the query result.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICEUNAVAILABLE = "FailedOperation.ServiceUnavailable"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACTION = "UnauthorizedOperation.UnauthorizedAction"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) DescribeMatchCodesWithContext(ctx context.Context, request *DescribeMatchCodesRequest) (response *DescribeMatchCodesResponse, err error) {
    if request == nil {
        request = NewDescribeMatchCodesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMatchCodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMatchCodesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMatchesRequest() (request *DescribeMatchesRequest) {
    request = &DescribeMatchesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gpm", APIVersion, "DescribeMatches")
    
    
    return
}

func NewDescribeMatchesResponse() (response *DescribeMatchesResponse) {
    response = &DescribeMatchesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMatches
// This API is used to query the matchmaking list and paginate the query result.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSFAILEDOPERATION = "FailedOperation.CLSFailedOperation"
//  FAILEDOPERATION_SERVICEUNAVAILABLE = "FailedOperation.ServiceUnavailable"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_VALUERANGELIMIT = "InvalidParameterValue.ValueRangeLimit"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CLSSERVICENOTACTIVATED = "ResourceUnavailable.CLSServiceNotActivated"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZEDOPERATION = "UnauthorizedOperation.CAMUnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACTION = "UnauthorizedOperation.UnauthorizedAction"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) DescribeMatches(request *DescribeMatchesRequest) (response *DescribeMatchesResponse, err error) {
    return c.DescribeMatchesWithContext(context.Background(), request)
}

// DescribeMatches
// This API is used to query the matchmaking list and paginate the query result.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSFAILEDOPERATION = "FailedOperation.CLSFailedOperation"
//  FAILEDOPERATION_SERVICEUNAVAILABLE = "FailedOperation.ServiceUnavailable"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_VALUERANGELIMIT = "InvalidParameterValue.ValueRangeLimit"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CLSSERVICENOTACTIVATED = "ResourceUnavailable.CLSServiceNotActivated"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZEDOPERATION = "UnauthorizedOperation.CAMUnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACTION = "UnauthorizedOperation.UnauthorizedAction"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) DescribeMatchesWithContext(ctx context.Context, request *DescribeMatchesRequest) (response *DescribeMatchesResponse, err error) {
    if request == nil {
        request = NewDescribeMatchesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMatches require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMatchesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMatchingProgressRequest() (request *DescribeMatchingProgressRequest) {
    request = &DescribeMatchingProgressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gpm", APIVersion, "DescribeMatchingProgress")
    
    
    return
}

func NewDescribeMatchingProgressResponse() (response *DescribeMatchingProgressResponse) {
    response = &DescribeMatchingProgressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMatchingProgress
// This API is used to query the matching progress.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICEUNAVAILABLE = "FailedOperation.ServiceUnavailable"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MATCHCODENOTFOUND = "InvalidParameterValue.MatchCodeNotFound"
//  INVALIDPARAMETERVALUE_MATCHTICKETIDNOTFOUND = "InvalidParameterValue.MatchTicketIdNotFound"
//  INVALIDPARAMETERVALUE_MATCHTICKETLIMIT = "InvalidParameterValue.MatchTicketLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACTION = "UnauthorizedOperation.UnauthorizedAction"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeMatchingProgress(request *DescribeMatchingProgressRequest) (response *DescribeMatchingProgressResponse, err error) {
    return c.DescribeMatchingProgressWithContext(context.Background(), request)
}

// DescribeMatchingProgress
// This API is used to query the matching progress.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICEUNAVAILABLE = "FailedOperation.ServiceUnavailable"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MATCHCODENOTFOUND = "InvalidParameterValue.MatchCodeNotFound"
//  INVALIDPARAMETERVALUE_MATCHTICKETIDNOTFOUND = "InvalidParameterValue.MatchTicketIdNotFound"
//  INVALIDPARAMETERVALUE_MATCHTICKETLIMIT = "InvalidParameterValue.MatchTicketLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACTION = "UnauthorizedOperation.UnauthorizedAction"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeMatchingProgressWithContext(ctx context.Context, request *DescribeMatchingProgressRequest) (response *DescribeMatchingProgressResponse, err error) {
    if request == nil {
        request = NewDescribeMatchingProgressRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMatchingProgress require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMatchingProgressResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRuleRequest() (request *DescribeRuleRequest) {
    request = &DescribeRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gpm", APIVersion, "DescribeRule")
    
    
    return
}

func NewDescribeRuleResponse() (response *DescribeRuleResponse) {
    response = &DescribeRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRule
// This API is used to query the rule details.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICEUNAVAILABLE = "FailedOperation.ServiceUnavailable"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_RULENOTFOUND = "InvalidParameterValue.RuleNotFound"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) DescribeRule(request *DescribeRuleRequest) (response *DescribeRuleResponse, err error) {
    return c.DescribeRuleWithContext(context.Background(), request)
}

// DescribeRule
// This API is used to query the rule details.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICEUNAVAILABLE = "FailedOperation.ServiceUnavailable"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_RULENOTFOUND = "InvalidParameterValue.RuleNotFound"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) DescribeRuleWithContext(ctx context.Context, request *DescribeRuleRequest) (response *DescribeRuleResponse, err error) {
    if request == nil {
        request = NewDescribeRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRulesRequest() (request *DescribeRulesRequest) {
    request = &DescribeRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gpm", APIVersion, "DescribeRules")
    
    
    return
}

func NewDescribeRulesResponse() (response *DescribeRulesResponse) {
    response = &DescribeRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRules
// This API is used to query the rule set list and paginate the query result.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICEUNAVAILABLE = "FailedOperation.ServiceUnavailable"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_RULENOTFOUND = "InvalidParameterValue.RuleNotFound"
//  INVALIDPARAMETERVALUE_VALUERANGELIMIT = "InvalidParameterValue.ValueRangeLimit"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACTION = "UnauthorizedOperation.UnauthorizedAction"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) DescribeRules(request *DescribeRulesRequest) (response *DescribeRulesResponse, err error) {
    return c.DescribeRulesWithContext(context.Background(), request)
}

// DescribeRules
// This API is used to query the rule set list and paginate the query result.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICEUNAVAILABLE = "FailedOperation.ServiceUnavailable"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_RULENOTFOUND = "InvalidParameterValue.RuleNotFound"
//  INVALIDPARAMETERVALUE_VALUERANGELIMIT = "InvalidParameterValue.ValueRangeLimit"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACTION = "UnauthorizedOperation.UnauthorizedAction"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) DescribeRulesWithContext(ctx context.Context, request *DescribeRulesRequest) (response *DescribeRulesResponse, err error) {
    if request == nil {
        request = NewDescribeRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTokenRequest() (request *DescribeTokenRequest) {
    request = &DescribeTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gpm", APIVersion, "DescribeToken")
    
    
    return
}

func NewDescribeTokenResponse() (response *DescribeTokenResponse) {
    response = &DescribeTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeToken
// This API is used to query the token of a Matchcode, which is used for verified the pushed message.
//
// error code that may be returned:
//  FAILEDOPERATION_SERVICEUNAVAILABLE = "FailedOperation.ServiceUnavailable"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_MATCHCODENOTFOUND = "InvalidParameterValue.MatchCodeNotFound"
//  INVALIDPARAMETERVALUE_MATCHNOTFOUND = "InvalidParameterValue.MatchNotFound"
func (c *Client) DescribeToken(request *DescribeTokenRequest) (response *DescribeTokenResponse, err error) {
    return c.DescribeTokenWithContext(context.Background(), request)
}

// DescribeToken
// This API is used to query the token of a Matchcode, which is used for verified the pushed message.
//
// error code that may be returned:
//  FAILEDOPERATION_SERVICEUNAVAILABLE = "FailedOperation.ServiceUnavailable"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_MATCHCODENOTFOUND = "InvalidParameterValue.MatchCodeNotFound"
//  INVALIDPARAMETERVALUE_MATCHNOTFOUND = "InvalidParameterValue.MatchNotFound"
func (c *Client) DescribeTokenWithContext(ctx context.Context, request *DescribeTokenRequest) (response *DescribeTokenResponse, err error) {
    if request == nil {
        request = NewDescribeTokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTokenResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMatchRequest() (request *ModifyMatchRequest) {
    request = &ModifyMatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gpm", APIVersion, "ModifyMatch")
    
    
    return
}

func NewModifyMatchResponse() (response *ModifyMatchResponse) {
    response = &ModifyMatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyMatch
// This API is used to modify a match.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICEUNAVAILABLE = "FailedOperation.ServiceUnavailable"
//  FAILEDOPERATION_TAGSOPEARTIONFAILED = "FailedOperation.TagsOpeartionFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MATCHNOTFOUND = "InvalidParameterValue.MatchNotFound"
//  INVALIDPARAMETERVALUE_RULENOTFOUND = "InvalidParameterValue.RuleNotFound"
//  INVALIDPARAMETERVALUE_TAGSLIMITPERMISSION = "InvalidParameterValue.TagsLimitPermission"
//  INVALIDPARAMETERVALUE_VALUERANGELIMIT = "InvalidParameterValue.ValueRangeLimit"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZEDOPERATION = "UnauthorizedOperation.CAMUnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACTION = "UnauthorizedOperation.UnauthorizedAction"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) ModifyMatch(request *ModifyMatchRequest) (response *ModifyMatchResponse, err error) {
    return c.ModifyMatchWithContext(context.Background(), request)
}

// ModifyMatch
// This API is used to modify a match.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICEUNAVAILABLE = "FailedOperation.ServiceUnavailable"
//  FAILEDOPERATION_TAGSOPEARTIONFAILED = "FailedOperation.TagsOpeartionFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MATCHNOTFOUND = "InvalidParameterValue.MatchNotFound"
//  INVALIDPARAMETERVALUE_RULENOTFOUND = "InvalidParameterValue.RuleNotFound"
//  INVALIDPARAMETERVALUE_TAGSLIMITPERMISSION = "InvalidParameterValue.TagsLimitPermission"
//  INVALIDPARAMETERVALUE_VALUERANGELIMIT = "InvalidParameterValue.ValueRangeLimit"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZEDOPERATION = "UnauthorizedOperation.CAMUnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACTION = "UnauthorizedOperation.UnauthorizedAction"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) ModifyMatchWithContext(ctx context.Context, request *ModifyMatchRequest) (response *ModifyMatchResponse, err error) {
    if request == nil {
        request = NewModifyMatchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMatch require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMatchResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRuleRequest() (request *ModifyRuleRequest) {
    request = &ModifyRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gpm", APIVersion, "ModifyRule")
    
    
    return
}

func NewModifyRuleResponse() (response *ModifyRuleResponse) {
    response = &ModifyRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyRule
// This API is used to modify a rule (including description and tag).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICEUNAVAILABLE = "FailedOperation.ServiceUnavailable"
//  FAILEDOPERATION_TAGSOPEARTIONFAILED = "FailedOperation.TagsOpeartionFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_RULENOTFOUND = "InvalidParameterValue.RuleNotFound"
//  INVALIDPARAMETERVALUE_VALUERANGELIMIT = "InvalidParameterValue.ValueRangeLimit"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) ModifyRule(request *ModifyRuleRequest) (response *ModifyRuleResponse, err error) {
    return c.ModifyRuleWithContext(context.Background(), request)
}

// ModifyRule
// This API is used to modify a rule (including description and tag).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICEUNAVAILABLE = "FailedOperation.ServiceUnavailable"
//  FAILEDOPERATION_TAGSOPEARTIONFAILED = "FailedOperation.TagsOpeartionFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_RULENOTFOUND = "InvalidParameterValue.RuleNotFound"
//  INVALIDPARAMETERVALUE_VALUERANGELIMIT = "InvalidParameterValue.ValueRangeLimit"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) ModifyRuleWithContext(ctx context.Context, request *ModifyRuleRequest) (response *ModifyRuleResponse, err error) {
    if request == nil {
        request = NewModifyRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTokenRequest() (request *ModifyTokenRequest) {
    request = &ModifyTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gpm", APIVersion, "ModifyToken")
    
    
    return
}

func NewModifyTokenResponse() (response *ModifyTokenResponse) {
    response = &ModifyTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyToken
// This API is used to modify the token of a Matchcode.
//
// error code that may be returned:
//  FAILEDOPERATION_SERVICEUNAVAILABLE = "FailedOperation.ServiceUnavailable"
//  INVALIDPARAMETERVALUE_TOKENCOMPATIBLESPANINVALID = "InvalidParameterValue.TokenCompatibleSpanInvalid"
//  INVALIDPARAMETERVALUE_TOKENLIMIT = "InvalidParameterValue.TokenLimit"
//  LIMITEXCEEDED_TOKENUPDATEEXCEED = "LimitExceeded.TokenUpdateExceed"
func (c *Client) ModifyToken(request *ModifyTokenRequest) (response *ModifyTokenResponse, err error) {
    return c.ModifyTokenWithContext(context.Background(), request)
}

// ModifyToken
// This API is used to modify the token of a Matchcode.
//
// error code that may be returned:
//  FAILEDOPERATION_SERVICEUNAVAILABLE = "FailedOperation.ServiceUnavailable"
//  INVALIDPARAMETERVALUE_TOKENCOMPATIBLESPANINVALID = "InvalidParameterValue.TokenCompatibleSpanInvalid"
//  INVALIDPARAMETERVALUE_TOKENLIMIT = "InvalidParameterValue.TokenLimit"
//  LIMITEXCEEDED_TOKENUPDATEEXCEED = "LimitExceeded.TokenUpdateExceed"
func (c *Client) ModifyTokenWithContext(ctx context.Context, request *ModifyTokenRequest) (response *ModifyTokenResponse, err error) {
    if request == nil {
        request = NewModifyTokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTokenResponse()
    err = c.Send(request, response)
    return
}

func NewStartMatchingRequest() (request *StartMatchingRequest) {
    request = &StartMatchingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gpm", APIVersion, "StartMatching")
    
    
    return
}

func NewStartMatchingResponse() (response *StartMatchingResponse) {
    response = &StartMatchingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartMatching
// This API is used to pass in one player or multiple players to initiate match. Players within the same request will be assigned to the same game session.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FREQUENCYSAMEPLAYERLIMITED = "FailedOperation.FrequencySamePlayerLimited"
//  FAILEDOPERATION_SERVICEUNAVAILABLE = "FailedOperation.ServiceUnavailable"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDRULESCRIPT = "InvalidParameterValue.InvalidRuleScript"
//  INVALIDPARAMETERVALUE_MATCHCODENOTFOUND = "InvalidParameterValue.MatchCodeNotFound"
//  INVALIDPARAMETERVALUE_MATCHFEILDVALUELIMIT = "InvalidParameterValue.MatchFeildValueLimit"
//  INVALIDPARAMETERVALUE_MATCHINVALIDCHARACTERS = "InvalidParameterValue.MatchInvalidCharacters"
//  INVALIDPARAMETERVALUE_MATCHPLAYERSLIMIT = "InvalidParameterValue.MatchPlayersLimit"
//  INVALIDPARAMETERVALUE_MATCHPLAYERSREPEATED = "InvalidParameterValue.MatchPlayersRepeated"
//  INVALIDPARAMETERVALUE_MATCHTICKETIDREPEATED = "InvalidParameterValue.MatchTicketIdRepeated"
//  INVALIDPARAMETERVALUE_RULENOTFOUND = "InvalidParameterValue.RuleNotFound"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACTION = "UnauthorizedOperation.UnauthorizedAction"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) StartMatching(request *StartMatchingRequest) (response *StartMatchingResponse, err error) {
    return c.StartMatchingWithContext(context.Background(), request)
}

// StartMatching
// This API is used to pass in one player or multiple players to initiate match. Players within the same request will be assigned to the same game session.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FREQUENCYSAMEPLAYERLIMITED = "FailedOperation.FrequencySamePlayerLimited"
//  FAILEDOPERATION_SERVICEUNAVAILABLE = "FailedOperation.ServiceUnavailable"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDRULESCRIPT = "InvalidParameterValue.InvalidRuleScript"
//  INVALIDPARAMETERVALUE_MATCHCODENOTFOUND = "InvalidParameterValue.MatchCodeNotFound"
//  INVALIDPARAMETERVALUE_MATCHFEILDVALUELIMIT = "InvalidParameterValue.MatchFeildValueLimit"
//  INVALIDPARAMETERVALUE_MATCHINVALIDCHARACTERS = "InvalidParameterValue.MatchInvalidCharacters"
//  INVALIDPARAMETERVALUE_MATCHPLAYERSLIMIT = "InvalidParameterValue.MatchPlayersLimit"
//  INVALIDPARAMETERVALUE_MATCHPLAYERSREPEATED = "InvalidParameterValue.MatchPlayersRepeated"
//  INVALIDPARAMETERVALUE_MATCHTICKETIDREPEATED = "InvalidParameterValue.MatchTicketIdRepeated"
//  INVALIDPARAMETERVALUE_RULENOTFOUND = "InvalidParameterValue.RuleNotFound"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACTION = "UnauthorizedOperation.UnauthorizedAction"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) StartMatchingWithContext(ctx context.Context, request *StartMatchingRequest) (response *StartMatchingResponse, err error) {
    if request == nil {
        request = NewStartMatchingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartMatching require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartMatchingResponse()
    err = c.Send(request, response)
    return
}

func NewStartMatchingBackfillRequest() (request *StartMatchingBackfillRequest) {
    request = &StartMatchingBackfillRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gpm", APIVersion, "StartMatchingBackfill")
    
    
    return
}

func NewStartMatchingBackfillResponse() (response *StartMatchingBackfillResponse) {
    response = &StartMatchingBackfillResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartMatchingBackfill
// This API is used to send a match backfill request, for which a MatchTicket will be searched to start a new match.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FREQUENCYSAMEPLAYERLIMITED = "FailedOperation.FrequencySamePlayerLimited"
//  FAILEDOPERATION_SERVICEUNAVAILABLE = "FailedOperation.ServiceUnavailable"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_GAMESERVERSESSIONREPEATED = "InvalidParameterValue.GameServerSessionRepeated"
//  INVALIDPARAMETERVALUE_INVALIDRULESCRIPT = "InvalidParameterValue.InvalidRuleScript"
//  INVALIDPARAMETERVALUE_MATCHCODENOTFOUND = "InvalidParameterValue.MatchCodeNotFound"
//  INVALIDPARAMETERVALUE_MATCHFEILDVALUELIMIT = "InvalidParameterValue.MatchFeildValueLimit"
//  INVALIDPARAMETERVALUE_MATCHINVALIDCHARACTERS = "InvalidParameterValue.MatchInvalidCharacters"
//  INVALIDPARAMETERVALUE_MATCHPLAYERSLIMIT = "InvalidParameterValue.MatchPlayersLimit"
//  INVALIDPARAMETERVALUE_MATCHPLAYERSREPEATED = "InvalidParameterValue.MatchPlayersRepeated"
//  INVALIDPARAMETERVALUE_MATCHTICKETIDNOTFOUND = "InvalidParameterValue.MatchTicketIdNotFound"
//  INVALIDPARAMETERVALUE_MATCHTICKETIDREPEATED = "InvalidParameterValue.MatchTicketIdRepeated"
//  INVALIDPARAMETERVALUE_RULENOTFOUND = "InvalidParameterValue.RuleNotFound"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACTION = "UnauthorizedOperation.UnauthorizedAction"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) StartMatchingBackfill(request *StartMatchingBackfillRequest) (response *StartMatchingBackfillResponse, err error) {
    return c.StartMatchingBackfillWithContext(context.Background(), request)
}

// StartMatchingBackfill
// This API is used to send a match backfill request, for which a MatchTicket will be searched to start a new match.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FREQUENCYSAMEPLAYERLIMITED = "FailedOperation.FrequencySamePlayerLimited"
//  FAILEDOPERATION_SERVICEUNAVAILABLE = "FailedOperation.ServiceUnavailable"
//  INTERNALERROR = "InternalError"
//  INVALIDACTION = "InvalidAction"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_GAMESERVERSESSIONREPEATED = "InvalidParameterValue.GameServerSessionRepeated"
//  INVALIDPARAMETERVALUE_INVALIDRULESCRIPT = "InvalidParameterValue.InvalidRuleScript"
//  INVALIDPARAMETERVALUE_MATCHCODENOTFOUND = "InvalidParameterValue.MatchCodeNotFound"
//  INVALIDPARAMETERVALUE_MATCHFEILDVALUELIMIT = "InvalidParameterValue.MatchFeildValueLimit"
//  INVALIDPARAMETERVALUE_MATCHINVALIDCHARACTERS = "InvalidParameterValue.MatchInvalidCharacters"
//  INVALIDPARAMETERVALUE_MATCHPLAYERSLIMIT = "InvalidParameterValue.MatchPlayersLimit"
//  INVALIDPARAMETERVALUE_MATCHPLAYERSREPEATED = "InvalidParameterValue.MatchPlayersRepeated"
//  INVALIDPARAMETERVALUE_MATCHTICKETIDNOTFOUND = "InvalidParameterValue.MatchTicketIdNotFound"
//  INVALIDPARAMETERVALUE_MATCHTICKETIDREPEATED = "InvalidParameterValue.MatchTicketIdRepeated"
//  INVALIDPARAMETERVALUE_RULENOTFOUND = "InvalidParameterValue.RuleNotFound"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  NOSUCHVERSION = "NoSuchVersion"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACTION = "UnauthorizedOperation.UnauthorizedAction"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDREGION = "UnsupportedRegion"
func (c *Client) StartMatchingBackfillWithContext(ctx context.Context, request *StartMatchingBackfillRequest) (response *StartMatchingBackfillResponse, err error) {
    if request == nil {
        request = NewStartMatchingBackfillRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartMatchingBackfill require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartMatchingBackfillResponse()
    err = c.Send(request, response)
    return
}
