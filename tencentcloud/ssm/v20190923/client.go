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

package v20190923

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-09-23"

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


func NewCreateProductSecretRequest() (request *CreateProductSecretRequest) {
    request = &CreateProductSecretRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "CreateProductSecret")
    
    
    return
}

func NewCreateProductSecretResponse() (response *CreateProductSecretResponse) {
    response = &CreateProductSecretResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateProductSecret
// This API is used to create a Tencent Cloud service credential.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSKMSERROR = "FailedOperation.AccessKmsError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TAGKEYSDUPLICATED = "InvalidParameterValue.TagKeysDuplicated"
//  INVALIDPARAMETERVALUE_TAGSNOTEXISTED = "InvalidParameterValue.TagsNotExisted"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_AUTOROTATEDRESOURCE = "OperationDenied.AutoRotatedResource"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_SECRETEXISTS = "ResourceInUse.SecretExists"
//  RESOURCEINUSE_VERSIONIDEXISTS = "ResourceInUse.VersionIdExists"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  RESOURCEUNAVAILABLE_RESOURCEDISABLED = "ResourceUnavailable.ResourceDisabled"
//  RESOURCEUNAVAILABLE_RESOURCEPENDINGDELETED = "ResourceUnavailable.ResourcePendingDeleted"
//  RESOURCEUNAVAILABLE_RESOURCEUNINITIALIZED = "ResourceUnavailable.ResourceUninitialized"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ACCESSKMSERROR = "UnauthorizedOperation.AccessKmsError"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateProductSecret(request *CreateProductSecretRequest) (response *CreateProductSecretResponse, err error) {
    return c.CreateProductSecretWithContext(context.Background(), request)
}

// CreateProductSecret
// This API is used to create a Tencent Cloud service credential.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSKMSERROR = "FailedOperation.AccessKmsError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TAGKEYSDUPLICATED = "InvalidParameterValue.TagKeysDuplicated"
//  INVALIDPARAMETERVALUE_TAGSNOTEXISTED = "InvalidParameterValue.TagsNotExisted"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_AUTOROTATEDRESOURCE = "OperationDenied.AutoRotatedResource"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_SECRETEXISTS = "ResourceInUse.SecretExists"
//  RESOURCEINUSE_VERSIONIDEXISTS = "ResourceInUse.VersionIdExists"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  RESOURCEUNAVAILABLE_RESOURCEDISABLED = "ResourceUnavailable.ResourceDisabled"
//  RESOURCEUNAVAILABLE_RESOURCEPENDINGDELETED = "ResourceUnavailable.ResourcePendingDeleted"
//  RESOURCEUNAVAILABLE_RESOURCEUNINITIALIZED = "ResourceUnavailable.ResourceUninitialized"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ACCESSKMSERROR = "UnauthorizedOperation.AccessKmsError"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateProductSecretWithContext(ctx context.Context, request *CreateProductSecretRequest) (response *CreateProductSecretResponse, err error) {
    if request == nil {
        request = NewCreateProductSecretRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateProductSecret require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateProductSecretResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSSHKeyPairSecretRequest() (request *CreateSSHKeyPairSecretRequest) {
    request = &CreateSSHKeyPairSecretRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "CreateSSHKeyPairSecret")
    
    
    return
}

func NewCreateSSHKeyPairSecretResponse() (response *CreateSSHKeyPairSecretResponse) {
    response = &CreateSSHKeyPairSecretResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSSHKeyPairSecret
// This API is used to create a secret that hosts SSH keys.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSKMSERROR = "FailedOperation.AccessKmsError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TAGKEYSDUPLICATED = "InvalidParameterValue.TagKeysDuplicated"
//  INVALIDPARAMETERVALUE_TAGSNOTEXISTED = "InvalidParameterValue.TagsNotExisted"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE_SECRETEXISTS = "ResourceInUse.SecretExists"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ACCESSKMSERROR = "UnauthorizedOperation.AccessKmsError"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSSHKeyPairSecret(request *CreateSSHKeyPairSecretRequest) (response *CreateSSHKeyPairSecretResponse, err error) {
    return c.CreateSSHKeyPairSecretWithContext(context.Background(), request)
}

// CreateSSHKeyPairSecret
// This API is used to create a secret that hosts SSH keys.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSKMSERROR = "FailedOperation.AccessKmsError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TAGKEYSDUPLICATED = "InvalidParameterValue.TagKeysDuplicated"
//  INVALIDPARAMETERVALUE_TAGSNOTEXISTED = "InvalidParameterValue.TagsNotExisted"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE_SECRETEXISTS = "ResourceInUse.SecretExists"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ACCESSKMSERROR = "UnauthorizedOperation.AccessKmsError"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSSHKeyPairSecretWithContext(ctx context.Context, request *CreateSSHKeyPairSecretRequest) (response *CreateSSHKeyPairSecretResponse, err error) {
    if request == nil {
        request = NewCreateSSHKeyPairSecretRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSSHKeyPairSecret require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSSHKeyPairSecretResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSecretRequest() (request *CreateSecretRequest) {
    request = &CreateSecretRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "CreateSecret")
    
    
    return
}

func NewCreateSecretResponse() (response *CreateSecretResponse) {
    response = &CreateSecretResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSecret
// This API is used to create a KMS-encrypted Secret. You can create and store up to 1,000 Secrets in each region.
//
// error code that may be returned:
//  FAILEDOPERATION_ACCESSKMSERROR = "FailedOperation.AccessKmsError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TAGKEYSDUPLICATED = "InvalidParameterValue.TagKeysDuplicated"
//  INVALIDPARAMETERVALUE_TAGSNOTEXISTED = "InvalidParameterValue.TagsNotExisted"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE_SECRETEXISTS = "ResourceInUse.SecretExists"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateSecret(request *CreateSecretRequest) (response *CreateSecretResponse, err error) {
    return c.CreateSecretWithContext(context.Background(), request)
}

// CreateSecret
// This API is used to create a KMS-encrypted Secret. You can create and store up to 1,000 Secrets in each region.
//
// error code that may be returned:
//  FAILEDOPERATION_ACCESSKMSERROR = "FailedOperation.AccessKmsError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TAGKEYSDUPLICATED = "InvalidParameterValue.TagKeysDuplicated"
//  INVALIDPARAMETERVALUE_TAGSNOTEXISTED = "InvalidParameterValue.TagsNotExisted"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE_SECRETEXISTS = "ResourceInUse.SecretExists"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateSecretWithContext(ctx context.Context, request *CreateSecretRequest) (response *CreateSecretResponse, err error) {
    if request == nil {
        request = NewCreateSecretRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSecret require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSecretResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSecretRequest() (request *DeleteSecretRequest) {
    request = &DeleteSecretRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "DeleteSecret")
    
    
    return
}

func NewDeleteSecretResponse() (response *DeleteSecretResponse) {
    response = &DeleteSecretResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteSecret
// This API is used to delete a Secret. You can set whether to delete the Secret immediately or on schedule using the `RecoveryWindowInDays` parameter. For a Secret to be deleted on schedule, its status will be `PendingDelete` before the scheduled deletion time. You can use `RestoreSecret` to restore a deleted Secret during this time. A deleted Secret will not be restorable after the scheduled deletion time. A Secret can only be deleted after being disabled using `DisableSecret`.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  RESOURCEUNAVAILABLE_RESOURCEUNINITIALIZED = "ResourceUnavailable.ResourceUninitialized"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteSecret(request *DeleteSecretRequest) (response *DeleteSecretResponse, err error) {
    return c.DeleteSecretWithContext(context.Background(), request)
}

// DeleteSecret
// This API is used to delete a Secret. You can set whether to delete the Secret immediately or on schedule using the `RecoveryWindowInDays` parameter. For a Secret to be deleted on schedule, its status will be `PendingDelete` before the scheduled deletion time. You can use `RestoreSecret` to restore a deleted Secret during this time. A deleted Secret will not be restorable after the scheduled deletion time. A Secret can only be deleted after being disabled using `DisableSecret`.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  RESOURCEUNAVAILABLE_RESOURCEUNINITIALIZED = "ResourceUnavailable.ResourceUninitialized"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteSecretWithContext(ctx context.Context, request *DeleteSecretRequest) (response *DeleteSecretResponse, err error) {
    if request == nil {
        request = NewDeleteSecretRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSecret require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSecretResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSecretVersionRequest() (request *DeleteSecretVersionRequest) {
    request = &DeleteSecretVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "DeleteSecretVersion")
    
    
    return
}

func NewDeleteSecretVersionResponse() (response *DeleteSecretVersionResponse) {
    response = &DeleteSecretVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteSecretVersion
// This API is used to directly delete a single credential version under the specified credential. The deletion takes effect immediately, and the credential version in all status can be deleted.
//
// This API is only applicable to user-defined credentials but not Tencent Cloud service credentials.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  RESOURCEUNAVAILABLE_RESOURCEUNINITIALIZED = "ResourceUnavailable.ResourceUninitialized"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteSecretVersion(request *DeleteSecretVersionRequest) (response *DeleteSecretVersionResponse, err error) {
    return c.DeleteSecretVersionWithContext(context.Background(), request)
}

// DeleteSecretVersion
// This API is used to directly delete a single credential version under the specified credential. The deletion takes effect immediately, and the credential version in all status can be deleted.
//
// This API is only applicable to user-defined credentials but not Tencent Cloud service credentials.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  RESOURCEUNAVAILABLE_RESOURCEUNINITIALIZED = "ResourceUnavailable.ResourceUninitialized"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteSecretVersionWithContext(ctx context.Context, request *DeleteSecretVersionRequest) (response *DeleteSecretVersionResponse, err error) {
    if request == nil {
        request = NewDeleteSecretVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSecretVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSecretVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAsyncRequestInfoRequest() (request *DescribeAsyncRequestInfoRequest) {
    request = &DescribeAsyncRequestInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "DescribeAsyncRequestInfo")
    
    
    return
}

func NewDescribeAsyncRequestInfoResponse() (response *DescribeAsyncRequestInfoResponse) {
    response = &DescribeAsyncRequestInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAsyncRequestInfo
// This API is used to query the execution result of an async task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSKMSERROR = "FailedOperation.AccessKmsError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TAGKEYSDUPLICATED = "InvalidParameterValue.TagKeysDuplicated"
//  INVALIDPARAMETERVALUE_TAGSNOTEXISTED = "InvalidParameterValue.TagsNotExisted"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_AUTOROTATEDRESOURCE = "OperationDenied.AutoRotatedResource"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_SECRETEXISTS = "ResourceInUse.SecretExists"
//  RESOURCEINUSE_VERSIONIDEXISTS = "ResourceInUse.VersionIdExists"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  RESOURCEUNAVAILABLE_RESOURCEDISABLED = "ResourceUnavailable.ResourceDisabled"
//  RESOURCEUNAVAILABLE_RESOURCEPENDINGDELETED = "ResourceUnavailable.ResourcePendingDeleted"
//  RESOURCEUNAVAILABLE_RESOURCEUNINITIALIZED = "ResourceUnavailable.ResourceUninitialized"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ACCESSKMSERROR = "UnauthorizedOperation.AccessKmsError"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAsyncRequestInfo(request *DescribeAsyncRequestInfoRequest) (response *DescribeAsyncRequestInfoResponse, err error) {
    return c.DescribeAsyncRequestInfoWithContext(context.Background(), request)
}

// DescribeAsyncRequestInfo
// This API is used to query the execution result of an async task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSKMSERROR = "FailedOperation.AccessKmsError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TAGKEYSDUPLICATED = "InvalidParameterValue.TagKeysDuplicated"
//  INVALIDPARAMETERVALUE_TAGSNOTEXISTED = "InvalidParameterValue.TagsNotExisted"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_AUTOROTATEDRESOURCE = "OperationDenied.AutoRotatedResource"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_SECRETEXISTS = "ResourceInUse.SecretExists"
//  RESOURCEINUSE_VERSIONIDEXISTS = "ResourceInUse.VersionIdExists"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  RESOURCEUNAVAILABLE_RESOURCEDISABLED = "ResourceUnavailable.ResourceDisabled"
//  RESOURCEUNAVAILABLE_RESOURCEPENDINGDELETED = "ResourceUnavailable.ResourcePendingDeleted"
//  RESOURCEUNAVAILABLE_RESOURCEUNINITIALIZED = "ResourceUnavailable.ResourceUninitialized"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ACCESSKMSERROR = "UnauthorizedOperation.AccessKmsError"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAsyncRequestInfoWithContext(ctx context.Context, request *DescribeAsyncRequestInfoRequest) (response *DescribeAsyncRequestInfoResponse, err error) {
    if request == nil {
        request = NewDescribeAsyncRequestInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAsyncRequestInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAsyncRequestInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRotationDetailRequest() (request *DescribeRotationDetailRequest) {
    request = &DescribeRotationDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "DescribeRotationDetail")
    
    
    return
}

func NewDescribeRotationDetailResponse() (response *DescribeRotationDetailResponse) {
    response = &DescribeRotationDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRotationDetail
// This API is used to query the details of a credential rotation policy.
//
// This API is only applicable to Tencent Cloud service credentials.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSKMSERROR = "FailedOperation.AccessKmsError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TAGKEYSDUPLICATED = "InvalidParameterValue.TagKeysDuplicated"
//  INVALIDPARAMETERVALUE_TAGSNOTEXISTED = "InvalidParameterValue.TagsNotExisted"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_AUTOROTATEDRESOURCE = "OperationDenied.AutoRotatedResource"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_SECRETEXISTS = "ResourceInUse.SecretExists"
//  RESOURCEINUSE_VERSIONIDEXISTS = "ResourceInUse.VersionIdExists"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  RESOURCEUNAVAILABLE_RESOURCEDISABLED = "ResourceUnavailable.ResourceDisabled"
//  RESOURCEUNAVAILABLE_RESOURCEPENDINGDELETED = "ResourceUnavailable.ResourcePendingDeleted"
//  RESOURCEUNAVAILABLE_RESOURCEUNINITIALIZED = "ResourceUnavailable.ResourceUninitialized"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ACCESSKMSERROR = "UnauthorizedOperation.AccessKmsError"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRotationDetail(request *DescribeRotationDetailRequest) (response *DescribeRotationDetailResponse, err error) {
    return c.DescribeRotationDetailWithContext(context.Background(), request)
}

// DescribeRotationDetail
// This API is used to query the details of a credential rotation policy.
//
// This API is only applicable to Tencent Cloud service credentials.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSKMSERROR = "FailedOperation.AccessKmsError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TAGKEYSDUPLICATED = "InvalidParameterValue.TagKeysDuplicated"
//  INVALIDPARAMETERVALUE_TAGSNOTEXISTED = "InvalidParameterValue.TagsNotExisted"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_AUTOROTATEDRESOURCE = "OperationDenied.AutoRotatedResource"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_SECRETEXISTS = "ResourceInUse.SecretExists"
//  RESOURCEINUSE_VERSIONIDEXISTS = "ResourceInUse.VersionIdExists"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  RESOURCEUNAVAILABLE_RESOURCEDISABLED = "ResourceUnavailable.ResourceDisabled"
//  RESOURCEUNAVAILABLE_RESOURCEPENDINGDELETED = "ResourceUnavailable.ResourcePendingDeleted"
//  RESOURCEUNAVAILABLE_RESOURCEUNINITIALIZED = "ResourceUnavailable.ResourceUninitialized"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ACCESSKMSERROR = "UnauthorizedOperation.AccessKmsError"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRotationDetailWithContext(ctx context.Context, request *DescribeRotationDetailRequest) (response *DescribeRotationDetailResponse, err error) {
    if request == nil {
        request = NewDescribeRotationDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRotationDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRotationDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRotationHistoryRequest() (request *DescribeRotationHistoryRequest) {
    request = &DescribeRotationHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "DescribeRotationHistory")
    
    
    return
}

func NewDescribeRotationHistoryResponse() (response *DescribeRotationHistoryResponse) {
    response = &DescribeRotationHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRotationHistory
// This API is used to query the historical versions of a rotated credential.
//
// This API is only applicable to Tencent Cloud service credentials.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSKMSERROR = "FailedOperation.AccessKmsError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TAGKEYSDUPLICATED = "InvalidParameterValue.TagKeysDuplicated"
//  INVALIDPARAMETERVALUE_TAGSNOTEXISTED = "InvalidParameterValue.TagsNotExisted"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_AUTOROTATEDRESOURCE = "OperationDenied.AutoRotatedResource"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_SECRETEXISTS = "ResourceInUse.SecretExists"
//  RESOURCEINUSE_VERSIONIDEXISTS = "ResourceInUse.VersionIdExists"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  RESOURCEUNAVAILABLE_RESOURCEDISABLED = "ResourceUnavailable.ResourceDisabled"
//  RESOURCEUNAVAILABLE_RESOURCEPENDINGDELETED = "ResourceUnavailable.ResourcePendingDeleted"
//  RESOURCEUNAVAILABLE_RESOURCEUNINITIALIZED = "ResourceUnavailable.ResourceUninitialized"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ACCESSKMSERROR = "UnauthorizedOperation.AccessKmsError"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRotationHistory(request *DescribeRotationHistoryRequest) (response *DescribeRotationHistoryResponse, err error) {
    return c.DescribeRotationHistoryWithContext(context.Background(), request)
}

// DescribeRotationHistory
// This API is used to query the historical versions of a rotated credential.
//
// This API is only applicable to Tencent Cloud service credentials.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSKMSERROR = "FailedOperation.AccessKmsError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TAGKEYSDUPLICATED = "InvalidParameterValue.TagKeysDuplicated"
//  INVALIDPARAMETERVALUE_TAGSNOTEXISTED = "InvalidParameterValue.TagsNotExisted"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_AUTOROTATEDRESOURCE = "OperationDenied.AutoRotatedResource"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_SECRETEXISTS = "ResourceInUse.SecretExists"
//  RESOURCEINUSE_VERSIONIDEXISTS = "ResourceInUse.VersionIdExists"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  RESOURCEUNAVAILABLE_RESOURCEDISABLED = "ResourceUnavailable.ResourceDisabled"
//  RESOURCEUNAVAILABLE_RESOURCEPENDINGDELETED = "ResourceUnavailable.ResourcePendingDeleted"
//  RESOURCEUNAVAILABLE_RESOURCEUNINITIALIZED = "ResourceUnavailable.ResourceUninitialized"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ACCESSKMSERROR = "UnauthorizedOperation.AccessKmsError"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRotationHistoryWithContext(ctx context.Context, request *DescribeRotationHistoryRequest) (response *DescribeRotationHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeRotationHistoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRotationHistory require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRotationHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecretRequest() (request *DescribeSecretRequest) {
    request = &DescribeSecretRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "DescribeSecret")
    
    
    return
}

func NewDescribeSecretResponse() (response *DescribeSecretResponse) {
    response = &DescribeSecretResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSecret
// This API is used to obtain the detailed attribute information of a Secret.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_ROLENOTEXIST = "OperationDenied.RoleNotExist"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeSecret(request *DescribeSecretRequest) (response *DescribeSecretResponse, err error) {
    return c.DescribeSecretWithContext(context.Background(), request)
}

// DescribeSecret
// This API is used to obtain the detailed attribute information of a Secret.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_ROLENOTEXIST = "OperationDenied.RoleNotExist"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeSecretWithContext(ctx context.Context, request *DescribeSecretRequest) (response *DescribeSecretResponse, err error) {
    if request == nil {
        request = NewDescribeSecretRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecret require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecretResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSupportedProductsRequest() (request *DescribeSupportedProductsRequest) {
    request = &DescribeSupportedProductsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "DescribeSupportedProducts")
    
    
    return
}

func NewDescribeSupportedProductsResponse() (response *DescribeSupportedProductsResponse) {
    response = &DescribeSupportedProductsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSupportedProducts
// This API is used to query the list of supported Tencent Cloud services.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSKMSERROR = "FailedOperation.AccessKmsError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TAGKEYSDUPLICATED = "InvalidParameterValue.TagKeysDuplicated"
//  INVALIDPARAMETERVALUE_TAGSNOTEXISTED = "InvalidParameterValue.TagsNotExisted"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_AUTOROTATEDRESOURCE = "OperationDenied.AutoRotatedResource"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_SECRETEXISTS = "ResourceInUse.SecretExists"
//  RESOURCEINUSE_VERSIONIDEXISTS = "ResourceInUse.VersionIdExists"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  RESOURCEUNAVAILABLE_RESOURCEDISABLED = "ResourceUnavailable.ResourceDisabled"
//  RESOURCEUNAVAILABLE_RESOURCEPENDINGDELETED = "ResourceUnavailable.ResourcePendingDeleted"
//  RESOURCEUNAVAILABLE_RESOURCEUNINITIALIZED = "ResourceUnavailable.ResourceUninitialized"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ACCESSKMSERROR = "UnauthorizedOperation.AccessKmsError"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSupportedProducts(request *DescribeSupportedProductsRequest) (response *DescribeSupportedProductsResponse, err error) {
    return c.DescribeSupportedProductsWithContext(context.Background(), request)
}

// DescribeSupportedProducts
// This API is used to query the list of supported Tencent Cloud services.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSKMSERROR = "FailedOperation.AccessKmsError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TAGKEYSDUPLICATED = "InvalidParameterValue.TagKeysDuplicated"
//  INVALIDPARAMETERVALUE_TAGSNOTEXISTED = "InvalidParameterValue.TagsNotExisted"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_AUTOROTATEDRESOURCE = "OperationDenied.AutoRotatedResource"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_SECRETEXISTS = "ResourceInUse.SecretExists"
//  RESOURCEINUSE_VERSIONIDEXISTS = "ResourceInUse.VersionIdExists"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  RESOURCEUNAVAILABLE_RESOURCEDISABLED = "ResourceUnavailable.ResourceDisabled"
//  RESOURCEUNAVAILABLE_RESOURCEPENDINGDELETED = "ResourceUnavailable.ResourcePendingDeleted"
//  RESOURCEUNAVAILABLE_RESOURCEUNINITIALIZED = "ResourceUnavailable.ResourceUninitialized"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ACCESSKMSERROR = "UnauthorizedOperation.AccessKmsError"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSupportedProductsWithContext(ctx context.Context, request *DescribeSupportedProductsRequest) (response *DescribeSupportedProductsResponse, err error) {
    if request == nil {
        request = NewDescribeSupportedProductsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSupportedProducts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSupportedProductsResponse()
    err = c.Send(request, response)
    return
}

func NewDisableSecretRequest() (request *DisableSecretRequest) {
    request = &DisableSecretRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "DisableSecret")
    
    
    return
}

func NewDisableSecretResponse() (response *DisableSecretResponse) {
    response = &DisableSecretResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisableSecret
// This API is used to disable a Secret and will change its status to `Disabled`. The plaintext of a disabled Secret cannot be obtained through APIs.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  RESOURCEUNAVAILABLE_RESOURCEUNINITIALIZED = "ResourceUnavailable.ResourceUninitialized"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DisableSecret(request *DisableSecretRequest) (response *DisableSecretResponse, err error) {
    return c.DisableSecretWithContext(context.Background(), request)
}

// DisableSecret
// This API is used to disable a Secret and will change its status to `Disabled`. The plaintext of a disabled Secret cannot be obtained through APIs.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  RESOURCEUNAVAILABLE_RESOURCEUNINITIALIZED = "ResourceUnavailable.ResourceUninitialized"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DisableSecretWithContext(ctx context.Context, request *DisableSecretRequest) (response *DisableSecretResponse, err error) {
    if request == nil {
        request = NewDisableSecretRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableSecret require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableSecretResponse()
    err = c.Send(request, response)
    return
}

func NewEnableSecretRequest() (request *EnableSecretRequest) {
    request = &EnableSecretRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "EnableSecret")
    
    
    return
}

func NewEnableSecretResponse() (response *EnableSecretResponse) {
    response = &EnableSecretResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EnableSecret
// This API is used to enable a Secret and will change its status to `Enabled`. You can call the `GetSecretValue` API to obtain the plaintext of this Secret. Secrets in `PendingDelete` status can only be enabled after being restored by using `RestoreSecret`.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  RESOURCEUNAVAILABLE_RESOURCEUNINITIALIZED = "ResourceUnavailable.ResourceUninitialized"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) EnableSecret(request *EnableSecretRequest) (response *EnableSecretResponse, err error) {
    return c.EnableSecretWithContext(context.Background(), request)
}

// EnableSecret
// This API is used to enable a Secret and will change its status to `Enabled`. You can call the `GetSecretValue` API to obtain the plaintext of this Secret. Secrets in `PendingDelete` status can only be enabled after being restored by using `RestoreSecret`.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  RESOURCEUNAVAILABLE_RESOURCEUNINITIALIZED = "ResourceUnavailable.ResourceUninitialized"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) EnableSecretWithContext(ctx context.Context, request *EnableSecretRequest) (response *EnableSecretResponse, err error) {
    if request == nil {
        request = NewEnableSecretRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableSecret require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableSecretResponse()
    err = c.Send(request, response)
    return
}

func NewGetRegionsRequest() (request *GetRegionsRequest) {
    request = &GetRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "GetRegions")
    
    
    return
}

func NewGetRegionsResponse() (response *GetRegionsResponse) {
    response = &GetRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetRegions
// This API is used to obtain the list of regions displayed on Console.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) GetRegions(request *GetRegionsRequest) (response *GetRegionsResponse, err error) {
    return c.GetRegionsWithContext(context.Background(), request)
}

// GetRegions
// This API is used to obtain the list of regions displayed on Console.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) GetRegionsWithContext(ctx context.Context, request *GetRegionsRequest) (response *GetRegionsResponse, err error) {
    if request == nil {
        request = NewGetRegionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetRegions require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewGetSSHKeyPairValueRequest() (request *GetSSHKeyPairValueRequest) {
    request = &GetSSHKeyPairValueRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "GetSSHKeyPairValue")
    
    
    return
}

func NewGetSSHKeyPairValueResponse() (response *GetSSHKeyPairValueResponse) {
    response = &GetSSHKeyPairValueResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetSSHKeyPairValue
// This API is used to obtain the plaintext value of the SSH key secret.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSKMSERROR = "FailedOperation.AccessKmsError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TAGKEYSDUPLICATED = "InvalidParameterValue.TagKeysDuplicated"
//  INVALIDPARAMETERVALUE_TAGSNOTEXISTED = "InvalidParameterValue.TagsNotExisted"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ACCESSKMSERROR = "UnauthorizedOperation.AccessKmsError"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetSSHKeyPairValue(request *GetSSHKeyPairValueRequest) (response *GetSSHKeyPairValueResponse, err error) {
    return c.GetSSHKeyPairValueWithContext(context.Background(), request)
}

// GetSSHKeyPairValue
// This API is used to obtain the plaintext value of the SSH key secret.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSKMSERROR = "FailedOperation.AccessKmsError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TAGKEYSDUPLICATED = "InvalidParameterValue.TagKeysDuplicated"
//  INVALIDPARAMETERVALUE_TAGSNOTEXISTED = "InvalidParameterValue.TagsNotExisted"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ACCESSKMSERROR = "UnauthorizedOperation.AccessKmsError"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetSSHKeyPairValueWithContext(ctx context.Context, request *GetSSHKeyPairValueRequest) (response *GetSSHKeyPairValueResponse, err error) {
    if request == nil {
        request = NewGetSSHKeyPairValueRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetSSHKeyPairValue require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetSSHKeyPairValueResponse()
    err = c.Send(request, response)
    return
}

func NewGetSecretValueRequest() (request *GetSecretValueRequest) {
    request = &GetSecretValueRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "GetSecretValue")
    
    
    return
}

func NewGetSecretValueResponse() (response *GetSecretValueResponse) {
    response = &GetSecretValueResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetSecretValue
// For user-defined credentials, this API is used to get the plaintext information of a credential by specifying the credential name and version.
//
// For Tencent Cloud service credentials such as MySQL credentials, this API is used to get the plaintext information of a previously rotated credential by specifying the credential name and historical version number. If you want to get the plaintext of the credential version currently in use, you need to specify the version number as `SSM_Current`.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSKMSERROR = "FailedOperation.AccessKmsError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_ROLENOTEXIST = "OperationDenied.RoleNotExist"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SECRETNOTEXIST = "ResourceNotFound.SecretNotExist"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  RESOURCEUNAVAILABLE_RESOURCEDISABLED = "ResourceUnavailable.ResourceDisabled"
//  RESOURCEUNAVAILABLE_RESOURCEPENDINGDELETED = "ResourceUnavailable.ResourcePendingDeleted"
//  RESOURCEUNAVAILABLE_RESOURCEUNINITIALIZED = "ResourceUnavailable.ResourceUninitialized"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) GetSecretValue(request *GetSecretValueRequest) (response *GetSecretValueResponse, err error) {
    return c.GetSecretValueWithContext(context.Background(), request)
}

// GetSecretValue
// For user-defined credentials, this API is used to get the plaintext information of a credential by specifying the credential name and version.
//
// For Tencent Cloud service credentials such as MySQL credentials, this API is used to get the plaintext information of a previously rotated credential by specifying the credential name and historical version number. If you want to get the plaintext of the credential version currently in use, you need to specify the version number as `SSM_Current`.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSKMSERROR = "FailedOperation.AccessKmsError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_ROLENOTEXIST = "OperationDenied.RoleNotExist"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SECRETNOTEXIST = "ResourceNotFound.SecretNotExist"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  RESOURCEUNAVAILABLE_RESOURCEDISABLED = "ResourceUnavailable.ResourceDisabled"
//  RESOURCEUNAVAILABLE_RESOURCEPENDINGDELETED = "ResourceUnavailable.ResourcePendingDeleted"
//  RESOURCEUNAVAILABLE_RESOURCEUNINITIALIZED = "ResourceUnavailable.ResourceUninitialized"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) GetSecretValueWithContext(ctx context.Context, request *GetSecretValueRequest) (response *GetSecretValueResponse, err error) {
    if request == nil {
        request = NewGetSecretValueRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetSecretValue require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetSecretValueResponse()
    err = c.Send(request, response)
    return
}

func NewGetServiceStatusRequest() (request *GetServiceStatusRequest) {
    request = &GetServiceStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "GetServiceStatus")
    
    
    return
}

func NewGetServiceStatusResponse() (response *GetServiceStatusResponse) {
    response = &GetServiceStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetServiceStatus
// This API is used to obtain the SecretsManager service status of a user.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) GetServiceStatus(request *GetServiceStatusRequest) (response *GetServiceStatusResponse, err error) {
    return c.GetServiceStatusWithContext(context.Background(), request)
}

// GetServiceStatus
// This API is used to obtain the SecretsManager service status of a user.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) GetServiceStatusWithContext(ctx context.Context, request *GetServiceStatusRequest) (response *GetServiceStatusResponse, err error) {
    if request == nil {
        request = NewGetServiceStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetServiceStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetServiceStatusResponse()
    err = c.Send(request, response)
    return
}

func NewListSecretVersionIdsRequest() (request *ListSecretVersionIdsRequest) {
    request = &ListSecretVersionIdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "ListSecretVersionIds")
    
    
    return
}

func NewListSecretVersionIdsResponse() (response *ListSecretVersionIdsResponse) {
    response = &ListSecretVersionIdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListSecretVersionIds
// This API is used to obtain list of versions of a Secret.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ListSecretVersionIds(request *ListSecretVersionIdsRequest) (response *ListSecretVersionIdsResponse, err error) {
    return c.ListSecretVersionIdsWithContext(context.Background(), request)
}

// ListSecretVersionIds
// This API is used to obtain list of versions of a Secret.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ListSecretVersionIdsWithContext(ctx context.Context, request *ListSecretVersionIdsRequest) (response *ListSecretVersionIdsResponse, err error) {
    if request == nil {
        request = NewListSecretVersionIdsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListSecretVersionIds require credential")
    }

    request.SetContext(ctx)
    
    response = NewListSecretVersionIdsResponse()
    err = c.Send(request, response)
    return
}

func NewListSecretsRequest() (request *ListSecretsRequest) {
    request = &ListSecretsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "ListSecrets")
    
    
    return
}

func NewListSecretsResponse() (response *ListSecretsResponse) {
    response = &ListSecretsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListSecrets
// This API is used to obtain the detailed information list of all Secrets. You can specify the filter fields and sorting order as needed.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ListSecrets(request *ListSecretsRequest) (response *ListSecretsResponse, err error) {
    return c.ListSecretsWithContext(context.Background(), request)
}

// ListSecrets
// This API is used to obtain the detailed information list of all Secrets. You can specify the filter fields and sorting order as needed.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ListSecretsWithContext(ctx context.Context, request *ListSecretsRequest) (response *ListSecretsResponse, err error) {
    if request == nil {
        request = NewListSecretsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListSecrets require credential")
    }

    request.SetContext(ctx)
    
    response = NewListSecretsResponse()
    err = c.Send(request, response)
    return
}

func NewPutSecretValueRequest() (request *PutSecretValueRequest) {
    request = &PutSecretValueRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "PutSecretValue")
    
    
    return
}

func NewPutSecretValueResponse() (response *PutSecretValueResponse) {
    response = &PutSecretValueResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PutSecretValue
// This API adds the new version of the credential content under the specified credential. One credential can have up to 10 versions. New versions can be added to credentials only in `Enabled` or `Disabled` status.
//
// This API is only applicable to user-defined credentials but not Tencent Cloud service credentials.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSKMSERROR = "FailedOperation.AccessKmsError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED_AUTOROTATEDRESOURCE = "OperationDenied.AutoRotatedResource"
//  RESOURCEINUSE_VERSIONIDEXISTS = "ResourceInUse.VersionIdExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) PutSecretValue(request *PutSecretValueRequest) (response *PutSecretValueResponse, err error) {
    return c.PutSecretValueWithContext(context.Background(), request)
}

// PutSecretValue
// This API adds the new version of the credential content under the specified credential. One credential can have up to 10 versions. New versions can be added to credentials only in `Enabled` or `Disabled` status.
//
// This API is only applicable to user-defined credentials but not Tencent Cloud service credentials.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSKMSERROR = "FailedOperation.AccessKmsError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED_AUTOROTATEDRESOURCE = "OperationDenied.AutoRotatedResource"
//  RESOURCEINUSE_VERSIONIDEXISTS = "ResourceInUse.VersionIdExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) PutSecretValueWithContext(ctx context.Context, request *PutSecretValueRequest) (response *PutSecretValueResponse, err error) {
    if request == nil {
        request = NewPutSecretValueRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PutSecretValue require credential")
    }

    request.SetContext(ctx)
    
    response = NewPutSecretValueResponse()
    err = c.Send(request, response)
    return
}

func NewRestoreSecretRequest() (request *RestoreSecretRequest) {
    request = &RestoreSecretRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "RestoreSecret")
    
    
    return
}

func NewRestoreSecretResponse() (response *RestoreSecretResponse) {
    response = &RestoreSecretResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RestoreSecret
// This API is used to restore a `PendingDelete` Secret, canceling its scheduled deletion. The restored Secret will be in `Disabled` status. You can call the `EnableSecret` API to enable this Secret again.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) RestoreSecret(request *RestoreSecretRequest) (response *RestoreSecretResponse, err error) {
    return c.RestoreSecretWithContext(context.Background(), request)
}

// RestoreSecret
// This API is used to restore a `PendingDelete` Secret, canceling its scheduled deletion. The restored Secret will be in `Disabled` status. You can call the `EnableSecret` API to enable this Secret again.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) RestoreSecretWithContext(ctx context.Context, request *RestoreSecretRequest) (response *RestoreSecretResponse, err error) {
    if request == nil {
        request = NewRestoreSecretRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestoreSecret require credential")
    }

    request.SetContext(ctx)
    
    response = NewRestoreSecretResponse()
    err = c.Send(request, response)
    return
}

func NewRotateProductSecretRequest() (request *RotateProductSecretRequest) {
    request = &RotateProductSecretRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "RotateProductSecret")
    
    
    return
}

func NewRotateProductSecretResponse() (response *RotateProductSecretResponse) {
    response = &RotateProductSecretResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RotateProductSecret
// This API is used to rotate secrets for Tencent Cloud services or Tencent Cloud API key pairs.
//
// Note that only the secrets with the "Enabled" status can be rotated.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSKMSERROR = "FailedOperation.AccessKmsError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TAGKEYSDUPLICATED = "InvalidParameterValue.TagKeysDuplicated"
//  INVALIDPARAMETERVALUE_TAGSNOTEXISTED = "InvalidParameterValue.TagsNotExisted"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCESSKEYOVERLIMIT = "OperationDenied.AccessKeyOverLimit"
//  OPERATIONDENIED_AUTOROTATEDRESOURCE = "OperationDenied.AutoRotatedResource"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_SECRETEXISTS = "ResourceInUse.SecretExists"
//  RESOURCEINUSE_VERSIONIDEXISTS = "ResourceInUse.VersionIdExists"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  RESOURCEUNAVAILABLE_RESOURCEDISABLED = "ResourceUnavailable.ResourceDisabled"
//  RESOURCEUNAVAILABLE_RESOURCEPENDINGDELETED = "ResourceUnavailable.ResourcePendingDeleted"
//  RESOURCEUNAVAILABLE_RESOURCEUNINITIALIZED = "ResourceUnavailable.ResourceUninitialized"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ACCESSKMSERROR = "UnauthorizedOperation.AccessKmsError"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RotateProductSecret(request *RotateProductSecretRequest) (response *RotateProductSecretResponse, err error) {
    return c.RotateProductSecretWithContext(context.Background(), request)
}

// RotateProductSecret
// This API is used to rotate secrets for Tencent Cloud services or Tencent Cloud API key pairs.
//
// Note that only the secrets with the "Enabled" status can be rotated.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSKMSERROR = "FailedOperation.AccessKmsError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TAGKEYSDUPLICATED = "InvalidParameterValue.TagKeysDuplicated"
//  INVALIDPARAMETERVALUE_TAGSNOTEXISTED = "InvalidParameterValue.TagsNotExisted"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCESSKEYOVERLIMIT = "OperationDenied.AccessKeyOverLimit"
//  OPERATIONDENIED_AUTOROTATEDRESOURCE = "OperationDenied.AutoRotatedResource"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_SECRETEXISTS = "ResourceInUse.SecretExists"
//  RESOURCEINUSE_VERSIONIDEXISTS = "ResourceInUse.VersionIdExists"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  RESOURCEUNAVAILABLE_RESOURCEDISABLED = "ResourceUnavailable.ResourceDisabled"
//  RESOURCEUNAVAILABLE_RESOURCEPENDINGDELETED = "ResourceUnavailable.ResourcePendingDeleted"
//  RESOURCEUNAVAILABLE_RESOURCEUNINITIALIZED = "ResourceUnavailable.ResourceUninitialized"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ACCESSKMSERROR = "UnauthorizedOperation.AccessKmsError"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RotateProductSecretWithContext(ctx context.Context, request *RotateProductSecretRequest) (response *RotateProductSecretResponse, err error) {
    if request == nil {
        request = NewRotateProductSecretRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RotateProductSecret require credential")
    }

    request.SetContext(ctx)
    
    response = NewRotateProductSecretResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateDescriptionRequest() (request *UpdateDescriptionRequest) {
    request = &UpdateDescriptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "UpdateDescription")
    
    
    return
}

func NewUpdateDescriptionResponse() (response *UpdateDescriptionResponse) {
    response = &UpdateDescriptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateDescription
// This API is used to update the description of a Secret. This API can only update Secrets in `Enabled` or `Disabled` status.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) UpdateDescription(request *UpdateDescriptionRequest) (response *UpdateDescriptionResponse, err error) {
    return c.UpdateDescriptionWithContext(context.Background(), request)
}

// UpdateDescription
// This API is used to update the description of a Secret. This API can only update Secrets in `Enabled` or `Disabled` status.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) UpdateDescriptionWithContext(ctx context.Context, request *UpdateDescriptionRequest) (response *UpdateDescriptionResponse, err error) {
    if request == nil {
        request = NewUpdateDescriptionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateDescription require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateDescriptionResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateRotationStatusRequest() (request *UpdateRotationStatusRequest) {
    request = &UpdateRotationStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "UpdateRotationStatus")
    
    
    return
}

func NewUpdateRotationStatusResponse() (response *UpdateRotationStatusResponse) {
    response = &UpdateRotationStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateRotationStatus
// This API is used to set a Tencent Cloud service credential rotation policy, including the following parameters:
//
// Specifies whether to enable rotation
//
// Rotation frequency
//
// Rotation start time
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSKMSERROR = "FailedOperation.AccessKmsError"
//  FAILEDOPERATION_ROTATIONFORBIDDEN = "FailedOperation.RotationForbidden"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TAGKEYSDUPLICATED = "InvalidParameterValue.TagKeysDuplicated"
//  INVALIDPARAMETERVALUE_TAGSNOTEXISTED = "InvalidParameterValue.TagsNotExisted"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_AUTOROTATEDRESOURCE = "OperationDenied.AutoRotatedResource"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_SECRETEXISTS = "ResourceInUse.SecretExists"
//  RESOURCEINUSE_VERSIONIDEXISTS = "ResourceInUse.VersionIdExists"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  RESOURCEUNAVAILABLE_RESOURCEDISABLED = "ResourceUnavailable.ResourceDisabled"
//  RESOURCEUNAVAILABLE_RESOURCEPENDINGDELETED = "ResourceUnavailable.ResourcePendingDeleted"
//  RESOURCEUNAVAILABLE_RESOURCEUNINITIALIZED = "ResourceUnavailable.ResourceUninitialized"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ACCESSKMSERROR = "UnauthorizedOperation.AccessKmsError"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateRotationStatus(request *UpdateRotationStatusRequest) (response *UpdateRotationStatusResponse, err error) {
    return c.UpdateRotationStatusWithContext(context.Background(), request)
}

// UpdateRotationStatus
// This API is used to set a Tencent Cloud service credential rotation policy, including the following parameters:
//
// Specifies whether to enable rotation
//
// Rotation frequency
//
// Rotation start time
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSKMSERROR = "FailedOperation.AccessKmsError"
//  FAILEDOPERATION_ROTATIONFORBIDDEN = "FailedOperation.RotationForbidden"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TAGKEYSDUPLICATED = "InvalidParameterValue.TagKeysDuplicated"
//  INVALIDPARAMETERVALUE_TAGSNOTEXISTED = "InvalidParameterValue.TagsNotExisted"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_AUTOROTATEDRESOURCE = "OperationDenied.AutoRotatedResource"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_SECRETEXISTS = "ResourceInUse.SecretExists"
//  RESOURCEINUSE_VERSIONIDEXISTS = "ResourceInUse.VersionIdExists"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  RESOURCEUNAVAILABLE_RESOURCEDISABLED = "ResourceUnavailable.ResourceDisabled"
//  RESOURCEUNAVAILABLE_RESOURCEPENDINGDELETED = "ResourceUnavailable.ResourcePendingDeleted"
//  RESOURCEUNAVAILABLE_RESOURCEUNINITIALIZED = "ResourceUnavailable.ResourceUninitialized"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ACCESSKMSERROR = "UnauthorizedOperation.AccessKmsError"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateRotationStatusWithContext(ctx context.Context, request *UpdateRotationStatusRequest) (response *UpdateRotationStatusResponse, err error) {
    if request == nil {
        request = NewUpdateRotationStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateRotationStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateRotationStatusResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateSecretRequest() (request *UpdateSecretRequest) {
    request = &UpdateSecretRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "UpdateSecret")
    
    
    return
}

func NewUpdateSecretResponse() (response *UpdateSecretResponse) {
    response = &UpdateSecretResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateSecret
// This API is used to update the credential content of the specified credential name and version number. Calling this API will encrypt the content of the new credential and overwrite the old content. Only credentials in `Enabled` or `Disabled` status can be updated.
//
// This API is only applicable to user-defined credentials but not Tencent Cloud service credentials.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSKMSERROR = "FailedOperation.AccessKmsError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_AUTOROTATEDRESOURCE = "OperationDenied.AutoRotatedResource"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) UpdateSecret(request *UpdateSecretRequest) (response *UpdateSecretResponse, err error) {
    return c.UpdateSecretWithContext(context.Background(), request)
}

// UpdateSecret
// This API is used to update the credential content of the specified credential name and version number. Calling this API will encrypt the content of the new credential and overwrite the old content. Only credentials in `Enabled` or `Disabled` status can be updated.
//
// This API is only applicable to user-defined credentials but not Tencent Cloud service credentials.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSKMSERROR = "FailedOperation.AccessKmsError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_AUTOROTATEDRESOURCE = "OperationDenied.AutoRotatedResource"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) UpdateSecretWithContext(ctx context.Context, request *UpdateSecretRequest) (response *UpdateSecretResponse, err error) {
    if request == nil {
        request = NewUpdateSecretRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateSecret require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateSecretResponse()
    err = c.Send(request, response)
    return
}
