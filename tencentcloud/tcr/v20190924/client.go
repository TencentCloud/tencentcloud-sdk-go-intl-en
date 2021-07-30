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
    if request == nil {
        request = NewCheckInstanceRequest()
    }
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
    if request == nil {
        request = NewCreateImmutableTagRulesRequest()
    }
    response = NewCreateImmutableTagRulesResponse()
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
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateMultipleSecurityPolicy(request *CreateMultipleSecurityPolicyRequest) (response *CreateMultipleSecurityPolicyResponse, err error) {
    if request == nil {
        request = NewCreateMultipleSecurityPolicyRequest()
    }
    response = NewCreateMultipleSecurityPolicyResponse()
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
    if request == nil {
        request = NewDeleteImmutableTagRulesRequest()
    }
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
    if request == nil {
        request = NewDeleteMultipleSecurityPolicyRequest()
    }
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
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
func (c *Client) DescribeImmutableTagRules(request *DescribeImmutableTagRulesRequest) (response *DescribeImmutableTagRulesResponse, err error) {
    if request == nil {
        request = NewDescribeImmutableTagRulesRequest()
    }
    response = NewDescribeImmutableTagRulesResponse()
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
    if request == nil {
        request = NewModifyImmutableTagRulesRequest()
    }
    response = NewModifyImmutableTagRulesResponse()
    err = c.Send(request, response)
    return
}
