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

package v20200309

import (
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2020-03-09"

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


func NewAssociateDDoSEipAddressRequest() (request *AssociateDDoSEipAddressRequest) {
    request = &AssociateDDoSEipAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "AssociateDDoSEipAddress")
    
    
    return
}

func NewAssociateDDoSEipAddressResponse() (response *AssociateDDoSEipAddressResponse) {
    response = &AssociateDDoSEipAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AssociateDDoSEipAddress
// This API is used to bind an EIP to an Anti-DDoS Advanced instance or a specified private IP of an ENI.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) AssociateDDoSEipAddress(request *AssociateDDoSEipAddressRequest) (response *AssociateDDoSEipAddressResponse, err error) {
    if request == nil {
        request = NewAssociateDDoSEipAddressRequest()
    }
    
    response = NewAssociateDDoSEipAddressResponse()
    err = c.Send(request, response)
    return
}

// AssociateDDoSEipAddress
// This API is used to bind an EIP to an Anti-DDoS Advanced instance or a specified private IP of an ENI.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) AssociateDDoSEipAddressWithContext(ctx context.Context, request *AssociateDDoSEipAddressRequest) (response *AssociateDDoSEipAddressResponse, err error) {
    if request == nil {
        request = NewAssociateDDoSEipAddressRequest()
    }
    request.SetContext(ctx)
    
    response = NewAssociateDDoSEipAddressResponse()
    err = c.Send(request, response)
    return
}

func NewAssociateDDoSEipLoadBalancerRequest() (request *AssociateDDoSEipLoadBalancerRequest) {
    request = &AssociateDDoSEipLoadBalancerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "AssociateDDoSEipLoadBalancer")
    
    
    return
}

func NewAssociateDDoSEipLoadBalancerResponse() (response *AssociateDDoSEipLoadBalancerResponse) {
    response = &AssociateDDoSEipLoadBalancerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AssociateDDoSEipLoadBalancer
// This API is used to bind an Anti-DDoS EIP to the specified private IP of a CLB instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) AssociateDDoSEipLoadBalancer(request *AssociateDDoSEipLoadBalancerRequest) (response *AssociateDDoSEipLoadBalancerResponse, err error) {
    if request == nil {
        request = NewAssociateDDoSEipLoadBalancerRequest()
    }
    
    response = NewAssociateDDoSEipLoadBalancerResponse()
    err = c.Send(request, response)
    return
}

// AssociateDDoSEipLoadBalancer
// This API is used to bind an Anti-DDoS EIP to the specified private IP of a CLB instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) AssociateDDoSEipLoadBalancerWithContext(ctx context.Context, request *AssociateDDoSEipLoadBalancerRequest) (response *AssociateDDoSEipLoadBalancerResponse, err error) {
    if request == nil {
        request = NewAssociateDDoSEipLoadBalancerRequest()
    }
    request.SetContext(ctx)
    
    response = NewAssociateDDoSEipLoadBalancerResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBlackWhiteIpListRequest() (request *CreateBlackWhiteIpListRequest) {
    request = &CreateBlackWhiteIpListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreateBlackWhiteIpList")
    
    
    return
}

func NewCreateBlackWhiteIpListResponse() (response *CreateBlackWhiteIpListResponse) {
    response = &CreateBlackWhiteIpListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBlackWhiteIpList
// This API is used to add an Anti-DDoS IP blocklist/allowlist.
//
// error code that may be returned:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateBlackWhiteIpList(request *CreateBlackWhiteIpListRequest) (response *CreateBlackWhiteIpListResponse, err error) {
    if request == nil {
        request = NewCreateBlackWhiteIpListRequest()
    }
    
    response = NewCreateBlackWhiteIpListResponse()
    err = c.Send(request, response)
    return
}

// CreateBlackWhiteIpList
// This API is used to add an Anti-DDoS IP blocklist/allowlist.
//
// error code that may be returned:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateBlackWhiteIpListWithContext(ctx context.Context, request *CreateBlackWhiteIpListRequest) (response *CreateBlackWhiteIpListResponse, err error) {
    if request == nil {
        request = NewCreateBlackWhiteIpListRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateBlackWhiteIpListResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBoundIPRequest() (request *CreateBoundIPRequest) {
    request = &CreateBoundIPRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreateBoundIP")
    
    
    return
}

func NewCreateBoundIPResponse() (response *CreateBoundIPResponse) {
    response = &CreateBoundIPResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBoundIP
// This API is used to bind an IP to an Anti-DDoS Pro instance Both single IP instances and multi-IP instances are available. Note that you should wait until the current binding or unbinding completes before using this async API for new operations.
//
// error code that may be returned:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateBoundIP(request *CreateBoundIPRequest) (response *CreateBoundIPResponse, err error) {
    if request == nil {
        request = NewCreateBoundIPRequest()
    }
    
    response = NewCreateBoundIPResponse()
    err = c.Send(request, response)
    return
}

// CreateBoundIP
// This API is used to bind an IP to an Anti-DDoS Pro instance Both single IP instances and multi-IP instances are available. Note that you should wait until the current binding or unbinding completes before using this async API for new operations.
//
// error code that may be returned:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateBoundIPWithContext(ctx context.Context, request *CreateBoundIPRequest) (response *CreateBoundIPResponse, err error) {
    if request == nil {
        request = NewCreateBoundIPRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateBoundIPResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDDoSAIRequest() (request *CreateDDoSAIRequest) {
    request = &CreateDDoSAIRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreateDDoSAI")
    
    
    return
}

func NewCreateDDoSAIResponse() (response *CreateDDoSAIResponse) {
    response = &CreateDDoSAIResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDDoSAI
// This API is used to set Anti-DDoS AI protection switches.
//
// error code that may be returned:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDDoSAI(request *CreateDDoSAIRequest) (response *CreateDDoSAIResponse, err error) {
    if request == nil {
        request = NewCreateDDoSAIRequest()
    }
    
    response = NewCreateDDoSAIResponse()
    err = c.Send(request, response)
    return
}

// CreateDDoSAI
// This API is used to set Anti-DDoS AI protection switches.
//
// error code that may be returned:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDDoSAIWithContext(ctx context.Context, request *CreateDDoSAIRequest) (response *CreateDDoSAIResponse, err error) {
    if request == nil {
        request = NewCreateDDoSAIRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateDDoSAIResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDDoSGeoIPBlockConfigRequest() (request *CreateDDoSGeoIPBlockConfigRequest) {
    request = &CreateDDoSGeoIPBlockConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreateDDoSGeoIPBlockConfig")
    
    
    return
}

func NewCreateDDoSGeoIPBlockConfigResponse() (response *CreateDDoSGeoIPBlockConfigResponse) {
    response = &CreateDDoSGeoIPBlockConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDDoSGeoIPBlockConfig
// This API is used to add an Anti-DDoS region blocking configuration.
//
// error code that may be returned:
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDDoSGeoIPBlockConfig(request *CreateDDoSGeoIPBlockConfigRequest) (response *CreateDDoSGeoIPBlockConfigResponse, err error) {
    if request == nil {
        request = NewCreateDDoSGeoIPBlockConfigRequest()
    }
    
    response = NewCreateDDoSGeoIPBlockConfigResponse()
    err = c.Send(request, response)
    return
}

// CreateDDoSGeoIPBlockConfig
// This API is used to add an Anti-DDoS region blocking configuration.
//
// error code that may be returned:
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDDoSGeoIPBlockConfigWithContext(ctx context.Context, request *CreateDDoSGeoIPBlockConfigRequest) (response *CreateDDoSGeoIPBlockConfigResponse, err error) {
    if request == nil {
        request = NewCreateDDoSGeoIPBlockConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateDDoSGeoIPBlockConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDDoSSpeedLimitConfigRequest() (request *CreateDDoSSpeedLimitConfigRequest) {
    request = &CreateDDoSSpeedLimitConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreateDDoSSpeedLimitConfig")
    
    
    return
}

func NewCreateDDoSSpeedLimitConfigResponse() (response *CreateDDoSSpeedLimitConfigResponse) {
    response = &CreateDDoSSpeedLimitConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDDoSSpeedLimitConfig
// This API is used to add Anti-DDoS access rate limit configurations.
//
// error code that may be returned:
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) CreateDDoSSpeedLimitConfig(request *CreateDDoSSpeedLimitConfigRequest) (response *CreateDDoSSpeedLimitConfigResponse, err error) {
    if request == nil {
        request = NewCreateDDoSSpeedLimitConfigRequest()
    }
    
    response = NewCreateDDoSSpeedLimitConfigResponse()
    err = c.Send(request, response)
    return
}

// CreateDDoSSpeedLimitConfig
// This API is used to add Anti-DDoS access rate limit configurations.
//
// error code that may be returned:
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) CreateDDoSSpeedLimitConfigWithContext(ctx context.Context, request *CreateDDoSSpeedLimitConfigRequest) (response *CreateDDoSSpeedLimitConfigResponse, err error) {
    if request == nil {
        request = NewCreateDDoSSpeedLimitConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateDDoSSpeedLimitConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDefaultAlarmThresholdRequest() (request *CreateDefaultAlarmThresholdRequest) {
    request = &CreateDefaultAlarmThresholdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreateDefaultAlarmThreshold")
    
    
    return
}

func NewCreateDefaultAlarmThresholdResponse() (response *CreateDefaultAlarmThresholdResponse) {
    response = &CreateDefaultAlarmThresholdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDefaultAlarmThreshold
// This API is used to set the default alarm threshold of an IP.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) CreateDefaultAlarmThreshold(request *CreateDefaultAlarmThresholdRequest) (response *CreateDefaultAlarmThresholdResponse, err error) {
    if request == nil {
        request = NewCreateDefaultAlarmThresholdRequest()
    }
    
    response = NewCreateDefaultAlarmThresholdResponse()
    err = c.Send(request, response)
    return
}

// CreateDefaultAlarmThreshold
// This API is used to set the default alarm threshold of an IP.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) CreateDefaultAlarmThresholdWithContext(ctx context.Context, request *CreateDefaultAlarmThresholdRequest) (response *CreateDefaultAlarmThresholdResponse, err error) {
    if request == nil {
        request = NewCreateDefaultAlarmThresholdRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateDefaultAlarmThresholdResponse()
    err = c.Send(request, response)
    return
}

func NewCreateIPAlarmThresholdConfigRequest() (request *CreateIPAlarmThresholdConfigRequest) {
    request = &CreateIPAlarmThresholdConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreateIPAlarmThresholdConfig")
    
    
    return
}

func NewCreateIPAlarmThresholdConfigResponse() (response *CreateIPAlarmThresholdConfigResponse) {
    response = &CreateIPAlarmThresholdConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateIPAlarmThresholdConfig
// This API is used to set the default alarm threshold of an IP.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) CreateIPAlarmThresholdConfig(request *CreateIPAlarmThresholdConfigRequest) (response *CreateIPAlarmThresholdConfigResponse, err error) {
    if request == nil {
        request = NewCreateIPAlarmThresholdConfigRequest()
    }
    
    response = NewCreateIPAlarmThresholdConfigResponse()
    err = c.Send(request, response)
    return
}

// CreateIPAlarmThresholdConfig
// This API is used to set the default alarm threshold of an IP.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) CreateIPAlarmThresholdConfigWithContext(ctx context.Context, request *CreateIPAlarmThresholdConfigRequest) (response *CreateIPAlarmThresholdConfigResponse, err error) {
    if request == nil {
        request = NewCreateIPAlarmThresholdConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateIPAlarmThresholdConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateL7RuleCertsRequest() (request *CreateL7RuleCertsRequest) {
    request = &CreateL7RuleCertsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreateL7RuleCerts")
    
    
    return
}

func NewCreateL7RuleCertsResponse() (response *CreateL7RuleCertsResponse) {
    response = &CreateL7RuleCertsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateL7RuleCerts
// This API is used to configure certificates with layer-7 forwarding rules in a batch for SSL testing.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateL7RuleCerts(request *CreateL7RuleCertsRequest) (response *CreateL7RuleCertsResponse, err error) {
    if request == nil {
        request = NewCreateL7RuleCertsRequest()
    }
    
    response = NewCreateL7RuleCertsResponse()
    err = c.Send(request, response)
    return
}

// CreateL7RuleCerts
// This API is used to configure certificates with layer-7 forwarding rules in a batch for SSL testing.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateL7RuleCertsWithContext(ctx context.Context, request *CreateL7RuleCertsRequest) (response *CreateL7RuleCertsResponse, err error) {
    if request == nil {
        request = NewCreateL7RuleCertsRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateL7RuleCertsResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePacketFilterConfigRequest() (request *CreatePacketFilterConfigRequest) {
    request = &CreatePacketFilterConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreatePacketFilterConfig")
    
    
    return
}

func NewCreatePacketFilterConfigResponse() (response *CreatePacketFilterConfigResponse) {
    response = &CreatePacketFilterConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePacketFilterConfig
// This API is used to add Anti-DDoS feature filtering rules.
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreatePacketFilterConfig(request *CreatePacketFilterConfigRequest) (response *CreatePacketFilterConfigResponse, err error) {
    if request == nil {
        request = NewCreatePacketFilterConfigRequest()
    }
    
    response = NewCreatePacketFilterConfigResponse()
    err = c.Send(request, response)
    return
}

// CreatePacketFilterConfig
// This API is used to add Anti-DDoS feature filtering rules.
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreatePacketFilterConfigWithContext(ctx context.Context, request *CreatePacketFilterConfigRequest) (response *CreatePacketFilterConfigResponse, err error) {
    if request == nil {
        request = NewCreatePacketFilterConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreatePacketFilterConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProtocolBlockConfigRequest() (request *CreateProtocolBlockConfigRequest) {
    request = &CreateProtocolBlockConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreateProtocolBlockConfig")
    
    
    return
}

func NewCreateProtocolBlockConfigResponse() (response *CreateProtocolBlockConfigResponse) {
    response = &CreateProtocolBlockConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateProtocolBlockConfig
// This API is used to set Anti-DDoS protocol blocking configurations.
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateProtocolBlockConfig(request *CreateProtocolBlockConfigRequest) (response *CreateProtocolBlockConfigResponse, err error) {
    if request == nil {
        request = NewCreateProtocolBlockConfigRequest()
    }
    
    response = NewCreateProtocolBlockConfigResponse()
    err = c.Send(request, response)
    return
}

// CreateProtocolBlockConfig
// This API is used to set Anti-DDoS protocol blocking configurations.
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateProtocolBlockConfigWithContext(ctx context.Context, request *CreateProtocolBlockConfigRequest) (response *CreateProtocolBlockConfigResponse, err error) {
    if request == nil {
        request = NewCreateProtocolBlockConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateProtocolBlockConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSchedulingDomainRequest() (request *CreateSchedulingDomainRequest) {
    request = &CreateSchedulingDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreateSchedulingDomain")
    
    
    return
}

func NewCreateSchedulingDomainResponse() (response *CreateSchedulingDomainResponse) {
    response = &CreateSchedulingDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSchedulingDomain
// This API is used to create a domain name for IP scheduling and switching.
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
func (c *Client) CreateSchedulingDomain(request *CreateSchedulingDomainRequest) (response *CreateSchedulingDomainResponse, err error) {
    if request == nil {
        request = NewCreateSchedulingDomainRequest()
    }
    
    response = NewCreateSchedulingDomainResponse()
    err = c.Send(request, response)
    return
}

// CreateSchedulingDomain
// This API is used to create a domain name for IP scheduling and switching.
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
func (c *Client) CreateSchedulingDomainWithContext(ctx context.Context, request *CreateSchedulingDomainRequest) (response *CreateSchedulingDomainResponse, err error) {
    if request == nil {
        request = NewCreateSchedulingDomainRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateSchedulingDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWaterPrintConfigRequest() (request *CreateWaterPrintConfigRequest) {
    request = &CreateWaterPrintConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreateWaterPrintConfig")
    
    
    return
}

func NewCreateWaterPrintConfigResponse() (response *CreateWaterPrintConfigResponse) {
    response = &CreateWaterPrintConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateWaterPrintConfig
// This API is used to add Anti-DDoS watermark configurations.
//
// error code that may be returned:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateWaterPrintConfig(request *CreateWaterPrintConfigRequest) (response *CreateWaterPrintConfigResponse, err error) {
    if request == nil {
        request = NewCreateWaterPrintConfigRequest()
    }
    
    response = NewCreateWaterPrintConfigResponse()
    err = c.Send(request, response)
    return
}

// CreateWaterPrintConfig
// This API is used to add Anti-DDoS watermark configurations.
//
// error code that may be returned:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateWaterPrintConfigWithContext(ctx context.Context, request *CreateWaterPrintConfigRequest) (response *CreateWaterPrintConfigResponse, err error) {
    if request == nil {
        request = NewCreateWaterPrintConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateWaterPrintConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWaterPrintKeyRequest() (request *CreateWaterPrintKeyRequest) {
    request = &CreateWaterPrintKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreateWaterPrintKey")
    
    
    return
}

func NewCreateWaterPrintKeyResponse() (response *CreateWaterPrintKeyResponse) {
    response = &CreateWaterPrintKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateWaterPrintKey
// This API is used to add Anti-DDoS watermark keys.
//
// error code that may be returned:
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) CreateWaterPrintKey(request *CreateWaterPrintKeyRequest) (response *CreateWaterPrintKeyResponse, err error) {
    if request == nil {
        request = NewCreateWaterPrintKeyRequest()
    }
    
    response = NewCreateWaterPrintKeyResponse()
    err = c.Send(request, response)
    return
}

// CreateWaterPrintKey
// This API is used to add Anti-DDoS watermark keys.
//
// error code that may be returned:
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) CreateWaterPrintKeyWithContext(ctx context.Context, request *CreateWaterPrintKeyRequest) (response *CreateWaterPrintKeyResponse, err error) {
    if request == nil {
        request = NewCreateWaterPrintKeyRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateWaterPrintKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBlackWhiteIpListRequest() (request *DeleteBlackWhiteIpListRequest) {
    request = &DeleteBlackWhiteIpListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DeleteBlackWhiteIpList")
    
    
    return
}

func NewDeleteBlackWhiteIpListResponse() (response *DeleteBlackWhiteIpListResponse) {
    response = &DeleteBlackWhiteIpListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteBlackWhiteIpList
// This API is used to delete an Anti-DDoS IP blocklist/allowlist.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteBlackWhiteIpList(request *DeleteBlackWhiteIpListRequest) (response *DeleteBlackWhiteIpListResponse, err error) {
    if request == nil {
        request = NewDeleteBlackWhiteIpListRequest()
    }
    
    response = NewDeleteBlackWhiteIpListResponse()
    err = c.Send(request, response)
    return
}

// DeleteBlackWhiteIpList
// This API is used to delete an Anti-DDoS IP blocklist/allowlist.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteBlackWhiteIpListWithContext(ctx context.Context, request *DeleteBlackWhiteIpListRequest) (response *DeleteBlackWhiteIpListResponse, err error) {
    if request == nil {
        request = NewDeleteBlackWhiteIpListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteBlackWhiteIpListResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDDoSGeoIPBlockConfigRequest() (request *DeleteDDoSGeoIPBlockConfigRequest) {
    request = &DeleteDDoSGeoIPBlockConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DeleteDDoSGeoIPBlockConfig")
    
    
    return
}

func NewDeleteDDoSGeoIPBlockConfigResponse() (response *DeleteDDoSGeoIPBlockConfigResponse) {
    response = &DeleteDDoSGeoIPBlockConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDDoSGeoIPBlockConfig
// This API is used to delete Anti-DDoS region blocking configurations.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteDDoSGeoIPBlockConfig(request *DeleteDDoSGeoIPBlockConfigRequest) (response *DeleteDDoSGeoIPBlockConfigResponse, err error) {
    if request == nil {
        request = NewDeleteDDoSGeoIPBlockConfigRequest()
    }
    
    response = NewDeleteDDoSGeoIPBlockConfigResponse()
    err = c.Send(request, response)
    return
}

// DeleteDDoSGeoIPBlockConfig
// This API is used to delete Anti-DDoS region blocking configurations.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteDDoSGeoIPBlockConfigWithContext(ctx context.Context, request *DeleteDDoSGeoIPBlockConfigRequest) (response *DeleteDDoSGeoIPBlockConfigResponse, err error) {
    if request == nil {
        request = NewDeleteDDoSGeoIPBlockConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteDDoSGeoIPBlockConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDDoSSpeedLimitConfigRequest() (request *DeleteDDoSSpeedLimitConfigRequest) {
    request = &DeleteDDoSSpeedLimitConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DeleteDDoSSpeedLimitConfig")
    
    
    return
}

func NewDeleteDDoSSpeedLimitConfigResponse() (response *DeleteDDoSSpeedLimitConfigResponse) {
    response = &DeleteDDoSSpeedLimitConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDDoSSpeedLimitConfig
// This API is used to delete Anti-DDoS access rate limit configurations.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteDDoSSpeedLimitConfig(request *DeleteDDoSSpeedLimitConfigRequest) (response *DeleteDDoSSpeedLimitConfigResponse, err error) {
    if request == nil {
        request = NewDeleteDDoSSpeedLimitConfigRequest()
    }
    
    response = NewDeleteDDoSSpeedLimitConfigResponse()
    err = c.Send(request, response)
    return
}

// DeleteDDoSSpeedLimitConfig
// This API is used to delete Anti-DDoS access rate limit configurations.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteDDoSSpeedLimitConfigWithContext(ctx context.Context, request *DeleteDDoSSpeedLimitConfigRequest) (response *DeleteDDoSSpeedLimitConfigResponse, err error) {
    if request == nil {
        request = NewDeleteDDoSSpeedLimitConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteDDoSSpeedLimitConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePacketFilterConfigRequest() (request *DeletePacketFilterConfigRequest) {
    request = &DeletePacketFilterConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DeletePacketFilterConfig")
    
    
    return
}

func NewDeletePacketFilterConfigResponse() (response *DeletePacketFilterConfigResponse) {
    response = &DeletePacketFilterConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeletePacketFilterConfig
// This API is used to delete Anti-DDoS feature filtering rules.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeletePacketFilterConfig(request *DeletePacketFilterConfigRequest) (response *DeletePacketFilterConfigResponse, err error) {
    if request == nil {
        request = NewDeletePacketFilterConfigRequest()
    }
    
    response = NewDeletePacketFilterConfigResponse()
    err = c.Send(request, response)
    return
}

// DeletePacketFilterConfig
// This API is used to delete Anti-DDoS feature filtering rules.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeletePacketFilterConfigWithContext(ctx context.Context, request *DeletePacketFilterConfigRequest) (response *DeletePacketFilterConfigResponse, err error) {
    if request == nil {
        request = NewDeletePacketFilterConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeletePacketFilterConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWaterPrintConfigRequest() (request *DeleteWaterPrintConfigRequest) {
    request = &DeleteWaterPrintConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DeleteWaterPrintConfig")
    
    
    return
}

func NewDeleteWaterPrintConfigResponse() (response *DeleteWaterPrintConfigResponse) {
    response = &DeleteWaterPrintConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteWaterPrintConfig
// This API is used to delete Anti-DDoS watermark configurations.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DeleteWaterPrintConfig(request *DeleteWaterPrintConfigRequest) (response *DeleteWaterPrintConfigResponse, err error) {
    if request == nil {
        request = NewDeleteWaterPrintConfigRequest()
    }
    
    response = NewDeleteWaterPrintConfigResponse()
    err = c.Send(request, response)
    return
}

// DeleteWaterPrintConfig
// This API is used to delete Anti-DDoS watermark configurations.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DeleteWaterPrintConfigWithContext(ctx context.Context, request *DeleteWaterPrintConfigRequest) (response *DeleteWaterPrintConfigResponse, err error) {
    if request == nil {
        request = NewDeleteWaterPrintConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteWaterPrintConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWaterPrintKeyRequest() (request *DeleteWaterPrintKeyRequest) {
    request = &DeleteWaterPrintKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DeleteWaterPrintKey")
    
    
    return
}

func NewDeleteWaterPrintKeyResponse() (response *DeleteWaterPrintKeyResponse) {
    response = &DeleteWaterPrintKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteWaterPrintKey
// This API is used to delete Anti-DDoS watermark keys.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DeleteWaterPrintKey(request *DeleteWaterPrintKeyRequest) (response *DeleteWaterPrintKeyResponse, err error) {
    if request == nil {
        request = NewDeleteWaterPrintKeyRequest()
    }
    
    response = NewDeleteWaterPrintKeyResponse()
    err = c.Send(request, response)
    return
}

// DeleteWaterPrintKey
// This API is used to delete Anti-DDoS watermark keys.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DeleteWaterPrintKeyWithContext(ctx context.Context, request *DeleteWaterPrintKeyRequest) (response *DeleteWaterPrintKeyResponse, err error) {
    if request == nil {
        request = NewDeleteWaterPrintKeyRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteWaterPrintKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBasicDeviceStatusRequest() (request *DescribeBasicDeviceStatusRequest) {
    request = &DescribeBasicDeviceStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeBasicDeviceStatus")
    
    
    return
}

func NewDescribeBasicDeviceStatusResponse() (response *DescribeBasicDeviceStatusResponse) {
    response = &DescribeBasicDeviceStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBasicDeviceStatus
// This API is used to querying the status of Anti-DDoS IP.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeBasicDeviceStatus(request *DescribeBasicDeviceStatusRequest) (response *DescribeBasicDeviceStatusResponse, err error) {
    if request == nil {
        request = NewDescribeBasicDeviceStatusRequest()
    }
    
    response = NewDescribeBasicDeviceStatusResponse()
    err = c.Send(request, response)
    return
}

// DescribeBasicDeviceStatus
// This API is used to querying the status of Anti-DDoS IP.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeBasicDeviceStatusWithContext(ctx context.Context, request *DescribeBasicDeviceStatusRequest) (response *DescribeBasicDeviceStatusResponse, err error) {
    if request == nil {
        request = NewDescribeBasicDeviceStatusRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeBasicDeviceStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBizTrendRequest() (request *DescribeBizTrendRequest) {
    request = &DescribeBizTrendRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeBizTrend")
    
    
    return
}

func NewDescribeBizTrendResponse() (response *DescribeBizTrendResponse) {
    response = &DescribeBizTrendResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBizTrend
// This API is used to get the traffic flow data collected in the specified period.
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
func (c *Client) DescribeBizTrend(request *DescribeBizTrendRequest) (response *DescribeBizTrendResponse, err error) {
    if request == nil {
        request = NewDescribeBizTrendRequest()
    }
    
    response = NewDescribeBizTrendResponse()
    err = c.Send(request, response)
    return
}

// DescribeBizTrend
// This API is used to get the traffic flow data collected in the specified period.
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
func (c *Client) DescribeBizTrendWithContext(ctx context.Context, request *DescribeBizTrendRequest) (response *DescribeBizTrendResponse, err error) {
    if request == nil {
        request = NewDescribeBizTrendRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeBizTrendResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBlackWhiteIpListRequest() (request *DescribeBlackWhiteIpListRequest) {
    request = &DescribeBlackWhiteIpListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeBlackWhiteIpList")
    
    
    return
}

func NewDescribeBlackWhiteIpListResponse() (response *DescribeBlackWhiteIpListResponse) {
    response = &DescribeBlackWhiteIpListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBlackWhiteIpList
// This API is used to get an Anti-DDoS IP blocklist/allowlist.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBlackWhiteIpList(request *DescribeBlackWhiteIpListRequest) (response *DescribeBlackWhiteIpListResponse, err error) {
    if request == nil {
        request = NewDescribeBlackWhiteIpListRequest()
    }
    
    response = NewDescribeBlackWhiteIpListResponse()
    err = c.Send(request, response)
    return
}

// DescribeBlackWhiteIpList
// This API is used to get an Anti-DDoS IP blocklist/allowlist.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBlackWhiteIpListWithContext(ctx context.Context, request *DescribeBlackWhiteIpListRequest) (response *DescribeBlackWhiteIpListResponse, err error) {
    if request == nil {
        request = NewDescribeBlackWhiteIpListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeBlackWhiteIpListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCCTrendRequest() (request *DescribeCCTrendRequest) {
    request = &DescribeCCTrendRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeCCTrend")
    
    
    return
}

func NewDescribeCCTrendResponse() (response *DescribeCCTrendResponse) {
    response = &DescribeCCTrendResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCCTrend
// This API is used to get CC attack data, including total QPS peaks, attack QPS, total number of requests and number of attack requests.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCCTrend(request *DescribeCCTrendRequest) (response *DescribeCCTrendResponse, err error) {
    if request == nil {
        request = NewDescribeCCTrendRequest()
    }
    
    response = NewDescribeCCTrendResponse()
    err = c.Send(request, response)
    return
}

// DescribeCCTrend
// This API is used to get CC attack data, including total QPS peaks, attack QPS, total number of requests and number of attack requests.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCCTrendWithContext(ctx context.Context, request *DescribeCCTrendRequest) (response *DescribeCCTrendResponse, err error) {
    if request == nil {
        request = NewDescribeCCTrendRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeCCTrendResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSTrendRequest() (request *DescribeDDoSTrendRequest) {
    request = &DescribeDDoSTrendRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeDDoSTrend")
    
    
    return
}

func NewDescribeDDoSTrendResponse() (response *DescribeDDoSTrendResponse) {
    response = &DescribeDDoSTrendResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDDoSTrend
// This API is used to get DDoS attack traffic bandwidth and attack packet rate.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDDoSTrend(request *DescribeDDoSTrendRequest) (response *DescribeDDoSTrendResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSTrendRequest()
    }
    
    response = NewDescribeDDoSTrendResponse()
    err = c.Send(request, response)
    return
}

// DescribeDDoSTrend
// This API is used to get DDoS attack traffic bandwidth and attack packet rate.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDDoSTrendWithContext(ctx context.Context, request *DescribeDDoSTrendRequest) (response *DescribeDDoSTrendResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSTrendRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDDoSTrendResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDefaultAlarmThresholdRequest() (request *DescribeDefaultAlarmThresholdRequest) {
    request = &DescribeDefaultAlarmThresholdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeDefaultAlarmThreshold")
    
    
    return
}

func NewDescribeDefaultAlarmThresholdResponse() (response *DescribeDefaultAlarmThresholdResponse) {
    response = &DescribeDefaultAlarmThresholdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDefaultAlarmThreshold
// This API is used to get the default alarm threshold of an IP.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDefaultAlarmThreshold(request *DescribeDefaultAlarmThresholdRequest) (response *DescribeDefaultAlarmThresholdResponse, err error) {
    if request == nil {
        request = NewDescribeDefaultAlarmThresholdRequest()
    }
    
    response = NewDescribeDefaultAlarmThresholdResponse()
    err = c.Send(request, response)
    return
}

// DescribeDefaultAlarmThreshold
// This API is used to get the default alarm threshold of an IP.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDefaultAlarmThresholdWithContext(ctx context.Context, request *DescribeDefaultAlarmThresholdRequest) (response *DescribeDefaultAlarmThresholdResponse, err error) {
    if request == nil {
        request = NewDescribeDefaultAlarmThresholdRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDefaultAlarmThresholdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeL7RulesBySSLCertIdRequest() (request *DescribeL7RulesBySSLCertIdRequest) {
    request = &DescribeL7RulesBySSLCertIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeL7RulesBySSLCertId")
    
    
    return
}

func NewDescribeL7RulesBySSLCertIdResponse() (response *DescribeL7RulesBySSLCertIdResponse) {
    response = &DescribeL7RulesBySSLCertIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeL7RulesBySSLCertId
// This API is used to query layer-7 rules matched with certificate IDs.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeL7RulesBySSLCertId(request *DescribeL7RulesBySSLCertIdRequest) (response *DescribeL7RulesBySSLCertIdResponse, err error) {
    if request == nil {
        request = NewDescribeL7RulesBySSLCertIdRequest()
    }
    
    response = NewDescribeL7RulesBySSLCertIdResponse()
    err = c.Send(request, response)
    return
}

// DescribeL7RulesBySSLCertId
// This API is used to query layer-7 rules matched with certificate IDs.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeL7RulesBySSLCertIdWithContext(ctx context.Context, request *DescribeL7RulesBySSLCertIdRequest) (response *DescribeL7RulesBySSLCertIdResponse, err error) {
    if request == nil {
        request = NewDescribeL7RulesBySSLCertIdRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeL7RulesBySSLCertIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListBGPIPInstancesRequest() (request *DescribeListBGPIPInstancesRequest) {
    request = &DescribeListBGPIPInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeListBGPIPInstances")
    
    
    return
}

func NewDescribeListBGPIPInstancesResponse() (response *DescribeListBGPIPInstancesResponse) {
    response = &DescribeListBGPIPInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListBGPIPInstances
// This API is used to get a list of Anti-DDoS Advanced instances.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListBGPIPInstances(request *DescribeListBGPIPInstancesRequest) (response *DescribeListBGPIPInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeListBGPIPInstancesRequest()
    }
    
    response = NewDescribeListBGPIPInstancesResponse()
    err = c.Send(request, response)
    return
}

// DescribeListBGPIPInstances
// This API is used to get a list of Anti-DDoS Advanced instances.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListBGPIPInstancesWithContext(ctx context.Context, request *DescribeListBGPIPInstancesRequest) (response *DescribeListBGPIPInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeListBGPIPInstancesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeListBGPIPInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListBGPInstancesRequest() (request *DescribeListBGPInstancesRequest) {
    request = &DescribeListBGPInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeListBGPInstances")
    
    
    return
}

func NewDescribeListBGPInstancesResponse() (response *DescribeListBGPInstancesResponse) {
    response = &DescribeListBGPInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListBGPInstances
// This API is used to get a list of Anti-DDoS Pro instances.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListBGPInstances(request *DescribeListBGPInstancesRequest) (response *DescribeListBGPInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeListBGPInstancesRequest()
    }
    
    response = NewDescribeListBGPInstancesResponse()
    err = c.Send(request, response)
    return
}

// DescribeListBGPInstances
// This API is used to get a list of Anti-DDoS Pro instances.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListBGPInstancesWithContext(ctx context.Context, request *DescribeListBGPInstancesRequest) (response *DescribeListBGPInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeListBGPInstancesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeListBGPInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListBlackWhiteIpListRequest() (request *DescribeListBlackWhiteIpListRequest) {
    request = &DescribeListBlackWhiteIpListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeListBlackWhiteIpList")
    
    
    return
}

func NewDescribeListBlackWhiteIpListResponse() (response *DescribeListBlackWhiteIpListResponse) {
    response = &DescribeListBlackWhiteIpListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListBlackWhiteIpList
// This API is used to get a list of Anti-DDoS IP blocklists/allowlists.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListBlackWhiteIpList(request *DescribeListBlackWhiteIpListRequest) (response *DescribeListBlackWhiteIpListResponse, err error) {
    if request == nil {
        request = NewDescribeListBlackWhiteIpListRequest()
    }
    
    response = NewDescribeListBlackWhiteIpListResponse()
    err = c.Send(request, response)
    return
}

// DescribeListBlackWhiteIpList
// This API is used to get a list of Anti-DDoS IP blocklists/allowlists.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListBlackWhiteIpListWithContext(ctx context.Context, request *DescribeListBlackWhiteIpListRequest) (response *DescribeListBlackWhiteIpListResponse, err error) {
    if request == nil {
        request = NewDescribeListBlackWhiteIpListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeListBlackWhiteIpListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListDDoSAIRequest() (request *DescribeListDDoSAIRequest) {
    request = &DescribeListDDoSAIRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeListDDoSAI")
    
    
    return
}

func NewDescribeListDDoSAIResponse() (response *DescribeListDDoSAIResponse) {
    response = &DescribeListDDoSAIResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListDDoSAI
// This API is used to get a list of Anti-DDoS AI protection switches.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListDDoSAI(request *DescribeListDDoSAIRequest) (response *DescribeListDDoSAIResponse, err error) {
    if request == nil {
        request = NewDescribeListDDoSAIRequest()
    }
    
    response = NewDescribeListDDoSAIResponse()
    err = c.Send(request, response)
    return
}

// DescribeListDDoSAI
// This API is used to get a list of Anti-DDoS AI protection switches.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListDDoSAIWithContext(ctx context.Context, request *DescribeListDDoSAIRequest) (response *DescribeListDDoSAIResponse, err error) {
    if request == nil {
        request = NewDescribeListDDoSAIRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeListDDoSAIResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListDDoSGeoIPBlockConfigRequest() (request *DescribeListDDoSGeoIPBlockConfigRequest) {
    request = &DescribeListDDoSGeoIPBlockConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeListDDoSGeoIPBlockConfig")
    
    
    return
}

func NewDescribeListDDoSGeoIPBlockConfigResponse() (response *DescribeListDDoSGeoIPBlockConfigResponse) {
    response = &DescribeListDDoSGeoIPBlockConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListDDoSGeoIPBlockConfig
// This API is used to get a list of Anti-DDoS region blocking configurations.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListDDoSGeoIPBlockConfig(request *DescribeListDDoSGeoIPBlockConfigRequest) (response *DescribeListDDoSGeoIPBlockConfigResponse, err error) {
    if request == nil {
        request = NewDescribeListDDoSGeoIPBlockConfigRequest()
    }
    
    response = NewDescribeListDDoSGeoIPBlockConfigResponse()
    err = c.Send(request, response)
    return
}

// DescribeListDDoSGeoIPBlockConfig
// This API is used to get a list of Anti-DDoS region blocking configurations.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListDDoSGeoIPBlockConfigWithContext(ctx context.Context, request *DescribeListDDoSGeoIPBlockConfigRequest) (response *DescribeListDDoSGeoIPBlockConfigResponse, err error) {
    if request == nil {
        request = NewDescribeListDDoSGeoIPBlockConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeListDDoSGeoIPBlockConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListDDoSSpeedLimitConfigRequest() (request *DescribeListDDoSSpeedLimitConfigRequest) {
    request = &DescribeListDDoSSpeedLimitConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeListDDoSSpeedLimitConfig")
    
    
    return
}

func NewDescribeListDDoSSpeedLimitConfigResponse() (response *DescribeListDDoSSpeedLimitConfigResponse) {
    response = &DescribeListDDoSSpeedLimitConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListDDoSSpeedLimitConfig
// This API is used to get a list of Anti-DDoS access rate limit configurations.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListDDoSSpeedLimitConfig(request *DescribeListDDoSSpeedLimitConfigRequest) (response *DescribeListDDoSSpeedLimitConfigResponse, err error) {
    if request == nil {
        request = NewDescribeListDDoSSpeedLimitConfigRequest()
    }
    
    response = NewDescribeListDDoSSpeedLimitConfigResponse()
    err = c.Send(request, response)
    return
}

// DescribeListDDoSSpeedLimitConfig
// This API is used to get a list of Anti-DDoS access rate limit configurations.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListDDoSSpeedLimitConfigWithContext(ctx context.Context, request *DescribeListDDoSSpeedLimitConfigRequest) (response *DescribeListDDoSSpeedLimitConfigResponse, err error) {
    if request == nil {
        request = NewDescribeListDDoSSpeedLimitConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeListDDoSSpeedLimitConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListIPAlarmConfigRequest() (request *DescribeListIPAlarmConfigRequest) {
    request = &DescribeListIPAlarmConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeListIPAlarmConfig")
    
    
    return
}

func NewDescribeListIPAlarmConfigResponse() (response *DescribeListIPAlarmConfigResponse) {
    response = &DescribeListIPAlarmConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListIPAlarmConfig
// This API is used to get a list of IP alarm threshold configurations.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListIPAlarmConfig(request *DescribeListIPAlarmConfigRequest) (response *DescribeListIPAlarmConfigResponse, err error) {
    if request == nil {
        request = NewDescribeListIPAlarmConfigRequest()
    }
    
    response = NewDescribeListIPAlarmConfigResponse()
    err = c.Send(request, response)
    return
}

// DescribeListIPAlarmConfig
// This API is used to get a list of IP alarm threshold configurations.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListIPAlarmConfigWithContext(ctx context.Context, request *DescribeListIPAlarmConfigRequest) (response *DescribeListIPAlarmConfigResponse, err error) {
    if request == nil {
        request = NewDescribeListIPAlarmConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeListIPAlarmConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListListenerRequest() (request *DescribeListListenerRequest) {
    request = &DescribeListListenerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeListListener")
    
    
    return
}

func NewDescribeListListenerResponse() (response *DescribeListListenerResponse) {
    response = &DescribeListListenerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListListener
// This API is used to get a list of forwarding listeners.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListListener(request *DescribeListListenerRequest) (response *DescribeListListenerResponse, err error) {
    if request == nil {
        request = NewDescribeListListenerRequest()
    }
    
    response = NewDescribeListListenerResponse()
    err = c.Send(request, response)
    return
}

// DescribeListListener
// This API is used to get a list of forwarding listeners.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListListenerWithContext(ctx context.Context, request *DescribeListListenerRequest) (response *DescribeListListenerResponse, err error) {
    if request == nil {
        request = NewDescribeListListenerRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeListListenerResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListPacketFilterConfigRequest() (request *DescribeListPacketFilterConfigRequest) {
    request = &DescribeListPacketFilterConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeListPacketFilterConfig")
    
    
    return
}

func NewDescribeListPacketFilterConfigResponse() (response *DescribeListPacketFilterConfigResponse) {
    response = &DescribeListPacketFilterConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListPacketFilterConfig
// This API is used to get a list of Anti-DDoS feature filtering rules.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListPacketFilterConfig(request *DescribeListPacketFilterConfigRequest) (response *DescribeListPacketFilterConfigResponse, err error) {
    if request == nil {
        request = NewDescribeListPacketFilterConfigRequest()
    }
    
    response = NewDescribeListPacketFilterConfigResponse()
    err = c.Send(request, response)
    return
}

// DescribeListPacketFilterConfig
// This API is used to get a list of Anti-DDoS feature filtering rules.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListPacketFilterConfigWithContext(ctx context.Context, request *DescribeListPacketFilterConfigRequest) (response *DescribeListPacketFilterConfigResponse, err error) {
    if request == nil {
        request = NewDescribeListPacketFilterConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeListPacketFilterConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListProtectThresholdConfigRequest() (request *DescribeListProtectThresholdConfigRequest) {
    request = &DescribeListProtectThresholdConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeListProtectThresholdConfig")
    
    
    return
}

func NewDescribeListProtectThresholdConfigResponse() (response *DescribeListProtectThresholdConfigResponse) {
    response = &DescribeListProtectThresholdConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListProtectThresholdConfig
// This API is used to get a list of protection threshold configurations for AI protection switch, protection level, and CC threshold switch.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListProtectThresholdConfig(request *DescribeListProtectThresholdConfigRequest) (response *DescribeListProtectThresholdConfigResponse, err error) {
    if request == nil {
        request = NewDescribeListProtectThresholdConfigRequest()
    }
    
    response = NewDescribeListProtectThresholdConfigResponse()
    err = c.Send(request, response)
    return
}

// DescribeListProtectThresholdConfig
// This API is used to get a list of protection threshold configurations for AI protection switch, protection level, and CC threshold switch.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListProtectThresholdConfigWithContext(ctx context.Context, request *DescribeListProtectThresholdConfigRequest) (response *DescribeListProtectThresholdConfigResponse, err error) {
    if request == nil {
        request = NewDescribeListProtectThresholdConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeListProtectThresholdConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListProtocolBlockConfigRequest() (request *DescribeListProtocolBlockConfigRequest) {
    request = &DescribeListProtocolBlockConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeListProtocolBlockConfig")
    
    
    return
}

func NewDescribeListProtocolBlockConfigResponse() (response *DescribeListProtocolBlockConfigResponse) {
    response = &DescribeListProtocolBlockConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListProtocolBlockConfig
// This API is used to get a list of Anti-DDoS protocol blocking configurations.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListProtocolBlockConfig(request *DescribeListProtocolBlockConfigRequest) (response *DescribeListProtocolBlockConfigResponse, err error) {
    if request == nil {
        request = NewDescribeListProtocolBlockConfigRequest()
    }
    
    response = NewDescribeListProtocolBlockConfigResponse()
    err = c.Send(request, response)
    return
}

// DescribeListProtocolBlockConfig
// This API is used to get a list of Anti-DDoS protocol blocking configurations.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListProtocolBlockConfigWithContext(ctx context.Context, request *DescribeListProtocolBlockConfigRequest) (response *DescribeListProtocolBlockConfigResponse, err error) {
    if request == nil {
        request = NewDescribeListProtocolBlockConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeListProtocolBlockConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListSchedulingDomainRequest() (request *DescribeListSchedulingDomainRequest) {
    request = &DescribeListSchedulingDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeListSchedulingDomain")
    
    
    return
}

func NewDescribeListSchedulingDomainResponse() (response *DescribeListSchedulingDomainResponse) {
    response = &DescribeListSchedulingDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListSchedulingDomain
// This API is used to get a list of intelligent scheduling domain names.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListSchedulingDomain(request *DescribeListSchedulingDomainRequest) (response *DescribeListSchedulingDomainResponse, err error) {
    if request == nil {
        request = NewDescribeListSchedulingDomainRequest()
    }
    
    response = NewDescribeListSchedulingDomainResponse()
    err = c.Send(request, response)
    return
}

// DescribeListSchedulingDomain
// This API is used to get a list of intelligent scheduling domain names.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListSchedulingDomainWithContext(ctx context.Context, request *DescribeListSchedulingDomainRequest) (response *DescribeListSchedulingDomainResponse, err error) {
    if request == nil {
        request = NewDescribeListSchedulingDomainRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeListSchedulingDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListWaterPrintConfigRequest() (request *DescribeListWaterPrintConfigRequest) {
    request = &DescribeListWaterPrintConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeListWaterPrintConfig")
    
    
    return
}

func NewDescribeListWaterPrintConfigResponse() (response *DescribeListWaterPrintConfigResponse) {
    response = &DescribeListWaterPrintConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListWaterPrintConfig
// This API is used to get a list of Anti-DDoS watermark configurations.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListWaterPrintConfig(request *DescribeListWaterPrintConfigRequest) (response *DescribeListWaterPrintConfigResponse, err error) {
    if request == nil {
        request = NewDescribeListWaterPrintConfigRequest()
    }
    
    response = NewDescribeListWaterPrintConfigResponse()
    err = c.Send(request, response)
    return
}

// DescribeListWaterPrintConfig
// This API is used to get a list of Anti-DDoS watermark configurations.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListWaterPrintConfigWithContext(ctx context.Context, request *DescribeListWaterPrintConfigRequest) (response *DescribeListWaterPrintConfigResponse, err error) {
    if request == nil {
        request = NewDescribeListWaterPrintConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeListWaterPrintConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateDDoSEipAddressRequest() (request *DisassociateDDoSEipAddressRequest) {
    request = &DisassociateDDoSEipAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DisassociateDDoSEipAddress")
    
    
    return
}

func NewDisassociateDDoSEipAddressResponse() (response *DisassociateDDoSEipAddressResponse) {
    response = &DisassociateDDoSEipAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisassociateDDoSEipAddress
// This API is used to unbind an Anti-DDoS EIP.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DisassociateDDoSEipAddress(request *DisassociateDDoSEipAddressRequest) (response *DisassociateDDoSEipAddressResponse, err error) {
    if request == nil {
        request = NewDisassociateDDoSEipAddressRequest()
    }
    
    response = NewDisassociateDDoSEipAddressResponse()
    err = c.Send(request, response)
    return
}

// DisassociateDDoSEipAddress
// This API is used to unbind an Anti-DDoS EIP.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DisassociateDDoSEipAddressWithContext(ctx context.Context, request *DisassociateDDoSEipAddressRequest) (response *DisassociateDDoSEipAddressResponse, err error) {
    if request == nil {
        request = NewDisassociateDDoSEipAddressRequest()
    }
    request.SetContext(ctx)
    
    response = NewDisassociateDDoSEipAddressResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDDoSGeoIPBlockConfigRequest() (request *ModifyDDoSGeoIPBlockConfigRequest) {
    request = &ModifyDDoSGeoIPBlockConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "ModifyDDoSGeoIPBlockConfig")
    
    
    return
}

func NewModifyDDoSGeoIPBlockConfigResponse() (response *ModifyDDoSGeoIPBlockConfigResponse) {
    response = &ModifyDDoSGeoIPBlockConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDDoSGeoIPBlockConfig
// This API is used to modify Anti-DDoS region blocking configurations.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyDDoSGeoIPBlockConfig(request *ModifyDDoSGeoIPBlockConfigRequest) (response *ModifyDDoSGeoIPBlockConfigResponse, err error) {
    if request == nil {
        request = NewModifyDDoSGeoIPBlockConfigRequest()
    }
    
    response = NewModifyDDoSGeoIPBlockConfigResponse()
    err = c.Send(request, response)
    return
}

// ModifyDDoSGeoIPBlockConfig
// This API is used to modify Anti-DDoS region blocking configurations.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyDDoSGeoIPBlockConfigWithContext(ctx context.Context, request *ModifyDDoSGeoIPBlockConfigRequest) (response *ModifyDDoSGeoIPBlockConfigResponse, err error) {
    if request == nil {
        request = NewModifyDDoSGeoIPBlockConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyDDoSGeoIPBlockConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDDoSSpeedLimitConfigRequest() (request *ModifyDDoSSpeedLimitConfigRequest) {
    request = &ModifyDDoSSpeedLimitConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "ModifyDDoSSpeedLimitConfig")
    
    
    return
}

func NewModifyDDoSSpeedLimitConfigResponse() (response *ModifyDDoSSpeedLimitConfigResponse) {
    response = &ModifyDDoSSpeedLimitConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDDoSSpeedLimitConfig
// This API is used to modify Anti-DDoS access rate limit configurations.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyDDoSSpeedLimitConfig(request *ModifyDDoSSpeedLimitConfigRequest) (response *ModifyDDoSSpeedLimitConfigResponse, err error) {
    if request == nil {
        request = NewModifyDDoSSpeedLimitConfigRequest()
    }
    
    response = NewModifyDDoSSpeedLimitConfigResponse()
    err = c.Send(request, response)
    return
}

// ModifyDDoSSpeedLimitConfig
// This API is used to modify Anti-DDoS access rate limit configurations.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyDDoSSpeedLimitConfigWithContext(ctx context.Context, request *ModifyDDoSSpeedLimitConfigRequest) (response *ModifyDDoSSpeedLimitConfigResponse, err error) {
    if request == nil {
        request = NewModifyDDoSSpeedLimitConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyDDoSSpeedLimitConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDomainUsrNameRequest() (request *ModifyDomainUsrNameRequest) {
    request = &ModifyDomainUsrNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "ModifyDomainUsrName")
    
    
    return
}

func NewModifyDomainUsrNameResponse() (response *ModifyDomainUsrNameResponse) {
    response = &ModifyDomainUsrNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDomainUsrName
// This API is used to modify intelligent scheduling domain names.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyDomainUsrName(request *ModifyDomainUsrNameRequest) (response *ModifyDomainUsrNameResponse, err error) {
    if request == nil {
        request = NewModifyDomainUsrNameRequest()
    }
    
    response = NewModifyDomainUsrNameResponse()
    err = c.Send(request, response)
    return
}

// ModifyDomainUsrName
// This API is used to modify intelligent scheduling domain names.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyDomainUsrNameWithContext(ctx context.Context, request *ModifyDomainUsrNameRequest) (response *ModifyDomainUsrNameResponse, err error) {
    if request == nil {
        request = NewModifyDomainUsrNameRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyDomainUsrNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNewDomainRulesRequest() (request *ModifyNewDomainRulesRequest) {
    request = &ModifyNewDomainRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "ModifyNewDomainRules")
    
    
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyNewDomainRules(request *ModifyNewDomainRulesRequest) (response *ModifyNewDomainRulesResponse, err error) {
    if request == nil {
        request = NewModifyNewDomainRulesRequest()
    }
    
    response = NewModifyNewDomainRulesResponse()
    err = c.Send(request, response)
    return
}

// ModifyNewDomainRules
// This API is used to modify layer-7 forwarding rules.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyNewDomainRulesWithContext(ctx context.Context, request *ModifyNewDomainRulesRequest) (response *ModifyNewDomainRulesResponse, err error) {
    if request == nil {
        request = NewModifyNewDomainRulesRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyNewDomainRulesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPacketFilterConfigRequest() (request *ModifyPacketFilterConfigRequest) {
    request = &ModifyPacketFilterConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "ModifyPacketFilterConfig")
    
    
    return
}

func NewModifyPacketFilterConfigResponse() (response *ModifyPacketFilterConfigResponse) {
    response = &ModifyPacketFilterConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyPacketFilterConfig
// This API is used to modify Anti-DDoS feature filtering rules.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyPacketFilterConfig(request *ModifyPacketFilterConfigRequest) (response *ModifyPacketFilterConfigResponse, err error) {
    if request == nil {
        request = NewModifyPacketFilterConfigRequest()
    }
    
    response = NewModifyPacketFilterConfigResponse()
    err = c.Send(request, response)
    return
}

// ModifyPacketFilterConfig
// This API is used to modify Anti-DDoS feature filtering rules.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyPacketFilterConfigWithContext(ctx context.Context, request *ModifyPacketFilterConfigRequest) (response *ModifyPacketFilterConfigResponse, err error) {
    if request == nil {
        request = NewModifyPacketFilterConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyPacketFilterConfigResponse()
    err = c.Send(request, response)
    return
}

func NewSwitchWaterPrintConfigRequest() (request *SwitchWaterPrintConfigRequest) {
    request = &SwitchWaterPrintConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "SwitchWaterPrintConfig")
    
    
    return
}

func NewSwitchWaterPrintConfigResponse() (response *SwitchWaterPrintConfigResponse) {
    response = &SwitchWaterPrintConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SwitchWaterPrintConfig
// This API is used to enable or disable Anti-DDoS watermark configurations.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SwitchWaterPrintConfig(request *SwitchWaterPrintConfigRequest) (response *SwitchWaterPrintConfigResponse, err error) {
    if request == nil {
        request = NewSwitchWaterPrintConfigRequest()
    }
    
    response = NewSwitchWaterPrintConfigResponse()
    err = c.Send(request, response)
    return
}

// SwitchWaterPrintConfig
// This API is used to enable or disable Anti-DDoS watermark configurations.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SwitchWaterPrintConfigWithContext(ctx context.Context, request *SwitchWaterPrintConfigRequest) (response *SwitchWaterPrintConfigResponse, err error) {
    if request == nil {
        request = NewSwitchWaterPrintConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewSwitchWaterPrintConfigResponse()
    err = c.Send(request, response)
    return
}
