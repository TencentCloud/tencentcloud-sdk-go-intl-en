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

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type AddExistedInstancesRequest struct {
	*tchttp.BaseRequest

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Instance list
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// Additional parameter to be set for the instance
	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitempty" name:"InstanceAdvancedSettings"`

	// Enhanced services. This parameter is used to specify whether to enable Cloud Security, Cloud Monitoring and other services. If this parameter is not specified, Cloud Monitor and Cloud Security are enabled by default.
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`

	// Node login information (currently only supports using Password or single KeyIds)
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// Security group to which the instance belongs. This parameter can be obtained from the `sgId` field returned by DescribeSecurityGroups. If this parameter is not specified, the default security group is bound. (Currently, you can only set a single sgId)
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`

	// 
	HostName *string `json:"HostName,omitempty" name:"HostName"`
}

func (r *AddExistedInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddExistedInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddExistedInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// IDs of failed nodes
	// Note: This field may return null, indicating that no valid value was found.
		FailedInstanceIds []*string `json:"FailedInstanceIds,omitempty" name:"FailedInstanceIds" list`

		// IDs of successful nodes
	// Note: This field may return null, indicating that no valid value was found.
		SuccInstanceIds []*string `json:"SuccInstanceIds,omitempty" name:"SuccInstanceIds" list`

		// IDs of (successful or failed) nodes that timed out
	// Note: This field may return null, indicating that no valid value was found.
		TimeoutInstanceIds []*string `json:"TimeoutInstanceIds,omitempty" name:"TimeoutInstanceIds" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddExistedInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddExistedInstancesResponse) FromJsonString(s string) error {
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

	// 
	TagSpecification []*TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification" list`

	// 
	ClusterStatus *string `json:"ClusterStatus,omitempty" name:"ClusterStatus"`

	// 
	Property *string `json:"Property,omitempty" name:"Property"`

	// Number of master nodes currently in the cluster
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
}

type ClusterAdvancedSettings struct {

	// Whether IPVS is enabled
	IPVS *bool `json:"IPVS,omitempty" name:"IPVS"`

	// Whether auto-scaling is enabled for nodes in the cluster (Enabling this function is not supported when you create a cluster)
	AsEnabled *bool `json:"AsEnabled,omitempty" name:"AsEnabled"`

	// Type of runtime component used by the cluster. The types include "docker" and "containerd". Default value: docker
	ContainerRuntime *string `json:"ContainerRuntime,omitempty" name:"ContainerRuntime"`

	// 
	NodeNameType *string `json:"NodeNameType,omitempty" name:"NodeNameType"`

	// Cluster custom parameter
	ExtraArgs *ClusterExtraArgs `json:"ExtraArgs,omitempty" name:"ExtraArgs"`

	// Cluster network type, which can be GR (Global Router) or VPC-CNI. The default value is GR.
	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`

	// Whether a cluster in VPC-CNI mode uses dynamic IP addresses. The default value is FALSE, which indicates that static IP addresses are used.
	IsNonStaticIpMode *bool `json:"IsNonStaticIpMode,omitempty" name:"IsNonStaticIpMode"`
}

type ClusterBasicSettings struct {

	// Cluster operating system. CentOS 7.2x86_64 or Ubuntu 16.04.1 LTSx86_64. Default value: Ubuntu 16.04.1 LTSx86_64
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

	// 
	TagSpecification []*TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification" list`

	// 
	OsCustomizeType *string `json:"OsCustomizeType,omitempty" name:"OsCustomizeType"`

	// 
	NeedWorkSecurityGroup *bool `json:"NeedWorkSecurityGroup,omitempty" name:"NeedWorkSecurityGroup"`
}

type ClusterCIDRSettings struct {

	// CIDR used to assign container and service IPs for the cluster. It cannot conflict with the VPC’s CIDR or the CIDRs of other clusters in the same VPC
	ClusterCIDR *string `json:"ClusterCIDR,omitempty" name:"ClusterCIDR"`

	// Whether to ignore ClusterCIDR conflict errors, which are not ignored by default
	IgnoreClusterCIDRConflict *bool `json:"IgnoreClusterCIDRConflict,omitempty" name:"IgnoreClusterCIDRConflict"`

	// Maximum number of pods on each node in the cluster
	MaxNodePodNum *uint64 `json:"MaxNodePodNum,omitempty" name:"MaxNodePodNum"`

	// Maximum number of cluster services
	MaxClusterServiceNum *uint64 `json:"MaxClusterServiceNum,omitempty" name:"MaxClusterServiceNum"`

	// The CIDR block used to assign cluster service IP addresses. It must conflict with neither the VPC CIDR block nor with CIDR blocks of other clusters in the same VPC instance. The IP range must be within the private network IP range, such as 10.1.0.0/14, 192.168.0.1/18, and 172.16.0.0/16.
	ServiceCIDR *string `json:"ServiceCIDR,omitempty" name:"ServiceCIDR"`

	// Subnet ID of the ENI in VPC-CNI network mode
	EniSubnetIds []*string `json:"EniSubnetIds,omitempty" name:"EniSubnetIds" list`

	// Repossession time of ENI IP addresses in VPC-CNI network mode, whose range is [300,15768000)
	ClaimExpiredSeconds *int64 `json:"ClaimExpiredSeconds,omitempty" name:"ClaimExpiredSeconds"`
}

type ClusterExtraArgs struct {

	// kube-apiserver custom parameter
	// Note: this field may return null, indicating that no valid value is obtained.
	KubeAPIServer []*string `json:"KubeAPIServer,omitempty" name:"KubeAPIServer" list`

	// kube-controller-manager custom parameter
	// Note: this field may return null, indicating that no valid value is obtained.
	KubeControllerManager []*string `json:"KubeControllerManager,omitempty" name:"KubeControllerManager" list`

	// kube-scheduler custom parameter
	// Note: this field may return null, indicating that no valid value is obtained.
	KubeScheduler []*string `json:"KubeScheduler,omitempty" name:"KubeScheduler" list`
}

type ClusterNetworkSettings struct {

	// CIDR used to assign container and service IPs for the cluster. It cannot conflict with the VPC’s CIDR or the CIDRs of other clusters in the same VPC.
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
}

type CreateClusterAsGroupRequest struct {
	*tchttp.BaseRequest

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The pass-through parameters for scaling group creation, in the format of a JSON string. For more information, see the [CreateAutoScalingGroup](https://cloud.tencent.com/document/api/377/20440) API. The **LaunchConfigurationId** is created with the LaunchConfigurePara parameter, which does not support data entry.
	AutoScalingGroupPara *string `json:"AutoScalingGroupPara,omitempty" name:"AutoScalingGroupPara"`

	// The pass-through parameters for launch configuration creation, in the format of a JSON string. For more information, see the [CreateLaunchConfiguration](https://cloud.tencent.com/document/api/377/20447) API. **ImageId** is not required as it is already included in the cluster dimension. **UserData** is not required as it’s set through the **UserScript**.
	LaunchConfigurePara *string `json:"LaunchConfigurePara,omitempty" name:"LaunchConfigurePara"`

	// Advanced configuration information of the node
	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitempty" name:"InstanceAdvancedSettings"`

	// Node label array
	Labels []*Label `json:"Labels,omitempty" name:"Labels" list`
}

func (r *CreateClusterAsGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClusterAsGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateClusterAsGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Launch configuration ID
		LaunchConfigurationId *string `json:"LaunchConfigurationId,omitempty" name:"LaunchConfigurationId"`

		// Scaling group ID
		AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClusterAsGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClusterAsGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateClusterEndpointRequest struct {
	*tchttp.BaseRequest

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The ID of the subnet where the cluster’s port is located (only needs to be entered when the non-public network access is enabled, and must be within the subnet of the cluster’s VPC). 
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Whether public network access is enabled or not (True = public network access, FALSE = private network access, with the default value as FALSE).
	IsExtranet *bool `json:"IsExtranet,omitempty" name:"IsExtranet"`
}

func (r *CreateClusterEndpointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClusterEndpointRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateClusterEndpointResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClusterEndpointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClusterEndpointResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateClusterEndpointVipRequest struct {
	*tchttp.BaseRequest

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Security policy opens single IP or CIDR to the Internet (for example: “192.168.1.0/24”, with “reject all” as the default).
	SecurityPolicies []*string `json:"SecurityPolicies,omitempty" name:"SecurityPolicies" list`
}

func (r *CreateClusterEndpointVipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClusterEndpointVipRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateClusterEndpointVipResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Request job’s FlowId
		RequestFlowId *int64 `json:"RequestFlowId,omitempty" name:"RequestFlowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClusterEndpointVipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClusterEndpointVipResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateClusterInstancesRequest struct {
	*tchttp.BaseRequest

	// Cluster ID. Enter the ClusterId field returned by the DescribeClusters API
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Pass-through parameter for CVM creation in the format of a JSON string. For more information, see the [RunInstances](https://cloud.tencent.com/document/product/213/15730) API.
	RunInstancePara *string `json:"RunInstancePara,omitempty" name:"RunInstancePara"`

	// Additional parameter to be set for the instance
	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitempty" name:"InstanceAdvancedSettings"`
}

func (r *CreateClusterInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClusterInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateClusterInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Instance ID
		InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClusterInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClusterInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateClusterRequest struct {
	*tchttp.BaseRequest

	// Container networking configuration information for the cluster
	ClusterCIDRSettings *ClusterCIDRSettings `json:"ClusterCIDRSettings,omitempty" name:"ClusterCIDRSettings"`

	// Cluster type. Managed cluster: MANAGED_CLUSTER; self-deployed cluster: INDEPENDENT_CLUSTER.
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// Pass-through parameter for CVM creation in the format of a JSON string. For more information, see the API for [creating a CVM instance](https://cloud.tencent.com/document/product/213/15730).
	RunInstancesForNode []*RunInstancesForNode `json:"RunInstancesForNode,omitempty" name:"RunInstancesForNode" list`

	// Basic configuration information of the cluster
	ClusterBasicSettings *ClusterBasicSettings `json:"ClusterBasicSettings,omitempty" name:"ClusterBasicSettings"`

	// Advanced configuration information of the cluster
	ClusterAdvancedSettings *ClusterAdvancedSettings `json:"ClusterAdvancedSettings,omitempty" name:"ClusterAdvancedSettings"`

	// Advanced configuration information of the node
	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitempty" name:"InstanceAdvancedSettings"`

	// Configuration information of an existing instance
	ExistedInstancesForNode []*ExistedInstancesForNode `json:"ExistedInstancesForNode,omitempty" name:"ExistedInstancesForNode" list`

	// CVM type and the corresponding data disk mounting configuration information.
	InstanceDataDiskMountSettings []*InstanceDataDiskMountSetting `json:"InstanceDataDiskMountSettings,omitempty" name:"InstanceDataDiskMountSettings" list`
}

func (r *CreateClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClusterRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Cluster ID
		ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClusterResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *CreateClusterRouteTableRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateClusterRouteTableResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClusterRouteTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClusterRouteTableResponse) FromJsonString(s string) error {
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
}

type DeleteClusterAsGroupsRequest struct {
	*tchttp.BaseRequest

	// The cluster ID, obtained through the [DescribeClusters](https://cloud.tencent.com/document/api/457/31862) API.
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Cluster scaling group ID list
	AutoScalingGroupIds []*string `json:"AutoScalingGroupIds,omitempty" name:"AutoScalingGroupIds" list`

	// Whether to keep nodes in the scaling group. Default to **false** (not keep)
	KeepInstance *bool `json:"KeepInstance,omitempty" name:"KeepInstance"`
}

func (r *DeleteClusterAsGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClusterAsGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterAsGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClusterAsGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClusterAsGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DeleteClusterEndpointRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterEndpointResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClusterEndpointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClusterEndpointResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DeleteClusterEndpointVipRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterEndpointVipResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClusterEndpointVipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClusterEndpointVipResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterInstancesRequest struct {
	*tchttp.BaseRequest

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// List of Instance IDs
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// Policy used to delete an instance in the cluster: `terminate` (terminates the instance. Only available for pay-as-you-go CVMs); `retain` (only removes it from the cluster. The instance will be retained.)
	InstanceDeleteMode *string `json:"InstanceDeleteMode,omitempty" name:"InstanceDeleteMode"`

	// 
	ForceDelete *bool `json:"ForceDelete,omitempty" name:"ForceDelete"`
}

func (r *DeleteClusterInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClusterInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// IDs of deleted instances
		SuccInstanceIds []*string `json:"SuccInstanceIds,omitempty" name:"SuccInstanceIds" list`

		// IDs of instances failed to be deleted
		FailedInstanceIds []*string `json:"FailedInstanceIds,omitempty" name:"FailedInstanceIds" list`

		// IDs of instances that cannot be found
		NotFoundInstanceIds []*string `json:"NotFoundInstanceIds,omitempty" name:"NotFoundInstanceIds" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClusterInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClusterInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterRequest struct {
	*tchttp.BaseRequest

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Policy used to delete an instance in the cluster: terminate (terminates the instance. Only available for instances on pay-as-you-go CVMs); retain (only removes it from the cluster. The instance will be retained.)
	InstanceDeleteMode *string `json:"InstanceDeleteMode,omitempty" name:"InstanceDeleteMode"`

	// Specifies the policy to deal with resources in the cluster when the cluster is deleted. It only supports CBS now. The default policy is to retain CBS disks.
	ResourceDeleteOptions []*ResourceDeleteOption `json:"ResourceDeleteOptions,omitempty" name:"ResourceDeleteOptions" list`
}

func (r *DeleteClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClusterRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClusterResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DeleteClusterRouteRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterRouteResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClusterRouteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClusterRouteResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DeleteClusterRouteTableRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterRouteTableResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClusterRouteTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClusterRouteTableResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeClusterEndpointStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterEndpointStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Queries cluster access port status (Created = successfully enabled; Creating = in the process of being enabled; NotFound = not enabled).
		Status *string `json:"Status,omitempty" name:"Status"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterEndpointStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterEndpointStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeClusterEndpointVipStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterEndpointVipStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Port operation status (Creating = in the process of creation; CreateFailed = creation has failed; Created = creation completed; Deleting = in the process of deletion; DeletedFailed = deletion has failed; Deleted = deletion completed; NotFound = operation not found)
		Status *string `json:"Status,omitempty" name:"Status"`

		// Reason for operation failure
		ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterEndpointVipStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterEndpointVipStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 
	InstanceRole *string `json:"InstanceRole,omitempty" name:"InstanceRole"`
}

func (r *DescribeClusterInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of instances in the cluster
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of instances in the cluster
		InstanceSet []*Instance `json:"InstanceSet,omitempty" name:"InstanceSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterRouteTablesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeClusterRouteTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterRouteTablesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterRouteTablesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of instances that match the filter condition(s).
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Object of cluster route table.
		RouteTableSet []*RouteTableInfo `json:"RouteTableSet,omitempty" name:"RouteTableSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterRouteTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterRouteTablesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterRoutesRequest struct {
	*tchttp.BaseRequest

	// Route table name.
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`
}

func (r *DescribeClusterRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterRoutesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterRoutesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of instances that match the filter condition(s).
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Object of cluster route.
		RouteSet []*RouteInfo `json:"RouteSet,omitempty" name:"RouteSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterRoutesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeClusterSecurityRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterSecurityResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Cluster’s account name
		UserName *string `json:"UserName,omitempty" name:"UserName"`

		// Cluster’s password
		Password *string `json:"Password,omitempty" name:"Password"`

		// Cluster’s access CA certificate
		CertificationAuthority *string `json:"CertificationAuthority,omitempty" name:"CertificationAuthority"`

		// Cluster’s access address
		ClusterExternalEndpoint *string `json:"ClusterExternalEndpoint,omitempty" name:"ClusterExternalEndpoint"`

		// Domain name accessed by the cluster
		Domain *string `json:"Domain,omitempty" name:"Domain"`

		// Cluster’s endpoint address
		PgwEndpoint *string `json:"PgwEndpoint,omitempty" name:"PgwEndpoint"`

		// Cluster’s access policy group
	// Note: This field may return null, indicating that no valid value was found.
		SecurityPolicy []*string `json:"SecurityPolicy,omitempty" name:"SecurityPolicy" list`

		// Cluster Kubeconfig file
	// Note: This field may return null, indicating that no valid value was found.
		Kubeconfig *string `json:"Kubeconfig,omitempty" name:"Kubeconfig"`

		// Access address of the cluster JnsGw
	// Note: This field may return null, indicating that no valid value was found.
		JnsGwEndpoint *string `json:"JnsGwEndpoint,omitempty" name:"JnsGwEndpoint"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterSecurityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterSecurityResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClustersRequest struct {
	*tchttp.BaseRequest

	// Cluster ID list (When it is empty,
	// all clusters under the account will be obtained)
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds" list`

	// Offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of output entries. Default value: 20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Filter condition. Currently, only filtering by a single ClusterName is supported
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClustersRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClustersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of clusters
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Cluster information list
		Clusters []*Cluster `json:"Clusters,omitempty" name:"Clusters" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClustersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeExistedInstancesRequest struct {
	*tchttp.BaseRequest

	// Cluster ID. Enter the `ClusterId` field returned when you call the DescribeClusters API (Only VPC ID obtained through `ClusterId` need filtering conditions. When comparing statuses, the nodes on all clusters in this region will be used for comparison. You cannot specify `InstanceIds` and `ClusterId` at the same time.)
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Query by one or more instance ID(s). Instance ID format: ins-xxxxxxxx. (Refer to section ID.N of the API overview for this parameter’s specific format.) Up to 100 instances are allowed for each request. You cannot specify InstanceIds and Filters at the same time.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// Filter condition. For fields and other information, see [the DescribeInstances API](https://cloud.tencent.com/document/api/213/15728). If a ClusterId has been set, then the cluster’s VPC ID will be attached as a query field. In this situation, if a "vpc-id" is specified in Filter, then the specified VPC ID must be consistent with the cluster’s VPC ID.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Filter by instance IP (Supports both private and public IPs)
	VagueIpAddress *string `json:"VagueIpAddress,omitempty" name:"VagueIpAddress"`

	// Filter by instance name
	VagueInstanceName *string `json:"VagueInstanceName,omitempty" name:"VagueInstanceName"`

	// Offset. Default value: 0. For more information on Offset, see the relevant section in the API [Introduction](https://cloud.tencent.com/document/api/213/15688).
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on Limit, see the relevant section in the API [Introduction](https://cloud.tencent.com/document/api/213/15688).
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeExistedInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeExistedInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeExistedInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Array of existing instance information.
	// Note: This field may return null, indicating that no valid values can be obtained.
		ExistedInstanceSet []*ExistedInstance `json:"ExistedInstanceSet,omitempty" name:"ExistedInstanceSet" list`

		// Number of instances that match the filter condition(s).
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExistedInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeExistedInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImagesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeImagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImagesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImagesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of images
	// Note: this field may return null, indicating that no valid values can be obtained.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Image information list
	// Note: this field may return null, indicating that no valid values can be obtained.
		ImageInstanceSet []*ImageInstance `json:"ImageInstanceSet,omitempty" name:"ImageInstanceSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImagesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRegionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of regions
	// Note: this field may return null, indicating that no valid values can be obtained.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// ## Region List
	// Note: this field may return null, indicating that no valid values can be obtained.
		RegionInstanceSet []*RegionInstance `json:"RegionInstanceSet,omitempty" name:"RegionInstanceSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRegionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeRouteTableConflictsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRouteTableConflictsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Whether there is a conflict in the route table.
		HasConflict *bool `json:"HasConflict,omitempty" name:"HasConflict"`

		// Route table conflict list.
	// Note: This field may return null, indicating that no valid values can be obtained.
		RouteTableConflictSet []*RouteTableConflict `json:"RouteTableConflictSet,omitempty" name:"RouteTableConflictSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRouteTableConflictsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRouteTableConflictsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnhancedService struct {

	// Enables cloud security service. If this parameter is not specified, the cloud security service will be enabled by default.
	SecurityService *RunSecurityServiceEnabled `json:"SecurityService,omitempty" name:"SecurityService"`

	// Enables cloud monitor service. If this parameter is not specified, the cloud monitor service will be enabled by default.
	MonitorService *RunMonitorServiceEnabled `json:"MonitorService,omitempty" name:"MonitorService"`
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

	// List of private IPs of the instance’s primary ENI.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses" list`

	// List of public IPs of the instance’s primary ENI.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitempty" name:"PublicIpAddresses" list`

	// Creation time, which follows the ISO8601 standard and uses UTC time. Format: YYYY-MM-DDThh:mm:ssZ.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Instance’s number of CPU cores. Unit: cores.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CPU *uint64 `json:"CPU,omitempty" name:"CPU"`

	// Instance’s memory capacity. Unit: GB.
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
}

type ExistedInstancesForNode struct {

	// Node role. Values: MASTER_ETCD, WORKER. You only need to specify MASTER_ETCD when creating a self-deployed cluster (INDEPENDENT_CLUSTER).
	NodeRole *string `json:"NodeRole,omitempty" name:"NodeRole"`

	// Reinstallation parameter of existing instances
	ExistedInstancesPara *ExistedInstancesPara `json:"ExistedInstancesPara,omitempty" name:"ExistedInstancesPara"`

	// Advanced node setting, which overrides the InstanceAdvancedSettings item set at the cluster level (currently valid for the ExtraArgs node custom parameter only)
	InstanceAdvancedSettingsOverride *InstanceAdvancedSettings `json:"InstanceAdvancedSettingsOverride,omitempty" name:"InstanceAdvancedSettingsOverride"`
}

type ExistedInstancesPara struct {

	// Cluster ID
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// Additional parameter to be set for the instance
	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitempty" name:"InstanceAdvancedSettings"`

	// Enhanced services. This parameter is used to specify whether to enable Cloud Security, Cloud Monitor and other services. If this parameter is not specified, Cloud Monitor and Cloud Security are enabled by default.
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`

	// Node login information (currently only supports using Password or single KeyIds)
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// Security group to which the instance belongs. This parameter can be obtained from the sgId field in the returned values of DescribeSecurityGroups. If this parameter is not specified, the default security group is bound. (Currently, you can only set a single sgId)
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`

	// 
	HostName *string `json:"HostName,omitempty" name:"HostName"`
}

type Filter struct {

	// Filters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Filter values.
	Values []*string `json:"Values,omitempty" name:"Values" list`
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

	// Data disk mount point. By default, no data disk is mounted. Data disks in ext3, ext4, or XFS file system formats will be mounted directly, while data disks in other file systems and unformatted data disks will automatically be formatted as ext4 and then mounted. Please back up your data in advance. This setting is only applicable to CVMs with a single data disk.
	// Note: This field may return null, indicating that no valid value was found.
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
	Labels []*Label `json:"Labels,omitempty" name:"Labels" list`

	// Data disk information
	// Note: This field may return null, indicating that no valid value was found.
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks" list`

	// Information about node custom parameters
	// Note: This field may return null, indicating that no valid value was found.
	ExtraArgs *InstanceExtraArgs `json:"ExtraArgs,omitempty" name:"ExtraArgs"`
}

type InstanceDataDiskMountSetting struct {

	// CVM instance type
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Data disk mounting information
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks" list`

	// Availability zone where the CVM instance is located
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type InstanceExtraArgs struct {

	// Kubelet custom parameter
	// Note: this field may return null, indicating that no valid value is obtained.
	Kubelet []*string `json:"Kubelet,omitempty" name:"Kubelet" list`
}

type Label struct {

	// 
	Name *string `json:"Name,omitempty" name:"Name"`

	// 
	Value *string `json:"Value,omitempty" name:"Value"`
}

type LoginSettings struct {

	// Login password of the instance. The password requirements vary among different operating systems: <br><li>For Linux instances, the password must be 8-30 characters long and contain at least two of the following types: [a-z], [A-Z], [0-9] and [( ) \` ~ ! @ # $ % ^ & *  - + = | { } [ ] : ; ' , . ? / ]. <br><li>For Windows instances, the password must be 12-30 characters long and contain at least three of the following categories: [a-z], [A-Z], [0-9] and [( ) \` ~ ! @ # $ % ^ & * - + = | { } [ ] : ; ' , . ? /]. <br><br>If this parameter is not specified, a random password will be generated and sent to you via the Message Center.
	// Note: this field may return null, indicating that no valid value is obtained.
	Password *string `json:"Password,omitempty" name:"Password"`

	// List of key IDs. After an instance is associated with a key, you can access the instance with the private key in the key pair. You can call [`DescribeKeyPairs`](https://cloud.tencent.com/document/api/213/15699) to obtain `KeyId`. A key and password cannot be specified at the same time. Windows instances do not support keys. Currently, you can only specify one key when purchasing an instance.
	// Note: this field may return null, indicating that no valid value is obtained.
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds" list`

	// Whether to keep the original settings of an image. You cannot specify this parameter and `Password` or `KeyIds.N` at the same time. You can specify this parameter as `TRUE` only when you create an instance using a custom image, a shared image, or an imported image. Valid values: <br><li>TRUE: keep the login settings of the image <br><li>FALSE: do not keep the login settings of the image <br><br>Default value: FALSE.
	// Note: This field may return null, indicating that no valid value is found.
	KeepImageLogin *string `json:"KeepImageLogin,omitempty" name:"KeepImageLogin"`
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
}

func (r *ModifyClusterAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClusterAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Project of the Cluster
	// Note: this field may return null, indicating that no valid values can be obtained.
		ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

		// Cluster name
	// Note: this field may return null, indicating that no valid values can be obtained.
		ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

		// Cluster description
	// Note: this field may return null, indicating that no valid values can be obtained.
		ClusterDesc *string `json:"ClusterDesc,omitempty" name:"ClusterDesc"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyClusterAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClusterAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterEndpointSPRequest struct {
	*tchttp.BaseRequest

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Security policy opens single IP or CIDR block to the Internet (for example: “192.168.1.0/24”, with “reject all” as the default).
	SecurityPolicies []*string `json:"SecurityPolicies,omitempty" name:"SecurityPolicies" list`
}

func (r *ModifyClusterEndpointSPRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClusterEndpointSPRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterEndpointSPResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyClusterEndpointSPResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClusterEndpointSPResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

type ResourceDeleteOption struct {

	// Resource type, for example `CBS`
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// Specifies the policy to deal with resources in the cluster when the cluster is deleted. It can be `terminate` or `retain`.
	DeleteMode *string `json:"DeleteMode,omitempty" name:"DeleteMode"`
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

type RunInstancesForNode struct {

	// Node role. Values: MASTER_ETCD, WORKER. You only need to specify MASTER_ETCD when creating a self-deployed cluster (INDEPENDENT_CLUSTER).
	NodeRole *string `json:"NodeRole,omitempty" name:"NodeRole"`

	// Pass-through parameter for CVM creation in the format of a JSON string. For more information, see the API for [creating a CVM instance](https://cloud.tencent.com/document/product/213/15730). Pass any parameter other than common parameters. ImageId will be replaced with the image corresponding to the TKE cluster operating system.
	RunInstancesPara []*string `json:"RunInstancesPara,omitempty" name:"RunInstancesPara" list`

	// An advanced node setting. This parameter overrides the InstanceAdvancedSettings item set at the cluster level and corresponds to RunInstancesPara in a one-to-one sequential manner (currently valid for the ExtraArgs node custom parameter only).
	InstanceAdvancedSettingsOverrides []*InstanceAdvancedSettings `json:"InstanceAdvancedSettingsOverrides,omitempty" name:"InstanceAdvancedSettingsOverrides" list`
}

type RunMonitorServiceEnabled struct {

	// Whether to enable [Cloud Monitor](/document/product/248). Valid values: <br><li>TRUE: enable Cloud Monitor <br><li>FALSE: do not enable Cloud Monitor <br><br>Default value: TRUE.
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

type RunSecurityServiceEnabled struct {

	// Whether to enable [Cloud Security](/document/product/296). Valid values: <br><li>TRUE: enable Cloud Security <br><li>FALSE: do not enable Cloud Security <br><br>Default value: TRUE.
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

type Tag struct {

	// Tag key.
	Key *string `json:"Key,omitempty" name:"Key"`

	// Tag value.
	Value *string `json:"Value,omitempty" name:"Value"`
}

type TagSpecification struct {

	// 
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`
}
