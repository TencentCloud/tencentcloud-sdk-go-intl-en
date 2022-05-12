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

package v20180419

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2018-04-19"

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


func NewAttachInstancesRequest() (request *AttachInstancesRequest) {
    request = &AttachInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "AttachInstances")
    
    
    return
}

func NewAttachInstancesResponse() (response *AttachInstancesResponse) {
    response = &AttachInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AttachInstances
// This API (AttachInstances) is used to add CVM instances to an auto scaling group.
//
// error code that may be returned:
//  FAILEDOPERATION_NOACTIVITYTOGENERATE = "FailedOperation.NoActivityToGenerate"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceId"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  RESOURCEINSUFFICIENT_AUTOSCALINGGROUPABOVEMAXSIZE = "ResourceInsufficient.AutoScalingGroupAboveMaxSize"
//  RESOURCEINSUFFICIENT_INSERVICEINSTANCEABOVEMAXSIZE = "ResourceInsufficient.InServiceInstanceAboveMaxSize"
//  RESOURCENOTFOUND_AUTOSCALINGGROUPNOTFOUND = "ResourceNotFound.AutoScalingGroupNotFound"
//  RESOURCENOTFOUND_INSTANCESNOTFOUND = "ResourceNotFound.InstancesNotFound"
//  RESOURCEUNAVAILABLE_AUTOSCALINGGROUPINACTIVITY = "ResourceUnavailable.AutoScalingGroupInActivity"
//  RESOURCEUNAVAILABLE_CVMVPCINCONSISTENT = "ResourceUnavailable.CvmVpcInconsistent"
//  RESOURCEUNAVAILABLE_INSTANCEINOPERATION = "ResourceUnavailable.InstanceInOperation"
//  RESOURCEUNAVAILABLE_INSTANCESALREADYINAUTOSCALINGGROUP = "ResourceUnavailable.InstancesAlreadyInAutoScalingGroup"
//  RESOURCEUNAVAILABLE_LOADBALANCERINOPERATION = "ResourceUnavailable.LoadBalancerInOperation"
func (c *Client) AttachInstances(request *AttachInstancesRequest) (response *AttachInstancesResponse, err error) {
    return c.AttachInstancesWithContext(context.Background(), request)
}

// AttachInstances
// This API (AttachInstances) is used to add CVM instances to an auto scaling group.
//
// error code that may be returned:
//  FAILEDOPERATION_NOACTIVITYTOGENERATE = "FailedOperation.NoActivityToGenerate"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceId"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  RESOURCEINSUFFICIENT_AUTOSCALINGGROUPABOVEMAXSIZE = "ResourceInsufficient.AutoScalingGroupAboveMaxSize"
//  RESOURCEINSUFFICIENT_INSERVICEINSTANCEABOVEMAXSIZE = "ResourceInsufficient.InServiceInstanceAboveMaxSize"
//  RESOURCENOTFOUND_AUTOSCALINGGROUPNOTFOUND = "ResourceNotFound.AutoScalingGroupNotFound"
//  RESOURCENOTFOUND_INSTANCESNOTFOUND = "ResourceNotFound.InstancesNotFound"
//  RESOURCEUNAVAILABLE_AUTOSCALINGGROUPINACTIVITY = "ResourceUnavailable.AutoScalingGroupInActivity"
//  RESOURCEUNAVAILABLE_CVMVPCINCONSISTENT = "ResourceUnavailable.CvmVpcInconsistent"
//  RESOURCEUNAVAILABLE_INSTANCEINOPERATION = "ResourceUnavailable.InstanceInOperation"
//  RESOURCEUNAVAILABLE_INSTANCESALREADYINAUTOSCALINGGROUP = "ResourceUnavailable.InstancesAlreadyInAutoScalingGroup"
//  RESOURCEUNAVAILABLE_LOADBALANCERINOPERATION = "ResourceUnavailable.LoadBalancerInOperation"
func (c *Client) AttachInstancesWithContext(ctx context.Context, request *AttachInstancesRequest) (response *AttachInstancesResponse, err error) {
    if request == nil {
        request = NewAttachInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AttachInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewAttachInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewAttachLoadBalancersRequest() (request *AttachLoadBalancersRequest) {
    request = &AttachLoadBalancersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "AttachLoadBalancers")
    
    
    return
}

func NewAttachLoadBalancersResponse() (response *AttachLoadBalancersResponse) {
    response = &AttachLoadBalancersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AttachLoadBalancers
// This API is used to add CLBs to a security group.
//
// error code that may be returned:
//  FAILEDOPERATION_NOACTIVITYTOGENERATE = "FailedOperation.NoActivityToGenerate"
//  INTERNALERROR_CALLLBERROR = "InternalError.CallLbError"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INVALIDPARAMETER_CONFLICT = "InvalidParameter.Conflict"
//  INVALIDPARAMETER_INSCENARIO = "InvalidParameter.InScenario"
//  INVALIDPARAMETER_MUSTONEPARAMETER = "InvalidParameter.MustOneParameter"
//  INVALIDPARAMETERVALUE_CLASSICLB = "InvalidParameterValue.ClassicLb"
//  INVALIDPARAMETERVALUE_DUPLICATEDFORWARDLB = "InvalidParameterValue.DuplicatedForwardLb"
//  INVALIDPARAMETERVALUE_FORWARDLB = "InvalidParameterValue.ForwardLb"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_INVALIDCLBREGION = "InvalidParameterValue.InvalidClbRegion"
//  INVALIDPARAMETERVALUE_TARGETPORTDUPLICATED = "InvalidParameterValue.TargetPortDuplicated"
//  LIMITEXCEEDED_AFTERATTACHLBLIMITEXCEEDED = "LimitExceeded.AfterAttachLbLimitExceeded"
//  MISSINGPARAMETER_INSCENARIO = "MissingParameter.InScenario"
//  RESOURCENOTFOUND_AUTOSCALINGGROUPNOTFOUND = "ResourceNotFound.AutoScalingGroupNotFound"
//  RESOURCENOTFOUND_LISTENERNOTFOUND = "ResourceNotFound.ListenerNotFound"
//  RESOURCENOTFOUND_LOADBALANCERNOTFOUND = "ResourceNotFound.LoadBalancerNotFound"
//  RESOURCENOTFOUND_LOCATIONNOTFOUND = "ResourceNotFound.LocationNotFound"
//  RESOURCEUNAVAILABLE_AUTOSCALINGGROUPINACTIVITY = "ResourceUnavailable.AutoScalingGroupInActivity"
//  RESOURCEUNAVAILABLE_LBBACKENDREGIONINCONSISTENT = "ResourceUnavailable.LbBackendRegionInconsistent"
//  RESOURCEUNAVAILABLE_LBPROJECTINCONSISTENT = "ResourceUnavailable.LbProjectInconsistent"
//  RESOURCEUNAVAILABLE_LBVPCINCONSISTENT = "ResourceUnavailable.LbVpcInconsistent"
//  RESOURCEUNAVAILABLE_LOADBALANCERINOPERATION = "ResourceUnavailable.LoadBalancerInOperation"
func (c *Client) AttachLoadBalancers(request *AttachLoadBalancersRequest) (response *AttachLoadBalancersResponse, err error) {
    return c.AttachLoadBalancersWithContext(context.Background(), request)
}

// AttachLoadBalancers
// This API is used to add CLBs to a security group.
//
// error code that may be returned:
//  FAILEDOPERATION_NOACTIVITYTOGENERATE = "FailedOperation.NoActivityToGenerate"
//  INTERNALERROR_CALLLBERROR = "InternalError.CallLbError"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INVALIDPARAMETER_CONFLICT = "InvalidParameter.Conflict"
//  INVALIDPARAMETER_INSCENARIO = "InvalidParameter.InScenario"
//  INVALIDPARAMETER_MUSTONEPARAMETER = "InvalidParameter.MustOneParameter"
//  INVALIDPARAMETERVALUE_CLASSICLB = "InvalidParameterValue.ClassicLb"
//  INVALIDPARAMETERVALUE_DUPLICATEDFORWARDLB = "InvalidParameterValue.DuplicatedForwardLb"
//  INVALIDPARAMETERVALUE_FORWARDLB = "InvalidParameterValue.ForwardLb"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_INVALIDCLBREGION = "InvalidParameterValue.InvalidClbRegion"
//  INVALIDPARAMETERVALUE_TARGETPORTDUPLICATED = "InvalidParameterValue.TargetPortDuplicated"
//  LIMITEXCEEDED_AFTERATTACHLBLIMITEXCEEDED = "LimitExceeded.AfterAttachLbLimitExceeded"
//  MISSINGPARAMETER_INSCENARIO = "MissingParameter.InScenario"
//  RESOURCENOTFOUND_AUTOSCALINGGROUPNOTFOUND = "ResourceNotFound.AutoScalingGroupNotFound"
//  RESOURCENOTFOUND_LISTENERNOTFOUND = "ResourceNotFound.ListenerNotFound"
//  RESOURCENOTFOUND_LOADBALANCERNOTFOUND = "ResourceNotFound.LoadBalancerNotFound"
//  RESOURCENOTFOUND_LOCATIONNOTFOUND = "ResourceNotFound.LocationNotFound"
//  RESOURCEUNAVAILABLE_AUTOSCALINGGROUPINACTIVITY = "ResourceUnavailable.AutoScalingGroupInActivity"
//  RESOURCEUNAVAILABLE_LBBACKENDREGIONINCONSISTENT = "ResourceUnavailable.LbBackendRegionInconsistent"
//  RESOURCEUNAVAILABLE_LBPROJECTINCONSISTENT = "ResourceUnavailable.LbProjectInconsistent"
//  RESOURCEUNAVAILABLE_LBVPCINCONSISTENT = "ResourceUnavailable.LbVpcInconsistent"
//  RESOURCEUNAVAILABLE_LOADBALANCERINOPERATION = "ResourceUnavailable.LoadBalancerInOperation"
func (c *Client) AttachLoadBalancersWithContext(ctx context.Context, request *AttachLoadBalancersRequest) (response *AttachLoadBalancersResponse, err error) {
    if request == nil {
        request = NewAttachLoadBalancersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AttachLoadBalancers require credential")
    }

    request.SetContext(ctx)
    
    response = NewAttachLoadBalancersResponse()
    err = c.Send(request, response)
    return
}

func NewClearLaunchConfigurationAttributesRequest() (request *ClearLaunchConfigurationAttributesRequest) {
    request = &ClearLaunchConfigurationAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "ClearLaunchConfigurationAttributes")
    
    
    return
}

func NewClearLaunchConfigurationAttributesResponse() (response *ClearLaunchConfigurationAttributesResponse) {
    response = &ClearLaunchConfigurationAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ClearLaunchConfigurationAttributes
// This API is used to clear specific attributes of the launch configuration.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDLAUNCHCONFIGURATIONID = "InvalidParameterValue.InvalidLaunchConfigurationId"
func (c *Client) ClearLaunchConfigurationAttributes(request *ClearLaunchConfigurationAttributesRequest) (response *ClearLaunchConfigurationAttributesResponse, err error) {
    return c.ClearLaunchConfigurationAttributesWithContext(context.Background(), request)
}

// ClearLaunchConfigurationAttributes
// This API is used to clear specific attributes of the launch configuration.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDLAUNCHCONFIGURATIONID = "InvalidParameterValue.InvalidLaunchConfigurationId"
func (c *Client) ClearLaunchConfigurationAttributesWithContext(ctx context.Context, request *ClearLaunchConfigurationAttributesRequest) (response *ClearLaunchConfigurationAttributesResponse, err error) {
    if request == nil {
        request = NewClearLaunchConfigurationAttributesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ClearLaunchConfigurationAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewClearLaunchConfigurationAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewCompleteLifecycleActionRequest() (request *CompleteLifecycleActionRequest) {
    request = &CompleteLifecycleActionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "CompleteLifecycleAction")
    
    
    return
}

func NewCompleteLifecycleActionResponse() (response *CompleteLifecycleActionResponse) {
    response = &CompleteLifecycleActionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CompleteLifecycleAction
// This API (CompleteLifecycleAction) is used to complete a lifecycle action.
//
// 
//
// * The result ("CONTINUE" or "ABANDON") of a specific lifecycle hook can be specified by calling this API. If this API is not called at all, the lifecycle hook will be processed based on the "DefaultResult" after timeout.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MUSTONEPARAMETER = "InvalidParameter.MustOneParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceId"
//  INVALIDPARAMETERVALUE_INVALIDLIFECYCLEHOOKID = "InvalidParameterValue.InvalidLifecycleHookId"
//  RESOURCENOTFOUND_LIFECYCLEHOOKINSTANCENOTFOUND = "ResourceNotFound.LifecycleHookInstanceNotFound"
//  RESOURCENOTFOUND_LIFECYCLEHOOKNOTFOUND = "ResourceNotFound.LifecycleHookNotFound"
//  RESOURCEUNAVAILABLE_LIFECYCLEACTIONRESULTHASSET = "ResourceUnavailable.LifecycleActionResultHasSet"
func (c *Client) CompleteLifecycleAction(request *CompleteLifecycleActionRequest) (response *CompleteLifecycleActionResponse, err error) {
    return c.CompleteLifecycleActionWithContext(context.Background(), request)
}

// CompleteLifecycleAction
// This API (CompleteLifecycleAction) is used to complete a lifecycle action.
//
// 
//
// * The result ("CONTINUE" or "ABANDON") of a specific lifecycle hook can be specified by calling this API. If this API is not called at all, the lifecycle hook will be processed based on the "DefaultResult" after timeout.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MUSTONEPARAMETER = "InvalidParameter.MustOneParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceId"
//  INVALIDPARAMETERVALUE_INVALIDLIFECYCLEHOOKID = "InvalidParameterValue.InvalidLifecycleHookId"
//  RESOURCENOTFOUND_LIFECYCLEHOOKINSTANCENOTFOUND = "ResourceNotFound.LifecycleHookInstanceNotFound"
//  RESOURCENOTFOUND_LIFECYCLEHOOKNOTFOUND = "ResourceNotFound.LifecycleHookNotFound"
//  RESOURCEUNAVAILABLE_LIFECYCLEACTIONRESULTHASSET = "ResourceUnavailable.LifecycleActionResultHasSet"
func (c *Client) CompleteLifecycleActionWithContext(ctx context.Context, request *CompleteLifecycleActionRequest) (response *CompleteLifecycleActionResponse, err error) {
    if request == nil {
        request = NewCompleteLifecycleActionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CompleteLifecycleAction require credential")
    }

    request.SetContext(ctx)
    
    response = NewCompleteLifecycleActionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAutoScalingGroupRequest() (request *CreateAutoScalingGroupRequest) {
    request = &CreateAutoScalingGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "CreateAutoScalingGroup")
    
    
    return
}

func NewCreateAutoScalingGroupResponse() (response *CreateAutoScalingGroupResponse) {
    response = &CreateAutoScalingGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAutoScalingGroup
// This API (CreateAutoScalingGroup) is used to create an auto scaling group.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLLBERROR = "InternalError.CallLbError"
//  INTERNALERROR_CALLTAGERROR = "InternalError.CallTagError"
//  INTERNALERROR_CALLTVPCERROR = "InternalError.CallTvpcError"
//  INTERNALERROR_CALLVPCERROR = "InternalError.CallVpcError"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INVALIDPARAMETER_INSCENARIO = "InvalidParameter.InScenario"
//  INVALIDPARAMETERVALUE_BASECAPACITYTOOLARGE = "InvalidParameterValue.BaseCapacityTooLarge"
//  INVALIDPARAMETERVALUE_CLASSICLB = "InvalidParameterValue.ClassicLb"
//  INVALIDPARAMETERVALUE_CVMERROR = "InvalidParameterValue.CvmError"
//  INVALIDPARAMETERVALUE_DUPLICATEDFORWARDLB = "InvalidParameterValue.DuplicatedForwardLb"
//  INVALIDPARAMETERVALUE_DUPLICATEDSUBNET = "InvalidParameterValue.DuplicatedSubnet"
//  INVALIDPARAMETERVALUE_FORWARDLB = "InvalidParameterValue.ForwardLb"
//  INVALIDPARAMETERVALUE_GROUPNAMEDUPLICATED = "InvalidParameterValue.GroupNameDuplicated"
//  INVALIDPARAMETERVALUE_INVALIDCLBREGION = "InvalidParameterValue.InvalidClbRegion"
//  INVALIDPARAMETERVALUE_INVALIDLAUNCHCONFIGURATIONID = "InvalidParameterValue.InvalidLaunchConfigurationId"
//  INVALIDPARAMETERVALUE_INVALIDSUBNETID = "InvalidParameterValue.InvalidSubnetId"
//  INVALIDPARAMETERVALUE_LAUNCHCONFIGURATIONNOTFOUND = "InvalidParameterValue.LaunchConfigurationNotFound"
//  INVALIDPARAMETERVALUE_LBPROJECTINCONSISTENT = "InvalidParameterValue.LbProjectInconsistent"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_ONLYVPC = "InvalidParameterValue.OnlyVpc"
//  INVALIDPARAMETERVALUE_PROJECTIDNOTFOUND = "InvalidParameterValue.ProjectIdNotFound"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_SIZE = "InvalidParameterValue.Size"
//  INVALIDPARAMETERVALUE_SUBNETIDS = "InvalidParameterValue.SubnetIds"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_ZONEMISMATCHREGION = "InvalidParameterValue.ZoneMismatchRegion"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_AUTOSCALINGGROUPLIMITEXCEEDED = "LimitExceeded.AutoScalingGroupLimitExceeded"
//  LIMITEXCEEDED_MAXSIZELIMITEXCEEDED = "LimitExceeded.MaxSizeLimitExceeded"
//  LIMITEXCEEDED_MINSIZELIMITEXCEEDED = "LimitExceeded.MinSizeLimitExceeded"
//  MISSINGPARAMETER_INSCENARIO = "MissingParameter.InScenario"
//  RESOURCENOTFOUND_LISTENERNOTFOUND = "ResourceNotFound.ListenerNotFound"
//  RESOURCENOTFOUND_LOADBALANCERNOTFOUND = "ResourceNotFound.LoadBalancerNotFound"
//  RESOURCENOTFOUND_LOCATIONNOTFOUND = "ResourceNotFound.LocationNotFound"
//  RESOURCEUNAVAILABLE_LAUNCHCONFIGURATIONSTATUSABNORMAL = "ResourceUnavailable.LaunchConfigurationStatusAbnormal"
//  RESOURCEUNAVAILABLE_LBBACKENDREGIONINCONSISTENT = "ResourceUnavailable.LbBackendRegionInconsistent"
//  RESOURCEUNAVAILABLE_LBVPCINCONSISTENT = "ResourceUnavailable.LbVpcInconsistent"
//  RESOURCEUNAVAILABLE_PROJECTINCONSISTENT = "ResourceUnavailable.ProjectInconsistent"
//  RESOURCEUNAVAILABLE_ZONEUNAVAILABLE = "ResourceUnavailable.ZoneUnavailable"
func (c *Client) CreateAutoScalingGroup(request *CreateAutoScalingGroupRequest) (response *CreateAutoScalingGroupResponse, err error) {
    return c.CreateAutoScalingGroupWithContext(context.Background(), request)
}

// CreateAutoScalingGroup
// This API (CreateAutoScalingGroup) is used to create an auto scaling group.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLLBERROR = "InternalError.CallLbError"
//  INTERNALERROR_CALLTAGERROR = "InternalError.CallTagError"
//  INTERNALERROR_CALLTVPCERROR = "InternalError.CallTvpcError"
//  INTERNALERROR_CALLVPCERROR = "InternalError.CallVpcError"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INVALIDPARAMETER_INSCENARIO = "InvalidParameter.InScenario"
//  INVALIDPARAMETERVALUE_BASECAPACITYTOOLARGE = "InvalidParameterValue.BaseCapacityTooLarge"
//  INVALIDPARAMETERVALUE_CLASSICLB = "InvalidParameterValue.ClassicLb"
//  INVALIDPARAMETERVALUE_CVMERROR = "InvalidParameterValue.CvmError"
//  INVALIDPARAMETERVALUE_DUPLICATEDFORWARDLB = "InvalidParameterValue.DuplicatedForwardLb"
//  INVALIDPARAMETERVALUE_DUPLICATEDSUBNET = "InvalidParameterValue.DuplicatedSubnet"
//  INVALIDPARAMETERVALUE_FORWARDLB = "InvalidParameterValue.ForwardLb"
//  INVALIDPARAMETERVALUE_GROUPNAMEDUPLICATED = "InvalidParameterValue.GroupNameDuplicated"
//  INVALIDPARAMETERVALUE_INVALIDCLBREGION = "InvalidParameterValue.InvalidClbRegion"
//  INVALIDPARAMETERVALUE_INVALIDLAUNCHCONFIGURATIONID = "InvalidParameterValue.InvalidLaunchConfigurationId"
//  INVALIDPARAMETERVALUE_INVALIDSUBNETID = "InvalidParameterValue.InvalidSubnetId"
//  INVALIDPARAMETERVALUE_LAUNCHCONFIGURATIONNOTFOUND = "InvalidParameterValue.LaunchConfigurationNotFound"
//  INVALIDPARAMETERVALUE_LBPROJECTINCONSISTENT = "InvalidParameterValue.LbProjectInconsistent"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_ONLYVPC = "InvalidParameterValue.OnlyVpc"
//  INVALIDPARAMETERVALUE_PROJECTIDNOTFOUND = "InvalidParameterValue.ProjectIdNotFound"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_SIZE = "InvalidParameterValue.Size"
//  INVALIDPARAMETERVALUE_SUBNETIDS = "InvalidParameterValue.SubnetIds"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_ZONEMISMATCHREGION = "InvalidParameterValue.ZoneMismatchRegion"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_AUTOSCALINGGROUPLIMITEXCEEDED = "LimitExceeded.AutoScalingGroupLimitExceeded"
//  LIMITEXCEEDED_MAXSIZELIMITEXCEEDED = "LimitExceeded.MaxSizeLimitExceeded"
//  LIMITEXCEEDED_MINSIZELIMITEXCEEDED = "LimitExceeded.MinSizeLimitExceeded"
//  MISSINGPARAMETER_INSCENARIO = "MissingParameter.InScenario"
//  RESOURCENOTFOUND_LISTENERNOTFOUND = "ResourceNotFound.ListenerNotFound"
//  RESOURCENOTFOUND_LOADBALANCERNOTFOUND = "ResourceNotFound.LoadBalancerNotFound"
//  RESOURCENOTFOUND_LOCATIONNOTFOUND = "ResourceNotFound.LocationNotFound"
//  RESOURCEUNAVAILABLE_LAUNCHCONFIGURATIONSTATUSABNORMAL = "ResourceUnavailable.LaunchConfigurationStatusAbnormal"
//  RESOURCEUNAVAILABLE_LBBACKENDREGIONINCONSISTENT = "ResourceUnavailable.LbBackendRegionInconsistent"
//  RESOURCEUNAVAILABLE_LBVPCINCONSISTENT = "ResourceUnavailable.LbVpcInconsistent"
//  RESOURCEUNAVAILABLE_PROJECTINCONSISTENT = "ResourceUnavailable.ProjectInconsistent"
//  RESOURCEUNAVAILABLE_ZONEUNAVAILABLE = "ResourceUnavailable.ZoneUnavailable"
func (c *Client) CreateAutoScalingGroupWithContext(ctx context.Context, request *CreateAutoScalingGroupRequest) (response *CreateAutoScalingGroupResponse, err error) {
    if request == nil {
        request = NewCreateAutoScalingGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAutoScalingGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAutoScalingGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAutoScalingGroupFromInstanceRequest() (request *CreateAutoScalingGroupFromInstanceRequest) {
    request = &CreateAutoScalingGroupFromInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "CreateAutoScalingGroupFromInstance")
    
    
    return
}

func NewCreateAutoScalingGroupFromInstanceResponse() (response *CreateAutoScalingGroupFromInstanceResponse) {
    response = &CreateAutoScalingGroupFromInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAutoScalingGroupFromInstance
// This API is used to create launch configurations and scaling groups based on an instance.
//
// 
//
// Note: for a scaling group that is created based on a monthly-subscribed instance, the instances added for scale-out are pay-as-you-go instance.
//
// error code that may be returned:
//  ACCOUNTQUALIFICATIONRESTRICTIONS = "AccountQualificationRestrictions"
//  CALLCVMERROR = "CallCvmError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLVPCERROR = "InternalError.CallVpcError"
//  INVALIDPARAMETER_INSCENARIO = "InvalidParameter.InScenario"
//  INVALIDPARAMETERVALUE_CVMCONFIGURATIONERROR = "InvalidParameterValue.CvmConfigurationError"
//  INVALIDPARAMETERVALUE_CVMERROR = "InvalidParameterValue.CvmError"
//  INVALIDPARAMETERVALUE_DUPLICATEDSUBNET = "InvalidParameterValue.DuplicatedSubnet"
//  INVALIDPARAMETERVALUE_FORWARDLB = "InvalidParameterValue.ForwardLb"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceId"
//  INVALIDPARAMETERVALUE_LAUNCHCONFIGURATIONNAMEDUPLICATED = "InvalidParameterValue.LaunchConfigurationNameDuplicated"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_SIZE = "InvalidParameterValue.Size"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED_AUTOSCALINGGROUPLIMITEXCEEDED = "LimitExceeded.AutoScalingGroupLimitExceeded"
//  LIMITEXCEEDED_DESIREDCAPACITYLIMITEXCEEDED = "LimitExceeded.DesiredCapacityLimitExceeded"
//  LIMITEXCEEDED_MAXSIZELIMITEXCEEDED = "LimitExceeded.MaxSizeLimitExceeded"
//  LIMITEXCEEDED_MINSIZELIMITEXCEEDED = "LimitExceeded.MinSizeLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCESNOTFOUND = "ResourceNotFound.InstancesNotFound"
//  RESOURCEUNAVAILABLE_LAUNCHCONFIGURATIONSTATUSABNORMAL = "ResourceUnavailable.LaunchConfigurationStatusAbnormal"
//  RESOURCEUNAVAILABLE_PROJECTINCONSISTENT = "ResourceUnavailable.ProjectInconsistent"
//  RESOURCEUNAVAILABLE_STOPPEDINSTANCENOTALLOWATTACH = "ResourceUnavailable.StoppedInstanceNotAllowAttach"
func (c *Client) CreateAutoScalingGroupFromInstance(request *CreateAutoScalingGroupFromInstanceRequest) (response *CreateAutoScalingGroupFromInstanceResponse, err error) {
    return c.CreateAutoScalingGroupFromInstanceWithContext(context.Background(), request)
}

// CreateAutoScalingGroupFromInstance
// This API is used to create launch configurations and scaling groups based on an instance.
//
// 
//
// Note: for a scaling group that is created based on a monthly-subscribed instance, the instances added for scale-out are pay-as-you-go instance.
//
// error code that may be returned:
//  ACCOUNTQUALIFICATIONRESTRICTIONS = "AccountQualificationRestrictions"
//  CALLCVMERROR = "CallCvmError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLVPCERROR = "InternalError.CallVpcError"
//  INVALIDPARAMETER_INSCENARIO = "InvalidParameter.InScenario"
//  INVALIDPARAMETERVALUE_CVMCONFIGURATIONERROR = "InvalidParameterValue.CvmConfigurationError"
//  INVALIDPARAMETERVALUE_CVMERROR = "InvalidParameterValue.CvmError"
//  INVALIDPARAMETERVALUE_DUPLICATEDSUBNET = "InvalidParameterValue.DuplicatedSubnet"
//  INVALIDPARAMETERVALUE_FORWARDLB = "InvalidParameterValue.ForwardLb"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceId"
//  INVALIDPARAMETERVALUE_LAUNCHCONFIGURATIONNAMEDUPLICATED = "InvalidParameterValue.LaunchConfigurationNameDuplicated"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_SIZE = "InvalidParameterValue.Size"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED_AUTOSCALINGGROUPLIMITEXCEEDED = "LimitExceeded.AutoScalingGroupLimitExceeded"
//  LIMITEXCEEDED_DESIREDCAPACITYLIMITEXCEEDED = "LimitExceeded.DesiredCapacityLimitExceeded"
//  LIMITEXCEEDED_MAXSIZELIMITEXCEEDED = "LimitExceeded.MaxSizeLimitExceeded"
//  LIMITEXCEEDED_MINSIZELIMITEXCEEDED = "LimitExceeded.MinSizeLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCESNOTFOUND = "ResourceNotFound.InstancesNotFound"
//  RESOURCEUNAVAILABLE_LAUNCHCONFIGURATIONSTATUSABNORMAL = "ResourceUnavailable.LaunchConfigurationStatusAbnormal"
//  RESOURCEUNAVAILABLE_PROJECTINCONSISTENT = "ResourceUnavailable.ProjectInconsistent"
//  RESOURCEUNAVAILABLE_STOPPEDINSTANCENOTALLOWATTACH = "ResourceUnavailable.StoppedInstanceNotAllowAttach"
func (c *Client) CreateAutoScalingGroupFromInstanceWithContext(ctx context.Context, request *CreateAutoScalingGroupFromInstanceRequest) (response *CreateAutoScalingGroupFromInstanceResponse, err error) {
    if request == nil {
        request = NewCreateAutoScalingGroupFromInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAutoScalingGroupFromInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAutoScalingGroupFromInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLifecycleHookRequest() (request *CreateLifecycleHookRequest) {
    request = &CreateLifecycleHookRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "CreateLifecycleHook")
    
    
    return
}

func NewCreateLifecycleHookResponse() (response *CreateLifecycleHookResponse) {
    response = &CreateLifecycleHookResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateLifecycleHook
// This API (CreateLifeCycleHook) is used to create a lifecycle hook.
//
// 
//
// * You can configure message notifications in the following format for lifecycle hooks, which will be sent to your CMQ queue by AS:
//
// 
//
// ```
//
// {
//
// 	"Service": "Tencent Cloud Auto Scaling",
//
// 	"Time": "2019-03-14T10:15:11Z",
//
// 	"AppId": "1251783334",
//
// 	"ActivityId": "asa-fznnvrja",
//
// 	"AutoScalingGroupId": "asg-rrrrtttt",
//
// 	"LifecycleHookId": "ash-xxxxyyyy",
//
// 	"LifecycleHookName": "my-hook",
//
// 	"LifecycleActionToken": "3080e1c9-0efe-4dd7-ad3b-90cd6618298f",
//
// 	"InstanceId": "ins-aaaabbbb",
//
// 	"LifecycleTransition": "INSTANCE_LAUNCHING",
//
// 	"NotificationMetadata": ""
//
// }
//
// ```
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLEEERROR = "InternalError.CalleeError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONFLICT = "InvalidParameter.Conflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FILTER = "InvalidParameterValue.Filter"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_LIFECYCLEHOOKNAMEDUPLICATED = "InvalidParameterValue.LifecycleHookNameDuplicated"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  LIMITEXCEEDED_QUOTANOTENOUGH = "LimitExceeded.QuotaNotEnough"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_AUTOSCALINGGROUPNOTFOUND = "ResourceNotFound.AutoScalingGroupNotFound"
//  RESOURCENOTFOUND_CMQQUEUENOTFOUND = "ResourceNotFound.CmqQueueNotFound"
//  RESOURCENOTFOUND_TDMQCMQQUEUENOTFOUND = "ResourceNotFound.TDMQCMQQueueNotFound"
//  RESOURCENOTFOUND_TDMQCMQTOPICNOTFOUND = "ResourceNotFound.TDMQCMQTopicNotFound"
//  RESOURCEUNAVAILABLE_CMQTOPICHASNOSUBSCRIBER = "ResourceUnavailable.CmqTopicHasNoSubscriber"
//  RESOURCEUNAVAILABLE_TDMQCMQTOPICHASNOSUBSCRIBER = "ResourceUnavailable.TDMQCMQTopicHasNoSubscriber"
func (c *Client) CreateLifecycleHook(request *CreateLifecycleHookRequest) (response *CreateLifecycleHookResponse, err error) {
    return c.CreateLifecycleHookWithContext(context.Background(), request)
}

// CreateLifecycleHook
// This API (CreateLifeCycleHook) is used to create a lifecycle hook.
//
// 
//
// * You can configure message notifications in the following format for lifecycle hooks, which will be sent to your CMQ queue by AS:
//
// 
//
// ```
//
// {
//
// 	"Service": "Tencent Cloud Auto Scaling",
//
// 	"Time": "2019-03-14T10:15:11Z",
//
// 	"AppId": "1251783334",
//
// 	"ActivityId": "asa-fznnvrja",
//
// 	"AutoScalingGroupId": "asg-rrrrtttt",
//
// 	"LifecycleHookId": "ash-xxxxyyyy",
//
// 	"LifecycleHookName": "my-hook",
//
// 	"LifecycleActionToken": "3080e1c9-0efe-4dd7-ad3b-90cd6618298f",
//
// 	"InstanceId": "ins-aaaabbbb",
//
// 	"LifecycleTransition": "INSTANCE_LAUNCHING",
//
// 	"NotificationMetadata": ""
//
// }
//
// ```
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLEEERROR = "InternalError.CalleeError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONFLICT = "InvalidParameter.Conflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FILTER = "InvalidParameterValue.Filter"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_LIFECYCLEHOOKNAMEDUPLICATED = "InvalidParameterValue.LifecycleHookNameDuplicated"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  LIMITEXCEEDED_QUOTANOTENOUGH = "LimitExceeded.QuotaNotEnough"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_AUTOSCALINGGROUPNOTFOUND = "ResourceNotFound.AutoScalingGroupNotFound"
//  RESOURCENOTFOUND_CMQQUEUENOTFOUND = "ResourceNotFound.CmqQueueNotFound"
//  RESOURCENOTFOUND_TDMQCMQQUEUENOTFOUND = "ResourceNotFound.TDMQCMQQueueNotFound"
//  RESOURCENOTFOUND_TDMQCMQTOPICNOTFOUND = "ResourceNotFound.TDMQCMQTopicNotFound"
//  RESOURCEUNAVAILABLE_CMQTOPICHASNOSUBSCRIBER = "ResourceUnavailable.CmqTopicHasNoSubscriber"
//  RESOURCEUNAVAILABLE_TDMQCMQTOPICHASNOSUBSCRIBER = "ResourceUnavailable.TDMQCMQTopicHasNoSubscriber"
func (c *Client) CreateLifecycleHookWithContext(ctx context.Context, request *CreateLifecycleHookRequest) (response *CreateLifecycleHookResponse, err error) {
    if request == nil {
        request = NewCreateLifecycleHookRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLifecycleHook require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLifecycleHookResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNotificationConfigurationRequest() (request *CreateNotificationConfigurationRequest) {
    request = &CreateNotificationConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "CreateNotificationConfiguration")
    
    
    return
}

func NewCreateNotificationConfigurationResponse() (response *CreateNotificationConfigurationResponse) {
    response = &CreateNotificationConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateNotificationConfiguration
// This API is used to create a notification policy.
//
// When the notification is sent to a CMQ topic or queue, the following contents are included:
//
// ```
//
// {
//
//     "Service": "Tencent Cloud Auto Scaling",
//
//     "CreatedTime": "2021-10-11T10:15:11Z", // Activity creation time
//
//     "AppId": "100000000",
//
//     "ActivityId": "asa-fznnvrja", // Scaling activity ID
//
//     "AutoScalingGroupId": "asg-pc2oqu2z", // Scaling group ID
//
//     "ActivityType": "SCALE_OUT",  // Scaling activity type
//
//     "StatusCode": "SUCCESSFUL",   // Scaling activity result
//
//     "Description": "Activity was launched in response to a difference between desired capacity and actual capacity,
//
//     scale out 1 instance(s).", // Scaling activity description
//
//     "StartTime": "2021-10-11T10:15:11Z",  // Activity starting time
//
//     "EndTime": "2021-10-11T10:15:32Z",    // Activity ending time
//
//     "DetailedStatusMessageSet": [ // A collection of failed attempts during the scaling process (Failed attempts are allowed in a successful scaling activity)
//
//         {
//
//             "Code": "InvalidInstanceType",
//
//             "Zone": "ap-guangzhou-2",
//
//             "InstanceId": "",
//
//             "InstanceChargeType": "POSTPAID_BY_HOUR",
//
//             "SubnetId": "subnet-4t5mgeuu",
//
//             "Message": "The specified instance type `S5.LARGE8` is invalid in `subnet-4t5mgeuu`, `ap-guangzhou-2`.",
//
//             "InstanceType": "S5.LARGE8",
//
//         }
//
//     ]
//
// }
//
// ```
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_CONFLICTNOTIFICATIONTARGET = "InvalidParameterValue.ConflictNotificationTarget"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_INVALIDNOTIFICATIONUSERGROUPID = "InvalidParameterValue.InvalidNotificationUserGroupId"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_USERGROUPIDNOTFOUND = "InvalidParameterValue.UserGroupIdNotFound"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUOTANOTENOUGH = "LimitExceeded.QuotaNotEnough"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_AUTOSCALINGGROUPNOTFOUND = "ResourceNotFound.AutoScalingGroupNotFound"
//  RESOURCENOTFOUND_CMQQUEUENOTFOUND = "ResourceNotFound.CmqQueueNotFound"
//  RESOURCENOTFOUND_TDMQCMQQUEUENOTFOUND = "ResourceNotFound.TDMQCMQQueueNotFound"
//  RESOURCENOTFOUND_TDMQCMQTOPICNOTFOUND = "ResourceNotFound.TDMQCMQTopicNotFound"
//  RESOURCEUNAVAILABLE_CMQTOPICHASNOSUBSCRIBER = "ResourceUnavailable.CmqTopicHasNoSubscriber"
//  RESOURCEUNAVAILABLE_TDMQCMQTOPICHASNOSUBSCRIBER = "ResourceUnavailable.TDMQCMQTopicHasNoSubscriber"
func (c *Client) CreateNotificationConfiguration(request *CreateNotificationConfigurationRequest) (response *CreateNotificationConfigurationResponse, err error) {
    return c.CreateNotificationConfigurationWithContext(context.Background(), request)
}

// CreateNotificationConfiguration
// This API is used to create a notification policy.
//
// When the notification is sent to a CMQ topic or queue, the following contents are included:
//
// ```
//
// {
//
//     "Service": "Tencent Cloud Auto Scaling",
//
//     "CreatedTime": "2021-10-11T10:15:11Z", // Activity creation time
//
//     "AppId": "100000000",
//
//     "ActivityId": "asa-fznnvrja", // Scaling activity ID
//
//     "AutoScalingGroupId": "asg-pc2oqu2z", // Scaling group ID
//
//     "ActivityType": "SCALE_OUT",  // Scaling activity type
//
//     "StatusCode": "SUCCESSFUL",   // Scaling activity result
//
//     "Description": "Activity was launched in response to a difference between desired capacity and actual capacity,
//
//     scale out 1 instance(s).", // Scaling activity description
//
//     "StartTime": "2021-10-11T10:15:11Z",  // Activity starting time
//
//     "EndTime": "2021-10-11T10:15:32Z",    // Activity ending time
//
//     "DetailedStatusMessageSet": [ // A collection of failed attempts during the scaling process (Failed attempts are allowed in a successful scaling activity)
//
//         {
//
//             "Code": "InvalidInstanceType",
//
//             "Zone": "ap-guangzhou-2",
//
//             "InstanceId": "",
//
//             "InstanceChargeType": "POSTPAID_BY_HOUR",
//
//             "SubnetId": "subnet-4t5mgeuu",
//
//             "Message": "The specified instance type `S5.LARGE8` is invalid in `subnet-4t5mgeuu`, `ap-guangzhou-2`.",
//
//             "InstanceType": "S5.LARGE8",
//
//         }
//
//     ]
//
// }
//
// ```
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_CONFLICTNOTIFICATIONTARGET = "InvalidParameterValue.ConflictNotificationTarget"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_INVALIDNOTIFICATIONUSERGROUPID = "InvalidParameterValue.InvalidNotificationUserGroupId"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_USERGROUPIDNOTFOUND = "InvalidParameterValue.UserGroupIdNotFound"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUOTANOTENOUGH = "LimitExceeded.QuotaNotEnough"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_AUTOSCALINGGROUPNOTFOUND = "ResourceNotFound.AutoScalingGroupNotFound"
//  RESOURCENOTFOUND_CMQQUEUENOTFOUND = "ResourceNotFound.CmqQueueNotFound"
//  RESOURCENOTFOUND_TDMQCMQQUEUENOTFOUND = "ResourceNotFound.TDMQCMQQueueNotFound"
//  RESOURCENOTFOUND_TDMQCMQTOPICNOTFOUND = "ResourceNotFound.TDMQCMQTopicNotFound"
//  RESOURCEUNAVAILABLE_CMQTOPICHASNOSUBSCRIBER = "ResourceUnavailable.CmqTopicHasNoSubscriber"
//  RESOURCEUNAVAILABLE_TDMQCMQTOPICHASNOSUBSCRIBER = "ResourceUnavailable.TDMQCMQTopicHasNoSubscriber"
func (c *Client) CreateNotificationConfigurationWithContext(ctx context.Context, request *CreateNotificationConfigurationRequest) (response *CreateNotificationConfigurationResponse, err error) {
    if request == nil {
        request = NewCreateNotificationConfigurationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNotificationConfiguration require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateNotificationConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateScalingPolicyRequest() (request *CreateScalingPolicyRequest) {
    request = &CreateScalingPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "CreateScalingPolicy")
    
    
    return
}

func NewCreateScalingPolicyResponse() (response *CreateScalingPolicyResponse) {
    response = &CreateScalingPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateScalingPolicy
// This API (CreateScalingPolicy) is used to create an alarm trigger policy.
//
// error code that may be returned:
//  INTERNALERROR_CALLNOTIFICATIONERROR = "InternalError.CallNotificationError"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_INVALIDNOTIFICATIONUSERGROUPID = "InvalidParameterValue.InvalidNotificationUserGroupId"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_SCALINGPOLICYNAMEDUPLICATE = "InvalidParameterValue.ScalingPolicyNameDuplicate"
//  INVALIDPARAMETERVALUE_THRESHOLDOUTOFRANGE = "InvalidParameterValue.ThresholdOutOfRange"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_USERGROUPIDNOTFOUND = "InvalidParameterValue.UserGroupIdNotFound"
//  LIMITEXCEEDED_QUOTANOTENOUGH = "LimitExceeded.QuotaNotEnough"
//  RESOURCENOTFOUND_AUTOSCALINGGROUPNOTFOUND = "ResourceNotFound.AutoScalingGroupNotFound"
func (c *Client) CreateScalingPolicy(request *CreateScalingPolicyRequest) (response *CreateScalingPolicyResponse, err error) {
    return c.CreateScalingPolicyWithContext(context.Background(), request)
}

// CreateScalingPolicy
// This API (CreateScalingPolicy) is used to create an alarm trigger policy.
//
// error code that may be returned:
//  INTERNALERROR_CALLNOTIFICATIONERROR = "InternalError.CallNotificationError"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_INVALIDNOTIFICATIONUSERGROUPID = "InvalidParameterValue.InvalidNotificationUserGroupId"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_SCALINGPOLICYNAMEDUPLICATE = "InvalidParameterValue.ScalingPolicyNameDuplicate"
//  INVALIDPARAMETERVALUE_THRESHOLDOUTOFRANGE = "InvalidParameterValue.ThresholdOutOfRange"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_USERGROUPIDNOTFOUND = "InvalidParameterValue.UserGroupIdNotFound"
//  LIMITEXCEEDED_QUOTANOTENOUGH = "LimitExceeded.QuotaNotEnough"
//  RESOURCENOTFOUND_AUTOSCALINGGROUPNOTFOUND = "ResourceNotFound.AutoScalingGroupNotFound"
func (c *Client) CreateScalingPolicyWithContext(ctx context.Context, request *CreateScalingPolicyRequest) (response *CreateScalingPolicyResponse, err error) {
    if request == nil {
        request = NewCreateScalingPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateScalingPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateScalingPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateScheduledActionRequest() (request *CreateScheduledActionRequest) {
    request = &CreateScheduledActionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "CreateScheduledAction")
    
    
    return
}

func NewCreateScheduledActionResponse() (response *CreateScheduledActionResponse) {
    response = &CreateScheduledActionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateScheduledAction
// This API (CreateScheduledAction) is used to create a scheduled task.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_CRONEXPRESSIONILLEGAL = "InvalidParameterValue.CronExpressionIllegal"
//  INVALIDPARAMETERVALUE_ENDTIMEBEFORESTARTTIME = "InvalidParameterValue.EndTimeBeforeStartTime"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_INVALIDSCHEDULEDACTIONNAMEINCLUDEILLEGALCHAR = "InvalidParameterValue.InvalidScheduledActionNameIncludeIllegalChar"
//  INVALIDPARAMETERVALUE_SCHEDULEDACTIONNAMEDUPLICATE = "InvalidParameterValue.ScheduledActionNameDuplicate"
//  INVALIDPARAMETERVALUE_SIZE = "InvalidParameterValue.Size"
//  INVALIDPARAMETERVALUE_STARTTIMEBEFORECURRENTTIME = "InvalidParameterValue.StartTimeBeforeCurrentTime"
//  INVALIDPARAMETERVALUE_TIMEFORMAT = "InvalidParameterValue.TimeFormat"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED_DESIREDCAPACITYLIMITEXCEEDED = "LimitExceeded.DesiredCapacityLimitExceeded"
//  LIMITEXCEEDED_MAXSIZELIMITEXCEEDED = "LimitExceeded.MaxSizeLimitExceeded"
//  LIMITEXCEEDED_MINSIZELIMITEXCEEDED = "LimitExceeded.MinSizeLimitExceeded"
//  LIMITEXCEEDED_QUOTANOTENOUGH = "LimitExceeded.QuotaNotEnough"
//  LIMITEXCEEDED_SCHEDULEDACTIONLIMITEXCEEDED = "LimitExceeded.ScheduledActionLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_AUTOSCALINGGROUPNOTFOUND = "ResourceNotFound.AutoScalingGroupNotFound"
func (c *Client) CreateScheduledAction(request *CreateScheduledActionRequest) (response *CreateScheduledActionResponse, err error) {
    return c.CreateScheduledActionWithContext(context.Background(), request)
}

// CreateScheduledAction
// This API (CreateScheduledAction) is used to create a scheduled task.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_CRONEXPRESSIONILLEGAL = "InvalidParameterValue.CronExpressionIllegal"
//  INVALIDPARAMETERVALUE_ENDTIMEBEFORESTARTTIME = "InvalidParameterValue.EndTimeBeforeStartTime"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_INVALIDSCHEDULEDACTIONNAMEINCLUDEILLEGALCHAR = "InvalidParameterValue.InvalidScheduledActionNameIncludeIllegalChar"
//  INVALIDPARAMETERVALUE_SCHEDULEDACTIONNAMEDUPLICATE = "InvalidParameterValue.ScheduledActionNameDuplicate"
//  INVALIDPARAMETERVALUE_SIZE = "InvalidParameterValue.Size"
//  INVALIDPARAMETERVALUE_STARTTIMEBEFORECURRENTTIME = "InvalidParameterValue.StartTimeBeforeCurrentTime"
//  INVALIDPARAMETERVALUE_TIMEFORMAT = "InvalidParameterValue.TimeFormat"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED_DESIREDCAPACITYLIMITEXCEEDED = "LimitExceeded.DesiredCapacityLimitExceeded"
//  LIMITEXCEEDED_MAXSIZELIMITEXCEEDED = "LimitExceeded.MaxSizeLimitExceeded"
//  LIMITEXCEEDED_MINSIZELIMITEXCEEDED = "LimitExceeded.MinSizeLimitExceeded"
//  LIMITEXCEEDED_QUOTANOTENOUGH = "LimitExceeded.QuotaNotEnough"
//  LIMITEXCEEDED_SCHEDULEDACTIONLIMITEXCEEDED = "LimitExceeded.ScheduledActionLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_AUTOSCALINGGROUPNOTFOUND = "ResourceNotFound.AutoScalingGroupNotFound"
func (c *Client) CreateScheduledActionWithContext(ctx context.Context, request *CreateScheduledActionRequest) (response *CreateScheduledActionResponse, err error) {
    if request == nil {
        request = NewCreateScheduledActionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateScheduledAction require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateScheduledActionResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAutoScalingGroupRequest() (request *DeleteAutoScalingGroupRequest) {
    request = &DeleteAutoScalingGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "DeleteAutoScalingGroup")
    
    
    return
}

func NewDeleteAutoScalingGroupResponse() (response *DeleteAutoScalingGroupResponse) {
    response = &DeleteAutoScalingGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAutoScalingGroup
// This API (DeleteAutoScalingGroup) is used to delete the specified auto scaling group that has no instances and remains inactive.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLERROR = "InternalError.CallError"
//  INTERNALERROR_CALLMONITORERROR = "InternalError.CallMonitorError"
//  INTERNALERROR_CALLTAGERROR = "InternalError.CallTagError"
//  INVALIDPARAMETER_CONFLICT = "InvalidParameter.Conflict"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  RESOURCEINUSE_ACTIVITYINPROGRESS = "ResourceInUse.ActivityInProgress"
//  RESOURCEINUSE_INSTANCEINGROUP = "ResourceInUse.InstanceInGroup"
//  RESOURCENOTFOUND_AUTOSCALINGGROUPNOTFOUND = "ResourceNotFound.AutoScalingGroupNotFound"
func (c *Client) DeleteAutoScalingGroup(request *DeleteAutoScalingGroupRequest) (response *DeleteAutoScalingGroupResponse, err error) {
    return c.DeleteAutoScalingGroupWithContext(context.Background(), request)
}

// DeleteAutoScalingGroup
// This API (DeleteAutoScalingGroup) is used to delete the specified auto scaling group that has no instances and remains inactive.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLERROR = "InternalError.CallError"
//  INTERNALERROR_CALLMONITORERROR = "InternalError.CallMonitorError"
//  INTERNALERROR_CALLTAGERROR = "InternalError.CallTagError"
//  INVALIDPARAMETER_CONFLICT = "InvalidParameter.Conflict"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  RESOURCEINUSE_ACTIVITYINPROGRESS = "ResourceInUse.ActivityInProgress"
//  RESOURCEINUSE_INSTANCEINGROUP = "ResourceInUse.InstanceInGroup"
//  RESOURCENOTFOUND_AUTOSCALINGGROUPNOTFOUND = "ResourceNotFound.AutoScalingGroupNotFound"
func (c *Client) DeleteAutoScalingGroupWithContext(ctx context.Context, request *DeleteAutoScalingGroupRequest) (response *DeleteAutoScalingGroupResponse, err error) {
    if request == nil {
        request = NewDeleteAutoScalingGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAutoScalingGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAutoScalingGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLaunchConfigurationRequest() (request *DeleteLaunchConfigurationRequest) {
    request = &DeleteLaunchConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "DeleteLaunchConfiguration")
    
    
    return
}

func NewDeleteLaunchConfigurationResponse() (response *DeleteLaunchConfigurationResponse) {
    response = &DeleteLaunchConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLaunchConfiguration
// This API (DeleteLaunchConfiguration) is used to delete a launch configuration.
//
// 
//
// * If the launch configuration is active in an auto scaling group, it cannot be deleted.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDLAUNCHCONFIGURATIONID = "InvalidParameterValue.InvalidLaunchConfigurationId"
//  RESOURCEINUSE_LAUNCHCONFIGURATIONIDINUSE = "ResourceInUse.LaunchConfigurationIdInUse"
//  RESOURCENOTFOUND_LAUNCHCONFIGURATIONIDNOTFOUND = "ResourceNotFound.LaunchConfigurationIdNotFound"
func (c *Client) DeleteLaunchConfiguration(request *DeleteLaunchConfigurationRequest) (response *DeleteLaunchConfigurationResponse, err error) {
    return c.DeleteLaunchConfigurationWithContext(context.Background(), request)
}

// DeleteLaunchConfiguration
// This API (DeleteLaunchConfiguration) is used to delete a launch configuration.
//
// 
//
// * If the launch configuration is active in an auto scaling group, it cannot be deleted.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDLAUNCHCONFIGURATIONID = "InvalidParameterValue.InvalidLaunchConfigurationId"
//  RESOURCEINUSE_LAUNCHCONFIGURATIONIDINUSE = "ResourceInUse.LaunchConfigurationIdInUse"
//  RESOURCENOTFOUND_LAUNCHCONFIGURATIONIDNOTFOUND = "ResourceNotFound.LaunchConfigurationIdNotFound"
func (c *Client) DeleteLaunchConfigurationWithContext(ctx context.Context, request *DeleteLaunchConfigurationRequest) (response *DeleteLaunchConfigurationResponse, err error) {
    if request == nil {
        request = NewDeleteLaunchConfigurationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLaunchConfiguration require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLaunchConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLifecycleHookRequest() (request *DeleteLifecycleHookRequest) {
    request = &DeleteLifecycleHookRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "DeleteLifecycleHook")
    
    
    return
}

func NewDeleteLifecycleHookResponse() (response *DeleteLifecycleHookResponse) {
    response = &DeleteLifecycleHookResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLifecycleHook
// This API (DeleteLifeCycleHook) is used to delete a lifecycle hook.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDLIFECYCLEHOOKID = "InvalidParameterValue.InvalidLifecycleHookId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_LIFECYCLEHOOKNOTFOUND = "ResourceNotFound.LifecycleHookNotFound"
func (c *Client) DeleteLifecycleHook(request *DeleteLifecycleHookRequest) (response *DeleteLifecycleHookResponse, err error) {
    return c.DeleteLifecycleHookWithContext(context.Background(), request)
}

// DeleteLifecycleHook
// This API (DeleteLifeCycleHook) is used to delete a lifecycle hook.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDLIFECYCLEHOOKID = "InvalidParameterValue.InvalidLifecycleHookId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_LIFECYCLEHOOKNOTFOUND = "ResourceNotFound.LifecycleHookNotFound"
func (c *Client) DeleteLifecycleHookWithContext(ctx context.Context, request *DeleteLifecycleHookRequest) (response *DeleteLifecycleHookResponse, err error) {
    if request == nil {
        request = NewDeleteLifecycleHookRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLifecycleHook require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLifecycleHookResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNotificationConfigurationRequest() (request *DeleteNotificationConfigurationRequest) {
    request = &DeleteNotificationConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "DeleteNotificationConfiguration")
    
    
    return
}

func NewDeleteNotificationConfigurationResponse() (response *DeleteNotificationConfigurationResponse) {
    response = &DeleteNotificationConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteNotificationConfiguration
// This API (DeleteNotificationConfiguration) is used to delete the specified notification.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGNOTIFICATIONID = "InvalidParameterValue.InvalidAutoScalingNotificationId"
//  RESOURCENOTFOUND_AUTOSCALINGNOTIFICATIONNOTFOUND = "ResourceNotFound.AutoScalingNotificationNotFound"
func (c *Client) DeleteNotificationConfiguration(request *DeleteNotificationConfigurationRequest) (response *DeleteNotificationConfigurationResponse, err error) {
    return c.DeleteNotificationConfigurationWithContext(context.Background(), request)
}

// DeleteNotificationConfiguration
// This API (DeleteNotificationConfiguration) is used to delete the specified notification.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGNOTIFICATIONID = "InvalidParameterValue.InvalidAutoScalingNotificationId"
//  RESOURCENOTFOUND_AUTOSCALINGNOTIFICATIONNOTFOUND = "ResourceNotFound.AutoScalingNotificationNotFound"
func (c *Client) DeleteNotificationConfigurationWithContext(ctx context.Context, request *DeleteNotificationConfigurationRequest) (response *DeleteNotificationConfigurationResponse, err error) {
    if request == nil {
        request = NewDeleteNotificationConfigurationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteNotificationConfiguration require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteNotificationConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteScalingPolicyRequest() (request *DeleteScalingPolicyRequest) {
    request = &DeleteScalingPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "DeleteScalingPolicy")
    
    
    return
}

func NewDeleteScalingPolicyResponse() (response *DeleteScalingPolicyResponse) {
    response = &DeleteScalingPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteScalingPolicy
// This API (DeleteScalingPolicy) is used to delete an alarm trigger policy.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGPOLICYID = "InvalidParameterValue.InvalidAutoScalingPolicyId"
//  RESOURCENOTFOUND_SCALINGPOLICYNOTFOUND = "ResourceNotFound.ScalingPolicyNotFound"
func (c *Client) DeleteScalingPolicy(request *DeleteScalingPolicyRequest) (response *DeleteScalingPolicyResponse, err error) {
    return c.DeleteScalingPolicyWithContext(context.Background(), request)
}

// DeleteScalingPolicy
// This API (DeleteScalingPolicy) is used to delete an alarm trigger policy.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGPOLICYID = "InvalidParameterValue.InvalidAutoScalingPolicyId"
//  RESOURCENOTFOUND_SCALINGPOLICYNOTFOUND = "ResourceNotFound.ScalingPolicyNotFound"
func (c *Client) DeleteScalingPolicyWithContext(ctx context.Context, request *DeleteScalingPolicyRequest) (response *DeleteScalingPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteScalingPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteScalingPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteScalingPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteScheduledActionRequest() (request *DeleteScheduledActionRequest) {
    request = &DeleteScheduledActionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "DeleteScheduledAction")
    
    
    return
}

func NewDeleteScheduledActionResponse() (response *DeleteScheduledActionResponse) {
    response = &DeleteScheduledActionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteScheduledAction
// This API (DeleteScheduledAction) is used to delete the specified scheduled task.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDSCHEDULEDACTIONID = "InvalidParameterValue.InvalidScheduledActionId"
//  RESOURCENOTFOUND_SCHEDULEDACTIONNOTFOUND = "ResourceNotFound.ScheduledActionNotFound"
func (c *Client) DeleteScheduledAction(request *DeleteScheduledActionRequest) (response *DeleteScheduledActionResponse, err error) {
    return c.DeleteScheduledActionWithContext(context.Background(), request)
}

// DeleteScheduledAction
// This API (DeleteScheduledAction) is used to delete the specified scheduled task.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDSCHEDULEDACTIONID = "InvalidParameterValue.InvalidScheduledActionId"
//  RESOURCENOTFOUND_SCHEDULEDACTIONNOTFOUND = "ResourceNotFound.ScheduledActionNotFound"
func (c *Client) DeleteScheduledActionWithContext(ctx context.Context, request *DeleteScheduledActionRequest) (response *DeleteScheduledActionResponse, err error) {
    if request == nil {
        request = NewDeleteScheduledActionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteScheduledAction require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteScheduledActionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountLimitsRequest() (request *DescribeAccountLimitsRequest) {
    request = &DescribeAccountLimitsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "DescribeAccountLimits")
    
    
    return
}

func NewDescribeAccountLimitsResponse() (response *DescribeAccountLimitsResponse) {
    response = &DescribeAccountLimitsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAccountLimits
// This API (DescribeAccountLimits) is used to query the limits of user's AS resources.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeAccountLimits(request *DescribeAccountLimitsRequest) (response *DescribeAccountLimitsResponse, err error) {
    return c.DescribeAccountLimitsWithContext(context.Background(), request)
}

// DescribeAccountLimits
// This API (DescribeAccountLimits) is used to query the limits of user's AS resources.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeAccountLimitsWithContext(ctx context.Context, request *DescribeAccountLimitsRequest) (response *DescribeAccountLimitsResponse, err error) {
    if request == nil {
        request = NewDescribeAccountLimitsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccountLimits require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccountLimitsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAutoScalingActivitiesRequest() (request *DescribeAutoScalingActivitiesRequest) {
    request = &DescribeAutoScalingActivitiesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "DescribeAutoScalingActivities")
    
    
    return
}

func NewDescribeAutoScalingActivitiesResponse() (response *DescribeAutoScalingActivitiesResponse) {
    response = &DescribeAutoScalingActivitiesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAutoScalingActivities
// This API (DescribeAutoScalingActivities) is used to query the activity history of an auto scaling group.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CONFLICT = "InvalidParameter.Conflict"
//  INVALIDPARAMETERVALUE_FILTER = "InvalidParameterValue.Filter"
//  INVALIDPARAMETERVALUE_INVALIDACTIVITYID = "InvalidParameterValue.InvalidActivityId"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  LIMITEXCEEDED_FILTERVALUESTOOLONG = "LimitExceeded.FilterValuesTooLong"
func (c *Client) DescribeAutoScalingActivities(request *DescribeAutoScalingActivitiesRequest) (response *DescribeAutoScalingActivitiesResponse, err error) {
    return c.DescribeAutoScalingActivitiesWithContext(context.Background(), request)
}

// DescribeAutoScalingActivities
// This API (DescribeAutoScalingActivities) is used to query the activity history of an auto scaling group.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CONFLICT = "InvalidParameter.Conflict"
//  INVALIDPARAMETERVALUE_FILTER = "InvalidParameterValue.Filter"
//  INVALIDPARAMETERVALUE_INVALIDACTIVITYID = "InvalidParameterValue.InvalidActivityId"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  LIMITEXCEEDED_FILTERVALUESTOOLONG = "LimitExceeded.FilterValuesTooLong"
func (c *Client) DescribeAutoScalingActivitiesWithContext(ctx context.Context, request *DescribeAutoScalingActivitiesRequest) (response *DescribeAutoScalingActivitiesResponse, err error) {
    if request == nil {
        request = NewDescribeAutoScalingActivitiesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAutoScalingActivities require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAutoScalingActivitiesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAutoScalingAdvicesRequest() (request *DescribeAutoScalingAdvicesRequest) {
    request = &DescribeAutoScalingAdvicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "DescribeAutoScalingAdvices")
    
    
    return
}

func NewDescribeAutoScalingAdvicesResponse() (response *DescribeAutoScalingAdvicesResponse) {
    response = &DescribeAutoScalingAdvicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAutoScalingAdvices
// This API is used to query suggestions for scaling group configurations.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
func (c *Client) DescribeAutoScalingAdvices(request *DescribeAutoScalingAdvicesRequest) (response *DescribeAutoScalingAdvicesResponse, err error) {
    return c.DescribeAutoScalingAdvicesWithContext(context.Background(), request)
}

// DescribeAutoScalingAdvices
// This API is used to query suggestions for scaling group configurations.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
func (c *Client) DescribeAutoScalingAdvicesWithContext(ctx context.Context, request *DescribeAutoScalingAdvicesRequest) (response *DescribeAutoScalingAdvicesResponse, err error) {
    if request == nil {
        request = NewDescribeAutoScalingAdvicesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAutoScalingAdvices require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAutoScalingAdvicesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAutoScalingGroupLastActivitiesRequest() (request *DescribeAutoScalingGroupLastActivitiesRequest) {
    request = &DescribeAutoScalingGroupLastActivitiesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "DescribeAutoScalingGroupLastActivities")
    
    
    return
}

func NewDescribeAutoScalingGroupLastActivitiesResponse() (response *DescribeAutoScalingGroupLastActivitiesResponse) {
    response = &DescribeAutoScalingGroupLastActivitiesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAutoScalingGroupLastActivities
// This API is used to query the latest activity history of an auto scaling group.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NORESOURCEPERMISSION = "InvalidParameterValue.NoResourcePermission"
func (c *Client) DescribeAutoScalingGroupLastActivities(request *DescribeAutoScalingGroupLastActivitiesRequest) (response *DescribeAutoScalingGroupLastActivitiesResponse, err error) {
    return c.DescribeAutoScalingGroupLastActivitiesWithContext(context.Background(), request)
}

// DescribeAutoScalingGroupLastActivities
// This API is used to query the latest activity history of an auto scaling group.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NORESOURCEPERMISSION = "InvalidParameterValue.NoResourcePermission"
func (c *Client) DescribeAutoScalingGroupLastActivitiesWithContext(ctx context.Context, request *DescribeAutoScalingGroupLastActivitiesRequest) (response *DescribeAutoScalingGroupLastActivitiesResponse, err error) {
    if request == nil {
        request = NewDescribeAutoScalingGroupLastActivitiesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAutoScalingGroupLastActivities require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAutoScalingGroupLastActivitiesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAutoScalingGroupsRequest() (request *DescribeAutoScalingGroupsRequest) {
    request = &DescribeAutoScalingGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "DescribeAutoScalingGroups")
    
    
    return
}

func NewDescribeAutoScalingGroupsResponse() (response *DescribeAutoScalingGroupsResponse) {
    response = &DescribeAutoScalingGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAutoScalingGroups
// This API (DescribeAutoScalingGroups) is used to query the information of auto scaling groups.
//
// 
//
// * You can query the details of auto scaling groups based on information such as auto scaling group ID, auto scaling group name, or launch configuration ID. For more information on filters, see `Filter`.
//
// * If the parameter is empty, a certain number (specified by `Limit` and 20 by default) of auto scaling groups of the current user will be returned.
//
// error code that may be returned:
//  INVALIDPARAMETER_CONFLICT = "InvalidParameter.Conflict"
//  INVALIDPARAMETERCONFLICT = "InvalidParameterConflict"
//  INVALIDPARAMETERVALUE_FILTER = "InvalidParameterValue.Filter"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDLAUNCHCONFIGURATIONID = "InvalidParameterValue.InvalidLaunchConfigurationId"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPERMISSION = "InvalidPermission"
func (c *Client) DescribeAutoScalingGroups(request *DescribeAutoScalingGroupsRequest) (response *DescribeAutoScalingGroupsResponse, err error) {
    return c.DescribeAutoScalingGroupsWithContext(context.Background(), request)
}

// DescribeAutoScalingGroups
// This API (DescribeAutoScalingGroups) is used to query the information of auto scaling groups.
//
// 
//
// * You can query the details of auto scaling groups based on information such as auto scaling group ID, auto scaling group name, or launch configuration ID. For more information on filters, see `Filter`.
//
// * If the parameter is empty, a certain number (specified by `Limit` and 20 by default) of auto scaling groups of the current user will be returned.
//
// error code that may be returned:
//  INVALIDPARAMETER_CONFLICT = "InvalidParameter.Conflict"
//  INVALIDPARAMETERCONFLICT = "InvalidParameterConflict"
//  INVALIDPARAMETERVALUE_FILTER = "InvalidParameterValue.Filter"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDLAUNCHCONFIGURATIONID = "InvalidParameterValue.InvalidLaunchConfigurationId"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPERMISSION = "InvalidPermission"
func (c *Client) DescribeAutoScalingGroupsWithContext(ctx context.Context, request *DescribeAutoScalingGroupsRequest) (response *DescribeAutoScalingGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeAutoScalingGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAutoScalingGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAutoScalingGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAutoScalingInstancesRequest() (request *DescribeAutoScalingInstancesRequest) {
    request = &DescribeAutoScalingInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "DescribeAutoScalingInstances")
    
    
    return
}

func NewDescribeAutoScalingInstancesResponse() (response *DescribeAutoScalingInstancesResponse) {
    response = &DescribeAutoScalingInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAutoScalingInstances
// This API (DescribeAutoScalingInstances) is used to query the information of instances associated with AS.
//
// 
//
// * You can query the details of instances based on information such as instance ID and auto scaling group ID. For more information on filters, see `Filter`.
//
// * If the parameter is empty, a certain number (specified by `Limit` and 20 by default) of instances of the current user will be returned.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CONFLICT = "InvalidParameter.Conflict"
//  INVALIDPARAMETERVALUE_FILTER = "InvalidParameterValue.Filter"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceId"
//  LIMITEXCEEDED_FILTERVALUESTOOLONG = "LimitExceeded.FilterValuesTooLong"
func (c *Client) DescribeAutoScalingInstances(request *DescribeAutoScalingInstancesRequest) (response *DescribeAutoScalingInstancesResponse, err error) {
    return c.DescribeAutoScalingInstancesWithContext(context.Background(), request)
}

// DescribeAutoScalingInstances
// This API (DescribeAutoScalingInstances) is used to query the information of instances associated with AS.
//
// 
//
// * You can query the details of instances based on information such as instance ID and auto scaling group ID. For more information on filters, see `Filter`.
//
// * If the parameter is empty, a certain number (specified by `Limit` and 20 by default) of instances of the current user will be returned.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CONFLICT = "InvalidParameter.Conflict"
//  INVALIDPARAMETERVALUE_FILTER = "InvalidParameterValue.Filter"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceId"
//  LIMITEXCEEDED_FILTERVALUESTOOLONG = "LimitExceeded.FilterValuesTooLong"
func (c *Client) DescribeAutoScalingInstancesWithContext(ctx context.Context, request *DescribeAutoScalingInstancesRequest) (response *DescribeAutoScalingInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeAutoScalingInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAutoScalingInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAutoScalingInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLifecycleHooksRequest() (request *DescribeLifecycleHooksRequest) {
    request = &DescribeLifecycleHooksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "DescribeLifecycleHooks")
    
    
    return
}

func NewDescribeLifecycleHooksResponse() (response *DescribeLifecycleHooksResponse) {
    response = &DescribeLifecycleHooksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLifecycleHooks
// This API (DescribeLifecycleHooks) is used to query the information of lifecycle hooks.
//
// 
//
// * You can query the details of lifecycle hooks based on information such as auto scaling group ID, lifecycle hook ID, or lifecycle hook name. For more information on filters, see `Filter`.
//
// * If the parameter is empty, a certain number (specified by `Limit` and 20 by default) of lifecycle hooks of the current user will be returned.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONFLICT = "InvalidParameter.Conflict"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDLIFECYCLEHOOKID = "InvalidParameterValue.InvalidLifecycleHookId"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeLifecycleHooks(request *DescribeLifecycleHooksRequest) (response *DescribeLifecycleHooksResponse, err error) {
    return c.DescribeLifecycleHooksWithContext(context.Background(), request)
}

// DescribeLifecycleHooks
// This API (DescribeLifecycleHooks) is used to query the information of lifecycle hooks.
//
// 
//
// * You can query the details of lifecycle hooks based on information such as auto scaling group ID, lifecycle hook ID, or lifecycle hook name. For more information on filters, see `Filter`.
//
// * If the parameter is empty, a certain number (specified by `Limit` and 20 by default) of lifecycle hooks of the current user will be returned.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONFLICT = "InvalidParameter.Conflict"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDLIFECYCLEHOOKID = "InvalidParameterValue.InvalidLifecycleHookId"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeLifecycleHooksWithContext(ctx context.Context, request *DescribeLifecycleHooksRequest) (response *DescribeLifecycleHooksResponse, err error) {
    if request == nil {
        request = NewDescribeLifecycleHooksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLifecycleHooks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLifecycleHooksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNotificationConfigurationsRequest() (request *DescribeNotificationConfigurationsRequest) {
    request = &DescribeNotificationConfigurationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "DescribeNotificationConfigurations")
    
    
    return
}

func NewDescribeNotificationConfigurationsResponse() (response *DescribeNotificationConfigurationsResponse) {
    response = &DescribeNotificationConfigurationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNotificationConfigurations
// This API (DescribeNotificationConfigurations) is used to query the information of one or more notifications.
//
// 
//
// You can query the details of notifications based on information such as notification ID and auto scaling group ID. For more information on filters, see `Filter`.
//
// If the parameter is empty, a certain number (specified by `Limit` and 20 by default) of notifications of the current user will be returned.
//
// error code that may be returned:
//  INVALIDPARAMETERCONFLICT = "InvalidParameterConflict"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGNOTIFICATIONID = "InvalidParameterValue.InvalidAutoScalingNotificationId"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
func (c *Client) DescribeNotificationConfigurations(request *DescribeNotificationConfigurationsRequest) (response *DescribeNotificationConfigurationsResponse, err error) {
    return c.DescribeNotificationConfigurationsWithContext(context.Background(), request)
}

// DescribeNotificationConfigurations
// This API (DescribeNotificationConfigurations) is used to query the information of one or more notifications.
//
// 
//
// You can query the details of notifications based on information such as notification ID and auto scaling group ID. For more information on filters, see `Filter`.
//
// If the parameter is empty, a certain number (specified by `Limit` and 20 by default) of notifications of the current user will be returned.
//
// error code that may be returned:
//  INVALIDPARAMETERCONFLICT = "InvalidParameterConflict"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGNOTIFICATIONID = "InvalidParameterValue.InvalidAutoScalingNotificationId"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
func (c *Client) DescribeNotificationConfigurationsWithContext(ctx context.Context, request *DescribeNotificationConfigurationsRequest) (response *DescribeNotificationConfigurationsResponse, err error) {
    if request == nil {
        request = NewDescribeNotificationConfigurationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNotificationConfigurations require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNotificationConfigurationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScalingPoliciesRequest() (request *DescribeScalingPoliciesRequest) {
    request = &DescribeScalingPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "DescribeScalingPolicies")
    
    
    return
}

func NewDescribeScalingPoliciesResponse() (response *DescribeScalingPoliciesResponse) {
    response = &DescribeScalingPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScalingPolicies
// This API (DescribeScalingPolicies) is used to query alarm trigger policies.
//
// error code that may be returned:
//  INTERNALERROR_CALLMONITORERROR = "InternalError.CallMonitorError"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INVALIDPARAMETER_CONFLICT = "InvalidParameter.Conflict"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGPOLICYID = "InvalidParameterValue.InvalidAutoScalingPolicyId"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DescribeScalingPolicies(request *DescribeScalingPoliciesRequest) (response *DescribeScalingPoliciesResponse, err error) {
    return c.DescribeScalingPoliciesWithContext(context.Background(), request)
}

// DescribeScalingPolicies
// This API (DescribeScalingPolicies) is used to query alarm trigger policies.
//
// error code that may be returned:
//  INTERNALERROR_CALLMONITORERROR = "InternalError.CallMonitorError"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INVALIDPARAMETER_CONFLICT = "InvalidParameter.Conflict"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGPOLICYID = "InvalidParameterValue.InvalidAutoScalingPolicyId"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DescribeScalingPoliciesWithContext(ctx context.Context, request *DescribeScalingPoliciesRequest) (response *DescribeScalingPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeScalingPoliciesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScalingPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScalingPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScheduledActionsRequest() (request *DescribeScheduledActionsRequest) {
    request = &DescribeScheduledActionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "DescribeScheduledActions")
    
    
    return
}

func NewDescribeScheduledActionsResponse() (response *DescribeScheduledActionsResponse) {
    response = &DescribeScheduledActionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScheduledActions
// This API (DescribeScheduledActions) is used to query the details of one or more scheduled tasks.
//
// 
//
// * You can query the details of scheduled tasks based on information such as scheduled task ID, scheduled task name, or auto scaling group ID. For more information on filters, see `Filter`.
//
// * If the parameter is empty, a certain number (specified by `Limit` and 20 by default) of scheduled tasks of the current user will be returned.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CONFLICT = "InvalidParameter.Conflict"
//  INVALIDPARAMETERVALUE_FILTER = "InvalidParameterValue.Filter"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDSCHEDULEDACTIONID = "InvalidParameterValue.InvalidScheduledActionId"
//  RESOURCENOTFOUND_SCHEDULEDACTIONNOTFOUND = "ResourceNotFound.ScheduledActionNotFound"
func (c *Client) DescribeScheduledActions(request *DescribeScheduledActionsRequest) (response *DescribeScheduledActionsResponse, err error) {
    return c.DescribeScheduledActionsWithContext(context.Background(), request)
}

// DescribeScheduledActions
// This API (DescribeScheduledActions) is used to query the details of one or more scheduled tasks.
//
// 
//
// * You can query the details of scheduled tasks based on information such as scheduled task ID, scheduled task name, or auto scaling group ID. For more information on filters, see `Filter`.
//
// * If the parameter is empty, a certain number (specified by `Limit` and 20 by default) of scheduled tasks of the current user will be returned.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CONFLICT = "InvalidParameter.Conflict"
//  INVALIDPARAMETERVALUE_FILTER = "InvalidParameterValue.Filter"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDSCHEDULEDACTIONID = "InvalidParameterValue.InvalidScheduledActionId"
//  RESOURCENOTFOUND_SCHEDULEDACTIONNOTFOUND = "ResourceNotFound.ScheduledActionNotFound"
func (c *Client) DescribeScheduledActionsWithContext(ctx context.Context, request *DescribeScheduledActionsRequest) (response *DescribeScheduledActionsResponse, err error) {
    if request == nil {
        request = NewDescribeScheduledActionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScheduledActions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScheduledActionsResponse()
    err = c.Send(request, response)
    return
}

func NewDetachInstancesRequest() (request *DetachInstancesRequest) {
    request = &DetachInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "DetachInstances")
    
    
    return
}

func NewDetachInstancesResponse() (response *DetachInstancesResponse) {
    response = &DetachInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DetachInstances
// This API is used to remove CVM instances from a scaling group. Instances removed via this API will not be terminated.
//
// * If the number of remaining `IN_SERVICE` instances in the scaling group is less than the minimum capacity, this API will return an error.
//
// * However, if the scaling group is in `DISABLED` status, the removal will not verify the relationship between the number of `IN_SERVICE` instances and the minimum capacity.
//
// * This removal will unassociate the CVM from the CLB instance that has been configured for the scaling group.
//
// error code that may be returned:
//  FAILEDOPERATION_NOACTIVITYTOGENERATE = "FailedOperation.NoActivityToGenerate"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceId"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  RESOURCEINSUFFICIENT_AUTOSCALINGGROUPBELOWMINSIZE = "ResourceInsufficient.AutoScalingGroupBelowMinSize"
//  RESOURCEINSUFFICIENT_INSERVICEINSTANCEBELOWMINSIZE = "ResourceInsufficient.InServiceInstanceBelowMinSize"
//  RESOURCENOTFOUND_AUTOSCALINGGROUPNOTFOUND = "ResourceNotFound.AutoScalingGroupNotFound"
//  RESOURCENOTFOUND_INSTANCESNOTINAUTOSCALINGGROUP = "ResourceNotFound.InstancesNotInAutoScalingGroup"
//  RESOURCEUNAVAILABLE_AUTOSCALINGGROUPINACTIVITY = "ResourceUnavailable.AutoScalingGroupInActivity"
//  RESOURCEUNAVAILABLE_INSTANCEINOPERATION = "ResourceUnavailable.InstanceInOperation"
//  RESOURCEUNAVAILABLE_LOADBALANCERINOPERATION = "ResourceUnavailable.LoadBalancerInOperation"
func (c *Client) DetachInstances(request *DetachInstancesRequest) (response *DetachInstancesResponse, err error) {
    return c.DetachInstancesWithContext(context.Background(), request)
}

// DetachInstances
// This API is used to remove CVM instances from a scaling group. Instances removed via this API will not be terminated.
//
// * If the number of remaining `IN_SERVICE` instances in the scaling group is less than the minimum capacity, this API will return an error.
//
// * However, if the scaling group is in `DISABLED` status, the removal will not verify the relationship between the number of `IN_SERVICE` instances and the minimum capacity.
//
// * This removal will unassociate the CVM from the CLB instance that has been configured for the scaling group.
//
// error code that may be returned:
//  FAILEDOPERATION_NOACTIVITYTOGENERATE = "FailedOperation.NoActivityToGenerate"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceId"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  RESOURCEINSUFFICIENT_AUTOSCALINGGROUPBELOWMINSIZE = "ResourceInsufficient.AutoScalingGroupBelowMinSize"
//  RESOURCEINSUFFICIENT_INSERVICEINSTANCEBELOWMINSIZE = "ResourceInsufficient.InServiceInstanceBelowMinSize"
//  RESOURCENOTFOUND_AUTOSCALINGGROUPNOTFOUND = "ResourceNotFound.AutoScalingGroupNotFound"
//  RESOURCENOTFOUND_INSTANCESNOTINAUTOSCALINGGROUP = "ResourceNotFound.InstancesNotInAutoScalingGroup"
//  RESOURCEUNAVAILABLE_AUTOSCALINGGROUPINACTIVITY = "ResourceUnavailable.AutoScalingGroupInActivity"
//  RESOURCEUNAVAILABLE_INSTANCEINOPERATION = "ResourceUnavailable.InstanceInOperation"
//  RESOURCEUNAVAILABLE_LOADBALANCERINOPERATION = "ResourceUnavailable.LoadBalancerInOperation"
func (c *Client) DetachInstancesWithContext(ctx context.Context, request *DetachInstancesRequest) (response *DetachInstancesResponse, err error) {
    if request == nil {
        request = NewDetachInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetachInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetachInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDetachLoadBalancersRequest() (request *DetachLoadBalancersRequest) {
    request = &DetachLoadBalancersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "DetachLoadBalancers")
    
    
    return
}

func NewDetachLoadBalancersResponse() (response *DetachLoadBalancersResponse) {
    response = &DetachLoadBalancersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DetachLoadBalancers
// This API is used to unbind one or more CLBs from a scaling group. This API will not terminate CLBs.
//
// error code that may be returned:
//  FAILEDOPERATION_NOACTIVITYTOGENERATE = "FailedOperation.NoActivityToGenerate"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INVALIDPARAMETER_CONFLICT = "InvalidParameter.Conflict"
//  INVALIDPARAMETER_INSCENARIO = "InvalidParameter.InScenario"
//  INVALIDPARAMETER_LOADBALANCERNOTINAUTOSCALINGGROUP = "InvalidParameter.LoadBalancerNotInAutoScalingGroup"
//  INVALIDPARAMETER_MUSTONEPARAMETER = "InvalidParameter.MustOneParameter"
//  INVALIDPARAMETERVALUE_CLASSICLB = "InvalidParameterValue.ClassicLb"
//  INVALIDPARAMETERVALUE_FORWARDLB = "InvalidParameterValue.ForwardLb"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_INVALIDCLBREGION = "InvalidParameterValue.InvalidClbRegion"
//  LIMITEXCEEDED_AFTERATTACHLBLIMITEXCEEDED = "LimitExceeded.AfterAttachLbLimitExceeded"
//  MISSINGPARAMETER_INSCENARIO = "MissingParameter.InScenario"
//  RESOURCENOTFOUND_AUTOSCALINGGROUPNOTFOUND = "ResourceNotFound.AutoScalingGroupNotFound"
//  RESOURCENOTFOUND_LISTENERNOTFOUND = "ResourceNotFound.ListenerNotFound"
//  RESOURCENOTFOUND_LOADBALANCERNOTFOUND = "ResourceNotFound.LoadBalancerNotFound"
//  RESOURCENOTFOUND_LOADBALANCERNOTINAUTOSCALINGGROUP = "ResourceNotFound.LoadBalancerNotInAutoScalingGroup"
//  RESOURCENOTFOUND_LOCATIONNOTFOUND = "ResourceNotFound.LocationNotFound"
//  RESOURCEUNAVAILABLE_AUTOSCALINGGROUPINACTIVITY = "ResourceUnavailable.AutoScalingGroupInActivity"
//  RESOURCEUNAVAILABLE_LBBACKENDREGIONINCONSISTENT = "ResourceUnavailable.LbBackendRegionInconsistent"
//  RESOURCEUNAVAILABLE_LBPROJECTINCONSISTENT = "ResourceUnavailable.LbProjectInconsistent"
//  RESOURCEUNAVAILABLE_LBVPCINCONSISTENT = "ResourceUnavailable.LbVpcInconsistent"
//  RESOURCEUNAVAILABLE_LOADBALANCERINOPERATION = "ResourceUnavailable.LoadBalancerInOperation"
func (c *Client) DetachLoadBalancers(request *DetachLoadBalancersRequest) (response *DetachLoadBalancersResponse, err error) {
    return c.DetachLoadBalancersWithContext(context.Background(), request)
}

// DetachLoadBalancers
// This API is used to unbind one or more CLBs from a scaling group. This API will not terminate CLBs.
//
// error code that may be returned:
//  FAILEDOPERATION_NOACTIVITYTOGENERATE = "FailedOperation.NoActivityToGenerate"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INVALIDPARAMETER_CONFLICT = "InvalidParameter.Conflict"
//  INVALIDPARAMETER_INSCENARIO = "InvalidParameter.InScenario"
//  INVALIDPARAMETER_LOADBALANCERNOTINAUTOSCALINGGROUP = "InvalidParameter.LoadBalancerNotInAutoScalingGroup"
//  INVALIDPARAMETER_MUSTONEPARAMETER = "InvalidParameter.MustOneParameter"
//  INVALIDPARAMETERVALUE_CLASSICLB = "InvalidParameterValue.ClassicLb"
//  INVALIDPARAMETERVALUE_FORWARDLB = "InvalidParameterValue.ForwardLb"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_INVALIDCLBREGION = "InvalidParameterValue.InvalidClbRegion"
//  LIMITEXCEEDED_AFTERATTACHLBLIMITEXCEEDED = "LimitExceeded.AfterAttachLbLimitExceeded"
//  MISSINGPARAMETER_INSCENARIO = "MissingParameter.InScenario"
//  RESOURCENOTFOUND_AUTOSCALINGGROUPNOTFOUND = "ResourceNotFound.AutoScalingGroupNotFound"
//  RESOURCENOTFOUND_LISTENERNOTFOUND = "ResourceNotFound.ListenerNotFound"
//  RESOURCENOTFOUND_LOADBALANCERNOTFOUND = "ResourceNotFound.LoadBalancerNotFound"
//  RESOURCENOTFOUND_LOADBALANCERNOTINAUTOSCALINGGROUP = "ResourceNotFound.LoadBalancerNotInAutoScalingGroup"
//  RESOURCENOTFOUND_LOCATIONNOTFOUND = "ResourceNotFound.LocationNotFound"
//  RESOURCEUNAVAILABLE_AUTOSCALINGGROUPINACTIVITY = "ResourceUnavailable.AutoScalingGroupInActivity"
//  RESOURCEUNAVAILABLE_LBBACKENDREGIONINCONSISTENT = "ResourceUnavailable.LbBackendRegionInconsistent"
//  RESOURCEUNAVAILABLE_LBPROJECTINCONSISTENT = "ResourceUnavailable.LbProjectInconsistent"
//  RESOURCEUNAVAILABLE_LBVPCINCONSISTENT = "ResourceUnavailable.LbVpcInconsistent"
//  RESOURCEUNAVAILABLE_LOADBALANCERINOPERATION = "ResourceUnavailable.LoadBalancerInOperation"
func (c *Client) DetachLoadBalancersWithContext(ctx context.Context, request *DetachLoadBalancersRequest) (response *DetachLoadBalancersResponse, err error) {
    if request == nil {
        request = NewDetachLoadBalancersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetachLoadBalancers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetachLoadBalancersResponse()
    err = c.Send(request, response)
    return
}

func NewEnableAutoScalingGroupRequest() (request *EnableAutoScalingGroupRequest) {
    request = &EnableAutoScalingGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "EnableAutoScalingGroup")
    
    
    return
}

func NewEnableAutoScalingGroupResponse() (response *EnableAutoScalingGroupResponse) {
    response = &EnableAutoScalingGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EnableAutoScalingGroup
// This API (EnableAutoScalingGroup) is used to enable the specified auto scaling group.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  RESOURCENOTFOUND_AUTOSCALINGGROUPNOTFOUND = "ResourceNotFound.AutoScalingGroupNotFound"
func (c *Client) EnableAutoScalingGroup(request *EnableAutoScalingGroupRequest) (response *EnableAutoScalingGroupResponse, err error) {
    return c.EnableAutoScalingGroupWithContext(context.Background(), request)
}

// EnableAutoScalingGroup
// This API (EnableAutoScalingGroup) is used to enable the specified auto scaling group.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  RESOURCENOTFOUND_AUTOSCALINGGROUPNOTFOUND = "ResourceNotFound.AutoScalingGroupNotFound"
func (c *Client) EnableAutoScalingGroupWithContext(ctx context.Context, request *EnableAutoScalingGroupRequest) (response *EnableAutoScalingGroupResponse, err error) {
    if request == nil {
        request = NewEnableAutoScalingGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableAutoScalingGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableAutoScalingGroupResponse()
    err = c.Send(request, response)
    return
}

func NewExecuteScalingPolicyRequest() (request *ExecuteScalingPolicyRequest) {
    request = &ExecuteScalingPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "ExecuteScalingPolicy")
    
    
    return
}

func NewExecuteScalingPolicyResponse() (response *ExecuteScalingPolicyResponse) {
    response = &ExecuteScalingPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExecuteScalingPolicy
// This API (ExecuteScalingPolicy) is used to execute a scaling policy.
//
// 
//
// * The scaling policy can be executed based on the scaling policy ID.
//
// * When the auto scaling group to which the scaling policy belongs is performing a scaling activity, the scaling policy will be rejected.
//
// error code that may be returned:
//  FAILEDOPERATION_NOACTIVITYTOGENERATE = "FailedOperation.NoActivityToGenerate"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGPOLICYID = "InvalidParameterValue.InvalidAutoScalingPolicyId"
//  RESOURCEINUSE_AUTOSCALINGGROUPNOTACTIVE = "ResourceInUse.AutoScalingGroupNotActive"
//  RESOURCENOTFOUND_AUTOSCALINGGROUPNOTFOUND = "ResourceNotFound.AutoScalingGroupNotFound"
//  RESOURCENOTFOUND_SCALINGPOLICYNOTFOUND = "ResourceNotFound.ScalingPolicyNotFound"
//  RESOURCEUNAVAILABLE_AUTOSCALINGGROUPABNORMALSTATUS = "ResourceUnavailable.AutoScalingGroupAbnormalStatus"
//  RESOURCEUNAVAILABLE_AUTOSCALINGGROUPINACTIVITY = "ResourceUnavailable.AutoScalingGroupInActivity"
func (c *Client) ExecuteScalingPolicy(request *ExecuteScalingPolicyRequest) (response *ExecuteScalingPolicyResponse, err error) {
    return c.ExecuteScalingPolicyWithContext(context.Background(), request)
}

// ExecuteScalingPolicy
// This API (ExecuteScalingPolicy) is used to execute a scaling policy.
//
// 
//
// * The scaling policy can be executed based on the scaling policy ID.
//
// * When the auto scaling group to which the scaling policy belongs is performing a scaling activity, the scaling policy will be rejected.
//
// error code that may be returned:
//  FAILEDOPERATION_NOACTIVITYTOGENERATE = "FailedOperation.NoActivityToGenerate"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGPOLICYID = "InvalidParameterValue.InvalidAutoScalingPolicyId"
//  RESOURCEINUSE_AUTOSCALINGGROUPNOTACTIVE = "ResourceInUse.AutoScalingGroupNotActive"
//  RESOURCENOTFOUND_AUTOSCALINGGROUPNOTFOUND = "ResourceNotFound.AutoScalingGroupNotFound"
//  RESOURCENOTFOUND_SCALINGPOLICYNOTFOUND = "ResourceNotFound.ScalingPolicyNotFound"
//  RESOURCEUNAVAILABLE_AUTOSCALINGGROUPABNORMALSTATUS = "ResourceUnavailable.AutoScalingGroupAbnormalStatus"
//  RESOURCEUNAVAILABLE_AUTOSCALINGGROUPINACTIVITY = "ResourceUnavailable.AutoScalingGroupInActivity"
func (c *Client) ExecuteScalingPolicyWithContext(ctx context.Context, request *ExecuteScalingPolicyRequest) (response *ExecuteScalingPolicyResponse, err error) {
    if request == nil {
        request = NewExecuteScalingPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExecuteScalingPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewExecuteScalingPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAutoScalingGroupRequest() (request *ModifyAutoScalingGroupRequest) {
    request = &ModifyAutoScalingGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "ModifyAutoScalingGroup")
    
    
    return
}

func NewModifyAutoScalingGroupResponse() (response *ModifyAutoScalingGroupResponse) {
    response = &ModifyAutoScalingGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAutoScalingGroup
// This API (ModifyAutoScalingGroup) is used to modify an auto scaling group.
//
// error code that may be returned:
//  INTERNALERROR_CALLVPCERROR = "InternalError.CallVpcError"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INVALIDPARAMETER_CONFLICT = "InvalidParameter.Conflict"
//  INVALIDPARAMETER_INSCENARIO = "InvalidParameter.InScenario"
//  INVALIDPARAMETERVALUE_BASECAPACITYTOOLARGE = "InvalidParameterValue.BaseCapacityTooLarge"
//  INVALIDPARAMETERVALUE_CVMERROR = "InvalidParameterValue.CvmError"
//  INVALIDPARAMETERVALUE_DUPLICATEDSUBNET = "InvalidParameterValue.DuplicatedSubnet"
//  INVALIDPARAMETERVALUE_GROUPNAMEDUPLICATED = "InvalidParameterValue.GroupNameDuplicated"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_INVALIDLAUNCHCONFIGURATIONID = "InvalidParameterValue.InvalidLaunchConfigurationId"
//  INVALIDPARAMETERVALUE_INVALIDSUBNETID = "InvalidParameterValue.InvalidSubnetId"
//  INVALIDPARAMETERVALUE_LAUNCHCONFIGURATIONNOTFOUND = "InvalidParameterValue.LaunchConfigurationNotFound"
//  INVALIDPARAMETERVALUE_LBPROJECTINCONSISTENT = "InvalidParameterValue.LbProjectInconsistent"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_ONLYVPC = "InvalidParameterValue.OnlyVpc"
//  INVALIDPARAMETERVALUE_PROJECTIDNOTFOUND = "InvalidParameterValue.ProjectIdNotFound"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_SIZE = "InvalidParameterValue.Size"
//  INVALIDPARAMETERVALUE_SUBNETIDS = "InvalidParameterValue.SubnetIds"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_MAXSIZELIMITEXCEEDED = "LimitExceeded.MaxSizeLimitExceeded"
//  LIMITEXCEEDED_MINSIZELIMITEXCEEDED = "LimitExceeded.MinSizeLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_INSCENARIO = "MissingParameter.InScenario"
//  RESOURCENOTFOUND_AUTOSCALINGGROUPNOTFOUND = "ResourceNotFound.AutoScalingGroupNotFound"
//  RESOURCEUNAVAILABLE_LAUNCHCONFIGURATIONSTATUSABNORMAL = "ResourceUnavailable.LaunchConfigurationStatusAbnormal"
//  RESOURCEUNAVAILABLE_PROJECTINCONSISTENT = "ResourceUnavailable.ProjectInconsistent"
func (c *Client) ModifyAutoScalingGroup(request *ModifyAutoScalingGroupRequest) (response *ModifyAutoScalingGroupResponse, err error) {
    return c.ModifyAutoScalingGroupWithContext(context.Background(), request)
}

// ModifyAutoScalingGroup
// This API (ModifyAutoScalingGroup) is used to modify an auto scaling group.
//
// error code that may be returned:
//  INTERNALERROR_CALLVPCERROR = "InternalError.CallVpcError"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INVALIDPARAMETER_CONFLICT = "InvalidParameter.Conflict"
//  INVALIDPARAMETER_INSCENARIO = "InvalidParameter.InScenario"
//  INVALIDPARAMETERVALUE_BASECAPACITYTOOLARGE = "InvalidParameterValue.BaseCapacityTooLarge"
//  INVALIDPARAMETERVALUE_CVMERROR = "InvalidParameterValue.CvmError"
//  INVALIDPARAMETERVALUE_DUPLICATEDSUBNET = "InvalidParameterValue.DuplicatedSubnet"
//  INVALIDPARAMETERVALUE_GROUPNAMEDUPLICATED = "InvalidParameterValue.GroupNameDuplicated"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_INVALIDLAUNCHCONFIGURATIONID = "InvalidParameterValue.InvalidLaunchConfigurationId"
//  INVALIDPARAMETERVALUE_INVALIDSUBNETID = "InvalidParameterValue.InvalidSubnetId"
//  INVALIDPARAMETERVALUE_LAUNCHCONFIGURATIONNOTFOUND = "InvalidParameterValue.LaunchConfigurationNotFound"
//  INVALIDPARAMETERVALUE_LBPROJECTINCONSISTENT = "InvalidParameterValue.LbProjectInconsistent"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_ONLYVPC = "InvalidParameterValue.OnlyVpc"
//  INVALIDPARAMETERVALUE_PROJECTIDNOTFOUND = "InvalidParameterValue.ProjectIdNotFound"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_SIZE = "InvalidParameterValue.Size"
//  INVALIDPARAMETERVALUE_SUBNETIDS = "InvalidParameterValue.SubnetIds"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_MAXSIZELIMITEXCEEDED = "LimitExceeded.MaxSizeLimitExceeded"
//  LIMITEXCEEDED_MINSIZELIMITEXCEEDED = "LimitExceeded.MinSizeLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_INSCENARIO = "MissingParameter.InScenario"
//  RESOURCENOTFOUND_AUTOSCALINGGROUPNOTFOUND = "ResourceNotFound.AutoScalingGroupNotFound"
//  RESOURCEUNAVAILABLE_LAUNCHCONFIGURATIONSTATUSABNORMAL = "ResourceUnavailable.LaunchConfigurationStatusAbnormal"
//  RESOURCEUNAVAILABLE_PROJECTINCONSISTENT = "ResourceUnavailable.ProjectInconsistent"
func (c *Client) ModifyAutoScalingGroupWithContext(ctx context.Context, request *ModifyAutoScalingGroupRequest) (response *ModifyAutoScalingGroupResponse, err error) {
    if request == nil {
        request = NewModifyAutoScalingGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAutoScalingGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAutoScalingGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDesiredCapacityRequest() (request *ModifyDesiredCapacityRequest) {
    request = &ModifyDesiredCapacityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "ModifyDesiredCapacity")
    
    
    return
}

func NewModifyDesiredCapacityResponse() (response *ModifyDesiredCapacityResponse) {
    response = &ModifyDesiredCapacityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDesiredCapacity
// This API (ModifyDesiredCapacity) is used to modify the desired number of instances in the specified auto scaling group.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_BASECAPACITYTOOLARGE = "InvalidParameterValue.BaseCapacityTooLarge"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_SIZE = "InvalidParameterValue.Size"
//  LIMITEXCEEDED_DESIREDCAPACITYLIMITEXCEEDED = "LimitExceeded.DesiredCapacityLimitExceeded"
//  RESOURCENOTFOUND_AUTOSCALINGGROUPNOTFOUND = "ResourceNotFound.AutoScalingGroupNotFound"
//  RESOURCEUNAVAILABLE_AUTOSCALINGGROUPABNORMALSTATUS = "ResourceUnavailable.AutoScalingGroupAbnormalStatus"
//  RESOURCEUNAVAILABLE_AUTOSCALINGGROUPDISABLED = "ResourceUnavailable.AutoScalingGroupDisabled"
func (c *Client) ModifyDesiredCapacity(request *ModifyDesiredCapacityRequest) (response *ModifyDesiredCapacityResponse, err error) {
    return c.ModifyDesiredCapacityWithContext(context.Background(), request)
}

// ModifyDesiredCapacity
// This API (ModifyDesiredCapacity) is used to modify the desired number of instances in the specified auto scaling group.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_BASECAPACITYTOOLARGE = "InvalidParameterValue.BaseCapacityTooLarge"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_SIZE = "InvalidParameterValue.Size"
//  LIMITEXCEEDED_DESIREDCAPACITYLIMITEXCEEDED = "LimitExceeded.DesiredCapacityLimitExceeded"
//  RESOURCENOTFOUND_AUTOSCALINGGROUPNOTFOUND = "ResourceNotFound.AutoScalingGroupNotFound"
//  RESOURCEUNAVAILABLE_AUTOSCALINGGROUPABNORMALSTATUS = "ResourceUnavailable.AutoScalingGroupAbnormalStatus"
//  RESOURCEUNAVAILABLE_AUTOSCALINGGROUPDISABLED = "ResourceUnavailable.AutoScalingGroupDisabled"
func (c *Client) ModifyDesiredCapacityWithContext(ctx context.Context, request *ModifyDesiredCapacityRequest) (response *ModifyDesiredCapacityResponse, err error) {
    if request == nil {
        request = NewModifyDesiredCapacityRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDesiredCapacity require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDesiredCapacityResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLaunchConfigurationAttributesRequest() (request *ModifyLaunchConfigurationAttributesRequest) {
    request = &ModifyLaunchConfigurationAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "ModifyLaunchConfigurationAttributes")
    
    
    return
}

func NewModifyLaunchConfigurationAttributesResponse() (response *ModifyLaunchConfigurationAttributesResponse) {
    response = &ModifyLaunchConfigurationAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLaunchConfigurationAttributes
// This API (ModifyLaunchConfigurationAttributes) is used to modify some attributes of a launch configuration.
//
// 
//
// * The changes of launch configuration do not affect the existing instances. New instances will be created based on the modified configuration.
//
// * This API supports modifying certain simple types of attributes.
//
// error code that may be returned:
//  INVALIDPARAMETER_CONFLICT = "InvalidParameter.Conflict"
//  INVALIDPARAMETER_HOSTNAMEUNAVAILABLE = "InvalidParameter.HostNameUnavailable"
//  INVALIDPARAMETER_INSCENARIO = "InvalidParameter.InScenario"
//  INVALIDPARAMETER_INVALIDCOMBINATION = "InvalidParameter.InvalidCombination"
//  INVALIDPARAMETER_PARAMETERMUSTBEDELETED = "InvalidParameter.ParameterMustBeDeleted"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CVMCONFIGURATIONERROR = "InvalidParameterValue.CvmConfigurationError"
//  INVALIDPARAMETERVALUE_HOSTNAMEILLEGAL = "InvalidParameterValue.HostNameIllegal"
//  INVALIDPARAMETERVALUE_IMAGENOTFOUND = "InvalidParameterValue.ImageNotFound"
//  INVALIDPARAMETERVALUE_INSTANCENAMEILLEGAL = "InvalidParameterValue.InstanceNameIllegal"
//  INVALIDPARAMETERVALUE_INSTANCETYPENOTSUPPORTED = "InvalidParameterValue.InstanceTypeNotSupported"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEID = "InvalidParameterValue.InvalidImageId"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCETYPE = "InvalidParameterValue.InvalidInstanceType"
//  INVALIDPARAMETERVALUE_INVALIDLAUNCHCONFIGURATIONID = "InvalidParameterValue.InvalidLaunchConfigurationId"
//  INVALIDPARAMETERVALUE_LAUNCHCONFIGURATIONNAMEDUPLICATED = "InvalidParameterValue.LaunchConfigurationNameDuplicated"
//  INVALIDPARAMETERVALUE_NOTSTRINGTYPEFLOAT = "InvalidParameterValue.NotStringTypeFloat"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_TOOSHORT = "InvalidParameterValue.TooShort"
//  INVALIDPARAMETERVALUE_USERDATAFORMATERROR = "InvalidParameterValue.UserDataFormatError"
//  INVALIDPARAMETERVALUE_USERDATASIZEEXCEEDED = "InvalidParameterValue.UserDataSizeExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_INSCENARIO = "MissingParameter.InScenario"
//  RESOURCENOTFOUND_LAUNCHCONFIGURATIONIDNOTFOUND = "ResourceNotFound.LaunchConfigurationIdNotFound"
func (c *Client) ModifyLaunchConfigurationAttributes(request *ModifyLaunchConfigurationAttributesRequest) (response *ModifyLaunchConfigurationAttributesResponse, err error) {
    return c.ModifyLaunchConfigurationAttributesWithContext(context.Background(), request)
}

// ModifyLaunchConfigurationAttributes
// This API (ModifyLaunchConfigurationAttributes) is used to modify some attributes of a launch configuration.
//
// 
//
// * The changes of launch configuration do not affect the existing instances. New instances will be created based on the modified configuration.
//
// * This API supports modifying certain simple types of attributes.
//
// error code that may be returned:
//  INVALIDPARAMETER_CONFLICT = "InvalidParameter.Conflict"
//  INVALIDPARAMETER_HOSTNAMEUNAVAILABLE = "InvalidParameter.HostNameUnavailable"
//  INVALIDPARAMETER_INSCENARIO = "InvalidParameter.InScenario"
//  INVALIDPARAMETER_INVALIDCOMBINATION = "InvalidParameter.InvalidCombination"
//  INVALIDPARAMETER_PARAMETERMUSTBEDELETED = "InvalidParameter.ParameterMustBeDeleted"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CVMCONFIGURATIONERROR = "InvalidParameterValue.CvmConfigurationError"
//  INVALIDPARAMETERVALUE_HOSTNAMEILLEGAL = "InvalidParameterValue.HostNameIllegal"
//  INVALIDPARAMETERVALUE_IMAGENOTFOUND = "InvalidParameterValue.ImageNotFound"
//  INVALIDPARAMETERVALUE_INSTANCENAMEILLEGAL = "InvalidParameterValue.InstanceNameIllegal"
//  INVALIDPARAMETERVALUE_INSTANCETYPENOTSUPPORTED = "InvalidParameterValue.InstanceTypeNotSupported"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEID = "InvalidParameterValue.InvalidImageId"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCETYPE = "InvalidParameterValue.InvalidInstanceType"
//  INVALIDPARAMETERVALUE_INVALIDLAUNCHCONFIGURATIONID = "InvalidParameterValue.InvalidLaunchConfigurationId"
//  INVALIDPARAMETERVALUE_LAUNCHCONFIGURATIONNAMEDUPLICATED = "InvalidParameterValue.LaunchConfigurationNameDuplicated"
//  INVALIDPARAMETERVALUE_NOTSTRINGTYPEFLOAT = "InvalidParameterValue.NotStringTypeFloat"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_TOOSHORT = "InvalidParameterValue.TooShort"
//  INVALIDPARAMETERVALUE_USERDATAFORMATERROR = "InvalidParameterValue.UserDataFormatError"
//  INVALIDPARAMETERVALUE_USERDATASIZEEXCEEDED = "InvalidParameterValue.UserDataSizeExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_INSCENARIO = "MissingParameter.InScenario"
//  RESOURCENOTFOUND_LAUNCHCONFIGURATIONIDNOTFOUND = "ResourceNotFound.LaunchConfigurationIdNotFound"
func (c *Client) ModifyLaunchConfigurationAttributesWithContext(ctx context.Context, request *ModifyLaunchConfigurationAttributesRequest) (response *ModifyLaunchConfigurationAttributesResponse, err error) {
    if request == nil {
        request = NewModifyLaunchConfigurationAttributesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLaunchConfigurationAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLaunchConfigurationAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLifecycleHookRequest() (request *ModifyLifecycleHookRequest) {
    request = &ModifyLifecycleHookRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "ModifyLifecycleHook")
    
    
    return
}

func NewModifyLifecycleHookResponse() (response *ModifyLifecycleHookResponse) {
    response = &ModifyLifecycleHookResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLifecycleHook
// This API is used to modify the lifecycle hook.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND_CMQQUEUENOTFOUND = "ResourceNotFound.CmqQueueNotFound"
//  RESOURCENOTFOUND_LIFECYCLEHOOKNOTFOUND = "ResourceNotFound.LifecycleHookNotFound"
//  RESOURCENOTFOUND_TDMQCMQQUEUENOTFOUND = "ResourceNotFound.TDMQCMQQueueNotFound"
//  RESOURCENOTFOUND_TDMQCMQTOPICNOTFOUND = "ResourceNotFound.TDMQCMQTopicNotFound"
//  RESOURCEUNAVAILABLE_TDMQCMQTOPICHASNOSUBSCRIBER = "ResourceUnavailable.TDMQCMQTopicHasNoSubscriber"
func (c *Client) ModifyLifecycleHook(request *ModifyLifecycleHookRequest) (response *ModifyLifecycleHookResponse, err error) {
    return c.ModifyLifecycleHookWithContext(context.Background(), request)
}

// ModifyLifecycleHook
// This API is used to modify the lifecycle hook.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND_CMQQUEUENOTFOUND = "ResourceNotFound.CmqQueueNotFound"
//  RESOURCENOTFOUND_LIFECYCLEHOOKNOTFOUND = "ResourceNotFound.LifecycleHookNotFound"
//  RESOURCENOTFOUND_TDMQCMQQUEUENOTFOUND = "ResourceNotFound.TDMQCMQQueueNotFound"
//  RESOURCENOTFOUND_TDMQCMQTOPICNOTFOUND = "ResourceNotFound.TDMQCMQTopicNotFound"
//  RESOURCEUNAVAILABLE_TDMQCMQTOPICHASNOSUBSCRIBER = "ResourceUnavailable.TDMQCMQTopicHasNoSubscriber"
func (c *Client) ModifyLifecycleHookWithContext(ctx context.Context, request *ModifyLifecycleHookRequest) (response *ModifyLifecycleHookResponse, err error) {
    if request == nil {
        request = NewModifyLifecycleHookRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLifecycleHook require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLifecycleHookResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLoadBalancerTargetAttributesRequest() (request *ModifyLoadBalancerTargetAttributesRequest) {
    request = &ModifyLoadBalancerTargetAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "ModifyLoadBalancerTargetAttributes")
    
    
    return
}

func NewModifyLoadBalancerTargetAttributesResponse() (response *ModifyLoadBalancerTargetAttributesResponse) {
    response = &ModifyLoadBalancerTargetAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLoadBalancerTargetAttributes
// This API is used to modify the target rule attributes of the CLB in the scaling group.
//
// error code that may be returned:
//  FAILEDOPERATION_NOACTIVITYTOGENERATE = "FailedOperation.NoActivityToGenerate"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INVALIDPARAMETER_CONFLICT = "InvalidParameter.Conflict"
//  INVALIDPARAMETER_INSCENARIO = "InvalidParameter.InScenario"
//  INVALIDPARAMETER_LOADBALANCERNOTINAUTOSCALINGGROUP = "InvalidParameter.LoadBalancerNotInAutoScalingGroup"
//  INVALIDPARAMETERVALUE_CLASSICLB = "InvalidParameterValue.ClassicLb"
//  INVALIDPARAMETERVALUE_FORWARDLB = "InvalidParameterValue.ForwardLb"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_INVALIDCLBREGION = "InvalidParameterValue.InvalidClbRegion"
//  INVALIDPARAMETERVALUE_TARGETPORTDUPLICATED = "InvalidParameterValue.TargetPortDuplicated"
//  LIMITEXCEEDED_AFTERATTACHLBLIMITEXCEEDED = "LimitExceeded.AfterAttachLbLimitExceeded"
//  MISSINGPARAMETER_INSCENARIO = "MissingParameter.InScenario"
//  RESOURCENOTFOUND_AUTOSCALINGGROUPNOTFOUND = "ResourceNotFound.AutoScalingGroupNotFound"
//  RESOURCENOTFOUND_LISTENERNOTFOUND = "ResourceNotFound.ListenerNotFound"
//  RESOURCENOTFOUND_LOADBALANCERNOTFOUND = "ResourceNotFound.LoadBalancerNotFound"
//  RESOURCENOTFOUND_LOADBALANCERNOTINAUTOSCALINGGROUP = "ResourceNotFound.LoadBalancerNotInAutoScalingGroup"
//  RESOURCENOTFOUND_LOCATIONNOTFOUND = "ResourceNotFound.LocationNotFound"
//  RESOURCEUNAVAILABLE_AUTOSCALINGGROUPINACTIVITY = "ResourceUnavailable.AutoScalingGroupInActivity"
//  RESOURCEUNAVAILABLE_LBBACKENDREGIONINCONSISTENT = "ResourceUnavailable.LbBackendRegionInconsistent"
//  RESOURCEUNAVAILABLE_LBPROJECTINCONSISTENT = "ResourceUnavailable.LbProjectInconsistent"
//  RESOURCEUNAVAILABLE_LBVPCINCONSISTENT = "ResourceUnavailable.LbVpcInconsistent"
//  RESOURCEUNAVAILABLE_LOADBALANCERINOPERATION = "ResourceUnavailable.LoadBalancerInOperation"
func (c *Client) ModifyLoadBalancerTargetAttributes(request *ModifyLoadBalancerTargetAttributesRequest) (response *ModifyLoadBalancerTargetAttributesResponse, err error) {
    return c.ModifyLoadBalancerTargetAttributesWithContext(context.Background(), request)
}

// ModifyLoadBalancerTargetAttributes
// This API is used to modify the target rule attributes of the CLB in the scaling group.
//
// error code that may be returned:
//  FAILEDOPERATION_NOACTIVITYTOGENERATE = "FailedOperation.NoActivityToGenerate"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INVALIDPARAMETER_CONFLICT = "InvalidParameter.Conflict"
//  INVALIDPARAMETER_INSCENARIO = "InvalidParameter.InScenario"
//  INVALIDPARAMETER_LOADBALANCERNOTINAUTOSCALINGGROUP = "InvalidParameter.LoadBalancerNotInAutoScalingGroup"
//  INVALIDPARAMETERVALUE_CLASSICLB = "InvalidParameterValue.ClassicLb"
//  INVALIDPARAMETERVALUE_FORWARDLB = "InvalidParameterValue.ForwardLb"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_INVALIDCLBREGION = "InvalidParameterValue.InvalidClbRegion"
//  INVALIDPARAMETERVALUE_TARGETPORTDUPLICATED = "InvalidParameterValue.TargetPortDuplicated"
//  LIMITEXCEEDED_AFTERATTACHLBLIMITEXCEEDED = "LimitExceeded.AfterAttachLbLimitExceeded"
//  MISSINGPARAMETER_INSCENARIO = "MissingParameter.InScenario"
//  RESOURCENOTFOUND_AUTOSCALINGGROUPNOTFOUND = "ResourceNotFound.AutoScalingGroupNotFound"
//  RESOURCENOTFOUND_LISTENERNOTFOUND = "ResourceNotFound.ListenerNotFound"
//  RESOURCENOTFOUND_LOADBALANCERNOTFOUND = "ResourceNotFound.LoadBalancerNotFound"
//  RESOURCENOTFOUND_LOADBALANCERNOTINAUTOSCALINGGROUP = "ResourceNotFound.LoadBalancerNotInAutoScalingGroup"
//  RESOURCENOTFOUND_LOCATIONNOTFOUND = "ResourceNotFound.LocationNotFound"
//  RESOURCEUNAVAILABLE_AUTOSCALINGGROUPINACTIVITY = "ResourceUnavailable.AutoScalingGroupInActivity"
//  RESOURCEUNAVAILABLE_LBBACKENDREGIONINCONSISTENT = "ResourceUnavailable.LbBackendRegionInconsistent"
//  RESOURCEUNAVAILABLE_LBPROJECTINCONSISTENT = "ResourceUnavailable.LbProjectInconsistent"
//  RESOURCEUNAVAILABLE_LBVPCINCONSISTENT = "ResourceUnavailable.LbVpcInconsistent"
//  RESOURCEUNAVAILABLE_LOADBALANCERINOPERATION = "ResourceUnavailable.LoadBalancerInOperation"
func (c *Client) ModifyLoadBalancerTargetAttributesWithContext(ctx context.Context, request *ModifyLoadBalancerTargetAttributesRequest) (response *ModifyLoadBalancerTargetAttributesResponse, err error) {
    if request == nil {
        request = NewModifyLoadBalancerTargetAttributesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLoadBalancerTargetAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLoadBalancerTargetAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLoadBalancersRequest() (request *ModifyLoadBalancersRequest) {
    request = &ModifyLoadBalancersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "ModifyLoadBalancers")
    
    
    return
}

func NewModifyLoadBalancersResponse() (response *ModifyLoadBalancersResponse) {
    response = &ModifyLoadBalancersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLoadBalancers
// This API is used to modify the cloud load balancers of a scaling group.
//
// 
//
// * This API can specify a new cloud load balancer configuration for the scaling group. The new configuration overwrites the original load balancer configuration.
//
// * To clear the cloud load balancer of the scaling group, specify only the scaling group ID but not the specific cloud load balancer.
//
// * This API modifies the cloud load balancer of the scaling group and generate a scaling activity to asynchronously modify the cloud load balancers of existing instances.
//
// error code that may be returned:
//  FAILEDOPERATION_NOACTIVITYTOGENERATE = "FailedOperation.NoActivityToGenerate"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INVALIDPARAMETER_CONFLICT = "InvalidParameter.Conflict"
//  INVALIDPARAMETER_INSCENARIO = "InvalidParameter.InScenario"
//  INVALIDPARAMETERVALUE_CLASSICLB = "InvalidParameterValue.ClassicLb"
//  INVALIDPARAMETERVALUE_DUPLICATEDFORWARDLB = "InvalidParameterValue.DuplicatedForwardLb"
//  INVALIDPARAMETERVALUE_FORWARDLB = "InvalidParameterValue.ForwardLb"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_INVALIDCLBREGION = "InvalidParameterValue.InvalidClbRegion"
//  INVALIDPARAMETERVALUE_LBPROJECTINCONSISTENT = "InvalidParameterValue.LbProjectInconsistent"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  MISSINGPARAMETER_INSCENARIO = "MissingParameter.InScenario"
//  RESOURCENOTFOUND_AUTOSCALINGGROUPNOTFOUND = "ResourceNotFound.AutoScalingGroupNotFound"
//  RESOURCENOTFOUND_LISTENERNOTFOUND = "ResourceNotFound.ListenerNotFound"
//  RESOURCENOTFOUND_LOADBALANCERNOTFOUND = "ResourceNotFound.LoadBalancerNotFound"
//  RESOURCENOTFOUND_LOCATIONNOTFOUND = "ResourceNotFound.LocationNotFound"
//  RESOURCEUNAVAILABLE_AUTOSCALINGGROUPINACTIVITY = "ResourceUnavailable.AutoScalingGroupInActivity"
//  RESOURCEUNAVAILABLE_LBBACKENDREGIONINCONSISTENT = "ResourceUnavailable.LbBackendRegionInconsistent"
//  RESOURCEUNAVAILABLE_LBVPCINCONSISTENT = "ResourceUnavailable.LbVpcInconsistent"
//  RESOURCEUNAVAILABLE_LOADBALANCERINOPERATION = "ResourceUnavailable.LoadBalancerInOperation"
func (c *Client) ModifyLoadBalancers(request *ModifyLoadBalancersRequest) (response *ModifyLoadBalancersResponse, err error) {
    return c.ModifyLoadBalancersWithContext(context.Background(), request)
}

// ModifyLoadBalancers
// This API is used to modify the cloud load balancers of a scaling group.
//
// 
//
// * This API can specify a new cloud load balancer configuration for the scaling group. The new configuration overwrites the original load balancer configuration.
//
// * To clear the cloud load balancer of the scaling group, specify only the scaling group ID but not the specific cloud load balancer.
//
// * This API modifies the cloud load balancer of the scaling group and generate a scaling activity to asynchronously modify the cloud load balancers of existing instances.
//
// error code that may be returned:
//  FAILEDOPERATION_NOACTIVITYTOGENERATE = "FailedOperation.NoActivityToGenerate"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INVALIDPARAMETER_CONFLICT = "InvalidParameter.Conflict"
//  INVALIDPARAMETER_INSCENARIO = "InvalidParameter.InScenario"
//  INVALIDPARAMETERVALUE_CLASSICLB = "InvalidParameterValue.ClassicLb"
//  INVALIDPARAMETERVALUE_DUPLICATEDFORWARDLB = "InvalidParameterValue.DuplicatedForwardLb"
//  INVALIDPARAMETERVALUE_FORWARDLB = "InvalidParameterValue.ForwardLb"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_INVALIDCLBREGION = "InvalidParameterValue.InvalidClbRegion"
//  INVALIDPARAMETERVALUE_LBPROJECTINCONSISTENT = "InvalidParameterValue.LbProjectInconsistent"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  MISSINGPARAMETER_INSCENARIO = "MissingParameter.InScenario"
//  RESOURCENOTFOUND_AUTOSCALINGGROUPNOTFOUND = "ResourceNotFound.AutoScalingGroupNotFound"
//  RESOURCENOTFOUND_LISTENERNOTFOUND = "ResourceNotFound.ListenerNotFound"
//  RESOURCENOTFOUND_LOADBALANCERNOTFOUND = "ResourceNotFound.LoadBalancerNotFound"
//  RESOURCENOTFOUND_LOCATIONNOTFOUND = "ResourceNotFound.LocationNotFound"
//  RESOURCEUNAVAILABLE_AUTOSCALINGGROUPINACTIVITY = "ResourceUnavailable.AutoScalingGroupInActivity"
//  RESOURCEUNAVAILABLE_LBBACKENDREGIONINCONSISTENT = "ResourceUnavailable.LbBackendRegionInconsistent"
//  RESOURCEUNAVAILABLE_LBVPCINCONSISTENT = "ResourceUnavailable.LbVpcInconsistent"
//  RESOURCEUNAVAILABLE_LOADBALANCERINOPERATION = "ResourceUnavailable.LoadBalancerInOperation"
func (c *Client) ModifyLoadBalancersWithContext(ctx context.Context, request *ModifyLoadBalancersRequest) (response *ModifyLoadBalancersResponse, err error) {
    if request == nil {
        request = NewModifyLoadBalancersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLoadBalancers require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLoadBalancersResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNotificationConfigurationRequest() (request *ModifyNotificationConfigurationRequest) {
    request = &ModifyNotificationConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "ModifyNotificationConfiguration")
    
    
    return
}

func NewModifyNotificationConfigurationResponse() (response *ModifyNotificationConfigurationResponse) {
    response = &ModifyNotificationConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyNotificationConfiguration
// This API is used to modify a notification policy.
//
// * The receiver type of the notification policy cannot be modified.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_CONFLICTNOTIFICATIONTARGET = "InvalidParameterValue.ConflictNotificationTarget"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGNOTIFICATIONID = "InvalidParameterValue.InvalidAutoScalingNotificationId"
//  INVALIDPARAMETERVALUE_INVALIDNOTIFICATIONUSERGROUPID = "InvalidParameterValue.InvalidNotificationUserGroupId"
//  INVALIDPARAMETERVALUE_USERGROUPIDNOTFOUND = "InvalidParameterValue.UserGroupIdNotFound"
//  RESOURCENOTFOUND_AUTOSCALINGNOTIFICATIONNOTFOUND = "ResourceNotFound.AutoScalingNotificationNotFound"
//  RESOURCENOTFOUND_CMQQUEUENOTFOUND = "ResourceNotFound.CmqQueueNotFound"
//  RESOURCENOTFOUND_TDMQCMQQUEUENOTFOUND = "ResourceNotFound.TDMQCMQQueueNotFound"
//  RESOURCENOTFOUND_TDMQCMQTOPICNOTFOUND = "ResourceNotFound.TDMQCMQTopicNotFound"
//  RESOURCEUNAVAILABLE_TDMQCMQTOPICHASNOSUBSCRIBER = "ResourceUnavailable.TDMQCMQTopicHasNoSubscriber"
func (c *Client) ModifyNotificationConfiguration(request *ModifyNotificationConfigurationRequest) (response *ModifyNotificationConfigurationResponse, err error) {
    return c.ModifyNotificationConfigurationWithContext(context.Background(), request)
}

// ModifyNotificationConfiguration
// This API is used to modify a notification policy.
//
// * The receiver type of the notification policy cannot be modified.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_CONFLICTNOTIFICATIONTARGET = "InvalidParameterValue.ConflictNotificationTarget"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGNOTIFICATIONID = "InvalidParameterValue.InvalidAutoScalingNotificationId"
//  INVALIDPARAMETERVALUE_INVALIDNOTIFICATIONUSERGROUPID = "InvalidParameterValue.InvalidNotificationUserGroupId"
//  INVALIDPARAMETERVALUE_USERGROUPIDNOTFOUND = "InvalidParameterValue.UserGroupIdNotFound"
//  RESOURCENOTFOUND_AUTOSCALINGNOTIFICATIONNOTFOUND = "ResourceNotFound.AutoScalingNotificationNotFound"
//  RESOURCENOTFOUND_CMQQUEUENOTFOUND = "ResourceNotFound.CmqQueueNotFound"
//  RESOURCENOTFOUND_TDMQCMQQUEUENOTFOUND = "ResourceNotFound.TDMQCMQQueueNotFound"
//  RESOURCENOTFOUND_TDMQCMQTOPICNOTFOUND = "ResourceNotFound.TDMQCMQTopicNotFound"
//  RESOURCEUNAVAILABLE_TDMQCMQTOPICHASNOSUBSCRIBER = "ResourceUnavailable.TDMQCMQTopicHasNoSubscriber"
func (c *Client) ModifyNotificationConfigurationWithContext(ctx context.Context, request *ModifyNotificationConfigurationRequest) (response *ModifyNotificationConfigurationResponse, err error) {
    if request == nil {
        request = NewModifyNotificationConfigurationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNotificationConfiguration require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNotificationConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyScalingPolicyRequest() (request *ModifyScalingPolicyRequest) {
    request = &ModifyScalingPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "ModifyScalingPolicy")
    
    
    return
}

func NewModifyScalingPolicyResponse() (response *ModifyScalingPolicyResponse) {
    response = &ModifyScalingPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyScalingPolicy
// This API (ModifyScalingPolicy) is used to modify an alarm trigger policy.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGPOLICYID = "InvalidParameterValue.InvalidAutoScalingPolicyId"
//  INVALIDPARAMETERVALUE_INVALIDNOTIFICATIONUSERGROUPID = "InvalidParameterValue.InvalidNotificationUserGroupId"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_SCALINGPOLICYNAMEDUPLICATE = "InvalidParameterValue.ScalingPolicyNameDuplicate"
//  RESOURCENOTFOUND_AUTOSCALINGGROUPNOTFOUND = "ResourceNotFound.AutoScalingGroupNotFound"
//  RESOURCENOTFOUND_SCALINGPOLICYNOTFOUND = "ResourceNotFound.ScalingPolicyNotFound"
func (c *Client) ModifyScalingPolicy(request *ModifyScalingPolicyRequest) (response *ModifyScalingPolicyResponse, err error) {
    return c.ModifyScalingPolicyWithContext(context.Background(), request)
}

// ModifyScalingPolicy
// This API (ModifyScalingPolicy) is used to modify an alarm trigger policy.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGPOLICYID = "InvalidParameterValue.InvalidAutoScalingPolicyId"
//  INVALIDPARAMETERVALUE_INVALIDNOTIFICATIONUSERGROUPID = "InvalidParameterValue.InvalidNotificationUserGroupId"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_SCALINGPOLICYNAMEDUPLICATE = "InvalidParameterValue.ScalingPolicyNameDuplicate"
//  RESOURCENOTFOUND_AUTOSCALINGGROUPNOTFOUND = "ResourceNotFound.AutoScalingGroupNotFound"
//  RESOURCENOTFOUND_SCALINGPOLICYNOTFOUND = "ResourceNotFound.ScalingPolicyNotFound"
func (c *Client) ModifyScalingPolicyWithContext(ctx context.Context, request *ModifyScalingPolicyRequest) (response *ModifyScalingPolicyResponse, err error) {
    if request == nil {
        request = NewModifyScalingPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyScalingPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyScalingPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyScheduledActionRequest() (request *ModifyScheduledActionRequest) {
    request = &ModifyScheduledActionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "ModifyScheduledAction")
    
    
    return
}

func NewModifyScheduledActionResponse() (response *ModifyScheduledActionResponse) {
    response = &ModifyScheduledActionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyScheduledAction
// This API (ModifyScheduledAction) is used to modify a scheduled task.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_CRONEXPRESSIONILLEGAL = "InvalidParameterValue.CronExpressionIllegal"
//  INVALIDPARAMETERVALUE_ENDTIMEBEFORESTARTTIME = "InvalidParameterValue.EndTimeBeforeStartTime"
//  INVALIDPARAMETERVALUE_INVALIDSCHEDULEDACTIONID = "InvalidParameterValue.InvalidScheduledActionId"
//  INVALIDPARAMETERVALUE_INVALIDSCHEDULEDACTIONNAMEINCLUDEILLEGALCHAR = "InvalidParameterValue.InvalidScheduledActionNameIncludeIllegalChar"
//  INVALIDPARAMETERVALUE_SCHEDULEDACTIONNAMEDUPLICATE = "InvalidParameterValue.ScheduledActionNameDuplicate"
//  INVALIDPARAMETERVALUE_SIZE = "InvalidParameterValue.Size"
//  INVALIDPARAMETERVALUE_STARTTIMEBEFORECURRENTTIME = "InvalidParameterValue.StartTimeBeforeCurrentTime"
//  INVALIDPARAMETERVALUE_TIMEFORMAT = "InvalidParameterValue.TimeFormat"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED_DESIREDCAPACITYLIMITEXCEEDED = "LimitExceeded.DesiredCapacityLimitExceeded"
//  LIMITEXCEEDED_MAXSIZELIMITEXCEEDED = "LimitExceeded.MaxSizeLimitExceeded"
//  LIMITEXCEEDED_MINSIZELIMITEXCEEDED = "LimitExceeded.MinSizeLimitExceeded"
//  LIMITEXCEEDED_SCHEDULEDACTIONLIMITEXCEEDED = "LimitExceeded.ScheduledActionLimitExceeded"
//  RESOURCENOTFOUND_SCHEDULEDACTIONNOTFOUND = "ResourceNotFound.ScheduledActionNotFound"
func (c *Client) ModifyScheduledAction(request *ModifyScheduledActionRequest) (response *ModifyScheduledActionResponse, err error) {
    return c.ModifyScheduledActionWithContext(context.Background(), request)
}

// ModifyScheduledAction
// This API (ModifyScheduledAction) is used to modify a scheduled task.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_CRONEXPRESSIONILLEGAL = "InvalidParameterValue.CronExpressionIllegal"
//  INVALIDPARAMETERVALUE_ENDTIMEBEFORESTARTTIME = "InvalidParameterValue.EndTimeBeforeStartTime"
//  INVALIDPARAMETERVALUE_INVALIDSCHEDULEDACTIONID = "InvalidParameterValue.InvalidScheduledActionId"
//  INVALIDPARAMETERVALUE_INVALIDSCHEDULEDACTIONNAMEINCLUDEILLEGALCHAR = "InvalidParameterValue.InvalidScheduledActionNameIncludeIllegalChar"
//  INVALIDPARAMETERVALUE_SCHEDULEDACTIONNAMEDUPLICATE = "InvalidParameterValue.ScheduledActionNameDuplicate"
//  INVALIDPARAMETERVALUE_SIZE = "InvalidParameterValue.Size"
//  INVALIDPARAMETERVALUE_STARTTIMEBEFORECURRENTTIME = "InvalidParameterValue.StartTimeBeforeCurrentTime"
//  INVALIDPARAMETERVALUE_TIMEFORMAT = "InvalidParameterValue.TimeFormat"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED_DESIREDCAPACITYLIMITEXCEEDED = "LimitExceeded.DesiredCapacityLimitExceeded"
//  LIMITEXCEEDED_MAXSIZELIMITEXCEEDED = "LimitExceeded.MaxSizeLimitExceeded"
//  LIMITEXCEEDED_MINSIZELIMITEXCEEDED = "LimitExceeded.MinSizeLimitExceeded"
//  LIMITEXCEEDED_SCHEDULEDACTIONLIMITEXCEEDED = "LimitExceeded.ScheduledActionLimitExceeded"
//  RESOURCENOTFOUND_SCHEDULEDACTIONNOTFOUND = "ResourceNotFound.ScheduledActionNotFound"
func (c *Client) ModifyScheduledActionWithContext(ctx context.Context, request *ModifyScheduledActionRequest) (response *ModifyScheduledActionResponse, err error) {
    if request == nil {
        request = NewModifyScheduledActionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyScheduledAction require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyScheduledActionResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveInstancesRequest() (request *RemoveInstancesRequest) {
    request = &RemoveInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "RemoveInstances")
    
    
    return
}

func NewRemoveInstancesResponse() (response *RemoveInstancesResponse) {
    response = &RemoveInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RemoveInstances
// This API is used to delete CVM instances from a scaling group. Instances that are automatically created through AS will be terminated, while those manually added to the scaling group will be removed and retained.
//
// * If the number of remaining `IN_SERVICE` instances in the scaling group is less than the minimum capacity, this API will return an error.
//
// * However, if the scaling group is in `DISABLED` status, the removal will not verify the relationship between the number of `IN_SERVICE` instances and the minimum capacity.
//
// * This removal will unassociate the CVM from the CLB instance that has been configured for the scaling group.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceId"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  RESOURCEINSUFFICIENT_AUTOSCALINGGROUPBELOWMINSIZE = "ResourceInsufficient.AutoScalingGroupBelowMinSize"
//  RESOURCEINSUFFICIENT_INSERVICEINSTANCEBELOWMINSIZE = "ResourceInsufficient.InServiceInstanceBelowMinSize"
//  RESOURCENOTFOUND_AUTOSCALINGGROUPNOTFOUND = "ResourceNotFound.AutoScalingGroupNotFound"
//  RESOURCENOTFOUND_INSTANCESNOTINAUTOSCALINGGROUP = "ResourceNotFound.InstancesNotInAutoScalingGroup"
//  RESOURCEUNAVAILABLE_AUTOSCALINGGROUPINACTIVITY = "ResourceUnavailable.AutoScalingGroupInActivity"
//  RESOURCEUNAVAILABLE_INSTANCEINOPERATION = "ResourceUnavailable.InstanceInOperation"
func (c *Client) RemoveInstances(request *RemoveInstancesRequest) (response *RemoveInstancesResponse, err error) {
    return c.RemoveInstancesWithContext(context.Background(), request)
}

// RemoveInstances
// This API is used to delete CVM instances from a scaling group. Instances that are automatically created through AS will be terminated, while those manually added to the scaling group will be removed and retained.
//
// * If the number of remaining `IN_SERVICE` instances in the scaling group is less than the minimum capacity, this API will return an error.
//
// * However, if the scaling group is in `DISABLED` status, the removal will not verify the relationship between the number of `IN_SERVICE` instances and the minimum capacity.
//
// * This removal will unassociate the CVM from the CLB instance that has been configured for the scaling group.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceId"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  RESOURCEINSUFFICIENT_AUTOSCALINGGROUPBELOWMINSIZE = "ResourceInsufficient.AutoScalingGroupBelowMinSize"
//  RESOURCEINSUFFICIENT_INSERVICEINSTANCEBELOWMINSIZE = "ResourceInsufficient.InServiceInstanceBelowMinSize"
//  RESOURCENOTFOUND_AUTOSCALINGGROUPNOTFOUND = "ResourceNotFound.AutoScalingGroupNotFound"
//  RESOURCENOTFOUND_INSTANCESNOTINAUTOSCALINGGROUP = "ResourceNotFound.InstancesNotInAutoScalingGroup"
//  RESOURCEUNAVAILABLE_AUTOSCALINGGROUPINACTIVITY = "ResourceUnavailable.AutoScalingGroupInActivity"
//  RESOURCEUNAVAILABLE_INSTANCEINOPERATION = "ResourceUnavailable.InstanceInOperation"
func (c *Client) RemoveInstancesWithContext(ctx context.Context, request *RemoveInstancesRequest) (response *RemoveInstancesResponse, err error) {
    if request == nil {
        request = NewRemoveInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewSetInstancesProtectionRequest() (request *SetInstancesProtectionRequest) {
    request = &SetInstancesProtectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "SetInstancesProtection")
    
    
    return
}

func NewSetInstancesProtectionResponse() (response *SetInstancesProtectionResponse) {
    response = &SetInstancesProtectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetInstancesProtection
// This API is used to enable scale-in protection for an instance.
//
// When scale-in protection is enabled, the instance will not be removed in scale-in activities triggered by replacement of unhealthy instances, alarm threshold reached, change of desired quantity, etc.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceId"
//  RESOURCENOTFOUND_AUTOSCALINGGROUPNOTFOUND = "ResourceNotFound.AutoScalingGroupNotFound"
//  RESOURCENOTFOUND_INSTANCESNOTINAUTOSCALINGGROUP = "ResourceNotFound.InstancesNotInAutoScalingGroup"
func (c *Client) SetInstancesProtection(request *SetInstancesProtectionRequest) (response *SetInstancesProtectionResponse, err error) {
    return c.SetInstancesProtectionWithContext(context.Background(), request)
}

// SetInstancesProtection
// This API is used to enable scale-in protection for an instance.
//
// When scale-in protection is enabled, the instance will not be removed in scale-in activities triggered by replacement of unhealthy instances, alarm threshold reached, change of desired quantity, etc.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceId"
//  RESOURCENOTFOUND_AUTOSCALINGGROUPNOTFOUND = "ResourceNotFound.AutoScalingGroupNotFound"
//  RESOURCENOTFOUND_INSTANCESNOTINAUTOSCALINGGROUP = "ResourceNotFound.InstancesNotInAutoScalingGroup"
func (c *Client) SetInstancesProtectionWithContext(ctx context.Context, request *SetInstancesProtectionRequest) (response *SetInstancesProtectionResponse, err error) {
    if request == nil {
        request = NewSetInstancesProtectionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetInstancesProtection require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetInstancesProtectionResponse()
    err = c.Send(request, response)
    return
}

func NewStartAutoScalingInstancesRequest() (request *StartAutoScalingInstancesRequest) {
    request = &StartAutoScalingInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "StartAutoScalingInstances")
    
    
    return
}

func NewStartAutoScalingInstancesResponse() (response *StartAutoScalingInstancesResponse) {
    response = &StartAutoScalingInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartAutoScalingInstances
// This API is used to start up CVM instances in a scaling group.
//
// * After startup, the instance will be in the `IN_SERVICE` status, which will increase the desired capacity. Please note that the desired capacity cannot exceed the maximum value.
//
// * This API supports batch operation. Up to 100 instances can be started up in each request.
//
// error code that may be returned:
//  FAILEDOPERATION_NOACTIVITYTOGENERATE = "FailedOperation.NoActivityToGenerate"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceId"
//  RESOURCEINSUFFICIENT_AUTOSCALINGGROUPABOVEMAXSIZE = "ResourceInsufficient.AutoScalingGroupAboveMaxSize"
//  RESOURCEINSUFFICIENT_INSERVICEINSTANCEABOVEMAXSIZE = "ResourceInsufficient.InServiceInstanceAboveMaxSize"
//  RESOURCENOTFOUND_AUTOSCALINGGROUPNOTFOUND = "ResourceNotFound.AutoScalingGroupNotFound"
//  RESOURCENOTFOUND_INSTANCESNOTINAUTOSCALINGGROUP = "ResourceNotFound.InstancesNotInAutoScalingGroup"
//  RESOURCEUNAVAILABLE_AUTOSCALINGGROUPINACTIVITY = "ResourceUnavailable.AutoScalingGroupInActivity"
//  RESOURCEUNAVAILABLE_LOADBALANCERINOPERATION = "ResourceUnavailable.LoadBalancerInOperation"
func (c *Client) StartAutoScalingInstances(request *StartAutoScalingInstancesRequest) (response *StartAutoScalingInstancesResponse, err error) {
    return c.StartAutoScalingInstancesWithContext(context.Background(), request)
}

// StartAutoScalingInstances
// This API is used to start up CVM instances in a scaling group.
//
// * After startup, the instance will be in the `IN_SERVICE` status, which will increase the desired capacity. Please note that the desired capacity cannot exceed the maximum value.
//
// * This API supports batch operation. Up to 100 instances can be started up in each request.
//
// error code that may be returned:
//  FAILEDOPERATION_NOACTIVITYTOGENERATE = "FailedOperation.NoActivityToGenerate"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceId"
//  RESOURCEINSUFFICIENT_AUTOSCALINGGROUPABOVEMAXSIZE = "ResourceInsufficient.AutoScalingGroupAboveMaxSize"
//  RESOURCEINSUFFICIENT_INSERVICEINSTANCEABOVEMAXSIZE = "ResourceInsufficient.InServiceInstanceAboveMaxSize"
//  RESOURCENOTFOUND_AUTOSCALINGGROUPNOTFOUND = "ResourceNotFound.AutoScalingGroupNotFound"
//  RESOURCENOTFOUND_INSTANCESNOTINAUTOSCALINGGROUP = "ResourceNotFound.InstancesNotInAutoScalingGroup"
//  RESOURCEUNAVAILABLE_AUTOSCALINGGROUPINACTIVITY = "ResourceUnavailable.AutoScalingGroupInActivity"
//  RESOURCEUNAVAILABLE_LOADBALANCERINOPERATION = "ResourceUnavailable.LoadBalancerInOperation"
func (c *Client) StartAutoScalingInstancesWithContext(ctx context.Context, request *StartAutoScalingInstancesRequest) (response *StartAutoScalingInstancesResponse, err error) {
    if request == nil {
        request = NewStartAutoScalingInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartAutoScalingInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartAutoScalingInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewStopAutoScalingInstancesRequest() (request *StopAutoScalingInstancesRequest) {
    request = &StopAutoScalingInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "StopAutoScalingInstances")
    
    
    return
}

func NewStopAutoScalingInstancesResponse() (response *StopAutoScalingInstancesResponse) {
    response = &StopAutoScalingInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopAutoScalingInstances
// This API is used to shut down CVM instances in a scaling group.
//
// * Use the `SOFT_FIRST` shutdown, which means the CVM will be forcibly shut down if the soft shutdown fails.
//
// * Shutting down instances in the `IN_SERVICE` status will reduce the desired capacity, but the desired capacity cannot be less than the minimum value.
//
// * To use the `STOP_CHARGING` shutdown, the instances you want to shut down must satisfy the conditions of [no charges when shut down](https://intl.cloud.tencent.com/document/product/213/19918?from_cn_redirect=1).
//
// * This API supports batch operation. Up to 100 instances can be shut down in each request.
//
// error code that may be returned:
//  CALLCVMERROR = "CallCvmError"
//  FAILEDOPERATION_NOACTIVITYTOGENERATE = "FailedOperation.NoActivityToGenerate"
//  INTERNALERROR_CALLEEERROR = "InternalError.CalleeError"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceId"
//  RESOURCEINSUFFICIENT_AUTOSCALINGGROUPBELOWMINSIZE = "ResourceInsufficient.AutoScalingGroupBelowMinSize"
//  RESOURCEINSUFFICIENT_INSERVICEINSTANCEBELOWMINSIZE = "ResourceInsufficient.InServiceInstanceBelowMinSize"
//  RESOURCENOTFOUND_AUTOSCALINGGROUPNOTFOUND = "ResourceNotFound.AutoScalingGroupNotFound"
//  RESOURCENOTFOUND_INSTANCESNOTINAUTOSCALINGGROUP = "ResourceNotFound.InstancesNotInAutoScalingGroup"
//  RESOURCEUNAVAILABLE_AUTOSCALINGGROUPINACTIVITY = "ResourceUnavailable.AutoScalingGroupInActivity"
//  RESOURCEUNAVAILABLE_INSTANCEINOPERATION = "ResourceUnavailable.InstanceInOperation"
//  RESOURCEUNAVAILABLE_INSTANCENOTSUPPORTSTOPCHARGING = "ResourceUnavailable.InstanceNotSupportStopCharging"
//  RESOURCEUNAVAILABLE_LOADBALANCERINOPERATION = "ResourceUnavailable.LoadBalancerInOperation"
func (c *Client) StopAutoScalingInstances(request *StopAutoScalingInstancesRequest) (response *StopAutoScalingInstancesResponse, err error) {
    return c.StopAutoScalingInstancesWithContext(context.Background(), request)
}

// StopAutoScalingInstances
// This API is used to shut down CVM instances in a scaling group.
//
// * Use the `SOFT_FIRST` shutdown, which means the CVM will be forcibly shut down if the soft shutdown fails.
//
// * Shutting down instances in the `IN_SERVICE` status will reduce the desired capacity, but the desired capacity cannot be less than the minimum value.
//
// * To use the `STOP_CHARGING` shutdown, the instances you want to shut down must satisfy the conditions of [no charges when shut down](https://intl.cloud.tencent.com/document/product/213/19918?from_cn_redirect=1).
//
// * This API supports batch operation. Up to 100 instances can be shut down in each request.
//
// error code that may be returned:
//  CALLCVMERROR = "CallCvmError"
//  FAILEDOPERATION_NOACTIVITYTOGENERATE = "FailedOperation.NoActivityToGenerate"
//  INTERNALERROR_CALLEEERROR = "InternalError.CalleeError"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALINGGROUPID = "InvalidParameterValue.InvalidAutoScalingGroupId"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceId"
//  RESOURCEINSUFFICIENT_AUTOSCALINGGROUPBELOWMINSIZE = "ResourceInsufficient.AutoScalingGroupBelowMinSize"
//  RESOURCEINSUFFICIENT_INSERVICEINSTANCEBELOWMINSIZE = "ResourceInsufficient.InServiceInstanceBelowMinSize"
//  RESOURCENOTFOUND_AUTOSCALINGGROUPNOTFOUND = "ResourceNotFound.AutoScalingGroupNotFound"
//  RESOURCENOTFOUND_INSTANCESNOTINAUTOSCALINGGROUP = "ResourceNotFound.InstancesNotInAutoScalingGroup"
//  RESOURCEUNAVAILABLE_AUTOSCALINGGROUPINACTIVITY = "ResourceUnavailable.AutoScalingGroupInActivity"
//  RESOURCEUNAVAILABLE_INSTANCEINOPERATION = "ResourceUnavailable.InstanceInOperation"
//  RESOURCEUNAVAILABLE_INSTANCENOTSUPPORTSTOPCHARGING = "ResourceUnavailable.InstanceNotSupportStopCharging"
//  RESOURCEUNAVAILABLE_LOADBALANCERINOPERATION = "ResourceUnavailable.LoadBalancerInOperation"
func (c *Client) StopAutoScalingInstancesWithContext(ctx context.Context, request *StopAutoScalingInstancesRequest) (response *StopAutoScalingInstancesResponse, err error) {
    if request == nil {
        request = NewStopAutoScalingInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopAutoScalingInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopAutoScalingInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeLifecycleHookRequest() (request *UpgradeLifecycleHookRequest) {
    request = &UpgradeLifecycleHookRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "UpgradeLifecycleHook")
    
    
    return
}

func NewUpgradeLifecycleHookResponse() (response *UpgradeLifecycleHookResponse) {
    response = &UpgradeLifecycleHookResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpgradeLifecycleHook
// This API (UpgradeLifecycleHook) is used to upgrade a lifecycle hook.
//
// 
//
// * This API is used to upgrade a lifecycle hook in a "completely overriding" manner, i.e., it uniformly sets a new configuration according to the API parameters regardless of the original parameters. If optional fields are left empty, their default values will be used.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLCMQERROR = "InternalError.CallCmqError"
//  INTERNALERROR_CALLSTSERROR = "InternalError.CallStsError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONFLICT = "InvalidParameter.Conflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FILTER = "InvalidParameterValue.Filter"
//  INVALIDPARAMETERVALUE_INVALIDLIFECYCLEHOOKID = "InvalidParameterValue.InvalidLifecycleHookId"
//  INVALIDPARAMETERVALUE_LIFECYCLEHOOKNAMEDUPLICATED = "InvalidParameterValue.LifecycleHookNameDuplicated"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_LIFECYCLEHOOKNOTFOUND = "ResourceNotFound.LifecycleHookNotFound"
//  RESOURCEUNAVAILABLE_CMQTOPICHASNOSUBSCRIBER = "ResourceUnavailable.CmqTopicHasNoSubscriber"
func (c *Client) UpgradeLifecycleHook(request *UpgradeLifecycleHookRequest) (response *UpgradeLifecycleHookResponse, err error) {
    return c.UpgradeLifecycleHookWithContext(context.Background(), request)
}

// UpgradeLifecycleHook
// This API (UpgradeLifecycleHook) is used to upgrade a lifecycle hook.
//
// 
//
// * This API is used to upgrade a lifecycle hook in a "completely overriding" manner, i.e., it uniformly sets a new configuration according to the API parameters regardless of the original parameters. If optional fields are left empty, their default values will be used.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLCMQERROR = "InternalError.CallCmqError"
//  INTERNALERROR_CALLSTSERROR = "InternalError.CallStsError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONFLICT = "InvalidParameter.Conflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FILTER = "InvalidParameterValue.Filter"
//  INVALIDPARAMETERVALUE_INVALIDLIFECYCLEHOOKID = "InvalidParameterValue.InvalidLifecycleHookId"
//  INVALIDPARAMETERVALUE_LIFECYCLEHOOKNAMEDUPLICATED = "InvalidParameterValue.LifecycleHookNameDuplicated"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_LIFECYCLEHOOKNOTFOUND = "ResourceNotFound.LifecycleHookNotFound"
//  RESOURCEUNAVAILABLE_CMQTOPICHASNOSUBSCRIBER = "ResourceUnavailable.CmqTopicHasNoSubscriber"
func (c *Client) UpgradeLifecycleHookWithContext(ctx context.Context, request *UpgradeLifecycleHookRequest) (response *UpgradeLifecycleHookResponse, err error) {
    if request == nil {
        request = NewUpgradeLifecycleHookRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeLifecycleHook require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeLifecycleHookResponse()
    err = c.Send(request, response)
    return
}
