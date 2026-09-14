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

package v20250115

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2025-01-15"

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


func NewCreateAccelerateAreasRequest() (request *CreateAccelerateAreasRequest) {
    request = &CreateAccelerateAreasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "CreateAccelerateAreas")
    
    
    return
}

func NewCreateAccelerateAreasResponse() (response *CreateAccelerateAreasResponse) {
    response = &CreateAccelerateAreasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAccelerateAreas
// This API is used to create an acceleration region.
//
// error code that may be returned:
//  INVALIDPARAMETER_INPUTREPEAT = "InvalidParameter.InputRepeat"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_THIRDAREALISTENERPORTSEGMENT = "InvalidParameterValue.ThirdAreaListenerPortSegment"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION_ACCELERATEREGIONREPEAT = "UnsupportedOperation.AccelerateRegionRepeat"
//  UNSUPPORTEDOPERATION_EXISTACCEPTACLPOLICY = "UnsupportedOperation.ExistAcceptAclPolicy"
//  UNSUPPORTEDOPERATION_EXISTDROPACCEPTACLRULE = "UnsupportedOperation.ExistDropAcceptAclRule"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTALLOWEDOPERATE = "UnsupportedOperation.InstanceStateNotAllowedOperate"
//  UNSUPPORTEDOPERATION_NOTSETCROSSBORDERTYPE = "UnsupportedOperation.NotSetCrossBorderType"
//  UNSUPPORTEDOPERATION_THIRDAREAHAVEPORTEQUALLISTENER = "UnsupportedOperation.ThirdAreaHavePortEqualListener"
//  UNSUPPORTEDOPERATION_THREENETWORKSACCELERATEAREAS = "UnsupportedOperation.ThreeNetworksAccelerateAreas"
//  UNSUPPORTEDOPERATION_UNABLECREATETHIRDPARTYNODES = "UnsupportedOperation.UnableCreateThirdPartyNodes"
func (c *Client) CreateAccelerateAreas(request *CreateAccelerateAreasRequest) (response *CreateAccelerateAreasResponse, err error) {
    return c.CreateAccelerateAreasWithContext(context.Background(), request)
}

// CreateAccelerateAreas
// This API is used to create an acceleration region.
//
// error code that may be returned:
//  INVALIDPARAMETER_INPUTREPEAT = "InvalidParameter.InputRepeat"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_THIRDAREALISTENERPORTSEGMENT = "InvalidParameterValue.ThirdAreaListenerPortSegment"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION_ACCELERATEREGIONREPEAT = "UnsupportedOperation.AccelerateRegionRepeat"
//  UNSUPPORTEDOPERATION_EXISTACCEPTACLPOLICY = "UnsupportedOperation.ExistAcceptAclPolicy"
//  UNSUPPORTEDOPERATION_EXISTDROPACCEPTACLRULE = "UnsupportedOperation.ExistDropAcceptAclRule"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTALLOWEDOPERATE = "UnsupportedOperation.InstanceStateNotAllowedOperate"
//  UNSUPPORTEDOPERATION_NOTSETCROSSBORDERTYPE = "UnsupportedOperation.NotSetCrossBorderType"
//  UNSUPPORTEDOPERATION_THIRDAREAHAVEPORTEQUALLISTENER = "UnsupportedOperation.ThirdAreaHavePortEqualListener"
//  UNSUPPORTEDOPERATION_THREENETWORKSACCELERATEAREAS = "UnsupportedOperation.ThreeNetworksAccelerateAreas"
//  UNSUPPORTEDOPERATION_UNABLECREATETHIRDPARTYNODES = "UnsupportedOperation.UnableCreateThirdPartyNodes"
func (c *Client) CreateAccelerateAreasWithContext(ctx context.Context, request *CreateAccelerateAreasRequest) (response *CreateAccelerateAreasResponse, err error) {
    if request == nil {
        request = NewCreateAccelerateAreasRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "CreateAccelerateAreas")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAccelerateAreas require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAccelerateAreasResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEndpointGroupRequest() (request *CreateEndpointGroupRequest) {
    request = &CreateEndpointGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "CreateEndpointGroup")
    
    
    return
}

func NewCreateEndpointGroupResponse() (response *CreateEndpointGroupResponse) {
    response = &CreateEndpointGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateEndpointGroup
// This API is used to create a terminal node group.
//
// error code that may be returned:
//  INVALIDPARAMETER_DISABLEHEALTHCHECKNOTCARRYPARAMETERS = "InvalidParameter.DisableHealthCheckNotCarryParameters"
//  INVALIDPARAMETER_ENABLEHEALTHCHECKNOTCARRYPARAMETERS = "InvalidParameter.EnableHealthCheckNotCarryParameters"
//  INVALIDPARAMETER_ENDPOINTGROUPCHECKTYPEHTTPCARRYPARAMETERS = "InvalidParameter.EndpointGroupCheckTypeHttpCarryParameters"
//  INVALIDPARAMETER_HTTPSORIGINSERVERPROTOCOLSUPPORTSPARAMETERS = "InvalidParameter.HttpsOriginServerProtocolSupportsParameters"
//  INVALIDPARAMETER_INPUTOUTOFRANGE = "InvalidParameter.InputOutOfRange"
//  INVALIDPARAMETER_INPUTREPEAT = "InvalidParameter.InputRepeat"
//  INVALIDPARAMETER_ONLYENDPOINTGROUPCHECKTYPEHTTPCARRYPARAMETERS = "InvalidParameter.OnlyEndpointGroupCheckTypeHttpCarryParameters"
//  INVALIDPARAMETER_TCPENDPOINTGROUPCANNOTCARRYPARAMETERS = "InvalidParameter.TcpEndpointGroupCannotCarryParameters"
//  INVALIDPARAMETER_TRANSPORTLAYERENDPOINTGROUPCANNOTCARRYPARAMETERS = "InvalidParameter.TransportLayerEndpointGroupCannotCarryParameters"
//  INVALIDPARAMETER_UDPENDPOINTGROUPCANNOTCARRYPARAMETERS = "InvalidParameter.UdpEndpointGroupCannotCarryParameters"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_LISTENERPORTNOTEQUAL = "InvalidParameterValue.ListenerPortNotEqual"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NOTMAPPINGMANYPORT = "InvalidParameterValue.NotMappingManyPort"
//  INVALIDPARAMETERVALUE_NOTTENCENTREGION = "InvalidParameterValue.NotTencentRegion"
//  INVALIDPARAMETERVALUE_NOTWITHINPUBLICNETWORK = "InvalidParameterValue.NotWithinPublicNetwork"
//  INVALIDPARAMETERVALUE_SPECIFICCHARACTERS = "InvalidParameterValue.SpecificCharacters"
//  INVALIDPARAMETERVALUE_THIRDAREALISTENERPORTSEGMENT = "InvalidParameterValue.ThirdAreaListenerPortSegment"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_WRONGATTRIBUTIONRELATIONSHIP = "InvalidParameterValue.WrongAttributionRelationship"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_APPLICATIONLAYERENDPOINTGROUPPARAMETER = "MissingParameter.ApplicationLayerEndpointGroupParameter"
//  MISSINGPARAMETER_CUSTOMCHECKTYPEPARAMETER = "MissingParameter.CustomCheckTypeParameter"
//  MISSINGPARAMETER_ENABLEHEALTHCHECKPARAMETER = "MissingParameter.EnableHealthCheckParameter"
//  MISSINGPARAMETER_ISPTYPE = "MissingParameter.IspType"
//  UNSUPPORTEDOPERATION_CURRENTACCOUNTNOTALLOWED = "UnsupportedOperation.CurrentAccountNotAllowed"
//  UNSUPPORTEDOPERATION_EXISTACCEPTACLPOLICY = "UnsupportedOperation.ExistAcceptAclPolicy"
//  UNSUPPORTEDOPERATION_EXISTDROPACCEPTACLRULE = "UnsupportedOperation.ExistDropAcceptAclRule"
//  UNSUPPORTEDOPERATION_HTTPVERSIONINCONSISTENT = "UnsupportedOperation.HttpVersionInconsistent"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
//  UNSUPPORTEDOPERATION_INTERNALENDPOINTFEATURENOTENABLED = "UnsupportedOperation.InternalEndpointFeatureNotEnabled"
//  UNSUPPORTEDOPERATION_IPV6NOTSUPPORT = "UnsupportedOperation.Ipv6NotSupport"
//  UNSUPPORTEDOPERATION_NOTSETCROSSBORDERTYPE = "UnsupportedOperation.NotSetCrossBorderType"
//  UNSUPPORTEDOPERATION_ONLYINTERNALCLB = "UnsupportedOperation.OnlyInternalClb"
//  UNSUPPORTEDOPERATION_SAMEDOMAIN = "UnsupportedOperation.SameDomain"
//  UNSUPPORTEDOPERATION_SAMEPUBLICIP = "UnsupportedOperation.SamePublicIp"
//  UNSUPPORTEDOPERATION_THIRDAREAHAVEPORTEQUALLISTENER = "UnsupportedOperation.ThirdAreaHavePortEqualListener"
//  UNSUPPORTEDOPERATION_THREENETWORKSENDPOINTGROUP = "UnsupportedOperation.ThreeNetworksEndpointGroup"
//  UNSUPPORTEDOPERATION_UNABLECREATETHIRDPARTYNODES = "UnsupportedOperation.UnableCreateThirdPartyNodes"
//  UNSUPPORTEDOPERATION_UNABLECROSSBORDER = "UnsupportedOperation.UnableCrossBorder"
//  UNSUPPORTEDOPERATION_VIRTUALENDPOINTGROUPUNSUPPORTEDTCPANDUDP = "UnsupportedOperation.VirtualEndpointGroupUnsupportedTcpAndUdp"
func (c *Client) CreateEndpointGroup(request *CreateEndpointGroupRequest) (response *CreateEndpointGroupResponse, err error) {
    return c.CreateEndpointGroupWithContext(context.Background(), request)
}

// CreateEndpointGroup
// This API is used to create a terminal node group.
//
// error code that may be returned:
//  INVALIDPARAMETER_DISABLEHEALTHCHECKNOTCARRYPARAMETERS = "InvalidParameter.DisableHealthCheckNotCarryParameters"
//  INVALIDPARAMETER_ENABLEHEALTHCHECKNOTCARRYPARAMETERS = "InvalidParameter.EnableHealthCheckNotCarryParameters"
//  INVALIDPARAMETER_ENDPOINTGROUPCHECKTYPEHTTPCARRYPARAMETERS = "InvalidParameter.EndpointGroupCheckTypeHttpCarryParameters"
//  INVALIDPARAMETER_HTTPSORIGINSERVERPROTOCOLSUPPORTSPARAMETERS = "InvalidParameter.HttpsOriginServerProtocolSupportsParameters"
//  INVALIDPARAMETER_INPUTOUTOFRANGE = "InvalidParameter.InputOutOfRange"
//  INVALIDPARAMETER_INPUTREPEAT = "InvalidParameter.InputRepeat"
//  INVALIDPARAMETER_ONLYENDPOINTGROUPCHECKTYPEHTTPCARRYPARAMETERS = "InvalidParameter.OnlyEndpointGroupCheckTypeHttpCarryParameters"
//  INVALIDPARAMETER_TCPENDPOINTGROUPCANNOTCARRYPARAMETERS = "InvalidParameter.TcpEndpointGroupCannotCarryParameters"
//  INVALIDPARAMETER_TRANSPORTLAYERENDPOINTGROUPCANNOTCARRYPARAMETERS = "InvalidParameter.TransportLayerEndpointGroupCannotCarryParameters"
//  INVALIDPARAMETER_UDPENDPOINTGROUPCANNOTCARRYPARAMETERS = "InvalidParameter.UdpEndpointGroupCannotCarryParameters"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_LISTENERPORTNOTEQUAL = "InvalidParameterValue.ListenerPortNotEqual"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NOTMAPPINGMANYPORT = "InvalidParameterValue.NotMappingManyPort"
//  INVALIDPARAMETERVALUE_NOTTENCENTREGION = "InvalidParameterValue.NotTencentRegion"
//  INVALIDPARAMETERVALUE_NOTWITHINPUBLICNETWORK = "InvalidParameterValue.NotWithinPublicNetwork"
//  INVALIDPARAMETERVALUE_SPECIFICCHARACTERS = "InvalidParameterValue.SpecificCharacters"
//  INVALIDPARAMETERVALUE_THIRDAREALISTENERPORTSEGMENT = "InvalidParameterValue.ThirdAreaListenerPortSegment"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_WRONGATTRIBUTIONRELATIONSHIP = "InvalidParameterValue.WrongAttributionRelationship"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_APPLICATIONLAYERENDPOINTGROUPPARAMETER = "MissingParameter.ApplicationLayerEndpointGroupParameter"
//  MISSINGPARAMETER_CUSTOMCHECKTYPEPARAMETER = "MissingParameter.CustomCheckTypeParameter"
//  MISSINGPARAMETER_ENABLEHEALTHCHECKPARAMETER = "MissingParameter.EnableHealthCheckParameter"
//  MISSINGPARAMETER_ISPTYPE = "MissingParameter.IspType"
//  UNSUPPORTEDOPERATION_CURRENTACCOUNTNOTALLOWED = "UnsupportedOperation.CurrentAccountNotAllowed"
//  UNSUPPORTEDOPERATION_EXISTACCEPTACLPOLICY = "UnsupportedOperation.ExistAcceptAclPolicy"
//  UNSUPPORTEDOPERATION_EXISTDROPACCEPTACLRULE = "UnsupportedOperation.ExistDropAcceptAclRule"
//  UNSUPPORTEDOPERATION_HTTPVERSIONINCONSISTENT = "UnsupportedOperation.HttpVersionInconsistent"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
//  UNSUPPORTEDOPERATION_INTERNALENDPOINTFEATURENOTENABLED = "UnsupportedOperation.InternalEndpointFeatureNotEnabled"
//  UNSUPPORTEDOPERATION_IPV6NOTSUPPORT = "UnsupportedOperation.Ipv6NotSupport"
//  UNSUPPORTEDOPERATION_NOTSETCROSSBORDERTYPE = "UnsupportedOperation.NotSetCrossBorderType"
//  UNSUPPORTEDOPERATION_ONLYINTERNALCLB = "UnsupportedOperation.OnlyInternalClb"
//  UNSUPPORTEDOPERATION_SAMEDOMAIN = "UnsupportedOperation.SameDomain"
//  UNSUPPORTEDOPERATION_SAMEPUBLICIP = "UnsupportedOperation.SamePublicIp"
//  UNSUPPORTEDOPERATION_THIRDAREAHAVEPORTEQUALLISTENER = "UnsupportedOperation.ThirdAreaHavePortEqualListener"
//  UNSUPPORTEDOPERATION_THREENETWORKSENDPOINTGROUP = "UnsupportedOperation.ThreeNetworksEndpointGroup"
//  UNSUPPORTEDOPERATION_UNABLECREATETHIRDPARTYNODES = "UnsupportedOperation.UnableCreateThirdPartyNodes"
//  UNSUPPORTEDOPERATION_UNABLECROSSBORDER = "UnsupportedOperation.UnableCrossBorder"
//  UNSUPPORTEDOPERATION_VIRTUALENDPOINTGROUPUNSUPPORTEDTCPANDUDP = "UnsupportedOperation.VirtualEndpointGroupUnsupportedTcpAndUdp"
func (c *Client) CreateEndpointGroupWithContext(ctx context.Context, request *CreateEndpointGroupRequest) (response *CreateEndpointGroupResponse, err error) {
    if request == nil {
        request = NewCreateEndpointGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "CreateEndpointGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEndpointGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEndpointGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateForwardingPolicyRequest() (request *CreateForwardingPolicyRequest) {
    request = &CreateForwardingPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "CreateForwardingPolicy")
    
    
    return
}

func NewCreateForwardingPolicyResponse() (response *CreateForwardingPolicyResponse) {
    response = &CreateForwardingPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateForwardingPolicy
// Create a layer-7 forwarding policy.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_FORWARDINGPOLICYHOSTCONFLICT = "InvalidParameterValue.ForwardingPolicyHostConflict"
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTALLOWEDOPERATE = "UnsupportedOperation.InstanceStateNotAllowedOperate"
//  UNSUPPORTEDOPERATION_TRANSPORTLAYERUNSUPPORTEDOPERATEFORWARDINGPOLICY = "UnsupportedOperation.TransportLayerUnsupportedOperateForwardingPolicy"
func (c *Client) CreateForwardingPolicy(request *CreateForwardingPolicyRequest) (response *CreateForwardingPolicyResponse, err error) {
    return c.CreateForwardingPolicyWithContext(context.Background(), request)
}

// CreateForwardingPolicy
// Create a layer-7 forwarding policy.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_FORWARDINGPOLICYHOSTCONFLICT = "InvalidParameterValue.ForwardingPolicyHostConflict"
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTALLOWEDOPERATE = "UnsupportedOperation.InstanceStateNotAllowedOperate"
//  UNSUPPORTEDOPERATION_TRANSPORTLAYERUNSUPPORTEDOPERATEFORWARDINGPOLICY = "UnsupportedOperation.TransportLayerUnsupportedOperateForwardingPolicy"
func (c *Client) CreateForwardingPolicyWithContext(ctx context.Context, request *CreateForwardingPolicyRequest) (response *CreateForwardingPolicyResponse, err error) {
    if request == nil {
        request = NewCreateForwardingPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "CreateForwardingPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateForwardingPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateForwardingPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateForwardingRuleRequest() (request *CreateForwardingRuleRequest) {
    request = &CreateForwardingRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "CreateForwardingRule")
    
    
    return
}

func NewCreateForwardingRuleResponse() (response *CreateForwardingRuleResponse) {
    response = &CreateForwardingRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateForwardingRule
// Create a Layer 7 forwarding rule
//
// error code that may be returned:
//  INVALIDPARAMETER_INPUTREPEAT = "InvalidParameter.InputRepeat"
//  INVALIDPARAMETERVALUE_ASCIICHARACTERS = "InvalidParameterValue.AsciiCharacters"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_FORWARDINGRULEPATHCONFLICT = "InvalidParameterValue.ForwardingRulePathConflict"
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  INVALIDPARAMETERVALUE_INTERNALRESERVEDFIELDS = "InvalidParameterValue.InternalReservedFields"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_SPECIALCHARACTERS = "InvalidParameterValue.SpecialCharacters"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_FORWARDGROUPFORWARDINGRULEUNSUPPORTEDDEFAULTENDPOINTGROUP = "UnsupportedOperation.ForwardGroupForwardingRuleUnsupportedDefaultEndpointGroup"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
//  UNSUPPORTEDOPERATION_TRANSPORTLAYERUNSUPPORTEDOPERATEFORWARDINGPOLICY = "UnsupportedOperation.TransportLayerUnsupportedOperateForwardingPolicy"
func (c *Client) CreateForwardingRule(request *CreateForwardingRuleRequest) (response *CreateForwardingRuleResponse, err error) {
    return c.CreateForwardingRuleWithContext(context.Background(), request)
}

// CreateForwardingRule
// Create a Layer 7 forwarding rule
//
// error code that may be returned:
//  INVALIDPARAMETER_INPUTREPEAT = "InvalidParameter.InputRepeat"
//  INVALIDPARAMETERVALUE_ASCIICHARACTERS = "InvalidParameterValue.AsciiCharacters"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_FORWARDINGRULEPATHCONFLICT = "InvalidParameterValue.ForwardingRulePathConflict"
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  INVALIDPARAMETERVALUE_INTERNALRESERVEDFIELDS = "InvalidParameterValue.InternalReservedFields"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_SPECIALCHARACTERS = "InvalidParameterValue.SpecialCharacters"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_FORWARDGROUPFORWARDINGRULEUNSUPPORTEDDEFAULTENDPOINTGROUP = "UnsupportedOperation.ForwardGroupForwardingRuleUnsupportedDefaultEndpointGroup"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
//  UNSUPPORTEDOPERATION_TRANSPORTLAYERUNSUPPORTEDOPERATEFORWARDINGPOLICY = "UnsupportedOperation.TransportLayerUnsupportedOperateForwardingPolicy"
func (c *Client) CreateForwardingRuleWithContext(ctx context.Context, request *CreateForwardingRuleRequest) (response *CreateForwardingRuleResponse, err error) {
    if request == nil {
        request = NewCreateForwardingRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "CreateForwardingRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateForwardingRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateForwardingRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGlobalAcceleratorRequest() (request *CreateGlobalAcceleratorRequest) {
    request = &CreateGlobalAcceleratorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "CreateGlobalAccelerator")
    
    
    return
}

func NewCreateGlobalAcceleratorResponse() (response *CreateGlobalAcceleratorResponse) {
    response = &CreateGlobalAcceleratorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateGlobalAccelerator
// This API is used to create a global acceleration instance.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEGLOBALACCELERATORFAILED = "FailedOperation.CreateGlobalAcceleratorFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENAME = "InvalidParameterValue.InstanceName"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_ACCOUNTFROZEN = "UnsupportedOperation.AccountFrozen"
//  UNSUPPORTEDOPERATION_BILLINGFAILED = "UnsupportedOperation.BillingFailed"
//  UNSUPPORTEDOPERATION_CROSSBORDERPROMISENOTSIGNED = "UnsupportedOperation.CrossBorderPromiseNotSigned"
//  UNSUPPORTEDOPERATION_CURRENTACCOUNTNOTALLOWED = "UnsupportedOperation.CurrentAccountNotAllowed"
//  UNSUPPORTEDOPERATION_ENABLECROSSBORDER = "UnsupportedOperation.EnableCrossBorder"
//  UNSUPPORTEDOPERATION_INSUFFICIENTFUNDS = "UnsupportedOperation.InsufficientFunds"
//  UNSUPPORTEDOPERATION_UNABLECROSSBORDER = "UnsupportedOperation.UnableCrossBorder"
func (c *Client) CreateGlobalAccelerator(request *CreateGlobalAcceleratorRequest) (response *CreateGlobalAcceleratorResponse, err error) {
    return c.CreateGlobalAcceleratorWithContext(context.Background(), request)
}

// CreateGlobalAccelerator
// This API is used to create a global acceleration instance.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEGLOBALACCELERATORFAILED = "FailedOperation.CreateGlobalAcceleratorFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENAME = "InvalidParameterValue.InstanceName"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_ACCOUNTFROZEN = "UnsupportedOperation.AccountFrozen"
//  UNSUPPORTEDOPERATION_BILLINGFAILED = "UnsupportedOperation.BillingFailed"
//  UNSUPPORTEDOPERATION_CROSSBORDERPROMISENOTSIGNED = "UnsupportedOperation.CrossBorderPromiseNotSigned"
//  UNSUPPORTEDOPERATION_CURRENTACCOUNTNOTALLOWED = "UnsupportedOperation.CurrentAccountNotAllowed"
//  UNSUPPORTEDOPERATION_ENABLECROSSBORDER = "UnsupportedOperation.EnableCrossBorder"
//  UNSUPPORTEDOPERATION_INSUFFICIENTFUNDS = "UnsupportedOperation.InsufficientFunds"
//  UNSUPPORTEDOPERATION_UNABLECROSSBORDER = "UnsupportedOperation.UnableCrossBorder"
func (c *Client) CreateGlobalAcceleratorWithContext(ctx context.Context, request *CreateGlobalAcceleratorRequest) (response *CreateGlobalAcceleratorResponse, err error) {
    if request == nil {
        request = NewCreateGlobalAcceleratorRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "CreateGlobalAccelerator")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateGlobalAccelerator require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateGlobalAcceleratorResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGlobalAcceleratorAccessLogRequest() (request *CreateGlobalAcceleratorAccessLogRequest) {
    request = &CreateGlobalAcceleratorAccessLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "CreateGlobalAcceleratorAccessLog")
    
    
    return
}

func NewCreateGlobalAcceleratorAccessLogResponse() (response *CreateGlobalAcceleratorAccessLogResponse) {
    response = &CreateGlobalAcceleratorAccessLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateGlobalAcceleratorAccessLog
// Create a GA access log
//
// error code that may be returned:
//  INVALIDPARAMETER_INPUTREPEAT = "InvalidParameter.InputRepeat"
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  INVALIDPARAMETERVALUE_LOGSETNOTEXIST = "InvalidParameterValue.LogsetNotExist"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_VPGHAGROUPNOTFOUND = "InvalidParameterValue.VpgHaGroupNotFound"
//  UNSUPPORTEDOPERATION_EXISTACCESSLOG = "UnsupportedOperation.ExistAccessLog"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
//  UNSUPPORTEDOPERATION_NOTAUTHORIZATION = "UnsupportedOperation.NotAuthorization"
func (c *Client) CreateGlobalAcceleratorAccessLog(request *CreateGlobalAcceleratorAccessLogRequest) (response *CreateGlobalAcceleratorAccessLogResponse, err error) {
    return c.CreateGlobalAcceleratorAccessLogWithContext(context.Background(), request)
}

// CreateGlobalAcceleratorAccessLog
// Create a GA access log
//
// error code that may be returned:
//  INVALIDPARAMETER_INPUTREPEAT = "InvalidParameter.InputRepeat"
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  INVALIDPARAMETERVALUE_LOGSETNOTEXIST = "InvalidParameterValue.LogsetNotExist"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_VPGHAGROUPNOTFOUND = "InvalidParameterValue.VpgHaGroupNotFound"
//  UNSUPPORTEDOPERATION_EXISTACCESSLOG = "UnsupportedOperation.ExistAccessLog"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
//  UNSUPPORTEDOPERATION_NOTAUTHORIZATION = "UnsupportedOperation.NotAuthorization"
func (c *Client) CreateGlobalAcceleratorAccessLogWithContext(ctx context.Context, request *CreateGlobalAcceleratorAccessLogRequest) (response *CreateGlobalAcceleratorAccessLogResponse, err error) {
    if request == nil {
        request = NewCreateGlobalAcceleratorAccessLogRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "CreateGlobalAcceleratorAccessLog")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateGlobalAcceleratorAccessLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateGlobalAcceleratorAccessLogResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGlobalAcceleratorAclPolicyRequest() (request *CreateGlobalAcceleratorAclPolicyRequest) {
    request = &CreateGlobalAcceleratorAclPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "CreateGlobalAcceleratorAclPolicy")
    
    
    return
}

func NewCreateGlobalAcceleratorAclPolicyResponse() (response *CreateGlobalAcceleratorAclPolicyResponse) {
    response = &CreateGlobalAcceleratorAclPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateGlobalAcceleratorAclPolicy
// Create access control policy
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION_EXISTTHIRDPARTYNODES = "UnsupportedOperation.ExistThirdPartyNodes"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTALLOWEDOPERATE = "UnsupportedOperation.InstanceStateNotAllowedOperate"
func (c *Client) CreateGlobalAcceleratorAclPolicy(request *CreateGlobalAcceleratorAclPolicyRequest) (response *CreateGlobalAcceleratorAclPolicyResponse, err error) {
    return c.CreateGlobalAcceleratorAclPolicyWithContext(context.Background(), request)
}

// CreateGlobalAcceleratorAclPolicy
// Create access control policy
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION_EXISTTHIRDPARTYNODES = "UnsupportedOperation.ExistThirdPartyNodes"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTALLOWEDOPERATE = "UnsupportedOperation.InstanceStateNotAllowedOperate"
func (c *Client) CreateGlobalAcceleratorAclPolicyWithContext(ctx context.Context, request *CreateGlobalAcceleratorAclPolicyRequest) (response *CreateGlobalAcceleratorAclPolicyResponse, err error) {
    if request == nil {
        request = NewCreateGlobalAcceleratorAclPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "CreateGlobalAcceleratorAclPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateGlobalAcceleratorAclPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateGlobalAcceleratorAclPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGlobalAcceleratorAclRuleRequest() (request *CreateGlobalAcceleratorAclRuleRequest) {
    request = &CreateGlobalAcceleratorAclRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "CreateGlobalAcceleratorAclRule")
    
    
    return
}

func NewCreateGlobalAcceleratorAclRuleResponse() (response *CreateGlobalAcceleratorAclRuleResponse) {
    response = &CreateGlobalAcceleratorAclRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateGlobalAcceleratorAclRule
// Create an ACL rule
//
// error code that may be returned:
//  INVALIDPARAMETER_INPUTDUPLICATEWITHEXISTINGDATA = "InvalidParameter.InputDuplicateWithExistingData"
//  INVALIDPARAMETER_INPUTOUTOFRANGE = "InvalidParameter.InputOutOfRange"
//  INVALIDPARAMETER_INPUTREPEAT = "InvalidParameter.InputRepeat"
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTALLOWEDOPERATE = "UnsupportedOperation.InstanceStateNotAllowedOperate"
//  UNSUPPORTEDOPERATION_REQUESTPARAMETERSERROR = "UnsupportedOperation.RequestParametersError"
func (c *Client) CreateGlobalAcceleratorAclRule(request *CreateGlobalAcceleratorAclRuleRequest) (response *CreateGlobalAcceleratorAclRuleResponse, err error) {
    return c.CreateGlobalAcceleratorAclRuleWithContext(context.Background(), request)
}

// CreateGlobalAcceleratorAclRule
// Create an ACL rule
//
// error code that may be returned:
//  INVALIDPARAMETER_INPUTDUPLICATEWITHEXISTINGDATA = "InvalidParameter.InputDuplicateWithExistingData"
//  INVALIDPARAMETER_INPUTOUTOFRANGE = "InvalidParameter.InputOutOfRange"
//  INVALIDPARAMETER_INPUTREPEAT = "InvalidParameter.InputRepeat"
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTALLOWEDOPERATE = "UnsupportedOperation.InstanceStateNotAllowedOperate"
//  UNSUPPORTEDOPERATION_REQUESTPARAMETERSERROR = "UnsupportedOperation.RequestParametersError"
func (c *Client) CreateGlobalAcceleratorAclRuleWithContext(ctx context.Context, request *CreateGlobalAcceleratorAclRuleRequest) (response *CreateGlobalAcceleratorAclRuleResponse, err error) {
    if request == nil {
        request = NewCreateGlobalAcceleratorAclRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "CreateGlobalAcceleratorAclRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateGlobalAcceleratorAclRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateGlobalAcceleratorAclRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateListenerRequest() (request *CreateListenerRequest) {
    request = &CreateListenerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "CreateListener")
    
    
    return
}

func NewCreateListenerResponse() (response *CreateListenerResponse) {
    response = &CreateListenerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateListener
// This API is used to create a listener.
//
// error code that may be returned:
//  FAILEDOPERATION_INSUFFICIENTRESOURCES = "FailedOperation.InsufficientResources"
//  INVALIDPARAMETER_APPLICATIONLAYERLISTENERCANNOTCARRYPARAMETERS = "InvalidParameter.ApplicationLayerListenerCannotCarryParameters"
//  INVALIDPARAMETER_HTTPLISTENERCANNOTCARRYPARAMETERS = "InvalidParameter.HttpListenerCannotCarryParameters"
//  INVALIDPARAMETER_HTTPSLISTENERCANCARRYPARAMETERS = "InvalidParameter.HttpsListenerCanCarryParameters"
//  INVALIDPARAMETER_SINGLEPORT = "InvalidParameter.SinglePort"
//  INVALIDPARAMETER_TCPLISTENERCARRYPARAMETERS = "InvalidParameter.TcpListenerCarryParameters"
//  INVALIDPARAMETER_TRANSPORTLAYERLISTENERCANNOTCARRYPARAMETERS = "InvalidParameter.TransportLayerListenerCannotCarryParameters"
//  INVALIDPARAMETERVALUE_INSTANCENAME = "InvalidParameterValue.InstanceName"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_SEGMENTPORTRANGE = "InvalidParameterValue.SegmentPortRange"
//  INVALIDPARAMETERVALUE_TCPSERIESLISTENERPORTEQUAL = "InvalidParameterValue.TcpSeriesListenerPortEqual"
//  INVALIDPARAMETERVALUE_THIRDAREALISTENERPORTEQUAL = "InvalidParameterValue.ThirdAreaListenerPortEqual"
//  INVALIDPARAMETERVALUE_THIRDAREALISTENERPORTSEGMENT = "InvalidParameterValue.ThirdAreaListenerPortSegment"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_UDPSERIESLISTENERPORTEQUAL = "InvalidParameterValue.UdpSeriesListenerPortEqual"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_CURRENTACCOUNTNOTALLOWED = "UnsupportedOperation.CurrentAccountNotAllowed"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTALLOWEDOPERATE = "UnsupportedOperation.InstanceStateNotAllowedOperate"
//  UNSUPPORTEDOPERATION_LISTENERPORTSEGMENT = "UnsupportedOperation.ListenerPortSegment"
//  UNSUPPORTEDOPERATION_ONLYTCPLISTENERSUPPORTTOA = "UnsupportedOperation.OnlyTcpListenerSupportToa"
//  UNSUPPORTEDOPERATION_TOAFEATURENOTENABLED = "UnsupportedOperation.TOAFeatureNotEnabled"
//  UNSUPPORTEDOPERATION_UDPLISTENERNOTSUPPORTPROXYPROTOCOL = "UnsupportedOperation.UdpListenerNotSupportProxyProtocol"
//  UNSUPPORTEDOPERATION_UNLAWFULCERTIFICATE = "UnsupportedOperation.UnlawfulCertificate"
func (c *Client) CreateListener(request *CreateListenerRequest) (response *CreateListenerResponse, err error) {
    return c.CreateListenerWithContext(context.Background(), request)
}

// CreateListener
// This API is used to create a listener.
//
// error code that may be returned:
//  FAILEDOPERATION_INSUFFICIENTRESOURCES = "FailedOperation.InsufficientResources"
//  INVALIDPARAMETER_APPLICATIONLAYERLISTENERCANNOTCARRYPARAMETERS = "InvalidParameter.ApplicationLayerListenerCannotCarryParameters"
//  INVALIDPARAMETER_HTTPLISTENERCANNOTCARRYPARAMETERS = "InvalidParameter.HttpListenerCannotCarryParameters"
//  INVALIDPARAMETER_HTTPSLISTENERCANCARRYPARAMETERS = "InvalidParameter.HttpsListenerCanCarryParameters"
//  INVALIDPARAMETER_SINGLEPORT = "InvalidParameter.SinglePort"
//  INVALIDPARAMETER_TCPLISTENERCARRYPARAMETERS = "InvalidParameter.TcpListenerCarryParameters"
//  INVALIDPARAMETER_TRANSPORTLAYERLISTENERCANNOTCARRYPARAMETERS = "InvalidParameter.TransportLayerListenerCannotCarryParameters"
//  INVALIDPARAMETERVALUE_INSTANCENAME = "InvalidParameterValue.InstanceName"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_SEGMENTPORTRANGE = "InvalidParameterValue.SegmentPortRange"
//  INVALIDPARAMETERVALUE_TCPSERIESLISTENERPORTEQUAL = "InvalidParameterValue.TcpSeriesListenerPortEqual"
//  INVALIDPARAMETERVALUE_THIRDAREALISTENERPORTEQUAL = "InvalidParameterValue.ThirdAreaListenerPortEqual"
//  INVALIDPARAMETERVALUE_THIRDAREALISTENERPORTSEGMENT = "InvalidParameterValue.ThirdAreaListenerPortSegment"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_UDPSERIESLISTENERPORTEQUAL = "InvalidParameterValue.UdpSeriesListenerPortEqual"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_CURRENTACCOUNTNOTALLOWED = "UnsupportedOperation.CurrentAccountNotAllowed"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTALLOWEDOPERATE = "UnsupportedOperation.InstanceStateNotAllowedOperate"
//  UNSUPPORTEDOPERATION_LISTENERPORTSEGMENT = "UnsupportedOperation.ListenerPortSegment"
//  UNSUPPORTEDOPERATION_ONLYTCPLISTENERSUPPORTTOA = "UnsupportedOperation.OnlyTcpListenerSupportToa"
//  UNSUPPORTEDOPERATION_TOAFEATURENOTENABLED = "UnsupportedOperation.TOAFeatureNotEnabled"
//  UNSUPPORTEDOPERATION_UDPLISTENERNOTSUPPORTPROXYPROTOCOL = "UnsupportedOperation.UdpListenerNotSupportProxyProtocol"
//  UNSUPPORTEDOPERATION_UNLAWFULCERTIFICATE = "UnsupportedOperation.UnlawfulCertificate"
func (c *Client) CreateListenerWithContext(ctx context.Context, request *CreateListenerRequest) (response *CreateListenerResponse, err error) {
    if request == nil {
        request = NewCreateListenerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "CreateListener")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateListener require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateListenerResponse()
    err = c.Send(request, response)
    return
}

func NewCreateListenerAdditionalCertRequest() (request *CreateListenerAdditionalCertRequest) {
    request = &CreateListenerAdditionalCertRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "CreateListenerAdditionalCert")
    
    
    return
}

func NewCreateListenerAdditionalCertResponse() (response *CreateListenerAdditionalCertResponse) {
    response = &CreateListenerAdditionalCertResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateListenerAdditionalCert
// Add an extension certificate.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_CERTIFICATESCONFLICT = "InvalidParameterValue.CertificatesConflict"
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
func (c *Client) CreateListenerAdditionalCert(request *CreateListenerAdditionalCertRequest) (response *CreateListenerAdditionalCertResponse, err error) {
    return c.CreateListenerAdditionalCertWithContext(context.Background(), request)
}

// CreateListenerAdditionalCert
// Add an extension certificate.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_CERTIFICATESCONFLICT = "InvalidParameterValue.CertificatesConflict"
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
func (c *Client) CreateListenerAdditionalCertWithContext(ctx context.Context, request *CreateListenerAdditionalCertRequest) (response *CreateListenerAdditionalCertResponse, err error) {
    if request == nil {
        request = NewCreateListenerAdditionalCertRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "CreateListenerAdditionalCert")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateListenerAdditionalCert require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateListenerAdditionalCertResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAccelerateAreasRequest() (request *DeleteAccelerateAreasRequest) {
    request = &DeleteAccelerateAreasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "DeleteAccelerateAreas")
    
    
    return
}

func NewDeleteAccelerateAreasResponse() (response *DeleteAccelerateAreasResponse) {
    response = &DeleteAccelerateAreasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAccelerateAreas
// Delete an acceleration region
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  RESOURCEINUSE_EDGEACCELERATEAREA = "ResourceInUse.EdgeAccelerateArea"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTALLOWEDOPERATE = "UnsupportedOperation.InstanceStateNotAllowedOperate"
func (c *Client) DeleteAccelerateAreas(request *DeleteAccelerateAreasRequest) (response *DeleteAccelerateAreasResponse, err error) {
    return c.DeleteAccelerateAreasWithContext(context.Background(), request)
}

// DeleteAccelerateAreas
// Delete an acceleration region
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  RESOURCEINUSE_EDGEACCELERATEAREA = "ResourceInUse.EdgeAccelerateArea"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTALLOWEDOPERATE = "UnsupportedOperation.InstanceStateNotAllowedOperate"
func (c *Client) DeleteAccelerateAreasWithContext(ctx context.Context, request *DeleteAccelerateAreasRequest) (response *DeleteAccelerateAreasResponse, err error) {
    if request == nil {
        request = NewDeleteAccelerateAreasRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "DeleteAccelerateAreas")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAccelerateAreas require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAccelerateAreasResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEndpointGroupsRequest() (request *DeleteEndpointGroupsRequest) {
    request = &DeleteEndpointGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "DeleteEndpointGroups")
    
    
    return
}

func NewDeleteEndpointGroupsResponse() (response *DeleteEndpointGroupsResponse) {
    response = &DeleteEndpointGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteEndpointGroups
// Delete a terminal node group.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  UNSUPPORTEDOPERATION_APPLICATIONLAYERENDPOINTGROUPNOTDELETE = "UnsupportedOperation.ApplicationLayerEndpointGroupNotDelete"
//  UNSUPPORTEDOPERATION_EXISTACCESSLOG = "UnsupportedOperation.ExistAccessLog"
//  UNSUPPORTEDOPERATION_EXISTFORWARDINGRULE = "UnsupportedOperation.ExistForwardingRule"
func (c *Client) DeleteEndpointGroups(request *DeleteEndpointGroupsRequest) (response *DeleteEndpointGroupsResponse, err error) {
    return c.DeleteEndpointGroupsWithContext(context.Background(), request)
}

// DeleteEndpointGroups
// Delete a terminal node group.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  UNSUPPORTEDOPERATION_APPLICATIONLAYERENDPOINTGROUPNOTDELETE = "UnsupportedOperation.ApplicationLayerEndpointGroupNotDelete"
//  UNSUPPORTEDOPERATION_EXISTACCESSLOG = "UnsupportedOperation.ExistAccessLog"
//  UNSUPPORTEDOPERATION_EXISTFORWARDINGRULE = "UnsupportedOperation.ExistForwardingRule"
func (c *Client) DeleteEndpointGroupsWithContext(ctx context.Context, request *DeleteEndpointGroupsRequest) (response *DeleteEndpointGroupsResponse, err error) {
    if request == nil {
        request = NewDeleteEndpointGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "DeleteEndpointGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteEndpointGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteEndpointGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteForwardingPolicyRequest() (request *DeleteForwardingPolicyRequest) {
    request = &DeleteForwardingPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "DeleteForwardingPolicy")
    
    
    return
}

func NewDeleteForwardingPolicyResponse() (response *DeleteForwardingPolicyResponse) {
    response = &DeleteForwardingPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteForwardingPolicy
// Delete a layer-7 forwarding policy.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  UNSUPPORTEDOPERATION_DEFAULTFORWARDINGPOLICYOPERATE = "UnsupportedOperation.DefaultForwardingPolicyOperate"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
//  UNSUPPORTEDOPERATION_TRANSPORTLAYERUNSUPPORTEDOPERATEFORWARDINGPOLICY = "UnsupportedOperation.TransportLayerUnsupportedOperateForwardingPolicy"
func (c *Client) DeleteForwardingPolicy(request *DeleteForwardingPolicyRequest) (response *DeleteForwardingPolicyResponse, err error) {
    return c.DeleteForwardingPolicyWithContext(context.Background(), request)
}

// DeleteForwardingPolicy
// Delete a layer-7 forwarding policy.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  UNSUPPORTEDOPERATION_DEFAULTFORWARDINGPOLICYOPERATE = "UnsupportedOperation.DefaultForwardingPolicyOperate"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
//  UNSUPPORTEDOPERATION_TRANSPORTLAYERUNSUPPORTEDOPERATEFORWARDINGPOLICY = "UnsupportedOperation.TransportLayerUnsupportedOperateForwardingPolicy"
func (c *Client) DeleteForwardingPolicyWithContext(ctx context.Context, request *DeleteForwardingPolicyRequest) (response *DeleteForwardingPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteForwardingPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "DeleteForwardingPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteForwardingPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteForwardingPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteForwardingRuleRequest() (request *DeleteForwardingRuleRequest) {
    request = &DeleteForwardingRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "DeleteForwardingRule")
    
    
    return
}

func NewDeleteForwardingRuleResponse() (response *DeleteForwardingRuleResponse) {
    response = &DeleteForwardingRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteForwardingRule
// Delete a Layer 7 forwarding rule
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  UNSUPPORTEDOPERATION_DEFAULTFORWARDINGPOLICYOPERATE = "UnsupportedOperation.DefaultForwardingPolicyOperate"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
//  UNSUPPORTEDOPERATION_TRANSPORTLAYERUNSUPPORTEDOPERATEFORWARDINGPOLICY = "UnsupportedOperation.TransportLayerUnsupportedOperateForwardingPolicy"
func (c *Client) DeleteForwardingRule(request *DeleteForwardingRuleRequest) (response *DeleteForwardingRuleResponse, err error) {
    return c.DeleteForwardingRuleWithContext(context.Background(), request)
}

// DeleteForwardingRule
// Delete a Layer 7 forwarding rule
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  UNSUPPORTEDOPERATION_DEFAULTFORWARDINGPOLICYOPERATE = "UnsupportedOperation.DefaultForwardingPolicyOperate"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
//  UNSUPPORTEDOPERATION_TRANSPORTLAYERUNSUPPORTEDOPERATEFORWARDINGPOLICY = "UnsupportedOperation.TransportLayerUnsupportedOperateForwardingPolicy"
func (c *Client) DeleteForwardingRuleWithContext(ctx context.Context, request *DeleteForwardingRuleRequest) (response *DeleteForwardingRuleResponse, err error) {
    if request == nil {
        request = NewDeleteForwardingRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "DeleteForwardingRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteForwardingRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteForwardingRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGlobalAcceleratorRequest() (request *DeleteGlobalAcceleratorRequest) {
    request = &DeleteGlobalAcceleratorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "DeleteGlobalAccelerator")
    
    
    return
}

func NewDeleteGlobalAcceleratorResponse() (response *DeleteGlobalAcceleratorResponse) {
    response = &DeleteGlobalAcceleratorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteGlobalAccelerator
// Deletes a global acceleration instance
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION_EXISTACCELERATORAREA = "UnsupportedOperation.ExistAcceleratorArea"
//  UNSUPPORTEDOPERATION_EXISTLISTENER = "UnsupportedOperation.ExistListener"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTALLOWEDOPERATE = "UnsupportedOperation.InstanceStateNotAllowedOperate"
func (c *Client) DeleteGlobalAccelerator(request *DeleteGlobalAcceleratorRequest) (response *DeleteGlobalAcceleratorResponse, err error) {
    return c.DeleteGlobalAcceleratorWithContext(context.Background(), request)
}

// DeleteGlobalAccelerator
// Deletes a global acceleration instance
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION_EXISTACCELERATORAREA = "UnsupportedOperation.ExistAcceleratorArea"
//  UNSUPPORTEDOPERATION_EXISTLISTENER = "UnsupportedOperation.ExistListener"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTALLOWEDOPERATE = "UnsupportedOperation.InstanceStateNotAllowedOperate"
func (c *Client) DeleteGlobalAcceleratorWithContext(ctx context.Context, request *DeleteGlobalAcceleratorRequest) (response *DeleteGlobalAcceleratorResponse, err error) {
    if request == nil {
        request = NewDeleteGlobalAcceleratorRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "DeleteGlobalAccelerator")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteGlobalAccelerator require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteGlobalAcceleratorResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGlobalAcceleratorAccessLogRequest() (request *DeleteGlobalAcceleratorAccessLogRequest) {
    request = &DeleteGlobalAcceleratorAccessLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "DeleteGlobalAcceleratorAccessLog")
    
    
    return
}

func NewDeleteGlobalAcceleratorAccessLogResponse() (response *DeleteGlobalAcceleratorAccessLogResponse) {
    response = &DeleteGlobalAcceleratorAccessLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteGlobalAcceleratorAccessLog
// This API is used to delete a GA log task.
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION_EXISTACCELERATORAREA = "UnsupportedOperation.ExistAcceleratorArea"
//  UNSUPPORTEDOPERATION_EXISTLISTENER = "UnsupportedOperation.ExistListener"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTALLOWEDOPERATE = "UnsupportedOperation.InstanceStateNotAllowedOperate"
func (c *Client) DeleteGlobalAcceleratorAccessLog(request *DeleteGlobalAcceleratorAccessLogRequest) (response *DeleteGlobalAcceleratorAccessLogResponse, err error) {
    return c.DeleteGlobalAcceleratorAccessLogWithContext(context.Background(), request)
}

// DeleteGlobalAcceleratorAccessLog
// This API is used to delete a GA log task.
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION_EXISTACCELERATORAREA = "UnsupportedOperation.ExistAcceleratorArea"
//  UNSUPPORTEDOPERATION_EXISTLISTENER = "UnsupportedOperation.ExistListener"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTALLOWEDOPERATE = "UnsupportedOperation.InstanceStateNotAllowedOperate"
func (c *Client) DeleteGlobalAcceleratorAccessLogWithContext(ctx context.Context, request *DeleteGlobalAcceleratorAccessLogRequest) (response *DeleteGlobalAcceleratorAccessLogResponse, err error) {
    if request == nil {
        request = NewDeleteGlobalAcceleratorAccessLogRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "DeleteGlobalAcceleratorAccessLog")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteGlobalAcceleratorAccessLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteGlobalAcceleratorAccessLogResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGlobalAcceleratorAclPolicyRequest() (request *DeleteGlobalAcceleratorAclPolicyRequest) {
    request = &DeleteGlobalAcceleratorAclPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "DeleteGlobalAcceleratorAclPolicy")
    
    
    return
}

func NewDeleteGlobalAcceleratorAclPolicyResponse() (response *DeleteGlobalAcceleratorAclPolicyResponse) {
    response = &DeleteGlobalAcceleratorAclPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteGlobalAcceleratorAclPolicy
// Delete access control policy
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTALLOWEDOPERATE = "UnsupportedOperation.InstanceStateNotAllowedOperate"
func (c *Client) DeleteGlobalAcceleratorAclPolicy(request *DeleteGlobalAcceleratorAclPolicyRequest) (response *DeleteGlobalAcceleratorAclPolicyResponse, err error) {
    return c.DeleteGlobalAcceleratorAclPolicyWithContext(context.Background(), request)
}

// DeleteGlobalAcceleratorAclPolicy
// Delete access control policy
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTALLOWEDOPERATE = "UnsupportedOperation.InstanceStateNotAllowedOperate"
func (c *Client) DeleteGlobalAcceleratorAclPolicyWithContext(ctx context.Context, request *DeleteGlobalAcceleratorAclPolicyRequest) (response *DeleteGlobalAcceleratorAclPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteGlobalAcceleratorAclPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "DeleteGlobalAcceleratorAclPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteGlobalAcceleratorAclPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteGlobalAcceleratorAclPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGlobalAcceleratorAclRuleRequest() (request *DeleteGlobalAcceleratorAclRuleRequest) {
    request = &DeleteGlobalAcceleratorAclRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "DeleteGlobalAcceleratorAclRule")
    
    
    return
}

func NewDeleteGlobalAcceleratorAclRuleResponse() (response *DeleteGlobalAcceleratorAclRuleResponse) {
    response = &DeleteGlobalAcceleratorAclRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteGlobalAcceleratorAclRule
// Delete ACL rule
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTALLOWEDOPERATE = "UnsupportedOperation.InstanceStateNotAllowedOperate"
func (c *Client) DeleteGlobalAcceleratorAclRule(request *DeleteGlobalAcceleratorAclRuleRequest) (response *DeleteGlobalAcceleratorAclRuleResponse, err error) {
    return c.DeleteGlobalAcceleratorAclRuleWithContext(context.Background(), request)
}

// DeleteGlobalAcceleratorAclRule
// Delete ACL rule
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTALLOWEDOPERATE = "UnsupportedOperation.InstanceStateNotAllowedOperate"
func (c *Client) DeleteGlobalAcceleratorAclRuleWithContext(ctx context.Context, request *DeleteGlobalAcceleratorAclRuleRequest) (response *DeleteGlobalAcceleratorAclRuleResponse, err error) {
    if request == nil {
        request = NewDeleteGlobalAcceleratorAclRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "DeleteGlobalAcceleratorAclRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteGlobalAcceleratorAclRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteGlobalAcceleratorAclRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteListenerRequest() (request *DeleteListenerRequest) {
    request = &DeleteListenerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "DeleteListener")
    
    
    return
}

func NewDeleteListenerResponse() (response *DeleteListenerResponse) {
    response = &DeleteListenerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteListener
// This API is used to delete a listener.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  UNSUPPORTEDOPERATION_EXISTACCESSLOG = "UnsupportedOperation.ExistAccessLog"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
func (c *Client) DeleteListener(request *DeleteListenerRequest) (response *DeleteListenerResponse, err error) {
    return c.DeleteListenerWithContext(context.Background(), request)
}

// DeleteListener
// This API is used to delete a listener.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  UNSUPPORTEDOPERATION_EXISTACCESSLOG = "UnsupportedOperation.ExistAccessLog"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
func (c *Client) DeleteListenerWithContext(ctx context.Context, request *DeleteListenerRequest) (response *DeleteListenerResponse, err error) {
    if request == nil {
        request = NewDeleteListenerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "DeleteListener")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteListener require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteListenerResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteListenerAdditionalCertRequest() (request *DeleteListenerAdditionalCertRequest) {
    request = &DeleteListenerAdditionalCertRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "DeleteListenerAdditionalCert")
    
    
    return
}

func NewDeleteListenerAdditionalCertResponse() (response *DeleteListenerAdditionalCertResponse) {
    response = &DeleteListenerAdditionalCertResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteListenerAdditionalCert
// Delete the extension certificate.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  UNSUPPORTEDOPERATION_CACERTIFICATESOPERATE = "UnsupportedOperation.CaCertificatesOperate"
//  UNSUPPORTEDOPERATION_DEFAULTCERTIFICATESOPERATE = "UnsupportedOperation.DefaultCertificatesOperate"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
func (c *Client) DeleteListenerAdditionalCert(request *DeleteListenerAdditionalCertRequest) (response *DeleteListenerAdditionalCertResponse, err error) {
    return c.DeleteListenerAdditionalCertWithContext(context.Background(), request)
}

// DeleteListenerAdditionalCert
// Delete the extension certificate.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  UNSUPPORTEDOPERATION_CACERTIFICATESOPERATE = "UnsupportedOperation.CaCertificatesOperate"
//  UNSUPPORTEDOPERATION_DEFAULTCERTIFICATESOPERATE = "UnsupportedOperation.DefaultCertificatesOperate"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
func (c *Client) DeleteListenerAdditionalCertWithContext(ctx context.Context, request *DeleteListenerAdditionalCertRequest) (response *DeleteListenerAdditionalCertResponse, err error) {
    if request == nil {
        request = NewDeleteListenerAdditionalCertRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "DeleteListenerAdditionalCert")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteListenerAdditionalCert require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteListenerAdditionalCertResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccelerateAreasRequest() (request *DescribeAccelerateAreasRequest) {
    request = &DescribeAccelerateAreasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "DescribeAccelerateAreas")
    
    
    return
}

func NewDescribeAccelerateAreasResponse() (response *DescribeAccelerateAreasResponse) {
    response = &DescribeAccelerateAreasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccelerateAreas
// Queries acceleration regions
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  UNSUPPORTEDOPERATION_CACERTIFICATESOPERATE = "UnsupportedOperation.CaCertificatesOperate"
//  UNSUPPORTEDOPERATION_DEFAULTCERTIFICATESOPERATE = "UnsupportedOperation.DefaultCertificatesOperate"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
func (c *Client) DescribeAccelerateAreas(request *DescribeAccelerateAreasRequest) (response *DescribeAccelerateAreasResponse, err error) {
    return c.DescribeAccelerateAreasWithContext(context.Background(), request)
}

// DescribeAccelerateAreas
// Queries acceleration regions
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  UNSUPPORTEDOPERATION_CACERTIFICATESOPERATE = "UnsupportedOperation.CaCertificatesOperate"
//  UNSUPPORTEDOPERATION_DEFAULTCERTIFICATESOPERATE = "UnsupportedOperation.DefaultCertificatesOperate"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
func (c *Client) DescribeAccelerateAreasWithContext(ctx context.Context, request *DescribeAccelerateAreasRequest) (response *DescribeAccelerateAreasResponse, err error) {
    if request == nil {
        request = NewDescribeAccelerateAreasRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "DescribeAccelerateAreas")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccelerateAreas require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccelerateAreasResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccelerateRegionsRequest() (request *DescribeAccelerateRegionsRequest) {
    request = &DescribeAccelerateRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "DescribeAccelerateRegions")
    
    
    return
}

func NewDescribeAccelerateRegionsResponse() (response *DescribeAccelerateRegionsResponse) {
    response = &DescribeAccelerateRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccelerateRegions
// Queries selectable acceleration regions.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  UNSUPPORTEDOPERATION_CACERTIFICATESOPERATE = "UnsupportedOperation.CaCertificatesOperate"
//  UNSUPPORTEDOPERATION_DEFAULTCERTIFICATESOPERATE = "UnsupportedOperation.DefaultCertificatesOperate"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
func (c *Client) DescribeAccelerateRegions(request *DescribeAccelerateRegionsRequest) (response *DescribeAccelerateRegionsResponse, err error) {
    return c.DescribeAccelerateRegionsWithContext(context.Background(), request)
}

// DescribeAccelerateRegions
// Queries selectable acceleration regions.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  UNSUPPORTEDOPERATION_CACERTIFICATESOPERATE = "UnsupportedOperation.CaCertificatesOperate"
//  UNSUPPORTEDOPERATION_DEFAULTCERTIFICATESOPERATE = "UnsupportedOperation.DefaultCertificatesOperate"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
func (c *Client) DescribeAccelerateRegionsWithContext(ctx context.Context, request *DescribeAccelerateRegionsRequest) (response *DescribeAccelerateRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeAccelerateRegionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "DescribeAccelerateRegions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccelerateRegions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccelerateRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessLogParamRequest() (request *DescribeAccessLogParamRequest) {
    request = &DescribeAccessLogParamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "DescribeAccessLogParam")
    
    
    return
}

func NewDescribeAccessLogParamResponse() (response *DescribeAccessLogParamResponse) {
    response = &DescribeAccessLogParamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccessLogParam
// View access log reporting parameters
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  UNSUPPORTEDOPERATION_CACERTIFICATESOPERATE = "UnsupportedOperation.CaCertificatesOperate"
//  UNSUPPORTEDOPERATION_DEFAULTCERTIFICATESOPERATE = "UnsupportedOperation.DefaultCertificatesOperate"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
func (c *Client) DescribeAccessLogParam(request *DescribeAccessLogParamRequest) (response *DescribeAccessLogParamResponse, err error) {
    return c.DescribeAccessLogParamWithContext(context.Background(), request)
}

// DescribeAccessLogParam
// View access log reporting parameters
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  UNSUPPORTEDOPERATION_CACERTIFICATESOPERATE = "UnsupportedOperation.CaCertificatesOperate"
//  UNSUPPORTEDOPERATION_DEFAULTCERTIFICATESOPERATE = "UnsupportedOperation.DefaultCertificatesOperate"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
func (c *Client) DescribeAccessLogParamWithContext(ctx context.Context, request *DescribeAccessLogParamRequest) (response *DescribeAccessLogParamResponse, err error) {
    if request == nil {
        request = NewDescribeAccessLogParamRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "DescribeAccessLogParam")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccessLogParam require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccessLogParamResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCrossBorderSettlementRequest() (request *DescribeCrossBorderSettlementRequest) {
    request = &DescribeCrossBorderSettlementRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "DescribeCrossBorderSettlement")
    
    
    return
}

func NewDescribeCrossBorderSettlementResponse() (response *DescribeCrossBorderSettlementResponse) {
    response = &DescribeCrossBorderSettlementResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCrossBorderSettlement
// Querying Cross-Border Bills
//
// error code that may be returned:
//  INVALIDPARAMETER_INPUTOUTOFRANGE = "InvalidParameter.InputOutOfRange"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
func (c *Client) DescribeCrossBorderSettlement(request *DescribeCrossBorderSettlementRequest) (response *DescribeCrossBorderSettlementResponse, err error) {
    return c.DescribeCrossBorderSettlementWithContext(context.Background(), request)
}

// DescribeCrossBorderSettlement
// Querying Cross-Border Bills
//
// error code that may be returned:
//  INVALIDPARAMETER_INPUTOUTOFRANGE = "InvalidParameter.InputOutOfRange"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
func (c *Client) DescribeCrossBorderSettlementWithContext(ctx context.Context, request *DescribeCrossBorderSettlementRequest) (response *DescribeCrossBorderSettlementResponse, err error) {
    if request == nil {
        request = NewDescribeCrossBorderSettlementRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "DescribeCrossBorderSettlement")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCrossBorderSettlement require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCrossBorderSettlementResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEndpointGroupsRequest() (request *DescribeEndpointGroupsRequest) {
    request = &DescribeEndpointGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "DescribeEndpointGroups")
    
    
    return
}

func NewDescribeEndpointGroupsResponse() (response *DescribeEndpointGroupsResponse) {
    response = &DescribeEndpointGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEndpointGroups
// Query a terminal node group.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeEndpointGroups(request *DescribeEndpointGroupsRequest) (response *DescribeEndpointGroupsResponse, err error) {
    return c.DescribeEndpointGroupsWithContext(context.Background(), request)
}

// DescribeEndpointGroups
// Query a terminal node group.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeEndpointGroupsWithContext(ctx context.Context, request *DescribeEndpointGroupsRequest) (response *DescribeEndpointGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeEndpointGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "DescribeEndpointGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEndpointGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEndpointGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeForwardingPolicyRequest() (request *DescribeForwardingPolicyRequest) {
    request = &DescribeForwardingPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "DescribeForwardingPolicy")
    
    
    return
}

func NewDescribeForwardingPolicyResponse() (response *DescribeForwardingPolicyResponse) {
    response = &DescribeForwardingPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeForwardingPolicy
// View a layer-7 forwarding policy
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeForwardingPolicy(request *DescribeForwardingPolicyRequest) (response *DescribeForwardingPolicyResponse, err error) {
    return c.DescribeForwardingPolicyWithContext(context.Background(), request)
}

// DescribeForwardingPolicy
// View a layer-7 forwarding policy
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeForwardingPolicyWithContext(ctx context.Context, request *DescribeForwardingPolicyRequest) (response *DescribeForwardingPolicyResponse, err error) {
    if request == nil {
        request = NewDescribeForwardingPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "DescribeForwardingPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeForwardingPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeForwardingPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeForwardingRuleRequest() (request *DescribeForwardingRuleRequest) {
    request = &DescribeForwardingRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "DescribeForwardingRule")
    
    
    return
}

func NewDescribeForwardingRuleResponse() (response *DescribeForwardingRuleResponse) {
    response = &DescribeForwardingRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeForwardingRule
// View a Layer 7 forwarding rule
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeForwardingRule(request *DescribeForwardingRuleRequest) (response *DescribeForwardingRuleResponse, err error) {
    return c.DescribeForwardingRuleWithContext(context.Background(), request)
}

// DescribeForwardingRule
// View a Layer 7 forwarding rule
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeForwardingRuleWithContext(ctx context.Context, request *DescribeForwardingRuleRequest) (response *DescribeForwardingRuleResponse, err error) {
    if request == nil {
        request = NewDescribeForwardingRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "DescribeForwardingRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeForwardingRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeForwardingRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGlobalAcceleratorAccessLogRequest() (request *DescribeGlobalAcceleratorAccessLogRequest) {
    request = &DescribeGlobalAcceleratorAccessLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "DescribeGlobalAcceleratorAccessLog")
    
    
    return
}

func NewDescribeGlobalAcceleratorAccessLogResponse() (response *DescribeGlobalAcceleratorAccessLogResponse) {
    response = &DescribeGlobalAcceleratorAccessLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGlobalAcceleratorAccessLog
// Query log tasks
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeGlobalAcceleratorAccessLog(request *DescribeGlobalAcceleratorAccessLogRequest) (response *DescribeGlobalAcceleratorAccessLogResponse, err error) {
    return c.DescribeGlobalAcceleratorAccessLogWithContext(context.Background(), request)
}

// DescribeGlobalAcceleratorAccessLog
// Query log tasks
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeGlobalAcceleratorAccessLogWithContext(ctx context.Context, request *DescribeGlobalAcceleratorAccessLogRequest) (response *DescribeGlobalAcceleratorAccessLogResponse, err error) {
    if request == nil {
        request = NewDescribeGlobalAcceleratorAccessLogRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "DescribeGlobalAcceleratorAccessLog")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGlobalAcceleratorAccessLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGlobalAcceleratorAccessLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGlobalAcceleratorAclPoliciesRequest() (request *DescribeGlobalAcceleratorAclPoliciesRequest) {
    request = &DescribeGlobalAcceleratorAclPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "DescribeGlobalAcceleratorAclPolicies")
    
    
    return
}

func NewDescribeGlobalAcceleratorAclPoliciesResponse() (response *DescribeGlobalAcceleratorAclPoliciesResponse) {
    response = &DescribeGlobalAcceleratorAclPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGlobalAcceleratorAclPolicies
// View the access control policy
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeGlobalAcceleratorAclPolicies(request *DescribeGlobalAcceleratorAclPoliciesRequest) (response *DescribeGlobalAcceleratorAclPoliciesResponse, err error) {
    return c.DescribeGlobalAcceleratorAclPoliciesWithContext(context.Background(), request)
}

// DescribeGlobalAcceleratorAclPolicies
// View the access control policy
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeGlobalAcceleratorAclPoliciesWithContext(ctx context.Context, request *DescribeGlobalAcceleratorAclPoliciesRequest) (response *DescribeGlobalAcceleratorAclPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeGlobalAcceleratorAclPoliciesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "DescribeGlobalAcceleratorAclPolicies")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGlobalAcceleratorAclPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGlobalAcceleratorAclPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGlobalAcceleratorAclRulesRequest() (request *DescribeGlobalAcceleratorAclRulesRequest) {
    request = &DescribeGlobalAcceleratorAclRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "DescribeGlobalAcceleratorAclRules")
    
    
    return
}

func NewDescribeGlobalAcceleratorAclRulesResponse() (response *DescribeGlobalAcceleratorAclRulesResponse) {
    response = &DescribeGlobalAcceleratorAclRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGlobalAcceleratorAclRules
// View ACL rules
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeGlobalAcceleratorAclRules(request *DescribeGlobalAcceleratorAclRulesRequest) (response *DescribeGlobalAcceleratorAclRulesResponse, err error) {
    return c.DescribeGlobalAcceleratorAclRulesWithContext(context.Background(), request)
}

// DescribeGlobalAcceleratorAclRules
// View ACL rules
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeGlobalAcceleratorAclRulesWithContext(ctx context.Context, request *DescribeGlobalAcceleratorAclRulesRequest) (response *DescribeGlobalAcceleratorAclRulesResponse, err error) {
    if request == nil {
        request = NewDescribeGlobalAcceleratorAclRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "DescribeGlobalAcceleratorAclRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGlobalAcceleratorAclRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGlobalAcceleratorAclRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGlobalAcceleratorsRequest() (request *DescribeGlobalAcceleratorsRequest) {
    request = &DescribeGlobalAcceleratorsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "DescribeGlobalAccelerators")
    
    
    return
}

func NewDescribeGlobalAcceleratorsResponse() (response *DescribeGlobalAcceleratorsResponse) {
    response = &DescribeGlobalAcceleratorsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGlobalAccelerators
// Modify a global acceleration instance
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEOFFSET = "InvalidParameterValue.InvalidParameterValueOffset"
func (c *Client) DescribeGlobalAccelerators(request *DescribeGlobalAcceleratorsRequest) (response *DescribeGlobalAcceleratorsResponse, err error) {
    return c.DescribeGlobalAcceleratorsWithContext(context.Background(), request)
}

// DescribeGlobalAccelerators
// Modify a global acceleration instance
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEOFFSET = "InvalidParameterValue.InvalidParameterValueOffset"
func (c *Client) DescribeGlobalAcceleratorsWithContext(ctx context.Context, request *DescribeGlobalAcceleratorsRequest) (response *DescribeGlobalAcceleratorsResponse, err error) {
    if request == nil {
        request = NewDescribeGlobalAcceleratorsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "DescribeGlobalAccelerators")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGlobalAccelerators require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGlobalAcceleratorsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListenersRequest() (request *DescribeListenersRequest) {
    request = &DescribeListenersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "DescribeListeners")
    
    
    return
}

func NewDescribeListenersResponse() (response *DescribeListenersResponse) {
    response = &DescribeListenersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeListeners
// This API is used to query listeners.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeListeners(request *DescribeListenersRequest) (response *DescribeListenersResponse, err error) {
    return c.DescribeListenersWithContext(context.Background(), request)
}

// DescribeListeners
// This API is used to query listeners.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeListenersWithContext(ctx context.Context, request *DescribeListenersRequest) (response *DescribeListenersResponse, err error) {
    if request == nil {
        request = NewDescribeListenersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "DescribeListeners")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeListeners require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeListenersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskResultRequest() (request *DescribeTaskResultRequest) {
    request = &DescribeTaskResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "DescribeTaskResult")
    
    
    return
}

func NewDescribeTaskResultResponse() (response *DescribeTaskResultResponse) {
    response = &DescribeTaskResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTaskResult
// Query asynchronous task result
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_TASKNOTFOUND = "InvalidParameterValue.TaskNotFound"
func (c *Client) DescribeTaskResult(request *DescribeTaskResultRequest) (response *DescribeTaskResultResponse, err error) {
    return c.DescribeTaskResultWithContext(context.Background(), request)
}

// DescribeTaskResult
// Query asynchronous task result
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_TASKNOTFOUND = "InvalidParameterValue.TaskNotFound"
func (c *Client) DescribeTaskResultWithContext(ctx context.Context, request *DescribeTaskResultRequest) (response *DescribeTaskResultResponse, err error) {
    if request == nil {
        request = NewDescribeTaskResultRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "DescribeTaskResult")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskResultResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccelerateAreasRequest() (request *ModifyAccelerateAreasRequest) {
    request = &ModifyAccelerateAreasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "ModifyAccelerateAreas")
    
    
    return
}

func NewModifyAccelerateAreasResponse() (response *ModifyAccelerateAreasResponse) {
    response = &ModifyAccelerateAreasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAccelerateAreas
// Modify acceleration region
//
// error code that may be returned:
//  INVALIDPARAMETER_INPUTDUPLICATEWITHEXISTINGDATA = "InvalidParameter.InputDuplicateWithExistingData"
//  INVALIDPARAMETER_INPUTREPEAT = "InvalidParameter.InputRepeat"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTALLOWEDOPERATE = "UnsupportedOperation.InstanceStateNotAllowedOperate"
func (c *Client) ModifyAccelerateAreas(request *ModifyAccelerateAreasRequest) (response *ModifyAccelerateAreasResponse, err error) {
    return c.ModifyAccelerateAreasWithContext(context.Background(), request)
}

// ModifyAccelerateAreas
// Modify acceleration region
//
// error code that may be returned:
//  INVALIDPARAMETER_INPUTDUPLICATEWITHEXISTINGDATA = "InvalidParameter.InputDuplicateWithExistingData"
//  INVALIDPARAMETER_INPUTREPEAT = "InvalidParameter.InputRepeat"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTALLOWEDOPERATE = "UnsupportedOperation.InstanceStateNotAllowedOperate"
func (c *Client) ModifyAccelerateAreasWithContext(ctx context.Context, request *ModifyAccelerateAreasRequest) (response *ModifyAccelerateAreasResponse, err error) {
    if request == nil {
        request = NewModifyAccelerateAreasRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "ModifyAccelerateAreas")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccelerateAreas require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAccelerateAreasResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccessLogStatusRequest() (request *ModifyAccessLogStatusRequest) {
    request = &ModifyAccessLogStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "ModifyAccessLogStatus")
    
    
    return
}

func NewModifyAccessLogStatusResponse() (response *ModifyAccessLogStatusResponse) {
    response = &ModifyAccessLogStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAccessLogStatus
// Modify the status of a log task
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_LOGSETNOTEXIST = "InvalidParameterValue.LogsetNotExist"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTALLOWEDOPERATE = "UnsupportedOperation.InstanceStateNotAllowedOperate"
//  UNSUPPORTEDOPERATION_NOTAUTHORIZATION = "UnsupportedOperation.NotAuthorization"
func (c *Client) ModifyAccessLogStatus(request *ModifyAccessLogStatusRequest) (response *ModifyAccessLogStatusResponse, err error) {
    return c.ModifyAccessLogStatusWithContext(context.Background(), request)
}

// ModifyAccessLogStatus
// Modify the status of a log task
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_LOGSETNOTEXIST = "InvalidParameterValue.LogsetNotExist"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTALLOWEDOPERATE = "UnsupportedOperation.InstanceStateNotAllowedOperate"
//  UNSUPPORTEDOPERATION_NOTAUTHORIZATION = "UnsupportedOperation.NotAuthorization"
func (c *Client) ModifyAccessLogStatusWithContext(ctx context.Context, request *ModifyAccessLogStatusRequest) (response *ModifyAccessLogStatusResponse, err error) {
    if request == nil {
        request = NewModifyAccessLogStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "ModifyAccessLogStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccessLogStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAccessLogStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEndpointGroupRequest() (request *ModifyEndpointGroupRequest) {
    request = &ModifyEndpointGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "ModifyEndpointGroup")
    
    
    return
}

func NewModifyEndpointGroupResponse() (response *ModifyEndpointGroupResponse) {
    response = &ModifyEndpointGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyEndpointGroup
// This API is used to modify a terminal node group.
//
// error code that may be returned:
//  INVALIDPARAMETER_DISABLEHEALTHCHECKNOTCARRYPARAMETERS = "InvalidParameter.DisableHealthCheckNotCarryParameters"
//  INVALIDPARAMETER_ENDPOINTGROUPCHECKTYPEHTTPCARRYPARAMETERS = "InvalidParameter.EndpointGroupCheckTypeHttpCarryParameters"
//  INVALIDPARAMETER_HTTPSORIGINSERVERPROTOCOLSUPPORTSPARAMETERS = "InvalidParameter.HttpsOriginServerProtocolSupportsParameters"
//  INVALIDPARAMETER_INPUTMUSTEXISTONE = "InvalidParameter.InputMustExistOne"
//  INVALIDPARAMETER_INPUTOUTOFRANGE = "InvalidParameter.InputOutOfRange"
//  INVALIDPARAMETER_INPUTREPEAT = "InvalidParameter.InputRepeat"
//  INVALIDPARAMETER_ONLYENDPOINTGROUPCHECKTYPEHTTPCARRYPARAMETERS = "InvalidParameter.OnlyEndpointGroupCheckTypeHttpCarryParameters"
//  INVALIDPARAMETER_TRANSPORTLAYERENDPOINTGROUPCANNOTCARRYPARAMETERS = "InvalidParameter.TransportLayerEndpointGroupCannotCarryParameters"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  INVALIDPARAMETERVALUE_INSTANCENAME = "InvalidParameterValue.InstanceName"
//  INVALIDPARAMETERVALUE_LENGTHRANGE = "InvalidParameterValue.LengthRange"
//  INVALIDPARAMETERVALUE_LISTENERPORTNOTEQUAL = "InvalidParameterValue.ListenerPortNotEqual"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NOTMAPPINGMANYPORT = "InvalidParameterValue.NotMappingManyPort"
//  INVALIDPARAMETERVALUE_NOTTENCENTREGION = "InvalidParameterValue.NotTencentRegion"
//  INVALIDPARAMETERVALUE_NOTWITHINPUBLICNETWORK = "InvalidParameterValue.NotWithinPublicNetwork"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_WRONGATTRIBUTIONRELATIONSHIP = "InvalidParameterValue.WrongAttributionRelationship"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER_CUSTOMCHECKTYPEPARAMETER = "MissingParameter.CustomCheckTypeParameter"
//  MISSINGPARAMETER_ENABLEHEALTHCHECKPARAMETER = "MissingParameter.EnableHealthCheckParameter"
//  MISSINGPARAMETER_INTERNALNETWORKSOURCE = "MissingParameter.InternalNetworkSource"
//  UNSUPPORTEDOPERATION_CURRENTACCOUNTNOTALLOWED = "UnsupportedOperation.CurrentAccountNotAllowed"
//  UNSUPPORTEDOPERATION_HTTPVERSIONINCONSISTENT = "UnsupportedOperation.HttpVersionInconsistent"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
//  UNSUPPORTEDOPERATION_INTERNALENDPOINTFEATURENOTENABLED = "UnsupportedOperation.InternalEndpointFeatureNotEnabled"
//  UNSUPPORTEDOPERATION_IPV6NOTSUPPORT = "UnsupportedOperation.Ipv6NotSupport"
//  UNSUPPORTEDOPERATION_ONLYINTERNALCLB = "UnsupportedOperation.OnlyInternalClb"
//  UNSUPPORTEDOPERATION_SAMEDOMAIN = "UnsupportedOperation.SameDomain"
//  UNSUPPORTEDOPERATION_SAMEPUBLICIP = "UnsupportedOperation.SamePublicIp"
func (c *Client) ModifyEndpointGroup(request *ModifyEndpointGroupRequest) (response *ModifyEndpointGroupResponse, err error) {
    return c.ModifyEndpointGroupWithContext(context.Background(), request)
}

// ModifyEndpointGroup
// This API is used to modify a terminal node group.
//
// error code that may be returned:
//  INVALIDPARAMETER_DISABLEHEALTHCHECKNOTCARRYPARAMETERS = "InvalidParameter.DisableHealthCheckNotCarryParameters"
//  INVALIDPARAMETER_ENDPOINTGROUPCHECKTYPEHTTPCARRYPARAMETERS = "InvalidParameter.EndpointGroupCheckTypeHttpCarryParameters"
//  INVALIDPARAMETER_HTTPSORIGINSERVERPROTOCOLSUPPORTSPARAMETERS = "InvalidParameter.HttpsOriginServerProtocolSupportsParameters"
//  INVALIDPARAMETER_INPUTMUSTEXISTONE = "InvalidParameter.InputMustExistOne"
//  INVALIDPARAMETER_INPUTOUTOFRANGE = "InvalidParameter.InputOutOfRange"
//  INVALIDPARAMETER_INPUTREPEAT = "InvalidParameter.InputRepeat"
//  INVALIDPARAMETER_ONLYENDPOINTGROUPCHECKTYPEHTTPCARRYPARAMETERS = "InvalidParameter.OnlyEndpointGroupCheckTypeHttpCarryParameters"
//  INVALIDPARAMETER_TRANSPORTLAYERENDPOINTGROUPCANNOTCARRYPARAMETERS = "InvalidParameter.TransportLayerEndpointGroupCannotCarryParameters"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  INVALIDPARAMETERVALUE_INSTANCENAME = "InvalidParameterValue.InstanceName"
//  INVALIDPARAMETERVALUE_LENGTHRANGE = "InvalidParameterValue.LengthRange"
//  INVALIDPARAMETERVALUE_LISTENERPORTNOTEQUAL = "InvalidParameterValue.ListenerPortNotEqual"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NOTMAPPINGMANYPORT = "InvalidParameterValue.NotMappingManyPort"
//  INVALIDPARAMETERVALUE_NOTTENCENTREGION = "InvalidParameterValue.NotTencentRegion"
//  INVALIDPARAMETERVALUE_NOTWITHINPUBLICNETWORK = "InvalidParameterValue.NotWithinPublicNetwork"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_WRONGATTRIBUTIONRELATIONSHIP = "InvalidParameterValue.WrongAttributionRelationship"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER_CUSTOMCHECKTYPEPARAMETER = "MissingParameter.CustomCheckTypeParameter"
//  MISSINGPARAMETER_ENABLEHEALTHCHECKPARAMETER = "MissingParameter.EnableHealthCheckParameter"
//  MISSINGPARAMETER_INTERNALNETWORKSOURCE = "MissingParameter.InternalNetworkSource"
//  UNSUPPORTEDOPERATION_CURRENTACCOUNTNOTALLOWED = "UnsupportedOperation.CurrentAccountNotAllowed"
//  UNSUPPORTEDOPERATION_HTTPVERSIONINCONSISTENT = "UnsupportedOperation.HttpVersionInconsistent"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
//  UNSUPPORTEDOPERATION_INTERNALENDPOINTFEATURENOTENABLED = "UnsupportedOperation.InternalEndpointFeatureNotEnabled"
//  UNSUPPORTEDOPERATION_IPV6NOTSUPPORT = "UnsupportedOperation.Ipv6NotSupport"
//  UNSUPPORTEDOPERATION_ONLYINTERNALCLB = "UnsupportedOperation.OnlyInternalClb"
//  UNSUPPORTEDOPERATION_SAMEDOMAIN = "UnsupportedOperation.SameDomain"
//  UNSUPPORTEDOPERATION_SAMEPUBLICIP = "UnsupportedOperation.SamePublicIp"
func (c *Client) ModifyEndpointGroupWithContext(ctx context.Context, request *ModifyEndpointGroupRequest) (response *ModifyEndpointGroupResponse, err error) {
    if request == nil {
        request = NewModifyEndpointGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "ModifyEndpointGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEndpointGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyEndpointGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyForwardingPolicyRequest() (request *ModifyForwardingPolicyRequest) {
    request = &ModifyForwardingPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "ModifyForwardingPolicy")
    
    
    return
}

func NewModifyForwardingPolicyResponse() (response *ModifyForwardingPolicyResponse) {
    response = &ModifyForwardingPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyForwardingPolicy
// Modify a layer-7 forwarding policy
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_DEFAULTFORWARDINGPOLICYOPERATE = "UnsupportedOperation.DefaultForwardingPolicyOperate"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
//  UNSUPPORTEDOPERATION_TRANSPORTLAYERUNSUPPORTEDOPERATEFORWARDINGPOLICY = "UnsupportedOperation.TransportLayerUnsupportedOperateForwardingPolicy"
func (c *Client) ModifyForwardingPolicy(request *ModifyForwardingPolicyRequest) (response *ModifyForwardingPolicyResponse, err error) {
    return c.ModifyForwardingPolicyWithContext(context.Background(), request)
}

// ModifyForwardingPolicy
// Modify a layer-7 forwarding policy
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_DEFAULTFORWARDINGPOLICYOPERATE = "UnsupportedOperation.DefaultForwardingPolicyOperate"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
//  UNSUPPORTEDOPERATION_TRANSPORTLAYERUNSUPPORTEDOPERATEFORWARDINGPOLICY = "UnsupportedOperation.TransportLayerUnsupportedOperateForwardingPolicy"
func (c *Client) ModifyForwardingPolicyWithContext(ctx context.Context, request *ModifyForwardingPolicyRequest) (response *ModifyForwardingPolicyResponse, err error) {
    if request == nil {
        request = NewModifyForwardingPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "ModifyForwardingPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyForwardingPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyForwardingPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyForwardingRuleRequest() (request *ModifyForwardingRuleRequest) {
    request = &ModifyForwardingRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "ModifyForwardingRule")
    
    
    return
}

func NewModifyForwardingRuleResponse() (response *ModifyForwardingRuleResponse) {
    response = &ModifyForwardingRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyForwardingRule
// This API is used to modify a Layer 7 forwarding rule.
//
// error code that may be returned:
//  INVALIDPARAMETER_INPUTREPEAT = "InvalidParameter.InputRepeat"
//  INVALIDPARAMETERVALUE_ASCIICHARACTERS = "InvalidParameterValue.AsciiCharacters"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_FORWARDINGRULEPATHCONFLICT = "InvalidParameterValue.ForwardingRulePathConflict"
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  INVALIDPARAMETERVALUE_INTERNALRESERVEDFIELDS = "InvalidParameterValue.InternalReservedFields"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_SPECIALCHARACTERS = "InvalidParameterValue.SpecialCharacters"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_DEFAULTFORWARDINGRULEOPERATE = "UnsupportedOperation.DefaultForwardingRuleOperate"
//  UNSUPPORTEDOPERATION_FORWARDGROUPFORWARDINGRULEUNSUPPORTEDDEFAULTENDPOINTGROUP = "UnsupportedOperation.ForwardGroupForwardingRuleUnsupportedDefaultEndpointGroup"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
//  UNSUPPORTEDOPERATION_TRANSPORTLAYERUNSUPPORTEDOPERATEFORWARDINGPOLICY = "UnsupportedOperation.TransportLayerUnsupportedOperateForwardingPolicy"
func (c *Client) ModifyForwardingRule(request *ModifyForwardingRuleRequest) (response *ModifyForwardingRuleResponse, err error) {
    return c.ModifyForwardingRuleWithContext(context.Background(), request)
}

// ModifyForwardingRule
// This API is used to modify a Layer 7 forwarding rule.
//
// error code that may be returned:
//  INVALIDPARAMETER_INPUTREPEAT = "InvalidParameter.InputRepeat"
//  INVALIDPARAMETERVALUE_ASCIICHARACTERS = "InvalidParameterValue.AsciiCharacters"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_FORWARDINGRULEPATHCONFLICT = "InvalidParameterValue.ForwardingRulePathConflict"
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  INVALIDPARAMETERVALUE_INTERNALRESERVEDFIELDS = "InvalidParameterValue.InternalReservedFields"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_SPECIALCHARACTERS = "InvalidParameterValue.SpecialCharacters"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_DEFAULTFORWARDINGRULEOPERATE = "UnsupportedOperation.DefaultForwardingRuleOperate"
//  UNSUPPORTEDOPERATION_FORWARDGROUPFORWARDINGRULEUNSUPPORTEDDEFAULTENDPOINTGROUP = "UnsupportedOperation.ForwardGroupForwardingRuleUnsupportedDefaultEndpointGroup"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
//  UNSUPPORTEDOPERATION_TRANSPORTLAYERUNSUPPORTEDOPERATEFORWARDINGPOLICY = "UnsupportedOperation.TransportLayerUnsupportedOperateForwardingPolicy"
func (c *Client) ModifyForwardingRuleWithContext(ctx context.Context, request *ModifyForwardingRuleRequest) (response *ModifyForwardingRuleResponse, err error) {
    if request == nil {
        request = NewModifyForwardingRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "ModifyForwardingRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyForwardingRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyForwardingRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyGlobalAcceleratorRequest() (request *ModifyGlobalAcceleratorRequest) {
    request = &ModifyGlobalAcceleratorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "ModifyGlobalAccelerator")
    
    
    return
}

func NewModifyGlobalAcceleratorResponse() (response *ModifyGlobalAcceleratorResponse) {
    response = &ModifyGlobalAcceleratorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyGlobalAccelerator
// Modify a global acceleration instance
//
// error code that may be returned:
//  INVALIDPARAMETER_INPUTMUSTEXISTONE = "InvalidParameter.InputMustExistOne"
//  UNSUPPORTEDOPERATION_ALREADYENABLECROSSBORDER = "UnsupportedOperation.AlreadyEnableCrossBorder"
//  UNSUPPORTEDOPERATION_CROSSBORDERPROMISENOTSIGNED = "UnsupportedOperation.CrossBorderPromiseNotSigned"
//  UNSUPPORTEDOPERATION_ENABLECROSSBORDER = "UnsupportedOperation.EnableCrossBorder"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTALLOWEDOPERATE = "UnsupportedOperation.InstanceStateNotAllowedOperate"
func (c *Client) ModifyGlobalAccelerator(request *ModifyGlobalAcceleratorRequest) (response *ModifyGlobalAcceleratorResponse, err error) {
    return c.ModifyGlobalAcceleratorWithContext(context.Background(), request)
}

// ModifyGlobalAccelerator
// Modify a global acceleration instance
//
// error code that may be returned:
//  INVALIDPARAMETER_INPUTMUSTEXISTONE = "InvalidParameter.InputMustExistOne"
//  UNSUPPORTEDOPERATION_ALREADYENABLECROSSBORDER = "UnsupportedOperation.AlreadyEnableCrossBorder"
//  UNSUPPORTEDOPERATION_CROSSBORDERPROMISENOTSIGNED = "UnsupportedOperation.CrossBorderPromiseNotSigned"
//  UNSUPPORTEDOPERATION_ENABLECROSSBORDER = "UnsupportedOperation.EnableCrossBorder"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTALLOWEDOPERATE = "UnsupportedOperation.InstanceStateNotAllowedOperate"
func (c *Client) ModifyGlobalAcceleratorWithContext(ctx context.Context, request *ModifyGlobalAcceleratorRequest) (response *ModifyGlobalAcceleratorResponse, err error) {
    if request == nil {
        request = NewModifyGlobalAcceleratorRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "ModifyGlobalAccelerator")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyGlobalAccelerator require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyGlobalAcceleratorResponse()
    err = c.Send(request, response)
    return
}

func NewModifyGlobalAcceleratorAccessLogRequest() (request *ModifyGlobalAcceleratorAccessLogRequest) {
    request = &ModifyGlobalAcceleratorAccessLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "ModifyGlobalAcceleratorAccessLog")
    
    
    return
}

func NewModifyGlobalAcceleratorAccessLogResponse() (response *ModifyGlobalAcceleratorAccessLogResponse) {
    response = &ModifyGlobalAcceleratorAccessLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyGlobalAcceleratorAccessLog
// Modify GA access logs
//
// error code that may be returned:
//  INVALIDPARAMETER_INPUTREPEAT = "InvalidParameter.InputRepeat"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_LOGSETNOTEXIST = "InvalidParameterValue.LogsetNotExist"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
//  UNSUPPORTEDOPERATION_NOTAUTHORIZATION = "UnsupportedOperation.NotAuthorization"
func (c *Client) ModifyGlobalAcceleratorAccessLog(request *ModifyGlobalAcceleratorAccessLogRequest) (response *ModifyGlobalAcceleratorAccessLogResponse, err error) {
    return c.ModifyGlobalAcceleratorAccessLogWithContext(context.Background(), request)
}

// ModifyGlobalAcceleratorAccessLog
// Modify GA access logs
//
// error code that may be returned:
//  INVALIDPARAMETER_INPUTREPEAT = "InvalidParameter.InputRepeat"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_LOGSETNOTEXIST = "InvalidParameterValue.LogsetNotExist"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
//  UNSUPPORTEDOPERATION_NOTAUTHORIZATION = "UnsupportedOperation.NotAuthorization"
func (c *Client) ModifyGlobalAcceleratorAccessLogWithContext(ctx context.Context, request *ModifyGlobalAcceleratorAccessLogRequest) (response *ModifyGlobalAcceleratorAccessLogResponse, err error) {
    if request == nil {
        request = NewModifyGlobalAcceleratorAccessLogRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "ModifyGlobalAcceleratorAccessLog")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyGlobalAcceleratorAccessLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyGlobalAcceleratorAccessLogResponse()
    err = c.Send(request, response)
    return
}

func NewModifyGlobalAcceleratorAclPolicyRequest() (request *ModifyGlobalAcceleratorAclPolicyRequest) {
    request = &ModifyGlobalAcceleratorAclPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "ModifyGlobalAcceleratorAclPolicy")
    
    
    return
}

func NewModifyGlobalAcceleratorAclPolicyResponse() (response *ModifyGlobalAcceleratorAclPolicyResponse) {
    response = &ModifyGlobalAcceleratorAclPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyGlobalAcceleratorAclPolicy
// Modify the status of an access control policy
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  UNSUPPORTEDOPERATION_DUPLICATEINSTANCESTATUS = "UnsupportedOperation.DuplicateInstanceStatus"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTALLOWEDOPERATE = "UnsupportedOperation.InstanceStateNotAllowedOperate"
func (c *Client) ModifyGlobalAcceleratorAclPolicy(request *ModifyGlobalAcceleratorAclPolicyRequest) (response *ModifyGlobalAcceleratorAclPolicyResponse, err error) {
    return c.ModifyGlobalAcceleratorAclPolicyWithContext(context.Background(), request)
}

// ModifyGlobalAcceleratorAclPolicy
// Modify the status of an access control policy
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  UNSUPPORTEDOPERATION_DUPLICATEINSTANCESTATUS = "UnsupportedOperation.DuplicateInstanceStatus"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTALLOWEDOPERATE = "UnsupportedOperation.InstanceStateNotAllowedOperate"
func (c *Client) ModifyGlobalAcceleratorAclPolicyWithContext(ctx context.Context, request *ModifyGlobalAcceleratorAclPolicyRequest) (response *ModifyGlobalAcceleratorAclPolicyResponse, err error) {
    if request == nil {
        request = NewModifyGlobalAcceleratorAclPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "ModifyGlobalAcceleratorAclPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyGlobalAcceleratorAclPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyGlobalAcceleratorAclPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyGlobalAcceleratorAclRuleRequest() (request *ModifyGlobalAcceleratorAclRuleRequest) {
    request = &ModifyGlobalAcceleratorAclRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "ModifyGlobalAcceleratorAclRule")
    
    
    return
}

func NewModifyGlobalAcceleratorAclRuleResponse() (response *ModifyGlobalAcceleratorAclRuleResponse) {
    response = &ModifyGlobalAcceleratorAclRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyGlobalAcceleratorAclRule
// Modify ACL rules
//
// error code that may be returned:
//  INVALIDPARAMETER_INPUTDUPLICATEWITHEXISTINGDATA = "InvalidParameter.InputDuplicateWithExistingData"
//  INVALIDPARAMETER_INPUTMUSTEXISTONE = "InvalidParameter.InputMustExistOne"
//  INVALIDPARAMETER_INPUTOUTOFRANGE = "InvalidParameter.InputOutOfRange"
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTALLOWEDOPERATE = "UnsupportedOperation.InstanceStateNotAllowedOperate"
func (c *Client) ModifyGlobalAcceleratorAclRule(request *ModifyGlobalAcceleratorAclRuleRequest) (response *ModifyGlobalAcceleratorAclRuleResponse, err error) {
    return c.ModifyGlobalAcceleratorAclRuleWithContext(context.Background(), request)
}

// ModifyGlobalAcceleratorAclRule
// Modify ACL rules
//
// error code that may be returned:
//  INVALIDPARAMETER_INPUTDUPLICATEWITHEXISTINGDATA = "InvalidParameter.InputDuplicateWithExistingData"
//  INVALIDPARAMETER_INPUTMUSTEXISTONE = "InvalidParameter.InputMustExistOne"
//  INVALIDPARAMETER_INPUTOUTOFRANGE = "InvalidParameter.InputOutOfRange"
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTALLOWEDOPERATE = "UnsupportedOperation.InstanceStateNotAllowedOperate"
func (c *Client) ModifyGlobalAcceleratorAclRuleWithContext(ctx context.Context, request *ModifyGlobalAcceleratorAclRuleRequest) (response *ModifyGlobalAcceleratorAclRuleResponse, err error) {
    if request == nil {
        request = NewModifyGlobalAcceleratorAclRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "ModifyGlobalAcceleratorAclRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyGlobalAcceleratorAclRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyGlobalAcceleratorAclRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyListenerRequest() (request *ModifyListenerRequest) {
    request = &ModifyListenerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "ModifyListener")
    
    
    return
}

func NewModifyListenerResponse() (response *ModifyListenerResponse) {
    response = &ModifyListenerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyListener
// Modify a listener
//
// error code that may be returned:
//  INVALIDPARAMETER_APPLICATIONLAYERLISTENERCANNOTCARRYPARAMETERS = "InvalidParameter.ApplicationLayerListenerCannotCarryParameters"
//  INVALIDPARAMETER_HTTPLISTENERCANNOTCARRYPARAMETERS = "InvalidParameter.HttpListenerCannotCarryParameters"
//  INVALIDPARAMETER_HTTPSLISTENERCANCARRYPARAMETERS = "InvalidParameter.HttpsListenerCanCarryParameters"
//  INVALIDPARAMETER_INPUTMUSTEXISTONE = "InvalidParameter.InputMustExistOne"
//  INVALIDPARAMETER_TCPLISTENERCARRYPARAMETERS = "InvalidParameter.TcpListenerCarryParameters"
//  INVALIDPARAMETER_TRANSPORTLAYERLISTENERCANNOTCARRYPARAMETERS = "InvalidParameter.TransportLayerListenerCannotCarryParameters"
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  INVALIDPARAMETERVALUE_INSTANCENAME = "InvalidParameterValue.InstanceName"
//  INVALIDPARAMETERVALUE_LENGTHRANGE = "InvalidParameterValue.LengthRange"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_CERTIFICATENOTCA = "UnsupportedOperation.CertificateNotCa"
//  UNSUPPORTEDOPERATION_CURRENTACCOUNTNOTALLOWED = "UnsupportedOperation.CurrentAccountNotAllowed"
//  UNSUPPORTEDOPERATION_HTTPVERSIONINCONSISTENT = "UnsupportedOperation.HttpVersionInconsistent"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
//  UNSUPPORTEDOPERATION_ONLYTCPLISTENERSUPPORTTOA = "UnsupportedOperation.OnlyTcpListenerSupportToa"
//  UNSUPPORTEDOPERATION_TOAFEATURENOTENABLED = "UnsupportedOperation.TOAFeatureNotEnabled"
//  UNSUPPORTEDOPERATION_UDPLISTENERNOTSUPPORTPROXYPROTOCOL = "UnsupportedOperation.UdpListenerNotSupportProxyProtocol"
//  UNSUPPORTEDOPERATION_UNLAWFULCERTIFICATE = "UnsupportedOperation.UnlawfulCertificate"
func (c *Client) ModifyListener(request *ModifyListenerRequest) (response *ModifyListenerResponse, err error) {
    return c.ModifyListenerWithContext(context.Background(), request)
}

// ModifyListener
// Modify a listener
//
// error code that may be returned:
//  INVALIDPARAMETER_APPLICATIONLAYERLISTENERCANNOTCARRYPARAMETERS = "InvalidParameter.ApplicationLayerListenerCannotCarryParameters"
//  INVALIDPARAMETER_HTTPLISTENERCANNOTCARRYPARAMETERS = "InvalidParameter.HttpListenerCannotCarryParameters"
//  INVALIDPARAMETER_HTTPSLISTENERCANCARRYPARAMETERS = "InvalidParameter.HttpsListenerCanCarryParameters"
//  INVALIDPARAMETER_INPUTMUSTEXISTONE = "InvalidParameter.InputMustExistOne"
//  INVALIDPARAMETER_TCPLISTENERCARRYPARAMETERS = "InvalidParameter.TcpListenerCarryParameters"
//  INVALIDPARAMETER_TRANSPORTLAYERLISTENERCANNOTCARRYPARAMETERS = "InvalidParameter.TransportLayerListenerCannotCarryParameters"
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  INVALIDPARAMETERVALUE_INSTANCENAME = "InvalidParameterValue.InstanceName"
//  INVALIDPARAMETERVALUE_LENGTHRANGE = "InvalidParameterValue.LengthRange"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_CERTIFICATENOTCA = "UnsupportedOperation.CertificateNotCa"
//  UNSUPPORTEDOPERATION_CURRENTACCOUNTNOTALLOWED = "UnsupportedOperation.CurrentAccountNotAllowed"
//  UNSUPPORTEDOPERATION_HTTPVERSIONINCONSISTENT = "UnsupportedOperation.HttpVersionInconsistent"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
//  UNSUPPORTEDOPERATION_ONLYTCPLISTENERSUPPORTTOA = "UnsupportedOperation.OnlyTcpListenerSupportToa"
//  UNSUPPORTEDOPERATION_TOAFEATURENOTENABLED = "UnsupportedOperation.TOAFeatureNotEnabled"
//  UNSUPPORTEDOPERATION_UDPLISTENERNOTSUPPORTPROXYPROTOCOL = "UnsupportedOperation.UdpListenerNotSupportProxyProtocol"
//  UNSUPPORTEDOPERATION_UNLAWFULCERTIFICATE = "UnsupportedOperation.UnlawfulCertificate"
func (c *Client) ModifyListenerWithContext(ctx context.Context, request *ModifyListenerRequest) (response *ModifyListenerResponse, err error) {
    if request == nil {
        request = NewModifyListenerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "ModifyListener")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyListener require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyListenerResponse()
    err = c.Send(request, response)
    return
}

func NewReplaceListenerAdditionalCertRequest() (request *ReplaceListenerAdditionalCertRequest) {
    request = &ReplaceListenerAdditionalCertRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "ReplaceListenerAdditionalCert")
    
    
    return
}

func NewReplaceListenerAdditionalCertResponse() (response *ReplaceListenerAdditionalCertResponse) {
    response = &ReplaceListenerAdditionalCertResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReplaceListenerAdditionalCert
// Replace the extension certificate.
//
// error code that may be returned:
//  INVALIDPARAMETER_INPUTREPEAT = "InvalidParameter.InputRepeat"
//  INVALIDPARAMETERVALUE_CERTIFICATESCONFLICT = "InvalidParameterValue.CertificatesConflict"
//  UNSUPPORTEDOPERATION_DOMAINMISMATCHED = "UnsupportedOperation.DomainMismatched"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
func (c *Client) ReplaceListenerAdditionalCert(request *ReplaceListenerAdditionalCertRequest) (response *ReplaceListenerAdditionalCertResponse, err error) {
    return c.ReplaceListenerAdditionalCertWithContext(context.Background(), request)
}

// ReplaceListenerAdditionalCert
// Replace the extension certificate.
//
// error code that may be returned:
//  INVALIDPARAMETER_INPUTREPEAT = "InvalidParameter.InputRepeat"
//  INVALIDPARAMETERVALUE_CERTIFICATESCONFLICT = "InvalidParameterValue.CertificatesConflict"
//  UNSUPPORTEDOPERATION_DOMAINMISMATCHED = "UnsupportedOperation.DomainMismatched"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
func (c *Client) ReplaceListenerAdditionalCertWithContext(ctx context.Context, request *ReplaceListenerAdditionalCertRequest) (response *ReplaceListenerAdditionalCertResponse, err error) {
    if request == nil {
        request = NewReplaceListenerAdditionalCertRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "ReplaceListenerAdditionalCert")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReplaceListenerAdditionalCert require credential")
    }

    request.SetContext(ctx)
    
    response = NewReplaceListenerAdditionalCertResponse()
    err = c.Send(request, response)
    return
}
