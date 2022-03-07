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

package v20190319

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2019-03-19"

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


func NewDescribeAuditTracksRequest() (request *DescribeAuditTracksRequest) {
    request = &DescribeAuditTracksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "DescribeAuditTracks")
    
    
    return
}

func NewDescribeAuditTracksResponse() (response *DescribeAuditTracksResponse) {
    response = &DescribeAuditTracksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAuditTracks
// This API is used to query the CloudAudit tracking set list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ALIASALREADYEXISTS = "InvalidParameterValue.AliasAlreadyExists"
//  LIMITEXCEEDED_OVERAMOUNT = "LimitExceeded.OverAmount"
func (c *Client) DescribeAuditTracks(request *DescribeAuditTracksRequest) (response *DescribeAuditTracksResponse, err error) {
    return c.DescribeAuditTracksWithContext(context.Background(), request)
}

// DescribeAuditTracks
// This API is used to query the CloudAudit tracking set list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ALIASALREADYEXISTS = "InvalidParameterValue.AliasAlreadyExists"
//  LIMITEXCEEDED_OVERAMOUNT = "LimitExceeded.OverAmount"
func (c *Client) DescribeAuditTracksWithContext(ctx context.Context, request *DescribeAuditTracksRequest) (response *DescribeAuditTracksResponse, err error) {
    if request == nil {
        request = NewDescribeAuditTracksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAuditTracks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuditTracksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEventsRequest() (request *DescribeEventsRequest) {
    request = &DescribeEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "DescribeEvents")
    
    
    return
}

func NewDescribeEventsResponse() (response *DescribeEventsResponse) {
    response = &DescribeEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEvents
// This API is used to query CloudAudit logs.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeEvents(request *DescribeEventsRequest) (response *DescribeEventsResponse, err error) {
    return c.DescribeEventsWithContext(context.Background(), request)
}

// DescribeEvents
// This API is used to query CloudAudit logs.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeEventsWithContext(ctx context.Context, request *DescribeEventsRequest) (response *DescribeEventsResponse, err error) {
    if request == nil {
        request = NewDescribeEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEventsResponse()
    err = c.Send(request, response)
    return
}
