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

package v20180808

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2018-08-08"

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


func NewBatchModifyIntlDomainDNSRequest() (request *BatchModifyIntlDomainDNSRequest) {
    request = &BatchModifyIntlDomainDNSRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "BatchModifyIntlDomainDNS")
    
    
    return
}

func NewBatchModifyIntlDomainDNSResponse() (response *BatchModifyIntlDomainDNSResponse) {
    response = &BatchModifyIntlDomainDNSResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BatchModifyIntlDomainDNS
// This API is used to bulk modify DNS servers for domains.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SETDOMAINDNSFAILED = "FailedOperation.SetDomainDnsFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_NEEDLOGIN = "InternalError.NeedLogin"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINNAMEISINVALID = "InvalidParameter.DomainNameIsInvalid"
//  INVALIDPARAMETER_UPTO4000 = "InvalidParameter.UpTo4000"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
//  MISSINGPARAMETER_DOMAINISEMPTY = "MissingParameter.DomainIsEmpty"
//  MISSINGPARAMETER_REPDATAISNONE = "MissingParameter.RepDataIsNone"
//  RESOURCENOTFOUND_DOMAINNOTFOUND = "ResourceNotFound.DomainNotFound"
//  RESOURCEUNAVAILABLE_DOMAINISMODIFYINGDNS = "ResourceUnavailable.DomainIsModifyingDNS"
//  UNSUPPORTEDOPERATION_MODIFYDOMAININFOOPERATEUNSUPPORTED = "UnsupportedOperation.ModifyDomainInfoOperateUnsupported"
//  UNSUPPORTEDOPERATION_MODIFYDOMAININFOUNSUPPORTED = "UnsupportedOperation.ModifyDomainInfoUnsupported"
//  UNSUPPORTEDOPERATION_MODIFYDOMAINUNSUPPORTED = "UnsupportedOperation.ModifyDomainUnsupported"
func (c *Client) BatchModifyIntlDomainDNS(request *BatchModifyIntlDomainDNSRequest) (response *BatchModifyIntlDomainDNSResponse, err error) {
    return c.BatchModifyIntlDomainDNSWithContext(context.Background(), request)
}

// BatchModifyIntlDomainDNS
// This API is used to bulk modify DNS servers for domains.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SETDOMAINDNSFAILED = "FailedOperation.SetDomainDnsFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_NEEDLOGIN = "InternalError.NeedLogin"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINNAMEISINVALID = "InvalidParameter.DomainNameIsInvalid"
//  INVALIDPARAMETER_UPTO4000 = "InvalidParameter.UpTo4000"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
//  MISSINGPARAMETER_DOMAINISEMPTY = "MissingParameter.DomainIsEmpty"
//  MISSINGPARAMETER_REPDATAISNONE = "MissingParameter.RepDataIsNone"
//  RESOURCENOTFOUND_DOMAINNOTFOUND = "ResourceNotFound.DomainNotFound"
//  RESOURCEUNAVAILABLE_DOMAINISMODIFYINGDNS = "ResourceUnavailable.DomainIsModifyingDNS"
//  UNSUPPORTEDOPERATION_MODIFYDOMAININFOOPERATEUNSUPPORTED = "UnsupportedOperation.ModifyDomainInfoOperateUnsupported"
//  UNSUPPORTEDOPERATION_MODIFYDOMAININFOUNSUPPORTED = "UnsupportedOperation.ModifyDomainInfoUnsupported"
//  UNSUPPORTEDOPERATION_MODIFYDOMAINUNSUPPORTED = "UnsupportedOperation.ModifyDomainUnsupported"
func (c *Client) BatchModifyIntlDomainDNSWithContext(ctx context.Context, request *BatchModifyIntlDomainDNSRequest) (response *BatchModifyIntlDomainDNSResponse, err error) {
    if request == nil {
        request = NewBatchModifyIntlDomainDNSRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchModifyIntlDomainDNS require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchModifyIntlDomainDNSResponse()
    err = c.Send(request, response)
    return
}

func NewBatchModifyIntlDomainInfoRequest() (request *BatchModifyIntlDomainInfoRequest) {
    request = &BatchModifyIntlDomainInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "BatchModifyIntlDomainInfo")
    
    
    return
}

func NewBatchModifyIntlDomainInfoResponse() (response *BatchModifyIntlDomainInfoResponse) {
    response = &BatchModifyIntlDomainInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BatchModifyIntlDomainInfo
// This API is used to bulk modify registrant information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_NEEDLOGIN = "InternalError.NeedLogin"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UPTO4000 = "InvalidParameter.UpTo4000"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
//  MISSINGPARAMETER_DOMAINISEMPTY = "MissingParameter.DomainIsEmpty"
//  MISSINGPARAMETER_REPDATAISNONE = "MissingParameter.RepDataIsNone"
//  MISSINGPARAMETER_TEMPLATEIDISEMPTY = "MissingParameter.TemplateIdIsEmpty"
//  RESOURCEINSUFFICIENT_OVERWORK = "ResourceInsufficient.Overwork"
//  RESOURCENOTFOUND_APPROVEDTEMPLATENOTFOUND = "ResourceNotFound.ApprovedTemplateNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTFOUND = "ResourceNotFound.TemplateNotFound"
func (c *Client) BatchModifyIntlDomainInfo(request *BatchModifyIntlDomainInfoRequest) (response *BatchModifyIntlDomainInfoResponse, err error) {
    return c.BatchModifyIntlDomainInfoWithContext(context.Background(), request)
}

// BatchModifyIntlDomainInfo
// This API is used to bulk modify registrant information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_NEEDLOGIN = "InternalError.NeedLogin"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UPTO4000 = "InvalidParameter.UpTo4000"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
//  MISSINGPARAMETER_DOMAINISEMPTY = "MissingParameter.DomainIsEmpty"
//  MISSINGPARAMETER_REPDATAISNONE = "MissingParameter.RepDataIsNone"
//  MISSINGPARAMETER_TEMPLATEIDISEMPTY = "MissingParameter.TemplateIdIsEmpty"
//  RESOURCEINSUFFICIENT_OVERWORK = "ResourceInsufficient.Overwork"
//  RESOURCENOTFOUND_APPROVEDTEMPLATENOTFOUND = "ResourceNotFound.ApprovedTemplateNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTFOUND = "ResourceNotFound.TemplateNotFound"
func (c *Client) BatchModifyIntlDomainInfoWithContext(ctx context.Context, request *BatchModifyIntlDomainInfoRequest) (response *BatchModifyIntlDomainInfoResponse, err error) {
    if request == nil {
        request = NewBatchModifyIntlDomainInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchModifyIntlDomainInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchModifyIntlDomainInfoResponse()
    err = c.Send(request, response)
    return
}

func NewCheckIntlDomainNewRequest() (request *CheckIntlDomainNewRequest) {
    request = &CheckIntlDomainNewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "CheckIntlDomainNew")
    
    
    return
}

func NewCheckIntlDomainNewResponse() (response *CheckIntlDomainNewResponse) {
    response = &CheckIntlDomainNewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckIntlDomainNew
// This API is used to check whether a domain is available for registration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHECKDOMAINFAILED = "FailedOperation.CheckDomainFailed"
//  INTERNALERROR_CHECKDNAVAILABLEERR = "InternalError.CheckDnAvailableErr"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_NEEDLOGIN = "InternalError.NeedLogin"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINNAMEISINVALID = "InvalidParameter.DomainNameIsInvalid"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
//  MISSINGPARAMETER_DOMAINISEMPTY = "MissingParameter.DomainIsEmpty"
func (c *Client) CheckIntlDomainNew(request *CheckIntlDomainNewRequest) (response *CheckIntlDomainNewResponse, err error) {
    return c.CheckIntlDomainNewWithContext(context.Background(), request)
}

// CheckIntlDomainNew
// This API is used to check whether a domain is available for registration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHECKDOMAINFAILED = "FailedOperation.CheckDomainFailed"
//  INTERNALERROR_CHECKDNAVAILABLEERR = "InternalError.CheckDnAvailableErr"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_NEEDLOGIN = "InternalError.NeedLogin"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINNAMEISINVALID = "InvalidParameter.DomainNameIsInvalid"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
//  MISSINGPARAMETER_DOMAINISEMPTY = "MissingParameter.DomainIsEmpty"
func (c *Client) CheckIntlDomainNewWithContext(ctx context.Context, request *CheckIntlDomainNewRequest) (response *CheckIntlDomainNewResponse, err error) {
    if request == nil {
        request = NewCheckIntlDomainNewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckIntlDomainNew require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckIntlDomainNewResponse()
    err = c.Send(request, response)
    return
}

func NewCreateIntlDomainBatchRequest() (request *CreateIntlDomainBatchRequest) {
    request = &CreateIntlDomainBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "CreateIntlDomainBatch")
    
    
    return
}

func NewCreateIntlDomainBatchResponse() (response *CreateIntlDomainBatchResponse) {
    response = &CreateIntlDomainBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateIntlDomainBatch
// This API is used to bulk register domains.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REGISTERDOMAIN = "FailedOperation.RegisterDomain"
//  FAILEDOPERATION_REGISTERDOMAINFAILED = "FailedOperation.RegisterDomainFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_NEEDLOGIN = "InternalError.NeedLogin"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CUSTOMDNSNOTALLOWED = "InvalidParameter.CustomDnsNotAllowed"
//  INVALIDPARAMETER_PACKAGERESOURCEIDINVALID = "InvalidParameter.PackageResourceIdInvalid"
//  INVALIDPARAMETER_UPTO4000 = "InvalidParameter.UpTo4000"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
//  MISSINGPARAMETER_DOMAINISEMPTY = "MissingParameter.DomainIsEmpty"
//  MISSINGPARAMETER_REPDATAISNONE = "MissingParameter.RepDataIsNone"
//  MISSINGPARAMETER_TEMPLATEIDISEMPTY = "MissingParameter.TemplateIdIsEmpty"
//  RESOURCEINSUFFICIENT_OVERWORK = "ResourceInsufficient.Overwork"
//  RESOURCENOTFOUND_TEMPLATENOTFOUND = "ResourceNotFound.TemplateNotFound"
func (c *Client) CreateIntlDomainBatch(request *CreateIntlDomainBatchRequest) (response *CreateIntlDomainBatchResponse, err error) {
    return c.CreateIntlDomainBatchWithContext(context.Background(), request)
}

// CreateIntlDomainBatch
// This API is used to bulk register domains.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REGISTERDOMAIN = "FailedOperation.RegisterDomain"
//  FAILEDOPERATION_REGISTERDOMAINFAILED = "FailedOperation.RegisterDomainFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_NEEDLOGIN = "InternalError.NeedLogin"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CUSTOMDNSNOTALLOWED = "InvalidParameter.CustomDnsNotAllowed"
//  INVALIDPARAMETER_PACKAGERESOURCEIDINVALID = "InvalidParameter.PackageResourceIdInvalid"
//  INVALIDPARAMETER_UPTO4000 = "InvalidParameter.UpTo4000"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
//  MISSINGPARAMETER_DOMAINISEMPTY = "MissingParameter.DomainIsEmpty"
//  MISSINGPARAMETER_REPDATAISNONE = "MissingParameter.RepDataIsNone"
//  MISSINGPARAMETER_TEMPLATEIDISEMPTY = "MissingParameter.TemplateIdIsEmpty"
//  RESOURCEINSUFFICIENT_OVERWORK = "ResourceInsufficient.Overwork"
//  RESOURCENOTFOUND_TEMPLATENOTFOUND = "ResourceNotFound.TemplateNotFound"
func (c *Client) CreateIntlDomainBatchWithContext(ctx context.Context, request *CreateIntlDomainBatchRequest) (response *CreateIntlDomainBatchResponse, err error) {
    if request == nil {
        request = NewCreateIntlDomainBatchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateIntlDomainBatch require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateIntlDomainBatchResponse()
    err = c.Send(request, response)
    return
}

func NewCreateIntlPhoneEmailRequest() (request *CreateIntlPhoneEmailRequest) {
    request = &CreateIntlPhoneEmailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "CreateIntlPhoneEmail")
    
    
    return
}

func NewCreateIntlPhoneEmailResponse() (response *CreateIntlPhoneEmailResponse) {
    response = &CreateIntlPhoneEmailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateIntlPhoneEmail
// This API is used to verify a mobile number or an email address via a verification code.
//
// error code that may be returned:
//  INTERNALERROR_ADDPHONEEMAILERR = "InternalError.AddPhoneEmailErr"
//  INTERNALERROR_CHECKVERIFYCODEERR = "InternalError.CheckVerifyCodeErr"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_NEEDLOGIN = "InternalError.NeedLogin"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INTERNALERROR_VERIFIEDPHONEEMAILERR = "InternalError.VerifiedPhoneEmailErr"
//  INTERNALERROR_VERIFYCODEERR = "InternalError.VerifyCodeErr"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
func (c *Client) CreateIntlPhoneEmail(request *CreateIntlPhoneEmailRequest) (response *CreateIntlPhoneEmailResponse, err error) {
    return c.CreateIntlPhoneEmailWithContext(context.Background(), request)
}

// CreateIntlPhoneEmail
// This API is used to verify a mobile number or an email address via a verification code.
//
// error code that may be returned:
//  INTERNALERROR_ADDPHONEEMAILERR = "InternalError.AddPhoneEmailErr"
//  INTERNALERROR_CHECKVERIFYCODEERR = "InternalError.CheckVerifyCodeErr"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_NEEDLOGIN = "InternalError.NeedLogin"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INTERNALERROR_VERIFIEDPHONEEMAILERR = "InternalError.VerifiedPhoneEmailErr"
//  INTERNALERROR_VERIFYCODEERR = "InternalError.VerifyCodeErr"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
func (c *Client) CreateIntlPhoneEmailWithContext(ctx context.Context, request *CreateIntlPhoneEmailRequest) (response *CreateIntlPhoneEmailResponse, err error) {
    if request == nil {
        request = NewCreateIntlPhoneEmailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateIntlPhoneEmail require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateIntlPhoneEmailResponse()
    err = c.Send(request, response)
    return
}

func NewCreateIntlTemplateRequest() (request *CreateIntlTemplateRequest) {
    request = &CreateIntlTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "CreateIntlTemplate")
    
    
    return
}

func NewCreateIntlTemplateResponse() (response *CreateIntlTemplateResponse) {
    response = &CreateIntlTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateIntlTemplate
// This API is used to add a registrant profile.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATETEMPLATEFAILED = "FailedOperation.CreateTemplateFailed"
//  FAILEDOPERATION_DESCRIBEDOMAINLISTFAILED = "FailedOperation.DescribeDomainListFailed"
//  FAILEDOPERATION_DESCRIBETEMPLATEFAILED = "FailedOperation.DescribeTemplateFailed"
//  FAILEDOPERATION_TEMPLATEMAXNUMFAILED = "FailedOperation.TemplateMaxNumFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_NEEDLOGIN = "InternalError.NeedLogin"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CERTIFICATECODEISINVALID = "InvalidParameter.CertificateCodeIsInvalid"
//  INVALIDPARAMETER_CERTIFICATEIMAGEISINVALID = "InvalidParameter.CertificateImageIsInvalid"
//  INVALIDPARAMETER_EMAILISINVALID = "InvalidParameter.EmailIsInvalid"
//  INVALIDPARAMETER_FIRSTNAMEISINVALID = "InvalidParameter.FirstNameIsInvalid"
//  INVALIDPARAMETER_LASTNAMEISINVALID = "InvalidParameter.LastNameIsInvalid"
//  INVALIDPARAMETER_NAMEISINVALID = "InvalidParameter.NameIsInvalid"
//  INVALIDPARAMETER_ORGISINVALID = "InvalidParameter.OrgIsInvalid"
//  INVALIDPARAMETER_REPTYPEISINVALID = "InvalidParameter.RepTypeIsInvalid"
//  INVALIDPARAMETER_STREETISINVALID = "InvalidParameter.StreetIsInvalid"
//  INVALIDPARAMETER_TELEPHONEISINVALID = "InvalidParameter.TelephoneIsInvalid"
//  INVALIDPARAMETER_USERTYPEISINVALID = "InvalidParameter.UserTypeIsInvalid"
//  INVALIDPARAMETER_ZIPCODEISINVALID = "InvalidParameter.ZipCodeIsInvalid"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
//  MISSINGPARAMETER_DOMAINISEMPTY = "MissingParameter.DomainIsEmpty"
//  MISSINGPARAMETER_REPDATAISNONE = "MissingParameter.RepDataIsNone"
//  MISSINGPARAMETER_TEMPLATEIDISEMPTY = "MissingParameter.TemplateIdIsEmpty"
//  MISSINGPARAMETER_TEMPLATEIDISEXIST = "MissingParameter.TemplateIdIsExist"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTFOUND = "ResourceNotFound.TemplateNotFound"
func (c *Client) CreateIntlTemplate(request *CreateIntlTemplateRequest) (response *CreateIntlTemplateResponse, err error) {
    return c.CreateIntlTemplateWithContext(context.Background(), request)
}

// CreateIntlTemplate
// This API is used to add a registrant profile.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATETEMPLATEFAILED = "FailedOperation.CreateTemplateFailed"
//  FAILEDOPERATION_DESCRIBEDOMAINLISTFAILED = "FailedOperation.DescribeDomainListFailed"
//  FAILEDOPERATION_DESCRIBETEMPLATEFAILED = "FailedOperation.DescribeTemplateFailed"
//  FAILEDOPERATION_TEMPLATEMAXNUMFAILED = "FailedOperation.TemplateMaxNumFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_NEEDLOGIN = "InternalError.NeedLogin"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CERTIFICATECODEISINVALID = "InvalidParameter.CertificateCodeIsInvalid"
//  INVALIDPARAMETER_CERTIFICATEIMAGEISINVALID = "InvalidParameter.CertificateImageIsInvalid"
//  INVALIDPARAMETER_EMAILISINVALID = "InvalidParameter.EmailIsInvalid"
//  INVALIDPARAMETER_FIRSTNAMEISINVALID = "InvalidParameter.FirstNameIsInvalid"
//  INVALIDPARAMETER_LASTNAMEISINVALID = "InvalidParameter.LastNameIsInvalid"
//  INVALIDPARAMETER_NAMEISINVALID = "InvalidParameter.NameIsInvalid"
//  INVALIDPARAMETER_ORGISINVALID = "InvalidParameter.OrgIsInvalid"
//  INVALIDPARAMETER_REPTYPEISINVALID = "InvalidParameter.RepTypeIsInvalid"
//  INVALIDPARAMETER_STREETISINVALID = "InvalidParameter.StreetIsInvalid"
//  INVALIDPARAMETER_TELEPHONEISINVALID = "InvalidParameter.TelephoneIsInvalid"
//  INVALIDPARAMETER_USERTYPEISINVALID = "InvalidParameter.UserTypeIsInvalid"
//  INVALIDPARAMETER_ZIPCODEISINVALID = "InvalidParameter.ZipCodeIsInvalid"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
//  MISSINGPARAMETER_DOMAINISEMPTY = "MissingParameter.DomainIsEmpty"
//  MISSINGPARAMETER_REPDATAISNONE = "MissingParameter.RepDataIsNone"
//  MISSINGPARAMETER_TEMPLATEIDISEMPTY = "MissingParameter.TemplateIdIsEmpty"
//  MISSINGPARAMETER_TEMPLATEIDISEXIST = "MissingParameter.TemplateIdIsExist"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTFOUND = "ResourceNotFound.TemplateNotFound"
func (c *Client) CreateIntlTemplateWithContext(ctx context.Context, request *CreateIntlTemplateRequest) (response *CreateIntlTemplateResponse, err error) {
    if request == nil {
        request = NewCreateIntlTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateIntlTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateIntlTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteIntlPhoneEmailRequest() (request *DeleteIntlPhoneEmailRequest) {
    request = &DeleteIntlPhoneEmailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "DeleteIntlPhoneEmail")
    
    
    return
}

func NewDeleteIntlPhoneEmailResponse() (response *DeleteIntlPhoneEmailResponse) {
    response = &DeleteIntlPhoneEmailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteIntlPhoneEmail
// This API is used to delete a mobile number or an email address.
//
// error code that may be returned:
//  INTERNALERROR_DELETEPHONEEMAILVERIFYERR = "InternalError.DeletePhoneEmailVerifyErr"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_NEEDLOGIN = "InternalError.NeedLogin"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
func (c *Client) DeleteIntlPhoneEmail(request *DeleteIntlPhoneEmailRequest) (response *DeleteIntlPhoneEmailResponse, err error) {
    return c.DeleteIntlPhoneEmailWithContext(context.Background(), request)
}

// DeleteIntlPhoneEmail
// This API is used to delete a mobile number or an email address.
//
// error code that may be returned:
//  INTERNALERROR_DELETEPHONEEMAILVERIFYERR = "InternalError.DeletePhoneEmailVerifyErr"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_NEEDLOGIN = "InternalError.NeedLogin"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
func (c *Client) DeleteIntlPhoneEmailWithContext(ctx context.Context, request *DeleteIntlPhoneEmailRequest) (response *DeleteIntlPhoneEmailResponse, err error) {
    if request == nil {
        request = NewDeleteIntlPhoneEmailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteIntlPhoneEmail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteIntlPhoneEmailResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteIntlTemplateRequest() (request *DeleteIntlTemplateRequest) {
    request = &DeleteIntlTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "DeleteIntlTemplate")
    
    
    return
}

func NewDeleteIntlTemplateResponse() (response *DeleteIntlTemplateResponse) {
    response = &DeleteIntlTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteIntlTemplate
// This API is used to delete a registrant profile.
//
// error code that may be returned:
//  INTERNALERROR_DELETETEMPLATEERR = "InternalError.DeleteTemplateErr"
//  INTERNALERROR_DESCRIBETEMPLATEERR = "InternalError.DescribeTemplateErr"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
func (c *Client) DeleteIntlTemplate(request *DeleteIntlTemplateRequest) (response *DeleteIntlTemplateResponse, err error) {
    return c.DeleteIntlTemplateWithContext(context.Background(), request)
}

// DeleteIntlTemplate
// This API is used to delete a registrant profile.
//
// error code that may be returned:
//  INTERNALERROR_DELETETEMPLATEERR = "InternalError.DeleteTemplateErr"
//  INTERNALERROR_DESCRIBETEMPLATEERR = "InternalError.DescribeTemplateErr"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
func (c *Client) DeleteIntlTemplateWithContext(ctx context.Context, request *DeleteIntlTemplateRequest) (response *DeleteIntlTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteIntlTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteIntlTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteIntlTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIntlBatchDetailStatusRequest() (request *DescribeIntlBatchDetailStatusRequest) {
    request = &DescribeIntlBatchDetailStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "DescribeIntlBatchDetailStatus")
    
    
    return
}

func NewDescribeIntlBatchDetailStatusResponse() (response *DescribeIntlBatchDetailStatusResponse) {
    response = &DescribeIntlBatchDetailStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIntlBatchDetailStatus
// This API is used to query the status of a bulk task.
//
// error code that may be returned:
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_NEEDLOGIN = "InternalError.NeedLogin"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
func (c *Client) DescribeIntlBatchDetailStatus(request *DescribeIntlBatchDetailStatusRequest) (response *DescribeIntlBatchDetailStatusResponse, err error) {
    return c.DescribeIntlBatchDetailStatusWithContext(context.Background(), request)
}

// DescribeIntlBatchDetailStatus
// This API is used to query the status of a bulk task.
//
// error code that may be returned:
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_NEEDLOGIN = "InternalError.NeedLogin"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
func (c *Client) DescribeIntlBatchDetailStatusWithContext(ctx context.Context, request *DescribeIntlBatchDetailStatusRequest) (response *DescribeIntlBatchDetailStatusResponse, err error) {
    if request == nil {
        request = NewDescribeIntlBatchDetailStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIntlBatchDetailStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIntlBatchDetailStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIntlBatchOperationLogsRequest() (request *DescribeIntlBatchOperationLogsRequest) {
    request = &DescribeIntlBatchOperationLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "DescribeIntlBatchOperationLogs")
    
    
    return
}

func NewDescribeIntlBatchOperationLogsResponse() (response *DescribeIntlBatchOperationLogsResponse) {
    response = &DescribeIntlBatchOperationLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIntlBatchOperationLogs
// This API is used to query the logs of bulk domain purchase.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_NEEDLOGIN = "InternalError.NeedLogin"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
func (c *Client) DescribeIntlBatchOperationLogs(request *DescribeIntlBatchOperationLogsRequest) (response *DescribeIntlBatchOperationLogsResponse, err error) {
    return c.DescribeIntlBatchOperationLogsWithContext(context.Background(), request)
}

// DescribeIntlBatchOperationLogs
// This API is used to query the logs of bulk domain purchase.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_NEEDLOGIN = "InternalError.NeedLogin"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
func (c *Client) DescribeIntlBatchOperationLogsWithContext(ctx context.Context, request *DescribeIntlBatchOperationLogsRequest) (response *DescribeIntlBatchOperationLogsResponse, err error) {
    if request == nil {
        request = NewDescribeIntlBatchOperationLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIntlBatchOperationLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIntlBatchOperationLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIntlDomainRequest() (request *DescribeIntlDomainRequest) {
    request = &DescribeIntlDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "DescribeIntlDomain")
    
    
    return
}

func NewDescribeIntlDomainResponse() (response *DescribeIntlDomainResponse) {
    response = &DescribeIntlDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIntlDomain
// This API is used to query domain information.
//
// error code that may be returned:
//  INTERNALERROR_DESCRIBEDOMAININFOERR = "InternalError.DescribeDomainInfoErr"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
func (c *Client) DescribeIntlDomain(request *DescribeIntlDomainRequest) (response *DescribeIntlDomainResponse, err error) {
    return c.DescribeIntlDomainWithContext(context.Background(), request)
}

// DescribeIntlDomain
// This API is used to query domain information.
//
// error code that may be returned:
//  INTERNALERROR_DESCRIBEDOMAININFOERR = "InternalError.DescribeDomainInfoErr"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
func (c *Client) DescribeIntlDomainWithContext(ctx context.Context, request *DescribeIntlDomainRequest) (response *DescribeIntlDomainResponse, err error) {
    if request == nil {
        request = NewDescribeIntlDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIntlDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIntlDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIntlDomainBatchDetailsRequest() (request *DescribeIntlDomainBatchDetailsRequest) {
    request = &DescribeIntlDomainBatchDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "DescribeIntlDomainBatchDetails")
    
    
    return
}

func NewDescribeIntlDomainBatchDetailsResponse() (response *DescribeIntlDomainBatchDetailsResponse) {
    response = &DescribeIntlDomainBatchDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIntlDomainBatchDetails
// This API is used to get the log details of bulk domain purchase.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeIntlDomainBatchDetails(request *DescribeIntlDomainBatchDetailsRequest) (response *DescribeIntlDomainBatchDetailsResponse, err error) {
    return c.DescribeIntlDomainBatchDetailsWithContext(context.Background(), request)
}

// DescribeIntlDomainBatchDetails
// This API is used to get the log details of bulk domain purchase.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeIntlDomainBatchDetailsWithContext(ctx context.Context, request *DescribeIntlDomainBatchDetailsRequest) (response *DescribeIntlDomainBatchDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeIntlDomainBatchDetailsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIntlDomainBatchDetails require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIntlDomainBatchDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIntlDomainListRequest() (request *DescribeIntlDomainListRequest) {
    request = &DescribeIntlDomainListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "DescribeIntlDomainList")
    
    
    return
}

func NewDescribeIntlDomainListResponse() (response *DescribeIntlDomainListResponse) {
    response = &DescribeIntlDomainListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIntlDomainList
// This API is used to query the "My domains" list.
//
// error code that may be returned:
//  INTERNALERROR_DESCRIBEDOMAINLISTERR = "InternalError.DescribeDomainListErr"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
func (c *Client) DescribeIntlDomainList(request *DescribeIntlDomainListRequest) (response *DescribeIntlDomainListResponse, err error) {
    return c.DescribeIntlDomainListWithContext(context.Background(), request)
}

// DescribeIntlDomainList
// This API is used to query the "My domains" list.
//
// error code that may be returned:
//  INTERNALERROR_DESCRIBEDOMAINLISTERR = "InternalError.DescribeDomainListErr"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
func (c *Client) DescribeIntlDomainListWithContext(ctx context.Context, request *DescribeIntlDomainListRequest) (response *DescribeIntlDomainListResponse, err error) {
    if request == nil {
        request = NewDescribeIntlDomainListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIntlDomainList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIntlDomainListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIntlDomainPriceNewListRequest() (request *DescribeIntlDomainPriceNewListRequest) {
    request = &DescribeIntlDomainPriceNewListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "DescribeIntlDomainPriceNewList")
    
    
    return
}

func NewDescribeIntlDomainPriceNewListResponse() (response *DescribeIntlDomainPriceNewListResponse) {
    response = &DescribeIntlDomainPriceNewListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIntlDomainPriceNewList
// This API is used to get the price list by domain suffix.
//
// error code that may be returned:
//  FAILEDOPERATION_DOMAINPRICELISTFAILED = "FailedOperation.DomainPriceListFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
func (c *Client) DescribeIntlDomainPriceNewList(request *DescribeIntlDomainPriceNewListRequest) (response *DescribeIntlDomainPriceNewListResponse, err error) {
    return c.DescribeIntlDomainPriceNewListWithContext(context.Background(), request)
}

// DescribeIntlDomainPriceNewList
// This API is used to get the price list by domain suffix.
//
// error code that may be returned:
//  FAILEDOPERATION_DOMAINPRICELISTFAILED = "FailedOperation.DomainPriceListFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
func (c *Client) DescribeIntlDomainPriceNewListWithContext(ctx context.Context, request *DescribeIntlDomainPriceNewListRequest) (response *DescribeIntlDomainPriceNewListResponse, err error) {
    if request == nil {
        request = NewDescribeIntlDomainPriceNewListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIntlDomainPriceNewList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIntlDomainPriceNewListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIntlPhoneEmailListRequest() (request *DescribeIntlPhoneEmailListRequest) {
    request = &DescribeIntlPhoneEmailListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "DescribeIntlPhoneEmailList")
    
    
    return
}

func NewDescribeIntlPhoneEmailListResponse() (response *DescribeIntlPhoneEmailListResponse) {
    response = &DescribeIntlPhoneEmailListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIntlPhoneEmailList
// This API is used to get the list of verified mobile numbers and email addresses.
//
// error code that may be returned:
//  INTERNALERROR_DESCRIBEPHONEEMAILLISTERR = "InternalError.DescribePhoneEmailListErr"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_NEEDLOGIN = "InternalError.NeedLogin"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
func (c *Client) DescribeIntlPhoneEmailList(request *DescribeIntlPhoneEmailListRequest) (response *DescribeIntlPhoneEmailListResponse, err error) {
    return c.DescribeIntlPhoneEmailListWithContext(context.Background(), request)
}

// DescribeIntlPhoneEmailList
// This API is used to get the list of verified mobile numbers and email addresses.
//
// error code that may be returned:
//  INTERNALERROR_DESCRIBEPHONEEMAILLISTERR = "InternalError.DescribePhoneEmailListErr"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_NEEDLOGIN = "InternalError.NeedLogin"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
func (c *Client) DescribeIntlPhoneEmailListWithContext(ctx context.Context, request *DescribeIntlPhoneEmailListRequest) (response *DescribeIntlPhoneEmailListResponse, err error) {
    if request == nil {
        request = NewDescribeIntlPhoneEmailListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIntlPhoneEmailList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIntlPhoneEmailListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIntlTemplateRequest() (request *DescribeIntlTemplateRequest) {
    request = &DescribeIntlTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "DescribeIntlTemplate")
    
    
    return
}

func NewDescribeIntlTemplateResponse() (response *DescribeIntlTemplateResponse) {
    response = &DescribeIntlTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIntlTemplate
// This API is used to get the details of a registrant profile.
//
// error code that may be returned:
//  INTERNALERROR_DESCRIBETEMPLATEERR = "InternalError.DescribeTemplateErr"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
func (c *Client) DescribeIntlTemplate(request *DescribeIntlTemplateRequest) (response *DescribeIntlTemplateResponse, err error) {
    return c.DescribeIntlTemplateWithContext(context.Background(), request)
}

// DescribeIntlTemplate
// This API is used to get the details of a registrant profile.
//
// error code that may be returned:
//  INTERNALERROR_DESCRIBETEMPLATEERR = "InternalError.DescribeTemplateErr"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
func (c *Client) DescribeIntlTemplateWithContext(ctx context.Context, request *DescribeIntlTemplateRequest) (response *DescribeIntlTemplateResponse, err error) {
    if request == nil {
        request = NewDescribeIntlTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIntlTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIntlTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIntlTemplateListRequest() (request *DescribeIntlTemplateListRequest) {
    request = &DescribeIntlTemplateListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "DescribeIntlTemplateList")
    
    
    return
}

func NewDescribeIntlTemplateListResponse() (response *DescribeIntlTemplateListResponse) {
    response = &DescribeIntlTemplateListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIntlTemplateList
// This API is used to get the list of registrant profiles.
//
// error code that may be returned:
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeIntlTemplateList(request *DescribeIntlTemplateListRequest) (response *DescribeIntlTemplateListResponse, err error) {
    return c.DescribeIntlTemplateListWithContext(context.Background(), request)
}

// DescribeIntlTemplateList
// This API is used to get the list of registrant profiles.
//
// error code that may be returned:
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeIntlTemplateListWithContext(ctx context.Context, request *DescribeIntlTemplateListRequest) (response *DescribeIntlTemplateListResponse, err error) {
    if request == nil {
        request = NewDescribeIntlTemplateListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIntlTemplateList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIntlTemplateListResponse()
    err = c.Send(request, response)
    return
}

func NewModifyOwnerIntlBatchRequest() (request *ModifyOwnerIntlBatchRequest) {
    request = &ModifyOwnerIntlBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "ModifyOwnerIntlBatch")
    
    
    return
}

func NewModifyOwnerIntlBatchResponse() (response *ModifyOwnerIntlBatchResponse) {
    response = &ModifyOwnerIntlBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyOwnerIntlBatch
// This API is used to bulk transfer domains to another account.
//
// error code that may be returned:
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INTERNALERROR_DOMAINTRANSFERERR = "InternalError.DomainTransferErr"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_NEEDLOGIN = "InternalError.NeedLogin"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UPTO4000 = "InvalidParameter.UpTo4000"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
//  RESOURCEINSUFFICIENT_OVERWORK = "ResourceInsufficient.Overwork"
func (c *Client) ModifyOwnerIntlBatch(request *ModifyOwnerIntlBatchRequest) (response *ModifyOwnerIntlBatchResponse, err error) {
    return c.ModifyOwnerIntlBatchWithContext(context.Background(), request)
}

// ModifyOwnerIntlBatch
// This API is used to bulk transfer domains to another account.
//
// error code that may be returned:
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INTERNALERROR_DOMAINTRANSFERERR = "InternalError.DomainTransferErr"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_NEEDLOGIN = "InternalError.NeedLogin"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UPTO4000 = "InvalidParameter.UpTo4000"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
//  RESOURCEINSUFFICIENT_OVERWORK = "ResourceInsufficient.Overwork"
func (c *Client) ModifyOwnerIntlBatchWithContext(ctx context.Context, request *ModifyOwnerIntlBatchRequest) (response *ModifyOwnerIntlBatchResponse, err error) {
    if request == nil {
        request = NewModifyOwnerIntlBatchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyOwnerIntlBatch require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyOwnerIntlBatchResponse()
    err = c.Send(request, response)
    return
}

func NewRenewIntlDomainBatchRequest() (request *RenewIntlDomainBatchRequest) {
    request = &RenewIntlDomainBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "RenewIntlDomainBatch")
    
    
    return
}

func NewRenewIntlDomainBatchResponse() (response *RenewIntlDomainBatchResponse) {
    response = &RenewIntlDomainBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RenewIntlDomainBatch
// This API is used to bulk renew domains.
//
// error code that may be returned:
//  FAILEDOPERATION_RENEWDOMAINFAILED = "FailedOperation.RenewDomainFailed"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_NEEDLOGIN = "InternalError.NeedLogin"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINNAMEISINVALID = "InvalidParameter.DomainNameIsInvalid"
//  INVALIDPARAMETER_UPTO4000 = "InvalidParameter.UpTo4000"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
//  RESOURCEINSUFFICIENT_OVERWORK = "ResourceInsufficient.Overwork"
func (c *Client) RenewIntlDomainBatch(request *RenewIntlDomainBatchRequest) (response *RenewIntlDomainBatchResponse, err error) {
    return c.RenewIntlDomainBatchWithContext(context.Background(), request)
}

// RenewIntlDomainBatch
// This API is used to bulk renew domains.
//
// error code that may be returned:
//  FAILEDOPERATION_RENEWDOMAINFAILED = "FailedOperation.RenewDomainFailed"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_NEEDLOGIN = "InternalError.NeedLogin"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINNAMEISINVALID = "InvalidParameter.DomainNameIsInvalid"
//  INVALIDPARAMETER_UPTO4000 = "InvalidParameter.UpTo4000"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
//  RESOURCEINSUFFICIENT_OVERWORK = "ResourceInsufficient.Overwork"
func (c *Client) RenewIntlDomainBatchWithContext(ctx context.Context, request *RenewIntlDomainBatchRequest) (response *RenewIntlDomainBatchResponse, err error) {
    if request == nil {
        request = NewRenewIntlDomainBatchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RenewIntlDomainBatch require credential")
    }

    request.SetContext(ctx)
    
    response = NewRenewIntlDomainBatchResponse()
    err = c.Send(request, response)
    return
}

func NewSendIntlPhoneEmailCodeRequest() (request *SendIntlPhoneEmailCodeRequest) {
    request = &SendIntlPhoneEmailCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "SendIntlPhoneEmailCode")
    
    
    return
}

func NewSendIntlPhoneEmailCodeResponse() (response *SendIntlPhoneEmailCodeResponse) {
    response = &SendIntlPhoneEmailCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SendIntlPhoneEmailCode
// This API is used to send a verification code to a mobile number or an email address.
//
// error code that may be returned:
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_NEEDLOGIN = "InternalError.NeedLogin"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INTERNALERROR_SENDVERIFYCODEERR = "InternalError.SendVerifyCodeErr"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
func (c *Client) SendIntlPhoneEmailCode(request *SendIntlPhoneEmailCodeRequest) (response *SendIntlPhoneEmailCodeResponse, err error) {
    return c.SendIntlPhoneEmailCodeWithContext(context.Background(), request)
}

// SendIntlPhoneEmailCode
// This API is used to send a verification code to a mobile number or an email address.
//
// error code that may be returned:
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_NEEDLOGIN = "InternalError.NeedLogin"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INTERNALERROR_SENDVERIFYCODEERR = "InternalError.SendVerifyCodeErr"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
func (c *Client) SendIntlPhoneEmailCodeWithContext(ctx context.Context, request *SendIntlPhoneEmailCodeRequest) (response *SendIntlPhoneEmailCodeResponse, err error) {
    if request == nil {
        request = NewSendIntlPhoneEmailCodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SendIntlPhoneEmailCode require credential")
    }

    request.SetContext(ctx)
    
    response = NewSendIntlPhoneEmailCodeResponse()
    err = c.Send(request, response)
    return
}

func NewSetIntlDomainAutoRenewRequest() (request *SetIntlDomainAutoRenewRequest) {
    request = &SetIntlDomainAutoRenewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "SetIntlDomainAutoRenew")
    
    
    return
}

func NewSetIntlDomainAutoRenewResponse() (response *SetIntlDomainAutoRenewResponse) {
    response = &SetIntlDomainAutoRenewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetIntlDomainAutoRenew
// This API is used to set auto-renewal.
//
// error code that may be returned:
//  INTERNALERROR_DESCRIBEDOMAININFOERR = "InternalError.DescribeDomainInfoErr"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INTERNALERROR_SETDOMAINAUTORENEWERR = "InternalError.SetDomainAutoRenewErr"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
func (c *Client) SetIntlDomainAutoRenew(request *SetIntlDomainAutoRenewRequest) (response *SetIntlDomainAutoRenewResponse, err error) {
    return c.SetIntlDomainAutoRenewWithContext(context.Background(), request)
}

// SetIntlDomainAutoRenew
// This API is used to set auto-renewal.
//
// error code that may be returned:
//  INTERNALERROR_DESCRIBEDOMAININFOERR = "InternalError.DescribeDomainInfoErr"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INTERNALERROR_SETDOMAINAUTORENEWERR = "InternalError.SetDomainAutoRenewErr"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
func (c *Client) SetIntlDomainAutoRenewWithContext(ctx context.Context, request *SetIntlDomainAutoRenewRequest) (response *SetIntlDomainAutoRenewResponse, err error) {
    if request == nil {
        request = NewSetIntlDomainAutoRenewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetIntlDomainAutoRenew require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetIntlDomainAutoRenewResponse()
    err = c.Send(request, response)
    return
}

func NewTransferInIntlDomainBatchRequest() (request *TransferInIntlDomainBatchRequest) {
    request = &TransferInIntlDomainBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "TransferInIntlDomainBatch")
    
    
    return
}

func NewTransferInIntlDomainBatchResponse() (response *TransferInIntlDomainBatchResponse) {
    response = &TransferInIntlDomainBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TransferInIntlDomainBatch
// This API is used to bulk transfer domains in.
//
// error code that may be returned:
//  FAILEDOPERATION_TRANSFERINDOMAINFAILED = "FailedOperation.TransferInDomainFailed"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_NEEDLOGIN = "InternalError.NeedLogin"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UPTO4000 = "InvalidParameter.UpTo4000"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
//  RESOURCEINSUFFICIENT_OVERWORK = "ResourceInsufficient.Overwork"
//  RESOURCENOTFOUND_TEMPLATENOTFOUND = "ResourceNotFound.TemplateNotFound"
func (c *Client) TransferInIntlDomainBatch(request *TransferInIntlDomainBatchRequest) (response *TransferInIntlDomainBatchResponse, err error) {
    return c.TransferInIntlDomainBatchWithContext(context.Background(), request)
}

// TransferInIntlDomainBatch
// This API is used to bulk transfer domains in.
//
// error code that may be returned:
//  FAILEDOPERATION_TRANSFERINDOMAINFAILED = "FailedOperation.TransferInDomainFailed"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_NEEDLOGIN = "InternalError.NeedLogin"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UPTO4000 = "InvalidParameter.UpTo4000"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
//  RESOURCEINSUFFICIENT_OVERWORK = "ResourceInsufficient.Overwork"
//  RESOURCENOTFOUND_TEMPLATENOTFOUND = "ResourceNotFound.TemplateNotFound"
func (c *Client) TransferInIntlDomainBatchWithContext(ctx context.Context, request *TransferInIntlDomainBatchRequest) (response *TransferInIntlDomainBatchResponse, err error) {
    if request == nil {
        request = NewTransferInIntlDomainBatchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TransferInIntlDomainBatch require credential")
    }

    request.SetContext(ctx)
    
    response = NewTransferInIntlDomainBatchResponse()
    err = c.Send(request, response)
    return
}

func NewTransferProhibitionIntlBatchRequest() (request *TransferProhibitionIntlBatchRequest) {
    request = &TransferProhibitionIntlBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "TransferProhibitionIntlBatch")
    
    
    return
}

func NewTransferProhibitionIntlBatchResponse() (response *TransferProhibitionIntlBatchResponse) {
    response = &TransferProhibitionIntlBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TransferProhibitionIntlBatch
// This API is used to bulk set transfer prohibition for domains.
//
// error code that may be returned:
//  INTERNALERROR_DESCRIBEDOMAININFOERR = "InternalError.DescribeDomainInfoErr"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INTERNALERROR_SETTRANSFERPROHIBITEDERR = "InternalError.SetTransferProhibitedErr"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
//  RESOURCEUNAVAILABLE_DOMAINISMODIFYINGDNS = "ResourceUnavailable.DomainIsModifyingDNS"
func (c *Client) TransferProhibitionIntlBatch(request *TransferProhibitionIntlBatchRequest) (response *TransferProhibitionIntlBatchResponse, err error) {
    return c.TransferProhibitionIntlBatchWithContext(context.Background(), request)
}

// TransferProhibitionIntlBatch
// This API is used to bulk set transfer prohibition for domains.
//
// error code that may be returned:
//  INTERNALERROR_DESCRIBEDOMAININFOERR = "InternalError.DescribeDomainInfoErr"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INTERNALERROR_SETTRANSFERPROHIBITEDERR = "InternalError.SetTransferProhibitedErr"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
//  RESOURCEUNAVAILABLE_DOMAINISMODIFYINGDNS = "ResourceUnavailable.DomainIsModifyingDNS"
func (c *Client) TransferProhibitionIntlBatchWithContext(ctx context.Context, request *TransferProhibitionIntlBatchRequest) (response *TransferProhibitionIntlBatchResponse, err error) {
    if request == nil {
        request = NewTransferProhibitionIntlBatchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TransferProhibitionIntlBatch require credential")
    }

    request.SetContext(ctx)
    
    response = NewTransferProhibitionIntlBatchResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateProhibitionIntlBatchRequest() (request *UpdateProhibitionIntlBatchRequest) {
    request = &UpdateProhibitionIntlBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "UpdateProhibitionIntlBatch")
    
    
    return
}

func NewUpdateProhibitionIntlBatchResponse() (response *UpdateProhibitionIntlBatchResponse) {
    response = &UpdateProhibitionIntlBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateProhibitionIntlBatch
// This API is used to bulk set update prohibition for domains.
//
// error code that may be returned:
//  INTERNALERROR_DESCRIBEDOMAININFOERR = "InternalError.DescribeDomainInfoErr"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INTERNALERROR_SETUPDATEPROHIBITEDERR = "InternalError.SetUpdateProhibitedErr"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
//  RESOURCEUNAVAILABLE_DOMAINISMODIFYINGDNS = "ResourceUnavailable.DomainIsModifyingDNS"
func (c *Client) UpdateProhibitionIntlBatch(request *UpdateProhibitionIntlBatchRequest) (response *UpdateProhibitionIntlBatchResponse, err error) {
    return c.UpdateProhibitionIntlBatchWithContext(context.Background(), request)
}

// UpdateProhibitionIntlBatch
// This API is used to bulk set update prohibition for domains.
//
// error code that may be returned:
//  INTERNALERROR_DESCRIBEDOMAININFOERR = "InternalError.DescribeDomainInfoErr"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INTERNALERROR_SETUPDATEPROHIBITEDERR = "InternalError.SetUpdateProhibitedErr"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
//  RESOURCEUNAVAILABLE_DOMAINISMODIFYINGDNS = "ResourceUnavailable.DomainIsModifyingDNS"
func (c *Client) UpdateProhibitionIntlBatchWithContext(ctx context.Context, request *UpdateProhibitionIntlBatchRequest) (response *UpdateProhibitionIntlBatchResponse, err error) {
    if request == nil {
        request = NewUpdateProhibitionIntlBatchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateProhibitionIntlBatch require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateProhibitionIntlBatchResponse()
    err = c.Send(request, response)
    return
}
