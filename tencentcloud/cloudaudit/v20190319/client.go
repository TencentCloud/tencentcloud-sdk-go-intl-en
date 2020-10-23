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

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewCreateAuditRequest() (request *CreateAuditRequest) {
    request = &CreateAuditRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "CreateAudit")
    return
}

func NewCreateAuditResponse() (response *CreateAuditResponse) {
    response = &CreateAuditResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Parameter requirements:
// 1. If the value of `IsCreateNewBucket` exists, `cosRegion` and `cosBucketName` are required.
// 2. If the value of `IsEnableCmqNotify` is 1, `IsCreateNewQueue`, `CmqRegion`, and `CmqQueueName` are required.
// 3. If the value of `IsEnableCmqNotify` is 0, `IsCreateNewQueue`, `CmqRegion`, and `CmqQueueName` cannot be passed in.
// 4. If the value of `IsEnableKmsEncry` is 1, `KmsRegion` and `KeyId` are required.
func (c *Client) CreateAudit(request *CreateAuditRequest) (response *CreateAuditResponse, err error) {
    if request == nil {
        request = NewCreateAuditRequest()
    }
    response = NewCreateAuditResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAuditRequest() (request *DeleteAuditRequest) {
    request = &DeleteAuditRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "DeleteAudit")
    return
}

func NewDeleteAuditResponse() (response *DeleteAuditResponse) {
    response = &DeleteAuditResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete a tracking set.
func (c *Client) DeleteAudit(request *DeleteAuditRequest) (response *DeleteAuditResponse, err error) {
    if request == nil {
        request = NewDeleteAuditRequest()
    }
    response = NewDeleteAuditResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuditRequest() (request *DescribeAuditRequest) {
    request = &DescribeAuditRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "DescribeAudit")
    return
}

func NewDescribeAuditResponse() (response *DescribeAuditResponse) {
    response = &DescribeAuditResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the details of a tracking set.
func (c *Client) DescribeAudit(request *DescribeAuditRequest) (response *DescribeAuditResponse, err error) {
    if request == nil {
        request = NewDescribeAuditRequest()
    }
    response = NewDescribeAuditResponse()
    err = c.Send(request, response)
    return
}

func NewGetAttributeKeyRequest() (request *GetAttributeKeyRequest) {
    request = &GetAttributeKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "GetAttributeKey")
    return
}

func NewGetAttributeKeyResponse() (response *GetAttributeKeyResponse) {
    response = &GetAttributeKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the valid values range of `AttributeKey`.
func (c *Client) GetAttributeKey(request *GetAttributeKeyRequest) (response *GetAttributeKeyResponse, err error) {
    if request == nil {
        request = NewGetAttributeKeyRequest()
    }
    response = NewGetAttributeKeyResponse()
    err = c.Send(request, response)
    return
}

func NewInquireAuditCreditRequest() (request *InquireAuditCreditRequest) {
    request = &InquireAuditCreditRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "InquireAuditCredit")
    return
}

func NewInquireAuditCreditResponse() (response *InquireAuditCreditResponse) {
    response = &InquireAuditCreditResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the number of tracking sets that can be created.
func (c *Client) InquireAuditCredit(request *InquireAuditCreditRequest) (response *InquireAuditCreditResponse, err error) {
    if request == nil {
        request = NewInquireAuditCreditRequest()
    }
    response = NewInquireAuditCreditResponse()
    err = c.Send(request, response)
    return
}

func NewListAuditsRequest() (request *ListAuditsRequest) {
    request = &ListAuditsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "ListAudits")
    return
}

func NewListAuditsResponse() (response *ListAuditsResponse) {
    response = &ListAuditsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the summary of tracking sets.
func (c *Client) ListAudits(request *ListAuditsRequest) (response *ListAuditsResponse, err error) {
    if request == nil {
        request = NewListAuditsRequest()
    }
    response = NewListAuditsResponse()
    err = c.Send(request, response)
    return
}

func NewListCmqEnableRegionRequest() (request *ListCmqEnableRegionRequest) {
    request = &ListCmqEnableRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "ListCmqEnableRegion")
    return
}

func NewListCmqEnableRegionResponse() (response *ListCmqEnableRegionResponse) {
    response = &ListCmqEnableRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query CloudAudit-enabled CMQ AZs.
func (c *Client) ListCmqEnableRegion(request *ListCmqEnableRegionRequest) (response *ListCmqEnableRegionResponse, err error) {
    if request == nil {
        request = NewListCmqEnableRegionRequest()
    }
    response = NewListCmqEnableRegionResponse()
    err = c.Send(request, response)
    return
}

func NewListCosEnableRegionRequest() (request *ListCosEnableRegionRequest) {
    request = &ListCosEnableRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "ListCosEnableRegion")
    return
}

func NewListCosEnableRegionResponse() (response *ListCosEnableRegionResponse) {
    response = &ListCosEnableRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query CloudAudit-enabled COS AZs.
func (c *Client) ListCosEnableRegion(request *ListCosEnableRegionRequest) (response *ListCosEnableRegionResponse, err error) {
    if request == nil {
        request = NewListCosEnableRegionRequest()
    }
    response = NewListCosEnableRegionResponse()
    err = c.Send(request, response)
    return
}

func NewLookUpEventsRequest() (request *LookUpEventsRequest) {
    request = &LookUpEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "LookUpEvents")
    return
}

func NewLookUpEventsResponse() (response *LookUpEventsResponse) {
    response = &LookUpEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to search for operation logs to help query relevant operation information.
func (c *Client) LookUpEvents(request *LookUpEventsRequest) (response *LookUpEventsResponse, err error) {
    if request == nil {
        request = NewLookUpEventsRequest()
    }
    response = NewLookUpEventsResponse()
    err = c.Send(request, response)
    return
}

func NewStartLoggingRequest() (request *StartLoggingRequest) {
    request = &StartLoggingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "StartLogging")
    return
}

func NewStartLoggingResponse() (response *StartLoggingResponse) {
    response = &StartLoggingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to enable a tracking set.
func (c *Client) StartLogging(request *StartLoggingRequest) (response *StartLoggingResponse, err error) {
    if request == nil {
        request = NewStartLoggingRequest()
    }
    response = NewStartLoggingResponse()
    err = c.Send(request, response)
    return
}

func NewStopLoggingRequest() (request *StopLoggingRequest) {
    request = &StopLoggingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "StopLogging")
    return
}

func NewStopLoggingResponse() (response *StopLoggingResponse) {
    response = &StopLoggingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to disable a tracking set.
func (c *Client) StopLogging(request *StopLoggingRequest) (response *StopLoggingResponse, err error) {
    if request == nil {
        request = NewStopLoggingRequest()
    }
    response = NewStopLoggingResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAuditRequest() (request *UpdateAuditRequest) {
    request = &UpdateAuditRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "UpdateAudit")
    return
}

func NewUpdateAuditResponse() (response *UpdateAuditResponse) {
    response = &UpdateAuditResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Parameter requirements:
// 1. If the value of `IsCreateNewBucket` exists, `cosRegion` and `cosBucketName` are required.
// 2. If the value of `IsEnableCmqNotify` is 1, `IsCreateNewQueue`, `CmqRegion`, and `CmqQueueName` are required.
// 3. If the value of `IsEnableCmqNotify` is 0, `IsCreateNewQueue`, `CmqRegion`, and `CmqQueueName` cannot be passed in.
// 4. If the value of `IsEnableKmsEncry` is 1, `KmsRegion` and `KeyId` are required.
func (c *Client) UpdateAudit(request *UpdateAuditRequest) (response *UpdateAuditResponse, err error) {
    if request == nil {
        request = NewUpdateAuditRequest()
    }
    response = NewUpdateAuditResponse()
    err = c.Send(request, response)
    return
}
