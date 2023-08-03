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

package v20220901

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2022-09-01"

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


func NewBindZoneToPlanRequest() (request *BindZoneToPlanRequest) {
    request = &BindZoneToPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "BindZoneToPlan")
    
    
    return
}

func NewBindZoneToPlanResponse() (response *BindZoneToPlanResponse) {
    response = &BindZoneToPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindZoneToPlan
// This API is used to bind a site to a plan.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PLANNOTFOUND = "InvalidParameter.PlanNotFound"
//  INVALIDPARAMETER_ZONEHASBEENBOUND = "InvalidParameter.ZoneHasBeenBound"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) BindZoneToPlan(request *BindZoneToPlanRequest) (response *BindZoneToPlanResponse, err error) {
    return c.BindZoneToPlanWithContext(context.Background(), request)
}

// BindZoneToPlan
// This API is used to bind a site to a plan.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PLANNOTFOUND = "InvalidParameter.PlanNotFound"
//  INVALIDPARAMETER_ZONEHASBEENBOUND = "InvalidParameter.ZoneHasBeenBound"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) BindZoneToPlanWithContext(ctx context.Context, request *BindZoneToPlanRequest) (response *BindZoneToPlanResponse, err error) {
    if request == nil {
        request = NewBindZoneToPlanRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindZoneToPlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindZoneToPlanResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAccelerationDomainRequest() (request *CreateAccelerationDomainRequest) {
    request = &CreateAccelerationDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateAccelerationDomain")
    
    
    return
}

func NewCreateAccelerationDomainResponse() (response *CreateAccelerationDomainResponse) {
    response = &CreateAccelerationDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAccelerationDomain
// This API is used to connect a domain to EdgeOne.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INVALIDPARAMETER_INVALIDACCELERATETYPE = "InvalidParameter.InvalidAccelerateType"
//  INVALIDPARAMETER_INVALIDAWSPRIVATEACCESS = "InvalidParameter.InvalidAwsPrivateAccess"
//  INVALIDPARAMETER_INVALIDCLIENTIPORIGIN = "InvalidParameter.InvalidClientIpOrigin"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_INVALIDPRIVATEACCESSPARAMS = "InvalidParameter.InvalidPrivateAccessParams"
//  INVALIDPARAMETER_INVALIDQUICBILLING = "InvalidParameter.InvalidQuicBilling"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETERVALUE_CONFLICTRECORD = "InvalidParameterValue.ConflictRecord"
//  INVALIDPARAMETERVALUE_CONFLICTWITHDNSSEC = "InvalidParameterValue.ConflictWithDNSSEC"
//  INVALIDPARAMETERVALUE_CONFLICTWITHNSRECORD = "InvalidParameterValue.ConflictWithNSRecord"
//  INVALIDPARAMETERVALUE_CONTENTSAMEASNAME = "InvalidParameterValue.ContentSameAsName"
//  INVALIDPARAMETERVALUE_DOMAINNOTMATCHZONE = "InvalidParameterValue.DomainNotMatchZone"
//  INVALIDPARAMETERVALUE_INVALIDDNSCONTENT = "InvalidParameterValue.InvalidDNSContent"
//  INVALIDPARAMETERVALUE_INVALIDDNSNAME = "InvalidParameterValue.InvalidDNSName"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDDISABLE = "OperationDenied.AccelerateMainlandDisable"
//  OPERATIONDENIED_DOMAINNOICP = "OperationDenied.DomainNoICP"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_INVALIDADVANCEDDEFENSEZONEAREA = "OperationDenied.InvalidAdvancedDefenseZoneArea"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_DNSRECORD = "ResourceInUse.DnsRecord"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_DOMAINALREADYEXISTS = "ResourceUnavailable.DomainAlreadyExists"
//  RESOURCESSOLDOUT_L7LACKOFRESOURCES = "ResourcesSoldOut.L7LackOfResources"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) CreateAccelerationDomain(request *CreateAccelerationDomainRequest) (response *CreateAccelerationDomainResponse, err error) {
    return c.CreateAccelerationDomainWithContext(context.Background(), request)
}

// CreateAccelerationDomain
// This API is used to connect a domain to EdgeOne.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INVALIDPARAMETER_INVALIDACCELERATETYPE = "InvalidParameter.InvalidAccelerateType"
//  INVALIDPARAMETER_INVALIDAWSPRIVATEACCESS = "InvalidParameter.InvalidAwsPrivateAccess"
//  INVALIDPARAMETER_INVALIDCLIENTIPORIGIN = "InvalidParameter.InvalidClientIpOrigin"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_INVALIDPRIVATEACCESSPARAMS = "InvalidParameter.InvalidPrivateAccessParams"
//  INVALIDPARAMETER_INVALIDQUICBILLING = "InvalidParameter.InvalidQuicBilling"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETERVALUE_CONFLICTRECORD = "InvalidParameterValue.ConflictRecord"
//  INVALIDPARAMETERVALUE_CONFLICTWITHDNSSEC = "InvalidParameterValue.ConflictWithDNSSEC"
//  INVALIDPARAMETERVALUE_CONFLICTWITHNSRECORD = "InvalidParameterValue.ConflictWithNSRecord"
//  INVALIDPARAMETERVALUE_CONTENTSAMEASNAME = "InvalidParameterValue.ContentSameAsName"
//  INVALIDPARAMETERVALUE_DOMAINNOTMATCHZONE = "InvalidParameterValue.DomainNotMatchZone"
//  INVALIDPARAMETERVALUE_INVALIDDNSCONTENT = "InvalidParameterValue.InvalidDNSContent"
//  INVALIDPARAMETERVALUE_INVALIDDNSNAME = "InvalidParameterValue.InvalidDNSName"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDDISABLE = "OperationDenied.AccelerateMainlandDisable"
//  OPERATIONDENIED_DOMAINNOICP = "OperationDenied.DomainNoICP"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_INVALIDADVANCEDDEFENSEZONEAREA = "OperationDenied.InvalidAdvancedDefenseZoneArea"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_DNSRECORD = "ResourceInUse.DnsRecord"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_DOMAINALREADYEXISTS = "ResourceUnavailable.DomainAlreadyExists"
//  RESOURCESSOLDOUT_L7LACKOFRESOURCES = "ResourcesSoldOut.L7LackOfResources"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) CreateAccelerationDomainWithContext(ctx context.Context, request *CreateAccelerationDomainRequest) (response *CreateAccelerationDomainResponse, err error) {
    if request == nil {
        request = NewCreateAccelerationDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAccelerationDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAccelerationDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAliasDomainRequest() (request *CreateAliasDomainRequest) {
    request = &CreateAliasDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateAliasDomain")
    
    
    return
}

func NewCreateAliasDomainResponse() (response *CreateAliasDomainResponse) {
    response = &CreateAliasDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAliasDomain
// This API is used to create an alias domain name.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INVALIDPARAMETER_ALIASDOMAINNOTSUPPORTSMCERT = "InvalidParameter.AliasDomainNotSupportSMCert"
//  INVALIDPARAMETER_CERTNOTMATCHDOMAIN = "InvalidParameter.CertNotMatchDomain"
//  INVALIDPARAMETER_INVALIDHTTPSCERTINFO = "InvalidParameter.InvalidHttpsCertInfo"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DOMAINISBLOCKED = "OperationDenied.DomainIsBlocked"
//  OPERATIONDENIED_DOMAINNOICP = "OperationDenied.DomainNoICP"
//  RESOURCEINUSE_ALIASNAME = "ResourceInUse.AliasName"
//  RESOURCEINUSE_DUPLICATENAME = "ResourceInUse.DuplicateName"
//  RESOURCEINUSE_ZONE = "ResourceInUse.Zone"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_DOMAINALREADYEXISTS = "ResourceUnavailable.DomainAlreadyExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_TARGETNAMEORIGINTYPECOS = "UnsupportedOperation.TargetNameOriginTypeCos"
func (c *Client) CreateAliasDomain(request *CreateAliasDomainRequest) (response *CreateAliasDomainResponse, err error) {
    return c.CreateAliasDomainWithContext(context.Background(), request)
}

// CreateAliasDomain
// This API is used to create an alias domain name.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INVALIDPARAMETER_ALIASDOMAINNOTSUPPORTSMCERT = "InvalidParameter.AliasDomainNotSupportSMCert"
//  INVALIDPARAMETER_CERTNOTMATCHDOMAIN = "InvalidParameter.CertNotMatchDomain"
//  INVALIDPARAMETER_INVALIDHTTPSCERTINFO = "InvalidParameter.InvalidHttpsCertInfo"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DOMAINISBLOCKED = "OperationDenied.DomainIsBlocked"
//  OPERATIONDENIED_DOMAINNOICP = "OperationDenied.DomainNoICP"
//  RESOURCEINUSE_ALIASNAME = "ResourceInUse.AliasName"
//  RESOURCEINUSE_DUPLICATENAME = "ResourceInUse.DuplicateName"
//  RESOURCEINUSE_ZONE = "ResourceInUse.Zone"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_DOMAINALREADYEXISTS = "ResourceUnavailable.DomainAlreadyExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_TARGETNAMEORIGINTYPECOS = "UnsupportedOperation.TargetNameOriginTypeCos"
func (c *Client) CreateAliasDomainWithContext(ctx context.Context, request *CreateAliasDomainRequest) (response *CreateAliasDomainResponse, err error) {
    if request == nil {
        request = NewCreateAliasDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAliasDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAliasDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApplicationProxyRequest() (request *CreateApplicationProxyRequest) {
    request = &CreateApplicationProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateApplicationProxy")
    
    
    return
}

func NewCreateApplicationProxyResponse() (response *CreateApplicationProxyResponse) {
    response = &CreateApplicationProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateApplicationProxy
// This API is used to create an application proxy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateApplicationProxy(request *CreateApplicationProxyRequest) (response *CreateApplicationProxyResponse, err error) {
    return c.CreateApplicationProxyWithContext(context.Background(), request)
}

// CreateApplicationProxy
// This API is used to create an application proxy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateApplicationProxyWithContext(ctx context.Context, request *CreateApplicationProxyRequest) (response *CreateApplicationProxyResponse, err error) {
    if request == nil {
        request = NewCreateApplicationProxyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApplicationProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApplicationProxyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApplicationProxyRuleRequest() (request *CreateApplicationProxyRuleRequest) {
    request = &CreateApplicationProxyRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateApplicationProxyRule")
    
    
    return
}

func NewCreateApplicationProxyRuleResponse() (response *CreateApplicationProxyRuleResponse) {
    response = &CreateApplicationProxyRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateApplicationProxyRule
// This API is used to create an application proxy rule.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateApplicationProxyRule(request *CreateApplicationProxyRuleRequest) (response *CreateApplicationProxyRuleResponse, err error) {
    return c.CreateApplicationProxyRuleWithContext(context.Background(), request)
}

// CreateApplicationProxyRule
// This API is used to create an application proxy rule.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateApplicationProxyRuleWithContext(ctx context.Context, request *CreateApplicationProxyRuleRequest) (response *CreateApplicationProxyRuleResponse, err error) {
    if request == nil {
        request = NewCreateApplicationProxyRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApplicationProxyRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApplicationProxyRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOriginGroupRequest() (request *CreateOriginGroupRequest) {
    request = &CreateOriginGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateOriginGroup")
    
    
    return
}

func NewCreateOriginGroupResponse() (response *CreateOriginGroupResponse) {
    response = &CreateOriginGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateOriginGroup
// This API is used to create an origin group.
//
// error code that may be returned:
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) CreateOriginGroup(request *CreateOriginGroupRequest) (response *CreateOriginGroupResponse, err error) {
    return c.CreateOriginGroupWithContext(context.Background(), request)
}

// CreateOriginGroup
// This API is used to create an origin group.
//
// error code that may be returned:
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) CreateOriginGroupWithContext(ctx context.Context, request *CreateOriginGroupRequest) (response *CreateOriginGroupResponse, err error) {
    if request == nil {
        request = NewCreateOriginGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOriginGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOriginGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePlanForZoneRequest() (request *CreatePlanForZoneRequest) {
    request = &CreatePlanForZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreatePlanForZone")
    
    
    return
}

func NewCreatePlanForZoneResponse() (response *CreatePlanForZoneResponse) {
    response = &CreatePlanForZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePlanForZone
// This API is used to purchase a plan for a new site.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSUFFICIENTACCOUNTBALANCE = "FailedOperation.InsufficientAccountBalance"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DOMAINNOICP = "OperationDenied.DomainNoICP"
func (c *Client) CreatePlanForZone(request *CreatePlanForZoneRequest) (response *CreatePlanForZoneResponse, err error) {
    return c.CreatePlanForZoneWithContext(context.Background(), request)
}

// CreatePlanForZone
// This API is used to purchase a plan for a new site.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSUFFICIENTACCOUNTBALANCE = "FailedOperation.InsufficientAccountBalance"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DOMAINNOICP = "OperationDenied.DomainNoICP"
func (c *Client) CreatePlanForZoneWithContext(ctx context.Context, request *CreatePlanForZoneRequest) (response *CreatePlanForZoneResponse, err error) {
    if request == nil {
        request = NewCreatePlanForZoneRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePlanForZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePlanForZoneResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrefetchTaskRequest() (request *CreatePrefetchTaskRequest) {
    request = &CreatePrefetchTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreatePrefetchTask")
    
    
    return
}

func NewCreatePrefetchTaskResponse() (response *CreatePrefetchTaskResponse) {
    response = &CreatePrefetchTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePrefetchTask
// This API is used to create a pre-warming task.
//
// error code that may be returned:
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DOMAINCONFIG = "InternalError.DomainConfig"
//  INTERNALERROR_FAILEDTOGENERATEURL = "InternalError.FailedToGenerateUrl"
//  INTERNALERROR_QUOTASYSTEM = "InternalError.QuotaSystem"
//  INVALIDPARAMETER_DOMAINNOTFOUND = "InvalidParameter.DomainNotFound"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETER_TARGET = "InvalidParameter.Target"
//  INVALIDPARAMETER_TASKNOTGENERATED = "InvalidParameter.TaskNotGenerated"
//  INVALIDPARAMETER_UPLOADURL = "InvalidParameter.UploadUrl"
//  LIMITEXCEEDED_BATCHQUOTA = "LimitExceeded.BatchQuota"
//  LIMITEXCEEDED_DAILYQUOTA = "LimitExceeded.DailyQuota"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) CreatePrefetchTask(request *CreatePrefetchTaskRequest) (response *CreatePrefetchTaskResponse, err error) {
    return c.CreatePrefetchTaskWithContext(context.Background(), request)
}

// CreatePrefetchTask
// This API is used to create a pre-warming task.
//
// error code that may be returned:
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DOMAINCONFIG = "InternalError.DomainConfig"
//  INTERNALERROR_FAILEDTOGENERATEURL = "InternalError.FailedToGenerateUrl"
//  INTERNALERROR_QUOTASYSTEM = "InternalError.QuotaSystem"
//  INVALIDPARAMETER_DOMAINNOTFOUND = "InvalidParameter.DomainNotFound"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETER_TARGET = "InvalidParameter.Target"
//  INVALIDPARAMETER_TASKNOTGENERATED = "InvalidParameter.TaskNotGenerated"
//  INVALIDPARAMETER_UPLOADURL = "InvalidParameter.UploadUrl"
//  LIMITEXCEEDED_BATCHQUOTA = "LimitExceeded.BatchQuota"
//  LIMITEXCEEDED_DAILYQUOTA = "LimitExceeded.DailyQuota"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) CreatePrefetchTaskWithContext(ctx context.Context, request *CreatePrefetchTaskRequest) (response *CreatePrefetchTaskResponse, err error) {
    if request == nil {
        request = NewCreatePrefetchTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrefetchTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrefetchTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePurgeTaskRequest() (request *CreatePurgeTaskRequest) {
    request = &CreatePurgeTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreatePurgeTask")
    
    
    return
}

func NewCreatePurgeTaskResponse() (response *CreatePurgeTaskResponse) {
    response = &CreatePurgeTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePurgeTask
// This API is used to create a cache purging task.
//
// error code that may be returned:
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DOMAINCONFIG = "InternalError.DomainConfig"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_QUOTASYSTEM = "InternalError.QuotaSystem"
//  INVALIDPARAMETER_DOMAINNOTFOUND = "InvalidParameter.DomainNotFound"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETER_TARGET = "InvalidParameter.Target"
//  INVALIDPARAMETER_TASKNOTGENERATED = "InvalidParameter.TaskNotGenerated"
//  INVALIDPARAMETER_UPLOADURL = "InvalidParameter.UploadUrl"
//  LIMITEXCEEDED_BATCHQUOTA = "LimitExceeded.BatchQuota"
//  LIMITEXCEEDED_DAILYQUOTA = "LimitExceeded.DailyQuota"
//  LIMITEXCEEDED_PACKNOTALLOW = "LimitExceeded.PackNotAllow"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) CreatePurgeTask(request *CreatePurgeTaskRequest) (response *CreatePurgeTaskResponse, err error) {
    return c.CreatePurgeTaskWithContext(context.Background(), request)
}

// CreatePurgeTask
// This API is used to create a cache purging task.
//
// error code that may be returned:
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DOMAINCONFIG = "InternalError.DomainConfig"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_QUOTASYSTEM = "InternalError.QuotaSystem"
//  INVALIDPARAMETER_DOMAINNOTFOUND = "InvalidParameter.DomainNotFound"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETER_TARGET = "InvalidParameter.Target"
//  INVALIDPARAMETER_TASKNOTGENERATED = "InvalidParameter.TaskNotGenerated"
//  INVALIDPARAMETER_UPLOADURL = "InvalidParameter.UploadUrl"
//  LIMITEXCEEDED_BATCHQUOTA = "LimitExceeded.BatchQuota"
//  LIMITEXCEEDED_DAILYQUOTA = "LimitExceeded.DailyQuota"
//  LIMITEXCEEDED_PACKNOTALLOW = "LimitExceeded.PackNotAllow"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) CreatePurgeTaskWithContext(ctx context.Context, request *CreatePurgeTaskRequest) (response *CreatePurgeTaskResponse, err error) {
    if request == nil {
        request = NewCreatePurgeTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePurgeTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePurgeTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRuleRequest() (request *CreateRuleRequest) {
    request = &CreateRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateRule")
    
    
    return
}

func NewCreateRuleResponse() (response *CreateRuleResponse) {
    response = &CreateRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRule
// This API is used to create a rule in the rule engine.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"
//  INVALIDPARAMETER_ERRINVALIDACTION = "InvalidParameter.ErrInvalidAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMACTION = "InvalidParameter.ErrInvalidActionParamAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMBADVALUETYPE = "InvalidParameter.ErrInvalidActionParamBadValueType"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMNAME = "InvalidParameter.ErrInvalidActionParamName"
//  INVALIDPARAMETER_ERRINVALIDACTIONTYPE = "InvalidParameter.ErrInvalidActionType"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONIGNORECASE = "InvalidParameter.ErrInvalidConditionIgnoreCase"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMEBADNAME = "InvalidParameter.ErrInvalidConditionNameBadName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMETARGETNOTSUPPORTNAME = "InvalidParameter.ErrInvalidConditionNameTargetNotSupportName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADREGULAR = "InvalidParameter.ErrInvalidConditionValueBadRegular"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUE = "InvalidParameter.ErrInvalidConditionValueBadValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUECONTAINFILENAMEEXTENSION = "InvalidParameter.ErrInvalidConditionValueBadValueContainFileNameExtension"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOLONGVALUE = "InvalidParameter.ErrInvalidConditionValueTooLongValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYREGULAR = "InvalidParameter.ErrInvalidConditionValueTooManyRegular"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYVALUES = "InvalidParameter.ErrInvalidConditionValueTooManyValues"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYWILDCARD = "InvalidParameter.ErrInvalidConditionValueTooManyWildcard"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEZEROLENGTH = "InvalidParameter.ErrInvalidConditionValueZeroLength"
//  INVALIDPARAMETER_HOSTNOTFOUND = "InvalidParameter.HostNotFound"
//  INVALIDPARAMETER_INVALIDAUTHENTICATION = "InvalidParameter.InvalidAuthentication"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESIGNPARAM = "InvalidParameter.InvalidAuthenticationTypeSignParam"
//  INVALIDPARAMETER_INVALIDBACKUPSERVERNAME = "InvalidParameter.InvalidBackupServerName"
//  INVALIDPARAMETER_INVALIDCACHECONFIGCACHE = "InvalidParameter.InvalidCacheConfigCache"
//  INVALIDPARAMETER_INVALIDCACHEKEY = "InvalidParameter.InvalidCacheKey"
//  INVALIDPARAMETER_INVALIDCACHEKEYQUERYSTRINGVALUE = "InvalidParameter.InvalidCacheKeyQueryStringValue"
//  INVALIDPARAMETER_INVALIDCACHETIME = "InvalidParameter.InvalidCacheTime"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINE = "InvalidParameter.InvalidDynamicRoutine"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINEBILLING = "InvalidParameter.InvalidDynamicRoutineBilling"
//  INVALIDPARAMETER_INVALIDERRORPAGEREDIRECTURL = "InvalidParameter.InvalidErrorPageRedirectUrl"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDHTTPSHSTSMAXAGE = "InvalidParameter.InvalidHttpsHstsMaxAge"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPOSTSIZEVALUE = "InvalidParameter.InvalidPostSizeValue"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAMEXFF = "InvalidParameter.InvalidRequestHeaderNameXff"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERVALUE = "InvalidParameter.InvalidRequestHeaderValue"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERNAME = "InvalidParameter.InvalidResponseHeaderName"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERVALUE = "InvalidParameter.InvalidResponseHeaderValue"
//  INVALIDPARAMETER_INVALIDRULEENGINEACTION = "InvalidParameter.InvalidRuleEngineAction"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGET = "InvalidParameter.InvalidRuleEngineTarget"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSEXTENSION = "InvalidParameter.InvalidRuleEngineTargetsExtension"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSURL = "InvalidParameter.InvalidRuleEngineTargetsUrl"
//  INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGEXPIRETIMELIMIT = "InvalidParameter.InvalidStandardDebugExpireTimeLimit"
//  INVALIDPARAMETER_INVALIDUPSTREAMREQUESTQUERYSTRINGVALUE = "InvalidParameter.InvalidUpstreamRequestQueryStringValue"
//  INVALIDPARAMETER_INVALIDURLREDIRECTHOST = "InvalidParameter.InvalidUrlRedirectHost"
//  INVALIDPARAMETER_INVALIDURLREDIRECTURL = "InvalidParameter.InvalidUrlRedirectUrl"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_ORIGINORIGINGROUPIDISREQUIRED = "InvalidParameter.OriginOriginGroupIdIsRequired"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_TASKSYSTEMERROR = "InvalidParameter.TaskSystemError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_INVALIDADVANCEDDEFENSESECURITYTYPE = "OperationDenied.InvalidAdvancedDefenseSecurityType"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) CreateRule(request *CreateRuleRequest) (response *CreateRuleResponse, err error) {
    return c.CreateRuleWithContext(context.Background(), request)
}

// CreateRule
// This API is used to create a rule in the rule engine.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"
//  INVALIDPARAMETER_ERRINVALIDACTION = "InvalidParameter.ErrInvalidAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMACTION = "InvalidParameter.ErrInvalidActionParamAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMBADVALUETYPE = "InvalidParameter.ErrInvalidActionParamBadValueType"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMNAME = "InvalidParameter.ErrInvalidActionParamName"
//  INVALIDPARAMETER_ERRINVALIDACTIONTYPE = "InvalidParameter.ErrInvalidActionType"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONIGNORECASE = "InvalidParameter.ErrInvalidConditionIgnoreCase"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMEBADNAME = "InvalidParameter.ErrInvalidConditionNameBadName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMETARGETNOTSUPPORTNAME = "InvalidParameter.ErrInvalidConditionNameTargetNotSupportName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADREGULAR = "InvalidParameter.ErrInvalidConditionValueBadRegular"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUE = "InvalidParameter.ErrInvalidConditionValueBadValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUECONTAINFILENAMEEXTENSION = "InvalidParameter.ErrInvalidConditionValueBadValueContainFileNameExtension"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOLONGVALUE = "InvalidParameter.ErrInvalidConditionValueTooLongValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYREGULAR = "InvalidParameter.ErrInvalidConditionValueTooManyRegular"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYVALUES = "InvalidParameter.ErrInvalidConditionValueTooManyValues"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYWILDCARD = "InvalidParameter.ErrInvalidConditionValueTooManyWildcard"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEZEROLENGTH = "InvalidParameter.ErrInvalidConditionValueZeroLength"
//  INVALIDPARAMETER_HOSTNOTFOUND = "InvalidParameter.HostNotFound"
//  INVALIDPARAMETER_INVALIDAUTHENTICATION = "InvalidParameter.InvalidAuthentication"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESIGNPARAM = "InvalidParameter.InvalidAuthenticationTypeSignParam"
//  INVALIDPARAMETER_INVALIDBACKUPSERVERNAME = "InvalidParameter.InvalidBackupServerName"
//  INVALIDPARAMETER_INVALIDCACHECONFIGCACHE = "InvalidParameter.InvalidCacheConfigCache"
//  INVALIDPARAMETER_INVALIDCACHEKEY = "InvalidParameter.InvalidCacheKey"
//  INVALIDPARAMETER_INVALIDCACHEKEYQUERYSTRINGVALUE = "InvalidParameter.InvalidCacheKeyQueryStringValue"
//  INVALIDPARAMETER_INVALIDCACHETIME = "InvalidParameter.InvalidCacheTime"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINE = "InvalidParameter.InvalidDynamicRoutine"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINEBILLING = "InvalidParameter.InvalidDynamicRoutineBilling"
//  INVALIDPARAMETER_INVALIDERRORPAGEREDIRECTURL = "InvalidParameter.InvalidErrorPageRedirectUrl"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDHTTPSHSTSMAXAGE = "InvalidParameter.InvalidHttpsHstsMaxAge"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPOSTSIZEVALUE = "InvalidParameter.InvalidPostSizeValue"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAMEXFF = "InvalidParameter.InvalidRequestHeaderNameXff"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERVALUE = "InvalidParameter.InvalidRequestHeaderValue"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERNAME = "InvalidParameter.InvalidResponseHeaderName"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERVALUE = "InvalidParameter.InvalidResponseHeaderValue"
//  INVALIDPARAMETER_INVALIDRULEENGINEACTION = "InvalidParameter.InvalidRuleEngineAction"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGET = "InvalidParameter.InvalidRuleEngineTarget"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSEXTENSION = "InvalidParameter.InvalidRuleEngineTargetsExtension"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSURL = "InvalidParameter.InvalidRuleEngineTargetsUrl"
//  INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGEXPIRETIMELIMIT = "InvalidParameter.InvalidStandardDebugExpireTimeLimit"
//  INVALIDPARAMETER_INVALIDUPSTREAMREQUESTQUERYSTRINGVALUE = "InvalidParameter.InvalidUpstreamRequestQueryStringValue"
//  INVALIDPARAMETER_INVALIDURLREDIRECTHOST = "InvalidParameter.InvalidUrlRedirectHost"
//  INVALIDPARAMETER_INVALIDURLREDIRECTURL = "InvalidParameter.InvalidUrlRedirectUrl"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_ORIGINORIGINGROUPIDISREQUIRED = "InvalidParameter.OriginOriginGroupIdIsRequired"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_TASKSYSTEMERROR = "InvalidParameter.TaskSystemError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_INVALIDADVANCEDDEFENSESECURITYTYPE = "OperationDenied.InvalidAdvancedDefenseSecurityType"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) CreateRuleWithContext(ctx context.Context, request *CreateRuleRequest) (response *CreateRuleResponse, err error) {
    if request == nil {
        request = NewCreateRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSecurityIPGroupRequest() (request *CreateSecurityIPGroupRequest) {
    request = &CreateSecurityIPGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateSecurityIPGroup")
    
    
    return
}

func NewCreateSecurityIPGroupResponse() (response *CreateSecurityIPGroupResponse) {
    response = &CreateSecurityIPGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSecurityIPGroup
// This API is used to create a security IP group.
//
// error code that may be returned:
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateSecurityIPGroup(request *CreateSecurityIPGroupRequest) (response *CreateSecurityIPGroupResponse, err error) {
    return c.CreateSecurityIPGroupWithContext(context.Background(), request)
}

// CreateSecurityIPGroup
// This API is used to create a security IP group.
//
// error code that may be returned:
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateSecurityIPGroupWithContext(ctx context.Context, request *CreateSecurityIPGroupRequest) (response *CreateSecurityIPGroupResponse, err error) {
    if request == nil {
        request = NewCreateSecurityIPGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSecurityIPGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSecurityIPGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateZoneRequest() (request *CreateZoneRequest) {
    request = &CreateZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateZone")
    
    
    return
}

func NewCreateZoneResponse() (response *CreateZoneResponse) {
    response = &CreateZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateZone
// This API is used to access a new site.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_LENGTHEXCEEDSLIMIT = "InvalidParameter.LengthExceedsLimit"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ZONESAMEASNAME = "InvalidParameterValue.ZoneSameAsName"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DOMAINISBLOCKED = "OperationDenied.DomainIsBlocked"
//  OPERATIONDENIED_RECORDISFORBIDDEN = "OperationDenied.RecordIsForbidden"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_ALIASDOMAIN = "ResourceInUse.AliasDomain"
//  RESOURCEINUSE_CNAME = "ResourceInUse.Cname"
//  RESOURCEINUSE_DNS = "ResourceInUse.Dns"
//  RESOURCEINUSE_GENERICHOST = "ResourceInUse.GenericHost"
//  RESOURCEINUSE_HOST = "ResourceInUse.Host"
//  RESOURCEINUSE_NS = "ResourceInUse.NS"
//  RESOURCEINUSE_OTHERS = "ResourceInUse.Others"
//  RESOURCEINUSE_OTHERSALIASDOMAIN = "ResourceInUse.OthersAliasDomain"
//  RESOURCEINUSE_OTHERSCNAME = "ResourceInUse.OthersCname"
//  RESOURCEINUSE_OTHERSHOST = "ResourceInUse.OthersHost"
//  RESOURCEINUSE_OTHERSNS = "ResourceInUse.OthersNS"
//  RESOURCEINUSE_SELFANDOTHERSCNAME = "ResourceInUse.SelfAndOthersCname"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) CreateZone(request *CreateZoneRequest) (response *CreateZoneResponse, err error) {
    return c.CreateZoneWithContext(context.Background(), request)
}

// CreateZone
// This API is used to access a new site.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_LENGTHEXCEEDSLIMIT = "InvalidParameter.LengthExceedsLimit"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ZONESAMEASNAME = "InvalidParameterValue.ZoneSameAsName"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DOMAINISBLOCKED = "OperationDenied.DomainIsBlocked"
//  OPERATIONDENIED_RECORDISFORBIDDEN = "OperationDenied.RecordIsForbidden"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_ALIASDOMAIN = "ResourceInUse.AliasDomain"
//  RESOURCEINUSE_CNAME = "ResourceInUse.Cname"
//  RESOURCEINUSE_DNS = "ResourceInUse.Dns"
//  RESOURCEINUSE_GENERICHOST = "ResourceInUse.GenericHost"
//  RESOURCEINUSE_HOST = "ResourceInUse.Host"
//  RESOURCEINUSE_NS = "ResourceInUse.NS"
//  RESOURCEINUSE_OTHERS = "ResourceInUse.Others"
//  RESOURCEINUSE_OTHERSALIASDOMAIN = "ResourceInUse.OthersAliasDomain"
//  RESOURCEINUSE_OTHERSCNAME = "ResourceInUse.OthersCname"
//  RESOURCEINUSE_OTHERSHOST = "ResourceInUse.OthersHost"
//  RESOURCEINUSE_OTHERSNS = "ResourceInUse.OthersNS"
//  RESOURCEINUSE_SELFANDOTHERSCNAME = "ResourceInUse.SelfAndOthersCname"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) CreateZoneWithContext(ctx context.Context, request *CreateZoneRequest) (response *CreateZoneResponse, err error) {
    if request == nil {
        request = NewCreateZoneRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateZoneResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAccelerationDomainsRequest() (request *DeleteAccelerationDomainsRequest) {
    request = &DeleteAccelerationDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteAccelerationDomains")
    
    
    return
}

func NewDeleteAccelerationDomainsResponse() (response *DeleteAccelerationDomainsResponse) {
    response = &DeleteAccelerationDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAccelerationDomains
// This API is used to batch remove accelerated domain names.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDDOMAINSTATUS = "InvalidParameterValue.InvalidDomainStatus"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE_DOMAINNOTFOUND = "ResourceUnavailable.DomainNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteAccelerationDomains(request *DeleteAccelerationDomainsRequest) (response *DeleteAccelerationDomainsResponse, err error) {
    return c.DeleteAccelerationDomainsWithContext(context.Background(), request)
}

// DeleteAccelerationDomains
// This API is used to batch remove accelerated domain names.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDDOMAINSTATUS = "InvalidParameterValue.InvalidDomainStatus"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE_DOMAINNOTFOUND = "ResourceUnavailable.DomainNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteAccelerationDomainsWithContext(ctx context.Context, request *DeleteAccelerationDomainsRequest) (response *DeleteAccelerationDomainsResponse, err error) {
    if request == nil {
        request = NewDeleteAccelerationDomainsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAccelerationDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAccelerationDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAliasDomainRequest() (request *DeleteAliasDomainRequest) {
    request = &DeleteAliasDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteAliasDomain")
    
    
    return
}

func NewDeleteAliasDomainResponse() (response *DeleteAliasDomainResponse) {
    response = &DeleteAliasDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAliasDomain
// This API is used to delete an alias domain name.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteAliasDomain(request *DeleteAliasDomainRequest) (response *DeleteAliasDomainResponse, err error) {
    return c.DeleteAliasDomainWithContext(context.Background(), request)
}

// DeleteAliasDomain
// This API is used to delete an alias domain name.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteAliasDomainWithContext(ctx context.Context, request *DeleteAliasDomainRequest) (response *DeleteAliasDomainResponse, err error) {
    if request == nil {
        request = NewDeleteAliasDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAliasDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAliasDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApplicationProxyRequest() (request *DeleteApplicationProxyRequest) {
    request = &DeleteApplicationProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteApplicationProxy")
    
    
    return
}

func NewDeleteApplicationProxyResponse() (response *DeleteApplicationProxyResponse) {
    response = &DeleteApplicationProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteApplicationProxy
// This API is used to delete an application proxy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteApplicationProxy(request *DeleteApplicationProxyRequest) (response *DeleteApplicationProxyResponse, err error) {
    return c.DeleteApplicationProxyWithContext(context.Background(), request)
}

// DeleteApplicationProxy
// This API is used to delete an application proxy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteApplicationProxyWithContext(ctx context.Context, request *DeleteApplicationProxyRequest) (response *DeleteApplicationProxyResponse, err error) {
    if request == nil {
        request = NewDeleteApplicationProxyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApplicationProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteApplicationProxyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApplicationProxyRuleRequest() (request *DeleteApplicationProxyRuleRequest) {
    request = &DeleteApplicationProxyRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteApplicationProxyRule")
    
    
    return
}

func NewDeleteApplicationProxyRuleResponse() (response *DeleteApplicationProxyRuleResponse) {
    response = &DeleteApplicationProxyRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteApplicationProxyRule
// This API is used to delete an application proxy rule.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteApplicationProxyRule(request *DeleteApplicationProxyRuleRequest) (response *DeleteApplicationProxyRuleResponse, err error) {
    return c.DeleteApplicationProxyRuleWithContext(context.Background(), request)
}

// DeleteApplicationProxyRule
// This API is used to delete an application proxy rule.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteApplicationProxyRuleWithContext(ctx context.Context, request *DeleteApplicationProxyRuleRequest) (response *DeleteApplicationProxyRuleResponse, err error) {
    if request == nil {
        request = NewDeleteApplicationProxyRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApplicationProxyRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteApplicationProxyRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteOriginGroupRequest() (request *DeleteOriginGroupRequest) {
    request = &DeleteOriginGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteOriginGroup")
    
    
    return
}

func NewDeleteOriginGroupResponse() (response *DeleteOriginGroupResponse) {
    response = &DeleteOriginGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteOriginGroup
// This API is used to delete an origin group.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteOriginGroup(request *DeleteOriginGroupRequest) (response *DeleteOriginGroupResponse, err error) {
    return c.DeleteOriginGroupWithContext(context.Background(), request)
}

// DeleteOriginGroup
// This API is used to delete an origin group.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteOriginGroupWithContext(ctx context.Context, request *DeleteOriginGroupRequest) (response *DeleteOriginGroupResponse, err error) {
    if request == nil {
        request = NewDeleteOriginGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteOriginGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteOriginGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRulesRequest() (request *DeleteRulesRequest) {
    request = &DeleteRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteRules")
    
    
    return
}

func NewDeleteRulesResponse() (response *DeleteRulesResponse) {
    response = &DeleteRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRules
// This API is used to batch delete rules from the rule engine.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_INVALIDHTTPS = "InvalidParameter.InvalidHttps"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDRULEENGINENOTFOUND = "InvalidParameter.InvalidRuleEngineNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteRules(request *DeleteRulesRequest) (response *DeleteRulesResponse, err error) {
    return c.DeleteRulesWithContext(context.Background(), request)
}

// DeleteRules
// This API is used to batch delete rules from the rule engine.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_INVALIDHTTPS = "InvalidParameter.InvalidHttps"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDRULEENGINENOTFOUND = "InvalidParameter.InvalidRuleEngineNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteRulesWithContext(ctx context.Context, request *DeleteRulesRequest) (response *DeleteRulesResponse, err error) {
    if request == nil {
        request = NewDeleteRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSecurityIPGroupRequest() (request *DeleteSecurityIPGroupRequest) {
    request = &DeleteSecurityIPGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteSecurityIPGroup")
    
    
    return
}

func NewDeleteSecurityIPGroupResponse() (response *DeleteSecurityIPGroupResponse) {
    response = &DeleteSecurityIPGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteSecurityIPGroup
// This API is used to delete a specified security IP group. Note that the security IP group cannot be deleted if it is referenced in a rule.
//
// error code that may be returned:
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteSecurityIPGroup(request *DeleteSecurityIPGroupRequest) (response *DeleteSecurityIPGroupResponse, err error) {
    return c.DeleteSecurityIPGroupWithContext(context.Background(), request)
}

// DeleteSecurityIPGroup
// This API is used to delete a specified security IP group. Note that the security IP group cannot be deleted if it is referenced in a rule.
//
// error code that may be returned:
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteSecurityIPGroupWithContext(ctx context.Context, request *DeleteSecurityIPGroupRequest) (response *DeleteSecurityIPGroupResponse, err error) {
    if request == nil {
        request = NewDeleteSecurityIPGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSecurityIPGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSecurityIPGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteZoneRequest() (request *DeleteZoneRequest) {
    request = &DeleteZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteZone")
    
    
    return
}

func NewDeleteZoneResponse() (response *DeleteZoneResponse) {
    response = &DeleteZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteZone
// This API is used to delete a site.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteZone(request *DeleteZoneRequest) (response *DeleteZoneResponse, err error) {
    return c.DeleteZoneWithContext(context.Background(), request)
}

// DeleteZone
// This API is used to delete a site.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteZoneWithContext(ctx context.Context, request *DeleteZoneRequest) (response *DeleteZoneResponse, err error) {
    if request == nil {
        request = NewDeleteZoneRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteZoneResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccelerationDomainsRequest() (request *DescribeAccelerationDomainsRequest) {
    request = &DescribeAccelerationDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeAccelerationDomains")
    
    
    return
}

func NewDescribeAccelerationDomainsResponse() (response *DescribeAccelerationDomainsResponse) {
    response = &DescribeAccelerationDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAccelerationDomains
// This API is used to query accelerated domain names. Paging, sorting and filtering are supported.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_DOMAINONTRAFFICSCHEDULING = "InvalidParameter.DomainOnTrafficScheduling"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeAccelerationDomains(request *DescribeAccelerationDomainsRequest) (response *DescribeAccelerationDomainsResponse, err error) {
    return c.DescribeAccelerationDomainsWithContext(context.Background(), request)
}

// DescribeAccelerationDomains
// This API is used to query accelerated domain names. Paging, sorting and filtering are supported.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_DOMAINONTRAFFICSCHEDULING = "InvalidParameter.DomainOnTrafficScheduling"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeAccelerationDomainsWithContext(ctx context.Context, request *DescribeAccelerationDomainsRequest) (response *DescribeAccelerationDomainsResponse, err error) {
    if request == nil {
        request = NewDescribeAccelerationDomainsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccelerationDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccelerationDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAliasDomainsRequest() (request *DescribeAliasDomainsRequest) {
    request = &DescribeAliasDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeAliasDomains")
    
    
    return
}

func NewDescribeAliasDomainsResponse() (response *DescribeAliasDomainsResponse) {
    response = &DescribeAliasDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAliasDomains
// This API is used to query the information of alias domain names.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeAliasDomains(request *DescribeAliasDomainsRequest) (response *DescribeAliasDomainsResponse, err error) {
    return c.DescribeAliasDomainsWithContext(context.Background(), request)
}

// DescribeAliasDomains
// This API is used to query the information of alias domain names.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeAliasDomainsWithContext(ctx context.Context, request *DescribeAliasDomainsRequest) (response *DescribeAliasDomainsResponse, err error) {
    if request == nil {
        request = NewDescribeAliasDomainsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAliasDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAliasDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationProxiesRequest() (request *DescribeApplicationProxiesRequest) {
    request = &DescribeApplicationProxiesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeApplicationProxies")
    
    
    return
}

func NewDescribeApplicationProxiesResponse() (response *DescribeApplicationProxiesResponse) {
    response = &DescribeApplicationProxiesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApplicationProxies
// This API is used to query the list of application proxies.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeApplicationProxies(request *DescribeApplicationProxiesRequest) (response *DescribeApplicationProxiesResponse, err error) {
    return c.DescribeApplicationProxiesWithContext(context.Background(), request)
}

// DescribeApplicationProxies
// This API is used to query the list of application proxies.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeApplicationProxiesWithContext(ctx context.Context, request *DescribeApplicationProxiesRequest) (response *DescribeApplicationProxiesResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationProxiesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplicationProxies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationProxiesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAvailablePlansRequest() (request *DescribeAvailablePlansRequest) {
    request = &DescribeAvailablePlansRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeAvailablePlans")
    
    
    return
}

func NewDescribeAvailablePlansResponse() (response *DescribeAvailablePlansResponse) {
    response = &DescribeAvailablePlansResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAvailablePlans
// This API is used to query plan options available for purchase.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
func (c *Client) DescribeAvailablePlans(request *DescribeAvailablePlansRequest) (response *DescribeAvailablePlansResponse, err error) {
    return c.DescribeAvailablePlansWithContext(context.Background(), request)
}

// DescribeAvailablePlans
// This API is used to query plan options available for purchase.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
func (c *Client) DescribeAvailablePlansWithContext(ctx context.Context, request *DescribeAvailablePlansRequest) (response *DescribeAvailablePlansResponse, err error) {
    if request == nil {
        request = NewDescribeAvailablePlansRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAvailablePlans require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAvailablePlansResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeContentQuotaRequest() (request *DescribeContentQuotaRequest) {
    request = &DescribeContentQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeContentQuota")
    
    
    return
}

func NewDescribeContentQuotaResponse() (response *DescribeContentQuotaResponse) {
    response = &DescribeContentQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeContentQuota
// This API is used to query content management quotas.
//
// error code that may be returned:
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeContentQuota(request *DescribeContentQuotaRequest) (response *DescribeContentQuotaResponse, err error) {
    return c.DescribeContentQuotaWithContext(context.Background(), request)
}

// DescribeContentQuota
// This API is used to query content management quotas.
//
// error code that may be returned:
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeContentQuotaWithContext(ctx context.Context, request *DescribeContentQuotaRequest) (response *DescribeContentQuotaResponse, err error) {
    if request == nil {
        request = NewDescribeContentQuotaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeContentQuota require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeContentQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSAttackDataRequest() (request *DescribeDDoSAttackDataRequest) {
    request = &DescribeDDoSAttackDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeDDoSAttackData")
    
    
    return
}

func NewDescribeDDoSAttackDataResponse() (response *DescribeDDoSAttackDataResponse) {
    response = &DescribeDDoSAttackDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDDoSAttackData
// This API is used to query the time-series data of DDoS attacks.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeDDoSAttackData(request *DescribeDDoSAttackDataRequest) (response *DescribeDDoSAttackDataResponse, err error) {
    return c.DescribeDDoSAttackDataWithContext(context.Background(), request)
}

// DescribeDDoSAttackData
// This API is used to query the time-series data of DDoS attacks.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeDDoSAttackDataWithContext(ctx context.Context, request *DescribeDDoSAttackDataRequest) (response *DescribeDDoSAttackDataResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSAttackDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDDoSAttackData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDDoSAttackDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSAttackEventRequest() (request *DescribeDDoSAttackEventRequest) {
    request = &DescribeDDoSAttackEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeDDoSAttackEvent")
    
    
    return
}

func NewDescribeDDoSAttackEventResponse() (response *DescribeDDoSAttackEventResponse) {
    response = &DescribeDDoSAttackEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDDoSAttackEvent
// This API is used to query DDoS attack events.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeDDoSAttackEvent(request *DescribeDDoSAttackEventRequest) (response *DescribeDDoSAttackEventResponse, err error) {
    return c.DescribeDDoSAttackEventWithContext(context.Background(), request)
}

// DescribeDDoSAttackEvent
// This API is used to query DDoS attack events.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeDDoSAttackEventWithContext(ctx context.Context, request *DescribeDDoSAttackEventRequest) (response *DescribeDDoSAttackEventResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSAttackEventRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDDoSAttackEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDDoSAttackEventResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSAttackTopDataRequest() (request *DescribeDDoSAttackTopDataRequest) {
    request = &DescribeDDoSAttackTopDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeDDoSAttackTopData")
    
    
    return
}

func NewDescribeDDoSAttackTopDataResponse() (response *DescribeDDoSAttackTopDataResponse) {
    response = &DescribeDDoSAttackTopDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDDoSAttackTopData
// This API is used to query the top-ranked DDoS attack data.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeDDoSAttackTopData(request *DescribeDDoSAttackTopDataRequest) (response *DescribeDDoSAttackTopDataResponse, err error) {
    return c.DescribeDDoSAttackTopDataWithContext(context.Background(), request)
}

// DescribeDDoSAttackTopData
// This API is used to query the top-ranked DDoS attack data.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeDDoSAttackTopDataWithContext(ctx context.Context, request *DescribeDDoSAttackTopDataRequest) (response *DescribeDDoSAttackTopDataResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSAttackTopDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDDoSAttackTopData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDDoSAttackTopDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDefaultCertificatesRequest() (request *DescribeDefaultCertificatesRequest) {
    request = &DescribeDefaultCertificatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeDefaultCertificates")
    
    
    return
}

func NewDescribeDefaultCertificatesResponse() (response *DescribeDefaultCertificatesResponse) {
    response = &DescribeDefaultCertificatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDefaultCertificates
// This API is used to query a list of default certificates.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE_ZONENOTFOUND = "ResourceUnavailable.ZoneNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeDefaultCertificates(request *DescribeDefaultCertificatesRequest) (response *DescribeDefaultCertificatesResponse, err error) {
    return c.DescribeDefaultCertificatesWithContext(context.Background(), request)
}

// DescribeDefaultCertificates
// This API is used to query a list of default certificates.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE_ZONENOTFOUND = "ResourceUnavailable.ZoneNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeDefaultCertificatesWithContext(ctx context.Context, request *DescribeDefaultCertificatesRequest) (response *DescribeDefaultCertificatesResponse, err error) {
    if request == nil {
        request = NewDescribeDefaultCertificatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDefaultCertificates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDefaultCertificatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostsSettingRequest() (request *DescribeHostsSettingRequest) {
    request = &DescribeHostsSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeHostsSetting")
    
    
    return
}

func NewDescribeHostsSettingResponse() (response *DescribeHostsSettingResponse) {
    response = &DescribeHostsSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeHostsSetting
// This API is used to query detailed domain name configuration.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeHostsSetting(request *DescribeHostsSettingRequest) (response *DescribeHostsSettingResponse, err error) {
    return c.DescribeHostsSettingWithContext(context.Background(), request)
}

// DescribeHostsSetting
// This API is used to query detailed domain name configuration.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeHostsSettingWithContext(ctx context.Context, request *DescribeHostsSettingRequest) (response *DescribeHostsSettingResponse, err error) {
    if request == nil {
        request = NewDescribeHostsSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHostsSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHostsSettingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIdentificationsRequest() (request *DescribeIdentificationsRequest) {
    request = &DescribeIdentificationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeIdentifications")
    
    
    return
}

func NewDescribeIdentificationsResponse() (response *DescribeIdentificationsResponse) {
    response = &DescribeIdentificationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIdentifications
// This API is used to query the verification information of a site.
//
// error code that may be returned:
//  OPERATIONDENIED_DOMAININSHARECNAMEGROUP = "OperationDenied.DomainInShareCnameGroup"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeIdentifications(request *DescribeIdentificationsRequest) (response *DescribeIdentificationsResponse, err error) {
    return c.DescribeIdentificationsWithContext(context.Background(), request)
}

// DescribeIdentifications
// This API is used to query the verification information of a site.
//
// error code that may be returned:
//  OPERATIONDENIED_DOMAININSHARECNAMEGROUP = "OperationDenied.DomainInShareCnameGroup"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeIdentificationsWithContext(ctx context.Context, request *DescribeIdentificationsRequest) (response *DescribeIdentificationsResponse, err error) {
    if request == nil {
        request = NewDescribeIdentificationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIdentifications require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIdentificationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOriginGroupRequest() (request *DescribeOriginGroupRequest) {
    request = &DescribeOriginGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeOriginGroup")
    
    
    return
}

func NewDescribeOriginGroupResponse() (response *DescribeOriginGroupResponse) {
    response = &DescribeOriginGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOriginGroup
// This API is used to obtain a list of origin groups.
//
// error code that may be returned:
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeOriginGroup(request *DescribeOriginGroupRequest) (response *DescribeOriginGroupResponse, err error) {
    return c.DescribeOriginGroupWithContext(context.Background(), request)
}

// DescribeOriginGroup
// This API is used to obtain a list of origin groups.
//
// error code that may be returned:
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeOriginGroupWithContext(ctx context.Context, request *DescribeOriginGroupRequest) (response *DescribeOriginGroupResponse, err error) {
    if request == nil {
        request = NewDescribeOriginGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOriginGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOriginGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOriginProtectionRequest() (request *DescribeOriginProtectionRequest) {
    request = &DescribeOriginProtectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeOriginProtection")
    
    
    return
}

func NewDescribeOriginProtectionResponse() (response *DescribeOriginProtectionResponse) {
    response = &DescribeOriginProtectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOriginProtection
// This API is used to query the origin protection configuration.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeOriginProtection(request *DescribeOriginProtectionRequest) (response *DescribeOriginProtectionResponse, err error) {
    return c.DescribeOriginProtectionWithContext(context.Background(), request)
}

// DescribeOriginProtection
// This API is used to query the origin protection configuration.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeOriginProtectionWithContext(ctx context.Context, request *DescribeOriginProtectionRequest) (response *DescribeOriginProtectionResponse, err error) {
    if request == nil {
        request = NewDescribeOriginProtectionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOriginProtection require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOriginProtectionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOverviewL7DataRequest() (request *DescribeOverviewL7DataRequest) {
    request = &DescribeOverviewL7DataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeOverviewL7Data")
    
    
    return
}

func NewDescribeOverviewL7DataResponse() (response *DescribeOverviewL7DataResponse) {
    response = &DescribeOverviewL7DataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOverviewL7Data
// This API is used to query the L7 traffic summary statistics recorded over time.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeOverviewL7Data(request *DescribeOverviewL7DataRequest) (response *DescribeOverviewL7DataResponse, err error) {
    return c.DescribeOverviewL7DataWithContext(context.Background(), request)
}

// DescribeOverviewL7Data
// This API is used to query the L7 traffic summary statistics recorded over time.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeOverviewL7DataWithContext(ctx context.Context, request *DescribeOverviewL7DataRequest) (response *DescribeOverviewL7DataResponse, err error) {
    if request == nil {
        request = NewDescribeOverviewL7DataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOverviewL7Data require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOverviewL7DataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrefetchTasksRequest() (request *DescribePrefetchTasksRequest) {
    request = &DescribePrefetchTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribePrefetchTasks")
    
    
    return
}

func NewDescribePrefetchTasksResponse() (response *DescribePrefetchTasksResponse) {
    response = &DescribePrefetchTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePrefetchTasks
// This API is used to query the pre-warming task status.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_DOMAINEMPTY = "UnauthorizedOperation.DomainEmpty"
func (c *Client) DescribePrefetchTasks(request *DescribePrefetchTasksRequest) (response *DescribePrefetchTasksResponse, err error) {
    return c.DescribePrefetchTasksWithContext(context.Background(), request)
}

// DescribePrefetchTasks
// This API is used to query the pre-warming task status.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_DOMAINEMPTY = "UnauthorizedOperation.DomainEmpty"
func (c *Client) DescribePrefetchTasksWithContext(ctx context.Context, request *DescribePrefetchTasksRequest) (response *DescribePrefetchTasksResponse, err error) {
    if request == nil {
        request = NewDescribePrefetchTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrefetchTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrefetchTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePurgeTasksRequest() (request *DescribePurgeTasksRequest) {
    request = &DescribePurgeTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribePurgeTasks")
    
    
    return
}

func NewDescribePurgeTasksResponse() (response *DescribePurgeTasksResponse) {
    response = &DescribePurgeTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePurgeTasks
// Querying the cache purging history
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribePurgeTasks(request *DescribePurgeTasksRequest) (response *DescribePurgeTasksResponse, err error) {
    return c.DescribePurgeTasksWithContext(context.Background(), request)
}

// DescribePurgeTasks
// Querying the cache purging history
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribePurgeTasksWithContext(ctx context.Context, request *DescribePurgeTasksRequest) (response *DescribePurgeTasksResponse, err error) {
    if request == nil {
        request = NewDescribePurgeTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePurgeTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePurgeTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRulesRequest() (request *DescribeRulesRequest) {
    request = &DescribeRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeRules")
    
    
    return
}

func NewDescribeRulesResponse() (response *DescribeRulesResponse) {
    response = &DescribeRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRules
// This API is used to query the rules in the rule engine.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeRules(request *DescribeRulesRequest) (response *DescribeRulesResponse, err error) {
    return c.DescribeRulesWithContext(context.Background(), request)
}

// DescribeRules
// This API is used to query the rules in the rule engine.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeRulesWithContext(ctx context.Context, request *DescribeRulesRequest) (response *DescribeRulesResponse, err error) {
    if request == nil {
        request = NewDescribeRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRulesSettingRequest() (request *DescribeRulesSettingRequest) {
    request = &DescribeRulesSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeRulesSetting")
    
    
    return
}

func NewDescribeRulesSettingResponse() (response *DescribeRulesSettingResponse) {
    response = &DescribeRulesSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRulesSetting
// This API is used to return the list of the settings of the rule engine that can be used for request match and their detailed recommended configuration information.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeRulesSetting(request *DescribeRulesSettingRequest) (response *DescribeRulesSettingResponse, err error) {
    return c.DescribeRulesSettingWithContext(context.Background(), request)
}

// DescribeRulesSetting
// This API is used to return the list of the settings of the rule engine that can be used for request match and their detailed recommended configuration information.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeRulesSettingWithContext(ctx context.Context, request *DescribeRulesSettingRequest) (response *DescribeRulesSettingResponse, err error) {
    if request == nil {
        request = NewDescribeRulesSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRulesSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRulesSettingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTimingL4DataRequest() (request *DescribeTimingL4DataRequest) {
    request = &DescribeTimingL4DataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeTimingL4Data")
    
    
    return
}

func NewDescribeTimingL4DataResponse() (response *DescribeTimingL4DataResponse) {
    response = &DescribeTimingL4DataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTimingL4Data
// This API is used to query the list of L4 traffic data recorded over time.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeTimingL4Data(request *DescribeTimingL4DataRequest) (response *DescribeTimingL4DataResponse, err error) {
    return c.DescribeTimingL4DataWithContext(context.Background(), request)
}

// DescribeTimingL4Data
// This API is used to query the list of L4 traffic data recorded over time.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeTimingL4DataWithContext(ctx context.Context, request *DescribeTimingL4DataRequest) (response *DescribeTimingL4DataResponse, err error) {
    if request == nil {
        request = NewDescribeTimingL4DataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTimingL4Data require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTimingL4DataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTimingL7AnalysisDataRequest() (request *DescribeTimingL7AnalysisDataRequest) {
    request = &DescribeTimingL7AnalysisDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeTimingL7AnalysisData")
    
    
    return
}

func NewDescribeTimingL7AnalysisDataResponse() (response *DescribeTimingL7AnalysisDataResponse) {
    response = &DescribeTimingL7AnalysisDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTimingL7AnalysisData
// This API is used to query the L7 data recorded over time.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeTimingL7AnalysisData(request *DescribeTimingL7AnalysisDataRequest) (response *DescribeTimingL7AnalysisDataResponse, err error) {
    return c.DescribeTimingL7AnalysisDataWithContext(context.Background(), request)
}

// DescribeTimingL7AnalysisData
// This API is used to query the L7 data recorded over time.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeTimingL7AnalysisDataWithContext(ctx context.Context, request *DescribeTimingL7AnalysisDataRequest) (response *DescribeTimingL7AnalysisDataResponse, err error) {
    if request == nil {
        request = NewDescribeTimingL7AnalysisDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTimingL7AnalysisData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTimingL7AnalysisDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTimingL7CacheDataRequest() (request *DescribeTimingL7CacheDataRequest) {
    request = &DescribeTimingL7CacheDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeTimingL7CacheData")
    
    
    return
}

func NewDescribeTimingL7CacheDataResponse() (response *DescribeTimingL7CacheDataResponse) {
    response = &DescribeTimingL7CacheDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTimingL7CacheData
// This API is used to query the time-series L7 cached data.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeTimingL7CacheData(request *DescribeTimingL7CacheDataRequest) (response *DescribeTimingL7CacheDataResponse, err error) {
    return c.DescribeTimingL7CacheDataWithContext(context.Background(), request)
}

// DescribeTimingL7CacheData
// This API is used to query the time-series L7 cached data.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeTimingL7CacheDataWithContext(ctx context.Context, request *DescribeTimingL7CacheDataRequest) (response *DescribeTimingL7CacheDataResponse, err error) {
    if request == nil {
        request = NewDescribeTimingL7CacheDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTimingL7CacheData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTimingL7CacheDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopL7AnalysisDataRequest() (request *DescribeTopL7AnalysisDataRequest) {
    request = &DescribeTopL7AnalysisDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeTopL7AnalysisData")
    
    
    return
}

func NewDescribeTopL7AnalysisDataResponse() (response *DescribeTopL7AnalysisDataResponse) {
    response = &DescribeTopL7AnalysisDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTopL7AnalysisData
// This API is used to query the top-ranked L7 traffic data.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeTopL7AnalysisData(request *DescribeTopL7AnalysisDataRequest) (response *DescribeTopL7AnalysisDataResponse, err error) {
    return c.DescribeTopL7AnalysisDataWithContext(context.Background(), request)
}

// DescribeTopL7AnalysisData
// This API is used to query the top-ranked L7 traffic data.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeTopL7AnalysisDataWithContext(ctx context.Context, request *DescribeTopL7AnalysisDataRequest) (response *DescribeTopL7AnalysisDataResponse, err error) {
    if request == nil {
        request = NewDescribeTopL7AnalysisDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopL7AnalysisData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopL7AnalysisDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopL7CacheDataRequest() (request *DescribeTopL7CacheDataRequest) {
    request = &DescribeTopL7CacheDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeTopL7CacheData")
    
    
    return
}

func NewDescribeTopL7CacheDataResponse() (response *DescribeTopL7CacheDataResponse) {
    response = &DescribeTopL7CacheDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTopL7CacheData
// This API is used to query the cached L7 top-ranked data.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeTopL7CacheData(request *DescribeTopL7CacheDataRequest) (response *DescribeTopL7CacheDataResponse, err error) {
    return c.DescribeTopL7CacheDataWithContext(context.Background(), request)
}

// DescribeTopL7CacheData
// This API is used to query the cached L7 top-ranked data.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeTopL7CacheDataWithContext(ctx context.Context, request *DescribeTopL7CacheDataRequest) (response *DescribeTopL7CacheDataResponse, err error) {
    if request == nil {
        request = NewDescribeTopL7CacheDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopL7CacheData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopL7CacheDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZoneSettingRequest() (request *DescribeZoneSettingRequest) {
    request = &DescribeZoneSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeZoneSetting")
    
    
    return
}

func NewDescribeZoneSettingResponse() (response *DescribeZoneSettingResponse) {
    response = &DescribeZoneSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeZoneSetting
// This API is used to query the site configuration.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_SETTINGINVALIDPARAM = "InvalidParameter.SettingInvalidParam"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeZoneSetting(request *DescribeZoneSettingRequest) (response *DescribeZoneSettingResponse, err error) {
    return c.DescribeZoneSettingWithContext(context.Background(), request)
}

// DescribeZoneSetting
// This API is used to query the site configuration.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_SETTINGINVALIDPARAM = "InvalidParameter.SettingInvalidParam"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeZoneSettingWithContext(ctx context.Context, request *DescribeZoneSettingRequest) (response *DescribeZoneSettingResponse, err error) {
    if request == nil {
        request = NewDescribeZoneSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeZoneSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeZoneSettingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZonesRequest() (request *DescribeZonesRequest) {
    request = &DescribeZonesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeZones")
    
    
    return
}

func NewDescribeZonesResponse() (response *DescribeZonesResponse) {
    response = &DescribeZonesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeZones
// This API is used to query the list of user sites.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeZones(request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
    return c.DescribeZonesWithContext(context.Background(), request)
}

// DescribeZones
// This API is used to query the list of user sites.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeZonesWithContext(ctx context.Context, request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
    if request == nil {
        request = NewDescribeZonesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeZones require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeZonesResponse()
    err = c.Send(request, response)
    return
}

func NewDownloadL4LogsRequest() (request *DownloadL4LogsRequest) {
    request = &DownloadL4LogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DownloadL4Logs")
    
    
    return
}

func NewDownloadL4LogsResponse() (response *DownloadL4LogsResponse) {
    response = &DownloadL4LogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DownloadL4Logs
// This API is used to download L4 logs.
//
// error code that may be returned:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DownloadL4Logs(request *DownloadL4LogsRequest) (response *DownloadL4LogsResponse, err error) {
    return c.DownloadL4LogsWithContext(context.Background(), request)
}

// DownloadL4Logs
// This API is used to download L4 logs.
//
// error code that may be returned:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DownloadL4LogsWithContext(ctx context.Context, request *DownloadL4LogsRequest) (response *DownloadL4LogsResponse, err error) {
    if request == nil {
        request = NewDownloadL4LogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DownloadL4Logs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDownloadL4LogsResponse()
    err = c.Send(request, response)
    return
}

func NewDownloadL7LogsRequest() (request *DownloadL7LogsRequest) {
    request = &DownloadL7LogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DownloadL7Logs")
    
    
    return
}

func NewDownloadL7LogsResponse() (response *DownloadL7LogsResponse) {
    response = &DownloadL7LogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DownloadL7Logs
// This API is used to download L7 logs.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DownloadL7Logs(request *DownloadL7LogsRequest) (response *DownloadL7LogsResponse, err error) {
    return c.DownloadL7LogsWithContext(context.Background(), request)
}

// DownloadL7Logs
// This API is used to download L7 logs.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DownloadL7LogsWithContext(ctx context.Context, request *DownloadL7LogsRequest) (response *DownloadL7LogsResponse, err error) {
    if request == nil {
        request = NewDownloadL7LogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DownloadL7Logs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDownloadL7LogsResponse()
    err = c.Send(request, response)
    return
}

func NewIdentifyZoneRequest() (request *IdentifyZoneRequest) {
    request = &IdentifyZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "IdentifyZone")
    
    
    return
}

func NewIdentifyZoneResponse() (response *IdentifyZoneResponse) {
    response = &IdentifyZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// IdentifyZone
// This API is used to verify ownership of the site.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) IdentifyZone(request *IdentifyZoneRequest) (response *IdentifyZoneResponse, err error) {
    return c.IdentifyZoneWithContext(context.Background(), request)
}

// IdentifyZone
// This API is used to verify ownership of the site.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) IdentifyZoneWithContext(ctx context.Context, request *IdentifyZoneRequest) (response *IdentifyZoneResponse, err error) {
    if request == nil {
        request = NewIdentifyZoneRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("IdentifyZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewIdentifyZoneResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccelerationDomainRequest() (request *ModifyAccelerationDomainRequest) {
    request = &ModifyAccelerationDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyAccelerationDomain")
    
    
    return
}

func NewModifyAccelerationDomainResponse() (response *ModifyAccelerationDomainResponse) {
    response = &ModifyAccelerationDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAccelerationDomain
// This API is used to modify an accelerated domain name.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_CONFLICTHOSTORIGIN = "InvalidParameter.ConflictHostOrigin"
//  INVALIDPARAMETER_INVALIDCLIENTIPORIGIN = "InvalidParameter.InvalidClientIpOrigin"
//  INVALIDPARAMETER_INVALIDHTTPS = "InvalidParameter.InvalidHttps"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_ORIGINISINNERIP = "InvalidParameter.OriginIsInnerIp"
//  INVALIDPARAMETERVALUE_CONFLICTRECORD = "InvalidParameterValue.ConflictRecord"
//  INVALIDPARAMETERVALUE_DOMAINNOTMATCHZONE = "InvalidParameterValue.DomainNotMatchZone"
//  INVALIDPARAMETERVALUE_INVALIDDOMAINSTATUS = "InvalidParameterValue.InvalidDomainStatus"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_DNSRECORD = "ResourceInUse.DnsRecord"
//  RESOURCEUNAVAILABLE_DOMAINNOTFOUND = "ResourceUnavailable.DomainNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) ModifyAccelerationDomain(request *ModifyAccelerationDomainRequest) (response *ModifyAccelerationDomainResponse, err error) {
    return c.ModifyAccelerationDomainWithContext(context.Background(), request)
}

// ModifyAccelerationDomain
// This API is used to modify an accelerated domain name.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_CONFLICTHOSTORIGIN = "InvalidParameter.ConflictHostOrigin"
//  INVALIDPARAMETER_INVALIDCLIENTIPORIGIN = "InvalidParameter.InvalidClientIpOrigin"
//  INVALIDPARAMETER_INVALIDHTTPS = "InvalidParameter.InvalidHttps"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_ORIGINISINNERIP = "InvalidParameter.OriginIsInnerIp"
//  INVALIDPARAMETERVALUE_CONFLICTRECORD = "InvalidParameterValue.ConflictRecord"
//  INVALIDPARAMETERVALUE_DOMAINNOTMATCHZONE = "InvalidParameterValue.DomainNotMatchZone"
//  INVALIDPARAMETERVALUE_INVALIDDOMAINSTATUS = "InvalidParameterValue.InvalidDomainStatus"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_DNSRECORD = "ResourceInUse.DnsRecord"
//  RESOURCEUNAVAILABLE_DOMAINNOTFOUND = "ResourceUnavailable.DomainNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) ModifyAccelerationDomainWithContext(ctx context.Context, request *ModifyAccelerationDomainRequest) (response *ModifyAccelerationDomainResponse, err error) {
    if request == nil {
        request = NewModifyAccelerationDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccelerationDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAccelerationDomainResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccelerationDomainStatusesRequest() (request *ModifyAccelerationDomainStatusesRequest) {
    request = &ModifyAccelerationDomainStatusesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyAccelerationDomainStatuses")
    
    
    return
}

func NewModifyAccelerationDomainStatusesResponse() (response *ModifyAccelerationDomainStatusesResponse) {
    response = &ModifyAccelerationDomainStatusesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAccelerationDomainStatuses
// This API is used to batch modify the status of accelerated domains.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INVALIDERRORPAGE = "InvalidParameter.InvalidErrorPage"
//  INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"
//  INVALIDPARAMETERVALUE_DOMAINNOTMATCHZONE = "InvalidParameterValue.DomainNotMatchZone"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE_DOMAINNOTFOUND = "ResourceUnavailable.DomainNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ModifyAccelerationDomainStatuses(request *ModifyAccelerationDomainStatusesRequest) (response *ModifyAccelerationDomainStatusesResponse, err error) {
    return c.ModifyAccelerationDomainStatusesWithContext(context.Background(), request)
}

// ModifyAccelerationDomainStatuses
// This API is used to batch modify the status of accelerated domains.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INVALIDERRORPAGE = "InvalidParameter.InvalidErrorPage"
//  INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"
//  INVALIDPARAMETERVALUE_DOMAINNOTMATCHZONE = "InvalidParameterValue.DomainNotMatchZone"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE_DOMAINNOTFOUND = "ResourceUnavailable.DomainNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ModifyAccelerationDomainStatusesWithContext(ctx context.Context, request *ModifyAccelerationDomainStatusesRequest) (response *ModifyAccelerationDomainStatusesResponse, err error) {
    if request == nil {
        request = NewModifyAccelerationDomainStatusesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccelerationDomainStatuses require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAccelerationDomainStatusesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAliasDomainRequest() (request *ModifyAliasDomainRequest) {
    request = &ModifyAliasDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyAliasDomain")
    
    
    return
}

func NewModifyAliasDomainResponse() (response *ModifyAliasDomainResponse) {
    response = &ModifyAliasDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAliasDomain
// This API is used to modify an alias domain name.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CERTNOTMATCHDOMAIN = "InvalidParameter.CertNotMatchDomain"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyAliasDomain(request *ModifyAliasDomainRequest) (response *ModifyAliasDomainResponse, err error) {
    return c.ModifyAliasDomainWithContext(context.Background(), request)
}

// ModifyAliasDomain
// This API is used to modify an alias domain name.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CERTNOTMATCHDOMAIN = "InvalidParameter.CertNotMatchDomain"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyAliasDomainWithContext(ctx context.Context, request *ModifyAliasDomainRequest) (response *ModifyAliasDomainResponse, err error) {
    if request == nil {
        request = NewModifyAliasDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAliasDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAliasDomainResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAliasDomainStatusRequest() (request *ModifyAliasDomainStatusRequest) {
    request = &ModifyAliasDomainStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyAliasDomainStatus")
    
    
    return
}

func NewModifyAliasDomainStatusResponse() (response *ModifyAliasDomainStatusResponse) {
    response = &ModifyAliasDomainStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAliasDomainStatus
// This API is used to modify the status of an alias domain name.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) ModifyAliasDomainStatus(request *ModifyAliasDomainStatusRequest) (response *ModifyAliasDomainStatusResponse, err error) {
    return c.ModifyAliasDomainStatusWithContext(context.Background(), request)
}

// ModifyAliasDomainStatus
// This API is used to modify the status of an alias domain name.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) ModifyAliasDomainStatusWithContext(ctx context.Context, request *ModifyAliasDomainStatusRequest) (response *ModifyAliasDomainStatusResponse, err error) {
    if request == nil {
        request = NewModifyAliasDomainStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAliasDomainStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAliasDomainStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationProxyRequest() (request *ModifyApplicationProxyRequest) {
    request = &ModifyApplicationProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyApplicationProxy")
    
    
    return
}

func NewModifyApplicationProxyResponse() (response *ModifyApplicationProxyResponse) {
    response = &ModifyApplicationProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyApplicationProxy
// This API is used to modify an application proxy.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_L4PROXYINBANNEDSTATUS = "OperationDenied.L4ProxyInBannedStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyApplicationProxy(request *ModifyApplicationProxyRequest) (response *ModifyApplicationProxyResponse, err error) {
    return c.ModifyApplicationProxyWithContext(context.Background(), request)
}

// ModifyApplicationProxy
// This API is used to modify an application proxy.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_L4PROXYINBANNEDSTATUS = "OperationDenied.L4ProxyInBannedStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyApplicationProxyWithContext(ctx context.Context, request *ModifyApplicationProxyRequest) (response *ModifyApplicationProxyResponse, err error) {
    if request == nil {
        request = NewModifyApplicationProxyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplicationProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApplicationProxyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationProxyRuleRequest() (request *ModifyApplicationProxyRuleRequest) {
    request = &ModifyApplicationProxyRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyApplicationProxyRule")
    
    
    return
}

func NewModifyApplicationProxyRuleResponse() (response *ModifyApplicationProxyRuleResponse) {
    response = &ModifyApplicationProxyRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyApplicationProxyRule
// This API is used to modify an application proxy rule.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_L4PROXYINBANNEDSTATUS = "OperationDenied.L4ProxyInBannedStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyApplicationProxyRule(request *ModifyApplicationProxyRuleRequest) (response *ModifyApplicationProxyRuleResponse, err error) {
    return c.ModifyApplicationProxyRuleWithContext(context.Background(), request)
}

// ModifyApplicationProxyRule
// This API is used to modify an application proxy rule.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_L4PROXYINBANNEDSTATUS = "OperationDenied.L4ProxyInBannedStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyApplicationProxyRuleWithContext(ctx context.Context, request *ModifyApplicationProxyRuleRequest) (response *ModifyApplicationProxyRuleResponse, err error) {
    if request == nil {
        request = NewModifyApplicationProxyRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplicationProxyRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApplicationProxyRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationProxyRuleStatusRequest() (request *ModifyApplicationProxyRuleStatusRequest) {
    request = &ModifyApplicationProxyRuleStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyApplicationProxyRuleStatus")
    
    
    return
}

func NewModifyApplicationProxyRuleStatusResponse() (response *ModifyApplicationProxyRuleStatusResponse) {
    response = &ModifyApplicationProxyRuleStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyApplicationProxyRuleStatus
// This API is used to modify the status of an application proxy rule.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_L4PROXYINBANNEDSTATUS = "OperationDenied.L4ProxyInBannedStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyApplicationProxyRuleStatus(request *ModifyApplicationProxyRuleStatusRequest) (response *ModifyApplicationProxyRuleStatusResponse, err error) {
    return c.ModifyApplicationProxyRuleStatusWithContext(context.Background(), request)
}

// ModifyApplicationProxyRuleStatus
// This API is used to modify the status of an application proxy rule.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_L4PROXYINBANNEDSTATUS = "OperationDenied.L4ProxyInBannedStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyApplicationProxyRuleStatusWithContext(ctx context.Context, request *ModifyApplicationProxyRuleStatusRequest) (response *ModifyApplicationProxyRuleStatusResponse, err error) {
    if request == nil {
        request = NewModifyApplicationProxyRuleStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplicationProxyRuleStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApplicationProxyRuleStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationProxyStatusRequest() (request *ModifyApplicationProxyStatusRequest) {
    request = &ModifyApplicationProxyStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyApplicationProxyStatus")
    
    
    return
}

func NewModifyApplicationProxyStatusResponse() (response *ModifyApplicationProxyStatusResponse) {
    response = &ModifyApplicationProxyStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyApplicationProxyStatus
// This API is used to modify the status of an application proxy.
//
// error code that may be returned:
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_L4PROXYINBANNEDSTATUS = "OperationDenied.L4ProxyInBannedStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyApplicationProxyStatus(request *ModifyApplicationProxyStatusRequest) (response *ModifyApplicationProxyStatusResponse, err error) {
    return c.ModifyApplicationProxyStatusWithContext(context.Background(), request)
}

// ModifyApplicationProxyStatus
// This API is used to modify the status of an application proxy.
//
// error code that may be returned:
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_L4PROXYINBANNEDSTATUS = "OperationDenied.L4ProxyInBannedStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyApplicationProxyStatusWithContext(ctx context.Context, request *ModifyApplicationProxyStatusRequest) (response *ModifyApplicationProxyStatusResponse, err error) {
    if request == nil {
        request = NewModifyApplicationProxyStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplicationProxyStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApplicationProxyStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyHostsCertificateRequest() (request *ModifyHostsCertificateRequest) {
    request = &ModifyHostsCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyHostsCertificate")
    
    
    return
}

func NewModifyHostsCertificateResponse() (response *ModifyHostsCertificateResponse) {
    response = &ModifyHostsCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyHostsCertificate
// This API is used to modify the certificate of a domain name.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDZONESTATUS = "FailedOperation.InvalidZoneStatus"
//  FAILEDOPERATION_MODIFYFAILED = "FailedOperation.ModifyFailed"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWERROR = "InternalError.UnknowError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_CERTNOTMATCHDOMAIN = "InvalidParameter.CertNotMatchDomain"
//  INVALIDPARAMETER_CERTTOEXPIRE = "InvalidParameter.CertToExpire"
//  INVALIDPARAMETER_CERTTOOSHORTKEYSIZE = "InvalidParameter.CertTooShortKeySize"
//  INVALIDPARAMETER_CNAMEWILDHOSTNOTALLOWAPPLYCERTIFICATE = "InvalidParameter.CnameWildHostNotAllowApplyCertificate"
//  INVALIDPARAMETER_HOSTSTATUSNOTALLOWAPPLYCERTIFICATE = "InvalidParameter.HostStatusNotAllowApplyCertificate"
//  INVALIDPARAMETER_INVALIDCERTINFO = "InvalidParameter.InvalidCertInfo"
//  INVALIDPARAMETER_INVALIDHTTPSCERTINFO = "InvalidParameter.InvalidHttpsCertInfo"
//  INVALIDPARAMETER_INVALIDHTTPSTLSVERSION = "InvalidParameter.InvalidHttpsTlsVersion"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE_CERTNOTFOUND = "ResourceUnavailable.CertNotFound"
//  RESOURCEUNAVAILABLE_HOSTNOTFOUND = "ResourceUnavailable.HostNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) ModifyHostsCertificate(request *ModifyHostsCertificateRequest) (response *ModifyHostsCertificateResponse, err error) {
    return c.ModifyHostsCertificateWithContext(context.Background(), request)
}

// ModifyHostsCertificate
// This API is used to modify the certificate of a domain name.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDZONESTATUS = "FailedOperation.InvalidZoneStatus"
//  FAILEDOPERATION_MODIFYFAILED = "FailedOperation.ModifyFailed"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWERROR = "InternalError.UnknowError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_CERTNOTMATCHDOMAIN = "InvalidParameter.CertNotMatchDomain"
//  INVALIDPARAMETER_CERTTOEXPIRE = "InvalidParameter.CertToExpire"
//  INVALIDPARAMETER_CERTTOOSHORTKEYSIZE = "InvalidParameter.CertTooShortKeySize"
//  INVALIDPARAMETER_CNAMEWILDHOSTNOTALLOWAPPLYCERTIFICATE = "InvalidParameter.CnameWildHostNotAllowApplyCertificate"
//  INVALIDPARAMETER_HOSTSTATUSNOTALLOWAPPLYCERTIFICATE = "InvalidParameter.HostStatusNotAllowApplyCertificate"
//  INVALIDPARAMETER_INVALIDCERTINFO = "InvalidParameter.InvalidCertInfo"
//  INVALIDPARAMETER_INVALIDHTTPSCERTINFO = "InvalidParameter.InvalidHttpsCertInfo"
//  INVALIDPARAMETER_INVALIDHTTPSTLSVERSION = "InvalidParameter.InvalidHttpsTlsVersion"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE_CERTNOTFOUND = "ResourceUnavailable.CertNotFound"
//  RESOURCEUNAVAILABLE_HOSTNOTFOUND = "ResourceUnavailable.HostNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) ModifyHostsCertificateWithContext(ctx context.Context, request *ModifyHostsCertificateRequest) (response *ModifyHostsCertificateResponse, err error) {
    if request == nil {
        request = NewModifyHostsCertificateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyHostsCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyHostsCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyOriginGroupRequest() (request *ModifyOriginGroupRequest) {
    request = &ModifyOriginGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyOriginGroup")
    
    
    return
}

func NewModifyOriginGroupResponse() (response *ModifyOriginGroupResponse) {
    response = &ModifyOriginGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyOriginGroup
// This API is used to modify an origin group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_L4STATUSNOTINONLINE = "OperationDenied.L4StatusNotInOnline"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ModifyOriginGroup(request *ModifyOriginGroupRequest) (response *ModifyOriginGroupResponse, err error) {
    return c.ModifyOriginGroupWithContext(context.Background(), request)
}

// ModifyOriginGroup
// This API is used to modify an origin group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_L4STATUSNOTINONLINE = "OperationDenied.L4StatusNotInOnline"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ModifyOriginGroupWithContext(ctx context.Context, request *ModifyOriginGroupRequest) (response *ModifyOriginGroupResponse, err error) {
    if request == nil {
        request = NewModifyOriginGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyOriginGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyOriginGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRuleRequest() (request *ModifyRuleRequest) {
    request = &ModifyRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyRule")
    
    
    return
}

func NewModifyRuleResponse() (response *ModifyRuleResponse) {
    response = &ModifyRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyRule
// This API is used to modify a rule in the rule engine.
//
// error code that may be returned:
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_CACHEKEYQUERYSTRINGTOOMANYVALUE = "InvalidParameter.CacheKeyQueryStringTooManyValue"
//  INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"
//  INVALIDPARAMETER_ERRACTIONUNSUPPORTTARGET = "InvalidParameter.ErrActionUnsupportTarget"
//  INVALIDPARAMETER_ERRINVALIDACTION = "InvalidParameter.ErrInvalidAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMACTION = "InvalidParameter.ErrInvalidActionParamAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMDUPLICATENAME = "InvalidParameter.ErrInvalidActionParamDuplicateName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMTOOMANYVALUES = "InvalidParameter.ErrInvalidActionParamTooManyValues"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONIGNORECASE = "InvalidParameter.ErrInvalidConditionIgnoreCase"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMETARGETNOTSUPPORTNAME = "InvalidParameter.ErrInvalidConditionNameTargetNotSupportName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADREGULAR = "InvalidParameter.ErrInvalidConditionValueBadRegular"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADURL = "InvalidParameter.ErrInvalidConditionValueBadUrl"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUE = "InvalidParameter.ErrInvalidConditionValueBadValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUECONTAINFILENAMEEXTENSION = "InvalidParameter.ErrInvalidConditionValueBadValueContainFileNameExtension"
//  INVALIDPARAMETER_HOSTNOTFOUND = "InvalidParameter.HostNotFound"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESECRETKEY = "InvalidParameter.InvalidAuthenticationTypeSecretKey"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESIGNPARAM = "InvalidParameter.InvalidAuthenticationTypeSignParam"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPETIMEFORMAT = "InvalidParameter.InvalidAuthenticationTypeTimeFormat"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPETIMEPARAM = "InvalidParameter.InvalidAuthenticationTypeTimeParam"
//  INVALIDPARAMETER_INVALIDBACKUPSERVERNAME = "InvalidParameter.InvalidBackupServerName"
//  INVALIDPARAMETER_INVALIDCACHEKEY = "InvalidParameter.InvalidCacheKey"
//  INVALIDPARAMETER_INVALIDCACHEKEYIGNORECASE = "InvalidParameter.InvalidCacheKeyIgnoreCase"
//  INVALIDPARAMETER_INVALIDCACHEKEYSCHEME = "InvalidParameter.InvalidCacheKeyScheme"
//  INVALIDPARAMETER_INVALIDCACHETIME = "InvalidParameter.InvalidCacheTime"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDERRORPAGEREDIRECTURL = "InvalidParameter.InvalidErrorPageRedirectUrl"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDMAXAGETIME = "InvalidParameter.InvalidMaxAgeTime"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAMEXFF = "InvalidParameter.InvalidRequestHeaderNameXff"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERVALUE = "InvalidParameter.InvalidRequestHeaderValue"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERNAME = "InvalidParameter.InvalidResponseHeaderName"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERVALUE = "InvalidParameter.InvalidResponseHeaderValue"
//  INVALIDPARAMETER_INVALIDRULEENGINEACTION = "InvalidParameter.InvalidRuleEngineAction"
//  INVALIDPARAMETER_INVALIDRULEENGINENOTFOUND = "InvalidParameter.InvalidRuleEngineNotFound"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGET = "InvalidParameter.InvalidRuleEngineTarget"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSEXTENSION = "InvalidParameter.InvalidRuleEngineTargetsExtension"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSURL = "InvalidParameter.InvalidRuleEngineTargetsUrl"
//  INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"
//  INVALIDPARAMETER_INVALIDUPSTREAMREQUESTQUERYSTRINGVALUE = "InvalidParameter.InvalidUpstreamRequestQueryStringValue"
//  INVALIDPARAMETER_INVALIDURLREDIRECTHOST = "InvalidParameter.InvalidUrlRedirectHost"
//  INVALIDPARAMETER_INVALIDURLREDIRECTURL = "InvalidParameter.InvalidUrlRedirectUrl"
//  INVALIDPARAMETER_KEYRULESINVALIDQUERYSTRINGVALUE = "InvalidParameter.KeyRulesInvalidQueryStringValue"
//  INVALIDPARAMETER_NOTSUPPORTTHISPRESET = "InvalidParameter.NotSupportThisPreset"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_INVALIDADVANCEDDEFENSESECURITYTYPE = "OperationDenied.InvalidAdvancedDefenseSecurityType"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ModifyRule(request *ModifyRuleRequest) (response *ModifyRuleResponse, err error) {
    return c.ModifyRuleWithContext(context.Background(), request)
}

// ModifyRule
// This API is used to modify a rule in the rule engine.
//
// error code that may be returned:
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_CACHEKEYQUERYSTRINGTOOMANYVALUE = "InvalidParameter.CacheKeyQueryStringTooManyValue"
//  INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"
//  INVALIDPARAMETER_ERRACTIONUNSUPPORTTARGET = "InvalidParameter.ErrActionUnsupportTarget"
//  INVALIDPARAMETER_ERRINVALIDACTION = "InvalidParameter.ErrInvalidAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMACTION = "InvalidParameter.ErrInvalidActionParamAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMDUPLICATENAME = "InvalidParameter.ErrInvalidActionParamDuplicateName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMTOOMANYVALUES = "InvalidParameter.ErrInvalidActionParamTooManyValues"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONIGNORECASE = "InvalidParameter.ErrInvalidConditionIgnoreCase"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMETARGETNOTSUPPORTNAME = "InvalidParameter.ErrInvalidConditionNameTargetNotSupportName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADREGULAR = "InvalidParameter.ErrInvalidConditionValueBadRegular"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADURL = "InvalidParameter.ErrInvalidConditionValueBadUrl"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUE = "InvalidParameter.ErrInvalidConditionValueBadValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUECONTAINFILENAMEEXTENSION = "InvalidParameter.ErrInvalidConditionValueBadValueContainFileNameExtension"
//  INVALIDPARAMETER_HOSTNOTFOUND = "InvalidParameter.HostNotFound"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESECRETKEY = "InvalidParameter.InvalidAuthenticationTypeSecretKey"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESIGNPARAM = "InvalidParameter.InvalidAuthenticationTypeSignParam"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPETIMEFORMAT = "InvalidParameter.InvalidAuthenticationTypeTimeFormat"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPETIMEPARAM = "InvalidParameter.InvalidAuthenticationTypeTimeParam"
//  INVALIDPARAMETER_INVALIDBACKUPSERVERNAME = "InvalidParameter.InvalidBackupServerName"
//  INVALIDPARAMETER_INVALIDCACHEKEY = "InvalidParameter.InvalidCacheKey"
//  INVALIDPARAMETER_INVALIDCACHEKEYIGNORECASE = "InvalidParameter.InvalidCacheKeyIgnoreCase"
//  INVALIDPARAMETER_INVALIDCACHEKEYSCHEME = "InvalidParameter.InvalidCacheKeyScheme"
//  INVALIDPARAMETER_INVALIDCACHETIME = "InvalidParameter.InvalidCacheTime"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDERRORPAGEREDIRECTURL = "InvalidParameter.InvalidErrorPageRedirectUrl"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDMAXAGETIME = "InvalidParameter.InvalidMaxAgeTime"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAMEXFF = "InvalidParameter.InvalidRequestHeaderNameXff"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERVALUE = "InvalidParameter.InvalidRequestHeaderValue"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERNAME = "InvalidParameter.InvalidResponseHeaderName"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERVALUE = "InvalidParameter.InvalidResponseHeaderValue"
//  INVALIDPARAMETER_INVALIDRULEENGINEACTION = "InvalidParameter.InvalidRuleEngineAction"
//  INVALIDPARAMETER_INVALIDRULEENGINENOTFOUND = "InvalidParameter.InvalidRuleEngineNotFound"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGET = "InvalidParameter.InvalidRuleEngineTarget"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSEXTENSION = "InvalidParameter.InvalidRuleEngineTargetsExtension"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSURL = "InvalidParameter.InvalidRuleEngineTargetsUrl"
//  INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"
//  INVALIDPARAMETER_INVALIDUPSTREAMREQUESTQUERYSTRINGVALUE = "InvalidParameter.InvalidUpstreamRequestQueryStringValue"
//  INVALIDPARAMETER_INVALIDURLREDIRECTHOST = "InvalidParameter.InvalidUrlRedirectHost"
//  INVALIDPARAMETER_INVALIDURLREDIRECTURL = "InvalidParameter.InvalidUrlRedirectUrl"
//  INVALIDPARAMETER_KEYRULESINVALIDQUERYSTRINGVALUE = "InvalidParameter.KeyRulesInvalidQueryStringValue"
//  INVALIDPARAMETER_NOTSUPPORTTHISPRESET = "InvalidParameter.NotSupportThisPreset"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_INVALIDADVANCEDDEFENSESECURITYTYPE = "OperationDenied.InvalidAdvancedDefenseSecurityType"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ModifyRuleWithContext(ctx context.Context, request *ModifyRuleRequest) (response *ModifyRuleResponse, err error) {
    if request == nil {
        request = NewModifyRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifySecurityIPGroupRequest() (request *ModifySecurityIPGroupRequest) {
    request = &ModifySecurityIPGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifySecurityIPGroup")
    
    
    return
}

func NewModifySecurityIPGroupResponse() (response *ModifySecurityIPGroupResponse) {
    response = &ModifySecurityIPGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySecurityIPGroup
// This API is used to modify a security IP group.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ModifySecurityIPGroup(request *ModifySecurityIPGroupRequest) (response *ModifySecurityIPGroupResponse, err error) {
    return c.ModifySecurityIPGroupWithContext(context.Background(), request)
}

// ModifySecurityIPGroup
// This API is used to modify a security IP group.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ModifySecurityIPGroupWithContext(ctx context.Context, request *ModifySecurityIPGroupRequest) (response *ModifySecurityIPGroupResponse, err error) {
    if request == nil {
        request = NewModifySecurityIPGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySecurityIPGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySecurityIPGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifySecurityPolicyRequest() (request *ModifySecurityPolicyRequest) {
    request = &ModifySecurityPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifySecurityPolicy")
    
    
    return
}

func NewModifySecurityPolicyResponse() (response *ModifySecurityPolicyResponse) {
    response = &ModifySecurityPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySecurityPolicy
// This API is used to modify the web and bot security configurations.
//
// error code that may be returned:
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifySecurityPolicy(request *ModifySecurityPolicyRequest) (response *ModifySecurityPolicyResponse, err error) {
    return c.ModifySecurityPolicyWithContext(context.Background(), request)
}

// ModifySecurityPolicy
// This API is used to modify the web and bot security configurations.
//
// error code that may be returned:
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifySecurityPolicyWithContext(ctx context.Context, request *ModifySecurityPolicyRequest) (response *ModifySecurityPolicyResponse, err error) {
    if request == nil {
        request = NewModifySecurityPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySecurityPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySecurityPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyZoneRequest() (request *ModifyZoneRequest) {
    request = &ModifyZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyZone")
    
    
    return
}

func NewModifyZoneResponse() (response *ModifyZoneResponse) {
    response = &ModifyZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyZone
// This API is used to modify a site.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_LENGTHEXCEEDSLIMIT = "InvalidParameter.LengthExceedsLimit"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ZONESAMEASNAME = "InvalidParameterValue.ZoneSameAsName"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DOMAININSHARECNAMEGROUP = "OperationDenied.DomainInShareCnameGroup"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_MULTIPLECNAMEZONE = "OperationDenied.MultipleCnameZone"
//  OPERATIONDENIED_NSNOTALLOWTRAFFICSTRATEGY = "OperationDenied.NSNotAllowTrafficStrategy"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  RESOURCEINUSE_NS = "ResourceInUse.NS"
//  RESOURCEINUSE_OTHERS = "ResourceInUse.Others"
//  RESOURCEINUSE_OTHERSALIASDOMAIN = "ResourceInUse.OthersAliasDomain"
//  RESOURCEINUSE_OTHERSNS = "ResourceInUse.OthersNS"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) ModifyZone(request *ModifyZoneRequest) (response *ModifyZoneResponse, err error) {
    return c.ModifyZoneWithContext(context.Background(), request)
}

// ModifyZone
// This API is used to modify a site.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_LENGTHEXCEEDSLIMIT = "InvalidParameter.LengthExceedsLimit"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ZONESAMEASNAME = "InvalidParameterValue.ZoneSameAsName"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DOMAININSHARECNAMEGROUP = "OperationDenied.DomainInShareCnameGroup"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_MULTIPLECNAMEZONE = "OperationDenied.MultipleCnameZone"
//  OPERATIONDENIED_NSNOTALLOWTRAFFICSTRATEGY = "OperationDenied.NSNotAllowTrafficStrategy"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  RESOURCEINUSE_NS = "ResourceInUse.NS"
//  RESOURCEINUSE_OTHERS = "ResourceInUse.Others"
//  RESOURCEINUSE_OTHERSALIASDOMAIN = "ResourceInUse.OthersAliasDomain"
//  RESOURCEINUSE_OTHERSNS = "ResourceInUse.OthersNS"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) ModifyZoneWithContext(ctx context.Context, request *ModifyZoneRequest) (response *ModifyZoneResponse, err error) {
    if request == nil {
        request = NewModifyZoneRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyZoneResponse()
    err = c.Send(request, response)
    return
}

func NewModifyZoneSettingRequest() (request *ModifyZoneSettingRequest) {
    request = &ModifyZoneSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyZoneSetting")
    
    
    return
}

func NewModifyZoneSettingResponse() (response *ModifyZoneSettingResponse) {
    response = &ModifyZoneSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyZoneSetting
// This API is used to modify the site configuration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWERROR = "InternalError.UnknowError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_CACHEKEYQUERYSTRINGTOOMANYVALUE = "InvalidParameter.CacheKeyQueryStringTooManyValue"
//  INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"
//  INVALIDPARAMETER_CLIENTIPCOUNTRYCONFLICTSWITHIPV6 = "InvalidParameter.ClientIpCountryConflictsWithIpv6"
//  INVALIDPARAMETER_GRPCREQUIREHTTP2 = "InvalidParameter.GrpcRequireHttp2"
//  INVALIDPARAMETER_INVALIDAWSPRIVATEACCESS = "InvalidParameter.InvalidAwsPrivateAccess"
//  INVALIDPARAMETER_INVALIDCACHECONFIGFOLLOWORIGIN = "InvalidParameter.InvalidCacheConfigFollowOrigin"
//  INVALIDPARAMETER_INVALIDCACHEKEYQUERYSTRINGVALUE = "InvalidParameter.InvalidCacheKeyQueryStringValue"
//  INVALIDPARAMETER_INVALIDCACHEONLYONSWITCH = "InvalidParameter.InvalidCacheOnlyOnSwitch"
//  INVALIDPARAMETER_INVALIDCACHETIME = "InvalidParameter.InvalidCacheTime"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINEBILLING = "InvalidParameter.InvalidDynamicRoutineBilling"
//  INVALIDPARAMETER_INVALIDHTTPS = "InvalidParameter.InvalidHttps"
//  INVALIDPARAMETER_INVALIDHTTPSCERTINFO = "InvalidParameter.InvalidHttpsCertInfo"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDHTTPSHSTSMAXAGE = "InvalidParameter.InvalidHttpsHstsMaxAge"
//  INVALIDPARAMETER_INVALIDHTTPSTLSVERSION = "InvalidParameter.InvalidHttpsTlsVersion"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPOSTMAXSIZEBILLING = "InvalidParameter.InvalidPostMaxSizeBilling"
//  INVALIDPARAMETER_INVALIDPOSTSIZEVALUE = "InvalidParameter.InvalidPostSizeValue"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDRESOURCEIDBILLING = "InvalidParameter.InvalidResourceIdBilling"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGCLIENTIP = "InvalidParameter.InvalidStandardDebugClientIp"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGEXPIRETIMELIMIT = "InvalidParameter.InvalidStandardDebugExpireTimeLimit"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_MULTIPLYLAYERNOTSUPPORTSMARTROUTING = "InvalidParameter.MultiplyLayerNotSupportSmartRouting"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_SETTINGINVALIDPARAM = "InvalidParameter.SettingInvalidParam"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDDISABLE = "OperationDenied.AccelerateMainlandDisable"
//  OPERATIONDENIED_ACCELERATEMAINLANDIPV6CONFLICT = "OperationDenied.AccelerateMainlandIpv6Conflict"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_POSTMAXSIZEQUOTANOTFOUND = "ResourceNotFound.PostMaxSizeQuotaNotFound"
//  RESOURCEUNAVAILABLE_CERTNOTFOUND = "ResourceUnavailable.CertNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) ModifyZoneSetting(request *ModifyZoneSettingRequest) (response *ModifyZoneSettingResponse, err error) {
    return c.ModifyZoneSettingWithContext(context.Background(), request)
}

// ModifyZoneSetting
// This API is used to modify the site configuration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWERROR = "InternalError.UnknowError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_CACHEKEYQUERYSTRINGTOOMANYVALUE = "InvalidParameter.CacheKeyQueryStringTooManyValue"
//  INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"
//  INVALIDPARAMETER_CLIENTIPCOUNTRYCONFLICTSWITHIPV6 = "InvalidParameter.ClientIpCountryConflictsWithIpv6"
//  INVALIDPARAMETER_GRPCREQUIREHTTP2 = "InvalidParameter.GrpcRequireHttp2"
//  INVALIDPARAMETER_INVALIDAWSPRIVATEACCESS = "InvalidParameter.InvalidAwsPrivateAccess"
//  INVALIDPARAMETER_INVALIDCACHECONFIGFOLLOWORIGIN = "InvalidParameter.InvalidCacheConfigFollowOrigin"
//  INVALIDPARAMETER_INVALIDCACHEKEYQUERYSTRINGVALUE = "InvalidParameter.InvalidCacheKeyQueryStringValue"
//  INVALIDPARAMETER_INVALIDCACHEONLYONSWITCH = "InvalidParameter.InvalidCacheOnlyOnSwitch"
//  INVALIDPARAMETER_INVALIDCACHETIME = "InvalidParameter.InvalidCacheTime"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINEBILLING = "InvalidParameter.InvalidDynamicRoutineBilling"
//  INVALIDPARAMETER_INVALIDHTTPS = "InvalidParameter.InvalidHttps"
//  INVALIDPARAMETER_INVALIDHTTPSCERTINFO = "InvalidParameter.InvalidHttpsCertInfo"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDHTTPSHSTSMAXAGE = "InvalidParameter.InvalidHttpsHstsMaxAge"
//  INVALIDPARAMETER_INVALIDHTTPSTLSVERSION = "InvalidParameter.InvalidHttpsTlsVersion"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPOSTMAXSIZEBILLING = "InvalidParameter.InvalidPostMaxSizeBilling"
//  INVALIDPARAMETER_INVALIDPOSTSIZEVALUE = "InvalidParameter.InvalidPostSizeValue"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDRESOURCEIDBILLING = "InvalidParameter.InvalidResourceIdBilling"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGCLIENTIP = "InvalidParameter.InvalidStandardDebugClientIp"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGEXPIRETIMELIMIT = "InvalidParameter.InvalidStandardDebugExpireTimeLimit"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_MULTIPLYLAYERNOTSUPPORTSMARTROUTING = "InvalidParameter.MultiplyLayerNotSupportSmartRouting"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_SETTINGINVALIDPARAM = "InvalidParameter.SettingInvalidParam"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDDISABLE = "OperationDenied.AccelerateMainlandDisable"
//  OPERATIONDENIED_ACCELERATEMAINLANDIPV6CONFLICT = "OperationDenied.AccelerateMainlandIpv6Conflict"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_POSTMAXSIZEQUOTANOTFOUND = "ResourceNotFound.PostMaxSizeQuotaNotFound"
//  RESOURCEUNAVAILABLE_CERTNOTFOUND = "ResourceUnavailable.CertNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) ModifyZoneSettingWithContext(ctx context.Context, request *ModifyZoneSettingRequest) (response *ModifyZoneSettingResponse, err error) {
    if request == nil {
        request = NewModifyZoneSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyZoneSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyZoneSettingResponse()
    err = c.Send(request, response)
    return
}

func NewModifyZoneStatusRequest() (request *ModifyZoneStatusRequest) {
    request = &ModifyZoneStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyZoneStatus")
    
    
    return
}

func NewModifyZoneStatusResponse() (response *ModifyZoneStatusResponse) {
    response = &ModifyZoneStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyZoneStatus
// This API is used to change the site status.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_L4PROXYINPROGRESSSTATUS = "OperationDenied.L4ProxyInProgressStatus"
//  OPERATIONDENIED_L4PROXYINSTOPPINGSTATUS = "OperationDenied.L4ProxyInStoppingStatus"
//  OPERATIONDENIED_L7HOSTINPROCESSSTATUS = "OperationDenied.L7HostInProcessStatus"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) ModifyZoneStatus(request *ModifyZoneStatusRequest) (response *ModifyZoneStatusResponse, err error) {
    return c.ModifyZoneStatusWithContext(context.Background(), request)
}

// ModifyZoneStatus
// This API is used to change the site status.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_L4PROXYINPROGRESSSTATUS = "OperationDenied.L4ProxyInProgressStatus"
//  OPERATIONDENIED_L4PROXYINSTOPPINGSTATUS = "OperationDenied.L4ProxyInStoppingStatus"
//  OPERATIONDENIED_L7HOSTINPROCESSSTATUS = "OperationDenied.L7HostInProcessStatus"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) ModifyZoneStatusWithContext(ctx context.Context, request *ModifyZoneStatusRequest) (response *ModifyZoneStatusResponse, err error) {
    if request == nil {
        request = NewModifyZoneStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyZoneStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyZoneStatusResponse()
    err = c.Send(request, response)
    return
}
