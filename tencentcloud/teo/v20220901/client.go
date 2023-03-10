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

func NewCheckCertificateRequest() (request *CheckCertificateRequest) {
    request = &CheckCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CheckCertificate")
    
    
    return
}

func NewCheckCertificateResponse() (response *CheckCertificateResponse) {
    response = &CheckCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckCertificate
// This API is used to verify a certificate.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_CERTCHAINERROR = "InvalidParameter.CertChainError"
//  INVALIDPARAMETER_CERTCHECKERROR = "InvalidParameter.CertCheckError"
//  INVALIDPARAMETER_CERTCOMPLETEERROR = "InvalidParameter.CertCompleteError"
//  INVALIDPARAMETER_CERTFORMATERROR = "InvalidParameter.CertFormatError"
//  INVALIDPARAMETER_CERTISEXPIRED = "InvalidParameter.CertIsExpired"
//  INVALIDPARAMETER_CERTNOCN = "InvalidParameter.CertNoCn"
//  INVALIDPARAMETER_CERTNOINFO = "InvalidParameter.CertNoInfo"
//  INVALIDPARAMETER_CERTNOTMATCHDOMAIN = "InvalidParameter.CertNotMatchDomain"
//  INVALIDPARAMETER_CERTNOTMATCHKEY = "InvalidParameter.CertNotMatchKey"
//  INVALIDPARAMETER_CERTNOTPEM = "InvalidParameter.CertNotPem"
//  INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"
//  INVALIDPARAMETER_CERTTOEXPIRE = "InvalidParameter.CertToExpire"
//  INVALIDPARAMETER_CERTTOOSHORTKEYSIZE = "InvalidParameter.CertTooShortKeySize"
//  INVALIDPARAMETER_CERTUNSUPPORTEDTYPE = "InvalidParameter.CertUnsupportedType"
//  INVALIDPARAMETER_INVALIDCERTINFO = "InvalidParameter.InvalidCertInfo"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CheckCertificate(request *CheckCertificateRequest) (response *CheckCertificateResponse, err error) {
    return c.CheckCertificateWithContext(context.Background(), request)
}

// CheckCertificate
// This API is used to verify a certificate.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_CERTCHAINERROR = "InvalidParameter.CertChainError"
//  INVALIDPARAMETER_CERTCHECKERROR = "InvalidParameter.CertCheckError"
//  INVALIDPARAMETER_CERTCOMPLETEERROR = "InvalidParameter.CertCompleteError"
//  INVALIDPARAMETER_CERTFORMATERROR = "InvalidParameter.CertFormatError"
//  INVALIDPARAMETER_CERTISEXPIRED = "InvalidParameter.CertIsExpired"
//  INVALIDPARAMETER_CERTNOCN = "InvalidParameter.CertNoCn"
//  INVALIDPARAMETER_CERTNOINFO = "InvalidParameter.CertNoInfo"
//  INVALIDPARAMETER_CERTNOTMATCHDOMAIN = "InvalidParameter.CertNotMatchDomain"
//  INVALIDPARAMETER_CERTNOTMATCHKEY = "InvalidParameter.CertNotMatchKey"
//  INVALIDPARAMETER_CERTNOTPEM = "InvalidParameter.CertNotPem"
//  INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"
//  INVALIDPARAMETER_CERTTOEXPIRE = "InvalidParameter.CertToExpire"
//  INVALIDPARAMETER_CERTTOOSHORTKEYSIZE = "InvalidParameter.CertTooShortKeySize"
//  INVALIDPARAMETER_CERTUNSUPPORTEDTYPE = "InvalidParameter.CertUnsupportedType"
//  INVALIDPARAMETER_INVALIDCERTINFO = "InvalidParameter.InvalidCertInfo"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CheckCertificateWithContext(ctx context.Context, request *CheckCertificateRequest) (response *CheckCertificateResponse, err error) {
    if request == nil {
        request = NewCheckCertificateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckCertificateResponse()
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
//  INVALIDPARAMETER_CERTNOTMATCHDOMAIN = "InvalidParameter.CertNotMatchDomain"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DOMAINISBLOCKED = "OperationDenied.DomainIsBlocked"
//  OPERATIONDENIED_DOMAINNOICP = "OperationDenied.DomainNoICP"
//  RESOURCEINUSE_DUPLICATENAME = "ResourceInUse.DuplicateName"
//  RESOURCEINUSE_ZONE = "ResourceInUse.Zone"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAliasDomain(request *CreateAliasDomainRequest) (response *CreateAliasDomainResponse, err error) {
    return c.CreateAliasDomainWithContext(context.Background(), request)
}

// CreateAliasDomain
// This API is used to create an alias domain name.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CERTNOTMATCHDOMAIN = "InvalidParameter.CertNotMatchDomain"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DOMAINISBLOCKED = "OperationDenied.DomainIsBlocked"
//  OPERATIONDENIED_DOMAINNOICP = "OperationDenied.DomainNoICP"
//  RESOURCEINUSE_DUPLICATENAME = "ResourceInUse.DuplicateName"
//  RESOURCEINUSE_ZONE = "ResourceInUse.Zone"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
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
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
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
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
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

func NewCreateCredentialRequest() (request *CreateCredentialRequest) {
    request = &CreateCredentialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateCredential")
    
    
    return
}

func NewCreateCredentialResponse() (response *CreateCredentialResponse) {
    response = &CreateCredentialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCredential
// Creates a credential for COS origin-pull.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateCredential(request *CreateCredentialRequest) (response *CreateCredentialResponse, err error) {
    return c.CreateCredentialWithContext(context.Background(), request)
}

// CreateCredential
// Creates a credential for COS origin-pull.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateCredentialWithContext(ctx context.Context, request *CreateCredentialRequest) (response *CreateCredentialResponse, err error) {
    if request == nil {
        request = NewCreateCredentialRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCredential require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCredentialResponse()
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
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  OPERATIONDENIED_DOMAINNOICP = "OperationDenied.DomainNoICP"
func (c *Client) CreatePlanForZone(request *CreatePlanForZoneRequest) (response *CreatePlanForZoneResponse, err error) {
    return c.CreatePlanForZoneWithContext(context.Background(), request)
}

// CreatePlanForZone
// This API is used to purchase a plan for a new site.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
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
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
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
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
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

func NewCreateReplayTaskRequest() (request *CreateReplayTaskRequest) {
    request = &CreateReplayTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateReplayTask")
    
    
    return
}

func NewCreateReplayTaskResponse() (response *CreateReplayTaskResponse) {
    response = &CreateReplayTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateReplayTask
// This API is used to create a replay task for purging or pre-warming URLs.
//
// error code that may be returned:
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  LIMITEXCEEDED_DAILYQUOTA = "LimitExceeded.DailyQuota"
func (c *Client) CreateReplayTask(request *CreateReplayTaskRequest) (response *CreateReplayTaskResponse, err error) {
    return c.CreateReplayTaskWithContext(context.Background(), request)
}

// CreateReplayTask
// This API is used to create a replay task for purging or pre-warming URLs.
//
// error code that may be returned:
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  LIMITEXCEEDED_DAILYQUOTA = "LimitExceeded.DailyQuota"
func (c *Client) CreateReplayTaskWithContext(ctx context.Context, request *CreateReplayTaskRequest) (response *CreateReplayTaskResponse, err error) {
    if request == nil {
        request = NewCreateReplayTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateReplayTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateReplayTaskResponse()
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
//  INVALIDPARAMETER_INVALIDERRORPAGEREDIRECTURL = "InvalidParameter.InvalidErrorPageRedirectUrl"
//  INVALIDPARAMETER_INVALIDHTTPSHSTSMAXAGE = "InvalidParameter.InvalidHttpsHstsMaxAge"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPOSTSIZEVALUE = "InvalidParameter.InvalidPostSizeValue"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAMEXFF = "InvalidParameter.InvalidRequestHeaderNameXff"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERNAME = "InvalidParameter.InvalidResponseHeaderName"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERVALUE = "InvalidParameter.InvalidResponseHeaderValue"
//  INVALIDPARAMETER_INVALIDRULEENGINEACTION = "InvalidParameter.InvalidRuleEngineAction"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGET = "InvalidParameter.InvalidRuleEngineTarget"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSEXTENSION = "InvalidParameter.InvalidRuleEngineTargetsExtension"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSURL = "InvalidParameter.InvalidRuleEngineTargetsUrl"
//  INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"
//  INVALIDPARAMETER_INVALIDURLREDIRECTHOST = "InvalidParameter.InvalidUrlRedirectHost"
//  INVALIDPARAMETER_INVALIDURLREDIRECTURL = "InvalidParameter.InvalidUrlRedirectUrl"
//  INVALIDPARAMETER_ORIGINORIGINGROUPIDISREQUIRED = "InvalidParameter.OriginOriginGroupIdIsRequired"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_TASKSYSTEMERROR = "InvalidParameter.TaskSystemError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) CreateRule(request *CreateRuleRequest) (response *CreateRuleResponse, err error) {
    return c.CreateRuleWithContext(context.Background(), request)
}

// CreateRule
// This API is used to create a rule in the rule engine.
//
// error code that may be returned:
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
//  INVALIDPARAMETER_INVALIDERRORPAGEREDIRECTURL = "InvalidParameter.InvalidErrorPageRedirectUrl"
//  INVALIDPARAMETER_INVALIDHTTPSHSTSMAXAGE = "InvalidParameter.InvalidHttpsHstsMaxAge"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPOSTSIZEVALUE = "InvalidParameter.InvalidPostSizeValue"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAMEXFF = "InvalidParameter.InvalidRequestHeaderNameXff"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERNAME = "InvalidParameter.InvalidResponseHeaderName"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERVALUE = "InvalidParameter.InvalidResponseHeaderValue"
//  INVALIDPARAMETER_INVALIDRULEENGINEACTION = "InvalidParameter.InvalidRuleEngineAction"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGET = "InvalidParameter.InvalidRuleEngineTarget"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSEXTENSION = "InvalidParameter.InvalidRuleEngineTargetsExtension"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSURL = "InvalidParameter.InvalidRuleEngineTargetsUrl"
//  INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"
//  INVALIDPARAMETER_INVALIDURLREDIRECTHOST = "InvalidParameter.InvalidUrlRedirectHost"
//  INVALIDPARAMETER_INVALIDURLREDIRECTURL = "InvalidParameter.InvalidUrlRedirectUrl"
//  INVALIDPARAMETER_ORIGINORIGINGROUPIDISREQUIRED = "InvalidParameter.OriginOriginGroupIdIsRequired"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_TASKSYSTEMERROR = "InvalidParameter.TaskSystemError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
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

func NewCreateSpeedTestingRequest() (request *CreateSpeedTestingRequest) {
    request = &CreateSpeedTestingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateSpeedTesting")
    
    
    return
}

func NewCreateSpeedTestingResponse() (response *CreateSpeedTestingResponse) {
    response = &CreateSpeedTestingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSpeedTesting
// This API is used to create a site testing task for a specified domain name.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDDOMAINSTATUS = "InvalidParameterValue.InvalidDomainStatus"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateSpeedTesting(request *CreateSpeedTestingRequest) (response *CreateSpeedTestingResponse, err error) {
    return c.CreateSpeedTestingWithContext(context.Background(), request)
}

// CreateSpeedTesting
// This API is used to create a site testing task for a specified domain name.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDDOMAINSTATUS = "InvalidParameterValue.InvalidDomainStatus"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateSpeedTestingWithContext(ctx context.Context, request *CreateSpeedTestingRequest) (response *CreateSpeedTestingResponse, err error) {
    if request == nil {
        request = NewCreateSpeedTestingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSpeedTesting require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSpeedTestingResponse()
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
func (c *Client) CreateZone(request *CreateZoneRequest) (response *CreateZoneResponse, err error) {
    return c.CreateZoneWithContext(context.Background(), request)
}

// CreateZone
// This API is used to access a new site.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
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
func (c *Client) DeleteApplicationProxy(request *DeleteApplicationProxyRequest) (response *DeleteApplicationProxyResponse, err error) {
    return c.DeleteApplicationProxyWithContext(context.Background(), request)
}

// DeleteApplicationProxy
// This API is used to delete an application proxy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteApplicationProxyRule(request *DeleteApplicationProxyRuleRequest) (response *DeleteApplicationProxyRuleResponse, err error) {
    return c.DeleteApplicationProxyRuleWithContext(context.Background(), request)
}

// DeleteApplicationProxyRule
// This API is used to delete an application proxy rule.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
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
func (c *Client) DeleteOriginGroup(request *DeleteOriginGroupRequest) (response *DeleteOriginGroupResponse, err error) {
    return c.DeleteOriginGroupWithContext(context.Background(), request)
}

// DeleteOriginGroup
// This API is used to delete an origin group.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
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
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_INVALIDRULEENGINENOTFOUND = "InvalidParameter.InvalidRuleEngineNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DeleteRules(request *DeleteRulesRequest) (response *DeleteRulesResponse, err error) {
    return c.DeleteRulesWithContext(context.Background(), request)
}

// DeleteRules
// This API is used to batch delete rules from the rule engine.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_INVALIDRULEENGINENOTFOUND = "InvalidParameter.InvalidRuleEngineNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
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
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteZone(request *DeleteZoneRequest) (response *DeleteZoneResponse, err error) {
    return c.DeleteZoneWithContext(context.Background(), request)
}

// DeleteZone
// This API is used to delete a site.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
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

func NewDescribeAddableEntityListRequest() (request *DescribeAddableEntityListRequest) {
    request = &DescribeAddableEntityListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeAddableEntityList")
    
    
    return
}

func NewDescribeAddableEntityListResponse() (response *DescribeAddableEntityListResponse) {
    response = &DescribeAddableEntityListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAddableEntityList
// This API is used to query available shipping entities.
//
// error code that may be returned:
//  RESOURCEUNAVAILABLE_PROXYZONENOTFOUND = "ResourceUnavailable.ProxyZoneNotFound"
func (c *Client) DescribeAddableEntityList(request *DescribeAddableEntityListRequest) (response *DescribeAddableEntityListResponse, err error) {
    return c.DescribeAddableEntityListWithContext(context.Background(), request)
}

// DescribeAddableEntityList
// This API is used to query available shipping entities.
//
// error code that may be returned:
//  RESOURCEUNAVAILABLE_PROXYZONENOTFOUND = "ResourceUnavailable.ProxyZoneNotFound"
func (c *Client) DescribeAddableEntityListWithContext(ctx context.Context, request *DescribeAddableEntityListRequest) (response *DescribeAddableEntityListResponse, err error) {
    if request == nil {
        request = NewDescribeAddableEntityListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAddableEntityList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAddableEntityListResponse()
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

func NewDescribeClientRuleListRequest() (request *DescribeClientRuleListRequest) {
    request = &DescribeClientRuleListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeClientRuleList")
    
    
    return
}

func NewDescribeClientRuleListResponse() (response *DescribeClientRuleListResponse) {
    response = &DescribeClientRuleListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClientRuleList
// This API is used to query the information of blocked clients.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeClientRuleList(request *DescribeClientRuleListRequest) (response *DescribeClientRuleListResponse, err error) {
    return c.DescribeClientRuleListWithContext(context.Background(), request)
}

// DescribeClientRuleList
// This API is used to query the information of blocked clients.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeClientRuleListWithContext(ctx context.Context, request *DescribeClientRuleListRequest) (response *DescribeClientRuleListResponse, err error) {
    if request == nil {
        request = NewDescribeClientRuleListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClientRuleList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClientRuleListResponse()
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

func NewDescribeDnsDataRequest() (request *DescribeDnsDataRequest) {
    request = &DescribeDnsDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeDnsData")
    
    
    return
}

func NewDescribeDnsDataResponse() (response *DescribeDnsDataResponse) {
    response = &DescribeDnsDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDnsData
// This API is used to get DNS requests.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeDnsData(request *DescribeDnsDataRequest) (response *DescribeDnsDataResponse, err error) {
    return c.DescribeDnsDataWithContext(context.Background(), request)
}

// DescribeDnsData
// This API is used to get DNS requests.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeDnsDataWithContext(ctx context.Context, request *DescribeDnsDataRequest) (response *DescribeDnsDataResponse, err error) {
    if request == nil {
        request = NewDescribeDnsDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDnsData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDnsDataResponse()
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
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeIdentifications(request *DescribeIdentificationsRequest) (response *DescribeIdentificationsResponse, err error) {
    return c.DescribeIdentificationsWithContext(context.Background(), request)
}

// DescribeIdentifications
// This API is used to query the verification information of a site.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
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

func NewDescribeLogSetsRequest() (request *DescribeLogSetsRequest) {
    request = &DescribeLogSetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeLogSets")
    
    
    return
}

func NewDescribeLogSetsResponse() (response *DescribeLogSetsResponse) {
    response = &DescribeLogSetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLogSets
// This API is used to get a list of logsets.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLogSets(request *DescribeLogSetsRequest) (response *DescribeLogSetsResponse, err error) {
    return c.DescribeLogSetsWithContext(context.Background(), request)
}

// DescribeLogSets
// This API is used to get a list of logsets.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLogSetsWithContext(ctx context.Context, request *DescribeLogSetsRequest) (response *DescribeLogSetsResponse, err error) {
    if request == nil {
        request = NewDescribeLogSetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogSets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogSetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogTopicTasksRequest() (request *DescribeLogTopicTasksRequest) {
    request = &DescribeLogTopicTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeLogTopicTasks")
    
    
    return
}

func NewDescribeLogTopicTasksResponse() (response *DescribeLogTopicTasksResponse) {
    response = &DescribeLogTopicTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLogTopicTasks
// This API is used to get a list of shipping tasks.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
func (c *Client) DescribeLogTopicTasks(request *DescribeLogTopicTasksRequest) (response *DescribeLogTopicTasksResponse, err error) {
    return c.DescribeLogTopicTasksWithContext(context.Background(), request)
}

// DescribeLogTopicTasks
// This API is used to get a list of shipping tasks.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
func (c *Client) DescribeLogTopicTasksWithContext(ctx context.Context, request *DescribeLogTopicTasksRequest) (response *DescribeLogTopicTasksResponse, err error) {
    if request == nil {
        request = NewDescribeLogTopicTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogTopicTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogTopicTasksResponse()
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
func (c *Client) DescribeRules(request *DescribeRulesRequest) (response *DescribeRulesResponse, err error) {
    return c.DescribeRulesWithContext(context.Background(), request)
}

// DescribeRules
// This API is used to query the rules in the rule engine.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
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
func (c *Client) DescribeRulesSetting(request *DescribeRulesSettingRequest) (response *DescribeRulesSettingResponse, err error) {
    return c.DescribeRulesSettingWithContext(context.Background(), request)
}

// DescribeRulesSetting
// This API is used to return the list of the settings of the rule engine that can be used for request match and their detailed recommended configuration information.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
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

func NewDescribeSingleL7AnalysisDataRequest() (request *DescribeSingleL7AnalysisDataRequest) {
    request = &DescribeSingleL7AnalysisDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeSingleL7AnalysisData")
    
    
    return
}

func NewDescribeSingleL7AnalysisDataResponse() (response *DescribeSingleL7AnalysisDataResponse) {
    response = &DescribeSingleL7AnalysisDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSingleL7AnalysisData
// This API is used to query the statistical calculation result of an L7 traffic metric. In common cases, only one value is returned for the query of statistical calculation results.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeSingleL7AnalysisData(request *DescribeSingleL7AnalysisDataRequest) (response *DescribeSingleL7AnalysisDataResponse, err error) {
    return c.DescribeSingleL7AnalysisDataWithContext(context.Background(), request)
}

// DescribeSingleL7AnalysisData
// This API is used to query the statistical calculation result of an L7 traffic metric. In common cases, only one value is returned for the query of statistical calculation results.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeSingleL7AnalysisDataWithContext(ctx context.Context, request *DescribeSingleL7AnalysisDataRequest) (response *DescribeSingleL7AnalysisDataResponse, err error) {
    if request == nil {
        request = NewDescribeSingleL7AnalysisDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSingleL7AnalysisData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSingleL7AnalysisDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSpeedTestingDetailsRequest() (request *DescribeSpeedTestingDetailsRequest) {
    request = &DescribeSpeedTestingDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeSpeedTestingDetails")
    
    
    return
}

func NewDescribeSpeedTestingDetailsResponse() (response *DescribeSpeedTestingDetailsResponse) {
    response = &DescribeSpeedTestingDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSpeedTestingDetails
// This API is used to query the site’s load speed across regions.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeSpeedTestingDetails(request *DescribeSpeedTestingDetailsRequest) (response *DescribeSpeedTestingDetailsResponse, err error) {
    return c.DescribeSpeedTestingDetailsWithContext(context.Background(), request)
}

// DescribeSpeedTestingDetails
// This API is used to query the site’s load speed across regions.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeSpeedTestingDetailsWithContext(ctx context.Context, request *DescribeSpeedTestingDetailsRequest) (response *DescribeSpeedTestingDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeSpeedTestingDetailsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSpeedTestingDetails require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSpeedTestingDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSpeedTestingMetricDataRequest() (request *DescribeSpeedTestingMetricDataRequest) {
    request = &DescribeSpeedTestingMetricDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeSpeedTestingMetricData")
    
    
    return
}

func NewDescribeSpeedTestingMetricDataResponse() (response *DescribeSpeedTestingMetricDataResponse) {
    response = &DescribeSpeedTestingMetricDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSpeedTestingMetricData
// This API is used to query the result of a site test.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSpeedTestingMetricData(request *DescribeSpeedTestingMetricDataRequest) (response *DescribeSpeedTestingMetricDataResponse, err error) {
    return c.DescribeSpeedTestingMetricDataWithContext(context.Background(), request)
}

// DescribeSpeedTestingMetricData
// This API is used to query the result of a site test.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSpeedTestingMetricDataWithContext(ctx context.Context, request *DescribeSpeedTestingMetricDataRequest) (response *DescribeSpeedTestingMetricDataResponse, err error) {
    if request == nil {
        request = NewDescribeSpeedTestingMetricDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSpeedTestingMetricData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSpeedTestingMetricDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSpeedTestingQuotaRequest() (request *DescribeSpeedTestingQuotaRequest) {
    request = &DescribeSpeedTestingQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeSpeedTestingQuota")
    
    
    return
}

func NewDescribeSpeedTestingQuotaResponse() (response *DescribeSpeedTestingQuotaResponse) {
    response = &DescribeSpeedTestingQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSpeedTestingQuota
// This API is used to query the quota limit on site tests.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSpeedTestingQuota(request *DescribeSpeedTestingQuotaRequest) (response *DescribeSpeedTestingQuotaResponse, err error) {
    return c.DescribeSpeedTestingQuotaWithContext(context.Background(), request)
}

// DescribeSpeedTestingQuota
// This API is used to query the quota limit on site tests.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSpeedTestingQuotaWithContext(ctx context.Context, request *DescribeSpeedTestingQuotaRequest) (response *DescribeSpeedTestingQuotaResponse, err error) {
    if request == nil {
        request = NewDescribeSpeedTestingQuotaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSpeedTestingQuota require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSpeedTestingQuotaResponse()
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

func NewDescribeWebManagedRulesDataRequest() (request *DescribeWebManagedRulesDataRequest) {
    request = &DescribeWebManagedRulesDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeWebManagedRulesData")
    
    
    return
}

func NewDescribeWebManagedRulesDataResponse() (response *DescribeWebManagedRulesDataResponse) {
    response = &DescribeWebManagedRulesDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWebManagedRulesData
// This API is used to query the WAF attack data recorded over time.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeWebManagedRulesData(request *DescribeWebManagedRulesDataRequest) (response *DescribeWebManagedRulesDataResponse, err error) {
    return c.DescribeWebManagedRulesDataWithContext(context.Background(), request)
}

// DescribeWebManagedRulesData
// This API is used to query the WAF attack data recorded over time.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeWebManagedRulesDataWithContext(ctx context.Context, request *DescribeWebManagedRulesDataRequest) (response *DescribeWebManagedRulesDataResponse, err error) {
    if request == nil {
        request = NewDescribeWebManagedRulesDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWebManagedRulesData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWebManagedRulesDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWebManagedRulesHitRuleDetailRequest() (request *DescribeWebManagedRulesHitRuleDetailRequest) {
    request = &DescribeWebManagedRulesHitRuleDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeWebManagedRulesHitRuleDetail")
    
    
    return
}

func NewDescribeWebManagedRulesHitRuleDetailResponse() (response *DescribeWebManagedRulesHitRuleDetailResponse) {
    response = &DescribeWebManagedRulesHitRuleDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWebManagedRulesHitRuleDetail
// This API is used to query the details of a hit WAF security rule.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeWebManagedRulesHitRuleDetail(request *DescribeWebManagedRulesHitRuleDetailRequest) (response *DescribeWebManagedRulesHitRuleDetailResponse, err error) {
    return c.DescribeWebManagedRulesHitRuleDetailWithContext(context.Background(), request)
}

// DescribeWebManagedRulesHitRuleDetail
// This API is used to query the details of a hit WAF security rule.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeWebManagedRulesHitRuleDetailWithContext(ctx context.Context, request *DescribeWebManagedRulesHitRuleDetailRequest) (response *DescribeWebManagedRulesHitRuleDetailResponse, err error) {
    if request == nil {
        request = NewDescribeWebManagedRulesHitRuleDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWebManagedRulesHitRuleDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWebManagedRulesHitRuleDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWebManagedRulesLogRequest() (request *DescribeWebManagedRulesLogRequest) {
    request = &DescribeWebManagedRulesLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeWebManagedRulesLog")
    
    
    return
}

func NewDescribeWebManagedRulesLogResponse() (response *DescribeWebManagedRulesLogResponse) {
    response = &DescribeWebManagedRulesLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWebManagedRulesLog
// This API is used to query web attack logs.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeWebManagedRulesLog(request *DescribeWebManagedRulesLogRequest) (response *DescribeWebManagedRulesLogResponse, err error) {
    return c.DescribeWebManagedRulesLogWithContext(context.Background(), request)
}

// DescribeWebManagedRulesLog
// This API is used to query web attack logs.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeWebManagedRulesLogWithContext(ctx context.Context, request *DescribeWebManagedRulesLogRequest) (response *DescribeWebManagedRulesLogResponse, err error) {
    if request == nil {
        request = NewDescribeWebManagedRulesLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWebManagedRulesLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWebManagedRulesLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWebProtectionClientIpListRequest() (request *DescribeWebProtectionClientIpListRequest) {
    request = &DescribeWebProtectionClientIpListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeWebProtectionClientIpList")
    
    
    return
}

func NewDescribeWebProtectionClientIpListResponse() (response *DescribeWebProtectionClientIpListResponse) {
    response = &DescribeWebProtectionClientIpListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWebProtectionClientIpList
// This API is used to query the information of CC attackers.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DescribeWebProtectionClientIpList(request *DescribeWebProtectionClientIpListRequest) (response *DescribeWebProtectionClientIpListResponse, err error) {
    return c.DescribeWebProtectionClientIpListWithContext(context.Background(), request)
}

// DescribeWebProtectionClientIpList
// This API is used to query the information of CC attackers.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DescribeWebProtectionClientIpListWithContext(ctx context.Context, request *DescribeWebProtectionClientIpListRequest) (response *DescribeWebProtectionClientIpListResponse, err error) {
    if request == nil {
        request = NewDescribeWebProtectionClientIpListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWebProtectionClientIpList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWebProtectionClientIpListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWebProtectionDataRequest() (request *DescribeWebProtectionDataRequest) {
    request = &DescribeWebProtectionDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeWebProtectionData")
    
    
    return
}

func NewDescribeWebProtectionDataResponse() (response *DescribeWebProtectionDataResponse) {
    response = &DescribeWebProtectionDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWebProtectionData
// This API is used to query the CC protection data recorded over time.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeWebProtectionData(request *DescribeWebProtectionDataRequest) (response *DescribeWebProtectionDataResponse, err error) {
    return c.DescribeWebProtectionDataWithContext(context.Background(), request)
}

// DescribeWebProtectionData
// This API is used to query the CC protection data recorded over time.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeWebProtectionDataWithContext(ctx context.Context, request *DescribeWebProtectionDataRequest) (response *DescribeWebProtectionDataResponse, err error) {
    if request == nil {
        request = NewDescribeWebProtectionDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWebProtectionData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWebProtectionDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWebProtectionHitRuleDetailRequest() (request *DescribeWebProtectionHitRuleDetailRequest) {
    request = &DescribeWebProtectionHitRuleDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeWebProtectionHitRuleDetail")
    
    
    return
}

func NewDescribeWebProtectionHitRuleDetailResponse() (response *DescribeWebProtectionHitRuleDetailResponse) {
    response = &DescribeWebProtectionHitRuleDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWebProtectionHitRuleDetail
// This API is used to query the details of a hit CC protection rule.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeWebProtectionHitRuleDetail(request *DescribeWebProtectionHitRuleDetailRequest) (response *DescribeWebProtectionHitRuleDetailResponse, err error) {
    return c.DescribeWebProtectionHitRuleDetailWithContext(context.Background(), request)
}

// DescribeWebProtectionHitRuleDetail
// This API is used to query the details of a hit CC protection rule.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeWebProtectionHitRuleDetailWithContext(ctx context.Context, request *DescribeWebProtectionHitRuleDetailRequest) (response *DescribeWebProtectionHitRuleDetailResponse, err error) {
    if request == nil {
        request = NewDescribeWebProtectionHitRuleDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWebProtectionHitRuleDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWebProtectionHitRuleDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWebProtectionTopDataRequest() (request *DescribeWebProtectionTopDataRequest) {
    request = &DescribeWebProtectionTopDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeWebProtectionTopData")
    
    
    return
}

func NewDescribeWebProtectionTopDataResponse() (response *DescribeWebProtectionTopDataResponse) {
    response = &DescribeWebProtectionTopDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWebProtectionTopData
// This API is used to query the top-ranked CC protection data.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeWebProtectionTopData(request *DescribeWebProtectionTopDataRequest) (response *DescribeWebProtectionTopDataResponse, err error) {
    return c.DescribeWebProtectionTopDataWithContext(context.Background(), request)
}

// DescribeWebProtectionTopData
// This API is used to query the top-ranked CC protection data.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeWebProtectionTopDataWithContext(ctx context.Context, request *DescribeWebProtectionTopDataRequest) (response *DescribeWebProtectionTopDataResponse, err error) {
    if request == nil {
        request = NewDescribeWebProtectionTopDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWebProtectionTopData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWebProtectionTopDataResponse()
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
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_L4PROXYINBANNEDSTATUS = "OperationDenied.L4ProxyInBannedStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyApplicationProxyStatus(request *ModifyApplicationProxyStatusRequest) (response *ModifyApplicationProxyStatusResponse, err error) {
    return c.ModifyApplicationProxyStatusWithContext(context.Background(), request)
}

// ModifyApplicationProxyStatus
// This API is used to modify the status of an application proxy.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
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

func NewModifyDefaultCertificateRequest() (request *ModifyDefaultCertificateRequest) {
    request = &ModifyDefaultCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyDefaultCertificate")
    
    
    return
}

func NewModifyDefaultCertificateResponse() (response *ModifyDefaultCertificateResponse) {
    response = &ModifyDefaultCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDefaultCertificate
// This example shows you how to modify the status of a default certificate.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDZONESTATUS = "FailedOperation.InvalidZoneStatus"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE_CERTNOTFOUND = "ResourceUnavailable.CertNotFound"
func (c *Client) ModifyDefaultCertificate(request *ModifyDefaultCertificateRequest) (response *ModifyDefaultCertificateResponse, err error) {
    return c.ModifyDefaultCertificateWithContext(context.Background(), request)
}

// ModifyDefaultCertificate
// This example shows you how to modify the status of a default certificate.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDZONESTATUS = "FailedOperation.InvalidZoneStatus"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE_CERTNOTFOUND = "ResourceUnavailable.CertNotFound"
func (c *Client) ModifyDefaultCertificateWithContext(ctx context.Context, request *ModifyDefaultCertificateRequest) (response *ModifyDefaultCertificateResponse, err error) {
    if request == nil {
        request = NewModifyDefaultCertificateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDefaultCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDefaultCertificateResponse()
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
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDZONESTATUS = "FailedOperation.InvalidZoneStatus"
//  INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_CERTNOTMATCHDOMAIN = "InvalidParameter.CertNotMatchDomain"
//  INVALIDPARAMETER_CERTTOEXPIRE = "InvalidParameter.CertToExpire"
//  INVALIDPARAMETER_CERTTOOSHORTKEYSIZE = "InvalidParameter.CertTooShortKeySize"
//  INVALIDPARAMETER_CNAMEWILDHOSTNOTALLOWAPPLYCERTIFICATE = "InvalidParameter.CnameWildHostNotAllowApplyCertificate"
//  INVALIDPARAMETER_HOSTSTATUSNOTALLOWAPPLYCERTIFICATE = "InvalidParameter.HostStatusNotAllowApplyCertificate"
//  INVALIDPARAMETER_INVALIDCERTINFO = "InvalidParameter.InvalidCertInfo"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE_HOSTNOTFOUND = "ResourceUnavailable.HostNotFound"
func (c *Client) ModifyHostsCertificate(request *ModifyHostsCertificateRequest) (response *ModifyHostsCertificateResponse, err error) {
    return c.ModifyHostsCertificateWithContext(context.Background(), request)
}

// ModifyHostsCertificate
// This API is used to modify the certificate of a domain name.
//
// error code that may be returned:
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDZONESTATUS = "FailedOperation.InvalidZoneStatus"
//  INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_CERTNOTMATCHDOMAIN = "InvalidParameter.CertNotMatchDomain"
//  INVALIDPARAMETER_CERTTOEXPIRE = "InvalidParameter.CertToExpire"
//  INVALIDPARAMETER_CERTTOOSHORTKEYSIZE = "InvalidParameter.CertTooShortKeySize"
//  INVALIDPARAMETER_CNAMEWILDHOSTNOTALLOWAPPLYCERTIFICATE = "InvalidParameter.CnameWildHostNotAllowApplyCertificate"
//  INVALIDPARAMETER_HOSTSTATUSNOTALLOWAPPLYCERTIFICATE = "InvalidParameter.HostStatusNotAllowApplyCertificate"
//  INVALIDPARAMETER_INVALIDCERTINFO = "InvalidParameter.InvalidCertInfo"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE_HOSTNOTFOUND = "ResourceUnavailable.HostNotFound"
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
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_L4STATUSNOTINONLINE = "OperationDenied.L4StatusNotInOnline"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyOriginGroup(request *ModifyOriginGroupRequest) (response *ModifyOriginGroupResponse, err error) {
    return c.ModifyOriginGroupWithContext(context.Background(), request)
}

// ModifyOriginGroup
// This API is used to modify an origin group.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_L4STATUSNOTINONLINE = "OperationDenied.L4StatusNotInOnline"
//  RESOURCENOTFOUND = "ResourceNotFound"
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
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ERRACTIONUNSUPPORTTARGET = "InvalidParameter.ErrActionUnsupportTarget"
//  INVALIDPARAMETER_ERRINVALIDACTION = "InvalidParameter.ErrInvalidAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMDUPLICATENAME = "InvalidParameter.ErrInvalidActionParamDuplicateName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMTOOMANYVALUES = "InvalidParameter.ErrInvalidActionParamTooManyValues"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONIGNORECASE = "InvalidParameter.ErrInvalidConditionIgnoreCase"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMETARGETNOTSUPPORTNAME = "InvalidParameter.ErrInvalidConditionNameTargetNotSupportName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADURL = "InvalidParameter.ErrInvalidConditionValueBadUrl"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUE = "InvalidParameter.ErrInvalidConditionValueBadValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUECONTAINFILENAMEEXTENSION = "InvalidParameter.ErrInvalidConditionValueBadValueContainFileNameExtension"
//  INVALIDPARAMETER_HOSTNOTFOUND = "InvalidParameter.HostNotFound"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESECRETKEY = "InvalidParameter.InvalidAuthenticationTypeSecretKey"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESIGNPARAM = "InvalidParameter.InvalidAuthenticationTypeSignParam"
//  INVALIDPARAMETER_INVALIDBACKUPSERVERNAME = "InvalidParameter.InvalidBackupServerName"
//  INVALIDPARAMETER_INVALIDCACHEKEY = "InvalidParameter.InvalidCacheKey"
//  INVALIDPARAMETER_INVALIDCACHETIME = "InvalidParameter.InvalidCacheTime"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDERRORPAGEREDIRECTURL = "InvalidParameter.InvalidErrorPageRedirectUrl"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAMEXFF = "InvalidParameter.InvalidRequestHeaderNameXff"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERNAME = "InvalidParameter.InvalidResponseHeaderName"
//  INVALIDPARAMETER_INVALIDRULEENGINEACTION = "InvalidParameter.InvalidRuleEngineAction"
//  INVALIDPARAMETER_INVALIDRULEENGINENOTFOUND = "InvalidParameter.InvalidRuleEngineNotFound"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGET = "InvalidParameter.InvalidRuleEngineTarget"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSEXTENSION = "InvalidParameter.InvalidRuleEngineTargetsExtension"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSURL = "InvalidParameter.InvalidRuleEngineTargetsUrl"
//  INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"
//  INVALIDPARAMETER_INVALIDURLREDIRECTURL = "InvalidParameter.InvalidUrlRedirectUrl"
//  INVALIDPARAMETER_KEYRULESINVALIDQUERYSTRINGVALUE = "InvalidParameter.KeyRulesInvalidQueryStringValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) ModifyRule(request *ModifyRuleRequest) (response *ModifyRuleResponse, err error) {
    return c.ModifyRuleWithContext(context.Background(), request)
}

// ModifyRule
// This API is used to modify a rule in the rule engine.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ERRACTIONUNSUPPORTTARGET = "InvalidParameter.ErrActionUnsupportTarget"
//  INVALIDPARAMETER_ERRINVALIDACTION = "InvalidParameter.ErrInvalidAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMDUPLICATENAME = "InvalidParameter.ErrInvalidActionParamDuplicateName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMTOOMANYVALUES = "InvalidParameter.ErrInvalidActionParamTooManyValues"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONIGNORECASE = "InvalidParameter.ErrInvalidConditionIgnoreCase"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMETARGETNOTSUPPORTNAME = "InvalidParameter.ErrInvalidConditionNameTargetNotSupportName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADURL = "InvalidParameter.ErrInvalidConditionValueBadUrl"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUE = "InvalidParameter.ErrInvalidConditionValueBadValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUECONTAINFILENAMEEXTENSION = "InvalidParameter.ErrInvalidConditionValueBadValueContainFileNameExtension"
//  INVALIDPARAMETER_HOSTNOTFOUND = "InvalidParameter.HostNotFound"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESECRETKEY = "InvalidParameter.InvalidAuthenticationTypeSecretKey"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESIGNPARAM = "InvalidParameter.InvalidAuthenticationTypeSignParam"
//  INVALIDPARAMETER_INVALIDBACKUPSERVERNAME = "InvalidParameter.InvalidBackupServerName"
//  INVALIDPARAMETER_INVALIDCACHEKEY = "InvalidParameter.InvalidCacheKey"
//  INVALIDPARAMETER_INVALIDCACHETIME = "InvalidParameter.InvalidCacheTime"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDERRORPAGEREDIRECTURL = "InvalidParameter.InvalidErrorPageRedirectUrl"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAMEXFF = "InvalidParameter.InvalidRequestHeaderNameXff"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERNAME = "InvalidParameter.InvalidResponseHeaderName"
//  INVALIDPARAMETER_INVALIDRULEENGINEACTION = "InvalidParameter.InvalidRuleEngineAction"
//  INVALIDPARAMETER_INVALIDRULEENGINENOTFOUND = "InvalidParameter.InvalidRuleEngineNotFound"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGET = "InvalidParameter.InvalidRuleEngineTarget"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSEXTENSION = "InvalidParameter.InvalidRuleEngineTargetsExtension"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSURL = "InvalidParameter.InvalidRuleEngineTargetsUrl"
//  INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"
//  INVALIDPARAMETER_INVALIDURLREDIRECTURL = "InvalidParameter.InvalidUrlRedirectUrl"
//  INVALIDPARAMETER_KEYRULESINVALIDQUERYSTRINGVALUE = "InvalidParameter.KeyRulesInvalidQueryStringValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
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

func NewModifyRulePriorityRequest() (request *ModifyRulePriorityRequest) {
    request = &ModifyRulePriorityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyRulePriority")
    
    
    return
}

func NewModifyRulePriorityResponse() (response *ModifyRulePriorityResponse) {
    response = &ModifyRulePriorityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyRulePriority
// This API is used to modify the priority of a rule in the rule engine.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDRULEENGINE = "InvalidParameter.InvalidRuleEngine"
//  INVALIDPARAMETER_INVALIDRULEENGINENOTFOUND = "InvalidParameter.InvalidRuleEngineNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) ModifyRulePriority(request *ModifyRulePriorityRequest) (response *ModifyRulePriorityResponse, err error) {
    return c.ModifyRulePriorityWithContext(context.Background(), request)
}

// ModifyRulePriority
// This API is used to modify the priority of a rule in the rule engine.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDRULEENGINE = "InvalidParameter.InvalidRuleEngine"
//  INVALIDPARAMETER_INVALIDRULEENGINENOTFOUND = "InvalidParameter.InvalidRuleEngineNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) ModifyRulePriorityWithContext(ctx context.Context, request *ModifyRulePriorityRequest) (response *ModifyRulePriorityResponse, err error) {
    if request == nil {
        request = NewModifyRulePriorityRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRulePriority require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRulePriorityResponse()
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
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) ModifySecurityPolicy(request *ModifySecurityPolicyRequest) (response *ModifySecurityPolicyResponse, err error) {
    return c.ModifySecurityPolicyWithContext(context.Background(), request)
}

// ModifySecurityPolicy
// This API is used to modify the web and bot security configurations.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
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

func NewModifySecurityWafGroupPolicyRequest() (request *ModifySecurityWafGroupPolicyRequest) {
    request = &ModifySecurityWafGroupPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifySecurityWafGroupPolicy")
    
    
    return
}

func NewModifySecurityWafGroupPolicyResponse() (response *ModifySecurityWafGroupPolicyResponse) {
    response = &ModifySecurityWafGroupPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySecurityWafGroupPolicy
// This API is used to modify a managed rule.
//
// error code that may be returned:
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
func (c *Client) ModifySecurityWafGroupPolicy(request *ModifySecurityWafGroupPolicyRequest) (response *ModifySecurityWafGroupPolicyResponse, err error) {
    return c.ModifySecurityWafGroupPolicyWithContext(context.Background(), request)
}

// ModifySecurityWafGroupPolicy
// This API is used to modify a managed rule.
//
// error code that may be returned:
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
func (c *Client) ModifySecurityWafGroupPolicyWithContext(ctx context.Context, request *ModifySecurityWafGroupPolicyRequest) (response *ModifySecurityWafGroupPolicyResponse, err error) {
    if request == nil {
        request = NewModifySecurityWafGroupPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySecurityWafGroupPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySecurityWafGroupPolicyResponse()
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
//  INVALIDPARAMETER_LENGTHEXCEEDSLIMIT = "InvalidParameter.LengthExceedsLimit"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ZONESAMEASNAME = "InvalidParameterValue.ZoneSameAsName"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_MULTIPLECNAMEZONE = "OperationDenied.MultipleCnameZone"
//  OPERATIONDENIED_NSNOTALLOWTRAFFICSTRATEGY = "OperationDenied.NSNotAllowTrafficStrategy"
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
//  INVALIDPARAMETER_LENGTHEXCEEDSLIMIT = "InvalidParameter.LengthExceedsLimit"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ZONESAMEASNAME = "InvalidParameterValue.ZoneSameAsName"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_MULTIPLECNAMEZONE = "OperationDenied.MultipleCnameZone"
//  OPERATIONDENIED_NSNOTALLOWTRAFFICSTRATEGY = "OperationDenied.NSNotAllowTrafficStrategy"
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
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_CACHEKEYQUERYSTRINGTOOMANYVALUE = "InvalidParameter.CacheKeyQueryStringTooManyValue"
//  INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"
//  INVALIDPARAMETER_GRPCREQUIREHTTP2 = "InvalidParameter.GrpcRequireHttp2"
//  INVALIDPARAMETER_INVALIDCACHECONFIGFOLLOWORIGIN = "InvalidParameter.InvalidCacheConfigFollowOrigin"
//  INVALIDPARAMETER_INVALIDCACHEKEYQUERYSTRINGVALUE = "InvalidParameter.InvalidCacheKeyQueryStringValue"
//  INVALIDPARAMETER_INVALIDCACHEONLYONSWITCH = "InvalidParameter.InvalidCacheOnlyOnSwitch"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINEBILLING = "InvalidParameter.InvalidDynamicRoutineBilling"
//  INVALIDPARAMETER_INVALIDHTTPSHSTSMAXAGE = "InvalidParameter.InvalidHttpsHstsMaxAge"
//  INVALIDPARAMETER_INVALIDHTTPSTLSVERSION = "InvalidParameter.InvalidHttpsTlsVersion"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPOSTMAXSIZEBILLING = "InvalidParameter.InvalidPostMaxSizeBilling"
//  INVALIDPARAMETER_INVALIDPOSTSIZEVALUE = "InvalidParameter.InvalidPostSizeValue"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDRESOURCEIDBILLING = "InvalidParameter.InvalidResourceIdBilling"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_SETTINGINVALIDPARAM = "InvalidParameter.SettingInvalidParam"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) ModifyZoneSetting(request *ModifyZoneSettingRequest) (response *ModifyZoneSettingResponse, err error) {
    return c.ModifyZoneSettingWithContext(context.Background(), request)
}

// ModifyZoneSetting
// This API is used to modify the site configuration.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_CACHEKEYQUERYSTRINGTOOMANYVALUE = "InvalidParameter.CacheKeyQueryStringTooManyValue"
//  INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"
//  INVALIDPARAMETER_GRPCREQUIREHTTP2 = "InvalidParameter.GrpcRequireHttp2"
//  INVALIDPARAMETER_INVALIDCACHECONFIGFOLLOWORIGIN = "InvalidParameter.InvalidCacheConfigFollowOrigin"
//  INVALIDPARAMETER_INVALIDCACHEKEYQUERYSTRINGVALUE = "InvalidParameter.InvalidCacheKeyQueryStringValue"
//  INVALIDPARAMETER_INVALIDCACHEONLYONSWITCH = "InvalidParameter.InvalidCacheOnlyOnSwitch"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINEBILLING = "InvalidParameter.InvalidDynamicRoutineBilling"
//  INVALIDPARAMETER_INVALIDHTTPSHSTSMAXAGE = "InvalidParameter.InvalidHttpsHstsMaxAge"
//  INVALIDPARAMETER_INVALIDHTTPSTLSVERSION = "InvalidParameter.InvalidHttpsTlsVersion"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPOSTMAXSIZEBILLING = "InvalidParameter.InvalidPostMaxSizeBilling"
//  INVALIDPARAMETER_INVALIDPOSTSIZEVALUE = "InvalidParameter.InvalidPostSizeValue"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDRESOURCEIDBILLING = "InvalidParameter.InvalidResourceIdBilling"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_SETTINGINVALIDPARAM = "InvalidParameter.SettingInvalidParam"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
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

func NewReclaimAliasDomainRequest() (request *ReclaimAliasDomainRequest) {
    request = &ReclaimAliasDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ReclaimAliasDomain")
    
    
    return
}

func NewReclaimAliasDomainResponse() (response *ReclaimAliasDomainResponse) {
    response = &ReclaimAliasDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReclaimAliasDomain
// To reclaim a site and its associated alias domain names, specify the "ZoneId" field; to reclaim an alias domain name that is already added, specify the "ZoneName" filed.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ReclaimAliasDomain(request *ReclaimAliasDomainRequest) (response *ReclaimAliasDomainResponse, err error) {
    return c.ReclaimAliasDomainWithContext(context.Background(), request)
}

// ReclaimAliasDomain
// To reclaim a site and its associated alias domain names, specify the "ZoneId" field; to reclaim an alias domain name that is already added, specify the "ZoneName" filed.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ReclaimAliasDomainWithContext(ctx context.Context, request *ReclaimAliasDomainRequest) (response *ReclaimAliasDomainResponse, err error) {
    if request == nil {
        request = NewReclaimAliasDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReclaimAliasDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewReclaimAliasDomainResponse()
    err = c.Send(request, response)
    return
}

func NewReclaimZoneRequest() (request *ReclaimZoneRequest) {
    request = &ReclaimZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ReclaimZone")
    
    
    return
}

func NewReclaimZoneResponse() (response *ReclaimZoneResponse) {
    response = &ReclaimZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReclaimZone
// This API is used to reclaim a site from other users after its ownership is verified.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ReclaimZone(request *ReclaimZoneRequest) (response *ReclaimZoneResponse, err error) {
    return c.ReclaimZoneWithContext(context.Background(), request)
}

// ReclaimZone
// This API is used to reclaim a site from other users after its ownership is verified.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ReclaimZoneWithContext(ctx context.Context, request *ReclaimZoneRequest) (response *ReclaimZoneResponse, err error) {
    if request == nil {
        request = NewReclaimZoneRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReclaimZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewReclaimZoneResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateOriginProtectionIPWhitelistRequest() (request *UpdateOriginProtectionIPWhitelistRequest) {
    request = &UpdateOriginProtectionIPWhitelistRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "UpdateOriginProtectionIPWhitelist")
    
    
    return
}

func NewUpdateOriginProtectionIPWhitelistResponse() (response *UpdateOriginProtectionIPWhitelistResponse) {
    response = &UpdateOriginProtectionIPWhitelistResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateOriginProtectionIPWhitelist
// This API is used to update the list of intermediate IPs.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_LATESTVERSIONNOW = "OperationDenied.LatestVersionNow"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) UpdateOriginProtectionIPWhitelist(request *UpdateOriginProtectionIPWhitelistRequest) (response *UpdateOriginProtectionIPWhitelistResponse, err error) {
    return c.UpdateOriginProtectionIPWhitelistWithContext(context.Background(), request)
}

// UpdateOriginProtectionIPWhitelist
// This API is used to update the list of intermediate IPs.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_LATESTVERSIONNOW = "OperationDenied.LatestVersionNow"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) UpdateOriginProtectionIPWhitelistWithContext(ctx context.Context, request *UpdateOriginProtectionIPWhitelistRequest) (response *UpdateOriginProtectionIPWhitelistResponse, err error) {
    if request == nil {
        request = NewUpdateOriginProtectionIPWhitelistRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateOriginProtectionIPWhitelist require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateOriginProtectionIPWhitelistResponse()
    err = c.Send(request, response)
    return
}
