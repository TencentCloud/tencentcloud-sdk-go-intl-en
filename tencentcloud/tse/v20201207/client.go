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

package v20201207

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2020-12-07"

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


func NewBindAutoScalerResourceStrategyToGroupsRequest() (request *BindAutoScalerResourceStrategyToGroupsRequest) {
    request = &BindAutoScalerResourceStrategyToGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "BindAutoScalerResourceStrategyToGroups")
    
    
    return
}

func NewBindAutoScalerResourceStrategyToGroupsResponse() (response *BindAutoScalerResourceStrategyToGroupsResponse) {
    response = &BindAutoScalerResourceStrategyToGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindAutoScalerResourceStrategyToGroups
// Bind auto scaling policies to gateway groupings in batch
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) BindAutoScalerResourceStrategyToGroups(request *BindAutoScalerResourceStrategyToGroupsRequest) (response *BindAutoScalerResourceStrategyToGroupsResponse, err error) {
    return c.BindAutoScalerResourceStrategyToGroupsWithContext(context.Background(), request)
}

// BindAutoScalerResourceStrategyToGroups
// Bind auto scaling policies to gateway groupings in batch
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) BindAutoScalerResourceStrategyToGroupsWithContext(ctx context.Context, request *BindAutoScalerResourceStrategyToGroupsRequest) (response *BindAutoScalerResourceStrategyToGroupsResponse, err error) {
    if request == nil {
        request = NewBindAutoScalerResourceStrategyToGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "BindAutoScalerResourceStrategyToGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindAutoScalerResourceStrategyToGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindAutoScalerResourceStrategyToGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewCloseWafProtectionRequest() (request *CloseWafProtectionRequest) {
    request = &CloseWafProtectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "CloseWafProtection")
    
    
    return
}

func NewCloseWafProtectionResponse() (response *CloseWafProtectionResponse) {
    response = &CloseWafProtectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CloseWafProtection
// Disables WAF protection
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CloseWafProtection(request *CloseWafProtectionRequest) (response *CloseWafProtectionResponse, err error) {
    return c.CloseWafProtectionWithContext(context.Background(), request)
}

// CloseWafProtection
// Disables WAF protection
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CloseWafProtectionWithContext(ctx context.Context, request *CloseWafProtectionRequest) (response *CloseWafProtectionResponse, err error) {
    if request == nil {
        request = NewCloseWafProtectionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "CloseWafProtection")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseWafProtection require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseWafProtectionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAutoScalerResourceStrategyRequest() (request *CreateAutoScalerResourceStrategyRequest) {
    request = &CreateAutoScalerResourceStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "CreateAutoScalerResourceStrategy")
    
    
    return
}

func NewCreateAutoScalerResourceStrategyResponse() (response *CreateAutoScalerResourceStrategyResponse) {
    response = &CreateAutoScalerResourceStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAutoScalerResourceStrategy
// Create AS policy
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CREATEERROR = "InternalError.CreateError"
//  INTERNALERROR_DECODEERROR = "InternalError.DecodeError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UPDATEERROR = "InternalError.UpdateError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) CreateAutoScalerResourceStrategy(request *CreateAutoScalerResourceStrategyRequest) (response *CreateAutoScalerResourceStrategyResponse, err error) {
    return c.CreateAutoScalerResourceStrategyWithContext(context.Background(), request)
}

// CreateAutoScalerResourceStrategy
// Create AS policy
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CREATEERROR = "InternalError.CreateError"
//  INTERNALERROR_DECODEERROR = "InternalError.DecodeError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UPDATEERROR = "InternalError.UpdateError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) CreateAutoScalerResourceStrategyWithContext(ctx context.Context, request *CreateAutoScalerResourceStrategyRequest) (response *CreateAutoScalerResourceStrategyResponse, err error) {
    if request == nil {
        request = NewCreateAutoScalerResourceStrategyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "CreateAutoScalerResourceStrategy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAutoScalerResourceStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAutoScalerResourceStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCloudNativeAPIGatewayRequest() (request *CreateCloudNativeAPIGatewayRequest) {
    request = &CreateCloudNativeAPIGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "CreateCloudNativeAPIGateway")
    
    
    return
}

func NewCreateCloudNativeAPIGatewayResponse() (response *CreateCloudNativeAPIGatewayResponse) {
    response = &CreateCloudNativeAPIGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCloudNativeAPIGateway
// Create a cloud native API gateway instance
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLS = "FailedOperation.Cls"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_RESOURCE = "FailedOperation.Resource"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CREATEERROR = "InternalError.CreateError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CLSNOTACTIVATED = "UnauthorizedOperation.ClsNotActivated"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateCloudNativeAPIGateway(request *CreateCloudNativeAPIGatewayRequest) (response *CreateCloudNativeAPIGatewayResponse, err error) {
    return c.CreateCloudNativeAPIGatewayWithContext(context.Background(), request)
}

// CreateCloudNativeAPIGateway
// Create a cloud native API gateway instance
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLS = "FailedOperation.Cls"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_RESOURCE = "FailedOperation.Resource"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CREATEERROR = "InternalError.CreateError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CLSNOTACTIVATED = "UnauthorizedOperation.ClsNotActivated"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateCloudNativeAPIGatewayWithContext(ctx context.Context, request *CreateCloudNativeAPIGatewayRequest) (response *CreateCloudNativeAPIGatewayResponse, err error) {
    if request == nil {
        request = NewCreateCloudNativeAPIGatewayRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "CreateCloudNativeAPIGateway")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudNativeAPIGateway require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudNativeAPIGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCloudNativeAPIGatewayCanaryRuleRequest() (request *CreateCloudNativeAPIGatewayCanaryRuleRequest) {
    request = &CreateCloudNativeAPIGatewayCanaryRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "CreateCloudNativeAPIGatewayCanaryRule")
    
    
    return
}

func NewCreateCloudNativeAPIGatewayCanaryRuleResponse() (response *CreateCloudNativeAPIGatewayCanaryRuleResponse) {
    response = &CreateCloudNativeAPIGatewayCanaryRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCloudNativeAPIGatewayCanaryRule
// Create a grayscale rule for the cloud-native gateway
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateCloudNativeAPIGatewayCanaryRule(request *CreateCloudNativeAPIGatewayCanaryRuleRequest) (response *CreateCloudNativeAPIGatewayCanaryRuleResponse, err error) {
    return c.CreateCloudNativeAPIGatewayCanaryRuleWithContext(context.Background(), request)
}

// CreateCloudNativeAPIGatewayCanaryRule
// Create a grayscale rule for the cloud-native gateway
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateCloudNativeAPIGatewayCanaryRuleWithContext(ctx context.Context, request *CreateCloudNativeAPIGatewayCanaryRuleRequest) (response *CreateCloudNativeAPIGatewayCanaryRuleResponse, err error) {
    if request == nil {
        request = NewCreateCloudNativeAPIGatewayCanaryRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "CreateCloudNativeAPIGatewayCanaryRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudNativeAPIGatewayCanaryRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudNativeAPIGatewayCanaryRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCloudNativeAPIGatewayCertificateRequest() (request *CreateCloudNativeAPIGatewayCertificateRequest) {
    request = &CreateCloudNativeAPIGatewayCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "CreateCloudNativeAPIGatewayCertificate")
    
    
    return
}

func NewCreateCloudNativeAPIGatewayCertificateResponse() (response *CreateCloudNativeAPIGatewayCertificateResponse) {
    response = &CreateCloudNativeAPIGatewayCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCloudNativeAPIGatewayCertificate
// This API is used to create a cloud-native gateway certificate
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  LIMITEXCEEDED_LBDOMAINS = "LimitExceeded.LBDomains"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateCloudNativeAPIGatewayCertificate(request *CreateCloudNativeAPIGatewayCertificateRequest) (response *CreateCloudNativeAPIGatewayCertificateResponse, err error) {
    return c.CreateCloudNativeAPIGatewayCertificateWithContext(context.Background(), request)
}

// CreateCloudNativeAPIGatewayCertificate
// This API is used to create a cloud-native gateway certificate
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  LIMITEXCEEDED_LBDOMAINS = "LimitExceeded.LBDomains"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateCloudNativeAPIGatewayCertificateWithContext(ctx context.Context, request *CreateCloudNativeAPIGatewayCertificateRequest) (response *CreateCloudNativeAPIGatewayCertificateResponse, err error) {
    if request == nil {
        request = NewCreateCloudNativeAPIGatewayCertificateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "CreateCloudNativeAPIGatewayCertificate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudNativeAPIGatewayCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudNativeAPIGatewayCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCloudNativeAPIGatewayPublicNetworkRequest() (request *CreateCloudNativeAPIGatewayPublicNetworkRequest) {
    request = &CreateCloudNativeAPIGatewayPublicNetworkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "CreateCloudNativeAPIGatewayPublicNetwork")
    
    
    return
}

func NewCreateCloudNativeAPIGatewayPublicNetworkResponse() (response *CreateCloudNativeAPIGatewayPublicNetworkResponse) {
    response = &CreateCloudNativeAPIGatewayPublicNetworkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCloudNativeAPIGatewayPublicNetwork
// Create a public network configuration
//
// error code that may be returned:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) CreateCloudNativeAPIGatewayPublicNetwork(request *CreateCloudNativeAPIGatewayPublicNetworkRequest) (response *CreateCloudNativeAPIGatewayPublicNetworkResponse, err error) {
    return c.CreateCloudNativeAPIGatewayPublicNetworkWithContext(context.Background(), request)
}

// CreateCloudNativeAPIGatewayPublicNetwork
// Create a public network configuration
//
// error code that may be returned:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) CreateCloudNativeAPIGatewayPublicNetworkWithContext(ctx context.Context, request *CreateCloudNativeAPIGatewayPublicNetworkRequest) (response *CreateCloudNativeAPIGatewayPublicNetworkResponse, err error) {
    if request == nil {
        request = NewCreateCloudNativeAPIGatewayPublicNetworkRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "CreateCloudNativeAPIGatewayPublicNetwork")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudNativeAPIGatewayPublicNetwork require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudNativeAPIGatewayPublicNetworkResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCloudNativeAPIGatewayRouteRequest() (request *CreateCloudNativeAPIGatewayRouteRequest) {
    request = &CreateCloudNativeAPIGatewayRouteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "CreateCloudNativeAPIGatewayRoute")
    
    
    return
}

func NewCreateCloudNativeAPIGatewayRouteResponse() (response *CreateCloudNativeAPIGatewayRouteResponse) {
    response = &CreateCloudNativeAPIGatewayRouteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCloudNativeAPIGatewayRoute
// This API is used to create a cloud-native gateway route.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateCloudNativeAPIGatewayRoute(request *CreateCloudNativeAPIGatewayRouteRequest) (response *CreateCloudNativeAPIGatewayRouteResponse, err error) {
    return c.CreateCloudNativeAPIGatewayRouteWithContext(context.Background(), request)
}

// CreateCloudNativeAPIGatewayRoute
// This API is used to create a cloud-native gateway route.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateCloudNativeAPIGatewayRouteWithContext(ctx context.Context, request *CreateCloudNativeAPIGatewayRouteRequest) (response *CreateCloudNativeAPIGatewayRouteResponse, err error) {
    if request == nil {
        request = NewCreateCloudNativeAPIGatewayRouteRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "CreateCloudNativeAPIGatewayRoute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudNativeAPIGatewayRoute require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudNativeAPIGatewayRouteResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCloudNativeAPIGatewayRouteRateLimitRequest() (request *CreateCloudNativeAPIGatewayRouteRateLimitRequest) {
    request = &CreateCloudNativeAPIGatewayRouteRateLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "CreateCloudNativeAPIGatewayRouteRateLimit")
    
    
    return
}

func NewCreateCloudNativeAPIGatewayRouteRateLimitResponse() (response *CreateCloudNativeAPIGatewayRouteRateLimitResponse) {
    response = &CreateCloudNativeAPIGatewayRouteRateLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCloudNativeAPIGatewayRouteRateLimit
// This API is used to create a cloud-native gateway traffic throttling plugin.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateCloudNativeAPIGatewayRouteRateLimit(request *CreateCloudNativeAPIGatewayRouteRateLimitRequest) (response *CreateCloudNativeAPIGatewayRouteRateLimitResponse, err error) {
    return c.CreateCloudNativeAPIGatewayRouteRateLimitWithContext(context.Background(), request)
}

// CreateCloudNativeAPIGatewayRouteRateLimit
// This API is used to create a cloud-native gateway traffic throttling plugin.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateCloudNativeAPIGatewayRouteRateLimitWithContext(ctx context.Context, request *CreateCloudNativeAPIGatewayRouteRateLimitRequest) (response *CreateCloudNativeAPIGatewayRouteRateLimitResponse, err error) {
    if request == nil {
        request = NewCreateCloudNativeAPIGatewayRouteRateLimitRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "CreateCloudNativeAPIGatewayRouteRateLimit")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudNativeAPIGatewayRouteRateLimit require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudNativeAPIGatewayRouteRateLimitResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCloudNativeAPIGatewayServiceRequest() (request *CreateCloudNativeAPIGatewayServiceRequest) {
    request = &CreateCloudNativeAPIGatewayServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "CreateCloudNativeAPIGatewayService")
    
    
    return
}

func NewCreateCloudNativeAPIGatewayServiceResponse() (response *CreateCloudNativeAPIGatewayServiceResponse) {
    response = &CreateCloudNativeAPIGatewayServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCloudNativeAPIGatewayService
// Create a cloud-native gateway service
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateCloudNativeAPIGatewayService(request *CreateCloudNativeAPIGatewayServiceRequest) (response *CreateCloudNativeAPIGatewayServiceResponse, err error) {
    return c.CreateCloudNativeAPIGatewayServiceWithContext(context.Background(), request)
}

// CreateCloudNativeAPIGatewayService
// Create a cloud-native gateway service
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateCloudNativeAPIGatewayServiceWithContext(ctx context.Context, request *CreateCloudNativeAPIGatewayServiceRequest) (response *CreateCloudNativeAPIGatewayServiceResponse, err error) {
    if request == nil {
        request = NewCreateCloudNativeAPIGatewayServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "CreateCloudNativeAPIGatewayService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudNativeAPIGatewayService require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudNativeAPIGatewayServiceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCloudNativeAPIGatewayServiceRateLimitRequest() (request *CreateCloudNativeAPIGatewayServiceRateLimitRequest) {
    request = &CreateCloudNativeAPIGatewayServiceRateLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "CreateCloudNativeAPIGatewayServiceRateLimit")
    
    
    return
}

func NewCreateCloudNativeAPIGatewayServiceRateLimitResponse() (response *CreateCloudNativeAPIGatewayServiceRateLimitResponse) {
    response = &CreateCloudNativeAPIGatewayServiceRateLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCloudNativeAPIGatewayServiceRateLimit
// This API is used to create a traffic throttling plugin for a cloud-native gateway
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateCloudNativeAPIGatewayServiceRateLimit(request *CreateCloudNativeAPIGatewayServiceRateLimitRequest) (response *CreateCloudNativeAPIGatewayServiceRateLimitResponse, err error) {
    return c.CreateCloudNativeAPIGatewayServiceRateLimitWithContext(context.Background(), request)
}

// CreateCloudNativeAPIGatewayServiceRateLimit
// This API is used to create a traffic throttling plugin for a cloud-native gateway
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateCloudNativeAPIGatewayServiceRateLimitWithContext(ctx context.Context, request *CreateCloudNativeAPIGatewayServiceRateLimitRequest) (response *CreateCloudNativeAPIGatewayServiceRateLimitResponse, err error) {
    if request == nil {
        request = NewCreateCloudNativeAPIGatewayServiceRateLimitRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "CreateCloudNativeAPIGatewayServiceRateLimit")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudNativeAPIGatewayServiceRateLimit require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudNativeAPIGatewayServiceRateLimitResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGovernanceLaneGroupsRequest() (request *CreateGovernanceLaneGroupsRequest) {
    request = &CreateGovernanceLaneGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "CreateGovernanceLaneGroups")
    
    
    return
}

func NewCreateGovernanceLaneGroupsResponse() (response *CreateGovernanceLaneGroupsResponse) {
    response = &CreateGovernanceLaneGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateGovernanceLaneGroups
// Create a lane group
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateGovernanceLaneGroups(request *CreateGovernanceLaneGroupsRequest) (response *CreateGovernanceLaneGroupsResponse, err error) {
    return c.CreateGovernanceLaneGroupsWithContext(context.Background(), request)
}

// CreateGovernanceLaneGroups
// Create a lane group
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateGovernanceLaneGroupsWithContext(ctx context.Context, request *CreateGovernanceLaneGroupsRequest) (response *CreateGovernanceLaneGroupsResponse, err error) {
    if request == nil {
        request = NewCreateGovernanceLaneGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "CreateGovernanceLaneGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateGovernanceLaneGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateGovernanceLaneGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNativeGatewayServerGroupRequest() (request *CreateNativeGatewayServerGroupRequest) {
    request = &CreateNativeGatewayServerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "CreateNativeGatewayServerGroup")
    
    
    return
}

func NewCreateNativeGatewayServerGroupResponse() (response *CreateNativeGatewayServerGroupResponse) {
    response = &CreateNativeGatewayServerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateNativeGatewayServerGroup
// Create a Cloud Native Gateway Engine group
//
// error code that may be returned:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
func (c *Client) CreateNativeGatewayServerGroup(request *CreateNativeGatewayServerGroupRequest) (response *CreateNativeGatewayServerGroupResponse, err error) {
    return c.CreateNativeGatewayServerGroupWithContext(context.Background(), request)
}

// CreateNativeGatewayServerGroup
// Create a Cloud Native Gateway Engine group
//
// error code that may be returned:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
func (c *Client) CreateNativeGatewayServerGroupWithContext(ctx context.Context, request *CreateNativeGatewayServerGroupRequest) (response *CreateNativeGatewayServerGroupResponse, err error) {
    if request == nil {
        request = NewCreateNativeGatewayServerGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "CreateNativeGatewayServerGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNativeGatewayServerGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateNativeGatewayServerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNativeGatewayServiceSourceRequest() (request *CreateNativeGatewayServiceSourceRequest) {
    request = &CreateNativeGatewayServiceSourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "CreateNativeGatewayServiceSource")
    
    
    return
}

func NewCreateNativeGatewayServiceSourceResponse() (response *CreateNativeGatewayServiceSourceResponse) {
    response = &CreateNativeGatewayServiceSourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateNativeGatewayServiceSource
// Create a gateway service source
//
// error code that may be returned:
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) CreateNativeGatewayServiceSource(request *CreateNativeGatewayServiceSourceRequest) (response *CreateNativeGatewayServiceSourceResponse, err error) {
    return c.CreateNativeGatewayServiceSourceWithContext(context.Background(), request)
}

// CreateNativeGatewayServiceSource
// Create a gateway service source
//
// error code that may be returned:
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) CreateNativeGatewayServiceSourceWithContext(ctx context.Context, request *CreateNativeGatewayServiceSourceRequest) (response *CreateNativeGatewayServiceSourceResponse, err error) {
    if request == nil {
        request = NewCreateNativeGatewayServiceSourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "CreateNativeGatewayServiceSource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNativeGatewayServiceSource require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateNativeGatewayServiceSourceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOrModifyCloudNativeAPIGatewayCORSRequest() (request *CreateOrModifyCloudNativeAPIGatewayCORSRequest) {
    request = &CreateOrModifyCloudNativeAPIGatewayCORSRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "CreateOrModifyCloudNativeAPIGatewayCORS")
    
    
    return
}

func NewCreateOrModifyCloudNativeAPIGatewayCORSResponse() (response *CreateOrModifyCloudNativeAPIGatewayCORSResponse) {
    response = &CreateOrModifyCloudNativeAPIGatewayCORSResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateOrModifyCloudNativeAPIGatewayCORS
// Create or edit a cloud-native gateway cross-domain configuration
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateOrModifyCloudNativeAPIGatewayCORS(request *CreateOrModifyCloudNativeAPIGatewayCORSRequest) (response *CreateOrModifyCloudNativeAPIGatewayCORSResponse, err error) {
    return c.CreateOrModifyCloudNativeAPIGatewayCORSWithContext(context.Background(), request)
}

// CreateOrModifyCloudNativeAPIGatewayCORS
// Create or edit a cloud-native gateway cross-domain configuration
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateOrModifyCloudNativeAPIGatewayCORSWithContext(ctx context.Context, request *CreateOrModifyCloudNativeAPIGatewayCORSRequest) (response *CreateOrModifyCloudNativeAPIGatewayCORSResponse, err error) {
    if request == nil {
        request = NewCreateOrModifyCloudNativeAPIGatewayCORSRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "CreateOrModifyCloudNativeAPIGatewayCORS")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOrModifyCloudNativeAPIGatewayCORS require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOrModifyCloudNativeAPIGatewayCORSResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWafDomainsRequest() (request *CreateWafDomainsRequest) {
    request = &CreateWafDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "CreateWafDomains")
    
    
    return
}

func NewCreateWafDomainsResponse() (response *CreateWafDomainsResponse) {
    response = &CreateWafDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateWafDomains
// Create a WAF-protected domain name
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateWafDomains(request *CreateWafDomainsRequest) (response *CreateWafDomainsResponse, err error) {
    return c.CreateWafDomainsWithContext(context.Background(), request)
}

// CreateWafDomains
// Create a WAF-protected domain name
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateWafDomainsWithContext(ctx context.Context, request *CreateWafDomainsRequest) (response *CreateWafDomainsResponse, err error) {
    if request == nil {
        request = NewCreateWafDomainsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "CreateWafDomains")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWafDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWafDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAutoScalerResourceStrategyRequest() (request *DeleteAutoScalerResourceStrategyRequest) {
    request = &DeleteAutoScalerResourceStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DeleteAutoScalerResourceStrategy")
    
    
    return
}

func NewDeleteAutoScalerResourceStrategyResponse() (response *DeleteAutoScalerResourceStrategyResponse) {
    response = &DeleteAutoScalerResourceStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAutoScalerResourceStrategy
// Delete AS policy
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteAutoScalerResourceStrategy(request *DeleteAutoScalerResourceStrategyRequest) (response *DeleteAutoScalerResourceStrategyResponse, err error) {
    return c.DeleteAutoScalerResourceStrategyWithContext(context.Background(), request)
}

// DeleteAutoScalerResourceStrategy
// Delete AS policy
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteAutoScalerResourceStrategyWithContext(ctx context.Context, request *DeleteAutoScalerResourceStrategyRequest) (response *DeleteAutoScalerResourceStrategyResponse, err error) {
    if request == nil {
        request = NewDeleteAutoScalerResourceStrategyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "DeleteAutoScalerResourceStrategy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAutoScalerResourceStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAutoScalerResourceStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCloudNativeAPIGatewayRequest() (request *DeleteCloudNativeAPIGatewayRequest) {
    request = &DeleteCloudNativeAPIGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DeleteCloudNativeAPIGateway")
    
    
    return
}

func NewDeleteCloudNativeAPIGatewayResponse() (response *DeleteCloudNativeAPIGatewayResponse) {
    response = &DeleteCloudNativeAPIGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCloudNativeAPIGateway
// Delete a cloud native API gateway instance
//
// error code that may be returned:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_CLSNOTACTIVATED = "UnauthorizedOperation.ClsNotActivated"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteCloudNativeAPIGateway(request *DeleteCloudNativeAPIGatewayRequest) (response *DeleteCloudNativeAPIGatewayResponse, err error) {
    return c.DeleteCloudNativeAPIGatewayWithContext(context.Background(), request)
}

// DeleteCloudNativeAPIGateway
// Delete a cloud native API gateway instance
//
// error code that may be returned:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_CLSNOTACTIVATED = "UnauthorizedOperation.ClsNotActivated"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteCloudNativeAPIGatewayWithContext(ctx context.Context, request *DeleteCloudNativeAPIGatewayRequest) (response *DeleteCloudNativeAPIGatewayResponse, err error) {
    if request == nil {
        request = NewDeleteCloudNativeAPIGatewayRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "DeleteCloudNativeAPIGateway")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCloudNativeAPIGateway require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCloudNativeAPIGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCloudNativeAPIGatewayCORSRequest() (request *DeleteCloudNativeAPIGatewayCORSRequest) {
    request = &DeleteCloudNativeAPIGatewayCORSRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DeleteCloudNativeAPIGatewayCORS")
    
    
    return
}

func NewDeleteCloudNativeAPIGatewayCORSResponse() (response *DeleteCloudNativeAPIGatewayCORSResponse) {
    response = &DeleteCloudNativeAPIGatewayCORSResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCloudNativeAPIGatewayCORS
// This API is used to delete a cloud-native gateway cross-domain plug-in.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteCloudNativeAPIGatewayCORS(request *DeleteCloudNativeAPIGatewayCORSRequest) (response *DeleteCloudNativeAPIGatewayCORSResponse, err error) {
    return c.DeleteCloudNativeAPIGatewayCORSWithContext(context.Background(), request)
}

// DeleteCloudNativeAPIGatewayCORS
// This API is used to delete a cloud-native gateway cross-domain plug-in.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteCloudNativeAPIGatewayCORSWithContext(ctx context.Context, request *DeleteCloudNativeAPIGatewayCORSRequest) (response *DeleteCloudNativeAPIGatewayCORSResponse, err error) {
    if request == nil {
        request = NewDeleteCloudNativeAPIGatewayCORSRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "DeleteCloudNativeAPIGatewayCORS")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCloudNativeAPIGatewayCORS require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCloudNativeAPIGatewayCORSResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCloudNativeAPIGatewayCanaryRuleRequest() (request *DeleteCloudNativeAPIGatewayCanaryRuleRequest) {
    request = &DeleteCloudNativeAPIGatewayCanaryRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DeleteCloudNativeAPIGatewayCanaryRule")
    
    
    return
}

func NewDeleteCloudNativeAPIGatewayCanaryRuleResponse() (response *DeleteCloudNativeAPIGatewayCanaryRuleResponse) {
    response = &DeleteCloudNativeAPIGatewayCanaryRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCloudNativeAPIGatewayCanaryRule
// This API is used to delete the grayscale rule of the cloud-native gateway.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteCloudNativeAPIGatewayCanaryRule(request *DeleteCloudNativeAPIGatewayCanaryRuleRequest) (response *DeleteCloudNativeAPIGatewayCanaryRuleResponse, err error) {
    return c.DeleteCloudNativeAPIGatewayCanaryRuleWithContext(context.Background(), request)
}

// DeleteCloudNativeAPIGatewayCanaryRule
// This API is used to delete the grayscale rule of the cloud-native gateway.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteCloudNativeAPIGatewayCanaryRuleWithContext(ctx context.Context, request *DeleteCloudNativeAPIGatewayCanaryRuleRequest) (response *DeleteCloudNativeAPIGatewayCanaryRuleResponse, err error) {
    if request == nil {
        request = NewDeleteCloudNativeAPIGatewayCanaryRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "DeleteCloudNativeAPIGatewayCanaryRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCloudNativeAPIGatewayCanaryRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCloudNativeAPIGatewayCanaryRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCloudNativeAPIGatewayCertificateRequest() (request *DeleteCloudNativeAPIGatewayCertificateRequest) {
    request = &DeleteCloudNativeAPIGatewayCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DeleteCloudNativeAPIGatewayCertificate")
    
    
    return
}

func NewDeleteCloudNativeAPIGatewayCertificateResponse() (response *DeleteCloudNativeAPIGatewayCertificateResponse) {
    response = &DeleteCloudNativeAPIGatewayCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCloudNativeAPIGatewayCertificate
// This API is used to delete a cloud-native gateway cert.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteCloudNativeAPIGatewayCertificate(request *DeleteCloudNativeAPIGatewayCertificateRequest) (response *DeleteCloudNativeAPIGatewayCertificateResponse, err error) {
    return c.DeleteCloudNativeAPIGatewayCertificateWithContext(context.Background(), request)
}

// DeleteCloudNativeAPIGatewayCertificate
// This API is used to delete a cloud-native gateway cert.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteCloudNativeAPIGatewayCertificateWithContext(ctx context.Context, request *DeleteCloudNativeAPIGatewayCertificateRequest) (response *DeleteCloudNativeAPIGatewayCertificateResponse, err error) {
    if request == nil {
        request = NewDeleteCloudNativeAPIGatewayCertificateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "DeleteCloudNativeAPIGatewayCertificate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCloudNativeAPIGatewayCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCloudNativeAPIGatewayCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCloudNativeAPIGatewayPublicNetworkRequest() (request *DeleteCloudNativeAPIGatewayPublicNetworkRequest) {
    request = &DeleteCloudNativeAPIGatewayPublicNetworkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DeleteCloudNativeAPIGatewayPublicNetwork")
    
    
    return
}

func NewDeleteCloudNativeAPIGatewayPublicNetworkResponse() (response *DeleteCloudNativeAPIGatewayPublicNetworkResponse) {
    response = &DeleteCloudNativeAPIGatewayPublicNetworkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCloudNativeAPIGatewayPublicNetwork
// Delete public network configuration
//
// error code that may be returned:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) DeleteCloudNativeAPIGatewayPublicNetwork(request *DeleteCloudNativeAPIGatewayPublicNetworkRequest) (response *DeleteCloudNativeAPIGatewayPublicNetworkResponse, err error) {
    return c.DeleteCloudNativeAPIGatewayPublicNetworkWithContext(context.Background(), request)
}

// DeleteCloudNativeAPIGatewayPublicNetwork
// Delete public network configuration
//
// error code that may be returned:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) DeleteCloudNativeAPIGatewayPublicNetworkWithContext(ctx context.Context, request *DeleteCloudNativeAPIGatewayPublicNetworkRequest) (response *DeleteCloudNativeAPIGatewayPublicNetworkResponse, err error) {
    if request == nil {
        request = NewDeleteCloudNativeAPIGatewayPublicNetworkRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "DeleteCloudNativeAPIGatewayPublicNetwork")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCloudNativeAPIGatewayPublicNetwork require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCloudNativeAPIGatewayPublicNetworkResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCloudNativeAPIGatewayRouteRequest() (request *DeleteCloudNativeAPIGatewayRouteRequest) {
    request = &DeleteCloudNativeAPIGatewayRouteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DeleteCloudNativeAPIGatewayRoute")
    
    
    return
}

func NewDeleteCloudNativeAPIGatewayRouteResponse() (response *DeleteCloudNativeAPIGatewayRouteResponse) {
    response = &DeleteCloudNativeAPIGatewayRouteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCloudNativeAPIGatewayRoute
// Delete a cloud-native gateway route
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteCloudNativeAPIGatewayRoute(request *DeleteCloudNativeAPIGatewayRouteRequest) (response *DeleteCloudNativeAPIGatewayRouteResponse, err error) {
    return c.DeleteCloudNativeAPIGatewayRouteWithContext(context.Background(), request)
}

// DeleteCloudNativeAPIGatewayRoute
// Delete a cloud-native gateway route
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteCloudNativeAPIGatewayRouteWithContext(ctx context.Context, request *DeleteCloudNativeAPIGatewayRouteRequest) (response *DeleteCloudNativeAPIGatewayRouteResponse, err error) {
    if request == nil {
        request = NewDeleteCloudNativeAPIGatewayRouteRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "DeleteCloudNativeAPIGatewayRoute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCloudNativeAPIGatewayRoute require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCloudNativeAPIGatewayRouteResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCloudNativeAPIGatewayRouteRateLimitRequest() (request *DeleteCloudNativeAPIGatewayRouteRateLimitRequest) {
    request = &DeleteCloudNativeAPIGatewayRouteRateLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DeleteCloudNativeAPIGatewayRouteRateLimit")
    
    
    return
}

func NewDeleteCloudNativeAPIGatewayRouteRateLimitResponse() (response *DeleteCloudNativeAPIGatewayRouteRateLimitResponse) {
    response = &DeleteCloudNativeAPIGatewayRouteRateLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCloudNativeAPIGatewayRouteRateLimit
// This API is used to delete a traffic throttling plugin of a cloud-native gateway (Route).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteCloudNativeAPIGatewayRouteRateLimit(request *DeleteCloudNativeAPIGatewayRouteRateLimitRequest) (response *DeleteCloudNativeAPIGatewayRouteRateLimitResponse, err error) {
    return c.DeleteCloudNativeAPIGatewayRouteRateLimitWithContext(context.Background(), request)
}

// DeleteCloudNativeAPIGatewayRouteRateLimit
// This API is used to delete a traffic throttling plugin of a cloud-native gateway (Route).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteCloudNativeAPIGatewayRouteRateLimitWithContext(ctx context.Context, request *DeleteCloudNativeAPIGatewayRouteRateLimitRequest) (response *DeleteCloudNativeAPIGatewayRouteRateLimitResponse, err error) {
    if request == nil {
        request = NewDeleteCloudNativeAPIGatewayRouteRateLimitRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "DeleteCloudNativeAPIGatewayRouteRateLimit")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCloudNativeAPIGatewayRouteRateLimit require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCloudNativeAPIGatewayRouteRateLimitResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCloudNativeAPIGatewayServiceRequest() (request *DeleteCloudNativeAPIGatewayServiceRequest) {
    request = &DeleteCloudNativeAPIGatewayServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DeleteCloudNativeAPIGatewayService")
    
    
    return
}

func NewDeleteCloudNativeAPIGatewayServiceResponse() (response *DeleteCloudNativeAPIGatewayServiceResponse) {
    response = &DeleteCloudNativeAPIGatewayServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCloudNativeAPIGatewayService
// This API is used to delete a cloud-native gateway service.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteCloudNativeAPIGatewayService(request *DeleteCloudNativeAPIGatewayServiceRequest) (response *DeleteCloudNativeAPIGatewayServiceResponse, err error) {
    return c.DeleteCloudNativeAPIGatewayServiceWithContext(context.Background(), request)
}

// DeleteCloudNativeAPIGatewayService
// This API is used to delete a cloud-native gateway service.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteCloudNativeAPIGatewayServiceWithContext(ctx context.Context, request *DeleteCloudNativeAPIGatewayServiceRequest) (response *DeleteCloudNativeAPIGatewayServiceResponse, err error) {
    if request == nil {
        request = NewDeleteCloudNativeAPIGatewayServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "DeleteCloudNativeAPIGatewayService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCloudNativeAPIGatewayService require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCloudNativeAPIGatewayServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCloudNativeAPIGatewayServiceRateLimitRequest() (request *DeleteCloudNativeAPIGatewayServiceRateLimitRequest) {
    request = &DeleteCloudNativeAPIGatewayServiceRateLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DeleteCloudNativeAPIGatewayServiceRateLimit")
    
    
    return
}

func NewDeleteCloudNativeAPIGatewayServiceRateLimitResponse() (response *DeleteCloudNativeAPIGatewayServiceRateLimitResponse) {
    response = &DeleteCloudNativeAPIGatewayServiceRateLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCloudNativeAPIGatewayServiceRateLimit
// This API is used to delete the traffic throttling plugin of a cloud-native gateway (Service).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteCloudNativeAPIGatewayServiceRateLimit(request *DeleteCloudNativeAPIGatewayServiceRateLimitRequest) (response *DeleteCloudNativeAPIGatewayServiceRateLimitResponse, err error) {
    return c.DeleteCloudNativeAPIGatewayServiceRateLimitWithContext(context.Background(), request)
}

// DeleteCloudNativeAPIGatewayServiceRateLimit
// This API is used to delete the traffic throttling plugin of a cloud-native gateway (Service).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteCloudNativeAPIGatewayServiceRateLimitWithContext(ctx context.Context, request *DeleteCloudNativeAPIGatewayServiceRateLimitRequest) (response *DeleteCloudNativeAPIGatewayServiceRateLimitResponse, err error) {
    if request == nil {
        request = NewDeleteCloudNativeAPIGatewayServiceRateLimitRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "DeleteCloudNativeAPIGatewayServiceRateLimit")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCloudNativeAPIGatewayServiceRateLimit require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCloudNativeAPIGatewayServiceRateLimitResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGovernanceLaneGroupsRequest() (request *DeleteGovernanceLaneGroupsRequest) {
    request = &DeleteGovernanceLaneGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DeleteGovernanceLaneGroups")
    
    
    return
}

func NewDeleteGovernanceLaneGroupsResponse() (response *DeleteGovernanceLaneGroupsResponse) {
    response = &DeleteGovernanceLaneGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteGovernanceLaneGroups
// Delete a lane group
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteGovernanceLaneGroups(request *DeleteGovernanceLaneGroupsRequest) (response *DeleteGovernanceLaneGroupsResponse, err error) {
    return c.DeleteGovernanceLaneGroupsWithContext(context.Background(), request)
}

// DeleteGovernanceLaneGroups
// Delete a lane group
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteGovernanceLaneGroupsWithContext(ctx context.Context, request *DeleteGovernanceLaneGroupsRequest) (response *DeleteGovernanceLaneGroupsResponse, err error) {
    if request == nil {
        request = NewDeleteGovernanceLaneGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "DeleteGovernanceLaneGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteGovernanceLaneGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteGovernanceLaneGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNativeGatewayServerGroupRequest() (request *DeleteNativeGatewayServerGroupRequest) {
    request = &DeleteNativeGatewayServerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DeleteNativeGatewayServerGroup")
    
    
    return
}

func NewDeleteNativeGatewayServerGroupResponse() (response *DeleteNativeGatewayServerGroupResponse) {
    response = &DeleteNativeGatewayServerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteNativeGatewayServerGroup
// Delete a Gateway Instance Group
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteNativeGatewayServerGroup(request *DeleteNativeGatewayServerGroupRequest) (response *DeleteNativeGatewayServerGroupResponse, err error) {
    return c.DeleteNativeGatewayServerGroupWithContext(context.Background(), request)
}

// DeleteNativeGatewayServerGroup
// Delete a Gateway Instance Group
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteNativeGatewayServerGroupWithContext(ctx context.Context, request *DeleteNativeGatewayServerGroupRequest) (response *DeleteNativeGatewayServerGroupResponse, err error) {
    if request == nil {
        request = NewDeleteNativeGatewayServerGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "DeleteNativeGatewayServerGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteNativeGatewayServerGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteNativeGatewayServerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNativeGatewayServiceSourceRequest() (request *DeleteNativeGatewayServiceSourceRequest) {
    request = &DeleteNativeGatewayServiceSourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DeleteNativeGatewayServiceSource")
    
    
    return
}

func NewDeleteNativeGatewayServiceSourceResponse() (response *DeleteNativeGatewayServiceSourceResponse) {
    response = &DeleteNativeGatewayServiceSourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteNativeGatewayServiceSource
// Delete a gateway service source instance
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CREATEERROR = "InternalError.CreateError"
//  INTERNALERROR_DECODEERROR = "InternalError.DecodeError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCEINUSE_GATEWAYSERVICESOURCEEXISTSERVICE = "ResourceInUse.GatewayServiceSourceExistService"
func (c *Client) DeleteNativeGatewayServiceSource(request *DeleteNativeGatewayServiceSourceRequest) (response *DeleteNativeGatewayServiceSourceResponse, err error) {
    return c.DeleteNativeGatewayServiceSourceWithContext(context.Background(), request)
}

// DeleteNativeGatewayServiceSource
// Delete a gateway service source instance
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CREATEERROR = "InternalError.CreateError"
//  INTERNALERROR_DECODEERROR = "InternalError.DecodeError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCEINUSE_GATEWAYSERVICESOURCEEXISTSERVICE = "ResourceInUse.GatewayServiceSourceExistService"
func (c *Client) DeleteNativeGatewayServiceSourceWithContext(ctx context.Context, request *DeleteNativeGatewayServiceSourceRequest) (response *DeleteNativeGatewayServiceSourceResponse, err error) {
    if request == nil {
        request = NewDeleteNativeGatewayServiceSourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "DeleteNativeGatewayServiceSource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteNativeGatewayServiceSource require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteNativeGatewayServiceSourceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWafDomainsRequest() (request *DeleteWafDomainsRequest) {
    request = &DeleteWafDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DeleteWafDomains")
    
    
    return
}

func NewDeleteWafDomainsResponse() (response *DeleteWafDomainsResponse) {
    response = &DeleteWafDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteWafDomains
// Delete a WAF-protected domain name
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteWafDomains(request *DeleteWafDomainsRequest) (response *DeleteWafDomainsResponse, err error) {
    return c.DeleteWafDomainsWithContext(context.Background(), request)
}

// DeleteWafDomains
// Delete a WAF-protected domain name
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteWafDomainsWithContext(ctx context.Context, request *DeleteWafDomainsRequest) (response *DeleteWafDomainsResponse, err error) {
    if request == nil {
        request = NewDeleteWafDomainsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "DeleteWafDomains")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteWafDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteWafDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAutoScalerResourceStrategiesRequest() (request *DescribeAutoScalerResourceStrategiesRequest) {
    request = &DescribeAutoScalerResourceStrategiesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeAutoScalerResourceStrategies")
    
    
    return
}

func NewDescribeAutoScalerResourceStrategiesResponse() (response *DescribeAutoScalerResourceStrategiesResponse) {
    response = &DescribeAutoScalerResourceStrategiesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAutoScalerResourceStrategies
// View AS policy list
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CREATEERROR = "InternalError.CreateError"
//  INTERNALERROR_DECODEERROR = "InternalError.DecodeError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UPDATEERROR = "InternalError.UpdateError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) DescribeAutoScalerResourceStrategies(request *DescribeAutoScalerResourceStrategiesRequest) (response *DescribeAutoScalerResourceStrategiesResponse, err error) {
    return c.DescribeAutoScalerResourceStrategiesWithContext(context.Background(), request)
}

// DescribeAutoScalerResourceStrategies
// View AS policy list
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CREATEERROR = "InternalError.CreateError"
//  INTERNALERROR_DECODEERROR = "InternalError.DecodeError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UPDATEERROR = "InternalError.UpdateError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) DescribeAutoScalerResourceStrategiesWithContext(ctx context.Context, request *DescribeAutoScalerResourceStrategiesRequest) (response *DescribeAutoScalerResourceStrategiesResponse, err error) {
    if request == nil {
        request = NewDescribeAutoScalerResourceStrategiesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "DescribeAutoScalerResourceStrategies")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAutoScalerResourceStrategies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAutoScalerResourceStrategiesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAutoScalerResourceStrategyBindingGroupsRequest() (request *DescribeAutoScalerResourceStrategyBindingGroupsRequest) {
    request = &DescribeAutoScalerResourceStrategyBindingGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeAutoScalerResourceStrategyBindingGroups")
    
    
    return
}

func NewDescribeAutoScalerResourceStrategyBindingGroupsResponse() (response *DescribeAutoScalerResourceStrategyBindingGroupsResponse) {
    response = &DescribeAutoScalerResourceStrategyBindingGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAutoScalerResourceStrategyBindingGroups
// View gateway groupings bound to an auto scaling policy
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_DECODEERROR = "InternalError.DecodeError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeAutoScalerResourceStrategyBindingGroups(request *DescribeAutoScalerResourceStrategyBindingGroupsRequest) (response *DescribeAutoScalerResourceStrategyBindingGroupsResponse, err error) {
    return c.DescribeAutoScalerResourceStrategyBindingGroupsWithContext(context.Background(), request)
}

// DescribeAutoScalerResourceStrategyBindingGroups
// View gateway groupings bound to an auto scaling policy
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_DECODEERROR = "InternalError.DecodeError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeAutoScalerResourceStrategyBindingGroupsWithContext(ctx context.Context, request *DescribeAutoScalerResourceStrategyBindingGroupsRequest) (response *DescribeAutoScalerResourceStrategyBindingGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeAutoScalerResourceStrategyBindingGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "DescribeAutoScalerResourceStrategyBindingGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAutoScalerResourceStrategyBindingGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAutoScalerResourceStrategyBindingGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudNativeAPIGatewayRequest() (request *DescribeCloudNativeAPIGatewayRequest) {
    request = &DescribeCloudNativeAPIGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGateway")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewayResponse() (response *DescribeCloudNativeAPIGatewayResponse) {
    response = &DescribeCloudNativeAPIGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudNativeAPIGateway
// This API is used to obtain cloud native API gateway instance information.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGateway(request *DescribeCloudNativeAPIGatewayRequest) (response *DescribeCloudNativeAPIGatewayResponse, err error) {
    return c.DescribeCloudNativeAPIGatewayWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGateway
// This API is used to obtain cloud native API gateway instance information.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewayRequest) (response *DescribeCloudNativeAPIGatewayResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewayRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "DescribeCloudNativeAPIGateway")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGateway require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudNativeAPIGatewayCORSRequest() (request *DescribeCloudNativeAPIGatewayCORSRequest) {
    request = &DescribeCloudNativeAPIGatewayCORSRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewayCORS")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewayCORSResponse() (response *DescribeCloudNativeAPIGatewayCORSResponse) {
    response = &DescribeCloudNativeAPIGatewayCORSResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudNativeAPIGatewayCORS
// Query cloud-native gateway cross-domain configuration
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayCORS(request *DescribeCloudNativeAPIGatewayCORSRequest) (response *DescribeCloudNativeAPIGatewayCORSResponse, err error) {
    return c.DescribeCloudNativeAPIGatewayCORSWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGatewayCORS
// Query cloud-native gateway cross-domain configuration
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayCORSWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewayCORSRequest) (response *DescribeCloudNativeAPIGatewayCORSResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewayCORSRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "DescribeCloudNativeAPIGatewayCORS")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGatewayCORS require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewayCORSResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudNativeAPIGatewayCanaryRulesRequest() (request *DescribeCloudNativeAPIGatewayCanaryRulesRequest) {
    request = &DescribeCloudNativeAPIGatewayCanaryRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewayCanaryRules")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewayCanaryRulesResponse() (response *DescribeCloudNativeAPIGatewayCanaryRulesResponse) {
    response = &DescribeCloudNativeAPIGatewayCanaryRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudNativeAPIGatewayCanaryRules
// Query the grayscale rule list of the cloud-native gateway
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayCanaryRules(request *DescribeCloudNativeAPIGatewayCanaryRulesRequest) (response *DescribeCloudNativeAPIGatewayCanaryRulesResponse, err error) {
    return c.DescribeCloudNativeAPIGatewayCanaryRulesWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGatewayCanaryRules
// Query the grayscale rule list of the cloud-native gateway
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayCanaryRulesWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewayCanaryRulesRequest) (response *DescribeCloudNativeAPIGatewayCanaryRulesResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewayCanaryRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "DescribeCloudNativeAPIGatewayCanaryRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGatewayCanaryRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewayCanaryRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudNativeAPIGatewayCertificateDetailsRequest() (request *DescribeCloudNativeAPIGatewayCertificateDetailsRequest) {
    request = &DescribeCloudNativeAPIGatewayCertificateDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewayCertificateDetails")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewayCertificateDetailsResponse() (response *DescribeCloudNativeAPIGatewayCertificateDetailsResponse) {
    response = &DescribeCloudNativeAPIGatewayCertificateDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudNativeAPIGatewayCertificateDetails
// Query the certificate detail of one cloud-native gateway
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayCertificateDetails(request *DescribeCloudNativeAPIGatewayCertificateDetailsRequest) (response *DescribeCloudNativeAPIGatewayCertificateDetailsResponse, err error) {
    return c.DescribeCloudNativeAPIGatewayCertificateDetailsWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGatewayCertificateDetails
// Query the certificate detail of one cloud-native gateway
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayCertificateDetailsWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewayCertificateDetailsRequest) (response *DescribeCloudNativeAPIGatewayCertificateDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewayCertificateDetailsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "DescribeCloudNativeAPIGatewayCertificateDetails")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGatewayCertificateDetails require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewayCertificateDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudNativeAPIGatewayCertificatesRequest() (request *DescribeCloudNativeAPIGatewayCertificatesRequest) {
    request = &DescribeCloudNativeAPIGatewayCertificatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewayCertificates")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewayCertificatesResponse() (response *DescribeCloudNativeAPIGatewayCertificatesResponse) {
    response = &DescribeCloudNativeAPIGatewayCertificatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudNativeAPIGatewayCertificates
// Query the certificate list of the cloud-native gateway
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayCertificates(request *DescribeCloudNativeAPIGatewayCertificatesRequest) (response *DescribeCloudNativeAPIGatewayCertificatesResponse, err error) {
    return c.DescribeCloudNativeAPIGatewayCertificatesWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGatewayCertificates
// Query the certificate list of the cloud-native gateway
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayCertificatesWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewayCertificatesRequest) (response *DescribeCloudNativeAPIGatewayCertificatesResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewayCertificatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "DescribeCloudNativeAPIGatewayCertificates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGatewayCertificates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewayCertificatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudNativeAPIGatewayConfigRequest() (request *DescribeCloudNativeAPIGatewayConfigRequest) {
    request = &DescribeCloudNativeAPIGatewayConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewayConfig")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewayConfigResponse() (response *DescribeCloudNativeAPIGatewayConfigResponse) {
    response = &DescribeCloudNativeAPIGatewayConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudNativeAPIGatewayConfig
// This API is used to obtain cloud native API gateway instance network configuration information
//
// error code that may be returned:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR_CREATEERROR = "InternalError.CreateError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayConfig(request *DescribeCloudNativeAPIGatewayConfigRequest) (response *DescribeCloudNativeAPIGatewayConfigResponse, err error) {
    return c.DescribeCloudNativeAPIGatewayConfigWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGatewayConfig
// This API is used to obtain cloud native API gateway instance network configuration information
//
// error code that may be returned:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR_CREATEERROR = "InternalError.CreateError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayConfigWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewayConfigRequest) (response *DescribeCloudNativeAPIGatewayConfigResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewayConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "DescribeCloudNativeAPIGatewayConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGatewayConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewayConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudNativeAPIGatewayInfoByIpRequest() (request *DescribeCloudNativeAPIGatewayInfoByIpRequest) {
    request = &DescribeCloudNativeAPIGatewayInfoByIpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewayInfoByIp")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewayInfoByIpResponse() (response *DescribeCloudNativeAPIGatewayInfoByIpResponse) {
    response = &DescribeCloudNativeAPIGatewayInfoByIpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudNativeAPIGatewayInfoByIp
// Query cloud native gateway instance information based on public IP address
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) DescribeCloudNativeAPIGatewayInfoByIp(request *DescribeCloudNativeAPIGatewayInfoByIpRequest) (response *DescribeCloudNativeAPIGatewayInfoByIpResponse, err error) {
    return c.DescribeCloudNativeAPIGatewayInfoByIpWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGatewayInfoByIp
// Query cloud native gateway instance information based on public IP address
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) DescribeCloudNativeAPIGatewayInfoByIpWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewayInfoByIpRequest) (response *DescribeCloudNativeAPIGatewayInfoByIpResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewayInfoByIpRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "DescribeCloudNativeAPIGatewayInfoByIp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGatewayInfoByIp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewayInfoByIpResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudNativeAPIGatewayNodesRequest() (request *DescribeCloudNativeAPIGatewayNodesRequest) {
    request = &DescribeCloudNativeAPIGatewayNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewayNodes")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewayNodesResponse() (response *DescribeCloudNativeAPIGatewayNodesResponse) {
    response = &DescribeCloudNativeAPIGatewayNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudNativeAPIGatewayNodes
// This API is used to get a cloud-native gateway node list
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETERVALUE_ACTION = "InvalidParameterValue.Action"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) DescribeCloudNativeAPIGatewayNodes(request *DescribeCloudNativeAPIGatewayNodesRequest) (response *DescribeCloudNativeAPIGatewayNodesResponse, err error) {
    return c.DescribeCloudNativeAPIGatewayNodesWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGatewayNodes
// This API is used to get a cloud-native gateway node list
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETERVALUE_ACTION = "InvalidParameterValue.Action"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) DescribeCloudNativeAPIGatewayNodesWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewayNodesRequest) (response *DescribeCloudNativeAPIGatewayNodesResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewayNodesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "DescribeCloudNativeAPIGatewayNodes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGatewayNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewayNodesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudNativeAPIGatewayPortsRequest() (request *DescribeCloudNativeAPIGatewayPortsRequest) {
    request = &DescribeCloudNativeAPIGatewayPortsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewayPorts")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewayPortsResponse() (response *DescribeCloudNativeAPIGatewayPortsResponse) {
    response = &DescribeCloudNativeAPIGatewayPortsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudNativeAPIGatewayPorts
// Retrieve port information of a cloud native API gateway instance
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
func (c *Client) DescribeCloudNativeAPIGatewayPorts(request *DescribeCloudNativeAPIGatewayPortsRequest) (response *DescribeCloudNativeAPIGatewayPortsResponse, err error) {
    return c.DescribeCloudNativeAPIGatewayPortsWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGatewayPorts
// Retrieve port information of a cloud native API gateway instance
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
func (c *Client) DescribeCloudNativeAPIGatewayPortsWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewayPortsRequest) (response *DescribeCloudNativeAPIGatewayPortsResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewayPortsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "DescribeCloudNativeAPIGatewayPorts")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGatewayPorts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewayPortsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudNativeAPIGatewayRouteRateLimitRequest() (request *DescribeCloudNativeAPIGatewayRouteRateLimitRequest) {
    request = &DescribeCloudNativeAPIGatewayRouteRateLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewayRouteRateLimit")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewayRouteRateLimitResponse() (response *DescribeCloudNativeAPIGatewayRouteRateLimitResponse) {
    response = &DescribeCloudNativeAPIGatewayRouteRateLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudNativeAPIGatewayRouteRateLimit
// Query the traffic throttling plugin of a cloud-native gateway (Route).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayRouteRateLimit(request *DescribeCloudNativeAPIGatewayRouteRateLimitRequest) (response *DescribeCloudNativeAPIGatewayRouteRateLimitResponse, err error) {
    return c.DescribeCloudNativeAPIGatewayRouteRateLimitWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGatewayRouteRateLimit
// Query the traffic throttling plugin of a cloud-native gateway (Route).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayRouteRateLimitWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewayRouteRateLimitRequest) (response *DescribeCloudNativeAPIGatewayRouteRateLimitResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewayRouteRateLimitRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "DescribeCloudNativeAPIGatewayRouteRateLimit")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGatewayRouteRateLimit require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewayRouteRateLimitResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudNativeAPIGatewayRoutesRequest() (request *DescribeCloudNativeAPIGatewayRoutesRequest) {
    request = &DescribeCloudNativeAPIGatewayRoutesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewayRoutes")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewayRoutesResponse() (response *DescribeCloudNativeAPIGatewayRoutesResponse) {
    response = &DescribeCloudNativeAPIGatewayRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudNativeAPIGatewayRoutes
// Query the routing list of the cloud-native gateway
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayRoutes(request *DescribeCloudNativeAPIGatewayRoutesRequest) (response *DescribeCloudNativeAPIGatewayRoutesResponse, err error) {
    return c.DescribeCloudNativeAPIGatewayRoutesWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGatewayRoutes
// Query the routing list of the cloud-native gateway
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayRoutesWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewayRoutesRequest) (response *DescribeCloudNativeAPIGatewayRoutesResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewayRoutesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "DescribeCloudNativeAPIGatewayRoutes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGatewayRoutes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewayRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudNativeAPIGatewayServiceRateLimitRequest() (request *DescribeCloudNativeAPIGatewayServiceRateLimitRequest) {
    request = &DescribeCloudNativeAPIGatewayServiceRateLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewayServiceRateLimit")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewayServiceRateLimitResponse() (response *DescribeCloudNativeAPIGatewayServiceRateLimitResponse) {
    response = &DescribeCloudNativeAPIGatewayServiceRateLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudNativeAPIGatewayServiceRateLimit
// This API is used to query the traffic throttling plugin of a cloud-native gateway (Service).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayServiceRateLimit(request *DescribeCloudNativeAPIGatewayServiceRateLimitRequest) (response *DescribeCloudNativeAPIGatewayServiceRateLimitResponse, err error) {
    return c.DescribeCloudNativeAPIGatewayServiceRateLimitWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGatewayServiceRateLimit
// This API is used to query the traffic throttling plugin of a cloud-native gateway (Service).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayServiceRateLimitWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewayServiceRateLimitRequest) (response *DescribeCloudNativeAPIGatewayServiceRateLimitResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewayServiceRateLimitRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "DescribeCloudNativeAPIGatewayServiceRateLimit")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGatewayServiceRateLimit require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewayServiceRateLimitResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudNativeAPIGatewayServicesRequest() (request *DescribeCloudNativeAPIGatewayServicesRequest) {
    request = &DescribeCloudNativeAPIGatewayServicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewayServices")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewayServicesResponse() (response *DescribeCloudNativeAPIGatewayServicesResponse) {
    response = &DescribeCloudNativeAPIGatewayServicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudNativeAPIGatewayServices
// Query the service list of the cloud-native gateway
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayServices(request *DescribeCloudNativeAPIGatewayServicesRequest) (response *DescribeCloudNativeAPIGatewayServicesResponse, err error) {
    return c.DescribeCloudNativeAPIGatewayServicesWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGatewayServices
// Query the service list of the cloud-native gateway
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayServicesWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewayServicesRequest) (response *DescribeCloudNativeAPIGatewayServicesResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewayServicesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "DescribeCloudNativeAPIGatewayServices")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGatewayServices require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewayServicesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudNativeAPIGatewayServicesLightRequest() (request *DescribeCloudNativeAPIGatewayServicesLightRequest) {
    request = &DescribeCloudNativeAPIGatewayServicesLightRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewayServicesLight")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewayServicesLightResponse() (response *DescribeCloudNativeAPIGatewayServicesLightResponse) {
    response = &DescribeCloudNativeAPIGatewayServicesLightResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudNativeAPIGatewayServicesLight
// Lightweight query the service list of the cloud-native gateway
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayServicesLight(request *DescribeCloudNativeAPIGatewayServicesLightRequest) (response *DescribeCloudNativeAPIGatewayServicesLightResponse, err error) {
    return c.DescribeCloudNativeAPIGatewayServicesLightWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGatewayServicesLight
// Lightweight query the service list of the cloud-native gateway
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayServicesLightWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewayServicesLightRequest) (response *DescribeCloudNativeAPIGatewayServicesLightResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewayServicesLightRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "DescribeCloudNativeAPIGatewayServicesLight")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGatewayServicesLight require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewayServicesLightResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudNativeAPIGatewayUpstreamRequest() (request *DescribeCloudNativeAPIGatewayUpstreamRequest) {
    request = &DescribeCloudNativeAPIGatewayUpstreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGatewayUpstream")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewayUpstreamResponse() (response *DescribeCloudNativeAPIGatewayUpstreamResponse) {
    response = &DescribeCloudNativeAPIGatewayUpstreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudNativeAPIGatewayUpstream
// This API is used to query the Upstream list in the service detail of a cloud-native gateway.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayUpstream(request *DescribeCloudNativeAPIGatewayUpstreamRequest) (response *DescribeCloudNativeAPIGatewayUpstreamResponse, err error) {
    return c.DescribeCloudNativeAPIGatewayUpstreamWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGatewayUpstream
// This API is used to query the Upstream list in the service detail of a cloud-native gateway.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeCloudNativeAPIGatewayUpstreamWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewayUpstreamRequest) (response *DescribeCloudNativeAPIGatewayUpstreamResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewayUpstreamRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "DescribeCloudNativeAPIGatewayUpstream")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGatewayUpstream require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewayUpstreamResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudNativeAPIGatewaysRequest() (request *DescribeCloudNativeAPIGatewaysRequest) {
    request = &DescribeCloudNativeAPIGatewaysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeCloudNativeAPIGateways")
    
    
    return
}

func NewDescribeCloudNativeAPIGatewaysResponse() (response *DescribeCloudNativeAPIGatewaysResponse) {
    response = &DescribeCloudNativeAPIGatewaysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudNativeAPIGateways
// This API is used to obtain the cloud native API gateway instance list.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  INTERNALERROR_HTTPSTATUSCODEERROR = "InternalError.HttpStatusCodeError"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCloudNativeAPIGateways(request *DescribeCloudNativeAPIGatewaysRequest) (response *DescribeCloudNativeAPIGatewaysResponse, err error) {
    return c.DescribeCloudNativeAPIGatewaysWithContext(context.Background(), request)
}

// DescribeCloudNativeAPIGateways
// This API is used to obtain the cloud native API gateway instance list.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  INTERNALERROR_HTTPSTATUSCODEERROR = "InternalError.HttpStatusCodeError"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCloudNativeAPIGatewaysWithContext(ctx context.Context, request *DescribeCloudNativeAPIGatewaysRequest) (response *DescribeCloudNativeAPIGatewaysResponse, err error) {
    if request == nil {
        request = NewDescribeCloudNativeAPIGatewaysRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "DescribeCloudNativeAPIGateways")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudNativeAPIGateways require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudNativeAPIGatewaysResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGovernanceLaneGroupsRequest() (request *DescribeGovernanceLaneGroupsRequest) {
    request = &DescribeGovernanceLaneGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeGovernanceLaneGroups")
    
    
    return
}

func NewDescribeGovernanceLaneGroupsResponse() (response *DescribeGovernanceLaneGroupsResponse) {
    response = &DescribeGovernanceLaneGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGovernanceLaneGroups
// Query lane group list
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  INTERNALERROR_HTTPSTATUSCODEERROR = "InternalError.HttpStatusCodeError"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeGovernanceLaneGroups(request *DescribeGovernanceLaneGroupsRequest) (response *DescribeGovernanceLaneGroupsResponse, err error) {
    return c.DescribeGovernanceLaneGroupsWithContext(context.Background(), request)
}

// DescribeGovernanceLaneGroups
// Query lane group list
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  INTERNALERROR_HTTPSTATUSCODEERROR = "InternalError.HttpStatusCodeError"
//  INTERNALERROR_IOERROR = "InternalError.IOError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeGovernanceLaneGroupsWithContext(ctx context.Context, request *DescribeGovernanceLaneGroupsRequest) (response *DescribeGovernanceLaneGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeGovernanceLaneGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "DescribeGovernanceLaneGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGovernanceLaneGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGovernanceLaneGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNativeGatewayServerGroupsRequest() (request *DescribeNativeGatewayServerGroupsRequest) {
    request = &DescribeNativeGatewayServerGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeNativeGatewayServerGroups")
    
    
    return
}

func NewDescribeNativeGatewayServerGroupsResponse() (response *DescribeNativeGatewayServerGroupsResponse) {
    response = &DescribeNativeGatewayServerGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNativeGatewayServerGroups
// Query cloud native gateway group information
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) DescribeNativeGatewayServerGroups(request *DescribeNativeGatewayServerGroupsRequest) (response *DescribeNativeGatewayServerGroupsResponse, err error) {
    return c.DescribeNativeGatewayServerGroupsWithContext(context.Background(), request)
}

// DescribeNativeGatewayServerGroups
// Query cloud native gateway group information
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) DescribeNativeGatewayServerGroupsWithContext(ctx context.Context, request *DescribeNativeGatewayServerGroupsRequest) (response *DescribeNativeGatewayServerGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeNativeGatewayServerGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "DescribeNativeGatewayServerGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNativeGatewayServerGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNativeGatewayServerGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNativeGatewayServiceSourcesRequest() (request *DescribeNativeGatewayServiceSourcesRequest) {
    request = &DescribeNativeGatewayServiceSourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeNativeGatewayServiceSources")
    
    
    return
}

func NewDescribeNativeGatewayServiceSourcesResponse() (response *DescribeNativeGatewayServiceSourcesResponse) {
    response = &DescribeNativeGatewayServiceSourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNativeGatewayServiceSources
// Query the instance list of the gateway service source
//
// error code that may be returned:
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) DescribeNativeGatewayServiceSources(request *DescribeNativeGatewayServiceSourcesRequest) (response *DescribeNativeGatewayServiceSourcesResponse, err error) {
    return c.DescribeNativeGatewayServiceSourcesWithContext(context.Background(), request)
}

// DescribeNativeGatewayServiceSources
// Query the instance list of the gateway service source
//
// error code that may be returned:
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) DescribeNativeGatewayServiceSourcesWithContext(ctx context.Context, request *DescribeNativeGatewayServiceSourcesRequest) (response *DescribeNativeGatewayServiceSourcesResponse, err error) {
    if request == nil {
        request = NewDescribeNativeGatewayServiceSourcesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "DescribeNativeGatewayServiceSources")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNativeGatewayServiceSources require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNativeGatewayServiceSourcesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOneCloudNativeAPIGatewayServiceRequest() (request *DescribeOneCloudNativeAPIGatewayServiceRequest) {
    request = &DescribeOneCloudNativeAPIGatewayServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeOneCloudNativeAPIGatewayService")
    
    
    return
}

func NewDescribeOneCloudNativeAPIGatewayServiceResponse() (response *DescribeOneCloudNativeAPIGatewayServiceResponse) {
    response = &DescribeOneCloudNativeAPIGatewayServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOneCloudNativeAPIGatewayService
// This API is used to obtain the service detail of the cloud-native gateway.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeOneCloudNativeAPIGatewayService(request *DescribeOneCloudNativeAPIGatewayServiceRequest) (response *DescribeOneCloudNativeAPIGatewayServiceResponse, err error) {
    return c.DescribeOneCloudNativeAPIGatewayServiceWithContext(context.Background(), request)
}

// DescribeOneCloudNativeAPIGatewayService
// This API is used to obtain the service detail of the cloud-native gateway.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeOneCloudNativeAPIGatewayServiceWithContext(ctx context.Context, request *DescribeOneCloudNativeAPIGatewayServiceRequest) (response *DescribeOneCloudNativeAPIGatewayServiceResponse, err error) {
    if request == nil {
        request = NewDescribeOneCloudNativeAPIGatewayServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "DescribeOneCloudNativeAPIGatewayService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOneCloudNativeAPIGatewayService require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOneCloudNativeAPIGatewayServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePublicAddressConfigRequest() (request *DescribePublicAddressConfigRequest) {
    request = &DescribePublicAddressConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribePublicAddressConfig")
    
    
    return
}

func NewDescribePublicAddressConfigResponse() (response *DescribePublicAddressConfigResponse) {
    response = &DescribePublicAddressConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePublicAddressConfig
// Query public IP address info
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) DescribePublicAddressConfig(request *DescribePublicAddressConfigRequest) (response *DescribePublicAddressConfigResponse, err error) {
    return c.DescribePublicAddressConfigWithContext(context.Background(), request)
}

// DescribePublicAddressConfig
// Query public IP address info
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) DescribePublicAddressConfigWithContext(ctx context.Context, request *DescribePublicAddressConfigRequest) (response *DescribePublicAddressConfigResponse, err error) {
    if request == nil {
        request = NewDescribePublicAddressConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "DescribePublicAddressConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePublicAddressConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePublicAddressConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePublicNetworkRequest() (request *DescribePublicNetworkRequest) {
    request = &DescribePublicNetworkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribePublicNetwork")
    
    
    return
}

func NewDescribePublicNetworkResponse() (response *DescribePublicNetworkResponse) {
    response = &DescribePublicNetworkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePublicNetwork
// Query the public network details of a cloud native API gateway instance
//
// error code that may be returned:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR_CREATEERROR = "InternalError.CreateError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribePublicNetwork(request *DescribePublicNetworkRequest) (response *DescribePublicNetworkResponse, err error) {
    return c.DescribePublicNetworkWithContext(context.Background(), request)
}

// DescribePublicNetwork
// Query the public network details of a cloud native API gateway instance
//
// error code that may be returned:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR_CREATEERROR = "InternalError.CreateError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED_LIMITEXCEEDED = "LimitExceeded.LimitExceeded"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribePublicNetworkWithContext(ctx context.Context, request *DescribePublicNetworkRequest) (response *DescribePublicNetworkResponse, err error) {
    if request == nil {
        request = NewDescribePublicNetworkRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "DescribePublicNetwork")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePublicNetwork require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePublicNetworkResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUpstreamHealthCheckConfigRequest() (request *DescribeUpstreamHealthCheckConfigRequest) {
    request = &DescribeUpstreamHealthCheckConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeUpstreamHealthCheckConfig")
    
    
    return
}

func NewDescribeUpstreamHealthCheckConfigResponse() (response *DescribeUpstreamHealthCheckConfigResponse) {
    response = &DescribeUpstreamHealthCheckConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUpstreamHealthCheckConfig
// This API is used to obtain the health check configuration of the cloud-native gateway service.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) DescribeUpstreamHealthCheckConfig(request *DescribeUpstreamHealthCheckConfigRequest) (response *DescribeUpstreamHealthCheckConfigResponse, err error) {
    return c.DescribeUpstreamHealthCheckConfigWithContext(context.Background(), request)
}

// DescribeUpstreamHealthCheckConfig
// This API is used to obtain the health check configuration of the cloud-native gateway service.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) DescribeUpstreamHealthCheckConfigWithContext(ctx context.Context, request *DescribeUpstreamHealthCheckConfigRequest) (response *DescribeUpstreamHealthCheckConfigResponse, err error) {
    if request == nil {
        request = NewDescribeUpstreamHealthCheckConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "DescribeUpstreamHealthCheckConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUpstreamHealthCheckConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUpstreamHealthCheckConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWafDomainsRequest() (request *DescribeWafDomainsRequest) {
    request = &DescribeWafDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeWafDomains")
    
    
    return
}

func NewDescribeWafDomainsResponse() (response *DescribeWafDomainsResponse) {
    response = &DescribeWafDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWafDomains
// Query a WAF-protected domain name
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeWafDomains(request *DescribeWafDomainsRequest) (response *DescribeWafDomainsResponse, err error) {
    return c.DescribeWafDomainsWithContext(context.Background(), request)
}

// DescribeWafDomains
// Query a WAF-protected domain name
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeWafDomainsWithContext(ctx context.Context, request *DescribeWafDomainsRequest) (response *DescribeWafDomainsResponse, err error) {
    if request == nil {
        request = NewDescribeWafDomainsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "DescribeWafDomains")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWafDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWafDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWafProtectionRequest() (request *DescribeWafProtectionRequest) {
    request = &DescribeWafProtectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "DescribeWafProtection")
    
    
    return
}

func NewDescribeWafProtectionResponse() (response *DescribeWafProtectionResponse) {
    response = &DescribeWafProtectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWafProtection
// Query WAF protection status
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeWafProtection(request *DescribeWafProtectionRequest) (response *DescribeWafProtectionResponse, err error) {
    return c.DescribeWafProtectionWithContext(context.Background(), request)
}

// DescribeWafProtection
// Query WAF protection status
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeWafProtectionWithContext(ctx context.Context, request *DescribeWafProtectionRequest) (response *DescribeWafProtectionResponse, err error) {
    if request == nil {
        request = NewDescribeWafProtectionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "DescribeWafProtection")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWafProtection require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWafProtectionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAutoScalerResourceStrategyRequest() (request *ModifyAutoScalerResourceStrategyRequest) {
    request = &ModifyAutoScalerResourceStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "ModifyAutoScalerResourceStrategy")
    
    
    return
}

func NewModifyAutoScalerResourceStrategyResponse() (response *ModifyAutoScalerResourceStrategyResponse) {
    response = &ModifyAutoScalerResourceStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAutoScalerResourceStrategy
// Update AS policy
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyAutoScalerResourceStrategy(request *ModifyAutoScalerResourceStrategyRequest) (response *ModifyAutoScalerResourceStrategyResponse, err error) {
    return c.ModifyAutoScalerResourceStrategyWithContext(context.Background(), request)
}

// ModifyAutoScalerResourceStrategy
// Update AS policy
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyAutoScalerResourceStrategyWithContext(ctx context.Context, request *ModifyAutoScalerResourceStrategyRequest) (response *ModifyAutoScalerResourceStrategyResponse, err error) {
    if request == nil {
        request = NewModifyAutoScalerResourceStrategyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "ModifyAutoScalerResourceStrategy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAutoScalerResourceStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAutoScalerResourceStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCloudNativeAPIGatewayRequest() (request *ModifyCloudNativeAPIGatewayRequest) {
    request = &ModifyCloudNativeAPIGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "ModifyCloudNativeAPIGateway")
    
    
    return
}

func NewModifyCloudNativeAPIGatewayResponse() (response *ModifyCloudNativeAPIGatewayResponse) {
    response = &ModifyCloudNativeAPIGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCloudNativeAPIGateway
// This API is used to modify the basic information of a cloud native API gateway instance.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyCloudNativeAPIGateway(request *ModifyCloudNativeAPIGatewayRequest) (response *ModifyCloudNativeAPIGatewayResponse, err error) {
    return c.ModifyCloudNativeAPIGatewayWithContext(context.Background(), request)
}

// ModifyCloudNativeAPIGateway
// This API is used to modify the basic information of a cloud native API gateway instance.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyCloudNativeAPIGatewayWithContext(ctx context.Context, request *ModifyCloudNativeAPIGatewayRequest) (response *ModifyCloudNativeAPIGatewayResponse, err error) {
    if request == nil {
        request = NewModifyCloudNativeAPIGatewayRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "ModifyCloudNativeAPIGateway")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCloudNativeAPIGateway require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCloudNativeAPIGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCloudNativeAPIGatewayCanaryRuleRequest() (request *ModifyCloudNativeAPIGatewayCanaryRuleRequest) {
    request = &ModifyCloudNativeAPIGatewayCanaryRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "ModifyCloudNativeAPIGatewayCanaryRule")
    
    
    return
}

func NewModifyCloudNativeAPIGatewayCanaryRuleResponse() (response *ModifyCloudNativeAPIGatewayCanaryRuleResponse) {
    response = &ModifyCloudNativeAPIGatewayCanaryRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCloudNativeAPIGatewayCanaryRule
// Modify the grayscale rule of the cloud-native gateway
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyCloudNativeAPIGatewayCanaryRule(request *ModifyCloudNativeAPIGatewayCanaryRuleRequest) (response *ModifyCloudNativeAPIGatewayCanaryRuleResponse, err error) {
    return c.ModifyCloudNativeAPIGatewayCanaryRuleWithContext(context.Background(), request)
}

// ModifyCloudNativeAPIGatewayCanaryRule
// Modify the grayscale rule of the cloud-native gateway
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyCloudNativeAPIGatewayCanaryRuleWithContext(ctx context.Context, request *ModifyCloudNativeAPIGatewayCanaryRuleRequest) (response *ModifyCloudNativeAPIGatewayCanaryRuleResponse, err error) {
    if request == nil {
        request = NewModifyCloudNativeAPIGatewayCanaryRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "ModifyCloudNativeAPIGatewayCanaryRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCloudNativeAPIGatewayCanaryRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCloudNativeAPIGatewayCanaryRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCloudNativeAPIGatewayCertificateRequest() (request *ModifyCloudNativeAPIGatewayCertificateRequest) {
    request = &ModifyCloudNativeAPIGatewayCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "ModifyCloudNativeAPIGatewayCertificate")
    
    
    return
}

func NewModifyCloudNativeAPIGatewayCertificateResponse() (response *ModifyCloudNativeAPIGatewayCertificateResponse) {
    response = &ModifyCloudNativeAPIGatewayCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCloudNativeAPIGatewayCertificate
// Update the cloud-native gateway certificate
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  LIMITEXCEEDED_LBDOMAINS = "LimitExceeded.LBDomains"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyCloudNativeAPIGatewayCertificate(request *ModifyCloudNativeAPIGatewayCertificateRequest) (response *ModifyCloudNativeAPIGatewayCertificateResponse, err error) {
    return c.ModifyCloudNativeAPIGatewayCertificateWithContext(context.Background(), request)
}

// ModifyCloudNativeAPIGatewayCertificate
// Update the cloud-native gateway certificate
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  LIMITEXCEEDED_LBDOMAINS = "LimitExceeded.LBDomains"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyCloudNativeAPIGatewayCertificateWithContext(ctx context.Context, request *ModifyCloudNativeAPIGatewayCertificateRequest) (response *ModifyCloudNativeAPIGatewayCertificateResponse, err error) {
    if request == nil {
        request = NewModifyCloudNativeAPIGatewayCertificateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "ModifyCloudNativeAPIGatewayCertificate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCloudNativeAPIGatewayCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCloudNativeAPIGatewayCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCloudNativeAPIGatewayRouteRequest() (request *ModifyCloudNativeAPIGatewayRouteRequest) {
    request = &ModifyCloudNativeAPIGatewayRouteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "ModifyCloudNativeAPIGatewayRoute")
    
    
    return
}

func NewModifyCloudNativeAPIGatewayRouteResponse() (response *ModifyCloudNativeAPIGatewayRouteResponse) {
    response = &ModifyCloudNativeAPIGatewayRouteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCloudNativeAPIGatewayRoute
// This API is used to modify a cloud-native gateway route.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_CREATEERROR = "MissingParameter.CreateError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyCloudNativeAPIGatewayRoute(request *ModifyCloudNativeAPIGatewayRouteRequest) (response *ModifyCloudNativeAPIGatewayRouteResponse, err error) {
    return c.ModifyCloudNativeAPIGatewayRouteWithContext(context.Background(), request)
}

// ModifyCloudNativeAPIGatewayRoute
// This API is used to modify a cloud-native gateway route.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_CREATEERROR = "MissingParameter.CreateError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyCloudNativeAPIGatewayRouteWithContext(ctx context.Context, request *ModifyCloudNativeAPIGatewayRouteRequest) (response *ModifyCloudNativeAPIGatewayRouteResponse, err error) {
    if request == nil {
        request = NewModifyCloudNativeAPIGatewayRouteRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "ModifyCloudNativeAPIGatewayRoute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCloudNativeAPIGatewayRoute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCloudNativeAPIGatewayRouteResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCloudNativeAPIGatewayRouteRateLimitRequest() (request *ModifyCloudNativeAPIGatewayRouteRateLimitRequest) {
    request = &ModifyCloudNativeAPIGatewayRouteRateLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "ModifyCloudNativeAPIGatewayRouteRateLimit")
    
    
    return
}

func NewModifyCloudNativeAPIGatewayRouteRateLimitResponse() (response *ModifyCloudNativeAPIGatewayRouteRateLimitResponse) {
    response = &ModifyCloudNativeAPIGatewayRouteRateLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCloudNativeAPIGatewayRouteRateLimit
// This API is used to modify the traffic throttling plugin of a cloud-native gateway (Route).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyCloudNativeAPIGatewayRouteRateLimit(request *ModifyCloudNativeAPIGatewayRouteRateLimitRequest) (response *ModifyCloudNativeAPIGatewayRouteRateLimitResponse, err error) {
    return c.ModifyCloudNativeAPIGatewayRouteRateLimitWithContext(context.Background(), request)
}

// ModifyCloudNativeAPIGatewayRouteRateLimit
// This API is used to modify the traffic throttling plugin of a cloud-native gateway (Route).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyCloudNativeAPIGatewayRouteRateLimitWithContext(ctx context.Context, request *ModifyCloudNativeAPIGatewayRouteRateLimitRequest) (response *ModifyCloudNativeAPIGatewayRouteRateLimitResponse, err error) {
    if request == nil {
        request = NewModifyCloudNativeAPIGatewayRouteRateLimitRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "ModifyCloudNativeAPIGatewayRouteRateLimit")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCloudNativeAPIGatewayRouteRateLimit require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCloudNativeAPIGatewayRouteRateLimitResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCloudNativeAPIGatewayServiceRequest() (request *ModifyCloudNativeAPIGatewayServiceRequest) {
    request = &ModifyCloudNativeAPIGatewayServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "ModifyCloudNativeAPIGatewayService")
    
    
    return
}

func NewModifyCloudNativeAPIGatewayServiceResponse() (response *ModifyCloudNativeAPIGatewayServiceResponse) {
    response = &ModifyCloudNativeAPIGatewayServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCloudNativeAPIGatewayService
// Modify a cloud-native gateway service
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyCloudNativeAPIGatewayService(request *ModifyCloudNativeAPIGatewayServiceRequest) (response *ModifyCloudNativeAPIGatewayServiceResponse, err error) {
    return c.ModifyCloudNativeAPIGatewayServiceWithContext(context.Background(), request)
}

// ModifyCloudNativeAPIGatewayService
// Modify a cloud-native gateway service
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyCloudNativeAPIGatewayServiceWithContext(ctx context.Context, request *ModifyCloudNativeAPIGatewayServiceRequest) (response *ModifyCloudNativeAPIGatewayServiceResponse, err error) {
    if request == nil {
        request = NewModifyCloudNativeAPIGatewayServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "ModifyCloudNativeAPIGatewayService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCloudNativeAPIGatewayService require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCloudNativeAPIGatewayServiceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCloudNativeAPIGatewayServiceRateLimitRequest() (request *ModifyCloudNativeAPIGatewayServiceRateLimitRequest) {
    request = &ModifyCloudNativeAPIGatewayServiceRateLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "ModifyCloudNativeAPIGatewayServiceRateLimit")
    
    
    return
}

func NewModifyCloudNativeAPIGatewayServiceRateLimitResponse() (response *ModifyCloudNativeAPIGatewayServiceRateLimitResponse) {
    response = &ModifyCloudNativeAPIGatewayServiceRateLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCloudNativeAPIGatewayServiceRateLimit
// This API is used to modify the traffic throttling plugin of a cloud-native gateway (Service).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyCloudNativeAPIGatewayServiceRateLimit(request *ModifyCloudNativeAPIGatewayServiceRateLimitRequest) (response *ModifyCloudNativeAPIGatewayServiceRateLimitResponse, err error) {
    return c.ModifyCloudNativeAPIGatewayServiceRateLimitWithContext(context.Background(), request)
}

// ModifyCloudNativeAPIGatewayServiceRateLimit
// This API is used to modify the traffic throttling plugin of a cloud-native gateway (Service).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyCloudNativeAPIGatewayServiceRateLimitWithContext(ctx context.Context, request *ModifyCloudNativeAPIGatewayServiceRateLimitRequest) (response *ModifyCloudNativeAPIGatewayServiceRateLimitResponse, err error) {
    if request == nil {
        request = NewModifyCloudNativeAPIGatewayServiceRateLimitRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "ModifyCloudNativeAPIGatewayServiceRateLimit")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCloudNativeAPIGatewayServiceRateLimit require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCloudNativeAPIGatewayServiceRateLimitResponse()
    err = c.Send(request, response)
    return
}

func NewModifyConsoleNetworkRequest() (request *ModifyConsoleNetworkRequest) {
    request = &ModifyConsoleNetworkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "ModifyConsoleNetwork")
    
    
    return
}

func NewModifyConsoleNetworkResponse() (response *ModifyConsoleNetworkResponse) {
    response = &ModifyConsoleNetworkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyConsoleNetwork
// Modify the network configuration of the Konga gateway instance
//
// error code that may be returned:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyConsoleNetwork(request *ModifyConsoleNetworkRequest) (response *ModifyConsoleNetworkResponse, err error) {
    return c.ModifyConsoleNetworkWithContext(context.Background(), request)
}

// ModifyConsoleNetwork
// Modify the network configuration of the Konga gateway instance
//
// error code that may be returned:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyConsoleNetworkWithContext(ctx context.Context, request *ModifyConsoleNetworkRequest) (response *ModifyConsoleNetworkResponse, err error) {
    if request == nil {
        request = NewModifyConsoleNetworkRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "ModifyConsoleNetwork")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyConsoleNetwork require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyConsoleNetworkResponse()
    err = c.Send(request, response)
    return
}

func NewModifyGovernanceLaneGroupsRequest() (request *ModifyGovernanceLaneGroupsRequest) {
    request = &ModifyGovernanceLaneGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "ModifyGovernanceLaneGroups")
    
    
    return
}

func NewModifyGovernanceLaneGroupsResponse() (response *ModifyGovernanceLaneGroupsResponse) {
    response = &ModifyGovernanceLaneGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyGovernanceLaneGroups
// Create a lane group
//
// error code that may be returned:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyGovernanceLaneGroups(request *ModifyGovernanceLaneGroupsRequest) (response *ModifyGovernanceLaneGroupsResponse, err error) {
    return c.ModifyGovernanceLaneGroupsWithContext(context.Background(), request)
}

// ModifyGovernanceLaneGroups
// Create a lane group
//
// error code that may be returned:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyGovernanceLaneGroupsWithContext(ctx context.Context, request *ModifyGovernanceLaneGroupsRequest) (response *ModifyGovernanceLaneGroupsResponse, err error) {
    if request == nil {
        request = NewModifyGovernanceLaneGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "ModifyGovernanceLaneGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyGovernanceLaneGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyGovernanceLaneGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNativeGatewayServerGroupRequest() (request *ModifyNativeGatewayServerGroupRequest) {
    request = &ModifyNativeGatewayServerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "ModifyNativeGatewayServerGroup")
    
    
    return
}

func NewModifyNativeGatewayServerGroupResponse() (response *ModifyNativeGatewayServerGroupResponse) {
    response = &ModifyNativeGatewayServerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyNativeGatewayServerGroup
// Modify the basic information of a cloud native API gateway instance group
//
// error code that may be returned:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyNativeGatewayServerGroup(request *ModifyNativeGatewayServerGroupRequest) (response *ModifyNativeGatewayServerGroupResponse, err error) {
    return c.ModifyNativeGatewayServerGroupWithContext(context.Background(), request)
}

// ModifyNativeGatewayServerGroup
// Modify the basic information of a cloud native API gateway instance group
//
// error code that may be returned:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyNativeGatewayServerGroupWithContext(ctx context.Context, request *ModifyNativeGatewayServerGroupRequest) (response *ModifyNativeGatewayServerGroupResponse, err error) {
    if request == nil {
        request = NewModifyNativeGatewayServerGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "ModifyNativeGatewayServerGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNativeGatewayServerGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNativeGatewayServerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNativeGatewayServiceSourceRequest() (request *ModifyNativeGatewayServiceSourceRequest) {
    request = &ModifyNativeGatewayServiceSourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "ModifyNativeGatewayServiceSource")
    
    
    return
}

func NewModifyNativeGatewayServiceSourceResponse() (response *ModifyNativeGatewayServiceSourceResponse) {
    response = &ModifyNativeGatewayServiceSourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyNativeGatewayServiceSource
// Modify the gateway service source
//
// error code that may be returned:
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) ModifyNativeGatewayServiceSource(request *ModifyNativeGatewayServiceSourceRequest) (response *ModifyNativeGatewayServiceSourceResponse, err error) {
    return c.ModifyNativeGatewayServiceSourceWithContext(context.Background(), request)
}

// ModifyNativeGatewayServiceSource
// Modify the gateway service source
//
// error code that may be returned:
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INVALIDPARAMETERVALUE_CREATEERROR = "InvalidParameterValue.CreateError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) ModifyNativeGatewayServiceSourceWithContext(ctx context.Context, request *ModifyNativeGatewayServiceSourceRequest) (response *ModifyNativeGatewayServiceSourceResponse, err error) {
    if request == nil {
        request = NewModifyNativeGatewayServiceSourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "ModifyNativeGatewayServiceSource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNativeGatewayServiceSource require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNativeGatewayServiceSourceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNetworkAccessStrategyRequest() (request *ModifyNetworkAccessStrategyRequest) {
    request = &ModifyNetworkAccessStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "ModifyNetworkAccessStrategy")
    
    
    return
}

func NewModifyNetworkAccessStrategyResponse() (response *ModifyNetworkAccessStrategyResponse) {
    response = &ModifyNetworkAccessStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyNetworkAccessStrategy
// Modify the access policy of the Kong cloud native API gateway instance to support allowlist or blocklist.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyNetworkAccessStrategy(request *ModifyNetworkAccessStrategyRequest) (response *ModifyNetworkAccessStrategyResponse, err error) {
    return c.ModifyNetworkAccessStrategyWithContext(context.Background(), request)
}

// ModifyNetworkAccessStrategy
// Modify the access policy of the Kong cloud native API gateway instance to support allowlist or blocklist.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyNetworkAccessStrategyWithContext(ctx context.Context, request *ModifyNetworkAccessStrategyRequest) (response *ModifyNetworkAccessStrategyResponse, err error) {
    if request == nil {
        request = NewModifyNetworkAccessStrategyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "ModifyNetworkAccessStrategy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNetworkAccessStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNetworkAccessStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNetworkBasicInfoRequest() (request *ModifyNetworkBasicInfoRequest) {
    request = &ModifyNetworkBasicInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "ModifyNetworkBasicInfo")
    
    
    return
}

func NewModifyNetworkBasicInfoResponse() (response *ModifyNetworkBasicInfoResponse) {
    response = &ModifyNetworkBasicInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyNetworkBasicInfo
// This API is used to modify the basic information of a cloud native API gateway instance network, such as bandwidth and description, as well as specification upgrade. Only modification of client public network or private network information is supported.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyNetworkBasicInfo(request *ModifyNetworkBasicInfoRequest) (response *ModifyNetworkBasicInfoResponse, err error) {
    return c.ModifyNetworkBasicInfoWithContext(context.Background(), request)
}

// ModifyNetworkBasicInfo
// This API is used to modify the basic information of a cloud native API gateway instance network, such as bandwidth and description, as well as specification upgrade. Only modification of client public network or private network information is supported.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyNetworkBasicInfoWithContext(ctx context.Context, request *ModifyNetworkBasicInfoRequest) (response *ModifyNetworkBasicInfoResponse, err error) {
    if request == nil {
        request = NewModifyNetworkBasicInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "ModifyNetworkBasicInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNetworkBasicInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNetworkBasicInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUpstreamNodeStatusRequest() (request *ModifyUpstreamNodeStatusRequest) {
    request = &ModifyUpstreamNodeStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "ModifyUpstreamNodeStatus")
    
    
    return
}

func NewModifyUpstreamNodeStatusResponse() (response *ModifyUpstreamNodeStatusResponse) {
    response = &ModifyUpstreamNodeStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyUpstreamNodeStatus
// Modify the node health status of the upstream instance for the cloud-native gateway
//
// error code that may be returned:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyUpstreamNodeStatus(request *ModifyUpstreamNodeStatusRequest) (response *ModifyUpstreamNodeStatusResponse, err error) {
    return c.ModifyUpstreamNodeStatusWithContext(context.Background(), request)
}

// ModifyUpstreamNodeStatus
// Modify the node health status of the upstream instance for the cloud-native gateway
//
// error code that may be returned:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_UPDATEERROR = "InvalidParameterValue.UpdateError"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyUpstreamNodeStatusWithContext(ctx context.Context, request *ModifyUpstreamNodeStatusRequest) (response *ModifyUpstreamNodeStatusResponse, err error) {
    if request == nil {
        request = NewModifyUpstreamNodeStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "ModifyUpstreamNodeStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUpstreamNodeStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUpstreamNodeStatusResponse()
    err = c.Send(request, response)
    return
}

func NewOpenWafProtectionRequest() (request *OpenWafProtectionRequest) {
    request = &OpenWafProtectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "OpenWafProtection")
    
    
    return
}

func NewOpenWafProtectionResponse() (response *OpenWafProtectionResponse) {
    response = &OpenWafProtectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OpenWafProtection
// Enable WAF protection
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) OpenWafProtection(request *OpenWafProtectionRequest) (response *OpenWafProtectionResponse, err error) {
    return c.OpenWafProtectionWithContext(context.Background(), request)
}

// OpenWafProtection
// Enable WAF protection
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) OpenWafProtectionWithContext(ctx context.Context, request *OpenWafProtectionRequest) (response *OpenWafProtectionResponse, err error) {
    if request == nil {
        request = NewOpenWafProtectionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "OpenWafProtection")
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenWafProtection require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenWafProtectionResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindAutoScalerResourceStrategyFromGroupsRequest() (request *UnbindAutoScalerResourceStrategyFromGroupsRequest) {
    request = &UnbindAutoScalerResourceStrategyFromGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "UnbindAutoScalerResourceStrategyFromGroups")
    
    
    return
}

func NewUnbindAutoScalerResourceStrategyFromGroupsResponse() (response *UnbindAutoScalerResourceStrategyFromGroupsResponse) {
    response = &UnbindAutoScalerResourceStrategyFromGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UnbindAutoScalerResourceStrategyFromGroups
// Unbind gateway groupings in batch with auto scaling policy
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) UnbindAutoScalerResourceStrategyFromGroups(request *UnbindAutoScalerResourceStrategyFromGroupsRequest) (response *UnbindAutoScalerResourceStrategyFromGroupsResponse, err error) {
    return c.UnbindAutoScalerResourceStrategyFromGroupsWithContext(context.Background(), request)
}

// UnbindAutoScalerResourceStrategyFromGroups
// Unbind gateway groupings in batch with auto scaling policy
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INTERNALERROR_OPERATIONFAILED = "InternalError.OperationFailed"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_GATEWAYID = "InvalidParameterValue.GatewayId"
//  INVALIDPARAMETERVALUE_QUERYERROR = "InvalidParameterValue.QueryError"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) UnbindAutoScalerResourceStrategyFromGroupsWithContext(ctx context.Context, request *UnbindAutoScalerResourceStrategyFromGroupsRequest) (response *UnbindAutoScalerResourceStrategyFromGroupsResponse, err error) {
    if request == nil {
        request = NewUnbindAutoScalerResourceStrategyFromGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "UnbindAutoScalerResourceStrategyFromGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindAutoScalerResourceStrategyFromGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindAutoScalerResourceStrategyFromGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCloudNativeAPIGatewayCertificateInfoRequest() (request *UpdateCloudNativeAPIGatewayCertificateInfoRequest) {
    request = &UpdateCloudNativeAPIGatewayCertificateInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "UpdateCloudNativeAPIGatewayCertificateInfo")
    
    
    return
}

func NewUpdateCloudNativeAPIGatewayCertificateInfoResponse() (response *UpdateCloudNativeAPIGatewayCertificateInfoResponse) {
    response = &UpdateCloudNativeAPIGatewayCertificateInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateCloudNativeAPIGatewayCertificateInfo
// Modify the certificate information of a cloud-native gateway
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  LIMITEXCEEDED_LBDOMAINS = "LimitExceeded.LBDomains"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) UpdateCloudNativeAPIGatewayCertificateInfo(request *UpdateCloudNativeAPIGatewayCertificateInfoRequest) (response *UpdateCloudNativeAPIGatewayCertificateInfoResponse, err error) {
    return c.UpdateCloudNativeAPIGatewayCertificateInfoWithContext(context.Background(), request)
}

// UpdateCloudNativeAPIGatewayCertificateInfo
// Modify the certificate information of a cloud-native gateway
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_ROLE = "FailedOperation.Role"
//  FAILEDOPERATION_VPC = "FailedOperation.Vpc"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYERROR = "InternalError.QueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_BADREQUESTFORMAT = "InvalidParameterValue.BadRequestFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OPERATIONFAILED = "InvalidParameterValue.OperationFailed"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXIST = "InvalidParameterValue.ResourceAlreadyExist"
//  INVALIDPARAMETERVALUE_SPECIFICATION = "InvalidParameterValue.Specification"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  LIMITEXCEEDED_LBDOMAINS = "LimitExceeded.LBDomains"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CAMPASSROLENOTEXIST = "UnauthorizedOperation.CamPassRoleNotExist"
//  UNAUTHORIZEDOPERATION_UIN = "UnauthorizedOperation.Uin"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) UpdateCloudNativeAPIGatewayCertificateInfoWithContext(ctx context.Context, request *UpdateCloudNativeAPIGatewayCertificateInfoRequest) (response *UpdateCloudNativeAPIGatewayCertificateInfoResponse, err error) {
    if request == nil {
        request = NewUpdateCloudNativeAPIGatewayCertificateInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "UpdateCloudNativeAPIGatewayCertificateInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateCloudNativeAPIGatewayCertificateInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateCloudNativeAPIGatewayCertificateInfoResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCloudNativeAPIGatewaySpecRequest() (request *UpdateCloudNativeAPIGatewaySpecRequest) {
    request = &UpdateCloudNativeAPIGatewaySpecRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "UpdateCloudNativeAPIGatewaySpec")
    
    
    return
}

func NewUpdateCloudNativeAPIGatewaySpecResponse() (response *UpdateCloudNativeAPIGatewaySpecResponse) {
    response = &UpdateCloudNativeAPIGatewaySpecResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateCloudNativeAPIGatewaySpec
// Modify the node specification information of a cloud native API gateway instance, such as node scaling or specification adjustment
//
// error code that may be returned:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) UpdateCloudNativeAPIGatewaySpec(request *UpdateCloudNativeAPIGatewaySpecRequest) (response *UpdateCloudNativeAPIGatewaySpecResponse, err error) {
    return c.UpdateCloudNativeAPIGatewaySpecWithContext(context.Background(), request)
}

// UpdateCloudNativeAPIGatewaySpec
// Modify the node specification information of a cloud native API gateway instance, such as node scaling or specification adjustment
//
// error code that may be returned:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSPARAMETER = "MissingParameter.MissParameter"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_FORBIDDEN = "ResourceNotFound.Forbidden"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) UpdateCloudNativeAPIGatewaySpecWithContext(ctx context.Context, request *UpdateCloudNativeAPIGatewaySpecRequest) (response *UpdateCloudNativeAPIGatewaySpecResponse, err error) {
    if request == nil {
        request = NewUpdateCloudNativeAPIGatewaySpecRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "UpdateCloudNativeAPIGatewaySpec")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateCloudNativeAPIGatewaySpec require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateCloudNativeAPIGatewaySpecResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateUpstreamHealthCheckConfigRequest() (request *UpdateUpstreamHealthCheckConfigRequest) {
    request = &UpdateUpstreamHealthCheckConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "UpdateUpstreamHealthCheckConfig")
    
    
    return
}

func NewUpdateUpstreamHealthCheckConfigResponse() (response *UpdateUpstreamHealthCheckConfigResponse) {
    response = &UpdateUpstreamHealthCheckConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateUpstreamHealthCheckConfig
// This API is used to update the health check configuration of the cloud-native gateway.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
func (c *Client) UpdateUpstreamHealthCheckConfig(request *UpdateUpstreamHealthCheckConfigRequest) (response *UpdateUpstreamHealthCheckConfigResponse, err error) {
    return c.UpdateUpstreamHealthCheckConfigWithContext(context.Background(), request)
}

// UpdateUpstreamHealthCheckConfig
// This API is used to update the health check configuration of the cloud-native gateway.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
func (c *Client) UpdateUpstreamHealthCheckConfigWithContext(ctx context.Context, request *UpdateUpstreamHealthCheckConfigRequest) (response *UpdateUpstreamHealthCheckConfigResponse, err error) {
    if request == nil {
        request = NewUpdateUpstreamHealthCheckConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "UpdateUpstreamHealthCheckConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateUpstreamHealthCheckConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateUpstreamHealthCheckConfigResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateUpstreamTargetsRequest() (request *UpdateUpstreamTargetsRequest) {
    request = &UpdateUpstreamTargetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tse", APIVersion, "UpdateUpstreamTargets")
    
    
    return
}

func NewUpdateUpstreamTargetsResponse() (response *UpdateUpstreamTargetsResponse) {
    response = &UpdateUpstreamTargetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateUpstreamTargets
// Refresh the upstream instance list of the gateway, only supported for the IPList service type
//
// error code that may be returned:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
func (c *Client) UpdateUpstreamTargets(request *UpdateUpstreamTargetsRequest) (response *UpdateUpstreamTargetsResponse, err error) {
    return c.UpdateUpstreamTargetsWithContext(context.Background(), request)
}

// UpdateUpstreamTargets
// Refresh the upstream instance list of the gateway, only supported for the IPList service type
//
// error code that may be returned:
//  FAILEDOPERATION_FAILEDOPERATION = "FailedOperation.FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
func (c *Client) UpdateUpstreamTargetsWithContext(ctx context.Context, request *UpdateUpstreamTargetsRequest) (response *UpdateUpstreamTargetsResponse, err error) {
    if request == nil {
        request = NewUpdateUpstreamTargetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tse", APIVersion, "UpdateUpstreamTargets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateUpstreamTargets require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateUpstreamTargetsResponse()
    err = c.Send(request, response)
    return
}
