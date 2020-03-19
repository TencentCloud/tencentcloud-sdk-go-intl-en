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

package v20190103

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type COSSettings struct {

	// COS `SecretId`
	CosSecretId *string `json:"CosSecretId,omitempty" name:"CosSecretId"`

	// COS `SecrectKey`
	CosSecretKey *string `json:"CosSecretKey,omitempty" name:"CosSecretKey"`

	// COS path to log
	LogOnCosPath *string `json:"LogOnCosPath,omitempty" name:"LogOnCosPath"`
}

type CdbInfo struct {

	// Database instance
	// Note: this field may return null, indicating that no valid values can be obtained.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Database IP
	// Note: this field may return null, indicating that no valid values can be obtained.
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Database port
	// Note: this field may return null, indicating that no valid values can be obtained.
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Database memory specification
	// Note: this field may return null, indicating that no valid values can be obtained.
	MemSize *int64 `json:"MemSize,omitempty" name:"MemSize"`

	// Database disk specification
	// Note: this field may return null, indicating that no valid values can be obtained.
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`

	// Service flag
	// Note: this field may return null, indicating that no valid values can be obtained.
	Service *string `json:"Service,omitempty" name:"Service"`

	// Expiration time
	// Note: this field may return null, indicating that no valid values can be obtained.
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// Application time
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApplyTime *string `json:"ApplyTime,omitempty" name:"ApplyTime"`

	// Payment type
	// Note: this field may return null, indicating that no valid values can be obtained.
	PayType *int64 `json:"PayType,omitempty" name:"PayType"`

	// Expiration flag
	// Note: this field may return null, indicating that no valid values can be obtained.
	ExpireFlag *bool `json:"ExpireFlag,omitempty" name:"ExpireFlag"`

	// Database status
	// Note: this field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Renewal flag
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsAutoRenew *int64 `json:"IsAutoRenew,omitempty" name:"IsAutoRenew"`

	// Database string
	// Note: this field may return null, indicating that no valid values can be obtained.
	SerialNo *string `json:"SerialNo,omitempty" name:"SerialNo"`

	// ZoneId
	// Note: this field may return null, indicating that no valid values can be obtained.
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// RegionId
	// Note: this field may return null, indicating that no valid values can be obtained.
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
}

type ClusterInstancesInfo struct {

	// ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// Cluster ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Title
	// Note: this field may return null, indicating that no valid values can be obtained.
	Ftitle *string `json:"Ftitle,omitempty" name:"Ftitle"`

	// Cluster name
	// Note: this field may return null, indicating that no valid values can be obtained.
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// Region ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// Region ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// User APPID
	// Note: this field may return null, indicating that no valid values can be obtained.
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// User UIN
	// Note: this field may return null, indicating that no valid values can be obtained.
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// Project ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Cluster `VPCID`
	// Note: this field may return null, indicating that no valid values can be obtained.
	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	SubnetId *int64 `json:"SubnetId,omitempty" name:"SubnetId"`

	// Instance status code. Value range:
	// <li>2: cluster running</li>
	// <li>3: creating cluster.</li>
	// <li>4: scaling out cluster.</li>
	// <li>5: adding router node in cluster.</li>
	// <li>6: installing component in cluster.</li>
	// <li>7: cluster executing command.</li>
	// <li>8: restarting service.</li>
	// <li>9: entering maintenance.</li>
	// <li>10: suspending service.</li>
	// <li>11: exiting maintenance.</li>
	// <li>12: exiting suspension.</li>
	// <li>13: delivering configuration.</li>
	// <li>14: terminating cluster.</li>
	// <li>15: terminating core node.</li>
	// <li>16: terminating task node.</li>
	// <li>17: terminating router node.</li>
	// <li>18: changing webproxy password.</li>
	// <li>19: isolating cluster.</li>
	// <li>20: resuming cluster.</li>
	// <li>21: repossessing cluster.</li>
	// <li>22: waiting for configuration adjustment.</li>
	// <li>23: cluster isolated.</li>
	// <li>24: removing node.</li>
	// <li>33: waiting for refund.</li>
	// <li>34: refunded.</li>
	// <li>301: creation failed.</li>
	// <li>302: scale-out failed.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Creation time
	// Note: this field may return null, indicating that no valid values can be obtained.
	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

	// Execution duration
	// Note: this field may return null, indicating that no valid values can be obtained.
	RunTime *string `json:"RunTime,omitempty" name:"RunTime"`

	// Cluster product configuration information
	// Note: this field may return null, indicating that no valid values can be obtained.
	Config *EmrProductConfigOutter `json:"Config,omitempty" name:"Config"`

	// Public IP of master node
	// Note: this field may return null, indicating that no valid values can be obtained.
	MasterIp *string `json:"MasterIp,omitempty" name:"MasterIp"`

	// EMR version
	// Note: this field may return null, indicating that no valid values can be obtained.
	EmrVersion *string `json:"EmrVersion,omitempty" name:"EmrVersion"`

	// Billing mode
	// Note: this field may return null, indicating that no valid values can be obtained.
	ChargeType *int64 `json:"ChargeType,omitempty" name:"ChargeType"`

	// Transaction version
	// Note: this field may return null, indicating that no valid values can be obtained.
	TradeVersion *int64 `json:"TradeVersion,omitempty" name:"TradeVersion"`

	// Resource order ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ResourceOrderId *int64 `json:"ResourceOrderId,omitempty" name:"ResourceOrderId"`

	// Whether this is a paid cluster
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsTradeCluster *int64 `json:"IsTradeCluster,omitempty" name:"IsTradeCluster"`

	// Alarm information for cluster error
	// Note: this field may return null, indicating that no valid values can be obtained.
	AlarmInfo *string `json:"AlarmInfo,omitempty" name:"AlarmInfo"`

	// Whether the new architecture is used
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsWoodpeckerCluster *int64 `json:"IsWoodpeckerCluster,omitempty" name:"IsWoodpeckerCluster"`

	// Metadatabase information
	// Note: this field may return null, indicating that no valid values can be obtained.
	MetaDb *string `json:"MetaDb,omitempty" name:"MetaDb"`

	// Tag information
	// Note: this field may return null, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`

	// Hive metadata
	// Note: this field may return null, indicating that no valid values can be obtained.
	HiveMetaDb *string `json:"HiveMetaDb,omitempty" name:"HiveMetaDb"`

	// 
	ServiceClass *string `json:"ServiceClass,omitempty" name:"ServiceClass"`
}

type CreateInstanceRequest struct {
	*tchttp.BaseRequest

	// Product ID. Different product IDs represent different EMR product versions. Valid values:
	// <li>1: EMR v1.3.1.</li>
	// <li>2: EMR v2.0.1.</li>
	// <li>4: EMR v2.1.0.</li>
	// <li>7: EMR v3.0.0.</li>
	ProductId *uint64 `json:"ProductId,omitempty" name:"ProductId"`

	// Configuration information of VPC. This parameter is used to specify the VPC ID, subnet ID, etc.
	VPCSettings *VPCSettings `json:"VPCSettings,omitempty" name:"VPCSettings"`

	// List of deployed components. Different `ProductIds` correspond to components on different versions. For example, when `ProductId` is 4, this parameter can be `Software.0=hadoop-2.8.4&Software.1=zookeeper-3.4.9`; when `ProductId` is 2, this parameter can be `Software.0=hadoop-2.7.3&Software.1=zookeeper-3.4.9`.
	Software []*string `json:"Software,omitempty" name:"Software" list`

	// Node resource specification.
	ResourceSpec *NewResourceSpec `json:"ResourceSpec,omitempty" name:"ResourceSpec"`

	// Whether to enable high node availability. Valid values:
	// <li>0: does not enable high availability of node.</li>
	// <li>1: enables high availability of node.</li>
	SupportHA *uint64 `json:"SupportHA,omitempty" name:"SupportHA"`

	// Instance name.
	// <li>Length limit: 6–36 characters.</li>
	// <li>Only letters, numbers, dashes (-), and underscores (_) are supported.</li>
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Instance billing mode. Valid values:
	// <li>0: pay-as-you-go.</li>
	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`

	// Instance location. This parameter is used to specify the AZ, project, and other attributes of the instance.
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// Purchase duration of instance, which needs to be used together with `TimeUnit`.
	// <li>When `PayMode` is 0, `TimeSpan` can only be 3,600.</li>
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// Time unit of instance purchase duration. Valid values:
	// <li>s: seconds. When `PayMode` is 0, `TimeUnit` can only be `s`.</li>
	// <li>m: month. When `PayMode` is 1, `TimeUnit` can only be `m`.</li>
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// Instance login settings. This parameter allows you to set the login password or key for your purchased node.
	// <li>If the key is set, the password will be only used for login to the native component WebUI.</li>
	// <li>If the key is not set, the password will be used for login to all purchased nodes and the native component WebUI.</li>
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// Parameter required for enabling COS access.
	COSSettings *COSSettings `json:"COSSettings,omitempty" name:"COSSettings"`

	// Security group to which an instance belongs in the format of `sg-xxxxxxxx`. This parameter can be obtained from the `SecurityGroupId` field in the return value of the [DescribeSecurityGroups](https://intl.cloud.tencent.com/document/api/215/15808) API.
	SgId *string `json:"SgId,omitempty" name:"SgId"`

	// Bootstrap script settings.
	PreExecutedFileSettings []*PreExecuteFileSettings `json:"PreExecutedFileSettings,omitempty" name:"PreExecutedFileSettings" list`

	// Whether auto-renewal is enabled. Valid values:
	// <li>0: auto-renewal not enabled.</li>
	// <li>1: auto-renewal enabled.</li>
	AutoRenew *uint64 `json:"AutoRenew,omitempty" name:"AutoRenew"`

	// Client token.
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// Whether to enable public IP access for master node. Valid values:
	// <li>NEED_MASTER_WAN: enables public IP for master node.</li>
	// <li>NOT_NEED_MASTER_WAN: does not enable.</li>Public IP is enabled for master node by default.
	NeedMasterWan *string `json:"NeedMasterWan,omitempty" name:"NeedMasterWan"`

	// Whether to enable remote public network login, i.e., port 22. When `SgId` is not empty, this parameter does not take effect.
	RemoteLoginAtCreate *int64 `json:"RemoteLoginAtCreate,omitempty" name:"RemoteLoginAtCreate"`

	// Whether to enable secure cluster. 0: no; other values: yes.
	CheckSecurity *int64 `json:"CheckSecurity,omitempty" name:"CheckSecurity"`

	// Accesses to external file system.
	ExtendFsField *string `json:"ExtendFsField,omitempty" name:"ExtendFsField"`

	// Tag description list. This parameter is used to bind a tag to a resource instance.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`

	// List of spread placement group IDs. Only one can be specified currently.
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitempty" name:"DisasterRecoverGroupIds" list`

	// CBS disk encryption at the cluster level. 0: not encrypted, 1: encrypted
	CbsEncrypt *uint64 `json:"CbsEncrypt,omitempty" name:"CbsEncrypt"`

	// Hive-shared metadatabase type. Valid values:
	// <li>EMR_DEFAULT_META: the cluster creates one by default.</li>
	// <li>EMR_EXIST_META: the cluster uses the specified EMR-MetaDB instance.</li>
	// <li>USER_CUSTOM_META: the cluster uses a custom MetaDB instance.</li>
	MetaType *string `json:"MetaType,omitempty" name:"MetaType"`

	// EMR-MetaDB instance
	UnifyMetaInstanceId *string `json:"UnifyMetaInstanceId,omitempty" name:"UnifyMetaInstanceId"`

	// Custom MetaDB instance information
	MetaDBInfo *CustomMetaInfo `json:"MetaDBInfo,omitempty" name:"MetaDBInfo"`
}

func (r *CreateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CustomMetaInfo struct {

	// JDBC connection to custom MetaDB instance beginning with `jdbc:mysql://`
	MetaDataJdbcUrl *string `json:"MetaDataJdbcUrl,omitempty" name:"MetaDataJdbcUrl"`

	// Custom MetaDB instance username
	MetaDataUser *string `json:"MetaDataUser,omitempty" name:"MetaDataUser"`

	// Custom MetaDB instance password
	MetaDataPass *string `json:"MetaDataPass,omitempty" name:"MetaDataPass"`
}

type DescribeClusterNodesRequest struct {
	*tchttp.BaseRequest

	// Cluster instance ID in the format of emr-xxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Node flag. Valid values:
	// <li>all: gets the information of nodes in all types except TencentDB information.</li>
	// <li>master: gets master node information.</li>
	// <li>core: gets core node information.</li>
	// <li>task: gets task node information.</li>
	// <li>common: gets common node information.</li>
	// <li>router: gets router node information.</li>
	// <li>db: gets TencentDB information in normal status.</li>
	// Note: only the above values are supported for the time being. Entering other values will cause errors.
	NodeFlag *string `json:"NodeFlag,omitempty" name:"NodeFlag"`

	// Page number. Default value: 0, indicating the first page.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results per page. Default value: 100. Maximum value: 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeClusterNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterNodesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterNodesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of queried nodes
		TotalCnt *int64 `json:"TotalCnt,omitempty" name:"TotalCnt"`

		// List of node details
	// Note: this field may return null, indicating that no valid values can be obtained.
		NodeList []*NodeHardwareInfo `json:"NodeList,omitempty" name:"NodeList" list`

		// List of tag keys owned by user
	// Note: this field may return null, indicating that no valid values can be obtained.
		TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterNodesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest

	// Cluster filtering policy. Valid values:
	// <li>clusterList: queries the list of clusters except terminated ones.</li>
	// <li>monitorManage: queries the list of clusters except those that have been terminated, are being created, or failed to be created.</li>
	// <li>cloudHardwareManage/componentManage: reserved fields with the same meaning as `monitorManage`.</li>
	DisplayStrategy *string `json:"DisplayStrategy,omitempty" name:"DisplayStrategy"`

	// Queries by one or more instance IDs in the format of `emr-xxxxxxxx`. For the format of this parameter, please see the `id.N` section in [API Overview](https://intl.cloud.tencent.com/document/api/213/15688). If no instance ID is entered, the list of all instances under this `APPID` will be returned.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// Page number. Default value: 0, indicating the first page.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results per page. Default value: 10. Maximum value: 100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// ID of the project to which the instance belongs. This parameter can be obtained from the `projectId` field in the return value of the `DescribeProject` API. If this value is -1, the list of all instances will be returned.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Sorting field. Valid values:
	// <li>clusterId: sorts by cluster ID.</li>
	// <li>addTime: sorts by instance creation time.</li>
	// <li>status: sorts by instance status code.</li>
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// Sorts according to `OrderField` in ascending or descending order. Valid values:
	// <li>0: descending order.</li>
	// <li>1: ascending order.</li>Default value: 0.
	Asc *int64 `json:"Asc,omitempty" name:"Asc"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible instances.
		TotalCnt *int64 `json:"TotalCnt,omitempty" name:"TotalCnt"`

		// List of EMR instance details.
	// Note: this field may return null, indicating that no valid values can be obtained.
		ClusterList []*ClusterInstancesInfo `json:"ClusterList,omitempty" name:"ClusterList" list`

		// List of tag keys associated to an instance.
	// Note: this field may return null, indicating that no valid values can be obtained.
		TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EmrProductConfigOutter struct {

	// Software information
	// Note: this field may return null, indicating that no valid values can be obtained.
	SoftInfo []*string `json:"SoftInfo,omitempty" name:"SoftInfo" list`

	// Number of master nodes
	// Note: this field may return null, indicating that no valid values can be obtained.
	MasterNodeSize *int64 `json:"MasterNodeSize,omitempty" name:"MasterNodeSize"`

	// Number of core nodes
	// Note: this field may return null, indicating that no valid values can be obtained.
	CoreNodeSize *int64 `json:"CoreNodeSize,omitempty" name:"CoreNodeSize"`

	// Number of task nodes
	// Note: this field may return null, indicating that no valid values can be obtained.
	TaskNodeSize *int64 `json:"TaskNodeSize,omitempty" name:"TaskNodeSize"`

	// Number of common nodes
	// Note: this field may return null, indicating that no valid values can be obtained.
	ComNodeSize *int64 `json:"ComNodeSize,omitempty" name:"ComNodeSize"`

	// Master node resource
	// Note: this field may return null, indicating that no valid values can be obtained.
	MasterResource *OutterResource `json:"MasterResource,omitempty" name:"MasterResource"`

	// Core node resource
	// Note: this field may return null, indicating that no valid values can be obtained.
	CoreResource *OutterResource `json:"CoreResource,omitempty" name:"CoreResource"`

	// Task node resource
	// Note: this field may return null, indicating that no valid values can be obtained.
	TaskResource *OutterResource `json:"TaskResource,omitempty" name:"TaskResource"`

	// Common node resource
	// Note: this field may return null, indicating that no valid values can be obtained.
	ComResource *OutterResource `json:"ComResource,omitempty" name:"ComResource"`

	// Whether COS is used
	// Note: this field may return null, indicating that no valid values can be obtained.
	OnCos *bool `json:"OnCos,omitempty" name:"OnCos"`

	// Billing mode
	// Note: this field may return null, indicating that no valid values can be obtained.
	ChargeType *int64 `json:"ChargeType,omitempty" name:"ChargeType"`

	// Number of router nodes
	// Note: this field may return null, indicating that no valid values can be obtained.
	RouterNodeSize *int64 `json:"RouterNodeSize,omitempty" name:"RouterNodeSize"`

	// Whether HA is supported
	// Note: this field may return null, indicating that no valid values can be obtained.
	SupportHA *bool `json:"SupportHA,omitempty" name:"SupportHA"`

	// Whether secure mode is supported
	// Note: this field may return null, indicating that no valid values can be obtained.
	SecurityOn *bool `json:"SecurityOn,omitempty" name:"SecurityOn"`

	// Security group name
	// Note: this field may return null, indicating that no valid values can be obtained.
	SecurityGroup *string `json:"SecurityGroup,omitempty" name:"SecurityGroup"`

	// Whether to enable CBS encryption
	// Note: this field may return null, indicating that no valid values can be obtained.
	CbsEncrypt *int64 `json:"CbsEncrypt,omitempty" name:"CbsEncrypt"`
}

type InquiryPriceCreateInstanceRequest struct {
	*tchttp.BaseRequest

	// Time unit of instance purchase duration. Valid values:
	// <li>s: seconds. When `PayMode` is 0, `TimeUnit` can only be `s`.</li>
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// Purchase duration of instance, which needs to be used together with `TimeUnit`.
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// Node specification queried for price.
	ResourceSpec *NewResourceSpec `json:"ResourceSpec,omitempty" name:"ResourceSpec"`

	// Currency.
	Currency *string `json:"Currency,omitempty" name:"Currency"`

	// Instance billing mode. Valid values:
	// <li>0: pay-as-you-go.</li>
	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`

	// Whether to enable high availability of node. Valid values:
	// <li>0: does not enable high availability of node.</li>
	// <li>1: enables high availability of node.</li>
	SupportHA *uint64 `json:"SupportHA,omitempty" name:"SupportHA"`

	// List of deployed components.
	Software []*string `json:"Software,omitempty" name:"Software" list`

	// Instance location. This parameter is used to specify the AZ, project, and other attributes of the instance.
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// Configuration information of VPC. This parameter is used to specify the VPC ID, subnet ID, etc.
	VPCSettings *VPCSettings `json:"VPCSettings,omitempty" name:"VPCSettings"`

	// Hive-shared metadatabase type. Valid values:
	// <li>EMR_DEFAULT_META: the cluster creates one by default.</li>
	// <li>EMR_EXIST_META: the cluster uses the specified EMR-MetaDB instance.</li>
	// <li>USER_CUSTOM_META: the cluster uses a custom MetaDB instance.</li>
	MetaType *string `json:"MetaType,omitempty" name:"MetaType"`

	// EMR-MetaDB instance
	UnifyMetaInstanceId *string `json:"UnifyMetaInstanceId,omitempty" name:"UnifyMetaInstanceId"`

	// Custom MetaDB instance information
	MetaDBInfo *CustomMetaInfo `json:"MetaDBInfo,omitempty" name:"MetaDBInfo"`

	// Product ID. Different product IDs represent different EMR product versions. Valid values:
	// <li>1: EMR v1.3.1.</li>
	// <li>2: EMR v2.0.1.</li>
	// <li>4: EMR v2.1.0.</li>
	// <li>7: EMR v3.0.0.</li>
	ProductId *uint64 `json:"ProductId,omitempty" name:"ProductId"`
}

func (r *InquiryPriceCreateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceCreateInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceCreateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Original price.
	// Note: this field may return null, indicating that no valid values can be obtained.
		OriginalCost *float64 `json:"OriginalCost,omitempty" name:"OriginalCost"`

		// Discounted price.
	// Note: this field may return null, indicating that no valid values can be obtained.
		DiscountCost *float64 `json:"DiscountCost,omitempty" name:"DiscountCost"`

		// Time unit of instance purchase duration. Valid values:
	// <li>s: seconds.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
		TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

		// Purchase duration of instance.
	// Note: this field may return null, indicating that no valid values can be obtained.
		TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceCreateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceCreateInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRenewInstanceRequest struct {
	*tchttp.BaseRequest

	// How long the instance will be renewed for, which needs to be used together with `TimeUnit`.
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// List of resource IDs of the node to be renewed. The resource ID is in the format of `emr-vm-xxxxxxxx`. A valid resource ID can be queried in the [console](https://console.cloud.tencent.com/emr/static/hardware).
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds" list`

	// Location of the instance. This parameter is used to specify the AZ, project, and other attributes of the instance.
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// Instance billing mode.
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// Unit of time for instance renewal.
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// Currency.
	Currency *string `json:"Currency,omitempty" name:"Currency"`
}

func (r *InquiryPriceRenewInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceRenewInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRenewInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Original price.
	// Note: this field may return null, indicating that no valid values can be obtained.
		OriginalCost *float64 `json:"OriginalCost,omitempty" name:"OriginalCost"`

		// Discounted price.
	// Note: this field may return null, indicating that no valid values can be obtained.
		DiscountCost *float64 `json:"DiscountCost,omitempty" name:"DiscountCost"`

		// Unit of time for instance renewal.
	// Note: this field may return null, indicating that no valid values can be obtained.
		TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

		// How long the instance will be renewed for.
	// Note: this field may return null, indicating that no valid values can be obtained.
		TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceRenewInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceRenewInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceScaleOutInstanceRequest struct {
	*tchttp.BaseRequest

	// Time unit of scale-out. Valid values:
	// <li>s: seconds. When `PayMode` is 0, `TimeUnit` can only be `s`.</li>
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// Duration of scale-out, which needs to be used together with `TimeUnit`.
	// <li>When `PayMode` is 0, `TimeSpan` can only be 3,600.</li>
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// ID of the AZ where an instance resides.
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// Instance billing mode. Valid values:
	// <li>0: pay-as-you-go.</li>
	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Number of core nodes added.
	CoreCount *uint64 `json:"CoreCount,omitempty" name:"CoreCount"`

	// Number of task nodes added.
	TaskCount *uint64 `json:"TaskCount,omitempty" name:"TaskCount"`

	// Currency.
	Currency *string `json:"Currency,omitempty" name:"Currency"`

	// Number of router nodes added.
	RouterCount *uint64 `json:"RouterCount,omitempty" name:"RouterCount"`
}

func (r *InquiryPriceScaleOutInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceScaleOutInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceScaleOutInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Original price.
	// Note: this field may return null, indicating that no valid values can be obtained.
		OriginalCost *string `json:"OriginalCost,omitempty" name:"OriginalCost"`

		// Discounted price.
	// Note: this field may return null, indicating that no valid values can be obtained.
		DiscountCost *string `json:"DiscountCost,omitempty" name:"DiscountCost"`

		// Time unit of scale-out. Valid values:
	// <li>s: seconds.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
		Unit *string `json:"Unit,omitempty" name:"Unit"`

		// Node specification queried for price.
	// Note: this field may return null, indicating that no valid values can be obtained.
		PriceSpec *PriceResource `json:"PriceSpec,omitempty" name:"PriceSpec"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceScaleOutInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceScaleOutInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceUpdateInstanceRequest struct {
	*tchttp.BaseRequest

	// Time unit of scaling. Valid values:
	// <li>s: seconds. When `PayMode` is 0, `TimeUnit` can only be `s`.</li>
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// Duration of scaling, which needs to be used together with `TimeUnit`.
	// <li>When `PayMode` is 0, `TimeSpan` can only be 3,600.</li>
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// Target node specification.
	UpdateSpec *UpdateInstanceSettings `json:"UpdateSpec,omitempty" name:"UpdateSpec"`

	// Instance billing mode. Valid values:
	// <li>0: pay-as-you-go.</li>
	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`

	// Instance location. This parameter is used to specify the AZ, project, and other attributes of the instance.
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// Currency.
	Currency *string `json:"Currency,omitempty" name:"Currency"`
}

func (r *InquiryPriceUpdateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceUpdateInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceUpdateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Original price.
	// Note: this field may return null, indicating that no valid values can be obtained.
		OriginalCost *float64 `json:"OriginalCost,omitempty" name:"OriginalCost"`

		// Discounted price.
	// Note: this field may return null, indicating that no valid values can be obtained.
		DiscountCost *float64 `json:"DiscountCost,omitempty" name:"DiscountCost"`

		// Time unit of scaling. Valid values:
	// <li>s: seconds.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
		TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

		// Duration of scaling.
	// Note: this field may return null, indicating that no valid values can be obtained.
		TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceUpdateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceUpdateInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LoginSettings struct {

	// Password
	Password *string `json:"Password,omitempty" name:"Password"`

	// Public Key
	PublicKeyId *string `json:"PublicKeyId,omitempty" name:"PublicKeyId"`
}

type MultiDisk struct {

	// Cloud disk type. Valid values: CLOUD_PREMIUM, CLOUD_SSD, CLOUD_BASIC
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// Cloud disk size
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`

	// Number of cloud disks of this type
	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type MultiDiskMC struct {

	// Number of cloud disks in this type
	// Note: this field may return null, indicating that no valid values can be obtained.
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// Disk type
	// Note: this field may return null, indicating that no valid values can be obtained.
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// Cloud disk size
	// Note: this field may return null, indicating that no valid values can be obtained.
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`
}

type NewResourceSpec struct {

	// Describes master node resource
	MasterResourceSpec *Resource `json:"MasterResourceSpec,omitempty" name:"MasterResourceSpec"`

	// Describes core node resource
	CoreResourceSpec *Resource `json:"CoreResourceSpec,omitempty" name:"CoreResourceSpec"`

	// Describes task node resource
	TaskResourceSpec *Resource `json:"TaskResourceSpec,omitempty" name:"TaskResourceSpec"`

	// Number of master nodes
	MasterCount *int64 `json:"MasterCount,omitempty" name:"MasterCount"`

	// Number of core nodes
	CoreCount *int64 `json:"CoreCount,omitempty" name:"CoreCount"`

	// Number of task nodes
	TaskCount *int64 `json:"TaskCount,omitempty" name:"TaskCount"`

	// Describes common node resource
	CommonResourceSpec *Resource `json:"CommonResourceSpec,omitempty" name:"CommonResourceSpec"`

	// Number of common nodes
	CommonCount *int64 `json:"CommonCount,omitempty" name:"CommonCount"`
}

type NodeHardwareInfo struct {

	// User `APPID`
	// Note: this field may return null, indicating that no valid values can be obtained.
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// Serial number
	// Note: this field may return null, indicating that no valid values can be obtained.
	SerialNo *string `json:"SerialNo,omitempty" name:"SerialNo"`

	// Machine instance ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	OrderNo *string `json:"OrderNo,omitempty" name:"OrderNo"`

	// Public IP bound to master node
	// Note: this field may return null, indicating that no valid values can be obtained.
	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`

	// Node type
	// Note: this field may return null, indicating that no valid values can be obtained.
	Flag *int64 `json:"Flag,omitempty" name:"Flag"`

	// Node specification
	// Note: this field may return null, indicating that no valid values can be obtained.
	Spec *string `json:"Spec,omitempty" name:"Spec"`

	// Number of node cores
	// Note: this field may return null, indicating that no valid values can be obtained.
	CpuNum *int64 `json:"CpuNum,omitempty" name:"CpuNum"`

	// Node memory size
	// Note: this field may return null, indicating that no valid values can be obtained.
	MemSize *int64 `json:"MemSize,omitempty" name:"MemSize"`

	// Node memory description
	// Note: this field may return null, indicating that no valid values can be obtained.
	MemDesc *string `json:"MemDesc,omitempty" name:"MemDesc"`

	// Node region
	// Note: this field may return null, indicating that no valid values can be obtained.
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// Node AZ
	// Note: this field may return null, indicating that no valid values can be obtained.
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// Application time
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApplyTime *string `json:"ApplyTime,omitempty" name:"ApplyTime"`

	// Release time
	// Note: this field may return null, indicating that no valid values can be obtained.
	FreeTime *string `json:"FreeTime,omitempty" name:"FreeTime"`

	// Disk size
	// Note: this field may return null, indicating that no valid values can be obtained.
	DiskSize *string `json:"DiskSize,omitempty" name:"DiskSize"`

	// Node description
	// Note: this field may return null, indicating that no valid values can be obtained.
	NameTag *string `json:"NameTag,omitempty" name:"NameTag"`

	// Services deployed on node
	// Note: this field may return null, indicating that no valid values can be obtained.
	Services *string `json:"Services,omitempty" name:"Services"`

	// Disk type
	// Note: this field may return null, indicating that no valid values can be obtained.
	StorageType *int64 `json:"StorageType,omitempty" name:"StorageType"`

	// System disk size
	// Note: this field may return null, indicating that no valid values can be obtained.
	RootSize *int64 `json:"RootSize,omitempty" name:"RootSize"`

	// Payment type
	// Note: this field may return null, indicating that no valid values can be obtained.
	ChargeType *int64 `json:"ChargeType,omitempty" name:"ChargeType"`

	// Database IP
	// Note: this field may return null, indicating that no valid values can be obtained.
	CdbIp *string `json:"CdbIp,omitempty" name:"CdbIp"`

	// Database port
	// Note: this field may return null, indicating that no valid values can be obtained.
	CdbPort *int64 `json:"CdbPort,omitempty" name:"CdbPort"`

	// Disk capacity
	// Note: this field may return null, indicating that no valid values can be obtained.
	HwDiskSize *int64 `json:"HwDiskSize,omitempty" name:"HwDiskSize"`

	// Disk capacity description
	// Note: this field may return null, indicating that no valid values can be obtained.
	HwDiskSizeDesc *string `json:"HwDiskSizeDesc,omitempty" name:"HwDiskSizeDesc"`

	// Memory capacity
	// Note: this field may return null, indicating that no valid values can be obtained.
	HwMemSize *int64 `json:"HwMemSize,omitempty" name:"HwMemSize"`

	// Memory capacity description
	// Note: this field may return null, indicating that no valid values can be obtained.
	HwMemSizeDesc *string `json:"HwMemSizeDesc,omitempty" name:"HwMemSizeDesc"`

	// Expiration time
	// Note: this field may return null, indicating that no valid values can be obtained.
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// Node resource ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	EmrResourceId *string `json:"EmrResourceId,omitempty" name:"EmrResourceId"`

	// Renewal flag
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsAutoRenew *int64 `json:"IsAutoRenew,omitempty" name:"IsAutoRenew"`

	// Device flag
	// Note: this field may return null, indicating that no valid values can be obtained.
	DeviceClass *string `json:"DeviceClass,omitempty" name:"DeviceClass"`

	// Support for configuration adjustment
	// Note: this field may return null, indicating that no valid values can be obtained.
	Mutable *int64 `json:"Mutable,omitempty" name:"Mutable"`

	// Multi-cloud disk
	// Note: this field may return null, indicating that no valid values can be obtained.
	MCMultiDisk []*MultiDiskMC `json:"MCMultiDisk,omitempty" name:"MCMultiDisk" list`

	// Database information
	// Note: this field may return null, indicating that no valid values can be obtained.
	CdbNodeInfo *CdbInfo `json:"CdbNodeInfo,omitempty" name:"CdbNodeInfo"`

	// Private IP
	// Note: this field may return null, indicating that no valid values can be obtained.
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Whether this node can be terminated. 1: yes, 0: no
	// Note: this field may return null, indicating that no valid values can be obtained.
	Destroyable *int64 `json:"Destroyable,omitempty" name:"Destroyable"`

	// Tags bound to node
	// Note: this field may return null, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`

	// 
	AutoFlag *int64 `json:"AutoFlag,omitempty" name:"AutoFlag"`
}

type OutterResource struct {

	// Specification
	// Note: this field may return null, indicating that no valid values can be obtained.
	Spec *string `json:"Spec,omitempty" name:"Spec"`

	// Specification name
	// Note: this field may return null, indicating that no valid values can be obtained.
	SpecName *string `json:"SpecName,omitempty" name:"SpecName"`

	// Disk type
	// Note: this field may return null, indicating that no valid values can be obtained.
	StorageType *int64 `json:"StorageType,omitempty" name:"StorageType"`

	// Disk type
	// Note: this field may return null, indicating that no valid values can be obtained.
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// System disk size
	// Note: this field may return null, indicating that no valid values can be obtained.
	RootSize *int64 `json:"RootSize,omitempty" name:"RootSize"`

	// Memory size
	// Note: this field may return null, indicating that no valid values can be obtained.
	MemSize *int64 `json:"MemSize,omitempty" name:"MemSize"`

	// Number of CPUs
	// Note: this field may return null, indicating that no valid values can be obtained.
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// Disk size
	// Note: this field may return null, indicating that no valid values can be obtained.
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// Specification
	// Note: this field may return null, indicating that no valid values can be obtained.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
}

type Placement struct {

	// ID of the project to which the instance belongs. You can call the `DescribeProject` API and see the `projectId` field in the response to get the value of this parameter. If it is left empty, the default project will be used.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// AZ where the instance resides, such as ap-guangzhou-1. You can call the `DescribeZones` API and see the `Zone` field to get the value of this parameter.
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type PreExecuteFileSettings struct {

	// COS path to script, which has been disused
	Path *string `json:"Path,omitempty" name:"Path"`

	// Execution script parameter
	Args []*string `json:"Args,omitempty" name:"Args" list`

	// COS bucket name, which has been disused
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// COS region name, which has been disused
	Region *string `json:"Region,omitempty" name:"Region"`

	// COS domain data, which has been disused
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Execution sequence
	RunOrder *int64 `json:"RunOrder,omitempty" name:"RunOrder"`

	// `resourceAfter` or `clusterAfter`
	WhenRun *string `json:"WhenRun,omitempty" name:"WhenRun"`

	// Script name, which has been disused
	CosFileName *string `json:"CosFileName,omitempty" name:"CosFileName"`

	// COS address of script
	CosFileURI *string `json:"CosFileURI,omitempty" name:"CosFileURI"`

	// COS `SecretId`
	CosSecretId *string `json:"CosSecretId,omitempty" name:"CosSecretId"`

	// COS `SecretKey`
	CosSecretKey *string `json:"CosSecretKey,omitempty" name:"CosSecretKey"`

	// COS `appid`, which has been disused
	AppId *string `json:"AppId,omitempty" name:"AppId"`
}

type PriceResource struct {

	// Target specification
	// Note: this field may return null, indicating that no valid values can be obtained.
	Spec *string `json:"Spec,omitempty" name:"Spec"`

	// Disk type
	// Note: this field may return null, indicating that no valid values can be obtained.
	StorageType *uint64 `json:"StorageType,omitempty" name:"StorageType"`

	// Disk type
	// Note: this field may return null, indicating that no valid values can be obtained.
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// System disk size
	// Note: this field may return null, indicating that no valid values can be obtained.
	RootSize *int64 `json:"RootSize,omitempty" name:"RootSize"`

	// Memory size
	// Note: this field may return null, indicating that no valid values can be obtained.
	MemSize *int64 `json:"MemSize,omitempty" name:"MemSize"`

	// Number of cores
	// Note: this field may return null, indicating that no valid values can be obtained.
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// Disk size
	// Note: this field may return null, indicating that no valid values can be obtained.
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// List of cloud disks
	// Note: this field may return null, indicating that no valid values can be obtained.
	MultiDisks []*MultiDisk `json:"MultiDisks,omitempty" name:"MultiDisks" list`

	// Number of disks
	// Note: this field may return null, indicating that no valid values can be obtained.
	DiskCnt *int64 `json:"DiskCnt,omitempty" name:"DiskCnt"`

	// Specification
	// Note: this field may return null, indicating that no valid values can be obtained.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Tag
	// Note: this field may return null, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`
}

type Resource struct {

	// Node specification description
	// Note: this field may return null, indicating that no valid values can be obtained.
	Spec *string `json:"Spec,omitempty" name:"Spec"`

	// Storage class
	// Note: this field may return null, indicating that no valid values can be obtained.
	StorageType *int64 `json:"StorageType,omitempty" name:"StorageType"`

	// Disk type
	// Note: this field may return null, indicating that no valid values can be obtained.
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// Memory capacity in MB
	// Note: this field may return null, indicating that no valid values can be obtained.
	MemSize *int64 `json:"MemSize,omitempty" name:"MemSize"`

	// Number of CPU cores
	// Note: this field may return null, indicating that no valid values can be obtained.
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// Data disk capacity
	// Note: this field may return null, indicating that no valid values can be obtained.
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// System disk capacity
	// Note: this field may return null, indicating that no valid values can be obtained.
	RootSize *int64 `json:"RootSize,omitempty" name:"RootSize"`

	// List of cloud disks. When the data disk is a cloud disk, `DiskType` and `DiskSize` are used directly; `MultiDisks` will be used for the excessive part
	// Note: this field may return null, indicating that no valid values can be obtained.
	MultiDisks []*MultiDisk `json:"MultiDisks,omitempty" name:"MultiDisks" list`

	// List of tags to be bound
	// Note: this field may return null, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`

	// Specification type
	// Note: this field may return null, indicating that no valid values can be obtained.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Number of local disks
	// Note: this field may return null, indicating that no valid values can be obtained.
	LocalDiskNum *uint64 `json:"LocalDiskNum,omitempty" name:"LocalDiskNum"`

	// Number of disks
	// Note: this field may return null, indicating that no valid values can be obtained.
	DiskNum *uint64 `json:"DiskNum,omitempty" name:"DiskNum"`
}

type ScaleOutInstanceRequest struct {
	*tchttp.BaseRequest

	// Time unit of scale-out. Valid values:
	// <li>s: seconds. When `PayMode` is 0, `TimeUnit` can only be `s`.</li>
	// <li>m: month. When `PayMode` is 1, `TimeUnit` can only be `m`.</li>
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// Duration of scale-out, which needs to be used together with `TimeUnit`.
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance billing mode. Valid values:
	// <li>0: pay-as-you-go.</li>
	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`

	// Client token.
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// Bootstrap script settings.
	PreExecutedFileSettings []*PreExecuteFileSettings `json:"PreExecutedFileSettings,omitempty" name:"PreExecutedFileSettings" list`

	// Number of task nodes added.
	TaskCount *uint64 `json:"TaskCount,omitempty" name:"TaskCount"`

	// Number of core nodes added.
	CoreCount *uint64 `json:"CoreCount,omitempty" name:"CoreCount"`

	// Process not required during scale-out.
	UnNecessaryNodeList []*uint64 `json:"UnNecessaryNodeList,omitempty" name:"UnNecessaryNodeList" list`

	// Number of router nodes added.
	RouterCount *uint64 `json:"RouterCount,omitempty" name:"RouterCount"`

	// Deployed service.
	// <li>`SoftDeployInfo` and `ServiceNodeInfo` are in the same group and mutually exclusive with `UnNecessaryNodeList`.</li>
	// <li>The combination of `SoftDeployInfo` and `ServiceNodeInfo` is recommended.</li>
	SoftDeployInfo []*uint64 `json:"SoftDeployInfo,omitempty" name:"SoftDeployInfo" list`

	// Started process.
	ServiceNodeInfo []*uint64 `json:"ServiceNodeInfo,omitempty" name:"ServiceNodeInfo" list`

	// List of spread placement group IDs. Only one can be specified currently.
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitempty" name:"DisasterRecoverGroupIds" list`

	// List of tags bound to added nodes.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`
}

func (r *ScaleOutInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ScaleOutInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ScaleOutInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Instance ID.
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// Order number.
	// Note: this field may return null, indicating that no valid values can be obtained.
		DealNames []*string `json:"DealNames,omitempty" name:"DealNames" list`

		// Client token.
	// Note: this field may return null, indicating that no valid values can be obtained.
		ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

		// Scaling workflow ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// Big order number.
	// Note: this field may return null, indicating that no valid values can be obtained.
		BillId *string `json:"BillId,omitempty" name:"BillId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ScaleOutInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ScaleOutInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Tag struct {

	// Tag key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type TerminateInstanceRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// ID of terminated node. This parameter is reserved and does not need to be configured.
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds" list`
}

func (r *TerminateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TerminateInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TerminateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TerminateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TerminateInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TerminateTasksRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// List of resource IDs of the node to be terminated. The resource ID is in the format of `emr-vm-xxxxxxxx`. A valid resource ID can be queried in the [console](https://console.cloud.tencent.com/emr/static/hardware).
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds" list`
}

func (r *TerminateTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TerminateTasksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TerminateTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TerminateTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TerminateTasksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateInstanceSettings struct {

	// Memory capacity in GB
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// Number of CPU cores
	CPUCores *uint64 `json:"CPUCores,omitempty" name:"CPUCores"`

	// Machine resource ID (EMR resource identifier)
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// Target machine specification
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
}

type VPCSettings struct {

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}
