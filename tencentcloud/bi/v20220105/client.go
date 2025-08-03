// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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

package v20220105

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2022-01-05"

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


func NewApplyEmbedIntervalRequest() (request *ApplyEmbedIntervalRequest) {
    request = &ApplyEmbedIntervalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "ApplyEmbedInterval")
    
    
    return
}

func NewApplyEmbedIntervalResponse() (response *ApplyEmbedIntervalResponse) {
    response = &ApplyEmbedIntervalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ApplyEmbedInterval
// This API is used to extend the available time of a token with strong authentication.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMBED = "InvalidParameter.Embed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) ApplyEmbedInterval(request *ApplyEmbedIntervalRequest) (response *ApplyEmbedIntervalResponse, err error) {
    return c.ApplyEmbedIntervalWithContext(context.Background(), request)
}

// ApplyEmbedInterval
// This API is used to extend the available time of a token with strong authentication.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMBED = "InvalidParameter.Embed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) ApplyEmbedIntervalWithContext(ctx context.Context, request *ApplyEmbedIntervalRequest) (response *ApplyEmbedIntervalResponse, err error) {
    if request == nil {
        request = NewApplyEmbedIntervalRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "ApplyEmbedInterval")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyEmbedInterval require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyEmbedIntervalResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDatasourceRequest() (request *CreateDatasourceRequest) {
    request = &CreateDatasourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bi", APIVersion, "CreateDatasource")
    
    
    return
}

func NewCreateDatasourceResponse() (response *CreateDatasourceResponse) {
    response = &CreateDatasourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDatasource
// This API is used to create a data source.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) CreateDatasource(request *CreateDatasourceRequest) (response *CreateDatasourceResponse, err error) {
    return c.CreateDatasourceWithContext(context.Background(), request)
}

// CreateDatasource
// This API is used to create a data source.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNAL = "InternalError.Internal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
func (c *Client) CreateDatasourceWithContext(ctx context.Context, request *CreateDatasourceRequest) (response *CreateDatasourceResponse, err error) {
    if request == nil {
        request = NewCreateDatasourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bi", APIVersion, "CreateDatasource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDatasource require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDatasourceResponse()
    err = c.Send(request, response)
    return
}
