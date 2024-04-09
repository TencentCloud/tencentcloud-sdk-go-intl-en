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


func NewAttachPluginRequest() (request *AttachPluginRequest) {
    request = &AttachPluginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "AttachPlugin")
    
    
    return
}

func NewAttachPluginResponse() (response *AttachPluginResponse) {
    response = &AttachPluginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AttachPlugin
// This API is used to bind a plugin to an API.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDINPUTJSON = "FailedOperation.InvalidInputJSON"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  FAILEDOPERATION_TASKALREADYEXIST = "FailedOperation.TaskAlreadyExist"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_RETURNABLEEXCEPTION = "InternalError.ReturnableException"
//  INVALIDPARAMETER_DUPLICATEPLUGINCONFIG = "InvalidParameter.DuplicatePluginConfig"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_DUPLICATEPLUGINCONFIG = "InvalidParameterValue.DuplicatePluginConfig"
//  INVALIDPARAMETERVALUE_INVALIDENVSTATUS = "InvalidParameterValue.InvalidEnvStatus"
//  INVALIDPARAMETERVALUE_INVALIDSERVICECONFIG = "InvalidParameterValue.InvalidServiceConfig"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_TRAFFICCONTROL = "InvalidParameterValue.TrafficControl"
//  LIMITEXCEEDED_SERVICECOUNTFORPLUGINLIMITEXCEEDED = "LimitExceeded.ServiceCountForPluginLimitExceeded"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDPLUGIN = "ResourceNotFound.InvalidPlugin"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
//  UNSUPPORTEDOPERATION_ATTACHPLUGIN = "UnsupportedOperation.AttachPlugin"
//  UNSUPPORTEDOPERATION_BASICSERVICENOTALLOWATTACHPLUGIN = "UnsupportedOperation.BasicServiceNotAllowAttachPlugin"
//  UNSUPPORTEDOPERATION_INVALIDENDPOINTTYPE = "UnsupportedOperation.InvalidEndpointType"
//  UNSUPPORTEDOPERATION_UNPACKERROR = "UnsupportedOperation.UnpackError"
func (c *Client) AttachPlugin(request *AttachPluginRequest) (response *AttachPluginResponse, err error) {
    return c.AttachPluginWithContext(context.Background(), request)
}

// AttachPlugin
// This API is used to bind a plugin to an API.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDINPUTJSON = "FailedOperation.InvalidInputJSON"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  FAILEDOPERATION_TASKALREADYEXIST = "FailedOperation.TaskAlreadyExist"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_RETURNABLEEXCEPTION = "InternalError.ReturnableException"
//  INVALIDPARAMETER_DUPLICATEPLUGINCONFIG = "InvalidParameter.DuplicatePluginConfig"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_DUPLICATEPLUGINCONFIG = "InvalidParameterValue.DuplicatePluginConfig"
//  INVALIDPARAMETERVALUE_INVALIDENVSTATUS = "InvalidParameterValue.InvalidEnvStatus"
//  INVALIDPARAMETERVALUE_INVALIDSERVICECONFIG = "InvalidParameterValue.InvalidServiceConfig"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_TRAFFICCONTROL = "InvalidParameterValue.TrafficControl"
//  LIMITEXCEEDED_SERVICECOUNTFORPLUGINLIMITEXCEEDED = "LimitExceeded.ServiceCountForPluginLimitExceeded"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDPLUGIN = "ResourceNotFound.InvalidPlugin"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
//  UNSUPPORTEDOPERATION_ATTACHPLUGIN = "UnsupportedOperation.AttachPlugin"
//  UNSUPPORTEDOPERATION_BASICSERVICENOTALLOWATTACHPLUGIN = "UnsupportedOperation.BasicServiceNotAllowAttachPlugin"
//  UNSUPPORTEDOPERATION_INVALIDENDPOINTTYPE = "UnsupportedOperation.InvalidEndpointType"
//  UNSUPPORTEDOPERATION_UNPACKERROR = "UnsupportedOperation.UnpackError"
func (c *Client) AttachPluginWithContext(ctx context.Context, request *AttachPluginRequest) (response *AttachPluginResponse, err error) {
    if request == nil {
        request = NewAttachPluginRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AttachPlugin require credential")
    }

    request.SetContext(ctx)
    
    response = NewAttachPluginResponse()
    err = c.Send(request, response)
    return
}

func NewBindApiAppRequest() (request *BindApiAppRequest) {
    request = &BindApiAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "BindApiApp")
    
    
    return
}

func NewBindApiAppResponse() (response *BindApiAppResponse) {
    response = &BindApiAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindApiApp
// This API is used to bind an application to an API.
//
// error code that may be returned:
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  FAILEDOPERATION_APIINOPERATION = "FailedOperation.ApiInOperation"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDENV = "InvalidParameterValue.InvalidEnv"
//  LIMITEXCEEDED_APIAPPCOUNTLIMITEXCEEDED = "LimitExceeded.ApiAppCountLimitExceeded"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDAPIAPP = "ResourceNotFound.InvalidApiApp"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_RESOURCEASSOCIATED = "UnsupportedOperation.ResourceAssociated"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDBINDENVIRONMENT = "UnsupportedOperation.UnsupportedBindEnvironment"
func (c *Client) BindApiApp(request *BindApiAppRequest) (response *BindApiAppResponse, err error) {
    return c.BindApiAppWithContext(context.Background(), request)
}

// BindApiApp
// This API is used to bind an application to an API.
//
// error code that may be returned:
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  FAILEDOPERATION_APIINOPERATION = "FailedOperation.ApiInOperation"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDENV = "InvalidParameterValue.InvalidEnv"
//  LIMITEXCEEDED_APIAPPCOUNTLIMITEXCEEDED = "LimitExceeded.ApiAppCountLimitExceeded"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDAPIAPP = "ResourceNotFound.InvalidApiApp"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_RESOURCEASSOCIATED = "UnsupportedOperation.ResourceAssociated"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDBINDENVIRONMENT = "UnsupportedOperation.UnsupportedBindEnvironment"
func (c *Client) BindApiAppWithContext(ctx context.Context, request *BindApiAppRequest) (response *BindApiAppResponse, err error) {
    if request == nil {
        request = NewBindApiAppRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindApiApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindApiAppResponse()
    err = c.Send(request, response)
    return
}

func NewBindEnvironmentRequest() (request *BindEnvironmentRequest) {
    request = &BindEnvironmentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "BindEnvironment")
    
    
    return
}

func NewBindEnvironmentResponse() (response *BindEnvironmentResponse) {
    response = &BindEnvironmentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindEnvironment
// This API is used to bind a usage plan to a service or API.
//
// After you publish a service to an environment, if the API requires authentication and can be called only when it is bound to a usage plan, you can use this API to bind a usage plan to the specified environment.
//
// Currently, a usage plan can be bound to an API; however, under the same service, usage plans bound to a service and usage plans bound to an API cannot coexist. Therefore, in an environment to which a service-level usage plan has already been bound, please use the `DemoteServiceUsagePlan` API to degrade it.
//
// error code that may be returned:
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDAPIIDS = "InvalidParameterValue.InvalidApiIds"
//  LIMITEXCEEDED_REQUESTLIMITEXCEEDED = "LimitExceeded.RequestLimitExceeded"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDBINDENVIRONMENT = "UnsupportedOperation.UnsupportedBindEnvironment"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDUNBINDENVIRONMENT = "UnsupportedOperation.UnsupportedUnBindEnvironment"
func (c *Client) BindEnvironment(request *BindEnvironmentRequest) (response *BindEnvironmentResponse, err error) {
    return c.BindEnvironmentWithContext(context.Background(), request)
}

// BindEnvironment
// This API is used to bind a usage plan to a service or API.
//
// After you publish a service to an environment, if the API requires authentication and can be called only when it is bound to a usage plan, you can use this API to bind a usage plan to the specified environment.
//
// Currently, a usage plan can be bound to an API; however, under the same service, usage plans bound to a service and usage plans bound to an API cannot coexist. Therefore, in an environment to which a service-level usage plan has already been bound, please use the `DemoteServiceUsagePlan` API to degrade it.
//
// error code that may be returned:
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDAPIIDS = "InvalidParameterValue.InvalidApiIds"
//  LIMITEXCEEDED_REQUESTLIMITEXCEEDED = "LimitExceeded.RequestLimitExceeded"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDBINDENVIRONMENT = "UnsupportedOperation.UnsupportedBindEnvironment"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDUNBINDENVIRONMENT = "UnsupportedOperation.UnsupportedUnBindEnvironment"
func (c *Client) BindEnvironmentWithContext(ctx context.Context, request *BindEnvironmentRequest) (response *BindEnvironmentResponse, err error) {
    if request == nil {
        request = NewBindEnvironmentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindEnvironment require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindEnvironmentResponse()
    err = c.Send(request, response)
    return
}

func NewBindIPStrategyRequest() (request *BindIPStrategyRequest) {
    request = &BindIPStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "BindIPStrategy")
    
    
    return
}

func NewBindIPStrategyResponse() (response *BindIPStrategyResponse) {
    response = &BindIPStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindIPStrategy
// This API is used to bind an IP policy to an API.
//
// error code that may be returned:
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETERVALUE_INVALIDENVSTATUS = "InvalidParameterValue.InvalidEnvStatus"
//  LIMITEXCEEDED_REQUESTLIMITEXCEEDED = "LimitExceeded.RequestLimitExceeded"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDIPSTRATEGY = "ResourceNotFound.InvalidIPStrategy"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) BindIPStrategy(request *BindIPStrategyRequest) (response *BindIPStrategyResponse, err error) {
    return c.BindIPStrategyWithContext(context.Background(), request)
}

// BindIPStrategy
// This API is used to bind an IP policy to an API.
//
// error code that may be returned:
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETERVALUE_INVALIDENVSTATUS = "InvalidParameterValue.InvalidEnvStatus"
//  LIMITEXCEEDED_REQUESTLIMITEXCEEDED = "LimitExceeded.RequestLimitExceeded"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDIPSTRATEGY = "ResourceNotFound.InvalidIPStrategy"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) BindIPStrategyWithContext(ctx context.Context, request *BindIPStrategyRequest) (response *BindIPStrategyResponse, err error) {
    if request == nil {
        request = NewBindIPStrategyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindIPStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindIPStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewBindSecretIdsRequest() (request *BindSecretIdsRequest) {
    request = &BindSecretIdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "BindSecretIds")
    
    
    return
}

func NewBindSecretIdsResponse() (response *BindSecretIdsResponse) {
    response = &BindSecretIdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindSecretIds
// This API is used to bind a key to a usage plan.
//
// You can bind a key to a usage plan and bind the usage plan to an environment where a service is published, so that callers can use the key to call APIs in the service. You can use this API to bind a key to a usage plan.
//
// error code that may be returned:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETERVALUE_INVALIDACCESSKEYIDS = "InvalidParameterValue.InvalidAccessKeyIds"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  LIMITEXCEEDED_ACCESSKEYCOUNTINUSAGEPLANLIMITEXCEEDED = "LimitExceeded.AccessKeyCountInUsagePlanLimitExceeded"
//  RESOURCENOTFOUND_INVALIDACCESSKEYID = "ResourceNotFound.InvalidAccessKeyId"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
//  UNSUPPORTEDOPERATION_ALREADYBINDUSAGEPLAN = "UnsupportedOperation.AlreadyBindUsagePlan"
//  UNSUPPORTEDOPERATION_INVALIDSTATUS = "UnsupportedOperation.InvalidStatus"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDBINDAPIKEY = "UnsupportedOperation.UnsupportedBindApiKey"
func (c *Client) BindSecretIds(request *BindSecretIdsRequest) (response *BindSecretIdsResponse, err error) {
    return c.BindSecretIdsWithContext(context.Background(), request)
}

// BindSecretIds
// This API is used to bind a key to a usage plan.
//
// You can bind a key to a usage plan and bind the usage plan to an environment where a service is published, so that callers can use the key to call APIs in the service. You can use this API to bind a key to a usage plan.
//
// error code that may be returned:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETERVALUE_INVALIDACCESSKEYIDS = "InvalidParameterValue.InvalidAccessKeyIds"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  LIMITEXCEEDED_ACCESSKEYCOUNTINUSAGEPLANLIMITEXCEEDED = "LimitExceeded.AccessKeyCountInUsagePlanLimitExceeded"
//  RESOURCENOTFOUND_INVALIDACCESSKEYID = "ResourceNotFound.InvalidAccessKeyId"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
//  UNSUPPORTEDOPERATION_ALREADYBINDUSAGEPLAN = "UnsupportedOperation.AlreadyBindUsagePlan"
//  UNSUPPORTEDOPERATION_INVALIDSTATUS = "UnsupportedOperation.InvalidStatus"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDBINDAPIKEY = "UnsupportedOperation.UnsupportedBindApiKey"
func (c *Client) BindSecretIdsWithContext(ctx context.Context, request *BindSecretIdsRequest) (response *BindSecretIdsResponse, err error) {
    if request == nil {
        request = NewBindSecretIdsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindSecretIds require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindSecretIdsResponse()
    err = c.Send(request, response)
    return
}

func NewBindSubDomainRequest() (request *BindSubDomainRequest) {
    request = &BindSubDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "BindSubDomain")
    
    
    return
}

func NewBindSubDomainResponse() (response *BindSubDomainResponse) {
    response = &BindSubDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindSubDomain
// This API is used to bind a custom domain name to a service.
//
// Each service in API Gateway provides a default domain name for users to call. If you want to use your own domain name, you can bind a custom domain name to the target service. After getting the ICP filing and configuring the CNAME record between the custom and default domain names, you can directly call the custom domain name.
//
// error code that may be returned:
//  FAILEDOPERATION_CERTIFICATEIDENTERPRISEWAITSUBMIT = "FailedOperation.CertificateIdEnterpriseWaitSubmit"
//  FAILEDOPERATION_CERTIFICATEIDERROR = "FailedOperation.CertificateIdError"
//  FAILEDOPERATION_CERTIFICATEIDEXPIRED = "FailedOperation.CertificateIdExpired"
//  FAILEDOPERATION_CERTIFICATEIDINFOERROR = "FailedOperation.CertificateIdInfoError"
//  FAILEDOPERATION_CERTIFICATEIDUNDERVERIFY = "FailedOperation.CertificateIdUnderVerify"
//  FAILEDOPERATION_CERTIFICATEIDUNKNOWNERROR = "FailedOperation.CertificateIdUnknownError"
//  FAILEDOPERATION_CERTIFICATEIDVERIFYFAIL = "FailedOperation.CertificateIdVerifyFail"
//  FAILEDOPERATION_CERTIFICATEISNULL = "FailedOperation.CertificateIsNull"
//  FAILEDOPERATION_DEFINEMAPPINGENVIRONMENTERROR = "FailedOperation.DefineMappingEnvironmentError"
//  FAILEDOPERATION_DEFINEMAPPINGNOTNULL = "FailedOperation.DefineMappingNotNull"
//  FAILEDOPERATION_DEFINEMAPPINGPARAMREPEAT = "FailedOperation.DefineMappingParamRepeat"
//  FAILEDOPERATION_DEFINEMAPPINGPATHERROR = "FailedOperation.DefineMappingPathError"
//  FAILEDOPERATION_DOMAINALREADYBINDOTHERSERVICE = "FailedOperation.DomainAlreadyBindOtherService"
//  FAILEDOPERATION_DOMAINALREADYBINDSERVICE = "FailedOperation.DomainAlreadyBindService"
//  FAILEDOPERATION_DOMAININBLACKLIST = "FailedOperation.DomainInBlackList"
//  FAILEDOPERATION_DOMAINNEEDBEIAN = "FailedOperation.DomainNeedBeian"
//  FAILEDOPERATION_DOMAINRESOLVEERROR = "FailedOperation.DomainResolveError"
//  FAILEDOPERATION_FORMATERROR = "FailedOperation.FormatError"
//  FAILEDOPERATION_ISDEFAULTMAPPING = "FailedOperation.IsDefaultMapping"
//  FAILEDOPERATION_NETSUBDOMAINERROR = "FailedOperation.NetSubDomainError"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  FAILEDOPERATION_SUBDOMAINFORMATERROR = "FailedOperation.SubDomainFormatError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDPROCOTOL = "InvalidParameterValue.InvalidProcotol"
//  LIMITEXCEEDED_EXCEEDEDDEFINEMAPPINGLIMIT = "LimitExceeded.ExceededDefineMappingLimit"
//  LIMITEXCEEDED_EXCEEDEDDOMAINLIMIT = "LimitExceeded.ExceededDomainLimit"
//  LIMITEXCEEDED_REQUESTLIMITEXCEEDED = "LimitExceeded.RequestLimitExceeded"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_FORCEHTTPS = "UnsupportedOperation.ForceHttps"
//  UNSUPPORTEDOPERATION_INVALIDSERVICETRADE = "UnsupportedOperation.InvalidServiceTrade"
func (c *Client) BindSubDomain(request *BindSubDomainRequest) (response *BindSubDomainResponse, err error) {
    return c.BindSubDomainWithContext(context.Background(), request)
}

// BindSubDomain
// This API is used to bind a custom domain name to a service.
//
// Each service in API Gateway provides a default domain name for users to call. If you want to use your own domain name, you can bind a custom domain name to the target service. After getting the ICP filing and configuring the CNAME record between the custom and default domain names, you can directly call the custom domain name.
//
// error code that may be returned:
//  FAILEDOPERATION_CERTIFICATEIDENTERPRISEWAITSUBMIT = "FailedOperation.CertificateIdEnterpriseWaitSubmit"
//  FAILEDOPERATION_CERTIFICATEIDERROR = "FailedOperation.CertificateIdError"
//  FAILEDOPERATION_CERTIFICATEIDEXPIRED = "FailedOperation.CertificateIdExpired"
//  FAILEDOPERATION_CERTIFICATEIDINFOERROR = "FailedOperation.CertificateIdInfoError"
//  FAILEDOPERATION_CERTIFICATEIDUNDERVERIFY = "FailedOperation.CertificateIdUnderVerify"
//  FAILEDOPERATION_CERTIFICATEIDUNKNOWNERROR = "FailedOperation.CertificateIdUnknownError"
//  FAILEDOPERATION_CERTIFICATEIDVERIFYFAIL = "FailedOperation.CertificateIdVerifyFail"
//  FAILEDOPERATION_CERTIFICATEISNULL = "FailedOperation.CertificateIsNull"
//  FAILEDOPERATION_DEFINEMAPPINGENVIRONMENTERROR = "FailedOperation.DefineMappingEnvironmentError"
//  FAILEDOPERATION_DEFINEMAPPINGNOTNULL = "FailedOperation.DefineMappingNotNull"
//  FAILEDOPERATION_DEFINEMAPPINGPARAMREPEAT = "FailedOperation.DefineMappingParamRepeat"
//  FAILEDOPERATION_DEFINEMAPPINGPATHERROR = "FailedOperation.DefineMappingPathError"
//  FAILEDOPERATION_DOMAINALREADYBINDOTHERSERVICE = "FailedOperation.DomainAlreadyBindOtherService"
//  FAILEDOPERATION_DOMAINALREADYBINDSERVICE = "FailedOperation.DomainAlreadyBindService"
//  FAILEDOPERATION_DOMAININBLACKLIST = "FailedOperation.DomainInBlackList"
//  FAILEDOPERATION_DOMAINNEEDBEIAN = "FailedOperation.DomainNeedBeian"
//  FAILEDOPERATION_DOMAINRESOLVEERROR = "FailedOperation.DomainResolveError"
//  FAILEDOPERATION_FORMATERROR = "FailedOperation.FormatError"
//  FAILEDOPERATION_ISDEFAULTMAPPING = "FailedOperation.IsDefaultMapping"
//  FAILEDOPERATION_NETSUBDOMAINERROR = "FailedOperation.NetSubDomainError"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  FAILEDOPERATION_SUBDOMAINFORMATERROR = "FailedOperation.SubDomainFormatError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDPROCOTOL = "InvalidParameterValue.InvalidProcotol"
//  LIMITEXCEEDED_EXCEEDEDDEFINEMAPPINGLIMIT = "LimitExceeded.ExceededDefineMappingLimit"
//  LIMITEXCEEDED_EXCEEDEDDOMAINLIMIT = "LimitExceeded.ExceededDomainLimit"
//  LIMITEXCEEDED_REQUESTLIMITEXCEEDED = "LimitExceeded.RequestLimitExceeded"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_FORCEHTTPS = "UnsupportedOperation.ForceHttps"
//  UNSUPPORTEDOPERATION_INVALIDSERVICETRADE = "UnsupportedOperation.InvalidServiceTrade"
func (c *Client) BindSubDomainWithContext(ctx context.Context, request *BindSubDomainRequest) (response *BindSubDomainResponse, err error) {
    if request == nil {
        request = NewBindSubDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindSubDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindSubDomainResponse()
    err = c.Send(request, response)
    return
}

func NewBuildAPIDocRequest() (request *BuildAPIDocRequest) {
    request = &BuildAPIDocRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "BuildAPIDoc")
    
    
    return
}

func NewBuildAPIDocResponse() (response *BuildAPIDocResponse) {
    response = &BuildAPIDocResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BuildAPIDoc
// This API is used to build an API document.
//
// error code that may be returned:
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  FAILEDOPERATION_CODINGERROR = "FailedOperation.CodingError"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  INTERNALERROR_RETURNABLEEXCEPTION = "InternalError.ReturnableException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDAPIDOC = "ResourceNotFound.InvalidApiDoc"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) BuildAPIDoc(request *BuildAPIDocRequest) (response *BuildAPIDocResponse, err error) {
    return c.BuildAPIDocWithContext(context.Background(), request)
}

// BuildAPIDoc
// This API is used to build an API document.
//
// error code that may be returned:
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  FAILEDOPERATION_CODINGERROR = "FailedOperation.CodingError"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  INTERNALERROR_RETURNABLEEXCEPTION = "InternalError.ReturnableException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDAPIDOC = "ResourceNotFound.InvalidApiDoc"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) BuildAPIDocWithContext(ctx context.Context, request *BuildAPIDocRequest) (response *BuildAPIDocResponse, err error) {
    if request == nil {
        request = NewBuildAPIDocRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BuildAPIDoc require credential")
    }

    request.SetContext(ctx)
    
    response = NewBuildAPIDocResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAPIDocRequest() (request *CreateAPIDocRequest) {
    request = &CreateAPIDocRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "CreateAPIDoc")
    
    
    return
}

func NewCreateAPIDocResponse() (response *CreateAPIDocResponse) {
    response = &CreateAPIDocResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAPIDoc
// This API is used to create an API document.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSKEYEXIST = "FailedOperation.AccessKeyExist"
//  FAILEDOPERATION_APIBINDENVIRONMEN = "FailedOperation.ApiBindEnvironmen"
//  FAILEDOPERATION_APIBINDENVIRONMENT = "FailedOperation.ApiBindEnvironment"
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  FAILEDOPERATION_CODINGERROR = "FailedOperation.CodingError"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_APIDOCLIMITEXCEEDED = "LimitExceeded.APIDocLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INVALIDACCESSKEYID = "ResourceNotFound.InvalidAccessKeyId"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAPIDoc(request *CreateAPIDocRequest) (response *CreateAPIDocResponse, err error) {
    return c.CreateAPIDocWithContext(context.Background(), request)
}

// CreateAPIDoc
// This API is used to create an API document.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSKEYEXIST = "FailedOperation.AccessKeyExist"
//  FAILEDOPERATION_APIBINDENVIRONMEN = "FailedOperation.ApiBindEnvironmen"
//  FAILEDOPERATION_APIBINDENVIRONMENT = "FailedOperation.ApiBindEnvironment"
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  FAILEDOPERATION_CODINGERROR = "FailedOperation.CodingError"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_APIDOCLIMITEXCEEDED = "LimitExceeded.APIDocLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INVALIDACCESSKEYID = "ResourceNotFound.InvalidAccessKeyId"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAPIDocWithContext(ctx context.Context, request *CreateAPIDocRequest) (response *CreateAPIDocResponse, err error) {
    if request == nil {
        request = NewCreateAPIDocRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAPIDoc require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAPIDocResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApiRequest() (request *CreateApiRequest) {
    request = &CreateApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "CreateApi")
    
    
    return
}

func NewCreateApiResponse() (response *CreateApiResponse) {
    response = &CreateApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateApi
// This API is used to create an API. Before creating an API, you need to create a service, as each API belongs to a certain service.
//
// error code that may be returned:
//  FAILEDOPERATION_APIINOPERATION = "FailedOperation.ApiInOperation"
//  FAILEDOPERATION_BACKENDDOMAINERROR = "FailedOperation.BackendDomainError"
//  FAILEDOPERATION_DOMAINNEEDBEIAN = "FailedOperation.DomainNeedBeian"
//  FAILEDOPERATION_EIAMERROR = "FailedOperation.EIAMError"
//  FAILEDOPERATION_EBERROR = "FailedOperation.EbError"
//  FAILEDOPERATION_GETROLEERROR = "FailedOperation.GetRoleError"
//  FAILEDOPERATION_SCFERROR = "FailedOperation.ScfError"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_CAUTHEXCEPTION = "InternalError.CauthException"
//  INTERNALERROR_CLBEXCEPTION = "InternalError.ClbException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_SCFEXCEPTION = "InternalError.ScfException"
//  INTERNALERROR_TSFEXCEPTION = "InternalError.TsfException"
//  INTERNALERROR_VPCEXCEPTION = "InternalError.VpcException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_ILLEGALPROXYIP = "InvalidParameterValue.IllegalProxyIp"
//  INVALIDPARAMETERVALUE_INVALIDAPIBUSINESSTYPE = "InvalidParameterValue.InvalidApiBusinessType"
//  INVALIDPARAMETERVALUE_INVALIDAPIREQUESTCONFIG = "InvalidParameterValue.InvalidApiRequestConfig"
//  INVALIDPARAMETERVALUE_INVALIDAPITYPE = "InvalidParameterValue.InvalidApiType"
//  INVALIDPARAMETERVALUE_INVALIDBACKENDPATH = "InvalidParameterValue.InvalidBackendPath"
//  INVALIDPARAMETERVALUE_INVALIDCLB = "InvalidParameterValue.InvalidClb"
//  INVALIDPARAMETERVALUE_INVALIDIPADDRESS = "InvalidParameterValue.InvalidIPAddress"
//  INVALIDPARAMETERVALUE_INVALIDPORT = "InvalidParameterValue.InvalidPort"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICKEY = "InvalidParameterValue.InvalidPublicKey"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_INVALIDREQUESTPARAMETERS = "InvalidParameterValue.InvalidRequestParameters"
//  INVALIDPARAMETERVALUE_INVALIDSCFCONFIG = "InvalidParameterValue.InvalidScfConfig"
//  INVALIDPARAMETERVALUE_INVALIDSERVICECONFIG = "InvalidParameterValue.InvalidServiceConfig"
//  INVALIDPARAMETERVALUE_INVALIDSERVICEMOCKRETURNMESSAGE = "InvalidParameterValue.InvalidServiceMockReturnMessage"
//  INVALIDPARAMETERVALUE_INVALIDSERVICEPARAM = "InvalidParameterValue.InvalidServiceParam"
//  INVALIDPARAMETERVALUE_INVALIDSERVICEPARAMETERS = "InvalidParameterValue.InvalidServiceParameters"
//  INVALIDPARAMETERVALUE_INVALIDSERVICETYPE = "InvalidParameterValue.InvalidServiceType"
//  INVALIDPARAMETERVALUE_INVALIDTSFCONFIG = "InvalidParameterValue.InvalidTsfConfig"
//  INVALIDPARAMETERVALUE_INVALIDUPSTREAM = "InvalidParameterValue.InvalidUpstream"
//  INVALIDPARAMETERVALUE_INVALIDURL = "InvalidParameterValue.InvalidUrl"
//  INVALIDPARAMETERVALUE_INVALIDVPCCONFIG = "InvalidParameterValue.InvalidVpcConfig"
//  INVALIDPARAMETERVALUE_INVALIDWSMETHOD = "InvalidParameterValue.InvalidWSMethod"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NOTINOPTIONS = "InvalidParameterValue.NotInOptions"
//  INVALIDPARAMETERVALUE_PARAMETERNOTMATCH = "InvalidParameterValue.ParameterNotMatch"
//  INVALIDPARAMETERVALUE_RANGEEXCEEDED = "InvalidParameterValue.RangeExceeded"
//  LIMITEXCEEDED_APICOUNTLIMITEXCEEDED = "LimitExceeded.ApiCountLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INVALIDOAUTHAPI = "ResourceNotFound.InvalidOauthApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_BASICSERVICENOMOREAPI = "UnsupportedOperation.BasicServiceNoMoreApi"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_INVALIDENDPOINTTYPE = "UnsupportedOperation.InvalidEndpointType"
//  UNSUPPORTEDOPERATION_SERVICEEXIST = "UnsupportedOperation.ServiceExist"
//  UNSUPPORTEDOPERATION_UNPACKERROR = "UnsupportedOperation.UnpackError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDNETTYPE = "UnsupportedOperation.UnsupportedNetType"
func (c *Client) CreateApi(request *CreateApiRequest) (response *CreateApiResponse, err error) {
    return c.CreateApiWithContext(context.Background(), request)
}

// CreateApi
// This API is used to create an API. Before creating an API, you need to create a service, as each API belongs to a certain service.
//
// error code that may be returned:
//  FAILEDOPERATION_APIINOPERATION = "FailedOperation.ApiInOperation"
//  FAILEDOPERATION_BACKENDDOMAINERROR = "FailedOperation.BackendDomainError"
//  FAILEDOPERATION_DOMAINNEEDBEIAN = "FailedOperation.DomainNeedBeian"
//  FAILEDOPERATION_EIAMERROR = "FailedOperation.EIAMError"
//  FAILEDOPERATION_EBERROR = "FailedOperation.EbError"
//  FAILEDOPERATION_GETROLEERROR = "FailedOperation.GetRoleError"
//  FAILEDOPERATION_SCFERROR = "FailedOperation.ScfError"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_CAUTHEXCEPTION = "InternalError.CauthException"
//  INTERNALERROR_CLBEXCEPTION = "InternalError.ClbException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_SCFEXCEPTION = "InternalError.ScfException"
//  INTERNALERROR_TSFEXCEPTION = "InternalError.TsfException"
//  INTERNALERROR_VPCEXCEPTION = "InternalError.VpcException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_ILLEGALPROXYIP = "InvalidParameterValue.IllegalProxyIp"
//  INVALIDPARAMETERVALUE_INVALIDAPIBUSINESSTYPE = "InvalidParameterValue.InvalidApiBusinessType"
//  INVALIDPARAMETERVALUE_INVALIDAPIREQUESTCONFIG = "InvalidParameterValue.InvalidApiRequestConfig"
//  INVALIDPARAMETERVALUE_INVALIDAPITYPE = "InvalidParameterValue.InvalidApiType"
//  INVALIDPARAMETERVALUE_INVALIDBACKENDPATH = "InvalidParameterValue.InvalidBackendPath"
//  INVALIDPARAMETERVALUE_INVALIDCLB = "InvalidParameterValue.InvalidClb"
//  INVALIDPARAMETERVALUE_INVALIDIPADDRESS = "InvalidParameterValue.InvalidIPAddress"
//  INVALIDPARAMETERVALUE_INVALIDPORT = "InvalidParameterValue.InvalidPort"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICKEY = "InvalidParameterValue.InvalidPublicKey"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_INVALIDREQUESTPARAMETERS = "InvalidParameterValue.InvalidRequestParameters"
//  INVALIDPARAMETERVALUE_INVALIDSCFCONFIG = "InvalidParameterValue.InvalidScfConfig"
//  INVALIDPARAMETERVALUE_INVALIDSERVICECONFIG = "InvalidParameterValue.InvalidServiceConfig"
//  INVALIDPARAMETERVALUE_INVALIDSERVICEMOCKRETURNMESSAGE = "InvalidParameterValue.InvalidServiceMockReturnMessage"
//  INVALIDPARAMETERVALUE_INVALIDSERVICEPARAM = "InvalidParameterValue.InvalidServiceParam"
//  INVALIDPARAMETERVALUE_INVALIDSERVICEPARAMETERS = "InvalidParameterValue.InvalidServiceParameters"
//  INVALIDPARAMETERVALUE_INVALIDSERVICETYPE = "InvalidParameterValue.InvalidServiceType"
//  INVALIDPARAMETERVALUE_INVALIDTSFCONFIG = "InvalidParameterValue.InvalidTsfConfig"
//  INVALIDPARAMETERVALUE_INVALIDUPSTREAM = "InvalidParameterValue.InvalidUpstream"
//  INVALIDPARAMETERVALUE_INVALIDURL = "InvalidParameterValue.InvalidUrl"
//  INVALIDPARAMETERVALUE_INVALIDVPCCONFIG = "InvalidParameterValue.InvalidVpcConfig"
//  INVALIDPARAMETERVALUE_INVALIDWSMETHOD = "InvalidParameterValue.InvalidWSMethod"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NOTINOPTIONS = "InvalidParameterValue.NotInOptions"
//  INVALIDPARAMETERVALUE_PARAMETERNOTMATCH = "InvalidParameterValue.ParameterNotMatch"
//  INVALIDPARAMETERVALUE_RANGEEXCEEDED = "InvalidParameterValue.RangeExceeded"
//  LIMITEXCEEDED_APICOUNTLIMITEXCEEDED = "LimitExceeded.ApiCountLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INVALIDOAUTHAPI = "ResourceNotFound.InvalidOauthApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_BASICSERVICENOMOREAPI = "UnsupportedOperation.BasicServiceNoMoreApi"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_INVALIDENDPOINTTYPE = "UnsupportedOperation.InvalidEndpointType"
//  UNSUPPORTEDOPERATION_SERVICEEXIST = "UnsupportedOperation.ServiceExist"
//  UNSUPPORTEDOPERATION_UNPACKERROR = "UnsupportedOperation.UnpackError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDNETTYPE = "UnsupportedOperation.UnsupportedNetType"
func (c *Client) CreateApiWithContext(ctx context.Context, request *CreateApiRequest) (response *CreateApiResponse, err error) {
    if request == nil {
        request = NewCreateApiRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApi require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApiResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApiAppRequest() (request *CreateApiAppRequest) {
    request = &CreateApiAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "CreateApiApp")
    
    
    return
}

func NewCreateApiAppResponse() (response *CreateApiAppResponse) {
    response = &CreateApiAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateApiApp
// This API is used to create an application.
//
// error code that may be returned:
//  FAILEDOPERATION_ACCESSKEYEXIST = "FailedOperation.AccessKeyExist"
//  FAILEDOPERATION_TAGBINDSERVICEERROR = "FailedOperation.TagBindServiceError"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  LIMITEXCEEDED_APIAPPCOUNTLIMITEXCEEDED = "LimitExceeded.ApiAppCountLimitExceeded"
//  LIMITEXCEEDED_APIKEYCOUNTLIMITEXCEEDED = "LimitExceeded.ApiKeyCountLimitExceeded"
//  UNSUPPORTEDOPERATION_UNPACKERROR = "UnsupportedOperation.UnpackError"
func (c *Client) CreateApiApp(request *CreateApiAppRequest) (response *CreateApiAppResponse, err error) {
    return c.CreateApiAppWithContext(context.Background(), request)
}

// CreateApiApp
// This API is used to create an application.
//
// error code that may be returned:
//  FAILEDOPERATION_ACCESSKEYEXIST = "FailedOperation.AccessKeyExist"
//  FAILEDOPERATION_TAGBINDSERVICEERROR = "FailedOperation.TagBindServiceError"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  LIMITEXCEEDED_APIAPPCOUNTLIMITEXCEEDED = "LimitExceeded.ApiAppCountLimitExceeded"
//  LIMITEXCEEDED_APIKEYCOUNTLIMITEXCEEDED = "LimitExceeded.ApiKeyCountLimitExceeded"
//  UNSUPPORTEDOPERATION_UNPACKERROR = "UnsupportedOperation.UnpackError"
func (c *Client) CreateApiAppWithContext(ctx context.Context, request *CreateApiAppRequest) (response *CreateApiAppResponse, err error) {
    if request == nil {
        request = NewCreateApiAppRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApiApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApiAppResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApiKeyRequest() (request *CreateApiKeyRequest) {
    request = &CreateApiKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "CreateApiKey")
    
    
    return
}

func NewCreateApiKeyResponse() (response *CreateApiKeyResponse) {
    response = &CreateApiKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateApiKey
// This API is used to create an API key pair.
//
// error code that may be returned:
//  FAILEDOPERATION_ACCESSKEYEXIST = "FailedOperation.AccessKeyExist"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  LIMITEXCEEDED_APIKEYCOUNTLIMITEXCEEDED = "LimitExceeded.ApiKeyCountLimitExceeded"
//  UNAUTHORIZEDOPERATION_UNCERTIFIEDUSER = "UnauthorizedOperation.UncertifiedUser"
//  UNSUPPORTEDOPERATION_UNPACKERROR = "UnsupportedOperation.UnpackError"
func (c *Client) CreateApiKey(request *CreateApiKeyRequest) (response *CreateApiKeyResponse, err error) {
    return c.CreateApiKeyWithContext(context.Background(), request)
}

// CreateApiKey
// This API is used to create an API key pair.
//
// error code that may be returned:
//  FAILEDOPERATION_ACCESSKEYEXIST = "FailedOperation.AccessKeyExist"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  LIMITEXCEEDED_APIKEYCOUNTLIMITEXCEEDED = "LimitExceeded.ApiKeyCountLimitExceeded"
//  UNAUTHORIZEDOPERATION_UNCERTIFIEDUSER = "UnauthorizedOperation.UncertifiedUser"
//  UNSUPPORTEDOPERATION_UNPACKERROR = "UnsupportedOperation.UnpackError"
func (c *Client) CreateApiKeyWithContext(ctx context.Context, request *CreateApiKeyRequest) (response *CreateApiKeyResponse, err error) {
    if request == nil {
        request = NewCreateApiKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApiKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApiKeyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateIPStrategyRequest() (request *CreateIPStrategyRequest) {
    request = &CreateIPStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "CreateIPStrategy")
    
    
    return
}

func NewCreateIPStrategyResponse() (response *CreateIPStrategyResponse) {
    response = &CreateIPStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateIPStrategy
// This API is used to create a service IP policy.
//
// error code that may be returned:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  LIMITEXCEEDED_IPSTRATEGYLIMITEXCEEDED = "LimitExceeded.IpStrategyLimitExceeded"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) CreateIPStrategy(request *CreateIPStrategyRequest) (response *CreateIPStrategyResponse, err error) {
    return c.CreateIPStrategyWithContext(context.Background(), request)
}

// CreateIPStrategy
// This API is used to create a service IP policy.
//
// error code that may be returned:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  LIMITEXCEEDED_IPSTRATEGYLIMITEXCEEDED = "LimitExceeded.IpStrategyLimitExceeded"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) CreateIPStrategyWithContext(ctx context.Context, request *CreateIPStrategyRequest) (response *CreateIPStrategyResponse, err error) {
    if request == nil {
        request = NewCreateIPStrategyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateIPStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateIPStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePluginRequest() (request *CreatePluginRequest) {
    request = &CreatePluginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "CreatePlugin")
    
    
    return
}

func NewCreatePluginResponse() (response *CreatePluginResponse) {
    response = &CreatePluginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePlugin
// This API is used to create an API Gateway plugin.
//
// error code that may be returned:
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  FAILEDOPERATION_TAGBINDSERVICEERROR = "FailedOperation.TagBindServiceError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_TSFEXCEPTION = "InternalError.TsfException"
//  INTERNALERROR_VPCEXCEPTION = "InternalError.VpcException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_DUPLICATEPLUGINCONFIG = "InvalidParameterValue.DuplicatePluginConfig"
//  INVALIDPARAMETERVALUE_INVALIDAPIREQUESTCONFIG = "InvalidParameterValue.InvalidApiRequestConfig"
//  INVALIDPARAMETERVALUE_INVALIDBACKENDPATH = "InvalidParameterValue.InvalidBackendPath"
//  INVALIDPARAMETERVALUE_INVALIDCLB = "InvalidParameterValue.InvalidClb"
//  INVALIDPARAMETERVALUE_INVALIDCONDITION = "InvalidParameterValue.InvalidCondition"
//  INVALIDPARAMETERVALUE_INVALIDIPADDRESS = "InvalidParameterValue.InvalidIPAddress"
//  INVALIDPARAMETERVALUE_INVALIDPLUGINCONFIG = "InvalidParameterValue.InvalidPluginConfig"
//  INVALIDPARAMETERVALUE_INVALIDPORT = "InvalidParameterValue.InvalidPort"
//  INVALIDPARAMETERVALUE_INVALIDSCFCONFIG = "InvalidParameterValue.InvalidScfConfig"
//  INVALIDPARAMETERVALUE_INVALIDSERVICECONFIG = "InvalidParameterValue.InvalidServiceConfig"
//  INVALIDPARAMETERVALUE_INVALIDSERVICEPARAM = "InvalidParameterValue.InvalidServiceParam"
//  INVALIDPARAMETERVALUE_INVALIDSERVICETYPE = "InvalidParameterValue.InvalidServiceType"
//  INVALIDPARAMETERVALUE_INVALIDUPSTREAM = "InvalidParameterValue.InvalidUpstream"
//  INVALIDPARAMETERVALUE_INVALIDURL = "InvalidParameterValue.InvalidUrl"
//  INVALIDPARAMETERVALUE_INVALIDVPCCONFIG = "InvalidParameterValue.InvalidVpcConfig"
//  INVALIDPARAMETERVALUE_LENGTHEXCEEDED = "InvalidParameterValue.LengthExceeded"
//  INVALIDPARAMETERVALUE_NOTINOPTIONS = "InvalidParameterValue.NotInOptions"
//  INVALIDPARAMETERVALUE_PARAMETERVALUELIMITEXCEEDED = "InvalidParameterValue.ParameterValueLimitExceeded"
//  INVALIDPARAMETERVALUE_RANGEEXCEEDED = "InvalidParameterValue.RangeExceeded"
//  MISSINGPARAMETER_BACKENDSPECIFICPARAM = "MissingParameter.BackendSpecificParam"
//  MISSINGPARAMETER_PLUGINCONFIG = "MissingParameter.PluginConfig"
//  UNAUTHORIZEDOPERATION_UNCERTIFIEDUSER = "UnauthorizedOperation.UncertifiedUser"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_UNPACKERROR = "UnsupportedOperation.UnpackError"
func (c *Client) CreatePlugin(request *CreatePluginRequest) (response *CreatePluginResponse, err error) {
    return c.CreatePluginWithContext(context.Background(), request)
}

// CreatePlugin
// This API is used to create an API Gateway plugin.
//
// error code that may be returned:
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  FAILEDOPERATION_TAGBINDSERVICEERROR = "FailedOperation.TagBindServiceError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_TSFEXCEPTION = "InternalError.TsfException"
//  INTERNALERROR_VPCEXCEPTION = "InternalError.VpcException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_DUPLICATEPLUGINCONFIG = "InvalidParameterValue.DuplicatePluginConfig"
//  INVALIDPARAMETERVALUE_INVALIDAPIREQUESTCONFIG = "InvalidParameterValue.InvalidApiRequestConfig"
//  INVALIDPARAMETERVALUE_INVALIDBACKENDPATH = "InvalidParameterValue.InvalidBackendPath"
//  INVALIDPARAMETERVALUE_INVALIDCLB = "InvalidParameterValue.InvalidClb"
//  INVALIDPARAMETERVALUE_INVALIDCONDITION = "InvalidParameterValue.InvalidCondition"
//  INVALIDPARAMETERVALUE_INVALIDIPADDRESS = "InvalidParameterValue.InvalidIPAddress"
//  INVALIDPARAMETERVALUE_INVALIDPLUGINCONFIG = "InvalidParameterValue.InvalidPluginConfig"
//  INVALIDPARAMETERVALUE_INVALIDPORT = "InvalidParameterValue.InvalidPort"
//  INVALIDPARAMETERVALUE_INVALIDSCFCONFIG = "InvalidParameterValue.InvalidScfConfig"
//  INVALIDPARAMETERVALUE_INVALIDSERVICECONFIG = "InvalidParameterValue.InvalidServiceConfig"
//  INVALIDPARAMETERVALUE_INVALIDSERVICEPARAM = "InvalidParameterValue.InvalidServiceParam"
//  INVALIDPARAMETERVALUE_INVALIDSERVICETYPE = "InvalidParameterValue.InvalidServiceType"
//  INVALIDPARAMETERVALUE_INVALIDUPSTREAM = "InvalidParameterValue.InvalidUpstream"
//  INVALIDPARAMETERVALUE_INVALIDURL = "InvalidParameterValue.InvalidUrl"
//  INVALIDPARAMETERVALUE_INVALIDVPCCONFIG = "InvalidParameterValue.InvalidVpcConfig"
//  INVALIDPARAMETERVALUE_LENGTHEXCEEDED = "InvalidParameterValue.LengthExceeded"
//  INVALIDPARAMETERVALUE_NOTINOPTIONS = "InvalidParameterValue.NotInOptions"
//  INVALIDPARAMETERVALUE_PARAMETERVALUELIMITEXCEEDED = "InvalidParameterValue.ParameterValueLimitExceeded"
//  INVALIDPARAMETERVALUE_RANGEEXCEEDED = "InvalidParameterValue.RangeExceeded"
//  MISSINGPARAMETER_BACKENDSPECIFICPARAM = "MissingParameter.BackendSpecificParam"
//  MISSINGPARAMETER_PLUGINCONFIG = "MissingParameter.PluginConfig"
//  UNAUTHORIZEDOPERATION_UNCERTIFIEDUSER = "UnauthorizedOperation.UncertifiedUser"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_UNPACKERROR = "UnsupportedOperation.UnpackError"
func (c *Client) CreatePluginWithContext(ctx context.Context, request *CreatePluginRequest) (response *CreatePluginResponse, err error) {
    if request == nil {
        request = NewCreatePluginRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePlugin require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePluginResponse()
    err = c.Send(request, response)
    return
}

func NewCreateServiceRequest() (request *CreateServiceRequest) {
    request = &CreateServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "CreateService")
    
    
    return
}

func NewCreateServiceResponse() (response *CreateServiceResponse) {
    response = &CreateServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateService
// This API is used to create a service.
//
// A service is the biggest usage unit in API Gateway. Each service can contain multiple APIs and one default domain name for invocation. You can also bind your own custom domain name to a service. 
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTEXIST = "FailedOperation.InstanceNotExist"
//  FAILEDOPERATION_TAGBINDSERVICEERROR = "FailedOperation.TagBindServiceError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_CAUTHEXCEPTION = "InternalError.CauthException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_VPCEXCEPTION = "InternalError.VpcException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_INVALIDVPCCONFIG = "InvalidParameterValue.InvalidVpcConfig"
//  INVALIDPARAMETERVALUE_NOTINOPTIONS = "InvalidParameterValue.NotInOptions"
//  INVALIDPARAMETERVALUE_PARAMETERVALUELIMITEXCEEDED = "InvalidParameterValue.ParameterValueLimitExceeded"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_IPSTRATEGYLIMITEXCEEDED = "LimitExceeded.IpStrategyLimitExceeded"
//  LIMITEXCEEDED_SERVICECOUNTLIMITEXCEEDED = "LimitExceeded.ServiceCountLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNCERTIFIEDUSER = "UnauthorizedOperation.UncertifiedUser"
//  UNSUPPORTEDOPERATION_ACCOUNTARREARS = "UnsupportedOperation.AccountArrears"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_UNPACKERROR = "UnsupportedOperation.UnpackError"
func (c *Client) CreateService(request *CreateServiceRequest) (response *CreateServiceResponse, err error) {
    return c.CreateServiceWithContext(context.Background(), request)
}

// CreateService
// This API is used to create a service.
//
// A service is the biggest usage unit in API Gateway. Each service can contain multiple APIs and one default domain name for invocation. You can also bind your own custom domain name to a service. 
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTEXIST = "FailedOperation.InstanceNotExist"
//  FAILEDOPERATION_TAGBINDSERVICEERROR = "FailedOperation.TagBindServiceError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_CAUTHEXCEPTION = "InternalError.CauthException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_VPCEXCEPTION = "InternalError.VpcException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_INVALIDVPCCONFIG = "InvalidParameterValue.InvalidVpcConfig"
//  INVALIDPARAMETERVALUE_NOTINOPTIONS = "InvalidParameterValue.NotInOptions"
//  INVALIDPARAMETERVALUE_PARAMETERVALUELIMITEXCEEDED = "InvalidParameterValue.ParameterValueLimitExceeded"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_IPSTRATEGYLIMITEXCEEDED = "LimitExceeded.IpStrategyLimitExceeded"
//  LIMITEXCEEDED_SERVICECOUNTLIMITEXCEEDED = "LimitExceeded.ServiceCountLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNCERTIFIEDUSER = "UnauthorizedOperation.UncertifiedUser"
//  UNSUPPORTEDOPERATION_ACCOUNTARREARS = "UnsupportedOperation.AccountArrears"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_UNPACKERROR = "UnsupportedOperation.UnpackError"
func (c *Client) CreateServiceWithContext(ctx context.Context, request *CreateServiceRequest) (response *CreateServiceResponse, err error) {
    if request == nil {
        request = NewCreateServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateService require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateServiceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUpstreamRequest() (request *CreateUpstreamRequest) {
    request = &CreateUpstreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "CreateUpstream")
    
    
    return
}

func NewCreateUpstreamResponse() (response *CreateUpstreamResponse) {
    response = &CreateUpstreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateUpstream
// This API is used to create an upstream.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_OPERATEUPSTREAM = "FailedOperation.OperateUpstream"
//  FAILEDOPERATION_TAGBINDSERVICEERROR = "FailedOperation.TagBindServiceError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDVPCCONFIG = "InvalidParameterValue.InvalidVpcConfig"
//  INVALIDPARAMETERVALUE_NOTINOPTIONS = "InvalidParameterValue.NotInOptions"
//  INVALIDPARAMETERVALUE_RANGEEXCEEDED = "InvalidParameterValue.RangeExceeded"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateUpstream(request *CreateUpstreamRequest) (response *CreateUpstreamResponse, err error) {
    return c.CreateUpstreamWithContext(context.Background(), request)
}

// CreateUpstream
// This API is used to create an upstream.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_OPERATEUPSTREAM = "FailedOperation.OperateUpstream"
//  FAILEDOPERATION_TAGBINDSERVICEERROR = "FailedOperation.TagBindServiceError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDVPCCONFIG = "InvalidParameterValue.InvalidVpcConfig"
//  INVALIDPARAMETERVALUE_NOTINOPTIONS = "InvalidParameterValue.NotInOptions"
//  INVALIDPARAMETERVALUE_RANGEEXCEEDED = "InvalidParameterValue.RangeExceeded"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateUpstreamWithContext(ctx context.Context, request *CreateUpstreamRequest) (response *CreateUpstreamResponse, err error) {
    if request == nil {
        request = NewCreateUpstreamRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUpstream require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUpstreamResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUsagePlanRequest() (request *CreateUsagePlanRequest) {
    request = &CreateUsagePlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "CreateUsagePlan")
    
    
    return
}

func NewCreateUsagePlanResponse() (response *CreateUsagePlanResponse) {
    response = &CreateUsagePlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateUsagePlan
// This API is used to create a usage plan.
//
// To use API Gateway, you need to create a usage plan and bind it to a service environment.
//
// error code that may be returned:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDMAXREQUESTNUM = "InvalidParameterValue.InvalidMaxRequestNum"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  LIMITEXCEEDED_USAGEPLANLIMITEXCEEDED = "LimitExceeded.UsagePlanLimitExceeded"
//  UNSUPPORTEDOPERATION_UNPACKERROR = "UnsupportedOperation.UnpackError"
func (c *Client) CreateUsagePlan(request *CreateUsagePlanRequest) (response *CreateUsagePlanResponse, err error) {
    return c.CreateUsagePlanWithContext(context.Background(), request)
}

// CreateUsagePlan
// This API is used to create a usage plan.
//
// To use API Gateway, you need to create a usage plan and bind it to a service environment.
//
// error code that may be returned:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDMAXREQUESTNUM = "InvalidParameterValue.InvalidMaxRequestNum"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  LIMITEXCEEDED_USAGEPLANLIMITEXCEEDED = "LimitExceeded.UsagePlanLimitExceeded"
//  UNSUPPORTEDOPERATION_UNPACKERROR = "UnsupportedOperation.UnpackError"
func (c *Client) CreateUsagePlanWithContext(ctx context.Context, request *CreateUsagePlanRequest) (response *CreateUsagePlanResponse, err error) {
    if request == nil {
        request = NewCreateUsagePlanRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUsagePlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUsagePlanResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAPIDocRequest() (request *DeleteAPIDocRequest) {
    request = &DeleteAPIDocRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DeleteAPIDoc")
    
    
    return
}

func NewDeleteAPIDocResponse() (response *DeleteAPIDocResponse) {
    response = &DeleteAPIDocResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAPIDoc
// This API is used to delete an API document.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
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
func (c *Client) DeleteAPIDoc(request *DeleteAPIDocRequest) (response *DeleteAPIDocResponse, err error) {
    return c.DeleteAPIDocWithContext(context.Background(), request)
}

// DeleteAPIDoc
// This API is used to delete an API document.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
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
func (c *Client) DeleteAPIDocWithContext(ctx context.Context, request *DeleteAPIDocRequest) (response *DeleteAPIDocResponse, err error) {
    if request == nil {
        request = NewDeleteAPIDocRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAPIDoc require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAPIDocResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApiRequest() (request *DeleteApiRequest) {
    request = &DeleteApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DeleteApi")
    
    
    return
}

func NewDeleteApiResponse() (response *DeleteApiResponse) {
    response = &DeleteApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteApi
// This API is used to delete a created API.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APIBINDENVIRONMEN = "FailedOperation.ApiBindEnvironmen"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_INVALIDSERVICETRADE = "UnsupportedOperation.InvalidServiceTrade"
//  UNSUPPORTEDOPERATION_RESOURCEASSOCIATED = "UnsupportedOperation.ResourceAssociated"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDDELETEAPI = "UnsupportedOperation.UnsupportedDeleteApi"
func (c *Client) DeleteApi(request *DeleteApiRequest) (response *DeleteApiResponse, err error) {
    return c.DeleteApiWithContext(context.Background(), request)
}

// DeleteApi
// This API is used to delete a created API.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APIBINDENVIRONMEN = "FailedOperation.ApiBindEnvironmen"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_INVALIDSERVICETRADE = "UnsupportedOperation.InvalidServiceTrade"
//  UNSUPPORTEDOPERATION_RESOURCEASSOCIATED = "UnsupportedOperation.ResourceAssociated"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDDELETEAPI = "UnsupportedOperation.UnsupportedDeleteApi"
func (c *Client) DeleteApiWithContext(ctx context.Context, request *DeleteApiRequest) (response *DeleteApiResponse, err error) {
    if request == nil {
        request = NewDeleteApiRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApi require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteApiResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApiAppRequest() (request *DeleteApiAppRequest) {
    request = &DeleteApiAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DeleteApiApp")
    
    
    return
}

func NewDeleteApiAppResponse() (response *DeleteApiAppResponse) {
    response = &DeleteApiAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteApiApp
// This API is used to delete a created application.
//
// error code that may be returned:
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  FAILEDOPERATION_APIINOPERATION = "FailedOperation.ApiInOperation"
//  FAILEDOPERATION_GETROLEERROR = "FailedOperation.GetRoleError"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDAPIAPP = "ResourceNotFound.InvalidApiApp"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_INVALIDSERVICETRADE = "UnsupportedOperation.InvalidServiceTrade"
//  UNSUPPORTEDOPERATION_RESOURCEASSOCIATED = "UnsupportedOperation.ResourceAssociated"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDDELETEAPI = "UnsupportedOperation.UnsupportedDeleteApi"
func (c *Client) DeleteApiApp(request *DeleteApiAppRequest) (response *DeleteApiAppResponse, err error) {
    return c.DeleteApiAppWithContext(context.Background(), request)
}

// DeleteApiApp
// This API is used to delete a created application.
//
// error code that may be returned:
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  FAILEDOPERATION_APIINOPERATION = "FailedOperation.ApiInOperation"
//  FAILEDOPERATION_GETROLEERROR = "FailedOperation.GetRoleError"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDAPIAPP = "ResourceNotFound.InvalidApiApp"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_INVALIDSERVICETRADE = "UnsupportedOperation.InvalidServiceTrade"
//  UNSUPPORTEDOPERATION_RESOURCEASSOCIATED = "UnsupportedOperation.ResourceAssociated"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDDELETEAPI = "UnsupportedOperation.UnsupportedDeleteApi"
func (c *Client) DeleteApiAppWithContext(ctx context.Context, request *DeleteApiAppRequest) (response *DeleteApiAppResponse, err error) {
    if request == nil {
        request = NewDeleteApiAppRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApiApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteApiAppResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApiKeyRequest() (request *DeleteApiKeyRequest) {
    request = &DeleteApiKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DeleteApiKey")
    
    
    return
}

func NewDeleteApiKeyResponse() (response *DeleteApiKeyResponse) {
    response = &DeleteApiKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteApiKey
// This API is used to delete an API key pair.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDACCESSKEYID = "ResourceNotFound.InvalidAccessKeyId"
//  UNSUPPORTEDOPERATION_INVALIDSTATUS = "UnsupportedOperation.InvalidStatus"
//  UNSUPPORTEDOPERATION_RESOURCEISINUSE = "UnsupportedOperation.ResourceIsInUse"
func (c *Client) DeleteApiKey(request *DeleteApiKeyRequest) (response *DeleteApiKeyResponse, err error) {
    return c.DeleteApiKeyWithContext(context.Background(), request)
}

// DeleteApiKey
// This API is used to delete an API key pair.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDACCESSKEYID = "ResourceNotFound.InvalidAccessKeyId"
//  UNSUPPORTEDOPERATION_INVALIDSTATUS = "UnsupportedOperation.InvalidStatus"
//  UNSUPPORTEDOPERATION_RESOURCEISINUSE = "UnsupportedOperation.ResourceIsInUse"
func (c *Client) DeleteApiKeyWithContext(ctx context.Context, request *DeleteApiKeyRequest) (response *DeleteApiKeyResponse, err error) {
    if request == nil {
        request = NewDeleteApiKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApiKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteApiKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteIPStrategyRequest() (request *DeleteIPStrategyRequest) {
    request = &DeleteIPStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DeleteIPStrategy")
    
    
    return
}

func NewDeleteIPStrategyResponse() (response *DeleteIPStrategyResponse) {
    response = &DeleteIPStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteIPStrategy
// This API is used to delete a service IP policy.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INVALIDIPSTRATEGY = "ResourceNotFound.InvalidIPStrategy"
func (c *Client) DeleteIPStrategy(request *DeleteIPStrategyRequest) (response *DeleteIPStrategyResponse, err error) {
    return c.DeleteIPStrategyWithContext(context.Background(), request)
}

// DeleteIPStrategy
// This API is used to delete a service IP policy.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INVALIDIPSTRATEGY = "ResourceNotFound.InvalidIPStrategy"
func (c *Client) DeleteIPStrategyWithContext(ctx context.Context, request *DeleteIPStrategyRequest) (response *DeleteIPStrategyResponse, err error) {
    if request == nil {
        request = NewDeleteIPStrategyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteIPStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteIPStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePluginRequest() (request *DeletePluginRequest) {
    request = &DeletePluginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DeletePlugin")
    
    
    return
}

func NewDeletePluginResponse() (response *DeletePluginResponse) {
    response = &DeletePluginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeletePlugin
// This API is used to delete an API Gateway plugin.
//
// error code that may be returned:
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_RETURNABLEEXCEPTION = "InternalError.ReturnableException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDPLUGIN = "ResourceNotFound.InvalidPlugin"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_UNPACKERROR = "UnsupportedOperation.UnpackError"
func (c *Client) DeletePlugin(request *DeletePluginRequest) (response *DeletePluginResponse, err error) {
    return c.DeletePluginWithContext(context.Background(), request)
}

// DeletePlugin
// This API is used to delete an API Gateway plugin.
//
// error code that may be returned:
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_RETURNABLEEXCEPTION = "InternalError.ReturnableException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDPLUGIN = "ResourceNotFound.InvalidPlugin"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_UNPACKERROR = "UnsupportedOperation.UnpackError"
func (c *Client) DeletePluginWithContext(ctx context.Context, request *DeletePluginRequest) (response *DeletePluginResponse, err error) {
    if request == nil {
        request = NewDeletePluginRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePlugin require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePluginResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteServiceRequest() (request *DeleteServiceRequest) {
    request = &DeleteServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DeleteService")
    
    
    return
}

func NewDeleteServiceResponse() (response *DeleteServiceResponse) {
    response = &DeleteServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteService
// This API is used to delete a service in API Gateway.
//
// error code that may be returned:
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_APILISTNOTEMPTY = "UnsupportedOperation.ApiListNotEmpty"
//  UNSUPPORTEDOPERATION_EXISTINGONLINEENVIRONMENT = "UnsupportedOperation.ExistingOnlineEnvironment"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_TAGSNOTEMPTY = "UnsupportedOperation.TagsNotEmpty"
//  UNSUPPORTEDOPERATION_UNPACKERROR = "UnsupportedOperation.UnpackError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDDELETESERVICE = "UnsupportedOperation.UnsupportedDeleteService"
//  UNSUPPORTEDOPERATION_USAGEPLANINUSE = "UnsupportedOperation.UsagePlanInUse"
func (c *Client) DeleteService(request *DeleteServiceRequest) (response *DeleteServiceResponse, err error) {
    return c.DeleteServiceWithContext(context.Background(), request)
}

// DeleteService
// This API is used to delete a service in API Gateway.
//
// error code that may be returned:
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_APILISTNOTEMPTY = "UnsupportedOperation.ApiListNotEmpty"
//  UNSUPPORTEDOPERATION_EXISTINGONLINEENVIRONMENT = "UnsupportedOperation.ExistingOnlineEnvironment"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_TAGSNOTEMPTY = "UnsupportedOperation.TagsNotEmpty"
//  UNSUPPORTEDOPERATION_UNPACKERROR = "UnsupportedOperation.UnpackError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDDELETESERVICE = "UnsupportedOperation.UnsupportedDeleteService"
//  UNSUPPORTEDOPERATION_USAGEPLANINUSE = "UnsupportedOperation.UsagePlanInUse"
func (c *Client) DeleteServiceWithContext(ctx context.Context, request *DeleteServiceRequest) (response *DeleteServiceResponse, err error) {
    if request == nil {
        request = NewDeleteServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteService require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteServiceSubDomainMappingRequest() (request *DeleteServiceSubDomainMappingRequest) {
    request = &DeleteServiceSubDomainMappingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DeleteServiceSubDomainMapping")
    
    
    return
}

func NewDeleteServiceSubDomainMappingResponse() (response *DeleteServiceSubDomainMappingResponse) {
    response = &DeleteServiceSubDomainMappingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteServiceSubDomainMapping
// This API is used to delete a custom domain name mapping in a service environment.
//
// You can use this API if you use a custom domain name and custom mapping. Please note that if you delete all mappings in all environments, a failure will be returned when this API is called.
//
// error code that may be returned:
//  FAILEDOPERATION_DELETEPATHMAPPINGSETERROR = "FailedOperation.DeletePathMappingSetError"
//  FAILEDOPERATION_SERVICENOTEXIST = "FailedOperation.ServiceNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  UNSUPPORTEDOPERATION_INVALIDSERVICETRADE = "UnsupportedOperation.InvalidServiceTrade"
func (c *Client) DeleteServiceSubDomainMapping(request *DeleteServiceSubDomainMappingRequest) (response *DeleteServiceSubDomainMappingResponse, err error) {
    return c.DeleteServiceSubDomainMappingWithContext(context.Background(), request)
}

// DeleteServiceSubDomainMapping
// This API is used to delete a custom domain name mapping in a service environment.
//
// You can use this API if you use a custom domain name and custom mapping. Please note that if you delete all mappings in all environments, a failure will be returned when this API is called.
//
// error code that may be returned:
//  FAILEDOPERATION_DELETEPATHMAPPINGSETERROR = "FailedOperation.DeletePathMappingSetError"
//  FAILEDOPERATION_SERVICENOTEXIST = "FailedOperation.ServiceNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  UNSUPPORTEDOPERATION_INVALIDSERVICETRADE = "UnsupportedOperation.InvalidServiceTrade"
func (c *Client) DeleteServiceSubDomainMappingWithContext(ctx context.Context, request *DeleteServiceSubDomainMappingRequest) (response *DeleteServiceSubDomainMappingResponse, err error) {
    if request == nil {
        request = NewDeleteServiceSubDomainMappingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteServiceSubDomainMapping require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteServiceSubDomainMappingResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUpstreamRequest() (request *DeleteUpstreamRequest) {
    request = &DeleteUpstreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DeleteUpstream")
    
    
    return
}

func NewDeleteUpstreamResponse() (response *DeleteUpstreamResponse) {
    response = &DeleteUpstreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteUpstream
// This API is used to delete an upstream. Note that you can only delete an upstream when it’s not bound with any APIs.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDDELETEUPSTREAM = "UnsupportedOperation.UnsupportedDeleteUpstream"
func (c *Client) DeleteUpstream(request *DeleteUpstreamRequest) (response *DeleteUpstreamResponse, err error) {
    return c.DeleteUpstreamWithContext(context.Background(), request)
}

// DeleteUpstream
// This API is used to delete an upstream. Note that you can only delete an upstream when it’s not bound with any APIs.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDDELETEUPSTREAM = "UnsupportedOperation.UnsupportedDeleteUpstream"
func (c *Client) DeleteUpstreamWithContext(ctx context.Context, request *DeleteUpstreamRequest) (response *DeleteUpstreamResponse, err error) {
    if request == nil {
        request = NewDeleteUpstreamRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUpstream require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUpstreamResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUsagePlanRequest() (request *DeleteUsagePlanRequest) {
    request = &DeleteUsagePlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DeleteUsagePlan")
    
    
    return
}

func NewDeleteUsagePlanResponse() (response *DeleteUsagePlanResponse) {
    response = &DeleteUsagePlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteUsagePlan
// This API is used to delete a usage plan.
//
// error code that may be returned:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
//  UNSUPPORTEDOPERATION_USAGEPLANINUSE = "UnsupportedOperation.UsagePlanInUse"
func (c *Client) DeleteUsagePlan(request *DeleteUsagePlanRequest) (response *DeleteUsagePlanResponse, err error) {
    return c.DeleteUsagePlanWithContext(context.Background(), request)
}

// DeleteUsagePlan
// This API is used to delete a usage plan.
//
// error code that may be returned:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
//  UNSUPPORTEDOPERATION_USAGEPLANINUSE = "UnsupportedOperation.UsagePlanInUse"
func (c *Client) DeleteUsagePlanWithContext(ctx context.Context, request *DeleteUsagePlanRequest) (response *DeleteUsagePlanResponse, err error) {
    if request == nil {
        request = NewDeleteUsagePlanRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUsagePlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUsagePlanResponse()
    err = c.Send(request, response)
    return
}

func NewDemoteServiceUsagePlanRequest() (request *DemoteServiceUsagePlanRequest) {
    request = &DemoteServiceUsagePlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DemoteServiceUsagePlan")
    
    
    return
}

func NewDemoteServiceUsagePlanResponse() (response *DemoteServiceUsagePlanResponse) {
    response = &DemoteServiceUsagePlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DemoteServiceUsagePlan
// This API is used to degrade a usage plan of a service in an environment to the API level.
//
// This operation will be denied if there are no APIs under the service.
//
// This operation will also be denied if the current environment has not been published.
//
// error code that may be returned:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
//  UNSUPPORTEDOPERATION_NOUSAGEPLANENV = "UnsupportedOperation.NoUsagePlanEnv"
func (c *Client) DemoteServiceUsagePlan(request *DemoteServiceUsagePlanRequest) (response *DemoteServiceUsagePlanResponse, err error) {
    return c.DemoteServiceUsagePlanWithContext(context.Background(), request)
}

// DemoteServiceUsagePlan
// This API is used to degrade a usage plan of a service in an environment to the API level.
//
// This operation will be denied if there are no APIs under the service.
//
// This operation will also be denied if the current environment has not been published.
//
// error code that may be returned:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
//  UNSUPPORTEDOPERATION_NOUSAGEPLANENV = "UnsupportedOperation.NoUsagePlanEnv"
func (c *Client) DemoteServiceUsagePlanWithContext(ctx context.Context, request *DemoteServiceUsagePlanRequest) (response *DemoteServiceUsagePlanResponse, err error) {
    if request == nil {
        request = NewDemoteServiceUsagePlanRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DemoteServiceUsagePlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewDemoteServiceUsagePlanResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAPIDocDetailRequest() (request *DescribeAPIDocDetailRequest) {
    request = &DescribeAPIDocDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeAPIDocDetail")
    
    
    return
}

func NewDescribeAPIDocDetailResponse() (response *DescribeAPIDocDetailResponse) {
    response = &DescribeAPIDocDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAPIDocDetail
// This API is used to query the details of an API document.
//
// error code that may be returned:
//  FAILEDOPERATION_CODINGERROR = "FailedOperation.CodingError"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDAPIDOC = "ResourceNotFound.InvalidApiDoc"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DescribeAPIDocDetail(request *DescribeAPIDocDetailRequest) (response *DescribeAPIDocDetailResponse, err error) {
    return c.DescribeAPIDocDetailWithContext(context.Background(), request)
}

// DescribeAPIDocDetail
// This API is used to query the details of an API document.
//
// error code that may be returned:
//  FAILEDOPERATION_CODINGERROR = "FailedOperation.CodingError"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDAPIDOC = "ResourceNotFound.InvalidApiDoc"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DescribeAPIDocDetailWithContext(ctx context.Context, request *DescribeAPIDocDetailRequest) (response *DescribeAPIDocDetailResponse, err error) {
    if request == nil {
        request = NewDescribeAPIDocDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAPIDocDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAPIDocDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAPIDocsRequest() (request *DescribeAPIDocsRequest) {
    request = &DescribeAPIDocsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeAPIDocs")
    
    
    return
}

func NewDescribeAPIDocsResponse() (response *DescribeAPIDocsResponse) {
    response = &DescribeAPIDocsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAPIDocs
// This API is used to query the list of API documents.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DescribeAPIDocs(request *DescribeAPIDocsRequest) (response *DescribeAPIDocsResponse, err error) {
    return c.DescribeAPIDocsWithContext(context.Background(), request)
}

// DescribeAPIDocs
// This API is used to query the list of API documents.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DescribeAPIDocsWithContext(ctx context.Context, request *DescribeAPIDocsRequest) (response *DescribeAPIDocsResponse, err error) {
    if request == nil {
        request = NewDescribeAPIDocsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAPIDocs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAPIDocsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAllPluginApisRequest() (request *DescribeAllPluginApisRequest) {
    request = &DescribeAllPluginApisRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeAllPluginApis")
    
    
    return
}

func NewDescribeAllPluginApisResponse() (response *DescribeAllPluginApisResponse) {
    response = &DescribeAllPluginApisResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAllPluginApis
// This API is used to list all APIs that can use this plugin, no matter whether the API is bound with the plugin.
//
// error code that may be returned:
//  FAILEDOPERATION_FORMATERROR = "FailedOperation.FormatError"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  RESOURCENOTFOUND_INVALIDPLUGIN = "ResourceNotFound.InvalidPlugin"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
func (c *Client) DescribeAllPluginApis(request *DescribeAllPluginApisRequest) (response *DescribeAllPluginApisResponse, err error) {
    return c.DescribeAllPluginApisWithContext(context.Background(), request)
}

// DescribeAllPluginApis
// This API is used to list all APIs that can use this plugin, no matter whether the API is bound with the plugin.
//
// error code that may be returned:
//  FAILEDOPERATION_FORMATERROR = "FailedOperation.FormatError"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  RESOURCENOTFOUND_INVALIDPLUGIN = "ResourceNotFound.InvalidPlugin"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
func (c *Client) DescribeAllPluginApisWithContext(ctx context.Context, request *DescribeAllPluginApisRequest) (response *DescribeAllPluginApisResponse, err error) {
    if request == nil {
        request = NewDescribeAllPluginApisRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAllPluginApis require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAllPluginApisResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiRequest() (request *DescribeApiRequest) {
    request = &DescribeApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApi")
    
    
    return
}

func NewDescribeApiResponse() (response *DescribeApiResponse) {
    response = &DescribeApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApi
// This API (`DescribeApi`) is used to query the details of the APIs users manage via Tencent Cloud API Gateway.
//
// error code that may be returned:
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  FAILEDOPERATION_EBERROR = "FailedOperation.EbError"
//  FAILEDOPERATION_SERVICENOTEXIST = "FailedOperation.ServiceNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBEXCEPTION = "InternalError.DbException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_INVALIDUPSTREAM = "InvalidParameterValue.InvalidUpstream"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DescribeApi(request *DescribeApiRequest) (response *DescribeApiResponse, err error) {
    return c.DescribeApiWithContext(context.Background(), request)
}

// DescribeApi
// This API (`DescribeApi`) is used to query the details of the APIs users manage via Tencent Cloud API Gateway.
//
// error code that may be returned:
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  FAILEDOPERATION_EBERROR = "FailedOperation.EbError"
//  FAILEDOPERATION_SERVICENOTEXIST = "FailedOperation.ServiceNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBEXCEPTION = "InternalError.DbException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_INVALIDUPSTREAM = "InvalidParameterValue.InvalidUpstream"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DescribeApiWithContext(ctx context.Context, request *DescribeApiRequest) (response *DescribeApiResponse, err error) {
    if request == nil {
        request = NewDescribeApiRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApi require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApiResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiAppRequest() (request *DescribeApiAppRequest) {
    request = &DescribeApiAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApiApp")
    
    
    return
}

func NewDescribeApiAppResponse() (response *DescribeApiAppResponse) {
    response = &DescribeApiAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApiApp
// This API is used to search for an application by application ID.
//
// error code that may be returned:
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
func (c *Client) DescribeApiApp(request *DescribeApiAppRequest) (response *DescribeApiAppResponse, err error) {
    return c.DescribeApiAppWithContext(context.Background(), request)
}

// DescribeApiApp
// This API is used to search for an application by application ID.
//
// error code that may be returned:
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
func (c *Client) DescribeApiAppWithContext(ctx context.Context, request *DescribeApiAppRequest) (response *DescribeApiAppResponse, err error) {
    if request == nil {
        request = NewDescribeApiAppRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApiApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApiAppResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiAppBindApisStatusRequest() (request *DescribeApiAppBindApisStatusRequest) {
    request = &DescribeApiAppBindApisStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApiAppBindApisStatus")
    
    
    return
}

func NewDescribeApiAppBindApisStatusResponse() (response *DescribeApiAppBindApisStatusResponse) {
    response = &DescribeApiAppBindApisStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApiAppBindApisStatus
// This API is used to query the list of APIs bound to an application.
//
// error code that may be returned:
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
func (c *Client) DescribeApiAppBindApisStatus(request *DescribeApiAppBindApisStatusRequest) (response *DescribeApiAppBindApisStatusResponse, err error) {
    return c.DescribeApiAppBindApisStatusWithContext(context.Background(), request)
}

// DescribeApiAppBindApisStatus
// This API is used to query the list of APIs bound to an application.
//
// error code that may be returned:
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
func (c *Client) DescribeApiAppBindApisStatusWithContext(ctx context.Context, request *DescribeApiAppBindApisStatusRequest) (response *DescribeApiAppBindApisStatusResponse, err error) {
    if request == nil {
        request = NewDescribeApiAppBindApisStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApiAppBindApisStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApiAppBindApisStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiAppsStatusRequest() (request *DescribeApiAppsStatusRequest) {
    request = &DescribeApiAppsStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApiAppsStatus")
    
    
    return
}

func NewDescribeApiAppsStatusResponse() (response *DescribeApiAppsStatusResponse) {
    response = &DescribeApiAppsStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApiAppsStatus
// This API is used to query the application list.
//
// error code that may be returned:
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_CAMEXCEPTION = "InternalError.CamException"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  INVALIDPARAMETERVALUE_INVALIDTAGVALUES = "InvalidParameterValue.InvalidTagValues"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_REQUESTPOSTERROR = "UnsupportedOperation.RequestPostError"
func (c *Client) DescribeApiAppsStatus(request *DescribeApiAppsStatusRequest) (response *DescribeApiAppsStatusResponse, err error) {
    return c.DescribeApiAppsStatusWithContext(context.Background(), request)
}

// DescribeApiAppsStatus
// This API is used to query the application list.
//
// error code that may be returned:
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_CAMEXCEPTION = "InternalError.CamException"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  INVALIDPARAMETERVALUE_INVALIDTAGVALUES = "InvalidParameterValue.InvalidTagValues"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_REQUESTPOSTERROR = "UnsupportedOperation.RequestPostError"
func (c *Client) DescribeApiAppsStatusWithContext(ctx context.Context, request *DescribeApiAppsStatusRequest) (response *DescribeApiAppsStatusResponse, err error) {
    if request == nil {
        request = NewDescribeApiAppsStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApiAppsStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApiAppsStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiBindApiAppsStatusRequest() (request *DescribeApiBindApiAppsStatusRequest) {
    request = &DescribeApiBindApiAppsStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApiBindApiAppsStatus")
    
    
    return
}

func NewDescribeApiBindApiAppsStatusResponse() (response *DescribeApiBindApiAppsStatusResponse) {
    response = &DescribeApiBindApiAppsStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApiBindApiAppsStatus
// This API is used to query the list of applications bound to an API.
//
// error code that may be returned:
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
func (c *Client) DescribeApiBindApiAppsStatus(request *DescribeApiBindApiAppsStatusRequest) (response *DescribeApiBindApiAppsStatusResponse, err error) {
    return c.DescribeApiBindApiAppsStatusWithContext(context.Background(), request)
}

// DescribeApiBindApiAppsStatus
// This API is used to query the list of applications bound to an API.
//
// error code that may be returned:
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
func (c *Client) DescribeApiBindApiAppsStatusWithContext(ctx context.Context, request *DescribeApiBindApiAppsStatusRequest) (response *DescribeApiBindApiAppsStatusResponse, err error) {
    if request == nil {
        request = NewDescribeApiBindApiAppsStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApiBindApiAppsStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApiBindApiAppsStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiEnvironmentStrategyRequest() (request *DescribeApiEnvironmentStrategyRequest) {
    request = &DescribeApiEnvironmentStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApiEnvironmentStrategy")
    
    
    return
}

func NewDescribeApiEnvironmentStrategyResponse() (response *DescribeApiEnvironmentStrategyResponse) {
    response = &DescribeApiEnvironmentStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApiEnvironmentStrategy
// This API is used to display the throttling policies bound to an API.
//
// error code that may be returned:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeApiEnvironmentStrategy(request *DescribeApiEnvironmentStrategyRequest) (response *DescribeApiEnvironmentStrategyResponse, err error) {
    return c.DescribeApiEnvironmentStrategyWithContext(context.Background(), request)
}

// DescribeApiEnvironmentStrategy
// This API is used to display the throttling policies bound to an API.
//
// error code that may be returned:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeApiEnvironmentStrategyWithContext(ctx context.Context, request *DescribeApiEnvironmentStrategyRequest) (response *DescribeApiEnvironmentStrategyResponse, err error) {
    if request == nil {
        request = NewDescribeApiEnvironmentStrategyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApiEnvironmentStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApiEnvironmentStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiForApiAppRequest() (request *DescribeApiForApiAppRequest) {
    request = &DescribeApiForApiAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApiForApiApp")
    
    
    return
}

func NewDescribeApiForApiAppResponse() (response *DescribeApiForApiAppResponse) {
    response = &DescribeApiForApiAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApiForApiApp
// This API is used to query the details of an API deployed at API Gateway.
//
// error code that may be returned:
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  FAILEDOPERATION_SERVICENOTEXIST = "FailedOperation.ServiceNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_RETURNABLEEXCEPTION = "InternalError.ReturnableException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_UNPACKERROR = "UnsupportedOperation.UnpackError"
func (c *Client) DescribeApiForApiApp(request *DescribeApiForApiAppRequest) (response *DescribeApiForApiAppResponse, err error) {
    return c.DescribeApiForApiAppWithContext(context.Background(), request)
}

// DescribeApiForApiApp
// This API is used to query the details of an API deployed at API Gateway.
//
// error code that may be returned:
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  FAILEDOPERATION_SERVICENOTEXIST = "FailedOperation.ServiceNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_RETURNABLEEXCEPTION = "InternalError.ReturnableException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_UNPACKERROR = "UnsupportedOperation.UnpackError"
func (c *Client) DescribeApiForApiAppWithContext(ctx context.Context, request *DescribeApiForApiAppRequest) (response *DescribeApiForApiAppResponse, err error) {
    if request == nil {
        request = NewDescribeApiForApiAppRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApiForApiApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApiForApiAppResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiKeyRequest() (request *DescribeApiKeyRequest) {
    request = &DescribeApiKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApiKey")
    
    
    return
}

func NewDescribeApiKeyResponse() (response *DescribeApiKeyResponse) {
    response = &DescribeApiKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApiKey
// This API is used to query the details of a key.
//
// After creating an API key, you can query its details by using this API.
//
// error code that may be returned:
//  FAILEDOPERATION_CAMEXCEPTION = "FailedOperation.CamException"
//  RESOURCENOTFOUND_INVALIDACCESSKEYID = "ResourceNotFound.InvalidAccessKeyId"
func (c *Client) DescribeApiKey(request *DescribeApiKeyRequest) (response *DescribeApiKeyResponse, err error) {
    return c.DescribeApiKeyWithContext(context.Background(), request)
}

// DescribeApiKey
// This API is used to query the details of a key.
//
// After creating an API key, you can query its details by using this API.
//
// error code that may be returned:
//  FAILEDOPERATION_CAMEXCEPTION = "FailedOperation.CamException"
//  RESOURCENOTFOUND_INVALIDACCESSKEYID = "ResourceNotFound.InvalidAccessKeyId"
func (c *Client) DescribeApiKeyWithContext(ctx context.Context, request *DescribeApiKeyRequest) (response *DescribeApiKeyResponse, err error) {
    if request == nil {
        request = NewDescribeApiKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApiKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApiKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiKeysStatusRequest() (request *DescribeApiKeysStatusRequest) {
    request = &DescribeApiKeysStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApiKeysStatus")
    
    
    return
}

func NewDescribeApiKeysStatusResponse() (response *DescribeApiKeysStatusResponse) {
    response = &DescribeApiKeysStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApiKeysStatus
// This API is used to query the information of one or more API keys.
//
//  
//
// error code that may be returned:
//  FAILEDOPERATION_CAMEXCEPTION = "FailedOperation.CamException"
//  INTERNALERROR_CAMEXCEPTION = "InternalError.CamException"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
func (c *Client) DescribeApiKeysStatus(request *DescribeApiKeysStatusRequest) (response *DescribeApiKeysStatusResponse, err error) {
    return c.DescribeApiKeysStatusWithContext(context.Background(), request)
}

// DescribeApiKeysStatus
// This API is used to query the information of one or more API keys.
//
//  
//
// error code that may be returned:
//  FAILEDOPERATION_CAMEXCEPTION = "FailedOperation.CamException"
//  INTERNALERROR_CAMEXCEPTION = "InternalError.CamException"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
func (c *Client) DescribeApiKeysStatusWithContext(ctx context.Context, request *DescribeApiKeysStatusRequest) (response *DescribeApiKeysStatusResponse, err error) {
    if request == nil {
        request = NewDescribeApiKeysStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApiKeysStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApiKeysStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiUsagePlanRequest() (request *DescribeApiUsagePlanRequest) {
    request = &DescribeApiUsagePlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApiUsagePlan")
    
    
    return
}

func NewDescribeApiUsagePlanResponse() (response *DescribeApiUsagePlanResponse) {
    response = &DescribeApiUsagePlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApiUsagePlan
// This API is used to query the details of API usage plans in a service.
//
// To make authentication and throttling for a service take effect, you need to bind a usage plan to it. This API is used to query all usage plans bound to a service and APIs under it.
//
// error code that may be returned:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeApiUsagePlan(request *DescribeApiUsagePlanRequest) (response *DescribeApiUsagePlanResponse, err error) {
    return c.DescribeApiUsagePlanWithContext(context.Background(), request)
}

// DescribeApiUsagePlan
// This API is used to query the details of API usage plans in a service.
//
// To make authentication and throttling for a service take effect, you need to bind a usage plan to it. This API is used to query all usage plans bound to a service and APIs under it.
//
// error code that may be returned:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeApiUsagePlanWithContext(ctx context.Context, request *DescribeApiUsagePlanRequest) (response *DescribeApiUsagePlanResponse, err error) {
    if request == nil {
        request = NewDescribeApiUsagePlanRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApiUsagePlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApiUsagePlanResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApisStatusRequest() (request *DescribeApisStatusRequest) {
    request = &DescribeApisStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApisStatus")
    
    
    return
}

func NewDescribeApisStatusResponse() (response *DescribeApisStatusResponse) {
    response = &DescribeApisStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApisStatus
// This API is used to view a certain API or the list of all APIs under a service and relevant information.
//
// error code that may be returned:
//  FAILEDOPERATION_SERVICENOTEXIST = "FailedOperation.ServiceNotExist"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_REQUESTPOSTERROR = "UnsupportedOperation.RequestPostError"
func (c *Client) DescribeApisStatus(request *DescribeApisStatusRequest) (response *DescribeApisStatusResponse, err error) {
    return c.DescribeApisStatusWithContext(context.Background(), request)
}

// DescribeApisStatus
// This API is used to view a certain API or the list of all APIs under a service and relevant information.
//
// error code that may be returned:
//  FAILEDOPERATION_SERVICENOTEXIST = "FailedOperation.ServiceNotExist"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_REQUESTPOSTERROR = "UnsupportedOperation.RequestPostError"
func (c *Client) DescribeApisStatusWithContext(ctx context.Context, request *DescribeApisStatusRequest) (response *DescribeApisStatusResponse, err error) {
    if request == nil {
        request = NewDescribeApisStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApisStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApisStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIPStrategyRequest() (request *DescribeIPStrategyRequest) {
    request = &DescribeIPStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeIPStrategy")
    
    
    return
}

func NewDescribeIPStrategyResponse() (response *DescribeIPStrategyResponse) {
    response = &DescribeIPStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIPStrategy
// This API is used to query IP policy details.
//
// error code that may be returned:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  RESOURCENOTFOUND_INVALIDIPSTRATEGY = "ResourceNotFound.InvalidIPStrategy"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeIPStrategy(request *DescribeIPStrategyRequest) (response *DescribeIPStrategyResponse, err error) {
    return c.DescribeIPStrategyWithContext(context.Background(), request)
}

// DescribeIPStrategy
// This API is used to query IP policy details.
//
// error code that may be returned:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  RESOURCENOTFOUND_INVALIDIPSTRATEGY = "ResourceNotFound.InvalidIPStrategy"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeIPStrategyWithContext(ctx context.Context, request *DescribeIPStrategyRequest) (response *DescribeIPStrategyResponse, err error) {
    if request == nil {
        request = NewDescribeIPStrategyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIPStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIPStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIPStrategyApisStatusRequest() (request *DescribeIPStrategyApisStatusRequest) {
    request = &DescribeIPStrategyApisStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeIPStrategyApisStatus")
    
    
    return
}

func NewDescribeIPStrategyApisStatusResponse() (response *DescribeIPStrategyApisStatusResponse) {
    response = &DescribeIPStrategyApisStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIPStrategyApisStatus
// This API is used to query the list of APIs to which an IP policy can be bound, i.e., the difference set between all APIs under the service and the APIs already bound to the policy.
//
// error code that may be returned:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDIPSTRATEGY = "ResourceNotFound.InvalidIPStrategy"
func (c *Client) DescribeIPStrategyApisStatus(request *DescribeIPStrategyApisStatusRequest) (response *DescribeIPStrategyApisStatusResponse, err error) {
    return c.DescribeIPStrategyApisStatusWithContext(context.Background(), request)
}

// DescribeIPStrategyApisStatus
// This API is used to query the list of APIs to which an IP policy can be bound, i.e., the difference set between all APIs under the service and the APIs already bound to the policy.
//
// error code that may be returned:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDIPSTRATEGY = "ResourceNotFound.InvalidIPStrategy"
func (c *Client) DescribeIPStrategyApisStatusWithContext(ctx context.Context, request *DescribeIPStrategyApisStatusRequest) (response *DescribeIPStrategyApisStatusResponse, err error) {
    if request == nil {
        request = NewDescribeIPStrategyApisStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIPStrategyApisStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIPStrategyApisStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIPStrategysStatusRequest() (request *DescribeIPStrategysStatusRequest) {
    request = &DescribeIPStrategysStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeIPStrategysStatus")
    
    
    return
}

func NewDescribeIPStrategysStatusResponse() (response *DescribeIPStrategysStatusResponse) {
    response = &DescribeIPStrategysStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIPStrategysStatus
// This API is used to query the list of service IP policies.
//
// error code that may be returned:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeIPStrategysStatus(request *DescribeIPStrategysStatusRequest) (response *DescribeIPStrategysStatusResponse, err error) {
    return c.DescribeIPStrategysStatusWithContext(context.Background(), request)
}

// DescribeIPStrategysStatus
// This API is used to query the list of service IP policies.
//
// error code that may be returned:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeIPStrategysStatusWithContext(ctx context.Context, request *DescribeIPStrategysStatusRequest) (response *DescribeIPStrategysStatusResponse, err error) {
    if request == nil {
        request = NewDescribeIPStrategysStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIPStrategysStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIPStrategysStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogSearchRequest() (request *DescribeLogSearchRequest) {
    request = &DescribeLogSearchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeLogSearch")
    
    
    return
}

func NewDescribeLogSearchResponse() (response *DescribeLogSearchResponse) {
    response = &DescribeLogSearchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLogSearch
// This API is used to search for logs.
//
// error code that may be returned:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_CLSSEARCHTIME = "UnsupportedOperation.ClsSearchTime"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DescribeLogSearch(request *DescribeLogSearchRequest) (response *DescribeLogSearchResponse, err error) {
    return c.DescribeLogSearchWithContext(context.Background(), request)
}

// DescribeLogSearch
// This API is used to search for logs.
//
// error code that may be returned:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_CLSSEARCHTIME = "UnsupportedOperation.ClsSearchTime"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DescribeLogSearchWithContext(ctx context.Context, request *DescribeLogSearchRequest) (response *DescribeLogSearchResponse, err error) {
    if request == nil {
        request = NewDescribeLogSearchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogSearch require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogSearchResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePluginRequest() (request *DescribePluginRequest) {
    request = &DescribePluginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribePlugin")
    
    
    return
}

func NewDescribePluginResponse() (response *DescribePluginResponse) {
    response = &DescribePluginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePlugin
// This API is used to query the plugin details by plugin ID.
//
// error code that may be returned:
//  INTERNALERROR_CAMEXCEPTION = "InternalError.CamException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDPLUGIN = "ResourceNotFound.InvalidPlugin"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
func (c *Client) DescribePlugin(request *DescribePluginRequest) (response *DescribePluginResponse, err error) {
    return c.DescribePluginWithContext(context.Background(), request)
}

// DescribePlugin
// This API is used to query the plugin details by plugin ID.
//
// error code that may be returned:
//  INTERNALERROR_CAMEXCEPTION = "InternalError.CamException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDPLUGIN = "ResourceNotFound.InvalidPlugin"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
func (c *Client) DescribePluginWithContext(ctx context.Context, request *DescribePluginRequest) (response *DescribePluginResponse, err error) {
    if request == nil {
        request = NewDescribePluginRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePlugin require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePluginResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePluginApisRequest() (request *DescribePluginApisRequest) {
    request = &DescribePluginApisRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribePluginApis")
    
    
    return
}

func NewDescribePluginApisResponse() (response *DescribePluginApisResponse) {
    response = &DescribePluginApisResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePluginApis
// This API is used to query APIs bound with a specified plugin.
//
// error code that may be returned:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  RESOURCENOTFOUND_INVALIDPLUGIN = "ResourceNotFound.InvalidPlugin"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
func (c *Client) DescribePluginApis(request *DescribePluginApisRequest) (response *DescribePluginApisResponse, err error) {
    return c.DescribePluginApisWithContext(context.Background(), request)
}

// DescribePluginApis
// This API is used to query APIs bound with a specified plugin.
//
// error code that may be returned:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  RESOURCENOTFOUND_INVALIDPLUGIN = "ResourceNotFound.InvalidPlugin"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
func (c *Client) DescribePluginApisWithContext(ctx context.Context, request *DescribePluginApisRequest) (response *DescribePluginApisResponse, err error) {
    if request == nil {
        request = NewDescribePluginApisRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePluginApis require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePluginApisResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePluginsByApiRequest() (request *DescribePluginsByApiRequest) {
    request = &DescribePluginsByApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribePluginsByApi")
    
    
    return
}

func NewDescribePluginsByApiResponse() (response *DescribePluginsByApiResponse) {
    response = &DescribePluginsByApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePluginsByApi
// This API is used to query all plug-ins bound with the API.
//
// error code that may be returned:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DescribePluginsByApi(request *DescribePluginsByApiRequest) (response *DescribePluginsByApiResponse, err error) {
    return c.DescribePluginsByApiWithContext(context.Background(), request)
}

// DescribePluginsByApi
// This API is used to query all plug-ins bound with the API.
//
// error code that may be returned:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DescribePluginsByApiWithContext(ctx context.Context, request *DescribePluginsByApiRequest) (response *DescribePluginsByApiResponse, err error) {
    if request == nil {
        request = NewDescribePluginsByApiRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePluginsByApi require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePluginsByApiResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceRequest() (request *DescribeServiceRequest) {
    request = &DescribeServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeService")
    
    
    return
}

func NewDescribeServiceResponse() (response *DescribeServiceResponse) {
    response = &DescribeServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeService
// This API is used to query the details of a service, such as its description, domain name, protocol, creation time, and releases.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CAMEXCEPTION = "FailedOperation.CamException"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMEXCEPTION = "InternalError.CamException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_RETURNABLEEXCEPTION = "InternalError.ReturnableException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
//  UNSUPPORTEDOPERATION_REQUESTPOSTERROR = "UnsupportedOperation.RequestPostError"
//  UNSUPPORTEDOPERATION_UNPACKERROR = "UnsupportedOperation.UnpackError"
func (c *Client) DescribeService(request *DescribeServiceRequest) (response *DescribeServiceResponse, err error) {
    return c.DescribeServiceWithContext(context.Background(), request)
}

// DescribeService
// This API is used to query the details of a service, such as its description, domain name, protocol, creation time, and releases.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CAMEXCEPTION = "FailedOperation.CamException"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMEXCEPTION = "InternalError.CamException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_RETURNABLEEXCEPTION = "InternalError.ReturnableException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
//  UNSUPPORTEDOPERATION_REQUESTPOSTERROR = "UnsupportedOperation.RequestPostError"
//  UNSUPPORTEDOPERATION_UNPACKERROR = "UnsupportedOperation.UnpackError"
func (c *Client) DescribeServiceWithContext(ctx context.Context, request *DescribeServiceRequest) (response *DescribeServiceResponse, err error) {
    if request == nil {
        request = NewDescribeServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeService require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceEnvironmentListRequest() (request *DescribeServiceEnvironmentListRequest) {
    request = &DescribeServiceEnvironmentListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServiceEnvironmentList")
    
    
    return
}

func NewDescribeServiceEnvironmentListResponse() (response *DescribeServiceEnvironmentListResponse) {
    response = &DescribeServiceEnvironmentListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeServiceEnvironmentList
// This API is used to query the list of environments under a service. All environments and their statuses under the queried service will be returned.
//
// error code that may be returned:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeServiceEnvironmentList(request *DescribeServiceEnvironmentListRequest) (response *DescribeServiceEnvironmentListResponse, err error) {
    return c.DescribeServiceEnvironmentListWithContext(context.Background(), request)
}

// DescribeServiceEnvironmentList
// This API is used to query the list of environments under a service. All environments and their statuses under the queried service will be returned.
//
// error code that may be returned:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeServiceEnvironmentListWithContext(ctx context.Context, request *DescribeServiceEnvironmentListRequest) (response *DescribeServiceEnvironmentListResponse, err error) {
    if request == nil {
        request = NewDescribeServiceEnvironmentListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServiceEnvironmentList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServiceEnvironmentListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceEnvironmentReleaseHistoryRequest() (request *DescribeServiceEnvironmentReleaseHistoryRequest) {
    request = &DescribeServiceEnvironmentReleaseHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServiceEnvironmentReleaseHistory")
    
    
    return
}

func NewDescribeServiceEnvironmentReleaseHistoryResponse() (response *DescribeServiceEnvironmentReleaseHistoryResponse) {
    response = &DescribeServiceEnvironmentReleaseHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeServiceEnvironmentReleaseHistory
// This API is used to query the release history in a service environment.
//
// A service can only be used when it is published to an environment after creation. This API is used to query the release history in an environment under a service.
//
// error code that may be returned:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeServiceEnvironmentReleaseHistory(request *DescribeServiceEnvironmentReleaseHistoryRequest) (response *DescribeServiceEnvironmentReleaseHistoryResponse, err error) {
    return c.DescribeServiceEnvironmentReleaseHistoryWithContext(context.Background(), request)
}

// DescribeServiceEnvironmentReleaseHistory
// This API is used to query the release history in a service environment.
//
// A service can only be used when it is published to an environment after creation. This API is used to query the release history in an environment under a service.
//
// error code that may be returned:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeServiceEnvironmentReleaseHistoryWithContext(ctx context.Context, request *DescribeServiceEnvironmentReleaseHistoryRequest) (response *DescribeServiceEnvironmentReleaseHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeServiceEnvironmentReleaseHistoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServiceEnvironmentReleaseHistory require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServiceEnvironmentReleaseHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceEnvironmentStrategyRequest() (request *DescribeServiceEnvironmentStrategyRequest) {
    request = &DescribeServiceEnvironmentStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServiceEnvironmentStrategy")
    
    
    return
}

func NewDescribeServiceEnvironmentStrategyResponse() (response *DescribeServiceEnvironmentStrategyResponse) {
    response = &DescribeServiceEnvironmentStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeServiceEnvironmentStrategy
// This API is used to display a service throttling policy.
//
// error code that may be returned:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeServiceEnvironmentStrategy(request *DescribeServiceEnvironmentStrategyRequest) (response *DescribeServiceEnvironmentStrategyResponse, err error) {
    return c.DescribeServiceEnvironmentStrategyWithContext(context.Background(), request)
}

// DescribeServiceEnvironmentStrategy
// This API is used to display a service throttling policy.
//
// error code that may be returned:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeServiceEnvironmentStrategyWithContext(ctx context.Context, request *DescribeServiceEnvironmentStrategyRequest) (response *DescribeServiceEnvironmentStrategyResponse, err error) {
    if request == nil {
        request = NewDescribeServiceEnvironmentStrategyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServiceEnvironmentStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServiceEnvironmentStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceForApiAppRequest() (request *DescribeServiceForApiAppRequest) {
    request = &DescribeServiceForApiAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServiceForApiApp")
    
    
    return
}

func NewDescribeServiceForApiAppResponse() (response *DescribeServiceForApiAppResponse) {
    response = &DescribeServiceForApiAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeServiceForApiApp
// This API is used to query the details of a service, such as its description, domain name, and protocol.
//
// error code that may be returned:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeServiceForApiApp(request *DescribeServiceForApiAppRequest) (response *DescribeServiceForApiAppResponse, err error) {
    return c.DescribeServiceForApiAppWithContext(context.Background(), request)
}

// DescribeServiceForApiApp
// This API is used to query the details of a service, such as its description, domain name, and protocol.
//
// error code that may be returned:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeServiceForApiAppWithContext(ctx context.Context, request *DescribeServiceForApiAppRequest) (response *DescribeServiceForApiAppResponse, err error) {
    if request == nil {
        request = NewDescribeServiceForApiAppRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServiceForApiApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServiceForApiAppResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceReleaseVersionRequest() (request *DescribeServiceReleaseVersionRequest) {
    request = &DescribeServiceReleaseVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServiceReleaseVersion")
    
    
    return
}

func NewDescribeServiceReleaseVersionResponse() (response *DescribeServiceReleaseVersionResponse) {
    response = &DescribeServiceReleaseVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeServiceReleaseVersion
// This API is used to query the list of all published versions under a service.
//
// A service is generally published on several versions. This API can be used to query the published versions.
//
// error code that may be returned:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeServiceReleaseVersion(request *DescribeServiceReleaseVersionRequest) (response *DescribeServiceReleaseVersionResponse, err error) {
    return c.DescribeServiceReleaseVersionWithContext(context.Background(), request)
}

// DescribeServiceReleaseVersion
// This API is used to query the list of all published versions under a service.
//
// A service is generally published on several versions. This API can be used to query the published versions.
//
// error code that may be returned:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeServiceReleaseVersionWithContext(ctx context.Context, request *DescribeServiceReleaseVersionRequest) (response *DescribeServiceReleaseVersionResponse, err error) {
    if request == nil {
        request = NewDescribeServiceReleaseVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServiceReleaseVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServiceReleaseVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceSubDomainMappingsRequest() (request *DescribeServiceSubDomainMappingsRequest) {
    request = &DescribeServiceSubDomainMappingsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServiceSubDomainMappings")
    
    
    return
}

func NewDescribeServiceSubDomainMappingsResponse() (response *DescribeServiceSubDomainMappingsResponse) {
    response = &DescribeServiceSubDomainMappingsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeServiceSubDomainMappings
// This API is used to query the path mappings of a custom domain name.
//
// In API Gateway, you can bind a custom domain name to a service and map its paths. You can customize different path mappings to up to 3 environments under the service. This API is used to query the list of path mappings of a custom domain name bound to a service.
//
// error code that may be returned:
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICENOTEXIST = "FailedOperation.ServiceNotExist"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeServiceSubDomainMappings(request *DescribeServiceSubDomainMappingsRequest) (response *DescribeServiceSubDomainMappingsResponse, err error) {
    return c.DescribeServiceSubDomainMappingsWithContext(context.Background(), request)
}

// DescribeServiceSubDomainMappings
// This API is used to query the path mappings of a custom domain name.
//
// In API Gateway, you can bind a custom domain name to a service and map its paths. You can customize different path mappings to up to 3 environments under the service. This API is used to query the list of path mappings of a custom domain name bound to a service.
//
// error code that may be returned:
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICENOTEXIST = "FailedOperation.ServiceNotExist"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeServiceSubDomainMappingsWithContext(ctx context.Context, request *DescribeServiceSubDomainMappingsRequest) (response *DescribeServiceSubDomainMappingsResponse, err error) {
    if request == nil {
        request = NewDescribeServiceSubDomainMappingsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServiceSubDomainMappings require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServiceSubDomainMappingsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceSubDomainsRequest() (request *DescribeServiceSubDomainsRequest) {
    request = &DescribeServiceSubDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServiceSubDomains")
    
    
    return
}

func NewDescribeServiceSubDomainsResponse() (response *DescribeServiceSubDomainsResponse) {
    response = &DescribeServiceSubDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeServiceSubDomains
// This API is used to query the list of custom domain names.
//
// In API Gateway, you can bind custom domain names to a service for service call. This API is used to query the list of custom domain names bound to a service.
//
// error code that may be returned:
//  FAILEDOPERATION_DESCRIBESERVICESUBDOMAINSERROR = "FailedOperation.DescribeServiceSubDomainsError"
//  FAILEDOPERATION_SERVICENOTEXIST = "FailedOperation.ServiceNotExist"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeServiceSubDomains(request *DescribeServiceSubDomainsRequest) (response *DescribeServiceSubDomainsResponse, err error) {
    return c.DescribeServiceSubDomainsWithContext(context.Background(), request)
}

// DescribeServiceSubDomains
// This API is used to query the list of custom domain names.
//
// In API Gateway, you can bind custom domain names to a service for service call. This API is used to query the list of custom domain names bound to a service.
//
// error code that may be returned:
//  FAILEDOPERATION_DESCRIBESERVICESUBDOMAINSERROR = "FailedOperation.DescribeServiceSubDomainsError"
//  FAILEDOPERATION_SERVICENOTEXIST = "FailedOperation.ServiceNotExist"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeServiceSubDomainsWithContext(ctx context.Context, request *DescribeServiceSubDomainsRequest) (response *DescribeServiceSubDomainsResponse, err error) {
    if request == nil {
        request = NewDescribeServiceSubDomainsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServiceSubDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServiceSubDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceUsagePlanRequest() (request *DescribeServiceUsagePlanRequest) {
    request = &DescribeServiceUsagePlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServiceUsagePlan")
    
    
    return
}

func NewDescribeServiceUsagePlanResponse() (response *DescribeServiceUsagePlanResponse) {
    response = &DescribeServiceUsagePlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeServiceUsagePlan
// This API is used to query the details of usage plans in a service.
//
// To make authentication and throttling for a service take effect, you need to bind a usage plan to it. This API is used to query all usage plans bound to a service.
//
// error code that may be returned:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeServiceUsagePlan(request *DescribeServiceUsagePlanRequest) (response *DescribeServiceUsagePlanResponse, err error) {
    return c.DescribeServiceUsagePlanWithContext(context.Background(), request)
}

// DescribeServiceUsagePlan
// This API is used to query the details of usage plans in a service.
//
// To make authentication and throttling for a service take effect, you need to bind a usage plan to it. This API is used to query all usage plans bound to a service.
//
// error code that may be returned:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeServiceUsagePlanWithContext(ctx context.Context, request *DescribeServiceUsagePlanRequest) (response *DescribeServiceUsagePlanResponse, err error) {
    if request == nil {
        request = NewDescribeServiceUsagePlanRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServiceUsagePlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServiceUsagePlanResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServicesStatusRequest() (request *DescribeServicesStatusRequest) {
    request = &DescribeServicesStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServicesStatus")
    
    
    return
}

func NewDescribeServicesStatusResponse() (response *DescribeServicesStatusResponse) {
    response = &DescribeServicesStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeServicesStatus
// This API is used to query the list of one or more services and return relevant domain name, time, and other information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CAMEXCEPTION = "FailedOperation.CamException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMEXCEPTION = "InternalError.CamException"
//  INTERNALERROR_DBEXCEPTION = "InternalError.DbException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_INVALIDTAGVALUES = "InvalidParameterValue.InvalidTagValues"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_REQUESTPOSTERROR = "UnsupportedOperation.RequestPostError"
func (c *Client) DescribeServicesStatus(request *DescribeServicesStatusRequest) (response *DescribeServicesStatusResponse, err error) {
    return c.DescribeServicesStatusWithContext(context.Background(), request)
}

// DescribeServicesStatus
// This API is used to query the list of one or more services and return relevant domain name, time, and other information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CAMEXCEPTION = "FailedOperation.CamException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMEXCEPTION = "InternalError.CamException"
//  INTERNALERROR_DBEXCEPTION = "InternalError.DbException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_INVALIDTAGVALUES = "InvalidParameterValue.InvalidTagValues"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_REQUESTPOSTERROR = "UnsupportedOperation.RequestPostError"
func (c *Client) DescribeServicesStatusWithContext(ctx context.Context, request *DescribeServicesStatusRequest) (response *DescribeServicesStatusResponse, err error) {
    if request == nil {
        request = NewDescribeServicesStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServicesStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServicesStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUpstreamBindApisRequest() (request *DescribeUpstreamBindApisRequest) {
    request = &DescribeUpstreamBindApisRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeUpstreamBindApis")
    
    
    return
}

func NewDescribeUpstreamBindApisResponse() (response *DescribeUpstreamBindApisResponse) {
    response = &DescribeUpstreamBindApisResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUpstreamBindApis
// This API is used to query APIs bound with an upstream.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUpstreamBindApis(request *DescribeUpstreamBindApisRequest) (response *DescribeUpstreamBindApisResponse, err error) {
    return c.DescribeUpstreamBindApisWithContext(context.Background(), request)
}

// DescribeUpstreamBindApis
// This API is used to query APIs bound with an upstream.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUpstreamBindApisWithContext(ctx context.Context, request *DescribeUpstreamBindApisRequest) (response *DescribeUpstreamBindApisResponse, err error) {
    if request == nil {
        request = NewDescribeUpstreamBindApisRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUpstreamBindApis require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUpstreamBindApisResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUpstreamsRequest() (request *DescribeUpstreamsRequest) {
    request = &DescribeUpstreamsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeUpstreams")
    
    
    return
}

func NewDescribeUpstreamsResponse() (response *DescribeUpstreamsResponse) {
    response = &DescribeUpstreamsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUpstreams
// This API is used to query details of upstreams under the current account.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CAMEXCEPTION = "FailedOperation.CamException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMEXCEPTION = "InternalError.CamException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INVALIDACCESSKEYID = "ResourceNotFound.InvalidAccessKeyId"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_REQUESTPOSTERROR = "UnsupportedOperation.RequestPostError"
func (c *Client) DescribeUpstreams(request *DescribeUpstreamsRequest) (response *DescribeUpstreamsResponse, err error) {
    return c.DescribeUpstreamsWithContext(context.Background(), request)
}

// DescribeUpstreams
// This API is used to query details of upstreams under the current account.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CAMEXCEPTION = "FailedOperation.CamException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMEXCEPTION = "InternalError.CamException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INVALIDACCESSKEYID = "ResourceNotFound.InvalidAccessKeyId"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_REQUESTPOSTERROR = "UnsupportedOperation.RequestPostError"
func (c *Client) DescribeUpstreamsWithContext(ctx context.Context, request *DescribeUpstreamsRequest) (response *DescribeUpstreamsResponse, err error) {
    if request == nil {
        request = NewDescribeUpstreamsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUpstreams require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUpstreamsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsagePlanRequest() (request *DescribeUsagePlanRequest) {
    request = &DescribeUsagePlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeUsagePlan")
    
    
    return
}

func NewDescribeUsagePlanResponse() (response *DescribeUsagePlanResponse) {
    response = &DescribeUsagePlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUsagePlan
// This API is used to query the details of a usage plan, such as its name, QPS, creation time, and bound environment.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
func (c *Client) DescribeUsagePlan(request *DescribeUsagePlanRequest) (response *DescribeUsagePlanResponse, err error) {
    return c.DescribeUsagePlanWithContext(context.Background(), request)
}

// DescribeUsagePlan
// This API is used to query the details of a usage plan, such as its name, QPS, creation time, and bound environment.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
func (c *Client) DescribeUsagePlanWithContext(ctx context.Context, request *DescribeUsagePlanRequest) (response *DescribeUsagePlanResponse, err error) {
    if request == nil {
        request = NewDescribeUsagePlanRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUsagePlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUsagePlanResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsagePlanEnvironmentsRequest() (request *DescribeUsagePlanEnvironmentsRequest) {
    request = &DescribeUsagePlanEnvironmentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeUsagePlanEnvironments")
    
    
    return
}

func NewDescribeUsagePlanEnvironmentsResponse() (response *DescribeUsagePlanEnvironmentsResponse) {
    response = &DescribeUsagePlanEnvironmentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUsagePlanEnvironments
// This API is used to query the list of environments bound to a usage plan.
//
// After binding a usage plan to environments, you can use this API to query all service environments bound to the usage plan.
//
// error code that may be returned:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
func (c *Client) DescribeUsagePlanEnvironments(request *DescribeUsagePlanEnvironmentsRequest) (response *DescribeUsagePlanEnvironmentsResponse, err error) {
    return c.DescribeUsagePlanEnvironmentsWithContext(context.Background(), request)
}

// DescribeUsagePlanEnvironments
// This API is used to query the list of environments bound to a usage plan.
//
// After binding a usage plan to environments, you can use this API to query all service environments bound to the usage plan.
//
// error code that may be returned:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
func (c *Client) DescribeUsagePlanEnvironmentsWithContext(ctx context.Context, request *DescribeUsagePlanEnvironmentsRequest) (response *DescribeUsagePlanEnvironmentsResponse, err error) {
    if request == nil {
        request = NewDescribeUsagePlanEnvironmentsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUsagePlanEnvironments require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUsagePlanEnvironmentsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsagePlanSecretIdsRequest() (request *DescribeUsagePlanSecretIdsRequest) {
    request = &DescribeUsagePlanSecretIdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeUsagePlanSecretIds")
    
    
    return
}

func NewDescribeUsagePlanSecretIdsResponse() (response *DescribeUsagePlanSecretIdsResponse) {
    response = &DescribeUsagePlanSecretIdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUsagePlanSecretIds
// This API is used to query the list of keys bound to a usage plan.
//
// In API Gateway, a usage plan can be bound to multiple key pairs. You can use this API to query the list of keys bound to it.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
func (c *Client) DescribeUsagePlanSecretIds(request *DescribeUsagePlanSecretIdsRequest) (response *DescribeUsagePlanSecretIdsResponse, err error) {
    return c.DescribeUsagePlanSecretIdsWithContext(context.Background(), request)
}

// DescribeUsagePlanSecretIds
// This API is used to query the list of keys bound to a usage plan.
//
// In API Gateway, a usage plan can be bound to multiple key pairs. You can use this API to query the list of keys bound to it.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
func (c *Client) DescribeUsagePlanSecretIdsWithContext(ctx context.Context, request *DescribeUsagePlanSecretIdsRequest) (response *DescribeUsagePlanSecretIdsResponse, err error) {
    if request == nil {
        request = NewDescribeUsagePlanSecretIdsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUsagePlanSecretIds require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUsagePlanSecretIdsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsagePlansStatusRequest() (request *DescribeUsagePlansStatusRequest) {
    request = &DescribeUsagePlansStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeUsagePlansStatus")
    
    
    return
}

func NewDescribeUsagePlansStatusResponse() (response *DescribeUsagePlansStatusResponse) {
    response = &DescribeUsagePlansStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUsagePlansStatus
// This API is used to query the list of usage plans.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
func (c *Client) DescribeUsagePlansStatus(request *DescribeUsagePlansStatusRequest) (response *DescribeUsagePlansStatusResponse, err error) {
    return c.DescribeUsagePlansStatusWithContext(context.Background(), request)
}

// DescribeUsagePlansStatus
// This API is used to query the list of usage plans.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
func (c *Client) DescribeUsagePlansStatusWithContext(ctx context.Context, request *DescribeUsagePlansStatusRequest) (response *DescribeUsagePlansStatusResponse, err error) {
    if request == nil {
        request = NewDescribeUsagePlansStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUsagePlansStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUsagePlansStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDetachPluginRequest() (request *DetachPluginRequest) {
    request = &DetachPluginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DetachPlugin")
    
    
    return
}

func NewDetachPluginResponse() (response *DetachPluginResponse) {
    response = &DetachPluginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DetachPlugin
// This API is used to unbind an API from the plugin.
//
// error code that may be returned:
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDAPIIDS = "InvalidParameterValue.InvalidApiIds"
//  INVALIDPARAMETERVALUE_INVALIDENVSTATUS = "InvalidParameterValue.InvalidEnvStatus"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDPLUGIN = "ResourceNotFound.InvalidPlugin"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
//  UNSUPPORTEDOPERATION_UNPACKERROR = "UnsupportedOperation.UnpackError"
func (c *Client) DetachPlugin(request *DetachPluginRequest) (response *DetachPluginResponse, err error) {
    return c.DetachPluginWithContext(context.Background(), request)
}

// DetachPlugin
// This API is used to unbind an API from the plugin.
//
// error code that may be returned:
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDAPIIDS = "InvalidParameterValue.InvalidApiIds"
//  INVALIDPARAMETERVALUE_INVALIDENVSTATUS = "InvalidParameterValue.InvalidEnvStatus"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDPLUGIN = "ResourceNotFound.InvalidPlugin"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
//  UNSUPPORTEDOPERATION_UNPACKERROR = "UnsupportedOperation.UnpackError"
func (c *Client) DetachPluginWithContext(ctx context.Context, request *DetachPluginRequest) (response *DetachPluginResponse, err error) {
    if request == nil {
        request = NewDetachPluginRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetachPlugin require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetachPluginResponse()
    err = c.Send(request, response)
    return
}

func NewDisableApiKeyRequest() (request *DisableApiKeyRequest) {
    request = &DisableApiKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DisableApiKey")
    
    
    return
}

func NewDisableApiKeyResponse() (response *DisableApiKeyResponse) {
    response = &DisableApiKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisableApiKey
// This API is used to disable an API key.
//
// error code that may be returned:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  RESOURCENOTFOUND_INVALIDACCESSKEYID = "ResourceNotFound.InvalidAccessKeyId"
//  UNSUPPORTEDOPERATION_UINNOTINWHITELIST = "UnsupportedOperation.UinNotInWhiteList"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDUPDATEAPIKEY = "UnsupportedOperation.UnsupportedUpdateApiKey"
func (c *Client) DisableApiKey(request *DisableApiKeyRequest) (response *DisableApiKeyResponse, err error) {
    return c.DisableApiKeyWithContext(context.Background(), request)
}

// DisableApiKey
// This API is used to disable an API key.
//
// error code that may be returned:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  RESOURCENOTFOUND_INVALIDACCESSKEYID = "ResourceNotFound.InvalidAccessKeyId"
//  UNSUPPORTEDOPERATION_UINNOTINWHITELIST = "UnsupportedOperation.UinNotInWhiteList"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDUPDATEAPIKEY = "UnsupportedOperation.UnsupportedUpdateApiKey"
func (c *Client) DisableApiKeyWithContext(ctx context.Context, request *DisableApiKeyRequest) (response *DisableApiKeyResponse, err error) {
    if request == nil {
        request = NewDisableApiKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableApiKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableApiKeyResponse()
    err = c.Send(request, response)
    return
}

func NewEnableApiKeyRequest() (request *EnableApiKeyRequest) {
    request = &EnableApiKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "EnableApiKey")
    
    
    return
}

func NewEnableApiKeyResponse() (response *EnableApiKeyResponse) {
    response = &EnableApiKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnableApiKey
// This API is used to enable a disabled API key.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INVALIDACCESSKEYID = "ResourceNotFound.InvalidAccessKeyId"
//  UNSUPPORTEDOPERATION_UINNOTINWHITELIST = "UnsupportedOperation.UinNotInWhiteList"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDUPDATEAPIKEY = "UnsupportedOperation.UnsupportedUpdateApiKey"
func (c *Client) EnableApiKey(request *EnableApiKeyRequest) (response *EnableApiKeyResponse, err error) {
    return c.EnableApiKeyWithContext(context.Background(), request)
}

// EnableApiKey
// This API is used to enable a disabled API key.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INVALIDACCESSKEYID = "ResourceNotFound.InvalidAccessKeyId"
//  UNSUPPORTEDOPERATION_UINNOTINWHITELIST = "UnsupportedOperation.UinNotInWhiteList"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDUPDATEAPIKEY = "UnsupportedOperation.UnsupportedUpdateApiKey"
func (c *Client) EnableApiKeyWithContext(ctx context.Context, request *EnableApiKeyRequest) (response *EnableApiKeyResponse, err error) {
    if request == nil {
        request = NewEnableApiKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableApiKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableApiKeyResponse()
    err = c.Send(request, response)
    return
}

func NewImportOpenApiRequest() (request *ImportOpenApiRequest) {
    request = &ImportOpenApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "ImportOpenApi")
    
    
    return
}

func NewImportOpenApiResponse() (response *ImportOpenApiResponse) {
    response = &ImportOpenApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ImportOpenApi
// This API is used to import an OpenAPI to API gateway. 
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDSERVICECONFIG = "InvalidParameterValue.InvalidServiceConfig"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDPARAMETER = "InvalidParameterValue.UnsupportedParameter"
//  LIMITEXCEEDED_APICOUNTLIMITEXCEEDED = "LimitExceeded.ApiCountLimitExceeded"
//  MISSINGPARAMETER_BACKENDSPECIFICPARAM = "MissingParameter.BackendSpecificParam"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_UNPACKERROR = "UnsupportedOperation.UnpackError"
func (c *Client) ImportOpenApi(request *ImportOpenApiRequest) (response *ImportOpenApiResponse, err error) {
    return c.ImportOpenApiWithContext(context.Background(), request)
}

// ImportOpenApi
// This API is used to import an OpenAPI to API gateway. 
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDSERVICECONFIG = "InvalidParameterValue.InvalidServiceConfig"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDPARAMETER = "InvalidParameterValue.UnsupportedParameter"
//  LIMITEXCEEDED_APICOUNTLIMITEXCEEDED = "LimitExceeded.ApiCountLimitExceeded"
//  MISSINGPARAMETER_BACKENDSPECIFICPARAM = "MissingParameter.BackendSpecificParam"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_UNPACKERROR = "UnsupportedOperation.UnpackError"
func (c *Client) ImportOpenApiWithContext(ctx context.Context, request *ImportOpenApiRequest) (response *ImportOpenApiResponse, err error) {
    if request == nil {
        request = NewImportOpenApiRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ImportOpenApi require credential")
    }

    request.SetContext(ctx)
    
    response = NewImportOpenApiResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAPIDocRequest() (request *ModifyAPIDocRequest) {
    request = &ModifyAPIDocRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifyAPIDoc")
    
    
    return
}

func NewModifyAPIDocResponse() (response *ModifyAPIDocResponse) {
    response = &ModifyAPIDocResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAPIDoc
// This API is used to modify an API document.
//
// error code that may be returned:
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  FAILEDOPERATION_CODINGERROR = "FailedOperation.CodingError"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDAPIDOC = "ResourceNotFound.InvalidApiDoc"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) ModifyAPIDoc(request *ModifyAPIDocRequest) (response *ModifyAPIDocResponse, err error) {
    return c.ModifyAPIDocWithContext(context.Background(), request)
}

// ModifyAPIDoc
// This API is used to modify an API document.
//
// error code that may be returned:
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  FAILEDOPERATION_CODINGERROR = "FailedOperation.CodingError"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDAPIDOC = "ResourceNotFound.InvalidApiDoc"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) ModifyAPIDocWithContext(ctx context.Context, request *ModifyAPIDocRequest) (response *ModifyAPIDocResponse, err error) {
    if request == nil {
        request = NewModifyAPIDocRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAPIDoc require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAPIDocResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApiRequest() (request *ModifyApiRequest) {
    request = &ModifyApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifyApi")
    
    
    return
}

func NewModifyApiResponse() (response *ModifyApiResponse) {
    response = &ModifyApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApi
// This API is used to modify an API. You can call this API to edit/modify a configured API. The modified API takes effect only after its service is published to the corresponding environment again.
//
// error code that may be returned:
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  FAILEDOPERATION_APIINOPERATION = "FailedOperation.ApiInOperation"
//  FAILEDOPERATION_BACKENDDOMAINERROR = "FailedOperation.BackendDomainError"
//  FAILEDOPERATION_DOMAINNEEDBEIAN = "FailedOperation.DomainNeedBeian"
//  FAILEDOPERATION_EIAMERROR = "FailedOperation.EIAMError"
//  FAILEDOPERATION_EBERROR = "FailedOperation.EbError"
//  FAILEDOPERATION_GETROLEERROR = "FailedOperation.GetRoleError"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_CLBEXCEPTION = "InternalError.ClbException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_SCFEXCEPTION = "InternalError.ScfException"
//  INTERNALERROR_TSFEXCEPTION = "InternalError.TsfException"
//  INTERNALERROR_VPCEXCEPTION = "InternalError.VpcException"
//  INVALIDPARAMETER_BASICSERVICENOTALLOWATTACHPLUGIN = "InvalidParameter.BasicServiceNotAllowAttachPlugin"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALPROXYIP = "InvalidParameterValue.IllegalProxyIp"
//  INVALIDPARAMETERVALUE_INVALIDAPIBUSINESSTYPE = "InvalidParameterValue.InvalidApiBusinessType"
//  INVALIDPARAMETERVALUE_INVALIDBACKENDPATH = "InvalidParameterValue.InvalidBackendPath"
//  INVALIDPARAMETERVALUE_INVALIDCLB = "InvalidParameterValue.InvalidClb"
//  INVALIDPARAMETERVALUE_INVALIDCONSTANTPARAMETERS = "InvalidParameterValue.InvalidConstantParameters"
//  INVALIDPARAMETERVALUE_INVALIDIPADDRESS = "InvalidParameterValue.InvalidIPAddress"
//  INVALIDPARAMETERVALUE_INVALIDMETHOD = "InvalidParameterValue.InvalidMethod"
//  INVALIDPARAMETERVALUE_INVALIDPORT = "InvalidParameterValue.InvalidPort"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICKEY = "InvalidParameterValue.InvalidPublicKey"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_INVALIDREQUESTPARAMETERS = "InvalidParameterValue.InvalidRequestParameters"
//  INVALIDPARAMETERVALUE_INVALIDSCFCONFIG = "InvalidParameterValue.InvalidScfConfig"
//  INVALIDPARAMETERVALUE_INVALIDSERVICECONFIG = "InvalidParameterValue.InvalidServiceConfig"
//  INVALIDPARAMETERVALUE_INVALIDSERVICEPARAM = "InvalidParameterValue.InvalidServiceParam"
//  INVALIDPARAMETERVALUE_INVALIDSERVICETYPE = "InvalidParameterValue.InvalidServiceType"
//  INVALIDPARAMETERVALUE_INVALIDUPSTREAM = "InvalidParameterValue.InvalidUpstream"
//  INVALIDPARAMETERVALUE_INVALIDURL = "InvalidParameterValue.InvalidUrl"
//  INVALIDPARAMETERVALUE_INVALIDVPCCONFIG = "InvalidParameterValue.InvalidVpcConfig"
//  INVALIDPARAMETERVALUE_INVALIDWSMETHOD = "InvalidParameterValue.InvalidWSMethod"
//  INVALIDPARAMETERVALUE_LENGTHEXCEEDED = "InvalidParameterValue.LengthExceeded"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NOTINOPTIONS = "InvalidParameterValue.NotInOptions"
//  INVALIDPARAMETERVALUE_PARAMETERNOTMATCH = "InvalidParameterValue.ParameterNotMatch"
//  INVALIDPARAMETERVALUE_RANGEEXCEEDED = "InvalidParameterValue.RangeExceeded"
//  LIMITEXCEEDED_APICOUNTLIMITEXCEEDED = "LimitExceeded.ApiCountLimitExceeded"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDOAUTHAPI = "ResourceNotFound.InvalidOauthApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_INVALIDENDPOINTTYPE = "UnsupportedOperation.InvalidEndpointType"
//  UNSUPPORTEDOPERATION_MODIFYEIAMAUTHAPI = "UnsupportedOperation.ModifyEIAMAuthApi"
//  UNSUPPORTEDOPERATION_MODIFYPROTOCOL = "UnsupportedOperation.ModifyProtocol"
//  UNSUPPORTEDOPERATION_RESOURCEASSOCIATED = "UnsupportedOperation.ResourceAssociated"
//  UNSUPPORTEDOPERATION_UNPACKERROR = "UnsupportedOperation.UnpackError"
func (c *Client) ModifyApi(request *ModifyApiRequest) (response *ModifyApiResponse, err error) {
    return c.ModifyApiWithContext(context.Background(), request)
}

// ModifyApi
// This API is used to modify an API. You can call this API to edit/modify a configured API. The modified API takes effect only after its service is published to the corresponding environment again.
//
// error code that may be returned:
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  FAILEDOPERATION_APIINOPERATION = "FailedOperation.ApiInOperation"
//  FAILEDOPERATION_BACKENDDOMAINERROR = "FailedOperation.BackendDomainError"
//  FAILEDOPERATION_DOMAINNEEDBEIAN = "FailedOperation.DomainNeedBeian"
//  FAILEDOPERATION_EIAMERROR = "FailedOperation.EIAMError"
//  FAILEDOPERATION_EBERROR = "FailedOperation.EbError"
//  FAILEDOPERATION_GETROLEERROR = "FailedOperation.GetRoleError"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_CLBEXCEPTION = "InternalError.ClbException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_SCFEXCEPTION = "InternalError.ScfException"
//  INTERNALERROR_TSFEXCEPTION = "InternalError.TsfException"
//  INTERNALERROR_VPCEXCEPTION = "InternalError.VpcException"
//  INVALIDPARAMETER_BASICSERVICENOTALLOWATTACHPLUGIN = "InvalidParameter.BasicServiceNotAllowAttachPlugin"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALPROXYIP = "InvalidParameterValue.IllegalProxyIp"
//  INVALIDPARAMETERVALUE_INVALIDAPIBUSINESSTYPE = "InvalidParameterValue.InvalidApiBusinessType"
//  INVALIDPARAMETERVALUE_INVALIDBACKENDPATH = "InvalidParameterValue.InvalidBackendPath"
//  INVALIDPARAMETERVALUE_INVALIDCLB = "InvalidParameterValue.InvalidClb"
//  INVALIDPARAMETERVALUE_INVALIDCONSTANTPARAMETERS = "InvalidParameterValue.InvalidConstantParameters"
//  INVALIDPARAMETERVALUE_INVALIDIPADDRESS = "InvalidParameterValue.InvalidIPAddress"
//  INVALIDPARAMETERVALUE_INVALIDMETHOD = "InvalidParameterValue.InvalidMethod"
//  INVALIDPARAMETERVALUE_INVALIDPORT = "InvalidParameterValue.InvalidPort"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICKEY = "InvalidParameterValue.InvalidPublicKey"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_INVALIDREQUESTPARAMETERS = "InvalidParameterValue.InvalidRequestParameters"
//  INVALIDPARAMETERVALUE_INVALIDSCFCONFIG = "InvalidParameterValue.InvalidScfConfig"
//  INVALIDPARAMETERVALUE_INVALIDSERVICECONFIG = "InvalidParameterValue.InvalidServiceConfig"
//  INVALIDPARAMETERVALUE_INVALIDSERVICEPARAM = "InvalidParameterValue.InvalidServiceParam"
//  INVALIDPARAMETERVALUE_INVALIDSERVICETYPE = "InvalidParameterValue.InvalidServiceType"
//  INVALIDPARAMETERVALUE_INVALIDUPSTREAM = "InvalidParameterValue.InvalidUpstream"
//  INVALIDPARAMETERVALUE_INVALIDURL = "InvalidParameterValue.InvalidUrl"
//  INVALIDPARAMETERVALUE_INVALIDVPCCONFIG = "InvalidParameterValue.InvalidVpcConfig"
//  INVALIDPARAMETERVALUE_INVALIDWSMETHOD = "InvalidParameterValue.InvalidWSMethod"
//  INVALIDPARAMETERVALUE_LENGTHEXCEEDED = "InvalidParameterValue.LengthExceeded"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NOTINOPTIONS = "InvalidParameterValue.NotInOptions"
//  INVALIDPARAMETERVALUE_PARAMETERNOTMATCH = "InvalidParameterValue.ParameterNotMatch"
//  INVALIDPARAMETERVALUE_RANGEEXCEEDED = "InvalidParameterValue.RangeExceeded"
//  LIMITEXCEEDED_APICOUNTLIMITEXCEEDED = "LimitExceeded.ApiCountLimitExceeded"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDOAUTHAPI = "ResourceNotFound.InvalidOauthApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_INVALIDENDPOINTTYPE = "UnsupportedOperation.InvalidEndpointType"
//  UNSUPPORTEDOPERATION_MODIFYEIAMAUTHAPI = "UnsupportedOperation.ModifyEIAMAuthApi"
//  UNSUPPORTEDOPERATION_MODIFYPROTOCOL = "UnsupportedOperation.ModifyProtocol"
//  UNSUPPORTEDOPERATION_RESOURCEASSOCIATED = "UnsupportedOperation.ResourceAssociated"
//  UNSUPPORTEDOPERATION_UNPACKERROR = "UnsupportedOperation.UnpackError"
func (c *Client) ModifyApiWithContext(ctx context.Context, request *ModifyApiRequest) (response *ModifyApiResponse, err error) {
    if request == nil {
        request = NewModifyApiRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApi require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApiResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApiAppRequest() (request *ModifyApiAppRequest) {
    request = &ModifyApiAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifyApiApp")
    
    
    return
}

func NewModifyApiAppResponse() (response *ModifyApiAppResponse) {
    response = &ModifyApiAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApiApp
// This API is used to modify a created API.
//
// error code that may be returned:
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  FAILEDOPERATION_APIINOPERATION = "FailedOperation.ApiInOperation"
//  FAILEDOPERATION_GETROLEERROR = "FailedOperation.GetRoleError"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_SCFEXCEPTION = "InternalError.ScfException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDAPIAPP = "ResourceNotFound.InvalidApiApp"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_INVALIDSERVICETRADE = "UnsupportedOperation.InvalidServiceTrade"
//  UNSUPPORTEDOPERATION_RESOURCEASSOCIATED = "UnsupportedOperation.ResourceAssociated"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDDELETEAPI = "UnsupportedOperation.UnsupportedDeleteApi"
func (c *Client) ModifyApiApp(request *ModifyApiAppRequest) (response *ModifyApiAppResponse, err error) {
    return c.ModifyApiAppWithContext(context.Background(), request)
}

// ModifyApiApp
// This API is used to modify a created API.
//
// error code that may be returned:
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  FAILEDOPERATION_APIINOPERATION = "FailedOperation.ApiInOperation"
//  FAILEDOPERATION_GETROLEERROR = "FailedOperation.GetRoleError"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_SCFEXCEPTION = "InternalError.ScfException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDAPIAPP = "ResourceNotFound.InvalidApiApp"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_INVALIDSERVICETRADE = "UnsupportedOperation.InvalidServiceTrade"
//  UNSUPPORTEDOPERATION_RESOURCEASSOCIATED = "UnsupportedOperation.ResourceAssociated"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDDELETEAPI = "UnsupportedOperation.UnsupportedDeleteApi"
func (c *Client) ModifyApiAppWithContext(ctx context.Context, request *ModifyApiAppRequest) (response *ModifyApiAppResponse, err error) {
    if request == nil {
        request = NewModifyApiAppRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApiApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApiAppResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApiEnvironmentStrategyRequest() (request *ModifyApiEnvironmentStrategyRequest) {
    request = &ModifyApiEnvironmentStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifyApiEnvironmentStrategy")
    
    
    return
}

func NewModifyApiEnvironmentStrategyResponse() (response *ModifyApiEnvironmentStrategyResponse) {
    response = &ModifyApiEnvironmentStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApiEnvironmentStrategy
// This API is used to modify an API throttling policy.
//
// error code that may be returned:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREQUESTPARAMETERS = "InvalidParameterValue.InvalidRequestParameters"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_RANGEEXCEEDED = "InvalidParameterValue.RangeExceeded"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
func (c *Client) ModifyApiEnvironmentStrategy(request *ModifyApiEnvironmentStrategyRequest) (response *ModifyApiEnvironmentStrategyResponse, err error) {
    return c.ModifyApiEnvironmentStrategyWithContext(context.Background(), request)
}

// ModifyApiEnvironmentStrategy
// This API is used to modify an API throttling policy.
//
// error code that may be returned:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREQUESTPARAMETERS = "InvalidParameterValue.InvalidRequestParameters"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_RANGEEXCEEDED = "InvalidParameterValue.RangeExceeded"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
func (c *Client) ModifyApiEnvironmentStrategyWithContext(ctx context.Context, request *ModifyApiEnvironmentStrategyRequest) (response *ModifyApiEnvironmentStrategyResponse, err error) {
    if request == nil {
        request = NewModifyApiEnvironmentStrategyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApiEnvironmentStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApiEnvironmentStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApiIncrementRequest() (request *ModifyApiIncrementRequest) {
    request = &ModifyApiIncrementRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifyApiIncrement")
    
    
    return
}

func NewModifyApiIncrementResponse() (response *ModifyApiIncrementResponse) {
    response = &ModifyApiIncrementResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApiIncrement
// This API is used to incrementally update an API and mainly called by programs (different from `ModifyApi`, which requires that full API parameters be passed in and is suitable for use in the console).
//
// error code that may be returned:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDAPIBUSINESSTYPE = "InvalidParameterValue.InvalidApiBusinessType"
//  INVALIDPARAMETERVALUE_NOTHINGMODIFYFOROAUTH = "InvalidParameterValue.NothingModifyForOauth"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) ModifyApiIncrement(request *ModifyApiIncrementRequest) (response *ModifyApiIncrementResponse, err error) {
    return c.ModifyApiIncrementWithContext(context.Background(), request)
}

// ModifyApiIncrement
// This API is used to incrementally update an API and mainly called by programs (different from `ModifyApi`, which requires that full API parameters be passed in and is suitable for use in the console).
//
// error code that may be returned:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDAPIBUSINESSTYPE = "InvalidParameterValue.InvalidApiBusinessType"
//  INVALIDPARAMETERVALUE_NOTHINGMODIFYFOROAUTH = "InvalidParameterValue.NothingModifyForOauth"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) ModifyApiIncrementWithContext(ctx context.Context, request *ModifyApiIncrementRequest) (response *ModifyApiIncrementResponse, err error) {
    if request == nil {
        request = NewModifyApiIncrementRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApiIncrement require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApiIncrementResponse()
    err = c.Send(request, response)
    return
}

func NewModifyIPStrategyRequest() (request *ModifyIPStrategyRequest) {
    request = &ModifyIPStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifyIPStrategy")
    
    
    return
}

func NewModifyIPStrategyResponse() (response *ModifyIPStrategyResponse) {
    response = &ModifyIPStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyIPStrategy
// This API is used to modify a service IP policy.
//
// error code that may be returned:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDIPSTRATEGY = "ResourceNotFound.InvalidIPStrategy"
func (c *Client) ModifyIPStrategy(request *ModifyIPStrategyRequest) (response *ModifyIPStrategyResponse, err error) {
    return c.ModifyIPStrategyWithContext(context.Background(), request)
}

// ModifyIPStrategy
// This API is used to modify a service IP policy.
//
// error code that may be returned:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDIPSTRATEGY = "ResourceNotFound.InvalidIPStrategy"
func (c *Client) ModifyIPStrategyWithContext(ctx context.Context, request *ModifyIPStrategyRequest) (response *ModifyIPStrategyResponse, err error) {
    if request == nil {
        request = NewModifyIPStrategyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyIPStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyIPStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPluginRequest() (request *ModifyPluginRequest) {
    request = &ModifyPluginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifyPlugin")
    
    
    return
}

func NewModifyPluginResponse() (response *ModifyPluginResponse) {
    response = &ModifyPluginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPlugin
// This API is used to modify a plugin.
//
// error code that may be returned:
//  FAILEDOPERATION_FORMATERROR = "FailedOperation.FormatError"
//  FAILEDOPERATION_GETROLEERROR = "FailedOperation.GetRoleError"
//  FAILEDOPERATION_INVALIDINPUTJSON = "FailedOperation.InvalidInputJSON"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_RETURNABLEEXCEPTION = "InternalError.ReturnableException"
//  INTERNALERROR_TSFEXCEPTION = "InternalError.TsfException"
//  INTERNALERROR_VPCEXCEPTION = "InternalError.VpcException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_DUPLICATEPLUGINCONFIG = "InvalidParameterValue.DuplicatePluginConfig"
//  INVALIDPARAMETERVALUE_INVALIDAPIREQUESTCONFIG = "InvalidParameterValue.InvalidApiRequestConfig"
//  INVALIDPARAMETERVALUE_INVALIDBACKENDPATH = "InvalidParameterValue.InvalidBackendPath"
//  INVALIDPARAMETERVALUE_INVALIDCLB = "InvalidParameterValue.InvalidClb"
//  INVALIDPARAMETERVALUE_INVALIDCONDITION = "InvalidParameterValue.InvalidCondition"
//  INVALIDPARAMETERVALUE_INVALIDIPADDRESS = "InvalidParameterValue.InvalidIPAddress"
//  INVALIDPARAMETERVALUE_INVALIDPLUGINCONFIG = "InvalidParameterValue.InvalidPluginConfig"
//  INVALIDPARAMETERVALUE_INVALIDPORT = "InvalidParameterValue.InvalidPort"
//  INVALIDPARAMETERVALUE_INVALIDREQUESTPARAMETERS = "InvalidParameterValue.InvalidRequestParameters"
//  INVALIDPARAMETERVALUE_INVALIDSCFCONFIG = "InvalidParameterValue.InvalidScfConfig"
//  INVALIDPARAMETERVALUE_INVALIDSERVICECONFIG = "InvalidParameterValue.InvalidServiceConfig"
//  INVALIDPARAMETERVALUE_INVALIDSERVICEPARAM = "InvalidParameterValue.InvalidServiceParam"
//  INVALIDPARAMETERVALUE_INVALIDSERVICETYPE = "InvalidParameterValue.InvalidServiceType"
//  INVALIDPARAMETERVALUE_INVALIDTSFCONFIG = "InvalidParameterValue.InvalidTsfConfig"
//  INVALIDPARAMETERVALUE_INVALIDUPSTREAM = "InvalidParameterValue.InvalidUpstream"
//  INVALIDPARAMETERVALUE_INVALIDURL = "InvalidParameterValue.InvalidUrl"
//  INVALIDPARAMETERVALUE_INVALIDVPCCONFIG = "InvalidParameterValue.InvalidVpcConfig"
//  INVALIDPARAMETERVALUE_LENGTHEXCEEDED = "InvalidParameterValue.LengthExceeded"
//  INVALIDPARAMETERVALUE_NOTINOPTIONS = "InvalidParameterValue.NotInOptions"
//  INVALIDPARAMETERVALUE_RANGEEXCEEDED = "InvalidParameterValue.RangeExceeded"
//  MISSINGPARAMETER_BACKENDSPECIFICPARAM = "MissingParameter.BackendSpecificParam"
//  MISSINGPARAMETER_PLUGINCONFIG = "MissingParameter.PluginConfig"
//  RESOURCENOTFOUND_INVALIDPLUGIN = "ResourceNotFound.InvalidPlugin"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_UNPACKERROR = "UnsupportedOperation.UnpackError"
func (c *Client) ModifyPlugin(request *ModifyPluginRequest) (response *ModifyPluginResponse, err error) {
    return c.ModifyPluginWithContext(context.Background(), request)
}

// ModifyPlugin
// This API is used to modify a plugin.
//
// error code that may be returned:
//  FAILEDOPERATION_FORMATERROR = "FailedOperation.FormatError"
//  FAILEDOPERATION_GETROLEERROR = "FailedOperation.GetRoleError"
//  FAILEDOPERATION_INVALIDINPUTJSON = "FailedOperation.InvalidInputJSON"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_RETURNABLEEXCEPTION = "InternalError.ReturnableException"
//  INTERNALERROR_TSFEXCEPTION = "InternalError.TsfException"
//  INTERNALERROR_VPCEXCEPTION = "InternalError.VpcException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_DUPLICATEPLUGINCONFIG = "InvalidParameterValue.DuplicatePluginConfig"
//  INVALIDPARAMETERVALUE_INVALIDAPIREQUESTCONFIG = "InvalidParameterValue.InvalidApiRequestConfig"
//  INVALIDPARAMETERVALUE_INVALIDBACKENDPATH = "InvalidParameterValue.InvalidBackendPath"
//  INVALIDPARAMETERVALUE_INVALIDCLB = "InvalidParameterValue.InvalidClb"
//  INVALIDPARAMETERVALUE_INVALIDCONDITION = "InvalidParameterValue.InvalidCondition"
//  INVALIDPARAMETERVALUE_INVALIDIPADDRESS = "InvalidParameterValue.InvalidIPAddress"
//  INVALIDPARAMETERVALUE_INVALIDPLUGINCONFIG = "InvalidParameterValue.InvalidPluginConfig"
//  INVALIDPARAMETERVALUE_INVALIDPORT = "InvalidParameterValue.InvalidPort"
//  INVALIDPARAMETERVALUE_INVALIDREQUESTPARAMETERS = "InvalidParameterValue.InvalidRequestParameters"
//  INVALIDPARAMETERVALUE_INVALIDSCFCONFIG = "InvalidParameterValue.InvalidScfConfig"
//  INVALIDPARAMETERVALUE_INVALIDSERVICECONFIG = "InvalidParameterValue.InvalidServiceConfig"
//  INVALIDPARAMETERVALUE_INVALIDSERVICEPARAM = "InvalidParameterValue.InvalidServiceParam"
//  INVALIDPARAMETERVALUE_INVALIDSERVICETYPE = "InvalidParameterValue.InvalidServiceType"
//  INVALIDPARAMETERVALUE_INVALIDTSFCONFIG = "InvalidParameterValue.InvalidTsfConfig"
//  INVALIDPARAMETERVALUE_INVALIDUPSTREAM = "InvalidParameterValue.InvalidUpstream"
//  INVALIDPARAMETERVALUE_INVALIDURL = "InvalidParameterValue.InvalidUrl"
//  INVALIDPARAMETERVALUE_INVALIDVPCCONFIG = "InvalidParameterValue.InvalidVpcConfig"
//  INVALIDPARAMETERVALUE_LENGTHEXCEEDED = "InvalidParameterValue.LengthExceeded"
//  INVALIDPARAMETERVALUE_NOTINOPTIONS = "InvalidParameterValue.NotInOptions"
//  INVALIDPARAMETERVALUE_RANGEEXCEEDED = "InvalidParameterValue.RangeExceeded"
//  MISSINGPARAMETER_BACKENDSPECIFICPARAM = "MissingParameter.BackendSpecificParam"
//  MISSINGPARAMETER_PLUGINCONFIG = "MissingParameter.PluginConfig"
//  RESOURCENOTFOUND_INVALIDPLUGIN = "ResourceNotFound.InvalidPlugin"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_UNPACKERROR = "UnsupportedOperation.UnpackError"
func (c *Client) ModifyPluginWithContext(ctx context.Context, request *ModifyPluginRequest) (response *ModifyPluginResponse, err error) {
    if request == nil {
        request = NewModifyPluginRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPlugin require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPluginResponse()
    err = c.Send(request, response)
    return
}

func NewModifyServiceRequest() (request *ModifyServiceRequest) {
    request = &ModifyServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifyService")
    
    
    return
}

func NewModifyServiceResponse() (response *ModifyServiceResponse) {
    response = &ModifyServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyService
// This API is used to modify the relevant information of a service. After a service is created, its name, description, and service type can be modified.
//
// error code that may be returned:
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_NOTINOPTIONS = "InvalidParameterValue.NotInOptions"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_MODIFYNETTYPE = "UnsupportedOperation.ModifyNetType"
//  UNSUPPORTEDOPERATION_REDUCENETTYPES = "UnsupportedOperation.ReduceNetTypes"
//  UNSUPPORTEDOPERATION_UNPACKERROR = "UnsupportedOperation.UnpackError"
func (c *Client) ModifyService(request *ModifyServiceRequest) (response *ModifyServiceResponse, err error) {
    return c.ModifyServiceWithContext(context.Background(), request)
}

// ModifyService
// This API is used to modify the relevant information of a service. After a service is created, its name, description, and service type can be modified.
//
// error code that may be returned:
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_NOTINOPTIONS = "InvalidParameterValue.NotInOptions"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_MODIFYNETTYPE = "UnsupportedOperation.ModifyNetType"
//  UNSUPPORTEDOPERATION_REDUCENETTYPES = "UnsupportedOperation.ReduceNetTypes"
//  UNSUPPORTEDOPERATION_UNPACKERROR = "UnsupportedOperation.UnpackError"
func (c *Client) ModifyServiceWithContext(ctx context.Context, request *ModifyServiceRequest) (response *ModifyServiceResponse, err error) {
    if request == nil {
        request = NewModifyServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyService require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyServiceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyServiceEnvironmentStrategyRequest() (request *ModifyServiceEnvironmentStrategyRequest) {
    request = &ModifyServiceEnvironmentStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifyServiceEnvironmentStrategy")
    
    
    return
}

func NewModifyServiceEnvironmentStrategyResponse() (response *ModifyServiceEnvironmentStrategyResponse) {
    response = &ModifyServiceEnvironmentStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyServiceEnvironmentStrategy
// This API is used to modify a service throttling policy.
//
// error code that may be returned:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_INVALIDREQUESTPARAMETERS = "InvalidParameterValue.InvalidRequestParameters"
//  INVALIDPARAMETERVALUE_RANGEEXCEEDED = "InvalidParameterValue.RangeExceeded"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) ModifyServiceEnvironmentStrategy(request *ModifyServiceEnvironmentStrategyRequest) (response *ModifyServiceEnvironmentStrategyResponse, err error) {
    return c.ModifyServiceEnvironmentStrategyWithContext(context.Background(), request)
}

// ModifyServiceEnvironmentStrategy
// This API is used to modify a service throttling policy.
//
// error code that may be returned:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_INVALIDREQUESTPARAMETERS = "InvalidParameterValue.InvalidRequestParameters"
//  INVALIDPARAMETERVALUE_RANGEEXCEEDED = "InvalidParameterValue.RangeExceeded"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) ModifyServiceEnvironmentStrategyWithContext(ctx context.Context, request *ModifyServiceEnvironmentStrategyRequest) (response *ModifyServiceEnvironmentStrategyResponse, err error) {
    if request == nil {
        request = NewModifyServiceEnvironmentStrategyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyServiceEnvironmentStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyServiceEnvironmentStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewModifySubDomainRequest() (request *ModifySubDomainRequest) {
    request = &ModifySubDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifySubDomain")
    
    
    return
}

func NewModifySubDomainResponse() (response *ModifySubDomainResponse) {
    response = &ModifySubDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySubDomain
// This API is used to modify the path mapping in the custom domain name settings of a service. The path mapping rule can be modified before it is bound to the custom domain name.
//
// error code that may be returned:
//  FAILEDOPERATION_CERTIFICATEIDBINDERROR = "FailedOperation.CertificateIdBindError"
//  FAILEDOPERATION_CERTIFICATEIDERROR = "FailedOperation.CertificateIdError"
//  FAILEDOPERATION_PATHMAPPINGSETERROR = "FailedOperation.PathMappingSetError"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  FAILEDOPERATION_SERVICENOTEXIST = "FailedOperation.ServiceNotExist"
//  FAILEDOPERATION_SETCUSTOMPATHMAPPINGERROR = "FailedOperation.SetCustomPathMappingError"
//  FAILEDOPERATION_SUBDOMAINFORMATERROR = "FailedOperation.SubDomainFormatError"
//  FAILEDOPERATION_UNKNOWNPROTOCOLTYPEERROR = "FailedOperation.UnknownProtocolTypeError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_FORCEHTTPS = "UnsupportedOperation.ForceHttps"
//  UNSUPPORTEDOPERATION_INVALIDSERVICETRADE = "UnsupportedOperation.InvalidServiceTrade"
func (c *Client) ModifySubDomain(request *ModifySubDomainRequest) (response *ModifySubDomainResponse, err error) {
    return c.ModifySubDomainWithContext(context.Background(), request)
}

// ModifySubDomain
// This API is used to modify the path mapping in the custom domain name settings of a service. The path mapping rule can be modified before it is bound to the custom domain name.
//
// error code that may be returned:
//  FAILEDOPERATION_CERTIFICATEIDBINDERROR = "FailedOperation.CertificateIdBindError"
//  FAILEDOPERATION_CERTIFICATEIDERROR = "FailedOperation.CertificateIdError"
//  FAILEDOPERATION_PATHMAPPINGSETERROR = "FailedOperation.PathMappingSetError"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  FAILEDOPERATION_SERVICENOTEXIST = "FailedOperation.ServiceNotExist"
//  FAILEDOPERATION_SETCUSTOMPATHMAPPINGERROR = "FailedOperation.SetCustomPathMappingError"
//  FAILEDOPERATION_SUBDOMAINFORMATERROR = "FailedOperation.SubDomainFormatError"
//  FAILEDOPERATION_UNKNOWNPROTOCOLTYPEERROR = "FailedOperation.UnknownProtocolTypeError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_FORCEHTTPS = "UnsupportedOperation.ForceHttps"
//  UNSUPPORTEDOPERATION_INVALIDSERVICETRADE = "UnsupportedOperation.InvalidServiceTrade"
func (c *Client) ModifySubDomainWithContext(ctx context.Context, request *ModifySubDomainRequest) (response *ModifySubDomainResponse, err error) {
    if request == nil {
        request = NewModifySubDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySubDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySubDomainResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUpstreamRequest() (request *ModifyUpstreamRequest) {
    request = &ModifyUpstreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifyUpstream")
    
    
    return
}

func NewModifyUpstreamResponse() (response *ModifyUpstreamResponse) {
    response = &ModifyUpstreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyUpstream
// This API is used to modify an upstream.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_OPERATEUPSTREAM = "FailedOperation.OperateUpstream"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDUPSTREAM = "InvalidParameterValue.InvalidUpstream"
//  INVALIDPARAMETERVALUE_INVALIDVPCCONFIG = "InvalidParameterValue.InvalidVpcConfig"
//  INVALIDPARAMETERVALUE_NOTINOPTIONS = "InvalidParameterValue.NotInOptions"
//  INVALIDPARAMETERVALUE_RANGEEXCEEDED = "InvalidParameterValue.RangeExceeded"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyUpstream(request *ModifyUpstreamRequest) (response *ModifyUpstreamResponse, err error) {
    return c.ModifyUpstreamWithContext(context.Background(), request)
}

// ModifyUpstream
// This API is used to modify an upstream.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_OPERATEUPSTREAM = "FailedOperation.OperateUpstream"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDUPSTREAM = "InvalidParameterValue.InvalidUpstream"
//  INVALIDPARAMETERVALUE_INVALIDVPCCONFIG = "InvalidParameterValue.InvalidVpcConfig"
//  INVALIDPARAMETERVALUE_NOTINOPTIONS = "InvalidParameterValue.NotInOptions"
//  INVALIDPARAMETERVALUE_RANGEEXCEEDED = "InvalidParameterValue.RangeExceeded"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyUpstreamWithContext(ctx context.Context, request *ModifyUpstreamRequest) (response *ModifyUpstreamResponse, err error) {
    if request == nil {
        request = NewModifyUpstreamRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUpstream require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUpstreamResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUsagePlanRequest() (request *ModifyUsagePlanRequest) {
    request = &ModifyUsagePlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifyUsagePlan")
    
    
    return
}

func NewModifyUsagePlanResponse() (response *ModifyUsagePlanResponse) {
    response = &ModifyUsagePlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyUsagePlan
// This API is used to modify the name, description, and QPS of a usage plan.
//
// error code that may be returned:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
func (c *Client) ModifyUsagePlan(request *ModifyUsagePlanRequest) (response *ModifyUsagePlanResponse, err error) {
    return c.ModifyUsagePlanWithContext(context.Background(), request)
}

// ModifyUsagePlan
// This API is used to modify the name, description, and QPS of a usage plan.
//
// error code that may be returned:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
func (c *Client) ModifyUsagePlanWithContext(ctx context.Context, request *ModifyUsagePlanRequest) (response *ModifyUsagePlanResponse, err error) {
    if request == nil {
        request = NewModifyUsagePlanRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUsagePlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUsagePlanResponse()
    err = c.Send(request, response)
    return
}

func NewReleaseServiceRequest() (request *ReleaseServiceRequest) {
    request = &ReleaseServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "ReleaseService")
    
    
    return
}

func NewReleaseServiceResponse() (response *ReleaseServiceResponse) {
    response = &ReleaseServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReleaseService
// This API is used to publish a service.
//
// An API Gateway service can only be called when it is published to an environment and takes effect after creation. This API is used to publish a service to an environment such as the `release` environment.
//
// error code that may be returned:
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  FAILEDOPERATION_TASKALREADYEXIST = "FailedOperation.TaskAlreadyExist"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  LIMITEXCEEDED_REQUESTLIMITEXCEEDED = "LimitExceeded.RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_UNPACKERROR = "UnsupportedOperation.UnpackError"
func (c *Client) ReleaseService(request *ReleaseServiceRequest) (response *ReleaseServiceResponse, err error) {
    return c.ReleaseServiceWithContext(context.Background(), request)
}

// ReleaseService
// This API is used to publish a service.
//
// An API Gateway service can only be called when it is published to an environment and takes effect after creation. This API is used to publish a service to an environment such as the `release` environment.
//
// error code that may be returned:
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  FAILEDOPERATION_TASKALREADYEXIST = "FailedOperation.TaskAlreadyExist"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  LIMITEXCEEDED_REQUESTLIMITEXCEEDED = "LimitExceeded.RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_UNPACKERROR = "UnsupportedOperation.UnpackError"
func (c *Client) ReleaseServiceWithContext(ctx context.Context, request *ReleaseServiceRequest) (response *ReleaseServiceResponse, err error) {
    if request == nil {
        request = NewReleaseServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReleaseService require credential")
    }

    request.SetContext(ctx)
    
    response = NewReleaseServiceResponse()
    err = c.Send(request, response)
    return
}

func NewResetAPIDocPasswordRequest() (request *ResetAPIDocPasswordRequest) {
    request = &ResetAPIDocPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "ResetAPIDocPassword")
    
    
    return
}

func NewResetAPIDocPasswordResponse() (response *ResetAPIDocPasswordResponse) {
    response = &ResetAPIDocPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetAPIDocPassword
// This API is used to reset the password of an API document.
//
// error code that may be returned:
//  FAILEDOPERATION_CODINGERROR = "FailedOperation.CodingError"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDAPIDOC = "ResourceNotFound.InvalidApiDoc"
func (c *Client) ResetAPIDocPassword(request *ResetAPIDocPasswordRequest) (response *ResetAPIDocPasswordResponse, err error) {
    return c.ResetAPIDocPasswordWithContext(context.Background(), request)
}

// ResetAPIDocPassword
// This API is used to reset the password of an API document.
//
// error code that may be returned:
//  FAILEDOPERATION_CODINGERROR = "FailedOperation.CodingError"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDAPIDOC = "ResourceNotFound.InvalidApiDoc"
func (c *Client) ResetAPIDocPasswordWithContext(ctx context.Context, request *ResetAPIDocPasswordRequest) (response *ResetAPIDocPasswordResponse, err error) {
    if request == nil {
        request = NewResetAPIDocPasswordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetAPIDocPassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetAPIDocPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewUnBindEnvironmentRequest() (request *UnBindEnvironmentRequest) {
    request = &UnBindEnvironmentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "UnBindEnvironment")
    
    
    return
}

func NewUnBindEnvironmentResponse() (response *UnBindEnvironmentResponse) {
    response = &UnBindEnvironmentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UnBindEnvironment
// This API is used to unbind a usage plan from a specified environment.
//
// error code that may be returned:
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDUNBINDENVIRONMENT = "UnsupportedOperation.UnsupportedUnBindEnvironment"
func (c *Client) UnBindEnvironment(request *UnBindEnvironmentRequest) (response *UnBindEnvironmentResponse, err error) {
    return c.UnBindEnvironmentWithContext(context.Background(), request)
}

// UnBindEnvironment
// This API is used to unbind a usage plan from a specified environment.
//
// error code that may be returned:
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDUNBINDENVIRONMENT = "UnsupportedOperation.UnsupportedUnBindEnvironment"
func (c *Client) UnBindEnvironmentWithContext(ctx context.Context, request *UnBindEnvironmentRequest) (response *UnBindEnvironmentResponse, err error) {
    if request == nil {
        request = NewUnBindEnvironmentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnBindEnvironment require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnBindEnvironmentResponse()
    err = c.Send(request, response)
    return
}

func NewUnBindIPStrategyRequest() (request *UnBindIPStrategyRequest) {
    request = &UnBindIPStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "UnBindIPStrategy")
    
    
    return
}

func NewUnBindIPStrategyResponse() (response *UnBindIPStrategyResponse) {
    response = &UnBindIPStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UnBindIPStrategy
// This API is used to unbind an IP policy from a service.
//
// error code that may be returned:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
func (c *Client) UnBindIPStrategy(request *UnBindIPStrategyRequest) (response *UnBindIPStrategyResponse, err error) {
    return c.UnBindIPStrategyWithContext(context.Background(), request)
}

// UnBindIPStrategy
// This API is used to unbind an IP policy from a service.
//
// error code that may be returned:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
func (c *Client) UnBindIPStrategyWithContext(ctx context.Context, request *UnBindIPStrategyRequest) (response *UnBindIPStrategyResponse, err error) {
    if request == nil {
        request = NewUnBindIPStrategyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnBindIPStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnBindIPStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewUnBindSecretIdsRequest() (request *UnBindSecretIdsRequest) {
    request = &UnBindSecretIdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "UnBindSecretIds")
    
    
    return
}

func NewUnBindSecretIdsResponse() (response *UnBindSecretIdsResponse) {
    response = &UnBindSecretIdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UnBindSecretIds
// This API is used to unbind a key from a usage plan.
//
// error code that may be returned:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  RESOURCENOTFOUND_INVALIDACCESSKEYID = "ResourceNotFound.InvalidAccessKeyId"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
//  UNSUPPORTEDOPERATION_REQUESTPOSTERROR = "UnsupportedOperation.RequestPostError"
func (c *Client) UnBindSecretIds(request *UnBindSecretIdsRequest) (response *UnBindSecretIdsResponse, err error) {
    return c.UnBindSecretIdsWithContext(context.Background(), request)
}

// UnBindSecretIds
// This API is used to unbind a key from a usage plan.
//
// error code that may be returned:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  RESOURCENOTFOUND_INVALIDACCESSKEYID = "ResourceNotFound.InvalidAccessKeyId"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
//  UNSUPPORTEDOPERATION_REQUESTPOSTERROR = "UnsupportedOperation.RequestPostError"
func (c *Client) UnBindSecretIdsWithContext(ctx context.Context, request *UnBindSecretIdsRequest) (response *UnBindSecretIdsResponse, err error) {
    if request == nil {
        request = NewUnBindSecretIdsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnBindSecretIds require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnBindSecretIdsResponse()
    err = c.Send(request, response)
    return
}

func NewUnBindSubDomainRequest() (request *UnBindSubDomainRequest) {
    request = &UnBindSubDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "UnBindSubDomain")
    
    
    return
}

func NewUnBindSubDomainResponse() (response *UnBindSubDomainResponse) {
    response = &UnBindSubDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UnBindSubDomain
// This API is used to unbind a custom domain name.
//
// After binding a custom domain name to a service by using API Gateway, you can use this API to unbind it.
//
// error code that may be returned:
//  FAILEDOPERATION_DOMAINNOTBINDSERVICE = "FailedOperation.DomainNotBindService"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  FAILEDOPERATION_SERVICENOTEXIST = "FailedOperation.ServiceNotExist"
//  FAILEDOPERATION_SUBDOMAINFORMATERROR = "FailedOperation.SubDomainFormatError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_INVALIDSTATUS = "UnsupportedOperation.InvalidStatus"
//  UNSUPPORTEDOPERATION_UNPACKERROR = "UnsupportedOperation.UnpackError"
func (c *Client) UnBindSubDomain(request *UnBindSubDomainRequest) (response *UnBindSubDomainResponse, err error) {
    return c.UnBindSubDomainWithContext(context.Background(), request)
}

// UnBindSubDomain
// This API is used to unbind a custom domain name.
//
// After binding a custom domain name to a service by using API Gateway, you can use this API to unbind it.
//
// error code that may be returned:
//  FAILEDOPERATION_DOMAINNOTBINDSERVICE = "FailedOperation.DomainNotBindService"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  FAILEDOPERATION_SERVICENOTEXIST = "FailedOperation.ServiceNotExist"
//  FAILEDOPERATION_SUBDOMAINFORMATERROR = "FailedOperation.SubDomainFormatError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_INVALIDSTATUS = "UnsupportedOperation.InvalidStatus"
//  UNSUPPORTEDOPERATION_UNPACKERROR = "UnsupportedOperation.UnpackError"
func (c *Client) UnBindSubDomainWithContext(ctx context.Context, request *UnBindSubDomainRequest) (response *UnBindSubDomainResponse, err error) {
    if request == nil {
        request = NewUnBindSubDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnBindSubDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnBindSubDomainResponse()
    err = c.Send(request, response)
    return
}

func NewUnReleaseServiceRequest() (request *UnReleaseServiceRequest) {
    request = &UnReleaseServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "UnReleaseService")
    
    
    return
}

func NewUnReleaseServiceResponse() (response *UnReleaseServiceResponse) {
    response = &UnReleaseServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UnReleaseService
// This API is used to deactivate a service.
//
// Only after a service is published to an environment can its APIs be called. You can call this API to deactivate a service in the release environment. Once deactivated, the service cannot be called.
//
// error code that may be returned:
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  LIMITEXCEEDED_REQUESTLIMITEXCEEDED = "LimitExceeded.RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_UNPACKERROR = "UnsupportedOperation.UnpackError"
func (c *Client) UnReleaseService(request *UnReleaseServiceRequest) (response *UnReleaseServiceResponse, err error) {
    return c.UnReleaseServiceWithContext(context.Background(), request)
}

// UnReleaseService
// This API is used to deactivate a service.
//
// Only after a service is published to an environment can its APIs be called. You can call this API to deactivate a service in the release environment. Once deactivated, the service cannot be called.
//
// error code that may be returned:
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  LIMITEXCEEDED_REQUESTLIMITEXCEEDED = "LimitExceeded.RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_UNPACKERROR = "UnsupportedOperation.UnpackError"
func (c *Client) UnReleaseServiceWithContext(ctx context.Context, request *UnReleaseServiceRequest) (response *UnReleaseServiceResponse, err error) {
    if request == nil {
        request = NewUnReleaseServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnReleaseService require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnReleaseServiceResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindApiAppRequest() (request *UnbindApiAppRequest) {
    request = &UnbindApiAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "UnbindApiApp")
    
    
    return
}

func NewUnbindApiAppResponse() (response *UnbindApiAppResponse) {
    response = &UnbindApiAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UnbindApiApp
// This API is used to unbind an application from an API.
//
// error code that may be returned:
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_RETURNABLEEXCEPTION = "InternalError.ReturnableException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDENV = "InvalidParameterValue.InvalidEnv"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDAPIAPP = "ResourceNotFound.InvalidApiApp"
//  UNSUPPORTEDOPERATION_RESOURCEASSOCIATED = "UnsupportedOperation.ResourceAssociated"
//  UNSUPPORTEDOPERATION_RESOURCEUNASSOCIATED = "UnsupportedOperation.ResourceUnassociated"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDBINDENVIRONMENT = "UnsupportedOperation.UnsupportedBindEnvironment"
func (c *Client) UnbindApiApp(request *UnbindApiAppRequest) (response *UnbindApiAppResponse, err error) {
    return c.UnbindApiAppWithContext(context.Background(), request)
}

// UnbindApiApp
// This API is used to unbind an application from an API.
//
// error code that may be returned:
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_RETURNABLEEXCEPTION = "InternalError.ReturnableException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDENV = "InvalidParameterValue.InvalidEnv"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDAPIAPP = "ResourceNotFound.InvalidApiApp"
//  UNSUPPORTEDOPERATION_RESOURCEASSOCIATED = "UnsupportedOperation.ResourceAssociated"
//  UNSUPPORTEDOPERATION_RESOURCEUNASSOCIATED = "UnsupportedOperation.ResourceUnassociated"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDBINDENVIRONMENT = "UnsupportedOperation.UnsupportedBindEnvironment"
func (c *Client) UnbindApiAppWithContext(ctx context.Context, request *UnbindApiAppRequest) (response *UnbindApiAppResponse, err error) {
    if request == nil {
        request = NewUnbindApiAppRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindApiApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindApiAppResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateApiAppKeyRequest() (request *UpdateApiAppKeyRequest) {
    request = &UpdateApiAppKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "UpdateApiAppKey")
    
    
    return
}

func NewUpdateApiAppKeyResponse() (response *UpdateApiAppKeyResponse) {
    response = &UpdateApiAppKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateApiAppKey
// This API is used to update an application key.
//
// error code that may be returned:
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  FAILEDOPERATION_APIINOPERATION = "FailedOperation.ApiInOperation"
//  FAILEDOPERATION_GETROLEERROR = "FailedOperation.GetRoleError"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_SCFEXCEPTION = "InternalError.ScfException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDAPIAPP = "ResourceNotFound.InvalidApiApp"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_INVALIDSERVICETRADE = "UnsupportedOperation.InvalidServiceTrade"
//  UNSUPPORTEDOPERATION_RESOURCEASSOCIATED = "UnsupportedOperation.ResourceAssociated"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDDELETEAPI = "UnsupportedOperation.UnsupportedDeleteApi"
func (c *Client) UpdateApiAppKey(request *UpdateApiAppKeyRequest) (response *UpdateApiAppKeyResponse, err error) {
    return c.UpdateApiAppKeyWithContext(context.Background(), request)
}

// UpdateApiAppKey
// This API is used to update an application key.
//
// error code that may be returned:
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  FAILEDOPERATION_APIINOPERATION = "FailedOperation.ApiInOperation"
//  FAILEDOPERATION_GETROLEERROR = "FailedOperation.GetRoleError"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_SCFEXCEPTION = "InternalError.ScfException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDAPIAPP = "ResourceNotFound.InvalidApiApp"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_INVALIDSERVICETRADE = "UnsupportedOperation.InvalidServiceTrade"
//  UNSUPPORTEDOPERATION_RESOURCEASSOCIATED = "UnsupportedOperation.ResourceAssociated"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDDELETEAPI = "UnsupportedOperation.UnsupportedDeleteApi"
func (c *Client) UpdateApiAppKeyWithContext(ctx context.Context, request *UpdateApiAppKeyRequest) (response *UpdateApiAppKeyResponse, err error) {
    if request == nil {
        request = NewUpdateApiAppKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateApiAppKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateApiAppKeyResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateApiKeyRequest() (request *UpdateApiKeyRequest) {
    request = &UpdateApiKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "UpdateApiKey")
    
    
    return
}

func NewUpdateApiKeyResponse() (response *UpdateApiKeyResponse) {
    response = &UpdateApiKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateApiKey
// This API is used to update a created API key pair.
//
// error code that may be returned:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDACCESSKEYID = "ResourceNotFound.InvalidAccessKeyId"
//  UNSUPPORTEDOPERATION_UINNOTINWHITELIST = "UnsupportedOperation.UinNotInWhiteList"
func (c *Client) UpdateApiKey(request *UpdateApiKeyRequest) (response *UpdateApiKeyResponse, err error) {
    return c.UpdateApiKeyWithContext(context.Background(), request)
}

// UpdateApiKey
// This API is used to update a created API key pair.
//
// error code that may be returned:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDACCESSKEYID = "ResourceNotFound.InvalidAccessKeyId"
//  UNSUPPORTEDOPERATION_UINNOTINWHITELIST = "UnsupportedOperation.UinNotInWhiteList"
func (c *Client) UpdateApiKeyWithContext(ctx context.Context, request *UpdateApiKeyRequest) (response *UpdateApiKeyResponse, err error) {
    if request == nil {
        request = NewUpdateApiKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateApiKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateApiKeyResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateServiceRequest() (request *UpdateServiceRequest) {
    request = &UpdateServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "UpdateService")
    
    
    return
}

func NewUpdateServiceResponse() (response *UpdateServiceResponse) {
    response = &UpdateServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateService
// u200dThis API is used to switch the running version of a service published in an environment to a specified version. After you create a service by using API Gateway and publish it to an environment, multiple versions will be generated during development. In this case, you can call this API to switch versions.
//
// error code that may be returned:
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_INVALIDSERVICETRADE = "UnsupportedOperation.InvalidServiceTrade"
func (c *Client) UpdateService(request *UpdateServiceRequest) (response *UpdateServiceResponse, err error) {
    return c.UpdateServiceWithContext(context.Background(), request)
}

// UpdateService
// u200dThis API is used to switch the running version of a service published in an environment to a specified version. After you create a service by using API Gateway and publish it to an environment, multiple versions will be generated during development. In this case, you can call this API to switch versions.
//
// error code that may be returned:
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_INVALIDSERVICETRADE = "UnsupportedOperation.InvalidServiceTrade"
func (c *Client) UpdateServiceWithContext(ctx context.Context, request *UpdateServiceRequest) (response *UpdateServiceResponse, err error) {
    if request == nil {
        request = NewUpdateServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateService require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateServiceResponse()
    err = c.Send(request, response)
    return
}
