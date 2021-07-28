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

package v20180416

import (
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2018-04-16"

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


func NewCopyFunctionRequest() (request *CopyFunctionRequest) {
    request = &CopyFunctionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "CopyFunction")
    return
}

func NewCopyFunctionResponse() (response *CopyFunctionResponse) {
    response = &CopyFunctionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CopyFunction
// This API is used to replicate a function. You can store the replicated function in a specified Region and Namespace.
//
// Note: This API **does not** replicate the following objects or attributes of the function:
//
// 1. Function trigger
//
// 2. Versions other than $LATEST
//
// 3. CLS target of the logs configured in the function
//
// 
//
// You can manually configure the function after replication as required.
//
// error code that may be returned:
//  FAILEDOPERATION_COPYFAILED = "FailedOperation.CopyFailed"
//  FAILEDOPERATION_COPYFUNCTION = "FailedOperation.CopyFunction"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXCEPTION = "InternalError.Exception"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETERVALUE_CODE = "InvalidParameterValue.Code"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_ENVIRONMENT = "InvalidParameterValue.Environment"
//  INVALIDPARAMETERVALUE_FUNCTIONNAME = "InvalidParameterValue.FunctionName"
//  INVALIDPARAMETERVALUE_HANDLER = "InvalidParameterValue.Handler"
//  INVALIDPARAMETERVALUE_RUNTIME = "InvalidParameterValue.Runtime"
//  LIMITEXCEEDED_FUNCTION = "LimitExceeded.Function"
//  LIMITEXCEEDED_MEMORY = "LimitExceeded.Memory"
//  LIMITEXCEEDED_TIMEOUT = "LimitExceeded.Timeout"
//  RESOURCEINUSE_FUNCTIONNAME = "ResourceInUse.FunctionName"
//  RESOURCENOTFOUND_FUNCTION = "ResourceNotFound.Function"
//  RESOURCENOTFOUND_FUNCTIONNAME = "ResourceNotFound.FunctionName"
//  RESOURCENOTFOUND_NAMESPACE = "ResourceNotFound.Namespace"
//  UNAUTHORIZEDOPERATION_CAM = "UnauthorizedOperation.CAM"
//  UNAUTHORIZEDOPERATION_REGION = "UnauthorizedOperation.Region"
//  UNSUPPORTEDOPERATION_REGION = "UnsupportedOperation.Region"
func (c *Client) CopyFunction(request *CopyFunctionRequest) (response *CopyFunctionResponse, err error) {
    if request == nil {
        request = NewCopyFunctionRequest()
    }
    response = NewCopyFunctionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAliasRequest() (request *CreateAliasRequest) {
    request = &CreateAliasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "CreateAlias")
    return
}

func NewCreateAliasResponse() (response *CreateAliasResponse) {
    response = &CreateAliasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAlias
// This API is used to create an alias for a function version. You can use the alias to mark a specific function version such as DEV/RELEASE. You can also modify the version pointed to by the alias at any time.
//
// An alias must point to a master version and can point to an additional version at the same time. If you specify an alias when invoking a function, the request will be sent to the versions pointed to by the alias. You can configure the ratio between the master version and additional version during request sending.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATEALIAS = "FailedOperation.CreateAlias"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_ROUTINGCONFIG = "InvalidParameter.RoutingConfig"
//  INVALIDPARAMETERVALUE_ADDITIONALVERSIONWEIGHTS = "InvalidParameterValue.AdditionalVersionWeights"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_ROUTINGCONFIG = "InvalidParameterValue.RoutingConfig"
//  LIMITEXCEEDED_ALIAS = "LimitExceeded.Alias"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_ALIAS = "ResourceInUse.Alias"
//  RESOURCENOTFOUND_FUNCTION = "ResourceNotFound.Function"
//  RESOURCENOTFOUND_FUNCTIONNAME = "ResourceNotFound.FunctionName"
//  RESOURCENOTFOUND_FUNCTIONVERSION = "ResourceNotFound.FunctionVersion"
//  UNAUTHORIZEDOPERATION_CAM = "UnauthorizedOperation.CAM"
func (c *Client) CreateAlias(request *CreateAliasRequest) (response *CreateAliasResponse, err error) {
    if request == nil {
        request = NewCreateAliasRequest()
    }
    response = NewCreateAliasResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFunctionRequest() (request *CreateFunctionRequest) {
    request = &CreateFunctionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "CreateFunction")
    return
}

func NewCreateFunctionResponse() (response *CreateFunctionResponse) {
    response = &CreateFunctionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateFunction
// This API is used to create a function based on the input parameters.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATEFUNCTION = "FailedOperation.CreateFunction"
//  FAILEDOPERATION_NAMESPACE = "FailedOperation.Namespace"
//  FAILEDOPERATION_OPENSERVICE = "FailedOperation.OpenService"
//  FAILEDOPERATION_QCSROLENOTFOUND = "FailedOperation.QcsRoleNotFound"
//  FAILEDOPERATION_TOTALCONCURRENCYMEMORYINPROGRESS = "FailedOperation.TotalConcurrencyMemoryInProgress"
//  FAILEDOPERATION_UNOPENEDSERVICE = "FailedOperation.UnOpenedService"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PAYLOAD = "InvalidParameter.Payload"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACTION = "InvalidParameterValue.Action"
//  INVALIDPARAMETERVALUE_CFSPARAMETERDUPLICATE = "InvalidParameterValue.CfsParameterDuplicate"
//  INVALIDPARAMETERVALUE_CFSPARAMETERERROR = "InvalidParameterValue.CfsParameterError"
//  INVALIDPARAMETERVALUE_CFSSTRUCTIONERROR = "InvalidParameterValue.CfsStructionError"
//  INVALIDPARAMETERVALUE_CLS = "InvalidParameterValue.Cls"
//  INVALIDPARAMETERVALUE_CODE = "InvalidParameterValue.Code"
//  INVALIDPARAMETERVALUE_CODESECRET = "InvalidParameterValue.CodeSecret"
//  INVALIDPARAMETERVALUE_CODESOURCE = "InvalidParameterValue.CodeSource"
//  INVALIDPARAMETERVALUE_COS = "InvalidParameterValue.Cos"
//  INVALIDPARAMETERVALUE_COSBUCKETNAME = "InvalidParameterValue.CosBucketName"
//  INVALIDPARAMETERVALUE_COSBUCKETREGION = "InvalidParameterValue.CosBucketRegion"
//  INVALIDPARAMETERVALUE_COSOBJECTNAME = "InvalidParameterValue.CosObjectName"
//  INVALIDPARAMETERVALUE_DEADLETTERCONFIG = "InvalidParameterValue.DeadLetterConfig"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_EIPCONFIG = "InvalidParameterValue.EipConfig"
//  INVALIDPARAMETERVALUE_ENVIRONMENT = "InvalidParameterValue.Environment"
//  INVALIDPARAMETERVALUE_ENVIRONMENTEXCEEDEDLIMIT = "InvalidParameterValue.EnvironmentExceededLimit"
//  INVALIDPARAMETERVALUE_ENVIRONMENTSYSTEMPROTECT = "InvalidParameterValue.EnvironmentSystemProtect"
//  INVALIDPARAMETERVALUE_FUNCTIONNAME = "InvalidParameterValue.FunctionName"
//  INVALIDPARAMETERVALUE_GITCOMMITID = "InvalidParameterValue.GitCommitId"
//  INVALIDPARAMETERVALUE_GITURL = "InvalidParameterValue.GitUrl"
//  INVALIDPARAMETERVALUE_HANDLER = "InvalidParameterValue.Handler"
//  INVALIDPARAMETERVALUE_LAYERS = "InvalidParameterValue.Layers"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MEMORY = "InvalidParameterValue.Memory"
//  INVALIDPARAMETERVALUE_MEMORYSIZE = "InvalidParameterValue.MemorySize"
//  INVALIDPARAMETERVALUE_NAMESPACE = "InvalidParameterValue.Namespace"
//  INVALIDPARAMETERVALUE_PUBLICNETCONFIG = "InvalidParameterValue.PublicNetConfig"
//  INVALIDPARAMETERVALUE_RUNTIME = "InvalidParameterValue.Runtime"
//  INVALIDPARAMETERVALUE_STAMP = "InvalidParameterValue.Stamp"
//  INVALIDPARAMETERVALUE_TEMPCOSOBJECTNAME = "InvalidParameterValue.TempCosObjectName"
//  INVALIDPARAMETERVALUE_TRACEENABLE = "InvalidParameterValue.TraceEnable"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  INVALIDPARAMETERVALUE_VPCNOTSETWHENOPENCFS = "InvalidParameterValue.VpcNotSetWhenOpenCfs"
//  INVALIDPARAMETERVALUE_ZIPFILE = "InvalidParameterValue.ZipFile"
//  INVALIDPARAMETERVALUE_ZIPFILEBASE64BINASCIIERROR = "InvalidParameterValue.ZipFileBase64BinasciiError"
//  LIMITEXCEEDED_EIP = "LimitExceeded.Eip"
//  LIMITEXCEEDED_FUNCTION = "LimitExceeded.Function"
//  LIMITEXCEEDED_INITTIMEOUT = "LimitExceeded.InitTimeout"
//  LIMITEXCEEDED_MEMORY = "LimitExceeded.Memory"
//  LIMITEXCEEDED_TIMEOUT = "LimitExceeded.Timeout"
//  MISSINGPARAMETER_CODE = "MissingParameter.Code"
//  MISSINGPARAMETER_RUNTIME = "MissingParameter.Runtime"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_FUNCTION = "ResourceInUse.Function"
//  RESOURCEINUSE_FUNCTIONNAME = "ResourceInUse.FunctionName"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CFSVPCNOTMATCH = "ResourceNotFound.CfsVpcNotMatch"
//  RESOURCENOTFOUND_CMQ = "ResourceNotFound.Cmq"
//  RESOURCENOTFOUND_DEMO = "ResourceNotFound.Demo"
//  RESOURCENOTFOUND_GETCFSMOUNTINSERROR = "ResourceNotFound.GetCfsMountInsError"
//  RESOURCENOTFOUND_GETCFSNOTMATCH = "ResourceNotFound.GetCfsNotMatch"
//  RESOURCENOTFOUND_LAYER = "ResourceNotFound.Layer"
//  RESOURCENOTFOUND_NAMESPACE = "ResourceNotFound.Namespace"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
//  RESOURCENOTFOUND_VPC = "ResourceNotFound.Vpc"
//  RESOURCEUNAVAILABLE_NAMESPACE = "ResourceUnavailable.Namespace"
//  UNAUTHORIZEDOPERATION_CAM = "UnauthorizedOperation.CAM"
//  UNAUTHORIZEDOPERATION_REGION = "UnauthorizedOperation.Region"
//  UNAUTHORIZEDOPERATION_ROLE = "UnauthorizedOperation.Role"
//  UNAUTHORIZEDOPERATION_TEMPCOSAPPID = "UnauthorizedOperation.TempCosAppid"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateFunction(request *CreateFunctionRequest) (response *CreateFunctionResponse, err error) {
    if request == nil {
        request = NewCreateFunctionRequest()
    }
    response = NewCreateFunctionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNamespaceRequest() (request *CreateNamespaceRequest) {
    request = &CreateNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "CreateNamespace")
    return
}

func NewCreateNamespaceResponse() (response *CreateNamespaceResponse) {
    response = &CreateNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateNamespace
// This API is used to create a namespace based on the input parameters.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATENAMESPACE = "FailedOperation.CreateNamespace"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEFAULTNAMESPACE = "InvalidParameterValue.DefaultNamespace"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_NAMESPACE = "InvalidParameterValue.Namespace"
//  INVALIDPARAMETERVALUE_NAMESPACEINVALID = "InvalidParameterValue.NamespaceInvalid"
//  INVALIDPARAMETERVALUE_STAMP = "InvalidParameterValue.Stamp"
//  LIMITEXCEEDED_NAMESPACE = "LimitExceeded.Namespace"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_NAMESPACE = "ResourceInUse.Namespace"
//  UNAUTHORIZEDOPERATION_CAM = "UnauthorizedOperation.CAM"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateNamespace(request *CreateNamespaceRequest) (response *CreateNamespaceResponse, err error) {
    if request == nil {
        request = NewCreateNamespaceRequest()
    }
    response = NewCreateNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTriggerRequest() (request *CreateTriggerRequest) {
    request = &CreateTriggerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "CreateTrigger")
    return
}

func NewCreateTriggerResponse() (response *CreateTriggerResponse) {
    response = &CreateTriggerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTrigger
// This API is used to create a trigger based on the input parameters.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APIGATEWAY = "FailedOperation.ApiGateway"
//  FAILEDOPERATION_APIGW = "FailedOperation.Apigw"
//  FAILEDOPERATION_COS = "FailedOperation.Cos"
//  FAILEDOPERATION_CREATETRIGGER = "FailedOperation.CreateTrigger"
//  INTERNALERROR_APIGATEWAY = "InternalError.ApiGateway"
//  INTERNALERROR_CKAFKA = "InternalError.Ckafka"
//  INTERNALERROR_CMQ = "InternalError.Cmq"
//  INTERNALERROR_COS = "InternalError.Cos"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_APIGATEWAY = "InvalidParameterValue.ApiGateway"
//  INVALIDPARAMETERVALUE_CDN = "InvalidParameterValue.Cdn"
//  INVALIDPARAMETERVALUE_CKAFKA = "InvalidParameterValue.Ckafka"
//  INVALIDPARAMETERVALUE_COS = "InvalidParameterValue.Cos"
//  INVALIDPARAMETERVALUE_CUSTOMARGUMENT = "InvalidParameterValue.CustomArgument"
//  INVALIDPARAMETERVALUE_ENABLE = "InvalidParameterValue.Enable"
//  INVALIDPARAMETERVALUE_FUNCTIONNAME = "InvalidParameterValue.FunctionName"
//  INVALIDPARAMETERVALUE_SERVICENAME = "InvalidParameterValue.ServiceName"
//  INVALIDPARAMETERVALUE_TRIGGERDESC = "InvalidParameterValue.TriggerDesc"
//  INVALIDPARAMETERVALUE_TRIGGERNAME = "InvalidParameterValue.TriggerName"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  INVALIDPARAMETERVALUE_ZIPFILE = "InvalidParameterValue.ZipFile"
//  LIMITEXCEEDED_CDN = "LimitExceeded.Cdn"
//  LIMITEXCEEDED_FUNCTIONONTOPIC = "LimitExceeded.FunctionOnTopic"
//  LIMITEXCEEDED_TRIGGER = "LimitExceeded.Trigger"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_CDN = "ResourceInUse.Cdn"
//  RESOURCEINUSE_CMQ = "ResourceInUse.Cmq"
//  RESOURCEINUSE_COS = "ResourceInUse.Cos"
//  RESOURCEINUSE_TRIGGER = "ResourceInUse.Trigger"
//  RESOURCEINUSE_TRIGGERNAME = "ResourceInUse.TriggerName"
//  RESOURCEINSUFFICIENT_COS = "ResourceInsufficient.COS"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CDN = "ResourceNotFound.Cdn"
//  RESOURCENOTFOUND_CKAFKA = "ResourceNotFound.Ckafka"
//  RESOURCENOTFOUND_CMQ = "ResourceNotFound.Cmq"
//  RESOURCENOTFOUND_COS = "ResourceNotFound.Cos"
//  RESOURCENOTFOUND_FUNCTION = "ResourceNotFound.Function"
//  RESOURCENOTFOUND_FUNCTIONNAME = "ResourceNotFound.FunctionName"
//  RESOURCENOTFOUND_FUNCTIONVERSION = "ResourceNotFound.FunctionVersion"
//  RESOURCENOTFOUND_NAMESPACE = "ResourceNotFound.Namespace"
//  RESOURCENOTFOUND_QUALIFIER = "ResourceNotFound.Qualifier"
//  RESOURCENOTFOUND_VERSION = "ResourceNotFound.Version"
//  UNAUTHORIZEDOPERATION_CAM = "UnauthorizedOperation.CAM"
//  UNAUTHORIZEDOPERATION_CREATETRIGGER = "UnauthorizedOperation.CreateTrigger"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CDN = "UnsupportedOperation.Cdn"
//  UNSUPPORTEDOPERATION_COS = "UnsupportedOperation.Cos"
//  UNSUPPORTEDOPERATION_TRIGGER = "UnsupportedOperation.Trigger"
func (c *Client) CreateTrigger(request *CreateTriggerRequest) (response *CreateTriggerResponse, err error) {
    if request == nil {
        request = NewCreateTriggerRequest()
    }
    response = NewCreateTriggerResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAliasRequest() (request *DeleteAliasRequest) {
    request = &DeleteAliasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "DeleteAlias")
    return
}

func NewDeleteAliasResponse() (response *DeleteAliasResponse) {
    response = &DeleteAliasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAlias
// This API is used to delete an alias of a function version.
//
// error code that may be returned:
//  FAILEDOPERATION_DELETEALIAS = "FailedOperation.DeleteAlias"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  RESOURCENOTFOUND_ALIAS = "ResourceNotFound.Alias"
//  RESOURCENOTFOUND_FUNCTION = "ResourceNotFound.Function"
//  RESOURCENOTFOUND_FUNCTIONNAME = "ResourceNotFound.FunctionName"
func (c *Client) DeleteAlias(request *DeleteAliasRequest) (response *DeleteAliasResponse, err error) {
    if request == nil {
        request = NewDeleteAliasRequest()
    }
    response = NewDeleteAliasResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteFunctionRequest() (request *DeleteFunctionRequest) {
    request = &DeleteFunctionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "DeleteFunction")
    return
}

func NewDeleteFunctionResponse() (response *DeleteFunctionResponse) {
    response = &DeleteFunctionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteFunction
// This API is used to delete a function based on the input parameters.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETEFUNCTION = "FailedOperation.DeleteFunction"
//  FAILEDOPERATION_FUNCTIONNAMESTATUSERROR = "FailedOperation.FunctionNameStatusError"
//  FAILEDOPERATION_FUNCTIONSTATUSERROR = "FailedOperation.FunctionStatusError"
//  FAILEDOPERATION_PROVISIONEDINPROGRESS = "FailedOperation.ProvisionedInProgress"
//  INTERNALERROR_CMQ = "InternalError.Cmq"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_PAYLOAD = "InvalidParameter.Payload"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FUNCTIONNAME = "InvalidParameterValue.FunctionName"
//  INVALIDPARAMETERVALUE_NAMESPACE = "InvalidParameterValue.Namespace"
//  INVALIDPARAMETERVALUE_QUALIFIER = "InvalidParameterValue.Qualifier"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FUNCTION = "ResourceNotFound.Function"
//  RESOURCENOTFOUND_FUNCTIONNAME = "ResourceNotFound.FunctionName"
//  RESOURCENOTFOUND_NAMESPACE = "ResourceNotFound.Namespace"
//  RESOURCENOTFOUND_QUALIFIER = "ResourceNotFound.Qualifier"
//  UNAUTHORIZEDOPERATION_CAM = "UnauthorizedOperation.CAM"
//  UNAUTHORIZEDOPERATION_DELETEFUNCTION = "UnauthorizedOperation.DeleteFunction"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ALIASBIND = "UnsupportedOperation.AliasBind"
func (c *Client) DeleteFunction(request *DeleteFunctionRequest) (response *DeleteFunctionResponse, err error) {
    if request == nil {
        request = NewDeleteFunctionRequest()
    }
    response = NewDeleteFunctionResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLayerVersionRequest() (request *DeleteLayerVersionRequest) {
    request = &DeleteLayerVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "DeleteLayerVersion")
    return
}

func NewDeleteLayerVersionResponse() (response *DeleteLayerVersionResponse) {
    response = &DeleteLayerVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLayerVersion
// This API is used to delete a specified version of a specified layer. The deleted version cannot be associated with a function, but the deletion does not affect functions that are referencing this layer.
//
// error code that may be returned:
//  FAILEDOPERATION_DELETELAYERVERSION = "FailedOperation.DeleteLayerVersion"
//  FAILEDOPERATION_OPERATIONCONFLICT = "FailedOperation.OperationConflict"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  RESOURCEINUSE_LAYERVERSION = "ResourceInUse.LayerVersion"
//  RESOURCENOTFOUND_LAYERVERSION = "ResourceNotFound.LayerVersion"
//  UNAUTHORIZEDOPERATION_CAM = "UnauthorizedOperation.CAM"
func (c *Client) DeleteLayerVersion(request *DeleteLayerVersionRequest) (response *DeleteLayerVersionResponse, err error) {
    if request == nil {
        request = NewDeleteLayerVersionRequest()
    }
    response = NewDeleteLayerVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNamespaceRequest() (request *DeleteNamespaceRequest) {
    request = &DeleteNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "DeleteNamespace")
    return
}

func NewDeleteNamespaceResponse() (response *DeleteNamespaceResponse) {
    response = &DeleteNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteNamespace
// This API is used to delete the specific namespace according to the parameters passed in.
//
// error code that may be returned:
//  FAILEDOPERATION_DELETENAMESPACE = "FailedOperation.DeleteNamespace"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NAMESPACE = "InvalidParameterValue.Namespace"
//  INVALIDPARAMETERVALUE_NAMESPACEINVALID = "InvalidParameterValue.NamespaceInvalid"
//  RESOURCEINUSE_NAMESPACE = "ResourceInUse.Namespace"
//  RESOURCENOTFOUND_NAMESPACE = "ResourceNotFound.Namespace"
//  UNAUTHORIZEDOPERATION_CAM = "UnauthorizedOperation.CAM"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteNamespace(request *DeleteNamespaceRequest) (response *DeleteNamespaceResponse, err error) {
    if request == nil {
        request = NewDeleteNamespaceRequest()
    }
    response = NewDeleteNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteProvisionedConcurrencyConfigRequest() (request *DeleteProvisionedConcurrencyConfigRequest) {
    request = &DeleteProvisionedConcurrencyConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "DeleteProvisionedConcurrencyConfig")
    return
}

func NewDeleteProvisionedConcurrencyConfigResponse() (response *DeleteProvisionedConcurrencyConfigResponse) {
    response = &DeleteProvisionedConcurrencyConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteProvisionedConcurrencyConfig
// This API is used to delete the provisioned concurrency configuration of a function version.
//
// error code that may be returned:
//  FAILEDOPERATION_PROVISIONEDINPROGRESS = "FailedOperation.ProvisionedInProgress"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_FUNCTIONNAME = "InvalidParameterValue.FunctionName"
//  INVALIDPARAMETERVALUE_QUALIFIER = "InvalidParameterValue.Qualifier"
//  RESOURCENOTFOUND_FUNCTION = "ResourceNotFound.Function"
//  RESOURCENOTFOUND_FUNCTIONVERSION = "ResourceNotFound.FunctionVersion"
//  RESOURCENOTFOUND_NAMESPACE = "ResourceNotFound.Namespace"
//  RESOURCENOTFOUND_VERSION = "ResourceNotFound.Version"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteProvisionedConcurrencyConfig(request *DeleteProvisionedConcurrencyConfigRequest) (response *DeleteProvisionedConcurrencyConfigResponse, err error) {
    if request == nil {
        request = NewDeleteProvisionedConcurrencyConfigRequest()
    }
    response = NewDeleteProvisionedConcurrencyConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteReservedConcurrencyConfigRequest() (request *DeleteReservedConcurrencyConfigRequest) {
    request = &DeleteReservedConcurrencyConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "DeleteReservedConcurrencyConfig")
    return
}

func NewDeleteReservedConcurrencyConfigResponse() (response *DeleteReservedConcurrencyConfigResponse) {
    response = &DeleteReservedConcurrencyConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteReservedConcurrencyConfig
// This API is used to delete the reserved concurrency configuration of a function.
//
// error code that may be returned:
//  FAILEDOPERATION_DEBUGMODESTATUS = "FailedOperation.DebugModeStatus"
//  FAILEDOPERATION_RESERVEDINPROGRESS = "FailedOperation.ReservedInProgress"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_FUNCTION = "ResourceNotFound.Function"
//  RESOURCENOTFOUND_NAMESPACE = "ResourceNotFound.Namespace"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteReservedConcurrencyConfig(request *DeleteReservedConcurrencyConfigRequest) (response *DeleteReservedConcurrencyConfigResponse, err error) {
    if request == nil {
        request = NewDeleteReservedConcurrencyConfigRequest()
    }
    response = NewDeleteReservedConcurrencyConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTriggerRequest() (request *DeleteTriggerRequest) {
    request = &DeleteTriggerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "DeleteTrigger")
    return
}

func NewDeleteTriggerResponse() (response *DeleteTriggerResponse) {
    response = &DeleteTriggerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteTrigger
// This API is used to delete an existing trigger based on the input parameters.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATETRIGGER = "FailedOperation.CreateTrigger"
//  FAILEDOPERATION_DELETETRIGGER = "FailedOperation.DeleteTrigger"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_APIGATEWAY = "InvalidParameterValue.ApiGateway"
//  INVALIDPARAMETERVALUE_CDN = "InvalidParameterValue.Cdn"
//  INVALIDPARAMETERVALUE_CMQ = "InvalidParameterValue.Cmq"
//  INVALIDPARAMETERVALUE_COS = "InvalidParameterValue.Cos"
//  INVALIDPARAMETERVALUE_FUNCTIONNAME = "InvalidParameterValue.FunctionName"
//  INVALIDPARAMETERVALUE_TRIGGERDESC = "InvalidParameterValue.TriggerDesc"
//  INVALIDPARAMETERVALUE_TRIGGERNAME = "InvalidParameterValue.TriggerName"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  RESOURCEINUSE_CDN = "ResourceInUse.Cdn"
//  RESOURCEINUSE_CMQ = "ResourceInUse.Cmq"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CDN = "ResourceNotFound.Cdn"
//  RESOURCENOTFOUND_CMQ = "ResourceNotFound.Cmq"
//  RESOURCENOTFOUND_FUNCTION = "ResourceNotFound.Function"
//  RESOURCENOTFOUND_FUNCTIONNAME = "ResourceNotFound.FunctionName"
//  RESOURCENOTFOUND_TIMER = "ResourceNotFound.Timer"
//  RESOURCENOTFOUND_TRIGGER = "ResourceNotFound.Trigger"
//  RESOURCENOTFOUND_VERSION = "ResourceNotFound.Version"
//  UNAUTHORIZEDOPERATION_CAM = "UnauthorizedOperation.CAM"
//  UNAUTHORIZEDOPERATION_DELETETRIGGER = "UnauthorizedOperation.DeleteTrigger"
//  UNSUPPORTEDOPERATION_CDN = "UnsupportedOperation.Cdn"
func (c *Client) DeleteTrigger(request *DeleteTriggerRequest) (response *DeleteTriggerResponse, err error) {
    if request == nil {
        request = NewDeleteTriggerRequest()
    }
    response = NewDeleteTriggerResponse()
    err = c.Send(request, response)
    return
}

func NewGetAccountRequest() (request *GetAccountRequest) {
    request = &GetAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "GetAccount")
    return
}

func NewGetAccountResponse() (response *GetAccountResponse) {
    response = &GetAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetAccount
// This API is used to get the account information.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_STAMP = "InvalidParameterValue.Stamp"
//  UNAUTHORIZEDOPERATION_CAM = "UnauthorizedOperation.CAM"
func (c *Client) GetAccount(request *GetAccountRequest) (response *GetAccountResponse, err error) {
    if request == nil {
        request = NewGetAccountRequest()
    }
    response = NewGetAccountResponse()
    err = c.Send(request, response)
    return
}

func NewGetAliasRequest() (request *GetAliasRequest) {
    request = &GetAliasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "GetAlias")
    return
}

func NewGetAliasResponse() (response *GetAliasResponse) {
    response = &GetAliasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetAlias
// This API is used to get the alias details such as the name, description, version, and routing information.
//
// error code that may be returned:
//  FAILEDOPERATION_GETALIAS = "FailedOperation.GetAlias"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_ROUTINGCONFIG = "InvalidParameter.RoutingConfig"
//  INVALIDPARAMETERVALUE_FUNCTIONNAME = "InvalidParameterValue.FunctionName"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_NAMESPACE = "InvalidParameterValue.Namespace"
//  RESOURCENOTFOUND_ALIAS = "ResourceNotFound.Alias"
//  RESOURCENOTFOUND_FUNCTION = "ResourceNotFound.Function"
//  RESOURCENOTFOUND_FUNCTIONNAME = "ResourceNotFound.FunctionName"
//  RESOURCENOTFOUND_NAMESPACE = "ResourceNotFound.Namespace"
//  UNAUTHORIZEDOPERATION_CAM = "UnauthorizedOperation.CAM"
func (c *Client) GetAlias(request *GetAliasRequest) (response *GetAliasResponse, err error) {
    if request == nil {
        request = NewGetAliasRequest()
    }
    response = NewGetAliasResponse()
    err = c.Send(request, response)
    return
}

func NewGetFunctionRequest() (request *GetFunctionRequest) {
    request = &GetFunctionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "GetFunction")
    return
}

func NewGetFunctionResponse() (response *GetFunctionResponse) {
    response = &GetFunctionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetFunction
// This API is used to obtain function details, such as name, code, handler, associated trigger, and timeout.
//
// error code that may be returned:
//  FAILEDOPERATION_APIGW = "FailedOperation.Apigw"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXCEPTION = "InternalError.Exception"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_PAYLOAD = "InvalidParameter.Payload"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CODESECRET = "InvalidParameterValue.CodeSecret"
//  INVALIDPARAMETERVALUE_FUNCTIONNAME = "InvalidParameterValue.FunctionName"
//  INVALIDPARAMETERVALUE_NAMESPACE = "InvalidParameterValue.Namespace"
//  INVALIDPARAMETERVALUE_QUALIFIER = "InvalidParameterValue.Qualifier"
//  RESOURCENOTFOUND_FUNCTION = "ResourceNotFound.Function"
//  RESOURCENOTFOUND_FUNCTIONNAME = "ResourceNotFound.FunctionName"
//  RESOURCENOTFOUND_NAMESPACE = "ResourceNotFound.Namespace"
//  RESOURCENOTFOUND_VERSION = "ResourceNotFound.Version"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAM = "UnauthorizedOperation.CAM"
//  UNAUTHORIZEDOPERATION_CODESECRET = "UnauthorizedOperation.CodeSecret"
func (c *Client) GetFunction(request *GetFunctionRequest) (response *GetFunctionResponse, err error) {
    if request == nil {
        request = NewGetFunctionRequest()
    }
    response = NewGetFunctionResponse()
    err = c.Send(request, response)
    return
}

func NewGetFunctionAddressRequest() (request *GetFunctionAddressRequest) {
    request = &GetFunctionAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "GetFunctionAddress")
    return
}

func NewGetFunctionAddressResponse() (response *GetFunctionAddressResponse) {
    response = &GetFunctionAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetFunctionAddress
// This API is used to obtain the download address of the function code package.
//
// error code that may be returned:
//  FAILEDOPERATION_FUNCTIONSTATUSERROR = "FailedOperation.FunctionStatusError"
//  FAILEDOPERATION_GETFUNCTIONADDRESS = "FailedOperation.GetFunctionAddress"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_PAYLOAD = "InvalidParameter.Payload"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FUNCTIONNAME = "InvalidParameterValue.FunctionName"
//  INVALIDPARAMETERVALUE_NAMESPACE = "InvalidParameterValue.Namespace"
//  INVALIDPARAMETERVALUE_QUALIFIER = "InvalidParameterValue.Qualifier"
//  RESOURCENOTFOUND_FUNCTION = "ResourceNotFound.Function"
//  RESOURCENOTFOUND_FUNCTIONNAME = "ResourceNotFound.FunctionName"
//  RESOURCENOTFOUND_FUNCTIONVERSION = "ResourceNotFound.FunctionVersion"
//  RESOURCENOTFOUND_NAMESPACE = "ResourceNotFound.Namespace"
//  RESOURCENOTFOUND_VERSION = "ResourceNotFound.Version"
//  UNAUTHORIZEDOPERATION_CAM = "UnauthorizedOperation.CAM"
//  UNAUTHORIZEDOPERATION_CODESECRET = "UnauthorizedOperation.CodeSecret"
func (c *Client) GetFunctionAddress(request *GetFunctionAddressRequest) (response *GetFunctionAddressResponse, err error) {
    if request == nil {
        request = NewGetFunctionAddressRequest()
    }
    response = NewGetFunctionAddressResponse()
    err = c.Send(request, response)
    return
}

func NewGetFunctionEventInvokeConfigRequest() (request *GetFunctionEventInvokeConfigRequest) {
    request = &GetFunctionEventInvokeConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "GetFunctionEventInvokeConfig")
    return
}

func NewGetFunctionEventInvokeConfigResponse() (response *GetFunctionEventInvokeConfigResponse) {
    response = &GetFunctionEventInvokeConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetFunctionEventInvokeConfig
// This API is used to get the async retry configuration of a function, including the number of retry attempts and message retention period.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FUNCTIONVERSIONSTATUSNOTACTIVE = "FailedOperation.FunctionVersionStatusNotActive"
//  FAILEDOPERATION_INSUFFICIENTBALANCE = "FailedOperation.InsufficientBalance"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FUNCTIONNAME = "InvalidParameterValue.FunctionName"
//  INVALIDPARAMETERVALUE_NAMESPACE = "InvalidParameterValue.Namespace"
//  INVALIDPARAMETERVALUE_QUALIFIER = "InvalidParameterValue.Qualifier"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FUNCTION = "ResourceNotFound.Function"
//  RESOURCENOTFOUND_FUNCTIONNAME = "ResourceNotFound.FunctionName"
//  RESOURCENOTFOUND_FUNCTIONVERSION = "ResourceNotFound.FunctionVersion"
//  RESOURCENOTFOUND_NAMESPACE = "ResourceNotFound.Namespace"
//  RESOURCENOTFOUND_VERSION = "ResourceNotFound.Version"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) GetFunctionEventInvokeConfig(request *GetFunctionEventInvokeConfigRequest) (response *GetFunctionEventInvokeConfigResponse, err error) {
    if request == nil {
        request = NewGetFunctionEventInvokeConfigRequest()
    }
    response = NewGetFunctionEventInvokeConfigResponse()
    err = c.Send(request, response)
    return
}

func NewGetFunctionLogsRequest() (request *GetFunctionLogsRequest) {
    request = &GetFunctionLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "GetFunctionLogs")
    return
}

func NewGetFunctionLogsResponse() (response *GetFunctionLogsResponse) {
    response = &GetFunctionLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetFunctionLogs
// This API is used to return function running logs according to the specified log query criteria.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ES = "InternalError.ES"
//  INTERNALERROR_EXCEPTION = "InternalError.Exception"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_PAYLOAD = "InvalidParameter.Payload"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATETIME = "InvalidParameterValue.DateTime"
//  INVALIDPARAMETERVALUE_OFFSET = "InvalidParameterValue.Offset"
//  INVALIDPARAMETERVALUE_ORDER = "InvalidParameterValue.Order"
//  INVALIDPARAMETERVALUE_ORDERBY = "InvalidParameterValue.OrderBy"
//  INVALIDPARAMETERVALUE_RETCODE = "InvalidParameterValue.RetCode"
//  INVALIDPARAMETERVALUE_STARTTIMEORENDTIME = "InvalidParameterValue.StartTimeOrEndTime"
//  LIMITEXCEEDED_OFFSET = "LimitExceeded.Offset"
//  RESOURCENOTFOUND_FUNCTION = "ResourceNotFound.Function"
//  RESOURCENOTFOUND_FUNCTIONNAME = "ResourceNotFound.FunctionName"
//  UNAUTHORIZEDOPERATION_CAM = "UnauthorizedOperation.CAM"
func (c *Client) GetFunctionLogs(request *GetFunctionLogsRequest) (response *GetFunctionLogsResponse, err error) {
    if request == nil {
        request = NewGetFunctionLogsRequest()
    }
    response = NewGetFunctionLogsResponse()
    err = c.Send(request, response)
    return
}

func NewGetLayerVersionRequest() (request *GetLayerVersionRequest) {
    request = &GetLayerVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "GetLayerVersion")
    return
}

func NewGetLayerVersionResponse() (response *GetLayerVersionResponse) {
    response = &GetLayerVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetLayerVersion
// This API is used to get the layer version details, including links used to download files in the layer.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  RESOURCENOTFOUND_LAYERVERSION = "ResourceNotFound.LayerVersion"
//  UNAUTHORIZEDOPERATION_CAM = "UnauthorizedOperation.CAM"
func (c *Client) GetLayerVersion(request *GetLayerVersionRequest) (response *GetLayerVersionResponse, err error) {
    if request == nil {
        request = NewGetLayerVersionRequest()
    }
    response = NewGetLayerVersionResponse()
    err = c.Send(request, response)
    return
}

func NewGetProvisionedConcurrencyConfigRequest() (request *GetProvisionedConcurrencyConfigRequest) {
    request = &GetProvisionedConcurrencyConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "GetProvisionedConcurrencyConfig")
    return
}

func NewGetProvisionedConcurrencyConfigResponse() (response *GetProvisionedConcurrencyConfigResponse) {
    response = &GetProvisionedConcurrencyConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetProvisionedConcurrencyConfig
// This API is used to get the provisioned concurrency details of a function or its specified version.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NAMESPACE = "InvalidParameterValue.Namespace"
//  INVALIDPARAMETERVALUE_QUALIFIER = "InvalidParameterValue.Qualifier"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FUNCTION = "ResourceNotFound.Function"
//  RESOURCENOTFOUND_FUNCTIONVERSION = "ResourceNotFound.FunctionVersion"
//  RESOURCENOTFOUND_NAMESPACE = "ResourceNotFound.Namespace"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) GetProvisionedConcurrencyConfig(request *GetProvisionedConcurrencyConfigRequest) (response *GetProvisionedConcurrencyConfigResponse, err error) {
    if request == nil {
        request = NewGetProvisionedConcurrencyConfigRequest()
    }
    response = NewGetProvisionedConcurrencyConfigResponse()
    err = c.Send(request, response)
    return
}

func NewGetReservedConcurrencyConfigRequest() (request *GetReservedConcurrencyConfigRequest) {
    request = &GetReservedConcurrencyConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "GetReservedConcurrencyConfig")
    return
}

func NewGetReservedConcurrencyConfigResponse() (response *GetReservedConcurrencyConfigResponse) {
    response = &GetReservedConcurrencyConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetReservedConcurrencyConfig
// This API is used to get the reserved concurrency details of a function.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FUNCTIONNAME = "InvalidParameterValue.FunctionName"
//  INVALIDPARAMETERVALUE_NAMESPACE = "InvalidParameterValue.Namespace"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FUNCTION = "ResourceNotFound.Function"
//  RESOURCENOTFOUND_NAMESPACE = "ResourceNotFound.Namespace"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) GetReservedConcurrencyConfig(request *GetReservedConcurrencyConfigRequest) (response *GetReservedConcurrencyConfigResponse, err error) {
    if request == nil {
        request = NewGetReservedConcurrencyConfigRequest()
    }
    response = NewGetReservedConcurrencyConfigResponse()
    err = c.Send(request, response)
    return
}

func NewInvokeRequest() (request *InvokeRequest) {
    request = &InvokeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "Invoke")
    return
}

func NewInvokeResponse() (response *InvokeResponse) {
    response = &InvokeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Invoke
// This API is used to run a function.
//
// error code that may be returned:
//  FAILEDOPERATION_FUNCTIONSTATUSERROR = "FailedOperation.FunctionStatusError"
//  FAILEDOPERATION_INVOKEFUNCTION = "FailedOperation.InvokeFunction"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_FUNCTIONNAME = "InvalidParameter.FunctionName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAM = "InvalidParameterValue.Param"
//  RESOURCENOTFOUND_FUNCTION = "ResourceNotFound.Function"
//  RESOURCENOTFOUND_FUNCTIONNAME = "ResourceNotFound.FunctionName"
//  RESOURCEUNAVAILABLE_INSUFFICIENTBALANCE = "ResourceUnavailable.InsufficientBalance"
//  UNAUTHORIZEDOPERATION_CAM = "UnauthorizedOperation.CAM"
func (c *Client) Invoke(request *InvokeRequest) (response *InvokeResponse, err error) {
    if request == nil {
        request = NewInvokeRequest()
    }
    response = NewInvokeResponse()
    err = c.Send(request, response)
    return
}

func NewInvokeFunctionRequest() (request *InvokeFunctionRequest) {
    request = &InvokeFunctionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "InvokeFunction")
    return
}

func NewInvokeFunctionResponse() (response *InvokeFunctionResponse) {
    response = &InvokeFunctionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InvokeFunction
//  This API is used to invoke functions synchronously.
//
// error code that may be returned:
//  FAILEDOPERATION_FUNCTIONSTATUSERROR = "FailedOperation.FunctionStatusError"
//  FAILEDOPERATION_INVOKEFUNCTION = "FailedOperation.InvokeFunction"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_FUNCTIONNAME = "InvalidParameter.FunctionName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAM = "InvalidParameterValue.Param"
//  RESOURCENOTFOUND_FUNCTION = "ResourceNotFound.Function"
//  RESOURCENOTFOUND_FUNCTIONNAME = "ResourceNotFound.FunctionName"
//  RESOURCENOTFOUND_QUALIFIER = "ResourceNotFound.Qualifier"
//  RESOURCEUNAVAILABLE_INSUFFICIENTBALANCE = "ResourceUnavailable.InsufficientBalance"
//  UNAUTHORIZEDOPERATION_CAM = "UnauthorizedOperation.CAM"
func (c *Client) InvokeFunction(request *InvokeFunctionRequest) (response *InvokeFunctionResponse, err error) {
    if request == nil {
        request = NewInvokeFunctionRequest()
    }
    response = NewInvokeFunctionResponse()
    err = c.Send(request, response)
    return
}

func NewListAliasesRequest() (request *ListAliasesRequest) {
    request = &ListAliasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "ListAliases")
    return
}

func NewListAliasesResponse() (response *ListAliasesResponse) {
    response = &ListAliasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListAliases
// This API is used to return the list of all aliases under a function. You can filter them by the specific function version.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETERVALUE_FUNCTIONNAME = "InvalidParameterValue.FunctionName"
//  INVALIDPARAMETERVALUE_NAMESPACE = "InvalidParameterValue.Namespace"
//  RESOURCENOTFOUND_FUNCTION = "ResourceNotFound.Function"
//  RESOURCENOTFOUND_FUNCTIONNAME = "ResourceNotFound.FunctionName"
//  RESOURCENOTFOUND_NAMESPACE = "ResourceNotFound.Namespace"
//  UNAUTHORIZEDOPERATION_CAM = "UnauthorizedOperation.CAM"
func (c *Client) ListAliases(request *ListAliasesRequest) (response *ListAliasesResponse, err error) {
    if request == nil {
        request = NewListAliasesRequest()
    }
    response = NewListAliasesResponse()
    err = c.Send(request, response)
    return
}

func NewListAsyncEventsRequest() (request *ListAsyncEventsRequest) {
    request = &ListAsyncEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "ListAsyncEvents")
    return
}

func NewListAsyncEventsResponse() (response *ListAsyncEventsResponse) {
    response = &ListAsyncEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListAsyncEvents
// This API is used to pull the list of async function events.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVOKETYPE = "InvalidParameterValue.InvokeType"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_ORDER = "InvalidParameterValue.Order"
//  INVALIDPARAMETERVALUE_STATUS = "InvalidParameterValue.Status"
//  RESOURCENOTFOUND_FUNCTION = "ResourceNotFound.Function"
//  RESOURCENOTFOUND_NAMESPACE = "ResourceNotFound.Namespace"
//  RESOURCENOTFOUND_VERSION = "ResourceNotFound.Version"
func (c *Client) ListAsyncEvents(request *ListAsyncEventsRequest) (response *ListAsyncEventsResponse, err error) {
    if request == nil {
        request = NewListAsyncEventsRequest()
    }
    response = NewListAsyncEventsResponse()
    err = c.Send(request, response)
    return
}

func NewListFunctionsRequest() (request *ListFunctionsRequest) {
    request = &ListFunctionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "ListFunctions")
    return
}

func NewListFunctionsResponse() (response *ListFunctionsResponse) {
    response = &ListFunctionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListFunctions
// This API is used to return relevant function information based on the input query parameters.
//
// error code that may be returned:
//  INTERNALERROR_EXCEPTION = "InternalError.Exception"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_PAYLOAD = "InvalidParameter.Payload"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_NAMESPACE = "InvalidParameterValue.Namespace"
//  INVALIDPARAMETERVALUE_OFFSET = "InvalidParameterValue.Offset"
//  INVALIDPARAMETERVALUE_ORDER = "InvalidParameterValue.Order"
//  RESOURCENOTFOUND_NAMESPACE = "ResourceNotFound.Namespace"
//  UNAUTHORIZEDOPERATION_CAM = "UnauthorizedOperation.CAM"
//  UNAUTHORIZEDOPERATION_REGION = "UnauthorizedOperation.Region"
func (c *Client) ListFunctions(request *ListFunctionsRequest) (response *ListFunctionsResponse, err error) {
    if request == nil {
        request = NewListFunctionsRequest()
    }
    response = NewListFunctionsResponse()
    err = c.Send(request, response)
    return
}

func NewListLayerVersionsRequest() (request *ListLayerVersionsRequest) {
    request = &ListLayerVersionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "ListLayerVersions")
    return
}

func NewListLayerVersionsResponse() (response *ListLayerVersionsResponse) {
    response = &ListLayerVersionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListLayerVersions
// This API is used to get the information of all versions of a specified layer.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  RESOURCENOTFOUND_LAYER = "ResourceNotFound.Layer"
//  UNAUTHORIZEDOPERATION_CAM = "UnauthorizedOperation.CAM"
func (c *Client) ListLayerVersions(request *ListLayerVersionsRequest) (response *ListLayerVersionsResponse, err error) {
    if request == nil {
        request = NewListLayerVersionsRequest()
    }
    response = NewListLayerVersionsResponse()
    err = c.Send(request, response)
    return
}

func NewListLayersRequest() (request *ListLayersRequest) {
    request = &ListLayersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "ListLayers")
    return
}

func NewListLayersResponse() (response *ListLayersResponse) {
    response = &ListLayersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListLayers
// This API is used to return the list of all layers, including the information of the latest version of each layer. You can filter them by the compatible runtime.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETERVALUE_RUNTIME = "InvalidParameterValue.Runtime"
//  INVALIDPARAMETERVALUE_STAMP = "InvalidParameterValue.Stamp"
//  LIMITEXCEEDED_LAYERS = "LimitExceeded.Layers"
//  UNAUTHORIZEDOPERATION_CAM = "UnauthorizedOperation.CAM"
func (c *Client) ListLayers(request *ListLayersRequest) (response *ListLayersResponse, err error) {
    if request == nil {
        request = NewListLayersRequest()
    }
    response = NewListLayersResponse()
    err = c.Send(request, response)
    return
}

func NewListNamespacesRequest() (request *ListNamespacesRequest) {
    request = &ListNamespacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "ListNamespaces")
    return
}

func NewListNamespacesResponse() (response *ListNamespacesResponse) {
    response = &ListNamespacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListNamespaces
// This API is used to display a namespace list.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_FILTERS = "InvalidParameterValue.Filters"
//  INVALIDPARAMETERVALUE_ORDER = "InvalidParameterValue.Order"
//  INVALIDPARAMETERVALUE_SEARCHKEY = "InvalidParameterValue.SearchKey"
//  INVALIDPARAMETERVALUE_STAMP = "InvalidParameterValue.Stamp"
func (c *Client) ListNamespaces(request *ListNamespacesRequest) (response *ListNamespacesResponse, err error) {
    if request == nil {
        request = NewListNamespacesRequest()
    }
    response = NewListNamespacesResponse()
    err = c.Send(request, response)
    return
}

func NewListTriggersRequest() (request *ListTriggersRequest) {
    request = &ListTriggersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "ListTriggers")
    return
}

func NewListTriggersResponse() (response *ListTriggersResponse) {
    response = &ListTriggersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListTriggers
// This API is used to get the function trigger list.
//
// error code that may be returned:
//  FAILEDOPERATION_APIGW = "FailedOperation.Apigw"
//  INVALIDPARAMETERVALUE_FILTERS = "InvalidParameterValue.Filters"
//  INVALIDPARAMETERVALUE_ORDER = "InvalidParameterValue.Order"
func (c *Client) ListTriggers(request *ListTriggersRequest) (response *ListTriggersResponse, err error) {
    if request == nil {
        request = NewListTriggersRequest()
    }
    response = NewListTriggersResponse()
    err = c.Send(request, response)
    return
}

func NewListVersionByFunctionRequest() (request *ListVersionByFunctionRequest) {
    request = &ListVersionByFunctionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "ListVersionByFunction")
    return
}

func NewListVersionByFunctionResponse() (response *ListVersionByFunctionResponse) {
    response = &ListVersionByFunctionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListVersionByFunction
// This API is used to query the function version based on the input parameters.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FUNCTIONNAME = "InvalidParameterValue.FunctionName"
//  INVALIDPARAMETERVALUE_NAMESPACE = "InvalidParameterValue.Namespace"
//  INVALIDPARAMETERVALUE_ORDER = "InvalidParameterValue.Order"
//  RESOURCENOTFOUND_FUNCTION = "ResourceNotFound.Function"
//  RESOURCENOTFOUND_FUNCTIONNAME = "ResourceNotFound.FunctionName"
//  RESOURCENOTFOUND_NAMESPACE = "ResourceNotFound.Namespace"
//  UNAUTHORIZEDOPERATION_CAM = "UnauthorizedOperation.CAM"
func (c *Client) ListVersionByFunction(request *ListVersionByFunctionRequest) (response *ListVersionByFunctionResponse, err error) {
    if request == nil {
        request = NewListVersionByFunctionRequest()
    }
    response = NewListVersionByFunctionResponse()
    err = c.Send(request, response)
    return
}

func NewPublishLayerVersionRequest() (request *PublishLayerVersionRequest) {
    request = &PublishLayerVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "PublishLayerVersion")
    return
}

func NewPublishLayerVersionResponse() (response *PublishLayerVersionResponse) {
    response = &PublishLayerVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PublishLayerVersion
// This API is used to create a version for a layer by using the given .zip file or COS object. Each time this API is called with the same layer name, a new version will be generated.
//
// error code that may be returned:
//  FAILEDOPERATION_PUBLISHLAYERVERSION = "FailedOperation.PublishLayerVersion"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXCEPTION = "InternalError.Exception"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETERVALUE_COMPATIBLERUNTIMES = "InvalidParameterValue.CompatibleRuntimes"
//  INVALIDPARAMETERVALUE_CONTENT = "InvalidParameterValue.Content"
//  INVALIDPARAMETERVALUE_COS = "InvalidParameterValue.Cos"
//  INVALIDPARAMETERVALUE_COSBUCKETNAME = "InvalidParameterValue.CosBucketName"
//  INVALIDPARAMETERVALUE_COSBUCKETREGION = "InvalidParameterValue.CosBucketRegion"
//  INVALIDPARAMETERVALUE_COSOBJECTNAME = "InvalidParameterValue.CosObjectName"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_LAYERNAME = "InvalidParameterValue.LayerName"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_RUNTIME = "InvalidParameterValue.Runtime"
//  INVALIDPARAMETERVALUE_STAMP = "InvalidParameterValue.Stamp"
//  INVALIDPARAMETERVALUE_TEMPCOSOBJECTNAME = "InvalidParameterValue.TempCosObjectName"
//  LIMITEXCEEDED_LAYERVERSIONS = "LimitExceeded.LayerVersions"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAM = "UnauthorizedOperation.CAM"
//  UNAUTHORIZEDOPERATION_NOTMC = "UnauthorizedOperation.NotMC"
//  UNSUPPORTEDOPERATION_COS = "UnsupportedOperation.Cos"
func (c *Client) PublishLayerVersion(request *PublishLayerVersionRequest) (response *PublishLayerVersionResponse, err error) {
    if request == nil {
        request = NewPublishLayerVersionRequest()
    }
    response = NewPublishLayerVersionResponse()
    err = c.Send(request, response)
    return
}

func NewPublishVersionRequest() (request *PublishVersionRequest) {
    request = &PublishVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "PublishVersion")
    return
}

func NewPublishVersionResponse() (response *PublishVersionResponse) {
    response = &PublishVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PublishVersion
// This API is used for users to release a new version of the function.
//
// error code that may be returned:
//  FAILEDOPERATION_PUBLISHVERSION = "FailedOperation.PublishVersion"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_FUNCTIONNAME = "InvalidParameterValue.FunctionName"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_FUNCTION = "ResourceNotFound.Function"
//  RESOURCENOTFOUND_FUNCTIONNAME = "ResourceNotFound.FunctionName"
//  RESOURCENOTFOUND_NAMESPACE = "ResourceNotFound.Namespace"
//  UNAUTHORIZEDOPERATION_CAM = "UnauthorizedOperation.CAM"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) PublishVersion(request *PublishVersionRequest) (response *PublishVersionResponse, err error) {
    if request == nil {
        request = NewPublishVersionRequest()
    }
    response = NewPublishVersionResponse()
    err = c.Send(request, response)
    return
}

func NewPutProvisionedConcurrencyConfigRequest() (request *PutProvisionedConcurrencyConfigRequest) {
    request = &PutProvisionedConcurrencyConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "PutProvisionedConcurrencyConfig")
    return
}

func NewPutProvisionedConcurrencyConfigResponse() (response *PutProvisionedConcurrencyConfigResponse) {
    response = &PutProvisionedConcurrencyConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PutProvisionedConcurrencyConfig
// This API is used to set the provisioned concurrency of a non-$LATEST version of a function.
//
// error code that may be returned:
//  FAILEDOPERATION_DEBUGMODESTATUS = "FailedOperation.DebugModeStatus"
//  FAILEDOPERATION_FUNCTIONVERSIONSTATUSNOTACTIVE = "FailedOperation.FunctionVersionStatusNotActive"
//  FAILEDOPERATION_INSUFFICIENTBALANCE = "FailedOperation.InsufficientBalance"
//  FAILEDOPERATION_PROVISIONEDINPROGRESS = "FailedOperation.ProvisionedInProgress"
//  FAILEDOPERATION_UNOPENEDSERVICE = "FailedOperation.UnOpenedService"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FUNCTIONNAME = "InvalidParameterValue.FunctionName"
//  INVALIDPARAMETERVALUE_QUALIFIER = "InvalidParameterValue.Qualifier"
//  LIMITEXCEEDED_FUNCTIONPROVISIONEDCONCURRENCYMEMORY = "LimitExceeded.FunctionProvisionedConcurrencyMemory"
//  LIMITEXCEEDED_FUNCTIONTOTALPROVISIONEDCONCURRENCYMEMORY = "LimitExceeded.FunctionTotalProvisionedConcurrencyMemory"
//  LIMITEXCEEDED_FUNCTIONTOTALPROVISIONEDCONCURRENCYNUM = "LimitExceeded.FunctionTotalProvisionedConcurrencyNum"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FUNCTION = "ResourceNotFound.Function"
//  RESOURCENOTFOUND_FUNCTIONNAME = "ResourceNotFound.FunctionName"
//  RESOURCENOTFOUND_FUNCTIONVERSION = "ResourceNotFound.FunctionVersion"
//  RESOURCENOTFOUND_NAMESPACE = "ResourceNotFound.Namespace"
//  RESOURCENOTFOUND_VERSION = "ResourceNotFound.Version"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) PutProvisionedConcurrencyConfig(request *PutProvisionedConcurrencyConfigRequest) (response *PutProvisionedConcurrencyConfigResponse, err error) {
    if request == nil {
        request = NewPutProvisionedConcurrencyConfigRequest()
    }
    response = NewPutProvisionedConcurrencyConfigResponse()
    err = c.Send(request, response)
    return
}

func NewPutReservedConcurrencyConfigRequest() (request *PutReservedConcurrencyConfigRequest) {
    request = &PutReservedConcurrencyConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "PutReservedConcurrencyConfig")
    return
}

func NewPutReservedConcurrencyConfigResponse() (response *PutReservedConcurrencyConfigResponse) {
    response = &PutReservedConcurrencyConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PutReservedConcurrencyConfig
// This API is used to set the reserved concurrency of a function.
//
// error code that may be returned:
//  FAILEDOPERATION_DEBUGMODESTATUS = "FailedOperation.DebugModeStatus"
//  FAILEDOPERATION_FUNCTIONVERSIONSTATUSNOTACTIVE = "FailedOperation.FunctionVersionStatusNotActive"
//  FAILEDOPERATION_INSUFFICIENTBALANCE = "FailedOperation.InsufficientBalance"
//  FAILEDOPERATION_RESERVEDINPROGRESS = "FailedOperation.ReservedInProgress"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_FUNCTIONRESERVEDCONCURRENCYMEMORY = "LimitExceeded.FunctionReservedConcurrencyMemory"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FUNCTION = "ResourceNotFound.Function"
//  RESOURCENOTFOUND_FUNCTIONNAME = "ResourceNotFound.FunctionName"
//  RESOURCENOTFOUND_NAMESPACE = "ResourceNotFound.Namespace"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) PutReservedConcurrencyConfig(request *PutReservedConcurrencyConfigRequest) (response *PutReservedConcurrencyConfigResponse, err error) {
    if request == nil {
        request = NewPutReservedConcurrencyConfigRequest()
    }
    response = NewPutReservedConcurrencyConfigResponse()
    err = c.Send(request, response)
    return
}

func NewPutTotalConcurrencyConfigRequest() (request *PutTotalConcurrencyConfigRequest) {
    request = &PutTotalConcurrencyConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "PutTotalConcurrencyConfig")
    return
}

func NewPutTotalConcurrencyConfigResponse() (response *PutTotalConcurrencyConfigResponse) {
    response = &PutTotalConcurrencyConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PutTotalConcurrencyConfig
// This API is used to modify the account concurrency quota.
//
// error code that may be returned:
//  FAILEDOPERATION_INSUFFICIENTBALANCE = "FailedOperation.InsufficientBalance"
//  FAILEDOPERATION_TOTALCONCURRENCYMEMORYINPROGRESS = "FailedOperation.TotalConcurrencyMemoryInProgress"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_TOTALCONCURRENCYMEMORY = "LimitExceeded.TotalConcurrencyMemory"
//  LIMITEXCEEDED_USERTOTALCONCURRENCYMEMORY = "LimitExceeded.UserTotalConcurrencyMemory"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_NAMESPACE = "ResourceNotFound.Namespace"
//  RESOURCENOTFOUND_TOTALCONCURRENCYMEMORY = "ResourceNotFound.TotalConcurrencyMemory"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) PutTotalConcurrencyConfig(request *PutTotalConcurrencyConfigRequest) (response *PutTotalConcurrencyConfigResponse, err error) {
    if request == nil {
        request = NewPutTotalConcurrencyConfigRequest()
    }
    response = NewPutTotalConcurrencyConfigResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateAsyncEventRequest() (request *TerminateAsyncEventRequest) {
    request = &TerminateAsyncEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "TerminateAsyncEvent")
    return
}

func NewTerminateAsyncEventResponse() (response *TerminateAsyncEventResponse) {
    response = &TerminateAsyncEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TerminateAsyncEvent
// This API is used to terminate a running async function event.
//
// error code that may be returned:
//  FAILEDOPERATION_ASYNCEVENTSTATUS = "FailedOperation.AsyncEventStatus"
//  RESOURCENOTFOUND_ASYNCEVENT = "ResourceNotFound.AsyncEvent"
//  RESOURCENOTFOUND_FUNCTION = "ResourceNotFound.Function"
//  RESOURCENOTFOUND_NAMESPACE = "ResourceNotFound.Namespace"
func (c *Client) TerminateAsyncEvent(request *TerminateAsyncEventRequest) (response *TerminateAsyncEventResponse, err error) {
    if request == nil {
        request = NewTerminateAsyncEventRequest()
    }
    response = NewTerminateAsyncEventResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAliasRequest() (request *UpdateAliasRequest) {
    request = &UpdateAliasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "UpdateAlias")
    return
}

func NewUpdateAliasResponse() (response *UpdateAliasResponse) {
    response = &UpdateAliasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateAlias
// This API is used to update the configuration of an alias.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UPDATEALIAS = "FailedOperation.UpdateAlias"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_ROUTINGCONFIG = "InvalidParameter.RoutingConfig"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ADDITIONALVERSIONWEIGHTS = "InvalidParameterValue.AdditionalVersionWeights"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"
//  INVALIDPARAMETERVALUE_ROUTINGCONFIG = "InvalidParameterValue.RoutingConfig"
//  RESOURCENOTFOUND_ALIAS = "ResourceNotFound.Alias"
//  RESOURCENOTFOUND_FUNCTION = "ResourceNotFound.Function"
//  RESOURCENOTFOUND_FUNCTIONNAME = "ResourceNotFound.FunctionName"
//  RESOURCENOTFOUND_FUNCTIONVERSION = "ResourceNotFound.FunctionVersion"
//  RESOURCENOTFOUND_NAMESPACE = "ResourceNotFound.Namespace"
//  UNAUTHORIZEDOPERATION_CAM = "UnauthorizedOperation.CAM"
func (c *Client) UpdateAlias(request *UpdateAliasRequest) (response *UpdateAliasResponse, err error) {
    if request == nil {
        request = NewUpdateAliasRequest()
    }
    response = NewUpdateAliasResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateFunctionCodeRequest() (request *UpdateFunctionCodeRequest) {
    request = &UpdateFunctionCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "UpdateFunctionCode")
    return
}

func NewUpdateFunctionCodeResponse() (response *UpdateFunctionCodeResponse) {
    response = &UpdateFunctionCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateFunctionCode
// This API is used to update the function code based on the input parameters.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FUNCTIONSTATUSERROR = "FailedOperation.FunctionStatusError"
//  FAILEDOPERATION_OPERATIONCONFLICT = "FailedOperation.OperationConflict"
//  FAILEDOPERATION_UPDATEFUNCTIONCODE = "FailedOperation.UpdateFunctionCode"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_PAYLOAD = "InvalidParameter.Payload"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CODE = "InvalidParameterValue.Code"
//  INVALIDPARAMETERVALUE_CODESECRET = "InvalidParameterValue.CodeSecret"
//  INVALIDPARAMETERVALUE_CODESOURCE = "InvalidParameterValue.CodeSource"
//  INVALIDPARAMETERVALUE_COS = "InvalidParameterValue.Cos"
//  INVALIDPARAMETERVALUE_COSBUCKETNAME = "InvalidParameterValue.CosBucketName"
//  INVALIDPARAMETERVALUE_COSBUCKETREGION = "InvalidParameterValue.CosBucketRegion"
//  INVALIDPARAMETERVALUE_COSOBJECTNAME = "InvalidParameterValue.CosObjectName"
//  INVALIDPARAMETERVALUE_FUNCTIONNAME = "InvalidParameterValue.FunctionName"
//  INVALIDPARAMETERVALUE_GITBRANCH = "InvalidParameterValue.GitBranch"
//  INVALIDPARAMETERVALUE_GITDIRECTORY = "InvalidParameterValue.GitDirectory"
//  INVALIDPARAMETERVALUE_GITPASSWORD = "InvalidParameterValue.GitPassword"
//  INVALIDPARAMETERVALUE_GITURL = "InvalidParameterValue.GitUrl"
//  INVALIDPARAMETERVALUE_GITUSERNAME = "InvalidParameterValue.GitUserName"
//  INVALIDPARAMETERVALUE_HANDLER = "InvalidParameterValue.Handler"
//  INVALIDPARAMETERVALUE_INLINEZIPFILE = "InvalidParameterValue.InlineZipFile"
//  INVALIDPARAMETERVALUE_NAMESPACE = "InvalidParameterValue.Namespace"
//  INVALIDPARAMETERVALUE_TEMPCOSOBJECTNAME = "InvalidParameterValue.TempCosObjectName"
//  INVALIDPARAMETERVALUE_ZIPFILE = "InvalidParameterValue.ZipFile"
//  INVALIDPARAMETERVALUE_ZIPFILEBASE64BINASCIIERROR = "InvalidParameterValue.ZipFileBase64BinasciiError"
//  RESOURCENOTFOUND_FUNCTION = "ResourceNotFound.Function"
//  RESOURCENOTFOUND_FUNCTIONNAME = "ResourceNotFound.FunctionName"
//  RESOURCENOTFOUND_NAMESPACE = "ResourceNotFound.Namespace"
//  UNAUTHORIZEDOPERATION_CAM = "UnauthorizedOperation.CAM"
//  UNAUTHORIZEDOPERATION_TEMPCOSAPPID = "UnauthorizedOperation.TempCosAppid"
//  UNAUTHORIZEDOPERATION_UPDATEFUNCTIONCODE = "UnauthorizedOperation.UpdateFunctionCode"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateFunctionCode(request *UpdateFunctionCodeRequest) (response *UpdateFunctionCodeResponse, err error) {
    if request == nil {
        request = NewUpdateFunctionCodeRequest()
    }
    response = NewUpdateFunctionCodeResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateFunctionConfigurationRequest() (request *UpdateFunctionConfigurationRequest) {
    request = &UpdateFunctionConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "UpdateFunctionConfiguration")
    return
}

func NewUpdateFunctionConfigurationResponse() (response *UpdateFunctionConfigurationResponse) {
    response = &UpdateFunctionConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateFunctionConfiguration
// This API is used to update the function configuration based on the input parameters.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DEBUGMODEUPDATETIMEOUTFAIL = "FailedOperation.DebugModeUpdateTimeOutFail"
//  FAILEDOPERATION_QCSROLENOTFOUND = "FailedOperation.QcsRoleNotFound"
//  FAILEDOPERATION_RESERVEDINPROGRESS = "FailedOperation.ReservedInProgress"
//  FAILEDOPERATION_UPDATEFUNCTIONCONFIGURATION = "FailedOperation.UpdateFunctionConfiguration"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_PAYLOAD = "InvalidParameter.Payload"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CFSPARAMETERDUPLICATE = "InvalidParameterValue.CfsParameterDuplicate"
//  INVALIDPARAMETERVALUE_CLS = "InvalidParameterValue.Cls"
//  INVALIDPARAMETERVALUE_CLSROLE = "InvalidParameterValue.ClsRole"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_EIPCONFIG = "InvalidParameterValue.EipConfig"
//  INVALIDPARAMETERVALUE_ENVIRONMENT = "InvalidParameterValue.Environment"
//  INVALIDPARAMETERVALUE_ENVIRONMENTEXCEEDEDLIMIT = "InvalidParameterValue.EnvironmentExceededLimit"
//  INVALIDPARAMETERVALUE_ENVIRONMENTSYSTEMPROTECT = "InvalidParameterValue.EnvironmentSystemProtect"
//  INVALIDPARAMETERVALUE_FUNCTIONNAME = "InvalidParameterValue.FunctionName"
//  INVALIDPARAMETERVALUE_HANDLER = "InvalidParameterValue.Handler"
//  INVALIDPARAMETERVALUE_LAYERS = "InvalidParameterValue.Layers"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MEMORY = "InvalidParameterValue.Memory"
//  INVALIDPARAMETERVALUE_MEMORYSIZE = "InvalidParameterValue.MemorySize"
//  INVALIDPARAMETERVALUE_NAMESPACE = "InvalidParameterValue.Namespace"
//  INVALIDPARAMETERVALUE_PUBLICNETCONFIG = "InvalidParameterValue.PublicNetConfig"
//  INVALIDPARAMETERVALUE_RUNTIME = "InvalidParameterValue.Runtime"
//  INVALIDPARAMETERVALUE_SYSTEMENVIRONMENT = "InvalidParameterValue.SystemEnvironment"
//  LIMITEXCEEDED_EIP = "LimitExceeded.Eip"
//  LIMITEXCEEDED_INITTIMEOUT = "LimitExceeded.InitTimeout"
//  LIMITEXCEEDED_MEMORY = "LimitExceeded.Memory"
//  LIMITEXCEEDED_TIMEOUT = "LimitExceeded.Timeout"
//  RESOURCENOTFOUND_CFSVPCNOTMATCH = "ResourceNotFound.CfsVpcNotMatch"
//  RESOURCENOTFOUND_CMQ = "ResourceNotFound.Cmq"
//  RESOURCENOTFOUND_FUNCTION = "ResourceNotFound.Function"
//  RESOURCENOTFOUND_FUNCTIONNAME = "ResourceNotFound.FunctionName"
//  RESOURCENOTFOUND_GETCFSNOTMATCH = "ResourceNotFound.GetCfsNotMatch"
//  RESOURCENOTFOUND_LAYER = "ResourceNotFound.Layer"
//  RESOURCENOTFOUND_NAMESPACE = "ResourceNotFound.Namespace"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
//  RESOURCENOTFOUND_VPC = "ResourceNotFound.Vpc"
//  UNAUTHORIZEDOPERATION_CAM = "UnauthorizedOperation.CAM"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateFunctionConfiguration(request *UpdateFunctionConfigurationRequest) (response *UpdateFunctionConfigurationResponse, err error) {
    if request == nil {
        request = NewUpdateFunctionConfigurationRequest()
    }
    response = NewUpdateFunctionConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateFunctionEventInvokeConfigRequest() (request *UpdateFunctionEventInvokeConfigRequest) {
    request = &UpdateFunctionEventInvokeConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "UpdateFunctionEventInvokeConfig")
    return
}

func NewUpdateFunctionEventInvokeConfigResponse() (response *UpdateFunctionEventInvokeConfigResponse) {
    response = &UpdateFunctionEventInvokeConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateFunctionEventInvokeConfig
// This API is used to update the async retry configuration of a function, including the number of retry attempts and message retention period.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FUNCTIONVERSIONSTATUSNOTACTIVE = "FailedOperation.FunctionVersionStatusNotActive"
//  FAILEDOPERATION_INSUFFICIENTBALANCE = "FailedOperation.InsufficientBalance"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ASYNCTRIGGERCONFIG = "InvalidParameterValue.AsyncTriggerConfig"
//  INVALIDPARAMETERVALUE_FUNCTIONNAME = "InvalidParameterValue.FunctionName"
//  INVALIDPARAMETERVALUE_NAMESPACE = "InvalidParameterValue.Namespace"
//  LIMITEXCEEDED_MSGTTL = "LimitExceeded.MsgTTL"
//  LIMITEXCEEDED_RETRYNUM = "LimitExceeded.RetryNum"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FUNCTION = "ResourceNotFound.Function"
//  RESOURCENOTFOUND_FUNCTIONNAME = "ResourceNotFound.FunctionName"
//  RESOURCENOTFOUND_FUNCTIONVERSION = "ResourceNotFound.FunctionVersion"
//  RESOURCENOTFOUND_NAMESPACE = "ResourceNotFound.Namespace"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) UpdateFunctionEventInvokeConfig(request *UpdateFunctionEventInvokeConfigRequest) (response *UpdateFunctionEventInvokeConfigResponse, err error) {
    if request == nil {
        request = NewUpdateFunctionEventInvokeConfigRequest()
    }
    response = NewUpdateFunctionEventInvokeConfigResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateNamespaceRequest() (request *UpdateNamespaceRequest) {
    request = &UpdateNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "UpdateNamespace")
    return
}

func NewUpdateNamespaceResponse() (response *UpdateNamespaceResponse) {
    response = &UpdateNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateNamespace
// This API is used to update a namespace.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  RESOURCENOTFOUND_NAMESPACE = "ResourceNotFound.Namespace"
func (c *Client) UpdateNamespace(request *UpdateNamespaceRequest) (response *UpdateNamespaceResponse, err error) {
    if request == nil {
        request = NewUpdateNamespaceRequest()
    }
    response = NewUpdateNamespaceResponse()
    err = c.Send(request, response)
    return
}
