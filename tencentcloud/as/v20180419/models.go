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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type Activity struct {
	// Auto scaling group ID.
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// Scaling activity ID.
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// Type of the scaling activity. Valid values:<br>
	// <li>`SCALE_OUT`: Scale out. <li>`SCALE_IN`: Scale in. <li>`ATTACH_INSTANCES`: Add instances. <li>`REMOVE_INSTANCES`: Terminate instances. <li>`DETACH_INSTANCES`: Remove instances. <li>`TERMINATE_INSTANCES_UNEXPECTEDLY`: Terminate instances in the CVM console. <li>`REPLACE_UNHEALTHY_INSTANCE`: Replace an unhealthy instance.
	// <li>`START_INSTANCES`: Starts up instances.
	// <li>`STOP_INSTANCES`: Shut down instances.
	// <li>`INVOKE_COMMAND`: Execute commands
	ActivityType *string `json:"ActivityType,omitnil,omitempty" name:"ActivityType"`

	// Scaling activity status. Value range:<br>
	// <li>INIT: initializing
	// <li>RUNNING: running
	// <li>SUCCESSFUL: succeeded
	// <li>PARTIALLY_SUCCESSFUL: partially succeeded
	// <li>FAILED: failed
	// <li>CANCELLED: canceled
	StatusCode *string `json:"StatusCode,omitnil,omitempty" name:"StatusCode"`

	// Description of the scaling activity status.
	StatusMessage *string `json:"StatusMessage,omitnil,omitempty" name:"StatusMessage"`

	// Cause of the scaling activity.
	Cause *string `json:"Cause,omitnil,omitempty" name:"Cause"`

	// Description of the scaling activity.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Start time of the scaling activity.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time of the scaling activity.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Creation time of the scaling activity.
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// This parameter has been deprecated.
	//
	// Deprecated: ActivityRelatedInstanceSet is deprecated.
	ActivityRelatedInstanceSet []*ActivtyRelatedInstance `json:"ActivityRelatedInstanceSet,omitnil,omitempty" name:"ActivityRelatedInstanceSet"`

	// Brief description of the scaling activity status.
	StatusMessageSimplified *string `json:"StatusMessageSimplified,omitnil,omitempty" name:"StatusMessageSimplified"`

	// Result of the lifecycle hook action in the scaling activity
	LifecycleActionResultSet []*LifecycleActionResultInfo `json:"LifecycleActionResultSet,omitnil,omitempty" name:"LifecycleActionResultSet"`

	// Detailed description of scaling activity status
	DetailedStatusMessageSet []*DetailedStatusMessage `json:"DetailedStatusMessageSet,omitnil,omitempty" name:"DetailedStatusMessageSet"`

	// Result of the command execution
	InvocationResultSet []*InvocationResult `json:"InvocationResultSet,omitnil,omitempty" name:"InvocationResultSet"`

	// Information set of the instances related to the scaling activity.
	RelatedInstanceSet []*RelatedInstance `json:"RelatedInstanceSet,omitnil,omitempty" name:"RelatedInstanceSet"`
}

type ActivtyRelatedInstance struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Status of the instance in the scaling activity. Value range:
	// <li>INIT: initializing
	// <li>RUNNING: running
	// <li>SUCCESSFUL: succeeded
	// <li>FAILED: failed
	InstanceStatus *string `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`
}

type Advice struct {
	// Problem Description
	Problem *string `json:"Problem,omitnil,omitempty" name:"Problem"`

	// Problem Details
	Detail *string `json:"Detail,omitnil,omitempty" name:"Detail"`

	// Recommended resolutions
	Solution *string `json:"Solution,omitnil,omitempty" name:"Solution"`

	// u200dRisk level of the scaling group configuration. Valid values: <br>
	// <li>WARNING<br>
	// <li>CRITICAL<br>
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`
}

// Predefined struct for user
type AttachInstancesRequestParams struct {
	// Auto scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// List of CVM instance IDs
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type AttachInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Auto scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// List of CVM instance IDs
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *AttachInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachInstancesResponseParams struct {
	// Scaling activity ID
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AttachInstancesResponse struct {
	*tchttp.BaseResponse
	Response *AttachInstancesResponseParams `json:"Response"`
}

func (r *AttachInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachLoadBalancersRequestParams struct {
	// Scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// List of classic CLB IDs. Up to 20 classic CLBs can be bound to a security group. `LoadBalancerIds` and `ForwardLoadBalancers` cannot be specified at the same time.
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// List of application CLBs. Up to 100 application CLBs can be bound to a scaling group. `LoadBalancerIds` and `ForwardLoadBalancers` cannot be specified at the same time.
	ForwardLoadBalancers []*ForwardLoadBalancer `json:"ForwardLoadBalancers,omitnil,omitempty" name:"ForwardLoadBalancers"`
}

type AttachLoadBalancersRequest struct {
	*tchttp.BaseRequest
	
	// Scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// List of classic CLB IDs. Up to 20 classic CLBs can be bound to a security group. `LoadBalancerIds` and `ForwardLoadBalancers` cannot be specified at the same time.
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// List of application CLBs. Up to 100 application CLBs can be bound to a scaling group. `LoadBalancerIds` and `ForwardLoadBalancers` cannot be specified at the same time.
	ForwardLoadBalancers []*ForwardLoadBalancer `json:"ForwardLoadBalancers,omitnil,omitempty" name:"ForwardLoadBalancers"`
}

func (r *AttachLoadBalancersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachLoadBalancersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "LoadBalancerIds")
	delete(f, "ForwardLoadBalancers")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachLoadBalancersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachLoadBalancersResponseParams struct {
	// Scaling activity ID
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AttachLoadBalancersResponse struct {
	*tchttp.BaseResponse
	Response *AttachLoadBalancersResponseParams `json:"Response"`
}

func (r *AttachLoadBalancersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachLoadBalancersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AutoScalingAdvice struct {
	// Scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// Scaling group warning level. Valid values:<br>
	// <li>NORMAL: Normal<br>
	// <li>WARNING: Warning<br>
	// <li>CRITICAL: Serious warning<br>
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// A collection of suggestions for scaling group configurations.
	Advices []*Advice `json:"Advices,omitnil,omitempty" name:"Advices"`
}

type AutoScalingGroup struct {
	// Auto scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// Auto scaling group name
	AutoScalingGroupName *string `json:"AutoScalingGroupName,omitnil,omitempty" name:"AutoScalingGroupName"`

	// Current scaling group status. Valid values:<br>
	// <li>NORMAL: Normal<br>
	// <li>CVM_ABNORMAL: Abnormal launch configuration<br>
	// <li>LB_ABNORMAL: Abnormal load balancer<br>
	// <li>LB_LISTENER_ABNORMAL: Abnormal load balancer listener<br>
	// <li>LB_LOCATION_ABNORMAL: Abnormal forwarding configuration of the load balancer listener<br>
	// <li>VPC_ABNORMAL: VPC network error<br>
	// <li>SUBNET_ABNORMAL: VPC subnet exception<br>
	// <li>INSUFFICIENT_BALANCE: Insufficient account balance<br>
	// <li>LB_BACKEND_REGION_NOT_MATCH: The CLB backend and the AS service are not in the same region.<br>
	// <li>LB_BACKEND_VPC_NOT_MATCH: The CLB instance and the scaling group are not in the same VPC.
	AutoScalingGroupStatus *string `json:"AutoScalingGroupStatus,omitnil,omitempty" name:"AutoScalingGroupStatus"`

	// Creation time in UTC format
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// Default cooldown period in seconds
	DefaultCooldown *int64 `json:"DefaultCooldown,omitnil,omitempty" name:"DefaultCooldown"`

	// Desired number of instances
	DesiredCapacity *int64 `json:"DesiredCapacity,omitnil,omitempty" name:"DesiredCapacity"`

	// Enabled status. Value range: `ENABLED`, `DISABLED`
	EnabledStatus *string `json:"EnabledStatus,omitnil,omitempty" name:"EnabledStatus"`

	// List of application load balancers
	ForwardLoadBalancerSet []*ForwardLoadBalancer `json:"ForwardLoadBalancerSet,omitnil,omitempty" name:"ForwardLoadBalancerSet"`

	// Number of instances
	InstanceCount *int64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// Number of instances in `IN_SERVICE` status
	InServiceInstanceCount *int64 `json:"InServiceInstanceCount,omitnil,omitempty" name:"InServiceInstanceCount"`

	// Launch configuration ID
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitnil,omitempty" name:"LaunchConfigurationId"`

	// Launch configuration name
	LaunchConfigurationName *string `json:"LaunchConfigurationName,omitnil,omitempty" name:"LaunchConfigurationName"`

	// List of Classic load balancer IDs
	LoadBalancerIdSet []*string `json:"LoadBalancerIdSet,omitnil,omitempty" name:"LoadBalancerIdSet"`

	// Maximum number of instances
	MaxSize *int64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`

	// Minimum number of instances
	MinSize *int64 `json:"MinSize,omitnil,omitempty" name:"MinSize"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// List of subnet IDs
	SubnetIdSet []*string `json:"SubnetIdSet,omitnil,omitempty" name:"SubnetIdSet"`

	// Termination policy
	TerminationPolicySet []*string `json:"TerminationPolicySet,omitnil,omitempty" name:"TerminationPolicySet"`

	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// List of availability zones
	ZoneSet []*string `json:"ZoneSet,omitnil,omitempty" name:"ZoneSet"`

	// Retry policy
	RetryPolicy *string `json:"RetryPolicy,omitnil,omitempty" name:"RetryPolicy"`

	// Whether the auto scaling group is performing a scaling activity. `IN_ACTIVITY` indicates yes, and `NOT_IN_ACTIVITY` indicates no.
	InActivityStatus *string `json:"InActivityStatus,omitnil,omitempty" name:"InActivityStatus"`

	// List of auto scaling group tags
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Service settings
	ServiceSettings *ServiceSettings `json:"ServiceSettings,omitnil,omitempty" name:"ServiceSettings"`

	// The number of IPv6 addresses that an instance has.
	Ipv6AddressCount *int64 `json:"Ipv6AddressCount,omitnil,omitempty" name:"Ipv6AddressCount"`

	// The policy applied when there are multiple availability zones/subnets
	// <br><li> PRIORITY: when creating instances, choose the availability zone/subnet based on the order in the list from top to bottom. If the first instance is successfully created in the availability zone/subnet of the highest priority, all instances will be created in this availability zone/subnet.
	// <br><li> EQUALITY: chooses the availability zone/subnet with the least instances for scale-out. This gives each availability zone/subnet an opportunity for scale-out and disperses the instances created during multiple scale-out operations across different availability zones/subnets.
	MultiZoneSubnetPolicy *string `json:"MultiZoneSubnetPolicy,omitnil,omitempty" name:"MultiZoneSubnetPolicy"`

	// Health check type of instances in a scaling group.<br><li>CVM: confirm whether an instance is healthy based on the network status. If the pinged instance is unreachable, the instance will be considered unhealthy. For more information, see [Instance Health Check](https://intl.cloud.tencent.com/document/product/377/8553?from_cn_redirect=1)<br><li>CLB: confirm whether an instance is healthy based on the CLB health check status. For more information, see [Health Check Overview](https://intl.cloud.tencent.com/document/product/214/6097?from_cn_redirect=1).
	HealthCheckType *string `json:"HealthCheckType,omitnil,omitempty" name:"HealthCheckType"`

	// Grace period of the CLB health check
	LoadBalancerHealthCheckGracePeriod *uint64 `json:"LoadBalancerHealthCheckGracePeriod,omitnil,omitempty" name:"LoadBalancerHealthCheckGracePeriod"`

	// Specifies how to assign instances. Valid values: `LAUNCH_CONFIGURATION` and `SPOT_MIXED`.
	// <br><li>`LAUNCH_CONFIGURATION`: the launch configuration mode.
	// <br><li>`SPOT_MIXED`: a mixed instance mode. Currently, this mode is supported only when the launch configuration takes the pay-as-you-go billing mode. With this mode, the scaling group can provision a combination of pay-as-you-go instances and spot instances to meet the configured capacity. Note that the billing mode of the associated launch configuration cannot be modified when this mode is used.
	InstanceAllocationPolicy *string `json:"InstanceAllocationPolicy,omitnil,omitempty" name:"InstanceAllocationPolicy"`

	// Specifies how to assign pay-as-you-go instances and spot instances.
	// A valid value will be returned only when `InstanceAllocationPolicy` is set to `SPOT_MIXED`.
	SpotMixedAllocationPolicy *SpotMixedAllocationPolicy `json:"SpotMixedAllocationPolicy,omitnil,omitempty" name:"SpotMixedAllocationPolicy"`

	// Indicates whether the capacity rebalancing feature is enabled. This parameter is only valid for spot instances in the scaling group. Valid values:
	// <br><li>`TRUE`: yes. Before the spot instances in the scaling group are about to be automatically repossessed, AS will terminate them. The scale-in hook (if configured) will take effect before the termination. After the termination process starts, AS will asynchronously initiate a scaling activity to meet the desired capacity.
	// <br><li>`FALSE`: no. AS will add instances to meet the desired capacity only after the spot instances are terminated.
	CapacityRebalance *bool `json:"CapacityRebalance,omitnil,omitempty" name:"CapacityRebalance"`
}

type AutoScalingGroupAbstract struct {
	// Scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// Auto scaling group name.
	AutoScalingGroupName *string `json:"AutoScalingGroupName,omitnil,omitempty" name:"AutoScalingGroupName"`
}

type AutoScalingNotification struct {
	// Auto scaling group ID.
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// List of user group IDs.
	NotificationUserGroupIds []*string `json:"NotificationUserGroupIds,omitnil,omitempty" name:"NotificationUserGroupIds"`

	// List of notification events.
	NotificationTypes []*string `json:"NotificationTypes,omitnil,omitempty" name:"NotificationTypes"`

	// Event notification ID.
	AutoScalingNotificationId *string `json:"AutoScalingNotificationId,omitnil,omitempty" name:"AutoScalingNotificationId"`

	// Notification receiver type.
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// CMQ queue name.
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// CMQ topic name.
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

// Predefined struct for user
type ClearLaunchConfigurationAttributesRequestParams struct {
	// Launch configuration ID
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitnil,omitempty" name:"LaunchConfigurationId"`

	// Whether to clear data disk information. This parameter is optional and the default value is `false`.
	// Setting it to `true` will clear data disks, which means that CVM newly created on this launch configuration will have no data disk.
	ClearDataDisks *bool `json:"ClearDataDisks,omitnil,omitempty" name:"ClearDataDisks"`

	// Whether to clear the CVM hostname settings. This parameter is optional and the default value is `false`.
	// Setting it to `true` will clear the hostname settings, which means that CVM newly created on this launch configuration will have no hostname.
	ClearHostNameSettings *bool `json:"ClearHostNameSettings,omitnil,omitempty" name:"ClearHostNameSettings"`

	// Whether to clear the CVM instance name settings. This parameter is optional and the default value is `false`.
	// Setting it to `true` will clear the instance name settings, which means that CVM newly created on this launch configuration will be named in the “as-{{AutoScalingGroupName}} format.
	ClearInstanceNameSettings *bool `json:"ClearInstanceNameSettings,omitnil,omitempty" name:"ClearInstanceNameSettings"`

	// Whether to clear placement group information. This parameter is optional. Default value: `false`.
	// `True` means clearing placement group information. After that, no placement groups are specified for CVMs created based on the information.
	ClearDisasterRecoverGroupIds *bool `json:"ClearDisasterRecoverGroupIds,omitnil,omitempty" name:"ClearDisasterRecoverGroupIds"`
}

type ClearLaunchConfigurationAttributesRequest struct {
	*tchttp.BaseRequest
	
	// Launch configuration ID
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitnil,omitempty" name:"LaunchConfigurationId"`

	// Whether to clear data disk information. This parameter is optional and the default value is `false`.
	// Setting it to `true` will clear data disks, which means that CVM newly created on this launch configuration will have no data disk.
	ClearDataDisks *bool `json:"ClearDataDisks,omitnil,omitempty" name:"ClearDataDisks"`

	// Whether to clear the CVM hostname settings. This parameter is optional and the default value is `false`.
	// Setting it to `true` will clear the hostname settings, which means that CVM newly created on this launch configuration will have no hostname.
	ClearHostNameSettings *bool `json:"ClearHostNameSettings,omitnil,omitempty" name:"ClearHostNameSettings"`

	// Whether to clear the CVM instance name settings. This parameter is optional and the default value is `false`.
	// Setting it to `true` will clear the instance name settings, which means that CVM newly created on this launch configuration will be named in the “as-{{AutoScalingGroupName}} format.
	ClearInstanceNameSettings *bool `json:"ClearInstanceNameSettings,omitnil,omitempty" name:"ClearInstanceNameSettings"`

	// Whether to clear placement group information. This parameter is optional. Default value: `false`.
	// `True` means clearing placement group information. After that, no placement groups are specified for CVMs created based on the information.
	ClearDisasterRecoverGroupIds *bool `json:"ClearDisasterRecoverGroupIds,omitnil,omitempty" name:"ClearDisasterRecoverGroupIds"`
}

func (r *ClearLaunchConfigurationAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClearLaunchConfigurationAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LaunchConfigurationId")
	delete(f, "ClearDataDisks")
	delete(f, "ClearHostNameSettings")
	delete(f, "ClearInstanceNameSettings")
	delete(f, "ClearDisasterRecoverGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ClearLaunchConfigurationAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ClearLaunchConfigurationAttributesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ClearLaunchConfigurationAttributesResponse struct {
	*tchttp.BaseResponse
	Response *ClearLaunchConfigurationAttributesResponseParams `json:"Response"`
}

func (r *ClearLaunchConfigurationAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClearLaunchConfigurationAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CompleteLifecycleActionRequestParams struct {
	// Lifecycle hook ID
	LifecycleHookId *string `json:"LifecycleHookId,omitnil,omitempty" name:"LifecycleHookId"`

	// Result of the lifecycle action. Value range: "CONTINUE", "ABANDON"
	LifecycleActionResult *string `json:"LifecycleActionResult,omitnil,omitempty" name:"LifecycleActionResult"`

	// Instance ID. Either "InstanceId" or "LifecycleActionToken" must be specified
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Either "InstanceId" or "LifecycleActionToken" must be specified
	LifecycleActionToken *string `json:"LifecycleActionToken,omitnil,omitempty" name:"LifecycleActionToken"`
}

type CompleteLifecycleActionRequest struct {
	*tchttp.BaseRequest
	
	// Lifecycle hook ID
	LifecycleHookId *string `json:"LifecycleHookId,omitnil,omitempty" name:"LifecycleHookId"`

	// Result of the lifecycle action. Value range: "CONTINUE", "ABANDON"
	LifecycleActionResult *string `json:"LifecycleActionResult,omitnil,omitempty" name:"LifecycleActionResult"`

	// Instance ID. Either "InstanceId" or "LifecycleActionToken" must be specified
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Either "InstanceId" or "LifecycleActionToken" must be specified
	LifecycleActionToken *string `json:"LifecycleActionToken,omitnil,omitempty" name:"LifecycleActionToken"`
}

func (r *CompleteLifecycleActionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CompleteLifecycleActionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LifecycleHookId")
	delete(f, "LifecycleActionResult")
	delete(f, "InstanceId")
	delete(f, "LifecycleActionToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CompleteLifecycleActionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CompleteLifecycleActionResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CompleteLifecycleActionResponse struct {
	*tchttp.BaseResponse
	Response *CompleteLifecycleActionResponseParams `json:"Response"`
}

func (r *CompleteLifecycleActionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CompleteLifecycleActionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAutoScalingGroupFromInstanceRequestParams struct {
	// The scaling group name. It must be unique under your account. The name can only contain letters, numbers, underscore, hyphen “-” and periods. It cannot exceed 55 bytes.
	AutoScalingGroupName *string `json:"AutoScalingGroupName,omitnil,omitempty" name:"AutoScalingGroupName"`

	// The instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The minimum number of instances. Value range: 0-2000.
	MinSize *int64 `json:"MinSize,omitnil,omitempty" name:"MinSize"`

	// The maximum number of instances. Value range: 0-2000.
	MaxSize *int64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`

	// The desired capacity. Its value must be greater than the minimum and smaller than the maximum.
	DesiredCapacity *int64 `json:"DesiredCapacity,omitnil,omitempty" name:"DesiredCapacity"`

	// Whether to inherit the instance tag. Default value: False
	InheritInstanceTag *bool `json:"InheritInstanceTag,omitnil,omitempty" name:"InheritInstanceTag"`
}

type CreateAutoScalingGroupFromInstanceRequest struct {
	*tchttp.BaseRequest
	
	// The scaling group name. It must be unique under your account. The name can only contain letters, numbers, underscore, hyphen “-” and periods. It cannot exceed 55 bytes.
	AutoScalingGroupName *string `json:"AutoScalingGroupName,omitnil,omitempty" name:"AutoScalingGroupName"`

	// The instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The minimum number of instances. Value range: 0-2000.
	MinSize *int64 `json:"MinSize,omitnil,omitempty" name:"MinSize"`

	// The maximum number of instances. Value range: 0-2000.
	MaxSize *int64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`

	// The desired capacity. Its value must be greater than the minimum and smaller than the maximum.
	DesiredCapacity *int64 `json:"DesiredCapacity,omitnil,omitempty" name:"DesiredCapacity"`

	// Whether to inherit the instance tag. Default value: False
	InheritInstanceTag *bool `json:"InheritInstanceTag,omitnil,omitempty" name:"InheritInstanceTag"`
}

func (r *CreateAutoScalingGroupFromInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAutoScalingGroupFromInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupName")
	delete(f, "InstanceId")
	delete(f, "MinSize")
	delete(f, "MaxSize")
	delete(f, "DesiredCapacity")
	delete(f, "InheritInstanceTag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAutoScalingGroupFromInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAutoScalingGroupFromInstanceResponseParams struct {
	// The scaling group ID.
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAutoScalingGroupFromInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateAutoScalingGroupFromInstanceResponseParams `json:"Response"`
}

func (r *CreateAutoScalingGroupFromInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAutoScalingGroupFromInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAutoScalingGroupRequestParams struct {
	// Auto scaling group name, which can only contain letters, numbers, underscores, hyphens ("-"), and decimal points with a maximum length of 55 bytes and must be unique under your account.
	AutoScalingGroupName *string `json:"AutoScalingGroupName,omitnil,omitempty" name:"AutoScalingGroupName"`

	// Launch configuration ID
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitnil,omitempty" name:"LaunchConfigurationId"`

	// Maximum number of instances. Value range: 0-2,000.
	MaxSize *uint64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`

	// Minimum number of instances. Value range: 0-2,000.
	MinSize *uint64 `json:"MinSize,omitnil,omitempty" name:"MinSize"`

	// VPC ID; if on a basic network, enter an empty string
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Default cooldown period in seconds. Default value: 300
	DefaultCooldown *uint64 `json:"DefaultCooldown,omitnil,omitempty" name:"DefaultCooldown"`

	// Desired number of instances. The number should be no larger than the maximum and no smaller than minimum number of instances
	DesiredCapacity *uint64 `json:"DesiredCapacity,omitnil,omitempty" name:"DesiredCapacity"`

	// List of classic CLB IDs. Currently, the maximum length is 20. You cannot specify LoadBalancerIds and ForwardLoadBalancers at the same time.
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// Project ID of an instance in a scaling group. The default project is used if it’s left blank.
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// List of application CLBs. Up to 100 CLBs are allowed. `LoadBalancerIds` and `ForwardLoadBalancers` cannot be specified at the same time.
	ForwardLoadBalancers []*ForwardLoadBalancer `json:"ForwardLoadBalancers,omitnil,omitempty" name:"ForwardLoadBalancers"`

	// List of subnet IDs. A subnet must be specified in the VPC scenario. If multiple subnets are entered, their priority will be determined by the order in which they are entered, and they will be tried one by one until instances can be successfully created.
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// Termination policy. Currently, the maximum length is 1. Value range: OLDEST_INSTANCE, NEWEST_INSTANCE. Default value: OLDEST_INSTANCE.
	// <br><li> OLDEST_INSTANCE: The oldest instance in the auto scaling group will be terminated first.
	// <br><li> NEWEST_INSTANCE: The newest instance in the auto scaling group will be terminated first.
	TerminationPolicies []*string `json:"TerminationPolicies,omitnil,omitempty" name:"TerminationPolicies"`

	// List of availability zones. An availability zone must be specified in the basic network scenario. If multiple availability zones are entered, their priority will be determined by the order in which they are entered, and they will be tried one by one until instances can be successfully created.
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// Retry policy. Valid values: `IMMEDIATE_RETRY` (default), `INCREMENTAL_INTERVALS`, `NO_RETRY`. A partially successful scaling is judged as a failed one.
	// <br><li>`IMMEDIATE_RETRY`: Retry immediately. Stop retrying after five consecutive failures.
	// <br><li>`INCREMENTAL_INTERVALS`: Retry at incremental intervals. As the number of consecutive failures increases, the retry interval gradually increases, ranging from seconds to one day.
	// <br><li>`NO_RETRY`: Do not retry. Actions are taken when the next call or alarm message comes.
	RetryPolicy *string `json:"RetryPolicy,omitnil,omitempty" name:"RetryPolicy"`

	// Availability zone verification policy. Value range: ALL, ANY. Default value: ANY.
	// <br><li> ALL: The verification will succeed only if all availability zones (Zone) or subnets (SubnetId) are available; otherwise, an error will be reported.
	// <br><li> ANY: The verification will success if any availability zone (Zone) or subnet (SubnetId) is available; otherwise, an error will be reported.
	// 
	// Common reasons why an availability zone or subnet is unavailable include stock-out of CVM instances or CBS cloud disks in the availability zone, insufficient quota in the availability zone, or insufficient IPs in the subnet.
	// If an availability zone or subnet in Zones/SubnetIds does not exist, a verification error will be reported regardless of the value of ZonesCheckPolicy.
	ZonesCheckPolicy *string `json:"ZonesCheckPolicy,omitnil,omitempty" name:"ZonesCheckPolicy"`

	// List of tag descriptions. In this parameter, you can specify the tags to be bound with a scaling group as well as corresponding resource instances. Each scaling group can have up to 30 tags.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Service settings such as unhealthy instance replacement.
	ServiceSettings *ServiceSettings `json:"ServiceSettings,omitnil,omitempty" name:"ServiceSettings"`

	// The number of IPv6 addresses that an instance has. Valid values: 0 and 1. Default value: 0.
	Ipv6AddressCount *int64 `json:"Ipv6AddressCount,omitnil,omitempty" name:"Ipv6AddressCount"`

	// Multi-availability zone/subnet policy. Valid values: PRIORITY and EQUALITY. Default value: PRIORITY.
	// <br><li> PRIORITY: when creating instances, choose the availability zone/subnet based on the order in the list from top to bottom. If the first instance is successfully created in the availability zone/subnet of the highest priority, all instances will be created in this availability zone/subnet.
	// <br><li>EQUALITY: instances created for scale-out are distributed to multiple availability zones/subnets, so as to keep the number of instances in different availability zone/subnet in balance.
	// 
	// Notes: 
	// <br><li> When the scaling group is based on the classic network, this policy applies to multiple availability zones. When the scaling group is based on a VPC, this policy applies to multiple subnets, and you do not need to consider availability zones. For example, if you have four subnets (A, B, C, and D) and A, B, and C are in availability zone 1 and D is in availability zone 2, you only need to decide the order of the four subnets, without worrying about the issue of availability zones.
	// <br><li> This policy is applicable to multiple availability zones/subnets, but is not applicable to multiple models with launch configurations. Specify the models according to the model priority.
	// <br><li> When creating instances based on the PRIORITY policy, apply the multi-model policy and then apply the multi-availability zones/subnet policy. For example, if you have models A and B and subnets 1, 2, and 3, creation will be attempted in the following order: A1, A2, A3, B1, B2, and B3. If A1 is sold out, A2 (not B1) is tried next.
	MultiZoneSubnetPolicy *string `json:"MultiZoneSubnetPolicy,omitnil,omitempty" name:"MultiZoneSubnetPolicy"`

	// Health check type of instances in a scaling group.<br><li>CVM: confirm whether an instance is healthy based on the network status. If the pinged instance is unreachable, the instance will be considered unhealthy. For more information, see [Instance Health Check](https://intl.cloud.tencent.com/document/product/377/8553?from_cn_redirect=1)<br><li>CLB: confirm whether an instance is healthy based on the CLB health check status. For more information, see [Health Check Overview](https://intl.cloud.tencent.com/document/product/214/6097?from_cn_redirect=1).<br>If the parameter is set to `CLB`, the scaling group will check both the network status and the CLB health check status. If the network check indicates unhealthy, the `HealthStatus` field will return `UNHEALTHY`. If the CLB health check indicates unhealthy, the `HealthStatus` field will return `CLB_UNHEALTHY`. If both checks indicate unhealthy, the `HealthStatus` field will return `UNHEALTHY|CLB_UNHEALTHY`. Default value: `CLB`.
	HealthCheckType *string `json:"HealthCheckType,omitnil,omitempty" name:"HealthCheckType"`

	// Grace period of the CLB health check during which the `IN_SERVICE` instances added will not be marked as `CLB_UNHEALTHY`.<br>Valid range: 0-7200, in seconds. Default value: `0`.
	LoadBalancerHealthCheckGracePeriod *uint64 `json:"LoadBalancerHealthCheckGracePeriod,omitnil,omitempty" name:"LoadBalancerHealthCheckGracePeriod"`

	// Specifies how to assign instances. Valid values: `LAUNCH_CONFIGURATION` and `SPOT_MIXED`; default value: `LAUNCH_CONFIGURATION`.
	// <br><li>`LAUNCH_CONFIGURATION`: the launch configuration mode.
	// <br><li>`SPOT_MIXED`: a mixed instance mode. Currently, this mode is supported only when the launch configuration takes the pay-as-you-go billing mode. With this mode, the scaling group can provision a combination of pay-as-you-go instances and spot instances to meet the configured capacity. Note that the billing mode of the associated launch configuration cannot be modified when this mode is used.
	InstanceAllocationPolicy *string `json:"InstanceAllocationPolicy,omitnil,omitempty" name:"InstanceAllocationPolicy"`

	// Specifies how to assign pay-as-you-go instances and spot instances.
	// This parameter is valid only when `InstanceAllocationPolicy ` is set to `SPOT_MIXED`.
	SpotMixedAllocationPolicy *SpotMixedAllocationPolicy `json:"SpotMixedAllocationPolicy,omitnil,omitempty" name:"SpotMixedAllocationPolicy"`

	// Indicates whether the capacity re-balancing feature is enabled. This parameter is only valid for spot instances in the scaling group. Valid values:
	// <br><li>`TRUE`: yes. Before the spot instances in the scaling group are about to be automatically repossessed, AS will terminate them. The scale-in hook (if configured) will take effect before the termination. After the termination process starts, AS will asynchronously initiate a scaling activity to meet the desired capacity.
	// <br><li>`FALSE`: no. In this case, AS will add instances to meet the desired capacity only after the spot instances are terminated.
	// 
	// Default value: `False`.
	CapacityRebalance *bool `json:"CapacityRebalance,omitnil,omitempty" name:"CapacityRebalance"`
}

type CreateAutoScalingGroupRequest struct {
	*tchttp.BaseRequest
	
	// Auto scaling group name, which can only contain letters, numbers, underscores, hyphens ("-"), and decimal points with a maximum length of 55 bytes and must be unique under your account.
	AutoScalingGroupName *string `json:"AutoScalingGroupName,omitnil,omitempty" name:"AutoScalingGroupName"`

	// Launch configuration ID
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitnil,omitempty" name:"LaunchConfigurationId"`

	// Maximum number of instances. Value range: 0-2,000.
	MaxSize *uint64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`

	// Minimum number of instances. Value range: 0-2,000.
	MinSize *uint64 `json:"MinSize,omitnil,omitempty" name:"MinSize"`

	// VPC ID; if on a basic network, enter an empty string
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Default cooldown period in seconds. Default value: 300
	DefaultCooldown *uint64 `json:"DefaultCooldown,omitnil,omitempty" name:"DefaultCooldown"`

	// Desired number of instances. The number should be no larger than the maximum and no smaller than minimum number of instances
	DesiredCapacity *uint64 `json:"DesiredCapacity,omitnil,omitempty" name:"DesiredCapacity"`

	// List of classic CLB IDs. Currently, the maximum length is 20. You cannot specify LoadBalancerIds and ForwardLoadBalancers at the same time.
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// Project ID of an instance in a scaling group. The default project is used if it’s left blank.
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// List of application CLBs. Up to 100 CLBs are allowed. `LoadBalancerIds` and `ForwardLoadBalancers` cannot be specified at the same time.
	ForwardLoadBalancers []*ForwardLoadBalancer `json:"ForwardLoadBalancers,omitnil,omitempty" name:"ForwardLoadBalancers"`

	// List of subnet IDs. A subnet must be specified in the VPC scenario. If multiple subnets are entered, their priority will be determined by the order in which they are entered, and they will be tried one by one until instances can be successfully created.
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// Termination policy. Currently, the maximum length is 1. Value range: OLDEST_INSTANCE, NEWEST_INSTANCE. Default value: OLDEST_INSTANCE.
	// <br><li> OLDEST_INSTANCE: The oldest instance in the auto scaling group will be terminated first.
	// <br><li> NEWEST_INSTANCE: The newest instance in the auto scaling group will be terminated first.
	TerminationPolicies []*string `json:"TerminationPolicies,omitnil,omitempty" name:"TerminationPolicies"`

	// List of availability zones. An availability zone must be specified in the basic network scenario. If multiple availability zones are entered, their priority will be determined by the order in which they are entered, and they will be tried one by one until instances can be successfully created.
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// Retry policy. Valid values: `IMMEDIATE_RETRY` (default), `INCREMENTAL_INTERVALS`, `NO_RETRY`. A partially successful scaling is judged as a failed one.
	// <br><li>`IMMEDIATE_RETRY`: Retry immediately. Stop retrying after five consecutive failures.
	// <br><li>`INCREMENTAL_INTERVALS`: Retry at incremental intervals. As the number of consecutive failures increases, the retry interval gradually increases, ranging from seconds to one day.
	// <br><li>`NO_RETRY`: Do not retry. Actions are taken when the next call or alarm message comes.
	RetryPolicy *string `json:"RetryPolicy,omitnil,omitempty" name:"RetryPolicy"`

	// Availability zone verification policy. Value range: ALL, ANY. Default value: ANY.
	// <br><li> ALL: The verification will succeed only if all availability zones (Zone) or subnets (SubnetId) are available; otherwise, an error will be reported.
	// <br><li> ANY: The verification will success if any availability zone (Zone) or subnet (SubnetId) is available; otherwise, an error will be reported.
	// 
	// Common reasons why an availability zone or subnet is unavailable include stock-out of CVM instances or CBS cloud disks in the availability zone, insufficient quota in the availability zone, or insufficient IPs in the subnet.
	// If an availability zone or subnet in Zones/SubnetIds does not exist, a verification error will be reported regardless of the value of ZonesCheckPolicy.
	ZonesCheckPolicy *string `json:"ZonesCheckPolicy,omitnil,omitempty" name:"ZonesCheckPolicy"`

	// List of tag descriptions. In this parameter, you can specify the tags to be bound with a scaling group as well as corresponding resource instances. Each scaling group can have up to 30 tags.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Service settings such as unhealthy instance replacement.
	ServiceSettings *ServiceSettings `json:"ServiceSettings,omitnil,omitempty" name:"ServiceSettings"`

	// The number of IPv6 addresses that an instance has. Valid values: 0 and 1. Default value: 0.
	Ipv6AddressCount *int64 `json:"Ipv6AddressCount,omitnil,omitempty" name:"Ipv6AddressCount"`

	// Multi-availability zone/subnet policy. Valid values: PRIORITY and EQUALITY. Default value: PRIORITY.
	// <br><li> PRIORITY: when creating instances, choose the availability zone/subnet based on the order in the list from top to bottom. If the first instance is successfully created in the availability zone/subnet of the highest priority, all instances will be created in this availability zone/subnet.
	// <br><li>EQUALITY: instances created for scale-out are distributed to multiple availability zones/subnets, so as to keep the number of instances in different availability zone/subnet in balance.
	// 
	// Notes: 
	// <br><li> When the scaling group is based on the classic network, this policy applies to multiple availability zones. When the scaling group is based on a VPC, this policy applies to multiple subnets, and you do not need to consider availability zones. For example, if you have four subnets (A, B, C, and D) and A, B, and C are in availability zone 1 and D is in availability zone 2, you only need to decide the order of the four subnets, without worrying about the issue of availability zones.
	// <br><li> This policy is applicable to multiple availability zones/subnets, but is not applicable to multiple models with launch configurations. Specify the models according to the model priority.
	// <br><li> When creating instances based on the PRIORITY policy, apply the multi-model policy and then apply the multi-availability zones/subnet policy. For example, if you have models A and B and subnets 1, 2, and 3, creation will be attempted in the following order: A1, A2, A3, B1, B2, and B3. If A1 is sold out, A2 (not B1) is tried next.
	MultiZoneSubnetPolicy *string `json:"MultiZoneSubnetPolicy,omitnil,omitempty" name:"MultiZoneSubnetPolicy"`

	// Health check type of instances in a scaling group.<br><li>CVM: confirm whether an instance is healthy based on the network status. If the pinged instance is unreachable, the instance will be considered unhealthy. For more information, see [Instance Health Check](https://intl.cloud.tencent.com/document/product/377/8553?from_cn_redirect=1)<br><li>CLB: confirm whether an instance is healthy based on the CLB health check status. For more information, see [Health Check Overview](https://intl.cloud.tencent.com/document/product/214/6097?from_cn_redirect=1).<br>If the parameter is set to `CLB`, the scaling group will check both the network status and the CLB health check status. If the network check indicates unhealthy, the `HealthStatus` field will return `UNHEALTHY`. If the CLB health check indicates unhealthy, the `HealthStatus` field will return `CLB_UNHEALTHY`. If both checks indicate unhealthy, the `HealthStatus` field will return `UNHEALTHY|CLB_UNHEALTHY`. Default value: `CLB`.
	HealthCheckType *string `json:"HealthCheckType,omitnil,omitempty" name:"HealthCheckType"`

	// Grace period of the CLB health check during which the `IN_SERVICE` instances added will not be marked as `CLB_UNHEALTHY`.<br>Valid range: 0-7200, in seconds. Default value: `0`.
	LoadBalancerHealthCheckGracePeriod *uint64 `json:"LoadBalancerHealthCheckGracePeriod,omitnil,omitempty" name:"LoadBalancerHealthCheckGracePeriod"`

	// Specifies how to assign instances. Valid values: `LAUNCH_CONFIGURATION` and `SPOT_MIXED`; default value: `LAUNCH_CONFIGURATION`.
	// <br><li>`LAUNCH_CONFIGURATION`: the launch configuration mode.
	// <br><li>`SPOT_MIXED`: a mixed instance mode. Currently, this mode is supported only when the launch configuration takes the pay-as-you-go billing mode. With this mode, the scaling group can provision a combination of pay-as-you-go instances and spot instances to meet the configured capacity. Note that the billing mode of the associated launch configuration cannot be modified when this mode is used.
	InstanceAllocationPolicy *string `json:"InstanceAllocationPolicy,omitnil,omitempty" name:"InstanceAllocationPolicy"`

	// Specifies how to assign pay-as-you-go instances and spot instances.
	// This parameter is valid only when `InstanceAllocationPolicy ` is set to `SPOT_MIXED`.
	SpotMixedAllocationPolicy *SpotMixedAllocationPolicy `json:"SpotMixedAllocationPolicy,omitnil,omitempty" name:"SpotMixedAllocationPolicy"`

	// Indicates whether the capacity re-balancing feature is enabled. This parameter is only valid for spot instances in the scaling group. Valid values:
	// <br><li>`TRUE`: yes. Before the spot instances in the scaling group are about to be automatically repossessed, AS will terminate them. The scale-in hook (if configured) will take effect before the termination. After the termination process starts, AS will asynchronously initiate a scaling activity to meet the desired capacity.
	// <br><li>`FALSE`: no. In this case, AS will add instances to meet the desired capacity only after the spot instances are terminated.
	// 
	// Default value: `False`.
	CapacityRebalance *bool `json:"CapacityRebalance,omitnil,omitempty" name:"CapacityRebalance"`
}

func (r *CreateAutoScalingGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAutoScalingGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupName")
	delete(f, "LaunchConfigurationId")
	delete(f, "MaxSize")
	delete(f, "MinSize")
	delete(f, "VpcId")
	delete(f, "DefaultCooldown")
	delete(f, "DesiredCapacity")
	delete(f, "LoadBalancerIds")
	delete(f, "ProjectId")
	delete(f, "ForwardLoadBalancers")
	delete(f, "SubnetIds")
	delete(f, "TerminationPolicies")
	delete(f, "Zones")
	delete(f, "RetryPolicy")
	delete(f, "ZonesCheckPolicy")
	delete(f, "Tags")
	delete(f, "ServiceSettings")
	delete(f, "Ipv6AddressCount")
	delete(f, "MultiZoneSubnetPolicy")
	delete(f, "HealthCheckType")
	delete(f, "LoadBalancerHealthCheckGracePeriod")
	delete(f, "InstanceAllocationPolicy")
	delete(f, "SpotMixedAllocationPolicy")
	delete(f, "CapacityRebalance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAutoScalingGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAutoScalingGroupResponseParams struct {
	// Auto scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAutoScalingGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateAutoScalingGroupResponseParams `json:"Response"`
}

func (r *CreateAutoScalingGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAutoScalingGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLaunchConfigurationRequestParams struct {
	// Display name of the launch configuration, which can contain letters, digits, underscores and hyphens (-), and dots. Up to of 60 bytes allowed..
	LaunchConfigurationName *string `json:"LaunchConfigurationName,omitnil,omitempty" name:"LaunchConfigurationName"`

	// [Image](https://intl.cloud.tencent.com/document/product/213/4940?from_cn_redirect=1) ID in the format of `img-xxx`. There are three types of images: <br/><li>Public images </li><li>Custom images </li><li>Shared images </li><br/>You can obtain the image IDs in the [CVM console](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE).</li><li>You can also use the [DescribeImages](https://intl.cloud.tencent.com/document/api/213/15715?from_cn_redirect=1) and look for `ImageId` in the response.</li>
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// Project ID of the launch configuration. The default project is used if it’s left blank.
	// Note that this project ID is not the same as the project ID of the scaling group. 
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Instance model. Different instance models specify different resource specifications. The specific value can be obtained by calling the [DescribeInstanceTypeConfigs](https://intl.cloud.tencent.com/document/api/213/15749?from_cn_redirect=1) API to get the latest specification table or referring to the descriptions in [Instance Types](https://intl.cloud.tencent.com/document/product/213/11518?from_cn_redirect=1).
	// `InstanceType` and `InstanceTypes` are mutually exclusive, and one and only one of them must be entered.
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// System disk configuration of the instance. If this parameter is not specified, the default value will be used.
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// Information of the instance's data disk configuration. If this parameter is not specified, no data disk is purchased by default. Up to 11 data disks can be supported.
	DataDisks []*DataDisk `json:"DataDisks,omitnil,omitempty" name:"DataDisks"`

	// Configuration of public network bandwidth. If this parameter is not specified, 0 Mbps will be used by default.
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// Login settings of the instance. You can use this parameter to set the login method, password, and key of the instance or keep the login settings of the original image. By default, a random password will be generated and sent to you via the Message Center.
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// The security group to which the instance belongs. This parameter can be obtained by calling the `SecurityGroupId` field in the returned value of [DescribeSecurityGroups](https://intl.cloud.tencent.com/document/api/215/15808?from_cn_redirect=1). If this parameter is not specified, no security group will be bound by default.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Enhanced services. You can use this parameter to specify whether to enable services such as Cloud Security and Cloud Monitor. If this parameter is not specified, Cloud Monitor and Cloud Security will be enabled by default.
	EnhancedService *EnhancedService `json:"EnhancedService,omitnil,omitempty" name:"EnhancedService"`

	// Base64-encoded custom data of up to 16 KB.
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`

	// Instance billing mode. CVM instances take `POSTPAID_BY_HOUR` by default. Valid values:
	// <br><li>POSTPAID_BY_HOUR: pay-as-you-go hourly
	// <br><li>SPOTPAID: spot instance
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// Market options of the instance, such as parameters related to spot instances. This parameter is required for spot instances.
	InstanceMarketOptions *InstanceMarketOptionsRequest `json:"InstanceMarketOptions,omitnil,omitempty" name:"InstanceMarketOptions"`

	// List of instance models. Different instance models specify different resource specifications. Up to 10 instance models can be supported.
	// `InstanceType` and `InstanceTypes` are mutually exclusive, and one and only one of them must be entered.
	InstanceTypes []*string `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// CAM role name. This parameter can be obtained from the `roleName` field returned by DescribeRoleList API.
	CamRoleName *string `json:"CamRoleName,omitnil,omitempty" name:"CamRoleName"`

	// Instance type verification policy. Value range: ALL, ANY. Default value: ANY.
	// <br><li> ALL: The verification will success only if all instance types (InstanceType) are available; otherwise, an error will be reported.
	// <br><li> ANY: The verification will success if any instance type (InstanceType) is available; otherwise, an error will be reported.
	// 
	// Common reasons why an instance type is unavailable include stock-out of the instance type or the corresponding cloud disk.
	// If a model in InstanceTypes does not exist or has been discontinued, a verification error will be reported regardless of the value of InstanceTypesCheckPolicy.
	InstanceTypesCheckPolicy *string `json:"InstanceTypesCheckPolicy,omitnil,omitempty" name:"InstanceTypesCheckPolicy"`

	// List of tags. This parameter is used to bind up to 10 tags to newly added instances.
	InstanceTags []*InstanceTag `json:"InstanceTags,omitnil,omitempty" name:"InstanceTags"`

	// List of tags. You can specify tags that you want to bind to the launch configuration. Each launch configuration can have up to 30 tags.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// CVM hostname settings.
	HostNameSettings *HostNameSettings `json:"HostNameSettings,omitnil,omitempty" name:"HostNameSettings"`

	// Settings of CVM instance names
	// If this field is configured in a launch configuration, the `InstanceName` of a CVM created by the scaling group will be generated according to the configuration; otherwise, it will be in the `as-{{AutoScalingGroupName }}` format.
	InstanceNameSettings *InstanceNameSettings `json:"InstanceNameSettings,omitnil,omitempty" name:"InstanceNameSettings"`

	// Details of the monthly subscription, including the purchase period, auto-renewal. It is required if the `InstanceChargeType` is `PREPAID`.
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// Selection policy of cloud disks. Default value: ORIGINAL. Valid values:
	// <br><li>ORIGINAL: uses the configured cloud disk type
	// <br><li>AUTOMATIC: automatically chooses an available cloud disk type
	DiskTypePolicy *string `json:"DiskTypePolicy,omitnil,omitempty" name:"DiskTypePolicy"`

	// HPC ID<br>
	// Note: This field is default to empty
	HpcClusterId *string `json:"HpcClusterId,omitnil,omitempty" name:"HpcClusterId"`

	// IPv6 public network bandwidth configuration. If the IPv6 address is available in the new instance, public network bandwidth can be allocated to the IPv6 address. This parameter is invalid if `Ipv6AddressCount` of the scaling group associated with the launch configuration is 0.
	IPv6InternetAccessible *IPv6InternetAccessible `json:"IPv6InternetAccessible,omitnil,omitempty" name:"IPv6InternetAccessible"`

	// Placement group ID. Only one is allowed.
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil,omitempty" name:"DisasterRecoverGroupIds"`
}

type CreateLaunchConfigurationRequest struct {
	*tchttp.BaseRequest
	
	// Display name of the launch configuration, which can contain letters, digits, underscores and hyphens (-), and dots. Up to of 60 bytes allowed..
	LaunchConfigurationName *string `json:"LaunchConfigurationName,omitnil,omitempty" name:"LaunchConfigurationName"`

	// [Image](https://intl.cloud.tencent.com/document/product/213/4940?from_cn_redirect=1) ID in the format of `img-xxx`. There are three types of images: <br/><li>Public images </li><li>Custom images </li><li>Shared images </li><br/>You can obtain the image IDs in the [CVM console](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE).</li><li>You can also use the [DescribeImages](https://intl.cloud.tencent.com/document/api/213/15715?from_cn_redirect=1) and look for `ImageId` in the response.</li>
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// Project ID of the launch configuration. The default project is used if it’s left blank.
	// Note that this project ID is not the same as the project ID of the scaling group. 
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Instance model. Different instance models specify different resource specifications. The specific value can be obtained by calling the [DescribeInstanceTypeConfigs](https://intl.cloud.tencent.com/document/api/213/15749?from_cn_redirect=1) API to get the latest specification table or referring to the descriptions in [Instance Types](https://intl.cloud.tencent.com/document/product/213/11518?from_cn_redirect=1).
	// `InstanceType` and `InstanceTypes` are mutually exclusive, and one and only one of them must be entered.
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// System disk configuration of the instance. If this parameter is not specified, the default value will be used.
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// Information of the instance's data disk configuration. If this parameter is not specified, no data disk is purchased by default. Up to 11 data disks can be supported.
	DataDisks []*DataDisk `json:"DataDisks,omitnil,omitempty" name:"DataDisks"`

	// Configuration of public network bandwidth. If this parameter is not specified, 0 Mbps will be used by default.
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// Login settings of the instance. You can use this parameter to set the login method, password, and key of the instance or keep the login settings of the original image. By default, a random password will be generated and sent to you via the Message Center.
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// The security group to which the instance belongs. This parameter can be obtained by calling the `SecurityGroupId` field in the returned value of [DescribeSecurityGroups](https://intl.cloud.tencent.com/document/api/215/15808?from_cn_redirect=1). If this parameter is not specified, no security group will be bound by default.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Enhanced services. You can use this parameter to specify whether to enable services such as Cloud Security and Cloud Monitor. If this parameter is not specified, Cloud Monitor and Cloud Security will be enabled by default.
	EnhancedService *EnhancedService `json:"EnhancedService,omitnil,omitempty" name:"EnhancedService"`

	// Base64-encoded custom data of up to 16 KB.
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`

	// Instance billing mode. CVM instances take `POSTPAID_BY_HOUR` by default. Valid values:
	// <br><li>POSTPAID_BY_HOUR: pay-as-you-go hourly
	// <br><li>SPOTPAID: spot instance
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// Market options of the instance, such as parameters related to spot instances. This parameter is required for spot instances.
	InstanceMarketOptions *InstanceMarketOptionsRequest `json:"InstanceMarketOptions,omitnil,omitempty" name:"InstanceMarketOptions"`

	// List of instance models. Different instance models specify different resource specifications. Up to 10 instance models can be supported.
	// `InstanceType` and `InstanceTypes` are mutually exclusive, and one and only one of them must be entered.
	InstanceTypes []*string `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// CAM role name. This parameter can be obtained from the `roleName` field returned by DescribeRoleList API.
	CamRoleName *string `json:"CamRoleName,omitnil,omitempty" name:"CamRoleName"`

	// Instance type verification policy. Value range: ALL, ANY. Default value: ANY.
	// <br><li> ALL: The verification will success only if all instance types (InstanceType) are available; otherwise, an error will be reported.
	// <br><li> ANY: The verification will success if any instance type (InstanceType) is available; otherwise, an error will be reported.
	// 
	// Common reasons why an instance type is unavailable include stock-out of the instance type or the corresponding cloud disk.
	// If a model in InstanceTypes does not exist or has been discontinued, a verification error will be reported regardless of the value of InstanceTypesCheckPolicy.
	InstanceTypesCheckPolicy *string `json:"InstanceTypesCheckPolicy,omitnil,omitempty" name:"InstanceTypesCheckPolicy"`

	// List of tags. This parameter is used to bind up to 10 tags to newly added instances.
	InstanceTags []*InstanceTag `json:"InstanceTags,omitnil,omitempty" name:"InstanceTags"`

	// List of tags. You can specify tags that you want to bind to the launch configuration. Each launch configuration can have up to 30 tags.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// CVM hostname settings.
	HostNameSettings *HostNameSettings `json:"HostNameSettings,omitnil,omitempty" name:"HostNameSettings"`

	// Settings of CVM instance names
	// If this field is configured in a launch configuration, the `InstanceName` of a CVM created by the scaling group will be generated according to the configuration; otherwise, it will be in the `as-{{AutoScalingGroupName }}` format.
	InstanceNameSettings *InstanceNameSettings `json:"InstanceNameSettings,omitnil,omitempty" name:"InstanceNameSettings"`

	// Details of the monthly subscription, including the purchase period, auto-renewal. It is required if the `InstanceChargeType` is `PREPAID`.
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// Selection policy of cloud disks. Default value: ORIGINAL. Valid values:
	// <br><li>ORIGINAL: uses the configured cloud disk type
	// <br><li>AUTOMATIC: automatically chooses an available cloud disk type
	DiskTypePolicy *string `json:"DiskTypePolicy,omitnil,omitempty" name:"DiskTypePolicy"`

	// HPC ID<br>
	// Note: This field is default to empty
	HpcClusterId *string `json:"HpcClusterId,omitnil,omitempty" name:"HpcClusterId"`

	// IPv6 public network bandwidth configuration. If the IPv6 address is available in the new instance, public network bandwidth can be allocated to the IPv6 address. This parameter is invalid if `Ipv6AddressCount` of the scaling group associated with the launch configuration is 0.
	IPv6InternetAccessible *IPv6InternetAccessible `json:"IPv6InternetAccessible,omitnil,omitempty" name:"IPv6InternetAccessible"`

	// Placement group ID. Only one is allowed.
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil,omitempty" name:"DisasterRecoverGroupIds"`
}

func (r *CreateLaunchConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLaunchConfigurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LaunchConfigurationName")
	delete(f, "ImageId")
	delete(f, "ProjectId")
	delete(f, "InstanceType")
	delete(f, "SystemDisk")
	delete(f, "DataDisks")
	delete(f, "InternetAccessible")
	delete(f, "LoginSettings")
	delete(f, "SecurityGroupIds")
	delete(f, "EnhancedService")
	delete(f, "UserData")
	delete(f, "InstanceChargeType")
	delete(f, "InstanceMarketOptions")
	delete(f, "InstanceTypes")
	delete(f, "CamRoleName")
	delete(f, "InstanceTypesCheckPolicy")
	delete(f, "InstanceTags")
	delete(f, "Tags")
	delete(f, "HostNameSettings")
	delete(f, "InstanceNameSettings")
	delete(f, "InstanceChargePrepaid")
	delete(f, "DiskTypePolicy")
	delete(f, "HpcClusterId")
	delete(f, "IPv6InternetAccessible")
	delete(f, "DisasterRecoverGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLaunchConfigurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLaunchConfigurationResponseParams struct {
	// This parameter is returned when a launch configuration is created through this API, indicating the launch configuration ID.
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitnil,omitempty" name:"LaunchConfigurationId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLaunchConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *CreateLaunchConfigurationResponseParams `json:"Response"`
}

func (r *CreateLaunchConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLaunchConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLifecycleHookRequestParams struct {
	// Auto scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// Lifecycle hook name, which can contain Chinese characters, letters, numbers, underscores (_), hyphens (-), and periods (.) with a maximum length of 128 bytes.
	LifecycleHookName *string `json:"LifecycleHookName,omitnil,omitempty" name:"LifecycleHookName"`

	// Scenario for the lifecycle hook. Valid values: "INSTANCE_LAUNCHING" and "INSTANCE_TERMINATING"
	LifecycleTransition *string `json:"LifecycleTransition,omitnil,omitempty" name:"LifecycleTransition"`

	// Defined actions when lifecycle hook times out. Valid values: "CONTINUE" and "ABANDON". Default value: "CONTINUE"
	DefaultResult *string `json:"DefaultResult,omitnil,omitempty" name:"DefaultResult"`

	// The maximum length of time (in seconds) that can elapse before the lifecycle hook times out. Value range: 30-7200. Default value: 300
	HeartbeatTimeout *int64 `json:"HeartbeatTimeout,omitnil,omitempty" name:"HeartbeatTimeout"`

	// Additional information of a notification that Auto Scaling sends to targets. This parameter is set when you configure a notification (default value: ""). Up to 1024 characters are allowed.
	NotificationMetadata *string `json:"NotificationMetadata,omitnil,omitempty" name:"NotificationMetadata"`

	// Notification target. `NotificationTarget` and `LifecycleCommand` cannot be specified at the same time.
	NotificationTarget *NotificationTarget `json:"NotificationTarget,omitnil,omitempty" name:"NotificationTarget"`

	// The scenario where the lifecycle hook is applied. `EXTENSION`: the lifecycle hook will be triggered when AttachInstances, DetachInstances or RemoveInstaces is called. `NORMAL`: the lifecycle hook is not triggered by the above APIs. 
	LifecycleTransitionType *string `json:"LifecycleTransitionType,omitnil,omitempty" name:"LifecycleTransitionType"`

	// Remote command execution object. `NotificationTarget` and `LifecycleCommand` cannot be specified at the same time.
	LifecycleCommand *LifecycleCommand `json:"LifecycleCommand,omitnil,omitempty" name:"LifecycleCommand"`
}

type CreateLifecycleHookRequest struct {
	*tchttp.BaseRequest
	
	// Auto scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// Lifecycle hook name, which can contain Chinese characters, letters, numbers, underscores (_), hyphens (-), and periods (.) with a maximum length of 128 bytes.
	LifecycleHookName *string `json:"LifecycleHookName,omitnil,omitempty" name:"LifecycleHookName"`

	// Scenario for the lifecycle hook. Valid values: "INSTANCE_LAUNCHING" and "INSTANCE_TERMINATING"
	LifecycleTransition *string `json:"LifecycleTransition,omitnil,omitempty" name:"LifecycleTransition"`

	// Defined actions when lifecycle hook times out. Valid values: "CONTINUE" and "ABANDON". Default value: "CONTINUE"
	DefaultResult *string `json:"DefaultResult,omitnil,omitempty" name:"DefaultResult"`

	// The maximum length of time (in seconds) that can elapse before the lifecycle hook times out. Value range: 30-7200. Default value: 300
	HeartbeatTimeout *int64 `json:"HeartbeatTimeout,omitnil,omitempty" name:"HeartbeatTimeout"`

	// Additional information of a notification that Auto Scaling sends to targets. This parameter is set when you configure a notification (default value: ""). Up to 1024 characters are allowed.
	NotificationMetadata *string `json:"NotificationMetadata,omitnil,omitempty" name:"NotificationMetadata"`

	// Notification target. `NotificationTarget` and `LifecycleCommand` cannot be specified at the same time.
	NotificationTarget *NotificationTarget `json:"NotificationTarget,omitnil,omitempty" name:"NotificationTarget"`

	// The scenario where the lifecycle hook is applied. `EXTENSION`: the lifecycle hook will be triggered when AttachInstances, DetachInstances or RemoveInstaces is called. `NORMAL`: the lifecycle hook is not triggered by the above APIs. 
	LifecycleTransitionType *string `json:"LifecycleTransitionType,omitnil,omitempty" name:"LifecycleTransitionType"`

	// Remote command execution object. `NotificationTarget` and `LifecycleCommand` cannot be specified at the same time.
	LifecycleCommand *LifecycleCommand `json:"LifecycleCommand,omitnil,omitempty" name:"LifecycleCommand"`
}

func (r *CreateLifecycleHookRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLifecycleHookRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "LifecycleHookName")
	delete(f, "LifecycleTransition")
	delete(f, "DefaultResult")
	delete(f, "HeartbeatTimeout")
	delete(f, "NotificationMetadata")
	delete(f, "NotificationTarget")
	delete(f, "LifecycleTransitionType")
	delete(f, "LifecycleCommand")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLifecycleHookRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLifecycleHookResponseParams struct {
	// Lifecycle hook ID
	LifecycleHookId *string `json:"LifecycleHookId,omitnil,omitempty" name:"LifecycleHookId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLifecycleHookResponse struct {
	*tchttp.BaseResponse
	Response *CreateLifecycleHookResponseParams `json:"Response"`
}

func (r *CreateLifecycleHookResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLifecycleHookResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNotificationConfigurationRequestParams struct {
	// Auto scaling group ID.
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// Notification type, i.e., the set of types of notifications to be subscribed to. Value range:
	// <li>SCALE_OUT_SUCCESSFUL: scale-out succeeded</li>
	// <li>SCALE_OUT_FAILED: scale-out failed</li>
	// <li>SCALE_IN_SUCCESSFUL: scale-in succeeded</li>
	// <li>SCALE_IN_FAILED: scale-in failed</li>
	// <li>REPLACE_UNHEALTHY_INSTANCE_SUCCESSFUL: unhealthy instance replacement succeeded</li>
	// <li>REPLACE_UNHEALTHY_INSTANCE_FAILED: unhealthy instance replacement failed</li>
	NotificationTypes []*string `json:"NotificationTypes,omitnil,omitempty" name:"NotificationTypes"`

	// Notification group ID, which is the set of user group IDs. You can query the user group IDs through the [ListGroups](https://intl.cloud.tencent.com/document/product/598/34589?from_cn_redirect=1) API.
	NotificationUserGroupIds []*string `json:"NotificationUserGroupIds,omitnil,omitempty" name:"NotificationUserGroupIds"`

	// Notification receiver type. Valid values:
	// <br><li>USER_GROUP:User group
	// <br><li>CMQ_QUEUE:CMQ queue
	// <br><li>CMQ_TOPIC:CMQ topic
	// <br><li>TDMQ_CMQ_TOPIC:TDMQ CMQ topic
	// <br><li>TDMQ_CMQ_QUEUE:TDMQ CMQ queue
	// 
	// Default value: `USER_GROUP`.
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// CMQ queue name. This parameter is required when `TargetType` is `CMQ_QUEUE` or `TDMQ_CMQ_QUEUE`.
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// CMQ topic name. This parameter is required when `TargetType` is `CMQ_TOPIC` or `TDMQ_CMQ_TOPIC`.
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

type CreateNotificationConfigurationRequest struct {
	*tchttp.BaseRequest
	
	// Auto scaling group ID.
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// Notification type, i.e., the set of types of notifications to be subscribed to. Value range:
	// <li>SCALE_OUT_SUCCESSFUL: scale-out succeeded</li>
	// <li>SCALE_OUT_FAILED: scale-out failed</li>
	// <li>SCALE_IN_SUCCESSFUL: scale-in succeeded</li>
	// <li>SCALE_IN_FAILED: scale-in failed</li>
	// <li>REPLACE_UNHEALTHY_INSTANCE_SUCCESSFUL: unhealthy instance replacement succeeded</li>
	// <li>REPLACE_UNHEALTHY_INSTANCE_FAILED: unhealthy instance replacement failed</li>
	NotificationTypes []*string `json:"NotificationTypes,omitnil,omitempty" name:"NotificationTypes"`

	// Notification group ID, which is the set of user group IDs. You can query the user group IDs through the [ListGroups](https://intl.cloud.tencent.com/document/product/598/34589?from_cn_redirect=1) API.
	NotificationUserGroupIds []*string `json:"NotificationUserGroupIds,omitnil,omitempty" name:"NotificationUserGroupIds"`

	// Notification receiver type. Valid values:
	// <br><li>USER_GROUP:User group
	// <br><li>CMQ_QUEUE:CMQ queue
	// <br><li>CMQ_TOPIC:CMQ topic
	// <br><li>TDMQ_CMQ_TOPIC:TDMQ CMQ topic
	// <br><li>TDMQ_CMQ_QUEUE:TDMQ CMQ queue
	// 
	// Default value: `USER_GROUP`.
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// CMQ queue name. This parameter is required when `TargetType` is `CMQ_QUEUE` or `TDMQ_CMQ_QUEUE`.
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// CMQ topic name. This parameter is required when `TargetType` is `CMQ_TOPIC` or `TDMQ_CMQ_TOPIC`.
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

func (r *CreateNotificationConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNotificationConfigurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "NotificationTypes")
	delete(f, "NotificationUserGroupIds")
	delete(f, "TargetType")
	delete(f, "QueueName")
	delete(f, "TopicName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNotificationConfigurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNotificationConfigurationResponseParams struct {
	// Notification ID.
	AutoScalingNotificationId *string `json:"AutoScalingNotificationId,omitnil,omitempty" name:"AutoScalingNotificationId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateNotificationConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *CreateNotificationConfigurationResponseParams `json:"Response"`
}

func (r *CreateNotificationConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNotificationConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateScalingPolicyRequestParams struct {
	// Auto scaling group ID.
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// Alarm trigger policy name.
	ScalingPolicyName *string `json:"ScalingPolicyName,omitnil,omitempty" name:"ScalingPolicyName"`

	// Scaling policy type. Valid values: <br><li>`SIMPLE` (default): A simple policy</li><li>`TARGET_TRACKING`: A target tracking policy</li>.
	ScalingPolicyType *string `json:"ScalingPolicyType,omitnil,omitempty" name:"ScalingPolicyType"`

	// The method to adjust the desired capacity after the alarm is triggered. It’s only available when `ScalingPolicyType` is `Simple`. Valid values: <br><li>`CHANGE_IN_CAPACITY`: Increase or decrease the desired capacity </li><li>`EXACT_CAPACITY`: Adjust to the specified desired capacity </li> <li>`PERCENT_CHANGE_IN_CAPACITY`: Adjust the desired capacity by percentage </li>
	AdjustmentType *string `json:"AdjustmentType,omitnil,omitempty" name:"AdjustmentType"`

	// Specifies how to adjust the number of desired capacity when the alarm is triggered. It’s only available when `ScalingPolicyType` is `Simple`. Values: <br><li>`AdjustmentType`=`CHANGE_IN_CAPACITY`: Number of instances to add (positive number) or remove (negative number). </li> <li>`AdjustmentType`=`EXACT_CAPACITY`: Set the desired capacity to the specified number. It must be ≥ 0. </li> <li>`AdjustmentType`=`PERCENT_CHANGE_IN_CAPACITY`: Percentage of instance number. Add instances (positive value) or remove instances (negative value) accordingly.
	AdjustmentValue *int64 `json:"AdjustmentValue,omitnil,omitempty" name:"AdjustmentValue"`

	// Cooldown period (in seconds). This parameter is only applicable to a simple policy. Default value: 300.
	Cooldown *uint64 `json:"Cooldown,omitnil,omitempty" name:"Cooldown"`

	// Alarm monitoring metric. It’s only available when `ScalingPolicyType` is `Simple`.
	MetricAlarm *MetricAlarm `json:"MetricAlarm,omitnil,omitempty" name:"MetricAlarm"`

	// Preset monitoring item. It’s only available when `ScalingPolicyType` is `TARGET_TRACKING`. Valid values: <br><li>ASG_AVG_CPU_UTILIZATION: Average CPU utilization</li><li>ASG_AVG_LAN_TRAFFIC_OUT: Average private bandwidth out</li><li>ASG_AVG_LAN_TRAFFIC_IN: Average private bandwidth in</li><li>ASG_AVG_WAN_TRAFFIC_OUT: Average public bandwidth out</li><li>ASG_AVG_WAN_TRAFFIC_IN: Average public bandwidth in</li>
	PredefinedMetricType *string `json:"PredefinedMetricType,omitnil,omitempty" name:"PredefinedMetricType"`

	// Target value. It’s only available when `ScalingPolicyType` is `TARGET_TRACKING`. Value ranges: <br><li>`ASG_AVG_CPU_UTILIZATION` (in %): [1, 100)</li><li>`ASG_AVG_LAN_TRAFFIC_OUT` (in Mbps): >0</li><li>`ASG_AVG_LAN_TRAFFIC_IN` (in Mbps): >0</li><li>`ASG_AVG_WAN_TRAFFIC_OUT` (in Mbps): >0</li><li>`ASG_AVG_WAN_TRAFFIC_IN` (in Mbps): >0</li>
	TargetValue *uint64 `json:"TargetValue,omitnil,omitempty" name:"TargetValue"`

	// Instance warm-up period (in seconds). It’s only available when `ScalingPolicyType` is `TARGET_TRACKING`. Value range: 0-3600. Default value: 300.
	EstimatedInstanceWarmup *uint64 `json:"EstimatedInstanceWarmup,omitnil,omitempty" name:"EstimatedInstanceWarmup"`

	// Whether to disable scale-in. It’s only available when `ScalingPolicyType` is `TARGET_TRACKING`. Valid values: <br><li>`true`: Do not scale in </li><li>`false` (default): Both scale-out and scale-in can be triggered.</li>
	DisableScaleIn *bool `json:"DisableScaleIn,omitnil,omitempty" name:"DisableScaleIn"`

	// This parameter is diused. Please use [CreateNotificationConfiguration](https://intl.cloud.tencent.com/document/api/377/33185?from_cn_redirect=1) instead.
	// Notification group ID, which is the set of user group IDs.
	NotificationUserGroupIds []*string `json:"NotificationUserGroupIds,omitnil,omitempty" name:"NotificationUserGroupIds"`
}

type CreateScalingPolicyRequest struct {
	*tchttp.BaseRequest
	
	// Auto scaling group ID.
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// Alarm trigger policy name.
	ScalingPolicyName *string `json:"ScalingPolicyName,omitnil,omitempty" name:"ScalingPolicyName"`

	// Scaling policy type. Valid values: <br><li>`SIMPLE` (default): A simple policy</li><li>`TARGET_TRACKING`: A target tracking policy</li>.
	ScalingPolicyType *string `json:"ScalingPolicyType,omitnil,omitempty" name:"ScalingPolicyType"`

	// The method to adjust the desired capacity after the alarm is triggered. It’s only available when `ScalingPolicyType` is `Simple`. Valid values: <br><li>`CHANGE_IN_CAPACITY`: Increase or decrease the desired capacity </li><li>`EXACT_CAPACITY`: Adjust to the specified desired capacity </li> <li>`PERCENT_CHANGE_IN_CAPACITY`: Adjust the desired capacity by percentage </li>
	AdjustmentType *string `json:"AdjustmentType,omitnil,omitempty" name:"AdjustmentType"`

	// Specifies how to adjust the number of desired capacity when the alarm is triggered. It’s only available when `ScalingPolicyType` is `Simple`. Values: <br><li>`AdjustmentType`=`CHANGE_IN_CAPACITY`: Number of instances to add (positive number) or remove (negative number). </li> <li>`AdjustmentType`=`EXACT_CAPACITY`: Set the desired capacity to the specified number. It must be ≥ 0. </li> <li>`AdjustmentType`=`PERCENT_CHANGE_IN_CAPACITY`: Percentage of instance number. Add instances (positive value) or remove instances (negative value) accordingly.
	AdjustmentValue *int64 `json:"AdjustmentValue,omitnil,omitempty" name:"AdjustmentValue"`

	// Cooldown period (in seconds). This parameter is only applicable to a simple policy. Default value: 300.
	Cooldown *uint64 `json:"Cooldown,omitnil,omitempty" name:"Cooldown"`

	// Alarm monitoring metric. It’s only available when `ScalingPolicyType` is `Simple`.
	MetricAlarm *MetricAlarm `json:"MetricAlarm,omitnil,omitempty" name:"MetricAlarm"`

	// Preset monitoring item. It’s only available when `ScalingPolicyType` is `TARGET_TRACKING`. Valid values: <br><li>ASG_AVG_CPU_UTILIZATION: Average CPU utilization</li><li>ASG_AVG_LAN_TRAFFIC_OUT: Average private bandwidth out</li><li>ASG_AVG_LAN_TRAFFIC_IN: Average private bandwidth in</li><li>ASG_AVG_WAN_TRAFFIC_OUT: Average public bandwidth out</li><li>ASG_AVG_WAN_TRAFFIC_IN: Average public bandwidth in</li>
	PredefinedMetricType *string `json:"PredefinedMetricType,omitnil,omitempty" name:"PredefinedMetricType"`

	// Target value. It’s only available when `ScalingPolicyType` is `TARGET_TRACKING`. Value ranges: <br><li>`ASG_AVG_CPU_UTILIZATION` (in %): [1, 100)</li><li>`ASG_AVG_LAN_TRAFFIC_OUT` (in Mbps): >0</li><li>`ASG_AVG_LAN_TRAFFIC_IN` (in Mbps): >0</li><li>`ASG_AVG_WAN_TRAFFIC_OUT` (in Mbps): >0</li><li>`ASG_AVG_WAN_TRAFFIC_IN` (in Mbps): >0</li>
	TargetValue *uint64 `json:"TargetValue,omitnil,omitempty" name:"TargetValue"`

	// Instance warm-up period (in seconds). It’s only available when `ScalingPolicyType` is `TARGET_TRACKING`. Value range: 0-3600. Default value: 300.
	EstimatedInstanceWarmup *uint64 `json:"EstimatedInstanceWarmup,omitnil,omitempty" name:"EstimatedInstanceWarmup"`

	// Whether to disable scale-in. It’s only available when `ScalingPolicyType` is `TARGET_TRACKING`. Valid values: <br><li>`true`: Do not scale in </li><li>`false` (default): Both scale-out and scale-in can be triggered.</li>
	DisableScaleIn *bool `json:"DisableScaleIn,omitnil,omitempty" name:"DisableScaleIn"`

	// This parameter is diused. Please use [CreateNotificationConfiguration](https://intl.cloud.tencent.com/document/api/377/33185?from_cn_redirect=1) instead.
	// Notification group ID, which is the set of user group IDs.
	NotificationUserGroupIds []*string `json:"NotificationUserGroupIds,omitnil,omitempty" name:"NotificationUserGroupIds"`
}

func (r *CreateScalingPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateScalingPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "ScalingPolicyName")
	delete(f, "ScalingPolicyType")
	delete(f, "AdjustmentType")
	delete(f, "AdjustmentValue")
	delete(f, "Cooldown")
	delete(f, "MetricAlarm")
	delete(f, "PredefinedMetricType")
	delete(f, "TargetValue")
	delete(f, "EstimatedInstanceWarmup")
	delete(f, "DisableScaleIn")
	delete(f, "NotificationUserGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateScalingPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateScalingPolicyResponseParams struct {
	// Alarm trigger policy ID.
	AutoScalingPolicyId *string `json:"AutoScalingPolicyId,omitnil,omitempty" name:"AutoScalingPolicyId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateScalingPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateScalingPolicyResponseParams `json:"Response"`
}

func (r *CreateScalingPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateScalingPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateScheduledActionRequestParams struct {
	// Auto scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// Scheduled task name, which can only contain letters, numbers, underscores, hyphens ("-"), and decimal points with a maximum length of 60 bytes and must be unique in an auto scaling group.
	ScheduledActionName *string `json:"ScheduledActionName,omitnil,omitempty" name:"ScheduledActionName"`

	// The maximum number of instances set for the auto scaling group when the scheduled task is triggered.
	MaxSize *uint64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`

	// The minimum number of instances set for the auto scaling group when the scheduled task is triggered.
	MinSize *uint64 `json:"MinSize,omitnil,omitempty" name:"MinSize"`

	// The desired number of instances set for the auto scaling group when the scheduled task is triggered.
	DesiredCapacity *uint64 `json:"DesiredCapacity,omitnil,omitempty" name:"DesiredCapacity"`

	// Initial triggered time of the scheduled task. The value is in `Beijing time` (UTC+8) in the format of `YYYY-MM-DDThh:mm:ss+08:00` according to the `ISO8601` standard.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time of the scheduled task. The value is in `Beijing time` (UTC+8) in the format of `YYYY-MM-DDThh:mm:ss+08:00` according to the `ISO8601` standard. <br><br>This parameter and `Recurrence` need to be specified at the same time. After the end time, the scheduled task will no longer take effect.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Repeating mode of the scheduled task, which is in standard cron format. <br><br>This parameter and `EndTime` need to be specified at the same time.
	Recurrence *string `json:"Recurrence,omitnil,omitempty" name:"Recurrence"`
}

type CreateScheduledActionRequest struct {
	*tchttp.BaseRequest
	
	// Auto scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// Scheduled task name, which can only contain letters, numbers, underscores, hyphens ("-"), and decimal points with a maximum length of 60 bytes and must be unique in an auto scaling group.
	ScheduledActionName *string `json:"ScheduledActionName,omitnil,omitempty" name:"ScheduledActionName"`

	// The maximum number of instances set for the auto scaling group when the scheduled task is triggered.
	MaxSize *uint64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`

	// The minimum number of instances set for the auto scaling group when the scheduled task is triggered.
	MinSize *uint64 `json:"MinSize,omitnil,omitempty" name:"MinSize"`

	// The desired number of instances set for the auto scaling group when the scheduled task is triggered.
	DesiredCapacity *uint64 `json:"DesiredCapacity,omitnil,omitempty" name:"DesiredCapacity"`

	// Initial triggered time of the scheduled task. The value is in `Beijing time` (UTC+8) in the format of `YYYY-MM-DDThh:mm:ss+08:00` according to the `ISO8601` standard.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time of the scheduled task. The value is in `Beijing time` (UTC+8) in the format of `YYYY-MM-DDThh:mm:ss+08:00` according to the `ISO8601` standard. <br><br>This parameter and `Recurrence` need to be specified at the same time. After the end time, the scheduled task will no longer take effect.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Repeating mode of the scheduled task, which is in standard cron format. <br><br>This parameter and `EndTime` need to be specified at the same time.
	Recurrence *string `json:"Recurrence,omitnil,omitempty" name:"Recurrence"`
}

func (r *CreateScheduledActionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateScheduledActionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "ScheduledActionName")
	delete(f, "MaxSize")
	delete(f, "MinSize")
	delete(f, "DesiredCapacity")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Recurrence")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateScheduledActionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateScheduledActionResponseParams struct {
	// Scheduled task ID
	ScheduledActionId *string `json:"ScheduledActionId,omitnil,omitempty" name:"ScheduledActionId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateScheduledActionResponse struct {
	*tchttp.BaseResponse
	Response *CreateScheduledActionResponseParams `json:"Response"`
}

func (r *CreateScheduledActionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateScheduledActionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataDisk struct {
	// Data disk type. See [Cloud Disk Types](https://intl.cloud.tencent.com/document/product/362/31636). Valid values:<br><li>`LOCAL_BASIC`: Local disk<br><li>`LOCAL_SSD`: Local SSD disk<br><li>`CLOUD_BASIC`: HDD cloud disk<br><li>`CLOUD_PREMIUM`: Premium cloud storage<br><li>`CLOUD_SSD`: SSD cloud disk<br><li>`CLOUD_HSSD`: Enhanced SSD<br><li>`CLOUD_TSSD`: Tremendous SSD<br><br>The default value should be the same as the `DiskType` field under `SystemDisk`.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// Data disk size (in GB). The minimum adjustment increment is 10 GB. The value range varies by data disk type. For more information on limits, see [CVM Instance Configuration](https://intl.cloud.tencent.com/document/product/213/2177?from_cn_redirect=1). The default value is 0, indicating that no data disk is purchased. For more information, see the product documentation.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// Data disk snapshot ID, such as `snap-l8psqwnt`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// Specifies whether the data disk is terminated along with the termination of the associated CVM instance.  Values: <br><li>`TRUE` (only available for pay-as-you-go cloud disks that are billed by hour) and `FALSE`.
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitnil,omitempty" name:"DeleteWithInstance"`

	// Data disk encryption. Valid values: <br><li>`TRUE`: Encrypted<br><li>`FALSE`: Not encrypted
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Encrypt *bool `json:"Encrypt,omitnil,omitempty" name:"Encrypt"`

	// Cloud disk performance (MB/s). This parameter is used to purchase extra performance for the cloud disk. For details on the feature and limits, see [Enhanced SSD Performance](https://intl.cloud.tencent.com/document/product/362/51896?from_cn_redirect=1#. E5.A2.9E.E5.BC.BA.E5.9E.8B-ssd-.E4.BA.91.E7.A1.AC.E7.9B.98.E9.A2.9D.E5.A4.96 .E6.80.A7.E8.83.BD).
	// This feature is only available to enhanced SSD (`CLOUD_HSSD`) and tremendous SSD (`CLOUD_TSSD`) disks with a capacity greater than 460 GB.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ThroughputPerformance *uint64 `json:"ThroughputPerformance,omitnil,omitempty" name:"ThroughputPerformance"`
}

// Predefined struct for user
type DeleteAutoScalingGroupRequestParams struct {
	// Auto scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`
}

type DeleteAutoScalingGroupRequest struct {
	*tchttp.BaseRequest
	
	// Auto scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`
}

func (r *DeleteAutoScalingGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAutoScalingGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAutoScalingGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAutoScalingGroupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAutoScalingGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAutoScalingGroupResponseParams `json:"Response"`
}

func (r *DeleteAutoScalingGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAutoScalingGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLaunchConfigurationRequestParams struct {
	// ID of the launch configuration to be deleted.
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitnil,omitempty" name:"LaunchConfigurationId"`
}

type DeleteLaunchConfigurationRequest struct {
	*tchttp.BaseRequest
	
	// ID of the launch configuration to be deleted.
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitnil,omitempty" name:"LaunchConfigurationId"`
}

func (r *DeleteLaunchConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLaunchConfigurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LaunchConfigurationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLaunchConfigurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLaunchConfigurationResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLaunchConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLaunchConfigurationResponseParams `json:"Response"`
}

func (r *DeleteLaunchConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLaunchConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLifecycleHookRequestParams struct {
	// Lifecycle hook ID
	LifecycleHookId *string `json:"LifecycleHookId,omitnil,omitempty" name:"LifecycleHookId"`
}

type DeleteLifecycleHookRequest struct {
	*tchttp.BaseRequest
	
	// Lifecycle hook ID
	LifecycleHookId *string `json:"LifecycleHookId,omitnil,omitempty" name:"LifecycleHookId"`
}

func (r *DeleteLifecycleHookRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLifecycleHookRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LifecycleHookId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLifecycleHookRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLifecycleHookResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLifecycleHookResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLifecycleHookResponseParams `json:"Response"`
}

func (r *DeleteLifecycleHookResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLifecycleHookResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNotificationConfigurationRequestParams struct {
	// ID of the notification to be deleted.
	AutoScalingNotificationId *string `json:"AutoScalingNotificationId,omitnil,omitempty" name:"AutoScalingNotificationId"`
}

type DeleteNotificationConfigurationRequest struct {
	*tchttp.BaseRequest
	
	// ID of the notification to be deleted.
	AutoScalingNotificationId *string `json:"AutoScalingNotificationId,omitnil,omitempty" name:"AutoScalingNotificationId"`
}

func (r *DeleteNotificationConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNotificationConfigurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingNotificationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNotificationConfigurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNotificationConfigurationResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteNotificationConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNotificationConfigurationResponseParams `json:"Response"`
}

func (r *DeleteNotificationConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNotificationConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteScalingPolicyRequestParams struct {
	// ID of the alarm policy to be deleted.
	AutoScalingPolicyId *string `json:"AutoScalingPolicyId,omitnil,omitempty" name:"AutoScalingPolicyId"`
}

type DeleteScalingPolicyRequest struct {
	*tchttp.BaseRequest
	
	// ID of the alarm policy to be deleted.
	AutoScalingPolicyId *string `json:"AutoScalingPolicyId,omitnil,omitempty" name:"AutoScalingPolicyId"`
}

func (r *DeleteScalingPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteScalingPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingPolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteScalingPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteScalingPolicyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteScalingPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteScalingPolicyResponseParams `json:"Response"`
}

func (r *DeleteScalingPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteScalingPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteScheduledActionRequestParams struct {
	// ID of the scheduled task to be deleted.
	ScheduledActionId *string `json:"ScheduledActionId,omitnil,omitempty" name:"ScheduledActionId"`
}

type DeleteScheduledActionRequest struct {
	*tchttp.BaseRequest
	
	// ID of the scheduled task to be deleted.
	ScheduledActionId *string `json:"ScheduledActionId,omitnil,omitempty" name:"ScheduledActionId"`
}

func (r *DeleteScheduledActionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteScheduledActionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScheduledActionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteScheduledActionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteScheduledActionResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteScheduledActionResponse struct {
	*tchttp.BaseResponse
	Response *DeleteScheduledActionResponseParams `json:"Response"`
}

func (r *DeleteScheduledActionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteScheduledActionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountLimitsRequestParams struct {

}

type DescribeAccountLimitsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAccountLimitsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountLimitsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountLimitsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountLimitsResponseParams struct {
	// Maximum number of launch configurations allowed for creation by the user account
	MaxNumberOfLaunchConfigurations *int64 `json:"MaxNumberOfLaunchConfigurations,omitnil,omitempty" name:"MaxNumberOfLaunchConfigurations"`

	// Current number of launch configurations under the user account
	NumberOfLaunchConfigurations *int64 `json:"NumberOfLaunchConfigurations,omitnil,omitempty" name:"NumberOfLaunchConfigurations"`

	// Maximum number of auto scaling groups allowed for creation by the user account
	MaxNumberOfAutoScalingGroups *int64 `json:"MaxNumberOfAutoScalingGroups,omitnil,omitempty" name:"MaxNumberOfAutoScalingGroups"`

	// Current number of auto scaling groups under the user account
	NumberOfAutoScalingGroups *int64 `json:"NumberOfAutoScalingGroups,omitnil,omitempty" name:"NumberOfAutoScalingGroups"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccountLimitsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccountLimitsResponseParams `json:"Response"`
}

func (r *DescribeAccountLimitsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountLimitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoScalingActivitiesRequestParams struct {
	// Queries by one or more scaling activity IDs in the format of `asa-5l2ejpfo`. The maximum quantity per request is 100. This parameter does not support specifying both `ActivityIds` and `Filters` at the same time.
	ActivityIds []*string `json:"ActivityIds,omitnil,omitempty" name:"ActivityIds"`

	// Filter.
	// <li> auto-scaling-group-id - String - Required: No - (Filter) Filter by auto scaling group ID.</li>
	// <li> activity-status-code - String - Required: No - (Filter) Filter by scaling activity status . (INIT: initializing | RUNNING: running | SUCCESSFUL: succeeded | PARTIALLY_SUCCESSFUL: partially succeeded | FAILED: failed | CANCELLED: canceled)</li>
	// <li> activity-type - String - Required: No - (Filter) Filter by scaling activity type. (SCALE_OUT: scale-out | SCALE_IN: scale-in | ATTACH_INSTANCES: adding an instance | REMOVE_INSTANCES: terminating an instance | DETACH_INSTANCES: removing an instance | TERMINATE_INSTANCES_UNEXPECTEDLY: terminating an instance in the CVM console | REPLACE_UNHEALTHY_INSTANCE: replacing an unhealthy instance | UPDATE_LOAD_BALANCERS: updating a load balancer)</li>
	// <li> activity-id - String - Required: No - (Filter) Filter by scaling activity ID.</li>
	// The maximum number of `Filters` per request is 10. The upper limit for `Filter.Values` is 5. This parameter does not support specifying both `ActivityIds` and `Filters` at the same time.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, see the relevant section in the API [overview](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Default value: 0. For more information on `Offset`, see the relevant section in the API [overview](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The earliest start time of the scaling activity, which will be ignored if ActivityIds is specified. The value is in `UTC time` in the format of `YYYY-MM-DDThh:mm:ssZ` according to the `ISO8601` standard.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// The latest end time of the scaling activity, which will be ignored if ActivityIds is specified. The value is in `UTC time` in the format of `YYYY-MM-DDThh:mm:ssZ` according to the `ISO8601` standard.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeAutoScalingActivitiesRequest struct {
	*tchttp.BaseRequest
	
	// Queries by one or more scaling activity IDs in the format of `asa-5l2ejpfo`. The maximum quantity per request is 100. This parameter does not support specifying both `ActivityIds` and `Filters` at the same time.
	ActivityIds []*string `json:"ActivityIds,omitnil,omitempty" name:"ActivityIds"`

	// Filter.
	// <li> auto-scaling-group-id - String - Required: No - (Filter) Filter by auto scaling group ID.</li>
	// <li> activity-status-code - String - Required: No - (Filter) Filter by scaling activity status . (INIT: initializing | RUNNING: running | SUCCESSFUL: succeeded | PARTIALLY_SUCCESSFUL: partially succeeded | FAILED: failed | CANCELLED: canceled)</li>
	// <li> activity-type - String - Required: No - (Filter) Filter by scaling activity type. (SCALE_OUT: scale-out | SCALE_IN: scale-in | ATTACH_INSTANCES: adding an instance | REMOVE_INSTANCES: terminating an instance | DETACH_INSTANCES: removing an instance | TERMINATE_INSTANCES_UNEXPECTEDLY: terminating an instance in the CVM console | REPLACE_UNHEALTHY_INSTANCE: replacing an unhealthy instance | UPDATE_LOAD_BALANCERS: updating a load balancer)</li>
	// <li> activity-id - String - Required: No - (Filter) Filter by scaling activity ID.</li>
	// The maximum number of `Filters` per request is 10. The upper limit for `Filter.Values` is 5. This parameter does not support specifying both `ActivityIds` and `Filters` at the same time.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, see the relevant section in the API [overview](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Default value: 0. For more information on `Offset`, see the relevant section in the API [overview](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The earliest start time of the scaling activity, which will be ignored if ActivityIds is specified. The value is in `UTC time` in the format of `YYYY-MM-DDThh:mm:ssZ` according to the `ISO8601` standard.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// The latest end time of the scaling activity, which will be ignored if ActivityIds is specified. The value is in `UTC time` in the format of `YYYY-MM-DDThh:mm:ssZ` according to the `ISO8601` standard.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *DescribeAutoScalingActivitiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoScalingActivitiesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ActivityIds")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAutoScalingActivitiesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoScalingActivitiesResponseParams struct {
	// Number of eligible scaling activities.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Information set of eligible scaling activities.
	ActivitySet []*Activity `json:"ActivitySet,omitnil,omitempty" name:"ActivitySet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAutoScalingActivitiesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAutoScalingActivitiesResponseParams `json:"Response"`
}

func (r *DescribeAutoScalingActivitiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoScalingActivitiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoScalingAdvicesRequestParams struct {
	// List of scaling groups to be queried. Upper limit: 100.
	AutoScalingGroupIds []*string `json:"AutoScalingGroupIds,omitnil,omitempty" name:"AutoScalingGroupIds"`
}

type DescribeAutoScalingAdvicesRequest struct {
	*tchttp.BaseRequest
	
	// List of scaling groups to be queried. Upper limit: 100.
	AutoScalingGroupIds []*string `json:"AutoScalingGroupIds,omitnil,omitempty" name:"AutoScalingGroupIds"`
}

func (r *DescribeAutoScalingAdvicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoScalingAdvicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAutoScalingAdvicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoScalingAdvicesResponseParams struct {
	// A collection of suggestions for scaling group configurations.
	AutoScalingAdviceSet []*AutoScalingAdvice `json:"AutoScalingAdviceSet,omitnil,omitempty" name:"AutoScalingAdviceSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAutoScalingAdvicesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAutoScalingAdvicesResponseParams `json:"Response"`
}

func (r *DescribeAutoScalingAdvicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoScalingAdvicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoScalingGroupLastActivitiesRequestParams struct {
	// ID list of an auto scaling group.
	AutoScalingGroupIds []*string `json:"AutoScalingGroupIds,omitnil,omitempty" name:"AutoScalingGroupIds"`
}

type DescribeAutoScalingGroupLastActivitiesRequest struct {
	*tchttp.BaseRequest
	
	// ID list of an auto scaling group.
	AutoScalingGroupIds []*string `json:"AutoScalingGroupIds,omitnil,omitempty" name:"AutoScalingGroupIds"`
}

func (r *DescribeAutoScalingGroupLastActivitiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoScalingGroupLastActivitiesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAutoScalingGroupLastActivitiesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoScalingGroupLastActivitiesResponseParams struct {
	// Information set of eligible scaling activities. Scaling groups without scaling activities are not returned. For example, if there are 50 auto scaling group IDs but only 45 records are returned, it indicates that 5 of the auto scaling groups do not have scaling activities.
	ActivitySet []*Activity `json:"ActivitySet,omitnil,omitempty" name:"ActivitySet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAutoScalingGroupLastActivitiesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAutoScalingGroupLastActivitiesResponseParams `json:"Response"`
}

func (r *DescribeAutoScalingGroupLastActivitiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoScalingGroupLastActivitiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoScalingGroupsRequestParams struct {
	// Queries by one or more auto scaling group IDs in the format of `asg-nkdwoui0`. The maximum quantity per request is 100. This parameter does not support specifying both `AutoScalingGroupIds` and `Filters` at the same time.
	AutoScalingGroupIds []*string `json:"AutoScalingGroupIds,omitnil,omitempty" name:"AutoScalingGroupIds"`

	// Filters.
	// <li> auto-scaling-group-id - String - Required: No - (Filter) Filter by auto scaling group ID.</li>
	// <li> auto-scaling-group-name - String - Required: No - (Filter) Filter by auto scaling group name.</li>
	// <li> vague-auto-scaling-group-name - String - Required: No - (Filter) Fuzzy search by auto scaling group name.</li>
	// <li> launch-configuration-id - String - Required: No - (Filter) Filter by launch configuration ID.</li>
	// <li> tag-key - String - Required: No - (Filter) Filter by tag key.</li>
	// <li> tag-value - String - Required: No - (Filter) Filter by tag value.</li>
	// <li> tag:tag-key - String - Required: No - (Filter) Filter by tag key-value pair. The tag-key should be replaced with a specified tag key. For more information, see example 2.</li>
	// The maximum number of `Filters` in each request is 10. The upper limit for `Filter.Values` is 5. This parameter cannot specify `AutoScalingGroupIds` and `Filters` at the same time.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, see the relevant section in the API [overview](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Default value: 0. For more information on `Offset`, see the relevant section in the API [overview](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeAutoScalingGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Queries by one or more auto scaling group IDs in the format of `asg-nkdwoui0`. The maximum quantity per request is 100. This parameter does not support specifying both `AutoScalingGroupIds` and `Filters` at the same time.
	AutoScalingGroupIds []*string `json:"AutoScalingGroupIds,omitnil,omitempty" name:"AutoScalingGroupIds"`

	// Filters.
	// <li> auto-scaling-group-id - String - Required: No - (Filter) Filter by auto scaling group ID.</li>
	// <li> auto-scaling-group-name - String - Required: No - (Filter) Filter by auto scaling group name.</li>
	// <li> vague-auto-scaling-group-name - String - Required: No - (Filter) Fuzzy search by auto scaling group name.</li>
	// <li> launch-configuration-id - String - Required: No - (Filter) Filter by launch configuration ID.</li>
	// <li> tag-key - String - Required: No - (Filter) Filter by tag key.</li>
	// <li> tag-value - String - Required: No - (Filter) Filter by tag value.</li>
	// <li> tag:tag-key - String - Required: No - (Filter) Filter by tag key-value pair. The tag-key should be replaced with a specified tag key. For more information, see example 2.</li>
	// The maximum number of `Filters` in each request is 10. The upper limit for `Filter.Values` is 5. This parameter cannot specify `AutoScalingGroupIds` and `Filters` at the same time.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, see the relevant section in the API [overview](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Default value: 0. For more information on `Offset`, see the relevant section in the API [overview](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeAutoScalingGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoScalingGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupIds")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAutoScalingGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoScalingGroupsResponseParams struct {
	// List of auto scaling group details.
	AutoScalingGroupSet []*AutoScalingGroup `json:"AutoScalingGroupSet,omitnil,omitempty" name:"AutoScalingGroupSet"`

	// Number of eligible auto scaling groups.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAutoScalingGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAutoScalingGroupsResponseParams `json:"Response"`
}

func (r *DescribeAutoScalingGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoScalingGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoScalingInstancesRequestParams struct {
	// IDs of the CVM instances to query. Up to 100 IDs can be queried at one time. `InstanceIds` and `Filters` can not be both specified.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Filter.
	// <li> instance-id - String - Required: No - (Filter) Filter by instance ID.</li>
	// <li> auto-scaling-group-id - String - Required: No - (Filter) Filter by auto scaling group ID.</li>
	// The maximum number of `Filters` per request is 10. The upper limit for `Filter.Values` is 5. This parameter does not support specifying both `InstanceIds` and `Filters` at the same time.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. Default value: 0. For more information on `Offset`, see the relevant section in the API [overview](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of returned results. Default value: `20`. Maximum value: `100`. For more information on `Limit`, see the relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeAutoScalingInstancesRequest struct {
	*tchttp.BaseRequest
	
	// IDs of the CVM instances to query. Up to 100 IDs can be queried at one time. `InstanceIds` and `Filters` can not be both specified.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Filter.
	// <li> instance-id - String - Required: No - (Filter) Filter by instance ID.</li>
	// <li> auto-scaling-group-id - String - Required: No - (Filter) Filter by auto scaling group ID.</li>
	// The maximum number of `Filters` per request is 10. The upper limit for `Filter.Values` is 5. This parameter does not support specifying both `InstanceIds` and `Filters` at the same time.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. Default value: 0. For more information on `Offset`, see the relevant section in the API [overview](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of returned results. Default value: `20`. Maximum value: `100`. For more information on `Limit`, see the relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeAutoScalingInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoScalingInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAutoScalingInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoScalingInstancesResponseParams struct {
	// List of instance details.
	AutoScalingInstanceSet []*Instance `json:"AutoScalingInstanceSet,omitnil,omitempty" name:"AutoScalingInstanceSet"`

	// Number of eligible instances.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAutoScalingInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAutoScalingInstancesResponseParams `json:"Response"`
}

func (r *DescribeAutoScalingInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoScalingInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLaunchConfigurationsRequestParams struct {
	// Queries by one or more launch configuration IDs in the format of `asc-ouy1ax38`. The maximum quantity per request is 100. This parameter does not support specifying both `LaunchConfigurationIds` and `Filters` at the same time.
	LaunchConfigurationIds []*string `json:"LaunchConfigurationIds,omitnil,omitempty" name:"LaunchConfigurationIds"`

	// Filters
	// <li> `launch-configuration-id` - String - Required: No - Filter by launch configuration ID.</li>
	// <li> `launch-configuration-name` - String - Required: No - Filter by launch configuration name.</li>
	// <li> `launch-configuration-name` - String - Required: No - Fuzzy search by launch configuration name.</li>
	// <li> `tag-key` - String - Required: No - Filter by the tag key.</li>
	// <li> `tag-value` - String - Required: No - Filter by the tag value.</li>
	// <li>tag:tag-key - String - Optional - Filter by tag key pair. Use a specific tag key to replace `tag-key`. See Example 3 for the detailed usage.</li>
	// </li>
	// The maximum number of `Filters` per request is 10. The upper limit for `Filter.Values` is 5. This parameter does not support specifying both `LaunchConfigurationIds` and `Filters` at the same time.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The number of returned results. Default value: `20`. Maximum value: `100`. For more information on `Limit`, see the relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The offset. Default value: `0`. For more information on `Offset`, see the relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeLaunchConfigurationsRequest struct {
	*tchttp.BaseRequest
	
	// Queries by one or more launch configuration IDs in the format of `asc-ouy1ax38`. The maximum quantity per request is 100. This parameter does not support specifying both `LaunchConfigurationIds` and `Filters` at the same time.
	LaunchConfigurationIds []*string `json:"LaunchConfigurationIds,omitnil,omitempty" name:"LaunchConfigurationIds"`

	// Filters
	// <li> `launch-configuration-id` - String - Required: No - Filter by launch configuration ID.</li>
	// <li> `launch-configuration-name` - String - Required: No - Filter by launch configuration name.</li>
	// <li> `launch-configuration-name` - String - Required: No - Fuzzy search by launch configuration name.</li>
	// <li> `tag-key` - String - Required: No - Filter by the tag key.</li>
	// <li> `tag-value` - String - Required: No - Filter by the tag value.</li>
	// <li>tag:tag-key - String - Optional - Filter by tag key pair. Use a specific tag key to replace `tag-key`. See Example 3 for the detailed usage.</li>
	// </li>
	// The maximum number of `Filters` per request is 10. The upper limit for `Filter.Values` is 5. This parameter does not support specifying both `LaunchConfigurationIds` and `Filters` at the same time.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The number of returned results. Default value: `20`. Maximum value: `100`. For more information on `Limit`, see the relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The offset. Default value: `0`. For more information on `Offset`, see the relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeLaunchConfigurationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLaunchConfigurationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LaunchConfigurationIds")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLaunchConfigurationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLaunchConfigurationsResponseParams struct {
	// Number of eligible launch configurations.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of launch configuration details.
	LaunchConfigurationSet []*LaunchConfiguration `json:"LaunchConfigurationSet,omitnil,omitempty" name:"LaunchConfigurationSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLaunchConfigurationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLaunchConfigurationsResponseParams `json:"Response"`
}

func (r *DescribeLaunchConfigurationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLaunchConfigurationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLifecycleHooksRequestParams struct {
	// Queries by one or more lifecycle hook IDs in the format of `ash-8azjzxcl`. The maximum quantity per request is 100. This parameter does not support specifying both `LifecycleHookIds` and `Filters` at the same time.
	LifecycleHookIds []*string `json:"LifecycleHookIds,omitnil,omitempty" name:"LifecycleHookIds"`

	// Filters.
	// <li> `lifecycle-hook-id` - String - Required: No - (Filter) Filter by lifecycle hook ID.</li>
	// <li> `lifecycle-hook-name` - String - Required: No - (Filter) Filter by lifecycle hook name.</li>
	// <li> `auto-scaling-group-id` - String - Required: No - (Filter) Filter by scaling group ID.</li>
	// Up to 10 filters can be included in a request and up to 5 values for each filter. It cannot be specified with `LifecycleHookIds` at the same time.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, see the relevant section in the API [overview](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Default value: 0. For more information on `Offset`, see the relevant section in the API [overview](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeLifecycleHooksRequest struct {
	*tchttp.BaseRequest
	
	// Queries by one or more lifecycle hook IDs in the format of `ash-8azjzxcl`. The maximum quantity per request is 100. This parameter does not support specifying both `LifecycleHookIds` and `Filters` at the same time.
	LifecycleHookIds []*string `json:"LifecycleHookIds,omitnil,omitempty" name:"LifecycleHookIds"`

	// Filters.
	// <li> `lifecycle-hook-id` - String - Required: No - (Filter) Filter by lifecycle hook ID.</li>
	// <li> `lifecycle-hook-name` - String - Required: No - (Filter) Filter by lifecycle hook name.</li>
	// <li> `auto-scaling-group-id` - String - Required: No - (Filter) Filter by scaling group ID.</li>
	// Up to 10 filters can be included in a request and up to 5 values for each filter. It cannot be specified with `LifecycleHookIds` at the same time.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, see the relevant section in the API [overview](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Default value: 0. For more information on `Offset`, see the relevant section in the API [overview](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeLifecycleHooksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLifecycleHooksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LifecycleHookIds")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLifecycleHooksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLifecycleHooksResponseParams struct {
	// Array of lifecycle hooks
	LifecycleHookSet []*LifecycleHook `json:"LifecycleHookSet,omitnil,omitempty" name:"LifecycleHookSet"`

	// Total quantity
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLifecycleHooksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLifecycleHooksResponseParams `json:"Response"`
}

func (r *DescribeLifecycleHooksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLifecycleHooksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotificationConfigurationsRequestParams struct {
	// Queries by one or more notification IDs in the format of asn-2sestqbr. The maximum number of instances per request is 100. This parameter does not support specifying both `AutoScalingNotificationIds` and `Filters` at the same time.
	AutoScalingNotificationIds []*string `json:"AutoScalingNotificationIds,omitnil,omitempty" name:"AutoScalingNotificationIds"`

	// Filter.
	// <li> auto-scaling-notification-id - String - Required: No - (Filter) Filter by notification ID.</li>
	// <li> auto-scaling-group-id - String - Required: No - (Filter) Filter by auto scaling group ID.</li>
	// The maximum number of `Filters` per request is 10. The upper limit for `Filter.Values` is 5. This parameter does not support specifying both `AutoScalingNotificationIds` and `Filters` at the same time.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, see the relevant section in the API [overview](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Default value: 0. For more information on `Offset`, see the relevant section in the API [overview](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeNotificationConfigurationsRequest struct {
	*tchttp.BaseRequest
	
	// Queries by one or more notification IDs in the format of asn-2sestqbr. The maximum number of instances per request is 100. This parameter does not support specifying both `AutoScalingNotificationIds` and `Filters` at the same time.
	AutoScalingNotificationIds []*string `json:"AutoScalingNotificationIds,omitnil,omitempty" name:"AutoScalingNotificationIds"`

	// Filter.
	// <li> auto-scaling-notification-id - String - Required: No - (Filter) Filter by notification ID.</li>
	// <li> auto-scaling-group-id - String - Required: No - (Filter) Filter by auto scaling group ID.</li>
	// The maximum number of `Filters` per request is 10. The upper limit for `Filter.Values` is 5. This parameter does not support specifying both `AutoScalingNotificationIds` and `Filters` at the same time.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, see the relevant section in the API [overview](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Default value: 0. For more information on `Offset`, see the relevant section in the API [overview](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeNotificationConfigurationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotificationConfigurationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingNotificationIds")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNotificationConfigurationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotificationConfigurationsResponseParams struct {
	// Number of eligible notifications.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of AS event notification details.
	AutoScalingNotificationSet []*AutoScalingNotification `json:"AutoScalingNotificationSet,omitnil,omitempty" name:"AutoScalingNotificationSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNotificationConfigurationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNotificationConfigurationsResponseParams `json:"Response"`
}

func (r *DescribeNotificationConfigurationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotificationConfigurationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScalingPoliciesRequestParams struct {
	// Queries by one or more alarm policy IDs in the format of asp-i9vkg894. The maximum number of instances per request is 100. This parameter does not support specifying both `AutoScalingPolicyIds` and `Filters` at the same time.
	AutoScalingPolicyIds []*string `json:"AutoScalingPolicyIds,omitnil,omitempty" name:"AutoScalingPolicyIds"`

	// Filters.
	// <li> `auto-scaling-policy-id` - String - Optional - Filter by the alarm policy ID.</li>
	// <li> `auto-scaling-group-id` - String - Optional - Filter by the scaling group ID.</li>
	// <li> `scaling-policy-name` - String - Optional - Filter by the alarm policy name.</li>
	// <li> `scaling-policy-type` - String - Optional - Filter by the alarm policy type. Valid values: `SIMPLE`, `TARGET_TRACKING`.</li>
	// The maximum number of `Filters` per request is 10. The upper limit for `Filter.Values` is 5. You cannot specify `AutoScalingPolicyIds` and `Filters` at the same time.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, see the relevant section in the API [overview](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Default value: 0. For more information on `Offset`, see the relevant section in the API [overview](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeScalingPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// Queries by one or more alarm policy IDs in the format of asp-i9vkg894. The maximum number of instances per request is 100. This parameter does not support specifying both `AutoScalingPolicyIds` and `Filters` at the same time.
	AutoScalingPolicyIds []*string `json:"AutoScalingPolicyIds,omitnil,omitempty" name:"AutoScalingPolicyIds"`

	// Filters.
	// <li> `auto-scaling-policy-id` - String - Optional - Filter by the alarm policy ID.</li>
	// <li> `auto-scaling-group-id` - String - Optional - Filter by the scaling group ID.</li>
	// <li> `scaling-policy-name` - String - Optional - Filter by the alarm policy name.</li>
	// <li> `scaling-policy-type` - String - Optional - Filter by the alarm policy type. Valid values: `SIMPLE`, `TARGET_TRACKING`.</li>
	// The maximum number of `Filters` per request is 10. The upper limit for `Filter.Values` is 5. You cannot specify `AutoScalingPolicyIds` and `Filters` at the same time.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, see the relevant section in the API [overview](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Default value: 0. For more information on `Offset`, see the relevant section in the API [overview](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeScalingPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScalingPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingPolicyIds")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScalingPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScalingPoliciesResponseParams struct {
	// List of AS alarm trigger policy details.
	ScalingPolicySet []*ScalingPolicy `json:"ScalingPolicySet,omitnil,omitempty" name:"ScalingPolicySet"`

	// Number of eligible notifications.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeScalingPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScalingPoliciesResponseParams `json:"Response"`
}

func (r *DescribeScalingPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScalingPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScheduledActionsRequestParams struct {
	// Queries by one or more scheduled task IDs in the format of asst-am691zxo. The maximum number of instances per request is 100. This parameter does not support specifying both ScheduledActionIds` and `Filters` at the same time.
	ScheduledActionIds []*string `json:"ScheduledActionIds,omitnil,omitempty" name:"ScheduledActionIds"`

	// Filter.
	// <li> scheduled-action-id - String - Required: No - (Filter) Filter by scheduled task ID.</li>
	// <li> scheduled-action-name - String - Required: No - (Filter) Filter by scheduled task name.</li>
	// <li> auto-scaling-group-id - String - Required: No - (Filter) Filter by auto scaling group ID.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. Default value: 0. For more information on `Offset`, see the relevant section in the API [overview](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, see the relevant section in the API [overview](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeScheduledActionsRequest struct {
	*tchttp.BaseRequest
	
	// Queries by one or more scheduled task IDs in the format of asst-am691zxo. The maximum number of instances per request is 100. This parameter does not support specifying both ScheduledActionIds` and `Filters` at the same time.
	ScheduledActionIds []*string `json:"ScheduledActionIds,omitnil,omitempty" name:"ScheduledActionIds"`

	// Filter.
	// <li> scheduled-action-id - String - Required: No - (Filter) Filter by scheduled task ID.</li>
	// <li> scheduled-action-name - String - Required: No - (Filter) Filter by scheduled task name.</li>
	// <li> auto-scaling-group-id - String - Required: No - (Filter) Filter by auto scaling group ID.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. Default value: 0. For more information on `Offset`, see the relevant section in the API [overview](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, see the relevant section in the API [overview](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeScheduledActionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScheduledActionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScheduledActionIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScheduledActionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScheduledActionsResponseParams struct {
	// Number of eligible scheduled tasks.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of scheduled task details.
	ScheduledActionSet []*ScheduledAction `json:"ScheduledActionSet,omitnil,omitempty" name:"ScheduledActionSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeScheduledActionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScheduledActionsResponseParams `json:"Response"`
}

func (r *DescribeScheduledActionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScheduledActionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachInstancesRequestParams struct {
	// Auto scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// List of CVM instance IDs
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type DetachInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Auto scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// List of CVM instance IDs
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *DetachInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachInstancesResponseParams struct {
	// Scaling activity ID
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DetachInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DetachInstancesResponseParams `json:"Response"`
}

func (r *DetachInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachLoadBalancersRequestParams struct {
	// Scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// List of classic CLB IDs. Up to 20 IDs are allowed. `LoadBalancerIds` and `ForwardLoadBalancerIdentifications` cannot be specified at the same time.
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// List of application CLB IDs. Up to 100 IDs are allowed. `LoadBalancerIds` and `ForwardLoadBalancerIdentifications` cannot be specified at the same time.
	ForwardLoadBalancerIdentifications []*ForwardLoadBalancerIdentification `json:"ForwardLoadBalancerIdentifications,omitnil,omitempty" name:"ForwardLoadBalancerIdentifications"`
}

type DetachLoadBalancersRequest struct {
	*tchttp.BaseRequest
	
	// Scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// List of classic CLB IDs. Up to 20 IDs are allowed. `LoadBalancerIds` and `ForwardLoadBalancerIdentifications` cannot be specified at the same time.
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// List of application CLB IDs. Up to 100 IDs are allowed. `LoadBalancerIds` and `ForwardLoadBalancerIdentifications` cannot be specified at the same time.
	ForwardLoadBalancerIdentifications []*ForwardLoadBalancerIdentification `json:"ForwardLoadBalancerIdentifications,omitnil,omitempty" name:"ForwardLoadBalancerIdentifications"`
}

func (r *DetachLoadBalancersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachLoadBalancersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "LoadBalancerIds")
	delete(f, "ForwardLoadBalancerIdentifications")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachLoadBalancersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachLoadBalancersResponseParams struct {
	// Scaling activity ID
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DetachLoadBalancersResponse struct {
	*tchttp.BaseResponse
	Response *DetachLoadBalancersResponseParams `json:"Response"`
}

func (r *DetachLoadBalancersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachLoadBalancersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetailedStatusMessage struct {
	// Error type
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// AZ information
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance billing mode
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Error message
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Instance type
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`
}

// Predefined struct for user
type DisableAutoScalingGroupRequestParams struct {
	// Scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`
}

type DisableAutoScalingGroupRequest struct {
	*tchttp.BaseRequest
	
	// Scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`
}

func (r *DisableAutoScalingGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableAutoScalingGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableAutoScalingGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableAutoScalingGroupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisableAutoScalingGroupResponse struct {
	*tchttp.BaseResponse
	Response *DisableAutoScalingGroupResponseParams `json:"Response"`
}

func (r *DisableAutoScalingGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableAutoScalingGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableAutoScalingGroupRequestParams struct {
	// Auto scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`
}

type EnableAutoScalingGroupRequest struct {
	*tchttp.BaseRequest
	
	// Auto scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`
}

func (r *EnableAutoScalingGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableAutoScalingGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableAutoScalingGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableAutoScalingGroupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EnableAutoScalingGroupResponse struct {
	*tchttp.BaseResponse
	Response *EnableAutoScalingGroupResponseParams `json:"Response"`
}

func (r *EnableAutoScalingGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableAutoScalingGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnhancedService struct {
	// Enables the Cloud Security service. If this parameter is not specified, the Cloud Security service will be enabled by default.
	SecurityService *RunSecurityServiceEnabled `json:"SecurityService,omitnil,omitempty" name:"SecurityService"`

	// Enables the Cloud Monitor service. If this parameter is not specified, the Cloud Monitor service will be enabled by default.
	MonitorService *RunMonitorServiceEnabled `json:"MonitorService,omitnil,omitempty" name:"MonitorService"`

	// Deprecated parameter.
	AutomationService []*RunAutomationServiceEnabled `json:"AutomationService,omitnil,omitempty" name:"AutomationService"`

	// Enable TAT service. If this parameter is not specified, the default logic is the same as that of the CVM instance. Note: This field may return `null`, indicating that no valid values can be obtained.
	AutomationToolsService *RunAutomationServiceEnabled `json:"AutomationToolsService,omitnil,omitempty" name:"AutomationToolsService"`
}

// Predefined struct for user
type ExecuteScalingPolicyRequestParams struct {
	// Auto-scaling policy ID. This parameter is not available to a target tracking policy.
	AutoScalingPolicyId *string `json:"AutoScalingPolicyId,omitnil,omitempty" name:"AutoScalingPolicyId"`

	// Whether to check if the auto scaling group is in the cooldown period. Default value: false
	HonorCooldown *bool `json:"HonorCooldown,omitnil,omitempty" name:"HonorCooldown"`

	// Source that triggers the scaling policy. Valid values: API and CLOUD_MONITOR. Default value: API. The value `CLOUD_MONITOR` is specific to the Cloud Monitor service.
	TriggerSource *string `json:"TriggerSource,omitnil,omitempty" name:"TriggerSource"`
}

type ExecuteScalingPolicyRequest struct {
	*tchttp.BaseRequest
	
	// Auto-scaling policy ID. This parameter is not available to a target tracking policy.
	AutoScalingPolicyId *string `json:"AutoScalingPolicyId,omitnil,omitempty" name:"AutoScalingPolicyId"`

	// Whether to check if the auto scaling group is in the cooldown period. Default value: false
	HonorCooldown *bool `json:"HonorCooldown,omitnil,omitempty" name:"HonorCooldown"`

	// Source that triggers the scaling policy. Valid values: API and CLOUD_MONITOR. Default value: API. The value `CLOUD_MONITOR` is specific to the Cloud Monitor service.
	TriggerSource *string `json:"TriggerSource,omitnil,omitempty" name:"TriggerSource"`
}

func (r *ExecuteScalingPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExecuteScalingPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingPolicyId")
	delete(f, "HonorCooldown")
	delete(f, "TriggerSource")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExecuteScalingPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExecuteScalingPolicyResponseParams struct {
	// Scaling activity ID
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ExecuteScalingPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ExecuteScalingPolicyResponseParams `json:"Response"`
}

func (r *ExecuteScalingPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExecuteScalingPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// Field to be filtered.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Filter value of the field.
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type ForwardLoadBalancer struct {
	// Load balancer ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// Application load balancer listener ID
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// List of target rule attributes
	TargetAttributes []*TargetAttribute `json:"TargetAttributes,omitnil,omitempty" name:"TargetAttributes"`

	// ID of a forwarding rule. This parameter is required for layer-7 listeners.
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// The region of CLB instance. It defaults to the region of AS service and is in the format of the common parameter `Region`, such as `ap-guangzhou`.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}

type ForwardLoadBalancerIdentification struct {
	// ID of the CLB
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// Application CLB listener ID
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// ID of a forwarding rule. This parameter is required for layer-7 listeners.
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`
}

type HostNameSettings struct {
	// Hostname of a CVM
	// <br><li>The `HostName` cannot start or end with a period (.) or hyphen (-), and cannot contain consecutive periods and hyphens.
	// <br><li>This field is unavailable to CVM instances.
	// <br><li>Other types of instances (such as Linux): the name contains 2 to 40 characters, and supports multiple periods (.). The string between two periods can consist of letters (case insensitive), numbers, and hyphens (-), and cannot be all numbers.
	// Note: this field may return `null`, indicating that no valid value is obtained.
	HostName *string `json:"HostName,omitnil,omitempty" name:"HostName"`

	// Type of CVM host name. Valid values: "ORIGINAL" and "UNIQUE". Default value: "ORIGINAL"
	// <br><li> ORIGINAL. Auto Scaling transfers the HostName set in input parameters to the CVM directly. CVM may adds serial numbers for the HostName. The HostName of instances within the auto scaling group may conflict.
	// <br><li> UNIQUE. The HostName set in input parameters is the prefix of a host name. Auto Scaling and CVM expand it. The HostName of an instance in the auto scaling group is unique.
	// Note: This field may return null, indicating that no valid values can be obtained.
	HostNameStyle *string `json:"HostNameStyle,omitnil,omitempty" name:"HostNameStyle"`
}

type IPv6InternetAccessible struct {
	// Network billing mode. Valid values: TRAFFIC_POSTPAID_BY_HOUR, BANDWIDTH_PACKAGE. Default value: TRAFFIC_POSTPAID_BY_HOUR. For the current account type, see [Account Type Description](https://intl.cloud.tencent.com/document/product/1199/49090?from_cn_redirect=1#judge).
	// <br><li> IPv6 supports `TRAFFIC_POSTPAID_BY_HOUR` under a bill-by-IP account.
	// <br><li> IPv6 supports `BANDWIDTH_PACKAGE` under a bill-by-CVM account.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	InternetChargeType *string `json:"InternetChargeType,omitnil,omitempty" name:"InternetChargeType"`

	// Outbound bandwidth cap of the public network (in Mbps). <br>It defaults to `0`, which indicates no public network bandwidth is allocated to IPv6. The value range of bandwidth caps varies with the model, availability zone and billing mode. For more information, see [Public Network Bandwidth Cap](https://intl.cloud.tencent.com/document/product/213/12523?from_cn_redirect=1).
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`

	// Bandwidth package ID. You can obtain the ID from the `BandwidthPackageId` field in the response of the [DescribeBandwidthPackages](https://intl.cloud.tencent.com/document/api/215/19209?from_cn_redirect=1) API.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil,omitempty" name:"BandwidthPackageId"`
}

type Instance struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Auto scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// Launch configuration ID
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitnil,omitempty" name:"LaunchConfigurationId"`

	// Launch configuration name
	LaunchConfigurationName *string `json:"LaunchConfigurationName,omitnil,omitempty" name:"LaunchConfigurationName"`

	// Lifecycle status. Valid values:<br>
	// <li>`IN_SERVICE`: The instance is running.
	// <li>`CREATING`: The instance is being created.
	// <li>`CREATION_FAILED`: The instance fails to be created.
	// <li>`TERMINATING`: The instance is being terminated.
	// <li>`TERMINATION_FAILED`: The instance fails to be terminated.
	// <li>`ATTACHING`: The instance is being bound.
	// <li>`ATTACH_FAILED`: The instance fails to be bound.
	// <li>`DETACHING`: The instance is being unbound.
	// <li>`DETACH_FAILED`: The instance fails to be unbound.
	// <li>`ATTACHING_LB`: The LB is being bound.
	// <li>DETACHING_LB: The LB is being unbound.
	// <li>`MODIFYING_LB`: The LB is being modified.
	// <li>`STARTING`: The instance is being started up.
	// <li>`START_FAILED`: The instance fails to be started up.
	// <li>`STOPPING`: The instance is being shut down.
	// <li>`STOP_FAILED`: The instance fails to be shut down.
	// <li>`STOPPED`: The instance is shut down.
	// <li>`IN_LAUNCHING_HOOK`: The lifecycle hook is being scaled out.
	// <li>`IN_TERMINATING_HOOK`: The lifecycle hook is being scaled in.
	LifeCycleState *string `json:"LifeCycleState,omitnil,omitempty" name:"LifeCycleState"`

	// Health status. Value range: HEALTHY, UNHEALTHY
	HealthStatus *string `json:"HealthStatus,omitnil,omitempty" name:"HealthStatus"`

	// Whether to add scale-in protection
	ProtectedFromScaleIn *bool `json:"ProtectedFromScaleIn,omitnil,omitempty" name:"ProtectedFromScaleIn"`

	// Availability zone
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Creation type. Value range: AUTO_CREATION, MANUAL_ATTACHING.
	CreationType *string `json:"CreationType,omitnil,omitempty" name:"CreationType"`

	// Instance addition time
	AddTime *string `json:"AddTime,omitnil,omitempty" name:"AddTime"`

	// Instance type
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Version number
	VersionNumber *int64 `json:"VersionNumber,omitnil,omitempty" name:"VersionNumber"`

	// Auto scaling group name
	AutoScalingGroupName *string `json:"AutoScalingGroupName,omitnil,omitempty" name:"AutoScalingGroupName"`

	// Warming up status. Valid values:
	// <li>`WAITING_ENTER_WARMUP`: The instance is waiting to be warmed up.
	// <li>`NO_NEED_WARMUP`: Warming up is not required.
	// <li>`IN_WARMUP`: The instance is being warmed up.
	// <li>`AFTER_WARMUP`: Warming up is completed.
	WarmupStatus *string `json:"WarmupStatus,omitnil,omitempty" name:"WarmupStatus"`

	// Placement group ID. Only one is allowed.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil,omitempty" name:"DisasterRecoverGroupIds"`
}

type InstanceChargePrepaid struct {
	// Purchased usage period of an instance in months. Value range: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36.
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Auto-renewal flag. Value range: <br><li>NOTIFY_AND_AUTO_RENEW: Notify expiry and renew automatically <br><li>NOTIFY_AND_MANUAL_RENEW: Notify expiry but not renew automatically <br><li>DISABLE_NOTIFY_AND_MANUAL_RENEW: Neither notify expiry nor renew automatically <br><br>Default value: NOTIFY_AND_MANUAL_RENEW. If this parameter is specified as NOTIFY_AND_AUTO_RENEW, the instance will be automatically renewed on a monthly basis if the account balance is sufficient.
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

type InstanceMarketOptionsRequest struct {
	// Bidding-related options
	SpotOptions *SpotMarketOptions `json:"SpotOptions,omitnil,omitempty" name:"SpotOptions"`

	// Market option type. Currently, this only supports the value "spot"
	// Note: This field may return null, indicating that no valid values can be obtained.
	MarketType *string `json:"MarketType,omitnil,omitempty" name:"MarketType"`
}

type InstanceNameSettings struct {
	// CVM instance name
	// 
	// The `InstanceName` cannot start or end with a dot (.) or hyphen (-), and cannot contain consecutive dots and hyphens.
	// The name contains 2 to 40 characters, and supports multiple dots (.). The string between two dots can consist of letters (case-insensitive), numbers, and hyphens (-), and cannot be all numbers.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Type of CVM instance name. Valid values: `ORIGINAL` and `UNIQUE`. Default value: `ORIGINAL`.
	// 
	// `ORIGINAL`: Auto Scaling sends the input parameter `InstanceName` to the CVM directly. The CVM may append a serial number to the `InstanceName`. The `InstanceName` of the instances within the scaling group may conflict.
	// 
	// `UNIQUE`: the input parameter `InstanceName` is the prefix of an instance name. Auto Scaling and CVM expand it. The `InstanceName` of an instance in the scaling group is unique.
	InstanceNameStyle *string `json:"InstanceNameStyle,omitnil,omitempty" name:"InstanceNameStyle"`
}

type InstanceTag struct {
	// Tag key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Tag value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type InternetAccessible struct {
	// Network billing method. Value range: <br><li>BANDWIDTH_PREPAID: Prepaid by bandwidth <br><li>TRAFFIC_POSTPAID_BY_HOUR: Postpaid by traffic on a per hour basis <br><li>BANDWIDTH_POSTPAID_BY_HOUR: Postpaid by bandwidth on a per hour basis <br><li>BANDWIDTH_PACKAGE: BWP user <br>Default value: TRAFFIC_POSTPAID_BY_HOUR.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InternetChargeType *string `json:"InternetChargeType,omitnil,omitempty" name:"InternetChargeType"`

	// The maximum outbound bandwidth in Mbps of the public network. The default value is 0 Mbps. The upper limit of bandwidth varies by model. For more information, see [Purchase Network Bandwidth](https://intl.cloud.tencent.com/document/product/213/509?from_cn_redirect=1).
	// Note: This field may return null, indicating that no valid values can be obtained.
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`

	// Whether to assign a public IP. Value range: <br><li>TRUE: Assign a public IP <br><li>FALSE: Do not assign a public IP <br><br>If the public network bandwidth is greater than 0 Mbps, you are free to choose whether to enable the public IP (which is enabled by default). If the public network bandwidth is 0 Mbps, no public IP will be allowed to be assigned.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PublicIpAssigned *bool `json:"PublicIpAssigned,omitnil,omitempty" name:"PublicIpAssigned"`

	// Bandwidth package ID. You can obtain the ID from the `BandwidthPackageId` field in the response of the [DescribeBandwidthPackages](https://intl.cloud.tencent.com/document/api/215/19209?from_cn_redirect=1) API.
	// Note: this field may return null, indicating that no valid value was found.
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil,omitempty" name:"BandwidthPackageId"`
}

type InvocationResult struct {
	// Instance ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Execution activity ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InvocationId *string `json:"InvocationId,omitnil,omitempty" name:"InvocationId"`

	// Execution task ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InvocationTaskId *string `json:"InvocationTaskId,omitnil,omitempty" name:"InvocationTaskId"`

	// Command ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CommandId *string `json:"CommandId,omitnil,omitempty" name:"CommandId"`

	// Execution Status
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskStatus *string `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// Execution exception information
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`
}

type LaunchConfiguration struct {
	// Project ID of the instance.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Launch configuration ID
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitnil,omitempty" name:"LaunchConfigurationId"`

	// Launch configuration name.
	LaunchConfigurationName *string `json:"LaunchConfigurationName,omitnil,omitempty" name:"LaunchConfigurationName"`

	// Instance model.
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Information of the instance's system disk configuration.
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// Information of the instance's data disk configuration.
	DataDisks []*DataDisk `json:"DataDisks,omitnil,omitempty" name:"DataDisks"`

	// Instance login settings.
	LoginSettings *LimitedLoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// Information of the public network bandwidth configuration.
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// Security group of the instance.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Auto scaling group associated with the launch configuration.
	AutoScalingGroupAbstractSet []*AutoScalingGroupAbstract `json:"AutoScalingGroupAbstractSet,omitnil,omitempty" name:"AutoScalingGroupAbstractSet"`

	// Custom data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`

	// Creation time of the launch configuration.
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// Conditions of enhancement services for the instance and their settings.
	EnhancedService *EnhancedService `json:"EnhancedService,omitnil,omitempty" name:"EnhancedService"`

	// Image ID.
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// Current status of the launch configuration. Value range: <br><li>NORMAL: normal <br><li>IMAGE_ABNORMAL: Exception with the image of the launch configuration <br><li>CBS_SNAP_ABNORMAL: Exception with the data disk snapshot of the launch configuration <br><li>SECURITY_GROUP_ABNORMAL: Exception with the security group of the launch configuration<br>
	LaunchConfigurationStatus *string `json:"LaunchConfigurationStatus,omitnil,omitempty" name:"LaunchConfigurationStatus"`

	// Instance billing mode. CVM instances take `POSTPAID_BY_HOUR` by default. Valid values:
	// <br><li>POSTPAID_BY_HOUR: pay-as-you-go hourly
	// <br><li>SPOTPAID: spot instance
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// Market options of the instance, such as parameters related to spot instances. This parameter is required for spot instances.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceMarketOptions *InstanceMarketOptionsRequest `json:"InstanceMarketOptions,omitnil,omitempty" name:"InstanceMarketOptions"`

	// List of instance models.
	InstanceTypes []*string `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// List of instance tags, which will be added to instances created by the scale-out activity. Up to 10 tags allowed.
	InstanceTags []*InstanceTag `json:"InstanceTags,omitnil,omitempty" name:"InstanceTags"`

	// Tag list.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Version
	VersionNumber *int64 `json:"VersionNumber,omitnil,omitempty" name:"VersionNumber"`

	// Update time
	UpdatedTime *string `json:"UpdatedTime,omitnil,omitempty" name:"UpdatedTime"`

	// CAM role name. This parameter can be obtained from the `roleName` field returned by DescribeRoleList API.
	CamRoleName *string `json:"CamRoleName,omitnil,omitempty" name:"CamRoleName"`

	// Value of InstanceTypesCheckPolicy upon the last operation.
	LastOperationInstanceTypesCheckPolicy *string `json:"LastOperationInstanceTypesCheckPolicy,omitnil,omitempty" name:"LastOperationInstanceTypesCheckPolicy"`

	// CVM hostname settings.
	HostNameSettings *HostNameSettings `json:"HostNameSettings,omitnil,omitempty" name:"HostNameSettings"`

	// Settings of CVM instance names
	InstanceNameSettings *InstanceNameSettings `json:"InstanceNameSettings,omitnil,omitempty" name:"InstanceNameSettings"`

	// Details of the monthly subscription, including the purchase period, auto-renewal. It is required if the `InstanceChargeType` is `PREPAID`.
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// Selection policy of cloud disks. Default value: ORIGINAL. Valid values:
	// <br><li>ORIGINAL: uses the configured cloud disk type
	// <br><li>AUTOMATIC: automatically chooses an available cloud disk type in the current availability zone
	DiskTypePolicy *string `json:"DiskTypePolicy,omitnil,omitempty" name:"DiskTypePolicy"`

	// HPC ID<br>
	// Note: This field is default to empty
	HpcClusterId *string `json:"HpcClusterId,omitnil,omitempty" name:"HpcClusterId"`

	// IPv6 public network bandwidth configuration.
	IPv6InternetAccessible *IPv6InternetAccessible `json:"IPv6InternetAccessible,omitnil,omitempty" name:"IPv6InternetAccessible"`
}

type LifecycleActionResultInfo struct {
	// ID of the lifecycle hook
	LifecycleHookId *string `json:"LifecycleHookId,omitnil,omitempty" name:"LifecycleHookId"`

	// ID of the instance
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Execution task ID. You can query the result by using the [DescribeInvocations](https://intl.cloud.tencent.com/document/api/1340/52679?from_cn_redirect=1) API of TAT. 
	InvocationId *string `json:"InvocationId,omitnil,omitempty" name:"InvocationId"`

	// Result of command invocation,
	// <li>`SUCCESSFUL`: Successful command invocation. It does mean that the task is successfully. You can query the task result with the `InvocationId.</li>
	// <li>`FAILED`: Failed to invoke the command</li>
	// <li>`NONE`</li>
	InvokeCommandResult *string `json:"InvokeCommandResult,omitnil,omitempty" name:"InvokeCommandResult"`

	// Notification result, which indicates whether it is successful to notify CMQ/TDMQ.<br>
	// <li>SUCCESSFUL: It is successful to notify CMQ/TDMQ.</li>
	// <li>FAILED: It is failed to notify CMQ/TDMQ.</li>
	// <li>NONE</li>
	NotificationResult *string `json:"NotificationResult,omitnil,omitempty" name:"NotificationResult"`

	// Result of the lifecyle hook action. Values: CONTINUE, ABANDON
	LifecycleActionResult *string `json:"LifecycleActionResult,omitnil,omitempty" name:"LifecycleActionResult"`

	// Reason of the result <br>
	// <li>`HEARTBEAT_TIMEOUT`: Heartbeat timed out. The setting of `DefaultResult` is used.</li>
	// <li>`NOTIFICATION_FAILURE`: Failed to send the notification. The setting of `DefaultResult` is used.</li>
	// <li>`CALL_INTERFACE`: Calls the `CompleteLifecycleAction` to set the result</li>
	// <li>ANOTHER_ACTION_ABANDON: It has been set to `ABANDON` by another operation.</li>
	// <li>COMMAND_CALL_FAILURE: Failed to call the command. The DefaultResult is applied.</li>
	// <li>COMMAND_EXEC_FINISH: Command completed</li>
	// <li>COMMAND_CALL_FAILURE: Failed to execute the command. The DefaultResult is applied.</li>
	// <li>COMMAND_EXEC_RESULT_CHECK_FAILURE: Failed to check the command result. The DefaultResult is applied.</li>
	ResultReason *string `json:"ResultReason,omitnil,omitempty" name:"ResultReason"`
}

type LifecycleCommand struct {
	// Remote command ID. It’s required to execute a command.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CommandId *string `json:"CommandId,omitnil,omitempty" name:"CommandId"`

	// Custom parameter. The field type is JSON encoded string. For example, {"varA": "222"}.
	// `key` is the name of the custom parameter and `value` is the default value. Both `key` and `value` are strings.
	// If this parameter is not specified, the `DefaultParameters` of `Command` is used.
	// Up to 20 customer parameters allowed. The parameter name can contain up to 64 characters, including [a-z], [A-Z], [0-9] and [-_].
	// Note: This field may return null, indicating that no valid values can be obtained.
	Parameters *string `json:"Parameters,omitnil,omitempty" name:"Parameters"`
}

type LifecycleHook struct {
	// Lifecycle hook ID
	LifecycleHookId *string `json:"LifecycleHookId,omitnil,omitempty" name:"LifecycleHookId"`

	// Lifecycle hook name
	LifecycleHookName *string `json:"LifecycleHookName,omitnil,omitempty" name:"LifecycleHookName"`

	// Auto scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// Default result of the lifecycle hook
	DefaultResult *string `json:"DefaultResult,omitnil,omitempty" name:"DefaultResult"`

	// Wait timeout period of the lifecycle hook
	HeartbeatTimeout *int64 `json:"HeartbeatTimeout,omitnil,omitempty" name:"HeartbeatTimeout"`

	// Applicable scenario of the lifecycle hook
	LifecycleTransition *string `json:"LifecycleTransition,omitnil,omitempty" name:"LifecycleTransition"`

	// Additional information for the notification target
	NotificationMetadata *string `json:"NotificationMetadata,omitnil,omitempty" name:"NotificationMetadata"`

	// Creation time
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// Notification target
	NotificationTarget *NotificationTarget `json:"NotificationTarget,omitnil,omitempty" name:"NotificationTarget"`

	// Applicable scenario of the lifecycle hook
	LifecycleTransitionType *string `json:"LifecycleTransitionType,omitnil,omitempty" name:"LifecycleTransitionType"`

	// Remote command execution object.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LifecycleCommand *LifecycleCommand `json:"LifecycleCommand,omitnil,omitempty" name:"LifecycleCommand"`
}

type LimitedLoginSettings struct {
	// List of key IDs.
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`
}

type LoginSettings struct {
	// Instance login password. <br><li>Linux: 8-16 characters. It should contain at least two sets of the following categories: [a-z], [A-Z], [0-9] and [()`~!@#$%^&*-+=|{}[]:;',.?/]. <br><li>Windows: 12-16 characters. It should contain at least three sets of the following categories: [a-z], [A-Z], [0-9] and [()`~!@#$%^&*-+={}[]:;',.?/]. <br><br>If this parameter is not specified, a random password is generated and sent to you via the Message Center.
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// List of key IDs. After an instance is associated with a key, you can access the instance with the private key in the key pair. You can call `DescribeKeyPairs` to obtain `KeyId`. Key and password cannot be specified at the same time. Windows instances do not support keys. Currently, you can only specify one key when purchasing an instance.
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`

	// Whether to keep the original settings of an image. It cannot be specified together with `Password` or `KeyIds.N`. You can specify this parameter as `TRUE` only when you create an instance using a custom image, a shared image, or an imported image. Valid values: <br><li>`TRUE`: Keep the login settings of the image <br><li>`FALSE` (Default): Do not keep the login settings of the image <br>
	KeepImageLogin *bool `json:"KeepImageLogin,omitnil,omitempty" name:"KeepImageLogin"`
}

type MetricAlarm struct {
	// Comparison operator. Value range: <br><li>GREATER_THAN: greater than </li><li>GREATER_THAN_OR_EQUAL_TO: greater than or equal to </li><li>LESS_THAN: less than </li><li> LESS_THAN_OR_EQUAL_TO: less than or equal to </li><li> EQUAL_TO: equal to </li> <li>NOT_EQUAL_TO: not equal to </li>
	ComparisonOperator *string `json:"ComparisonOperator,omitnil,omitempty" name:"ComparisonOperator"`

	// Metric name. Value range: <br><li>CPU_UTILIZATION: CPU utilization </li><li>MEM_UTILIZATION: memory utilization </li><li>LAN_TRAFFIC_OUT: private network outbound bandwidth </li><li>LAN_TRAFFIC_IN: private network inbound bandwidth </li><li>WAN_TRAFFIC_OUT: public network outbound bandwidth </li><li>WAN_TRAFFIC_IN: public network inbound bandwidth </li>
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// Alarming threshold: <br><li>CPU_UTILIZATION: [1, 100] in % </li><li>MEM_UTILIZATION: [1, 100] in % </li><li>LAN_TRAFFIC_OUT: >0 in Mbps </li><li>LAN_TRAFFIC_IN: >0 in Mbps </li><li>WAN_TRAFFIC_OUT: >0 in Mbps </li><li>WAN_TRAFFIC_IN: >0 in Mbps </li>
	Threshold *uint64 `json:"Threshold,omitnil,omitempty" name:"Threshold"`

	// Time period in seconds. Enumerated values: 60, 300.
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Number of repetitions. Value range: [1, 10]
	ContinuousTime *uint64 `json:"ContinuousTime,omitnil,omitempty" name:"ContinuousTime"`

	// Statistics type. Value range: <br><li>AVERAGE: average </li><li>MAXIMUM: maximum <li>MINIMUM: minimum </li><br> Default value: AVERAGE
	Statistic *string `json:"Statistic,omitnil,omitempty" name:"Statistic"`

	// Exact alarming threshold. This parameter is only used in API outputs. Values: <br><li>`CPU_UTILIZATION` (in %): (0, 100]</li><li>`MEM_UTILIZATION` (in %): (0, 100]</li><li>`LAN_TRAFFIC_OUT` (in Mbps): > 0</li><li>`LAN_TRAFFIC_IN` (in Mbps): > 0</li><li>`WAN_TRAFFIC_OUT` (in Mbps): > 0</li><li>`WAN_TRAFFIC_IN` (in Mbps): > 0</li>
	PreciseThreshold *float64 `json:"PreciseThreshold,omitnil,omitempty" name:"PreciseThreshold"`
}

// Predefined struct for user
type ModifyAutoScalingGroupRequestParams struct {
	// Auto scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// Auto scaling group name, which can only contain letters, numbers, underscores, hyphens ("-"), and decimal points with a maximum length of 55 bytes and must be unique under your account.
	AutoScalingGroupName *string `json:"AutoScalingGroupName,omitnil,omitempty" name:"AutoScalingGroupName"`

	// Default cooldown period in seconds. Default value: 300
	DefaultCooldown *uint64 `json:"DefaultCooldown,omitnil,omitempty" name:"DefaultCooldown"`

	// Desired number of instances. The number should be no larger than the maximum and no smaller than minimum number of instances
	DesiredCapacity *uint64 `json:"DesiredCapacity,omitnil,omitempty" name:"DesiredCapacity"`

	// Launch configuration ID
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitnil,omitempty" name:"LaunchConfigurationId"`

	// Maximum number of instances. Value range: 0-2,000.
	MaxSize *uint64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`

	// Minimum number of instances. Value range: 0-2,000.
	MinSize *uint64 `json:"MinSize,omitnil,omitempty" name:"MinSize"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// List of subnet IDs
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// Termination policy. Currently, the maximum length is 1. Value range: OLDEST_INSTANCE, NEWEST_INSTANCE.
	// <br><li> OLDEST_INSTANCE: The oldest instance in the auto scaling group will be terminated first.
	// <br><li> NEWEST_INSTANCE: The newest instance in the auto scaling group will be terminated first.
	TerminationPolicies []*string `json:"TerminationPolicies,omitnil,omitempty" name:"TerminationPolicies"`

	// VPC ID. This field is left empty for basic networks. You need to specify SubnetIds when modifying the network of the auto scaling group to a VPC with a specified VPC ID. Specify Zones when modifying the network to a basic network.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// List of availability zones
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// Retry policy. Valid values: `IMMEDIATE_RETRY` (default), `INCREMENTAL_INTERVALS`, `NO_RETRY`. A partially successful scaling is judged as a failed one.
	// <br><li>
	// `IMMEDIATE_RETRY`: Retrying immediately in a short period of time and stopping after five consecutive failures.
	// <br><li>
	// `INCREMENTAL_INTERVALS`: Retrying at incremental intervals. As the number of consecutive failures increases, the retry interval gradually increases, ranging from seconds to one day.
	// <br><li>`NO_RETRY`: Do not retry. Actions are taken when the next call or alarm message comes.
	RetryPolicy *string `json:"RetryPolicy,omitnil,omitempty" name:"RetryPolicy"`

	// Availability zone verification policy. Value range: ALL, ANY. Default value: ANY. This will work when the resource-related fields (launch configuration, availability zone, or subnet) of the auto scaling group are actually modified.
	// <br><li> ALL: The verification will succeed only if all availability zones (Zone) or subnets (SubnetId) are available; otherwise, an error will be reported.
	// <br><li> ANY: The verification will success if any availability zone (Zone) or subnet (SubnetId) is available; otherwise, an error will be reported.
	// 
	// Common reasons why an availability zone or subnet is unavailable include stock-out of CVM instances or CBS cloud disks in the availability zone, insufficient quota in the availability zone, or insufficient IPs in the subnet.
	// If an availability zone or subnet in Zones/SubnetIds does not exist, a verification error will be reported regardless of the value of ZonesCheckPolicy.
	ZonesCheckPolicy *string `json:"ZonesCheckPolicy,omitnil,omitempty" name:"ZonesCheckPolicy"`

	// Service settings such as unhealthy instance replacement.
	ServiceSettings *ServiceSettings `json:"ServiceSettings,omitnil,omitempty" name:"ServiceSettings"`

	// The number of IPv6 addresses that an instance has. Valid values: 0 and 1.
	Ipv6AddressCount *int64 `json:"Ipv6AddressCount,omitnil,omitempty" name:"Ipv6AddressCount"`

	// Multi-availability zone/subnet policy. Valid values: `PRIORITY` and `EQUALITY`. Default value: `PRIORITY`.
	// <br><li>`PRIORITY`: When an instance is being created, the availability zone/subnet is chosen from top to bottom in the list. The first availability zone/subnet is always used as long as instances can be created.
	// <br><li>`EQUALITY`: Instances created for scaling out are distributed to multiple availability zones/subnets, so as to keep the number of instances in different availability zone/subnet in balance.
	// 
	// Notes:
	// <br><li> When the scaling group is based on the classic network, this policy applies to multiple availability zones. When the scaling group is based on a VPC, this policy applies to multiple subnets, and you do not need to consider availability zones. For example, if you have four subnets (A, B, C, and D) and A, B, and C are in availability zone 1 and D is in availability zone 2, you only need to decide the order of the four subnets, without worrying about the issue of availability zones.
	// <br><li> This policy is applicable to multiple availability zones/subnets, but is not applicable to multiple models with launch configurations. Specify the models according to the model priority.
	// <br><li> When `PRIORITY` policy is used, the multi-model policy prevails the multi-availability zones/subnet policy. For example, if you have Model A/B, and Subnet 1/2/3, the model-subnet combinations are tried in the following order: A1 -> A2 -> A3 -> B1 -> B2 -> B3. If A1 is sold out, A2 (not B1) is tried next.
	MultiZoneSubnetPolicy *string `json:"MultiZoneSubnetPolicy,omitnil,omitempty" name:"MultiZoneSubnetPolicy"`

	// Health check type of instances in a scaling group.<br><li>CVM: confirm whether an instance is healthy based on the network status. If the pinged instance is unreachable, the instance will be considered unhealthy. For more information, see [Instance Health Check](https://intl.cloud.tencent.com/document/product/377/8553?from_cn_redirect=1)<br><li>CLB: confirm whether an instance is healthy based on the CLB health check status. For more information, see [Health Check Overview](https://intl.cloud.tencent.com/document/product/214/6097?from_cn_redirect=1).
	HealthCheckType *string `json:"HealthCheckType,omitnil,omitempty" name:"HealthCheckType"`

	// Grace period of the CLB health check
	LoadBalancerHealthCheckGracePeriod *uint64 `json:"LoadBalancerHealthCheckGracePeriod,omitnil,omitempty" name:"LoadBalancerHealthCheckGracePeriod"`

	// Specifies how to assign instances. Valid values: `LAUNCH_CONFIGURATION` and `SPOT_MIXED`.
	// <br><li>`LAUNCH_CONFIGURATION`: the launch configuration mode.
	// <br><li>`SPOT_MIXED`: a mixed instance mode. Currently, this mode is supported only when the launch configuration takes the pay-as-you-go billing mode. With this mode, the scaling group can provision a combination of pay-as-you-go instances and spot instances to meet the configured capacity. Note that the billing mode of the associated launch configuration cannot be modified when this mode is used.
	InstanceAllocationPolicy *string `json:"InstanceAllocationPolicy,omitnil,omitempty" name:"InstanceAllocationPolicy"`

	// Specifies how to assign pay-as-you-go instances and spot instances.
	// This parameter is valid only when `InstanceAllocationPolicy` is set to `SPOT_MIXED`.
	SpotMixedAllocationPolicy *SpotMixedAllocationPolicy `json:"SpotMixedAllocationPolicy,omitnil,omitempty" name:"SpotMixedAllocationPolicy"`

	// Indicates whether the capacity rebalancing feature is enabled. This parameter is only valid for spot instances in the scaling group. Valid values:
	// <br><li>`TRUE`: yes. Before the spot instances in the scaling group are about to be automatically repossessed, AS will terminate them. The scale-in hook (if configured) will take effect before the termination. After the termination process starts, AS will asynchronously initiate a scaling activity to meet the desired capacity.
	// <br><li>`FALSE`: no. In this case, AS will add instances to meet the desired capacity only after the spot instances are terminated.
	CapacityRebalance *bool `json:"CapacityRebalance,omitnil,omitempty" name:"CapacityRebalance"`
}

type ModifyAutoScalingGroupRequest struct {
	*tchttp.BaseRequest
	
	// Auto scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// Auto scaling group name, which can only contain letters, numbers, underscores, hyphens ("-"), and decimal points with a maximum length of 55 bytes and must be unique under your account.
	AutoScalingGroupName *string `json:"AutoScalingGroupName,omitnil,omitempty" name:"AutoScalingGroupName"`

	// Default cooldown period in seconds. Default value: 300
	DefaultCooldown *uint64 `json:"DefaultCooldown,omitnil,omitempty" name:"DefaultCooldown"`

	// Desired number of instances. The number should be no larger than the maximum and no smaller than minimum number of instances
	DesiredCapacity *uint64 `json:"DesiredCapacity,omitnil,omitempty" name:"DesiredCapacity"`

	// Launch configuration ID
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitnil,omitempty" name:"LaunchConfigurationId"`

	// Maximum number of instances. Value range: 0-2,000.
	MaxSize *uint64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`

	// Minimum number of instances. Value range: 0-2,000.
	MinSize *uint64 `json:"MinSize,omitnil,omitempty" name:"MinSize"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// List of subnet IDs
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// Termination policy. Currently, the maximum length is 1. Value range: OLDEST_INSTANCE, NEWEST_INSTANCE.
	// <br><li> OLDEST_INSTANCE: The oldest instance in the auto scaling group will be terminated first.
	// <br><li> NEWEST_INSTANCE: The newest instance in the auto scaling group will be terminated first.
	TerminationPolicies []*string `json:"TerminationPolicies,omitnil,omitempty" name:"TerminationPolicies"`

	// VPC ID. This field is left empty for basic networks. You need to specify SubnetIds when modifying the network of the auto scaling group to a VPC with a specified VPC ID. Specify Zones when modifying the network to a basic network.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// List of availability zones
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// Retry policy. Valid values: `IMMEDIATE_RETRY` (default), `INCREMENTAL_INTERVALS`, `NO_RETRY`. A partially successful scaling is judged as a failed one.
	// <br><li>
	// `IMMEDIATE_RETRY`: Retrying immediately in a short period of time and stopping after five consecutive failures.
	// <br><li>
	// `INCREMENTAL_INTERVALS`: Retrying at incremental intervals. As the number of consecutive failures increases, the retry interval gradually increases, ranging from seconds to one day.
	// <br><li>`NO_RETRY`: Do not retry. Actions are taken when the next call or alarm message comes.
	RetryPolicy *string `json:"RetryPolicy,omitnil,omitempty" name:"RetryPolicy"`

	// Availability zone verification policy. Value range: ALL, ANY. Default value: ANY. This will work when the resource-related fields (launch configuration, availability zone, or subnet) of the auto scaling group are actually modified.
	// <br><li> ALL: The verification will succeed only if all availability zones (Zone) or subnets (SubnetId) are available; otherwise, an error will be reported.
	// <br><li> ANY: The verification will success if any availability zone (Zone) or subnet (SubnetId) is available; otherwise, an error will be reported.
	// 
	// Common reasons why an availability zone or subnet is unavailable include stock-out of CVM instances or CBS cloud disks in the availability zone, insufficient quota in the availability zone, or insufficient IPs in the subnet.
	// If an availability zone or subnet in Zones/SubnetIds does not exist, a verification error will be reported regardless of the value of ZonesCheckPolicy.
	ZonesCheckPolicy *string `json:"ZonesCheckPolicy,omitnil,omitempty" name:"ZonesCheckPolicy"`

	// Service settings such as unhealthy instance replacement.
	ServiceSettings *ServiceSettings `json:"ServiceSettings,omitnil,omitempty" name:"ServiceSettings"`

	// The number of IPv6 addresses that an instance has. Valid values: 0 and 1.
	Ipv6AddressCount *int64 `json:"Ipv6AddressCount,omitnil,omitempty" name:"Ipv6AddressCount"`

	// Multi-availability zone/subnet policy. Valid values: `PRIORITY` and `EQUALITY`. Default value: `PRIORITY`.
	// <br><li>`PRIORITY`: When an instance is being created, the availability zone/subnet is chosen from top to bottom in the list. The first availability zone/subnet is always used as long as instances can be created.
	// <br><li>`EQUALITY`: Instances created for scaling out are distributed to multiple availability zones/subnets, so as to keep the number of instances in different availability zone/subnet in balance.
	// 
	// Notes:
	// <br><li> When the scaling group is based on the classic network, this policy applies to multiple availability zones. When the scaling group is based on a VPC, this policy applies to multiple subnets, and you do not need to consider availability zones. For example, if you have four subnets (A, B, C, and D) and A, B, and C are in availability zone 1 and D is in availability zone 2, you only need to decide the order of the four subnets, without worrying about the issue of availability zones.
	// <br><li> This policy is applicable to multiple availability zones/subnets, but is not applicable to multiple models with launch configurations. Specify the models according to the model priority.
	// <br><li> When `PRIORITY` policy is used, the multi-model policy prevails the multi-availability zones/subnet policy. For example, if you have Model A/B, and Subnet 1/2/3, the model-subnet combinations are tried in the following order: A1 -> A2 -> A3 -> B1 -> B2 -> B3. If A1 is sold out, A2 (not B1) is tried next.
	MultiZoneSubnetPolicy *string `json:"MultiZoneSubnetPolicy,omitnil,omitempty" name:"MultiZoneSubnetPolicy"`

	// Health check type of instances in a scaling group.<br><li>CVM: confirm whether an instance is healthy based on the network status. If the pinged instance is unreachable, the instance will be considered unhealthy. For more information, see [Instance Health Check](https://intl.cloud.tencent.com/document/product/377/8553?from_cn_redirect=1)<br><li>CLB: confirm whether an instance is healthy based on the CLB health check status. For more information, see [Health Check Overview](https://intl.cloud.tencent.com/document/product/214/6097?from_cn_redirect=1).
	HealthCheckType *string `json:"HealthCheckType,omitnil,omitempty" name:"HealthCheckType"`

	// Grace period of the CLB health check
	LoadBalancerHealthCheckGracePeriod *uint64 `json:"LoadBalancerHealthCheckGracePeriod,omitnil,omitempty" name:"LoadBalancerHealthCheckGracePeriod"`

	// Specifies how to assign instances. Valid values: `LAUNCH_CONFIGURATION` and `SPOT_MIXED`.
	// <br><li>`LAUNCH_CONFIGURATION`: the launch configuration mode.
	// <br><li>`SPOT_MIXED`: a mixed instance mode. Currently, this mode is supported only when the launch configuration takes the pay-as-you-go billing mode. With this mode, the scaling group can provision a combination of pay-as-you-go instances and spot instances to meet the configured capacity. Note that the billing mode of the associated launch configuration cannot be modified when this mode is used.
	InstanceAllocationPolicy *string `json:"InstanceAllocationPolicy,omitnil,omitempty" name:"InstanceAllocationPolicy"`

	// Specifies how to assign pay-as-you-go instances and spot instances.
	// This parameter is valid only when `InstanceAllocationPolicy` is set to `SPOT_MIXED`.
	SpotMixedAllocationPolicy *SpotMixedAllocationPolicy `json:"SpotMixedAllocationPolicy,omitnil,omitempty" name:"SpotMixedAllocationPolicy"`

	// Indicates whether the capacity rebalancing feature is enabled. This parameter is only valid for spot instances in the scaling group. Valid values:
	// <br><li>`TRUE`: yes. Before the spot instances in the scaling group are about to be automatically repossessed, AS will terminate them. The scale-in hook (if configured) will take effect before the termination. After the termination process starts, AS will asynchronously initiate a scaling activity to meet the desired capacity.
	// <br><li>`FALSE`: no. In this case, AS will add instances to meet the desired capacity only after the spot instances are terminated.
	CapacityRebalance *bool `json:"CapacityRebalance,omitnil,omitempty" name:"CapacityRebalance"`
}

func (r *ModifyAutoScalingGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoScalingGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "AutoScalingGroupName")
	delete(f, "DefaultCooldown")
	delete(f, "DesiredCapacity")
	delete(f, "LaunchConfigurationId")
	delete(f, "MaxSize")
	delete(f, "MinSize")
	delete(f, "ProjectId")
	delete(f, "SubnetIds")
	delete(f, "TerminationPolicies")
	delete(f, "VpcId")
	delete(f, "Zones")
	delete(f, "RetryPolicy")
	delete(f, "ZonesCheckPolicy")
	delete(f, "ServiceSettings")
	delete(f, "Ipv6AddressCount")
	delete(f, "MultiZoneSubnetPolicy")
	delete(f, "HealthCheckType")
	delete(f, "LoadBalancerHealthCheckGracePeriod")
	delete(f, "InstanceAllocationPolicy")
	delete(f, "SpotMixedAllocationPolicy")
	delete(f, "CapacityRebalance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAutoScalingGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAutoScalingGroupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAutoScalingGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAutoScalingGroupResponseParams `json:"Response"`
}

func (r *ModifyAutoScalingGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoScalingGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDesiredCapacityRequestParams struct {
	// Auto scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// Desired capacity
	DesiredCapacity *uint64 `json:"DesiredCapacity,omitnil,omitempty" name:"DesiredCapacity"`

	// Minimum number of instances. Value range: 0-2000.
	MinSize *uint64 `json:"MinSize,omitnil,omitempty" name:"MinSize"`

	// Maximum number of instances. Value range: 0-2000.
	MaxSize *uint64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`
}

type ModifyDesiredCapacityRequest struct {
	*tchttp.BaseRequest
	
	// Auto scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// Desired capacity
	DesiredCapacity *uint64 `json:"DesiredCapacity,omitnil,omitempty" name:"DesiredCapacity"`

	// Minimum number of instances. Value range: 0-2000.
	MinSize *uint64 `json:"MinSize,omitnil,omitempty" name:"MinSize"`

	// Maximum number of instances. Value range: 0-2000.
	MaxSize *uint64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`
}

func (r *ModifyDesiredCapacityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDesiredCapacityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "DesiredCapacity")
	delete(f, "MinSize")
	delete(f, "MaxSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDesiredCapacityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDesiredCapacityResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDesiredCapacityResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDesiredCapacityResponseParams `json:"Response"`
}

func (r *ModifyDesiredCapacityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDesiredCapacityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLaunchConfigurationAttributesRequestParams struct {
	// Launch configuration ID
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitnil,omitempty" name:"LaunchConfigurationId"`

	// [Image](https://intl.cloud.tencent.com/document/product/213/4940?from_cn_redirect=1) ID in the format of `img-xxx`. There are three types of images: <br/><li>Public images </li><li>Custom images </li><li>Shared images </li><br/>You can obtain the image IDs in the [CVM console](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE).</li><li>You can also use the [DescribeImages](https://intl.cloud.tencent.com/document/api/213/15715?from_cn_redirect=1) and look for `ImageId` in the response.</li>
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// List of instance types. Each type specifies different resource specifications. This list contains up to 10 instance types.
	// The launch configuration uses `InstanceType` to indicate one single instance type and `InstanceTypes` to indicate multiple instance types. Specifying the `InstanceTypes` field will invalidate the original `InstanceType`.
	InstanceTypes []*string `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// Instance type verification policy which works when InstanceTypes is actually modified. Value range: ALL, ANY. Default value: ANY.
	// <br><li> ALL: The verification will success only if all instance types (InstanceType) are available; otherwise, an error will be reported.
	// <br><li> ANY: The verification will success if any instance type (InstanceType) is available; otherwise, an error will be reported.
	// 
	// Common reasons why an instance type is unavailable include stock-out of the instance type or the corresponding cloud disk.
	// If a model in InstanceTypes does not exist or has been discontinued, a verification error will be reported regardless of the value of InstanceTypesCheckPolicy.
	InstanceTypesCheckPolicy *string `json:"InstanceTypesCheckPolicy,omitnil,omitempty" name:"InstanceTypesCheckPolicy"`

	// Display name of the launch configuration, which can contain Chinese characters, letters, numbers, underscores, separators ("-"), and decimal points with a maximum length of 60 bytes.
	LaunchConfigurationName *string `json:"LaunchConfigurationName,omitnil,omitempty" name:"LaunchConfigurationName"`

	// Base64-encoded custom data of up to 16 KB. If you want to clear `UserData`, set it to an empty string.
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`

	// Security group to which the instance belongs. This parameter can be obtained from the `SecurityGroupId` field in the response of the [`DescribeSecurityGroups`](https://intl.cloud.tencent.com/document/api/215/15808?from_cn_redirect=1) API.
	// At least one security group is required for this parameter. The security group specified is sequential.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Information of the public network bandwidth configuration.
	// When the public outbound network bandwidth is 0 Mbps, assigning a public IP is not allowed. Accordingly, if a public IP is assigned, the new public network outbound bandwidth must be greater than 0 Mbps.
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// Instance billing mode. Valid values:
	// <br><li>POSTPAID_BY_HOUR: pay-as-you-go hourly
	// <br><li>SPOTPAID: spot instance
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// Parameter setting for the prepaid mode (monthly subscription mode). This parameter can specify the renewal period, whether to set the auto-renewal, and other attributes of the monthly-subscribed instances.
	// This parameter is required when changing the instance billing mode to monthly subscription. It will be automatically discarded after you choose another billing mode.
	// This field requires passing in the `Period` field. Other fields that are not passed in will use their default values.
	// This field can be modified only when the current billing mode is monthly subscription.
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// Market-related options for instances, such as parameters related to spot instances.
	// This parameter is required when changing the instance billing mode to spot instance. It will be automatically discarded after you choose another instance billing mode.
	// This field requires passing in the `MaxPrice` field under the `SpotOptions`. Other fields that are not passed in will use their default values.
	// This field can be modified only when the current billing mode is spot instance.
	InstanceMarketOptions *InstanceMarketOptionsRequest `json:"InstanceMarketOptions,omitnil,omitempty" name:"InstanceMarketOptions"`

	// Selection policy of cloud disks. Default value: ORIGINAL. Valid values:
	// <br><li>ORIGINAL: uses the configured cloud disk type
	// <br><li>AUTOMATIC: automatically chooses an available cloud disk type
	DiskTypePolicy *string `json:"DiskTypePolicy,omitnil,omitempty" name:"DiskTypePolicy"`

	// Instance system disk configurations
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// Configuration information of instance data disks.
	// Up to 11 data disks can be specified and will be collectively modified. Please provide all the new values for the modification.
	// The default data disk should be the same as the system disk.
	DataDisks []*DataDisk `json:"DataDisks,omitnil,omitempty" name:"DataDisks"`

	// CVM hostname settings.
	// This field is not supported for Windows instances.
	// This field requires passing the `HostName` field. Other fields that are not passed in will use their default values.
	HostNameSettings *HostNameSettings `json:"HostNameSettings,omitnil,omitempty" name:"HostNameSettings"`

	// Settings of CVM instance names. 
	// If this field is configured in a launch configuration, the `InstanceName` of a CVM created by the scaling group will be generated according to the configuration; otherwise, it will be in the `as-{{AutoScalingGroupName }}` format.
	// This field requires passing in the `InstanceName` field. Other fields that are not passed in will use their default values.
	InstanceNameSettings *InstanceNameSettings `json:"InstanceNameSettings,omitnil,omitempty" name:"InstanceNameSettings"`

	// Specifies whether to enable additional services, such as security services and monitoring service.
	EnhancedService *EnhancedService `json:"EnhancedService,omitnil,omitempty" name:"EnhancedService"`

	// CAM role name. This parameter can be obtained from the `roleName` field returned by DescribeRoleList API.
	CamRoleName *string `json:"CamRoleName,omitnil,omitempty" name:"CamRoleName"`

	// HPC ID<br>
	// Note: This field is default to empty
	HpcClusterId *string `json:"HpcClusterId,omitnil,omitempty" name:"HpcClusterId"`

	// IPv6 public network bandwidth configuration. If the IPv6 address is available in the new instance, public network bandwidth can be allocated to the IPv6 address. This parameter is invalid if `Ipv6AddressCount` of the scaling group associated with the launch configuration is 0.
	IPv6InternetAccessible *IPv6InternetAccessible `json:"IPv6InternetAccessible,omitnil,omitempty" name:"IPv6InternetAccessible"`

	// Placement group ID. Only one is allowed.
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil,omitempty" name:"DisasterRecoverGroupIds"`
}

type ModifyLaunchConfigurationAttributesRequest struct {
	*tchttp.BaseRequest
	
	// Launch configuration ID
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitnil,omitempty" name:"LaunchConfigurationId"`

	// [Image](https://intl.cloud.tencent.com/document/product/213/4940?from_cn_redirect=1) ID in the format of `img-xxx`. There are three types of images: <br/><li>Public images </li><li>Custom images </li><li>Shared images </li><br/>You can obtain the image IDs in the [CVM console](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE).</li><li>You can also use the [DescribeImages](https://intl.cloud.tencent.com/document/api/213/15715?from_cn_redirect=1) and look for `ImageId` in the response.</li>
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// List of instance types. Each type specifies different resource specifications. This list contains up to 10 instance types.
	// The launch configuration uses `InstanceType` to indicate one single instance type and `InstanceTypes` to indicate multiple instance types. Specifying the `InstanceTypes` field will invalidate the original `InstanceType`.
	InstanceTypes []*string `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// Instance type verification policy which works when InstanceTypes is actually modified. Value range: ALL, ANY. Default value: ANY.
	// <br><li> ALL: The verification will success only if all instance types (InstanceType) are available; otherwise, an error will be reported.
	// <br><li> ANY: The verification will success if any instance type (InstanceType) is available; otherwise, an error will be reported.
	// 
	// Common reasons why an instance type is unavailable include stock-out of the instance type or the corresponding cloud disk.
	// If a model in InstanceTypes does not exist or has been discontinued, a verification error will be reported regardless of the value of InstanceTypesCheckPolicy.
	InstanceTypesCheckPolicy *string `json:"InstanceTypesCheckPolicy,omitnil,omitempty" name:"InstanceTypesCheckPolicy"`

	// Display name of the launch configuration, which can contain Chinese characters, letters, numbers, underscores, separators ("-"), and decimal points with a maximum length of 60 bytes.
	LaunchConfigurationName *string `json:"LaunchConfigurationName,omitnil,omitempty" name:"LaunchConfigurationName"`

	// Base64-encoded custom data of up to 16 KB. If you want to clear `UserData`, set it to an empty string.
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`

	// Security group to which the instance belongs. This parameter can be obtained from the `SecurityGroupId` field in the response of the [`DescribeSecurityGroups`](https://intl.cloud.tencent.com/document/api/215/15808?from_cn_redirect=1) API.
	// At least one security group is required for this parameter. The security group specified is sequential.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Information of the public network bandwidth configuration.
	// When the public outbound network bandwidth is 0 Mbps, assigning a public IP is not allowed. Accordingly, if a public IP is assigned, the new public network outbound bandwidth must be greater than 0 Mbps.
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// Instance billing mode. Valid values:
	// <br><li>POSTPAID_BY_HOUR: pay-as-you-go hourly
	// <br><li>SPOTPAID: spot instance
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// Parameter setting for the prepaid mode (monthly subscription mode). This parameter can specify the renewal period, whether to set the auto-renewal, and other attributes of the monthly-subscribed instances.
	// This parameter is required when changing the instance billing mode to monthly subscription. It will be automatically discarded after you choose another billing mode.
	// This field requires passing in the `Period` field. Other fields that are not passed in will use their default values.
	// This field can be modified only when the current billing mode is monthly subscription.
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// Market-related options for instances, such as parameters related to spot instances.
	// This parameter is required when changing the instance billing mode to spot instance. It will be automatically discarded after you choose another instance billing mode.
	// This field requires passing in the `MaxPrice` field under the `SpotOptions`. Other fields that are not passed in will use their default values.
	// This field can be modified only when the current billing mode is spot instance.
	InstanceMarketOptions *InstanceMarketOptionsRequest `json:"InstanceMarketOptions,omitnil,omitempty" name:"InstanceMarketOptions"`

	// Selection policy of cloud disks. Default value: ORIGINAL. Valid values:
	// <br><li>ORIGINAL: uses the configured cloud disk type
	// <br><li>AUTOMATIC: automatically chooses an available cloud disk type
	DiskTypePolicy *string `json:"DiskTypePolicy,omitnil,omitempty" name:"DiskTypePolicy"`

	// Instance system disk configurations
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// Configuration information of instance data disks.
	// Up to 11 data disks can be specified and will be collectively modified. Please provide all the new values for the modification.
	// The default data disk should be the same as the system disk.
	DataDisks []*DataDisk `json:"DataDisks,omitnil,omitempty" name:"DataDisks"`

	// CVM hostname settings.
	// This field is not supported for Windows instances.
	// This field requires passing the `HostName` field. Other fields that are not passed in will use their default values.
	HostNameSettings *HostNameSettings `json:"HostNameSettings,omitnil,omitempty" name:"HostNameSettings"`

	// Settings of CVM instance names. 
	// If this field is configured in a launch configuration, the `InstanceName` of a CVM created by the scaling group will be generated according to the configuration; otherwise, it will be in the `as-{{AutoScalingGroupName }}` format.
	// This field requires passing in the `InstanceName` field. Other fields that are not passed in will use their default values.
	InstanceNameSettings *InstanceNameSettings `json:"InstanceNameSettings,omitnil,omitempty" name:"InstanceNameSettings"`

	// Specifies whether to enable additional services, such as security services and monitoring service.
	EnhancedService *EnhancedService `json:"EnhancedService,omitnil,omitempty" name:"EnhancedService"`

	// CAM role name. This parameter can be obtained from the `roleName` field returned by DescribeRoleList API.
	CamRoleName *string `json:"CamRoleName,omitnil,omitempty" name:"CamRoleName"`

	// HPC ID<br>
	// Note: This field is default to empty
	HpcClusterId *string `json:"HpcClusterId,omitnil,omitempty" name:"HpcClusterId"`

	// IPv6 public network bandwidth configuration. If the IPv6 address is available in the new instance, public network bandwidth can be allocated to the IPv6 address. This parameter is invalid if `Ipv6AddressCount` of the scaling group associated with the launch configuration is 0.
	IPv6InternetAccessible *IPv6InternetAccessible `json:"IPv6InternetAccessible,omitnil,omitempty" name:"IPv6InternetAccessible"`

	// Placement group ID. Only one is allowed.
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil,omitempty" name:"DisasterRecoverGroupIds"`
}

func (r *ModifyLaunchConfigurationAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLaunchConfigurationAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LaunchConfigurationId")
	delete(f, "ImageId")
	delete(f, "InstanceTypes")
	delete(f, "InstanceTypesCheckPolicy")
	delete(f, "LaunchConfigurationName")
	delete(f, "UserData")
	delete(f, "SecurityGroupIds")
	delete(f, "InternetAccessible")
	delete(f, "InstanceChargeType")
	delete(f, "InstanceChargePrepaid")
	delete(f, "InstanceMarketOptions")
	delete(f, "DiskTypePolicy")
	delete(f, "SystemDisk")
	delete(f, "DataDisks")
	delete(f, "HostNameSettings")
	delete(f, "InstanceNameSettings")
	delete(f, "EnhancedService")
	delete(f, "CamRoleName")
	delete(f, "HpcClusterId")
	delete(f, "IPv6InternetAccessible")
	delete(f, "DisasterRecoverGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLaunchConfigurationAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLaunchConfigurationAttributesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLaunchConfigurationAttributesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLaunchConfigurationAttributesResponseParams `json:"Response"`
}

func (r *ModifyLaunchConfigurationAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLaunchConfigurationAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLifecycleHookRequestParams struct {
	// Lifecycle hook ID.
	LifecycleHookId *string `json:"LifecycleHookId,omitnil,omitempty" name:"LifecycleHookId"`

	// Lifecycle hook name.
	LifecycleHookName *string `json:"LifecycleHookName,omitnil,omitempty" name:"LifecycleHookName"`

	// The time when the lifecycle hook is applied. Valid values:
	// <li> `INSTANCE_LAUNCHING`: After the instance launch
	// <li> `INSTANCE_TERMINATING`: Before the instance termination
	LifecycleTransition *string `json:"LifecycleTransition,omitnil,omitempty" name:"LifecycleTransition"`

	// Actions after the lifecycle hook times out. Valid values:
	// <li> `CONTINUE`: Continue the scaling activity after the timeout
	// <li> `ABANDON`: Terminate the scaling activity after the timeout
	DefaultResult *string `json:"DefaultResult,omitnil,omitempty" name:"DefaultResult"`

	// The maximum length of time (in seconds) that can elapse before the lifecycle hook times out. Value range: 30 - 7,200 seconds.
	HeartbeatTimeout *uint64 `json:"HeartbeatTimeout,omitnil,omitempty" name:"HeartbeatTimeout"`

	// Additional information sent by AS to the notification target.
	NotificationMetadata *string `json:"NotificationMetadata,omitnil,omitempty" name:"NotificationMetadata"`

	// The scenario where the lifecycle hook is applied. `EXTENSION`: The lifecycle hook will be triggered when `AttachInstances`, `DetachInstances` or `RemoveInstances` is called. `NORMAL`: The lifecycle hook is not triggered by the above APIs.
	LifecycleTransitionType *string `json:"LifecycleTransitionType,omitnil,omitempty" name:"LifecycleTransitionType"`

	// Information of the notification target.
	NotificationTarget *NotificationTarget `json:"NotificationTarget,omitnil,omitempty" name:"NotificationTarget"`

	// Remote command execution object.
	LifecycleCommand *LifecycleCommand `json:"LifecycleCommand,omitnil,omitempty" name:"LifecycleCommand"`
}

type ModifyLifecycleHookRequest struct {
	*tchttp.BaseRequest
	
	// Lifecycle hook ID.
	LifecycleHookId *string `json:"LifecycleHookId,omitnil,omitempty" name:"LifecycleHookId"`

	// Lifecycle hook name.
	LifecycleHookName *string `json:"LifecycleHookName,omitnil,omitempty" name:"LifecycleHookName"`

	// The time when the lifecycle hook is applied. Valid values:
	// <li> `INSTANCE_LAUNCHING`: After the instance launch
	// <li> `INSTANCE_TERMINATING`: Before the instance termination
	LifecycleTransition *string `json:"LifecycleTransition,omitnil,omitempty" name:"LifecycleTransition"`

	// Actions after the lifecycle hook times out. Valid values:
	// <li> `CONTINUE`: Continue the scaling activity after the timeout
	// <li> `ABANDON`: Terminate the scaling activity after the timeout
	DefaultResult *string `json:"DefaultResult,omitnil,omitempty" name:"DefaultResult"`

	// The maximum length of time (in seconds) that can elapse before the lifecycle hook times out. Value range: 30 - 7,200 seconds.
	HeartbeatTimeout *uint64 `json:"HeartbeatTimeout,omitnil,omitempty" name:"HeartbeatTimeout"`

	// Additional information sent by AS to the notification target.
	NotificationMetadata *string `json:"NotificationMetadata,omitnil,omitempty" name:"NotificationMetadata"`

	// The scenario where the lifecycle hook is applied. `EXTENSION`: The lifecycle hook will be triggered when `AttachInstances`, `DetachInstances` or `RemoveInstances` is called. `NORMAL`: The lifecycle hook is not triggered by the above APIs.
	LifecycleTransitionType *string `json:"LifecycleTransitionType,omitnil,omitempty" name:"LifecycleTransitionType"`

	// Information of the notification target.
	NotificationTarget *NotificationTarget `json:"NotificationTarget,omitnil,omitempty" name:"NotificationTarget"`

	// Remote command execution object.
	LifecycleCommand *LifecycleCommand `json:"LifecycleCommand,omitnil,omitempty" name:"LifecycleCommand"`
}

func (r *ModifyLifecycleHookRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLifecycleHookRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LifecycleHookId")
	delete(f, "LifecycleHookName")
	delete(f, "LifecycleTransition")
	delete(f, "DefaultResult")
	delete(f, "HeartbeatTimeout")
	delete(f, "NotificationMetadata")
	delete(f, "LifecycleTransitionType")
	delete(f, "NotificationTarget")
	delete(f, "LifecycleCommand")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLifecycleHookRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLifecycleHookResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLifecycleHookResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLifecycleHookResponseParams `json:"Response"`
}

func (r *ModifyLifecycleHookResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLifecycleHookResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoadBalancerTargetAttributesRequestParams struct {
	// Scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// List of application CLBs to modify. Up to 100 CLBs allowed.
	ForwardLoadBalancers []*ForwardLoadBalancer `json:"ForwardLoadBalancers,omitnil,omitempty" name:"ForwardLoadBalancers"`
}

type ModifyLoadBalancerTargetAttributesRequest struct {
	*tchttp.BaseRequest
	
	// Scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// List of application CLBs to modify. Up to 100 CLBs allowed.
	ForwardLoadBalancers []*ForwardLoadBalancer `json:"ForwardLoadBalancers,omitnil,omitempty" name:"ForwardLoadBalancers"`
}

func (r *ModifyLoadBalancerTargetAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancerTargetAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "ForwardLoadBalancers")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLoadBalancerTargetAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoadBalancerTargetAttributesResponseParams struct {
	// Scaling activity ID
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLoadBalancerTargetAttributesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLoadBalancerTargetAttributesResponseParams `json:"Response"`
}

func (r *ModifyLoadBalancerTargetAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancerTargetAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoadBalancersRequestParams struct {
	// Auto scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// List of classic CLB IDs. Currently, the maximum length is 20. You cannot specify LoadBalancerIds and ForwardLoadBalancers at the same time.
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// List of application CLBs. Up to 100 CLBs are allowed. `LoadBalancerIds` and `ForwardLoadBalancers` cannot be specified at the same time.
	ForwardLoadBalancers []*ForwardLoadBalancer `json:"ForwardLoadBalancers,omitnil,omitempty" name:"ForwardLoadBalancers"`

	// CLB verification policy. Valid values: "ALL" and "DIFF". Default value: "ALL"
	// <br><li> ALL. Verification is successful only when all CLBs are valid. Otherwise, verification fails.
	// <br><li> DIFF. Only the changes in the CLB parameters are verified. If valid, the verification is successful. Otherwise, verification fails.
	LoadBalancersCheckPolicy *string `json:"LoadBalancersCheckPolicy,omitnil,omitempty" name:"LoadBalancersCheckPolicy"`
}

type ModifyLoadBalancersRequest struct {
	*tchttp.BaseRequest
	
	// Auto scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// List of classic CLB IDs. Currently, the maximum length is 20. You cannot specify LoadBalancerIds and ForwardLoadBalancers at the same time.
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// List of application CLBs. Up to 100 CLBs are allowed. `LoadBalancerIds` and `ForwardLoadBalancers` cannot be specified at the same time.
	ForwardLoadBalancers []*ForwardLoadBalancer `json:"ForwardLoadBalancers,omitnil,omitempty" name:"ForwardLoadBalancers"`

	// CLB verification policy. Valid values: "ALL" and "DIFF". Default value: "ALL"
	// <br><li> ALL. Verification is successful only when all CLBs are valid. Otherwise, verification fails.
	// <br><li> DIFF. Only the changes in the CLB parameters are verified. If valid, the verification is successful. Otherwise, verification fails.
	LoadBalancersCheckPolicy *string `json:"LoadBalancersCheckPolicy,omitnil,omitempty" name:"LoadBalancersCheckPolicy"`
}

func (r *ModifyLoadBalancersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "LoadBalancerIds")
	delete(f, "ForwardLoadBalancers")
	delete(f, "LoadBalancersCheckPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLoadBalancersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoadBalancersResponseParams struct {
	// Scaling activity ID
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLoadBalancersResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLoadBalancersResponseParams `json:"Response"`
}

func (r *ModifyLoadBalancersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNotificationConfigurationRequestParams struct {
	// ID of the notification to be modified.
	AutoScalingNotificationId *string `json:"AutoScalingNotificationId,omitnil,omitempty" name:"AutoScalingNotificationId"`

	// Notification type, i.e., the set of types of notifications to be subscribed to. Value range:
	// <li>SCALE_OUT_SUCCESSFUL: scale-out succeeded</li>
	// <li>SCALE_OUT_FAILED: scale-out failed</li>
	// <li>SCALE_IN_SUCCESSFUL: scale-in succeeded</li>
	// <li>SCALE_IN_FAILED: scale-in failed</li>
	// <li>REPLACE_UNHEALTHY_INSTANCE_SUCCESSFUL: unhealthy instance replacement succeeded</li>
	// <li>REPLACE_UNHEALTHY_INSTANCE_FAILED: unhealthy instance replacement failed</li>
	NotificationTypes []*string `json:"NotificationTypes,omitnil,omitempty" name:"NotificationTypes"`

	// Notification group ID, which is the set of user group IDs. You can query the user group IDs through the [ListGroups](https://intl.cloud.tencent.com/document/product/598/34589?from_cn_redirect=1) API.
	NotificationUserGroupIds []*string `json:"NotificationUserGroupIds,omitnil,omitempty" name:"NotificationUserGroupIds"`

	// CMQ or TDMQ CMQ queue name.
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// CMQ or TDMQ CMQ toipc name.
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

type ModifyNotificationConfigurationRequest struct {
	*tchttp.BaseRequest
	
	// ID of the notification to be modified.
	AutoScalingNotificationId *string `json:"AutoScalingNotificationId,omitnil,omitempty" name:"AutoScalingNotificationId"`

	// Notification type, i.e., the set of types of notifications to be subscribed to. Value range:
	// <li>SCALE_OUT_SUCCESSFUL: scale-out succeeded</li>
	// <li>SCALE_OUT_FAILED: scale-out failed</li>
	// <li>SCALE_IN_SUCCESSFUL: scale-in succeeded</li>
	// <li>SCALE_IN_FAILED: scale-in failed</li>
	// <li>REPLACE_UNHEALTHY_INSTANCE_SUCCESSFUL: unhealthy instance replacement succeeded</li>
	// <li>REPLACE_UNHEALTHY_INSTANCE_FAILED: unhealthy instance replacement failed</li>
	NotificationTypes []*string `json:"NotificationTypes,omitnil,omitempty" name:"NotificationTypes"`

	// Notification group ID, which is the set of user group IDs. You can query the user group IDs through the [ListGroups](https://intl.cloud.tencent.com/document/product/598/34589?from_cn_redirect=1) API.
	NotificationUserGroupIds []*string `json:"NotificationUserGroupIds,omitnil,omitempty" name:"NotificationUserGroupIds"`

	// CMQ or TDMQ CMQ queue name.
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// CMQ or TDMQ CMQ toipc name.
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

func (r *ModifyNotificationConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNotificationConfigurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingNotificationId")
	delete(f, "NotificationTypes")
	delete(f, "NotificationUserGroupIds")
	delete(f, "QueueName")
	delete(f, "TopicName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNotificationConfigurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNotificationConfigurationResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNotificationConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNotificationConfigurationResponseParams `json:"Response"`
}

func (r *ModifyNotificationConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNotificationConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyScalingPolicyRequestParams struct {
	// Alarm policy ID.
	AutoScalingPolicyId *string `json:"AutoScalingPolicyId,omitnil,omitempty" name:"AutoScalingPolicyId"`

	// Alarm policy name.
	ScalingPolicyName *string `json:"ScalingPolicyName,omitnil,omitempty" name:"ScalingPolicyName"`

	// The method to adjust the desired capacity after the alarm is triggered. It’s only available when `ScalingPolicyType` is `Simple`. Valid values: <br><li>`CHANGE_IN_CAPACITY`: Increase or decrease the desired capacity </li><li>`EXACT_CAPACITY`: Adjust to the specified desired capacity </li> <li>`PERCENT_CHANGE_IN_CAPACITY`: Adjust the desired capacity by percentage </li>
	AdjustmentType *string `json:"AdjustmentType,omitnil,omitempty" name:"AdjustmentType"`

	// Specifies how to adjust the number of desired capacity when the alarm is triggered. It’s only available when `ScalingPolicyType` is `Simple`. Values: <br><li>`AdjustmentType`=`CHANGE_IN_CAPACITY`: Number of instances to add (positive number) or remove (negative number). </li> <li>`AdjustmentType`=`EXACT_CAPACITY`: Set the desired capacity to the specified number. It must be ≥ 0. </li> <li>`AdjustmentType`=`PERCENT_CHANGE_IN_CAPACITY`: Percentage of instance number. Add instances (positive value) or remove instances (negative value) accordingly.
	AdjustmentValue *int64 `json:"AdjustmentValue,omitnil,omitempty" name:"AdjustmentValue"`

	// Cooldown period (in seconds). It’s only available when `ScalingPolicyType` is `Simple`.
	Cooldown *uint64 `json:"Cooldown,omitnil,omitempty" name:"Cooldown"`

	// Alarm monitoring metric. It’s only available when `ScalingPolicyType` is `Simple`.
	MetricAlarm *MetricAlarm `json:"MetricAlarm,omitnil,omitempty" name:"MetricAlarm"`

	// Preset monitoring item. It’s only available when `ScalingPolicyType` is `TARGET_TRACKING`. Valid values: <br><li>ASG_AVG_CPU_UTILIZATION: Average CPU utilization</li><li>ASG_AVG_LAN_TRAFFIC_OUT: Average private bandwidth out</li><li>ASG_AVG_LAN_TRAFFIC_IN: Average private bandwidth in</li><li>ASG_AVG_WAN_TRAFFIC_OUT: Average public bandwidth out</li><li>ASG_AVG_WAN_TRAFFIC_IN: Average public bandwidth in</li>
	PredefinedMetricType *string `json:"PredefinedMetricType,omitnil,omitempty" name:"PredefinedMetricType"`

	// Target value. It’s only available when `ScalingPolicyType` is `TARGET_TRACKING`. Value ranges: <br><li>`ASG_AVG_CPU_UTILIZATION` (in %): [1, 100)</li><li>`ASG_AVG_LAN_TRAFFIC_OUT` (in Mbps): >0</li><li>`ASG_AVG_LAN_TRAFFIC_IN` (in Mbps): >0</li><li>`ASG_AVG_WAN_TRAFFIC_OUT` (in Mbps): >0</li><li>`ASG_AVG_WAN_TRAFFIC_IN` (in Mbps): >0</li>
	TargetValue *uint64 `json:"TargetValue,omitnil,omitempty" name:"TargetValue"`

	// Instance warm-up period (in seconds). It’s only available when `ScalingPolicyType` is `TARGET_TRACKING`. Value range: 0-3600.
	EstimatedInstanceWarmup *uint64 `json:"EstimatedInstanceWarmup,omitnil,omitempty" name:"EstimatedInstanceWarmup"`

	// Whether to disable scale-in. It’s only available when `ScalingPolicyType` is `TARGET_TRACKING`. Valid values: <br><li>`true`: Scaling in is not allowed.</li><li>`false`: Allows both scale-out and scale-in</li>
	DisableScaleIn *bool `json:"DisableScaleIn,omitnil,omitempty" name:"DisableScaleIn"`

	// This parameter is diused. Please use [CreateNotificationConfiguration](https://intl.cloud.tencent.com/document/api/377/33185?from_cn_redirect=1) instead.
	// Notification group ID, which is the set of user group IDs.
	NotificationUserGroupIds []*string `json:"NotificationUserGroupIds,omitnil,omitempty" name:"NotificationUserGroupIds"`
}

type ModifyScalingPolicyRequest struct {
	*tchttp.BaseRequest
	
	// Alarm policy ID.
	AutoScalingPolicyId *string `json:"AutoScalingPolicyId,omitnil,omitempty" name:"AutoScalingPolicyId"`

	// Alarm policy name.
	ScalingPolicyName *string `json:"ScalingPolicyName,omitnil,omitempty" name:"ScalingPolicyName"`

	// The method to adjust the desired capacity after the alarm is triggered. It’s only available when `ScalingPolicyType` is `Simple`. Valid values: <br><li>`CHANGE_IN_CAPACITY`: Increase or decrease the desired capacity </li><li>`EXACT_CAPACITY`: Adjust to the specified desired capacity </li> <li>`PERCENT_CHANGE_IN_CAPACITY`: Adjust the desired capacity by percentage </li>
	AdjustmentType *string `json:"AdjustmentType,omitnil,omitempty" name:"AdjustmentType"`

	// Specifies how to adjust the number of desired capacity when the alarm is triggered. It’s only available when `ScalingPolicyType` is `Simple`. Values: <br><li>`AdjustmentType`=`CHANGE_IN_CAPACITY`: Number of instances to add (positive number) or remove (negative number). </li> <li>`AdjustmentType`=`EXACT_CAPACITY`: Set the desired capacity to the specified number. It must be ≥ 0. </li> <li>`AdjustmentType`=`PERCENT_CHANGE_IN_CAPACITY`: Percentage of instance number. Add instances (positive value) or remove instances (negative value) accordingly.
	AdjustmentValue *int64 `json:"AdjustmentValue,omitnil,omitempty" name:"AdjustmentValue"`

	// Cooldown period (in seconds). It’s only available when `ScalingPolicyType` is `Simple`.
	Cooldown *uint64 `json:"Cooldown,omitnil,omitempty" name:"Cooldown"`

	// Alarm monitoring metric. It’s only available when `ScalingPolicyType` is `Simple`.
	MetricAlarm *MetricAlarm `json:"MetricAlarm,omitnil,omitempty" name:"MetricAlarm"`

	// Preset monitoring item. It’s only available when `ScalingPolicyType` is `TARGET_TRACKING`. Valid values: <br><li>ASG_AVG_CPU_UTILIZATION: Average CPU utilization</li><li>ASG_AVG_LAN_TRAFFIC_OUT: Average private bandwidth out</li><li>ASG_AVG_LAN_TRAFFIC_IN: Average private bandwidth in</li><li>ASG_AVG_WAN_TRAFFIC_OUT: Average public bandwidth out</li><li>ASG_AVG_WAN_TRAFFIC_IN: Average public bandwidth in</li>
	PredefinedMetricType *string `json:"PredefinedMetricType,omitnil,omitempty" name:"PredefinedMetricType"`

	// Target value. It’s only available when `ScalingPolicyType` is `TARGET_TRACKING`. Value ranges: <br><li>`ASG_AVG_CPU_UTILIZATION` (in %): [1, 100)</li><li>`ASG_AVG_LAN_TRAFFIC_OUT` (in Mbps): >0</li><li>`ASG_AVG_LAN_TRAFFIC_IN` (in Mbps): >0</li><li>`ASG_AVG_WAN_TRAFFIC_OUT` (in Mbps): >0</li><li>`ASG_AVG_WAN_TRAFFIC_IN` (in Mbps): >0</li>
	TargetValue *uint64 `json:"TargetValue,omitnil,omitempty" name:"TargetValue"`

	// Instance warm-up period (in seconds). It’s only available when `ScalingPolicyType` is `TARGET_TRACKING`. Value range: 0-3600.
	EstimatedInstanceWarmup *uint64 `json:"EstimatedInstanceWarmup,omitnil,omitempty" name:"EstimatedInstanceWarmup"`

	// Whether to disable scale-in. It’s only available when `ScalingPolicyType` is `TARGET_TRACKING`. Valid values: <br><li>`true`: Scaling in is not allowed.</li><li>`false`: Allows both scale-out and scale-in</li>
	DisableScaleIn *bool `json:"DisableScaleIn,omitnil,omitempty" name:"DisableScaleIn"`

	// This parameter is diused. Please use [CreateNotificationConfiguration](https://intl.cloud.tencent.com/document/api/377/33185?from_cn_redirect=1) instead.
	// Notification group ID, which is the set of user group IDs.
	NotificationUserGroupIds []*string `json:"NotificationUserGroupIds,omitnil,omitempty" name:"NotificationUserGroupIds"`
}

func (r *ModifyScalingPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyScalingPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingPolicyId")
	delete(f, "ScalingPolicyName")
	delete(f, "AdjustmentType")
	delete(f, "AdjustmentValue")
	delete(f, "Cooldown")
	delete(f, "MetricAlarm")
	delete(f, "PredefinedMetricType")
	delete(f, "TargetValue")
	delete(f, "EstimatedInstanceWarmup")
	delete(f, "DisableScaleIn")
	delete(f, "NotificationUserGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyScalingPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyScalingPolicyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyScalingPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyScalingPolicyResponseParams `json:"Response"`
}

func (r *ModifyScalingPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyScalingPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyScheduledActionRequestParams struct {
	// ID of the scheduled task to be edited
	ScheduledActionId *string `json:"ScheduledActionId,omitnil,omitempty" name:"ScheduledActionId"`

	// Scheduled task name, which can only contain letters, numbers, underscores, hyphens ("-"), and decimal points with a maximum length of 60 bytes and must be unique in an auto scaling group.
	ScheduledActionName *string `json:"ScheduledActionName,omitnil,omitempty" name:"ScheduledActionName"`

	// The maximum number of instances set for the auto scaling group when the scheduled task is triggered.
	MaxSize *uint64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`

	// The minimum number of instances set for the auto scaling group when the scheduled task is triggered.
	MinSize *uint64 `json:"MinSize,omitnil,omitempty" name:"MinSize"`

	// The desired number of instances set for the auto scaling group when the scheduled task is triggered.
	DesiredCapacity *uint64 `json:"DesiredCapacity,omitnil,omitempty" name:"DesiredCapacity"`

	// Initial triggered time of the scheduled task. The value is in `Beijing time` (UTC+8) in the format of `YYYY-MM-DDThh:mm:ss+08:00` according to the `ISO8601` standard.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time of the scheduled task. The value is in `Beijing time` (UTC+8) in the format of `YYYY-MM-DDThh:mm:ss+08:00` according to the `ISO8601` standard. <br>This parameter and `Recurrence` need to be specified at the same time. After the end time, the scheduled task will no longer take effect.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Repeating mode of the scheduled task, which is in standard cron format. <br>This parameter and `EndTime` need to be specified at the same time.
	Recurrence *string `json:"Recurrence,omitnil,omitempty" name:"Recurrence"`
}

type ModifyScheduledActionRequest struct {
	*tchttp.BaseRequest
	
	// ID of the scheduled task to be edited
	ScheduledActionId *string `json:"ScheduledActionId,omitnil,omitempty" name:"ScheduledActionId"`

	// Scheduled task name, which can only contain letters, numbers, underscores, hyphens ("-"), and decimal points with a maximum length of 60 bytes and must be unique in an auto scaling group.
	ScheduledActionName *string `json:"ScheduledActionName,omitnil,omitempty" name:"ScheduledActionName"`

	// The maximum number of instances set for the auto scaling group when the scheduled task is triggered.
	MaxSize *uint64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`

	// The minimum number of instances set for the auto scaling group when the scheduled task is triggered.
	MinSize *uint64 `json:"MinSize,omitnil,omitempty" name:"MinSize"`

	// The desired number of instances set for the auto scaling group when the scheduled task is triggered.
	DesiredCapacity *uint64 `json:"DesiredCapacity,omitnil,omitempty" name:"DesiredCapacity"`

	// Initial triggered time of the scheduled task. The value is in `Beijing time` (UTC+8) in the format of `YYYY-MM-DDThh:mm:ss+08:00` according to the `ISO8601` standard.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time of the scheduled task. The value is in `Beijing time` (UTC+8) in the format of `YYYY-MM-DDThh:mm:ss+08:00` according to the `ISO8601` standard. <br>This parameter and `Recurrence` need to be specified at the same time. After the end time, the scheduled task will no longer take effect.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Repeating mode of the scheduled task, which is in standard cron format. <br>This parameter and `EndTime` need to be specified at the same time.
	Recurrence *string `json:"Recurrence,omitnil,omitempty" name:"Recurrence"`
}

func (r *ModifyScheduledActionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyScheduledActionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScheduledActionId")
	delete(f, "ScheduledActionName")
	delete(f, "MaxSize")
	delete(f, "MinSize")
	delete(f, "DesiredCapacity")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Recurrence")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyScheduledActionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyScheduledActionResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyScheduledActionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyScheduledActionResponseParams `json:"Response"`
}

func (r *ModifyScheduledActionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyScheduledActionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NotificationTarget struct {
	// Target type. Valid values: `CMQ_QUEUE`, `CMQ_TOPIC`, `TDMQ_CMQ_QUEUE` and `TDMQ_CMQ_TOPIC`.
	// <li> CMQ_QUEUE: Tencent Cloud message queue - queue model.</li>
	// <li> CMQ_TOPIC: Tencent Cloud message queue - topic model.</li>
	// <li> TDMQ_CMQ_QUEUE: Tencent Cloud TDMQ message queue - queue model.</li>
	// <li> TDMQ_CMQ_TOPIC: Tencent Cloud TDMQ message queue - topic model.</li>
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// Queue name. This parameter is required when `TargetType` is `CMQ_QUEUE` or `TDMQ_CMQ_QUEUE`.
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// Topic name. This parameter is required when `TargetType` is `CMQ_TOPIC` or `TDMQ_CMQ_TOPIC`.
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

type RelatedInstance struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Status of the instance in the scaling activity. Valid values:
	// `INIT`: Initializing
	// `RUNNING`: u200dProcessing u200dthe instance
	// `SUCCESSFUL`: Task succeeded on the instance
	// `FAILED`: Task failed on the instance
	InstanceStatus *string `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`
}

// Predefined struct for user
type RemoveInstancesRequestParams struct {
	// Auto scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// List of CVM instance IDs
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type RemoveInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Auto scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// List of CVM instance IDs
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *RemoveInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveInstancesResponseParams struct {
	// Scaling activity ID
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemoveInstancesResponse struct {
	*tchttp.BaseResponse
	Response *RemoveInstancesResponseParams `json:"Response"`
}

func (r *RemoveInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunAutomationServiceEnabled struct {
	// Whether to enable [TencentCloud Automation Tools](https://intl.cloud.tencent.com/document/product/1340?from_cn_redirect=1). Valid values:<br><li>`TRUE`: Enable<br><li>`FALSE`: Not enable.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type RunMonitorServiceEnabled struct {
	// Whether to enable the [Cloud Monitor](https://intl.cloud.tencent.com/document/product/248?from_cn_redirect=1) service. Value range: <br><li>TRUE: Cloud Monitor is enabled <br><li>FALSE: Cloud Monitor is disabled <br><br>Default value: TRUE. |
	// Note: This field may return null, indicating that no valid values can be obtained.
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type RunSecurityServiceEnabled struct {
	// Whether to enable the [Cloud Security](https://intl.cloud.tencent.com/document/product/296?from_cn_redirect=1) service. Value range: <br><li>TRUE: Cloud Security is enabled <br><li>FALSE: Cloud Security is disabled <br><br>Default value: TRUE.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

// Predefined struct for user
type ScaleInInstancesRequestParams struct {
	// Scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// Number of instances to be reduced
	ScaleInNumber *uint64 `json:"ScaleInNumber,omitnil,omitempty" name:"ScaleInNumber"`
}

type ScaleInInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// Number of instances to be reduced
	ScaleInNumber *uint64 `json:"ScaleInNumber,omitnil,omitempty" name:"ScaleInNumber"`
}

func (r *ScaleInInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleInInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "ScaleInNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScaleInInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleInInstancesResponseParams struct {
	// Scaling activity ID
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ScaleInInstancesResponse struct {
	*tchttp.BaseResponse
	Response *ScaleInInstancesResponseParams `json:"Response"`
}

func (r *ScaleInInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleInInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleOutInstancesRequestParams struct {
	// Scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// Number of instances to be added
	ScaleOutNumber *uint64 `json:"ScaleOutNumber,omitnil,omitempty" name:"ScaleOutNumber"`
}

type ScaleOutInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// Number of instances to be added
	ScaleOutNumber *uint64 `json:"ScaleOutNumber,omitnil,omitempty" name:"ScaleOutNumber"`
}

func (r *ScaleOutInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleOutInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "ScaleOutNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScaleOutInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleOutInstancesResponseParams struct {
	// Scaling activity ID
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ScaleOutInstancesResponse struct {
	*tchttp.BaseResponse
	Response *ScaleOutInstancesResponseParams `json:"Response"`
}

func (r *ScaleOutInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleOutInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScalingPolicy struct {
	// Auto scaling group ID.
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// Alarm trigger policy ID.
	AutoScalingPolicyId *string `json:"AutoScalingPolicyId,omitnil,omitempty" name:"AutoScalingPolicyId"`

	// Scaling policy type. Valid values:
	// - `SIMPLE`: A simple policy.
	// - `TARGET_TRACKING`: A target tracking policy.
	ScalingPolicyType *string `json:"ScalingPolicyType,omitnil,omitempty" name:"ScalingPolicyType"`

	// Alarm trigger policy name.
	ScalingPolicyName *string `json:"ScalingPolicyName,omitnil,omitempty" name:"ScalingPolicyName"`

	// The method to adjust the desired capacity after the alarm is triggered. It’s only available when `ScalingPolicyType` is `Simple`. Valid values: <br><li>`CHANGE_IN_CAPACITY`: Increase or decrease the desired capacity </li><li>`EXACT_CAPACITY`: Adjust to the specified desired capacity </li> <li>`PERCENT_CHANGE_IN_CAPACITY`: Adjust the desired capacity by percentage </li>
	AdjustmentType *string `json:"AdjustmentType,omitnil,omitempty" name:"AdjustmentType"`

	// The adjusted value of desired capacity after the alarm is triggered. This parameter is only applicable to a simple policy.
	AdjustmentValue *int64 `json:"AdjustmentValue,omitnil,omitempty" name:"AdjustmentValue"`

	// Cooldown period. This parameter is only applicable to a simple policy.
	Cooldown *uint64 `json:"Cooldown,omitnil,omitempty" name:"Cooldown"`

	// Alarm monitoring metrics of a simple policy.
	MetricAlarm *MetricAlarm `json:"MetricAlarm,omitnil,omitempty" name:"MetricAlarm"`

	// Preset monitoring item. It’s only available when `ScalingPolicyType` is `TARGET_TRACKING`. Valid values: <br><li>ASG_AVG_CPU_UTILIZATION: Average CPU utilization</li><li>ASG_AVG_LAN_TRAFFIC_OUT: Average private bandwidth out</li><li>ASG_AVG_LAN_TRAFFIC_IN: Average private bandwidth in</li><li>ASG_AVG_WAN_TRAFFIC_OUT: Average public bandwidth out</li><li>ASG_AVG_WAN_TRAFFIC_IN: Average public bandwidth in</li>
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	PredefinedMetricType *string `json:"PredefinedMetricType,omitnil,omitempty" name:"PredefinedMetricType"`

	// Target value. It’s only available when `ScalingPolicyType` is `TARGET_TRACKING`. Value ranges: <br><li>`ASG_AVG_CPU_UTILIZATION` (in %): [1, 100)</li><li>`ASG_AVG_LAN_TRAFFIC_OUT` (in Mbps): >0</li><li>`ASG_AVG_LAN_TRAFFIC_IN` (in Mbps): >0</li><li>`ASG_AVG_WAN_TRAFFIC_OUT` (in Mbps): >0</li><li>`ASG_AVG_WAN_TRAFFIC_IN` (in Mbps): >0</li>
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	TargetValue *uint64 `json:"TargetValue,omitnil,omitempty" name:"TargetValue"`

	// Instance warm-up period (in seconds). It’s only available when `ScalingPolicyType` is `TARGET_TRACKING`. Value range: 0-3600.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	EstimatedInstanceWarmup *uint64 `json:"EstimatedInstanceWarmup,omitnil,omitempty" name:"EstimatedInstanceWarmup"`

	// Whether to disable scale-in. It’s only available when `ScalingPolicyType` is `TARGET_TRACKING`. Valid values: <br><li>`true`: Scaling in is not allowed.</li><li>`false`: Allows both scale-out and scale-in</li>
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	DisableScaleIn *bool `json:"DisableScaleIn,omitnil,omitempty" name:"DisableScaleIn"`

	// List of alarm monitoring metrics. This parameter is only applicable to a target tracking policy.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	MetricAlarms []*MetricAlarm `json:"MetricAlarms,omitnil,omitempty" name:"MetricAlarms"`

	// Notification group ID, which is the set of user group IDs.
	NotificationUserGroupIds []*string `json:"NotificationUserGroupIds,omitnil,omitempty" name:"NotificationUserGroupIds"`
}

type ScheduledAction struct {
	// Scheduled task ID.
	ScheduledActionId *string `json:"ScheduledActionId,omitnil,omitempty" name:"ScheduledActionId"`

	// Scheduled task name.
	ScheduledActionName *string `json:"ScheduledActionName,omitnil,omitempty" name:"ScheduledActionName"`

	// ID of the auto scaling group where the scheduled task is located.
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// Start time of the scheduled task. The value is in `Beijing time` (UTC+8) in the format of `YYYY-MM-DDThh:mm:ss+08:00` according to the `ISO8601` standard.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Repeating mode of the scheduled task.
	Recurrence *string `json:"Recurrence,omitnil,omitempty" name:"Recurrence"`

	// End time of the scheduled task. The value is in `Beijing time` (UTC+8) in the format of `YYYY-MM-DDThh:mm:ss+08:00` according to the `ISO8601` standard.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Maximum number of instances set by the scheduled task.
	MaxSize *uint64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`

	// Desired number of instances set by the scheduled task.
	DesiredCapacity *uint64 `json:"DesiredCapacity,omitnil,omitempty" name:"DesiredCapacity"`

	// Minimum number of instances set by the scheduled task.
	MinSize *uint64 `json:"MinSize,omitnil,omitempty" name:"MinSize"`

	// Creation time of the scheduled task. The value is in `UTC time` in the format of `YYYY-MM-DDThh:mm:ssZ` according to the `ISO8601` standard.
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// Specifies how the scheduled action is executed. <br><li>`CRONTAB`: execute repeatedly <br><li>`ONCE`: execute only once
	ScheduledType *string `json:"ScheduledType,omitnil,omitempty" name:"ScheduledType"`
}

type ServiceSettings struct {
	// Enables unhealthy instance replacement. If this feature is enabled, AS will replace instances that are flagged as unhealthy by Cloud Monitor. If this parameter is not specified, the value will be False by default.
	ReplaceMonitorUnhealthy *bool `json:"ReplaceMonitorUnhealthy,omitnil,omitempty" name:"ReplaceMonitorUnhealthy"`

	// Valid values: 
	// CLASSIC_SCALING: this is the typical scaling method, which creates and terminates instances to perform scaling operations. 
	// WAKE_UP_STOPPED_SCALING: this scaling method first tries to start stopped instances. If the number of instances woken up is insufficient, the system creates new instances for scale-out. For scale-in, instances are terminated as in the typical method. You can use the StopAutoScalingInstances API to stop instances in the scaling group. Scale-out operations triggered by alarms will still create new instances.
	// Default value: CLASSIC_SCALING
	ScalingMode *string `json:"ScalingMode,omitnil,omitempty" name:"ScalingMode"`

	// Enable unhealthy instance replacement. If this feature is enabled, AS will replace instances that are found unhealthy in the CLB health check. If this parameter is not specified, the default value `False` will be used.
	ReplaceLoadBalancerUnhealthy *bool `json:"ReplaceLoadBalancerUnhealthy,omitnil,omitempty" name:"ReplaceLoadBalancerUnhealthy"`
}

// Predefined struct for user
type SetInstancesProtectionRequestParams struct {
	// Auto scaling group ID.
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// Instance ID.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Whether to enable scale-in protection for this instance
	ProtectedFromScaleIn *bool `json:"ProtectedFromScaleIn,omitnil,omitempty" name:"ProtectedFromScaleIn"`
}

type SetInstancesProtectionRequest struct {
	*tchttp.BaseRequest
	
	// Auto scaling group ID.
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// Instance ID.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Whether to enable scale-in protection for this instance
	ProtectedFromScaleIn *bool `json:"ProtectedFromScaleIn,omitnil,omitempty" name:"ProtectedFromScaleIn"`
}

func (r *SetInstancesProtectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetInstancesProtectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "InstanceIds")
	delete(f, "ProtectedFromScaleIn")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetInstancesProtectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetInstancesProtectionResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetInstancesProtectionResponse struct {
	*tchttp.BaseResponse
	Response *SetInstancesProtectionResponseParams `json:"Response"`
}

func (r *SetInstancesProtectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetInstancesProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SpotMarketOptions struct {
	// Bidding price such as "1.05"
	MaxPrice *string `json:"MaxPrice,omitnil,omitempty" name:"MaxPrice"`

	// Bid request type. Currently, only "one-time" type is supported. Default value: one-time
	// Note: This field may return null, indicating that no valid values can be obtained.
	SpotInstanceType *string `json:"SpotInstanceType,omitnil,omitempty" name:"SpotInstanceType"`
}

type SpotMixedAllocationPolicy struct {
	// The minimum number of the scaling group’s capacity that must be fulfilled by pay-as-you-go instances. It defaults to 0 if not specified. Its value cannot exceed the max capacity of the scaling group.
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	BaseCapacity *uint64 `json:"BaseCapacity,omitnil,omitempty" name:"BaseCapacity"`

	// Controls the percentage of pay-as-you-go instances for the additional capacity beyond `BaseCapacity`. Valid range: 0-100. The value 0 indicates that only spot instances are provisioned, while the value 100 indicates that only pay-as-you-go instances are provisioned. It defaults to 70 if not specified. The number of pay-as-you-go instances calculated on the percentage should be rounded up.
	// For example, if the desired capacity is 3, the `BaseCapacity` is set to 1, and the `OnDemandPercentageAboveBaseCapacity` is set to 1, the scaling group will have 2 pay-as-you-go instance (one comes from the base capacity, and the other comes from the rounded up value of the proportion), and 1 spot instance.
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	OnDemandPercentageAboveBaseCapacity *uint64 `json:"OnDemandPercentageAboveBaseCapacity,omitnil,omitempty" name:"OnDemandPercentageAboveBaseCapacity"`

	// Specifies how to assign spot instances in a mixed instance mode. Valid values: `COST_OPTIMIZED` and `CAPACITY_OPTIMIZED`; default value: `COST_OPTIMIZED`.
	// <br><li>`COST_OPTIMIZED`: the lowest cost policy. For each model in the launch configuration, AS tries to purchase it based on the lowest unit price per core in each availability zone. If the purchase failed, try the second-lowest unit price.
	// <br><li>`CAPACITY_OPTIMIZED`: the optimal capacity policy. For each model in the launch configuration, AS tries to purchase it based on the largest stock in each availability zone, minimizing the automatic repossession probability of spot instances.
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	SpotAllocationStrategy *string `json:"SpotAllocationStrategy,omitnil,omitempty" name:"SpotAllocationStrategy"`

	// Whether to replace with pay-as-you go instances. Valid values:
	// <br><li>`TRUE`: yes. After the purchase of spot instances failed due to insufficient stock and other reasons, purchase pay-as-you-go instances.
	// <br><li>`FALSE`: no. The scaling group only tries the configured model of spot instances when it needs to add spot instances.
	// 
	// Default value: `TRUE`.
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	CompensateWithBaseInstance *bool `json:"CompensateWithBaseInstance,omitnil,omitempty" name:"CompensateWithBaseInstance"`
}

// Predefined struct for user
type StartAutoScalingInstancesRequestParams struct {
	// The scaling group ID.
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// The list of the CVM instances you want to start up.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type StartAutoScalingInstancesRequest struct {
	*tchttp.BaseRequest
	
	// The scaling group ID.
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// The list of the CVM instances you want to start up.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *StartAutoScalingInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartAutoScalingInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartAutoScalingInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartAutoScalingInstancesResponseParams struct {
	// The scaling activity ID.
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartAutoScalingInstancesResponse struct {
	*tchttp.BaseResponse
	Response *StartAutoScalingInstancesResponseParams `json:"Response"`
}

func (r *StartAutoScalingInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartAutoScalingInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopAutoScalingInstancesRequestParams struct {
	// The scaling group ID.
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// The list of the CVM instances you want to shut down.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Whether the shutdown instances will be charged. Valid values:  
	// KEEP_CHARGING: keep charging after shutdown.  
	// STOP_CHARGING: stop charging after shutdown.
	// Default value: KEEP_CHARGING.
	StoppedMode *string `json:"StoppedMode,omitnil,omitempty" name:"StoppedMode"`
}

type StopAutoScalingInstancesRequest struct {
	*tchttp.BaseRequest
	
	// The scaling group ID.
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// The list of the CVM instances you want to shut down.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Whether the shutdown instances will be charged. Valid values:  
	// KEEP_CHARGING: keep charging after shutdown.  
	// STOP_CHARGING: stop charging after shutdown.
	// Default value: KEEP_CHARGING.
	StoppedMode *string `json:"StoppedMode,omitnil,omitempty" name:"StoppedMode"`
}

func (r *StopAutoScalingInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopAutoScalingInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "InstanceIds")
	delete(f, "StoppedMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopAutoScalingInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopAutoScalingInstancesResponseParams struct {
	// The scaling activity ID.
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopAutoScalingInstancesResponse struct {
	*tchttp.BaseResponse
	Response *StopAutoScalingInstancesResponseParams `json:"Response"`
}

func (r *StopAutoScalingInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopAutoScalingInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SystemDisk struct {
	// System disk type. For more information on limits of system disk types, see [Cloud Disk Types](https://intl.cloud.tencent.com/document/product/362/31636). Valid values:<br><li>`LOCAL_BASIC`: local disk <br><li>`LOCAL_SSD`: local SSD disk <br><li>`CLOUD_BASIC`: HDD cloud disk <br><li>`CLOUD_PREMIUM`: premium cloud storage<br><li>`CLOUD_SSD`: SSD cloud disk <br><br>Default value: `CLOUD_PREMIUM`.
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// System disk size in GB. Default value: 50
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`
}

type Tag struct {
	// Tag key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Tag value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// Type of the resource binded to the tag. Currently supported types include "auto-scaling-group"
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`
}

type TargetAttribute struct {
	// Port
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Weight
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`
}

// Predefined struct for user
type UpgradeLaunchConfigurationRequestParams struct {
	// Launch configuration ID.
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitnil,omitempty" name:"LaunchConfigurationId"`

	// [Image](https://intl.cloud.tencent.com/document/product/213/4940?from_cn_redirect=1) ID in the format of `img-xxx`. There are three types of images: <br/><li>Public images </li><li>Custom images </li><li>Shared images </li><br/>You can obtain the image IDs in the [CVM console](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE).</li><li>You can also use the [DescribeImages](https://intl.cloud.tencent.com/document/api/213/15715?from_cn_redirect=1) and look for `ImageId` in the response.</li>
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// List of instance models. Different instance models specify different resource specifications. Up to 5 instance models can be supported.
	InstanceTypes []*string `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// Display name of the launch configuration, which can contain letters, digits, underscores and hyphens (-), and dots. Up to of 60 bytes allowed..
	LaunchConfigurationName *string `json:"LaunchConfigurationName,omitnil,omitempty" name:"LaunchConfigurationName"`

	// Information of the instance's data disk configuration. If this parameter is not specified, no data disk is purchased by default. Up to 11 data disks can be supported.
	DataDisks []*DataDisk `json:"DataDisks,omitnil,omitempty" name:"DataDisks"`

	// Enhanced services. You can use this parameter to specify whether to enable services such as Cloud Security and Cloud Monitor. If this parameter is not specified, Cloud Monitor and Cloud Security will be enabled by default.
	EnhancedService *EnhancedService `json:"EnhancedService,omitnil,omitempty" name:"EnhancedService"`

	// Instance billing type. CVM instances are POSTPAID_BY_HOUR by default.
	// <br><li>POSTPAID_BY_HOUR: Pay-as-you-go on an hourly basis
	// <br><li>SPOTPAID: Bidding
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// Market options of the instance, such as parameters related to spot instances. This parameter is required for spot instances.
	InstanceMarketOptions *InstanceMarketOptionsRequest `json:"InstanceMarketOptions,omitnil,omitempty" name:"InstanceMarketOptions"`

	// Instance type verification policy. Value range: ALL, ANY. Default value: ANY.
	// <br><li> ALL: The verification will success only if all instance types (InstanceType) are available; otherwise, an error will be reported.
	// <br><li> ANY: The verification will success if any instance type (InstanceType) is available; otherwise, an error will be reported.
	// 
	// Common reasons why an instance type is unavailable include stock-out of the instance type or the corresponding cloud disk.
	// If a model in InstanceTypes does not exist or has been discontinued, a verification error will be reported regardless of the value of InstanceTypesCheckPolicy.
	InstanceTypesCheckPolicy *string `json:"InstanceTypesCheckPolicy,omitnil,omitempty" name:"InstanceTypesCheckPolicy"`

	// Configuration of public network bandwidth. If this parameter is not specified, 0 Mbps will be used by default.
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// Login settings of the instance. You can use this parameter to set the login method, password, and key of the instance or keep the login settings of the original image. By default, a random password will be generated and sent to you via the Message Center.
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// Project ID of the instance. Leave it blank as the default.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// The security group to which the instance belongs. This parameter can be obtained by calling the `SecurityGroupId` field in the returned value of [DescribeSecurityGroups](https://intl.cloud.tencent.com/document/api/215/15808?from_cn_redirect=1). If this parameter is not specified, no security group will be bound by default.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// System disk configuration of the instance. If this parameter is not specified, the default value will be used.
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// Base64-encoded custom data of up to 16 KB.
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`

	// List of tags. This parameter is used to bind up to 10 tags to newly added instances.
	InstanceTags []*InstanceTag `json:"InstanceTags,omitnil,omitempty" name:"InstanceTags"`

	// CAM role name, which can be obtained from the roleName field in the return value of the DescribeRoleList API.
	CamRoleName *string `json:"CamRoleName,omitnil,omitempty" name:"CamRoleName"`

	// CVM hostname settings.
	HostNameSettings *HostNameSettings `json:"HostNameSettings,omitnil,omitempty" name:"HostNameSettings"`

	// Settings of CVM instance names
	InstanceNameSettings *InstanceNameSettings `json:"InstanceNameSettings,omitnil,omitempty" name:"InstanceNameSettings"`

	// Details of the monthly subscription, including the purchase period, auto-renewal. It is required if the `InstanceChargeType` is `PREPAID`.
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// Selection policy of cloud disks. Default value: ORIGINAL. Valid values:
	// <br><li>ORIGINAL: uses the configured cloud disk type
	// <br><li>AUTOMATIC: automatically chooses an available cloud disk type
	DiskTypePolicy *string `json:"DiskTypePolicy,omitnil,omitempty" name:"DiskTypePolicy"`

	// IPv6 public network bandwidth configuration. If the IPv6 address is available in the new instance, public network bandwidth can be allocated to the IPv6 address. This parameter is invalid if `Ipv6AddressCount` of the scaling group associated with the launch configuration is 0.
	IPv6InternetAccessible *IPv6InternetAccessible `json:"IPv6InternetAccessible,omitnil,omitempty" name:"IPv6InternetAccessible"`
}

type UpgradeLaunchConfigurationRequest struct {
	*tchttp.BaseRequest
	
	// Launch configuration ID.
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitnil,omitempty" name:"LaunchConfigurationId"`

	// [Image](https://intl.cloud.tencent.com/document/product/213/4940?from_cn_redirect=1) ID in the format of `img-xxx`. There are three types of images: <br/><li>Public images </li><li>Custom images </li><li>Shared images </li><br/>You can obtain the image IDs in the [CVM console](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE).</li><li>You can also use the [DescribeImages](https://intl.cloud.tencent.com/document/api/213/15715?from_cn_redirect=1) and look for `ImageId` in the response.</li>
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// List of instance models. Different instance models specify different resource specifications. Up to 5 instance models can be supported.
	InstanceTypes []*string `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// Display name of the launch configuration, which can contain letters, digits, underscores and hyphens (-), and dots. Up to of 60 bytes allowed..
	LaunchConfigurationName *string `json:"LaunchConfigurationName,omitnil,omitempty" name:"LaunchConfigurationName"`

	// Information of the instance's data disk configuration. If this parameter is not specified, no data disk is purchased by default. Up to 11 data disks can be supported.
	DataDisks []*DataDisk `json:"DataDisks,omitnil,omitempty" name:"DataDisks"`

	// Enhanced services. You can use this parameter to specify whether to enable services such as Cloud Security and Cloud Monitor. If this parameter is not specified, Cloud Monitor and Cloud Security will be enabled by default.
	EnhancedService *EnhancedService `json:"EnhancedService,omitnil,omitempty" name:"EnhancedService"`

	// Instance billing type. CVM instances are POSTPAID_BY_HOUR by default.
	// <br><li>POSTPAID_BY_HOUR: Pay-as-you-go on an hourly basis
	// <br><li>SPOTPAID: Bidding
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// Market options of the instance, such as parameters related to spot instances. This parameter is required for spot instances.
	InstanceMarketOptions *InstanceMarketOptionsRequest `json:"InstanceMarketOptions,omitnil,omitempty" name:"InstanceMarketOptions"`

	// Instance type verification policy. Value range: ALL, ANY. Default value: ANY.
	// <br><li> ALL: The verification will success only if all instance types (InstanceType) are available; otherwise, an error will be reported.
	// <br><li> ANY: The verification will success if any instance type (InstanceType) is available; otherwise, an error will be reported.
	// 
	// Common reasons why an instance type is unavailable include stock-out of the instance type or the corresponding cloud disk.
	// If a model in InstanceTypes does not exist or has been discontinued, a verification error will be reported regardless of the value of InstanceTypesCheckPolicy.
	InstanceTypesCheckPolicy *string `json:"InstanceTypesCheckPolicy,omitnil,omitempty" name:"InstanceTypesCheckPolicy"`

	// Configuration of public network bandwidth. If this parameter is not specified, 0 Mbps will be used by default.
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// Login settings of the instance. You can use this parameter to set the login method, password, and key of the instance or keep the login settings of the original image. By default, a random password will be generated and sent to you via the Message Center.
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// Project ID of the instance. Leave it blank as the default.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// The security group to which the instance belongs. This parameter can be obtained by calling the `SecurityGroupId` field in the returned value of [DescribeSecurityGroups](https://intl.cloud.tencent.com/document/api/215/15808?from_cn_redirect=1). If this parameter is not specified, no security group will be bound by default.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// System disk configuration of the instance. If this parameter is not specified, the default value will be used.
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// Base64-encoded custom data of up to 16 KB.
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`

	// List of tags. This parameter is used to bind up to 10 tags to newly added instances.
	InstanceTags []*InstanceTag `json:"InstanceTags,omitnil,omitempty" name:"InstanceTags"`

	// CAM role name, which can be obtained from the roleName field in the return value of the DescribeRoleList API.
	CamRoleName *string `json:"CamRoleName,omitnil,omitempty" name:"CamRoleName"`

	// CVM hostname settings.
	HostNameSettings *HostNameSettings `json:"HostNameSettings,omitnil,omitempty" name:"HostNameSettings"`

	// Settings of CVM instance names
	InstanceNameSettings *InstanceNameSettings `json:"InstanceNameSettings,omitnil,omitempty" name:"InstanceNameSettings"`

	// Details of the monthly subscription, including the purchase period, auto-renewal. It is required if the `InstanceChargeType` is `PREPAID`.
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// Selection policy of cloud disks. Default value: ORIGINAL. Valid values:
	// <br><li>ORIGINAL: uses the configured cloud disk type
	// <br><li>AUTOMATIC: automatically chooses an available cloud disk type
	DiskTypePolicy *string `json:"DiskTypePolicy,omitnil,omitempty" name:"DiskTypePolicy"`

	// IPv6 public network bandwidth configuration. If the IPv6 address is available in the new instance, public network bandwidth can be allocated to the IPv6 address. This parameter is invalid if `Ipv6AddressCount` of the scaling group associated with the launch configuration is 0.
	IPv6InternetAccessible *IPv6InternetAccessible `json:"IPv6InternetAccessible,omitnil,omitempty" name:"IPv6InternetAccessible"`
}

func (r *UpgradeLaunchConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeLaunchConfigurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LaunchConfigurationId")
	delete(f, "ImageId")
	delete(f, "InstanceTypes")
	delete(f, "LaunchConfigurationName")
	delete(f, "DataDisks")
	delete(f, "EnhancedService")
	delete(f, "InstanceChargeType")
	delete(f, "InstanceMarketOptions")
	delete(f, "InstanceTypesCheckPolicy")
	delete(f, "InternetAccessible")
	delete(f, "LoginSettings")
	delete(f, "ProjectId")
	delete(f, "SecurityGroupIds")
	delete(f, "SystemDisk")
	delete(f, "UserData")
	delete(f, "InstanceTags")
	delete(f, "CamRoleName")
	delete(f, "HostNameSettings")
	delete(f, "InstanceNameSettings")
	delete(f, "InstanceChargePrepaid")
	delete(f, "DiskTypePolicy")
	delete(f, "IPv6InternetAccessible")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeLaunchConfigurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeLaunchConfigurationResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeLaunchConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeLaunchConfigurationResponseParams `json:"Response"`
}

func (r *UpgradeLaunchConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeLaunchConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeLifecycleHookRequestParams struct {
	// Lifecycle hook ID
	LifecycleHookId *string `json:"LifecycleHookId,omitnil,omitempty" name:"LifecycleHookId"`

	// Lifecycle hook name
	LifecycleHookName *string `json:"LifecycleHookName,omitnil,omitempty" name:"LifecycleHookName"`

	// Scenario for the lifecycle hook. Value range: "INSTANCE_LAUNCHING", "INSTANCE_TERMINATING"
	LifecycleTransition *string `json:"LifecycleTransition,omitnil,omitempty" name:"LifecycleTransition"`

	// Defines the action to be taken by the auto scaling group upon lifecycle hook timeout. Value range: "CONTINUE", "ABANDON". Default value: "CONTINUE"
	DefaultResult *string `json:"DefaultResult,omitnil,omitempty" name:"DefaultResult"`

	// The maximum length of time (in seconds) that can elapse before the lifecycle hook times out. Value range: 30-7200. Default value: 300
	HeartbeatTimeout *int64 `json:"HeartbeatTimeout,omitnil,omitempty" name:"HeartbeatTimeout"`

	// Additional information of a notification that Auto Scaling sends to targets. This parameter is set when you configure a notification (default value: "").
	NotificationMetadata *string `json:"NotificationMetadata,omitnil,omitempty" name:"NotificationMetadata"`

	// Notification result. `NotificationTarget` and `LifecycleCommand` cannot be specified at the same time.
	NotificationTarget *NotificationTarget `json:"NotificationTarget,omitnil,omitempty" name:"NotificationTarget"`

	// The scenario where the lifecycle hook is applied. `EXTENSION`: the lifecycle hook will be triggered when AttachInstances, DetachInstances or RemoveInstaces is called. `NORMAL`: the lifecycle hook is not triggered by the above APIs. 
	LifecycleTransitionType *string `json:"LifecycleTransitionType,omitnil,omitempty" name:"LifecycleTransitionType"`

	// Remote command execution object. `NotificationTarget` and `LifecycleCommand` cannot be specified at the same time.
	LifecycleCommand *LifecycleCommand `json:"LifecycleCommand,omitnil,omitempty" name:"LifecycleCommand"`
}

type UpgradeLifecycleHookRequest struct {
	*tchttp.BaseRequest
	
	// Lifecycle hook ID
	LifecycleHookId *string `json:"LifecycleHookId,omitnil,omitempty" name:"LifecycleHookId"`

	// Lifecycle hook name
	LifecycleHookName *string `json:"LifecycleHookName,omitnil,omitempty" name:"LifecycleHookName"`

	// Scenario for the lifecycle hook. Value range: "INSTANCE_LAUNCHING", "INSTANCE_TERMINATING"
	LifecycleTransition *string `json:"LifecycleTransition,omitnil,omitempty" name:"LifecycleTransition"`

	// Defines the action to be taken by the auto scaling group upon lifecycle hook timeout. Value range: "CONTINUE", "ABANDON". Default value: "CONTINUE"
	DefaultResult *string `json:"DefaultResult,omitnil,omitempty" name:"DefaultResult"`

	// The maximum length of time (in seconds) that can elapse before the lifecycle hook times out. Value range: 30-7200. Default value: 300
	HeartbeatTimeout *int64 `json:"HeartbeatTimeout,omitnil,omitempty" name:"HeartbeatTimeout"`

	// Additional information of a notification that Auto Scaling sends to targets. This parameter is set when you configure a notification (default value: "").
	NotificationMetadata *string `json:"NotificationMetadata,omitnil,omitempty" name:"NotificationMetadata"`

	// Notification result. `NotificationTarget` and `LifecycleCommand` cannot be specified at the same time.
	NotificationTarget *NotificationTarget `json:"NotificationTarget,omitnil,omitempty" name:"NotificationTarget"`

	// The scenario where the lifecycle hook is applied. `EXTENSION`: the lifecycle hook will be triggered when AttachInstances, DetachInstances or RemoveInstaces is called. `NORMAL`: the lifecycle hook is not triggered by the above APIs. 
	LifecycleTransitionType *string `json:"LifecycleTransitionType,omitnil,omitempty" name:"LifecycleTransitionType"`

	// Remote command execution object. `NotificationTarget` and `LifecycleCommand` cannot be specified at the same time.
	LifecycleCommand *LifecycleCommand `json:"LifecycleCommand,omitnil,omitempty" name:"LifecycleCommand"`
}

func (r *UpgradeLifecycleHookRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeLifecycleHookRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LifecycleHookId")
	delete(f, "LifecycleHookName")
	delete(f, "LifecycleTransition")
	delete(f, "DefaultResult")
	delete(f, "HeartbeatTimeout")
	delete(f, "NotificationMetadata")
	delete(f, "NotificationTarget")
	delete(f, "LifecycleTransitionType")
	delete(f, "LifecycleCommand")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeLifecycleHookRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeLifecycleHookResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeLifecycleHookResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeLifecycleHookResponseParams `json:"Response"`
}

func (r *UpgradeLifecycleHookResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeLifecycleHookResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}