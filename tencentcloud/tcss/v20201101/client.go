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

package v20201101

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2020-11-01"

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


func NewAddAndPublishNetworkFirewallPolicyDetailRequest() (request *AddAndPublishNetworkFirewallPolicyDetailRequest) {
    request = &AddAndPublishNetworkFirewallPolicyDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "AddAndPublishNetworkFirewallPolicyDetail")
    
    
    return
}

func NewAddAndPublishNetworkFirewallPolicyDetailResponse() (response *AddAndPublishNetworkFirewallPolicyDetailResponse) {
    response = &AddAndPublishNetworkFirewallPolicyDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddAndPublishNetworkFirewallPolicyDetail
// This API is used to create a task to add and publish a network policy in the container network.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AddAndPublishNetworkFirewallPolicyDetail(request *AddAndPublishNetworkFirewallPolicyDetailRequest) (response *AddAndPublishNetworkFirewallPolicyDetailResponse, err error) {
    return c.AddAndPublishNetworkFirewallPolicyDetailWithContext(context.Background(), request)
}

// AddAndPublishNetworkFirewallPolicyDetail
// This API is used to create a task to add and publish a network policy in the container network.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AddAndPublishNetworkFirewallPolicyDetailWithContext(ctx context.Context, request *AddAndPublishNetworkFirewallPolicyDetailRequest) (response *AddAndPublishNetworkFirewallPolicyDetailResponse, err error) {
    if request == nil {
        request = NewAddAndPublishNetworkFirewallPolicyDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddAndPublishNetworkFirewallPolicyDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddAndPublishNetworkFirewallPolicyDetailResponse()
    err = c.Send(request, response)
    return
}

func NewAddAndPublishNetworkFirewallPolicyYamlDetailRequest() (request *AddAndPublishNetworkFirewallPolicyYamlDetailRequest) {
    request = &AddAndPublishNetworkFirewallPolicyYamlDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "AddAndPublishNetworkFirewallPolicyYamlDetail")
    
    
    return
}

func NewAddAndPublishNetworkFirewallPolicyYamlDetailResponse() (response *AddAndPublishNetworkFirewallPolicyYamlDetailResponse) {
    response = &AddAndPublishNetworkFirewallPolicyYamlDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddAndPublishNetworkFirewallPolicyYamlDetail
// This API is used to create a task to configure and publish a YAML network policy in the container network.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AddAndPublishNetworkFirewallPolicyYamlDetail(request *AddAndPublishNetworkFirewallPolicyYamlDetailRequest) (response *AddAndPublishNetworkFirewallPolicyYamlDetailResponse, err error) {
    return c.AddAndPublishNetworkFirewallPolicyYamlDetailWithContext(context.Background(), request)
}

// AddAndPublishNetworkFirewallPolicyYamlDetail
// This API is used to create a task to configure and publish a YAML network policy in the container network.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AddAndPublishNetworkFirewallPolicyYamlDetailWithContext(ctx context.Context, request *AddAndPublishNetworkFirewallPolicyYamlDetailRequest) (response *AddAndPublishNetworkFirewallPolicyYamlDetailResponse, err error) {
    if request == nil {
        request = NewAddAndPublishNetworkFirewallPolicyYamlDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddAndPublishNetworkFirewallPolicyYamlDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddAndPublishNetworkFirewallPolicyYamlDetailResponse()
    err = c.Send(request, response)
    return
}

func NewAddAssetImageRegistryRegistryDetailRequest() (request *AddAssetImageRegistryRegistryDetailRequest) {
    request = &AddAssetImageRegistryRegistryDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "AddAssetImageRegistryRegistryDetail")
    
    
    return
}

func NewAddAssetImageRegistryRegistryDetailResponse() (response *AddAssetImageRegistryRegistryDetailResponse) {
    response = &AddAssetImageRegistryRegistryDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddAssetImageRegistryRegistryDetail
// This API is used to add the details of an image repository.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AddAssetImageRegistryRegistryDetail(request *AddAssetImageRegistryRegistryDetailRequest) (response *AddAssetImageRegistryRegistryDetailResponse, err error) {
    return c.AddAssetImageRegistryRegistryDetailWithContext(context.Background(), request)
}

// AddAssetImageRegistryRegistryDetail
// This API is used to add the details of an image repository.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AddAssetImageRegistryRegistryDetailWithContext(ctx context.Context, request *AddAssetImageRegistryRegistryDetailRequest) (response *AddAssetImageRegistryRegistryDetailResponse, err error) {
    if request == nil {
        request = NewAddAssetImageRegistryRegistryDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddAssetImageRegistryRegistryDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddAssetImageRegistryRegistryDetailResponse()
    err = c.Send(request, response)
    return
}

func NewAddComplianceAssetPolicySetToWhitelistRequest() (request *AddComplianceAssetPolicySetToWhitelistRequest) {
    request = &AddComplianceAssetPolicySetToWhitelistRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "AddComplianceAssetPolicySetToWhitelist")
    
    
    return
}

func NewAddComplianceAssetPolicySetToWhitelistResponse() (response *AddComplianceAssetPolicySetToWhitelistResponse) {
    response = &AddComplianceAssetPolicySetToWhitelistResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddComplianceAssetPolicySetToWhitelist
// This API is used to ignore the specified asset IDs and check item IDs so as to hide the assets contained in the specified check items.
//
// `AddCompliancePolicyItemToWhitelist` is the reference API. Except for the input field, others should be the same, and if not, it may be due to the definition.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) AddComplianceAssetPolicySetToWhitelist(request *AddComplianceAssetPolicySetToWhitelistRequest) (response *AddComplianceAssetPolicySetToWhitelistResponse, err error) {
    return c.AddComplianceAssetPolicySetToWhitelistWithContext(context.Background(), request)
}

// AddComplianceAssetPolicySetToWhitelist
// This API is used to ignore the specified asset IDs and check item IDs so as to hide the assets contained in the specified check items.
//
// `AddCompliancePolicyItemToWhitelist` is the reference API. Except for the input field, others should be the same, and if not, it may be due to the definition.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) AddComplianceAssetPolicySetToWhitelistWithContext(ctx context.Context, request *AddComplianceAssetPolicySetToWhitelistRequest) (response *AddComplianceAssetPolicySetToWhitelistResponse, err error) {
    if request == nil {
        request = NewAddComplianceAssetPolicySetToWhitelistRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddComplianceAssetPolicySetToWhitelist require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddComplianceAssetPolicySetToWhitelistResponse()
    err = c.Send(request, response)
    return
}

func NewAddCompliancePolicyAssetSetToWhitelistRequest() (request *AddCompliancePolicyAssetSetToWhitelistRequest) {
    request = &AddCompliancePolicyAssetSetToWhitelistRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "AddCompliancePolicyAssetSetToWhitelist")
    
    
    return
}

func NewAddCompliancePolicyAssetSetToWhitelistResponse() (response *AddCompliancePolicyAssetSetToWhitelistResponse) {
    response = &AddCompliancePolicyAssetSetToWhitelistResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddCompliancePolicyAssetSetToWhitelist
// This API is used to ignore the specified asset IDs and check item IDs so as to hide the assets contained in the specified check items.
//
// `AddCompliancePolicyItemToWhitelist` is the reference API. Except for the input field, others should be the same, and if not, it may be due to the definition.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) AddCompliancePolicyAssetSetToWhitelist(request *AddCompliancePolicyAssetSetToWhitelistRequest) (response *AddCompliancePolicyAssetSetToWhitelistResponse, err error) {
    return c.AddCompliancePolicyAssetSetToWhitelistWithContext(context.Background(), request)
}

// AddCompliancePolicyAssetSetToWhitelist
// This API is used to ignore the specified asset IDs and check item IDs so as to hide the assets contained in the specified check items.
//
// `AddCompliancePolicyItemToWhitelist` is the reference API. Except for the input field, others should be the same, and if not, it may be due to the definition.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) AddCompliancePolicyAssetSetToWhitelistWithContext(ctx context.Context, request *AddCompliancePolicyAssetSetToWhitelistRequest) (response *AddCompliancePolicyAssetSetToWhitelistResponse, err error) {
    if request == nil {
        request = NewAddCompliancePolicyAssetSetToWhitelistRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddCompliancePolicyAssetSetToWhitelist require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddCompliancePolicyAssetSetToWhitelistResponse()
    err = c.Send(request, response)
    return
}

func NewAddCompliancePolicyItemToWhitelistRequest() (request *AddCompliancePolicyItemToWhitelistRequest) {
    request = &AddCompliancePolicyItemToWhitelistRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "AddCompliancePolicyItemToWhitelist")
    
    
    return
}

func NewAddCompliancePolicyItemToWhitelistResponse() (response *AddCompliancePolicyItemToWhitelistResponse) {
    response = &AddCompliancePolicyItemToWhitelistResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddCompliancePolicyItemToWhitelist
// This API is used to add the specified check item IDs to the allowlist so as to hide the failure result.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) AddCompliancePolicyItemToWhitelist(request *AddCompliancePolicyItemToWhitelistRequest) (response *AddCompliancePolicyItemToWhitelistResponse, err error) {
    return c.AddCompliancePolicyItemToWhitelistWithContext(context.Background(), request)
}

// AddCompliancePolicyItemToWhitelist
// This API is used to add the specified check item IDs to the allowlist so as to hide the failure result.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) AddCompliancePolicyItemToWhitelistWithContext(ctx context.Context, request *AddCompliancePolicyItemToWhitelistRequest) (response *AddCompliancePolicyItemToWhitelistResponse, err error) {
    if request == nil {
        request = NewAddCompliancePolicyItemToWhitelistRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddCompliancePolicyItemToWhitelist require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddCompliancePolicyItemToWhitelistResponse()
    err = c.Send(request, response)
    return
}

func NewAddEditAbnormalProcessRuleRequest() (request *AddEditAbnormalProcessRuleRequest) {
    request = &AddEditAbnormalProcessRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "AddEditAbnormalProcessRule")
    
    
    return
}

func NewAddEditAbnormalProcessRuleResponse() (response *AddEditAbnormalProcessRuleResponse) {
    response = &AddEditAbnormalProcessRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddEditAbnormalProcessRule
// This API is used to add or edit an abnormal process policy at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTIFYPOLICYCHANGEFAILED = "FailedOperation.NotifyPolicyChangeFailed"
//  FAILEDOPERATION_RULECONFIGTOOMANY = "FailedOperation.RuleConfigTooMany"
//  FAILEDOPERATION_RULEINFOREPEAT = "FailedOperation.RuleInfoRepeat"
//  FAILEDOPERATION_RULENAMEREPEAT = "FailedOperation.RuleNameRepeat"
//  FAILEDOPERATION_RULESELECTIMAGEOUTRANGE = "FailedOperation.RuleSelectImageOutRange"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_RULEINFOINVALID = "InvalidParameter.RuleInfoInValid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AddEditAbnormalProcessRule(request *AddEditAbnormalProcessRuleRequest) (response *AddEditAbnormalProcessRuleResponse, err error) {
    return c.AddEditAbnormalProcessRuleWithContext(context.Background(), request)
}

// AddEditAbnormalProcessRule
// This API is used to add or edit an abnormal process policy at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTIFYPOLICYCHANGEFAILED = "FailedOperation.NotifyPolicyChangeFailed"
//  FAILEDOPERATION_RULECONFIGTOOMANY = "FailedOperation.RuleConfigTooMany"
//  FAILEDOPERATION_RULEINFOREPEAT = "FailedOperation.RuleInfoRepeat"
//  FAILEDOPERATION_RULENAMEREPEAT = "FailedOperation.RuleNameRepeat"
//  FAILEDOPERATION_RULESELECTIMAGEOUTRANGE = "FailedOperation.RuleSelectImageOutRange"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_RULEINFOINVALID = "InvalidParameter.RuleInfoInValid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AddEditAbnormalProcessRuleWithContext(ctx context.Context, request *AddEditAbnormalProcessRuleRequest) (response *AddEditAbnormalProcessRuleResponse, err error) {
    if request == nil {
        request = NewAddEditAbnormalProcessRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddEditAbnormalProcessRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddEditAbnormalProcessRuleResponse()
    err = c.Send(request, response)
    return
}

func NewAddEditAccessControlRuleRequest() (request *AddEditAccessControlRuleRequest) {
    request = &AddEditAccessControlRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "AddEditAccessControlRule")
    
    
    return
}

func NewAddEditAccessControlRuleResponse() (response *AddEditAccessControlRuleResponse) {
    response = &AddEditAccessControlRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddEditAccessControlRule
// This API is used to add or edit an access control policy at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_NOTIFYPOLICYCHANGEFAILED = "FailedOperation.NotifyPolicyChangeFailed"
//  FAILEDOPERATION_RULECONFIGTOOMANY = "FailedOperation.RuleConfigTooMany"
//  FAILEDOPERATION_RULEINFOREPEAT = "FailedOperation.RuleInfoRepeat"
//  FAILEDOPERATION_RULENAMEREPEAT = "FailedOperation.RuleNameRepeat"
//  FAILEDOPERATION_RULESELECTIMAGEOUTRANGE = "FailedOperation.RuleSelectImageOutRange"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_RULEINFOINVALID = "InvalidParameter.RuleInfoInValid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AddEditAccessControlRule(request *AddEditAccessControlRuleRequest) (response *AddEditAccessControlRuleResponse, err error) {
    return c.AddEditAccessControlRuleWithContext(context.Background(), request)
}

// AddEditAccessControlRule
// This API is used to add or edit an access control policy at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_NOTIFYPOLICYCHANGEFAILED = "FailedOperation.NotifyPolicyChangeFailed"
//  FAILEDOPERATION_RULECONFIGTOOMANY = "FailedOperation.RuleConfigTooMany"
//  FAILEDOPERATION_RULEINFOREPEAT = "FailedOperation.RuleInfoRepeat"
//  FAILEDOPERATION_RULENAMEREPEAT = "FailedOperation.RuleNameRepeat"
//  FAILEDOPERATION_RULESELECTIMAGEOUTRANGE = "FailedOperation.RuleSelectImageOutRange"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_RULEINFOINVALID = "InvalidParameter.RuleInfoInValid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AddEditAccessControlRuleWithContext(ctx context.Context, request *AddEditAccessControlRuleRequest) (response *AddEditAccessControlRuleResponse, err error) {
    if request == nil {
        request = NewAddEditAccessControlRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddEditAccessControlRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddEditAccessControlRuleResponse()
    err = c.Send(request, response)
    return
}

func NewAddEditImageAutoAuthorizedRuleRequest() (request *AddEditImageAutoAuthorizedRuleRequest) {
    request = &AddEditImageAutoAuthorizedRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "AddEditImageAutoAuthorizedRule")
    
    
    return
}

func NewAddEditImageAutoAuthorizedRuleResponse() (response *AddEditImageAutoAuthorizedRuleResponse) {
    response = &AddEditImageAutoAuthorizedRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddEditImageAutoAuthorizedRule
// This API is used to add or edit an automatic licensing rule for local images.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRRULENOTFIND = "FailedOperation.ErrRuleNotFind"
//  FAILEDOPERATION_RULENOTFIND = "FailedOperation.RuleNotFind"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) AddEditImageAutoAuthorizedRule(request *AddEditImageAutoAuthorizedRuleRequest) (response *AddEditImageAutoAuthorizedRuleResponse, err error) {
    return c.AddEditImageAutoAuthorizedRuleWithContext(context.Background(), request)
}

// AddEditImageAutoAuthorizedRule
// This API is used to add or edit an automatic licensing rule for local images.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRRULENOTFIND = "FailedOperation.ErrRuleNotFind"
//  FAILEDOPERATION_RULENOTFIND = "FailedOperation.RuleNotFind"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) AddEditImageAutoAuthorizedRuleWithContext(ctx context.Context, request *AddEditImageAutoAuthorizedRuleRequest) (response *AddEditImageAutoAuthorizedRuleResponse, err error) {
    if request == nil {
        request = NewAddEditImageAutoAuthorizedRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddEditImageAutoAuthorizedRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddEditImageAutoAuthorizedRuleResponse()
    err = c.Send(request, response)
    return
}

func NewAddEditReverseShellWhiteListRequest() (request *AddEditReverseShellWhiteListRequest) {
    request = &AddEditReverseShellWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "AddEditReverseShellWhiteList")
    
    
    return
}

func NewAddEditReverseShellWhiteListResponse() (response *AddEditReverseShellWhiteListResponse) {
    response = &AddEditReverseShellWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddEditReverseShellWhiteList
// This API is used to add or edit an allowed reverse shell at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_NOTIFYPOLICYCHANGEFAILED = "FailedOperation.NotifyPolicyChangeFailed"
//  FAILEDOPERATION_RULECONFIGTOOMANY = "FailedOperation.RuleConfigTooMany"
//  FAILEDOPERATION_RULEINFOREPEAT = "FailedOperation.RuleInfoRepeat"
//  FAILEDOPERATION_RULENAMEREPEAT = "FailedOperation.RuleNameRepeat"
//  FAILEDOPERATION_RULESELECTIMAGEOUTRANGE = "FailedOperation.RuleSelectImageOutRange"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRIPNOVALID = "InvalidParameter.ErrIpNoValid"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PORTNOVALID = "InvalidParameter.PortNoValid"
//  INVALIDPARAMETER_REVERSHELLKEYFIELDALLEMPTY = "InvalidParameter.ReverShellKeyFieldAllEmpty"
//  INVALIDPARAMETER_RULEINFOINVALID = "InvalidParameter.RuleInfoInValid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTHLIMIT = "InvalidParameterValue.LengthLimit"
func (c *Client) AddEditReverseShellWhiteList(request *AddEditReverseShellWhiteListRequest) (response *AddEditReverseShellWhiteListResponse, err error) {
    return c.AddEditReverseShellWhiteListWithContext(context.Background(), request)
}

// AddEditReverseShellWhiteList
// This API is used to add or edit an allowed reverse shell at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_NOTIFYPOLICYCHANGEFAILED = "FailedOperation.NotifyPolicyChangeFailed"
//  FAILEDOPERATION_RULECONFIGTOOMANY = "FailedOperation.RuleConfigTooMany"
//  FAILEDOPERATION_RULEINFOREPEAT = "FailedOperation.RuleInfoRepeat"
//  FAILEDOPERATION_RULENAMEREPEAT = "FailedOperation.RuleNameRepeat"
//  FAILEDOPERATION_RULESELECTIMAGEOUTRANGE = "FailedOperation.RuleSelectImageOutRange"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRIPNOVALID = "InvalidParameter.ErrIpNoValid"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PORTNOVALID = "InvalidParameter.PortNoValid"
//  INVALIDPARAMETER_REVERSHELLKEYFIELDALLEMPTY = "InvalidParameter.ReverShellKeyFieldAllEmpty"
//  INVALIDPARAMETER_RULEINFOINVALID = "InvalidParameter.RuleInfoInValid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTHLIMIT = "InvalidParameterValue.LengthLimit"
func (c *Client) AddEditReverseShellWhiteListWithContext(ctx context.Context, request *AddEditReverseShellWhiteListRequest) (response *AddEditReverseShellWhiteListResponse, err error) {
    if request == nil {
        request = NewAddEditReverseShellWhiteListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddEditReverseShellWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddEditReverseShellWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewAddEditRiskSyscallWhiteListRequest() (request *AddEditRiskSyscallWhiteListRequest) {
    request = &AddEditRiskSyscallWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "AddEditRiskSyscallWhiteList")
    
    
    return
}

func NewAddEditRiskSyscallWhiteListResponse() (response *AddEditRiskSyscallWhiteListResponse) {
    response = &AddEditRiskSyscallWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddEditRiskSyscallWhiteList
// This API is used to add or edit an allowed high-risk syscall at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_NOTIFYPOLICYCHANGEFAILED = "FailedOperation.NotifyPolicyChangeFailed"
//  FAILEDOPERATION_RULECONFIGTOOMANY = "FailedOperation.RuleConfigTooMany"
//  FAILEDOPERATION_RULEINFOREPEAT = "FailedOperation.RuleInfoRepeat"
//  FAILEDOPERATION_RULENAMEREPEAT = "FailedOperation.RuleNameRepeat"
//  FAILEDOPERATION_RULESELECTIMAGEOUTRANGE = "FailedOperation.RuleSelectImageOutRange"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_RULEINFOINVALID = "InvalidParameter.RuleInfoInValid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTHLIMIT = "InvalidParameterValue.LengthLimit"
func (c *Client) AddEditRiskSyscallWhiteList(request *AddEditRiskSyscallWhiteListRequest) (response *AddEditRiskSyscallWhiteListResponse, err error) {
    return c.AddEditRiskSyscallWhiteListWithContext(context.Background(), request)
}

// AddEditRiskSyscallWhiteList
// This API is used to add or edit an allowed high-risk syscall at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_NOTIFYPOLICYCHANGEFAILED = "FailedOperation.NotifyPolicyChangeFailed"
//  FAILEDOPERATION_RULECONFIGTOOMANY = "FailedOperation.RuleConfigTooMany"
//  FAILEDOPERATION_RULEINFOREPEAT = "FailedOperation.RuleInfoRepeat"
//  FAILEDOPERATION_RULENAMEREPEAT = "FailedOperation.RuleNameRepeat"
//  FAILEDOPERATION_RULESELECTIMAGEOUTRANGE = "FailedOperation.RuleSelectImageOutRange"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_RULEINFOINVALID = "InvalidParameter.RuleInfoInValid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTHLIMIT = "InvalidParameterValue.LengthLimit"
func (c *Client) AddEditRiskSyscallWhiteListWithContext(ctx context.Context, request *AddEditRiskSyscallWhiteListRequest) (response *AddEditRiskSyscallWhiteListResponse, err error) {
    if request == nil {
        request = NewAddEditRiskSyscallWhiteListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddEditRiskSyscallWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddEditRiskSyscallWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewAddEditWarningRulesRequest() (request *AddEditWarningRulesRequest) {
    request = &AddEditWarningRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "AddEditWarningRules")
    
    
    return
}

func NewAddEditWarningRulesResponse() (response *AddEditWarningRulesResponse) {
    response = &AddEditWarningRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddEditWarningRules
// This API is used to add or edit an alert policy.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) AddEditWarningRules(request *AddEditWarningRulesRequest) (response *AddEditWarningRulesResponse, err error) {
    return c.AddEditWarningRulesWithContext(context.Background(), request)
}

// AddEditWarningRules
// This API is used to add or edit an alert policy.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) AddEditWarningRulesWithContext(ctx context.Context, request *AddEditWarningRulesRequest) (response *AddEditWarningRulesResponse, err error) {
    if request == nil {
        request = NewAddEditWarningRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddEditWarningRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddEditWarningRulesResponse()
    err = c.Send(request, response)
    return
}

func NewAddEscapeWhiteListRequest() (request *AddEscapeWhiteListRequest) {
    request = &AddEscapeWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "AddEscapeWhiteList")
    
    
    return
}

func NewAddEscapeWhiteListResponse() (response *AddEscapeWhiteListResponse) {
    response = &AddEscapeWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddEscapeWhiteList
// This API is used to add an allowed escape.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) AddEscapeWhiteList(request *AddEscapeWhiteListRequest) (response *AddEscapeWhiteListResponse, err error) {
    return c.AddEscapeWhiteListWithContext(context.Background(), request)
}

// AddEscapeWhiteList
// This API is used to add an allowed escape.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) AddEscapeWhiteListWithContext(ctx context.Context, request *AddEscapeWhiteListRequest) (response *AddEscapeWhiteListResponse, err error) {
    if request == nil {
        request = NewAddEscapeWhiteListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddEscapeWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddEscapeWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewAddIgnoreVulRequest() (request *AddIgnoreVulRequest) {
    request = &AddIgnoreVulRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "AddIgnoreVul")
    
    
    return
}

func NewAddIgnoreVulResponse() (response *AddIgnoreVulResponse) {
    response = &AddIgnoreVulResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddIgnoreVul
// This API is used to add ignored vulnerabilities in a vulnerability scan.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AddIgnoreVul(request *AddIgnoreVulRequest) (response *AddIgnoreVulResponse, err error) {
    return c.AddIgnoreVulWithContext(context.Background(), request)
}

// AddIgnoreVul
// This API is used to add ignored vulnerabilities in a vulnerability scan.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AddIgnoreVulWithContext(ctx context.Context, request *AddIgnoreVulRequest) (response *AddIgnoreVulResponse, err error) {
    if request == nil {
        request = NewAddIgnoreVulRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddIgnoreVul require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddIgnoreVulResponse()
    err = c.Send(request, response)
    return
}

func NewAddNetworkFirewallPolicyDetailRequest() (request *AddNetworkFirewallPolicyDetailRequest) {
    request = &AddNetworkFirewallPolicyDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "AddNetworkFirewallPolicyDetail")
    
    
    return
}

func NewAddNetworkFirewallPolicyDetailResponse() (response *AddNetworkFirewallPolicyDetailResponse) {
    response = &AddNetworkFirewallPolicyDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddNetworkFirewallPolicyDetail
// This API is used to create a task to add a network policy in the container network.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AddNetworkFirewallPolicyDetail(request *AddNetworkFirewallPolicyDetailRequest) (response *AddNetworkFirewallPolicyDetailResponse, err error) {
    return c.AddNetworkFirewallPolicyDetailWithContext(context.Background(), request)
}

// AddNetworkFirewallPolicyDetail
// This API is used to create a task to add a network policy in the container network.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AddNetworkFirewallPolicyDetailWithContext(ctx context.Context, request *AddNetworkFirewallPolicyDetailRequest) (response *AddNetworkFirewallPolicyDetailResponse, err error) {
    if request == nil {
        request = NewAddNetworkFirewallPolicyDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddNetworkFirewallPolicyDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddNetworkFirewallPolicyDetailResponse()
    err = c.Send(request, response)
    return
}

func NewAddNetworkFirewallPolicyYamlDetailRequest() (request *AddNetworkFirewallPolicyYamlDetailRequest) {
    request = &AddNetworkFirewallPolicyYamlDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "AddNetworkFirewallPolicyYamlDetail")
    
    
    return
}

func NewAddNetworkFirewallPolicyYamlDetailResponse() (response *AddNetworkFirewallPolicyYamlDetailResponse) {
    response = &AddNetworkFirewallPolicyYamlDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddNetworkFirewallPolicyYamlDetail
// This API is used to create a task to add a YAML network policy in the container network.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AddNetworkFirewallPolicyYamlDetail(request *AddNetworkFirewallPolicyYamlDetailRequest) (response *AddNetworkFirewallPolicyYamlDetailResponse, err error) {
    return c.AddNetworkFirewallPolicyYamlDetailWithContext(context.Background(), request)
}

// AddNetworkFirewallPolicyYamlDetail
// This API is used to create a task to add a YAML network policy in the container network.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AddNetworkFirewallPolicyYamlDetailWithContext(ctx context.Context, request *AddNetworkFirewallPolicyYamlDetailRequest) (response *AddNetworkFirewallPolicyYamlDetailResponse, err error) {
    if request == nil {
        request = NewAddNetworkFirewallPolicyYamlDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddNetworkFirewallPolicyYamlDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddNetworkFirewallPolicyYamlDetailResponse()
    err = c.Send(request, response)
    return
}

func NewCheckNetworkFirewallPolicyYamlRequest() (request *CheckNetworkFirewallPolicyYamlRequest) {
    request = &CheckNetworkFirewallPolicyYamlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "CheckNetworkFirewallPolicyYaml")
    
    
    return
}

func NewCheckNetworkFirewallPolicyYamlResponse() (response *CheckNetworkFirewallPolicyYamlResponse) {
    response = &CheckNetworkFirewallPolicyYamlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckNetworkFirewallPolicyYaml
// This API is used to create a task to check a YAML network policy in the container network.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CheckNetworkFirewallPolicyYaml(request *CheckNetworkFirewallPolicyYamlRequest) (response *CheckNetworkFirewallPolicyYamlResponse, err error) {
    return c.CheckNetworkFirewallPolicyYamlWithContext(context.Background(), request)
}

// CheckNetworkFirewallPolicyYaml
// This API is used to create a task to check a YAML network policy in the container network.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CheckNetworkFirewallPolicyYamlWithContext(ctx context.Context, request *CheckNetworkFirewallPolicyYamlRequest) (response *CheckNetworkFirewallPolicyYamlResponse, err error) {
    if request == nil {
        request = NewCheckNetworkFirewallPolicyYamlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckNetworkFirewallPolicyYaml require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckNetworkFirewallPolicyYamlResponse()
    err = c.Send(request, response)
    return
}

func NewCheckRepeatAssetImageRegistryRequest() (request *CheckRepeatAssetImageRegistryRequest) {
    request = &CheckRepeatAssetImageRegistryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "CheckRepeatAssetImageRegistry")
    
    
    return
}

func NewCheckRepeatAssetImageRegistryResponse() (response *CheckRepeatAssetImageRegistryResponse) {
    response = &CheckRepeatAssetImageRegistryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckRepeatAssetImageRegistry
// This API is used to check whether an image repository name is duplicated.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CheckRepeatAssetImageRegistry(request *CheckRepeatAssetImageRegistryRequest) (response *CheckRepeatAssetImageRegistryResponse, err error) {
    return c.CheckRepeatAssetImageRegistryWithContext(context.Background(), request)
}

// CheckRepeatAssetImageRegistry
// This API is used to check whether an image repository name is duplicated.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CheckRepeatAssetImageRegistryWithContext(ctx context.Context, request *CheckRepeatAssetImageRegistryRequest) (response *CheckRepeatAssetImageRegistryResponse, err error) {
    if request == nil {
        request = NewCheckRepeatAssetImageRegistryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckRepeatAssetImageRegistry require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckRepeatAssetImageRegistryResponse()
    err = c.Send(request, response)
    return
}

func NewConfirmNetworkFirewallPolicyRequest() (request *ConfirmNetworkFirewallPolicyRequest) {
    request = &ConfirmNetworkFirewallPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "ConfirmNetworkFirewallPolicy")
    
    
    return
}

func NewConfirmNetworkFirewallPolicyResponse() (response *ConfirmNetworkFirewallPolicyResponse) {
    response = &ConfirmNetworkFirewallPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ConfirmNetworkFirewallPolicy
// This API is used to create a task to confirm a network policy in the container network.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ConfirmNetworkFirewallPolicy(request *ConfirmNetworkFirewallPolicyRequest) (response *ConfirmNetworkFirewallPolicyResponse, err error) {
    return c.ConfirmNetworkFirewallPolicyWithContext(context.Background(), request)
}

// ConfirmNetworkFirewallPolicy
// This API is used to create a task to confirm a network policy in the container network.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ConfirmNetworkFirewallPolicyWithContext(ctx context.Context, request *ConfirmNetworkFirewallPolicyRequest) (response *ConfirmNetworkFirewallPolicyResponse, err error) {
    if request == nil {
        request = NewConfirmNetworkFirewallPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ConfirmNetworkFirewallPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewConfirmNetworkFirewallPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAbnormalProcessRulesExportJobRequest() (request *CreateAbnormalProcessRulesExportJobRequest) {
    request = &CreateAbnormalProcessRulesExportJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "CreateAbnormalProcessRulesExportJob")
    
    
    return
}

func NewCreateAbnormalProcessRulesExportJobResponse() (response *CreateAbnormalProcessRulesExportJobResponse) {
    response = &CreateAbnormalProcessRulesExportJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAbnormalProcessRulesExportJob
// This API is used to export abnormal process rules.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateAbnormalProcessRulesExportJob(request *CreateAbnormalProcessRulesExportJobRequest) (response *CreateAbnormalProcessRulesExportJobResponse, err error) {
    return c.CreateAbnormalProcessRulesExportJobWithContext(context.Background(), request)
}

// CreateAbnormalProcessRulesExportJob
// This API is used to export abnormal process rules.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateAbnormalProcessRulesExportJobWithContext(ctx context.Context, request *CreateAbnormalProcessRulesExportJobRequest) (response *CreateAbnormalProcessRulesExportJobResponse, err error) {
    if request == nil {
        request = NewCreateAbnormalProcessRulesExportJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAbnormalProcessRulesExportJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAbnormalProcessRulesExportJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAccessControlsRuleExportJobRequest() (request *CreateAccessControlsRuleExportJobRequest) {
    request = &CreateAccessControlsRuleExportJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "CreateAccessControlsRuleExportJob")
    
    
    return
}

func NewCreateAccessControlsRuleExportJobResponse() (response *CreateAccessControlsRuleExportJobResponse) {
    response = &CreateAccessControlsRuleExportJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAccessControlsRuleExportJob
// This API is used to export file tampering detection rules.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateAccessControlsRuleExportJob(request *CreateAccessControlsRuleExportJobRequest) (response *CreateAccessControlsRuleExportJobResponse, err error) {
    return c.CreateAccessControlsRuleExportJobWithContext(context.Background(), request)
}

// CreateAccessControlsRuleExportJob
// This API is used to export file tampering detection rules.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateAccessControlsRuleExportJobWithContext(ctx context.Context, request *CreateAccessControlsRuleExportJobRequest) (response *CreateAccessControlsRuleExportJobResponse, err error) {
    if request == nil {
        request = NewCreateAccessControlsRuleExportJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAccessControlsRuleExportJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAccessControlsRuleExportJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAssetImageRegistryScanTaskRequest() (request *CreateAssetImageRegistryScanTaskRequest) {
    request = &CreateAssetImageRegistryScanTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "CreateAssetImageRegistryScanTask")
    
    
    return
}

func NewCreateAssetImageRegistryScanTaskResponse() (response *CreateAssetImageRegistryScanTaskResponse) {
    response = &CreateAssetImageRegistryScanTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAssetImageRegistryScanTask
// This API is used to create an image scan task for an image repository.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateAssetImageRegistryScanTask(request *CreateAssetImageRegistryScanTaskRequest) (response *CreateAssetImageRegistryScanTaskResponse, err error) {
    return c.CreateAssetImageRegistryScanTaskWithContext(context.Background(), request)
}

// CreateAssetImageRegistryScanTask
// This API is used to create an image scan task for an image repository.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateAssetImageRegistryScanTaskWithContext(ctx context.Context, request *CreateAssetImageRegistryScanTaskRequest) (response *CreateAssetImageRegistryScanTaskResponse, err error) {
    if request == nil {
        request = NewCreateAssetImageRegistryScanTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAssetImageRegistryScanTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAssetImageRegistryScanTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAssetImageRegistryScanTaskOneKeyRequest() (request *CreateAssetImageRegistryScanTaskOneKeyRequest) {
    request = &CreateAssetImageRegistryScanTaskOneKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "CreateAssetImageRegistryScanTaskOneKey")
    
    
    return
}

func NewCreateAssetImageRegistryScanTaskOneKeyResponse() (response *CreateAssetImageRegistryScanTaskOneKeyResponse) {
    response = &CreateAssetImageRegistryScanTaskOneKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAssetImageRegistryScanTaskOneKey
// This API is used to create a quick image scan task for an image repository.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateAssetImageRegistryScanTaskOneKey(request *CreateAssetImageRegistryScanTaskOneKeyRequest) (response *CreateAssetImageRegistryScanTaskOneKeyResponse, err error) {
    return c.CreateAssetImageRegistryScanTaskOneKeyWithContext(context.Background(), request)
}

// CreateAssetImageRegistryScanTaskOneKey
// This API is used to create a quick image scan task for an image repository.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateAssetImageRegistryScanTaskOneKeyWithContext(ctx context.Context, request *CreateAssetImageRegistryScanTaskOneKeyRequest) (response *CreateAssetImageRegistryScanTaskOneKeyResponse, err error) {
    if request == nil {
        request = NewCreateAssetImageRegistryScanTaskOneKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAssetImageRegistryScanTaskOneKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAssetImageRegistryScanTaskOneKeyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAssetImageScanSettingRequest() (request *CreateAssetImageScanSettingRequest) {
    request = &CreateAssetImageScanSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "CreateAssetImageScanSetting")
    
    
    return
}

func NewCreateAssetImageScanSettingResponse() (response *CreateAssetImageScanSettingResponse) {
    response = &CreateAssetImageScanSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAssetImageScanSetting
// This API is used to set an image scan.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateAssetImageScanSetting(request *CreateAssetImageScanSettingRequest) (response *CreateAssetImageScanSettingResponse, err error) {
    return c.CreateAssetImageScanSettingWithContext(context.Background(), request)
}

// CreateAssetImageScanSetting
// This API is used to set an image scan.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateAssetImageScanSettingWithContext(ctx context.Context, request *CreateAssetImageScanSettingRequest) (response *CreateAssetImageScanSettingResponse, err error) {
    if request == nil {
        request = NewCreateAssetImageScanSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAssetImageScanSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAssetImageScanSettingResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAssetImageScanTaskRequest() (request *CreateAssetImageScanTaskRequest) {
    request = &CreateAssetImageScanTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "CreateAssetImageScanTask")
    
    
    return
}

func NewCreateAssetImageScanTaskResponse() (response *CreateAssetImageScanTaskResponse) {
    response = &CreateAssetImageScanTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAssetImageScanTask
// This API is used to create an image scan task.
//
// error code that may be returned:
//  FAILEDOPERATION_ERRALREADYSCANNING = "FailedOperation.ErrAlreadyScanning"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateAssetImageScanTask(request *CreateAssetImageScanTaskRequest) (response *CreateAssetImageScanTaskResponse, err error) {
    return c.CreateAssetImageScanTaskWithContext(context.Background(), request)
}

// CreateAssetImageScanTask
// This API is used to create an image scan task.
//
// error code that may be returned:
//  FAILEDOPERATION_ERRALREADYSCANNING = "FailedOperation.ErrAlreadyScanning"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateAssetImageScanTaskWithContext(ctx context.Context, request *CreateAssetImageScanTaskRequest) (response *CreateAssetImageScanTaskResponse, err error) {
    if request == nil {
        request = NewCreateAssetImageScanTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAssetImageScanTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAssetImageScanTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAssetImageVirusExportJobRequest() (request *CreateAssetImageVirusExportJobRequest) {
    request = &CreateAssetImageVirusExportJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "CreateAssetImageVirusExportJob")
    
    
    return
}

func NewCreateAssetImageVirusExportJobResponse() (response *CreateAssetImageVirusExportJobResponse) {
    response = &CreateAssetImageVirusExportJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAssetImageVirusExportJob
// This API is used to create a task to export the list of trojans in a local image.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateAssetImageVirusExportJob(request *CreateAssetImageVirusExportJobRequest) (response *CreateAssetImageVirusExportJobResponse, err error) {
    return c.CreateAssetImageVirusExportJobWithContext(context.Background(), request)
}

// CreateAssetImageVirusExportJob
// This API is used to create a task to export the list of trojans in a local image.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateAssetImageVirusExportJobWithContext(ctx context.Context, request *CreateAssetImageVirusExportJobRequest) (response *CreateAssetImageVirusExportJobResponse, err error) {
    if request == nil {
        request = NewCreateAssetImageVirusExportJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAssetImageVirusExportJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAssetImageVirusExportJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCheckComponentRequest() (request *CreateCheckComponentRequest) {
    request = &CreateCheckComponentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "CreateCheckComponent")
    
    
    return
}

func NewCreateCheckComponentResponse() (response *CreateCheckComponentResponse) {
    response = &CreateCheckComponentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCheckComponent
// This API is used to install the check component and create a defender.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateCheckComponent(request *CreateCheckComponentRequest) (response *CreateCheckComponentResponse, err error) {
    return c.CreateCheckComponentWithContext(context.Background(), request)
}

// CreateCheckComponent
// This API is used to install the check component and create a defender.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateCheckComponentWithContext(ctx context.Context, request *CreateCheckComponentRequest) (response *CreateCheckComponentResponse, err error) {
    if request == nil {
        request = NewCreateCheckComponentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCheckComponent require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCheckComponentResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClusterCheckTaskRequest() (request *CreateClusterCheckTaskRequest) {
    request = &CreateClusterCheckTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "CreateClusterCheckTask")
    
    
    return
}

func NewCreateClusterCheckTaskResponse() (response *CreateClusterCheckTaskResponse) {
    response = &CreateClusterCheckTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateClusterCheckTask
// This API is used to create a cluster check task to check it for risk items.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateClusterCheckTask(request *CreateClusterCheckTaskRequest) (response *CreateClusterCheckTaskResponse, err error) {
    return c.CreateClusterCheckTaskWithContext(context.Background(), request)
}

// CreateClusterCheckTask
// This API is used to create a cluster check task to check it for risk items.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateClusterCheckTaskWithContext(ctx context.Context, request *CreateClusterCheckTaskRequest) (response *CreateClusterCheckTaskResponse, err error) {
    if request == nil {
        request = NewCreateClusterCheckTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateClusterCheckTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateClusterCheckTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateComplianceTaskRequest() (request *CreateComplianceTaskRequest) {
    request = &CreateComplianceTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "CreateComplianceTask")
    
    
    return
}

func NewCreateComplianceTaskResponse() (response *CreateComplianceTaskResponse) {
    response = &CreateComplianceTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateComplianceTask
// This API is used to create a compliance check task for another check triggered at the asset level.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateComplianceTask(request *CreateComplianceTaskRequest) (response *CreateComplianceTaskResponse, err error) {
    return c.CreateComplianceTaskWithContext(context.Background(), request)
}

// CreateComplianceTask
// This API is used to create a compliance check task for another check triggered at the asset level.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateComplianceTaskWithContext(ctx context.Context, request *CreateComplianceTaskRequest) (response *CreateComplianceTaskResponse, err error) {
    if request == nil {
        request = NewCreateComplianceTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateComplianceTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateComplianceTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateComponentExportJobRequest() (request *CreateComponentExportJobRequest) {
    request = &CreateComponentExportJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "CreateComponentExportJob")
    
    
    return
}

func NewCreateComponentExportJobResponse() (response *CreateComponentExportJobResponse) {
    response = &CreateComponentExportJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateComponentExportJob
// This API is used to export the list of components in a local image.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateComponentExportJob(request *CreateComponentExportJobRequest) (response *CreateComponentExportJobResponse, err error) {
    return c.CreateComponentExportJobWithContext(context.Background(), request)
}

// CreateComponentExportJob
// This API is used to export the list of components in a local image.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateComponentExportJobWithContext(ctx context.Context, request *CreateComponentExportJobRequest) (response *CreateComponentExportJobResponse, err error) {
    if request == nil {
        request = NewCreateComponentExportJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateComponentExportJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateComponentExportJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDefenceVulExportJobRequest() (request *CreateDefenceVulExportJobRequest) {
    request = &CreateDefenceVulExportJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "CreateDefenceVulExportJob")
    
    
    return
}

func NewCreateDefenceVulExportJobResponse() (response *CreateDefenceVulExportJobResponse) {
    response = &CreateDefenceVulExportJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDefenceVulExportJob
// This API is used to create a task to export vulnerabilities that can be prevented.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateDefenceVulExportJob(request *CreateDefenceVulExportJobRequest) (response *CreateDefenceVulExportJobResponse, err error) {
    return c.CreateDefenceVulExportJobWithContext(context.Background(), request)
}

// CreateDefenceVulExportJob
// This API is used to create a task to export vulnerabilities that can be prevented.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateDefenceVulExportJobWithContext(ctx context.Context, request *CreateDefenceVulExportJobRequest) (response *CreateDefenceVulExportJobResponse, err error) {
    if request == nil {
        request = NewCreateDefenceVulExportJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDefenceVulExportJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDefenceVulExportJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEmergencyVulExportJobRequest() (request *CreateEmergencyVulExportJobRequest) {
    request = &CreateEmergencyVulExportJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "CreateEmergencyVulExportJob")
    
    
    return
}

func NewCreateEmergencyVulExportJobResponse() (response *CreateEmergencyVulExportJobResponse) {
    response = &CreateEmergencyVulExportJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateEmergencyVulExportJob
// This API is used to create a task to export emergency vulnerabilities.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateEmergencyVulExportJob(request *CreateEmergencyVulExportJobRequest) (response *CreateEmergencyVulExportJobResponse, err error) {
    return c.CreateEmergencyVulExportJobWithContext(context.Background(), request)
}

// CreateEmergencyVulExportJob
// This API is used to create a task to export emergency vulnerabilities.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateEmergencyVulExportJobWithContext(ctx context.Context, request *CreateEmergencyVulExportJobRequest) (response *CreateEmergencyVulExportJobResponse, err error) {
    if request == nil {
        request = NewCreateEmergencyVulExportJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEmergencyVulExportJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEmergencyVulExportJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEscapeEventsExportJobRequest() (request *CreateEscapeEventsExportJobRequest) {
    request = &CreateEscapeEventsExportJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "CreateEscapeEventsExportJob")
    
    
    return
}

func NewCreateEscapeEventsExportJobResponse() (response *CreateEscapeEventsExportJobResponse) {
    response = &CreateEscapeEventsExportJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateEscapeEventsExportJob
// This API is used to create a task to asynchronously export escape events.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateEscapeEventsExportJob(request *CreateEscapeEventsExportJobRequest) (response *CreateEscapeEventsExportJobResponse, err error) {
    return c.CreateEscapeEventsExportJobWithContext(context.Background(), request)
}

// CreateEscapeEventsExportJob
// This API is used to create a task to asynchronously export escape events.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateEscapeEventsExportJobWithContext(ctx context.Context, request *CreateEscapeEventsExportJobRequest) (response *CreateEscapeEventsExportJobResponse, err error) {
    if request == nil {
        request = NewCreateEscapeEventsExportJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEscapeEventsExportJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEscapeEventsExportJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEscapeWhiteListExportJobRequest() (request *CreateEscapeWhiteListExportJobRequest) {
    request = &CreateEscapeWhiteListExportJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "CreateEscapeWhiteListExportJob")
    
    
    return
}

func NewCreateEscapeWhiteListExportJobResponse() (response *CreateEscapeWhiteListExportJobResponse) {
    response = &CreateEscapeWhiteListExportJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateEscapeWhiteListExportJob
// This API is used to create a task to export the allowlist of escapes.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateEscapeWhiteListExportJob(request *CreateEscapeWhiteListExportJobRequest) (response *CreateEscapeWhiteListExportJobResponse, err error) {
    return c.CreateEscapeWhiteListExportJobWithContext(context.Background(), request)
}

// CreateEscapeWhiteListExportJob
// This API is used to create a task to export the allowlist of escapes.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateEscapeWhiteListExportJobWithContext(ctx context.Context, request *CreateEscapeWhiteListExportJobRequest) (response *CreateEscapeWhiteListExportJobResponse, err error) {
    if request == nil {
        request = NewCreateEscapeWhiteListExportJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEscapeWhiteListExportJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEscapeWhiteListExportJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateExportComplianceStatusListJobRequest() (request *CreateExportComplianceStatusListJobRequest) {
    request = &CreateExportComplianceStatusListJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "CreateExportComplianceStatusListJob")
    
    
    return
}

func NewCreateExportComplianceStatusListJobResponse() (response *CreateExportComplianceStatusListJobResponse) {
    response = &CreateExportComplianceStatusListJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateExportComplianceStatusListJob
// This API is used to create a task to export security compliance information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATANOTFOUND = "InvalidParameterValue.DataNotFound"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) CreateExportComplianceStatusListJob(request *CreateExportComplianceStatusListJobRequest) (response *CreateExportComplianceStatusListJobResponse, err error) {
    return c.CreateExportComplianceStatusListJobWithContext(context.Background(), request)
}

// CreateExportComplianceStatusListJob
// This API is used to create a task to export security compliance information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATANOTFOUND = "InvalidParameterValue.DataNotFound"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) CreateExportComplianceStatusListJobWithContext(ctx context.Context, request *CreateExportComplianceStatusListJobRequest) (response *CreateExportComplianceStatusListJobResponse, err error) {
    if request == nil {
        request = NewCreateExportComplianceStatusListJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateExportComplianceStatusListJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateExportComplianceStatusListJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateHostExportJobRequest() (request *CreateHostExportJobRequest) {
    request = &CreateHostExportJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "CreateHostExportJob")
    
    
    return
}

func NewCreateHostExportJobResponse() (response *CreateHostExportJobResponse) {
    response = &CreateHostExportJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateHostExportJob
// This API is used to create a task to export the list of servers.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateHostExportJob(request *CreateHostExportJobRequest) (response *CreateHostExportJobResponse, err error) {
    return c.CreateHostExportJobWithContext(context.Background(), request)
}

// CreateHostExportJob
// This API is used to create a task to export the list of servers.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateHostExportJobWithContext(ctx context.Context, request *CreateHostExportJobRequest) (response *CreateHostExportJobResponse, err error) {
    if request == nil {
        request = NewCreateHostExportJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateHostExportJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateHostExportJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateImageExportJobRequest() (request *CreateImageExportJobRequest) {
    request = &CreateImageExportJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "CreateImageExportJob")
    
    
    return
}

func NewCreateImageExportJobResponse() (response *CreateImageExportJobResponse) {
    response = &CreateImageExportJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateImageExportJob
// This API is used to create an image export task.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateImageExportJob(request *CreateImageExportJobRequest) (response *CreateImageExportJobResponse, err error) {
    return c.CreateImageExportJobWithContext(context.Background(), request)
}

// CreateImageExportJob
// This API is used to create an image export task.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateImageExportJobWithContext(ctx context.Context, request *CreateImageExportJobRequest) (response *CreateImageExportJobResponse, err error) {
    if request == nil {
        request = NewCreateImageExportJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateImageExportJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateImageExportJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateK8sApiAbnormalEventExportJobRequest() (request *CreateK8sApiAbnormalEventExportJobRequest) {
    request = &CreateK8sApiAbnormalEventExportJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "CreateK8sApiAbnormalEventExportJob")
    
    
    return
}

func NewCreateK8sApiAbnormalEventExportJobResponse() (response *CreateK8sApiAbnormalEventExportJobResponse) {
    response = &CreateK8sApiAbnormalEventExportJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateK8sApiAbnormalEventExportJob
// This API is used to create K8sApi abnormal event export jobs.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateK8sApiAbnormalEventExportJob(request *CreateK8sApiAbnormalEventExportJobRequest) (response *CreateK8sApiAbnormalEventExportJobResponse, err error) {
    return c.CreateK8sApiAbnormalEventExportJobWithContext(context.Background(), request)
}

// CreateK8sApiAbnormalEventExportJob
// This API is used to create K8sApi abnormal event export jobs.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateK8sApiAbnormalEventExportJobWithContext(ctx context.Context, request *CreateK8sApiAbnormalEventExportJobRequest) (response *CreateK8sApiAbnormalEventExportJobResponse, err error) {
    if request == nil {
        request = NewCreateK8sApiAbnormalEventExportJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateK8sApiAbnormalEventExportJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateK8sApiAbnormalEventExportJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateK8sApiAbnormalRuleExportJobRequest() (request *CreateK8sApiAbnormalRuleExportJobRequest) {
    request = &CreateK8sApiAbnormalRuleExportJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "CreateK8sApiAbnormalRuleExportJob")
    
    
    return
}

func NewCreateK8sApiAbnormalRuleExportJobResponse() (response *CreateK8sApiAbnormalRuleExportJobResponse) {
    response = &CreateK8sApiAbnormalRuleExportJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateK8sApiAbnormalRuleExportJob
// This API is used to export rules of K8sApi exceptions. 
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateK8sApiAbnormalRuleExportJob(request *CreateK8sApiAbnormalRuleExportJobRequest) (response *CreateK8sApiAbnormalRuleExportJobResponse, err error) {
    return c.CreateK8sApiAbnormalRuleExportJobWithContext(context.Background(), request)
}

// CreateK8sApiAbnormalRuleExportJob
// This API is used to export rules of K8sApi exceptions. 
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateK8sApiAbnormalRuleExportJobWithContext(ctx context.Context, request *CreateK8sApiAbnormalRuleExportJobRequest) (response *CreateK8sApiAbnormalRuleExportJobResponse, err error) {
    if request == nil {
        request = NewCreateK8sApiAbnormalRuleExportJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateK8sApiAbnormalRuleExportJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateK8sApiAbnormalRuleExportJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateK8sApiAbnormalRuleInfoRequest() (request *CreateK8sApiAbnormalRuleInfoRequest) {
    request = &CreateK8sApiAbnormalRuleInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "CreateK8sApiAbnormalRuleInfo")
    
    
    return
}

func NewCreateK8sApiAbnormalRuleInfoResponse() (response *CreateK8sApiAbnormalRuleInfoResponse) {
    response = &CreateK8sApiAbnormalRuleInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateK8sApiAbnormalRuleInfo
// This API is used to create K8sApi abnormal event rules.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateK8sApiAbnormalRuleInfo(request *CreateK8sApiAbnormalRuleInfoRequest) (response *CreateK8sApiAbnormalRuleInfoResponse, err error) {
    return c.CreateK8sApiAbnormalRuleInfoWithContext(context.Background(), request)
}

// CreateK8sApiAbnormalRuleInfo
// This API is used to create K8sApi abnormal event rules.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateK8sApiAbnormalRuleInfoWithContext(ctx context.Context, request *CreateK8sApiAbnormalRuleInfoRequest) (response *CreateK8sApiAbnormalRuleInfoResponse, err error) {
    if request == nil {
        request = NewCreateK8sApiAbnormalRuleInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateK8sApiAbnormalRuleInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateK8sApiAbnormalRuleInfoResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNetworkFirewallClusterRefreshRequest() (request *CreateNetworkFirewallClusterRefreshRequest) {
    request = &CreateNetworkFirewallClusterRefreshRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "CreateNetworkFirewallClusterRefresh")
    
    
    return
}

func NewCreateNetworkFirewallClusterRefreshResponse() (response *CreateNetworkFirewallClusterRefreshResponse) {
    response = &CreateNetworkFirewallClusterRefreshResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateNetworkFirewallClusterRefresh
// This API is used to distribute a refresh task in the container network cluster.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateNetworkFirewallClusterRefresh(request *CreateNetworkFirewallClusterRefreshRequest) (response *CreateNetworkFirewallClusterRefreshResponse, err error) {
    return c.CreateNetworkFirewallClusterRefreshWithContext(context.Background(), request)
}

// CreateNetworkFirewallClusterRefresh
// This API is used to distribute a refresh task in the container network cluster.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateNetworkFirewallClusterRefreshWithContext(ctx context.Context, request *CreateNetworkFirewallClusterRefreshRequest) (response *CreateNetworkFirewallClusterRefreshResponse, err error) {
    if request == nil {
        request = NewCreateNetworkFirewallClusterRefreshRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNetworkFirewallClusterRefresh require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateNetworkFirewallClusterRefreshResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNetworkFirewallPolicyDiscoverRequest() (request *CreateNetworkFirewallPolicyDiscoverRequest) {
    request = &CreateNetworkFirewallPolicyDiscoverRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "CreateNetworkFirewallPolicyDiscover")
    
    
    return
}

func NewCreateNetworkFirewallPolicyDiscoverResponse() (response *CreateNetworkFirewallPolicyDiscoverResponse) {
    response = &CreateNetworkFirewallPolicyDiscoverResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateNetworkFirewallPolicyDiscover
// This API is used to create a task to sync a network policy from the container network cluster.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateNetworkFirewallPolicyDiscover(request *CreateNetworkFirewallPolicyDiscoverRequest) (response *CreateNetworkFirewallPolicyDiscoverResponse, err error) {
    return c.CreateNetworkFirewallPolicyDiscoverWithContext(context.Background(), request)
}

// CreateNetworkFirewallPolicyDiscover
// This API is used to create a task to sync a network policy from the container network cluster.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateNetworkFirewallPolicyDiscoverWithContext(ctx context.Context, request *CreateNetworkFirewallPolicyDiscoverRequest) (response *CreateNetworkFirewallPolicyDiscoverResponse, err error) {
    if request == nil {
        request = NewCreateNetworkFirewallPolicyDiscoverRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNetworkFirewallPolicyDiscover require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateNetworkFirewallPolicyDiscoverResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNetworkFirewallPublishRequest() (request *CreateNetworkFirewallPublishRequest) {
    request = &CreateNetworkFirewallPublishRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "CreateNetworkFirewallPublish")
    
    
    return
}

func NewCreateNetworkFirewallPublishResponse() (response *CreateNetworkFirewallPublishResponse) {
    response = &CreateNetworkFirewallPublishResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateNetworkFirewallPublish
// This API is used to create a task to publish a network policy in the container network.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateNetworkFirewallPublish(request *CreateNetworkFirewallPublishRequest) (response *CreateNetworkFirewallPublishResponse, err error) {
    return c.CreateNetworkFirewallPublishWithContext(context.Background(), request)
}

// CreateNetworkFirewallPublish
// This API is used to create a task to publish a network policy in the container network.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateNetworkFirewallPublishWithContext(ctx context.Context, request *CreateNetworkFirewallPublishRequest) (response *CreateNetworkFirewallPublishResponse, err error) {
    if request == nil {
        request = NewCreateNetworkFirewallPublishRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNetworkFirewallPublish require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateNetworkFirewallPublishResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNetworkFirewallUndoPublishRequest() (request *CreateNetworkFirewallUndoPublishRequest) {
    request = &CreateNetworkFirewallUndoPublishRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "CreateNetworkFirewallUndoPublish")
    
    
    return
}

func NewCreateNetworkFirewallUndoPublishResponse() (response *CreateNetworkFirewallUndoPublishResponse) {
    response = &CreateNetworkFirewallUndoPublishResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateNetworkFirewallUndoPublish
// This API is used to create a task to revoke a network policy in the container network.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateNetworkFirewallUndoPublish(request *CreateNetworkFirewallUndoPublishRequest) (response *CreateNetworkFirewallUndoPublishResponse, err error) {
    return c.CreateNetworkFirewallUndoPublishWithContext(context.Background(), request)
}

// CreateNetworkFirewallUndoPublish
// This API is used to create a task to revoke a network policy in the container network.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateNetworkFirewallUndoPublishWithContext(ctx context.Context, request *CreateNetworkFirewallUndoPublishRequest) (response *CreateNetworkFirewallUndoPublishResponse, err error) {
    if request == nil {
        request = NewCreateNetworkFirewallUndoPublishRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNetworkFirewallUndoPublish require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateNetworkFirewallUndoPublishResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOrModifyPostPayCoresRequest() (request *CreateOrModifyPostPayCoresRequest) {
    request = &CreateOrModifyPostPayCoresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "CreateOrModifyPostPayCores")
    
    
    return
}

func NewCreateOrModifyPostPayCoresResponse() (response *CreateOrModifyPostPayCoresResponse) {
    response = &CreateOrModifyPostPayCoresResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateOrModifyPostPayCores
// This API is used to create or edit the upper limit for elastic billing.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateOrModifyPostPayCores(request *CreateOrModifyPostPayCoresRequest) (response *CreateOrModifyPostPayCoresResponse, err error) {
    return c.CreateOrModifyPostPayCoresWithContext(context.Background(), request)
}

// CreateOrModifyPostPayCores
// This API is used to create or edit the upper limit for elastic billing.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateOrModifyPostPayCoresWithContext(ctx context.Context, request *CreateOrModifyPostPayCoresRequest) (response *CreateOrModifyPostPayCoresResponse, err error) {
    if request == nil {
        request = NewCreateOrModifyPostPayCoresRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOrModifyPostPayCores require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOrModifyPostPayCoresResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProcessEventsExportJobRequest() (request *CreateProcessEventsExportJobRequest) {
    request = &CreateProcessEventsExportJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "CreateProcessEventsExportJob")
    
    
    return
}

func NewCreateProcessEventsExportJobResponse() (response *CreateProcessEventsExportJobResponse) {
    response = &CreateProcessEventsExportJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateProcessEventsExportJob
// This API is used to create a task to asynchronously export abnormal process events.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateProcessEventsExportJob(request *CreateProcessEventsExportJobRequest) (response *CreateProcessEventsExportJobResponse, err error) {
    return c.CreateProcessEventsExportJobWithContext(context.Background(), request)
}

// CreateProcessEventsExportJob
// This API is used to create a task to asynchronously export abnormal process events.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateProcessEventsExportJobWithContext(ctx context.Context, request *CreateProcessEventsExportJobRequest) (response *CreateProcessEventsExportJobResponse, err error) {
    if request == nil {
        request = NewCreateProcessEventsExportJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateProcessEventsExportJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateProcessEventsExportJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRefreshTaskRequest() (request *CreateRefreshTaskRequest) {
    request = &CreateRefreshTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "CreateRefreshTask")
    
    
    return
}

func NewCreateRefreshTaskResponse() (response *CreateRefreshTaskResponse) {
    response = &CreateRefreshTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRefreshTask
// This API is used to distribute a task to refresh the asset information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateRefreshTask(request *CreateRefreshTaskRequest) (response *CreateRefreshTaskResponse, err error) {
    return c.CreateRefreshTaskWithContext(context.Background(), request)
}

// CreateRefreshTask
// This API is used to distribute a task to refresh the asset information.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateRefreshTaskWithContext(ctx context.Context, request *CreateRefreshTaskRequest) (response *CreateRefreshTaskResponse, err error) {
    if request == nil {
        request = NewCreateRefreshTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRefreshTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRefreshTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRiskDnsEventExportJobRequest() (request *CreateRiskDnsEventExportJobRequest) {
    request = &CreateRiskDnsEventExportJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "CreateRiskDnsEventExportJob")
    
    
    return
}

func NewCreateRiskDnsEventExportJobResponse() (response *CreateRiskDnsEventExportJobResponse) {
    response = &CreateRiskDnsEventExportJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRiskDnsEventExportJob
// This API is used to export malicious request events.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateRiskDnsEventExportJob(request *CreateRiskDnsEventExportJobRequest) (response *CreateRiskDnsEventExportJobResponse, err error) {
    return c.CreateRiskDnsEventExportJobWithContext(context.Background(), request)
}

// CreateRiskDnsEventExportJob
// This API is used to export malicious request events.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateRiskDnsEventExportJobWithContext(ctx context.Context, request *CreateRiskDnsEventExportJobRequest) (response *CreateRiskDnsEventExportJobResponse, err error) {
    if request == nil {
        request = NewCreateRiskDnsEventExportJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRiskDnsEventExportJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRiskDnsEventExportJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSearchTemplateRequest() (request *CreateSearchTemplateRequest) {
    request = &CreateSearchTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "CreateSearchTemplate")
    
    
    return
}

func NewCreateSearchTemplateResponse() (response *CreateSearchTemplateResponse) {
    response = &CreateSearchTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSearchTemplate
// This API is used to add a search template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateSearchTemplate(request *CreateSearchTemplateRequest) (response *CreateSearchTemplateResponse, err error) {
    return c.CreateSearchTemplateWithContext(context.Background(), request)
}

// CreateSearchTemplate
// This API is used to add a search template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateSearchTemplateWithContext(ctx context.Context, request *CreateSearchTemplateRequest) (response *CreateSearchTemplateResponse, err error) {
    if request == nil {
        request = NewCreateSearchTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSearchTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSearchTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSystemVulExportJobRequest() (request *CreateSystemVulExportJobRequest) {
    request = &CreateSystemVulExportJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "CreateSystemVulExportJob")
    
    
    return
}

func NewCreateSystemVulExportJobResponse() (response *CreateSystemVulExportJobResponse) {
    response = &CreateSystemVulExportJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSystemVulExportJob
// This API is used to create a task to export system vulnerabilities.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateSystemVulExportJob(request *CreateSystemVulExportJobRequest) (response *CreateSystemVulExportJobResponse, err error) {
    return c.CreateSystemVulExportJobWithContext(context.Background(), request)
}

// CreateSystemVulExportJob
// This API is used to create a task to export system vulnerabilities.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateSystemVulExportJobWithContext(ctx context.Context, request *CreateSystemVulExportJobRequest) (response *CreateSystemVulExportJobResponse, err error) {
    if request == nil {
        request = NewCreateSystemVulExportJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSystemVulExportJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSystemVulExportJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVirusScanAgainRequest() (request *CreateVirusScanAgainRequest) {
    request = &CreateVirusScanAgainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "CreateVirusScanAgain")
    
    
    return
}

func NewCreateVirusScanAgainResponse() (response *CreateVirusScanAgainResponse) {
    response = &CreateVirusScanAgainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateVirusScanAgain
// This API is used to perform another virus scan at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateVirusScanAgain(request *CreateVirusScanAgainRequest) (response *CreateVirusScanAgainResponse, err error) {
    return c.CreateVirusScanAgainWithContext(context.Background(), request)
}

// CreateVirusScanAgain
// This API is used to perform another virus scan at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateVirusScanAgainWithContext(ctx context.Context, request *CreateVirusScanAgainRequest) (response *CreateVirusScanAgainResponse, err error) {
    if request == nil {
        request = NewCreateVirusScanAgainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVirusScanAgain require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateVirusScanAgainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVirusScanTaskRequest() (request *CreateVirusScanTaskRequest) {
    request = &CreateVirusScanTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "CreateVirusScanTask")
    
    
    return
}

func NewCreateVirusScanTaskResponse() (response *CreateVirusScanTaskResponse) {
    response = &CreateVirusScanTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateVirusScanTask
// This API is used to perform a quick virus scan at runtime.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateVirusScanTask(request *CreateVirusScanTaskRequest) (response *CreateVirusScanTaskResponse, err error) {
    return c.CreateVirusScanTaskWithContext(context.Background(), request)
}

// CreateVirusScanTask
// This API is used to perform a quick virus scan at runtime.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateVirusScanTaskWithContext(ctx context.Context, request *CreateVirusScanTaskRequest) (response *CreateVirusScanTaskResponse, err error) {
    if request == nil {
        request = NewCreateVirusScanTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVirusScanTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateVirusScanTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVulContainerExportJobRequest() (request *CreateVulContainerExportJobRequest) {
    request = &CreateVulContainerExportJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "CreateVulContainerExportJob")
    
    
    return
}

func NewCreateVulContainerExportJobResponse() (response *CreateVulContainerExportJobResponse) {
    response = &CreateVulContainerExportJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateVulContainerExportJob
// This API is used to create a task to export containers affected by vulnerabilities.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateVulContainerExportJob(request *CreateVulContainerExportJobRequest) (response *CreateVulContainerExportJobResponse, err error) {
    return c.CreateVulContainerExportJobWithContext(context.Background(), request)
}

// CreateVulContainerExportJob
// This API is used to create a task to export containers affected by vulnerabilities.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateVulContainerExportJobWithContext(ctx context.Context, request *CreateVulContainerExportJobRequest) (response *CreateVulContainerExportJobResponse, err error) {
    if request == nil {
        request = NewCreateVulContainerExportJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVulContainerExportJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateVulContainerExportJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVulDefenceEventExportJobRequest() (request *CreateVulDefenceEventExportJobRequest) {
    request = &CreateVulDefenceEventExportJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "CreateVulDefenceEventExportJob")
    
    
    return
}

func NewCreateVulDefenceEventExportJobResponse() (response *CreateVulDefenceEventExportJobResponse) {
    response = &CreateVulDefenceEventExportJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateVulDefenceEventExportJob
// This API is used to create an exploit prevention event export task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateVulDefenceEventExportJob(request *CreateVulDefenceEventExportJobRequest) (response *CreateVulDefenceEventExportJobResponse, err error) {
    return c.CreateVulDefenceEventExportJobWithContext(context.Background(), request)
}

// CreateVulDefenceEventExportJob
// This API is used to create an exploit prevention event export task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateVulDefenceEventExportJobWithContext(ctx context.Context, request *CreateVulDefenceEventExportJobRequest) (response *CreateVulDefenceEventExportJobResponse, err error) {
    if request == nil {
        request = NewCreateVulDefenceEventExportJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVulDefenceEventExportJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateVulDefenceEventExportJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVulDefenceHostExportJobRequest() (request *CreateVulDefenceHostExportJobRequest) {
    request = &CreateVulDefenceHostExportJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "CreateVulDefenceHostExportJob")
    
    
    return
}

func NewCreateVulDefenceHostExportJobResponse() (response *CreateVulDefenceHostExportJobResponse) {
    response = &CreateVulDefenceHostExportJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateVulDefenceHostExportJob
// This API is used to create a task to export servers with exploit prevention enabled.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateVulDefenceHostExportJob(request *CreateVulDefenceHostExportJobRequest) (response *CreateVulDefenceHostExportJobResponse, err error) {
    return c.CreateVulDefenceHostExportJobWithContext(context.Background(), request)
}

// CreateVulDefenceHostExportJob
// This API is used to create a task to export servers with exploit prevention enabled.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateVulDefenceHostExportJobWithContext(ctx context.Context, request *CreateVulDefenceHostExportJobRequest) (response *CreateVulDefenceHostExportJobResponse, err error) {
    if request == nil {
        request = NewCreateVulDefenceHostExportJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVulDefenceHostExportJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateVulDefenceHostExportJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVulExportJobRequest() (request *CreateVulExportJobRequest) {
    request = &CreateVulExportJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "CreateVulExportJob")
    
    
    return
}

func NewCreateVulExportJobResponse() (response *CreateVulExportJobResponse) {
    response = &CreateVulExportJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateVulExportJob
// This API is used to export the list of vulnerabilities in a local image.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateVulExportJob(request *CreateVulExportJobRequest) (response *CreateVulExportJobResponse, err error) {
    return c.CreateVulExportJobWithContext(context.Background(), request)
}

// CreateVulExportJob
// This API is used to export the list of vulnerabilities in a local image.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateVulExportJobWithContext(ctx context.Context, request *CreateVulExportJobRequest) (response *CreateVulExportJobResponse, err error) {
    if request == nil {
        request = NewCreateVulExportJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVulExportJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateVulExportJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVulImageExportJobRequest() (request *CreateVulImageExportJobRequest) {
    request = &CreateVulImageExportJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "CreateVulImageExportJob")
    
    
    return
}

func NewCreateVulImageExportJobResponse() (response *CreateVulImageExportJobResponse) {
    response = &CreateVulImageExportJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateVulImageExportJob
// This API is used to create a task to export images affected by vulnerabilities.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateVulImageExportJob(request *CreateVulImageExportJobRequest) (response *CreateVulImageExportJobResponse, err error) {
    return c.CreateVulImageExportJobWithContext(context.Background(), request)
}

// CreateVulImageExportJob
// This API is used to create a task to export images affected by vulnerabilities.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateVulImageExportJobWithContext(ctx context.Context, request *CreateVulImageExportJobRequest) (response *CreateVulImageExportJobResponse, err error) {
    if request == nil {
        request = NewCreateVulImageExportJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVulImageExportJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateVulImageExportJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVulScanTaskRequest() (request *CreateVulScanTaskRequest) {
    request = &CreateVulScanTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "CreateVulScanTask")
    
    
    return
}

func NewCreateVulScanTaskResponse() (response *CreateVulScanTaskResponse) {
    response = &CreateVulScanTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateVulScanTask
// This API is used to create a vulnerability scan task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateVulScanTask(request *CreateVulScanTaskRequest) (response *CreateVulScanTaskResponse, err error) {
    return c.CreateVulScanTaskWithContext(context.Background(), request)
}

// CreateVulScanTask
// This API is used to create a vulnerability scan task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateVulScanTaskWithContext(ctx context.Context, request *CreateVulScanTaskRequest) (response *CreateVulScanTaskResponse, err error) {
    if request == nil {
        request = NewCreateVulScanTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVulScanTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateVulScanTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWebVulExportJobRequest() (request *CreateWebVulExportJobRequest) {
    request = &CreateWebVulExportJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "CreateWebVulExportJob")
    
    
    return
}

func NewCreateWebVulExportJobResponse() (response *CreateWebVulExportJobResponse) {
    response = &CreateWebVulExportJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateWebVulExportJob
// This API is used to create a task to export web vulnerabilities.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateWebVulExportJob(request *CreateWebVulExportJobRequest) (response *CreateWebVulExportJobResponse, err error) {
    return c.CreateWebVulExportJobWithContext(context.Background(), request)
}

// CreateWebVulExportJob
// This API is used to create a task to export web vulnerabilities.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateWebVulExportJobWithContext(ctx context.Context, request *CreateWebVulExportJobRequest) (response *CreateWebVulExportJobResponse, err error) {
    if request == nil {
        request = NewCreateWebVulExportJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWebVulExportJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWebVulExportJobResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAbnormalProcessRulesRequest() (request *DeleteAbnormalProcessRulesRequest) {
    request = &DeleteAbnormalProcessRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DeleteAbnormalProcessRules")
    
    
    return
}

func NewDeleteAbnormalProcessRulesResponse() (response *DeleteAbnormalProcessRulesResponse) {
    response = &DeleteAbnormalProcessRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAbnormalProcessRules
// This API is used to delete an abnormal process policy at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteAbnormalProcessRules(request *DeleteAbnormalProcessRulesRequest) (response *DeleteAbnormalProcessRulesResponse, err error) {
    return c.DeleteAbnormalProcessRulesWithContext(context.Background(), request)
}

// DeleteAbnormalProcessRules
// This API is used to delete an abnormal process policy at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteAbnormalProcessRulesWithContext(ctx context.Context, request *DeleteAbnormalProcessRulesRequest) (response *DeleteAbnormalProcessRulesResponse, err error) {
    if request == nil {
        request = NewDeleteAbnormalProcessRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAbnormalProcessRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAbnormalProcessRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAccessControlRulesRequest() (request *DeleteAccessControlRulesRequest) {
    request = &DeleteAccessControlRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DeleteAccessControlRules")
    
    
    return
}

func NewDeleteAccessControlRulesResponse() (response *DeleteAccessControlRulesResponse) {
    response = &DeleteAccessControlRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAccessControlRules
// This API is used to delete an access control policy at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteAccessControlRules(request *DeleteAccessControlRulesRequest) (response *DeleteAccessControlRulesResponse, err error) {
    return c.DeleteAccessControlRulesWithContext(context.Background(), request)
}

// DeleteAccessControlRules
// This API is used to delete an access control policy at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteAccessControlRulesWithContext(ctx context.Context, request *DeleteAccessControlRulesRequest) (response *DeleteAccessControlRulesResponse, err error) {
    if request == nil {
        request = NewDeleteAccessControlRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAccessControlRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAccessControlRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteComplianceAssetPolicySetFromWhitelistRequest() (request *DeleteComplianceAssetPolicySetFromWhitelistRequest) {
    request = &DeleteComplianceAssetPolicySetFromWhitelistRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DeleteComplianceAssetPolicySetFromWhitelist")
    
    
    return
}

func NewDeleteComplianceAssetPolicySetFromWhitelistResponse() (response *DeleteComplianceAssetPolicySetFromWhitelistResponse) {
    response = &DeleteComplianceAssetPolicySetFromWhitelistResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteComplianceAssetPolicySetFromWhitelist
// This API is used to unignore the specified asset IDs and check item IDs so as to show the assets contained in the specified check items.
//
// `AddCompliancePolicyAssetSetToWhitelist` is the reference API. Except for the input field, others should be the same, and if not, it may be due to the definition.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DeleteComplianceAssetPolicySetFromWhitelist(request *DeleteComplianceAssetPolicySetFromWhitelistRequest) (response *DeleteComplianceAssetPolicySetFromWhitelistResponse, err error) {
    return c.DeleteComplianceAssetPolicySetFromWhitelistWithContext(context.Background(), request)
}

// DeleteComplianceAssetPolicySetFromWhitelist
// This API is used to unignore the specified asset IDs and check item IDs so as to show the assets contained in the specified check items.
//
// `AddCompliancePolicyAssetSetToWhitelist` is the reference API. Except for the input field, others should be the same, and if not, it may be due to the definition.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DeleteComplianceAssetPolicySetFromWhitelistWithContext(ctx context.Context, request *DeleteComplianceAssetPolicySetFromWhitelistRequest) (response *DeleteComplianceAssetPolicySetFromWhitelistResponse, err error) {
    if request == nil {
        request = NewDeleteComplianceAssetPolicySetFromWhitelistRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteComplianceAssetPolicySetFromWhitelist require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteComplianceAssetPolicySetFromWhitelistResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCompliancePolicyAssetSetFromWhitelistRequest() (request *DeleteCompliancePolicyAssetSetFromWhitelistRequest) {
    request = &DeleteCompliancePolicyAssetSetFromWhitelistRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DeleteCompliancePolicyAssetSetFromWhitelist")
    
    
    return
}

func NewDeleteCompliancePolicyAssetSetFromWhitelistResponse() (response *DeleteCompliancePolicyAssetSetFromWhitelistResponse) {
    response = &DeleteCompliancePolicyAssetSetFromWhitelistResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCompliancePolicyAssetSetFromWhitelist
// This API is used to unignore the specified asset IDs and check item IDs so as to show the assets contained in the specified check items.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DeleteCompliancePolicyAssetSetFromWhitelist(request *DeleteCompliancePolicyAssetSetFromWhitelistRequest) (response *DeleteCompliancePolicyAssetSetFromWhitelistResponse, err error) {
    return c.DeleteCompliancePolicyAssetSetFromWhitelistWithContext(context.Background(), request)
}

// DeleteCompliancePolicyAssetSetFromWhitelist
// This API is used to unignore the specified asset IDs and check item IDs so as to show the assets contained in the specified check items.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DeleteCompliancePolicyAssetSetFromWhitelistWithContext(ctx context.Context, request *DeleteCompliancePolicyAssetSetFromWhitelistRequest) (response *DeleteCompliancePolicyAssetSetFromWhitelistResponse, err error) {
    if request == nil {
        request = NewDeleteCompliancePolicyAssetSetFromWhitelistRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCompliancePolicyAssetSetFromWhitelist require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCompliancePolicyAssetSetFromWhitelistResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCompliancePolicyItemFromWhitelistRequest() (request *DeleteCompliancePolicyItemFromWhitelistRequest) {
    request = &DeleteCompliancePolicyItemFromWhitelistRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DeleteCompliancePolicyItemFromWhitelist")
    
    
    return
}

func NewDeleteCompliancePolicyItemFromWhitelistResponse() (response *DeleteCompliancePolicyItemFromWhitelistResponse) {
    response = &DeleteCompliancePolicyItemFromWhitelistResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCompliancePolicyItemFromWhitelist
// This API is used to remove the specified check item from the allowlist.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DeleteCompliancePolicyItemFromWhitelist(request *DeleteCompliancePolicyItemFromWhitelistRequest) (response *DeleteCompliancePolicyItemFromWhitelistResponse, err error) {
    return c.DeleteCompliancePolicyItemFromWhitelistWithContext(context.Background(), request)
}

// DeleteCompliancePolicyItemFromWhitelist
// This API is used to remove the specified check item from the allowlist.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DeleteCompliancePolicyItemFromWhitelistWithContext(ctx context.Context, request *DeleteCompliancePolicyItemFromWhitelistRequest) (response *DeleteCompliancePolicyItemFromWhitelistResponse, err error) {
    if request == nil {
        request = NewDeleteCompliancePolicyItemFromWhitelistRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCompliancePolicyItemFromWhitelist require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCompliancePolicyItemFromWhitelistResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEscapeWhiteListRequest() (request *DeleteEscapeWhiteListRequest) {
    request = &DeleteEscapeWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DeleteEscapeWhiteList")
    
    
    return
}

func NewDeleteEscapeWhiteListResponse() (response *DeleteEscapeWhiteListResponse) {
    response = &DeleteEscapeWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteEscapeWhiteList
// This API is used to delete an allowed escape.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteEscapeWhiteList(request *DeleteEscapeWhiteListRequest) (response *DeleteEscapeWhiteListResponse, err error) {
    return c.DeleteEscapeWhiteListWithContext(context.Background(), request)
}

// DeleteEscapeWhiteList
// This API is used to delete an allowed escape.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteEscapeWhiteListWithContext(ctx context.Context, request *DeleteEscapeWhiteListRequest) (response *DeleteEscapeWhiteListResponse, err error) {
    if request == nil {
        request = NewDeleteEscapeWhiteListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteEscapeWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteEscapeWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteIgnoreVulRequest() (request *DeleteIgnoreVulRequest) {
    request = &DeleteIgnoreVulRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DeleteIgnoreVul")
    
    
    return
}

func NewDeleteIgnoreVulResponse() (response *DeleteIgnoreVulResponse) {
    response = &DeleteIgnoreVulResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteIgnoreVul
// This API is used to unignore vulnerabilities in a vulnerability scan.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteIgnoreVul(request *DeleteIgnoreVulRequest) (response *DeleteIgnoreVulResponse, err error) {
    return c.DeleteIgnoreVulWithContext(context.Background(), request)
}

// DeleteIgnoreVul
// This API is used to unignore vulnerabilities in a vulnerability scan.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteIgnoreVulWithContext(ctx context.Context, request *DeleteIgnoreVulRequest) (response *DeleteIgnoreVulResponse, err error) {
    if request == nil {
        request = NewDeleteIgnoreVulRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteIgnoreVul require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteIgnoreVulResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteK8sApiAbnormalRuleRequest() (request *DeleteK8sApiAbnormalRuleRequest) {
    request = &DeleteK8sApiAbnormalRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DeleteK8sApiAbnormalRule")
    
    
    return
}

func NewDeleteK8sApiAbnormalRuleResponse() (response *DeleteK8sApiAbnormalRuleResponse) {
    response = &DeleteK8sApiAbnormalRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteK8sApiAbnormalRule
// This API is used to delete a K8sApi abnormal event rules.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteK8sApiAbnormalRule(request *DeleteK8sApiAbnormalRuleRequest) (response *DeleteK8sApiAbnormalRuleResponse, err error) {
    return c.DeleteK8sApiAbnormalRuleWithContext(context.Background(), request)
}

// DeleteK8sApiAbnormalRule
// This API is used to delete a K8sApi abnormal event rules.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteK8sApiAbnormalRuleWithContext(ctx context.Context, request *DeleteK8sApiAbnormalRuleRequest) (response *DeleteK8sApiAbnormalRuleResponse, err error) {
    if request == nil {
        request = NewDeleteK8sApiAbnormalRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteK8sApiAbnormalRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteK8sApiAbnormalRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMachineRequest() (request *DeleteMachineRequest) {
    request = &DeleteMachineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DeleteMachine")
    
    
    return
}

func NewDeleteMachineResponse() (response *DeleteMachineResponse) {
    response = &DeleteMachineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteMachine
// This API is used to uninstall the agent.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENTOFFLINE = "FailedOperation.AgentOffline"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteMachine(request *DeleteMachineRequest) (response *DeleteMachineResponse, err error) {
    return c.DeleteMachineWithContext(context.Background(), request)
}

// DeleteMachine
// This API is used to uninstall the agent.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENTOFFLINE = "FailedOperation.AgentOffline"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteMachineWithContext(ctx context.Context, request *DeleteMachineRequest) (response *DeleteMachineResponse, err error) {
    if request == nil {
        request = NewDeleteMachineRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMachine require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMachineResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNetworkFirewallPolicyDetailRequest() (request *DeleteNetworkFirewallPolicyDetailRequest) {
    request = &DeleteNetworkFirewallPolicyDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DeleteNetworkFirewallPolicyDetail")
    
    
    return
}

func NewDeleteNetworkFirewallPolicyDetailResponse() (response *DeleteNetworkFirewallPolicyDetailResponse) {
    response = &DeleteNetworkFirewallPolicyDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteNetworkFirewallPolicyDetail
// This API is used to create a task to delete a network policy in the container network.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteNetworkFirewallPolicyDetail(request *DeleteNetworkFirewallPolicyDetailRequest) (response *DeleteNetworkFirewallPolicyDetailResponse, err error) {
    return c.DeleteNetworkFirewallPolicyDetailWithContext(context.Background(), request)
}

// DeleteNetworkFirewallPolicyDetail
// This API is used to create a task to delete a network policy in the container network.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteNetworkFirewallPolicyDetailWithContext(ctx context.Context, request *DeleteNetworkFirewallPolicyDetailRequest) (response *DeleteNetworkFirewallPolicyDetailResponse, err error) {
    if request == nil {
        request = NewDeleteNetworkFirewallPolicyDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteNetworkFirewallPolicyDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteNetworkFirewallPolicyDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteReverseShellEventsRequest() (request *DeleteReverseShellEventsRequest) {
    request = &DeleteReverseShellEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DeleteReverseShellEvents")
    
    
    return
}

func NewDeleteReverseShellEventsResponse() (response *DeleteReverseShellEventsResponse) {
    response = &DeleteReverseShellEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteReverseShellEvents
// This API is used to delete a reverse shell event at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteReverseShellEvents(request *DeleteReverseShellEventsRequest) (response *DeleteReverseShellEventsResponse, err error) {
    return c.DeleteReverseShellEventsWithContext(context.Background(), request)
}

// DeleteReverseShellEvents
// This API is used to delete a reverse shell event at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteReverseShellEventsWithContext(ctx context.Context, request *DeleteReverseShellEventsRequest) (response *DeleteReverseShellEventsResponse, err error) {
    if request == nil {
        request = NewDeleteReverseShellEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteReverseShellEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteReverseShellEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteReverseShellWhiteListsRequest() (request *DeleteReverseShellWhiteListsRequest) {
    request = &DeleteReverseShellWhiteListsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DeleteReverseShellWhiteLists")
    
    
    return
}

func NewDeleteReverseShellWhiteListsResponse() (response *DeleteReverseShellWhiteListsResponse) {
    response = &DeleteReverseShellWhiteListsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteReverseShellWhiteLists
// This API is used to delete an allowed reverse shell at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteReverseShellWhiteLists(request *DeleteReverseShellWhiteListsRequest) (response *DeleteReverseShellWhiteListsResponse, err error) {
    return c.DeleteReverseShellWhiteListsWithContext(context.Background(), request)
}

// DeleteReverseShellWhiteLists
// This API is used to delete an allowed reverse shell at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteReverseShellWhiteListsWithContext(ctx context.Context, request *DeleteReverseShellWhiteListsRequest) (response *DeleteReverseShellWhiteListsResponse, err error) {
    if request == nil {
        request = NewDeleteReverseShellWhiteListsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteReverseShellWhiteLists require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteReverseShellWhiteListsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRiskSyscallEventsRequest() (request *DeleteRiskSyscallEventsRequest) {
    request = &DeleteRiskSyscallEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DeleteRiskSyscallEvents")
    
    
    return
}

func NewDeleteRiskSyscallEventsResponse() (response *DeleteRiskSyscallEventsResponse) {
    response = &DeleteRiskSyscallEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRiskSyscallEvents
// This API is used to delete a high-risk syscall event at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteRiskSyscallEvents(request *DeleteRiskSyscallEventsRequest) (response *DeleteRiskSyscallEventsResponse, err error) {
    return c.DeleteRiskSyscallEventsWithContext(context.Background(), request)
}

// DeleteRiskSyscallEvents
// This API is used to delete a high-risk syscall event at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteRiskSyscallEventsWithContext(ctx context.Context, request *DeleteRiskSyscallEventsRequest) (response *DeleteRiskSyscallEventsResponse, err error) {
    if request == nil {
        request = NewDeleteRiskSyscallEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRiskSyscallEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRiskSyscallEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRiskSyscallWhiteListsRequest() (request *DeleteRiskSyscallWhiteListsRequest) {
    request = &DeleteRiskSyscallWhiteListsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DeleteRiskSyscallWhiteLists")
    
    
    return
}

func NewDeleteRiskSyscallWhiteListsResponse() (response *DeleteRiskSyscallWhiteListsResponse) {
    response = &DeleteRiskSyscallWhiteListsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRiskSyscallWhiteLists
// This API is used to delete an allowed high-risk syscall at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteRiskSyscallWhiteLists(request *DeleteRiskSyscallWhiteListsRequest) (response *DeleteRiskSyscallWhiteListsResponse, err error) {
    return c.DeleteRiskSyscallWhiteListsWithContext(context.Background(), request)
}

// DeleteRiskSyscallWhiteLists
// This API is used to delete an allowed high-risk syscall at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteRiskSyscallWhiteListsWithContext(ctx context.Context, request *DeleteRiskSyscallWhiteListsRequest) (response *DeleteRiskSyscallWhiteListsResponse, err error) {
    if request == nil {
        request = NewDeleteRiskSyscallWhiteListsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRiskSyscallWhiteLists require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRiskSyscallWhiteListsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSearchTemplateRequest() (request *DeleteSearchTemplateRequest) {
    request = &DeleteSearchTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DeleteSearchTemplate")
    
    
    return
}

func NewDeleteSearchTemplateResponse() (response *DeleteSearchTemplateResponse) {
    response = &DeleteSearchTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSearchTemplate
// This API is used to delete a search template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteSearchTemplate(request *DeleteSearchTemplateRequest) (response *DeleteSearchTemplateResponse, err error) {
    return c.DeleteSearchTemplateWithContext(context.Background(), request)
}

// DeleteSearchTemplate
// This API is used to delete a search template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteSearchTemplateWithContext(ctx context.Context, request *DeleteSearchTemplateRequest) (response *DeleteSearchTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteSearchTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSearchTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSearchTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeABTestConfigRequest() (request *DescribeABTestConfigRequest) {
    request = &DescribeABTestConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeABTestConfig")
    
    
    return
}

func NewDescribeABTestConfigResponse() (response *DescribeABTestConfigResponse) {
    response = &DescribeABTestConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeABTestConfig
// This API is used to get the current canary configuration of the user.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeABTestConfig(request *DescribeABTestConfigRequest) (response *DescribeABTestConfigResponse, err error) {
    return c.DescribeABTestConfigWithContext(context.Background(), request)
}

// DescribeABTestConfig
// This API is used to get the current canary configuration of the user.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeABTestConfigWithContext(ctx context.Context, request *DescribeABTestConfigRequest) (response *DescribeABTestConfigResponse, err error) {
    if request == nil {
        request = NewDescribeABTestConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeABTestConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeABTestConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAbnormalProcessDetailRequest() (request *DescribeAbnormalProcessDetailRequest) {
    request = &DescribeAbnormalProcessDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAbnormalProcessDetail")
    
    
    return
}

func NewDescribeAbnormalProcessDetailResponse() (response *DescribeAbnormalProcessDetailResponse) {
    response = &DescribeAbnormalProcessDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAbnormalProcessDetail
// This API is used to query the details of an abnormal process event at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAbnormalProcessDetail(request *DescribeAbnormalProcessDetailRequest) (response *DescribeAbnormalProcessDetailResponse, err error) {
    return c.DescribeAbnormalProcessDetailWithContext(context.Background(), request)
}

// DescribeAbnormalProcessDetail
// This API is used to query the details of an abnormal process event at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAbnormalProcessDetailWithContext(ctx context.Context, request *DescribeAbnormalProcessDetailRequest) (response *DescribeAbnormalProcessDetailResponse, err error) {
    if request == nil {
        request = NewDescribeAbnormalProcessDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAbnormalProcessDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAbnormalProcessDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAbnormalProcessEventTendencyRequest() (request *DescribeAbnormalProcessEventTendencyRequest) {
    request = &DescribeAbnormalProcessEventTendencyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAbnormalProcessEventTendency")
    
    
    return
}

func NewDescribeAbnormalProcessEventTendencyResponse() (response *DescribeAbnormalProcessEventTendencyResponse) {
    response = &DescribeAbnormalProcessEventTendencyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAbnormalProcessEventTendency
// This API is used to query the trend of pending abnormal process events.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAbnormalProcessEventTendency(request *DescribeAbnormalProcessEventTendencyRequest) (response *DescribeAbnormalProcessEventTendencyResponse, err error) {
    return c.DescribeAbnormalProcessEventTendencyWithContext(context.Background(), request)
}

// DescribeAbnormalProcessEventTendency
// This API is used to query the trend of pending abnormal process events.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAbnormalProcessEventTendencyWithContext(ctx context.Context, request *DescribeAbnormalProcessEventTendencyRequest) (response *DescribeAbnormalProcessEventTendencyResponse, err error) {
    if request == nil {
        request = NewDescribeAbnormalProcessEventTendencyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAbnormalProcessEventTendency require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAbnormalProcessEventTendencyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAbnormalProcessEventsRequest() (request *DescribeAbnormalProcessEventsRequest) {
    request = &DescribeAbnormalProcessEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAbnormalProcessEvents")
    
    
    return
}

func NewDescribeAbnormalProcessEventsResponse() (response *DescribeAbnormalProcessEventsResponse) {
    response = &DescribeAbnormalProcessEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAbnormalProcessEvents
// This API is used to query the list of abnormal process events at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAbnormalProcessEvents(request *DescribeAbnormalProcessEventsRequest) (response *DescribeAbnormalProcessEventsResponse, err error) {
    return c.DescribeAbnormalProcessEventsWithContext(context.Background(), request)
}

// DescribeAbnormalProcessEvents
// This API is used to query the list of abnormal process events at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAbnormalProcessEventsWithContext(ctx context.Context, request *DescribeAbnormalProcessEventsRequest) (response *DescribeAbnormalProcessEventsResponse, err error) {
    if request == nil {
        request = NewDescribeAbnormalProcessEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAbnormalProcessEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAbnormalProcessEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAbnormalProcessEventsExportRequest() (request *DescribeAbnormalProcessEventsExportRequest) {
    request = &DescribeAbnormalProcessEventsExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAbnormalProcessEventsExport")
    
    
    return
}

func NewDescribeAbnormalProcessEventsExportResponse() (response *DescribeAbnormalProcessEventsExportResponse) {
    response = &DescribeAbnormalProcessEventsExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAbnormalProcessEventsExport
// 接口已废弃
//
// 
//
// This API is used to query and export the list of abnormal process events at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAbnormalProcessEventsExport(request *DescribeAbnormalProcessEventsExportRequest) (response *DescribeAbnormalProcessEventsExportResponse, err error) {
    return c.DescribeAbnormalProcessEventsExportWithContext(context.Background(), request)
}

// DescribeAbnormalProcessEventsExport
// 接口已废弃
//
// 
//
// This API is used to query and export the list of abnormal process events at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAbnormalProcessEventsExportWithContext(ctx context.Context, request *DescribeAbnormalProcessEventsExportRequest) (response *DescribeAbnormalProcessEventsExportResponse, err error) {
    if request == nil {
        request = NewDescribeAbnormalProcessEventsExportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAbnormalProcessEventsExport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAbnormalProcessEventsExportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAbnormalProcessLevelSummaryRequest() (request *DescribeAbnormalProcessLevelSummaryRequest) {
    request = &DescribeAbnormalProcessLevelSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAbnormalProcessLevelSummary")
    
    
    return
}

func NewDescribeAbnormalProcessLevelSummaryResponse() (response *DescribeAbnormalProcessLevelSummaryResponse) {
    response = &DescribeAbnormalProcessLevelSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAbnormalProcessLevelSummary
// This API is used to count the number of pending abnormal process events at each severity level.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAbnormalProcessLevelSummary(request *DescribeAbnormalProcessLevelSummaryRequest) (response *DescribeAbnormalProcessLevelSummaryResponse, err error) {
    return c.DescribeAbnormalProcessLevelSummaryWithContext(context.Background(), request)
}

// DescribeAbnormalProcessLevelSummary
// This API is used to count the number of pending abnormal process events at each severity level.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAbnormalProcessLevelSummaryWithContext(ctx context.Context, request *DescribeAbnormalProcessLevelSummaryRequest) (response *DescribeAbnormalProcessLevelSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeAbnormalProcessLevelSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAbnormalProcessLevelSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAbnormalProcessLevelSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAbnormalProcessRuleDetailRequest() (request *DescribeAbnormalProcessRuleDetailRequest) {
    request = &DescribeAbnormalProcessRuleDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAbnormalProcessRuleDetail")
    
    
    return
}

func NewDescribeAbnormalProcessRuleDetailResponse() (response *DescribeAbnormalProcessRuleDetailResponse) {
    response = &DescribeAbnormalProcessRuleDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAbnormalProcessRuleDetail
// This API is used to query the details of an abnormal process policy at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_RULENOTFIND = "FailedOperation.RuleNotFind"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAbnormalProcessRuleDetail(request *DescribeAbnormalProcessRuleDetailRequest) (response *DescribeAbnormalProcessRuleDetailResponse, err error) {
    return c.DescribeAbnormalProcessRuleDetailWithContext(context.Background(), request)
}

// DescribeAbnormalProcessRuleDetail
// This API is used to query the details of an abnormal process policy at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_RULENOTFIND = "FailedOperation.RuleNotFind"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAbnormalProcessRuleDetailWithContext(ctx context.Context, request *DescribeAbnormalProcessRuleDetailRequest) (response *DescribeAbnormalProcessRuleDetailResponse, err error) {
    if request == nil {
        request = NewDescribeAbnormalProcessRuleDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAbnormalProcessRuleDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAbnormalProcessRuleDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAbnormalProcessRulesRequest() (request *DescribeAbnormalProcessRulesRequest) {
    request = &DescribeAbnormalProcessRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAbnormalProcessRules")
    
    
    return
}

func NewDescribeAbnormalProcessRulesResponse() (response *DescribeAbnormalProcessRulesResponse) {
    response = &DescribeAbnormalProcessRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAbnormalProcessRules
// This API is used to query the list of abnormal process policies at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAbnormalProcessRules(request *DescribeAbnormalProcessRulesRequest) (response *DescribeAbnormalProcessRulesResponse, err error) {
    return c.DescribeAbnormalProcessRulesWithContext(context.Background(), request)
}

// DescribeAbnormalProcessRules
// This API is used to query the list of abnormal process policies at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAbnormalProcessRulesWithContext(ctx context.Context, request *DescribeAbnormalProcessRulesRequest) (response *DescribeAbnormalProcessRulesResponse, err error) {
    if request == nil {
        request = NewDescribeAbnormalProcessRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAbnormalProcessRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAbnormalProcessRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAbnormalProcessRulesExportRequest() (request *DescribeAbnormalProcessRulesExportRequest) {
    request = &DescribeAbnormalProcessRulesExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAbnormalProcessRulesExport")
    
    
    return
}

func NewDescribeAbnormalProcessRulesExportResponse() (response *DescribeAbnormalProcessRulesExportResponse) {
    response = &DescribeAbnormalProcessRulesExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAbnormalProcessRulesExport
// 接口已废弃
//
// 
//
// This API is used to query and export the list of abnormal process policies at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAbnormalProcessRulesExport(request *DescribeAbnormalProcessRulesExportRequest) (response *DescribeAbnormalProcessRulesExportResponse, err error) {
    return c.DescribeAbnormalProcessRulesExportWithContext(context.Background(), request)
}

// DescribeAbnormalProcessRulesExport
// 接口已废弃
//
// 
//
// This API is used to query and export the list of abnormal process policies at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAbnormalProcessRulesExportWithContext(ctx context.Context, request *DescribeAbnormalProcessRulesExportRequest) (response *DescribeAbnormalProcessRulesExportResponse, err error) {
    if request == nil {
        request = NewDescribeAbnormalProcessRulesExportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAbnormalProcessRulesExport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAbnormalProcessRulesExportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessControlDetailRequest() (request *DescribeAccessControlDetailRequest) {
    request = &DescribeAccessControlDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAccessControlDetail")
    
    
    return
}

func NewDescribeAccessControlDetailResponse() (response *DescribeAccessControlDetailResponse) {
    response = &DescribeAccessControlDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccessControlDetail
// This API is used to query the details of an access control event at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAccessControlDetail(request *DescribeAccessControlDetailRequest) (response *DescribeAccessControlDetailResponse, err error) {
    return c.DescribeAccessControlDetailWithContext(context.Background(), request)
}

// DescribeAccessControlDetail
// This API is used to query the details of an access control event at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAccessControlDetailWithContext(ctx context.Context, request *DescribeAccessControlDetailRequest) (response *DescribeAccessControlDetailResponse, err error) {
    if request == nil {
        request = NewDescribeAccessControlDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccessControlDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccessControlDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessControlEventsRequest() (request *DescribeAccessControlEventsRequest) {
    request = &DescribeAccessControlEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAccessControlEvents")
    
    
    return
}

func NewDescribeAccessControlEventsResponse() (response *DescribeAccessControlEventsResponse) {
    response = &DescribeAccessControlEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccessControlEvents
// This API is used to query the list of access control events at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAccessControlEvents(request *DescribeAccessControlEventsRequest) (response *DescribeAccessControlEventsResponse, err error) {
    return c.DescribeAccessControlEventsWithContext(context.Background(), request)
}

// DescribeAccessControlEvents
// This API is used to query the list of access control events at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAccessControlEventsWithContext(ctx context.Context, request *DescribeAccessControlEventsRequest) (response *DescribeAccessControlEventsResponse, err error) {
    if request == nil {
        request = NewDescribeAccessControlEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccessControlEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccessControlEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessControlEventsExportRequest() (request *DescribeAccessControlEventsExportRequest) {
    request = &DescribeAccessControlEventsExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAccessControlEventsExport")
    
    
    return
}

func NewDescribeAccessControlEventsExportResponse() (response *DescribeAccessControlEventsExportResponse) {
    response = &DescribeAccessControlEventsExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccessControlEventsExport
// This API is used to export the list of access control events at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAccessControlEventsExport(request *DescribeAccessControlEventsExportRequest) (response *DescribeAccessControlEventsExportResponse, err error) {
    return c.DescribeAccessControlEventsExportWithContext(context.Background(), request)
}

// DescribeAccessControlEventsExport
// This API is used to export the list of access control events at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAccessControlEventsExportWithContext(ctx context.Context, request *DescribeAccessControlEventsExportRequest) (response *DescribeAccessControlEventsExportResponse, err error) {
    if request == nil {
        request = NewDescribeAccessControlEventsExportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccessControlEventsExport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccessControlEventsExportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessControlRuleDetailRequest() (request *DescribeAccessControlRuleDetailRequest) {
    request = &DescribeAccessControlRuleDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAccessControlRuleDetail")
    
    
    return
}

func NewDescribeAccessControlRuleDetailResponse() (response *DescribeAccessControlRuleDetailResponse) {
    response = &DescribeAccessControlRuleDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccessControlRuleDetail
// This API is used to query the details of an access control policy at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_RULENOTFIND = "FailedOperation.RuleNotFind"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAccessControlRuleDetail(request *DescribeAccessControlRuleDetailRequest) (response *DescribeAccessControlRuleDetailResponse, err error) {
    return c.DescribeAccessControlRuleDetailWithContext(context.Background(), request)
}

// DescribeAccessControlRuleDetail
// This API is used to query the details of an access control policy at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_RULENOTFIND = "FailedOperation.RuleNotFind"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAccessControlRuleDetailWithContext(ctx context.Context, request *DescribeAccessControlRuleDetailRequest) (response *DescribeAccessControlRuleDetailResponse, err error) {
    if request == nil {
        request = NewDescribeAccessControlRuleDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccessControlRuleDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccessControlRuleDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessControlRulesRequest() (request *DescribeAccessControlRulesRequest) {
    request = &DescribeAccessControlRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAccessControlRules")
    
    
    return
}

func NewDescribeAccessControlRulesResponse() (response *DescribeAccessControlRulesResponse) {
    response = &DescribeAccessControlRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccessControlRules
// This API is used to query the list of access control policies at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAccessControlRules(request *DescribeAccessControlRulesRequest) (response *DescribeAccessControlRulesResponse, err error) {
    return c.DescribeAccessControlRulesWithContext(context.Background(), request)
}

// DescribeAccessControlRules
// This API is used to query the list of access control policies at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAccessControlRulesWithContext(ctx context.Context, request *DescribeAccessControlRulesRequest) (response *DescribeAccessControlRulesResponse, err error) {
    if request == nil {
        request = NewDescribeAccessControlRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccessControlRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccessControlRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessControlRulesExportRequest() (request *DescribeAccessControlRulesExportRequest) {
    request = &DescribeAccessControlRulesExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAccessControlRulesExport")
    
    
    return
}

func NewDescribeAccessControlRulesExportResponse() (response *DescribeAccessControlRulesExportResponse) {
    response = &DescribeAccessControlRulesExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccessControlRulesExport
// 接口已废弃
//
// 
//
// This API is used to export the list of access control policies at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAccessControlRulesExport(request *DescribeAccessControlRulesExportRequest) (response *DescribeAccessControlRulesExportResponse, err error) {
    return c.DescribeAccessControlRulesExportWithContext(context.Background(), request)
}

// DescribeAccessControlRulesExport
// 接口已废弃
//
// 
//
// This API is used to export the list of access control policies at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAccessControlRulesExportWithContext(ctx context.Context, request *DescribeAccessControlRulesExportRequest) (response *DescribeAccessControlRulesExportResponse, err error) {
    if request == nil {
        request = NewDescribeAccessControlRulesExportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccessControlRulesExport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccessControlRulesExportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAffectedClusterCountRequest() (request *DescribeAffectedClusterCountRequest) {
    request = &DescribeAffectedClusterCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAffectedClusterCount")
    
    
    return
}

func NewDescribeAffectedClusterCountResponse() (response *DescribeAffectedClusterCountResponse) {
    response = &DescribeAffectedClusterCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAffectedClusterCount
// This API is used to get and return the number of affected clusters.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAffectedClusterCount(request *DescribeAffectedClusterCountRequest) (response *DescribeAffectedClusterCountResponse, err error) {
    return c.DescribeAffectedClusterCountWithContext(context.Background(), request)
}

// DescribeAffectedClusterCount
// This API is used to get and return the number of affected clusters.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAffectedClusterCountWithContext(ctx context.Context, request *DescribeAffectedClusterCountRequest) (response *DescribeAffectedClusterCountResponse, err error) {
    if request == nil {
        request = NewDescribeAffectedClusterCountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAffectedClusterCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAffectedClusterCountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAffectedNodeListRequest() (request *DescribeAffectedNodeListRequest) {
    request = &DescribeAffectedNodeListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAffectedNodeList")
    
    
    return
}

func NewDescribeAffectedNodeListResponse() (response *DescribeAffectedNodeListResponse) {
    response = &DescribeAffectedNodeListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAffectedNodeList
// This API is used to query affected node types and return the node list.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAffectedNodeList(request *DescribeAffectedNodeListRequest) (response *DescribeAffectedNodeListResponse, err error) {
    return c.DescribeAffectedNodeListWithContext(context.Background(), request)
}

// DescribeAffectedNodeList
// This API is used to query affected node types and return the node list.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAffectedNodeListWithContext(ctx context.Context, request *DescribeAffectedNodeListRequest) (response *DescribeAffectedNodeListResponse, err error) {
    if request == nil {
        request = NewDescribeAffectedNodeListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAffectedNodeList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAffectedNodeListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAffectedWorkloadListRequest() (request *DescribeAffectedWorkloadListRequest) {
    request = &DescribeAffectedWorkloadListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAffectedWorkloadList")
    
    
    return
}

func NewDescribeAffectedWorkloadListResponse() (response *DescribeAffectedWorkloadListResponse) {
    response = &DescribeAffectedWorkloadListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAffectedWorkloadList
// This API is used to query affected workload types and return the workload list.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAffectedWorkloadList(request *DescribeAffectedWorkloadListRequest) (response *DescribeAffectedWorkloadListResponse, err error) {
    return c.DescribeAffectedWorkloadListWithContext(context.Background(), request)
}

// DescribeAffectedWorkloadList
// This API is used to query affected workload types and return the workload list.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAffectedWorkloadListWithContext(ctx context.Context, request *DescribeAffectedWorkloadListRequest) (response *DescribeAffectedWorkloadListResponse, err error) {
    if request == nil {
        request = NewDescribeAffectedWorkloadListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAffectedWorkloadList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAffectedWorkloadListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentDaemonSetCmdRequest() (request *DescribeAgentDaemonSetCmdRequest) {
    request = &DescribeAgentDaemonSetCmdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAgentDaemonSetCmd")
    
    
    return
}

func NewDescribeAgentDaemonSetCmdResponse() (response *DescribeAgentDaemonSetCmdResponse) {
    response = &DescribeAgentDaemonSetCmdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAgentDaemonSetCmd
// This API is used to query parallel container installation commands.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeAgentDaemonSetCmd(request *DescribeAgentDaemonSetCmdRequest) (response *DescribeAgentDaemonSetCmdResponse, err error) {
    return c.DescribeAgentDaemonSetCmdWithContext(context.Background(), request)
}

// DescribeAgentDaemonSetCmd
// This API is used to query parallel container installation commands.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeAgentDaemonSetCmdWithContext(ctx context.Context, request *DescribeAgentDaemonSetCmdRequest) (response *DescribeAgentDaemonSetCmdResponse, err error) {
    if request == nil {
        request = NewDescribeAgentDaemonSetCmdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgentDaemonSetCmd require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgentDaemonSetCmdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentInstallCommandRequest() (request *DescribeAgentInstallCommandRequest) {
    request = &DescribeAgentInstallCommandRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAgentInstallCommand")
    
    
    return
}

func NewDescribeAgentInstallCommandResponse() (response *DescribeAgentInstallCommandResponse) {
    response = &DescribeAgentInstallCommandResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAgentInstallCommand
// This API is used to query agent installation commands.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeAgentInstallCommand(request *DescribeAgentInstallCommandRequest) (response *DescribeAgentInstallCommandResponse, err error) {
    return c.DescribeAgentInstallCommandWithContext(context.Background(), request)
}

// DescribeAgentInstallCommand
// This API is used to query agent installation commands.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeAgentInstallCommandWithContext(ctx context.Context, request *DescribeAgentInstallCommandRequest) (response *DescribeAgentInstallCommandResponse, err error) {
    if request == nil {
        request = NewDescribeAgentInstallCommandRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgentInstallCommand require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgentInstallCommandResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetAppServiceListRequest() (request *DescribeAssetAppServiceListRequest) {
    request = &DescribeAssetAppServiceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetAppServiceList")
    
    
    return
}

func NewDescribeAssetAppServiceListResponse() (response *DescribeAssetAppServiceListResponse) {
    response = &DescribeAssetAppServiceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetAppServiceList
// This API is used to query the list of application services.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetAppServiceList(request *DescribeAssetAppServiceListRequest) (response *DescribeAssetAppServiceListResponse, err error) {
    return c.DescribeAssetAppServiceListWithContext(context.Background(), request)
}

// DescribeAssetAppServiceList
// This API is used to query the list of application services.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetAppServiceListWithContext(ctx context.Context, request *DescribeAssetAppServiceListRequest) (response *DescribeAssetAppServiceListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetAppServiceListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetAppServiceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetAppServiceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetClusterListRequest() (request *DescribeAssetClusterListRequest) {
    request = &DescribeAssetClusterListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetClusterList")
    
    
    return
}

func NewDescribeAssetClusterListResponse() (response *DescribeAssetClusterListResponse) {
    response = &DescribeAssetClusterListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetClusterList
// This API is used to query the list of clusters.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAssetClusterList(request *DescribeAssetClusterListRequest) (response *DescribeAssetClusterListResponse, err error) {
    return c.DescribeAssetClusterListWithContext(context.Background(), request)
}

// DescribeAssetClusterList
// This API is used to query the list of clusters.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAssetClusterListWithContext(ctx context.Context, request *DescribeAssetClusterListRequest) (response *DescribeAssetClusterListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetClusterListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetClusterList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetClusterListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetComponentListRequest() (request *DescribeAssetComponentListRequest) {
    request = &DescribeAssetComponentListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetComponentList")
    
    
    return
}

func NewDescribeAssetComponentListResponse() (response *DescribeAssetComponentListResponse) {
    response = &DescribeAssetComponentListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetComponentList
// This API is used to query the list of container components.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetComponentList(request *DescribeAssetComponentListRequest) (response *DescribeAssetComponentListResponse, err error) {
    return c.DescribeAssetComponentListWithContext(context.Background(), request)
}

// DescribeAssetComponentList
// This API is used to query the list of container components.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetComponentListWithContext(ctx context.Context, request *DescribeAssetComponentListRequest) (response *DescribeAssetComponentListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetComponentListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetComponentList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetComponentListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetContainerDetailRequest() (request *DescribeAssetContainerDetailRequest) {
    request = &DescribeAssetContainerDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetContainerDetail")
    
    
    return
}

func NewDescribeAssetContainerDetailResponse() (response *DescribeAssetContainerDetailResponse) {
    response = &DescribeAssetContainerDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetContainerDetail
// This API is used to query the information of a container.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetContainerDetail(request *DescribeAssetContainerDetailRequest) (response *DescribeAssetContainerDetailResponse, err error) {
    return c.DescribeAssetContainerDetailWithContext(context.Background(), request)
}

// DescribeAssetContainerDetail
// This API is used to query the information of a container.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetContainerDetailWithContext(ctx context.Context, request *DescribeAssetContainerDetailRequest) (response *DescribeAssetContainerDetailResponse, err error) {
    if request == nil {
        request = NewDescribeAssetContainerDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetContainerDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetContainerDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetContainerListRequest() (request *DescribeAssetContainerListRequest) {
    request = &DescribeAssetContainerListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetContainerList")
    
    
    return
}

func NewDescribeAssetContainerListResponse() (response *DescribeAssetContainerListResponse) {
    response = &DescribeAssetContainerListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetContainerList
// This API is used to query the list of containers.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetContainerList(request *DescribeAssetContainerListRequest) (response *DescribeAssetContainerListResponse, err error) {
    return c.DescribeAssetContainerListWithContext(context.Background(), request)
}

// DescribeAssetContainerList
// This API is used to query the list of containers.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetContainerListWithContext(ctx context.Context, request *DescribeAssetContainerListRequest) (response *DescribeAssetContainerListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetContainerListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetContainerList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetContainerListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetDBServiceListRequest() (request *DescribeAssetDBServiceListRequest) {
    request = &DescribeAssetDBServiceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetDBServiceList")
    
    
    return
}

func NewDescribeAssetDBServiceListResponse() (response *DescribeAssetDBServiceListResponse) {
    response = &DescribeAssetDBServiceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetDBServiceList
// This API is used to query the list of database services.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetDBServiceList(request *DescribeAssetDBServiceListRequest) (response *DescribeAssetDBServiceListResponse, err error) {
    return c.DescribeAssetDBServiceListWithContext(context.Background(), request)
}

// DescribeAssetDBServiceList
// This API is used to query the list of database services.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetDBServiceListWithContext(ctx context.Context, request *DescribeAssetDBServiceListRequest) (response *DescribeAssetDBServiceListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetDBServiceListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetDBServiceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetDBServiceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetHostDetailRequest() (request *DescribeAssetHostDetailRequest) {
    request = &DescribeAssetHostDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetHostDetail")
    
    
    return
}

func NewDescribeAssetHostDetailResponse() (response *DescribeAssetHostDetailResponse) {
    response = &DescribeAssetHostDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetHostDetail
// This API is used to query the details of a server.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetHostDetail(request *DescribeAssetHostDetailRequest) (response *DescribeAssetHostDetailResponse, err error) {
    return c.DescribeAssetHostDetailWithContext(context.Background(), request)
}

// DescribeAssetHostDetail
// This API is used to query the details of a server.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetHostDetailWithContext(ctx context.Context, request *DescribeAssetHostDetailRequest) (response *DescribeAssetHostDetailResponse, err error) {
    if request == nil {
        request = NewDescribeAssetHostDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetHostDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetHostDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetHostListRequest() (request *DescribeAssetHostListRequest) {
    request = &DescribeAssetHostListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetHostList")
    
    
    return
}

func NewDescribeAssetHostListResponse() (response *DescribeAssetHostListResponse) {
    response = &DescribeAssetHostListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetHostList
// This API is used to query the list of servers.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetHostList(request *DescribeAssetHostListRequest) (response *DescribeAssetHostListResponse, err error) {
    return c.DescribeAssetHostListWithContext(context.Background(), request)
}

// DescribeAssetHostList
// This API is used to query the list of servers.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetHostListWithContext(ctx context.Context, request *DescribeAssetHostListRequest) (response *DescribeAssetHostListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetHostListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetHostList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetHostListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageBindRuleInfoRequest() (request *DescribeAssetImageBindRuleInfoRequest) {
    request = &DescribeAssetImageBindRuleInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageBindRuleInfo")
    
    
    return
}

func NewDescribeAssetImageBindRuleInfoResponse() (response *DescribeAssetImageBindRuleInfoResponse) {
    response = &DescribeAssetImageBindRuleInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetImageBindRuleInfo
// This API is used to query the list of rules bound to images, including runtime access control and abnormal process rules.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageBindRuleInfo(request *DescribeAssetImageBindRuleInfoRequest) (response *DescribeAssetImageBindRuleInfoResponse, err error) {
    return c.DescribeAssetImageBindRuleInfoWithContext(context.Background(), request)
}

// DescribeAssetImageBindRuleInfo
// This API is used to query the list of rules bound to images, including runtime access control and abnormal process rules.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageBindRuleInfoWithContext(ctx context.Context, request *DescribeAssetImageBindRuleInfoRequest) (response *DescribeAssetImageBindRuleInfoResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageBindRuleInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageBindRuleInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageBindRuleInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageDetailRequest() (request *DescribeAssetImageDetailRequest) {
    request = &DescribeAssetImageDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageDetail")
    
    
    return
}

func NewDescribeAssetImageDetailResponse() (response *DescribeAssetImageDetailResponse) {
    response = &DescribeAssetImageDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetImageDetail
// This API is used to query the details of an image.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageDetail(request *DescribeAssetImageDetailRequest) (response *DescribeAssetImageDetailResponse, err error) {
    return c.DescribeAssetImageDetailWithContext(context.Background(), request)
}

// DescribeAssetImageDetail
// This API is used to query the details of an image.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageDetailWithContext(ctx context.Context, request *DescribeAssetImageDetailRequest) (response *DescribeAssetImageDetailResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageHostListRequest() (request *DescribeAssetImageHostListRequest) {
    request = &DescribeAssetImageHostListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageHostList")
    
    
    return
}

func NewDescribeAssetImageHostListResponse() (response *DescribeAssetImageHostListResponse) {
    response = &DescribeAssetImageHostListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetImageHostList
// This API is used to query the servers associated with an image.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageHostList(request *DescribeAssetImageHostListRequest) (response *DescribeAssetImageHostListResponse, err error) {
    return c.DescribeAssetImageHostListWithContext(context.Background(), request)
}

// DescribeAssetImageHostList
// This API is used to query the servers associated with an image.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageHostListWithContext(ctx context.Context, request *DescribeAssetImageHostListRequest) (response *DescribeAssetImageHostListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageHostListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageHostList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageHostListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageListRequest() (request *DescribeAssetImageListRequest) {
    request = &DescribeAssetImageListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageList")
    
    
    return
}

func NewDescribeAssetImageListResponse() (response *DescribeAssetImageListResponse) {
    response = &DescribeAssetImageListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetImageList
// This API is used to query the list of images.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageList(request *DescribeAssetImageListRequest) (response *DescribeAssetImageListResponse, err error) {
    return c.DescribeAssetImageListWithContext(context.Background(), request)
}

// DescribeAssetImageList
// This API is used to query the list of images.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageListWithContext(ctx context.Context, request *DescribeAssetImageListRequest) (response *DescribeAssetImageListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageListExportRequest() (request *DescribeAssetImageListExportRequest) {
    request = &DescribeAssetImageListExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageListExport")
    
    
    return
}

func NewDescribeAssetImageListExportResponse() (response *DescribeAssetImageListExportResponse) {
    response = &DescribeAssetImageListExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetImageListExport
// 接口已废弃
//
// 
//
// This API is used to export the list of images.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageListExport(request *DescribeAssetImageListExportRequest) (response *DescribeAssetImageListExportResponse, err error) {
    return c.DescribeAssetImageListExportWithContext(context.Background(), request)
}

// DescribeAssetImageListExport
// 接口已废弃
//
// 
//
// This API is used to export the list of images.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageListExportWithContext(ctx context.Context, request *DescribeAssetImageListExportRequest) (response *DescribeAssetImageListExportResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageListExportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageListExport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageListExportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageRegistryAssetStatusRequest() (request *DescribeAssetImageRegistryAssetStatusRequest) {
    request = &DescribeAssetImageRegistryAssetStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRegistryAssetStatus")
    
    
    return
}

func NewDescribeAssetImageRegistryAssetStatusResponse() (response *DescribeAssetImageRegistryAssetStatusResponse) {
    response = &DescribeAssetImageRegistryAssetStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetImageRegistryAssetStatus
// This API is used to view the update progress of the assets in an image repository.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryAssetStatus(request *DescribeAssetImageRegistryAssetStatusRequest) (response *DescribeAssetImageRegistryAssetStatusResponse, err error) {
    return c.DescribeAssetImageRegistryAssetStatusWithContext(context.Background(), request)
}

// DescribeAssetImageRegistryAssetStatus
// This API is used to view the update progress of the assets in an image repository.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryAssetStatusWithContext(ctx context.Context, request *DescribeAssetImageRegistryAssetStatusRequest) (response *DescribeAssetImageRegistryAssetStatusResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageRegistryAssetStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageRegistryAssetStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageRegistryAssetStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageRegistryDetailRequest() (request *DescribeAssetImageRegistryDetailRequest) {
    request = &DescribeAssetImageRegistryDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRegistryDetail")
    
    
    return
}

func NewDescribeAssetImageRegistryDetailResponse() (response *DescribeAssetImageRegistryDetailResponse) {
    response = &DescribeAssetImageRegistryDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetImageRegistryDetail
// This API is used to query the image repository details.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryDetail(request *DescribeAssetImageRegistryDetailRequest) (response *DescribeAssetImageRegistryDetailResponse, err error) {
    return c.DescribeAssetImageRegistryDetailWithContext(context.Background(), request)
}

// DescribeAssetImageRegistryDetail
// This API is used to query the image repository details.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryDetailWithContext(ctx context.Context, request *DescribeAssetImageRegistryDetailRequest) (response *DescribeAssetImageRegistryDetailResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageRegistryDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageRegistryDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageRegistryDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageRegistryListRequest() (request *DescribeAssetImageRegistryListRequest) {
    request = &DescribeAssetImageRegistryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRegistryList")
    
    
    return
}

func NewDescribeAssetImageRegistryListResponse() (response *DescribeAssetImageRegistryListResponse) {
    response = &DescribeAssetImageRegistryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetImageRegistryList
// This API is used to query the list of image repositories.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryList(request *DescribeAssetImageRegistryListRequest) (response *DescribeAssetImageRegistryListResponse, err error) {
    return c.DescribeAssetImageRegistryListWithContext(context.Background(), request)
}

// DescribeAssetImageRegistryList
// This API is used to query the list of image repositories.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryListWithContext(ctx context.Context, request *DescribeAssetImageRegistryListRequest) (response *DescribeAssetImageRegistryListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageRegistryListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageRegistryList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageRegistryListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageRegistryListExportRequest() (request *DescribeAssetImageRegistryListExportRequest) {
    request = &DescribeAssetImageRegistryListExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRegistryListExport")
    
    
    return
}

func NewDescribeAssetImageRegistryListExportResponse() (response *DescribeAssetImageRegistryListExportResponse) {
    response = &DescribeAssetImageRegistryListExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetImageRegistryListExport
// This API is used to export the list of images for an image repository.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryListExport(request *DescribeAssetImageRegistryListExportRequest) (response *DescribeAssetImageRegistryListExportResponse, err error) {
    return c.DescribeAssetImageRegistryListExportWithContext(context.Background(), request)
}

// DescribeAssetImageRegistryListExport
// This API is used to export the list of images for an image repository.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryListExportWithContext(ctx context.Context, request *DescribeAssetImageRegistryListExportRequest) (response *DescribeAssetImageRegistryListExportResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageRegistryListExportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageRegistryListExport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageRegistryListExportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageRegistryRegistryDetailRequest() (request *DescribeAssetImageRegistryRegistryDetailRequest) {
    request = &DescribeAssetImageRegistryRegistryDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRegistryRegistryDetail")
    
    
    return
}

func NewDescribeAssetImageRegistryRegistryDetailResponse() (response *DescribeAssetImageRegistryRegistryDetailResponse) {
    response = &DescribeAssetImageRegistryRegistryDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetImageRegistryRegistryDetail
// This API is used to view the details of an image repository.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryRegistryDetail(request *DescribeAssetImageRegistryRegistryDetailRequest) (response *DescribeAssetImageRegistryRegistryDetailResponse, err error) {
    return c.DescribeAssetImageRegistryRegistryDetailWithContext(context.Background(), request)
}

// DescribeAssetImageRegistryRegistryDetail
// This API is used to view the details of an image repository.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryRegistryDetailWithContext(ctx context.Context, request *DescribeAssetImageRegistryRegistryDetailRequest) (response *DescribeAssetImageRegistryRegistryDetailResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageRegistryRegistryDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageRegistryRegistryDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageRegistryRegistryDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageRegistryRegistryListRequest() (request *DescribeAssetImageRegistryRegistryListRequest) {
    request = &DescribeAssetImageRegistryRegistryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRegistryRegistryList")
    
    
    return
}

func NewDescribeAssetImageRegistryRegistryListResponse() (response *DescribeAssetImageRegistryRegistryListResponse) {
    response = &DescribeAssetImageRegistryRegistryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetImageRegistryRegistryList
// This API is used to query the list of image registries.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryRegistryList(request *DescribeAssetImageRegistryRegistryListRequest) (response *DescribeAssetImageRegistryRegistryListResponse, err error) {
    return c.DescribeAssetImageRegistryRegistryListWithContext(context.Background(), request)
}

// DescribeAssetImageRegistryRegistryList
// This API is used to query the list of image registries.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryRegistryListWithContext(ctx context.Context, request *DescribeAssetImageRegistryRegistryListRequest) (response *DescribeAssetImageRegistryRegistryListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageRegistryRegistryListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageRegistryRegistryList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageRegistryRegistryListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageRegistryRiskInfoListRequest() (request *DescribeAssetImageRegistryRiskInfoListRequest) {
    request = &DescribeAssetImageRegistryRiskInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRegistryRiskInfoList")
    
    
    return
}

func NewDescribeAssetImageRegistryRiskInfoListResponse() (response *DescribeAssetImageRegistryRiskInfoListResponse) {
    response = &DescribeAssetImageRegistryRiskInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetImageRegistryRiskInfoList
// This API is used to query the list of image high-risk behaviors of an image repository.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryRiskInfoList(request *DescribeAssetImageRegistryRiskInfoListRequest) (response *DescribeAssetImageRegistryRiskInfoListResponse, err error) {
    return c.DescribeAssetImageRegistryRiskInfoListWithContext(context.Background(), request)
}

// DescribeAssetImageRegistryRiskInfoList
// This API is used to query the list of image high-risk behaviors of an image repository.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryRiskInfoListWithContext(ctx context.Context, request *DescribeAssetImageRegistryRiskInfoListRequest) (response *DescribeAssetImageRegistryRiskInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageRegistryRiskInfoListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageRegistryRiskInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageRegistryRiskInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageRegistryRiskListExportRequest() (request *DescribeAssetImageRegistryRiskListExportRequest) {
    request = &DescribeAssetImageRegistryRiskListExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRegistryRiskListExport")
    
    
    return
}

func NewDescribeAssetImageRegistryRiskListExportResponse() (response *DescribeAssetImageRegistryRiskListExportResponse) {
    response = &DescribeAssetImageRegistryRiskListExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetImageRegistryRiskListExport
// This API is used to export the list of sensitive data for an image repository.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryRiskListExport(request *DescribeAssetImageRegistryRiskListExportRequest) (response *DescribeAssetImageRegistryRiskListExportResponse, err error) {
    return c.DescribeAssetImageRegistryRiskListExportWithContext(context.Background(), request)
}

// DescribeAssetImageRegistryRiskListExport
// This API is used to export the list of sensitive data for an image repository.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryRiskListExportWithContext(ctx context.Context, request *DescribeAssetImageRegistryRiskListExportRequest) (response *DescribeAssetImageRegistryRiskListExportResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageRegistryRiskListExportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageRegistryRiskListExport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageRegistryRiskListExportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageRegistryScanStatusOneKeyRequest() (request *DescribeAssetImageRegistryScanStatusOneKeyRequest) {
    request = &DescribeAssetImageRegistryScanStatusOneKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRegistryScanStatusOneKey")
    
    
    return
}

func NewDescribeAssetImageRegistryScanStatusOneKeyResponse() (response *DescribeAssetImageRegistryScanStatusOneKeyResponse) {
    response = &DescribeAssetImageRegistryScanStatusOneKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetImageRegistryScanStatusOneKey
// This API is used to query the quick image scanning status of an image repository.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryScanStatusOneKey(request *DescribeAssetImageRegistryScanStatusOneKeyRequest) (response *DescribeAssetImageRegistryScanStatusOneKeyResponse, err error) {
    return c.DescribeAssetImageRegistryScanStatusOneKeyWithContext(context.Background(), request)
}

// DescribeAssetImageRegistryScanStatusOneKey
// This API is used to query the quick image scanning status of an image repository.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryScanStatusOneKeyWithContext(ctx context.Context, request *DescribeAssetImageRegistryScanStatusOneKeyRequest) (response *DescribeAssetImageRegistryScanStatusOneKeyResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageRegistryScanStatusOneKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageRegistryScanStatusOneKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageRegistryScanStatusOneKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageRegistrySummaryRequest() (request *DescribeAssetImageRegistrySummaryRequest) {
    request = &DescribeAssetImageRegistrySummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRegistrySummary")
    
    
    return
}

func NewDescribeAssetImageRegistrySummaryResponse() (response *DescribeAssetImageRegistrySummaryResponse) {
    response = &DescribeAssetImageRegistrySummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetImageRegistrySummary
// This API is used to query the image statistics of an image repository.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistrySummary(request *DescribeAssetImageRegistrySummaryRequest) (response *DescribeAssetImageRegistrySummaryResponse, err error) {
    return c.DescribeAssetImageRegistrySummaryWithContext(context.Background(), request)
}

// DescribeAssetImageRegistrySummary
// This API is used to query the image statistics of an image repository.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistrySummaryWithContext(ctx context.Context, request *DescribeAssetImageRegistrySummaryRequest) (response *DescribeAssetImageRegistrySummaryResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageRegistrySummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageRegistrySummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageRegistrySummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageRegistryVirusListRequest() (request *DescribeAssetImageRegistryVirusListRequest) {
    request = &DescribeAssetImageRegistryVirusListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRegistryVirusList")
    
    
    return
}

func NewDescribeAssetImageRegistryVirusListResponse() (response *DescribeAssetImageRegistryVirusListResponse) {
    response = &DescribeAssetImageRegistryVirusListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetImageRegistryVirusList
// This API is used to query the list of viruses and trojans in an image repository.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryVirusList(request *DescribeAssetImageRegistryVirusListRequest) (response *DescribeAssetImageRegistryVirusListResponse, err error) {
    return c.DescribeAssetImageRegistryVirusListWithContext(context.Background(), request)
}

// DescribeAssetImageRegistryVirusList
// This API is used to query the list of viruses and trojans in an image repository.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryVirusListWithContext(ctx context.Context, request *DescribeAssetImageRegistryVirusListRequest) (response *DescribeAssetImageRegistryVirusListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageRegistryVirusListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageRegistryVirusList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageRegistryVirusListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageRegistryVirusListExportRequest() (request *DescribeAssetImageRegistryVirusListExportRequest) {
    request = &DescribeAssetImageRegistryVirusListExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRegistryVirusListExport")
    
    
    return
}

func NewDescribeAssetImageRegistryVirusListExportResponse() (response *DescribeAssetImageRegistryVirusListExportResponse) {
    response = &DescribeAssetImageRegistryVirusListExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetImageRegistryVirusListExport
// This API is used to export the list of trojan information for an image repository.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryVirusListExport(request *DescribeAssetImageRegistryVirusListExportRequest) (response *DescribeAssetImageRegistryVirusListExportResponse, err error) {
    return c.DescribeAssetImageRegistryVirusListExportWithContext(context.Background(), request)
}

// DescribeAssetImageRegistryVirusListExport
// This API is used to export the list of trojan information for an image repository.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryVirusListExportWithContext(ctx context.Context, request *DescribeAssetImageRegistryVirusListExportRequest) (response *DescribeAssetImageRegistryVirusListExportResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageRegistryVirusListExportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageRegistryVirusListExport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageRegistryVirusListExportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageRegistryVulListRequest() (request *DescribeAssetImageRegistryVulListRequest) {
    request = &DescribeAssetImageRegistryVulListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRegistryVulList")
    
    
    return
}

func NewDescribeAssetImageRegistryVulListResponse() (response *DescribeAssetImageRegistryVulListResponse) {
    response = &DescribeAssetImageRegistryVulListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetImageRegistryVulList
// This API is used to query the list of image vulnerabilities of an image repository
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryVulList(request *DescribeAssetImageRegistryVulListRequest) (response *DescribeAssetImageRegistryVulListResponse, err error) {
    return c.DescribeAssetImageRegistryVulListWithContext(context.Background(), request)
}

// DescribeAssetImageRegistryVulList
// This API is used to query the list of image vulnerabilities of an image repository
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryVulListWithContext(ctx context.Context, request *DescribeAssetImageRegistryVulListRequest) (response *DescribeAssetImageRegistryVulListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageRegistryVulListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageRegistryVulList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageRegistryVulListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageRegistryVulListExportRequest() (request *DescribeAssetImageRegistryVulListExportRequest) {
    request = &DescribeAssetImageRegistryVulListExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRegistryVulListExport")
    
    
    return
}

func NewDescribeAssetImageRegistryVulListExportResponse() (response *DescribeAssetImageRegistryVulListExportResponse) {
    response = &DescribeAssetImageRegistryVulListExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetImageRegistryVulListExport
// This API is used to export the list of vulnerabilities for an image repository.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryVulListExport(request *DescribeAssetImageRegistryVulListExportRequest) (response *DescribeAssetImageRegistryVulListExportResponse, err error) {
    return c.DescribeAssetImageRegistryVulListExportWithContext(context.Background(), request)
}

// DescribeAssetImageRegistryVulListExport
// This API is used to export the list of vulnerabilities for an image repository.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryVulListExportWithContext(ctx context.Context, request *DescribeAssetImageRegistryVulListExportRequest) (response *DescribeAssetImageRegistryVulListExportResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageRegistryVulListExportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageRegistryVulListExport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageRegistryVulListExportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageRiskListRequest() (request *DescribeAssetImageRiskListRequest) {
    request = &DescribeAssetImageRiskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRiskList")
    
    
    return
}

func NewDescribeAssetImageRiskListResponse() (response *DescribeAssetImageRiskListResponse) {
    response = &DescribeAssetImageRiskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetImageRiskList
// This API is used to query the list of risks in an image.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRiskList(request *DescribeAssetImageRiskListRequest) (response *DescribeAssetImageRiskListResponse, err error) {
    return c.DescribeAssetImageRiskListWithContext(context.Background(), request)
}

// DescribeAssetImageRiskList
// This API is used to query the list of risks in an image.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRiskListWithContext(ctx context.Context, request *DescribeAssetImageRiskListRequest) (response *DescribeAssetImageRiskListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageRiskListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageRiskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageRiskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageRiskListExportRequest() (request *DescribeAssetImageRiskListExportRequest) {
    request = &DescribeAssetImageRiskListExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRiskListExport")
    
    
    return
}

func NewDescribeAssetImageRiskListExportResponse() (response *DescribeAssetImageRiskListExportResponse) {
    response = &DescribeAssetImageRiskListExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetImageRiskListExport
// This API is used to export the list of risks in an image.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRiskListExport(request *DescribeAssetImageRiskListExportRequest) (response *DescribeAssetImageRiskListExportResponse, err error) {
    return c.DescribeAssetImageRiskListExportWithContext(context.Background(), request)
}

// DescribeAssetImageRiskListExport
// This API is used to export the list of risks in an image.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRiskListExportWithContext(ctx context.Context, request *DescribeAssetImageRiskListExportRequest) (response *DescribeAssetImageRiskListExportResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageRiskListExportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageRiskListExport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageRiskListExportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageScanSettingRequest() (request *DescribeAssetImageScanSettingRequest) {
    request = &DescribeAssetImageScanSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageScanSetting")
    
    
    return
}

func NewDescribeAssetImageScanSettingResponse() (response *DescribeAssetImageScanSettingResponse) {
    response = &DescribeAssetImageScanSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetImageScanSetting
// This API is used to get the image scan settings.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
func (c *Client) DescribeAssetImageScanSetting(request *DescribeAssetImageScanSettingRequest) (response *DescribeAssetImageScanSettingResponse, err error) {
    return c.DescribeAssetImageScanSettingWithContext(context.Background(), request)
}

// DescribeAssetImageScanSetting
// This API is used to get the image scan settings.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
func (c *Client) DescribeAssetImageScanSettingWithContext(ctx context.Context, request *DescribeAssetImageScanSettingRequest) (response *DescribeAssetImageScanSettingResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageScanSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageScanSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageScanSettingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageScanStatusRequest() (request *DescribeAssetImageScanStatusRequest) {
    request = &DescribeAssetImageScanStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageScanStatus")
    
    
    return
}

func NewDescribeAssetImageScanStatusResponse() (response *DescribeAssetImageScanStatusResponse) {
    response = &DescribeAssetImageScanStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetImageScanStatus
// This API is used to query the image scanning status.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageScanStatus(request *DescribeAssetImageScanStatusRequest) (response *DescribeAssetImageScanStatusResponse, err error) {
    return c.DescribeAssetImageScanStatusWithContext(context.Background(), request)
}

// DescribeAssetImageScanStatus
// This API is used to query the image scanning status.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageScanStatusWithContext(ctx context.Context, request *DescribeAssetImageScanStatusRequest) (response *DescribeAssetImageScanStatusResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageScanStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageScanStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageScanStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageScanTaskRequest() (request *DescribeAssetImageScanTaskRequest) {
    request = &DescribeAssetImageScanTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageScanTask")
    
    
    return
}

func NewDescribeAssetImageScanTaskResponse() (response *DescribeAssetImageScanTaskResponse) {
    response = &DescribeAssetImageScanTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetImageScanTask
// This API is used to query the ID of a running quick image scan task.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageScanTask(request *DescribeAssetImageScanTaskRequest) (response *DescribeAssetImageScanTaskResponse, err error) {
    return c.DescribeAssetImageScanTaskWithContext(context.Background(), request)
}

// DescribeAssetImageScanTask
// This API is used to query the ID of a running quick image scan task.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageScanTaskWithContext(ctx context.Context, request *DescribeAssetImageScanTaskRequest) (response *DescribeAssetImageScanTaskResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageScanTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageScanTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageScanTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageSimpleListRequest() (request *DescribeAssetImageSimpleListRequest) {
    request = &DescribeAssetImageSimpleListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageSimpleList")
    
    
    return
}

func NewDescribeAssetImageSimpleListResponse() (response *DescribeAssetImageSimpleListResponse) {
    response = &DescribeAssetImageSimpleListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetImageSimpleList
// This API is used to query the brief information list of an image.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageSimpleList(request *DescribeAssetImageSimpleListRequest) (response *DescribeAssetImageSimpleListResponse, err error) {
    return c.DescribeAssetImageSimpleListWithContext(context.Background(), request)
}

// DescribeAssetImageSimpleList
// This API is used to query the brief information list of an image.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageSimpleListWithContext(ctx context.Context, request *DescribeAssetImageSimpleListRequest) (response *DescribeAssetImageSimpleListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageSimpleListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageSimpleList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageSimpleListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageVirusListRequest() (request *DescribeAssetImageVirusListRequest) {
    request = &DescribeAssetImageVirusListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageVirusList")
    
    
    return
}

func NewDescribeAssetImageVirusListResponse() (response *DescribeAssetImageVirusListResponse) {
    response = &DescribeAssetImageVirusListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetImageVirusList
// This API is used to query the list of viruses in an image.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageVirusList(request *DescribeAssetImageVirusListRequest) (response *DescribeAssetImageVirusListResponse, err error) {
    return c.DescribeAssetImageVirusListWithContext(context.Background(), request)
}

// DescribeAssetImageVirusList
// This API is used to query the list of viruses in an image.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageVirusListWithContext(ctx context.Context, request *DescribeAssetImageVirusListRequest) (response *DescribeAssetImageVirusListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageVirusListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageVirusList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageVirusListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageVirusListExportRequest() (request *DescribeAssetImageVirusListExportRequest) {
    request = &DescribeAssetImageVirusListExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageVirusListExport")
    
    
    return
}

func NewDescribeAssetImageVirusListExportResponse() (response *DescribeAssetImageVirusListExportResponse) {
    response = &DescribeAssetImageVirusListExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetImageVirusListExport
// This API is used to export the list of trojans in an image.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageVirusListExport(request *DescribeAssetImageVirusListExportRequest) (response *DescribeAssetImageVirusListExportResponse, err error) {
    return c.DescribeAssetImageVirusListExportWithContext(context.Background(), request)
}

// DescribeAssetImageVirusListExport
// This API is used to export the list of trojans in an image.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageVirusListExportWithContext(ctx context.Context, request *DescribeAssetImageVirusListExportRequest) (response *DescribeAssetImageVirusListExportResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageVirusListExportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageVirusListExport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageVirusListExportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageVulListRequest() (request *DescribeAssetImageVulListRequest) {
    request = &DescribeAssetImageVulListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageVulList")
    
    
    return
}

func NewDescribeAssetImageVulListResponse() (response *DescribeAssetImageVulListResponse) {
    response = &DescribeAssetImageVulListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetImageVulList
// This API is used to query the list of vulnerabilities in an image.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageVulList(request *DescribeAssetImageVulListRequest) (response *DescribeAssetImageVulListResponse, err error) {
    return c.DescribeAssetImageVulListWithContext(context.Background(), request)
}

// DescribeAssetImageVulList
// This API is used to query the list of vulnerabilities in an image.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageVulListWithContext(ctx context.Context, request *DescribeAssetImageVulListRequest) (response *DescribeAssetImageVulListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageVulListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageVulList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageVulListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageVulListExportRequest() (request *DescribeAssetImageVulListExportRequest) {
    request = &DescribeAssetImageVulListExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageVulListExport")
    
    
    return
}

func NewDescribeAssetImageVulListExportResponse() (response *DescribeAssetImageVulListExportResponse) {
    response = &DescribeAssetImageVulListExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetImageVulListExport
// This API is used to export the list of vulnerabilities in an image.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageVulListExport(request *DescribeAssetImageVulListExportRequest) (response *DescribeAssetImageVulListExportResponse, err error) {
    return c.DescribeAssetImageVulListExportWithContext(context.Background(), request)
}

// DescribeAssetImageVulListExport
// This API is used to export the list of vulnerabilities in an image.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageVulListExportWithContext(ctx context.Context, request *DescribeAssetImageVulListExportRequest) (response *DescribeAssetImageVulListExportResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageVulListExportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageVulListExport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageVulListExportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetPortListRequest() (request *DescribeAssetPortListRequest) {
    request = &DescribeAssetPortListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetPortList")
    
    
    return
}

func NewDescribeAssetPortListResponse() (response *DescribeAssetPortListResponse) {
    response = &DescribeAssetPortListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetPortList
// This API is used to query the list of occupied ports.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetPortList(request *DescribeAssetPortListRequest) (response *DescribeAssetPortListResponse, err error) {
    return c.DescribeAssetPortListWithContext(context.Background(), request)
}

// DescribeAssetPortList
// This API is used to query the list of occupied ports.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetPortListWithContext(ctx context.Context, request *DescribeAssetPortListRequest) (response *DescribeAssetPortListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetPortListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetPortList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetPortListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetProcessListRequest() (request *DescribeAssetProcessListRequest) {
    request = &DescribeAssetProcessListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetProcessList")
    
    
    return
}

func NewDescribeAssetProcessListResponse() (response *DescribeAssetProcessListResponse) {
    response = &DescribeAssetProcessListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetProcessList
// This API is used to query the list of processes.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetProcessList(request *DescribeAssetProcessListRequest) (response *DescribeAssetProcessListResponse, err error) {
    return c.DescribeAssetProcessListWithContext(context.Background(), request)
}

// DescribeAssetProcessList
// This API is used to query the list of processes.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetProcessListWithContext(ctx context.Context, request *DescribeAssetProcessListRequest) (response *DescribeAssetProcessListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetProcessListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetProcessList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetProcessListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetSummaryRequest() (request *DescribeAssetSummaryRequest) {
    request = &DescribeAssetSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetSummary")
    
    
    return
}

func NewDescribeAssetSummaryResponse() (response *DescribeAssetSummaryResponse) {
    response = &DescribeAssetSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetSummary
// This API is used to query the statistics of containers and images under an account.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetSummary(request *DescribeAssetSummaryRequest) (response *DescribeAssetSummaryResponse, err error) {
    return c.DescribeAssetSummaryWithContext(context.Background(), request)
}

// DescribeAssetSummary
// This API is used to query the statistics of containers and images under an account.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetSummaryWithContext(ctx context.Context, request *DescribeAssetSummaryRequest) (response *DescribeAssetSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeAssetSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetSyncLastTimeRequest() (request *DescribeAssetSyncLastTimeRequest) {
    request = &DescribeAssetSyncLastTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetSyncLastTime")
    
    
    return
}

func NewDescribeAssetSyncLastTimeResponse() (response *DescribeAssetSyncLastTimeResponse) {
    response = &DescribeAssetSyncLastTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetSyncLastTime
// This API is used to query the last asset sync time.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DescribeAssetSyncLastTime(request *DescribeAssetSyncLastTimeRequest) (response *DescribeAssetSyncLastTimeResponse, err error) {
    return c.DescribeAssetSyncLastTimeWithContext(context.Background(), request)
}

// DescribeAssetSyncLastTime
// This API is used to query the last asset sync time.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DescribeAssetSyncLastTimeWithContext(ctx context.Context, request *DescribeAssetSyncLastTimeRequest) (response *DescribeAssetSyncLastTimeResponse, err error) {
    if request == nil {
        request = NewDescribeAssetSyncLastTimeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetSyncLastTime require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetSyncLastTimeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetWebServiceListRequest() (request *DescribeAssetWebServiceListRequest) {
    request = &DescribeAssetWebServiceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetWebServiceList")
    
    
    return
}

func NewDescribeAssetWebServiceListResponse() (response *DescribeAssetWebServiceListResponse) {
    response = &DescribeAssetWebServiceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetWebServiceList
// This API is used to query the list of web services.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetWebServiceList(request *DescribeAssetWebServiceListRequest) (response *DescribeAssetWebServiceListResponse, err error) {
    return c.DescribeAssetWebServiceListWithContext(context.Background(), request)
}

// DescribeAssetWebServiceList
// This API is used to query the list of web services.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetWebServiceListWithContext(ctx context.Context, request *DescribeAssetWebServiceListRequest) (response *DescribeAssetWebServiceListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetWebServiceListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetWebServiceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetWebServiceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAutoAuthorizedRuleHostRequest() (request *DescribeAutoAuthorizedRuleHostRequest) {
    request = &DescribeAutoAuthorizedRuleHostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAutoAuthorizedRuleHost")
    
    
    return
}

func NewDescribeAutoAuthorizedRuleHostResponse() (response *DescribeAutoAuthorizedRuleHostResponse) {
    response = &DescribeAutoAuthorizedRuleHostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAutoAuthorizedRuleHost
// This API is used to query the servers licensed according to an automatic licensing rule.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_ERRRULENOTFIND = "FailedOperation.ErrRuleNotFind"
//  FAILEDOPERATION_RULENOTFIND = "FailedOperation.RuleNotFind"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAutoAuthorizedRuleHost(request *DescribeAutoAuthorizedRuleHostRequest) (response *DescribeAutoAuthorizedRuleHostResponse, err error) {
    return c.DescribeAutoAuthorizedRuleHostWithContext(context.Background(), request)
}

// DescribeAutoAuthorizedRuleHost
// This API is used to query the servers licensed according to an automatic licensing rule.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_ERRRULENOTFIND = "FailedOperation.ErrRuleNotFind"
//  FAILEDOPERATION_RULENOTFIND = "FailedOperation.RuleNotFind"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAutoAuthorizedRuleHostWithContext(ctx context.Context, request *DescribeAutoAuthorizedRuleHostRequest) (response *DescribeAutoAuthorizedRuleHostResponse, err error) {
    if request == nil {
        request = NewDescribeAutoAuthorizedRuleHostRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAutoAuthorizedRuleHost require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAutoAuthorizedRuleHostResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCheckItemListRequest() (request *DescribeCheckItemListRequest) {
    request = &DescribeCheckItemListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeCheckItemList")
    
    
    return
}

func NewDescribeCheckItemListResponse() (response *DescribeCheckItemListResponse) {
    response = &DescribeCheckItemListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCheckItemList
// This API is used to query all check items and return the total number and list of check items.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeCheckItemList(request *DescribeCheckItemListRequest) (response *DescribeCheckItemListResponse, err error) {
    return c.DescribeCheckItemListWithContext(context.Background(), request)
}

// DescribeCheckItemList
// This API is used to query all check items and return the total number and list of check items.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeCheckItemListWithContext(ctx context.Context, request *DescribeCheckItemListRequest) (response *DescribeCheckItemListResponse, err error) {
    if request == nil {
        request = NewDescribeCheckItemListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCheckItemList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCheckItemListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterDetailRequest() (request *DescribeClusterDetailRequest) {
    request = &DescribeClusterDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeClusterDetail")
    
    
    return
}

func NewDescribeClusterDetailResponse() (response *DescribeClusterDetailResponse) {
    response = &DescribeClusterDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterDetail
// This API is used to query the details of a cluster.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeClusterDetail(request *DescribeClusterDetailRequest) (response *DescribeClusterDetailResponse, err error) {
    return c.DescribeClusterDetailWithContext(context.Background(), request)
}

// DescribeClusterDetail
// This API is used to query the details of a cluster.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeClusterDetailWithContext(ctx context.Context, request *DescribeClusterDetailRequest) (response *DescribeClusterDetailResponse, err error) {
    if request == nil {
        request = NewDescribeClusterDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterSummaryRequest() (request *DescribeClusterSummaryRequest) {
    request = &DescribeClusterSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeClusterSummary")
    
    
    return
}

func NewDescribeClusterSummaryResponse() (response *DescribeClusterSummaryResponse) {
    response = &DescribeClusterSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterSummary
// This API is used to query the overview of cluster assets.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeClusterSummary(request *DescribeClusterSummaryRequest) (response *DescribeClusterSummaryResponse, err error) {
    return c.DescribeClusterSummaryWithContext(context.Background(), request)
}

// DescribeClusterSummary
// This API is used to query the overview of cluster assets.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeClusterSummaryWithContext(ctx context.Context, request *DescribeClusterSummaryRequest) (response *DescribeClusterSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeClusterSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComplianceAssetDetailInfoRequest() (request *DescribeComplianceAssetDetailInfoRequest) {
    request = &DescribeComplianceAssetDetailInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeComplianceAssetDetailInfo")
    
    
    return
}

func NewDescribeComplianceAssetDetailInfoResponse() (response *DescribeComplianceAssetDetailInfoResponse) {
    response = &DescribeComplianceAssetDetailInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeComplianceAssetDetailInfo
// This API is used to query the details of an asset.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DescribeComplianceAssetDetailInfo(request *DescribeComplianceAssetDetailInfoRequest) (response *DescribeComplianceAssetDetailInfoResponse, err error) {
    return c.DescribeComplianceAssetDetailInfoWithContext(context.Background(), request)
}

// DescribeComplianceAssetDetailInfo
// This API is used to query the details of an asset.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DescribeComplianceAssetDetailInfoWithContext(ctx context.Context, request *DescribeComplianceAssetDetailInfoRequest) (response *DescribeComplianceAssetDetailInfoResponse, err error) {
    if request == nil {
        request = NewDescribeComplianceAssetDetailInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeComplianceAssetDetailInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeComplianceAssetDetailInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComplianceAssetListRequest() (request *DescribeComplianceAssetListRequest) {
    request = &DescribeComplianceAssetListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeComplianceAssetList")
    
    
    return
}

func NewDescribeComplianceAssetListResponse() (response *DescribeComplianceAssetListResponse) {
    response = &DescribeComplianceAssetListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeComplianceAssetList
// This API is used to query the list of assets of a certain type.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DescribeComplianceAssetList(request *DescribeComplianceAssetListRequest) (response *DescribeComplianceAssetListResponse, err error) {
    return c.DescribeComplianceAssetListWithContext(context.Background(), request)
}

// DescribeComplianceAssetList
// This API is used to query the list of assets of a certain type.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DescribeComplianceAssetListWithContext(ctx context.Context, request *DescribeComplianceAssetListRequest) (response *DescribeComplianceAssetListResponse, err error) {
    if request == nil {
        request = NewDescribeComplianceAssetListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeComplianceAssetList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeComplianceAssetListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComplianceAssetPolicyItemListRequest() (request *DescribeComplianceAssetPolicyItemListRequest) {
    request = &DescribeComplianceAssetPolicyItemListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeComplianceAssetPolicyItemList")
    
    
    return
}

func NewDescribeComplianceAssetPolicyItemListResponse() (response *DescribeComplianceAssetPolicyItemListResponse) {
    response = &DescribeComplianceAssetPolicyItemListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeComplianceAssetPolicyItemList
// This API is used to query the list of check items of an asset.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DescribeComplianceAssetPolicyItemList(request *DescribeComplianceAssetPolicyItemListRequest) (response *DescribeComplianceAssetPolicyItemListResponse, err error) {
    return c.DescribeComplianceAssetPolicyItemListWithContext(context.Background(), request)
}

// DescribeComplianceAssetPolicyItemList
// This API is used to query the list of check items of an asset.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DescribeComplianceAssetPolicyItemListWithContext(ctx context.Context, request *DescribeComplianceAssetPolicyItemListRequest) (response *DescribeComplianceAssetPolicyItemListResponse, err error) {
    if request == nil {
        request = NewDescribeComplianceAssetPolicyItemListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeComplianceAssetPolicyItemList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeComplianceAssetPolicyItemListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCompliancePeriodTaskListRequest() (request *DescribeCompliancePeriodTaskListRequest) {
    request = &DescribeCompliancePeriodTaskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeCompliancePeriodTaskList")
    
    
    return
}

func NewDescribeCompliancePeriodTaskListResponse() (response *DescribeCompliancePeriodTaskListResponse) {
    response = &DescribeCompliancePeriodTaskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCompliancePeriodTaskList
// This API is used to query the list of scheduled tasks.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DescribeCompliancePeriodTaskList(request *DescribeCompliancePeriodTaskListRequest) (response *DescribeCompliancePeriodTaskListResponse, err error) {
    return c.DescribeCompliancePeriodTaskListWithContext(context.Background(), request)
}

// DescribeCompliancePeriodTaskList
// This API is used to query the list of scheduled tasks.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DescribeCompliancePeriodTaskListWithContext(ctx context.Context, request *DescribeCompliancePeriodTaskListRequest) (response *DescribeCompliancePeriodTaskListResponse, err error) {
    if request == nil {
        request = NewDescribeCompliancePeriodTaskListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCompliancePeriodTaskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCompliancePeriodTaskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCompliancePolicyItemAffectedAssetListRequest() (request *DescribeCompliancePolicyItemAffectedAssetListRequest) {
    request = &DescribeCompliancePolicyItemAffectedAssetListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeCompliancePolicyItemAffectedAssetList")
    
    
    return
}

func NewDescribeCompliancePolicyItemAffectedAssetListResponse() (response *DescribeCompliancePolicyItemAffectedAssetListResponse) {
    response = &DescribeCompliancePolicyItemAffectedAssetListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCompliancePolicyItemAffectedAssetList
// This API is used to apply the asset level in the "check item + asset" two-level structure.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DescribeCompliancePolicyItemAffectedAssetList(request *DescribeCompliancePolicyItemAffectedAssetListRequest) (response *DescribeCompliancePolicyItemAffectedAssetListResponse, err error) {
    return c.DescribeCompliancePolicyItemAffectedAssetListWithContext(context.Background(), request)
}

// DescribeCompliancePolicyItemAffectedAssetList
// This API is used to apply the asset level in the "check item + asset" two-level structure.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DescribeCompliancePolicyItemAffectedAssetListWithContext(ctx context.Context, request *DescribeCompliancePolicyItemAffectedAssetListRequest) (response *DescribeCompliancePolicyItemAffectedAssetListResponse, err error) {
    if request == nil {
        request = NewDescribeCompliancePolicyItemAffectedAssetListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCompliancePolicyItemAffectedAssetList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCompliancePolicyItemAffectedAssetListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCompliancePolicyItemAffectedSummaryRequest() (request *DescribeCompliancePolicyItemAffectedSummaryRequest) {
    request = &DescribeCompliancePolicyItemAffectedSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeCompliancePolicyItemAffectedSummary")
    
    
    return
}

func NewDescribeCompliancePolicyItemAffectedSummaryResponse() (response *DescribeCompliancePolicyItemAffectedSummaryResponse) {
    response = &DescribeCompliancePolicyItemAffectedSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCompliancePolicyItemAffectedSummary
// This API is used to apply the check item level in the "check item + asset" two-level structure.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DescribeCompliancePolicyItemAffectedSummary(request *DescribeCompliancePolicyItemAffectedSummaryRequest) (response *DescribeCompliancePolicyItemAffectedSummaryResponse, err error) {
    return c.DescribeCompliancePolicyItemAffectedSummaryWithContext(context.Background(), request)
}

// DescribeCompliancePolicyItemAffectedSummary
// This API is used to apply the check item level in the "check item + asset" two-level structure.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DescribeCompliancePolicyItemAffectedSummaryWithContext(ctx context.Context, request *DescribeCompliancePolicyItemAffectedSummaryRequest) (response *DescribeCompliancePolicyItemAffectedSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeCompliancePolicyItemAffectedSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCompliancePolicyItemAffectedSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCompliancePolicyItemAffectedSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComplianceScanFailedAssetListRequest() (request *DescribeComplianceScanFailedAssetListRequest) {
    request = &DescribeComplianceScanFailedAssetListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeComplianceScanFailedAssetList")
    
    
    return
}

func NewDescribeComplianceScanFailedAssetListResponse() (response *DescribeComplianceScanFailedAssetListResponse) {
    response = &DescribeComplianceScanFailedAssetListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeComplianceScanFailedAssetList
// This API is used to query the aggregate information of the pass rate at the asset level, the first level in the "asset + check item" two-level structure.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DescribeComplianceScanFailedAssetList(request *DescribeComplianceScanFailedAssetListRequest) (response *DescribeComplianceScanFailedAssetListResponse, err error) {
    return c.DescribeComplianceScanFailedAssetListWithContext(context.Background(), request)
}

// DescribeComplianceScanFailedAssetList
// This API is used to query the aggregate information of the pass rate at the asset level, the first level in the "asset + check item" two-level structure.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DescribeComplianceScanFailedAssetListWithContext(ctx context.Context, request *DescribeComplianceScanFailedAssetListRequest) (response *DescribeComplianceScanFailedAssetListResponse, err error) {
    if request == nil {
        request = NewDescribeComplianceScanFailedAssetListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeComplianceScanFailedAssetList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeComplianceScanFailedAssetListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComplianceTaskAssetSummaryRequest() (request *DescribeComplianceTaskAssetSummaryRequest) {
    request = &DescribeComplianceTaskAssetSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeComplianceTaskAssetSummary")
    
    
    return
}

func NewDescribeComplianceTaskAssetSummaryResponse() (response *DescribeComplianceTaskAssetSummaryResponse) {
    response = &DescribeComplianceTaskAssetSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeComplianceTaskAssetSummary
// This API is used to query the aggregated information of the asset pass rate in the last task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DescribeComplianceTaskAssetSummary(request *DescribeComplianceTaskAssetSummaryRequest) (response *DescribeComplianceTaskAssetSummaryResponse, err error) {
    return c.DescribeComplianceTaskAssetSummaryWithContext(context.Background(), request)
}

// DescribeComplianceTaskAssetSummary
// This API is used to query the aggregated information of the asset pass rate in the last task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DescribeComplianceTaskAssetSummaryWithContext(ctx context.Context, request *DescribeComplianceTaskAssetSummaryRequest) (response *DescribeComplianceTaskAssetSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeComplianceTaskAssetSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeComplianceTaskAssetSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeComplianceTaskAssetSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComplianceTaskPolicyItemSummaryListRequest() (request *DescribeComplianceTaskPolicyItemSummaryListRequest) {
    request = &DescribeComplianceTaskPolicyItemSummaryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeComplianceTaskPolicyItemSummaryList")
    
    
    return
}

func NewDescribeComplianceTaskPolicyItemSummaryListResponse() (response *DescribeComplianceTaskPolicyItemSummaryListResponse) {
    response = &DescribeComplianceTaskPolicyItemSummaryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeComplianceTaskPolicyItemSummaryList
// This API is used to query the list of aggregated information of check items identified in the last task in line with the "check item + asset" two-level structure.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeComplianceTaskPolicyItemSummaryList(request *DescribeComplianceTaskPolicyItemSummaryListRequest) (response *DescribeComplianceTaskPolicyItemSummaryListResponse, err error) {
    return c.DescribeComplianceTaskPolicyItemSummaryListWithContext(context.Background(), request)
}

// DescribeComplianceTaskPolicyItemSummaryList
// This API is used to query the list of aggregated information of check items identified in the last task in line with the "check item + asset" two-level structure.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeComplianceTaskPolicyItemSummaryListWithContext(ctx context.Context, request *DescribeComplianceTaskPolicyItemSummaryListRequest) (response *DescribeComplianceTaskPolicyItemSummaryListResponse, err error) {
    if request == nil {
        request = NewDescribeComplianceTaskPolicyItemSummaryListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeComplianceTaskPolicyItemSummaryList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeComplianceTaskPolicyItemSummaryListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComplianceWhitelistItemListRequest() (request *DescribeComplianceWhitelistItemListRequest) {
    request = &DescribeComplianceWhitelistItemListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeComplianceWhitelistItemList")
    
    
    return
}

func NewDescribeComplianceWhitelistItemListResponse() (response *DescribeComplianceWhitelistItemListResponse) {
    response = &DescribeComplianceWhitelistItemListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeComplianceWhitelistItemList
// This API is used to query the allowlist.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DescribeComplianceWhitelistItemList(request *DescribeComplianceWhitelistItemListRequest) (response *DescribeComplianceWhitelistItemListResponse, err error) {
    return c.DescribeComplianceWhitelistItemListWithContext(context.Background(), request)
}

// DescribeComplianceWhitelistItemList
// This API is used to query the allowlist.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DescribeComplianceWhitelistItemListWithContext(ctx context.Context, request *DescribeComplianceWhitelistItemListRequest) (response *DescribeComplianceWhitelistItemListResponse, err error) {
    if request == nil {
        request = NewDescribeComplianceWhitelistItemListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeComplianceWhitelistItemList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeComplianceWhitelistItemListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeContainerAssetSummaryRequest() (request *DescribeContainerAssetSummaryRequest) {
    request = &DescribeContainerAssetSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeContainerAssetSummary")
    
    
    return
}

func NewDescribeContainerAssetSummaryResponse() (response *DescribeContainerAssetSummaryResponse) {
    response = &DescribeContainerAssetSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeContainerAssetSummary
// This API is used to query the asset overview.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeContainerAssetSummary(request *DescribeContainerAssetSummaryRequest) (response *DescribeContainerAssetSummaryResponse, err error) {
    return c.DescribeContainerAssetSummaryWithContext(context.Background(), request)
}

// DescribeContainerAssetSummary
// This API is used to query the asset overview.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeContainerAssetSummaryWithContext(ctx context.Context, request *DescribeContainerAssetSummaryRequest) (response *DescribeContainerAssetSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeContainerAssetSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeContainerAssetSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeContainerAssetSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeContainerSecEventSummaryRequest() (request *DescribeContainerSecEventSummaryRequest) {
    request = &DescribeContainerSecEventSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeContainerSecEventSummary")
    
    
    return
}

func NewDescribeContainerSecEventSummaryResponse() (response *DescribeContainerSecEventSummaryResponse) {
    response = &DescribeContainerSecEventSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeContainerSecEventSummary
// This API is used to query the overview of pending events.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeContainerSecEventSummary(request *DescribeContainerSecEventSummaryRequest) (response *DescribeContainerSecEventSummaryResponse, err error) {
    return c.DescribeContainerSecEventSummaryWithContext(context.Background(), request)
}

// DescribeContainerSecEventSummary
// This API is used to query the overview of pending events.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeContainerSecEventSummaryWithContext(ctx context.Context, request *DescribeContainerSecEventSummaryRequest) (response *DescribeContainerSecEventSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeContainerSecEventSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeContainerSecEventSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeContainerSecEventSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeESAggregationsRequest() (request *DescribeESAggregationsRequest) {
    request = &DescribeESAggregationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeESAggregations")
    
    
    return
}

func NewDescribeESAggregationsResponse() (response *DescribeESAggregationsResponse) {
    response = &DescribeESAggregationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeESAggregations
// This API is used to get the aggregation result of the ES field.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeESAggregations(request *DescribeESAggregationsRequest) (response *DescribeESAggregationsResponse, err error) {
    return c.DescribeESAggregationsWithContext(context.Background(), request)
}

// DescribeESAggregations
// This API is used to get the aggregation result of the ES field.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeESAggregationsWithContext(ctx context.Context, request *DescribeESAggregationsRequest) (response *DescribeESAggregationsResponse, err error) {
    if request == nil {
        request = NewDescribeESAggregationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeESAggregations require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeESAggregationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeESHitsRequest() (request *DescribeESHitsRequest) {
    request = &DescribeESHitsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeESHits")
    
    
    return
}

func NewDescribeESHitsResponse() (response *DescribeESHitsResponse) {
    response = &DescribeESHitsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeESHits
// This API is used to get the list of ES query files.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeESHits(request *DescribeESHitsRequest) (response *DescribeESHitsResponse, err error) {
    return c.DescribeESHitsWithContext(context.Background(), request)
}

// DescribeESHits
// This API is used to get the list of ES query files.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeESHitsWithContext(ctx context.Context, request *DescribeESHitsRequest) (response *DescribeESHitsResponse, err error) {
    if request == nil {
        request = NewDescribeESHitsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeESHits require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeESHitsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEmergencyVulListRequest() (request *DescribeEmergencyVulListRequest) {
    request = &DescribeEmergencyVulListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeEmergencyVulList")
    
    
    return
}

func NewDescribeEmergencyVulListResponse() (response *DescribeEmergencyVulListResponse) {
    response = &DescribeEmergencyVulListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEmergencyVulList
// This API is used to query the list of emergency vulnerabilities.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeEmergencyVulList(request *DescribeEmergencyVulListRequest) (response *DescribeEmergencyVulListResponse, err error) {
    return c.DescribeEmergencyVulListWithContext(context.Background(), request)
}

// DescribeEmergencyVulList
// This API is used to query the list of emergency vulnerabilities.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeEmergencyVulListWithContext(ctx context.Context, request *DescribeEmergencyVulListRequest) (response *DescribeEmergencyVulListResponse, err error) {
    if request == nil {
        request = NewDescribeEmergencyVulListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEmergencyVulList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEmergencyVulListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEscapeEventDetailRequest() (request *DescribeEscapeEventDetailRequest) {
    request = &DescribeEscapeEventDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeEscapeEventDetail")
    
    
    return
}

func NewDescribeEscapeEventDetailResponse() (response *DescribeEscapeEventDetailResponse) {
    response = &DescribeEscapeEventDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEscapeEventDetail
// This API is used to query the details of a container escape event.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
func (c *Client) DescribeEscapeEventDetail(request *DescribeEscapeEventDetailRequest) (response *DescribeEscapeEventDetailResponse, err error) {
    return c.DescribeEscapeEventDetailWithContext(context.Background(), request)
}

// DescribeEscapeEventDetail
// This API is used to query the details of a container escape event.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
func (c *Client) DescribeEscapeEventDetailWithContext(ctx context.Context, request *DescribeEscapeEventDetailRequest) (response *DescribeEscapeEventDetailResponse, err error) {
    if request == nil {
        request = NewDescribeEscapeEventDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEscapeEventDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEscapeEventDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEscapeEventInfoRequest() (request *DescribeEscapeEventInfoRequest) {
    request = &DescribeEscapeEventInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeEscapeEventInfo")
    
    
    return
}

func NewDescribeEscapeEventInfoResponse() (response *DescribeEscapeEventInfoResponse) {
    response = &DescribeEscapeEventInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEscapeEventInfo
// This API is used to query the list of container escape events.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeEscapeEventInfo(request *DescribeEscapeEventInfoRequest) (response *DescribeEscapeEventInfoResponse, err error) {
    return c.DescribeEscapeEventInfoWithContext(context.Background(), request)
}

// DescribeEscapeEventInfo
// This API is used to query the list of container escape events.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeEscapeEventInfoWithContext(ctx context.Context, request *DescribeEscapeEventInfoRequest) (response *DescribeEscapeEventInfoResponse, err error) {
    if request == nil {
        request = NewDescribeEscapeEventInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEscapeEventInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEscapeEventInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEscapeEventTendencyRequest() (request *DescribeEscapeEventTendencyRequest) {
    request = &DescribeEscapeEventTendencyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeEscapeEventTendency")
    
    
    return
}

func NewDescribeEscapeEventTendencyResponse() (response *DescribeEscapeEventTendencyResponse) {
    response = &DescribeEscapeEventTendencyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEscapeEventTendency
// This API is used to query the trend of pending escape events.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeEscapeEventTendency(request *DescribeEscapeEventTendencyRequest) (response *DescribeEscapeEventTendencyResponse, err error) {
    return c.DescribeEscapeEventTendencyWithContext(context.Background(), request)
}

// DescribeEscapeEventTendency
// This API is used to query the trend of pending escape events.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeEscapeEventTendencyWithContext(ctx context.Context, request *DescribeEscapeEventTendencyRequest) (response *DescribeEscapeEventTendencyResponse, err error) {
    if request == nil {
        request = NewDescribeEscapeEventTendencyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEscapeEventTendency require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEscapeEventTendencyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEscapeEventTypeSummaryRequest() (request *DescribeEscapeEventTypeSummaryRequest) {
    request = &DescribeEscapeEventTypeSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeEscapeEventTypeSummary")
    
    
    return
}

func NewDescribeEscapeEventTypeSummaryResponse() (response *DescribeEscapeEventTypeSummaryResponse) {
    response = &DescribeEscapeEventTypeSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEscapeEventTypeSummary
// This API is used to query the types of container escape events and the number of pending events.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) DescribeEscapeEventTypeSummary(request *DescribeEscapeEventTypeSummaryRequest) (response *DescribeEscapeEventTypeSummaryResponse, err error) {
    return c.DescribeEscapeEventTypeSummaryWithContext(context.Background(), request)
}

// DescribeEscapeEventTypeSummary
// This API is used to query the types of container escape events and the number of pending events.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) DescribeEscapeEventTypeSummaryWithContext(ctx context.Context, request *DescribeEscapeEventTypeSummaryRequest) (response *DescribeEscapeEventTypeSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeEscapeEventTypeSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEscapeEventTypeSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEscapeEventTypeSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEscapeEventsExportRequest() (request *DescribeEscapeEventsExportRequest) {
    request = &DescribeEscapeEventsExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeEscapeEventsExport")
    
    
    return
}

func NewDescribeEscapeEventsExportResponse() (response *DescribeEscapeEventsExportResponse) {
    response = &DescribeEscapeEventsExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEscapeEventsExport
// 接口已废弃
//
// 
//
// This API is used to export the list of container escape events.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeEscapeEventsExport(request *DescribeEscapeEventsExportRequest) (response *DescribeEscapeEventsExportResponse, err error) {
    return c.DescribeEscapeEventsExportWithContext(context.Background(), request)
}

// DescribeEscapeEventsExport
// 接口已废弃
//
// 
//
// This API is used to export the list of container escape events.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeEscapeEventsExportWithContext(ctx context.Context, request *DescribeEscapeEventsExportRequest) (response *DescribeEscapeEventsExportResponse, err error) {
    if request == nil {
        request = NewDescribeEscapeEventsExportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEscapeEventsExport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEscapeEventsExportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEscapeRuleInfoRequest() (request *DescribeEscapeRuleInfoRequest) {
    request = &DescribeEscapeRuleInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeEscapeRuleInfo")
    
    
    return
}

func NewDescribeEscapeRuleInfoResponse() (response *DescribeEscapeRuleInfoResponse) {
    response = &DescribeEscapeRuleInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEscapeRuleInfo
// This API is used to query the information of a container escape scan rule.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeEscapeRuleInfo(request *DescribeEscapeRuleInfoRequest) (response *DescribeEscapeRuleInfoResponse, err error) {
    return c.DescribeEscapeRuleInfoWithContext(context.Background(), request)
}

// DescribeEscapeRuleInfo
// This API is used to query the information of a container escape scan rule.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeEscapeRuleInfoWithContext(ctx context.Context, request *DescribeEscapeRuleInfoRequest) (response *DescribeEscapeRuleInfoResponse, err error) {
    if request == nil {
        request = NewDescribeEscapeRuleInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEscapeRuleInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEscapeRuleInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEscapeSafeStateRequest() (request *DescribeEscapeSafeStateRequest) {
    request = &DescribeEscapeSafeStateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeEscapeSafeState")
    
    
    return
}

func NewDescribeEscapeSafeStateResponse() (response *DescribeEscapeSafeStateResponse) {
    response = &DescribeEscapeSafeStateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEscapeSafeState
// This API is used to query the container escape security status.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeEscapeSafeState(request *DescribeEscapeSafeStateRequest) (response *DescribeEscapeSafeStateResponse, err error) {
    return c.DescribeEscapeSafeStateWithContext(context.Background(), request)
}

// DescribeEscapeSafeState
// This API is used to query the container escape security status.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeEscapeSafeStateWithContext(ctx context.Context, request *DescribeEscapeSafeStateRequest) (response *DescribeEscapeSafeStateResponse, err error) {
    if request == nil {
        request = NewDescribeEscapeSafeStateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEscapeSafeState require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEscapeSafeStateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEscapeWhiteListRequest() (request *DescribeEscapeWhiteListRequest) {
    request = &DescribeEscapeWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeEscapeWhiteList")
    
    
    return
}

func NewDescribeEscapeWhiteListResponse() (response *DescribeEscapeWhiteListResponse) {
    response = &DescribeEscapeWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEscapeWhiteList
// This API is used to query the allowlist of escapes.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeEscapeWhiteList(request *DescribeEscapeWhiteListRequest) (response *DescribeEscapeWhiteListResponse, err error) {
    return c.DescribeEscapeWhiteListWithContext(context.Background(), request)
}

// DescribeEscapeWhiteList
// This API is used to query the allowlist of escapes.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeEscapeWhiteListWithContext(ctx context.Context, request *DescribeEscapeWhiteListRequest) (response *DescribeEscapeWhiteListResponse, err error) {
    if request == nil {
        request = NewDescribeEscapeWhiteListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEscapeWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEscapeWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExportJobDownloadURLRequest() (request *DescribeExportJobDownloadURLRequest) {
    request = &DescribeExportJobDownloadURLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeExportJobDownloadURL")
    
    
    return
}

func NewDescribeExportJobDownloadURLResponse() (response *DescribeExportJobDownloadURLResponse) {
    response = &DescribeExportJobDownloadURLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeExportJobDownloadURL
// This API is used to query the URL to download the result of an exportation task.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeExportJobDownloadURL(request *DescribeExportJobDownloadURLRequest) (response *DescribeExportJobDownloadURLResponse, err error) {
    return c.DescribeExportJobDownloadURLWithContext(context.Background(), request)
}

// DescribeExportJobDownloadURL
// This API is used to query the URL to download the result of an exportation task.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeExportJobDownloadURLWithContext(ctx context.Context, request *DescribeExportJobDownloadURLRequest) (response *DescribeExportJobDownloadURLResponse, err error) {
    if request == nil {
        request = NewDescribeExportJobDownloadURLRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExportJobDownloadURL require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExportJobDownloadURLResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExportJobManageListRequest() (request *DescribeExportJobManageListRequest) {
    request = &DescribeExportJobManageListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeExportJobManageList")
    
    
    return
}

func NewDescribeExportJobManageListResponse() (response *DescribeExportJobManageListResponse) {
    response = &DescribeExportJobManageListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeExportJobManageList
// This API is used to query the export job management list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeExportJobManageList(request *DescribeExportJobManageListRequest) (response *DescribeExportJobManageListResponse, err error) {
    return c.DescribeExportJobManageListWithContext(context.Background(), request)
}

// DescribeExportJobManageList
// This API is used to query the export job management list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeExportJobManageListWithContext(ctx context.Context, request *DescribeExportJobManageListRequest) (response *DescribeExportJobManageListResponse, err error) {
    if request == nil {
        request = NewDescribeExportJobManageListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExportJobManageList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExportJobManageListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExportJobResultRequest() (request *DescribeExportJobResultRequest) {
    request = &DescribeExportJobResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeExportJobResult")
    
    
    return
}

func NewDescribeExportJobResultResponse() (response *DescribeExportJobResultResponse) {
    response = &DescribeExportJobResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeExportJobResult
// This API is used to query the result of an export task.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeExportJobResult(request *DescribeExportJobResultRequest) (response *DescribeExportJobResultResponse, err error) {
    return c.DescribeExportJobResultWithContext(context.Background(), request)
}

// DescribeExportJobResult
// This API is used to query the result of an export task.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeExportJobResultWithContext(ctx context.Context, request *DescribeExportJobResultRequest) (response *DescribeExportJobResultResponse, err error) {
    if request == nil {
        request = NewDescribeExportJobResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExportJobResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExportJobResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageAuthorizedInfoRequest() (request *DescribeImageAuthorizedInfoRequest) {
    request = &DescribeImageAuthorizedInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeImageAuthorizedInfo")
    
    
    return
}

func NewDescribeImageAuthorizedInfoResponse() (response *DescribeImageAuthorizedInfoResponse) {
    response = &DescribeImageAuthorizedInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeImageAuthorizedInfo
// This API is used to query the image licensing information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
func (c *Client) DescribeImageAuthorizedInfo(request *DescribeImageAuthorizedInfoRequest) (response *DescribeImageAuthorizedInfoResponse, err error) {
    return c.DescribeImageAuthorizedInfoWithContext(context.Background(), request)
}

// DescribeImageAuthorizedInfo
// This API is used to query the image licensing information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
func (c *Client) DescribeImageAuthorizedInfoWithContext(ctx context.Context, request *DescribeImageAuthorizedInfoRequest) (response *DescribeImageAuthorizedInfoResponse, err error) {
    if request == nil {
        request = NewDescribeImageAuthorizedInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImageAuthorizedInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImageAuthorizedInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageAutoAuthorizedLogListRequest() (request *DescribeImageAutoAuthorizedLogListRequest) {
    request = &DescribeImageAutoAuthorizedLogListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeImageAutoAuthorizedLogList")
    
    
    return
}

func NewDescribeImageAutoAuthorizedLogListResponse() (response *DescribeImageAutoAuthorizedLogListResponse) {
    response = &DescribeImageAutoAuthorizedLogListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeImageAutoAuthorizedLogList
// This API is used to query the list of automatic image licensing results.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeImageAutoAuthorizedLogList(request *DescribeImageAutoAuthorizedLogListRequest) (response *DescribeImageAutoAuthorizedLogListResponse, err error) {
    return c.DescribeImageAutoAuthorizedLogListWithContext(context.Background(), request)
}

// DescribeImageAutoAuthorizedLogList
// This API is used to query the list of automatic image licensing results.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeImageAutoAuthorizedLogListWithContext(ctx context.Context, request *DescribeImageAutoAuthorizedLogListRequest) (response *DescribeImageAutoAuthorizedLogListResponse, err error) {
    if request == nil {
        request = NewDescribeImageAutoAuthorizedLogListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImageAutoAuthorizedLogList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImageAutoAuthorizedLogListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageAutoAuthorizedRuleRequest() (request *DescribeImageAutoAuthorizedRuleRequest) {
    request = &DescribeImageAutoAuthorizedRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeImageAutoAuthorizedRule")
    
    
    return
}

func NewDescribeImageAutoAuthorizedRuleResponse() (response *DescribeImageAutoAuthorizedRuleResponse) {
    response = &DescribeImageAutoAuthorizedRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeImageAutoAuthorizedRule
// This API is used to query an automatic licensing rule for local images.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeImageAutoAuthorizedRule(request *DescribeImageAutoAuthorizedRuleRequest) (response *DescribeImageAutoAuthorizedRuleResponse, err error) {
    return c.DescribeImageAutoAuthorizedRuleWithContext(context.Background(), request)
}

// DescribeImageAutoAuthorizedRule
// This API is used to query an automatic licensing rule for local images.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeImageAutoAuthorizedRuleWithContext(ctx context.Context, request *DescribeImageAutoAuthorizedRuleRequest) (response *DescribeImageAutoAuthorizedRuleResponse, err error) {
    if request == nil {
        request = NewDescribeImageAutoAuthorizedRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImageAutoAuthorizedRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImageAutoAuthorizedRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageAutoAuthorizedTaskListRequest() (request *DescribeImageAutoAuthorizedTaskListRequest) {
    request = &DescribeImageAutoAuthorizedTaskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeImageAutoAuthorizedTaskList")
    
    
    return
}

func NewDescribeImageAutoAuthorizedTaskListResponse() (response *DescribeImageAutoAuthorizedTaskListResponse) {
    response = &DescribeImageAutoAuthorizedTaskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeImageAutoAuthorizedTaskList
// This API is used to query the list of automatic image licensing tasks.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeImageAutoAuthorizedTaskList(request *DescribeImageAutoAuthorizedTaskListRequest) (response *DescribeImageAutoAuthorizedTaskListResponse, err error) {
    return c.DescribeImageAutoAuthorizedTaskListWithContext(context.Background(), request)
}

// DescribeImageAutoAuthorizedTaskList
// This API is used to query the list of automatic image licensing tasks.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeImageAutoAuthorizedTaskListWithContext(ctx context.Context, request *DescribeImageAutoAuthorizedTaskListRequest) (response *DescribeImageAutoAuthorizedTaskListResponse, err error) {
    if request == nil {
        request = NewDescribeImageAutoAuthorizedTaskListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImageAutoAuthorizedTaskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImageAutoAuthorizedTaskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageComponentListRequest() (request *DescribeImageComponentListRequest) {
    request = &DescribeImageComponentListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeImageComponentList")
    
    
    return
}

func NewDescribeImageComponentListResponse() (response *DescribeImageComponentListResponse) {
    response = &DescribeImageComponentListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeImageComponentList
// This API is used to query the list of components in an local image.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeImageComponentList(request *DescribeImageComponentListRequest) (response *DescribeImageComponentListResponse, err error) {
    return c.DescribeImageComponentListWithContext(context.Background(), request)
}

// DescribeImageComponentList
// This API is used to query the list of components in an local image.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeImageComponentListWithContext(ctx context.Context, request *DescribeImageComponentListRequest) (response *DescribeImageComponentListResponse, err error) {
    if request == nil {
        request = NewDescribeImageComponentListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImageComponentList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImageComponentListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageRegistryNamespaceListRequest() (request *DescribeImageRegistryNamespaceListRequest) {
    request = &DescribeImageRegistryNamespaceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeImageRegistryNamespaceList")
    
    
    return
}

func NewDescribeImageRegistryNamespaceListResponse() (response *DescribeImageRegistryNamespaceListResponse) {
    response = &DescribeImageRegistryNamespaceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeImageRegistryNamespaceList
// This API is used to query the list of namespaces for an image repository.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeImageRegistryNamespaceList(request *DescribeImageRegistryNamespaceListRequest) (response *DescribeImageRegistryNamespaceListResponse, err error) {
    return c.DescribeImageRegistryNamespaceListWithContext(context.Background(), request)
}

// DescribeImageRegistryNamespaceList
// This API is used to query the list of namespaces for an image repository.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeImageRegistryNamespaceListWithContext(ctx context.Context, request *DescribeImageRegistryNamespaceListRequest) (response *DescribeImageRegistryNamespaceListResponse, err error) {
    if request == nil {
        request = NewDescribeImageRegistryNamespaceListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImageRegistryNamespaceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImageRegistryNamespaceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageRegistryTimingScanTaskRequest() (request *DescribeImageRegistryTimingScanTaskRequest) {
    request = &DescribeImageRegistryTimingScanTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeImageRegistryTimingScanTask")
    
    
    return
}

func NewDescribeImageRegistryTimingScanTaskResponse() (response *DescribeImageRegistryTimingScanTaskResponse) {
    response = &DescribeImageRegistryTimingScanTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeImageRegistryTimingScanTask
// This API is used to view a scheduled task of an image repository.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeImageRegistryTimingScanTask(request *DescribeImageRegistryTimingScanTaskRequest) (response *DescribeImageRegistryTimingScanTaskResponse, err error) {
    return c.DescribeImageRegistryTimingScanTaskWithContext(context.Background(), request)
}

// DescribeImageRegistryTimingScanTask
// This API is used to view a scheduled task of an image repository.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeImageRegistryTimingScanTaskWithContext(ctx context.Context, request *DescribeImageRegistryTimingScanTaskRequest) (response *DescribeImageRegistryTimingScanTaskResponse, err error) {
    if request == nil {
        request = NewDescribeImageRegistryTimingScanTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImageRegistryTimingScanTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImageRegistryTimingScanTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageRiskSummaryRequest() (request *DescribeImageRiskSummaryRequest) {
    request = &DescribeImageRiskSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeImageRiskSummary")
    
    
    return
}

func NewDescribeImageRiskSummaryResponse() (response *DescribeImageRiskSummaryResponse) {
    response = &DescribeImageRiskSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeImageRiskSummary
// This API is used to query the overview of local image risks.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeImageRiskSummary(request *DescribeImageRiskSummaryRequest) (response *DescribeImageRiskSummaryResponse, err error) {
    return c.DescribeImageRiskSummaryWithContext(context.Background(), request)
}

// DescribeImageRiskSummary
// This API is used to query the overview of local image risks.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeImageRiskSummaryWithContext(ctx context.Context, request *DescribeImageRiskSummaryRequest) (response *DescribeImageRiskSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeImageRiskSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImageRiskSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImageRiskSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageRiskTendencyRequest() (request *DescribeImageRiskTendencyRequest) {
    request = &DescribeImageRiskTendencyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeImageRiskTendency")
    
    
    return
}

func NewDescribeImageRiskTendencyResponse() (response *DescribeImageRiskTendencyResponse) {
    response = &DescribeImageRiskTendencyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeImageRiskTendency
// This API is used to query the trend of local image risks.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATARANGE = "InvalidParameterValue.DataRange"
func (c *Client) DescribeImageRiskTendency(request *DescribeImageRiskTendencyRequest) (response *DescribeImageRiskTendencyResponse, err error) {
    return c.DescribeImageRiskTendencyWithContext(context.Background(), request)
}

// DescribeImageRiskTendency
// This API is used to query the trend of local image risks.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATARANGE = "InvalidParameterValue.DataRange"
func (c *Client) DescribeImageRiskTendencyWithContext(ctx context.Context, request *DescribeImageRiskTendencyRequest) (response *DescribeImageRiskTendencyResponse, err error) {
    if request == nil {
        request = NewDescribeImageRiskTendencyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImageRiskTendency require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImageRiskTendencyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageSimpleListRequest() (request *DescribeImageSimpleListRequest) {
    request = &DescribeImageSimpleListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeImageSimpleList")
    
    
    return
}

func NewDescribeImageSimpleListResponse() (response *DescribeImageSimpleListResponse) {
    response = &DescribeImageSimpleListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeImageSimpleList
// This API is used to query the list of all images.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeImageSimpleList(request *DescribeImageSimpleListRequest) (response *DescribeImageSimpleListResponse, err error) {
    return c.DescribeImageSimpleListWithContext(context.Background(), request)
}

// DescribeImageSimpleList
// This API is used to query the list of all images.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeImageSimpleListWithContext(ctx context.Context, request *DescribeImageSimpleListRequest) (response *DescribeImageSimpleListResponse, err error) {
    if request == nil {
        request = NewDescribeImageSimpleListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImageSimpleList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImageSimpleListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIndexListRequest() (request *DescribeIndexListRequest) {
    request = &DescribeIndexListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeIndexList")
    
    
    return
}

func NewDescribeIndexListResponse() (response *DescribeIndexListResponse) {
    response = &DescribeIndexListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIndexList
// This API is used to get the index list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeIndexList(request *DescribeIndexListRequest) (response *DescribeIndexListResponse, err error) {
    return c.DescribeIndexListWithContext(context.Background(), request)
}

// DescribeIndexList
// This API is used to get the index list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeIndexListWithContext(ctx context.Context, request *DescribeIndexListRequest) (response *DescribeIndexListResponse, err error) {
    if request == nil {
        request = NewDescribeIndexListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIndexList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIndexListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInspectionReportRequest() (request *DescribeInspectionReportRequest) {
    request = &DescribeInspectionReportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeInspectionReport")
    
    
    return
}

func NewDescribeInspectionReportResponse() (response *DescribeInspectionReportResponse) {
    response = &DescribeInspectionReportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInspectionReport
// This API is used to query check reports.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeInspectionReport(request *DescribeInspectionReportRequest) (response *DescribeInspectionReportResponse, err error) {
    return c.DescribeInspectionReportWithContext(context.Background(), request)
}

// DescribeInspectionReport
// This API is used to query check reports.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeInspectionReportWithContext(ctx context.Context, request *DescribeInspectionReportRequest) (response *DescribeInspectionReportResponse, err error) {
    if request == nil {
        request = NewDescribeInspectionReportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInspectionReport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInspectionReportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeK8sApiAbnormalEventInfoRequest() (request *DescribeK8sApiAbnormalEventInfoRequest) {
    request = &DescribeK8sApiAbnormalEventInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeK8sApiAbnormalEventInfo")
    
    
    return
}

func NewDescribeK8sApiAbnormalEventInfoResponse() (response *DescribeK8sApiAbnormalEventInfoResponse) {
    response = &DescribeK8sApiAbnormalEventInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeK8sApiAbnormalEventInfo
// Querying details of a K8s API exception event
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeK8sApiAbnormalEventInfo(request *DescribeK8sApiAbnormalEventInfoRequest) (response *DescribeK8sApiAbnormalEventInfoResponse, err error) {
    return c.DescribeK8sApiAbnormalEventInfoWithContext(context.Background(), request)
}

// DescribeK8sApiAbnormalEventInfo
// Querying details of a K8s API exception event
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeK8sApiAbnormalEventInfoWithContext(ctx context.Context, request *DescribeK8sApiAbnormalEventInfoRequest) (response *DescribeK8sApiAbnormalEventInfoResponse, err error) {
    if request == nil {
        request = NewDescribeK8sApiAbnormalEventInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeK8sApiAbnormalEventInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeK8sApiAbnormalEventInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeK8sApiAbnormalEventListRequest() (request *DescribeK8sApiAbnormalEventListRequest) {
    request = &DescribeK8sApiAbnormalEventListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeK8sApiAbnormalEventList")
    
    
    return
}

func NewDescribeK8sApiAbnormalEventListResponse() (response *DescribeK8sApiAbnormalEventListResponse) {
    response = &DescribeK8sApiAbnormalEventListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeK8sApiAbnormalEventList
// This API is used to query the K8sApi abnormal event list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeK8sApiAbnormalEventList(request *DescribeK8sApiAbnormalEventListRequest) (response *DescribeK8sApiAbnormalEventListResponse, err error) {
    return c.DescribeK8sApiAbnormalEventListWithContext(context.Background(), request)
}

// DescribeK8sApiAbnormalEventList
// This API is used to query the K8sApi abnormal event list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeK8sApiAbnormalEventListWithContext(ctx context.Context, request *DescribeK8sApiAbnormalEventListRequest) (response *DescribeK8sApiAbnormalEventListResponse, err error) {
    if request == nil {
        request = NewDescribeK8sApiAbnormalEventListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeK8sApiAbnormalEventList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeK8sApiAbnormalEventListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeK8sApiAbnormalRuleInfoRequest() (request *DescribeK8sApiAbnormalRuleInfoRequest) {
    request = &DescribeK8sApiAbnormalRuleInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeK8sApiAbnormalRuleInfo")
    
    
    return
}

func NewDescribeK8sApiAbnormalRuleInfoResponse() (response *DescribeK8sApiAbnormalRuleInfoResponse) {
    response = &DescribeK8sApiAbnormalRuleInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeK8sApiAbnormalRuleInfo
// This API is used to query K8sApi abnormal request rule details.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeK8sApiAbnormalRuleInfo(request *DescribeK8sApiAbnormalRuleInfoRequest) (response *DescribeK8sApiAbnormalRuleInfoResponse, err error) {
    return c.DescribeK8sApiAbnormalRuleInfoWithContext(context.Background(), request)
}

// DescribeK8sApiAbnormalRuleInfo
// This API is used to query K8sApi abnormal request rule details.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeK8sApiAbnormalRuleInfoWithContext(ctx context.Context, request *DescribeK8sApiAbnormalRuleInfoRequest) (response *DescribeK8sApiAbnormalRuleInfoResponse, err error) {
    if request == nil {
        request = NewDescribeK8sApiAbnormalRuleInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeK8sApiAbnormalRuleInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeK8sApiAbnormalRuleInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeK8sApiAbnormalRuleListRequest() (request *DescribeK8sApiAbnormalRuleListRequest) {
    request = &DescribeK8sApiAbnormalRuleListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeK8sApiAbnormalRuleList")
    
    
    return
}

func NewDescribeK8sApiAbnormalRuleListResponse() (response *DescribeK8sApiAbnormalRuleListResponse) {
    response = &DescribeK8sApiAbnormalRuleListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeK8sApiAbnormalRuleList
// This API is used to the K8sApi abnormal request rule list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeK8sApiAbnormalRuleList(request *DescribeK8sApiAbnormalRuleListRequest) (response *DescribeK8sApiAbnormalRuleListResponse, err error) {
    return c.DescribeK8sApiAbnormalRuleListWithContext(context.Background(), request)
}

// DescribeK8sApiAbnormalRuleList
// This API is used to the K8sApi abnormal request rule list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeK8sApiAbnormalRuleListWithContext(ctx context.Context, request *DescribeK8sApiAbnormalRuleListRequest) (response *DescribeK8sApiAbnormalRuleListResponse, err error) {
    if request == nil {
        request = NewDescribeK8sApiAbnormalRuleListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeK8sApiAbnormalRuleList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeK8sApiAbnormalRuleListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeK8sApiAbnormalRuleScopeListRequest() (request *DescribeK8sApiAbnormalRuleScopeListRequest) {
    request = &DescribeK8sApiAbnormalRuleScopeListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeK8sApiAbnormalRuleScopeList")
    
    
    return
}

func NewDescribeK8sApiAbnormalRuleScopeListResponse() (response *DescribeK8sApiAbnormalRuleScopeListResponse) {
    response = &DescribeK8sApiAbnormalRuleScopeListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeK8sApiAbnormalRuleScopeList
// This API is used to query rules for K8s API exceptions. 
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeK8sApiAbnormalRuleScopeList(request *DescribeK8sApiAbnormalRuleScopeListRequest) (response *DescribeK8sApiAbnormalRuleScopeListResponse, err error) {
    return c.DescribeK8sApiAbnormalRuleScopeListWithContext(context.Background(), request)
}

// DescribeK8sApiAbnormalRuleScopeList
// This API is used to query rules for K8s API exceptions. 
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeK8sApiAbnormalRuleScopeListWithContext(ctx context.Context, request *DescribeK8sApiAbnormalRuleScopeListRequest) (response *DescribeK8sApiAbnormalRuleScopeListResponse, err error) {
    if request == nil {
        request = NewDescribeK8sApiAbnormalRuleScopeListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeK8sApiAbnormalRuleScopeList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeK8sApiAbnormalRuleScopeListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeK8sApiAbnormalSummaryRequest() (request *DescribeK8sApiAbnormalSummaryRequest) {
    request = &DescribeK8sApiAbnormalSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeK8sApiAbnormalSummary")
    
    
    return
}

func NewDescribeK8sApiAbnormalSummaryResponse() (response *DescribeK8sApiAbnormalSummaryResponse) {
    response = &DescribeK8sApiAbnormalSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeK8sApiAbnormalSummary
// This API is used to query the statistics of K8sApi abnormal events.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeK8sApiAbnormalSummary(request *DescribeK8sApiAbnormalSummaryRequest) (response *DescribeK8sApiAbnormalSummaryResponse, err error) {
    return c.DescribeK8sApiAbnormalSummaryWithContext(context.Background(), request)
}

// DescribeK8sApiAbnormalSummary
// This API is used to query the statistics of K8sApi abnormal events.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeK8sApiAbnormalSummaryWithContext(ctx context.Context, request *DescribeK8sApiAbnormalSummaryRequest) (response *DescribeK8sApiAbnormalSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeK8sApiAbnormalSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeK8sApiAbnormalSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeK8sApiAbnormalSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeK8sApiAbnormalTendencyRequest() (request *DescribeK8sApiAbnormalTendencyRequest) {
    request = &DescribeK8sApiAbnormalTendencyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeK8sApiAbnormalTendency")
    
    
    return
}

func NewDescribeK8sApiAbnormalTendencyResponse() (response *DescribeK8sApiAbnormalTendencyResponse) {
    response = &DescribeK8sApiAbnormalTendencyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeK8sApiAbnormalTendency
// This API is used to query the trend of K8sApi abnormal events.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeK8sApiAbnormalTendency(request *DescribeK8sApiAbnormalTendencyRequest) (response *DescribeK8sApiAbnormalTendencyResponse, err error) {
    return c.DescribeK8sApiAbnormalTendencyWithContext(context.Background(), request)
}

// DescribeK8sApiAbnormalTendency
// This API is used to query the trend of K8sApi abnormal events.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeK8sApiAbnormalTendencyWithContext(ctx context.Context, request *DescribeK8sApiAbnormalTendencyRequest) (response *DescribeK8sApiAbnormalTendencyResponse, err error) {
    if request == nil {
        request = NewDescribeK8sApiAbnormalTendencyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeK8sApiAbnormalTendency require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeK8sApiAbnormalTendencyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogStorageStatisticRequest() (request *DescribeLogStorageStatisticRequest) {
    request = &DescribeLogStorageStatisticRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeLogStorageStatistic")
    
    
    return
}

func NewDescribeLogStorageStatisticResponse() (response *DescribeLogStorageStatisticResponse) {
    response = &DescribeLogStorageStatisticResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLogStorageStatistic
// This API is used to get the statistics of the log search usage.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeLogStorageStatistic(request *DescribeLogStorageStatisticRequest) (response *DescribeLogStorageStatisticResponse, err error) {
    return c.DescribeLogStorageStatisticWithContext(context.Background(), request)
}

// DescribeLogStorageStatistic
// This API is used to get the statistics of the log search usage.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeLogStorageStatisticWithContext(ctx context.Context, request *DescribeLogStorageStatisticRequest) (response *DescribeLogStorageStatisticResponse, err error) {
    if request == nil {
        request = NewDescribeLogStorageStatisticRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogStorageStatistic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogStorageStatisticResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNetworkFirewallAuditRecordRequest() (request *DescribeNetworkFirewallAuditRecordRequest) {
    request = &DescribeNetworkFirewallAuditRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeNetworkFirewallAuditRecord")
    
    
    return
}

func NewDescribeNetworkFirewallAuditRecordResponse() (response *DescribeNetworkFirewallAuditRecordResponse) {
    response = &DescribeNetworkFirewallAuditRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNetworkFirewallAuditRecord
// This API is used to query the list of cluster policy audits.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeNetworkFirewallAuditRecord(request *DescribeNetworkFirewallAuditRecordRequest) (response *DescribeNetworkFirewallAuditRecordResponse, err error) {
    return c.DescribeNetworkFirewallAuditRecordWithContext(context.Background(), request)
}

// DescribeNetworkFirewallAuditRecord
// This API is used to query the list of cluster policy audits.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeNetworkFirewallAuditRecordWithContext(ctx context.Context, request *DescribeNetworkFirewallAuditRecordRequest) (response *DescribeNetworkFirewallAuditRecordResponse, err error) {
    if request == nil {
        request = NewDescribeNetworkFirewallAuditRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNetworkFirewallAuditRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNetworkFirewallAuditRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNetworkFirewallClusterListRequest() (request *DescribeNetworkFirewallClusterListRequest) {
    request = &DescribeNetworkFirewallClusterListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeNetworkFirewallClusterList")
    
    
    return
}

func NewDescribeNetworkFirewallClusterListResponse() (response *DescribeNetworkFirewallClusterListResponse) {
    response = &DescribeNetworkFirewallClusterListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNetworkFirewallClusterList
// This API is used to query the list of clusters.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeNetworkFirewallClusterList(request *DescribeNetworkFirewallClusterListRequest) (response *DescribeNetworkFirewallClusterListResponse, err error) {
    return c.DescribeNetworkFirewallClusterListWithContext(context.Background(), request)
}

// DescribeNetworkFirewallClusterList
// This API is used to query the list of clusters.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeNetworkFirewallClusterListWithContext(ctx context.Context, request *DescribeNetworkFirewallClusterListRequest) (response *DescribeNetworkFirewallClusterListResponse, err error) {
    if request == nil {
        request = NewDescribeNetworkFirewallClusterListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNetworkFirewallClusterList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNetworkFirewallClusterListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNetworkFirewallClusterRefreshStatusRequest() (request *DescribeNetworkFirewallClusterRefreshStatusRequest) {
    request = &DescribeNetworkFirewallClusterRefreshStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeNetworkFirewallClusterRefreshStatus")
    
    
    return
}

func NewDescribeNetworkFirewallClusterRefreshStatusResponse() (response *DescribeNetworkFirewallClusterRefreshStatusResponse) {
    response = &DescribeNetworkFirewallClusterRefreshStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNetworkFirewallClusterRefreshStatus
// This API is used to query the progress of the asset query task in the container network.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeNetworkFirewallClusterRefreshStatus(request *DescribeNetworkFirewallClusterRefreshStatusRequest) (response *DescribeNetworkFirewallClusterRefreshStatusResponse, err error) {
    return c.DescribeNetworkFirewallClusterRefreshStatusWithContext(context.Background(), request)
}

// DescribeNetworkFirewallClusterRefreshStatus
// This API is used to query the progress of the asset query task in the container network.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeNetworkFirewallClusterRefreshStatusWithContext(ctx context.Context, request *DescribeNetworkFirewallClusterRefreshStatusRequest) (response *DescribeNetworkFirewallClusterRefreshStatusResponse, err error) {
    if request == nil {
        request = NewDescribeNetworkFirewallClusterRefreshStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNetworkFirewallClusterRefreshStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNetworkFirewallClusterRefreshStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNetworkFirewallNamespaceLabelListRequest() (request *DescribeNetworkFirewallNamespaceLabelListRequest) {
    request = &DescribeNetworkFirewallNamespaceLabelListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeNetworkFirewallNamespaceLabelList")
    
    
    return
}

func NewDescribeNetworkFirewallNamespaceLabelListResponse() (response *DescribeNetworkFirewallNamespaceLabelListResponse) {
    response = &DescribeNetworkFirewallNamespaceLabelListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNetworkFirewallNamespaceLabelList
// This API is used to query the list of cluster network namespace labels.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeNetworkFirewallNamespaceLabelList(request *DescribeNetworkFirewallNamespaceLabelListRequest) (response *DescribeNetworkFirewallNamespaceLabelListResponse, err error) {
    return c.DescribeNetworkFirewallNamespaceLabelListWithContext(context.Background(), request)
}

// DescribeNetworkFirewallNamespaceLabelList
// This API is used to query the list of cluster network namespace labels.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeNetworkFirewallNamespaceLabelListWithContext(ctx context.Context, request *DescribeNetworkFirewallNamespaceLabelListRequest) (response *DescribeNetworkFirewallNamespaceLabelListResponse, err error) {
    if request == nil {
        request = NewDescribeNetworkFirewallNamespaceLabelListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNetworkFirewallNamespaceLabelList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNetworkFirewallNamespaceLabelListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNetworkFirewallPodLabelsListRequest() (request *DescribeNetworkFirewallPodLabelsListRequest) {
    request = &DescribeNetworkFirewallPodLabelsListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeNetworkFirewallPodLabelsList")
    
    
    return
}

func NewDescribeNetworkFirewallPodLabelsListResponse() (response *DescribeNetworkFirewallPodLabelsListResponse) {
    response = &DescribeNetworkFirewallPodLabelsListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNetworkFirewallPodLabelsList
// This API is used to query cluster network Pod labels.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeNetworkFirewallPodLabelsList(request *DescribeNetworkFirewallPodLabelsListRequest) (response *DescribeNetworkFirewallPodLabelsListResponse, err error) {
    return c.DescribeNetworkFirewallPodLabelsListWithContext(context.Background(), request)
}

// DescribeNetworkFirewallPodLabelsList
// This API is used to query cluster network Pod labels.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeNetworkFirewallPodLabelsListWithContext(ctx context.Context, request *DescribeNetworkFirewallPodLabelsListRequest) (response *DescribeNetworkFirewallPodLabelsListResponse, err error) {
    if request == nil {
        request = NewDescribeNetworkFirewallPodLabelsListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNetworkFirewallPodLabelsList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNetworkFirewallPodLabelsListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNetworkFirewallPolicyDetailRequest() (request *DescribeNetworkFirewallPolicyDetailRequest) {
    request = &DescribeNetworkFirewallPolicyDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeNetworkFirewallPolicyDetail")
    
    
    return
}

func NewDescribeNetworkFirewallPolicyDetailResponse() (response *DescribeNetworkFirewallPolicyDetailResponse) {
    response = &DescribeNetworkFirewallPolicyDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNetworkFirewallPolicyDetail
// This API is used to view the details of a policy in the container network cluster.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeNetworkFirewallPolicyDetail(request *DescribeNetworkFirewallPolicyDetailRequest) (response *DescribeNetworkFirewallPolicyDetailResponse, err error) {
    return c.DescribeNetworkFirewallPolicyDetailWithContext(context.Background(), request)
}

// DescribeNetworkFirewallPolicyDetail
// This API is used to view the details of a policy in the container network cluster.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeNetworkFirewallPolicyDetailWithContext(ctx context.Context, request *DescribeNetworkFirewallPolicyDetailRequest) (response *DescribeNetworkFirewallPolicyDetailResponse, err error) {
    if request == nil {
        request = NewDescribeNetworkFirewallPolicyDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNetworkFirewallPolicyDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNetworkFirewallPolicyDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNetworkFirewallPolicyDiscoverRequest() (request *DescribeNetworkFirewallPolicyDiscoverRequest) {
    request = &DescribeNetworkFirewallPolicyDiscoverRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeNetworkFirewallPolicyDiscover")
    
    
    return
}

func NewDescribeNetworkFirewallPolicyDiscoverResponse() (response *DescribeNetworkFirewallPolicyDiscoverResponse) {
    response = &DescribeNetworkFirewallPolicyDiscoverResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNetworkFirewallPolicyDiscover
// This API is used to query the progress of a network policy sync task in the container network.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeNetworkFirewallPolicyDiscover(request *DescribeNetworkFirewallPolicyDiscoverRequest) (response *DescribeNetworkFirewallPolicyDiscoverResponse, err error) {
    return c.DescribeNetworkFirewallPolicyDiscoverWithContext(context.Background(), request)
}

// DescribeNetworkFirewallPolicyDiscover
// This API is used to query the progress of a network policy sync task in the container network.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeNetworkFirewallPolicyDiscoverWithContext(ctx context.Context, request *DescribeNetworkFirewallPolicyDiscoverRequest) (response *DescribeNetworkFirewallPolicyDiscoverResponse, err error) {
    if request == nil {
        request = NewDescribeNetworkFirewallPolicyDiscoverRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNetworkFirewallPolicyDiscover require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNetworkFirewallPolicyDiscoverResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNetworkFirewallPolicyListRequest() (request *DescribeNetworkFirewallPolicyListRequest) {
    request = &DescribeNetworkFirewallPolicyListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeNetworkFirewallPolicyList")
    
    
    return
}

func NewDescribeNetworkFirewallPolicyListResponse() (response *DescribeNetworkFirewallPolicyListResponse) {
    response = &DescribeNetworkFirewallPolicyListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNetworkFirewallPolicyList
// This API is used to query the list of cluster network policies.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeNetworkFirewallPolicyList(request *DescribeNetworkFirewallPolicyListRequest) (response *DescribeNetworkFirewallPolicyListResponse, err error) {
    return c.DescribeNetworkFirewallPolicyListWithContext(context.Background(), request)
}

// DescribeNetworkFirewallPolicyList
// This API is used to query the list of cluster network policies.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeNetworkFirewallPolicyListWithContext(ctx context.Context, request *DescribeNetworkFirewallPolicyListRequest) (response *DescribeNetworkFirewallPolicyListResponse, err error) {
    if request == nil {
        request = NewDescribeNetworkFirewallPolicyListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNetworkFirewallPolicyList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNetworkFirewallPolicyListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNetworkFirewallPolicyStatusRequest() (request *DescribeNetworkFirewallPolicyStatusRequest) {
    request = &DescribeNetworkFirewallPolicyStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeNetworkFirewallPolicyStatus")
    
    
    return
}

func NewDescribeNetworkFirewallPolicyStatusResponse() (response *DescribeNetworkFirewallPolicyStatusResponse) {
    response = &DescribeNetworkFirewallPolicyStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNetworkFirewallPolicyStatus
// This API is used to query the execution status of a network policy in the container network.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeNetworkFirewallPolicyStatus(request *DescribeNetworkFirewallPolicyStatusRequest) (response *DescribeNetworkFirewallPolicyStatusResponse, err error) {
    return c.DescribeNetworkFirewallPolicyStatusWithContext(context.Background(), request)
}

// DescribeNetworkFirewallPolicyStatus
// This API is used to query the execution status of a network policy in the container network.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeNetworkFirewallPolicyStatusWithContext(ctx context.Context, request *DescribeNetworkFirewallPolicyStatusRequest) (response *DescribeNetworkFirewallPolicyStatusResponse, err error) {
    if request == nil {
        request = NewDescribeNetworkFirewallPolicyStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNetworkFirewallPolicyStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNetworkFirewallPolicyStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNetworkFirewallPolicyYamlDetailRequest() (request *DescribeNetworkFirewallPolicyYamlDetailRequest) {
    request = &DescribeNetworkFirewallPolicyYamlDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeNetworkFirewallPolicyYamlDetail")
    
    
    return
}

func NewDescribeNetworkFirewallPolicyYamlDetailResponse() (response *DescribeNetworkFirewallPolicyYamlDetailResponse) {
    response = &DescribeNetworkFirewallPolicyYamlDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNetworkFirewallPolicyYamlDetail
// This API is used to view the details of a YAML network policy in the container network cluster.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeNetworkFirewallPolicyYamlDetail(request *DescribeNetworkFirewallPolicyYamlDetailRequest) (response *DescribeNetworkFirewallPolicyYamlDetailResponse, err error) {
    return c.DescribeNetworkFirewallPolicyYamlDetailWithContext(context.Background(), request)
}

// DescribeNetworkFirewallPolicyYamlDetail
// This API is used to view the details of a YAML network policy in the container network cluster.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeNetworkFirewallPolicyYamlDetailWithContext(ctx context.Context, request *DescribeNetworkFirewallPolicyYamlDetailRequest) (response *DescribeNetworkFirewallPolicyYamlDetailResponse, err error) {
    if request == nil {
        request = NewDescribeNetworkFirewallPolicyYamlDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNetworkFirewallPolicyYamlDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNetworkFirewallPolicyYamlDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNewestVulRequest() (request *DescribeNewestVulRequest) {
    request = &DescribeNewestVulRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeNewestVul")
    
    
    return
}

func NewDescribeNewestVulResponse() (response *DescribeNewestVulResponse) {
    response = &DescribeNewestVulResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNewestVul
// This API is used to query the latest list of vulnerabilities.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeNewestVul(request *DescribeNewestVulRequest) (response *DescribeNewestVulResponse, err error) {
    return c.DescribeNewestVulWithContext(context.Background(), request)
}

// DescribeNewestVul
// This API is used to query the latest list of vulnerabilities.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeNewestVulWithContext(ctx context.Context, request *DescribeNewestVulRequest) (response *DescribeNewestVulResponse, err error) {
    if request == nil {
        request = NewDescribeNewestVulRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNewestVul require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNewestVulResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePostPayDetailRequest() (request *DescribePostPayDetailRequest) {
    request = &DescribePostPayDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribePostPayDetail")
    
    
    return
}

func NewDescribePostPayDetailResponse() (response *DescribePostPayDetailResponse) {
    response = &DescribePostPayDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePostPayDetail
// This API is used to query the pay-as-you-go billing details.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
func (c *Client) DescribePostPayDetail(request *DescribePostPayDetailRequest) (response *DescribePostPayDetailResponse, err error) {
    return c.DescribePostPayDetailWithContext(context.Background(), request)
}

// DescribePostPayDetail
// This API is used to query the pay-as-you-go billing details.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
func (c *Client) DescribePostPayDetailWithContext(ctx context.Context, request *DescribePostPayDetailRequest) (response *DescribePostPayDetailResponse, err error) {
    if request == nil {
        request = NewDescribePostPayDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePostPayDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePostPayDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProVersionInfoRequest() (request *DescribeProVersionInfoRequest) {
    request = &DescribeProVersionInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeProVersionInfo")
    
    
    return
}

func NewDescribeProVersionInfoResponse() (response *DescribeProVersionInfoResponse) {
    response = &DescribeProVersionInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProVersionInfo
// This API is used to check whether the Pro Edition needs to be purchased.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRROLENOTEXIST = "InternalError.ErrRoleNotExist"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DescribeProVersionInfo(request *DescribeProVersionInfoRequest) (response *DescribeProVersionInfoResponse, err error) {
    return c.DescribeProVersionInfoWithContext(context.Background(), request)
}

// DescribeProVersionInfo
// This API is used to check whether the Pro Edition needs to be purchased.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRROLENOTEXIST = "InternalError.ErrRoleNotExist"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DescribeProVersionInfoWithContext(ctx context.Context, request *DescribeProVersionInfoRequest) (response *DescribeProVersionInfoResponse, err error) {
    if request == nil {
        request = NewDescribeProVersionInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProVersionInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProVersionInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePromotionActivityRequest() (request *DescribePromotionActivityRequest) {
    request = &DescribePromotionActivityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribePromotionActivity")
    
    
    return
}

func NewDescribePromotionActivityResponse() (response *DescribePromotionActivityResponse) {
    response = &DescribePromotionActivityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePromotionActivity
// This API is used to query promotions.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePromotionActivity(request *DescribePromotionActivityRequest) (response *DescribePromotionActivityResponse, err error) {
    return c.DescribePromotionActivityWithContext(context.Background(), request)
}

// DescribePromotionActivity
// This API is used to query promotions.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePromotionActivityWithContext(ctx context.Context, request *DescribePromotionActivityRequest) (response *DescribePromotionActivityResponse, err error) {
    if request == nil {
        request = NewDescribePromotionActivityRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePromotionActivity require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePromotionActivityResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePublicKeyRequest() (request *DescribePublicKeyRequest) {
    request = &DescribePublicKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribePublicKey")
    
    
    return
}

func NewDescribePublicKeyResponse() (response *DescribePublicKeyResponse) {
    response = &DescribePublicKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePublicKey
// This API is used to get the public key.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePublicKey(request *DescribePublicKeyRequest) (response *DescribePublicKeyResponse, err error) {
    return c.DescribePublicKeyWithContext(context.Background(), request)
}

// DescribePublicKey
// This API is used to get the public key.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePublicKeyWithContext(ctx context.Context, request *DescribePublicKeyRequest) (response *DescribePublicKeyResponse, err error) {
    if request == nil {
        request = NewDescribePublicKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePublicKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePublicKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePurchaseStateInfoRequest() (request *DescribePurchaseStateInfoRequest) {
    request = &DescribePurchaseStateInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribePurchaseStateInfo")
    
    
    return
}

func NewDescribePurchaseStateInfoResponse() (response *DescribePurchaseStateInfoResponse) {
    response = &DescribePurchaseStateInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePurchaseStateInfo
// This API is used to check whether TCSS is purchased.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRROLENOTEXIST = "InternalError.ErrRoleNotExist"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePurchaseStateInfo(request *DescribePurchaseStateInfoRequest) (response *DescribePurchaseStateInfoResponse, err error) {
    return c.DescribePurchaseStateInfoWithContext(context.Background(), request)
}

// DescribePurchaseStateInfo
// This API is used to check whether TCSS is purchased.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRROLENOTEXIST = "InternalError.ErrRoleNotExist"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePurchaseStateInfoWithContext(ctx context.Context, request *DescribePurchaseStateInfoRequest) (response *DescribePurchaseStateInfoResponse, err error) {
    if request == nil {
        request = NewDescribePurchaseStateInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePurchaseStateInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePurchaseStateInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRefreshTaskRequest() (request *DescribeRefreshTaskRequest) {
    request = &DescribeRefreshTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeRefreshTask")
    
    
    return
}

func NewDescribeRefreshTaskResponse() (response *DescribeRefreshTaskResponse) {
    response = &DescribeRefreshTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRefreshTask
// This API is used to query a refresh task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeRefreshTask(request *DescribeRefreshTaskRequest) (response *DescribeRefreshTaskResponse, err error) {
    return c.DescribeRefreshTaskWithContext(context.Background(), request)
}

// DescribeRefreshTask
// This API is used to query a refresh task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeRefreshTaskWithContext(ctx context.Context, request *DescribeRefreshTaskRequest) (response *DescribeRefreshTaskResponse, err error) {
    if request == nil {
        request = NewDescribeRefreshTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRefreshTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRefreshTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReverseShellDetailRequest() (request *DescribeReverseShellDetailRequest) {
    request = &DescribeReverseShellDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeReverseShellDetail")
    
    
    return
}

func NewDescribeReverseShellDetailResponse() (response *DescribeReverseShellDetailResponse) {
    response = &DescribeReverseShellDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeReverseShellDetail
// This API is used to query the details of a reverse shell event at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATANOTFOUND = "InvalidParameterValue.DataNotFound"
func (c *Client) DescribeReverseShellDetail(request *DescribeReverseShellDetailRequest) (response *DescribeReverseShellDetailResponse, err error) {
    return c.DescribeReverseShellDetailWithContext(context.Background(), request)
}

// DescribeReverseShellDetail
// This API is used to query the details of a reverse shell event at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATANOTFOUND = "InvalidParameterValue.DataNotFound"
func (c *Client) DescribeReverseShellDetailWithContext(ctx context.Context, request *DescribeReverseShellDetailRequest) (response *DescribeReverseShellDetailResponse, err error) {
    if request == nil {
        request = NewDescribeReverseShellDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReverseShellDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReverseShellDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReverseShellEventsRequest() (request *DescribeReverseShellEventsRequest) {
    request = &DescribeReverseShellEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeReverseShellEvents")
    
    
    return
}

func NewDescribeReverseShellEventsResponse() (response *DescribeReverseShellEventsResponse) {
    response = &DescribeReverseShellEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeReverseShellEvents
// This API is used to query the list of reverse shell events at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeReverseShellEvents(request *DescribeReverseShellEventsRequest) (response *DescribeReverseShellEventsResponse, err error) {
    return c.DescribeReverseShellEventsWithContext(context.Background(), request)
}

// DescribeReverseShellEvents
// This API is used to query the list of reverse shell events at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeReverseShellEventsWithContext(ctx context.Context, request *DescribeReverseShellEventsRequest) (response *DescribeReverseShellEventsResponse, err error) {
    if request == nil {
        request = NewDescribeReverseShellEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReverseShellEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReverseShellEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReverseShellEventsExportRequest() (request *DescribeReverseShellEventsExportRequest) {
    request = &DescribeReverseShellEventsExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeReverseShellEventsExport")
    
    
    return
}

func NewDescribeReverseShellEventsExportResponse() (response *DescribeReverseShellEventsExportResponse) {
    response = &DescribeReverseShellEventsExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeReverseShellEventsExport
// This API is used to query and export the list of reverse shell events at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeReverseShellEventsExport(request *DescribeReverseShellEventsExportRequest) (response *DescribeReverseShellEventsExportResponse, err error) {
    return c.DescribeReverseShellEventsExportWithContext(context.Background(), request)
}

// DescribeReverseShellEventsExport
// This API is used to query and export the list of reverse shell events at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeReverseShellEventsExportWithContext(ctx context.Context, request *DescribeReverseShellEventsExportRequest) (response *DescribeReverseShellEventsExportResponse, err error) {
    if request == nil {
        request = NewDescribeReverseShellEventsExportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReverseShellEventsExport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReverseShellEventsExportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReverseShellWhiteListDetailRequest() (request *DescribeReverseShellWhiteListDetailRequest) {
    request = &DescribeReverseShellWhiteListDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeReverseShellWhiteListDetail")
    
    
    return
}

func NewDescribeReverseShellWhiteListDetailResponse() (response *DescribeReverseShellWhiteListDetailResponse) {
    response = &DescribeReverseShellWhiteListDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeReverseShellWhiteListDetail
// This API is used to query the details of the allowlist of reverse shells at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATANOTFOUND = "InvalidParameterValue.DataNotFound"
func (c *Client) DescribeReverseShellWhiteListDetail(request *DescribeReverseShellWhiteListDetailRequest) (response *DescribeReverseShellWhiteListDetailResponse, err error) {
    return c.DescribeReverseShellWhiteListDetailWithContext(context.Background(), request)
}

// DescribeReverseShellWhiteListDetail
// This API is used to query the details of the allowlist of reverse shells at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATANOTFOUND = "InvalidParameterValue.DataNotFound"
func (c *Client) DescribeReverseShellWhiteListDetailWithContext(ctx context.Context, request *DescribeReverseShellWhiteListDetailRequest) (response *DescribeReverseShellWhiteListDetailResponse, err error) {
    if request == nil {
        request = NewDescribeReverseShellWhiteListDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReverseShellWhiteListDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReverseShellWhiteListDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReverseShellWhiteListsRequest() (request *DescribeReverseShellWhiteListsRequest) {
    request = &DescribeReverseShellWhiteListsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeReverseShellWhiteLists")
    
    
    return
}

func NewDescribeReverseShellWhiteListsResponse() (response *DescribeReverseShellWhiteListsResponse) {
    response = &DescribeReverseShellWhiteListsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeReverseShellWhiteLists
// This API is used to query the allowlist of reverse shells at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeReverseShellWhiteLists(request *DescribeReverseShellWhiteListsRequest) (response *DescribeReverseShellWhiteListsResponse, err error) {
    return c.DescribeReverseShellWhiteListsWithContext(context.Background(), request)
}

// DescribeReverseShellWhiteLists
// This API is used to query the allowlist of reverse shells at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeReverseShellWhiteListsWithContext(ctx context.Context, request *DescribeReverseShellWhiteListsRequest) (response *DescribeReverseShellWhiteListsResponse, err error) {
    if request == nil {
        request = NewDescribeReverseShellWhiteListsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReverseShellWhiteLists require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReverseShellWhiteListsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskListRequest() (request *DescribeRiskListRequest) {
    request = &DescribeRiskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeRiskList")
    
    
    return
}

func NewDescribeRiskListResponse() (response *DescribeRiskListResponse) {
    response = &DescribeRiskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRiskList
// This API is used to query the list of risk items identified in the last task and filter them by special field.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeRiskList(request *DescribeRiskListRequest) (response *DescribeRiskListResponse, err error) {
    return c.DescribeRiskListWithContext(context.Background(), request)
}

// DescribeRiskList
// This API is used to query the list of risk items identified in the last task and filter them by special field.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeRiskListWithContext(ctx context.Context, request *DescribeRiskListRequest) (response *DescribeRiskListResponse, err error) {
    if request == nil {
        request = NewDescribeRiskListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskSyscallDetailRequest() (request *DescribeRiskSyscallDetailRequest) {
    request = &DescribeRiskSyscallDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeRiskSyscallDetail")
    
    
    return
}

func NewDescribeRiskSyscallDetailResponse() (response *DescribeRiskSyscallDetailResponse) {
    response = &DescribeRiskSyscallDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRiskSyscallDetail
// This API is used to query the details of a high-risk syscall event.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATANOTFOUND = "InvalidParameterValue.DataNotFound"
func (c *Client) DescribeRiskSyscallDetail(request *DescribeRiskSyscallDetailRequest) (response *DescribeRiskSyscallDetailResponse, err error) {
    return c.DescribeRiskSyscallDetailWithContext(context.Background(), request)
}

// DescribeRiskSyscallDetail
// This API is used to query the details of a high-risk syscall event.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATANOTFOUND = "InvalidParameterValue.DataNotFound"
func (c *Client) DescribeRiskSyscallDetailWithContext(ctx context.Context, request *DescribeRiskSyscallDetailRequest) (response *DescribeRiskSyscallDetailResponse, err error) {
    if request == nil {
        request = NewDescribeRiskSyscallDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskSyscallDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskSyscallDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskSyscallEventsRequest() (request *DescribeRiskSyscallEventsRequest) {
    request = &DescribeRiskSyscallEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeRiskSyscallEvents")
    
    
    return
}

func NewDescribeRiskSyscallEventsResponse() (response *DescribeRiskSyscallEventsResponse) {
    response = &DescribeRiskSyscallEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRiskSyscallEvents
// This API is used to query the list of high-risk syscalls at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRiskSyscallEvents(request *DescribeRiskSyscallEventsRequest) (response *DescribeRiskSyscallEventsResponse, err error) {
    return c.DescribeRiskSyscallEventsWithContext(context.Background(), request)
}

// DescribeRiskSyscallEvents
// This API is used to query the list of high-risk syscalls at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRiskSyscallEventsWithContext(ctx context.Context, request *DescribeRiskSyscallEventsRequest) (response *DescribeRiskSyscallEventsResponse, err error) {
    if request == nil {
        request = NewDescribeRiskSyscallEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskSyscallEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskSyscallEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskSyscallEventsExportRequest() (request *DescribeRiskSyscallEventsExportRequest) {
    request = &DescribeRiskSyscallEventsExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeRiskSyscallEventsExport")
    
    
    return
}

func NewDescribeRiskSyscallEventsExportResponse() (response *DescribeRiskSyscallEventsExportResponse) {
    response = &DescribeRiskSyscallEventsExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRiskSyscallEventsExport
// This API is used to export the list of high-risk syscalls at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRiskSyscallEventsExport(request *DescribeRiskSyscallEventsExportRequest) (response *DescribeRiskSyscallEventsExportResponse, err error) {
    return c.DescribeRiskSyscallEventsExportWithContext(context.Background(), request)
}

// DescribeRiskSyscallEventsExport
// This API is used to export the list of high-risk syscalls at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRiskSyscallEventsExportWithContext(ctx context.Context, request *DescribeRiskSyscallEventsExportRequest) (response *DescribeRiskSyscallEventsExportResponse, err error) {
    if request == nil {
        request = NewDescribeRiskSyscallEventsExportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskSyscallEventsExport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskSyscallEventsExportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskSyscallNamesRequest() (request *DescribeRiskSyscallNamesRequest) {
    request = &DescribeRiskSyscallNamesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeRiskSyscallNames")
    
    
    return
}

func NewDescribeRiskSyscallNamesResponse() (response *DescribeRiskSyscallNamesResponse) {
    response = &DescribeRiskSyscallNamesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRiskSyscallNames
// This API is used to query the list of names of high-risk syscalls at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRiskSyscallNames(request *DescribeRiskSyscallNamesRequest) (response *DescribeRiskSyscallNamesResponse, err error) {
    return c.DescribeRiskSyscallNamesWithContext(context.Background(), request)
}

// DescribeRiskSyscallNames
// This API is used to query the list of names of high-risk syscalls at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRiskSyscallNamesWithContext(ctx context.Context, request *DescribeRiskSyscallNamesRequest) (response *DescribeRiskSyscallNamesResponse, err error) {
    if request == nil {
        request = NewDescribeRiskSyscallNamesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskSyscallNames require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskSyscallNamesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskSyscallWhiteListDetailRequest() (request *DescribeRiskSyscallWhiteListDetailRequest) {
    request = &DescribeRiskSyscallWhiteListDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeRiskSyscallWhiteListDetail")
    
    
    return
}

func NewDescribeRiskSyscallWhiteListDetailResponse() (response *DescribeRiskSyscallWhiteListDetailResponse) {
    response = &DescribeRiskSyscallWhiteListDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRiskSyscallWhiteListDetail
// This API is used to query the details of the allowlist of high-risk syscalls at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATANOTFOUND = "InvalidParameterValue.DataNotFound"
func (c *Client) DescribeRiskSyscallWhiteListDetail(request *DescribeRiskSyscallWhiteListDetailRequest) (response *DescribeRiskSyscallWhiteListDetailResponse, err error) {
    return c.DescribeRiskSyscallWhiteListDetailWithContext(context.Background(), request)
}

// DescribeRiskSyscallWhiteListDetail
// This API is used to query the details of the allowlist of high-risk syscalls at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATANOTFOUND = "InvalidParameterValue.DataNotFound"
func (c *Client) DescribeRiskSyscallWhiteListDetailWithContext(ctx context.Context, request *DescribeRiskSyscallWhiteListDetailRequest) (response *DescribeRiskSyscallWhiteListDetailResponse, err error) {
    if request == nil {
        request = NewDescribeRiskSyscallWhiteListDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskSyscallWhiteListDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskSyscallWhiteListDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskSyscallWhiteListsRequest() (request *DescribeRiskSyscallWhiteListsRequest) {
    request = &DescribeRiskSyscallWhiteListsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeRiskSyscallWhiteLists")
    
    
    return
}

func NewDescribeRiskSyscallWhiteListsResponse() (response *DescribeRiskSyscallWhiteListsResponse) {
    response = &DescribeRiskSyscallWhiteListsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRiskSyscallWhiteLists
// This API is used to query the allowlist of high-risk syscalls at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeRiskSyscallWhiteLists(request *DescribeRiskSyscallWhiteListsRequest) (response *DescribeRiskSyscallWhiteListsResponse, err error) {
    return c.DescribeRiskSyscallWhiteListsWithContext(context.Background(), request)
}

// DescribeRiskSyscallWhiteLists
// This API is used to query the allowlist of high-risk syscalls at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeRiskSyscallWhiteListsWithContext(ctx context.Context, request *DescribeRiskSyscallWhiteListsRequest) (response *DescribeRiskSyscallWhiteListsResponse, err error) {
    if request == nil {
        request = NewDescribeRiskSyscallWhiteListsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskSyscallWhiteLists require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskSyscallWhiteListsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScanIgnoreVulListRequest() (request *DescribeScanIgnoreVulListRequest) {
    request = &DescribeScanIgnoreVulListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeScanIgnoreVulList")
    
    
    return
}

func NewDescribeScanIgnoreVulListResponse() (response *DescribeScanIgnoreVulListResponse) {
    response = &DescribeScanIgnoreVulListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeScanIgnoreVulList
// This API is used to query the list of vulnerabilities ignored in a scan.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeScanIgnoreVulList(request *DescribeScanIgnoreVulListRequest) (response *DescribeScanIgnoreVulListResponse, err error) {
    return c.DescribeScanIgnoreVulListWithContext(context.Background(), request)
}

// DescribeScanIgnoreVulList
// This API is used to query the list of vulnerabilities ignored in a scan.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeScanIgnoreVulListWithContext(ctx context.Context, request *DescribeScanIgnoreVulListRequest) (response *DescribeScanIgnoreVulListResponse, err error) {
    if request == nil {
        request = NewDescribeScanIgnoreVulListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScanIgnoreVulList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScanIgnoreVulListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSearchExportListRequest() (request *DescribeSearchExportListRequest) {
    request = &DescribeSearchExportListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeSearchExportList")
    
    
    return
}

func NewDescribeSearchExportListResponse() (response *DescribeSearchExportListResponse) {
    response = &DescribeSearchExportListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSearchExportList
// This API is used to export the list of ES query files.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSearchExportList(request *DescribeSearchExportListRequest) (response *DescribeSearchExportListResponse, err error) {
    return c.DescribeSearchExportListWithContext(context.Background(), request)
}

// DescribeSearchExportList
// This API is used to export the list of ES query files.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSearchExportListWithContext(ctx context.Context, request *DescribeSearchExportListRequest) (response *DescribeSearchExportListResponse, err error) {
    if request == nil {
        request = NewDescribeSearchExportListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSearchExportList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSearchExportListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSearchLogsRequest() (request *DescribeSearchLogsRequest) {
    request = &DescribeSearchLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeSearchLogs")
    
    
    return
}

func NewDescribeSearchLogsResponse() (response *DescribeSearchLogsResponse) {
    response = &DescribeSearchLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSearchLogs
// This API is used to get historical search records.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSearchLogs(request *DescribeSearchLogsRequest) (response *DescribeSearchLogsResponse, err error) {
    return c.DescribeSearchLogsWithContext(context.Background(), request)
}

// DescribeSearchLogs
// This API is used to get historical search records.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSearchLogsWithContext(ctx context.Context, request *DescribeSearchLogsRequest) (response *DescribeSearchLogsResponse, err error) {
    if request == nil {
        request = NewDescribeSearchLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSearchLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSearchLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSearchTemplatesRequest() (request *DescribeSearchTemplatesRequest) {
    request = &DescribeSearchTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeSearchTemplates")
    
    
    return
}

func NewDescribeSearchTemplatesResponse() (response *DescribeSearchTemplatesResponse) {
    response = &DescribeSearchTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSearchTemplates
// This API is used to get the list of search templates.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSearchTemplates(request *DescribeSearchTemplatesRequest) (response *DescribeSearchTemplatesResponse, err error) {
    return c.DescribeSearchTemplatesWithContext(context.Background(), request)
}

// DescribeSearchTemplates
// This API is used to get the list of search templates.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSearchTemplatesWithContext(ctx context.Context, request *DescribeSearchTemplatesRequest) (response *DescribeSearchTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeSearchTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSearchTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSearchTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecEventsTendencyRequest() (request *DescribeSecEventsTendencyRequest) {
    request = &DescribeSecEventsTendencyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeSecEventsTendency")
    
    
    return
}

func NewDescribeSecEventsTendencyResponse() (response *DescribeSecEventsTendencyResponse) {
    response = &DescribeSecEventsTendencyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecEventsTendency
// This API is used to query the trend of security events at runtime.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATARANGE = "InvalidParameterValue.DataRange"
func (c *Client) DescribeSecEventsTendency(request *DescribeSecEventsTendencyRequest) (response *DescribeSecEventsTendencyResponse, err error) {
    return c.DescribeSecEventsTendencyWithContext(context.Background(), request)
}

// DescribeSecEventsTendency
// This API is used to query the trend of security events at runtime.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATARANGE = "InvalidParameterValue.DataRange"
func (c *Client) DescribeSecEventsTendencyWithContext(ctx context.Context, request *DescribeSecEventsTendencyRequest) (response *DescribeSecEventsTendencyResponse, err error) {
    if request == nil {
        request = NewDescribeSecEventsTendencyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecEventsTendency require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecEventsTendencyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecLogAlertMsgRequest() (request *DescribeSecLogAlertMsgRequest) {
    request = &DescribeSecLogAlertMsgRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeSecLogAlertMsg")
    
    
    return
}

func NewDescribeSecLogAlertMsgResponse() (response *DescribeSecLogAlertMsgResponse) {
    response = &DescribeSecLogAlertMsgResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecLogAlertMsg
// This API is used to query a security log alert message.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSecLogAlertMsg(request *DescribeSecLogAlertMsgRequest) (response *DescribeSecLogAlertMsgResponse, err error) {
    return c.DescribeSecLogAlertMsgWithContext(context.Background(), request)
}

// DescribeSecLogAlertMsg
// This API is used to query a security log alert message.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSecLogAlertMsgWithContext(ctx context.Context, request *DescribeSecLogAlertMsgRequest) (response *DescribeSecLogAlertMsgResponse, err error) {
    if request == nil {
        request = NewDescribeSecLogAlertMsgRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecLogAlertMsg require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecLogAlertMsgResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecLogCleanSettingInfoRequest() (request *DescribeSecLogCleanSettingInfoRequest) {
    request = &DescribeSecLogCleanSettingInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeSecLogCleanSettingInfo")
    
    
    return
}

func NewDescribeSecLogCleanSettingInfoResponse() (response *DescribeSecLogCleanSettingInfoResponse) {
    response = &DescribeSecLogCleanSettingInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecLogCleanSettingInfo
// This API is used to query the settings of security log cleanup.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSecLogCleanSettingInfo(request *DescribeSecLogCleanSettingInfoRequest) (response *DescribeSecLogCleanSettingInfoResponse, err error) {
    return c.DescribeSecLogCleanSettingInfoWithContext(context.Background(), request)
}

// DescribeSecLogCleanSettingInfo
// This API is used to query the settings of security log cleanup.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSecLogCleanSettingInfoWithContext(ctx context.Context, request *DescribeSecLogCleanSettingInfoRequest) (response *DescribeSecLogCleanSettingInfoResponse, err error) {
    if request == nil {
        request = NewDescribeSecLogCleanSettingInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecLogCleanSettingInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecLogCleanSettingInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecLogDeliveryClsOptionsRequest() (request *DescribeSecLogDeliveryClsOptionsRequest) {
    request = &DescribeSecLogDeliveryClsOptionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeSecLogDeliveryClsOptions")
    
    
    return
}

func NewDescribeSecLogDeliveryClsOptionsResponse() (response *DescribeSecLogDeliveryClsOptionsResponse) {
    response = &DescribeSecLogDeliveryClsOptionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecLogDeliveryClsOptions
// This API is used to query the options of security log delivery to CLS.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSecLogDeliveryClsOptions(request *DescribeSecLogDeliveryClsOptionsRequest) (response *DescribeSecLogDeliveryClsOptionsResponse, err error) {
    return c.DescribeSecLogDeliveryClsOptionsWithContext(context.Background(), request)
}

// DescribeSecLogDeliveryClsOptions
// This API is used to query the options of security log delivery to CLS.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSecLogDeliveryClsOptionsWithContext(ctx context.Context, request *DescribeSecLogDeliveryClsOptionsRequest) (response *DescribeSecLogDeliveryClsOptionsResponse, err error) {
    if request == nil {
        request = NewDescribeSecLogDeliveryClsOptionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecLogDeliveryClsOptions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecLogDeliveryClsOptionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecLogDeliveryClsSettingRequest() (request *DescribeSecLogDeliveryClsSettingRequest) {
    request = &DescribeSecLogDeliveryClsSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeSecLogDeliveryClsSetting")
    
    
    return
}

func NewDescribeSecLogDeliveryClsSettingResponse() (response *DescribeSecLogDeliveryClsSettingResponse) {
    response = &DescribeSecLogDeliveryClsSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecLogDeliveryClsSetting
// This API is used to query the settings of security log delivery to CLS.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSecLogDeliveryClsSetting(request *DescribeSecLogDeliveryClsSettingRequest) (response *DescribeSecLogDeliveryClsSettingResponse, err error) {
    return c.DescribeSecLogDeliveryClsSettingWithContext(context.Background(), request)
}

// DescribeSecLogDeliveryClsSetting
// This API is used to query the settings of security log delivery to CLS.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSecLogDeliveryClsSettingWithContext(ctx context.Context, request *DescribeSecLogDeliveryClsSettingRequest) (response *DescribeSecLogDeliveryClsSettingResponse, err error) {
    if request == nil {
        request = NewDescribeSecLogDeliveryClsSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecLogDeliveryClsSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecLogDeliveryClsSettingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecLogDeliveryKafkaOptionsRequest() (request *DescribeSecLogDeliveryKafkaOptionsRequest) {
    request = &DescribeSecLogDeliveryKafkaOptionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeSecLogDeliveryKafkaOptions")
    
    
    return
}

func NewDescribeSecLogDeliveryKafkaOptionsResponse() (response *DescribeSecLogDeliveryKafkaOptionsResponse) {
    response = &DescribeSecLogDeliveryKafkaOptionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecLogDeliveryKafkaOptions
// This API is used to query the options of security log delivery to Kafka.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSecLogDeliveryKafkaOptions(request *DescribeSecLogDeliveryKafkaOptionsRequest) (response *DescribeSecLogDeliveryKafkaOptionsResponse, err error) {
    return c.DescribeSecLogDeliveryKafkaOptionsWithContext(context.Background(), request)
}

// DescribeSecLogDeliveryKafkaOptions
// This API is used to query the options of security log delivery to Kafka.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSecLogDeliveryKafkaOptionsWithContext(ctx context.Context, request *DescribeSecLogDeliveryKafkaOptionsRequest) (response *DescribeSecLogDeliveryKafkaOptionsResponse, err error) {
    if request == nil {
        request = NewDescribeSecLogDeliveryKafkaOptionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecLogDeliveryKafkaOptions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecLogDeliveryKafkaOptionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecLogDeliveryKafkaSettingRequest() (request *DescribeSecLogDeliveryKafkaSettingRequest) {
    request = &DescribeSecLogDeliveryKafkaSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeSecLogDeliveryKafkaSetting")
    
    
    return
}

func NewDescribeSecLogDeliveryKafkaSettingResponse() (response *DescribeSecLogDeliveryKafkaSettingResponse) {
    response = &DescribeSecLogDeliveryKafkaSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecLogDeliveryKafkaSetting
// This API is used to query the settings of security log delivery to Kafka.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSecLogDeliveryKafkaSetting(request *DescribeSecLogDeliveryKafkaSettingRequest) (response *DescribeSecLogDeliveryKafkaSettingResponse, err error) {
    return c.DescribeSecLogDeliveryKafkaSettingWithContext(context.Background(), request)
}

// DescribeSecLogDeliveryKafkaSetting
// This API is used to query the settings of security log delivery to Kafka.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSecLogDeliveryKafkaSettingWithContext(ctx context.Context, request *DescribeSecLogDeliveryKafkaSettingRequest) (response *DescribeSecLogDeliveryKafkaSettingResponse, err error) {
    if request == nil {
        request = NewDescribeSecLogDeliveryKafkaSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecLogDeliveryKafkaSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecLogDeliveryKafkaSettingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecLogJoinObjectListRequest() (request *DescribeSecLogJoinObjectListRequest) {
    request = &DescribeSecLogJoinObjectListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeSecLogJoinObjectList")
    
    
    return
}

func NewDescribeSecLogJoinObjectListResponse() (response *DescribeSecLogJoinObjectListResponse) {
    response = &DescribeSecLogJoinObjectListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecLogJoinObjectList
// This API is used to query the list of accessed security log objects.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSecLogJoinObjectList(request *DescribeSecLogJoinObjectListRequest) (response *DescribeSecLogJoinObjectListResponse, err error) {
    return c.DescribeSecLogJoinObjectListWithContext(context.Background(), request)
}

// DescribeSecLogJoinObjectList
// This API is used to query the list of accessed security log objects.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSecLogJoinObjectListWithContext(ctx context.Context, request *DescribeSecLogJoinObjectListRequest) (response *DescribeSecLogJoinObjectListResponse, err error) {
    if request == nil {
        request = NewDescribeSecLogJoinObjectListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecLogJoinObjectList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecLogJoinObjectListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecLogJoinTypeListRequest() (request *DescribeSecLogJoinTypeListRequest) {
    request = &DescribeSecLogJoinTypeListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeSecLogJoinTypeList")
    
    
    return
}

func NewDescribeSecLogJoinTypeListResponse() (response *DescribeSecLogJoinTypeListResponse) {
    response = &DescribeSecLogJoinTypeListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecLogJoinTypeList
// This API is used to query the list of security log access types.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSecLogJoinTypeList(request *DescribeSecLogJoinTypeListRequest) (response *DescribeSecLogJoinTypeListResponse, err error) {
    return c.DescribeSecLogJoinTypeListWithContext(context.Background(), request)
}

// DescribeSecLogJoinTypeList
// This API is used to query the list of security log access types.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSecLogJoinTypeListWithContext(ctx context.Context, request *DescribeSecLogJoinTypeListRequest) (response *DescribeSecLogJoinTypeListResponse, err error) {
    if request == nil {
        request = NewDescribeSecLogJoinTypeListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecLogJoinTypeList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecLogJoinTypeListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecLogKafkaUINRequest() (request *DescribeSecLogKafkaUINRequest) {
    request = &DescribeSecLogKafkaUINRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeSecLogKafkaUIN")
    
    
    return
}

func NewDescribeSecLogKafkaUINResponse() (response *DescribeSecLogKafkaUINResponse) {
    response = &DescribeSecLogKafkaUINResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecLogKafkaUIN
// This API is used to query the UIN of a Kafka security log.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSecLogKafkaUIN(request *DescribeSecLogKafkaUINRequest) (response *DescribeSecLogKafkaUINResponse, err error) {
    return c.DescribeSecLogKafkaUINWithContext(context.Background(), request)
}

// DescribeSecLogKafkaUIN
// This API is used to query the UIN of a Kafka security log.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSecLogKafkaUINWithContext(ctx context.Context, request *DescribeSecLogKafkaUINRequest) (response *DescribeSecLogKafkaUINResponse, err error) {
    if request == nil {
        request = NewDescribeSecLogKafkaUINRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecLogKafkaUIN require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecLogKafkaUINResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecLogVasInfoRequest() (request *DescribeSecLogVasInfoRequest) {
    request = &DescribeSecLogVasInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeSecLogVasInfo")
    
    
    return
}

func NewDescribeSecLogVasInfoResponse() (response *DescribeSecLogVasInfoResponse) {
    response = &DescribeSecLogVasInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecLogVasInfo
// This API is used to query the information of the security log product.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSecLogVasInfo(request *DescribeSecLogVasInfoRequest) (response *DescribeSecLogVasInfoResponse, err error) {
    return c.DescribeSecLogVasInfoWithContext(context.Background(), request)
}

// DescribeSecLogVasInfo
// This API is used to query the information of the security log product.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSecLogVasInfoWithContext(ctx context.Context, request *DescribeSecLogVasInfoRequest) (response *DescribeSecLogVasInfoResponse, err error) {
    if request == nil {
        request = NewDescribeSecLogVasInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecLogVasInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecLogVasInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSupportDefenceVulRequest() (request *DescribeSupportDefenceVulRequest) {
    request = &DescribeSupportDefenceVulRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeSupportDefenceVul")
    
    
    return
}

func NewDescribeSupportDefenceVulResponse() (response *DescribeSupportDefenceVulResponse) {
    response = &DescribeSupportDefenceVulResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSupportDefenceVul
// This API is used to query the list of vulnerabilities that can be prevented
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeSupportDefenceVul(request *DescribeSupportDefenceVulRequest) (response *DescribeSupportDefenceVulResponse, err error) {
    return c.DescribeSupportDefenceVulWithContext(context.Background(), request)
}

// DescribeSupportDefenceVul
// This API is used to query the list of vulnerabilities that can be prevented
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeSupportDefenceVulWithContext(ctx context.Context, request *DescribeSupportDefenceVulRequest) (response *DescribeSupportDefenceVulResponse, err error) {
    if request == nil {
        request = NewDescribeSupportDefenceVulRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSupportDefenceVul require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSupportDefenceVulResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSystemVulListRequest() (request *DescribeSystemVulListRequest) {
    request = &DescribeSystemVulListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeSystemVulList")
    
    
    return
}

func NewDescribeSystemVulListResponse() (response *DescribeSystemVulListResponse) {
    response = &DescribeSystemVulListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSystemVulList
// This API is used to query the list of system vulnerabilities.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSystemVulList(request *DescribeSystemVulListRequest) (response *DescribeSystemVulListResponse, err error) {
    return c.DescribeSystemVulListWithContext(context.Background(), request)
}

// DescribeSystemVulList
// This API is used to query the list of system vulnerabilities.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSystemVulListWithContext(ctx context.Context, request *DescribeSystemVulListRequest) (response *DescribeSystemVulListResponse, err error) {
    if request == nil {
        request = NewDescribeSystemVulListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSystemVulList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSystemVulListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskResultSummaryRequest() (request *DescribeTaskResultSummaryRequest) {
    request = &DescribeTaskResultSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeTaskResultSummary")
    
    
    return
}

func NewDescribeTaskResultSummaryResponse() (response *DescribeTaskResultSummaryResponse) {
    response = &DescribeTaskResultSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTaskResultSummary
// This API is used to query the overview of a check result and return the number of affected nodes in the last seven days.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTaskResultSummary(request *DescribeTaskResultSummaryRequest) (response *DescribeTaskResultSummaryResponse, err error) {
    return c.DescribeTaskResultSummaryWithContext(context.Background(), request)
}

// DescribeTaskResultSummary
// This API is used to query the overview of a check result and return the number of affected nodes in the last seven days.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTaskResultSummaryWithContext(ctx context.Context, request *DescribeTaskResultSummaryRequest) (response *DescribeTaskResultSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeTaskResultSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskResultSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskResultSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTcssSummaryRequest() (request *DescribeTcssSummaryRequest) {
    request = &DescribeTcssSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeTcssSummary")
    
    
    return
}

func NewDescribeTcssSummaryResponse() (response *DescribeTcssSummaryResponse) {
    response = &DescribeTcssSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTcssSummary
// This API is used to query the TCSS overview.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeTcssSummary(request *DescribeTcssSummaryRequest) (response *DescribeTcssSummaryResponse, err error) {
    return c.DescribeTcssSummaryWithContext(context.Background(), request)
}

// DescribeTcssSummary
// This API is used to query the TCSS overview.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeTcssSummaryWithContext(ctx context.Context, request *DescribeTcssSummaryRequest) (response *DescribeTcssSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeTcssSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTcssSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTcssSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUnauthorizedCoresTendencyRequest() (request *DescribeUnauthorizedCoresTendencyRequest) {
    request = &DescribeUnauthorizedCoresTendencyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeUnauthorizedCoresTendency")
    
    
    return
}

func NewDescribeUnauthorizedCoresTendencyResponse() (response *DescribeUnauthorizedCoresTendencyResponse) {
    response = &DescribeUnauthorizedCoresTendencyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUnauthorizedCoresTendency
// This API is used to query the trend of daily unlicensed cores.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeUnauthorizedCoresTendency(request *DescribeUnauthorizedCoresTendencyRequest) (response *DescribeUnauthorizedCoresTendencyResponse, err error) {
    return c.DescribeUnauthorizedCoresTendencyWithContext(context.Background(), request)
}

// DescribeUnauthorizedCoresTendency
// This API is used to query the trend of daily unlicensed cores.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeUnauthorizedCoresTendencyWithContext(ctx context.Context, request *DescribeUnauthorizedCoresTendencyRequest) (response *DescribeUnauthorizedCoresTendencyResponse, err error) {
    if request == nil {
        request = NewDescribeUnauthorizedCoresTendencyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUnauthorizedCoresTendency require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUnauthorizedCoresTendencyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUnfinishRefreshTaskRequest() (request *DescribeUnfinishRefreshTaskRequest) {
    request = &DescribeUnfinishRefreshTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeUnfinishRefreshTask")
    
    
    return
}

func NewDescribeUnfinishRefreshTaskResponse() (response *DescribeUnfinishRefreshTaskResponse) {
    response = &DescribeUnfinishRefreshTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUnfinishRefreshTask
// This API is used to query the information of an unfinished asset refreshing task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeUnfinishRefreshTask(request *DescribeUnfinishRefreshTaskRequest) (response *DescribeUnfinishRefreshTaskResponse, err error) {
    return c.DescribeUnfinishRefreshTaskWithContext(context.Background(), request)
}

// DescribeUnfinishRefreshTask
// This API is used to query the information of an unfinished asset refreshing task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeUnfinishRefreshTaskWithContext(ctx context.Context, request *DescribeUnfinishRefreshTaskRequest) (response *DescribeUnfinishRefreshTaskResponse, err error) {
    if request == nil {
        request = NewDescribeUnfinishRefreshTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUnfinishRefreshTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUnfinishRefreshTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserClusterRequest() (request *DescribeUserClusterRequest) {
    request = &DescribeUserClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeUserCluster")
    
    
    return
}

func NewDescribeUserClusterResponse() (response *DescribeUserClusterResponse) {
    response = &DescribeUserClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserCluster
// This API is used to query the information of a cluster on the Security Dashboard and Cluster Security pages.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeUserCluster(request *DescribeUserClusterRequest) (response *DescribeUserClusterResponse, err error) {
    return c.DescribeUserClusterWithContext(context.Background(), request)
}

// DescribeUserCluster
// This API is used to query the information of a cluster on the Security Dashboard and Cluster Security pages.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeUserClusterWithContext(ctx context.Context, request *DescribeUserClusterRequest) (response *DescribeUserClusterResponse, err error) {
    if request == nil {
        request = NewDescribeUserClusterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserClusterResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeValueAddedSrvInfoRequest() (request *DescribeValueAddedSrvInfoRequest) {
    request = &DescribeValueAddedSrvInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeValueAddedSrvInfo")
    
    
    return
}

func NewDescribeValueAddedSrvInfoResponse() (response *DescribeValueAddedSrvInfoResponse) {
    response = &DescribeValueAddedSrvInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeValueAddedSrvInfo
// This API is used to query the information of the required value-added service.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
func (c *Client) DescribeValueAddedSrvInfo(request *DescribeValueAddedSrvInfoRequest) (response *DescribeValueAddedSrvInfoResponse, err error) {
    return c.DescribeValueAddedSrvInfoWithContext(context.Background(), request)
}

// DescribeValueAddedSrvInfo
// This API is used to query the information of the required value-added service.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
func (c *Client) DescribeValueAddedSrvInfoWithContext(ctx context.Context, request *DescribeValueAddedSrvInfoRequest) (response *DescribeValueAddedSrvInfoResponse, err error) {
    if request == nil {
        request = NewDescribeValueAddedSrvInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeValueAddedSrvInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeValueAddedSrvInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVirusAutoIsolateSampleDetailRequest() (request *DescribeVirusAutoIsolateSampleDetailRequest) {
    request = &DescribeVirusAutoIsolateSampleDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVirusAutoIsolateSampleDetail")
    
    
    return
}

func NewDescribeVirusAutoIsolateSampleDetailResponse() (response *DescribeVirusAutoIsolateSampleDetailResponse) {
    response = &DescribeVirusAutoIsolateSampleDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVirusAutoIsolateSampleDetail
// This API is used to query the details of an automatically isolated trojan sample.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeVirusAutoIsolateSampleDetail(request *DescribeVirusAutoIsolateSampleDetailRequest) (response *DescribeVirusAutoIsolateSampleDetailResponse, err error) {
    return c.DescribeVirusAutoIsolateSampleDetailWithContext(context.Background(), request)
}

// DescribeVirusAutoIsolateSampleDetail
// This API is used to query the details of an automatically isolated trojan sample.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeVirusAutoIsolateSampleDetailWithContext(ctx context.Context, request *DescribeVirusAutoIsolateSampleDetailRequest) (response *DescribeVirusAutoIsolateSampleDetailResponse, err error) {
    if request == nil {
        request = NewDescribeVirusAutoIsolateSampleDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVirusAutoIsolateSampleDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVirusAutoIsolateSampleDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVirusAutoIsolateSampleDownloadURLRequest() (request *DescribeVirusAutoIsolateSampleDownloadURLRequest) {
    request = &DescribeVirusAutoIsolateSampleDownloadURLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVirusAutoIsolateSampleDownloadURL")
    
    
    return
}

func NewDescribeVirusAutoIsolateSampleDownloadURLResponse() (response *DescribeVirusAutoIsolateSampleDownloadURLResponse) {
    response = &DescribeVirusAutoIsolateSampleDownloadURLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVirusAutoIsolateSampleDownloadURL
// This API is used to query the download URL of an automatically isolated trojan sample.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeVirusAutoIsolateSampleDownloadURL(request *DescribeVirusAutoIsolateSampleDownloadURLRequest) (response *DescribeVirusAutoIsolateSampleDownloadURLResponse, err error) {
    return c.DescribeVirusAutoIsolateSampleDownloadURLWithContext(context.Background(), request)
}

// DescribeVirusAutoIsolateSampleDownloadURL
// This API is used to query the download URL of an automatically isolated trojan sample.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeVirusAutoIsolateSampleDownloadURLWithContext(ctx context.Context, request *DescribeVirusAutoIsolateSampleDownloadURLRequest) (response *DescribeVirusAutoIsolateSampleDownloadURLResponse, err error) {
    if request == nil {
        request = NewDescribeVirusAutoIsolateSampleDownloadURLRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVirusAutoIsolateSampleDownloadURL require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVirusAutoIsolateSampleDownloadURLResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVirusAutoIsolateSampleListRequest() (request *DescribeVirusAutoIsolateSampleListRequest) {
    request = &DescribeVirusAutoIsolateSampleListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVirusAutoIsolateSampleList")
    
    
    return
}

func NewDescribeVirusAutoIsolateSampleListResponse() (response *DescribeVirusAutoIsolateSampleListResponse) {
    response = &DescribeVirusAutoIsolateSampleListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVirusAutoIsolateSampleList
// This API is used to query the list of automatically isolated trojan samples.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeVirusAutoIsolateSampleList(request *DescribeVirusAutoIsolateSampleListRequest) (response *DescribeVirusAutoIsolateSampleListResponse, err error) {
    return c.DescribeVirusAutoIsolateSampleListWithContext(context.Background(), request)
}

// DescribeVirusAutoIsolateSampleList
// This API is used to query the list of automatically isolated trojan samples.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeVirusAutoIsolateSampleListWithContext(ctx context.Context, request *DescribeVirusAutoIsolateSampleListRequest) (response *DescribeVirusAutoIsolateSampleListResponse, err error) {
    if request == nil {
        request = NewDescribeVirusAutoIsolateSampleListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVirusAutoIsolateSampleList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVirusAutoIsolateSampleListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVirusAutoIsolateSettingRequest() (request *DescribeVirusAutoIsolateSettingRequest) {
    request = &DescribeVirusAutoIsolateSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVirusAutoIsolateSetting")
    
    
    return
}

func NewDescribeVirusAutoIsolateSettingResponse() (response *DescribeVirusAutoIsolateSettingResponse) {
    response = &DescribeVirusAutoIsolateSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVirusAutoIsolateSetting
// This API is used to query the settings of automatic trojan isolation.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeVirusAutoIsolateSetting(request *DescribeVirusAutoIsolateSettingRequest) (response *DescribeVirusAutoIsolateSettingResponse, err error) {
    return c.DescribeVirusAutoIsolateSettingWithContext(context.Background(), request)
}

// DescribeVirusAutoIsolateSetting
// This API is used to query the settings of automatic trojan isolation.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeVirusAutoIsolateSettingWithContext(ctx context.Context, request *DescribeVirusAutoIsolateSettingRequest) (response *DescribeVirusAutoIsolateSettingResponse, err error) {
    if request == nil {
        request = NewDescribeVirusAutoIsolateSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVirusAutoIsolateSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVirusAutoIsolateSettingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVirusDetailRequest() (request *DescribeVirusDetailRequest) {
    request = &DescribeVirusDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVirusDetail")
    
    
    return
}

func NewDescribeVirusDetailResponse() (response *DescribeVirusDetailResponse) {
    response = &DescribeVirusDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVirusDetail
// This API is used to query the information of a trojan file at runtime.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVirusDetail(request *DescribeVirusDetailRequest) (response *DescribeVirusDetailResponse, err error) {
    return c.DescribeVirusDetailWithContext(context.Background(), request)
}

// DescribeVirusDetail
// This API is used to query the information of a trojan file at runtime.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVirusDetailWithContext(ctx context.Context, request *DescribeVirusDetailRequest) (response *DescribeVirusDetailResponse, err error) {
    if request == nil {
        request = NewDescribeVirusDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVirusDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVirusDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVirusEventTendencyRequest() (request *DescribeVirusEventTendencyRequest) {
    request = &DescribeVirusEventTendencyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVirusEventTendency")
    
    
    return
}

func NewDescribeVirusEventTendencyResponse() (response *DescribeVirusEventTendencyResponse) {
    response = &DescribeVirusEventTendencyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVirusEventTendency
// This API is used to query the trend of trojan events.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeVirusEventTendency(request *DescribeVirusEventTendencyRequest) (response *DescribeVirusEventTendencyResponse, err error) {
    return c.DescribeVirusEventTendencyWithContext(context.Background(), request)
}

// DescribeVirusEventTendency
// This API is used to query the trend of trojan events.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeVirusEventTendencyWithContext(ctx context.Context, request *DescribeVirusEventTendencyRequest) (response *DescribeVirusEventTendencyResponse, err error) {
    if request == nil {
        request = NewDescribeVirusEventTendencyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVirusEventTendency require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVirusEventTendencyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVirusListRequest() (request *DescribeVirusListRequest) {
    request = &DescribeVirusListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVirusList")
    
    
    return
}

func NewDescribeVirusListResponse() (response *DescribeVirusListResponse) {
    response = &DescribeVirusListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVirusList
// This API is used to query the list of virus scanning events at runtime.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVirusList(request *DescribeVirusListRequest) (response *DescribeVirusListResponse, err error) {
    return c.DescribeVirusListWithContext(context.Background(), request)
}

// DescribeVirusList
// This API is used to query the list of virus scanning events at runtime.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVirusListWithContext(ctx context.Context, request *DescribeVirusListRequest) (response *DescribeVirusListResponse, err error) {
    if request == nil {
        request = NewDescribeVirusListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVirusList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVirusListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVirusManualScanEstimateTimeoutRequest() (request *DescribeVirusManualScanEstimateTimeoutRequest) {
    request = &DescribeVirusManualScanEstimateTimeoutRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVirusManualScanEstimateTimeout")
    
    
    return
}

func NewDescribeVirusManualScanEstimateTimeoutResponse() (response *DescribeVirusManualScanEstimateTimeoutResponse) {
    response = &DescribeVirusManualScanEstimateTimeoutResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVirusManualScanEstimateTimeout
// This API is used to query the estimated timeout period of a quick trojan scan.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeVirusManualScanEstimateTimeout(request *DescribeVirusManualScanEstimateTimeoutRequest) (response *DescribeVirusManualScanEstimateTimeoutResponse, err error) {
    return c.DescribeVirusManualScanEstimateTimeoutWithContext(context.Background(), request)
}

// DescribeVirusManualScanEstimateTimeout
// This API is used to query the estimated timeout period of a quick trojan scan.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeVirusManualScanEstimateTimeoutWithContext(ctx context.Context, request *DescribeVirusManualScanEstimateTimeoutRequest) (response *DescribeVirusManualScanEstimateTimeoutResponse, err error) {
    if request == nil {
        request = NewDescribeVirusManualScanEstimateTimeoutRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVirusManualScanEstimateTimeout require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVirusManualScanEstimateTimeoutResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVirusMonitorSettingRequest() (request *DescribeVirusMonitorSettingRequest) {
    request = &DescribeVirusMonitorSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVirusMonitorSetting")
    
    
    return
}

func NewDescribeVirusMonitorSettingResponse() (response *DescribeVirusMonitorSettingResponse) {
    response = &DescribeVirusMonitorSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVirusMonitorSetting
// This API is used to query the real-time monitoring settings of virus scanning at runtime.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVirusMonitorSetting(request *DescribeVirusMonitorSettingRequest) (response *DescribeVirusMonitorSettingResponse, err error) {
    return c.DescribeVirusMonitorSettingWithContext(context.Background(), request)
}

// DescribeVirusMonitorSetting
// This API is used to query the real-time monitoring settings of virus scanning at runtime.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVirusMonitorSettingWithContext(ctx context.Context, request *DescribeVirusMonitorSettingRequest) (response *DescribeVirusMonitorSettingResponse, err error) {
    if request == nil {
        request = NewDescribeVirusMonitorSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVirusMonitorSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVirusMonitorSettingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVirusSampleDownloadUrlRequest() (request *DescribeVirusSampleDownloadUrlRequest) {
    request = &DescribeVirusSampleDownloadUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVirusSampleDownloadUrl")
    
    
    return
}

func NewDescribeVirusSampleDownloadUrlResponse() (response *DescribeVirusSampleDownloadUrlResponse) {
    response = &DescribeVirusSampleDownloadUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVirusSampleDownloadUrl
// This API is used to query the download URL of a trojan sample.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVirusSampleDownloadUrl(request *DescribeVirusSampleDownloadUrlRequest) (response *DescribeVirusSampleDownloadUrlResponse, err error) {
    return c.DescribeVirusSampleDownloadUrlWithContext(context.Background(), request)
}

// DescribeVirusSampleDownloadUrl
// This API is used to query the download URL of a trojan sample.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVirusSampleDownloadUrlWithContext(ctx context.Context, request *DescribeVirusSampleDownloadUrlRequest) (response *DescribeVirusSampleDownloadUrlResponse, err error) {
    if request == nil {
        request = NewDescribeVirusSampleDownloadUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVirusSampleDownloadUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVirusSampleDownloadUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVirusScanSettingRequest() (request *DescribeVirusScanSettingRequest) {
    request = &DescribeVirusScanSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVirusScanSetting")
    
    
    return
}

func NewDescribeVirusScanSettingResponse() (response *DescribeVirusScanSettingResponse) {
    response = &DescribeVirusScanSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVirusScanSetting
// This API is used to query virus scanning settings at runtime.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVirusScanSetting(request *DescribeVirusScanSettingRequest) (response *DescribeVirusScanSettingResponse, err error) {
    return c.DescribeVirusScanSettingWithContext(context.Background(), request)
}

// DescribeVirusScanSetting
// This API is used to query virus scanning settings at runtime.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVirusScanSettingWithContext(ctx context.Context, request *DescribeVirusScanSettingRequest) (response *DescribeVirusScanSettingResponse, err error) {
    if request == nil {
        request = NewDescribeVirusScanSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVirusScanSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVirusScanSettingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVirusScanTaskStatusRequest() (request *DescribeVirusScanTaskStatusRequest) {
    request = &DescribeVirusScanTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVirusScanTaskStatus")
    
    
    return
}

func NewDescribeVirusScanTaskStatusResponse() (response *DescribeVirusScanTaskStatusResponse) {
    response = &DescribeVirusScanTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVirusScanTaskStatus
// This API is used to query the status of a virus scanning task at runtime.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVirusScanTaskStatus(request *DescribeVirusScanTaskStatusRequest) (response *DescribeVirusScanTaskStatusResponse, err error) {
    return c.DescribeVirusScanTaskStatusWithContext(context.Background(), request)
}

// DescribeVirusScanTaskStatus
// This API is used to query the status of a virus scanning task at runtime.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVirusScanTaskStatusWithContext(ctx context.Context, request *DescribeVirusScanTaskStatusRequest) (response *DescribeVirusScanTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeVirusScanTaskStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVirusScanTaskStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVirusScanTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVirusScanTimeoutSettingRequest() (request *DescribeVirusScanTimeoutSettingRequest) {
    request = &DescribeVirusScanTimeoutSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVirusScanTimeoutSetting")
    
    
    return
}

func NewDescribeVirusScanTimeoutSettingResponse() (response *DescribeVirusScanTimeoutSettingResponse) {
    response = &DescribeVirusScanTimeoutSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVirusScanTimeoutSetting
// This API is used to query the timeout settings of a file scan at runtime.
//
// error code that may be returned:
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
func (c *Client) DescribeVirusScanTimeoutSetting(request *DescribeVirusScanTimeoutSettingRequest) (response *DescribeVirusScanTimeoutSettingResponse, err error) {
    return c.DescribeVirusScanTimeoutSettingWithContext(context.Background(), request)
}

// DescribeVirusScanTimeoutSetting
// This API is used to query the timeout settings of a file scan at runtime.
//
// error code that may be returned:
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
func (c *Client) DescribeVirusScanTimeoutSettingWithContext(ctx context.Context, request *DescribeVirusScanTimeoutSettingRequest) (response *DescribeVirusScanTimeoutSettingResponse, err error) {
    if request == nil {
        request = NewDescribeVirusScanTimeoutSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVirusScanTimeoutSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVirusScanTimeoutSettingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVirusSummaryRequest() (request *DescribeVirusSummaryRequest) {
    request = &DescribeVirusSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVirusSummary")
    
    
    return
}

func NewDescribeVirusSummaryResponse() (response *DescribeVirusSummaryResponse) {
    response = &DescribeVirusSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVirusSummary
// This API is used to query the trojan overview at runtime.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVirusSummary(request *DescribeVirusSummaryRequest) (response *DescribeVirusSummaryResponse, err error) {
    return c.DescribeVirusSummaryWithContext(context.Background(), request)
}

// DescribeVirusSummary
// This API is used to query the trojan overview at runtime.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVirusSummaryWithContext(ctx context.Context, request *DescribeVirusSummaryRequest) (response *DescribeVirusSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeVirusSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVirusSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVirusSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVirusTaskListRequest() (request *DescribeVirusTaskListRequest) {
    request = &DescribeVirusTaskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVirusTaskList")
    
    
    return
}

func NewDescribeVirusTaskListResponse() (response *DescribeVirusTaskListResponse) {
    response = &DescribeVirusTaskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVirusTaskList
// This API is used to query the list of virus scanning tasks at runtime.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVirusTaskList(request *DescribeVirusTaskListRequest) (response *DescribeVirusTaskListResponse, err error) {
    return c.DescribeVirusTaskListWithContext(context.Background(), request)
}

// DescribeVirusTaskList
// This API is used to query the list of virus scanning tasks at runtime.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVirusTaskListWithContext(ctx context.Context, request *DescribeVirusTaskListRequest) (response *DescribeVirusTaskListResponse, err error) {
    if request == nil {
        request = NewDescribeVirusTaskListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVirusTaskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVirusTaskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulContainerListRequest() (request *DescribeVulContainerListRequest) {
    request = &DescribeVulContainerListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulContainerList")
    
    
    return
}

func NewDescribeVulContainerListResponse() (response *DescribeVulContainerListResponse) {
    response = &DescribeVulContainerListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVulContainerList
// This API is used to query the list of containers affected by vulnerabilities.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVulContainerList(request *DescribeVulContainerListRequest) (response *DescribeVulContainerListResponse, err error) {
    return c.DescribeVulContainerListWithContext(context.Background(), request)
}

// DescribeVulContainerList
// This API is used to query the list of containers affected by vulnerabilities.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVulContainerListWithContext(ctx context.Context, request *DescribeVulContainerListRequest) (response *DescribeVulContainerListResponse, err error) {
    if request == nil {
        request = NewDescribeVulContainerListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulContainerList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulContainerListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulDefenceEventRequest() (request *DescribeVulDefenceEventRequest) {
    request = &DescribeVulDefenceEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulDefenceEvent")
    
    
    return
}

func NewDescribeVulDefenceEventResponse() (response *DescribeVulDefenceEventResponse) {
    response = &DescribeVulDefenceEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVulDefenceEvent
// This API is used to query the list of exploit prevention events.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeVulDefenceEvent(request *DescribeVulDefenceEventRequest) (response *DescribeVulDefenceEventResponse, err error) {
    return c.DescribeVulDefenceEventWithContext(context.Background(), request)
}

// DescribeVulDefenceEvent
// This API is used to query the list of exploit prevention events.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeVulDefenceEventWithContext(ctx context.Context, request *DescribeVulDefenceEventRequest) (response *DescribeVulDefenceEventResponse, err error) {
    if request == nil {
        request = NewDescribeVulDefenceEventRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulDefenceEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulDefenceEventResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulDefenceEventDetailRequest() (request *DescribeVulDefenceEventDetailRequest) {
    request = &DescribeVulDefenceEventDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulDefenceEventDetail")
    
    
    return
}

func NewDescribeVulDefenceEventDetailResponse() (response *DescribeVulDefenceEventDetailResponse) {
    response = &DescribeVulDefenceEventDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVulDefenceEventDetail
// This API is used to query the details of an exploit prevention event.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVulDefenceEventDetail(request *DescribeVulDefenceEventDetailRequest) (response *DescribeVulDefenceEventDetailResponse, err error) {
    return c.DescribeVulDefenceEventDetailWithContext(context.Background(), request)
}

// DescribeVulDefenceEventDetail
// This API is used to query the details of an exploit prevention event.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVulDefenceEventDetailWithContext(ctx context.Context, request *DescribeVulDefenceEventDetailRequest) (response *DescribeVulDefenceEventDetailResponse, err error) {
    if request == nil {
        request = NewDescribeVulDefenceEventDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulDefenceEventDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulDefenceEventDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulDefenceEventTendencyRequest() (request *DescribeVulDefenceEventTendencyRequest) {
    request = &DescribeVulDefenceEventTendencyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulDefenceEventTendency")
    
    
    return
}

func NewDescribeVulDefenceEventTendencyResponse() (response *DescribeVulDefenceEventTendencyResponse) {
    response = &DescribeVulDefenceEventTendencyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVulDefenceEventTendency
// This API is used to query the trend of exploit prevention events.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeVulDefenceEventTendency(request *DescribeVulDefenceEventTendencyRequest) (response *DescribeVulDefenceEventTendencyResponse, err error) {
    return c.DescribeVulDefenceEventTendencyWithContext(context.Background(), request)
}

// DescribeVulDefenceEventTendency
// This API is used to query the trend of exploit prevention events.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeVulDefenceEventTendencyWithContext(ctx context.Context, request *DescribeVulDefenceEventTendencyRequest) (response *DescribeVulDefenceEventTendencyResponse, err error) {
    if request == nil {
        request = NewDescribeVulDefenceEventTendencyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulDefenceEventTendency require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulDefenceEventTendencyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulDefenceHostRequest() (request *DescribeVulDefenceHostRequest) {
    request = &DescribeVulDefenceHostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulDefenceHost")
    
    
    return
}

func NewDescribeVulDefenceHostResponse() (response *DescribeVulDefenceHostResponse) {
    response = &DescribeVulDefenceHostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVulDefenceHost
// This API is used to query the list of servers with exploit prevention enabled.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeVulDefenceHost(request *DescribeVulDefenceHostRequest) (response *DescribeVulDefenceHostResponse, err error) {
    return c.DescribeVulDefenceHostWithContext(context.Background(), request)
}

// DescribeVulDefenceHost
// This API is used to query the list of servers with exploit prevention enabled.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeVulDefenceHostWithContext(ctx context.Context, request *DescribeVulDefenceHostRequest) (response *DescribeVulDefenceHostResponse, err error) {
    if request == nil {
        request = NewDescribeVulDefenceHostRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulDefenceHost require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulDefenceHostResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulDefencePluginRequest() (request *DescribeVulDefencePluginRequest) {
    request = &DescribeVulDefencePluginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulDefencePlugin")
    
    
    return
}

func NewDescribeVulDefencePluginResponse() (response *DescribeVulDefencePluginResponse) {
    response = &DescribeVulDefencePluginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVulDefencePlugin
// This API is used to query the list of exploit prevention plugins.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeVulDefencePlugin(request *DescribeVulDefencePluginRequest) (response *DescribeVulDefencePluginResponse, err error) {
    return c.DescribeVulDefencePluginWithContext(context.Background(), request)
}

// DescribeVulDefencePlugin
// This API is used to query the list of exploit prevention plugins.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeVulDefencePluginWithContext(ctx context.Context, request *DescribeVulDefencePluginRequest) (response *DescribeVulDefencePluginResponse, err error) {
    if request == nil {
        request = NewDescribeVulDefencePluginRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulDefencePlugin require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulDefencePluginResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulDefenceSettingRequest() (request *DescribeVulDefenceSettingRequest) {
    request = &DescribeVulDefenceSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulDefenceSetting")
    
    
    return
}

func NewDescribeVulDefenceSettingResponse() (response *DescribeVulDefenceSettingResponse) {
    response = &DescribeVulDefenceSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVulDefenceSetting
// This API is used to query the exploit prevention settings.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeVulDefenceSetting(request *DescribeVulDefenceSettingRequest) (response *DescribeVulDefenceSettingResponse, err error) {
    return c.DescribeVulDefenceSettingWithContext(context.Background(), request)
}

// DescribeVulDefenceSetting
// This API is used to query the exploit prevention settings.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeVulDefenceSettingWithContext(ctx context.Context, request *DescribeVulDefenceSettingRequest) (response *DescribeVulDefenceSettingResponse, err error) {
    if request == nil {
        request = NewDescribeVulDefenceSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulDefenceSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulDefenceSettingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulDetailRequest() (request *DescribeVulDetailRequest) {
    request = &DescribeVulDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulDetail")
    
    
    return
}

func NewDescribeVulDetailResponse() (response *DescribeVulDetailResponse) {
    response = &DescribeVulDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVulDetail
// This API is used to query vulnerability details.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVulDetail(request *DescribeVulDetailRequest) (response *DescribeVulDetailResponse, err error) {
    return c.DescribeVulDetailWithContext(context.Background(), request)
}

// DescribeVulDetail
// This API is used to query vulnerability details.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVulDetailWithContext(ctx context.Context, request *DescribeVulDetailRequest) (response *DescribeVulDetailResponse, err error) {
    if request == nil {
        request = NewDescribeVulDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulIgnoreLocalImageListRequest() (request *DescribeVulIgnoreLocalImageListRequest) {
    request = &DescribeVulIgnoreLocalImageListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulIgnoreLocalImageList")
    
    
    return
}

func NewDescribeVulIgnoreLocalImageListResponse() (response *DescribeVulIgnoreLocalImageListResponse) {
    response = &DescribeVulIgnoreLocalImageListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVulIgnoreLocalImageList
// This API is used to query the list of local images ignored in a vulnerability scan.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVulIgnoreLocalImageList(request *DescribeVulIgnoreLocalImageListRequest) (response *DescribeVulIgnoreLocalImageListResponse, err error) {
    return c.DescribeVulIgnoreLocalImageListWithContext(context.Background(), request)
}

// DescribeVulIgnoreLocalImageList
// This API is used to query the list of local images ignored in a vulnerability scan.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVulIgnoreLocalImageListWithContext(ctx context.Context, request *DescribeVulIgnoreLocalImageListRequest) (response *DescribeVulIgnoreLocalImageListResponse, err error) {
    if request == nil {
        request = NewDescribeVulIgnoreLocalImageListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulIgnoreLocalImageList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulIgnoreLocalImageListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulIgnoreRegistryImageListRequest() (request *DescribeVulIgnoreRegistryImageListRequest) {
    request = &DescribeVulIgnoreRegistryImageListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulIgnoreRegistryImageList")
    
    
    return
}

func NewDescribeVulIgnoreRegistryImageListResponse() (response *DescribeVulIgnoreRegistryImageListResponse) {
    response = &DescribeVulIgnoreRegistryImageListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVulIgnoreRegistryImageList
// This API is used to query the list of repository images ignored in a vulnerability scan.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVulIgnoreRegistryImageList(request *DescribeVulIgnoreRegistryImageListRequest) (response *DescribeVulIgnoreRegistryImageListResponse, err error) {
    return c.DescribeVulIgnoreRegistryImageListWithContext(context.Background(), request)
}

// DescribeVulIgnoreRegistryImageList
// This API is used to query the list of repository images ignored in a vulnerability scan.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVulIgnoreRegistryImageListWithContext(ctx context.Context, request *DescribeVulIgnoreRegistryImageListRequest) (response *DescribeVulIgnoreRegistryImageListResponse, err error) {
    if request == nil {
        request = NewDescribeVulIgnoreRegistryImageListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulIgnoreRegistryImageList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulIgnoreRegistryImageListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulImageListRequest() (request *DescribeVulImageListRequest) {
    request = &DescribeVulImageListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulImageList")
    
    
    return
}

func NewDescribeVulImageListResponse() (response *DescribeVulImageListResponse) {
    response = &DescribeVulImageListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVulImageList
// This API is used to query the list of images affected by vulnerabilities.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVulImageList(request *DescribeVulImageListRequest) (response *DescribeVulImageListResponse, err error) {
    return c.DescribeVulImageListWithContext(context.Background(), request)
}

// DescribeVulImageList
// This API is used to query the list of images affected by vulnerabilities.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVulImageListWithContext(ctx context.Context, request *DescribeVulImageListRequest) (response *DescribeVulImageListResponse, err error) {
    if request == nil {
        request = NewDescribeVulImageListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulImageList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulImageListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulImageSummaryRequest() (request *DescribeVulImageSummaryRequest) {
    request = &DescribeVulImageSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulImageSummary")
    
    
    return
}

func NewDescribeVulImageSummaryResponse() (response *DescribeVulImageSummaryResponse) {
    response = &DescribeVulImageSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVulImageSummary
// This API is used to query the statistics of images affected by vulnerabilities.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVulImageSummary(request *DescribeVulImageSummaryRequest) (response *DescribeVulImageSummaryResponse, err error) {
    return c.DescribeVulImageSummaryWithContext(context.Background(), request)
}

// DescribeVulImageSummary
// This API is used to query the statistics of images affected by vulnerabilities.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVulImageSummaryWithContext(ctx context.Context, request *DescribeVulImageSummaryRequest) (response *DescribeVulImageSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeVulImageSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulImageSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulImageSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulLevelImageSummaryRequest() (request *DescribeVulLevelImageSummaryRequest) {
    request = &DescribeVulLevelImageSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulLevelImageSummary")
    
    
    return
}

func NewDescribeVulLevelImageSummaryResponse() (response *DescribeVulLevelImageSummaryResponse) {
    response = &DescribeVulLevelImageSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVulLevelImageSummary
// This API is used to query the numbers of images affected by emergency vulnerabilities at each severity level.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVulLevelImageSummary(request *DescribeVulLevelImageSummaryRequest) (response *DescribeVulLevelImageSummaryResponse, err error) {
    return c.DescribeVulLevelImageSummaryWithContext(context.Background(), request)
}

// DescribeVulLevelImageSummary
// This API is used to query the numbers of images affected by emergency vulnerabilities at each severity level.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVulLevelImageSummaryWithContext(ctx context.Context, request *DescribeVulLevelImageSummaryRequest) (response *DescribeVulLevelImageSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeVulLevelImageSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulLevelImageSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulLevelImageSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulLevelSummaryRequest() (request *DescribeVulLevelSummaryRequest) {
    request = &DescribeVulLevelSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulLevelSummary")
    
    
    return
}

func NewDescribeVulLevelSummaryResponse() (response *DescribeVulLevelSummaryResponse) {
    response = &DescribeVulLevelSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVulLevelSummary
// This API is used to query the numbers of vulnerabilities at each severity level.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVulLevelSummary(request *DescribeVulLevelSummaryRequest) (response *DescribeVulLevelSummaryResponse, err error) {
    return c.DescribeVulLevelSummaryWithContext(context.Background(), request)
}

// DescribeVulLevelSummary
// This API is used to query the numbers of vulnerabilities at each severity level.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVulLevelSummaryWithContext(ctx context.Context, request *DescribeVulLevelSummaryRequest) (response *DescribeVulLevelSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeVulLevelSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulLevelSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulLevelSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulRegistryImageListRequest() (request *DescribeVulRegistryImageListRequest) {
    request = &DescribeVulRegistryImageListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulRegistryImageList")
    
    
    return
}

func NewDescribeVulRegistryImageListResponse() (response *DescribeVulRegistryImageListResponse) {
    response = &DescribeVulRegistryImageListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVulRegistryImageList
// This API is used to query the list of repository images affected by vulnerabilities.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVulRegistryImageList(request *DescribeVulRegistryImageListRequest) (response *DescribeVulRegistryImageListResponse, err error) {
    return c.DescribeVulRegistryImageListWithContext(context.Background(), request)
}

// DescribeVulRegistryImageList
// This API is used to query the list of repository images affected by vulnerabilities.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVulRegistryImageListWithContext(ctx context.Context, request *DescribeVulRegistryImageListRequest) (response *DescribeVulRegistryImageListResponse, err error) {
    if request == nil {
        request = NewDescribeVulRegistryImageListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulRegistryImageList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulRegistryImageListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulScanAuthorizedImageSummaryRequest() (request *DescribeVulScanAuthorizedImageSummaryRequest) {
    request = &DescribeVulScanAuthorizedImageSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulScanAuthorizedImageSummary")
    
    
    return
}

func NewDescribeVulScanAuthorizedImageSummaryResponse() (response *DescribeVulScanAuthorizedImageSummaryResponse) {
    response = &DescribeVulScanAuthorizedImageSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVulScanAuthorizedImageSummary
// This API is used to count the number of licensed but not scanned images on the vulnerability scanning page.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVulScanAuthorizedImageSummary(request *DescribeVulScanAuthorizedImageSummaryRequest) (response *DescribeVulScanAuthorizedImageSummaryResponse, err error) {
    return c.DescribeVulScanAuthorizedImageSummaryWithContext(context.Background(), request)
}

// DescribeVulScanAuthorizedImageSummary
// This API is used to count the number of licensed but not scanned images on the vulnerability scanning page.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVulScanAuthorizedImageSummaryWithContext(ctx context.Context, request *DescribeVulScanAuthorizedImageSummaryRequest) (response *DescribeVulScanAuthorizedImageSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeVulScanAuthorizedImageSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulScanAuthorizedImageSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulScanAuthorizedImageSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulScanInfoRequest() (request *DescribeVulScanInfoRequest) {
    request = &DescribeVulScanInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulScanInfo")
    
    
    return
}

func NewDescribeVulScanInfoResponse() (response *DescribeVulScanInfoResponse) {
    response = &DescribeVulScanInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVulScanInfo
// This API is used to query the information of a vulnerability scan task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVulScanInfo(request *DescribeVulScanInfoRequest) (response *DescribeVulScanInfoResponse, err error) {
    return c.DescribeVulScanInfoWithContext(context.Background(), request)
}

// DescribeVulScanInfo
// This API is used to query the information of a vulnerability scan task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVulScanInfoWithContext(ctx context.Context, request *DescribeVulScanInfoRequest) (response *DescribeVulScanInfoResponse, err error) {
    if request == nil {
        request = NewDescribeVulScanInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulScanInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulScanInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulScanLocalImageListRequest() (request *DescribeVulScanLocalImageListRequest) {
    request = &DescribeVulScanLocalImageListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulScanLocalImageList")
    
    
    return
}

func NewDescribeVulScanLocalImageListResponse() (response *DescribeVulScanLocalImageListResponse) {
    response = &DescribeVulScanLocalImageListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVulScanLocalImageList
// This API is used to query the list of local images in a vulnerability scan task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVulScanLocalImageList(request *DescribeVulScanLocalImageListRequest) (response *DescribeVulScanLocalImageListResponse, err error) {
    return c.DescribeVulScanLocalImageListWithContext(context.Background(), request)
}

// DescribeVulScanLocalImageList
// This API is used to query the list of local images in a vulnerability scan task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVulScanLocalImageListWithContext(ctx context.Context, request *DescribeVulScanLocalImageListRequest) (response *DescribeVulScanLocalImageListResponse, err error) {
    if request == nil {
        request = NewDescribeVulScanLocalImageListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulScanLocalImageList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulScanLocalImageListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulSummaryRequest() (request *DescribeVulSummaryRequest) {
    request = &DescribeVulSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulSummary")
    
    
    return
}

func NewDescribeVulSummaryResponse() (response *DescribeVulSummaryResponse) {
    response = &DescribeVulSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVulSummary
// This API is used to query the overview of vulnerability risks.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVulSummary(request *DescribeVulSummaryRequest) (response *DescribeVulSummaryResponse, err error) {
    return c.DescribeVulSummaryWithContext(context.Background(), request)
}

// DescribeVulSummary
// This API is used to query the overview of vulnerability risks.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVulSummaryWithContext(ctx context.Context, request *DescribeVulSummaryRequest) (response *DescribeVulSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeVulSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulTendencyRequest() (request *DescribeVulTendencyRequest) {
    request = &DescribeVulTendencyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulTendency")
    
    
    return
}

func NewDescribeVulTendencyResponse() (response *DescribeVulTendencyResponse) {
    response = &DescribeVulTendencyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVulTendency
// This API is used to query the trend of critical and high-risk vulnerabilities in local and repository images.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVulTendency(request *DescribeVulTendencyRequest) (response *DescribeVulTendencyResponse, err error) {
    return c.DescribeVulTendencyWithContext(context.Background(), request)
}

// DescribeVulTendency
// This API is used to query the trend of critical and high-risk vulnerabilities in local and repository images.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVulTendencyWithContext(ctx context.Context, request *DescribeVulTendencyRequest) (response *DescribeVulTendencyResponse, err error) {
    if request == nil {
        request = NewDescribeVulTendencyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulTendency require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulTendencyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulTopRankingRequest() (request *DescribeVulTopRankingRequest) {
    request = &DescribeVulTopRankingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVulTopRanking")
    
    
    return
}

func NewDescribeVulTopRankingResponse() (response *DescribeVulTopRankingResponse) {
    response = &DescribeVulTopRankingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVulTopRanking
// This API is used to query the list of top vulnerabilities.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DescribeVulTopRanking(request *DescribeVulTopRankingRequest) (response *DescribeVulTopRankingResponse, err error) {
    return c.DescribeVulTopRankingWithContext(context.Background(), request)
}

// DescribeVulTopRanking
// This API is used to query the list of top vulnerabilities.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DescribeVulTopRankingWithContext(ctx context.Context, request *DescribeVulTopRankingRequest) (response *DescribeVulTopRankingResponse, err error) {
    if request == nil {
        request = NewDescribeVulTopRankingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulTopRanking require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulTopRankingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWarningRulesRequest() (request *DescribeWarningRulesRequest) {
    request = &DescribeWarningRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeWarningRules")
    
    
    return
}

func NewDescribeWarningRulesResponse() (response *DescribeWarningRulesResponse) {
    response = &DescribeWarningRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWarningRules
// This API is used to get the list of alert policies.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeWarningRules(request *DescribeWarningRulesRequest) (response *DescribeWarningRulesResponse, err error) {
    return c.DescribeWarningRulesWithContext(context.Background(), request)
}

// DescribeWarningRules
// This API is used to get the list of alert policies.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeWarningRulesWithContext(ctx context.Context, request *DescribeWarningRulesRequest) (response *DescribeWarningRulesResponse, err error) {
    if request == nil {
        request = NewDescribeWarningRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWarningRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWarningRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWebVulListRequest() (request *DescribeWebVulListRequest) {
    request = &DescribeWebVulListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeWebVulList")
    
    
    return
}

func NewDescribeWebVulListResponse() (response *DescribeWebVulListResponse) {
    response = &DescribeWebVulListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWebVulList
// This API is used to query the list of web application vulnerabilities.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeWebVulList(request *DescribeWebVulListRequest) (response *DescribeWebVulListResponse, err error) {
    return c.DescribeWebVulListWithContext(context.Background(), request)
}

// DescribeWebVulList
// This API is used to query the list of web application vulnerabilities.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeWebVulListWithContext(ctx context.Context, request *DescribeWebVulListRequest) (response *DescribeWebVulListResponse, err error) {
    if request == nil {
        request = NewDescribeWebVulListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWebVulList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWebVulListResponse()
    err = c.Send(request, response)
    return
}

func NewExportVirusListRequest() (request *ExportVirusListRequest) {
    request = &ExportVirusListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "ExportVirusList")
    
    
    return
}

func NewExportVirusListResponse() (response *ExportVirusListResponse) {
    response = &ExportVirusListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ExportVirusList
// This API is used to export the list of virus scanning events at runtime.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ExportVirusList(request *ExportVirusListRequest) (response *ExportVirusListResponse, err error) {
    return c.ExportVirusListWithContext(context.Background(), request)
}

// ExportVirusList
// This API is used to export the list of virus scanning events at runtime.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ExportVirusListWithContext(ctx context.Context, request *ExportVirusListRequest) (response *ExportVirusListResponse, err error) {
    if request == nil {
        request = NewExportVirusListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportVirusList require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportVirusListResponse()
    err = c.Send(request, response)
    return
}

func NewInitializeUserComplianceEnvironmentRequest() (request *InitializeUserComplianceEnvironmentRequest) {
    request = &InitializeUserComplianceEnvironmentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "InitializeUserComplianceEnvironment")
    
    
    return
}

func NewInitializeUserComplianceEnvironmentResponse() (response *InitializeUserComplianceEnvironmentResponse) {
    response = &InitializeUserComplianceEnvironmentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InitializeUserComplianceEnvironment
// This API is used to initialize the compliance baseline environment and create necessary data and options.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) InitializeUserComplianceEnvironment(request *InitializeUserComplianceEnvironmentRequest) (response *InitializeUserComplianceEnvironmentResponse, err error) {
    return c.InitializeUserComplianceEnvironmentWithContext(context.Background(), request)
}

// InitializeUserComplianceEnvironment
// This API is used to initialize the compliance baseline environment and create necessary data and options.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) InitializeUserComplianceEnvironmentWithContext(ctx context.Context, request *InitializeUserComplianceEnvironmentRequest) (response *InitializeUserComplianceEnvironmentResponse, err error) {
    if request == nil {
        request = NewInitializeUserComplianceEnvironmentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InitializeUserComplianceEnvironment require credential")
    }

    request.SetContext(ctx)
    
    response = NewInitializeUserComplianceEnvironmentResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAbnormalProcessRuleStatusRequest() (request *ModifyAbnormalProcessRuleStatusRequest) {
    request = &ModifyAbnormalProcessRuleStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyAbnormalProcessRuleStatus")
    
    
    return
}

func NewModifyAbnormalProcessRuleStatusResponse() (response *ModifyAbnormalProcessRuleStatusResponse) {
    response = &ModifyAbnormalProcessRuleStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAbnormalProcessRuleStatus
// This API is used to change the status of an abnormal process policy at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAbnormalProcessRuleStatus(request *ModifyAbnormalProcessRuleStatusRequest) (response *ModifyAbnormalProcessRuleStatusResponse, err error) {
    return c.ModifyAbnormalProcessRuleStatusWithContext(context.Background(), request)
}

// ModifyAbnormalProcessRuleStatus
// This API is used to change the status of an abnormal process policy at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAbnormalProcessRuleStatusWithContext(ctx context.Context, request *ModifyAbnormalProcessRuleStatusRequest) (response *ModifyAbnormalProcessRuleStatusResponse, err error) {
    if request == nil {
        request = NewModifyAbnormalProcessRuleStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAbnormalProcessRuleStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAbnormalProcessRuleStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAbnormalProcessStatusRequest() (request *ModifyAbnormalProcessStatusRequest) {
    request = &ModifyAbnormalProcessStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyAbnormalProcessStatus")
    
    
    return
}

func NewModifyAbnormalProcessStatusResponse() (response *ModifyAbnormalProcessStatusResponse) {
    response = &ModifyAbnormalProcessStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAbnormalProcessStatus
// This API is used to change the status of an abnormal process event.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAbnormalProcessStatus(request *ModifyAbnormalProcessStatusRequest) (response *ModifyAbnormalProcessStatusResponse, err error) {
    return c.ModifyAbnormalProcessStatusWithContext(context.Background(), request)
}

// ModifyAbnormalProcessStatus
// This API is used to change the status of an abnormal process event.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAbnormalProcessStatusWithContext(ctx context.Context, request *ModifyAbnormalProcessStatusRequest) (response *ModifyAbnormalProcessStatusResponse, err error) {
    if request == nil {
        request = NewModifyAbnormalProcessStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAbnormalProcessStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAbnormalProcessStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccessControlRuleStatusRequest() (request *ModifyAccessControlRuleStatusRequest) {
    request = &ModifyAccessControlRuleStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyAccessControlRuleStatus")
    
    
    return
}

func NewModifyAccessControlRuleStatusResponse() (response *ModifyAccessControlRuleStatusResponse) {
    response = &ModifyAccessControlRuleStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAccessControlRuleStatus
// This API is used to change the status of an access control policy at runtime, i.e., enable or disable it.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAccessControlRuleStatus(request *ModifyAccessControlRuleStatusRequest) (response *ModifyAccessControlRuleStatusResponse, err error) {
    return c.ModifyAccessControlRuleStatusWithContext(context.Background(), request)
}

// ModifyAccessControlRuleStatus
// This API is used to change the status of an access control policy at runtime, i.e., enable or disable it.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAccessControlRuleStatusWithContext(ctx context.Context, request *ModifyAccessControlRuleStatusRequest) (response *ModifyAccessControlRuleStatusResponse, err error) {
    if request == nil {
        request = NewModifyAccessControlRuleStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccessControlRuleStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAccessControlRuleStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccessControlStatusRequest() (request *ModifyAccessControlStatusRequest) {
    request = &ModifyAccessControlStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyAccessControlStatus")
    
    
    return
}

func NewModifyAccessControlStatusResponse() (response *ModifyAccessControlStatusResponse) {
    response = &ModifyAccessControlStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAccessControlStatus
// This API is used to change the status of an access control event at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAccessControlStatus(request *ModifyAccessControlStatusRequest) (response *ModifyAccessControlStatusResponse, err error) {
    return c.ModifyAccessControlStatusWithContext(context.Background(), request)
}

// ModifyAccessControlStatus
// This API is used to change the status of an access control event at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAccessControlStatusWithContext(ctx context.Context, request *ModifyAccessControlStatusRequest) (response *ModifyAccessControlStatusResponse, err error) {
    if request == nil {
        request = NewModifyAccessControlStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccessControlStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAccessControlStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAssetRequest() (request *ModifyAssetRequest) {
    request = &ModifyAssetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyAsset")
    
    
    return
}

func NewModifyAssetResponse() (response *ModifyAssetResponse) {
    response = &ModifyAssetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAsset
// This API is used to refresh server assets.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAsset(request *ModifyAssetRequest) (response *ModifyAssetResponse, err error) {
    return c.ModifyAssetWithContext(context.Background(), request)
}

// ModifyAsset
// This API is used to refresh server assets.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAssetWithContext(ctx context.Context, request *ModifyAssetRequest) (response *ModifyAssetResponse, err error) {
    if request == nil {
        request = NewModifyAssetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAsset require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAssetResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAssetImageRegistryScanStopRequest() (request *ModifyAssetImageRegistryScanStopRequest) {
    request = &ModifyAssetImageRegistryScanStopRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyAssetImageRegistryScanStop")
    
    
    return
}

func NewModifyAssetImageRegistryScanStopResponse() (response *ModifyAssetImageRegistryScanStopResponse) {
    response = &ModifyAssetImageRegistryScanStopResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAssetImageRegistryScanStop
// This API is used to stop an image scan task for an image repository.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAssetImageRegistryScanStop(request *ModifyAssetImageRegistryScanStopRequest) (response *ModifyAssetImageRegistryScanStopResponse, err error) {
    return c.ModifyAssetImageRegistryScanStopWithContext(context.Background(), request)
}

// ModifyAssetImageRegistryScanStop
// This API is used to stop an image scan task for an image repository.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAssetImageRegistryScanStopWithContext(ctx context.Context, request *ModifyAssetImageRegistryScanStopRequest) (response *ModifyAssetImageRegistryScanStopResponse, err error) {
    if request == nil {
        request = NewModifyAssetImageRegistryScanStopRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAssetImageRegistryScanStop require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAssetImageRegistryScanStopResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAssetImageRegistryScanStopOneKeyRequest() (request *ModifyAssetImageRegistryScanStopOneKeyRequest) {
    request = &ModifyAssetImageRegistryScanStopOneKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyAssetImageRegistryScanStopOneKey")
    
    
    return
}

func NewModifyAssetImageRegistryScanStopOneKeyResponse() (response *ModifyAssetImageRegistryScanStopOneKeyResponse) {
    response = &ModifyAssetImageRegistryScanStopOneKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAssetImageRegistryScanStopOneKey
// This API is used to stop a quick image scan task for an image repository.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAssetImageRegistryScanStopOneKey(request *ModifyAssetImageRegistryScanStopOneKeyRequest) (response *ModifyAssetImageRegistryScanStopOneKeyResponse, err error) {
    return c.ModifyAssetImageRegistryScanStopOneKeyWithContext(context.Background(), request)
}

// ModifyAssetImageRegistryScanStopOneKey
// This API is used to stop a quick image scan task for an image repository.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAssetImageRegistryScanStopOneKeyWithContext(ctx context.Context, request *ModifyAssetImageRegistryScanStopOneKeyRequest) (response *ModifyAssetImageRegistryScanStopOneKeyResponse, err error) {
    if request == nil {
        request = NewModifyAssetImageRegistryScanStopOneKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAssetImageRegistryScanStopOneKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAssetImageRegistryScanStopOneKeyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAssetImageScanStopRequest() (request *ModifyAssetImageScanStopRequest) {
    request = &ModifyAssetImageScanStopRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyAssetImageScanStop")
    
    
    return
}

func NewModifyAssetImageScanStopResponse() (response *ModifyAssetImageScanStopResponse) {
    response = &ModifyAssetImageScanStopResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAssetImageScanStop
// This API is used to stop an image scan.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAssetImageScanStop(request *ModifyAssetImageScanStopRequest) (response *ModifyAssetImageScanStopResponse, err error) {
    return c.ModifyAssetImageScanStopWithContext(context.Background(), request)
}

// ModifyAssetImageScanStop
// This API is used to stop an image scan.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAssetImageScanStopWithContext(ctx context.Context, request *ModifyAssetImageScanStopRequest) (response *ModifyAssetImageScanStopResponse, err error) {
    if request == nil {
        request = NewModifyAssetImageScanStopRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAssetImageScanStop require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAssetImageScanStopResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCompliancePeriodTaskRequest() (request *ModifyCompliancePeriodTaskRequest) {
    request = &ModifyCompliancePeriodTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyCompliancePeriodTask")
    
    
    return
}

func NewModifyCompliancePeriodTaskResponse() (response *ModifyCompliancePeriodTaskResponse) {
    response = &ModifyCompliancePeriodTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCompliancePeriodTask
// This API is used to modify the settings of a scheduled task, including the check cycle and the status of the compliance benchmark.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ModifyCompliancePeriodTask(request *ModifyCompliancePeriodTaskRequest) (response *ModifyCompliancePeriodTaskResponse, err error) {
    return c.ModifyCompliancePeriodTaskWithContext(context.Background(), request)
}

// ModifyCompliancePeriodTask
// This API is used to modify the settings of a scheduled task, including the check cycle and the status of the compliance benchmark.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ModifyCompliancePeriodTaskWithContext(ctx context.Context, request *ModifyCompliancePeriodTaskRequest) (response *ModifyCompliancePeriodTaskResponse, err error) {
    if request == nil {
        request = NewModifyCompliancePeriodTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCompliancePeriodTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCompliancePeriodTaskResponse()
    err = c.Send(request, response)
    return
}

func NewModifyContainerNetStatusRequest() (request *ModifyContainerNetStatusRequest) {
    request = &ModifyContainerNetStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyContainerNetStatus")
    
    
    return
}

func NewModifyContainerNetStatusResponse() (response *ModifyContainerNetStatusResponse) {
    response = &ModifyContainerNetStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyContainerNetStatus
// This API is used to isolate a container.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ModifyContainerNetStatus(request *ModifyContainerNetStatusRequest) (response *ModifyContainerNetStatusResponse, err error) {
    return c.ModifyContainerNetStatusWithContext(context.Background(), request)
}

// ModifyContainerNetStatus
// This API is used to isolate a container.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ModifyContainerNetStatusWithContext(ctx context.Context, request *ModifyContainerNetStatusRequest) (response *ModifyContainerNetStatusResponse, err error) {
    if request == nil {
        request = NewModifyContainerNetStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyContainerNetStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyContainerNetStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEscapeEventStatusRequest() (request *ModifyEscapeEventStatusRequest) {
    request = &ModifyEscapeEventStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyEscapeEventStatus")
    
    
    return
}

func NewModifyEscapeEventStatusResponse() (response *ModifyEscapeEventStatusResponse) {
    response = &ModifyEscapeEventStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyEscapeEventStatus
// This API is used to change the status of a container escape scan event.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyEscapeEventStatus(request *ModifyEscapeEventStatusRequest) (response *ModifyEscapeEventStatusResponse, err error) {
    return c.ModifyEscapeEventStatusWithContext(context.Background(), request)
}

// ModifyEscapeEventStatus
// This API is used to change the status of a container escape scan event.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyEscapeEventStatusWithContext(ctx context.Context, request *ModifyEscapeEventStatusRequest) (response *ModifyEscapeEventStatusResponse, err error) {
    if request == nil {
        request = NewModifyEscapeEventStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEscapeEventStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyEscapeEventStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEscapeRuleRequest() (request *ModifyEscapeRuleRequest) {
    request = &ModifyEscapeRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyEscapeRule")
    
    
    return
}

func NewModifyEscapeRuleResponse() (response *ModifyEscapeRuleResponse) {
    response = &ModifyEscapeRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyEscapeRule
// This API is used to modify the information of a container escape scan rule.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_NOTIFYPOLICYCHANGEFAILED = "FailedOperation.NotifyPolicyChangeFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) ModifyEscapeRule(request *ModifyEscapeRuleRequest) (response *ModifyEscapeRuleResponse, err error) {
    return c.ModifyEscapeRuleWithContext(context.Background(), request)
}

// ModifyEscapeRule
// This API is used to modify the information of a container escape scan rule.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_NOTIFYPOLICYCHANGEFAILED = "FailedOperation.NotifyPolicyChangeFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) ModifyEscapeRuleWithContext(ctx context.Context, request *ModifyEscapeRuleRequest) (response *ModifyEscapeRuleResponse, err error) {
    if request == nil {
        request = NewModifyEscapeRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEscapeRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyEscapeRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEscapeWhiteListRequest() (request *ModifyEscapeWhiteListRequest) {
    request = &ModifyEscapeWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyEscapeWhiteList")
    
    
    return
}

func NewModifyEscapeWhiteListResponse() (response *ModifyEscapeWhiteListResponse) {
    response = &ModifyEscapeWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyEscapeWhiteList
// This API is used to modify an allowed escape.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyEscapeWhiteList(request *ModifyEscapeWhiteListRequest) (response *ModifyEscapeWhiteListResponse, err error) {
    return c.ModifyEscapeWhiteListWithContext(context.Background(), request)
}

// ModifyEscapeWhiteList
// This API is used to modify an allowed escape.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyEscapeWhiteListWithContext(ctx context.Context, request *ModifyEscapeWhiteListRequest) (response *ModifyEscapeWhiteListResponse, err error) {
    if request == nil {
        request = NewModifyEscapeWhiteListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEscapeWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyEscapeWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewModifyImageAuthorizedRequest() (request *ModifyImageAuthorizedRequest) {
    request = &ModifyImageAuthorizedRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyImageAuthorized")
    
    
    return
}

func NewModifyImageAuthorizedResponse() (response *ModifyImageAuthorizedResponse) {
    response = &ModifyImageAuthorizedResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyImageAuthorized
// This API is used to batch license images to be scanned (v2.0).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHORIZEDNOTENOUGH = "FailedOperation.AuthorizedNotEnough"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyImageAuthorized(request *ModifyImageAuthorizedRequest) (response *ModifyImageAuthorizedResponse, err error) {
    return c.ModifyImageAuthorizedWithContext(context.Background(), request)
}

// ModifyImageAuthorized
// This API is used to batch license images to be scanned (v2.0).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHORIZEDNOTENOUGH = "FailedOperation.AuthorizedNotEnough"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyImageAuthorizedWithContext(ctx context.Context, request *ModifyImageAuthorizedRequest) (response *ModifyImageAuthorizedResponse, err error) {
    if request == nil {
        request = NewModifyImageAuthorizedRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyImageAuthorized require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyImageAuthorizedResponse()
    err = c.Send(request, response)
    return
}

func NewModifyK8sApiAbnormalEventStatusRequest() (request *ModifyK8sApiAbnormalEventStatusRequest) {
    request = &ModifyK8sApiAbnormalEventStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyK8sApiAbnormalEventStatus")
    
    
    return
}

func NewModifyK8sApiAbnormalEventStatusResponse() (response *ModifyK8sApiAbnormalEventStatusResponse) {
    response = &ModifyK8sApiAbnormalEventStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyK8sApiAbnormalEventStatus
// This API is used to modify the status of K8sApi exception events.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyK8sApiAbnormalEventStatus(request *ModifyK8sApiAbnormalEventStatusRequest) (response *ModifyK8sApiAbnormalEventStatusResponse, err error) {
    return c.ModifyK8sApiAbnormalEventStatusWithContext(context.Background(), request)
}

// ModifyK8sApiAbnormalEventStatus
// This API is used to modify the status of K8sApi exception events.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyK8sApiAbnormalEventStatusWithContext(ctx context.Context, request *ModifyK8sApiAbnormalEventStatusRequest) (response *ModifyK8sApiAbnormalEventStatusResponse, err error) {
    if request == nil {
        request = NewModifyK8sApiAbnormalEventStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyK8sApiAbnormalEventStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyK8sApiAbnormalEventStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyK8sApiAbnormalRuleInfoRequest() (request *ModifyK8sApiAbnormalRuleInfoRequest) {
    request = &ModifyK8sApiAbnormalRuleInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyK8sApiAbnormalRuleInfo")
    
    
    return
}

func NewModifyK8sApiAbnormalRuleInfoResponse() (response *ModifyK8sApiAbnormalRuleInfoResponse) {
    response = &ModifyK8sApiAbnormalRuleInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyK8sApiAbnormalRuleInfo
// This API is used to modify the information of K8sApi abnormal rules.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyK8sApiAbnormalRuleInfo(request *ModifyK8sApiAbnormalRuleInfoRequest) (response *ModifyK8sApiAbnormalRuleInfoResponse, err error) {
    return c.ModifyK8sApiAbnormalRuleInfoWithContext(context.Background(), request)
}

// ModifyK8sApiAbnormalRuleInfo
// This API is used to modify the information of K8sApi abnormal rules.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyK8sApiAbnormalRuleInfoWithContext(ctx context.Context, request *ModifyK8sApiAbnormalRuleInfoRequest) (response *ModifyK8sApiAbnormalRuleInfoResponse, err error) {
    if request == nil {
        request = NewModifyK8sApiAbnormalRuleInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyK8sApiAbnormalRuleInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyK8sApiAbnormalRuleInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyK8sApiAbnormalRuleStatusRequest() (request *ModifyK8sApiAbnormalRuleStatusRequest) {
    request = &ModifyK8sApiAbnormalRuleStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyK8sApiAbnormalRuleStatus")
    
    
    return
}

func NewModifyK8sApiAbnormalRuleStatusResponse() (response *ModifyK8sApiAbnormalRuleStatusResponse) {
    response = &ModifyK8sApiAbnormalRuleStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyK8sApiAbnormalRuleStatus
// This API is used to modify the status of K8sApi abnormal event rules.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyK8sApiAbnormalRuleStatus(request *ModifyK8sApiAbnormalRuleStatusRequest) (response *ModifyK8sApiAbnormalRuleStatusResponse, err error) {
    return c.ModifyK8sApiAbnormalRuleStatusWithContext(context.Background(), request)
}

// ModifyK8sApiAbnormalRuleStatus
// This API is used to modify the status of K8sApi abnormal event rules.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyK8sApiAbnormalRuleStatusWithContext(ctx context.Context, request *ModifyK8sApiAbnormalRuleStatusRequest) (response *ModifyK8sApiAbnormalRuleStatusResponse, err error) {
    if request == nil {
        request = NewModifyK8sApiAbnormalRuleStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyK8sApiAbnormalRuleStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyK8sApiAbnormalRuleStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyReverseShellStatusRequest() (request *ModifyReverseShellStatusRequest) {
    request = &ModifyReverseShellStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyReverseShellStatus")
    
    
    return
}

func NewModifyReverseShellStatusResponse() (response *ModifyReverseShellStatusResponse) {
    response = &ModifyReverseShellStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyReverseShellStatus
// This API is used to change the status of a reverse shell event.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATANOTFOUND = "InvalidParameterValue.DataNotFound"
func (c *Client) ModifyReverseShellStatus(request *ModifyReverseShellStatusRequest) (response *ModifyReverseShellStatusResponse, err error) {
    return c.ModifyReverseShellStatusWithContext(context.Background(), request)
}

// ModifyReverseShellStatus
// This API is used to change the status of a reverse shell event.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATANOTFOUND = "InvalidParameterValue.DataNotFound"
func (c *Client) ModifyReverseShellStatusWithContext(ctx context.Context, request *ModifyReverseShellStatusRequest) (response *ModifyReverseShellStatusResponse, err error) {
    if request == nil {
        request = NewModifyReverseShellStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyReverseShellStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyReverseShellStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRiskSyscallStatusRequest() (request *ModifyRiskSyscallStatusRequest) {
    request = &ModifyRiskSyscallStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyRiskSyscallStatus")
    
    
    return
}

func NewModifyRiskSyscallStatusResponse() (response *ModifyRiskSyscallStatusResponse) {
    response = &ModifyRiskSyscallStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRiskSyscallStatus
// This API is used to change the status of a high-risk syscall event.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATANOTFOUND = "InvalidParameterValue.DataNotFound"
func (c *Client) ModifyRiskSyscallStatus(request *ModifyRiskSyscallStatusRequest) (response *ModifyRiskSyscallStatusResponse, err error) {
    return c.ModifyRiskSyscallStatusWithContext(context.Background(), request)
}

// ModifyRiskSyscallStatus
// This API is used to change the status of a high-risk syscall event.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATANOTFOUND = "InvalidParameterValue.DataNotFound"
func (c *Client) ModifyRiskSyscallStatusWithContext(ctx context.Context, request *ModifyRiskSyscallStatusRequest) (response *ModifyRiskSyscallStatusResponse, err error) {
    if request == nil {
        request = NewModifyRiskSyscallStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRiskSyscallStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRiskSyscallStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifySecLogCleanSettingInfoRequest() (request *ModifySecLogCleanSettingInfoRequest) {
    request = &ModifySecLogCleanSettingInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "ModifySecLogCleanSettingInfo")
    
    
    return
}

func NewModifySecLogCleanSettingInfoResponse() (response *ModifySecLogCleanSettingInfoResponse) {
    response = &ModifySecLogCleanSettingInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySecLogCleanSettingInfo
// This API is used to modify the settings of security log cleanup.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifySecLogCleanSettingInfo(request *ModifySecLogCleanSettingInfoRequest) (response *ModifySecLogCleanSettingInfoResponse, err error) {
    return c.ModifySecLogCleanSettingInfoWithContext(context.Background(), request)
}

// ModifySecLogCleanSettingInfo
// This API is used to modify the settings of security log cleanup.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifySecLogCleanSettingInfoWithContext(ctx context.Context, request *ModifySecLogCleanSettingInfoRequest) (response *ModifySecLogCleanSettingInfoResponse, err error) {
    if request == nil {
        request = NewModifySecLogCleanSettingInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySecLogCleanSettingInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySecLogCleanSettingInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifySecLogDeliveryClsSettingRequest() (request *ModifySecLogDeliveryClsSettingRequest) {
    request = &ModifySecLogDeliveryClsSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "ModifySecLogDeliveryClsSetting")
    
    
    return
}

func NewModifySecLogDeliveryClsSettingResponse() (response *ModifySecLogDeliveryClsSettingResponse) {
    response = &ModifySecLogDeliveryClsSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySecLogDeliveryClsSetting
// This API is used to update the configuration of security log delivery to CLS.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifySecLogDeliveryClsSetting(request *ModifySecLogDeliveryClsSettingRequest) (response *ModifySecLogDeliveryClsSettingResponse, err error) {
    return c.ModifySecLogDeliveryClsSettingWithContext(context.Background(), request)
}

// ModifySecLogDeliveryClsSetting
// This API is used to update the configuration of security log delivery to CLS.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifySecLogDeliveryClsSettingWithContext(ctx context.Context, request *ModifySecLogDeliveryClsSettingRequest) (response *ModifySecLogDeliveryClsSettingResponse, err error) {
    if request == nil {
        request = NewModifySecLogDeliveryClsSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySecLogDeliveryClsSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySecLogDeliveryClsSettingResponse()
    err = c.Send(request, response)
    return
}

func NewModifySecLogDeliveryKafkaSettingRequest() (request *ModifySecLogDeliveryKafkaSettingRequest) {
    request = &ModifySecLogDeliveryKafkaSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "ModifySecLogDeliveryKafkaSetting")
    
    
    return
}

func NewModifySecLogDeliveryKafkaSettingResponse() (response *ModifySecLogDeliveryKafkaSettingResponse) {
    response = &ModifySecLogDeliveryKafkaSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySecLogDeliveryKafkaSetting
// This API is used to update the settings of security log delivery to Kafka.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifySecLogDeliveryKafkaSetting(request *ModifySecLogDeliveryKafkaSettingRequest) (response *ModifySecLogDeliveryKafkaSettingResponse, err error) {
    return c.ModifySecLogDeliveryKafkaSettingWithContext(context.Background(), request)
}

// ModifySecLogDeliveryKafkaSetting
// This API is used to update the settings of security log delivery to Kafka.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifySecLogDeliveryKafkaSettingWithContext(ctx context.Context, request *ModifySecLogDeliveryKafkaSettingRequest) (response *ModifySecLogDeliveryKafkaSettingResponse, err error) {
    if request == nil {
        request = NewModifySecLogDeliveryKafkaSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySecLogDeliveryKafkaSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySecLogDeliveryKafkaSettingResponse()
    err = c.Send(request, response)
    return
}

func NewModifySecLogJoinObjectsRequest() (request *ModifySecLogJoinObjectsRequest) {
    request = &ModifySecLogJoinObjectsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "ModifySecLogJoinObjects")
    
    
    return
}

func NewModifySecLogJoinObjectsResponse() (response *ModifySecLogJoinObjectsResponse) {
    response = &ModifySecLogJoinObjectsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySecLogJoinObjects
// This API is used to modify an accessed security log object.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifySecLogJoinObjects(request *ModifySecLogJoinObjectsRequest) (response *ModifySecLogJoinObjectsResponse, err error) {
    return c.ModifySecLogJoinObjectsWithContext(context.Background(), request)
}

// ModifySecLogJoinObjects
// This API is used to modify an accessed security log object.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifySecLogJoinObjectsWithContext(ctx context.Context, request *ModifySecLogJoinObjectsRequest) (response *ModifySecLogJoinObjectsResponse, err error) {
    if request == nil {
        request = NewModifySecLogJoinObjectsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySecLogJoinObjects require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySecLogJoinObjectsResponse()
    err = c.Send(request, response)
    return
}

func NewModifySecLogJoinStateRequest() (request *ModifySecLogJoinStateRequest) {
    request = &ModifySecLogJoinStateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "ModifySecLogJoinState")
    
    
    return
}

func NewModifySecLogJoinStateResponse() (response *ModifySecLogJoinStateResponse) {
    response = &ModifySecLogJoinStateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySecLogJoinState
// This API is used to change the security log access status.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifySecLogJoinState(request *ModifySecLogJoinStateRequest) (response *ModifySecLogJoinStateResponse, err error) {
    return c.ModifySecLogJoinStateWithContext(context.Background(), request)
}

// ModifySecLogJoinState
// This API is used to change the security log access status.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifySecLogJoinStateWithContext(ctx context.Context, request *ModifySecLogJoinStateRequest) (response *ModifySecLogJoinStateResponse, err error) {
    if request == nil {
        request = NewModifySecLogJoinStateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySecLogJoinState require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySecLogJoinStateResponse()
    err = c.Send(request, response)
    return
}

func NewModifySecLogKafkaUINRequest() (request *ModifySecLogKafkaUINRequest) {
    request = &ModifySecLogKafkaUINRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "ModifySecLogKafkaUIN")
    
    
    return
}

func NewModifySecLogKafkaUINResponse() (response *ModifySecLogKafkaUINResponse) {
    response = &ModifySecLogKafkaUINResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySecLogKafkaUIN
// This API is used to modify the UIN of a Kafka security log.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifySecLogKafkaUIN(request *ModifySecLogKafkaUINRequest) (response *ModifySecLogKafkaUINResponse, err error) {
    return c.ModifySecLogKafkaUINWithContext(context.Background(), request)
}

// ModifySecLogKafkaUIN
// This API is used to modify the UIN of a Kafka security log.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifySecLogKafkaUINWithContext(ctx context.Context, request *ModifySecLogKafkaUINRequest) (response *ModifySecLogKafkaUINResponse, err error) {
    if request == nil {
        request = NewModifySecLogKafkaUINRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySecLogKafkaUIN require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySecLogKafkaUINResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVirusAutoIsolateExampleSwitchRequest() (request *ModifyVirusAutoIsolateExampleSwitchRequest) {
    request = &ModifyVirusAutoIsolateExampleSwitchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyVirusAutoIsolateExampleSwitch")
    
    
    return
}

func NewModifyVirusAutoIsolateExampleSwitchResponse() (response *ModifyVirusAutoIsolateExampleSwitchResponse) {
    response = &ModifyVirusAutoIsolateExampleSwitchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyVirusAutoIsolateExampleSwitch
// This API is used to enable/disable automatic trojan sample isolation.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyVirusAutoIsolateExampleSwitch(request *ModifyVirusAutoIsolateExampleSwitchRequest) (response *ModifyVirusAutoIsolateExampleSwitchResponse, err error) {
    return c.ModifyVirusAutoIsolateExampleSwitchWithContext(context.Background(), request)
}

// ModifyVirusAutoIsolateExampleSwitch
// This API is used to enable/disable automatic trojan sample isolation.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyVirusAutoIsolateExampleSwitchWithContext(ctx context.Context, request *ModifyVirusAutoIsolateExampleSwitchRequest) (response *ModifyVirusAutoIsolateExampleSwitchResponse, err error) {
    if request == nil {
        request = NewModifyVirusAutoIsolateExampleSwitchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVirusAutoIsolateExampleSwitch require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyVirusAutoIsolateExampleSwitchResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVirusAutoIsolateSettingRequest() (request *ModifyVirusAutoIsolateSettingRequest) {
    request = &ModifyVirusAutoIsolateSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyVirusAutoIsolateSetting")
    
    
    return
}

func NewModifyVirusAutoIsolateSettingResponse() (response *ModifyVirusAutoIsolateSettingResponse) {
    response = &ModifyVirusAutoIsolateSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyVirusAutoIsolateSetting
// This API is used to modify the settings of automatic trojan isolation.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyVirusAutoIsolateSetting(request *ModifyVirusAutoIsolateSettingRequest) (response *ModifyVirusAutoIsolateSettingResponse, err error) {
    return c.ModifyVirusAutoIsolateSettingWithContext(context.Background(), request)
}

// ModifyVirusAutoIsolateSetting
// This API is used to modify the settings of automatic trojan isolation.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyVirusAutoIsolateSettingWithContext(ctx context.Context, request *ModifyVirusAutoIsolateSettingRequest) (response *ModifyVirusAutoIsolateSettingResponse, err error) {
    if request == nil {
        request = NewModifyVirusAutoIsolateSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVirusAutoIsolateSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyVirusAutoIsolateSettingResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVirusFileStatusRequest() (request *ModifyVirusFileStatusRequest) {
    request = &ModifyVirusFileStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyVirusFileStatus")
    
    
    return
}

func NewModifyVirusFileStatusResponse() (response *ModifyVirusFileStatusResponse) {
    response = &ModifyVirusFileStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyVirusFileStatus
// This API is used to update the status of a trojan file at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATANOTFOUND = "InvalidParameterValue.DataNotFound"
func (c *Client) ModifyVirusFileStatus(request *ModifyVirusFileStatusRequest) (response *ModifyVirusFileStatusResponse, err error) {
    return c.ModifyVirusFileStatusWithContext(context.Background(), request)
}

// ModifyVirusFileStatus
// This API is used to update the status of a trojan file at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATANOTFOUND = "InvalidParameterValue.DataNotFound"
func (c *Client) ModifyVirusFileStatusWithContext(ctx context.Context, request *ModifyVirusFileStatusRequest) (response *ModifyVirusFileStatusResponse, err error) {
    if request == nil {
        request = NewModifyVirusFileStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVirusFileStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyVirusFileStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVirusMonitorSettingRequest() (request *ModifyVirusMonitorSettingRequest) {
    request = &ModifyVirusMonitorSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyVirusMonitorSetting")
    
    
    return
}

func NewModifyVirusMonitorSettingResponse() (response *ModifyVirusMonitorSettingResponse) {
    response = &ModifyVirusMonitorSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyVirusMonitorSetting
// This API is used to update the real-time monitoring settings of virus scanning at runtime.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyVirusMonitorSetting(request *ModifyVirusMonitorSettingRequest) (response *ModifyVirusMonitorSettingResponse, err error) {
    return c.ModifyVirusMonitorSettingWithContext(context.Background(), request)
}

// ModifyVirusMonitorSetting
// This API is used to update the real-time monitoring settings of virus scanning at runtime.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyVirusMonitorSettingWithContext(ctx context.Context, request *ModifyVirusMonitorSettingRequest) (response *ModifyVirusMonitorSettingResponse, err error) {
    if request == nil {
        request = NewModifyVirusMonitorSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVirusMonitorSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyVirusMonitorSettingResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVirusScanSettingRequest() (request *ModifyVirusScanSettingRequest) {
    request = &ModifyVirusScanSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyVirusScanSetting")
    
    
    return
}

func NewModifyVirusScanSettingResponse() (response *ModifyVirusScanSettingResponse) {
    response = &ModifyVirusScanSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyVirusScanSetting
// This API is used to update virus scanning settings at runtime.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyVirusScanSetting(request *ModifyVirusScanSettingRequest) (response *ModifyVirusScanSettingResponse, err error) {
    return c.ModifyVirusScanSettingWithContext(context.Background(), request)
}

// ModifyVirusScanSetting
// This API is used to update virus scanning settings at runtime.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyVirusScanSettingWithContext(ctx context.Context, request *ModifyVirusScanSettingRequest) (response *ModifyVirusScanSettingResponse, err error) {
    if request == nil {
        request = NewModifyVirusScanSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVirusScanSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyVirusScanSettingResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVirusScanTimeoutSettingRequest() (request *ModifyVirusScanTimeoutSettingRequest) {
    request = &ModifyVirusScanTimeoutSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyVirusScanTimeoutSetting")
    
    
    return
}

func NewModifyVirusScanTimeoutSettingResponse() (response *ModifyVirusScanTimeoutSettingResponse) {
    response = &ModifyVirusScanTimeoutSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyVirusScanTimeoutSetting
// This API is used to modify the timeout settings of a file scan at runtime.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyVirusScanTimeoutSetting(request *ModifyVirusScanTimeoutSettingRequest) (response *ModifyVirusScanTimeoutSettingResponse, err error) {
    return c.ModifyVirusScanTimeoutSettingWithContext(context.Background(), request)
}

// ModifyVirusScanTimeoutSetting
// This API is used to modify the timeout settings of a file scan at runtime.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyVirusScanTimeoutSettingWithContext(ctx context.Context, request *ModifyVirusScanTimeoutSettingRequest) (response *ModifyVirusScanTimeoutSettingResponse, err error) {
    if request == nil {
        request = NewModifyVirusScanTimeoutSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVirusScanTimeoutSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyVirusScanTimeoutSettingResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVulDefenceEventStatusRequest() (request *ModifyVulDefenceEventStatusRequest) {
    request = &ModifyVulDefenceEventStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyVulDefenceEventStatus")
    
    
    return
}

func NewModifyVulDefenceEventStatusResponse() (response *ModifyVulDefenceEventStatusResponse) {
    response = &ModifyVulDefenceEventStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyVulDefenceEventStatus
// This API is used to change the status of an exploit prevention event.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyVulDefenceEventStatus(request *ModifyVulDefenceEventStatusRequest) (response *ModifyVulDefenceEventStatusResponse, err error) {
    return c.ModifyVulDefenceEventStatusWithContext(context.Background(), request)
}

// ModifyVulDefenceEventStatus
// This API is used to change the status of an exploit prevention event.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyVulDefenceEventStatusWithContext(ctx context.Context, request *ModifyVulDefenceEventStatusRequest) (response *ModifyVulDefenceEventStatusResponse, err error) {
    if request == nil {
        request = NewModifyVulDefenceEventStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVulDefenceEventStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyVulDefenceEventStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVulDefenceSettingRequest() (request *ModifyVulDefenceSettingRequest) {
    request = &ModifyVulDefenceSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyVulDefenceSetting")
    
    
    return
}

func NewModifyVulDefenceSettingResponse() (response *ModifyVulDefenceSettingResponse) {
    response = &ModifyVulDefenceSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyVulDefenceSetting
// This API is used to edit the exploit prevention settings.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyVulDefenceSetting(request *ModifyVulDefenceSettingRequest) (response *ModifyVulDefenceSettingResponse, err error) {
    return c.ModifyVulDefenceSettingWithContext(context.Background(), request)
}

// ModifyVulDefenceSetting
// This API is used to edit the exploit prevention settings.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyVulDefenceSettingWithContext(ctx context.Context, request *ModifyVulDefenceSettingRequest) (response *ModifyVulDefenceSettingResponse, err error) {
    if request == nil {
        request = NewModifyVulDefenceSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVulDefenceSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyVulDefenceSettingResponse()
    err = c.Send(request, response)
    return
}

func NewOpenTcssTrialRequest() (request *OpenTcssTrialRequest) {
    request = &OpenTcssTrialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "OpenTcssTrial")
    
    
    return
}

func NewOpenTcssTrialResponse() (response *OpenTcssTrialResponse) {
    response = &OpenTcssTrialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OpenTcssTrial
// This API is used to activate TCSS trial.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) OpenTcssTrial(request *OpenTcssTrialRequest) (response *OpenTcssTrialResponse, err error) {
    return c.OpenTcssTrialWithContext(context.Background(), request)
}

// OpenTcssTrial
// This API is used to activate TCSS trial.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) OpenTcssTrialWithContext(ctx context.Context, request *OpenTcssTrialRequest) (response *OpenTcssTrialResponse, err error) {
    if request == nil {
        request = NewOpenTcssTrialRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenTcssTrial require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenTcssTrialResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveAssetImageRegistryRegistryDetailRequest() (request *RemoveAssetImageRegistryRegistryDetailRequest) {
    request = &RemoveAssetImageRegistryRegistryDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "RemoveAssetImageRegistryRegistryDetail")
    
    
    return
}

func NewRemoveAssetImageRegistryRegistryDetailResponse() (response *RemoveAssetImageRegistryRegistryDetailResponse) {
    response = &RemoveAssetImageRegistryRegistryDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemoveAssetImageRegistryRegistryDetail
// This API is used to delete the details of an image repository.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) RemoveAssetImageRegistryRegistryDetail(request *RemoveAssetImageRegistryRegistryDetailRequest) (response *RemoveAssetImageRegistryRegistryDetailResponse, err error) {
    return c.RemoveAssetImageRegistryRegistryDetailWithContext(context.Background(), request)
}

// RemoveAssetImageRegistryRegistryDetail
// This API is used to delete the details of an image repository.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) RemoveAssetImageRegistryRegistryDetailWithContext(ctx context.Context, request *RemoveAssetImageRegistryRegistryDetailRequest) (response *RemoveAssetImageRegistryRegistryDetailResponse, err error) {
    if request == nil {
        request = NewRemoveAssetImageRegistryRegistryDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveAssetImageRegistryRegistryDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveAssetImageRegistryRegistryDetailResponse()
    err = c.Send(request, response)
    return
}

func NewRenewImageAuthorizeStateRequest() (request *RenewImageAuthorizeStateRequest) {
    request = &RenewImageAuthorizeStateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "RenewImageAuthorizeState")
    
    
    return
}

func NewRenewImageAuthorizeStateResponse() (response *RenewImageAuthorizeStateResponse) {
    response = &RenewImageAuthorizeStateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RenewImageAuthorizeState
// This API is used to license an image to be scanned.
//
// error code that may be returned:
//  FAILEDOPERATION_AUTHORIZEDNOTENOUGH = "FailedOperation.AuthorizedNotEnough"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) RenewImageAuthorizeState(request *RenewImageAuthorizeStateRequest) (response *RenewImageAuthorizeStateResponse, err error) {
    return c.RenewImageAuthorizeStateWithContext(context.Background(), request)
}

// RenewImageAuthorizeState
// This API is used to license an image to be scanned.
//
// error code that may be returned:
//  FAILEDOPERATION_AUTHORIZEDNOTENOUGH = "FailedOperation.AuthorizedNotEnough"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) RenewImageAuthorizeStateWithContext(ctx context.Context, request *RenewImageAuthorizeStateRequest) (response *RenewImageAuthorizeStateResponse, err error) {
    if request == nil {
        request = NewRenewImageAuthorizeStateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RenewImageAuthorizeState require credential")
    }

    request.SetContext(ctx)
    
    response = NewRenewImageAuthorizeStateResponse()
    err = c.Send(request, response)
    return
}

func NewResetSecLogTopicConfigRequest() (request *ResetSecLogTopicConfigRequest) {
    request = &ResetSecLogTopicConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "ResetSecLogTopicConfig")
    
    
    return
}

func NewResetSecLogTopicConfigResponse() (response *ResetSecLogTopicConfigResponse) {
    response = &ResetSecLogTopicConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetSecLogTopicConfig
// This API is used to reset a security log topic.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ResetSecLogTopicConfig(request *ResetSecLogTopicConfigRequest) (response *ResetSecLogTopicConfigResponse, err error) {
    return c.ResetSecLogTopicConfigWithContext(context.Background(), request)
}

// ResetSecLogTopicConfig
// This API is used to reset a security log topic.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ResetSecLogTopicConfigWithContext(ctx context.Context, request *ResetSecLogTopicConfigRequest) (response *ResetSecLogTopicConfigResponse, err error) {
    if request == nil {
        request = NewResetSecLogTopicConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetSecLogTopicConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetSecLogTopicConfigResponse()
    err = c.Send(request, response)
    return
}

func NewScanComplianceAssetsRequest() (request *ScanComplianceAssetsRequest) {
    request = &ScanComplianceAssetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "ScanComplianceAssets")
    
    
    return
}

func NewScanComplianceAssetsResponse() (response *ScanComplianceAssetsResponse) {
    response = &ScanComplianceAssetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ScanComplianceAssets
// This API is used to check the specified asset again.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ScanComplianceAssets(request *ScanComplianceAssetsRequest) (response *ScanComplianceAssetsResponse, err error) {
    return c.ScanComplianceAssetsWithContext(context.Background(), request)
}

// ScanComplianceAssets
// This API is used to check the specified asset again.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ScanComplianceAssetsWithContext(ctx context.Context, request *ScanComplianceAssetsRequest) (response *ScanComplianceAssetsResponse, err error) {
    if request == nil {
        request = NewScanComplianceAssetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ScanComplianceAssets require credential")
    }

    request.SetContext(ctx)
    
    response = NewScanComplianceAssetsResponse()
    err = c.Send(request, response)
    return
}

func NewScanComplianceAssetsByPolicyItemRequest() (request *ScanComplianceAssetsByPolicyItemRequest) {
    request = &ScanComplianceAssetsByPolicyItemRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "ScanComplianceAssetsByPolicyItem")
    
    
    return
}

func NewScanComplianceAssetsByPolicyItemResponse() (response *ScanComplianceAssetsByPolicyItemResponse) {
    response = &ScanComplianceAssetsByPolicyItemResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ScanComplianceAssetsByPolicyItem
// This API is used to check the specified asset again with the specified check item and return the ID of the created compliance check task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ScanComplianceAssetsByPolicyItem(request *ScanComplianceAssetsByPolicyItemRequest) (response *ScanComplianceAssetsByPolicyItemResponse, err error) {
    return c.ScanComplianceAssetsByPolicyItemWithContext(context.Background(), request)
}

// ScanComplianceAssetsByPolicyItem
// This API is used to check the specified asset again with the specified check item and return the ID of the created compliance check task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ScanComplianceAssetsByPolicyItemWithContext(ctx context.Context, request *ScanComplianceAssetsByPolicyItemRequest) (response *ScanComplianceAssetsByPolicyItemResponse, err error) {
    if request == nil {
        request = NewScanComplianceAssetsByPolicyItemRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ScanComplianceAssetsByPolicyItem require credential")
    }

    request.SetContext(ctx)
    
    response = NewScanComplianceAssetsByPolicyItemResponse()
    err = c.Send(request, response)
    return
}

func NewScanCompliancePolicyItemsRequest() (request *ScanCompliancePolicyItemsRequest) {
    request = &ScanCompliancePolicyItemsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "ScanCompliancePolicyItems")
    
    
    return
}

func NewScanCompliancePolicyItemsResponse() (response *ScanCompliancePolicyItemsResponse) {
    response = &ScanCompliancePolicyItemsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ScanCompliancePolicyItems
// This API is used to check all the assets of the specified check item again and return the ID of the created compliance check task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ScanCompliancePolicyItems(request *ScanCompliancePolicyItemsRequest) (response *ScanCompliancePolicyItemsResponse, err error) {
    return c.ScanCompliancePolicyItemsWithContext(context.Background(), request)
}

// ScanCompliancePolicyItems
// This API is used to check all the assets of the specified check item again and return the ID of the created compliance check task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ScanCompliancePolicyItemsWithContext(ctx context.Context, request *ScanCompliancePolicyItemsRequest) (response *ScanCompliancePolicyItemsResponse, err error) {
    if request == nil {
        request = NewScanCompliancePolicyItemsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ScanCompliancePolicyItems require credential")
    }

    request.SetContext(ctx)
    
    response = NewScanCompliancePolicyItemsResponse()
    err = c.Send(request, response)
    return
}

func NewScanComplianceScanFailedAssetsRequest() (request *ScanComplianceScanFailedAssetsRequest) {
    request = &ScanComplianceScanFailedAssetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "ScanComplianceScanFailedAssets")
    
    
    return
}

func NewScanComplianceScanFailedAssetsResponse() (response *ScanComplianceScanFailedAssetsResponse) {
    response = &ScanComplianceScanFailedAssetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ScanComplianceScanFailedAssets
// This API is used to check all the failed check items of the specified asset again and return the ID of the created compliance check task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ScanComplianceScanFailedAssets(request *ScanComplianceScanFailedAssetsRequest) (response *ScanComplianceScanFailedAssetsResponse, err error) {
    return c.ScanComplianceScanFailedAssetsWithContext(context.Background(), request)
}

// ScanComplianceScanFailedAssets
// This API is used to check all the failed check items of the specified asset again and return the ID of the created compliance check task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ScanComplianceScanFailedAssetsWithContext(ctx context.Context, request *ScanComplianceScanFailedAssetsRequest) (response *ScanComplianceScanFailedAssetsResponse, err error) {
    if request == nil {
        request = NewScanComplianceScanFailedAssetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ScanComplianceScanFailedAssets require credential")
    }

    request.SetContext(ctx)
    
    response = NewScanComplianceScanFailedAssetsResponse()
    err = c.Send(request, response)
    return
}

func NewSetCheckModeRequest() (request *SetCheckModeRequest) {
    request = &SetCheckModeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "SetCheckMode")
    
    
    return
}

func NewSetCheckModeResponse() (response *SetCheckModeResponse) {
    response = &SetCheckModeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetCheckMode
// This API is used to set the check mode and automatic check.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) SetCheckMode(request *SetCheckModeRequest) (response *SetCheckModeResponse, err error) {
    return c.SetCheckModeWithContext(context.Background(), request)
}

// SetCheckMode
// This API is used to set the check mode and automatic check.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) SetCheckModeWithContext(ctx context.Context, request *SetCheckModeRequest) (response *SetCheckModeResponse, err error) {
    if request == nil {
        request = NewSetCheckModeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetCheckMode require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetCheckModeResponse()
    err = c.Send(request, response)
    return
}

func NewStopVirusScanTaskRequest() (request *StopVirusScanTaskRequest) {
    request = &StopVirusScanTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "StopVirusScanTask")
    
    
    return
}

func NewStopVirusScanTaskResponse() (response *StopVirusScanTaskResponse) {
    response = &StopVirusScanTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopVirusScanTask
// This API is used to stop a trojan scan task at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) StopVirusScanTask(request *StopVirusScanTaskRequest) (response *StopVirusScanTaskResponse, err error) {
    return c.StopVirusScanTaskWithContext(context.Background(), request)
}

// StopVirusScanTask
// This API is used to stop a trojan scan task at runtime.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) StopVirusScanTaskWithContext(ctx context.Context, request *StopVirusScanTaskRequest) (response *StopVirusScanTaskResponse, err error) {
    if request == nil {
        request = NewStopVirusScanTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopVirusScanTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopVirusScanTaskResponse()
    err = c.Send(request, response)
    return
}

func NewStopVulScanTaskRequest() (request *StopVulScanTaskRequest) {
    request = &StopVulScanTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "StopVulScanTask")
    
    
    return
}

func NewStopVulScanTaskResponse() (response *StopVulScanTaskResponse) {
    response = &StopVulScanTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopVulScanTask
// This API is used to stop a vulnerability scan task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) StopVulScanTask(request *StopVulScanTaskRequest) (response *StopVulScanTaskResponse, err error) {
    return c.StopVulScanTaskWithContext(context.Background(), request)
}

// StopVulScanTask
// This API is used to stop a vulnerability scan task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) StopVulScanTaskWithContext(ctx context.Context, request *StopVulScanTaskRequest) (response *StopVulScanTaskResponse, err error) {
    if request == nil {
        request = NewStopVulScanTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopVulScanTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopVulScanTaskResponse()
    err = c.Send(request, response)
    return
}

func NewSwitchImageAutoAuthorizedRuleRequest() (request *SwitchImageAutoAuthorizedRuleRequest) {
    request = &SwitchImageAutoAuthorizedRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "SwitchImageAutoAuthorizedRule")
    
    
    return
}

func NewSwitchImageAutoAuthorizedRuleResponse() (response *SwitchImageAutoAuthorizedRuleResponse) {
    response = &SwitchImageAutoAuthorizedRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SwitchImageAutoAuthorizedRule
// This API is used to enable/disable automatic licensing for local images.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRRULENOTFIND = "FailedOperation.ErrRuleNotFind"
//  FAILEDOPERATION_RULENOTFIND = "FailedOperation.RuleNotFind"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SwitchImageAutoAuthorizedRule(request *SwitchImageAutoAuthorizedRuleRequest) (response *SwitchImageAutoAuthorizedRuleResponse, err error) {
    return c.SwitchImageAutoAuthorizedRuleWithContext(context.Background(), request)
}

// SwitchImageAutoAuthorizedRule
// This API is used to enable/disable automatic licensing for local images.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRRULENOTFIND = "FailedOperation.ErrRuleNotFind"
//  FAILEDOPERATION_RULENOTFIND = "FailedOperation.RuleNotFind"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SwitchImageAutoAuthorizedRuleWithContext(ctx context.Context, request *SwitchImageAutoAuthorizedRuleRequest) (response *SwitchImageAutoAuthorizedRuleResponse, err error) {
    if request == nil {
        request = NewSwitchImageAutoAuthorizedRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SwitchImageAutoAuthorizedRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewSwitchImageAutoAuthorizedRuleResponse()
    err = c.Send(request, response)
    return
}

func NewSyncAssetImageRegistryAssetRequest() (request *SyncAssetImageRegistryAssetRequest) {
    request = &SyncAssetImageRegistryAssetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "SyncAssetImageRegistryAsset")
    
    
    return
}

func NewSyncAssetImageRegistryAssetResponse() (response *SyncAssetImageRegistryAssetResponse) {
    response = &SyncAssetImageRegistryAssetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SyncAssetImageRegistryAsset
// This API is used to refresh the assets in an image repository.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SyncAssetImageRegistryAsset(request *SyncAssetImageRegistryAssetRequest) (response *SyncAssetImageRegistryAssetResponse, err error) {
    return c.SyncAssetImageRegistryAssetWithContext(context.Background(), request)
}

// SyncAssetImageRegistryAsset
// This API is used to refresh the assets in an image repository.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SyncAssetImageRegistryAssetWithContext(ctx context.Context, request *SyncAssetImageRegistryAssetRequest) (response *SyncAssetImageRegistryAssetResponse, err error) {
    if request == nil {
        request = NewSyncAssetImageRegistryAssetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SyncAssetImageRegistryAsset require credential")
    }

    request.SetContext(ctx)
    
    response = NewSyncAssetImageRegistryAssetResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAndPublishNetworkFirewallPolicyDetailRequest() (request *UpdateAndPublishNetworkFirewallPolicyDetailRequest) {
    request = &UpdateAndPublishNetworkFirewallPolicyDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "UpdateAndPublishNetworkFirewallPolicyDetail")
    
    
    return
}

func NewUpdateAndPublishNetworkFirewallPolicyDetailResponse() (response *UpdateAndPublishNetworkFirewallPolicyDetailResponse) {
    response = &UpdateAndPublishNetworkFirewallPolicyDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateAndPublishNetworkFirewallPolicyDetail
// This API is used to create a task to update and publish a network policy in the container network.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) UpdateAndPublishNetworkFirewallPolicyDetail(request *UpdateAndPublishNetworkFirewallPolicyDetailRequest) (response *UpdateAndPublishNetworkFirewallPolicyDetailResponse, err error) {
    return c.UpdateAndPublishNetworkFirewallPolicyDetailWithContext(context.Background(), request)
}

// UpdateAndPublishNetworkFirewallPolicyDetail
// This API is used to create a task to update and publish a network policy in the container network.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) UpdateAndPublishNetworkFirewallPolicyDetailWithContext(ctx context.Context, request *UpdateAndPublishNetworkFirewallPolicyDetailRequest) (response *UpdateAndPublishNetworkFirewallPolicyDetailResponse, err error) {
    if request == nil {
        request = NewUpdateAndPublishNetworkFirewallPolicyDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateAndPublishNetworkFirewallPolicyDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateAndPublishNetworkFirewallPolicyDetailResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAndPublishNetworkFirewallPolicyYamlDetailRequest() (request *UpdateAndPublishNetworkFirewallPolicyYamlDetailRequest) {
    request = &UpdateAndPublishNetworkFirewallPolicyYamlDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "UpdateAndPublishNetworkFirewallPolicyYamlDetail")
    
    
    return
}

func NewUpdateAndPublishNetworkFirewallPolicyYamlDetailResponse() (response *UpdateAndPublishNetworkFirewallPolicyYamlDetailResponse) {
    response = &UpdateAndPublishNetworkFirewallPolicyYamlDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateAndPublishNetworkFirewallPolicyYamlDetail
// This API is used to create a task to update and publish a YAML network policy in the container network.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) UpdateAndPublishNetworkFirewallPolicyYamlDetail(request *UpdateAndPublishNetworkFirewallPolicyYamlDetailRequest) (response *UpdateAndPublishNetworkFirewallPolicyYamlDetailResponse, err error) {
    return c.UpdateAndPublishNetworkFirewallPolicyYamlDetailWithContext(context.Background(), request)
}

// UpdateAndPublishNetworkFirewallPolicyYamlDetail
// This API is used to create a task to update and publish a YAML network policy in the container network.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) UpdateAndPublishNetworkFirewallPolicyYamlDetailWithContext(ctx context.Context, request *UpdateAndPublishNetworkFirewallPolicyYamlDetailRequest) (response *UpdateAndPublishNetworkFirewallPolicyYamlDetailResponse, err error) {
    if request == nil {
        request = NewUpdateAndPublishNetworkFirewallPolicyYamlDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateAndPublishNetworkFirewallPolicyYamlDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateAndPublishNetworkFirewallPolicyYamlDetailResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAssetImageRegistryRegistryDetailRequest() (request *UpdateAssetImageRegistryRegistryDetailRequest) {
    request = &UpdateAssetImageRegistryRegistryDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "UpdateAssetImageRegistryRegistryDetail")
    
    
    return
}

func NewUpdateAssetImageRegistryRegistryDetailResponse() (response *UpdateAssetImageRegistryRegistryDetailResponse) {
    response = &UpdateAssetImageRegistryRegistryDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateAssetImageRegistryRegistryDetail
// This API is used to update the details of an image repository.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateAssetImageRegistryRegistryDetail(request *UpdateAssetImageRegistryRegistryDetailRequest) (response *UpdateAssetImageRegistryRegistryDetailResponse, err error) {
    return c.UpdateAssetImageRegistryRegistryDetailWithContext(context.Background(), request)
}

// UpdateAssetImageRegistryRegistryDetail
// This API is used to update the details of an image repository.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateAssetImageRegistryRegistryDetailWithContext(ctx context.Context, request *UpdateAssetImageRegistryRegistryDetailRequest) (response *UpdateAssetImageRegistryRegistryDetailResponse, err error) {
    if request == nil {
        request = NewUpdateAssetImageRegistryRegistryDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateAssetImageRegistryRegistryDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateAssetImageRegistryRegistryDetailResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateImageRegistryTimingScanTaskRequest() (request *UpdateImageRegistryTimingScanTaskRequest) {
    request = &UpdateImageRegistryTimingScanTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "UpdateImageRegistryTimingScanTask")
    
    
    return
}

func NewUpdateImageRegistryTimingScanTaskResponse() (response *UpdateImageRegistryTimingScanTaskResponse) {
    response = &UpdateImageRegistryTimingScanTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateImageRegistryTimingScanTask
// This API is used to update a scheduled task for an image repository.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateImageRegistryTimingScanTask(request *UpdateImageRegistryTimingScanTaskRequest) (response *UpdateImageRegistryTimingScanTaskResponse, err error) {
    return c.UpdateImageRegistryTimingScanTaskWithContext(context.Background(), request)
}

// UpdateImageRegistryTimingScanTask
// This API is used to update a scheduled task for an image repository.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateImageRegistryTimingScanTaskWithContext(ctx context.Context, request *UpdateImageRegistryTimingScanTaskRequest) (response *UpdateImageRegistryTimingScanTaskResponse, err error) {
    if request == nil {
        request = NewUpdateImageRegistryTimingScanTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateImageRegistryTimingScanTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateImageRegistryTimingScanTaskResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateNetworkFirewallPolicyDetailRequest() (request *UpdateNetworkFirewallPolicyDetailRequest) {
    request = &UpdateNetworkFirewallPolicyDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "UpdateNetworkFirewallPolicyDetail")
    
    
    return
}

func NewUpdateNetworkFirewallPolicyDetailResponse() (response *UpdateNetworkFirewallPolicyDetailResponse) {
    response = &UpdateNetworkFirewallPolicyDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateNetworkFirewallPolicyDetail
// This API is used to create a task to update a network policy in the container network.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) UpdateNetworkFirewallPolicyDetail(request *UpdateNetworkFirewallPolicyDetailRequest) (response *UpdateNetworkFirewallPolicyDetailResponse, err error) {
    return c.UpdateNetworkFirewallPolicyDetailWithContext(context.Background(), request)
}

// UpdateNetworkFirewallPolicyDetail
// This API is used to create a task to update a network policy in the container network.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) UpdateNetworkFirewallPolicyDetailWithContext(ctx context.Context, request *UpdateNetworkFirewallPolicyDetailRequest) (response *UpdateNetworkFirewallPolicyDetailResponse, err error) {
    if request == nil {
        request = NewUpdateNetworkFirewallPolicyDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateNetworkFirewallPolicyDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateNetworkFirewallPolicyDetailResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateNetworkFirewallPolicyYamlDetailRequest() (request *UpdateNetworkFirewallPolicyYamlDetailRequest) {
    request = &UpdateNetworkFirewallPolicyYamlDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcss", APIVersion, "UpdateNetworkFirewallPolicyYamlDetail")
    
    
    return
}

func NewUpdateNetworkFirewallPolicyYamlDetailResponse() (response *UpdateNetworkFirewallPolicyYamlDetailResponse) {
    response = &UpdateNetworkFirewallPolicyYamlDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateNetworkFirewallPolicyYamlDetail
// This API is used to create a task to update a YAML network policy in the container network.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) UpdateNetworkFirewallPolicyYamlDetail(request *UpdateNetworkFirewallPolicyYamlDetailRequest) (response *UpdateNetworkFirewallPolicyYamlDetailResponse, err error) {
    return c.UpdateNetworkFirewallPolicyYamlDetailWithContext(context.Background(), request)
}

// UpdateNetworkFirewallPolicyYamlDetail
// This API is used to create a task to update a YAML network policy in the container network.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) UpdateNetworkFirewallPolicyYamlDetailWithContext(ctx context.Context, request *UpdateNetworkFirewallPolicyYamlDetailRequest) (response *UpdateNetworkFirewallPolicyYamlDetailResponse, err error) {
    if request == nil {
        request = NewUpdateNetworkFirewallPolicyYamlDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateNetworkFirewallPolicyYamlDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateNetworkFirewallPolicyYamlDetailResponse()
    err = c.Send(request, response)
    return
}
