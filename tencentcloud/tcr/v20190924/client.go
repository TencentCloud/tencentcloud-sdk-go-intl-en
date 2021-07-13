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

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
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
