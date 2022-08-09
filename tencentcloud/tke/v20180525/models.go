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

package v20180525

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type AcquireClusterAdminRoleRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type AcquireClusterAdminRoleRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *AcquireClusterAdminRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AcquireClusterAdminRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AcquireClusterAdminRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AcquireClusterAdminRoleResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AcquireClusterAdminRoleResponse struct {
	*tchttp.BaseResponse
	Response *AcquireClusterAdminRoleResponseParams `json:"Response"`
}

func (r *AcquireClusterAdminRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AcquireClusterAdminRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddExistedInstancesRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Instance list. Spot instance is not supported.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Detailed information of the instance
	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitempty" name:"InstanceAdvancedSettings"`

	// Enhanced services. This parameter is used to specify whether to enable Cloud Security, Cloud Monitoring and other services. If this parameter is not specified, Cloud Monitor and Cloud Security are enabled by default.
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`

	// Node login information (currently only supports using Password or single KeyIds)
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// When reinstalling the system, you can specify the HostName of the modified instance (when the cluster is in HostName mode, this parameter is required, and the rule name is the same as the [Create CVM Instance](https://intl.cloud.tencent.com/document/product/213/15730?from_cn_redirect=1) API HostName except for uppercase letters not being supported.
	HostName *string `json:"HostName,omitempty" name:"HostName"`

	// Security group to which the instance belongs. This parameter can be obtained from the `sgId` field returned by DescribeSecurityGroups. If this parameter is not specified, the default security group is bound. (Currently, you can only set a single sgId)
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// Node pool options
	NodePool *NodePoolOption `json:"NodePool,omitempty" name:"NodePool"`

	// Skips the specified verification. Valid values: GlobalRouteCIDRCheck, VpcCniCIDRCheck
	SkipValidateOptions []*string `json:"SkipValidateOptions,omitempty" name:"SkipValidateOptions"`

	// This parameter is used to customize the configuration of an instance, which corresponds to the `InstanceIds` one-to-one in sequence. If this parameter is passed in, the default parameter `InstanceAdvancedSettings` will be overwritten and will not take effect. If this parameter is not passed in, the `InstanceAdvancedSettings` will take effect for each instance.
	// 
	// The array length of `InstanceAdvancedSettingsOverride` should be the same as the array length of `InstanceIds`. If its array length is greater than the `InstanceIds` array length, an error will be reported. If its array length is less than the `InstanceIds` array length, the instance without corresponding configuration will use the default configuration.
	InstanceAdvancedSettingsOverrides []*InstanceAdvancedSettings `json:"InstanceAdvancedSettingsOverrides,omitempty" name:"InstanceAdvancedSettingsOverrides"`

	// Node image (it is required when creating a node)
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
}

type AddExistedInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Instance list. Spot instance is not supported.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Detailed information of the instance
	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitempty" name:"InstanceAdvancedSettings"`

	// Enhanced services. This parameter is used to specify whether to enable Cloud Security, Cloud Monitoring and other services. If this parameter is not specified, Cloud Monitor and Cloud Security are enabled by default.
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`

	// Node login information (currently only supports using Password or single KeyIds)
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// When reinstalling the system, you can specify the HostName of the modified instance (when the cluster is in HostName mode, this parameter is required, and the rule name is the same as the [Create CVM Instance](https://intl.cloud.tencent.com/document/product/213/15730?from_cn_redirect=1) API HostName except for uppercase letters not being supported.
	HostName *string `json:"HostName,omitempty" name:"HostName"`

	// Security group to which the instance belongs. This parameter can be obtained from the `sgId` field returned by DescribeSecurityGroups. If this parameter is not specified, the default security group is bound. (Currently, you can only set a single sgId)
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// Node pool options
	NodePool *NodePoolOption `json:"NodePool,omitempty" name:"NodePool"`

	// Skips the specified verification. Valid values: GlobalRouteCIDRCheck, VpcCniCIDRCheck
	SkipValidateOptions []*string `json:"SkipValidateOptions,omitempty" name:"SkipValidateOptions"`

	// This parameter is used to customize the configuration of an instance, which corresponds to the `InstanceIds` one-to-one in sequence. If this parameter is passed in, the default parameter `InstanceAdvancedSettings` will be overwritten and will not take effect. If this parameter is not passed in, the `InstanceAdvancedSettings` will take effect for each instance.
	// 
	// The array length of `InstanceAdvancedSettingsOverride` should be the same as the array length of `InstanceIds`. If its array length is greater than the `InstanceIds` array length, an error will be reported. If its array length is less than the `InstanceIds` array length, the instance without corresponding configuration will use the default configuration.
	InstanceAdvancedSettingsOverrides []*InstanceAdvancedSettings `json:"InstanceAdvancedSettingsOverrides,omitempty" name:"InstanceAdvancedSettingsOverrides"`

	// Node image (it is required when creating a node)
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
}

func (r *AddExistedInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddExistedInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceIds")
	delete(f, "InstanceAdvancedSettings")
	delete(f, "EnhancedService")
	delete(f, "LoginSettings")
	delete(f, "HostName")
	delete(f, "SecurityGroupIds")
	delete(f, "NodePool")
	delete(f, "SkipValidateOptions")
	delete(f, "InstanceAdvancedSettingsOverrides")
	delete(f, "ImageId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddExistedInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddExistedInstancesResponseParams struct {
	// IDs of failed nodes
	// Note: This field may return null, indicating that no valid value was found.
	FailedInstanceIds []*string `json:"FailedInstanceIds,omitempty" name:"FailedInstanceIds"`

	// IDs of successful nodes
	// Note: This field may return null, indicating that no valid value was found.
	SuccInstanceIds []*string `json:"SuccInstanceIds,omitempty" name:"SuccInstanceIds"`

	// IDs of (successful or failed) nodes that timed out
	// Note: This field may return null, indicating that no valid value was found.
	TimeoutInstanceIds []*string `json:"TimeoutInstanceIds,omitempty" name:"TimeoutInstanceIds"`

	// Causes of the failure to add a node to a cluster
	// Note: this field may return `null`, indicating that no valid value is obtained.
	FailedReasons []*string `json:"FailedReasons,omitempty" name:"FailedReasons"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddExistedInstancesResponse struct {
	*tchttp.BaseResponse
	Response *AddExistedInstancesResponseParams `json:"Response"`
}

func (r *AddExistedInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddExistedInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddNodeToNodePoolRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Node pool ID
	NodePoolId *string `json:"NodePoolId,omitempty" name:"NodePoolId"`

	// Node ID
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type AddNodeToNodePoolRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Node pool ID
	NodePoolId *string `json:"NodePoolId,omitempty" name:"NodePoolId"`

	// Node ID
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *AddNodeToNodePoolRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddNodeToNodePoolRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NodePoolId")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddNodeToNodePoolRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddNodeToNodePoolResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddNodeToNodePoolResponse struct {
	*tchttp.BaseResponse
	Response *AddNodeToNodePoolResponseParams `json:"Response"`
}

func (r *AddNodeToNodePoolResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddNodeToNodePoolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddVpcCniSubnetsRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The subnets added for the cluster container network
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// ID of the VPC where the cluster resides
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

type AddVpcCniSubnetsRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The subnets added for the cluster container network
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// ID of the VPC where the cluster resides
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

func (r *AddVpcCniSubnetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddVpcCniSubnetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SubnetIds")
	delete(f, "VpcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddVpcCniSubnetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddVpcCniSubnetsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddVpcCniSubnetsResponse struct {
	*tchttp.BaseResponse
	Response *AddVpcCniSubnetsResponseParams `json:"Response"`
}

func (r *AddVpcCniSubnetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddVpcCniSubnetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AutoScalingGroupRange struct {
	// Minimum number of pods in a scaling group
	MinSize *int64 `json:"MinSize,omitempty" name:"MinSize"`

	// Maximum number of pods in a scaling group
	MaxSize *int64 `json:"MaxSize,omitempty" name:"MaxSize"`
}

type AutoUpgradeClusterLevel struct {
	// Whether to enable Auto Cluster Upgrade
	IsAutoUpgrade *bool `json:"IsAutoUpgrade,omitempty" name:"IsAutoUpgrade"`
}

type AutoscalingAdded struct {
	// Number of nodes that are being added
	Joining *int64 `json:"Joining,omitempty" name:"Joining"`

	// Number of nodes that are being initialized
	Initializing *int64 `json:"Initializing,omitempty" name:"Initializing"`

	// Number of normal nodes
	Normal *int64 `json:"Normal,omitempty" name:"Normal"`

	// Total number of nodes
	Total *int64 `json:"Total,omitempty" name:"Total"`
}

// Predefined struct for user
type CheckEdgeClusterCIDRRequestParams struct {
	// Cluster VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Cluster Pod CIDR block
	PodCIDR *string `json:"PodCIDR,omitempty" name:"PodCIDR"`

	// Cluster service CIDR block
	ServiceCIDR *string `json:"ServiceCIDR,omitempty" name:"ServiceCIDR"`
}

type CheckEdgeClusterCIDRRequest struct {
	*tchttp.BaseRequest
	
	// Cluster VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Cluster Pod CIDR block
	PodCIDR *string `json:"PodCIDR,omitempty" name:"PodCIDR"`

	// Cluster service CIDR block
	ServiceCIDR *string `json:"ServiceCIDR,omitempty" name:"ServiceCIDR"`
}

func (r *CheckEdgeClusterCIDRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckEdgeClusterCIDRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "PodCIDR")
	delete(f, "ServiceCIDR")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckEdgeClusterCIDRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckEdgeClusterCIDRResponseParams struct {
	// Return code. Valid values:
	// -1: Internal error
	// 0: No conflict
	// 1: Conflict between VPC and serviceCIDR
	// 2: Conflict between VPC and podCIDR
	// 3: Conflict between serviceCIDR and podCIDR
	ConflictCode *int64 `json:"ConflictCode,omitempty" name:"ConflictCode"`

	// CIDR block conflict description
	ConflictMsg *string `json:"ConflictMsg,omitempty" name:"ConflictMsg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CheckEdgeClusterCIDRResponse struct {
	*tchttp.BaseResponse
	Response *CheckEdgeClusterCIDRResponseParams `json:"Response"`
}

func (r *CheckEdgeClusterCIDRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckEdgeClusterCIDRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckInstancesUpgradeAbleRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Specify the node list to check. If it’s not passed in, all nodes of the cluster will be checked.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Upgrade type
	UpgradeType *string `json:"UpgradeType,omitempty" name:"UpgradeType"`

	// Pagination offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Pagination limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Filtering
	Filter []*Filter `json:"Filter,omitempty" name:"Filter"`
}

type CheckInstancesUpgradeAbleRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Specify the node list to check. If it’s not passed in, all nodes of the cluster will be checked.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Upgrade type
	UpgradeType *string `json:"UpgradeType,omitempty" name:"UpgradeType"`

	// Pagination offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Pagination limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Filtering
	Filter []*Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *CheckInstancesUpgradeAbleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckInstancesUpgradeAbleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceIds")
	delete(f, "UpgradeType")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckInstancesUpgradeAbleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckInstancesUpgradeAbleResponseParams struct {
	// The current minor version of cluster Master
	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`

	// The latest minor version of cluster Master corresponding major version
	LatestVersion *string `json:"LatestVersion,omitempty" name:"LatestVersion"`

	// List of nodes that can be upgraded
	// Note: this field may return `null`, indicating that no valid value is obtained.
	UpgradeAbleInstances []*UpgradeAbleInstancesItem `json:"UpgradeAbleInstances,omitempty" name:"UpgradeAbleInstances"`

	// Total number
	// Note: this field may return `null`, indicating that no valid value is obtained.
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CheckInstancesUpgradeAbleResponse struct {
	*tchttp.BaseResponse
	Response *CheckInstancesUpgradeAbleResponseParams `json:"Response"`
}

func (r *CheckInstancesUpgradeAbleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckInstancesUpgradeAbleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Cluster struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Cluster name
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// Cluster description
	ClusterDescription *string `json:"ClusterDescription,omitempty" name:"ClusterDescription"`

	// Cluster version. The default value is 1.10.5.
	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`

	// Cluster operating system. centOS 7.2x86_64 or ubuntu 16.04.1 LTSx86_64. Default value: ubuntu 16.04.1 LTSx86_64
	ClusterOs *string `json:"ClusterOs,omitempty" name:"ClusterOs"`

	// Cluster type. Managed cluster: MANAGED_CLUSTER; Self-deployed cluster: INDEPENDENT_CLUSTER.
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// Cluster network-related parameters
	ClusterNetworkSettings *ClusterNetworkSettings `json:"ClusterNetworkSettings,omitempty" name:"ClusterNetworkSettings"`

	// Current number of nodes in the cluster
	ClusterNodeNum *uint64 `json:"ClusterNodeNum,omitempty" name:"ClusterNodeNum"`

	// ID of the project to which the cluster belongs
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Tag description list.
	TagSpecification []*TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification"`

	// Cluster status (`Running`, `Creating`, `Idling` or `Abnormal`)
	ClusterStatus *string `json:"ClusterStatus,omitempty" name:"ClusterStatus"`

	// Cluster attributes (including a map of different cluster attributes, with attribute fields including NodeNameType (lan-ip mode and hostname mode, with lan-ip mode as default))
	Property *string `json:"Property,omitempty" name:"Property"`

	// Number of primary nodes currently in the cluster
	ClusterMaterNodeNum *uint64 `json:"ClusterMaterNodeNum,omitempty" name:"ClusterMaterNodeNum"`

	// ID of the image used by the cluster
	// Note: this field may return null, indicating that no valid value is obtained.
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// Container image tag
	// Note: This field may return null, indicating that no valid value was found.
	OsCustomizeType *string `json:"OsCustomizeType,omitempty" name:"OsCustomizeType"`

	// Runtime environment of the cluster. Values can be `docker` or `containerd`.
	// Note: this field may return null, indicating that no valid value is obtained.
	ContainerRuntime *string `json:"ContainerRuntime,omitempty" name:"ContainerRuntime"`

	// Creation time
	// Note: this field may return null, indicating that no valid value is obtained.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Whether Deletion Protection is enabled
	// Note: this field may return null, indicating that no valid value is obtained.
	DeletionProtection *bool `json:"DeletionProtection,omitempty" name:"DeletionProtection"`

	// Specifies whether the cluster supports external nodes.
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	EnableExternalNode *bool `json:"EnableExternalNode,omitempty" name:"EnableExternalNode"`

	// Cluster models. It’s valid for managed clusters.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ClusterLevel *string `json:"ClusterLevel,omitempty" name:"ClusterLevel"`

	// The target cluster model for auto-upgrade
	// Note: this field may return null, indicating that no valid value is obtained.
	AutoUpgradeClusterLevel *bool `json:"AutoUpgradeClusterLevel,omitempty" name:"AutoUpgradeClusterLevel"`

	// Whether to enable qGPU Sharing
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	QGPUShareEnable *bool `json:"QGPUShareEnable,omitempty" name:"QGPUShareEnable"`
}

type ClusterAdvancedSettings struct {
	// Whether IPVS is enabled
	IPVS *bool `json:"IPVS,omitempty" name:"IPVS"`

	// Whether auto-scaling is enabled for nodes in the cluster (Enabling this function is not supported when you create a cluster)
	AsEnabled *bool `json:"AsEnabled,omitempty" name:"AsEnabled"`

	// Type of runtime component used by the cluster. The types include "docker" and "containerd". Default value: docker
	ContainerRuntime *string `json:"ContainerRuntime,omitempty" name:"ContainerRuntime"`

	// NodeName type for a node in a cluster (This includes the two forms of **hostname** and **lan-ip**, with the default as **lan-ip**. If **hostname** is used, you need to set the HostName parameter when creating a node, and the InstanceName needs to be the same as the HostName.)
	NodeNameType *string `json:"NodeNameType,omitempty" name:"NodeNameType"`

	// Cluster custom parameter
	ExtraArgs *ClusterExtraArgs `json:"ExtraArgs,omitempty" name:"ExtraArgs"`

	// Cluster network type, which can be GR (Global Router) or VPC-CNI. The default value is GR.
	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`

	// Whether a cluster in VPC-CNI mode uses dynamic IP addresses. The default value is FALSE, which indicates that static IP addresses are used.
	IsNonStaticIpMode *bool `json:"IsNonStaticIpMode,omitempty" name:"IsNonStaticIpMode"`

	// Indicates whether to enable cluster deletion protection.
	DeletionProtection *bool `json:"DeletionProtection,omitempty" name:"DeletionProtection"`

	// Cluster network proxy model, which is only used when ipvs-bpf mode is used. At present, TKE cluster supports three network proxy modes including `iptables`, `ipvs` and `ipvs-bpf` and their parameter setting relationships are as follows:
	// `iptables`: do not set IPVS and KubeProxyMode.
	// `ipvs`: set IPVS to `true` and do not set KubeProxyMode.
	// `ipvs-bpf`: set KubeProxyMode to `kube-proxy-bpf`.
	// The following conditions are required to use ipvs-bpf network mode:
	// 1. The cluster version must be v1.14 or later.
	// 2. The system image must be Tencent Linux 2.4.
	KubeProxyMode *string `json:"KubeProxyMode,omitempty" name:"KubeProxyMode"`

	// Indicates whether to enable auditing
	AuditEnabled *bool `json:"AuditEnabled,omitempty" name:"AuditEnabled"`

	// Specifies the ID of logset to which the audit logs are uploaded.
	AuditLogsetId *string `json:"AuditLogsetId,omitempty" name:"AuditLogsetId"`

	// Specifies the ID of topic to which the audit logs are uploaded.
	AuditLogTopicId *string `json:"AuditLogTopicId,omitempty" name:"AuditLogTopicId"`

	// Specifies the ENI type. Values: `tke-route-eni` (multi-IP shared ENI); `tke-direct-eni` (independent ENI). It defaults to `tke-route-eni`.
	VpcCniType *string `json:"VpcCniType,omitempty" name:"VpcCniType"`

	// Runtime version
	RuntimeVersion *string `json:"RuntimeVersion,omitempty" name:"RuntimeVersion"`

	// Indicates whether to enable the custom mode for the node’s pod CIDR range
	EnableCustomizedPodCIDR *bool `json:"EnableCustomizedPodCIDR,omitempty" name:"EnableCustomizedPodCIDR"`

	// The basic number of Pods in custom mode
	BasePodNumber *int64 `json:"BasePodNumber,omitempty" name:"BasePodNumber"`

	// Specifies whether to enable Cilium. If it’s left empty, Cilium is not enabled. If `clusterIP` is passed in, it means to enable Cilium to support the clusterIP service type.
	CiliumMode *string `json:"CiliumMode,omitempty" name:"CiliumMode"`

	// Whether it is a dual-stack cluster in VPC-CNI mode. Default value: `false`, which indicates it is not a dual-stack cluster.
	IsDualStack *bool `json:"IsDualStack,omitempty" name:"IsDualStack"`

	// Whether to enable qGPU Sharing
	QGPUShareEnable *bool `json:"QGPUShareEnable,omitempty" name:"QGPUShareEnable"`
}

type ClusterAsGroup struct {
	// Scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// Scaling group status (`enabled`, `enabling`, `disabled`, `disabling`, `updating`, `deleting`, `scaleDownEnabling`, `scaleDownDisabling`)
	Status *string `json:"Status,omitempty" name:"Status"`

	// Whether the node is set to unschedulable
	// Note: this field may return null, indicating that no valid value was found.
	IsUnschedulable *bool `json:"IsUnschedulable,omitempty" name:"IsUnschedulable"`

	// Scaling group label list
	// Note: this field may return null, indicating that no valid value was found.
	Labels []*Label `json:"Labels,omitempty" name:"Labels"`

	// Creation time
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type ClusterAsGroupAttribute struct {
	// Scaling group ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// Whether it is enabled
	AutoScalingGroupEnabled *bool `json:"AutoScalingGroupEnabled,omitempty" name:"AutoScalingGroupEnabled"`

	// Maximum and minimum number of pods in a scaling group
	AutoScalingGroupRange *AutoScalingGroupRange `json:"AutoScalingGroupRange,omitempty" name:"AutoScalingGroupRange"`
}

type ClusterAsGroupOption struct {
	// Whether to enable scale-in
	// Note: this field may return null, indicating that no valid value was found.
	IsScaleDownEnabled *bool `json:"IsScaleDownEnabled,omitempty" name:"IsScaleDownEnabled"`

	// The scale-out method when there are multiple scaling groups. `random`: select a random scaling group. `most-pods`: choose the scaling group that can schedule the most pods. `least-waste`: select the scaling group that can ensure the fewest remaining resources after Pod scheduling.. The default value is `random`.)
	// Note: this field may return null, indicating that no valid value was found.
	Expander *string `json:"Expander,omitempty" name:"Expander"`

	// Max concurrent scale-in volume
	// Note: this field may return null, indicating that no valid value was found.
	MaxEmptyBulkDelete *int64 `json:"MaxEmptyBulkDelete,omitempty" name:"MaxEmptyBulkDelete"`

	// Number of minutes after cluster scale-out when the system starts judging whether to perform scale-in
	// Note: this field may return null, indicating that no valid value was found.
	ScaleDownDelay *int64 `json:"ScaleDownDelay,omitempty" name:"ScaleDownDelay"`

	// Number of consecutive minutes of idleness after which the node is subject to scale-in (default value: 10)
	// Note: this field may return null, indicating that no valid value was found.
	ScaleDownUnneededTime *int64 `json:"ScaleDownUnneededTime,omitempty" name:"ScaleDownUnneededTime"`

	// Percentage of node resource usage below which the node is considered to be idle (default value: 50)
	// Note: this field may return null, indicating that no valid value was found.
	ScaleDownUtilizationThreshold *int64 `json:"ScaleDownUtilizationThreshold,omitempty" name:"ScaleDownUtilizationThreshold"`

	// During scale-in, ignore nodes with local storage pods (default value: False)
	// Note: this field may return null, indicating that no valid value was found.
	SkipNodesWithLocalStorage *bool `json:"SkipNodesWithLocalStorage,omitempty" name:"SkipNodesWithLocalStorage"`

	// During scale-in, ignore nodes with pods in the kube-system namespace that are not managed by DaemonSet (default value: False)
	// Note: this field may return null, indicating that no valid value was found.
	SkipNodesWithSystemPods *bool `json:"SkipNodesWithSystemPods,omitempty" name:"SkipNodesWithSystemPods"`

	// Whether to ignore DaemonSet pods by default when calculating resource usage (default value: False: do not ignore)
	// Note: this field may return null, indicating that no valid value was found.
	IgnoreDaemonSetsUtilization *bool `json:"IgnoreDaemonSetsUtilization,omitempty" name:"IgnoreDaemonSetsUtilization"`

	// Number at which CA health detection is triggered (default value: 3). After the number specified in OkTotalUnreadyCount is exceeded, CA will perform health detection.
	// Note: this field may return null, indicating that no valid value was found.
	OkTotalUnreadyCount *int64 `json:"OkTotalUnreadyCount,omitempty" name:"OkTotalUnreadyCount"`

	// Max percentage of unready nodes. After the max percentage is exceeded, CA will stop operation.
	// Note: this field may return null, indicating that no valid value was found.
	MaxTotalUnreadyPercentage *int64 `json:"MaxTotalUnreadyPercentage,omitempty" name:"MaxTotalUnreadyPercentage"`

	// Amount of time before unready nodes become eligible for scale-in
	// Note: this field may return null, indicating that no valid value was found.
	ScaleDownUnreadyTime *int64 `json:"ScaleDownUnreadyTime,omitempty" name:"ScaleDownUnreadyTime"`

	// Waiting time before CA deletes nodes that are not registered in Kubernetes
	// Note: this field may return null, indicating that no valid value was found.
	UnregisteredNodeRemovalTime *int64 `json:"UnregisteredNodeRemovalTime,omitempty" name:"UnregisteredNodeRemovalTime"`
}

type ClusterBasicSettings struct {
	// Cluster operating system. Public image (enter the image ID) and custom image (enter the image name) are supported. For details, see https://intl.cloud.tencent.com/document/product/457/68289?from_cn_redirect=1
	ClusterOs *string `json:"ClusterOs,omitempty" name:"ClusterOs"`

	// Cluster version. The default value is 1.10.5.
	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`

	// Cluster name
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// Cluster description
	ClusterDescription *string `json:"ClusterDescription,omitempty" name:"ClusterDescription"`

	// VPC ID, in the format of vpc-xxx, which is required when you create an empty managed cluster.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// ID of the project to which the new resources in the cluster belong.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Tag description list. This parameter is used to bind a tag to a resource instance. Currently, a tag can only be bound to cluster instances.
	TagSpecification []*TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification"`

	// Container image tag, `DOCKER_CUSTOMIZE` (container customized tag), `GENERAL` (general tag, default value)
	OsCustomizeType *string `json:"OsCustomizeType,omitempty" name:"OsCustomizeType"`

	// Whether to enable the node’s default security group (default: `No`, Alpha feature)
	NeedWorkSecurityGroup *bool `json:"NeedWorkSecurityGroup,omitempty" name:"NeedWorkSecurityGroup"`

	// When the Cilium Overlay add-on is selected, TKE will take two IPs from the subnet to create the private network CLB.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Cluster specifications available for managed clusters
	ClusterLevel *string `json:"ClusterLevel,omitempty" name:"ClusterLevel"`

	// Auto cluster upgrade for managed clusters
	AutoUpgradeClusterLevel *AutoUpgradeClusterLevel `json:"AutoUpgradeClusterLevel,omitempty" name:"AutoUpgradeClusterLevel"`
}

type ClusterCIDRSettings struct {
	// CIDR used to assign container and service IPs for the cluster. It cannot conflict with the VPC's CIDR or the CIDRs of other clusters in the same VPC
	ClusterCIDR *string `json:"ClusterCIDR,omitempty" name:"ClusterCIDR"`

	// Whether to ignore ClusterCIDR conflict errors, which are not ignored by default
	IgnoreClusterCIDRConflict *bool `json:"IgnoreClusterCIDRConflict,omitempty" name:"IgnoreClusterCIDRConflict"`

	// Maximum number of Pods on each node. Value range: 16 to 256. When its power is not 2, it will round upward to the closest power of 2.
	MaxNodePodNum *uint64 `json:"MaxNodePodNum,omitempty" name:"MaxNodePodNum"`

	// The maximum number of services in a cluster. The range is from 32 to 32768. When its power is not 2, it will round upward to the closest power of 2. Default value is 256.
	MaxClusterServiceNum *uint64 `json:"MaxClusterServiceNum,omitempty" name:"MaxClusterServiceNum"`

	// The CIDR block used to assign cluster service IP addresses. It must conflict with neither the VPC CIDR block nor with CIDR blocks of other clusters in the same VPC instance. The IP range must be within the private network IP range, such as 10.1.0.0/14, 192.168.0.1/18, and 172.16.0.0/16.
	ServiceCIDR *string `json:"ServiceCIDR,omitempty" name:"ServiceCIDR"`

	// Subnet ID of the ENI in VPC-CNI network mode
	EniSubnetIds []*string `json:"EniSubnetIds,omitempty" name:"EniSubnetIds"`

	// Repossession time of ENI IP addresses in VPC-CNI network mode, whose range is [300,15768000)
	ClaimExpiredSeconds *int64 `json:"ClaimExpiredSeconds,omitempty" name:"ClaimExpiredSeconds"`

	// Whether to ignore ServiceCIDR conflict errors. It is only valid in VPC-CNI mode. Default value: `false`.
	IgnoreServiceCIDRConflict *bool `json:"IgnoreServiceCIDRConflict,omitempty" name:"IgnoreServiceCIDRConflict"`
}

type ClusterCondition struct {
	// Process type
	Type *string `json:"Type,omitempty" name:"Type"`

	// Process status
	Status *string `json:"Status,omitempty" name:"Status"`

	// Last time when the status is probed
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	LastProbeTime *string `json:"LastProbeTime,omitempty" name:"LastProbeTime"`

	// Last time when transiting to the process
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	LastTransitionTime *string `json:"LastTransitionTime,omitempty" name:"LastTransitionTime"`

	// Reasons for transiting to the process
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// More information on transition
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitempty" name:"Message"`
}

type ClusterCredential struct {
	// CA root certificate
	CACert *string `json:"CACert,omitempty" name:"CACert"`

	// Token for authentication
	Token *string `json:"Token,omitempty" name:"Token"`
}

type ClusterExtraArgs struct {
	// kube-apiserver custom parameter, in the format of ["k1=v1", "k1=v2"], for example: ["max-requests-inflight=500","feature-gates=PodShareProcessNamespace=true,DynamicKubeletConfig=true"].
	// Note: this field may return `null`, indicating that no valid value is obtained.
	KubeAPIServer []*string `json:"KubeAPIServer,omitempty" name:"KubeAPIServer"`

	// kube-controller-manager custom parameter
	// Note: this field may return null, indicating that no valid value is obtained.
	KubeControllerManager []*string `json:"KubeControllerManager,omitempty" name:"KubeControllerManager"`

	// kube-scheduler custom parameter
	// Note: this field may return null, indicating that no valid value is obtained.
	KubeScheduler []*string `json:"KubeScheduler,omitempty" name:"KubeScheduler"`

	// etcd custom parameter, which is only effective for self-deployed cluster.
	// Note: this field may return `null`, indicating that no valid value is obtained.
	Etcd []*string `json:"Etcd,omitempty" name:"Etcd"`
}

type ClusterLevelAttribute struct {
	// Cluster model
	Name *string `json:"Name,omitempty" name:"Name"`

	// Model name
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// Number of nodes
	NodeCount *uint64 `json:"NodeCount,omitempty" name:"NodeCount"`

	// Number of Pods
	PodCount *uint64 `json:"PodCount,omitempty" name:"PodCount"`

	// Number of ConfigMap
	ConfigMapCount *uint64 `json:"ConfigMapCount,omitempty" name:"ConfigMapCount"`

	// Number of CRDs
	CRDCount *uint64 `json:"CRDCount,omitempty" name:"CRDCount"`

	// Whether it is enabled
	Enable *bool `json:"Enable,omitempty" name:"Enable"`

	// Number of other resources
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	OtherCount *uint64 `json:"OtherCount,omitempty" name:"OtherCount"`
}

type ClusterLevelChangeRecord struct {
	// Record ID
	ID *string `json:"ID,omitempty" name:"ID"`

	// Cluster ID
	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`

	// Status (valid values: `trading`, `upgrading`, `success`, `failed`)
	Status *string `json:"Status,omitempty" name:"Status"`

	// Status description
	Message *string `json:"Message,omitempty" name:"Message"`

	// Original model
	OldLevel *string `json:"OldLevel,omitempty" name:"OldLevel"`

	// New model
	NewLevel *string `json:"NewLevel,omitempty" name:"NewLevel"`

	// Trigger type (valid values: `manual`, `auto`)
	TriggerType *string `json:"TriggerType,omitempty" name:"TriggerType"`

	// Start time
	StartedAt *string `json:"StartedAt,omitempty" name:"StartedAt"`

	// End time
	EndedAt *string `json:"EndedAt,omitempty" name:"EndedAt"`
}

type ClusterNetworkSettings struct {
	// CIDR used to assign container and service IPs for the cluster. It cannot conflict with the VPC's CIDR or the CIDRs of other clusters in the same VPC.
	ClusterCIDR *string `json:"ClusterCIDR,omitempty" name:"ClusterCIDR"`

	// Whether to ignore ClusterCIDR conflict errors. It defaults to not ignore.
	IgnoreClusterCIDRConflict *bool `json:"IgnoreClusterCIDRConflict,omitempty" name:"IgnoreClusterCIDRConflict"`

	// Maximum number of pods on each node in the cluster. Default value: 256
	MaxNodePodNum *uint64 `json:"MaxNodePodNum,omitempty" name:"MaxNodePodNum"`

	// Maximum number of cluster services. Default value: 256
	MaxClusterServiceNum *uint64 `json:"MaxClusterServiceNum,omitempty" name:"MaxClusterServiceNum"`

	// Whether IPVS is enabled. Default value: disabled
	Ipvs *bool `json:"Ipvs,omitempty" name:"Ipvs"`

	// Cluster VPC ID, which is required when you create an empty cluster; otherwise, it is automatically set to be consistent with that of the nodes in the cluster
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Whether CNI is enabled for network plugin(s). Default value: enabled
	Cni *bool `json:"Cni,omitempty" name:"Cni"`

	// The network mode of service. This parameter is only applicable to ipvs+bpf mode.
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	KubeProxyMode *string `json:"KubeProxyMode,omitempty" name:"KubeProxyMode"`

	// The IP range for service assignment. It cannot conflict with the VPC’s CIDR block nor the CIDR blocks of other clusters in the same VPC.
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	ServiceCIDR *string `json:"ServiceCIDR,omitempty" name:"ServiceCIDR"`

	// The container subnet associated with the cluster
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	Subnets []*string `json:"Subnets,omitempty" name:"Subnets"`

	// Whether to ignore ServiceCIDR conflict errors. It is only valid in VPC-CNI mode. Default value: `false`.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	IgnoreServiceCIDRConflict *bool `json:"IgnoreServiceCIDRConflict,omitempty" name:"IgnoreServiceCIDRConflict"`

	// Whether it is a dual-stack cluster in VPC-CNI mode. Default value: `false`, which indicates it is not a dual-stack cluster.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	IsDualStack *bool `json:"IsDualStack,omitempty" name:"IsDualStack"`

	// It is used to automatically assign the IP ranges for the service.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Ipv6ServiceCIDR *string `json:"Ipv6ServiceCIDR,omitempty" name:"Ipv6ServiceCIDR"`
}

type ClusterStatus struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Cluster status
	ClusterState *string `json:"ClusterState,omitempty" name:"ClusterState"`

	// Status of nodes in the cluster
	ClusterInstanceState *string `json:"ClusterInstanceState,omitempty" name:"ClusterInstanceState"`

	// Indicates whether the monitoring service is enabled for the cluster
	ClusterBMonitor *bool `json:"ClusterBMonitor,omitempty" name:"ClusterBMonitor"`

	// Number of cluster nodes being created. "-1" indicates that the request to obtain the node status timed out. "-2" indicates that the request failed.
	ClusterInitNodeNum *int64 `json:"ClusterInitNodeNum,omitempty" name:"ClusterInitNodeNum"`

	// Number of running nodes in the cluster. "-1" indicates that the request to obtain the node status timed out. "-2" indicates that the request failed.
	ClusterRunningNodeNum *int64 `json:"ClusterRunningNodeNum,omitempty" name:"ClusterRunningNodeNum"`

	// Number of abnormal nodes in the cluster.  "-1" indicates that the request to obtain the node status timed out. "-2" indicates that the request failed.
	ClusterFailedNodeNum *int64 `json:"ClusterFailedNodeNum,omitempty" name:"ClusterFailedNodeNum"`

	// Number of shutdown nodes in the cluster.  "-1" indicates that the request to obtain the node status timed out. "-2" indicates that the request failed.
	// Note: this field may return `null`, indicating that no valid value can be found.
	ClusterClosedNodeNum *int64 `json:"ClusterClosedNodeNum,omitempty" name:"ClusterClosedNodeNum"`

	// Number of nodes being shut down in the cluster.  "-1" indicates that the request to obtain the node status timed out. "-2" indicates that the request failed.
	// Note: this field may return `null`, indicating that no valid value can be found.
	ClusterClosingNodeNum *int64 `json:"ClusterClosingNodeNum,omitempty" name:"ClusterClosingNodeNum"`

	// Indicates whether to enable cluster deletion protection
	// Note: this field may return `null`, indicating that no valid value can be found.
	ClusterDeletionProtection *bool `json:"ClusterDeletionProtection,omitempty" name:"ClusterDeletionProtection"`

	// Indicates whether the cluster is auditable
	// Note: this field may return `null`, indicating that no valid value can be found.
	ClusterAuditEnabled *bool `json:"ClusterAuditEnabled,omitempty" name:"ClusterAuditEnabled"`
}

type ClusterVersion struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The list of cluster major version, such as 1.18.4
	Versions []*string `json:"Versions,omitempty" name:"Versions"`
}

type CommonName struct {
	// User UIN
	SubaccountUin *string `json:"SubaccountUin,omitempty" name:"SubaccountUin"`

	// The CommonName in the certificate of the client corresponding to the sub-account
	CN *string `json:"CN,omitempty" name:"CN"`
}

// Predefined struct for user
type CreateClusterEndpointRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The ID of the subnet where the cluster's port is located (only needs to be entered when the non-public network access is enabled, and must be within the subnet of the cluster's VPC). 
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Whether public network access is enabled or not (True = public network access, FALSE = private network access, with the default value as FALSE).
	IsExtranet *bool `json:"IsExtranet,omitempty" name:"IsExtranet"`

	// The domain name
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// The security group in use. Required only for public network access.
	SecurityGroup *string `json:"SecurityGroup,omitempty" name:"SecurityGroup"`

	// The LB parameter. Required only for public network access.
	ExtensiveParameters *string `json:"ExtensiveParameters,omitempty" name:"ExtensiveParameters"`
}

type CreateClusterEndpointRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The ID of the subnet where the cluster's port is located (only needs to be entered when the non-public network access is enabled, and must be within the subnet of the cluster's VPC). 
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Whether public network access is enabled or not (True = public network access, FALSE = private network access, with the default value as FALSE).
	IsExtranet *bool `json:"IsExtranet,omitempty" name:"IsExtranet"`

	// The domain name
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// The security group in use. Required only for public network access.
	SecurityGroup *string `json:"SecurityGroup,omitempty" name:"SecurityGroup"`

	// The LB parameter. Required only for public network access.
	ExtensiveParameters *string `json:"ExtensiveParameters,omitempty" name:"ExtensiveParameters"`
}

func (r *CreateClusterEndpointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterEndpointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SubnetId")
	delete(f, "IsExtranet")
	delete(f, "Domain")
	delete(f, "SecurityGroup")
	delete(f, "ExtensiveParameters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClusterEndpointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterEndpointResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateClusterEndpointResponse struct {
	*tchttp.BaseResponse
	Response *CreateClusterEndpointResponseParams `json:"Response"`
}

func (r *CreateClusterEndpointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterEndpointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterEndpointVipRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Security policy opens single IP or CIDR to the Internet (for example: '192.168.1.0/24', with 'reject all' as the default).
	SecurityPolicies []*string `json:"SecurityPolicies,omitempty" name:"SecurityPolicies"`
}

type CreateClusterEndpointVipRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Security policy opens single IP or CIDR to the Internet (for example: '192.168.1.0/24', with 'reject all' as the default).
	SecurityPolicies []*string `json:"SecurityPolicies,omitempty" name:"SecurityPolicies"`
}

func (r *CreateClusterEndpointVipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterEndpointVipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SecurityPolicies")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClusterEndpointVipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterEndpointVipResponseParams struct {
	// Request job's FlowId
	RequestFlowId *int64 `json:"RequestFlowId,omitempty" name:"RequestFlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateClusterEndpointVipResponse struct {
	*tchttp.BaseResponse
	Response *CreateClusterEndpointVipResponseParams `json:"Response"`
}

func (r *CreateClusterEndpointVipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterEndpointVipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterInstancesRequestParams struct {
	// Cluster ID. Enter the ClusterId field returned by the DescribeClusters API
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Pass-through parameter for CVM creation in the format of a JSON string. To ensure the idempotence of requests for adding cluster nodes, you need to add the ClientToken field in this parameter. For more information, see the documentation for [RunInstances](https://intl.cloud.tencent.com/document/product/213/15730?from_cn_redirect=1) API.
	RunInstancePara *string `json:"RunInstancePara,omitempty" name:"RunInstancePara"`

	// Additional parameter to be set for the instance
	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitempty" name:"InstanceAdvancedSettings"`

	// Skips the specified verification. Valid values: GlobalRouteCIDRCheck, VpcCniCIDRCheck
	SkipValidateOptions []*string `json:"SkipValidateOptions,omitempty" name:"SkipValidateOptions"`
}

type CreateClusterInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID. Enter the ClusterId field returned by the DescribeClusters API
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Pass-through parameter for CVM creation in the format of a JSON string. To ensure the idempotence of requests for adding cluster nodes, you need to add the ClientToken field in this parameter. For more information, see the documentation for [RunInstances](https://intl.cloud.tencent.com/document/product/213/15730?from_cn_redirect=1) API.
	RunInstancePara *string `json:"RunInstancePara,omitempty" name:"RunInstancePara"`

	// Additional parameter to be set for the instance
	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitempty" name:"InstanceAdvancedSettings"`

	// Skips the specified verification. Valid values: GlobalRouteCIDRCheck, VpcCniCIDRCheck
	SkipValidateOptions []*string `json:"SkipValidateOptions,omitempty" name:"SkipValidateOptions"`
}

func (r *CreateClusterInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "RunInstancePara")
	delete(f, "InstanceAdvancedSettings")
	delete(f, "SkipValidateOptions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClusterInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterInstancesResponseParams struct {
	// Instance ID
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateClusterInstancesResponse struct {
	*tchttp.BaseResponse
	Response *CreateClusterInstancesResponseParams `json:"Response"`
}

func (r *CreateClusterInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterNodePoolFromExistingAsgRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Scaling group ID
	AutoscalingGroupId *string `json:"AutoscalingGroupId,omitempty" name:"AutoscalingGroupId"`
}

type CreateClusterNodePoolFromExistingAsgRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Scaling group ID
	AutoscalingGroupId *string `json:"AutoscalingGroupId,omitempty" name:"AutoscalingGroupId"`
}

func (r *CreateClusterNodePoolFromExistingAsgRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterNodePoolFromExistingAsgRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "AutoscalingGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClusterNodePoolFromExistingAsgRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterNodePoolFromExistingAsgResponseParams struct {
	// Node pool ID
	NodePoolId *string `json:"NodePoolId,omitempty" name:"NodePoolId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateClusterNodePoolFromExistingAsgResponse struct {
	*tchttp.BaseResponse
	Response *CreateClusterNodePoolFromExistingAsgResponseParams `json:"Response"`
}

func (r *CreateClusterNodePoolFromExistingAsgResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterNodePoolFromExistingAsgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterNodePoolRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// AS group parameters
	AutoScalingGroupPara *string `json:"AutoScalingGroupPara,omitempty" name:"AutoScalingGroupPara"`

	// Running parameters
	LaunchConfigurePara *string `json:"LaunchConfigurePara,omitempty" name:"LaunchConfigurePara"`

	// Sample parameters
	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitempty" name:"InstanceAdvancedSettings"`

	// Indicates whether to enable auto scaling
	EnableAutoscale *bool `json:"EnableAutoscale,omitempty" name:"EnableAutoscale"`

	// Node pool name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Labels
	Labels []*Label `json:"Labels,omitempty" name:"Labels"`

	// Taints
	Taints []*Taint `json:"Taints,omitempty" name:"Taints"`

	// Node pool runtime type and version
	ContainerRuntime *string `json:"ContainerRuntime,omitempty" name:"ContainerRuntime"`

	// Runtime version
	RuntimeVersion *string `json:"RuntimeVersion,omitempty" name:"RuntimeVersion"`

	// Node pool operating system (enter the image ID for a custom image, and enter the OS name for a public image)
	NodePoolOs *string `json:"NodePoolOs,omitempty" name:"NodePoolOs"`

	// Container image tag, `DOCKER_CUSTOMIZE` (container customized tag), `GENERAL` (general tag, default value)
	OsCustomizeType *string `json:"OsCustomizeType,omitempty" name:"OsCustomizeType"`

	// Resource tag
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Whether Deletion Protection is enabled
	DeletionProtection *bool `json:"DeletionProtection,omitempty" name:"DeletionProtection"`
}

type CreateClusterNodePoolRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// AS group parameters
	AutoScalingGroupPara *string `json:"AutoScalingGroupPara,omitempty" name:"AutoScalingGroupPara"`

	// Running parameters
	LaunchConfigurePara *string `json:"LaunchConfigurePara,omitempty" name:"LaunchConfigurePara"`

	// Sample parameters
	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitempty" name:"InstanceAdvancedSettings"`

	// Indicates whether to enable auto scaling
	EnableAutoscale *bool `json:"EnableAutoscale,omitempty" name:"EnableAutoscale"`

	// Node pool name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Labels
	Labels []*Label `json:"Labels,omitempty" name:"Labels"`

	// Taints
	Taints []*Taint `json:"Taints,omitempty" name:"Taints"`

	// Node pool runtime type and version
	ContainerRuntime *string `json:"ContainerRuntime,omitempty" name:"ContainerRuntime"`

	// Runtime version
	RuntimeVersion *string `json:"RuntimeVersion,omitempty" name:"RuntimeVersion"`

	// Node pool operating system (enter the image ID for a custom image, and enter the OS name for a public image)
	NodePoolOs *string `json:"NodePoolOs,omitempty" name:"NodePoolOs"`

	// Container image tag, `DOCKER_CUSTOMIZE` (container customized tag), `GENERAL` (general tag, default value)
	OsCustomizeType *string `json:"OsCustomizeType,omitempty" name:"OsCustomizeType"`

	// Resource tag
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Whether Deletion Protection is enabled
	DeletionProtection *bool `json:"DeletionProtection,omitempty" name:"DeletionProtection"`
}

func (r *CreateClusterNodePoolRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterNodePoolRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "AutoScalingGroupPara")
	delete(f, "LaunchConfigurePara")
	delete(f, "InstanceAdvancedSettings")
	delete(f, "EnableAutoscale")
	delete(f, "Name")
	delete(f, "Labels")
	delete(f, "Taints")
	delete(f, "ContainerRuntime")
	delete(f, "RuntimeVersion")
	delete(f, "NodePoolOs")
	delete(f, "OsCustomizeType")
	delete(f, "Tags")
	delete(f, "DeletionProtection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClusterNodePoolRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterNodePoolResponseParams struct {
	// Node pool ID
	NodePoolId *string `json:"NodePoolId,omitempty" name:"NodePoolId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateClusterNodePoolResponse struct {
	*tchttp.BaseResponse
	Response *CreateClusterNodePoolResponseParams `json:"Response"`
}

func (r *CreateClusterNodePoolResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterNodePoolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterRequestParams struct {
	// Container networking configuration information for the cluster
	ClusterCIDRSettings *ClusterCIDRSettings `json:"ClusterCIDRSettings,omitempty" name:"ClusterCIDRSettings"`

	// Cluster type. Managed cluster: MANAGED_CLUSTER; self-deployed cluster: INDEPENDENT_CLUSTER.
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// Pass-through parameter for CVM creation in the format of a JSON string. For more information, see the API for [creating a CVM instance](https://intl.cloud.tencent.com/document/product/213/15730?from_cn_redirect=1).
	RunInstancesForNode []*RunInstancesForNode `json:"RunInstancesForNode,omitempty" name:"RunInstancesForNode"`

	// Basic configuration information of the cluster
	ClusterBasicSettings *ClusterBasicSettings `json:"ClusterBasicSettings,omitempty" name:"ClusterBasicSettings"`

	// Advanced configuration information of the cluster
	ClusterAdvancedSettings *ClusterAdvancedSettings `json:"ClusterAdvancedSettings,omitempty" name:"ClusterAdvancedSettings"`

	// Advanced configuration information of the node
	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitempty" name:"InstanceAdvancedSettings"`

	// The configuration information for existing instances. All instances must be in the same VPC. Up to 100 instances are allowed in one VPC. Spot instances are not supported.
	ExistedInstancesForNode []*ExistedInstancesForNode `json:"ExistedInstancesForNode,omitempty" name:"ExistedInstancesForNode"`

	// CVM type and the corresponding data disk mounting configuration information.
	InstanceDataDiskMountSettings []*InstanceDataDiskMountSetting `json:"InstanceDataDiskMountSettings,omitempty" name:"InstanceDataDiskMountSettings"`

	// Information of the add-on to be installed
	ExtensionAddons []*ExtensionAddon `json:"ExtensionAddons,omitempty" name:"ExtensionAddons"`
}

type CreateClusterRequest struct {
	*tchttp.BaseRequest
	
	// Container networking configuration information for the cluster
	ClusterCIDRSettings *ClusterCIDRSettings `json:"ClusterCIDRSettings,omitempty" name:"ClusterCIDRSettings"`

	// Cluster type. Managed cluster: MANAGED_CLUSTER; self-deployed cluster: INDEPENDENT_CLUSTER.
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// Pass-through parameter for CVM creation in the format of a JSON string. For more information, see the API for [creating a CVM instance](https://intl.cloud.tencent.com/document/product/213/15730?from_cn_redirect=1).
	RunInstancesForNode []*RunInstancesForNode `json:"RunInstancesForNode,omitempty" name:"RunInstancesForNode"`

	// Basic configuration information of the cluster
	ClusterBasicSettings *ClusterBasicSettings `json:"ClusterBasicSettings,omitempty" name:"ClusterBasicSettings"`

	// Advanced configuration information of the cluster
	ClusterAdvancedSettings *ClusterAdvancedSettings `json:"ClusterAdvancedSettings,omitempty" name:"ClusterAdvancedSettings"`

	// Advanced configuration information of the node
	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitempty" name:"InstanceAdvancedSettings"`

	// The configuration information for existing instances. All instances must be in the same VPC. Up to 100 instances are allowed in one VPC. Spot instances are not supported.
	ExistedInstancesForNode []*ExistedInstancesForNode `json:"ExistedInstancesForNode,omitempty" name:"ExistedInstancesForNode"`

	// CVM type and the corresponding data disk mounting configuration information.
	InstanceDataDiskMountSettings []*InstanceDataDiskMountSetting `json:"InstanceDataDiskMountSettings,omitempty" name:"InstanceDataDiskMountSettings"`

	// Information of the add-on to be installed
	ExtensionAddons []*ExtensionAddon `json:"ExtensionAddons,omitempty" name:"ExtensionAddons"`
}

func (r *CreateClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterCIDRSettings")
	delete(f, "ClusterType")
	delete(f, "RunInstancesForNode")
	delete(f, "ClusterBasicSettings")
	delete(f, "ClusterAdvancedSettings")
	delete(f, "InstanceAdvancedSettings")
	delete(f, "ExistedInstancesForNode")
	delete(f, "InstanceDataDiskMountSettings")
	delete(f, "ExtensionAddons")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterResponseParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateClusterResponse struct {
	*tchttp.BaseResponse
	Response *CreateClusterResponseParams `json:"Response"`
}

func (r *CreateClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterRouteTableRequestParams struct {
	// Route table name
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`

	// Route table CIDR
	RouteTableCidrBlock *string `json:"RouteTableCidrBlock,omitempty" name:"RouteTableCidrBlock"`

	// VPC bound to the route table
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Whether to ignore CIDR conflicts
	IgnoreClusterCidrConflict *int64 `json:"IgnoreClusterCidrConflict,omitempty" name:"IgnoreClusterCidrConflict"`
}

type CreateClusterRouteTableRequest struct {
	*tchttp.BaseRequest
	
	// Route table name
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`

	// Route table CIDR
	RouteTableCidrBlock *string `json:"RouteTableCidrBlock,omitempty" name:"RouteTableCidrBlock"`

	// VPC bound to the route table
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Whether to ignore CIDR conflicts
	IgnoreClusterCidrConflict *int64 `json:"IgnoreClusterCidrConflict,omitempty" name:"IgnoreClusterCidrConflict"`
}

func (r *CreateClusterRouteTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterRouteTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableName")
	delete(f, "RouteTableCidrBlock")
	delete(f, "VpcId")
	delete(f, "IgnoreClusterCidrConflict")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClusterRouteTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterRouteTableResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateClusterRouteTableResponse struct {
	*tchttp.BaseResponse
	Response *CreateClusterRouteTableResponseParams `json:"Response"`
}

func (r *CreateClusterRouteTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterRouteTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateECMInstancesRequestParams struct {
	// Cluster ID
	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`

	// Module ID
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// Instance AZ, number of instances and ISP
	ZoneInstanceCountISPSet []*ECMZoneInstanceCountISP `json:"ZoneInstanceCountISPSet,omitempty" name:"ZoneInstanceCountISPSet"`

	// Password
	Password *string `json:"Password,omitempty" name:"Password"`

	// Public network bandwidth
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`

	// Image ID
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Host name
	HostName *string `json:"HostName,omitempty" name:"HostName"`

	// Enhanced service (including CWP and Cloud Monitoring)
	EnhancedService *ECMEnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`

	// Custom script
	UserData *string `json:"UserData,omitempty" name:"UserData"`

	// Instance extension information
	External *string `json:"External,omitempty" name:"External"`

	// Security group of the instance
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
}

type CreateECMInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`

	// Module ID
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// Instance AZ, number of instances and ISP
	ZoneInstanceCountISPSet []*ECMZoneInstanceCountISP `json:"ZoneInstanceCountISPSet,omitempty" name:"ZoneInstanceCountISPSet"`

	// Password
	Password *string `json:"Password,omitempty" name:"Password"`

	// Public network bandwidth
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`

	// Image ID
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Host name
	HostName *string `json:"HostName,omitempty" name:"HostName"`

	// Enhanced service (including CWP and Cloud Monitoring)
	EnhancedService *ECMEnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`

	// Custom script
	UserData *string `json:"UserData,omitempty" name:"UserData"`

	// Instance extension information
	External *string `json:"External,omitempty" name:"External"`

	// Security group of the instance
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
}

func (r *CreateECMInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateECMInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterID")
	delete(f, "ModuleId")
	delete(f, "ZoneInstanceCountISPSet")
	delete(f, "Password")
	delete(f, "InternetMaxBandwidthOut")
	delete(f, "ImageId")
	delete(f, "InstanceName")
	delete(f, "HostName")
	delete(f, "EnhancedService")
	delete(f, "UserData")
	delete(f, "External")
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateECMInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateECMInstancesResponseParams struct {
	// ECM ID list
	EcmIdSet []*string `json:"EcmIdSet,omitempty" name:"EcmIdSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateECMInstancesResponse struct {
	*tchttp.BaseResponse
	Response *CreateECMInstancesResponseParams `json:"Response"`
}

func (r *CreateECMInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateECMInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusAlertRuleRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Alarm configurations
	AlertRule *PrometheusAlertRuleDetail `json:"AlertRule,omitempty" name:"AlertRule"`
}

type CreatePrometheusAlertRuleRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Alarm configurations
	AlertRule *PrometheusAlertRuleDetail `json:"AlertRule,omitempty" name:"AlertRule"`
}

func (r *CreatePrometheusAlertRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusAlertRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AlertRule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrometheusAlertRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusAlertRuleResponseParams struct {
	// Alarm ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePrometheusAlertRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrometheusAlertRuleResponseParams `json:"Response"`
}

func (r *CreatePrometheusAlertRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusAlertRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTKEEdgeClusterRequestParams struct {

	K8SVersion *string `json:"K8SVersion,omitempty" name:"K8SVersion"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Cluster name
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// Cluster Pod CIDR block
	PodCIDR *string `json:"PodCIDR,omitempty" name:"PodCIDR"`

	// Cluster service CIDR block
	ServiceCIDR *string `json:"ServiceCIDR,omitempty" name:"ServiceCIDR"`

	// Cluster description
	ClusterDesc *string `json:"ClusterDesc,omitempty" name:"ClusterDesc"`

	// Cluster advanced settings
	ClusterAdvancedSettings *EdgeClusterAdvancedSettings `json:"ClusterAdvancedSettings,omitempty" name:"ClusterAdvancedSettings"`

	// Maximum number of Pods on the node
	MaxNodePodNum *int64 `json:"MaxNodePodNum,omitempty" name:"MaxNodePodNum"`

	// Public LB of the TKE Edge cluster
	PublicLB *EdgeClusterPublicLB `json:"PublicLB,omitempty" name:"PublicLB"`
}

type CreateTKEEdgeClusterRequest struct {
	*tchttp.BaseRequest
	
	K8SVersion *string `json:"K8SVersion,omitempty" name:"K8SVersion"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Cluster name
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// Cluster Pod CIDR block
	PodCIDR *string `json:"PodCIDR,omitempty" name:"PodCIDR"`

	// Cluster service CIDR block
	ServiceCIDR *string `json:"ServiceCIDR,omitempty" name:"ServiceCIDR"`

	// Cluster description
	ClusterDesc *string `json:"ClusterDesc,omitempty" name:"ClusterDesc"`

	// Cluster advanced settings
	ClusterAdvancedSettings *EdgeClusterAdvancedSettings `json:"ClusterAdvancedSettings,omitempty" name:"ClusterAdvancedSettings"`

	// Maximum number of Pods on the node
	MaxNodePodNum *int64 `json:"MaxNodePodNum,omitempty" name:"MaxNodePodNum"`

	// Public LB of the TKE Edge cluster
	PublicLB *EdgeClusterPublicLB `json:"PublicLB,omitempty" name:"PublicLB"`
}

func (r *CreateTKEEdgeClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTKEEdgeClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "K8SVersion")
	delete(f, "VpcId")
	delete(f, "ClusterName")
	delete(f, "PodCIDR")
	delete(f, "ServiceCIDR")
	delete(f, "ClusterDesc")
	delete(f, "ClusterAdvancedSettings")
	delete(f, "MaxNodePodNum")
	delete(f, "PublicLB")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTKEEdgeClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTKEEdgeClusterResponseParams struct {
	// TKE Edge cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateTKEEdgeClusterResponse struct {
	*tchttp.BaseResponse
	Response *CreateTKEEdgeClusterResponseParams `json:"Response"`
}

func (r *CreateTKEEdgeClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTKEEdgeClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataDisk struct {
	// Disk type
	// Note: this field may return null, indicating that no valid values can be obtained.
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// File system (ext3/ext4/xfs)
	// Note: This field may return null, indicating that no valid value was found.
	FileSystem *string `json:"FileSystem,omitempty" name:"FileSystem"`

	// Disk size (G)
	// Note: This field may return null, indicating that no valid value was found.
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// Whether the disk is auto-formatted and mounted
	// Note: this field may return `null`, indicating that no valid value is obtained.
	AutoFormatAndMount *bool `json:"AutoFormatAndMount,omitempty" name:"AutoFormatAndMount"`

	// Mounting directory
	// Note: This field may return null, indicating that no valid value was found.
	MountTarget *string `json:"MountTarget,omitempty" name:"MountTarget"`

	// Mounted device name or partition name (only required when adding an existing node)
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	DiskPartition *string `json:"DiskPartition,omitempty" name:"DiskPartition"`
}

// Predefined struct for user
type DeleteClusterAsGroupsRequestParams struct {
	// The cluster ID, obtained through the [DescribeClusters](https://intl.cloud.tencent.com/document/api/457/31862?from_cn_redirect=1) API.
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Cluster scaling group ID list
	AutoScalingGroupIds []*string `json:"AutoScalingGroupIds,omitempty" name:"AutoScalingGroupIds"`

	// Whether to keep nodes in the scaling group. Default to **false** (not keep)
	KeepInstance *bool `json:"KeepInstance,omitempty" name:"KeepInstance"`
}

type DeleteClusterAsGroupsRequest struct {
	*tchttp.BaseRequest
	
	// The cluster ID, obtained through the [DescribeClusters](https://intl.cloud.tencent.com/document/api/457/31862?from_cn_redirect=1) API.
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Cluster scaling group ID list
	AutoScalingGroupIds []*string `json:"AutoScalingGroupIds,omitempty" name:"AutoScalingGroupIds"`

	// Whether to keep nodes in the scaling group. Default to **false** (not keep)
	KeepInstance *bool `json:"KeepInstance,omitempty" name:"KeepInstance"`
}

func (r *DeleteClusterAsGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterAsGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "AutoScalingGroupIds")
	delete(f, "KeepInstance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteClusterAsGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterAsGroupsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteClusterAsGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteClusterAsGroupsResponseParams `json:"Response"`
}

func (r *DeleteClusterAsGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterAsGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterEndpointRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Whether public network access is enabled or not (True = public network access, FALSE = private network access, with the default value as FALSE).
	IsExtranet *bool `json:"IsExtranet,omitempty" name:"IsExtranet"`
}

type DeleteClusterEndpointRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Whether public network access is enabled or not (True = public network access, FALSE = private network access, with the default value as FALSE).
	IsExtranet *bool `json:"IsExtranet,omitempty" name:"IsExtranet"`
}

func (r *DeleteClusterEndpointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterEndpointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "IsExtranet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteClusterEndpointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterEndpointResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteClusterEndpointResponse struct {
	*tchttp.BaseResponse
	Response *DeleteClusterEndpointResponseParams `json:"Response"`
}

func (r *DeleteClusterEndpointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterEndpointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterEndpointVipRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type DeleteClusterEndpointVipRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DeleteClusterEndpointVipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterEndpointVipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteClusterEndpointVipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterEndpointVipResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteClusterEndpointVipResponse struct {
	*tchttp.BaseResponse
	Response *DeleteClusterEndpointVipResponseParams `json:"Response"`
}

func (r *DeleteClusterEndpointVipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterEndpointVipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterInstancesRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// List of Instance IDs
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Policy used to delete an instance in the cluster: `terminate` (terminates the instance. Only available for pay-as-you-go CVMs); `retain` (only removes it from the cluster. The instance will be retained.)
	InstanceDeleteMode *string `json:"InstanceDeleteMode,omitempty" name:"InstanceDeleteMode"`

	// Whether or not there is forced deletion (when a node is initialized, the parameters can be specified as TRUE)
	ForceDelete *bool `json:"ForceDelete,omitempty" name:"ForceDelete"`
}

type DeleteClusterInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// List of Instance IDs
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Policy used to delete an instance in the cluster: `terminate` (terminates the instance. Only available for pay-as-you-go CVMs); `retain` (only removes it from the cluster. The instance will be retained.)
	InstanceDeleteMode *string `json:"InstanceDeleteMode,omitempty" name:"InstanceDeleteMode"`

	// Whether or not there is forced deletion (when a node is initialized, the parameters can be specified as TRUE)
	ForceDelete *bool `json:"ForceDelete,omitempty" name:"ForceDelete"`
}

func (r *DeleteClusterInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceIds")
	delete(f, "InstanceDeleteMode")
	delete(f, "ForceDelete")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteClusterInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterInstancesResponseParams struct {
	// IDs of deleted instances
	SuccInstanceIds []*string `json:"SuccInstanceIds,omitempty" name:"SuccInstanceIds"`

	// IDs of instances failed to be deleted
	FailedInstanceIds []*string `json:"FailedInstanceIds,omitempty" name:"FailedInstanceIds"`

	// IDs of instances that cannot be found
	NotFoundInstanceIds []*string `json:"NotFoundInstanceIds,omitempty" name:"NotFoundInstanceIds"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteClusterInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteClusterInstancesResponseParams `json:"Response"`
}

func (r *DeleteClusterInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterNodePoolRequestParams struct {
	// ClusterId of a node pool
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// IDs of node pools to delete
	NodePoolIds []*string `json:"NodePoolIds,omitempty" name:"NodePoolIds"`

	// Indicates whether nodes in a node pool are retained when the node pool is deleted. (The nodes are removed from the cluster. However, the corresponding instances will not be terminated.)
	KeepInstance *bool `json:"KeepInstance,omitempty" name:"KeepInstance"`
}

type DeleteClusterNodePoolRequest struct {
	*tchttp.BaseRequest
	
	// ClusterId of a node pool
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// IDs of node pools to delete
	NodePoolIds []*string `json:"NodePoolIds,omitempty" name:"NodePoolIds"`

	// Indicates whether nodes in a node pool are retained when the node pool is deleted. (The nodes are removed from the cluster. However, the corresponding instances will not be terminated.)
	KeepInstance *bool `json:"KeepInstance,omitempty" name:"KeepInstance"`
}

func (r *DeleteClusterNodePoolRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterNodePoolRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NodePoolIds")
	delete(f, "KeepInstance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteClusterNodePoolRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterNodePoolResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteClusterNodePoolResponse struct {
	*tchttp.BaseResponse
	Response *DeleteClusterNodePoolResponseParams `json:"Response"`
}

func (r *DeleteClusterNodePoolResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterNodePoolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Policy used to delete an instance in the cluster: terminate (terminates the instance. Only available for instances on pay-as-you-go CVMs); retain (only removes it from the cluster. The instance will be retained.)
	InstanceDeleteMode *string `json:"InstanceDeleteMode,omitempty" name:"InstanceDeleteMode"`

	// Specifies the policy to deal with resources in the cluster when the cluster is deleted. It only supports CBS now. The default policy is to retain CBS disks.
	ResourceDeleteOptions []*ResourceDeleteOption `json:"ResourceDeleteOptions,omitempty" name:"ResourceDeleteOptions"`
}

type DeleteClusterRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Policy used to delete an instance in the cluster: terminate (terminates the instance. Only available for instances on pay-as-you-go CVMs); retain (only removes it from the cluster. The instance will be retained.)
	InstanceDeleteMode *string `json:"InstanceDeleteMode,omitempty" name:"InstanceDeleteMode"`

	// Specifies the policy to deal with resources in the cluster when the cluster is deleted. It only supports CBS now. The default policy is to retain CBS disks.
	ResourceDeleteOptions []*ResourceDeleteOption `json:"ResourceDeleteOptions,omitempty" name:"ResourceDeleteOptions"`
}

func (r *DeleteClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceDeleteMode")
	delete(f, "ResourceDeleteOptions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteClusterResponse struct {
	*tchttp.BaseResponse
	Response *DeleteClusterResponseParams `json:"Response"`
}

func (r *DeleteClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterRouteRequestParams struct {
	// Route table name.
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`

	// Next hop address.
	GatewayIp *string `json:"GatewayIp,omitempty" name:"GatewayIp"`

	// Destination CIDR.
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" name:"DestinationCidrBlock"`
}

type DeleteClusterRouteRequest struct {
	*tchttp.BaseRequest
	
	// Route table name.
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`

	// Next hop address.
	GatewayIp *string `json:"GatewayIp,omitempty" name:"GatewayIp"`

	// Destination CIDR.
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" name:"DestinationCidrBlock"`
}

func (r *DeleteClusterRouteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterRouteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableName")
	delete(f, "GatewayIp")
	delete(f, "DestinationCidrBlock")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteClusterRouteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterRouteResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteClusterRouteResponse struct {
	*tchttp.BaseResponse
	Response *DeleteClusterRouteResponseParams `json:"Response"`
}

func (r *DeleteClusterRouteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterRouteTableRequestParams struct {
	// Route table name
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`
}

type DeleteClusterRouteTableRequest struct {
	*tchttp.BaseRequest
	
	// Route table name
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`
}

func (r *DeleteClusterRouteTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterRouteTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteClusterRouteTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterRouteTableResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteClusterRouteTableResponse struct {
	*tchttp.BaseResponse
	Response *DeleteClusterRouteTableResponseParams `json:"Response"`
}

func (r *DeleteClusterRouteTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterRouteTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteECMInstancesRequestParams struct {
	// Cluster ID
	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`

	// IDs of ECMs to be deleted
	EcmIdSet []*string `json:"EcmIdSet,omitempty" name:"EcmIdSet"`
}

type DeleteECMInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`

	// IDs of ECMs to be deleted
	EcmIdSet []*string `json:"EcmIdSet,omitempty" name:"EcmIdSet"`
}

func (r *DeleteECMInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteECMInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterID")
	delete(f, "EcmIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteECMInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteECMInstancesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteECMInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteECMInstancesResponseParams `json:"Response"`
}

func (r *DeleteECMInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteECMInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEdgeCVMInstancesRequestParams struct {
	// Cluster ID
	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`

	// IDs of CVMs to be deleted
	CvmIdSet []*string `json:"CvmIdSet,omitempty" name:"CvmIdSet"`
}

type DeleteEdgeCVMInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`

	// IDs of CVMs to be deleted
	CvmIdSet []*string `json:"CvmIdSet,omitempty" name:"CvmIdSet"`
}

func (r *DeleteEdgeCVMInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEdgeCVMInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterID")
	delete(f, "CvmIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteEdgeCVMInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEdgeCVMInstancesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteEdgeCVMInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteEdgeCVMInstancesResponseParams `json:"Response"`
}

func (r *DeleteEdgeCVMInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEdgeCVMInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEdgeClusterInstancesRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Array of instance IDs to be deleted
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type DeleteEdgeClusterInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Array of instance IDs to be deleted
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DeleteEdgeClusterInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEdgeClusterInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteEdgeClusterInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEdgeClusterInstancesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteEdgeClusterInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteEdgeClusterInstancesResponseParams `json:"Response"`
}

func (r *DeleteEdgeClusterInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEdgeClusterInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusAlertRuleRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// The ID list of alarm rules
	AlertIds []*string `json:"AlertIds,omitempty" name:"AlertIds"`
}

type DeletePrometheusAlertRuleRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// The ID list of alarm rules
	AlertIds []*string `json:"AlertIds,omitempty" name:"AlertIds"`
}

func (r *DeletePrometheusAlertRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusAlertRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AlertIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePrometheusAlertRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusAlertRuleResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeletePrometheusAlertRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeletePrometheusAlertRuleResponseParams `json:"Response"`
}

func (r *DeletePrometheusAlertRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusAlertRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTKEEdgeClusterRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type DeleteTKEEdgeClusterRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DeleteTKEEdgeClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTKEEdgeClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTKEEdgeClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTKEEdgeClusterResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteTKEEdgeClusterResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTKEEdgeClusterResponseParams `json:"Response"`
}

func (r *DeleteTKEEdgeClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTKEEdgeClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAvailableClusterVersionRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// List of cluster IDs
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`
}

type DescribeAvailableClusterVersionRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// List of cluster IDs
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`
}

func (r *DescribeAvailableClusterVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAvailableClusterVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ClusterIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAvailableClusterVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAvailableClusterVersionResponseParams struct {
	// Upgradable cluster version
	// Note: this field may return `null`, indicating that no valid value is obtained.
	Versions []*string `json:"Versions,omitempty" name:"Versions"`

	// Cluster information
	// Note: this field may return `null`, indicating that no valid value is obtained.
	Clusters []*ClusterVersion `json:"Clusters,omitempty" name:"Clusters"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAvailableClusterVersionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAvailableClusterVersionResponseParams `json:"Response"`
}

func (r *DescribeAvailableClusterVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAvailableClusterVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAvailableTKEEdgeVersionRequestParams struct {

}

type DescribeAvailableTKEEdgeVersionRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAvailableTKEEdgeVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAvailableTKEEdgeVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAvailableTKEEdgeVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAvailableTKEEdgeVersionResponseParams struct {
	// Version list
	Versions []*string `json:"Versions,omitempty" name:"Versions"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAvailableTKEEdgeVersionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAvailableTKEEdgeVersionResponseParams `json:"Response"`
}

func (r *DescribeAvailableTKEEdgeVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAvailableTKEEdgeVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterAsGroupOptionRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type DescribeClusterAsGroupOptionRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterAsGroupOptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterAsGroupOptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterAsGroupOptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterAsGroupOptionResponseParams struct {
	// Cluster auto scaling attributes
	// Note: this field may return null, indicating that no valid value was found.
	ClusterAsGroupOption *ClusterAsGroupOption `json:"ClusterAsGroupOption,omitempty" name:"ClusterAsGroupOption"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClusterAsGroupOptionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterAsGroupOptionResponseParams `json:"Response"`
}

func (r *DescribeClusterAsGroupOptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterAsGroupOptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterAsGroupsRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Scaling group ID list. If this value is null, it indicates that all cluster-associated scaling groups are pulled.
	AutoScalingGroupIds []*string `json:"AutoScalingGroupIds,omitempty" name:"AutoScalingGroupIds"`

	// Offset. This value defaults to 0. For more information on Offset, see the relevant sections in API [Overview](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. This value defaults to 20. The maximum is 100. For more information on Limit, see the relevant sections in API [Overview](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeClusterAsGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Scaling group ID list. If this value is null, it indicates that all cluster-associated scaling groups are pulled.
	AutoScalingGroupIds []*string `json:"AutoScalingGroupIds,omitempty" name:"AutoScalingGroupIds"`

	// Offset. This value defaults to 0. For more information on Offset, see the relevant sections in API [Overview](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. This value defaults to 20. The maximum is 100. For more information on Limit, see the relevant sections in API [Overview](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeClusterAsGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterAsGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "AutoScalingGroupIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterAsGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterAsGroupsResponseParams struct {
	// Total number of scaling groups associated with the cluster
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Cluster-associated scaling group list
	ClusterAsGroupSet []*ClusterAsGroup `json:"ClusterAsGroupSet,omitempty" name:"ClusterAsGroupSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClusterAsGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterAsGroupsResponseParams `json:"Response"`
}

func (r *DescribeClusterAsGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterAsGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterAuthenticationOptionsRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type DescribeClusterAuthenticationOptionsRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterAuthenticationOptionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterAuthenticationOptionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterAuthenticationOptionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterAuthenticationOptionsResponseParams struct {
	// ServiceAccount authentication configuration
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ServiceAccounts *ServiceAccountAuthenticationOptions `json:"ServiceAccounts,omitempty" name:"ServiceAccounts"`

	// Result of the last modification. Values: `Updating`, `Success`, `Failed` or `TimeOut`.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	LatestOperationState *string `json:"LatestOperationState,omitempty" name:"LatestOperationState"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClusterAuthenticationOptionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterAuthenticationOptionsResponseParams `json:"Response"`
}

func (r *DescribeClusterAuthenticationOptionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterAuthenticationOptionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterCommonNamesRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Sub-account. Up to 50 sub-accounts can be passed in at a time.
	SubaccountUins []*string `json:"SubaccountUins,omitempty" name:"SubaccountUins"`

	// Role ID. Up to 50 role IDs can be passed in at a time.
	RoleIds []*string `json:"RoleIds,omitempty" name:"RoleIds"`
}

type DescribeClusterCommonNamesRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Sub-account. Up to 50 sub-accounts can be passed in at a time.
	SubaccountUins []*string `json:"SubaccountUins,omitempty" name:"SubaccountUins"`

	// Role ID. Up to 50 role IDs can be passed in at a time.
	RoleIds []*string `json:"RoleIds,omitempty" name:"RoleIds"`
}

func (r *DescribeClusterCommonNamesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterCommonNamesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SubaccountUins")
	delete(f, "RoleIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterCommonNamesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterCommonNamesResponseParams struct {
	// The CommonName in the certificate of the client corresponding to the sub-account UIN
	CommonNames []*CommonName `json:"CommonNames,omitempty" name:"CommonNames"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClusterCommonNamesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterCommonNamesResponseParams `json:"Response"`
}

func (r *DescribeClusterCommonNamesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterCommonNamesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterEndpointStatusRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Whether public network access is enabled or not (True = public network access, FALSE = private network access, with the default value as FALSE).
	IsExtranet *bool `json:"IsExtranet,omitempty" name:"IsExtranet"`
}

type DescribeClusterEndpointStatusRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Whether public network access is enabled or not (True = public network access, FALSE = private network access, with the default value as FALSE).
	IsExtranet *bool `json:"IsExtranet,omitempty" name:"IsExtranet"`
}

func (r *DescribeClusterEndpointStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterEndpointStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "IsExtranet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterEndpointStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterEndpointStatusResponseParams struct {
	// The status of cluster access port. It can be `Created` (enabled); `Creating` (enabling) and `NotFound` (not enabled)
	// Note: this field may return `null`, indicating that no valid value is obtained.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Details of the error occurred while opening the access port
	// Note: this field may return `null`, indicating that no valid value is obtained.
	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClusterEndpointStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterEndpointStatusResponseParams `json:"Response"`
}

func (r *DescribeClusterEndpointStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterEndpointStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterEndpointVipStatusRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type DescribeClusterEndpointVipStatusRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterEndpointVipStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterEndpointVipStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterEndpointVipStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterEndpointVipStatusResponseParams struct {
	// Port operation status (Creating = in the process of creation; CreateFailed = creation has failed; Created = creation completed; Deleting = in the process of deletion; DeletedFailed = deletion has failed; Deleted = deletion completed; NotFound = operation not found)
	Status *string `json:"Status,omitempty" name:"Status"`

	// Reason for operation failure
	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClusterEndpointVipStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterEndpointVipStatusResponseParams `json:"Response"`
}

func (r *DescribeClusterEndpointVipStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterEndpointVipStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterInstancesRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of output entries. Default value: 20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// List of instance IDs to be obtained. This parameter is empty by default, which indicates that all instances in the cluster will be pulled.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Node role. Valid values are MASTER, WORKER, ETCD, MASTER_ETCD, and ALL. Default value: WORKER.
	InstanceRole *string `json:"InstanceRole,omitempty" name:"InstanceRole"`

	// Filters include `nodepool-id` and `nodepool-instance-type` (how the instance is added to the pool). For `nodepool-instance-type`, the values can be `MANUALLY_ADDED`, `AUTOSCALING_ADDED` and `ALL`.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeClusterInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of output entries. Default value: 20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// List of instance IDs to be obtained. This parameter is empty by default, which indicates that all instances in the cluster will be pulled.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Node role. Valid values are MASTER, WORKER, ETCD, MASTER_ETCD, and ALL. Default value: WORKER.
	InstanceRole *string `json:"InstanceRole,omitempty" name:"InstanceRole"`

	// Filters include `nodepool-id` and `nodepool-instance-type` (how the instance is added to the pool). For `nodepool-instance-type`, the values can be `MANUALLY_ADDED`, `AUTOSCALING_ADDED` and `ALL`.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeClusterInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "InstanceIds")
	delete(f, "InstanceRole")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterInstancesResponseParams struct {
	// Total number of instances in the cluster
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of instances in the cluster
	InstanceSet []*Instance `json:"InstanceSet,omitempty" name:"InstanceSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClusterInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterInstancesResponseParams `json:"Response"`
}

func (r *DescribeClusterInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterKubeconfigRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Defaults to `false`, which means to obtain the kubeconfig of private network
	IsExtranet *bool `json:"IsExtranet,omitempty" name:"IsExtranet"`
}

type DescribeClusterKubeconfigRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Defaults to `false`, which means to obtain the kubeconfig of private network
	IsExtranet *bool `json:"IsExtranet,omitempty" name:"IsExtranet"`
}

func (r *DescribeClusterKubeconfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterKubeconfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "IsExtranet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterKubeconfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterKubeconfigResponseParams struct {
	// Sub-account kubeconfig file, used to access the cluster kube-apiserver directly
	Kubeconfig *string `json:"Kubeconfig,omitempty" name:"Kubeconfig"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClusterKubeconfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterKubeconfigResponseParams `json:"Response"`
}

func (r *DescribeClusterKubeconfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterKubeconfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterLevelAttributeRequestParams struct {
	// Cluster ID (available for cluster model adjustment)
	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
}

type DescribeClusterLevelAttributeRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID (available for cluster model adjustment)
	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
}

func (r *DescribeClusterLevelAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterLevelAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterLevelAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterLevelAttributeResponseParams struct {
	// Total number
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Cluster model
	Items []*ClusterLevelAttribute `json:"Items,omitempty" name:"Items"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClusterLevelAttributeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterLevelAttributeResponseParams `json:"Response"`
}

func (r *DescribeClusterLevelAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterLevelAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterLevelChangeRecordsRequestParams struct {
	// Cluster ID
	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`

	// Start time
	StartAt *string `json:"StartAt,omitempty" name:"StartAt"`

	// End time
	EndAt *string `json:"EndAt,omitempty" name:"EndAt"`

	// Offset. Default value: `0`
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of output entries. Default value: `20`
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeClusterLevelChangeRecordsRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`

	// Start time
	StartAt *string `json:"StartAt,omitempty" name:"StartAt"`

	// End time
	EndAt *string `json:"EndAt,omitempty" name:"EndAt"`

	// Offset. Default value: `0`
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of output entries. Default value: `20`
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeClusterLevelChangeRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterLevelChangeRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterID")
	delete(f, "StartAt")
	delete(f, "EndAt")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterLevelChangeRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterLevelChangeRecordsResponseParams struct {
	// Total number
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Cluster model
	Items []*ClusterLevelChangeRecord `json:"Items,omitempty" name:"Items"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClusterLevelChangeRecordsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterLevelChangeRecordsResponseParams `json:"Response"`
}

func (r *DescribeClusterLevelChangeRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterLevelChangeRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterNodePoolDetailRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Node pool ID
	NodePoolId *string `json:"NodePoolId,omitempty" name:"NodePoolId"`
}

type DescribeClusterNodePoolDetailRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Node pool ID
	NodePoolId *string `json:"NodePoolId,omitempty" name:"NodePoolId"`
}

func (r *DescribeClusterNodePoolDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterNodePoolDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NodePoolId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterNodePoolDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterNodePoolDetailResponseParams struct {
	// Node pool details
	NodePool *NodePool `json:"NodePool,omitempty" name:"NodePool"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClusterNodePoolDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterNodePoolDetailResponseParams `json:"Response"`
}

func (r *DescribeClusterNodePoolDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterNodePoolDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterNodePoolsRequestParams struct {
	// ClusterId (cluster ID)
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// ·  NodePoolsName
	//     Filters by the node pool name
	//     Type: String
	//     Required: No
	// 
	// ·  NodePoolsId
	//     Filters by the node pool ID
	//     Type: String
	//     Required: No
	// 
	// ·  tags
	//     Filters by key-value pairs of tags
	//     Type: String
	//     Required: No
	// 
	// ·  tag:tag-key
	//     Filters by key-value pairs of tags
	//     Type: String
	//     Required: No
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeClusterNodePoolsRequest struct {
	*tchttp.BaseRequest
	
	// ClusterId (cluster ID)
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// ·  NodePoolsName
	//     Filters by the node pool name
	//     Type: String
	//     Required: No
	// 
	// ·  NodePoolsId
	//     Filters by the node pool ID
	//     Type: String
	//     Required: No
	// 
	// ·  tags
	//     Filters by key-value pairs of tags
	//     Type: String
	//     Required: No
	// 
	// ·  tag:tag-key
	//     Filters by key-value pairs of tags
	//     Type: String
	//     Required: No
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeClusterNodePoolsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterNodePoolsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterNodePoolsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterNodePoolsResponseParams struct {
	// NodePools (node pool list)
	// Note: this field may return `null`, indicating that no valid value is obtained.
	NodePoolSet []*NodePool `json:"NodePoolSet,omitempty" name:"NodePoolSet"`

	// Total resources
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClusterNodePoolsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterNodePoolsResponseParams `json:"Response"`
}

func (r *DescribeClusterNodePoolsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterNodePoolsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterRouteTablesRequestParams struct {

}

type DescribeClusterRouteTablesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeClusterRouteTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterRouteTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterRouteTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterRouteTablesResponseParams struct {
	// Number of instances that match the filter condition(s).
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Object of cluster route table.
	RouteTableSet []*RouteTableInfo `json:"RouteTableSet,omitempty" name:"RouteTableSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClusterRouteTablesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterRouteTablesResponseParams `json:"Response"`
}

func (r *DescribeClusterRouteTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterRouteTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterRoutesRequestParams struct {
	// Route table name.
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`

	// Filtering conditions, which are optional. Currently, only filtering by GatewayIP is supported.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeClusterRoutesRequest struct {
	*tchttp.BaseRequest
	
	// Route table name.
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`

	// Filtering conditions, which are optional. Currently, only filtering by GatewayIP is supported.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeClusterRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableName")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterRoutesResponseParams struct {
	// Number of instances that match the filter condition(s).
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Object of cluster route.
	RouteSet []*RouteInfo `json:"RouteSet,omitempty" name:"RouteSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClusterRoutesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterRoutesResponseParams `json:"Response"`
}

func (r *DescribeClusterRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterSecurityRequestParams struct {
	// Cluster ID. Enter the ClusterId field returned by the DescribeClusters API
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type DescribeClusterSecurityRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID. Enter the ClusterId field returned by the DescribeClusters API
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterSecurityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterSecurityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterSecurityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterSecurityResponseParams struct {
	// Cluster's account name
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Cluster's password
	Password *string `json:"Password,omitempty" name:"Password"`

	// Cluster's access CA certificate
	CertificationAuthority *string `json:"CertificationAuthority,omitempty" name:"CertificationAuthority"`

	// Cluster's access address
	ClusterExternalEndpoint *string `json:"ClusterExternalEndpoint,omitempty" name:"ClusterExternalEndpoint"`

	// Domain name accessed by the cluster
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Cluster's endpoint address
	PgwEndpoint *string `json:"PgwEndpoint,omitempty" name:"PgwEndpoint"`

	// Cluster's access policy group
	// Note: This field may return null, indicating that no valid value was found.
	SecurityPolicy []*string `json:"SecurityPolicy,omitempty" name:"SecurityPolicy"`

	// Cluster Kubeconfig file
	// Note: This field may return null, indicating that no valid value was found.
	Kubeconfig *string `json:"Kubeconfig,omitempty" name:"Kubeconfig"`

	// Access address of the cluster JnsGw
	// Note: This field may return null, indicating that no valid value was found.
	JnsGwEndpoint *string `json:"JnsGwEndpoint,omitempty" name:"JnsGwEndpoint"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClusterSecurityResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterSecurityResponseParams `json:"Response"`
}

func (r *DescribeClusterSecurityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterSecurityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterStatusRequestParams struct {
	// Cluster ID list. All clusters are pulled if it is left empty.
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`
}

type DescribeClusterStatusRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID list. All clusters are pulled if it is left empty.
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`
}

func (r *DescribeClusterStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterStatusResponseParams struct {
	// Cluster status list
	ClusterStatusSet []*ClusterStatus `json:"ClusterStatusSet,omitempty" name:"ClusterStatusSet"`

	// Number of clusters
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClusterStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterStatusResponseParams `json:"Response"`
}

func (r *DescribeClusterStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClustersRequestParams struct {
	// Cluster ID list (When it is empty,
	// all clusters under the account will be obtained)
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`

	// Offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of output entries. Default value: 20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// ·  ClusterName
	//     Filters by the cluster name
	//     Type: String
	//     Required: No
	// 
	// ·  ClusterType
	//     Filters by the cluster type
	//     Type: String
	//     Required: No
	// 
	// ·  ClusterStatus
	//     Filters by the cluster status
	//     Type: String
	//     Required: No
	// 
	// ·  Tags
	//     Filters by key-value pairs of tags
	//     Type: String
	//     Required: No
	// 
	// ·  vpc-id
	//     Filters by the VPC ID
	//     Type: String
	//     Required: No
	// 
	// ·  tag-key
	//     Filters by the tag key
	//     Type: String
	//     Required: No
	// 
	// ·  tag-value
	//     Filters by the tag value
	//     Type: String
	//     Required: No
	// 
	// ·  tag:tag-key
	//     Filters by key-value pairs of tags
	//     Type: String
	//     Required: No
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Cluster type, such as `MANAGED_CLUSTER`
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
}

type DescribeClustersRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID list (When it is empty,
	// all clusters under the account will be obtained)
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`

	// Offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of output entries. Default value: 20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// ·  ClusterName
	//     Filters by the cluster name
	//     Type: String
	//     Required: No
	// 
	// ·  ClusterType
	//     Filters by the cluster type
	//     Type: String
	//     Required: No
	// 
	// ·  ClusterStatus
	//     Filters by the cluster status
	//     Type: String
	//     Required: No
	// 
	// ·  Tags
	//     Filters by key-value pairs of tags
	//     Type: String
	//     Required: No
	// 
	// ·  vpc-id
	//     Filters by the VPC ID
	//     Type: String
	//     Required: No
	// 
	// ·  tag-key
	//     Filters by the tag key
	//     Type: String
	//     Required: No
	// 
	// ·  tag-value
	//     Filters by the tag value
	//     Type: String
	//     Required: No
	// 
	// ·  tag:tag-key
	//     Filters by key-value pairs of tags
	//     Type: String
	//     Required: No
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Cluster type, such as `MANAGED_CLUSTER`
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
}

func (r *DescribeClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterIds")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "ClusterType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClustersResponseParams struct {
	// Total number of clusters
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Cluster information list
	Clusters []*Cluster `json:"Clusters,omitempty" name:"Clusters"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClustersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClustersResponseParams `json:"Response"`
}

func (r *DescribeClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeECMInstancesRequestParams struct {
	// Cluster ID
	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`

	// Filter condition
	// Only filtering by an ECM ID is supported
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeECMInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`

	// Filter condition
	// Only filtering by an ECM ID is supported
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeECMInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeECMInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterID")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeECMInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeECMInstancesResponseParams struct {
	// Number of instances matched the condition
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of the returned instance information
	InstanceInfoSet []*string `json:"InstanceInfoSet,omitempty" name:"InstanceInfoSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeECMInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeECMInstancesResponseParams `json:"Response"`
}

func (r *DescribeECMInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeECMInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeAvailableExtraArgsRequestParams struct {
	// Cluster version
	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`
}

type DescribeEdgeAvailableExtraArgsRequest struct {
	*tchttp.BaseRequest
	
	// Cluster version
	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`
}

func (r *DescribeEdgeAvailableExtraArgsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeAvailableExtraArgsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgeAvailableExtraArgsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeAvailableExtraArgsResponseParams struct {
	// Cluster version
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`

	// Available custom parameters
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	AvailableExtraArgs *EdgeAvailableExtraArgs `json:"AvailableExtraArgs,omitempty" name:"AvailableExtraArgs"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEdgeAvailableExtraArgsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgeAvailableExtraArgsResponseParams `json:"Response"`
}

func (r *DescribeEdgeAvailableExtraArgsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeAvailableExtraArgsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeCVMInstancesRequestParams struct {
	// Cluster ID
	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`

	// Filter condition
	// Only `cvm-id` is supported.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeEdgeCVMInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`

	// Filter condition
	// Only `cvm-id` is supported.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeEdgeCVMInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeCVMInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterID")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgeCVMInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeCVMInstancesResponseParams struct {
	// Number of instances matched the condition
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of the returned instance information
	InstanceInfoSet []*string `json:"InstanceInfoSet,omitempty" name:"InstanceInfoSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEdgeCVMInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgeCVMInstancesResponseParams `json:"Response"`
}

func (r *DescribeEdgeCVMInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeCVMInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeClusterExtraArgsRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type DescribeEdgeClusterExtraArgsRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeEdgeClusterExtraArgsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeClusterExtraArgsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgeClusterExtraArgsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeClusterExtraArgsResponseParams struct {
	// Custom parameters of the cluster
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ClusterExtraArgs *EdgeClusterExtraArgs `json:"ClusterExtraArgs,omitempty" name:"ClusterExtraArgs"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEdgeClusterExtraArgsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgeClusterExtraArgsResponseParams `json:"Response"`
}

func (r *DescribeEdgeClusterExtraArgsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeClusterExtraArgsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeClusterInstancesRequestParams struct {
	// Cluster ID
	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`

	// Max number of returned entries
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Filter condition. Only `NodeName` is supported.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeEdgeClusterInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`

	// Max number of returned entries
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Filter condition. Only `NodeName` is supported.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeEdgeClusterInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeClusterInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterID")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgeClusterInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeClusterInstancesResponseParams struct {
	// Total number of nodes in the cluster
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Array of node information
	InstanceInfoSet *string `json:"InstanceInfoSet,omitempty" name:"InstanceInfoSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEdgeClusterInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgeClusterInstancesResponseParams `json:"Response"`
}

func (r *DescribeEdgeClusterInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeClusterInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnableVpcCniProgressRequestParams struct {
	// ID of the cluster for which you want to enable the VPC-CNI mode
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type DescribeEnableVpcCniProgressRequest struct {
	*tchttp.BaseRequest
	
	// ID of the cluster for which you want to enable the VPC-CNI mode
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeEnableVpcCniProgressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnableVpcCniProgressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEnableVpcCniProgressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnableVpcCniProgressResponseParams struct {
	// Task status, which can be `Running`, `Succeed`, or `Failed`.
	Status *string `json:"Status,omitempty" name:"Status"`

	// The description for the task status when the task status is “Failed”, for example, failed to install the IPAMD component.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEnableVpcCniProgressResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEnableVpcCniProgressResponseParams `json:"Response"`
}

func (r *DescribeEnableVpcCniProgressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnableVpcCniProgressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExistedInstancesRequestParams struct {
	// Cluster ID. Enter the `ClusterId` field returned when you call the DescribeClusters API (Only VPC ID obtained through `ClusterId` need filtering conditions. When comparing statuses, the nodes on all clusters in this region will be used for comparison. You cannot specify `InstanceIds` and `ClusterId` at the same time.)
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Query by one or more instance ID(s). Instance ID format: ins-xxxxxxxx. (Refer to section ID.N of the API overview for this parameter's specific format.) Up to 100 instances are allowed for each request. You cannot specify InstanceIds and Filters at the same time.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Filter condition. For fields and other information, see [the DescribeInstances API](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1). If a ClusterId has been set, then the cluster's VPC ID will be attached as a query field. In this situation, if a "vpc-id" is specified in Filter, then the specified VPC ID must be consistent with the cluster's VPC ID.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Filter by instance IP (Supports both private and public IPs)
	VagueIpAddress *string `json:"VagueIpAddress,omitempty" name:"VagueIpAddress"`

	// Filter by instance name
	VagueInstanceName *string `json:"VagueInstanceName,omitempty" name:"VagueInstanceName"`

	// Offset. Default value: 0. For more information on Offset, see the relevant section in the API [Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on Limit, see the relevant section in the API [Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Filter by multiple instance IPs
	IpAddresses []*string `json:"IpAddresses,omitempty" name:"IpAddresses"`
}

type DescribeExistedInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID. Enter the `ClusterId` field returned when you call the DescribeClusters API (Only VPC ID obtained through `ClusterId` need filtering conditions. When comparing statuses, the nodes on all clusters in this region will be used for comparison. You cannot specify `InstanceIds` and `ClusterId` at the same time.)
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Query by one or more instance ID(s). Instance ID format: ins-xxxxxxxx. (Refer to section ID.N of the API overview for this parameter's specific format.) Up to 100 instances are allowed for each request. You cannot specify InstanceIds and Filters at the same time.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Filter condition. For fields and other information, see [the DescribeInstances API](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1). If a ClusterId has been set, then the cluster's VPC ID will be attached as a query field. In this situation, if a "vpc-id" is specified in Filter, then the specified VPC ID must be consistent with the cluster's VPC ID.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Filter by instance IP (Supports both private and public IPs)
	VagueIpAddress *string `json:"VagueIpAddress,omitempty" name:"VagueIpAddress"`

	// Filter by instance name
	VagueInstanceName *string `json:"VagueInstanceName,omitempty" name:"VagueInstanceName"`

	// Offset. Default value: 0. For more information on Offset, see the relevant section in the API [Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on Limit, see the relevant section in the API [Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Filter by multiple instance IPs
	IpAddresses []*string `json:"IpAddresses,omitempty" name:"IpAddresses"`
}

func (r *DescribeExistedInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExistedInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceIds")
	delete(f, "Filters")
	delete(f, "VagueIpAddress")
	delete(f, "VagueInstanceName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "IpAddresses")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExistedInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExistedInstancesResponseParams struct {
	// Array of existing instance information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExistedInstanceSet []*ExistedInstance `json:"ExistedInstanceSet,omitempty" name:"ExistedInstanceSet"`

	// Number of instances that match the filter condition(s).
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeExistedInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExistedInstancesResponseParams `json:"Response"`
}

func (r *DescribeExistedInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExistedInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImagesRequestParams struct {

}

type DescribeImagesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeImagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImagesResponseParams struct {
	// Number of images
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Image information list
	// Note: this field may return null, indicating that no valid values can be obtained.
	ImageInstanceSet []*ImageInstance `json:"ImageInstanceSet,omitempty" name:"ImageInstanceSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeImagesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImagesResponseParams `json:"Response"`
}

func (r *DescribeImagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusInstanceRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribePrometheusInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribePrometheusInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusInstanceResponseParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance name
	Name *string `json:"Name,omitempty" name:"Name"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// COS bucket name
	COSBucket *string `json:"COSBucket,omitempty" name:"COSBucket"`

	// Data query address
	QueryAddress *string `json:"QueryAddress,omitempty" name:"QueryAddress"`

	// The grafana related information in the instance
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	Grafana *PrometheusGrafanaInfo `json:"Grafana,omitempty" name:"Grafana"`

	// Custom alertmanager
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	AlertManagerUrl *string `json:"AlertManagerUrl,omitempty" name:"AlertManagerUrl"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePrometheusInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusInstanceResponseParams `json:"Response"`
}

func (r *DescribePrometheusInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegionsRequestParams struct {

}

type DescribeRegionsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegionsResponseParams struct {
	// Number of regions
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// ## Region List
	// Note: this field may return null, indicating that no valid values can be obtained.
	RegionInstanceSet []*RegionInstance `json:"RegionInstanceSet,omitempty" name:"RegionInstanceSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRegionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRegionsResponseParams `json:"Response"`
}

func (r *DescribeRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceUsageRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type DescribeResourceUsageRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeResourceUsageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceUsageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourceUsageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceUsageResponseParams struct {
	// CRD usage
	CRDUsage *ResourceUsage `json:"CRDUsage,omitempty" name:"CRDUsage"`

	// Pod usage
	PodUsage *uint64 `json:"PodUsage,omitempty" name:"PodUsage"`

	// ConfigMap usage
	ConfigMapUsage *uint64 `json:"ConfigMapUsage,omitempty" name:"ConfigMapUsage"`

	// Other resource usage
	OtherUsage *ResourceUsage `json:"OtherUsage,omitempty" name:"OtherUsage"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeResourceUsageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourceUsageResponseParams `json:"Response"`
}

func (r *DescribeResourceUsageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRouteTableConflictsRequestParams struct {
	// Route table CIDR
	RouteTableCidrBlock *string `json:"RouteTableCidrBlock,omitempty" name:"RouteTableCidrBlock"`

	// VPC bound to the route table
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

type DescribeRouteTableConflictsRequest struct {
	*tchttp.BaseRequest
	
	// Route table CIDR
	RouteTableCidrBlock *string `json:"RouteTableCidrBlock,omitempty" name:"RouteTableCidrBlock"`

	// VPC bound to the route table
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

func (r *DescribeRouteTableConflictsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRouteTableConflictsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableCidrBlock")
	delete(f, "VpcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRouteTableConflictsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRouteTableConflictsResponseParams struct {
	// Whether there is a conflict in the route table.
	HasConflict *bool `json:"HasConflict,omitempty" name:"HasConflict"`

	// Route table conflict list.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RouteTableConflictSet []*RouteTableConflict `json:"RouteTableConflictSet,omitempty" name:"RouteTableConflictSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRouteTableConflictsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRouteTableConflictsResponseParams `json:"Response"`
}

func (r *DescribeRouteTableConflictsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRouteTableConflictsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTKEEdgeClusterCredentialRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type DescribeTKEEdgeClusterCredentialRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeTKEEdgeClusterCredentialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTKEEdgeClusterCredentialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTKEEdgeClusterCredentialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTKEEdgeClusterCredentialResponseParams struct {
	// Access address of the cluster
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Addresses []*IPAddress `json:"Addresses,omitempty" name:"Addresses"`

	// Cluster authentication information
	Credential *ClusterCredential `json:"Credential,omitempty" name:"Credential"`

	// Public network access information of the cluster
	PublicLB *EdgeClusterPublicLB `json:"PublicLB,omitempty" name:"PublicLB"`

	// Private network access information of the cluster
	InternalLB *EdgeClusterInternalLB `json:"InternalLB,omitempty" name:"InternalLB"`

	// CoreDns deployment information of the cluster
	CoreDns *string `json:"CoreDns,omitempty" name:"CoreDns"`

	// Multi-region health check deployment information of the cluster
	HealthRegion *string `json:"HealthRegion,omitempty" name:"HealthRegion"`

	// Health check deployment information of the cluster
	Health *string `json:"Health,omitempty" name:"Health"`

	// Whether to deploy GridDaemon to support headless service
	GridDaemon *string `json:"GridDaemon,omitempty" name:"GridDaemon"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTKEEdgeClusterCredentialResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTKEEdgeClusterCredentialResponseParams `json:"Response"`
}

func (r *DescribeTKEEdgeClusterCredentialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTKEEdgeClusterCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTKEEdgeClusterStatusRequestParams struct {
	// Edge compute cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type DescribeTKEEdgeClusterStatusRequest struct {
	*tchttp.BaseRequest
	
	// Edge compute cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeTKEEdgeClusterStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTKEEdgeClusterStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTKEEdgeClusterStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTKEEdgeClusterStatusResponseParams struct {
	// Current cluster status
	Phase *string `json:"Phase,omitempty" name:"Phase"`

	// Array of cluster processes
	Conditions []*ClusterCondition `json:"Conditions,omitempty" name:"Conditions"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTKEEdgeClusterStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTKEEdgeClusterStatusResponseParams `json:"Response"`
}

func (r *DescribeTKEEdgeClusterStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTKEEdgeClusterStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTKEEdgeClustersRequestParams struct {
	// Cluster ID list (when it is empty,
	// all clusters under the account are obtained)
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`

	// Offset. Default value: `0`
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of output entries. Default value: `20`
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Filter condition (only filtering by a single ClusterName is supported)
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeTKEEdgeClustersRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID list (when it is empty,
	// all clusters under the account are obtained)
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`

	// Offset. Default value: `0`
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of output entries. Default value: `20`
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Filter condition (only filtering by a single ClusterName is supported)
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeTKEEdgeClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTKEEdgeClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterIds")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTKEEdgeClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTKEEdgeClustersResponseParams struct {
	// Total number of clusters
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Cluster information list
	Clusters []*EdgeCluster `json:"Clusters,omitempty" name:"Clusters"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTKEEdgeClustersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTKEEdgeClustersResponseParams `json:"Response"`
}

func (r *DescribeTKEEdgeClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTKEEdgeClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTKEEdgeExternalKubeconfigRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type DescribeTKEEdgeExternalKubeconfigRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeTKEEdgeExternalKubeconfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTKEEdgeExternalKubeconfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTKEEdgeExternalKubeconfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTKEEdgeExternalKubeconfigResponseParams struct {
	// Kubeconfig file content
	Kubeconfig *string `json:"Kubeconfig,omitempty" name:"Kubeconfig"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTKEEdgeExternalKubeconfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTKEEdgeExternalKubeconfigResponseParams `json:"Response"`
}

func (r *DescribeTKEEdgeExternalKubeconfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTKEEdgeExternalKubeconfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTKEEdgeScriptRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// ENI
	Interface *string `json:"Interface,omitempty" name:"Interface"`

	// Name of the name
	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`

	// Node configuration in JSON format 
	Config *string `json:"Config,omitempty" name:"Config"`
}

type DescribeTKEEdgeScriptRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// ENI
	Interface *string `json:"Interface,omitempty" name:"Interface"`

	// Name of the name
	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`

	// Node configuration in JSON format 
	Config *string `json:"Config,omitempty" name:"Config"`
}

func (r *DescribeTKEEdgeScriptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTKEEdgeScriptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Interface")
	delete(f, "NodeName")
	delete(f, "Config")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTKEEdgeScriptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTKEEdgeScriptResponseParams struct {
	// Whether to download the link
	Link *string `json:"Link,omitempty" name:"Link"`

	// Whether to download the desired token
	Token *string `json:"Token,omitempty" name:"Token"`

	// Whether to download the command
	Command *string `json:"Command,omitempty" name:"Command"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTKEEdgeScriptResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTKEEdgeScriptResponseParams `json:"Response"`
}

func (r *DescribeTKEEdgeScriptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTKEEdgeScriptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVersionsRequestParams struct {

}

type DescribeVersionsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeVersionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVersionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVersionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVersionsResponseParams struct {
	// Number of versions
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Version list
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	VersionInstanceSet []*VersionInstance `json:"VersionInstanceSet,omitempty" name:"VersionInstanceSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVersionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVersionsResponseParams `json:"Response"`
}

func (r *DescribeVersionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcCniPodLimitsRequestParams struct {
	// The availability zone of the model to query, for example, `ap-guangzhou-3`. This field is left empty by default, that is, do not filter by the availability zone.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// The instance family to query, for example, `S5`. This field is left empty by default, that is, do not filter by the instance family.
	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`

	// The instance model to query, for example, `S5.LARGE8`. This field is empty by default, that is, do not filter by instance type.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
}

type DescribeVpcCniPodLimitsRequest struct {
	*tchttp.BaseRequest
	
	// The availability zone of the model to query, for example, `ap-guangzhou-3`. This field is left empty by default, that is, do not filter by the availability zone.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// The instance family to query, for example, `S5`. This field is left empty by default, that is, do not filter by the instance family.
	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`

	// The instance model to query, for example, `S5.LARGE8`. This field is empty by default, that is, do not filter by instance type.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
}

func (r *DescribeVpcCniPodLimitsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcCniPodLimitsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "InstanceFamily")
	delete(f, "InstanceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpcCniPodLimitsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcCniPodLimitsResponseParams struct {
	// The number of the models
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The model information and the maximum supported number of Pods in the VPC-CNI mode
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	PodLimitsInstanceSet []*PodLimitsInstance `json:"PodLimitsInstanceSet,omitempty" name:"PodLimitsInstanceSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVpcCniPodLimitsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpcCniPodLimitsResponseParams `json:"Response"`
}

func (r *DescribeVpcCniPodLimitsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcCniPodLimitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableClusterDeletionProtectionRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type DisableClusterDeletionProtectionRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DisableClusterDeletionProtectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableClusterDeletionProtectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableClusterDeletionProtectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableClusterDeletionProtectionResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DisableClusterDeletionProtectionResponse struct {
	*tchttp.BaseResponse
	Response *DisableClusterDeletionProtectionResponseParams `json:"Response"`
}

func (r *DisableClusterDeletionProtectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableClusterDeletionProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ECMEnhancedService struct {
	// Whether Cloud Monitoring is enabled
	SecurityService *ECMRunMonitorServiceEnabled `json:"SecurityService,omitempty" name:"SecurityService"`

	// Whether Cloud Workload Protection is enabled
	MonitorService *ECMRunSecurityServiceEnabled `json:"MonitorService,omitempty" name:"MonitorService"`
}

type ECMRunMonitorServiceEnabled struct {
	// Whether it is enabled
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

type ECMRunSecurityServiceEnabled struct {
	// Whether it is enabled
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`

	// CWP version. Valid values: `0` (CWP Pro), `1` (CWP Pro)
	Version *int64 `json:"Version,omitempty" name:"Version"`
}

type ECMZoneInstanceCountISP struct {
	// Instance AZ
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Number of instances to be created in the current AZ
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// ISP
	ISP *string `json:"ISP,omitempty" name:"ISP"`
}

type EdgeArgsFlag struct {
	// Parameter name
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Parameter type
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Parameter description
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Usage *string `json:"Usage,omitempty" name:"Usage"`

	// Default value of the parameter
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Default *string `json:"Default,omitempty" name:"Default"`

	// Valid value or range. Options: `[]` (it indicates a range, for example, “[1, 5]” indicates the parameter must be equal or larger than 1, and be equal or smaller than 5), and `()` (it indicates a valid value, for example, “('aa', 'bb')” indicates the parameter must be “aa” or “bb”. If it is left empty, the verification can be skipped.)
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Constraint *string `json:"Constraint,omitempty" name:"Constraint"`
}

type EdgeAvailableExtraArgs struct {
	// kube-apiserver custom parameter
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	KubeAPIServer []*EdgeArgsFlag `json:"KubeAPIServer,omitempty" name:"KubeAPIServer"`

	// kube-controller-manager custom parameter
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	KubeControllerManager []*EdgeArgsFlag `json:"KubeControllerManager,omitempty" name:"KubeControllerManager"`

	// kube-scheduler custom parameter
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	KubeScheduler []*EdgeArgsFlag `json:"KubeScheduler,omitempty" name:"KubeScheduler"`

	// kubelet custom parameter
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Kubelet []*EdgeArgsFlag `json:"Kubelet,omitempty" name:"Kubelet"`
}

type EdgeCluster struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Cluster name
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Cluster Pod CIDR block
	PodCIDR *string `json:"PodCIDR,omitempty" name:"PodCIDR"`

	// Cluster service CIDR block
	ServiceCIDR *string `json:"ServiceCIDR,omitempty" name:"ServiceCIDR"`


	K8SVersion *string `json:"K8SVersion,omitempty" name:"K8SVersion"`

	// Cluster status
	Status *string `json:"Status,omitempty" name:"Status"`

	// Cluster description
	ClusterDesc *string `json:"ClusterDesc,omitempty" name:"ClusterDesc"`

	// Cluster creation time
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Edge cluster version
	EdgeClusterVersion *string `json:"EdgeClusterVersion,omitempty" name:"EdgeClusterVersion"`

	// Maximum number of Pods on the node
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	MaxNodePodNum *int64 `json:"MaxNodePodNum,omitempty" name:"MaxNodePodNum"`
}

type EdgeClusterAdvancedSettings struct {
	// Custom parameters of the cluster
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ExtraArgs *EdgeClusterExtraArgs `json:"ExtraArgs,omitempty" name:"ExtraArgs"`
}

type EdgeClusterExtraArgs struct {
	// kube-apiserver custom parameter, in the format of ["k1=v1", "k1=v2"], for example: ["max-requests-inflight=500","feature-gates=PodShareProcessNamespace=true,DynamicKubeletConfig=true"]
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	KubeAPIServer []*string `json:"KubeAPIServer,omitempty" name:"KubeAPIServer"`

	// kube-controller-manager custom parameter
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	KubeControllerManager []*string `json:"KubeControllerManager,omitempty" name:"KubeControllerManager"`

	// kube-scheduler custom parameter
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	KubeScheduler []*string `json:"KubeScheduler,omitempty" name:"KubeScheduler"`
}

type EdgeClusterInternalLB struct {
	// Whether the private LB is enabled
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`

	// ID of the subnet associated with the private LB
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	SubnetId []*string `json:"SubnetId,omitempty" name:"SubnetId"`
}

type EdgeClusterPublicLB struct {
	// Whether the public LB is enabled
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`

	// Public network CIDR block allowed to access
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	AllowFromCidrs []*string `json:"AllowFromCidrs,omitempty" name:"AllowFromCidrs"`
}

// Predefined struct for user
type EnableClusterDeletionProtectionRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type EnableClusterDeletionProtectionRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *EnableClusterDeletionProtectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableClusterDeletionProtectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableClusterDeletionProtectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableClusterDeletionProtectionResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type EnableClusterDeletionProtectionResponse struct {
	*tchttp.BaseResponse
	Response *EnableClusterDeletionProtectionResponseParams `json:"Response"`
}

func (r *EnableClusterDeletionProtectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableClusterDeletionProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableVpcCniNetworkTypeRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The VPC-CNI mode. `tke-route-eni`: Multi-IP ENI, `tke-direct-eni`: Independent ENI
	VpcCniType *string `json:"VpcCniType,omitempty" name:"VpcCniType"`

	// Whether to enable static IP address
	EnableStaticIp *bool `json:"EnableStaticIp,omitempty" name:"EnableStaticIp"`

	// The container subnet being used
	Subnets []*string `json:"Subnets,omitempty" name:"Subnets"`

	// Specifies when to release the IP after the Pod termination in static IP mode. It must be longer than 300 seconds. If this parameter is left empty, the IP address will never be released.
	ExpiredSeconds *uint64 `json:"ExpiredSeconds,omitempty" name:"ExpiredSeconds"`
}

type EnableVpcCniNetworkTypeRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The VPC-CNI mode. `tke-route-eni`: Multi-IP ENI, `tke-direct-eni`: Independent ENI
	VpcCniType *string `json:"VpcCniType,omitempty" name:"VpcCniType"`

	// Whether to enable static IP address
	EnableStaticIp *bool `json:"EnableStaticIp,omitempty" name:"EnableStaticIp"`

	// The container subnet being used
	Subnets []*string `json:"Subnets,omitempty" name:"Subnets"`

	// Specifies when to release the IP after the Pod termination in static IP mode. It must be longer than 300 seconds. If this parameter is left empty, the IP address will never be released.
	ExpiredSeconds *uint64 `json:"ExpiredSeconds,omitempty" name:"ExpiredSeconds"`
}

func (r *EnableVpcCniNetworkTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableVpcCniNetworkTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "VpcCniType")
	delete(f, "EnableStaticIp")
	delete(f, "Subnets")
	delete(f, "ExpiredSeconds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableVpcCniNetworkTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableVpcCniNetworkTypeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type EnableVpcCniNetworkTypeResponse struct {
	*tchttp.BaseResponse
	Response *EnableVpcCniNetworkTypeResponseParams `json:"Response"`
}

func (r *EnableVpcCniNetworkTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableVpcCniNetworkTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnhancedService struct {
	// Enables cloud security service. If this parameter is not specified, the cloud security service will be enabled by default.
	SecurityService *RunSecurityServiceEnabled `json:"SecurityService,omitempty" name:"SecurityService"`

	// Enables cloud monitor service. If this parameter is not specified, the cloud monitor service will be enabled by default.
	MonitorService *RunMonitorServiceEnabled `json:"MonitorService,omitempty" name:"MonitorService"`

	// Enables the TAT service. If this parameter is not specified, the TAT service will not be enabled.
	AutomationService *RunAutomationServiceEnabled `json:"AutomationService,omitempty" name:"AutomationService"`
}

type ExistedInstance struct {
	// Whether the instance supports being added to the cluster (TRUE: support; FALSE: not support).
	// Note: This field may return null, indicating that no valid values can be obtained.
	Usable *bool `json:"Usable,omitempty" name:"Usable"`

	// Reason that the instance does not support being added.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UnusableReason *string `json:"UnusableReason,omitempty" name:"UnusableReason"`

	// ID of the cluster in which the instance currently resides.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AlreadyInCluster *string `json:"AlreadyInCluster,omitempty" name:"AlreadyInCluster"`

	// Instance ID, in the format of ins-xxxxxxxx.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// List of private IPs of the instance's primary ENI.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`

	// List of public IPs of the instance's primary ENI.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitempty" name:"PublicIpAddresses"`

	// Creation time, which follows the ISO8601 standard and uses UTC time. Format: YYYY-MM-DDThh:mm:ssZ.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Instance's number of CPU cores. Unit: cores.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CPU *uint64 `json:"CPU,omitempty" name:"CPU"`

	// Instance's memory capacity. Unit: GB.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// Operating system name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OsName *string `json:"OsName,omitempty" name:"OsName"`

	// Instance model.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Auto scaling group ID
	// Note: This field may return null, indicating that no valid value was found.
	AutoscalingGroupId *string `json:"AutoscalingGroupId,omitempty" name:"AutoscalingGroupId"`

	// Instance billing method. Valid values: POSTPAID_BY_HOUR (pay-as-you-go hourly); CDHPAID (billed based on CDH, i.e., only CDH is billed but not the instances on CDH)
	// Note: This field may return null, indicating that no valid value was found.
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// IPv6 address of the instance
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	IPv6Addresses []*string `json:"IPv6Addresses,omitempty" name:"IPv6Addresses"`
}

type ExistedInstancesForNode struct {
	// Node role. Values: MASTER_ETCD, WORKER. You only need to specify MASTER_ETCD when creating a self-deployed cluster (INDEPENDENT_CLUSTER).
	NodeRole *string `json:"NodeRole,omitempty" name:"NodeRole"`

	// Reinstallation parameter of existing instances
	ExistedInstancesPara *ExistedInstancesPara `json:"ExistedInstancesPara,omitempty" name:"ExistedInstancesPara"`

	// Advanced node setting, which overrides the InstanceAdvancedSettings item set at the cluster level (currently valid for the ExtraArgs node custom parameter only)
	InstanceAdvancedSettingsOverride *InstanceAdvancedSettings `json:"InstanceAdvancedSettingsOverride,omitempty" name:"InstanceAdvancedSettingsOverride"`

	// When the custom PodCIDR mode is enabled for the cluster, you can specify the maximum number of pods per node.
	DesiredPodNumbers []*int64 `json:"DesiredPodNumbers,omitempty" name:"DesiredPodNumbers"`
}

type ExistedInstancesPara struct {
	// Cluster ID
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Additional parameter to be set for the instance
	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitempty" name:"InstanceAdvancedSettings"`

	// Enhanced services. This parameter is used to specify whether to enable Cloud Security, Cloud Monitor and other services. If this parameter is not specified, Cloud Monitor and Cloud Security are enabled by default.
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`

	// Node login information (currently only supports using Password or single KeyIds)
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// Security group to which the instance belongs. This parameter can be obtained from the sgId field in the returned values of DescribeSecurityGroups. If this parameter is not specified, the default security group is bound. (Currently, you can only set a single sgId)
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// When reinstalling the system, you can specify the HostName of the modified instance (when the cluster is in HostName mode, this parameter is required, and the rule name is the same as the [Create CVM Instance](https://intl.cloud.tencent.com/document/product/213/15730?from_cn_redirect=1) API HostName except for uppercase letters not being supported.
	HostName *string `json:"HostName,omitempty" name:"HostName"`
}

type ExtensionAddon struct {
	// Add-on name
	AddonName *string `json:"AddonName,omitempty" name:"AddonName"`

	// Add-on information (description of the add-on resource object in JSON string format)
	AddonParam *string `json:"AddonParam,omitempty" name:"AddonParam"`
}

type Filter struct {
	// Filters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Filter values.
	Values []*string `json:"Values,omitempty" name:"Values"`
}

// Predefined struct for user
type ForwardTKEEdgeApplicationRequestV3RequestParams struct {
	// Access to request the cluster add-on
	Method *string `json:"Method,omitempty" name:"Method"`

	// Path to request the cluster add-on
	Path *string `json:"Path,omitempty" name:"Path"`

	// Data format allowed to receive the requested cluster add-on
	Accept *string `json:"Accept,omitempty" name:"Accept"`

	// Data format for requesting the cluster add-on
	ContentType *string `json:"ContentType,omitempty" name:"ContentType"`

	// Data sent to request the cluster add-on
	RequestBody *string `json:"RequestBody,omitempty" name:"RequestBody"`

	// Cluster name (for example, `cls-1234abcd`)
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// Whether to encode the request content
	EncodedBody *string `json:"EncodedBody,omitempty" name:"EncodedBody"`
}

type ForwardTKEEdgeApplicationRequestV3Request struct {
	*tchttp.BaseRequest
	
	// Access to request the cluster add-on
	Method *string `json:"Method,omitempty" name:"Method"`

	// Path to request the cluster add-on
	Path *string `json:"Path,omitempty" name:"Path"`

	// Data format allowed to receive the requested cluster add-on
	Accept *string `json:"Accept,omitempty" name:"Accept"`

	// Data format for requesting the cluster add-on
	ContentType *string `json:"ContentType,omitempty" name:"ContentType"`

	// Data sent to request the cluster add-on
	RequestBody *string `json:"RequestBody,omitempty" name:"RequestBody"`

	// Cluster name (for example, `cls-1234abcd`)
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// Whether to encode the request content
	EncodedBody *string `json:"EncodedBody,omitempty" name:"EncodedBody"`
}

func (r *ForwardTKEEdgeApplicationRequestV3Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ForwardTKEEdgeApplicationRequestV3Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Method")
	delete(f, "Path")
	delete(f, "Accept")
	delete(f, "ContentType")
	delete(f, "RequestBody")
	delete(f, "ClusterName")
	delete(f, "EncodedBody")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ForwardTKEEdgeApplicationRequestV3Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ForwardTKEEdgeApplicationRequestV3ResponseParams struct {
	// Data returned after requesting the cluster add-on
	ResponseBody *string `json:"ResponseBody,omitempty" name:"ResponseBody"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ForwardTKEEdgeApplicationRequestV3Response struct {
	*tchttp.BaseResponse
	Response *ForwardTKEEdgeApplicationRequestV3ResponseParams `json:"Response"`
}

func (r *ForwardTKEEdgeApplicationRequestV3Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ForwardTKEEdgeApplicationRequestV3Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetClusterLevelPriceRequestParams struct {
	// The cluster model. It’s used for price query.
	ClusterLevel *string `json:"ClusterLevel,omitempty" name:"ClusterLevel"`
}

type GetClusterLevelPriceRequest struct {
	*tchttp.BaseRequest
	
	// The cluster model. It’s used for price query.
	ClusterLevel *string `json:"ClusterLevel,omitempty" name:"ClusterLevel"`
}

func (r *GetClusterLevelPriceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetClusterLevelPriceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterLevel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetClusterLevelPriceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetClusterLevelPriceResponseParams struct {
	// Discount price (unit: US cent)
	Cost *uint64 `json:"Cost,omitempty" name:"Cost"`

	// Original price (unit: US cent)
	TotalCost *uint64 `json:"TotalCost,omitempty" name:"TotalCost"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetClusterLevelPriceResponse struct {
	*tchttp.BaseResponse
	Response *GetClusterLevelPriceResponseParams `json:"Response"`
}

func (r *GetClusterLevelPriceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetClusterLevelPriceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetUpgradeInstanceProgressRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Maximum number of nodes to be queried
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// The starting node for the query
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type GetUpgradeInstanceProgressRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Maximum number of nodes to be queried
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// The starting node for the query
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *GetUpgradeInstanceProgressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetUpgradeInstanceProgressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetUpgradeInstanceProgressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetUpgradeInstanceProgressResponseParams struct {
	// Total nodes to upgrade
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// Total upgraded nodes
	Done *int64 `json:"Done,omitempty" name:"Done"`

	// The lifecycle of the upgrade task
	// process: running
	// paused: stopped
	// pausing: stopping
	// done: completed
	// timeout: timed out
	// aborted: canceled
	LifeState *string `json:"LifeState,omitempty" name:"LifeState"`

	// Details of upgrade progress of each node
	Instances []*InstanceUpgradeProgressItem `json:"Instances,omitempty" name:"Instances"`

	// Current cluster status
	ClusterStatus *InstanceUpgradeClusterStatus `json:"ClusterStatus,omitempty" name:"ClusterStatus"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetUpgradeInstanceProgressResponse struct {
	*tchttp.BaseResponse
	Response *GetUpgradeInstanceProgressResponseParams `json:"Response"`
}

func (r *GetUpgradeInstanceProgressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetUpgradeInstanceProgressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IPAddress struct {
	// Type. Valid values: `advertise`, `public`, and others
	Type *string `json:"Type,omitempty" name:"Type"`

	// IP Address
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Network port
	Port *uint64 `json:"Port,omitempty" name:"Port"`
}

type ImageInstance struct {
	// Image alias
	// Note: this field may return null, indicating that no valid values can be obtained.
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// Operating system name
	// Note: this field may return null, indicating that no valid values can be obtained.
	OsName *string `json:"OsName,omitempty" name:"OsName"`

	// Image ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// Container image tag, **DOCKER_CUSTOMIZE** (container customized tag), **GENERAL** (general tag, default value)
	// Note: this field may return null, indicating that no valid values can be obtained.
	OsCustomizeType *string `json:"OsCustomizeType,omitempty" name:"OsCustomizeType"`
}

type Instance struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Node role: MASTER, WORKER, ETCD, MASTER_ETCD, and ALL. Default value: WORKER
	InstanceRole *string `json:"InstanceRole,omitempty" name:"InstanceRole"`

	// Reason for instance exception (or initialization)
	FailedReason *string `json:"FailedReason,omitempty" name:"FailedReason"`

	// Instance status (running, initializing, or failed)
	InstanceState *string `json:"InstanceState,omitempty" name:"InstanceState"`

	// Whether the instance is drained
	// Note: this field may return null, indicating that no valid value is obtained.
	DrainStatus *string `json:"DrainStatus,omitempty" name:"DrainStatus"`

	// Node settings
	// Note: this field may return null, indicating that no valid value is obtained.
	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitempty" name:"InstanceAdvancedSettings"`

	// Creation time
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Node private IP
	// Note: this field may return null, indicating that no valid values can be obtained.
	LanIP *string `json:"LanIP,omitempty" name:"LanIP"`

	// Resource pool ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	NodePoolId *string `json:"NodePoolId,omitempty" name:"NodePoolId"`

	// ID of the auto-scaling group
	// Note: this field may return null, indicating that no valid value is obtained.
	AutoscalingGroupId *string `json:"AutoscalingGroupId,omitempty" name:"AutoscalingGroupId"`
}

type InstanceAdvancedSettings struct {
	// Data disk mount point. By default, no data disk is mounted. Data disks in ext3, ext4, or XFS file system formats will be mounted directly, while data disks in other file systems and unformatted data disks will automatically be formatted as ext4 (xfs for tlinux system) and then mounted. Please back up your data in advance. This setting is only applicable to CVMs with a single data disk.
	// Note: in multi-disk scenarios, use the DataDisks data structure below to set the corresponding information, such as cloud disk type, cloud disk size, mount path, and whether to perform formatting.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	MountTarget *string `json:"MountTarget,omitempty" name:"MountTarget"`

	// Specified value of dockerd --graph. Default value: /var/lib/docker
	// Note: This field may return null, indicating that no valid value was found.
	DockerGraphPath *string `json:"DockerGraphPath,omitempty" name:"DockerGraphPath"`

	// Base64-encoded user script, which will be executed after the K8s component starts running. You need to ensure the reentrant and retry logic of the script. The script and its log files can be viewed at the node path: /data/ccs_userscript/. If you want to initialize nodes before adding them to the scheduling list, you can use this parameter together with the unschedulable parameter. After the final initialization of userScript is completed, add the kubectl uncordon nodename --kubeconfig=/root/.kube/config command to enable the node for scheduling.
	// Note: This field may return null, indicating that no valid value was found.
	UserScript *string `json:"UserScript,omitempty" name:"UserScript"`

	// Sets whether the added node is schedulable. 0 (default): schedulable; other values: unschedulable. After node initialization is completed, you can run kubectl uncordon nodename to enable this node for scheduling.
	Unschedulable *int64 `json:"Unschedulable,omitempty" name:"Unschedulable"`

	// Node label array
	// Note: This field may return null, indicating that no valid value was found.
	Labels []*Label `json:"Labels,omitempty" name:"Labels"`

	// Mounting information of multiple data disks. When you create a node, ensure that the CVM purchase parameter specifies the information required for the purchase of multiple data disks. For example, the `DataDisks` under `RunInstancesPara` of the `CreateClusterInstances` API should be configured accordingly (Referto document of CreateClusterInstances API). When you add an existing node, ensure that the specified partition exists in the node.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`

	// Information about node custom parameters
	// Note: This field may return null, indicating that no valid value was found.
	ExtraArgs *InstanceExtraArgs `json:"ExtraArgs,omitempty" name:"ExtraArgs"`

	// When the custom PodCIDR mode is enabled for the cluster, you can specify the maximum number of pods per node.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DesiredPodNumber *int64 `json:"DesiredPodNumber,omitempty" name:"DesiredPodNumber"`

	// Specifies the base64-encoded custom script to be executed before initialization of the node. It’s only valid for adding existing nodes for now.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	PreStartUserScript *string `json:"PreStartUserScript,omitempty" name:"PreStartUserScript"`

	// Node taint
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Taints []*Taint `json:"Taints,omitempty" name:"Taints"`
}

type InstanceDataDiskMountSetting struct {
	// CVM instance type
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Data disk mounting information
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`

	// Availability zone where the CVM instance is located
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type InstanceExtraArgs struct {
	// Kubelet custom parameter, in the format of ["k1=v1", "k1=v2"], for example: ["root-dir=/var/lib/kubelet","feature-gates=PodShareProcessNamespace=true,DynamicKubeletConfig=true"].
	// Note: this field may return `null`, indicating that no valid value is obtained.
	Kubelet []*string `json:"Kubelet,omitempty" name:"Kubelet"`
}

type InstanceUpgradeClusterStatus struct {
	// Total Pods
	PodTotal *int64 `json:"PodTotal,omitempty" name:"PodTotal"`

	// Total number of NotReady Pods
	NotReadyPod *int64 `json:"NotReadyPod,omitempty" name:"NotReadyPod"`
}

type InstanceUpgradePreCheckResult struct {
	// Whether the check is passed
	CheckPass *bool `json:"CheckPass,omitempty" name:"CheckPass"`

	// Array of check items
	Items []*InstanceUpgradePreCheckResultItem `json:"Items,omitempty" name:"Items"`

	// List of independent pods on this node
	SinglePods []*string `json:"SinglePods,omitempty" name:"SinglePods"`
}

type InstanceUpgradePreCheckResultItem struct {
	// The namespace of the workload
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// Workload type
	WorkLoadKind *string `json:"WorkLoadKind,omitempty" name:"WorkLoadKind"`

	// Workload name
	WorkLoadName *string `json:"WorkLoadName,omitempty" name:"WorkLoadName"`

	// The number of running pods in the workload before draining the node
	Before *uint64 `json:"Before,omitempty" name:"Before"`

	// The number of running pods in the workload after draining the node
	After *uint64 `json:"After,omitempty" name:"After"`

	// The pod list of the workload on this node
	Pods []*string `json:"Pods,omitempty" name:"Pods"`
}

type InstanceUpgradeProgressItem struct {
	// Node instance ID
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`

	// Task lifecycle
	// process: running
	// paused: stopped
	// pausing: stopping
	// done: completed
	// timeout: timed out
	// aborted: canceled
	// pending: not started
	LifeState *string `json:"LifeState,omitempty" name:"LifeState"`

	// Upgrade start time
	// Note: this field may return `null`, indicating that no valid value is obtained.
	StartAt *string `json:"StartAt,omitempty" name:"StartAt"`

	// Upgrade end time
	// Note: this field may return `null`, indicating that no valid value is obtained.
	EndAt *string `json:"EndAt,omitempty" name:"EndAt"`

	// Check result before upgrading
	CheckResult *InstanceUpgradePreCheckResult `json:"CheckResult,omitempty" name:"CheckResult"`

	// Upgrade steps details
	Detail []*TaskStepInfo `json:"Detail,omitempty" name:"Detail"`
}

type Label struct {
	// Name in map list
	Name *string `json:"Name,omitempty" name:"Name"`

	// Value in map list
	Value *string `json:"Value,omitempty" name:"Value"`
}

type LoginSettings struct {
	// Login password of the instance. <br><li>For Linux instances, the password must include 8-30 characters, and contain at least two of the following character sets: [a-z], [A-Z], [0-9] and [()\`~!@#$%^&*-+=|{}[]:;',.?/]. <br><li>For Windows instances, the password must include 12-30 characters, and contain at least three of the following character sets: [a-z], [A-Z], [0-9] and [()\`~!@#$%^&*-+=|{}[]:;',.?/]. <br><br>If it's not specified, the user needs to set the login password using the **Reset password** option in the CVM console or calling the API `ResetInstancesPassword` to complete the creation of the CVM instance(s).
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Password *string `json:"Password,omitempty" name:"Password"`

	// List of key IDs. After an instance is associated with a key, you can access the instance with the private key in the key pair. You can call [`DescribeKeyPairs`](https://intl.cloud.tencent.com/document/api/213/15699?from_cn_redirect=1) to obtain `KeyId`. A key and password cannot be specified at the same time. Windows instances do not support keys. Currently, you can only specify one key when purchasing an instance.
	// Note: this field may return null, indicating that no valid value is obtained.
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`

	// Whether to keep the original settings of an image. You cannot specify this parameter and `Password` or `KeyIds.N` at the same time. You can specify this parameter as `TRUE` only when you create an instance using a custom image, a shared image, or an imported image. Valid values: <br><li>TRUE: keep the login settings of the image <br><li>FALSE: do not keep the login settings of the image <br><br>Default value: FALSE.
	// Note: This field may return null, indicating that no valid value is found.
	KeepImageLogin *string `json:"KeepImageLogin,omitempty" name:"KeepImageLogin"`
}

type ManuallyAdded struct {
	// Number of nodes that are being added
	Joining *int64 `json:"Joining,omitempty" name:"Joining"`

	// Number of nodes that are being initialized
	Initializing *int64 `json:"Initializing,omitempty" name:"Initializing"`

	// Number of normal nodes
	Normal *int64 `json:"Normal,omitempty" name:"Normal"`

	// Total number of nodes
	Total *int64 `json:"Total,omitempty" name:"Total"`
}

// Predefined struct for user
type ModifyClusterAsGroupAttributeRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Cluster-associated scaling group attributes
	ClusterAsGroupAttribute *ClusterAsGroupAttribute `json:"ClusterAsGroupAttribute,omitempty" name:"ClusterAsGroupAttribute"`
}

type ModifyClusterAsGroupAttributeRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Cluster-associated scaling group attributes
	ClusterAsGroupAttribute *ClusterAsGroupAttribute `json:"ClusterAsGroupAttribute,omitempty" name:"ClusterAsGroupAttribute"`
}

func (r *ModifyClusterAsGroupAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterAsGroupAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ClusterAsGroupAttribute")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterAsGroupAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterAsGroupAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyClusterAsGroupAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterAsGroupAttributeResponseParams `json:"Response"`
}

func (r *ModifyClusterAsGroupAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterAsGroupAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterAsGroupOptionAttributeRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Cluster auto scaling attributes
	ClusterAsGroupOption *ClusterAsGroupOption `json:"ClusterAsGroupOption,omitempty" name:"ClusterAsGroupOption"`
}

type ModifyClusterAsGroupOptionAttributeRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Cluster auto scaling attributes
	ClusterAsGroupOption *ClusterAsGroupOption `json:"ClusterAsGroupOption,omitempty" name:"ClusterAsGroupOption"`
}

func (r *ModifyClusterAsGroupOptionAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterAsGroupOptionAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ClusterAsGroupOption")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterAsGroupOptionAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterAsGroupOptionAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyClusterAsGroupOptionAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterAsGroupOptionAttributeResponseParams `json:"Response"`
}

func (r *ModifyClusterAsGroupOptionAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterAsGroupOptionAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterAttributeRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Project of the Cluster
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Cluster name
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// Cluster description
	ClusterDesc *string `json:"ClusterDesc,omitempty" name:"ClusterDesc"`

	// Cluster specification
	ClusterLevel *string `json:"ClusterLevel,omitempty" name:"ClusterLevel"`

	// Auto-upgrades cluster specification
	AutoUpgradeClusterLevel *AutoUpgradeClusterLevel `json:"AutoUpgradeClusterLevel,omitempty" name:"AutoUpgradeClusterLevel"`

	// Whether to enable qGPU Sharing
	QGPUShareEnable *bool `json:"QGPUShareEnable,omitempty" name:"QGPUShareEnable"`
}

type ModifyClusterAttributeRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Project of the Cluster
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Cluster name
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// Cluster description
	ClusterDesc *string `json:"ClusterDesc,omitempty" name:"ClusterDesc"`

	// Cluster specification
	ClusterLevel *string `json:"ClusterLevel,omitempty" name:"ClusterLevel"`

	// Auto-upgrades cluster specification
	AutoUpgradeClusterLevel *AutoUpgradeClusterLevel `json:"AutoUpgradeClusterLevel,omitempty" name:"AutoUpgradeClusterLevel"`

	// Whether to enable qGPU Sharing
	QGPUShareEnable *bool `json:"QGPUShareEnable,omitempty" name:"QGPUShareEnable"`
}

func (r *ModifyClusterAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ProjectId")
	delete(f, "ClusterName")
	delete(f, "ClusterDesc")
	delete(f, "ClusterLevel")
	delete(f, "AutoUpgradeClusterLevel")
	delete(f, "QGPUShareEnable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterAttributeResponseParams struct {
	// Project of the Cluster
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Cluster name
	// Note: this field may return null, indicating that no valid values can be obtained.
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// Cluster description
	// Note: this field may return null, indicating that no valid values can be obtained.
	ClusterDesc *string `json:"ClusterDesc,omitempty" name:"ClusterDesc"`

	// Cluster specification
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ClusterLevel *string `json:"ClusterLevel,omitempty" name:"ClusterLevel"`

	// Auto-upgrades cluster specification
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	AutoUpgradeClusterLevel *AutoUpgradeClusterLevel `json:"AutoUpgradeClusterLevel,omitempty" name:"AutoUpgradeClusterLevel"`

	// Whether to enable qGPU Sharing
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	QGPUShareEnable *bool `json:"QGPUShareEnable,omitempty" name:"QGPUShareEnable"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyClusterAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterAttributeResponseParams `json:"Response"`
}

func (r *ModifyClusterAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterAuthenticationOptionsRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// ServiceAccount authentication configuration
	ServiceAccounts *ServiceAccountAuthenticationOptions `json:"ServiceAccounts,omitempty" name:"ServiceAccounts"`
}

type ModifyClusterAuthenticationOptionsRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// ServiceAccount authentication configuration
	ServiceAccounts *ServiceAccountAuthenticationOptions `json:"ServiceAccounts,omitempty" name:"ServiceAccounts"`
}

func (r *ModifyClusterAuthenticationOptionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterAuthenticationOptionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ServiceAccounts")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterAuthenticationOptionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterAuthenticationOptionsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyClusterAuthenticationOptionsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterAuthenticationOptionsResponseParams `json:"Response"`
}

func (r *ModifyClusterAuthenticationOptionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterAuthenticationOptionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterEndpointSPRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Security policy opens single IP or CIDR block to the Internet (for example: '192.168.1.0/24', with 'reject all' as the default).
	SecurityPolicies []*string `json:"SecurityPolicies,omitempty" name:"SecurityPolicies"`
}

type ModifyClusterEndpointSPRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Security policy opens single IP or CIDR block to the Internet (for example: '192.168.1.0/24', with 'reject all' as the default).
	SecurityPolicies []*string `json:"SecurityPolicies,omitempty" name:"SecurityPolicies"`
}

func (r *ModifyClusterEndpointSPRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterEndpointSPRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SecurityPolicies")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterEndpointSPRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterEndpointSPResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyClusterEndpointSPResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterEndpointSPResponseParams `json:"Response"`
}

func (r *ModifyClusterEndpointSPResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterEndpointSPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterNodePoolRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Node pool ID
	NodePoolId *string `json:"NodePoolId,omitempty" name:"NodePoolId"`

	// Name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Maximum number of nodes
	MaxNodesNum *int64 `json:"MaxNodesNum,omitempty" name:"MaxNodesNum"`

	// Minimum number of nodes
	MinNodesNum *int64 `json:"MinNodesNum,omitempty" name:"MinNodesNum"`

	// Labels
	Labels []*Label `json:"Labels,omitempty" name:"Labels"`

	// Taints
	Taints []*Taint `json:"Taints,omitempty" name:"Taints"`

	// Indicates whether auto scaling is enabled.
	EnableAutoscale *bool `json:"EnableAutoscale,omitempty" name:"EnableAutoscale"`

	// Operating system name
	OsName *string `json:"OsName,omitempty" name:"OsName"`

	// Image tag, `DOCKER_CUSTOMIZE` (container customized tag), `GENERAL` (general tag, default value)
	OsCustomizeType *string `json:"OsCustomizeType,omitempty" name:"OsCustomizeType"`

	// Node custom parameter
	ExtraArgs *InstanceExtraArgs `json:"ExtraArgs,omitempty" name:"ExtraArgs"`

	// Resource tag
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Sets whether the added node is schedulable. 0 (default): schedulable; other values: unschedulable. After node initialization is completed, you can run `kubectl uncordon nodename` to enable this node for scheduling.
	Unschedulable *int64 `json:"Unschedulable,omitempty" name:"Unschedulable"`

	// Whether Deletion Protection is enabled
	DeletionProtection *bool `json:"DeletionProtection,omitempty" name:"DeletionProtection"`
}

type ModifyClusterNodePoolRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Node pool ID
	NodePoolId *string `json:"NodePoolId,omitempty" name:"NodePoolId"`

	// Name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Maximum number of nodes
	MaxNodesNum *int64 `json:"MaxNodesNum,omitempty" name:"MaxNodesNum"`

	// Minimum number of nodes
	MinNodesNum *int64 `json:"MinNodesNum,omitempty" name:"MinNodesNum"`

	// Labels
	Labels []*Label `json:"Labels,omitempty" name:"Labels"`

	// Taints
	Taints []*Taint `json:"Taints,omitempty" name:"Taints"`

	// Indicates whether auto scaling is enabled.
	EnableAutoscale *bool `json:"EnableAutoscale,omitempty" name:"EnableAutoscale"`

	// Operating system name
	OsName *string `json:"OsName,omitempty" name:"OsName"`

	// Image tag, `DOCKER_CUSTOMIZE` (container customized tag), `GENERAL` (general tag, default value)
	OsCustomizeType *string `json:"OsCustomizeType,omitempty" name:"OsCustomizeType"`

	// Node custom parameter
	ExtraArgs *InstanceExtraArgs `json:"ExtraArgs,omitempty" name:"ExtraArgs"`

	// Resource tag
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Sets whether the added node is schedulable. 0 (default): schedulable; other values: unschedulable. After node initialization is completed, you can run `kubectl uncordon nodename` to enable this node for scheduling.
	Unschedulable *int64 `json:"Unschedulable,omitempty" name:"Unschedulable"`

	// Whether Deletion Protection is enabled
	DeletionProtection *bool `json:"DeletionProtection,omitempty" name:"DeletionProtection"`
}

func (r *ModifyClusterNodePoolRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterNodePoolRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NodePoolId")
	delete(f, "Name")
	delete(f, "MaxNodesNum")
	delete(f, "MinNodesNum")
	delete(f, "Labels")
	delete(f, "Taints")
	delete(f, "EnableAutoscale")
	delete(f, "OsName")
	delete(f, "OsCustomizeType")
	delete(f, "ExtraArgs")
	delete(f, "Tags")
	delete(f, "Unschedulable")
	delete(f, "DeletionProtection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterNodePoolRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterNodePoolResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyClusterNodePoolResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterNodePoolResponseParams `json:"Response"`
}

func (r *ModifyClusterNodePoolResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterNodePoolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNodePoolInstanceTypesRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Node pool ID
	NodePoolId *string `json:"NodePoolId,omitempty" name:"NodePoolId"`

	// List of instance types
	InstanceTypes []*string `json:"InstanceTypes,omitempty" name:"InstanceTypes"`
}

type ModifyNodePoolInstanceTypesRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Node pool ID
	NodePoolId *string `json:"NodePoolId,omitempty" name:"NodePoolId"`

	// List of instance types
	InstanceTypes []*string `json:"InstanceTypes,omitempty" name:"InstanceTypes"`
}

func (r *ModifyNodePoolInstanceTypesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNodePoolInstanceTypesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NodePoolId")
	delete(f, "InstanceTypes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNodePoolInstanceTypesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNodePoolInstanceTypesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyNodePoolInstanceTypesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNodePoolInstanceTypesResponseParams `json:"Response"`
}

func (r *ModifyNodePoolInstanceTypesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNodePoolInstanceTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusAlertRuleRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Alarm configurations
	AlertRule *PrometheusAlertRuleDetail `json:"AlertRule,omitempty" name:"AlertRule"`
}

type ModifyPrometheusAlertRuleRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Alarm configurations
	AlertRule *PrometheusAlertRuleDetail `json:"AlertRule,omitempty" name:"AlertRule"`
}

func (r *ModifyPrometheusAlertRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusAlertRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AlertRule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPrometheusAlertRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusAlertRuleResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyPrometheusAlertRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPrometheusAlertRuleResponseParams `json:"Response"`
}

func (r *ModifyPrometheusAlertRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusAlertRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NodeCountSummary struct {
	// Nodes that are manually managed
	// Note: this field may return `null`, indicating that no valid value is obtained.
	ManuallyAdded *ManuallyAdded `json:"ManuallyAdded,omitempty" name:"ManuallyAdded"`

	// Nodes that are automatically managed
	// Note: this field may return `null`, indicating that no valid value is obtained.
	AutoscalingAdded *AutoscalingAdded `json:"AutoscalingAdded,omitempty" name:"AutoscalingAdded"`
}

type NodePool struct {
	// Node pool ID
	NodePoolId *string `json:"NodePoolId,omitempty" name:"NodePoolId"`

	// Node pool name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Cluster instance ID
	ClusterInstanceId *string `json:"ClusterInstanceId,omitempty" name:"ClusterInstanceId"`

	// The lifecycle state of the current node pool. Valid values: `creating`, `normal`, `updating`, `deleting`, and `deleted`.
	LifeState *string `json:"LifeState,omitempty" name:"LifeState"`

	// Launch configuration ID
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitempty" name:"LaunchConfigurationId"`

	// Auto-scaling group ID
	AutoscalingGroupId *string `json:"AutoscalingGroupId,omitempty" name:"AutoscalingGroupId"`

	// Labels
	Labels []*Label `json:"Labels,omitempty" name:"Labels"`

	// Array of taint
	Taints []*Taint `json:"Taints,omitempty" name:"Taints"`

	// Node list
	NodeCountSummary *NodeCountSummary `json:"NodeCountSummary,omitempty" name:"NodeCountSummary"`


	AutoscalingGroupStatus *string `json:"AutoscalingGroupStatus,omitempty" name:"AutoscalingGroupStatus"`

	// Maximum number of nodes
	// Note: this field may return `null`, indicating that no valid value is obtained.
	MaxNodesNum *int64 `json:"MaxNodesNum,omitempty" name:"MaxNodesNum"`

	// Minimum number of nodes
	// Note: this field may return `null`, indicating that no valid value is obtained.
	MinNodesNum *int64 `json:"MinNodesNum,omitempty" name:"MinNodesNum"`

	// Desired number of nodes
	// Note: this field may return `null`, indicating that no valid value is obtained.
	DesiredNodesNum *int64 `json:"DesiredNodesNum,omitempty" name:"DesiredNodesNum"`

	// The operating system of the node pool
	// Note: this field may return `null`, indicating that no valid value is obtained.
	NodePoolOs *string `json:"NodePoolOs,omitempty" name:"NodePoolOs"`

	// Container image tag, `DOCKER_CUSTOMIZE` (container customized tag), `GENERAL` (general tag, default value)
	// Note: this field may return `null`, indicating that no valid value is obtained.
	OsCustomizeType *string `json:"OsCustomizeType,omitempty" name:"OsCustomizeType"`

	// Image ID
	// Note: this field may return `null`, indicating that no valid value is obtained.
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// This parameter is required when the custom PodCIDR mode is enabled for the cluster.
	// Note: this field may return `null`, indicating that no valid value is obtained.
	DesiredPodNum *int64 `json:"DesiredPodNum,omitempty" name:"DesiredPodNum"`

	// Custom script
	// Note: this field may return `null`, indicating that no valid value is obtained.
	UserScript *string `json:"UserScript,omitempty" name:"UserScript"`

	// Resource tag
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Whether Deletion Protection is enabled
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DeletionProtection *bool `json:"DeletionProtection,omitempty" name:"DeletionProtection"`
}

type NodePoolOption struct {
	// Whether to add to the node pool.
	AddToNodePool *bool `json:"AddToNodePool,omitempty" name:"AddToNodePool"`

	// Node pool ID
	NodePoolId *string `json:"NodePoolId,omitempty" name:"NodePoolId"`

	// Whether to inherit the node pool configuration.
	InheritConfigurationFromNodePool *bool `json:"InheritConfigurationFromNodePool,omitempty" name:"InheritConfigurationFromNodePool"`
}

type PodLimitsByType struct {
	// The number of Pods supported by a TKE shared ENI in non-static IP address mode
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	TKERouteENINonStaticIP *int64 `json:"TKERouteENINonStaticIP,omitempty" name:"TKERouteENINonStaticIP"`

	// The number of Pods supported by a TKE shared ENI in static IP address mode
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	TKERouteENIStaticIP *int64 `json:"TKERouteENIStaticIP,omitempty" name:"TKERouteENIStaticIP"`

	// The number of Pods supported by TKE independent ENI mode
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	TKEDirectENI *int64 `json:"TKEDirectENI,omitempty" name:"TKEDirectENI"`
}

type PodLimitsInstance struct {
	// The availability zone where the model is located
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// The instance family to which the model belongs
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`

	// Instance type
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// The maximum number of Pods in the VPC-CNI mode supported by the model
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	PodLimits *PodLimitsByType `json:"PodLimits,omitempty" name:"PodLimits"`
}

type PrometheusAlertRule struct {
	// Rule name
	Name *string `json:"Name,omitempty" name:"Name"`

	// PromQL contents
	Rule *string `json:"Rule,omitempty" name:"Rule"`

	// Additional labels
	Labels []*Label `json:"Labels,omitempty" name:"Labels"`

	// Alarm delivery template
	Template *string `json:"Template,omitempty" name:"Template"`

	// Duration
	For *string `json:"For,omitempty" name:"For"`

	// Rule description
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	Describe *string `json:"Describe,omitempty" name:"Describe"`

	// Refer to annotations in prometheus rule
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	Annotations []*Label `json:"Annotations,omitempty" name:"Annotations"`

	// Alarm rule status
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	RuleState *int64 `json:"RuleState,omitempty" name:"RuleState"`
}

type PrometheusAlertRuleDetail struct {
	// Rule name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Rule list
	Rules []*PrometheusAlertRule `json:"Rules,omitempty" name:"Rules"`

	// Last modification time
	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`

	// Alarm delivery methods
	Notification *PrometheusNotification `json:"Notification,omitempty" name:"Notification"`

	// Alarm rule ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// If the alarm is delivered via a template, the TemplateId is the template ID.
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// Alarm interval
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	Interval *string `json:"Interval,omitempty" name:"Interval"`
}

type PrometheusGrafanaInfo struct {
	// Whether it is enabled
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`

	// Domain name. It will be effective only when the public network access is enabled.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// The private network or public network address
	Address *string `json:"Address,omitempty" name:"Address"`

	// Whether the public network access is enabled.
	// `close`: the public network access is not enabled
	// `opening`: the public network access is being enabled
	// `open`: the public network access is enabled
	Internet *string `json:"Internet,omitempty" name:"Internet"`

	// The user name of the grafana admin
	AdminUser *string `json:"AdminUser,omitempty" name:"AdminUser"`
}

type PrometheusNotification struct {
	// Whether it is enabled
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`

	// Convergence time
	RepeatInterval *string `json:"RepeatInterval,omitempty" name:"RepeatInterval"`

	// Start time
	TimeRangeStart *string `json:"TimeRangeStart,omitempty" name:"TimeRangeStart"`

	// End time
	TimeRangeEnd *string `json:"TimeRangeEnd,omitempty" name:"TimeRangeEnd"`

	// Alarm delivery method. Valid values: `SMS`, `EMAIL`, `CALL`, and `WECHAT`
	// It respectively represents SMS, email, phone calls, and WeChat.
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	NotifyWay []*string `json:"NotifyWay,omitempty" name:"NotifyWay"`

	// The alarm recipient group (user group)
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	ReceiverGroups []*uint64 `json:"ReceiverGroups,omitempty" name:"ReceiverGroups"`

	// The alarm sequence of phone calls
	// This parameter is used when you specify `CALL` for `NotifyWay`.
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	PhoneNotifyOrder []*uint64 `json:"PhoneNotifyOrder,omitempty" name:"PhoneNotifyOrder"`

	// The number of phone call alarms
	// This parameter is used when you specify `CALL` for `NotifyWay`.
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	PhoneCircleTimes *int64 `json:"PhoneCircleTimes,omitempty" name:"PhoneCircleTimes"`

	// Dialing interval in seconds within one polling
	// This parameter is used when you specify `CALL` for `NotifyWay`.
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	PhoneInnerInterval *int64 `json:"PhoneInnerInterval,omitempty" name:"PhoneInnerInterval"`

	// Polling interval in seconds
	// This parameter is used when you specify `CALL` for `NotifyWay`.
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	PhoneCircleInterval *int64 `json:"PhoneCircleInterval,omitempty" name:"PhoneCircleInterval"`

	// Phone call alarm arrival notification
	// This parameter is used when you specify `CALL` for `NotifyWay`.
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	PhoneArriveNotice *bool `json:"PhoneArriveNotice,omitempty" name:"PhoneArriveNotice"`

	// Channel type. Default value: `amp`. The following channels are supported:
	// amp
	// webhook
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	Type *string `json:"Type,omitempty" name:"Type"`

	// This parameter is required if `Type` is `webhook`.
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	WebHook *string `json:"WebHook,omitempty" name:"WebHook"`
}

type RegionInstance struct {
	// Region name
	// Note: this field may return null, indicating that no valid values can be obtained.
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// Region ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// Region status
	// Note: this field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Status of region-related features (return all attributes in JSON format)
	// Note: this field may return null, indicating that no valid values can be obtained.
	FeatureGates *string `json:"FeatureGates,omitempty" name:"FeatureGates"`

	// Region abbreviation
	// Note: this field may return null, indicating that no valid values can be obtained.
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// Whitelisted location
	// Note: this field may return null, indicating that no valid values can be obtained.
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

// Predefined struct for user
type RemoveNodeFromNodePoolRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Node pool ID
	NodePoolId *string `json:"NodePoolId,omitempty" name:"NodePoolId"`

	// The node ID list. Up to 100 nodes can be removed at a time.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type RemoveNodeFromNodePoolRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Node pool ID
	NodePoolId *string `json:"NodePoolId,omitempty" name:"NodePoolId"`

	// The node ID list. Up to 100 nodes can be removed at a time.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *RemoveNodeFromNodePoolRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveNodeFromNodePoolRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NodePoolId")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveNodeFromNodePoolRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveNodeFromNodePoolResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RemoveNodeFromNodePoolResponse struct {
	*tchttp.BaseResponse
	Response *RemoveNodeFromNodePoolResponseParams `json:"Response"`
}

func (r *RemoveNodeFromNodePoolResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveNodeFromNodePoolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceDeleteOption struct {
	// Resource type, for example `CBS`
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// Specifies the policy to deal with resources in the cluster when the cluster is deleted. It can be `terminate` or `retain`.
	DeleteMode *string `json:"DeleteMode,omitempty" name:"DeleteMode"`
}

type ResourceUsage struct {
	// Resource type
	Name *string `json:"Name,omitempty" name:"Name"`

	// Resource usage
	Usage *uint64 `json:"Usage,omitempty" name:"Usage"`

	// Resource usage details
	Details []*ResourceUsageDetail `json:"Details,omitempty" name:"Details"`
}

type ResourceUsageDetail struct {
	// Resource name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Resource usage
	Usage *uint64 `json:"Usage,omitempty" name:"Usage"`
}

type RouteInfo struct {
	// Route table name.
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`

	// Destination CIDR.
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" name:"DestinationCidrBlock"`

	// Next hop address.
	GatewayIp *string `json:"GatewayIp,omitempty" name:"GatewayIp"`
}

type RouteTableConflict struct {
	// Route table type.
	RouteTableType *string `json:"RouteTableType,omitempty" name:"RouteTableType"`

	// Route table CIDR.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RouteTableCidrBlock *string `json:"RouteTableCidrBlock,omitempty" name:"RouteTableCidrBlock"`

	// Route table name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`

	// Route table ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
}

type RouteTableInfo struct {
	// Route table name.
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`

	// Route table CIDR.
	RouteTableCidrBlock *string `json:"RouteTableCidrBlock,omitempty" name:"RouteTableCidrBlock"`

	// VPC instance ID.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

type RunAutomationServiceEnabled struct {
	// Whether to enable the TAT service. Valid values: <br><li>`TRUE`: yes;<br><li>`FALSE`: no<br><br>Default: `FALSE`.
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

type RunInstancesForNode struct {
	// Node role. Values: MASTER_ETCD, WORKER. You only need to specify MASTER_ETCD when creating a self-deployed cluster (INDEPENDENT_CLUSTER).
	NodeRole *string `json:"NodeRole,omitempty" name:"NodeRole"`

	// Pass-through parameter for CVM creation in the format of a JSON string. For more information, see the API for [creating a CVM instance](https://intl.cloud.tencent.com/document/product/213/15730?from_cn_redirect=1). Pass any parameter other than common parameters. ImageId will be replaced with the image corresponding to the TKE cluster operating system.
	RunInstancesPara []*string `json:"RunInstancesPara,omitempty" name:"RunInstancesPara"`

	// An advanced node setting. This parameter overrides the InstanceAdvancedSettings item set at the cluster level and corresponds to RunInstancesPara in a one-to-one sequential manner (currently valid for the ExtraArgs node custom parameter only).
	InstanceAdvancedSettingsOverrides []*InstanceAdvancedSettings `json:"InstanceAdvancedSettingsOverrides,omitempty" name:"InstanceAdvancedSettingsOverrides"`
}

type RunMonitorServiceEnabled struct {
	// Whether to enable [Cloud Monitor](https://intl.cloud.tencent.com/document/product/248?from_cn_redirect=1). Valid values: <br><li>TRUE: enable Cloud Monitor <br><li>FALSE: do not enable Cloud Monitor <br><br>Default value: TRUE.
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

type RunSecurityServiceEnabled struct {
	// Whether to enable [Cloud Security](https://intl.cloud.tencent.com/document/product/296?from_cn_redirect=1). Valid values: <br><li>TRUE: enable Cloud Security <br><li>FALSE: do not enable Cloud Security <br><br>Default value: TRUE.
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

type ServiceAccountAuthenticationOptions struct {
	// service-account-issuer
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Issuer *string `json:"Issuer,omitempty" name:"Issuer"`

	// service-account-jwks-uri
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	JWKSURI *string `json:"JWKSURI,omitempty" name:"JWKSURI"`

	// If it is set to `true`, a RABC rule is automatically created to allow anonymous users to access `/.well-known/openid-configuration` and `/openid/v1/jwks`.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	AutoCreateDiscoveryAnonymousAuth *bool `json:"AutoCreateDiscoveryAnonymousAuth,omitempty" name:"AutoCreateDiscoveryAnonymousAuth"`
}

// Predefined struct for user
type SetNodePoolNodeProtectionRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Node pool ID
	NodePoolId *string `json:"NodePoolId,omitempty" name:"NodePoolId"`

	// Node ID
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Whether the node needs removal protection
	ProtectedFromScaleIn *bool `json:"ProtectedFromScaleIn,omitempty" name:"ProtectedFromScaleIn"`
}

type SetNodePoolNodeProtectionRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Node pool ID
	NodePoolId *string `json:"NodePoolId,omitempty" name:"NodePoolId"`

	// Node ID
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Whether the node needs removal protection
	ProtectedFromScaleIn *bool `json:"ProtectedFromScaleIn,omitempty" name:"ProtectedFromScaleIn"`
}

func (r *SetNodePoolNodeProtectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetNodePoolNodeProtectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NodePoolId")
	delete(f, "InstanceIds")
	delete(f, "ProtectedFromScaleIn")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetNodePoolNodeProtectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetNodePoolNodeProtectionResponseParams struct {
	// ID of the node that has successfully set the removal protection
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SucceedInstanceIds []*string `json:"SucceedInstanceIds,omitempty" name:"SucceedInstanceIds"`

	// ID of the node that fails to set the removal protection
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	FailedInstanceIds []*string `json:"FailedInstanceIds,omitempty" name:"FailedInstanceIds"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetNodePoolNodeProtectionResponse struct {
	*tchttp.BaseResponse
	Response *SetNodePoolNodeProtectionResponseParams `json:"Response"`
}

func (r *SetNodePoolNodeProtectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetNodePoolNodeProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {
	// Tag key.
	Key *string `json:"Key,omitempty" name:"Key"`

	// Tag value.
	Value *string `json:"Value,omitempty" name:"Value"`
}

type TagSpecification struct {
	// The type of resource that the tag is bound to. The type currently supported is `cluster`.
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// List of tag pairs
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type Taint struct {
	// Key of the taint
	Key *string `json:"Key,omitempty" name:"Key"`

	// Value of the taint
	Value *string `json:"Value,omitempty" name:"Value"`

	// Effect of the taint
	Effect *string `json:"Effect,omitempty" name:"Effect"`
}

type TaskStepInfo struct {
	// Step name
	Step *string `json:"Step,omitempty" name:"Step"`

	// Lifecycle
	// pending: the step is not started
	// running: the step is in progress
	// success: the step is completed
	// failed: the step failed
	LifeState *string `json:"LifeState,omitempty" name:"LifeState"`

	// Step start time
	// Note: this field may return `null`, indicating that no valid value is obtained.
	StartAt *string `json:"StartAt,omitempty" name:"StartAt"`

	// Step end time
	// Note: this field may return `null`, indicating that no valid value is obtained.
	EndAt *string `json:"EndAt,omitempty" name:"EndAt"`

	// If the lifecycle of the step is failed, this field will display the error information.
	// Note: this field may return `null`, indicating that no valid value is obtained.
	FailedMsg *string `json:"FailedMsg,omitempty" name:"FailedMsg"`
}

// Predefined struct for user
type UpdateClusterVersionRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The version that needs to upgrade to
	DstVersion *string `json:"DstVersion,omitempty" name:"DstVersion"`

	// Cluster custom parameter
	ExtraArgs *ClusterExtraArgs `json:"ExtraArgs,omitempty" name:"ExtraArgs"`

	// The maximum tolerable number of unavailable pods
	MaxNotReadyPercent *float64 `json:"MaxNotReadyPercent,omitempty" name:"MaxNotReadyPercent"`

	// Whether to skip the precheck
	SkipPreCheck *bool `json:"SkipPreCheck,omitempty" name:"SkipPreCheck"`
}

type UpdateClusterVersionRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The version that needs to upgrade to
	DstVersion *string `json:"DstVersion,omitempty" name:"DstVersion"`

	// Cluster custom parameter
	ExtraArgs *ClusterExtraArgs `json:"ExtraArgs,omitempty" name:"ExtraArgs"`

	// The maximum tolerable number of unavailable pods
	MaxNotReadyPercent *float64 `json:"MaxNotReadyPercent,omitempty" name:"MaxNotReadyPercent"`

	// Whether to skip the precheck
	SkipPreCheck *bool `json:"SkipPreCheck,omitempty" name:"SkipPreCheck"`
}

func (r *UpdateClusterVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateClusterVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "DstVersion")
	delete(f, "ExtraArgs")
	delete(f, "MaxNotReadyPercent")
	delete(f, "SkipPreCheck")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateClusterVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateClusterVersionResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateClusterVersionResponse struct {
	*tchttp.BaseResponse
	Response *UpdateClusterVersionResponseParams `json:"Response"`
}

func (r *UpdateClusterVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateClusterVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeAbleInstancesItem struct {
	// Node ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// The current version of the node
	Version *string `json:"Version,omitempty" name:"Version"`

	// The latest minor version of the current version
	// Note: this field may return `null`, indicating that no valid value is obtained.
	LatestVersion *string `json:"LatestVersion,omitempty" name:"LatestVersion"`
}

// Predefined struct for user
type UpgradeClusterInstancesRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// create: starting an upgrade task
	// pause: pausing the task
	// resume: continuing the task
	// abort: stopping the task
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// Upgrade type. It’s only required when `Operation` is set as `create`.
	// reset: the reinstallation and upgrade of major version
	// hot: the hot upgrade of minor version
	// major: in-place upgrade of major version
	UpgradeType *string `json:"UpgradeType,omitempty" name:"UpgradeType"`

	// List of nodes that need to upgrade
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// This parameter is used when the node joins the cluster again. Refer to the API of creating one or more cluster nodes.
	ResetParam *UpgradeNodeResetParam `json:"ResetParam,omitempty" name:"ResetParam"`

	// Whether to skip the pre-upgrade check of the node
	SkipPreCheck *bool `json:"SkipPreCheck,omitempty" name:"SkipPreCheck"`

	// The maximum tolerable proportion of unavailable pods
	MaxNotReadyPercent *float64 `json:"MaxNotReadyPercent,omitempty" name:"MaxNotReadyPercent"`
}

type UpgradeClusterInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// create: starting an upgrade task
	// pause: pausing the task
	// resume: continuing the task
	// abort: stopping the task
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// Upgrade type. It’s only required when `Operation` is set as `create`.
	// reset: the reinstallation and upgrade of major version
	// hot: the hot upgrade of minor version
	// major: in-place upgrade of major version
	UpgradeType *string `json:"UpgradeType,omitempty" name:"UpgradeType"`

	// List of nodes that need to upgrade
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// This parameter is used when the node joins the cluster again. Refer to the API of creating one or more cluster nodes.
	ResetParam *UpgradeNodeResetParam `json:"ResetParam,omitempty" name:"ResetParam"`

	// Whether to skip the pre-upgrade check of the node
	SkipPreCheck *bool `json:"SkipPreCheck,omitempty" name:"SkipPreCheck"`

	// The maximum tolerable proportion of unavailable pods
	MaxNotReadyPercent *float64 `json:"MaxNotReadyPercent,omitempty" name:"MaxNotReadyPercent"`
}

func (r *UpgradeClusterInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeClusterInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Operation")
	delete(f, "UpgradeType")
	delete(f, "InstanceIds")
	delete(f, "ResetParam")
	delete(f, "SkipPreCheck")
	delete(f, "MaxNotReadyPercent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeClusterInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeClusterInstancesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpgradeClusterInstancesResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeClusterInstancesResponseParams `json:"Response"`
}

func (r *UpgradeClusterInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeClusterInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeNodeResetParam struct {
	// Additional parameters set for the instance
	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitempty" name:"InstanceAdvancedSettings"`

	// Enhanced services. You can use this parameter to specify whether to enable services such as Cloud Security and Cloud Monitor. If this parameter is not specified, Cloud Monitor and Cloud Security will be enabled by default.
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`

	// Node login information. For now, it only supports Password or a single KeyIds
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// Security group to which the instance belongs. This parameter can be obtained from the `sgId` field in the response of `DescribeSecurityGroups`. If this parameter is not specified, the default security group is bound. (Currently, you can only set a single sgId.)
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
}

type VersionInstance struct {
	// Version name
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Version Info
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Version *string `json:"Version,omitempty" name:"Version"`

	// Remark
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}