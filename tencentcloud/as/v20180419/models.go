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

	// Scaling activity type. Valid values:
	// <li>SCALE_OUT: Scale out an instance.</li>
	// <li>SCALE_IN: Scale in an instance.</li>
	// <li>ATTACH_INSTANCES: Add an instance.</li>
	// <li>REMOVE_INSTANCES: Terminate an instance.</li>
	// <li>DETACH_INSTANCES: Remove an instance.</li>
	// <li>TERMINATE_INSTANCES_UNEXPECTEDLY: Terminate an instance in the CVM console.</li>
	// <li>REPLACE_UNHEALTHY_INSTANCE: Replace an unhealthy instance.</li>
	// <li>START_INSTANCES: Start an instance.</li>
	// <li>STOP_INSTANCES: Stop an instance.</li>
	// <li>INVOKE_COMMAND: Execute a command.</li>
	ActivityType *string `json:"ActivityType,omitnil,omitempty" name:"ActivityType"`

	// Scaling activity status. Valid values:
	// <li>INIT: initializing.</li>
	// <li>RUNNING: running.</li>
	// <li>SUCCESSFUL: successful.</li>
	// <li>PARTIALLY_SUCCESSFUL: partially successful.</li>
	// <li>FAILED: failed.</li>
	// <li>CANCELLED: canceled.</li>
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

	// Status of the instance in the scaling activity. Valid values: <br><li>INIT: initializing;</li> <li>RUNNING: instance in operation;</li> <li>SUCCESSFUL: activity successful;</li> <li>FAILED: activity failed.
	InstanceStatus *string `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`
}

type Advice struct {
	// Problem Description
	Problem *string `json:"Problem,omitnil,omitempty" name:"Problem"`

	// Problem Details
	Detail *string `json:"Detail,omitnil,omitempty" name:"Detail"`

	// Recommended resolutions
	Solution *string `json:"Solution,omitnil,omitempty" name:"Solution"`

	// Scaling suggestion warning level. Valid values:
	// <li>WARNING: warning.</li>
	// <li>CRITICAL: critical.</li>
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// Scaling group warning level. Valid values:
	// <li>NORMAL: normal.</li>
	// <li>WARNING: warning.</li>
	// <li>CRITICAL: critical.</li>
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// A collection of suggestions for scaling group configurations.
	Advices []*Advice `json:"Advices,omitnil,omitempty" name:"Advices"`
}

type AutoScalingGroup struct {
	// Auto scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// Auto scaling group name
	AutoScalingGroupName *string `json:"AutoScalingGroupName,omitnil,omitempty" name:"AutoScalingGroupName"`

	// Current status of the scaling group. Valid values:
	// <li>NORMAL: The scaling group is normal.</li>
	// <li>CVM_ABNORMAL: The launch configuration is abnormal.</li>
	// <li>LB_ABNORMAL: The CLB is abnormal.</li>
	// <li>LB_LISTENER_ABNORMAL: The CLB listener is abnormal.</li>
	// <li>LB_LOCATION_ABNORMAL: The forwarding configuration of the CLB listener is abnormal.</li>
	// <li>VPC_ABNORMAL: The VPC is abnormal.</li>
	// <li>SUBNET_ABNORMAL: The VPC subnet is abnormal.</li>
	// <li>INSUFFICIENT_BALANCE: The balance is insufficient.</li>
	// <li>LB_BACKEND_REGION_NOT_MATCH: The region of the CLB instance backend does not match that of the AS service.</li>
	// <li>LB_BACKEND_VPC_NOT_MATCH: The VPC of the CLB instance does not match that of the scaling group.</li>
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

	// Multi-AZ/subnet policy.
	// <li>PRIORITY: The instances are attempted to be created taking the order of the AZ/subnet list as the priority. If the highest-priority AZ/subnet can create instances successfully, instances can always be created in that AZ/subnet.</li>
	// <li>EQUALITY: Select the AZ/subnet with the least number of instances for scale-out. In this way, each AZ/subnet has an opportunity for scale-out. Instances produced from multiple scale-out operations will be distributed to multiple AZs/subnets.</li>
	MultiZoneSubnetPolicy *string `json:"MultiZoneSubnetPolicy,omitnil,omitempty" name:"MultiZoneSubnetPolicy"`

	// Scaling group instance health check type, whose valid values include:
	// <li>CVM: Determine whether an instance is unhealthy based on its network status. An unhealthy network status is indicated by an event where instances become unreachable via PING. Detailed criteria of judgment can be found in [Instance Health Check](https://intl.cloud.tencent.com/document/product/377/8553?from_cn_redirect=1).</li>
	// <li>CLB: Determine whether an instance is unhealthy based on the health check status of CLB. For principles behind CLB health checks, see [Health Check](https://intl.cloud.tencent.com/document/product/214/6097?from_cn_redirect=1).</li>
	HealthCheckType *string `json:"HealthCheckType,omitnil,omitempty" name:"HealthCheckType"`

	// Grace period of the CLB health check
	LoadBalancerHealthCheckGracePeriod *uint64 `json:"LoadBalancerHealthCheckGracePeriod,omitnil,omitempty" name:"LoadBalancerHealthCheckGracePeriod"`

	// Instance assignment policy, whose valid values include LAUNCH_CONFIGURATION and SPOT_MIXED.
	// <li>LAUNCH_CONFIGURATION: Represent the traditional mode of assigning instances according to the launch configuration.</li>
	// <li>SPOT_MIXED: Represent the spot mixed mode. Currently, this mode is supported only when the launch configuration is set to the pay-as-you-go billing mode. In the mixed mode, the scaling group will scale out pay-as-you-go models or spot models based on the predefined settings. When the mixed mode is used, the billing type of the associated launch configuration cannot be modified.</li>
	InstanceAllocationPolicy *string `json:"InstanceAllocationPolicy,omitnil,omitempty" name:"InstanceAllocationPolicy"`

	// Specifies how to assign pay-as-you-go instances and spot instances.
	// A valid value will be returned only when `InstanceAllocationPolicy` is set to `SPOT_MIXED`.
	SpotMixedAllocationPolicy *SpotMixedAllocationPolicy `json:"SpotMixedAllocationPolicy,omitnil,omitempty" name:"SpotMixedAllocationPolicy"`

	// Capacity rebalancing feature, which is applicable only to spot instances within the scaling group. Valid values:
	// <li>TRUE: Enable this feature. When spot instances in the scaling group are about to be automatically recycled by the spot instance service, AS proactively initiates the termination process of the spot instances. If there is a configured scale-in hook, it will be triggered before termination. After the termination process starts, AS asynchronously initiates the scale-out to reach the expected number of instances.</li>
	// <li>FALSE: Disable this feature. AS waits for the spot instance to be terminated before scaling out to reach the number of instances expected by the scaling group.</li>
	CapacityRebalance *bool `json:"CapacityRebalance,omitnil,omitempty" name:"CapacityRebalance"`

	// Instance name sequencing settings.
	// Note: This field may return null, indicating that no valid value can be obtained.
	InstanceNameIndexSettings *InstanceNameIndexSettings `json:"InstanceNameIndexSettings,omitnil,omitempty" name:"InstanceNameIndexSettings"`
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
type CancelInstanceRefreshRequestParams struct {
	// Scaling group ID.
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// Refresh activity ID.
	RefreshActivityId *string `json:"RefreshActivityId,omitnil,omitempty" name:"RefreshActivityId"`
}

type CancelInstanceRefreshRequest struct {
	*tchttp.BaseRequest
	
	// Scaling group ID.
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// Refresh activity ID.
	RefreshActivityId *string `json:"RefreshActivityId,omitnil,omitempty" name:"RefreshActivityId"`
}

func (r *CancelInstanceRefreshRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelInstanceRefreshRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "RefreshActivityId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelInstanceRefreshRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelInstanceRefreshResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CancelInstanceRefreshResponse struct {
	*tchttp.BaseResponse
	Response *CancelInstanceRefreshResponseParams `json:"Response"`
}

func (r *CancelInstanceRefreshResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelInstanceRefreshResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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
	// Setting it to `true` will clear the instance name settings, which means that CVM newly created on this launch configuration will be named in the as-{{AutoScalingGroupName}} format.
	ClearInstanceNameSettings *bool `json:"ClearInstanceNameSettings,omitnil,omitempty" name:"ClearInstanceNameSettings"`

	// Whether to clear placement group information. This parameter is optional. Default value: `false`.
	// `True` means clearing placement group information. After that, no placement groups are specified for CVMs created based on the information.
	ClearDisasterRecoverGroupIds *bool `json:"ClearDisasterRecoverGroupIds,omitnil,omitempty" name:"ClearDisasterRecoverGroupIds"`

	// Whether to clear the instance tag list. This parameter is optional, and its default value is false.
	// If true is filled in, it indicates that the instance tag list should be cleared. After the list is cleared, the CVMs created based on this will not be bound to the tags in the list.
	ClearInstanceTags *bool `json:"ClearInstanceTags,omitnil,omitempty" name:"ClearInstanceTags"`

	// Whether to clear metadata, optional, defaults to false. Setting it to true will clear metadata, the CVMs created based on this will not be associated with custom metadata.
	ClearMetadata *bool `json:"ClearMetadata,omitnil,omitempty" name:"ClearMetadata"`
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
	// Setting it to `true` will clear the instance name settings, which means that CVM newly created on this launch configuration will be named in the as-{{AutoScalingGroupName}} format.
	ClearInstanceNameSettings *bool `json:"ClearInstanceNameSettings,omitnil,omitempty" name:"ClearInstanceNameSettings"`

	// Whether to clear placement group information. This parameter is optional. Default value: `false`.
	// `True` means clearing placement group information. After that, no placement groups are specified for CVMs created based on the information.
	ClearDisasterRecoverGroupIds *bool `json:"ClearDisasterRecoverGroupIds,omitnil,omitempty" name:"ClearDisasterRecoverGroupIds"`

	// Whether to clear the instance tag list. This parameter is optional, and its default value is false.
	// If true is filled in, it indicates that the instance tag list should be cleared. After the list is cleared, the CVMs created based on this will not be bound to the tags in the list.
	ClearInstanceTags *bool `json:"ClearInstanceTags,omitnil,omitempty" name:"ClearInstanceTags"`

	// Whether to clear metadata, optional, defaults to false. Setting it to true will clear metadata, the CVMs created based on this will not be associated with custom metadata.
	ClearMetadata *bool `json:"ClearMetadata,omitnil,omitempty" name:"ClearMetadata"`
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
	delete(f, "ClearInstanceTags")
	delete(f, "ClearMetadata")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ClearLaunchConfigurationAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ClearLaunchConfigurationAttributesResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// Termination policy, whose maximum length is currently 1. Valid values: OLDEST_INSTANCE and NEWEST_INSTANCE. Default value: OLDEST_INSTANCE.
	// <li>OLDEST_INSTANCE: Terminate the oldest instance in the scaling group first.</li>
	// <li>NEWEST_INSTANCE: Terminate the newest instance in the scaling group first.</li>
	TerminationPolicies []*string `json:"TerminationPolicies,omitnil,omitempty" name:"TerminationPolicies"`

	// List of availability zones. An availability zone must be specified in the basic network scenario. If multiple availability zones are entered, their priority will be determined by the order in which they are entered, and they will be tried one by one until instances can be successfully created.
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// Retry policy. Valid values: IMMEDIATE_RETRY, INCREMENTAL_INTERVALS, and NO_RETRY. Default value: IMMEDIATE_RETRY. A partially successful scaling activity is considered a failed activity.
	// <li>IMMEDIATE_RETRY: Immediately retry or quickly retry in a short period. There will be no retry anymore after a certain number of consecutive failures (5).</li>
	// <li>INCREMENTAL_INTERVALS: Retry at an incremental interval. As the number of continuous failures increase, the retry interval gradually increases. The interval for the first 10 retries is the same as the immediate retry mode, and that for the subsequent retries gradually increases to 10 minutes, 30 minutes, 60 minutes, or 1 day.</li>
	// <li>NO_RETRY: There will be no retry until another user call or alarm information is received.</li>  
	RetryPolicy *string `json:"RetryPolicy,omitnil,omitempty" name:"RetryPolicy"`

	// AZ verification policy. Valid values: ALL and ANY. Default value: ANY.
	// <li>ALL: Verification passes if all AZs or subnets are available; otherwise, a verification error will be reported.<li>
	// <li>ANY: Verification passes if any AZ or subnet is available; otherwise, a verification error will be reported.</li>
	// 
	// Common reasons for unavailable AZs or subnets include the CVM InstanceType in the AZ being sold out, the CBS cloud disk in the AZ being sold out, insufficient quota in the AZ, and insufficient IP addresses in the subnet.
	// If there is no AZ or subnet in Zones/SubnetIds, a verification error will be reported regardless of the values of ZonesCheckPolicy.
	ZonesCheckPolicy *string `json:"ZonesCheckPolicy,omitnil,omitempty" name:"ZonesCheckPolicy"`

	// List of tag descriptions. In this parameter, you can specify the tags to be bound with a scaling group as well as corresponding resource instances. Each scaling group can have up to 30 tags.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Service settings such as unhealthy instance replacement.
	ServiceSettings *ServiceSettings `json:"ServiceSettings,omitnil,omitempty" name:"ServiceSettings"`

	// The number of IPv6 addresses that an instance has. Valid values: 0 and 1. Default value: 0.
	Ipv6AddressCount *int64 `json:"Ipv6AddressCount,omitnil,omitempty" name:"Ipv6AddressCount"`

	// Multi-AZ/multi-subnet policy, whose valid values include PRIORITY and EQUALITY, with the default value being PRIORITY.
	// <li>PRIORITY: The instances are attempted to be created taking the order of the AZ/subnet list as the priority. If instances can be successfully created in the highest-priority AZ/subnet, they will always be created in that AZ/subnet.</li>
	// <li>EQUALITY: The instances added through scale-out will be distributed across multiple AZs/subnets to ensure a relatively balanced number of instances in each AZ/subnet after scaling out.</li>
	// 
	// Points to consider regarding this policy:
	// <li>When the scaling group is based on a classic network, this policy applies to the multi-AZ; when the scaling group is based on a VPC network, this policy applies to the multi-subnet, in this case, the AZs are no longer considered. For example, if there are four subnets labeled A, B, C, and D, where A, B, and C are in AZ 1 and D is in AZ 2, the subnets A, B, C, and D are considered for sorting without regard to AZs 1 and 2.</li>
	// <li>This policy applies to the multi-AZ/multi-subnet and not to the InstanceTypes parameter of the launch configuration, which is selected according to the priority policy.</li>
	// <li>When instances are created according to the PRIORITY policy, ensure the policy for multiple models first, followed by the policy for the multi-AZ/subnet. For example, with models A and B and subnets 1, 2, and 3, attempts will be made in the order of A1, A2, A3, B1, B2, and B3. If A1 is sold out, A2 will be attempted (instead of B1).</li>
	MultiZoneSubnetPolicy *string `json:"MultiZoneSubnetPolicy,omitnil,omitempty" name:"MultiZoneSubnetPolicy"`

	// Health check type for scaling group instances. Valid values:
	// <li>CVM: Determine whether an instance is unhealthy based on its network status. An unhealthy network status is indicated by an event where instances become unreachable via PING. For detailed criteria of judgment, see [Instance Health Check](https://www.tencentcloud.com/document/product/377/8553?lang=en&pg=).</li>
	// <li>CLB: Determine whether an instance is unhealthy based on the health check status of CLB. For principles behind CLB health checks, see [Health Check Overview](https://www.tencentcloud.com/document/product/214/6097?from_search=1&lang=en&pg=).</li>
	// If CLB is selected, the scaling group will check both the instance's network status and the CLB's health check status. If the instance's network status is unhealthy, the instance will be marked as UNHEALTHY. If the CLB's health check status is abnormal, the instance will be marked as CLB_UNHEALTHY. If both of them are abnormal, the instance will be marked as UNHEALTHY|CLB_UNHEALTHY. Default value: CLB.
	HealthCheckType *string `json:"HealthCheckType,omitnil,omitempty" name:"HealthCheckType"`

	// Grace period of the CLB health check during which the `IN_SERVICE` instances added will not be marked as `CLB_UNHEALTHY`.<br>Valid range: 0-7200, in seconds. Default value: `0`.
	LoadBalancerHealthCheckGracePeriod *uint64 `json:"LoadBalancerHealthCheckGracePeriod,omitnil,omitempty" name:"LoadBalancerHealthCheckGracePeriod"`

	// Instance assignment policy. Valid values: LAUNCH_CONFIGURATION and SPOT_MIXED. Default value: LAUNCH_CONFIGURATION.
	// <li>LAUNCH_CONFIGURATION: Represent the traditional mode of assigning instances according to the launch configuration.</li>
	// <li>SPOT_MIXED: Represent the spot mixed mode. Currently, this mode is supported only when the launch configuration is set to the pay-as-you-go billing mode. In the mixed mode, the scaling group will scale out pay-as-you-go models or spot models based on the predefined settings. When the mixed mode is used, the billing type of the associated launch configuration cannot be modified.</li>
	InstanceAllocationPolicy *string `json:"InstanceAllocationPolicy,omitnil,omitempty" name:"InstanceAllocationPolicy"`

	// Specifies how to assign pay-as-you-go instances and spot instances.
	// This parameter is valid only when `InstanceAllocationPolicy ` is set to `SPOT_MIXED`.
	SpotMixedAllocationPolicy *SpotMixedAllocationPolicy `json:"SpotMixedAllocationPolicy,omitnil,omitempty" name:"SpotMixedAllocationPolicy"`

	// Capacity rebalancing feature, which is applicable only to spot instances within the scaling group. Valid values:
	// <li>TRUE: Enable this feature. When spot instances in the scaling group are about to be automatically recycled by the spot instance service, AS proactively initiates the termination process of the spot instances. If there is a configured scale-in hook, it will be triggered before termination. After the termination process starts, AS asynchronously initiates the scale-out to reach the expected number of instances.</li>
	// <li>FALSE: Disable this feature. AS waits for the spot instance to be terminated before scaling out to reach the number of instances expected by the scaling group.</li>
	// 
	// Default value: FALSE.
	CapacityRebalance *bool `json:"CapacityRebalance,omitnil,omitempty" name:"CapacityRebalance"`

	// Instance name sequencing settings. If this parameter is not specified, the default is not enabled. When enabled, an incremental numeric sequence will be appended to the names of instances automatically created within the scaling group.
	InstanceNameIndexSettings *InstanceNameIndexSettings `json:"InstanceNameIndexSettings,omitnil,omitempty" name:"InstanceNameIndexSettings"`
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

	// Termination policy, whose maximum length is currently 1. Valid values: OLDEST_INSTANCE and NEWEST_INSTANCE. Default value: OLDEST_INSTANCE.
	// <li>OLDEST_INSTANCE: Terminate the oldest instance in the scaling group first.</li>
	// <li>NEWEST_INSTANCE: Terminate the newest instance in the scaling group first.</li>
	TerminationPolicies []*string `json:"TerminationPolicies,omitnil,omitempty" name:"TerminationPolicies"`

	// List of availability zones. An availability zone must be specified in the basic network scenario. If multiple availability zones are entered, their priority will be determined by the order in which they are entered, and they will be tried one by one until instances can be successfully created.
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// Retry policy. Valid values: IMMEDIATE_RETRY, INCREMENTAL_INTERVALS, and NO_RETRY. Default value: IMMEDIATE_RETRY. A partially successful scaling activity is considered a failed activity.
	// <li>IMMEDIATE_RETRY: Immediately retry or quickly retry in a short period. There will be no retry anymore after a certain number of consecutive failures (5).</li>
	// <li>INCREMENTAL_INTERVALS: Retry at an incremental interval. As the number of continuous failures increase, the retry interval gradually increases. The interval for the first 10 retries is the same as the immediate retry mode, and that for the subsequent retries gradually increases to 10 minutes, 30 minutes, 60 minutes, or 1 day.</li>
	// <li>NO_RETRY: There will be no retry until another user call or alarm information is received.</li>  
	RetryPolicy *string `json:"RetryPolicy,omitnil,omitempty" name:"RetryPolicy"`

	// AZ verification policy. Valid values: ALL and ANY. Default value: ANY.
	// <li>ALL: Verification passes if all AZs or subnets are available; otherwise, a verification error will be reported.<li>
	// <li>ANY: Verification passes if any AZ or subnet is available; otherwise, a verification error will be reported.</li>
	// 
	// Common reasons for unavailable AZs or subnets include the CVM InstanceType in the AZ being sold out, the CBS cloud disk in the AZ being sold out, insufficient quota in the AZ, and insufficient IP addresses in the subnet.
	// If there is no AZ or subnet in Zones/SubnetIds, a verification error will be reported regardless of the values of ZonesCheckPolicy.
	ZonesCheckPolicy *string `json:"ZonesCheckPolicy,omitnil,omitempty" name:"ZonesCheckPolicy"`

	// List of tag descriptions. In this parameter, you can specify the tags to be bound with a scaling group as well as corresponding resource instances. Each scaling group can have up to 30 tags.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Service settings such as unhealthy instance replacement.
	ServiceSettings *ServiceSettings `json:"ServiceSettings,omitnil,omitempty" name:"ServiceSettings"`

	// The number of IPv6 addresses that an instance has. Valid values: 0 and 1. Default value: 0.
	Ipv6AddressCount *int64 `json:"Ipv6AddressCount,omitnil,omitempty" name:"Ipv6AddressCount"`

	// Multi-AZ/multi-subnet policy, whose valid values include PRIORITY and EQUALITY, with the default value being PRIORITY.
	// <li>PRIORITY: The instances are attempted to be created taking the order of the AZ/subnet list as the priority. If instances can be successfully created in the highest-priority AZ/subnet, they will always be created in that AZ/subnet.</li>
	// <li>EQUALITY: The instances added through scale-out will be distributed across multiple AZs/subnets to ensure a relatively balanced number of instances in each AZ/subnet after scaling out.</li>
	// 
	// Points to consider regarding this policy:
	// <li>When the scaling group is based on a classic network, this policy applies to the multi-AZ; when the scaling group is based on a VPC network, this policy applies to the multi-subnet, in this case, the AZs are no longer considered. For example, if there are four subnets labeled A, B, C, and D, where A, B, and C are in AZ 1 and D is in AZ 2, the subnets A, B, C, and D are considered for sorting without regard to AZs 1 and 2.</li>
	// <li>This policy applies to the multi-AZ/multi-subnet and not to the InstanceTypes parameter of the launch configuration, which is selected according to the priority policy.</li>
	// <li>When instances are created according to the PRIORITY policy, ensure the policy for multiple models first, followed by the policy for the multi-AZ/subnet. For example, with models A and B and subnets 1, 2, and 3, attempts will be made in the order of A1, A2, A3, B1, B2, and B3. If A1 is sold out, A2 will be attempted (instead of B1).</li>
	MultiZoneSubnetPolicy *string `json:"MultiZoneSubnetPolicy,omitnil,omitempty" name:"MultiZoneSubnetPolicy"`

	// Health check type for scaling group instances. Valid values:
	// <li>CVM: Determine whether an instance is unhealthy based on its network status. An unhealthy network status is indicated by an event where instances become unreachable via PING. For detailed criteria of judgment, see [Instance Health Check](https://www.tencentcloud.com/document/product/377/8553?lang=en&pg=).</li>
	// <li>CLB: Determine whether an instance is unhealthy based on the health check status of CLB. For principles behind CLB health checks, see [Health Check Overview](https://www.tencentcloud.com/document/product/214/6097?from_search=1&lang=en&pg=).</li>
	// If CLB is selected, the scaling group will check both the instance's network status and the CLB's health check status. If the instance's network status is unhealthy, the instance will be marked as UNHEALTHY. If the CLB's health check status is abnormal, the instance will be marked as CLB_UNHEALTHY. If both of them are abnormal, the instance will be marked as UNHEALTHY|CLB_UNHEALTHY. Default value: CLB.
	HealthCheckType *string `json:"HealthCheckType,omitnil,omitempty" name:"HealthCheckType"`

	// Grace period of the CLB health check during which the `IN_SERVICE` instances added will not be marked as `CLB_UNHEALTHY`.<br>Valid range: 0-7200, in seconds. Default value: `0`.
	LoadBalancerHealthCheckGracePeriod *uint64 `json:"LoadBalancerHealthCheckGracePeriod,omitnil,omitempty" name:"LoadBalancerHealthCheckGracePeriod"`

	// Instance assignment policy. Valid values: LAUNCH_CONFIGURATION and SPOT_MIXED. Default value: LAUNCH_CONFIGURATION.
	// <li>LAUNCH_CONFIGURATION: Represent the traditional mode of assigning instances according to the launch configuration.</li>
	// <li>SPOT_MIXED: Represent the spot mixed mode. Currently, this mode is supported only when the launch configuration is set to the pay-as-you-go billing mode. In the mixed mode, the scaling group will scale out pay-as-you-go models or spot models based on the predefined settings. When the mixed mode is used, the billing type of the associated launch configuration cannot be modified.</li>
	InstanceAllocationPolicy *string `json:"InstanceAllocationPolicy,omitnil,omitempty" name:"InstanceAllocationPolicy"`

	// Specifies how to assign pay-as-you-go instances and spot instances.
	// This parameter is valid only when `InstanceAllocationPolicy ` is set to `SPOT_MIXED`.
	SpotMixedAllocationPolicy *SpotMixedAllocationPolicy `json:"SpotMixedAllocationPolicy,omitnil,omitempty" name:"SpotMixedAllocationPolicy"`

	// Capacity rebalancing feature, which is applicable only to spot instances within the scaling group. Valid values:
	// <li>TRUE: Enable this feature. When spot instances in the scaling group are about to be automatically recycled by the spot instance service, AS proactively initiates the termination process of the spot instances. If there is a configured scale-in hook, it will be triggered before termination. After the termination process starts, AS asynchronously initiates the scale-out to reach the expected number of instances.</li>
	// <li>FALSE: Disable this feature. AS waits for the spot instance to be terminated before scaling out to reach the number of instances expected by the scaling group.</li>
	// 
	// Default value: FALSE.
	CapacityRebalance *bool `json:"CapacityRebalance,omitnil,omitempty" name:"CapacityRebalance"`

	// Instance name sequencing settings. If this parameter is not specified, the default is not enabled. When enabled, an incremental numeric sequence will be appended to the names of instances automatically created within the scaling group.
	InstanceNameIndexSettings *InstanceNameIndexSettings `json:"InstanceNameIndexSettings,omitnil,omitempty" name:"InstanceNameIndexSettings"`
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
	delete(f, "InstanceNameIndexSettings")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAutoScalingGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAutoScalingGroupResponseParams struct {
	// Auto scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// Project ID of the launch configuration. The default project is used if it is left blank.
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
	// <li>POSTPAID_BY_HOUR: pay-as-you-go hourly</li>
	// <li>SPOTPAID: spot instance</li>
	// <li> CDCPAID: dedicated cluster</li>
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// Market options of the instance, such as parameters related to spot instances. This parameter is required for spot instances.
	InstanceMarketOptions *InstanceMarketOptionsRequest `json:"InstanceMarketOptions,omitnil,omitempty" name:"InstanceMarketOptions"`

	// List of instance models. Different instance models specify different resource specifications. Up to 10 instance models can be supported.
	// `InstanceType` and `InstanceTypes` are mutually exclusive, and one and only one of them must be entered.
	InstanceTypes []*string `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// CAM role name. This parameter can be obtained from the `roleName` field returned by DescribeRoleList API.
	CamRoleName *string `json:"CamRoleName,omitnil,omitempty" name:"CamRoleName"`

	// InstanceType verification policy, whose valid values include ALL and ANY, with the default value being ANY.
	// <li>ALL: Verification passes if all InstanceTypes are available; otherwise, a verification error will be reported.</li>
	// <li>ANY: Verification passes if any InstanceType is available; otherwise, a verification error will be reported.</li>
	// 
	// Common reasons for unavailable InstanceTypes include the InstanceType being sold out, and the corresponding cloud disk being sold out.
	// If a model in InstanceTypes does not exist or has been abolished, a verification error will be reported regardless of the valid values set for InstanceTypesCheckPolicy.
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

	// Cloud disk type selection policy, whose default value is ORIGINAL. Valid values:
	// <li>ORIGINAL: Use the set cloud disk type.</li>
	// <li>AUTOMATIC: Automatically select the currently available cloud disk type.</li>
	DiskTypePolicy *string `json:"DiskTypePolicy,omitnil,omitempty" name:"DiskTypePolicy"`

	// HPC ID<br>
	// Note: This field is default to empty
	HpcClusterId *string `json:"HpcClusterId,omitnil,omitempty" name:"HpcClusterId"`

	// IPv6 public network bandwidth configuration. If the IPv6 address is available in the new instance, public network bandwidth can be allocated to the IPv6 address. This parameter is invalid if `Ipv6AddressCount` of the scaling group associated with the launch configuration is 0.
	IPv6InternetAccessible *IPv6InternetAccessible `json:"IPv6InternetAccessible,omitnil,omitempty" name:"IPv6InternetAccessible"`

	// Placement group ID. Only one is allowed.
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil,omitempty" name:"DisasterRecoverGroupIds"`

	// Image family name. Either image ID or image family name should be filled in, and only one of which can be filled.
	ImageFamily *string `json:"ImageFamily,omitnil,omitempty" name:"ImageFamily"`

	// CDC ID.
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`

	// Custom metadata.
	Metadata *Metadata `json:"Metadata,omitnil,omitempty" name:"Metadata"`
}

type CreateLaunchConfigurationRequest struct {
	*tchttp.BaseRequest
	
	// Display name of the launch configuration, which can contain letters, digits, underscores and hyphens (-), and dots. Up to of 60 bytes allowed..
	LaunchConfigurationName *string `json:"LaunchConfigurationName,omitnil,omitempty" name:"LaunchConfigurationName"`

	// [Image](https://intl.cloud.tencent.com/document/product/213/4940?from_cn_redirect=1) ID in the format of `img-xxx`. There are three types of images: <br/><li>Public images </li><li>Custom images </li><li>Shared images </li><br/>You can obtain the image IDs in the [CVM console](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE).</li><li>You can also use the [DescribeImages](https://intl.cloud.tencent.com/document/api/213/15715?from_cn_redirect=1) and look for `ImageId` in the response.</li>
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// Project ID of the launch configuration. The default project is used if it is left blank.
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
	// <li>POSTPAID_BY_HOUR: pay-as-you-go hourly</li>
	// <li>SPOTPAID: spot instance</li>
	// <li> CDCPAID: dedicated cluster</li>
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// Market options of the instance, such as parameters related to spot instances. This parameter is required for spot instances.
	InstanceMarketOptions *InstanceMarketOptionsRequest `json:"InstanceMarketOptions,omitnil,omitempty" name:"InstanceMarketOptions"`

	// List of instance models. Different instance models specify different resource specifications. Up to 10 instance models can be supported.
	// `InstanceType` and `InstanceTypes` are mutually exclusive, and one and only one of them must be entered.
	InstanceTypes []*string `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// CAM role name. This parameter can be obtained from the `roleName` field returned by DescribeRoleList API.
	CamRoleName *string `json:"CamRoleName,omitnil,omitempty" name:"CamRoleName"`

	// InstanceType verification policy, whose valid values include ALL and ANY, with the default value being ANY.
	// <li>ALL: Verification passes if all InstanceTypes are available; otherwise, a verification error will be reported.</li>
	// <li>ANY: Verification passes if any InstanceType is available; otherwise, a verification error will be reported.</li>
	// 
	// Common reasons for unavailable InstanceTypes include the InstanceType being sold out, and the corresponding cloud disk being sold out.
	// If a model in InstanceTypes does not exist or has been abolished, a verification error will be reported regardless of the valid values set for InstanceTypesCheckPolicy.
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

	// Cloud disk type selection policy, whose default value is ORIGINAL. Valid values:
	// <li>ORIGINAL: Use the set cloud disk type.</li>
	// <li>AUTOMATIC: Automatically select the currently available cloud disk type.</li>
	DiskTypePolicy *string `json:"DiskTypePolicy,omitnil,omitempty" name:"DiskTypePolicy"`

	// HPC ID<br>
	// Note: This field is default to empty
	HpcClusterId *string `json:"HpcClusterId,omitnil,omitempty" name:"HpcClusterId"`

	// IPv6 public network bandwidth configuration. If the IPv6 address is available in the new instance, public network bandwidth can be allocated to the IPv6 address. This parameter is invalid if `Ipv6AddressCount` of the scaling group associated with the launch configuration is 0.
	IPv6InternetAccessible *IPv6InternetAccessible `json:"IPv6InternetAccessible,omitnil,omitempty" name:"IPv6InternetAccessible"`

	// Placement group ID. Only one is allowed.
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil,omitempty" name:"DisasterRecoverGroupIds"`

	// Image family name. Either image ID or image family name should be filled in, and only one of which can be filled.
	ImageFamily *string `json:"ImageFamily,omitnil,omitempty" name:"ImageFamily"`

	// CDC ID.
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`

	// Custom metadata.
	Metadata *Metadata `json:"Metadata,omitnil,omitempty" name:"Metadata"`
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
	delete(f, "ImageFamily")
	delete(f, "DedicatedClusterId")
	delete(f, "Metadata")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLaunchConfigurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLaunchConfigurationResponseParams struct {
	// This parameter is returned when a launch configuration is created through this API, indicating the launch configuration ID.
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitnil,omitempty" name:"LaunchConfigurationId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// The method to adjust the desired capacity after the alarm is triggered. It is only available when `ScalingPolicyType` is `Simple`. Valid values: <br><li>`CHANGE_IN_CAPACITY`: Increase or decrease the desired capacity </li><li>`EXACT_CAPACITY`: Adjust to the specified desired capacity </li> <li>`PERCENT_CHANGE_IN_CAPACITY`: Adjust the desired capacity by percentage </li>
	AdjustmentType *string `json:"AdjustmentType,omitnil,omitempty" name:"AdjustmentType"`

	// Adjustment value for the expected number of instances after an alarm is triggered, which is applicable only to simple policies.
	// <li>When AdjustmentType is set to CHANGE_IN_CAPACITY, a positive value of AdjustmentValue indicates an increase in the number of instances after the alarm is triggered, while a negative value indicates a decrease in the number of instances after the alarm is triggered.</li>
	// <li>When AdjustmentType is set to EXACT_CAPACITY, the value of AdjustmentValue indicates the new desired number of instances after the alarm is triggered. It should be greater than or equal to 0.</li>
	// <li>When AdjustmentType is set to PERCENT_CHANGE_IN_CAPACITY, a positive value of AdjustmentValue indicates an increase in the number of instances by a percentage after the alarm is triggered, while a negative value indicates a decrease in the number of instances by a percentage after the alarm is triggered. Unit: %.</li>
	AdjustmentValue *int64 `json:"AdjustmentValue,omitnil,omitempty" name:"AdjustmentValue"`

	// Cooldown period (in seconds). This parameter is only applicable to a simple policy. Default value: 300.
	Cooldown *uint64 `json:"Cooldown,omitnil,omitempty" name:"Cooldown"`

	// Alarm monitoring metric. It is only available when `ScalingPolicyType` is `Simple`.
	MetricAlarm *MetricAlarm `json:"MetricAlarm,omitnil,omitempty" name:"MetricAlarm"`

	// Predefined monitoring item, applicable only to target tracking policies. Valid values:
	// <li>ASG_AVG_CPU_UTILIZATION: average CPU utilization</li>
	// <li>ASG_AVG_LAN_TRAFFIC_OUT: average outbound private network bandwidth</li>
	// <li>ASG_AVG_LAN_TRAFFIC_IN: average inbound private network bandwidth</li>
	// <li>ASG_AVG_WAN_TRAFFIC_OUT: average outbound public network bandwidth</li>
	// <li>ASG_AVG_WAN_TRAFFIC_IN: average inbound public network bandwidth</li>
	PredefinedMetricType *string `json:"PredefinedMetricType,omitnil,omitempty" name:"PredefinedMetricType"`

	// Target value, which is applicable only to target tracking policies.
	// <li>ASG_AVG_CPU_UTILIZATION: value range: [1, 100); unit: %.</li>
	// <li>ASG_AVG_LAN_TRAFFIC_OUT: value range: > 0; unit: Mbps.</li>
	// <li>ASG_AVG_LAN_TRAFFIC_IN: value range: > 0; unit: Mbps.</li>
	// <li>ASG_AVG_WAN_TRAFFIC_OUT: value range: > 0; unit: Mbps.</li>
	// <li>ASG_AVG_WAN_TRAFFIC_IN: value range: > 0; unit: Mbps.</li>
	TargetValue *uint64 `json:"TargetValue,omitnil,omitempty" name:"TargetValue"`

	// Instance warm-up period (in seconds). It is only available when `ScalingPolicyType` is `TARGET_TRACKING`. Value range: 0-3600. Default value: 300.
	EstimatedInstanceWarmup *uint64 `json:"EstimatedInstanceWarmup,omitnil,omitempty" name:"EstimatedInstanceWarmup"`

	// Whether to disable scale-in, which is applicable only to target tracking policies. Default value: false. Valid values:
	// <li>true: Target tracking policies trigger only scale-out.</li>
	// <li>false: Target tracking policies trigger both scale-out and scale-in.</li>
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

	// The method to adjust the desired capacity after the alarm is triggered. It is only available when `ScalingPolicyType` is `Simple`. Valid values: <br><li>`CHANGE_IN_CAPACITY`: Increase or decrease the desired capacity </li><li>`EXACT_CAPACITY`: Adjust to the specified desired capacity </li> <li>`PERCENT_CHANGE_IN_CAPACITY`: Adjust the desired capacity by percentage </li>
	AdjustmentType *string `json:"AdjustmentType,omitnil,omitempty" name:"AdjustmentType"`

	// Adjustment value for the expected number of instances after an alarm is triggered, which is applicable only to simple policies.
	// <li>When AdjustmentType is set to CHANGE_IN_CAPACITY, a positive value of AdjustmentValue indicates an increase in the number of instances after the alarm is triggered, while a negative value indicates a decrease in the number of instances after the alarm is triggered.</li>
	// <li>When AdjustmentType is set to EXACT_CAPACITY, the value of AdjustmentValue indicates the new desired number of instances after the alarm is triggered. It should be greater than or equal to 0.</li>
	// <li>When AdjustmentType is set to PERCENT_CHANGE_IN_CAPACITY, a positive value of AdjustmentValue indicates an increase in the number of instances by a percentage after the alarm is triggered, while a negative value indicates a decrease in the number of instances by a percentage after the alarm is triggered. Unit: %.</li>
	AdjustmentValue *int64 `json:"AdjustmentValue,omitnil,omitempty" name:"AdjustmentValue"`

	// Cooldown period (in seconds). This parameter is only applicable to a simple policy. Default value: 300.
	Cooldown *uint64 `json:"Cooldown,omitnil,omitempty" name:"Cooldown"`

	// Alarm monitoring metric. It is only available when `ScalingPolicyType` is `Simple`.
	MetricAlarm *MetricAlarm `json:"MetricAlarm,omitnil,omitempty" name:"MetricAlarm"`

	// Predefined monitoring item, applicable only to target tracking policies. Valid values:
	// <li>ASG_AVG_CPU_UTILIZATION: average CPU utilization</li>
	// <li>ASG_AVG_LAN_TRAFFIC_OUT: average outbound private network bandwidth</li>
	// <li>ASG_AVG_LAN_TRAFFIC_IN: average inbound private network bandwidth</li>
	// <li>ASG_AVG_WAN_TRAFFIC_OUT: average outbound public network bandwidth</li>
	// <li>ASG_AVG_WAN_TRAFFIC_IN: average inbound public network bandwidth</li>
	PredefinedMetricType *string `json:"PredefinedMetricType,omitnil,omitempty" name:"PredefinedMetricType"`

	// Target value, which is applicable only to target tracking policies.
	// <li>ASG_AVG_CPU_UTILIZATION: value range: [1, 100); unit: %.</li>
	// <li>ASG_AVG_LAN_TRAFFIC_OUT: value range: > 0; unit: Mbps.</li>
	// <li>ASG_AVG_LAN_TRAFFIC_IN: value range: > 0; unit: Mbps.</li>
	// <li>ASG_AVG_WAN_TRAFFIC_OUT: value range: > 0; unit: Mbps.</li>
	// <li>ASG_AVG_WAN_TRAFFIC_IN: value range: > 0; unit: Mbps.</li>
	TargetValue *uint64 `json:"TargetValue,omitnil,omitempty" name:"TargetValue"`

	// Instance warm-up period (in seconds). It is only available when `ScalingPolicyType` is `TARGET_TRACKING`. Value range: 0-3600. Default value: 300.
	EstimatedInstanceWarmup *uint64 `json:"EstimatedInstanceWarmup,omitnil,omitempty" name:"EstimatedInstanceWarmup"`

	// Whether to disable scale-in, which is applicable only to target tracking policies. Default value: false. Valid values:
	// <li>true: Target tracking policies trigger only scale-out.</li>
	// <li>false: Target tracking policies trigger both scale-out and scale-in.</li>
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// Data disk type. For restrictions on the data disk type, see [cloud block storage types](https://intl.cloud.tencent.com/document/product/362/2353?from_cn_redirect=1). Valid values:
	// <li>LOCAL_BASIC: Local Disk.</li>
	// <li>LOCAL_SSD: Local SSD.</li>
	// <li>CLOUD_BASIC: Basic Cloud Disk.</li>
	// <li>CLOUD_PREMIUM: Premium Disk.</li>
	// <li>CLOUD_SSD: Cloud SSD.</li>
	// <li>CLOUD_HSSD: Enhanced SSD.</li>
	// <li>CLOUD_TSSD: Tremendous SSD.</li>
	// The default value is consistent with the system disk type (SystemDisk.DiskType).
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// Data disk size, in GB. The minimum adjustment step size is 10 GB. The value range varies according to the data disk type. For specific restrictions, see [CVM instance configuration](https://intl.cloud.tencent.com/document/product/213/2177?from_cn_redirect=1). Default value: 0, which means that no data disk is purchased. For more restrictions, see the product documentation.
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// Data disk snapshot ID, such as `snap-l8psqwnt`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// Whether the data disk is terminated along with the instance. Valid values:
	// <li>TRUE: When the instance is terminated, the data disk is also terminated. This option is only supported for hourly postpaid cloud disks.</li>
	// <li>FALSE: When the instance is terminated, the data disk is retained.</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitnil,omitempty" name:"DeleteWithInstance"`

	// Whether the data disk is encrypted. Valid values:
	// <li>TRUE: Encrypted.</li>
	// <li>FALSE: Not encrypted.</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Encrypt *bool `json:"Encrypt,omitnil,omitempty" name:"Encrypt"`

	// Cloud disk performance (MB/s). This parameter is used to purchase extra performance for the cloud disk. For details on the feature and limits, see [Enhanced SSD Performance](https://intl.cloud.tencent.com/document/product/362/51896?from_cn_redirect=1#. E5.A2.9E.E5.BC.BA.E5.9E.8B-ssd-.E4.BA.91.E7.A1.AC.E7.9B.98.E9.A2.9D.E5.A4.96 .E6.80.A7.E8.83.BD).
	// This feature is only available to enhanced SSD (`CLOUD_HSSD`) and tremendous SSD (`CLOUD_TSSD`) disks with a capacity greater than 460 GB.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ThroughputPerformance *uint64 `json:"ThroughputPerformance,omitnil,omitempty" name:"ThroughputPerformance"`

	// Burst performance: Whether to enable burst performance. The default value is false.
	// 
	// Note: This feature is in beta test and requires a ticket to be submitted for usage.
	// Note: This field may return null, indicating that no valid values can be obtained.
	BurstPerformance *bool `json:"BurstPerformance,omitnil,omitempty" name:"BurstPerformance"`
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
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// ID list of auto scaling groups.
	AutoScalingGroupIds []*string `json:"AutoScalingGroupIds,omitnil,omitempty" name:"AutoScalingGroupIds"`

	// Excludes cancelled type activities when querying. Default value is false, which means cancelled type activities are not excluded.
	ExcludeCancelledActivity *bool `json:"ExcludeCancelledActivity,omitnil,omitempty" name:"ExcludeCancelledActivity"`
}

type DescribeAutoScalingGroupLastActivitiesRequest struct {
	*tchttp.BaseRequest
	
	// ID list of auto scaling groups.
	AutoScalingGroupIds []*string `json:"AutoScalingGroupIds,omitnil,omitempty" name:"AutoScalingGroupIds"`

	// Excludes cancelled type activities when querying. Default value is false, which means cancelled type activities are not excluded.
	ExcludeCancelledActivity *bool `json:"ExcludeCancelledActivity,omitnil,omitempty" name:"ExcludeCancelledActivity"`
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
	delete(f, "ExcludeCancelledActivity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAutoScalingGroupLastActivitiesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoScalingGroupLastActivitiesResponseParams struct {
	// Information set of eligible scaling activities. Scaling groups without scaling activities are not returned. For example, if there are 50 auto scaling group IDs but only 45 records are returned, it indicates that 5 of the auto scaling groups do not have scaling activities.
	ActivitySet []*Activity `json:"ActivitySet,omitnil,omitempty" name:"ActivitySet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// Filter criteria
	// 
	// <li>launch-configuration-id - String - required: no - (filter condition) filter by launch configuration ID.</li>
	// <li>launch-configuration-name - String - required: no - (filter condition) filter by launch configuration name.</li>
	// <li>vague-launch-configuration-name - String - required: no - (filter condition) fuzzy search by launch configuration name.</li>
	// <li>tag-key - String - required: no - (filter condition) filter by tag key.</li>
	// <li>tag-value - String - required: no - (filter condition) filter by tag value.</li>
	// <li>tag:tag-key - String - required: no - (filter condition) filter by Tag key-value pair. Replace tag-key with a specific tag key. See Example 3 for usage.</li>
	// The maximum number of `Filters` per request is 10, and the maximum number of `Filter.Values` is 5. The parameter does not support specifying both `LaunchConfigurationIds` and `Filters`.
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

	// Filter criteria
	// 
	// <li>launch-configuration-id - String - required: no - (filter condition) filter by launch configuration ID.</li>
	// <li>launch-configuration-name - String - required: no - (filter condition) filter by launch configuration name.</li>
	// <li>vague-launch-configuration-name - String - required: no - (filter condition) fuzzy search by launch configuration name.</li>
	// <li>tag-key - String - required: no - (filter condition) filter by tag key.</li>
	// <li>tag-value - String - required: no - (filter condition) filter by tag value.</li>
	// <li>tag:tag-key - String - required: no - (filter condition) filter by Tag key-value pair. Replace tag-key with a specific tag key. See Example 3 for usage.</li>
	// The maximum number of `Filters` per request is 10, and the maximum number of `Filter.Values` is 5. The parameter does not support specifying both `LaunchConfigurationIds` and `Filters`.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
type DescribeRefreshActivitiesRequestParams struct {
	// List of refresh activity IDs. IDs are formatted like: 'asr-5l2ejpfo'. The upper limit per request is 100. Parameters do not support specifying both RefreshActivityIds and Filters simultaneously.
	RefreshActivityIds []*string `json:"RefreshActivityIds,omitnil,omitempty" name:"RefreshActivityIds"`

	// Filter criteria
	// 
	// <li> auto-scaling-group-id - String - Required: No - (Filter) Filter by auto scaling group ID.</li>
	// <li> refresh-activity-status-code - String - Required: No - (Filter)Filter based on refresh activity status. (INIT: Initialization | RUNNING:Running | SUCCESSFUL: Successful Activity | FAILED_PAUSE: Failed & Paused | AUTO_PAUSE: Auto Paused | MANUAL_PAUSE: Manually Paused | CANCELLED: Activity Cancelled | FAILED: Activity Failed)</li>
	// <li> refresh-activity-type - String - Required: No - (Filter) Filter by refresh activity types. (NORMAL: Regular Refresh Activity | ROLLBACK: Rollback Refresh Activity)</li>
	// <li> refresh-activity-id - String - Required: No - (Filter) Filter by refresh activity ID.</li>
	// <li>The upper limit of Filters per request is 10, and that of Filter.Values is 5. The RefreshActivityIds and Filters parameters cannot be specified at the same time.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Number of returned pieces. Default value: 20. Maximum value: 100. For further information on Limit, please refer to relevant sections in API [Overview] (https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset, 0 by default. For further information on Offset, please refer to relevant sections in API [Overview] (https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeRefreshActivitiesRequest struct {
	*tchttp.BaseRequest
	
	// List of refresh activity IDs. IDs are formatted like: 'asr-5l2ejpfo'. The upper limit per request is 100. Parameters do not support specifying both RefreshActivityIds and Filters simultaneously.
	RefreshActivityIds []*string `json:"RefreshActivityIds,omitnil,omitempty" name:"RefreshActivityIds"`

	// Filter criteria
	// 
	// <li> auto-scaling-group-id - String - Required: No - (Filter) Filter by auto scaling group ID.</li>
	// <li> refresh-activity-status-code - String - Required: No - (Filter)Filter based on refresh activity status. (INIT: Initialization | RUNNING:Running | SUCCESSFUL: Successful Activity | FAILED_PAUSE: Failed & Paused | AUTO_PAUSE: Auto Paused | MANUAL_PAUSE: Manually Paused | CANCELLED: Activity Cancelled | FAILED: Activity Failed)</li>
	// <li> refresh-activity-type - String - Required: No - (Filter) Filter by refresh activity types. (NORMAL: Regular Refresh Activity | ROLLBACK: Rollback Refresh Activity)</li>
	// <li> refresh-activity-id - String - Required: No - (Filter) Filter by refresh activity ID.</li>
	// <li>The upper limit of Filters per request is 10, and that of Filter.Values is 5. The RefreshActivityIds and Filters parameters cannot be specified at the same time.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Number of returned pieces. Default value: 20. Maximum value: 100. For further information on Limit, please refer to relevant sections in API [Overview] (https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset, 0 by default. For further information on Offset, please refer to relevant sections in API [Overview] (https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeRefreshActivitiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRefreshActivitiesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RefreshActivityIds")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRefreshActivitiesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRefreshActivitiesResponseParams struct {
	// Number of refresh activities that meet the conditions.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// A collection of information about refresh activities that meet the conditions.
	RefreshActivitySet []*RefreshActivity `json:"RefreshActivitySet,omitnil,omitempty" name:"RefreshActivitySet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRefreshActivitiesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRefreshActivitiesResponseParams `json:"Response"`
}

func (r *DescribeRefreshActivitiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRefreshActivitiesResponse) FromJsonString(s string) error {
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	//
	// Deprecated: AutomationService is deprecated.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

// Predefined struct for user
type ExitStandbyRequestParams struct {
	// Scaling group ID.
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// List of CVM instances in standby status.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type ExitStandbyRequest struct {
	*tchttp.BaseRequest
	
	// Scaling group ID.
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// List of CVM instances in standby status.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *ExitStandbyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExitStandbyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExitStandbyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExitStandbyResponseParams struct {
	// Scaling activity ID.
	// Note: This field may return null, indicating that no valid value can be obtained.
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ExitStandbyResponse struct {
	*tchttp.BaseResponse
	Response *ExitStandbyResponseParams `json:"Response"`
}

func (r *ExitStandbyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExitStandbyResponse) FromJsonString(s string) error {
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
	// CVM HostName.
	// <li>Dots (.) and hyphens (-) cannot be used as the first or last character of HostName, and cannot be used consecutively.</li>
	// <li>Windows instances are not supported.</li>
	// <li>Instances of other types (e.g., Linux): The length of the character should be within the range of [2, 40]. Multiple dots (.) are allowed. Each segment between dot marks can consist of letters (case-insensitive), digits, and hyphens (-). Using only digits is not allowed.</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	HostName *string `json:"HostName,omitnil,omitempty" name:"HostName"`

	// The style of the CVM HostName. Valid values include ORIGINAL and UNIQUE, and the default value is ORIGINAL.
	// <li>ORIGINAL: AS passes HostName filled in the input parameters to CVM. CVM may append serial numbers to HostName, which can result in conflicts with HostName of instances in the scaling group.</li>
	// <li> UNIQUE: HostName filled in the input parameters acts as a prefix for the HostName. AS and CVM will expand this prefix to ensure that HostName of the instance in the scaling group is unique.</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	HostNameStyle *string `json:"HostNameStyle,omitnil,omitempty" name:"HostNameStyle"`

	// HostName suffix for CVM.
	// <li>Dots (.) and hyphens (-) cannot be used as the first or last character of HostNameSuffix, and cannot be used consecutively.</li>
	// <li>Windows instances are not supported.</li>
	// <li>Instances of other types (e.g., Linux): The length of the character should be within the range of [1, 37], and the combined length with HostName should not exceed 39. Multiple dots (.) are allowed. Each segment between dots can consist of letters (case-insensitive), digits, and hyphens (-).</li>
	// Assume the suffix name is suffix and the original HostName is test.0, then the final HostName is test.0.suffix.
	// Note: This field may return null, indicating that no valid values can be obtained.
	HostNameSuffix *string `json:"HostNameSuffix,omitnil,omitempty" name:"HostNameSuffix"`
}

type IPv6InternetAccessible struct {
	// Network billing mode. Valid values: TRAFFIC_POSTPAID_BY_HOUR, BANDWIDTH_PACKAGE. Default value: TRAFFIC_POSTPAID_BY_HOUR. For the current account type, see [Account Type Description](https://intl.cloud.tencent.com/document/product/684/15246?from_cn_redirect=1#judge).
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

	// Auto-renewal flag. Valid values: <li>NOTIFY_AND_AUTO_RENEW: Notify upon expiration and automatically renew.</li> <li>NOTIFY_AND_MANUAL_RENEW: Notify upon expiration but do not auto-renew.</li> <li>DISABLE_NOTIFY_AND_MANUAL_RENEW: Do not notify and do not auto-renew</li> Default value: NOTIFY_AND_MANUAL_RENEW. If this parameter is set to NOTIFY_AND_AUTO_RENEW, and the account balance is sufficient, the instance will automatically renew monthly upon its expiration date.
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

type InstanceMarketOptionsRequest struct {
	// Bidding-related options
	SpotOptions *SpotMarketOptions `json:"SpotOptions,omitnil,omitempty" name:"SpotOptions"`

	// Market option type. The value can only be spot currently.
	MarketType *string `json:"MarketType,omitnil,omitempty" name:"MarketType"`
}

type InstanceNameIndexSettings struct {
	// Whether to enable instance creation sequencing, which is disabled by default. Valid values: <li>TRUE: Indicates that instance creation sequencing is enabled. <li>FALSE: Indicates that instance creation sequencing is disabled.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// Initial sequence number, with a value range of [0, 99,999,999]. When the sequence number exceeds this range after incrementing, scale-out activities will fail. <li>Upon the first enabling of instance name sequencing: The default value is 0. <li>Upon the enabling of instance name sequencing (not for the first time): If this parameter is not specified, the historical sequence number will be carried forward. Lowering the initial sequence number may result in duplicate instance name sequences within the scaling group.
	// Note: This field may return null, indicating that no valid value can be obtained.
	BeginIndex *int64 `json:"BeginIndex,omitnil,omitempty" name:"BeginIndex"`
}

type InstanceNameSettings struct {
	// CVM instance name. Value range: 2-108.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Type of CVM instance name. Valid values: `ORIGINAL` and `UNIQUE`. Default value: `ORIGINAL`.
	// 
	// `ORIGINAL`: Auto Scaling sends the input parameter `InstanceName` to the CVM directly. The CVM may append a serial number to the `InstanceName`. The `InstanceName` of the instances within the scaling group may conflict.
	// 
	// `UNIQUE`: the input parameter `InstanceName` is the prefix of an instance name. Auto Scaling and CVM expand it. The `InstanceName` of an instance in the scaling group is unique.
	InstanceNameStyle *string `json:"InstanceNameStyle,omitnil,omitempty" name:"InstanceNameStyle"`

	// CVM instance name suffix. The length of the character is within the range of [1, 105], and the combined length with InstanceName should not exceed 107.
	// 
	// Assume the suffix name is suffix and the original instance name is test.0, then the final instance name is test.0.suffix.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceNameSuffix *string `json:"InstanceNameSuffix,omitnil,omitempty" name:"InstanceNameSuffix"`
}

type InstanceTag struct {
	// Tag key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Tag value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type InternetAccessible struct {
	// Network billing type. Valid values: <li>BANDWIDTH_PREPAID: prepaid by bandwidth;</li> <li>TRAFFIC_POSTPAID_BY_HOUR: postpaid by traffic per hour;</li> <li>BANDWIDTH_POSTPAID_BY_HOUR: postpaid by bandwidth per hour;</li> <li>BANDWIDTH_PACKAGE: bandwidth package users.</li> Default value: TRAFFIC_POSTPAID_BY_HOUR.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InternetChargeType *string `json:"InternetChargeType,omitnil,omitempty" name:"InternetChargeType"`

	// The maximum outbound bandwidth in Mbps of the public network. The default value is 0 Mbps. The upper limit of bandwidth varies by model. For more information, see [Purchase Network Bandwidth](https://intl.cloud.tencent.com/document/product/213/509?from_cn_redirect=1).
	// Note: This field may return null, indicating that no valid values can be obtained.
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`

	// Whether to assign a public IP address. Valid values: <li>TRUE: Allocate a public IP address.</li> <li>FALSE: Do not allocate a public IP address.</li> When the public network bandwidth is greater than 0 Mbps, you can choose whether to enable this feature based on your needs. By default, this feature is enabled. When the public network bandwidth is 0, public IP address assignment is not allowed.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PublicIpAssigned *bool `json:"PublicIpAssigned,omitnil,omitempty" name:"PublicIpAssigned"`

	// Bandwidth package ID. You can obtain the ID from the `BandwidthPackageId` field in the response of the [DescribeBandwidthPackages](https://intl.cloud.tencent.com/document/api/215/19209?from_cn_redirect=1) API.
	// Note: this field may return null, indicating that no valid value was found.
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil,omitempty" name:"BandwidthPackageId"`
}

type InvocationResult struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Execution activity ID.
	InvocationId *string `json:"InvocationId,omitnil,omitempty" name:"InvocationId"`

	// Task ID.
	InvocationTaskId *string `json:"InvocationTaskId,omitnil,omitempty" name:"InvocationTaskId"`

	// Command ID.
	CommandId *string `json:"CommandId,omitnil,omitempty" name:"CommandId"`

	// Specifies the execution task status.
	TaskStatus *string `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// Specifies the exception information during execution.
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

	// Current status of the launch configuration. Valid values: <li>NORMAL: Normal.</li> <li>IMAGE_ABNORMAL: Image exception in the launch configuration.</li> <li>CBS_SNAP_ABNORMAL: Exception with data disk snapshot in the launch configuration.</li> <li>SECURITY_GROUP_ABNORMAL: Security group exception in the launch configuration.</li>
	LaunchConfigurationStatus *string `json:"LaunchConfigurationStatus,omitnil,omitempty" name:"LaunchConfigurationStatus"`

	// Instance billing type, with the CVM default value processed as POSTPAID_BY_HOUR. <li>POSTPAID_BY_HOUR: Hourly postpaid billing.</li> <li>SPOTPAID: Spot billing.</li>
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

	// Cloud disk type selection policy. Valid values: <li>ORIGINAL: Use the set cloud disk type.</li> <li>AUTOMATIC: Automatically select available cloud disk types in the current availability zone.</li>
	DiskTypePolicy *string `json:"DiskTypePolicy,omitnil,omitempty" name:"DiskTypePolicy"`

	// HPC ID<br>
	// Note: This field is default to empty
	HpcClusterId *string `json:"HpcClusterId,omitnil,omitempty" name:"HpcClusterId"`

	// IPv6 public network bandwidth configuration.
	IPv6InternetAccessible *IPv6InternetAccessible `json:"IPv6InternetAccessible,omitnil,omitempty" name:"IPv6InternetAccessible"`

	// Placement group ID, supporting specification of only one.
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil,omitempty" name:"DisasterRecoverGroupIds"`

	// Image family name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ImageFamily *string `json:"ImageFamily,omitnil,omitempty" name:"ImageFamily"`

	// CDC ID.
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`
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
	// Instance login password. The password complexity requirements vary according to the operating system type. The details are as follows:
	// - For a Linux system, the password should contain 8 to 30 characters consisting of at least two of the four character types: lowercase letters, uppercase letters, digits, and special characters.
	// - For a Windows system, the password should contain 12 to 30 characters consisting of at least three of the four character types: lowercase letters, uppercase letters, digits, and special characters.
	// - If this parameter is not specified, the system will generate a random password and notify the user via the message centerSupported special characters: ( ) ` ~ ! @ # $ % ^ & * - + = | { } [ ] : ; ' , . ? /
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// List of key IDs. After an instance is associated with a key, you can access the instance with the private key in the key pair. You can call `DescribeKeyPairs` to obtain `KeyId`. Key and password cannot be specified at the same time. Windows instances do not support keys. Currently, you can only specify one key when purchasing an instance.
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`

	// Retain the original settings of the image. This parameter cannot be specified simultaneously with Password or KeyIds.N. It can only be set to TRUE when you create an instance by using a custom image, shared image, or externally imported image. Valid values:
	// <li>TRUE: Retain the login settings of the image.</li>
	// <li>FALSE: Do not retain the login settings of the image.</li> Default value: FALSE.
	KeepImageLogin *bool `json:"KeepImageLogin,omitnil,omitempty" name:"KeepImageLogin"`
}

type Metadata struct {
	// Custom metadata key-value pair list.
	Items []*MetadataItem `json:"Items,omitnil,omitempty" name:"Items"`
}

type MetadataItem struct {
	// Custom metadata key.
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Custom metadata value.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type MetricAlarm struct {
	// Comparison operator. Value range: <br><li>GREATER_THAN: greater than </li><li>GREATER_THAN_OR_EQUAL_TO: greater than or equal to </li><li>LESS_THAN: less than </li><li> LESS_THAN_OR_EQUAL_TO: less than or equal to </li><li> EQUAL_TO: equal to </li> <li>NOT_EQUAL_TO: not equal to </li>
	ComparisonOperator *string `json:"ComparisonOperator,omitnil,omitempty" name:"ComparisonOperator"`

	// Metric names, with the following optional fields: <br><li>CPU_UTILIZATION: CPU utilization.</li> <li>MEM_UTILIZATION: Memory utilization.</li> <li>LAN_TRAFFIC_OUT: Private network outbound bandwidth.</li> <li>LAN_TRAFFIC_IN: Private network inbound bandwidth.</li> <li>WAN_TRAFFIC_OUT: Public network outbound bandwidth.</li> <li>WAN_TRAFFIC_IN: Public network inbound bandwidth.</li> <li>TCP_CURR_ESTAB: TCP connections.</li>
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// Alarm threshold values: <br><li>CPU_UTILIZATION: [1, 100], Unit: %.</li> <li>MEM_UTILIZATION: [1, 100], Unit: %.</li> <li>LAN_TRAFFIC_OUT: >0, Unit: Mbps.</li> <li>LAN_TRAFFIC_IN: >0, Unit: Mbps.</li> <li>WAN_TRAFFIC_OUT: >0, Unit: Mbps.</li> <li>WAN_TRAFFIC_IN: >0, Unit: Mbps.</li> <li>TCP_CURR_ESTAB: >0, Unit: Count.</li>
	Threshold *uint64 `json:"Threshold,omitnil,omitempty" name:"Threshold"`

	// Time period in seconds. Enumerated values: 60, 300.
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Number of repetitions. Value range: [1, 10]
	ContinuousTime *uint64 `json:"ContinuousTime,omitnil,omitempty" name:"ContinuousTime"`

	// Statistics type. Value range: <br><li>AVERAGE: average </li><li>MAXIMUM: maximum <li>MINIMUM: minimum </li><br> Default value: AVERAGE
	Statistic *string `json:"Statistic,omitnil,omitempty" name:"Statistic"`

	// Precise alarm threshold values. This parameter is not used as an input argument but is used solely as an output parameter for the query API: <br><li>CPU_UTILIZATION: (0, 100], Unit: %.</li> <li>MEM_UTILIZATION: (0, 100], Unit: %.</li> <li>LAN_TRAFFIC_OUT: >0, Unit: Mbps.</li> <li>LAN_TRAFFIC_IN: >0, Unit: Mbps.</li> <li>WAN_TRAFFIC_OUT: >0, Unit: Mbps.</li> <li>WAN_TRAFFIC_IN: >0, Unit: Mbps.</li> <li>TCP_CURR_ESTAB: >0, Unit: Count.</li>
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

	// Termination policy, whose maximum length is currently 1. Valid values include OLDEST_INSTANCE and NEWEST_INSTANCE.
	// <li>OLDEST_INSTANCE: Terminate the oldest instance in the scaling group first.</li>
	// <li>NEWEST_INSTANCE: Terminate the newest instance in the scaling group first.</li>
	TerminationPolicies []*string `json:"TerminationPolicies,omitnil,omitempty" name:"TerminationPolicies"`

	// VPC ID. This field is left empty for basic networks. You need to specify SubnetIds when modifying the network of the auto scaling group to a VPC with a specified VPC ID. Specify Zones when modifying the network to a basic network.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// List of availability zones
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// Retry policy, whose valid values include IMMEDIATE_RETRY, INCREMENTAL_INTERVALS, and NO_RETRY, with the default value being IMMEDIATE_RETRY. A partially successful scaling activity is considered a failed activity.
	// <li>IMMEDIATE_RETRY: Immediately retry, and quickly retry in a short period. There will be no retry anymore after a certain number of consecutive failures (5).</li>
	// <li>INCREMENTAL_INTERVALS: Retry with incremental intervals. As the number of consecutive failures increases, the retry intervals gradually become longer, ranging from seconds to one day.</li>
	// <li>NO_RETRY: There will be no retry until another user call or alarm information is received.</li>
	RetryPolicy *string `json:"RetryPolicy,omitnil,omitempty" name:"RetryPolicy"`

	// AZ verification policy, whose valid values include ALL and ANY, with the default value being ANY. This policy comes into effect when actual changes are made to resource-related fields in the scaling group (such as launch configuration, AZ, or subnet).
	// <li>ALL: Verification passes if all AZs or subnets are available; otherwise, a verification error will be reported.<li>
	// <li>ANY: Verification passes if any AZ or subnet is available; otherwise, a verification error will be reported.</li>
	// 
	// Common reasons for unavailable AZs or subnets include the CVM InstanceType in the AZ being sold out, the CBS cloud disk in the AZ being sold out, insufficient quota in the AZ, and insufficient IP addresses in the subnet.
	// If there is no AZ or subnet in Zones/SubnetIds, a verification error will be reported regardless of the values of ZonesCheckPolicy.
	ZonesCheckPolicy *string `json:"ZonesCheckPolicy,omitnil,omitempty" name:"ZonesCheckPolicy"`

	// Service settings such as unhealthy instance replacement.
	ServiceSettings *ServiceSettings `json:"ServiceSettings,omitnil,omitempty" name:"ServiceSettings"`

	// The number of IPv6 addresses that an instance has. Valid values: 0 and 1.
	Ipv6AddressCount *int64 `json:"Ipv6AddressCount,omitnil,omitempty" name:"Ipv6AddressCount"`

	// Multi-AZ/multi-subnet policy, whose valid values include PRIORITY and EQUALITY, with the default value being PRIORITY.
	// <li>PRIORITY: Instances are attempted to be created taking the order of the AZ/subnet list as the priority. If the highest-priority AZ/subnet can create instances successfully, instances can always be created in that AZ/subnet.</li>
	// <li>EQUALITY: The instances added through scale-out will be distributed across multiple AZs/subnets to ensure a relatively balanced number of instances in each AZ/subnet after scaling out.</li>
	// 
	// Points to consider regarding this policy:
	// <li>When the scaling group is based on a classic network, this policy applies to the multi-AZ; when the scaling group is based on a VPC network, this policy applies to the multi-subnet, in this case, the AZs are no longer considered. For example, if there are four subnets labeled A, B, C, and D, where A, B, and C are in AZ 1 and D is in AZ 2, the subnets A, B, C, and D are considered for sorting without regard to AZs 1 and 2.</li>
	// <li>This policy applies to the multi-AZ/multi-subnet and not to the InstanceTypes parameter of the launch configuration, which is selected according to the priority policy.</li>
	// <li>When instances are created according to the PRIORITY policy, ensure the policy for multiple models first, followed by the policy for the multi-AZ/subnet. For example, with models A and B and subnets 1, 2, and 3, attempts will be made in the order of A1, A2, A3, B1, B2, and B3. If A1 is sold out, A2 will be attempted (instead of B1).</li>
	MultiZoneSubnetPolicy *string `json:"MultiZoneSubnetPolicy,omitnil,omitempty" name:"MultiZoneSubnetPolicy"`

	// Scaling group instance health check type, whose valid values include:
	// <li>CVM: Determines whether an instance is unhealthy based on its network status. An unhealthy network status is indicated by an event where instances become unreachable via PING. Detailed criteria of judgment can be found in [Instance Health Check](https://intl.cloud.tencent.com/document/product/377/8553?from_cn_redirect=1).</li>
	// <li>CLB: Determines whether an instance is unhealthy based on the health check status of CLB. For principles behind CLB health checks, see [Health Check](https://intl.cloud.tencent.com/document/product/214/6097?from_cn_redirect=1).</li>
	HealthCheckType *string `json:"HealthCheckType,omitnil,omitempty" name:"HealthCheckType"`

	// Grace period of the CLB health check
	LoadBalancerHealthCheckGracePeriod *uint64 `json:"LoadBalancerHealthCheckGracePeriod,omitnil,omitempty" name:"LoadBalancerHealthCheckGracePeriod"`

	// Instance assignment policy, whose valid values include LAUNCH_CONFIGURATION and SPOT_MIXED.
	// <li>LAUNCH_CONFIGURATION: Represent the traditional mode of assigning instances according to the launch configuration.</li>
	// <li>SPOT_MIXED: Represent the spot mixed mode. Currently, this mode is supported only when the launch configuration is set to the pay-as-you-go billing mode. In the mixed mode, the scaling group will scale out pay-as-you-go models or spot models based on the predefined settings. When the mixed mode is used, the billing type of the associated launch configuration cannot be modified.</li>
	InstanceAllocationPolicy *string `json:"InstanceAllocationPolicy,omitnil,omitempty" name:"InstanceAllocationPolicy"`

	// Specifies how to assign pay-as-you-go instances and spot instances.
	// This parameter is valid only when `InstanceAllocationPolicy` is set to `SPOT_MIXED`.
	SpotMixedAllocationPolicy *SpotMixedAllocationPolicy `json:"SpotMixedAllocationPolicy,omitnil,omitempty" name:"SpotMixedAllocationPolicy"`

	// Capacity rebalancing feature, which is applicable only to spot instances within the scaling group. Valid values:
	// <li>TRUE: Enable this feature. When spot instances in the scaling group are about to be automatically recycled by the spot instance service, AS proactively initiates the termination process of the spot instances. If there is a configured scale-in hook, it will be triggered before termination. After the termination process starts, AS asynchronously initiates the scale-out to reach the expected number of instances.</li>
	// <li>FALSE: Disable this feature. AS waits for the spot instance to be terminated before scaling out to reach the number of instances expected by the scaling group.</li>
	CapacityRebalance *bool `json:"CapacityRebalance,omitnil,omitempty" name:"CapacityRebalance"`

	// Instance name sequencing settings. When enabled, an incremental numeric sequence will be appended to the names of instances automatically created within the scaling group.
	InstanceNameIndexSettings *InstanceNameIndexSettings `json:"InstanceNameIndexSettings,omitnil,omitempty" name:"InstanceNameIndexSettings"`
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

	// Termination policy, whose maximum length is currently 1. Valid values include OLDEST_INSTANCE and NEWEST_INSTANCE.
	// <li>OLDEST_INSTANCE: Terminate the oldest instance in the scaling group first.</li>
	// <li>NEWEST_INSTANCE: Terminate the newest instance in the scaling group first.</li>
	TerminationPolicies []*string `json:"TerminationPolicies,omitnil,omitempty" name:"TerminationPolicies"`

	// VPC ID. This field is left empty for basic networks. You need to specify SubnetIds when modifying the network of the auto scaling group to a VPC with a specified VPC ID. Specify Zones when modifying the network to a basic network.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// List of availability zones
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// Retry policy, whose valid values include IMMEDIATE_RETRY, INCREMENTAL_INTERVALS, and NO_RETRY, with the default value being IMMEDIATE_RETRY. A partially successful scaling activity is considered a failed activity.
	// <li>IMMEDIATE_RETRY: Immediately retry, and quickly retry in a short period. There will be no retry anymore after a certain number of consecutive failures (5).</li>
	// <li>INCREMENTAL_INTERVALS: Retry with incremental intervals. As the number of consecutive failures increases, the retry intervals gradually become longer, ranging from seconds to one day.</li>
	// <li>NO_RETRY: There will be no retry until another user call or alarm information is received.</li>
	RetryPolicy *string `json:"RetryPolicy,omitnil,omitempty" name:"RetryPolicy"`

	// AZ verification policy, whose valid values include ALL and ANY, with the default value being ANY. This policy comes into effect when actual changes are made to resource-related fields in the scaling group (such as launch configuration, AZ, or subnet).
	// <li>ALL: Verification passes if all AZs or subnets are available; otherwise, a verification error will be reported.<li>
	// <li>ANY: Verification passes if any AZ or subnet is available; otherwise, a verification error will be reported.</li>
	// 
	// Common reasons for unavailable AZs or subnets include the CVM InstanceType in the AZ being sold out, the CBS cloud disk in the AZ being sold out, insufficient quota in the AZ, and insufficient IP addresses in the subnet.
	// If there is no AZ or subnet in Zones/SubnetIds, a verification error will be reported regardless of the values of ZonesCheckPolicy.
	ZonesCheckPolicy *string `json:"ZonesCheckPolicy,omitnil,omitempty" name:"ZonesCheckPolicy"`

	// Service settings such as unhealthy instance replacement.
	ServiceSettings *ServiceSettings `json:"ServiceSettings,omitnil,omitempty" name:"ServiceSettings"`

	// The number of IPv6 addresses that an instance has. Valid values: 0 and 1.
	Ipv6AddressCount *int64 `json:"Ipv6AddressCount,omitnil,omitempty" name:"Ipv6AddressCount"`

	// Multi-AZ/multi-subnet policy, whose valid values include PRIORITY and EQUALITY, with the default value being PRIORITY.
	// <li>PRIORITY: Instances are attempted to be created taking the order of the AZ/subnet list as the priority. If the highest-priority AZ/subnet can create instances successfully, instances can always be created in that AZ/subnet.</li>
	// <li>EQUALITY: The instances added through scale-out will be distributed across multiple AZs/subnets to ensure a relatively balanced number of instances in each AZ/subnet after scaling out.</li>
	// 
	// Points to consider regarding this policy:
	// <li>When the scaling group is based on a classic network, this policy applies to the multi-AZ; when the scaling group is based on a VPC network, this policy applies to the multi-subnet, in this case, the AZs are no longer considered. For example, if there are four subnets labeled A, B, C, and D, where A, B, and C are in AZ 1 and D is in AZ 2, the subnets A, B, C, and D are considered for sorting without regard to AZs 1 and 2.</li>
	// <li>This policy applies to the multi-AZ/multi-subnet and not to the InstanceTypes parameter of the launch configuration, which is selected according to the priority policy.</li>
	// <li>When instances are created according to the PRIORITY policy, ensure the policy for multiple models first, followed by the policy for the multi-AZ/subnet. For example, with models A and B and subnets 1, 2, and 3, attempts will be made in the order of A1, A2, A3, B1, B2, and B3. If A1 is sold out, A2 will be attempted (instead of B1).</li>
	MultiZoneSubnetPolicy *string `json:"MultiZoneSubnetPolicy,omitnil,omitempty" name:"MultiZoneSubnetPolicy"`

	// Scaling group instance health check type, whose valid values include:
	// <li>CVM: Determines whether an instance is unhealthy based on its network status. An unhealthy network status is indicated by an event where instances become unreachable via PING. Detailed criteria of judgment can be found in [Instance Health Check](https://intl.cloud.tencent.com/document/product/377/8553?from_cn_redirect=1).</li>
	// <li>CLB: Determines whether an instance is unhealthy based on the health check status of CLB. For principles behind CLB health checks, see [Health Check](https://intl.cloud.tencent.com/document/product/214/6097?from_cn_redirect=1).</li>
	HealthCheckType *string `json:"HealthCheckType,omitnil,omitempty" name:"HealthCheckType"`

	// Grace period of the CLB health check
	LoadBalancerHealthCheckGracePeriod *uint64 `json:"LoadBalancerHealthCheckGracePeriod,omitnil,omitempty" name:"LoadBalancerHealthCheckGracePeriod"`

	// Instance assignment policy, whose valid values include LAUNCH_CONFIGURATION and SPOT_MIXED.
	// <li>LAUNCH_CONFIGURATION: Represent the traditional mode of assigning instances according to the launch configuration.</li>
	// <li>SPOT_MIXED: Represent the spot mixed mode. Currently, this mode is supported only when the launch configuration is set to the pay-as-you-go billing mode. In the mixed mode, the scaling group will scale out pay-as-you-go models or spot models based on the predefined settings. When the mixed mode is used, the billing type of the associated launch configuration cannot be modified.</li>
	InstanceAllocationPolicy *string `json:"InstanceAllocationPolicy,omitnil,omitempty" name:"InstanceAllocationPolicy"`

	// Specifies how to assign pay-as-you-go instances and spot instances.
	// This parameter is valid only when `InstanceAllocationPolicy` is set to `SPOT_MIXED`.
	SpotMixedAllocationPolicy *SpotMixedAllocationPolicy `json:"SpotMixedAllocationPolicy,omitnil,omitempty" name:"SpotMixedAllocationPolicy"`

	// Capacity rebalancing feature, which is applicable only to spot instances within the scaling group. Valid values:
	// <li>TRUE: Enable this feature. When spot instances in the scaling group are about to be automatically recycled by the spot instance service, AS proactively initiates the termination process of the spot instances. If there is a configured scale-in hook, it will be triggered before termination. After the termination process starts, AS asynchronously initiates the scale-out to reach the expected number of instances.</li>
	// <li>FALSE: Disable this feature. AS waits for the spot instance to be terminated before scaling out to reach the number of instances expected by the scaling group.</li>
	CapacityRebalance *bool `json:"CapacityRebalance,omitnil,omitempty" name:"CapacityRebalance"`

	// Instance name sequencing settings. When enabled, an incremental numeric sequence will be appended to the names of instances automatically created within the scaling group.
	InstanceNameIndexSettings *InstanceNameIndexSettings `json:"InstanceNameIndexSettings,omitnil,omitempty" name:"InstanceNameIndexSettings"`
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
	delete(f, "InstanceNameIndexSettings")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAutoScalingGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAutoScalingGroupResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// InstanceType verification policy, which is effective when actual modification is made to InstanceTypes. Valid values include ALL and ANY and the default value is ANY.
	// <li>ALL: Verification passes if all InstanceTypes are available; otherwise, a verification error will be reported.</li>
	// <li>ANY: Verification passes if any InstanceType is available; otherwise, a verification error will be reported.</li>
	// Common reasons for unavailable InstanceTypes include the InstanceType being sold out, and the corresponding cloud disk being sold out.
	// If a model in InstanceTypes does not exist or has been abolished, a verification error will be reported regardless of the valid values set for InstanceTypesCheckPolicy.
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
	// <li>POSTPAID_BY_HOUR: pay-as-you-go hourly</li>
	// <li>SPOTPAID: spot instance</li>
	// <li> CDCPAID: dedicated cluster</li>
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

	// Cloud disk type selection policy. Valid values:
	// <li>ORIGINAL: Use the set cloud disk type.</li>
	// <li>AUTOMATIC: Automatically select the currently available cloud disk type.</li>
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

	// Instance login settings, which include passwords, keys, or the original login settings inherited from the image. <br>Please note that specifying new login settings will overwrite the existing ones. For instance, if you previously used a password for login and then use this parameter to switch the login settings to a key, the original password will be removed.
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// Instance tag list. By specifying this parameter, the instances added through scale-out can be bound to the tag. Up to 10 Tags can be specified.
	// This parameter will overwrite the original instance tag list. To add new tags, you need to pass the new tags along with the original tags.
	InstanceTags []*InstanceTag `json:"InstanceTags,omitnil,omitempty" name:"InstanceTags"`

	// Image family name.
	ImageFamily *string `json:"ImageFamily,omitnil,omitempty" name:"ImageFamily"`

	// Cloud Dedicated Cluster (CDC) ID.
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`

	// Custom metadata.
	Metadata *Metadata `json:"Metadata,omitnil,omitempty" name:"Metadata"`
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

	// InstanceType verification policy, which is effective when actual modification is made to InstanceTypes. Valid values include ALL and ANY and the default value is ANY.
	// <li>ALL: Verification passes if all InstanceTypes are available; otherwise, a verification error will be reported.</li>
	// <li>ANY: Verification passes if any InstanceType is available; otherwise, a verification error will be reported.</li>
	// Common reasons for unavailable InstanceTypes include the InstanceType being sold out, and the corresponding cloud disk being sold out.
	// If a model in InstanceTypes does not exist or has been abolished, a verification error will be reported regardless of the valid values set for InstanceTypesCheckPolicy.
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
	// <li>POSTPAID_BY_HOUR: pay-as-you-go hourly</li>
	// <li>SPOTPAID: spot instance</li>
	// <li> CDCPAID: dedicated cluster</li>
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

	// Cloud disk type selection policy. Valid values:
	// <li>ORIGINAL: Use the set cloud disk type.</li>
	// <li>AUTOMATIC: Automatically select the currently available cloud disk type.</li>
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

	// Instance login settings, which include passwords, keys, or the original login settings inherited from the image. <br>Please note that specifying new login settings will overwrite the existing ones. For instance, if you previously used a password for login and then use this parameter to switch the login settings to a key, the original password will be removed.
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// Instance tag list. By specifying this parameter, the instances added through scale-out can be bound to the tag. Up to 10 Tags can be specified.
	// This parameter will overwrite the original instance tag list. To add new tags, you need to pass the new tags along with the original tags.
	InstanceTags []*InstanceTag `json:"InstanceTags,omitnil,omitempty" name:"InstanceTags"`

	// Image family name.
	ImageFamily *string `json:"ImageFamily,omitnil,omitempty" name:"ImageFamily"`

	// Cloud Dedicated Cluster (CDC) ID.
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`

	// Custom metadata.
	Metadata *Metadata `json:"Metadata,omitnil,omitempty" name:"Metadata"`
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
	delete(f, "LoginSettings")
	delete(f, "InstanceTags")
	delete(f, "ImageFamily")
	delete(f, "DedicatedClusterId")
	delete(f, "Metadata")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLaunchConfigurationAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLaunchConfigurationAttributesResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// Scenario for entering the lifecycle hook. Valid values:
	// <li>INSTANCE_LAUNCHING: after the instance is launched.</li>
	// <li>INSTANCE_TERMINATING: before the instance is terminated.</li>
	LifecycleTransition *string `json:"LifecycleTransition,omitnil,omitempty" name:"LifecycleTransition"`

	// Action to be taken by the scaling group in case of lifecycle hook timeout. Valid values:
	// <li>CONTINUE: Continue the scaling activity after timeout.</li>
	// <li>ABANDON: Terminate the scaling activity after timeout.</li>
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

	// Scenario for entering the lifecycle hook. Valid values:
	// <li>INSTANCE_LAUNCHING: after the instance is launched.</li>
	// <li>INSTANCE_TERMINATING: before the instance is terminated.</li>
	LifecycleTransition *string `json:"LifecycleTransition,omitnil,omitempty" name:"LifecycleTransition"`

	// Action to be taken by the scaling group in case of lifecycle hook timeout. Valid values:
	// <li>CONTINUE: Continue the scaling activity after timeout.</li>
	// <li>ABANDON: Terminate the scaling activity after timeout.</li>
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
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// CLB verification policy. Valid values: ALL and DIFF. Default value: ALL.
	// <li>ALL: The CLB passes the verification only when all CLB parameters are valid. Otherwise, a verification error occurs.</li>
	// <li>DIFF: The CLB passes the verification only when the CLB parameters with changes are valid. Otherwise, a verification error occurs.</li>
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

	// CLB verification policy. Valid values: ALL and DIFF. Default value: ALL.
	// <li>ALL: The CLB passes the verification only when all CLB parameters are valid. Otherwise, a verification error occurs.</li>
	// <li>DIFF: The CLB passes the verification only when the CLB parameters with changes are valid. Otherwise, a verification error occurs.</li>
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// The adjustment value for the expected number of instances after an alarm is triggered. It applies only to simple policies. <li>When AdjustmentType is CHANGE_IN_CAPACITY, a positive AdjustmentValue indicates an increase in the number of instances after the alarm is triggered, and a negative AdjustmentValue indicates a decrease in the number of instances after the alarm is triggered.</li> <li>When AdjustmentType is EXACT_CAPACITY, the value of AdjustmentValue represents the expected number of instances after the alarm is triggered, which should be greater than or equal to 0.</li> <li>When AdjustmentType is PERCENT_CHANGE_IN_CAPACITY, a positive AdjustmentValue indicates an increase in the number of instances by percentage after the alarm is triggered, and a negative AdjustmentValue indicates a decrease in the number of instances by percentage after the alarm is triggered. The unit is: %.</li>
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

	// The adjustment value for the expected number of instances after an alarm is triggered. It applies only to simple policies. <li>When AdjustmentType is CHANGE_IN_CAPACITY, a positive AdjustmentValue indicates an increase in the number of instances after the alarm is triggered, and a negative AdjustmentValue indicates a decrease in the number of instances after the alarm is triggered.</li> <li>When AdjustmentType is EXACT_CAPACITY, the value of AdjustmentValue represents the expected number of instances after the alarm is triggered, which should be greater than or equal to 0.</li> <li>When AdjustmentType is PERCENT_CHANGE_IN_CAPACITY, a positive AdjustmentValue indicates an increase in the number of instances by percentage after the alarm is triggered, and a negative AdjustmentValue indicates a decrease in the number of instances by percentage after the alarm is triggered. The unit is: %.</li>
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
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

type RefreshActivity struct {
	// Scaling group ID.
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// Refresh activity ID.
	RefreshActivityId *string `json:"RefreshActivityId,omitnil,omitempty" name:"RefreshActivityId"`

	// Original refresh activity ID, which exists only in the rollback refresh activity.
	// Note: This field may return null, indicating that no valid value can be obtained.
	OriginRefreshActivityId *string `json:"OriginRefreshActivityId,omitnil,omitempty" name:"OriginRefreshActivityId"`

	// Refresh batch information list.
	RefreshBatchSet []*RefreshBatch `json:"RefreshBatchSet,omitnil,omitempty" name:"RefreshBatchSet"`

	// Refresh mode.
	RefreshMode *string `json:"RefreshMode,omitnil,omitempty" name:"RefreshMode"`

	// Instance update setting parameters.
	RefreshSettings *RefreshSettings `json:"RefreshSettings,omitnil,omitempty" name:"RefreshSettings"`

	// Refresh activity type. Valid values:
	// <li>NORMAL: normal refresh activity.</li>
	// <li>ROLLBACK: rollback refresh activity.</li>
	ActivityType *string `json:"ActivityType,omitnil,omitempty" name:"ActivityType"`

	// Refresh activity status. Valid values:
	// <li>INIT: initializing.</li>
	// <li>RUNNING: running.</li>
	// <li>SUCCESSFUL: successful.</li>
	// <li>FAILED_PAUSE: paused due to the failure of a refresh batch.</li>
	// <li>AUTO_PAUSE: automatically paused due to the pause policy.</li>
	// <li>MANUAL_PAUSE: manually paused.</li>
	// <li>CANCELLED: canceled.</li>
	// <li>FAILED: failed.</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Current refresh batch number. For example, a value of 2 indicates that the current activity is refreshing the second batch of instances.
	// Note: This field may return null, indicating that no valid value can be obtained.
	CurrentRefreshBatchNum *uint64 `json:"CurrentRefreshBatchNum,omitnil,omitempty" name:"CurrentRefreshBatchNum"`

	// Refresh activity start time.
	// Note: This field may return null, indicating that no valid value can be obtained.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Refresh activity end time.
	// Note: This field may return null, indicating that no valid value can be obtained.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Refresh activity creation time.
	// Note: This field may return null, indicating that no valid value can be obtained.
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`
}

type RefreshBatch struct {
	// Refresh batch number. For example, a value of 2 indicates that the current batch of instances will be refreshed in the second batch.
	RefreshBatchNum *uint64 `json:"RefreshBatchNum,omitnil,omitempty" name:"RefreshBatchNum"`

	// Refresh batch status. Valid values: <li>WAITING: pending refresh;</li> <li>INIT: initializing;</li> <li>RUNNING: refreshing;</li> <li>FAILED: refresh failed;</li> <li>PARTIALLY_SUCCESSFUL: partially successful in the batch;</li> <li>CANCELLED: cancelled;</li> <li>SUCCESSFUL: refresh successful.</li>
	RefreshBatchStatus *string `json:"RefreshBatchStatus,omitnil,omitempty" name:"RefreshBatchStatus"`

	// List of instances linked to a refresh batch.
	RefreshBatchRelatedInstanceSet []*RefreshBatchRelatedInstance `json:"RefreshBatchRelatedInstanceSet,omitnil,omitempty" name:"RefreshBatchRelatedInstanceSet"`

	// Refresh batch start time.
	// Note: This field may return null, indicating that no valid value can be obtained.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Refresh batch end time.
	// Note: This field may return null, indicating that no valid value can be obtained.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type RefreshBatchRelatedInstance struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Refresh instance status. If an instance is removed or destroyed during the refresh process, its status will be updated to NOT_FOUND. Valid values: <br><li>WAITING: pending refresh.</li> <li>INIT: Initializing.</li> <li>RUNNING: Refreshing in progress.</li> <li>FAILED: Refresh failed.</li> <li>CANCELLED: Canceled.</li> <li>SUCCESSFUL: Refreshed.</li> <li>NOT_FOUND: Instance not found.
	InstanceStatus *string `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// The most recent scaling activity ID during instance refresh can be queried via the DescribeAutoScalingActivities API.
	// Please note that scaling activities differ from instance refresh activities; a single instance refresh activity may involve multiple scaling activities.
	// Note: This field may return null, indicating that no valid value can be obtained.
	LastActivityId *string `json:"LastActivityId,omitnil,omitempty" name:"LastActivityId"`

	// Instance refresh status information.
	// Note: This field may return null, indicating that no valid value can be obtained.
	InstanceStatusMessage *string `json:"InstanceStatusMessage,omitnil,omitempty" name:"InstanceStatusMessage"`
}

type RefreshSettings struct {
	// Rolling update settings parameters. RefreshMode is the rolling update. This parameter must be filled in.Note: This field may return null, indicating that no valid value can be obtained.
	RollingUpdateSettings *RollingUpdateSettings `json:"RollingUpdateSettings,omitnil,omitempty" name:"RollingUpdateSettings"`

	// Whether to enable the backend service health check for the instance. Default value: FALSE. This parameter is valid only for the scaling group bound to an application-based CLB. After this feature is enabled, if the instance fails the check after refresh, the port weight of the CLB will be always 0, and it will be marked as a refresh failure. Valid values: <li>TRUE: enable;</li> <li>FALSE: disable.</li>
	CheckInstanceTargetHealth *bool `json:"CheckInstanceTargetHealth,omitnil,omitempty" name:"CheckInstanceTargetHealth"`
}

type RelatedInstance struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Status of the instance in the scaling activity. Valid values:
	// `INIT`: Initializing
	// `RUNNING`: Processing the instance
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

// Predefined struct for user
type ResumeInstanceRefreshRequestParams struct {
	// Scaling group ID.
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// Refresh activity ID.
	RefreshActivityId *string `json:"RefreshActivityId,omitnil,omitempty" name:"RefreshActivityId"`

	// Recovery mode of instances that have failed to be refreshed in the current batch. If there are no failed instances, this parameter is invalid. Default value: RETRY. Valid values: <li>RETRY: Retry instances that have failed to be refreshed in the current batch.</li> <li>CONTINUE: Skip instances that have failed to be refreshed in the current batch.</li>
	ResumeMode *string `json:"ResumeMode,omitnil,omitempty" name:"ResumeMode"`
}

type ResumeInstanceRefreshRequest struct {
	*tchttp.BaseRequest
	
	// Scaling group ID.
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// Refresh activity ID.
	RefreshActivityId *string `json:"RefreshActivityId,omitnil,omitempty" name:"RefreshActivityId"`

	// Recovery mode of instances that have failed to be refreshed in the current batch. If there are no failed instances, this parameter is invalid. Default value: RETRY. Valid values: <li>RETRY: Retry instances that have failed to be refreshed in the current batch.</li> <li>CONTINUE: Skip instances that have failed to be refreshed in the current batch.</li>
	ResumeMode *string `json:"ResumeMode,omitnil,omitempty" name:"ResumeMode"`
}

func (r *ResumeInstanceRefreshRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeInstanceRefreshRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "RefreshActivityId")
	delete(f, "ResumeMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResumeInstanceRefreshRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeInstanceRefreshResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResumeInstanceRefreshResponse struct {
	*tchttp.BaseResponse
	Response *ResumeInstanceRefreshResponseParams `json:"Response"`
}

func (r *ResumeInstanceRefreshResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeInstanceRefreshResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RollbackInstanceRefreshRequestParams struct {
	// Scaling group ID.
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// Refresh settings.
	RefreshSettings *RefreshSettings `json:"RefreshSettings,omitnil,omitempty" name:"RefreshSettings"`

	// Original refresh activity ID.
	OriginRefreshActivityId *string `json:"OriginRefreshActivityId,omitnil,omitempty" name:"OriginRefreshActivityId"`

	// Refresh mode, currently, only rolling updates are supported, with the default value being ROLLING_UPDATE_RESET.
	RefreshMode *string `json:"RefreshMode,omitnil,omitempty" name:"RefreshMode"`
}

type RollbackInstanceRefreshRequest struct {
	*tchttp.BaseRequest
	
	// Scaling group ID.
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// Refresh settings.
	RefreshSettings *RefreshSettings `json:"RefreshSettings,omitnil,omitempty" name:"RefreshSettings"`

	// Original refresh activity ID.
	OriginRefreshActivityId *string `json:"OriginRefreshActivityId,omitnil,omitempty" name:"OriginRefreshActivityId"`

	// Refresh mode, currently, only rolling updates are supported, with the default value being ROLLING_UPDATE_RESET.
	RefreshMode *string `json:"RefreshMode,omitnil,omitempty" name:"RefreshMode"`
}

func (r *RollbackInstanceRefreshRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollbackInstanceRefreshRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "RefreshSettings")
	delete(f, "OriginRefreshActivityId")
	delete(f, "RefreshMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RollbackInstanceRefreshRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RollbackInstanceRefreshResponseParams struct {
	// Refresh activity ID.
	RefreshActivityId *string `json:"RefreshActivityId,omitnil,omitempty" name:"RefreshActivityId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RollbackInstanceRefreshResponse struct {
	*tchttp.BaseResponse
	Response *RollbackInstanceRefreshResponseParams `json:"Response"`
}

func (r *RollbackInstanceRefreshResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollbackInstanceRefreshResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RollingUpdateSettings struct {
	// Batch quantity. The batch quantity should be a positive integer greater than 0, but cannot exceed the total number of instances pending refresh.
	BatchNumber *uint64 `json:"BatchNumber,omitnil,omitempty" name:"BatchNumber"`

	// Pause policy between batches. Default value: Automatic. Valid values:
	// <li>FIRST_BATCH_PAUSE: Pause after the first batch of updates is completed.</li>
	// <li>BATCH_INTERVAL_PAUSE: Pause between batches.</li>
	// <li>AUTOMATIC: Do not pause.</li>
	BatchPause *string `json:"BatchPause,omitnil,omitempty" name:"BatchPause"`

	// The maximum additional quantity of instances. After this parameter is set, create a batch of additional pay-as-you-go instances according to the launch configuration before the rolling update starts. After the rolling update is completed, the additional instances will be terminated.This parameter is used to ensure a certain number of instances available during the rolling update. The maximum additional quantity of instances cannot exceed the number of refreshing instances in a single batch of the rolling update. The rollback process does not support this parameter currently.
	MaxSurge *int64 `json:"MaxSurge,omitnil,omitempty" name:"MaxSurge"`
}

type RunAutomationServiceEnabled struct {
	// Whether to enable [TencentCloud Automation Tools](https://intl.cloud.tencent.com/document/product/1340?from_cn_redirect=1). Valid values:<br><li>`TRUE`: Enable<br><li>`FALSE`: Not enable.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type RunMonitorServiceEnabled struct {
	// Whether to enable [Tencent Cloud Observability Platform (formerly Cloud Monitor)](https://intl.cloud.tencent.com/document/product/248?from_cn_redirect=1) (TCOP). Valid values:
	// <li>TRUE: enable TCOP.</li>
	// <li>FALSE: disable TCOP.</li>
	// Default value: TRUE.
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type RunSecurityServiceEnabled struct {
	// Whether to enable [Cloud Workload Protection Platform (CWPP)](https://intl.cloud.tencent.com/document/product/296?from_cn_redirect=1). Valid values:
	// <li>TRUE: enable CWPP.</li>
	// <li>FALSE: disable CWPP.</li>
	// Default value: TRUE.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// Scheduled task execution type. Valid values:
	// <li>CRONTAB: repeated execution.</li>
	// <li>ONCE: single execution.</li>
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

	// Replace mode of unhealthy replacement service. Valid values:
	// RECREATE: Rebuild an instance to replace the original unhealthy instance.
	// RESET: Performing a system reinstallation on unhealthy instances to keep information such as data disks, private IP addresses, and instance IDs unchanged. The instance login settings, HostName, enhanced services, and UserData will remain consistent with the current launch configuration.
	// Default value: RECREATE.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ReplaceMode *string `json:"ReplaceMode,omitnil,omitempty" name:"ReplaceMode"`

	// Automatic instance tag update. The default value is false. If this feature is enabled, tags of running instances in a scaling group will be updated as well if the scaling group tags are updated. (This feature takes effect for tag creation and editing but not tag deletion.) The update does not take effect immediately due to certain latency.
	AutoUpdateInstanceTags *bool `json:"AutoUpdateInstanceTags,omitnil,omitempty" name:"AutoUpdateInstanceTags"`
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
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// Spot instance type. The value can only be one-time currently. Default value: one-time.
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
type StartInstanceRefreshRequestParams struct {
	// Scaling group ID.
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// Refresh settings.
	RefreshSettings *RefreshSettings `json:"RefreshSettings,omitnil,omitempty" name:"RefreshSettings"`

	// Refresh mode. Valid values:
	// <li>ROLLING_UPDATE_RESET: Reinstall the system for rolling updates.</li>
	// <li>ROLLING_UPDATE_REPLACE: Create an instance and replace the old instance with it for rolling updates. This mode does not support the rollback API currently.</li>
	RefreshMode *string `json:"RefreshMode,omitnil,omitempty" name:"RefreshMode"`
}

type StartInstanceRefreshRequest struct {
	*tchttp.BaseRequest
	
	// Scaling group ID.
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// Refresh settings.
	RefreshSettings *RefreshSettings `json:"RefreshSettings,omitnil,omitempty" name:"RefreshSettings"`

	// Refresh mode. Valid values:
	// <li>ROLLING_UPDATE_RESET: Reinstall the system for rolling updates.</li>
	// <li>ROLLING_UPDATE_REPLACE: Create an instance and replace the old instance with it for rolling updates. This mode does not support the rollback API currently.</li>
	RefreshMode *string `json:"RefreshMode,omitnil,omitempty" name:"RefreshMode"`
}

func (r *StartInstanceRefreshRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartInstanceRefreshRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "RefreshSettings")
	delete(f, "RefreshMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartInstanceRefreshRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartInstanceRefreshResponseParams struct {
	// Refresh activity ID.
	RefreshActivityId *string `json:"RefreshActivityId,omitnil,omitempty" name:"RefreshActivityId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartInstanceRefreshResponse struct {
	*tchttp.BaseResponse
	Response *StartInstanceRefreshResponseParams `json:"Response"`
}

func (r *StartInstanceRefreshResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartInstanceRefreshResponse) FromJsonString(s string) error {
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

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

// Predefined struct for user
type StopInstanceRefreshRequestParams struct {
	// Scaling group ID.
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// Refresh activity ID.
	RefreshActivityId *string `json:"RefreshActivityId,omitnil,omitempty" name:"RefreshActivityId"`
}

type StopInstanceRefreshRequest struct {
	*tchttp.BaseRequest
	
	// Scaling group ID.
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// Refresh activity ID.
	RefreshActivityId *string `json:"RefreshActivityId,omitnil,omitempty" name:"RefreshActivityId"`
}

func (r *StopInstanceRefreshRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopInstanceRefreshRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "RefreshActivityId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopInstanceRefreshRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopInstanceRefreshResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopInstanceRefreshResponse struct {
	*tchttp.BaseResponse
	Response *StopInstanceRefreshResponseParams `json:"Response"`
}

func (r *StopInstanceRefreshResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopInstanceRefreshResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SystemDisk struct {
	// System disk type. For restrictions on the system disk type, see [cloud block storage types](https://intl.cloud.tencent.com/document/product/362/2353?from_cn_redirect=1). Valid values:
	// <li>LOCAL_BASIC: Local Disk.</li>
	// <li>LOCAL_SSD: Local SSD.</li>
	// <li>CLOUD_BASIC: Basic Cloud Disk.</li>
	// <li>CLOUD_PREMIUM: Premium Disk.</li>
	// <li>CLOUD_SSD: Cloud SSD.</li>
	// <li>Default value: CLOUD_PREMIUM.</li>
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// System disk size, in GB. Default value: 50.
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`
}

type Tag struct {
	// Tag key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Tag value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// Resource type bound with tags. Valid values: auto-scaling-group and launch-configuration.
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

	// This parameter is now invalid and should not be used. Upgrading the launch configuration API does not allow modification or overwriting of the LoginSettings parameter. LoginSettings will not change after upgrade.
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

	// This parameter is now invalid and should not be used. Upgrading the launch configuration API does not allow modification or overwriting of the LoginSettings parameter. LoginSettings will not change after upgrade.
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
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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