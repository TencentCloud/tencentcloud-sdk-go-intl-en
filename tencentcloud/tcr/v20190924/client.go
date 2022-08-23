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
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
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
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
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
func (c *Client) CreateReplicationInstance(request *CreateReplicationInstanceRequest) (response *CreateReplicationInstanceResponse, err error) {
    return c.CreateReplicationInstanceWithContext(context.Background(), request)
}

// CreateReplicationInstance
// This API is used to create a replication instance.
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
