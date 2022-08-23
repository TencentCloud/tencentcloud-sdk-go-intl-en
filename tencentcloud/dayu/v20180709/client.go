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

package v20180709

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2018-07-09"

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


func NewCreateBasicDDoSAlarmThresholdRequest() (request *CreateBasicDDoSAlarmThresholdRequest) {
    request = &CreateBasicDDoSAlarmThresholdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "CreateBasicDDoSAlarmThreshold")
    
    
    return
}

func NewCreateBasicDDoSAlarmThresholdResponse() (response *CreateBasicDDoSAlarmThresholdResponse) {
    response = &CreateBasicDDoSAlarmThresholdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBasicDDoSAlarmThreshold
// This API is used to set the DDoS alarm threshold for Anti-DDoS Basic, which is only supported for Anti-DDoS Basic.
func (c *Client) CreateBasicDDoSAlarmThreshold(request *CreateBasicDDoSAlarmThresholdRequest) (response *CreateBasicDDoSAlarmThresholdResponse, err error) {
    return c.CreateBasicDDoSAlarmThresholdWithContext(context.Background(), request)
}

// CreateBasicDDoSAlarmThreshold
// This API is used to set the DDoS alarm threshold for Anti-DDoS Basic, which is only supported for Anti-DDoS Basic.
func (c *Client) CreateBasicDDoSAlarmThresholdWithContext(ctx context.Context, request *CreateBasicDDoSAlarmThresholdRequest) (response *CreateBasicDDoSAlarmThresholdResponse, err error) {
    if request == nil {
        request = NewCreateBasicDDoSAlarmThresholdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBasicDDoSAlarmThreshold require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBasicDDoSAlarmThresholdResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBoundIPRequest() (request *CreateBoundIPRequest) {
    request = &CreateBoundIPRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "CreateBoundIP")
    
    
    return
}

func NewCreateBoundIPResponse() (response *CreateBoundIPResponse) {
    response = &CreateBoundIPResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBoundIP
// This API is used to bind an IP to an Anti-DDoS Pro instance, which supports both single IP instances and multi-IP instances. It should be noted that this API is async; therefore, if a binding/unbinding operation is in progress, new binding/unbinding operations cannot be initiated.
//
// error code that may be returned:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateBoundIP(request *CreateBoundIPRequest) (response *CreateBoundIPResponse, err error) {
    return c.CreateBoundIPWithContext(context.Background(), request)
}

// CreateBoundIP
// This API is used to bind an IP to an Anti-DDoS Pro instance, which supports both single IP instances and multi-IP instances. It should be noted that this API is async; therefore, if a binding/unbinding operation is in progress, new binding/unbinding operations cannot be initiated.
//
// error code that may be returned:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateBoundIPWithContext(ctx context.Context, request *CreateBoundIPRequest) (response *CreateBoundIPResponse, err error) {
    if request == nil {
        request = NewCreateBoundIPRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBoundIP require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBoundIPResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCCFrequencyRulesRequest() (request *CreateCCFrequencyRulesRequest) {
    request = &CreateCCFrequencyRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "CreateCCFrequencyRules")
    
    
    return
}

func NewCreateCCFrequencyRulesResponse() (response *CreateCCFrequencyRulesResponse) {
    response = &CreateCCFrequencyRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCCFrequencyRules
// This API is used to add an access frequency control rule for CC protection.
//
// error code that may be returned:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCCFrequencyRules(request *CreateCCFrequencyRulesRequest) (response *CreateCCFrequencyRulesResponse, err error) {
    return c.CreateCCFrequencyRulesWithContext(context.Background(), request)
}

// CreateCCFrequencyRules
// This API is used to add an access frequency control rule for CC protection.
//
// error code that may be returned:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCCFrequencyRulesWithContext(ctx context.Context, request *CreateCCFrequencyRulesRequest) (response *CreateCCFrequencyRulesResponse, err error) {
    if request == nil {
        request = NewCreateCCFrequencyRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCCFrequencyRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCCFrequencyRulesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCCSelfDefinePolicyRequest() (request *CreateCCSelfDefinePolicyRequest) {
    request = &CreateCCSelfDefinePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "CreateCCSelfDefinePolicy")
    
    
    return
}

func NewCreateCCSelfDefinePolicyResponse() (response *CreateCCSelfDefinePolicyResponse) {
    response = &CreateCCSelfDefinePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCCSelfDefinePolicy
// This API is used to create a custom CC policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateCCSelfDefinePolicy(request *CreateCCSelfDefinePolicyRequest) (response *CreateCCSelfDefinePolicyResponse, err error) {
    return c.CreateCCSelfDefinePolicyWithContext(context.Background(), request)
}

// CreateCCSelfDefinePolicy
// This API is used to create a custom CC policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateCCSelfDefinePolicyWithContext(ctx context.Context, request *CreateCCSelfDefinePolicyRequest) (response *CreateCCSelfDefinePolicyResponse, err error) {
    if request == nil {
        request = NewCreateCCSelfDefinePolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCCSelfDefinePolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCCSelfDefinePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDDoSPolicyRequest() (request *CreateDDoSPolicyRequest) {
    request = &CreateDDoSPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "CreateDDoSPolicy")
    
    
    return
}

func NewCreateDDoSPolicyResponse() (response *CreateDDoSPolicyResponse) {
    response = &CreateDDoSPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDDoSPolicy
// This API is used to add an advanced DDoS policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateDDoSPolicy(request *CreateDDoSPolicyRequest) (response *CreateDDoSPolicyResponse, err error) {
    return c.CreateDDoSPolicyWithContext(context.Background(), request)
}

// CreateDDoSPolicy
// This API is used to add an advanced DDoS policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateDDoSPolicyWithContext(ctx context.Context, request *CreateDDoSPolicyRequest) (response *CreateDDoSPolicyResponse, err error) {
    if request == nil {
        request = NewCreateDDoSPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDDoSPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDDoSPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDDoSPolicyCaseRequest() (request *CreateDDoSPolicyCaseRequest) {
    request = &CreateDDoSPolicyCaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "CreateDDoSPolicyCase")
    
    
    return
}

func NewCreateDDoSPolicyCaseResponse() (response *CreateDDoSPolicyCaseResponse) {
    response = &CreateDDoSPolicyCaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDDoSPolicyCase
// This API is used to add a policy scenario.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateDDoSPolicyCase(request *CreateDDoSPolicyCaseRequest) (response *CreateDDoSPolicyCaseResponse, err error) {
    return c.CreateDDoSPolicyCaseWithContext(context.Background(), request)
}

// CreateDDoSPolicyCase
// This API is used to add a policy scenario.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateDDoSPolicyCaseWithContext(ctx context.Context, request *CreateDDoSPolicyCaseRequest) (response *CreateDDoSPolicyCaseResponse, err error) {
    if request == nil {
        request = NewCreateDDoSPolicyCaseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDDoSPolicyCase require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDDoSPolicyCaseResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstanceNameRequest() (request *CreateInstanceNameRequest) {
    request = &CreateInstanceNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "CreateInstanceName")
    
    
    return
}

func NewCreateInstanceNameResponse() (response *CreateInstanceNameResponse) {
    response = &CreateInstanceNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateInstanceName
// This API is used to rename a resource instance, which supports single IP instances, multi-IP instances, Anti-DDoS Advanced, and Anti-DDoS Ultimate.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateInstanceName(request *CreateInstanceNameRequest) (response *CreateInstanceNameResponse, err error) {
    return c.CreateInstanceNameWithContext(context.Background(), request)
}

// CreateInstanceName
// This API is used to rename a resource instance, which supports single IP instances, multi-IP instances, Anti-DDoS Advanced, and Anti-DDoS Ultimate.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateInstanceNameWithContext(ctx context.Context, request *CreateInstanceNameRequest) (response *CreateInstanceNameResponse, err error) {
    if request == nil {
        request = NewCreateInstanceNameRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInstanceName require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInstanceNameResponse()
    err = c.Send(request, response)
    return
}

func NewCreateL4HealthConfigRequest() (request *CreateL4HealthConfigRequest) {
    request = &CreateL4HealthConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "CreateL4HealthConfig")
    
    
    return
}

func NewCreateL4HealthConfigResponse() (response *CreateL4HealthConfigResponse) {
    response = &CreateL4HealthConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateL4HealthConfig
// This API is used to upload layer-4 health check configuration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateL4HealthConfig(request *CreateL4HealthConfigRequest) (response *CreateL4HealthConfigResponse, err error) {
    return c.CreateL4HealthConfigWithContext(context.Background(), request)
}

// CreateL4HealthConfig
// This API is used to upload layer-4 health check configuration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateL4HealthConfigWithContext(ctx context.Context, request *CreateL4HealthConfigRequest) (response *CreateL4HealthConfigResponse, err error) {
    if request == nil {
        request = NewCreateL4HealthConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateL4HealthConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateL4HealthConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateL4RulesRequest() (request *CreateL4RulesRequest) {
    request = &CreateL4RulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "CreateL4Rules")
    
    
    return
}

func NewCreateL4RulesResponse() (response *CreateL4RulesResponse) {
    response = &CreateL4RulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateL4Rules
// This API is used to add a layer-4 forwarding rule.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateL4Rules(request *CreateL4RulesRequest) (response *CreateL4RulesResponse, err error) {
    return c.CreateL4RulesWithContext(context.Background(), request)
}

// CreateL4Rules
// This API is used to add a layer-4 forwarding rule.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateL4RulesWithContext(ctx context.Context, request *CreateL4RulesRequest) (response *CreateL4RulesResponse, err error) {
    if request == nil {
        request = NewCreateL4RulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateL4Rules require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateL4RulesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateL7CCRuleRequest() (request *CreateL7CCRuleRequest) {
    request = &CreateL7CCRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "CreateL7CCRule")
    
    
    return
}

func NewCreateL7CCRuleResponse() (response *CreateL7CCRuleResponse) {
    response = &CreateL7CCRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateL7CCRule
// This API is used to add a custom frequency control rule for layer-7 CC access (it works in the IP + Host dimension and does not support specific URIs). As it has been disused, please call the new `CreateCCFrequencyRules` API, which supports both IP + Host and URI.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateL7CCRule(request *CreateL7CCRuleRequest) (response *CreateL7CCRuleResponse, err error) {
    return c.CreateL7CCRuleWithContext(context.Background(), request)
}

// CreateL7CCRule
// This API is used to add a custom frequency control rule for layer-7 CC access (it works in the IP + Host dimension and does not support specific URIs). As it has been disused, please call the new `CreateCCFrequencyRules` API, which supports both IP + Host and URI.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateL7CCRuleWithContext(ctx context.Context, request *CreateL7CCRuleRequest) (response *CreateL7CCRuleResponse, err error) {
    if request == nil {
        request = NewCreateL7CCRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateL7CCRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateL7CCRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateL7HealthConfigRequest() (request *CreateL7HealthConfigRequest) {
    request = &CreateL7HealthConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "CreateL7HealthConfig")
    
    
    return
}

func NewCreateL7HealthConfigResponse() (response *CreateL7HealthConfigResponse) {
    response = &CreateL7HealthConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateL7HealthConfig
// This API is used to upload layer-7 health check configuration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateL7HealthConfig(request *CreateL7HealthConfigRequest) (response *CreateL7HealthConfigResponse, err error) {
    return c.CreateL7HealthConfigWithContext(context.Background(), request)
}

// CreateL7HealthConfig
// This API is used to upload layer-7 health check configuration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateL7HealthConfigWithContext(ctx context.Context, request *CreateL7HealthConfigRequest) (response *CreateL7HealthConfigResponse, err error) {
    if request == nil {
        request = NewCreateL7HealthConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateL7HealthConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateL7HealthConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateL7RuleCertRequest() (request *CreateL7RuleCertRequest) {
    request = &CreateL7RuleCertRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "CreateL7RuleCert")
    
    
    return
}

func NewCreateL7RuleCertResponse() (response *CreateL7RuleCertResponse) {
    response = &CreateL7RuleCertResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateL7RuleCert
// This API is used to configure a certificate for a layer-7 forwarding rule.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) CreateL7RuleCert(request *CreateL7RuleCertRequest) (response *CreateL7RuleCertResponse, err error) {
    return c.CreateL7RuleCertWithContext(context.Background(), request)
}

// CreateL7RuleCert
// This API is used to configure a certificate for a layer-7 forwarding rule.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) CreateL7RuleCertWithContext(ctx context.Context, request *CreateL7RuleCertRequest) (response *CreateL7RuleCertResponse, err error) {
    if request == nil {
        request = NewCreateL7RuleCertRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateL7RuleCert require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateL7RuleCertResponse()
    err = c.Send(request, response)
    return
}

func NewCreateL7RulesRequest() (request *CreateL7RulesRequest) {
    request = &CreateL7RulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "CreateL7Rules")
    
    
    return
}

func NewCreateL7RulesResponse() (response *CreateL7RulesResponse) {
    response = &CreateL7RulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateL7Rules
// This API is used to add a layer-7 (website) forwarding rule.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) CreateL7Rules(request *CreateL7RulesRequest) (response *CreateL7RulesResponse, err error) {
    return c.CreateL7RulesWithContext(context.Background(), request)
}

// CreateL7Rules
// This API is used to add a layer-7 (website) forwarding rule.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) CreateL7RulesWithContext(ctx context.Context, request *CreateL7RulesRequest) (response *CreateL7RulesResponse, err error) {
    if request == nil {
        request = NewCreateL7RulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateL7Rules require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateL7RulesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateL7RulesUploadRequest() (request *CreateL7RulesUploadRequest) {
    request = &CreateL7RulesUploadRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "CreateL7RulesUpload")
    
    
    return
}

func NewCreateL7RulesUploadResponse() (response *CreateL7RulesUploadResponse) {
    response = &CreateL7RulesUploadResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateL7RulesUpload
// This API is used to upload layer-7 forwarding rules in batches.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateL7RulesUpload(request *CreateL7RulesUploadRequest) (response *CreateL7RulesUploadResponse, err error) {
    return c.CreateL7RulesUploadWithContext(context.Background(), request)
}

// CreateL7RulesUpload
// This API is used to upload layer-7 forwarding rules in batches.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateL7RulesUploadWithContext(ctx context.Context, request *CreateL7RulesUploadRequest) (response *CreateL7RulesUploadResponse, err error) {
    if request == nil {
        request = NewCreateL7RulesUploadRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateL7RulesUpload require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateL7RulesUploadResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNetReturnRequest() (request *CreateNetReturnRequest) {
    request = &CreateNetReturnRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "CreateNetReturn")
    
    
    return
}

func NewCreateNetReturnResponse() (response *CreateNetReturnResponse) {
    response = &CreateNetReturnResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateNetReturn
// This API is used to switch to the real server in Anti-DDoS Ultimate.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateNetReturn(request *CreateNetReturnRequest) (response *CreateNetReturnResponse, err error) {
    return c.CreateNetReturnWithContext(context.Background(), request)
}

// CreateNetReturn
// This API is used to switch to the real server in Anti-DDoS Ultimate.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateNetReturnWithContext(ctx context.Context, request *CreateNetReturnRequest) (response *CreateNetReturnResponse, err error) {
    if request == nil {
        request = NewCreateNetReturnRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNetReturn require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateNetReturnResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNewL7RulesUploadRequest() (request *CreateNewL7RulesUploadRequest) {
    request = &CreateNewL7RulesUploadRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "CreateNewL7RulesUpload")
    
    
    return
}

func NewCreateNewL7RulesUploadResponse() (response *CreateNewL7RulesUploadResponse) {
    response = &CreateNewL7RulesUploadResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateNewL7RulesUpload
// This API is used to batch upload Layer-7 forwarding rules.
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateNewL7RulesUpload(request *CreateNewL7RulesUploadRequest) (response *CreateNewL7RulesUploadResponse, err error) {
    return c.CreateNewL7RulesUploadWithContext(context.Background(), request)
}

// CreateNewL7RulesUpload
// This API is used to batch upload Layer-7 forwarding rules.
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateNewL7RulesUploadWithContext(ctx context.Context, request *CreateNewL7RulesUploadRequest) (response *CreateNewL7RulesUploadResponse, err error) {
    if request == nil {
        request = NewCreateNewL7RulesUploadRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNewL7RulesUpload require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateNewL7RulesUploadResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUnblockIpRequest() (request *CreateUnblockIpRequest) {
    request = &CreateUnblockIpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "CreateUnblockIp")
    
    
    return
}

func NewCreateUnblockIpResponse() (response *CreateUnblockIpResponse) {
    response = &CreateUnblockIpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateUnblockIp
// This API is used to unblock an IP.
//
// error code that may be returned:
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) CreateUnblockIp(request *CreateUnblockIpRequest) (response *CreateUnblockIpResponse, err error) {
    return c.CreateUnblockIpWithContext(context.Background(), request)
}

// CreateUnblockIp
// This API is used to unblock an IP.
//
// error code that may be returned:
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) CreateUnblockIpWithContext(ctx context.Context, request *CreateUnblockIpRequest) (response *CreateUnblockIpResponse, err error) {
    if request == nil {
        request = NewCreateUnblockIpRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUnblockIp require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUnblockIpResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCCFrequencyRulesRequest() (request *DeleteCCFrequencyRulesRequest) {
    request = &DeleteCCFrequencyRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DeleteCCFrequencyRules")
    
    
    return
}

func NewDeleteCCFrequencyRulesResponse() (response *DeleteCCFrequencyRulesResponse) {
    response = &DeleteCCFrequencyRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCCFrequencyRules
// This API is used to delete an access frequency control rule for CC protection.
//
// error code that may be returned:
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DeleteCCFrequencyRules(request *DeleteCCFrequencyRulesRequest) (response *DeleteCCFrequencyRulesResponse, err error) {
    return c.DeleteCCFrequencyRulesWithContext(context.Background(), request)
}

// DeleteCCFrequencyRules
// This API is used to delete an access frequency control rule for CC protection.
//
// error code that may be returned:
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DeleteCCFrequencyRulesWithContext(ctx context.Context, request *DeleteCCFrequencyRulesRequest) (response *DeleteCCFrequencyRulesResponse, err error) {
    if request == nil {
        request = NewDeleteCCFrequencyRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCCFrequencyRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCCFrequencyRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCCSelfDefinePolicyRequest() (request *DeleteCCSelfDefinePolicyRequest) {
    request = &DeleteCCSelfDefinePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DeleteCCSelfDefinePolicy")
    
    
    return
}

func NewDeleteCCSelfDefinePolicyResponse() (response *DeleteCCSelfDefinePolicyResponse) {
    response = &DeleteCCSelfDefinePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCCSelfDefinePolicy
// This API is used to delete a custom CC policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteCCSelfDefinePolicy(request *DeleteCCSelfDefinePolicyRequest) (response *DeleteCCSelfDefinePolicyResponse, err error) {
    return c.DeleteCCSelfDefinePolicyWithContext(context.Background(), request)
}

// DeleteCCSelfDefinePolicy
// This API is used to delete a custom CC policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteCCSelfDefinePolicyWithContext(ctx context.Context, request *DeleteCCSelfDefinePolicyRequest) (response *DeleteCCSelfDefinePolicyResponse, err error) {
    if request == nil {
        request = NewDeleteCCSelfDefinePolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCCSelfDefinePolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCCSelfDefinePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDDoSPolicyRequest() (request *DeleteDDoSPolicyRequest) {
    request = &DeleteDDoSPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DeleteDDoSPolicy")
    
    
    return
}

func NewDeleteDDoSPolicyResponse() (response *DeleteDDoSPolicyResponse) {
    response = &DeleteDDoSPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDDoSPolicy
// This API is used to delete an advanced DDoS protection policy.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DeleteDDoSPolicy(request *DeleteDDoSPolicyRequest) (response *DeleteDDoSPolicyResponse, err error) {
    return c.DeleteDDoSPolicyWithContext(context.Background(), request)
}

// DeleteDDoSPolicy
// This API is used to delete an advanced DDoS protection policy.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DeleteDDoSPolicyWithContext(ctx context.Context, request *DeleteDDoSPolicyRequest) (response *DeleteDDoSPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteDDoSPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDDoSPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDDoSPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDDoSPolicyCaseRequest() (request *DeleteDDoSPolicyCaseRequest) {
    request = &DeleteDDoSPolicyCaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DeleteDDoSPolicyCase")
    
    
    return
}

func NewDeleteDDoSPolicyCaseResponse() (response *DeleteDDoSPolicyCaseResponse) {
    response = &DeleteDDoSPolicyCaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDDoSPolicyCase
// This API is used to delete a policy scenario.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DeleteDDoSPolicyCase(request *DeleteDDoSPolicyCaseRequest) (response *DeleteDDoSPolicyCaseResponse, err error) {
    return c.DeleteDDoSPolicyCaseWithContext(context.Background(), request)
}

// DeleteDDoSPolicyCase
// This API is used to delete a policy scenario.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DeleteDDoSPolicyCaseWithContext(ctx context.Context, request *DeleteDDoSPolicyCaseRequest) (response *DeleteDDoSPolicyCaseResponse, err error) {
    if request == nil {
        request = NewDeleteDDoSPolicyCaseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDDoSPolicyCase require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDDoSPolicyCaseResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteL4RulesRequest() (request *DeleteL4RulesRequest) {
    request = &DeleteL4RulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DeleteL4Rules")
    
    
    return
}

func NewDeleteL4RulesResponse() (response *DeleteL4RulesResponse) {
    response = &DeleteL4RulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteL4Rules
// This API is used to delete one or more layer-4 forwarding rules.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DeleteL4Rules(request *DeleteL4RulesRequest) (response *DeleteL4RulesResponse, err error) {
    return c.DeleteL4RulesWithContext(context.Background(), request)
}

// DeleteL4Rules
// This API is used to delete one or more layer-4 forwarding rules.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DeleteL4RulesWithContext(ctx context.Context, request *DeleteL4RulesRequest) (response *DeleteL4RulesResponse, err error) {
    if request == nil {
        request = NewDeleteL4RulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteL4Rules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteL4RulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteL7RulesRequest() (request *DeleteL7RulesRequest) {
    request = &DeleteL7RulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DeleteL7Rules")
    
    
    return
}

func NewDeleteL7RulesResponse() (response *DeleteL7RulesResponse) {
    response = &DeleteL7RulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteL7Rules
// This API is used to delete one or more layer-7 forwarding rules.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DeleteL7Rules(request *DeleteL7RulesRequest) (response *DeleteL7RulesResponse, err error) {
    return c.DeleteL7RulesWithContext(context.Background(), request)
}

// DeleteL7Rules
// This API is used to delete one or more layer-7 forwarding rules.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DeleteL7RulesWithContext(ctx context.Context, request *DeleteL7RulesRequest) (response *DeleteL7RulesResponse, err error) {
    if request == nil {
        request = NewDeleteL7RulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteL7Rules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteL7RulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeActionLogRequest() (request *DescribeActionLogRequest) {
    request = &DescribeActionLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeActionLog")
    
    
    return
}

func NewDescribeActionLogResponse() (response *DescribeActionLogResponse) {
    response = &DescribeActionLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeActionLog
// This API is used to get operation logs.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeActionLog(request *DescribeActionLogRequest) (response *DescribeActionLogResponse, err error) {
    return c.DescribeActionLogWithContext(context.Background(), request)
}

// DescribeActionLog
// This API is used to get operation logs.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeActionLogWithContext(ctx context.Context, request *DescribeActionLogRequest) (response *DescribeActionLogResponse, err error) {
    if request == nil {
        request = NewDescribeActionLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeActionLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeActionLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBGPIPL7RuleMaxCntRequest() (request *DescribeBGPIPL7RuleMaxCntRequest) {
    request = &DescribeBGPIPL7RuleMaxCntRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeBGPIPL7RuleMaxCnt")
    
    
    return
}

func NewDescribeBGPIPL7RuleMaxCntResponse() (response *DescribeBGPIPL7RuleMaxCntResponse) {
    response = &DescribeBGPIPL7RuleMaxCntResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBGPIPL7RuleMaxCnt
// This API is used to get the maximum number of layer-7 rules that can be added for Anti-DDoS Advanced.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeBGPIPL7RuleMaxCnt(request *DescribeBGPIPL7RuleMaxCntRequest) (response *DescribeBGPIPL7RuleMaxCntResponse, err error) {
    return c.DescribeBGPIPL7RuleMaxCntWithContext(context.Background(), request)
}

// DescribeBGPIPL7RuleMaxCnt
// This API is used to get the maximum number of layer-7 rules that can be added for Anti-DDoS Advanced.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeBGPIPL7RuleMaxCntWithContext(ctx context.Context, request *DescribeBGPIPL7RuleMaxCntRequest) (response *DescribeBGPIPL7RuleMaxCntResponse, err error) {
    if request == nil {
        request = NewDescribeBGPIPL7RuleMaxCntRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBGPIPL7RuleMaxCnt require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBGPIPL7RuleMaxCntResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaradDataRequest() (request *DescribeBaradDataRequest) {
    request = &DescribeBaradDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeBaradData")
    
    
    return
}

func NewDescribeBaradDataResponse() (response *DescribeBaradDataResponse) {
    response = &DescribeBaradDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBaradData
// This API is used to provide business forwarding metric data of Anti-DDoS services.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeBaradData(request *DescribeBaradDataRequest) (response *DescribeBaradDataResponse, err error) {
    return c.DescribeBaradDataWithContext(context.Background(), request)
}

// DescribeBaradData
// This API is used to provide business forwarding metric data of Anti-DDoS services.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeBaradDataWithContext(ctx context.Context, request *DescribeBaradDataRequest) (response *DescribeBaradDataResponse, err error) {
    if request == nil {
        request = NewDescribeBaradDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBaradData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBaradDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBasicCCThresholdRequest() (request *DescribeBasicCCThresholdRequest) {
    request = &DescribeBasicCCThresholdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeBasicCCThreshold")
    
    
    return
}

func NewDescribeBasicCCThresholdResponse() (response *DescribeBasicCCThresholdResponse) {
    response = &DescribeBasicCCThresholdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBasicCCThreshold
// This API is used to get the CC protection threshold of Anti-DDoS Basic.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeBasicCCThreshold(request *DescribeBasicCCThresholdRequest) (response *DescribeBasicCCThresholdResponse, err error) {
    return c.DescribeBasicCCThresholdWithContext(context.Background(), request)
}

// DescribeBasicCCThreshold
// This API is used to get the CC protection threshold of Anti-DDoS Basic.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeBasicCCThresholdWithContext(ctx context.Context, request *DescribeBasicCCThresholdRequest) (response *DescribeBasicCCThresholdResponse, err error) {
    if request == nil {
        request = NewDescribeBasicCCThresholdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBasicCCThreshold require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBasicCCThresholdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBasicDeviceThresholdRequest() (request *DescribeBasicDeviceThresholdRequest) {
    request = &DescribeBasicDeviceThresholdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeBasicDeviceThreshold")
    
    
    return
}

func NewDescribeBasicDeviceThresholdResponse() (response *DescribeBasicDeviceThresholdResponse) {
    response = &DescribeBasicDeviceThresholdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBasicDeviceThreshold
// This API is used to get the blackhole threshold of Anti-DDoS Basic.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeBasicDeviceThreshold(request *DescribeBasicDeviceThresholdRequest) (response *DescribeBasicDeviceThresholdResponse, err error) {
    return c.DescribeBasicDeviceThresholdWithContext(context.Background(), request)
}

// DescribeBasicDeviceThreshold
// This API is used to get the blackhole threshold of Anti-DDoS Basic.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeBasicDeviceThresholdWithContext(ctx context.Context, request *DescribeBasicDeviceThresholdRequest) (response *DescribeBasicDeviceThresholdResponse, err error) {
    if request == nil {
        request = NewDescribeBasicDeviceThresholdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBasicDeviceThreshold require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBasicDeviceThresholdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBizHttpStatusRequest() (request *DescribeBizHttpStatusRequest) {
    request = &DescribeBizHttpStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeBizHttpStatus")
    
    
    return
}

func NewDescribeBizHttpStatusResponse() (response *DescribeBizHttpStatusResponse) {
    response = &DescribeBizHttpStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBizHttpStatus
// This API is used to get the statistics on the status codes of business traffic.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeBizHttpStatus(request *DescribeBizHttpStatusRequest) (response *DescribeBizHttpStatusResponse, err error) {
    return c.DescribeBizHttpStatusWithContext(context.Background(), request)
}

// DescribeBizHttpStatus
// This API is used to get the statistics on the status codes of business traffic.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeBizHttpStatusWithContext(ctx context.Context, request *DescribeBizHttpStatusRequest) (response *DescribeBizHttpStatusResponse, err error) {
    if request == nil {
        request = NewDescribeBizHttpStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBizHttpStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBizHttpStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCCAlarmThresholdRequest() (request *DescribeCCAlarmThresholdRequest) {
    request = &DescribeCCAlarmThresholdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeCCAlarmThreshold")
    
    
    return
}

func NewDescribeCCAlarmThresholdResponse() (response *DescribeCCAlarmThresholdResponse) {
    response = &DescribeCCAlarmThresholdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCCAlarmThreshold
// This API is used to get the alarm notification threshold set for CC attacks in Anti-DDoS Pro, Anti-DDoS Advanced, Anti-DDoS Ultimate, and Chess Shield.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeCCAlarmThreshold(request *DescribeCCAlarmThresholdRequest) (response *DescribeCCAlarmThresholdResponse, err error) {
    return c.DescribeCCAlarmThresholdWithContext(context.Background(), request)
}

// DescribeCCAlarmThreshold
// This API is used to get the alarm notification threshold set for CC attacks in Anti-DDoS Pro, Anti-DDoS Advanced, Anti-DDoS Ultimate, and Chess Shield.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeCCAlarmThresholdWithContext(ctx context.Context, request *DescribeCCAlarmThresholdRequest) (response *DescribeCCAlarmThresholdResponse, err error) {
    if request == nil {
        request = NewDescribeCCAlarmThresholdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCCAlarmThreshold require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCCAlarmThresholdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCCEvListRequest() (request *DescribeCCEvListRequest) {
    request = &DescribeCCEvListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeCCEvList")
    
    
    return
}

func NewDescribeCCEvListResponse() (response *DescribeCCEvListResponse) {
    response = &DescribeCCEvListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCCEvList
// This API is used to get the CC attack event list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeCCEvList(request *DescribeCCEvListRequest) (response *DescribeCCEvListResponse, err error) {
    return c.DescribeCCEvListWithContext(context.Background(), request)
}

// DescribeCCEvList
// This API is used to get the CC attack event list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeCCEvListWithContext(ctx context.Context, request *DescribeCCEvListRequest) (response *DescribeCCEvListResponse, err error) {
    if request == nil {
        request = NewDescribeCCEvListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCCEvList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCCEvListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCCFrequencyRulesRequest() (request *DescribeCCFrequencyRulesRequest) {
    request = &DescribeCCFrequencyRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeCCFrequencyRules")
    
    
    return
}

func NewDescribeCCFrequencyRulesResponse() (response *DescribeCCFrequencyRulesResponse) {
    response = &DescribeCCFrequencyRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCCFrequencyRules
// This API is used to get an access frequency control rule for CC protection.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeCCFrequencyRules(request *DescribeCCFrequencyRulesRequest) (response *DescribeCCFrequencyRulesResponse, err error) {
    return c.DescribeCCFrequencyRulesWithContext(context.Background(), request)
}

// DescribeCCFrequencyRules
// This API is used to get an access frequency control rule for CC protection.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeCCFrequencyRulesWithContext(ctx context.Context, request *DescribeCCFrequencyRulesRequest) (response *DescribeCCFrequencyRulesResponse, err error) {
    if request == nil {
        request = NewDescribeCCFrequencyRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCCFrequencyRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCCFrequencyRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCCIpAllowDenyRequest() (request *DescribeCCIpAllowDenyRequest) {
    request = &DescribeCCIpAllowDenyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeCCIpAllowDeny")
    
    
    return
}

func NewDescribeCCIpAllowDenyResponse() (response *DescribeCCIpAllowDenyResponse) {
    response = &DescribeCCIpAllowDenyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCCIpAllowDeny
// This API is used to get the CC IP blocklist/allowlist.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeCCIpAllowDeny(request *DescribeCCIpAllowDenyRequest) (response *DescribeCCIpAllowDenyResponse, err error) {
    return c.DescribeCCIpAllowDenyWithContext(context.Background(), request)
}

// DescribeCCIpAllowDeny
// This API is used to get the CC IP blocklist/allowlist.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeCCIpAllowDenyWithContext(ctx context.Context, request *DescribeCCIpAllowDenyRequest) (response *DescribeCCIpAllowDenyResponse, err error) {
    if request == nil {
        request = NewDescribeCCIpAllowDenyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCCIpAllowDeny require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCCIpAllowDenyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCCSelfDefinePolicyRequest() (request *DescribeCCSelfDefinePolicyRequest) {
    request = &DescribeCCSelfDefinePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeCCSelfDefinePolicy")
    
    
    return
}

func NewDescribeCCSelfDefinePolicyResponse() (response *DescribeCCSelfDefinePolicyResponse) {
    response = &DescribeCCSelfDefinePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCCSelfDefinePolicy
// This API is used to get a custom CC policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeCCSelfDefinePolicy(request *DescribeCCSelfDefinePolicyRequest) (response *DescribeCCSelfDefinePolicyResponse, err error) {
    return c.DescribeCCSelfDefinePolicyWithContext(context.Background(), request)
}

// DescribeCCSelfDefinePolicy
// This API is used to get a custom CC policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeCCSelfDefinePolicyWithContext(ctx context.Context, request *DescribeCCSelfDefinePolicyRequest) (response *DescribeCCSelfDefinePolicyResponse, err error) {
    if request == nil {
        request = NewDescribeCCSelfDefinePolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCCSelfDefinePolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCCSelfDefinePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCCTrendRequest() (request *DescribeCCTrendRequest) {
    request = &DescribeCCTrendRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeCCTrend")
    
    
    return
}

func NewDescribeCCTrendResponse() (response *DescribeCCTrendResponse) {
    response = &DescribeCCTrendResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCCTrend
// This API is used to get CC attack metric data, including total requests peak (QPS) and attack requests (QPS).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeCCTrend(request *DescribeCCTrendRequest) (response *DescribeCCTrendResponse, err error) {
    return c.DescribeCCTrendWithContext(context.Background(), request)
}

// DescribeCCTrend
// This API is used to get CC attack metric data, including total requests peak (QPS) and attack requests (QPS).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeCCTrendWithContext(ctx context.Context, request *DescribeCCTrendRequest) (response *DescribeCCTrendResponse, err error) {
    if request == nil {
        request = NewDescribeCCTrendRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCCTrend require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCCTrendResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCCUrlAllowRequest() (request *DescribeCCUrlAllowRequest) {
    request = &DescribeCCUrlAllowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeCCUrlAllow")
    
    
    return
}

func NewDescribeCCUrlAllowResponse() (response *DescribeCCUrlAllowResponse) {
    response = &DescribeCCUrlAllowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCCUrlAllow
// This API is used to get the CC URL allowlist.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeCCUrlAllow(request *DescribeCCUrlAllowRequest) (response *DescribeCCUrlAllowResponse, err error) {
    return c.DescribeCCUrlAllowWithContext(context.Background(), request)
}

// DescribeCCUrlAllow
// This API is used to get the CC URL allowlist.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeCCUrlAllowWithContext(ctx context.Context, request *DescribeCCUrlAllowRequest) (response *DescribeCCUrlAllowResponse, err error) {
    if request == nil {
        request = NewDescribeCCUrlAllowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCCUrlAllow require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCCUrlAllowResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSAlarmThresholdRequest() (request *DescribeDDoSAlarmThresholdRequest) {
    request = &DescribeDDoSAlarmThresholdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeDDoSAlarmThreshold")
    
    
    return
}

func NewDescribeDDoSAlarmThresholdResponse() (response *DescribeDDoSAlarmThresholdResponse) {
    response = &DescribeDDoSAlarmThresholdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDDoSAlarmThreshold
// This API is used to get the alarm notification threshold set for DDoS attacks in Anti-DDoS Pro, Anti-DDoS Advanced, Anti-DDoS Ultimate, and Chess Shield.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDDoSAlarmThreshold(request *DescribeDDoSAlarmThresholdRequest) (response *DescribeDDoSAlarmThresholdResponse, err error) {
    return c.DescribeDDoSAlarmThresholdWithContext(context.Background(), request)
}

// DescribeDDoSAlarmThreshold
// This API is used to get the alarm notification threshold set for DDoS attacks in Anti-DDoS Pro, Anti-DDoS Advanced, Anti-DDoS Ultimate, and Chess Shield.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDDoSAlarmThresholdWithContext(ctx context.Context, request *DescribeDDoSAlarmThresholdRequest) (response *DescribeDDoSAlarmThresholdResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSAlarmThresholdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDDoSAlarmThreshold require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDDoSAlarmThresholdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSAttackIPRegionMapRequest() (request *DescribeDDoSAttackIPRegionMapRequest) {
    request = &DescribeDDoSAttackIPRegionMapRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeDDoSAttackIPRegionMap")
    
    
    return
}

func NewDescribeDDoSAttackIPRegionMapResponse() (response *DescribeDDoSAttackIPRegionMapResponse) {
    response = &DescribeDDoSAttackIPRegionMapResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDDoSAttackIPRegionMap
// This API is used to get the geographical distribution map of DDoS attack source IPs. It supports display by global regions and Chinese provinces.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDDoSAttackIPRegionMap(request *DescribeDDoSAttackIPRegionMapRequest) (response *DescribeDDoSAttackIPRegionMapResponse, err error) {
    return c.DescribeDDoSAttackIPRegionMapWithContext(context.Background(), request)
}

// DescribeDDoSAttackIPRegionMap
// This API is used to get the geographical distribution map of DDoS attack source IPs. It supports display by global regions and Chinese provinces.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDDoSAttackIPRegionMapWithContext(ctx context.Context, request *DescribeDDoSAttackIPRegionMapRequest) (response *DescribeDDoSAttackIPRegionMapResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSAttackIPRegionMapRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDDoSAttackIPRegionMap require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDDoSAttackIPRegionMapResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSAttackSourceRequest() (request *DescribeDDoSAttackSourceRequest) {
    request = &DescribeDDoSAttackSourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeDDoSAttackSource")
    
    
    return
}

func NewDescribeDDoSAttackSourceResponse() (response *DescribeDDoSAttackSourceResponse) {
    response = &DescribeDDoSAttackSourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDDoSAttackSource
// This API is used to get the DDoS attack source list.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDDoSAttackSource(request *DescribeDDoSAttackSourceRequest) (response *DescribeDDoSAttackSourceResponse, err error) {
    return c.DescribeDDoSAttackSourceWithContext(context.Background(), request)
}

// DescribeDDoSAttackSource
// This API is used to get the DDoS attack source list.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDDoSAttackSourceWithContext(ctx context.Context, request *DescribeDDoSAttackSourceRequest) (response *DescribeDDoSAttackSourceResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSAttackSourceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDDoSAttackSource require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDDoSAttackSourceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSCountRequest() (request *DescribeDDoSCountRequest) {
    request = &DescribeDDoSCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeDDoSCount")
    
    
    return
}

func NewDescribeDDoSCountResponse() (response *DescribeDDoSCountResponse) {
    response = &DescribeDDoSCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDDoSCount
// This API is used to get the DDoS attack proportion analysis.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDDoSCount(request *DescribeDDoSCountRequest) (response *DescribeDDoSCountResponse, err error) {
    return c.DescribeDDoSCountWithContext(context.Background(), request)
}

// DescribeDDoSCount
// This API is used to get the DDoS attack proportion analysis.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDDoSCountWithContext(ctx context.Context, request *DescribeDDoSCountRequest) (response *DescribeDDoSCountResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSCountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDDoSCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDDoSCountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSDefendStatusRequest() (request *DescribeDDoSDefendStatusRequest) {
    request = &DescribeDDoSDefendStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeDDoSDefendStatus")
    
    
    return
}

func NewDescribeDDoSDefendStatusResponse() (response *DescribeDDoSDefendStatusResponse) {
    response = &DescribeDDoSDefendStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDDoSDefendStatus
// This API is used to get the DDoS protection status (temporarily disabled status). It is supported for Anti-DDoS Basic, single IP instance, multi-IP instance, Anti-DDoS Advanced, and Anti-DDoS Ultimate. It is used to query whether DDoS protection is temporarily disabled, and if yes, return parameters such as temporary disablement duration.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDDoSDefendStatus(request *DescribeDDoSDefendStatusRequest) (response *DescribeDDoSDefendStatusResponse, err error) {
    return c.DescribeDDoSDefendStatusWithContext(context.Background(), request)
}

// DescribeDDoSDefendStatus
// This API is used to get the DDoS protection status (temporarily disabled status). It is supported for Anti-DDoS Basic, single IP instance, multi-IP instance, Anti-DDoS Advanced, and Anti-DDoS Ultimate. It is used to query whether DDoS protection is temporarily disabled, and if yes, return parameters such as temporary disablement duration.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDDoSDefendStatusWithContext(ctx context.Context, request *DescribeDDoSDefendStatusRequest) (response *DescribeDDoSDefendStatusResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSDefendStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDDoSDefendStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDDoSDefendStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSEvInfoRequest() (request *DescribeDDoSEvInfoRequest) {
    request = &DescribeDDoSEvInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeDDoSEvInfo")
    
    
    return
}

func NewDescribeDDoSEvInfoResponse() (response *DescribeDDoSEvInfoResponse) {
    response = &DescribeDDoSEvInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDDoSEvInfo
// This API is used to get details of a specific DDoS attack.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDDoSEvInfo(request *DescribeDDoSEvInfoRequest) (response *DescribeDDoSEvInfoResponse, err error) {
    return c.DescribeDDoSEvInfoWithContext(context.Background(), request)
}

// DescribeDDoSEvInfo
// This API is used to get details of a specific DDoS attack.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDDoSEvInfoWithContext(ctx context.Context, request *DescribeDDoSEvInfoRequest) (response *DescribeDDoSEvInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSEvInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDDoSEvInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDDoSEvInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSEvListRequest() (request *DescribeDDoSEvListRequest) {
    request = &DescribeDDoSEvListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeDDoSEvList")
    
    
    return
}

func NewDescribeDDoSEvListResponse() (response *DescribeDDoSEvListResponse) {
    response = &DescribeDDoSEvListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDDoSEvList
// This API is used to get the DDoS attack event list.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDDoSEvList(request *DescribeDDoSEvListRequest) (response *DescribeDDoSEvListResponse, err error) {
    return c.DescribeDDoSEvListWithContext(context.Background(), request)
}

// DescribeDDoSEvList
// This API is used to get the DDoS attack event list.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDDoSEvListWithContext(ctx context.Context, request *DescribeDDoSEvListRequest) (response *DescribeDDoSEvListResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSEvListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDDoSEvList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDDoSEvListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSIpLogRequest() (request *DescribeDDoSIpLogRequest) {
    request = &DescribeDDoSIpLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeDDoSIpLog")
    
    
    return
}

func NewDescribeDDoSIpLogResponse() (response *DescribeDDoSIpLogResponse) {
    response = &DescribeDDoSIpLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDDoSIpLog
// This API is used to get a DDoS IP attack log.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDDoSIpLog(request *DescribeDDoSIpLogRequest) (response *DescribeDDoSIpLogResponse, err error) {
    return c.DescribeDDoSIpLogWithContext(context.Background(), request)
}

// DescribeDDoSIpLog
// This API is used to get a DDoS IP attack log.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDDoSIpLogWithContext(ctx context.Context, request *DescribeDDoSIpLogRequest) (response *DescribeDDoSIpLogResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSIpLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDDoSIpLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDDoSIpLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSNetCountRequest() (request *DescribeDDoSNetCountRequest) {
    request = &DescribeDDoSNetCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeDDoSNetCount")
    
    
    return
}

func NewDescribeDDoSNetCountResponse() (response *DescribeDDoSNetCountResponse) {
    response = &DescribeDDoSNetCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDDoSNetCount
// This API is used to get the DDoS attack proportion analysis for an Anti-DDoS Ultimate resource.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDDoSNetCount(request *DescribeDDoSNetCountRequest) (response *DescribeDDoSNetCountResponse, err error) {
    return c.DescribeDDoSNetCountWithContext(context.Background(), request)
}

// DescribeDDoSNetCount
// This API is used to get the DDoS attack proportion analysis for an Anti-DDoS Ultimate resource.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDDoSNetCountWithContext(ctx context.Context, request *DescribeDDoSNetCountRequest) (response *DescribeDDoSNetCountResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSNetCountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDDoSNetCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDDoSNetCountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSNetEvInfoRequest() (request *DescribeDDoSNetEvInfoRequest) {
    request = &DescribeDDoSNetEvInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeDDoSNetEvInfo")
    
    
    return
}

func NewDescribeDDoSNetEvInfoResponse() (response *DescribeDDoSNetEvInfoResponse) {
    response = &DescribeDDoSNetEvInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDDoSNetEvInfo
// This API is used to get the DDoS attack event details of an Anti-DDoS Ultimate resource.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDDoSNetEvInfo(request *DescribeDDoSNetEvInfoRequest) (response *DescribeDDoSNetEvInfoResponse, err error) {
    return c.DescribeDDoSNetEvInfoWithContext(context.Background(), request)
}

// DescribeDDoSNetEvInfo
// This API is used to get the DDoS attack event details of an Anti-DDoS Ultimate resource.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDDoSNetEvInfoWithContext(ctx context.Context, request *DescribeDDoSNetEvInfoRequest) (response *DescribeDDoSNetEvInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSNetEvInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDDoSNetEvInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDDoSNetEvInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSNetEvListRequest() (request *DescribeDDoSNetEvListRequest) {
    request = &DescribeDDoSNetEvListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeDDoSNetEvList")
    
    
    return
}

func NewDescribeDDoSNetEvListResponse() (response *DescribeDDoSNetEvListResponse) {
    response = &DescribeDDoSNetEvListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDDoSNetEvList
// This API is used to get the DDoS attack event list of an Anti-DDoS Ultimate resource.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDDoSNetEvList(request *DescribeDDoSNetEvListRequest) (response *DescribeDDoSNetEvListResponse, err error) {
    return c.DescribeDDoSNetEvListWithContext(context.Background(), request)
}

// DescribeDDoSNetEvList
// This API is used to get the DDoS attack event list of an Anti-DDoS Ultimate resource.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDDoSNetEvListWithContext(ctx context.Context, request *DescribeDDoSNetEvListRequest) (response *DescribeDDoSNetEvListResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSNetEvListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDDoSNetEvList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDDoSNetEvListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSNetIpLogRequest() (request *DescribeDDoSNetIpLogRequest) {
    request = &DescribeDDoSNetIpLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeDDoSNetIpLog")
    
    
    return
}

func NewDescribeDDoSNetIpLogResponse() (response *DescribeDDoSNetIpLogResponse) {
    response = &DescribeDDoSNetIpLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDDoSNetIpLog
// This API is used to get the DDoS IP attack logs of an Anti-DDoS Ultimate resource.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDDoSNetIpLog(request *DescribeDDoSNetIpLogRequest) (response *DescribeDDoSNetIpLogResponse, err error) {
    return c.DescribeDDoSNetIpLogWithContext(context.Background(), request)
}

// DescribeDDoSNetIpLog
// This API is used to get the DDoS IP attack logs of an Anti-DDoS Ultimate resource.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDDoSNetIpLogWithContext(ctx context.Context, request *DescribeDDoSNetIpLogRequest) (response *DescribeDDoSNetIpLogResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSNetIpLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDDoSNetIpLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDDoSNetIpLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSNetTrendRequest() (request *DescribeDDoSNetTrendRequest) {
    request = &DescribeDDoSNetTrendRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeDDoSNetTrend")
    
    
    return
}

func NewDescribeDDoSNetTrendResponse() (response *DescribeDDoSNetTrendResponse) {
    response = &DescribeDDoSNetTrendResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDDoSNetTrend
// This API is used to get the DDoS attack metric data of an Anti-DDoS Ultimate resource.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDDoSNetTrend(request *DescribeDDoSNetTrendRequest) (response *DescribeDDoSNetTrendResponse, err error) {
    return c.DescribeDDoSNetTrendWithContext(context.Background(), request)
}

// DescribeDDoSNetTrend
// This API is used to get the DDoS attack metric data of an Anti-DDoS Ultimate resource.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDDoSNetTrendWithContext(ctx context.Context, request *DescribeDDoSNetTrendRequest) (response *DescribeDDoSNetTrendResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSNetTrendRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDDoSNetTrend require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDDoSNetTrendResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSPolicyRequest() (request *DescribeDDoSPolicyRequest) {
    request = &DescribeDDoSPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeDDoSPolicy")
    
    
    return
}

func NewDescribeDDoSPolicyResponse() (response *DescribeDDoSPolicyResponse) {
    response = &DescribeDDoSPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDDoSPolicy
// This API is used to get an advanced DDoS policy.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDDoSPolicy(request *DescribeDDoSPolicyRequest) (response *DescribeDDoSPolicyResponse, err error) {
    return c.DescribeDDoSPolicyWithContext(context.Background(), request)
}

// DescribeDDoSPolicy
// This API is used to get an advanced DDoS policy.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDDoSPolicyWithContext(ctx context.Context, request *DescribeDDoSPolicyRequest) (response *DescribeDDoSPolicyResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDDoSPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDDoSPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSTrendRequest() (request *DescribeDDoSTrendRequest) {
    request = &DescribeDDoSTrendRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeDDoSTrend")
    
    
    return
}

func NewDescribeDDoSTrendResponse() (response *DescribeDDoSTrendResponse) {
    response = &DescribeDDoSTrendResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDDoSTrend
// This API is used to get the data of DDoS attack traffic bandwidth and attack packet rate.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDDoSTrend(request *DescribeDDoSTrendRequest) (response *DescribeDDoSTrendResponse, err error) {
    return c.DescribeDDoSTrendWithContext(context.Background(), request)
}

// DescribeDDoSTrend
// This API is used to get the data of DDoS attack traffic bandwidth and attack packet rate.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDDoSTrendWithContext(ctx context.Context, request *DescribeDDoSTrendRequest) (response *DescribeDDoSTrendResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSTrendRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDDoSTrend require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDDoSTrendResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSUsedStatisRequest() (request *DescribeDDoSUsedStatisRequest) {
    request = &DescribeDDoSUsedStatisRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeDDoSUsedStatis")
    
    
    return
}

func NewDescribeDDoSUsedStatisResponse() (response *DescribeDDoSUsedStatisResponse) {
    response = &DescribeDDoSUsedStatisResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDDoSUsedStatis
// This API is used to count the number of days of Anti-DDoS resource use and number of DDoS attacks defended against.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDDoSUsedStatis(request *DescribeDDoSUsedStatisRequest) (response *DescribeDDoSUsedStatisResponse, err error) {
    return c.DescribeDDoSUsedStatisWithContext(context.Background(), request)
}

// DescribeDDoSUsedStatis
// This API is used to count the number of days of Anti-DDoS resource use and number of DDoS attacks defended against.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDDoSUsedStatisWithContext(ctx context.Context, request *DescribeDDoSUsedStatisRequest) (response *DescribeDDoSUsedStatisResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSUsedStatisRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDDoSUsedStatis require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDDoSUsedStatisResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIPProductInfoRequest() (request *DescribeIPProductInfoRequest) {
    request = &DescribeIPProductInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeIPProductInfo")
    
    
    return
}

func NewDescribeIPProductInfoResponse() (response *DescribeIPProductInfoResponse) {
    response = &DescribeIPProductInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIPProductInfo
// This API is used to get the Tencent Cloud asset information corresponding to an IP of a single IP instance or multi-IP instance, which is supported only for IPs of single IP instances and multi-IP instances.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeIPProductInfo(request *DescribeIPProductInfoRequest) (response *DescribeIPProductInfoResponse, err error) {
    return c.DescribeIPProductInfoWithContext(context.Background(), request)
}

// DescribeIPProductInfo
// This API is used to get the Tencent Cloud asset information corresponding to an IP of a single IP instance or multi-IP instance, which is supported only for IPs of single IP instances and multi-IP instances.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeIPProductInfoWithContext(ctx context.Context, request *DescribeIPProductInfoRequest) (response *DescribeIPProductInfoResponse, err error) {
    if request == nil {
        request = NewDescribeIPProductInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIPProductInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIPProductInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInsurePacksRequest() (request *DescribeInsurePacksRequest) {
    request = &DescribeInsurePacksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeInsurePacks")
    
    
    return
}

func NewDescribeInsurePacksResponse() (response *DescribeInsurePacksResponse) {
    response = &DescribeInsurePacksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInsurePacks
// This API is used to get the guarantee package list.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeInsurePacks(request *DescribeInsurePacksRequest) (response *DescribeInsurePacksResponse, err error) {
    return c.DescribeInsurePacksWithContext(context.Background(), request)
}

// DescribeInsurePacks
// This API is used to get the guarantee package list.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeInsurePacksWithContext(ctx context.Context, request *DescribeInsurePacksRequest) (response *DescribeInsurePacksResponse, err error) {
    if request == nil {
        request = NewDescribeInsurePacksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInsurePacks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInsurePacksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIpBlockListRequest() (request *DescribeIpBlockListRequest) {
    request = &DescribeIpBlockListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeIpBlockList")
    
    
    return
}

func NewDescribeIpBlockListResponse() (response *DescribeIpBlockListResponse) {
    response = &DescribeIpBlockListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIpBlockList
// This API is used to get the blocked IP list.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeIpBlockList(request *DescribeIpBlockListRequest) (response *DescribeIpBlockListResponse, err error) {
    return c.DescribeIpBlockListWithContext(context.Background(), request)
}

// DescribeIpBlockList
// This API is used to get the blocked IP list.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeIpBlockListWithContext(ctx context.Context, request *DescribeIpBlockListRequest) (response *DescribeIpBlockListResponse, err error) {
    if request == nil {
        request = NewDescribeIpBlockListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIpBlockList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIpBlockListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIpUnBlockListRequest() (request *DescribeIpUnBlockListRequest) {
    request = &DescribeIpUnBlockListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeIpUnBlockList")
    
    
    return
}

func NewDescribeIpUnBlockListResponse() (response *DescribeIpUnBlockListResponse) {
    response = &DescribeIpUnBlockListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIpUnBlockList
// This API is used to get the IP unblocking records.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeIpUnBlockList(request *DescribeIpUnBlockListRequest) (response *DescribeIpUnBlockListResponse, err error) {
    return c.DescribeIpUnBlockListWithContext(context.Background(), request)
}

// DescribeIpUnBlockList
// This API is used to get the IP unblocking records.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeIpUnBlockListWithContext(ctx context.Context, request *DescribeIpUnBlockListRequest) (response *DescribeIpUnBlockListResponse, err error) {
    if request == nil {
        request = NewDescribeIpUnBlockListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIpUnBlockList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIpUnBlockListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeL4HealthConfigRequest() (request *DescribeL4HealthConfigRequest) {
    request = &DescribeL4HealthConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeL4HealthConfig")
    
    
    return
}

func NewDescribeL4HealthConfigResponse() (response *DescribeL4HealthConfigResponse) {
    response = &DescribeL4HealthConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeL4HealthConfig
// This API is used to export the layer-4 health check configuration.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeL4HealthConfig(request *DescribeL4HealthConfigRequest) (response *DescribeL4HealthConfigResponse, err error) {
    return c.DescribeL4HealthConfigWithContext(context.Background(), request)
}

// DescribeL4HealthConfig
// This API is used to export the layer-4 health check configuration.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeL4HealthConfigWithContext(ctx context.Context, request *DescribeL4HealthConfigRequest) (response *DescribeL4HealthConfigResponse, err error) {
    if request == nil {
        request = NewDescribeL4HealthConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeL4HealthConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeL4HealthConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeL4RulesErrHealthRequest() (request *DescribeL4RulesErrHealthRequest) {
    request = &DescribeL4RulesErrHealthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeL4RulesErrHealth")
    
    
    return
}

func NewDescribeL4RulesErrHealthResponse() (response *DescribeL4RulesErrHealthResponse) {
    response = &DescribeL4RulesErrHealthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeL4RulesErrHealth
// This API is used to get the exception result of a layer-4 forwarding rule health check.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeL4RulesErrHealth(request *DescribeL4RulesErrHealthRequest) (response *DescribeL4RulesErrHealthResponse, err error) {
    return c.DescribeL4RulesErrHealthWithContext(context.Background(), request)
}

// DescribeL4RulesErrHealth
// This API is used to get the exception result of a layer-4 forwarding rule health check.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeL4RulesErrHealthWithContext(ctx context.Context, request *DescribeL4RulesErrHealthRequest) (response *DescribeL4RulesErrHealthResponse, err error) {
    if request == nil {
        request = NewDescribeL4RulesErrHealthRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeL4RulesErrHealth require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeL4RulesErrHealthResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeL7HealthConfigRequest() (request *DescribeL7HealthConfigRequest) {
    request = &DescribeL7HealthConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeL7HealthConfig")
    
    
    return
}

func NewDescribeL7HealthConfigResponse() (response *DescribeL7HealthConfigResponse) {
    response = &DescribeL7HealthConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeL7HealthConfig
// This API is used to export the layer-7 health check configuration.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeL7HealthConfig(request *DescribeL7HealthConfigRequest) (response *DescribeL7HealthConfigResponse, err error) {
    return c.DescribeL7HealthConfigWithContext(context.Background(), request)
}

// DescribeL7HealthConfig
// This API is used to export the layer-7 health check configuration.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeL7HealthConfigWithContext(ctx context.Context, request *DescribeL7HealthConfigRequest) (response *DescribeL7HealthConfigResponse, err error) {
    if request == nil {
        request = NewDescribeL7HealthConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeL7HealthConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeL7HealthConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePackIndexRequest() (request *DescribePackIndexRequest) {
    request = &DescribePackIndexRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribePackIndex")
    
    
    return
}

func NewDescribePackIndexResponse() (response *DescribePackIndexResponse) {
    response = &DescribePackIndexResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePackIndex
// This API is used to get the product overview statistics. It is supported for Anti-DDoS Pro, Anti-DDoS Advanced, and Anti-DDoS Ultimate.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribePackIndex(request *DescribePackIndexRequest) (response *DescribePackIndexResponse, err error) {
    return c.DescribePackIndexWithContext(context.Background(), request)
}

// DescribePackIndex
// This API is used to get the product overview statistics. It is supported for Anti-DDoS Pro, Anti-DDoS Advanced, and Anti-DDoS Ultimate.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribePackIndexWithContext(ctx context.Context, request *DescribePackIndexRequest) (response *DescribePackIndexResponse, err error) {
    if request == nil {
        request = NewDescribePackIndexRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePackIndex require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePackIndexResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePcapRequest() (request *DescribePcapRequest) {
    request = &DescribePcapRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribePcap")
    
    
    return
}

func NewDescribePcapResponse() (response *DescribePcapResponse) {
    response = &DescribePcapResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePcap
// This API is used to download the PCAP packet of an attack event.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribePcap(request *DescribePcapRequest) (response *DescribePcapResponse, err error) {
    return c.DescribePcapWithContext(context.Background(), request)
}

// DescribePcap
// This API is used to download the PCAP packet of an attack event.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribePcapWithContext(ctx context.Context, request *DescribePcapRequest) (response *DescribePcapResponse, err error) {
    if request == nil {
        request = NewDescribePcapRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePcap require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePcapResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePolicyCaseRequest() (request *DescribePolicyCaseRequest) {
    request = &DescribePolicyCaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribePolicyCase")
    
    
    return
}

func NewDescribePolicyCaseResponse() (response *DescribePolicyCaseResponse) {
    response = &DescribePolicyCaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePolicyCase
// This API is used to get the policy scenario.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribePolicyCase(request *DescribePolicyCaseRequest) (response *DescribePolicyCaseResponse, err error) {
    return c.DescribePolicyCaseWithContext(context.Background(), request)
}

// DescribePolicyCase
// This API is used to get the policy scenario.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribePolicyCaseWithContext(ctx context.Context, request *DescribePolicyCaseRequest) (response *DescribePolicyCaseResponse, err error) {
    if request == nil {
        request = NewDescribePolicyCaseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePolicyCase require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePolicyCaseResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResIpListRequest() (request *DescribeResIpListRequest) {
    request = &DescribeResIpListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeResIpList")
    
    
    return
}

func NewDescribeResIpListResponse() (response *DescribeResIpListResponse) {
    response = &DescribeResIpListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeResIpList
// This API is used to get the IP list of a resource.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeResIpList(request *DescribeResIpListRequest) (response *DescribeResIpListResponse, err error) {
    return c.DescribeResIpListWithContext(context.Background(), request)
}

// DescribeResIpList
// This API is used to get the IP list of a resource.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeResIpListWithContext(ctx context.Context, request *DescribeResIpListRequest) (response *DescribeResIpListResponse, err error) {
    if request == nil {
        request = NewDescribeResIpListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResIpList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResIpListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourceListRequest() (request *DescribeResourceListRequest) {
    request = &DescribeResourceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeResourceList")
    
    
    return
}

func NewDescribeResourceListResponse() (response *DescribeResourceListResponse) {
    response = &DescribeResourceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeResourceList
// This API is used to get the resource list.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeResourceList(request *DescribeResourceListRequest) (response *DescribeResourceListResponse, err error) {
    return c.DescribeResourceListWithContext(context.Background(), request)
}

// DescribeResourceList
// This API is used to get the resource list.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeResourceListWithContext(ctx context.Context, request *DescribeResourceListRequest) (response *DescribeResourceListResponse, err error) {
    if request == nil {
        request = NewDescribeResourceListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResourceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResourceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRuleSetsRequest() (request *DescribeRuleSetsRequest) {
    request = &DescribeRuleSetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeRuleSets")
    
    
    return
}

func NewDescribeRuleSetsResponse() (response *DescribeRuleSetsResponse) {
    response = &DescribeRuleSetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRuleSets
// This API is used to get the number of rules of a resource.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeRuleSets(request *DescribeRuleSetsRequest) (response *DescribeRuleSetsResponse, err error) {
    return c.DescribeRuleSetsWithContext(context.Background(), request)
}

// DescribeRuleSets
// This API is used to get the number of rules of a resource.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeRuleSetsWithContext(ctx context.Context, request *DescribeRuleSetsRequest) (response *DescribeRuleSetsResponse, err error) {
    if request == nil {
        request = NewDescribeRuleSetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRuleSets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRuleSetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSchedulingDomainListRequest() (request *DescribeSchedulingDomainListRequest) {
    request = &DescribeSchedulingDomainListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeSchedulingDomainList")
    
    
    return
}

func NewDescribeSchedulingDomainListResponse() (response *DescribeSchedulingDomainListResponse) {
    response = &DescribeSchedulingDomainListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSchedulingDomainList
// Get scheduling domain name list
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeSchedulingDomainList(request *DescribeSchedulingDomainListRequest) (response *DescribeSchedulingDomainListResponse, err error) {
    return c.DescribeSchedulingDomainListWithContext(context.Background(), request)
}

// DescribeSchedulingDomainList
// Get scheduling domain name list
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeSchedulingDomainListWithContext(ctx context.Context, request *DescribeSchedulingDomainListRequest) (response *DescribeSchedulingDomainListResponse, err error) {
    if request == nil {
        request = NewDescribeSchedulingDomainListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSchedulingDomainList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSchedulingDomainListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecIndexRequest() (request *DescribeSecIndexRequest) {
    request = &DescribeSecIndexRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeSecIndex")
    
    
    return
}

func NewDescribeSecIndexResponse() (response *DescribeSecIndexResponse) {
    response = &DescribeSecIndexResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSecIndex
// This API is used to get the security statistics of the current month.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeSecIndex(request *DescribeSecIndexRequest) (response *DescribeSecIndexResponse, err error) {
    return c.DescribeSecIndexWithContext(context.Background(), request)
}

// DescribeSecIndex
// This API is used to get the security statistics of the current month.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeSecIndexWithContext(ctx context.Context, request *DescribeSecIndexRequest) (response *DescribeSecIndexResponse, err error) {
    if request == nil {
        request = NewDescribeSecIndexRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecIndex require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecIndexResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSourceIpSegmentRequest() (request *DescribeSourceIpSegmentRequest) {
    request = &DescribeSourceIpSegmentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeSourceIpSegment")
    
    
    return
}

func NewDescribeSourceIpSegmentResponse() (response *DescribeSourceIpSegmentResponse) {
    response = &DescribeSourceIpSegmentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSourceIpSegment
// This API is used to get the intermediate IP range. It is supported for Anti-DDoS Advanced and Anti-DDoS Ultimate.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeSourceIpSegment(request *DescribeSourceIpSegmentRequest) (response *DescribeSourceIpSegmentResponse, err error) {
    return c.DescribeSourceIpSegmentWithContext(context.Background(), request)
}

// DescribeSourceIpSegment
// This API is used to get the intermediate IP range. It is supported for Anti-DDoS Advanced and Anti-DDoS Ultimate.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeSourceIpSegmentWithContext(ctx context.Context, request *DescribeSourceIpSegmentRequest) (response *DescribeSourceIpSegmentResponse, err error) {
    if request == nil {
        request = NewDescribeSourceIpSegmentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSourceIpSegment require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSourceIpSegmentResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTransmitStatisRequest() (request *DescribeTransmitStatisRequest) {
    request = &DescribeTransmitStatisRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeTransmitStatis")
    
    
    return
}

func NewDescribeTransmitStatisResponse() (response *DescribeTransmitStatisResponse) {
    response = &DescribeTransmitStatisResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTransmitStatis
// This API is used to get the business forwarding statistics, including forwarded traffic and packet forwarding rate.
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTransmitStatis(request *DescribeTransmitStatisRequest) (response *DescribeTransmitStatisResponse, err error) {
    return c.DescribeTransmitStatisWithContext(context.Background(), request)
}

// DescribeTransmitStatis
// This API is used to get the business forwarding statistics, including forwarded traffic and packet forwarding rate.
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTransmitStatisWithContext(ctx context.Context, request *DescribeTransmitStatisRequest) (response *DescribeTransmitStatisResponse, err error) {
    if request == nil {
        request = NewDescribeTransmitStatisRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTransmitStatis require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTransmitStatisResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUnBlockStatisRequest() (request *DescribeUnBlockStatisRequest) {
    request = &DescribeUnBlockStatisRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeUnBlockStatis")
    
    
    return
}

func NewDescribeUnBlockStatisResponse() (response *DescribeUnBlockStatisResponse) {
    response = &DescribeUnBlockStatisResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUnBlockStatis
// This API is used to get the number of blackhole unblockings.
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUnBlockStatis(request *DescribeUnBlockStatisRequest) (response *DescribeUnBlockStatisResponse, err error) {
    return c.DescribeUnBlockStatisWithContext(context.Background(), request)
}

// DescribeUnBlockStatis
// This API is used to get the number of blackhole unblockings.
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUnBlockStatisWithContext(ctx context.Context, request *DescribeUnBlockStatisRequest) (response *DescribeUnBlockStatisResponse, err error) {
    if request == nil {
        request = NewDescribeUnBlockStatisRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUnBlockStatis require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUnBlockStatisResponse()
    err = c.Send(request, response)
    return
}

func NewDescribleL4RulesRequest() (request *DescribleL4RulesRequest) {
    request = &DescribleL4RulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribleL4Rules")
    
    
    return
}

func NewDescribleL4RulesResponse() (response *DescribleL4RulesResponse) {
    response = &DescribleL4RulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribleL4Rules
// This API is used to get a layer-4 forwarding rule.
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribleL4Rules(request *DescribleL4RulesRequest) (response *DescribleL4RulesResponse, err error) {
    return c.DescribleL4RulesWithContext(context.Background(), request)
}

// DescribleL4Rules
// This API is used to get a layer-4 forwarding rule.
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribleL4RulesWithContext(ctx context.Context, request *DescribleL4RulesRequest) (response *DescribleL4RulesResponse, err error) {
    if request == nil {
        request = NewDescribleL4RulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribleL4Rules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribleL4RulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribleL7RulesRequest() (request *DescribleL7RulesRequest) {
    request = &DescribleL7RulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribleL7Rules")
    
    
    return
}

func NewDescribleL7RulesResponse() (response *DescribleL7RulesResponse) {
    response = &DescribleL7RulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribleL7Rules
// This API is used to get a layer-7 forwarding rule.
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribleL7Rules(request *DescribleL7RulesRequest) (response *DescribleL7RulesResponse, err error) {
    return c.DescribleL7RulesWithContext(context.Background(), request)
}

// DescribleL7Rules
// This API is used to get a layer-7 forwarding rule.
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribleL7RulesWithContext(ctx context.Context, request *DescribleL7RulesRequest) (response *DescribleL7RulesResponse, err error) {
    if request == nil {
        request = NewDescribleL7RulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribleL7Rules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribleL7RulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribleRegionCountRequest() (request *DescribleRegionCountRequest) {
    request = &DescribleRegionCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "DescribleRegionCount")
    
    
    return
}

func NewDescribleRegionCountResponse() (response *DescribleRegionCountResponse) {
    response = &DescribleRegionCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribleRegionCount
// This API is used to get the number of resource instances in a region.
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribleRegionCount(request *DescribleRegionCountRequest) (response *DescribleRegionCountResponse, err error) {
    return c.DescribleRegionCountWithContext(context.Background(), request)
}

// DescribleRegionCount
// This API is used to get the number of resource instances in a region.
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribleRegionCountWithContext(ctx context.Context, request *DescribleRegionCountRequest) (response *DescribleRegionCountResponse, err error) {
    if request == nil {
        request = NewDescribleRegionCountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribleRegionCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribleRegionCountResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCCAlarmThresholdRequest() (request *ModifyCCAlarmThresholdRequest) {
    request = &ModifyCCAlarmThresholdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyCCAlarmThreshold")
    
    
    return
}

func NewModifyCCAlarmThresholdResponse() (response *ModifyCCAlarmThresholdResponse) {
    response = &ModifyCCAlarmThresholdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCCAlarmThreshold
// This API is used to set the alarm notification threshold for CC attacks in Anti-DDoS Pro, Anti-DDoS Advanced, Anti-DDoS Ultimate, and Chess Shield.
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCCAlarmThreshold(request *ModifyCCAlarmThresholdRequest) (response *ModifyCCAlarmThresholdResponse, err error) {
    return c.ModifyCCAlarmThresholdWithContext(context.Background(), request)
}

// ModifyCCAlarmThreshold
// This API is used to set the alarm notification threshold for CC attacks in Anti-DDoS Pro, Anti-DDoS Advanced, Anti-DDoS Ultimate, and Chess Shield.
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCCAlarmThresholdWithContext(ctx context.Context, request *ModifyCCAlarmThresholdRequest) (response *ModifyCCAlarmThresholdResponse, err error) {
    if request == nil {
        request = NewModifyCCAlarmThresholdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCCAlarmThreshold require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCCAlarmThresholdResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCCFrequencyRulesRequest() (request *ModifyCCFrequencyRulesRequest) {
    request = &ModifyCCFrequencyRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyCCFrequencyRules")
    
    
    return
}

func NewModifyCCFrequencyRulesResponse() (response *ModifyCCFrequencyRulesResponse) {
    response = &ModifyCCFrequencyRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCCFrequencyRules
// This API is used to modify an access frequency control rule for CC protection.
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCCFrequencyRules(request *ModifyCCFrequencyRulesRequest) (response *ModifyCCFrequencyRulesResponse, err error) {
    return c.ModifyCCFrequencyRulesWithContext(context.Background(), request)
}

// ModifyCCFrequencyRules
// This API is used to modify an access frequency control rule for CC protection.
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCCFrequencyRulesWithContext(ctx context.Context, request *ModifyCCFrequencyRulesRequest) (response *ModifyCCFrequencyRulesResponse, err error) {
    if request == nil {
        request = NewModifyCCFrequencyRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCCFrequencyRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCCFrequencyRulesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCCFrequencyRulesStatusRequest() (request *ModifyCCFrequencyRulesStatusRequest) {
    request = &ModifyCCFrequencyRulesStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyCCFrequencyRulesStatus")
    
    
    return
}

func NewModifyCCFrequencyRulesStatusResponse() (response *ModifyCCFrequencyRulesStatusResponse) {
    response = &ModifyCCFrequencyRulesStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCCFrequencyRulesStatus
// This API is used to enable or disable an access frequency control rule for CC protection.
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCCFrequencyRulesStatus(request *ModifyCCFrequencyRulesStatusRequest) (response *ModifyCCFrequencyRulesStatusResponse, err error) {
    return c.ModifyCCFrequencyRulesStatusWithContext(context.Background(), request)
}

// ModifyCCFrequencyRulesStatus
// This API is used to enable or disable an access frequency control rule for CC protection.
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCCFrequencyRulesStatusWithContext(ctx context.Context, request *ModifyCCFrequencyRulesStatusRequest) (response *ModifyCCFrequencyRulesStatusResponse, err error) {
    if request == nil {
        request = NewModifyCCFrequencyRulesStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCCFrequencyRulesStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCCFrequencyRulesStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCCHostProtectionRequest() (request *ModifyCCHostProtectionRequest) {
    request = &ModifyCCHostProtectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyCCHostProtection")
    
    
    return
}

func NewModifyCCHostProtectionResponse() (response *ModifyCCHostProtectionResponse) {
    response = &ModifyCCHostProtectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCCHostProtection
// This API is used to enable or disable CC domain name protection.
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCCHostProtection(request *ModifyCCHostProtectionRequest) (response *ModifyCCHostProtectionResponse, err error) {
    return c.ModifyCCHostProtectionWithContext(context.Background(), request)
}

// ModifyCCHostProtection
// This API is used to enable or disable CC domain name protection.
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCCHostProtectionWithContext(ctx context.Context, request *ModifyCCHostProtectionRequest) (response *ModifyCCHostProtectionResponse, err error) {
    if request == nil {
        request = NewModifyCCHostProtectionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCCHostProtection require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCCHostProtectionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCCIpAllowDenyRequest() (request *ModifyCCIpAllowDenyRequest) {
    request = &ModifyCCIpAllowDenyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyCCIpAllowDeny")
    
    
    return
}

func NewModifyCCIpAllowDenyResponse() (response *ModifyCCIpAllowDenyResponse) {
    response = &ModifyCCIpAllowDenyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCCIpAllowDeny
// This API is used to add/remove a CC IP to/from the blocklist/allowlist.
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCCIpAllowDeny(request *ModifyCCIpAllowDenyRequest) (response *ModifyCCIpAllowDenyResponse, err error) {
    return c.ModifyCCIpAllowDenyWithContext(context.Background(), request)
}

// ModifyCCIpAllowDeny
// This API is used to add/remove a CC IP to/from the blocklist/allowlist.
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCCIpAllowDenyWithContext(ctx context.Context, request *ModifyCCIpAllowDenyRequest) (response *ModifyCCIpAllowDenyResponse, err error) {
    if request == nil {
        request = NewModifyCCIpAllowDenyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCCIpAllowDeny require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCCIpAllowDenyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCCLevelRequest() (request *ModifyCCLevelRequest) {
    request = &ModifyCCLevelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyCCLevel")
    
    
    return
}

func NewModifyCCLevelResponse() (response *ModifyCCLevelResponse) {
    response = &ModifyCCLevelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCCLevel
// This API is used to modify CC protection level.
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCCLevel(request *ModifyCCLevelRequest) (response *ModifyCCLevelResponse, err error) {
    return c.ModifyCCLevelWithContext(context.Background(), request)
}

// ModifyCCLevel
// This API is used to modify CC protection level.
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCCLevelWithContext(ctx context.Context, request *ModifyCCLevelRequest) (response *ModifyCCLevelResponse, err error) {
    if request == nil {
        request = NewModifyCCLevelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCCLevel require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCCLevelResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCCPolicySwitchRequest() (request *ModifyCCPolicySwitchRequest) {
    request = &ModifyCCPolicySwitchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyCCPolicySwitch")
    
    
    return
}

func NewModifyCCPolicySwitchResponse() (response *ModifyCCPolicySwitchResponse) {
    response = &ModifyCCPolicySwitchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCCPolicySwitch
// This API is used to enable or disable a custom CC policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyCCPolicySwitch(request *ModifyCCPolicySwitchRequest) (response *ModifyCCPolicySwitchResponse, err error) {
    return c.ModifyCCPolicySwitchWithContext(context.Background(), request)
}

// ModifyCCPolicySwitch
// This API is used to enable or disable a custom CC policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyCCPolicySwitchWithContext(ctx context.Context, request *ModifyCCPolicySwitchRequest) (response *ModifyCCPolicySwitchResponse, err error) {
    if request == nil {
        request = NewModifyCCPolicySwitchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCCPolicySwitch require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCCPolicySwitchResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCCSelfDefinePolicyRequest() (request *ModifyCCSelfDefinePolicyRequest) {
    request = &ModifyCCSelfDefinePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyCCSelfDefinePolicy")
    
    
    return
}

func NewModifyCCSelfDefinePolicyResponse() (response *ModifyCCSelfDefinePolicyResponse) {
    response = &ModifyCCSelfDefinePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCCSelfDefinePolicy
// This API is used to modify a custom CC policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyCCSelfDefinePolicy(request *ModifyCCSelfDefinePolicyRequest) (response *ModifyCCSelfDefinePolicyResponse, err error) {
    return c.ModifyCCSelfDefinePolicyWithContext(context.Background(), request)
}

// ModifyCCSelfDefinePolicy
// This API is used to modify a custom CC policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyCCSelfDefinePolicyWithContext(ctx context.Context, request *ModifyCCSelfDefinePolicyRequest) (response *ModifyCCSelfDefinePolicyResponse, err error) {
    if request == nil {
        request = NewModifyCCSelfDefinePolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCCSelfDefinePolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCCSelfDefinePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCCThresholdRequest() (request *ModifyCCThresholdRequest) {
    request = &ModifyCCThresholdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyCCThreshold")
    
    
    return
}

func NewModifyCCThresholdResponse() (response *ModifyCCThresholdResponse) {
    response = &ModifyCCThresholdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCCThreshold
// This API is used to modify the CC protection threshold.
//
// error code that may be returned:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyCCThreshold(request *ModifyCCThresholdRequest) (response *ModifyCCThresholdResponse, err error) {
    return c.ModifyCCThresholdWithContext(context.Background(), request)
}

// ModifyCCThreshold
// This API is used to modify the CC protection threshold.
//
// error code that may be returned:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyCCThresholdWithContext(ctx context.Context, request *ModifyCCThresholdRequest) (response *ModifyCCThresholdResponse, err error) {
    if request == nil {
        request = NewModifyCCThresholdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCCThreshold require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCCThresholdResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCCUrlAllowRequest() (request *ModifyCCUrlAllowRequest) {
    request = &ModifyCCUrlAllowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyCCUrlAllow")
    
    
    return
}

func NewModifyCCUrlAllowResponse() (response *ModifyCCUrlAllowResponse) {
    response = &ModifyCCUrlAllowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCCUrlAllow
// This API is used to add/remove a CC URL to/from the allowlist.
//
// error code that may be returned:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyCCUrlAllow(request *ModifyCCUrlAllowRequest) (response *ModifyCCUrlAllowResponse, err error) {
    return c.ModifyCCUrlAllowWithContext(context.Background(), request)
}

// ModifyCCUrlAllow
// This API is used to add/remove a CC URL to/from the allowlist.
//
// error code that may be returned:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyCCUrlAllowWithContext(ctx context.Context, request *ModifyCCUrlAllowRequest) (response *ModifyCCUrlAllowResponse, err error) {
    if request == nil {
        request = NewModifyCCUrlAllowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCCUrlAllow require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCCUrlAllowResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDDoSAIStatusRequest() (request *ModifyDDoSAIStatusRequest) {
    request = &ModifyDDoSAIStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyDDoSAIStatus")
    
    
    return
}

func NewModifyDDoSAIStatusResponse() (response *ModifyDDoSAIStatusResponse) {
    response = &ModifyDDoSAIStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDDoSAIStatus
// This API is used to read or modify DDoS AI protection status.
//
// error code that may be returned:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyDDoSAIStatus(request *ModifyDDoSAIStatusRequest) (response *ModifyDDoSAIStatusResponse, err error) {
    return c.ModifyDDoSAIStatusWithContext(context.Background(), request)
}

// ModifyDDoSAIStatus
// This API is used to read or modify DDoS AI protection status.
//
// error code that may be returned:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyDDoSAIStatusWithContext(ctx context.Context, request *ModifyDDoSAIStatusRequest) (response *ModifyDDoSAIStatusResponse, err error) {
    if request == nil {
        request = NewModifyDDoSAIStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDDoSAIStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDDoSAIStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDDoSAlarmThresholdRequest() (request *ModifyDDoSAlarmThresholdRequest) {
    request = &ModifyDDoSAlarmThresholdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyDDoSAlarmThreshold")
    
    
    return
}

func NewModifyDDoSAlarmThresholdResponse() (response *ModifyDDoSAlarmThresholdResponse) {
    response = &ModifyDDoSAlarmThresholdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDDoSAlarmThreshold
// This API is used to set the alarm notification threshold for DDoS attacks in Anti-DDoS Pro, Anti-DDoS Advanced, Anti-DDoS Ultimate, and Chess Shield.
//
// error code that may be returned:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyDDoSAlarmThreshold(request *ModifyDDoSAlarmThresholdRequest) (response *ModifyDDoSAlarmThresholdResponse, err error) {
    return c.ModifyDDoSAlarmThresholdWithContext(context.Background(), request)
}

// ModifyDDoSAlarmThreshold
// This API is used to set the alarm notification threshold for DDoS attacks in Anti-DDoS Pro, Anti-DDoS Advanced, Anti-DDoS Ultimate, and Chess Shield.
//
// error code that may be returned:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyDDoSAlarmThresholdWithContext(ctx context.Context, request *ModifyDDoSAlarmThresholdRequest) (response *ModifyDDoSAlarmThresholdResponse, err error) {
    if request == nil {
        request = NewModifyDDoSAlarmThresholdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDDoSAlarmThreshold require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDDoSAlarmThresholdResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDDoSDefendStatusRequest() (request *ModifyDDoSDefendStatusRequest) {
    request = &ModifyDDoSDefendStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyDDoSDefendStatus")
    
    
    return
}

func NewModifyDDoSDefendStatusResponse() (response *ModifyDDoSDefendStatusResponse) {
    response = &ModifyDDoSDefendStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDDoSDefendStatus
// This API is used to enable or disable DDoS. It can disable DDoS protection for a period of time, which will be automatically enabled after the period of time elapses.
//
// error code that may be returned:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyDDoSDefendStatus(request *ModifyDDoSDefendStatusRequest) (response *ModifyDDoSDefendStatusResponse, err error) {
    return c.ModifyDDoSDefendStatusWithContext(context.Background(), request)
}

// ModifyDDoSDefendStatus
// This API is used to enable or disable DDoS. It can disable DDoS protection for a period of time, which will be automatically enabled after the period of time elapses.
//
// error code that may be returned:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyDDoSDefendStatusWithContext(ctx context.Context, request *ModifyDDoSDefendStatusRequest) (response *ModifyDDoSDefendStatusResponse, err error) {
    if request == nil {
        request = NewModifyDDoSDefendStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDDoSDefendStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDDoSDefendStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDDoSLevelRequest() (request *ModifyDDoSLevelRequest) {
    request = &ModifyDDoSLevelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyDDoSLevel")
    
    
    return
}

func NewModifyDDoSLevelResponse() (response *ModifyDDoSLevelResponse) {
    response = &ModifyDDoSLevelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDDoSLevel
// This API is used to read or modify DDoS protection level.
//
// error code that may be returned:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyDDoSLevel(request *ModifyDDoSLevelRequest) (response *ModifyDDoSLevelResponse, err error) {
    return c.ModifyDDoSLevelWithContext(context.Background(), request)
}

// ModifyDDoSLevel
// This API is used to read or modify DDoS protection level.
//
// error code that may be returned:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyDDoSLevelWithContext(ctx context.Context, request *ModifyDDoSLevelRequest) (response *ModifyDDoSLevelResponse, err error) {
    if request == nil {
        request = NewModifyDDoSLevelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDDoSLevel require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDDoSLevelResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDDoSPolicyRequest() (request *ModifyDDoSPolicyRequest) {
    request = &ModifyDDoSPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyDDoSPolicy")
    
    
    return
}

func NewModifyDDoSPolicyResponse() (response *ModifyDDoSPolicyResponse) {
    response = &ModifyDDoSPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDDoSPolicy
// This API is used to modify an advanced DDoS policy.
//
// error code that may be returned:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyDDoSPolicy(request *ModifyDDoSPolicyRequest) (response *ModifyDDoSPolicyResponse, err error) {
    return c.ModifyDDoSPolicyWithContext(context.Background(), request)
}

// ModifyDDoSPolicy
// This API is used to modify an advanced DDoS policy.
//
// error code that may be returned:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyDDoSPolicyWithContext(ctx context.Context, request *ModifyDDoSPolicyRequest) (response *ModifyDDoSPolicyResponse, err error) {
    if request == nil {
        request = NewModifyDDoSPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDDoSPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDDoSPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDDoSPolicyCaseRequest() (request *ModifyDDoSPolicyCaseRequest) {
    request = &ModifyDDoSPolicyCaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyDDoSPolicyCase")
    
    
    return
}

func NewModifyDDoSPolicyCaseResponse() (response *ModifyDDoSPolicyCaseResponse) {
    response = &ModifyDDoSPolicyCaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDDoSPolicyCase
// This API is used to modify a policy scenario.
//
// error code that may be returned:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyDDoSPolicyCase(request *ModifyDDoSPolicyCaseRequest) (response *ModifyDDoSPolicyCaseResponse, err error) {
    return c.ModifyDDoSPolicyCaseWithContext(context.Background(), request)
}

// ModifyDDoSPolicyCase
// This API is used to modify a policy scenario.
//
// error code that may be returned:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyDDoSPolicyCaseWithContext(ctx context.Context, request *ModifyDDoSPolicyCaseRequest) (response *ModifyDDoSPolicyCaseResponse, err error) {
    if request == nil {
        request = NewModifyDDoSPolicyCaseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDDoSPolicyCase require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDDoSPolicyCaseResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDDoSPolicyNameRequest() (request *ModifyDDoSPolicyNameRequest) {
    request = &ModifyDDoSPolicyNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyDDoSPolicyName")
    
    
    return
}

func NewModifyDDoSPolicyNameResponse() (response *ModifyDDoSPolicyNameResponse) {
    response = &ModifyDDoSPolicyNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDDoSPolicyName
// This API is used to rename an advanced DDoS policy.
//
// error code that may be returned:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyDDoSPolicyName(request *ModifyDDoSPolicyNameRequest) (response *ModifyDDoSPolicyNameResponse, err error) {
    return c.ModifyDDoSPolicyNameWithContext(context.Background(), request)
}

// ModifyDDoSPolicyName
// This API is used to rename an advanced DDoS policy.
//
// error code that may be returned:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyDDoSPolicyNameWithContext(ctx context.Context, request *ModifyDDoSPolicyNameRequest) (response *ModifyDDoSPolicyNameResponse, err error) {
    if request == nil {
        request = NewModifyDDoSPolicyNameRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDDoSPolicyName require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDDoSPolicyNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDDoSSwitchRequest() (request *ModifyDDoSSwitchRequest) {
    request = &ModifyDDoSSwitchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyDDoSSwitch")
    
    
    return
}

func NewModifyDDoSSwitchResponse() (response *ModifyDDoSSwitchResponse) {
    response = &ModifyDDoSSwitchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDDoSSwitch
// This API is used to enable or disable DDoS protection, which is only supported for Anti-DDoS Basic.
//
// error code that may be returned:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyDDoSSwitch(request *ModifyDDoSSwitchRequest) (response *ModifyDDoSSwitchResponse, err error) {
    return c.ModifyDDoSSwitchWithContext(context.Background(), request)
}

// ModifyDDoSSwitch
// This API is used to enable or disable DDoS protection, which is only supported for Anti-DDoS Basic.
//
// error code that may be returned:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyDDoSSwitchWithContext(ctx context.Context, request *ModifyDDoSSwitchRequest) (response *ModifyDDoSSwitchResponse, err error) {
    if request == nil {
        request = NewModifyDDoSSwitchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDDoSSwitch require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDDoSSwitchResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDDoSThresholdRequest() (request *ModifyDDoSThresholdRequest) {
    request = &ModifyDDoSThresholdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyDDoSThreshold")
    
    
    return
}

func NewModifyDDoSThresholdResponse() (response *ModifyDDoSThresholdResponse) {
    response = &ModifyDDoSThresholdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDDoSThreshold
// This API is used to modify the DDoS cleansing threshold.
//
// error code that may be returned:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyDDoSThreshold(request *ModifyDDoSThresholdRequest) (response *ModifyDDoSThresholdResponse, err error) {
    return c.ModifyDDoSThresholdWithContext(context.Background(), request)
}

// ModifyDDoSThreshold
// This API is used to modify the DDoS cleansing threshold.
//
// error code that may be returned:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyDDoSThresholdWithContext(ctx context.Context, request *ModifyDDoSThresholdRequest) (response *ModifyDDoSThresholdResponse, err error) {
    if request == nil {
        request = NewModifyDDoSThresholdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDDoSThreshold require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDDoSThresholdResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDDoSWaterKeyRequest() (request *ModifyDDoSWaterKeyRequest) {
    request = &ModifyDDoSWaterKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyDDoSWaterKey")
    
    
    return
}

func NewModifyDDoSWaterKeyResponse() (response *ModifyDDoSWaterKeyResponse) {
    response = &ModifyDDoSWaterKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDDoSWaterKey
// This API is used to add, delete, enable, or disable a watermark key.
//
// error code that may be returned:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyDDoSWaterKey(request *ModifyDDoSWaterKeyRequest) (response *ModifyDDoSWaterKeyResponse, err error) {
    return c.ModifyDDoSWaterKeyWithContext(context.Background(), request)
}

// ModifyDDoSWaterKey
// This API is used to add, delete, enable, or disable a watermark key.
//
// error code that may be returned:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyDDoSWaterKeyWithContext(ctx context.Context, request *ModifyDDoSWaterKeyRequest) (response *ModifyDDoSWaterKeyResponse, err error) {
    if request == nil {
        request = NewModifyDDoSWaterKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDDoSWaterKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDDoSWaterKeyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyElasticLimitRequest() (request *ModifyElasticLimitRequest) {
    request = &ModifyElasticLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyElasticLimit")
    
    
    return
}

func NewModifyElasticLimitResponse() (response *ModifyElasticLimitResponse) {
    response = &ModifyElasticLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyElasticLimit
// This API is used to modify the elastic protection threshold.
//
// error code that may be returned:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyElasticLimit(request *ModifyElasticLimitRequest) (response *ModifyElasticLimitResponse, err error) {
    return c.ModifyElasticLimitWithContext(context.Background(), request)
}

// ModifyElasticLimit
// This API is used to modify the elastic protection threshold.
//
// error code that may be returned:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyElasticLimitWithContext(ctx context.Context, request *ModifyElasticLimitRequest) (response *ModifyElasticLimitResponse, err error) {
    if request == nil {
        request = NewModifyElasticLimitRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyElasticLimit require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyElasticLimitResponse()
    err = c.Send(request, response)
    return
}

func NewModifyL4HealthRequest() (request *ModifyL4HealthRequest) {
    request = &ModifyL4HealthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyL4Health")
    
    
    return
}

func NewModifyL4HealthResponse() (response *ModifyL4HealthResponse) {
    response = &ModifyL4HealthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyL4Health
// This API is used to modify the health check parameters of a layer-4 forwarding rule. It is supported for Anti-DDoS Advanced and Anti-DDoS Ultimate.
//
// error code that may be returned:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyL4Health(request *ModifyL4HealthRequest) (response *ModifyL4HealthResponse, err error) {
    return c.ModifyL4HealthWithContext(context.Background(), request)
}

// ModifyL4Health
// This API is used to modify the health check parameters of a layer-4 forwarding rule. It is supported for Anti-DDoS Advanced and Anti-DDoS Ultimate.
//
// error code that may be returned:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyL4HealthWithContext(ctx context.Context, request *ModifyL4HealthRequest) (response *ModifyL4HealthResponse, err error) {
    if request == nil {
        request = NewModifyL4HealthRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyL4Health require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyL4HealthResponse()
    err = c.Send(request, response)
    return
}

func NewModifyL4KeepTimeRequest() (request *ModifyL4KeepTimeRequest) {
    request = &ModifyL4KeepTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyL4KeepTime")
    
    
    return
}

func NewModifyL4KeepTimeResponse() (response *ModifyL4KeepTimeResponse) {
    response = &ModifyL4KeepTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyL4KeepTime
// This API is used to modify the session persistence of a layer-4 forwarding rule. It is supported for Anti-DDoS Advanced and Anti-DDoS Ultimate.
//
// error code that may be returned:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyL4KeepTime(request *ModifyL4KeepTimeRequest) (response *ModifyL4KeepTimeResponse, err error) {
    return c.ModifyL4KeepTimeWithContext(context.Background(), request)
}

// ModifyL4KeepTime
// This API is used to modify the session persistence of a layer-4 forwarding rule. It is supported for Anti-DDoS Advanced and Anti-DDoS Ultimate.
//
// error code that may be returned:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyL4KeepTimeWithContext(ctx context.Context, request *ModifyL4KeepTimeRequest) (response *ModifyL4KeepTimeResponse, err error) {
    if request == nil {
        request = NewModifyL4KeepTimeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyL4KeepTime require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyL4KeepTimeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyL4RulesRequest() (request *ModifyL4RulesRequest) {
    request = &ModifyL4RulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyL4Rules")
    
    
    return
}

func NewModifyL4RulesResponse() (response *ModifyL4RulesResponse) {
    response = &ModifyL4RulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyL4Rules
// This API is used to modify a layer-4 forwarding rule.
//
// error code that may be returned:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyL4Rules(request *ModifyL4RulesRequest) (response *ModifyL4RulesResponse, err error) {
    return c.ModifyL4RulesWithContext(context.Background(), request)
}

// ModifyL4Rules
// This API is used to modify a layer-4 forwarding rule.
//
// error code that may be returned:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyL4RulesWithContext(ctx context.Context, request *ModifyL4RulesRequest) (response *ModifyL4RulesResponse, err error) {
    if request == nil {
        request = NewModifyL4RulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyL4Rules require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyL4RulesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyL7RulesRequest() (request *ModifyL7RulesRequest) {
    request = &ModifyL7RulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyL7Rules")
    
    
    return
}

func NewModifyL7RulesResponse() (response *ModifyL7RulesResponse) {
    response = &ModifyL7RulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyL7Rules
// This API is used to modify the layer-7 forwarding rules.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyL7Rules(request *ModifyL7RulesRequest) (response *ModifyL7RulesResponse, err error) {
    return c.ModifyL7RulesWithContext(context.Background(), request)
}

// ModifyL7Rules
// This API is used to modify the layer-7 forwarding rules.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyL7RulesWithContext(ctx context.Context, request *ModifyL7RulesRequest) (response *ModifyL7RulesResponse, err error) {
    if request == nil {
        request = NewModifyL7RulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyL7Rules require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyL7RulesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNetReturnSwitchRequest() (request *ModifyNetReturnSwitchRequest) {
    request = &ModifyNetReturnSwitchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyNetReturnSwitch")
    
    
    return
}

func NewModifyNetReturnSwitchResponse() (response *ModifyNetReturnSwitchResponse) {
    response = &ModifyNetReturnSwitchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyNetReturnSwitch
// This API is used to switch a client to the real server and set the switch duration when the client is under attack or blocked.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyNetReturnSwitch(request *ModifyNetReturnSwitchRequest) (response *ModifyNetReturnSwitchResponse, err error) {
    return c.ModifyNetReturnSwitchWithContext(context.Background(), request)
}

// ModifyNetReturnSwitch
// This API is used to switch a client to the real server and set the switch duration when the client is under attack or blocked.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyNetReturnSwitchWithContext(ctx context.Context, request *ModifyNetReturnSwitchRequest) (response *ModifyNetReturnSwitchResponse, err error) {
    if request == nil {
        request = NewModifyNetReturnSwitchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNetReturnSwitch require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNetReturnSwitchResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNewDomainRulesRequest() (request *ModifyNewDomainRulesRequest) {
    request = &ModifyNewDomainRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyNewDomainRules")
    
    
    return
}

func NewModifyNewDomainRulesResponse() (response *ModifyNewDomainRulesResponse) {
    response = &ModifyNewDomainRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyNewDomainRules
// This API is used to modify layer-7 forwarding rules.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) ModifyNewDomainRules(request *ModifyNewDomainRulesRequest) (response *ModifyNewDomainRulesResponse, err error) {
    return c.ModifyNewDomainRulesWithContext(context.Background(), request)
}

// ModifyNewDomainRules
// This API is used to modify layer-7 forwarding rules.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) ModifyNewDomainRulesWithContext(ctx context.Context, request *ModifyNewDomainRulesRequest) (response *ModifyNewDomainRulesResponse, err error) {
    if request == nil {
        request = NewModifyNewDomainRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNewDomainRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNewDomainRulesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNewL4RuleRequest() (request *ModifyNewL4RuleRequest) {
    request = &ModifyNewL4RuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyNewL4Rule")
    
    
    return
}

func NewModifyNewL4RuleResponse() (response *ModifyNewL4RuleResponse) {
    response = &ModifyNewL4RuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyNewL4Rule
// This API is used to modify layer-4 forwarding rules.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) ModifyNewL4Rule(request *ModifyNewL4RuleRequest) (response *ModifyNewL4RuleResponse, err error) {
    return c.ModifyNewL4RuleWithContext(context.Background(), request)
}

// ModifyNewL4Rule
// This API is used to modify layer-4 forwarding rules.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) ModifyNewL4RuleWithContext(ctx context.Context, request *ModifyNewL4RuleRequest) (response *ModifyNewL4RuleResponse, err error) {
    if request == nil {
        request = NewModifyNewL4RuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNewL4Rule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNewL4RuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyResBindDDoSPolicyRequest() (request *ModifyResBindDDoSPolicyRequest) {
    request = &ModifyResBindDDoSPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyResBindDDoSPolicy")
    
    
    return
}

func NewModifyResBindDDoSPolicyResponse() (response *ModifyResBindDDoSPolicyResponse) {
    response = &ModifyResBindDDoSPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyResBindDDoSPolicy
// This API is used to bind an advanced DDoS policy to an instance.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) ModifyResBindDDoSPolicy(request *ModifyResBindDDoSPolicyRequest) (response *ModifyResBindDDoSPolicyResponse, err error) {
    return c.ModifyResBindDDoSPolicyWithContext(context.Background(), request)
}

// ModifyResBindDDoSPolicy
// This API is used to bind an advanced DDoS policy to an instance.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) ModifyResBindDDoSPolicyWithContext(ctx context.Context, request *ModifyResBindDDoSPolicyRequest) (response *ModifyResBindDDoSPolicyResponse, err error) {
    if request == nil {
        request = NewModifyResBindDDoSPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyResBindDDoSPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyResBindDDoSPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyResourceRenewFlagRequest() (request *ModifyResourceRenewFlagRequest) {
    request = &ModifyResourceRenewFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyResourceRenewFlag")
    
    
    return
}

func NewModifyResourceRenewFlagResponse() (response *ModifyResourceRenewFlagResponse) {
    response = &ModifyResourceRenewFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyResourceRenewFlag
// This API is used to enable or disable auto-renewal for a resource.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyResourceRenewFlag(request *ModifyResourceRenewFlagRequest) (response *ModifyResourceRenewFlagResponse, err error) {
    return c.ModifyResourceRenewFlagWithContext(context.Background(), request)
}

// ModifyResourceRenewFlag
// This API is used to enable or disable auto-renewal for a resource.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyResourceRenewFlagWithContext(ctx context.Context, request *ModifyResourceRenewFlagRequest) (response *ModifyResourceRenewFlagResponse, err error) {
    if request == nil {
        request = NewModifyResourceRenewFlagRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyResourceRenewFlag require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyResourceRenewFlagResponse()
    err = c.Send(request, response)
    return
}
