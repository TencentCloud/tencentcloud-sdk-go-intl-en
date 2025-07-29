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

package v20220331

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2022-03-31"

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


func NewCreateApiImportUserJobRequest() (request *CreateApiImportUserJobRequest) {
    request = &CreateApiImportUserJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ciam", APIVersion, "CreateApiImportUserJob")
    
    
    return
}

func NewCreateApiImportUserJobResponse() (response *CreateApiImportUserJobResponse) {
    response = &CreateApiImportUserJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateApiImportUserJob
// This API is used to create an API for user import task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATAFLOWTOOMANYREQUESTS = "FailedOperation.DataFlowTooManyRequests"
//  FAILEDOPERATION_IMPORTUSERISEMPTY = "FailedOperation.ImportUserIsEmpty"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateApiImportUserJob(request *CreateApiImportUserJobRequest) (response *CreateApiImportUserJobResponse, err error) {
    return c.CreateApiImportUserJobWithContext(context.Background(), request)
}

// CreateApiImportUserJob
// This API is used to create an API for user import task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATAFLOWTOOMANYREQUESTS = "FailedOperation.DataFlowTooManyRequests"
//  FAILEDOPERATION_IMPORTUSERISEMPTY = "FailedOperation.ImportUserIsEmpty"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateApiImportUserJobWithContext(ctx context.Context, request *CreateApiImportUserJobRequest) (response *CreateApiImportUserJobResponse, err error) {
    if request == nil {
        request = NewCreateApiImportUserJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ciam", APIVersion, "CreateApiImportUserJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApiImportUserJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApiImportUserJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFileExportUserJobRequest() (request *CreateFileExportUserJobRequest) {
    request = &CreateFileExportUserJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ciam", APIVersion, "CreateFileExportUserJob")
    
    
    return
}

func NewCreateFileExportUserJobResponse() (response *CreateFileExportUserJobResponse) {
    response = &CreateFileExportUserJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateFileExportUserJob
// This API is used to create a file for user export task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATAFLOWTOOMANYREQUESTS = "FailedOperation.DataFlowTooManyRequests"
//  FAILEDOPERATION_IMPORTUSERISEMPTY = "FailedOperation.ImportUserIsEmpty"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateFileExportUserJob(request *CreateFileExportUserJobRequest) (response *CreateFileExportUserJobResponse, err error) {
    return c.CreateFileExportUserJobWithContext(context.Background(), request)
}

// CreateFileExportUserJob
// This API is used to create a file for user export task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATAFLOWTOOMANYREQUESTS = "FailedOperation.DataFlowTooManyRequests"
//  FAILEDOPERATION_IMPORTUSERISEMPTY = "FailedOperation.ImportUserIsEmpty"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateFileExportUserJobWithContext(ctx context.Context, request *CreateFileExportUserJobRequest) (response *CreateFileExportUserJobResponse, err error) {
    if request == nil {
        request = NewCreateFileExportUserJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ciam", APIVersion, "CreateFileExportUserJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFileExportUserJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFileExportUserJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserRequest() (request *CreateUserRequest) {
    request = &CreateUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ciam", APIVersion, "CreateUser")
    
    
    return
}

func NewCreateUserResponse() (response *CreateUserResponse) {
    response = &CreateUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateUser
// This API is used to create a user.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ATTRIBUTEFORMATERROR = "FailedOperation.AttributeFormatError"
//  FAILEDOPERATION_EMAILALREADYEXISTS = "FailedOperation.EmailAlreadyExists"
//  FAILEDOPERATION_EMAILISNULL = "FailedOperation.EmailIsNull"
//  FAILEDOPERATION_INDEXEDATTRIBUTETOOLONG = "FailedOperation.IndexedAttributeTooLong"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_PASSWORDISNULL = "FailedOperation.PasswordIsNull"
//  FAILEDOPERATION_PHONENUMBERALREADYEXISTS = "FailedOperation.PhoneNumberAlreadyExists"
//  FAILEDOPERATION_PHONENUMBERISNULL = "FailedOperation.PhoneNumberIsNull"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  FAILEDOPERATION_USERGROUPNOTFOUND = "FailedOperation.UserGroupNotFound"
//  FAILEDOPERATION_USERNAMEALREADYEXISTS = "FailedOperation.UserNameAlreadyExists"
//  FAILEDOPERATION_USERNAMEISNULL = "FailedOperation.UserNameIsNull"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateUser(request *CreateUserRequest) (response *CreateUserResponse, err error) {
    return c.CreateUserWithContext(context.Background(), request)
}

// CreateUser
// This API is used to create a user.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ATTRIBUTEFORMATERROR = "FailedOperation.AttributeFormatError"
//  FAILEDOPERATION_EMAILALREADYEXISTS = "FailedOperation.EmailAlreadyExists"
//  FAILEDOPERATION_EMAILISNULL = "FailedOperation.EmailIsNull"
//  FAILEDOPERATION_INDEXEDATTRIBUTETOOLONG = "FailedOperation.IndexedAttributeTooLong"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_PASSWORDISNULL = "FailedOperation.PasswordIsNull"
//  FAILEDOPERATION_PHONENUMBERALREADYEXISTS = "FailedOperation.PhoneNumberAlreadyExists"
//  FAILEDOPERATION_PHONENUMBERISNULL = "FailedOperation.PhoneNumberIsNull"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  FAILEDOPERATION_USERGROUPNOTFOUND = "FailedOperation.UserGroupNotFound"
//  FAILEDOPERATION_USERNAMEALREADYEXISTS = "FailedOperation.UserNameAlreadyExists"
//  FAILEDOPERATION_USERNAMEISNULL = "FailedOperation.UserNameIsNull"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateUserWithContext(ctx context.Context, request *CreateUserRequest) (response *CreateUserResponse, err error) {
    if request == nil {
        request = NewCreateUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ciam", APIVersion, "CreateUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUserResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUsersRequest() (request *DeleteUsersRequest) {
    request = &DeleteUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ciam", APIVersion, "DeleteUsers")
    
    
    return
}

func NewDeleteUsersResponse() (response *DeleteUsersResponse) {
    response = &DeleteUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteUsers
// This API is used to delete users in batches.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteUsers(request *DeleteUsersRequest) (response *DeleteUsersResponse, err error) {
    return c.DeleteUsersWithContext(context.Background(), request)
}

// DeleteUsers
// This API is used to delete users in batches.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteUsersWithContext(ctx context.Context, request *DeleteUsersRequest) (response *DeleteUsersResponse, err error) {
    if request == nil {
        request = NewDeleteUsersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ciam", APIVersion, "DeleteUsers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUsers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUsersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserRequest() (request *DescribeUserRequest) {
    request = &DescribeUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ciam", APIVersion, "DescribeUser")
    
    
    return
}

func NewDescribeUserResponse() (response *DescribeUserResponse) {
    response = &DescribeUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUser
// This API is used to query the user information with multiple conditions.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUERYUSERSPARAMETERMUSTINWHITELIST = "FailedOperation.QueryUsersParameterMustInWhitelist"
//  FAILEDOPERATION_QUERYUSERSPARAMETERREPEAT = "FailedOperation.QueryUsersParameterRepeat"
//  FAILEDOPERATION_QUERYUSERSSORTPARAMETERMUSTINWHITELIST = "FailedOperation.QueryUsersSortParameterMustInWhitelist"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUser(request *DescribeUserRequest) (response *DescribeUserResponse, err error) {
    return c.DescribeUserWithContext(context.Background(), request)
}

// DescribeUser
// This API is used to query the user information with multiple conditions.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUERYUSERSPARAMETERMUSTINWHITELIST = "FailedOperation.QueryUsersParameterMustInWhitelist"
//  FAILEDOPERATION_QUERYUSERSPARAMETERREPEAT = "FailedOperation.QueryUsersParameterRepeat"
//  FAILEDOPERATION_QUERYUSERSSORTPARAMETERMUSTINWHITELIST = "FailedOperation.QueryUsersSortParameterMustInWhitelist"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUserWithContext(ctx context.Context, request *DescribeUserRequest) (response *DescribeUserResponse, err error) {
    if request == nil {
        request = NewDescribeUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ciam", APIVersion, "DescribeUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserByIdRequest() (request *DescribeUserByIdRequest) {
    request = &DescribeUserByIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ciam", APIVersion, "DescribeUserById")
    
    
    return
}

func NewDescribeUserByIdResponse() (response *DescribeUserByIdResponse) {
    response = &DescribeUserByIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserById
// This API is used to query a user by ID.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUserById(request *DescribeUserByIdRequest) (response *DescribeUserByIdResponse, err error) {
    return c.DescribeUserByIdWithContext(context.Background(), request)
}

// DescribeUserById
// This API is used to query a user by ID.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUserByIdWithContext(ctx context.Context, request *DescribeUserByIdRequest) (response *DescribeUserByIdResponse, err error) {
    if request == nil {
        request = NewDescribeUserByIdRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ciam", APIVersion, "DescribeUserById")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserById require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserByIdResponse()
    err = c.Send(request, response)
    return
}

func NewLinkAccountRequest() (request *LinkAccountRequest) {
    request = &LinkAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ciam", APIVersion, "LinkAccount")
    
    
    return
}

func NewLinkAccountResponse() (response *LinkAccountResponse) {
    response = &LinkAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// LinkAccount
// This API is used to merge accounts.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCOUNTALREADYBELINKED = "FailedOperation.AccountAlreadyBeLinked"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTATUSENUM = "FailedOperation.InvalidUserStatusEnum"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_PRIMARYUSERNOTFOUND = "FailedOperation.PrimaryUserNotFound"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  FAILEDOPERATION_SECONDARYUSERNOTFOUND = "FailedOperation.SecondaryUserNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) LinkAccount(request *LinkAccountRequest) (response *LinkAccountResponse, err error) {
    return c.LinkAccountWithContext(context.Background(), request)
}

// LinkAccount
// This API is used to merge accounts.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCOUNTALREADYBELINKED = "FailedOperation.AccountAlreadyBeLinked"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTATUSENUM = "FailedOperation.InvalidUserStatusEnum"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_PRIMARYUSERNOTFOUND = "FailedOperation.PrimaryUserNotFound"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  FAILEDOPERATION_SECONDARYUSERNOTFOUND = "FailedOperation.SecondaryUserNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) LinkAccountWithContext(ctx context.Context, request *LinkAccountRequest) (response *LinkAccountResponse, err error) {
    if request == nil {
        request = NewLinkAccountRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ciam", APIVersion, "LinkAccount")
    
    if c.GetCredential() == nil {
        return nil, errors.New("LinkAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewLinkAccountResponse()
    err = c.Send(request, response)
    return
}

func NewListJobsRequest() (request *ListJobsRequest) {
    request = &ListJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ciam", APIVersion, "ListJobs")
    
    
    return
}

func NewListJobsResponse() (response *ListJobsResponse) {
    response = &ListJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListJobs
// This API is used to query the details of a task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATAFLOWTOOMANYREQUESTS = "FailedOperation.DataFlowTooManyRequests"
//  FAILEDOPERATION_IMPORTUSERISEMPTY = "FailedOperation.ImportUserIsEmpty"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ListJobs(request *ListJobsRequest) (response *ListJobsResponse, err error) {
    return c.ListJobsWithContext(context.Background(), request)
}

// ListJobs
// This API is used to query the details of a task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATAFLOWTOOMANYREQUESTS = "FailedOperation.DataFlowTooManyRequests"
//  FAILEDOPERATION_IMPORTUSERISEMPTY = "FailedOperation.ImportUserIsEmpty"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ListJobsWithContext(ctx context.Context, request *ListJobsRequest) (response *ListJobsResponse, err error) {
    if request == nil {
        request = NewListJobsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ciam", APIVersion, "ListJobs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewListJobsResponse()
    err = c.Send(request, response)
    return
}

func NewListLogMessageByConditionRequest() (request *ListLogMessageByConditionRequest) {
    request = &ListLogMessageByConditionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ciam", APIVersion, "ListLogMessageByCondition")
    
    
    return
}

func NewListLogMessageByConditionResponse() (response *ListLogMessageByConditionResponse) {
    response = &ListLogMessageByConditionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListLogMessageByCondition
// This API is used to query a specified log.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ListLogMessageByCondition(request *ListLogMessageByConditionRequest) (response *ListLogMessageByConditionResponse, err error) {
    return c.ListLogMessageByConditionWithContext(context.Background(), request)
}

// ListLogMessageByCondition
// This API is used to query a specified log.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ListLogMessageByConditionWithContext(ctx context.Context, request *ListLogMessageByConditionRequest) (response *ListLogMessageByConditionResponse, err error) {
    if request == nil {
        request = NewListLogMessageByConditionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ciam", APIVersion, "ListLogMessageByCondition")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListLogMessageByCondition require credential")
    }

    request.SetContext(ctx)
    
    response = NewListLogMessageByConditionResponse()
    err = c.Send(request, response)
    return
}

func NewListUserRequest() (request *ListUserRequest) {
    request = &ListUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ciam", APIVersion, "ListUser")
    
    
    return
}

func NewListUserResponse() (response *ListUserResponse) {
    response = &ListUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListUser
// This API is used to query the list of users.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ListUser(request *ListUserRequest) (response *ListUserResponse, err error) {
    return c.ListUserWithContext(context.Background(), request)
}

// ListUser
// This API is used to query the list of users.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ListUserWithContext(ctx context.Context, request *ListUserRequest) (response *ListUserResponse, err error) {
    if request == nil {
        request = NewListUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ciam", APIVersion, "ListUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewListUserResponse()
    err = c.Send(request, response)
    return
}

func NewListUserByPropertyRequest() (request *ListUserByPropertyRequest) {
    request = &ListUserByPropertyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ciam", APIVersion, "ListUserByProperty")
    
    
    return
}

func NewListUserByPropertyResponse() (response *ListUserByPropertyResponse) {
    response = &ListUserByPropertyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListUserByProperty
// This API is used to query the list of users by attribute.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ListUserByProperty(request *ListUserByPropertyRequest) (response *ListUserByPropertyResponse, err error) {
    return c.ListUserByPropertyWithContext(context.Background(), request)
}

// ListUserByProperty
// This API is used to query the list of users by attribute.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ListUserByPropertyWithContext(ctx context.Context, request *ListUserByPropertyRequest) (response *ListUserByPropertyResponse, err error) {
    if request == nil {
        request = NewListUserByPropertyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ciam", APIVersion, "ListUserByProperty")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListUserByProperty require credential")
    }

    request.SetContext(ctx)
    
    response = NewListUserByPropertyResponse()
    err = c.Send(request, response)
    return
}

func NewResetPasswordRequest() (request *ResetPasswordRequest) {
    request = &ResetPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ciam", APIVersion, "ResetPassword")
    
    
    return
}

func NewResetPasswordResponse() (response *ResetPasswordResponse) {
    response = &ResetPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetPassword
// This API is used to reset the password for a user.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  FAILEDOPERATION_USERISFREEZE = "FailedOperation.UserIsFreeze"
//  FAILEDOPERATION_USERNOTFOUND = "FailedOperation.UserNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ResetPassword(request *ResetPasswordRequest) (response *ResetPasswordResponse, err error) {
    return c.ResetPasswordWithContext(context.Background(), request)
}

// ResetPassword
// This API is used to reset the password for a user.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  FAILEDOPERATION_USERISFREEZE = "FailedOperation.UserIsFreeze"
//  FAILEDOPERATION_USERNOTFOUND = "FailedOperation.UserNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ResetPasswordWithContext(ctx context.Context, request *ResetPasswordRequest) (response *ResetPasswordResponse, err error) {
    if request == nil {
        request = NewResetPasswordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ciam", APIVersion, "ResetPassword")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetPassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewSetPasswordRequest() (request *SetPasswordRequest) {
    request = &SetPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ciam", APIVersion, "SetPassword")
    
    
    return
}

func NewSetPasswordResponse() (response *SetPasswordResponse) {
    response = &SetPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetPassword
// This API is used to set the password for a user.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_PASSWORDISNULL = "FailedOperation.PasswordIsNull"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  FAILEDOPERATION_USERISFREEZE = "FailedOperation.UserIsFreeze"
//  FAILEDOPERATION_USERNOTFOUND = "FailedOperation.UserNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SetPassword(request *SetPasswordRequest) (response *SetPasswordResponse, err error) {
    return c.SetPasswordWithContext(context.Background(), request)
}

// SetPassword
// This API is used to set the password for a user.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_PASSWORDISNULL = "FailedOperation.PasswordIsNull"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  FAILEDOPERATION_USERISFREEZE = "FailedOperation.UserIsFreeze"
//  FAILEDOPERATION_USERNOTFOUND = "FailedOperation.UserNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SetPasswordWithContext(ctx context.Context, request *SetPasswordRequest) (response *SetPasswordResponse, err error) {
    if request == nil {
        request = NewSetPasswordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ciam", APIVersion, "SetPassword")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetPassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateUserRequest() (request *UpdateUserRequest) {
    request = &UpdateUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ciam", APIVersion, "UpdateUser")
    
    
    return
}

func NewUpdateUserResponse() (response *UpdateUserResponse) {
    response = &UpdateUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateUser
// This API is used to update a user.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ATTRIBUTEFORMATERROR = "FailedOperation.AttributeFormatError"
//  FAILEDOPERATION_EMAILALREADYEXISTS = "FailedOperation.EmailAlreadyExists"
//  FAILEDOPERATION_INDEXEDATTRIBUTETOOLONG = "FailedOperation.IndexedAttributeTooLong"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_PHONENUMBERALREADYEXISTS = "FailedOperation.PhoneNumberAlreadyExists"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  FAILEDOPERATION_USERGROUPNOTFOUND = "FailedOperation.UserGroupNotFound"
//  FAILEDOPERATION_USERNAMEALREADYEXISTS = "FailedOperation.UserNameAlreadyExists"
//  FAILEDOPERATION_USERNOTFOUND = "FailedOperation.UserNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateUser(request *UpdateUserRequest) (response *UpdateUserResponse, err error) {
    return c.UpdateUserWithContext(context.Background(), request)
}

// UpdateUser
// This API is used to update a user.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ATTRIBUTEFORMATERROR = "FailedOperation.AttributeFormatError"
//  FAILEDOPERATION_EMAILALREADYEXISTS = "FailedOperation.EmailAlreadyExists"
//  FAILEDOPERATION_INDEXEDATTRIBUTETOOLONG = "FailedOperation.IndexedAttributeTooLong"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_PHONENUMBERALREADYEXISTS = "FailedOperation.PhoneNumberAlreadyExists"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  FAILEDOPERATION_USERGROUPNOTFOUND = "FailedOperation.UserGroupNotFound"
//  FAILEDOPERATION_USERNAMEALREADYEXISTS = "FailedOperation.UserNameAlreadyExists"
//  FAILEDOPERATION_USERNOTFOUND = "FailedOperation.UserNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateUserWithContext(ctx context.Context, request *UpdateUserRequest) (response *UpdateUserResponse, err error) {
    if request == nil {
        request = NewUpdateUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ciam", APIVersion, "UpdateUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateUserResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateUserStatusRequest() (request *UpdateUserStatusRequest) {
    request = &UpdateUserStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ciam", APIVersion, "UpdateUserStatus")
    
    
    return
}

func NewUpdateUserStatusResponse() (response *UpdateUserStatusResponse) {
    response = &UpdateUserStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateUserStatus
// This API is used to update the status of a user.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTATUSENUM = "FailedOperation.InvalidUserStatusEnum"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  FAILEDOPERATION_USERNOTFOUND = "FailedOperation.UserNotFound"
//  FAILEDOPERATION_USERSTATUSREQUIRED = "FailedOperation.UserStatusRequired"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateUserStatus(request *UpdateUserStatusRequest) (response *UpdateUserStatusResponse, err error) {
    return c.UpdateUserStatusWithContext(context.Background(), request)
}

// UpdateUserStatus
// This API is used to update the status of a user.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_INVALIDUSERSTATUSENUM = "FailedOperation.InvalidUserStatusEnum"
//  FAILEDOPERATION_INVALIDUSERSTORE = "FailedOperation.InvalidUserStore"
//  FAILEDOPERATION_QUOTALIMITEXCEEDED = "FailedOperation.QuotaLimitExceeded"
//  FAILEDOPERATION_USERNOTFOUND = "FailedOperation.UserNotFound"
//  FAILEDOPERATION_USERSTATUSREQUIRED = "FailedOperation.UserStatusRequired"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_FREQUENTREQUEST = "RequestLimitExceeded.FrequentRequest"
//  REQUESTLIMITEXCEEDED_REPEATREQUEST = "RequestLimitExceeded.RepeatRequest"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateUserStatusWithContext(ctx context.Context, request *UpdateUserStatusRequest) (response *UpdateUserStatusResponse, err error) {
    if request == nil {
        request = NewUpdateUserStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ciam", APIVersion, "UpdateUserStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateUserStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateUserStatusResponse()
    err = c.Send(request, response)
    return
}
