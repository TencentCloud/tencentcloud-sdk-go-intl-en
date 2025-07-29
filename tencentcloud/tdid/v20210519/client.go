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

package v20210519

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2021-05-19"

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


func NewAddLabelRequest() (request *AddLabelRequest) {
    request = &AddLabelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "AddLabel")
    
    
    return
}

func NewAddLabelResponse() (response *AddLabelResponse) {
    response = &AddLabelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddLabel
// This API is used to add a label to a DID.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AddLabel(request *AddLabelRequest) (response *AddLabelResponse, err error) {
    return c.AddLabelWithContext(context.Background(), request)
}

// AddLabel
// This API is used to add a label to a DID.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AddLabelWithContext(ctx context.Context, request *AddLabelRequest) (response *AddLabelResponse, err error) {
    if request == nil {
        request = NewAddLabelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "AddLabel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddLabel require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddLabelResponse()
    err = c.Send(request, response)
    return
}

func NewCancelAuthorityIssuerRequest() (request *CancelAuthorityIssuerRequest) {
    request = &CancelAuthorityIssuerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "CancelAuthorityIssuer")
    
    
    return
}

func NewCancelAuthorityIssuerResponse() (response *CancelAuthorityIssuerResponse) {
    response = &CancelAuthorityIssuerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CancelAuthorityIssuer
// This API is used to revoke the certification of an issuing authority.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CancelAuthorityIssuer(request *CancelAuthorityIssuerRequest) (response *CancelAuthorityIssuerResponse, err error) {
    return c.CancelAuthorityIssuerWithContext(context.Background(), request)
}

// CancelAuthorityIssuer
// This API is used to revoke the certification of an issuing authority.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CancelAuthorityIssuerWithContext(ctx context.Context, request *CancelAuthorityIssuerRequest) (response *CancelAuthorityIssuerResponse, err error) {
    if request == nil {
        request = NewCancelAuthorityIssuerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "CancelAuthorityIssuer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelAuthorityIssuer require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelAuthorityIssuerResponse()
    err = c.Send(request, response)
    return
}

func NewCheckChainRequest() (request *CheckChainRequest) {
    request = &CheckChainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "CheckChain")
    
    
    return
}

func NewCheckChainResponse() (response *CheckChainResponse) {
    response = &CheckChainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckChain
// This API is used to get blockchain information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CheckChain(request *CheckChainRequest) (response *CheckChainResponse, err error) {
    return c.CheckChainWithContext(context.Background(), request)
}

// CheckChain
// This API is used to get blockchain information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CheckChainWithContext(ctx context.Context, request *CheckChainRequest) (response *CheckChainResponse, err error) {
    if request == nil {
        request = NewCheckChainRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "CheckChain")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckChain require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckChainResponse()
    err = c.Send(request, response)
    return
}

func NewCheckDidDeployRequest() (request *CheckDidDeployRequest) {
    request = &CheckDidDeployRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "CheckDidDeploy")
    
    
    return
}

func NewCheckDidDeployResponse() (response *CheckDidDeployResponse) {
    response = &CheckDidDeployResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckDidDeploy
// This API is used to query a deployment task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CheckDidDeploy(request *CheckDidDeployRequest) (response *CheckDidDeployResponse, err error) {
    return c.CheckDidDeployWithContext(context.Background(), request)
}

// CheckDidDeploy
// This API is used to query a deployment task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CheckDidDeployWithContext(ctx context.Context, request *CheckDidDeployRequest) (response *CheckDidDeployResponse, err error) {
    if request == nil {
        request = NewCheckDidDeployRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "CheckDidDeploy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckDidDeploy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckDidDeployResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCredentialRequest() (request *CreateCredentialRequest) {
    request = &CreateCredentialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "CreateCredential")
    
    
    return
}

func NewCreateCredentialResponse() (response *CreateCredentialResponse) {
    response = &CreateCredentialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCredential
// This API is used to create a credential.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCredential(request *CreateCredentialRequest) (response *CreateCredentialResponse, err error) {
    return c.CreateCredentialWithContext(context.Background(), request)
}

// CreateCredential
// This API is used to create a credential.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCredentialWithContext(ctx context.Context, request *CreateCredentialRequest) (response *CreateCredentialResponse, err error) {
    if request == nil {
        request = NewCreateCredentialRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "CreateCredential")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCredential require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCredentialResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDidServiceRequest() (request *CreateDidServiceRequest) {
    request = &CreateDidServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "CreateDidService")
    
    
    return
}

func NewCreateDidServiceResponse() (response *CreateDidServiceResponse) {
    response = &CreateDidServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDidService
// This API is used to create a DID service.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDidService(request *CreateDidServiceRequest) (response *CreateDidServiceResponse, err error) {
    return c.CreateDidServiceWithContext(context.Background(), request)
}

// CreateDidService
// This API is used to create a DID service.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDidServiceWithContext(ctx context.Context, request *CreateDidServiceRequest) (response *CreateDidServiceResponse, err error) {
    if request == nil {
        request = NewCreateDidServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "CreateDidService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDidService require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDidServiceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLabelRequest() (request *CreateLabelRequest) {
    request = &CreateLabelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "CreateLabel")
    
    
    return
}

func NewCreateLabelResponse() (response *CreateLabelResponse) {
    response = &CreateLabelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLabel
// This API is used to create a label.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateLabel(request *CreateLabelRequest) (response *CreateLabelResponse, err error) {
    return c.CreateLabelWithContext(context.Background(), request)
}

// CreateLabel
// This API is used to create a label.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateLabelWithContext(ctx context.Context, request *CreateLabelRequest) (response *CreateLabelResponse, err error) {
    if request == nil {
        request = NewCreateLabelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "CreateLabel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLabel require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLabelResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSelectiveCredentialRequest() (request *CreateSelectiveCredentialRequest) {
    request = &CreateSelectiveCredentialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "CreateSelectiveCredential")
    
    
    return
}

func NewCreateSelectiveCredentialResponse() (response *CreateSelectiveCredentialResponse) {
    response = &CreateSelectiveCredentialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSelectiveCredential
// This API is used to create a selective disclosure credential.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSelectiveCredential(request *CreateSelectiveCredentialRequest) (response *CreateSelectiveCredentialResponse, err error) {
    return c.CreateSelectiveCredentialWithContext(context.Background(), request)
}

// CreateSelectiveCredential
// This API is used to create a selective disclosure credential.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSelectiveCredentialWithContext(ctx context.Context, request *CreateSelectiveCredentialRequest) (response *CreateSelectiveCredentialResponse, err error) {
    if request == nil {
        request = NewCreateSelectiveCredentialRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "CreateSelectiveCredential")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSelectiveCredential require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSelectiveCredentialResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTDidRequest() (request *CreateTDidRequest) {
    request = &CreateTDidRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "CreateTDid")
    
    
    return
}

func NewCreateTDidResponse() (response *CreateTDidResponse) {
    response = &CreateTDidResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTDid
// This API is used to create an organization DID.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
func (c *Client) CreateTDid(request *CreateTDidRequest) (response *CreateTDidResponse, err error) {
    return c.CreateTDidWithContext(context.Background(), request)
}

// CreateTDid
// This API is used to create an organization DID.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
func (c *Client) CreateTDidWithContext(ctx context.Context, request *CreateTDidRequest) (response *CreateTDidResponse, err error) {
    if request == nil {
        request = NewCreateTDidRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "CreateTDid")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTDid require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTDidResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTDidByPrivateKeyRequest() (request *CreateTDidByPrivateKeyRequest) {
    request = &CreateTDidByPrivateKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "CreateTDidByPrivateKey")
    
    
    return
}

func NewCreateTDidByPrivateKeyResponse() (response *CreateTDidByPrivateKeyResponse) {
    response = &CreateTDidByPrivateKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTDidByPrivateKey
// This API is used to generate a TDID by private key.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateTDidByPrivateKey(request *CreateTDidByPrivateKeyRequest) (response *CreateTDidByPrivateKeyResponse, err error) {
    return c.CreateTDidByPrivateKeyWithContext(context.Background(), request)
}

// CreateTDidByPrivateKey
// This API is used to generate a TDID by private key.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateTDidByPrivateKeyWithContext(ctx context.Context, request *CreateTDidByPrivateKeyRequest) (response *CreateTDidByPrivateKeyResponse, err error) {
    if request == nil {
        request = NewCreateTDidByPrivateKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "CreateTDidByPrivateKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTDidByPrivateKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTDidByPrivateKeyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTDidByPublicKeyRequest() (request *CreateTDidByPublicKeyRequest) {
    request = &CreateTDidByPublicKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "CreateTDidByPublicKey")
    
    
    return
}

func NewCreateTDidByPublicKeyResponse() (response *CreateTDidByPublicKeyResponse) {
    response = &CreateTDidByPublicKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTDidByPublicKey
//  This API is used to generate a TDID by public key.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateTDidByPublicKey(request *CreateTDidByPublicKeyRequest) (response *CreateTDidByPublicKeyResponse, err error) {
    return c.CreateTDidByPublicKeyWithContext(context.Background(), request)
}

// CreateTDidByPublicKey
//  This API is used to generate a TDID by public key.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateTDidByPublicKeyWithContext(ctx context.Context, request *CreateTDidByPublicKeyRequest) (response *CreateTDidByPublicKeyResponse, err error) {
    if request == nil {
        request = NewCreateTDidByPublicKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "CreateTDidByPublicKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTDidByPublicKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTDidByPublicKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDeployByNameRequest() (request *DeployByNameRequest) {
    request = &DeployByNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "DeployByName")
    
    
    return
}

func NewDeployByNameResponse() (response *DeployByNameResponse) {
    response = &DeployByNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeployByName
// This API is used to deploy a TDID contract by name.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeployByName(request *DeployByNameRequest) (response *DeployByNameResponse, err error) {
    return c.DeployByNameWithContext(context.Background(), request)
}

// DeployByName
// This API is used to deploy a TDID contract by name.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeployByNameWithContext(ctx context.Context, request *DeployByNameRequest) (response *DeployByNameResponse, err error) {
    if request == nil {
        request = NewDeployByNameRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "DeployByName")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeployByName require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeployByNameResponse()
    err = c.Send(request, response)
    return
}

func NewDownCptRequest() (request *DownCptRequest) {
    request = &DownCptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "DownCpt")
    
    
    return
}

func NewDownCptResponse() (response *DownCptResponse) {
    response = &DownCptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DownCpt
// This API is used to download a claim protocol type (CPT).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DownCpt(request *DownCptRequest) (response *DownCptResponse, err error) {
    return c.DownCptWithContext(context.Background(), request)
}

// DownCpt
// This API is used to download a claim protocol type (CPT).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DownCptWithContext(ctx context.Context, request *DownCptRequest) (response *DownCptResponse, err error) {
    if request == nil {
        request = NewDownCptRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "DownCpt")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DownCpt require credential")
    }

    request.SetContext(ctx)
    
    response = NewDownCptResponse()
    err = c.Send(request, response)
    return
}

func NewEnableHashRequest() (request *EnableHashRequest) {
    request = &EnableHashRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "EnableHash")
    
    
    return
}

func NewEnableHashResponse() (response *EnableHashResponse) {
    response = &EnableHashResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnableHash
// This API is used to enable a contract.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) EnableHash(request *EnableHashRequest) (response *EnableHashResponse, err error) {
    return c.EnableHashWithContext(context.Background(), request)
}

// EnableHash
// This API is used to enable a contract.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) EnableHashWithContext(ctx context.Context, request *EnableHashRequest) (response *EnableHashResponse, err error) {
    if request == nil {
        request = NewEnableHashRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "EnableHash")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableHash require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableHashResponse()
    err = c.Send(request, response)
    return
}

func NewGetAgencyTDidRequest() (request *GetAgencyTDidRequest) {
    request = &GetAgencyTDidRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "GetAgencyTDid")
    
    
    return
}

func NewGetAgencyTDidResponse() (response *GetAgencyTDidResponse) {
    response = &GetAgencyTDidResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetAgencyTDid
// 该接口已废弃
//
// 
//
// This API is used to get the DID details of the current organization.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetAgencyTDid(request *GetAgencyTDidRequest) (response *GetAgencyTDidResponse, err error) {
    return c.GetAgencyTDidWithContext(context.Background(), request)
}

// GetAgencyTDid
// 该接口已废弃
//
// 
//
// This API is used to get the DID details of the current organization.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetAgencyTDidWithContext(ctx context.Context, request *GetAgencyTDidRequest) (response *GetAgencyTDidResponse, err error) {
    if request == nil {
        request = NewGetAgencyTDidRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "GetAgencyTDid")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetAgencyTDid require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetAgencyTDidResponse()
    err = c.Send(request, response)
    return
}

func NewGetAuthoritiesListRequest() (request *GetAuthoritiesListRequest) {
    request = &GetAuthoritiesListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "GetAuthoritiesList")
    
    
    return
}

func NewGetAuthoritiesListResponse() (response *GetAuthoritiesListResponse) {
    response = &GetAuthoritiesListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetAuthoritiesList
// This API is used to query authorities.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetAuthoritiesList(request *GetAuthoritiesListRequest) (response *GetAuthoritiesListResponse, err error) {
    return c.GetAuthoritiesListWithContext(context.Background(), request)
}

// GetAuthoritiesList
// This API is used to query authorities.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetAuthoritiesListWithContext(ctx context.Context, request *GetAuthoritiesListRequest) (response *GetAuthoritiesListResponse, err error) {
    if request == nil {
        request = NewGetAuthoritiesListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "GetAuthoritiesList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetAuthoritiesList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetAuthoritiesListResponse()
    err = c.Send(request, response)
    return
}

func NewGetAuthorityIssuerRequest() (request *GetAuthorityIssuerRequest) {
    request = &GetAuthorityIssuerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "GetAuthorityIssuer")
    
    
    return
}

func NewGetAuthorityIssuerResponse() (response *GetAuthorityIssuerResponse) {
    response = &GetAuthorityIssuerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetAuthorityIssuer
// This API is used to get the information of an issuing authority.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetAuthorityIssuer(request *GetAuthorityIssuerRequest) (response *GetAuthorityIssuerResponse, err error) {
    return c.GetAuthorityIssuerWithContext(context.Background(), request)
}

// GetAuthorityIssuer
// This API is used to get the information of an issuing authority.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetAuthorityIssuerWithContext(ctx context.Context, request *GetAuthorityIssuerRequest) (response *GetAuthorityIssuerResponse, err error) {
    if request == nil {
        request = NewGetAuthorityIssuerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "GetAuthorityIssuer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetAuthorityIssuer require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetAuthorityIssuerResponse()
    err = c.Send(request, response)
    return
}

func NewGetConsortiumClusterListRequest() (request *GetConsortiumClusterListRequest) {
    request = &GetConsortiumClusterListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "GetConsortiumClusterList")
    
    
    return
}

func NewGetConsortiumClusterListResponse() (response *GetConsortiumClusterListResponse) {
    response = &GetConsortiumClusterListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetConsortiumClusterList
// This API is used to query the BCOS networks of a consortium.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetConsortiumClusterList(request *GetConsortiumClusterListRequest) (response *GetConsortiumClusterListResponse, err error) {
    return c.GetConsortiumClusterListWithContext(context.Background(), request)
}

// GetConsortiumClusterList
// This API is used to query the BCOS networks of a consortium.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetConsortiumClusterListWithContext(ctx context.Context, request *GetConsortiumClusterListRequest) (response *GetConsortiumClusterListResponse, err error) {
    if request == nil {
        request = NewGetConsortiumClusterListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "GetConsortiumClusterList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetConsortiumClusterList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetConsortiumClusterListResponse()
    err = c.Send(request, response)
    return
}

func NewGetConsortiumListRequest() (request *GetConsortiumListRequest) {
    request = &GetConsortiumListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "GetConsortiumList")
    
    
    return
}

func NewGetConsortiumListResponse() (response *GetConsortiumListResponse) {
    response = &GetConsortiumListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetConsortiumList
// This API is used to query consortiums.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetConsortiumList(request *GetConsortiumListRequest) (response *GetConsortiumListResponse, err error) {
    return c.GetConsortiumListWithContext(context.Background(), request)
}

// GetConsortiumList
// This API is used to query consortiums.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetConsortiumListWithContext(ctx context.Context, request *GetConsortiumListRequest) (response *GetConsortiumListResponse, err error) {
    if request == nil {
        request = NewGetConsortiumListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "GetConsortiumList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetConsortiumList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetConsortiumListResponse()
    err = c.Send(request, response)
    return
}

func NewGetCptInfoRequest() (request *GetCptInfoRequest) {
    request = &GetCptInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "GetCptInfo")
    
    
    return
}

func NewGetCptInfoResponse() (response *GetCptInfoResponse) {
    response = &GetCptInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetCptInfo
// This API is used to get the information of a claim protocol type (CPT).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetCptInfo(request *GetCptInfoRequest) (response *GetCptInfoResponse, err error) {
    return c.GetCptInfoWithContext(context.Background(), request)
}

// GetCptInfo
// This API is used to get the information of a claim protocol type (CPT).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetCptInfoWithContext(ctx context.Context, request *GetCptInfoRequest) (response *GetCptInfoResponse, err error) {
    if request == nil {
        request = NewGetCptInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "GetCptInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetCptInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetCptInfoResponse()
    err = c.Send(request, response)
    return
}

func NewGetCptListRequest() (request *GetCptListRequest) {
    request = &GetCptListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "GetCptList")
    
    
    return
}

func NewGetCptListResponse() (response *GetCptListResponse) {
    response = &GetCptListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetCptList
// This API is used to query claim protocol types (CPT).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetCptList(request *GetCptListRequest) (response *GetCptListResponse, err error) {
    return c.GetCptListWithContext(context.Background(), request)
}

// GetCptList
// This API is used to query claim protocol types (CPT).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetCptListWithContext(ctx context.Context, request *GetCptListRequest) (response *GetCptListResponse, err error) {
    if request == nil {
        request = NewGetCptListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "GetCptList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetCptList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetCptListResponse()
    err = c.Send(request, response)
    return
}

func NewGetCredentialCptRankRequest() (request *GetCredentialCptRankRequest) {
    request = &GetCredentialCptRankRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "GetCredentialCptRank")
    
    
    return
}

func NewGetCredentialCptRankResponse() (response *GetCredentialCptRankResponse) {
    response = &GetCredentialCptRankResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetCredentialCptRank
// This API is used to get the rankings of claim protocol types (CPT).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetCredentialCptRank(request *GetCredentialCptRankRequest) (response *GetCredentialCptRankResponse, err error) {
    return c.GetCredentialCptRankWithContext(context.Background(), request)
}

// GetCredentialCptRank
// This API is used to get the rankings of claim protocol types (CPT).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetCredentialCptRankWithContext(ctx context.Context, request *GetCredentialCptRankRequest) (response *GetCredentialCptRankResponse, err error) {
    if request == nil {
        request = NewGetCredentialCptRankRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "GetCredentialCptRank")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetCredentialCptRank require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetCredentialCptRankResponse()
    err = c.Send(request, response)
    return
}

func NewGetCredentialIssueRankRequest() (request *GetCredentialIssueRankRequest) {
    request = &GetCredentialIssueRankRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "GetCredentialIssueRank")
    
    
    return
}

func NewGetCredentialIssueRankResponse() (response *GetCredentialIssueRankResponse) {
    response = &GetCredentialIssueRankResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetCredentialIssueRank
// This API is used to get the rankings of issuers.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetCredentialIssueRank(request *GetCredentialIssueRankRequest) (response *GetCredentialIssueRankResponse, err error) {
    return c.GetCredentialIssueRankWithContext(context.Background(), request)
}

// GetCredentialIssueRank
// This API is used to get the rankings of issuers.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetCredentialIssueRankWithContext(ctx context.Context, request *GetCredentialIssueRankRequest) (response *GetCredentialIssueRankResponse, err error) {
    if request == nil {
        request = NewGetCredentialIssueRankRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "GetCredentialIssueRank")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetCredentialIssueRank require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetCredentialIssueRankResponse()
    err = c.Send(request, response)
    return
}

func NewGetCredentialIssueTrendRequest() (request *GetCredentialIssueTrendRequest) {
    request = &GetCredentialIssueTrendRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "GetCredentialIssueTrend")
    
    
    return
}

func NewGetCredentialIssueTrendResponse() (response *GetCredentialIssueTrendResponse) {
    response = &GetCredentialIssueTrendResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetCredentialIssueTrend
// This API is used to query credential issuing trends.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetCredentialIssueTrend(request *GetCredentialIssueTrendRequest) (response *GetCredentialIssueTrendResponse, err error) {
    return c.GetCredentialIssueTrendWithContext(context.Background(), request)
}

// GetCredentialIssueTrend
// This API is used to query credential issuing trends.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetCredentialIssueTrendWithContext(ctx context.Context, request *GetCredentialIssueTrendRequest) (response *GetCredentialIssueTrendResponse, err error) {
    if request == nil {
        request = NewGetCredentialIssueTrendRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "GetCredentialIssueTrend")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetCredentialIssueTrend require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetCredentialIssueTrendResponse()
    err = c.Send(request, response)
    return
}

func NewGetCredentialStatusRequest() (request *GetCredentialStatusRequest) {
    request = &GetCredentialStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "GetCredentialStatus")
    
    
    return
}

func NewGetCredentialStatusResponse() (response *GetCredentialStatusResponse) {
    response = &GetCredentialStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetCredentialStatus
// This API is used to query the on-chain status of a credential.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetCredentialStatus(request *GetCredentialStatusRequest) (response *GetCredentialStatusResponse, err error) {
    return c.GetCredentialStatusWithContext(context.Background(), request)
}

// GetCredentialStatus
// This API is used to query the on-chain status of a credential.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetCredentialStatusWithContext(ctx context.Context, request *GetCredentialStatusRequest) (response *GetCredentialStatusResponse, err error) {
    if request == nil {
        request = NewGetCredentialStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "GetCredentialStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetCredentialStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetCredentialStatusResponse()
    err = c.Send(request, response)
    return
}

func NewGetDataPanelRequest() (request *GetDataPanelRequest) {
    request = &GetDataPanelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "GetDataPanel")
    
    
    return
}

func NewGetDataPanelResponse() (response *GetDataPanelResponse) {
    response = &GetDataPanelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetDataPanel
// This API is used to query the overall statistics of a network.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetDataPanel(request *GetDataPanelRequest) (response *GetDataPanelResponse, err error) {
    return c.GetDataPanelWithContext(context.Background(), request)
}

// GetDataPanel
// This API is used to query the overall statistics of a network.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetDataPanelWithContext(ctx context.Context, request *GetDataPanelRequest) (response *GetDataPanelResponse, err error) {
    if request == nil {
        request = NewGetDataPanelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "GetDataPanel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDataPanel require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDataPanelResponse()
    err = c.Send(request, response)
    return
}

func NewGetDeployInfoRequest() (request *GetDeployInfoRequest) {
    request = &GetDeployInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "GetDeployInfo")
    
    
    return
}

func NewGetDeployInfoResponse() (response *GetDeployInfoResponse) {
    response = &GetDeployInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetDeployInfo
// This API is used to query the deployment information of a contract.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetDeployInfo(request *GetDeployInfoRequest) (response *GetDeployInfoResponse, err error) {
    return c.GetDeployInfoWithContext(context.Background(), request)
}

// GetDeployInfo
// This API is used to query the deployment information of a contract.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetDeployInfoWithContext(ctx context.Context, request *GetDeployInfoRequest) (response *GetDeployInfoResponse, err error) {
    if request == nil {
        request = NewGetDeployInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "GetDeployInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDeployInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDeployInfoResponse()
    err = c.Send(request, response)
    return
}

func NewGetDeployListRequest() (request *GetDeployListRequest) {
    request = &GetDeployListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "GetDeployList")
    
    
    return
}

func NewGetDeployListResponse() (response *GetDeployListResponse) {
    response = &GetDeployListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetDeployList
// This API is used to query deployed contracts.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetDeployList(request *GetDeployListRequest) (response *GetDeployListResponse, err error) {
    return c.GetDeployListWithContext(context.Background(), request)
}

// GetDeployList
// This API is used to query deployed contracts.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetDeployListWithContext(ctx context.Context, request *GetDeployListRequest) (response *GetDeployListResponse, err error) {
    if request == nil {
        request = NewGetDeployListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "GetDeployList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDeployList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDeployListResponse()
    err = c.Send(request, response)
    return
}

func NewGetDidClusterDetailRequest() (request *GetDidClusterDetailRequest) {
    request = &GetDidClusterDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "GetDidClusterDetail")
    
    
    return
}

func NewGetDidClusterDetailResponse() (response *GetDidClusterDetailResponse) {
    response = &GetDidClusterDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetDidClusterDetail
// This API is used to get the information of a DID blockchain network.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetDidClusterDetail(request *GetDidClusterDetailRequest) (response *GetDidClusterDetailResponse, err error) {
    return c.GetDidClusterDetailWithContext(context.Background(), request)
}

// GetDidClusterDetail
// This API is used to get the information of a DID blockchain network.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetDidClusterDetailWithContext(ctx context.Context, request *GetDidClusterDetailRequest) (response *GetDidClusterDetailResponse, err error) {
    if request == nil {
        request = NewGetDidClusterDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "GetDidClusterDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDidClusterDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDidClusterDetailResponse()
    err = c.Send(request, response)
    return
}

func NewGetDidClusterListRequest() (request *GetDidClusterListRequest) {
    request = &GetDidClusterListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "GetDidClusterList")
    
    
    return
}

func NewGetDidClusterListResponse() (response *GetDidClusterListResponse) {
    response = &GetDidClusterListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetDidClusterList
// This API is used to query your DID networks.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetDidClusterList(request *GetDidClusterListRequest) (response *GetDidClusterListResponse, err error) {
    return c.GetDidClusterListWithContext(context.Background(), request)
}

// GetDidClusterList
// This API is used to query your DID networks.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetDidClusterListWithContext(ctx context.Context, request *GetDidClusterListRequest) (response *GetDidClusterListResponse, err error) {
    if request == nil {
        request = NewGetDidClusterListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "GetDidClusterList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDidClusterList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDidClusterListResponse()
    err = c.Send(request, response)
    return
}

func NewGetDidDetailRequest() (request *GetDidDetailRequest) {
    request = &GetDidDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "GetDidDetail")
    
    
    return
}

func NewGetDidDetailResponse() (response *GetDidDetailResponse) {
    response = &GetDidDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetDidDetail
// This API is used to get the information of a DID.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetDidDetail(request *GetDidDetailRequest) (response *GetDidDetailResponse, err error) {
    return c.GetDidDetailWithContext(context.Background(), request)
}

// GetDidDetail
// This API is used to get the information of a DID.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetDidDetailWithContext(ctx context.Context, request *GetDidDetailRequest) (response *GetDidDetailResponse, err error) {
    if request == nil {
        request = NewGetDidDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "GetDidDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDidDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDidDetailResponse()
    err = c.Send(request, response)
    return
}

func NewGetDidDocumentRequest() (request *GetDidDocumentRequest) {
    request = &GetDidDocumentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "GetDidDocument")
    
    
    return
}

func NewGetDidDocumentResponse() (response *GetDidDocumentResponse) {
    response = &GetDidDocumentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetDidDocument
// This API is used to get the document of a DID.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetDidDocument(request *GetDidDocumentRequest) (response *GetDidDocumentResponse, err error) {
    return c.GetDidDocumentWithContext(context.Background(), request)
}

// GetDidDocument
// This API is used to get the document of a DID.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetDidDocumentWithContext(ctx context.Context, request *GetDidDocumentRequest) (response *GetDidDocumentResponse, err error) {
    if request == nil {
        request = NewGetDidDocumentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "GetDidDocument")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDidDocument require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDidDocumentResponse()
    err = c.Send(request, response)
    return
}

func NewGetDidListRequest() (request *GetDidListRequest) {
    request = &GetDidListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "GetDidList")
    
    
    return
}

func NewGetDidListResponse() (response *GetDidListResponse) {
    response = &GetDidListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetDidList
// This API is used to query DIDs.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetDidList(request *GetDidListRequest) (response *GetDidListResponse, err error) {
    return c.GetDidListWithContext(context.Background(), request)
}

// GetDidList
// This API is used to query DIDs.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetDidListWithContext(ctx context.Context, request *GetDidListRequest) (response *GetDidListResponse, err error) {
    if request == nil {
        request = NewGetDidListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "GetDidList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDidList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDidListResponse()
    err = c.Send(request, response)
    return
}

func NewGetDidRegisterTrendRequest() (request *GetDidRegisterTrendRequest) {
    request = &GetDidRegisterTrendRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "GetDidRegisterTrend")
    
    
    return
}

func NewGetDidRegisterTrendResponse() (response *GetDidRegisterTrendResponse) {
    response = &GetDidRegisterTrendResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetDidRegisterTrend
// This API is used to query DID registration trends.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetDidRegisterTrend(request *GetDidRegisterTrendRequest) (response *GetDidRegisterTrendResponse, err error) {
    return c.GetDidRegisterTrendWithContext(context.Background(), request)
}

// GetDidRegisterTrend
// This API is used to query DID registration trends.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetDidRegisterTrendWithContext(ctx context.Context, request *GetDidRegisterTrendRequest) (response *GetDidRegisterTrendResponse, err error) {
    if request == nil {
        request = NewGetDidRegisterTrendRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "GetDidRegisterTrend")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDidRegisterTrend require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDidRegisterTrendResponse()
    err = c.Send(request, response)
    return
}

func NewGetDidServiceDetailRequest() (request *GetDidServiceDetailRequest) {
    request = &GetDidServiceDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "GetDidServiceDetail")
    
    
    return
}

func NewGetDidServiceDetailResponse() (response *GetDidServiceDetailResponse) {
    response = &GetDidServiceDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetDidServiceDetail
// This API is used to get the information of a DID service.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetDidServiceDetail(request *GetDidServiceDetailRequest) (response *GetDidServiceDetailResponse, err error) {
    return c.GetDidServiceDetailWithContext(context.Background(), request)
}

// GetDidServiceDetail
// This API is used to get the information of a DID service.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetDidServiceDetailWithContext(ctx context.Context, request *GetDidServiceDetailRequest) (response *GetDidServiceDetailResponse, err error) {
    if request == nil {
        request = NewGetDidServiceDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "GetDidServiceDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDidServiceDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDidServiceDetailResponse()
    err = c.Send(request, response)
    return
}

func NewGetDidServiceListRequest() (request *GetDidServiceListRequest) {
    request = &GetDidServiceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "GetDidServiceList")
    
    
    return
}

func NewGetDidServiceListResponse() (response *GetDidServiceListResponse) {
    response = &GetDidServiceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetDidServiceList
// This API is used to query DID services.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetDidServiceList(request *GetDidServiceListRequest) (response *GetDidServiceListResponse, err error) {
    return c.GetDidServiceListWithContext(context.Background(), request)
}

// GetDidServiceList
// This API is used to query DID services.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetDidServiceListWithContext(ctx context.Context, request *GetDidServiceListRequest) (response *GetDidServiceListResponse, err error) {
    if request == nil {
        request = NewGetDidServiceListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "GetDidServiceList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDidServiceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDidServiceListResponse()
    err = c.Send(request, response)
    return
}

func NewGetGroupListRequest() (request *GetGroupListRequest) {
    request = &GetGroupListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "GetGroupList")
    
    
    return
}

func NewGetGroupListResponse() (response *GetGroupListResponse) {
    response = &GetGroupListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetGroupList
// This API is used to query main groups.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetGroupList(request *GetGroupListRequest) (response *GetGroupListResponse, err error) {
    return c.GetGroupListWithContext(context.Background(), request)
}

// GetGroupList
// This API is used to query main groups.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetGroupListWithContext(ctx context.Context, request *GetGroupListRequest) (response *GetGroupListResponse, err error) {
    if request == nil {
        request = NewGetGroupListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "GetGroupList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetGroupList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetGroupListResponse()
    err = c.Send(request, response)
    return
}

func NewGetLabelListRequest() (request *GetLabelListRequest) {
    request = &GetLabelListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "GetLabelList")
    
    
    return
}

func NewGetLabelListResponse() (response *GetLabelListResponse) {
    response = &GetLabelListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetLabelList
// This API is used to query labels.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetLabelList(request *GetLabelListRequest) (response *GetLabelListResponse, err error) {
    return c.GetLabelListWithContext(context.Background(), request)
}

// GetLabelList
// This API is used to query labels.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetLabelListWithContext(ctx context.Context, request *GetLabelListRequest) (response *GetLabelListResponse, err error) {
    if request == nil {
        request = NewGetLabelListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "GetLabelList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetLabelList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetLabelListResponse()
    err = c.Send(request, response)
    return
}

func NewGetPolicyListRequest() (request *GetPolicyListRequest) {
    request = &GetPolicyListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "GetPolicyList")
    
    
    return
}

func NewGetPolicyListResponse() (response *GetPolicyListResponse) {
    response = &GetPolicyListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetPolicyList
// This API is used to query disclosure policies.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetPolicyList(request *GetPolicyListRequest) (response *GetPolicyListResponse, err error) {
    return c.GetPolicyListWithContext(context.Background(), request)
}

// GetPolicyList
// This API is used to query disclosure policies.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetPolicyListWithContext(ctx context.Context, request *GetPolicyListRequest) (response *GetPolicyListResponse, err error) {
    if request == nil {
        request = NewGetPolicyListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "GetPolicyList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetPolicyList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetPolicyListResponse()
    err = c.Send(request, response)
    return
}

func NewGetPublicKeyRequest() (request *GetPublicKeyRequest) {
    request = &GetPublicKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "GetPublicKey")
    
    
    return
}

func NewGetPublicKeyResponse() (response *GetPublicKeyResponse) {
    response = &GetPublicKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetPublicKey
// This API is used to get the public key of a DID.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetPublicKey(request *GetPublicKeyRequest) (response *GetPublicKeyResponse, err error) {
    return c.GetPublicKeyWithContext(context.Background(), request)
}

// GetPublicKey
// This API is used to get the public key of a DID.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetPublicKeyWithContext(ctx context.Context, request *GetPublicKeyRequest) (response *GetPublicKeyResponse, err error) {
    if request == nil {
        request = NewGetPublicKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "GetPublicKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetPublicKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetPublicKeyResponse()
    err = c.Send(request, response)
    return
}

func NewQueryPolicyRequest() (request *QueryPolicyRequest) {
    request = &QueryPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "QueryPolicy")
    
    
    return
}

func NewQueryPolicyResponse() (response *QueryPolicyResponse) {
    response = &QueryPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryPolicy
// This API is used to get the information of a disclosure policy.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) QueryPolicy(request *QueryPolicyRequest) (response *QueryPolicyResponse, err error) {
    return c.QueryPolicyWithContext(context.Background(), request)
}

// QueryPolicy
// This API is used to get the information of a disclosure policy.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) QueryPolicyWithContext(ctx context.Context, request *QueryPolicyRequest) (response *QueryPolicyResponse, err error) {
    if request == nil {
        request = NewQueryPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "QueryPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewRecognizeAuthorityIssuerRequest() (request *RecognizeAuthorityIssuerRequest) {
    request = &RecognizeAuthorityIssuerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "RecognizeAuthorityIssuer")
    
    
    return
}

func NewRecognizeAuthorityIssuerResponse() (response *RecognizeAuthorityIssuerResponse) {
    response = &RecognizeAuthorityIssuerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RecognizeAuthorityIssuer
// This API is used to certify an issuing authority.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RecognizeAuthorityIssuer(request *RecognizeAuthorityIssuerRequest) (response *RecognizeAuthorityIssuerResponse, err error) {
    return c.RecognizeAuthorityIssuerWithContext(context.Background(), request)
}

// RecognizeAuthorityIssuer
// This API is used to certify an issuing authority.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RecognizeAuthorityIssuerWithContext(ctx context.Context, request *RecognizeAuthorityIssuerRequest) (response *RecognizeAuthorityIssuerResponse, err error) {
    if request == nil {
        request = NewRecognizeAuthorityIssuerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "RecognizeAuthorityIssuer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecognizeAuthorityIssuer require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecognizeAuthorityIssuerResponse()
    err = c.Send(request, response)
    return
}

func NewRegisterClaimPolicyRequest() (request *RegisterClaimPolicyRequest) {
    request = &RegisterClaimPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "RegisterClaimPolicy")
    
    
    return
}

func NewRegisterClaimPolicyResponse() (response *RegisterClaimPolicyResponse) {
    response = &RegisterClaimPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RegisterClaimPolicy
// This API is used to register a disclosure policy.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RegisterClaimPolicy(request *RegisterClaimPolicyRequest) (response *RegisterClaimPolicyResponse, err error) {
    return c.RegisterClaimPolicyWithContext(context.Background(), request)
}

// RegisterClaimPolicy
// This API is used to register a disclosure policy.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RegisterClaimPolicyWithContext(ctx context.Context, request *RegisterClaimPolicyRequest) (response *RegisterClaimPolicyResponse, err error) {
    if request == nil {
        request = NewRegisterClaimPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "RegisterClaimPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RegisterClaimPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewRegisterClaimPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewRegisterCptRequest() (request *RegisterCptRequest) {
    request = &RegisterCptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "RegisterCpt")
    
    
    return
}

func NewRegisterCptResponse() (response *RegisterCptResponse) {
    response = &RegisterCptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RegisterCpt
// This API is used to create a claim protocol type (CPT).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RegisterCpt(request *RegisterCptRequest) (response *RegisterCptResponse, err error) {
    return c.RegisterCptWithContext(context.Background(), request)
}

// RegisterCpt
// This API is used to create a claim protocol type (CPT).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RegisterCptWithContext(ctx context.Context, request *RegisterCptRequest) (response *RegisterCptResponse, err error) {
    if request == nil {
        request = NewRegisterCptRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "RegisterCpt")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RegisterCpt require credential")
    }

    request.SetContext(ctx)
    
    response = NewRegisterCptResponse()
    err = c.Send(request, response)
    return
}

func NewRegisterIssuerRequest() (request *RegisterIssuerRequest) {
    request = &RegisterIssuerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "RegisterIssuer")
    
    
    return
}

func NewRegisterIssuerResponse() (response *RegisterIssuerResponse) {
    response = &RegisterIssuerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RegisterIssuer
// This API is used to register an issuing authority.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RegisterIssuer(request *RegisterIssuerRequest) (response *RegisterIssuerResponse, err error) {
    return c.RegisterIssuerWithContext(context.Background(), request)
}

// RegisterIssuer
// This API is used to register an issuing authority.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RegisterIssuerWithContext(ctx context.Context, request *RegisterIssuerRequest) (response *RegisterIssuerResponse, err error) {
    if request == nil {
        request = NewRegisterIssuerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "RegisterIssuer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RegisterIssuer require credential")
    }

    request.SetContext(ctx)
    
    response = NewRegisterIssuerResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveHashRequest() (request *RemoveHashRequest) {
    request = &RemoveHashRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "RemoveHash")
    
    
    return
}

func NewRemoveHashResponse() (response *RemoveHashResponse) {
    response = &RemoveHashResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemoveHash
// This API is used to delete a contract.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RemoveHash(request *RemoveHashRequest) (response *RemoveHashResponse, err error) {
    return c.RemoveHashWithContext(context.Background(), request)
}

// RemoveHash
// This API is used to delete a contract.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RemoveHashWithContext(ctx context.Context, request *RemoveHashRequest) (response *RemoveHashResponse, err error) {
    if request == nil {
        request = NewRemoveHashRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "RemoveHash")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveHash require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveHashResponse()
    err = c.Send(request, response)
    return
}

func NewSetCredentialStatusRequest() (request *SetCredentialStatusRequest) {
    request = &SetCredentialStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "SetCredentialStatus")
    
    
    return
}

func NewSetCredentialStatusResponse() (response *SetCredentialStatusResponse) {
    response = &SetCredentialStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetCredentialStatus
// This API is used to change the on-chain status of a credential.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SetCredentialStatus(request *SetCredentialStatusRequest) (response *SetCredentialStatusResponse, err error) {
    return c.SetCredentialStatusWithContext(context.Background(), request)
}

// SetCredentialStatus
// This API is used to change the on-chain status of a credential.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SetCredentialStatusWithContext(ctx context.Context, request *SetCredentialStatusRequest) (response *SetCredentialStatusResponse, err error) {
    if request == nil {
        request = NewSetCredentialStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "SetCredentialStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetCredentialStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetCredentialStatusResponse()
    err = c.Send(request, response)
    return
}

func NewVerifyCredentialRequest() (request *VerifyCredentialRequest) {
    request = &VerifyCredentialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdid", APIVersion, "VerifyCredential")
    
    
    return
}

func NewVerifyCredentialResponse() (response *VerifyCredentialResponse) {
    response = &VerifyCredentialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// VerifyCredential
// This API is used to verify a credential.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) VerifyCredential(request *VerifyCredentialRequest) (response *VerifyCredentialResponse, err error) {
    return c.VerifyCredentialWithContext(context.Background(), request)
}

// VerifyCredential
// This API is used to verify a credential.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) VerifyCredentialWithContext(ctx context.Context, request *VerifyCredentialRequest) (response *VerifyCredentialResponse, err error) {
    if request == nil {
        request = NewVerifyCredentialRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdid", APIVersion, "VerifyCredential")
    
    if c.GetCredential() == nil {
        return nil, errors.New("VerifyCredential require credential")
    }

    request.SetContext(ctx)
    
    response = NewVerifyCredentialResponse()
    err = c.Send(request, response)
    return
}
