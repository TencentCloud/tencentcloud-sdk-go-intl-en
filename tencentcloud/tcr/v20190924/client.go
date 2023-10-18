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

package v20190924

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2019-09-24"

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


func NewCheckInstanceRequest() (request *CheckInstanceRequest) {
    request = &CheckInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "CheckInstance")
    
    
    return
}

func NewCheckInstanceResponse() (response *CheckInstanceResponse) {
    response = &CheckInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckInstance
// This API is used to verify the information of the Enterprise Edition instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORNAMEEXISTS = "InvalidParameter.ErrorNameExists"
//  INVALIDPARAMETER_ERRORREGISTRYNAME = "InvalidParameter.ErrorRegistryName"
//  INVALIDPARAMETER_ERRORTAGOVERLIMIT = "InvalidParameter.ErrorTagOverLimit"
//  INVALIDPARAMETER_UNSUPPORTEDREGION = "InvalidParameter.UnsupportedRegion"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CheckInstance(request *CheckInstanceRequest) (response *CheckInstanceResponse, err error) {
    return c.CheckInstanceWithContext(context.Background(), request)
}

// CheckInstance
// This API is used to verify the information of the Enterprise Edition instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORNAMEEXISTS = "InvalidParameter.ErrorNameExists"
//  INVALIDPARAMETER_ERRORREGISTRYNAME = "InvalidParameter.ErrorRegistryName"
//  INVALIDPARAMETER_ERRORTAGOVERLIMIT = "InvalidParameter.ErrorTagOverLimit"
//  INVALIDPARAMETER_UNSUPPORTEDREGION = "InvalidParameter.UnsupportedRegion"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CheckInstanceWithContext(ctx context.Context, request *CheckInstanceRequest) (response *CheckInstanceResponse, err error) {
    if request == nil {
        request = NewCheckInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCheckInstanceNameRequest() (request *CheckInstanceNameRequest) {
    request = &CheckInstanceNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "CheckInstanceName")
    
    
    return
}

func NewCheckInstanceNameResponse() (response *CheckInstanceNameResponse) {
    response = &CheckInstanceNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckInstanceName
// This API is used to check whether the name of the instance to be created meets the specifications.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORNAMEEXISTS = "InvalidParameter.ErrorNameExists"
//  INVALIDPARAMETER_ERRORREGISTRYNAME = "InvalidParameter.ErrorRegistryName"
//  INVALIDPARAMETER_ERRORTAGOVERLIMIT = "InvalidParameter.ErrorTagOverLimit"
//  INVALIDPARAMETER_UNSUPPORTEDREGION = "InvalidParameter.UnsupportedRegion"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CheckInstanceName(request *CheckInstanceNameRequest) (response *CheckInstanceNameResponse, err error) {
    return c.CheckInstanceNameWithContext(context.Background(), request)
}

// CheckInstanceName
// This API is used to check whether the name of the instance to be created meets the specifications.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORNAMEEXISTS = "InvalidParameter.ErrorNameExists"
//  INVALIDPARAMETER_ERRORREGISTRYNAME = "InvalidParameter.ErrorRegistryName"
//  INVALIDPARAMETER_ERRORTAGOVERLIMIT = "InvalidParameter.ErrorTagOverLimit"
//  INVALIDPARAMETER_UNSUPPORTEDREGION = "InvalidParameter.UnsupportedRegion"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CheckInstanceNameWithContext(ctx context.Context, request *CheckInstanceNameRequest) (response *CheckInstanceNameResponse, err error) {
    if request == nil {
        request = NewCheckInstanceNameRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckInstanceName require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckInstanceNameResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCustomAccountRequest() (request *CreateCustomAccountRequest) {
    request = &CreateCustomAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "CreateCustomAccount")
    
    
    return
}

func NewCreateCustomAccountResponse() (response *CreateCustomAccountResponse) {
    response = &CreateCustomAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCustomAccount
// This API is used to create a custom account.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_QUOTAOVERLIMIT = "OperationDenied.QuotaOverLimit"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCustomAccount(request *CreateCustomAccountRequest) (response *CreateCustomAccountResponse, err error) {
    return c.CreateCustomAccountWithContext(context.Background(), request)
}

// CreateCustomAccount
// This API is used to create a custom account.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_QUOTAOVERLIMIT = "OperationDenied.QuotaOverLimit"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCustomAccountWithContext(ctx context.Context, request *CreateCustomAccountRequest) (response *CreateCustomAccountResponse, err error) {
    if request == nil {
        request = NewCreateCustomAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCustomAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCustomAccountResponse()
    err = c.Send(request, response)
    return
}

func NewCreateImageAccelerationServiceRequest() (request *CreateImageAccelerationServiceRequest) {
    request = &CreateImageAccelerationServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "CreateImageAccelerationService")
    
    
    return
}

func NewCreateImageAccelerationServiceResponse() (response *CreateImageAccelerationServiceResponse) {
    response = &CreateImageAccelerationServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateImageAccelerationService
// This API is used to create an image acceleration service.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateImageAccelerationService(request *CreateImageAccelerationServiceRequest) (response *CreateImageAccelerationServiceResponse, err error) {
    return c.CreateImageAccelerationServiceWithContext(context.Background(), request)
}

// CreateImageAccelerationService
// This API is used to create an image acceleration service.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateImageAccelerationServiceWithContext(ctx context.Context, request *CreateImageAccelerationServiceRequest) (response *CreateImageAccelerationServiceResponse, err error) {
    if request == nil {
        request = NewCreateImageAccelerationServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateImageAccelerationService require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateImageAccelerationServiceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateImmutableTagRulesRequest() (request *CreateImmutableTagRulesRequest) {
    request = &CreateImmutableTagRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "CreateImmutableTagRules")
    
    
    return
}

func NewCreateImmutableTagRulesResponse() (response *CreateImmutableTagRulesResponse) {
    response = &CreateImmutableTagRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateImmutableTagRules
// This API is used to create the tag immutability rule.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateImmutableTagRules(request *CreateImmutableTagRulesRequest) (response *CreateImmutableTagRulesResponse, err error) {
    return c.CreateImmutableTagRulesWithContext(context.Background(), request)
}

// CreateImmutableTagRules
// This API is used to create the tag immutability rule.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateImmutableTagRulesWithContext(ctx context.Context, request *CreateImmutableTagRulesRequest) (response *CreateImmutableTagRulesResponse, err error) {
    if request == nil {
        request = NewCreateImmutableTagRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateImmutableTagRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateImmutableTagRulesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstanceRequest() (request *CreateInstanceRequest) {
    request = &CreateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "CreateInstance")
    
    
    return
}

func NewCreateInstanceResponse() (response *CreateInstanceResponse) {
    response = &CreateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateInstance
// This API is used to create an instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETDBDATAERROR = "FailedOperation.GetDBDataError"
//  FAILEDOPERATION_TRADEFAILED = "FailedOperation.TradeFailed"
//  FAILEDOPERATION_VALIDATEREGISTRYNAMEFAIL = "FailedOperation.ValidateRegistryNameFail"
//  FAILEDOPERATION_VALIDATESUPPORTEDREGIONFAIL = "FailedOperation.ValidateSupportedRegionFail"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORNAMEEXISTS = "InvalidParameter.ErrorNameExists"
//  INVALIDPARAMETER_ERRORNAMEILLEGAL = "InvalidParameter.ErrorNameIllegal"
//  INVALIDPARAMETER_ERRORNAMERESERVED = "InvalidParameter.ErrorNameReserved"
//  INVALIDPARAMETER_ERRORREGISTRYNAME = "InvalidParameter.ErrorRegistryName"
//  INVALIDPARAMETER_ERRORTAGOVERLIMIT = "InvalidParameter.ErrorTagOverLimit"
//  INVALIDPARAMETER_UNSUPPORTEDREGION = "InvalidParameter.UnsupportedRegion"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_QUOTAOVERLIMIT = "OperationDenied.QuotaOverLimit"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateInstance(request *CreateInstanceRequest) (response *CreateInstanceResponse, err error) {
    return c.CreateInstanceWithContext(context.Background(), request)
}

// CreateInstance
// This API is used to create an instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETDBDATAERROR = "FailedOperation.GetDBDataError"
//  FAILEDOPERATION_TRADEFAILED = "FailedOperation.TradeFailed"
//  FAILEDOPERATION_VALIDATEREGISTRYNAMEFAIL = "FailedOperation.ValidateRegistryNameFail"
//  FAILEDOPERATION_VALIDATESUPPORTEDREGIONFAIL = "FailedOperation.ValidateSupportedRegionFail"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORNAMEEXISTS = "InvalidParameter.ErrorNameExists"
//  INVALIDPARAMETER_ERRORNAMEILLEGAL = "InvalidParameter.ErrorNameIllegal"
//  INVALIDPARAMETER_ERRORNAMERESERVED = "InvalidParameter.ErrorNameReserved"
//  INVALIDPARAMETER_ERRORREGISTRYNAME = "InvalidParameter.ErrorRegistryName"
//  INVALIDPARAMETER_ERRORTAGOVERLIMIT = "InvalidParameter.ErrorTagOverLimit"
//  INVALIDPARAMETER_UNSUPPORTEDREGION = "InvalidParameter.UnsupportedRegion"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_QUOTAOVERLIMIT = "OperationDenied.QuotaOverLimit"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateInstanceWithContext(ctx context.Context, request *CreateInstanceRequest) (response *CreateInstanceResponse, err error) {
    if request == nil {
        request = NewCreateInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstanceCustomizedDomainRequest() (request *CreateInstanceCustomizedDomainRequest) {
    request = &CreateInstanceCustomizedDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "CreateInstanceCustomizedDomain")
    
    
    return
}

func NewCreateInstanceCustomizedDomainResponse() (response *CreateInstanceCustomizedDomainResponse) {
    response = &CreateInstanceCustomizedDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateInstanceCustomizedDomain
// This API is used to create a custom domain name.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTEDREGION = "InvalidParameter.UnsupportedRegion"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateInstanceCustomizedDomain(request *CreateInstanceCustomizedDomainRequest) (response *CreateInstanceCustomizedDomainResponse, err error) {
    return c.CreateInstanceCustomizedDomainWithContext(context.Background(), request)
}

// CreateInstanceCustomizedDomain
// This API is used to create a custom domain name.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTEDREGION = "InvalidParameter.UnsupportedRegion"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateInstanceCustomizedDomainWithContext(ctx context.Context, request *CreateInstanceCustomizedDomainRequest) (response *CreateInstanceCustomizedDomainResponse, err error) {
    if request == nil {
        request = NewCreateInstanceCustomizedDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInstanceCustomizedDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInstanceCustomizedDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstanceTokenRequest() (request *CreateInstanceTokenRequest) {
    request = &CreateInstanceTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "CreateInstanceToken")
    
    
    return
}

func NewCreateInstanceTokenResponse() (response *CreateInstanceTokenResponse) {
    response = &CreateInstanceTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateInstanceToken
// This API is used to create a temporary or long-term instance access credential.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateInstanceToken(request *CreateInstanceTokenRequest) (response *CreateInstanceTokenResponse, err error) {
    return c.CreateInstanceTokenWithContext(context.Background(), request)
}

// CreateInstanceToken
// This API is used to create a temporary or long-term instance access credential.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateInstanceTokenWithContext(ctx context.Context, request *CreateInstanceTokenRequest) (response *CreateInstanceTokenResponse, err error) {
    if request == nil {
        request = NewCreateInstanceTokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInstanceToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInstanceTokenResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMultipleSecurityPolicyRequest() (request *CreateMultipleSecurityPolicyRequest) {
    request = &CreateMultipleSecurityPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "CreateMultipleSecurityPolicy")
    
    
    return
}

func NewCreateMultipleSecurityPolicyResponse() (response *CreateMultipleSecurityPolicyResponse) {
    response = &CreateMultipleSecurityPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateMultipleSecurityPolicy
// This API is used to create multiple public network access allowlist policies of the TCR instance.
//
// error code that may be returned:
//  FAILEDOPERATION_DEPENDENCEERROR = "FailedOperation.DependenceError"
//  FAILEDOPERATION_ERRORGETDBDATAERROR = "FailedOperation.ErrorGetDBDataError"
//  FAILEDOPERATION_GETDBDATAERROR = "FailedOperation.GetDBDataError"
//  FAILEDOPERATION_GETTCRCLIENT = "FailedOperation.GetTcrClient"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateMultipleSecurityPolicy(request *CreateMultipleSecurityPolicyRequest) (response *CreateMultipleSecurityPolicyResponse, err error) {
    return c.CreateMultipleSecurityPolicyWithContext(context.Background(), request)
}

// CreateMultipleSecurityPolicy
// This API is used to create multiple public network access allowlist policies of the TCR instance.
//
// error code that may be returned:
//  FAILEDOPERATION_DEPENDENCEERROR = "FailedOperation.DependenceError"
//  FAILEDOPERATION_ERRORGETDBDATAERROR = "FailedOperation.ErrorGetDBDataError"
//  FAILEDOPERATION_GETDBDATAERROR = "FailedOperation.GetDBDataError"
//  FAILEDOPERATION_GETTCRCLIENT = "FailedOperation.GetTcrClient"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateMultipleSecurityPolicyWithContext(ctx context.Context, request *CreateMultipleSecurityPolicyRequest) (response *CreateMultipleSecurityPolicyResponse, err error) {
    if request == nil {
        request = NewCreateMultipleSecurityPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMultipleSecurityPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMultipleSecurityPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNamespaceRequest() (request *CreateNamespaceRequest) {
    request = &CreateNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "CreateNamespace")
    
    
    return
}

func NewCreateNamespaceResponse() (response *CreateNamespaceResponse) {
    response = &CreateNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateNamespace
// This API is used to create a namespace in an Enterprise Edition instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DEPENDENCEERROR = "FailedOperation.DependenceError"
//  FAILEDOPERATION_ERRORTCRINVALIDMEDIATYPE = "FailedOperation.ErrorTcrInvalidMediaType"
//  FAILEDOPERATION_ERRORTCRRESOURCECONFLICT = "FailedOperation.ErrorTcrResourceConflict"
//  FAILEDOPERATION_ERRORTCRUNAUTHORIZED = "FailedOperation.ErrorTcrUnauthorized"
//  FAILEDOPERATION_OPERATIONCANCEL = "FailedOperation.OperationCancel"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateNamespace(request *CreateNamespaceRequest) (response *CreateNamespaceResponse, err error) {
    return c.CreateNamespaceWithContext(context.Background(), request)
}

// CreateNamespace
// This API is used to create a namespace in an Enterprise Edition instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DEPENDENCEERROR = "FailedOperation.DependenceError"
//  FAILEDOPERATION_ERRORTCRINVALIDMEDIATYPE = "FailedOperation.ErrorTcrInvalidMediaType"
//  FAILEDOPERATION_ERRORTCRRESOURCECONFLICT = "FailedOperation.ErrorTcrResourceConflict"
//  FAILEDOPERATION_ERRORTCRUNAUTHORIZED = "FailedOperation.ErrorTcrUnauthorized"
//  FAILEDOPERATION_OPERATIONCANCEL = "FailedOperation.OperationCancel"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateNamespaceWithContext(ctx context.Context, request *CreateNamespaceRequest) (response *CreateNamespaceResponse, err error) {
    if request == nil {
        request = NewCreateNamespaceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNamespace require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateReplicationInstanceRequest() (request *CreateReplicationInstanceRequest) {
    request = &CreateReplicationInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "CreateReplicationInstance")
    
    
    return
}

func NewCreateReplicationInstanceResponse() (response *CreateReplicationInstanceResponse) {
    response = &CreateReplicationInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateReplicationInstance
// This API is used to create a replication instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_REPLICATIONEXISTS = "InvalidParameter.ReplicationExists"
//  INVALIDPARAMETER_UNSUPPORTEDREGION = "InvalidParameter.UnsupportedRegion"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateReplicationInstance(request *CreateReplicationInstanceRequest) (response *CreateReplicationInstanceResponse, err error) {
    return c.CreateReplicationInstanceWithContext(context.Background(), request)
}

// CreateReplicationInstance
// This API is used to create a replication instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_REPLICATIONEXISTS = "InvalidParameter.ReplicationExists"
//  INVALIDPARAMETER_UNSUPPORTEDREGION = "InvalidParameter.UnsupportedRegion"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateReplicationInstanceWithContext(ctx context.Context, request *CreateReplicationInstanceRequest) (response *CreateReplicationInstanceResponse, err error) {
    if request == nil {
        request = NewCreateReplicationInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateReplicationInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateReplicationInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRepositoryRequest() (request *CreateRepositoryRequest) {
    request = &CreateRepositoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "CreateRepository")
    
    
    return
}

func NewCreateRepositoryResponse() (response *CreateRepositoryResponse) {
    response = &CreateRepositoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRepository
// This API is used to create an image repository in an Enterprise Edition instance.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONCANCEL = "FailedOperation.OperationCancel"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateRepository(request *CreateRepositoryRequest) (response *CreateRepositoryResponse, err error) {
    return c.CreateRepositoryWithContext(context.Background(), request)
}

// CreateRepository
// This API is used to create an image repository in an Enterprise Edition instance.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONCANCEL = "FailedOperation.OperationCancel"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateRepositoryWithContext(ctx context.Context, request *CreateRepositoryRequest) (response *CreateRepositoryResponse, err error) {
    if request == nil {
        request = NewCreateRepositoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRepository require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRepositoryResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSecurityPolicyRequest() (request *CreateSecurityPolicyRequest) {
    request = &CreateSecurityPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "CreateSecurityPolicy")
    
    
    return
}

func NewCreateSecurityPolicyResponse() (response *CreateSecurityPolicyResponse) {
    response = &CreateSecurityPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSecurityPolicy
// This API is used to create a public network access allowlist policy for an instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSecurityPolicy(request *CreateSecurityPolicyRequest) (response *CreateSecurityPolicyResponse, err error) {
    return c.CreateSecurityPolicyWithContext(context.Background(), request)
}

// CreateSecurityPolicy
// This API is used to create a public network access allowlist policy for an instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSecurityPolicyWithContext(ctx context.Context, request *CreateSecurityPolicyRequest) (response *CreateSecurityPolicyResponse, err error) {
    if request == nil {
        request = NewCreateSecurityPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSecurityPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSecurityPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateServiceAccountRequest() (request *CreateServiceAccountRequest) {
    request = &CreateServiceAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "CreateServiceAccount")
    
    
    return
}

func NewCreateServiceAccountResponse() (response *CreateServiceAccountResponse) {
    response = &CreateServiceAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateServiceAccount
// This API is used to create a service account.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_QUOTAOVERLIMIT = "OperationDenied.QuotaOverLimit"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateServiceAccount(request *CreateServiceAccountRequest) (response *CreateServiceAccountResponse, err error) {
    return c.CreateServiceAccountWithContext(context.Background(), request)
}

// CreateServiceAccount
// This API is used to create a service account.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_QUOTAOVERLIMIT = "OperationDenied.QuotaOverLimit"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateServiceAccountWithContext(ctx context.Context, request *CreateServiceAccountRequest) (response *CreateServiceAccountResponse, err error) {
    if request == nil {
        request = NewCreateServiceAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateServiceAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateServiceAccountResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSignatureRequest() (request *CreateSignatureRequest) {
    request = &CreateSignatureRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "CreateSignature")
    
    
    return
}

func NewCreateSignatureResponse() (response *CreateSignatureResponse) {
    response = &CreateSignatureResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSignature
// This API is used to create a signature for an image tag.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONCANCEL = "FailedOperation.OperationCancel"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSignature(request *CreateSignatureRequest) (response *CreateSignatureResponse, err error) {
    return c.CreateSignatureWithContext(context.Background(), request)
}

// CreateSignature
// This API is used to create a signature for an image tag.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONCANCEL = "FailedOperation.OperationCancel"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSignatureWithContext(ctx context.Context, request *CreateSignatureRequest) (response *CreateSignatureResponse, err error) {
    if request == nil {
        request = NewCreateSignatureRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSignature require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSignatureResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSignaturePolicyRequest() (request *CreateSignaturePolicyRequest) {
    request = &CreateSignaturePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "CreateSignaturePolicy")
    
    
    return
}

func NewCreateSignaturePolicyResponse() (response *CreateSignaturePolicyResponse) {
    response = &CreateSignaturePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSignaturePolicy
// This API is used to create an image signature policy.
//
// error code that may be returned:
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSignaturePolicy(request *CreateSignaturePolicyRequest) (response *CreateSignaturePolicyResponse, err error) {
    return c.CreateSignaturePolicyWithContext(context.Background(), request)
}

// CreateSignaturePolicy
// This API is used to create an image signature policy.
//
// error code that may be returned:
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSignaturePolicyWithContext(ctx context.Context, request *CreateSignaturePolicyRequest) (response *CreateSignaturePolicyResponse, err error) {
    if request == nil {
        request = NewCreateSignaturePolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSignaturePolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSignaturePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTagRetentionExecutionRequest() (request *CreateTagRetentionExecutionRequest) {
    request = &CreateTagRetentionExecutionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "CreateTagRetentionExecution")
    
    
    return
}

func NewCreateTagRetentionExecutionResponse() (response *CreateTagRetentionExecutionResponse) {
    response = &CreateTagRetentionExecutionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTagRetentionExecution
// This API is used to execute tag retention manually.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateTagRetentionExecution(request *CreateTagRetentionExecutionRequest) (response *CreateTagRetentionExecutionResponse, err error) {
    return c.CreateTagRetentionExecutionWithContext(context.Background(), request)
}

// CreateTagRetentionExecution
// This API is used to execute tag retention manually.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateTagRetentionExecutionWithContext(ctx context.Context, request *CreateTagRetentionExecutionRequest) (response *CreateTagRetentionExecutionResponse, err error) {
    if request == nil {
        request = NewCreateTagRetentionExecutionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTagRetentionExecution require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTagRetentionExecutionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTagRetentionRuleRequest() (request *CreateTagRetentionRuleRequest) {
    request = &CreateTagRetentionRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "CreateTagRetentionRule")
    
    
    return
}

func NewCreateTagRetentionRuleResponse() (response *CreateTagRetentionRuleResponse) {
    response = &CreateTagRetentionRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTagRetentionRule
// This API is used to create a tag retention rule.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateTagRetentionRule(request *CreateTagRetentionRuleRequest) (response *CreateTagRetentionRuleResponse, err error) {
    return c.CreateTagRetentionRuleWithContext(context.Background(), request)
}

// CreateTagRetentionRule
// This API is used to create a tag retention rule.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateTagRetentionRuleWithContext(ctx context.Context, request *CreateTagRetentionRuleRequest) (response *CreateTagRetentionRuleResponse, err error) {
    if request == nil {
        request = NewCreateTagRetentionRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTagRetentionRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTagRetentionRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWebhookTriggerRequest() (request *CreateWebhookTriggerRequest) {
    request = &CreateWebhookTriggerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "CreateWebhookTrigger")
    
    
    return
}

func NewCreateWebhookTriggerResponse() (response *CreateWebhookTriggerResponse) {
    response = &CreateWebhookTriggerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateWebhookTrigger
// This API is used to create a trigger.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
func (c *Client) CreateWebhookTrigger(request *CreateWebhookTriggerRequest) (response *CreateWebhookTriggerResponse, err error) {
    return c.CreateWebhookTriggerWithContext(context.Background(), request)
}

// CreateWebhookTrigger
// This API is used to create a trigger.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
func (c *Client) CreateWebhookTriggerWithContext(ctx context.Context, request *CreateWebhookTriggerRequest) (response *CreateWebhookTriggerResponse, err error) {
    if request == nil {
        request = NewCreateWebhookTriggerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWebhookTrigger require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWebhookTriggerResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCustomAccountRequest() (request *DeleteCustomAccountRequest) {
    request = &DeleteCustomAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteCustomAccount")
    
    
    return
}

func NewDeleteCustomAccountResponse() (response *DeleteCustomAccountResponse) {
    response = &DeleteCustomAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCustomAccount
// This API is used to delete a custom account.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteCustomAccount(request *DeleteCustomAccountRequest) (response *DeleteCustomAccountResponse, err error) {
    return c.DeleteCustomAccountWithContext(context.Background(), request)
}

// DeleteCustomAccount
// This API is used to delete a custom account.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteCustomAccountWithContext(ctx context.Context, request *DeleteCustomAccountRequest) (response *DeleteCustomAccountResponse, err error) {
    if request == nil {
        request = NewDeleteCustomAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCustomAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCustomAccountResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteImageRequest() (request *DeleteImageRequest) {
    request = &DeleteImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteImage")
    
    
    return
}

func NewDeleteImageResponse() (response *DeleteImageResponse) {
    response = &DeleteImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteImage
// This API is used to delete the specified image.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteImage(request *DeleteImageRequest) (response *DeleteImageResponse, err error) {
    return c.DeleteImageWithContext(context.Background(), request)
}

// DeleteImage
// This API is used to delete the specified image.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteImageWithContext(ctx context.Context, request *DeleteImageRequest) (response *DeleteImageResponse, err error) {
    if request == nil {
        request = NewDeleteImageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteImage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteImageResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteImageAccelerateServiceRequest() (request *DeleteImageAccelerateServiceRequest) {
    request = &DeleteImageAccelerateServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteImageAccelerateService")
    
    
    return
}

func NewDeleteImageAccelerateServiceResponse() (response *DeleteImageAccelerateServiceResponse) {
    response = &DeleteImageAccelerateServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteImageAccelerateService
// This API is used to delete an image acceleration service.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteImageAccelerateService(request *DeleteImageAccelerateServiceRequest) (response *DeleteImageAccelerateServiceResponse, err error) {
    return c.DeleteImageAccelerateServiceWithContext(context.Background(), request)
}

// DeleteImageAccelerateService
// This API is used to delete an image acceleration service.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteImageAccelerateServiceWithContext(ctx context.Context, request *DeleteImageAccelerateServiceRequest) (response *DeleteImageAccelerateServiceResponse, err error) {
    if request == nil {
        request = NewDeleteImageAccelerateServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteImageAccelerateService require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteImageAccelerateServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteImmutableTagRulesRequest() (request *DeleteImmutableTagRulesRequest) {
    request = &DeleteImmutableTagRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteImmutableTagRules")
    
    
    return
}

func NewDeleteImmutableTagRulesResponse() (response *DeleteImmutableTagRulesResponse) {
    response = &DeleteImmutableTagRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteImmutableTagRules
//  This API is used to delete the tag immutability rule.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteImmutableTagRules(request *DeleteImmutableTagRulesRequest) (response *DeleteImmutableTagRulesResponse, err error) {
    return c.DeleteImmutableTagRulesWithContext(context.Background(), request)
}

// DeleteImmutableTagRules
//  This API is used to delete the tag immutability rule.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteImmutableTagRulesWithContext(ctx context.Context, request *DeleteImmutableTagRulesRequest) (response *DeleteImmutableTagRulesResponse, err error) {
    if request == nil {
        request = NewDeleteImmutableTagRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteImmutableTagRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteImmutableTagRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteInstanceRequest() (request *DeleteInstanceRequest) {
    request = &DeleteInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteInstance")
    
    
    return
}

func NewDeleteInstanceResponse() (response *DeleteInstanceResponse) {
    response = &DeleteInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteInstance
// This API is used to delete a TCR Enterprise Edition instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteInstance(request *DeleteInstanceRequest) (response *DeleteInstanceResponse, err error) {
    return c.DeleteInstanceWithContext(context.Background(), request)
}

// DeleteInstance
// This API is used to delete a TCR Enterprise Edition instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteInstanceWithContext(ctx context.Context, request *DeleteInstanceRequest) (response *DeleteInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteInstanceCustomizedDomainRequest() (request *DeleteInstanceCustomizedDomainRequest) {
    request = &DeleteInstanceCustomizedDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteInstanceCustomizedDomain")
    
    
    return
}

func NewDeleteInstanceCustomizedDomainResponse() (response *DeleteInstanceCustomizedDomainResponse) {
    response = &DeleteInstanceCustomizedDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteInstanceCustomizedDomain
// This API is used to delete a custom domain name.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTEDREGION = "InvalidParameter.UnsupportedRegion"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteInstanceCustomizedDomain(request *DeleteInstanceCustomizedDomainRequest) (response *DeleteInstanceCustomizedDomainResponse, err error) {
    return c.DeleteInstanceCustomizedDomainWithContext(context.Background(), request)
}

// DeleteInstanceCustomizedDomain
// This API is used to delete a custom domain name.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTEDREGION = "InvalidParameter.UnsupportedRegion"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteInstanceCustomizedDomainWithContext(ctx context.Context, request *DeleteInstanceCustomizedDomainRequest) (response *DeleteInstanceCustomizedDomainResponse, err error) {
    if request == nil {
        request = NewDeleteInstanceCustomizedDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteInstanceCustomizedDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteInstanceCustomizedDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteInstanceTokenRequest() (request *DeleteInstanceTokenRequest) {
    request = &DeleteInstanceTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteInstanceToken")
    
    
    return
}

func NewDeleteInstanceTokenResponse() (response *DeleteInstanceTokenResponse) {
    response = &DeleteInstanceTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteInstanceToken
// This API is used to delete a long-term access credential.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteInstanceToken(request *DeleteInstanceTokenRequest) (response *DeleteInstanceTokenResponse, err error) {
    return c.DeleteInstanceTokenWithContext(context.Background(), request)
}

// DeleteInstanceToken
// This API is used to delete a long-term access credential.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteInstanceTokenWithContext(ctx context.Context, request *DeleteInstanceTokenRequest) (response *DeleteInstanceTokenResponse, err error) {
    if request == nil {
        request = NewDeleteInstanceTokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteInstanceToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteInstanceTokenResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMultipleSecurityPolicyRequest() (request *DeleteMultipleSecurityPolicyRequest) {
    request = &DeleteMultipleSecurityPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteMultipleSecurityPolicy")
    
    
    return
}

func NewDeleteMultipleSecurityPolicyResponse() (response *DeleteMultipleSecurityPolicyResponse) {
    response = &DeleteMultipleSecurityPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteMultipleSecurityPolicy
// This API is used to delete multiple public network access allowlist policies of the instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteMultipleSecurityPolicy(request *DeleteMultipleSecurityPolicyRequest) (response *DeleteMultipleSecurityPolicyResponse, err error) {
    return c.DeleteMultipleSecurityPolicyWithContext(context.Background(), request)
}

// DeleteMultipleSecurityPolicy
// This API is used to delete multiple public network access allowlist policies of the instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteMultipleSecurityPolicyWithContext(ctx context.Context, request *DeleteMultipleSecurityPolicyRequest) (response *DeleteMultipleSecurityPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteMultipleSecurityPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMultipleSecurityPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMultipleSecurityPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNamespaceRequest() (request *DeleteNamespaceRequest) {
    request = &DeleteNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteNamespace")
    
    
    return
}

func NewDeleteNamespaceResponse() (response *DeleteNamespaceResponse) {
    response = &DeleteNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteNamespace
// This API is used to delete a namespace.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORGETDBDATAERROR = "FailedOperation.ErrorGetDBDataError"
//  FAILEDOPERATION_ERRORTCRRESOURCECONFLICT = "FailedOperation.ErrorTcrResourceConflict"
//  FAILEDOPERATION_ERRORTCRUNAUTHORIZED = "FailedOperation.ErrorTcrUnauthorized"
//  FAILEDOPERATION_GETDBDATAERROR = "FailedOperation.GetDBDataError"
//  FAILEDOPERATION_PRECONDITIONFAILED = "FailedOperation.PreconditionFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteNamespace(request *DeleteNamespaceRequest) (response *DeleteNamespaceResponse, err error) {
    return c.DeleteNamespaceWithContext(context.Background(), request)
}

// DeleteNamespace
// This API is used to delete a namespace.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORGETDBDATAERROR = "FailedOperation.ErrorGetDBDataError"
//  FAILEDOPERATION_ERRORTCRRESOURCECONFLICT = "FailedOperation.ErrorTcrResourceConflict"
//  FAILEDOPERATION_ERRORTCRUNAUTHORIZED = "FailedOperation.ErrorTcrUnauthorized"
//  FAILEDOPERATION_GETDBDATAERROR = "FailedOperation.GetDBDataError"
//  FAILEDOPERATION_PRECONDITIONFAILED = "FailedOperation.PreconditionFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteNamespaceWithContext(ctx context.Context, request *DeleteNamespaceRequest) (response *DeleteNamespaceResponse, err error) {
    if request == nil {
        request = NewDeleteNamespaceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteNamespace require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteReplicationInstanceRequest() (request *DeleteReplicationInstanceRequest) {
    request = &DeleteReplicationInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteReplicationInstance")
    
    
    return
}

func NewDeleteReplicationInstanceResponse() (response *DeleteReplicationInstanceResponse) {
    response = &DeleteReplicationInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteReplicationInstance
// This API is used to delete a replica instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteReplicationInstance(request *DeleteReplicationInstanceRequest) (response *DeleteReplicationInstanceResponse, err error) {
    return c.DeleteReplicationInstanceWithContext(context.Background(), request)
}

// DeleteReplicationInstance
// This API is used to delete a replica instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteReplicationInstanceWithContext(ctx context.Context, request *DeleteReplicationInstanceRequest) (response *DeleteReplicationInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteReplicationInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteReplicationInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteReplicationInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRepositoryRequest() (request *DeleteRepositoryRequest) {
    request = &DeleteRepositoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteRepository")
    
    
    return
}

func NewDeleteRepositoryResponse() (response *DeleteRepositoryResponse) {
    response = &DeleteRepositoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRepository
// This API is used to delete an image repository.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRepository(request *DeleteRepositoryRequest) (response *DeleteRepositoryResponse, err error) {
    return c.DeleteRepositoryWithContext(context.Background(), request)
}

// DeleteRepository
// This API is used to delete an image repository.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRepositoryWithContext(ctx context.Context, request *DeleteRepositoryRequest) (response *DeleteRepositoryResponse, err error) {
    if request == nil {
        request = NewDeleteRepositoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRepository require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRepositoryResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRepositoryTagsRequest() (request *DeleteRepositoryTagsRequest) {
    request = &DeleteRepositoryTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteRepositoryTags")
    
    
    return
}

func NewDeleteRepositoryTagsResponse() (response *DeleteRepositoryTagsResponse) {
    response = &DeleteRepositoryTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRepositoryTags
// This API is used to batch delete repository tags in an Enterprise Edition instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRepositoryTags(request *DeleteRepositoryTagsRequest) (response *DeleteRepositoryTagsResponse, err error) {
    return c.DeleteRepositoryTagsWithContext(context.Background(), request)
}

// DeleteRepositoryTags
// This API is used to batch delete repository tags in an Enterprise Edition instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRepositoryTagsWithContext(ctx context.Context, request *DeleteRepositoryTagsRequest) (response *DeleteRepositoryTagsResponse, err error) {
    if request == nil {
        request = NewDeleteRepositoryTagsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRepositoryTags require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRepositoryTagsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSecurityPolicyRequest() (request *DeleteSecurityPolicyRequest) {
    request = &DeleteSecurityPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteSecurityPolicy")
    
    
    return
}

func NewDeleteSecurityPolicyResponse() (response *DeleteSecurityPolicyResponse) {
    response = &DeleteSecurityPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSecurityPolicy
// This API is used to delete a public network access allow policy.
//
// 
//
// Note: When both `PolicyIndex` and `CidrBlock` are specified, `CidrBlock` takes the higher priority
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteSecurityPolicy(request *DeleteSecurityPolicyRequest) (response *DeleteSecurityPolicyResponse, err error) {
    return c.DeleteSecurityPolicyWithContext(context.Background(), request)
}

// DeleteSecurityPolicy
// This API is used to delete a public network access allow policy.
//
// 
//
// Note: When both `PolicyIndex` and `CidrBlock` are specified, `CidrBlock` takes the higher priority
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteSecurityPolicyWithContext(ctx context.Context, request *DeleteSecurityPolicyRequest) (response *DeleteSecurityPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteSecurityPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSecurityPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSecurityPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteServiceAccountRequest() (request *DeleteServiceAccountRequest) {
    request = &DeleteServiceAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteServiceAccount")
    
    
    return
}

func NewDeleteServiceAccountResponse() (response *DeleteServiceAccountResponse) {
    response = &DeleteServiceAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteServiceAccount
// This API is used to delete a service account.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteServiceAccount(request *DeleteServiceAccountRequest) (response *DeleteServiceAccountResponse, err error) {
    return c.DeleteServiceAccountWithContext(context.Background(), request)
}

// DeleteServiceAccount
// This API is used to delete a service account.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteServiceAccountWithContext(ctx context.Context, request *DeleteServiceAccountRequest) (response *DeleteServiceAccountResponse, err error) {
    if request == nil {
        request = NewDeleteServiceAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteServiceAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteServiceAccountResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSignaturePolicyRequest() (request *DeleteSignaturePolicyRequest) {
    request = &DeleteSignaturePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteSignaturePolicy")
    
    
    return
}

func NewDeleteSignaturePolicyResponse() (response *DeleteSignaturePolicyResponse) {
    response = &DeleteSignaturePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSignaturePolicy
// This API is used to delete a namespace signing policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORGETDBDATAERROR = "FailedOperation.ErrorGetDBDataError"
//  FAILEDOPERATION_ERRORTCRRESOURCECONFLICT = "FailedOperation.ErrorTcrResourceConflict"
//  FAILEDOPERATION_ERRORTCRUNAUTHORIZED = "FailedOperation.ErrorTcrUnauthorized"
//  FAILEDOPERATION_GETDBDATAERROR = "FailedOperation.GetDBDataError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteSignaturePolicy(request *DeleteSignaturePolicyRequest) (response *DeleteSignaturePolicyResponse, err error) {
    return c.DeleteSignaturePolicyWithContext(context.Background(), request)
}

// DeleteSignaturePolicy
// This API is used to delete a namespace signing policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORGETDBDATAERROR = "FailedOperation.ErrorGetDBDataError"
//  FAILEDOPERATION_ERRORTCRRESOURCECONFLICT = "FailedOperation.ErrorTcrResourceConflict"
//  FAILEDOPERATION_ERRORTCRUNAUTHORIZED = "FailedOperation.ErrorTcrUnauthorized"
//  FAILEDOPERATION_GETDBDATAERROR = "FailedOperation.GetDBDataError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteSignaturePolicyWithContext(ctx context.Context, request *DeleteSignaturePolicyRequest) (response *DeleteSignaturePolicyResponse, err error) {
    if request == nil {
        request = NewDeleteSignaturePolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSignaturePolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSignaturePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTagRetentionRuleRequest() (request *DeleteTagRetentionRuleRequest) {
    request = &DeleteTagRetentionRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteTagRetentionRule")
    
    
    return
}

func NewDeleteTagRetentionRuleResponse() (response *DeleteTagRetentionRuleResponse) {
    response = &DeleteTagRetentionRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTagRetentionRule
// This API is used to delete a tag retention rule.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteTagRetentionRule(request *DeleteTagRetentionRuleRequest) (response *DeleteTagRetentionRuleResponse, err error) {
    return c.DeleteTagRetentionRuleWithContext(context.Background(), request)
}

// DeleteTagRetentionRule
// This API is used to delete a tag retention rule.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteTagRetentionRuleWithContext(ctx context.Context, request *DeleteTagRetentionRuleRequest) (response *DeleteTagRetentionRuleResponse, err error) {
    if request == nil {
        request = NewDeleteTagRetentionRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTagRetentionRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTagRetentionRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWebhookTriggerRequest() (request *DeleteWebhookTriggerRequest) {
    request = &DeleteWebhookTriggerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteWebhookTrigger")
    
    
    return
}

func NewDeleteWebhookTriggerResponse() (response *DeleteWebhookTriggerResponse) {
    response = &DeleteWebhookTriggerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteWebhookTrigger
// This API is used to delete a trigger.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
func (c *Client) DeleteWebhookTrigger(request *DeleteWebhookTriggerRequest) (response *DeleteWebhookTriggerResponse, err error) {
    return c.DeleteWebhookTriggerWithContext(context.Background(), request)
}

// DeleteWebhookTrigger
// This API is used to delete a trigger.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
func (c *Client) DeleteWebhookTriggerWithContext(ctx context.Context, request *DeleteWebhookTriggerRequest) (response *DeleteWebhookTriggerResponse, err error) {
    if request == nil {
        request = NewDeleteWebhookTriggerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteWebhookTrigger require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteWebhookTriggerResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeChartDownloadInfoRequest() (request *DescribeChartDownloadInfoRequest) {
    request = &DescribeChartDownloadInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeChartDownloadInfo")
    
    
    return
}

func NewDescribeChartDownloadInfoResponse() (response *DescribeChartDownloadInfoResponse) {
    response = &DescribeChartDownloadInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeChartDownloadInfo
// This API is used to return the chart download information in an Enterprise Edition instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeChartDownloadInfo(request *DescribeChartDownloadInfoRequest) (response *DescribeChartDownloadInfoResponse, err error) {
    return c.DescribeChartDownloadInfoWithContext(context.Background(), request)
}

// DescribeChartDownloadInfo
// This API is used to return the chart download information in an Enterprise Edition instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeChartDownloadInfoWithContext(ctx context.Context, request *DescribeChartDownloadInfoRequest) (response *DescribeChartDownloadInfoResponse, err error) {
    if request == nil {
        request = NewDescribeChartDownloadInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeChartDownloadInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeChartDownloadInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomAccountsRequest() (request *DescribeCustomAccountsRequest) {
    request = &DescribeCustomAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeCustomAccounts")
    
    
    return
}

func NewDescribeCustomAccountsResponse() (response *DescribeCustomAccountsResponse) {
    response = &DescribeCustomAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCustomAccounts
// This API is used to query custom accounts.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCustomAccounts(request *DescribeCustomAccountsRequest) (response *DescribeCustomAccountsResponse, err error) {
    return c.DescribeCustomAccountsWithContext(context.Background(), request)
}

// DescribeCustomAccounts
// This API is used to query custom accounts.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCustomAccountsWithContext(ctx context.Context, request *DescribeCustomAccountsRequest) (response *DescribeCustomAccountsResponse, err error) {
    if request == nil {
        request = NewDescribeCustomAccountsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCustomAccounts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCustomAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExternalEndpointStatusRequest() (request *DescribeExternalEndpointStatusRequest) {
    request = &DescribeExternalEndpointStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeExternalEndpointStatus")
    
    
    return
}

func NewDescribeExternalEndpointStatusResponse() (response *DescribeExternalEndpointStatusResponse) {
    response = &DescribeExternalEndpointStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeExternalEndpointStatus
// This API is used to query the public network access entry status of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_ERRORGETDBDATAERROR = "FailedOperation.ErrorGetDBDataError"
//  FAILEDOPERATION_GETDBDATAERROR = "FailedOperation.GetDBDataError"
//  FAILEDOPERATION_GETTCRCLIENT = "FailedOperation.GetTcrClient"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeExternalEndpointStatus(request *DescribeExternalEndpointStatusRequest) (response *DescribeExternalEndpointStatusResponse, err error) {
    return c.DescribeExternalEndpointStatusWithContext(context.Background(), request)
}

// DescribeExternalEndpointStatus
// This API is used to query the public network access entry status of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_ERRORGETDBDATAERROR = "FailedOperation.ErrorGetDBDataError"
//  FAILEDOPERATION_GETDBDATAERROR = "FailedOperation.GetDBDataError"
//  FAILEDOPERATION_GETTCRCLIENT = "FailedOperation.GetTcrClient"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeExternalEndpointStatusWithContext(ctx context.Context, request *DescribeExternalEndpointStatusRequest) (response *DescribeExternalEndpointStatusResponse, err error) {
    if request == nil {
        request = NewDescribeExternalEndpointStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExternalEndpointStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExternalEndpointStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGCJobsRequest() (request *DescribeGCJobsRequest) {
    request = &DescribeGCJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeGCJobs")
    
    
    return
}

func NewDescribeGCJobsResponse() (response *DescribeGCJobsResponse) {
    response = &DescribeGCJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGCJobs
// This API is used to query the last ten garbage collection (GC) records.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeGCJobs(request *DescribeGCJobsRequest) (response *DescribeGCJobsResponse, err error) {
    return c.DescribeGCJobsWithContext(context.Background(), request)
}

// DescribeGCJobs
// This API is used to query the last ten garbage collection (GC) records.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeGCJobsWithContext(ctx context.Context, request *DescribeGCJobsRequest) (response *DescribeGCJobsResponse, err error) {
    if request == nil {
        request = NewDescribeGCJobsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGCJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGCJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageAccelerateServiceRequest() (request *DescribeImageAccelerateServiceRequest) {
    request = &DescribeImageAccelerateServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeImageAccelerateService")
    
    
    return
}

func NewDescribeImageAccelerateServiceResponse() (response *DescribeImageAccelerateServiceResponse) {
    response = &DescribeImageAccelerateServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeImageAccelerateService
// This API is used to query the status of an image acceleration service.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeImageAccelerateService(request *DescribeImageAccelerateServiceRequest) (response *DescribeImageAccelerateServiceResponse, err error) {
    return c.DescribeImageAccelerateServiceWithContext(context.Background(), request)
}

// DescribeImageAccelerateService
// This API is used to query the status of an image acceleration service.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeImageAccelerateServiceWithContext(ctx context.Context, request *DescribeImageAccelerateServiceRequest) (response *DescribeImageAccelerateServiceResponse, err error) {
    if request == nil {
        request = NewDescribeImageAccelerateServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImageAccelerateService require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImageAccelerateServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageManifestsRequest() (request *DescribeImageManifestsRequest) {
    request = &DescribeImageManifestsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeImageManifests")
    
    
    return
}

func NewDescribeImageManifestsResponse() (response *DescribeImageManifestsResponse) {
    response = &DescribeImageManifestsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeImageManifests
// This API is used to query the manifest information of a container image.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeImageManifests(request *DescribeImageManifestsRequest) (response *DescribeImageManifestsResponse, err error) {
    return c.DescribeImageManifestsWithContext(context.Background(), request)
}

// DescribeImageManifests
// This API is used to query the manifest information of a container image.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeImageManifestsWithContext(ctx context.Context, request *DescribeImageManifestsRequest) (response *DescribeImageManifestsResponse, err error) {
    if request == nil {
        request = NewDescribeImageManifestsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImageManifests require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImageManifestsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImagesRequest() (request *DescribeImagesRequest) {
    request = &DescribeImagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeImages")
    
    
    return
}

func NewDescribeImagesResponse() (response *DescribeImagesResponse) {
    response = &DescribeImagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeImages
// This API is used to query the list of image tags or the information of the specified container image.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
func (c *Client) DescribeImages(request *DescribeImagesRequest) (response *DescribeImagesResponse, err error) {
    return c.DescribeImagesWithContext(context.Background(), request)
}

// DescribeImages
// This API is used to query the list of image tags or the information of the specified container image.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
func (c *Client) DescribeImagesWithContext(ctx context.Context, request *DescribeImagesRequest) (response *DescribeImagesResponse, err error) {
    if request == nil {
        request = NewDescribeImagesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImages require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImagesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImmutableTagRulesRequest() (request *DescribeImmutableTagRulesRequest) {
    request = &DescribeImmutableTagRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeImmutableTagRules")
    
    
    return
}

func NewDescribeImmutableTagRulesResponse() (response *DescribeImmutableTagRulesResponse) {
    response = &DescribeImmutableTagRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeImmutableTagRules
// This API is used to list the tag immutability rule.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
func (c *Client) DescribeImmutableTagRules(request *DescribeImmutableTagRulesRequest) (response *DescribeImmutableTagRulesResponse, err error) {
    return c.DescribeImmutableTagRulesWithContext(context.Background(), request)
}

// DescribeImmutableTagRules
// This API is used to list the tag immutability rule.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
func (c *Client) DescribeImmutableTagRulesWithContext(ctx context.Context, request *DescribeImmutableTagRulesRequest) (response *DescribeImmutableTagRulesResponse, err error) {
    if request == nil {
        request = NewDescribeImmutableTagRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImmutableTagRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImmutableTagRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceAllNamespacesRequest() (request *DescribeInstanceAllNamespacesRequest) {
    request = &DescribeInstanceAllNamespacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeInstanceAllNamespaces")
    
    
    return
}

func NewDescribeInstanceAllNamespacesResponse() (response *DescribeInstanceAllNamespacesResponse) {
    response = &DescribeInstanceAllNamespacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceAllNamespaces
// This API is used to query the list of all namespaces in an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_DEPENDENCEERROR = "FailedOperation.DependenceError"
//  FAILEDOPERATION_ERRORGETDBDATAERROR = "FailedOperation.ErrorGetDBDataError"
//  FAILEDOPERATION_OPERATIONCANCEL = "FailedOperation.OperationCancel"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInstanceAllNamespaces(request *DescribeInstanceAllNamespacesRequest) (response *DescribeInstanceAllNamespacesResponse, err error) {
    return c.DescribeInstanceAllNamespacesWithContext(context.Background(), request)
}

// DescribeInstanceAllNamespaces
// This API is used to query the list of all namespaces in an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_DEPENDENCEERROR = "FailedOperation.DependenceError"
//  FAILEDOPERATION_ERRORGETDBDATAERROR = "FailedOperation.ErrorGetDBDataError"
//  FAILEDOPERATION_OPERATIONCANCEL = "FailedOperation.OperationCancel"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInstanceAllNamespacesWithContext(ctx context.Context, request *DescribeInstanceAllNamespacesRequest) (response *DescribeInstanceAllNamespacesResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceAllNamespacesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceAllNamespaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceAllNamespacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceCustomizedDomainRequest() (request *DescribeInstanceCustomizedDomainRequest) {
    request = &DescribeInstanceCustomizedDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeInstanceCustomizedDomain")
    
    
    return
}

func NewDescribeInstanceCustomizedDomainResponse() (response *DescribeInstanceCustomizedDomainResponse) {
    response = &DescribeInstanceCustomizedDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceCustomizedDomain
// This API is used to query the list of custom domain names of an instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTEDREGION = "InvalidParameter.UnsupportedRegion"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInstanceCustomizedDomain(request *DescribeInstanceCustomizedDomainRequest) (response *DescribeInstanceCustomizedDomainResponse, err error) {
    return c.DescribeInstanceCustomizedDomainWithContext(context.Background(), request)
}

// DescribeInstanceCustomizedDomain
// This API is used to query the list of custom domain names of an instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTEDREGION = "InvalidParameter.UnsupportedRegion"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInstanceCustomizedDomainWithContext(ctx context.Context, request *DescribeInstanceCustomizedDomainRequest) (response *DescribeInstanceCustomizedDomainResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceCustomizedDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceCustomizedDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceCustomizedDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceStatusRequest() (request *DescribeInstanceStatusRequest) {
    request = &DescribeInstanceStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeInstanceStatus")
    
    
    return
}

func NewDescribeInstanceStatusResponse() (response *DescribeInstanceStatusResponse) {
    response = &DescribeInstanceStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceStatus
// This API is used to query the current status and process information of an instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInstanceStatus(request *DescribeInstanceStatusRequest) (response *DescribeInstanceStatusResponse, err error) {
    return c.DescribeInstanceStatusWithContext(context.Background(), request)
}

// DescribeInstanceStatus
// This API is used to query the current status and process information of an instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInstanceStatusWithContext(ctx context.Context, request *DescribeInstanceStatusRequest) (response *DescribeInstanceStatusResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceTokenRequest() (request *DescribeInstanceTokenRequest) {
    request = &DescribeInstanceTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeInstanceToken")
    
    
    return
}

func NewDescribeInstanceTokenResponse() (response *DescribeInstanceTokenResponse) {
    response = &DescribeInstanceTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceToken
// This API is used to query the information of long-term access credentials.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInstanceToken(request *DescribeInstanceTokenRequest) (response *DescribeInstanceTokenResponse, err error) {
    return c.DescribeInstanceTokenWithContext(context.Background(), request)
}

// DescribeInstanceToken
// This API is used to query the information of long-term access credentials.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInstanceTokenWithContext(ctx context.Context, request *DescribeInstanceTokenRequest) (response *DescribeInstanceTokenResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceTokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceTokenResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeInstances")
    
    
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstances
// This API is used to query the instance information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINSUFFICIENT_ERRORINSTANCENOTRUNNING = "ResourceInsufficient.ErrorInstanceNotRunning"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    return c.DescribeInstancesWithContext(context.Background(), request)
}

// DescribeInstances
// This API is used to query the instance information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINSUFFICIENT_ERRORINSTANCENOTRUNNING = "ResourceInsufficient.ErrorInstanceNotRunning"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInstancesWithContext(ctx context.Context, request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInternalEndpointsRequest() (request *DescribeInternalEndpointsRequest) {
    request = &DescribeInternalEndpointsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeInternalEndpoints")
    
    
    return
}

func NewDescribeInternalEndpointsResponse() (response *DescribeInternalEndpointsResponse) {
    response = &DescribeInternalEndpointsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInternalEndpoints
// This API is used to query the VPC URLs for private network access to an instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInternalEndpoints(request *DescribeInternalEndpointsRequest) (response *DescribeInternalEndpointsResponse, err error) {
    return c.DescribeInternalEndpointsWithContext(context.Background(), request)
}

// DescribeInternalEndpoints
// This API is used to query the VPC URLs for private network access to an instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInternalEndpointsWithContext(ctx context.Context, request *DescribeInternalEndpointsRequest) (response *DescribeInternalEndpointsResponse, err error) {
    if request == nil {
        request = NewDescribeInternalEndpointsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInternalEndpoints require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInternalEndpointsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNamespacesRequest() (request *DescribeNamespacesRequest) {
    request = &DescribeNamespacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeNamespaces")
    
    
    return
}

func NewDescribeNamespacesResponse() (response *DescribeNamespacesResponse) {
    response = &DescribeNamespacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNamespaces
// This API is used to query the namespace list or the information of the specified namespace.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DEPENDENCEERROR = "FailedOperation.DependenceError"
//  FAILEDOPERATION_ERRORGETDBDATAERROR = "FailedOperation.ErrorGetDBDataError"
//  FAILEDOPERATION_ERRORTCRRESOURCECONFLICT = "FailedOperation.ErrorTcrResourceConflict"
//  FAILEDOPERATION_ERRORTCRUNAUTHORIZED = "FailedOperation.ErrorTcrUnauthorized"
//  FAILEDOPERATION_GETDBDATAERROR = "FailedOperation.GetDBDataError"
//  FAILEDOPERATION_OPERATIONCANCEL = "FailedOperation.OperationCancel"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeNamespaces(request *DescribeNamespacesRequest) (response *DescribeNamespacesResponse, err error) {
    return c.DescribeNamespacesWithContext(context.Background(), request)
}

// DescribeNamespaces
// This API is used to query the namespace list or the information of the specified namespace.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DEPENDENCEERROR = "FailedOperation.DependenceError"
//  FAILEDOPERATION_ERRORGETDBDATAERROR = "FailedOperation.ErrorGetDBDataError"
//  FAILEDOPERATION_ERRORTCRRESOURCECONFLICT = "FailedOperation.ErrorTcrResourceConflict"
//  FAILEDOPERATION_ERRORTCRUNAUTHORIZED = "FailedOperation.ErrorTcrUnauthorized"
//  FAILEDOPERATION_GETDBDATAERROR = "FailedOperation.GetDBDataError"
//  FAILEDOPERATION_OPERATIONCANCEL = "FailedOperation.OperationCancel"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeNamespacesWithContext(ctx context.Context, request *DescribeNamespacesRequest) (response *DescribeNamespacesResponse, err error) {
    if request == nil {
        request = NewDescribeNamespacesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNamespaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNamespacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRegionsRequest() (request *DescribeRegionsRequest) {
    request = &DescribeRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeRegions")
    
    
    return
}

func NewDescribeRegionsResponse() (response *DescribeRegionsResponse) {
    response = &DescribeRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRegions
// This API is used to get the available regions in TCR.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRegions(request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
    return c.DescribeRegionsWithContext(context.Background(), request)
}

// DescribeRegions
// This API is used to get the available regions in TCR.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRegionsWithContext(ctx context.Context, request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeRegionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRegions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReplicationInstanceCreateTasksRequest() (request *DescribeReplicationInstanceCreateTasksRequest) {
    request = &DescribeReplicationInstanceCreateTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeReplicationInstanceCreateTasks")
    
    
    return
}

func NewDescribeReplicationInstanceCreateTasksResponse() (response *DescribeReplicationInstanceCreateTasksResponse) {
    response = &DescribeReplicationInstanceCreateTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeReplicationInstanceCreateTasks
// This API is used to query the task status of creating a replication instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeReplicationInstanceCreateTasks(request *DescribeReplicationInstanceCreateTasksRequest) (response *DescribeReplicationInstanceCreateTasksResponse, err error) {
    return c.DescribeReplicationInstanceCreateTasksWithContext(context.Background(), request)
}

// DescribeReplicationInstanceCreateTasks
// This API is used to query the task status of creating a replication instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeReplicationInstanceCreateTasksWithContext(ctx context.Context, request *DescribeReplicationInstanceCreateTasksRequest) (response *DescribeReplicationInstanceCreateTasksResponse, err error) {
    if request == nil {
        request = NewDescribeReplicationInstanceCreateTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReplicationInstanceCreateTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReplicationInstanceCreateTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReplicationInstanceSyncStatusRequest() (request *DescribeReplicationInstanceSyncStatusRequest) {
    request = &DescribeReplicationInstanceSyncStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeReplicationInstanceSyncStatus")
    
    
    return
}

func NewDescribeReplicationInstanceSyncStatusResponse() (response *DescribeReplicationInstanceSyncStatusResponse) {
    response = &DescribeReplicationInstanceSyncStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeReplicationInstanceSyncStatus
// This API is used to query the synchronization status of a replication instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeReplicationInstanceSyncStatus(request *DescribeReplicationInstanceSyncStatusRequest) (response *DescribeReplicationInstanceSyncStatusResponse, err error) {
    return c.DescribeReplicationInstanceSyncStatusWithContext(context.Background(), request)
}

// DescribeReplicationInstanceSyncStatus
// This API is used to query the synchronization status of a replication instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeReplicationInstanceSyncStatusWithContext(ctx context.Context, request *DescribeReplicationInstanceSyncStatusRequest) (response *DescribeReplicationInstanceSyncStatusResponse, err error) {
    if request == nil {
        request = NewDescribeReplicationInstanceSyncStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReplicationInstanceSyncStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReplicationInstanceSyncStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReplicationInstancesRequest() (request *DescribeReplicationInstancesRequest) {
    request = &DescribeReplicationInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeReplicationInstances")
    
    
    return
}

func NewDescribeReplicationInstancesResponse() (response *DescribeReplicationInstancesResponse) {
    response = &DescribeReplicationInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeReplicationInstances
// This API is used to query the list of replication instances.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINSUFFICIENT_ERRORINSTANCENOTRUNNING = "ResourceInsufficient.ErrorInstanceNotRunning"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeReplicationInstances(request *DescribeReplicationInstancesRequest) (response *DescribeReplicationInstancesResponse, err error) {
    return c.DescribeReplicationInstancesWithContext(context.Background(), request)
}

// DescribeReplicationInstances
// This API is used to query the list of replication instances.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINSUFFICIENT_ERRORINSTANCENOTRUNNING = "ResourceInsufficient.ErrorInstanceNotRunning"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeReplicationInstancesWithContext(ctx context.Context, request *DescribeReplicationInstancesRequest) (response *DescribeReplicationInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeReplicationInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReplicationInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReplicationInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRepositoriesRequest() (request *DescribeRepositoriesRequest) {
    request = &DescribeRepositoriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeRepositories")
    
    
    return
}

func NewDescribeRepositoriesResponse() (response *DescribeRepositoriesResponse) {
    response = &DescribeRepositoriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRepositories
// This API is used to query the image repository list or the information of the specified image repository.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EMPTYCOREBODY = "FailedOperation.EmptyCoreBody"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRepositories(request *DescribeRepositoriesRequest) (response *DescribeRepositoriesResponse, err error) {
    return c.DescribeRepositoriesWithContext(context.Background(), request)
}

// DescribeRepositories
// This API is used to query the image repository list or the information of the specified image repository.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EMPTYCOREBODY = "FailedOperation.EmptyCoreBody"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRepositoriesWithContext(ctx context.Context, request *DescribeRepositoriesRequest) (response *DescribeRepositoriesResponse, err error) {
    if request == nil {
        request = NewDescribeRepositoriesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRepositories require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRepositoriesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityPoliciesRequest() (request *DescribeSecurityPoliciesRequest) {
    request = &DescribeSecurityPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeSecurityPolicies")
    
    
    return
}

func NewDescribeSecurityPoliciesResponse() (response *DescribeSecurityPoliciesResponse) {
    response = &DescribeSecurityPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecurityPolicies
// This API is used to query the public network access allowlist policies of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_ERRORGETDBDATAERROR = "FailedOperation.ErrorGetDBDataError"
//  FAILEDOPERATION_GETDBDATAERROR = "FailedOperation.GetDBDataError"
//  FAILEDOPERATION_GETSECURITYPOLICYFAIL = "FailedOperation.GetSecurityPolicyFail"
//  FAILEDOPERATION_GETTCRCLIENT = "FailedOperation.GetTcrClient"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSecurityPolicies(request *DescribeSecurityPoliciesRequest) (response *DescribeSecurityPoliciesResponse, err error) {
    return c.DescribeSecurityPoliciesWithContext(context.Background(), request)
}

// DescribeSecurityPolicies
// This API is used to query the public network access allowlist policies of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_ERRORGETDBDATAERROR = "FailedOperation.ErrorGetDBDataError"
//  FAILEDOPERATION_GETDBDATAERROR = "FailedOperation.GetDBDataError"
//  FAILEDOPERATION_GETSECURITYPOLICYFAIL = "FailedOperation.GetSecurityPolicyFail"
//  FAILEDOPERATION_GETTCRCLIENT = "FailedOperation.GetTcrClient"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSecurityPoliciesWithContext(ctx context.Context, request *DescribeSecurityPoliciesRequest) (response *DescribeSecurityPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityPoliciesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceAccountsRequest() (request *DescribeServiceAccountsRequest) {
    request = &DescribeServiceAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeServiceAccounts")
    
    
    return
}

func NewDescribeServiceAccountsResponse() (response *DescribeServiceAccountsResponse) {
    response = &DescribeServiceAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeServiceAccounts
// This API is used to query service accounts.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeServiceAccounts(request *DescribeServiceAccountsRequest) (response *DescribeServiceAccountsResponse, err error) {
    return c.DescribeServiceAccountsWithContext(context.Background(), request)
}

// DescribeServiceAccounts
// This API is used to query service accounts.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeServiceAccountsWithContext(ctx context.Context, request *DescribeServiceAccountsRequest) (response *DescribeServiceAccountsResponse, err error) {
    if request == nil {
        request = NewDescribeServiceAccountsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServiceAccounts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServiceAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTagRetentionExecutionRequest() (request *DescribeTagRetentionExecutionRequest) {
    request = &DescribeTagRetentionExecutionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeTagRetentionExecution")
    
    
    return
}

func NewDescribeTagRetentionExecutionResponse() (response *DescribeTagRetentionExecutionResponse) {
    response = &DescribeTagRetentionExecutionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTagRetentionExecution
// This API is used to query tag retention execution records.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTagRetentionExecution(request *DescribeTagRetentionExecutionRequest) (response *DescribeTagRetentionExecutionResponse, err error) {
    return c.DescribeTagRetentionExecutionWithContext(context.Background(), request)
}

// DescribeTagRetentionExecution
// This API is used to query tag retention execution records.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTagRetentionExecutionWithContext(ctx context.Context, request *DescribeTagRetentionExecutionRequest) (response *DescribeTagRetentionExecutionResponse, err error) {
    if request == nil {
        request = NewDescribeTagRetentionExecutionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTagRetentionExecution require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTagRetentionExecutionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTagRetentionExecutionTaskRequest() (request *DescribeTagRetentionExecutionTaskRequest) {
    request = &DescribeTagRetentionExecutionTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeTagRetentionExecutionTask")
    
    
    return
}

func NewDescribeTagRetentionExecutionTaskResponse() (response *DescribeTagRetentionExecutionTaskResponse) {
    response = &DescribeTagRetentionExecutionTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTagRetentionExecutionTask
// This API is used to query tag retention execution tasks.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTagRetentionExecutionTask(request *DescribeTagRetentionExecutionTaskRequest) (response *DescribeTagRetentionExecutionTaskResponse, err error) {
    return c.DescribeTagRetentionExecutionTaskWithContext(context.Background(), request)
}

// DescribeTagRetentionExecutionTask
// This API is used to query tag retention execution tasks.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTagRetentionExecutionTaskWithContext(ctx context.Context, request *DescribeTagRetentionExecutionTaskRequest) (response *DescribeTagRetentionExecutionTaskResponse, err error) {
    if request == nil {
        request = NewDescribeTagRetentionExecutionTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTagRetentionExecutionTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTagRetentionExecutionTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTagRetentionRulesRequest() (request *DescribeTagRetentionRulesRequest) {
    request = &DescribeTagRetentionRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeTagRetentionRules")
    
    
    return
}

func NewDescribeTagRetentionRulesResponse() (response *DescribeTagRetentionRulesResponse) {
    response = &DescribeTagRetentionRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTagRetentionRules
// This API is used to query tag retention rules.
//
// error code that may be returned:
//  FAILEDOPERATION_EMPTYCOREBODY = "FailedOperation.EmptyCoreBody"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTagRetentionRules(request *DescribeTagRetentionRulesRequest) (response *DescribeTagRetentionRulesResponse, err error) {
    return c.DescribeTagRetentionRulesWithContext(context.Background(), request)
}

// DescribeTagRetentionRules
// This API is used to query tag retention rules.
//
// error code that may be returned:
//  FAILEDOPERATION_EMPTYCOREBODY = "FailedOperation.EmptyCoreBody"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTagRetentionRulesWithContext(ctx context.Context, request *DescribeTagRetentionRulesRequest) (response *DescribeTagRetentionRulesResponse, err error) {
    if request == nil {
        request = NewDescribeTagRetentionRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTagRetentionRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTagRetentionRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWebhookTriggerRequest() (request *DescribeWebhookTriggerRequest) {
    request = &DescribeWebhookTriggerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeWebhookTrigger")
    
    
    return
}

func NewDescribeWebhookTriggerResponse() (response *DescribeWebhookTriggerResponse) {
    response = &DescribeWebhookTriggerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWebhookTrigger
// This API is used to query triggers.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
func (c *Client) DescribeWebhookTrigger(request *DescribeWebhookTriggerRequest) (response *DescribeWebhookTriggerResponse, err error) {
    return c.DescribeWebhookTriggerWithContext(context.Background(), request)
}

// DescribeWebhookTrigger
// This API is used to query triggers.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
func (c *Client) DescribeWebhookTriggerWithContext(ctx context.Context, request *DescribeWebhookTriggerRequest) (response *DescribeWebhookTriggerResponse, err error) {
    if request == nil {
        request = NewDescribeWebhookTriggerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWebhookTrigger require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWebhookTriggerResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWebhookTriggerLogRequest() (request *DescribeWebhookTriggerLogRequest) {
    request = &DescribeWebhookTriggerLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeWebhookTriggerLog")
    
    
    return
}

func NewDescribeWebhookTriggerLogResponse() (response *DescribeWebhookTriggerLogResponse) {
    response = &DescribeWebhookTriggerLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWebhookTriggerLog
// This API is used to get trigger logs.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
func (c *Client) DescribeWebhookTriggerLog(request *DescribeWebhookTriggerLogRequest) (response *DescribeWebhookTriggerLogResponse, err error) {
    return c.DescribeWebhookTriggerLogWithContext(context.Background(), request)
}

// DescribeWebhookTriggerLog
// This API is used to get trigger logs.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
func (c *Client) DescribeWebhookTriggerLogWithContext(ctx context.Context, request *DescribeWebhookTriggerLogRequest) (response *DescribeWebhookTriggerLogResponse, err error) {
    if request == nil {
        request = NewDescribeWebhookTriggerLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWebhookTriggerLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWebhookTriggerLogResponse()
    err = c.Send(request, response)
    return
}

func NewDownloadHelmChartRequest() (request *DownloadHelmChartRequest) {
    request = &DownloadHelmChartRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DownloadHelmChart")
    
    
    return
}

func NewDownloadHelmChartResponse() (response *DownloadHelmChartResponse) {
    response = &DownloadHelmChartResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DownloadHelmChart
// This API is used to download a Helm chart in TCR.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DownloadHelmChart(request *DownloadHelmChartRequest) (response *DownloadHelmChartResponse, err error) {
    return c.DownloadHelmChartWithContext(context.Background(), request)
}

// DownloadHelmChart
// This API is used to download a Helm chart in TCR.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DownloadHelmChartWithContext(ctx context.Context, request *DownloadHelmChartRequest) (response *DownloadHelmChartResponse, err error) {
    if request == nil {
        request = NewDownloadHelmChartRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DownloadHelmChart require credential")
    }

    request.SetContext(ctx)
    
    response = NewDownloadHelmChartResponse()
    err = c.Send(request, response)
    return
}

func NewManageExternalEndpointRequest() (request *ManageExternalEndpointRequest) {
    request = &ManageExternalEndpointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "ManageExternalEndpoint")
    
    
    return
}

func NewManageExternalEndpointResponse() (response *ManageExternalEndpointResponse) {
    response = &ManageExternalEndpointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ManageExternalEndpoint
// This API is used to manage the public network access of an instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ManageExternalEndpoint(request *ManageExternalEndpointRequest) (response *ManageExternalEndpointResponse, err error) {
    return c.ManageExternalEndpointWithContext(context.Background(), request)
}

// ManageExternalEndpoint
// This API is used to manage the public network access of an instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ManageExternalEndpointWithContext(ctx context.Context, request *ManageExternalEndpointRequest) (response *ManageExternalEndpointResponse, err error) {
    if request == nil {
        request = NewManageExternalEndpointRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ManageExternalEndpoint require credential")
    }

    request.SetContext(ctx)
    
    response = NewManageExternalEndpointResponse()
    err = c.Send(request, response)
    return
}

func NewManageInternalEndpointRequest() (request *ManageInternalEndpointRequest) {
    request = &ManageInternalEndpointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "ManageInternalEndpoint")
    
    
    return
}

func NewManageInternalEndpointResponse() (response *ManageInternalEndpointResponse) {
    response = &ManageInternalEndpointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ManageInternalEndpoint
// This API is used to manage VPC URLs for private network access to an instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINSUFFICIENT_ERRORVPCDNSSTATUS = "ResourceInsufficient.ErrorVpcDnsStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ManageInternalEndpoint(request *ManageInternalEndpointRequest) (response *ManageInternalEndpointResponse, err error) {
    return c.ManageInternalEndpointWithContext(context.Background(), request)
}

// ManageInternalEndpoint
// This API is used to manage VPC URLs for private network access to an instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINSUFFICIENT_ERRORVPCDNSSTATUS = "ResourceInsufficient.ErrorVpcDnsStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ManageInternalEndpointWithContext(ctx context.Context, request *ManageInternalEndpointRequest) (response *ManageInternalEndpointResponse, err error) {
    if request == nil {
        request = NewManageInternalEndpointRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ManageInternalEndpoint require credential")
    }

    request.SetContext(ctx)
    
    response = NewManageInternalEndpointResponse()
    err = c.Send(request, response)
    return
}

func NewManageReplicationRequest() (request *ManageReplicationRequest) {
    request = &ManageReplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "ManageReplication")
    
    
    return
}

func NewManageReplicationResponse() (response *ManageReplicationResponse) {
    response = &ManageReplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ManageReplication
// This API is used to manage the instance synchronization rule.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ManageReplication(request *ManageReplicationRequest) (response *ManageReplicationResponse, err error) {
    return c.ManageReplicationWithContext(context.Background(), request)
}

// ManageReplication
// This API is used to manage the instance synchronization rule.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ManageReplicationWithContext(ctx context.Context, request *ManageReplicationRequest) (response *ManageReplicationResponse, err error) {
    if request == nil {
        request = NewManageReplicationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ManageReplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewManageReplicationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCustomAccountRequest() (request *ModifyCustomAccountRequest) {
    request = &ModifyCustomAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "ModifyCustomAccount")
    
    
    return
}

func NewModifyCustomAccountResponse() (response *ModifyCustomAccountResponse) {
    response = &ModifyCustomAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCustomAccount
// This API is used to update a custom account.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyCustomAccount(request *ModifyCustomAccountRequest) (response *ModifyCustomAccountResponse, err error) {
    return c.ModifyCustomAccountWithContext(context.Background(), request)
}

// ModifyCustomAccount
// This API is used to update a custom account.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyCustomAccountWithContext(ctx context.Context, request *ModifyCustomAccountRequest) (response *ModifyCustomAccountResponse, err error) {
    if request == nil {
        request = NewModifyCustomAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCustomAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCustomAccountResponse()
    err = c.Send(request, response)
    return
}

func NewModifyImmutableTagRulesRequest() (request *ModifyImmutableTagRulesRequest) {
    request = &ModifyImmutableTagRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "ModifyImmutableTagRules")
    
    
    return
}

func NewModifyImmutableTagRulesResponse() (response *ModifyImmutableTagRulesResponse) {
    response = &ModifyImmutableTagRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyImmutableTagRules
// This API is used to update the tag immutability rule.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
func (c *Client) ModifyImmutableTagRules(request *ModifyImmutableTagRulesRequest) (response *ModifyImmutableTagRulesResponse, err error) {
    return c.ModifyImmutableTagRulesWithContext(context.Background(), request)
}

// ModifyImmutableTagRules
// This API is used to update the tag immutability rule.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
func (c *Client) ModifyImmutableTagRulesWithContext(ctx context.Context, request *ModifyImmutableTagRulesRequest) (response *ModifyImmutableTagRulesResponse, err error) {
    if request == nil {
        request = NewModifyImmutableTagRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyImmutableTagRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyImmutableTagRulesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceRequest() (request *ModifyInstanceRequest) {
    request = &ModifyInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "ModifyInstance")
    
    
    return
}

func NewModifyInstanceResponse() (response *ModifyInstanceResponse) {
    response = &ModifyInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstance
// This API is used to update instance information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyInstance(request *ModifyInstanceRequest) (response *ModifyInstanceResponse, err error) {
    return c.ModifyInstanceWithContext(context.Background(), request)
}

// ModifyInstance
// This API is used to update instance information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyInstanceWithContext(ctx context.Context, request *ModifyInstanceRequest) (response *ModifyInstanceResponse, err error) {
    if request == nil {
        request = NewModifyInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceTokenRequest() (request *ModifyInstanceTokenRequest) {
    request = &ModifyInstanceTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "ModifyInstanceToken")
    
    
    return
}

func NewModifyInstanceTokenResponse() (response *ModifyInstanceTokenResponse) {
    response = &ModifyInstanceTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstanceToken
// This API is used to update the status of the specified long-term access credential in an instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyInstanceToken(request *ModifyInstanceTokenRequest) (response *ModifyInstanceTokenResponse, err error) {
    return c.ModifyInstanceTokenWithContext(context.Background(), request)
}

// ModifyInstanceToken
// This API is used to update the status of the specified long-term access credential in an instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyInstanceTokenWithContext(ctx context.Context, request *ModifyInstanceTokenRequest) (response *ModifyInstanceTokenResponse, err error) {
    if request == nil {
        request = NewModifyInstanceTokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceTokenResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNamespaceRequest() (request *ModifyNamespaceRequest) {
    request = &ModifyNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "ModifyNamespace")
    
    
    return
}

func NewModifyNamespaceResponse() (response *ModifyNamespaceResponse) {
    response = &ModifyNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyNamespace
// This API is used to update the information of a namespace. Currently, only the namespace access level can be modified.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyNamespace(request *ModifyNamespaceRequest) (response *ModifyNamespaceResponse, err error) {
    return c.ModifyNamespaceWithContext(context.Background(), request)
}

// ModifyNamespace
// This API is used to update the information of a namespace. Currently, only the namespace access level can be modified.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyNamespaceWithContext(ctx context.Context, request *ModifyNamespaceRequest) (response *ModifyNamespaceResponse, err error) {
    if request == nil {
        request = NewModifyNamespaceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNamespace require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRepositoryRequest() (request *ModifyRepositoryRequest) {
    request = &ModifyRepositoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "ModifyRepository")
    
    
    return
}

func NewModifyRepositoryResponse() (response *ModifyRepositoryResponse) {
    response = &ModifyRepositoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRepository
// This API is used to update the information of an image repository. The repository description can be modified.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyRepository(request *ModifyRepositoryRequest) (response *ModifyRepositoryResponse, err error) {
    return c.ModifyRepositoryWithContext(context.Background(), request)
}

// ModifyRepository
// This API is used to update the information of an image repository. The repository description can be modified.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyRepositoryWithContext(ctx context.Context, request *ModifyRepositoryRequest) (response *ModifyRepositoryResponse, err error) {
    if request == nil {
        request = NewModifyRepositoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRepository require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRepositoryResponse()
    err = c.Send(request, response)
    return
}

func NewModifySecurityPolicyRequest() (request *ModifySecurityPolicyRequest) {
    request = &ModifySecurityPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "ModifySecurityPolicy")
    
    
    return
}

func NewModifySecurityPolicyResponse() (response *ModifySecurityPolicyResponse) {
    response = &ModifySecurityPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySecurityPolicy
// This API is used to update the public network access allowlist of an instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifySecurityPolicy(request *ModifySecurityPolicyRequest) (response *ModifySecurityPolicyResponse, err error) {
    return c.ModifySecurityPolicyWithContext(context.Background(), request)
}

// ModifySecurityPolicy
// This API is used to update the public network access allowlist of an instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifySecurityPolicyWithContext(ctx context.Context, request *ModifySecurityPolicyRequest) (response *ModifySecurityPolicyResponse, err error) {
    if request == nil {
        request = NewModifySecurityPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySecurityPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySecurityPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyServiceAccountRequest() (request *ModifyServiceAccountRequest) {
    request = &ModifyServiceAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "ModifyServiceAccount")
    
    
    return
}

func NewModifyServiceAccountResponse() (response *ModifyServiceAccountResponse) {
    response = &ModifyServiceAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyServiceAccount
// This API is used to update a service account.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyServiceAccount(request *ModifyServiceAccountRequest) (response *ModifyServiceAccountResponse, err error) {
    return c.ModifyServiceAccountWithContext(context.Background(), request)
}

// ModifyServiceAccount
// This API is used to update a service account.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyServiceAccountWithContext(ctx context.Context, request *ModifyServiceAccountRequest) (response *ModifyServiceAccountResponse, err error) {
    if request == nil {
        request = NewModifyServiceAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyServiceAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyServiceAccountResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTagRetentionRuleRequest() (request *ModifyTagRetentionRuleRequest) {
    request = &ModifyTagRetentionRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "ModifyTagRetentionRule")
    
    
    return
}

func NewModifyTagRetentionRuleResponse() (response *ModifyTagRetentionRuleResponse) {
    response = &ModifyTagRetentionRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTagRetentionRule
// This API is used to update a tag retention rule.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyTagRetentionRule(request *ModifyTagRetentionRuleRequest) (response *ModifyTagRetentionRuleResponse, err error) {
    return c.ModifyTagRetentionRuleWithContext(context.Background(), request)
}

// ModifyTagRetentionRule
// This API is used to update a tag retention rule.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyTagRetentionRuleWithContext(ctx context.Context, request *ModifyTagRetentionRuleRequest) (response *ModifyTagRetentionRuleResponse, err error) {
    if request == nil {
        request = NewModifyTagRetentionRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTagRetentionRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTagRetentionRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWebhookTriggerRequest() (request *ModifyWebhookTriggerRequest) {
    request = &ModifyWebhookTriggerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "ModifyWebhookTrigger")
    
    
    return
}

func NewModifyWebhookTriggerResponse() (response *ModifyWebhookTriggerResponse) {
    response = &ModifyWebhookTriggerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyWebhookTrigger
// This API is used to update a trigger.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
func (c *Client) ModifyWebhookTrigger(request *ModifyWebhookTriggerRequest) (response *ModifyWebhookTriggerResponse, err error) {
    return c.ModifyWebhookTriggerWithContext(context.Background(), request)
}

// ModifyWebhookTrigger
// This API is used to update a trigger.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
func (c *Client) ModifyWebhookTriggerWithContext(ctx context.Context, request *ModifyWebhookTriggerRequest) (response *ModifyWebhookTriggerResponse, err error) {
    if request == nil {
        request = NewModifyWebhookTriggerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyWebhookTrigger require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyWebhookTriggerResponse()
    err = c.Send(request, response)
    return
}

func NewRenewInstanceRequest() (request *RenewInstanceRequest) {
    request = &RenewInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "RenewInstance")
    
    
    return
}

func NewRenewInstanceResponse() (response *RenewInstanceResponse) {
    response = &RenewInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RenewInstance
// This API is used to renew a prepaid instance or change the billing mode from pay-as-you-go billing to monthly subscription billing.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORNAMEEXISTS = "InvalidParameter.ErrorNameExists"
//  INVALIDPARAMETER_ERRORREGISTRYNAME = "InvalidParameter.ErrorRegistryName"
//  INVALIDPARAMETER_ERRORTAGOVERLIMIT = "InvalidParameter.ErrorTagOverLimit"
//  INVALIDPARAMETER_UNSUPPORTEDREGION = "InvalidParameter.UnsupportedRegion"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RenewInstance(request *RenewInstanceRequest) (response *RenewInstanceResponse, err error) {
    return c.RenewInstanceWithContext(context.Background(), request)
}

// RenewInstance
// This API is used to renew a prepaid instance or change the billing mode from pay-as-you-go billing to monthly subscription billing.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORNAMEEXISTS = "InvalidParameter.ErrorNameExists"
//  INVALIDPARAMETER_ERRORREGISTRYNAME = "InvalidParameter.ErrorRegistryName"
//  INVALIDPARAMETER_ERRORTAGOVERLIMIT = "InvalidParameter.ErrorTagOverLimit"
//  INVALIDPARAMETER_UNSUPPORTEDREGION = "InvalidParameter.UnsupportedRegion"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RenewInstanceWithContext(ctx context.Context, request *RenewInstanceRequest) (response *RenewInstanceResponse, err error) {
    if request == nil {
        request = NewRenewInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RenewInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewRenewInstanceResponse()
    err = c.Send(request, response)
    return
}
