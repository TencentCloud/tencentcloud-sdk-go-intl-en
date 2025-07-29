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

package v20180125

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2018-01-25"

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


func NewAddAntiFakeUrlRequest() (request *AddAntiFakeUrlRequest) {
    request = &AddAntiFakeUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "AddAntiFakeUrl")
    
    
    return
}

func NewAddAntiFakeUrlResponse() (response *AddAntiFakeUrlResponse) {
    response = &AddAntiFakeUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddAntiFakeUrl
// Add tamper-proof URL
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_SPECIFICATIONERR = "LimitExceeded.SpecificationErr"
func (c *Client) AddAntiFakeUrl(request *AddAntiFakeUrlRequest) (response *AddAntiFakeUrlResponse, err error) {
    return c.AddAntiFakeUrlWithContext(context.Background(), request)
}

// AddAntiFakeUrl
// Add tamper-proof URL
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_SPECIFICATIONERR = "LimitExceeded.SpecificationErr"
func (c *Client) AddAntiFakeUrlWithContext(ctx context.Context, request *AddAntiFakeUrlRequest) (response *AddAntiFakeUrlResponse, err error) {
    if request == nil {
        request = NewAddAntiFakeUrlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "AddAntiFakeUrl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddAntiFakeUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddAntiFakeUrlResponse()
    err = c.Send(request, response)
    return
}

func NewAddAntiInfoLeakRulesRequest() (request *AddAntiInfoLeakRulesRequest) {
    request = &AddAntiInfoLeakRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "AddAntiInfoLeakRules")
    
    
    return
}

func NewAddAntiInfoLeakRulesResponse() (response *AddAntiInfoLeakRulesResponse) {
    response = &AddAntiInfoLeakRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddAntiInfoLeakRules
// Add information leakage prevention rule
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_SPECIFICATIONERR = "LimitExceeded.SpecificationErr"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddAntiInfoLeakRules(request *AddAntiInfoLeakRulesRequest) (response *AddAntiInfoLeakRulesResponse, err error) {
    return c.AddAntiInfoLeakRulesWithContext(context.Background(), request)
}

// AddAntiInfoLeakRules
// Add information leakage prevention rule
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_SPECIFICATIONERR = "LimitExceeded.SpecificationErr"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddAntiInfoLeakRulesWithContext(ctx context.Context, request *AddAntiInfoLeakRulesRequest) (response *AddAntiInfoLeakRulesResponse, err error) {
    if request == nil {
        request = NewAddAntiInfoLeakRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "AddAntiInfoLeakRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddAntiInfoLeakRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddAntiInfoLeakRulesResponse()
    err = c.Send(request, response)
    return
}

func NewAddCustomRuleRequest() (request *AddCustomRuleRequest) {
    request = &AddCustomRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "AddCustomRule")
    
    
    return
}

func NewAddCustomRuleResponse() (response *AddCustomRuleResponse) {
    response = &AddCustomRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddCustomRule
// Add access control (from custom policy)
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_SPECIFICATIONERR = "LimitExceeded.SpecificationErr"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddCustomRule(request *AddCustomRuleRequest) (response *AddCustomRuleResponse, err error) {
    return c.AddCustomRuleWithContext(context.Background(), request)
}

// AddCustomRule
// Add access control (from custom policy)
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_SPECIFICATIONERR = "LimitExceeded.SpecificationErr"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddCustomRuleWithContext(ctx context.Context, request *AddCustomRuleRequest) (response *AddCustomRuleResponse, err error) {
    if request == nil {
        request = NewAddCustomRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "AddCustomRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddCustomRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddCustomRuleResponse()
    err = c.Send(request, response)
    return
}

func NewAddCustomWhiteRuleRequest() (request *AddCustomWhiteRuleRequest) {
    request = &AddCustomWhiteRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "AddCustomWhiteRule")
    
    
    return
}

func NewAddCustomWhiteRuleResponse() (response *AddCustomWhiteRuleResponse) {
    response = &AddCustomWhiteRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddCustomWhiteRule
// Add precision allowlist rules
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_SPECIFICATIONERR = "LimitExceeded.SpecificationErr"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddCustomWhiteRule(request *AddCustomWhiteRuleRequest) (response *AddCustomWhiteRuleResponse, err error) {
    return c.AddCustomWhiteRuleWithContext(context.Background(), request)
}

// AddCustomWhiteRule
// Add precision allowlist rules
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_SPECIFICATIONERR = "LimitExceeded.SpecificationErr"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddCustomWhiteRuleWithContext(ctx context.Context, request *AddCustomWhiteRuleRequest) (response *AddCustomWhiteRuleResponse, err error) {
    if request == nil {
        request = NewAddCustomWhiteRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "AddCustomWhiteRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddCustomWhiteRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddCustomWhiteRuleResponse()
    err = c.Send(request, response)
    return
}

func NewAddSpartaProtectionRequest() (request *AddSpartaProtectionRequest) {
    request = &AddSpartaProtectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "AddSpartaProtection")
    
    
    return
}

func NewAddSpartaProtectionResponse() (response *AddSpartaProtectionResponse) {
    response = &AddSpartaProtectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddSpartaProtection
// Add SaaS WAF protection domain
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REDISOPERATIONFAILED = "FailedOperation.RedisOperationFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ASYNCHRONOUSCALLFAILED = "InternalError.AsynchronousCallFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CERTIFICATIONPARAMETERERR = "InvalidParameter.CertificationParameterErr"
//  INVALIDPARAMETER_DOMAINEXCEEDSLIMITERR = "InvalidParameter.DomainExceedsLimitErr"
//  INVALIDPARAMETER_DOMAINNOTRECORD = "InvalidParameter.DomainNotRecord"
//  INVALIDPARAMETER_PORTPARAMETERERR = "InvalidParameter.PortParameterErr"
//  INVALIDPARAMETER_PROTECTIONDOMAINPARAMETERERR = "InvalidParameter.ProtectionDomainParameterErr"
//  INVALIDPARAMETER_TLSPARAMETERERR = "InvalidParameter.TLSParameterErr"
//  INVALIDPARAMETER_UNAUTHORIZEDOPERATIONPARAMETERERR = "InvalidParameter.UnauthorizedOperationParameterErr"
//  INVALIDPARAMETER_UPSTREAMPARAMETERERR = "InvalidParameter.UpstreamParameterErr"
//  INVALIDPARAMETER_XFFRESETPARAMETERERR = "InvalidParameter.XFFResetParameterErr"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_EMPTYERR = "ResourceInUse.EmptyErr"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) AddSpartaProtection(request *AddSpartaProtectionRequest) (response *AddSpartaProtectionResponse, err error) {
    return c.AddSpartaProtectionWithContext(context.Background(), request)
}

// AddSpartaProtection
// Add SaaS WAF protection domain
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REDISOPERATIONFAILED = "FailedOperation.RedisOperationFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ASYNCHRONOUSCALLFAILED = "InternalError.AsynchronousCallFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CERTIFICATIONPARAMETERERR = "InvalidParameter.CertificationParameterErr"
//  INVALIDPARAMETER_DOMAINEXCEEDSLIMITERR = "InvalidParameter.DomainExceedsLimitErr"
//  INVALIDPARAMETER_DOMAINNOTRECORD = "InvalidParameter.DomainNotRecord"
//  INVALIDPARAMETER_PORTPARAMETERERR = "InvalidParameter.PortParameterErr"
//  INVALIDPARAMETER_PROTECTIONDOMAINPARAMETERERR = "InvalidParameter.ProtectionDomainParameterErr"
//  INVALIDPARAMETER_TLSPARAMETERERR = "InvalidParameter.TLSParameterErr"
//  INVALIDPARAMETER_UNAUTHORIZEDOPERATIONPARAMETERERR = "InvalidParameter.UnauthorizedOperationParameterErr"
//  INVALIDPARAMETER_UPSTREAMPARAMETERERR = "InvalidParameter.UpstreamParameterErr"
//  INVALIDPARAMETER_XFFRESETPARAMETERERR = "InvalidParameter.XFFResetParameterErr"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_EMPTYERR = "ResourceInUse.EmptyErr"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) AddSpartaProtectionWithContext(ctx context.Context, request *AddSpartaProtectionRequest) (response *AddSpartaProtectionResponse, err error) {
    if request == nil {
        request = NewAddSpartaProtectionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "AddSpartaProtection")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddSpartaProtection require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddSpartaProtectionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDealsRequest() (request *CreateDealsRequest) {
    request = &CreateDealsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "CreateDeals")
    
    
    return
}

func NewCreateDealsResponse() (response *CreateDealsResponse) {
    response = &CreateDealsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDeals
// Billing Resource Purchase, Renewal Order API
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDeals(request *CreateDealsRequest) (response *CreateDealsResponse, err error) {
    return c.CreateDealsWithContext(context.Background(), request)
}

// CreateDeals
// Billing Resource Purchase, Renewal Order API
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDealsWithContext(ctx context.Context, request *CreateDealsRequest) (response *CreateDealsResponse, err error) {
    if request == nil {
        request = NewCreateDealsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "CreateDeals")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDeals require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDealsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateHostRequest() (request *CreateHostRequest) {
    request = &CreateHostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "CreateHost")
    
    
    return
}

func NewCreateHostResponse() (response *CreateHostResponse) {
    response = &CreateHostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateHost
// Add a protection domain in CLB-WAF
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateHost(request *CreateHostRequest) (response *CreateHostResponse, err error) {
    return c.CreateHostWithContext(context.Background(), request)
}

// CreateHost
// Add a protection domain in CLB-WAF
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateHostWithContext(ctx context.Context, request *CreateHostRequest) (response *CreateHostResponse, err error) {
    if request == nil {
        request = NewCreateHostRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "CreateHost")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateHost require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateHostResponse()
    err = c.Send(request, response)
    return
}

func NewCreateIpAccessControlRequest() (request *CreateIpAccessControlRequest) {
    request = &CreateIpAccessControlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "CreateIpAccessControl")
    
    
    return
}

func NewCreateIpAccessControlResponse() (response *CreateIpAccessControlResponse) {
    response = &CreateIpAccessControlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateIpAccessControl
// This API is used to add WAF IP allowlists/blocklists.
//
// error code that may be returned:
//  FAILEDOPERATION_THENUMBEROFADDEDBLACKANDWHITELISTEXCEEDSTHEUPPERLIMIT = "FailedOperation.TheNumberOfAddedBlackAndWhiteListExceedsTheUpperLimit"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_SPECIFICATIONERR = "LimitExceeded.SpecificationErr"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateIpAccessControl(request *CreateIpAccessControlRequest) (response *CreateIpAccessControlResponse, err error) {
    return c.CreateIpAccessControlWithContext(context.Background(), request)
}

// CreateIpAccessControl
// This API is used to add WAF IP allowlists/blocklists.
//
// error code that may be returned:
//  FAILEDOPERATION_THENUMBEROFADDEDBLACKANDWHITELISTEXCEEDSTHEUPPERLIMIT = "FailedOperation.TheNumberOfAddedBlackAndWhiteListExceedsTheUpperLimit"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_SPECIFICATIONERR = "LimitExceeded.SpecificationErr"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateIpAccessControlWithContext(ctx context.Context, request *CreateIpAccessControlRequest) (response *CreateIpAccessControlResponse, err error) {
    if request == nil {
        request = NewCreateIpAccessControlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "CreateIpAccessControl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateIpAccessControl require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateIpAccessControlResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAntiFakeUrlRequest() (request *DeleteAntiFakeUrlRequest) {
    request = &DeleteAntiFakeUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DeleteAntiFakeUrl")
    
    
    return
}

func NewDeleteAntiFakeUrlResponse() (response *DeleteAntiFakeUrlResponse) {
    response = &DeleteAntiFakeUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAntiFakeUrl
// Delete tamper-proof URL
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteAntiFakeUrl(request *DeleteAntiFakeUrlRequest) (response *DeleteAntiFakeUrlResponse, err error) {
    return c.DeleteAntiFakeUrlWithContext(context.Background(), request)
}

// DeleteAntiFakeUrl
// Delete tamper-proof URL
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteAntiFakeUrlWithContext(ctx context.Context, request *DeleteAntiFakeUrlRequest) (response *DeleteAntiFakeUrlResponse, err error) {
    if request == nil {
        request = NewDeleteAntiFakeUrlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DeleteAntiFakeUrl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAntiFakeUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAntiFakeUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAntiInfoLeakRuleRequest() (request *DeleteAntiInfoLeakRuleRequest) {
    request = &DeleteAntiInfoLeakRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DeleteAntiInfoLeakRule")
    
    
    return
}

func NewDeleteAntiInfoLeakRuleResponse() (response *DeleteAntiInfoLeakRuleResponse) {
    response = &DeleteAntiInfoLeakRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAntiInfoLeakRule
// Delete information leakage prevention rule
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
func (c *Client) DeleteAntiInfoLeakRule(request *DeleteAntiInfoLeakRuleRequest) (response *DeleteAntiInfoLeakRuleResponse, err error) {
    return c.DeleteAntiInfoLeakRuleWithContext(context.Background(), request)
}

// DeleteAntiInfoLeakRule
// Delete information leakage prevention rule
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
func (c *Client) DeleteAntiInfoLeakRuleWithContext(ctx context.Context, request *DeleteAntiInfoLeakRuleRequest) (response *DeleteAntiInfoLeakRuleResponse, err error) {
    if request == nil {
        request = NewDeleteAntiInfoLeakRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DeleteAntiInfoLeakRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAntiInfoLeakRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAntiInfoLeakRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCCRuleRequest() (request *DeleteCCRuleRequest) {
    request = &DeleteCCRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DeleteCCRule")
    
    
    return
}

func NewDeleteCCRuleResponse() (response *DeleteCCRuleResponse) {
    response = &DeleteCCRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCCRule
// WAF CC V2 deletion API
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCCRule(request *DeleteCCRuleRequest) (response *DeleteCCRuleResponse, err error) {
    return c.DeleteCCRuleWithContext(context.Background(), request)
}

// DeleteCCRule
// WAF CC V2 deletion API
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCCRuleWithContext(ctx context.Context, request *DeleteCCRuleRequest) (response *DeleteCCRuleResponse, err error) {
    if request == nil {
        request = NewDeleteCCRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DeleteCCRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCCRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCCRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCustomRuleRequest() (request *DeleteCustomRuleRequest) {
    request = &DeleteCustomRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DeleteCustomRule")
    
    
    return
}

func NewDeleteCustomRuleResponse() (response *DeleteCustomRuleResponse) {
    response = &DeleteCustomRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCustomRule
// Delete custom rule
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCustomRule(request *DeleteCustomRuleRequest) (response *DeleteCustomRuleResponse, err error) {
    return c.DeleteCustomRuleWithContext(context.Background(), request)
}

// DeleteCustomRule
// Delete custom rule
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCustomRuleWithContext(ctx context.Context, request *DeleteCustomRuleRequest) (response *DeleteCustomRuleResponse, err error) {
    if request == nil {
        request = NewDeleteCustomRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DeleteCustomRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCustomRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCustomRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCustomWhiteRuleRequest() (request *DeleteCustomWhiteRuleRequest) {
    request = &DeleteCustomWhiteRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DeleteCustomWhiteRule")
    
    
    return
}

func NewDeleteCustomWhiteRuleResponse() (response *DeleteCustomWhiteRuleResponse) {
    response = &DeleteCustomWhiteRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCustomWhiteRule
// Delete precision allowlist rules
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCustomWhiteRule(request *DeleteCustomWhiteRuleRequest) (response *DeleteCustomWhiteRuleResponse, err error) {
    return c.DeleteCustomWhiteRuleWithContext(context.Background(), request)
}

// DeleteCustomWhiteRule
// Delete precision allowlist rules
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCustomWhiteRuleWithContext(ctx context.Context, request *DeleteCustomWhiteRuleRequest) (response *DeleteCustomWhiteRuleResponse, err error) {
    if request == nil {
        request = NewDeleteCustomWhiteRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DeleteCustomWhiteRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCustomWhiteRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCustomWhiteRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteHostRequest() (request *DeleteHostRequest) {
    request = &DeleteHostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DeleteHost")
    
    
    return
}

func NewDeleteHostResponse() (response *DeleteHostResponse) {
    response = &DeleteHostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteHost
// This API is used to delete a domain name protected by CLB WAF. Batch operation is supported.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteHost(request *DeleteHostRequest) (response *DeleteHostResponse, err error) {
    return c.DeleteHostWithContext(context.Background(), request)
}

// DeleteHost
// This API is used to delete a domain name protected by CLB WAF. Batch operation is supported.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteHostWithContext(ctx context.Context, request *DeleteHostRequest) (response *DeleteHostResponse, err error) {
    if request == nil {
        request = NewDeleteHostRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DeleteHost")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteHost require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteHostResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteIpAccessControlV2Request() (request *DeleteIpAccessControlV2Request) {
    request = &DeleteIpAccessControlV2Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DeleteIpAccessControlV2")
    
    
    return
}

func NewDeleteIpAccessControlV2Response() (response *DeleteIpAccessControlV2Response) {
    response = &DeleteIpAccessControlV2Response{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteIpAccessControlV2
// This API is used to delete latest versions of WAF IP allowlists/blocklists.
//
// error code that may be returned:
//  FAILEDOPERATION_THENUMBEROFONETIMEDELETIONSREACHEDTHEUPPERLIMIT = "FailedOperation.TheNumberOfOneTimeDeletionsReachedTheUpperLimit"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteIpAccessControlV2(request *DeleteIpAccessControlV2Request) (response *DeleteIpAccessControlV2Response, err error) {
    return c.DeleteIpAccessControlV2WithContext(context.Background(), request)
}

// DeleteIpAccessControlV2
// This API is used to delete latest versions of WAF IP allowlists/blocklists.
//
// error code that may be returned:
//  FAILEDOPERATION_THENUMBEROFONETIMEDELETIONSREACHEDTHEUPPERLIMIT = "FailedOperation.TheNumberOfOneTimeDeletionsReachedTheUpperLimit"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteIpAccessControlV2WithContext(ctx context.Context, request *DeleteIpAccessControlV2Request) (response *DeleteIpAccessControlV2Response, err error) {
    if request == nil {
        request = NewDeleteIpAccessControlV2Request()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DeleteIpAccessControlV2")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteIpAccessControlV2 require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteIpAccessControlV2Response()
    err = c.Send(request, response)
    return
}

func NewDeleteSessionRequest() (request *DeleteSessionRequest) {
    request = &DeleteSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DeleteSession")
    
    
    return
}

func NewDeleteSessionResponse() (response *DeleteSessionResponse) {
    response = &DeleteSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSession
// Delete CC attack session settings
//
// error code that may be returned:
//  FAILEDOPERATION_SESSIONINUSED = "FailedOperation.SessionInUsed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteSession(request *DeleteSessionRequest) (response *DeleteSessionResponse, err error) {
    return c.DeleteSessionWithContext(context.Background(), request)
}

// DeleteSession
// Delete CC attack session settings
//
// error code that may be returned:
//  FAILEDOPERATION_SESSIONINUSED = "FailedOperation.SessionInUsed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteSessionWithContext(ctx context.Context, request *DeleteSessionRequest) (response *DeleteSessionResponse, err error) {
    if request == nil {
        request = NewDeleteSessionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DeleteSession")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSession require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSessionResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSpartaProtectionRequest() (request *DeleteSpartaProtectionRequest) {
    request = &DeleteSpartaProtectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DeleteSpartaProtection")
    
    
    return
}

func NewDeleteSpartaProtectionResponse() (response *DeleteSpartaProtectionResponse) {
    response = &DeleteSpartaProtectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSpartaProtection
// This API is used to delete a domain name protected by SaaS WAF.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ASYNCHRONOUSCALLFAILED = "InternalError.AsynchronousCallFailed"
//  INVALIDPARAMETER_UNAUTHORIZEDOPERATIONPARAMETERERR = "InvalidParameter.UnauthorizedOperationParameterErr"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteSpartaProtection(request *DeleteSpartaProtectionRequest) (response *DeleteSpartaProtectionResponse, err error) {
    return c.DeleteSpartaProtectionWithContext(context.Background(), request)
}

// DeleteSpartaProtection
// This API is used to delete a domain name protected by SaaS WAF.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ASYNCHRONOUSCALLFAILED = "InternalError.AsynchronousCallFailed"
//  INVALIDPARAMETER_UNAUTHORIZEDOPERATIONPARAMETERERR = "InvalidParameter.UnauthorizedOperationParameterErr"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteSpartaProtectionWithContext(ctx context.Context, request *DeleteSpartaProtectionRequest) (response *DeleteSpartaProtectionResponse, err error) {
    if request == nil {
        request = NewDeleteSpartaProtectionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DeleteSpartaProtection")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSpartaProtection require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSpartaProtectionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAntiFakeRulesRequest() (request *DescribeAntiFakeRulesRequest) {
    request = &DescribeAntiFakeRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeAntiFakeRules")
    
    
    return
}

func NewDescribeAntiFakeRulesResponse() (response *DescribeAntiFakeRulesResponse) {
    response = &DescribeAntiFakeRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAntiFakeRules
// Obtain a tamper-proof URL
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAntiFakeRules(request *DescribeAntiFakeRulesRequest) (response *DescribeAntiFakeRulesResponse, err error) {
    return c.DescribeAntiFakeRulesWithContext(context.Background(), request)
}

// DescribeAntiFakeRules
// Obtain a tamper-proof URL
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAntiFakeRulesWithContext(ctx context.Context, request *DescribeAntiFakeRulesRequest) (response *DescribeAntiFakeRulesResponse, err error) {
    if request == nil {
        request = NewDescribeAntiFakeRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DescribeAntiFakeRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAntiFakeRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAntiFakeRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAntiInfoLeakageRulesRequest() (request *DescribeAntiInfoLeakageRulesRequest) {
    request = &DescribeAntiInfoLeakageRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeAntiInfoLeakageRules")
    
    
    return
}

func NewDescribeAntiInfoLeakageRulesResponse() (response *DescribeAntiInfoLeakageRulesResponse) {
    response = &DescribeAntiInfoLeakageRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAntiInfoLeakageRules
// Obtain the information leakage prevention rule list
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
func (c *Client) DescribeAntiInfoLeakageRules(request *DescribeAntiInfoLeakageRulesRequest) (response *DescribeAntiInfoLeakageRulesResponse, err error) {
    return c.DescribeAntiInfoLeakageRulesWithContext(context.Background(), request)
}

// DescribeAntiInfoLeakageRules
// Obtain the information leakage prevention rule list
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
func (c *Client) DescribeAntiInfoLeakageRulesWithContext(ctx context.Context, request *DescribeAntiInfoLeakageRulesRequest) (response *DescribeAntiInfoLeakageRulesResponse, err error) {
    if request == nil {
        request = NewDescribeAntiInfoLeakageRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DescribeAntiInfoLeakageRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAntiInfoLeakageRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAntiInfoLeakageRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAttackOverviewRequest() (request *DescribeAttackOverviewRequest) {
    request = &DescribeAttackOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeAttackOverview")
    
    
    return
}

func NewDescribeAttackOverviewResponse() (response *DescribeAttackOverviewResponse) {
    response = &DescribeAttackOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAttackOverview
// This API is used to describe the attack overview.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAttackOverview(request *DescribeAttackOverviewRequest) (response *DescribeAttackOverviewResponse, err error) {
    return c.DescribeAttackOverviewWithContext(context.Background(), request)
}

// DescribeAttackOverview
// This API is used to describe the attack overview.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAttackOverviewWithContext(ctx context.Context, request *DescribeAttackOverviewRequest) (response *DescribeAttackOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeAttackOverviewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DescribeAttackOverview")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAttackOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAttackOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAttackTypeRequest() (request *DescribeAttackTypeRequest) {
    request = &DescribeAttackTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeAttackType")
    
    
    return
}

func NewDescribeAttackTypeResponse() (response *DescribeAttackTypeResponse) {
    response = &DescribeAttackTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAttackType
// Query the top N attack types for a specified domain
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAttackType(request *DescribeAttackTypeRequest) (response *DescribeAttackTypeResponse, err error) {
    return c.DescribeAttackTypeWithContext(context.Background(), request)
}

// DescribeAttackType
// Query the top N attack types for a specified domain
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAttackTypeWithContext(ctx context.Context, request *DescribeAttackTypeRequest) (response *DescribeAttackTypeResponse, err error) {
    if request == nil {
        request = NewDescribeAttackTypeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DescribeAttackType")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAttackType require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAttackTypeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBatchIpAccessControlRequest() (request *DescribeBatchIpAccessControlRequest) {
    request = &DescribeBatchIpAccessControlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeBatchIpAccessControl")
    
    
    return
}

func NewDescribeBatchIpAccessControlResponse() (response *DescribeBatchIpAccessControlResponse) {
    response = &DescribeBatchIpAccessControlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBatchIpAccessControl
// This API is used to query the IP blocklist and allowlist for WAF batch protection.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
func (c *Client) DescribeBatchIpAccessControl(request *DescribeBatchIpAccessControlRequest) (response *DescribeBatchIpAccessControlResponse, err error) {
    return c.DescribeBatchIpAccessControlWithContext(context.Background(), request)
}

// DescribeBatchIpAccessControl
// This API is used to query the IP blocklist and allowlist for WAF batch protection.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
func (c *Client) DescribeBatchIpAccessControlWithContext(ctx context.Context, request *DescribeBatchIpAccessControlRequest) (response *DescribeBatchIpAccessControlResponse, err error) {
    if request == nil {
        request = NewDescribeBatchIpAccessControlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DescribeBatchIpAccessControl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBatchIpAccessControl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBatchIpAccessControlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCCRuleRequest() (request *DescribeCCRuleRequest) {
    request = &DescribeCCRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeCCRule")
    
    
    return
}

func NewDescribeCCRuleResponse() (response *DescribeCCRuleResponse) {
    response = &DescribeCCRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCCRule
// WAF CC V2 query API
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeCCRule(request *DescribeCCRuleRequest) (response *DescribeCCRuleResponse, err error) {
    return c.DescribeCCRuleWithContext(context.Background(), request)
}

// DescribeCCRule
// WAF CC V2 query API
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeCCRuleWithContext(ctx context.Context, request *DescribeCCRuleRequest) (response *DescribeCCRuleResponse, err error) {
    if request == nil {
        request = NewDescribeCCRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DescribeCCRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCCRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCCRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCCRuleListRequest() (request *DescribeCCRuleListRequest) {
    request = &DescribeCCRuleListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeCCRuleList")
    
    
    return
}

func NewDescribeCCRuleListResponse() (response *DescribeCCRuleListResponse) {
    response = &DescribeCCRuleListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCCRuleList
// Query CC rules based on multiple conditions
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCCRuleList(request *DescribeCCRuleListRequest) (response *DescribeCCRuleListResponse, err error) {
    return c.DescribeCCRuleListWithContext(context.Background(), request)
}

// DescribeCCRuleList
// Query CC rules based on multiple conditions
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCCRuleListWithContext(ctx context.Context, request *DescribeCCRuleListRequest) (response *DescribeCCRuleListResponse, err error) {
    if request == nil {
        request = NewDescribeCCRuleListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DescribeCCRuleList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCCRuleList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCCRuleListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCertificateVerifyResultRequest() (request *DescribeCertificateVerifyResultRequest) {
    request = &DescribeCertificateVerifyResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeCertificateVerifyResult")
    
    
    return
}

func NewDescribeCertificateVerifyResultResponse() (response *DescribeCertificateVerifyResultResponse) {
    response = &DescribeCertificateVerifyResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCertificateVerifyResult
// Obtain certificate inspection result
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeCertificateVerifyResult(request *DescribeCertificateVerifyResultRequest) (response *DescribeCertificateVerifyResultResponse, err error) {
    return c.DescribeCertificateVerifyResultWithContext(context.Background(), request)
}

// DescribeCertificateVerifyResult
// Obtain certificate inspection result
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeCertificateVerifyResultWithContext(ctx context.Context, request *DescribeCertificateVerifyResultRequest) (response *DescribeCertificateVerifyResultResponse, err error) {
    if request == nil {
        request = NewDescribeCertificateVerifyResultRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DescribeCertificateVerifyResult")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCertificateVerifyResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCertificateVerifyResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCiphersDetailRequest() (request *DescribeCiphersDetailRequest) {
    request = &DescribeCiphersDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeCiphersDetail")
    
    
    return
}

func NewDescribeCiphersDetailResponse() (response *DescribeCiphersDetailResponse) {
    response = &DescribeCiphersDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCiphersDetail
// Query SaaS WAF cipher suite information
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCiphersDetail(request *DescribeCiphersDetailRequest) (response *DescribeCiphersDetailResponse, err error) {
    return c.DescribeCiphersDetailWithContext(context.Background(), request)
}

// DescribeCiphersDetail
// Query SaaS WAF cipher suite information
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCiphersDetailWithContext(ctx context.Context, request *DescribeCiphersDetailRequest) (response *DescribeCiphersDetailResponse, err error) {
    if request == nil {
        request = NewDescribeCiphersDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DescribeCiphersDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCiphersDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCiphersDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomRuleListRequest() (request *DescribeCustomRuleListRequest) {
    request = &DescribeCustomRuleListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeCustomRuleList")
    
    
    return
}

func NewDescribeCustomRuleListResponse() (response *DescribeCustomRuleListResponse) {
    response = &DescribeCustomRuleListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCustomRuleList
// Obtain the access control policy list in the protection configuration
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCustomRuleList(request *DescribeCustomRuleListRequest) (response *DescribeCustomRuleListResponse, err error) {
    return c.DescribeCustomRuleListWithContext(context.Background(), request)
}

// DescribeCustomRuleList
// Obtain the access control policy list in the protection configuration
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCustomRuleListWithContext(ctx context.Context, request *DescribeCustomRuleListRequest) (response *DescribeCustomRuleListResponse, err error) {
    if request == nil {
        request = NewDescribeCustomRuleListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DescribeCustomRuleList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCustomRuleList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCustomRuleListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomWhiteRuleRequest() (request *DescribeCustomWhiteRuleRequest) {
    request = &DescribeCustomWhiteRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeCustomWhiteRule")
    
    
    return
}

func NewDescribeCustomWhiteRuleResponse() (response *DescribeCustomWhiteRuleResponse) {
    response = &DescribeCustomWhiteRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCustomWhiteRule
// Obtain the precision allowlist policy list in the protection configuration
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCustomWhiteRule(request *DescribeCustomWhiteRuleRequest) (response *DescribeCustomWhiteRuleResponse, err error) {
    return c.DescribeCustomWhiteRuleWithContext(context.Background(), request)
}

// DescribeCustomWhiteRule
// Obtain the precision allowlist policy list in the protection configuration
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCustomWhiteRuleWithContext(ctx context.Context, request *DescribeCustomWhiteRuleRequest) (response *DescribeCustomWhiteRuleResponse, err error) {
    if request == nil {
        request = NewDescribeCustomWhiteRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DescribeCustomWhiteRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCustomWhiteRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCustomWhiteRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainCountInfoRequest() (request *DescribeDomainCountInfoRequest) {
    request = &DescribeDomainCountInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeDomainCountInfo")
    
    
    return
}

func NewDescribeDomainCountInfoResponse() (response *DescribeDomainCountInfoResponse) {
    response = &DescribeDomainCountInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDomainCountInfo
// Obtain domain overview
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDomainCountInfo(request *DescribeDomainCountInfoRequest) (response *DescribeDomainCountInfoResponse, err error) {
    return c.DescribeDomainCountInfoWithContext(context.Background(), request)
}

// DescribeDomainCountInfo
// Obtain domain overview
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDomainCountInfoWithContext(ctx context.Context, request *DescribeDomainCountInfoRequest) (response *DescribeDomainCountInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDomainCountInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DescribeDomainCountInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomainCountInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainCountInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainDetailsClbRequest() (request *DescribeDomainDetailsClbRequest) {
    request = &DescribeDomainDetailsClbRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeDomainDetailsClb")
    
    
    return
}

func NewDescribeDomainDetailsClbResponse() (response *DescribeDomainDetailsClbResponse) {
    response = &DescribeDomainDetailsClbResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDomainDetailsClb
// Obtain CLB WAF domain details
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDomainDetailsClb(request *DescribeDomainDetailsClbRequest) (response *DescribeDomainDetailsClbResponse, err error) {
    return c.DescribeDomainDetailsClbWithContext(context.Background(), request)
}

// DescribeDomainDetailsClb
// Obtain CLB WAF domain details
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDomainDetailsClbWithContext(ctx context.Context, request *DescribeDomainDetailsClbRequest) (response *DescribeDomainDetailsClbResponse, err error) {
    if request == nil {
        request = NewDescribeDomainDetailsClbRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DescribeDomainDetailsClb")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomainDetailsClb require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainDetailsClbResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainDetailsSaasRequest() (request *DescribeDomainDetailsSaasRequest) {
    request = &DescribeDomainDetailsSaasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeDomainDetailsSaas")
    
    
    return
}

func NewDescribeDomainDetailsSaasResponse() (response *DescribeDomainDetailsSaasResponse) {
    response = &DescribeDomainDetailsSaasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDomainDetailsSaas
// Query individual SaaS WAF domain details
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDomainDetailsSaas(request *DescribeDomainDetailsSaasRequest) (response *DescribeDomainDetailsSaasResponse, err error) {
    return c.DescribeDomainDetailsSaasWithContext(context.Background(), request)
}

// DescribeDomainDetailsSaas
// Query individual SaaS WAF domain details
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDomainDetailsSaasWithContext(ctx context.Context, request *DescribeDomainDetailsSaasRequest) (response *DescribeDomainDetailsSaasResponse, err error) {
    if request == nil {
        request = NewDescribeDomainDetailsSaasRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DescribeDomainDetailsSaas")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomainDetailsSaas require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainDetailsSaasResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainVerifyResultRequest() (request *DescribeDomainVerifyResultRequest) {
    request = &DescribeDomainVerifyResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeDomainVerifyResult")
    
    
    return
}

func NewDescribeDomainVerifyResultResponse() (response *DescribeDomainVerifyResultResponse) {
    response = &DescribeDomainVerifyResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDomainVerifyResult
// Obtain the result of adding domain operation
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeDomainVerifyResult(request *DescribeDomainVerifyResultRequest) (response *DescribeDomainVerifyResultResponse, err error) {
    return c.DescribeDomainVerifyResultWithContext(context.Background(), request)
}

// DescribeDomainVerifyResult
// Obtain the result of adding domain operation
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeDomainVerifyResultWithContext(ctx context.Context, request *DescribeDomainVerifyResultRequest) (response *DescribeDomainVerifyResultResponse, err error) {
    if request == nil {
        request = NewDescribeDomainVerifyResultRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DescribeDomainVerifyResult")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomainVerifyResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainVerifyResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainsRequest() (request *DescribeDomainsRequest) {
    request = &DescribeDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeDomains")
    
    
    return
}

func NewDescribeDomainsResponse() (response *DescribeDomainsResponse) {
    response = &DescribeDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDomains
// Query detailed information of all user domains
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDomains(request *DescribeDomainsRequest) (response *DescribeDomainsResponse, err error) {
    return c.DescribeDomainsWithContext(context.Background(), request)
}

// DescribeDomains
// Query detailed information of all user domains
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDomainsWithContext(ctx context.Context, request *DescribeDomainsRequest) (response *DescribeDomainsResponse, err error) {
    if request == nil {
        request = NewDescribeDomainsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DescribeDomains")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFindDomainListRequest() (request *DescribeFindDomainListRequest) {
    request = &DescribeFindDomainListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeFindDomainList")
    
    
    return
}

func NewDescribeFindDomainListResponse() (response *DescribeFindDomainListResponse) {
    response = &DescribeFindDomainListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFindDomainList
// Obtain discovered domain name list API
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeFindDomainList(request *DescribeFindDomainListRequest) (response *DescribeFindDomainListResponse, err error) {
    return c.DescribeFindDomainListWithContext(context.Background(), request)
}

// DescribeFindDomainList
// Obtain discovered domain name list API
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeFindDomainListWithContext(ctx context.Context, request *DescribeFindDomainListRequest) (response *DescribeFindDomainListResponse, err error) {
    if request == nil {
        request = NewDescribeFindDomainListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DescribeFindDomainList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFindDomainList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFindDomainListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHistogramRequest() (request *DescribeHistogramRequest) {
    request = &DescribeHistogramRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeHistogram")
    
    
    return
}

func NewDescribeHistogramResponse() (response *DescribeHistogramResponse) {
    response = &DescribeHistogramResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHistogram
// Query various conditions of cluster analysis
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeHistogram(request *DescribeHistogramRequest) (response *DescribeHistogramResponse, err error) {
    return c.DescribeHistogramWithContext(context.Background(), request)
}

// DescribeHistogram
// Query various conditions of cluster analysis
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeHistogramWithContext(ctx context.Context, request *DescribeHistogramRequest) (response *DescribeHistogramResponse, err error) {
    if request == nil {
        request = NewDescribeHistogramRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DescribeHistogram")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHistogram require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHistogramResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostRequest() (request *DescribeHostRequest) {
    request = &DescribeHostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeHost")
    
    
    return
}

func NewDescribeHostResponse() (response *DescribeHostResponse) {
    response = &DescribeHostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHost
// Obtain protection domain details in CLB-WAF
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeHost(request *DescribeHostRequest) (response *DescribeHostResponse, err error) {
    return c.DescribeHostWithContext(context.Background(), request)
}

// DescribeHost
// Obtain protection domain details in CLB-WAF
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeHostWithContext(ctx context.Context, request *DescribeHostRequest) (response *DescribeHostResponse, err error) {
    if request == nil {
        request = NewDescribeHostRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DescribeHost")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHost require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHostResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostLimitRequest() (request *DescribeHostLimitRequest) {
    request = &DescribeHostLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeHostLimit")
    
    
    return
}

func NewDescribeHostLimitResponse() (response *DescribeHostLimitResponse) {
    response = &DescribeHostLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHostLimit
// Firstly verify when adding a domain whether a package has been purchased, whether the limit of purchased package has not been reached, and whether the domain has already been added
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeHostLimit(request *DescribeHostLimitRequest) (response *DescribeHostLimitResponse, err error) {
    return c.DescribeHostLimitWithContext(context.Background(), request)
}

// DescribeHostLimit
// Firstly verify when adding a domain whether a package has been purchased, whether the limit of purchased package has not been reached, and whether the domain has already been added
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeHostLimitWithContext(ctx context.Context, request *DescribeHostLimitRequest) (response *DescribeHostLimitResponse, err error) {
    if request == nil {
        request = NewDescribeHostLimitRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DescribeHostLimit")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHostLimit require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHostLimitResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostsRequest() (request *DescribeHostsRequest) {
    request = &DescribeHostsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeHosts")
    
    
    return
}

func NewDescribeHostsResponse() (response *DescribeHostsResponse) {
    response = &DescribeHostsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHosts
// Obtain protection domain list in CLB-WAF
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeHosts(request *DescribeHostsRequest) (response *DescribeHostsResponse, err error) {
    return c.DescribeHostsWithContext(context.Background(), request)
}

// DescribeHosts
// Obtain protection domain list in CLB-WAF
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeHostsWithContext(ctx context.Context, request *DescribeHostsRequest) (response *DescribeHostsResponse, err error) {
    if request == nil {
        request = NewDescribeHostsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DescribeHosts")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHosts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHostsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeInstances")
    
    
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstances
// Query detailed information of all user instances
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    return c.DescribeInstancesWithContext(context.Background(), request)
}

// DescribeInstances
// Query detailed information of all user instances
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInstancesWithContext(ctx context.Context, request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DescribeInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIpAccessControlRequest() (request *DescribeIpAccessControlRequest) {
    request = &DescribeIpAccessControlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeIpAccessControl")
    
    
    return
}

func NewDescribeIpAccessControlResponse() (response *DescribeIpAccessControlResponse) {
    response = &DescribeIpAccessControlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIpAccessControl
// WAF IP blocklist/allowlist query
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
func (c *Client) DescribeIpAccessControl(request *DescribeIpAccessControlRequest) (response *DescribeIpAccessControlResponse, err error) {
    return c.DescribeIpAccessControlWithContext(context.Background(), request)
}

// DescribeIpAccessControl
// WAF IP blocklist/allowlist query
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
func (c *Client) DescribeIpAccessControlWithContext(ctx context.Context, request *DescribeIpAccessControlRequest) (response *DescribeIpAccessControlResponse, err error) {
    if request == nil {
        request = NewDescribeIpAccessControlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DescribeIpAccessControl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIpAccessControl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIpAccessControlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModuleStatusRequest() (request *DescribeModuleStatusRequest) {
    request = &DescribeModuleStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeModuleStatus")
    
    
    return
}

func NewDescribeModuleStatusResponse() (response *DescribeModuleStatusResponse) {
    response = &DescribeModuleStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeModuleStatus
// Query the switch status of each WAF basic security module, check if each module is enabled
//
// error code that may be returned:
//  FAILEDOPERATION_MONGOOPERATIONFAILED = "FailedOperation.MongoOperationFailed"
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeModuleStatus(request *DescribeModuleStatusRequest) (response *DescribeModuleStatusResponse, err error) {
    return c.DescribeModuleStatusWithContext(context.Background(), request)
}

// DescribeModuleStatus
// Query the switch status of each WAF basic security module, check if each module is enabled
//
// error code that may be returned:
//  FAILEDOPERATION_MONGOOPERATIONFAILED = "FailedOperation.MongoOperationFailed"
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeModuleStatusWithContext(ctx context.Context, request *DescribeModuleStatusRequest) (response *DescribeModuleStatusResponse, err error) {
    if request == nil {
        request = NewDescribeModuleStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DescribeModuleStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeModuleStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeModuleStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeObjectsRequest() (request *DescribeObjectsRequest) {
    request = &DescribeObjectsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeObjects")
    
    
    return
}

func NewDescribeObjectsResponse() (response *DescribeObjectsResponse) {
    response = &DescribeObjectsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeObjects
// View protected object list
//
// error code that may be returned:
//  FAILEDOPERATION_MONGOOPERATIONFAILED = "FailedOperation.MongoOperationFailed"
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeObjects(request *DescribeObjectsRequest) (response *DescribeObjectsResponse, err error) {
    return c.DescribeObjectsWithContext(context.Background(), request)
}

// DescribeObjects
// View protected object list
//
// error code that may be returned:
//  FAILEDOPERATION_MONGOOPERATIONFAILED = "FailedOperation.MongoOperationFailed"
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeObjectsWithContext(ctx context.Context, request *DescribeObjectsRequest) (response *DescribeObjectsResponse, err error) {
    if request == nil {
        request = NewDescribeObjectsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DescribeObjects")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeObjects require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeObjectsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePeakPointsRequest() (request *DescribePeakPointsRequest) {
    request = &DescribePeakPointsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribePeakPoints")
    
    
    return
}

func NewDescribePeakPointsResponse() (response *DescribePeakPointsResponse) {
    response = &DescribePeakPointsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePeakPoints
// Query business and attack summary trends
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePeakPoints(request *DescribePeakPointsRequest) (response *DescribePeakPointsResponse, err error) {
    return c.DescribePeakPointsWithContext(context.Background(), request)
}

// DescribePeakPoints
// Query business and attack summary trends
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePeakPointsWithContext(ctx context.Context, request *DescribePeakPointsRequest) (response *DescribePeakPointsResponse, err error) {
    if request == nil {
        request = NewDescribePeakPointsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DescribePeakPoints")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePeakPoints require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePeakPointsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePolicyStatusRequest() (request *DescribePolicyStatusRequest) {
    request = &DescribePolicyStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribePolicyStatus")
    
    
    return
}

func NewDescribePolicyStatusResponse() (response *DescribePolicyStatusResponse) {
    response = &DescribePolicyStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePolicyStatus
// Obtain protection status and the effective instance ID
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePolicyStatus(request *DescribePolicyStatusRequest) (response *DescribePolicyStatusResponse, err error) {
    return c.DescribePolicyStatusWithContext(context.Background(), request)
}

// DescribePolicyStatus
// Obtain protection status and the effective instance ID
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePolicyStatusWithContext(ctx context.Context, request *DescribePolicyStatusRequest) (response *DescribePolicyStatusResponse, err error) {
    if request == nil {
        request = NewDescribePolicyStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DescribePolicyStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePolicyStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePolicyStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePortsRequest() (request *DescribePortsRequest) {
    request = &DescribePortsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribePorts")
    
    
    return
}

func NewDescribePortsResponse() (response *DescribePortsResponse) {
    response = &DescribePortsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePorts
// Obtain the SaaS-type WAF protection port list
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribePorts(request *DescribePortsRequest) (response *DescribePortsResponse, err error) {
    return c.DescribePortsWithContext(context.Background(), request)
}

// DescribePorts
// Obtain the SaaS-type WAF protection port list
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribePortsWithContext(ctx context.Context, request *DescribePortsRequest) (response *DescribePortsResponse, err error) {
    if request == nil {
        request = NewDescribePortsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DescribePorts")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePorts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePortsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRuleLimitRequest() (request *DescribeRuleLimitRequest) {
    request = &DescribeRuleLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeRuleLimit")
    
    
    return
}

func NewDescribeRuleLimitResponse() (response *DescribeRuleLimitResponse) {
    response = &DescribeRuleLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRuleLimit
// Obtain specific specification limits for each module
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRuleLimit(request *DescribeRuleLimitRequest) (response *DescribeRuleLimitResponse, err error) {
    return c.DescribeRuleLimitWithContext(context.Background(), request)
}

// DescribeRuleLimit
// Obtain specific specification limits for each module
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRuleLimitWithContext(ctx context.Context, request *DescribeRuleLimitRequest) (response *DescribeRuleLimitResponse, err error) {
    if request == nil {
        request = NewDescribeRuleLimitRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DescribeRuleLimit")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRuleLimit require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRuleLimitResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSessionRequest() (request *DescribeSessionRequest) {
    request = &DescribeSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeSession")
    
    
    return
}

func NewDescribeSessionResponse() (response *DescribeSessionResponse) {
    response = &DescribeSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSession
// WAF session definition query API
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSession(request *DescribeSessionRequest) (response *DescribeSessionResponse, err error) {
    return c.DescribeSessionWithContext(context.Background(), request)
}

// DescribeSession
// WAF session definition query API
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSessionWithContext(ctx context.Context, request *DescribeSessionRequest) (response *DescribeSessionResponse, err error) {
    if request == nil {
        request = NewDescribeSessionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DescribeSession")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSession require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSessionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSpartaProtectionInfoRequest() (request *DescribeSpartaProtectionInfoRequest) {
    request = &DescribeSpartaProtectionInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeSpartaProtectionInfo")
    
    
    return
}

func NewDescribeSpartaProtectionInfoResponse() (response *DescribeSpartaProtectionInfoResponse) {
    response = &DescribeSpartaProtectionInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSpartaProtectionInfo
// WAF Sparta - Obtain protection domain information
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeSpartaProtectionInfo(request *DescribeSpartaProtectionInfoRequest) (response *DescribeSpartaProtectionInfoResponse, err error) {
    return c.DescribeSpartaProtectionInfoWithContext(context.Background(), request)
}

// DescribeSpartaProtectionInfo
// WAF Sparta - Obtain protection domain information
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeSpartaProtectionInfoWithContext(ctx context.Context, request *DescribeSpartaProtectionInfoRequest) (response *DescribeSpartaProtectionInfoResponse, err error) {
    if request == nil {
        request = NewDescribeSpartaProtectionInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DescribeSpartaProtectionInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSpartaProtectionInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSpartaProtectionInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTlsVersionRequest() (request *DescribeTlsVersionRequest) {
    request = &DescribeTlsVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeTlsVersion")
    
    
    return
}

func NewDescribeTlsVersionResponse() (response *DescribeTlsVersionResponse) {
    response = &DescribeTlsVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTlsVersion
// This API is used to query TLS versions supported by SaaS WAF.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTlsVersion(request *DescribeTlsVersionRequest) (response *DescribeTlsVersionResponse, err error) {
    return c.DescribeTlsVersionWithContext(context.Background(), request)
}

// DescribeTlsVersion
// This API is used to query TLS versions supported by SaaS WAF.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTlsVersionWithContext(ctx context.Context, request *DescribeTlsVersionRequest) (response *DescribeTlsVersionResponse, err error) {
    if request == nil {
        request = NewDescribeTlsVersionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DescribeTlsVersion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTlsVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTlsVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopAttackDomainRequest() (request *DescribeTopAttackDomainRequest) {
    request = &DescribeTopAttackDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeTopAttackDomain")
    
    
    return
}

func NewDescribeTopAttackDomainResponse() (response *DescribeTopAttackDomainResponse) {
    response = &DescribeTopAttackDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTopAttackDomain
// Query top 5 attack domains
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeTopAttackDomain(request *DescribeTopAttackDomainRequest) (response *DescribeTopAttackDomainResponse, err error) {
    return c.DescribeTopAttackDomainWithContext(context.Background(), request)
}

// DescribeTopAttackDomain
// Query top 5 attack domains
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeTopAttackDomainWithContext(ctx context.Context, request *DescribeTopAttackDomainRequest) (response *DescribeTopAttackDomainResponse, err error) {
    if request == nil {
        request = NewDescribeTopAttackDomainRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DescribeTopAttackDomain")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopAttackDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopAttackDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserClbWafRegionsRequest() (request *DescribeUserClbWafRegionsRequest) {
    request = &DescribeUserClbWafRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeUserClbWafRegions")
    
    
    return
}

func NewDescribeUserClbWafRegionsResponse() (response *DescribeUserClbWafRegionsResponse) {
    response = &DescribeUserClbWafRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserClbWafRegions
// During the addition and modification of Domain Configuration for CLB-type WAF, it is required to display the supported region list for CLB-type WAF (clb-waf) through DescribeUserClbWafRegions to obtain the currently available region list for the customer.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeUserClbWafRegions(request *DescribeUserClbWafRegionsRequest) (response *DescribeUserClbWafRegionsResponse, err error) {
    return c.DescribeUserClbWafRegionsWithContext(context.Background(), request)
}

// DescribeUserClbWafRegions
// During the addition and modification of Domain Configuration for CLB-type WAF, it is required to display the supported region list for CLB-type WAF (clb-waf) through DescribeUserClbWafRegions to obtain the currently available region list for the customer.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeUserClbWafRegionsWithContext(ctx context.Context, request *DescribeUserClbWafRegionsRequest) (response *DescribeUserClbWafRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeUserClbWafRegionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DescribeUserClbWafRegions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserClbWafRegions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserClbWafRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserDomainInfoRequest() (request *DescribeUserDomainInfoRequest) {
    request = &DescribeUserDomainInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeUserDomainInfo")
    
    
    return
}

func NewDescribeUserDomainInfoResponse() (response *DescribeUserDomainInfoResponse) {
    response = &DescribeUserDomainInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserDomainInfo
// Query Domain Information for SaaS and CLB
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUserDomainInfo(request *DescribeUserDomainInfoRequest) (response *DescribeUserDomainInfoResponse, err error) {
    return c.DescribeUserDomainInfoWithContext(context.Background(), request)
}

// DescribeUserDomainInfo
// Query Domain Information for SaaS and CLB
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUserDomainInfoWithContext(ctx context.Context, request *DescribeUserDomainInfoRequest) (response *DescribeUserDomainInfoResponse, err error) {
    if request == nil {
        request = NewDescribeUserDomainInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DescribeUserDomainInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserDomainInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserDomainInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserLevelRequest() (request *DescribeUserLevelRequest) {
    request = &DescribeUserLevelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeUserLevel")
    
    
    return
}

func NewDescribeUserLevelResponse() (response *DescribeUserLevelResponse) {
    response = &DescribeUserLevelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserLevel
// Obtain the user protection rule level
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeUserLevel(request *DescribeUserLevelRequest) (response *DescribeUserLevelResponse, err error) {
    return c.DescribeUserLevelWithContext(context.Background(), request)
}

// DescribeUserLevel
// Obtain the user protection rule level
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeUserLevelWithContext(ctx context.Context, request *DescribeUserLevelRequest) (response *DescribeUserLevelResponse, err error) {
    if request == nil {
        request = NewDescribeUserLevelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DescribeUserLevel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserLevel require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserLevelResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVipInfoRequest() (request *DescribeVipInfoRequest) {
    request = &DescribeVipInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeVipInfo")
    
    
    return
}

func NewDescribeVipInfoResponse() (response *DescribeVipInfoResponse) {
    response = &DescribeVipInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVipInfo
// Query VIP information based on filter criteria
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeVipInfo(request *DescribeVipInfoRequest) (response *DescribeVipInfoResponse, err error) {
    return c.DescribeVipInfoWithContext(context.Background(), request)
}

// DescribeVipInfo
// Query VIP information based on filter criteria
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeVipInfoWithContext(ctx context.Context, request *DescribeVipInfoRequest) (response *DescribeVipInfoResponse, err error) {
    if request == nil {
        request = NewDescribeVipInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DescribeVipInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVipInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVipInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWebshellStatusRequest() (request *DescribeWebshellStatusRequest) {
    request = &DescribeWebshellStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeWebshellStatus")
    
    
    return
}

func NewDescribeWebshellStatusResponse() (response *DescribeWebshellStatusResponse) {
    response = &DescribeWebshellStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWebshellStatus
// Obtain the webshell status of a domain
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeWebshellStatus(request *DescribeWebshellStatusRequest) (response *DescribeWebshellStatusResponse, err error) {
    return c.DescribeWebshellStatusWithContext(context.Background(), request)
}

// DescribeWebshellStatus
// Obtain the webshell status of a domain
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeWebshellStatusWithContext(ctx context.Context, request *DescribeWebshellStatusRequest) (response *DescribeWebshellStatusResponse, err error) {
    if request == nil {
        request = NewDescribeWebshellStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "DescribeWebshellStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWebshellStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWebshellStatusResponse()
    err = c.Send(request, response)
    return
}

func NewFreshAntiFakeUrlRequest() (request *FreshAntiFakeUrlRequest) {
    request = &FreshAntiFakeUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "FreshAntiFakeUrl")
    
    
    return
}

func NewFreshAntiFakeUrlResponse() (response *FreshAntiFakeUrlResponse) {
    response = &FreshAntiFakeUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// FreshAntiFakeUrl
// Refresh a tamper-proof URL
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) FreshAntiFakeUrl(request *FreshAntiFakeUrlRequest) (response *FreshAntiFakeUrlResponse, err error) {
    return c.FreshAntiFakeUrlWithContext(context.Background(), request)
}

// FreshAntiFakeUrl
// Refresh a tamper-proof URL
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) FreshAntiFakeUrlWithContext(ctx context.Context, request *FreshAntiFakeUrlRequest) (response *FreshAntiFakeUrlResponse, err error) {
    if request == nil {
        request = NewFreshAntiFakeUrlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "FreshAntiFakeUrl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("FreshAntiFakeUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewFreshAntiFakeUrlResponse()
    err = c.Send(request, response)
    return
}

func NewGenerateDealsAndPayNewRequest() (request *GenerateDealsAndPayNewRequest) {
    request = &GenerateDealsAndPayNewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "GenerateDealsAndPayNew")
    
    
    return
}

func NewGenerateDealsAndPayNewResponse() (response *GenerateDealsAndPayNewResponse) {
    response = &GenerateDealsAndPayNewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GenerateDealsAndPayNew
// Billing Resource Purchase, Renewal Order API
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GenerateDealsAndPayNew(request *GenerateDealsAndPayNewRequest) (response *GenerateDealsAndPayNewResponse, err error) {
    return c.GenerateDealsAndPayNewWithContext(context.Background(), request)
}

// GenerateDealsAndPayNew
// Billing Resource Purchase, Renewal Order API
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GenerateDealsAndPayNewWithContext(ctx context.Context, request *GenerateDealsAndPayNewRequest) (response *GenerateDealsAndPayNewResponse, err error) {
    if request == nil {
        request = NewGenerateDealsAndPayNewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "GenerateDealsAndPayNew")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GenerateDealsAndPayNew require credential")
    }

    request.SetContext(ctx)
    
    response = NewGenerateDealsAndPayNewResponse()
    err = c.Send(request, response)
    return
}

func NewGetAttackHistogramRequest() (request *GetAttackHistogramRequest) {
    request = &GetAttackHistogramRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "GetAttackHistogram")
    
    
    return
}

func NewGetAttackHistogramResponse() (response *GetAttackHistogramResponse) {
    response = &GetAttackHistogramResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetAttackHistogram
// This API is used to generate a bar chart for the generation time of attack logs.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSDBOPERATIONFAILED = "FailedOperation.CLSDBOperationFailed"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERR = "InternalError.UnknownErr"
//  INVALIDPARAMETER_LOGICERR = "InvalidParameter.LogicErr"
//  INVALIDPARAMETER_QUERYSTRINGSYNTAXERR = "InvalidParameter.QueryStringSyntaxErr"
//  INVALIDPARAMETER_SQLSYNTAXERR = "InvalidParameter.SQLSyntaxErr"
//  INVALIDPARAMETER_TYPEMISMATCH = "InvalidParameter.TypeMismatch"
func (c *Client) GetAttackHistogram(request *GetAttackHistogramRequest) (response *GetAttackHistogramResponse, err error) {
    return c.GetAttackHistogramWithContext(context.Background(), request)
}

// GetAttackHistogram
// This API is used to generate a bar chart for the generation time of attack logs.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSDBOPERATIONFAILED = "FailedOperation.CLSDBOperationFailed"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERR = "InternalError.UnknownErr"
//  INVALIDPARAMETER_LOGICERR = "InvalidParameter.LogicErr"
//  INVALIDPARAMETER_QUERYSTRINGSYNTAXERR = "InvalidParameter.QueryStringSyntaxErr"
//  INVALIDPARAMETER_SQLSYNTAXERR = "InvalidParameter.SQLSyntaxErr"
//  INVALIDPARAMETER_TYPEMISMATCH = "InvalidParameter.TypeMismatch"
func (c *Client) GetAttackHistogramWithContext(ctx context.Context, request *GetAttackHistogramRequest) (response *GetAttackHistogramResponse, err error) {
    if request == nil {
        request = NewGetAttackHistogramRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "GetAttackHistogram")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetAttackHistogram require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetAttackHistogramResponse()
    err = c.Send(request, response)
    return
}

func NewGetAttackTotalCountRequest() (request *GetAttackTotalCountRequest) {
    request = &GetAttackTotalCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "GetAttackTotalCount")
    
    
    return
}

func NewGetAttackTotalCountResponse() (response *GetAttackTotalCountResponse) {
    response = &GetAttackTotalCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetAttackTotalCount
// Display total attack count by querying based on conditions
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSDBOPERATIONFAILED = "FailedOperation.CLSDBOperationFailed"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERR = "InternalError.UnknownErr"
//  INVALIDPARAMETER_LOGICERR = "InvalidParameter.LogicErr"
//  INVALIDPARAMETER_QUERYSTRINGSYNTAXERR = "InvalidParameter.QueryStringSyntaxErr"
//  INVALIDPARAMETER_SQLSYNTAXERR = "InvalidParameter.SQLSyntaxErr"
//  INVALIDPARAMETER_TYPEMISMATCH = "InvalidParameter.TypeMismatch"
func (c *Client) GetAttackTotalCount(request *GetAttackTotalCountRequest) (response *GetAttackTotalCountResponse, err error) {
    return c.GetAttackTotalCountWithContext(context.Background(), request)
}

// GetAttackTotalCount
// Display total attack count by querying based on conditions
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSDBOPERATIONFAILED = "FailedOperation.CLSDBOperationFailed"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERR = "InternalError.UnknownErr"
//  INVALIDPARAMETER_LOGICERR = "InvalidParameter.LogicErr"
//  INVALIDPARAMETER_QUERYSTRINGSYNTAXERR = "InvalidParameter.QueryStringSyntaxErr"
//  INVALIDPARAMETER_SQLSYNTAXERR = "InvalidParameter.SQLSyntaxErr"
//  INVALIDPARAMETER_TYPEMISMATCH = "InvalidParameter.TypeMismatch"
func (c *Client) GetAttackTotalCountWithContext(ctx context.Context, request *GetAttackTotalCountRequest) (response *GetAttackTotalCountResponse, err error) {
    if request == nil {
        request = NewGetAttackTotalCountRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "GetAttackTotalCount")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetAttackTotalCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetAttackTotalCountResponse()
    err = c.Send(request, response)
    return
}

func NewGetInstanceQpsLimitRequest() (request *GetInstanceQpsLimitRequest) {
    request = &GetInstanceQpsLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "GetInstanceQpsLimit")
    
    
    return
}

func NewGetInstanceQpsLimitResponse() (response *GetInstanceQpsLimitResponse) {
    response = &GetInstanceQpsLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetInstanceQpsLimit
// Obtain the elastic QPS limit of package instances
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSDBOPERATIONFAILED = "FailedOperation.CLSDBOperationFailed"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERR = "InternalError.UnknownErr"
//  INVALIDPARAMETER_LOGICERR = "InvalidParameter.LogicErr"
//  INVALIDPARAMETER_QUERYSTRINGSYNTAXERR = "InvalidParameter.QueryStringSyntaxErr"
//  INVALIDPARAMETER_SQLSYNTAXERR = "InvalidParameter.SQLSyntaxErr"
//  INVALIDPARAMETER_TYPEMISMATCH = "InvalidParameter.TypeMismatch"
func (c *Client) GetInstanceQpsLimit(request *GetInstanceQpsLimitRequest) (response *GetInstanceQpsLimitResponse, err error) {
    return c.GetInstanceQpsLimitWithContext(context.Background(), request)
}

// GetInstanceQpsLimit
// Obtain the elastic QPS limit of package instances
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSDBOPERATIONFAILED = "FailedOperation.CLSDBOperationFailed"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERR = "InternalError.UnknownErr"
//  INVALIDPARAMETER_LOGICERR = "InvalidParameter.LogicErr"
//  INVALIDPARAMETER_QUERYSTRINGSYNTAXERR = "InvalidParameter.QueryStringSyntaxErr"
//  INVALIDPARAMETER_SQLSYNTAXERR = "InvalidParameter.SQLSyntaxErr"
//  INVALIDPARAMETER_TYPEMISMATCH = "InvalidParameter.TypeMismatch"
func (c *Client) GetInstanceQpsLimitWithContext(ctx context.Context, request *GetInstanceQpsLimitRequest) (response *GetInstanceQpsLimitResponse, err error) {
    if request == nil {
        request = NewGetInstanceQpsLimitRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "GetInstanceQpsLimit")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetInstanceQpsLimit require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetInstanceQpsLimitResponse()
    err = c.Send(request, response)
    return
}

func NewImportIpAccessControlRequest() (request *ImportIpAccessControlRequest) {
    request = &ImportIpAccessControlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ImportIpAccessControl")
    
    
    return
}

func NewImportIpAccessControlResponse() (response *ImportIpAccessControlResponse) {
    response = &ImportIpAccessControlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ImportIpAccessControl
// This API is used to import IP allowlists/blocklists.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSDBOPERATIONFAILED = "FailedOperation.CLSDBOperationFailed"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERR = "InternalError.UnknownErr"
//  INVALIDPARAMETER_LOGICERR = "InvalidParameter.LogicErr"
//  INVALIDPARAMETER_QUERYSTRINGSYNTAXERR = "InvalidParameter.QueryStringSyntaxErr"
//  INVALIDPARAMETER_SQLSYNTAXERR = "InvalidParameter.SQLSyntaxErr"
//  INVALIDPARAMETER_TYPEMISMATCH = "InvalidParameter.TypeMismatch"
func (c *Client) ImportIpAccessControl(request *ImportIpAccessControlRequest) (response *ImportIpAccessControlResponse, err error) {
    return c.ImportIpAccessControlWithContext(context.Background(), request)
}

// ImportIpAccessControl
// This API is used to import IP allowlists/blocklists.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSDBOPERATIONFAILED = "FailedOperation.CLSDBOperationFailed"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERR = "InternalError.UnknownErr"
//  INVALIDPARAMETER_LOGICERR = "InvalidParameter.LogicErr"
//  INVALIDPARAMETER_QUERYSTRINGSYNTAXERR = "InvalidParameter.QueryStringSyntaxErr"
//  INVALIDPARAMETER_SQLSYNTAXERR = "InvalidParameter.SQLSyntaxErr"
//  INVALIDPARAMETER_TYPEMISMATCH = "InvalidParameter.TypeMismatch"
func (c *Client) ImportIpAccessControlWithContext(ctx context.Context, request *ImportIpAccessControlRequest) (response *ImportIpAccessControlResponse, err error) {
    if request == nil {
        request = NewImportIpAccessControlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "ImportIpAccessControl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ImportIpAccessControl require credential")
    }

    request.SetContext(ctx)
    
    response = NewImportIpAccessControlResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAntiFakeUrlRequest() (request *ModifyAntiFakeUrlRequest) {
    request = &ModifyAntiFakeUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyAntiFakeUrl")
    
    
    return
}

func NewModifyAntiFakeUrlResponse() (response *ModifyAntiFakeUrlResponse) {
    response = &ModifyAntiFakeUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAntiFakeUrl
// Edit a tamper-proof URL
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyAntiFakeUrl(request *ModifyAntiFakeUrlRequest) (response *ModifyAntiFakeUrlResponse, err error) {
    return c.ModifyAntiFakeUrlWithContext(context.Background(), request)
}

// ModifyAntiFakeUrl
// Edit a tamper-proof URL
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyAntiFakeUrlWithContext(ctx context.Context, request *ModifyAntiFakeUrlRequest) (response *ModifyAntiFakeUrlResponse, err error) {
    if request == nil {
        request = NewModifyAntiFakeUrlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "ModifyAntiFakeUrl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAntiFakeUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAntiFakeUrlResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAntiFakeUrlStatusRequest() (request *ModifyAntiFakeUrlStatusRequest) {
    request = &ModifyAntiFakeUrlStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyAntiFakeUrlStatus")
    
    
    return
}

func NewModifyAntiFakeUrlStatusResponse() (response *ModifyAntiFakeUrlStatusResponse) {
    response = &ModifyAntiFakeUrlStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAntiFakeUrlStatus
// Toggle tamper-proof switch
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAntiFakeUrlStatus(request *ModifyAntiFakeUrlStatusRequest) (response *ModifyAntiFakeUrlStatusResponse, err error) {
    return c.ModifyAntiFakeUrlStatusWithContext(context.Background(), request)
}

// ModifyAntiFakeUrlStatus
// Toggle tamper-proof switch
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAntiFakeUrlStatusWithContext(ctx context.Context, request *ModifyAntiFakeUrlStatusRequest) (response *ModifyAntiFakeUrlStatusResponse, err error) {
    if request == nil {
        request = NewModifyAntiFakeUrlStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "ModifyAntiFakeUrlStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAntiFakeUrlStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAntiFakeUrlStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAntiInfoLeakRuleStatusRequest() (request *ModifyAntiInfoLeakRuleStatusRequest) {
    request = &ModifyAntiInfoLeakRuleStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyAntiInfoLeakRuleStatus")
    
    
    return
}

func NewModifyAntiInfoLeakRuleStatusResponse() (response *ModifyAntiInfoLeakRuleStatusResponse) {
    response = &ModifyAntiInfoLeakRuleStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAntiInfoLeakRuleStatus
// Information leakage prevention toggle rule switch
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyAntiInfoLeakRuleStatus(request *ModifyAntiInfoLeakRuleStatusRequest) (response *ModifyAntiInfoLeakRuleStatusResponse, err error) {
    return c.ModifyAntiInfoLeakRuleStatusWithContext(context.Background(), request)
}

// ModifyAntiInfoLeakRuleStatus
// Information leakage prevention toggle rule switch
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyAntiInfoLeakRuleStatusWithContext(ctx context.Context, request *ModifyAntiInfoLeakRuleStatusRequest) (response *ModifyAntiInfoLeakRuleStatusResponse, err error) {
    if request == nil {
        request = NewModifyAntiInfoLeakRuleStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "ModifyAntiInfoLeakRuleStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAntiInfoLeakRuleStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAntiInfoLeakRuleStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAntiInfoLeakRulesRequest() (request *ModifyAntiInfoLeakRulesRequest) {
    request = &ModifyAntiInfoLeakRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyAntiInfoLeakRules")
    
    
    return
}

func NewModifyAntiInfoLeakRulesResponse() (response *ModifyAntiInfoLeakRulesResponse) {
    response = &ModifyAntiInfoLeakRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAntiInfoLeakRules
// Edit an information leakage prevention rule
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyAntiInfoLeakRules(request *ModifyAntiInfoLeakRulesRequest) (response *ModifyAntiInfoLeakRulesResponse, err error) {
    return c.ModifyAntiInfoLeakRulesWithContext(context.Background(), request)
}

// ModifyAntiInfoLeakRules
// Edit an information leakage prevention rule
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyAntiInfoLeakRulesWithContext(ctx context.Context, request *ModifyAntiInfoLeakRulesRequest) (response *ModifyAntiInfoLeakRulesResponse, err error) {
    if request == nil {
        request = NewModifyAntiInfoLeakRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "ModifyAntiInfoLeakRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAntiInfoLeakRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAntiInfoLeakRulesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApiAnalyzeStatusRequest() (request *ModifyApiAnalyzeStatusRequest) {
    request = &ModifyApiAnalyzeStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyApiAnalyzeStatus")
    
    
    return
}

func NewModifyApiAnalyzeStatusResponse() (response *ModifyApiAnalyzeStatusResponse) {
    response = &ModifyApiAnalyzeStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApiAnalyzeStatus
// API analysis page switch
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyApiAnalyzeStatus(request *ModifyApiAnalyzeStatusRequest) (response *ModifyApiAnalyzeStatusResponse, err error) {
    return c.ModifyApiAnalyzeStatusWithContext(context.Background(), request)
}

// ModifyApiAnalyzeStatus
// API analysis page switch
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyApiAnalyzeStatusWithContext(ctx context.Context, request *ModifyApiAnalyzeStatusRequest) (response *ModifyApiAnalyzeStatusResponse, err error) {
    if request == nil {
        request = NewModifyApiAnalyzeStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "ModifyApiAnalyzeStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApiAnalyzeStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApiAnalyzeStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBotStatusRequest() (request *ModifyBotStatusRequest) {
    request = &ModifyBotStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyBotStatus")
    
    
    return
}

func NewModifyBotStatusResponse() (response *ModifyBotStatusResponse) {
    response = &ModifyBotStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyBotStatus
// Bot_V2 bot master switch update
//
// error code that may be returned:
//  AUTHFAILURE_ERRCODENOPURCHASED = "AuthFailure.ErrCodeNoPurchased"
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION_INVALIDREQUEST = "UnsupportedOperation.InvalidRequest"
func (c *Client) ModifyBotStatus(request *ModifyBotStatusRequest) (response *ModifyBotStatusResponse, err error) {
    return c.ModifyBotStatusWithContext(context.Background(), request)
}

// ModifyBotStatus
// Bot_V2 bot master switch update
//
// error code that may be returned:
//  AUTHFAILURE_ERRCODENOPURCHASED = "AuthFailure.ErrCodeNoPurchased"
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION_INVALIDREQUEST = "UnsupportedOperation.InvalidRequest"
func (c *Client) ModifyBotStatusWithContext(ctx context.Context, request *ModifyBotStatusRequest) (response *ModifyBotStatusResponse, err error) {
    if request == nil {
        request = NewModifyBotStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "ModifyBotStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBotStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBotStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCustomRuleRequest() (request *ModifyCustomRuleRequest) {
    request = &ModifyCustomRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyCustomRule")
    
    
    return
}

func NewModifyCustomRuleResponse() (response *ModifyCustomRuleResponse) {
    response = &ModifyCustomRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCustomRule
// This API is used to edit a custom rule.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCustomRule(request *ModifyCustomRuleRequest) (response *ModifyCustomRuleResponse, err error) {
    return c.ModifyCustomRuleWithContext(context.Background(), request)
}

// ModifyCustomRule
// This API is used to edit a custom rule.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCustomRuleWithContext(ctx context.Context, request *ModifyCustomRuleRequest) (response *ModifyCustomRuleResponse, err error) {
    if request == nil {
        request = NewModifyCustomRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "ModifyCustomRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCustomRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCustomRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCustomRuleStatusRequest() (request *ModifyCustomRuleStatusRequest) {
    request = &ModifyCustomRuleStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyCustomRuleStatus")
    
    
    return
}

func NewModifyCustomRuleStatusResponse() (response *ModifyCustomRuleStatusResponse) {
    response = &ModifyCustomRuleStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCustomRuleStatus
// Enable or disable access control (from custom policy)
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCustomRuleStatus(request *ModifyCustomRuleStatusRequest) (response *ModifyCustomRuleStatusResponse, err error) {
    return c.ModifyCustomRuleStatusWithContext(context.Background(), request)
}

// ModifyCustomRuleStatus
// Enable or disable access control (from custom policy)
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCustomRuleStatusWithContext(ctx context.Context, request *ModifyCustomRuleStatusRequest) (response *ModifyCustomRuleStatusResponse, err error) {
    if request == nil {
        request = NewModifyCustomRuleStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "ModifyCustomRuleStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCustomRuleStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCustomRuleStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCustomWhiteRuleRequest() (request *ModifyCustomWhiteRuleRequest) {
    request = &ModifyCustomWhiteRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyCustomWhiteRule")
    
    
    return
}

func NewModifyCustomWhiteRuleResponse() (response *ModifyCustomWhiteRuleResponse) {
    response = &ModifyCustomWhiteRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCustomWhiteRule
// This API is used to edit a precise allowlist.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCustomWhiteRule(request *ModifyCustomWhiteRuleRequest) (response *ModifyCustomWhiteRuleResponse, err error) {
    return c.ModifyCustomWhiteRuleWithContext(context.Background(), request)
}

// ModifyCustomWhiteRule
// This API is used to edit a precise allowlist.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCustomWhiteRuleWithContext(ctx context.Context, request *ModifyCustomWhiteRuleRequest) (response *ModifyCustomWhiteRuleResponse, err error) {
    if request == nil {
        request = NewModifyCustomWhiteRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "ModifyCustomWhiteRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCustomWhiteRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCustomWhiteRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCustomWhiteRuleStatusRequest() (request *ModifyCustomWhiteRuleStatusRequest) {
    request = &ModifyCustomWhiteRuleStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyCustomWhiteRuleStatus")
    
    
    return
}

func NewModifyCustomWhiteRuleStatusResponse() (response *ModifyCustomWhiteRuleStatusResponse) {
    response = &ModifyCustomWhiteRuleStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCustomWhiteRuleStatus
// Enable or disable a precision allowlist
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCustomWhiteRuleStatus(request *ModifyCustomWhiteRuleStatusRequest) (response *ModifyCustomWhiteRuleStatusResponse, err error) {
    return c.ModifyCustomWhiteRuleStatusWithContext(context.Background(), request)
}

// ModifyCustomWhiteRuleStatus
// Enable or disable a precision allowlist
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCustomWhiteRuleStatusWithContext(ctx context.Context, request *ModifyCustomWhiteRuleStatusRequest) (response *ModifyCustomWhiteRuleStatusResponse, err error) {
    if request == nil {
        request = NewModifyCustomWhiteRuleStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "ModifyCustomWhiteRuleStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCustomWhiteRuleStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCustomWhiteRuleStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDomainIpv6StatusRequest() (request *ModifyDomainIpv6StatusRequest) {
    request = &ModifyDomainIpv6StatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyDomainIpv6Status")
    
    
    return
}

func NewModifyDomainIpv6StatusResponse() (response *ModifyDomainIpv6StatusResponse) {
    response = &ModifyDomainIpv6StatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDomainIpv6Status
// Toggle the IPv6 switch
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_DOMAINIPV6INCONFIGERR = "ResourceUnavailable.DomainIpv6InConfigErr"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDomainIpv6Status(request *ModifyDomainIpv6StatusRequest) (response *ModifyDomainIpv6StatusResponse, err error) {
    return c.ModifyDomainIpv6StatusWithContext(context.Background(), request)
}

// ModifyDomainIpv6Status
// Toggle the IPv6 switch
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_DOMAINIPV6INCONFIGERR = "ResourceUnavailable.DomainIpv6InConfigErr"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDomainIpv6StatusWithContext(ctx context.Context, request *ModifyDomainIpv6StatusRequest) (response *ModifyDomainIpv6StatusResponse, err error) {
    if request == nil {
        request = NewModifyDomainIpv6StatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "ModifyDomainIpv6Status")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDomainIpv6Status require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDomainIpv6StatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDomainPostActionRequest() (request *ModifyDomainPostActionRequest) {
    request = &ModifyDomainPostActionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyDomainPostAction")
    
    
    return
}

func NewModifyDomainPostActionResponse() (response *ModifyDomainPostActionResponse) {
    response = &ModifyDomainPostActionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDomainPostAction
// This API is used to modify the domain shipping status.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CKAFKAINTERNALERROR = "FailedOperation.CKafkaInternalError"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDomainPostAction(request *ModifyDomainPostActionRequest) (response *ModifyDomainPostActionResponse, err error) {
    return c.ModifyDomainPostActionWithContext(context.Background(), request)
}

// ModifyDomainPostAction
// This API is used to modify the domain shipping status.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CKAFKAINTERNALERROR = "FailedOperation.CKafkaInternalError"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDomainPostActionWithContext(ctx context.Context, request *ModifyDomainPostActionRequest) (response *ModifyDomainPostActionResponse, err error) {
    if request == nil {
        request = NewModifyDomainPostActionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "ModifyDomainPostAction")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDomainPostAction require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDomainPostActionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDomainsCLSStatusRequest() (request *ModifyDomainsCLSStatusRequest) {
    request = &ModifyDomainsCLSStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyDomainsCLSStatus")
    
    
    return
}

func NewModifyDomainsCLSStatusResponse() (response *ModifyDomainsCLSStatusResponse) {
    response = &ModifyDomainsCLSStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDomainsCLSStatus
// Enable or disable access log for domain list
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDomainsCLSStatus(request *ModifyDomainsCLSStatusRequest) (response *ModifyDomainsCLSStatusResponse, err error) {
    return c.ModifyDomainsCLSStatusWithContext(context.Background(), request)
}

// ModifyDomainsCLSStatus
// Enable or disable access log for domain list
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDomainsCLSStatusWithContext(ctx context.Context, request *ModifyDomainsCLSStatusRequest) (response *ModifyDomainsCLSStatusResponse, err error) {
    if request == nil {
        request = NewModifyDomainsCLSStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "ModifyDomainsCLSStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDomainsCLSStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDomainsCLSStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyHostRequest() (request *ModifyHostRequest) {
    request = &ModifyHostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyHost")
    
    
    return
}

func NewModifyHostResponse() (response *ModifyHostResponse) {
    response = &ModifyHostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyHost
// This API is used to edit the configuration of domain names protected by CLB WAF.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyHost(request *ModifyHostRequest) (response *ModifyHostResponse, err error) {
    return c.ModifyHostWithContext(context.Background(), request)
}

// ModifyHost
// This API is used to edit the configuration of domain names protected by CLB WAF.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyHostWithContext(ctx context.Context, request *ModifyHostRequest) (response *ModifyHostResponse, err error) {
    if request == nil {
        request = NewModifyHostRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "ModifyHost")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyHost require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyHostResponse()
    err = c.Send(request, response)
    return
}

func NewModifyHostFlowModeRequest() (request *ModifyHostFlowModeRequest) {
    request = &ModifyHostFlowModeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyHostFlowMode")
    
    
    return
}

func NewModifyHostFlowModeResponse() (response *ModifyHostFlowModeResponse) {
    response = &ModifyHostFlowModeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyHostFlowMode
// This API is used to set the traffic mode for domain names protected by CLB WAF. The mode can be mirror mode or cleaning mode.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyHostFlowMode(request *ModifyHostFlowModeRequest) (response *ModifyHostFlowModeResponse, err error) {
    return c.ModifyHostFlowModeWithContext(context.Background(), request)
}

// ModifyHostFlowMode
// This API is used to set the traffic mode for domain names protected by CLB WAF. The mode can be mirror mode or cleaning mode.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyHostFlowModeWithContext(ctx context.Context, request *ModifyHostFlowModeRequest) (response *ModifyHostFlowModeResponse, err error) {
    if request == nil {
        request = NewModifyHostFlowModeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "ModifyHostFlowMode")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyHostFlowMode require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyHostFlowModeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyHostModeRequest() (request *ModifyHostModeRequest) {
    request = &ModifyHostModeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyHostMode")
    
    
    return
}

func NewModifyHostModeResponse() (response *ModifyHostModeResponse) {
    response = &ModifyHostModeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyHostMode
// Set CLB WAF protection domain status
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyHostMode(request *ModifyHostModeRequest) (response *ModifyHostModeResponse, err error) {
    return c.ModifyHostModeWithContext(context.Background(), request)
}

// ModifyHostMode
// Set CLB WAF protection domain status
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyHostModeWithContext(ctx context.Context, request *ModifyHostModeRequest) (response *ModifyHostModeResponse, err error) {
    if request == nil {
        request = NewModifyHostModeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "ModifyHostMode")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyHostMode require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyHostModeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyHostStatusRequest() (request *ModifyHostStatusRequest) {
    request = &ModifyHostStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyHostStatus")
    
    
    return
}

func NewModifyHostStatusResponse() (response *ModifyHostStatusResponse) {
    response = &ModifyHostStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyHostStatus
// This API is used to enable or disable CLB WAF for a protected domain name.
//
// Batch operation is supported.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyHostStatus(request *ModifyHostStatusRequest) (response *ModifyHostStatusResponse, err error) {
    return c.ModifyHostStatusWithContext(context.Background(), request)
}

// ModifyHostStatus
// This API is used to enable or disable CLB WAF for a protected domain name.
//
// Batch operation is supported.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyHostStatusWithContext(ctx context.Context, request *ModifyHostStatusRequest) (response *ModifyHostStatusResponse, err error) {
    if request == nil {
        request = NewModifyHostStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "ModifyHostStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyHostStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyHostStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceElasticModeRequest() (request *ModifyInstanceElasticModeRequest) {
    request = &ModifyInstanceElasticModeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyInstanceElasticMode")
    
    
    return
}

func NewModifyInstanceElasticModeResponse() (response *ModifyInstanceElasticModeResponse) {
    response = &ModifyInstanceElasticModeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstanceElasticMode
// Modify the QPS elastic billing switch for an instance
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyInstanceElasticMode(request *ModifyInstanceElasticModeRequest) (response *ModifyInstanceElasticModeResponse, err error) {
    return c.ModifyInstanceElasticModeWithContext(context.Background(), request)
}

// ModifyInstanceElasticMode
// Modify the QPS elastic billing switch for an instance
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyInstanceElasticModeWithContext(ctx context.Context, request *ModifyInstanceElasticModeRequest) (response *ModifyInstanceElasticModeResponse, err error) {
    if request == nil {
        request = NewModifyInstanceElasticModeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "ModifyInstanceElasticMode")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceElasticMode require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceElasticModeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceNameRequest() (request *ModifyInstanceNameRequest) {
    request = &ModifyInstanceNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyInstanceName")
    
    
    return
}

func NewModifyInstanceNameResponse() (response *ModifyInstanceNameResponse) {
    response = &ModifyInstanceNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstanceName
// Modify instance name
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyInstanceName(request *ModifyInstanceNameRequest) (response *ModifyInstanceNameResponse, err error) {
    return c.ModifyInstanceNameWithContext(context.Background(), request)
}

// ModifyInstanceName
// Modify instance name
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyInstanceNameWithContext(ctx context.Context, request *ModifyInstanceNameRequest) (response *ModifyInstanceNameResponse, err error) {
    if request == nil {
        request = NewModifyInstanceNameRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "ModifyInstanceName")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceName require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceQpsLimitRequest() (request *ModifyInstanceQpsLimitRequest) {
    request = &ModifyInstanceQpsLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyInstanceQpsLimit")
    
    
    return
}

func NewModifyInstanceQpsLimitResponse() (response *ModifyInstanceQpsLimitResponse) {
    response = &ModifyInstanceQpsLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstanceQpsLimit
// Set elastic QPS limit for package instances
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyInstanceQpsLimit(request *ModifyInstanceQpsLimitRequest) (response *ModifyInstanceQpsLimitResponse, err error) {
    return c.ModifyInstanceQpsLimitWithContext(context.Background(), request)
}

// ModifyInstanceQpsLimit
// Set elastic QPS limit for package instances
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyInstanceQpsLimitWithContext(ctx context.Context, request *ModifyInstanceQpsLimitRequest) (response *ModifyInstanceQpsLimitResponse, err error) {
    if request == nil {
        request = NewModifyInstanceQpsLimitRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "ModifyInstanceQpsLimit")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceQpsLimit require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceQpsLimitResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceRenewFlagRequest() (request *ModifyInstanceRenewFlagRequest) {
    request = &ModifyInstanceRenewFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyInstanceRenewFlag")
    
    
    return
}

func NewModifyInstanceRenewFlagResponse() (response *ModifyInstanceRenewFlagResponse) {
    response = &ModifyInstanceRenewFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstanceRenewFlag
// Enable or disable auto-renewal for instance
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyInstanceRenewFlag(request *ModifyInstanceRenewFlagRequest) (response *ModifyInstanceRenewFlagResponse, err error) {
    return c.ModifyInstanceRenewFlagWithContext(context.Background(), request)
}

// ModifyInstanceRenewFlag
// Enable or disable auto-renewal for instance
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyInstanceRenewFlagWithContext(ctx context.Context, request *ModifyInstanceRenewFlagRequest) (response *ModifyInstanceRenewFlagResponse, err error) {
    if request == nil {
        request = NewModifyInstanceRenewFlagRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "ModifyInstanceRenewFlag")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceRenewFlag require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceRenewFlagResponse()
    err = c.Send(request, response)
    return
}

func NewModifyIpAccessControlRequest() (request *ModifyIpAccessControlRequest) {
    request = &ModifyIpAccessControlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyIpAccessControl")
    
    
    return
}

func NewModifyIpAccessControlResponse() (response *ModifyIpAccessControlResponse) {
    response = &ModifyIpAccessControlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyIpAccessControl
// This API is used to edit WAF IP allowlists/blocklists.
//
// error code that may be returned:
//  FAILEDOPERATION_THENUMBEROFADDEDBLACKANDWHITELISTEXCEEDSTHEUPPERLIMIT = "FailedOperation.TheNumberOfAddedBlackAndWhiteListExceedsTheUpperLimit"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_SPECIFICATIONERR = "LimitExceeded.SpecificationErr"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyIpAccessControl(request *ModifyIpAccessControlRequest) (response *ModifyIpAccessControlResponse, err error) {
    return c.ModifyIpAccessControlWithContext(context.Background(), request)
}

// ModifyIpAccessControl
// This API is used to edit WAF IP allowlists/blocklists.
//
// error code that may be returned:
//  FAILEDOPERATION_THENUMBEROFADDEDBLACKANDWHITELISTEXCEEDSTHEUPPERLIMIT = "FailedOperation.TheNumberOfAddedBlackAndWhiteListExceedsTheUpperLimit"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_SPECIFICATIONERR = "LimitExceeded.SpecificationErr"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyIpAccessControlWithContext(ctx context.Context, request *ModifyIpAccessControlRequest) (response *ModifyIpAccessControlResponse, err error) {
    if request == nil {
        request = NewModifyIpAccessControlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "ModifyIpAccessControl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyIpAccessControl require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyIpAccessControlResponse()
    err = c.Send(request, response)
    return
}

func NewModifyModuleStatusRequest() (request *ModifyModuleStatusRequest) {
    request = &ModifyModuleStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyModuleStatus")
    
    
    return
}

func NewModifyModuleStatusResponse() (response *ModifyModuleStatusResponse) {
    response = &ModifyModuleStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyModuleStatus
// Set the switch for the basic security module under a certain domain
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyModuleStatus(request *ModifyModuleStatusRequest) (response *ModifyModuleStatusResponse, err error) {
    return c.ModifyModuleStatusWithContext(context.Background(), request)
}

// ModifyModuleStatus
// Set the switch for the basic security module under a certain domain
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyModuleStatusWithContext(ctx context.Context, request *ModifyModuleStatusRequest) (response *ModifyModuleStatusResponse, err error) {
    if request == nil {
        request = NewModifyModuleStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "ModifyModuleStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyModuleStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyModuleStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyObjectRequest() (request *ModifyObjectRequest) {
    request = &ModifyObjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyObject")
    
    
    return
}

func NewModifyObjectResponse() (response *ModifyObjectResponse) {
    response = &ModifyObjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyObject
// Modify protection object
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyObject(request *ModifyObjectRequest) (response *ModifyObjectResponse, err error) {
    return c.ModifyObjectWithContext(context.Background(), request)
}

// ModifyObject
// Modify protection object
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyObjectWithContext(ctx context.Context, request *ModifyObjectRequest) (response *ModifyObjectResponse, err error) {
    if request == nil {
        request = NewModifyObjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "ModifyObject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyObject require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyObjectResponse()
    err = c.Send(request, response)
    return
}

func NewModifyProtectionStatusRequest() (request *ModifyProtectionStatusRequest) {
    request = &ModifyProtectionStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyProtectionStatus")
    
    
    return
}

func NewModifyProtectionStatusResponse() (response *ModifyProtectionStatusResponse) {
    response = &ModifyProtectionStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyProtectionStatus
// This API is used to obtain the enabling status of the basic security protection module of WAF.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_SPECIFICATIONERR = "LimitExceeded.SpecificationErr"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyProtectionStatus(request *ModifyProtectionStatusRequest) (response *ModifyProtectionStatusResponse, err error) {
    return c.ModifyProtectionStatusWithContext(context.Background(), request)
}

// ModifyProtectionStatus
// This API is used to obtain the enabling status of the basic security protection module of WAF.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_SPECIFICATIONERR = "LimitExceeded.SpecificationErr"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyProtectionStatusWithContext(ctx context.Context, request *ModifyProtectionStatusRequest) (response *ModifyProtectionStatusResponse, err error) {
    if request == nil {
        request = NewModifyProtectionStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "ModifyProtectionStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyProtectionStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyProtectionStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifySpartaProtectionRequest() (request *ModifySpartaProtectionRequest) {
    request = &ModifySpartaProtectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifySpartaProtection")
    
    
    return
}

func NewModifySpartaProtectionResponse() (response *ModifySpartaProtectionResponse) {
    response = &ModifySpartaProtectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySpartaProtection
// This API is used to edit the configuration of domain names protected by SaaS WAF.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MYSQLDBOPERATIONFAILED = "FailedOperation.MysqlDBOperationFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ASYNCHRONOUSCALLFAILED = "InternalError.AsynchronousCallFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CERTIFICATIONPARAMETERERR = "InvalidParameter.CertificationParameterErr"
//  INVALIDPARAMETER_PORTPARAMETERERR = "InvalidParameter.PortParameterErr"
//  INVALIDPARAMETER_PROTECTIONDOMAINPARAMETERERR = "InvalidParameter.ProtectionDomainParameterErr"
//  INVALIDPARAMETER_SUPPORTTLSCONFFAILED = "InvalidParameter.SupportTLSConfFailed"
//  INVALIDPARAMETER_TLSPARAMETERERR = "InvalidParameter.TLSParameterErr"
//  INVALIDPARAMETER_UPSTREAMPARAMETERERR = "InvalidParameter.UpstreamParameterErr"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifySpartaProtection(request *ModifySpartaProtectionRequest) (response *ModifySpartaProtectionResponse, err error) {
    return c.ModifySpartaProtectionWithContext(context.Background(), request)
}

// ModifySpartaProtection
// This API is used to edit the configuration of domain names protected by SaaS WAF.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MYSQLDBOPERATIONFAILED = "FailedOperation.MysqlDBOperationFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ASYNCHRONOUSCALLFAILED = "InternalError.AsynchronousCallFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CERTIFICATIONPARAMETERERR = "InvalidParameter.CertificationParameterErr"
//  INVALIDPARAMETER_PORTPARAMETERERR = "InvalidParameter.PortParameterErr"
//  INVALIDPARAMETER_PROTECTIONDOMAINPARAMETERERR = "InvalidParameter.ProtectionDomainParameterErr"
//  INVALIDPARAMETER_SUPPORTTLSCONFFAILED = "InvalidParameter.SupportTLSConfFailed"
//  INVALIDPARAMETER_TLSPARAMETERERR = "InvalidParameter.TLSParameterErr"
//  INVALIDPARAMETER_UPSTREAMPARAMETERERR = "InvalidParameter.UpstreamParameterErr"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifySpartaProtectionWithContext(ctx context.Context, request *ModifySpartaProtectionRequest) (response *ModifySpartaProtectionResponse, err error) {
    if request == nil {
        request = NewModifySpartaProtectionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "ModifySpartaProtection")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySpartaProtection require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySpartaProtectionResponse()
    err = c.Send(request, response)
    return
}

func NewModifySpartaProtectionModeRequest() (request *ModifySpartaProtectionModeRequest) {
    request = &ModifySpartaProtectionModeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifySpartaProtectionMode")
    
    
    return
}

func NewModifySpartaProtectionModeResponse() (response *ModifySpartaProtectionModeResponse) {
    response = &ModifySpartaProtectionModeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySpartaProtectionMode
// Set WAF protection status
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_SPECIFICATIONERR = "LimitExceeded.SpecificationErr"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifySpartaProtectionMode(request *ModifySpartaProtectionModeRequest) (response *ModifySpartaProtectionModeResponse, err error) {
    return c.ModifySpartaProtectionModeWithContext(context.Background(), request)
}

// ModifySpartaProtectionMode
// Set WAF protection status
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_SPECIFICATIONERR = "LimitExceeded.SpecificationErr"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifySpartaProtectionModeWithContext(ctx context.Context, request *ModifySpartaProtectionModeRequest) (response *ModifySpartaProtectionModeResponse, err error) {
    if request == nil {
        request = NewModifySpartaProtectionModeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "ModifySpartaProtectionMode")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySpartaProtectionMode require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySpartaProtectionModeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUserLevelRequest() (request *ModifyUserLevelRequest) {
    request = &ModifyUserLevelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyUserLevel")
    
    
    return
}

func NewModifyUserLevelResponse() (response *ModifyUserLevelResponse) {
    response = &ModifyUserLevelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyUserLevel
// Modify the user protection rule level
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyUserLevel(request *ModifyUserLevelRequest) (response *ModifyUserLevelResponse, err error) {
    return c.ModifyUserLevelWithContext(context.Background(), request)
}

// ModifyUserLevel
// Modify the user protection rule level
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyUserLevelWithContext(ctx context.Context, request *ModifyUserLevelRequest) (response *ModifyUserLevelResponse, err error) {
    if request == nil {
        request = NewModifyUserLevelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "ModifyUserLevel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUserLevel require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUserLevelResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUserSignatureRuleRequest() (request *ModifyUserSignatureRuleRequest) {
    request = &ModifyUserSignatureRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyUserSignatureRule")
    
    
    return
}

func NewModifyUserSignatureRuleResponse() (response *ModifyUserSignatureRuleResponse) {
    response = &ModifyUserSignatureRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyUserSignatureRule
// Modify user protection rules, turn on/off specific rules
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyUserSignatureRule(request *ModifyUserSignatureRuleRequest) (response *ModifyUserSignatureRuleResponse, err error) {
    return c.ModifyUserSignatureRuleWithContext(context.Background(), request)
}

// ModifyUserSignatureRule
// Modify user protection rules, turn on/off specific rules
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyUserSignatureRuleWithContext(ctx context.Context, request *ModifyUserSignatureRuleRequest) (response *ModifyUserSignatureRuleResponse, err error) {
    if request == nil {
        request = NewModifyUserSignatureRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "ModifyUserSignatureRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUserSignatureRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUserSignatureRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWebshellStatusRequest() (request *ModifyWebshellStatusRequest) {
    request = &ModifyWebshellStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyWebshellStatus")
    
    
    return
}

func NewModifyWebshellStatusResponse() (response *ModifyWebshellStatusResponse) {
    response = &ModifyWebshellStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyWebshellStatus
// Set the Webshell status of a domain.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyWebshellStatus(request *ModifyWebshellStatusRequest) (response *ModifyWebshellStatusResponse, err error) {
    return c.ModifyWebshellStatusWithContext(context.Background(), request)
}

// ModifyWebshellStatus
// Set the Webshell status of a domain.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyWebshellStatusWithContext(ctx context.Context, request *ModifyWebshellStatusRequest) (response *ModifyWebshellStatusResponse, err error) {
    if request == nil {
        request = NewModifyWebshellStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "ModifyWebshellStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyWebshellStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyWebshellStatusResponse()
    err = c.Send(request, response)
    return
}

func NewRefreshAccessCheckResultRequest() (request *RefreshAccessCheckResultRequest) {
    request = &RefreshAccessCheckResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "RefreshAccessCheckResult")
    
    
    return
}

func NewRefreshAccessCheckResultResponse() (response *RefreshAccessCheckResultResponse) {
    response = &RefreshAccessCheckResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RefreshAccessCheckResult
// Refresh integration check results. The backend will generate integration check tasks
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RefreshAccessCheckResult(request *RefreshAccessCheckResultRequest) (response *RefreshAccessCheckResultResponse, err error) {
    return c.RefreshAccessCheckResultWithContext(context.Background(), request)
}

// RefreshAccessCheckResult
// Refresh integration check results. The backend will generate integration check tasks
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RefreshAccessCheckResultWithContext(ctx context.Context, request *RefreshAccessCheckResultRequest) (response *RefreshAccessCheckResultResponse, err error) {
    if request == nil {
        request = NewRefreshAccessCheckResultRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "RefreshAccessCheckResult")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RefreshAccessCheckResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewRefreshAccessCheckResultResponse()
    err = c.Send(request, response)
    return
}

func NewSearchAttackLogRequest() (request *SearchAttackLogRequest) {
    request = &SearchAttackLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "SearchAttackLog")
    
    
    return
}

func NewSearchAttackLogResponse() (response *SearchAttackLogResponse) {
    response = &SearchAttackLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SearchAttackLog
// The new version of the CLS API has parameter changes, with query changed to query_string to support Lucene syntax for API search queries.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSDBOPERATIONFAILED = "FailedOperation.CLSDBOperationFailed"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERR = "InternalError.UnknownErr"
//  INVALIDPARAMETER_LOGICERR = "InvalidParameter.LogicErr"
//  INVALIDPARAMETER_QUERYSTRINGSYNTAXERR = "InvalidParameter.QueryStringSyntaxErr"
//  INVALIDPARAMETER_TYPEMISMATCH = "InvalidParameter.TypeMismatch"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SearchAttackLog(request *SearchAttackLogRequest) (response *SearchAttackLogResponse, err error) {
    return c.SearchAttackLogWithContext(context.Background(), request)
}

// SearchAttackLog
// The new version of the CLS API has parameter changes, with query changed to query_string to support Lucene syntax for API search queries.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSDBOPERATIONFAILED = "FailedOperation.CLSDBOperationFailed"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERR = "InternalError.UnknownErr"
//  INVALIDPARAMETER_LOGICERR = "InvalidParameter.LogicErr"
//  INVALIDPARAMETER_QUERYSTRINGSYNTAXERR = "InvalidParameter.QueryStringSyntaxErr"
//  INVALIDPARAMETER_TYPEMISMATCH = "InvalidParameter.TypeMismatch"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SearchAttackLogWithContext(ctx context.Context, request *SearchAttackLogRequest) (response *SearchAttackLogResponse, err error) {
    if request == nil {
        request = NewSearchAttackLogRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "SearchAttackLog")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchAttackLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchAttackLogResponse()
    err = c.Send(request, response)
    return
}

func NewSwitchElasticModeRequest() (request *SwitchElasticModeRequest) {
    request = &SwitchElasticModeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "SwitchElasticMode")
    
    
    return
}

func NewSwitchElasticModeResponse() (response *SwitchElasticModeResponse) {
    response = &SwitchElasticModeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SwitchElasticMode
// Toggle elasticity switch
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) SwitchElasticMode(request *SwitchElasticModeRequest) (response *SwitchElasticModeResponse, err error) {
    return c.SwitchElasticModeWithContext(context.Background(), request)
}

// SwitchElasticMode
// Toggle elasticity switch
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) SwitchElasticModeWithContext(ctx context.Context, request *SwitchElasticModeRequest) (response *SwitchElasticModeResponse, err error) {
    if request == nil {
        request = NewSwitchElasticModeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "SwitchElasticMode")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SwitchElasticMode require credential")
    }

    request.SetContext(ctx)
    
    response = NewSwitchElasticModeResponse()
    err = c.Send(request, response)
    return
}

func NewUpsertCCRuleRequest() (request *UpsertCCRuleRequest) {
    request = &UpsertCCRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "UpsertCCRule")
    
    
    return
}

func NewUpsertCCRuleResponse() (response *UpsertCCRuleResponse) {
    response = &UpsertCCRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpsertCCRule
// WAF CC V2 upsert API
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_SPECIFICATIONERR = "LimitExceeded.SpecificationErr"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpsertCCRule(request *UpsertCCRuleRequest) (response *UpsertCCRuleResponse, err error) {
    return c.UpsertCCRuleWithContext(context.Background(), request)
}

// UpsertCCRule
// WAF CC V2 upsert API
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_SPECIFICATIONERR = "LimitExceeded.SpecificationErr"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpsertCCRuleWithContext(ctx context.Context, request *UpsertCCRuleRequest) (response *UpsertCCRuleResponse, err error) {
    if request == nil {
        request = NewUpsertCCRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "UpsertCCRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpsertCCRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpsertCCRuleResponse()
    err = c.Send(request, response)
    return
}

func NewUpsertSessionRequest() (request *UpsertSessionRequest) {
    request = &UpsertSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "UpsertSession")
    
    
    return
}

func NewUpsertSessionResponse() (response *UpsertSessionResponse) {
    response = &UpsertSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpsertSession
// WAF session definition upsert API
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpsertSession(request *UpsertSessionRequest) (response *UpsertSessionResponse, err error) {
    return c.UpsertSessionWithContext(context.Background(), request)
}

// UpsertSession
// WAF session definition upsert API
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpsertSessionWithContext(ctx context.Context, request *UpsertSessionRequest) (response *UpsertSessionResponse, err error) {
    if request == nil {
        request = NewUpsertSessionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "waf", APIVersion, "UpsertSession")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpsertSession require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpsertSessionResponse()
    err = c.Send(request, response)
    return
}
