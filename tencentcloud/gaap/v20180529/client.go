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

package v20180529

import (
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2018-05-29"

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


func NewAddRealServersRequest() (request *AddRealServersRequest) {
    request = &AddRealServersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "AddRealServers")
    return
}

func NewAddRealServersResponse() (response *AddRealServersResponse) {
    response = &AddRealServersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddRealServers
// This API is used to add the information of the origin server (server), which supports IP or the domain name.
//
// error code that may be returned:
//  AUTHFAILURE_NOTENTERPRISEAUTHORIZATION = "AuthFailure.NotEnterpriseAuthorization"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATERS = "InvalidParameterValue.DuplicateRS"
//  INVALIDPARAMETERVALUE_INVALIDTAGS = "InvalidParameterValue.InvalidTags"
//  INVALIDPARAMETERVALUE_PROJECTIDNOTBELONG = "InvalidParameterValue.ProjectIdNotBelong"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) AddRealServers(request *AddRealServersRequest) (response *AddRealServersResponse, err error) {
    if request == nil {
        request = NewAddRealServersRequest()
    }
    response = NewAddRealServersResponse()
    err = c.Send(request, response)
    return
}

func NewBindListenerRealServersRequest() (request *BindListenerRealServersRequest) {
    request = &BindListenerRealServersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "BindListenerRealServers")
    return
}

func NewBindListenerRealServersResponse() (response *BindListenerRealServersResponse) {
    response = &BindListenerRealServersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindListenerRealServers
// This API (BindListenerRealServers) is used for the TCP/UDP listener to bind/unbind the origin server.
//
// Note: This API unbinds the previously bound origin servers, and binds the origin servers selected for this call. For example, the previously bound origin servers are A, B and C, and the origin servers selected for this call are C, D and E, then the origin servers bound after this call will be C, D and E.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GROUPSTATUSNOTINRUNING = "FailedOperation.GroupStatusNotInRuning"
//  FAILEDOPERATION_INSTANCESTATUSNOTINRUNING = "FailedOperation.InstanceStatusNotInRuning"
//  FAILEDOPERATION_LIMITREALSERVERNUM = "FailedOperation.LimitRealServerNum"
//  FAILEDOPERATION_LISTENERHASTASK = "FailedOperation.ListenerHasTask"
//  FAILEDOPERATION_PROXYSTATUSNOTINRUNING = "FailedOperation.ProxyStatusNotInRuning"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_REALSERVERNOTBELONG = "InvalidParameterValue.RealServerNotBelong"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) BindListenerRealServers(request *BindListenerRealServersRequest) (response *BindListenerRealServersResponse, err error) {
    if request == nil {
        request = NewBindListenerRealServersRequest()
    }
    response = NewBindListenerRealServersResponse()
    err = c.Send(request, response)
    return
}

func NewBindRuleRealServersRequest() (request *BindRuleRealServersRequest) {
    request = &BindRuleRealServersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "BindRuleRealServers")
    return
}

func NewBindRuleRealServersResponse() (response *BindRuleRealServersResponse) {
    response = &BindRuleRealServersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindRuleRealServers
// This API is used to bind an origin server to the forwarding rules of layer-7 listeners. Note: This API unbinds all previously bound origin servers before binding those selected.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINSTATUSNOTINRUNNING = "FailedOperation.DomainStatusNotInRunning"
//  FAILEDOPERATION_INSTANCESTATUSNOTINRUNING = "FailedOperation.InstanceStatusNotInRuning"
//  FAILEDOPERATION_LISTENERHASTASK = "FailedOperation.ListenerHasTask"
//  FAILEDOPERATION_LISTENERSTATUSERROR = "FailedOperation.ListenerStatusError"
//  FAILEDOPERATION_NOTSUPPORTPROXYGROUP = "FailedOperation.NotSupportProxyGroup"
//  FAILEDOPERATION_PROXYSTATUSNOTINRUNING = "FailedOperation.ProxyStatusNotInRuning"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_REALSERVERNOTBELONG = "InvalidParameterValue.RealServerNotBelong"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) BindRuleRealServers(request *BindRuleRealServersRequest) (response *BindRuleRealServersResponse, err error) {
    if request == nil {
        request = NewBindRuleRealServersRequest()
    }
    response = NewBindRuleRealServersResponse()
    err = c.Send(request, response)
    return
}

func NewCheckProxyCreateRequest() (request *CheckProxyCreateRequest) {
    request = &CheckProxyCreateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "CheckProxyCreate")
    return
}

func NewCheckProxyCreateResponse() (response *CheckProxyCreateResponse) {
    response = &CheckProxyCreateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckProxyCreate
// This API (CheckProxyCreate) is used to query whether an acceleration connection with the specified configuration can be created.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTH = "InvalidParameterValue.InvalidBandwidth"
//  INVALIDPARAMETERVALUE_INVALIDCONCURRENCY = "InvalidParameterValue.InvalidConcurrency"
//  INVALIDPARAMETERVALUE_UNKNOWNACCESSREGION = "InvalidParameterValue.UnknownAccessRegion"
//  INVALIDPARAMETERVALUE_UNKNOWNDESTREGION = "InvalidParameterValue.UnknownDestRegion"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CheckProxyCreate(request *CheckProxyCreateRequest) (response *CheckProxyCreateResponse, err error) {
    if request == nil {
        request = NewCheckProxyCreateRequest()
    }
    response = NewCheckProxyCreateResponse()
    err = c.Send(request, response)
    return
}

func NewCloseProxiesRequest() (request *CloseProxiesRequest) {
    request = &CloseProxiesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "CloseProxies")
    return
}

func NewCloseProxiesResponse() (response *CloseProxiesResponse) {
    response = &CloseProxiesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CloseProxies
// This API (CloseProxies) is used to disable connections. If disabled, no traffic will be generated, but the basic configuration fee will still be incurred.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCOUNTBALANCEINSUFFICIENT = "FailedOperation.AccountBalanceInsufficient"
//  FAILEDOPERATION_DUPLICATEDREQUEST = "FailedOperation.DuplicatedRequest"
//  FAILEDOPERATION_NOTSUPPORTOLDVERSIONPROXY = "FailedOperation.NotSupportOldVersionProxy"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CloseProxies(request *CloseProxiesRequest) (response *CloseProxiesResponse, err error) {
    if request == nil {
        request = NewCloseProxiesRequest()
    }
    response = NewCloseProxiesResponse()
    err = c.Send(request, response)
    return
}

func NewCloseProxyGroupRequest() (request *CloseProxyGroupRequest) {
    request = &CloseProxyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "CloseProxyGroup")
    return
}

func NewCloseProxyGroupResponse() (response *CloseProxyGroupResponse) {
    response = &CloseProxyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CloseProxyGroup
// This API is used to disable a connection group. Once disabled, the connection group will no longer generate traffic, but the basic connection configuration fees will still be incurred every day.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCOUNTBALANCEINSUFFICIENT = "FailedOperation.AccountBalanceInsufficient"
//  FAILEDOPERATION_DUPLICATEDREQUEST = "FailedOperation.DuplicatedRequest"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CloseProxyGroup(request *CloseProxyGroupRequest) (response *CloseProxyGroupResponse, err error) {
    if request == nil {
        request = NewCloseProxyGroupRequest()
    }
    response = NewCloseProxyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCloseSecurityPolicyRequest() (request *CloseSecurityPolicyRequest) {
    request = &CloseSecurityPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "CloseSecurityPolicy")
    return
}

func NewCloseSecurityPolicyResponse() (response *CloseSecurityPolicyResponse) {
    response = &CloseSecurityPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CloseSecurityPolicy
// This API is used to disable security policies.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACTIONISDOING = "FailedOperation.ActionIsDoing"
//  FAILEDOPERATION_PROXYSECURITYALREADYCLOSE = "FailedOperation.ProxySecurityAlreadyClose"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CloseSecurityPolicy(request *CloseSecurityPolicyRequest) (response *CloseSecurityPolicyResponse, err error) {
    if request == nil {
        request = NewCloseSecurityPolicyRequest()
    }
    response = NewCloseSecurityPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCertificateRequest() (request *CreateCertificateRequest) {
    request = &CreateCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "CreateCertificate")
    return
}

func NewCreateCertificateResponse() (response *CreateCertificateResponse) {
    response = &CreateCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCertificate
// This API (CreateCertificate) is used to create the GAAP certificates and configuration files, including basic authentication configuration files, client CA certificates, server SSL certificates, GAAP SSL certificates, and origin server CA certificates.
//
// error code that may be returned:
//  AUTHFAILURE_NOTENTERPRISEAUTHORIZATION = "AuthFailure.NotEnterpriseAuthorization"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDCERTIFICATECONTENT = "InvalidParameterValue.InvalidCertificateContent"
//  INVALIDPARAMETERVALUE_INVALIDCERTIFICATEKEY = "InvalidParameterValue.InvalidCertificateKey"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCertificate(request *CreateCertificateRequest) (response *CreateCertificateResponse, err error) {
    if request == nil {
        request = NewCreateCertificateRequest()
    }
    response = NewCreateCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDomainRequest() (request *CreateDomainRequest) {
    request = &CreateDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "CreateDomain")
    return
}

func NewCreateDomainResponse() (response *CreateDomainResponse) {
    response = &CreateDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDomain
// This API (CreateDomain) is used to create the access domain name for the HTTP/HTTPS listener. Clients request the backend data by accessing this domain.
//
// This API only supports connections of version 3.0.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINALREADYEXISTED = "FailedOperation.DomainAlreadyExisted"
//  FAILEDOPERATION_INSTANCESTATUSNOTINRUNING = "FailedOperation.InstanceStatusNotInRuning"
//  FAILEDOPERATION_INVALIDLISTENERPROTOCOL = "FailedOperation.InvalidListenerProtocol"
//  FAILEDOPERATION_LIMITNUMOFDOMAINS = "FailedOperation.LimitNumofDomains"
//  FAILEDOPERATION_LISTENERHASTASK = "FailedOperation.ListenerHasTask"
//  FAILEDOPERATION_PROXYVERSIONNOTSUPPORT = "FailedOperation.ProxyVersionNotSupport"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateDomain(request *CreateDomainRequest) (response *CreateDomainResponse, err error) {
    if request == nil {
        request = NewCreateDomainRequest()
    }
    response = NewCreateDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDomainErrorPageInfoRequest() (request *CreateDomainErrorPageInfoRequest) {
    request = &CreateDomainErrorPageInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "CreateDomainErrorPageInfo")
    return
}

func NewCreateDomainErrorPageInfoResponse() (response *CreateDomainErrorPageInfoResponse) {
    response = &CreateDomainErrorPageInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDomainErrorPageInfo
// This API is used to customize the error code of an error response to the specified domain name.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDLISTENERPROTOCOL = "FailedOperation.InvalidListenerProtocol"
//  FAILEDOPERATION_LISTENERHASTASK = "FailedOperation.ListenerHasTask"
//  FAILEDOPERATION_PROXYSTATUSNOTINRUNING = "FailedOperation.ProxyStatusNotInRuning"
//  FAILEDOPERATION_PROXYVERSIONNOTSUPPORT = "FailedOperation.ProxyVersionNotSupport"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateDomainErrorPageInfo(request *CreateDomainErrorPageInfoRequest) (response *CreateDomainErrorPageInfoResponse, err error) {
    if request == nil {
        request = NewCreateDomainErrorPageInfoRequest()
    }
    response = NewCreateDomainErrorPageInfoResponse()
    err = c.Send(request, response)
    return
}

func NewCreateHTTPListenerRequest() (request *CreateHTTPListenerRequest) {
    request = &CreateHTTPListenerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "CreateHTTPListener")
    return
}

func NewCreateHTTPListenerResponse() (response *CreateHTTPListenerResponse) {
    response = &CreateHTTPListenerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateHTTPListener
// This API (CreateHTTPListener) is used to create an HTTP listener in the connection instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GROUPSTATUSNOTINRUNING = "FailedOperation.GroupStatusNotInRuning"
//  FAILEDOPERATION_INSTANCESTATUSNOTINRUNING = "FailedOperation.InstanceStatusNotInRuning"
//  FAILEDOPERATION_LIMITNUMOFLISTENER = "FailedOperation.LimitNumofListener"
//  FAILEDOPERATION_LISTENERHASTASK = "FailedOperation.ListenerHasTask"
//  FAILEDOPERATION_NOTSUPPORTPROXYGROUP = "FailedOperation.NotSupportProxyGroup"
//  FAILEDOPERATION_PROXYSTATUSNOTINRUNING = "FailedOperation.ProxyStatusNotInRuning"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATEDLISTENERPORT = "InvalidParameterValue.DuplicatedListenerPort"
//  INVALIDPARAMETERVALUE_INVALIDLISTENERPORT = "InvalidParameterValue.InvalidListenerPort"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateHTTPListener(request *CreateHTTPListenerRequest) (response *CreateHTTPListenerResponse, err error) {
    if request == nil {
        request = NewCreateHTTPListenerRequest()
    }
    response = NewCreateHTTPListenerResponse()
    err = c.Send(request, response)
    return
}

func NewCreateHTTPSListenerRequest() (request *CreateHTTPSListenerRequest) {
    request = &CreateHTTPSListenerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "CreateHTTPSListener")
    return
}

func NewCreateHTTPSListenerResponse() (response *CreateHTTPSListenerResponse) {
    response = &CreateHTTPSListenerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateHTTPSListener
// This API (CreateHTTPListener) is used to create an HTTPS listener in the connection instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GROUPSTATUSNOTINRUNING = "FailedOperation.GroupStatusNotInRuning"
//  FAILEDOPERATION_INSTANCESTATUSNOTINRUNING = "FailedOperation.InstanceStatusNotInRuning"
//  FAILEDOPERATION_LIMITNUMOFLISTENER = "FailedOperation.LimitNumofListener"
//  FAILEDOPERATION_LISTENERHASTASK = "FailedOperation.ListenerHasTask"
//  FAILEDOPERATION_PROXYSTATUSNOTINRUNING = "FailedOperation.ProxyStatusNotInRuning"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATEDLISTENERPORT = "InvalidParameterValue.DuplicatedListenerPort"
//  INVALIDPARAMETERVALUE_INVALIDCERTIFICATEID = "InvalidParameterValue.InvalidCertificateId"
//  INVALIDPARAMETERVALUE_INVALIDLISTENERPORT = "InvalidParameterValue.InvalidListenerPort"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateHTTPSListener(request *CreateHTTPSListenerRequest) (response *CreateHTTPSListenerResponse, err error) {
    if request == nil {
        request = NewCreateHTTPSListenerRequest()
    }
    response = NewCreateHTTPSListenerResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProxyRequest() (request *CreateProxyRequest) {
    request = &CreateProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "CreateProxy")
    return
}

func NewCreateProxyResponse() (response *CreateProxyResponse) {
    response = &CreateProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateProxy
// This API (CreateProxy) is used to create an acceleration connection with specified configuration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCOUNTBALANCEINSUFFICIENT = "FailedOperation.AccountBalanceInsufficient"
//  FAILEDOPERATION_DUPLICATEDREQUEST = "FailedOperation.DuplicatedRequest"
//  FAILEDOPERATION_LIMITNUMOFLISTENER = "FailedOperation.LimitNumofListener"
//  FAILEDOPERATION_PROXYSELLOUT = "FailedOperation.ProxySellOut"
//  FAILEDOPERATION_USERNOTAUTHENTICATED = "FailedOperation.UserNotAuthenticated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTH = "InvalidParameterValue.InvalidBandwidth"
//  INVALIDPARAMETERVALUE_INVALIDCONCURRENCY = "InvalidParameterValue.InvalidConcurrency"
//  INVALIDPARAMETERVALUE_INVALIDTAGS = "InvalidParameterValue.InvalidTags"
//  INVALIDPARAMETERVALUE_PROJECTIDNOTBELONG = "InvalidParameterValue.ProjectIdNotBelong"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateProxy(request *CreateProxyRequest) (response *CreateProxyResponse, err error) {
    if request == nil {
        request = NewCreateProxyRequest()
    }
    response = NewCreateProxyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProxyGroupRequest() (request *CreateProxyGroupRequest) {
    request = &CreateProxyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "CreateProxyGroup")
    return
}

func NewCreateProxyGroupResponse() (response *CreateProxyGroupResponse) {
    response = &CreateProxyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateProxyGroup
// This API (CreateProxyGroup) is used to create a connection group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LIMITNUMOFPROXIESINGROUP = "FailedOperation.LimitNumofProxiesInGroup"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDTAGS = "InvalidParameterValue.InvalidTags"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateProxyGroup(request *CreateProxyGroupRequest) (response *CreateProxyGroupResponse, err error) {
    if request == nil {
        request = NewCreateProxyGroupRequest()
    }
    response = NewCreateProxyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProxyGroupDomainRequest() (request *CreateProxyGroupDomainRequest) {
    request = &CreateProxyGroupDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "CreateProxyGroupDomain")
    return
}

func NewCreateProxyGroupDomainResponse() (response *CreateProxyGroupDomainResponse) {
    response = &CreateProxyGroupDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateProxyGroupDomain
// This API (CreateProxyGroupDomain) is used to create the connection group domain name, and enable the domain name resolution.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateProxyGroupDomain(request *CreateProxyGroupDomainRequest) (response *CreateProxyGroupDomainResponse, err error) {
    if request == nil {
        request = NewCreateProxyGroupDomainRequest()
    }
    response = NewCreateProxyGroupDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRuleRequest() (request *CreateRuleRequest) {
    request = &CreateRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "CreateRule")
    return
}

func NewCreateRuleResponse() (response *CreateRuleResponse) {
    response = &CreateRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRule
// This API (CreateRule) is used to create the forwarding rules of HTTP/HTTPS listeners.
//
// error code that may be returned:
//  AUTHFAILURE_NOTENTERPRISEAUTHORIZATION = "AuthFailure.NotEnterpriseAuthorization"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCESTATUSNOTINRUNING = "FailedOperation.InstanceStatusNotInRuning"
//  FAILEDOPERATION_LIMITNUMOFRULES = "FailedOperation.LimitNumofRules"
//  FAILEDOPERATION_LISTENERHASTASK = "FailedOperation.ListenerHasTask"
//  FAILEDOPERATION_NOTSUPPORTOLDVERSIONPROXY = "FailedOperation.NotSupportOldVersionProxy"
//  FAILEDOPERATION_RULEALREADYEXISTED = "FailedOperation.RuleAlreadyExisted"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateRule(request *CreateRuleRequest) (response *CreateRuleResponse, err error) {
    if request == nil {
        request = NewCreateRuleRequest()
    }
    response = NewCreateRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSecurityPolicyRequest() (request *CreateSecurityPolicyRequest) {
    request = &CreateSecurityPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "CreateSecurityPolicy")
    return
}

func NewCreateSecurityPolicyResponse() (response *CreateSecurityPolicyResponse) {
    response = &CreateSecurityPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSecurityPolicy
// This API is used to create security policies.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROXYSECURITYPOLICYEXISTED = "FailedOperation.ProxySecurityPolicyExisted"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateSecurityPolicy(request *CreateSecurityPolicyRequest) (response *CreateSecurityPolicyResponse, err error) {
    if request == nil {
        request = NewCreateSecurityPolicyRequest()
    }
    response = NewCreateSecurityPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSecurityRulesRequest() (request *CreateSecurityRulesRequest) {
    request = &CreateSecurityRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "CreateSecurityRules")
    return
}

func NewCreateSecurityRulesResponse() (response *CreateSecurityRulesResponse) {
    response = &CreateSecurityRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSecurityRules
// This API is used to add security policy rules.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROXYSECURITYPOLICYDEFAULTRULE = "FailedOperation.ProxySecurityPolicyDefaultRule"
//  FAILEDOPERATION_PROXYSECURITYPOLICYDUPLICATEDRULE = "FailedOperation.ProxySecurityPolicyDuplicatedRule"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateSecurityRules(request *CreateSecurityRulesRequest) (response *CreateSecurityRulesResponse, err error) {
    if request == nil {
        request = NewCreateSecurityRulesRequest()
    }
    response = NewCreateSecurityRulesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTCPListenersRequest() (request *CreateTCPListenersRequest) {
    request = &CreateTCPListenersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "CreateTCPListeners")
    return
}

func NewCreateTCPListenersResponse() (response *CreateTCPListenersResponse) {
    response = &CreateTCPListenersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTCPListeners
// This API (CreateTCPListeners) is used to batch create TCP listeners of single connections or connection groups.
//
// error code that may be returned:
//  AUTHFAILURE_NOTENTERPRISEAUTHORIZATION = "AuthFailure.NotEnterpriseAuthorization"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DUPLICATEDREQUEST = "FailedOperation.DuplicatedRequest"
//  FAILEDOPERATION_GROUPSTATUSNOTINRUNING = "FailedOperation.GroupStatusNotInRuning"
//  FAILEDOPERATION_INSTANCESTATUSNOTINRUNING = "FailedOperation.InstanceStatusNotInRuning"
//  FAILEDOPERATION_LIMITNUMOFLISTENER = "FailedOperation.LimitNumofListener"
//  FAILEDOPERATION_LISTENERHASTASK = "FailedOperation.ListenerHasTask"
//  FAILEDOPERATION_OPERATELIMITNUMOFLISTENER = "FailedOperation.OperateLimitNumofListener"
//  FAILEDOPERATION_PROXYSTATUSNOTINRUNING = "FailedOperation.ProxyStatusNotInRuning"
//  FAILEDOPERATION_RESOURCEUPGRADING = "FailedOperation.ResourceUpgrading"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATEDLISTENERPORT = "InvalidParameterValue.DuplicatedListenerPort"
//  INVALIDPARAMETERVALUE_INVALIDLISTENERPORT = "InvalidParameterValue.InvalidListenerPort"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateTCPListeners(request *CreateTCPListenersRequest) (response *CreateTCPListenersResponse, err error) {
    if request == nil {
        request = NewCreateTCPListenersRequest()
    }
    response = NewCreateTCPListenersResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUDPListenersRequest() (request *CreateUDPListenersRequest) {
    request = &CreateUDPListenersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "CreateUDPListeners")
    return
}

func NewCreateUDPListenersResponse() (response *CreateUDPListenersResponse) {
    response = &CreateUDPListenersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateUDPListeners
// This API (CreateTCPListeners) is used to batch create UDP listeners of single connections or connection groups.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GROUPSTATUSNOTINRUNING = "FailedOperation.GroupStatusNotInRuning"
//  FAILEDOPERATION_INSTANCESTATUSNOTINRUNING = "FailedOperation.InstanceStatusNotInRuning"
//  FAILEDOPERATION_LIMITNUMOFLISTENER = "FailedOperation.LimitNumofListener"
//  FAILEDOPERATION_LISTENERHASTASK = "FailedOperation.ListenerHasTask"
//  FAILEDOPERATION_OPERATELIMITNUMOFLISTENER = "FailedOperation.OperateLimitNumofListener"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATEDLISTENERPORT = "InvalidParameterValue.DuplicatedListenerPort"
//  INVALIDPARAMETERVALUE_INVALIDLISTENERPORT = "InvalidParameterValue.InvalidListenerPort"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateUDPListeners(request *CreateUDPListenersRequest) (response *CreateUDPListenersResponse, err error) {
    if request == nil {
        request = NewCreateUDPListenersRequest()
    }
    response = NewCreateUDPListenersResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCertificateRequest() (request *DeleteCertificateRequest) {
    request = &DeleteCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DeleteCertificate")
    return
}

func NewDeleteCertificateResponse() (response *DeleteCertificateResponse) {
    response = &DeleteCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCertificate
// This API (DeleteCertificate) is used to delete certificates.
//
// error code that may be returned:
//  FAILEDOPERATION_CERTIFICATEISUSING = "FailedOperation.CertificateIsUsing"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCertificate(request *DeleteCertificateRequest) (response *DeleteCertificateResponse, err error) {
    if request == nil {
        request = NewDeleteCertificateRequest()
    }
    response = NewDeleteCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDomainRequest() (request *DeleteDomainRequest) {
    request = &DeleteDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DeleteDomain")
    return
}

func NewDeleteDomainResponse() (response *DeleteDomainResponse) {
    response = &DeleteDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDomain
// This API (DeleteDomain) is only applicable to layer-7 listeners. It is used to delete the domain names of this listener, and all the rules under the domain name. All rules bound to the origin server are unbound automatically.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINSTATUSNOTINRUNNING = "FailedOperation.DomainStatusNotInRunning"
//  FAILEDOPERATION_GROUPSTATUSNOTINRUNING = "FailedOperation.GroupStatusNotInRuning"
//  FAILEDOPERATION_INSTANCESTATUSNOTINRUNING = "FailedOperation.InstanceStatusNotInRuning"
//  FAILEDOPERATION_LISTENERHASTASK = "FailedOperation.ListenerHasTask"
//  FAILEDOPERATION_PROXYSTATUSNOTINRUNING = "FailedOperation.ProxyStatusNotInRuning"
//  FAILEDOPERATION_REALSERVERALREADYBOUND = "FailedOperation.RealServerAlreadyBound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteDomain(request *DeleteDomainRequest) (response *DeleteDomainResponse, err error) {
    if request == nil {
        request = NewDeleteDomainRequest()
    }
    response = NewDeleteDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDomainErrorPageInfoRequest() (request *DeleteDomainErrorPageInfoRequest) {
    request = &DeleteDomainErrorPageInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DeleteDomainErrorPageInfo")
    return
}

func NewDeleteDomainErrorPageInfoResponse() (response *DeleteDomainErrorPageInfoResponse) {
    response = &DeleteDomainErrorPageInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDomainErrorPageInfo
// This API is used to delete a custom error code for a domain name.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LISTENERHASTASK = "FailedOperation.ListenerHasTask"
//  FAILEDOPERATION_LISTENERSTATUSERROR = "FailedOperation.ListenerStatusError"
//  FAILEDOPERATION_PROXYVERSIONNOTSUPPORT = "FailedOperation.ProxyVersionNotSupport"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteDomainErrorPageInfo(request *DeleteDomainErrorPageInfoRequest) (response *DeleteDomainErrorPageInfoResponse, err error) {
    if request == nil {
        request = NewDeleteDomainErrorPageInfoRequest()
    }
    response = NewDeleteDomainErrorPageInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteListenersRequest() (request *DeleteListenersRequest) {
    request = &DeleteListenersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DeleteListeners")
    return
}

func NewDeleteListenersResponse() (response *DeleteListenersResponse) {
    response = &DeleteListenersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteListeners
// This API (DeleteListeners) is used to batch delete the listeners of a connection or connection group, including layer-4/7 listeners.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GROUPSTATUSNOTINRUNING = "FailedOperation.GroupStatusNotInRuning"
//  FAILEDOPERATION_INSTANCESTATUSNOTINRUNING = "FailedOperation.InstanceStatusNotInRuning"
//  FAILEDOPERATION_LISTENERHASTASK = "FailedOperation.ListenerHasTask"
//  FAILEDOPERATION_REALSERVERALREADYBOUND = "FailedOperation.RealServerAlreadyBound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteListeners(request *DeleteListenersRequest) (response *DeleteListenersResponse, err error) {
    if request == nil {
        request = NewDeleteListenersRequest()
    }
    response = NewDeleteListenersResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteProxyGroupRequest() (request *DeleteProxyGroupRequest) {
    request = &DeleteProxyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DeleteProxyGroup")
    return
}

func NewDeleteProxyGroupResponse() (response *DeleteProxyGroupResponse) {
    response = &DeleteProxyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteProxyGroup
// This API (DeleteProxyGroup) is used to delete a connection group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETEPROXYGROUPPROXYREMAINED = "FailedOperation.DeleteProxyGroupProxyRemained"
//  FAILEDOPERATION_REALSERVERALREADYBOUND = "FailedOperation.RealServerAlreadyBound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteProxyGroup(request *DeleteProxyGroupRequest) (response *DeleteProxyGroupResponse, err error) {
    if request == nil {
        request = NewDeleteProxyGroupRequest()
    }
    response = NewDeleteProxyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRuleRequest() (request *DeleteRuleRequest) {
    request = &DeleteRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DeleteRule")
    return
}

func NewDeleteRuleResponse() (response *DeleteRuleResponse) {
    response = &DeleteRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRule
// This API (DeleteRule) is used to delete the forwarding rules of HTTP/HTTPS listeners.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINSTATUSNOTINRUNNING = "FailedOperation.DomainStatusNotInRunning"
//  FAILEDOPERATION_INSTANCESTATUSNOTINRUNING = "FailedOperation.InstanceStatusNotInRuning"
//  FAILEDOPERATION_LISTENERHASTASK = "FailedOperation.ListenerHasTask"
//  FAILEDOPERATION_PROXYSTATUSNOTINRUNING = "FailedOperation.ProxyStatusNotInRuning"
//  FAILEDOPERATION_REALSERVERALREADYBOUND = "FailedOperation.RealServerAlreadyBound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteRule(request *DeleteRuleRequest) (response *DeleteRuleResponse, err error) {
    if request == nil {
        request = NewDeleteRuleRequest()
    }
    response = NewDeleteRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSecurityPolicyRequest() (request *DeleteSecurityPolicyRequest) {
    request = &DeleteSecurityPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DeleteSecurityPolicy")
    return
}

func NewDeleteSecurityPolicyResponse() (response *DeleteSecurityPolicyResponse) {
    response = &DeleteSecurityPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteSecurityPolicy
// This API is used to delete security policies.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteSecurityPolicy(request *DeleteSecurityPolicyRequest) (response *DeleteSecurityPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteSecurityPolicyRequest()
    }
    response = NewDeleteSecurityPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSecurityRulesRequest() (request *DeleteSecurityRulesRequest) {
    request = &DeleteSecurityRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DeleteSecurityRules")
    return
}

func NewDeleteSecurityRulesResponse() (response *DeleteSecurityRulesResponse) {
    response = &DeleteSecurityRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteSecurityRules
// This API is used to delete security policy rules.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteSecurityRules(request *DeleteSecurityRulesRequest) (response *DeleteSecurityRulesResponse, err error) {
    if request == nil {
        request = NewDeleteSecurityRulesRequest()
    }
    response = NewDeleteSecurityRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessRegionsRequest() (request *DescribeAccessRegionsRequest) {
    request = &DescribeAccessRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeAccessRegions")
    return
}

func NewDescribeAccessRegionsResponse() (response *DescribeAccessRegionsResponse) {
    response = &DescribeAccessRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAccessRegions
// This API (DescribeAccessRegions) is used to query acceleration region (client access region).
//
// error code that may be returned:
//  AUTHFAILURE_NOTENTERPRISEAUTHORIZATION = "AuthFailure.NotEnterpriseAuthorization"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCEUPGRADING = "FailedOperation.ResourceUpgrading"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeAccessRegions(request *DescribeAccessRegionsRequest) (response *DescribeAccessRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeAccessRegionsRequest()
    }
    response = NewDescribeAccessRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessRegionsByDestRegionRequest() (request *DescribeAccessRegionsByDestRegionRequest) {
    request = &DescribeAccessRegionsByDestRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeAccessRegionsByDestRegion")
    return
}

func NewDescribeAccessRegionsByDestRegionResponse() (response *DescribeAccessRegionsByDestRegionResponse) {
    response = &DescribeAccessRegionsByDestRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAccessRegionsByDestRegion
// This API (DescribeAccessRegionsByDestRegion) is used to query the list of the available acceleration regions based on the origin server region.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_UNKNOWNDESTREGION = "InvalidParameterValue.UnknownDestRegion"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeAccessRegionsByDestRegion(request *DescribeAccessRegionsByDestRegionRequest) (response *DescribeAccessRegionsByDestRegionResponse, err error) {
    if request == nil {
        request = NewDescribeAccessRegionsByDestRegionRequest()
    }
    response = NewDescribeAccessRegionsByDestRegionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCertificateDetailRequest() (request *DescribeCertificateDetailRequest) {
    request = &DescribeCertificateDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeCertificateDetail")
    return
}

func NewDescribeCertificateDetailResponse() (response *DescribeCertificateDetailResponse) {
    response = &DescribeCertificateDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCertificateDetail
// This API (DescribeCertificateDetail) is used to query certificate details, including the certificate ID, name, type, content, key, and other information.
//
// error code that may be returned:
//  AUTHFAILURE_NOTENTERPRISEAUTHORIZATION = "AuthFailure.NotEnterpriseAuthorization"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCertificateDetail(request *DescribeCertificateDetailRequest) (response *DescribeCertificateDetailResponse, err error) {
    if request == nil {
        request = NewDescribeCertificateDetailRequest()
    }
    response = NewDescribeCertificateDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCertificatesRequest() (request *DescribeCertificatesRequest) {
    request = &DescribeCertificatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeCertificates")
    return
}

func NewDescribeCertificatesResponse() (response *DescribeCertificatesResponse) {
    response = &DescribeCertificatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCertificates
// This API (DescribeCertificates) is used to query the list of available certificates.
//
// error code that may be returned:
//  AUTHFAILURE_NOTENTERPRISEAUTHORIZATION = "AuthFailure.NotEnterpriseAuthorization"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCertificates(request *DescribeCertificatesRequest) (response *DescribeCertificatesResponse, err error) {
    if request == nil {
        request = NewDescribeCertificatesRequest()
    }
    response = NewDescribeCertificatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCountryAreaMappingRequest() (request *DescribeCountryAreaMappingRequest) {
    request = &DescribeCountryAreaMappingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeCountryAreaMapping")
    return
}

func NewDescribeCountryAreaMappingResponse() (response *DescribeCountryAreaMappingResponse) {
    response = &DescribeCountryAreaMappingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCountryAreaMapping
// This API (DescribeCountryAreaMapping) is used to obtain the country/region code mapping table.
//
// error code that may be returned:
//  AUTHFAILURE_NOTENTERPRISEAUTHORIZATION = "AuthFailure.NotEnterpriseAuthorization"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeCountryAreaMapping(request *DescribeCountryAreaMappingRequest) (response *DescribeCountryAreaMappingResponse, err error) {
    if request == nil {
        request = NewDescribeCountryAreaMappingRequest()
    }
    response = NewDescribeCountryAreaMappingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDestRegionsRequest() (request *DescribeDestRegionsRequest) {
    request = &DescribeDestRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeDestRegions")
    return
}

func NewDescribeDestRegionsResponse() (response *DescribeDestRegionsResponse) {
    response = &DescribeDestRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDestRegions
// This API (DescribeDestRegions) is used to query an origin server region (i.e., the region in which the origin server locates).
//
// error code that may be returned:
//  AUTHFAILURE_NOTENTERPRISEAUTHORIZATION = "AuthFailure.NotEnterpriseAuthorization"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDestRegions(request *DescribeDestRegionsRequest) (response *DescribeDestRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeDestRegionsRequest()
    }
    response = NewDescribeDestRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainErrorPageInfoRequest() (request *DescribeDomainErrorPageInfoRequest) {
    request = &DescribeDomainErrorPageInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeDomainErrorPageInfo")
    return
}

func NewDescribeDomainErrorPageInfoResponse() (response *DescribeDomainErrorPageInfoResponse) {
    response = &DescribeDomainErrorPageInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDomainErrorPageInfo
// This API is used to query the custom error response to a domain name.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTSUPPORTPROXYGROUP = "FailedOperation.NotSupportProxyGroup"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDomainErrorPageInfo(request *DescribeDomainErrorPageInfoRequest) (response *DescribeDomainErrorPageInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDomainErrorPageInfoRequest()
    }
    response = NewDescribeDomainErrorPageInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainErrorPageInfoByIdsRequest() (request *DescribeDomainErrorPageInfoByIdsRequest) {
    request = &DescribeDomainErrorPageInfoByIdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeDomainErrorPageInfoByIds")
    return
}

func NewDescribeDomainErrorPageInfoByIdsResponse() (response *DescribeDomainErrorPageInfoByIdsResponse) {
    response = &DescribeDomainErrorPageInfoByIdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDomainErrorPageInfoByIds
// This API is used to query the corresponding error response by a custom error ID.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTSUPPORTPROXYGROUP = "FailedOperation.NotSupportProxyGroup"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDomainErrorPageInfoByIds(request *DescribeDomainErrorPageInfoByIdsRequest) (response *DescribeDomainErrorPageInfoByIdsResponse, err error) {
    if request == nil {
        request = NewDescribeDomainErrorPageInfoByIdsRequest()
    }
    response = NewDescribeDomainErrorPageInfoByIdsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupAndStatisticsProxyRequest() (request *DescribeGroupAndStatisticsProxyRequest) {
    request = &DescribeGroupAndStatisticsProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeGroupAndStatisticsProxy")
    return
}

func NewDescribeGroupAndStatisticsProxyResponse() (response *DescribeGroupAndStatisticsProxyResponse) {
    response = &DescribeGroupAndStatisticsProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGroupAndStatisticsProxy
// This is an internal API. It is used to query the information of connections and connection groups from which the statistics can be derived.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCECANNOTACCESS = "FailedOperation.ResourceCanNotAccess"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeGroupAndStatisticsProxy(request *DescribeGroupAndStatisticsProxyRequest) (response *DescribeGroupAndStatisticsProxyResponse, err error) {
    if request == nil {
        request = NewDescribeGroupAndStatisticsProxyRequest()
    }
    response = NewDescribeGroupAndStatisticsProxyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupDomainConfigRequest() (request *DescribeGroupDomainConfigRequest) {
    request = &DescribeGroupDomainConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeGroupDomainConfig")
    return
}

func NewDescribeGroupDomainConfigResponse() (response *DescribeGroupDomainConfigResponse) {
    response = &DescribeGroupDomainConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGroupDomainConfig
// This API (DescribeGroupDomainConfig) is used to obtain the domain name resolution configuration details of a connection group.
//
// error code that may be returned:
//  AUTHFAILURE_NOTENTERPRISEAUTHORIZATION = "AuthFailure.NotEnterpriseAuthorization"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCEUPGRADING = "FailedOperation.ResourceUpgrading"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeGroupDomainConfig(request *DescribeGroupDomainConfigRequest) (response *DescribeGroupDomainConfigResponse, err error) {
    if request == nil {
        request = NewDescribeGroupDomainConfigRequest()
    }
    response = NewDescribeGroupDomainConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHTTPListenersRequest() (request *DescribeHTTPListenersRequest) {
    request = &DescribeHTTPListenersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeHTTPListeners")
    return
}

func NewDescribeHTTPListenersResponse() (response *DescribeHTTPListenersResponse) {
    response = &DescribeHTTPListenersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeHTTPListeners
// This API (DescribeHTTPListeners) is used to query HTTP listener information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTSUPPORTOLDVERSIONPROXY = "FailedOperation.NotSupportOldVersionProxy"
//  FAILEDOPERATION_NOTSUPPORTPROXYGROUP = "FailedOperation.NotSupportProxyGroup"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeHTTPListeners(request *DescribeHTTPListenersRequest) (response *DescribeHTTPListenersResponse, err error) {
    if request == nil {
        request = NewDescribeHTTPListenersRequest()
    }
    response = NewDescribeHTTPListenersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHTTPSListenersRequest() (request *DescribeHTTPSListenersRequest) {
    request = &DescribeHTTPSListenersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeHTTPSListeners")
    return
}

func NewDescribeHTTPSListenersResponse() (response *DescribeHTTPSListenersResponse) {
    response = &DescribeHTTPSListenersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeHTTPSListeners
// This API (DescribeHTTPSListeners) is used to query HTTPS listener information.
//
// error code that may be returned:
//  AUTHFAILURE_NOTENTERPRISEAUTHORIZATION = "AuthFailure.NotEnterpriseAuthorization"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTSUPPORTOLDVERSIONPROXY = "FailedOperation.NotSupportOldVersionProxy"
//  FAILEDOPERATION_NOTSUPPORTPROXYGROUP = "FailedOperation.NotSupportProxyGroup"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeHTTPSListeners(request *DescribeHTTPSListenersRequest) (response *DescribeHTTPSListenersResponse, err error) {
    if request == nil {
        request = NewDescribeHTTPSListenersRequest()
    }
    response = NewDescribeHTTPSListenersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListenerRealServersRequest() (request *DescribeListenerRealServersRequest) {
    request = &DescribeListenerRealServersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeListenerRealServers")
    return
}

func NewDescribeListenerRealServersResponse() (response *DescribeListenerRealServersResponse) {
    response = &DescribeListenerRealServersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListenerRealServers
// This API (DescribeListenerRealServers) is used to query the origin server list of TCP/UDP listeners, including the list of bound origin servers and origin servers that can be bound.
//
// error code that may be returned:
//  AUTHFAILURE_NOTENTERPRISEAUTHORIZATION = "AuthFailure.NotEnterpriseAuthorization"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeListenerRealServers(request *DescribeListenerRealServersRequest) (response *DescribeListenerRealServersResponse, err error) {
    if request == nil {
        request = NewDescribeListenerRealServersRequest()
    }
    response = NewDescribeListenerRealServersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListenerStatisticsRequest() (request *DescribeListenerStatisticsRequest) {
    request = &DescribeListenerStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeListenerStatistics")
    return
}

func NewDescribeListenerStatisticsResponse() (response *DescribeListenerStatisticsResponse) {
    response = &DescribeListenerStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListenerStatistics
// This API is used to query listener statistics, including inbound/outbound bandwidth, inbound/outbound packets, and concurrence data. It supports granularities of 300, 3,600, and 86,400. Default value is the highest within the granularity range.
//
// error code that may be returned:
//  AUTHFAILURE_NOTENTERPRISEAUTHORIZATION = "AuthFailure.NotEnterpriseAuthorization"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCECANNOTACCESS = "FailedOperation.ResourceCanNotAccess"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeListenerStatistics(request *DescribeListenerStatisticsRequest) (response *DescribeListenerStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeListenerStatisticsRequest()
    }
    response = NewDescribeListenerStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProxiesRequest() (request *DescribeProxiesRequest) {
    request = &DescribeProxiesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeProxies")
    return
}

func NewDescribeProxiesResponse() (response *DescribeProxiesResponse) {
    response = &DescribeProxiesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProxies
// This API (DescribeProxies) is used to query the connection instance list.
//
// error code that may be returned:
//  AUTHFAILURE_NOTENTERPRISEAUTHORIZATION = "AuthFailure.NotEnterpriseAuthorization"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeProxies(request *DescribeProxiesRequest) (response *DescribeProxiesResponse, err error) {
    if request == nil {
        request = NewDescribeProxiesRequest()
    }
    response = NewDescribeProxiesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProxiesStatusRequest() (request *DescribeProxiesStatusRequest) {
    request = &DescribeProxiesStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeProxiesStatus")
    return
}

func NewDescribeProxiesStatusResponse() (response *DescribeProxiesStatusResponse) {
    response = &DescribeProxiesStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProxiesStatus
// This API (DescribeProxiesStatus) is used to query the connection status list.
//
// error code that may be returned:
//  AUTHFAILURE_NOTENTERPRISEAUTHORIZATION = "AuthFailure.NotEnterpriseAuthorization"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCOUNTBALANCEINSUFFICIENT = "FailedOperation.AccountBalanceInsufficient"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeProxiesStatus(request *DescribeProxiesStatusRequest) (response *DescribeProxiesStatusResponse, err error) {
    if request == nil {
        request = NewDescribeProxiesStatusRequest()
    }
    response = NewDescribeProxiesStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProxyAndStatisticsListenersRequest() (request *DescribeProxyAndStatisticsListenersRequest) {
    request = &DescribeProxyAndStatisticsListenersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeProxyAndStatisticsListeners")
    return
}

func NewDescribeProxyAndStatisticsListenersResponse() (response *DescribeProxyAndStatisticsListenersResponse) {
    response = &DescribeProxyAndStatisticsListenersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProxyAndStatisticsListeners
// This is an internal API. It is used to query the information of connections and listeners from which the statistics can be derived.
//
// error code that may be returned:
//  AUTHFAILURE_NOTENTERPRISEAUTHORIZATION = "AuthFailure.NotEnterpriseAuthorization"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeProxyAndStatisticsListeners(request *DescribeProxyAndStatisticsListenersRequest) (response *DescribeProxyAndStatisticsListenersResponse, err error) {
    if request == nil {
        request = NewDescribeProxyAndStatisticsListenersRequest()
    }
    response = NewDescribeProxyAndStatisticsListenersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProxyDetailRequest() (request *DescribeProxyDetailRequest) {
    request = &DescribeProxyDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeProxyDetail")
    return
}

func NewDescribeProxyDetailResponse() (response *DescribeProxyDetailResponse) {
    response = &DescribeProxyDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProxyDetail
// This API (DescribeProxyDetail) is used to query connection details.
//
// error code that may be returned:
//  AUTHFAILURE_NOTENTERPRISEAUTHORIZATION = "AuthFailure.NotEnterpriseAuthorization"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCECANNOTACCESS = "FailedOperation.ResourceCanNotAccess"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PROJECTIDNOTBELONG = "InvalidParameterValue.ProjectIdNotBelong"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeProxyDetail(request *DescribeProxyDetailRequest) (response *DescribeProxyDetailResponse, err error) {
    if request == nil {
        request = NewDescribeProxyDetailRequest()
    }
    response = NewDescribeProxyDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProxyGroupDetailsRequest() (request *DescribeProxyGroupDetailsRequest) {
    request = &DescribeProxyGroupDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeProxyGroupDetails")
    return
}

func NewDescribeProxyGroupDetailsResponse() (response *DescribeProxyGroupDetailsResponse) {
    response = &DescribeProxyGroupDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProxyGroupDetails
// This API (DescribeProxyGroupDetails) is used to query connection group details.
//
// error code that may be returned:
//  AUTHFAILURE_NOTENTERPRISEAUTHORIZATION = "AuthFailure.NotEnterpriseAuthorization"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeProxyGroupDetails(request *DescribeProxyGroupDetailsRequest) (response *DescribeProxyGroupDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeProxyGroupDetailsRequest()
    }
    response = NewDescribeProxyGroupDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProxyGroupListRequest() (request *DescribeProxyGroupListRequest) {
    request = &DescribeProxyGroupListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeProxyGroupList")
    return
}

func NewDescribeProxyGroupListResponse() (response *DescribeProxyGroupListResponse) {
    response = &DescribeProxyGroupListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProxyGroupList
// This API (DescribeProxyGroupList) is used to pull the list of connection groups and the basic information of each connection group.
//
// error code that may be returned:
//  AUTHFAILURE_NOTENTERPRISEAUTHORIZATION = "AuthFailure.NotEnterpriseAuthorization"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeProxyGroupList(request *DescribeProxyGroupListRequest) (response *DescribeProxyGroupListResponse, err error) {
    if request == nil {
        request = NewDescribeProxyGroupListRequest()
    }
    response = NewDescribeProxyGroupListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProxyGroupStatisticsRequest() (request *DescribeProxyGroupStatisticsRequest) {
    request = &DescribeProxyGroupStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeProxyGroupStatistics")
    return
}

func NewDescribeProxyGroupStatisticsResponse() (response *DescribeProxyGroupStatisticsResponse) {
    response = &DescribeProxyGroupStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProxyGroupStatistics
// This API is used to query listener statistics, including inbound/outbound bandwidth, inbound/outbound packets, and concurrence data. It supports granularities of 300, 3,600, and 86,400. Default value is the highest within the granularity range.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCECANNOTACCESS = "FailedOperation.ResourceCanNotAccess"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeProxyGroupStatistics(request *DescribeProxyGroupStatisticsRequest) (response *DescribeProxyGroupStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeProxyGroupStatisticsRequest()
    }
    response = NewDescribeProxyGroupStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProxyStatisticsRequest() (request *DescribeProxyStatisticsRequest) {
    request = &DescribeProxyStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeProxyStatistics")
    return
}

func NewDescribeProxyStatisticsResponse() (response *DescribeProxyStatisticsResponse) {
    response = &DescribeProxyStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProxyStatistics
// This API is used to query listener statistics, including inbound/outbound bandwidth, inbound/outbound packets, concurrence, packet loss, and latency data. It supports granularities of 300, 3,600, and 86,400. Default value is the highest within the granularity range.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCECANNOTACCESS = "FailedOperation.ResourceCanNotAccess"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeProxyStatistics(request *DescribeProxyStatisticsRequest) (response *DescribeProxyStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeProxyStatisticsRequest()
    }
    response = NewDescribeProxyStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRealServerStatisticsRequest() (request *DescribeRealServerStatisticsRequest) {
    request = &DescribeRealServerStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeRealServerStatistics")
    return
}

func NewDescribeRealServerStatisticsResponse() (response *DescribeRealServerStatisticsResponse) {
    response = &DescribeRealServerStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRealServerStatistics
// This API is used to query the statistics of an origin server's health check results. Origin server status is displayed as 1 (normal) or 0 (exceptional). The queried origin server must be bound to a listener or rule, and the ID of the bound listener or rule must be specified for the query. This API supports displaying origin server status statistics at a 1-minute granularity.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRealServerStatistics(request *DescribeRealServerStatisticsRequest) (response *DescribeRealServerStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeRealServerStatisticsRequest()
    }
    response = NewDescribeRealServerStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRealServersRequest() (request *DescribeRealServersRequest) {
    request = &DescribeRealServersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeRealServers")
    return
}

func NewDescribeRealServersResponse() (response *DescribeRealServersResponse) {
    response = &DescribeRealServersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRealServers
// This API is used to query origin server information. It can query all origin servers under the specified project name, and supports fuzzy query by specified IPs or domain names.
//
// error code that may be returned:
//  AUTHFAILURE_NOTENTERPRISEAUTHORIZATION = "AuthFailure.NotEnterpriseAuthorization"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeRealServers(request *DescribeRealServersRequest) (response *DescribeRealServersResponse, err error) {
    if request == nil {
        request = NewDescribeRealServersRequest()
    }
    response = NewDescribeRealServersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRealServersStatusRequest() (request *DescribeRealServersStatusRequest) {
    request = &DescribeRealServersStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeRealServersStatus")
    return
}

func NewDescribeRealServersStatusResponse() (response *DescribeRealServersStatusResponse) {
    response = &DescribeRealServersStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRealServersStatus
// This API (DescribeRealServersStatus) is used to query whether an origin server has been bound to a rule or listener.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REALSERVERNOTINPROJECT = "FailedOperation.RealServerNotInProject"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_REALSERVERNOTBELONG = "InvalidParameterValue.RealServerNotBelong"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRealServersStatus(request *DescribeRealServersStatusRequest) (response *DescribeRealServersStatusResponse, err error) {
    if request == nil {
        request = NewDescribeRealServersStatusRequest()
    }
    response = NewDescribeRealServersStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRegionAndPriceRequest() (request *DescribeRegionAndPriceRequest) {
    request = &DescribeRegionAndPriceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeRegionAndPrice")
    return
}

func NewDescribeRegionAndPriceResponse() (response *DescribeRegionAndPriceResponse) {
    response = &DescribeRegionAndPriceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRegionAndPrice
// This API (DescribeRegionAndPrice) is used to obtain the origin server regions and the bandwidth price gradient.
//
// error code that may be returned:
//  AUTHFAILURE_NOTENTERPRISEAUTHORIZATION = "AuthFailure.NotEnterpriseAuthorization"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeRegionAndPrice(request *DescribeRegionAndPriceRequest) (response *DescribeRegionAndPriceResponse, err error) {
    if request == nil {
        request = NewDescribeRegionAndPriceRequest()
    }
    response = NewDescribeRegionAndPriceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourcesByTagRequest() (request *DescribeResourcesByTagRequest) {
    request = &DescribeResourcesByTagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeResourcesByTag")
    return
}

func NewDescribeResourcesByTagResponse() (response *DescribeResourcesByTagResponse) {
    response = &DescribeResourcesByTagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeResourcesByTag
// This API (DescribeResourcesByTag) is used to query corresponding resource information by tags, including connection, connection group, and origin server.
//
// error code that may be returned:
//  AUTHFAILURE_NOTENTERPRISEAUTHORIZATION = "AuthFailure.NotEnterpriseAuthorization"
//  FAILEDOPERATION_DUPLICATEDREQUEST = "FailedOperation.DuplicatedRequest"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeResourcesByTag(request *DescribeResourcesByTagRequest) (response *DescribeResourcesByTagResponse, err error) {
    if request == nil {
        request = NewDescribeResourcesByTagRequest()
    }
    response = NewDescribeResourcesByTagResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRuleRealServersRequest() (request *DescribeRuleRealServersRequest) {
    request = &DescribeRuleRealServersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeRuleRealServers")
    return
}

func NewDescribeRuleRealServersResponse() (response *DescribeRuleRealServersResponse) {
    response = &DescribeRuleRealServersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRuleRealServers
// This API (DescribeRuleRealServers) is used to query forwarding rules-related origin server information, including information of origin servers that can be bound and have been bound.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCEUPGRADING = "FailedOperation.ResourceUpgrading"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRuleRealServers(request *DescribeRuleRealServersRequest) (response *DescribeRuleRealServersResponse, err error) {
    if request == nil {
        request = NewDescribeRuleRealServersRequest()
    }
    response = NewDescribeRuleRealServersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRulesRequest() (request *DescribeRulesRequest) {
    request = &DescribeRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeRules")
    return
}

func NewDescribeRulesResponse() (response *DescribeRulesResponse) {
    response = &DescribeRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRules
// This API (DescribeRules) is used to query all rule information of a listener, including the domain names, paths, and list of bound origin servers. For version 3.0 connections, this API returns the advanced authentication configuration information of the domain name.
//
// error code that may be returned:
//  AUTHFAILURE_NOTENTERPRISEAUTHORIZATION = "AuthFailure.NotEnterpriseAuthorization"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeRules(request *DescribeRulesRequest) (response *DescribeRulesResponse, err error) {
    if request == nil {
        request = NewDescribeRulesRequest()
    }
    response = NewDescribeRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRulesByRuleIdsRequest() (request *DescribeRulesByRuleIdsRequest) {
    request = &DescribeRulesByRuleIdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeRulesByRuleIds")
    return
}

func NewDescribeRulesByRuleIdsResponse() (response *DescribeRulesByRuleIdsResponse) {
    response = &DescribeRulesByRuleIdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRulesByRuleIds
// This API is used to pull the list of rules based on rule ID. It supports pulling 1 to 10 rules at a time.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCECANNOTACCESS = "FailedOperation.ResourceCanNotAccess"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRulesByRuleIds(request *DescribeRulesByRuleIdsRequest) (response *DescribeRulesByRuleIdsResponse, err error) {
    if request == nil {
        request = NewDescribeRulesByRuleIdsRequest()
    }
    response = NewDescribeRulesByRuleIdsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityPolicyDetailRequest() (request *DescribeSecurityPolicyDetailRequest) {
    request = &DescribeSecurityPolicyDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeSecurityPolicyDetail")
    return
}

func NewDescribeSecurityPolicyDetailResponse() (response *DescribeSecurityPolicyDetailResponse) {
    response = &DescribeSecurityPolicyDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSecurityPolicyDetail
// This API is used to obtain security policy details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeSecurityPolicyDetail(request *DescribeSecurityPolicyDetailRequest) (response *DescribeSecurityPolicyDetailResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityPolicyDetailRequest()
    }
    response = NewDescribeSecurityPolicyDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityRulesRequest() (request *DescribeSecurityRulesRequest) {
    request = &DescribeSecurityRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeSecurityRules")
    return
}

func NewDescribeSecurityRulesResponse() (response *DescribeSecurityRulesResponse) {
    response = &DescribeSecurityRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSecurityRules
// This API is used to query the list of security rules based on security rule ID. It supports querying 1 to 20 security rules at a time.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCECANNOTACCESS = "FailedOperation.ResourceCanNotAccess"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSecurityRules(request *DescribeSecurityRulesRequest) (response *DescribeSecurityRulesResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityRulesRequest()
    }
    response = NewDescribeSecurityRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTCPListenersRequest() (request *DescribeTCPListenersRequest) {
    request = &DescribeTCPListenersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeTCPListeners")
    return
}

func NewDescribeTCPListenersResponse() (response *DescribeTCPListenersResponse) {
    response = &DescribeTCPListenersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTCPListeners
// This API (DescribeTCPListeners) is used to query the TCP listener information of a single connection or connection group.
//
// error code that may be returned:
//  AUTHFAILURE_NOTENTERPRISEAUTHORIZATION = "AuthFailure.NotEnterpriseAuthorization"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTCPListeners(request *DescribeTCPListenersRequest) (response *DescribeTCPListenersResponse, err error) {
    if request == nil {
        request = NewDescribeTCPListenersRequest()
    }
    response = NewDescribeTCPListenersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUDPListenersRequest() (request *DescribeUDPListenersRequest) {
    request = &DescribeUDPListenersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DescribeUDPListeners")
    return
}

func NewDescribeUDPListenersResponse() (response *DescribeUDPListenersResponse) {
    response = &DescribeUDPListenersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUDPListeners
// This API (DescribeUDPListeners) is used to query the UDP listener information of a single connection or connection group.
//
// error code that may be returned:
//  AUTHFAILURE_NOTENTERPRISEAUTHORIZATION = "AuthFailure.NotEnterpriseAuthorization"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeUDPListeners(request *DescribeUDPListenersRequest) (response *DescribeUDPListenersResponse, err error) {
    if request == nil {
        request = NewDescribeUDPListenersRequest()
    }
    response = NewDescribeUDPListenersResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyProxiesRequest() (request *DestroyProxiesRequest) {
    request = &DestroyProxiesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "DestroyProxies")
    return
}

func NewDestroyProxiesResponse() (response *DestroyProxiesResponse) {
    response = &DestroyProxiesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DestroyProxies
// This API (DestroyProxies) is used to terminate a connection. If terminated, no fees will be incurred.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCOUNTBALANCEINSUFFICIENT = "FailedOperation.AccountBalanceInsufficient"
//  FAILEDOPERATION_ACTIONISDOING = "FailedOperation.ActionIsDoing"
//  FAILEDOPERATION_BELONGDIFFERENTGROUP = "FailedOperation.BelongDifferentGroup"
//  FAILEDOPERATION_DUPLICATEDREQUEST = "FailedOperation.DuplicatedRequest"
//  FAILEDOPERATION_REALSERVERALREADYBOUND = "FailedOperation.RealServerAlreadyBound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DestroyProxies(request *DestroyProxiesRequest) (response *DestroyProxiesResponse, err error) {
    if request == nil {
        request = NewDestroyProxiesRequest()
    }
    response = NewDestroyProxiesResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceCreateProxyRequest() (request *InquiryPriceCreateProxyRequest) {
    request = &InquiryPriceCreateProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "InquiryPriceCreateProxy")
    return
}

func NewInquiryPriceCreateProxyResponse() (response *InquiryPriceCreateProxyResponse) {
    response = &InquiryPriceCreateProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquiryPriceCreateProxy
// This API (InquiryPriceCreateProxy) is used to create the price inquiries of acceleration connections.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCEUPGRADING = "FailedOperation.ResourceUpgrading"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTH = "InvalidParameterValue.InvalidBandwidth"
//  INVALIDPARAMETERVALUE_INVALIDCONCURRENCY = "InvalidParameterValue.InvalidConcurrency"
//  INVALIDPARAMETERVALUE_UNKNOWNACCESSREGION = "InvalidParameterValue.UnknownAccessRegion"
//  INVALIDPARAMETERVALUE_UNKNOWNDESTREGION = "InvalidParameterValue.UnknownDestRegion"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) InquiryPriceCreateProxy(request *InquiryPriceCreateProxyRequest) (response *InquiryPriceCreateProxyResponse, err error) {
    if request == nil {
        request = NewInquiryPriceCreateProxyRequest()
    }
    response = NewInquiryPriceCreateProxyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCertificateRequest() (request *ModifyCertificateRequest) {
    request = &ModifyCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "ModifyCertificate")
    return
}

func NewModifyCertificateResponse() (response *ModifyCertificateResponse) {
    response = &ModifyCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCertificate
// This API (ModifyCertificate) is used to modify a domain name certificate of a listener. domain name certificate. This API is only applicable to connections of version 3.0.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINSTATUSNOTINRUNNING = "FailedOperation.DomainStatusNotInRunning"
//  FAILEDOPERATION_LISTENERHASTASK = "FailedOperation.ListenerHasTask"
//  FAILEDOPERATION_PROXYSTATUSNOTINRUNING = "FailedOperation.ProxyStatusNotInRuning"
//  FAILEDOPERATION_PROXYVERSIONNOTSUPPORT = "FailedOperation.ProxyVersionNotSupport"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyCertificate(request *ModifyCertificateRequest) (response *ModifyCertificateResponse, err error) {
    if request == nil {
        request = NewModifyCertificateRequest()
    }
    response = NewModifyCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCertificateAttributesRequest() (request *ModifyCertificateAttributesRequest) {
    request = &ModifyCertificateAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "ModifyCertificateAttributes")
    return
}

func NewModifyCertificateAttributesResponse() (response *ModifyCertificateAttributesResponse) {
    response = &ModifyCertificateAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCertificateAttributes
// This API is used to modify certificate name and content.
//
// error code that may be returned:
//  FAILEDOPERATION_CERTIFICATEISUSING = "FailedOperation.CertificateIsUsing"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCertificateAttributes(request *ModifyCertificateAttributesRequest) (response *ModifyCertificateAttributesResponse, err error) {
    if request == nil {
        request = NewModifyCertificateAttributesRequest()
    }
    response = NewModifyCertificateAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDomainRequest() (request *ModifyDomainRequest) {
    request = &ModifyDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "ModifyDomain")
    return
}

func NewModifyDomainResponse() (response *ModifyDomainResponse) {
    response = &ModifyDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDomain
// This API (ModifyDomain) is used to modify domain names of listeners. For connections of version 3.0, it supports modifying certificates of the domain names.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINALREADYEXISTED = "FailedOperation.DomainAlreadyExisted"
//  FAILEDOPERATION_DOMAINSTATUSNOTINRUNNING = "FailedOperation.DomainStatusNotInRunning"
//  FAILEDOPERATION_INSTANCESTATUSNOTINRUNING = "FailedOperation.InstanceStatusNotInRuning"
//  FAILEDOPERATION_INVALIDLISTENERPROTOCOL = "FailedOperation.InvalidListenerProtocol"
//  FAILEDOPERATION_LISTENERHASTASK = "FailedOperation.ListenerHasTask"
//  FAILEDOPERATION_PROXYSTATUSNOTINRUNING = "FailedOperation.ProxyStatusNotInRuning"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyDomain(request *ModifyDomainRequest) (response *ModifyDomainResponse, err error) {
    if request == nil {
        request = NewModifyDomainRequest()
    }
    response = NewModifyDomainResponse()
    err = c.Send(request, response)
    return
}

func NewModifyGroupDomainConfigRequest() (request *ModifyGroupDomainConfigRequest) {
    request = &ModifyGroupDomainConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "ModifyGroupDomainConfig")
    return
}

func NewModifyGroupDomainConfigResponse() (response *ModifyGroupDomainConfigResponse) {
    response = &ModifyGroupDomainConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyGroupDomainConfig
// This API (ModifyGroupDomainConfig) is used to configure the nearest access domain name of a connection group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyGroupDomainConfig(request *ModifyGroupDomainConfigRequest) (response *ModifyGroupDomainConfigResponse, err error) {
    if request == nil {
        request = NewModifyGroupDomainConfigRequest()
    }
    response = NewModifyGroupDomainConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyHTTPListenerAttributeRequest() (request *ModifyHTTPListenerAttributeRequest) {
    request = &ModifyHTTPListenerAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "ModifyHTTPListenerAttribute")
    return
}

func NewModifyHTTPListenerAttributeResponse() (response *ModifyHTTPListenerAttributeResponse) {
    response = &ModifyHTTPListenerAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyHTTPListenerAttribute
// This API (ModifyHTTPListenerAttribute) is used to modify the HTTP listener configuration information of a connection. Currently only supports modifying listener name.
//
// Note: Grouped connections currently do not support HTTP/HTTPS listeners.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LISTENERSTATUSERROR = "FailedOperation.ListenerStatusError"
//  FAILEDOPERATION_NOTSUPPORTOLDVERSIONPROXY = "FailedOperation.NotSupportOldVersionProxy"
//  FAILEDOPERATION_NOTSUPPORTPROXYGROUP = "FailedOperation.NotSupportProxyGroup"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyHTTPListenerAttribute(request *ModifyHTTPListenerAttributeRequest) (response *ModifyHTTPListenerAttributeResponse, err error) {
    if request == nil {
        request = NewModifyHTTPListenerAttributeRequest()
    }
    response = NewModifyHTTPListenerAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyHTTPSListenerAttributeRequest() (request *ModifyHTTPSListenerAttributeRequest) {
    request = &ModifyHTTPSListenerAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "ModifyHTTPSListenerAttribute")
    return
}

func NewModifyHTTPSListenerAttributeResponse() (response *ModifyHTTPSListenerAttributeResponse) {
    response = &ModifyHTTPSListenerAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyHTTPSListenerAttribute
// This API (ModifyHTTPSListenerAttribute) is used to modify HTTPS listener configurations. It currently do not support connection groups and connections of version 1.0.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GROUPSTATUSNOTINRUNING = "FailedOperation.GroupStatusNotInRuning"
//  FAILEDOPERATION_INSTANCESTATUSNOTINRUNING = "FailedOperation.InstanceStatusNotInRuning"
//  FAILEDOPERATION_LISTENERHASTASK = "FailedOperation.ListenerHasTask"
//  FAILEDOPERATION_NOTSUPPORTOLDVERSIONPROXY = "FailedOperation.NotSupportOldVersionProxy"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyHTTPSListenerAttribute(request *ModifyHTTPSListenerAttributeRequest) (response *ModifyHTTPSListenerAttributeResponse, err error) {
    if request == nil {
        request = NewModifyHTTPSListenerAttributeRequest()
    }
    response = NewModifyHTTPSListenerAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyProxiesAttributeRequest() (request *ModifyProxiesAttributeRequest) {
    request = &ModifyProxiesAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "ModifyProxiesAttribute")
    return
}

func NewModifyProxiesAttributeResponse() (response *ModifyProxiesAttributeResponse) {
    response = &ModifyProxiesAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyProxiesAttribute
// This API (ModifyProxiesAttribute) is used to modify instance attributes (currently only supports modifying connection name).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCOUNTBALANCEINSUFFICIENT = "FailedOperation.AccountBalanceInsufficient"
//  FAILEDOPERATION_DUPLICATEDREQUEST = "FailedOperation.DuplicatedRequest"
//  FAILEDOPERATION_RESOURCECANNOTACCESS = "FailedOperation.ResourceCanNotAccess"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PROJECTIDNOTBELONG = "InvalidParameterValue.ProjectIdNotBelong"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyProxiesAttribute(request *ModifyProxiesAttributeRequest) (response *ModifyProxiesAttributeResponse, err error) {
    if request == nil {
        request = NewModifyProxiesAttributeRequest()
    }
    response = NewModifyProxiesAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyProxiesProjectRequest() (request *ModifyProxiesProjectRequest) {
    request = &ModifyProxiesProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "ModifyProxiesProject")
    return
}

func NewModifyProxiesProjectResponse() (response *ModifyProxiesProjectResponse) {
    response = &ModifyProxiesProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyProxiesProject
// This API (ModifyProxiesProject) is used to modify the project to which a connection belongs.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCOUNTBALANCEINSUFFICIENT = "FailedOperation.AccountBalanceInsufficient"
//  FAILEDOPERATION_DUPLICATEDREQUEST = "FailedOperation.DuplicatedRequest"
//  FAILEDOPERATION_NOTSUPPORTOLDVERSIONPROXY = "FailedOperation.NotSupportOldVersionProxy"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyProxiesProject(request *ModifyProxiesProjectRequest) (response *ModifyProxiesProjectResponse, err error) {
    if request == nil {
        request = NewModifyProxiesProjectRequest()
    }
    response = NewModifyProxiesProjectResponse()
    err = c.Send(request, response)
    return
}

func NewModifyProxyConfigurationRequest() (request *ModifyProxyConfigurationRequest) {
    request = &ModifyProxyConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "ModifyProxyConfiguration")
    return
}

func NewModifyProxyConfigurationResponse() (response *ModifyProxyConfigurationResponse) {
    response = &ModifyProxyConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyProxyConfiguration
// This API (ModifyProxyConfiguration) is used to modify connection configurations. You can expand or reduce the capacity based on current business requirements. It only supports connections with a Scalarable of 1. Scalarable can be obtained via DescribeProxies API.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCOUNTBALANCEINSUFFICIENT = "FailedOperation.AccountBalanceInsufficient"
//  FAILEDOPERATION_DUPLICATEDREQUEST = "FailedOperation.DuplicatedRequest"
//  FAILEDOPERATION_NOTSUPPORTSCALAR = "FailedOperation.NotSupportScalar"
//  FAILEDOPERATION_REALSERVERALREADYBOUND = "FailedOperation.RealServerAlreadyBound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyProxyConfiguration(request *ModifyProxyConfigurationRequest) (response *ModifyProxyConfigurationResponse, err error) {
    if request == nil {
        request = NewModifyProxyConfigurationRequest()
    }
    response = NewModifyProxyConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyProxyGroupAttributeRequest() (request *ModifyProxyGroupAttributeRequest) {
    request = &ModifyProxyGroupAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "ModifyProxyGroupAttribute")
    return
}

func NewModifyProxyGroupAttributeResponse() (response *ModifyProxyGroupAttributeResponse) {
    response = &ModifyProxyGroupAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyProxyGroupAttribute
// This API (ModifyProxyGroupAttribute) is used to modify connection group attributes. It currently only supports modifying connection group name.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyProxyGroupAttribute(request *ModifyProxyGroupAttributeRequest) (response *ModifyProxyGroupAttributeResponse, err error) {
    if request == nil {
        request = NewModifyProxyGroupAttributeRequest()
    }
    response = NewModifyProxyGroupAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRealServerNameRequest() (request *ModifyRealServerNameRequest) {
    request = &ModifyRealServerNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "ModifyRealServerName")
    return
}

func NewModifyRealServerNameResponse() (response *ModifyRealServerNameResponse) {
    response = &ModifyRealServerNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyRealServerName
// This API (ModifyRealServerName) is used to modify origin server names.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyRealServerName(request *ModifyRealServerNameRequest) (response *ModifyRealServerNameResponse, err error) {
    if request == nil {
        request = NewModifyRealServerNameRequest()
    }
    response = NewModifyRealServerNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRuleAttributeRequest() (request *ModifyRuleAttributeRequest) {
    request = &ModifyRuleAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "ModifyRuleAttribute")
    return
}

func NewModifyRuleAttributeResponse() (response *ModifyRuleAttributeResponse) {
    response = &ModifyRuleAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyRuleAttribute
// This API (ModifyRuleAttribute) is used to modify forwarding rule information, including health check configuration and forwarding policies.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCESTATUSNOTINRUNING = "FailedOperation.InstanceStatusNotInRuning"
//  FAILEDOPERATION_LISTENERHASTASK = "FailedOperation.ListenerHasTask"
//  FAILEDOPERATION_RULEALREADYEXISTED = "FailedOperation.RuleAlreadyExisted"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyRuleAttribute(request *ModifyRuleAttributeRequest) (response *ModifyRuleAttributeResponse, err error) {
    if request == nil {
        request = NewModifyRuleAttributeRequest()
    }
    response = NewModifyRuleAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifySecurityRuleRequest() (request *ModifySecurityRuleRequest) {
    request = &ModifySecurityRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "ModifySecurityRule")
    return
}

func NewModifySecurityRuleResponse() (response *ModifySecurityRuleResponse) {
    response = &ModifySecurityRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySecurityRule
// This API is used to modify security policy rule names.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROXYSECURITYPOLICYDEFAULTRULE = "FailedOperation.ProxySecurityPolicyDefaultRule"
//  FAILEDOPERATION_PROXYSECURITYPOLICYDUPLICATEDRULE = "FailedOperation.ProxySecurityPolicyDuplicatedRule"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifySecurityRule(request *ModifySecurityRuleRequest) (response *ModifySecurityRuleResponse, err error) {
    if request == nil {
        request = NewModifySecurityRuleRequest()
    }
    response = NewModifySecurityRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTCPListenerAttributeRequest() (request *ModifyTCPListenerAttributeRequest) {
    request = &ModifyTCPListenerAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "ModifyTCPListenerAttribute")
    return
}

func NewModifyTCPListenerAttributeResponse() (response *ModifyTCPListenerAttributeResponse) {
    response = &ModifyTCPListenerAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyTCPListenerAttribute
// This API (ModifyTCPListenerAttribute) is used to modify the TCP listener configuration of a connection instance, including health check configuration and scheduling policies.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GROUPSTATUSNOTINRUNING = "FailedOperation.GroupStatusNotInRuning"
//  FAILEDOPERATION_INSTANCESTATUSNOTINRUNING = "FailedOperation.InstanceStatusNotInRuning"
//  FAILEDOPERATION_LISTENERHASTASK = "FailedOperation.ListenerHasTask"
//  FAILEDOPERATION_LISTENERSTATUSERROR = "FailedOperation.ListenerStatusError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyTCPListenerAttribute(request *ModifyTCPListenerAttributeRequest) (response *ModifyTCPListenerAttributeResponse, err error) {
    if request == nil {
        request = NewModifyTCPListenerAttributeRequest()
    }
    response = NewModifyTCPListenerAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUDPListenerAttributeRequest() (request *ModifyUDPListenerAttributeRequest) {
    request = &ModifyUDPListenerAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "ModifyUDPListenerAttribute")
    return
}

func NewModifyUDPListenerAttributeResponse() (response *ModifyUDPListenerAttributeResponse) {
    response = &ModifyUDPListenerAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyUDPListenerAttribute
// This API (ModifyUDPListenerAttribute) is used to modify the UDP listener configuration of a connection instance, including modification of listener names and scheduling policies.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GROUPSTATUSNOTINRUNING = "FailedOperation.GroupStatusNotInRuning"
//  FAILEDOPERATION_INSTANCESTATUSNOTINRUNING = "FailedOperation.InstanceStatusNotInRuning"
//  FAILEDOPERATION_LISTENERHASTASK = "FailedOperation.ListenerHasTask"
//  FAILEDOPERATION_LISTENERSTATUSERROR = "FailedOperation.ListenerStatusError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyUDPListenerAttribute(request *ModifyUDPListenerAttributeRequest) (response *ModifyUDPListenerAttributeResponse, err error) {
    if request == nil {
        request = NewModifyUDPListenerAttributeRequest()
    }
    response = NewModifyUDPListenerAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewOpenProxiesRequest() (request *OpenProxiesRequest) {
    request = &OpenProxiesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "OpenProxies")
    return
}

func NewOpenProxiesResponse() (response *OpenProxiesResponse) {
    response = &OpenProxiesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// OpenProxies
// This API (OpenProxies) is used to enable one or more connections.
//
// error code that may be returned:
//  AUTHFAILURE_NOTENTERPRISEAUTHORIZATION = "AuthFailure.NotEnterpriseAuthorization"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCOUNTBALANCEINSUFFICIENT = "FailedOperation.AccountBalanceInsufficient"
//  FAILEDOPERATION_DUPLICATEDREQUEST = "FailedOperation.DuplicatedRequest"
//  FAILEDOPERATION_NOTSUPPORTOLDVERSIONPROXY = "FailedOperation.NotSupportOldVersionProxy"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) OpenProxies(request *OpenProxiesRequest) (response *OpenProxiesResponse, err error) {
    if request == nil {
        request = NewOpenProxiesRequest()
    }
    response = NewOpenProxiesResponse()
    err = c.Send(request, response)
    return
}

func NewOpenProxyGroupRequest() (request *OpenProxyGroupRequest) {
    request = &OpenProxyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "OpenProxyGroup")
    return
}

func NewOpenProxyGroupResponse() (response *OpenProxyGroupResponse) {
    response = &OpenProxyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// OpenProxyGroup
// This API is used to enable all connections in a connection group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCOUNTBALANCEINSUFFICIENT = "FailedOperation.AccountBalanceInsufficient"
//  FAILEDOPERATION_DUPLICATEDREQUEST = "FailedOperation.DuplicatedRequest"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) OpenProxyGroup(request *OpenProxyGroupRequest) (response *OpenProxyGroupResponse, err error) {
    if request == nil {
        request = NewOpenProxyGroupRequest()
    }
    response = NewOpenProxyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewOpenSecurityPolicyRequest() (request *OpenSecurityPolicyRequest) {
    request = &OpenSecurityPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "OpenSecurityPolicy")
    return
}

func NewOpenSecurityPolicyResponse() (response *OpenSecurityPolicyResponse) {
    response = &OpenSecurityPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// OpenSecurityPolicy
// This API is used to enable security policies.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACTIONISDOING = "FailedOperation.ActionIsDoing"
//  FAILEDOPERATION_PROXYSECURITYALREADYOPEN = "FailedOperation.ProxySecurityAlreadyOpen"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) OpenSecurityPolicy(request *OpenSecurityPolicyRequest) (response *OpenSecurityPolicyResponse, err error) {
    if request == nil {
        request = NewOpenSecurityPolicyRequest()
    }
    response = NewOpenSecurityPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveRealServersRequest() (request *RemoveRealServersRequest) {
    request = &RemoveRealServersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "RemoveRealServers")
    return
}

func NewRemoveRealServersResponse() (response *RemoveRealServersResponse) {
    response = &RemoveRealServersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RemoveRealServers
// This API is used to delete the added origin server (server) IP or domain name.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REALSERVERALREADYBOUND = "FailedOperation.RealServerAlreadyBound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) RemoveRealServers(request *RemoveRealServersRequest) (response *RemoveRealServersResponse, err error) {
    if request == nil {
        request = NewRemoveRealServersRequest()
    }
    response = NewRemoveRealServersResponse()
    err = c.Send(request, response)
    return
}

func NewSetAuthenticationRequest() (request *SetAuthenticationRequest) {
    request = &SetAuthenticationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gaap", APIVersion, "SetAuthentication")
    return
}

func NewSetAuthenticationResponse() (response *SetAuthenticationResponse) {
    response = &SetAuthenticationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetAuthentication
// This API (SetAuthentication) is used for the advanced authentication configuration of connections, including authentication methods and their certificates. If only supports connections of version 3.0.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACTIONISDOING = "FailedOperation.ActionIsDoing"
//  FAILEDOPERATION_LISTENERHASTASK = "FailedOperation.ListenerHasTask"
//  FAILEDOPERATION_LISTENERSTATUSERROR = "FailedOperation.ListenerStatusError"
//  FAILEDOPERATION_NOTSUPPORTOLDVERSIONPROXY = "FailedOperation.NotSupportOldVersionProxy"
//  FAILEDOPERATION_PROXYVERSIONNOTSUPPORT = "FailedOperation.ProxyVersionNotSupport"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDCERTIFICATEID = "InvalidParameterValue.InvalidCertificateId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) SetAuthentication(request *SetAuthenticationRequest) (response *SetAuthenticationResponse, err error) {
    if request == nil {
        request = NewSetAuthenticationRequest()
    }
    response = NewSetAuthenticationResponse()
    err = c.Send(request, response)
    return
}
