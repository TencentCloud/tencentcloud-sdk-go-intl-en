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


func NewAllocateCreditPoolRequest() (request *AllocateCreditPoolRequest) {
    request = &AllocateCreditPoolRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "AllocateCreditPool")
    
    
    return
}

func NewAllocateCreditPoolResponse() (response *AllocateCreditPoolResponse) {
    response = &AllocateCreditPoolResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AllocateCreditPool
// This API is used to allocate credit pools to second-level resellers by distributors.
//
// error code that may be returned:
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) AllocateCreditPool(request *AllocateCreditPoolRequest) (response *AllocateCreditPoolResponse, err error) {
    return c.AllocateCreditPoolWithContext(context.Background(), request)
}

// AllocateCreditPool
// This API is used to allocate credit pools to second-level resellers by distributors.
//
// error code that may be returned:
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) AllocateCreditPoolWithContext(ctx context.Context, request *AllocateCreditPoolRequest) (response *AllocateCreditPoolResponse, err error) {
    if request == nil {
        request = NewAllocateCreditPoolRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AllocateCreditPool require credential")
    }

    request.SetContext(ctx)
    
    response = NewAllocateCreditPoolResponse()
    err = c.Send(request, response)
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
//  INVALIDPARAMETERVALUE_CREDITAMOUNTOUTOFRANGE = "InvalidParameterValue.CreditAmountOutOfRange"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
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
//  INVALIDPARAMETERVALUE_CREDITAMOUNTOUTOFRANGE = "InvalidParameterValue.CreditAmountOutOfRange"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
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

func NewApproveClientApplyRequest() (request *ApproveClientApplyRequest) {
    request = &ApproveClientApplyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "ApproveClientApply")
    
    
    return
}

func NewApproveClientApplyResponse() (response *ApproveClientApplyResponse) {
    response = &ApproveClientApplyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ApproveClientApply
// Description: This API is used by resellers to review applications to become sub-customers. Note: This API is used to apply for the allowlist. If needed, please contact your business representative.
//
// 
//
// Callable roles: Reseller, Distributer, Second-level reseller
//
// error code that may be returned:
//  FAILEDOPERATION_BINDORGANIZATIONACCOUNT = "FailedOperation.BindOrganizationAccount"
//  FAILEDOPERATION_CLIENTBALANCEISLESSOREQUALZERO = "FailedOperation.ClientBalanceIsLessOrEqualZero"
//  FAILEDOPERATION_CLIENTBUYSP = "FailedOperation.ClientBuySP"
//  FAILEDOPERATION_CLIENTCREATESHAREUNIT = "FailedOperation.ClientCreateShareUnit"
//  FAILEDOPERATION_CLIENTJOINSHAREUNIT = "FailedOperation.ClientJoinShareUnit"
//  FAILEDOPERATION_CLIENTNOTAPPLY = "FailedOperation.ClientNotApply"
//  FAILEDOPERATION_EXCEEDMAXBINDCOUNT = "FailedOperation.ExceedMaxBindCount"
//  FAILEDOPERATION_UINALREADYKA = "FailedOperation.UinAlreadyKA"
//  FAILEDOPERATION_UINNOTRESELLER = "FailedOperation.UinNotReseller"
//  INVALIDPARAMETERVALUE_INVALIDUIN = "InvalidParameterValue.InvalidUin"
//  INVALIDPARAMETERVALUE_UINALREADYCLIENT = "InvalidParameterValue.UinAlreadyClient"
//  INVALIDPARAMETERVALUE_UINISSUBACCOUNT = "InvalidParameterValue.UinIsSubAccount"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) ApproveClientApply(request *ApproveClientApplyRequest) (response *ApproveClientApplyResponse, err error) {
    return c.ApproveClientApplyWithContext(context.Background(), request)
}

// ApproveClientApply
// Description: This API is used by resellers to review applications to become sub-customers. Note: This API is used to apply for the allowlist. If needed, please contact your business representative.
//
// 
//
// Callable roles: Reseller, Distributer, Second-level reseller
//
// error code that may be returned:
//  FAILEDOPERATION_BINDORGANIZATIONACCOUNT = "FailedOperation.BindOrganizationAccount"
//  FAILEDOPERATION_CLIENTBALANCEISLESSOREQUALZERO = "FailedOperation.ClientBalanceIsLessOrEqualZero"
//  FAILEDOPERATION_CLIENTBUYSP = "FailedOperation.ClientBuySP"
//  FAILEDOPERATION_CLIENTCREATESHAREUNIT = "FailedOperation.ClientCreateShareUnit"
//  FAILEDOPERATION_CLIENTJOINSHAREUNIT = "FailedOperation.ClientJoinShareUnit"
//  FAILEDOPERATION_CLIENTNOTAPPLY = "FailedOperation.ClientNotApply"
//  FAILEDOPERATION_EXCEEDMAXBINDCOUNT = "FailedOperation.ExceedMaxBindCount"
//  FAILEDOPERATION_UINALREADYKA = "FailedOperation.UinAlreadyKA"
//  FAILEDOPERATION_UINNOTRESELLER = "FailedOperation.UinNotReseller"
//  INVALIDPARAMETERVALUE_INVALIDUIN = "InvalidParameterValue.InvalidUin"
//  INVALIDPARAMETERVALUE_UINALREADYCLIENT = "InvalidParameterValue.UinAlreadyClient"
//  INVALIDPARAMETERVALUE_UINISSUBACCOUNT = "InvalidParameterValue.UinIsSubAccount"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) ApproveClientApplyWithContext(ctx context.Context, request *ApproveClientApplyRequest) (response *ApproveClientApplyResponse, err error) {
    if request == nil {
        request = NewApproveClientApplyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApproveClientApply require credential")
    }

    request.SetContext(ctx)
    
    response = NewApproveClientApplyResponse()
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
// This API is used to create Tencent Cloud customer accounts for distributor/second-level resellers.After the account is created, it will be automatically bound to the partner account.Note:
//
// 1. Create a Tencent Cloud account. The entered email address and mobile phone number need to be verified by the partner for validity.
//
// 2. Customers need to add personal information when logging in for the first time.
//
// 3. This interface needs to be applied for allowlist usage. Please contact the channel manager to initiate the application process.
//
// 
//
// Callable roles: distributor, second-level reseller, reseller
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_EXCEEDMAXBINDCOUNT = "FailedOperation.ExceedMaxBindCount"
//  FAILEDOPERATION_MAILISREGISTERED = "FailedOperation.MailIsRegistered"
//  FAILEDOPERATION_PHONEBINDUPPER = "FailedOperation.PhoneBindUpper"
//  FAILEDOPERATION_TRADEINFOINCORRECT = "FailedOperation.TradeInfoIncorrect"
//  FAILEDOPERATION_VERIFICATIONCODEILLEGAL = "FailedOperation.VerificationCodeIllegal"
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
//  INVALIDPARAMETERVALUE_TRADEINFOEMPTY = "InvalidParameterValue.TradeInfoEmpty"
//  INVALIDPARAMETERVALUE_UNSUPPORTAREA = "InvalidParameterValue.UnSupportArea"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAccount(request *CreateAccountRequest) (response *CreateAccountResponse, err error) {
    return c.CreateAccountWithContext(context.Background(), request)
}

// CreateAccount
// This API is used to create Tencent Cloud customer accounts for distributor/second-level resellers.After the account is created, it will be automatically bound to the partner account.Note:
//
// 1. Create a Tencent Cloud account. The entered email address and mobile phone number need to be verified by the partner for validity.
//
// 2. Customers need to add personal information when logging in for the first time.
//
// 3. This interface needs to be applied for allowlist usage. Please contact the channel manager to initiate the application process.
//
// 
//
// Callable roles: distributor, second-level reseller, reseller
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_EXCEEDMAXBINDCOUNT = "FailedOperation.ExceedMaxBindCount"
//  FAILEDOPERATION_MAILISREGISTERED = "FailedOperation.MailIsRegistered"
//  FAILEDOPERATION_PHONEBINDUPPER = "FailedOperation.PhoneBindUpper"
//  FAILEDOPERATION_TRADEINFOINCORRECT = "FailedOperation.TradeInfoIncorrect"
//  FAILEDOPERATION_VERIFICATIONCODEILLEGAL = "FailedOperation.VerificationCodeIllegal"
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
//  INVALIDPARAMETERVALUE_TRADEINFOEMPTY = "InvalidParameterValue.TradeInfoEmpty"
//  INVALIDPARAMETERVALUE_UNSUPPORTAREA = "InvalidParameterValue.UnSupportArea"
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

func NewCreateAndSendClientInvitationMailRequest() (request *CreateAndSendClientInvitationMailRequest) {
    request = &CreateAndSendClientInvitationMailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "CreateAndSendClientInvitationMail")
    
    
    return
}

func NewCreateAndSendClientInvitationMailResponse() (response *CreateAndSendClientInvitationMailResponse) {
    response = &CreateAndSendClientInvitationMailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAndSendClientInvitationMail
// This API is used to apply for the allowlist. If needed, please contact your business representative.Directions:
//
// 1.This API is used to create an invitation link, which you can send to a specified email address.
//
// 2.Customer need to click the invitation link in the email, fill in and submit the required information.
//
// 3.You can review the customer's application in customer management  after submission.
//
// 
//
// Note:This API is used to manually send the invitation link to the customer if the specified email does not receive it.
//
// error code that may be returned:
//  FAILEDOPERATION_SENDMAILLIMIT180 = "FailedOperation.SendMailLimit180"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) CreateAndSendClientInvitationMail(request *CreateAndSendClientInvitationMailRequest) (response *CreateAndSendClientInvitationMailResponse, err error) {
    return c.CreateAndSendClientInvitationMailWithContext(context.Background(), request)
}

// CreateAndSendClientInvitationMail
// This API is used to apply for the allowlist. If needed, please contact your business representative.Directions:
//
// 1.This API is used to create an invitation link, which you can send to a specified email address.
//
// 2.Customer need to click the invitation link in the email, fill in and submit the required information.
//
// 3.You can review the customer's application in customer management  after submission.
//
// 
//
// Note:This API is used to manually send the invitation link to the customer if the specified email does not receive it.
//
// error code that may be returned:
//  FAILEDOPERATION_SENDMAILLIMIT180 = "FailedOperation.SendMailLimit180"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) CreateAndSendClientInvitationMailWithContext(ctx context.Context, request *CreateAndSendClientInvitationMailRequest) (response *CreateAndSendClientInvitationMailResponse, err error) {
    if request == nil {
        request = NewCreateAndSendClientInvitationMailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAndSendClientInvitationMail require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAndSendClientInvitationMailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillDetailRequest() (request *DescribeBillDetailRequest) {
    request = &DescribeBillDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "DescribeBillDetail")
    
    
    return
}

func NewDescribeBillDetailResponse() (response *DescribeBillDetailResponse) {
    response = &DescribeBillDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillDetail
// This API is used to query bill details by customers.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDMONTH = "InvalidParameterValue.InvalidMonth"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) DescribeBillDetail(request *DescribeBillDetailRequest) (response *DescribeBillDetailResponse, err error) {
    return c.DescribeBillDetailWithContext(context.Background(), request)
}

// DescribeBillDetail
// This API is used to query bill details by customers.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDMONTH = "InvalidParameterValue.InvalidMonth"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) DescribeBillDetailWithContext(ctx context.Context, request *DescribeBillDetailRequest) (response *DescribeBillDetailResponse, err error) {
    if request == nil {
        request = NewDescribeBillDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillDownloadUrlRequest() (request *DescribeBillDownloadUrlRequest) {
    request = &DescribeBillDownloadUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "DescribeBillDownloadUrl")
    
    
    return
}

func NewDescribeBillDownloadUrlResponse() (response *DescribeBillDownloadUrlResponse) {
    response = &DescribeBillDownloadUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillDownloadUrl
// This API is used to download billing files and return billing file URLs by customers.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBillDownloadUrl(request *DescribeBillDownloadUrlRequest) (response *DescribeBillDownloadUrlResponse, err error) {
    return c.DescribeBillDownloadUrlWithContext(context.Background(), request)
}

// DescribeBillDownloadUrl
// This API is used to download billing files and return billing file URLs by customers.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBillDownloadUrlWithContext(ctx context.Context, request *DescribeBillDownloadUrlRequest) (response *DescribeBillDownloadUrlResponse, err error) {
    if request == nil {
        request = NewDescribeBillDownloadUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillDownloadUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillDownloadUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillSummaryRequest() (request *DescribeBillSummaryRequest) {
    request = &DescribeBillSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "DescribeBillSummary")
    
    
    return
}

func NewDescribeBillSummaryResponse() (response *DescribeBillSummaryResponse) {
    response = &DescribeBillSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillSummary
// External API for the L1 billing of the customer billing center
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDDIMENSION = "InvalidParameterValue.InvalidDimension"
//  INVALIDPARAMETERVALUE_INVALIDMONTH = "InvalidParameterValue.InvalidMonth"
func (c *Client) DescribeBillSummary(request *DescribeBillSummaryRequest) (response *DescribeBillSummaryResponse, err error) {
    return c.DescribeBillSummaryWithContext(context.Background(), request)
}

// DescribeBillSummary
// External API for the L1 billing of the customer billing center
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDDIMENSION = "InvalidParameterValue.InvalidDimension"
//  INVALIDPARAMETERVALUE_INVALIDMONTH = "InvalidParameterValue.InvalidMonth"
func (c *Client) DescribeBillSummaryWithContext(ctx context.Context, request *DescribeBillSummaryRequest) (response *DescribeBillSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeBillSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillSummaryByPayModeRequest() (request *DescribeBillSummaryByPayModeRequest) {
    request = &DescribeBillSummaryByPayModeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "DescribeBillSummaryByPayMode")
    
    
    return
}

func NewDescribeBillSummaryByPayModeResponse() (response *DescribeBillSummaryByPayModeResponse) {
    response = &DescribeBillSummaryByPayModeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillSummaryByPayMode
// This API is used to obtain the total amount of customer bills by payment mode.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDMONTH = "InvalidParameterValue.InvalidMonth"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) DescribeBillSummaryByPayMode(request *DescribeBillSummaryByPayModeRequest) (response *DescribeBillSummaryByPayModeResponse, err error) {
    return c.DescribeBillSummaryByPayModeWithContext(context.Background(), request)
}

// DescribeBillSummaryByPayMode
// This API is used to obtain the total amount of customer bills by payment mode.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDMONTH = "InvalidParameterValue.InvalidMonth"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) DescribeBillSummaryByPayModeWithContext(ctx context.Context, request *DescribeBillSummaryByPayModeRequest) (response *DescribeBillSummaryByPayModeResponse, err error) {
    if request == nil {
        request = NewDescribeBillSummaryByPayModeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillSummaryByPayMode require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillSummaryByPayModeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillSummaryByProductRequest() (request *DescribeBillSummaryByProductRequest) {
    request = &DescribeBillSummaryByProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "DescribeBillSummaryByProduct")
    
    
    return
}

func NewDescribeBillSummaryByProductResponse() (response *DescribeBillSummaryByProductResponse) {
    response = &DescribeBillSummaryByProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillSummaryByProduct
// This API is used to obtain the total amount of customer bills by product.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDMONTH = "InvalidParameterValue.InvalidMonth"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) DescribeBillSummaryByProduct(request *DescribeBillSummaryByProductRequest) (response *DescribeBillSummaryByProductResponse, err error) {
    return c.DescribeBillSummaryByProductWithContext(context.Background(), request)
}

// DescribeBillSummaryByProduct
// This API is used to obtain the total amount of customer bills by product.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDMONTH = "InvalidParameterValue.InvalidMonth"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) DescribeBillSummaryByProductWithContext(ctx context.Context, request *DescribeBillSummaryByProductRequest) (response *DescribeBillSummaryByProductResponse, err error) {
    if request == nil {
        request = NewDescribeBillSummaryByProductRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillSummaryByProduct require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillSummaryByProductResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillSummaryByRegionRequest() (request *DescribeBillSummaryByRegionRequest) {
    request = &DescribeBillSummaryByRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "DescribeBillSummaryByRegion")
    
    
    return
}

func NewDescribeBillSummaryByRegionResponse() (response *DescribeBillSummaryByRegionResponse) {
    response = &DescribeBillSummaryByRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillSummaryByRegion
// This API is used to obtain the total amount of customer bills by region.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDMONTH = "InvalidParameterValue.InvalidMonth"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) DescribeBillSummaryByRegion(request *DescribeBillSummaryByRegionRequest) (response *DescribeBillSummaryByRegionResponse, err error) {
    return c.DescribeBillSummaryByRegionWithContext(context.Background(), request)
}

// DescribeBillSummaryByRegion
// This API is used to obtain the total amount of customer bills by region.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDMONTH = "InvalidParameterValue.InvalidMonth"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) DescribeBillSummaryByRegionWithContext(ctx context.Context, request *DescribeBillSummaryByRegionRequest) (response *DescribeBillSummaryByRegionResponse, err error) {
    if request == nil {
        request = NewDescribeBillSummaryByRegionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillSummaryByRegion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillSummaryByRegionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomerBillDetailRequest() (request *DescribeCustomerBillDetailRequest) {
    request = &DescribeCustomerBillDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "DescribeCustomerBillDetail")
    
    
    return
}

func NewDescribeCustomerBillDetailResponse() (response *DescribeCustomerBillDetailResponse) {
    response = &DescribeCustomerBillDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCustomerBillDetail
// This API is used to query the customer bill details by resellers.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDMONTH = "InvalidParameterValue.InvalidMonth"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_NOTCUSTOMERUIN = "UnauthorizedOperation.NotCustomerUin"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) DescribeCustomerBillDetail(request *DescribeCustomerBillDetailRequest) (response *DescribeCustomerBillDetailResponse, err error) {
    return c.DescribeCustomerBillDetailWithContext(context.Background(), request)
}

// DescribeCustomerBillDetail
// This API is used to query the customer bill details by resellers.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDMONTH = "InvalidParameterValue.InvalidMonth"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_NOTCUSTOMERUIN = "UnauthorizedOperation.NotCustomerUin"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) DescribeCustomerBillDetailWithContext(ctx context.Context, request *DescribeCustomerBillDetailRequest) (response *DescribeCustomerBillDetailResponse, err error) {
    if request == nil {
        request = NewDescribeCustomerBillDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCustomerBillDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCustomerBillDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomerBillDownloadUrlRequest() (request *DescribeCustomerBillDownloadUrlRequest) {
    request = &DescribeCustomerBillDownloadUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "DescribeCustomerBillDownloadUrl")
    
    
    return
}

func NewDescribeCustomerBillDownloadUrlResponse() (response *DescribeCustomerBillDownloadUrlResponse) {
    response = &DescribeCustomerBillDownloadUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCustomerBillDownloadUrl
// This API is used to get the URL for downloading the customer bill file by reseller. The download conditions are as follows:
//
// 1. Detailed bills (billDetail and billDetailPack) can be downloaded starting from June 2022; resource bills (billResource and billResourcePack) can be downloaded starting from November 2023.
//
// 2. Bill packages (billDetailPack and billResourcePack) can only be downloaded after billing.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) DescribeCustomerBillDownloadUrl(request *DescribeCustomerBillDownloadUrlRequest) (response *DescribeCustomerBillDownloadUrlResponse, err error) {
    return c.DescribeCustomerBillDownloadUrlWithContext(context.Background(), request)
}

// DescribeCustomerBillDownloadUrl
// This API is used to get the URL for downloading the customer bill file by reseller. The download conditions are as follows:
//
// 1. Detailed bills (billDetail and billDetailPack) can be downloaded starting from June 2022; resource bills (billResource and billResourcePack) can be downloaded starting from November 2023.
//
// 2. Bill packages (billDetailPack and billResourcePack) can only be downloaded after billing.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) DescribeCustomerBillDownloadUrlWithContext(ctx context.Context, request *DescribeCustomerBillDownloadUrlRequest) (response *DescribeCustomerBillDownloadUrlResponse, err error) {
    if request == nil {
        request = NewDescribeCustomerBillDownloadUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCustomerBillDownloadUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCustomerBillDownloadUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomerBillSummaryRequest() (request *DescribeCustomerBillSummaryRequest) {
    request = &DescribeCustomerBillSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "DescribeCustomerBillSummary")
    
    
    return
}

func NewDescribeCustomerBillSummaryResponse() (response *DescribeCustomerBillSummaryResponse) {
    response = &DescribeCustomerBillSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCustomerBillSummary
// This API is used to query the total amount of customer bills.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDMONTH = "InvalidParameterValue.InvalidMonth"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) DescribeCustomerBillSummary(request *DescribeCustomerBillSummaryRequest) (response *DescribeCustomerBillSummaryResponse, err error) {
    return c.DescribeCustomerBillSummaryWithContext(context.Background(), request)
}

// DescribeCustomerBillSummary
// This API is used to query the total amount of customer bills.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDMONTH = "InvalidParameterValue.InvalidMonth"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) DescribeCustomerBillSummaryWithContext(ctx context.Context, request *DescribeCustomerBillSummaryRequest) (response *DescribeCustomerBillSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeCustomerBillSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCustomerBillSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCustomerBillSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomerInfoRequest() (request *DescribeCustomerInfoRequest) {
    request = &DescribeCustomerInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "DescribeCustomerInfo")
    
    
    return
}

func NewDescribeCustomerInfoResponse() (response *DescribeCustomerInfoResponse) {
    response = &DescribeCustomerInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCustomerInfo
// This API is used to query the customer information.
//
// error code that may be returned:
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) DescribeCustomerInfo(request *DescribeCustomerInfoRequest) (response *DescribeCustomerInfoResponse, err error) {
    return c.DescribeCustomerInfoWithContext(context.Background(), request)
}

// DescribeCustomerInfo
// This API is used to query the customer information.
//
// error code that may be returned:
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) DescribeCustomerInfoWithContext(ctx context.Context, request *DescribeCustomerInfoRequest) (response *DescribeCustomerInfoResponse, err error) {
    if request == nil {
        request = NewDescribeCustomerInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCustomerInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCustomerInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomerUinRequest() (request *DescribeCustomerUinRequest) {
    request = &DescribeCustomerUinRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "DescribeCustomerUin")
    
    
    return
}

func NewDescribeCustomerUinResponse() (response *DescribeCustomerUinResponse) {
    response = &DescribeCustomerUinResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCustomerUin
// This API is used to query the list of customer UINs.
//
// error code that may be returned:
//  INVALIDPARAMETER_PAGE = "InvalidParameter.Page"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) DescribeCustomerUin(request *DescribeCustomerUinRequest) (response *DescribeCustomerUinResponse, err error) {
    return c.DescribeCustomerUinWithContext(context.Background(), request)
}

// DescribeCustomerUin
// This API is used to query the list of customer UINs.
//
// error code that may be returned:
//  INVALIDPARAMETER_PAGE = "InvalidParameter.Page"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) DescribeCustomerUinWithContext(ctx context.Context, request *DescribeCustomerUinRequest) (response *DescribeCustomerUinResponse, err error) {
    if request == nil {
        request = NewDescribeCustomerUinRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCustomerUin require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCustomerUinResponse()
    err = c.Send(request, response)
    return
}

func NewForceQNRequest() (request *ForceQNRequest) {
    request = &ForceQNRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "ForceQN")
    
    
    return
}

func NewForceQNResponse() (response *ForceQNResponse) {
    response = &ForceQNResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ForceQN
// This API is used to set and cancel forced service suspension.
//
// Note:Reseller need to be allowlisted to use the API, please contact your business representative to apply for allowlist.
//
// error code that may be returned:
//  INVALIDPARAMETER_PAGE = "InvalidParameter.Page"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) ForceQN(request *ForceQNRequest) (response *ForceQNResponse, err error) {
    return c.ForceQNWithContext(context.Background(), request)
}

// ForceQN
// This API is used to set and cancel forced service suspension.
//
// Note:Reseller need to be allowlisted to use the API, please contact your business representative to apply for allowlist.
//
// error code that may be returned:
//  INVALIDPARAMETER_PAGE = "InvalidParameter.Page"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) ForceQNWithContext(ctx context.Context, request *ForceQNRequest) (response *ForceQNResponse, err error) {
    if request == nil {
        request = NewForceQNRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ForceQN require credential")
    }

    request.SetContext(ctx)
    
    response = NewForceQNResponse()
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

func NewGetTradeConfigListRequest() (request *GetTradeConfigListRequest) {
    request = &GetTradeConfigListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "GetTradeConfigList")
    
    
    return
}

func NewGetTradeConfigListResponse() (response *GetTradeConfigListResponse) {
    response = &GetTradeConfigListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetTradeConfigList
// This API is used to query industry information, including layer-1 industry and layer-2 industry.
//
// error code that may be returned:
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
func (c *Client) GetTradeConfigList(request *GetTradeConfigListRequest) (response *GetTradeConfigListResponse, err error) {
    return c.GetTradeConfigListWithContext(context.Background(), request)
}

// GetTradeConfigList
// This API is used to query industry information, including layer-1 industry and layer-2 industry.
//
// error code that may be returned:
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
func (c *Client) GetTradeConfigListWithContext(ctx context.Context, request *GetTradeConfigListRequest) (response *GetTradeConfigListResponse, err error) {
    if request == nil {
        request = NewGetTradeConfigListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTradeConfigList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTradeConfigListResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClientRemarkRequest() (request *ModifyClientRemarkRequest) {
    request = &ModifyClientRemarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "ModifyClientRemark")
    
    
    return
}

func NewModifyClientRemarkResponse() (response *ModifyClientRemarkResponse) {
    response = &ModifyClientRemarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyClientRemark
// This API is used to modify customer remarks.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDUIN = "InvalidParameterValue.InvalidUin"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) ModifyClientRemark(request *ModifyClientRemarkRequest) (response *ModifyClientRemarkResponse, err error) {
    return c.ModifyClientRemarkWithContext(context.Background(), request)
}

// ModifyClientRemark
// This API is used to modify customer remarks.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDUIN = "InvalidParameterValue.InvalidUin"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) ModifyClientRemarkWithContext(ctx context.Context, request *ModifyClientRemarkRequest) (response *ModifyClientRemarkResponse, err error) {
    if request == nil {
        request = NewModifyClientRemarkRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClientRemark require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClientRemarkResponse()
    err = c.Send(request, response)
    return
}

func NewQueryAccountVerificationStatusRequest() (request *QueryAccountVerificationStatusRequest) {
    request = &QueryAccountVerificationStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "QueryAccountVerificationStatus")
    
    
    return
}

func NewQueryAccountVerificationStatusResponse() (response *QueryAccountVerificationStatusResponse) {
    response = &QueryAccountVerificationStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryAccountVerificationStatus
// This API is used to query the account verification status.
//
// error code that may be returned:
//  FAILEDOPERATION_UININVALID = "FailedOperation.UinInvalid"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_NOTCUSTOMERUIN = "UnauthorizedOperation.NotCustomerUin"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) QueryAccountVerificationStatus(request *QueryAccountVerificationStatusRequest) (response *QueryAccountVerificationStatusResponse, err error) {
    return c.QueryAccountVerificationStatusWithContext(context.Background(), request)
}

// QueryAccountVerificationStatus
// This API is used to query the account verification status.
//
// error code that may be returned:
//  FAILEDOPERATION_UININVALID = "FailedOperation.UinInvalid"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_NOTCUSTOMERUIN = "UnauthorizedOperation.NotCustomerUin"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) QueryAccountVerificationStatusWithContext(ctx context.Context, request *QueryAccountVerificationStatusRequest) (response *QueryAccountVerificationStatusResponse, err error) {
    if request == nil {
        request = NewQueryAccountVerificationStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryAccountVerificationStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryAccountVerificationStatusResponse()
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
//  FAILEDOPERATION_UINNOTAGENT = "FailedOperation.UinNotAgent"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_UINLIST = "InvalidParameterValue.UinList"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION_NOTCUSTOMERUIN = "UnauthorizedOperation.NotCustomerUin"
func (c *Client) QueryCreditByUinList(request *QueryCreditByUinListRequest) (response *QueryCreditByUinListResponse, err error) {
    return c.QueryCreditByUinListWithContext(context.Background(), request)
}

// QueryCreditByUinList
// This API is used to query the credit of users in the list.
//
// error code that may be returned:
//  FAILEDOPERATION_UINNOTAGENT = "FailedOperation.UinNotAgent"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_UINLIST = "InvalidParameterValue.UinList"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION_NOTCUSTOMERUIN = "UnauthorizedOperation.NotCustomerUin"
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

func NewQueryCreditQuotaRequest() (request *QueryCreditQuotaRequest) {
    request = &QueryCreditQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "QueryCreditQuota")
    
    
    return
}

func NewQueryCreditQuotaResponse() (response *QueryCreditQuotaResponse) {
    response = &QueryCreditQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryCreditQuota
// This API is used to query customer credits.
//
// error code that may be returned:
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
func (c *Client) QueryCreditQuota(request *QueryCreditQuotaRequest) (response *QueryCreditQuotaResponse, err error) {
    return c.QueryCreditQuotaWithContext(context.Background(), request)
}

// QueryCreditQuota
// This API is used to query customer credits.
//
// error code that may be returned:
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
func (c *Client) QueryCreditQuotaWithContext(ctx context.Context, request *QueryCreditQuotaRequest) (response *QueryCreditQuotaResponse, err error) {
    if request == nil {
        request = NewQueryCreditQuotaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryCreditQuota require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryCreditQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewQueryCustomerBillingQuotaRequest() (request *QueryCustomerBillingQuotaRequest) {
    request = &QueryCustomerBillingQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "QueryCustomerBillingQuota")
    
    
    return
}

func NewQueryCustomerBillingQuotaResponse() (response *QueryCustomerBillingQuotaResponse) {
    response = &QueryCustomerBillingQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryCustomerBillingQuota
// Description: This API is used for a sub-customer to real-time query its own total credit and remaining credit in USD.
//
// 
//
// Callable roles: Sub-customer
//
// error code that may be returned:
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) QueryCustomerBillingQuota(request *QueryCustomerBillingQuotaRequest) (response *QueryCustomerBillingQuotaResponse, err error) {
    return c.QueryCustomerBillingQuotaWithContext(context.Background(), request)
}

// QueryCustomerBillingQuota
// Description: This API is used for a sub-customer to real-time query its own total credit and remaining credit in USD.
//
// 
//
// Callable roles: Sub-customer
//
// error code that may be returned:
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) QueryCustomerBillingQuotaWithContext(ctx context.Context, request *QueryCustomerBillingQuotaRequest) (response *QueryCustomerBillingQuotaResponse, err error) {
    if request == nil {
        request = NewQueryCustomerBillingQuotaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryCustomerBillingQuota require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryCustomerBillingQuotaResponse()
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
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) QueryCustomersCredit(request *QueryCustomersCreditRequest) (response *QueryCustomersCreditResponse, err error) {
    return c.QueryCustomersCreditWithContext(context.Background(), request)
}

// QueryCustomersCredit
// This API is used for a partner to the credits and basic information of cutomers.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
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
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
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
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
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
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) QueryPartnerCredit(request *QueryPartnerCreditRequest) (response *QueryPartnerCreditResponse, err error) {
    return c.QueryPartnerCreditWithContext(context.Background(), request)
}

// QueryPartnerCredit
// This API is used for a partner to query its own total credit, available credit, and used credit in USD.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
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

func NewQueryPendingClientsV2Request() (request *QueryPendingClientsV2Request) {
    request = &QueryPendingClientsV2Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "QueryPendingClientsV2")
    
    
    return
}

func NewQueryPendingClientsV2Response() (response *QueryPendingClientsV2Response) {
    response = &QueryPendingClientsV2Response{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryPendingClientsV2
// Description: This API is used by resellers to query the list of sub-customers pending review. Note: This API is used to apply for the allowlist. If needed, please contact your business representative.
//
// 
//
// Callable roles: Reseller, Distributer, Second-level reseller
//
// error code that may be returned:
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
func (c *Client) QueryPendingClientsV2(request *QueryPendingClientsV2Request) (response *QueryPendingClientsV2Response, err error) {
    return c.QueryPendingClientsV2WithContext(context.Background(), request)
}

// QueryPendingClientsV2
// Description: This API is used by resellers to query the list of sub-customers pending review. Note: This API is used to apply for the allowlist. If needed, please contact your business representative.
//
// 
//
// Callable roles: Reseller, Distributer, Second-level reseller
//
// error code that may be returned:
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
func (c *Client) QueryPendingClientsV2WithContext(ctx context.Context, request *QueryPendingClientsV2Request) (response *QueryPendingClientsV2Response, err error) {
    if request == nil {
        request = NewQueryPendingClientsV2Request()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryPendingClientsV2 require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryPendingClientsV2Response()
    err = c.Send(request, response)
    return
}

func NewQueryPolicyProductListByCodeRequest() (request *QueryPolicyProductListByCodeRequest) {
    request = &QueryPolicyProductListByCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "QueryPolicyProductListByCode")
    
    
    return
}

func NewQueryPolicyProductListByCodeResponse() (response *QueryPolicyProductListByCodeResponse) {
    response = &QueryPolicyProductListByCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryPolicyProductListByCode
// This API is used to query the product list information within the specified policy range. To call this API, contact your account manager to add it to the allowlist.
//
// error code that may be returned:
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
func (c *Client) QueryPolicyProductListByCode(request *QueryPolicyProductListByCodeRequest) (response *QueryPolicyProductListByCodeResponse, err error) {
    return c.QueryPolicyProductListByCodeWithContext(context.Background(), request)
}

// QueryPolicyProductListByCode
// This API is used to query the product list information within the specified policy range. To call this API, contact your account manager to add it to the allowlist.
//
// error code that may be returned:
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
func (c *Client) QueryPolicyProductListByCodeWithContext(ctx context.Context, request *QueryPolicyProductListByCodeRequest) (response *QueryPolicyProductListByCodeResponse, err error) {
    if request == nil {
        request = NewQueryPolicyProductListByCodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryPolicyProductListByCode require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryPolicyProductListByCodeResponse()
    err = c.Send(request, response)
    return
}

func NewQueryVoucherAmountByUinRequest() (request *QueryVoucherAmountByUinRequest) {
    request = &QueryVoucherAmountByUinRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "QueryVoucherAmountByUin")
    
    
    return
}

func NewQueryVoucherAmountByUinResponse() (response *QueryVoucherAmountByUinResponse) {
    response = &QueryVoucherAmountByUinResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryVoucherAmountByUin
// This API is used to query the voucher quota based on the customer UIN.
//
// error code that may be returned:
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
func (c *Client) QueryVoucherAmountByUin(request *QueryVoucherAmountByUinRequest) (response *QueryVoucherAmountByUinResponse, err error) {
    return c.QueryVoucherAmountByUinWithContext(context.Background(), request)
}

// QueryVoucherAmountByUin
// This API is used to query the voucher quota based on the customer UIN.
//
// error code that may be returned:
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
func (c *Client) QueryVoucherAmountByUinWithContext(ctx context.Context, request *QueryVoucherAmountByUinRequest) (response *QueryVoucherAmountByUinResponse, err error) {
    if request == nil {
        request = NewQueryVoucherAmountByUinRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryVoucherAmountByUin require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryVoucherAmountByUinResponse()
    err = c.Send(request, response)
    return
}

func NewQueryVoucherListByUinRequest() (request *QueryVoucherListByUinRequest) {
    request = &QueryVoucherListByUinRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "QueryVoucherListByUin")
    
    
    return
}

func NewQueryVoucherListByUinResponse() (response *QueryVoucherListByUinResponse) {
    response = &QueryVoucherListByUinResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryVoucherListByUin
// This API is used to query the voucher list based on the customer UIN.
//
// error code that may be returned:
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
func (c *Client) QueryVoucherListByUin(request *QueryVoucherListByUinRequest) (response *QueryVoucherListByUinResponse, err error) {
    return c.QueryVoucherListByUinWithContext(context.Background(), request)
}

// QueryVoucherListByUin
// This API is used to query the voucher list based on the customer UIN.
//
// error code that may be returned:
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
func (c *Client) QueryVoucherListByUinWithContext(ctx context.Context, request *QueryVoucherListByUinRequest) (response *QueryVoucherListByUinResponse, err error) {
    if request == nil {
        request = NewQueryVoucherListByUinRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryVoucherListByUin require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryVoucherListByUinResponse()
    err = c.Send(request, response)
    return
}

func NewQueryVoucherPoolRequest() (request *QueryVoucherPoolRequest) {
    request = &QueryVoucherPoolRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "QueryVoucherPool")
    
    
    return
}

func NewQueryVoucherPoolResponse() (response *QueryVoucherPoolResponse) {
    response = &QueryVoucherPoolResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryVoucherPool
// This API is used to query the voucher quota pool.
//
// error code that may be returned:
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
func (c *Client) QueryVoucherPool(request *QueryVoucherPoolRequest) (response *QueryVoucherPoolResponse, err error) {
    return c.QueryVoucherPoolWithContext(context.Background(), request)
}

// QueryVoucherPool
// This API is used to query the voucher quota pool.
//
// error code that may be returned:
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
func (c *Client) QueryVoucherPoolWithContext(ctx context.Context, request *QueryVoucherPoolRequest) (response *QueryVoucherPoolResponse, err error) {
    if request == nil {
        request = NewQueryVoucherPoolRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryVoucherPool require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryVoucherPoolResponse()
    err = c.Send(request, response)
    return
}

func NewSendVerifyCodeRequest() (request *SendVerifyCodeRequest) {
    request = &SendVerifyCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "SendVerifyCode")
    
    
    return
}

func NewSendVerifyCodeResponse() (response *SendVerifyCodeResponse) {
    response = &SendVerifyCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SendVerifyCode
// This API is used to send a verification code for account registration.
//
// error code that may be returned:
//  FAILEDOPERATION_PHONEBINDUPPER = "FailedOperation.PhoneBindUpper"
//  FAILEDOPERATION_SENDVERIFYCODELIMIT = "FailedOperation.SendVerifyCodeLimit"
//  FAILEDOPERATION_SENDVERIFYCODELIMIT60 = "FailedOperation.SendVerifyCodeLimit60"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_UNSUPPORTAREA = "InvalidParameterValue.UnSupportArea"
func (c *Client) SendVerifyCode(request *SendVerifyCodeRequest) (response *SendVerifyCodeResponse, err error) {
    return c.SendVerifyCodeWithContext(context.Background(), request)
}

// SendVerifyCode
// This API is used to send a verification code for account registration.
//
// error code that may be returned:
//  FAILEDOPERATION_PHONEBINDUPPER = "FailedOperation.PhoneBindUpper"
//  FAILEDOPERATION_SENDVERIFYCODELIMIT = "FailedOperation.SendVerifyCodeLimit"
//  FAILEDOPERATION_SENDVERIFYCODELIMIT60 = "FailedOperation.SendVerifyCodeLimit60"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_UNSUPPORTAREA = "InvalidParameterValue.UnSupportArea"
func (c *Client) SendVerifyCodeWithContext(ctx context.Context, request *SendVerifyCodeRequest) (response *SendVerifyCodeResponse, err error) {
    if request == nil {
        request = NewSendVerifyCodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SendVerifyCode require credential")
    }

    request.SetContext(ctx)
    
    response = NewSendVerifyCodeResponse()
    err = c.Send(request, response)
    return
}
