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

package v20210409

import (
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2021-04-09"

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


func NewAssignClientCreditRequest() (request *AssignClientCreditRequest) {
    request = &AssignClientCreditRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ip", APIVersion, "AssignClientCredit")
    
    return
}

func NewAssignClientCreditResponse() (response *AssignClientCreditResponse) {
    response = &AssignClientCreditResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AssignClientCredit
// This API is used for a partner to set credit for a customer, such as increasing or lowering the credit and setting it to 0.
//
// 1. The credit is valid permanently and will not be zeroed regularly.
//
// 2. The customer's service will be suspended when its available credit sets to 0, so caution should be exercised with this operation.
//
// 3. To prevent the customer from making new purchases without affecting their use of previously purchased products, the partner can set their available credit to 0 after obtaining the non-stop feature privilege from the channel manager.
//
// 4. The set credit is an increase to the current available credit and cannot exceed the remaining allocable credit. Setting the credit to a negative value indicates to repossess it. The available credit can be set to 0 at the minimum.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) AssignClientCredit(request *AssignClientCreditRequest) (response *AssignClientCreditResponse, err error) {
    if request == nil {
        request = NewAssignClientCreditRequest()
    }
    response = NewAssignClientCreditResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAccountRequest() (request *CreateAccountRequest) {
    request = &CreateAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ip", APIVersion, "CreateAccount")
    
    return
}

func NewCreateAccountResponse() (response *CreateAccountResponse) {
    response = &CreateAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAccount
// This API is used to create a Tencent Cloud account in the International Partner platform for a customer. After registration, the customer will be automatically bound to the partner account.
//
// 
//
// Notes:<br>
//
// 1. To create the Tencent Cloud account, the partner should enter and verify the customer’s email address and mobile number.<br>
//
// 2. The customer needs to complete personal information after the first login.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_MAILISREGISTERED = "FailedOperation.MailIsRegistered"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAccount(request *CreateAccountRequest) (response *CreateAccountResponse, err error) {
    if request == nil {
        request = NewCreateAccountRequest()
    }
    response = NewCreateAccountResponse()
    err = c.Send(request, response)
    return
}

func NewGetCountryCodesRequest() (request *GetCountryCodesRequest) {
    request = &GetCountryCodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ip", APIVersion, "GetCountryCodes")
    
    return
}

func NewGetCountryCodesResponse() (response *GetCountryCodesResponse) {
    response = &GetCountryCodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetCountryCodes
// This API is used to obtain country and region codes.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) GetCountryCodes(request *GetCountryCodesRequest) (response *GetCountryCodesResponse, err error) {
    if request == nil {
        request = NewGetCountryCodesRequest()
    }
    response = NewGetCountryCodesResponse()
    err = c.Send(request, response)
    return
}

func NewQueryAgentCreditRequest() (request *QueryAgentCreditRequest) {
    request = &QueryAgentCreditRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ip", APIVersion, "QueryAgentCredit")
    
    return
}

func NewQueryAgentCreditResponse() (response *QueryAgentCreditResponse) {
    response = &QueryAgentCreditResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryAgentCredit
// This API is used for a partner to query its own total credit, available credit, and used credit in USD.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) QueryAgentCredit(request *QueryAgentCreditRequest) (response *QueryAgentCreditResponse, err error) {
    if request == nil {
        request = NewQueryAgentCreditRequest()
    }
    response = NewQueryAgentCreditResponse()
    err = c.Send(request, response)
    return
}

func NewQueryClientListRequest() (request *QueryClientListRequest) {
    request = &QueryClientListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ip", APIVersion, "QueryClientList")
    
    return
}

func NewQueryClientListResponse() (response *QueryClientListResponse) {
    response = &QueryClientListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryClientList
// This API is used for a partner to query a customer's credit and basic information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) QueryClientList(request *QueryClientListRequest) (response *QueryClientListResponse, err error) {
    if request == nil {
        request = NewQueryClientListRequest()
    }
    response = NewQueryClientListResponse()
    err = c.Send(request, response)
    return
}

func NewQueryCreditHistoryRequest() (request *QueryCreditHistoryRequest) {
    request = &QueryCreditHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ip", APIVersion, "QueryCreditHistory")
    
    return
}

func NewQueryCreditHistoryResponse() (response *QueryCreditHistoryResponse) {
    response = &QueryCreditHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryCreditHistory
// This API is used to query all the credit allocation records of a single customer.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) QueryCreditHistory(request *QueryCreditHistoryRequest) (response *QueryCreditHistoryResponse, err error) {
    if request == nil {
        request = NewQueryCreditHistoryRequest()
    }
    response = NewQueryCreditHistoryResponse()
    err = c.Send(request, response)
    return
}
