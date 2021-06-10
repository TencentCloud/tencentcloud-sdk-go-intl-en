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

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
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
//  LIMITEXCEEDED_REQUESTLIMITEXCEEDED = "LimitExceeded.RequestLimitExceeded"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDBINDENVIRONMENT = "UnsupportedOperation.UnsupportedBindEnvironment"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDUNBINDENVIRONMENT = "UnsupportedOperation.UnsupportedUnBindEnvironment"
func (c *Client) BindEnvironment(request *BindEnvironmentRequest) (response *BindEnvironmentResponse, err error) {
    if request == nil {
        request = NewBindEnvironmentRequest()
    }
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
//  LIMITEXCEEDED_REQUESTLIMITEXCEEDED = "LimitExceeded.RequestLimitExceeded"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDIPSTRATEGY = "ResourceNotFound.InvalidIPStrategy"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) BindIPStrategy(request *BindIPStrategyRequest) (response *BindIPStrategyResponse, err error) {
    if request == nil {
        request = NewBindIPStrategyRequest()
    }
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
//  RESOURCENOTFOUND_INVALIDACCESSKEYID = "ResourceNotFound.InvalidAccessKeyId"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
//  UNSUPPORTEDOPERATION_ALREADYBINDUSAGEPLAN = "UnsupportedOperation.AlreadyBindUsagePlan"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDBINDAPIKEY = "UnsupportedOperation.UnsupportedBindApiKey"
func (c *Client) BindSecretIds(request *BindSecretIdsRequest) (response *BindSecretIdsResponse, err error) {
    if request == nil {
        request = NewBindSecretIdsRequest()
    }
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
//  FAILEDOPERATION_DOMAINNEEDBEIAN = "FailedOperation.DomainNeedBeian"
//  FAILEDOPERATION_DOMAINRESOLVEERROR = "FailedOperation.DomainResolveError"
//  FAILEDOPERATION_FORMATERROR = "FailedOperation.FormatError"
//  FAILEDOPERATION_ISDEFAULTMAPPING = "FailedOperation.IsDefaultMapping"
//  FAILEDOPERATION_NETSUBDOMAINERROR = "FailedOperation.NetSubDomainError"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SUBDOMAINFORMATERROR = "FailedOperation.SubDomainFormatError"
//  INTERNALERROR = "InternalError"
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
    if request == nil {
        request = NewBindSubDomainRequest()
    }
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
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDAPIDOC = "ResourceNotFound.InvalidApiDoc"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) BuildAPIDoc(request *BuildAPIDocRequest) (response *BuildAPIDocResponse, err error) {
    if request == nil {
        request = NewBuildAPIDocRequest()
    }
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
    if request == nil {
        request = NewCreateAPIDocRequest()
    }
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
//  FAILEDOPERATION_EBERROR = "FailedOperation.EbError"
//  FAILEDOPERATION_GETROLEERROR = "FailedOperation.GetRoleError"
//  FAILEDOPERATION_SCFERROR = "FailedOperation.ScfError"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_SCFEXCEPTION = "InternalError.ScfException"
//  INTERNALERROR_TSFEXCEPTION = "InternalError.TsfException"
//  INTERNALERROR_VPCEXCEPTION = "InternalError.VpcException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDAPIBUSINESSTYPE = "InvalidParameterValue.InvalidApiBusinessType"
//  INVALIDPARAMETERVALUE_INVALIDAPIREQUESTCONFIG = "InvalidParameterValue.InvalidApiRequestConfig"
//  INVALIDPARAMETERVALUE_INVALIDAPITYPE = "InvalidParameterValue.InvalidApiType"
//  INVALIDPARAMETERVALUE_INVALIDCLB = "InvalidParameterValue.InvalidClb"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICKEY = "InvalidParameterValue.InvalidPublicKey"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_INVALIDREQUESTPARAMETERS = "InvalidParameterValue.InvalidRequestParameters"
//  INVALIDPARAMETERVALUE_INVALIDSCFCONFIG = "InvalidParameterValue.InvalidScfConfig"
//  INVALIDPARAMETERVALUE_INVALIDSERVICECONFIG = "InvalidParameterValue.InvalidServiceConfig"
//  INVALIDPARAMETERVALUE_INVALIDSERVICEMOCKRETURNMESSAGE = "InvalidParameterValue.InvalidServiceMockReturnMessage"
//  INVALIDPARAMETERVALUE_INVALIDSERVICEPARAMETERS = "InvalidParameterValue.InvalidServiceParameters"
//  INVALIDPARAMETERVALUE_INVALIDTSFCONFIG = "InvalidParameterValue.InvalidTsfConfig"
//  INVALIDPARAMETERVALUE_INVALIDURL = "InvalidParameterValue.InvalidUrl"
//  INVALIDPARAMETERVALUE_PARAMETERNOTMATCH = "InvalidParameterValue.ParameterNotMatch"
//  INVALIDPARAMETERVALUE_RANGEEXCEEDED = "InvalidParameterValue.RangeExceeded"
//  LIMITEXCEEDED_APICOUNTLIMITEXCEEDED = "LimitExceeded.ApiCountLimitExceeded"
//  RESOURCENOTFOUND_INVALIDOAUTHAPI = "ResourceNotFound.InvalidOauthApi"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_INVALIDENDPOINTTYPE = "UnsupportedOperation.InvalidEndpointType"
func (c *Client) CreateApi(request *CreateApiRequest) (response *CreateApiResponse, err error) {
    if request == nil {
        request = NewCreateApiRequest()
    }
    response = NewCreateApiResponse()
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
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  LIMITEXCEEDED_APIKEYCOUNTLIMITEXCEEDED = "LimitExceeded.ApiKeyCountLimitExceeded"
func (c *Client) CreateApiKey(request *CreateApiKeyRequest) (response *CreateApiKeyResponse, err error) {
    if request == nil {
        request = NewCreateApiKeyRequest()
    }
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
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  LIMITEXCEEDED_IPSTRATEGYLIMITEXCEEDED = "LimitExceeded.IpStrategyLimitExceeded"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) CreateIPStrategy(request *CreateIPStrategyRequest) (response *CreateIPStrategyResponse, err error) {
    if request == nil {
        request = NewCreateIPStrategyRequest()
    }
    response = NewCreateIPStrategyResponse()
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
// The maximum unit in API Gateway is service. Multiple APIs can be created in one service, and each service has a default domain name for users to call. You can also bind your own custom domain name to a service.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTEXIST = "FailedOperation.InstanceNotExist"
//  FAILEDOPERATION_TAGBINDSERVICEERROR = "FailedOperation.TagBindServiceError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_VPCEXCEPTION = "InternalError.VpcException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
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
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) CreateService(request *CreateServiceRequest) (response *CreateServiceResponse, err error) {
    if request == nil {
        request = NewCreateServiceRequest()
    }
    response = NewCreateServiceResponse()
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
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  LIMITEXCEEDED_USAGEPLANLIMITEXCEEDED = "LimitExceeded.UsagePlanLimitExceeded"
func (c *Client) CreateUsagePlan(request *CreateUsagePlanRequest) (response *CreateUsagePlanResponse, err error) {
    if request == nil {
        request = NewCreateUsagePlanRequest()
    }
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
    if request == nil {
        request = NewDeleteAPIDocRequest()
    }
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
//  UNSUPPORTEDOPERATION_INVALIDSERVICETRADE = "UnsupportedOperation.InvalidServiceTrade"
func (c *Client) DeleteApi(request *DeleteApiRequest) (response *DeleteApiResponse, err error) {
    if request == nil {
        request = NewDeleteApiRequest()
    }
    response = NewDeleteApiResponse()
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
    if request == nil {
        request = NewDeleteApiKeyRequest()
    }
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
    if request == nil {
        request = NewDeleteIPStrategyRequest()
    }
    response = NewDeleteIPStrategyResponse()
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
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDDELETESERVICE = "UnsupportedOperation.UnsupportedDeleteService"
func (c *Client) DeleteService(request *DeleteServiceRequest) (response *DeleteServiceResponse, err error) {
    if request == nil {
        request = NewDeleteServiceRequest()
    }
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
//  UNSUPPORTEDOPERATION_INVALIDSERVICETRADE = "UnsupportedOperation.InvalidServiceTrade"
func (c *Client) DeleteServiceSubDomainMapping(request *DeleteServiceSubDomainMappingRequest) (response *DeleteServiceSubDomainMappingResponse, err error) {
    if request == nil {
        request = NewDeleteServiceSubDomainMappingRequest()
    }
    response = NewDeleteServiceSubDomainMappingResponse()
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
    if request == nil {
        request = NewDeleteUsagePlanRequest()
    }
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
//  UNSUPPORTEDOPERATION_NOUSAGEPLANENV = "UnsupportedOperation.NoUsagePlanEnv"
func (c *Client) DemoteServiceUsagePlan(request *DemoteServiceUsagePlanRequest) (response *DemoteServiceUsagePlanResponse, err error) {
    if request == nil {
        request = NewDemoteServiceUsagePlanRequest()
    }
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
    if request == nil {
        request = NewDescribeAPIDocDetailRequest()
    }
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
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DescribeAPIDocs(request *DescribeAPIDocsRequest) (response *DescribeAPIDocsResponse, err error) {
    if request == nil {
        request = NewDescribeAPIDocsRequest()
    }
    response = NewDescribeAPIDocsResponse()
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
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DescribeApi(request *DescribeApiRequest) (response *DescribeApiResponse, err error) {
    if request == nil {
        request = NewDescribeApiRequest()
    }
    response = NewDescribeApiResponse()
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
    if request == nil {
        request = NewDescribeApiEnvironmentStrategyRequest()
    }
    response = NewDescribeApiEnvironmentStrategyResponse()
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
//  RESOURCENOTFOUND_INVALIDACCESSKEYID = "ResourceNotFound.InvalidAccessKeyId"
func (c *Client) DescribeApiKey(request *DescribeApiKeyRequest) (response *DescribeApiKeyResponse, err error) {
    if request == nil {
        request = NewDescribeApiKeyRequest()
    }
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
// This API is used to query the list of keys.
//
// If you have created multiple API key pairs, you can use this API to query the information of one or more keys. This API does not display the `secretKey`.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
func (c *Client) DescribeApiKeysStatus(request *DescribeApiKeysStatusRequest) (response *DescribeApiKeysStatusResponse, err error) {
    if request == nil {
        request = NewDescribeApiKeysStatusRequest()
    }
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
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeApiUsagePlan(request *DescribeApiUsagePlanRequest) (response *DescribeApiUsagePlanResponse, err error) {
    if request == nil {
        request = NewDescribeApiUsagePlanRequest()
    }
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
func (c *Client) DescribeApisStatus(request *DescribeApisStatusRequest) (response *DescribeApisStatusResponse, err error) {
    if request == nil {
        request = NewDescribeApisStatusRequest()
    }
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
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  RESOURCENOTFOUND_INVALIDIPSTRATEGY = "ResourceNotFound.InvalidIPStrategy"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeIPStrategy(request *DescribeIPStrategyRequest) (response *DescribeIPStrategyResponse, err error) {
    if request == nil {
        request = NewDescribeIPStrategyRequest()
    }
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
//  RESOURCENOTFOUND_INVALIDIPSTRATEGY = "ResourceNotFound.InvalidIPStrategy"
func (c *Client) DescribeIPStrategyApisStatus(request *DescribeIPStrategyApisStatusRequest) (response *DescribeIPStrategyApisStatusResponse, err error) {
    if request == nil {
        request = NewDescribeIPStrategyApisStatusRequest()
    }
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
    if request == nil {
        request = NewDescribeIPStrategysStatusRequest()
    }
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
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
func (c *Client) DescribeLogSearch(request *DescribeLogSearchRequest) (response *DescribeLogSearchResponse, err error) {
    if request == nil {
        request = NewDescribeLogSearchRequest()
    }
    response = NewDescribeLogSearchResponse()
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
func (c *Client) DescribeService(request *DescribeServiceRequest) (response *DescribeServiceResponse, err error) {
    if request == nil {
        request = NewDescribeServiceRequest()
    }
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
    if request == nil {
        request = NewDescribeServiceEnvironmentListRequest()
    }
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
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeServiceEnvironmentReleaseHistory(request *DescribeServiceEnvironmentReleaseHistoryRequest) (response *DescribeServiceEnvironmentReleaseHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeServiceEnvironmentReleaseHistoryRequest()
    }
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
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeServiceEnvironmentStrategy(request *DescribeServiceEnvironmentStrategyRequest) (response *DescribeServiceEnvironmentStrategyResponse, err error) {
    if request == nil {
        request = NewDescribeServiceEnvironmentStrategyRequest()
    }
    response = NewDescribeServiceEnvironmentStrategyResponse()
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
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeServiceReleaseVersion(request *DescribeServiceReleaseVersionRequest) (response *DescribeServiceReleaseVersionResponse, err error) {
    if request == nil {
        request = NewDescribeServiceReleaseVersionRequest()
    }
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
func (c *Client) DescribeServiceSubDomainMappings(request *DescribeServiceSubDomainMappingsRequest) (response *DescribeServiceSubDomainMappingsResponse, err error) {
    if request == nil {
        request = NewDescribeServiceSubDomainMappingsRequest()
    }
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
    if request == nil {
        request = NewDescribeServiceSubDomainsRequest()
    }
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
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeServiceUsagePlan(request *DescribeServiceUsagePlanRequest) (response *DescribeServiceUsagePlanResponse, err error) {
    if request == nil {
        request = NewDescribeServiceUsagePlanRequest()
    }
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
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
func (c *Client) DescribeServicesStatus(request *DescribeServicesStatusRequest) (response *DescribeServicesStatusResponse, err error) {
    if request == nil {
        request = NewDescribeServicesStatusRequest()
    }
    response = NewDescribeServicesStatusResponse()
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
    if request == nil {
        request = NewDescribeUsagePlanRequest()
    }
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
    if request == nil {
        request = NewDescribeUsagePlanEnvironmentsRequest()
    }
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
    if request == nil {
        request = NewDescribeUsagePlanSecretIdsRequest()
    }
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
    if request == nil {
        request = NewDescribeUsagePlansStatusRequest()
    }
    response = NewDescribeUsagePlansStatusResponse()
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
    if request == nil {
        request = NewDisableApiKeyRequest()
    }
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
    if request == nil {
        request = NewEnableApiKeyRequest()
    }
    response = NewEnableApiKeyResponse()
    err = c.Send(request, response)
    return
}

func NewGenerateApiDocumentRequest() (request *GenerateApiDocumentRequest) {
    request = &GenerateApiDocumentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "GenerateApiDocument")
    return
}

func NewGenerateApiDocumentResponse() (response *GenerateApiDocumentResponse) {
    response = &GenerateApiDocumentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GenerateApiDocument
// This API is used to automatically generate API documents and SDKs. One document and one SDK will be generated for each environment under each service, respectively.
//
// error code that may be returned:
//  FAILEDOPERATION_GENERATEAPIDOCUMENTERROR = "FailedOperation.GenerateApiDocumentError"
//  FAILEDOPERATION_SERVICENOTEXIST = "FailedOperation.ServiceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDGENLANGUAGE = "InvalidParameterValue.InvalidGenLanguage"
//  INVALIDPARAMETERVALUE_NOTINOPTIONS = "InvalidParameterValue.NotInOptions"
func (c *Client) GenerateApiDocument(request *GenerateApiDocumentRequest) (response *GenerateApiDocumentResponse, err error) {
    if request == nil {
        request = NewGenerateApiDocumentRequest()
    }
    response = NewGenerateApiDocumentResponse()
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
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDAPIDOC = "ResourceNotFound.InvalidApiDoc"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) ModifyAPIDoc(request *ModifyAPIDocRequest) (response *ModifyAPIDocResponse, err error) {
    if request == nil {
        request = NewModifyAPIDocRequest()
    }
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
//  FAILEDOPERATION_EBERROR = "FailedOperation.EbError"
//  FAILEDOPERATION_GETROLEERROR = "FailedOperation.GetRoleError"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_SCFEXCEPTION = "InternalError.ScfException"
//  INTERNALERROR_TSFEXCEPTION = "InternalError.TsfException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
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
//  INVALIDPARAMETERVALUE_INVALIDSERVICECONFIG = "InvalidParameterValue.InvalidServiceConfig"
//  INVALIDPARAMETERVALUE_INVALIDSERVICEPARAM = "InvalidParameterValue.InvalidServiceParam"
//  INVALIDPARAMETERVALUE_INVALIDSERVICETYPE = "InvalidParameterValue.InvalidServiceType"
//  INVALIDPARAMETERVALUE_INVALIDURL = "InvalidParameterValue.InvalidUrl"
//  INVALIDPARAMETERVALUE_INVALIDWSMETHOD = "InvalidParameterValue.InvalidWSMethod"
//  INVALIDPARAMETERVALUE_LENGTHEXCEEDED = "InvalidParameterValue.LengthExceeded"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_PARAMETERNOTMATCH = "InvalidParameterValue.ParameterNotMatch"
//  INVALIDPARAMETERVALUE_RANGEEXCEEDED = "InvalidParameterValue.RangeExceeded"
//  LIMITEXCEEDED_APICOUNTLIMITEXCEEDED = "LimitExceeded.ApiCountLimitExceeded"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDOAUTHAPI = "ResourceNotFound.InvalidOauthApi"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_INVALIDENDPOINTTYPE = "UnsupportedOperation.InvalidEndpointType"
//  UNSUPPORTEDOPERATION_MODIFYPROTOCOL = "UnsupportedOperation.ModifyProtocol"
//  UNSUPPORTEDOPERATION_RESOURCEASSOCIATED = "UnsupportedOperation.ResourceAssociated"
func (c *Client) ModifyApi(request *ModifyApiRequest) (response *ModifyApiResponse, err error) {
    if request == nil {
        request = NewModifyApiRequest()
    }
    response = NewModifyApiResponse()
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
//  INVALIDPARAMETERVALUE_RANGEEXCEEDED = "InvalidParameterValue.RangeExceeded"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
func (c *Client) ModifyApiEnvironmentStrategy(request *ModifyApiEnvironmentStrategyRequest) (response *ModifyApiEnvironmentStrategyResponse, err error) {
    if request == nil {
        request = NewModifyApiEnvironmentStrategyRequest()
    }
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
//  INVALIDPARAMETERVALUE_INVALIDAPIBUSINESSTYPE = "InvalidParameterValue.InvalidApiBusinessType"
//  INVALIDPARAMETERVALUE_NOTHINGMODIFYFOROAUTH = "InvalidParameterValue.NothingModifyForOauth"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
func (c *Client) ModifyApiIncrement(request *ModifyApiIncrementRequest) (response *ModifyApiIncrementResponse, err error) {
    if request == nil {
        request = NewModifyApiIncrementRequest()
    }
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
    if request == nil {
        request = NewModifyIPStrategyRequest()
    }
    response = NewModifyIPStrategyResponse()
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
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_REDUCENETTYPES = "UnsupportedOperation.ReduceNetTypes"
func (c *Client) ModifyService(request *ModifyServiceRequest) (response *ModifyServiceResponse, err error) {
    if request == nil {
        request = NewModifyServiceRequest()
    }
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
//  INVALIDPARAMETERVALUE_RANGEEXCEEDED = "InvalidParameterValue.RangeExceeded"
func (c *Client) ModifyServiceEnvironmentStrategy(request *ModifyServiceEnvironmentStrategyRequest) (response *ModifyServiceEnvironmentStrategyResponse, err error) {
    if request == nil {
        request = NewModifyServiceEnvironmentStrategyRequest()
    }
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
//  FAILEDOPERATION_SERVICENOTEXIST = "FailedOperation.ServiceNotExist"
//  FAILEDOPERATION_SETCUSTOMPATHMAPPINGERROR = "FailedOperation.SetCustomPathMappingError"
//  FAILEDOPERATION_SUBDOMAINFORMATERROR = "FailedOperation.SubDomainFormatError"
//  FAILEDOPERATION_UNKNOWNPROTOCOLTYPEERROR = "FailedOperation.UnknownProtocolTypeError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  UNSUPPORTEDOPERATION_FORCEHTTPS = "UnsupportedOperation.ForceHttps"
//  UNSUPPORTEDOPERATION_INVALIDSERVICETRADE = "UnsupportedOperation.InvalidServiceTrade"
func (c *Client) ModifySubDomain(request *ModifySubDomainRequest) (response *ModifySubDomainResponse, err error) {
    if request == nil {
        request = NewModifySubDomainRequest()
    }
    response = NewModifySubDomainResponse()
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
    if request == nil {
        request = NewModifyUsagePlanRequest()
    }
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
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  LIMITEXCEEDED_REQUESTLIMITEXCEEDED = "LimitExceeded.RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ReleaseService(request *ReleaseServiceRequest) (response *ReleaseServiceResponse, err error) {
    if request == nil {
        request = NewReleaseServiceRequest()
    }
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
//  RESOURCENOTFOUND_INVALIDAPIDOC = "ResourceNotFound.InvalidApiDoc"
func (c *Client) ResetAPIDocPassword(request *ResetAPIDocPasswordRequest) (response *ResetAPIDocPasswordResponse, err error) {
    if request == nil {
        request = NewResetAPIDocPasswordRequest()
    }
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
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDUNBINDENVIRONMENT = "UnsupportedOperation.UnsupportedUnBindEnvironment"
func (c *Client) UnBindEnvironment(request *UnBindEnvironmentRequest) (response *UnBindEnvironmentResponse, err error) {
    if request == nil {
        request = NewUnBindEnvironmentRequest()
    }
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
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDUNBINDENVIRONMENT = "UnsupportedOperation.UnsupportedUnBindEnvironment"
func (c *Client) UnBindIPStrategy(request *UnBindIPStrategyRequest) (response *UnBindIPStrategyResponse, err error) {
    if request == nil {
        request = NewUnBindIPStrategyRequest()
    }
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
func (c *Client) UnBindSecretIds(request *UnBindSecretIdsRequest) (response *UnBindSecretIdsResponse, err error) {
    if request == nil {
        request = NewUnBindSecretIdsRequest()
    }
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
//  FAILEDOPERATION_SERVICENOTEXIST = "FailedOperation.ServiceNotExist"
//  FAILEDOPERATION_SUBDOMAINFORMATERROR = "FailedOperation.SubDomainFormatError"
//  UNSUPPORTEDOPERATION_INVALIDSTATUS = "UnsupportedOperation.InvalidStatus"
func (c *Client) UnBindSubDomain(request *UnBindSubDomainRequest) (response *UnBindSubDomainResponse, err error) {
    if request == nil {
        request = NewUnBindSubDomainRequest()
    }
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
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  LIMITEXCEEDED_REQUESTLIMITEXCEEDED = "LimitExceeded.RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) UnReleaseService(request *UnReleaseServiceRequest) (response *UnReleaseServiceResponse, err error) {
    if request == nil {
        request = NewUnReleaseServiceRequest()
    }
    response = NewUnReleaseServiceResponse()
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
//  RESOURCENOTFOUND_INVALIDACCESSKEYID = "ResourceNotFound.InvalidAccessKeyId"
//  UNSUPPORTEDOPERATION_UINNOTINWHITELIST = "UnsupportedOperation.UinNotInWhiteList"
func (c *Client) UpdateApiKey(request *UpdateApiKeyRequest) (response *UpdateApiKeyResponse, err error) {
    if request == nil {
        request = NewUpdateApiKeyRequest()
    }
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
// This API is used to switch the running version of a service published in an environment to a specified version. After you create a service by using API Gateway and publish it to an environment, multiple versions will be generated during development. In this case, you can call this API to switch versions.
//
// error code that may be returned:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_INVALIDSERVICETRADE = "UnsupportedOperation.InvalidServiceTrade"
func (c *Client) UpdateService(request *UpdateServiceRequest) (response *UpdateServiceResponse, err error) {
    if request == nil {
        request = NewUpdateServiceRequest()
    }
    response = NewUpdateServiceResponse()
    err = c.Send(request, response)
    return
}
