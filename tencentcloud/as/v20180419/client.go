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

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
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

// This API (AttachInstances) is used to add CVM instances to an auto scaling group.
func (c *Client) AttachInstances(request *AttachInstancesRequest) (response *AttachInstancesResponse, err error) {
    if request == nil {
        request = NewAttachInstancesRequest()
    }
    response = NewAttachInstancesResponse()
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

// This API (CompleteLifecycleAction) is used to complete a lifecycle action.
// 
// * The result ("CONTINUE" or "ABANDON") of a specific lifecycle hook can be specified by calling this API. If this API is not called at all, the lifecycle hook will be processed based on the "DefaultResult" after timeout.
func (c *Client) CompleteLifecycleAction(request *CompleteLifecycleActionRequest) (response *CompleteLifecycleActionResponse, err error) {
    if request == nil {
        request = NewCompleteLifecycleActionRequest()
    }
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

// This API (CreateAutoScalingGroup) is used to create an auto scaling group.
func (c *Client) CreateAutoScalingGroup(request *CreateAutoScalingGroupRequest) (response *CreateAutoScalingGroupResponse, err error) {
    if request == nil {
        request = NewCreateAutoScalingGroupRequest()
    }
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

// This API is used to create launch configurations and scaling groups based on an instance.
// 
// Note: for a scaling group that is created based on a monthly-subscribed instance, the instances added for scale-out are pay-as-you-go instance.
func (c *Client) CreateAutoScalingGroupFromInstance(request *CreateAutoScalingGroupFromInstanceRequest) (response *CreateAutoScalingGroupFromInstanceResponse, err error) {
    if request == nil {
        request = NewCreateAutoScalingGroupFromInstanceRequest()
    }
    response = NewCreateAutoScalingGroupFromInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLaunchConfigurationRequest() (request *CreateLaunchConfigurationRequest) {
    request = &CreateLaunchConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "CreateLaunchConfiguration")
    return
}

func NewCreateLaunchConfigurationResponse() (response *CreateLaunchConfigurationResponse) {
    response = &CreateLaunchConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (CreateLaunchConfiguration) is used to create a launch configuration.
// 
// * A few fields of a launch configuration can be modified through `ModifyLaunchConfigurationAttributes`. To use a new launch configuration, it is recommended to create it from scratch.
// 
// * You can create up to 20 launch configurations for each project. For more information, see [Usage Limits](https://intl.cloud.tencent.com/document/product/377/3120?from_cn_redirect=1).
func (c *Client) CreateLaunchConfiguration(request *CreateLaunchConfigurationRequest) (response *CreateLaunchConfigurationResponse, err error) {
    if request == nil {
        request = NewCreateLaunchConfigurationRequest()
    }
    response = NewCreateLaunchConfigurationResponse()
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

// This API (CreateLifeCycleHook) is used to create a lifecycle hook.
// 
// * You can configure message notifications in the following format for lifecycle hooks, which will be sent to your CMQ queue by AS:
// 
// ```
// {
// 	"Service": "Tencent Cloud Auto Scaling",
// 	"Time": "2019-03-14T10:15:11Z",
// 	"AppId": "1251783334",
// 	"ActivityId": "asa-fznnvrja",
// 	"AutoScalingGroupId": "asg-rrrrtttt",
// 	"LifecycleHookId": "ash-xxxxyyyy",
// 	"LifecycleHookName": "my-hook",
// 	"LifecycleActionToken": "3080e1c9-0efe-4dd7-ad3b-90cd6618298f",
// 	"InstanceId": "ins-aaaabbbb",
// 	"LifecycleTransition": "INSTANCE_LAUNCHING",
// 	"NotificationMetadata": ""
// }
// ```
func (c *Client) CreateLifecycleHook(request *CreateLifecycleHookRequest) (response *CreateLifecycleHookResponse, err error) {
    if request == nil {
        request = NewCreateLifecycleHookRequest()
    }
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

// This API (CreateNotificationConfiguration) is used to create a notification.
func (c *Client) CreateNotificationConfiguration(request *CreateNotificationConfigurationRequest) (response *CreateNotificationConfigurationResponse, err error) {
    if request == nil {
        request = NewCreateNotificationConfigurationRequest()
    }
    response = NewCreateNotificationConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePaiInstanceRequest() (request *CreatePaiInstanceRequest) {
    request = &CreatePaiInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "CreatePaiInstance")
    return
}

func NewCreatePaiInstanceResponse() (response *CreatePaiInstanceResponse) {
    response = &CreatePaiInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (CreatePaiInstance) is used to create a PAI instance.
func (c *Client) CreatePaiInstance(request *CreatePaiInstanceRequest) (response *CreatePaiInstanceResponse, err error) {
    if request == nil {
        request = NewCreatePaiInstanceRequest()
    }
    response = NewCreatePaiInstanceResponse()
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

// This API (CreateScalingPolicy) is used to create an alarm trigger policy.
func (c *Client) CreateScalingPolicy(request *CreateScalingPolicyRequest) (response *CreateScalingPolicyResponse, err error) {
    if request == nil {
        request = NewCreateScalingPolicyRequest()
    }
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

// This API (CreateScheduledAction) is used to create a scheduled task.
func (c *Client) CreateScheduledAction(request *CreateScheduledActionRequest) (response *CreateScheduledActionResponse, err error) {
    if request == nil {
        request = NewCreateScheduledActionRequest()
    }
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

// This API (DeleteAutoScalingGroup) is used to delete the specified auto scaling group that has no instances and remains inactive.
func (c *Client) DeleteAutoScalingGroup(request *DeleteAutoScalingGroupRequest) (response *DeleteAutoScalingGroupResponse, err error) {
    if request == nil {
        request = NewDeleteAutoScalingGroupRequest()
    }
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

// This API (DeleteLaunchConfiguration) is used to delete a launch configuration.
// 
// * If the launch configuration is active in an auto scaling group, it cannot be deleted.
func (c *Client) DeleteLaunchConfiguration(request *DeleteLaunchConfigurationRequest) (response *DeleteLaunchConfigurationResponse, err error) {
    if request == nil {
        request = NewDeleteLaunchConfigurationRequest()
    }
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

// This API (DeleteLifeCycleHook) is used to delete a lifecycle hook.
func (c *Client) DeleteLifecycleHook(request *DeleteLifecycleHookRequest) (response *DeleteLifecycleHookResponse, err error) {
    if request == nil {
        request = NewDeleteLifecycleHookRequest()
    }
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

// This API (DeleteNotificationConfiguration) is used to delete the specified notification.
func (c *Client) DeleteNotificationConfiguration(request *DeleteNotificationConfigurationRequest) (response *DeleteNotificationConfigurationResponse, err error) {
    if request == nil {
        request = NewDeleteNotificationConfigurationRequest()
    }
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

// This API (DeleteScalingPolicy) is used to delete an alarm trigger policy.
func (c *Client) DeleteScalingPolicy(request *DeleteScalingPolicyRequest) (response *DeleteScalingPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteScalingPolicyRequest()
    }
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

// This API (DeleteScheduledAction) is used to delete the specified scheduled task.
func (c *Client) DeleteScheduledAction(request *DeleteScheduledActionRequest) (response *DeleteScheduledActionResponse, err error) {
    if request == nil {
        request = NewDeleteScheduledActionRequest()
    }
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

// This API (DescribeAccountLimits) is used to query the limits of user's AS resources.
func (c *Client) DescribeAccountLimits(request *DescribeAccountLimitsRequest) (response *DescribeAccountLimitsResponse, err error) {
    if request == nil {
        request = NewDescribeAccountLimitsRequest()
    }
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

// This API (DescribeAutoScalingActivities) is used to query the activity history of an auto scaling group.
func (c *Client) DescribeAutoScalingActivities(request *DescribeAutoScalingActivitiesRequest) (response *DescribeAutoScalingActivitiesResponse, err error) {
    if request == nil {
        request = NewDescribeAutoScalingActivitiesRequest()
    }
    response = NewDescribeAutoScalingActivitiesResponse()
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

// This API is used to query the latest activity history of an auto scaling group.
func (c *Client) DescribeAutoScalingGroupLastActivities(request *DescribeAutoScalingGroupLastActivitiesRequest) (response *DescribeAutoScalingGroupLastActivitiesResponse, err error) {
    if request == nil {
        request = NewDescribeAutoScalingGroupLastActivitiesRequest()
    }
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

// This API (DescribeAutoScalingGroups) is used to query the information of auto scaling groups.
// 
// * You can query the details of auto scaling groups based on information such as auto scaling group ID, auto scaling group name, or launch configuration ID. For more information on filters, see `Filter`.
// * If the parameter is empty, a certain number (specified by `Limit` and 20 by default) of auto scaling groups of the current user will be returned.
func (c *Client) DescribeAutoScalingGroups(request *DescribeAutoScalingGroupsRequest) (response *DescribeAutoScalingGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeAutoScalingGroupsRequest()
    }
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

// This API (DescribeAutoScalingInstances) is used to query the information of instances associated with AS.
// 
// * You can query the details of instances based on information such as instance ID and auto scaling group ID. For more information on filters, see `Filter`.
// * If the parameter is empty, a certain number (specified by `Limit` and 20 by default) of instances of the current user will be returned.
func (c *Client) DescribeAutoScalingInstances(request *DescribeAutoScalingInstancesRequest) (response *DescribeAutoScalingInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeAutoScalingInstancesRequest()
    }
    response = NewDescribeAutoScalingInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLaunchConfigurationsRequest() (request *DescribeLaunchConfigurationsRequest) {
    request = &DescribeLaunchConfigurationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "DescribeLaunchConfigurations")
    return
}

func NewDescribeLaunchConfigurationsResponse() (response *DescribeLaunchConfigurationsResponse) {
    response = &DescribeLaunchConfigurationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (DescribeLaunchConfigurations) is used to query the information of launch configurations.
// 
// * You can query the launch configuration details based on information such as launch configuration ID and name. For more information on filters, see `Filter`.
// * If the parameter is empty, a certain number (specified by `Limit` and 20 by default) of launch configurations of the current user will be returned.
func (c *Client) DescribeLaunchConfigurations(request *DescribeLaunchConfigurationsRequest) (response *DescribeLaunchConfigurationsResponse, err error) {
    if request == nil {
        request = NewDescribeLaunchConfigurationsRequest()
    }
    response = NewDescribeLaunchConfigurationsResponse()
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

// This API (DescribeLifecycleHooks) is used to query the information of lifecycle hooks.
// 
// * You can query the details of lifecycle hooks based on information such as auto scaling group ID, lifecycle hook ID, or lifecycle hook name. For more information on filters, see `Filter`.
// * If the parameter is empty, a certain number (specified by `Limit` and 20 by default) of lifecycle hooks of the current user will be returned.
func (c *Client) DescribeLifecycleHooks(request *DescribeLifecycleHooksRequest) (response *DescribeLifecycleHooksResponse, err error) {
    if request == nil {
        request = NewDescribeLifecycleHooksRequest()
    }
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

// This API (DescribeNotificationConfigurations) is used to query the information of one or more notifications.
// 
// You can query the details of notifications based on information such as notification ID and auto scaling group ID. For more information on filters, see `Filter`.
// If the parameter is empty, a certain number (specified by `Limit` and 20 by default) of notifications of the current user will be returned.
func (c *Client) DescribeNotificationConfigurations(request *DescribeNotificationConfigurationsRequest) (response *DescribeNotificationConfigurationsResponse, err error) {
    if request == nil {
        request = NewDescribeNotificationConfigurationsRequest()
    }
    response = NewDescribeNotificationConfigurationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePaiInstancesRequest() (request *DescribePaiInstancesRequest) {
    request = &DescribePaiInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "DescribePaiInstances")
    return
}

func NewDescribePaiInstancesResponse() (response *DescribePaiInstancesResponse) {
    response = &DescribePaiInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (DescribePaiInstances) is used to query the information of PAI instances.
// 
// * You can query the detailed information of PAI instances based on information such as instance ID and instance domain name. For more information on filters, see `Filter`.
// * If the parameter is empty, a certain number (specified by `Limit` and 20 by default) of PAI instances of the current user will be returned.
func (c *Client) DescribePaiInstances(request *DescribePaiInstancesRequest) (response *DescribePaiInstancesResponse, err error) {
    if request == nil {
        request = NewDescribePaiInstancesRequest()
    }
    response = NewDescribePaiInstancesResponse()
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

// This API (DescribeScalingPolicies) is used to query alarm trigger policies.
func (c *Client) DescribeScalingPolicies(request *DescribeScalingPoliciesRequest) (response *DescribeScalingPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeScalingPoliciesRequest()
    }
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

// This API (DescribeScheduledActions) is used to query the details of one or more scheduled tasks.
// 
// * You can query the details of scheduled tasks based on information such as scheduled task ID, scheduled task name, or auto scaling group ID. For more information on filters, see `Filter`.
// * If the parameter is empty, a certain number (specified by `Limit` and 20 by default) of scheduled tasks of the current user will be returned.
func (c *Client) DescribeScheduledActions(request *DescribeScheduledActionsRequest) (response *DescribeScheduledActionsResponse, err error) {
    if request == nil {
        request = NewDescribeScheduledActionsRequest()
    }
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

// This API is used to remove CVM instances from a scaling group. Instances removed via this API will not be terminated.
// * If the number of remaining `IN_SERVICE` instances in the scaling group is less than the minimum capacity, an error will be reported.
// * However, if the scaling group is in `DISABLED` status, the removal will not verify the relationship between the number of `IN_SERVICE` instances and the minimum capacity.
func (c *Client) DetachInstances(request *DetachInstancesRequest) (response *DetachInstancesResponse, err error) {
    if request == nil {
        request = NewDetachInstancesRequest()
    }
    response = NewDetachInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDisableAutoScalingGroupRequest() (request *DisableAutoScalingGroupRequest) {
    request = &DisableAutoScalingGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "DisableAutoScalingGroup")
    return
}

func NewDisableAutoScalingGroupResponse() (response *DisableAutoScalingGroupResponse) {
    response = &DisableAutoScalingGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (DisableAutoScalingGroup) is used to disable the specified auto scaling group.
func (c *Client) DisableAutoScalingGroup(request *DisableAutoScalingGroupRequest) (response *DisableAutoScalingGroupResponse, err error) {
    if request == nil {
        request = NewDisableAutoScalingGroupRequest()
    }
    response = NewDisableAutoScalingGroupResponse()
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

// This API (EnableAutoScalingGroup) is used to enable the specified auto scaling group.
func (c *Client) EnableAutoScalingGroup(request *EnableAutoScalingGroupRequest) (response *EnableAutoScalingGroupResponse, err error) {
    if request == nil {
        request = NewEnableAutoScalingGroupRequest()
    }
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

// This API (ExecuteScalingPolicy) is used to execute a scaling policy.
// 
// * The scaling policy can be executed based on the scaling policy ID.
// * When the auto scaling group to which the scaling policy belongs is performing a scaling activity, the scaling policy will be rejected.
func (c *Client) ExecuteScalingPolicy(request *ExecuteScalingPolicyRequest) (response *ExecuteScalingPolicyResponse, err error) {
    if request == nil {
        request = NewExecuteScalingPolicyRequest()
    }
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

// This API (ModifyAutoScalingGroup) is used to modify an auto scaling group.
func (c *Client) ModifyAutoScalingGroup(request *ModifyAutoScalingGroupRequest) (response *ModifyAutoScalingGroupResponse, err error) {
    if request == nil {
        request = NewModifyAutoScalingGroupRequest()
    }
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

// This API (ModifyDesiredCapacity) is used to modify the desired number of instances in the specified auto scaling group.
func (c *Client) ModifyDesiredCapacity(request *ModifyDesiredCapacityRequest) (response *ModifyDesiredCapacityResponse, err error) {
    if request == nil {
        request = NewModifyDesiredCapacityRequest()
    }
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

// This API (ModifyLaunchConfigurationAttributes) is used to modify some attributes of a launch configuration.
// 
// * The changes of launch configuration do not affect the existing instances. New instances will be created based on the modified configuration.
// * This API supports modifying certain simple types of attributes.
func (c *Client) ModifyLaunchConfigurationAttributes(request *ModifyLaunchConfigurationAttributesRequest) (response *ModifyLaunchConfigurationAttributesResponse, err error) {
    if request == nil {
        request = NewModifyLaunchConfigurationAttributesRequest()
    }
    response = NewModifyLaunchConfigurationAttributesResponse()
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

// This API (ModifyLoadBalancers) is used to modify the load balancers of an auto scaling group.
// 
// * This API can specify a new load balancer configuration for the auto scaling group. The new configuration overwrites the original load balancer configuration.
// * If you want to clear the load balancer for the auto scaling group, specify only the auto scaling group ID but not the specific load balancer when calling this API.
// * This API modifies the load balancer of the auto scaling group and generate a scaling activity to asynchronously modify the load balancers of existing instances.
func (c *Client) ModifyLoadBalancers(request *ModifyLoadBalancersRequest) (response *ModifyLoadBalancersResponse, err error) {
    if request == nil {
        request = NewModifyLoadBalancersRequest()
    }
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

// This API (ModifyNotificationConfiguration) is used to modify a notification.
func (c *Client) ModifyNotificationConfiguration(request *ModifyNotificationConfigurationRequest) (response *ModifyNotificationConfigurationResponse, err error) {
    if request == nil {
        request = NewModifyNotificationConfigurationRequest()
    }
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

// This API (ModifyScalingPolicy) is used to modify an alarm trigger policy.
func (c *Client) ModifyScalingPolicy(request *ModifyScalingPolicyRequest) (response *ModifyScalingPolicyResponse, err error) {
    if request == nil {
        request = NewModifyScalingPolicyRequest()
    }
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

// This API (ModifyScheduledAction) is used to modify a scheduled task.
func (c *Client) ModifyScheduledAction(request *ModifyScheduledActionRequest) (response *ModifyScheduledActionResponse, err error) {
    if request == nil {
        request = NewModifyScheduledActionRequest()
    }
    response = NewModifyScheduledActionResponse()
    err = c.Send(request, response)
    return
}

func NewPreviewPaiDomainNameRequest() (request *PreviewPaiDomainNameRequest) {
    request = &PreviewPaiDomainNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "PreviewPaiDomainName")
    return
}

func NewPreviewPaiDomainNameResponse() (response *PreviewPaiDomainNameResponse) {
    response = &PreviewPaiDomainNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (PreviewPaiDomainName) is used to preview a PAI domain name.
func (c *Client) PreviewPaiDomainName(request *PreviewPaiDomainNameRequest) (response *PreviewPaiDomainNameResponse, err error) {
    if request == nil {
        request = NewPreviewPaiDomainNameRequest()
    }
    response = NewPreviewPaiDomainNameResponse()
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

// This API is used to remove CVM instances from a scaling group. Instances that are automatically created through AS will be terminated, while those manually added to the scaling group will be removed and retained.
// * If the number of remaining `IN_SERVICE` instances in the scaling group is less than the minimum capacity, an error will be reported.
// * However, if the scaling group is in `DISABLED` status, the removal will not verify the relationship between the number of `IN_SERVICE` instances and the minimum capacity.
func (c *Client) RemoveInstances(request *RemoveInstancesRequest) (response *RemoveInstancesResponse, err error) {
    if request == nil {
        request = NewRemoveInstancesRequest()
    }
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

// This API (SetInstancesProtection) is used to enable scale-in protection for an instance.
// When an instance has scale-in protection enabled, it will not be removed when scaling is triggered by replacement of unhealthy instances, alarm trigger policy, threshold change, etc.
func (c *Client) SetInstancesProtection(request *SetInstancesProtectionRequest) (response *SetInstancesProtectionResponse, err error) {
    if request == nil {
        request = NewSetInstancesProtectionRequest()
    }
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

// This API is used to start up CVM instances in a scaling group.
// * After startup, the instance will be in the `IN_SERVICE` status, which will increase the desired capacity. Please note that the desired capacity cannot exceed the maximum value.
// * This API supports batch operation. Up to 100 instances can be started up in each request.
func (c *Client) StartAutoScalingInstances(request *StartAutoScalingInstancesRequest) (response *StartAutoScalingInstancesResponse, err error) {
    if request == nil {
        request = NewStartAutoScalingInstancesRequest()
    }
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

// This API is used to shut down CVM instances in a scaling group.
// * Use the `SOFT_FIRST` shutdown, which means the CVM will be forcibly shut down if the soft shutdown fails.
// * Shutting down instances in the `IN_SERVICE` status will reduce the desired capacity, but the desired capacity cannot be less than the minimum value.
// * To use the `STOP_CHARGING` shutdown, the instances you want to shut down must satisfy the conditions of [no charges when shut down](https://intl.cloud.tencent.com/document/product/213/19918?from_cn_redirect=1).
// * This API supports batch operation. Up to 100 instances can be shut down in each request.
func (c *Client) StopAutoScalingInstances(request *StopAutoScalingInstancesRequest) (response *StopAutoScalingInstancesResponse, err error) {
    if request == nil {
        request = NewStopAutoScalingInstancesRequest()
    }
    response = NewStopAutoScalingInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeLaunchConfigurationRequest() (request *UpgradeLaunchConfigurationRequest) {
    request = &UpgradeLaunchConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "UpgradeLaunchConfiguration")
    return
}

func NewUpgradeLaunchConfigurationResponse() (response *UpgradeLaunchConfigurationResponse) {
    response = &UpgradeLaunchConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (UpgradeLaunchConfiguration) is used to upgrade a launch configuration.
// 
// * This API is used to upgrade a launch configuration in a "completely overriding" manner, i.e., it uniformly sets a new configuration according to the API parameters regardless of the original parameters. If optional fields are left empty, their default values will be used.
// * After the launch configuration is upgraded, the existing instances that have been created by it will not be changed, but new instances will be created according to the new configuration.
func (c *Client) UpgradeLaunchConfiguration(request *UpgradeLaunchConfigurationRequest) (response *UpgradeLaunchConfigurationResponse, err error) {
    if request == nil {
        request = NewUpgradeLaunchConfigurationRequest()
    }
    response = NewUpgradeLaunchConfigurationResponse()
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

// This API (UpgradeLifecycleHook) is used to upgrade a lifecycle hook.
// 
// * This API is used to upgrade a lifecycle hook in a "completely overriding" manner, i.e., it uniformly sets a new configuration according to the API parameters regardless of the original parameters. If optional fields are left empty, their default values will be used.
func (c *Client) UpgradeLifecycleHook(request *UpgradeLifecycleHookRequest) (response *UpgradeLifecycleHookResponse, err error) {
    if request == nil {
        request = NewUpgradeLifecycleHookRequest()
    }
    response = NewUpgradeLifecycleHookResponse()
    err = c.Send(request, response)
    return
}
