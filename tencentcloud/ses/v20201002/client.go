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

package v20201002

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2020-10-02"

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


func NewBatchSendEmailRequest() (request *BatchSendEmailRequest) {
    request = &BatchSendEmailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ses", APIVersion, "BatchSendEmail")
    
    
    return
}

func NewBatchSendEmailResponse() (response *BatchSendEmailResponse) {
    response = &BatchSendEmailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchSendEmail
// This API is used to send a TEXT or HTML email to multiple recipients at a time for marketing or notification purposes. By default, you can send emails using a template only. You need to create a recipient group with email addresses first and then send emails by group ID. SES supports scheduled and recurring email sending tasks. You need to pass in `TimedParam` for a scheduled task and `CycleParam` for a recurring one.
//
// error code that may be returned:
//  FAILEDOPERATION_EMAILCONTENTTOOLARGE = "FailedOperation.EmailContentToolarge"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_BEGINTIMEBEFORENOW = "InvalidParameterValue.BeginTimeBeforeNow"
//  INVALIDPARAMETERVALUE_EMAILCONTENTISWRONG = "InvalidParameterValue.EmailContentIsWrong"
//  INVALIDPARAMETERVALUE_SUBJECTLENGTHERROR = "InvalidParameterValue.SubjectLengthError"
//  INVALIDPARAMETERVALUE_TEMPLATEDATAERROR = "InvalidParameterValue.TemplateDataError"
//  INVALIDPARAMETERVALUE_TEMPLATENOTMATCHDATA = "InvalidParameterValue.TemplateNotMatchData"
//  MISSINGPARAMETER_CYCLEPARAMNECESSARY = "MissingParameter.CycleParamNecessary"
//  MISSINGPARAMETER_SENDPARAMNECESSARY = "MissingParameter.SendParamNecessary"
//  MISSINGPARAMETER_TIMEDPARAMNECESSARY = "MissingParameter.TimedParamNecessary"
//  OPERATIONDENIED_RECEIVERNOTEXIST = "OperationDenied.ReceiverNotExist"
//  OPERATIONDENIED_RECEIVERSTATUSERROR = "OperationDenied.ReceiverStatusError"
//  OPERATIONDENIED_SENDADDRESSSTATUSERROR = "OperationDenied.SendAddressStatusError"
//  OPERATIONDENIED_TEMPLATESTATUSERROR = "OperationDenied.TemplateStatusError"
func (c *Client) BatchSendEmail(request *BatchSendEmailRequest) (response *BatchSendEmailResponse, err error) {
    return c.BatchSendEmailWithContext(context.Background(), request)
}

// BatchSendEmail
// This API is used to send a TEXT or HTML email to multiple recipients at a time for marketing or notification purposes. By default, you can send emails using a template only. You need to create a recipient group with email addresses first and then send emails by group ID. SES supports scheduled and recurring email sending tasks. You need to pass in `TimedParam` for a scheduled task and `CycleParam` for a recurring one.
//
// error code that may be returned:
//  FAILEDOPERATION_EMAILCONTENTTOOLARGE = "FailedOperation.EmailContentToolarge"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_BEGINTIMEBEFORENOW = "InvalidParameterValue.BeginTimeBeforeNow"
//  INVALIDPARAMETERVALUE_EMAILCONTENTISWRONG = "InvalidParameterValue.EmailContentIsWrong"
//  INVALIDPARAMETERVALUE_SUBJECTLENGTHERROR = "InvalidParameterValue.SubjectLengthError"
//  INVALIDPARAMETERVALUE_TEMPLATEDATAERROR = "InvalidParameterValue.TemplateDataError"
//  INVALIDPARAMETERVALUE_TEMPLATENOTMATCHDATA = "InvalidParameterValue.TemplateNotMatchData"
//  MISSINGPARAMETER_CYCLEPARAMNECESSARY = "MissingParameter.CycleParamNecessary"
//  MISSINGPARAMETER_SENDPARAMNECESSARY = "MissingParameter.SendParamNecessary"
//  MISSINGPARAMETER_TIMEDPARAMNECESSARY = "MissingParameter.TimedParamNecessary"
//  OPERATIONDENIED_RECEIVERNOTEXIST = "OperationDenied.ReceiverNotExist"
//  OPERATIONDENIED_RECEIVERSTATUSERROR = "OperationDenied.ReceiverStatusError"
//  OPERATIONDENIED_SENDADDRESSSTATUSERROR = "OperationDenied.SendAddressStatusError"
//  OPERATIONDENIED_TEMPLATESTATUSERROR = "OperationDenied.TemplateStatusError"
func (c *Client) BatchSendEmailWithContext(ctx context.Context, request *BatchSendEmailRequest) (response *BatchSendEmailResponse, err error) {
    if request == nil {
        request = NewBatchSendEmailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ses", APIVersion, "BatchSendEmail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchSendEmail require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchSendEmailResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAddressUnsubscribeConfigRequest() (request *CreateAddressUnsubscribeConfigRequest) {
    request = &CreateAddressUnsubscribeConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ses", APIVersion, "CreateAddressUnsubscribeConfig")
    
    
    return
}

func NewCreateAddressUnsubscribeConfigResponse() (response *CreateAddressUnsubscribeConfigResponse) {
    response = &CreateAddressUnsubscribeConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAddressUnsubscribeConfig
// This API is used to create an address-level unsubscribe configuration.
//
// error code that may be returned:
//  FAILEDOPERATION_EMAILCONTENTTOOLARGE = "FailedOperation.EmailContentToolarge"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_BEGINTIMEBEFORENOW = "InvalidParameterValue.BeginTimeBeforeNow"
//  INVALIDPARAMETERVALUE_EMAILCONTENTISWRONG = "InvalidParameterValue.EmailContentIsWrong"
//  INVALIDPARAMETERVALUE_SUBJECTLENGTHERROR = "InvalidParameterValue.SubjectLengthError"
//  INVALIDPARAMETERVALUE_TEMPLATEDATAERROR = "InvalidParameterValue.TemplateDataError"
//  INVALIDPARAMETERVALUE_TEMPLATENOTMATCHDATA = "InvalidParameterValue.TemplateNotMatchData"
//  MISSINGPARAMETER_CYCLEPARAMNECESSARY = "MissingParameter.CycleParamNecessary"
//  MISSINGPARAMETER_SENDPARAMNECESSARY = "MissingParameter.SendParamNecessary"
//  MISSINGPARAMETER_TIMEDPARAMNECESSARY = "MissingParameter.TimedParamNecessary"
//  OPERATIONDENIED_RECEIVERNOTEXIST = "OperationDenied.ReceiverNotExist"
//  OPERATIONDENIED_RECEIVERSTATUSERROR = "OperationDenied.ReceiverStatusError"
//  OPERATIONDENIED_SENDADDRESSSTATUSERROR = "OperationDenied.SendAddressStatusError"
//  OPERATIONDENIED_TEMPLATESTATUSERROR = "OperationDenied.TemplateStatusError"
func (c *Client) CreateAddressUnsubscribeConfig(request *CreateAddressUnsubscribeConfigRequest) (response *CreateAddressUnsubscribeConfigResponse, err error) {
    return c.CreateAddressUnsubscribeConfigWithContext(context.Background(), request)
}

// CreateAddressUnsubscribeConfig
// This API is used to create an address-level unsubscribe configuration.
//
// error code that may be returned:
//  FAILEDOPERATION_EMAILCONTENTTOOLARGE = "FailedOperation.EmailContentToolarge"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_BEGINTIMEBEFORENOW = "InvalidParameterValue.BeginTimeBeforeNow"
//  INVALIDPARAMETERVALUE_EMAILCONTENTISWRONG = "InvalidParameterValue.EmailContentIsWrong"
//  INVALIDPARAMETERVALUE_SUBJECTLENGTHERROR = "InvalidParameterValue.SubjectLengthError"
//  INVALIDPARAMETERVALUE_TEMPLATEDATAERROR = "InvalidParameterValue.TemplateDataError"
//  INVALIDPARAMETERVALUE_TEMPLATENOTMATCHDATA = "InvalidParameterValue.TemplateNotMatchData"
//  MISSINGPARAMETER_CYCLEPARAMNECESSARY = "MissingParameter.CycleParamNecessary"
//  MISSINGPARAMETER_SENDPARAMNECESSARY = "MissingParameter.SendParamNecessary"
//  MISSINGPARAMETER_TIMEDPARAMNECESSARY = "MissingParameter.TimedParamNecessary"
//  OPERATIONDENIED_RECEIVERNOTEXIST = "OperationDenied.ReceiverNotExist"
//  OPERATIONDENIED_RECEIVERSTATUSERROR = "OperationDenied.ReceiverStatusError"
//  OPERATIONDENIED_SENDADDRESSSTATUSERROR = "OperationDenied.SendAddressStatusError"
//  OPERATIONDENIED_TEMPLATESTATUSERROR = "OperationDenied.TemplateStatusError"
func (c *Client) CreateAddressUnsubscribeConfigWithContext(ctx context.Context, request *CreateAddressUnsubscribeConfigRequest) (response *CreateAddressUnsubscribeConfigResponse, err error) {
    if request == nil {
        request = NewCreateAddressUnsubscribeConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ses", APIVersion, "CreateAddressUnsubscribeConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAddressUnsubscribeConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAddressUnsubscribeConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEmailAddressRequest() (request *CreateEmailAddressRequest) {
    request = &CreateEmailAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ses", APIVersion, "CreateEmailAddress")
    
    
    return
}

func NewCreateEmailAddressResponse() (response *CreateEmailAddressResponse) {
    response = &CreateEmailAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateEmailAddress
// After the sender domain is verified, you need a sender address to send emails. For example, if your sender domain is mail.qcloud.com, your sender address can be service@mail.qcloud.com. If you want to display your name (such as "Tencent Cloud") in the inbox list of the recipients, the sender address should be in the format of `Tencent Cloud <email address>`. Please note that there must be a space between your name and the first angle bracket.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALEMAILADDRESS = "InvalidParameterValue.IllegalEmailAddress"
//  INVALIDPARAMETERVALUE_ILLEGALSENDERNAME = "InvalidParameterValue.IllegalSenderName"
//  INVALIDPARAMETERVALUE_REPEATEMAILADDRESS = "InvalidParameterValue.RepeatEmailAddress"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DOMAINNOTVERIFIED = "OperationDenied.DomainNotVerified"
//  OPERATIONDENIED_EXCEEDSENDERLIMIT = "OperationDenied.ExceedSenderLimit"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) CreateEmailAddress(request *CreateEmailAddressRequest) (response *CreateEmailAddressResponse, err error) {
    return c.CreateEmailAddressWithContext(context.Background(), request)
}

// CreateEmailAddress
// After the sender domain is verified, you need a sender address to send emails. For example, if your sender domain is mail.qcloud.com, your sender address can be service@mail.qcloud.com. If you want to display your name (such as "Tencent Cloud") in the inbox list of the recipients, the sender address should be in the format of `Tencent Cloud <email address>`. Please note that there must be a space between your name and the first angle bracket.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALEMAILADDRESS = "InvalidParameterValue.IllegalEmailAddress"
//  INVALIDPARAMETERVALUE_ILLEGALSENDERNAME = "InvalidParameterValue.IllegalSenderName"
//  INVALIDPARAMETERVALUE_REPEATEMAILADDRESS = "InvalidParameterValue.RepeatEmailAddress"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DOMAINNOTVERIFIED = "OperationDenied.DomainNotVerified"
//  OPERATIONDENIED_EXCEEDSENDERLIMIT = "OperationDenied.ExceedSenderLimit"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) CreateEmailAddressWithContext(ctx context.Context, request *CreateEmailAddressRequest) (response *CreateEmailAddressResponse, err error) {
    if request == nil {
        request = NewCreateEmailAddressRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ses", APIVersion, "CreateEmailAddress")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEmailAddress require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEmailAddressResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEmailIdentityRequest() (request *CreateEmailIdentityRequest) {
    request = &CreateEmailIdentityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ses", APIVersion, "CreateEmailIdentity")
    
    
    return
}

func NewCreateEmailIdentityResponse() (response *CreateEmailIdentityResponse) {
    response = &CreateEmailIdentityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateEmailIdentity
// This API is used to create a sender domain. Before you can send an email using Tencent Cloud SES, you must create a sender domain as your identity. It can be the domain of your website or mobile app. You must verify the domain to prove that you own it and authorize Tencent Cloud SES to use it to send emails.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CREATEDBYOTHER = "InvalidParameterValue.CreatedByOther"
//  INVALIDPARAMETERVALUE_INVALIDEMAILIDENTITY = "InvalidParameterValue.InvalidEmailIdentity"
//  INVALIDPARAMETERVALUE_REPEATCREATION = "InvalidParameterValue.RepeatCreation"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_EXCEEDDOMAINLIMIT = "OperationDenied.ExceedDomainLimit"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) CreateEmailIdentity(request *CreateEmailIdentityRequest) (response *CreateEmailIdentityResponse, err error) {
    return c.CreateEmailIdentityWithContext(context.Background(), request)
}

// CreateEmailIdentity
// This API is used to create a sender domain. Before you can send an email using Tencent Cloud SES, you must create a sender domain as your identity. It can be the domain of your website or mobile app. You must verify the domain to prove that you own it and authorize Tencent Cloud SES to use it to send emails.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CREATEDBYOTHER = "InvalidParameterValue.CreatedByOther"
//  INVALIDPARAMETERVALUE_INVALIDEMAILIDENTITY = "InvalidParameterValue.InvalidEmailIdentity"
//  INVALIDPARAMETERVALUE_REPEATCREATION = "InvalidParameterValue.RepeatCreation"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_EXCEEDDOMAINLIMIT = "OperationDenied.ExceedDomainLimit"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) CreateEmailIdentityWithContext(ctx context.Context, request *CreateEmailIdentityRequest) (response *CreateEmailIdentityResponse, err error) {
    if request == nil {
        request = NewCreateEmailIdentityRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ses", APIVersion, "CreateEmailIdentity")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEmailIdentity require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEmailIdentityResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEmailTemplateRequest() (request *CreateEmailTemplateRequest) {
    request = &CreateEmailTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ses", APIVersion, "CreateEmailTemplate")
    
    
    return
}

func NewCreateEmailTemplateResponse() (response *CreateEmailTemplateResponse) {
    response = &CreateEmailTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateEmailTemplate
// This API is used to create a TEXT or HTML email template. To create an HTML template, ensure that it does not include external CSS files. You can use {{variable name}} to specify a variable in the template.
//
// Note: Only an approved template can be used to send emails.
//
// error code that may be returned:
//  FAILEDOPERATION_EXCEEDTEMPLATELIMIT = "FailedOperation.ExceedTemplateLimit"
//  FAILEDOPERATION_TEMPLATECONTENTTOOLARGE = "FailedOperation.TemplateContentToolarge"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TEMPLATECONTENTISNULL = "InvalidParameterValue.TemplateContentIsNULL"
//  INVALIDPARAMETERVALUE_TEMPLATECONTENTISWRONG = "InvalidParameterValue.TemplateContentIsWrong"
//  INVALIDPARAMETERVALUE_TEMPLATENAMEILLEGAL = "InvalidParameterValue.TemplateNameIllegal"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) CreateEmailTemplate(request *CreateEmailTemplateRequest) (response *CreateEmailTemplateResponse, err error) {
    return c.CreateEmailTemplateWithContext(context.Background(), request)
}

// CreateEmailTemplate
// This API is used to create a TEXT or HTML email template. To create an HTML template, ensure that it does not include external CSS files. You can use {{variable name}} to specify a variable in the template.
//
// Note: Only an approved template can be used to send emails.
//
// error code that may be returned:
//  FAILEDOPERATION_EXCEEDTEMPLATELIMIT = "FailedOperation.ExceedTemplateLimit"
//  FAILEDOPERATION_TEMPLATECONTENTTOOLARGE = "FailedOperation.TemplateContentToolarge"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TEMPLATECONTENTISNULL = "InvalidParameterValue.TemplateContentIsNULL"
//  INVALIDPARAMETERVALUE_TEMPLATECONTENTISWRONG = "InvalidParameterValue.TemplateContentIsWrong"
//  INVALIDPARAMETERVALUE_TEMPLATENAMEILLEGAL = "InvalidParameterValue.TemplateNameIllegal"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) CreateEmailTemplateWithContext(ctx context.Context, request *CreateEmailTemplateRequest) (response *CreateEmailTemplateResponse, err error) {
    if request == nil {
        request = NewCreateEmailTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ses", APIVersion, "CreateEmailTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEmailTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEmailTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateReceiverRequest() (request *CreateReceiverRequest) {
    request = &CreateReceiverRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ses", APIVersion, "CreateReceiver")
    
    
    return
}

func NewCreateReceiverResponse() (response *CreateReceiverResponse) {
    response = &CreateReceiverResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateReceiver
// This API is used to create a recipient group, which is the list of target email addresses for batch sending emails. After creating a group, you need to upload recipient email addresses. Then, you can create a sending task and select the group to batch send emails.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_RECEIVERDESCILLEGAL = "InvalidParameterValue.ReceiverDescIllegal"
//  INVALIDPARAMETERVALUE_RECEIVERNAMEILLEGAL = "InvalidParameterValue.ReceiverNameIllegal"
//  INVALIDPARAMETERVALUE_REPEATRECEIVERNAME = "InvalidParameterValue.RepeatReceiverName"
//  LIMITEXCEEDED_EXCEEDRECEIVERLIMIT = "LimitExceeded.ExceedReceiverLimit"
func (c *Client) CreateReceiver(request *CreateReceiverRequest) (response *CreateReceiverResponse, err error) {
    return c.CreateReceiverWithContext(context.Background(), request)
}

// CreateReceiver
// This API is used to create a recipient group, which is the list of target email addresses for batch sending emails. After creating a group, you need to upload recipient email addresses. Then, you can create a sending task and select the group to batch send emails.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_RECEIVERDESCILLEGAL = "InvalidParameterValue.ReceiverDescIllegal"
//  INVALIDPARAMETERVALUE_RECEIVERNAMEILLEGAL = "InvalidParameterValue.ReceiverNameIllegal"
//  INVALIDPARAMETERVALUE_REPEATRECEIVERNAME = "InvalidParameterValue.RepeatReceiverName"
//  LIMITEXCEEDED_EXCEEDRECEIVERLIMIT = "LimitExceeded.ExceedReceiverLimit"
func (c *Client) CreateReceiverWithContext(ctx context.Context, request *CreateReceiverRequest) (response *CreateReceiverResponse, err error) {
    if request == nil {
        request = NewCreateReceiverRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ses", APIVersion, "CreateReceiver")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateReceiver require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateReceiverResponse()
    err = c.Send(request, response)
    return
}

func NewCreateReceiverDetailRequest() (request *CreateReceiverDetailRequest) {
    request = &CreateReceiverDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ses", APIVersion, "CreateReceiverDetail")
    
    
    return
}

func NewCreateReceiverDetailResponse() (response *CreateReceiverDetailResponse) {
    response = &CreateReceiverDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateReceiverDetail
// This API is used to add recipient email addresses (up to 20,000 at a time) to a recipient group. This will be processed asynchronously. If the data volume is large, it may take some time to upload. You can check the recipient group for the upload status and upload quantity. This API has basically the same feature as that of `CreateReceiverDetailWithData` except that it doesn't support uploading template parameters for email sending. You need to first call the `CreateReceiver` API to create a recipient group, then call this API to pass in recipient addresses, and finally call the `BatchSendEmail` API to batch send emails. This API supports adding more recipient addresses during upload but not address deduplication, so you need to make sure that the recipient addresses are not duplicate by yourself. This API can request up to 20,000 recipient addresses at a time, but the recipient group can contain up to 50,000 addresses currently.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICENOTAVAILABLE = "FailedOperation.ServiceNotAvailable"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_TEMPLATEDATAERROR = "InvalidParameterValue.TemplateDataError"
//  INVALIDPARAMETERVALUE_TEMPLATEDATAINCONSISTENT = "InvalidParameterValue.TemplateDataInconsistent"
//  LIMITEXCEEDED_RECEIVERDETAILCOUNTLIMIT = "LimitExceeded.ReceiverDetailCountLimit"
//  LIMITEXCEEDED_RECEIVERDETAILREQUESTLIMIT = "LimitExceeded.ReceiverDetailRequestLimit"
//  MISSINGPARAMETER_EMAILSNECESSARY = "MissingParameter.EmailsNecessary"
//  MISSINGPARAMETER_RECEIVERIDNECESSARY = "MissingParameter.ReceiverIdNecessary"
//  OPERATIONDENIED_RECEIVERISOPERATING = "OperationDenied.ReceiverIsOperating"
//  OPERATIONDENIED_RECEIVERNOTEXIST = "OperationDenied.ReceiverNotExist"
func (c *Client) CreateReceiverDetail(request *CreateReceiverDetailRequest) (response *CreateReceiverDetailResponse, err error) {
    return c.CreateReceiverDetailWithContext(context.Background(), request)
}

// CreateReceiverDetail
// This API is used to add recipient email addresses (up to 20,000 at a time) to a recipient group. This will be processed asynchronously. If the data volume is large, it may take some time to upload. You can check the recipient group for the upload status and upload quantity. This API has basically the same feature as that of `CreateReceiverDetailWithData` except that it doesn't support uploading template parameters for email sending. You need to first call the `CreateReceiver` API to create a recipient group, then call this API to pass in recipient addresses, and finally call the `BatchSendEmail` API to batch send emails. This API supports adding more recipient addresses during upload but not address deduplication, so you need to make sure that the recipient addresses are not duplicate by yourself. This API can request up to 20,000 recipient addresses at a time, but the recipient group can contain up to 50,000 addresses currently.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICENOTAVAILABLE = "FailedOperation.ServiceNotAvailable"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_TEMPLATEDATAERROR = "InvalidParameterValue.TemplateDataError"
//  INVALIDPARAMETERVALUE_TEMPLATEDATAINCONSISTENT = "InvalidParameterValue.TemplateDataInconsistent"
//  LIMITEXCEEDED_RECEIVERDETAILCOUNTLIMIT = "LimitExceeded.ReceiverDetailCountLimit"
//  LIMITEXCEEDED_RECEIVERDETAILREQUESTLIMIT = "LimitExceeded.ReceiverDetailRequestLimit"
//  MISSINGPARAMETER_EMAILSNECESSARY = "MissingParameter.EmailsNecessary"
//  MISSINGPARAMETER_RECEIVERIDNECESSARY = "MissingParameter.ReceiverIdNecessary"
//  OPERATIONDENIED_RECEIVERISOPERATING = "OperationDenied.ReceiverIsOperating"
//  OPERATIONDENIED_RECEIVERNOTEXIST = "OperationDenied.ReceiverNotExist"
func (c *Client) CreateReceiverDetailWithContext(ctx context.Context, request *CreateReceiverDetailRequest) (response *CreateReceiverDetailResponse, err error) {
    if request == nil {
        request = NewCreateReceiverDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ses", APIVersion, "CreateReceiverDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateReceiverDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateReceiverDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAddressUnsubscribeConfigRequest() (request *DeleteAddressUnsubscribeConfigRequest) {
    request = &DeleteAddressUnsubscribeConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ses", APIVersion, "DeleteAddressUnsubscribeConfig")
    
    
    return
}

func NewDeleteAddressUnsubscribeConfigResponse() (response *DeleteAddressUnsubscribeConfigResponse) {
    response = &DeleteAddressUnsubscribeConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAddressUnsubscribeConfig
// Remove address-level unsubscribe configuration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICENOTAVAILABLE = "FailedOperation.ServiceNotAvailable"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_TEMPLATEDATAERROR = "InvalidParameterValue.TemplateDataError"
//  INVALIDPARAMETERVALUE_TEMPLATEDATAINCONSISTENT = "InvalidParameterValue.TemplateDataInconsistent"
//  LIMITEXCEEDED_RECEIVERDETAILCOUNTLIMIT = "LimitExceeded.ReceiverDetailCountLimit"
//  LIMITEXCEEDED_RECEIVERDETAILREQUESTLIMIT = "LimitExceeded.ReceiverDetailRequestLimit"
//  MISSINGPARAMETER_EMAILSNECESSARY = "MissingParameter.EmailsNecessary"
//  MISSINGPARAMETER_RECEIVERIDNECESSARY = "MissingParameter.ReceiverIdNecessary"
//  OPERATIONDENIED_RECEIVERISOPERATING = "OperationDenied.ReceiverIsOperating"
//  OPERATIONDENIED_RECEIVERNOTEXIST = "OperationDenied.ReceiverNotExist"
func (c *Client) DeleteAddressUnsubscribeConfig(request *DeleteAddressUnsubscribeConfigRequest) (response *DeleteAddressUnsubscribeConfigResponse, err error) {
    return c.DeleteAddressUnsubscribeConfigWithContext(context.Background(), request)
}

// DeleteAddressUnsubscribeConfig
// Remove address-level unsubscribe configuration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICENOTAVAILABLE = "FailedOperation.ServiceNotAvailable"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_TEMPLATEDATAERROR = "InvalidParameterValue.TemplateDataError"
//  INVALIDPARAMETERVALUE_TEMPLATEDATAINCONSISTENT = "InvalidParameterValue.TemplateDataInconsistent"
//  LIMITEXCEEDED_RECEIVERDETAILCOUNTLIMIT = "LimitExceeded.ReceiverDetailCountLimit"
//  LIMITEXCEEDED_RECEIVERDETAILREQUESTLIMIT = "LimitExceeded.ReceiverDetailRequestLimit"
//  MISSINGPARAMETER_EMAILSNECESSARY = "MissingParameter.EmailsNecessary"
//  MISSINGPARAMETER_RECEIVERIDNECESSARY = "MissingParameter.ReceiverIdNecessary"
//  OPERATIONDENIED_RECEIVERISOPERATING = "OperationDenied.ReceiverIsOperating"
//  OPERATIONDENIED_RECEIVERNOTEXIST = "OperationDenied.ReceiverNotExist"
func (c *Client) DeleteAddressUnsubscribeConfigWithContext(ctx context.Context, request *DeleteAddressUnsubscribeConfigRequest) (response *DeleteAddressUnsubscribeConfigResponse, err error) {
    if request == nil {
        request = NewDeleteAddressUnsubscribeConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ses", APIVersion, "DeleteAddressUnsubscribeConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAddressUnsubscribeConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAddressUnsubscribeConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBlackListRequest() (request *DeleteBlackListRequest) {
    request = &DeleteBlackListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ses", APIVersion, "DeleteBlackList")
    
    
    return
}

func NewDeleteBlackListResponse() (response *DeleteBlackListResponse) {
    response = &DeleteBlackListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteBlackList
// This API is used to unblocklist email addresses. If you confirm that a blocklisted recipient address is valid and active, you can remove it from Tencent Cloud’s address blocklist database.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DeleteBlackList(request *DeleteBlackListRequest) (response *DeleteBlackListResponse, err error) {
    return c.DeleteBlackListWithContext(context.Background(), request)
}

// DeleteBlackList
// This API is used to unblocklist email addresses. If you confirm that a blocklisted recipient address is valid and active, you can remove it from Tencent Cloud’s address blocklist database.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DeleteBlackListWithContext(ctx context.Context, request *DeleteBlackListRequest) (response *DeleteBlackListResponse, err error) {
    if request == nil {
        request = NewDeleteBlackListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ses", APIVersion, "DeleteBlackList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteBlackList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteBlackListResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEmailAddressRequest() (request *DeleteEmailAddressRequest) {
    request = &DeleteEmailAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ses", APIVersion, "DeleteEmailAddress")
    
    
    return
}

func NewDeleteEmailAddressResponse() (response *DeleteEmailAddressResponse) {
    response = &DeleteEmailAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteEmailAddress
// This API is used to delete a sender address.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NOSUCHSENDER = "InvalidParameterValue.NoSuchSender"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DeleteEmailAddress(request *DeleteEmailAddressRequest) (response *DeleteEmailAddressResponse, err error) {
    return c.DeleteEmailAddressWithContext(context.Background(), request)
}

// DeleteEmailAddress
// This API is used to delete a sender address.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NOSUCHSENDER = "InvalidParameterValue.NoSuchSender"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DeleteEmailAddressWithContext(ctx context.Context, request *DeleteEmailAddressRequest) (response *DeleteEmailAddressResponse, err error) {
    if request == nil {
        request = NewDeleteEmailAddressRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ses", APIVersion, "DeleteEmailAddress")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteEmailAddress require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteEmailAddressResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEmailIdentityRequest() (request *DeleteEmailIdentityRequest) {
    request = &DeleteEmailIdentityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ses", APIVersion, "DeleteEmailIdentity")
    
    
    return
}

func NewDeleteEmailIdentityResponse() (response *DeleteEmailIdentityResponse) {
    response = &DeleteEmailIdentityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteEmailIdentity
// This API is used to delete a sender domain. After deleted, the sender domain can no longer be used to send emails.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DeleteEmailIdentity(request *DeleteEmailIdentityRequest) (response *DeleteEmailIdentityResponse, err error) {
    return c.DeleteEmailIdentityWithContext(context.Background(), request)
}

// DeleteEmailIdentity
// This API is used to delete a sender domain. After deleted, the sender domain can no longer be used to send emails.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DeleteEmailIdentityWithContext(ctx context.Context, request *DeleteEmailIdentityRequest) (response *DeleteEmailIdentityResponse, err error) {
    if request == nil {
        request = NewDeleteEmailIdentityRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ses", APIVersion, "DeleteEmailIdentity")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteEmailIdentity require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteEmailIdentityResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEmailTemplateRequest() (request *DeleteEmailTemplateRequest) {
    request = &DeleteEmailTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ses", APIVersion, "DeleteEmailTemplate")
    
    
    return
}

func NewDeleteEmailTemplateResponse() (response *DeleteEmailTemplateResponse) {
    response = &DeleteEmailTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteEmailTemplate
// This API is used to delete an email template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DeleteEmailTemplate(request *DeleteEmailTemplateRequest) (response *DeleteEmailTemplateResponse, err error) {
    return c.DeleteEmailTemplateWithContext(context.Background(), request)
}

// DeleteEmailTemplate
// This API is used to delete an email template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DeleteEmailTemplateWithContext(ctx context.Context, request *DeleteEmailTemplateRequest) (response *DeleteEmailTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteEmailTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ses", APIVersion, "DeleteEmailTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteEmailTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteEmailTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteReceiverRequest() (request *DeleteReceiverRequest) {
    request = &DeleteReceiverRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ses", APIVersion, "DeleteReceiver")
    
    
    return
}

func NewDeleteReceiverResponse() (response *DeleteReceiverResponse) {
    response = &DeleteReceiverResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteReceiver
// This API is used to delete a recipient group and all recipient email addresses in the group based on the recipient group ID.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER_RECEIVERIDNECESSARY = "MissingParameter.ReceiverIdNecessary"
//  OPERATIONDENIED_RECEIVERNOTEXIST = "OperationDenied.ReceiverNotExist"
func (c *Client) DeleteReceiver(request *DeleteReceiverRequest) (response *DeleteReceiverResponse, err error) {
    return c.DeleteReceiverWithContext(context.Background(), request)
}

// DeleteReceiver
// This API is used to delete a recipient group and all recipient email addresses in the group based on the recipient group ID.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER_RECEIVERIDNECESSARY = "MissingParameter.ReceiverIdNecessary"
//  OPERATIONDENIED_RECEIVERNOTEXIST = "OperationDenied.ReceiverNotExist"
func (c *Client) DeleteReceiverWithContext(ctx context.Context, request *DeleteReceiverRequest) (response *DeleteReceiverResponse, err error) {
    if request == nil {
        request = NewDeleteReceiverRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ses", APIVersion, "DeleteReceiver")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteReceiver require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteReceiverResponse()
    err = c.Send(request, response)
    return
}

func NewGetEmailIdentityRequest() (request *GetEmailIdentityRequest) {
    request = &GetEmailIdentityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ses", APIVersion, "GetEmailIdentity")
    
    
    return
}

func NewGetEmailIdentityResponse() (response *GetEmailIdentityResponse) {
    response = &GetEmailIdentityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetEmailIdentity
// This API is used to get the configuration details of a sender domain.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NOTEXISTDOMAIN = "InvalidParameterValue.NotExistDomain"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) GetEmailIdentity(request *GetEmailIdentityRequest) (response *GetEmailIdentityResponse, err error) {
    return c.GetEmailIdentityWithContext(context.Background(), request)
}

// GetEmailIdentity
// This API is used to get the configuration details of a sender domain.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NOTEXISTDOMAIN = "InvalidParameterValue.NotExistDomain"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) GetEmailIdentityWithContext(ctx context.Context, request *GetEmailIdentityRequest) (response *GetEmailIdentityResponse, err error) {
    if request == nil {
        request = NewGetEmailIdentityRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ses", APIVersion, "GetEmailIdentity")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetEmailIdentity require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetEmailIdentityResponse()
    err = c.Send(request, response)
    return
}

func NewGetEmailTemplateRequest() (request *GetEmailTemplateRequest) {
    request = &GetEmailTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ses", APIVersion, "GetEmailTemplate")
    
    
    return
}

func NewGetEmailTemplateResponse() (response *GetEmailTemplateResponse) {
    response = &GetEmailTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetEmailTemplate
// This API is used to get the details of a template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TEMPLATENOTEXIST = "InvalidParameterValue.TemplateNotExist"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetEmailTemplate(request *GetEmailTemplateRequest) (response *GetEmailTemplateResponse, err error) {
    return c.GetEmailTemplateWithContext(context.Background(), request)
}

// GetEmailTemplate
// This API is used to get the details of a template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TEMPLATENOTEXIST = "InvalidParameterValue.TemplateNotExist"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetEmailTemplateWithContext(ctx context.Context, request *GetEmailTemplateRequest) (response *GetEmailTemplateResponse, err error) {
    if request == nil {
        request = NewGetEmailTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ses", APIVersion, "GetEmailTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetEmailTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetEmailTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewGetSendEmailStatusRequest() (request *GetSendEmailStatusRequest) {
    request = &GetSendEmailStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ses", APIVersion, "GetSendEmailStatus")
    
    
    return
}

func NewGetSendEmailStatusResponse() (response *GetSendEmailStatusResponse) {
    response = &GetSendEmailStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetSendEmailStatus
// This API is used to get email sending status. Only data within 30 days can be queried.
//
// Default API request rate limit: 1 request/sec.
//
// error code that may be returned:
//  FAILEDOPERATION_EMAILADDRINBLACKLIST = "FailedOperation.EmailAddrInBlacklist"
//  FAILEDOPERATION_EMAILCONTENTTOOLARGE = "FailedOperation.EmailContentToolarge"
//  FAILEDOPERATION_EXCEEDSENDLIMIT = "FailedOperation.ExceedSendLimit"
//  FAILEDOPERATION_FREQUENCYLIMIT = "FailedOperation.FrequencyLimit"
//  FAILEDOPERATION_HIGHREJECTIONRATE = "FailedOperation.HighRejectionRate"
//  FAILEDOPERATION_INCORRECTEMAIL = "FailedOperation.IncorrectEmail"
//  FAILEDOPERATION_INSUFFICIENTBALANCE = "FailedOperation.InsufficientBalance"
//  FAILEDOPERATION_INSUFFICIENTQUOTA = "FailedOperation.InsufficientQuota"
//  FAILEDOPERATION_INVALIDLIMIT = "FailedOperation.InvalidLimit"
//  FAILEDOPERATION_INVALIDTEMPLATEID = "FailedOperation.InvalidTemplateID"
//  FAILEDOPERATION_NOTAUTHENTICATEDSENDER = "FailedOperation.NotAuthenticatedSender"
//  FAILEDOPERATION_NOTSUPPORTDATE = "FailedOperation.NotSupportDate"
//  FAILEDOPERATION_PROTOCOLCHECKERR = "FailedOperation.ProtocolCheckErr"
//  FAILEDOPERATION_TEMPORARYBLOCKED = "FailedOperation.TemporaryBlocked"
//  FAILEDOPERATION_TOOMANYRECIPIENTS = "FailedOperation.TooManyRecipients"
//  FAILEDOPERATION_UNSUPPORTMAILTYPE = "FailedOperation.UnsupportMailType"
//  FAILEDOPERATION_WITHOUTPERMISSION = "FailedOperation.WithOutPermission"
//  FAILEDOPERATION_WRONGCONTENTJSON = "FailedOperation.WrongContentJson"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EMAILADDRESSISNULL = "InvalidParameterValue.EmailAddressIsNULL"
//  INVALIDPARAMETERVALUE_EMAILCONTENTISWRONG = "InvalidParameterValue.EmailContentIsWrong"
//  INVALIDPARAMETERVALUE_WRONGDATE = "InvalidParameterValue.WrongDate"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetSendEmailStatus(request *GetSendEmailStatusRequest) (response *GetSendEmailStatusResponse, err error) {
    return c.GetSendEmailStatusWithContext(context.Background(), request)
}

// GetSendEmailStatus
// This API is used to get email sending status. Only data within 30 days can be queried.
//
// Default API request rate limit: 1 request/sec.
//
// error code that may be returned:
//  FAILEDOPERATION_EMAILADDRINBLACKLIST = "FailedOperation.EmailAddrInBlacklist"
//  FAILEDOPERATION_EMAILCONTENTTOOLARGE = "FailedOperation.EmailContentToolarge"
//  FAILEDOPERATION_EXCEEDSENDLIMIT = "FailedOperation.ExceedSendLimit"
//  FAILEDOPERATION_FREQUENCYLIMIT = "FailedOperation.FrequencyLimit"
//  FAILEDOPERATION_HIGHREJECTIONRATE = "FailedOperation.HighRejectionRate"
//  FAILEDOPERATION_INCORRECTEMAIL = "FailedOperation.IncorrectEmail"
//  FAILEDOPERATION_INSUFFICIENTBALANCE = "FailedOperation.InsufficientBalance"
//  FAILEDOPERATION_INSUFFICIENTQUOTA = "FailedOperation.InsufficientQuota"
//  FAILEDOPERATION_INVALIDLIMIT = "FailedOperation.InvalidLimit"
//  FAILEDOPERATION_INVALIDTEMPLATEID = "FailedOperation.InvalidTemplateID"
//  FAILEDOPERATION_NOTAUTHENTICATEDSENDER = "FailedOperation.NotAuthenticatedSender"
//  FAILEDOPERATION_NOTSUPPORTDATE = "FailedOperation.NotSupportDate"
//  FAILEDOPERATION_PROTOCOLCHECKERR = "FailedOperation.ProtocolCheckErr"
//  FAILEDOPERATION_TEMPORARYBLOCKED = "FailedOperation.TemporaryBlocked"
//  FAILEDOPERATION_TOOMANYRECIPIENTS = "FailedOperation.TooManyRecipients"
//  FAILEDOPERATION_UNSUPPORTMAILTYPE = "FailedOperation.UnsupportMailType"
//  FAILEDOPERATION_WITHOUTPERMISSION = "FailedOperation.WithOutPermission"
//  FAILEDOPERATION_WRONGCONTENTJSON = "FailedOperation.WrongContentJson"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EMAILADDRESSISNULL = "InvalidParameterValue.EmailAddressIsNULL"
//  INVALIDPARAMETERVALUE_EMAILCONTENTISWRONG = "InvalidParameterValue.EmailContentIsWrong"
//  INVALIDPARAMETERVALUE_WRONGDATE = "InvalidParameterValue.WrongDate"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetSendEmailStatusWithContext(ctx context.Context, request *GetSendEmailStatusRequest) (response *GetSendEmailStatusResponse, err error) {
    if request == nil {
        request = NewGetSendEmailStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ses", APIVersion, "GetSendEmailStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetSendEmailStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetSendEmailStatusResponse()
    err = c.Send(request, response)
    return
}

func NewGetStatisticsReportRequest() (request *GetStatisticsReportRequest) {
    request = &GetStatisticsReportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ses", APIVersion, "GetStatisticsReport")
    
    
    return
}

func NewGetStatisticsReportResponse() (response *GetStatisticsReportResponse) {
    response = &GetStatisticsReportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetStatisticsReport
// This API is used to get the email sending statistics over a recent period, including data on sent emails, delivery success rate, open rate, bounce rate, and so on.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_WRONGDATE = "InvalidParameterValue.WrongDate"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) GetStatisticsReport(request *GetStatisticsReportRequest) (response *GetStatisticsReportResponse, err error) {
    return c.GetStatisticsReportWithContext(context.Background(), request)
}

// GetStatisticsReport
// This API is used to get the email sending statistics over a recent period, including data on sent emails, delivery success rate, open rate, bounce rate, and so on.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_WRONGDATE = "InvalidParameterValue.WrongDate"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) GetStatisticsReportWithContext(ctx context.Context, request *GetStatisticsReportRequest) (response *GetStatisticsReportResponse, err error) {
    if request == nil {
        request = NewGetStatisticsReportRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ses", APIVersion, "GetStatisticsReport")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetStatisticsReport require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetStatisticsReportResponse()
    err = c.Send(request, response)
    return
}

func NewListAddressUnsubscribeConfigRequest() (request *ListAddressUnsubscribeConfigRequest) {
    request = &ListAddressUnsubscribeConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ses", APIVersion, "ListAddressUnsubscribeConfig")
    
    
    return
}

func NewListAddressUnsubscribeConfigResponse() (response *ListAddressUnsubscribeConfigResponse) {
    response = &ListAddressUnsubscribeConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListAddressUnsubscribeConfig
// This API is used to get the address and unsubscribe configuration list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_WRONGDATE = "InvalidParameterValue.WrongDate"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ListAddressUnsubscribeConfig(request *ListAddressUnsubscribeConfigRequest) (response *ListAddressUnsubscribeConfigResponse, err error) {
    return c.ListAddressUnsubscribeConfigWithContext(context.Background(), request)
}

// ListAddressUnsubscribeConfig
// This API is used to get the address and unsubscribe configuration list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_WRONGDATE = "InvalidParameterValue.WrongDate"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ListAddressUnsubscribeConfigWithContext(ctx context.Context, request *ListAddressUnsubscribeConfigRequest) (response *ListAddressUnsubscribeConfigResponse, err error) {
    if request == nil {
        request = NewListAddressUnsubscribeConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ses", APIVersion, "ListAddressUnsubscribeConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAddressUnsubscribeConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAddressUnsubscribeConfigResponse()
    err = c.Send(request, response)
    return
}

func NewListBlackEmailAddressRequest() (request *ListBlackEmailAddressRequest) {
    request = &ListBlackEmailAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ses", APIVersion, "ListBlackEmailAddress")
    
    
    return
}

func NewListBlackEmailAddressResponse() (response *ListBlackEmailAddressResponse) {
    response = &ListBlackEmailAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListBlackEmailAddress
// The API is used to get blocklisted addresses. In the case of a hard bounce, Tencent Cloud will blocklist the recipient address and do not allow any user to send emails to this address. If you confirm that this is a misjudgment, you can remove it from the blocklist.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDLIMIT = "FailedOperation.InvalidLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_WRONGDATE = "InvalidParameterValue.WrongDate"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ListBlackEmailAddress(request *ListBlackEmailAddressRequest) (response *ListBlackEmailAddressResponse, err error) {
    return c.ListBlackEmailAddressWithContext(context.Background(), request)
}

// ListBlackEmailAddress
// The API is used to get blocklisted addresses. In the case of a hard bounce, Tencent Cloud will blocklist the recipient address and do not allow any user to send emails to this address. If you confirm that this is a misjudgment, you can remove it from the blocklist.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDLIMIT = "FailedOperation.InvalidLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_WRONGDATE = "InvalidParameterValue.WrongDate"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ListBlackEmailAddressWithContext(ctx context.Context, request *ListBlackEmailAddressRequest) (response *ListBlackEmailAddressResponse, err error) {
    if request == nil {
        request = NewListBlackEmailAddressRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ses", APIVersion, "ListBlackEmailAddress")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListBlackEmailAddress require credential")
    }

    request.SetContext(ctx)
    
    response = NewListBlackEmailAddressResponse()
    err = c.Send(request, response)
    return
}

func NewListEmailAddressRequest() (request *ListEmailAddressRequest) {
    request = &ListEmailAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ses", APIVersion, "ListEmailAddress")
    
    
    return
}

func NewListEmailAddressResponse() (response *ListEmailAddressResponse) {
    response = &ListEmailAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListEmailAddress
// This API is used to get the list of sender addresses.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ListEmailAddress(request *ListEmailAddressRequest) (response *ListEmailAddressResponse, err error) {
    return c.ListEmailAddressWithContext(context.Background(), request)
}

// ListEmailAddress
// This API is used to get the list of sender addresses.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ListEmailAddressWithContext(ctx context.Context, request *ListEmailAddressRequest) (response *ListEmailAddressResponse, err error) {
    if request == nil {
        request = NewListEmailAddressRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ses", APIVersion, "ListEmailAddress")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListEmailAddress require credential")
    }

    request.SetContext(ctx)
    
    response = NewListEmailAddressResponse()
    err = c.Send(request, response)
    return
}

func NewListEmailIdentitiesRequest() (request *ListEmailIdentitiesRequest) {
    request = &ListEmailIdentitiesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ses", APIVersion, "ListEmailIdentities")
    
    
    return
}

func NewListEmailIdentitiesResponse() (response *ListEmailIdentitiesResponse) {
    response = &ListEmailIdentitiesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListEmailIdentities
// This API is used to get the list of sender domains, including verified and unverified domains.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ListEmailIdentities(request *ListEmailIdentitiesRequest) (response *ListEmailIdentitiesResponse, err error) {
    return c.ListEmailIdentitiesWithContext(context.Background(), request)
}

// ListEmailIdentities
// This API is used to get the list of sender domains, including verified and unverified domains.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ListEmailIdentitiesWithContext(ctx context.Context, request *ListEmailIdentitiesRequest) (response *ListEmailIdentitiesResponse, err error) {
    if request == nil {
        request = NewListEmailIdentitiesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ses", APIVersion, "ListEmailIdentities")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListEmailIdentities require credential")
    }

    request.SetContext(ctx)
    
    response = NewListEmailIdentitiesResponse()
    err = c.Send(request, response)
    return
}

func NewListEmailTemplatesRequest() (request *ListEmailTemplatesRequest) {
    request = &ListEmailTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ses", APIVersion, "ListEmailTemplates")
    
    
    return
}

func NewListEmailTemplatesResponse() (response *ListEmailTemplatesResponse) {
    response = &ListEmailTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListEmailTemplates
// This API is used to get the list of email templates.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDLIMIT = "FailedOperation.InvalidLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ListEmailTemplates(request *ListEmailTemplatesRequest) (response *ListEmailTemplatesResponse, err error) {
    return c.ListEmailTemplatesWithContext(context.Background(), request)
}

// ListEmailTemplates
// This API is used to get the list of email templates.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDLIMIT = "FailedOperation.InvalidLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ListEmailTemplatesWithContext(ctx context.Context, request *ListEmailTemplatesRequest) (response *ListEmailTemplatesResponse, err error) {
    if request == nil {
        request = NewListEmailTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ses", APIVersion, "ListEmailTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListEmailTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewListEmailTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewListReceiversRequest() (request *ListReceiversRequest) {
    request = &ListReceiversRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ses", APIVersion, "ListReceivers")
    
    
    return
}

func NewListReceiversResponse() (response *ListReceiversResponse) {
    response = &ListReceiversResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListReceivers
// This API is used to query recipient groups. It supports pagination, fuzzy query, and query by status.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDLIMIT = "FailedOperation.InvalidLimit"
//  INTERNALERROR = "InternalError"
func (c *Client) ListReceivers(request *ListReceiversRequest) (response *ListReceiversResponse, err error) {
    return c.ListReceiversWithContext(context.Background(), request)
}

// ListReceivers
// This API is used to query recipient groups. It supports pagination, fuzzy query, and query by status.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDLIMIT = "FailedOperation.InvalidLimit"
//  INTERNALERROR = "InternalError"
func (c *Client) ListReceiversWithContext(ctx context.Context, request *ListReceiversRequest) (response *ListReceiversResponse, err error) {
    if request == nil {
        request = NewListReceiversRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ses", APIVersion, "ListReceivers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListReceivers require credential")
    }

    request.SetContext(ctx)
    
    response = NewListReceiversResponse()
    err = c.Send(request, response)
    return
}

func NewListSendTasksRequest() (request *ListSendTasksRequest) {
    request = &ListSendTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ses", APIVersion, "ListSendTasks")
    
    
    return
}

func NewListSendTasksResponse() (response *ListSendTasksResponse) {
    response = &ListSendTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListSendTasks
// This API is used to query batch email sending tasks (including immediate, scheduled, and recurring tasks) by page. You can query task data including the number of emails requested to be sent, the number of sent emails, the number of cached emails, and task status.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDLIMIT = "FailedOperation.InvalidLimit"
func (c *Client) ListSendTasks(request *ListSendTasksRequest) (response *ListSendTasksResponse, err error) {
    return c.ListSendTasksWithContext(context.Background(), request)
}

// ListSendTasks
// This API is used to query batch email sending tasks (including immediate, scheduled, and recurring tasks) by page. You can query task data including the number of emails requested to be sent, the number of sent emails, the number of cached emails, and task status.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDLIMIT = "FailedOperation.InvalidLimit"
func (c *Client) ListSendTasksWithContext(ctx context.Context, request *ListSendTasksRequest) (response *ListSendTasksResponse, err error) {
    if request == nil {
        request = NewListSendTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ses", APIVersion, "ListSendTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListSendTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewListSendTasksResponse()
    err = c.Send(request, response)
    return
}

func NewSendEmailRequest() (request *SendEmailRequest) {
    request = &SendEmailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ses", APIVersion, "SendEmail")
    
    
    return
}

func NewSendEmailResponse() (response *SendEmailResponse) {
    response = &SendEmailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SendEmail
// This API is used to send an HTML or TEXT email triggered for authentication or transaction. By default, you can send emails using a template only.
//
// error code that may be returned:
//  FAILEDOPERATION_ATTACHCONTENTTOOLARGE = "FailedOperation.AttachContentToolarge"
//  FAILEDOPERATION_EMAILADDRINBLACKLIST = "FailedOperation.EmailAddrInBlacklist"
//  FAILEDOPERATION_EMAILCONTENTTOOLARGE = "FailedOperation.EmailContentToolarge"
//  FAILEDOPERATION_EXCEEDSENDLIMIT = "FailedOperation.ExceedSendLimit"
//  FAILEDOPERATION_FREQUENCYLIMIT = "FailedOperation.FrequencyLimit"
//  FAILEDOPERATION_HIGHREJECTIONRATE = "FailedOperation.HighRejectionRate"
//  FAILEDOPERATION_INCORRECTEMAIL = "FailedOperation.IncorrectEmail"
//  FAILEDOPERATION_INCORRECTSENDER = "FailedOperation.IncorrectSender"
//  FAILEDOPERATION_INSUFFICIENTBALANCE = "FailedOperation.InsufficientBalance"
//  FAILEDOPERATION_INSUFFICIENTQUOTA = "FailedOperation.InsufficientQuota"
//  FAILEDOPERATION_INVALIDATTACHNAME = "FailedOperation.InvalidAttachName"
//  FAILEDOPERATION_INVALIDTEMPLATEID = "FailedOperation.InvalidTemplateID"
//  FAILEDOPERATION_MISSINGEMAILCONTENT = "FailedOperation.MissingEmailContent"
//  FAILEDOPERATION_NOATTACHPERMISSION = "FailedOperation.NoAttachPermission"
//  FAILEDOPERATION_NOTAUTHENTICATEDSENDER = "FailedOperation.NotAuthenticatedSender"
//  FAILEDOPERATION_PROTOCOLCHECKERR = "FailedOperation.ProtocolCheckErr"
//  FAILEDOPERATION_RECEIVERHASUNSUBSCRIBED = "FailedOperation.ReceiverHasUnsubscribed"
//  FAILEDOPERATION_REJECTEDBYRECIPIENTS = "FailedOperation.RejectedByRecipients"
//  FAILEDOPERATION_SENDEMAILERR = "FailedOperation.SendEmailErr"
//  FAILEDOPERATION_TEMPORARYBLOCKED = "FailedOperation.TemporaryBlocked"
//  FAILEDOPERATION_TOOMANYATTACHMENTS = "FailedOperation.TooManyAttachments"
//  FAILEDOPERATION_TOOMANYRECIPIENTS = "FailedOperation.TooManyRecipients"
//  FAILEDOPERATION_UNSUPPORTMAILTYPE = "FailedOperation.UnsupportMailType"
//  FAILEDOPERATION_WITHOUTPERMISSION = "FailedOperation.WithOutPermission"
//  FAILEDOPERATION_WRONGCONTENTJSON = "FailedOperation.WrongContentJson"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ATTACHCONTENTISWRONG = "InvalidParameterValue.AttachContentIsWrong"
//  INVALIDPARAMETERVALUE_EMAILADDRESSISNULL = "InvalidParameterValue.EmailAddressIsNULL"
//  INVALIDPARAMETERVALUE_EMAILCONTENTISWRONG = "InvalidParameterValue.EmailContentIsWrong"
//  INVALIDPARAMETERVALUE_INVALIDEMAILIDENTITY = "InvalidParameterValue.InvalidEmailIdentity"
//  INVALIDPARAMETERVALUE_RECEIVEREMAILINVALID = "InvalidParameterValue.ReceiverEmailInvalid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SendEmail(request *SendEmailRequest) (response *SendEmailResponse, err error) {
    return c.SendEmailWithContext(context.Background(), request)
}

// SendEmail
// This API is used to send an HTML or TEXT email triggered for authentication or transaction. By default, you can send emails using a template only.
//
// error code that may be returned:
//  FAILEDOPERATION_ATTACHCONTENTTOOLARGE = "FailedOperation.AttachContentToolarge"
//  FAILEDOPERATION_EMAILADDRINBLACKLIST = "FailedOperation.EmailAddrInBlacklist"
//  FAILEDOPERATION_EMAILCONTENTTOOLARGE = "FailedOperation.EmailContentToolarge"
//  FAILEDOPERATION_EXCEEDSENDLIMIT = "FailedOperation.ExceedSendLimit"
//  FAILEDOPERATION_FREQUENCYLIMIT = "FailedOperation.FrequencyLimit"
//  FAILEDOPERATION_HIGHREJECTIONRATE = "FailedOperation.HighRejectionRate"
//  FAILEDOPERATION_INCORRECTEMAIL = "FailedOperation.IncorrectEmail"
//  FAILEDOPERATION_INCORRECTSENDER = "FailedOperation.IncorrectSender"
//  FAILEDOPERATION_INSUFFICIENTBALANCE = "FailedOperation.InsufficientBalance"
//  FAILEDOPERATION_INSUFFICIENTQUOTA = "FailedOperation.InsufficientQuota"
//  FAILEDOPERATION_INVALIDATTACHNAME = "FailedOperation.InvalidAttachName"
//  FAILEDOPERATION_INVALIDTEMPLATEID = "FailedOperation.InvalidTemplateID"
//  FAILEDOPERATION_MISSINGEMAILCONTENT = "FailedOperation.MissingEmailContent"
//  FAILEDOPERATION_NOATTACHPERMISSION = "FailedOperation.NoAttachPermission"
//  FAILEDOPERATION_NOTAUTHENTICATEDSENDER = "FailedOperation.NotAuthenticatedSender"
//  FAILEDOPERATION_PROTOCOLCHECKERR = "FailedOperation.ProtocolCheckErr"
//  FAILEDOPERATION_RECEIVERHASUNSUBSCRIBED = "FailedOperation.ReceiverHasUnsubscribed"
//  FAILEDOPERATION_REJECTEDBYRECIPIENTS = "FailedOperation.RejectedByRecipients"
//  FAILEDOPERATION_SENDEMAILERR = "FailedOperation.SendEmailErr"
//  FAILEDOPERATION_TEMPORARYBLOCKED = "FailedOperation.TemporaryBlocked"
//  FAILEDOPERATION_TOOMANYATTACHMENTS = "FailedOperation.TooManyAttachments"
//  FAILEDOPERATION_TOOMANYRECIPIENTS = "FailedOperation.TooManyRecipients"
//  FAILEDOPERATION_UNSUPPORTMAILTYPE = "FailedOperation.UnsupportMailType"
//  FAILEDOPERATION_WITHOUTPERMISSION = "FailedOperation.WithOutPermission"
//  FAILEDOPERATION_WRONGCONTENTJSON = "FailedOperation.WrongContentJson"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ATTACHCONTENTISWRONG = "InvalidParameterValue.AttachContentIsWrong"
//  INVALIDPARAMETERVALUE_EMAILADDRESSISNULL = "InvalidParameterValue.EmailAddressIsNULL"
//  INVALIDPARAMETERVALUE_EMAILCONTENTISWRONG = "InvalidParameterValue.EmailContentIsWrong"
//  INVALIDPARAMETERVALUE_INVALIDEMAILIDENTITY = "InvalidParameterValue.InvalidEmailIdentity"
//  INVALIDPARAMETERVALUE_RECEIVEREMAILINVALID = "InvalidParameterValue.ReceiverEmailInvalid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SendEmailWithContext(ctx context.Context, request *SendEmailRequest) (response *SendEmailResponse, err error) {
    if request == nil {
        request = NewSendEmailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ses", APIVersion, "SendEmail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SendEmail require credential")
    }

    request.SetContext(ctx)
    
    response = NewSendEmailResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAddressUnsubscribeConfigRequest() (request *UpdateAddressUnsubscribeConfigRequest) {
    request = &UpdateAddressUnsubscribeConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ses", APIVersion, "UpdateAddressUnsubscribeConfig")
    
    
    return
}

func NewUpdateAddressUnsubscribeConfigResponse() (response *UpdateAddressUnsubscribeConfigResponse) {
    response = &UpdateAddressUnsubscribeConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateAddressUnsubscribeConfig
// This API is used to update address-level unsubscribe configurations.
//
// error code that may be returned:
//  FAILEDOPERATION_ATTACHCONTENTTOOLARGE = "FailedOperation.AttachContentToolarge"
//  FAILEDOPERATION_EMAILADDRINBLACKLIST = "FailedOperation.EmailAddrInBlacklist"
//  FAILEDOPERATION_EMAILCONTENTTOOLARGE = "FailedOperation.EmailContentToolarge"
//  FAILEDOPERATION_EXCEEDSENDLIMIT = "FailedOperation.ExceedSendLimit"
//  FAILEDOPERATION_FREQUENCYLIMIT = "FailedOperation.FrequencyLimit"
//  FAILEDOPERATION_HIGHREJECTIONRATE = "FailedOperation.HighRejectionRate"
//  FAILEDOPERATION_INCORRECTEMAIL = "FailedOperation.IncorrectEmail"
//  FAILEDOPERATION_INCORRECTSENDER = "FailedOperation.IncorrectSender"
//  FAILEDOPERATION_INSUFFICIENTBALANCE = "FailedOperation.InsufficientBalance"
//  FAILEDOPERATION_INSUFFICIENTQUOTA = "FailedOperation.InsufficientQuota"
//  FAILEDOPERATION_INVALIDATTACHNAME = "FailedOperation.InvalidAttachName"
//  FAILEDOPERATION_INVALIDTEMPLATEID = "FailedOperation.InvalidTemplateID"
//  FAILEDOPERATION_MISSINGEMAILCONTENT = "FailedOperation.MissingEmailContent"
//  FAILEDOPERATION_NOATTACHPERMISSION = "FailedOperation.NoAttachPermission"
//  FAILEDOPERATION_NOTAUTHENTICATEDSENDER = "FailedOperation.NotAuthenticatedSender"
//  FAILEDOPERATION_PROTOCOLCHECKERR = "FailedOperation.ProtocolCheckErr"
//  FAILEDOPERATION_RECEIVERHASUNSUBSCRIBED = "FailedOperation.ReceiverHasUnsubscribed"
//  FAILEDOPERATION_REJECTEDBYRECIPIENTS = "FailedOperation.RejectedByRecipients"
//  FAILEDOPERATION_SENDEMAILERR = "FailedOperation.SendEmailErr"
//  FAILEDOPERATION_TEMPORARYBLOCKED = "FailedOperation.TemporaryBlocked"
//  FAILEDOPERATION_TOOMANYATTACHMENTS = "FailedOperation.TooManyAttachments"
//  FAILEDOPERATION_TOOMANYRECIPIENTS = "FailedOperation.TooManyRecipients"
//  FAILEDOPERATION_UNSUPPORTMAILTYPE = "FailedOperation.UnsupportMailType"
//  FAILEDOPERATION_WITHOUTPERMISSION = "FailedOperation.WithOutPermission"
//  FAILEDOPERATION_WRONGCONTENTJSON = "FailedOperation.WrongContentJson"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ATTACHCONTENTISWRONG = "InvalidParameterValue.AttachContentIsWrong"
//  INVALIDPARAMETERVALUE_EMAILADDRESSISNULL = "InvalidParameterValue.EmailAddressIsNULL"
//  INVALIDPARAMETERVALUE_EMAILCONTENTISWRONG = "InvalidParameterValue.EmailContentIsWrong"
//  INVALIDPARAMETERVALUE_INVALIDEMAILIDENTITY = "InvalidParameterValue.InvalidEmailIdentity"
//  INVALIDPARAMETERVALUE_RECEIVEREMAILINVALID = "InvalidParameterValue.ReceiverEmailInvalid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateAddressUnsubscribeConfig(request *UpdateAddressUnsubscribeConfigRequest) (response *UpdateAddressUnsubscribeConfigResponse, err error) {
    return c.UpdateAddressUnsubscribeConfigWithContext(context.Background(), request)
}

// UpdateAddressUnsubscribeConfig
// This API is used to update address-level unsubscribe configurations.
//
// error code that may be returned:
//  FAILEDOPERATION_ATTACHCONTENTTOOLARGE = "FailedOperation.AttachContentToolarge"
//  FAILEDOPERATION_EMAILADDRINBLACKLIST = "FailedOperation.EmailAddrInBlacklist"
//  FAILEDOPERATION_EMAILCONTENTTOOLARGE = "FailedOperation.EmailContentToolarge"
//  FAILEDOPERATION_EXCEEDSENDLIMIT = "FailedOperation.ExceedSendLimit"
//  FAILEDOPERATION_FREQUENCYLIMIT = "FailedOperation.FrequencyLimit"
//  FAILEDOPERATION_HIGHREJECTIONRATE = "FailedOperation.HighRejectionRate"
//  FAILEDOPERATION_INCORRECTEMAIL = "FailedOperation.IncorrectEmail"
//  FAILEDOPERATION_INCORRECTSENDER = "FailedOperation.IncorrectSender"
//  FAILEDOPERATION_INSUFFICIENTBALANCE = "FailedOperation.InsufficientBalance"
//  FAILEDOPERATION_INSUFFICIENTQUOTA = "FailedOperation.InsufficientQuota"
//  FAILEDOPERATION_INVALIDATTACHNAME = "FailedOperation.InvalidAttachName"
//  FAILEDOPERATION_INVALIDTEMPLATEID = "FailedOperation.InvalidTemplateID"
//  FAILEDOPERATION_MISSINGEMAILCONTENT = "FailedOperation.MissingEmailContent"
//  FAILEDOPERATION_NOATTACHPERMISSION = "FailedOperation.NoAttachPermission"
//  FAILEDOPERATION_NOTAUTHENTICATEDSENDER = "FailedOperation.NotAuthenticatedSender"
//  FAILEDOPERATION_PROTOCOLCHECKERR = "FailedOperation.ProtocolCheckErr"
//  FAILEDOPERATION_RECEIVERHASUNSUBSCRIBED = "FailedOperation.ReceiverHasUnsubscribed"
//  FAILEDOPERATION_REJECTEDBYRECIPIENTS = "FailedOperation.RejectedByRecipients"
//  FAILEDOPERATION_SENDEMAILERR = "FailedOperation.SendEmailErr"
//  FAILEDOPERATION_TEMPORARYBLOCKED = "FailedOperation.TemporaryBlocked"
//  FAILEDOPERATION_TOOMANYATTACHMENTS = "FailedOperation.TooManyAttachments"
//  FAILEDOPERATION_TOOMANYRECIPIENTS = "FailedOperation.TooManyRecipients"
//  FAILEDOPERATION_UNSUPPORTMAILTYPE = "FailedOperation.UnsupportMailType"
//  FAILEDOPERATION_WITHOUTPERMISSION = "FailedOperation.WithOutPermission"
//  FAILEDOPERATION_WRONGCONTENTJSON = "FailedOperation.WrongContentJson"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ATTACHCONTENTISWRONG = "InvalidParameterValue.AttachContentIsWrong"
//  INVALIDPARAMETERVALUE_EMAILADDRESSISNULL = "InvalidParameterValue.EmailAddressIsNULL"
//  INVALIDPARAMETERVALUE_EMAILCONTENTISWRONG = "InvalidParameterValue.EmailContentIsWrong"
//  INVALIDPARAMETERVALUE_INVALIDEMAILIDENTITY = "InvalidParameterValue.InvalidEmailIdentity"
//  INVALIDPARAMETERVALUE_RECEIVEREMAILINVALID = "InvalidParameterValue.ReceiverEmailInvalid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateAddressUnsubscribeConfigWithContext(ctx context.Context, request *UpdateAddressUnsubscribeConfigRequest) (response *UpdateAddressUnsubscribeConfigResponse, err error) {
    if request == nil {
        request = NewUpdateAddressUnsubscribeConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ses", APIVersion, "UpdateAddressUnsubscribeConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateAddressUnsubscribeConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateAddressUnsubscribeConfigResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateEmailIdentityRequest() (request *UpdateEmailIdentityRequest) {
    request = &UpdateEmailIdentityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ses", APIVersion, "UpdateEmailIdentity")
    
    
    return
}

func NewUpdateEmailIdentityResponse() (response *UpdateEmailIdentityResponse) {
    response = &UpdateEmailIdentityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateEmailIdentity
// This API is used to verify whether your DNS configuration is correct.
//
// error code that may be returned:
//  FAILEDOPERATION_SERVICENOTAVAILABLE = "FailedOperation.ServiceNotAvailable"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NOTEXISTDOMAIN = "InvalidParameterValue.NotExistDomain"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) UpdateEmailIdentity(request *UpdateEmailIdentityRequest) (response *UpdateEmailIdentityResponse, err error) {
    return c.UpdateEmailIdentityWithContext(context.Background(), request)
}

// UpdateEmailIdentity
// This API is used to verify whether your DNS configuration is correct.
//
// error code that may be returned:
//  FAILEDOPERATION_SERVICENOTAVAILABLE = "FailedOperation.ServiceNotAvailable"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NOTEXISTDOMAIN = "InvalidParameterValue.NotExistDomain"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) UpdateEmailIdentityWithContext(ctx context.Context, request *UpdateEmailIdentityRequest) (response *UpdateEmailIdentityResponse, err error) {
    if request == nil {
        request = NewUpdateEmailIdentityRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ses", APIVersion, "UpdateEmailIdentity")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateEmailIdentity require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateEmailIdentityResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateEmailSmtpPassWordRequest() (request *UpdateEmailSmtpPassWordRequest) {
    request = &UpdateEmailSmtpPassWordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ses", APIVersion, "UpdateEmailSmtpPassWord")
    
    
    return
}

func NewUpdateEmailSmtpPassWordResponse() (response *UpdateEmailSmtpPassWordResponse) {
    response = &UpdateEmailSmtpPassWordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateEmailSmtpPassWord
// This API is used to set the SMTP password. Initially, no SMTP password is set for your email address, so emails cannot be sent over SMTP. To send emails over SMTP, you must set the SMTP password. The set password can be changed subsequently.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDSMTPPASSWORD = "InvalidParameterValue.InvalidSmtpPassWord"
//  INVALIDPARAMETERVALUE_NOSUCHSENDER = "InvalidParameterValue.NoSuchSender"
//  OPERATIONDENIED_REPEATPASSWORD = "OperationDenied.RepeatPassWord"
func (c *Client) UpdateEmailSmtpPassWord(request *UpdateEmailSmtpPassWordRequest) (response *UpdateEmailSmtpPassWordResponse, err error) {
    return c.UpdateEmailSmtpPassWordWithContext(context.Background(), request)
}

// UpdateEmailSmtpPassWord
// This API is used to set the SMTP password. Initially, no SMTP password is set for your email address, so emails cannot be sent over SMTP. To send emails over SMTP, you must set the SMTP password. The set password can be changed subsequently.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDSMTPPASSWORD = "InvalidParameterValue.InvalidSmtpPassWord"
//  INVALIDPARAMETERVALUE_NOSUCHSENDER = "InvalidParameterValue.NoSuchSender"
//  OPERATIONDENIED_REPEATPASSWORD = "OperationDenied.RepeatPassWord"
func (c *Client) UpdateEmailSmtpPassWordWithContext(ctx context.Context, request *UpdateEmailSmtpPassWordRequest) (response *UpdateEmailSmtpPassWordResponse, err error) {
    if request == nil {
        request = NewUpdateEmailSmtpPassWordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ses", APIVersion, "UpdateEmailSmtpPassWord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateEmailSmtpPassWord require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateEmailSmtpPassWordResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateEmailTemplateRequest() (request *UpdateEmailTemplateRequest) {
    request = &UpdateEmailTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ses", APIVersion, "UpdateEmailTemplate")
    
    
    return
}

func NewUpdateEmailTemplateResponse() (response *UpdateEmailTemplateResponse) {
    response = &UpdateEmailTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateEmailTemplate
// This API is used to update an email template. An updated template must be approved again before it can be used.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TEMPLATECONTENTISNULL = "InvalidParameterValue.TemplateContentIsNULL"
//  INVALIDPARAMETERVALUE_TEMPLATECONTENTISWRONG = "InvalidParameterValue.TemplateContentIsWrong"
//  INVALIDPARAMETERVALUE_TEMPLATENAMEILLEGAL = "InvalidParameterValue.TemplateNameIllegal"
//  INVALIDPARAMETERVALUE_TEMPLATENAMEISNULL = "InvalidParameterValue.TemplateNameIsNULL"
//  INVALIDPARAMETERVALUE_TEMPLATENOTEXIST = "InvalidParameterValue.TemplateNotExist"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) UpdateEmailTemplate(request *UpdateEmailTemplateRequest) (response *UpdateEmailTemplateResponse, err error) {
    return c.UpdateEmailTemplateWithContext(context.Background(), request)
}

// UpdateEmailTemplate
// This API is used to update an email template. An updated template must be approved again before it can be used.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TEMPLATECONTENTISNULL = "InvalidParameterValue.TemplateContentIsNULL"
//  INVALIDPARAMETERVALUE_TEMPLATECONTENTISWRONG = "InvalidParameterValue.TemplateContentIsWrong"
//  INVALIDPARAMETERVALUE_TEMPLATENAMEILLEGAL = "InvalidParameterValue.TemplateNameIllegal"
//  INVALIDPARAMETERVALUE_TEMPLATENAMEISNULL = "InvalidParameterValue.TemplateNameIsNULL"
//  INVALIDPARAMETERVALUE_TEMPLATENOTEXIST = "InvalidParameterValue.TemplateNotExist"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) UpdateEmailTemplateWithContext(ctx context.Context, request *UpdateEmailTemplateRequest) (response *UpdateEmailTemplateResponse, err error) {
    if request == nil {
        request = NewUpdateEmailTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ses", APIVersion, "UpdateEmailTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateEmailTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateEmailTemplateResponse()
    err = c.Send(request, response)
    return
}
