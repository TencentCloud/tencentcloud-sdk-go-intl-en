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

package v20220928

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2022-09-28"

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


func NewAllocateCustomerCreditRequest() (request *AllocateCustomerCreditRequest) {
    request = &AllocateCustomerCreditRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "AllocateCustomerCredit")
    
    
    return
}

func NewAllocateCustomerCreditResponse() (response *AllocateCustomerCreditResponse) {
    response = &AllocateCustomerCreditResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AllocateCustomerCredit
// This API is used for a partner to set credit for a customer, such as increasing or lowering the credit and setting it to 0.
//
// 1. The credit is valid permanently and will not be zeroed regularly.
//
// 2. The customer's service will be suspended when its available credit is set to 0, so caution should be exercised with this operation.
//
// 3. To prevent the customer from making new purchases without affecting their use of previously purchased products, the partner can set their available credit to 0 after obtaining the non-stop feature privilege from the channel manager.
//
// 4. The set credit is an increment of the current available credit and cannot exceed the remaining allocable credit. Setting the credit to a negative value indicates that it will be repossessed. The available credit can be set to 0 at the minimum.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) AllocateCustomerCredit(request *AllocateCustomerCreditRequest) (response *AllocateCustomerCreditResponse, err error) {
    return c.AllocateCustomerCreditWithContext(context.Background(), request)
}

// AllocateCustomerCredit
// This API is used for a partner to set credit for a customer, such as increasing or lowering the credit and setting it to 0.
//
// 1. The credit is valid permanently and will not be zeroed regularly.
//
// 2. The customer's service will be suspended when its available credit is set to 0, so caution should be exercised with this operation.
//
// 3. To prevent the customer from making new purchases without affecting their use of previously purchased products, the partner can set their available credit to 0 after obtaining the non-stop feature privilege from the channel manager.
//
// 4. The set credit is an increment of the current available credit and cannot exceed the remaining allocable credit. Setting the credit to a negative value indicates that it will be repossessed. The available credit can be set to 0 at the minimum.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) AllocateCustomerCreditWithContext(ctx context.Context, request *AllocateCustomerCreditRequest) (response *AllocateCustomerCreditResponse, err error) {
    if request == nil {
        request = NewAllocateCustomerCreditRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AllocateCustomerCredit require credential")
    }

    request.SetContext(ctx)
    
    response = NewAllocateCustomerCreditResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAccountRequest() (request *CreateAccountRequest) {
    request = &CreateAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "CreateAccount")
    
    
    return
}

func NewCreateAccountResponse() (response *CreateAccountResponse) {
    response = &CreateAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAccount
// This API is used to create a Tencent Cloud account on the partner platform for a customer. After registration, the customer will be automatically bound to the partner account.
//
// 
//
// Notes:<br>
//
// 1. The partner should verify the entered email address and mobile number for creating a Tencent Cloud account.<br>
//
// 2. The customer needs to complete personal information after the first login.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_MAILISREGISTERED = "FailedOperation.MailIsRegistered"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ACCOUNTTYPECONTENTINCORRECT = "InvalidParameter.AccountTypeContentIncorrect"
//  INVALIDPARAMETER_AREACONTENTINCORRECT = "InvalidParameter.AreaContentIncorrect"
//  INVALIDPARAMETER_AREAFORMATINCORRECT = "InvalidParameter.AreaFormatIncorrect"
//  INVALIDPARAMETER_CONFIRMPASSWORDCONTENTINCORRECT = "InvalidParameter.ConfirmPasswordContentIncorrect"
//  INVALIDPARAMETER_COUNTRYCODECONTENTINCORRECT = "InvalidParameter.CountryCodeContentIncorrect"
//  INVALIDPARAMETER_COUNTRYCODEFORMATINCORRECT = "InvalidParameter.CountryCodeFormatIncorrect"
//  INVALIDPARAMETER_MAILFORMATINCORRECT = "InvalidParameter.MailFormatIncorrect"
//  INVALIDPARAMETER_PASSWORDCONTENTINCORRECT = "InvalidParameter.PasswordContentIncorrect"
//  INVALIDPARAMETER_PASSWORDFORMATINCORRECT = "InvalidParameter.PasswordFormatIncorrect"
//  INVALIDPARAMETER_PHONENUMFORMATINCORRECT = "InvalidParameter.PhoneNumFormatIncorrect"
//  INVALIDPARAMETERVALUE_ACCOUNTTYPEEMPTY = "InvalidParameterValue.AccountTypeEmpty"
//  INVALIDPARAMETERVALUE_AREAEMPTY = "InvalidParameterValue.AreaEmpty"
//  INVALIDPARAMETERVALUE_COUNTRYCODEEMPTY = "InvalidParameterValue.CountryCodeEmpty"
//  INVALIDPARAMETERVALUE_MAILEMPTY = "InvalidParameterValue.MailEmpty"
//  INVALIDPARAMETERVALUE_PASSWORDEMPTY = "InvalidParameterValue.PasswordEmpty"
//  INVALIDPARAMETERVALUE_PHONENUMEMPTY = "InvalidParameterValue.PhoneNumEmpty"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAccount(request *CreateAccountRequest) (response *CreateAccountResponse, err error) {
    return c.CreateAccountWithContext(context.Background(), request)
}

// CreateAccount
// This API is used to create a Tencent Cloud account on the partner platform for a customer. After registration, the customer will be automatically bound to the partner account.
//
// 
//
// Notes:<br>
//
// 1. The partner should verify the entered email address and mobile number for creating a Tencent Cloud account.<br>
//
// 2. The customer needs to complete personal information after the first login.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_MAILISREGISTERED = "FailedOperation.MailIsRegistered"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ACCOUNTTYPECONTENTINCORRECT = "InvalidParameter.AccountTypeContentIncorrect"
//  INVALIDPARAMETER_AREACONTENTINCORRECT = "InvalidParameter.AreaContentIncorrect"
//  INVALIDPARAMETER_AREAFORMATINCORRECT = "InvalidParameter.AreaFormatIncorrect"
//  INVALIDPARAMETER_CONFIRMPASSWORDCONTENTINCORRECT = "InvalidParameter.ConfirmPasswordContentIncorrect"
//  INVALIDPARAMETER_COUNTRYCODECONTENTINCORRECT = "InvalidParameter.CountryCodeContentIncorrect"
//  INVALIDPARAMETER_COUNTRYCODEFORMATINCORRECT = "InvalidParameter.CountryCodeFormatIncorrect"
//  INVALIDPARAMETER_MAILFORMATINCORRECT = "InvalidParameter.MailFormatIncorrect"
//  INVALIDPARAMETER_PASSWORDCONTENTINCORRECT = "InvalidParameter.PasswordContentIncorrect"
//  INVALIDPARAMETER_PASSWORDFORMATINCORRECT = "InvalidParameter.PasswordFormatIncorrect"
//  INVALIDPARAMETER_PHONENUMFORMATINCORRECT = "InvalidParameter.PhoneNumFormatIncorrect"
//  INVALIDPARAMETERVALUE_ACCOUNTTYPEEMPTY = "InvalidParameterValue.AccountTypeEmpty"
//  INVALIDPARAMETERVALUE_AREAEMPTY = "InvalidParameterValue.AreaEmpty"
//  INVALIDPARAMETERVALUE_COUNTRYCODEEMPTY = "InvalidParameterValue.CountryCodeEmpty"
//  INVALIDPARAMETERVALUE_MAILEMPTY = "InvalidParameterValue.MailEmpty"
//  INVALIDPARAMETERVALUE_PASSWORDEMPTY = "InvalidParameterValue.PasswordEmpty"
//  INVALIDPARAMETERVALUE_PHONENUMEMPTY = "InvalidParameterValue.PhoneNumEmpty"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAccountWithContext(ctx context.Context, request *CreateAccountRequest) (response *CreateAccountResponse, err error) {
    if request == nil {
        request = NewCreateAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAccountResponse()
    err = c.Send(request, response)
    return
}

func NewGetCountryCodesRequest() (request *GetCountryCodesRequest) {
    request = &GetCountryCodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "GetCountryCodes")
    
    
    return
}

func NewGetCountryCodesResponse() (response *GetCountryCodesResponse) {
    response = &GetCountryCodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetCountryCodes
// This API is used to obtain country/region codes.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) GetCountryCodes(request *GetCountryCodesRequest) (response *GetCountryCodesResponse, err error) {
    return c.GetCountryCodesWithContext(context.Background(), request)
}

// GetCountryCodes
// This API is used to obtain country/region codes.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) GetCountryCodesWithContext(ctx context.Context, request *GetCountryCodesRequest) (response *GetCountryCodesResponse, err error) {
    if request == nil {
        request = NewGetCountryCodesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetCountryCodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetCountryCodesResponse()
    err = c.Send(request, response)
    return
}

func NewQueryCreditAllocationHistoryRequest() (request *QueryCreditAllocationHistoryRequest) {
    request = &QueryCreditAllocationHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "QueryCreditAllocationHistory")
    
    
    return
}

func NewQueryCreditAllocationHistoryResponse() (response *QueryCreditAllocationHistoryResponse) {
    response = &QueryCreditAllocationHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryCreditAllocationHistory
// This API is used to query all the credit allocation records of a single customer.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) QueryCreditAllocationHistory(request *QueryCreditAllocationHistoryRequest) (response *QueryCreditAllocationHistoryResponse, err error) {
    return c.QueryCreditAllocationHistoryWithContext(context.Background(), request)
}

// QueryCreditAllocationHistory
// This API is used to query all the credit allocation records of a single customer.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) QueryCreditAllocationHistoryWithContext(ctx context.Context, request *QueryCreditAllocationHistoryRequest) (response *QueryCreditAllocationHistoryResponse, err error) {
    if request == nil {
        request = NewQueryCreditAllocationHistoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryCreditAllocationHistory require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryCreditAllocationHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewQueryCreditByUinListRequest() (request *QueryCreditByUinListRequest) {
    request = &QueryCreditByUinListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "QueryCreditByUinList")
    
    
    return
}

func NewQueryCreditByUinListResponse() (response *QueryCreditByUinListResponse) {
    response = &QueryCreditByUinListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryCreditByUinList
// This API is used to query the credit of users in the list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) QueryCreditByUinList(request *QueryCreditByUinListRequest) (response *QueryCreditByUinListResponse, err error) {
    return c.QueryCreditByUinListWithContext(context.Background(), request)
}

// QueryCreditByUinList
// This API is used to query the credit of users in the list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) QueryCreditByUinListWithContext(ctx context.Context, request *QueryCreditByUinListRequest) (response *QueryCreditByUinListResponse, err error) {
    if request == nil {
        request = NewQueryCreditByUinListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryCreditByUinList require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryCreditByUinListResponse()
    err = c.Send(request, response)
    return
}

func NewQueryCustomersCreditRequest() (request *QueryCustomersCreditRequest) {
    request = &QueryCustomersCreditRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "QueryCustomersCredit")
    
    
    return
}

func NewQueryCustomersCreditResponse() (response *QueryCustomersCreditResponse) {
    response = &QueryCustomersCreditResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryCustomersCredit
// This API is used for a partner to the credits and basic information of cutomers.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) QueryCustomersCredit(request *QueryCustomersCreditRequest) (response *QueryCustomersCreditResponse, err error) {
    return c.QueryCustomersCreditWithContext(context.Background(), request)
}

// QueryCustomersCredit
// This API is used for a partner to the credits and basic information of cutomers.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) QueryCustomersCreditWithContext(ctx context.Context, request *QueryCustomersCreditRequest) (response *QueryCustomersCreditResponse, err error) {
    if request == nil {
        request = NewQueryCustomersCreditRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryCustomersCredit require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryCustomersCreditResponse()
    err = c.Send(request, response)
    return
}

func NewQueryDirectCustomersCreditRequest() (request *QueryDirectCustomersCreditRequest) {
    request = &QueryDirectCustomersCreditRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "QueryDirectCustomersCredit")
    
    
    return
}

func NewQueryDirectCustomersCreditResponse() (response *QueryDirectCustomersCreditResponse) {
    response = &QueryDirectCustomersCreditResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryDirectCustomersCredit
// This API is used to query the credits of direct customers.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) QueryDirectCustomersCredit(request *QueryDirectCustomersCreditRequest) (response *QueryDirectCustomersCreditResponse, err error) {
    return c.QueryDirectCustomersCreditWithContext(context.Background(), request)
}

// QueryDirectCustomersCredit
// This API is used to query the credits of direct customers.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) QueryDirectCustomersCreditWithContext(ctx context.Context, request *QueryDirectCustomersCreditRequest) (response *QueryDirectCustomersCreditResponse, err error) {
    if request == nil {
        request = NewQueryDirectCustomersCreditRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryDirectCustomersCredit require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryDirectCustomersCreditResponse()
    err = c.Send(request, response)
    return
}

func NewQueryPartnerCreditRequest() (request *QueryPartnerCreditRequest) {
    request = &QueryPartnerCreditRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "QueryPartnerCredit")
    
    
    return
}

func NewQueryPartnerCreditResponse() (response *QueryPartnerCreditResponse) {
    response = &QueryPartnerCreditResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryPartnerCredit
// This API is used for a partner to query its own total credit, available credit, and used credit in USD.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) QueryPartnerCredit(request *QueryPartnerCreditRequest) (response *QueryPartnerCreditResponse, err error) {
    return c.QueryPartnerCreditWithContext(context.Background(), request)
}

// QueryPartnerCredit
// This API is used for a partner to query its own total credit, available credit, and used credit in USD.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) QueryPartnerCreditWithContext(ctx context.Context, request *QueryPartnerCreditRequest) (response *QueryPartnerCreditResponse, err error) {
    if request == nil {
        request = NewQueryPartnerCreditRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryPartnerCredit require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryPartnerCreditResponse()
    err = c.Send(request, response)
    return
}
