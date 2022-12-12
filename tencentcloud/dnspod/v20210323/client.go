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

package v20210323

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2021-03-23"

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


func NewCreateDomainRequest() (request *CreateDomainRequest) {
    request = &CreateDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "CreateDomain")
    
    
    return
}

func NewCreateDomainResponse() (response *CreateDomainResponse) {
    response = &CreateDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDomain
// This API is used to add a domain.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINEXISTS = "FailedOperation.DomainExists"
//  FAILEDOPERATION_DOMAINOWNEDBYOTHERUSER = "FailedOperation.DomainOwnedByOtherUser"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_ACCOUNTISBANNED = "InvalidParameter.AccountIsBanned"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINNOTREGED = "InvalidParameter.DomainNotReged"
//  INVALIDPARAMETER_EMAILNOTVERIFIED = "InvalidParameter.EmailNotVerified"
//  INVALIDPARAMETER_TOOLSDOMAININVALID = "InvalidParameter.ToolsDomainInvalid"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
func (c *Client) CreateDomain(request *CreateDomainRequest) (response *CreateDomainResponse, err error) {
    return c.CreateDomainWithContext(context.Background(), request)
}

// CreateDomain
// This API is used to add a domain.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINEXISTS = "FailedOperation.DomainExists"
//  FAILEDOPERATION_DOMAINOWNEDBYOTHERUSER = "FailedOperation.DomainOwnedByOtherUser"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_ACCOUNTISBANNED = "InvalidParameter.AccountIsBanned"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINNOTREGED = "InvalidParameter.DomainNotReged"
//  INVALIDPARAMETER_EMAILNOTVERIFIED = "InvalidParameter.EmailNotVerified"
//  INVALIDPARAMETER_TOOLSDOMAININVALID = "InvalidParameter.ToolsDomainInvalid"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
func (c *Client) CreateDomainWithContext(ctx context.Context, request *CreateDomainRequest) (response *CreateDomainResponse, err error) {
    if request == nil {
        request = NewCreateDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRecordRequest() (request *CreateRecordRequest) {
    request = &CreateRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "CreateRecord")
    
    
    return
}

func NewCreateRecordResponse() (response *CreateRecordResponse) {
    response = &CreateRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRecord
// This API is used to add a record.
//
// error code that may be returned:
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_ACCOUNTISBANNED = "InvalidParameter.AccountIsBanned"
//  INVALIDPARAMETER_CUSTOMMESSAGE = "InvalidParameter.CustomMessage"
//  INVALIDPARAMETER_DNSSECADDCNAMEERROR = "InvalidParameter.DnssecAddCnameError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINNOTALLOWEDMODIFYRECORDS = "InvalidParameter.DomainNotAllowedModifyRecords"
//  INVALIDPARAMETER_DOMAINNOTBEIAN = "InvalidParameter.DomainNotBeian"
//  INVALIDPARAMETER_DOMAINRECORDEXIST = "InvalidParameter.DomainRecordExist"
//  INVALIDPARAMETER_EMAILNOTVERIFIED = "InvalidParameter.EmailNotVerified"
//  INVALIDPARAMETER_INVALIDWEIGHT = "InvalidParameter.InvalidWeight"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_MOBILENOTVERIFIED = "InvalidParameter.MobileNotVerified"
//  INVALIDPARAMETER_MXINVALID = "InvalidParameter.MxInvalid"
//  INVALIDPARAMETER_RECORDLINEINVALID = "InvalidParameter.RecordLineInvalid"
//  INVALIDPARAMETER_RECORDTYPEINVALID = "InvalidParameter.RecordTypeInvalid"
//  INVALIDPARAMETER_RECORDVALUEINVALID = "InvalidParameter.RecordValueInvalid"
//  INVALIDPARAMETER_RECORDVALUELENGTHINVALID = "InvalidParameter.RecordValueLengthInvalid"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_SUBDOMAININVALID = "InvalidParameter.SubdomainInvalid"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_URLVALUEILLEGAL = "InvalidParameter.UrlValueIllegal"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_AAAACOUNTLIMIT = "LimitExceeded.AAAACountLimit"
//  LIMITEXCEEDED_ATNSRECORDLIMIT = "LimitExceeded.AtNsRecordLimit"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  LIMITEXCEEDED_HIDDENURLEXCEEDED = "LimitExceeded.HiddenUrlExceeded"
//  LIMITEXCEEDED_NSCOUNTLIMIT = "LimitExceeded.NsCountLimit"
//  LIMITEXCEEDED_RECORDTTLLIMIT = "LimitExceeded.RecordTtlLimit"
//  LIMITEXCEEDED_SRVCOUNTLIMIT = "LimitExceeded.SrvCountLimit"
//  LIMITEXCEEDED_SUBDOMAINLEVELLIMIT = "LimitExceeded.SubdomainLevelLimit"
//  LIMITEXCEEDED_SUBDOMAINROLLLIMIT = "LimitExceeded.SubdomainRollLimit"
//  LIMITEXCEEDED_SUBDOMAINWCARDLIMIT = "LimitExceeded.SubdomainWcardLimit"
//  LIMITEXCEEDED_URLCOUNTLIMIT = "LimitExceeded.UrlCountLimit"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_IPINBLACKLISTNOTALLOWED = "OperationDenied.IPInBlacklistNotAllowed"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTADMIN = "OperationDenied.NotAdmin"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
func (c *Client) CreateRecord(request *CreateRecordRequest) (response *CreateRecordResponse, err error) {
    return c.CreateRecordWithContext(context.Background(), request)
}

// CreateRecord
// This API is used to add a record.
//
// error code that may be returned:
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_ACCOUNTISBANNED = "InvalidParameter.AccountIsBanned"
//  INVALIDPARAMETER_CUSTOMMESSAGE = "InvalidParameter.CustomMessage"
//  INVALIDPARAMETER_DNSSECADDCNAMEERROR = "InvalidParameter.DnssecAddCnameError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINNOTALLOWEDMODIFYRECORDS = "InvalidParameter.DomainNotAllowedModifyRecords"
//  INVALIDPARAMETER_DOMAINNOTBEIAN = "InvalidParameter.DomainNotBeian"
//  INVALIDPARAMETER_DOMAINRECORDEXIST = "InvalidParameter.DomainRecordExist"
//  INVALIDPARAMETER_EMAILNOTVERIFIED = "InvalidParameter.EmailNotVerified"
//  INVALIDPARAMETER_INVALIDWEIGHT = "InvalidParameter.InvalidWeight"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_MOBILENOTVERIFIED = "InvalidParameter.MobileNotVerified"
//  INVALIDPARAMETER_MXINVALID = "InvalidParameter.MxInvalid"
//  INVALIDPARAMETER_RECORDLINEINVALID = "InvalidParameter.RecordLineInvalid"
//  INVALIDPARAMETER_RECORDTYPEINVALID = "InvalidParameter.RecordTypeInvalid"
//  INVALIDPARAMETER_RECORDVALUEINVALID = "InvalidParameter.RecordValueInvalid"
//  INVALIDPARAMETER_RECORDVALUELENGTHINVALID = "InvalidParameter.RecordValueLengthInvalid"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_SUBDOMAININVALID = "InvalidParameter.SubdomainInvalid"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_URLVALUEILLEGAL = "InvalidParameter.UrlValueIllegal"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_AAAACOUNTLIMIT = "LimitExceeded.AAAACountLimit"
//  LIMITEXCEEDED_ATNSRECORDLIMIT = "LimitExceeded.AtNsRecordLimit"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  LIMITEXCEEDED_HIDDENURLEXCEEDED = "LimitExceeded.HiddenUrlExceeded"
//  LIMITEXCEEDED_NSCOUNTLIMIT = "LimitExceeded.NsCountLimit"
//  LIMITEXCEEDED_RECORDTTLLIMIT = "LimitExceeded.RecordTtlLimit"
//  LIMITEXCEEDED_SRVCOUNTLIMIT = "LimitExceeded.SrvCountLimit"
//  LIMITEXCEEDED_SUBDOMAINLEVELLIMIT = "LimitExceeded.SubdomainLevelLimit"
//  LIMITEXCEEDED_SUBDOMAINROLLLIMIT = "LimitExceeded.SubdomainRollLimit"
//  LIMITEXCEEDED_SUBDOMAINWCARDLIMIT = "LimitExceeded.SubdomainWcardLimit"
//  LIMITEXCEEDED_URLCOUNTLIMIT = "LimitExceeded.UrlCountLimit"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_IPINBLACKLISTNOTALLOWED = "OperationDenied.IPInBlacklistNotAllowed"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTADMIN = "OperationDenied.NotAdmin"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
func (c *Client) CreateRecordWithContext(ctx context.Context, request *CreateRecordRequest) (response *CreateRecordResponse, err error) {
    if request == nil {
        request = NewCreateRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDomainRequest() (request *DeleteDomainRequest) {
    request = &DeleteDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "DeleteDomain")
    
    
    return
}

func NewDeleteDomainResponse() (response *DeleteDomainResponse) {
    response = &DeleteDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDomain
// This API is used to delete a domain.
//
// error code that may be returned:
//  FAILEDOPERATION_DOMAINISKEYDOMAIN = "FailedOperation.DomainIsKeyDomain"
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_DOMAINISVIP = "FailedOperation.DomainIsVip"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_ACCOUNTISBANNED = "InvalidParameter.AccountIsBanned"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININEFFECTORINVALIDATED = "InvalidParameter.DomainInEffectOrInvalidated"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
func (c *Client) DeleteDomain(request *DeleteDomainRequest) (response *DeleteDomainResponse, err error) {
    return c.DeleteDomainWithContext(context.Background(), request)
}

// DeleteDomain
// This API is used to delete a domain.
//
// error code that may be returned:
//  FAILEDOPERATION_DOMAINISKEYDOMAIN = "FailedOperation.DomainIsKeyDomain"
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_DOMAINISVIP = "FailedOperation.DomainIsVip"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_ACCOUNTISBANNED = "InvalidParameter.AccountIsBanned"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININEFFECTORINVALIDATED = "InvalidParameter.DomainInEffectOrInvalidated"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
func (c *Client) DeleteDomainWithContext(ctx context.Context, request *DeleteDomainRequest) (response *DeleteDomainResponse, err error) {
    if request == nil {
        request = NewDeleteDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRecordRequest() (request *DeleteRecordRequest) {
    request = &DeleteRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "DeleteRecord")
    
    
    return
}

func NewDeleteRecordResponse() (response *DeleteRecordResponse) {
    response = &DeleteRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRecord
// This API is used to delete a record.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINNOTALLOWEDMODIFYRECORDS = "InvalidParameter.DomainNotAllowedModifyRecords"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_RECORDIDINVALID = "InvalidParameter.RecordIdInvalid"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTADMIN = "OperationDenied.NotAdmin"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
func (c *Client) DeleteRecord(request *DeleteRecordRequest) (response *DeleteRecordResponse, err error) {
    return c.DeleteRecordWithContext(context.Background(), request)
}

// DeleteRecord
// This API is used to delete a record.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINNOTALLOWEDMODIFYRECORDS = "InvalidParameter.DomainNotAllowedModifyRecords"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_RECORDIDINVALID = "InvalidParameter.RecordIdInvalid"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTADMIN = "OperationDenied.NotAdmin"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
func (c *Client) DeleteRecordWithContext(ctx context.Context, request *DeleteRecordRequest) (response *DeleteRecordResponse, err error) {
    if request == nil {
        request = NewDeleteRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainRequest() (request *DescribeDomainRequest) {
    request = &DescribeDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribeDomain")
    
    
    return
}

func NewDescribeDomainResponse() (response *DescribeDomainResponse) {
    response = &DescribeDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDomain
// This API is used to get the information of a domain.
//
// error code that may be returned:
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) DescribeDomain(request *DescribeDomainRequest) (response *DescribeDomainResponse, err error) {
    return c.DescribeDomainWithContext(context.Background(), request)
}

// DescribeDomain
// This API is used to get the information of a domain.
//
// error code that may be returned:
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) DescribeDomainWithContext(ctx context.Context, request *DescribeDomainRequest) (response *DescribeDomainResponse, err error) {
    if request == nil {
        request = NewDescribeDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordRequest() (request *DescribeRecordRequest) {
    request = &DescribeRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribeRecord")
    
    
    return
}

func NewDescribeRecordResponse() (response *DescribeRecordResponse) {
    response = &DescribeRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRecord
// This API is used to get the information of a record.
//
// error code that may be returned:
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_RECORDIDINVALID = "InvalidParameter.RecordIdInvalid"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTADMIN = "OperationDenied.NotAdmin"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
func (c *Client) DescribeRecord(request *DescribeRecordRequest) (response *DescribeRecordResponse, err error) {
    return c.DescribeRecordWithContext(context.Background(), request)
}

// DescribeRecord
// This API is used to get the information of a record.
//
// error code that may be returned:
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_RECORDIDINVALID = "InvalidParameter.RecordIdInvalid"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTADMIN = "OperationDenied.NotAdmin"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
func (c *Client) DescribeRecordWithContext(ctx context.Context, request *DescribeRecordRequest) (response *DescribeRecordResponse, err error) {
    if request == nil {
        request = NewDescribeRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordListRequest() (request *DescribeRecordListRequest) {
    request = &DescribeRecordListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribeRecordList")
    
    
    return
}

func NewDescribeRecordListResponse() (response *DescribeRecordListResponse) {
    response = &DescribeRecordListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRecordList
// This API is used to get the DNS records of a domain.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTDOMAINOWNER = "FailedOperation.NotDomainOwner"
//  FAILEDOPERATION_NOTREALNAMEDUSER = "FailedOperation.NotRealNamedUser"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETER_PARAMINVALID = "InvalidParameter.ParamInvalid"
//  INVALIDPARAMETER_RECORDLINEINVALID = "InvalidParameter.RecordLineInvalid"
//  INVALIDPARAMETER_RECORDTYPEINVALID = "InvalidParameter.RecordTypeInvalid"
//  INVALIDPARAMETER_RESULTMORETHAN500 = "InvalidParameter.ResultMoreThan500"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_LIMITINVALID = "InvalidParameterValue.LimitInvalid"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
//  RESOURCENOTFOUND_NODATAOFRECORD = "ResourceNotFound.NoDataOfRecord"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeRecordList(request *DescribeRecordListRequest) (response *DescribeRecordListResponse, err error) {
    return c.DescribeRecordListWithContext(context.Background(), request)
}

// DescribeRecordList
// This API is used to get the DNS records of a domain.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTDOMAINOWNER = "FailedOperation.NotDomainOwner"
//  FAILEDOPERATION_NOTREALNAMEDUSER = "FailedOperation.NotRealNamedUser"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETER_PARAMINVALID = "InvalidParameter.ParamInvalid"
//  INVALIDPARAMETER_RECORDLINEINVALID = "InvalidParameter.RecordLineInvalid"
//  INVALIDPARAMETER_RECORDTYPEINVALID = "InvalidParameter.RecordTypeInvalid"
//  INVALIDPARAMETER_RESULTMORETHAN500 = "InvalidParameter.ResultMoreThan500"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_LIMITINVALID = "InvalidParameterValue.LimitInvalid"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
//  RESOURCENOTFOUND_NODATAOFRECORD = "ResourceNotFound.NoDataOfRecord"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeRecordListWithContext(ctx context.Context, request *DescribeRecordListRequest) (response *DescribeRecordListResponse, err error) {
    if request == nil {
        request = NewDescribeRecordListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecordList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecordListResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDomainRemarkRequest() (request *ModifyDomainRemarkRequest) {
    request = &ModifyDomainRemarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "ModifyDomainRemark")
    
    
    return
}

func NewModifyDomainRemarkResponse() (response *ModifyDomainRemarkResponse) {
    response = &ModifyDomainRemarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDomainRemark
// This API is used to set the remarks of a domain.
//
// error code that may be returned:
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_REMARKTOOLONG = "InvalidParameter.RemarkTooLong"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
func (c *Client) ModifyDomainRemark(request *ModifyDomainRemarkRequest) (response *ModifyDomainRemarkResponse, err error) {
    return c.ModifyDomainRemarkWithContext(context.Background(), request)
}

// ModifyDomainRemark
// This API is used to set the remarks of a domain.
//
// error code that may be returned:
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_REMARKTOOLONG = "InvalidParameter.RemarkTooLong"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
func (c *Client) ModifyDomainRemarkWithContext(ctx context.Context, request *ModifyDomainRemarkRequest) (response *ModifyDomainRemarkResponse, err error) {
    if request == nil {
        request = NewModifyDomainRemarkRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDomainRemark require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDomainRemarkResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDomainStatusRequest() (request *ModifyDomainStatusRequest) {
    request = &ModifyDomainStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "ModifyDomainStatus")
    
    
    return
}

func NewModifyDomainStatusResponse() (response *ModifyDomainStatusResponse) {
    response = &ModifyDomainStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDomainStatus
// This API is used to modify the status of a domain.
//
// error code that may be returned:
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_TOOLSDOMAININVALID = "InvalidParameter.ToolsDomainInvalid"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) ModifyDomainStatus(request *ModifyDomainStatusRequest) (response *ModifyDomainStatusResponse, err error) {
    return c.ModifyDomainStatusWithContext(context.Background(), request)
}

// ModifyDomainStatus
// This API is used to modify the status of a domain.
//
// error code that may be returned:
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_TOOLSDOMAININVALID = "InvalidParameter.ToolsDomainInvalid"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) ModifyDomainStatusWithContext(ctx context.Context, request *ModifyDomainStatusRequest) (response *ModifyDomainStatusResponse, err error) {
    if request == nil {
        request = NewModifyDomainStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDomainStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDomainStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRecordRequest() (request *ModifyRecordRequest) {
    request = &ModifyRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "ModifyRecord")
    
    
    return
}

func NewModifyRecordResponse() (response *ModifyRecordResponse) {
    response = &ModifyRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyRecord
// This API is used to modify a record.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_ACCOUNTISBANNED = "InvalidParameter.AccountIsBanned"
//  INVALIDPARAMETER_CUSTOMMESSAGE = "InvalidParameter.CustomMessage"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINNOTALLOWEDMODIFYRECORDS = "InvalidParameter.DomainNotAllowedModifyRecords"
//  INVALIDPARAMETER_DOMAINNOTBEIAN = "InvalidParameter.DomainNotBeian"
//  INVALIDPARAMETER_DOMAINRECORDEXIST = "InvalidParameter.DomainRecordExist"
//  INVALIDPARAMETER_EMAILNOTVERIFIED = "InvalidParameter.EmailNotVerified"
//  INVALIDPARAMETER_INVALIDWEIGHT = "InvalidParameter.InvalidWeight"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_MOBILENOTVERIFIED = "InvalidParameter.MobileNotVerified"
//  INVALIDPARAMETER_MXINVALID = "InvalidParameter.MxInvalid"
//  INVALIDPARAMETER_RECORDIDINVALID = "InvalidParameter.RecordIdInvalid"
//  INVALIDPARAMETER_RECORDLINEINVALID = "InvalidParameter.RecordLineInvalid"
//  INVALIDPARAMETER_RECORDTYPEINVALID = "InvalidParameter.RecordTypeInvalid"
//  INVALIDPARAMETER_RECORDVALUEINVALID = "InvalidParameter.RecordValueInvalid"
//  INVALIDPARAMETER_RECORDVALUELENGTHINVALID = "InvalidParameter.RecordValueLengthInvalid"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_SUBDOMAININVALID = "InvalidParameter.SubdomainInvalid"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_URLVALUEILLEGAL = "InvalidParameter.UrlValueIllegal"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_AAAACOUNTLIMIT = "LimitExceeded.AAAACountLimit"
//  LIMITEXCEEDED_ATNSRECORDLIMIT = "LimitExceeded.AtNsRecordLimit"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  LIMITEXCEEDED_HIDDENURLEXCEEDED = "LimitExceeded.HiddenUrlExceeded"
//  LIMITEXCEEDED_NSCOUNTLIMIT = "LimitExceeded.NsCountLimit"
//  LIMITEXCEEDED_RECORDTTLLIMIT = "LimitExceeded.RecordTtlLimit"
//  LIMITEXCEEDED_SRVCOUNTLIMIT = "LimitExceeded.SrvCountLimit"
//  LIMITEXCEEDED_SUBDOMAINLEVELLIMIT = "LimitExceeded.SubdomainLevelLimit"
//  LIMITEXCEEDED_SUBDOMAINROLLLIMIT = "LimitExceeded.SubdomainRollLimit"
//  LIMITEXCEEDED_SUBDOMAINWCARDLIMIT = "LimitExceeded.SubdomainWcardLimit"
//  LIMITEXCEEDED_URLCOUNTLIMIT = "LimitExceeded.UrlCountLimit"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_IPINBLACKLISTNOTALLOWED = "OperationDenied.IPInBlacklistNotAllowed"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTADMIN = "OperationDenied.NotAdmin"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
func (c *Client) ModifyRecord(request *ModifyRecordRequest) (response *ModifyRecordResponse, err error) {
    return c.ModifyRecordWithContext(context.Background(), request)
}

// ModifyRecord
// This API is used to modify a record.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_ACCOUNTISBANNED = "InvalidParameter.AccountIsBanned"
//  INVALIDPARAMETER_CUSTOMMESSAGE = "InvalidParameter.CustomMessage"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINNOTALLOWEDMODIFYRECORDS = "InvalidParameter.DomainNotAllowedModifyRecords"
//  INVALIDPARAMETER_DOMAINNOTBEIAN = "InvalidParameter.DomainNotBeian"
//  INVALIDPARAMETER_DOMAINRECORDEXIST = "InvalidParameter.DomainRecordExist"
//  INVALIDPARAMETER_EMAILNOTVERIFIED = "InvalidParameter.EmailNotVerified"
//  INVALIDPARAMETER_INVALIDWEIGHT = "InvalidParameter.InvalidWeight"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_MOBILENOTVERIFIED = "InvalidParameter.MobileNotVerified"
//  INVALIDPARAMETER_MXINVALID = "InvalidParameter.MxInvalid"
//  INVALIDPARAMETER_RECORDIDINVALID = "InvalidParameter.RecordIdInvalid"
//  INVALIDPARAMETER_RECORDLINEINVALID = "InvalidParameter.RecordLineInvalid"
//  INVALIDPARAMETER_RECORDTYPEINVALID = "InvalidParameter.RecordTypeInvalid"
//  INVALIDPARAMETER_RECORDVALUEINVALID = "InvalidParameter.RecordValueInvalid"
//  INVALIDPARAMETER_RECORDVALUELENGTHINVALID = "InvalidParameter.RecordValueLengthInvalid"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_SUBDOMAININVALID = "InvalidParameter.SubdomainInvalid"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_URLVALUEILLEGAL = "InvalidParameter.UrlValueIllegal"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_AAAACOUNTLIMIT = "LimitExceeded.AAAACountLimit"
//  LIMITEXCEEDED_ATNSRECORDLIMIT = "LimitExceeded.AtNsRecordLimit"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  LIMITEXCEEDED_HIDDENURLEXCEEDED = "LimitExceeded.HiddenUrlExceeded"
//  LIMITEXCEEDED_NSCOUNTLIMIT = "LimitExceeded.NsCountLimit"
//  LIMITEXCEEDED_RECORDTTLLIMIT = "LimitExceeded.RecordTtlLimit"
//  LIMITEXCEEDED_SRVCOUNTLIMIT = "LimitExceeded.SrvCountLimit"
//  LIMITEXCEEDED_SUBDOMAINLEVELLIMIT = "LimitExceeded.SubdomainLevelLimit"
//  LIMITEXCEEDED_SUBDOMAINROLLLIMIT = "LimitExceeded.SubdomainRollLimit"
//  LIMITEXCEEDED_SUBDOMAINWCARDLIMIT = "LimitExceeded.SubdomainWcardLimit"
//  LIMITEXCEEDED_URLCOUNTLIMIT = "LimitExceeded.UrlCountLimit"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_IPINBLACKLISTNOTALLOWED = "OperationDenied.IPInBlacklistNotAllowed"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTADMIN = "OperationDenied.NotAdmin"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
func (c *Client) ModifyRecordWithContext(ctx context.Context, request *ModifyRecordRequest) (response *ModifyRecordResponse, err error) {
    if request == nil {
        request = NewModifyRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRecordResponse()
    err = c.Send(request, response)
    return
}
